## PLAN-REVIEW - s9 evidence-thicken plan rev7 must revise; marker removal is real, but legacy-v2 observation, timing semantics, and B-opaque authority remain unresolved

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-plan-m3-review-r8
PARENT_DISPATCH_ID: s9-plan-m3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the master scope flag is already filed; the remaining plan corrections are mechanical
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s9-evidence-thicken-plan
PLAN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s9-plan-m3/PLAN-planner-20260713-191000.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-7.planner, m-1.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must revise - rev7 genuinely removes marker inference and narrows timing, but its pre-v3 undeclared rule makes current Row-3 observations unreachable, two timing bytes contradict the narrowed contract, and master has not ruled on B-opaque scope

PLAN_REVIEW_VERDICT: must-revise

Rev7 closes r7 F1: the executable T1/T4/commit/acceptance bytes now remove VCS-marker inference and constrain `rootHealth` to worker/root health. It also chooses the honest narrow timing model and moves the unlicensed accepted-opaque row out of the build. The m-7 r2 owner leg is now technically approved and contains the required runtime handoff and version negatives. Three bounded issues remain before this PLAN is executable and authority-clean.

### Blocking Findings

#### F1 - the pre-v3 fail-closed row makes the promised current Row-3 behavior unreachable

Rev7 keeps all `lane_vcs` loci and the `RegistryEnv` handoff out of the fence until master activation (`plan:37-38`), yet T4's table dispatches an absent `lane_vcs` directly to `check-machinery-vcs-capability-undeclared` (`plan:207-219`). At `39474d0` every lane is necessarily pre-v3 and no observe component has a `LaneVCS` field. Therefore every claimless report takes the undeclared machinery row before git observation; the clean/matching `turn-attribution-unavailable` E0 row, mismatch veto, malformed-porcelain row, and existing section-13 regression suite cannot run.

That contradicts T4's claim that the restored Option-2 floor and mismatch observation land regardless, and Step 4's requirement that the existing `turn-attribution-unavailable` suite remains unchanged. Pin one legacy-v2 rule through the authority chain. Safe options include: attempt the existing canonical git observation for v2 and treat success normally while every failure remains machinery-fault/no opaque acceptance; or hold all of T4 until the master-activated v3 handoff. Do not silently convert every current report into a machinery failure while claiming the existing Row-3 floor remains live. If the chosen compatibility rule differs from m-7 r2's proposed v2-nil consumer behavior, include it in the master reconciliation request rather than deciding it within m-3.

#### F2 - timing and hostile-origin bytes still contradict the narrowed contract

The narrowed section correctly says executor timing is host-produced and conductor-canonicalized, not independently derived (`plan:156-159`). Two later bytes still state the old claim:

- The origin-table introduction says class, `MachineryFault`, **and timing are derived from origin, never from a returned token** (`plan:143`).
- Acceptance says executor timing is **re-derived** (`plan:285`).

Replace both with the precise split: conductor origins produce timing; executor origin supplies a bounded host value that is canonicalized and checked for tuple consistency; timing is diagnostic and never selects disposition.

The second hostile-origin test also conflicts with the matrix. The rule says a conductor-policy verdict carrying `check-machinery-*` becomes typed `check-machinery-verdict-origin-class-mismatch` with `MachineryFault:true` (`plan:154`). The test says it "does not acquire `MachineryFault:true`" (`plan:171`). Pin the intended result: the original refusal cannot be reclassified as machinery by its token, but the **invalid cross-origin verdict itself** must either fault through the typed machinery edge or remain the policy disposition. One outcome only, matching the table and terminal assertions.

#### F3 - B-opaque scope disposition and cross-domain activation are still pending master

m-7 r2 now has an approving technical countersign (`s9-vcskind-m7/DESIGN-REVIEW-implementer-20260713-162229.md`), so rev7's repeated statement that the amendment is currently must-revise is stale. The technical approval explicitly remains non-activating: master must reconcile the owner leg with the corrected m-3 consumer contract.

The separate `...-191500` scope flag asks master whether ledgering the accepted-opaque row is a permitted carry or an unratified reduction of the dispatched "opaque-lane detection + labeling" bullet. No master answer exists yet. PLAN approval would implicitly decide that question from this pair. Wait for the master relay, then either activate the ruled `lane_vcs` bytes and include B-opaque, or record the master-ratified carry and exact fail-closed s9 scope. Update the fence/order/T4 table to that ruling before requesting approval.

### Prior Findings Closed

- Marker-derived capability and opaque acceptance are absent from production instructions; `rootHealth` is worker health only and the text-sweep guard is explicit.
- `output-truncated`, origin families, tuple consistency, identity binding, signal-class derivation, and I-PH redaction remain closed.
- The timing contract is correctly narrowed in its primary section; no new executor/engine seam is invented.
- m-7 r2 technically closes runtime handoff, map cloning, v2/v3 boundaries, malformed forms, ceiling, and end-to-end fixture obligations.
- The Option-2 E0 invariant, T1/T2 mechanics, serialized-probe removal, master rulings, owner holds, and governance ledger remain otherwise closed.

### Revision Acceptance Bar

1. Pin current/pre-v3 observation behavior so the live Row-3 floor and mismatch tests are reachable without any opaque acceptance.
2. Make every timing and cross-origin test byte agree with the narrowed origin contract.
3. Consume master's B-opaque scope and lane-VCS activation ruling; do not infer it from technical approval.
4. Preserve all closed mechanics, then return PLAN rev8. `SCOPE_DIFF` remains after approval and before any token.

ACTIONS_GIT_REF: none - read-only PLAN-REVIEW of `s9-evidence-thicken-plan` rev7; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` edit, branch, commit, implementation token, or merge
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0

Next requested action: m-3.planner reconciles legacy-v2 observation and timing bytes, then folds master's B-opaque/lane-VCS ruling into PLAN rev8; implementation remains held.
