# Sprint s2 — Slice-2: thicken store/engine + the owed-item projection

**RUN_ID:** `s2` · **Repo:** `frank/` (branch `main`, baseline tag `s1-close` = main@f0dcb85) · **Ceremony:** medium · **Opened:** 2026-07-03

## Mandate (from master `s2-dispatch`)

Thicken the **ENGINE** against the **LOCKED m-1 store contract** — build **against** it, do **not** redefine it:

- **Full crash-recovery, phases 0–4** (S1 shipped crash-atomic commit + dumb-replay recovery; S2 completes the reified phase machinery: validate genesis → scan/quarantine → rebuild projections → restore runtime tables → open).
- **Durable FIFO** — the intake queue's full durability + ordering under crash/restart.
- **GC / genesis** — store genesis + garbage-collection of superseded/consumed derived artifacts, crash-safe (canonical records are never GC'd in v3.0; retain-everything posture, m-7 §10).
- **The owed-item-as-typed-record projection** — an owed-item `record_kind` (m-1 fidelity) + a `project()` over it; **`open` = owed-record with no disposition-record** (materialize-first). **`OI-S1-F11-SWEEP` is its first real customer** and closes at the S2 exit gate.

Authorizing relay: `../master/relays/s2-dispatch/PLAN-orchestrator-planner-20260703-223913.md` (r2)
(read-only reference — the master governance trail lives in cwd-parent `master/`, not here).

## Spec (read-only references — ABSOLUTE paths; never edit; escalate spec problems via s2.orchestrator-planner to master)

- Charter dispatch (scope/gate/exit of record): `/Users/jack/Programming/harness/master/relays/s2-dispatch/PLAN-orchestrator-planner-20260703-223913.md`
- Engine spec: `/Users/jack/Programming/harness/master/ARCHITECTURE.md` §C4 (+ §C4.3 claim boundary / I-PH)
- m-7 conductor-core design-of-record (DESIGN-LOCKED): `/Users/jack/Programming/harness/master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md` — for S2: §2.2 (durable FIFO), §5 (recovery phases 0–4), §6 (fault/quarantine posture), §10 (genesis + GC), §13 (F9/F10/F11)
- LOCKED m-1 store contract: `/Users/jack/Programming/harness/master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md`
- S1 continuity (the code this slice thickens): `docs/sprints/2026-07-03-s1-slice-1/` (design r5, plan r3, RECONCILE.md) + the S1 source at tag `s1-close`
- Sequencing: `/Users/jack/Programming/harness/ROADMAP.md` (Step-1)

**Guide:** m-7 (`m-7.planner`, via operator hand-relay) — engine implementation guidance; it guided S1 and is this team's continuity into the S1 design.
**Fidelity:** m-1 keeps authority over the owed-item `record_kind`, the store layout, and store-API fidelity; `m-1.implementer` fidelity-reviews store-touches **before** their dispatch.

## Team

| Seat | Address | Role |
|---|---|---|
| Slice orchestrator | `s2.orchestrator-planner` | decompose, route, sequence (this seat; never implements) |
| Slice orch-reviewer | `s2.orchestrator-reviewer` | adversarial review of orchestration relays (visibility gate) |
| Build pair — planner | `s2-core.planner` | audits, designs, plans the slice |
| Build pair — implementer | `s2-core.implementer` | independent audit, design/plan review, implements on dispatch |

One pair only (same rationale as S1): every S2 item lives on the single serialized commit loop — recovery, FIFO, GC/genesis, and the projection share `internal/store`, `internal/intake`, `internal/recover`, `internal/gate`, `internal/engine` and the same asserting exit-gate fixtures; parallel pairs would collide on everything.

## Gate model (F2 — non-bootstrap; conditioned delegation)

The S2 plan runs on the **normal pair Implementer plan-review + conditioned delegated dispatch** — no bootstrap guide+VP plan-gate. `DISPATCH IMPL` is delegated **only** under: {Implementer plan-review = approve · no deviation from the dispatched scope/boundary · no hard escalation trigger · no cross-slice collision · no locked-contract or design-of-record amendment}. **m-1 fidelity-reviews store-touches before their dispatch.** If **any** condition fails — including any task touching an OUT item — **escalate back to master (CTO + m-7 guide + VP)**; never self-dispatch through it. The S2-close sign-off is the operator's, exercised separately. Merge is never implied by green fixtures.

## Phase sequence

1. AUDIT — pair onboards to the S1 code + locked spec, confirms the build surface (paired, independent)
2. DESIGN — pair-internal design of the S2 implementation (spec is locked; this designs the *code*, not the contracts)
3. PLAN — locked plan with exit-gate fixtures as acceptance criteria; m-1 fidelity review of store-touches
4. IMPL — via delegated `DISPATCH IMPL` under the F2 conditions above
5. REVIEW-FOLD → S2 exit gate → SITREP to master → operator S2-close sign-off

## Exit gate (HARD acceptance, E2 fixtures — from the s2-dispatch, verbatim scope)

- **Recovery phases 0–4** crash-tested (crash at each phase boundary → correct resumption; no lost/double intake or delivery).
- **Durable FIFO** — order + exactly-once preserved across crash/restart.
- **GC/genesis** crash-safe (genesis idempotent; GC never drops a live record; crash mid-GC recovers).
- **The owed-item projection** — `open = owed-record with no disposition-record` holds; **OI-S1-F11-SWEEP surfaces as open**; a dispositioned owed-item drops from the open set; materialize-first honored.
- **F9/F11 re-run under the new recovery machinery** — the crash matrix passes against full phase-0–4 recovery, and **OI-S1-F11-SWEEP's full class×point sweep is discharged here** (the S1 owed item closes at this gate).
- **No regression** of the S1 gate invariants (serialized-loop kill, crash-atomicity, I-PH, enum byte-exactness, guardrail).

## Scope OUT (do not build — escalate if seemingly needed; escalation blocks any delegated dispatch that touches these)

Full FieldSpec registry + 62-check linter refactor (**S3**) · the **MCP live-adapter + protocol-envelope FieldSpec** (deferred on operator record — no live testbed; S1 is built-but-unwired) · observe-as-send (**Step 2**) · routing execution (**Step 3**) · all consumer schema — observe/routing/archetype/ODB (**S4**).

## Honesty framing (must hold in code AND docs)

Still **provenance + transport, not verified work** ("done" = `self_reported`; observe is Step 2). S2 adds durability hardening + the first governance primitive (the owed-item projection) — real value, but the projection **guards recorded owed-items only; it does NOT make an unrecorded observation impossible to miss** (materialize-first). Confusion-resistance stays tool-mediated (removes affordance, not access); D5 shell-routed confusion is the accepted residual and must be stated beside every exclusivity-shaped claim. Do not over-claim.
