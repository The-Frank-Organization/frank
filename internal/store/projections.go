package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackli/frank/internal/crashpoint"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/fsio"
	"github.com/jackli/frank/internal/record"
)

func (s *Store) AppendIndex(row string) error {
	return appendUnique(filepath.Join(s.Root, "projections", "INDEX.md"), []byte(row))
}

func (s *Store) RebuildProjections() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	entries, err := readRedo(s.Root)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		recordPath := filepath.Join(s.Root, "records", entry.RelayID+".json")
		if _, err := os.Stat(recordPath); err != nil {
			continue
		}
		if err := s.applyIntents(entry.Intents); err != nil {
			return err
		}
	}
	if err := resetMailboxes(filepath.Join(s.Root, "mailboxes")); err != nil {
		return err
	}
	records, err := s.recordsLocked()
	if err != nil {
		return err
	}
	for _, rec := range records {
		intents, err := canonicalProjectionIntents(rec)
		if err != nil {
			return err
		}
		if err := s.applyIntents(intents); err != nil {
			return err
		}
	}
	if err := s.applyIntents([]Intent{OwedOpenProjectionIntent(records)}); err != nil {
		return err
	}
	return nil
}

func (s *Store) applyIntents(intents []Intent) error {
	for _, intent := range intents {
		switch intent.Kind {
		case IntentIndex:
			crashpoint.Hit("pre_projection_write")
			if err := appendUnique(filepath.Join(s.Root, "projections", intent.Path), intent.Payload); err != nil {
				return err
			}
			crashpoint.Hit("post_projection_write")
		case IntentRender:
			crashpoint.Hit("pre_projection_write")
			if err := fsio.WriteFileAtomic(filepath.Join(s.Root, "projections"), intent.Path, intent.Payload); err != nil {
				return err
			}
			crashpoint.Hit("post_projection_write")
		case IntentMailbox:
			crashpoint.Hit("pre_projection_write")
			crashpoint.Hit("pre_delivery_write")
			if err := appendUnique(filepath.Join(s.Root, "mailboxes", intent.Path), intent.Payload); err != nil {
				return err
			}
			crashpoint.Hit("post_delivery_write")
			crashpoint.Hit("post_projection_write")
		case IntentOutbox:
			crashpoint.Hit("pre_projection_write")
			if err := fsio.WriteFileAtomic(s.Root, filepath.Join("outbox", intent.Path), intent.Payload); err != nil {
				return err
			}
			crashpoint.Hit("post_projection_write")
		case IntentConfig:
			crashpoint.Hit("pre_projection_write")
			if err := fsio.WriteFileAtomic(s.Root, intent.Path, intent.Payload); err != nil {
				return err
			}
			crashpoint.Hit("post_projection_write")
		}
	}
	return nil
}

func DefaultProjectionIntents(rec record.Record) []Intent {
	intents, _ := DefaultProjectionIntentsStrict(rec)
	return intents
}

func DefaultProjectionIntentsStrict(rec record.Record) ([]Intent, error) {
	relayID := rec.Envelope.RelayID
	if relayID == "" {
		return nil, nil
	}
	phase := safeSegment(rec.Headers["PHASE"])
	if phase == "" {
		phase = "record"
	}
	role := safeSegment(rec.Envelope.Role)
	if role == "" {
		role = "seat"
	}
	dispatchID := safeSegment(rec.Envelope.DispatchID)
	if dispatchID == "" {
		dispatchID = "unassigned"
	}
	renderPath := filepath.Join("relays", dispatchID, fmt.Sprintf("%s-%s-%s.md", phase, role, relayID))
	toList, ccList, err := decodedHeaderRecipients(rec)
	if err != nil {
		return nil, err
	}
	toDisplay := joinRecipients(toList)
	ccDisplay := joinRecipients(ccList)
	if toDisplay == "" && !hasAddressListHeader(rec, "TO") {
		toDisplay = rec.Envelope.To
	}
	render := []byte(fmt.Sprintf("## %s\n\nFROM: %s\nTO: %s\nCC: %s\nSUBJECT: %s\n\n%s\n", relayID, rec.Envelope.From, toDisplay, ccDisplay, rec.Headers["SUBJECT"], rec.Body))
	intents := []Intent{
		{Kind: IntentIndex, Path: "INDEX.md", Payload: []byte(fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n", relayID, rec.Headers["PHASE"], rec.Envelope.From, toDisplay, ccDisplay, rec.Envelope.DeliveryState))},
		{Kind: IntentRender, Path: renderPath, Payload: render},
	}
	if rec.Envelope.DeliveryState == record.Accepted {
		recipients, err := DeliveryRecipients(rec)
		if err != nil {
			return nil, err
		}
		for _, recipient := range recipients {
			intents = append(intents, Intent{Kind: IntentMailbox, Path: safeMailbox(recipient) + ".jsonl", Payload: []byte(relayID + "\n")})
		}
	}
	return intents, nil
}

func DeliveryRecipients(rec record.Record) ([]string, error) {
	toList, ccList, err := decodedHeaderRecipients(rec)
	if err != nil {
		return nil, err
	}
	seen := map[string]bool{}
	var recipients []string
	add := func(value string) {
		value = strings.TrimSpace(value)
		if value == "" || seen[value] {
			return
		}
		seen[value] = true
		recipients = append(recipients, value)
	}
	if hasAddressListHeader(rec, "TO") {
		for _, value := range toList {
			add(value)
		}
	} else {
		add(rec.Envelope.To)
	}
	for _, value := range ccList {
		add(value)
	}
	return recipients, nil
}

func hasAddressListHeader(rec record.Record, header string) bool {
	return strings.TrimSpace(rec.Headers[header]) != ""
}

func decodedHeaderRecipients(rec record.Record) ([]string, []string, error) {
	toList, err := addressList("TO", rec.Headers["TO"])
	if err != nil {
		return nil, nil, err
	}
	ccList, err := addressList("CC", rec.Headers["CC"])
	if err != nil {
		return nil, nil, err
	}
	return toList, ccList, nil
}

func addressList(header, raw string) ([]string, error) {
	values, err := fieldspec.DecodeAddressList(raw)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", header, err)
	}
	return values, nil
}

func joinRecipients(values []string) string {
	return strings.Join(values, ", ")
}

func canonicalProjectionIntents(rec record.Record) ([]Intent, error) {
	if rec.Headers["record_kind"] == "config_change" {
		return ConfigChangeIntentsStrict(rec)
	}
	intents, err := DefaultProjectionIntentsStrict(rec)
	if err != nil {
		return nil, err
	}
	if rec.Body == "" {
		return intents, nil
	}
	var outbox struct {
		ItemID string `json:"item_id"`
	}
	if err := json.Unmarshal([]byte(rec.Body), &outbox); err != nil || outbox.ItemID == "" {
		return intents, nil
	}
	intents = append(intents, Intent{Kind: IntentOutbox, Path: outbox.ItemID + ".json", Payload: []byte(rec.Body)})
	return intents, nil
}

func resetMailboxes(dir string) error {
	if err := os.RemoveAll(dir); err != nil {
		return err
	}
	return os.MkdirAll(dir, 0o755)
}

func OwedProjectionIntentsForCandidate(st *Store, cand record.Record) []Intent {
	records, err := st.Records()
	if err != nil {
		return nil
	}
	records = append(records, cand)
	return []Intent{OwedOpenProjectionIntent(records)}
}

func OwedOpenProjectionIntent(records []record.Record) Intent {
	type owedRow struct {
		RelayID         string
		Owner           string
		Source          string
		TargetSurface   string
		DispositionPath string
	}
	disposed := map[string]bool{}
	for _, rec := range records {
		if rec.Headers["record_kind"] == "owed_disposition" && rec.Headers["disposes_owed"] != "" && rec.Envelope.DeliveryState == record.Accepted {
			disposed[rec.Headers["disposes_owed"]] = true
		}
	}
	var rows []owedRow
	for _, rec := range records {
		if rec.Headers["record_kind"] != "owed_item" || rec.Envelope.DeliveryState != record.Accepted {
			continue
		}
		if disposed[rec.Envelope.RelayID] {
			continue
		}
		rows = append(rows, owedRow{
			RelayID:         rec.Envelope.RelayID,
			Owner:           rec.Headers["owner"],
			Source:          rec.Headers["source"],
			TargetSurface:   rec.Headers["target_surface"],
			DispositionPath: rec.Headers["disposition_path"],
		})
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].RelayID < rows[j].RelayID })
	var b strings.Builder
	b.WriteString("# Open owed items\n\n")
	b.WriteString("| relay_id | owner | source | target_surface | disposition_path |\n")
	b.WriteString("|---|---|---|---|---|\n")
	for _, row := range rows {
		fmt.Fprintf(&b, "| %s | %s | %s | %s | %s |\n", row.RelayID, row.Owner, row.Source, row.TargetSurface, row.DispositionPath)
	}
	return Intent{Kind: IntentRender, Path: filepath.Join("owed", "OPEN.md"), Payload: []byte(b.String())}
}

// ProjectBucketB is the live, non-interrupting digest for records whose
// committed gate_category belongs to the pinned operator-configured B set.
// It is a saved query over immutable record tags: no mailbox is created and a
// monotonic raise that rewrites the category to A cannot remain in this view.
func (s *Store) ProjectBucketB(reg *fieldspec.Registry) ([]string, error) {
	if reg == nil {
		return nil, fmt.Errorf("bucket B registry required")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	records, err := s.recordsLocked()
	if err != nil {
		return nil, err
	}
	var relayIDs []string
	for _, rec := range records {
		if rec.Envelope.DeliveryState != record.Accepted {
			continue
		}
		class, _ := reg.ClassifyGateCategory(rec.Headers["gate_category"], false)
		if class == "B" {
			relayIDs = append(relayIDs, rec.Envelope.RelayID)
		}
	}
	return relayIDs, nil
}

// ProjectBucketC is the low-priority FYI saved query. Only records that name
// the operator on CC, and not on TO, belong here; delivery remains ordinary
// mailbox delivery and creates no decision obligation.
func (s *Store) ProjectBucketC() ([]string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	records, err := s.recordsLocked()
	if err != nil {
		return nil, err
	}
	var relayIDs []string
	for _, rec := range records {
		if rec.Envelope.DeliveryState != record.Accepted {
			continue
		}
		to, cc, err := decodedHeaderRecipients(rec)
		if err != nil {
			return nil, err
		}
		if !hasAddressListHeader(rec, "TO") {
			to = append(to, rec.Envelope.To)
		}
		if containsRecipient(to, "operator") || !containsRecipient(cc, "operator") {
			continue
		}
		relayIDs = append(relayIDs, rec.Envelope.RelayID)
	}
	return relayIDs, nil
}

// ProjectBucketD is the acceptance-bounce return view for an authoring seat.
// Egress failures are excluded because they occur after acceptance and remain
// on the local A-gate resummon path.
func (s *Store) ProjectBucketD(author string) ([]string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	records, err := s.recordsLocked()
	if err != nil {
		return nil, err
	}
	var relayIDs []string
	for _, rec := range records {
		if rec.Envelope.DeliveryState != record.Rejected || rec.Envelope.From != author || !bucketDFailingEdge(rec.Headers["failing_edge"]) {
			continue
		}
		relayIDs = append(relayIDs, rec.Envelope.RelayID)
	}
	return relayIDs, nil
}

func bucketDFailingEdge(edge string) bool {
	switch edge {
	case "form-validation", "lineage", "observe-predicate", "declared-vs-observed":
		return true
	default:
		return false
	}
}

func containsRecipient(values []string, recipient string) bool {
	for _, value := range values {
		if strings.TrimSpace(value) == recipient {
			return true
		}
	}
	return false
}

func appendUnique(path string, payload []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	current, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	for _, line := range strings.SplitAfter(string(current), "\n") {
		if line == string(payload) {
			return nil
		}
	}
	if len(payload) == 0 {
		return nil
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	if err := fsio.AppendFsync(f, payload); err != nil {
		_ = f.Close()
		return err
	}
	return f.Close()
}

func safeSegment(value string) string {
	value = strings.TrimSpace(value)
	value = strings.ReplaceAll(value, "/", "_")
	value = strings.ReplaceAll(value, string(filepath.Separator), "_")
	return value
}

func safeMailbox(value string) string {
	return safeSegment(value)
}
