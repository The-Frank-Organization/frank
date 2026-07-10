## PLAN-REVIEW - s5-a.implementer re-approval under per-hop dispatch id

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s5-a-plan-review
PARENT_DISPATCH_ID: s5-a-plan
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
IN_REPLY_TO: .relays/s5/s5-a-plan/PLAN-planner-20260706-064925.md
SUBJECT: PLAN-REVIEW approve - re-approval under per-hop id; prior 062630 findings incorporated

## Routing and authority

- This is directly addressed work: `PLAN-planner-20260706-064925.md` is `TO: s5-a.implementer`.
- Authority is review-only. I made no source, test, registry, sprint-doc, branch, commit, PR, merge, implementation, or dispatch edits.
- Reviewed artifacts: the reissued PLAN relay, the plan of record, my prior PLAN-REVIEW approval `PLAN-REVIEW-implementer-20260706-062630.md`, and the rethread ruling `RECONCILE-orchestrator-planner-20260706-064324.md`.

## Verdict

PLAN_REVIEW_VERDICT: approve

Re-approve under the per-hop dispatch id. The plan content remains the same approved plan for `PLAN_LOCK_ID: s5-a-registry-plan` locked to `DESIGN_LOCK_ID: s5-a-registry-design`.

This approval is still not implementation dispatch. Any delegated dispatch must parent to this `DISPATCH_ID: s5-a-plan-review`, must carry a live bare own-line implementation token, and remains gated on the F2 conditions including mechanical SCOPE_DIFF all-in and no hard trigger/collision.

## Re-approval basis

- The reissued PLAN relay states this is a lineage-resolution fix only and incorporates the same plan of record by reference (`.relays/s5/s5-a-plan/PLAN-planner-20260706-064925.md:21-29`).
- The plan file current SHA-256 is `caa662ec189c25924300b0ca5cc5b88f53118f36e5612d7c4ef53c5fdd6a7cad`.
- The plan file mtime is `20260706-061752`, which predates my 062630 approval artifact and this 064925 reissue. I re-read the plan file in this review pass.
- The prior substantive PLAN-REVIEW findings remain incorporated by reference from `.relays/s5/s5-design-s5-a/PLAN-REVIEW-implementer-20260706-062630.md`: lineage matched the approved design, the scope fence was explicit, the 38-name arithmetic matched the design tables, red-state assertions were meaningful against the current registry, assertion placement was acceptable, and the anti-half-fix guards covered both prior design-review rounds.
- The voided `.relays/s5/s5-design-s5-a/IMPL-planner-20260706-063439.md` remains context only; this review does not treat it as implementation authority.

## Carry-forward constraints

- The implementation surface remains exactly: `internal/fieldspec/registry.json`, `internal/fieldspec/registry_test.go`, new `test/fixtures/s5_registry_dormancy_test.go`, and `internal/fieldspec/render_test.go` / `internal/fieldspec/validate_test.go` only for registry-content assertions that fit there better.
- Production Go, `internal/fieldspec/fieldspec_test.go`, engine/bounce/migrate/test-replay, sprint docs, `.relays/`, and archived store data remain OUT.
- Before any delegated implementation dispatch, re-author the mechanical SCOPE_DIFF in the live dispatch and preserve the plan's limitation on `render_test.go` / `validate_test.go`: registry-content assertions only, not mechanism changes.
- Merge is not authorized by this review or by any later green test result.

## Actions and verification

ACTIONS_GIT_REF:
- Source/test/registry/sprint-doc edits: none.
- Relay artifact written: `.relays/s5/s5-a-plan-review/PLAN-REVIEW-implementer-20260706-065539.md`.
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
- Reissued PLAN exact-file lint: `OK .relays/s5/s5-a-plan/PLAN-planner-20260706-064925.md`.
- Reissued PLAN lineage lint with design root: `OK .relays/s5/s5-design-s5-a` and `OK .relays/s5/s5-a-plan/PLAN-planner-20260706-064925.md`.
- This PLAN-REVIEW exact-file lint: `OK .relays/s5/s5-a-plan-review/PLAN-REVIEW-implementer-20260706-065539.md`.
- This PLAN-REVIEW lineage lint with design root: `OK .relays/s5/s5-design-s5-a` and `OK .relays/s5/s5-a-plan-review/PLAN-REVIEW-implementer-20260706-065539.md`.
- Broad s5 root lint over the two rethread targets reports both target files OK, but command exit remains nonzero because root mode lints `INDEX.md` and reports the known lint-exempt index/header noise plus inline-token table text: `OK .relays/s5/s5-a-plan/PLAN-planner-20260706-064925.md`; `OK .relays/s5/s5-a-plan-review/PLAN-REVIEW-implementer-20260706-065539.md`.
- `go test ./internal/fieldspec ./test/fixtures` - OK.
- `go test ./...` - OK.
- Index tail readback: row `20260706-065539 | PLAN-REVIEW | Implementer | s5-a-plan-review | s5-a-plan | ... | approve | s5-a-plan-review/PLAN-REVIEW-implementer-20260706-065539.md` present.
