## PLAN-REVIEW — APPROVE r4: all four bounded r3 findings are folded without changing the ruled design, build scope, or terminal gates

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s12-build-plan-review-4
PARENT_DISPATCH_ID: s12-build-plan-4
RUN_ID: s12
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this pair review is complete; delegated dispatch remains independently blocked on the operator's exact `194914` disposition
FILED_AT_LOCAL: 20260817-212251
PLAN_LOCK_ID: s12-h16-fix-plan
DESIGN_LOCK_ID: h16-outcome-split-design
PLAN_REVIEW_VERDICT: approve
IN_REPLY_TO: s12-build/PLAN-planner-20260817-211202.md
ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — the r4 battery now executes both operator-only `attempt_resolution` authority boundaries at raw ingress
- migration/backfill/destructive-write/canonical-data-repair: yes — the locked plan changes canonical chain, repair, version, and registry-compatibility machinery
- money/inventory/orders/planning/accounting/trust-critical-state: yes — relay, credential, marker, and binding truth are trust-critical state
- AI-or-automation-acts-downstream: yes — retry/recovery machinery performs automated canonical and binding actions
- worker/scheduler/queue/retry/async-side-effect: yes — Class-G/Class-D retries, startup drains, and post-commit effects
- cross-repo/service-contract/generated-schema/shared-API-event: yes — fieldspec, Outcome, and auth-refusal surfaces are shared contracts
- user-visible-control-with-materializer/downstream-consumer: yes — ceremony and Outcome consumers have materialized downstream effects
- test-runtime-role-mismatch: no — r4 adds worker and role-spoof-resistant raw-ingress negatives for both operator-only shapes plus operator controls
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: no — r4 pins the inherited routed-lint output, makes substrate ownership consistent, and states the independent governance tail accurately
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no ceremony downgrade or risk waiver is requested by this review
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none
FROM: s12.implementer
TO: s12.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
SUBJECT: r4 approve — F1 through F4 close at plan grain; operator `194914` disposition remains the delegated-dispatch gate

## Verdict

`PLAN_REVIEW_VERDICT: approve`. The locked r4 plan artifact hashes exactly `f1de9e91e2c7d7c87faa471886d186295bbf7251acae54e98919c97e1ade9696`. Each of the four findings from `s12-build-plan-review-3` is closed at plan grain, while the rev21 contract, four commissioned legs, frozen-byte fences, E2 target, and terminal OPERATOR-ONLY MERGE-GATE remain unchanged.

This is review authority only. It is not `DISPATCH IMPL`, does not grant the still-open operator disposition requested by `194914`, and does not claim the independently sequenced rev21 owner reruns or refreshed Master+VP join are complete.

## Finding closure

### F1 — closed: the plan now states both governance truths

Exit criterion 6a distinguishes the pair's own dispatch gates (this approval plus the operator's `194914` disposition) from the independently sequenced rev21 owner/join tail under master `202541`. It says the latter does not gate the pair's build tasks, must be observed at lane exit, and must be reported honestly in the merge-decision return. The live trail advanced after r4 was filed: m-2's second-position rerun at `211431` received pair approval at `212218`, so the owner reruns are now complete and the refreshed Master+VP join remains owed. That concurrency does not falsify r4's filing-time statement or change the pair's dispatch gates.

### F2 — closed: lint is an executable expected-output oracle, not a false clean claim

Global constraint 40 requires exact-file `OK`, exactly the two immutable r1 routed errors plus `OK` for the named target, no other routed error, and operator acceptance of the nonzero result through the exact `194914` disposition. Denial or changed disposition holds dispatch and end review. Task 16 and exit criterion 6 cite that oracle. Fresh review execution reproduced precisely those two inherited errors and `OK` for r4; the incoming r4 relay and pair INDEX lint exactly `OK`.

The phrase “operator-granted r1 errors” occurs inside this explicitly conditional oracle. Because the same sentence says the nonzero exit is accepted only under the operator's exact waiver and the relay says that disposition is still open, it cannot be read as a present grant.

### F3 — closed: one substrate-ownership model is used at every named locus

Task 0 creates the isolated code worktree and shared-tree capture files without modifying a source, test, or registry file. Code commits remain worktree-only. The shared `main` source tree is not checked out or source-mutated by s12, while its relay root is append-written for captures and relays and banked by master's checkpoint cadence; no capture or relay copy is committed on `s12-h16-fix`. The file map now correctly identifies Task 4's functional v9 downgrade edit to `internal/store/genesis.go`.

### F4 — closed: both operator-only shapes have raw-ingress authority coverage

Task 9 now requires a non-operator worker and a role-spoof-resistant credential case for both `attempt_resolution` shapes: reopen against a park and disposition against a marker. Each must commit REJECTED with the exact existing authority violation, preserve supplied headers under the rejected-record discipline, and have zero fold/state effect; operator controls must commit ACCEPTED with the expected effects. It binds the existing class only and explicitly stops for owner routing if the live path cannot express the refusal.

Live-source inspection confirms this is implementable without a new mechanism: submit decoding accepts the full record, `seat.Stamp` overwrites claimed FROM/role from authenticated seat metadata before validation, registry validation emits the existing `seat-scope` authority class, and `rejectAtEdge` retains the candidate record and headers while replacing the body with the violation. Thus “supplied evidence preserved per the rejected-record discipline” is consistent with global constraint 37's supplied-header rule; it does not require preservation of the submitted body.

## Passed carried dimensions

- Rev21 remains bound at `cc8bcff3f4d04d45eb7cc4250714022f4050ab95bddae55f12fe17bdd7003d05`; m-2's four enum members, two operator scope entries, no `derived-work-*` seat entries, by-target carrier, and `mint-predecessor-mismatch` realization remain faithful.
- The v9 reader/transition/downgrade consumers, `type: string`, exact provenance, genesis handling, ceremony freedom, and `-bless`/`operatorSubmit` exclusions remain sound.
- The task decomposition, 48-case matrix, 31-to-34 registry regression, four lane legs, frozen exit-oracle/INV fences, continuous T15 rule, branch isolation, and terminal operator-only merge gate remain coherent.

BOUNDARY_CONTRACT: not applicable to this review-only artifact — it writes no runtime object; the approved PLAN carries the raw-ingress authority proof in its E2 boundary contract
ACTIONS_GIT_REF: docs-workspace disk action — this PLAN-REVIEW relay plus one append-only `frank/.relays/s12/INDEX.md` row; no source, test, plan, roadmap, branch, commit, PR, merge, deploy, or master-trail edit by this seat
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s12/INDEX.md
 M frank/docs/sprints/active/2026-08-17-s12-h16-fix/ROADMAP.md
 M frank/docs/sprints/active/2026-08-17-s12-h16-fix/plans/PL-s12-build-plan-20260817.md
 M master/domains/m-2-forms-determinism/README.md
 M master/relays/CHECKPOINTS.md
 M master/relays/INDEX.md
?? frank/.relays/s12/s12-build/PLAN-REVIEW-implementer-20260817-212251.md
?? frank/.relays/s12/s12-build/PLAN-planner-20260817-211202.md
?? master/relays/step3-h16-h26-lane/DESIGN-REVIEW-implementer-m2-20260817-212218.md
?? master/relays/step3-h16-h26-lane/DESIGN-planner-m2-20260817-211431.md

## Verification

- Plan SHA-256: `f1de9e91e2c7d7c87faa471886d186295bbf7251acae54e98919c97e1ade9696`.
- Exact r4 relay lint with `--no-freshness`: OK; pair INDEX lint before filing: OK.
- Target-scoped routed lint: exactly r1's two immutable errors, then `OK` for r4; no new error.
- Direct live-source inspection covered raw payload decoding, channel stamping, existing `seat-scope` validation, and rejected-record preservation behavior.
- Live pair and master INDEX EOFs were re-read through m-2's concurrent `212218` pair approval; the refreshed Master+VP join remains owed, and the concurrency does not change this plan verdict.

Next requested action: s12.planner performs the commissioned SCOPE_DIFF. Any `DISPATCH IMPL` remains blocked until the operator files the exact disposition requested by `194914`; absent a live bare own-line dispatch token, s12.implementer performs no implementation.
