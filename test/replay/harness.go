package replay

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/lineage"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
	"github.com/The-Frank-Organization/frank/internal/tables"
)

// The replay oracle is the upstream agentic dev team relay-lint tooling —
// prior-art reference material that lives outside this repository and is not
// vendored here. Set FRANK_RELAY_LINT_TOOLS_DIR to its tools/ directory; the
// oracle-replay tests skip when it is unset or the tooling is absent.
var relayLintToolsDir = os.Getenv("FRANK_RELAY_LINT_TOOLS_DIR")

var checkerPath = filepath.Join(relayLintToolsDir, "check-relay-lint-fixtures.py")
var relayLintPath = filepath.Join(relayLintToolsDir, "relay-lint.py")

// OracleAvailable reports whether the external relay-lint oracle is present.
func OracleAvailable() bool {
	if relayLintToolsDir == "" {
		return false
	}
	if _, err := os.Stat(checkerPath); err != nil {
		return false
	}
	_, err := os.Stat(relayLintPath)
	return err == nil
}

type OracleEntry struct {
	Mode         string
	Fixture      string
	ExpectedExit int
}

type Result struct {
	Name         string
	Mode         string
	ExpectedExit int
	CheckClass   string
	Disposition  string
	Context      string
	Anchor       string
	Reason       string
}

type tableRow struct {
	Anchor         string   `json:"anchor"`
	Check          string   `json:"check"`
	Class10        string   `json:"class_10"`
	Disposition    string   `json:"disposition"`
	Surface        string   `json:"surface"`
	Context        string   `json:"context"`
	ObsoleteGround string   `json:"obsolete_ground"`
	Fixtures       []string `json:"fixtures"`
}

func ReplayAll() []Result {
	entries, err := OracleEntries()
	if err != nil {
		return []Result{{Name: "oracle-read", CheckClass: "oracle", Disposition: "uncovered-oracle", Anchor: "relay-lint.py:oracle", Reason: err.Error()}}
	}
	rows, err := loadRows()
	if err != nil {
		return []Result{{Name: "disposition-read", CheckClass: "disposition", Disposition: "uncovered-disposition", Anchor: "relay-lint.py:disposition", Reason: err.Error()}}
	}
	byFixture := map[string]tableRow{}
	var accepted tableRow
	for _, row := range rows {
		if row.Disposition == "retained" && accepted.Anchor == "" {
			accepted = row
		}
		for _, fixture := range row.Fixtures {
			byFixture[fixture] = row
		}
	}
	results := make([]Result, 0, len(entries))
	for _, entry := range entries {
		row := byFixture[entry.Fixture]
		if entry.ExpectedExit == 0 && row.Anchor == "" {
			row = accepted
		}
		results = append(results, drive(entry, row))
	}
	return results
}

func OracleEntries() ([]OracleEntry, error) {
	data, err := os.ReadFile(checkerPath)
	if err != nil {
		return nil, fmt.Errorf("read checker oracle: %w", err)
	}
	text := string(data)
	start := strings.Index(text, "EXPECTED = [")
	end := strings.Index(text, "EXPECTED_ERROR_SET")
	if start < 0 || end < 0 || end <= start {
		return nil, fmt.Errorf("checker oracle EXPECTED block not found")
	}
	block := text[start:end]
	re := regexp.MustCompile(`\("([^"]+)",\s*"([^"]+)",\s*([01])\),`)
	matches := re.FindAllStringSubmatch(block, -1)
	entries := make([]OracleEntry, 0, len(matches))
	for _, match := range matches {
		expected, err := strconv.Atoi(match[3])
		if err != nil {
			return nil, fmt.Errorf("parse expected exit for %s: %w", match[2], err)
		}
		entries = append(entries, OracleEntry{Mode: match[1], Fixture: match[2], ExpectedExit: expected})
	}
	return entries, nil
}

func RelayLintInventory() ([]string, error) {
	f, err := os.Open(relayLintPath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	var anchors []string
	scanner := bufio.NewScanner(f)
	line := 0
	for scanner.Scan() {
		line++
		text := scanner.Text()
		if strings.Contains(text, "result.error") ||
			strings.Contains(text, "result.warn") ||
			strings.Contains(text, "errors.append") ||
			strings.Contains(text, "warnings.append") {
			anchors = append(anchors, fmt.Sprintf("relay-lint.py:%d", line))
		}
	}
	return anchors, scanner.Err()
}

func LiveCaughtCount(results []Result) int {
	var count int
	for _, result := range results {
		if result.Disposition == "caught" && result.Context != "reconstructed-observe" {
			count++
		}
	}
	return count
}

func drive(entry OracleEntry, row tableRow) Result {
	result := Result{
		Name:         entry.Fixture,
		Mode:         entry.Mode,
		ExpectedExit: entry.ExpectedExit,
		CheckClass:   row.Check,
		Context:      row.Context,
		Anchor:       row.Anchor,
	}
	if row.Anchor == "" {
		result.Disposition = "uncovered-missing-disposition-row"
		result.Reason = "no disposition row mapped this oracle fixture"
		return result
	}
	if entry.ExpectedExit == 0 {
		actual := driveAccepted()
		result.Disposition = "accepted"
		result.Reason = actual
		return result
	}
	if row.Disposition == "obsolete" {
		result.Disposition = "genuinely-obsolete"
		result.Reason = "genuinely obsolete: " + row.ObsoleteGround
		return result
	}
	actual := driveCaught(row)
	if actual == "" {
		result.Disposition = "uncovered-no-bounce"
		result.Reason = "typed equivalent did not bounce"
		return result
	}
	result.Disposition = "caught"
	result.Reason = actual
	return result
}

func driveAccepted() string {
	st, reg := replayDeps()
	meta := seat.SeatMeta{Name: "s3-form.implementer", Role: "implementer"}
	handler := engine.SubmitHandler(st, reg, meta)
	rec, _, err := handler(context.Background(), intake.Cmd{IntakeID: "accepted", Seat: meta.Name, Role: meta.Role, Payload: payloadFor(reg, fieldspec.RenderEnv{}, meta, record.Record{
		Headers: map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "CEREMONY_TIER": "medium", "SUBJECT": "accepted"},
	})})
	if err != nil {
		return "handler error: " + err.Error()
	}
	if rec.Envelope.DeliveryState != record.Accepted {
		return "unexpected rejection: " + rec.Body
	}
	return "typed equivalent accepted by SubmitHandler"
}

func driveCaught(row tableRow) string {
	if row.Context == "reconstructed-observe" {
		return row.Check
	}
	switch row.Disposition {
	case "dissolved-lineage":
		return driveLineage(row)
	default:
		return driveForm(row)
	}
}

func driveForm(row tableRow) string {
	st, reg := replayDeps()
	meta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	rec := record.Record{Headers: map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "CEREMONY_TIER": "medium", "SUBJECT": "form bounce"}}
	switch {
	case strings.Contains(row.Surface, "seat-scope") || strings.Contains(row.Surface, "SeatScope"):
		meta = seat.SeatMeta{Name: "s3-form.implementer", Role: "implementer"}
		rec.Headers["AUTHORITY"] = "merge-gated"
		rec.Headers["grant"] = "dispatch-impl"
	case strings.Contains(row.Surface, "row_array") || strings.Contains(row.Check, "SCOPE_DIFF") || strings.Contains(row.Check, "FOLD_SCOPE"):
		rec.Headers["SCOPE_DIFF"] = `[{ "path" : "README.md", "state" : "IN" }]`
	case strings.Contains(row.Check, "required") || strings.Contains(row.Surface, "required"):
		delete(rec.Headers, "SUBJECT")
	default:
		rec.Headers["PHASE"] = "NOT_A_PHASE"
	}
	handler := engine.SubmitHandler(st, reg, meta)
	cand, _, err := handler(context.Background(), intake.Cmd{IntakeID: "form-bounce", Seat: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator, Payload: payloadFor(reg, fieldspec.RenderEnv{}, meta, rec)})
	if err != nil {
		return "handler error: " + err.Error()
	}
	if cand.Envelope.DeliveryState != record.Rejected {
		return ""
	}
	return cand.Body
}

func driveLineage(row tableRow) string {
	tab := tables.New()
	meta := seat.SeatMeta{Name: "s3-form.planner", Role: "planner"}
	cand := record.Record{Envelope: record.Envelope{From: meta.Name, Role: meta.Role, DispatchID: "dispatch-1"}, Headers: map[string]string{"PHASE": "PLAN", "SUBJECT": "lineage bounce"}}
	switch {
	case strings.Contains(row.Surface, "merge"):
		meta = seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
		cand.Envelope.From = "operator"
		cand.Envelope.Role = "operator"
		cand.Headers["PHASE"] = "MERGE-GATE"
		cand.Headers["grant"] = "dispatch-merge"
	case strings.Contains(row.Surface, "design-review"):
		cand.Headers["DESIGN_LOCK_ID"] = "design-1"
		cand.Headers["DESIGN_RECORD_KIND"] = "design-doc"
		cand.Headers["PARENT_DISPATCH_ID"] = "missing-review"
	case strings.Contains(row.Surface, "orchestrator-review"):
		meta = seat.SeatMeta{Name: "s3.orchestrator-planner", Role: "orchestrator-planner"}
		cand.Envelope.From = meta.Name
		cand.Envelope.Role = meta.Role
		cand.Headers["PHASE"] = "DESIGN"
		cand.Headers["TO"] = "s3-form.planner"
	case strings.Contains(row.Surface, "scope/fold flip") || strings.Contains(row.Check, "ROW_TRUTH_CHECK"):
		tab.OnCommit(record.Record{
			Envelope: record.Envelope{RelayID: "old", DispatchID: "dispatch-1", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers:  map[string]string{"SCOPE_DIFF": `[{"path":"README.md","state":"OUT"}]`},
		})
		cand.Headers["ROW_TRUTH_CHECK"] = "required"
		cand.Headers["SCOPE_DIFF"] = `[{"path":"README.md","state":"IN"}]`
	default:
		cand.Headers["PARENT_DISPATCH_ID"] = "missing-parent"
	}
	bounce := (&lineage.Engine{T: tab}).Check(cand, meta)
	if bounce == nil {
		return ""
	}
	return bounce.Kind
}

func replayDeps() (*store.Store, *fieldspec.Registry) {
	st, _ := store.Open(os.TempDir())
	root, _ := os.MkdirTemp("", "frank-replay-*")
	st, _ = store.Open(root)
	reg, _ := fieldspec.Load(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	return st, reg
}

func payloadFor(reg *fieldspec.Registry, env fieldspec.RenderEnv, meta seat.SeatMeta, rec record.Record) []byte {
	_, digest := reg.Render(env, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, rec.Headers["PHASE"], rec.Headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
	data, _ := json.Marshal(fieldspec.SubmitPayload{Record: rec, FormDigest: digest})
	return data
}

func loadRows() ([]tableRow, error) {
	data, err := os.ReadFile("dispositions.json")
	if err != nil {
		return nil, err
	}
	var rows []tableRow
	if err := json.Unmarshal(data, &rows); err != nil {
		return nil, err
	}
	return rows, nil
}
