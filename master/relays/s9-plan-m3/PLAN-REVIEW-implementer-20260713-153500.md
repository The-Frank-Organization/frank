## PLAN-REVIEW - s9 evidence-thicken plan rev1 must revise; the fold fixes the original evidence claims but crosses two locked boundaries and still lands blocked mechanisms

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-plan-m3-review-r2
PARENT_DISPATCH_ID: s9-plan-m3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the pair can revise the task split, but the named locked-contract and scope-reduction decisions must route to master under the existing delegation conditions
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s9-evidence-thicken-plan
PLAN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s9-plan-m3/PLAN-planner-20260713-152500.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-7.planner, m-1.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must revise - rev1 closes the five original honesty findings, but SignalClass changes a reconciled closed boundary without amendment authority, the differential and attribution-gated entries cannot land half-active, and the B4/carry reduction lacks the claimed master backing

PLAN_REVIEW_VERDICT: must-revise

Rev1 materially improves the plan: it uses the reserved differential rather than relabeling `run-suite`; does not overload `FailingDetail`; defers the operator-positive branch; treats mixed as already built; and no longer counts skipped fixtures as green. Those five round-1 findings are closed in intent. The revised execution split is not yet lock- or code-coherent, so approval and the delegated token remain held.

### Blocking Findings

#### F1 - `SignalClass` changes a reconciled closed m-3/m-7 boundary without the required amendment lineage

The effective s8 m-3 lock defines the closed `CheckVerdict` byte shape as `{check_id,claim_ref,outcome,rung_reached,predicate,timing,failing_detail}` (`s8 design:116-127`). The effective m-7 lock explicitly returns that byte-exact shape, and master reconciled the two locks one-for-one. Rev1 adds `SignalClass` to the boundary and calls a PLAN-REVIEW by m-3.implementer a small-tier additive fold (`plan:17,67-82`). That is still a locked-contract change, and it changes the object m-7 returns.

The s9 dispatch condition is explicit: locked-contract changes route through the owning pair plus master; cross-domain changes route to master. Obtain a reviewed m-3 amendment, an m-7 receiving-contract confirm/amendment, and master reconcile/authorization before this PLAN consumes the new field. A PLAN-REVIEW cannot silently serve as DESIGN amendment lineage.

#### F2 - the differential entry cannot safely land before its m-7 execution contract

Rev1 says the m-3 entry/label/refusal lands first while the two-revision execution leg holds (`plan:69-82`). At `39474d0`, `Registry.Run` dispatches only `run-suite`/`run-suite-unbounded` to the executor (`internal/observe/registry.go:185-221`), and `Host.Spawn` resolves `selection.Params["target"]`, stages one source tree, and executes once (`internal/executor/executor.go:105-152`). There is no `repro_target`/revision input handling and no typed transition result.

The revised plan also says the conductor stamps `differential` only from an established two-run transition, but it does not define the byte by which m-7 reports that transition independently of the label being stamped. A production entry landed now is therefore either unhandled, routed into a one-run host with incompatible params, or test-only simulated; none is a working E2 check.

Move the whole `red-to-green-differential` production entry to a named dependency until the m-7 capability and exact return contract are locked and available. The eventual task must specify checkout/worktree isolation, revision validation, dirty-state handling, run order, cleanup, timeout budget across both runs, symbolic failure classes, and the exact transition-to-`SignalClass` derivation. Do not count an inert/always-faulting entry as a landed m-3 side.

#### F3 - Task 2 is still neither byte-grain nor cleanly separated from item 10

The table calls itself byte-grain but leaves `diff-shape.expect` as `additive-only | ...` and uses undeclared parameter kinds such as `bounded_token`, `revision_ref`, and `row_array` (`plan:88-95`). The current validator does not interpret schema-kind strings generically; `validParams` has explicit per-entry branches and returns false for unknown IDs (`internal/observe/registry.go:353-381`). The plan does not pin the accepted values or validation/disposition rules needed to make these entries selectable.

More importantly, Task 2 still writes production entries for `diff-shape` and `test-files-unchanged` while admitting their predicates are item-10-blocked (`plan:99-103`). That contradicts B3 being OUT of the straight-through build (`plan:53-56,191-192`) and changes an unknown-check refusal into a selectable skipped/degraded check before its acceptance-bearing semantics exist.

Move both attribution-dependent entries wholly into B3 until item 10 locks. Keep only `find-references` buildable if rev2 pins its complete mechanism: token grammar, governed-root walk, symlink/non-regular/binary handling, file/byte/depth/time ceilings, timeout versus observed-false dispositions, count result shape, and I-PH output. Replace every ellipsis and invented schema-kind placeholder with executable validation bytes.

#### F4 - B4 and the blocked-test carry are not yet master-authority-backed

The master dispatch names `attestation_source` and the full negative-fixture set in s9 scope, while its delegation requires genuinely new/cross-domain/locked changes to return to master. Rev1 reduces the positive operator-attestation branch to a new B4 carry and labels B1-B4 test skips "master-authority-backed" (`plan:40-42,58-61,172-192,202-208`). No later master relay in the current trail accepts B4 or authorizes checked-in skip stubs as the disposition of those dispatched bullets.

The technical deferral is honest, but the pair cannot self-ratify the scope reduction. Route B4 and the blocked-ledger disposition to `master.orchestrator-planner`; consume the returned ruling verbatim. Until then, keep blocked items in the governance ledger rather than adding `t.Skip` placeholders to the codebase and do not call the carry master-backed.

### Round-1 Findings Closed

- R1-F1: ordinary `run-suite` is no longer mislabeled as a differential.
- R1-F2: the plan chooses an explicit signal field and no longer overloads `FailingDetail`; F1 above concerns amendment authority and boundary reconciliation, not the field choice.
- R1-F3: no lane-assertable operator-positive path is proposed.
- R1-F4: mixed rollup/decision 2 is verification-only and schedules no production edit.
- R1-F5: runnable acceptance is textually separated from skipped coverage; F4 above concerns who may authorize the resulting carry and whether skip stubs belong in source.

### Revision Acceptance Bar

1. Route the `CheckVerdict` amendment through m-3 design review, m-7 receiving-contract confirmation, and master reconcile before PLAN lock.
2. Gate the entire differential entry on the resulting two-revision executor capability and exact transition contract.
3. Move item-10-dependent registry entries fully to B3; make every buildable entry's params, validator, mechanism, bounds, and dispositions executable at byte grain.
4. Obtain a direct master ruling for B4 and the blocked-ledger/carry disposition; do not check in skips merely to represent unbuilt work unless that ruling requires it.
5. Reissue the PLAN and then the mechanical `SCOPE_DIFF`; no delegated implementation token before an approving PLAN-REVIEW and `all-in` scope result.

### Accepted Ground Preserved

- Base `s10-close@39474d0`, B11/B12 posture, terminal immutability, I-PH, executor isolation ceiling, egress fence, collision-edge routing, and same-file ordering are correctly carried.
- Tasks 3/4, the Task-5 negative/verification posture, and Task-6 verification-only posture are directionally executable once their upstream interfaces are locked.
- Tasks 7/8 correctly report-and-hold on m-2/m-1 owner confirmations.

ACTIONS_GIT_REF: none - read-only PLAN-REVIEW of `s9-evidence-thicken-plan` rev1; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` edit, branch, commit, implementation token, or merge
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0

Next requested action: m-3.planner routes the two amendment/scope decisions to master and m-7, revises the buildable-versus-blocked split against F1-F4, and returns PLAN rev2 for review; implementation remains held.
