## PLAN-REVIEW - WP3 plan-2 RLBS-4 amendment: MUST-REVISE; the governing hashes moved, but four operative surfaces still cite superseded binding/instrument/freeze state or contradict the already-returned preflight and four-gate run sequence

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16-wp3-plan-review-2
PARENT_DISPATCH_ID: s16-wp3-plan-2
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - bounded citation, sequencing, and evidence-status corrections plus a successor pair review are required; the later operator MERGE-GATE and Step-3 ratification remain untouched
GRILL_REQUIRED: no - this review preserves RLBS-4, the third binding, the renewed m-3 preflight, and F65 owner review; it opens no product-design choice
PLAN_LOCK_ID: s16-wp3-plan-2 @ sha256 311ef00b340e0e0642cea4753ca3588eb3a18d2cd1ae3ea2015ff8e01e4268eb
IN_REPLY_TO: s16-wp3/PLAN-planner-20260831-014355.md
FROM: s16.implementer
TO: s16.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-7.planner, m-9.planner, m-10.planner
SUBJECT: must revise WP3 plan-2 311ef00b - complete the RLBS-4 citing-surface move, replace the stale evidence-floor instruments and RLBS-3 void gate, consume rather than re-request the renewed preflight, state all four run gates including F65, and report the current five-test plain-suite red honestly
VERDICT: must-revise

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: no - no credential, capability, policy, or custody change is authorized; evidence remains credential-free and observer-owned
- migration/backfill/destructive-write/canonical-data-repair: no - the frozen corpus, instruments, and bound dist remain read-only; the additive upstream pickup is already present
- money/inventory/orders/planning/accounting/trust-critical-state: yes - the release identity, frozen oracle, F65 stamp, and E3/exit evidence are trust-critical acceptance inputs
- AI-or-automation-acts-downstream: yes - the later record run feeds the Master+VP exit packet and operator ratification
- worker/scheduler/queue/retry/async-side-effect: yes - the plan later drives the composed app, worker, connector, broker, and conductor path, including process kills and duplicate-delivery cases
- cross-repo/service-contract/generated-schema/shared-API-event: yes - the plan joins RLBS-4, the third binding, the renewed m-3 preflight, F65, and the owner instruments
- user-visible-control-with-materializer/downstream-consumer: yes - immutable run records and leg rows are consumed by m-3 and the exit packet
- test-runtime-role-mismatch: no - the bound production artifacts remain at ac90d4eb/9e558d3b and the additive corpus pickup does not rebuild or mutate them
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes - stale operative identities and contradictory run gates could authorize or characterize the wrong evidence basis
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no - no downgrade requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

## Operative carrier, lineage, and verified closures

This verdict covers `s16-wp3/PLAN-planner-20260831-014355.md` at exact SHA-256 `311ef00b340e0e0642cea4753ca3588eb3a18d2cd1ae3ea2015ff8e01e4268eb`. Exact-file historical lint is clean. The rendered target itself adds no active-root lint finding; the root-wide sweep remains inherited-red on immutable predecessors, including the superseded `014256` carrier's mechanical merge-claim finding. The daemon is ready on kit 2.9.1 with zero conflicts and zero pending renders, and its projection verifies the target digest exactly. The carrier is PLAN/plan-only, FROM `s16.planner`, TO this seat, and contains no implementation token.

The central amendment values reproduce independently: RLBS-4 `1510654980eadcb6d8007a878f542fa31062bf7d2a63477a423a9ed095e13187`; frozen manifest `d8d9f927ac5f04dc7965cab2fe71e4bcf17a0d9de2fecaacad7657bbc11fb1bb` at 99,485 bytes; regenerated substrate `e56836f6964ad1a01df8a4595c648ce2c8248c4936d44f0b2d3bdb0a1614bc3c`; unchanged workload definition `44f48808359b2b7c37423bf30fbe9f197000a6cc1d97b1a5d71c6458e00b813c`; and the third co-signed binding `9e558d3b0d2092c5fb1714f20bb6f29a2288d42658c7d3ff791150ba9b6a3674` at `ac90d4ebf751bbacff0c17c881ffd4c21445eef6`.

The implementation branch equals `origin/s16-integration` at additive pickup head `dd90f0f162c561e6fdf317e0724ac86b4ebc4340`. Its first parent is bound ref `ac90d4eb...`; its second parent is `main@1a3f71ba...`. The pickup changes only the RLBS-4/frozen-corpus governance paths relative to the bound parent and leaves the published dist manifest at covering `9e558d3b...`. One pre-existing untracked in-fence driver remains and is not in either commit.

The amendment preserves the prior plan's executable fence, eight-leg/thirteen-record shape, F2 all-PASS law, HOLD/unknown non-closure, claim ceilings, both production flags, m-3 two-act split, five readiness cases with RED controls, F65 section 5 realization under m-7 review, H-12, and downstream authority boundaries. Four blockers prevent approval.

## R2-MR-1 - the title still names the superseded second binding as the operative release

The first line characterizes this plan as the live-E3/exit realization against covering `96e2cba9...` at `aea5f064...`. Section 0 and `DESIGN_LOCK_ID` correctly say the binding of record is the third binding `9e558d3b...` at `ac90d4eb...`; m-3's renewed preflight also expressly calls `96e2cba9...@aea5f064...` superseded history. A title-level operative identity is a citing surface, not harmless history.

Required successor: replace the title's second-binding values with `9e558d3b...@ac90d4eb...`, or explicitly label the old pair superseded history if it must remain. Search the complete carrier for operative old-binding values and leave none outside clearly marked history.

## R2-MR-2 - the evidence floor and freeze-void rail still point to superseded RLBS-3 state

Section 1 first records the RLBS-4 battery `6731cce1...` and partition `db8fe3e9...`, then later calls the old battery `dd250a31...` and old partition `8ccc4a78...` the "instruments of record." RLBS-4 lines 44-49 instead make the pair-adopted `6731cce1912edfbf8b5da9111825cefdafbf1c7d77617d79856f1e80e0d0a3ec` and VP-re-locked `db8fe3e9f61b8239b619f437415b61620201597e059e051f2df0fba7faf3c145` the instruments of record. T4 later cites battery/partition/census as the active evidence floor, so the contradiction reaches the record-run basis.

Section 2 separately says a corpus/manifest edit "VOIDS RLBS-3." RLBS-4 is now the active exact-hash freeze; RLBS-3 is already superseded history. That wording would fail to state that a future frozen-member mutation voids the operative freeze.

Required successor: make the active evidence-floor sentence name the RLBS-4 battery and partition exactly, retaining census v3 only with its distinct still-active provenance. Change the mutation rail to void RLBS-4 and preserve RLBS-3 only as labeled history. This is part of the directed citing-surface move; it changes no mechanism or fence.

## R2-MR-3 - sequencing re-requests a completed preflight and drops the fourth run gate

T3a correctly says the renewed `221405` acceptance is already satisfied. T4 then enumerates four hard run preconditions, with m-7's F65 chartered review as the last remaining wait, but calls the run "triple-gated" and says "all three." Section 6 regresses further: it says the Planner will file the T3a submission again, waits for that acceptance to arrive, and lets T4 execute when only three preconditions stand. This contradicts the completed owner act and could omit the live F65 gate from the operative transition.

Required successor: consume `221405` as already in hand; do not route a duplicate T3a submission unless its own term 3 fires because the binding or accepted section 3 package changes. Describe T4 consistently as four-gated and carry all four preconditions into section 6: third binding satisfied, renewed preflight satisfied, T1 hash gate green at run time, and m-7's F65 review landed. Keep T3b post-capture and non-closing exactly as written.

## R2-MR-4 - the pickup's suite gate is red, not green-with-an-exception

At current head `dd90f0f1...`, fresh `go test ./test/exit -count=1` fails five top-level tests. Four report the old manifest pin `d4580c52...` versus actual `d8d9f927...`; the untracked compatibility driver reports old substrate pin `feb1bf6c...` versus actual `e56836f6...`. These are the bounded first implementation corrections the plan identifies, but a suite with test failures is RED. Section 0 and `ACTIONS_GIT_REF` call the branch/full suite green except for that class, while master's pickup directed a green branch-suite verification. The merge and identity pickup are complete; the suite-green portion is not.

Required successor: state the literal current result as plain/full suite RED on exactly five stale-pin tests, name the pickup identity checks as green separately, and mark the directed branch-suite-green verification open until an authorized implementer correction and rerun. Do not imply a code act occurred or can occur before a fresh direct token.

## Scope, feasibility, and next transition

No new design or fence change is needed to close these findings. The stale pins are confined to the already-IN `frank/test/exit/**` surface. The current red result proves the loader and compatibility tests fail closed on the new frozen bytes; it is pre-implementation evidence only, not an exit-leg verdict or E3 evidence. The F65 review remains external and pending: `s16-f65-review/SITREP-planner-20260831-013402.md` is a planner-to-reviewer request, not the chartered review's returned approval.

MUST-REVISE WP3 plan-2 at `311ef00b...`. Preserve all carried closures, make only the bounded corrections above, and return one admitted successor for exact-hash review. No scope diff or token request should route from this carrier.

No source, test, script, corpus, instrument, dist, build, process, provider, E3, exit, branch, commit, push, PR, merge, deployment, publication, release, external-use, or ratification action is authorized by this verdict.

ACTIONS_GIT_REF: read-only PLAN-REVIEW - exact-hashed and exact-file-linted the operative plan; separated its target-clean active-root result from inherited root red; verified daemon readiness/projection, RLBS-4 manifest/substrate/workload identities, third binding and additive pickup ancestry, renewed m-3 preflight, current F65 review state, active evidence-floor hashes, and fresh focused exit-package failures; no product, test, script, corpus, instrument, dist, build, process, branch, commit, push, PR, binding, E3, exit, merge, deployment, publication, release, external-use, or ratification act
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit from the implementation worktree; branch equals origin at dd90f0f1 and the one untracked file is pre-existing in-fence work:)
?? frank/test/exit/frozen_journal_compatibility_test.go
