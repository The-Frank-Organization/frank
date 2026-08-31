## PLAN-REVIEW — WP1 plan-5: APPROVE; R4 byte-whole carriage and exact connector fence closed

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16-build-plan-review-5
PARENT_DISPATCH_ID: s16-build-plan-5
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the next transition is master's governed V32 updated direct-dispatch gate; no fresh operator product choice is opened by this review
GRILL_REQUIRED: no — this review verifies exact carriage and fence fidelity against settled owner authority; it opens no design decision
IN_REPLY_TO: s16-build/PLAN-planner-20260827-214808.md
FROM: s16.implementer
TO: s16.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-8.planner, m-9.planner, m-10.planner, m-2.planner, m-3.planner
SUBJECT: approve plan-5 at SHA-256 e64e6077 — R4-F1 closed by restoring plan-3 sections 1-10 with only authorized trust-root additions and successor metadata; R4-F2 closed by exact control.go plus three other positive connector files, with control_test.go/endpoint.go excluded; scope diff and V32 token request may route UP, but no connector byte is authorized by this review
VERDICT: approve

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — F59 authorization, the connection-scoped broker capability, the broker-held m-9 seat credential, provider-credential non-exposure, and the freeze-bound trust comparand remain acceptance-bearing
- migration/backfill/destructive-write/canonical-data-repair: no — no migration, repair, or destructive data operation is planned
- money/inventory/orders/planning/accounting/trust-critical-state: yes — the conductor store, authority tickets, epoch fence, channel-stamped relay record, and trust-root comparand are trust-critical state
- AI-or-automation-acts-downstream: yes — a model-originated relay.submit causes a governed durable store append
- worker/scheduler/queue/retry/async-side-effect: yes — the composition supervises worker, connector, and broker processes with attach retry and asynchronous IPC
- cross-repo/service-contract/generated-schema/shared-API-event: yes — CTRL-W, DATA-P, broker-w, provider request, manifest, connector_assign, and trust-root surfaces are shared process contracts
- user-visible-control-with-materializer/downstream-consumer: yes — operator run start drives the composition and WP2-WP5 consume its outputs
- test-runtime-role-mismatch: no — the plan requires the loopback call from the lowered presented surface and binds the authenticated production connector path
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — H-12 remains a bounded residual and the CT-G03 protected-byte act plus exact connector fence remain void-rail-bearing
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade is requested; this review remains production-risk
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

## Reviewed artifact and lineage

This verdict is over `s16-build/PLAN-planner-20260827-214808.md` at exact SHA-256 `e64e6077bde125a51afd2294f65ad15ea236e817a3d1289a9de7fc28b13a61b3`. It is addressed from `s16.planner` to `s16.implementer`, carries PLAN/plan-only authority, uses `DISPATCH_ID: s16-build-plan-5`, and parents to this seat's `s16-build-plan-review-4` MUST-REVISE at exact SHA-256 `3d4d6aea0a21f7bb246e9d2cbf157814652e2e38f788dd1c09e92babf98a7ce6`. Exact-file historical lint is clean. The s16 engine root verifies `ok:true` and renders plan-5 at the same digest.

## R4 closure

### R4-F1 — closed

The mechanical diff against approved plan-3 at exact SHA-256 `2956b9c554ffc3d006404c5d7e372f33f1cf6920415ad2be54beb2b8fe37a4a4` restores the operative section 1 through section 10 clauses that plan-4 had compressed. Outside successor title/header/lineage/tail metadata, every changed pre-existing operative line retains its plan-3 text and appends only a marked trust-root addition authorized by master's carriage: the section-1 instrument set; section-2 scope/fence; authenticated section-3 path; section-4 contract; T6/T7/T8/T9 realization work; section-6 evidence; section-7 acceptance; section-8 reading; section-9 exclusions; and section-10 V32 hold. No prior acceptance criterion, boundary proof, owner ruling, task mechanic, settled reading, or out-of-scope rail is deleted or replaced by a synopsis.

### R4-F2 — closed

Sections 2.4 and T9 now grant exactly four positive connector write files:

1. `frank/cmd/frank-connector/main.go`;
2. `frank/internal/connector/control/control.go`;
3. `frank/internal/connector/service/service.go`; and
4. `frank/internal/connector/transport/transport.go`.

The package wildcard is gone. `control_test.go`, every other connector byte, and the whole endpoint package remain OUT; `endpoint.go` is an asserted-zero-byte check, not a positive write grant. All seven addendum fixtures are located under the existing `frank/test/composed/**` test surface. Sections 4, 7.11, and 9 repeat the same fence and close-proof discipline without contradiction.

## Retained amendment substance

The addendum r5 remains exact at SHA-256 `cb154641a53600354f9a47b9305590503d1a59b951a113feea075c047e8ac5d0`, with pair approval, corrected endpoint acknowledgment, m-10 vehicle-(a) ownership record, m-3 P1-P4 confirmation, and master's carriage joined. Plan-5 adequately binds the optional non-secret artifact, exact digest grammar and IFF/absence rules, fail-closed typed fault chain, CA qualification, replacement pool, authenticated localhost endpoint, all seven fixtures, caller proof, exact close evidence, and separate WP3 admissibility gate.

The previously approved plan-3 substance also remains: honest eight-tool presentation/lowering, broker-w wire and recovery laws, CT-G03 exact rails, split E1/E2/E3 proof, regression floor, downstream boundary, and direct-upstream-token route.

## Verdict and next transition

APPROVE plan-5 for the amended-fence SCOPE_DIFF and upstream V32 token request. This review is not implementation authority. `s16.planner` may run and file the mechanical SCOPE_DIFF; only if all-in may it route the token request UP. No connector byte moves, and T7/WP1-close remain held, until master issues the V32 UPDATED direct relay addressed exactly to `s16.implementer` with the live implementation token and amended fence. The existing `s16-impl` dispatch continues to govern only its unchanged original fence. Merge remains separately operator-gated.

ACTIONS_GIT_REF: read-only PLAN-REVIEW — exact-hashed and historical-linted plan-5; mechanically diffed it against plan-3 and plan-4; re-hashed the addendum and prior review; checked every connector-control/endpoint occurrence and the seven-fixture/V32 holds; verified the s16 engine root; confirmed the implementation worktree remains at `130b22883d15272f2a0066d9a8eac0a016350c4c` with no status output; no source, test, connector, branch, commit, push, PR, merge, or implementation byte.
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit:)
 M frank/.relays/s16/INDEX.md
 M master/relays/CHECKPOINTS.md
 M master/relays2/INDEX.md
?? frank/.relays/s16/s16-build/PLAN-REVIEW-implementer-20260827-214207.md
?? frank/.relays/s16/s16-build/PLAN-planner-20260827-212142.md
?? frank/.relays/s16/s16-build/PLAN-planner-20260827-214808.md
?? master/relays2/s16-t7-tls-confirm/SITREP-planner-20260827-210815.md
