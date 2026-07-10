## SITREP — lint-residue DISPOSITIONED: operator granted a scoped waiver for the superseded r2 PLAN file; the lint gate is satisfied on record; dispatch now waits on m-1 alone

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s2-core-plan
PARENT_DISPATCH_ID: s2-core-design
RUN_ID: s2
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the flagged gate was exercised (operator waiver below)
IN_REPLY_TO: s2-core-plan/SITREP-planner-20260704-033900.md
FROM: s2.orchestrator-planner
TO: s2-core.planner
CC: operator, s2-core.implementer, s2.orchestrator-reviewer
SUBJECT: Operator waiver on record for the superseded r2 gated PLAN's root-mode lint error — scoped to that one file; r3 chain is the chain of record; remaining dispatch gates = m-1 fidelity approve + your SCOPE_DIFF

**Operator decision (2026-07-04, in response to the disposition question put to them with the three options + my recommendation): the operator selected "Scoped waiver"** — the structural root-mode lint error on the superseded `s2-core-plan/PLAN-planner-20260704-030751.md` (design-lock parent edge, the defect r3 fixed) is **waived by the operator for that one named file**. Recorded verbatim-equivalent per the S1 pattern (s1-exit-gate lint-flag disposition (a)): selection made from options {scoped waiver — recommended; current-chain ruling; full-root-clean}, so the waiver is informed and scoped, not a standing gate-loosening.

**Basis on record:**
- The r2 file is superseded; nothing consumes it; the live chain runs r3-only (`s2-core-plan-lock-r3` → re-affirming approve `s2-core-plan-review-implementer-r3` → the approving DESIGN-REVIEW — full lineage-gate walk resolves; both r3 files exact-file AND root-mode OK, pair-verified + planner-rerun E2).
- Append-only trail discipline forbids rewriting/removing the superseded relay; a scoped waiver is the protocol-named bypass for a structural lint error ("unless the operator explicitly waives").
- Precedent: the S1 close used exactly this shape (operator waiver line for a named file's root-mode flag; exact-file lint OK; classified trail-cleanliness, not code evidence).
- **Scope limit:** this waiver covers ONLY `PLAN-planner-20260704-030751.md`'s design-lock-parent error. Any NEW root-mode error on any live relay gates as normal; future supersession residues need their own disposition (no standing rule was created — the operator chose waiver over the current-chain ruling).

**Standing state:** lint gate SATISFIED on record. Remaining before your delegated dispatch: (1) the m-1 fidelity approve in `.relays/s2/s2-fidelity-m1/` (substantive gate, pending); (2) your mechanical SCOPE_DIFF over the plan's file list (README row citing the fence ruling `…-032200.md`) with all-in. Then `DISPATCH IMPL` parenting `s2-core-plan-review-implementer-r3` under the standing conditions. No dispatch token in this relay.

ACTIONS_GIT_REF: none — report-only disposition record; this relay + its INDEX row under gitignored .relays/ (the ledger entry rides its own commit, cited in RECONCILE.md).
FINAL_GIT_STATUS_SHORT: none — clean tree
