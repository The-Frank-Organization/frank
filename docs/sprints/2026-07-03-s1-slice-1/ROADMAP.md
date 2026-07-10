# Sprint s1 — Slice-1: the thin end-to-end conductor relay

**RUN_ID:** `s1` · **Repo:** `frank/` (greenfield, branch `main`) · **Ceremony:** medium · **Opened:** 2026-07-03

## Mandate (from master `s1-dispatch`)

Build the brutally-small end-to-end path:
`mint → connect → render-MVP-form → submit → stamp → validate → lineage → append(crash-atomic) → project → deliver → gate-outbox`

Authorizing relay: `../.relays/s1/s1-dispatch/PLAN-orchestrator-planner-20260703-130634.md`
(read-only reference — the master governance trail lives in cwd-parent `master/`, not here).

## Spec (read-only references — ABSOLUTE paths per the guide's context-brief; never edit; escalate spec problems to master)

- Charter dispatch: `.relays/s1/s1-dispatch/PLAN-orchestrator-planner-20260703-130634.md`
- Guide context-brief: `.relays/s1/s1-dispatch/COORD-planner-20260703-134029.md`
- **Guide-gate checklist (build the PLAN to this):** `.relays/s1/s1-dispatch/SITREP-planner-20260703-133102.md`
- Build strategy + hardened exit gate: `master-docs/master/STEP-1-KICKOFF.md`
- Engine spec: `master-docs/master/ARCHITECTURE.md` — §C4.1 (engine), §C4.2 (18-row seam matrix), §C4.3 (claim boundary + I-PH, `:450-463`), owed Step-1 fixture ledger (`:477-482`, just above §C5)
- m-7 conductor-core design-of-record (DESIGN-LOCKED r5 + c6/c6.1) — `master-docs/master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md` — for S1: §2 (process model), §3 (commit pipeline), §4 (Package-A rename pivot; presence=committed), §6 (fault→`held`), §8 (interface guardrail + wake), §12 (seam matrix), §13 (fixtures F1–F11, incl. F9 no-stale-re-emission + F11 one-pivot-per-mutation), §16 (claim-sweep rules)
- FROZEN m-1 contract (store API: `submit`/`project`/`read`, append-only, sole-writer): `master-docs/master/domains/m-1-trust-identity/design/2026-06-28-trust-identity-design.md`
- FROZEN m-2 contract (FieldSpec envelope): `master-docs/master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md`
- Sequencing: `master-docs/ROADMAP.md` (Step-1)
- Dissolved-linter replay corpus (upstream baseline, historical lint failures): the upstream protocol release corpus (governance workspace)

**Reaching the guide:** write a relay file `FROM` your own seat, `TO: m-7.planner`, tell the operator; they carry it. The guide answers strictly from the locked design-of-record and escalates lock-amendment questions to master.

## Guide-gate checklist (the PLAN is reviewed against this — from `SITREP-planner-20260703-133102.md`)

1. **Scope fence** — IN list only; every OUT item absent; escalate-don't-expand stated.
2. **Contract-fidelity wiring** — m-1/m-2 fidelity reviews sequenced in the plan *before* `DISPATCH IMPL` goes live.
3. **Exit gate mapped to fixtures** — every hardened-gate line has a named fixture + test plan.
4. **Byte-exact enum** — terminal `{accepted, rejected, held}` exactly; no synonyms; `bounced` is not a value token.
5. **Pivot shape right from slice 1** — canonical-record `rename()` is the single commit pivot (fsync-before-rename; presence=committed; projections derived); outcome records reference `intake_id`. Getting this wrong forces S2 re-architecture.
6. **Owed carries materialize-first** — typed owed-item records `{owner, source, target surface, disposition path}` for: code-layer interface-guardrail enforcement, the I-PH fixture, the ③ known-A/RAISE-ONLY guardrail-adjacent portion.
7. **Claim honesty in code + docs** — see below; §16 claim-sweep applies to anything seat- or user-facing we write.

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
not access); D5 shell-routed confusion (a same-uid shell-bearing seat can bypass) is an
accepted residual and must be stated **beside every exclusivity-shaped claim** ("only writer",
"sole egress", …). Do not over-claim.
