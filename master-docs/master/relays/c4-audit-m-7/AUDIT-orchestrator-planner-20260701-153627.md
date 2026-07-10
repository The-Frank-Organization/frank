## Team m-7 — Conductor-Core (AUDIT)

ROLE: Orchestrator Planner
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c4-audit-m-7
PARENT_DISPATCH_ID: conductor-core-standup
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — read-only audit; surface operator-judgment items in findings
FROM: master.orchestrator-planner
TO: m-7.planner, m-7.implementer
CC: master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)

Phase scope — **AUDIT (read-only)**, opening **Cycle c4 — the conductor-core re-baseline cycle** (the substrate the design review found was nobody's domain). Inspect source and docs, run safe read-only commands, produce an independent paired audit with findings. **Not in scope:** any edits, branches, commits, PRs, scaffolding, or prototype code — and specifically **no `pcode/` writes, no spike.** Still the design-of-record research + design phase; no implementation exists or is authorized. This audit does not open Step-1 PLAN.

**Fresh-domain, full audit.** Unlike m-5's c3 (which had a prior narrow engagement to consolidate), m-7 is a **brand-new domain with no design-of-record**. This audit is two things: **(a)** the **runtime-substrate prior-art sweep** m-7 has never run, and **(b)** mapping the `DESIGN-REVIEW-2026-07-01.md` **§2A requirement set** to auditable form — for each substrate mechanism, what **exists to promote** vs what is **net-new**.

Pair roles & research method:
- **m-7.planner** (design-lead): lead the audit; surface the substrate design questions. Use parallel agents + websearch + a deep-research workflow for the prior-art sweep (process/concurrency models, serialized-commit / single-writer loops, crash-atomic multi-file commit + recovery, tool-surface restriction under attach, config-integrity-at-startup).
- **m-7.implementer** (adversarial design-reviewer): run an **INDEPENDENT** audit — do not mirror the planner — and pressure-test every substrate claim against the **LOCKED attach + interface-guardrail posture** (confusion-resistant, **NOT** wrap, **NOT** adversarial isolation). Hunt the **over-reach** (m-7 silently re-owning a policy decision) and **under-reach** (substrate left homeless) failure modes from your boot.
- **Independent paired audit:** each seat audits separately, then reconciles. The Planner does not spawn, direct, or simulate the Implementer; the two seats are independent operator-relayed sessions.

**Domain context.** m-7 owns the **ENGINE** — the running program the six policy domains ride on. The one-line boundary: *conductor-core owns the ENGINE (how things run); the six domains own the CONTRACTS (what is valid/required/gated); conductor-core EXECUTES their contracts — right order, right atomicity, behind the right interface — and does NOT re-own policy.* Scope = `DESIGN-REVIEW` §2A (`master/domains/m-7-conductor-core/README.md`).

**You build on the LOCKED c1+c2+c3 contract — do NOT reopen m-1..m-6** (`master/ARCHITECTURE.md` §1–§C3). m-7 **hosts + executes** those contracts; it consumes, it does not renegotiate. The hosted seams you audit against (contract owners in parens):
- **m-1** store append + channel-stamped FROM stamp (schema + stamp contract = m-1's; the append/stamp *execution inside the serialized loop* = m-7's);
- **m-2** fill-time-authority form render + **phase-split required-set** (FieldSpec + predicate = m-2's; render forbidden-options-absent + validate = m-7's; a Step-1 with no observe writer must not demand observe-owned fields);
- **m-3** observe hook + **decision-② fail-closed** on authority-class `self_reported` (observe/evidence contract = m-3's; hosting the hook + class-conditional fail-closed = m-7's);
- **pure-judgment A-floor** (floor contract = m-2/m-6's; enforce at fill/submit = m-7's);
- **m-4/m-5/m-6** routing-record / archetype-spawn / human-surface (contracts theirs; sequence/host = m-7's).

**Sources to audit** (cross-check the export's distillation against real source):
- **The upstream protocol's current state:** the upstream protocol release corpus (not vendored) — the running orchestrator today: how it commits relays, orders writes, loads config, restricts (or fails to restrict) the agent's file surface. Our LIVE team (`master/`) is itself a running instance — audit how *this* relay store actually gets written.
- **jcode:** `references/jcode/` — the swarm/channels bus, per-seat connection binding (`ClientConnectionInfo`), the write path + any serialized/queued commit; cross-ref `references/README.md` (m-1/m-7 runtime rows).
- **claude-code:** `references/claude-code/` — the per-agent JSON inbox + **lockfile** + poll-and-inject delivery (the crash-atomicity + concurrent-write prior-art most directly analogous to our multi-mailbox commit).
- **Runtime research (primary-source-verified):** `master/RUNTIME-RESEARCH.md` — esp. **§8** (the `srt`/attach facts) and **:772-780** (the confirmed `srt` + Codex app-server facts). This is the load-bearing evidence for whether attach can deliver the interface guardrail.
- **Crash-atomicity / serialized-commit prior-art** (websearch): single-writer append logs, atomic multi-file commit (write-tmp-then-rename, journaling), recovery/reconciliation, corrupt-record quarantine.
- **The finding set:** `master/DESIGN-REVIEW-2026-07-01.md` §2A (your requirement list) + `GRILL-LOCK-deployment-fork-2026-07-01.md` (D2/D2b — the interface guardrail that makes the confused-agent model hold).
- **The locked m-1..m-6 domain design docs** (`master/domains/*/design/`) as **CONTRACT INPUTS** — m-7 hosts these contracts; it does not invent substitutes.

**Design question to resolve (surfaced for c4 DESIGN, not locked in audit):**
> What is the minimal conductor-core **ENGINE** — the process/concurrency model + the single-threaded **serialized commit loop**, **crash-atomic** multi-file commit + recovery + **internal-fault disposition**, **trusted config load + integrity**, **attach/pipe lifecycle + interface-guardrail enforcement** (seat tool surface = only `submit()`/`project()`/`read()`; raw store/config paths absent), **local-outbox-only** send, and **store genesis/GC** — that EXECUTES the six locked policy contracts in the right order with the right atomicity **under the LOCKED attach + interface-guardrail posture (confusion-resistant)**, rides existing runtimes in Step-1, and extends to the standalone runtime (later steps) without re-cutting?

**Hard acceptance criteria:**
1. A **4-bucket verdict** (still-open / already-closed / product-overlapped / recommended-next) on the frank conductor-core substrate vs what **the upstream protocol** (relay commit/config today), **jcode** (bus + connection-binding + write path), **claude-code** (inbox + lockfile + poll), and the **srt/attach** facts already provide.
2. The **process/concurrency model + serialized commit loop** — how `submit()` read-validate-append and `verify()` check-and-burn serialize into one critical section, and how that **kills the two-honest-seats double-accept race** (see the claim-boundary note below — this concurrency invariant is the one *licensed* "by construction").
3. **Crash-atomicity + recovery** — atomic multi-file commit (record + INDEX + N mailboxes) + recovery/reconciliation + corrupt-record quarantine; the **internal-fault disposition** (a trusted check that throws/times-out → held/fail-closed for authority records; never silent-accept, never brick).
4. The **interface-guardrail realization under attach** — HOW the seat tool surface is held to only `submit()`/`project()`/`read()` with raw store/config paths **absent**, given same-uid attach (the fork grill-lock D2/D2b). Name the **negative fixture** shape (a confused seat *reaching for* a raw store/config path finds no such tool).
5. **Trusted config load + integrity** — config conductor-owned, loaded once at trusted startup, integrity-checked, absent from every seat tool surface.
6. The **SEAM INVENTORY** (precursor to the DESIGN-phase seam matrix) — for each hosted contract: `{contract owner · contract doc/section · m-7 execution obligation}`. The full matrix (adding **biting negative fixtures** + a **contract-question-raised?** flag) is a **DESIGN-lock** deliverable; audit surfaces the inventory and flags any contract that is **under-specified** for m-7 to execute against (→ a contract question for a *targeted* COORD, never a silent reinterpretation).
7. **`away`/parked-lane + restart** touch-points m-7 must recover (conductor-restart seat-binding recovery; store genesis) — surfaced, not designed.

**c4 GUARDRAILS (VP boot constraints, `conductor-core-standup/RECONCILE-orchestrator-reviewer-20260701-152055`):**
- **Phase band = AUDIT + DESIGN only.** No build / PLAN / IMPL / code / `pcode/` / spike.
- **Do NOT reopen the locked c1+c2+c3 contract** (m-1..m-6). Host/consume; do not renegotiate. A genuine contract question → **surface it** (targeted COORD at DESIGN); never silently reinterpret a policy contract.
- **Pair-artifact requirement (F4):** BOTH m-7.planner and m-7.implementer return an audit artifact, OR one explicitly reconciled pair artifact. No single-seat audit.
- **CLAIM BOUNDARY (the crux of this whole re-baseline).** The posture is **attach + interface guardrail = CONFUSION-RESISTANT**. Draw the line precisely:
  - **LICENSED "by construction":** the **concurrency invariant** — a single-threaded serialized commit loop structurally prevents two *honest, racing* seats from both passing a check-and-burn (double-accept / replay). That is a property of the trusted engine's own control flow, and it holds regardless of seat behavior. Claim it.
  - **FORBIDDEN "by construction":** any **adversarial-containment** claim — "sole-writer / unbypassable / a malicious seat cannot reach the store." Under same-uid attach a confused-*with-a-shell* or malicious seat can write the store directly; the guardrail only removes the *tool-surface* path (confused-agent-resistant), it does not *contain* a code-executing adversary. Adversarial isolation / wrap / "by-construction sole-writer" remain **SHELVED** unless a later operator-gated spike changes that.
- **SEAM MATRIX** (with biting negative fixtures) is required **before design-lock** — this audit produces the inventory precursor.
- **GRILL_REQUIRED: yes** will be set at DESIGN (substrate semantics are cross-domain + hard-to-reverse).

**Not authorized / not claimed:** no Step-1 PLAN, no code/source/`pcode/`, no spike, no branch/commit/PR, no locked-contract reopen, no concrete substrate **lock** in audit (surface for the c4 DESIGN-lock, m-7-owned), no operator decision reopened.

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py master/relays/c4-audit-m-7/AUDIT-orchestrator-planner-20260701-153627.md` — OK
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays/c4-audit-m-7` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this audit dispatch (`master/relays/c4-audit-m-7/AUDIT-orchestrator-planner-20260701-153627.md`) + appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-7.planner + m-7.implementer each run the independent substrate audit (4-bucket verdict + §2A mechanism sweep + seam inventory), then pair-reconcile; audit-reconcile re-engages the VP.
