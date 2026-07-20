## PLAN-REVIEW - s10 comms-spine PLAN r2 approved; lineage-only blocker folded

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s10-plan-m6-review-r2
PARENT_DISPATCH_ID: s10-plan-m6
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - pair plan-review; implementation token remains delegated to m-6.planner and merge remains operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
DESIGN_RECORD_KIND: design-doc
DESIGN_LOCK_ID: c3-design-m-6-human-surface-scheduler
PLAN_LOCK_ID: s10-comms-spine-plan
IN_REPLY_TO: s10-plan-m6/PLAN-planner-20260712-193605.md
FROM: m-6.implementer
TO: m-6.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-3.planner, m-7.planner, m-7.implementer, m-1.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: approve - r2 folds the design-doc parent edge by pointing to c6-fix-m-6-review-r2; technical plan body remains accepted and unchanged

PLAN_REVIEW_VERDICT: approve

I reviewed `s10-plan-m6/PLAN-planner-20260712-193605.md` against my r1 `must-revise-lineage-only` relay, the unchanged plan doc `master/domains/m-6-human-surface-scheduler/plan/2026-07-12-s10-comms-spine-plan.md`, the s10 dispatch, Q6xQ4 reconcile, and live `frank@8941889`.

The r2 PLAN is approved. This approval closes the pair PLAN-REVIEW gate only. It grants no implementation authority by itself and no merge authority.

## Finding Closed

F1 - CLOSED: the design-doc PLAN parent edge now resolves to the latest same-owner approving design-review edge.

- r1 blocker: `PARENT_DISPATCH_ID: s10-dispatch` could not satisfy the design-review lineage gate for a `DESIGN_RECORD_KIND: design-doc` PLAN.
- r2 fix: `PARENT_DISPATCH_ID: c6-fix-m-6-review-r2`.
- Confirmed parent relay: `master/relays/c6-fix-m-6/DESIGN-REVIEW-implementer-20260702-211652.md` carries `DISPATCH_ID: c6-fix-m-6-review-r2`, `DESIGN_REVIEW_VERDICT: approve`, `FROM: m-6.implementer`, and `DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler`.
- r2 keeps `IN_REPLY_TO: s10-dispatch/PLAN-orchestrator-planner-20260712-205011.md`, so the orchestrator work request remains visible while the structural parent edge is lint-valid.

## Technical Review Carried Forward

The plan-doc contract remains technically approved:

- Scope matches the s10 dispatch: minimum Bucket-A spine only, with ODB render/capture, park, validated operator reply, local re-observe on wake, exactly-once wake, deterministic resummon, and EXIT LEG 3.
- Out-of-scope stays preserved: B/C/D projections, elaborate-more, away-token machinery, live egress activation, TUI/email-client UX, and the full 8a freeze/reissue branch remain s11 or later.
- SEQ-2 and PARK-ACROSS-V8 remain coherent: fresh v8 genesis for dogfood, T1 transition before any park fixture, no in-slice parked-v7 gate crossing the v8 bump.
- Cross-domain ownership remains explicit: m-2 owns the one governed v7->v8 transition, m-7 owns the capability marker and scheduler/A-2 dedupe fidelity, m-3 owns re-observe and sunset-hook fidelity, and master owns any beyond-T1 schema or locked-contract/cross-domain shape change.
- Acceptance criteria remain concrete enough for red-first implementation: old-reader refusal, byte-exact terminal enum, operator-FROM validation, already-resolved and invalid-choice negatives, stale-done rejection, exactly-once crash legs, content-hash resummon dedupe, both s10 sunsets gone, and I-PH clean.

## Token Conditions

m-6.planner may issue the delegated implementation token only from a new addressed relay that:

1. parents to `s10-plan-m6-review-r2`;
2. preserves delegation conditions (a)-(g), especially no schema beyond T1 without master and no merge without operator grant;
3. carries the required mechanical `SCOPE_DIFF` before the token;
4. treats the superseded r1 PLAN as non-token-bearing and uses r2 as the reviewed plan of record.

Any scope deviation, locked-contract change, registry/schema change beyond T1, unresolved owner-confirm, or failed root-targeted lint on the new dispatch relay routes back to master before implementation.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s10-plan-m6/PLAN-planner-20260712-193605.md` - OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays master/relays/s10-plan-m6/PLAN-planner-20260712-193605.md` - target r2 PLAN reports OK; broad-root output still includes known `INDEX.md` noise, the superseded r1 PLAN lineage error, and historical merge-lineage noise.
- `diff -u master/relays/s10-plan-m6/PLAN-planner-20260712-185404.md master/relays/s10-plan-m6/PLAN-planner-20260712-193605.md` - confirms the substantive change is the lineage reissue: title/prose updates plus `PARENT_DISPATCH_ID: c6-fix-m-6-review-r2`; the plan doc was not edited.
- Read `master/domains/m-6-human-surface-scheduler/plan/2026-07-12-s10-comms-spine-plan.md` - rechecked global constraints, T1-T11, boundary contract, out-of-scope, acceptance criteria, and self-review.
- Read `master/relays/s10-plan-m6/PLAN-REVIEW-implementer-20260712-190723.md` - confirmed r1 requested only the lineage correction and expected approve if the technical body remained unchanged.
- `git -C frank status --short && git -C frank rev-parse --short HEAD` - clean output, then `8941889`.
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`.
- Post-write exact relay lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s10-plan-m6/PLAN-REVIEW-implementer-20260712-193841.md` - OK.

ACTIONS_GIT_REF: wrote this PLAN-REVIEW relay and appended one `master/relays/INDEX.md` row; no `frank/` source/test edit, no plan-doc edit, no branch/worktree, no commit, no implementation dispatch, no merge.
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `8941889`; cwd root is a docs workspace and not a git repo.
Next requested action: m-6.planner may issue the delegated implementation-token relay parented to `s10-plan-m6-review-r2`, with the required mechanical scope diff and all s10 dispatch conditions preserved.
