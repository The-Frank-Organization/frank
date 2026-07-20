## DESIGN-REVIEW - m-10 stage-1 IPC/manifest seam r7 must revise: recovery transition substate is incomplete and consumed tickets are still voided

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r8
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - both findings are bounded crash-state/disposition corrections under the locked recovery and F59 choices; Master+VP retain interface-lock authority
GRILL_REQUIRED: no - this re-review does not reopen the operator-locked topology, threat ceiling, F59, F60, or revocation-first choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260716-103500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260716-082658.md
SUBJECT: must-revise exact d0e56da7... - role-exact reveal proof closes, but recovery lacks total epoch-transition substates/ordering and the canonical retirement still maps consumed unknown-outcome tickets to VOID

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner - I reviewed the exact revised bytes at SHA-256 `d0e56da7a67b569354ea62998eb3c6e336caf8e3bbc6c66eb9e817fc17eabaa8`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, the r7 basis hash, and both consumed sibling hashes pass. R7-F1 closes. The keyed one-mint guard is the right recovery basis, but the recovery matrix is not yet total across §B.5 substates; the full crash-disposition sweep also found one F59 contradiction in the canonical retirement transaction.

## Findings

### R8-F1 - Recovery cannot execute its stated §B.5 action for all post-retirement crash cuts

The matrix says a committed `RETIRED_PENDING_REAP` state resumes/reconciles the same epoch and transition by ID (`2026-07-16-mvp-ipc-manifest-seam-contract.md:84`), then says broker adoption/fresh spawn happens afterward (`:87`). §B.5 recovery, however, is explicitly performed on control re-establishment (`:115`), so reconciliation cannot precede CI-1 adoption/spawn.

The matrix also leaves two durable transition substates unresolved:

- crash after the keyed retirement transaction commits E+1/G+1 (`:95`) but before step 2 creates/proposes an `epoch_transition_id` (`:96,111`): there is no transition ID to “resume/reconcile”; and
- crash after the broker transition reaches terminal INSTALLED but before reap/spawn/lease: branch (b) names a non-terminal transition, yet the correct action is to resume lifecycle step 3/4 under the installed epoch without another handshake or mint.

The keyed `(run_id, retiring generation_id)` guard prevents a second retirement mint, but does not by itself choose the correct broker/lifecycle continuation. A T4 implementation cannot derive a total ordering from the current branch plus the common suffix.

Required return: make broker control establishment precede any §B.5 query/reconciliation and make branch (b) total over transition substate. Either durably allocate the stable `epoch_transition_id`/ledger row inside the retirement transaction, eliminating the no-ID cut, or explicitly define “retirement committed, no transition row” ⇒ after broker adoption/spawn, propose a fresh ID for the already-current E+1. Then distinguish non-terminal transition ⇒ resume/abort per §B.5; INSTALLED ⇒ continue at reap/spawn with no new transition; ABORTED ⇒ propose the permitted fresh attempt for the same E+1 before lifecycle continuation. In every subcase, broker E+1 install must precede successor lease/worker assign, and no subcase re-mints.

### R8-F2 - Canonical retirement maps all tickets to VOID, contradicting the F59 unknown-outcome state

The canonical retirement transaction still says row parking includes `tickets → VOID` (`:95`). D.4 requires a state-sensitive split: ISSUED/unconsumed tickets become VOID, but a CONSUMED ticket with no outcome parks both the ticket and tool-call row as `UNKNOWN_TOOL_OUTCOME` (`:196-199`). The new recovery suffix further describes consumed-ticket-no-outcome as outside the retirement transaction's scope (`:87`), while B.3 says the retirement transaction both fences and records crash disposition in one transaction (`:74-78`) and the matrix says retirement atomically owns its row parking (`:82-83`).

If `tickets → VOID` is applied literally after consume, the honest unknown-effect record is erased and an unsafe retry can be mistaken for a clean pre-execution death. If it is excluded and parked later, the claimed atomic crash disposition is false.

Required return: make §B.4 step 1 state-sensitive and atomic: `ISSUED → VOID`; `CONSUMED` with no recorded outcome → `tool_authorizations.state = UNKNOWN_TOOL_OUTCOME` AND matching `tool_calls.state = UNKNOWN_TOOL_OUTCOME`; terminal outcome rows remain terminal. Remove consumed-ticket-no-outcome from the “outside retirement scope” example for an active-generation retirement. Preserve the same mapping in app-main recovery branches that do not execute a retirement transaction, before new admission.

## What closes from review r7

- R7-F1 closes: `assign` is now only worker/presenter-facing, the broker is explicitly the verifier-only recipient, m-8 remains generation-blind, and §B.1/§B.4 use the same exact reveal set.
- R7-F2 closes at the retirement-identity level: `(run_id, retiring generation_id)` is the one-mint guard; active leased, already-retired, pre-lease wash-out, and initial E=1 branches are separately named; recovery no longer unconditionally re-retires.
- Spawn-only fail-closed child recovery, connector bootstrap/READY gating, R4-F2, and the prior broker CI, manifest/F63, DATA-P pairing, PREPARING-ledger, attempt/event ingress, and canonical wire-counter repairs remain present.

## Gate disposition

This verdict is byte-bound to `d0e56da7a67b569354ea62998eb3c6e336caf8e3bbc6c66eb9e817fc17eabaa8`. Any revision requires a new SHA and fresh uniquely-parented DESIGN-REVIEW. Consumer confirmations, the Master+VP interface lock, stage-3/5 designs and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact design SHA-256 recomputed: `d0e56da7a67b569354ea62998eb3c6e336caf8e3bbc6c66eb9e817fc17eabaa8`.
- Ratified amendment SHA-256 recomputed: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Pair-approved m-7 r6 SHA-256 recomputed: `f072bd996da0c85b1be9b67fad880e7395ab03de7142cb87fc6864d4f67a100e`.
- Pair-approved m-3 r2 SHA-256 recomputed: `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44`.
- Incoming DESIGN relay exact-file lint: OK.
- Exact crash-cut sweep: recovery branch/action/common suffix at `:82-87`; retirement before transition proposal at `:95-96,111`; control-recovery prerequisite at `:115`; canonical tickets-to-VOID at `:95`; consumed UNKNOWN requirement at `:196-199`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`
RELAY_LINT: OK - exact-file mode exit 0
Next requested action: m-10.planner folds R8-F1 and R8-F2 into fresh design bytes, recomputes the SHA-256, files a new uniquely-parented DESIGN relay, and requests fresh m-10.implementer review; do not route consumer confirmations on the current bytes.
