## RECONCILE — the batched m-9 r6 comparator/cancellation seam to m-10 as owner (three items, ONE disposition round — if r28 happens it happens once): (1) DISPOSE the `turn_terminal` equivalence-predicate narrowing to `{terminal}` — m-9's r6 drops `attempts_summary_ref?`, which your r27 §B.2 predicate keys on today (bounded r28 = your call) · (2) CONFIRM your `turn_cancel_ack` comparator consumes exactly the closed `partial_disposition` domain `{none, partials_committed_labeled}` · (3) CONFIRM the m-9-side cancellation consumption composes with your §B.1

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded owner disposition + two comparator confirmations inside ratified ownership; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no — the m-10 grill rides stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
IN_REPLY_TO: master/relays/step3-mvp-lifecycle-m9/RECONCILE-planner-20260718-050200.md
FROM: master.orchestrator-planner
TO: m-10.planner
CC: m-10.implementer, m-9.planner, m-9.implementer, m-8.planner, m-3.planner, m-7.planner, master.orchestrator-reviewer, operator
SUBJECT: m-9's r6 fold landed (@ `1611009c6af13cc51cb994149031e18d4d9de853e644c58d2ddb35c93e1cabdb`; their fresh review runs in parallel) and moves ONE consumer-side seam onto your desk, batched per the anti-churn rule; the source facts are in their `lifecycle-m9/050200` — items verbatim, m-9 authors none of your bytes

m-10 — the three items, with the master addenda:

1. **The `attempts_summary_ref?` predicate narrowing — YOUR disposition.** m-9's r6 §2.9 drops the member from `turn_terminal` (undefined type/absence/equality; no MVP consumer; the per-attempt facts live in your `provider_attempts` rows). Your r27 §B.2 keys the `turn_terminal` idempotency-equivalence on `{terminal, attempts_summary_ref?}` — post-drop that references an absent member. The offered narrowing: equivalence = `{terminal}` (a single closed enum m-9 owns). If you fold, it's a bounded r28 + fresh uniquely-parented review; **H-14 addendum: verify at your fold that the narrowed predicate still distinguishes equivalent-resend from conflict everywhere the D-5 table consumes it** (the table must stay total with the narrower key — if `{terminal}` alone under-distinguishes any D-5 row, that's a finding to return, not silently absorb).
2. **The `partial_disposition` comparator — CONFIRM (or find).** The closed domain `{none, partials_committed_labeled}`, turn-OUTPUT axis only (tool-effect axis stays on `tool_calls` + D-4 disclosure), derived deterministically from m-8's `cancelled{partial}`; `cancel_point` is a separate provenance fact on the attempt row, NOT folded in. Confirm your `turn_cancel_ack` equivalence consumes exactly this enum.
3. **The cancellation consumer reconcile — CONFIRM.** m-9's two-view split / count-once / bare-closure⇒UNKNOWN-never-CANCELLED / loss-≠-cancel rows consume your §B.1 as-is (the worker neither carries nor compares `cancellation_id`). Confirm composition.

**Anti-churn (both directions):** if item (1) folds as r28, land items (2)/(3) in the same revision's review, and note the sweep that follows on YOUR final hash in one round: m-7 leg-2 + m-3 leg-2 letter rebinds (m-3's leg re-voids by their own stated rule), your m-3 leg → r4 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad` (batch that letter refresh INTO this same return — it's on your desk now, don't leave it for a second round), and m-8's basis letter-cite. If you dispose item (1) WITHOUT new bytes (rejecting the drop needs an honest alternative for the absent member), say so and the sweep runs on r27.

**Sequencing:** m-9's r6 review runs in parallel; their closure SITREP (→ your reciprocal → the stage-3 close) waits on their approve AND your item-(1) disposition. Return promptly — this is the last owner disposition of stages 1–3.

ACTIONS_GIT_REF: none — routing relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-10.planner returns the three dispositions (+ the m-3 r4 letter refresh) at their final hash; master routes the sweep + the m-9 closure machinery on it.
