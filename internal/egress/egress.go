// Package egress contains the Step-1 outbox drain and scanner.
//
// PRESENT-BUT-DORMANT: no production package imports this package in Step-1.
// Activation belongs to the away-mode bridge. Detection is best-effort and
// encoding-evadable; the outbox is the only egress path the governance system
// offers. Origin values are runtime-only renderer facts, never serialized or
// parsed from record bytes.
package egress

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sort"

	"github.com/jackli/frank/internal/gate"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

type Class string

const (
	ClassConfidentiality Class = "confidentiality"
	ClassSafety          Class = "safety"
	ClassOther           Class = "other"

	DispositionReady   = "egress_ready"
	DispositionBlocked = "egress_blocked"
)

type Finding struct {
	Field string
	Class Class
}

func (f Finding) String() string {
	return f.Field + ":" + string(f.Class)
}

type Verdict struct {
	Blocked  bool
	Findings []Finding
}

type Origin struct {
	ConductorODB bool
}

type Item struct {
	Meta   gate.OutboxItem
	Source record.Record
	Dest   string
	Field  RenderedField
}

type Result struct {
	ItemID          string
	SourceRecordRef string
	Dest            string
	Disposition     string
	Findings        []Finding
}

type Report struct {
	Results []Result
}

func Scan(item Item, rules Rules) Verdict {
	rules = rules.withDefaults()
	var verdict Verdict
	for _, classify := range rules.Classifiers {
		for _, finding := range classify(item) {
			verdict.Findings = append(verdict.Findings, finding)
			if !findingAllowed(item, finding) {
				verdict.Blocked = true
			}
		}
	}
	return verdict
}

func findingAllowed(item Item, finding Finding) bool {
	return finding.Class == ClassConfidentiality &&
		item.Field.Name == "model_name" &&
		item.Field.Origin.ConductorODB &&
		item.Dest == "operator"
}

var (
	errOutboxRead  = errors.New("outbox-read-error")
	errItemRead    = errors.New("outbox-item-read-error")
	errItemDecode  = errors.New("outbox-item-decode-error")
	errSourceRead  = errors.New("source-record-read-error")
	errRender      = errors.New("egress-render-error")
	errEmptyItemID = errors.New("outbox-item-id-empty")
)

func Drain(st *store.Store, rules Rules, render Renderer) (Report, error) {
	if render == nil {
		render = DefaultRenderer
	}
	entries, err := os.ReadDir(filepath.Join(st.Root, "outbox"))
	if err != nil {
		return Report{}, errOutboxRead
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })

	var report Report
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		meta, err := readOutboxItem(filepath.Join(st.Root, "outbox", entry.Name()))
		if err != nil {
			return Report{}, err
		}
		source, err := st.Read(meta.SourceRecordRef)
		if err != nil {
			return Report{}, errSourceRead
		}
		rendered, err := render(meta, source)
		if err != nil {
			return Report{}, errRender
		}
		result := Result{
			ItemID:          meta.ItemID,
			SourceRecordRef: meta.SourceRecordRef,
			Dest:            rendered.Dest,
			Disposition:     DispositionReady,
		}
		for _, field := range rendered.Fields {
			item := Item{Meta: meta, Source: source, Dest: rendered.Dest, Field: field}
			verdict := Scan(item, rules)
			result.Findings = append(result.Findings, verdict.Findings...)
			if verdict.Blocked {
				result.Disposition = DispositionBlocked
			}
		}
		report.Results = append(report.Results, result)
	}
	return report, nil
}

func readOutboxItem(path string) (gate.OutboxItem, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return gate.OutboxItem{}, errItemRead
	}
	var item gate.OutboxItem
	if err := json.Unmarshal(data, &item); err != nil {
		return gate.OutboxItem{}, errItemDecode
	}
	if item.ItemID == "" {
		return gate.OutboxItem{}, errEmptyItemID
	}
	return item, nil
}
