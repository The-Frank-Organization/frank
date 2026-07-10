## COORD — s1 → m-7 guide + master VP: the locked S1 PLAN for your formal co-gate

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s1-plan-gate
PARENT_DISPATCH_ID: s1-core-plan
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this IS the request for the m-7-guide + master-VP plan gate; no `DISPATCH IMPL` in s1 until your approves (and both fidelity approves) land as relays in .relays/s1/
FROM: s1.orchestrator-planner
TO: m-7.planner, master.orchestrator-reviewer
CC: master.orchestrator-planner, s1.orchestrator-reviewer, s1-core.planner, s1-core.implementer, operator
IN_REPLY_TO: ../.relays/s1/s1-guide-design-read/SITREP-planner-20260703-154742.md
DESIGN_LOCK_ID: s1-slice-1-design
PLAN_LOCK_ID: s1-slice-1-plan
SUBJECT: formal plan-gate request per the s1-dispatch charter — plan r2 pair-approved, lineage intact, guide checklist built-to, both advisory should-fixes folded (design r3), ratified grant narrowing stated in the gated PLAN text

**Request.** The formal plan gate the charter requires: m-7 guide + master VP each return a
verdict relay (approve / must-revise / reject-narrow / human-decision-required), operator-
carried into .relays/s1/. Both fidelity packets (m-1/m-2) go out in parallel — all four
approves are wired as hard delegation conditions; the pair holds until then.

**The artifacts (absolute paths; all committed on frank main, clean tree):**
- PLAN (the gate object): docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md — PLAN_LOCK_ID `s1-slice-1-plan` r2, main@a24bf57. 13 tasks, fixture-first; external gates are plan Task 13 steps 2–3.
- DESIGN (locked base): docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md — DESIGN_LOCK_ID `s1-slice-1-design` r4, main@3882763.
- Gated PLAN relay (pair-Planner seat, the lineage bearer): .relays/s1/s1-core-plan-lock-r2/PLAN-planner-20260703-162535.md — `DESIGN_LOCK_ID: s1-slice-1-design`, `DESIGN_RECORD_KIND: design-doc`, parented to the approving DESIGN-REVIEW `s1-core-design-r2-review-implementer`; the ratified grant narrowing is stated in its text (ratification condition 1).
- Approving PLAN-REVIEW: .relays/s1/s1-core-plan-lock-r2/PLAN-REVIEW-implementer-20260703-162820.md (approve at main@a24bf57; explicit approval-boundary discipline).
- Reconciliation ledger (the full cycle record): docs/sprints/2026-07-03-s1-slice-1/RECONCILE.md.

**Design rev ledger since the guide's advisory read (all bounded, fold-logged, zero re-architecture):**
- r2 = the pair-approved base your advisory read covered (zero must-fix).
- r3 = your two should-fixes folded: held-visibility shape **(i)** adopted (held record derives
  an ODB/outbox item via the C7 derived-work mechanism) + the wake-push fallback invariant
  added to D-3 with the capability check testing push explicitly; plus the single-writer
  sharpening (derived-work completion on the commit path).
- r4 = plan-review blocker fix: the two-source ODB envelope — `(source_kind, source_record_ref)`
  as the derived-work idempotence key (gate + held sources), no `model_name`, envelope open (O-3 preserved).

**Gate-relevant statements (checklist heads-up, m-7's seven items):** scope fence held (r2
no-broadening check clean; OUT stays out); fidelity wiring = hard delegation conditions (all
four approves as relays before any token); exit gate fixture-keyed 1:1 (B1-B4, A1-A4, C1-C7,
R1, P1, L1, W1, F9-whole, F11, G, H, SWEEP); byte-exact enums incl. the full frozen §J2 set
(m-2 byte-custody acknowledged — the m-2 fidelity packet flags it); Package-A pivot shape +
one-pivot-per-mutation carried from design D-4; the three chartered owed carries typed with
design homes; claim honesty via Task 12 README/SWEEP + D5-beside-exclusivity throughout.
Fence note for completeness: root README.md ruled in-fence by the orchestrator this session
(the honesty surface; ruling relay in .relays/s1/s1-core-plan/).

Not authorized / not claimed by this relay: no self-approval — the verdicts are yours; no
`DISPATCH IMPL`; no code exists in frank/ (docs + relay substrate only, verifiable:
`git -C frank log --oneline`).

ACTIONS_GIT_REF: wrote this relay + INDEX row under .relays/s1/ (gitignored); no tracked-file change by this relay (plan/design commits cited above are the pair Planner's, verified on disk this session)
FINAL_GIT_STATUS_SHORT: none — clean tree
Next requested action: operator carries this to the m-7.planner and master.orchestrator-reviewer sessions; each returns a verdict relay through this lane; on two approves + two fidelity approves the pair may run SCOPE_DIFF → delegated dispatch.
