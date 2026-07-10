## COORD — s1 → m-7 guide: early advisory read of the pair-approved S1 design (ahead of the formal plan gate)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s1-guide-design-read
PARENT_DISPATCH_ID: s1-core-design
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this is advisory; your formal gate (with the master VP) remains at the S1 PLAN
FROM: s1.orchestrator-planner
TO: m-7.planner
CC: master.orchestrator-reviewer, s1.orchestrator-reviewer, s1-core.planner, s1-core.implementer, operator
IN_REPLY_TO: ../.relays/s1/s1-guide-q1/SITREP-planner-20260703-141628.md
SUBJECT: the S1 design is pair-approved (r2) — operator suggests you see it now, in parallel with PLAN drafting, so domain findings fold in before the plan gate instead of at it

**Request.** An early advisory read of the locked S1 design — not a gate, no verdict owed;
`must-fix-before-plan` findings (if any) come back through this lane and fold into the PLAN
the pair is drafting now. Your formal co-gate with the VP still happens on the PLAN itself.

**The artifact:** docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md
— DESIGN_DOC_ID `s1-slice-1-design`, r2, committed frank main@5622516. Pair lineage:
Implementer `DESIGN_REVIEW_VERDICT: approve`
(.relays/s1/s1-core-design-r2-review/DESIGN-REVIEW-implementer-20260703-152445.md),
after an r1 must-revise whose five blockers were folded + fold-verified (§8 fold-log in the doc).

**Where your domain eyes matter most (the r2 deltas + the choices that touch m-7 locked text):**
1. **C7 / crash-window fold (r1 blocker-1):** the accepted gate record is the durable
   derived-work intent; missing park/outbox follow-ups complete before channels open
   (an extension of your R-3 rebuild-before-open to derived-work completion), deduped by
   `gate_record_ref`; park and outbox stay distinct mutation classes (no compound pivot).
   Does this read as faithful to §4/§6/F11, and is the R-3 extension the right shape?
2. **Your six constraints as landed:** R-1 checksummed-from-first-commit, R-2 outcome→`intake_id`,
   R-3 rebuild-before-open, O-1..O-3 (outbox loop-mutation / produce-only claim-honest / open
   ODB envelope, no `model_name`); F9 named whole. Spot-check the doc's D-4/D-5/D-7/D-9.
3. **B1 claim narrowing (r1 blocker-4):** B1 = the no-tool-path property; direct record-file
   injection detection/quarantine explicitly S2; direct-file leg limited to torn/staging
   cleanup; D5 residual stated plainly. Consistent with the §16 claim boundary?
4. **MVP FieldSpec gate enums:** the full frozen §J2 `gate_category` default set locked
   byte-exact in DESIGN (`routing_escalation` preserved as owed non-member per
   ARCHITECTURE :481). Any m-7-side concern with carrying the full set this early?
5. **Blocker-2 grant narrowing (orchestrator-ratified, conditions bound to PLAN, operator veto
   open):** S1 renders `grant` on operator/orchestrator forms only (frozen m-2 :177 verbatim);
   pair-Planner conditional render + the m-2 :167 walk land S3. Flag now if you see an m-7-side
   reason this forecloses anything.
6. **Stack:** conductor = Go (operator decision D-1) with the MCP Go SDK per-seat-socket
   capability check as first implementation task, fallback invariants named (per-seat channels,
   stamped identity, registry exactly the 3 verbs). Any conductor-core concern with the stack
   or the fallback shape?

**Context for cross-reading (all in frank/, read-only for you):** the reconciliation ledger
docs/sprints/2026-07-03-s1-slice-1/RECONCILE.md (entries 2–4: audit reconcile, your q1 answers
folded, design reconcile + ratification); the PROCEED-TO-PLAN dispatch
(.relays/s1/s1-core-plan/PLAN-orchestrator-planner-20260703-153356.md) — its delegation
conditions hard-wire your + the VP's plan-gate approves and the m-1/m-2 fidelity approves
before any `DISPATCH IMPL`.

Not requested / not granted: no verdict obligation, no lock reopen, no scope change; the pair
does not pause PLAN drafting for this read unless you return a must-fix-before-plan finding.

ACTIONS_GIT_REF: wrote this relay + an INDEX.md row under .relays/s1/ (gitignored substrate); no tracked-file change (tree clean at main@876e279)
FINAL_GIT_STATUS_SHORT: none — clean tree
Next requested action: operator carries this to the m-7.planner session; guide reads at own cadence; findings (if any) return through this lane and fold into the PLAN before or at the formal gate.
