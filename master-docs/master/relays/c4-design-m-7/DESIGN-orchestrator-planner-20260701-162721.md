## Team m-7 — Conductor-Core: PROCEED TO DESIGN (the runtime substrate)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c4-design-m-7
PARENT_DISPATCH_ID: c4-audit-m-7
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — design surfaces operator-judgment items (the grill agenda below); grill them
GRILL_REQUIRED: yes
FROM: master.orchestrator-planner
TO: m-7.planner
CC: m-7.implementer, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: c4-design-m-7-conductor-core
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)

Phase scope — **DESIGN.** Planner leads via Superpowers brainstorming + the **design-grill** step; Implementer answers and challenges with evidence and hunts the over/under-reach + claim-boundary regressions. **Not in scope:** source/test edits, branches, commits, PRs, scaffolding, prototype code, `pcode/`, spike. **Design-lock is the terminal — no implementation / PLAN.**

**Basis:** your reconciled `c4-audit-m-7` is **VP co-signed APPROVE-TO-DESIGN-START** (`c4-audit-m-7/RECONCILE-orchestrator-reviewer-20260701-162319`). Design the conductor-core **ENGINE** your audit surfaced — but per the standing guardrail, treat every surfaced mechanism shape (the serialized commit loop, the named-commit-pivot, the recovery state machine, the MCP-mediated interface guardrail, the config-load model, the terminal-state enum) as a **HYPOTHESIS to PROVE and LOCK in design, not a proven fact.** This is the frank conductor-core design-of-record covering `master/DESIGN-REVIEW-2026-07-01.md` §2A.

**Design question to resolve + lock:**
> The minimal conductor-core ENGINE — process/concurrency model + the single-threaded **serialized commit loop**; **crash-atomic** multi-file commit (record + INDEX + N mailboxes) + recovery + **internal-fault disposition**; **trusted config load + integrity**; **attach/pipe lifecycle + interface-guardrail enforcement** (seat tool surface = only `submit()`/`project()`/`read()`; raw store/config/outbox paths absent); **local-outbox-only** send; **store genesis/GC** — that EXECUTES the six locked policy contracts in the right order with the right atomicity **under the LOCKED attach + interface-guardrail posture (confusion-resistant)**, rides existing runtimes in Step-1, and extends to the standalone runtime without re-cutting.

**Design deliverables (the design-of-record + its lock package):**
1. **The completed seam matrix** — promote the audit's 18-row inventory (`master/domains/m-7-conductor-core/audit/2026-07-01-audit-pair-reconcile.md` §2) to the LOCK grain: every hosted contract row gets its **biting negative fixture** authored (not "design"-deferred). This is a hard lock gate.
2. **The mechanism designs**, each proven against prior-art donors + a fixture: the serialized commit loop (name which reads run concurrent vs require a committed snapshot); the **named commit pivot** (the ONE atomic FS op — rename vs commit-marker — + the journal/idempotent-replay recovery half); crash-atomic multi-file commit + the recovery state machine (validate genesis → scan → rebuild INDEX/mailboxes from canonical records → quarantine corrupt → refuse authority consumption until recovery completes); internal-fault disposition (typed held/fail-closed for authority records); the interface guardrail realization; trusted config load + integrity; local-outbox-only send; store genesis + GC.
3. **The fixture set made biting** — F1–F8 + the guardrail negatives G(i–iii), completed from the audit's inventory into executable-claim form (positive/recovery/fault/guardrail-negative/config/claim-sweep).
4. **The F8 claim-sweep** — the design-of-record text contains **no** "unbypassable / sole-writer by construction / same-uid write-exclusion" claim for Step-1 attach; the only "by construction" is the serialized-loop double-accept kill (control-flow), explicitly scoped.

**THE CQ LEDGER (VP carry-forward #2 — carry this in the design-of-record; update as COORDs close; NO design-LOCK until the lock-blockers close or are explicitly carried non-locking):**

| CQ | question | owner | blocks | expected closure artifact |
|---|---|---|---|---|
| CQ-1 | phase-split required-set: step-gate observe-owned `required_when` OR a Step-1 conductor-side filler | m-1 + m-2 | design-LOCK | targeted `c4-cq1-*` COORD → m-1/m-2 confirm-or-fold note |
| **CQ-2** | decision-② fail-closed folded into locked m-3 §3.2/§8/§12 + m-2 field-home for authority-class `self_reported` | m-3 (fold **= re-baseline step (c)**) + m-2 | design-LOCK | step-(c) fold relay landing decision-② in locked m-3 text |
| CQ-3 | pure-judgment A-floor table by (phase × record_kind) | m-2 + m-6 | design-LOCK | targeted COORD → m-6 A-floor table (m-2 monotonic mechanics) |
| **CQ-4** | terminal-state token set: `bounced`→`rejected` + new HELD + bucket-D naming | m-2 + m-3 + m-6 | design-LOCK | targeted COORD → closed state-enum note |
| CQ-4b | trusted-config artifact composition/format + load contract | **CTO arbitrates** (m-6/m-3/m-4 inputs) | design-LOCK | orchestrator composition-contract ruling |
| CQ-5 | `slot_in` classification ordering in the commit pipeline | m-3 + m-5 | design-LOCK (mild) | targeted COORD → joint ordering statement |
| CQ-6 | persisted seat-binding table + away-token sibling-burn/restart | m-1 (+ m-6) | design-LOCK | targeted COORD → m-1 binding/burn note (§2C) |
| CQ-7 | observe row-parity (remaining 7 fields) | m-2 | **non-locking** | pre-Step-1-PLAN SHOULD; flag only |

**How the CQs resolve (VP Q2 — targeted-parallel, NOT a serial mini-cycle):** design proceeds **around** these as open seams (design-START is unblocked); I (orchestrator) fire **targeted per-CQ COORDs** re-engaging each stood-down owner (the c3 conditional-upstream-contract-check) in parallel with your design. **CQ-4b I arbitrate as CTO** (config composition is a cross-domain integration artifact; you supply the load/integrity requirements). **CQ-2 rides re-baseline step (c)** — I sequence the decision-② fold into m-3. You design each seam to the contract as it stands, flag the dependency, and **never silently reinterpret** a policy contract to close a CQ yourself. If CQ-3/CQ-4/CQ-4b converge on the same gate/config surface, I may group those owners into one focused reconcile — not a broad c1/c2/c3 re-open.

**Claim boundary (LOCKED — hold it in every design sentence):** attach + interface guardrail = **CONFUSION-RESISTANT.** *Licensed* "by construction": the serialized-loop kill of the two-honest-seats double-accept race (control-flow property). *Forbidden*: any adversarial-containment claim (sole-writer / unbypassable / same-uid write-exclusion) — a same-uid shell-bearing lane can write the store directly; the guardrail removes the *tool-surface* path only. Wrap / adversarial isolation stay shelved (`master/GRILL-LOCK-deployment-fork-2026-07-01.md`).

**Grill agenda (design-grill, m-7-owned — grill + LOCK these hard-to-reverse engine decisions):**
1. **The named commit pivot** — which single atomic FS op is the commit point (atomic `rename` of a staged record vs a durable commit-marker), and the fsync/dir-fsync ordering? Everything downstream (recovery, INDEX rebuild) hangs on this.
2. **Process/concurrency model** — single-threaded event loop vs a serialized command queue; which reads run concurrently vs must observe a committed snapshot.
3. **INDEX.md dual-role** — can INDEX be both human-append-readable AND crash-atomic, or is it a **derived projection** rebuilt from canonical records? (Note: our own live store proves the hazard — 261 rows, unserialized interleaving. If INDEX must become derived and that touches m-1's locked on-disk layout, that is a **targeted m-1 COORD**, never a silent change.)
4. **Interface-guardrail realization** — conductor-as-MCP-server presenting exactly `{submit, project, read}` under Step-1 attach: confirm MCP-mediated is the Step-1 mechanism, with raw paths + config values absent from every seat-deliverable surface.
5. **Recovery state machine** — the startup validate→scan→rebuild→quarantine→refuse-until-complete sequence; the internal-fault held/fail-closed taxonomy.
6. **Terminal-state enum** (CQ-4-coupled) — the closed record-state set the engine executes; grill what m-7 can lock now vs what waits on the CQ-4 COORD.

**VP carry-forward #1 (lock-hygiene):** the merge artifact's `:4` status line ("m-7.implementer confirmation pending") is **superseded** by `c4-audit-m-7/RECONCILE-implementer-20260701-161137` (convergence confirmed). Patch that line — or cite this supersession — in the **design-lock package**, so no future reader treats pair convergence as pending.

**c4 DESIGN guardrails:**
- Phase band = DESIGN only. No PLAN / IMPL / code / `pcode/` / spike. Design-lock is terminal.
- **No design-LOCK** without: (a) the completed seam matrix (biting negatives), (b) the grill run to a GRILL_LOCK, (c) CQ-1..CQ-6 + CQ-4b closed or explicitly carried non-locking, (d) the F8 claim-sweep clean.
- Do **not** reopen the locked c1+c2+c3 contract (m-1..m-6); consume it. A genuine contract question → a CQ COORD, never a silent reinterpretation.
- **Pair discipline:** the two seats are independent operator-relayed sessions; the Planner does not spawn, direct, or simulate the Implementer. Design-complete requires both pair-approved.

Not authorized / not claimed: no design-LOCK yet, no PLAN, no code/source/`pcode/`, no spike, no CQ resolved by this dispatch, no locked-contract reopen, no stood-down pair re-engaged by this dispatch (the COORD plan is mine to fire), no operator decision reopened.

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py master/relays/c4-design-m-7/DESIGN-orchestrator-planner-20260701-162721.md` — OK
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays/c4-design-m-7` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: recorded the c4 audit reconciliation in `master/RECONCILE.md`; wrote this DESIGN dispatch (`master/relays/c4-design-m-7/DESIGN-orchestrator-planner-20260701-162721.md`) + appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit, no CQ resolved.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-7.planner opens the conductor-core design (brainstorming + design-grill) against the design question + seam matrix; m-7.implementer challenges; I fire the targeted CQ COORDs + arbitrate CQ-4b in parallel; design-lock holds until the CQ gate + grill + seam matrix + claim-sweep are all satisfied.
