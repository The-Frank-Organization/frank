## AUDIT — s2-core paired independent audit: onboard to the S1 build surface + map the S2 mandate onto it (read-only)

ROLE: Orchestrator Planner
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: s2-core-audit
PARENT_DISPATCH_ID: s2-dispatch
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s2.orchestrator-planner
TO: s2-core.planner, s2-core.implementer
CC: s2.orchestrator-reviewer, operator
SUBJECT: AUDIT dispatch — paired, independent: onboard to the S1 code (you did not build it); map every S2 exit-gate line to locked spec + current code; enumerate the m-1 store-touch surface; 4-bucket verdicts per S2 IN item

**Operator hold RELEASED (2026-07-04, "alright lets go issue" — recorded in RECONCILE.md).** This is the first work dispatch of RUN_ID s2.

### What this is

Paired **independent** audits (do not coordinate before writing; reconciliation is my job). Read-only: no code, no design, no plan — findings + evidence. You are auditing the **S1 baseline** (`main`, tag `s1-close` = f0dcb85; scaffold commit 6ceeb5d on top) against the **S2 mandate** (docs/sprints/2026-07-03-s2-slice-2/ROADMAP.md — the s2-dispatch scope verbatim).

### Scope (IN for this audit)

1. **Onboard to the S1 code first — you did not build it.** Read the s1 sprint docs (docs/sprints/2026-07-03-s1-slice-1/: design r5, plan r3, RECONCILE.md — the deviation rulings and the owed item live there) and the S1 source (internal/* + test/*). Fresh adversarial eyes are the point: anything in the S1 code that looks load-bearing-but-fragile for S2's machinery is a finding.
2. **Spec-to-exit-gate map.** Every S2 exit-gate line (ROADMAP) maps to locked spec text — m-7 design §2.2 (durable FIFO), §5 (recovery phases 0–4), §6 (fault/quarantine), §10 (genesis + GC), §13 (F9/F10/F11); ARCHITECTURE §C4; m-1 §5/§6 — with file:line cites. Name any exit-gate line you cannot ground in locked text (spec gap ⇒ escalate, never improvise).
3. **4-bucket verdict per S2 IN item** (recovery phases 0–4 · durable FIFO · GC/genesis · owed-item projection): what does S1 already ship (already-closed — do not rebuild; e.g. rebuild-before-open discipline, derived-work completion scan, intake−outcomes re-enqueue, checksum field, crashpoint registry + child-SIGKILL harness, fsio rename-pivot counters), what is genuinely still-open, what narrows. The S1 recovery is deliberately dumb replay — the audit maps which of phases 0–4 exist in embryonic form vs not at all (genesis validation, checksum quarantine disposition, reified phase ordering, runtime-table restore as a phase, segment rotation).
4. **The owed-item projection surface.** Locked semantics: `open = owed-record with no disposition-record`, materialize-first; first customer = `OI-S1-F11-SWEEP` (typed record in s1 RECONCILE.md entry 11 — owner, source, target surface, disposition path). Guide advisory on record (s1 design §9.4): at S2 the C7 derived-work scan becomes an **instance of** the owed-item projection, not a parallel mechanism — audit what that implies for internal/gate/derived.go.
5. **m-1 store-touch enumeration (the fidelity surface).** List every store-layout/record-shape touch S2 plausibly needs: the owed-item `record_kind` (m-1's authority — enumerate the proposal surface, do not fix its shape), disposition-record shape, genesis record, quarantine/ location, GC target set, journal segment rotation. This list is what m-1.implementer fidelity-reviews before any dispatch — completeness here saves a bounce later.
6. **Claim-boundary probes.** (a) "Durable FIFO — exactly-once" exit-gate wording vs m-7 §2.2's "at-least-once intake, exactly-once effect" — pin the honest claim. (b) GC exit-gate line "never drops a live record" vs m-7 §10's posture (canonical records NEVER GC'd; GC compacts only derived artifacts + drained journal segments) — pin the GC target set from locked text. (c) Materialize-first honesty: the projection guards recorded owed-items only.
7. **Duplicate/already-built gate** (protocol standard): anything S2-shaped already present, feature-flagged, or dead-pathed — recommend promote/reuse, never rebuild. The S1 crash harness (crashpoint.Names() + FRANK_TEST_CRASHPOINT child-SIGKILL) was explicitly built to be the S2 full-sweep machinery — confirm its reusability for the OI-S1-F11-SWEEP full class×point sweep and name any gap.

### Out of scope (escalate via me, never absorb)

Everything on the ROADMAP OUT list (S3 registry/linter · MCP live-adapter (deferred, no testbed) · observe (Step 2) · routing (Step 3) · consumer schemas (S4)) · any locked-contract or design-of-record amendment · any edit under ../master or ../extracted · any code/design/plan work in this phase.

### Deliverable (each seat, independently)

One lint-clean AUDIT relay: `.relays/s2/s2-core-audit/AUDIT-<role>-<YYYYMMDD-HHMMSS>.md`, FROM your seat, TO s2.orchestrator-planner, CC operator. Carrying: the spec-to-gate map (file:line), the four 4-bucket verdicts with PRIMARY_BUCKET lines, the m-1 store-touch enumeration, the claim-boundary probe answers, reuse/duplicate findings, any spec gaps or escalations, evidence levels per claim (E1 cites; E2 for anything you executed — running `go test` yourself is encouraged), and `FINAL_GIT_STATUS_SHORT` (read-only proof). Append your INDEX row (end-of-file, write order).

### Acceptance criteria (for this audit)

- Every S2 exit-gate line mapped to locked text or flagged as a gap — none unaccounted.
- Four PRIMARY_BUCKET verdicts with evidence; no bookkeeping contradicting a recommendation.
- m-1 store-touch list complete enough for the fidelity packet.
- Claim-boundary probes answered from locked text, not judgment.
- No file modified anywhere (both trees clean; frank/ at 6ceeb5d unless the operator advances it).

Operator-judgment items: none new — D5 residual restated (accepted, on record); the MCP deferral stands.

ACTIONS_GIT_REF: none — read-only dispatch authored as this relay file + an INDEX row; .relays/ is gitignored operational substrate, no tracked-file edit in this action.
FINAL_GIT_STATUS_SHORT: none — clean tree
