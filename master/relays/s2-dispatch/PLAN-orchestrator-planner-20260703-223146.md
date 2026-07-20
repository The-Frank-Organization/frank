## PLAN — Slice-2 build dispatch [SUPERSEDED] (master → s2 slice-team; thicken store/engine + the owed-item projection)

**SUPERSEDED 2026-07-03 by `s2-dispatch/PLAN-orchestrator-planner-20260703-223913.md`** (new team per operator; dispatch-root lint fix per VP F-S2-1 `…-223502`). This draft never took effect; the close-gate wording below is neutralized so the dispatch root stays lint-clean.

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
IN_REPLY_TO: frank/.relays/s1/s1-merge-gate/RECONCILE-orchestrator-planner-20260703-220652.md
FROM: master.orchestrator-planner
TO: s2.orchestrator-planner, m-7.planner
CC: master.orchestrator-reviewer, operator, m-1.planner, m-1.implementer, m-7.implementer
SUBJECT: S2 = thicken store/engine (full recovery phases 0–4 + durable FIFO + GC/genesis) + the owed-item-as-typed-record projection; guide m-7 (refines the decomposition's m-1); F2 conditioned delegation; OI-S1-F11-SWEEP is the projection's first customer

**What this is.** The master dispatch for **Slice-2**, executing the VP-approved + operator-ratified `step1-plan` decomposition (S1 closed 2026-07-03, tag `s1-close`). S1 built + proved the relay spine at E2; **S2 hardens the engine and lands the first governance primitive** (the owed-item projection). Spec = `ARCHITECTURE.md` §C4 + the m-1/m-7 domain docs.

**Deferral on record (operator, 2026-07-03):** the **MCP live-adapter + fuller-FieldSpec "wire-it-up" work is deliberately deferred** — S1 is built-but-unwired (a real socket daemon, E2-proven, but no live agent session has connected; it speaks frank's own socket protocol + a minimal FieldSpec dialect). The operator chose S2 (engine-thickening, fully E2-testable now) over pulling the live-integration forward, since there's no live testbed to hand just now. The wire-it-up slice returns when a testbed exists.

### To the s2 slice-team — your charter
- **Team = the s1 team continued** (same seats; the `frank/` code + context carry over). **Use `/orchestrator-planner`**; scaffold a new **`s2`** sprint via `sprint-doc-setup` in `frank/` (or continue the s1 sprint tree — your call). Your relays live in `frank/`.
- **Spec = read-only reference in cwd:** `ARCHITECTURE.md` §C4 (the engine) + the m-1/m-7 domain docs. Escalate spec problems to master — do not self-amend a locked design.
- **Build on `main`** (S1 baseline `f0dcb85`).
- **Your m-x guide = m-7** (conductor-core) — see the guide note below.

### Guide (refines the decomposition — flagged to the VP)
The decomposition listed S2's guide as **m-1**. S2's actual content — **recovery phases 0–4, durable FIFO, GC/genesis, and the owed-item projection (the generalized m-7 held/burn pattern)** — is m-7's engine domain, and m-7 built S1. So: **m-7 is the primary guide**; **m-1 is consulted** on the store-record/API surface and is the **fidelity reviewer** for anything that touches its locked store contract (the new owed-item `record_kind` especially — F3 pattern, scoped to store-touches). This changes a VP-approved item — the **VP is CC'd for visibility** (not a re-gate); VP, flag if you disagree.

### Slice-2 scope (IN)
- **Full crash-recovery — phases 0–4** (S1 shipped the crash-atomic commit + a representative recovery; S2 completes the phase machinery).
- **Durable FIFO** — the intake queue's full durability + ordering guarantees under crash/restart.
- **GC / genesis** — store genesis + garbage-collection of superseded/consumed records, crash-safe.
- **The owed-item-as-typed-record projection** — an owed-item `record_kind` + a `project()` over it, `open = owed-record with no disposition-record` (materialize-first; guards only *recorded* owed-items). **`OI-S1-F11-SWEEP` is its first real customer** (already materialized in the s1 ledger) — the projection must surface it as an open owed-item dispositioned to the S2 exit gate.

### Slice-2 scope (OUT — deferred; do not build)
Full FieldSpec registry + 62-check linter refactor (**S3**) · all consumer schema — observe/routing/archetype/ODB (**S4**) · observe-as-send (**Step 2**) · routing execution (**Step 3**) · **the MCP live-adapter + protocol-envelope FieldSpec** (deferred, no testbed — see above). If a task seems to need an OUT item, **escalate to master**.

### Slice-2 exit gate (HARD acceptance)
- **Recovery phases 0–4** crash-tested (crash at each phase boundary → correct resumption; no lost/double intake or delivery).
- **Durable FIFO** — order + at-least-once/exactly-once preserved across crash/restart.
- **GC/genesis** crash-safe (genesis idempotent; GC never drops a live record; crash mid-GC recovers).
- **The owed-item projection** — `open = owed-record with no disposition-record` holds; **OI-S1-F11-SWEEP surfaces as open**; a dispositioned owed-item drops from the open set; **materialize-first** honored (unrecorded observations are not claimed as guarded).
- **F9/F11 re-run under the new recovery machinery** (the guide's standing note) — the crash matrix passes against the full phase-0–4 recovery, and **`OI-S1-F11-SWEEP`'s full class×point sweep is discharged** here (the S1 owed item closes at this gate).
- **No regression** of the S1-scoped gate invariants (serialized-loop kill, crash-atomicity, I-PH, enum byte-exactness, guardrail).

### Plan-gate (F2 — non-bootstrap; conditioned delegation)
Produce your **S2 PLAN** (in `frank/`); it runs on the **normal pair Implementer plan-review + delegated dispatch** — no mandatory guide+VP plan-gate this time. Your `DISPATCH IMPL` is authorized under the delegated conditions: **{Implementer plan-review = approve; no deviation from this dispatch's scope/boundary; no hard escalation trigger; no cross-slice collision; no locked-contract or design-of-record amendment}**. If **any** condition fails — a scope/boundary deviation, a hard trigger, a touch/amendment of a locked contract or the design-of-record, or a cross-slice collision — **escalate back to master (CTO + m-7 guide + VP)**; do not self-dispatch through it. (m-7 as guide will naturally review as the domain owner; m-1 fidelity-reviews store-touches before their dispatch.)

### Framing (honesty — code + any doc)
Still **provenance + transport, not verified work** (observe is Step 2). S2 adds **durability hardening + the first governance primitive** (the owed-item projection) — a real, valuable step, but it does **not** make an *unrecorded* observation impossible to miss (materialize-first). Confusion-resistance stays tool-mediated; D5 residual. Do not over-claim.

### Deliverable format
The thickened engine on `main` in `frank/` (commits) + the S2 exit-gate fixtures green (E2) + the OI-S1-F11-SWEEP full sweep discharged; your build relays in `frank/`; a SITREP back to master at the S2 exit gate.

### Operator-judgment items
- **residual risk (accepted):** D5 shell-routed confusion (restated).
- **escalation posture:** no auth/data/migration (greenfield store); the F2 delegated conditions gate any code; the S2-close sign-off is the operator's, exercised separately.
- **deferred (recorded):** the MCP live-adapter + fuller FieldSpec — returns with a testbed.

### Not authorized by this relay
No S2-close authority, no scope expansion beyond the IN list, no locked-design amendment. `DISPATCH IMPL` is delegated **only** under the conditions above; a failed condition escalates to master.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py <this file>` — run below.
- Executes `step1-plan` (VP-approved decomposition) + the S1 close (`s1-merge-gate/…-220652`, verified in `master/RECONCILE.md` § S1). Guide-refinement + MCP-deferral noted for VP visibility.
- Pointers: `master/ARCHITECTURE.md` §C4, `master/STEP-1-KICKOFF.md`, `ROADMAP.md`, `frank/` (baseline tag `s1-close`).

ACTIONS_GIT_REF: wrote this s2-dispatch relay + an `INDEX.md` row; cwd is not a git repo (docs workspace) so no sha — files on disk. No code, no dispatch token, no `frank/` write.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` is on branch `main`, clean tree, baseline tag `s1-close`.
Next requested action: operator relays this to the s2 slice-team (the s1 seats continued) + m-7.planner; s2 scaffolds via sprint-doc-setup in `frank/`, plans S2, and dispatches under the F2 delegated conditions (escalating to master only on a trigger); SITREP back to master at the S2 exit gate.
