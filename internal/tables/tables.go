package tables

import (
	"encoding/json"
	"sort"
	"sync"

	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

type T struct {
	Records         []record.Record
	ByRelay         map[string]record.Record
	ByDispatch      map[string][]record.Record
	AcceptedParents map[string]string
	Grants          map[string][]record.Record
	Verdicts        map[string][]record.Record
	Locks           map[string][]record.Record
	MergeGates      map[string][]record.Record
	Waivers         []record.Record
	OutcomeByIntake map[string]record.Record
	ContentHash     map[string]string
	CompletionIndex map[string]map[string]bool
	ParkedLanes     map[string]bool
}

type Live struct {
	mu      sync.RWMutex
	current *T
}

func New() *T {
	return &T{
		ByRelay:         map[string]record.Record{},
		ByDispatch:      map[string][]record.Record{},
		AcceptedParents: map[string]string{},
		Grants:          map[string][]record.Record{},
		Verdicts:        map[string][]record.Record{},
		Locks:           map[string][]record.Record{},
		MergeGates:      map[string][]record.Record{},
		OutcomeByIntake: map[string]record.Record{},
		ContentHash:     map[string]string{},
		CompletionIndex: map[string]map[string]bool{
			"park":       {},
			"outbox":     {},
			"incident":   {},
			"quarantine": {},
			"owed":       {},
		},
		ParkedLanes: map[string]bool{},
	}
}

func NewLive(initial *T) *Live {
	l := &Live{}
	l.Publish(initial)
	return l
}

func (l *Live) Snapshot() *T {
	if l == nil {
		return New()
	}
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.current == nil {
		return New()
	}
	return l.current.Clone()
}

func (l *Live) Publish(next *T) {
	if l == nil {
		return
	}
	if next == nil {
		next = New()
	}
	l.mu.Lock()
	l.current = next.Clone()
	l.mu.Unlock()
}

func (t *T) Clone() *T {
	if t == nil {
		return New()
	}
	out := New()
	out.Records = cloneRecords(t.Records)
	out.ByRelay = cloneRecordMap(t.ByRelay)
	out.ByDispatch = cloneRecordSliceMap(t.ByDispatch)
	out.AcceptedParents = cloneStringMap(t.AcceptedParents)
	out.Grants = cloneRecordSliceMap(t.Grants)
	out.Verdicts = cloneRecordSliceMap(t.Verdicts)
	out.Locks = cloneRecordSliceMap(t.Locks)
	out.MergeGates = cloneRecordSliceMap(t.MergeGates)
	out.Waivers = cloneRecords(t.Waivers)
	out.OutcomeByIntake = cloneRecordMap(t.OutcomeByIntake)
	out.ContentHash = cloneStringMap(t.ContentHash)
	out.CompletionIndex = cloneBoolMapMap(t.CompletionIndex)
	out.ParkedLanes = cloneBoolMap(t.ParkedLanes)
	return out
}

func Build(st *store.Store) (*T, error) {
	records, err := st.Records()
	if err != nil {
		return nil, err
	}
	t := New()
	for _, rec := range records {
		t.OnCommit(rec)
	}
	t.hydrateContentHashes(st)
	return t, nil
}

func (t *T) hydrateContentHashes(st *store.Store) {
	journal, err := intake.Open(st.Root)
	if err != nil {
		return
	}
	entries, err := journal.ReadAll()
	if err != nil {
		return
	}
	for _, entry := range entries {
		if entry.IntakeID == "" || entry.ContentHash == "" {
			continue
		}
		if _, ok := t.OutcomeByIntake[entry.IntakeID]; ok {
			t.ContentHash[entry.ContentHash] = entry.IntakeID
		}
	}
}

func (t *T) OnCommit(rec record.Record) {
	if t == nil {
		return
	}
	rec = verifiedShape(rec)
	t.Records = append(t.Records, rec)
	sort.Slice(t.Records, func(i, j int) bool {
		return t.Records[i].Envelope.RelayID < t.Records[j].Envelope.RelayID
	})
	t.ByRelay[rec.Envelope.RelayID] = rec
	if rec.Envelope.DispatchID != "" {
		t.ByDispatch[rec.Envelope.DispatchID] = append(t.ByDispatch[rec.Envelope.DispatchID], rec)
	}
	if rec.Envelope.DeliveryState == record.Accepted {
		t.AcceptedParents[rec.Envelope.RelayID] = rec.Envelope.RelayID
		if rec.Envelope.DispatchID != "" {
			t.AcceptedParents[rec.Envelope.DispatchID] = rec.Envelope.RelayID
		}
	}
	if rec.Envelope.IntakeID != "" {
		t.OutcomeByIntake[rec.Envelope.IntakeID] = rec
	}
	if rec.Headers["grant"] != "" {
		t.Grants[rec.Envelope.DispatchID] = append(t.Grants[rec.Envelope.DispatchID], rec)
	}
	if rec.Headers["resolves_gate"] != "" || rec.Headers["DESIGN_REVIEW_VERDICT"] != "" {
		t.Verdicts[rec.Envelope.DispatchID] = append(t.Verdicts[rec.Envelope.DispatchID], rec)
	}
	if rec.Headers["DESIGN_LOCK_ID"] != "" || rec.Headers["DESIGN_DOC_ID"] != "" {
		t.Locks[rec.Envelope.DispatchID] = append(t.Locks[rec.Envelope.DispatchID], rec)
	}
	if rec.Headers["PHASE"] == "MERGE-GATE" {
		t.MergeGates[rec.Envelope.DispatchID] = append(t.MergeGates[rec.Envelope.DispatchID], rec)
	}
	if rec.Headers["ORCH_REVIEW_WAIVER"] != "" {
		t.Waivers = append(t.Waivers, rec)
	}
	t.updateCompletion(rec)
}

func verifiedShape(rec record.Record) record.Record {
	data, err := record.Seal(rec)
	if err != nil {
		return rec
	}
	verified, err := record.Verify(data)
	if err != nil {
		return rec
	}
	return verified
}

func (t *T) updateCompletion(rec record.Record) {
	if rec.Headers["parks_gate"] != "" {
		t.CompletionIndex["park"][rec.Headers["parks_gate"]] = true
		t.ParkedLanes[rec.Headers["parks_gate"]] = true
	}
	if rec.Headers["quarantined_ref"] != "" {
		t.CompletionIndex["incident"][rec.Headers["quarantined_ref"]] = true
	}
	if rec.Headers["disposes_owed"] != "" {
		t.CompletionIndex["owed"][rec.Headers["disposes_owed"]] = true
	}
	var outbox struct {
		SourceRecordRef string `json:"source_record_ref"`
	}
	if rec.Body != "" && json.Unmarshal([]byte(rec.Body), &outbox) == nil && outbox.SourceRecordRef != "" {
		t.CompletionIndex["outbox"][outbox.SourceRecordRef] = true
	}
}

func cloneRecords(records []record.Record) []record.Record {
	out := make([]record.Record, len(records))
	for i, rec := range records {
		out[i] = cloneRecord(rec)
	}
	return out
}

func cloneRecordMap(in map[string]record.Record) map[string]record.Record {
	out := make(map[string]record.Record, len(in))
	for key, value := range in {
		out[key] = cloneRecord(value)
	}
	return out
}

func cloneRecordSliceMap(in map[string][]record.Record) map[string][]record.Record {
	out := make(map[string][]record.Record, len(in))
	for key, values := range in {
		out[key] = cloneRecords(values)
	}
	return out
}

func cloneRecord(rec record.Record) record.Record {
	rec.Headers = cloneStringMap(rec.Headers)
	return rec
}

func cloneStringMap(in map[string]string) map[string]string {
	if in == nil {
		return nil
	}
	out := make(map[string]string, len(in))
	for key, value := range in {
		out[key] = value
	}
	return out
}

func cloneBoolMapMap(in map[string]map[string]bool) map[string]map[string]bool {
	out := make(map[string]map[string]bool, len(in))
	for key, values := range in {
		out[key] = cloneBoolMap(values)
	}
	return out
}

func cloneBoolMap(in map[string]bool) map[string]bool {
	if in == nil {
		return nil
	}
	out := make(map[string]bool, len(in))
	for key, value := range in {
		out[key] = value
	}
	return out
}
