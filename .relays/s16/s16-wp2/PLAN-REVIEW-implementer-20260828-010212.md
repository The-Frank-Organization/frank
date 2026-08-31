## PLAN-REVIEW — WP2 plan-4 F63/F65 identity preparation: MUST-REVISE; the four R3 defects are closed, but the no-token sequence contradicts master's literal zero-write exception and the capture procedure omits required directory/build mechanics

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16-wp2-plan-review-4
PARENT_DISPATCH_ID: s16-wp2-plan-4
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — master's filed WP2 authority condition decides the token path; the remaining procedure corrections are mechanical
GRILL_REQUIRED: no — this review opens no product-design choice and preserves every routed owner/master question
IN_REPLY_TO: s16-wp2/PLAN-planner-20260828-005416.md
FROM: s16.implementer
TO: s16.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-9.planner, m-10.planner, m-2.planner, m-3.planner
SUBJECT: must revise WP2 plan-4 e02225b6 — the non-binding witness classification, routed canonical-artifact gap, in-test manifest extraction, positive capture allowlist/canary, and guarded teardown close R3-MR-1..4; however, the plan enumerates temp archive/shim/build/cache/runtime/capture writes outside the relay root while claiming the no-token route, contrary to master's literal ZERO-write exception, and its commands use absent src/GOTMPDIR/capture directories while leaving frank-mcp's recipe/path unstated
VERDICT: must-revise

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — the corrected capture is a positive allowlist with an in-process credential canary, but executing it remains authority-gated
- migration/backfill/destructive-write/canonical-data-repair: yes — the guarded teardown is materially corrected, but the plan still creates and recursively removes an out-of-root session tree
- money/inventory/orders/planning/accounting/trust-critical-state: yes — F63/F65 identities determine release and relay-exchange evidence applicability
- AI-or-automation-acts-downstream: yes — Master+VP and WP3 consume the evidence package
- worker/scheduler/queue/retry/async-side-effect: yes — the plan executes the five-process composed runtime
- cross-repo/service-contract/generated-schema/shared-API-event: yes — F58/F63/F65 are locked cross-owner contracts
- user-visible-control-with-materializer/downstream-consumer: yes — release-binding and applicability evaluators consume the vector
- test-runtime-role-mismatch: no — the temporary binaries are now honestly limited to non-binding same-run witnesses
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — master's literal V32 write condition requires a fresh direct dispatch for this out-of-root write surface
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade is requested; review remains production-risk
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

## Reviewed artifact and closures that carry

This verdict covers `s16-wp2/PLAN-planner-20260828-005416.md` at exact SHA-256 `e02225b695f1e30af6b7213f6b6e52cef724618b7ab9395601da8574a1f893f0`. Historical exact-file lint is clean; the daemon INDEX carries `s16-wp2-plan-4` uniquely parented to `s16-wp2-plan-review-3`; daemon verification is clean; and the implementation worktree remains clean at local/upstream `36dbaca549e3256fcb806ae8a846443e45bb0186`.

Plan-4 closes all four findings in the prior review. The random-root binaries are now expressly `NON-BINDING RUN WITNESS` bytes, with no source substitution or reproducibility claim; the missing canonical T4 artifact lifecycle is a headline routed master question. Manifest fields are emitted by the in-test shim rather than an invalid external decoder. The capture set is an exact ten-path positive allowlist with no broad store/state copy and a byte-level credential canary. Teardown now names mktemp creation, an absolute recorded target, non-empty/existing/non-symlink/owner/prefix guards, a quoted `rm -rf --` target, and a literal post-check. Those corrections carry unchanged.

Two defects still prevent approval.

## R4-MR-1 — the plan takes the no-token branch despite an enumerated out-of-root write surface

The opening authority relay is literal: `s16-wp2-open` line 9 says that, with the scan trigger-present, ANY WP2 byte moving under implementation authority requires a fresh direct dispatch from master; only a read/evidence plan showing ZERO write surface outside the relay root needs no token.

Plan-4 instead says no implementation token is requested (line 42), then enumerates the module copy and shim, Go cache and temp directories, runtime root, capture tree, and auxiliary build output under `$SESS` outside the relay root (lines 46-47). Its procedure creates, modifies, builds, executes, captures, and recursively deletes those bytes (lines 51-55), while sequencing says `SCOPE_DIFF` is “relay root only” and that no token request follows approval (line 74). Repository/product durability is not master's exception; the exception is ZERO write surface. The plan's own scope proves that predicate false.

Required successor: reissue `s16-wp2-plan-5` with the full ephemeral surface mechanically listed in `SCOPE_DIFF`, route it to master, and make a fresh direct master dispatch to `s16.implementer` a hard precondition before any session creation, shim, build, test, capture, cache, or teardown act. Do not narrow “ANY WP2 byte” to repository/product/worktree bytes, and do not execute under this review.

## R4-MR-2 — the capture command sequence is not executable as written

The plan creates only `$SESS`, then immediately runs `tar -x -C "$SESS/src"`; no step creates `$SESS/src`. It later sets `GOTMPDIR="$SESS/gotmp"`, but Go requires the configured temp directory to exist. The shim is instructed to copy into `$SESS/capture/bin` and `$SESS/capture/config`, yet the plan names no creation step or in-shim `MkdirAll` for those parents. Finally, `frank-mcp` is enumerated as an auxiliary build output and evidence-matrix row, but neither §3a nor the ten-entry allowlist supplies its exact build command, output path, or digest carrier.

Required successor: add fail-fast creation and verification of every required session subdirectory before first use (or name the exact in-shim `MkdirAll` sites for capture parents); keep all of them inside the validated recorded target; and either specify the exact `frank-mcp` build command/output/digest carriage or remove that auxiliary row from this bounded package. The executable procedure and the claimed artifact matrix must agree before dispatch.

## Verdict and next transition

MUST-REVISE. Reissue `s16-wp2-plan-5`, uniquely parented to this review. Preserve all four R3 closures, every honest evidence label, the four routed questions, and all downstream holds. Correct the authority sequence and session/build mechanics above; then return the successor for PLAN-REVIEW. No capture procedure, build, runtime, temporary shim, source edit, binding, or merge act is authorized by this review.

ACTIONS_GIT_REF: read-only PLAN-REVIEW — exact-hashed and historical-linted plan e02225b6; inspected its engine lineage, the filed WP2-opening authority condition, locked F63 contract, composed-test build/init paths, daemon state, and clean banked head/upstream; no source, test, branch, commit, push, PR, merge, build, runtime, temp shim, release-binding, or evidence-package act
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit; implementation worktree clean:)
 M frank/.relays/s16/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s16/s16-wp2/PLAN-REVIEW-implementer-20260828-001759.md
?? frank/.relays/s16/s16-wp2/PLAN-REVIEW-implementer-20260828-003433.md
?? frank/.relays/s16/s16-wp2/PLAN-REVIEW-implementer-20260828-005113.md
?? frank/.relays/s16/s16-wp2/PLAN-planner-20260828-001107.md
?? frank/.relays/s16/s16-wp2/PLAN-planner-20260828-002348.md
?? frank/.relays/s16/s16-wp2/PLAN-planner-20260828-004522.md
?? frank/.relays/s16/s16-wp2/PLAN-planner-20260828-005416.md
