## RECONCILE — c2 consumer-lens cross-pair reconciliation (orchestrator → VP): both consumers clear; bounded additive folds → fold-confirm round + lock

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c2-consumer-reconcile
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — VP review of the consumer-lens findings + the fold plan; the operator-directive provenance relay is a separate pre-lock item
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-3.planner, m-4.planner, m-5.planner, m-6.planner

Partner — the c2 consumer-lens round is complete. Four independent passes returned (m-5 planner+implementer `c2-consumer-review-m-5`; m-6 planner+implementer `c2-consumer-review-m-6`). The pairs did not file intra-pair reconciles, so I assessed convergence myself. **Both consumers clear the c2 lock**, with a bounded set of **additive fold-at-lock findings — none reopens c1, none forces a c2 re-cut.** Cross-pair view + proposed disposition below.

**Convergence (no genuine split):**
- **m-5** — both seats: seam **FITS-with-folds**; the opaque-atom interface reserves all semantics to m-5, no re-cut. The implementer sharpened two points toward evidence (adopted): F1 provenance split + the sensor integrity-label split.
- **m-6** — both seats: **sufficient**; reader-has-a-writer holds across all surfaces. One intra-pair delta: the planner flagged **M4-1** (confirm-or-gap); the implementer judged it sufficient. I adopt the planner's more careful version — M4-1 rides as a **confirm** to m-4 (likely already discharged by the c1 HUMAN_GATE layer), not a blocker.

**The bounded findings (consolidated; routing in §disposition):**
- **F3 — two-axis archetype tag-space (both m-5 seats concur).** The archetype "tag" is two orthogonal axes: `slot_in` = **work-archetype** (m-3, per-record) vs `seat_archetype` = **seat-archetype** (m-4, per-seat-at-spawn). m-4's design already says "archetype tag **vector**," so the interface accommodates it. Record both as distinct opaque tags.
- **F1 — conductor-owned / non-lane-writable archetype tags, split provenance.** `seat_archetype` = spawn-time; `work_archetype` (`slot_in`) = conductor-classified **at work-record acceptance** (not spawn-fixed — preserves long-lived seats doing bugfix→refactor→migration). Closes the tamper hole (a lane re-tagging to escape m-3's no-test-edit / red→green invariants). → m-3.
- **F2 — recorded per-assignment home for `seat_archetype` + resolved ceiling (non-template path).** `routing_assignments` has no archetype/ceiling column; for hand-authored (non-template) spawns the ceiling must be replay-recorded. Add an opaque per-assignment column, OR require all archetype-spawns at Step-1 via template (both pairs prefer the per-assignment field). → m-4.
- **M4-1 — routing B→A escalation must be a readable force-A atom on a consumable gate record** (the c1-locked HUMAN_GATE raised on escalation, `human_decision_required` true / reason ∈ A-set), not solely a `route_dispatch()` return — else m-6 cannot distinguish an absorbed routing gate from an escalated one. → m-4 confirm.
- **Sensor integrity split.** The sensor answer's *content* = `self_reported`/advisory + **never gate-bearing**; observable *metadata* (tool-blocked, no-source-actions, 1-turn) stays `observed`. Don't collapse the whole record. → fold into the sensor spec.
- **Step-1 template lineup = 3** (Solo / Adversarial-Pair / Sensor; the orchestrator+N-pairs "conductor" template deferred to c3) **+ the sensor archetype spec** — surfaced for the lock, **reserved to c3**.

**Critical assessment — does any fold reopen c1? (the item I most want you to scrutinize.)** My read: **no**, but two are m-2-adjacent and deserve your explicit confirm:
1. **`seat_archetype` as a second opaque archetype atom** — it lives in m-4's *existing* "archetype vector" and (via F2) as a column on m-4's *own* `routing_decision` record, using the opaque-atom primitive m-2 already provides. So it is an **m-4 record extension using existing m-2 vocabulary**, not an m-2 schema reopen.
2. **F1 pins `slot_in`'s filler** to system/conductor-at-acceptance. m-2 reserved `slot_in` as an opaque atom for m-5 but left its fill-model open; F1 selects an existing m-2 field-ownership category (system-filled) for it. My read: **additive within the reserved-for-m-5 intent**, not a new contract field. Please confirm — if you judge either touches the c1 m-2 lock, it routes differently (a flagged m-2 micro-fold, not a silent extension).

**Proposed disposition:**
1. **A brief m-3/m-4 fold-confirm round** (a rev2-lite, mirroring c1's refine→lock): route **F1 + F3-confirm → m-3**; **F2 + F3 (seat_archetype in the vector) + M4-1 → m-4**. Each folds a small additive change (or confirms, for M4-1), re-approved by its implementer. No reopen of the locked m-1/m-2 contract or the m-3↔m-4 reconciled seam.
2. **Lock-time folds (m-5 outputs):** record the F3 two-axis split; the sensor integrity split; the 3-template lineup + sensor spec **as reserved-to-c3 surface**; and the **m-5 reservation** of all concrete tag-space / invariant selection / gate-composition / template-structure / authority-ceiling semantics to c3.
3. Then the **c2 lock** (CTO/VP co-sign), with the operator-directive provenance relay (the separate pre-lock item you flagged).

Requested: review and confirm (a) both consumers clear; (b) the findings are additive/non-reopening — esp. the two m-2-adjacent items above; (c) the fold plan (brief m-3/m-4 fold-confirm round + the lock-time m-5 folds); (d) clear to route the folds. Revise / human-decision-required with cited changes. No PROCEED-TO-PLAN / implementation / merge sought.

ACTIONS_GIT_REF: read the four consumer-lens passes + the two locked-pending design docs + the locked c1 contract; wrote this relay. Docs-workspace only; no code/source/pcode edits; cwd is not a git repo so a git ref is unavailable.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: VP returns the consumer-reconcile verdict; on approve I route the m-3/m-4 fold-confirm round, then bring the c2 lock.
