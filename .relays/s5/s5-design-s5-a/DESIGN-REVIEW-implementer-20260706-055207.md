## DESIGN-REVIEW - s5-a implementer review of registry-pass design

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s5-design-s5-a
PARENT_DISPATCH_ID: s5-design-s5-a
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: s5-a-registry-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: .relays/s5/s5-design-s5-a/DESIGN-planner-20260706-053601.md
FROM: s5-a.implementer
TO: s5-a.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer, operator
SUBJECT: DESIGN-REVIEW must-revise - stale riding-leg closure, missing on_timeout row, and test-scope contradiction

## Verdict

DESIGN_REVIEW_VERDICT: must-revise

The design is close on the hard code-shape questions. I approve the direction of D-1, D-4, D-5, and D-6 after live-code checks: unknown row keys are loader-transparent, `any_row` needs a gate-referenceable carrier, `"yes"` is the right bool byte, and `{"not":{"phase_in":[]}}` parses and evaluates as a total required_when predicate even for empty PHASE headers. Those points are not the blockers.

The blockers are scope fidelity issues in the design artifact. They would make the PLAN lock stale or internally contradictory if approved as written.

## Findings

1. must-revise - the design missed the later master riding-leg closure and now carries stale scope variants.

The design cites `.relays/s5/s5-escalations/RECONCILE-orchestrator-planner-20260706-052214.md` as the escalation authority and leaves the m-1/m-4/m-6 riding legs open or integration-pending. A newer authority already existed before this design request: `.relays/s5/s5-escalations/RECONCILE-orchestrator-planner-20260706-053113.md`. It closes all three riding legs, says there are zero open m-x legs, makes owed rows operator-only, removes `genesis` from every seat scope including operator, and adds adjacent scope checks for `gate_resolution` and `disposition`.

That contradicts the design's §6/§9 shape: it keeps the operator `record_kind` scope unchanged with all seven tokens, leaves fallback variants for owed/genesis, and says the riding legs do not block the review. The revised design must fold the 053113 reconcile directly: no pending m-1/m-4/m-6 legs, owed rows operator-only, `genesis` in no rendered scope, and explicit handling of the `gate_resolution` / `disposition` adjacent flag.

Evidence:
- design request: `.relays/s5/s5-design-s5-a/DESIGN-planner-20260706-053601.md:20-29`.
- design doc stale variants: `docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md:110` and `:135-143`.
- latest authority: `.relays/s5/s5-escalations/RECONCILE-orchestrator-planner-20260706-053113.md:19-32`.

2. must-revise - the confirmed `on_timeout` row is omitted from the row tables and fixture count.

The 052753 update and the 052214 reconcile both require M-3(d): `on_timeout` as a valueless reserved row, with the m-6 policy floor annotation that no value may ever mean auto-approve or auto-resolve. The design's Block B table includes `decision_deadline` but no `on_timeout`, and `rg on_timeout` over the design doc returns no design-row hit.

This is not just a naming detail. It invalidates the claimed row count and the [VP-W3] 37-name render absence fixture: if `on_timeout` is in scope, it must be named in the 35 new rows and included in the enumerated absence list, or the design must explicitly route it out with authority.

Evidence:
- required by update: `.relays/s5/s5-design-s5-a/DESIGN-orchestrator-planner-20260706-052753.md:21-32`.
- required by master reconcile: `.relays/s5/s5-escalations/RECONCILE-orchestrator-planner-20260706-052214.md:39-41`.
- omitted from Block B: `docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md:64-81`.
- count/fixture claims that must change if the row is added: `docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md:15` and `:118`.

3. must-revise - acceptance criterion 6 contradicts the authorized test surface.

R-s5-2 explicitly puts `registry_test.go`, `render_test.go`, and `validate_test.go` registry-content fixtures in s5-a's write surface. The design repeats those files in §7, but its acceptance criterion 6 says no `.go` file under `internal/` or `cmd/` may be modified by s5-a. That would forbid the very internal fieldspec test edits the design and dispatch require.

The fix is small: change the boundary to "no production `.go` under `internal/` or `cmd/`; `internal/fieldspec/*_test.go` is in-scope for registry-content tests." If the intent is to avoid all internal test edits, then §7 must move those cases elsewhere and explain how the registry-content assertions remain covered.

Evidence:
- R-s5-2 authorized test surface: `.relays/s5/s5-reconcile-audits/RECONCILE-orchestrator-planner-20260706-045327.md:37-39`.
- design §7 requires those files: `docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md:114-127`.
- conflicting acceptance criterion: `docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md:145-153`.

## Non-blocking checks

- D-1 annotation key: acceptable as a bytes-only annotation path. `json.Unmarshal` into `FieldSpec` ignores unknown keys, so `Load` does not choke; because `FieldSpec` has no `Annotation` field, tests for annotation presence must inspect raw JSON rather than loaded structs. Evidence: `internal/fieldspec/registry.go:24-41`, `:64-81`.
- D-4/D-5 routing assignment predicate: acceptable for the degraded shape. `any_row` parses only when the array field is gate-referenceable, and the live row-array parser returns `[]map[string]string`; `"yes"` is the correct bool byte. Evidence: `internal/fieldspec/predicate.go:121-149`, `internal/fieldspec/canonical.go:19-24`, `internal/fieldspec/canonical.go:34-38`.
- D-6 total predicate: acceptable. `phase_in` accepts an empty string array; `not(phase_in [])` evaluates true because no PHASE value can be contained in an empty list, including an omitted/empty PHASE header. Evidence: `internal/fieldspec/predicate.go:112-120`, `:166-173`, `:304-309`, `:337-344`.
- Observe visibility trap: the design's ACTIONS/FINAL fix is correct. `Render` uses `DefaultLayers()` without observe and only `visible_when` hides otherwise renderable fields, so required_when alone is insufficient. Evidence: `internal/fieldspec/render.go:30-57` and current rows at `internal/fieldspec/registry.json:111-112`.

## Required revision checklist

- Fold `.relays/s5/s5-escalations/RECONCILE-orchestrator-planner-20260706-053113.md` into the design basis and remove stale riding-leg language.
- Update `record_kind` scope design: owed rows operator-only; `genesis` removed from all seat scopes; address `gate_resolution` and `disposition` per the adjacent flag.
- Add or explicitly re-route `on_timeout`; if added, update row counts, [VP-W3] enumerated names, and fixture claims.
- Resolve the `.go` test-surface contradiction by distinguishing production `.go` from authorized registry-content test files.

## Verification

- Reviewed the addressed DESIGN relay, design doc, 045327/052753 s5 design dispatches, s5-reconcile-audits, s5-fidelity, and both s5-escalations reconciles.
- Reviewed live fieldspec code for the D-1/D-4/D-5/D-6 pressure points: registry loader, predicate parser/evaluator, renderer, validator, canonical row-array parser, and current registry rows.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s5/s5-design-s5-a/DESIGN-REVIEW-implementer-20260706-055207.md` - OK.
- `go test ./internal/fieldspec` - OK.
- Pre-write `git rev-parse --short HEAD`: `67ee23e`.
- Pre-write `git status --short --untracked-files=all`:
  - `?? docs/sprints/2026-07-06-s5-slice-5/RECONCILE.md`
  - `?? docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md`
  - `?? docs/sprints/2026-07-06-s5-slice-5/designs/s5-b-mechanisms-design.md`

ACTIONS_GIT_REF: wrote `.relays/s5/s5-design-s5-a/DESIGN-REVIEW-implementer-20260706-055207.md` and appended `.relays/s5/INDEX.md`; no design-doc, registry, source, test, branch, commit, PR, merge, or store edits made by this review.
FINAL_GIT_STATUS_SHORT:
?? docs/sprints/2026-07-06-s5-slice-5/RECONCILE.md
?? docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md
?? docs/sprints/2026-07-06-s5-slice-5/designs/s5-b-mechanisms-design.md

Next requested action: s5-a.planner revises the design doc and re-requests DESIGN-REVIEW against the same `DESIGN_DOC_ID` or an explicitly superseding one.
