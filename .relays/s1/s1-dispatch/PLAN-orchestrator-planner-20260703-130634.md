## PLAN — Slice-1 build dispatch (master → s1 slice-team; the design→build handoff)

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s1-dispatch
PARENT_DISPATCH_ID: step1-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no for this dispatch — operator ratified 2026-07-03; the S1 plan is guide(m-7)+VP-gated before any `DISPATCH IMPL`; merge remains a separate human gate at S1 close
DELEGATED_DISPATCH_AUTHORITY: no
GRILL_REQUIRED: no
IN_REPLY_TO: step1-plan/RECONCILE-orchestrator-reviewer-20260703-125826.md
FROM: master.orchestrator-planner
TO: s1.orchestrator-planner, m-7.planner
CC: master.orchestrator-reviewer, operator, m-1.implementer, m-2.implementer, m-7.implementer
SUBJECT: Stand up the s1 slice-team (via /orchestrator-planner, in frank/) + plan Slice-1 (thin end-to-end relay) — S1-scoped hardened gate; frozen m-1/m-2 contracts + fidelity edge; guide+VP gate before impl

**What this is.** The master CTO's dispatch executing the **VP-approved + operator-ratified** `step1-plan` decomposition (r2 `…-125826`; ratified 2026-07-03). It stands up the **s1 slice-team** to plan + implement **Slice-1**, the thin end-to-end relay. No `DISPATCH IMPL`, no code authorized here — the s1 plan is gated (m-7 guide + VP) first.

**Ratification + convention record (CC VP for visibility, not re-gate).** The operator ratified the build-execution model on 2026-07-03, with the correction that **slice-team relays live in `frank/`, governed by `sprint-doc-setup`** (not under `master-docs/master/relays/`). This dispatch records that; the VP is CC'd for visibility on the relay-location change (a Cardinal-#2 correction, operator-directed).

### To the s1 slice-team — your charter
- **Use `/orchestrator-planner`** — it carries the protocol + `sprint-doc-setup` (your relay/doc substrate). **Stand up your team in `frank/`** at whatever structure + granularity you judge right — that's yours to decide. Suggested `RUN_ID: s1` to keep the trail legible; your relays live in `frank/`, not the master governance trail.
- **Spec = read-only reference in cwd:** `master-docs/master/ARCHITECTURE.md` §C4 (the engine) + the m-1 / m-2 / m-7 domain docs. **Do not edit governance docs**; escalate spec problems to master — do not self-amend a locked design.
- **Code lands in `frank/`** on `main` (empty repo — your first commits).
- **Your m-x guide = m-7** (conductor-core): feeds the locked design, answers domain questions, co-gates your S1 plan. Reach the guide via `m-7.planner`.

### Slice-1 scope (IN)
The brutally-small end-to-end path: `mint → connect → render-MVP-form → submit → stamp → validate → lineage → append(crash-atomic) → project → deliver → gate-outbox`. Concretely:
- engine skeleton — the single-threaded **serialized commit loop** + **crash-atomic rename commit**;
- channel/identity — `connect` binds the channel; **`FROM` is system-stamped**, not lane-supplied;
- the **interface guardrail** — seats reach only `submit` / `project` / `read`; raw store/config paths absent from the seat tool surface — **+ I-PH path hygiene**;
- a **tiny MVP FieldSpec** — only the fields this slice needs (the full registry is S3);
- minimal **validate** (required-set / enum / seat-scope) + minimal **lineage** (parent-edge presence + validity);
- terminal **append** `{accepted | rejected | held}` (byte-exact enum) + **projection rebuild** + `read`;
- **deliver** via `project`; one **gate** produces a local outbox / ODB item;
- a **minimal dissolved-linter replay** over the historical upstream lint failures the MVP FieldSpec covers.

### Slice-1 scope (OUT — deferred; do not build)
Full FieldSpec registry + 62-check linter refactor + the **full** "~33 checks dissolve" replay (S3) · full recovery phases 0–4 + FIFO durability + GC/genesis + the owed-item projection (S2) · all consumer schema fields — observe/routing/archetype/ODB (S4) · observe-as-send (Step-2) · routing execution (Step-3). If a task seems to need an OUT item, **escalate to master** — do not expand scope.

### Frozen contracts + the fidelity edge (F3)
- **Consumed, locked:** the m-1 store API (`submit`/`project`/`read`, append-only, sole-writer) + the m-2 FieldSpec envelope. Build **against** them; **no re-design**.
- **Fidelity gate:** `m-1.implementer` (store-API usage fidelity) + `m-2.implementer` (FieldSpec-envelope usage fidelity) are **required reviewers of your consuming surface**; each returns a fidelity-review approve **before your `DISPATCH IMPL` is live**. This reviews your *usage*, never their locked design. A fidelity finding blocks your dispatch until your usage is corrected — the contract itself is not changed. If you believe the contract is wrong, escalate to master.

### Boundary contract
```
Writes: append-only relay-store records (terminal {accepted,rejected,held}) + a local outbox/ODB item
Reads: the locked m-1 store API + m-2 FieldSpec envelope
Target entity: the conductor's committed relay + its rebuilt projection
Downstream consumer: seat inboxes (via project) + the local outbox
Contract: submit/project/read; channel-stamped FROM; byte-exact {accepted,rejected,held}; MVP FieldSpec; crash-atomic commit
Proof: the S1-scoped hardened exit gate (E2 fixtures)
No-consumer action: n/a — S1's consumer is the relay flow itself
```

### Slice-1 exit gate — the S1-SCOPED hardened gate (HARD acceptance)
Baseline: accepted only through the conductor · `FROM` system-stamped · form/lint before delivery · gate produces a local outbox item. Promoted:
- **Adversarial:** forged `FROM` rejected · forbidden enum absent-then-rejected · invalid parent rejected · **duplicate-sibling double-accept killed** (the serialized-loop guarantee, actually exercised).
- **Crash-atomicity (crash it for real):** `kill -9` mid-commit and mid-delivery → exactly-once outcome + re-issued wake · crash after intake-fsync (no lost intake) · crash before/after the atomic rename (presence = committed) · corrupt-projection rebuild from canonical · replayed intake-id (no double-emission).
- **Dissolved-linter replay (S1-minimal):** the MVP-covered historical upstream lint failures run through the MVP validator are caught-or-genuinely-obsolete. (The full ~33-check replay is an S3 gate.)
- **Path-hygiene (I-PH):** no seat-facing output — every bounce/error included — contains a canonical store path.
- **Liveness:** inbox = durable truth, pipe `write()` = nudge; a busy/dead seat still receives via `project()` on reconnect; a lost wake never sleeps a parked lane forever.
- **park/wake:** a gated lane parks (consumes nothing) + wakes on the operator's verdict.

### Plan-gate (F2 — S1/bootstrap)
Produce your **S1 PLAN** (in `frank/`); it is gated by **m-7 (guide) + the VP** before your `DISPATCH IMPL`. **No code before that gate.** Your internal plan-review + delegated dispatch are ordinary protocol *below* that gate.

### Owed §C4 carries landing in S1 (materialize-first)
The **code-layer interface-guardrail enforcement** (the one genuinely new build item) + the **I-PH path-hygiene fixture**; plus the ③ known-A / RAISE-ONLY NF fixture (guardrail-adjacent portion). Write each as a **typed owed-item record** `{owner, source, target surface, disposition path}` before treating it as covered.

### Framing (honesty — must hold in code AND any doc)
Step-1 = **provenance + transport, not verified work** (observe is Step-2; "done" here is `self_reported`). Only the **serialized-loop double-accept kill** (and, with a constrained grammar, **R2**) are *operationally live* in S1; the others are *recorded*, not enforced, until their later layers. Confusion-resistance is **tool-mediated** — removes affordance, not access; **D5** (shell-routed confusion) is an accepted residual. Do not let S1 code or docs over-claim.

### Deliverable format
The first working conductor relay in `frank/` (branch + first commits) + the S1-scoped exit-gate fixtures green (E2); your build relays in `frank/`; a SITREP back to master at the S1 exit gate.

### Operator-judgment items
- **residual risk (accepted):** D5 shell-routed confusion — restated so S1 code does not over-claim.
- **escalation posture:** no auth/data/migration in S1 (greenfield store, empty repo); the S1 plan is gated before any code; merge is a separate human gate at S1 close.

### Not authorized by this relay
No `DISPATCH IMPL`, no merge, no code, no scope expansion, no `frank/` write beyond the sprint scaffold + the planned S1 build after the guide+VP gate. The s1 plan must pass the guide+VP gate first.

## Verification
- `python3 ~/.claude/skills/tools/relay-lint.py <this file>` — run below.
- Executes `step1-plan` r2 (VP `approve` `…-125826`); ratification 2026-07-03; convention per `STEP-1-KICKOFF.md` (build-execution model).
- Pointers: `master-docs/master/ARCHITECTURE.md` §C4 + §C4.3/I-PH, `master-docs/master/STEP-1-KICKOFF.md`, `ROADMAP.md` Step-1, `frank/` (empty repo, branch `main`, 0 commits).

ACTIONS_GIT_REF: wrote this s1-dispatch relay + an `INDEX.md` row; cwd is not a git repo (docs workspace) so no sha — files on disk. No code, no dispatch token, no `frank/` write yet.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` is the code repo, empty (branch `main`, 0 commits, `git status --short` empty).
Next requested action: operator relays this to a fresh `s1.orchestrator-planner` session (CC m-7.planner) to stand up the s1 team via `/orchestrator-planner` in `frank/`; s1 decides its own team structure, plans S1, and returns the plan for the m-7+VP gate before any `DISPATCH IMPL`.
