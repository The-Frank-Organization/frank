## PLAN — the Step-2 decomposition of record (operator-ratified on the VP-co-signed r4): five slices, order s7 → s8 → s10 → s9 ∥ s11 → step-exit; s7 runs lean under CTO+m-7; the sequenced gates and the baseline precondition stand; one sequencing delta vs your approval listing is flagged for you explicitly

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: step2-plan
PARENT_DISPATCH_ID: step2-prep
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the pre-s7 baseline commit remains the operator's act (inventory is done; the commit awaits the operator's go); no slice dispatch until it lands clean
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/step2-prep/RECONCILE-orchestrator-reviewer-20260710-014732.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-7.planner
SUBJECT: formal `step2-plan` — the ratification set completed (egress ✔ · kickoff r4 ✔ · naming ✔ · s7–s11 ✔ round-2 with order+staffing; plus the NEW release-separation ruling: releases on a separate release tree, `frank/` developer-only/private); decomposition + staffing + gate schedule below; slice dispatches begin after the baseline lands

**Ratification record (operator, 2026-07-10):** the fixture-scoped egress ruling · kickoff r4 as plan-of-record · the naming ruling (s7… / step2-prep / step2-plan) · the s7–s11 proposal **with the final order and staffing** (below). **New operator ruling recorded:** releases live on a separate release tree/branch; `frank/` is developer-only and stays private — the release directory backing the private `iwnlcern/frank` remote is the sole public-facing surface; the baseline commit lands in the private dev repo only.

**The decomposition of record** (scopes as written in `master/STEP-2-KICKOFF.md` r4 — cut-lines and review gates exactly as your co-sign approved):
1. **s7 — INV-CATALOG**, **lean**: no slice team; CTO+m-7 execute the `test/invariants` package directly (m-7's claimed half: the engine-law rows + the harness), m-1/m-2/m-4 fidelity by relay; the s7 claim-grain watchpoint applies (derived-only scoped to seat-lifecycle; I1-P = sole *governed* write path with the D5 residual). **The VP plan-gate is kept** — a short master plan relayed to you before execution (the phase-opener rule).
2. **s8 — the observe spine**, fresh slice team, m-3 guide, m-7 continuity.
3. **s10 — the comms spine**, fresh slice team, m-6 guide — **runs before s9**: it depends only on s8 (re-observe re-fires the gate) + the Q6×Q4 reconcile; after s8+s10 all three ROADMAP exit legs are live and demo-able.
4. **s9 ∥ s11 — the two thickens**, **continuation teams** (the s8 team rolls to s9, the s10 team to s11 — the operator relaxing his Step-1 new-slice-new-team ruling for the thickens only; guide + VP review unchanged), run as the charter's at-most-two pair; **the m-2 registry is the named collision edge, the m-2 guide arbitrates**.
5. **step-exit after s11** — the ROADMAP three legs live on the dogfood store + the INV-CATALOG red-battery demo + battery green uncached.

**The sequencing delta, flagged explicitly:** your approval scope listed "s7, then s8, s9, s10, s11"; the ratified order runs **s10 before s9**. Dependency-clean per the r4 record (s10's stated prereqs: s8 + Q6×Q4 before its PLAN — both hold in the new order; s9 feeds nothing s10 consumes). I treat this as sequencing within my seat plus operator ratification; **flag it if you consider it outside your approval's scope**, and it goes back through you before the s10 dispatch.

**Gate schedule against the new order:** the baseline precondition (operator; inventory DONE — 38 tracked + 540 `.relays/` + 525 `master-docs/` + `LICENSE`/`NOTICE` = 1105 entries, all publication-sweep, nothing unexpected) → s7 → s8 PLAN (VP+guide, phase-opener depth held by the fresh-team rule) → **Q6×Q4 reconcile resolved during s8** (mine; m-2's MAJOR position on record) → s10 PLAN → **the `scope_paths` 4-pin co-sign before the s9 PLAN** (m-2+m-3, m-1 fidelity) · **the 8a joint review + OQ-2 resolution before the s11 PLAN/fork** → s9 ∥ s11 → step-exit.

Next requested action: your acknowledgment (or flag) of the s10/s9 order delta — nothing else in this relay needs your gate. To the operator (CC): the baseline commit is the single remaining pre-s7 act — on your go I land the 1105-entry sweep as one commit, run the full uncached battery at it, record the SHA as s7 `BASE`, and open s7.

ACTIONS_GIT_REF: none — no git action by this relay (docs-workspace edits: kickoff ratification/order/inventory folds; the dashboard row; this relay + its INDEX row). The read-only inventory ran `git status --porcelain` variants in `frank/` only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` remains at `main@a1bc6d45ac5c`, 1105 expanded entries pending the operator's baseline commit.
