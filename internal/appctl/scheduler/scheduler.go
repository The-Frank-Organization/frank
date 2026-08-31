// Package scheduler owns serialized run and turn admission decisions.
package scheduler

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strings"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/settle"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

type CreateAuthIDFunc func() (string, error)

type Scheduler struct {
	applier      *applier.Host
	createAuthID CreateAuthIDFunc
}

func New(host *applier.Host, sources ...CreateAuthIDFunc) *Scheduler {
	source := mintCreateAuthID
	if len(sources) == 1 && sources[0] != nil {
		source = sources[0]
	}
	return &Scheduler{applier: host, createAuthID: source}
}

func mintCreateAuthID() (string, error) {
	value := make([]byte, 16)
	if _, err := rand.Read(value); err != nil {
		return "", err
	}
	return hex.EncodeToString(value), nil
}

type AdmissionDecision string

const (
	AdmissionCommitted      AdmissionDecision = "committed"
	AdmissionTaskOverflow   AdmissionDecision = "task_input_frame_overflow"
	AdmissionResumeTerminal AdmissionDecision = "resume_frame_overflow"
)

type AdmitRequest struct {
	RunID, TurnID, GenerationID, TurnEpoch string
	AdmissionRef                           appipc.AdmissionRef
	PredecessorTurnID                      string
	SessionLogPath                         string
	ConnectorReady                         bool
	EncodedSize                            int
	AdmissionRefEncodedSize                int
	At                                     int64
}

type AdmitResult struct {
	Decision AdmissionDecision
	Body     appipc.TurnOpenBody
}

type admitEvent struct {
	request      AdmitRequest
	createAuthID CreateAuthIDFunc
}

func (event admitEvent) RunID() string { return event.request.RunID }

func (scheduler *Scheduler) Admit(ctx context.Context, request AdmitRequest) (AdmitResult, error) {
	if scheduler == nil || scheduler.applier == nil || request.RunID == "" || request.TurnID == "" || request.GenerationID == "" || request.TurnEpoch == "" {
		return AdmitResult{}, errors.New("scheduler: invalid admission")
	}
	if err := validateAdmission(request); err != nil {
		return AdmitResult{}, err
	}
	result, err := scheduler.applier.Apply(ctx, admitEvent{request: request, createAuthID: scheduler.createAuthID})
	if err != nil {
		return AdmitResult{}, err
	}
	return result.Value.(AdmitResult), nil
}

func validateAdmission(request AdmitRequest) error {
	if request.PredecessorTurnID != "" && request.SessionLogPath == "" {
		return errors.New("scheduler: admission surface is incomplete")
	}
	if _, err := appipc.ParseCounter(request.TurnEpoch); err != nil {
		return err
	}
	switch request.AdmissionRef.Kind {
	case appipc.AdmissionWakeRelay:
		if request.AdmissionRef.RelayID == nil || *request.AdmissionRef.RelayID == "" || request.AdmissionRef.TaskInput != nil {
			return errors.New("scheduler: invalid wake admission reference")
		}
	case appipc.AdmissionOperatorInput:
		if request.AdmissionRef.TaskInput == nil || request.AdmissionRef.RelayID != nil {
			return errors.New("scheduler: invalid operator admission reference")
		}
	default:
		return errors.New("scheduler: unknown admission reference")
	}
	return nil
}

func (event admitEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	request := event.request
	continuation := request.PredecessorTurnID != ""
	if !request.ConnectorReady {
		return applier.Result{}, errors.New("scheduler: connector is not READY")
	}
	var runState, runPhase, runSessionLogPath, currentEpoch, workerState string
	var attachResult *string
	err := tx.QueryRowContext(ctx, `SELECT r.state,r.run_phase,r.session_log_path,e.turn_epoch,w.state,w.attach_result FROM runs r JOIN epochs e ON e.run_id=r.run_id JOIN workers w ON w.run_id=r.run_id AND w.generation_id=? WHERE r.run_id=?`, request.GenerationID, request.RunID).Scan(&runState, &runPhase, &runSessionLogPath, &currentEpoch, &workerState, &attachResult)
	if err != nil {
		return applier.Result{}, err
	}
	if runState != "ADMITTED" && runState != "ACTIVE" {
		return applier.Result{}, errors.New("scheduler: run is terminal")
	}
	if unpad(currentEpoch) != request.TurnEpoch || workerState != "LEASED" || attachResult == nil || *attachResult != appipc.AttachOK {
		return applier.Result{}, errors.New("scheduler: worker tuple is not admissible")
	}
	var workerLeases, activeTurns int
	if err := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM leases WHERE run_id=? AND lease_kind='worker' AND generation_id=? AND turn_epoch=? AND state='ACTIVE'`, request.RunID, request.GenerationID, currentEpoch).Scan(&workerLeases); err != nil {
		return applier.Result{}, err
	}
	if err := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM turns WHERE run_id=? AND state IN ('ADMITTED','ACTIVE')`, request.RunID).Scan(&activeTurns); err != nil {
		return applier.Result{}, err
	}
	if workerLeases != 1 || activeTurns != 0 {
		return applier.Result{}, errors.New("scheduler: lease or active-turn gate failed")
	}
	if (!continuation && runPhase != "created") || (continuation && runPhase != "create_authorized" && runPhase != "established") {
		return applier.Result{}, errors.New("scheduler: run phase is not admissible")
	}
	if !continuation && runSessionLogPath == "" {
		return applier.Result{}, errors.New("scheduler: initial session log path is absent")
	}
	encodedRef, err := appipc.MarshalJCS(request.AdmissionRef)
	if err != nil {
		return applier.Result{}, err
	}
	var predecessor any
	var snapshot any
	var settlement *appipc.SettlementManifest
	runDisposition := appipc.RunDispositionFresh
	createAuthID := ""
	sessionLogPath := runSessionLogPath
	if continuation {
		runDisposition = appipc.RunDispositionResume
		sessionLogPath = request.SessionLogPath
		predecessor = request.PredecessorTurnID
		var inherited []byte
		var predecessorState string
		if err := tx.QueryRowContext(ctx, `SELECT admission_ref,state,create_auth_id FROM turns WHERE run_id=? AND turn_id=?`, request.RunID, request.PredecessorTurnID).Scan(&inherited, &predecessorState, &createAuthID); err != nil {
			return applier.Result{}, err
		}
		if predecessorState != "INTERRUPTED" {
			return applier.Result{}, errors.New("scheduler: continuation predecessor is not interrupted")
		}
		var inheritedRef appipc.AdmissionRef
		if err := json.Unmarshal(inherited, &inheritedRef); err != nil {
			return applier.Result{}, err
		}
		request.AdmissionRef = inheritedRef
		encodedRef = inherited
		produced, _, err := (settle.Producer{}).Produce(ctx, tx, request.RunID, request.PredecessorTurnID, request.TurnID)
		if err != nil {
			return applier.Result{}, err
		}
		settlement = &produced
		snapshot, err = appipc.MarshalJCS(map[string]any{"session_log_path": sessionLogPath, "settlement_manifest": produced})
		if err != nil {
			return applier.Result{}, err
		}
	} else {
		if event.createAuthID == nil {
			return applier.Result{}, errors.New("scheduler: create authorization mint is unavailable")
		}
		createAuthID, err = event.createAuthID()
		if err != nil || len(createAuthID) != 32 {
			return applier.Result{}, errors.New("scheduler: create authorization mint failed")
		}
		if _, err := hex.DecodeString(createAuthID); err != nil || createAuthID != strings.ToLower(createAuthID) {
			return applier.Result{}, errors.New("scheduler: create authorization mint is non-canonical")
		}
		var collisions int
		if err := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM turns WHERE create_auth_id=?`, createAuthID).Scan(&collisions); err != nil {
			return applier.Result{}, err
		}
		if collisions != 0 {
			return applier.Result{}, errors.New("scheduler: create authorization mint collided")
		}
	}
	parked, err := snapshotParked(ctx, tx, request.RunID)
	if err != nil {
		return applier.Result{}, err
	}
	body := appipc.TurnOpenBody{TurnID: request.TurnID, AdmissionRef: request.AdmissionRef, ParkedUnknown: parked, RunDisposition: runDisposition, CreateAuthID: createAuthID, SessionLogPath: sessionLogPath}
	if continuation {
		body.PredecessorTurnID = &request.PredecessorTurnID
		body.SettlementManifest = settlement
	}
	actualSize, err := encodedTurnOpenSize(request.RunID, request.TurnEpoch, body)
	if err != nil {
		return applier.Result{}, err
	}
	effectiveSize := actualSize
	if request.EncodedSize > effectiveSize {
		effectiveSize = request.EncodedSize
	}
	refSize := len(encodedRef)
	if request.AdmissionRefEncodedSize > refSize {
		refSize = request.AdmissionRefEncodedSize
	}
	if (!continuation && refSize > appipc.AdmissionRefEncMax) || appipc.ClassifyTurnOpenSize(effectiveSize, continuation) == appipc.TurnOpenTaskInputOverflow {
		refusal := appipc.AdmissionRefusedBody{Reason: string(AdmissionTaskOverflow)}
		return applier.Result{Value: AdmitResult{Decision: AdmissionTaskOverflow}, Emissions: []applier.Emission{{Kind: "admission_refused", Value: refusal}}, NoMutation: true}, nil
	}
	if appipc.ClassifyTurnOpenSize(effectiveSize, continuation) == appipc.TurnOpenResumeOverflow {
		_, err := tx.ExecContext(ctx, `UPDATE runs SET state='FAILED',stop_reason='resume_frame_overflow',resume_action='operator_new_run',updated_at=? WHERE run_id=?`, request.At, request.RunID)
		return applier.Result{Value: AdmitResult{Decision: AdmissionResumeTerminal}}, err
	}
	if !continuation {
		result, err := tx.ExecContext(ctx, `UPDATE runs SET run_phase='create_authorized',updated_at=? WHERE run_id=? AND run_phase='created'`, request.At, request.RunID)
		if err != nil {
			return applier.Result{}, err
		}
		changed, _ := result.RowsAffected()
		if changed != 1 {
			return applier.Result{}, errors.New("scheduler: create authorization was already consumed")
		}
	}
	_, err = tx.ExecContext(ctx, `INSERT INTO turns(turn_id,run_id,turn_epoch,state,admission_ref,run_disposition,create_auth_id,predecessor_turn_id,resume_snapshot,resume_disposition,created_at) VALUES(?,?,?,?,?,?,?,?,?,?,?)`, request.TurnID, request.RunID, currentEpoch, "ACTIVE", encodedRef, runDisposition, createAuthID, predecessor, snapshot, "PENDING", request.At)
	if err != nil {
		return applier.Result{}, err
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO leases(run_id,lease_kind,lease_id,generation_id,turn_epoch,state,granted_at) VALUES(?,?,?,?,?,'ACTIVE',?)
		ON CONFLICT(run_id,lease_kind) DO UPDATE SET lease_id=excluded.lease_id,generation_id=excluded.generation_id,turn_epoch=excluded.turn_epoch,state='ACTIVE',granted_at=excluded.granted_at,released_at=NULL`, request.RunID, "turn", "turn:"+request.TurnID, request.GenerationID, currentEpoch, request.At); err != nil {
		return applier.Result{}, err
	}
	if !continuation && request.AdmissionRef.Kind == appipc.AdmissionWakeRelay {
		if request.AdmissionRef.RelayID == nil {
			return applier.Result{}, errors.New("scheduler: wake admission lacks relay")
		}
		result, err := tx.ExecContext(ctx, `UPDATE wake_schedule SET disposition='ADMITTED',admitted_turn_id=? WHERE relay_id=? AND run_id=? AND disposition='PENDING'`, request.TurnID, *request.AdmissionRef.RelayID, request.RunID)
		if err != nil {
			return applier.Result{}, err
		}
		changed, _ := result.RowsAffected()
		if changed != 1 {
			return applier.Result{}, errors.New("scheduler: wake is absent or already consumed")
		}
	}
	for _, row := range parked {
		if _, err := tx.ExecContext(ctx, `INSERT INTO turn_disclosures(disclosing_turn_id,source_turn_id,tool_call_id,ticket_id,state_at_disclosure,canonical_tool_name,canonical_args_digest) VALUES(?,?,?,?,?,?,?)`, request.TurnID, row.TurnID, row.ToolCallID, row.TicketID, row.State, row.CanonicalToolName, row.CanonicalArgsDigest); err != nil {
			return applier.Result{}, err
		}
	}
	return applier.Result{Value: AdmitResult{Decision: AdmissionCommitted, Body: body}, Emissions: []applier.Emission{{Kind: "turn_open", Value: body}}}, nil
}

func encodedTurnOpenSize(runID, epoch string, body appipc.TurnOpenBody) (int, error) {
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		return 0, err
	}
	payload, err := registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlW, Type: "turn_open", Seq: "0", RunID: &runID, TurnEpoch: &epoch, Body: body})
	if err != nil {
		return 0, err
	}
	return len(payload), nil
}

func snapshotParked(ctx context.Context, tx *store.Tx, runID string) ([]appipc.ParkedUnknown, error) {
	rows, err := tx.QueryContext(ctx, `SELECT c.turn_id,c.tool_call_id,a.ticket_id,c.state,c.canonical_tool_name,c.canonical_args_digest FROM tool_calls c JOIN tool_authorizations a ON a.run_id=c.run_id AND a.turn_id=c.turn_id AND a.tool_call_id=c.tool_call_id WHERE c.run_id=? AND c.state IN ('UNKNOWN_TOOL_OUTCOME','PARTIAL_TOOL_EFFECT') ORDER BY c.tool_call_id,a.ticket_id`, runID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]appipc.ParkedUnknown, 0)
	for rows.Next() {
		var row appipc.ParkedUnknown
		if err := rows.Scan(&row.TurnID, &row.ToolCallID, &row.TicketID, &row.State, &row.CanonicalToolName, &row.CanonicalArgsDigest); err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, rows.Err()
}

func unpad(value string) string {
	result, _ := appipc.UnpadCounter(value)
	return result
}
