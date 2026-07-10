# m-6 — Human Surface & Scheduler

**Pair:** design-lead `m-6.planner` + adversarial design-reviewer `m-6.implementer`.
**Engagement:** **c3 full domain design** (the design-of-record). m-6 ran the c1 + c2 consumer-lenses against the
locked foundation/runtime contracts but never designed its own domain — c3 is its full design, the **last Step-0
design domain**, co-designed with m-5 (lead).

## Owns (the durable domain)
Email-governance + meeting-collaboration **surfaces**, **gate→email buckets**, the **Owner Decision Brief** (ODB),
and **scheduler park/wake**. The relay/inbox governance graph as first-class operator-facing comms.

## c3 scope (VP-approved — `c3-decomp` 20260630-051448)
The full m-6 design-of-record:
- **Gate→email buckets** — the mapping from m-2 HUMAN_GATE / `gate_category` (A/B set, §J) to human-surface
  buckets; the bucket taxonomy + per-bucket surface behavior + the operator decision-capture path.
- **Owner Decision Brief** — the operator-facing decision surface (port of agent-scripts' ODB): what a gate
  presents to the human + the decision-capture/return path (the J1 `hold_and_resummon` surface).
- **Email-governance + meeting-collaboration surfaces** — the design-of-record (mechanism builds Step 2, full
  email-client UX builds Step 4). Apply the jcode negative-look (no GUI noise/filler) + the codex positive-look.
- **Scheduler park/wake** — park a gated/awaiting lane, wake-on-reply; the away-mode external-inbox bridge
  (egress-gated via m-3; §J1 away-mode is the egress chokepoint's first external send).
- **Seam with m-5 (Seam A/B)** — consume m-5's per-archetype **human-mode vocabulary** (declare-before-bind, VP
  F2) and host the **interjection surface** (steer / side-question / interrupt; m-5 owns the sensor archetype,
  m-4 the routing).

## Consumes (locked upstream — do NOT reopen)
- **m-1** addressing graph (who a gate routes to) + the §J operator-judgment defaults.
- **m-2** HUMAN_GATE fields + `gate_category` + the monotonic HUMAN_GATE floor.
- **m-3** egress/content-safety gate (for any external send / the away-mode bridge).
- **m-4** routing (for the interjection side-question archetype).

## Roadmap mapping (`ROADMAP.md`)
m-6 design-of-record = Step 0 (now). m-6 *mechanism* (inbox/outbox + scheduler) = Step 2. Full *email-client UX*
= Step 4. c3 designs the contracts now; build is later ("designed early, executed later").

## Status
- **c3 LOCKED** (`c3-lock` VP co-sign `RECONCILE-orchestrator-reviewer-20260630-191315`; orchestrator lock `190627`/`191525`) — pair-**approved** (`DESIGN-REVIEW-...-133839`) + seam **converged** (`123022`⊕`131856`) + **Seam C RESOLVED (A)**: the away-token cell folded over **m-1-owned mint/verify** (`182600`/`183008`); the COORD-182600 build carries recorded (§12); GRILL_LOCK §9 folded. **All cells closed.** Post-lock review-driven consistency folds — **c4** CQ-3/4/4b + CQ-6 (`024620`), **c5** claim-sweep + decisions ③/④/⑤ (`134742`–`134745`), **c6** cleanup F3/F6/F7/F8/F9 (`c6-fix-m-6`) — folded into the doc §13 fold-log; **lock invariants unchanged**. No PLAN/IMPL (c3 terminal = design-lock).
### History (superseded by the c3 lock — chronological trail, all gates since cleared)
*(These bullets are a past-tense progression log; the current state is the `c3 LOCKED` bullet above. Nothing below is a live/pending gate.)*
- **c3 DESIGN doc r2 — both must-revise blockers folded (historical)** (`c3-design-m-6`) — `design/2026-06-30-human-surface-scheduler-design.md`. Implementer must-revise `131702`: **F1 (seam stale)** resolved by m-5's retraction of `125604` → converged **FOUR-class `123022`⊕`131856`** (doc reverted; the spurious `131747` bind-confirm WITHDRAWN); **F2 (egress A/D double-assignment)** folded an explicit `failing_edge`+stage precedence rule (§2). rev-2 was **subsequently approved** (`DESIGN-REVIEW-implementer-20260630-185424`); the then-**Held** away-mode token-bridge cell (OQ-1) was **resolved Seam C = A** (m-1 answer `182600`/`183008`). Both gates cleared; the doc is now **c3 LOCKED** (see above).
- **c3 full-domain audit RECONCILED (historical)** (`c3-audit-m-6`) — planner `053651` + implementer `053053` independent passes reconciled into `RECONCILE-planner-20260630-054107.md` (F4). Verdict: **promote-and-bind** (thin local-first projection over locked m-1..m-4, not a new gate system). One divergence resolved (away-mode inbound verdict trust → DESIGN-phase confirm-or-gap with m-1). Two c3-DESIGN COORDs (m-5 human-mode vocabulary + m-1 inbound verdict-token bridge) **were subsequently accepted and run**; c3 DESIGN completed and **locked**.
- **c3 full-domain audit dispatched (historical)** `c3-audit-m-6` (20260630-051950) — m-6 had no prior design-of-record at the time.

## Layout
- `audit/` — c3 AUDIT artifacts.  `design/` — c3 DESIGN docs + grill locks.
