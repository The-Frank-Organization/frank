// Package journal implements the worker-owned, one-file-per-run session log.
package journal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/jackli/frank/internal/worker/jcs"
	"github.com/jackli/frank/internal/worker/wire"
)

const (
	KindRunOpen           = "run_open"
	KindTurnScope         = "turn_scope"
	KindObjectiveRef      = "objective_ref"
	KindWorkspaceSnapshot = "workspace_snapshot"
	KindInputItem         = "input_item"
	KindToolCall          = "tool_call"
	KindToolResult        = "tool_result"
	KindProviderOutput    = "provider_output"
	KindCompactionEvent   = "compaction_event"
	KindRoundMarker       = "round_marker"
)

var recordKinds = []string{
	KindRunOpen,
	KindTurnScope,
	KindObjectiveRef,
	KindWorkspaceSnapshot,
	KindInputItem,
	KindToolCall,
	KindToolResult,
	KindProviderOutput,
	KindCompactionEvent,
	KindRoundMarker,
}

// Record holds the common envelope and the kind-specific top-level members.
// Fields remain raw canonical JSON so content values are not interpreted by
// the journal.
type Record struct {
	Seq          string
	Kind         string
	GenerationID string
	TurnID       string
	RoundIndex   string
	TSMonotonic  string
	RecordDigest string
	Fields       map[string]json.RawMessage
}

// RecordKinds returns the stable reduced closed union in normative order.
func RecordKinds() []string {
	return append([]string(nil), recordKinds...)
}

var commonMembers = map[string]struct{}{
	"seq": {}, "kind": {}, "generation_id": {}, "turn_id": {},
	"round_index": {}, "ts_monotonic": {}, "record_digest": {},
}

type schema struct {
	required   []string
	optional   []string
	turnID     presence
	roundIndex presence
}

type presence uint8

const (
	presenceAbsent presence = iota
	presenceRequired
)

var schemas = map[string]schema{
	KindRunOpen: {
		required: []string{"run_id", "run_manifest_digest", "turn_epoch", "create_auth_id"},
	},
	KindTurnScope: {
		required: []string{"admission_ref_kind"}, optional: []string{"predecessor_turn_id"}, turnID: presenceRequired,
	},
	KindObjectiveRef: {
		required: []string{"objective_locator", "constraint_refs"}, turnID: presenceRequired,
	},
	KindWorkspaceSnapshot: {
		required: []string{"workspace_root_id", "snapshot_id"}, turnID: presenceRequired,
	},
	KindInputItem: {
		required: []string{"role", "item_index", "content"}, optional: []string{"source_tool_call_id"}, turnID: presenceRequired, roundIndex: presenceRequired,
	},
	KindToolCall: {
		required: []string{"tool_call_id", "canonical_tool_name", "canonical_args_digest", "args"}, turnID: presenceRequired, roundIndex: presenceRequired,
	},
	KindToolResult: {
		required: []string{"tool_call_id", "content", "truncated"}, turnID: presenceRequired, roundIndex: presenceRequired,
	},
	KindProviderOutput: {
		required: []string{"attempt_id", "item_index", "content"}, turnID: presenceRequired, roundIndex: presenceRequired,
	},
	KindCompactionEvent: {
		required: []string{"tier", "template_id", "template_version", "affected_seq"}, optional: []string{"summary_item_index"}, turnID: presenceRequired, roundIndex: presenceRequired,
	},
	KindRoundMarker: {
		required: []string{"round_index", "first_seq", "last_seq", "marker_digest"}, turnID: presenceRequired,
	},
}

// FinalizeRecord computes and installs the self-excluding record digest.
func FinalizeRecord(record Record) (Record, error) {
	record.RecordDigest = ""
	if err := validateRecord(record, false); err != nil {
		return Record{}, err
	}
	encoded, err := marshalRecordObject(record, false)
	if err != nil {
		return Record{}, err
	}
	digest, err := jcs.Digest(encoded)
	if err != nil {
		return Record{}, fmt.Errorf("digest record: %w", err)
	}
	record.RecordDigest = digest
	return record, nil
}

// MarshalRecord returns one canonical record without its line terminator.
func MarshalRecord(record Record) ([]byte, error) {
	if err := validateRecord(record, true); err != nil {
		return nil, err
	}
	return marshalRecordObject(record, true)
}

// DecodeRecord requires one canonical record without its line terminator.
func DecodeRecord(input []byte) (Record, error) {
	canonical, err := jcs.Canonicalize(input)
	if err != nil {
		return Record{}, fmt.Errorf("invalid record JSON: %w", err)
	}
	if !bytes.Equal(input, canonical) {
		return Record{}, errors.New("record is not canonical JSON")
	}
	var object map[string]json.RawMessage
	if err := json.Unmarshal(input, &object); err != nil {
		return Record{}, fmt.Errorf("decode record object: %w", err)
	}
	record := Record{Fields: make(map[string]json.RawMessage)}
	if record.Seq, err = takeString(object, "seq", true); err != nil {
		return Record{}, err
	}
	if record.Kind, err = takeString(object, "kind", true); err != nil {
		return Record{}, err
	}
	if record.GenerationID, err = takeString(object, "generation_id", true); err != nil {
		return Record{}, err
	}
	if record.TurnID, err = takeString(object, "turn_id", false); err != nil {
		return Record{}, err
	}
	if record.Kind != KindRoundMarker {
		if record.RoundIndex, err = takeString(object, "round_index", false); err != nil {
			return Record{}, err
		}
	}
	if record.TSMonotonic, err = takeString(object, "ts_monotonic", true); err != nil {
		return Record{}, err
	}
	if record.RecordDigest, err = takeString(object, "record_digest", true); err != nil {
		return Record{}, err
	}
	for member, raw := range object {
		_, common := commonMembers[member]
		if record.Kind == KindRoundMarker && member == "round_index" {
			common = false
		}
		if common {
			continue
		}
		record.Fields[member] = append(json.RawMessage(nil), raw...)
	}
	if err := validateRecord(record, true); err != nil {
		return Record{}, err
	}
	return record, nil
}

// VerifyRecordDigest recomputes the advisory self-excluding digest.
func VerifyRecordDigest(record Record) (bool, error) {
	stored := record.RecordDigest
	if !lowerHex(stored, 64) {
		return false, errors.New("record_digest is not lowercase SHA-256")
	}
	record.RecordDigest = ""
	encoded, err := marshalRecordObject(record, false)
	if err != nil {
		return false, err
	}
	digest, err := jcs.Digest(encoded)
	if err != nil {
		return false, err
	}
	return digest == stored, nil
}

func marshalRecordObject(record Record, includeDigest bool) ([]byte, error) {
	object := make(map[string]json.RawMessage, len(record.Fields)+7)
	object["seq"] = marshalString(record.Seq)
	object["kind"] = marshalString(record.Kind)
	object["generation_id"] = marshalString(record.GenerationID)
	if record.TurnID != "" {
		object["turn_id"] = marshalString(record.TurnID)
	}
	if record.RoundIndex != "" {
		object["round_index"] = marshalString(record.RoundIndex)
	}
	object["ts_monotonic"] = marshalString(record.TSMonotonic)
	if includeDigest {
		object["record_digest"] = marshalString(record.RecordDigest)
	}
	for member, raw := range record.Fields {
		object[member] = append(json.RawMessage(nil), raw...)
	}
	raw, err := json.Marshal(object)
	if err != nil {
		return nil, fmt.Errorf("marshal record: %w", err)
	}
	canonical, err := jcs.Canonicalize(raw)
	if err != nil {
		return nil, fmt.Errorf("canonicalize record: %w", err)
	}
	return canonical, nil
}

func validateRecord(record Record, requireDigest bool) error {
	spec, known := schemas[record.Kind]
	if !known {
		return fmt.Errorf("unknown record kind %q", record.Kind)
	}
	if _, err := wire.ParseCounter(record.Seq); err != nil {
		return fmt.Errorf("invalid seq: %w", err)
	}
	if record.Kind == KindRunOpen && record.Seq != "0" {
		return errors.New("run_open seq must be 0")
	}
	if record.GenerationID == "" {
		return errors.New("generation_id is empty")
	}
	if _, err := wire.ParseCounter(record.TSMonotonic); err != nil {
		return fmt.Errorf("invalid ts_monotonic: %w", err)
	}
	if err := validatePresence("turn_id", record.TurnID, spec.turnID); err != nil {
		return err
	}
	if err := validatePresence("round_index", record.RoundIndex, spec.roundIndex); err != nil {
		return err
	}
	if record.RoundIndex != "" {
		if _, err := wire.ParseCounter(record.RoundIndex); err != nil {
			return fmt.Errorf("invalid round_index: %w", err)
		}
	}
	if requireDigest && !lowerHex(record.RecordDigest, 64) {
		return errors.New("record_digest is not lowercase SHA-256")
	}
	allowed := make(map[string]struct{}, len(spec.required)+len(spec.optional))
	for _, member := range spec.required {
		allowed[member] = struct{}{}
		if _, exists := record.Fields[member]; !exists {
			return fmt.Errorf("%s lacks required member %q", record.Kind, member)
		}
	}
	for _, member := range spec.optional {
		allowed[member] = struct{}{}
	}
	for member, raw := range record.Fields {
		if _, exists := allowed[member]; !exists {
			return fmt.Errorf("unknown member %q for %s", member, record.Kind)
		}
		if _, err := jcs.Canonicalize(raw); err != nil {
			return fmt.Errorf("invalid member %q: %w", member, err)
		}
	}
	return validatePayload(record)
}

func validatePresence(member, value string, required presence) error {
	if required == presenceRequired && value == "" {
		return fmt.Errorf("%s is required", member)
	}
	if required == presenceAbsent && value != "" {
		return fmt.Errorf("%s must be absent", member)
	}
	return nil
}

func validatePayload(record Record) error {
	stringMember := func(name string) (string, error) { return rawStringValue(record.Fields[name], name) }
	counterMember := func(name string) error {
		value, err := stringMember(name)
		if err != nil {
			return err
		}
		if _, err := wire.ParseCounter(value); err != nil {
			return fmt.Errorf("%s is not a canonical counter", name)
		}
		return nil
	}
	switch record.Kind {
	case KindRunOpen:
		for _, member := range []string{"run_id", "run_manifest_digest", "turn_epoch", "create_auth_id"} {
			if _, err := stringMember(member); err != nil {
				return err
			}
		}
		manifest, _ := stringMember("run_manifest_digest")
		auth, _ := stringMember("create_auth_id")
		if !lowerHex(manifest, 64) || !lowerHex(auth, 32) {
			return errors.New("run_open digest or create_auth_id grammar mismatch")
		}
		return counterMember("turn_epoch")
	case KindTurnScope:
		kind, err := stringMember("admission_ref_kind")
		if err != nil || (kind != "wake_relay" && kind != "operator_input") {
			return errors.New("turn_scope admission_ref_kind is invalid")
		}
		if raw, exists := record.Fields["predecessor_turn_id"]; exists {
			if _, err := rawStringValue(raw, "predecessor_turn_id"); err != nil {
				return err
			}
		}
	case KindObjectiveRef:
		if _, err := stringMember("objective_locator"); err != nil {
			return err
		}
		return validateStringArray(record.Fields["constraint_refs"], "constraint_refs")
	case KindWorkspaceSnapshot:
		for _, member := range []string{"workspace_root_id", "snapshot_id"} {
			if _, err := stringMember(member); err != nil {
				return err
			}
		}
	case KindInputItem:
		role, err := stringMember("role")
		if err != nil || (role != "user" && role != "assistant" && role != "tool_result") {
			return errors.New("input_item role is invalid")
		}
		if err := counterMember("item_index"); err != nil {
			return err
		}
		_, sourceExists := record.Fields["source_tool_call_id"]
		if (role == "tool_result") != sourceExists {
			return errors.New("input_item source_tool_call_id conditional presence mismatch")
		}
		if sourceExists {
			if _, err := stringMember("source_tool_call_id"); err != nil {
				return err
			}
		}
	case KindToolCall:
		for _, member := range []string{"tool_call_id", "canonical_tool_name", "canonical_args_digest"} {
			if _, err := stringMember(member); err != nil {
				return err
			}
		}
		digest, _ := stringMember("canonical_args_digest")
		if !lowerHex(digest, 64) {
			return errors.New("canonical_args_digest is not lowercase SHA-256")
		}
		canonicalArgs, err := jcs.Canonicalize(record.Fields["args"])
		if err != nil || len(canonicalArgs) == 0 || canonicalArgs[0] != '{' {
			return errors.New("tool_call args must be a JSON object")
		}
	case KindToolResult:
		if _, err := stringMember("tool_call_id"); err != nil {
			return err
		}
		var truncated bool
		if err := json.Unmarshal(record.Fields["truncated"], &truncated); err != nil {
			return errors.New("tool_result truncated must be boolean")
		}
	case KindProviderOutput:
		if _, err := stringMember("attempt_id"); err != nil {
			return err
		}
		return counterMember("item_index")
	case KindCompactionEvent:
		tier, err := stringMember("tier")
		if err != nil || (tier != "evict" && tier != "summarize") {
			return errors.New("compaction_event tier is invalid")
		}
		for _, member := range []string{"template_id", "template_version"} {
			if _, err := stringMember(member); err != nil {
				return err
			}
		}
		if err := validateCounterArray(record.Fields["affected_seq"], "affected_seq"); err != nil {
			return err
		}
		if raw, exists := record.Fields["summary_item_index"]; exists {
			value, err := rawStringValue(raw, "summary_item_index")
			if err != nil {
				return err
			}
			if _, err := wire.ParseCounter(value); err != nil {
				return errors.New("summary_item_index is not a canonical counter")
			}
		}
	case KindRoundMarker:
		for _, member := range []string{"round_index", "first_seq", "last_seq"} {
			if err := counterMember(member); err != nil {
				return err
			}
		}
		digest, err := stringMember("marker_digest")
		if err != nil || !lowerHex(digest, 64) {
			return errors.New("marker_digest is not lowercase SHA-256")
		}
	}
	return nil
}

func takeString(object map[string]json.RawMessage, name string, required bool) (string, error) {
	raw, exists := object[name]
	if !exists {
		if required {
			return "", fmt.Errorf("record lacks %s", name)
		}
		return "", nil
	}
	return rawStringValue(raw, name)
}

func rawStringValue(raw json.RawMessage, name string) (string, error) {
	var value string
	if err := json.Unmarshal(raw, &value); err != nil {
		return "", fmt.Errorf("%s must be a string", name)
	}
	return value, nil
}

func validateStringArray(raw json.RawMessage, name string) error {
	var values []string
	if err := json.Unmarshal(raw, &values); err != nil {
		return fmt.Errorf("%s must be a string array", name)
	}
	return nil
}

func validateCounterArray(raw json.RawMessage, name string) error {
	var values []string
	if err := json.Unmarshal(raw, &values); err != nil {
		return fmt.Errorf("%s must be a counter-string array", name)
	}
	for _, value := range values {
		if _, err := wire.ParseCounter(value); err != nil {
			return fmt.Errorf("%s contains a malformed counter", name)
		}
	}
	return nil
}

func marshalString(value string) json.RawMessage {
	raw, _ := json.Marshal(value)
	return raw
}

func lowerHex(value string, length int) bool {
	if len(value) != length {
		return false
	}
	for _, character := range value {
		if character >= '0' && character <= '9' {
			continue
		}
		if character < 'a' || character > 'f' {
			return false
		}
	}
	return true
}

// TrustState is the transition table's three-column input domain.
type TrustState string

const (
	TrustClean             TrustState = "clean"
	TrustDigestMismatch    TrustState = "digest_mismatch"
	TrustStructuralFailure TrustState = "structural_failure"
)

// TransitionAction is the reader action selected by a table cell.
type TransitionAction string

const (
	ActionContinue         TransitionAction = "continue"
	ActionDegrade          TransitionAction = "degrade"
	ActionResolveObjective TransitionAction = "resolve_objective"
)

// TransitionRow is one cell in the total per-kind trust transition table.
type TransitionRow struct {
	Subject string
	State   TrustState
	Action  TransitionAction
}

var transitionSubjects = []string{
	"input_item:user", "input_item:assistant", "input_item:tool_result",
	KindToolCall, KindToolResult, KindProviderOutput, KindCompactionEvent,
	KindObjectiveRef, KindWorkspaceSnapshot, KindTurnScope, KindRunOpen, KindRoundMarker,
}

// TransitionSubjects returns the table's stable subject domain.
func TransitionSubjects() []string {
	return append([]string(nil), transitionSubjects...)
}

// TransitionRows materializes every subject by every trust state exactly once.
func TransitionRows() []TransitionRow {
	states := []TrustState{TrustClean, TrustDigestMismatch, TrustStructuralFailure}
	rows := make([]TransitionRow, 0, len(transitionSubjects)*len(states))
	for _, subject := range transitionSubjects {
		for _, state := range states {
			rows = append(rows, TransitionRow{Subject: subject, State: state, Action: transitionAction(subject, state)})
		}
	}
	return rows
}

func transitionAction(subject string, state TrustState) TransitionAction {
	if state == TrustClean {
		return ActionContinue
	}
	if state == TrustStructuralFailure {
		return ActionDegrade
	}
	switch subject {
	case "input_item:user", "input_item:assistant", KindToolCall:
		return ActionContinue
	case KindObjectiveRef:
		return ActionResolveObjective
	default:
		return ActionDegrade
	}
}

func transitionSubject(record Record) (string, error) {
	if record.Kind != KindInputItem {
		return record.Kind, nil
	}
	role, err := rawStringValue(record.Fields["role"], "role")
	if err != nil {
		return "", err
	}
	return KindInputItem + ":" + role, nil
}

// BuildRoundMarker creates the marker immediately after a contiguous member
// interval. Its digest binds the ordered stored {seq, record_digest} pairs.
func BuildRoundMarker(seq, generationID, turnID, roundIndex, tsMonotonic string, members []Record) (Record, error) {
	if len(members) == 0 {
		return Record{}, errors.New("round has no members")
	}
	markerDigest, err := roundMarkerDigest(members)
	if err != nil {
		return Record{}, err
	}
	record := Record{
		Seq: seq, Kind: KindRoundMarker, GenerationID: generationID,
		TurnID: turnID, TSMonotonic: tsMonotonic,
		Fields: map[string]json.RawMessage{
			"round_index":   marshalString(roundIndex),
			"first_seq":     marshalString(members[0].Seq),
			"last_seq":      marshalString(members[len(members)-1].Seq),
			"marker_digest": marshalString(markerDigest),
		},
	}
	finalized, err := FinalizeRecord(record)
	if err != nil {
		return Record{}, err
	}
	if err := HonourRoundMarker(finalized, members, nil); err != nil {
		return Record{}, fmt.Errorf("build round marker: %w", err)
	}
	return finalized, nil
}

// HonourRoundMarker verifies interval, key, exclusion, sequence, and digest.
func HonourRoundMarker(marker Record, members, acceptedOutside []Record) error {
	if marker.Kind != KindRoundMarker {
		return errors.New("record is not a round_marker")
	}
	if err := validateRecord(marker, true); err != nil {
		return err
	}
	if ok, err := VerifyRecordDigest(marker); err != nil || !ok {
		return errors.New("round_marker record digest mismatch")
	}
	if len(members) == 0 {
		return errors.New("round marker interval is empty")
	}
	roundIndex, _ := rawStringValue(marker.Fields["round_index"], "round_index")
	firstSeq, _ := rawStringValue(marker.Fields["first_seq"], "first_seq")
	lastSeq, _ := rawStringValue(marker.Fields["last_seq"], "last_seq")
	first, _ := wire.ParseCounter(firstSeq)
	last, _ := wire.ParseCounter(lastSeq)
	markerSeq, _ := wire.ParseCounter(marker.Seq)
	if markerSeq != last+1 || members[0].Seq != firstSeq || members[len(members)-1].Seq != lastSeq {
		return errors.New("round marker coordinates mismatch")
	}
	if uint64(len(members)) != last-first+1 {
		return errors.New("round marker interval is not contiguous")
	}
	for index, member := range members {
		seq, err := wire.ParseCounter(member.Seq)
		if err != nil || seq != first+uint64(index) {
			return errors.New("round member sequence is not contiguous")
		}
		if !admitEligible(member.Kind) || member.TurnID != marker.TurnID || member.RoundIndex != roundIndex {
			return errors.New("round member key or eligibility mismatch")
		}
	}
	for _, record := range acceptedOutside {
		if admitEligible(record.Kind) && record.TurnID == marker.TurnID && record.RoundIndex == roundIndex {
			return errors.New("same-key admit-eligible record exists outside marker interval")
		}
	}
	wantDigest, err := roundMarkerDigest(members)
	if err != nil {
		return err
	}
	gotDigest, _ := rawStringValue(marker.Fields["marker_digest"], "marker_digest")
	if gotDigest != wantDigest {
		return errors.New("round marker digest mismatch")
	}
	return nil
}

func roundMarkerDigest(members []Record) (string, error) {
	pairs := make([]map[string]string, 0, len(members))
	for _, member := range members {
		if !lowerHex(member.RecordDigest, 64) {
			return "", errors.New("round member lacks record_digest")
		}
		pairs = append(pairs, map[string]string{"seq": member.Seq, "record_digest": member.RecordDigest})
	}
	raw, err := json.Marshal(pairs)
	if err != nil {
		return "", err
	}
	return jcs.Digest(raw)
}

func admitEligible(kind string) bool {
	switch kind {
	case KindInputItem, KindToolCall, KindToolResult, KindProviderOutput, KindCompactionEvent:
		return true
	default:
		return false
	}
}

// SortedFieldNames is useful for diagnostics without exposing field contents.
func SortedFieldNames(record Record) []string {
	names := make([]string, 0, len(record.Fields))
	for name := range record.Fields {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func faultSafeDetail(err error) string {
	if err == nil {
		return ""
	}
	return strings.SplitN(err.Error(), ":", 2)[0]
}
