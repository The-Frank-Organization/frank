## PLAN — Step-1 build decomposition + build-execution model (for the VP gate; pre-dispatch)

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: step1-plan
PARENT_DISPATCH_ID: step1-prep
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — VP gate on this decomposition + operator ratification before any slice-team spins up
DELEGATED_DISPATCH_AUTHORITY: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator
SUBJECT: Step-1 decomposition — 4 sequential vertical slices, team-per-slice (m-x-guided), Slice-1 spec ready to cut on GO; gate requested before any dispatch

**What this gate decides.** This relay decomposes the Step-1 build and records the operator-decided build-execution model, for your gate. **No code, no slice-team, no dispatch token is issued here** — on your GO I cut the Slice-1 team-boot as a *separate* relay. Spec = `ARCHITECTURE.md` §C4 + the 7 domain docs; approach = `STEP-1-KICKOFF.md`; sequence = `ROADMAP.md` Step-1. This is the first PLAN of the build.

### 1. Build-execution model (operator-decided 2026-07-03)
Each slice is implemented by its **own orchestrator-team instance**, guided by the domain's **m-x leader**. Three layers:
- **master** (this team: CTO + VP + m-1…m-7) — *governs*: decomposes, gates, integrates, owns the architecture-of-record.
- **m-x guide** — the master-team pair that owns the slice's core domain; feeds the slice-team the locked design, answers domain questions, reviews its output for contract-fidelity. Does not itself type the code.
- **slice-team** — a *separate* orchestrator-team instance (its own planner/reviewer + implementer pairs) that plans + implements the slice against the frozen contracts.

This **dogfoods the product's own nested-orchestrator-team vision** (roadmap Step-5) as the build method itself — master governs → m-x guides → slice-team implements.

Proposed conventions (please gate these too):
- slice-team RUN_ID = `s<N>` (Slice-1 = `s1`); seats `s1.orchestrator-planner` / `s1.orchestrator-reviewer` / `s1.<pair>`.
- **one durable relay trail:** slice-team relays live under `master/relays/` namespaced by `s<N>-*` dispatch IDs (one INDEX, one auditable history) — not a separate hidden root.
- **the m-x guide bridges master↔slice-team:** the master CTO dispatches the slice PLAN *to* the slice-team's orchestrator-planner **and** the m-x guide; the guide + the VP gate the slice-team's internal plan before its own `DISPATCH IMPL`. Pair-lineage inside the slice-team is ordinary protocol.
- **one slice at a time (at most two)**, per the operator's cadence; a slice's exit gate is green before the next spins up.

### 2. Decomposition — 4 sequential sections, VERTICAL-SLICE-FIRST
Per `STEP-1-KICKOFF.md` §"Proposed section decomposition" — de-risk *composition* first (one working relay before any taxonomy thickens):
- **S1 — the thin end-to-end relay.** Guide **m-7** (conductor-core owns the engine *and* the composition). The brutally-small path `mint → connect → render-MVP-form → submit → stamp → validate → lineage → append(crash-atomic) → project → deliver → gate-outbox`. Ships with its adversarial + crash/replay fixtures. **The de-risk slice.**
- **S2 — thicken store/engine.** Guide **m-1** (+ m-7 on the loop). Full recovery (phases 0–4), durable FIFO, GC/genesis, and the **owed-item-as-typed-record projection** (materialize-first, scoped to *recorded* items).
- **S3 — thicken forms/lineage.** Guide **m-2**. Full FieldSpec registry + the 62-check linter refactor (with the dissolved-linter replay), `schema_version`, migrator registry.
- **S4 — consumer schema slices + §C4 fixtures.** Guides **m-3 / m-4 / m-5 / m-6** (each its own slice as it comes up). Step-gated-off observe fields (m-3), routing-record schema (m-4), archetype atoms + ceilings (m-5), gate/ODB fields (m-6) + the owed §C4 fixtures.

Sequence rationale = dependency + de-risk: S1 proves the composition; S2 hardens the substrate S1 skeletons; S3 generalizes the form/lineage S1 stubs; S4 lands the consumer contracts on a now-real store. **I fully dispatch only S1 now**; S2–S4 are sequenced, dispatched as each predecessor's exit gate goes green.

### 3. Slice-1 dispatch spec (ready to cut on your GO)
- **Guide:** m-7. **Slice-team:** `s1`.
- **Scope (in):** engine skeleton (single-threaded serialized commit loop + crash-atomic rename commit) · channel/identity (connect = channel-stamped FROM) · the interface guardrail (seats reach only `submit`/`project`/`read`; raw store/config paths absent from the tool surface) + I-PH path hygiene · a **tiny MVP FieldSpec** (only the fields the slice needs) · minimal validate (required-set / enum / seat-scope) · minimal lineage (parent-edge presence + validity) · terminal append `{accepted|rejected|held}` · projection rebuild + `read` · deliver via `project` · one gate producing a local outbox/ODB item.
- **Scope (out — deferred):** full FieldSpec registry + linter refactor (S3) · full recovery phases + FIFO durability + GC (S2) · the owed-item projection (S2) · all consumer schema fields (S4) · observe-as-send (Step-2) · routing execution (Step-3).
- **Consumed, frozen contracts:** m-1 store API (`submit`/`project`/`read`, append-only, sole-writer) + the m-2 FieldSpec envelope — **locked; the slice builds against them, m-1/m-2 implementers review for fidelity, no re-design.** (Boundary contract lives at the slice layer, not this one; each slice dispatch carries its own.)
- **Deliverable:** the first working conductor relay in `frank/` (first commits) + the exit-gate fixtures green.
- **Framing (honesty — must hold in code AND any doc/marketing):** Step-1 = **provenance + transport, not verified work** (observe is Step-2; "done" here is `self_reported`). Of the four sanctioned by-construction claims, only the **serialized-loop double-accept kill** (and, with a constrained grammar, **R2**) are *operationally live* in S1; the others are *recorded*, not enforced, until their later layers. Confusion-resistance is **tool-mediated** (removes affordance, not access; D5 = shell-routed confusion survives).

### 4. Slice-1 exit gate — HARDENED (the binding acceptance test)
Baseline (roadmap): accepted only through the conductor · `FROM` system-stamped not lane-supplied · form/lint before delivery · gate produces a local outbox item.
Promoted (adversarial + crash + replay):
- **Adversarial:** forged `FROM` rejected · forbidden enum absent-then-rejected · invalid parent rejected · **duplicate-sibling double-accept killed** (the serialized-loop guarantee, actually exercised).
- **Crash-atomicity (crash it for real):** `kill -9` mid-commit and mid-delivery → exactly-once outcome + re-issued wake · crash after intake-fsync (no lost intake) · crash before/after the atomic rename (presence = committed) · corrupt-projection rebuild from canonical · replayed intake-id (no double-emission).
- **Path-hygiene (I-PH):** no seat-facing output — every bounce/error included — contains a canonical store path.
- **Liveness:** inbox = durable truth, pipe `write()` = nudge; a busy/dead seat at write time still receives via `project()` on reconnect; a lost wake never sleeps a parked lane forever.
- **park/wake:** a gated lane parks (consumes nothing) + wakes on the operator's verdict.

### 5. Owed §C4 carries — mapped to sections (so none silently drops)
- **code-layer interface-guardrail enforcement** (the one genuinely new build item) + **I-PH path-hygiene fixture** → **S1**.
- ③ known-A / RAISE-ONLY NF fixture → **S1** (guardrail-adjacent) + **S4** (schema surface).
- R2 `gate_referenceable`-per-column negative fixtures → **S3** (with the FieldSpec registry); the constrained R2 grammar partially in **S1**.
- `GRILL_REQUIRED` FieldSpec row → **S3**.
- (optional) `routing_escalation` §J2 member → **S4** (m-4 routing-record slice).
- ⑤ ODB model-name egress fixture → **S4** (m-6 gate/ODB) — dormant, egress-shaped scan.

Per the **materialize-first PLAN rule**: each carry is first written as a typed owed-item record `{owner, source, target surface, disposition path}`; S2's projection guards only *recorded* items — it does **not** make an *unrecorded* observation impossible to miss.

### 6. The gate — what I need from you (VP)
GO / amend on:
- **(a)** the 4-section sequence + the vertical-slice-first ordering;
- **(b)** the team-per-slice build-execution model + the `s<N>` relay conventions + the m-x-guide bridge;
- **(c)** the **m-7-leads-Slice-1** call (engine owns the composition);
- **(d)** the **hardened exit gate as S1's binding acceptance** (crash / adversarial / replay / I-PH — not the happy-path 4);
- **(e)** the **frozen-contracts** stance (m-1/m-2 consumed as locked; fidelity-review, not re-design);
- **(f)** the honesty framing riding into code (`self_reported` done; only serialized-kill + constrained-R2 live).

On GO I cut the `s1` team-boot dispatch (`TO s1.orchestrator-planner + m-7`, `CC` you + operator) as a separate relay; nothing spins up before then.

### 7. Operator-judgment items
- **broad-scope / product-structure:** the whole build's decomposition + the team-per-slice recursion — this relay *is* the escalation; operator ratifies.
- **ambiguous product semantics:** none new (design locked).
- **residual risk accepted:** D5 shell-routed confusion (tool-mediated, not malice-proof) — already operator-accepted; restated so S1 code does not over-claim.

## Verification
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root=master-docs/master/relays <this file>` — run below.
- Pointers: `master/STEP-1-KICKOFF.md` (approach + the build-execution-model note added this relay), `master/ARCHITECTURE.md` §C4 + §C4.3/I-PH, `ROADMAP.md` Step-1, `frank/` (empty repo, branch `main`, 0 commits).

ACTIONS_GIT_REF: wrote this decomposition relay + a build-execution-model note in `master/STEP-1-KICKOFF.md` + an `INDEX.md` row; cwd is not a git repo (docs workspace) so no sha — files on disk. No code, no dispatch token, no slice-team spun up.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` is the code repo, empty (branch `main`, 0 commits, `git status --short` empty).
Next requested action: VP gates this decomposition (RECONCILE GO/amend under `step1-plan/`); on GO the CTO cuts the `s1` Slice-1 team-boot dispatch; first code lands in `frank/`.
