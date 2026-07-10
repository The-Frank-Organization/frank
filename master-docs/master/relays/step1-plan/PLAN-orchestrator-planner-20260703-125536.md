## PLAN — Step-1 build decomposition + build-execution model (r2 — VP amendments folded; for re-gate + operator ratification)

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: step1-plan
PARENT_DISPATCH_ID: step1-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — operator ratifies the build-execution model + VP re-gates before any s1 team-boot
DELEGATED_DISPATCH_AUTHORITY: no
GRILL_REQUIRED: no
IN_REPLY_TO: step1-plan/RECONCILE-orchestrator-reviewer-20260703-125139.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator
SUBJECT: Step-1 decomposition r2 — dissolved-linter replay re-scoped (S1-minimal / S3-full), guide+VP gate narrowed to bootstrap, m-1/m-2 fidelity edge wired in; re-gate + operator ratification

**Amendments folded (VP RECONCILE `…-125139` → revise).** Three narrow blockers, all accepted; each fixed below. Your concurrences carry unchanged: vertical-slice-first, m-7 as S1 composition guide, `frank/` as the code repo, and the Step-1 honesty boundary (`self_reported`, tool-mediated confusion-resistance, D5 residual). Still **no code, no slice-team, no dispatch token** — this re-gates.
- **F1 → dissolved-linter replay re-scoped.** The *full* "~33 checks dissolve" replay is an **S3** gate (it needs the 62-check refactor + full FieldSpec registry); **S1** carries a **minimal** replay over only its MVP-FieldSpec-covered historical failures. S1's gate is now labeled the **S1-scoped** hardened gate — not the full one. `STEP-1-KICKOFF.md` exit-gate section updated to match.
- **F2 → guide+VP plan-gate narrowed to S1/bootstrap.** It is **not** a standing per-slice approval gate. S2+ use the **normal pair Implementer plan-review + delegated dispatch**; a slice plan escalates back to master **only** on named triggers.
- **F3 → m-1/m-2 fidelity edge wired into the S1 contract.** The S1 boundary/plan-review contract names **m-1.implementer** (store-API usage fidelity) + **m-2.implementer** (FieldSpec-envelope usage fidelity) as **required reviewers of the slice's consuming surface**, their approve a **precondition to the s1 slice-team's delegated dispatch** — usage-fidelity only, never reopening the locked contracts. Boot addressing updated to wire them in.

### 1. Build-execution model (operator-decided 2026-07-03; VP-amended F2)
Each slice is implemented by its **own orchestrator-team instance**, guided by the domain's **m-x leader** — dogfooding the product's nested-orchestrator-team vision (roadmap Step-5) as the build method:
- **master** (CTO + VP + m-1…m-7) — *governs*: decomposes, gates, integrates, owns the architecture-of-record.
- **m-x guide** — the master-team pair owning the slice's core domain; feeds the locked design, answers domain questions, reviews output for contract-fidelity; does not type the code.
- **slice-team** — a *separate* orchestrator-team instance (own planner/reviewer + pairs) that plans + implements against the frozen contracts.

Conventions (please gate; operator ratifies):
- slice-team RUN_ID `s<N>` (Slice-1 = `s1`); seats `s1.orchestrator-planner` / `s1.orchestrator-reviewer` / `s1.<pair>`.
- **one durable relay trail:** slice-team relays under `master/relays/`, namespaced `s<N>-*` (one INDEX, one auditable history).
- **the m-x guide bridges master↔slice-team:** the master CTO dispatches the slice PLAN to the slice-team's orchestrator-planner **and** the m-x guide.
- **plan-gate depth is risk-scaled, not blanket (F2):**
  - **S1 / bootstrap:** the m-x guide **and** the VP gate the slice-team's internal plan before its `DISPATCH IMPL` — justified by first-code / unproven-composition risk.
  - **S2 and later:** the **pair Implementer plan-review** is the plan gate; the master slice-dispatch grants `DELEGATED_DISPATCH_AUTHORITY: yes` with conditions `{Implementer plan-review = approve; no deviation from the slice's dispatched scope/boundary; no hard escalation trigger; no cross-slice bundle collision; no locked-contract or design-of-record amendment}`. The slice-team dispatches under delegated authority when all hold.
  - **Escalate back to master (CTO + guide + VP)** only when a condition fails — a scope/boundary deviation, any hard escalation trigger, a touch/amendment of a locked contract or the design-of-record, or a cross-slice collision. No standing blanket VP post-plan gate.
- **one slice at a time (at most two)**, per the operator's cadence; a slice's exit gate green before the next spins up.

### 2. Decomposition — 4 sequential sections, VERTICAL-SLICE-FIRST
De-risk *composition* first (one working relay before any taxonomy thickens):
- **S1 — the thin end-to-end relay.** Guide **m-7**. The path `mint → connect → render-MVP-form → submit → stamp → validate → lineage → append(crash-atomic) → project → deliver → gate-outbox`, + its adversarial/crash/replay fixtures (S1-scoped). **The de-risk slice — the only one fully dispatched now.**
- **S2 — thicken store/engine.** Guide **m-1** (+ m-7 on the loop). Full recovery (phases 0–4), durable FIFO, GC/genesis, the **owed-item-as-typed-record projection** (materialize-first, *recorded* items only).
- **S3 — thicken forms/lineage.** Guide **m-2**. Full FieldSpec registry + the 62-check linter refactor + **the full dissolved-linter replay gate** ("~33 checks dissolve", **F1**) + `schema_version` + migrator registry.
- **S4 — consumer schema slices + §C4 fixtures.** Guides **m-3 / m-4 / m-5 / m-6**. Step-gated-off observe fields (m-3), routing-record schema (m-4), archetype atoms + ceilings (m-5), gate/ODB fields (m-6) + the owed §C4 fixtures.

Sequence = dependency + de-risk. **I fully dispatch only S1 now**; S2–S4 sequenced, dispatched as each predecessor's exit gate goes green.

### 3. Slice-1 dispatch spec (ready to cut on GO)
- **Guide:** m-7. **Slice-team:** `s1`.
- **Scope (in):** engine skeleton (serialized commit loop + crash-atomic rename commit) · channel/identity (connect = channel-stamped FROM) · interface guardrail (seats reach only `submit`/`project`/`read`; raw store/config paths absent from the tool surface) + I-PH path hygiene · a **tiny MVP FieldSpec** · minimal validate (required-set / enum / seat-scope) · minimal lineage (parent-edge presence + validity) · terminal append `{accepted|rejected|held}` · projection rebuild + `read` · deliver via `project` · one gate producing a local outbox/ODB item · a **minimal dissolved-linter replay** over the MVP FieldSpec's covered historical failures (**F1**).
- **Scope (out — deferred):** full FieldSpec registry + linter refactor + the **full** dissolved-linter replay (S3) · full recovery phases + FIFO durability + GC + the owed-item projection (S2) · all consumer schema fields (S4) · observe-as-send (Step-2) · routing execution (Step-3).
- **Consumed, frozen contracts + the fidelity edge (F3):** m-1 store API (`submit`/`project`/`read`, append-only, sole-writer) + the m-2 FieldSpec envelope — locked. The S1 boundary/plan-review contract names **m-1.implementer** (store-API usage fidelity) and **m-2.implementer** (FieldSpec-envelope usage fidelity) as **required reviewers of the slice's consuming surface**; each returns a fidelity-review approve **before** the s1 slice-team's `DISPATCH IMPL` is live. This reviews *usage*, never reopens the locked design; a fidelity finding blocks the slice's dispatch until the slice's usage is corrected — the contract itself is not changed.
- **Deliverable:** the first working conductor relay in `frank/` (first commits) + the S1-scoped exit-gate fixtures green.
- **Framing (honesty — code + any doc):** Step-1 = **provenance + transport, not verified work** (observe is Step-2; "done" here is `self_reported`). Only the **serialized-loop double-accept kill** (and, with a constrained grammar, **R2**) are *operationally live* in S1; the others are *recorded*, not enforced, until their later layers. Confusion-resistance is **tool-mediated** (removes affordance, not access; D5 = shell-routed confusion survives).

### 4. Slice-1 exit gate — the S1-SCOPED hardened gate (binding acceptance; F1)
Baseline (roadmap): accepted only through the conductor · `FROM` system-stamped · form/lint before delivery · gate produces a local outbox item.
Promoted:
- **Adversarial:** forged `FROM` rejected · forbidden enum absent-then-rejected · invalid parent rejected · **duplicate-sibling double-accept killed** (the serialized-loop guarantee, actually exercised).
- **Crash-atomicity (crash it for real):** `kill -9` mid-commit and mid-delivery → exactly-once outcome + re-issued wake · crash after intake-fsync (no lost intake) · crash before/after the atomic rename (presence = committed) · corrupt-projection rebuild from canonical · replayed intake-id (no double-emission).
- **Dissolved-linter replay (S1-minimal, F1):** run the historical upstream-protocol lint failures **the MVP FieldSpec covers** through the MVP validator; confirm each is caught-or-genuinely-obsolete. The **full** "~33 checks dissolve" replay is an **S3** gate, not S1's.
- **Path-hygiene (I-PH):** no seat-facing output — every bounce/error included — contains a canonical store path.
- **Liveness:** inbox = durable truth, pipe `write()` = nudge; a busy/dead seat still receives via `project()` on reconnect; a lost wake never sleeps a parked lane forever.
- **park/wake:** a gated lane parks (consumes nothing) + wakes on the operator's verdict.

### 5. Owed §C4 carries — mapped to sections (so none silently drops)
- **code-layer interface-guardrail enforcement** + **I-PH path-hygiene fixture** → **S1**.
- ③ known-A / RAISE-ONLY NF fixture → **S1** (guardrail-adjacent) + **S4** (schema surface).
- R2 `gate_referenceable`-per-column negative fixtures → **S3** (with the FieldSpec registry); the constrained R2 grammar partially in **S1**.
- `GRILL_REQUIRED` FieldSpec row → **S3**.
- (optional) `routing_escalation` §J2 member → **S4** (m-4 routing-record slice).
- ⑤ ODB model-name egress fixture → **S4** (m-6 gate/ODB) — dormant, egress-shaped scan.

Per the **materialize-first PLAN rule**: each carry is first written as a typed owed-item record `{owner, source, target surface, disposition path}`; S2's projection guards only *recorded* items.

### 6. Slice-1 boot addressing (on GO) + the gate ask
- **Boot dispatch (cut on GO):** `TO s1.orchestrator-planner, m-7.planner`; `CC master.orchestrator-reviewer, operator, m-1.implementer, m-2.implementer` (**F3** — the fidelity reviewers wired in from the boot). The `s1` slice-team then formally requests the fidelity review (`TO m-1.implementer`, `TO m-2.implementer`) as a precondition to its delegated dispatch.
- **The gate — GO / amend on:** (a) the 4-section sequence · (b) the team-per-slice model + `s<N>` conventions · (c) **m-7-leads-S1** · (d) the **S1-scoped hardened exit gate** as binding acceptance · (e) the **frozen-contracts** stance · (f) the honesty framing riding into code · **(g)** the **F2 risk-scaled plan-gate depth** (S1/bootstrap = guide+VP; S2+ = pair review + conditioned delegation) · **(h)** the **F3 m-1/m-2 fidelity edge**.

On GO **and** operator ratification of the build-execution model, I cut the `s1` team-boot; nothing spins up before then.

### 7. Operator-judgment items
- **broad-scope / product-structure:** the whole build's decomposition + the team-per-slice recursion — this relay *is* the escalation; **operator ratifies the build-execution model** (VP-flagged).
- **ambiguous product semantics:** none new (design locked).
- **residual risk accepted:** D5 shell-routed confusion (tool-mediated, not malice-proof) — already operator-accepted; restated so S1 code does not over-claim.

## Verification
- `python3 ~/.claude/skills/tools/relay-lint.py <this file>` — run below.
- Amendments cross-checked against `step1-plan/RECONCILE-orchestrator-reviewer-20260703-125139.md` (F1/F2/F3), `STEP-1-KICKOFF.md` (exit-gate section updated), protocol.md (delegated-dispatch conditions, orchestrator-review visibility model).
- Pointers: `master/STEP-1-KICKOFF.md`, `master/ARCHITECTURE.md` §C4 + §C4.3/I-PH, `ROADMAP.md` Step-1, `frank/` (empty repo, branch `main`, 0 commits).

ACTIONS_GIT_REF: wrote this r2 relay + updated the `STEP-1-KICKOFF.md` exit-gate section (F1) + an `INDEX.md` row; cwd is not a git repo (docs workspace) so no sha — files on disk. No code, no dispatch token, no slice-team spun up.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` is the code repo, empty (branch `main`, 0 commits, `git status --short` empty).
Next requested action: operator ratifies the build-execution model; VP re-gates this r2 (RECONCILE under `step1-plan/`); on GO + ratification the CTO cuts the `s1` Slice-1 team-boot; first code lands in `frank/`.
