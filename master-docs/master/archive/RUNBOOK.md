# RUNBOOK — standing up a governed standing team on a new app

**What this is.** A self-contained init procedure for replicating the `master`-team structure (the
standing-team extension of the **agentic dev team protocol**) on a different app/project.
It was distilled from the worked setup in this repo (charter `CLAUDE.md`,
`CYCLE-PLAYBOOK.md` Parts A–D, relay trail `master/relays/`). Copy this one file into the new project's
cwd and hand it to a fresh **orchestrator-planner** session with the instruction: *"init the standing
team from RUNBOOK.md."*

**Who executes it.** The orchestrator-planner seat (CTO). The operator is the trusted transport between
sessions; every other seat comes online via a boot relay the operator hand-relays.

---

## 0. The model in one paragraph

The stock protocol supplies the protocol physics — role skills, lint-clean file relays, authority/evidence/ceremony
rules, pair adversarialism, VP visibility. This runbook adds the four layers that made the standing team
work: a **persistence layer** (a charter in `CLAUDE.md` that survives compaction and auto-loads every
session, plus a visible durable relay root), an **org layer** (durable domain ownership across N pairs
instead of per-sprint bundles), a **process-memory layer** (a self-amending cycle playbook), and a
**scope governor** (phase-banded cycles — e.g. AUDIT+DESIGN only, terminating at design-lock). No
installed skill is modified; everything is layered via the charter and boot relays.

---

## 1. Preconditions (verify before anything else)

1. **Skillset installed** at `~/.claude/skills/` (and `~/.codex/skills/` if any lane runs Codex):
   `orchestrator-planner`, `orchestrator-reviewer`, `agent-pair-planner`, `agent-pair-implementer`,
   `design-grill`, `sprint-doc-setup`, each with its `protocol.md`. Superpowers installed (the role
   skills hard-require it).
2. **relay-lint reachable:** `python3 ~/.claude/skills/tools/relay-lint.py --help` exits 0. Never
   hardcode an absolute home path inside project docs — the toolchain has moved hosts before.
3. **The cwd decision:** pick the project root the whole team will run in. ALL team sessions run in this
   cwd — that is what makes the charter auto-load. If the app has a code repo, decide now whether docs
   live beside it or in a sibling docs workspace (in the reference setup, docs live in cwd and the code
   repo `pcode/` stays untouched during design phases).

---

## 2. Parameter block (fill in before executing; everything downstream derives from it)

| Parameter | Reference value | Yours |
|---|---|---|
| `RUN_ID` | `master` (standing, undated — NOT a dated sprint id) | |
| Mission / product | "design frank, the conductor" | |
| **Phase-band** | AUDIT + DESIGN only; terminate at design-lock | |
| Forbidden zone | no code, no writes to the code repo (`pcode/`) | |
| Team root dir | `master/` in cwd | |
| Domain pairs `m-1..m-N` | 6 at standup (a 7th added later by evidence) | |
| Foundational domains (lock first) | m-1 trust/identity, m-2 forms/schema | |
| Orchestrator-planner model/lane | Claude Opus-class | |
| Orchestrator-reviewer model/lane | a *different* vendor/lane (GPT-class) | |
| Pair lanes | distinct models/lanes per pair where possible | |
| First cycle `c1` scope | audit current state + design the foundations | |

Rules of thumb baked into the reference values:

- **Decompose by durable ownership, not by task.** Each domain owns interfaces others consume; write the
  consumes/collision edges into the org table from day one.
- **Foundational domains are the ones everyone consumes** (identity/store, schema). They lock first;
  consumers design against them. Co-foundational domains with a load-bearing seam lock **together**.
- **Cross-domain integration, the architecture-of-record, and collision arbitration stay with the
  CTO + VP** — never delegated to a pair.
- **The VP is a different model/vendor lane than the CTO.** The adversarial review only bites if the
  reviewer doesn't share the planner's blind spots.
- Start with the smallest pair set that covers the foundations; **domains can be added mid-flight** when
  evidence shows a gap (the reference team stood up m-7 Conductor-Core only after an adversarial review
  found the runtime substrate was nobody's domain). Adding a domain = VP-approved decomposition relay +
  domain charter + boot relays; it is a normal, sanctioned move.

---

## 3. Setup sequence (once per team)

Run in this order. Every authority-bearing step is VP-reviewed **before** execution — including the
setup steps themselves; the setup runs through the same relay protocol it later governs.

```
SETUP
  1. boot the VP            (boot relay → operator hand-relays to a fresh reviewer session)
  2. ORG-DECOMP             propose the org/domain decomposition → VP review → fold
  3. ROADMAP                draft the build sequence → VP review (multi-round OK) → persist as ROADMAP.md
  4. CHARTER                write CLAUDE.md (+ AGENTS.md symlink)  ← the persistence step; do not skip
  5. LAYOUT                 create the team root + domain dirs + relay root + INDEX.md
  6. boot the pairs         one boot relay per seat, foundational pairs first
```

Notes on each:

- **(1, 6) Boot relays** grant no work authority (`AUTHORITY: report-only`, `PHASE: SITREP`). Work
  authority arrives only with the phase dispatches. See §6 for the template.
- **(2) ORG-DECOMP** goes in its own dispatch thread (`<RUN_ID>-org-decomp`). Expect operator pushback
  to simplify — the reference team dropped a 4-tier recursive org for flat + lean because *the operator
  is the transport layer*; every extra tier is a hand-relay tax.
- **(3) ROADMAP** lives at the top level, marked rough/step-0. Sequencing principles that held up:
  own the differentiator first; foundations before consumers — but design foundations *with their
  consumer contracts sketched and reviewed* (the "designed-early, executed-later" rule); ride existing
  infrastructure before replacing it; every step gets an **observable exit test**, not a vibe.
- **(4) CHARTER is the load-bearing step.** The role skills carry only the generic protocol; the team's
  org, addressing, domain map, layout, and cardinal rules live in `CLAUDE.md` so they survive context
  compaction and load into every session in the cwd. Symlink `AGENTS.md → CLAUDE.md` for non-Claude
  lanes. §4 has the skeleton.
- **(5) LAYOUT — deliberate override of `sprint-doc-setup`.** The orchestrator-planner skill's init
  directive says to run `sprint-doc-setup`; for a *standing* team you override its two transient
  defaults: no dated `docs/sprints/…` tree (use durable domain dirs) and no hidden gitignored
  `.relays/` (use a visible, durable `<team-root>/relays/` — the relay trail IS the design-of-record).
  Everything else (relay schema, lint compatibility, INDEX) is kept `relay-lint`-compatible.

### The layout to create

```
<team-root>/                          # e.g. master/ — persistent, in cwd
  README.md                           # LIVING dashboard: current phase, org status, decisions, pointers
  ARCHITECTURE.md                     # CTO/VP integrated architecture-of-record (grows per cycle)
  RECONCILE.md                        # cross-pair reconciliation ledger
  CYCLE-PLAYBOOK.md                   # process memory: worked example per cycle + distilled template
  domains/m-<n>-<slug>/
    README.md                         # domain charter: scope, boundaries, consumes-edges, status
    audit/   <YYYY-MM-DD>-<topic>.md
    design/  <YYYY-MM-DD>-<topic>.md
  relays/                             # relay root — visible, durable, never hidden
    INDEX.md                          # append-only routing log (lint-exempt)
    boot/<RUN_ID>-boot-<seat>/<PHASE>-<ROLE>-<YYYYMMDD-HHMMSS>.md
    <DISPATCH_ID>/<PHASE>-<ROLE>-<YYYYMMDD-HHMMSS>.md
```

---

## 4. Charter skeleton (`CLAUDE.md`, symlinked to `AGENTS.md`)

Adapt; keep every section — each earned its place.

```markdown
# <RUN_ID> — <mission> governing team (charter)

**This file is the standing team's charter** — the team architecture that lives *outside* the installed
skill, persisting across context compactions and loading into every session in this cwd.

> **Freshly-started session?** Read this charter, then your boot relay under
> `<team-root>/relays/boot/` (it names your seat), then load your role's installed skill. Act only
> within your seat, your domain, and the current phase. If your address is in neither TO nor CC of a
> relay, it is not yours.

## Cardinal rules (all seats)
1. **Phase scope = <PHASE-BAND> only.** <what terminates the band, e.g. per-domain design-lock>.
2. **Docs live in <docs location>, never in <code repo>.**
3. <the product-DNA rule — the principles to port, the systems not to copy>
4. **Relay discipline.** Every substantive handoff is a lint-clean file relay under
   `<team-root>/relays/`. The operator hand-relays between sessions; a seat never proxy-authors
   another seat's FROM. Authority-bearing orchestrator relays CC `<RUN_ID>.orchestrator-reviewer`.

## The product we're governing
<2–5 lines: what is being built, the locked high-level decisions, where the full design state lives,
what the shipped baseline is>

## Org chart + addressing
| Seat | Role | Address |
<operator / CTO=orchestrator-planner / VP=orchestrator-reviewer / pairs m-1..m-N>
The pairs are independent operator-relayed sessions (distinct models/lanes); never spawned as
subagents. The only sanctioned subagent spawn is read-only reviewer lenses.

## Domains (durable ownership)
| Pair | Domain | Owns | Consumes (collision edges) |
<one row per pair; foundational domains flagged; integration + arbitration reserved to CTO+VP>

## Persistent layout + naming
<the §3 layout, verbatim for this project; RUN_ID; address scheme
 `<RUN_ID>.orchestrator-planner|-reviewer`, `m-<n>.planner|.implementer`;
 DISPATCH_ID: boot = `<RUN_ID>-boot-<seat>`, work = `<cycle>-<phase>-<pair>`;
 INDEX.md columns: | time | phase | role | dispatch | parent | from | to | cc | status | file |>

## Where the rules live
<installed role skills + protocol.md; relay-lint path; external design-state paths>

## Roadmap & status
Sequencing: `ROADMAP.md`. Live status: `<team-root>/README.md`. Process: `<team-root>/CYCLE-PLAYBOOK.md`.
```

**Seat-role recasting under a truncated phase-band.** In an AUDIT+DESIGN band, `agent-pair-planner` is
the domain's **design-lead** and `agent-pair-implementer` is its **adversarial design-reviewer — NOT a
builder**. State this recast in the charter AND in every boot relay; the skill names say "implementer"
and sessions will drift toward building without the explicit hold. Reject any pair's
"READY FOR PROCEED-TO-PLAN" framing if the band has no PLAN.

---

## 5. The cycle template (repeat per initiative)

One **cycle** = one initiative run across the domains (`c1`, `c2`, …), scoped to the phase-band. The
skeleton, including the phases the standing team added on top of the stock protocol (marked ✦):

```
CYCLE c<n>
  1. AUDIT             orchestrator → pairs; planner + adversarial implementer each return
  2. AUDIT-RECONCILE   orchestrator reconciles → VP review → PROCEED-TO-<next> (or revise)
  3. DESIGN            orchestrator → pairs (GRILL_REQUIRED where semantics are unsettled /
                       cross-domain / hard-to-reverse); intra-pair review iterates rev0→rev1;
                       ✦ seed cross-pair COORD threads EARLY, with the agenda written down
  4. ✦ CONSUMER-REVIEW boot + dispatch the CONSUMING domains to review interface sketches
                       BEFORE lock (cheapest place to catch a writer with no reader)
  5. RECONCILE         orchestrator folds design + consumer findings → VP co-review (multi-round OK)
  6. REFINE (rev2)     pairs fold findings, RE-AFFIRM shared contracts to each other, re-review
  7. ✦ JOINT-LOCK      co-foundational domains lock TOGETHER; VP approves
                       (may gate on operator judgment items)
  8. ✦ OPERATOR GATE   judgment calls go to the operator as free-form prose — options +
                       recommendation + rejected alternatives, never a menu widget —
                       operator ratifies → orchestrator folds → VP close-confirm
  9. CLOSE             close declaration + pair stand-down + seal dashboard / ledger /
                       ARCHITECTURE.md + append the cycle's worked example to CYCLE-PLAYBOOK.md
```

**Invariant rules (non-negotiable; each was enforced the hard way):**

1. **Loop in the VP before executing any authority-bearing orchestrator decision.** Every broad SET and
   every reconciliation CCs the VP. Visibility, not approval — never wait on an approve to write the relay.
2. **Foundations lock first; consumers design against them.** Co-foundational domains lock together.
3. **Consumer review precedes design-lock.**
4. **Phase discipline is explicit and enforced.** State the band; reject self-advancement by pairs.
5. **The orchestrator actively polls.** Independent sessions stall silently — no "I'm done" ≠ done.
   Poke a quiet seat rather than waiting.
6. **Adjudicate intra-pair splits on evidence, not on which role said it.**
7. **Every substantive handoff is a lint-clean file relay.** No proxy-authoring; the operator is the
   transport.
8. **Never front-run a gate in a status doc.** The durable status must match the live decision.

**Later-cycle patterns (adopt when the situation matches):**

- **Bounded mid-cycle engagement** — a future domain can join a cycle for a reserved slice; write the
  "must not become the full design; the lock RESERVES your ownership to c<later>" guardrail into both
  the boot and the dispatch, and have the VP re-state it.
- **Fold-confirm round (rev2-lite)** — consumer findings routed back as a bounded *additive* fold,
  implementer-re-approved, with an "anything stronger → flagged micro-fold to the foundation owner"
  tripwire.
- **Terminal-cycle capstone** — when a cycle's domains are the last layer (no downstream to review),
  replace consumer-review with (a) conditional re-engagement of specific locked upstreams and (b) a
  blocking **integration-completeness capstone** (consume-graph acyclic + writer-backed; seams closed;
  locked invariants intact; deferrals recorded) folded into the lock co-sign.
- **Conditional-upstream-contract-check** — a locked domain can answer a new cross-domain question
  *without reopening* iff it reserved the extension point. Corollary for every design: **reserve
  versioned, present-but-null extension seams from day one.**
- **Adversarial pre-build review before leaving the design band** — an independent multi-lens ×
  multi-verifier fleet over the whole design-of-record, empowered to RETRACT prior certifications and
  return NO-GO. Treat a NO-GO as a bounded re-baseline with named root causes (which may include
  standing up a new domain), not a rewrite.
- **Lock-text hygiene** — when a lock reserves a space to a later owner, mark every example value as
  non-locking candidate vocabulary and state the operative lock rule once. After folding a resolved
  cell, grep-sweep the whole doc for its stale status strings.

---

## 6. Relay mechanics (the substrate)

**Headers, every relay** (then body, then the two trailer fields):

```
ROLE · PHASE · AUTHORITY · DISPATCH_ID · PARENT_DISPATCH_ID · RUN_ID · CEREMONY_TIER ·
EVIDENCE_TARGET · HUMAN_GATE_REQUIRED · FROM · TO · CC · SUBJECT
...body...
ACTIONS_GIT_REF: <ref | none — honest>
FINAL_GIT_STATUS_SHORT: <output | "unavailable — cwd is not a git repo (docs workspace)">
```

- File naming: `<PHASE>-<ROLE>-<YYYYMMDD-HHMMSS>.md` under `<team-root>/relays/<DISPATCH_ID>/`.
- Lint before every handoff: `python3 <skills-root>/tools/relay-lint.py <file>` (exit 0 = OK). Scoped
  lineage check: `--relay-root <team-root>/relays/<DISPATCH_ID>` (filter out INDEX.md — lint-exempt).
- `PARENT_DISPATCH_ID` is the load-bearing lineage field; an addressing mistake with correct PARENT
  reconciles transparently, the reverse does not.
- Cross-pair seams get their own COORD thread (`c<n>-design-m<a>-m<b>-coord`).

**Boot relay template** (one per seat; adapt the reference shape):

```markdown
## BOOT — initialize <seat> for RUN_ID <RUN_ID> (<domain>)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: <RUN_ID>-boot-<seat>
PARENT_DISPATCH_ID: <RUN_ID>-boot
RUN_ID: <RUN_ID>
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: <RUN_ID>.orchestrator-planner
TO: <seat address>
CC: operator, <RUN_ID>.orchestrator-reviewer
SUBJECT: BOOT — initialize <seat> for RUN_ID <RUN_ID> (<domain>)

You are <seat> — <role recast, e.g. "design-lead (Planner)"> on the <domain> domain. <1–3 lines of
domain framing + why it exists.>

Load **<role skill>** (+ `protocol.md`).
Read the team charter first: `CLAUDE.md` / `AGENTS.md` (auto-loaded in this cwd).
Read your domain charter: `<team-root>/domains/<dir>/README.md`.
Your pair partner (<recast, e.g. "adversarial design-reviewer, NOT a builder">) is <partner address>.

**THE ONE-LINE BOUNDARY:** <the single sentence separating this domain from its neighbors.>
**Scope of your first cycle (<cycle-id>) — <PHASE-BAND> only:** <bullets.>
**Constraints you inherit:** <VP boot constraints; claim boundaries; GRILL_REQUIRED; the phase hold.>
**Audit corpus:** <what to read/audit against.>

Sprint root: <team-root>/ (docs in cwd, never <code repo>). Relay root: <team-root>/relays/.
INDEX: <team-root>/relays/INDEX.md. relay-lint: <skills-root>/tools/relay-lint.py.
Current authority: **report-only onboarding.** This boot grants no work authority.
Acknowledge identity, loaded skill, reachable relay root + lint, and your boundary; then stand by for
the audit dispatch (DISPATCH_ID <cycle>-audit-<pair>).
ACTIONS_GIT_REF: none — report-only boot onboarding
FINAL_GIT_STATUS_SHORT: <honest value>
```

The boot relay is the **per-seat configuration layer**: generic skill + charter + boot relay fully
specify a seat, so the installed skills never need editing.

---

## 7. Pre-learned friction (apply from day one — each of these cost a round-trip)

- **relay-lint tripwires:** avoid `<word> = <token>` phrasings near reserved verbs (`merge = A` parses
  as a git claim — write `merge_decision`); a bare "backfill" reads as an action-claim; always set
  `ACTIONS_GIT_REF` / `FINAL_GIT_STATUS_SHORT` honestly.
- **INDEX.md is volatile** (async appends from pair sessions): re-read the tail immediately before
  appending; orchestrator runs a completeness + dangling check at each reconcile and backfills.
- **Pairs skip the intra-pair reconcile** (two independent passes ≠ reconciled): make the
  pair-reconcile relay an explicit dispatch deliverable and hold the gate on it (the F4 bar: two
  artifacts PLUS reconciliation, or one reconciled artifact).
- **Operator judgment items:** free-form prose with recommendation + alternatives + rejections; no
  `AskUserQuestion` menus. "Explore, then decide" — don't pre-bake one answer into a lock.
- **Operator directives via chat aren't channel-stamped:** cite "operator-directed by current session
  context" or have the operator drop a stamped `FROM: operator` relay.
- **Orchestrator edits to a pair's doc:** only under explicit operator direction, self-attributed,
  per-closure — never a precedent. Substantive changes go through the owning pair.
- **Declare-before-bind COORDs:** read the thread's current state before declaring anything "final."
- **Don't hardcode host paths** (skills root, linter) in durable docs; hosts change.

---

## 8. Init checklist (what "done" looks like for setup)

- [ ] Preconditions verified (skills + Superpowers installed, relay-lint exit 0)
- [ ] Parameter block filled; operator confirmed RUN_ID, phase-band, domain split, model lanes
- [ ] VP booted (boot relay hand-relayed; acknowledgment received)
- [ ] ORG-DECOMP dispatched, VP-reviewed, folded
- [ ] ROADMAP.md drafted, VP-reviewed (steps with observable exit tests)
- [ ] Charter written to CLAUDE.md; `AGENTS.md` symlink created; re-read from a throwaway session to
      confirm it auto-loads
- [ ] Layout created: team root, domain dirs + charters, `relays/` + `INDEX.md` header row, `boot/`
- [ ] Dashboard `README.md` seeded (current phase = SETUP → c1)
- [ ] Boot relays emitted for the c1 pairs (foundational first), all lint-clean, INDEX rows appended
- [ ] Per-seat relay pointers printed for the operator to hand-relay
- [ ] c1 AUDIT dispatches drafted (held until all c1 seats acknowledge boot)

From here, run cycles per §5, append each cycle's worked example to `CYCLE-PLAYBOOK.md` at close, and
keep the dashboard truthful to the live gate state — never ahead of it.
