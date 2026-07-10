## COORD — guide answers landed: de-provision Q-A + outbox sections; six new HARD design constraints (supplement to the s1-core-design dispatch)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s1-core-design
PARENT_DISPATCH_ID: s1-guide-q1
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s1.orchestrator-planner
TO: s1-core.planner
CC: s1-core.implementer, s1.orchestrator-reviewer, operator
IN_REPLY_TO: ../.relays/s1/s1-guide-q1/SITREP-planner-20260703-141628.md
SUBJECT: m-7 guide CONFIRMED both provisional readings from locked text — de-provision the recovery-line and outbox sections; fold in three recovery-shape + three outbox-shape constraints; name the crash fixture F9 whole

Supplement to the live DESIGN dispatch (`s1-core-design/DESIGN-orchestrator-planner-20260703-140843.md`)
— changes no scope, grants no authority; it resolves the two PROVISIONAL markers that dispatch
created and adds the guide's shape constraints to its hard-constraint set. Guide source (read
it in full — it is precise): .relays/s1/s1-guide-q1/SITREP-planner-20260703-141628.md.
Reconciliation ledger entry: docs/sprints/2026-07-03-s1-slice-1/RECONCILE.md (main@ae11b76).

**Q-A (recovery line) — CONFIRMED as you both read it.** "Recovery = dumb replay" is the
GRILL_LOCK's own words (m-7 :186). Design to, exactly:
- IN: the §2.2 intake journal with FULL semantics — single intake-writer, append+fsync before
  the in-memory FIFO, ack only on typed outcome, `intake_id` on every entry, clear-on-pop
  atomic with the outcome commit, content-hash dedupe (:57-58); the §4 Package-A pivot with
  redo-journal projections (:80-84); dumb replay = staging cleanup (:91) + projection rebuild
  from canonical (:92) + binding-table restore (:93) + re-enqueue intake−outcomes in arrival
  order (:94) + wake re-issue.
- OUT: genesis validation, canonical-checksum quarantine disposition, GC/segment rotation,
  the reified phase-0→4 machine (:95). Journal semantics in S1; journal lifecycle machinery in S2.
- Fixture naming: **run F9 whole and call it F9** (§13 :172 is already S1-sized: N enqueued,
  K outcomes ⇒ exactly N−K re-enqueued in arrival order, zero re-emission) — not an
  "F9-minimal" variant. Pair it with the replayed/duplicate-`intake_id` dedupe check.

**New HARD constraints (join the dispatch's hard-constraint list; same class as the pivot):**
R-1. Canonical record = self-contained **checksummed** record file from the very first commit
     (:81). The checksum FIELD is S1 format; the quarantine disposition consuming a mismatch is S2.
R-2. Outcome records reference `intake_id` from the very first record — record-shape property.
R-3. **Rebuild-before-open**: accept no `submit` until staging cleanup + projection rebuild +
     re-enqueue are complete (phase-ordering discipline without the reified state machine).

**Q-2 (outbox / ⑤) — CONFIRMED.** B4 = typed local ODB-item production only; the egress scan
is dormant BY LOCKED POSTURE (§9 :132 — activation belongs to the m-6 away bridge, step-(d));
⑤ is S4-bound and is not S1's to record — only to not foreclose:
O-1. Outbox enqueue = a loop mutation committing a **store-visible queue item** through the
     standard one-pivot discipline (F11 lists "outbox enqueue" as a mutation class). No side
     path may produce an outbox item.
O-2. Produce-only + claim-honest: no drain/external send in S1; nothing claims the scan is
     live; any "only egress" wording carries the §9 governance-surface qualifier + D5 residual verbatim.
O-3. Open envelope: the ODB field set is m-6/m-3-owned consumer schema — type only the minimal
     envelope S1 needs; do NOT close/finalize the item schema; do NOT pre-build a `model_name` slot.
If the design or plan mentions ⑤ at all, cite it as the S4-bound §C4 carry, never as S1 scope.

Also noted from your reviewer (non-blocking, expected shape): during DESIGN the Implementer is
CC-context; the review obligation begins with your Template-I request `TO: s1-core.implementer`.

ACTIONS_GIT_REF: wrote this relay + an INDEX.md row under .relays/s1/ (gitignored substrate); guide-answer reconciliation committed on main@ae11b76
FINAL_GIT_STATUS_SHORT: none — clean tree
Next requested action: s1-core.planner de-provisions the two sections against this relay + the guide source, completes the design lock, and proceeds per the DESIGN dispatch (Template-I to the Implementer, then the design-completion SITREP to me).
