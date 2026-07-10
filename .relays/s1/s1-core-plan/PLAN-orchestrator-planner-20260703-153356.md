## Team s1-core — Slice-1: PROCEED TO PLAN (blocker-2 narrowing ratified; external gates wired into delegation conditions)

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s1-core-plan
PARENT_DISPATCH_ID: s1-core-design
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — the S1 PLAN is m-7-guide + master-VP gated and m-1/m-2 fidelity-reviewed before any `DISPATCH IMPL`; merge is a separate human gate at S1 close
FROM: s1.orchestrator-planner
TO: s1-core.planner
CC: s1-core.implementer, s1.orchestrator-reviewer, operator
IN_REPLY_TO: s1-core-design/SITREP-planner-20260703-152941.md
APPROVED_DESIGN_DOC_ID: s1-slice-1-design
APPROVING_DESIGN_REVIEW_DISPATCH_ID: s1-core-design-r2-review-implementer
PLAN_LOCK_ID: s1-slice-1-plan
BUNDLE_ID: s1-slice-1
OWNER: s1-core
REPO: frank/ (this repo)
BASE: main@876e279
TARGET_BRANCH: main
DELEGATED_DISPATCH_AUTHORITY: yes
DELEGATED_DISPATCH_CONDITIONS: Implementer plan review = approve; pair-Planner dispatch PARENT_DISPATCH_ID points to that approving PLAN-REVIEW relay; that review parents to the pair-Planner PLAN; SCOPE_DIFF_RESULT = all-in; no hard trigger; no boundary-contract deviation; AND the external gates — m-7-guide approve + master-VP approve of the locked plan, and m-1.implementer + m-2.implementer fidelity approves of the consuming surface — are each present as relays in .relays/s1/ (operator-carried) before the token is issued. Absent any one of these, do not dispatch; relay to s1.orchestrator-planner and wait.

Implementer phase scope — PLAN-REVIEW after the plan is drafted.
Current scope: answer design/plan questions, review the Planner's plan, findings inline or as a PLAN-REVIEW relay.
Not in current scope: source/test edits, implementation branches, commits (except tracked sprint plan docs by the pair Planner), PRs, scaffolding, or prototype implementation.
Implementation begins only after a current relay under the active run's RELAY_ROOT contains the exact literal token `DISPATCH IMPL` bare, unfenced, un-backticked, alone on its own line, and addressed to the Implementer in `TO`. Urgency is not dispatch; inline, quoted, fenced, CC-only, cross-read, or non-addressee mentions are inert.

Approved design context:
DESIGN_DOC_ID `s1-slice-1-design` r2 at docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md (main@5622516), Implementer verdict approve (`s1-core-design-r2-review/DESIGN-REVIEW-implementer-20260703-152445.md`), five r1 blockers folded + fold-verified, scope-absorption clean. Operator decision D-1: conductor stack = Go. Guide constraints R-1..R-3 / O-1..O-3 honored; F9 named whole.

This PROCEED-TO-PLAN is sequencing only. It does not carry the gated design-doc lock. You (pair Planner) emit the gated PLAN from `FROM: s1-core.planner` with `DESIGN_LOCK_ID: s1-slice-1-design`, `DESIGN_RECORD_KIND: design-doc`, and `PARENT_DISPATCH_ID: s1-core-design-r2-review-implementer` (the approving DESIGN-REVIEW).

**Blocker-2 narrowing: RATIFIED** (reconciliation ledger entry 4, main@876e279), on three conditions the PLAN must carry:
1. The narrowing is stated explicitly in the PLAN text so the m-7-guide + master-VP gate reviews it (their gate is the authoritative check above this ratification).
2. Claim honesty: any S1 doc/surface describing grants states the S3 landing plainly (pair-Planner conditional-render lands with the full lineage engine).
3. No S1 schema/format decision forecloses the S3 conditional-render landing.
Operator veto path remains open until the plan gate; do not treat the ratification as final until the external gates pass.

LOCKED scope (from the approved design; the plan concretizes into tasks/files):
1. The S1 slice per DESIGN_DOC_ID s1-slice-1-design — engine (serialized commit loop, intake journal full semantics, Package-A rename pivot, dumb-replay recovery, rebuild-before-open), channel/identity (attach-time mint, system-stamped FROM, binding table), interface guardrail (registry exactly {submit, project, read} + I-PH), MVP FieldSpec (as designed, incl. full frozen §J2 gate_category set byte-exact), validate + lineage minimal, terminal append + projection rebuild + read, deliver via project, gate → typed local outbox item (O-1..O-3), park/wake.
2. The full fixture suite keyed 1:1 to the exit gate: B1-B4, A1-A4, C1-C7, R1, P1, L1, W1, F9, F11, G, H, SWEEP (E2, runnable locally).
3. The three chartered owed carries at their design homes, materialize-first (typed records into the plan).

Scope list for delegated-dispatch SCOPE_DIFF (the fence; the locked plan enumerates concrete paths within it):
- new Go source/test/fixture/build files under the frank/ repo root (module files, cmd/, internal-style dirs — as the plan names them)
- docs/sprints/2026-07-03-s1-slice-1/plans/** (the plan doc; pair-Planner-committed)
- .relays/s1/** (relay substrate; not code)
- OUT for IMPL: docs/sprints/** other than plans/ + the IMPL-phase results/ artifacts; anything under ../master or ../extracted (read-only always); .gitignore and existing tracked docs except where the plan explicitly lists them.

Fold into the plan (each as a named task or acceptance criterion):
- The guide-gate checklist as the plan's rubric — all 7 items demonstrably satisfied in the plan text (ROADMAP.md §guide-gate; the guide pre-published it to be built to).
- External-gate sequencing as explicit plan steps: (a) pair plan-review; (b) plan relayed to m-7.planner + master.orchestrator-reviewer via operator; (c) consuming-surface fidelity packet to m-1.implementer + m-2.implementer (the audits' E1 contract enumerations are the review object; include the DI-2 realization record for m-1); (d) only after all approves → delegated `DISPATCH IMPL`.
- MCP Go SDK per-seat-socket capability check as the FIRST implementation task (de-risk; fallback + invariants already named in the design — per-seat channels, stamped identity, exactly the 3-verb registry).
- m-1 §13 carries as acceptance detail (per the completion SITREP's open-items list).
- The S1-minimal replay (R1) task with its caught / genuinely-obsolete / uncovered-S3 reporting shape.
- Red→green fixture ordering per Superpowers writing-plans/TDD where the shape allows (fixtures F9/F11/C-matrix are the asserting spine).

Boundary contract (locked, unchanged from the charter):
- Writes: append-only relay-store records (terminal byte-exact {accepted,rejected,held}) + a local outbox/ODB item
- Reads: the locked m-1 store API + m-2 FieldSpec envelope
- Target entity: the conductor's committed relay + its rebuilt projection
- Downstream consumer: seat inboxes (via project) + the local outbox
- Contract: submit/project/read; channel-stamped FROM; MVP FieldSpec; crash-atomic commit
- Proof: the S1-scoped hardened exit gate (E2 fixtures)
- No-consumer action: n/a — S1's consumer is the relay flow itself

Out of scope:
- Everything on the ROADMAP scope-OUT list; the ⑤ carry (cite as S4-bound if mentioned); m-2 §9 migrators; organic master-trail corpus as R1 gate inputs; pair-Planner grant rendering (ratified narrowing); any edit to ../master or ../extracted.

Tests/verification plan must include:
- every exit-gate fixture id with its red→green command; the kill -9 crash matrix run for real (child-process kills at the design's injection points); the I-PH grep fixture over ALL fixture outputs; the SWEEP claim-honesty check over seat-/user-facing strings; `go test ./...` (or the design's equivalent) as the E2 umbrella.

Ceremony downgrade record: none — medium tier maintained; no downgrade proposed.

Deliverable:
Locked plan via Superpowers writing-plans at docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md (tracked, pair-Planner-committed) + the gated PLAN relay from your seat (lineage per above) + Implementer PLAN-REVIEW relay + SCOPE_DIFF (all-in required for later delegation) + ACTIONS_GIT_REF + FINAL_GIT_STATUS_SHORT. Then HOLD for the external gates; I relay the locked plan to the m-7 guide + master VP and the fidelity packet to m-1/m-2. No `DISPATCH IMPL` before all external approves land in .relays/s1/.

ACTIONS_GIT_REF: wrote this relay + an INDEX.md row under .relays/s1/ (gitignored substrate); reconciliation + ratification committed on main@876e279
FINAL_GIT_STATUS_SHORT: none — clean tree
