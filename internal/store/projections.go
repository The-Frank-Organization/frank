package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackli/frank/internal/crashpoint"
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
	records, err := s.recordsLocked()
	if err != nil {
		return err
	}
	for _, rec := range records {
		if err := s.applyIntents(canonicalProjectionIntents(rec)); err != nil {
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
		}
	}
	return nil
}

func DefaultProjectionIntents(rec record.Record) []Intent {
	relayID := rec.Envelope.RelayID
	if relayID == "" {
		return nil
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
	render := []byte(fmt.Sprintf("## %s\n\nFROM: %s\nTO: %s\nSUBJECT: %s\n\n%s\n", relayID, rec.Envelope.From, rec.Envelope.To, rec.Headers["SUBJECT"], rec.Body))
	intents := []Intent{
		{Kind: IntentIndex, Path: "INDEX.md", Payload: []byte(fmt.Sprintf("| %s | %s | %s | %s | %s |\n", relayID, rec.Headers["PHASE"], rec.Envelope.From, rec.Envelope.To, rec.Envelope.DeliveryState))},
		{Kind: IntentRender, Path: renderPath, Payload: render},
	}
	if rec.Envelope.To != "" {
		intents = append(intents, Intent{Kind: IntentMailbox, Path: safeMailbox(rec.Envelope.To) + ".jsonl", Payload: []byte(relayID + "\n")})
	}
	return intents
}

func canonicalProjectionIntents(rec record.Record) []Intent {
	intents := DefaultProjectionIntents(rec)
	if rec.Body == "" {
		return intents
	}
	var outbox struct {
		ItemID string `json:"item_id"`
	}
	if err := json.Unmarshal([]byte(rec.Body), &outbox); err != nil || outbox.ItemID == "" {
		return intents
	}
	intents = append(intents, Intent{Kind: IntentOutbox, Path: outbox.ItemID + ".json", Payload: []byte(rec.Body)})
	return intents
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
		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s |\n", row.RelayID, row.Owner, row.Source, row.TargetSurface, row.DispositionPath))
	}
	return Intent{Kind: IntentRender, Path: filepath.Join("owed", "OPEN.md"), Payload: []byte(b.String())}
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
