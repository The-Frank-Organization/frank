## PLAN-REVIEW - WP3 plan-1: MUST-REVISE; the frozen-suite/hash/fence direction carries, but the formal E3 ruling is sequenced before the captures it must judge and the bound release cannot enter the current production app under the zero-production-byte fence

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16-wp3-plan-review-1
PARENT_DISPATCH_ID: s16-wp3-plan
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - bounded owner/master corrections and a successor pair review are required; the later operator MERGE-GATE and Step-3 ratification remain untouched
GRILL_REQUIRED: no - this review preserves the frozen oracle and owner locks and makes no product-design choice for them
PLAN_LOCK_ID: s16-wp3-plan @ sha256 b976244eae5be96babee99f90dc3aa421a91d99d7be978dabff23ab8e607b737
IN_REPLY_TO: s16-wp3/PLAN-planner-20260828-144354.md
FROM: s16.implementer
TO: s16.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-7.planner, m-9.planner, m-10.planner
SUBJECT: must revise WP3 plan-1 b976244e - split m-3 pre-run shape approval from post-capture formal admissibility, carry all five readiness preconditions, and resolve the unavailable production release carrier plus four-member F63 versus closed E3-vector collision before any implementation token
VERDICT: must-revise

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: no - no credential or authority change is authorized; evidence must remain credential-free and observer-owned
- migration/backfill/destructive-write/canonical-data-repair: no - frozen inputs and dist artifacts remain read-only; fault injection is run-scoped
- money/inventory/orders/planning/accounting/trust-critical-state: yes - E3 applicability and the release/exit verdicts are trust-critical acceptance evidence
- AI-or-automation-acts-downstream: yes - these records feed the Master+VP exit packet and operator ratification
- worker/scheduler/queue/retry/async-side-effect: yes - the proposed suite drives process kills, duplicate delivery, and composed worker/app/broker behavior
- cross-repo/service-contract/generated-schema/shared-API-event: yes - the plan consumes closed m-3, m-7, m-8, m-9, m-10, F63/F65, RLBS-3, and frozen-oracle contracts
- user-visible-control-with-materializer/downstream-consumer: yes - the emitted records are consumed by the exit packet
- test-runtime-role-mismatch: yes - the current production app cannot consume the bound release vector, so a test-only manifest path would not prove the claimed bound composed runtime
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes - the release-carrier and broker-in-E3-vector seams are unresolved owner/master decisions, not implementer choices
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no - no downgrade requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

## Reviewed carrier and closures

This verdict covers corrected carrier s16-wp3/PLAN-planner-20260828-144354.md at exact SHA-256 b976244eae5be96babee99f90dc3aa421a91d99d7be978dabff23ab8e607b737. Historical exact-file lint is clean. The daemon ledger records sequence 72 with admits_against_seq=71, so it explicitly corrects immutable lint-red predecessor PLAN-planner-20260828-144331.md. It is PLAN/plan-only, addressed TO this seat, and carries no implementation token.

The governing hashes reproduce: plan r7 044fe4ccefa3b55e16c897afecfe1f6bca9971a2a7c905aac804bd53713c726d; RLBS-3 289a0da01d9ba654f10b01865d54d9174aa862ba5a3eec21439e0ca3ef4d93c3; frozen manifest d4580c52675038049471e2fd4ea813c42604b21b0032a9ba5f39fa794f972639; workload definition 44f48808359b2b7c37423bf30fbe9f197000a6cc1d97b1a5d71c6458e00b813c. The thirteen-id inventory, 32/100 budget, eight-leg indivisibility, F2 all-PASS law, claim ceilings, HOLD/unknown non-closure, frozen-byte prohibition, H-12 posture, platform dist path, zero production/seam/corpus bytes, and fresh-master-dispatch requirement carry.

Four blockers prevent approval.

## R1-MR-1 - T3 asks m-3 for a formal verdict before the evidence exists

T3 submits a record-run design, calls the response the formal E3-admissibility ruling, and T4 makes it a hard precondition of the run that creates captures and E3 records. The m-3 owner records state the opposite temporal fact: the injected-CA posture is a pre-signaled admissible shape, but the formal ruling renders against the live run and its captures, and nothing pre-approves an unexecuted run (s16-t7-tls-confirm/SITREP-planner-20260827-204725.md; s16-t7-tls-inst/SITREP-planner-20260827-154918.md).

Required successor: split two owner acts. Before execution, m-3 accepts the concrete run design, recorder boundary, capture plan, and applicability inputs as preflight shape. After the bound run produces immutable captures and candidate m3.e3_observation.v2 records, m-3 evaluates those exact bytes/context and returns the formal ruling. The run starts only after RELEASE-BINDING, owner preflight, and runtime hash gate; no E3/exit closure proceeds before the post-capture formal ruling.

## R1-MR-2 - the five m-3 readiness preconditions are not executable acceptance criteria

The readiness return s16-pm-readiness/SITREP-planner-20260826-203227.md says the plan must bake in five items. The plan does not require all three observer-reachable capture points (pre-attach frozen request core; worker effect descriptor plus captured invocation/env; committed stamped relay read at the observer boundary), an observer-owned output root the composed seats cannot write, fresh F84 re-derivation, incomplete capture mapping to unknown/HOLD with machines 2/5 non-gating, and a first-class identity-manifest input with mutation invalidation.

Required successor: make those five named acceptance cases, including RED controls for seat-write denial, copied-digest refusal, incomplete-capture unknown/HOLD, and each bound-identity mutation. Outside the conductor store plus a separate process does not establish the non-seat-writable recorder boundary.

## R1-MR-3 - no production path consumes the pending bound release under this fence

T2/T4 require bound dist artifacts on the composed production path, while section 2 forbids every production byte. At banked head 4e179fc3, frank-app still hashes the literal frank-mvp-development and writes a working-tree placeholder ReleaseBinding (frank/cmd/frank-app/main.go:141-150). It has no CLI/config input for a completed binding. A test helper can manufacture a truthful manifest in memory, but that is not the claimed production release-carrier path; the plan's test-runtime-role-mismatch answer becomes true.

Required successor: wait for Q-CARRIER and point to an existing production consumption seam, or route owner-approved production realization as a separately fenced predecessor. If production bytes must move, amend the fence and re-review. A test-only injection cannot substitute.

## R1-MR-4 - the four-member F63 set does not fit the closed E3 v2 vector

T3 says the applicability vector carries all four F63 digests: app, worker, connector, broker. The closed m3.e3_observation.v2 matrix and production manifest.BuildDigests admit only app_main_build_digest, m9_worker_build_digest, and m8_build_digest, XOR one covering release_digest. broker_build_digest would be malformed. The canonical release artifact has four member digests but no ruled covering digest, and the owner proposal called covering form later-release. The pair cannot invent encoding or leave broker outside applicability.

Required successor: route the collision through master to m-3/m-10 with RELEASE-BINDING/Q-CARRIER. Return with one owner-authorized closed form that mechanically binds every ratified F63 member and is consumed by the production app and m-3 evaluator. A schema or production-contract change is a predecessor act, not exit-test discretion.

## Verdict and next transition

MUST-REVISE. Preserve all closures. Resolve R1-MR-3/R1-MR-4 before any implementation token; fold R1-MR-1/R1-MR-2 into a successor plan and return it for exact-hash review. No source, test, script, corpus, dist, build, process, provider, branch, commit, push, PR, binding, E3, exit, merge, or release action is authorized.

ACTIONS_GIT_REF: read-only PLAN-REVIEW - exact-hashed and historical-linted corrected plan; verified daemon correction lineage, parent, plan r7, RLBS-3, frozen manifest/workload hashes, m-3 readiness and ruling records, composed helpers, production release construction, closed E3 and manifest vectors, dist manifest, and clean/equal local-upstream implementation head; no product, test, script, corpus, dist, branch, commit, push, PR, binding, E3, merge, or release act
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit; implementation worktree clean/equal origin:)
 M frank/.relays/s16/INDEX.md
?? frank/.relays/s16/s16-wp2b/SITREP-implementer-20260828-144120.md
?? frank/.relays/s16/s16-wp3/PLAN-planner-20260828-144331.md
?? frank/.relays/s16/s16-wp3/PLAN-planner-20260828-144354.md
