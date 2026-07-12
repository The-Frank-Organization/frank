package invariants_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"testing"
)

type catalog struct {
	Version          string           `json:"version"`
	Status           string           `json:"status"`
	ChangeConvention changeConvention `json:"change_convention"`
	Discovery        discovery        `json:"discovery"`
	Laws             []catalogLaw     `json:"laws"`
}

type changeConvention struct {
	SingleWriter  string `json:"single_writer"`
	OwnerFidelity string `json:"owner_fidelity"`
	Section7Claim string `json:"section_7_claim"`
	Section7Carry string `json:"section_7_carry"`
}

type discovery struct {
	Scan                  discoveryScan         `json:"scan"`
	BoundaryFiles         []string              `json:"boundary_files"`
	Recognizers           discoveryRecognizers  `json:"recognizers"`
	SiteCensus            []string              `json:"site_census"`
	OutputFamilies        []outputFamily        `json:"output_families"`
	SinkPatterns          []sinkPattern         `json:"sink_patterns"`
	CanonicalPathFamilies canonicalPathFamilies `json:"canonical_path_families"`
}

type discoveryScan struct {
	Root                string   `json:"root"`
	ExcludeDirPrefixes  []string `json:"exclude_dir_prefixes"`
	ExcludeFileSuffixes []string `json:"exclude_file_suffixes"`
	IncludeOnlySuffixes []string `json:"include_only_suffixes"`
}

type discoveryContext struct {
	PathPrefix        string   `json:"path_prefix"`
	CallSelectors     []string `json:"call_selectors"`
	IdentCalls        []string `json:"ident_calls,omitempty"`
	ConnReceiverCalls []string `json:"conn_receiver_calls,omitempty"`
}

type discoveryRecognizers struct {
	ChannelContext            discoveryContext `json:"channel_context"`
	MCPContext                discoveryContext `json:"mcp_context"`
	ProtocolSwitchTagSelector string           `json:"protocol_switch_tag_selector"`
	TreeWideIdioms            []string         `json:"tree_wide_idioms"`
}

type canonicalPathFamilies struct {
	Rows []canonicalPathFamily `json:"rows"`
}

type outputFamily struct {
	ID        string   `json:"id"`
	CarveOuts []string `json:"carve_outs"`
}

type sinkPattern struct {
	ID            string `json:"id"`
	Family        string `json:"family"`
	Symbol        string `json:"symbol"`
	ExpectedSites int    `json:"expected_sites"`
}

type catalogLaw struct {
	ID             string   `json:"id"`
	Test           string   `json:"test"`
	Owners         []string `json:"owners"`
	FidelityReview []string `json:"fidelity_review"`
	Claim          string   `json:"claim"`
}

func requireCatalogLaw(t *testing.T, testName string) (catalog, catalogLaw) {
	t.Helper()
	cat := loadCatalog(t)
	validateCatalog(t, cat)
	for _, law := range cat.Laws {
		if law.Test == testName {
			return cat, law
		}
	}
	t.Fatalf("catalog missing law for %s", testName)
	return catalog{}, catalogLaw{}
}

func loadCatalog(t *testing.T) catalog {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(repoRoot(t), "test", "invariants", "catalog.v1.json"))
	if err != nil {
		t.Fatalf("read catalog: %v", err)
	}
	var cat catalog
	if err := json.Unmarshal(data, &cat); err != nil {
		t.Fatalf("decode catalog: %v", err)
	}
	return cat
}

func validateCatalog(t *testing.T, cat catalog) {
	t.Helper()
	if cat.Version != "s8-v1" || cat.Status != "section7-pinned" {
		t.Fatalf("catalog version/status = %q/%q, want s8-v1/section7-pinned", cat.Version, cat.Status)
	}
	if cat.ChangeConvention.SingleWriter == "" || cat.ChangeConvention.OwnerFidelity == "" {
		t.Fatalf("catalog change convention incomplete: %+v", cat.ChangeConvention)
	}
	if cat.ChangeConvention.Section7Claim != "pinned-at-s8: load enforces digest and member/provenance shape; owner review remains a relay/design-review gate" ||
		cat.ChangeConvention.Section7Carry != "discharged-at-s8" {
		t.Fatalf("catalog section-7 staging = %+v", cat.ChangeConvention)
	}
	if len(cat.Discovery.SiteCensus) != 17 || len(cat.Discovery.OutputFamilies) != 6 || len(cat.Discovery.SinkPatterns) != 6 || len(cat.Discovery.CanonicalPathFamilies.Rows) != 12 {
		t.Fatalf("catalog discovery cardinalities = sites:%d families:%d sinks:%d paths:%d", len(cat.Discovery.SiteCensus), len(cat.Discovery.OutputFamilies), len(cat.Discovery.SinkPatterns), len(cat.Discovery.CanonicalPathFamilies.Rows))
	}
	if cat.Discovery.Scan.Root != "." || !reflect.DeepEqual(cat.Discovery.Scan.ExcludeDirPrefixes, []string{"."}) || !reflect.DeepEqual(cat.Discovery.Scan.ExcludeFileSuffixes, []string{"_test.go"}) || !reflect.DeepEqual(cat.Discovery.Scan.IncludeOnlySuffixes, []string{".go"}) {
		t.Fatalf("catalog scan descriptor = %+v", cat.Discovery.Scan)
	}
	wantTests := []string{
		"TestLawCanonicalWins",
		"TestLawDerivedOnlyActivation",
		"TestLawIntakeOutcomeOneToOne",
		"TestLawOnePivotPerMutation",
		"TestLawPathHygiene",
		"TestLawR2NoModelPredicate",
		"TestLawRebuildBeforeOpen",
		"TestLawSoleGovernedWriter",
		"TestLawTerminalEnumByteExact",
		"TestLawThreeVerbSurface",
	}
	var gotTests []string
	seenIDs := map[string]bool{}
	for _, law := range cat.Laws {
		if law.ID == "" || law.Test == "" || law.Claim == "" || len(law.Owners) == 0 {
			t.Fatalf("incomplete law row: %+v", law)
		}
		if seenIDs[law.ID] {
			t.Fatalf("duplicate law id %q", law.ID)
		}
		seenIDs[law.ID] = true
		gotTests = append(gotTests, law.Test)
	}
	sort.Strings(gotTests)
	if !reflect.DeepEqual(gotTests, wantTests) {
		t.Fatalf("catalog tests = %v, want %v", gotTests, wantTests)
	}
}

func repoRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
}
