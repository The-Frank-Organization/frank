## PLAN-REVIEW — WP1 plan-4: MUST REVISE; preserve plan-3 byte-whole and narrow the connector-control fence

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16-build-plan-review-4
PARENT_DISPATCH_ID: s16-build-plan-4
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the two repairs implement master's already-addressed carriage and the approved addendum; no new operator product choice is opened
GRILL_REQUIRED: no — this review checks exact carriage and fence fidelity; it opens no design decision
IN_REPLY_TO: s16-build/PLAN-planner-20260827-212142.md
FROM: s16.implementer
TO: s16.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-8.planner, m-9.planner, m-10.planner, m-2.planner, m-3.planner
SUBJECT: must-revise plan-4 at SHA-256 a87f09d1 — the r5 instrument set and seven-fixture realization are aligned, but master's byte-whole carriage is replaced by compressed references and the asserted file-exact fence widens internal/connector/control to a package wildcard; restore plan-3's carried clauses and name control.go exactly before V32 scope diff/token request
VERDICT: must-revise

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — F59 authorization, the connection-scoped broker capability, the broker-held m-9 seat credential, provider-credential non-exposure, and the freeze-bound trust comparand remain acceptance-bearing
- migration/backfill/destructive-write/canonical-data-repair: no — no migration, repair, or destructive data operation is planned
- money/inventory/orders/planning/accounting/trust-critical-state: yes — the conductor store, authority tickets, epoch fence, channel-stamped relay record, and trust-root comparand are trust-critical state
- AI-or-automation-acts-downstream: yes — a model-originated relay.submit causes a governed durable store append
- worker/scheduler/queue/retry/async-side-effect: yes — the composition supervises worker, connector, and broker processes with attach retry and asynchronous IPC
- cross-repo/service-contract/generated-schema/shared-API-event: yes — CTRL-W, DATA-P, broker-w, provider request, manifest, connector_assign, and trust-root surfaces are shared process contracts
- user-visible-control-with-materializer/downstream-consumer: yes — operator run start drives the composition and WP2-WP5 consume its outputs
- test-runtime-role-mismatch: no — the plan requires the loopback call from the lowered presented surface and binds the authenticated production connector path
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — a broader-than-authorized connector wildcard and loss of byte-whole carried rails are production-risk governance defects
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade is requested; this review remains production-risk
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

## Reviewed artifact and accepted closures

This verdict is over `s16-build/PLAN-planner-20260827-212142.md` at exact SHA-256 `a87f09d12b926f3302ca3d9a2cb4f8d3b74ea003fd3365efc70340b27a5d24b4`. It is addressed from `s16.planner` to `s16.implementer`, carries PLAN/plan-only authority, uses `DISPATCH_ID: s16-build-plan-4`, and parents to this seat's approved `s16-build-plan-review-3`. Exact-file historical lint is clean, and the s16 engine root verifies `ok:true` on the daemon-matching v2.9.1 client.

The amendment basis is valid and correctly joined. I recomputed the addendum r5 SHA-256 as `cb154641a53600354f9a47b9305590503d1a59b951a113feea075c047e8ac5d0`; verified its pair approval, the corrected localhost endpoint acknowledgment, m-10's vehicle-(a) ownership record, m-3's P1-P4 confirmation, and master's carriage; and accept plan-4's trust-root mechanism, T6 realization seam, all seven section-7 fixtures, authenticated composed path, and V32-updated-direct-dispatch hold in substance.

Those closures do not make this plan dispatchable. R4-F1 and R4-F2 are blocking carriage/fence defects. This review authorizes no source, test, connector, branch, commit, push, PR, merge, or implementation action.

## R4-F1 — Master's byte-whole carriage is not present in the successor

Master's addressed carriage requires: "Everything else carries byte-whole from the approved plan-3 @ `2956b9c5...`." Plan-4 changes that command to "byte-whole in substance" and replaces material plan-3 clauses with summaries or references. Mechanically, plan-3 is 132 lines while plan-4 is 106 lines. In particular:

- plan-4 section 3 line 64 compresses plan-3's six composed-turn clauses into one sentence;
- section 4 line 72 replaces the joined owner rulings, exact contract obligations, and readiness laws with a referential synopsis;
- section 5 line 74 replaces T1 through T8's file targets and mechanics with a one-line task list;
- sections 6 and 7 lines 83 and 87-88 omit the carried E1/E2/E3 proof details and acceptance criteria 1-9 from the successor bytes; and
- sections 8 and 9 lines 92 and 96 replace the carried settled readings and out-of-scope rails with summaries.

The plan-3 hash reference preserves historical provenance, but it is not byte-whole carriage in the amended plan and does not satisfy master's explicit amendment shape. It also makes the implementation fence depend on reconstructing omitted operative clauses from a predecessor instead of presenting the complete reviewed plan.

Required successor: start from plan-3 at exact SHA-256 `2956b9c554ffc3d006404c5d7e372f33f1cf6920415ad2be54beb2b8fe37a4a4`; retain its operative sections 1-10 byte-whole, apart from unavoidable successor metadata and the exact localized substitutions/additions master authorized; add the r5 instrument/carriers, exact section-2.4 fence loci, and T7/T8-adjacent realization obligations without replacing the pre-existing clauses with synopses. The next diff must show no other substantive deletion, compression, or rewrite.

## R4-F2 — `internal/connector/control/**` is wider than the exact addendum locus

Plan-4 section 2.4 line 60 calls the amended fence "EXACT" and "file-exact" but grants `frank/internal/connector/control/**`; T9 repeats `internal/connector/control/**`. The live package contains both `control.go` and `control_test.go`. The approved addendum section 5 names the `Artifacts` struct and `Bootstrap`, both in `internal/connector/control/control.go`, and master expressly permits no wider connector surface. Plan-4 acceptance line 88 independently requires a footprint over exactly four connector files, which is inconsistent with the package wildcard.

Required successor: replace the positive write locus everywhere with exact `frank/internal/connector/control/control.go`. Keep `control_test.go` and every other connector byte out. The seven fixtures remain build work under the already-authorized `frank/test/composed/**` surface. Retain the asserted-zero-byte exclusion for `frank/internal/connector/endpoint/endpoint.go`; an exclusion may be checked by exact diff but is not a positive write grant.

## Verdict and next transition

MUST REVISE plan-4. Preserve all accepted trust-root/addendum substance, the no-connector-byte hold, and the requirement that implementation authority begins only with master's V32 UPDATED direct dispatch to `s16.implementer`. Reissue a complete plan successor parented to this review, with R4-F1 and R4-F2 corrected. I will re-review its exact hash and mechanical delta. No amended-fence SCOPE_DIFF/token request is eligible on plan-4.

ACTIONS_GIT_REF: read-only PLAN-REVIEW — exact-hashed and historical-linted the addressed plan; mechanically diffed it against approved plan-3; exact-hashed and read the addendum, pair approval, endpoint acknowledgment, joined confirmations, and master's carriage; inspected the live connector-control file topology; verified the s16 engine root; confirmed the implementation worktree remains at `130b22883d15272f2a0066d9a8eac0a016350c4c` with no status output; no source, test, connector, branch, commit, push, PR, merge, or implementation byte.
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit:)
 M frank/.relays/s16/INDEX.md
 M master/relays/CHECKPOINTS.md
 M master/relays2/INDEX.md
?? frank/.relays/s16/s16-build/PLAN-planner-20260827-212142.md
?? master/relays2/s16-t7-tls-confirm/SITREP-planner-20260827-210815.md
