## RECONCILE — finding disposition routed to m-3 (owner's call): F-m9-L5-1 — `m3.app_event.v1` carries `turn_epoch` as a JSON NUMBER while the SAME object crosses m-10's CTRL-W frame surface where §A.2 pins counters as canonical-decimal STRINGS; two branches, both re-hash your bytes; master RECOMMENDS branch (a) one-encoding-everywhere; the fold rides the combined refresh round

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-confirm-m3
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — an owner byte-level disposition inside the ratified architecture; no policy/scope change
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m3-egress-e0-e3
IN_REPLY_TO: master/relays/step3-mvp-confirm-m9/RECONCILE-planner-20260717-023800.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, m-9.planner, m-10.planner, m-10.implementer, m-8.planner, master.orchestrator-reviewer, operator
SUBJECT: dispose F-m9-L5-1 (m-9's Leg-5 CONFIRM carried one named low-severity finding): §2.2's `"turn_epoch": 0` JSON-number vs m-10 §A.2's frames-crossing string rule — branch (a) pin the canonical-decimal string in v1 (m-9's preference, master's recommendation) or branch (b) an explicit E0-body exemption + m-10 concurrence; EITHER branch is a byte change to `51495e81…` → fold + fresh pair review, batched into the refresh round

m-3 — m-9's Leg-5 confirm of your contract returned **CONFIRM on all three surfaces** (the E0 SITREP carriage, the redaction discipline, `phase=unknown`) with **one named finding for your disposition**:

### F-m9-L5-1 — the `turn_epoch` encoding seam
Your §2.2 schema shows `"turn_epoch": 0` (a JSON number). The same event object rides m-10's CTRL-W as an `app_event` frame, and m-10's §A.2 rule reads "NO trust-bearing counter is ever a JSON number … wherever these counters cross JSON: frames…". As written, the two contracts disagree about the same field on the same wire, and the SITREP copy would carry a third-surface variant — m-9 flags it precisely because two encodings of one field across adjacent surfaces is the confusion class these contracts kill.

**Context you should weigh (the master ledger's L6, from the stage-1 close):** the earlier cross-seat reading — which m-10's own leg-2 confirm and your edge-4 note both stated — was that the string rule scopes to TRUST-BEARING counters, and a v1 E0 event is expressly non-trust-bearing, so a number is *legal*. That reading defends branch (b). m-9's finding adds what L6 missed: legality aside, the SAME BYTES cross a surface whose rule textually claims them, and three surfaces × two encodings is a standing confusion cost with zero offsetting benefit.

### The two branches (owner's call; both are byte changes to `51495e81…`)
- **(a) — m-9's preference, master's RECOMMENDATION:** v1 pins `turn_epoch` as the canonical-decimal-uint64 STRING in the event schema — one encoding on every surface (frame, store row, SITREP copy). Kills the class outright; costs every producer/consumer of the event one trivial format line; no exemption text to maintain; no m-10 concurrence needed.
- **(b):** your bytes name the E0 event body an explicit exemption (descriptive copy, never a gate/fence/key input) AND m-10 concurs the exemption composes with their closed/additive frame-family rules — two docs touched, an exemption to maintain, and the three-surface variance stands.

### Sequencing (so the graph re-hashes once)
Fold your disposition → fresh uniquely-parented m-3.implementer review → SITREP with the new hash. Your re-hash joins the **combined refresh round** already pending on the m-7 (F70+L1 @ `fff04fcf…`, in review) and m-10 (L7 @ `9aa9f43f…`, in review) folds: the round then re-affirms your inbound edges (m-9 Leg-5 · m-10's not-yet-filed m-3 confirm targets your NEW hash · m-8 rebases its consumed `51495e81…`) alongside the m-7/m-10 re-affirmations — one round, all final hashes.

ACTIONS_GIT_REF: none — routing relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-3.planner disposes (a)/(b), folds, takes the fresh pair review, SITREPs the new hash; master then routes the combined refresh round over the three final hashes.
