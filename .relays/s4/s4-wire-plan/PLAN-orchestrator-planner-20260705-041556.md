## PLAN — PROCEED-TO-PLAN for s4-wire (sequencing only; the gated PLAN lock stays in the pair seat; conditioned delegated dispatch)

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s4-wire-plan
PARENT_DISPATCH_ID: s4-wire-design-complete
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: satisfied for design — the operator GRILL_LOCK `s4-grill-s4-wire` is on record pre-lock (VP watchpoint W1 normalization); the s4-close and §7-authorization gates remain the operator's downstream
DELEGATED_DISPATCH_AUTHORITY: yes — conditions below
IN_REPLY_TO: s4-wire-design/SITREP-planner-20260705-042500.md
FROM: s4.orchestrator-planner
TO: s4-wire.planner
CC: s4-wire.implementer, s4.orchestrator-reviewer, operator
SUBJECT: PROCEED-TO-PLAN — design r3 accepted at this seat (approving review + grill + guide folds verified on disk + my own battery run); draft the gated PLAN; delegated dispatch conditioned on {Implementer PLAN-REVIEW approve · m-1 fidelity approve ON RECORD · SCOPE_DIFF all-in · no trigger/collision/amendment/OUT}

**Sequencing only.** This relay references the approved design and authorizes PLAN drafting;
it does NOT carry the design lock. You (the pair Planner) emit the gated `PHASE: PLAN` relay
from your own seat with `DESIGN_LOCK_ID: s4-slice-4-design`, `DESIGN_RECORD_KIND: design-doc`,
`GRILL_LOCK_ID: s4-grill-s4-wire` referenced, and `PARENT_DISPATCH_ID:
s4-wire-design-r3-review-implementer` (the approving DESIGN-REVIEW) — **keep that parent edge
across ALL plan revisions** (the S2 r2-lineage lesson; S3 held it under revision pressure).

**Verified at my seat before this relay (my own reads/runs):** revision trail r1+r2 = c28dab7,
queue-note = 059b4c5, r3 = 2ef9437 — each docs-only, the design file alone (`git show --stat`);
the lineage-gate chain resolves (r3 approve → r3 request → r2 must-revise → r2 request → my DESIGN
dispatch); GRILL_LOCK §12 carries the six operator rows verbatim with the dispatch grill-rule
honored (guide answers entered as resolved rows, no c1–c6 reopen, zero amendments needed);
the m-7 guide relay confirms all six from locked text; battery at 2ef9437 = 20 packages ok
(uncached) + vet clean, my run.

### Conditions for your delegated `DISPATCH IMPL` (all must hold; any failure escalates to me)
1. Implementer PLAN-REVIEW approve, parenting your gated PLAN (which parents the approving
   DESIGN-REVIEW).
2. **m-1 fidelity approve ON RECORD in `.relays/s4/s4-fidelity-m1/` before ANY dispatch** —
   the packet is routed in parallel (`s4-fidelity-m1/SITREP-orchestrator-planner-20260705-041556.md`);
   five items incl. the `config_change` record_kind + provenance (ruling condition 4) and the
   `config/`-as-derived-projection store-shape posture. A must-revise folds bounded (the S2/S3
   F-M1 pattern). If m-1 prescribes shapes, the plan carries them VERBATIM or routes back.
3. Mechanical SCOPE_DIFF all-in vs the plan's file list. Root-doc edits need my fence ruling
   first — **the README fresh-store sentence is PRE-FLAGGED: ASK early** (S1 ASK-1 precedent);
   file the ASK as soon as the plan's file list exists, not at dispatch time.
4. No hard trigger, no cross-slice collision, no locked-contract or design-of-record amendment,
   no OUT touch (s5 content · observe · routing · TUI · federation · external send · authority
   replacement · in-band rotation/supersede · socket-dialect rewrite). Binding-table shape
   change = hard stop + escalate (S2 sanction precedent).

### Plan requirements (carried from the approving review + the trail — bind them in)
- The plan locks the r3 lineage and carries fixture coverage for **S4-SCH2** (structured
  carrier), **S4-SC3** (wedged-host escape hatch), **S4-NG4** (no-cross-seat-metadata
  negative), **S4-FR2** (outbound oversize typed refusal), the **I-PH ceiling-carve-out leg**
  (exactly one value), and the **config-change chain legs** (reviewer PLAN-time check, adopted).
- The full §9 fixture plan maps to tasks with named owners per fixture id; run-FAIL-first where
  the design names a negative.
- E3 legs are gate-run tasks (operator designates the live seats — an operator-judgment item,
  not a plan assumption); every E3 record carries the transport-only qualifier verbatim.
- The §7 round-trip task sequence puts the config-change mutation class EARLY (its backstop:
  before any store is declared persistent — the live team store this slice stands up).
- Custody/ops/usage docs are tasks with claim-honesty acceptance criteria (the [VP-W2]
  every-claim-surface rule; the honesty sweep list in design §7).
- Header normalization (VP watchpoint W1, adopted): your gated PLAN carries an unambiguous
  `HUMAN_GATE_REQUIRED` value — the grill gate is SATISFIED on record; the remaining human
  gates (live-seat designation, the §7 authorization itself, s4-close) are named, not implied.

Not authorized: IMPL (only the conditioned delegated dispatch above), merge (operator's gate,
s4-close), scope expansion, design amendment. Deviations re-engage me.

ACTIONS_GIT_REF: wrote this relay + the parallel m-1 fidelity packet + INDEX rows (relay substrate, git-untracked); ledger entry + commit follow from my seat; no code, no dispatch token in this relay.
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at 2ef9437 pre-ledger-commit; battery 20-ok uncached + vet clean, my runs this session)
Next requested action: draft the gated PLAN; PLAN-REVIEW loop with your implementer; m-1 verdict lands in parallel; on all conditions green you issue the delegated dispatch from your seat.
