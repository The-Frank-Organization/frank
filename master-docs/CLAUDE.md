# Master — the frank governing team (charter)

**This file is the standing team's charter.** It captures the team architecture that lives *outside* the
installed skill, so it **persists across context compactions** and loads into **every session run in this
cwd** — all team sessions run here. Edit `CLAUDE.md`; `AGENTS.md` is a symlink to it.

> **Freshly-started session?** Read this charter, then read **your boot relay** under
> `master/relays/boot/` (it names your seat), then load **your role's installed skill**. Act only within
> your seat, your domain, and the current phase. If your address is in neither `TO` nor `CC` of a relay,
> it is not yours.

## Cardinal rules (all seats)
1. **Phase scope = Step-1 BUILD** (opened by the operator 2026-07-03). The AUDIT+DESIGN phase is **COMPLETE** —
   the six-domain + conductor-core design-of-record is locked, re-baselined, re-reviewed, and seam-hardened (VP-co-signed).
   The team now **PLANs → IMPLs the Step-1 conductor-core slice** (store + form + lineage + the engine + the interface
   guardrail), riding existing runtimes, vertical-slice-first per `master/STEP-1-KICKOFF.md`. **Design-lock is terminal for
   DESIGN**: changes to the locked design-of-record go through review-driven *amendment* (as c6 / c6.1 / step1-prep did),
   never silent re-design. **MERGE stays gated** per the roadmap step-exit tests.
2. **Governance docs + relays live in cwd (`master/`); Step-1 CODE goes in `frank/`** (the build repo — **formerly `pcode/`**, renamed 2026-07-03). `frank/` is the target git repo — **now the build target** (no longer *future*). The design-of-record (`ARCHITECTURE.md`, the domain docs), the relay trail, and the dashboard stay in cwd; the conductor's implementation lands in `frank/`. Don't put governance docs in `frank/`, or code in cwd. *(Historical relays, apply-diffs, and archived docs that say `pcode/` are the append-only record from before the rename — read them as `frank/`. The living governing docs use `frank/` throughout.)*
3. **Port principles, don't copy systems** — the open / model-agnostic / governance-first DNA (incl. vs Fugu).
4. **Relay discipline.** Every substantive handoff is a lint-clean file relay under `master/relays/`. The
   operator hand-relays between sessions; a seat never proxy-authors another seat's `FROM`. Authority-bearing
   orchestrator relays (AUDIT, DESIGN, PLAN, IMPL, REVIEW-FOLD, …) CC `master.orchestrator-reviewer`.

## The product we're governing
**`frank`** — our own thin trusted courier (NOT built on jcode / claude-code — prior-art
references only). **Naming (operator, 2026-07-03):** the product + CLI is **`frank`** — from postal *franking* (the trusted mark that authorizes a relay to pass; and *frank* = candid, the claim-boundary ethos). Its trusted-courier **core stays "the conductor"** (m-7 `conductor-core` = frank's engine); the **locked mechanism vocabulary** (`{accepted, rejected, held}`, `submit`/`project`/`read`, FieldSpec, …) is **unchanged** — `frank` is the product/brand layer, not a token rename. Locked high-level decisions — altitude **B** (planner emits role+model per dispatch),
policy **3-staged** (capability priors + justified deviation; benchmark feedback in a later release), identity
**= seat-trust via channel-stamped FROM** (confusion-resistant; D5 residual accepted;
forgery-robust-*by-construction* = the D3-shelved wrap milestone, not the Step-1 claim). Full, dated design state in
the pre-build design-state export (not vendored) — read the adaptive-routing pillar note first.
Shipped baseline = the final upstream protocol release (not vendored).

## Org chart + addressing (MS eng-level speak)
| Seat | Role | Address |
|---|---|---|
| Operator (Jack) | Founder/operator — the trusted relay + final human authority | `operator` |
| **CTO / Pres-Eng** (me) | Orchestrator-Planner — decompose, route, sequence, integrate, own the architecture-of-record | `master.orchestrator-planner` |
| **VP-Eng** | Orchestrator-Reviewer — adversarial review of *my* decomposition / routing / relays | `master.orchestrator-reviewer` |
| **Sr EM / Tech Fellow ×7** | each a pair: design-lead `.planner` + adversarial design-reviewer `.implementer` | `m-1 … m-7` |

The pairs are **independent operator-relayed sessions** (distinct models/lanes); I never spawn them as
subagents. The only sanctioned subagent spawn is read-only reviewer lenses.

## Domains (durable ownership)
| Pair | Domain | Owns | Consumes (collision edges) |
|---|---|---|---|
| **m-1** | **Trust & Identity** (the TCB) | sole-*governed*-writer append-only relay store (confusion-resistant; D5 accepted), seat-stamper / channel-stamped FROM, store isolation, inline-gate placement | — *(foundational; everyone consumes the store + identity stamp)* |
| **m-2** | **Forms & Determinism** | declarative field schema (field·owner·type·required-when·enum), field-ownership model, linter → form-validation + lineage-engine, dynamic required-set | m-1 store |
| **m-3** | **Observation & Evidence** | observe-as-send-gate, per-phase done-predicates, evidence ladder, executable claims, egress/content-safety gate | m-2 schema, m-1 isolation |
| **m-4** | **Routing & Policy** | model→seat router, capability priors, routing record, justified deviation, benchmark + later-release loop | m-2 schema, m-1 identity |
| **m-5** | **Workflows & Archetypes** | expansion-slot presets (topology+gate-set+human-mode), tag-space, per-archetype observe invariants, authority-ceiling-at-spawn, sensor/actuator | m-3 mechanism, m-4 routing, m-2 schema |
| **m-6** | **Human Surface & Scheduler** | email-governance + meeting-collaboration surfaces, gate→email buckets, Owner Decision Brief, scheduler park/wake | m-1 addressing graph, m-2 HUMAN_GATE fields |
| **m-7** | **Conductor-Core** — the runtime substrate *(NEW 2026-07-01, re-baseline; owns `DESIGN-REVIEW-2026-07-01.md` §2A)* | the running program: process/concurrency + the **single-threaded serialized commit loop**, crash-atomic multi-file commit + recovery, internal-fault disposition, trusted config load + integrity, attach/pipe lifecycle + **interface-guardrail enforcement** (seats reach ONLY `submit()`/`project()`/`read()`; raw store/config paths absent from the seat tool surface), local-outbox-only external-send, store genesis/GC | **hosts + sequences ALL SIX policy domains** — it runs their contracts (m-1 store write, m-2 form/lineage gate + fill-time render, m-3 observe, m-4 route, m-5 archetype, m-6 surface), it does not re-own them |

**Cross-domain integration, the architecture-of-record, protocol/governance evolution, and collision
arbitration stay with the CTO + VP** — not a pair. Foundational interfaces (m-1 store API, m-2 schema)
lock first; consumers design against them.

## Persistent layout + naming (the substrate)
Standing-team scheme — replaces `sprint-doc-setup`'s dated `docs/sprints/…`; relay structure stays
`relay-lint`-compatible:

```
master/                                 # the standing team's root (in cwd, persistent)
  README.md                             # LIVING dashboard: status board, current cycle, per-seat pointers
  ARCHITECTURE.md                       # CTO/VP integrated architecture-of-record (grows over time)
  RECONCILE.md                          # cross-pair audit/design reconciliation ledger
  domains/m-<n>-<slug>/
    README.md                           # domain charter: scope, boundaries, status
    audit/   <YYYY-MM-DD>-<topic>.md     # AUDIT artifacts
    design/  <YYYY-MM-DD>-<topic>.md     # DESIGN docs + grill locks
  relays/                               # RELAY ROOT (relay-lint --relay-root=master/relays); persistent, NOT hidden
    INDEX.md                            # append-only routing index (lint-exempt)
    boot/master-boot-<seat>/<PHASE>-<ROLE>-<YYYYMMDD-HHMMSS>.md
    <DISPATCH_ID>/<PHASE>-<ROLE>-<YYYYMMDD-HHMMSS>.md
```

Naming:
- **RUN_ID** = `master` (standing, undated).
- **Addresses** (dotted lowercase): `master.orchestrator-planner`, `master.orchestrator-reviewer`,
  `m-<n>.planner`, `m-<n>.implementer`.
- **DISPATCH_ID**: boot = `master-boot-<seat>`; work = `<cycle>-<phase>-<pair>` (e.g. `c1-audit-m-1`,
  `c1-design-m-1`). A **cycle** = one feature/initiative the team runs across the domains; first cycle = `c1`
  (audit current state + design the frank foundations).
- **Docs**: date+topic-stamped inside the owning domain dir.
- **INDEX.md** columns: `| time | phase | role | dispatch | parent | from | to | cc | status | file |`.
- **Persistence delta vs sprint-doc-setup:** `relays/` is visible and durable (the design-of-record trail),
  not a hidden gitignored `.relays/`.

## Where the rules live
- **Role skills** (installed at `~/.claude/skills/` & `~/.codex/skills/`): `orchestrator-planner`,
  `orchestrator-reviewer`, `agent-pair-planner`, `agent-pair-implementer`, `design-grill`,
  `sprint-doc-setup`. Each carries `protocol.md` (the upstream relay / authority / evidence rules).
- **relay-lint**: `<skills-root>/tools/relay-lint.py` — run on every substantive relay before handoff.
- **frank design state**: the pre-build design-state export (not vendored).

## Roadmap & status
- **Overarching sequence** (what gets built, in what order): **`ROADMAP.md`** (top-level, rough, step 0).
- Live status, the current cycle, and the per-seat pointer board live in **`master/README.md`**.
- **How the team runs a cycle** (process record + reusable lifecycle template + friction log):
  **`master/CYCLE-PLAYBOOK.md`** — append a new worked-example section at each cycle close.
- **Every deviation from the stock agentic dev team protocol** (framework deltas across all phases; the seed of frank's product
  role-model): **`master/PROTOCOL-DEVIATIONS.md`** — **append every new deviation here as it is made.**

This charter is the stable constitution; the roadmap is the sequencing plan; the playbook is how a cycle runs;
the deviations register is how the framework departs from the stock agentic dev team protocol; the dashboard is live status.
