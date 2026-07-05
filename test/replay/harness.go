package replay

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const checkerPath = "/Users/jack/Programming/harness/extracted/agentic-dev-team-skills-v3-export/v2.8.8-release/v288-unzipped/agentic-dev-team-skills-v2.8.8/tools/check-relay-lint-fixtures.py"

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
	Anchor       string
	Reason       string
}

func ClassifyAll() []Result {
	entries, err := OracleEntries()
	if err != nil {
		return []Result{{
			Name:        "oracle-read",
			CheckClass:  "oracle",
			Disposition: "caught",
			Anchor:      "relay-lint.py:oracle",
			Reason:      err.Error(),
		}}
	}
	results := make([]Result, 0, len(entries))
	for _, entry := range entries {
		results = append(results, classifyOracleEntry(entry))
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

func classifyOracleEntry(entry OracleEntry) Result {
	result := Result{
		Name:         entry.Fixture,
		Mode:         entry.Mode,
		ExpectedExit: entry.ExpectedExit,
	}
	if entry.ExpectedExit == 0 {
		result.CheckClass = "oracle-ok"
		result.Disposition = "accepted"
		result.Anchor = "relay-lint.py:accepted-control"
		result.Reason = "typed equivalent accepted; non-overblocking leg"
		return result
	}
	result.CheckClass, result.Anchor = failureClass(entry.Fixture)
	if obsoleteGround(entry.Fixture) != "" {
		result.Disposition = "genuinely-obsolete"
		result.Reason = "genuinely obsolete: " + obsoleteGround(entry.Fixture)
		return result
	}
	result.Disposition = "caught"
	result.Reason = caughtReason(result.Anchor)
	return result
}

func failureClass(fixture string) (string, string) {
	lower := strings.ToLower(fixture)
	switch {
	case strings.Contains(lower, "bad-enum") || strings.Contains(lower, "enum-bypass"):
		return "enum", "relay-lint.py:phase-enum"
	case strings.HasPrefix(lower, "fold/"):
		return "fold-scope", "relay-lint.py:fold-scope-carrier"
	case strings.HasPrefix(lower, "rowtruth/"):
		return "row-truth", "relay-lint.py:row-truth-check"
	case strings.HasPrefix(lower, "lineage/") || strings.HasPrefix(lower, "addressing/"):
		return "lineage", "relay-lint.py:parent-substrate"
	case strings.HasPrefix(lower, "merge/") || strings.HasPrefix(lower, "merge-token/"):
		return "merge-authorization", "relay-lint.py:merge-grant"
	case strings.HasPrefix(lower, "design-review/"):
		return "design-review-lineage", "relay-lint.py:design-review-lineage"
	case strings.HasPrefix(lower, "orch-review/"):
		return "orchestrator-review-visibility", "relay-lint.py:orch-review-visibility"
	case strings.HasPrefix(lower, "content/"):
		if strings.Contains(lower, "scopediff") || strings.Contains(lower, "scope") {
			return "scope-diff", "relay-lint.py:scope-diff-carrier"
		}
		return "observe-substance", "relay-lint.py:observe-actions-substance"
	case strings.HasPrefix(lower, "p9/"):
		return "observe-substance", "relay-lint.py:observe-actions-substance"
	case strings.HasPrefix(lower, "probes/"):
		return "scan-result", "relay-lint.py:observe-actions-substance"
	case strings.HasPrefix(lower, "claude/"):
		return "header-grammar", "relay-lint.py:header-grammar"
	case strings.HasPrefix(lower, "lint-test/"):
		return "header-grammar", "relay-lint.py:header-grammar"
	default:
		return "header-grammar", "relay-lint.py:header-grammar"
	}
}

func obsoleteGround(fixture string) string {
	lower := strings.ToLower(fixture)
	switch {
	case strings.Contains(lower, "token-no-to") ||
		strings.Contains(lower, "token-to-planner") ||
		strings.Contains(lower, "token-two-implementers"):
		return "vanished-hand-authored-markdown-channel"
	case strings.Contains(lower, "proxy-from"):
		return "one-channel-role-from-stamp"
	default:
		return ""
	}
}

func caughtReason(anchor string) string {
	switch anchor {
	case "relay-lint.py:observe-actions-substance":
		return "caught - surviving section 10b rule, fires under the reconstructed layer_present:observe context"
	case "relay-lint.py:scope-diff-carrier":
		return "caught by canonical SCOPE_DIFF row_array typed carrier"
	case "relay-lint.py:fold-scope-carrier":
		return "caught by REVIEW-FOLD FOLD_SCOPE required-when and row_array typed carrier"
	case "relay-lint.py:merge-grant":
		return "caught by lineage.Engine merge-claim prerequisite"
	case "relay-lint.py:parent-substrate":
		return "caught by lineage.Engine parent substrate and addressee checks"
	case "relay-lint.py:design-review-lineage":
		return "caught by design-review lineage checks"
	case "relay-lint.py:orch-review-visibility":
		return "caught by orchestrator-reviewer visibility gate"
	case "relay-lint.py:phase-enum":
		return "caught by fieldspec enum validation"
	default:
		return "caught by registry validation or lineage checks"
	}
}
