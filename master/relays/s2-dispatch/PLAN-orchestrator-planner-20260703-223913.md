## PLAN — Slice-2 build dispatch (r2 — supersedes `…-223146`; new team per operator; dispatch-root lint-clean per VP F-S2-1)

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s2-dispatch
PARENT_DISPATCH_ID: step1-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no for this dispatch — operator directed S2 (2026-07-03); S2 runs on F2 (pair plan-review + conditioned delegation); the S2-close sign-off is the operator's, exercised separately
DELEGATED_DISPATCH_AUTHORITY: yes — conditions below (F2, non-bootstrap)
GRILL_REQUIRED: no
IN_REPLY_TO: s2-dispatch/RECONCILE-orchestrator-reviewer-20260703-223502.md
FROM: master.orchestrator-planner
TO: s2.orchestrator-planner, m-7.planner
CC: master.orchestrator-reviewer, operator, m-1.planner, m-1.implementer, m-7.implementer
SUBJECT: S2 = thicken store/engine (full recovery phases 0–4 + durable FIFO + GC/genesis) + the owed-item-as-typed-record projection; NEW slice-team; guide m-7 (refines the decomposition's m-1); F2 conditioned delegation; OI-S1-F11-SWEEP is the projection's first customer

**Supersedes `s2-dispatch/PLAN-orchestrator-planner-20260703-223146.md`** — two folds: (1) **new team** (operator: a new sprint should be a new team — the truer reading of the build-execution model, "each slice = its own orchestrator-team instance"); (2) **dispatch-root lint-clean** (VP F-S2-1 `…-223502`: the prior copy tripped the root-mode merge-claim detector on merge-form prose; reworded here, not waived — the S1 waiver was scoped to the S1 trail). VP substantive review otherwise all-green; its three watchpoints are folded below.

**What this is.** The master dispatch for **Slice-2**, executing the VP-approved + operator-ratified `step1-plan` decomposition. S1 closed 2026-07-03 (tag `s1-close`); S1 built + proved the relay spine at E2. **S2 hardens the engine and lands the first governance primitive** (the owed-item projection). Spec = `ARCHITECTURE.md` §C4 + the m-1/m-7 domain docs.

**Deferral on record (operator, 2026-07-03):** the **MCP live-adapter + fuller-FieldSpec "wire-it-up" work is deliberately deferred** — S1 is built-but-unwired (a real socket daemon, E2-proven, but no live agent session has connected; it speaks frank's own socket protocol + a minimal FieldSpec dialect). The operator chose S2 (engine-thickening, fully E2-testable now) over pulling live-integration forward, since there is no live testbed just now. The wire-it-up slice returns when a testbed exists.

### To the s2 slice-team — your charter
- **You are a NEW slice-team for S2** (fresh orchestrator + pair) — a new sprint, a new team. **Use `/orchestrator-planner`**; scaffold a new **`s2`** sprint via `sprint-doc-setup` in `frank/`. Your relays live in `frank/`.
- **Onboard first — you did NOT build S1.** The `frank/` code carries over (you build on the S1 baseline, tag `s1-close`), but it is new to you: **read the S1 source + the s1 sprint docs (`docs/sprints/2026-07-03-s1-slice-1/`) before planning.** Fresh adversarial eyes on that code are a benefit; unfamiliarity is the cost — m-7 (your guide, who guided S1) is your continuity into the S1 design.
- **Spec = read-only reference in cwd:** `ARCHITECTURE.md` §C4 (the engine) + the m-1/m-7 domain docs. Escalate spec problems to master — do not self-amend a locked design.
- **Build on `main`** (the S1 baseline).

### Guide + contract boundary (refines the decomposition; VP-concurred)
The decomposition listed S2's guide as **m-1**, but S2's live content — recovery phases 0–4, durable FIFO, GC/genesis, the owed-item projection (the generalized m-7 held/burn pattern) — is m-7's engine domain, and m-7 guided S1. So **m-7 is the primary guide.** **m-1's authority is explicit and unchanged (VP watchpoint):** m-1 retains authority over the **owed-item `record_kind`, the store layout, and store-API fidelity**, and is the **fidelity reviewer** for anything touching its locked store contract — m-7 guides the engine *implementation* but must **not** redefine the m-1 store contract. (VP CC'd; it concurred with this refinement in `…-223502`.)

### Slice-2 scope (IN)
- **Full crash-recovery — phases 0–4** (S1 shipped crash-atomic commit + representative recovery; S2 completes the phase machinery).
- **Durable FIFO** — the intake queue's full durability + ordering under crash/restart.
- **GC / genesis** — store genesis + garbage-collection of superseded/consumed records, crash-safe.
- **The owed-item-as-typed-record projection** — an owed-item `record_kind` (m-1 fidelity) + a `project()` over it, `open = owed-record with no disposition-record` (materialize-first; **guards recorded owed-items only — it does NOT make an unrecorded observation impossible to miss**, VP watchpoint). **`OI-S1-F11-SWEEP` is its first real customer** (materialized in the s1 ledger) — the projection must surface it as an open owed-item dispositioned to the S2 exit gate.

### Slice-2 scope (OUT — deferred; do not build; escalate before any delegated dispatch that touches these)
Full FieldSpec registry + 62-check linter refactor (**S3**) · the MCP live-adapter + protocol-envelope FieldSpec (**deferred, no testbed**) · observe-as-send (**Step 2**) · routing execution (**Step 3**) · all consumer schema — observe/routing/archetype/ODB (**S4**). If a task seems to need one, **escalate to master** — do not self-dispatch through it.

### Slice-2 exit gate (HARD acceptance)
- **Recovery phases 0–4** crash-tested (crash at each phase boundary → correct resumption; no lost/double intake or delivery).
- **Durable FIFO** — order + exactly-once preserved across crash/restart.
- **GC/genesis** crash-safe (genesis idempotent; GC never drops a live record; crash mid-GC recovers).
- **The owed-item projection** — `open = owed-record with no disposition-record` holds; **OI-S1-F11-SWEEP surfaces as open**; a dispositioned owed-item drops from the open set; materialize-first honored.
- **F9/F11 re-run under the new recovery machinery** (guide's standing note) — the crash matrix passes against full phase-0–4 recovery, and **OI-S1-F11-SWEEP's full class×point sweep is discharged here** (the S1 owed item closes at this gate).
- **No regression** of the S1 gate invariants (serialized-loop kill, crash-atomicity, I-PH, enum byte-exactness, guardrail).

### Plan-gate (F2 — non-bootstrap; conditioned delegation)
Produce your **S2 PLAN** (in `frank/`); it runs on the **normal pair Implementer plan-review + delegated dispatch** — no bootstrap guide+VP plan-gate. Your `DISPATCH IMPL` is authorized under: **{Implementer plan-review = approve; no deviation from this dispatch's scope/boundary; no hard escalation trigger; no cross-slice collision; no locked-contract or design-of-record amendment}**. If **any** fails — a scope/boundary deviation, a hard trigger, a touch/amendment of a locked contract or the design-of-record, or a cross-slice collision — **escalate back to master (CTO + m-7 guide + VP)**; do not self-dispatch through it. (m-7 guides as domain owner; m-1 fidelity-reviews store-touches before their dispatch.)

### Framing (honesty — code + any doc)
Still **provenance + transport, not verified work** (observe is Step 2). S2 adds **durability hardening + the first governance primitive** (the owed-item projection) — real value, but it does **not** make an *unrecorded* observation impossible to miss (materialize-first). Confusion-resistance stays tool-mediated; D5 residual. Do not over-claim.

### Deliverable format
The thickened engine on `main` in `frank/` (commits) + the S2 exit-gate fixtures green (E2) + the OI-S1-F11-SWEEP full sweep discharged; your build relays in `frank/`; a SITREP back to master at the S2 exit gate.

### Operator-judgment items
- **residual risk (accepted):** D5 shell-routed confusion (restated).
- **escalation posture:** no auth/data/migration (greenfield store); the F2 delegated conditions gate any code; the S2-close sign-off is the operator's, exercised separately.
- **deferred (recorded):** the MCP live-adapter + fuller FieldSpec — returns with a testbed.

### Not authorized by this relay
No S2-close authority, no scope expansion beyond the IN list, no locked-design amendment. `DISPATCH IMPL` is delegated **only** under the conditions above; a failed condition escalates to master.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py <this file>` + `--relay-root master/relays/s2-dispatch` — both run below (target: exact-file AND dispatch-root clean, F-S2-1 closed).
- Executes `step1-plan` (VP-approved decomposition) + the S1 close (verified in `master/RECONCILE.md` § S1). Guide-refinement + MCP-deferral VP-reviewed (`…-223502`, all-green ex-lint); watchpoints folded.
- Pointers: `master/ARCHITECTURE.md` §C4, `master/STEP-1-KICKOFF.md`, `ROADMAP.md`, `frank/` (baseline tag `s1-close`).

ACTIONS_GIT_REF: wrote this superseding s2-dispatch relay + an `INDEX.md` row; cwd is not a git repo (docs workspace) so no sha — files on disk. No code, no dispatch token, no `frank/` write.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` is on branch `main`, clean tree, baseline tag `s1-close`.
Next requested action: operator relays this to a fresh `s2.orchestrator-planner` session + m-7.planner; s2 scaffolds a new `s2` sprint via sprint-doc-setup in `frank/`, onboards to the S1 code, plans S2, and dispatches under the F2 delegated conditions (escalating to master only on a trigger); SITREP back to master at the S2 exit gate.
