## PLAN-REVIEW — s16a WP2 PLAN r12 must revise the owner-side matrix and per-commit gate

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16a-build-plan-review-7
PARENT_DISPATCH_ID: s16a-build-12
RUN_ID: s16a
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — s16a.planner can issue a bounded successor; WP2 implementation and merge remain separately gated
PLAN_LOCK_ID: s16a-build-12
IN_REPLY_TO: s16a-build/PLAN-planner-20260825-011947.md
FROM: s16a.implementer
TO: s16a.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
SUBJECT: must-revise — restore B08 W and A16 CP, enumerate all cross-side rows, and fix the impossible tagged-battery-green law
PLAN_REVIEW_VERDICT: must-revise

`PLAN_REVIEW_VERDICT: must-revise` for exact PLAN r12 SHA-256 `2c766189f7a4d489493e4a4acdbdee31610494b06a5d9c2b6280b5feb6f15f23` at `s16a-build/PLAN-planner-20260825-011947.md`.

The engine admitted those exact bytes under `s16a-build-12`, parented to `s16a-impl-4`. The implementation worktree is clean on `s16a-conformance` at `f70406668d3db96d882f879d90c6569c7be858b3`, with the accepted 21 GREEN / 43 RED / 64 coda. The plan correctly targets 38 distinct rows and a 59 GREEN / 5 RED close. Two execution blockers remain.

## Blocking findings

- **S16A-WP2-PR-F1 — two required owner halves are omitted.** Locked r7 §2 / r9 require A16 on CP+W and B08 on CN+W. R12's W enumeration has 25 halves and omits B08's W total-consumption half; CP has 14 and omits A16's CP half. The full cross-side set is `{A13,A14,A16,B08,C01,C08}`, not the four in §4. Calling their second half “conforming” contradicts the ledger, which requires both halves to move. The SUBJECT's `W 18` is also unsupported.
- **S16A-WP2-PR-F2 — §3.1's tagged-battery green gate is impossible.** The tagged battery begins WP2 with 43 intentional REDs and retains five at WP2 close. Requiring it to be “green at every commit” cannot be met before WP5 and contradicts 59/5. The tagged instrument must run and derive a valid moving census, not exit green while planned rows remain RED.

## Required bounded successor

1. Reissue the WP2 plan parented to this review, preserving 38 distinct rows and the 59/5 close bar.
2. Add B08 W, yielding 26 W halves; add A16 CP, yielding 15 CP halves; retain 3 CN halves. State all six cross-side rows and use first/second or paired-half sequencing, never “conforming-half.” Correct the SUBJECT/count prose. The complete matrix is 44 owner halves over 38 rows.
3. Make each commit require its named row/cluster GREEN for the contract reason, a valid tagged 64-row census with only then-planned remaining REDs, and the plain suite plus vet GREEN. Preserve exact 59/5 at close and the no-unexpected-delta rule.
4. Preserve all other r12 substance: non-conforming-side-only, coupled-test restriction, census modernization, A-2 carries, WP3/WP4 exclusions, A14 registration request, boundary contract, no E3, and no push/PR/merge.

## Accepted surfaces not reopened

- Pins re-compute: plan r7 `5fd00b98727e2f518cba0976f8da7a2fc4cffd6c124ca593b3dde7bb4d02193f`; r9 `acfd358ae3ccb6b250da1dca3fa24625b8e60e4afce8cdf26a60d0ca41f2cf06`; A-1/E-1 `4d9b0cb07d88e1014cb4c710dcca13b14357c6a03da54de3f83de42c252164f1`; A-2 `899157bd9de7ea166bc968ba1edd1bb0b9855207996f6b0a106baedc404fcd29`; charter `41c43b279eae4d6c4ad97cea7e3fd27b5e309eeefc44db5a64e2f3570fe1bfa2`; WP0(c) `b04507774a424a2e6c10ec2b4630666ddbccd2f49cadbde957348a2322085bac`.
- The distinct-row arithmetic is sound once the owner halves are complete: 38 REDs turn GREEN, moving 21/43 to 59/5; only A12, C07, C09, D01, D02 remain RED.
- A-2 §3 authorizes the bounded `census.py` modernization while preserving sentinel and bijection.
- The named production trees, coupled seam literals, and census tool stay inside the charter fence. WP3/WP4 rows, contracts/ledger/master, `go.mod`/`go.sum`, governed s8, exit fixtures, E3, push, PR, and merge remain out.

Boundary-contract disposition: target, consumers, E2 proof, and close criterion suffice. Approval is withheld because A16/B08 writers are omitted and the per-commit proof is not executable.

Remaining gates: no SCOPE_DIFF or implementation token may issue on this verdict. A corrected successor needs approval, then mechanical SCOPE_DIFF and a separately addressed implementation relay. Merge remains operator-only at WP5.

Tests / verification: E1 exact-byte review; r7/r9/A-2 reconciliation; mechanical counts (r12 W 25, CP 14, CN 3 versus required W 26, CP 15, CN 3; 44 owner halves over 38 rows); engine admission/INDEX/hash; clean implementation-worktree head/status. No source test or implementation action was authorized.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — WP2 repairs authority and fail-closed seams; review-only here
- migration/backfill/destructive-write/canonical-data-repair: yes — C04 changes persistence; no review-time write
- money/inventory/orders/planning/accounting/trust-critical-state: yes — trust-critical seam
- AI-or-automation-acts-downstream: yes — worker and tool-authority behavior changes
- worker/scheduler/queue/retry/async-side-effect: yes — worker, scheduler, connector, epoch, attempt lifecycles
- cross-repo/service-contract/generated-schema/shared-API-event: yes — registered cross-module contracts
- user-visible-control-with-materializer/downstream-consumer: no
- test-runtime-role-mismatch: yes — intentionally RED battery is F2's basis
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — medium ceremony retained; F1/F2 require revision
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade or risk acceptance requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

No implementation action: no production/test byte, branch state, dependency, stage, commit, push, PR, merge, provider, credential, store, or runtime state changed.

ACTIONS_GIT_REF: engine-lane governance act only — drafted under `frank/.relays/s16a/.engine/drafts/s16a.implementer/` for submission; no source/test or implementation-worktree action claimed
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s16a/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s16a/s16a-build/PLAN-REVIEW-implementer-20260824-225758.md
?? frank/.relays/s16a/s16a-build/PLAN-REVIEW-implementer-20260824-233824.md
?? frank/.relays/s16a/s16a-build/PLAN-planner-20260824-225355.md
?? frank/.relays/s16a/s16a-build/PLAN-planner-20260824-232655.md
?? frank/.relays/s16a/s16a-build/PLAN-planner-20260825-011947.md
?? frank/.relays/s16a/s16a-impl/IMPL-planner-20260824-234004.md
?? frank/.relays/s16a/s16a-impl/SITREP-implementer-20260825-011602.md
