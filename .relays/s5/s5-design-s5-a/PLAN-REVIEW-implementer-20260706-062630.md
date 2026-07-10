## PLAN-REVIEW - s5-a.implementer review of s5-a-registry-plan

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s5-design-s5-a
PARENT_DISPATCH_ID: s5-design-s5-a
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_LOCK_ID: s5-a-registry-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s5-a-registry-plan
PLAN_REVIEW_VERDICT: approve
FROM: s5-a.implementer
TO: s5-a.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer
IN_REPLY_TO: .relays/s5/s5-design-s5-a/PLAN-planner-20260706-061650.md
SUBJECT: PLAN-REVIEW approve - plan matches approved registry design; 38-name guard and TDD red checks hold

## Routing and authority

- This PLAN-REVIEW is directly addressed work: `PLAN-planner-20260706-061650.md` is `TO: s5-a.implementer`.
- Authority is review-only. I made no source, test, registry, sprint-doc, branch, commit, PR, merge, implementation, or dispatch edits.
- Reviewed artifacts: the planner PLAN relay, the plan of record `docs/sprints/2026-07-06-s5-slice-5/plans/s5-a-registry-pass-plan.md`, the approved rev2 design doc, the approving DESIGN-REVIEW `DESIGN-REVIEW-implementer-20260706-060559.md`, and the orchestrator PROCEED-TO-PLAN relay.

## Verdict

PLAN_REVIEW_VERDICT: approve

The plan is approved for `PLAN_LOCK_ID: s5-a-registry-plan` locked to `DESIGN_LOCK_ID: s5-a-registry-design`.

This approval is not implementation dispatch. The planner relay correctly states that delegated implementation authority can follow only after the mechanical SCOPE_DIFF is all-in, no hard trigger/collision exists, and the dispatch parents this PLAN-REVIEW.

## Review checks

- Lineage is correct: the PLAN relay locks `DESIGN_LOCK_ID: s5-a-registry-design`, carries `DESIGN_RECORD_KIND: design-doc`, and replies to the approving DESIGN-REVIEW; that review approved the same `DESIGN_DOC_ID` and explicitly did not authorize implementation (`PLAN-planner-20260706-061650.md:12-19`, `DESIGN-REVIEW-implementer-20260706-060559.md:12-27`).
- Scope fence is explicit and matches the approved design plus orchestrator F2 surface: `registry.json`, `registry_test.go`, the new `test/fixtures/s5_registry_dormancy_test.go`, and `render_test.go`/`validate_test.go` only for registry-content cases that fit there better. Production Go, `fieldspec_test.go`, engine/bounce/migrate/test-replay, sprint docs, `.relays/`, and archived store data stay OUT (`s5-a-registry-pass-plan.md:8-15`, `PLAN-orchestrator-planner-20260706-061037.md:21-29`).
- The 38-name slice is consistent with the design tables: 36 new rows (Block A 12 + Block B 13 + Block C 9 + Block D 2) plus the two edited observe action-report rows `ACTIONS_GIT_REF` and `FINAL_GIT_STATUS_SHORT` (`s5-a-registry-pass-design.md:44-83`, `:85-118`, `:122-126`).
- The current base proves the Step-1 red-state plan is meaningful: live registry is version `s3-fieldspec-v2`, 47 rows, 14 named enums, no new candidate row IDs, `gate_category_A` has 8 tokens, `gate_category` has 13 with no `routing_escalation`, and `record_kind` still renders `genesis` plus owed/gate/disposition to `*` (`internal/fieldspec/registry.json:90`, `:111-113`; Python registry readback during this review).
- The proposed red assertions should bite before the registry edit: row count/version/enum/token assertions fail in `registry_test.go`, the 38-name dormancy sweep fails because the row IDs are absent, the OI-S4 scope legs fail against the current `record_kind` seat_scope, and the EVIDENCE_TARGET total-required leg fails before the planned predicate edit (`registry_test.go:14-48`, `s5-a-registry-pass-plan.md:21-29`).
- Assertion placement is acceptable: byte-exact registry shape belongs in `registry_test.go`; render sweep, digest determinism, OI-S4 scope legs, payload controls, and raw-JSON annotation checks fit the new fixture package because `test/fixtures` already imports `internal/fieldspec`; `render_test.go` and `validate_test.go` remain available only for narrow registry-content cases (`test/fixtures/s4_config_change_test.go:1-15`, `render_test.go:10-55`, `validate_test.go:10-46`).
- The anti-half-fix guards cover both prior review rounds: the omitted `on_timeout` row, stale 82/37/35 count class, raw JSON annotation mechanics, `"yes"` not `"true"` predicate byte, m-2 in-pass confirms, and the corrected `on_timeout` boundary that leaves submit-path typed rejection to s5-b (`s5-a-registry-pass-plan.md:36-46`, `DESIGN-REVIEW-implementer-20260706-060559.md:37-51`).

## Hard-trigger review

No unresolved operator judgment is open at PLAN time. The plan has explicit stop points for any OUT file, cross-slice collision, design-contract amendment, or in-pass m-2 confirm that changes row content before the IMPL report. Those are the correct escalation points.

## Non-blocking carry into dispatch

Before issuing any delegated implementation dispatch, run the promised mechanical SCOPE_DIFF against this plan's file map and the orchestrator's dispatched surface. If `render_test.go` or `validate_test.go` is treated as in-scope, the SCOPE_DIFF should preserve the plan's limitation: registry-content assertions only, not mechanism changes.

## Actions and verification

ACTIONS_GIT_REF:
- Source/test/registry/sprint-doc edits: none.
- Relay artifact written: `.relays/s5/s5-design-s5-a/PLAN-REVIEW-implementer-20260706-062630.md`.
- Index row appended: `.relays/s5/INDEX.md`.
- Pre-write `git status --short --untracked-files=all`:
  `?? docs/sprints/2026-07-06-s5-slice-5/RECONCILE.md`
  `?? docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md`
  `?? docs/sprints/2026-07-06-s5-slice-5/designs/s5-b-mechanisms-design.md`
  `?? docs/sprints/2026-07-06-s5-slice-5/plans/s5-a-registry-pass-plan.md`
  `?? docs/sprints/2026-07-06-s5-slice-5/plans/s5-b-mechanisms-plan.md`

FINAL_GIT_STATUS_SHORT:
?? docs/sprints/2026-07-06-s5-slice-5/RECONCILE.md
?? docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md
?? docs/sprints/2026-07-06-s5-slice-5/designs/s5-b-mechanisms-design.md
?? docs/sprints/2026-07-06-s5-slice-5/plans/s5-a-registry-pass-plan.md
?? docs/sprints/2026-07-06-s5-slice-5/plans/s5-b-mechanisms-plan.md

VERIFICATION:
- Planner PLAN exact-file lint: `OK .relays/s5/s5-design-s5-a/PLAN-planner-20260706-061650.md`.
- Planner PLAN root lineage lint: `OK .relays/s5/s5-design-s5-a` and `OK .relays/s5/s5-design-s5-a/PLAN-planner-20260706-061650.md`.
- `go test ./internal/fieldspec ./test/fixtures` - OK.
- `go test ./...` - OK.
- This PLAN-REVIEW exact-file lint: `OK .relays/s5/s5-design-s5-a/PLAN-REVIEW-implementer-20260706-062630.md`.
- This PLAN-REVIEW root lineage lint: `OK .relays/s5/s5-design-s5-a` and `OK .relays/s5/s5-design-s5-a/PLAN-REVIEW-implementer-20260706-062630.md`.
- Index tail readback: row `20260706-062630 | PLAN-REVIEW | Implementer | s5-design-s5-a | ... | approve | s5-design-s5-a/PLAN-REVIEW-implementer-20260706-062630.md` present.
