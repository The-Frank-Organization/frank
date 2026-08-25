package fixtures_test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"sort"
	"strings"
	"testing"
)

const (
	s12DogfoodNestedEnv          = "FRANK_DOGFOOD_NESTED"
	s12NestedAttestation         = "nested-green attests the courier's gate adjudication over the nested suite; the process-level crash/ceremony cuts are proven ONLY by the outer run."
	s12OuterReachabilityArgument = "the exit-gate test carries the same FRANK_DOGFOOD_NESTED guard; reaching its outer body proves the variable was unset in the outer invocation's shared parent environment and licenses every same-guard fixture's outer execution"
)

// s12OuterRunOnlyFixtures is the authoritative nested narrowing. Each row names
// a process-level proof that the outer run already executes against real bytes.
var s12OuterRunOnlyFixtures = map[string]string{
	"TestH16CeremonyConflictShapesRetainInBandOperatorPaths":             "starts a real conductor to prove in-band operator recovery paths",
	"TestH16CeremonyConflictSelectsOnceRealizesAndRefusesReplay":         "runs ceremony subprocesses to prove realization and replay refusal",
	"TestH16CeremonyStateActionMatrixResolvedBranches":                   "runs ceremony subprocesses across the resolved-state action matrix",
	"TestH16CeremonyCustodyAuthorityIgnoresDisputedRoleBits":             "runs ceremony subprocesses to prove canonical custody authority",
	"TestH16CeremonyLockFirstAliasAndSocketDiagnostic":                   "runs ceremony subprocesses against real root locks and sockets",
	"TestH16CeremonyConcurrentProcessesHaveOneLockWinner":                "races two real ceremony processes to prove single lock ownership",
	"TestH16CeremonyCrashCutsConvergeWithoutDuplicateAnchor":             "crashes and restarts real ceremony processes at three cut points",
	"TestH16CeremonyRetryCreatesDistinctCanonicalRotationEachTime":       "runs repeated retry subprocesses to prove credential rotation",
	"TestH16CeremonyRetryInterruptedPivotResumesOrdinaryMarkerMachinery": "crashes and resumes a real retry subprocess around pivot persistence",
	"TestH16CeremonyRetryScopeIsCanonicalAndFailClosed":                  "runs retry subprocesses across canonical and refused scopes",
	"TestH16CeremonyRetryRejectsAuthorityFieldsAtIntake":                 "runs retry subprocesses to prove authority override refusal",
	"TestH16CeremonyRetryReplyCrashRotatesAgainWithoutCredentialReuse":   "crashes real retry subprocesses across reply delivery cuts",
	"TestH16StartupFoldsRealizedMintEvidenceBeforeServing":               "starts and restarts real conductors across startup evidence cuts",
}

func s12DogfoodNested() bool {
	return os.Getenv(s12DogfoodNestedEnv) != ""
}

func s12DogfoodOuterOnlyDecision(name string, nested bool) (bool, string) {
	if !nested {
		return false, ""
	}
	why, ok := s12OuterRunOnlyFixtures[name]
	return ok, why
}

func s12SkipOuterRunOnly(t *testing.T) bool {
	t.Helper()
	name, _, _ := strings.Cut(t.Name(), "/")
	skip, why := s12DogfoodOuterOnlyDecision(name, s12DogfoodNested())
	if !skip {
		return false
	}
	t.Skipf("%s outer-run-only reason: %s", s12NestedAttestation, why)
	return true
}

func s12AssertOuterDogfoodReachability(t *testing.T) {
	t.Helper()
	if s12DogfoodNested() {
		t.Fatalf("exit-gate outer body reached with %s set", s12DogfoodNestedEnv)
	}
	t.Log(s12OuterReachabilityArgument)
}

func s12InspectOuterRunOnlyBindings(sources map[string]string, registry map[string]string) error {
	paths := make([]string, 0, len(sources))
	for path := range sources {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	found := make(map[string]string, len(registry))
	for _, path := range paths {
		file, err := parser.ParseFile(token.NewFileSet(), path, sources[path], 0)
		if err != nil {
			return fmt.Errorf("parse %s: %w", path, err)
		}
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Body == nil {
				continue
			}
			name := fn.Name.Name
			if _, registered := registry[name]; !registered {
				continue
			}
			if firstPath := found[name]; firstPath != "" {
				return fmt.Errorf("%s: duplicate test bodies in %s and %s", name, firstPath, path)
			}
			found[name] = path

			guardIndex := 0
			if name == "TestH16StartupFoldsRealizedMintEvidenceBeforeServing" {
				if len(fn.Body.List) == 0 || !s12IsStartupCrashChildPrefix(fn.Body.List[0]) {
					return fmt.Errorf("%s: missing startup crash-child prefix in %s", name, path)
				}
				guardIndex = 1
			}
			if len(fn.Body.List) <= guardIndex || !s12IsOuterRunOnlyGuard(fn.Body.List[guardIndex]) {
				return fmt.Errorf("%s: missing executable s12SkipOuterRunOnly(t) guard in %s", name, path)
			}
		}
	}

	names := make([]string, 0, len(registry))
	for name := range registry {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		if found[name] == "" {
			return fmt.Errorf("%s: registered outer-run-only test body not found", name)
		}
	}
	return nil
}

func s12IsOuterRunOnlyGuard(stmt ast.Stmt) bool {
	guard, ok := stmt.(*ast.IfStmt)
	if !ok || guard.Init != nil || guard.Else != nil || len(guard.Body.List) != 1 {
		return false
	}
	call, ok := guard.Cond.(*ast.CallExpr)
	if !ok || len(call.Args) != 1 {
		return false
	}
	helper, ok := call.Fun.(*ast.Ident)
	if !ok || helper.Name != "s12SkipOuterRunOnly" {
		return false
	}
	argument, argumentOK := call.Args[0].(*ast.Ident)
	ret, returnOK := guard.Body.List[0].(*ast.ReturnStmt)
	return argumentOK && argument.Name == "t" && returnOK && len(ret.Results) == 0
}

func s12IsStartupCrashChildPrefix(stmt ast.Stmt) bool {
	crashGuard, ok := stmt.(*ast.IfStmt)
	if !ok || crashGuard.Init != nil || crashGuard.Else != nil || len(crashGuard.Body.List) != 2 {
		return false
	}
	comparison, ok := crashGuard.Cond.(*ast.BinaryExpr)
	if !ok || comparison.Op != token.EQL {
		return false
	}
	getenv, ok := comparison.X.(*ast.CallExpr)
	if !ok || len(getenv.Args) != 1 {
		return false
	}
	selector, ok := getenv.Fun.(*ast.SelectorExpr)
	if !ok || selector.Sel.Name != "Getenv" {
		return false
	}
	packageName, packageOK := selector.X.(*ast.Ident)
	envName, envOK := getenv.Args[0].(*ast.BasicLit)
	wantValue, valueOK := comparison.Y.(*ast.BasicLit)
	if !packageOK || packageName.Name != "os" || !envOK || envName.Kind != token.STRING || envName.Value != `"FRANK_H16_MINT_CRASH_CHILD"` || !valueOK || wantValue.Kind != token.STRING || wantValue.Value != `"1"` {
		return false
	}
	callStmt, ok := crashGuard.Body.List[0].(*ast.ExprStmt)
	if !ok {
		return false
	}
	crashCall, ok := callStmt.X.(*ast.CallExpr)
	if !ok || len(crashCall.Args) != 1 {
		return false
	}
	crashHelper, ok := crashCall.Fun.(*ast.Ident)
	if !ok || crashHelper.Name != "h16CrashAfterMintPersistence" {
		return false
	}
	argument, argumentOK := crashCall.Args[0].(*ast.Ident)
	ret, returnOK := crashGuard.Body.List[1].(*ast.ReturnStmt)
	return argumentOK && argument.Name == "t" && returnOK && len(ret.Results) == 0
}

func TestS12OuterRunOnlyRegistryBindsNamedTestBodies(t *testing.T) {
	sources := make(map[string]string)
	for _, path := range []string{"ceremony_test.go", "ceremony_retry_test.go", "h16_startup_evidence_test.go"} {
		raw, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}
		sources[path] = string(raw)
	}
	if err := s12InspectOuterRunOnlyBindings(sources, s12OuterRunOnlyFixtures); err != nil {
		t.Fatal(err)
	}
}

func TestS12OuterRunOnlyRegistryBindingRejectsMissingCall(t *testing.T) {
	const source = `package fixtures_test

import "testing"

func TestSyntheticOuterOnly(t *testing.T) {
	t.Log("shared helper call intentionally absent")
}
`
	err := s12InspectOuterRunOnlyBindings(
		map[string]string{"synthetic_test.go": source},
		map[string]string{"TestSyntheticOuterOnly": "synthetic negative proof"},
	)
	if err == nil {
		t.Fatal("missing s12SkipOuterRunOnly(t) call was accepted")
	}
	if !strings.Contains(err.Error(), "TestSyntheticOuterOnly: missing executable s12SkipOuterRunOnly(t) guard") {
		t.Fatalf("missing-call error = %q", err)
	}
}

func TestS12OuterRunOnlyRegistryBindingRejectsNestedIneffectiveCall(t *testing.T) {
	const source = `package fixtures_test

import "testing"

func TestSyntheticOuterOnly(t *testing.T) {
	_ = func() {
		if s12SkipOuterRunOnly(t) {
			return
		}
	}
	t.Log("outer body remains unguarded")
}
`
	err := s12InspectOuterRunOnlyBindings(
		map[string]string{"synthetic_test.go": source},
		map[string]string{"TestSyntheticOuterOnly": "synthetic nested-occurrence negative proof"},
	)
	if err == nil {
		t.Fatal("nested ineffective s12SkipOuterRunOnly(t) call was accepted")
	}
	if !strings.Contains(err.Error(), "TestSyntheticOuterOnly: missing executable s12SkipOuterRunOnly(t) guard") {
		t.Fatalf("nested-call error = %q", err)
	}
}

func TestS12DogfoodNestedOuterOnlyContract(t *testing.T) {
	want := map[string]string{
		"TestH16CeremonyConflictShapesRetainInBandOperatorPaths":             "starts a real conductor to prove in-band operator recovery paths",
		"TestH16CeremonyConflictSelectsOnceRealizesAndRefusesReplay":         "runs ceremony subprocesses to prove realization and replay refusal",
		"TestH16CeremonyStateActionMatrixResolvedBranches":                   "runs ceremony subprocesses across the resolved-state action matrix",
		"TestH16CeremonyCustodyAuthorityIgnoresDisputedRoleBits":             "runs ceremony subprocesses to prove canonical custody authority",
		"TestH16CeremonyLockFirstAliasAndSocketDiagnostic":                   "runs ceremony subprocesses against real root locks and sockets",
		"TestH16CeremonyConcurrentProcessesHaveOneLockWinner":                "races two real ceremony processes to prove single lock ownership",
		"TestH16CeremonyCrashCutsConvergeWithoutDuplicateAnchor":             "crashes and restarts real ceremony processes at three cut points",
		"TestH16CeremonyRetryCreatesDistinctCanonicalRotationEachTime":       "runs repeated retry subprocesses to prove credential rotation",
		"TestH16CeremonyRetryInterruptedPivotResumesOrdinaryMarkerMachinery": "crashes and resumes a real retry subprocess around pivot persistence",
		"TestH16CeremonyRetryScopeIsCanonicalAndFailClosed":                  "runs retry subprocesses across canonical and refused scopes",
		"TestH16CeremonyRetryRejectsAuthorityFieldsAtIntake":                 "runs retry subprocesses to prove authority override refusal",
		"TestH16CeremonyRetryReplyCrashRotatesAgainWithoutCredentialReuse":   "crashes real retry subprocesses across reply delivery cuts",
		"TestH16StartupFoldsRealizedMintEvidenceBeforeServing":               "starts and restarts real conductors across startup evidence cuts",
	}
	if len(s12OuterRunOnlyFixtures) != len(want) {
		t.Fatalf("outer-run-only fixtures=%d, want %d: %#v", len(s12OuterRunOnlyFixtures), len(want), s12OuterRunOnlyFixtures)
	}
	for name, wantWhy := range want {
		skip, why := s12DogfoodOuterOnlyDecision(name, true)
		if !skip || why != wantWhy {
			t.Errorf("nested decision for %s = (%v, %q), want (true, %q)", name, skip, why, wantWhy)
		}
		if skip, _ := s12DogfoodOuterOnlyDecision(name, false); skip {
			t.Errorf("outer decision skipped %s", name)
		}
	}
	if skip, why := s12DogfoodOuterOnlyDecision("TestUnlisted", true); skip || why != "" {
		t.Fatalf("unlisted nested decision = (%v, %q), want (false, empty)", skip, why)
	}
	const wantAttestation = "nested-green attests the courier's gate adjudication over the nested suite; the process-level crash/ceremony cuts are proven ONLY by the outer run."
	if s12NestedAttestation != wantAttestation {
		t.Fatalf("nested attestation = %q, want %q", s12NestedAttestation, wantAttestation)
	}
	if s12OuterReachabilityArgument == "" {
		t.Fatal("outer reachability argument is empty")
	}
}
