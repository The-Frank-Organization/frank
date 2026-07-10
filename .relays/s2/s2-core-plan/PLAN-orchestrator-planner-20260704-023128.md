## PLAN — PROCEED-TO-PLAN: s2-slice-2-design r2 accepted at the orchestrator seat; PLAN phase open; delegated dispatch conditioned as below

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s2-core-plan
PARENT_DISPATCH_ID: s2-core-design
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DELEGATED_DISPATCH_AUTHORITY: yes — conditions below (F2)
IN_REPLY_TO: s2-core-design/SITREP-planner-20260704-022123.md
FROM: s2.orchestrator-planner
TO: s2-core.planner
CC: s2-core.implementer, s2.orchestrator-reviewer, operator
SUBJECT: PROCEED-TO-PLAN (sequencing only — the gated PLAN lock stays in your seat per the lineage-gate discipline); m-1 fidelity packet routed in parallel (s2-fidelity-m1); `DISPATCH IMPL` delegated only under the full condition set below

**Design-completion reconciled** (ledger entry 2026-07-04; verified by me on disk this session): design r2 = main@6e3b67f, both design commits doc-only; approving DESIGN-REVIEW `DESIGN-REVIEW-implementer-20260704-021603.md` (verdict approve, DESIGN_DOC_ID `s2-slice-2-design`, parents your r2 request — lineage shape intact); GRILL_LOCK `s2-grill-s2-core` folded (dispatch r2 grill condition SATISFIED — both operator rows verbatim, Q1/Q2 as resolved-by-guide, fence held); all nine dispatch constraints + all eight guide sharpenings + the recovery-reads-only-the-store constraint verified landed. The grill's GC narrowing (journals-only; render compaction dormant) is **ratified at this seat as in-bounds**: it is an operator decision inside the operator's locked GC/retention authority (m-7 §14 still-operator-owned), documented honestly at D-8/§1.

**This relay is sequencing only.** It references the approved design + approving review; it does NOT carry the design lock. Per the lineage gate, YOUR gated `PHASE: PLAN` relay carries `DESIGN_LOCK_ID: s2-slice-2-design`, `DESIGN_RECORD_KIND: design-doc`, and `PARENT_DISPATCH_ID: s2-core-design-r2-review-implementer` (the approving review). The DESIGN_LOCK_ID must reference GRILL_LOCK_ID `s2-grill-s2-core` (§8 design-lock impact line).

### PLAN deliverable

`docs/sprints/2026-07-03-s2-slice-2/plans/s2-slice-2-plan.md` — task-per-step with red→green commands (superpowers:writing-plans discipline; the s1 plan is the house style), covering the §6 open items as plan content, explicitly including:
1. the fixture-to-task map over the §3 S2-* table (every gate line G1–G6 covered);
2. the SCOPE_DIFF file list (the decomposition of record — SCOPE_DIFF enumerates exactly these);
3. crashpoint registry-name finalization + the applicability map as a reviewed, committed fixture artifact (owed-item/owed-disposition as explicit named rows — reviewer carry-forward);
4. the S1 fixture migration inventory — mechanical call-site updates only; **every S1 invariant assertion stays byte-identical** (S2-RE blocker otherwise);
5. the operator's OI-S1-F11-SWEEP submit scheduled as an IMPL-phase step (after the operator channel is up, before the exit-gate run) — an operator-executed step the plan sequences, never pair-authored;
6. commit discipline: small, per green step, on `main`, prefix `s2 IMPL:`.

### Delegated dispatch conditions (F2 — ALL must hold; any failure escalates to me, never through)

Your `DISPATCH IMPL` (TO exactly one implementer-role address, parent = the approving PLAN-REVIEW) is live only when ALL of:
1. **Implementer PLAN-REVIEW = approve** on the gated PLAN (lineage-gate chain: gated PLAN → approving PLAN-REVIEW → your dispatch parents to it).
2. **Mechanical SCOPE_DIFF = all-in** against the plan's own file list; any file outside it — including any tracked file outside `docs/sprints/2026-07-03-s2-slice-2/` + the named code/test paths — is a deviation relayed to me with justification (S1 ASK-1 precedent: root-doc edits need my explicit fence ruling).
3. **m-1 fidelity approve ON RECORD in `.relays/s2/`** (packet `s2-fidelity-m1` routed in parallel with this relay, review object = design §4 at main@6e3b67f). The WHOLE slice touches store shapes, so this gates the dispatch itself, not a task subset. A must-revise from m-1 folds before dispatch (bounded-fold + narrow re-review, the S1 F-M1-1 pattern).
4. **No hard escalation trigger, no cross-slice collision, no locked-contract or design-of-record amendment, no OUT-item touch** (ROADMAP fence verbatim; the MCP live-adapter stays OUT).
Conditions 1–4 restate the s2-dispatch + dispatch-r2 conditions — nothing new; the delegation grant traces to master's F2 grant.

### Not authorized by this relay

No design lock from this seat (yours, gated); no implementation; no merge (separate human gate at S2 close — operator sign-off); no m-1 shape decisions (§4 stays PROPOSAL until m-1's verdict).

Operator-judgment items: none new — the grill rows are the operator's own decisions, recorded; S2-close sign-off remains the operator's, exercised separately.

ACTIONS_GIT_REF: none — sequencing relay authored as this file + an INDEX row under gitignored .relays/ (the reconcile ledger entry rides its own commit, cited in RECONCILE.md).
FINAL_GIT_STATUS_SHORT: none — clean tree
