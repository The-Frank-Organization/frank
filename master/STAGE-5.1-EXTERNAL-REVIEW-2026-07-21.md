# Stage-5.1 external review (third-party chat model + research harness) — 2026-07-21

**Provenance.** Independent review commissioned by the operator before the stage-6 gate, run by a
third-party chat model with a web-research harness, against the vendored governance workspace at
`frank-dev` `main@c78da38` (the full `master/` tree). Pasted into the master session verbatim by the
operator 2026-07-21. Stored here byte-for-byte as the citable basis for `STEP-3-STAGE6-AMENDMENT.md`.
This is an external artifact — its claims are the reviewer's, not master's; master's disposition of
each finding is recorded in the amendment, not here.

---

## Overall verdict

Frank has a real and differentiated project inside it, and the July architecture reframe is the correct
direction — the conductor should remain a small isolated governance plane, not an agent-runtime + provider
gateway + supervisor + database at once; the native harness should sit above it as a replaceable reference
implementation. That boundary is the strongest architectural decision in the current design.

The reviewer does NOT think the current design should be ratified unchanged under the label "MVP of the
Frank harness." It is closer to a sophisticated **governance-plane integration prototype**: it proves the
relay boundary, a designated provider-send path, and durable one-shot dispatch bookkeeping; it deliberately
does NOT contain the dominant local authority (`bash`); it deliberately loses the model's working context
on worker replacement; and its exit test mostly proves protocol composition, not that the resulting coding
agent is useful, recoverable, or acceptably safe.

"As an MVP of Frank's governance kernel, the design is strong. As an MVP coding-agent product, it is
upside-down: the most elaborate parts are control-plane mechanics, while the security boundary, local tool
semantics, resumability, and task-level product proof remain deferred." As of `main@c78da381`, both stage-6
halves are complete but the operator gate is pending and T4 authority is not issued — the last sensible
moment to correct the milestone rather than patch around it.

## What is genuinely strong
1. **The conductor boundary is correct** — demoting the conductor from application hub to one isolated
   service; each component has a coherent job; the conductor stays useful outside the native harness (a
   governance plane between heterogeneous harnesses).
2. **Exceptionally honest about evidence** — reporter/observer/validator/recorder/authorization-point/
   effect-point/failure separation; the H-17 census; explicit `UNKNOWN_TOOL_OUTCOME`/`UNKNOWN_PROVIDER_OUTCOME`,
   cancellation-distinct-from-failure, "no event" when the observer died. Refuses to manufacture plausible
   tool results. More serious about outcomes/indeterminacy than e.g. Microsoft's Agent Governance Toolkit,
   which records attempts rather than verified outcomes.
3. **The one-shot execution binding is a good primitive** — ticket bound to run/turn/epoch/call-id/tool-name/
   args-digest, consumed atomically before invocation; refusal to silently replay consumed-but-unrecorded.
4. **Governance outside the summarizer is sound** — reloading constraints/objective from owner-controlled
   sources, not a lossy summary; directly addresses the Governance Decay result (0%→30–59% violations).
5. **Provider handling substantially better specified than most adapters** — freeze→authorize→attach→send,
   canonical endpoint restrictions, reserved-auth-header exclusion, zero credential resolution on denial,
   normalized terminal outcomes, explicit body hash.
6. **The project already knows most of its weaknesses** — the hardening backlog names TCB fragmentation,
   sandboxing, effect descriptors, sequence-level policy, signed evidence, formal modeling, stale status
   projections, authority provenance; already found H-16 relabeling + H-26 unlocked minting. "The sequencing
   is wrong" — several Step-4 items are prerequisites for calling Step-3 a harness MVP.

## The biggest problems

### 1. The threat model does not close under prompt injection
The design treats an agent acting on hostile content as a *confused* agent (in scope), yet gives `bash`
ambient host authority, no sandbox, no network restriction, no allow/block list, unpinned cwd/env, same-user
access to peer app-side processes, and daemon-escape from process-group teardown — and states arbitrary
network egress from `bash` is outside the governed route. Those positions cannot coexist: a prompt-injected
agent need not become a "malicious autonomous principal"; it need only follow an untrusted instruction
(inspect peer processes, read outside the workspace, `curl` an exfil endpoint, modify another process's
config, start a detached process, bypass the connector). Prompt injection commonly causes exactly such
deliberate-looking actions. Deployed systems increasingly treat filesystem + network isolation as baseline
(OpenAI Codex managed sandbox + no open-ended outbound; Anthropic: both FS and network isolation needed).
Catching `setsid` escapees requiring a full OS sandbox is the reason to pull the sandbox forward, not defer.
**Recommendation:** make the MVP Linux-first and place every local tool behind a sandbox executor (separate
UID/container or namespace, workspace-only FS, network denied by default, no broker/m-8/m-10 sockets, minimal
env, process/output limits). The local-backend abstraction is already the right seam — implement the first
concrete backend as the sandbox, not the ambient host; a macOS backend can honestly offer weaker guarantees
later. Without this, the separate m-8/broker processes are software boundaries, not security boundaries.

### 2. The "exact request" proof has two holes
(a) **Local effects bind model arguments, not the actual execution context.** F59 binds the canonical
model-provided args (`bash` ≈ `{command, timeout}`) while cwd/env stay ambient. It proves "this command
string was authorized+invoked," not "this exact action, in this dir, env, FS root, shell, network policy,
workspace state." Same for file tools (`"path":"foo"` depends on root/symlinks/canonicalization). The ticket
should bind an **effect descriptor** — `{action, canonical_resource, workspace_root_id, cwd, env_profile_digest,
backend_id, network_policy_id, canonical_args_digest, one_shot}`. For local-write, at minimum bind the
canonical confined path, workspace/snapshot identity, and exact content/patch digest.
(b) **The live E3 record does not bind the observed send to `frozen_core_digest`.** m-8 computes a good
frozen-core digest (method/endpoint/headers/body-hash/body-len) but `m3.e3_observation.v1` does not list it as
an attempt-scoped field — so an E3 can be applicable to the correct attempt/release without mechanically
proving the observed request was the exact request m-8 froze. A lock-level gap. Add the frozen digest to the
m-10 provider-attempt record, m-8 terminal, attempt-scoped E0, attempt-scoped E3, and the composite exit
proof; the observer independently derives it. Matches recent Proof-of-Execution work.

### 3. Losing the whole working context on replacement is too weak for a harness MVP
The transcript is memory-only; on replacement the objective is recovered but the replacement starts a fresh
turn with no transcript/summary. Avoids a second authority store (good) but the consequences are severe
(repeats exploration, repeats completed non-obvious actions, loses settled results, more tokens/time,
possibly contradictory strategy; long-horizon work fragile exactly when recovery is needed). Persist a typed,
append-only **model-visible run journal** (attempt input-item hashes/refs, provider-visible output items,
complete tool calls, settled results, compaction events + template/version, objective + hard-constraint refs,
workspace snapshot/revision ids, unknown-effect markers). Checkpoint at settled tool-round boundaries; a crash
inside an attempt/effect still parks UNKNOWN. This can be a projection over canonical events + immutable
content blobs, preserving no-second-outcome-truth. (cf. OpenAI Agents SDK externalized state; LangGraph
step-boundary persistence.)

### 4. The broker protocol is too complex for the value it provides in the MVP
Broker survives app-main failure, retains the seat, suspends, is later adopted; its epoch transition includes
stable op/transition ids, PROPOSED→PREPARING→CROSSERS_DURABLE→INSTALLED, frozen crossing set, durable crossing
rows, recognition-based recovery, survivor-vs-fresh behavior, completion-before-install, aborted attempts,
withheld payloads, rediscovered outcomes. Impressive, but a custom distributed protocol with enough states/
crash-windows to be a project by itself. For the MVP, surviving an app-main crash has limited value (the
broker suspends anyway; push is best-effort; the durable store supports rediscovery; the worker is being
replaced; product value is unproven). Simpler rule: app-main failure kills all children incl the broker;
restart makes a new broker + rebinds the same logical seat; no op crosses a worker epoch; at replacement stop
admissions + short bounded drain; unresolved → typed UNKNOWN/connection-lost; recovery via rediscovery +
content-hash dedup. If broker survival is non-negotiable, make the bounded TLA+/Alloy model (H-24) a **pre-T4
gate**. H-16 relabeling + H-26 unlocked writer already show how subtle state-machine errors survive prose
review. "Simplify first; model-check what remains."

### 5. Local tool contracts far less mature than the governance contracts
The schemas are tiny; a schema is not a behavioral contract. Before freezing the worker as a coding-agent
design, give executable semantics to: workspace-root/path resolution, absolute paths + `..`, symlink
traversal, binary/invalid-UTF-8 files, max file/output size, atomic writes, permissions/mode preservation,
newline normalization, edit-multiple-match behavior, patch path-stripping + out-of-workspace files, shell
identity, fixed cwd, env allowlist, stdout-vs-stderr, exit-code/signal handling, process/resource limits,
background-process behavior, network policy. These affect reliability more than many digest-serialization
decisions. Port observable behavior + adversarial tests from the best existing coding agents.
**Add a model-visible surface identity** — the catalog digest excludes descriptions; add a `model_surface_digest`
over system/developer instructions + tool names/descriptions/schemas + provider-specific tool lowering +
compaction template/version + policy messages + relevant env facts; bind it to each attempt + E3. A binary
digest says "which executable?"; a surface digest says "which behavioral interface did the model receive?"

### 6. The exit test proves plumbing, not product viability
The three-leg bar (provider-send route; exact F59-bound local write; composed governed turn) is a good
protocol test, not enough to call the result a viable coding harness. It does not test task completion,
governance overhead/latency, F59 round-trip cost, context-loss damage, compaction of code-relevant state,
operator recovery from UNKNOWN, native-vs-simpler-baseline, or multi-agent help-vs-hurt (the scaling study:
multi-agent degrades sequential reasoning 39–70%; CodeTeam 34.6%/42.3% pass despite orchestration). **Add a
product/evaluation exit leg:** 10–20 small real-repo tasks; tests as success; simpler single-agent baseline;
completion rate, wall-clock, tokens, tool-calls, durable-record count, governance latency, operator
interventions, unknown outcomes, recovery after injected process deaths, duplicate-effect prevention; one
adversarial repo with an indirect prompt injection; one governed handoff to another seat. (Harness succeeds
through feedback loops + executable acceptance criteria + real testers, not architecture alone.)

### 7. The authority seam is mechanically useful but substantively empty
The eight-tool set is a dispatch identity, not a capability ceiling; because `bash` is included it does not
restrict effects. F59 mostly proves name-in-set + well-formed + run/turn/epoch + not-consumed = execution
integrity, not an authority decision — calling the milestone "governed local tool execution" overstates it.
Include ONE real fixed local policy without the whole Step-4 system: paths outside the workspace denied;
network denied from local execution; destructive FS op requires a hold; write to a protected path requires
operator approval; external publication / `git push` unavailable; secret-bearing paths unavailable to the
sandbox. Hard-coded + operator-ratified. Gives the user an observable reason for the machinery. Longer term,
static authority narrows by intent, never expands by it.

### 8. E3 is structurally bound but semantically too prose-heavy
Prose `claim` + binary `observed_outcome`; the evaluator rigorously checks applicability but does not make the
claim machine-verifiable. Before DSSE signatures, define typed predicates: `provider_request_matches_frozen_core`,
`provider_deny_caused_zero_transport`, `local_invocation_matches_effect_descriptor`,
`relay_record_committed_with_stamped_sender`, `no_alternate_credentialed_provider_route_observed` — each with
version, required structured inputs, observed facts, evidence locator, pass/fail/unknown, observer identity,
exact request/effect digest. Signing prose proves who signed the prose; typed claims should precede/accompany
the planned DSSE/SPIFFE/SLSA work. Observer identity is convention-bound (acceptable for internal MVP if the
claim stays modest).

## What the reviewer would ratify vs hold
**Ratify the architectural direction:** conductor isolated + narrow; provider bypasses it; m-10 owns lifecycle
not policy; m-8 holds provider creds; m-9 replaceable; stores separated by state family; unknown outcomes
first-class; one lane + one active turn; fixed catalog acceptable; native+MCP share contracts; H-17 census in
the lock.
**Hold the current all-artifact lock** for a bounded amendment: (1) a sandboxed local executor, Linux-only ok;
(2) canonical effect descriptors incl execution context; (3) frozen_core_digest joined through m-8/m-10/E3;
(4) typed model-visible run journal with settled-round recovery; (5) typed E3 claim ids; (6) a task-level
evaluation exit leg; (7) simplified broker restart OR blocking formal verification; (8) H-16 + H-26 repaired
before T4.
**Change lock granularity:** hard-lock cross-process wire contracts, durable identifiers, canonical encodings,
effect-record schemas; keep prompt wording, tool descriptions, compaction behavior, local-tool ergonomics,
retry tuning, user-surface behavior revisionable during the first implementation/evaluation cycle. Freezing
all ten realization documents under one full-reopen rule risks turning normal product learning into
constitutional amendment.

## Revised MVP shape + revised exit test
Sandbox executor between m-9 and the host (workspace-only FS, network-off, no peer sockets, effect descriptor
+ F59 ticket). Durable run journal + visible recovery state in m-10. Five exit proofs: **Utility** (completes a
representative repo task + tests, vs a baseline), **Containment** (a prompt-injected repo cannot read outside
the workspace, reach arbitrary network, access peer resources, or leave an untracked daemon), **Durability**
(crashes at every boundary → deterministic resume from last settled round or a visible typed unknown; never
duplicate effects or fabricated outcomes), **Governance** (provider request joined frozen-digest→authorization
→observed effect; local effect joined descriptor→ticket→outcome; relay committed with stamped identity),
**Operability** (operator sees what happened, why it stopped, which effects are unknown, what safely resumes).

## Direction of the project
- **The moat is the governance plane, not the model loop.** The coding loop/adapters/sandbox/checkpointing/
  tool-calling/subagents are commoditizing (OpenAI Agents SDK; Symphony orchestration spec). Frank's stronger
  position: "a durable governance plane between agent harnesses and execution sandboxes, with system-derived
  identity, lineage, authority decisions, effect bindings, outcome uncertainty, and evidence status." The
  native harness is the reference client, not the only environment Frank's governance is available in.
- **Keep MCP, add A2A as adapters, not foundations.** MCP annotations are untrusted hints. A2A (>150 orgs) is
  a task/message carrier, not courier authority. Frank adapter at the trust boundary; frank-stamped identity
  internally; explicit self-reported-vs-observed translation. Don't redesign the conductor around A2A.
- **Don't rush the multi-agent topology.** Keep the single-worker MVP until metrics exist; topology by task
  shape, not as an inherent upgrade. Establish a strong single-agent baseline + the cost of one governed
  handoff first.
- **Delay the Rust multiplexer.** Current risks are containment, recovery, evidence binding, usefulness — keep
  the surface all-Go / existing terminals until the runtime has real users.

## Final judgment
Architecturally promising, unusually self-critical, differentiated at the courier/evidence layer. The reframe
was right; the H-17 census is excellent; crash/unknown treatment is ahead of much of the ecosystem. But the
MVP spends too much complexity preserving app-process continuity and too little establishing the three
properties a coding-agent user cares about immediately: (1) the agent cannot escape its workspace/network
boundary; (2) a crash does not erase useful working state; (3) the task is completed at acceptable cost/latency.
Promote containment, resumable state, exact effect binding, and task-level evaluation into the MVP; simplify or
postpone the broker-survival protocol.

**Recommendation:** HOLD the stage-6 lock for a narrowly scoped amendment if this milestone will be called the
Frank harness MVP. Otherwise ratify under the narrower name "Step-3 governed-runtime integration prototype,"
and make sandboxing, run journaling, and request/effect digest joining hard blockers before external or
security-sensitive use. Frank's strongest future is the small, hard-to-bypass layer that lets many
orchestrators make fewer unverifiable claims.

## Cited sources (reviewer-supplied)
1. Microsoft Agent Governance Toolkit — LIMITATIONS.md (records attempts, not verified outcomes).
2. arXiv 2606.22528 — Governance Decay (compaction 0%→30–59% prohibited-action violations; pinning → 0%).
3. openai.com/index/running-codex-safely — managed sandbox, no open-ended outbound.
4. anthropic.com/engineering/claude-code-sandboxing — FS + network isolation both required.
5. arXiv 2607.05397 — Proof of Execution (exclusive effector credentialing + causally joined record).
6. openai.com/index/the-next-evolution-of-the-agents-sdk — externalized state, checkpoint/rehydrate.
7. docs.langchain.com/oss/javascript/langgraph/persistence — step-boundary persistence.
8. arXiv 2512.08296 — scaling study (multi-agent degrades sequential reasoning 39–70%).
9. arXiv 2606.22082 — CodeTeam (34.6%/42.3% pass rates).
10. openai.com/index/harness-engineering — feedback loops + executable acceptance criteria + real testers.
11. arXiv 2606.22916 — intent-governed authorization (narrow by intent, never expand).
12. blog.modelcontextprotocol.io/posts/2026-03-16-tool-annotations — annotations are untrusted hints.
13. linuxfoundation.org — A2A surpasses 150 organizations.
