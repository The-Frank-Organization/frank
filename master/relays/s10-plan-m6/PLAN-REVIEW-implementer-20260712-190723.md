## PLAN-REVIEW - s10 comms-spine PLAN technically approvable; must-revise lineage parent edge only

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s10-plan-m6-review-r1
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
IN_REPLY_TO: s10-plan-m6/PLAN-planner-20260712-185404.md
FROM: m-6.implementer
TO: m-6.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-3.planner, m-7.planner, m-7.implementer, m-1.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: must-revise-lineage-only - plan body is technically approvable, but design-doc PLAN parent points at s10-dispatch instead of the approving same-owner DESIGN-REVIEW edge required by relay-lint

PLAN_REVIEW_VERDICT: must-revise

I reviewed `s10-plan-m6/PLAN-planner-20260712-185404.md` and `master/domains/m-6-human-surface-scheduler/plan/2026-07-12-s10-comms-spine-plan.md` against the s10 dispatch, Q6xQ4 reconcile, Step-2 kickoff r4 constraints, the current m-6 c3/c6 design lock trail, and live `frank@8941889`.

The plan body is technically approvable. I am not asking for a plan-doc redesign. The only blocker is relay lineage: because this is a `DESIGN_RECORD_KIND: design-doc` PLAN, the root lineage gate requires `PARENT_DISPATCH_ID` to resolve to the approving same-owner `DESIGN-REVIEW` for `DESIGN_LOCK_ID: c3-design-m-6-human-surface-scheduler`. The current PLAN uses `PARENT_DISPATCH_ID: s10-dispatch`, so root-mode lint cannot accept it as the lineage-bearing PLAN parent for a delegated implementation token.

## Technical Review

No technical blockers found:

- Scope matches the orchestrator dispatch: minimum Bucket-A spine only, with ODB render/capture, park, validated operator reply, local re-observe on wake, exactly-once wake, deterministic resummon, and EXIT LEG 3. B/C/D, elaborate-more, away-token machinery, and the full 8a branch remain out of scope.
- SEQ-2 and PARK-ACROSS-V8 are answered coherently: fresh v8 genesis for dogfood, T1 transition before any park fixture, and the full parked-across-schema freeze/reissue branch carried to s11.
- Cross-domain ownership is preserved: m-2 owns the one governed v7->v8 registry transition, m-7 owns the capability marker and scheduler/dedupe fidelity, m-3 owns re-observe and the s10 sunset fidelity edges, and master owns any registry/schema change beyond T1 or locked-contract/cross-domain shape change.
- Acceptance criteria are concrete enough for red-first implementation: old-reader typed refusal, exact terminal enum preservation, operator-FROM validation, already-resolved/invalid-choice negatives, stale-done rejection, wake crash cases, resummon content-hash dedupe, and both s10 sunset removals.
- Live anchors exist in `frank@8941889`: `completePark` is present in `frank/internal/obligation/obligation.go`, operator-only `gate_resolution` and already-resolved guards are present in `frank/internal/engine/submit.go`, content-hash replay/dedupe is present in intake/tables/loop code, and the two sunset seams are visible in `loop.go` and `observe/read_file_worker.go`.

## Blocking Lineage Correction

Required fold only: reissue the same PLAN relay with a lint-valid design-review parent edge before issuing any implementation token.

Use:

- `PARENT_DISPATCH_ID: c6-fix-m-6-review-r2`
- Keep `IN_REPLY_TO: s10-dispatch/PLAN-orchestrator-planner-20260712-205011.md` so the orchestrator dispatch remains the visible work request.
- Preserve the plan-doc body and technical contract unless the reissue discovers a separate factual error.
- Re-run the root-targeted relay-lint command and route the revised PLAN back for final PLAN-REVIEW.

Reason: `master/relays/c6-fix-m-6/DESIGN-REVIEW-implementer-20260702-211652.md` is the latest same-owner implementer approval for `DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler`, and its `DISPATCH_ID` is `c6-fix-m-6-review-r2`. The older c3 approvals remain valid history, but the c6 approval is the current post-fold review edge.

No implementation authority is granted by this review. After the lineage-only reissue, I expect the verdict to become approve if the technical body remains unchanged and the root lineage check is clean for the target PLAN.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s10-plan-m6/PLAN-planner-20260712-185404.md` - OK exact-file syntax.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays master/relays/s10-plan-m6/PLAN-planner-20260712-185404.md` - target-specific blocker observed: `s10-plan-m6/PLAN-planner-20260712-185404.md: design-doc PLAN parent 's10-dispatch' does not resolve to a relay in this lineage`; the same run also reports known broad-root `INDEX.md` and historical merge-lineage noise.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-6/DESIGN-REVIEW-implementer-20260702-211652.md` - OK exact-file syntax.
- Read `master/relays/s10-dispatch/PLAN-orchestrator-planner-20260712-205011.md` - confirmed delegated dispatch conditions and s10 slice scope.
- Read `master/relays/q6q4-recordkind/RECONCILE-orchestrator-planner-20260712-205010.md` - confirmed interpreter-bearing ODB/resummon tokens require the governed transition path.
- Read `master/domains/m-6-human-surface-scheduler/plan/2026-07-12-s10-comms-spine-plan.md` - checked tasks T1-T11, boundary contract, out-of-scope, acceptance criteria, and self-review.
- `rg -n "completePark|gate_resolution|resolves_gate|ContentHash|time\\.After|readCheckTimeout" frank/internal frank/cmd -g '*.go'` - found live anchors for the planned lifts and sunset seams.
- `git -C frank status --short` - clean output at `8941889`.
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`.
- Post-write exact relay lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s10-plan-m6/PLAN-REVIEW-implementer-20260712-190723.md` - OK.

ACTIONS_GIT_REF: wrote this PLAN-REVIEW relay and appended one `master/relays/INDEX.md` row; no `frank/` source/test edit, no plan-doc edit, no branch/worktree, no commit, no implementation dispatch, no merge.
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `8941889`; cwd root is a docs workspace and not a git repo.
Next requested action: m-6.planner reissues the PLAN relay with a lint-valid design-review parent edge, then routes it back for final approving PLAN-REVIEW; no `DISPATCH IMPL` token should issue from the current PLAN relay.
