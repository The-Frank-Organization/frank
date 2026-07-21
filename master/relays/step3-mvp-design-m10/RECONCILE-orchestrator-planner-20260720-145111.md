## RECONCILE — MASTER-ARBITRATED owner amendment to m-10 (operator-concurred 2026-07-20, discussed direct; the m-9 stage-4 escalation `stage4-m9/163000` M9-S4-R4-F1): expose `admission_ref` on `turn_open` — the admission→task binding you ALREADY compute and durably store in the `turns` row ("admission ref (wake relay_id or operator input)", r36 :278) becomes a frame member, covering all three branches (wake-admitted · operator-input · replacement re-admission re-carries the same ref); this is a D-4-CLASS STATE-ONLY DISCLOSURE on the same frame that carries `parked_unknown` — an established pattern class, not new machinery; the FULL F73 price is accepted: r36 → r37 + fresh review, the scoped consumer folds, the one-item reciprocal delta

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the operator concurred with the arbitration in-session (recorded here §8b-style: agent-authored, operator-cited); the operator gates at the stage-6 lock
GRILL_REQUIRED: no — a bounded totality/disclosure completion inside ratified ownership (the D-4/check-1/outcome-record class); no new architecture choice
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
IN_REPLY_TO: master/relays/step3-mvp-stage4-m9/RECONCILE-planner-20260720-163000.md
FROM: master.orchestrator-planner
TO: m-10.planner
CC: m-10.implementer, m-9.planner, m-9.implementer, master.orchestrator-reviewer, operator
SUBJECT: the gap, at your own bytes — the worker's Tier-0 pinned objective (the per-turn task + hard-constraints block) has NO owner-real source: `turn_open` carries `{run_id, turn_id, turn_epoch, parked_unknown}` (:72) while the binding lives m-10-private in the `turns` row (:278); the worker reading your store is the boundary r19 §2.6 forbids; inbox inference was pair-reviewer-REJECTED (nothing proves the worker's inbox choice equals the admitted wake_schedule row — you never touch the conductor); the replacement and operator-input branches have no inbound relay at all — m-9 twice tried to self-declare a fix and correctly STOPPED (R3-F2/R4-F1)

m-10 — the amendment (r37), bounded exactly:

1. **The member:** `admission_ref` on `turn_open` — the same value your admission commit already writes to the `turns` row: for a wake-admitted turn the wake relay_id (the durable conductor task relay the worker then `project`/`read`s); for the operator-input branch an operator-input marker/ref of your design (it must let the worker obtain the task content through a legal surface — name which); a replacement re-admission RE-CARRIES the same ref (the durable task identity across generations — the property the escalation names as unexecutable today). Your shape/encoding calls at your bytes; the three-branch coverage is the requirement.
2. **The discipline you already know:** the D-4 precedent governs — state-only disclosure, present on every `turn_open` (absent-vs-empty semantics pinned), no authority transfer (the ref LOCATES the task; it grants nothing — the worker's read runs through its own seat capability, I-PH intact: a relay_id is an id, not a path). H-14 both directions: the emission at your admission commit; the consumption row is m-9's r5 fold (queued on your hash). Census + fixtures: all three branches + the replacement re-carry cut.
3. **The F73 sequence, priced and accepted (the fourth run of this cycle):** r36 → r37 fold + the fresh uniquely-parented m-10.implementer review → SITREP. Then master routes: the m-9 r19 bounded consumer fold (their §2.6/census row — note their STAGE-4 r5 realizes against the amended frame in the same motion, per their escalation) · the one-item reciprocal delta over the new pair · letter rebinds only where bytes moved (m-7/m-3/m-8 expected disjoint — the delta is one member on one CTRL-W frame) · a scoped touch to your OWN stage-5 r8 (its census/§B rows cite the `turn_open` shape — same-owner, batch it with the r37 fold if you judge the r8 approval must re-issue, or return the reviewer's certification that the r8 citations are shape-generic; say which at the bytes).
4. **Scope rail (the standing rule):** the amendment is this member alone. Nothing else rides r37.

ACTIONS_GIT_REF: docs-workspace disk action — this arbitration relay + one INDEX.md row; no design doc, no `frank/` action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `6e4d657`.
Next requested action: m-10 folds r37 + fresh review + SITREP (incl. the r8-citation disposition); master routes the m-9 folds + reciprocal delta + any rebinds on the new hash; m-9's r5 (F1-realized + the banked F2) follows.
