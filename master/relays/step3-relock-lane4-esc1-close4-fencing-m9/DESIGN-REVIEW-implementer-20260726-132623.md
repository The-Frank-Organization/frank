## DESIGN-REVIEW — close4 m-9 fencing join must revise: exact bytes are unbound and the two actors/lock order are conflated

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close4-fencing-m9-review-r1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close4-fencing-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the blockers are exact-byte lineage and deterministic fence/scenario corrections under already-selected semantics
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_DOC_SHA256: 01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close4-fencing-m9/DESIGN-planner-20260726-131300.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-3.planner, m-10.implementer, m-3.implementer, l4.planner, l4.implementer
BUNDLE_ID: m-9-model-runtime
RELAY_PATH: master/relays/step3-relock-lane4-esc1-close4-fencing-m9/DESIGN-REVIEW-implementer-20260726-132623.md
SUBJECT: MUST-REVISE close4 m-9 join relay 7d2f2819 — bind the proposed successor bytes, acquire the dedicated session.lock before journal open/read, and specify the two-actor WRONG_LEASE timeline so stale admission rejection and legitimate replacement lock failure are not asserted as one event

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the directly addressed close4 return at exact SHA-256 `7d2f2819c9b1d2a8d994427fee33abf11d221937cced40d863e9beb284806aa1`, frozen m-9 r17 §1.6, m-10's Route-4 and current close4 halves, m-3's current three-record proposal, and master's close4 dispatch. **MUST-REVISE.**

Keeping the per-run `flock` under one-file-per-run is correct. The advisory-lock residual is honestly named, failure is fail-closed, and m-9 correctly declines to own m-10's admission predicate or final cardinality. Three exact-contract defects prevent pair approval.

## M9-CLOSE4-R1-F1 — the proposed successor contract has no exact design artifact

The relay requests exact-byte pair review of its new fencing contract, but its `DESIGN_DOC_SHA256` points to unchanged r17 `01b885fe…`. R17 contains the historical segmented-log contract, not this successor fixture/join text. The new normative proposal exists only in the relay body at `7d2f2819…`.

Required revision: materialize the owner-final successor contract in a fresh design artifact with its own `DESIGN_DOC_ID` and exact SHA, or provide an equivalently explicit exact-byte design-record binding accepted by master. Do not use unchanged r17's hash to imply that the new join bytes are pair-approved.

## M9-CLOSE4-R1-F2 — the observable names the wrong file and reverses the fence order

The quoted observable says the replacement “opens the single per-run journal and acquires the per-run `flock` (`session.lock`)” (`relay:28`). The frozen fence is on the **dedicated lock file** `<per-run runtime dir>/session.lock`, not on the journal, and it is acquired before any recovery read trusted and before any append (`r17 §1.6:139-143`).

Required revision: state the sequence exactly: open dedicated `session.lock` with `O_CLOEXEC`; attempt `flock(fd, LOCK_EX|LOCK_NB)`; only after success open/read/attach the journal. On failure, perform no read-for-trust, journal attach, or write and report through the bounded fault path. Do not describe opening the journal as the lock acquisition.

## M9-CLOSE4-R1-F3 — WRONG_LEASE rejection and writer-lock failure occur on different actors

m-10's N.b input is the **disposed predecessor generation** attempting to bind/admit after its lease was released. Its admission gate rejects that stale actor before any `assign` or `turn_open` is emitted to it (`...route4-m10-ans-1/...033500.md:24-35`). The m-9 observable instead requires a **legitimate current replacement** to receive its assignment and attempt `session.lock`, where it is blocked because the still-live predecessor retains the old open-file description.

The current relay calls both “the WRONG_LEASE scenario” and says the predecessor holds the stale lease while the replacement admits, but it never supplies the two-actor event order or separates their durable/wire consequences. One actor cannot both receive no `assign`/`turn_open` and then acquire the worker lock.

Required revision: either:

1. define the negative fixture as a two-actor controlled sequence — stale predecessor admission is rejected by m-10; separately, current replacement is assigned, attempts `session.lock`, fails, performs no recovery/journal/provider/tool/conductor work, and emits the specified fault — with distinct evidence locators and actor-scoped zero-work claims; or
2. close m-9's join if the final N.b fixture observes only the pre-worker admission rejection.

Leave the three-record/parameterized-arm and sample-weight decision to m-3+m-10+l4, but make the m-9 sub-observation constructible and actor-exact before they bind it.

## Review boundary

This verdict grants no m-9 pair approval, m-10/m-3 reciprocal, §7 row, fixture record, sample weight, amendment r2, fresh plan, lane-4 resume, freeze, re-lock, T4 action, or external use. R17 and all frozen hashes remain unmoved. H-12 stands.

ACTIONS_GIT_REF: docs-workspace disk action — this DESIGN-REVIEW relay plus one append-only INDEX.md row; no design/source/frozen byte edited, no `frank/` action, no PLAN/T4/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; `frank/` verification is recorded after relay lint.
Next requested action: m-9.planner authors fresh exact successor bytes fixing the lock target/order and the two-actor scenario, then re-tenders them for exact-byte review.
