## SITREP — CQ-closure COORD plan (full-pair rigor): 3 scoped clusters re-engaging 5 stood-down pairs; VP review before I fire

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: c4-cq-coord
PARENT_DISPATCH_ID: c4-design-m-7
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — operator already directed full-pair rigor; this is the routing decomposition, VP-reviewed, operator on CC
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-7.planner, m-7.implementer, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-5.planner, m-5.implementer, m-6.planner, m-6.implementer

Partner — m-7 is **DESIGN-COMPLETE r3, pair-approved** (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md`; grill-locked `c4-grill-m-7`; seam matrix biting; F8 semantic claim-sweep clean; two adversarial revise cycles caught + fixed the sole-egress overclaim twice). It is **holding at the CQ gate** — correct pair discipline, no CQ self-resolved, no locked contract reopened. Design-LOCK is gated on **8 CQs** (CQ-1/2/3/4/4b/5/6/8; CQ-7 non-locking). Per your audit-reconcile co-sign, I own firing the targeted COORDs.

**Operator has directed FULL-PAIR RIGOR** for the CQ closures (upgrade from the targeted-parallel baseline): each cluster re-engages the owning **full pair** (planner + adversarial implementer), so every contract change is held to the same adversarial standard that caught m-7's own overclaims. Before I fire — and re-engage 5 stood-down pairs (~10 sessions) — I bring you the **routing decomposition** for review, same decompose→review→fire spine every standup has run.

**Proposed 3 scoped COORD clusters** (each SCOPED to its CQs — confirm the m-7 reading OR produce the missing artifact, **without** reopening the rest of the locked contract; **not** a broad c1/c2/c3 re-open):

- **COORD-1 · `c4-cq-gateconfig` — the schema/gate/config interlock.** CQs: **CQ-2** (decision-② fail-closed fold — this *is* re-baseline step (c), sequenced here since the owners are in the room), **CQ-3** (pure-judgment A-floor table — *produce*: it doesn't exist yet), **CQ-4** (terminal-state token set: `bounced`→`rejected` + HELD registry home + bucket-D naming — *produce/settle*), **CQ-4b** (trusted-config artifact composition — **CTO-arbitrated**, I supply the ruling from the locked m-2/m-3/m-4/m-6 config inputs + m-7's §7 load/integrity requirements). Re-engage **m-2 + m-3 + m-6 full pairs** + CTO. The heavy, interlocking cluster (all touch the gate/config surface — your sanctioned grouping).
- **COORD-2 · `c4-cq-m1` — the m-1 store-provenance + phase-split.** CQs: **CQ-1** (phase-split required-set — m-1 §5 Step-1 boundary + m-2 `required_when`; m-2 co-signs from COORD-1), **CQ-6** (persisted seat-binding table + away-token sibling-burn scope — m-6 co-signs the token edge), **CQ-8** (INDEX derived-authority — confirm m-7's §4 reading of m-1 §6's "append-only INDEX.md": layout unchanged, only crash-recovery provenance changes). Re-engage **m-1 full pair**. Mostly confirm-type (m-7 proposed the readings; m-1 ratifies-or-corrects).
- **COORD-3 · `c4-cq-slotin` — slot_in classification ordering.** CQ: **CQ-5** (m-7 proposes classify **post-gates / pre-observe** in the commit pipeline; the done-predicate reads it). Re-engage **m-5 full pair**; **m-3** (already up in COORD-1) contributes the observe-pipeline-ordering half. Mild.

**Sequencing:** fire COORD-1 + COORD-2 in parallel (the heavy + the m-1 blockers); COORD-3 rides alongside (m-3 shared with COORD-1). Each closes to a scoped confirm-or-produce artifact that folds into the m-7 design-lock package. m-7 stays holding until all `blocks: design-LOCK` CQs close (or I explicitly carry one non-locking).

**Where I want you to push (VP asks):**
- **Q1 — the clustering.** Is the 3-cluster split right, or do you re-cluster? Two cross-cuts I want your eye on: **CQ-1** spans m-1 (§5 boundary) + m-2 (`required_when`) — I put it in COORD-2 (m-1-led) with m-2 co-signing from COORD-1; and **CQ-5** spans m-3 + m-5 — I put it in COORD-3 (m-5) with m-3 contributing. Lead-plus-co-sign, or should either be its own two-pair COORD?
- **Q2 — CQ-2 placement.** I fold the decision-② fix (= re-baseline step (c)) into COORD-1 because m-3 (fold owner) + m-2 (field-home) are both re-engaged there. Concur, or should step (c) run as its own dedicated fold thread (it's the "fold the 5 decisions" re-baseline work, of which decision-② is one)?
- **Q3 — CQ-4b arbitration shape.** I'll draft the config-composition ruling as CTO (one artifact vs per-domain files + the load contract) from the locked config inputs + m-7 §7, and route it INTO COORD-1 for the m-2/m-3/m-6 authors to build against. Right shape, or arbitrate it as a standalone CTO ruling first, then feed all three COORDs?
- **Q4 — scope-hold.** Confirm each COORD is SCOPED to its CQs (confirm-or-produce, no locked-contract re-open); the re-engaged pairs answer the bounded question and stand back down — this is the conditional-upstream-contract-check at scale, not a reopened cycle.

On your concurrence I write the three COORD dispatches + my CQ-4b ruling draft, and the operator spins up the pairs. If you re-cluster, I fold before firing.

Not authorized / not claimed: no CQ resolved by this SITREP, no pair re-engaged yet (the dispatches follow your review), no design-LOCK, no PLAN, no code/source/`pcode/`, no spike, no locked-contract reopen, no operator decision reopened.

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py master/relays/c4-cq-coord/SITREP-orchestrator-planner-20260702-011058.md` — OK
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-coord` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: read the m-7 design-of-record (r3) + the CQ ledger (§15); wrote this COORD-plan SITREP + appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit, no CQ resolved, no pair re-engaged.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: VP review of the CQ-closure COORD decomposition (Q1–Q4); on concurrence I write the 3 COORD dispatches + the CQ-4b ruling draft and the operator re-engages the pairs.
