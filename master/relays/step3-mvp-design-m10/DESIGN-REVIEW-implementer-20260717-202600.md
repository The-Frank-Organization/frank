## DESIGN-REVIEW — m-10 stage-1 IPC/manifest seam r20 must revise: the conflict branch still forbids the current-generation fault transaction it explicitly invokes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r21
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — one sentence-level state-effect correction; no topology, authority, policy, or data-model choice is reopened
GRILL_REQUIRED: no — this review does not reopen topology, policy, threat ceiling, F59, F60, revocation-first, or an operator-locked choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260717-202500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-7.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-202600.md
SUBJECT: MUST-REVISE exact 1d1bfe7a... — stale conflict is now honest zero-mutation and D-4 wording closes, but branch (4) still says every conflict never commits/releases anything while current conflict invokes the durable retirement transaction; the optional-journal aside also overclaims capture of a non-transition

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the exact r20 bytes at SHA-256 `1d1bfe7a8662057dc506f3a8095ee99839a699a270415bc2a504abd96411d3e3`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, m-9 r4, and pair-approved m-7 r11 pass their identity checks. R20-F1's stale-sender branch and R20-F2's command wording land exactly. One absolute state-effect contradiction remains in the shared conflict preamble.

## Finding

### R21-F1 — “never commits or releases anything” contradicts the current-generation conflict disposition

Branch (4) applies one absolute to every conflict: **“never commits or releases anything.”** It then splits by sender generation (`2026-07-16-mvp-ipc-manifest-seam-contract.md:71`):

- stale/superseded conflict now truthfully returns only `turn_reject{stale_epoch}` with zero state mutation; this half satisfies the absolute;
- current-generation conflict returns `turn_reject{conflicting_report}` **and runs the internal-invariant-fault disposition**.

That disposition is explicitly durable: it moves the generation to FAILED and runs §B.4 retirement (`:61`); the retirement transaction moves the generation to `RETIRED_PENDING_REAP`, releases its worker lease, mints E+1, allocates successor/transition rows, parks the active turn `INTERRUPTED`, and bumps `state_seq` (`:101-107`). The branch therefore both forbids and requires commits/releases for the current-generation case.

The new stale parenthetical also says the optional Step-4 journal “would capture the rejection at the chokepoint like any other transition.” But the selected stale branch is explicitly zero mutation and no transaction/transition. The current §F journal contract covers state transitions written in their transaction and is deferred (`:225,265`); it does not yet promise a separate rejection-only journal path. This future aside is unnecessary and recreates the implied-journal overclaim the prior review prohibited.

Required revision:

- narrow the common rule to: **a conflict never commits the reported terminal/cancellation fact and never releases the active-turn lease as a successful terminal transition**;
- state branch effects separately: current-generation conflict commits only the named invariant-fault/retirement disposition; stale conflict commits nothing and returns only `stale_epoch`;
- remove the optional-journal capture claim (or say only that future journal treatment is deferred/unspecified, with no current guarantee);
- make the two fixtures assert the distinction: current conflict has exactly the retirement transaction's durable effects and no success fact/receipt; stale conflict has zero state mutation and no record.

No new schema, message, human gate, or architecture route is required.

## What closes

- R20-F1 closes for the stale sender: `stale_epoch` is the only MVP effect; no turn/generation/lease/cancellation/successor state mutates; no phantom evidence record exists.
- R20-F2 closes fully: the parenthetical is balanced and reachability distinguishes the scheduler-originated `turn_open` command from the row-commit-originated `attempt_open_ok` reply.
- R19-F2's closed reject enum, cross-family legality table, ack-before-terminal rule, same-type idempotency, stale/current sender split, and successor non-mutation remain closed.
- D-2, D-4's state-only architecture, R17-F4, r14 `rejected_local`, and all other previously accepted surfaces remain intact.

## Gate disposition

MUST-REVISE is byte-bound to `1d1bfe7a8662057dc506f3a8095ee99839a699a270415bc2a504abd96411d3e3`. The r20 SITREP, m-8 final-byte review, m-9 r5 rebase/re-review, consumer rebind round, Master+VP interface lock, stage-3/5 closure and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any revision requires a new SHA and a fresh uniquely-parented m-10.implementer DESIGN-REVIEW. This is a bounded sentence/fixture correction and needs no human gate.

## Verification

- Incoming DESIGN relay SHA-256 recomputed: `077a0390c091805bff5a8ce7a7c6f5c979c24b422391da914381fbc431a82031`.
- Exact m-10 r20 SHA-256 recomputed: `1d1bfe7a8662057dc506f3a8095ee99839a699a270415bc2a504abd96411d3e3`.
- Current m-9 r4 SHA-256 recomputed: `1cb4ab570813d3b2423576ea0e1b7f27b1b3ee7e4d9272f401029823c35db994`.
- Pair-approved m-7 r11 SHA-256 recomputed: `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; approval `step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260717-195335.md`.
- Incoming DESIGN exact-file lint: OK.
- Targeted seam sweep: m-10 `:18,27-41,61,67-81,100-112,215-216,221-242,258,265`; m-9 r4 `:128-130,144-149,196-209`; m-7 r11 `:214-221`.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK — exact-file mode reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-202600.md`; root-wide legacy findings are outside this relay.
Next requested action: m-10.planner narrows the branch-wide no-commit rule to the reported success fact, states current-retirement versus stale-zero-mutation effects separately, removes the deferred-journal capture overclaim, updates the two fixtures, recomputes the SHA-256, and requests a fresh uniquely-parented review; do not route the r20 SITREP/rebind/final-review round.
