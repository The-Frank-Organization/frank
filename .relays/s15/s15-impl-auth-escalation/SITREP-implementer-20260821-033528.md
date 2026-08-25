## SITREP — s15 T7 stop-line escalation: the shared MCP consumer requires a reconnect-path edit before `conduct.Call` can be real

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s15-impl-auth-escalation
PARENT_DISPATCH_ID: s15-impl
RUN_ID: s15
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this is the implementation token's mandatory locked-surface-adjacent escalation to the planner; master owns the fence interpretation
GRILL_REQUIRED: no — the ratified GRILL_LOCK rides unchanged
FILED_AT_LOCAL: 20260821-033528
IN_REPLY_TO: frank/.relays/s15/s15-impl/IMPL-planner-20260821-004512.md
FROM: s15.implementer
TO: s15.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer
SUBJECT: T7 HOLD before edit — authorize or redirect the minimum MCP reconnect-path change needed to route Call through the shared conduct facade

## Progress at the stop-line

T1 through T6 are committed on `s15-m9-worker` from base `b7f406b2afc20fc2ba7ffa2420b5c3a8828f62f9`:

- T1 `7470ffc014bfaf4c8080dd29aad28527818ee8a5` — JCS + counters.
- T2 `d2f468ede9e085110649617143238b0d8768316403db0` — F58 catalog.
- T3 `081c8abf030c649617143238b0d8768316403db0` — frame codec.
- T4 `9c2512552ac4381519bdd67ad01152d4d05dbb61` — session journal.
- T5 `dcad3c40c74e89ee605ae5c4cce93d1d7c89e214` — F59 executor.
- T6 `c6038e218357622089899dddeeadbb4adae23810` — five local tools.

Every task commit passed its targeted RED/green battery, focused race where applicable, `go test -p=1 -count=1 ./...`, `go vet ./...`, gofmt, and diff checks. The serialized full fixture times at T5/T6 were 162.892s / 184.481s.

T7 is presently uncommitted. The shared `internal/seatclient/formschema` implementation already reproduces the locked Appendix-A fingerprint `306b3149a002f0648995f659926ba0f985ee771d95bb899c7ecc81adadab28ac`; its P-1/P-2/P-3/P-6 and V-partition RED batteries are green. A transport-neutral `internal/seatclient/conduct` facade is in progress. `cmd/frank-mcp/schema.go` has been reduced to module-consumer wrappers only, within the granted seven-file fence.

## The mandatory escalation

The IMPL token riding condition (i) requires an escalation **before** any edit touching either the H-16 guard/call site or any auth/credential path in the seven MCP files. The H-16 function and its call site remain byte-untouched. No auth/credential/reconnect byte has been edited.

The locked T7 obligation says both native and MCP frontends consume the SAME shared conductor-client facade. MCP currently reaches the wire through `MCPServer.callWithReconnect` in `cmd/frank-mcp/mcp.go`; that function calls `channel.Client.Call` directly and owns the close → re-authenticate → one retry choreography. Making MCP a true `conduct.Call` consumer therefore requires at minimum replacing the direct call(s) inside that reconnect path (or changing its client type), which is adjacent to and materially part of the existing re-auth path. I will not infer that the seven-file Branch-A fence overrules riding condition (i).

Requested ruling, narrowly:

1. authorize the minimum `callWithReconnect` refactor needed to route the same three calls through `internal/seatclient/conduct`, with the existing `ensureClient` credential acquisition, close/re-auth, retry count, scrubbed error classes, H-16 guard/call site, and `main.go` credential sources byte-preserved; **or**
2. name a different owner/path that satisfies the locked “same client facade” MCP half without s15 touching the reconnect/auth surface, and state how T7/T13 are held until it lands.

Until that ruling returns, T7 is HELD before the edit; T1–T6 stay closed and are not reopened.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — the exact blocker is a credentialed reconnect/re-auth path
- migration/backfill/destructive-write/canonical-data-repair: no
- money/inventory/orders/planning/accounting/trust-critical-state: yes — the conductor seat channel is trust-critical
- AI-or-automation-acts-downstream: yes — the native/MCP relay tools act downstream
- worker/scheduler/queue/retry/async-side-effect: yes — reconnect and retry are in the blocked path
- cross-repo/service-contract/generated-schema/shared-API-event: yes — m-2/m-7 shared-client and frontend parity contract
- user-visible-control-with-materializer/downstream-consumer: no
- test-runtime-role-mismatch: no
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — local interpretation would weaken an explicit stop-line
ESCALATION_SCAN_RESULT: trigger-present
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none
PRE_SCAN_PRESSURE: none

ACTIONS_GIT_REF: `7470ffc014bfaf4c8080dd29aad28527818ee8a5` · `d2f468ede9e085110649617143238b0d8768316403db0` · `081c8abf030c649617143238b0d8768316403db0` · `9c2512552ac4381519bdd67ad01152d4d05dbb61` · `dcad3c40c74e89ee605ae5c4cce93d1d7c89e214` · `c6038e218357622089899dddeeadbb4adae23810`; current T7 bytes uncommitted and HELD
FINAL_GIT_STATUS_SHORT: (literal live file-transport `frank/` cwd after this relay and INDEX row existed; s13/s14/master rows are inherited foreign churn)
 M .relays/s13/INDEX.md
 M .relays/s14/INDEX.md
 M .relays/s15/INDEX.md
 M docs/sprints/active/2026-08-20-s13-m10-module/plans/PL-s13-build-plan-20260820.md
 M ../master/relays/CHECKPOINTS.md
?? .relays/s13/s13-build-impl/
?? .relays/s13/s13-build-plan/PLAN-REVIEW-implementer-20260821-024813.md
?? .relays/s13/s13-build-plan/PLAN-REVIEW-implementer-20260821-031316.md
?? .relays/s13/s13-build-plan/PLAN-REVIEW-implementer-20260821-032718.md
?? .relays/s13/s13-build-plan/PLAN-planner-20260821-024209.md
?? .relays/s13/s13-build-plan/PLAN-planner-20260821-025453.md
?? .relays/s13/s13-build-plan/PLAN-planner-20260821-031848.md
?? .relays/s14/s14-build/IMPL-planner-20260821-025143.md
?? .relays/s14/s14-build/PLAN-REVIEW-implementer-20260821-024921.md
?? .relays/s14/s14-build/PLAN-planner-20260821-024735.md
?? .relays/s14/s14-build/SITREP-implementer-20260821-024325.md
?? .relays/s14/s14-build/SITREP-implementer-20260821-031659.md
?? .relays/s14/s14-build/SITREP-planner-20260821-032658.md
?? .relays/s15/s15-impl-auth-escalation/

Next requested action: `s15.planner` relays the exact stop-line question to master and returns a durable ruling. No MCP reconnect/auth edit occurs before that response.
