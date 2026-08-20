## PLAN-REVIEW — MUST-REVISE r3: owner rulings and rev21 bytes are folded correctly, but governance truth, the routed-lint oracle, substrate ownership, and operator-only negative coverage need one bounded successor

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s12-build-plan-review-3
PARENT_DISPATCH_ID: s12-build-plan-3
RUN_ID: s12
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — s12.planner can file the bounded successor; the separate r1 structural-error waiver remains an open operator decision and still blocks delegated dispatch
FILED_AT_LOCAL: 20260817-205921
PLAN_LOCK_ID: s12-h16-fix-plan
DESIGN_LOCK_ID: h16-outcome-split-design
PLAN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: s12-build/PLAN-planner-20260817-204945.md
ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — the missing negative covers operator-only authority on canonical attempt-resolution records
- migration/backfill/destructive-write/canonical-data-repair: yes — the locked plan changes canonical chain, repair, and registry compatibility machinery
- money/inventory/orders/planning/accounting/trust-critical-state: yes — relay, credential, and binding truth are trust-critical state
- AI-or-automation-acts-downstream: yes — retry/recovery machinery performs automated canonical and binding actions
- worker/scheduler/queue/retry/async-side-effect: yes — Class-G/Class-D retries, startup drains, and post-commit effects
- cross-repo/service-contract/generated-schema/shared-API-event: yes — fieldspec, Outcome, and auth refusal surfaces are shared contracts
- user-visible-control-with-materializer/downstream-consumer: yes — ceremony and Outcome consumers have materialized downstream effects
- test-runtime-role-mismatch: yes — the plan asserts operator-only authority but has no non-operator raw-ingress execution leg
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — the routed-lint waiver oracle and branch-versus-shared substrate ownership are contradictory as written
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
SUBJECT: r3 must revise on four bounded accuracy and coverage defects; the ruled record-kind realization and rev21 rejection byte pass review

## Verdict

`PLAN_REVIEW_VERDICT: must-revise`. The r3 fold correctly transcribes m-2 `201520` section 1, binds `mint-predecessor-mismatch` at exact rev21 `cc8bcff3f4d04d45eb7cc4250714022f4050ab95bddae55f12fe17bdd7003d05`, includes the v9 reader/transition/genesis consumers, and preserves the four commissioned legs. Four bounded defects remain before an approving review can parent delegated dispatch.

## Blocking findings

### F1 — "Nothing else is outstanding" is false; distinguish task unblocks from the still-open rev21 governance tail

The target says before dispatch only this review and the operator's `194914` waiver remain, and that "nothing else is outstanding" (`PLAN-planner-20260817-204945.md:1,56`). Master `202541` explicitly pre-sequences a fresh rev21 tail after m-7 pair approval: m-1 then m-2 exact-hash delta reruns, then a refreshed Master+VP join (`RECONCILE-orchestrator-planner-20260817-202541.md:26-34`). The live trail confirms that tail is not complete: m-1's first rerun received a bounded must-revise at `205246`; its `205628` successor is filed, while the m-2 rerun and refreshed join are not yet present.

This does **not** revoke s12's task unblocks: master `202541` unblocks Tasks 4/7–15, and m-1's review expressly says its filing correction does not withdraw those unblocks or Task 10's byte. Required successor: state both truths. The pair's dispatch gates are review + the operator's exact r1-error disposition; independently, the rev21 owner/join tail remains open under master's sequence and must be observed at lane exit unless master changes its routing. Remove the absolute "nothing else outstanding" claim.

### F2 — The routed-lint success predicate cannot produce the claimed clean command result

Plan boundary proof says "routed relay-lint clean over `frank/.relays/s12`" (`PL-...md:25`); Task 16 demands "target-scoped clean for every live relay" (`:276`); exit criterion 6 repeats target-scoped routed lint clean (`:292`). Fresh E2 execution disproves a zero-error/zero-exit reading:

- exact-file lint of r3: OK;
- INDEX lint: OK;
- target-scoped routed invocation naming r3: exit 1, the two immutable r1 errors, then `OK` for r3;
- full-root routed invocation: exit 1 with exactly the same two r1 errors.

The pending operator waiver can authorize consumption despite those exact structural errors; it cannot make the linter return clean. Required successor: pin the executable oracle honestly — exact-file lint OK for each consumed/live successor; routed output contains **only** the two explicitly granted r1 errors and an `OK` line for the target, with no other errors; the nonzero exit is accepted only by the exact operator waiver. If the operator denies or changes the disposition, dispatch holds. Replace every unqualified "routed clean" claim with this expected-output rule.

### F3 — The branch/shared-substrate instructions contradict each other

ROADMAP says captures and relay substrate live in the shared `main` tree and the work branch carries code only (`ROADMAP.md:14-16`); Task 16 agrees that no substrate copy is committed (`PL-...md:276`). But the file map still says battery captures are "committed on the branch" (`PL-...md:65`), and Task 0 says the shared tree is "never ... mutated by this build" immediately before directing captures and relays to be written there (`:89`). Task 0 also labels its files "none modified" while creating captures (`:87`). These are mutually exclusive execution instructions.

Required successor: make one byte-consistent model everywhere: code edits/tests/commits occur in the isolated worktree; shared `main` is never checked out or source-mutated by s12, but the shared relay root is append-written for relays/captures and banked by master's checkpoint cadence; no relay/capture copy is committed on `s12-h16-fix`. Correct the Task 0 file declaration and the stale `internal/store/genesis.go` file-map comment (`:47`) to reflect Task 4's functional v9 downgrade edit.

### F4 — Operator-only `attempt_resolution` has implementation prose but no negative authority battery

Rev21 section 5 requires both operator records to be operator-only and commit-time authority-validated together with instance state (`h16-outcome-split.md:135-138`). The m-2 ruling scopes `attempt_resolution` to the operator and the plan says authority is validated (`PL-...md:75,209`), but Task 9's RED set tests state/instance failures only (`:206-208`). A defect accepting a worker/planner raw submit of either the reopen shape or the disposition shape could pass every named Task 9 test.

Required successor: add a real raw-ingress negative for both resolved-target shapes from at least one non-operator seat (prefer a worker and one role-spoof-resistant case): committed REJECTED through the existing ruled/registry authority path, exact existing violation byte asserted, supplied evidence preserved as applicable, and zero fold/state effect. Also assert the operator control accepts the same otherwise-valid shapes. This adds no new class or mechanism; if the live existing class cannot express the required refusal, stop and route that locked-byte gap rather than inventing one.

## Passed review dimensions

- Plan artifact SHA-256 is exactly `aebf19203f330a8aa22d7cec23e6d5b2221b6a1907051049af0dfe37b148a421`; rev21 is exactly `cc8bcff3f4d04d45eb7cc4250714022f4050ab95bddae55f12fe17bdd7003d05`.
- m-2's four enum members, two operator scope entries, no `derived-work-*` seat entries, and by-target `attempt_resolution` carrier are transcribed faithfully.
- `type: string`, the v9 reader/transition/downgrade consumers, exact provenance shape, `genesis`, ceremony freedom, and `-bless`/`operatorSubmit` exclusions are sound at plan grain.
- The task decomposition, 48-case matrix, 31-to-34 registry-driven regression, four lane legs, frozen-oracle/INV fences, E2 target, and terminal operator-only merge gate remain coherent.

BOUNDARY_CONTRACT: not applicable to this review-only artifact — it writes no runtime object; F4 requires the target PLAN's existing boundary contract to gain the missing authority proof
ACTIONS_GIT_REF: docs-workspace disk action — this PLAN-REVIEW relay plus one append-only frank/.relays/s12/INDEX.md row; no source, test, plan, roadmap, branch, commit, PR, merge, deploy, or master-trail edit by this seat
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s12/INDEX.md
 M frank/docs/sprints/active/2026-08-17-s12-h16-fix/ROADMAP.md
 M frank/docs/sprints/active/2026-08-17-s12-h16-fix/plans/PL-s12-build-plan-20260817.md
 M master/domains/m-2-forms-determinism/README.md
 M master/relays/CHECKPOINTS.md
 M master/relays/INDEX.md
?? frank/.relays/s12/s12-build/PLAN-REVIEW-implementer-20260817-205921.md
?? frank/.relays/s12/s12-build/PLAN-planner-20260817-204945.md
?? master/relays/step3-h16-h26-lane/DESIGN-REVIEW-implementer-m1-20260817-205246.md
?? master/relays/step3-h16-h26-lane/DESIGN-planner-m1-20260817-204807.md
?? master/relays/step3-h16-h26-lane/DESIGN-planner-m1-20260817-205628.md

## Verification

- Exact r3 relay lint: OK with `--no-freshness`; pair INDEX lint: OK.
- Target-scoped and full-root routed lint: both exit 1 solely on r1's two immutable errors; target-scoped output separately prints `OK` for r3.
- Direct hashes and live source/registry inspection performed; no runtime/source mutation.
- Live master INDEX re-read through m-1's `205628` successor; no owner ruling or s12 task unblock has been withdrawn.

Next requested action: s12.planner files r4 under `s12-build-plan-4`, folding F1–F4 without changing the ruled design or scope, and returns it for `s12-build-plan-review-4`. Delegated dispatch remains blocked by this verdict and independently by the unresolved operator disposition on `194914`.
