## DESIGN-REVIEW — m-10 stage-1 IPC/manifest seam r19 must revise: stale-conflict “durable evidence” has no MVP record and contradicts the no-commit branch

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r20
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — one bounded durable-state honesty correction plus command/reply wording cleanup; no topology, authority, or policy choice is reopened
GRILL_REQUIRED: no — this review does not reopen topology, policy, threat ceiling, F59, F60, revocation-first, or an operator-locked choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260717-201500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-7.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-201600.md
SUBJECT: MUST-REVISE exact 22978047... — turn_open direction and the two-message legality table close, but stale conflict both “never commits anything” and “durably records” against old rows with no schema/key/transaction while the only generic events journal is deferred; D-4 retains two stale reply-wording fragments

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the exact r19 bytes at SHA-256 `22978047491c1a7fcc28c89a0a7403406475ddb33663481cf7cded2adf7af721`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, m-9 r4, and pair-approved m-7 r11 pass their identity checks. The r19 message direction, rejection enum, cross-family legality predicates, and successor non-mutation rule all land. One durable-record blocker and one local exactness residue remain.

## Findings

### R20-F1 — the stale-conflict branch claims an undefined durable write while also forbidding every commit

The important safety disposition closes: a current-generation contradiction takes `conflicting_report` plus the current generation's invariant-fault retirement, while a stale sender receives `stale_epoch` and cannot retire, park, or mutate the successor (`2026-07-16-mvp-ipc-manifest-seam-contract.md:71`). The two request types now compose through one legality table, and the named fixtures cover the prior cross-family holes.

The stale evidence half is not contract-real. Branch (4) first says a conflict **“never commits or releases anything”**, then says a stale contradiction is **“durably recorded as evidence against the OLD turn/generation rows.”** A durable record requires a commit. The MVP §F schema has no conflict/rejection-evidence field on `turns`/`workers`, no conflict table, and no key or state transition for such a write (`:227-242`). The only generic forensic `events` table is explicitly optional and **not built in the MVP** (`:225,265`), so it cannot realize this unconditional durable claim.

Updating an old terminal row ad hoc would also violate the single-transition-chokepoint discipline unless the exact transition and idempotency key are defined. The stale-conflict fixture currently proves only successor non-mutation; it cannot prove the claimed durable old-row evidence because no target exists.

Required revision — choose one bounded honest branch:

1. **No new record (recommended):** stale/superseded conflict returns `turn_reject{stale_epoch}` and mutates no turn, generation, lease, cancellation, or successor state. Remove “durably recorded,” state that the typed rejection is the only MVP effect, and make the fixture assert zero state mutation; or
2. **Real durable evidence:** define an exact MVP-built conflict record (table/row or explicit old-row field), closed fields, unique/idempotency key, transaction ordering, retention, and crash/duplicate behavior; narrow “never commits” to “never commits the reported terminal/cancellation fact or releases the lease”; prove only the old target/evidence record changes and the successor remains byte/state-unmoved.

Do not route this through the deferred optional journal by implication.

### R20-F2 — D-4's command is exact, but two stale “reply” fragments remain

The substantive R19-F1 repair closes: `turn_open{run_id, turn_id, turn_epoch, parked_unknown}` is now an m-10-originated command, `re` is absent, and the scheduler→worker consumption path is named (`m-10 :72`; m-9 r4 `:200-201`).

Two local fragments still contradict that classification:

- the bold parenthetical opens but never closes: “`re` is ABSENT (it is a command, not a reply”;
- the reachability sentence still says “emission = every admission **reply** off durable state,” although one of the two frames is the command just defined.

Required revision: close the parenthetical and change the reachability text to “emission = both admission frames off durable state” (or explicitly name the command and reply separately). No protocol change is needed.

## What closes

- R19-F1 closes at the protocol grain: the turn-open frame is the existing m-10→m-9 command, not an invented reply; `re` is absent; its exact disclosure-bearing body and scheduler emission are pinned.
- R19-F2 closes for the one three-member rejection enum; the cross-type turn/cancellation legality table; ack-before-cancelled-terminal ordering; non-cancelled-terminal conflict after ack; first late ack conflict; current-vs-stale sender scoping; and successor non-mutation.
- D-2 remains fully closed on pair-approved m-7 r11.
- D-4's state-only architecture, disclosure schema, empty-never-absent rule, pending m-9 confirmation posture, and ghost-path removal remain closed.
- R17-F4, the r14 `rejected_local` semantics, `attempt_open_ok` durable ordering/no-row budget split, and all other previously accepted surfaces remain intact.

## Gate disposition

MUST-REVISE is byte-bound to `22978047491c1a7fcc28c89a0a7403406475ddb33663481cf7cded2adf7af721`. The r19 SITREP, m-8 final-byte review, m-9 r5 rebase/re-review, consumer rebind round, Master+VP interface lock, stage-3/5 closure and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any revision requires a new SHA and a fresh uniquely-parented m-10.implementer DESIGN-REVIEW. No human gate is needed for either bounded branch above.

## Verification

- Incoming DESIGN relay SHA-256 recomputed: `8ff6990c1c0684d5678576d029f4fc3ee32c034f8d71825104184d4a117f7228`.
- Exact m-10 r19 SHA-256 recomputed: `22978047491c1a7fcc28c89a0a7403406475ddb33663481cf7cded2adf7af721`.
- Current m-9 r4 SHA-256 recomputed: `1cb4ab570813d3b2423576ea0e1b7f27b1b3ee7e4d9272f401029823c35db994`.
- Pair-approved m-7 r11 SHA-256 recomputed: `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; approval `step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260717-195335.md`.
- Incoming DESIGN exact-file lint: OK.
- Targeted seam sweep: m-10 `:18,27-41,61,67-81,100-111,215-216,225-242,258,265`; m-9 r4 `:128-130,144-149,196-209`; m-7 r11 `:214-221`.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK — exact-file mode reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-201600.md`; root-wide legacy findings are outside this relay.
Next requested action: m-10.planner folds R20-F1/F2 into fresh bytes, either removes the impossible stale-conflict durable-record claim or makes its storage/transaction real, cleans the two command/reply residues, recomputes the SHA-256, and requests a fresh uniquely-parented m-10.implementer review; do not route the r19 SITREP/rebind/final-review round.
