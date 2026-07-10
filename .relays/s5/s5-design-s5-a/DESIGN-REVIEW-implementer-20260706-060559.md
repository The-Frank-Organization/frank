## DESIGN-REVIEW - s5-a implementer approval of registry-pass design rev2

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
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: .relays/s5/s5-design-s5-a/DESIGN-planner-20260706-060325.md
FROM: s5-a.implementer
TO: s5-a.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer, operator
SUBJECT: DESIGN-REVIEW approve - rev2 folds on_timeout boundary and row-count correction

## Verdict

DESIGN_REVIEW_VERDICT: approve

The rev2 design is approved for PLAN. The two blockers from `DESIGN-REVIEW-implementer-20260706-060104.md` are folded: the `on_timeout` row no longer claims current enum validation on lane-supplied system/system_only payloads, and §7 now instructs `registry_test.go` to assert row count 83. The earlier must-revise blockers are also folded: 053113 is the first authority, stale riding-leg variants are gone, `on_timeout` is included in the row set, and the internal fieldspec test surface is correctly in scope.

This approval is design-review approval only. It does not authorize implementation, dispatch, merge, or any code/registry edits.

## Scope Proof

- DESIGN_DOC_ID matches the request and design doc: `s5-a-registry-design`.
- Reviewed artifact: `docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md`.
- Parent relay reviewed: `.relays/s5/s5-design-s5-a/DESIGN-planner-20260706-060325.md`.
- Prior review folded: `.relays/s5/s5-design-s5-a/DESIGN-REVIEW-implementer-20260706-060104.md`.
- Latest master scope folded: `.relays/s5/s5-escalations/RECONCILE-orchestrator-planner-20260706-053113.md`.

## Checks

- Prior blocker 1 closed: `on_timeout` now states the accurate boundary: loadable valueless row, Step-1 protection by render absence, lane-supplied rejection dependent on s5-b DEF-2, and conductor-internal conformance owned by the later scheduler writer. Evidence: design doc line 80; live validator ordering at `internal/fieldspec/validate.go:31-53` and skip at `:115-120`.
- Prior blocker 2 closed: §7 now says row count 83 in the `registry_test.go` instruction. Evidence: design doc line 124.
- Stale-count scan reviewed: remaining 82/37/35 hits are historical fold-log deltas, not live acceptance criteria. Evidence: design doc lines 175 and 182.
- Prior rev0 blockers remain folded: 053113 is first in the basis, record_kind scope is final with `genesis` in no scope and narrowed operator-only tokens, `on_timeout` is present, and criterion 6 now distinguishes production `.go` from authorized `internal/fieldspec/*_test.go`.

## Remaining Integration Gates

These are not design-review blockers, but they must remain explicit in PLAN/integration:

- m-2 in-pass shaping for MR-1 final name/type of `gate_category_pick`.
- m-2 in-pass confirmation of D-1 annotation-key mechanics and D-6 total predicate.
- m-2 in-pass ruling for the `disposition` scope narrowing.
- s5-b remains owner of DEF-2 typed-REJECT submit-path enforcement and all engine-side mechanics.

## Verification

- Reviewed the addressed rev2 DESIGN relay and current design doc.
- Re-read the 060104 must-revise relay.
- Re-checked the current design text for `on_timeout`, stale row counts, and acceptance criteria.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s5/s5-design-s5-a/DESIGN-REVIEW-implementer-20260706-060559.md` - OK.
- `go test ./internal/fieldspec` - OK.
- Index row present in `.relays/s5/INDEX.md`.

ACTIONS_GIT_REF: wrote `.relays/s5/s5-design-s5-a/DESIGN-REVIEW-implementer-20260706-060559.md` and appended `.relays/s5/INDEX.md`; no design-doc, registry, source, test, branch, commit, PR, merge, or store edits made by this review.
FINAL_GIT_STATUS_SHORT:
?? docs/sprints/2026-07-06-s5-slice-5/RECONCILE.md
?? docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md
?? docs/sprints/2026-07-06-s5-slice-5/designs/s5-b-mechanisms-design.md

Next requested action: s5-a.planner may report design-complete to s5.orchestrator-planner and hold for PROCEED-TO-PLAN.
