## RECONCILE — THE JOINED B10 RULING CARRIED DOWN (m-9 rules, m-8 concurs, zero conflict): the byte-fidelity observable at the persisted journal is DECODE-ROUND-TRIP BYTE-EQUALITY through the CLOSED two-branch item carrier; B10 lands LAST as the bounded battery-predicate fold + the W fix; the hold is LIFTED on these exact pins

**The ruling of record** is m-9's `master/relays2/s16a-wp2-b10/SITREP-planner-20260825-063617.md`, adopted verbatim; m-8's K6 concur `…/SITREP-planner-20260825-063610.md` CONFIRMS with no decline (K6 = byte-FIDELITY, never raw-byte literalism; m-8's own r12 `:46`/`:69` already carries opaque provider items as base64 — the endorsed carrier is the established shape). Until m-9's WP2-close design registration folds the carrier into the worker design's journal surface, the ruling relay IS the carrier's definition of record (the effective-time rule); the owed registration is tracked at `master/RESIDUALS.md` row R-S16A-B10-CARRIER-REG, joined with the A-2 §4 A14-recipe registration at the same gate.

**The pins that bind your B10 fold (m-9 §1-§2, restated exactly):**
1. **The carrier (CLOSED):** the journal item record carries EXACTLY ONE of — (a) **verbatim**: the item's raw bytes as-is, permitted IFF `bytes == JCS(decode(bytes))`; (b) **`raw_b64`**: RFC 4648 §4 STANDARD base64, padded, no line-wrap, for everything else (non-UTF-8, non-JSON, AND valid-but-non-canonical JSON). No third encoding; both members, neither member, or any other encoding = a malformed record, fail-closed at the checksum-reader like any non-decoding record.
2. **The predicate (LOSSLESS, at the worker's persisted record):** live `Runner.Run` → read persisted `session.log` → parse under the CANONICAL decode (recovery's own decode) → extract the carrier member → decode → `bytes.Equal(original)`. NEVER `bytes.Contains` on the file's raw bytes.
3. **The two legs (the anti-vacuity control on the discriminator):** (i) the needle `{0xff,0xfe,0x00,'x'}` ⇒ the record carries `raw_b64` AND NOT the verbatim member ⇒ decode == the needle; (ii) an exact-JCS item ⇒ persisted verbatim AND NOT `raw_b64` ⇒ bytes == original. Both red today; both greens reachable; neither satisfiable by weakening the other branch.

**Execution:** B10 lands LAST per the standing order — the bounded battery-predicate fold + the W (worker-side carrier) fix, under a bounded plan successor if your cadence requires one; every byte stays inside the standing fences (branch + worktree, no push, no merge). The raw-bytes-in-`session.log` alternative is DECLINED by the owner — do not implement it in any form. Any implementation question the pins do not answer routes UP, not sideways.

**For the record:** the escalation chain was exemplary end-to-end — the implementer HELD without weakening, fabricating, or amending a foreign format; the planner routed with the defect ownership honestly shared; the owner banked the missed representability check as their own. The remaining WP2 rows proceed on the standing r14 token unchanged.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: s16a-wp2-b10-ruling
PARENT_DISPATCH_ID: s16a-wp2-b10-hold
RUN_ID: s16a
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the routed owner ruling carried down; no operator choice is opened; the operator's next gate remains the terminal WP5 MERGE-GATE
IN_REPLY_TO: s16a-wp2/SITREP-planner-20260825-062725.md
FROM: master.orchestrator-planner
TO: s16a.planner, s16a.implementer
CC: master.orchestrator-reviewer, operator, m-9.planner, m-8.planner, m-9.implementer
SUBJECT: B10 hold LIFTED on the joined ruling — closed two-branch carrier (verbatim-iff-exact-JCS | raw_b64 RFC4648§4 padded no-wrap) + canonical-decode round-trip predicate with the two-leg anti-vacuity cut; raw-bytes declined; B10 lands last; registration owed at WP2-close rides R-S16A-B10-CARRIER-REG

ACTIONS_GIT_REF: engine-lane governance act — this carriage drafted at .engine/drafts/master.orchestrator-planner/ and submitted through relay submit; one registry row appended to master/RESIDUALS.md this act (banked by the accompanying checkpoint); no source, test, branch, or commit byte from this seat.
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit:)
 M master/RESIDUALS.md
 M master/relays/CHECKPOINTS.md
 M master/relays2/INDEX.md
?? master/relays2/s16a-wp2-b10/SITREP-planner-20260825-063610.md
?? master/relays2/s16a-wp2-b10/SITREP-planner-20260825-063617.md
