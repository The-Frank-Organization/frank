package fieldspec_test

import (
	"slices"
	"testing"

	"github.com/jackli/frank/internal/fieldspec"
)

func TestRenderV2GrantAndAuthoritySeatScope(t *testing.T) {
	reg := loadRegistry(t)
	env := fieldspec.RenderEnv{ConfigDigest: "cfg-1"}

	pairForm, _ := reg.Render(env, fieldspec.SeatMeta{Name: "s1-core.implementer", Role: "implementer"}, "MERGE-GATE", "medium", fieldspec.ClosedGrantState)
	if pairForm.HasField("grant") {
		t.Fatalf("pair form unexpectedly renders grant")
	}
	if pairForm.OptionAllowed("AUTHORITY", "merge-gated") {
		t.Fatalf("pair form unexpectedly allows merge-gated")
	}

	reviewerForm, _ := reg.Render(env, fieldspec.SeatMeta{Name: "s3.orchestrator-reviewer", Role: "orchestrator-reviewer"}, "MERGE-GATE", "medium", fieldspec.ClosedGrantState)
	if reviewerForm.HasField("grant") {
		t.Fatalf("reviewer form unexpectedly renders grant")
	}
	if reviewerForm.OptionAllowed("AUTHORITY", "merge-gated") {
		t.Fatalf("reviewer form unexpectedly allows merge-gated")
	}

	operatorForm, _ := reg.Render(env, fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}, "MERGE-GATE", "medium", fieldspec.ClosedGrantState)
	if !operatorForm.OptionAllowed("grant", "dispatch-impl") || !operatorForm.OptionAllowed("grant", "dispatch-merge") {
		t.Fatalf("operator merge-gate grant options = %#v", operatorForm.Fields["grant"].Options)
	}

	orchestratorForm, _ := reg.Render(env, fieldspec.SeatMeta{Name: "s3.orchestrator-planner", Role: "orchestrator-planner"}, "IMPL", "medium", fieldspec.ClosedGrantState)
	if !orchestratorForm.OptionAllowed("grant", "dispatch-impl") {
		t.Fatalf("orchestrator form missing dispatch-impl: %#v", orchestratorForm.Fields["grant"].Options)
	}
	if orchestratorForm.OptionAllowed("grant", "dispatch-merge") {
		t.Fatalf("orchestrator IMPL form unexpectedly allows dispatch-merge")
	}
}

func TestRenderV2MonotonicFloorAndDigestContext(t *testing.T) {
	reg := loadRegistry(t)
	env := fieldspec.RenderEnv{ConfigDigest: "cfg-1"}
	seat := fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}

	form, digest := reg.Render(env, seat, "IMPL", "medium", fieldspec.ClosedGrantState)
	if got := form.Fields["HUMAN_GATE_REQUIRED"].Options; !slices.Equal(got, []string{"no", "yes"}) {
		t.Fatalf("HUMAN_GATE_REQUIRED options = %v", got)
	}

	_, byConfig := reg.Render(fieldspec.RenderEnv{ConfigDigest: "cfg-2"}, seat, "IMPL", "medium", fieldspec.ClosedGrantState)
	assertDigestChanged(t, digest, byConfig, "config digest")
	_, bySeat := reg.Render(env, fieldspec.SeatMeta{Name: "s3.orchestrator-planner", Role: "orchestrator-planner"}, "IMPL", "medium", fieldspec.ClosedGrantState)
	assertDigestChanged(t, digest, bySeat, "seat")
	_, byPhase := reg.Render(env, seat, "PLAN", "medium", fieldspec.ClosedGrantState)
	assertDigestChanged(t, digest, byPhase, "phase")
	_, byTier := reg.Render(env, seat, "IMPL", "large", fieldspec.ClosedGrantState)
	assertDigestChanged(t, digest, byTier, "tier")

	changed := loadRegistry(t)
	changed.NamedEnums["PHASE"] = append(append([]string(nil), changed.NamedEnums["PHASE"]...), "EXTRA")
	changed.Phase = append([]string(nil), changed.NamedEnums["PHASE"]...)
	_, byRegistry := changed.Render(env, seat, "IMPL", "medium", fieldspec.ClosedGrantState)
	assertDigestChanged(t, digest, byRegistry, "registry options")
}

func TestRenderV2ParentCandidatesDigestExempt(t *testing.T) {
	reg := loadRegistry(t)
	seat := fieldspec.SeatMeta{Name: "s1-core.implementer", Role: "implementer"}

	envA := fieldspec.RenderEnv{
		ConfigDigest: "cfg-1",
		ParentCandidates: func(fieldspec.SeatMeta) ([]string, string) {
			return []string{"parent-a"}, "parent-a"
		},
	}
	envB := fieldspec.RenderEnv{
		ConfigDigest: "cfg-1",
		ParentCandidates: func(fieldspec.SeatMeta) ([]string, string) {
			return []string{"parent-b", "parent-c"}, "parent-b"
		},
	}

	formA, digestA := reg.Render(envA, seat, "IMPL", "medium", fieldspec.ClosedGrantState)
	formB, digestB := reg.Render(envB, seat, "IMPL", "medium", fieldspec.ClosedGrantState)
	if digestA != digestB {
		t.Fatalf("candidate-set contents changed digest: %s vs %s", digestA, digestB)
	}
	parentA := formA.Fields["PARENT_DISPATCH_ID"]
	parentB := formB.Fields["PARENT_DISPATCH_ID"]
	if !parentA.DigestExempt || !parentB.DigestExempt {
		t.Fatalf("parent picker not marked digest-exempt: %#v %#v", parentA, parentB)
	}
	if slices.Equal(parentA.Options, parentB.Options) || parentA.Default == parentB.Default {
		t.Fatalf("parent candidate contents did not render distinctly: %#v %#v", parentA, parentB)
	}
}

func assertDigestChanged(t *testing.T, before, after, dimension string) {
	t.Helper()
	if before == after {
		t.Fatalf("digest unchanged after %s changed", dimension)
	}
}
