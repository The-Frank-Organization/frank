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
	PathHygiene      pathHygiene      `json:"path_hygiene"`
	Laws             []catalogLaw     `json:"laws"`
}

type changeConvention struct {
	SingleWriter  string `json:"single_writer"`
	OwnerFidelity string `json:"owner_fidelity"`
	Section7Claim string `json:"section_7_claim"`
	Section7Carry string `json:"section_7_carry"`
}

type pathHygiene struct {
	Families     []outputFamily `json:"families"`
	SinkPatterns []sinkPattern  `json:"sink_patterns"`
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
	if cat.Version != "s7-v1" || cat.Status != "convention-only" {
		t.Fatalf("catalog version/status = %q/%q, want s7-v1/convention-only", cat.Version, cat.Status)
	}
	if cat.ChangeConvention.SingleWriter == "" || cat.ChangeConvention.OwnerFidelity == "" {
		t.Fatalf("catalog change convention incomplete: %+v", cat.ChangeConvention)
	}
	if cat.ChangeConvention.Section7Claim != "not-claimed-in-s7" ||
		cat.ChangeConvention.Section7Carry != "s8-section-7-digest-pinning" {
		t.Fatalf("catalog section-7 staging = %+v", cat.ChangeConvention)
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
