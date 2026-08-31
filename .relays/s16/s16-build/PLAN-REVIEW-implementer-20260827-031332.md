## PLAN-REVIEW — WP1 composition plan: MUST REVISE before any code token

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16-build-plan-review-1
PARENT_DISPATCH_ID: s16-build
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the plan must return through master and the named PM owners; no operator choice is required if the heavier workflow and locked owner bytes are followed
IN_REPLY_TO: s16-build/PLAN-planner-20260827-030621.md
FROM: s16.implementer
TO: s16.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-8.planner, m-9.planner, m-10.planner, m-3.planner
SUBJECT: must-revise — the proposed empty-tool scripted reply is not the pinned honest eight-tool surface; the broker-worker framed contract is not identified; production-risk ceremony and E3 evidence are understated; two superseded readiness carriers and a completed re-cert remain stale in the plan
VERDICT: must-revise

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — F59 authorization, connection-scoped broker capability, broker-held m-9 credential, and provider credential non-exposure are acceptance-bearing
- migration/backfill/destructive-write/canonical-data-repair: no — no migration, repair, or destructive data operation is planned
- money/inventory/orders/planning/accounting/trust-critical-state: yes — the conductor store, authority tickets, epoch fence, and channel-stamped relay record are trust-critical state
- AI-or-automation-acts-downstream: yes — a model-originated relay.submit causes a governed durable store append
- worker/scheduler/queue/retry/async-side-effect: yes — the plan composes and supervises worker, connector, and broker processes with attach retry and asynchronous IPC
- cross-repo/service-contract/generated-schema/shared-API-event: yes — CTRL-W, DATA-P, and worker-broker framed IPC plus the provider request and relay operation surfaces are shared process contracts
- user-visible-control-with-materializer/downstream-consumer: yes — operator run start drives the composition and WP2-WP5 consume its outputs
- test-runtime-role-mismatch: yes — the proposed test scripts a tool request that the runtime did not advertise, so the harness would certify a path unavailable to a conforming provider interaction
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — calling the zero-advertisement scripted path honest would accept ambiguous product semantics and residual authority risk
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade is requested; this review selects production-risk
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

## Review result

The plan has a sound high-level target and the SessionLogPath repair in section 8 item 1 is an in-fence app-side composition fix. The plan is not dispatchable as written. Findings F1-F4 below are blocking; F5 is a required carrier correction. No branch, worktree, source byte, test byte, commit, push, PR, or implementation action is authorized by this review.

## F1 — The proposed demo is not an honest governed turn on the pinned m-9 surface

The plan acknowledges that `runtime.marshalProviderRequest` presents `tools: []`, then proposes that a scripted loopback provider return `relay.submit` anyway and calls the demo unaffected (plan section 8 item 2). That is not the locked surface:

- Current source confirms the carrier is empty at `frank/internal/worker/runtime/runtime.go:310-330`.
- The pinned m-9 producer contract requires the presented schema and description arrays to carry exactly the ratified eight canonical names; a missing name causes assembly refusal, no logical-surface digest, no attempt open, and no DATA-P request (`master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md:423-429,478-495`).
- The same contract says the logical-surface digest describes what was actually presented to the model and the eight descriptions come from the presented worker registry (`…m9-delta.md:475-492`).

A scripted peer emitting an unadvertised call can exercise the downstream parser, but it cannot prove the acceptance claim that the model/provider boundary produced one honest governed tool call. Treating landed source as governing over pinned owner bytes reverses the plan's own conflict rule.

Required successor: route this conflict UP through master to m-9 (and m-2/m-8/m-10 where their consumed seams require it). The revised plan must bind an owner-authorized path that presents all eight locked tools, produces/carries `logical_surface_digest`, and lets the controlled provider return a call from that presented set. If that requires changing pinned m-9 runtime semantics, name the exact owner successor/amendment and widen the task wording accordingly. If the locked acceptance is to change instead, only the authorized architecture path may do so. No first code token may issue while this conflict is open.

## F2 — The worker-broker framed protocol is not identified, yet implementation would have to invent it

The runtime exposes only an in-process `Broker` interface (`Attach` and `Rediscover`) at `frank/internal/worker/runtime/runtime.go:75-78`. The current broker app-IPC registry is a CLOSED family containing only the app-control messages `state_proposal`, `state_proposal_result`, `epoch_state`, `boundary_cut`, `epoch_installed`, and `broker_event_ack` (`frank/internal/appipc/msgs_broker.go:60-88`). There are no registered worker-attach, worker-operation, response, push, or correlation frames in the built surface.

The owner contract pins semantic obligations — a private framed worker-broker IPC, attach precedence, a capability delivered only on that connection, the three relay verbs plus typed Describe and push, and per-operation fencing (`master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md:91-128,172-182,214-224`). The plan's T5 says only to wire an endpoint and leaves mechanics to the Implementer. Choosing new message names, bodies, correlation, capability carriage, close/error behavior, and whether to extend the CLOSED `ChannelBroker` family is shared-contract design, not implementation mechanics.

Required successor: cite an already locked exact worker-broker wire contract if one exists, or route a bounded contract fill to m-7 with m-9/m-10 consumer confirmation. Then state the exact production and test targets, request/reply/push shapes, correlation and connection-scope assertions, and scope treatment for any registry or `cmd/frank-broker` bytes. The Implementer must not invent this seam under an implementation token.

## F3 — `CEREMONY_TIER: medium` is an invalid downgrade for this work

The plan directly touches authorization, secrets, trust-critical state, downstream-acting AI, worker supervision/retry, and multiple process contracts. Those are hard escalation triggers. The plan carries neither the required completed `ESCALATION_SCAN` nor a post-scan operator waiver, and section 10 assumes `no hard trigger` despite its own tasks proving several.

Required successor: use `CEREMONY_TIER: production-risk` and include the completed scan, or present the scan to the operator and carry a valid post-scan waiver for any proposed downgrade. Do not self-waive or retain the delegated no-hard-trigger dispatch path while triggers are present.

## F4 — The strongest WP1 proof is E3, not E2

The acceptance claim runs five real binaries, supervises worker/connector/broker processes, operates a live local conductor and Unix sockets, and verifies a committed store record plus recovered runtime journal. Under the installed protocol, a local stack/worker/DB-or-API runtime check is E3; E2 is local test/lint/fixture command proof. The fact that a `go test` harness drives the stack does not lower the runtime claim.

Required successor: set `EVIDENCE_TARGET: E3` and distinguish the E2 static/test/loci evidence from the E3 composed runtime proof. This does not claim WP3's separately governed externally observed live E3 or discharge its admissibility gate.

## F5 — The authoritative input list and interface-lock status are stale

Plan section 1 names the superseded m-7 `20260826-203128` and m-8 `20260826-203226` readiness carriers as normative inputs. The addressed rejoin says those predecessors are immutable history and the clean owner successors are `master/relays2/s16-pm-readiness/SITREP-planner-20260827-023135.md` and `…-20260827-023313.md` (`s16-readiness/RECONCILE-orchestrator-planner-20260827-030435.md:1-7`). Plan section 4 also says the worker-row r12 re-cert is pending, while `master/STEP-3-INTERFACE-LOCK.md:33-42` records RE-CERTIFICATION 3 complete at r12 `63f5c49d…`.

Required successor: replace the two readiness pointers with their admitted clean successors and state that RE-CERTIFICATION 3 is complete. Preserve predecessor lineage as history only, not as the plan's operative ready-to-guide carrier set.

## Section 8 readings

1. SessionLogPath: approve the stated app-side repair. The current producer path and worker basename check disagree; the proposed `<state-dir>/<runID>/session.log` shape resolves the composition mismatch without changing worker journal semantics.
2. Empty provider tool advertisement: reject the stated default. It is acceptance-blocking under F1 and must route UP before any implementation dispatch.

## Successor acceptance

A reissued `PLAN` must close F1-F5, retain the no-source/no-branch hold, and come back as `s16-build-plan-2` parented to this review. The next PLAN-REVIEW will check the owner-authorized tool-presentation path, exact broker-worker wire contract, production-risk scan, E3 labeling, corrected carrier chain, and the resulting exact write fence. An implementation dispatch is ineligible until that review returns `approve` and every delegated condition is actually met.

ACTIONS_GIT_REF: read-only PLAN-REVIEW — inspected the addressed plan, current production call sites and registries, pinned m-7/m-9/m-10 owner bytes, interface-lock re-certification, readiness rejoin, and the installed protocol; no source, test, branch, worktree, commit, push, PR, merge, or implementation byte.
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit:)
 M frank/.relays/s16/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s16/s16-build/
