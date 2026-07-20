## RECONCILE — m-9 leg-3 REBIND to m-10 FINAL r14 `a2663a79…`: RE-AFFIRM — the `rejected_local` seam matches my `132400` forward mapping byte-for-byte; nothing else crosses CTRL-W/F59; the r3 basis takes r14 in the open review cycle

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-confirm-m9
PARENT_DISPATCH_ID: step3-mvp-confirm-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a byte-bound rebind over pair-approved final bytes; the operator gates at the Master+VP interface-lock
FROM: m-9.planner
TO: master.orchestrator-planner
CC: m-9.implementer, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
BUNDLE_ID: m-9-model-runtime
IN_REPLY_TO: step3-mvp-confirm-m9/RECONCILE-orchestrator-planner-20260717-183010
SUBJECT: leg-3 REBIND — m-10 r12 `111ab95a…` → FINAL r14 `a2663a7964fb23e5c92eeb3b3ecf530b80c1b3108813a398941724fd6a25c5b7` (r13/r12 void): the r14 `rejected_local` fifth disposition + terminal `REJECTED_LOCAL` row + the R14-F1 emission fix match my `132400` mapping byte-for-byte; the L8 cosmetics + assign + §D F59 rows are unmoved; the lifecycle-half r3 m-10 basis rebinds to r14 in the open review cycle

**Leg-3 rebind — m-10 FINAL r14 @ `a2663a7964fb23e5c92eeb3b3ecf530b80c1b3108813a398941724fd6a25c5b7` (SHA verified): RE-AFFIRM.**
Verified at the bytes against the forward mapping I pinned in `step3-mvp-design-m8/RECONCILE-planner-20260717-132400`:
1. **The `rejected_local(<reject_reason>)` fifth CTRL-C disposition + terminal `REJECTED_LOCAL` `provider_attempts` state — MATCH, byte-for-byte.** m-10 §B.1: a received `rejected_local` closes the row terminal `REJECTED_LOCAL` **on its own, with NO m-9 `attempt_stream_end` expected** — exactly my mapping (no stream-end for a no-stream reply); reject-reason tokens **m-8-owned + disjoint from m-3's policy tokens** (`denied` stays policy-only); **never `UNKNOWN`** (reserved for genuine indeterminacy); no stream, resolver never invoked, zero transport counters.
2. **The R14-F1 emission fix STRENGTHENS my accounting model, does not disturb it.** `rejected_local` is emitted immediately on the deterministic pre-freeze failure, **before the typed DATA-P return completes** — so the durable row closes before my worker sees the reply. My side is unchanged: my `attempt_open` created the row; I record a typed attempt failure on the DATA-P reject reply, emit no stream-end, count one toward §2a; the row is already terminal `REJECTED_LOCAL` m-10-side. Consistent.
3. **Nothing else crosses my CTRL-W or F59 surface.** The worker `assign{run_id, turn_epoch, manifest_digest, generation_id, broker_worker_endpoint}` is **byte-unmoved** (5 fields, post-lease-bind); the §D F59 issue/consume/outcome rows are unmoved; the L8 changes are cosmetic. Confirmed.

**Leg-3 hash chain (for the close table):** `79fcf742…` (dispatch basis) → `111ab95a…` (r12 refresh-round) → **`a2663a79…` (r14 FINAL; r13/r12 void)**.

**Owner-leg resolution:** r14 is the **owner-real confirmation of m-8's PROPOSED `rejected_local` CTRL-C disposition + terminal row state** — the last row of my lifecycle-half r3 §7 route-back table (the m-8→m-10 ask this half consumes). That row is now **RESOLVED**; the three m-9-authored-consumer deltas still open are **D-2 (m-10 attach-ready gate) · D-3 (m-7 attach taxonomy) · D-4 (m-10 parked-`UNKNOWN_TOOL_OUTCOME` gate)**.

**The r3 basis fold (in the open cycle, not a separate spin):** my lifecycle-half r3 (`d51ce074…`) is in its fresh m-9.implementer re-review right now. Per your "fold it in the open cycle with the rest," I **do not edit r3 mid-review** (that would void the in-flight review hash); the m-10 basis row rebinds `111ab95a…` → `a2663a79…` and the §7 `rejected_local` owner-leg flips to RESOLVED **in the r3 cycle's next revision** — batched with the re-review outcome (folded with findings on must-revise, or noted at closure on approve). The rebind is a basis-hash + owner-leg-status change only; my §2.2 no-stream-reject semantics already consume exactly r14's shape (via the m-8 r4 seam), so no byte meaning moves.

Duplicate/already-built gate: not applicable — a byte-bound rebind.
Boundary contract: not applicable — no artifact beyond this relay + the batched r3 basis update.

ACTIONS_GIT_REF: wrote only this relay + its INDEX.md row; no frank/ edit; no code; r3 deliberately NOT edited mid-review (the r14 basis rebinds in the next r3 revision)
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); frank/ main clean at 502e06c (s11-close)
Next requested action: master holds the leg-3 rebind for the close table; the r3 cycle carries the r14 basis + the resolved rejected_local owner-leg in its next revision; D-2/D-3/D-4 remain the open owner deltas.
