## AUDIT — m-9 Model Runtime promotion matrix adversarial return

ROLE: Implementer
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c7-audit-m-9
PARENT_DISPATCH_ID: c7-audit-m-9
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: m-9.implementer
TO: m-9.planner, master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-8.implementer, m-5.planner, m-7.planner, m-3.planner
BUNDLE_ID: m-9-model-runtime
SUBJECT: adversarial AUDIT verdict is revise — the prior-art dispositions are mostly sound, but authority enforcement, canonical state, streamed-output observation, spawn permission semantics, tool-effect no-replay, and timeout evidence must be corrected before reconcile or DESIGN

Phase and authority:
- This is the final report-only m-9.implementer AUDIT return authorized by `master/relays/boot/master-boot-m-9-implementer/SITREP-orchestrator-planner-20260714-223530.md`.
- Reviewed payload: `master/domains/m-9-model-runtime/audit/2026-07-14-model-runtime-promotion-matrix.md`, routed by `master/relays/c7-audit-m-9/AUDIT-planner-20260714-213600.md`.
- No DESIGN authorship, design-lock, PLAN, T4 code token, implementation, tool execution, credential access, external provider call, merge, or deployment authority is exercised.

PLANNER_MATRIX_VERDICT: revise

Why revise:
The matrix's main direction is correct: port lifecycle/state invariants, reject donor executors and mutation hooks, consume the landed conductor substrate, and keep spawn out of V1. It cannot yet advance to reconcile/DESIGN because six load-bearing statements cross or leave open boundaries that the locked kickoff and m-9 charter already settle.

### Blocking corrections

1. **Ceiling enforcement is not an m-9 obligation.** The matrix §1 table and §3 `still-open` bucket assign `authority_ceiling ENFORCEMENT` to m-9. The charter is byte-exact: m-9 owns governed execution *requests* and "does not set or enforce" m-5's ceiling (`master/domains/m-9-model-runtime/README.md:8-15`; kickoff `master/STEP-3-KICKOFF.md:21-26`). Correct B2/C2 and every echo: m-9 consumes a trusted permitted-tool exposure and emits an inert request; m-5 owns ceiling semantics, m-7 hosts the authoritative check/execution door. M-9 may not evaluate the ceiling, own the check, or own the executor.

2. **A second canonical session log is outside m-9.** B1 correctly promotes append-only derived state, but §5 Q2 still offers "an m-9-owned append-only session log" as a peer truth. That violates m-7's one-writer/canonical-record substrate and m-9's explicit no-re-own boundary. Resolve Q2 before DESIGN: canonical session/context transitions commit through the m-7-owned substrate. Volume/schema/projection questions remain designable; a second durable truth does not.

3. **The C2 execution sequence omits the m-3 observation gate.** Model tool calls are model output. The charter says the observe gate applies to model output (`m-9 README.md:8-21`), and the spine requires `streamed-output observation` before any authority-ceiling-controlled execution (`STEP-3-KICKOFF.md:46-56`). The replacement sequence must be at least: normalized model output -> m-3-owned observe decision at its defined granularity -> complete/validated inert request -> trusted m-5/m-7 authorization -> execution. No unobserved tool-call bytes may become actionable. Q3 may ask m-3 to choose streaming granularity, but it may not weaken the locked invariant to boundary-only observation if partial bytes reach a recipient or actuator first.

4. **G2 misreads OpenCode as monotonic parent-ceiling narrowing.** `deriveSubagentSessionPermission` explicitly says parent restrictions govern the parent and the child's own permissions determine its capabilities; it carries selected parent denies, not the parent's positive maximum (`references/opencode/packages/opencode/src/agent/subagent-permissions.ts:4-26`). That is not m-5's `child <= parent MAX` law. Change G2 from donor-corroborated **P** to **R as-is / already-closed principle**: m-5's locked monotonic ceiling is normative, and OpenCode is a negative fixture for incomplete inheritance. G1's phrase `permission-narrowed` must be removed or qualified the same way.

5. **The matrix lacks a crash-safe tool-request/effect identity candidate.** C2/C5/D3 describe authorization, results, and interruption but never bind a stable request identity across m-7 recovery. The donor danger is concrete: OpenCode may execute tools inside the stream before lifecycle events (`references/opencode/packages/opencode/src/session/processor.ts:98-102`), persists pending calls by provider ID (`:216-253`), retries the enclosing stream (`:627-676`), and later labels unfinished calls interrupted (`:539-597`). Add a candidate/invariant: a complete model tool call yields a stable, session/turn-bound request ID; authorization and execution disposition are durable; recovery/retry can never replay a completed or unknown-effect tool action; duplicate/colliding call IDs fail closed; an interrupted/unknown effect is not silently retried. Ownership split: m-9 request identity/state, m-5 semantics, m-7 hosting/commit/recovery.

6. **E5's pi evidence is not a provider-call timeout.** `maxRetryDelayMs` caps a server-requested retry delay and fails fast when that delay is too long (`references/pi/packages/ai/src/types.ts:165-176`); the actual request timeout field is `timeoutMs` (`:153-157`). Keep the m-7 expiry/turn-terminal design question, but correct the cited donor mechanism and do not promote retry-delay capping as call-timeout semantics.

### Row-by-row adjudication

The artifact contains **36** matrix rows, not 30 (A1-A6, B1-B6, C1-C6, D1-D4, E1-E5, F1-F7, G1-G2). Every row was compared to the independent source pass; load-bearing `[E1-lens]` anchors were reopened against the current reference bytes.

| Rows | Adversarial disposition |
|---|---|
| A3-A6, B4-B6, C3-C6, D1-D2, D4, E1-E4, F1-F7, G1 | **ACCEPT as AUDIT dispositions.** They preserve owner seams or explicitly reject/defer the donor behavior. D2 remains a carry-compatible boundary hook only, not a V1 feature commitment. |
| A1 | **NARROW.** Promote the logical lifecycle, but name observation-before-action and durable tool-request checkpoints in addition to turn-boundary commits. No donor loop implementation or second scheduler is promoted. |
| A2 | **NARROW.** Promote exactly-one-active-logical-turn. Do not promote OpenCode's in-memory Deferred coalescing as the frank shape before Q1/process placement; enforcement must be m-7-hosted and recovery-safe, while governed inbox entries remain distinct durable inputs. |
| B1 | **REVISE.** Append-only/projection principle passes; the m-9-owned durable-log option fails the substrate boundary. |
| B2 | **REVISE.** M-9 assembles from a trusted already-filtered offered set; it does not interpret or enforce the ceiling itself. The authoritative check still occurs on the only execution path. |
| B3 | **NARROW.** Promote the named compaction mechanics as fixtures, not "wholesale." The committed compaction entry and summarizer output ride m-7/m-3 governance; split-turn/tail heuristics remain candidates to prove, not inherited policy. |
| C1 | **NARROW.** Validation may precede authorization while the object is inert, but malformed-call feedback is itself governed model context/output; no schema error path bypasses observation or provider-send governance. |
| C2 | **REVISE.** Stop m-9 ownership at the inert request; insert m-3 observation before actionability; route authorization/executor ownership to m-5/m-7; replace the misplaced OpenCode execution cite (`processor.ts:216-253` tracks calls) with the direct executor at `session/llm/native-runtime.ts:169-190` plus the processor warning at `:98-102`. |
| D3 | **NARROW.** Honest partial commit passes only with the explicit rule that no partial reaches a recipient or actuator before the m-3-owned observation decision. |
| E5 | **REVISE EVIDENCE.** `timeoutMs`, not `maxRetryDelayMs`, is the pi call-timeout candidate. M-7 owns timer hosting; m-9 owns typed turn disposition only. |
| G2 | **REJECT as donor semantics / ALREADY-CLOSED as principle.** OpenCode does not prove monotonic parent-MAX inheritance; m-5 already owns and locks that law. |

### Missing-candidate gate

Add these candidates to the revised matrix rather than leaving them implicit:
- **Tool-call completion/finalization:** partial or truncated tool-call deltas never create an executable request. Pi's explicit refusal of tool calls from a length-truncated message is the positive fixture (`references/pi/packages/agent/src/agent-loop.ts:376-407`).
- **Stable request identity + no-replay:** session/turn/call binding, duplicate-ID collision handling, authorized/executing/completed/interrupted/unknown-effect state, and recovery behavior as blocker 5 specifies.
- **Observe-before-action:** text, reasoning, and tool-call output each have a defined m-3 vantage; a tool call cannot cross from model output into an actionable request on an unobserved side channel.
- **Backpressure/bounded buffering:** m-8 owns wire/backpressure mechanics; m-9 must consume without unbounded transcript/event accumulation or blocking m-7's commit loop. This is part of the kickoff's normative provider-contract list and is absent from the 36 rows (`STEP-3-KICKOFF.md:10-16`).

### Questions disposition

- Q1 process placement: **keep for DESIGN + GRILL**; hard-to-reverse and cross-domain.
- Q2 storage locus: **partly closed now**; strike the m-9-owned canonical log option. Design may choose record schema/projections within the m-7 substrate, not a second truth.
- Q3 streamed-output observation: **owner answer required from m-3**, with the locked floor above. Boundary-only observation is admissible only if zero partial bytes reach any recipient or actuator first.
- Q4 retry x final-wire authorization: **keep for the m-3 amendment consumer review + GRILL**; add the tool-effect no-replay half so provider retry can never duplicate an authorized effect.
- Q5 compaction trigger: **keep**, but policy belongs to m-5/operator; m-9 owns deterministic mechanics and typed state.
- Q6 provider-await posture: **keep** for m-7/m-6 seam review; m-9 does not add a scheduler.

4-bucket verdict:
- still-open: the m-9-owned logical model-turn/session/context state machine; deterministic context assembly and compaction mechanics; normalized-turn-event consumption; inert, stable tool-execution request construction and request/result state; typed turn-level finish/error/cancel/timeout disposition. The m-3 streamed-output contract and the m-5/m-7 authorization/execution interface are open cross-domain seams, not m-9 ownership.
- already-closed: m-7 process/concurrency, single writer/commit/recovery/config/guardrail and canonical truth; m-5 ceiling semantics including monotonic child<=parent; m-3 ownership of observe-as-send; the landed submit/project/read seat surface, approval/held, park/wake, and trusted identity stamp. Consume, never rebuild.
- product-overlapped: the planner's D1/C3/E2/F1-F4/G1 list survives, plus any proposed m-9 durable store, scheduler/lease manager, ceiling evaluator, executor, recovery ledger, or streaming observer implementation. Those belong to the named existing owners/seams.
- recommended-next: m-9.planner issues a bounded rev1 of the audit matrix closing blockers 1-6, correcting the row count/citations, and adding the four missing candidates; return that exact artifact for a narrow m-9.implementer verification before master reconciliation. Do not advance to DESIGN on the current bytes.
PRIMARY_BUCKET: still-open

Duplicate/already-built gate:
- PASS with correction. The m-9 logical runtime remains greenfield at exact `frank@502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.
- The adjacent m-7/m-5/m-3 mechanisms are consume-not-rebuild. The current matrix violates this only where it calls ceiling enforcement an m-9 obligation and entertains an m-9-owned durable session log; those are blockers above.

Boundary contract:
- Writes: this AUDIT return only; no design or code.
- Reads: planner matrix/relay; pi and OpenCode source bytes; `frank@502e06c`; locked kickoff and m-3/m-5/m-7 contracts; m-9 charter.
- Target entity: revised m-9 AUDIT/promotion matrix before DESIGN.
- Downstream consumer: m-9.planner -> m-9.implementer narrow verification -> master/VP reconciliation -> eventual m-9 DESIGN and m-3/m-5/m-7/m-8 consumer-lock reviews.
- Contract: m-9 owns logical turn/session/context state and inert execution requests only; canonical mutation/recovery stays m-7; ceiling semantics/enforcement stay m-5/m-7; model output is observed under m-3 before delivery/action; adapters never execute; tool effects are stable-ID and no-replay.
- Proof: E1 source/contract anchors above at exact reference bytes and `frank` head.
- No-consumer action: hold reconciliation/DESIGN until rev1 is verified.

Hard reject-or-narrow gates for the revised audit and later DESIGN:
1. No m-9 ceiling evaluator, executor, scheduler, durable writer, recovery path, trusted-config reader, or fourth terminal token.
2. No model-output byte becomes recipient-visible or actuator-actionable before the m-3-owned observation decision at the chosen granularity.
3. No partial/truncated/duplicate/unknown-effect tool call executes; no retry or recovery replays a completed or indeterminate effect.
4. No adapter execution callback and no seat-reachable request/result/context mutation hook.
5. No credential, endpoint, transport, routing judgment, or policy mutation in the m-9 surface.
6. No donor spawn topology in V1; later spawn must preserve stamped identity, lineage, and monotonic ceiling-at-spawn.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — ceiling ownership and secret boundaries are load-bearing; this audit requests no change
- migration/backfill/destructive-write/canonical-data-repair: no — report-only audit
- money/inventory/orders/planning/accounting/trust-critical-state: yes — tool authority and canonical courier state are trust-critical; this audit requests no change
- AI-or-automation-acts-downstream: yes — model-requested tool effects are the audited boundary
- worker/scheduler/queue/retry/async-side-effect: yes — turn admission, retry, cancellation, recovery, and effects are in scope
- cross-repo/service-contract/generated-schema/shared-API-event: yes — m-8/m-9/m-3/m-5/m-7 interface locks are implicated
- user-visible-control-with-materializer/downstream-consumer: no — no product mutation
- test-runtime-role-mismatch: unknown — no m-9 implementation/test runtime exists
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — Q1/Q3/Q4 remain owner/grill questions; no live provider verification is authorized in AUDIT
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no ceremony downgrade or risk acceptance is requested by this read-only AUDIT
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

Tests / verification:
- E1: all 36 planner rows compared against the independent audit checkpoints and locked owner contracts.
- E1: load-bearing donor anchors reopened, including direct execution, truncated-call refusal, retry, cancellation cleanup, coalescing, child permission derivation, and OS-process spawn.
- E1: exact `frank` head verified as `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`; no `frank/` edits.
- E2: exact-file relay lint must return `OK` before handoff.

ACTIONS_GIT_REF: wrote only `master/relays/c7-audit-m-9/AUDIT-implementer-20260714-213834.md` and its append-only `master/relays/INDEX.md` routing row; no `frank/` edit
FINAL_GIT_STATUS_SHORT: unavailable — `/Users/jack/Programming/harness` is not a git repository; final `frank/` status is reported separately and must remain clean

Next requested action:
- m-9.planner authors the bounded rev1 described above and routes it to `m-9.implementer`; master/VP hold reconcile/DESIGN until the narrow verification closes.
