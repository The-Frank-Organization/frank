// Package resume reconciles m-10 settlement evidence with the worker journal.
// It owns neither the settlement rows nor their persistence.
package resume

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/jackli/frank/internal/worker/jcs"
	"github.com/jackli/frank/internal/worker/journal"
	"github.com/jackli/frank/internal/worker/wire"
)

type EffectKind string

const (
	KindTool     EffectKind = "tool"
	KindProvider EffectKind = "provider"
)

type SettlementClass string

const (
	SettledWithContent  SettlementClass = "settled_with_content"
	DeterminateNoResume SettlementClass = "determinate_no_resume"
	Uncertain           SettlementClass = "uncertain"
)

const (
	DispositionResumable = "resumable"
	DispositionDegraded  = "degraded"
	ResumeActionReDerive = "re_derive"
)

type FirstAction string

const (
	ContinueModelLoop        FirstAction = "continue_model_loop"
	SurfaceAndTerminalize    FirstAction = "surface_and_terminalize"
	SurfaceUncertainTool     FirstAction = "surface_uncertain_tool"
	SurfaceUncertainProvider FirstAction = "surface_uncertain_provider"
	ReDerive                 FirstAction = "re_derive"
)

type SettlementEntry struct {
	Kind       EffectKind
	Class      SettlementClass
	RunID      string
	TurnID     string
	EffectID   string
	ArgsDigest string
	Terminal   string
	SettledAt  time.Time
}

type SettlementManifest struct {
	Version           int                 `json:"settlement_manifest_v"`
	RunID             string              `json:"run_id"`
	ProducedForTurnID string              `json:"produced_for_turn_id"`
	Entries           []ManifestEntryWire `json:"entries"`
}

type ManifestEntryWire struct {
	Kind        EffectKind      `json:"kind"`
	Class       SettlementClass `json:"class"`
	RunID       string          `json:"run_id"`
	TurnID      string          `json:"turn_id"`
	ToolCallID  string          `json:"tool_call_id,omitempty"`
	ArgsDigest  string          `json:"args_digest,omitempty"`
	AttemptID   string          `json:"attempt_id,omitempty"`
	Terminal    string          `json:"terminal"`
	VoidReason  string          `json:"void_reason,omitempty"`
	CancelPoint string          `json:"cancel_point,omitempty"`
}

// DecodeSettlementManifest consumes m-10's payload-free, closed S-4 wire
// union. It rejects non-canonical bytes and never returns a partial manifest.
func DecodeSettlementManifest(input []byte) (SettlementManifest, error) {
	canonical, err := jcs.Canonicalize(input)
	if err != nil || !bytes.Equal(input, canonical) {
		return SettlementManifest{}, ErrManifest
	}
	var manifest SettlementManifest
	decoder := json.NewDecoder(bytes.NewReader(input))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&manifest); err != nil {
		return SettlementManifest{}, ErrManifest
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return SettlementManifest{}, ErrManifest
	}
	if manifest.Version != 1 || manifest.RunID == "" || manifest.ProducedForTurnID == "" || manifest.Entries == nil {
		return SettlementManifest{}, ErrManifest
	}
	seen := make(map[string]struct{}, len(manifest.Entries))
	for _, entry := range manifest.Entries {
		if err := validateManifestWireEntry(manifest.RunID, entry); err != nil {
			return SettlementManifest{}, err
		}
		key := string(entry.Kind) + "\x00" + entry.RunID + "\x00" + entry.TurnID + "\x00" + manifestEffectID(entry)
		if _, duplicate := seen[key]; duplicate {
			return SettlementManifest{}, ErrManifest
		}
		seen[key] = struct{}{}
	}
	return manifest, nil
}

func validateManifestWireEntry(topRunID string, entry ManifestEntryWire) error {
	if entry.RunID == "" || entry.RunID != topRunID || entry.TurnID == "" || entry.Terminal == "" {
		return ErrManifest
	}
	if entry.Class != SettledWithContent && entry.Class != DeterminateNoResume && entry.Class != Uncertain {
		return ErrManifest
	}
	switch entry.Kind {
	case KindTool:
		if entry.ToolCallID == "" || !lowerHex(entry.ArgsDigest, 64) || entry.AttemptID != "" || entry.CancelPoint != "" {
			return ErrManifest
		}
		if (entry.VoidReason != "") != (entry.Terminal == "VOID") || (entry.VoidReason != "" && !validVoidReason(entry.VoidReason)) {
			return ErrManifest
		}
	case KindProvider:
		if entry.AttemptID == "" || entry.ToolCallID != "" || entry.ArgsDigest != "" || entry.VoidReason != "" {
			return ErrManifest
		}
		if (entry.CancelPoint != "") != (entry.Terminal == "CANCELLED") {
			return ErrManifest
		}
		if entry.CancelPoint != "" && entry.CancelPoint != "pre_transport" && entry.CancelPoint != "post_invocation" {
			return ErrManifest
		}
	default:
		return ErrManifest
	}
	return nil
}

func validVoidReason(reason string) bool {
	switch reason {
	case "run_not_admitted", "turn_inactive", "lease_invalid", "denied_above_set", "expired":
		return true
	default:
		return false
	}
}

func manifestEffectID(entry ManifestEntryWire) string {
	if entry.Kind == KindTool {
		return entry.ToolCallID
	}
	return entry.AttemptID
}

type ContentObservation struct {
	Present     bool
	DigestValid bool
	DurableAt   time.Time
	InspectedAt time.Time
}

type Decision struct {
	Disposition  string
	ResumeAction string
	FirstAction  FirstAction
	Trusted      bool
	Prohibition  string
}

var (
	ErrManifest       = errors.New("resume: invalid settlement entry")
	ErrTrustWindow    = errors.New("resume: trust-window violation")
	ErrPrefix         = errors.New("resume: frozen prefix unavailable")
	ErrEditSurface    = errors.New("resume: edit is outside the worker journal")
	ErrClassPromotion = errors.New("resume: edited record crosses trust classes")
)

// Reconcile applies the total five-branch first-action table. Settlement-time
// evidence and resume-time presence are separate conjuncts.
func Reconcile(entry SettlementEntry, content ContentObservation) (Decision, error) {
	if err := validateEntry(entry); err != nil {
		return Decision{}, err
	}
	if content.InspectedAt.IsZero() || content.InspectedAt.Before(entry.SettledAt) {
		return Decision{}, ErrTrustWindow
	}
	switch entry.Class {
	case SettledWithContent:
		if !content.Present {
			return degradedDecision(), nil
		}
		if content.DurableAt.IsZero() || content.DurableAt.After(entry.SettledAt) || !content.DigestValid {
			return Decision{}, ErrTrustWindow
		}
		return Decision{Disposition: DispositionResumable, FirstAction: ContinueModelLoop, Trusted: true, Prohibition: "not_a_resend"}, nil
	case DeterminateNoResume:
		return Decision{Disposition: DispositionResumable, FirstAction: SurfaceAndTerminalize, Prohibition: "no_automatic_replacement_attempt"}, nil
	case Uncertain:
		if entry.Kind == KindTool {
			return Decision{Disposition: DispositionResumable, FirstAction: SurfaceUncertainTool, Prohibition: "no_silent_reexecute"}, nil
		}
		return Decision{Disposition: DispositionResumable, FirstAction: SurfaceUncertainProvider, Prohibition: "no_auto_resend"}, nil
	default:
		return Decision{}, ErrManifest
	}
}

func degradedDecision() Decision {
	return Decision{Disposition: DispositionDegraded, ResumeAction: ResumeActionReDerive, FirstAction: ReDerive, Prohibition: "no_silent_resume_or_reconstruction"}
}

func validateEntry(entry SettlementEntry) error {
	if entry.RunID == "" || entry.TurnID == "" || entry.EffectID == "" || entry.SettledAt.IsZero() {
		return ErrManifest
	}
	if entry.Kind != KindTool && entry.Kind != KindProvider {
		return ErrManifest
	}
	if entry.Kind == KindTool {
		if !lowerHex(entry.ArgsDigest, 64) {
			return ErrManifest
		}
	} else if entry.ArgsDigest != "" {
		return ErrManifest
	}
	if entry.Class != SettledWithContent && entry.Class != DeterminateNoResume && entry.Class != Uncertain {
		return ErrManifest
	}
	return nil
}

type ContentReadyInput struct {
	Seq            string
	RunID          string
	TurnEpoch      string
	TurnID         string
	AttemptID      string
	RoundIdentity  string
	SeqHWM         string
	GenerationID   string
	ContentFsyncAt time.Time
	MarkerFsyncAt  time.Time
	EmitAt         time.Time
}

type contentReadyBody struct {
	TurnID        string `json:"turn_id"`
	AttemptID     string `json:"attempt_id"`
	RoundIdentity string `json:"round_identity"`
	SeqHWM        string `json:"seq_hwm"`
	GenerationID  string `json:"generation_id"`
}

// ContentReady constructs the one-way frame only after both durability
// linearization points. Raw content is deliberately absent from the frame.
func ContentReady(input ContentReadyInput) (wire.Frame, error) {
	if input.RunID == "" || input.TurnID == "" || input.AttemptID == "" || input.GenerationID == "" ||
		!lowerHex(input.RoundIdentity, 64) || !orderedTimes(input.ContentFsyncAt, input.MarkerFsyncAt, input.EmitAt) {
		return wire.Frame{}, errors.New("resume: invalid content_ready evidence")
	}
	if _, err := wire.ParseCounter(input.Seq); err != nil {
		return wire.Frame{}, err
	}
	if _, err := wire.ParseCounter(input.TurnEpoch); err != nil {
		return wire.Frame{}, err
	}
	if _, err := wire.ParseCounter(input.SeqHWM); err != nil {
		return wire.Frame{}, err
	}
	body, err := json.Marshal(contentReadyBody{input.TurnID, input.AttemptID, input.RoundIdentity, input.SeqHWM, input.GenerationID})
	if err != nil {
		return wire.Frame{}, err
	}
	return wire.Frame{Version: wire.FrameVersion, Channel: wire.ChannelCTRLW, Type: "content_ready", Seq: input.Seq, RunID: input.RunID, TurnEpoch: input.TurnEpoch, Body: body}, nil
}

func orderedTimes(first, second, third time.Time) bool {
	return !first.IsZero() && !second.IsZero() && !third.IsZero() && !second.Before(first) && !third.Before(second)
}

type dispositionBody struct {
	TurnID       string `json:"turn_id"`
	Disposition  string `json:"disposition"`
	ResumeAction string `json:"resume_action,omitempty"`
}

// WorkGate prevents provider, tool, and conductor work until m-10 returns its
// post-commit disposition pair. A conflict response is represented by Commit:
// the carried committed pair is adopted, never the locally reported pair.
type WorkGate struct {
	mu        sync.Mutex
	runID     string
	turnID    string
	turnEpoch string
	reported  bool
	committed dispositionBody
}

func NewWorkGate(runID, turnID, turnEpoch string) *WorkGate {
	return &WorkGate{runID: runID, turnID: turnID, turnEpoch: turnEpoch}
}

func (gate *WorkGate) Report(seq, disposition string) (wire.Frame, error) {
	gate.mu.Lock()
	defer gate.mu.Unlock()
	if gate.runID == "" || gate.turnID == "" {
		return wire.Frame{}, errors.New("resume: incomplete disposition identity")
	}
	if _, err := wire.ParseCounter(gate.turnEpoch); err != nil {
		return wire.Frame{}, err
	}
	if _, err := wire.ParseCounter(seq); err != nil {
		return wire.Frame{}, err
	}
	action, err := actionForDisposition(disposition)
	if err != nil {
		return wire.Frame{}, err
	}
	body := dispositionBody{TurnID: gate.turnID, Disposition: disposition, ResumeAction: action}
	raw, err := json.Marshal(body)
	if err != nil {
		return wire.Frame{}, err
	}
	gate.reported = true
	return wire.Frame{Version: wire.FrameVersion, Channel: wire.ChannelCTRLW, Type: "report_resume_disposition", Seq: seq, RunID: gate.runID, TurnEpoch: gate.turnEpoch, Body: raw}, nil
}

func (gate *WorkGate) Commit(disposition, resumeAction string) error {
	gate.mu.Lock()
	defer gate.mu.Unlock()
	if !gate.reported {
		return errors.New("resume: disposition committed before report")
	}
	want, err := actionForDisposition(disposition)
	if err != nil || want != resumeAction {
		return errors.New("resume: invalid committed disposition pair")
	}
	gate.committed = dispositionBody{TurnID: gate.turnID, Disposition: disposition, ResumeAction: resumeAction}
	return nil
}

func (gate *WorkGate) WorkAllowed() bool {
	gate.mu.Lock()
	defer gate.mu.Unlock()
	return gate.committed.Disposition != ""
}

func (gate *WorkGate) CommittedDisposition() string {
	gate.mu.Lock()
	defer gate.mu.Unlock()
	return gate.committed.Disposition
}

func actionForDisposition(disposition string) (string, error) {
	switch disposition {
	case DispositionResumable:
		return "", nil
	case DispositionDegraded:
		return ResumeActionReDerive, nil
	default:
		return "", errors.New("resume: invalid disposition")
	}
}

// ExtractCanonicalPrefix stops at the fixture-pinned full round key. It never
// examines bytes after that marker and excludes advisory record_digest values
// from the compared canonical record-content sequence.
func ExtractCanonicalPrefix(data []byte, identity journal.Identity, predecessorTurnID, resumedRoundIndex string) ([]byte, error) {
	if predecessorTurnID == "" {
		return nil, ErrPrefix
	}
	if _, err := wire.ParseCounter(resumedRoundIndex); err != nil {
		return nil, ErrPrefix
	}
	position := 0
	var bounded []byte
	found := false
	for position < len(data) {
		relative := bytes.IndexByte(data[position:], '\n')
		if relative < 0 {
			return nil, ErrPrefix
		}
		end := position + relative
		line := data[position:end]
		record, err := journal.DecodeRecord(line)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrPrefix, err)
		}
		bounded = append(bounded, line...)
		bounded = append(bounded, '\n')
		if record.Kind == journal.KindRoundMarker && record.TurnID == predecessorTurnID && rawString(record.Fields["round_index"]) == resumedRoundIndex {
			found = true
			break
		}
		position = end + 1
	}
	if !found {
		return nil, ErrPrefix
	}
	recovery := journal.Recover(bounded, identity)
	if recovery.Disposition != journal.DispositionResumable || recovery.Boundary.Kind != journal.BoundaryMarker || recovery.Boundary.Offset != int64(len(bounded)) {
		return nil, ErrPrefix
	}
	var compared bytes.Buffer
	for _, line := range recovery.Records {
		var object map[string]json.RawMessage
		if err := json.Unmarshal(line, &object); err != nil {
			return nil, ErrPrefix
		}
		delete(object, "record_digest")
		raw, err := json.Marshal(object)
		if err != nil {
			return nil, err
		}
		canonical, err := jcs.Canonicalize(raw)
		if err != nil {
			return nil, err
		}
		compared.Write(canonical)
		compared.WriteByte('\n')
	}
	return compared.Bytes(), nil
}

type EditSource string

const (
	SourceJournal       EditSource = "m9_journal"
	SourceM10Settlement EditSource = "m10_settlement"
	SourceM1Governed    EditSource = "m1_governed_record"
)

type EditAssessment struct {
	Disposition            string
	ResumeAction           string
	PresentAsOriginalTruth bool
}

func AssessJournalEdit(before, after journal.Record) (EditAssessment, error) {
	return AssessEdit(SourceJournal, before, after)
}

// AssessEdit enforces the MVP's journal-only edit surface and prevents a
// changed record from crossing into a more-trusted content class.
func AssessEdit(source EditSource, before, after journal.Record) (EditAssessment, error) {
	if source != SourceJournal {
		return EditAssessment{}, ErrEditSurface
	}
	beforeClass, err := recordTrustClass(before)
	if err != nil {
		return EditAssessment{Disposition: DispositionDegraded, ResumeAction: ResumeActionReDerive}, err
	}
	afterClass, err := recordTrustClass(after)
	if err != nil || beforeClass != afterClass {
		return EditAssessment{Disposition: DispositionDegraded, ResumeAction: ResumeActionReDerive}, ErrClassPromotion
	}
	if beforeClass == "model_transcript" {
		return EditAssessment{Disposition: DispositionResumable}, nil
	}
	return EditAssessment{Disposition: DispositionDegraded, ResumeAction: ResumeActionReDerive}, nil
}

func recordTrustClass(record journal.Record) (string, error) {
	switch record.Kind {
	case journal.KindProviderOutput, journal.KindToolResult:
		return "external_truth", nil
	case journal.KindInputItem:
		role := rawString(record.Fields["role"])
		switch role {
		case "tool_result":
			return "external_truth", nil
		case "user", "assistant":
			return "model_transcript", nil
		default:
			return "", errors.New("resume: unknown input_item role")
		}
	case journal.KindToolCall:
		return "model_transcript", nil
	case journal.KindCompactionEvent:
		return "governance_summary", nil
	case journal.KindRunOpen, journal.KindTurnScope, journal.KindObjectiveRef, journal.KindWorkspaceSnapshot, journal.KindRoundMarker:
		return "framing_reference", nil
	default:
		return "", errors.New("resume: unknown journal kind")
	}
}

func rawString(raw json.RawMessage) string {
	var value string
	_ = json.Unmarshal(raw, &value)
	return value
}

func lowerHex(value string, length int) bool {
	if len(value) != length {
		return false
	}
	for _, character := range value {
		if !strings.ContainsRune("0123456789abcdef", character) {
			return false
		}
	}
	return true
}
