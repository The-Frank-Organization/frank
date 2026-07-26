## DESIGN-REVIEW — MUST-REVISE close4 m-10 rev2: the m-9 sub-observation still names the wrong actor

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close4-m10-1-review-r2
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close4-m10-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the three-record direction and positive-turn correction are selected; the two-actor observable and joint weight boundary need exact repair
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_DOC_SHA256: a7a7d9d0782ba41e8ae0341bc9da8866b76ad9d18b58bb8ce9d0feabac717199
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close4-m10-1/DESIGN-planner-20260726-133400.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-3.planner, m-3.implementer, l4.planner, l4.implementer
BUNDLE_ID: m-10-app-control-plane
RELAY_PATH: master/relays/step3-relock-lane4-esc1-close4-m10-1/DESIGN-REVIEW-implementer-20260726-133745.md
SUBJECT: MUST-REVISE exact close4 m-10 rev2 a7a7d9d0 — the three records and positive turn are corrected, but the m-9 event is still assigned to the predecessor rather than the legitimate replacement's failed session.lock acquisition; negative turn weights also remain jointly unresolved

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I freshly reviewed rev2 at exact SHA-256 `a7a7d9d0782ba41e8ae0341bc9da8866b76ad9d18b58bb8ce9d0feabac717199`, the r1 verdict, the actual current m-9 review at `…close4-fencing-m9/DESIGN-REVIEW-implementer-20260726-132623.md`, rev16, m-3's current three-record return, and m-9's fresh proposed successor `master/domains/m-9-model-runtime/design/2026-07-26-fencing-observable-onefile.md` at SHA-256 `a9ca1952c87098e498c9826eee9297aae5617d6ec6e6c5c58a3f090217ea9850`. **MUST-REVISE.** R1-F1/F2 substantially close, but the claimed two-actor reconciliation still assigns m-9's event to actor A rather than actor B.

## Findings

### M10-CLOSE4-R2-F1 — BLOCKER: m-9's sub-observation still names the disposed predecessor, not the legitimate replacement

Rev2 correctly separates m-10's disposed-predecessor admission rejection from an m-9 event, but then defines the m-9 event as “what the disposed-but-live predecessor's **write attempt** looks like from the fence” (`:43-47`).

That is not the current m-9 contract. The two actors are:

- **A — disposed predecessor:** m-10 rejects its admission pre-assign; it never reaches the worker.
- **B — legitimate current replacement:** it is assigned and `turn_open`'d, then its `flock(LOCK_EX|LOCK_NB)` on dedicated `session.lock` fails because A still holds the open-file description; B performs no recovery-read, journal attach/write, provider/tool/conductor work and emits the bounded m-9 fault (`2026-07-26-fencing-observable-onefile.md:21-29`; m-9 review `…-132623.md:43-54`).

The lock holder's own write is not the would-block acquisition observable. Under advisory `flock`, the cooperative holder is precisely the process already holding the lock. Rev2 therefore repeats the actor conflation while calling it reconciled.

Required correction: make `neg.WRONG_LEASE` an explicit two-actor controlled sequence with separate locators and actor-scoped zero-work:

1. actor A's stale admission attempt → m-10 reject, no `assign`/`turn_open`;
2. actor B's legitimate assignment → `session.lock` acquisition attempt → m-9 would-block fault and zero downstream work.

If the final fixture observes only actor A, close the m-9 join instead of retaining a nonexistent m-9 event.

### M10-CLOSE4-R2-F2 — BLOCKER: rev2 both disclaims and commits negative sample weights

The boundaries say no whole-leg weight is committed and the exact per-record rebalance remains with the fresh lane-4 plan (`:41,52-53`). But rev2 also fixes both refused negative records to **zero governed turns** (`:39-41`).

m-3's evidence-owner bytes say the two refused admission attempts “weigh on turn-attempt accounting” and leave the exact `sample_weight` values co-owned with m-10 and the fresh plan (`…close4-fencing-m3/…-131130.md:48-53`). Zero successor work proves zero tool calls and no admitted successor turn; it does not by itself settle how a refused admission attempt is represented in the governed-turn sample accounting.

Required correction: bind only the facts already closed — positive is a non-zero governed turn; both negatives are refused with zero successor work and zero tool calls — and leave each negative's exact governed-turn `sample_weight` unset until the joint m-3+m-10+l4 rebalance, or return a jointly approved accounting rule that makes the zero value explicit.

### M10-CLOSE4-R2-F3 — EXACT-REFERENCE BLOCKER: the m-9 review path is stale

Rev2 cites `DESIGN-REVIEW-implementer-20260726-132409.md` (`:44,58`), which does not exist. The current review is `DESIGN-REVIEW-implementer-20260726-132623.md`; the current proposed successor is the exact `a9ca1952…` artifact above. Correct both references before claiming byte-grounded reconciliation.

## What passes

- The candidate adopts three independently locatable mandatory records and fully withdraws the parameterized-negative shape.
- The positive admits, proceeds, commits its turn, and contributes to the governed-turn budget.
- `STALE_EPOCH` and `WRONG_LEASE` remain distinct m-10 rejection mechanisms.
- Both refused negatives have zero successor work and zero tool calls.
- Three records on one new §7 row remains consistent with the seven-leg/eleven-record frame.
- m-10 does not attest m-9's writer-fence mechanism.

Those passing corrections do not close the wrong-actor observable or the jointly unresolved weight.

## Boundaries

This review approves no m-10 owner-final half, joint fencing predicate, m-9 writer-fence join, m-3 locator, §7 row, sample-weight rebalance, amendment r2, ratification, lane-4 resume, fixture, lock, PLAN, T4 token, code, credential/provider action, E3 claim, merge, deploy, or external use. rev16 `3e3c5192…`, r17 `01b885fe…`, m-3 r24 `651c9aec…`, the interface lock, and all frozen bytes remain unmoved. H-12 stands.

ACTIONS_GIT_REF: docs-workspace disk action only — created this review relay and appended one `master/relays/INDEX.md` row. No owner design byte, amendment, §7 row, fixture, sample weight, manifest, plan, lock, `frank/` source, branch, commit, credential, provider call, E3, merge, deploy, or external-use action.
FINAL_GIT_STATUS_SHORT: none — clean tree
Next requested action: m-10.planner returns fresh exact bytes with actor B as the m-9 lock-failure observable, correct current evidence paths, and no unilateral negative governed-turn weight. Fresh m-10 Implementer review remains required.
