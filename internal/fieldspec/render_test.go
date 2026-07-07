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

	pairPlanner := fieldspec.SeatMeta{Name: "s3-form.planner", Role: "planner"}
	closedPairPlanner, _ := reg.Render(env, pairPlanner, "IMPL", "medium", fieldspec.ClosedGrantState)
	if closedPairPlanner.HasField("grant") {
		t.Fatalf("pair planner grant rendered before approved plan-review chain")
	}
	openPairPlanner, _ := reg.Render(env, pairPlanner, "IMPL", "medium", func(fieldspec.SeatMeta) bool { return true })
	if !openPairPlanner.OptionAllowed("grant", "dispatch-impl") {
		t.Fatalf("pair planner grant did not render after grant state opened: %#v", openPairPlanner.Fields["grant"].Options)
	}
	if openPairPlanner.OptionAllowed("grant", "dispatch-merge") {
		t.Fatalf("pair planner IMPL form unexpectedly allows dispatch-merge")
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
	raised, _ := reg.Render(fieldspec.RenderEnv{ConfigDigest: "cfg-1", MonotonicFloors: map[string]string{"HUMAN_GATE_REQUIRED": "yes", "GRILL_REQUIRED": "yes"}}, seat, "IMPL", "medium", fieldspec.ClosedGrantState)
	if got := raised.Fields["HUMAN_GATE_REQUIRED"].Options; !slices.Equal(got, []string{"yes"}) {
		t.Fatalf("raised HUMAN_GATE_REQUIRED options = %v, want [yes]", got)
	}
	if got := raised.Fields["GRILL_REQUIRED"].Options; !slices.Equal(got, []string{"yes"}) {
		t.Fatalf("raised GRILL_REQUIRED options = %v, want [yes]", got)
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

func TestRenderStableDigestIgnoresConductorVolatileClasses(t *testing.T) {
	reg := loadRegistry(t)
	seat := fieldspec.SeatMeta{Name: "s3-form.planner", Role: "planner"}
	baseEnv := fieldspec.RenderEnv{
		ConfigDigest: "cfg-1",
		ParentCandidates: func(fieldspec.SeatMeta) ([]string, string) {
			return []string{"parent-a"}, "parent-a"
		},
		RecipientCandidates: func(fieldspec.SeatMeta) []string {
			return []string{"seat-a", "seat-b"}
		},
	}
	volatileEnv := fieldspec.RenderEnv{
		ConfigDigest: "cfg-1",
		ParentCandidates: func(fieldspec.SeatMeta) ([]string, string) {
			return []string{"parent-c", "parent-d"}, "parent-c"
		},
		RecipientCandidates: func(fieldspec.SeatMeta) []string {
			return []string{"seat-c"}
		},
		MonotonicFloors: map[string]string{"HUMAN_GATE_REQUIRED": "yes", "GRILL_REQUIRED": "yes"},
	}

	closedForm, closedDigest := reg.Render(baseEnv, seat, "IMPL", "medium", fieldspec.ClosedGrantState)
	openForm, openDigest := reg.Render(volatileEnv, seat, "IMPL", "medium", func(fieldspec.SeatMeta) bool { return true })
	if closedDigest != openDigest {
		t.Fatalf("volatile options changed digest: %s vs %s", closedDigest, openDigest)
	}
	if closedForm.HasField("grant") {
		t.Fatalf("grant rendered before grant state opened")
	}
	grant := openForm.Fields["grant"]
	if !grant.ConductorVolatile || !grant.DigestExempt || len(grant.Options) == 0 {
		t.Fatalf("open grant field = %#v, want volatile options in rendered form", grant)
	}
	for _, id := range []string{"TO", "CC", "HUMAN_GATE_REQUIRED", "GRILL_REQUIRED"} {
		field, ok := openForm.Fields[id]
		if !ok {
			t.Fatalf("open form missing %s", id)
		}
		if !field.ConductorVolatile || !field.DigestExempt {
			t.Fatalf("%s field = %#v, want conductor volatile digest-exempt", id, field)
		}
	}
	if slices.Equal(closedForm.Fields["TO"].Options, openForm.Fields["TO"].Options) ||
		slices.Equal(closedForm.Fields["HUMAN_GATE_REQUIRED"].Options, openForm.Fields["HUMAN_GATE_REQUIRED"].Options) {
		t.Fatalf("volatile rendered options did not differ: closed=%#v open=%#v", closedForm.Fields, openForm.Fields)
	}

	_, byConfig := reg.Render(fieldspec.RenderEnv{ConfigDigest: "cfg-2"}, seat, "IMPL", "medium", fieldspec.ClosedGrantState)
	assertDigestChanged(t, closedDigest, byConfig, "config digest")
	changed := loadRegistry(t)
	changed.NamedEnums["PHASE"] = append(append([]string(nil), changed.NamedEnums["PHASE"]...), "EXTRA")
	changed.Phase = append([]string(nil), changed.NamedEnums["PHASE"]...)
	_, byStaticEnum := changed.Render(baseEnv, seat, "IMPL", "medium", fieldspec.ClosedGrantState)
	assertDigestChanged(t, closedDigest, byStaticEnum, "static enum shape")
}

func TestRenderParentDispatchIDAbsentFromEveryForm(t *testing.T) {
	reg := loadRegistry(t)
	seats := []fieldspec.SeatMeta{
		{Name: "operator", Role: "operator", IsOperator: true},
		{Name: "s6-core.planner", Role: "planner"},
		{Name: "s6-core.implementer", Role: "implementer"},
		{Name: "s6.orchestrator-planner", Role: "orchestrator-planner"},
	}
	for _, seat := range seats {
		for _, phase := range reg.NamedEnums["PHASE"] {
			form, _ := reg.Render(fieldspec.RenderEnv{}, seat, phase, "medium", func(fieldspec.SeatMeta) bool { return true })
			if form.HasField("PARENT_DISPATCH_ID") {
				t.Fatalf("%s/%s rendered PARENT_DISPATCH_ID", seat.Name, phase)
			}
		}
	}
}

func assertDigestChanged(t *testing.T, before, after, dimension string) {
	t.Helper()
	if before == after {
		t.Fatalf("digest unchanged after %s changed", dimension)
	}
}
