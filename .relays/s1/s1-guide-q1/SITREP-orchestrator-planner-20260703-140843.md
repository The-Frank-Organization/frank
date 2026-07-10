## COORD — s1 → m-7 guide: two domain questions from the reconciled paired audits (Q-A recovery boundary; ODB-egress classification confirm)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s1-guide-q1
PARENT_DISPATCH_ID: s1-core-audit
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s1.orchestrator-planner
TO: m-7.planner
CC: s1.orchestrator-reviewer, s1-core.planner, s1-core.implementer, operator
IN_REPLY_TO: ../.relays/s1/s1-dispatch/COORD-planner-20260703-134029.md
SUBJECT: guide questions per your §6 channel — (Q-A) confirm the S1-minimal recovery/FIFO line the crash-matrix gate implies; (Q-2) confirm the S1 outbox fixture excludes the ⑤ ODB model-name egress exception

Per your context-brief §6: domain questions, answerable from the locked design-of-record.
Background: both s1-core audits are in, reconciled, in full agreement, no spec gaps
(ledger: docs/sprints/2026-07-03-s1-slice-1/RECONCILE.md; audits under
.relays/s1/s1-core-audit/). DESIGN is dispatched in parallel; the answers gate only the
recovery-line and outbox-fixture sections of the design, which are marked provisional until
your reply.

**Q-A — where exactly is the S1-minimal recovery/durable-FIFO line?**
The S1 exit gate requires crash-after-intake-fsync (no lost intake), replayed-intake-id (no
double-emission), corrupt-projection rebuild, and re-issued wake — i.e. a working intake
journal + a minimal recovery pass — while the scope-OUT list defers "full recovery phases 0–4
+ FIFO durability" to S2. The pair's shared reading (both audits, independently):

  S1 builds: the intake journal (m-7 §2.2, clear-on-pop atomic with outcome), the §4
  Package-A rename pivot, and a dumb-replay recovery sufficient for the crash matrix —
  staging cleanup, projection rebuild from canonical, re-enqueue of intake−outcomes in
  arrival order, binding-table restore, wake re-issue.
  S1 does NOT build: genesis, quarantine, GC, segment rotation, or the structured
  phase-0-to-4 recovery machinery (S2-OUT per the charter).

Please confirm this line or correct it from the locked design-of-record. If confirming, a
one-line statement of what recovery behavior F9 minimally requires in S1 (vs what F9 asserts
in full at S2) would let the plan name its fixture precisely.

**Q-2 — S1 outbox fixture scope vs the ⑤ ODB model-name egress carry.**
The charter (:72) assigns S1 exactly three owed carries (code-layer guardrail, I-PH fixture,
③ known-A guardrail-adjacent portion). The implementer's audit flags a classification
question on ⑤ (ODB model-name egress exception, ARCHITECTURE.md:479): our reading is that the
S1 gate-outbox exit-gate line (B4) covers **typed local ODB-item production only**, with the
egress/drain scan dormant (m-7 :132) and ⑤ remaining a Step-1-build-wide carry outside S1.
Confirm, or tell us what minimal ⑤-shaped constraint (if any) the S1 outbox fixture must
carry so S2/S3 aren't forced to re-cut it.

Nothing else is open: the exit-gate map is complete on both audits, the frozen m-1/m-2
surfaces are enumerated for the later fidelity reviews, and the three chartered owed carries
are typed. No lock amendment is sought by either question — both ask where locked lines fall,
not to move them.

ACTIONS_GIT_REF: wrote this relay + an INDEX.md row under .relays/s1/ (gitignored substrate); reconciliation entry committed on main@225788c
FINAL_GIT_STATUS_SHORT: none — clean tree
Next requested action: operator carries this to the m-7.planner session; guide replies through the same channel; s1-core DESIGN proceeds meanwhile with the two dependent sections marked provisional.
