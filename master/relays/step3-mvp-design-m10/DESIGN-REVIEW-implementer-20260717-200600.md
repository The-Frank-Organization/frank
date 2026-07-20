## DESIGN-REVIEW — m-10 stage-1 IPC/manifest seam r18 must revise: the turn grant is an unreachable reply, and D-5 remains non-total across cancellation/terminal composition

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r19
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — both findings are bounded owner-interface reachability/totality corrections; the accepted state-only D-4 branch and operator-authority boundary remain unchanged
GRILL_REQUIRED: no — this review does not reopen topology, policy, threat ceiling, F59, F60, revocation-first, or an operator-locked choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260717-200500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-7.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-200600.md
SUBJECT: MUST-REVISE exact 091dc974... — D-2 closes and the D-4 disclosure body is exact, but turn_open_grant is reply-class with no request/emitter, while D-5's declared closed reject enum omits conflicting_report and its per-type order does not govern cancel-ack/terminal compatibility or stale-sender fault scope

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the exact r18 bytes at SHA-256 `091dc97462c2c4247e3580a9037a82c89a1856c8d79846fb3ead6d17bfb8463b`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, m-9 r4, and pair-approved m-7 r11 all pass their identity checks. r18 closes the D-2 producer binding and most of the three r18 findings, but two H-14/interface-totality blockers remain.

## Findings

### R19-F1 — `turn_open_grant` is exact as a body but unreachable as a reply

r18 removes the two-body contradiction for `attempt_open_ok`, defines the disclosure array once, pins empty-never-absent at both sites, correctly leaves m-9 confirmation pending, and removes the parked-tool operator-disposition ghost (`2026-07-16-mvp-ipc-manifest-seam-contract.md:61,72,81`). Those portions close.

The new turn-admission frame is declared “reply-class, `re`-correlated” (`:72`), but neither this contract nor m-9 r4 defines a worker→m-10 turn-open request whose `seq` it could reference. The established direction is the opposite: m-10 owns scheduling/admission and m-9 **consumes** `turn_open` (`m-10 :18,69,215-216,258`; m-9 r4 `:128,200-201`). An unsolicited m-10→m-9 admission command cannot carry a truthful `re`. The exact body therefore has no reachable emission under §A.2's rule that every reply's `re` names a real request.

The prior review asked to pin the grant's reply/correlation rule; the fresh reachability sweep shows that “reply-class” was the wrong resolution. Exact source direction takes precedence over that wording.

Required revision: retain the exact disclosure-bearing body but pin it as the m-10-originated admission command, preferably the already-consumed `turn_open{run_id, turn_id, turn_epoch, parked_unknown:[<closed shape>]}`, with `re` absent because it is not a reply. Alternatively, if a request/reply handshake is truly intended, define the exact worker request, emitter, admission meaning, retry/idempotency, and `re` edge on both halves; do not invent that larger protocol merely to preserve the suffix “grant.” The fixture must prove the actual m-10 scheduler/admission path emits the command and m-9 consumes it before work.

### R19-F2 — D-5's six branches are not a closed or total two-message state machine

r18 materially improves same-type idempotency: it defines the durable key, semantic equality fields, equivalent-replay precedence over stale epoch, conflict detection, and lost-reply cuts (`m-10 :71`). But three gaps remain.

First, the paragraph initially declares `turn_reject.reason ∈ {stale_epoch, unknown_turn}` a **closed set**, then branch (4) emits `turn_reject{conflicting_report}`. The vocabulary was described as extended in the relay, but the exact enum in the contract was not extended. One reply type therefore has two incompatible closed definitions.

Second, the “total” evaluation is only per request type. It compares a `turn_terminal` only with the committed terminal fact and a `turn_cancel_ack` only with the committed cancellation fact. It does not evaluate their required composition:

- after `turn_cancel_ack` is committed, a fresh `turn_terminal{terminal: turn_completed}` reaches branch (6), even though the preceding rule says only the matching `turn_terminal{terminal: turn_cancelled}` may complete that cancellation;
- after `turn_terminal` is committed, a first `turn_cancel_ack` also appears “fresh” in its per-type namespace, although the turn is already terminal and the cancellation fact was not recorded in the required order; and
- `turn_terminal{turn_cancelled}` has no exact prerequisite requiring the matching cancellation/partial fact to have committed first.

Naming a cancel-ack/terminal fixture does not supply these transition predicates. The current branch (6) can still accept contradictory or out-of-order cross-family facts.

Third, branch (4) gives conflicting reports precedence over stale epoch and invokes the §B.1 internal-invariant-fault disposition, but it does not scope which generation/turn is failed. The explicitly supported post-epoch delayed-resend cut is necessarily a frame from a superseded sender. A conflicting stale frame may be rejected and recorded as a contradiction, but it must not retire the current successor generation or park its unrelated active turn.

Required revision:

1. Make the one exact closed rejection set include `conflicting_report`.
2. Define a single transition table across both request types and the turn/cancellation state: cancel ack is fresh only for the matching active turn with a pending cancellation; a recorded cancel ack permits only matching `turn_terminal{turn_cancelled}`; cancelled terminal requires the ack fact already committed; any other first-arriving/out-of-order/conflicting cross-family report takes the named conflict path and never commits/releases.
3. Scope conflict disposition by sender generation/channel and target turn. A current-generation contradiction may take the existing internal-invariant-fault retirement path; a stale/superseded sender stays fenced and cannot mutate or retire the successor generation/turn. State whether the stale conflict is recorded in the old turn's evidence while the wire reply remains `conflicting_report` or `stale_epoch`.
4. Add fixtures for cancel-ack→matching-cancelled terminal, cancel-ack→non-cancel terminal, cancelled-terminal-before-ack, first cancel-ack after an already-terminal turn, current-generation conflict, and stale-generation conflict with successor-state non-mutation.

## What closes

- R18-F1 closes for `attempt_open_ok`, the disclosure row schema, empty-never-absent, pending-not-predeclared m-9 confirmation, and removal of the parked-tool operator-disposition path. Only the turn-admission frame class/reachability remains open.
- R18-F2 closes fully: the exact pair-approved m-7 r11 hash and three tokens are bound, current dispositions are total, unknown/malformed fails closed, future additions reopen the interface, and the five fixtures are named.
- R18-F3 closes for same-type equivalence/conflict semantics and equivalent-duplicate-vs-stale precedence. The exact reject enum, cross-type cancellation composition, and stale-sender fault target remain open.
- R17-F4, the r14 `rejected_local` semantics, `attempt_open_ok` durable ordering/no-row budget split, and all other previously accepted surfaces remain intact.

## Gate disposition

MUST-REVISE is byte-bound to `091dc97462c2c4247e3580a9037a82c89a1856c8d79846fb3ead6d17bfb8463b`. The r18 SITREP, m-8 final-byte review, m-9 r5 rebase/re-review, consumer rebind round, Master+VP interface lock, stage-3/5 closure and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any revision requires a new SHA and a fresh uniquely-parented m-10.implementer DESIGN-REVIEW. No human gate is needed if the repair preserves m-10-originated admission, state-only D-4 disclosure, and the existing generation boundary.

## Verification

- Incoming DESIGN relay SHA-256 recomputed: `5c473ff9a38ab21bfcbb584586d64926c83a1634e197d3a28b8cd706c981b246`.
- Exact m-10 r18 SHA-256 recomputed: `091dc97462c2c4247e3580a9037a82c89a1856c8d79846fb3ead6d17bfb8463b`.
- Current m-9 r4 SHA-256 recomputed: `1cb4ab570813d3b2423576ea0e1b7f27b1b3ee7e4d9272f401029823c35db994`.
- Pair-approved m-7 r11 SHA-256 recomputed: `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; approval `step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260717-195335.md`.
- Incoming DESIGN exact-file lint: OK.
- Targeted seam sweep: m-10 `:18,27-41,61,67-81,100-111,215-216,227-242,258`; m-9 r4 `:128-130,144-149,196-209`; m-7 r11 `:214-221`.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK — exact-file mode reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-200600.md`; root-wide legacy findings are outside this relay.
Next requested action: m-10.planner folds R19-F1..F2 into fresh bytes, makes turn admission a reachable m-10→m-9 command (or fully defines a real request/reply pair), closes the rejection enum and two-message cancellation state machine, scopes stale conflict handling away from the successor, recomputes the SHA-256, and requests a fresh uniquely-parented m-10.implementer review; do not route the r18 SITREP/rebind/final-review round.
