# Sprint s3 — Slice-3: the full form system (FieldSpec registry + linter dissolution)

**RUN_ID:** `s3` · **Repo:** `frank/` (branch `main`, baseline tag `s2-close` = main@b322b6d; code surface verified identical at HEAD 7a8b9d7 — ledger-docs-only delta) · **Ceremony:** medium · **Opened:** 2026-07-04

## Mandate (from master `s3-dispatch`)

**S3 makes frank speak the real protocol.** Build the **FORM SYSTEM** against the **LOCKED m-2 design-of-record** — build against it, do **not** redefine it:

- **Full FieldSpec registry** — the complete typed-envelope field catalog (field · owner ∈ {system, seat_scoped_enum, agent_enum_pick, free_text} · type · required-when · enum · seat_scope · gate_referenceable), replacing the S1 MVP dialect (`internal/fieldspec`'s flat 6-enum registry); registry-driven render + validate; **fill-time authority** (forbidden options absent from the rendered form, not rejected after).
- **The 62-check linter dissolution** — the upstream `relay-lint`'s checks re-homed per the m-2 §10 dissolve/survive map: each check **dissolved into form-validation/lineage** or **retained as an explicit post-submit check** or **genuinely obsolete** — a per-check disposition table for all 62, no silent drops.
- **The FULL dissolved-linter replay (the S1-deferred F1 gate):** the historical upstream lint-failure corpus (243 fixtures, 14 categories) through the new validation; every failure **caught-or-genuinely-obsolete**; replaces S1's name-heuristic classifier (`test/replay/classmap.go`) whose deferral bucket is literally named `uncovered-S3`.
- **`schema_version` + the migrator registry** — the version stamp in its `system_only` envelope home (stamped since S1) + the migration mechanism (registry + read-time apply path + a fixture proving a v(n)→v(n+1) walk + fixtured refusal/bounce legs). Zero real migrators; the mechanism ships first.
- **Owed §C4 carries landing here:** **R2 `gate_referenceable`-per-column negative fixtures** (no model-derived predicate enters a gate through any registry column) · the **`GRILL_REQUIRED` FieldSpec row** (m-6-F6).
- **Re-render/drift** — a seat holding a stale rendered form gets a bounce carrying "re-render" (S1 seeded the digest mechanism; S3 keeps it correct under a *changing* registry).

Authorizing relay: `../.relays/s3/s3-dispatch/PLAN-orchestrator-planner-20260704-150904.md`
(read-only reference — the master governance trail lives in cwd-parent `master/`, not here).

## Spec (read-only references — ABSOLUTE paths; never edit; escalate spec problems via s3.orchestrator-planner to master)

- Charter dispatch (scope/gate/exit of record): `.relays/s3/s3-dispatch/PLAN-orchestrator-planner-20260704-150904.md`
- **m-2 design-of-record (PRIMARY — this slice IS this domain)** — `master-docs/master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md` — FieldSpec shape §4; predicate vocabulary §5; X-overflow §6; versioning §9; the 62-check dissolve/survive map §10; GATE-1 §11; consumer contract §12; ACs §14; gate/delivery/ODB/routing/computed-field specs §17; readiness folds §18
- Engine spec: `master-docs/master/ARCHITECTURE.md` §C4 (+ §C4.3 claim boundary / I-PH; the §C4 owed-carry ledger this slice discharges)
- m-7 conductor-core design-of-record (consulted contract — trusted-config seam): `master-docs/master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md`
- m-1 store contract (consumed via the locked API only): `master-docs/master/domains/m-1-trust-identity/design/2026-06-28-trust-identity-design.md`
- the upstream linter + fixture corpus (read-only replay input, DO-NOT-COPY as design): the upstream protocol release corpus (governance workspace, `tools/`) (relay-lint.py = the 62 checks; relay-lint-fixtures/ = the corpus, 243 fixtures)
- S1+S2 continuity (the code this slice thickens): `docs/sprints/2026-07-03-s1-slice-1/` + `docs/sprints/2026-07-03-s2-slice-2/` + the source at tag `s2-close`
- Sequencing: `master-docs/ROADMAP.md` (Step-1)

**Guide:** m-2 (`m-2.planner`, via operator hand-relay) — the domain continuity; this slice IS m-2's domain.
**Consult:** m-7 (`m-7.planner`) on engine/config seams — the registry rides the trusted config (per-domain sections, single top-level digest, loaded once at trusted startup); consult, don't improvise.
**Fidelity:** `m-1.implementer` fidelity-reviews any store-API touches before their dispatch (expected light — S3 is mostly above the store). **[VP-W] Lineage movement is an m-1 fidelity trigger even inside m-2-owned modules:** changes to `PARENT`, `parent_picker`, candidate-set derivation, system-filled lineage fields, or store *query* semantics require m-1 fidelity review before delegated-dispatch eligibility.

## Team

| Seat | Address | Role |
|---|---|---|
| Slice orchestrator | `s3.orchestrator-planner` | decompose, route, sequence (this seat; never implements) |
| Slice orch-reviewer | `s3.orchestrator-reviewer` | adversarial review of orchestration relays (visibility gate) |
| Build pair — planner | `s3-form.planner` | audits, designs, plans the slice |
| Build pair — implementer | `s3-form.implementer` | independent audit, design/plan review, implements on dispatch |

One pair only (same rationale as S1/S2): every S3 item lives on one schema + one validate path — the registry, the dissolution, the replay, versioning, and the R2/GRILL_REQUIRED carries all share `internal/fieldspec`, `internal/engine/submit.go`, `internal/lineage`, `internal/config`, `test/replay` and the same asserting exit-gate fixtures; parallel pairs would collide on everything.

## Gate model (F2 — non-bootstrap; conditioned delegation)

The S3 plan runs on the **normal pair Implementer plan-review + conditioned delegated dispatch** — no bootstrap guide+VP plan-gate. `DISPATCH IMPL` is delegated **only** under: {Implementer plan-review = approve · no deviation from the dispatched scope/boundary · no hard escalation trigger · no cross-slice collision · no locked-contract or design-of-record amendment}. **m-1 fidelity-reviews store-touches (incl. lineage-movement per the VP watchpoint) before their dispatch.** If **any** condition fails — including any task touching an OUT item — **escalate back to master (CTO + m-2 guide + VP)**; never self-dispatch through it. The S3-close sign-off is the operator's, exercised separately. Merge is never implied by green fixtures.

## Phase sequence

1. AUDIT — pair onboards to the S1+S2 code + the m-2 locked spec, confirms the build surface (paired, independent)
2. DESIGN — pair-internal design of the S3 implementation (spec is locked; this designs the *code*, not the contracts)
3. PLAN — locked plan with exit-gate fixtures as acceptance criteria; m-1 fidelity review of store/lineage touches
4. IMPL — via delegated `DISPATCH IMPL` under the F2 conditions above
5. REVIEW-FOLD → S3 exit gate → SITREP to master → operator S3-close sign-off

## Exit gate (HARD acceptance, E2 fixtures — from the s3-dispatch, verbatim scope)

- **Registry live end-to-end:** a real relay in the team's actual header vocabulary (ROLE/PHASE/AUTHORITY/DISPATCH_ID/lineage/…) renders, validates, commits, projects through frank — the S1 thin path running the *full* protocol envelope.
- **Fill-time authority proven by negatives:** forbidden enum members **absent from the rendered form** per seat-scope; a forged/out-of-scope field submission bounces pre-append.
- **The per-check disposition table complete** (all 62: dissolved / retained / obsolete — each with its fixture or rationale) + **the FULL replay green**: the historical corpus caught-or-genuinely-obsolete, no silent drops. **[VP-W] Obsolete-adjudication rule:** every *obsolete* disposition names the **concrete vanished surface, legacy-only path, or replaced invariant**; an obsolete disposition resting on a **design-of-record change** escalates to master before S3 close.
- **R2 negatives green** (`gate_referenceable`-per-column) · **GRILL_REQUIRED row present + rendered**.
- **`schema_version` + migrators:** a fixture-proven v(n)→v(n+1) migration walk; **[VP-W] plus a fixtured refusal/bounce leg** for unknown/future, unversioned, and mismatched records (bounce/re-render, never silent coercion). A backward/downgrade migrator is **not** required — introducing one is a scope expansion (escalate).
- **No regression:** S1+S2 suites green; enum byte-exact; I-PH on every new surface (registry errors + bounce text included); the owed-item projection + recovery/FIFO/GC untouched-and-green.
- **materialize-first, with teeth:** any S3 finding meant to be *guarded* is first a typed owed-item record — the live S2 mechanism exists; use it where practical (the OI-S1-F11-SWEEP precedent).

## Scope OUT (do not build — escalate if seemingly needed; escalation blocks any delegated dispatch that touches these)

MCP live-adapter / wire-up (**next after S3, operator's call**) · observe-as-send + evidence fields (**Step 2**) · routing execution (**Step 3**) · consumer schema slices — observe/routing/archetype/ODB field *content* (**S4**) · any TUI/runtime work. **[VP-W] The S4 line, precisely:** S3 **may** define registry slots, types, ownership classes, validation mechanics, and referenceability; S3 **may not** choose sibling-owned consumer *semantics* (what an observe/routing/archetype/ODB field means or which values it takes) — expression capacity in, content out.

## Honesty framing (must hold in code AND docs)

Still **provenance + transport, not verified work** (observe is Step 2). S3 adds the *form system* — determinism + fill-time authority — not evidence or done-ness. The linter dissolution claim is proven **by the executed replay**, never asserted. Confusion-resistance stays tool-mediated (removes affordance, not access); D5 shell-routed confusion is the accepted residual and must be stated beside every exclusivity-shaped claim. I-PH: no canonical store/config/outbox path in any seat-delivered surface, bounce text included. Do not over-claim.
