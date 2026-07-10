## DESIGN — s2-core design dispatch r2 (supersedes `…-004400`; VP F1 folded: GRILL_REQUIRED flipped to yes + pre-lock GRILL_LOCK requirement)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: s2-core-design
PARENT_DISPATCH_ID: s2-core-audit
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: yes
IN_REPLY_TO: s2-core-design/RECONCILE-orchestrator-reviewer-20260704-004823.md
FROM: s2.orchestrator-planner
TO: s2-core.planner
CC: s2-core.implementer, s2.orchestrator-reviewer, operator
SUBJECT: DESIGN dispatch r2 — same nine hard constraints; now GRILL_REQUIRED: yes (VP F1, verified correct): run design-grill with the operator before lock; GRILL_LOCK_ID folds into DESIGN_LOCK_ID; drafting may proceed provisionally meanwhile

**Supersedes `s2-core-design/DESIGN-orchestrator-planner-20260704-004400.md`.** One fold: the VP's blocking finding F1 (`RECONCILE-orchestrator-reviewer-20260704-004823.md`) — verified correct by me against the design-grill trigger rules before folding: this design is new-feature/still-open at medium tier (trigger met on its own), AND carries a cross-domain boundary contract (the m-1 store-touch surface), hard-to-reverse on-disk data decisions (owed/disposition record shapes, genesis record, quarantine/ + journal-segment layout), and two unsettled questions with multiple downstream choices hanging on them (Q1/Q2). `GRILL_REQUIRED: no` was wrong; it is now **yes**. Everything else in the r1 dispatch — scope, the nine hard constraints, provisional-Q1/Q2 handling, OUT list, the lineage-gate discipline, acceptance criteria — carries forward **verbatim by reference**; re-read r1 for them (it stays on disk as the constraint text of record; this r2 changes ONLY the grill posture).

### The grill requirement (the r2 delta — binding)

1. **Run the `design-grill` step with the operator before any design lock.** One question at a time, recommended answer per question, codebase-answerable questions answered from the code (not asked). The product is a durable **GRILL_LOCK artifact** (GRILL_LOCK_ID + sources + resolved decisions + rejected alternatives + still-operator-owned + design-lock impact) folded into the design record.
2. **No `DESIGN_LOCK_ID`, no design-review approve consumed toward PLAN, no PROCEED-TO-PLAN request, until the GRILL_LOCK exists** and its decisions (plus any guide-answer deltas from s2-guide-q1) are folded. Drafting the design doc meanwhile is fine and encouraged — provisional work only, exactly as the r1 provisional-Q1/Q2 rule already framed it.
3. **Grill agenda floor (must resolve or explicitly carry each):** Q1 genesis/config-digest scope; Q2 owed-record authorship/provenance path; the m-1 fidelity proposal boundaries (owed-item `record_kind` + disposition shape — what the pair proposes vs what only m-1 may fix); the on-disk layout commitments (quarantine/, journal segmentation/rotation). If the guide's s2-guide-q1 answers land before or during the grill, fold them as resolved-by-guide rows rather than re-asking the operator.
4. **Grill fence — do NOT re-open locked decisions.** The commit pivot, recovery phase set, fault taxonomy, GC retain-everything posture, store verbs, and enum are c1/c4 operator-grilled locks (m-1/m-7 GRILL_LOCKs); the grill stress-tests S2's *new* surface against them, it never re-litigates them. A grill answer that would amend locked text is an escalation to master via me, not a grill outcome.

### Everything else: unchanged from r1 (by reference, binding as written there)

Scope + deliverable (`designs/s2-slice-2-design.md`) · the nine hard constraints (promote-don't-rebuild; reified phase machine + boundary crashpoints; quarantine-replaces-fail-stop incl. live-path sweep; single intake-writer; genesis/GC claim-pinned; projection generalizes derived.go, record_kind = proposal to m-1; claim honesty; harness extension; performance-shape findings) · provisional-pending-guide handling · OUT list · the design-review lineage discipline (your design-review request goes TO s2-core.implementer; approve parents your DESIGN relay on the same DESIGN_DOC_ID) · acceptance criteria — all per `DESIGN-orchestrator-planner-20260704-004400.md`, which this r2 supersedes as the dispatch of record while incorporating its body by reference.

Operator-judgment items: the grill itself is operator-facing by design (they are the grillee); Q2 already carries operator visibility. No new judgment items beyond the standing set (D5 restated; MCP deferral stands).

ACTIONS_GIT_REF: none — superseding dispatch authored as this relay file + an INDEX row under gitignored .relays/; no tracked-file edit in this action (the ledger entry rides its own commit, cited in RECONCILE.md).
FINAL_GIT_STATUS_SHORT: none — clean tree
