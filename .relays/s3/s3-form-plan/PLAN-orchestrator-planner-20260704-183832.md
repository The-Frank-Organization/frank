## PLAN — PROCEED-TO-PLAN for s3-form (sequencing only; the gated PLAN lock stays in the pair-Planner seat)

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s3-form-plan
PARENT_DISPATCH_ID: s3-form-design
RUN_ID: s3
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the S3-close sign-off stays the operator's separate gate; merge is never implied
DELEGATED_DISPATCH_AUTHORITY: yes — conditions below (F2, non-bootstrap)
FROM: s3.orchestrator-planner
TO: s3-form.planner
CC: s3-form.implementer, s3.orchestrator-reviewer, operator
SUBJECT: PROCEED-TO-PLAN — design r4 accepted at this seat (approving DESIGN-REVIEW verified, GRILL_LOCK satisfied, battery my-run green at 291ab08); draft the S3 PLAN; delegated `DISPATCH IMPL` under the standing F2 conditions incl. m-1 fidelity approve ON RECORD before any dispatch

**Design-completion reconciled at this seat (my own verification, this session):** doc r4 = main@291ab08, all four design commits docs-only (git log --stat); the design-review lineage-gate chain resolves — approving DESIGN-REVIEW `s3-form-design-r4-review-implementer` (verdict approve, DESIGN_DOC_ID `s3-slice-3-design`) parents your r4 request, which parents the r3 must-revise per the narrow-round pattern; GRILL_LOCK `s3-grill-s3-form` folded at §8 with the two operator rows verbatim (Q1 JSON-in-string carrier; Q2 generated disposition pair) + the S2-store non-decision recorded with the operator's words + the grill fence honored (resolved-by-guide/consult/master rows entered as resolved); battery 18 packages ok (uncached) + vet clean at 291ab08 — my own runs, matching yours and the reviewer's. This relay is **sequencing only**: it references the approved design + approving review but carries **no** design-doc lock — the design-review lineage gate watches YOUR seat.

**Draft the S3 PLAN** (`docs/sprints/2026-07-04-s3-slice-3/plans/s3-slice-3-plan.md`): tasks, file list, exit-gate fixtures as acceptance criteria (the D-12 inventory ids are the skeleton), boundary contracts, OUT lines. Your gated PLAN relay carries `DESIGN_LOCK_ID: s3-slice-3-design`, `DESIGN_RECORD_KIND: design-doc`, `PARENT_DISPATCH_ID: s3-form-design-r4-review-implementer` (the approving review — the S2 r2-lineage lesson: keep that exact edge across ALL plan revisions; a narrow plan-fix round re-parents its lock to the approving DESIGN-REVIEW, never to the intervening PLAN-REVIEW). The pair Implementer's PLAN-REVIEW is the plan gate.

**Delegated `DISPATCH IMPL` — yours to issue ONLY under ALL of (the standing F2 set + the slice-specific gates):**
1. Implementer PLAN-REVIEW = approve (lint-clean, correct lineage: review parents your gated PLAN).
2. Mechanical SCOPE_DIFF all-in vs the plan's file list. **Root-doc edits (README.md or any file outside the sprint tree + source) need my fence ruling first** — the S1 ASK-1 / S2 precedent; note the ratified fresh-store qualifier will likely want a README line, so ASK early rather than deviate.
3. **m-1 fidelity approve ON RECORD in .relays/s3/ before ANY dispatch** — the packet is routed in parallel (`s3-fidelity-m1/SITREP-orchestrator-planner-20260704-183833.md`, review object = design §4's seven proposals at main@291ab08). The VP watchpoint binds: PARENT/parent_picker/candidate-set/lineage-field/store-query movement is m-1's gate even inside m-2-owned modules. A must-revise folds bounded (the S2 F-M1 pattern) and re-routes narrow.
4. No hard escalation trigger · no cross-slice collision · no locked-contract/design-of-record amendment · no OUT-item touch (MCP wire-up · observe · routing execution · S4 consumer-field content · TUI/runtime · **the §7 config-change record — master-ruled OUT, OI-S3-CONFIG-CHANGE carries it**). Any failure escalates to master via me; no self-dispatch through it.

**Carried into the PLAN as binding (from the design + rulings — restate, don't re-derive):** the fresh-store qualifier on every claim surface incl. the exit-gate line; the Q2 replay-label guardrails (reconstructed-observe-context wording, no live-coverage tally); the both-flags R2 test column; the D-10 lock-bearing dependent-required grain; I-PH on registry errors + re-render bounce text + the served schema payload; battery green at every commit.

Not authorized by this relay: no design lock from my seat, no IMPL, no merge authority, no scope expansion. Merge stays the operator's separate gate at S3 close.

ACTIONS_GIT_REF: none — sequencing relay only; this file + the parallel fidelity packet + INDEX rows under gitignored .relays/; no tracked-file edit in this action (the reconciliation ledger entry + its commit are cited in RECONCILE.md).
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at 291ab08)
