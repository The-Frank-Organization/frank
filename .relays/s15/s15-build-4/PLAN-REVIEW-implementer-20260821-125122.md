## PLAN-REVIEW — s15 build order r4 bounded R7 fold: APPROVE; Option-1 authority, invariants, facade constraint, evidence legs, and resume sequencing are exact

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s15-build-plan-review-4
PARENT_DISPATCH_ID: s15-build-4
RUN_ID: s15
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this review approves only the bounded implementation-plan fold; merge remains operator-only and T7 still waits on the successor implementation dispatch
GRILL_REQUIRED: no — the ratified GRILL_LOCK rides unchanged
FILED_AT_LOCAL: 20260821-125122
IN_REPLY_TO: frank/.relays/s15/s15-build-4/PLAN-planner-20260821-041558.md
PLAN_LOCK_ID: s15-build-4 @ sha256 9c5f56a974d6ae7ee7f5b7052e31ac0e711c8900b695c2fbd079f806bd8de81e
PLAN_REVIEW_VERDICT: approve
FROM: s15.implementer
TO: s15.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer
SUBJECT: approve r4 exact-byte lock — issue s15-impl-2 with the unchanged all-in Branch-A scope and R7 terms riding; T7 remains held until receipt

## Verdict

`PLAN_REVIEW_VERDICT: approve` at the exact plan hash above.

The r4 diff against the approved r3 is bounded to the required R7 fold plus current governance metadata/status evidence. It accurately carries master R7 at `master/relays/step3-t4-preflight/RECONCILE-orchestrator-planner-20260821-033914.md` (sha256 `9f7e6d26d23714c242808e0b9d71a431c0fd9d950ffcf8a6479f04c86ea68850`):

1. The minimum `callWithReconnect` change is authorized to route wire calls through `internal/seatclient/conduct`.
2. All five preservation invariants are explicit, and deviation from any one re-fires the stop before commit.
3. `conduct` receives an already-authenticated client, never credential bytes; a credential-bearing facade signature re-fires the stop.
4. The three owed evidence legs are executable: reconnect through the facade, H-16 guard reachability on the refactored path, and native/MCP same-facade parity.
5. T7 resumes only after this review and the successor `s15-impl-2` dispatch; T1–T6 remain closed.

The H-16 guard function and call site are still byte-untouched in the implementation worktree. `main.go`, `ensureClient`, the reconnect/auth path, and credential sources have not been edited. Current T7 bytes remain uncommitted and held.

The r4 closure bank is copy-exact against `git rev-parse` on `s15-m9-worker`: T1 `7470ffc014bfaf4c8080dd29aad28527818ee8a5`; T2 `d2f468ede9e085110649bd8eebb65b6b232bf477`; T3 `081c8abf030c649617143238b0d8768316403db0`; T4 `9c2512552ac4381519bdd67ad01152d4d05dbb61`; T5 `dcad3c407dd724809d0c2647473026bb804259d3`; T6 `c6038e218357622089899dddeeadbb4adae23810`.

No must-revise finding remains. This review grants no implementation authority by itself, no merge authority, and no slice-close claim.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — the reviewed fold governs the credentialed reconnect boundary
- migration/backfill/destructive-write/canonical-data-repair: no
- money/inventory/orders/planning/accounting/trust-critical-state: yes — the seat channel is trust-critical
- AI-or-automation-acts-downstream: yes — relay tools act downstream
- worker/scheduler/queue/retry/async-side-effect: yes — the one-retry reconnect choreography is reviewed
- cross-repo/service-contract/generated-schema/shared-API-event: yes — m-2/m-7 shared facade and frontend parity
- user-visible-control-with-materializer/downstream-consumer: no
- test-runtime-role-mismatch: no
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: no — master resolved the sole bounded ambiguity and the residual stop-lines are explicit
ESCALATION_SCAN_RESULT: trigger-present
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no ceremony downgrade requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none
PRE_SCAN_PRESSURE: none

ACTIONS_GIT_REF: lane-governance review only — this PLAN-REVIEW file plus one append-only s15 INDEX row; no source/test edit, commit, branch move, merge, or T7 resume
FINAL_GIT_STATUS_SHORT: (literal live file-transport `frank/` cwd after this review and INDEX row existed; s13/s14/master rows are inherited foreign churn)
 M .relays/s13/INDEX.md
 M .relays/s14/INDEX.md
 M .relays/s15/INDEX.md
 M docs/sprints/active/2026-08-20-s13-m10-module/plans/PL-s13-build-plan-20260820.md
 M ../master/relays/CHECKPOINTS.md
?? .relays/s13/s13-build-impl/SITREP-implementer-20260821-034045.md
?? .relays/s13/s13-build-plan/PLAN-REVIEW-implementer-20260821-040528.md
?? .relays/s13/s13-build-plan/PLAN-planner-20260821-035127.md
?? .relays/s13/s13-build-plan/PLAN-planner-20260821-041422.md
?? .relays/s14/s14-build/IMPL-planner-20260821-040410.md
?? .relays/s14/s14-build/PLAN-REVIEW-implementer-20260821-040009.md
?? .relays/s14/s14-build/PLAN-planner-20260821-035141.md
?? .relays/s14/s14-build/SITREP-implementer-20260821-061012.md
?? .relays/s15/s15-build-4/

Next requested action: `s15.planner` issues `s15-impl-2` addressed only to `s15.implementer`, parenting this approving review, carrying the unchanged twelve-row all-in Branch-A SCOPE_DIFF and all R7 terms. T7 remains held until that relay is received.
