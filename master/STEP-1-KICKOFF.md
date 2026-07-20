# Step-1 KICKOFF — `frank` conductor-core / "automated operator-relay" (PLAN-approach; pre-decomposition)

**Status: PLAN-APPROACH DRAFT (2026-07-03).** The durable starting point for the Step-1 PLAN — scope, build strategy, exit gate, and the external-review refinements folded in. The formal decomposition + VP-gated PLAN dispatch happens after the charter transition + compaction. Design-of-record: `ARCHITECTURE.md` (§C4 = the engine) + the 7 domain docs; sequencing: `ROADMAP.md` Step-1.

## Goal (one line)
**Remove the operator-as-transport.** **`frank`** (the conductor) becomes an MCP server that existing agent sessions (Claude Code / Codex) connect to; seats call `submit()` / receive via `project()` instead of the operator hand-relaying files. *(The Step-1 build lands in the `frank/` repo — formerly `pcode/`; the CLI verb is `frank`.)* Step-1 = the trusted courier's spine: **store + form + lineage**. It does **not** run its own agents (Step-3) and does **not** verify done-ness (Step-2).

## Framing (external-review honesty — GPT §3/§4)
- Step-1's pitch is **provenance + transport automation, NOT "verified work"** — observe-as-send is Step-2, so "done" in Step-1 is `self_reported` (honest fallback).
- Of the four sanctioned by-construction guarantees, **only the serialized-loop kill (and, with a constrained grammar, R2) are *operationally live* in Step-1**; observer-selected controls + authority-ceilings are *recorded* in Step-1 but their runtime lands with the archetype/runtime layers. Step-1 docs/marketing must not imply otherwise.
- **Confusion-resistance is tool-mediated** (ARCHITECTURE §C4.3, I-PH) — the guardrail removes affordance, not access; a confused shell-bearing seat can route around it (D5). Step-1 enforces the **path-hygiene invariant (I-PH)** so the conductor's own bounces/errors/projections never leak a store path.

## Build-execution model — team-per-slice, m-x-guided (operator-decided 2026-07-03)
**Each slice is implemented by its own orchestrator-team instance, guided by the domain's m-x leader.** Three layers — dogfooding the product's own nested-orchestrator-team vision (roadmap Step-5) as the build method itself:
- **master** (CTO + VP + m-1…m-7) — *governs*: decomposes, gates, integrates, owns the architecture-of-record.
- **m-x guide** — the master-team pair owning the slice's core domain; feeds the slice-team the locked design, answers domain questions, reviews output for contract-fidelity; does not type the code.
- **slice-team** — a *separate* orchestrator-team instance (own planner/reviewer + pairs) that plans + implements the slice against the frozen contracts.

Conventions: slice-team RUN_ID `s<N>` (Slice-1 = `s1`). **Slice-team relays live in `frank/`, governed by the `sprint-doc-setup` skill** — its `docs/sprints/<date>-<slice>/` tree + gitignored `.relays/<RUN_ID>/` substrate — so build relays live *with* the code, not in the governance trail (Cardinal #2). The master CTO's dispatch act (the design→build handoff) is recorded in `master/relays/`; the slice-team then scaffolds via `sprint-doc-setup` and hosts its own boot + PLAN + review + IMPL relays in `frank/`. **Plan-gate depth is risk-scaled (VP-amended F2):** the m-x guide + VP gate the slice-team's plan for **S1/bootstrap only**; **S2+** run on normal pair Implementer plan-review + conditioned delegated dispatch, escalating back to master only on named triggers (scope/boundary deviation · hard trigger · locked-contract or design-of-record amendment · cross-slice collision). One slice at a time (at most two). *(VP-approved in `step1-plan` r2 `…-125826`; operator-ratified 2026-07-03 with the `frank/` relay-location correction.)*

## Build strategy — VERTICAL SLICE FIRST (external-review, GPT rec)
**Do NOT build m-1, then m-2, then m-7 as separate castles.** Build the thinnest end-to-end path through *all* layers first, then thicken. This gets a working conductor end-to-end fast and de-risks the composition (the hard part), not the taxonomy.

**Slice 1 — the brutally small end-to-end relay** (the first section):
```
mint seat → connect seat (channel = identity) → render allowed submit schema (MVP FieldSpec)
  → submit relay → conductor stamps FROM/ROLE/PARENT/id/time
  → validate (minimal required-set / enum / seat-scope) → lineage check (minimal parent-edge)
  → append terminal {accepted|rejected|held} (crash-atomic rename) → rebuild/read projection
  → deliver to target inbox (project) → a gate produces a local ODB/outbox item
```
Everything in Slice 1 is the **minimum** to make one relay flow: a **tiny MVP FieldSpec** (just the fields the slice needs — the full FieldSpec registry + the 62-check linter refactor come *after* the first relay works), a minimal lineage check, the crash-atomic commit, the interface guardrail + I-PH path hygiene.

**Then thicken (subsequent sections):** full FieldSpec registry + linter refactor · full crash-recovery (phases 0–4) + durable FIFO · the **owed-item-as-typed-record projection** (promote early — see below) · the consumer schema slices (m-3 observe fields step-gated off, m-4 routing-record schema, m-5 archetype atoms + ceilings, m-6 gate/ODB fields) · the §C4 owed fixtures.

## Promote into Step-1 early (both reviews + our own conclusion)
**The owed-item-as-typed-record projection.** An owed-item `record_kind` + a `project()` over it, `open = owed-record with no disposition-record` (the m-7 `held`/burn pattern generalized). It is the smallest, most immediately-valuable governance primitive that proves the architecture *before* observe-as-send exists. **Scope (VP-corrected):** it makes silent drop impossible for a ***recorded*** owed-item (`open = owed-record with no disposition-record`) — it does **NOT** make an *unrecorded* observation impossible to miss; materializing the record is an intake/triage step the projection does not replace. (Unqualified "impossible-by-projection" is the same overclaim shape this package exists to scrub — noted, and fixed.) **PLAN rule (materialize-first):** any review finding / carry / external-review sharpening meant to be *guarded* by the projection must first be written as a typed owed-item record — `{owner, source relay/file, target surface, disposition path}`; the projection guards only *after* that record exists. Build it in the "thicken" phase, right after Slice 1.

## Exit gate — HARDENED (external-review; the current 4 criteria are too happy-path)
**Baseline (roadmap):** (1) a relay is accepted **only** through the conductor; (2) `FROM` is system-stamped, not lane-supplied; (3) form/lint validation runs **before** delivery; (4) a gate produces a local outbox item.

**Promoted into the exit gate (adversarial + crash + replay):**
- **Adversarial:** forged `FROM` rejected · forbidden enum absent-then-rejected · invalid parent rejected · **duplicate-sibling double-accept killed** (the serialized-loop guarantee, actually exercised).
- **Crash-atomicity (the headline claim — actually crash it):** `kill -9` **mid-commit** and **mid-delivery** → verify **exactly-once** outcome AND a **re-issued wake** (Fable liveness) · crash after intake-fsync (no lost intake) · crash before/after the atomic rename (presence=committed holds) · **corrupt-projection rebuild** from canonical · **replayed intake-id** (no double-emission).
- **Dissolved-linter replay (Fable — argument→evidence):** run historical upstream lint failures through form-validation; confirm each is **caught-or-genuinely-obsolete**. **Slice-scoped (VP RECONCILE 2026-07-03):** the *full* "~33 checks dissolve" validation is an **S3 gate** — it needs the 62-check linter refactor + full FieldSpec registry, which land in S3. **S1 carries a minimal replay** over only the historical failures its **MVP FieldSpec** covers — proving the replay harness runs and the MVP-covered subset is caught-or-obsolete. So S1's exit gate is the **S1-scoped** hardened gate, not the full one.
- **Path-hygiene (I-PH):** no seat-facing output — every bounce/error included — contains a canonical store path.
- **Liveness:** the **inbox is the durable truth**, the pipe `write()` is a **nudge**; a busy/dead seat at write time still receives via `project()` on reconnect; a lost wake never sleeps a parked lane forever.
- **park/wake:** a gated lane parks (consumes nothing) + wakes on the operator's verdict.

## Design refinements folded (external-review — small, cheap, do at build)
- **Liveness contract stated:** inbox durable / pipe = nudge / recovery re-issues nudges.
- **Lineage bounce ergonomics:** distinguish "parent unknown, possibly in-flight — recompose + resubmit" from "parent invalid — dead edge" in the **reason field** (not a new state), so seats don't retry-loop on dead edges or abandon live ones.
- **FieldSpec drift:** a seat holding a stale rendered form gets a bounce carrying **"re-render."**
- **Explicit non-goals (write them down):** inter-seat confidentiality is **unclaimable** under same-uid/D5 (an explicit non-claim, not a silent gap); relay **content** passes the courier **unscanned** (the egress scan is dormant + egress-shaped) — the right non-goal for same-trust-domain seats, but "governance-first courier" invites the wrong assumption, so state it.

## Owed Step-1-build carries (from the §C4 ledger — inherited by this build)
③ known-A / RAISE-ONLY NF fixture · ⑤ ODB model-name egress fixture · R2 `gate_referenceable`-per-column negative fixtures · `GRILL_REQUIRED` FieldSpec row · (optional) `routing_escalation` §J2 member · **the code-layer interface-guardrail enforcement** (the one genuinely new build item) · **I-PH path-hygiene fixture** (external-review).

## Proposed section decomposition (for the PLAN — revised to vertical-slice)
**Renumber of record (operator, 2026-07-05):** slice RUN_IDs count *team instances*. After S1–S3 closed, the operator elected the **wire-up as s4** (the MCP shim + live sessions + the §7 config-change record — the slice that ends the operator-as-transport; chartered as the operator's fork in the s3-dispatch). **Section 4 below (consumer schema slices + fixtures) is therefore s5**, built *with s4 in use* — over the wired conductor, its registry rows landing as real §7 config-change records, generating the first usage data. Step-1 ends after s5 + the roadmap step-exit test.
1. **Slice-1 end-to-end** (thin path above; engine skeleton + channel/identity + guardrail + I-PH + MVP FieldSpec + minimal lineage + crash-atomic commit + local outbox) → the first working relay + the adversarial/crash fixtures around *it*.
2. **Thicken the store/engine:** full recovery (phases 0–4) + durable FIFO + GC/genesis + the owed-item projection.
3. **Thicken forms/lineage:** full FieldSpec registry + the 62-check linter refactor (with the dissolved-linter replay) + `schema_version` + migrator registry.
4. **Consumer schema slices + fixtures:** m-3 (step-gated observe fields) · m-4 (routing-record schema) · m-5 (archetype atoms + ceilings) · m-6 (gate/ODB fields) + the §C4 owed fixtures.

*(Order is dependency + de-risk: one working relay before any taxonomy is thickened — the "don't become ceremonially complete before the first conductor exists" rule.)*
