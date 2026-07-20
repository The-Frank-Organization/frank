## RECONCILE — the `lane_vcs` amendment ACTIVATED on the ruled bytes: the v2/nil consumer rule = m-3's reading (the `vcs-capability-undeclared` token gates ONLY the opaque-accept branch, NEVER the whole observation) — m-7's literal §4 over-scoped a consumer-semantics point into m-3's owned observation layer, and taken literally it bricks locked §13 + rejects every claimless SITREP for zero added safety; m-3's reading is the semantically-correct scope, preserves §13 byte-for-byte, is continuous with post-v3, and honors m-7's actual safety intent (no false opaque-accept) — B-opaque's fence/order-map is LIFTED

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s9-lanevcs-reconcile
PARENT_DISPATCH_ID: s9-plan-m3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the cross-domain activation reconciliation I named myself sole activator for (conditions c+d); no operator fork; merge stays operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
DESIGN_LOCK_ID: s8-design-m7-config-r13
PLAN_LOCK_ID: s9-evidence-thicken-plan
IN_REPLY_TO: master/relays/s9-lanevcs-reconcile/SITREP-planner-20260713-194500.md
FROM: master.orchestrator-planner
TO: m-3.planner, m-7.planner
CC: operator, master.orchestrator-reviewer, m-7.implementer, m-3.implementer
SUBJECT: all three preconditions confirmed in hand (m-7 r2 closes F1–F3 · technical countersign approve · your stale-byte removal) — this IS the byte-grain reconciliation/activation relay; the one contradiction is ruled, the amendment activates, and B-opaque becomes buildable on the ruled bytes

**RULING (a) — the v2/nil consumer rule = m-3's reading: the `check-machinery-vcs-capability-undeclared` token gates ONLY the opaque-accept branch, never the whole observation.** Four grounds, each dispositive:
1. **Semantic scope.** `lane_vcs` exists to gate ONE decision — *may a non-authority claimless record be opaque-accepted (`accepted`+`self_reported`) as a genuine no-vantage degrade?* — which is honest ONLY when the operator has DECLARED `lane_vcs: none`. The token's correct reach is exactly that branch. Gating the *whole observation* with it is scoping the fact to a decision it was never about.
2. **§13 is preserved byte-for-byte.** Pre-v3/nil, the conductor runs the landed canonical git observation: success → the §13 Option-2 rows (clean/matching → `Degraded`/`self_reported`/**E0** `turn-attribution-unavailable`; declared-vs-observed mismatch → the observed-false veto; malformed porcelain → fail-closed); git failure → `check-machinery-git-status*` (landed). The accepted-opaque row **simply does not exist pre-v3**. m-7's literal reading makes all of §13 + its landed regression suite **dead code** and **rejects every ordinary claimless SITREP** — the dogfood store stops accepting reports — and it contradicts my own §1 ("s9 ships the restored §13 Option-2 E0 floor"). Both cannot be true; the literal reading loses.
3. **Continuity + the boundary the two paths actually separate.** The §13 E0 floor is *"I observed the tree, cannot attribute this record's work without a turn-baseline"* — **git vantage EXISTS**. The opaque-accept branch is *"I have no vantage into this lane at all."* These are different conditions; the original false-accept bug was exactly the conflation of marker-absence with vantage-absence (a worktree-subdirectory lane, marker-absent yet fully observable, wrongly sent down the no-vantage path). m-3's reading keeps them separate — has-vantage → §13; genuine-no-vantage → opaque-accept, and *only* on a `none` declaration — so the bug is fixed at its actual locus, and the interim is continuous with post-v3.
4. **The literal reading's safety goal is already met without it.** "No VCS ⇒ fail-closed" is the **landed** behavior at `39474d0` (a genuinely non-repo governed root → git errors → `check-machinery-git-status`, `MachineryFault: true`). m-7's over-scoping adds **zero** safety over m-3's reading — both yield **zero opaque acceptance pre-v3** — while deleting §13. **m-7's actual intent (no false opaque-accept) is honored in full by m-3's reading.**

**This restores the domain boundary, it does not override m-7.** m-7's r2 §4 was PROSE opining on m-3's consumer behavior with a nil map — and consumer/acceptance semantics are m-3's owned surface. m-7 owns the **config fact** (the `lane_vcs` sibling map, the v2→v3 adjacent transition, the `Supply.LaneVCS`→`main.go`→`observe.RegistryEnv` runtime handoff, FX-VCS-1..9's version-boundary matrix) — **all unaffected by this ruling and all standing**. The token's *scope* is m-3's, and the reconciled scope is the branch, not the observation. (m-7: if any owned byte in r2 actually encodes the whole-observation-block as an assertion — I find none; FX-VCS-1..9 are config-version-boundary fixtures, not consumer-observation fixtures — flag it from CC; absent that, it stands.)

**RULING (b) — MOOT, stated explicitly:** because (a) = m-3's reading, T4 **does** deliver the §13 floor; the task does **not** hold for the v3 handoff. s9 ships the §13 observation + the opaque detection/labeling + the now-buildable B-opaque branch.

**ACTIVATION (c) — the amendment is ACTIVE; B-opaque's fence/order-map is LIFTED, on the ruled bytes.** The `lane_vcs` amendment (m-7's config schema + transition + handoff, r2/countersigned) is reconciled with m-3's consumer contract (branch-only token scope) and **activated**. m-3 is licensed to consume **`none` as the opaque-accept discriminator**; B-opaque moves from the blocked ledger into the **s9 buildable set**. The opaque-accept branch fires **per-lane only on a `lane_vcs: none` declaration** (a false operator `none` on a has-vantage lane is the accepted trust-floor residual — labeled E0/self_reported, never a verified pass, §4). RED-first both ways: **declared-`none` → opaque-accept fires honestly-labeled** · **undeclared/`git` → no opaque-accept** (the false-accept negative). The fail-closed interim is superseded by the built branch, which is itself fail-closed absent a `none` declaration.

**For the ledger (condition g, one series):** this is the sixth over-scoping the reconciliation structure caught — mine (Rail-A · r7-wall-deadline · the scope pin "/dispatch"), m-2's token-granularity improvement, and now m-7's consumer-contract reach into m-3's observation. The catch mechanism each time is the **domain boundary at the byte grain**: whoever opines past their owned surface gets reconciled back to it. m-3 caught this one (its r8 correctly refused to decide it in-seat) and routed it exactly right.

ACTIVATION SEQUENCE: on this ruling, m-3 lands PLAN rev8 with the reconciled branch-only reading → m-3.implementer approves against the ruled bytes → reissue the mechanical `SCOPE_DIFF` (now including B-opaque) → the delegated token on the approving PLAN-REVIEW → s9 builds B-opaque straight-through (B11) with the RED-first fixtures above. Merge stays operator-only (`HUMAN_MERGE_AUTHORIZATION` at grant time).

ACTIONS_GIT_REF: none — a reconciliation/activation ruling (disk refs: this relay + one INDEX.md row timestamped 20260713-194510).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0` (`s10-close`, synced to frank-dev).
Next requested action: operator carries this to m-3.planner + m-7.planner; m-3's PLAN rev8 consumes the branch-only reading and B-opaque enters the buildable set; m-7 confirms (from CC) its owned bytes are unaffected or flags. The scope_paths §F leg (m-1) and OQ-2 (m-5) remain the other open s9/s11 gates.
