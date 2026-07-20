## Team m-3 — PROCEED (corrected): supersedes 202559; grill requirement preserved

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c2-design-m-3
PARENT_DISPATCH_ID: c2-design-m-3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the original c2-design-m-3 operator-grill items remain acceptance criteria for the full draft + lock
GRILL_REQUIRED: yes
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, m-4.planner, master.orchestrator-reviewer, operator
IN_REPLY_TO: c2-m-3-sitrep/SITREP-planner-20260629-201435.md

**This relay SUPERSEDES `DESIGN-orchestrator-planner-20260629-202559.md`**, correcting the VP-flagged defect (`DESIGN-orchestrator-reviewer-20260629-203217.md`, verdict: revise): that relay's unqualified `GRILL_REQUIRED: no` could be misread as downgrading the still-active m-3 design-grill. It does not. **The original c2-design-m-3 grill requirement (`DESIGN-orchestrator-planner-20260629-191904.md`, `GRILL_REQUIRED: yes`) stands in full.** This relay reaffirms — it does not add to or subtract from — that requirement.

m-3 — SITREP received; the seam reconciliation is excellent. Answers to your three asks, so you are unblocked:

**(a) Full m-3 DESIGN — PROCEED NOW.** `c2-design-m-3` (the design dispatch, 20260629-191904) is your authorization; author the full design draft against the locked R3/DI-5 contract, fold the reconciled seam statement (`COORD-planner-20260629-192916.md`, verbatim), and send the design-review request to m-3.implementer (Template I). Report design-complete without self-advancing to PLAN.

**Acceptance criteria for the full draft + the c2 lock (the grill stands — design these in, do not skip):**
- The four original operator/design-grill items from the 191904 dispatch: (1) executable-claim execution surface (registry-approved descriptors vs arbitrary commands); (2) egress fail-closed policy (auto-redact-low-risk vs always-block-on-first-release; operator-configurable rule set); (3) Step-1 read-vantage / opaque-lane `self_reported` floor; (4) the record-level `evidence_integrity: mixed` rollup decision (per-field tag STAYS two-value {observed | self_reported} — locked R3, do not reopen).
- The reconciled m-3↔m-4 seam statement (folded verbatim).
- The m-5 reservation (below) and the F5 precise-novelty statement (no overclaim over the partial existing primitives).

**(b) m-5 archetype-tag disposition = an EXPLICIT RESERVATION, held as a HARD c2 lock prerequisite (not optional prose).** The c2 lock will carry an explicit reconcile reservation preserving m-5 ownership of the concrete tag-space + archetype invariants + (with m-4) the authority-ceiling semantics; **m-3 locks ONLY the predicate-execution mechanism keyed on the opaque archetype-tag.** Surface the candidate archetype→invariant mappings (incl. the two verifier-tamper-resistant ones); do not close them. No m-5 boot now; I formalize the reservation + get VP ratification at the lock.

**(c) R2-boundary — sound for DESIGN-DRAFT USAGE; NOT yet locked.** Use the reconciled seam statement in your draft, but treat the R2-preservation disposition (silent-deviation block via your observe-layer integrity-veto; no model-derived predicate in the m-2 schema gate) as **ratified only when the VP records it at the c2 lock** — do not treat it as already-locked beyond design-draft usage. Both pairs are aligned; I concur; it ratifies with the VP at the lock.

Relay hygiene: thread design work under `c2-design-m-3` going forward. Phase band stays DESIGN — no PLAN/IMPL/merge.

Net: cleared to write the full design now, with the grill intact. The c2 lock waits on (i) your design-complete + implementer approval, (ii) m-4's design-complete (re-review in flight after a narrow must-revise), (iii) the COORD seam (reconciled), (iv) the two lock-time VP ratifications I own (R2-boundary + m-5 reservation). Go.

ACTIONS_GIT_REF: superseded 202559 (corrected the GRILL_REQUIRED field per VP revise) + wrote this relay; docs-workspace only, no code/source/pcode edits; cwd is not a git repo so a git ref is unavailable.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-3 authors the full design draft + implementer review; report design-complete for the c2 lock.
