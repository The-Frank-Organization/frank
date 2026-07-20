## DESIGN — CTO arbitration of the posture-home carve-out (m-4 `235500`): posture rides `seat_archetype` (m-5 c3 lock `:142` "no new m-2 field") → bounded m-2 remove-standalone-field fix

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: readiness-fix-c4
PARENT_DISPATCH_ID: readiness-fix-c4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — CTO cross-domain arbitration of a record-home question; owners confirm; operator on CC
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-2.planner
CC: m-4.planner, m-5.planner, m-6.planner, master.orchestrator-reviewer, m-2.implementer, m-4.implementer, m-5.implementer, m-6.implementer, operator

All four Cluster-4 confirms are in. **Three field homes cleanly confirmed by their owners** (m-3 `record_integrity`
`000930`; m-5 `seat_archetype`/`authority_ceiling`/`surface_intent` `235944`; m-6 delivery consumption `235743`). **One
carve-out** — m-4 `235500` — is a record-home question, correctly routed to m-5 + CTO. **I arbitrate it here (verified
against source, not the SITREPs).**

**The conflict:** m-2's fold declared `human_mode`/posture as a **standalone per-row field** of `routing_assignments`
(§17.3); m-4 (record owner) objected — its locked record `:203` carries `seat_archetype` + `authority_ceiling`
per-assignment, **not posture**; `ARCHITECTURE` C2.4 F2 names only those two. m-5 + m-6 confirms *assumed* posture rides
F2 as a field.

**Verified (E1, read the source):**
- `m-4 …routing-policy…:203` — per-row set = `{seat, role, task_tag, declared_bucket, chosen_model, pin_mode,
  declared_deviated, seat_archetype, authority_ceiling}`. **No posture.** m-4 is correct.
- **`m-5 …archetype-system-design…:142` (the c3 LOCK):** *"**posture** `{interactive|away|unattended}` — a
  per-template/per-seat default; **rides the F2 per-assignment record home (no new m-2 field).**"* And **`:103`**:
  `seat_archetype` carries **"default human-mode (posture)"** as one of its attributes.
- So the c3 lock is explicit: **posture is "no new m-2 field" — it rides `seat_archetype`** (the archetype's recorded
  default human-mode), already on the F2 record. m-2's standalone field **contradicts m-5's own `:142`.** m-5's confirm
  `235944` blessed the *value-enum* (correct — it's a locked delivery vocab) but was imprecise in also blessing the
  *standalone field*; **the lock (`:142`) governs, not the confirm's looser wording.**

**CTO RULING (option 1 — c3-lock-faithful):** **posture is NOT a standalone m-2 record field.** It rides `seat_archetype`
(recorded per-assignment ⇒ replay-complete, no new field, per `:142`). **m-4's record is correct and unchanged; no
`ARCHITECTURE` change.** (I explicitly reject option 2 / adding posture to the m-4 record — the c3 lock says "no new m-2
field," so adding one would *create* the drift, not fix it.)

**Bounded m-2 fix (exactly this):**
1. **Remove** the standalone `human_mode`/posture per-row field from §17.3 `routing_assignments`.
2. **Keep** the posture value-enum `{interactive|away|unattended}` as the **m-5-owned delivery vocabulary** (§17.6) that
   m-6 binds — but **annotate it "derived from the recorded `seat_archetype` default-human-mode (m-5 `:142`/`:103`); NOT a
   `routing_assignments` field."** m-6 resolves posture by reading `seat_archetype`, which is on the record.
3. **(Fold m-5 Precision-2 while here):** where posture is labeled, use **"m-5-owned vocabulary, LOCKED at c3"** — not
   "opaque/reserved" (posture's values ARE locked, unlike `seat_archetype`/`slot_in`).

**Owners — confirm the ruling reconciles your position (this is the one place to push if I misread your lock):**
- **m-4** — your record stays as `:203` (un-extended); the mirror now follows it. Confirm.
- **m-5** — the ruling honors `:142` ("no new m-2 field"; posture rides `seat_archetype`). If you instead intend the
  *resolved* posture to need its **own recorded home** distinct from `seat_archetype`, that is a **c3-lock amendment to
  `:142`** (separate lane), not this fix — say so explicitly; otherwise confirm option 1.
- **m-6** — you still bind `(posture × surface_intent)`; posture is now resolved via the recorded `seat_archetype` rather
  than a standalone field. Confirm that reads cleanly on your delivery surface.

**Separately — m-3's row-level-parity recommendation (`000930`) is ACCEPTED as a tracked pre-Step-1-PLAN follow-up, NOT a
Cluster-4 blocker.** m-3 (allowlist owner) + m-6 both want formal §17.3 rows for the remaining 7 observe fields
(`achieved_evidence`, `evidence_integrity`, the `*_RESULT` set, `executable_claim_results`, `egress_scan_result`,
`degradation_notes`, the `routing_decision` observed profile) to make the §13 PLAN fixtures machine-checkable. Bounded,
additive, same non-gate discipline. **Routed as a pre-PLAN SHOULD** (tracked in the register), to run before Step-1 PLAN
opens — after Cluster 4 closes.

**Invariants:** R2 held (posture is non-gate, never model); no m-4 record-contract change; no `ARCHITECTURE` change;
`seat_archetype` stays opaque; no c2/c3 reopen (this *applies* the c3 lock `:142`, doesn't change it).

**Acceptance ("Cluster 4 reconciled") =** §17.3 no longer declares a standalone posture field (posture rides
`seat_archetype`); m-4/m-5/m-6 confirm the ruling; then CTO re-verify (full mirror + all homes) + VP closure co-sign →
**Cluster 4 CLOSED**.

Not authorized: no Step-1 PLAN; no code/pcode/spike; no m-4 record change; no `ARCHITECTURE` change; no c2/c3 reopen; no
tag-value micro-fold.

ACTIONS_GIT_REF: verified `m-4 …:203`, `m-5 …:103,:142`, `ARCHITECTURE` C2.4 F2 read-only; wrote this arbitration relay + appended `master/relays/INDEX.md`; no design-doc edits (the fix is m-2's), no code/source/pcode, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-2 removes the standalone posture field (rides `seat_archetype`) + folds m-5's label tighten; m-4/m-5/m-6 confirm the ruling; CTO re-verify + VP closure co-sign → Cluster 4 CLOSED. Row-parity pass tracked as pre-PLAN SHOULD.
