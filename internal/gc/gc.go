package gc

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/The-Frank-Organization/frank/internal/config"
	"github.com/The-Frank-Organization/frank/internal/crashpoint"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/obligation"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func Drained(segs []intake.Segment, tables *obligation.Tables) []intake.Segment {
	if tables == nil || len(segs) == 0 {
		return nil
	}
	activeSeq := segs[0].Seq
	for _, seg := range segs {
		if seg.Seq > activeSeq {
			activeSeq = seg.Seq
		}
	}
	var drained []intake.Segment
	for _, seg := range segs {
		if seg.Seq == activeSeq {
			continue
		}
		ok := true
		for _, entry := range seg.Entries {
			if _, exists := tables.OutcomeByIntake[entry.IntakeID]; !exists {
				ok = false
				break
			}
		}
		if ok {
			drained = append(drained, seg)
		}
	}
	return drained
}

func Pass(st *store.Store, tables *obligation.Tables, cfg config.EngineConfig) error {
	if !cfg.GCEnabled {
		return nil
	}
	journal, err := intake.OpenWithConfig(st.Root, cfg)
	if err != nil {
		return err
	}
	segments, err := journal.Segments()
	if err != nil {
		return err
	}
	drained := Drained(segments, tables)
	if len(drained) == 0 {
		return nil
	}
	names := make([]string, 0, len(drained))
	for _, seg := range drained {
		names = append(names, filepath.Base(seg.Path))
	}
	marked, err := markerCovers(st, names)
	if err != nil {
		return err
	}
	if !marked {
		body, err := json.Marshal(struct {
			Segments []string `json:"segments"`
		}{Segments: names})
		if err != nil {
			return err
		}
		crashpoint.Hit("pre_gc_marker")
		_, err = st.Commit(record.Record{
			Envelope: record.Envelope{
				RelayID:       fmt.Sprintf("gc-%06d", drained[0].Seq),
				DispatchID:    "gc",
				From:          "system",
				Role:          "system",
				DeliveryState: record.Accepted,
				SchemaVersion: 1,
			},
			Headers: map[string]string{"record_kind": "gc_marker"},
			Body:    string(body),
		}, nil)
		if err != nil {
			return err
		}
	}
	crashpoint.Hit("post_gc_marker")
	for _, seg := range drained {
		crashpoint.Hit("pre_gc_unlink")
		if err := os.Remove(seg.Path); err != nil && !os.IsNotExist(err) {
			return err
		}
		crashpoint.Hit("post_gc_unlink")
	}
	return nil
}

func markerCovers(st *store.Store, names []string) (bool, error) {
	if len(names) == 0 {
		return true, nil
	}
	records, err := st.Records()
	if err != nil {
		return false, err
	}
	marked := map[string]bool{}
	for _, rec := range records {
		if rec.Headers["record_kind"] != "gc_marker" {
			continue
		}
		var body struct {
			Segments []string `json:"segments"`
		}
		if err := json.Unmarshal([]byte(rec.Body), &body); err != nil {
			continue
		}
		for _, segment := range body.Segments {
			marked[segment] = true
		}
	}
	for _, name := range names {
		if !marked[name] {
			return false, nil
		}
	}
	return true, nil
}
