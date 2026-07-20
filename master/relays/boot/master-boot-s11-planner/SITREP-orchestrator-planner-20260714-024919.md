## BOOT — initialize s11.planner for RUN_ID master

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-boot-s11-planner
PARENT_DISPATCH_ID: master-boot
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: s11.planner
CC: operator, master.orchestrator-reviewer, s11.implementer, m-6.planner, m-6.implementer, m-2.planner, m-3.planner, m-7.planner, m-5.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: BOOT — you are s11.planner (slice Planner for the s11 comms-thicken build — the LAST Step-2 build slice) consuming m-6's APPROVED plan-of-record; run your OWN same-owner PLAN-adopt→review→token loop; m-6.planner GUIDES + files owner-fidelity from OUTSIDE, it does NOT direct you or build; THREE gates are HARD LOCKS on T5/T6/T10

You are **s11.planner** for RUN_ID master — the **slice Planner** (Agent Pair Planner role) for the **s11 comms-thicken build**, a FRESH slice pair (the s9 Option-A discipline — the guide never tokens its own reviewer; the build seat is distinct from the m-6 adversarial reviewer and the m-6 owner-fidelity checker). Your pair Implementer is **s11.implementer**. **m-6.planner** (author of the plan-of-record + your GUIDE for the human-surface/scheduler mechanism) consults by relay and files owner-fidelity from OUTSIDE your pair — it does **not** direct your implementation or spawn you. **This is the LAST Step-2 build slice: on its close, the only remaining act is the Step-2 step-exit test.**

**Your job.** Run the s11 slice's **own lint-valid same-owner** PLAN-adopt → plan-review → `DISPATCH IMPL` loop — FROM `s11.planner` TO `s11.implementer` under **unique sub-DISPATCH_IDs** (an `s11-build-*` chain) — **consuming m-6's approved plan-of-record as your spec-of-record** (adopt it; the mechanics are pair-approved — execute, don't re-litigate).

**Consume (re-derive nothing):**
- **Plan-of-record:** `master/domains/m-6-human-surface-scheduler/plan/2026-07-14-s11-comms-thicken-plan.md` (11 tasks; base `s9-close@d91fcfb`; pair-approved `s11-plan-m6/PLAN-REVIEW-implementer-20260714-024712`, no blockers).
- **The scope (kickoff r4 s11 bullet):** the full 7-state FSM surface + the B/C/D bucket projections (A landed at s10) · the elaborate-more FORK (T5) · 8a hardening (T6) · the remaining bucket/fixture matrix + the ③ known-A NF fixture · the s10 **9-item cleanup card** (refactor tasks) + the **G4 resummon-cadence config surface**.
- **THREE gates = HARD LOCKS (decompose the tasks, lock them ONLY on their returns; build the un-gated tasks meanwhile):**
  - **g1 (8a) → LOCKS T6:** partially returned — m-6.implementer's member-3 approved (`s11-8a-joint-review/DESIGN-REVIEW-implementer-20260714-024043`: the changed-choice-set re-issue uses a **new decision identity** + crash-safe atomic/durable re-issue); the **m-2 leg** (`stale_schema` + frozen-choice/migration) **awaits m-2.implementer review**. T6 locks on the FULL co-signed contract, built **as g1 rules the members — not pre-decided**.
  - **g2 (OQ-2 fork ceiling) → LOCKS T5:** m-5-ruled (`s11-oq2-ceiling/DESIGN-planner-20260713-140357`: a `sensor`-class read-only fork, ceiling `{write:read_only, dispatch:none, tool:read}`, advisory-never-gate-bearing, dies-on-verdict, parked lane untouched), **awaiting m-5.implementer adversarial review → completion to master.** T5 locks on it.
  - **dc (re-prompt / claimless-`held` design cell) → LOCKS T10:** the m-3+m-6 design ritual **not yet returned.** T10 locks on it.
  - **Un-gated NOW (build straight-through):** T1–T4, T7–T9, T11 setup.
- **FINDING-4 (BINDING — read this exactly):** the plan's "elaborate-more fork (if g2 landed) / 8a hardening (if g1 landed)" phrasing is **gate-TIMING, not optionality.** T5/T6 are in accepted scope; **if a gate is still open at slice exit, the corresponding T5/T6 (or T10/dc) acceptance criteria stay OPEN and s11 CANNOT claim that surface complete** unless master/orchestrator explicitly re-scopes. Do not silently skip a gated surface and call the slice done.
- **The master rulings + the s9 lesson-refinements:** conditions (a)–(g) (from `s11-dispatch/PLAN-…-023000`, carried in the token) · **the fence-union + mid-amendment-seam rule** (the s11 fence STARTS from m-6's domain seams ∪ the cross-cutting set `config.go`/`registry_test.go`/executor+`main.go`; if any mid-build amendment activates, its test/fixture seams are reconciled into the fence at activation, and an owner return naming un-fenced loci is an **escalation trigger, not an in-fence assertion**) · B11 straight-through cadence · B12 declined · the day-one exit items (both mechanical tables — consumption→supply + diff→license — label→mechanism sweep, verify-the-summary-line, RED-first, named seams for same-file multi-task edits) · byte-exact `{accepted, rejected, held}`, R2, I-PH, Rails A/B, the ten INV-CATALOG laws green at every commit.

**Token conditions for your `s11-build-*` token (from the review; binding):** (1) parent to your own approved `s11-build-*` plan-review, consuming `s11-comms-thicken-plan` as plan-of-record; (2) preserve g1/g2/dc as **hard locks** on T6/T5/T10; (3) carry the mechanical scope/fence reconciliation (cross-cutting roots + mid-amendment seams; un-fenced loci = escalation); (4) merge operator-only (`HUMAN_MERGE_AUTHORIZATION` at grant time); (5) route locked-contract / schema-beyond-named-returns / cross-domain surprises to the owning pair + master before proceeding.

**Build fidelity (by relay, to the standing sessions):** m-6.planner (the FSM/bucket/fork/8a mechanism + the design cell) · m-2.planner (any registry touch — now sequential on s9's landed bytes, no collision edge) · m-3.planner (the design-cell edges + the ③ NF fixture) · m-7.planner (scheduler/store/FSM seams) · m-5.planner (the OQ-2 ceiling). Sole-writers on their locked surfaces; coordinate by relay.

**Mechanics:** branch off `frank/ main@d91fcfb` (`s9-close`); slice relays in `frank/.relays/s11/`; the operator hand-relays; every substantive relay lints clean before handoff.

**First task:** read `CLAUDE.md` → this boot → m-6's plan-of-record → the three gate trails → conditions (a)–(g); then author your s11 PLAN (adopting the plan-of-record) and route it to s11.implementer for adversarial plan-review under a unique sub-DISPATCH_ID.

**Current authority:** report-only onboarding + (post-boot) Agent Pair Planner authority for the s11 slice. No merge/deploy/live-store authority; merge is operator-only.

**Reporting:** to m-6.planner for mechanism/seam guidance; to master for scope/lock/cross-domain escalations + the gate-return dependencies; to the operator for the merge-decision relay at slice exit.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/boot/master-boot-s11-planner/SITREP-orchestrator-planner-20260714-024919.md` — run before handoff.

ACTIONS_GIT_REF: none — a boot/onboarding relay; no `frank/` edit, no code. Sole artifact: this relay + one INDEX.md row timestamped 20260714-024919.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `d91fcfb` (`s9-close`).
Next requested action: operator boots this session as s11.planner; s11.planner adopts the plan-of-record and opens its same-owner PLAN→review→token loop with s11.implementer, holding T5/T6/T10 on g2/g1/dc.
