## PLAN — Slice-3 build dispatch (master → s3 slice-team; thicken forms/lineage: the full FieldSpec registry + the linter dissolution)

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s3-dispatch
PARENT_DISPATCH_ID: step1-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no for this dispatch — operator directed S3 (2026-07-04, post-S2-close); S3 runs on F2 (pair plan-review + conditioned delegation); the S3-close sign-off is the operator's, exercised separately
DELEGATED_DISPATCH_AUTHORITY: yes — conditions below (F2, non-bootstrap)
GRILL_REQUIRED: no
IN_REPLY_TO: .relays/s2/s2-exit-gate/SITREP-orchestrator-planner-20260704-152000.md
FROM: master.orchestrator-planner
TO: s3.orchestrator-planner, m-2.planner
CC: master.orchestrator-reviewer, operator, m-2.implementer, m-7.planner, m-1.implementer
SUBJECT: S3 = thicken forms/lineage — full FieldSpec registry + the 62-check linter dissolution (with the FULL dissolved-linter replay) + schema_version/migrator registry + the R2/GRILL_REQUIRED owed carries; NEW slice-team; guide m-2; F2 conditioned delegation

**What this is.** The master dispatch for **Slice-3**, executing the VP-approved + operator-ratified `step1-plan` decomposition. **VP pre-handoff review: APPROVE** (`s3-dispatch/RECONCILE-orchestrator-reviewer-20260704-151823`); its four watchpoints are folded inline below (marked **[VP-W]**). S1 (spine) and S2 (engine: recovery 0–4, durable FIFO, GC/genesis, the owed-item projection) are **CLOSED, complete at E2** — baseline **`main`, tag `s2-close`**, zero owed items riding in. **S3 makes frank speak the real protocol:** the full FieldSpec registry replaces the S1 MVP dialect, and the upstream linter dissolves into form-validation — the precondition for any future live wire-up (real relays cannot validate through frank until this lands). Spec = `ARCHITECTURE.md` §C4 + the **m-2** domain doc (primary) + m-1/m-7 as consumed contracts.

### To the s3 slice-team — your charter
- **You are a NEW slice-team for S3** (fresh orchestrator + pair; new sprint = new team). **Use `/orchestrator-planner`**; scaffold a new **`s3`** sprint via `sprint-doc-setup` in `frank/`. Your relays live in `frank/`.
- **Onboard first — you built neither S1 nor S2.** Read the S1+S2 source and both sprint ledgers (`docs/sprints/2026-07-03-s1-slice-1/`, `…-s2-slice-2/`) before planning. Fresh adversarial eyes are a benefit (S2's audits found 2 latent S1 races — the precedent to live up to); **m-2 (your guide) is the domain continuity**, m-7 knows the engine.
- **Spec = read-only reference in cwd:** the **m-2 design-of-record** (FieldSpec registry, field-ownership model, linter→form-validation dissolution, dynamic required-set) + `ARCHITECTURE.md` §C4. Escalate spec problems to master — do not self-amend a locked design.
- **Build on `main`** (post-`s2-close`), on a branch; the close-time integration is the operator's separate gate (the S2 layered-authority precedent).

### Guide + contract boundaries
- **m-2 is the primary guide** — this slice IS m-2's domain (the registry, the ownership model, the dissolution, the dynamic required-set).
- **m-7 consulted** on the engine/config seams: the registry rides the **trusted config** (per-domain-authored sections, single top-level digest, loaded once at trusted startup — §C4.1); the engine *executes* the m-2 contract, it does not re-own it. Any change to the config-load/digest mechanics is an m-7 seam — consult, don't improvise.
- **m-1.implementer fidelity** for any store-API touches (the S2 precedent; expected to be light — S3 is mostly above the store).
- **[VP-W] Lineage movement is an m-1 fidelity trigger even inside m-2-owned modules:** if the lineage-check re-home changes `PARENT`, `parent_picker`, candidate-set derivation, system-filled lineage fields, or store *query* semantics, **m-1 fidelity review is required before the s3 planner treats that plan as delegated-dispatch eligible** — the module boundary does not decide the contract boundary.

### Slice-3 scope (IN)
- **Full FieldSpec registry** — the complete typed-envelope field catalog (field · owner ∈ {system, seat-scoped-enum, agent-enum-pick, free-text} · type · required-when · enum), replacing the S1 MVP dialect; registry-driven render + validate; **fill-time authority** (forbidden options absent from the rendered form, not rejected after).
- **The 62-check linter dissolution** — the upstream `relay-lint`'s checks re-homed: each check either **dissolves into form-validation/lineage** (structurally impossible to violate through the form) or is **retained as an explicit post-submit check** or is **genuinely obsolete** — a per-check disposition table, no silent drops.
- **The FULL dissolved-linter replay (the S1-deferred F1 gate):** run the historical upstream lint-failure corpus through the new validation; every failure **caught-or-genuinely-obsolete**, validating the "~33 checks dissolve" claim by execution.
- **`schema_version` + the migrator registry** — the version stamp in its `system_only` envelope home (already stamped since S1/s2-amend) + the migration mechanism (registry + apply path + a fixture proving a v(n)→v(n+1) walk). Zero real migrators exist yet — the mechanism ships first (the public-release intent: versioned schemas + migration procedure from day one).
- **Owed §C4 carries landing here:** **R2 `gate_referenceable`-per-column negative fixtures** (no model-derived predicate can enter a gate through any registry column — proven by negative fixtures, per the locked R2 grammar) · the **`GRILL_REQUIRED` FieldSpec row**.
- **Re-render/drift** — a seat holding a stale rendered form gets a bounce carrying "re-render" (S1 seeded this; S3 must keep it correct under a *changing* registry).

### Slice-3 scope (OUT — deferred; escalate before any delegated dispatch that touches these)
The MCP live-adapter / wire-up (**next after S3, operator's call**) · observe-as-send + evidence fields (**Step 2**) · routing execution (**Step 3**) · consumer schema slices — observe/routing/archetype/ODB field *content* (**S4**) · any TUI/runtime work. **[VP-W] The S4 line, precisely:** S3 **may** define registry slots, types, ownership classes, validation mechanics, and referenceability; S3 **may not** choose sibling-owned consumer *semantics* (what an observe/routing/archetype/ODB field means or which values it takes) — expression capacity in, content out. If a task seems to need an OUT item, **escalate to master**.

### Slice-3 exit gate (HARD acceptance)
- **Registry live end-to-end:** a real relay in the team's actual header vocabulary (ROLE/PHASE/AUTHORITY/DISPATCH_ID/lineage/…) renders, validates, commits, projects through frank — the S1 thin path running the *full* protocol envelope.
- **Fill-time authority proven by negatives:** forbidden enum members are **absent from the rendered form** per seat-scope; a forged/out-of-scope field submission bounces pre-append.
- **The per-check disposition table complete** (all 62: dissolved / retained / obsolete — each with its fixture or rationale) + **the FULL replay green**: the historical corpus caught-or-genuinely-obsolete, no silent drops. **[VP-W] The obsolete-adjudication rule:** every *obsolete* disposition must name the **concrete vanished surface, legacy-only path, or replaced invariant** that makes the old check impossible or irrelevant — evidence, not a label; an obsolete disposition that rests on a **design-of-record change** (not a vanished implementation surface) **escalates to master before S3 close**.
- **R2 negatives green** (`gate_referenceable`-per-column) · **GRILL_REQUIRED row present + rendered**.
- **`schema_version` + migrators:** a fixture-proven v(n)→v(n+1) migration walk; **[VP-W] plus a fixtured refusal/bounce leg** for unknown/future, unversioned, and mismatched records (bounce/re-render, never silent coercion). A backward/downgrade migrator is **not** required — introducing one would be a scope expansion (escalate).
- **No regression:** S1+S2 suites green; enum byte-exact; I-PH on every new surface (registry errors + bounce text included); the owed-item projection + recovery/FIFO/GC untouched-and-green.
- **materialize-first, now with teeth:** any S3 finding meant to be *guarded* is first a typed owed-item record — **the live S2 mechanism exists; use it where practical** (the OI-S1-F11-SWEEP precedent — govern the build through the product when the cost is low).

### Plan-gate (F2 — non-bootstrap; conditioned delegation)
Produce your **S3 PLAN** (in `frank/`); the **pair Implementer plan-review** is the plan gate. Your `DISPATCH IMPL` is delegated **only** under: **{Implementer plan-review = approve; no deviation from this dispatch's scope/boundary; no hard escalation trigger; no cross-slice collision; no locked-contract or design-of-record amendment}**. Any failure — including any OUT-item touch — **escalates back to master (CTO + m-2 guide + VP)**; do not self-dispatch through it. (m-2 guides as domain owner; m-7 consulted on config seams; m-1 fidelity on store touches.)

### Framing (honesty — code + any doc)
Still **provenance + transport, not verified work** (observe is Step 2). S3 adds the *form system* — determinism + fill-time authority — not evidence or done-ness. The linter dissolution claim is proven **by the executed replay**, not asserted. Confusion-resistance stays tool-mediated; D5 residual. Do not over-claim.

### Deliverable format
The full form system on a branch in `frank/` + the S3 exit-gate fixtures green (E2) + the disposition table + the replay report; your build relays in `frank/`; a SITREP back to master at the S3 exit gate; the close-time integration is the operator's separate gate.

### Operator-judgment items
- **residual risk (accepted):** D5 shell-routed confusion (restated).
- **escalation posture:** greenfield above the store; F2 conditions gate any code; the S3-close sign-off is the operator's, exercised separately.
- **sequencing note (recorded):** S3 is the prerequisite that makes the deferred MCP wire-up *meaningful* — after S3, frank validates the team's real protocol; the wire-up slice becomes "connect the real thing."

### Not authorized by this relay
No S3-close authority, no scope expansion beyond the IN list, no locked-design amendment. `DISPATCH IMPL` is delegated **only** under the conditions above; a failed condition escalates to master.

## Verification
- `python3 ~/.claude/skills/tools/relay-lint.py <this file>` + `--relay-root .relays/s3/s3-dispatch` — both run below (target: clean both modes).
- Executes `step1-plan` (VP-approved decomposition, S3 = forms/lineage, guide m-2) + the S2 close (`master-docs/master/RECONCILE.md` § S2; tag `s2-close`; zero owed items riding in).
- Pointers: the m-2 design-of-record, `master-docs/master/ARCHITECTURE.md` §C4 (+ the §C4 owed-carry ledger rows this slice discharges), `master-docs/master/STEP-1-KICKOFF.md`, `frank/` (baseline `main`, tag `s2-close`).

ACTIONS_GIT_REF: wrote this s3-dispatch relay + the s3 boot + `INDEX.md` rows + a dashboard update; cwd is not a git repo (docs workspace) so no sha — files on disk. No code, no dispatch token, no `frank/` write.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` is on `main`, clean tree, baseline tag `s2-close`.
Next requested action: operator relays the s3 boot to a fresh `s3.orchestrator-planner` session, then this dispatch; s3 onboards to the S1+S2 code, scaffolds `s3`, plans S3, and dispatches under the F2 conditions; SITREP back to master at the S3 exit gate.
