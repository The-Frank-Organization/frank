# Sprint s1 — Slice-1: the thin end-to-end conductor relay

**RUN_ID:** `s1` · **Repo:** `frank/` (greenfield, branch `main`) · **Ceremony:** medium · **Opened:** 2026-07-03

## Mandate (from master `s1-dispatch`)

Build the brutally-small end-to-end path:
`mint → connect → render-MVP-form → submit → stamp → validate → lineage → append(crash-atomic) → project → deliver → gate-outbox`

Authorizing relay: `../master/relays/s1-dispatch/PLAN-orchestrator-planner-20260703-130634.md`
(read-only reference — the master governance trail lives in cwd-parent `master/`, not here).

## Spec (read-only references — never edit; escalate spec problems to master)

- `../master/ARCHITECTURE.md` §C4 (the engine) + §C4.3/I-PH (path hygiene)
- `../master/domains/m-1-*/` (store API: `submit`/`project`/`read`, append-only, sole-writer) — FROZEN contract
- `../master/domains/m-2-*/` (FieldSpec envelope) — FROZEN contract
- `../master/domains/m-7-*/` (conductor-core) — our guide domain
- `../master/STEP-1-KICKOFF.md`

## Team

| Seat | Address | Role |
|---|---|---|
| Slice orchestrator | `s1.orchestrator-planner` | decompose, route, sequence (this seat; never implements) |
| Slice orch-reviewer | `s1.orchestrator-reviewer` | adversarial review of orchestration relays (visibility gate) |
| Build pair — planner | `s1-core.planner` | audits, designs, plans the slice |
| Build pair — implementer | `s1-core.implementer` | independent audit, design/plan review, implements on dispatch |

One pair only: the slice is a single serialized commit loop — every component shares core
files and the same asserting exit-gate fixtures; parallel pairs would collide on everything.

External gates (master seats, reached via operator hand-relay):
- **Plan gate:** `m-7.planner` (guide) + `master.orchestrator-reviewer` (VP) must approve the S1 plan **before any `DISPATCH IMPL`**.
- **Fidelity gate:** `m-1.implementer` + `m-2.implementer` review our *usage* of their frozen contracts before `DISPATCH IMPL` is live.
- **Merge:** separate human gate at S1 close.

## Phase sequence

1. AUDIT — pair reads the locked spec + confirms greenfield state (paired, independent)
2. DESIGN — pair-internal design of the slice implementation (spec is locked; this designs the *code*, not the contracts)
3. PLAN — locked plan with exit-gate fixtures as acceptance criteria
4. **EXTERNAL PLAN GATE** — m-7 guide + VP approve; m-1/m-2 fidelity approve
5. IMPL — after gates, via delegated `DISPATCH IMPL`
6. REVIEW-FOLD → SITREP to master at exit gate → human merge gate

## Exit gate (S1-scoped hardened — HARD acceptance, E2 fixtures)

Baseline: accepted-only-through-conductor · system-stamped `FROM` · form/lint before delivery · gate → local outbox item.
Promoted: forged-FROM rejected · forbidden-enum absent-then-rejected · invalid-parent rejected ·
duplicate-sibling double-accept killed · `kill -9` crash matrix (mid-commit, mid-delivery,
post-intake-fsync, around-rename, corrupt-projection rebuild, replayed intake-id) ·
S1-minimal dissolved-linter replay · I-PH path hygiene (no canonical store path in any
seat-facing output) · liveness (inbox = durable truth; pipe write = nudge) · park/wake.

## Scope OUT (do not build — escalate if seemingly needed)

Full FieldSpec registry / 62-check linter / full ~33-check replay (S3) · recovery phases 0–4 +
FIFO durability + GC/genesis + owed-item projection (S2) · consumer schema fields (S4) ·
observe-as-send (Step-2) · routing execution (Step-3).

## Honesty framing (must hold in code AND docs)

Step-1 = provenance + transport, not verified work ("done" = `self_reported`). Only the
serialized-loop double-accept kill (and constrained-grammar R2) are operationally live in S1;
the rest is recorded, not enforced. Confusion-resistance is tool-mediated (removes affordance,
not access); D5 shell-routed confusion is an accepted residual. Do not over-claim.
