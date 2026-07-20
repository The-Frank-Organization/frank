## SITREP — m-9.implementer / Model Runtime adversarial AUDIT waiting on planner matrix

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-boot-m-9-implementer-wait
PARENT_DISPATCH_ID: master-boot-m-9-implementer
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: m-9.implementer
TO: m-9.planner, master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-8.implementer, m-5.planner, m-7.planner, m-3.planner
BUNDLE_ID: m-9-model-runtime
SUBJECT: independent m-9 adversarial source audit complete; final four-bucket AUDIT is correctly withheld because m-9.planner's promotion matrix is not yet present in the relay trail

Phase:
Report-only wait SITREP inside the authorized DESIGN-only AUDIT boot. This is not the final adversarial AUDIT, does not advance phase, and carries no design authorship, design-lock, PLAN, implementation, external-provider-call, tool-execution, credential, merge, or deployment authority.

Current artifact:
- Addressed boot: `master/relays/boot/master-boot-m-9-implementer/SITREP-orchestrator-planner-20260714-223530.md`.
- Required upstream planner boot exists at `master/relays/boot/master-boot-m-9-planner/SITREP-orchestrator-planner-20260714-223520.md`.
- Required upstream output does not: repeated `find`/`rg` checks found no m-9 planner AUDIT or promotion-matrix relay, and `master/domains/m-9-model-runtime/` still contains only its charter.
- Independent source pass completed against pi's turn/tool loop and context contracts, OpenCode's session processor/tool registry/task topology, `frank/` at exact `s11-close@502e06c`, and the locked m-3/m-5/m-7 contracts.

Independent adversarial checkpoints already established (E1; provisional until compared with the planner matrix):

1. **Promote the lifecycle grammar, not either donor loop.** Pi's outer follow-up loop and inner tool/steering loop make turn continuation and stop points explicit (`references/pi/packages/agent/src/agent-loop.ts:155-274`); OpenCode likewise has an explicit session step loop with compaction and terminal detection (`references/opencode/packages/opencode/src/session/prompt.ts:1081-1180`). The useful promotion is a frank-owned logical turn/session state machine. Its state transitions must run on m-7's host and committed truth; m-9 must not create a second scheduler, serialized writer, recovery path, or session database.

2. **Reject donor-owned transcript mutation as canonical state.** Pi mutates an in-memory message array with partial and final assistant messages while emitting raw stream events (`references/pi/packages/agent/src/agent-loop.ts:281-373`); OpenCode writes reasoning/tool/step parts directly during stream processing (`references/opencode/packages/opencode/src/session/processor.ts:215-335,421-469`). Frank may adapt the partial/final lifecycle as an m-8-normalized-event consumer, but any canonical session/context transition must be expressed through the m-7-owned commit substrate, not an m-9-owned store or direct mutable side truth.

3. **The streamed-model-output observation seam is still open.** The locked rule is no clean observation, no outgoing record (`master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:29-37`), and kickoff extends that gate to model output. Landed `observe.Gate` accepts a cloned `record.Record` candidate and computes relay-field stamps/terminal disposition (`frank/internal/observe/gate.go:21-50,90-149`); no streamed-model-event observation API exists. Pi emits text/thinking/tool-call deltas directly (`references/pi/packages/agent/src/agent-loop.ts:319-345`), while OpenCode persists stream deltas before any frank observation (`references/opencode/packages/opencode/src/session/processor.ts:278-305`). Therefore m-9 cannot claim it can merely wire to the existing API: master must route the missing model-output observation contract to m-3 ownership/consumer review, including partial-stream granularity and the rule that unobserved bytes reach no recipient or actuator.

4. **Parsed tool calls must become inert request objects, never donor executors.** Pi looks up a tool, validates arguments, optionally consults `beforeToolCall`, then directly invokes `tool.execute` (`references/pi/packages/agent/src/agent-loop.ts:602-709`; optional-hook contract at `references/pi/packages/agent/src/types.ts:52-63,261-267`). OpenCode adapts native calls straight into `item.execute` (`references/opencode/packages/opencode/src/session/llm/native-runtime.ts:169-190`), and its processor warns the SDK may execute tools before some lifecycle events (`references/opencode/packages/opencode/src/session/processor.ts:98-102`). Those are negative fixtures. M-9 may parse/normalize and request, but the request must remain inert until an owner-defined m-5/m-7 authority/tool-exposure decision has succeeded.

5. **The model-tool authorization interface is not already landed.** Frank's authenticated channel constructs a trusted-side `ToolSet`, rejects unknown tool names, and directly calls the selected function (`frank/internal/channel/server.go:277-336,391-415`). The production surface is byte-exactly `submit/project/read` (`frank/cmd/frank/main.go:396-438,543-548`; invariant `frank/test/invariants/terminal_surface_test.go:106-179`). M-5 locks a spawn-time MAX ceiling whose absent tool axis defaults to `none` (`master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:82-95`), but its own lock says uniform standalone tool blocking was a later enforcement tier (`...archetype-system-design.md:158-164`). Thus “existing trusted authority/tool-exposure path” is true for the three seat verbs, but overstates the landed capability for arbitrary model-tool requests. This is a potential spec/interface mismatch: m-5/m-7 must own or amend the authorization seam; m-9 must not silently fill it.

6. **Reject m-9-owned parallel/sequential execution policy.** Pi chooses sequential versus parallel execution and starts allowed calls itself (`references/pi/packages/agent/src/agent-loop.ts:413-555`; `references/pi/packages/agent/src/types.ts:250-267`). The scheduling concept is useful fixture material, but ownership of whether, where, and concurrently a tool may run belongs after the trusted authority decision and on m-7's host. No m-9 promotion should include `executeToolCalls*`, direct goroutine/task spawning, or a second mutation queue.

7. **Reject OpenCode's in-process child-session topology and permission semantics.** Its task tool creates a child session and recursively prompts it in-process (`references/opencode/packages/opencode/src/tool/task.ts:121-199`), including background execution/result injection (`.../task.ts:202-228`). More sharply, the child rules explicitly say parent agent restrictions govern only the parent while the child's own permissions determine its capabilities (`references/opencode/packages/opencode/src/agent/subagent-permissions.ts:4-26`). Frank's m-5 ceiling is monotonic and may only tighten below the spawn-time MAX (`master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:82-95`); m-7 owns pane/seat spawn and records the ceiling (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:159-162`). In-process child sessions are reject/defer, not an m-9 session feature; native governed spawn remains the kickoff's named Step-3 carry.

8. **Adapt context/compaction as deterministic logical state, reject arbitrary hooks and silent lane changes.** Pi allows arbitrary `transformContext`, per-turn model replacement, dynamic API-key lookup, and payload/response hooks (`references/pi/packages/agent/src/types.ts:140-222`; `references/pi/packages/agent/src/agent.ts:432-466`). OpenCode similarly mutates messages through a plugin hook before the request (`references/opencode/packages/opencode/src/session/prompt.ts:1252-1286`). M-9 can own context assembly/compaction policy, but not unrecorded transforms, credential resolution, provider selection, or model swaps. The input must bind the exact handed lane/snapshot; context changes must be deterministic/replayable and committed on m-7's substrate.

9. **Adapt cancellation/retry/partial-stream semantics with an explicit no-replay boundary.** Pi stops on typed error/abort and refuses truncated tool calls (`references/pi/packages/agent/src/agent-loop.ts:196-214,376-407`). OpenCode wraps the provider stream in retry while its own comment acknowledges tool execution can occur inside that stream (`references/opencode/packages/opencode/src/session/processor.ts:98-102,627-676`). A frank design must therefore distinguish provider-request retry from authorized tool execution: no stream retry may replay a tool effect, duplicate a committed tool result, or surface an unobserved partial. This belongs in the m-8/m-9 contract and m-7-hosted idempotency boundary, not donor control flow.

10. **Defer steering, soft interrupt, and native spawn beyond V1.** Pi's steering queue is polled only after the current assistant turn's tool calls finish (`references/pi/packages/agent/src/types.ts:224-248`), a behavior useful as a counterexample for true boundary steering. Kickoff explicitly keeps native governed spawn, steer-at-boundary, and soft-interrupt/cancel-redeliver as non-terminal Step-3 carries (`master/STEP-3-KICKOFF.md:64-67`). They must not be smuggled into the minimal V1 turn loop or treated as already solved by donor queues.

Duplicate/already-built gate:
- `frank/internal` and `frank/cmd` contain no provider adapter, normalized model-event contract, model-turn/session runtime, or inert model-tool-request authorization surface. The landed overlap is boundary substrate only: authenticated per-seat tool exposure, the m-7 serialized loop, record-shaped observe-as-send, and the locked authority-ceiling data/semantics.
- The substrate is already closed and must be consumed, not rebuilt: m-7's loop serializes jobs and commits through the canonical store (`frank/internal/engine/loop.go:84-109,130-168`), while the submit path stamps identity, validates form/lineage, runs observation, then returns the candidate for the loop's commit (`frank/internal/engine/submit.go:51-130`).
- Result: m-9 is still open, with product overlap at the m-3/m-5/m-7 seams. The two absent runtime-facing seams above must be routed to their owners rather than absorbed by m-9.

Done:
- Routing and phase authority verified.
- `using-superpowers` and the full Agent Pair Implementer skill/protocol applied.
- Locked kickoff, m-9 charter, Cycle Playbook Part B, and locked m-3/m-5/m-7 anchors read.
- Independent adversarial pass completed with byte evidence at exact `frank` head `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.
- Incoming boot target returned `OK` from the requested exact-file lint invocation; root-context diagnostics were pre-existing INDEX/lineage noise.

Not done:
- No final four-bucket AUDIT verdict issued.
- No row-by-row adversarial comparison against m-9.planner's promotion matrix performed, because that artifact is absent.
- No design authored or reviewed; no design lock, plan, code, tool execution, credential access, provider call, external send, branch, commit, PR, merge, or deployment action performed.

Blocked:
- The comparative half of the boot assignment requires m-9.planner's actual promotion matrix. Inventing its rows would violate the file-first and evidence-over-seniority rules.

Scope drift risk:
- High if a final verdict is issued before the planner bytes exist: the final review could miss a direct executor, m-7 loop/store re-ownership, in-process child session, silent lane swap, or an output-observation bypass. Low while this SITREP holds the lane.

Boundary contract status:
- Writes: final m-9 adversarial AUDIT relay only after the planner matrix arrives.
- Reads: planner promotion matrix plus the audited pi/OpenCode/frank/m-x source sets.
- Target entity: master reconciliation input for the m-9 pre-design AUDIT.
- Downstream consumer: m-9.planner and master/VP AUDIT reconciliation.
- Contract: every promoted/adapted/rejected row must preserve m-7 hosting/commit ownership, keep parsed tools inert until m-5/m-7 authorization, route model output through m-3 observation, and reject in-process spawn.
- Proof: E1 file/line comparison in the eventual AUDIT.
- No-consumer action: defer the final verdict; do not speculate.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — model-tool authorization and the credential boundary are adjacent; this SITREP makes no change
- migration/backfill/destructive-write/canonical-data-repair: no — read-only source audit
- money/inventory/orders/planning/accounting/trust-critical-state: yes — authority and trusted-courier state are load-bearing; this SITREP makes no change
- AI-or-automation-acts-downstream: yes — the audited runtime would execute model-requested tools only after trusted authorization
- worker/scheduler/queue/retry/async-side-effect: yes — turn loop, retry, cancellation, and tool execution are the subject; this SITREP makes no change
- cross-repo/service-contract/generated-schema/shared-API-event: yes — m-8/m-9/m-3/m-5/m-7 interfaces are cross-domain contracts
- user-visible-control-with-materializer/downstream-consumer: no — read-only audit
- test-runtime-role-mismatch: unknown — no implementation or test runtime exists yet
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — partial-output observation granularity and tool-authorization ownership require owner resolution
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no ceremony downgrade or risk acceptance is requested by this read-only wait SITREP
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

Tests / verification:
- E1: exact boot relay and all required local anchors inspected.
- E1: reference and frank byte anchors above inspected at exact `frank` head `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b` (`s11-close`).
- E2: `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/boot/master-boot-m-9-implementer/SITREP-orchestrator-planner-20260714-223530.md` returned target-file `OK` (root reported pre-existing INDEX/lineage noise).

ACTIONS_GIT_REF: wrote only `master/relays/boot/master-boot-m-9-implementer/SITREP-implementer-20260714-212251.md` and appended its routing row to `master/relays/INDEX.md`; no `frank/` edit; `frank` remained at `502e06c` with clean `git status --short` before the relay write.
FINAL_GIT_STATUS_SHORT: unavailable — `/Users/jack/Programming/harness` is not a git repository; final `frank/` status is reported separately and must remain clean

Next requested action:
- m-9.planner authors and routes the lint-clean AUDIT/promotion matrix to `m-9.implementer`; the operator then re-presents that exact path. I will reopen the current bytes, compare every promote/adapt/reject row against the independent checkpoints above, and return the final four-bucket adversarial AUDIT to m-9.planner + master.
