# m-9 Model Runtime — AUDIT: the audit/promotion matrix (design-only)

**Owner:** m-9 (Model Runtime) — authored by m-9.planner (design-lead); adversarial return owed by m-9.implementer
**Cycle/phase:** Step-3 pre-build design sequence, §6 step 1 (AUDIT) · **Tier:** medium · **Evidence:** E1 · **Date:** 2026-07-14
**Dispatch:** `c7-audit-m-9` (parent `master-boot-m-9-planner`) · **Status:** **REV1** — the m-9.implementer adversarial return (`c7-audit-m-9/AUDIT-implementer-20260714-213834`, verdict revise) is FOLDED: all six blockers verified-then-accepted, row adjudications applied, four missing candidates added (§6 fold-log). Routed back for the narrow rev1 verification; no design authority, no lock
**Spec-of-record consumed (re-derived nothing):** `master/STEP-3-KICKOFF.md` (LOCKED, VP-co-signed `step3-prep/RECONCILE-orchestrator-reviewer-20260714-222000`) §1/§2/§5/§6; the m-9 charter `master/domains/m-9-model-runtime/README.md`; locked m-7 (`c4-design-m-7-conductor-core`), m-5 (`c3` archetype design), m-3 (`c2` observe design); the m-8 charter (emerging contract, design not yet authored); `master/RUNTIME-RESEARCH.md` (§0/§2/§6.2/§12).
**Method note (evidence honesty):** pi/opencode/frank surveys were produced by three read-only research lenses (sanctioned spawn class); their file:line claims are E0-as-received. I spot-checked the load-bearing anchors myself (pi `agent-loop.ts:155/170/174`, `ai/src/types.ts:138` `onPayload`, `agent/src/types.ts:201` `getApiKey`; opencode `effect/runner.ts:115-122` coalescing, `provider/provider.ts` apiKey options-bag `:308/:612`, `session/processor.ts:29` doom-loop; frank anchors cross-checked against the kickoff's own E1 cites). Rows below marked **[E1]** where the anchor is kickoff-cited or spot-checked; **[E1-lens]** where the citation is from a lens report I did not independently re-open. The implementer's adversarial return should treat [E1-lens] rows as re-verification targets. *(rev1: the implementer reopened the load-bearing [E1-lens] anchors in its return; the two evidence contradictions it found — G2 `subagent-permissions.ts:4-26`, E5 `timeoutMs` vs `maxRetryDelayMs` `ai/types.ts:153-176` — plus the C2 replacement cites (`native-runtime.ts:169-190`, `processor.ts:98-102`) were then re-verified by the planner directly against the bytes before folding.)*

---

## 0. Frame — what this audit decides

Decision (B), as folded (kickoff §1): pi/opencode are **prior art + conformance-fixture sources, NOT the spec**. This matrix walks their turn/session/tool-loop/context contracts plus the landed frank interfaces, and dispositions each candidate mechanism into:

- **PROMOTE** — adopt the contract/invariant shape as m-9 prior art (fresh implementation, frank-owned contract).
- **ADAPT-WITH-GOVERNANCE** — the mechanism survives, but only re-shaped around a named frank gate/seam.
- **REJECT** — the mechanism violates a frank invariant or crosses into another domain's locked ownership; named replacement or owner given.

Every row names the governance seam it crosses. The two charter lines this matrix enforces everywhere: **(i)** a parsed tool call stays **INERT** until the existing trusted authority/tool-exposure path authorizes it; **(ii)** m-9 runs **ON** m-7's substrate and never re-owns loop/recovery/config/guardrail, m-5's ceiling, or credentials/endpoints/egress/routing.

**Spawn-topology correction (honesty over the boot's shorthand):** pi's reference subagent impl is **not** in-process — it spawns separate `pi` OS processes (`packages/coding-agent/examples/extensions/subagent/index.ts:294,333-339` [E1-lens]). What frank rejects is more precise: **anonymous, ungoverned spawn** — children with no seat identity, no channel-stamped FROM, no ceiling-at-spawn record (pi: anonymous `--no-session` children; opencode: in-process child sessions distinguished only by `session_id`/`parent_id`, `tool/task.ts:142-158` [E1-lens]). The reject rationale is the missing governance, not the process boundary per se.

---

## 1. The duplicate/already-built gate (frank at `s11-close@502e06c`)

**Greenfield CONFIRMED for the m-9 surface.** Exhaustive sweep of `frank/internal/` + `frank/cmd/` for provider/LLM/turn-loop/session-state/context-window/compaction code: **zero hits**; all near-matches are false positives (intake "completion" channels, fs-observation "lanes", the egress classifier's model-name byte patterns, operator-decision "prompters") [E1-lens, sweep query enumerated in the survey]. No `net/http` client exists; the only network surface is the unix-socket seat channel.

**Adjacent machinery that ALREADY EXISTS and must be consumed, never rebuilt** (the already-closed list — rebuilding any of these is a review blocker):

| landed mechanism | where | what m-9 does with it |
|---|---|---|
| serialized commit loop + intake writer | `internal/engine/loop.go:41-109`, `internal/intake/writer.go:87-105` | ALL m-9 state mutations commit through `intake.Writer.Submit` → the one loop; m-9 never opens a second write path |
| internal-subsystem extension pattern | `ResummonHandler`/`ApprovalHandler` wrappers, `internal/engine/resummon.go:265-298`, `approval.go:148-189`, composed `cmd/frank/main.go:282` | the candidate integration shape for m-9-originated commands (own Verb, `Seat=="system"` required) — a DESIGN question, not a decided answer (§5 Q1) |
| `ServiceWhileBlocked` + expiry scheduler | `loop.go:211-250`, `internal/engine/expiry.go` | the landed precedent for awaiting a long operation under operator-disposable timeout — provider-await is structurally similar (§5 Q6) |
| observe gate + pluggable predicate | `internal/observe/gate.go:90-150`, `Env.Evaluate` `gate.go:38-41`, inline at `submit.go:115` | m-9's turn record enters observe as a candidate; NO streaming observer exists today (§5 Q3) |
| authority headers, enforcement-absent | `fieldspec/registry.json:161-166` (`slot_in`, `seat_archetype`, `authority_ceiling`, `capability_prior_snapshot`, `template_ref`); no landed code reads `authority_ceiling` to gate anything | the data model for "inert until authorized" is landed; the ENFORCEMENT is a **still-open cross-domain seam owned by m-5 (semantics) + m-7 (hosting the authoritative check/execution door)** — NOT an m-9 obligation (charter: m-9 "does not set or enforce" the ceiling). m-9's still-open piece is the **inert governed request** + consuming the authorization result *(rev1 blocker-1 fix)* |
| operator approval gate | `ApprovalPrompter`, `approval.go:76-89,177-187`, default-denied | the landed ask/hold machinery — opencode's permission-ask maps here (row C3); no second approval subsystem |
| park/wake + nudge | `resummon.go:49-230`, `channel.Server.PushTo` `server.go:167-180` | the wake path; follow-up queueing overlaps the mailbox (row D1) |
| egress outbox + classifier | `internal/egress/egress.go:100-144`, `rules.go:22-44`, model-name→operator-only carve-out `egress.go:84-89` | confirms kickoff §1: the away-email class **cannot** front a provider request; the provider-request egress class is the m-3 owner amendment (§1a) m-9 consumer-reviews |
| trusted config, no provider contract | `internal/config/config.go:185-261`; member enum `{engine, fieldspec, catalog}` hardcoded `submit.go:463` | confirms kickoff §1b: NO credential/endpoint/secret/serving-profile member exists — the m-7 owner amendment gates any lock |
| seat surface + channel-stamped FROM | `cmd/frank-mcp/mcp.go:134,325-363`, `seat.Stamp` at `submit.go:57` | whatever process shape m-9 takes, its identity/guardrail story rides this, unchanged |

**FieldSpec detail sharpening the kickoff cite:** `chosen_model` is a **column** of the `routing_assignments` row_array at `registry.json:175` (cols incl. `declared_bucket, chosen_model, pin_mode, seat_archetype, authority_ceiling`), not a top-level scalar; the model-identity scalar is `model_name` (`:173`, system-owned, render-absent). **No provider/serving/compat/lane field exists anywhere** — consistent with the mandatory m-4/m-2 exact-lane amendment (kickoff §3) [E1-lens; kickoff-consistent].

---

## 2. THE PROMOTION MATRIX

Verdict key: **P** = promote · **A** = adapt-with-governance · **R** = reject. Seam = the governance boundary the row crosses.

### A. Turn-loop state machine (the m-9 core)

| # | candidate (source, evidence) | verdict | disposition + governance seam |
|---|---|---|---|
| A1 | **Two-level loop shape**: outer follow-up restart / inner turn iterator `while (hasMoreToolCalls \|\| pending)` — pi `agent-loop.ts:155-275` [E1]; opencode `runLoop` storage-re-deriving `while(true)` `prompt.ts:1081-1341` [E1-lens] | **P** (narrowed) | The **logical lifecycle** m-9 designs fresh: assemble → adapter call → consume normalized stream → **m-3 observation** → tool-request round → iterate to typed finish. *(rev1 narrow:)* turn boundaries are frank commit points AND **durable tool-request checkpoints** (C8) are named lifecycle stations — no donor loop implementation and no second scheduler is promoted. Seam: m-9→m-7 (commit), m-9→m-3 (observe-before-action, C9). |
| A2 | **Exactly-one-active-turn-per-session**: opencode `Runner.ensureRunning` returns the in-flight run's Deferred instead of starting a second turn `effect/runner.ts:115-138` [E1]; pi busy-throw `agent-harness.ts:609` [E1-lens] | **P** (invariant only, narrowed) | Promote **exactly-one-active-logical-turn** as the invariant. *(rev1 narrow:)* opencode's **in-memory Deferred coalescing is NOT promoted as the frank shape** before Q1/process placement resolves — enforcement of the invariant must be recovery-safe and rides the m-7-hosted substrate, and governed inbox entries remain distinct durable inputs (never silently merged by a coalescer). Seam: m-7 (hosting/recovery), Q1. |
| A3 | **Turn lifecycle event vocabulary** (`agent_start/turn_start/message_update/tool_execution_*/agent_end`, pi `types.ts:415-430` [E1-lens]; opencode part-typed events [E1-lens]) | **A** | m-9 owns a normalized **turn-event vocabulary** layered over m-8's normalized provider events — but shaped for consumers frank actually has: m-3 observation + the projection surfaces, not a UI SDK. Seam: m-8↔m-9 (who owns which event layer — consumer-lock set). |
| A4 | **Typed finish/stop conditions**: pi `StopReason {stop,length,toolUse,error,aborted}` `ai/types.ts:380` [E1-lens]; opencode finish + `lastUser.id > lastAssistant.id` re-derivation `prompt.ts:1111-1130` [E1-lens] | **A** | Wire-level finish/error vocabulary is **m-8's contract** (charter); m-9 owns the **turn-level terminal semantics** (no live tool requests + no newer inbound + typed finish ⇒ turn complete). Both promoted as fixture sources. Seam: m-8↔m-9; also maps onto frank's exactly-one-outcome discipline (m-7 §6). |
| A5 | **Loop state re-derived from persisted truth each iteration** (opencode re-reads storage per step; steering = persist+coalesce) [E1-lens] | **P** | Isomorphic to frank's store-is-truth/projections-derived DNA (m-7 §4). Promote as the m-9 session-state principle: the turn loop's authoritative state is the committed record set, memory is cache. Seam: WHERE the session/transcript record set lives is a design question (§5 Q2), the principle is not. |
| A6 | **Bounded turns + doom-loop tripwire**: opencode `maxSteps`/`MAX_STEPS_PROMPT` `prompt.ts:1178-1179` [E1-lens], `DOOM_LOOP_THRESHOLD=3` identical calls → forced ask `processor.ts:29,353-380` [E1] | **A** | Promote both as governed-runtime tripwires, re-homed: exhaustion/repetition ⇒ a **typed outcome / operator gate** (frank's held/approval machinery), never an inline UI ask or a silent prompt injection. Seam: m-6 surface consumes; m-5 may own thresholds as archetype policy. |

### B. Session / context state

| # | candidate | verdict | disposition + seam |
|---|---|---|---|
| B1 | **Append-only session history + derived projection**: pi session tree w/ typed entries + leaf pointer, compaction as an ENTRY, history projected out but never deleted (`harness/types.ts:409-420`, `session.ts:57-135`) [E1-lens]; opencode event-sourced writes + projector (`projector.ts:262-322`) [E1-lens] | **P** (revised) | The strongest convergence in the corpus: both donors independently landed on frank's own append-only + derived-projection shape. Promote pi's **compaction-as-recorded-entry** (never destructive) specifically. *(rev1 blocker-2 fix:)* **canonical session/context transitions commit through the m-7-owned substrate — an m-9-owned durable session log as a second canonical truth is REJECTED** (violates one-writer/canonical-record + m-9's no-re-own of crash-atomic commit/recovery). What stays designable: record schema, projection shape, volume management — WITHIN the substrate (§5 Q2, narrowed). |
| B2 | **Deterministic per-turn context assembly** (system prompt + filtered history + tool defs): pi `buildSessionContext` `session.ts:125-135` [E1-lens]; opencode `prompt.ts:1257-1286` [E1-lens] | **A** (revised) | Promote determinism; add the frank obligations: (i) assembly is **reproducible + snapshot/digest-pinned** (kickoff §4 pinning discipline applies to what was actually sent); (ii) *(rev1 blocker-1 fix:)* m-9 assembles the offered tool set **from a trusted, already-ceiling-filtered exposure it receives** (the codex ToolExposure principle m-7 §8.1 already promotes) — **m-9 does not interpret or enforce the ceiling itself**; the authoritative check sits on the only execution path, owned m-5 (semantics) / m-7 (hosting) (C2). Seam: m-5/m-7 (the exposure contract m-9 consumes). |
| B3 | **Compaction/summarization mechanics**: pi trigger math (`contextTokens > window − reserve`), real-usage-first token accounting, valid cut points never mid-tool-result, split-turn handling (`compaction.ts:200-203,160-197,265-381`) [E1-lens]; opencode overflow math + tail preservation + overflow-replay + prune (`overflow.ts:10-34`, `compaction.ts:188-239,310-326,243-287`) [E1-lens] | **A** (narrowed) | *(rev1 narrow:)* Promote the **named mechanics as conformance FIXTURES, not wholesale policy** — the split-turn/tail-preservation/prune heuristics are candidates the DESIGN must prove, not inherited defaults. The governance re-shape stands: **compaction is a lossy context mutation and must be a committed, observable event** (before/after digests, trigger recorded) riding m-7/m-3 governance, never silent; **the summarizer turn is itself a governed model turn** passing the same gates (it is a provider send!); opencode's auto-continue synthetic user message must become a **system-stamped** injection, never forged author-side (m-1 FROM discipline). Seam: m-3 (observability), m-1 (provenance), m-8 (the summary call is a normal governed send), m-5/operator (trigger POLICY — RUNTIME-RESEARCH §12 #6); m-9 owns deterministic mechanics + typed state only. |
| B4 | **Token/limit facts from the catalog, usage from the wire** (opencode `limit.input`/`context` per model; pi real-usage-first) [E1-lens] | **P** | Facts (`context_window`, `max_output`) come from the **m-8 lane catalog**, usage from m-8's normalized usage events; m-9 hardcodes nothing. Seam: m-8 single-writer facts (kickoff §3). |
| B5 | **Reasoning-replay round-trip**: strict servers reject multi-turn tool replay missing `reasoning_content` (RUNTIME-RESEARCH §6.2, primary-sourced) [E1] | **A** | m-8 owns reasoning-replay normalization (charter); m-9's session state must **round-trip opaque reasoning blobs** across turns/reopens without inspecting them. Conformance fixture: the ≥8-step close/reopen canary (RUNTIME-RESEARCH §12 #2/#5). Seam: m-8↔m-9 contract; m-3 lane-qualification consumes the canary. |
| B6 | **Session fork/revert + filesystem snapshot restore**: opencode `fork`/`revert`/`unrevert` + snap restore + hard-delete cleanup (`revert.ts:70-134`, `prompt.ts:1056`) [E1-lens]; pi tree branching (`moveTo`) [E1-lens] | **R** (V1) / carry | Out of the spine: revert **hard-deletes history and mutates the workspace** — collides with append-only + observe; fork has no V1 consumer. Record as a named carry (fork may return with native governed spawn). Seam: m-7 store immutability; m-3 observation of workspace mutations. |

### C. Tool-call path (the charter's sharpest line)

| # | candidate | verdict | disposition + seam |
|---|---|---|---|
| C1 | **Schema validation of parsed args before anything else**: pi TypeBox compile+cache, coerce, formatted error back to model (`ai/utils/validation.ts:278-310,232-241`) [E1-lens]; opencode AI-SDK inputSchema [E1-lens] | **P** (narrowed) | Deterministic validation is step one; a malformed call **never reaches the authority path** — it returns a typed error result to the model. *(rev1 narrow:)* validation may precede authorization only while the object is **inert**, and the **malformed-call feedback is itself governed model context/output** — the schema-error path bypasses neither observation nor provider-send governance (the error result rides the next governed send like any other context). Seam: m-3/m-7 (the feedback path is governed); the validator is m-9-owned deterministic code, not model-supplied. |
| C2 | **Execute-on-parse, in-process**: pi `executePreparedToolCall` calls `tool.execute` directly (`agent-loop.ts:668-709`) [E1-lens]; opencode's **direct executor** wires `item.execute` straight into the SDK tool shape (`session/llm/native-runtime.ts:169-190` [E1]), and the AI SDK **may execute tools internally before emitting start-step events** (the in-source warning, `processor.ts:98-102` [E1]) | **R** (revised) | THE charter line. *(rev1 blocker-1/-3 fix — the replacement sequence, with m-9 ownership STOPPING at the inert request:)* normalized model output → **m-3-owned observe decision at m-3's defined granularity (C9 — no unobserved tool-call bytes become actionable)** → complete/validated **inert request with stable identity (C7/C8)** → **trusted m-5/m-7 authorization** (m-5 ceiling semantics, fail-closed, absent-default = floor; m-7 hosts the check + the execution door + the unprivileged executor — the m-3 suite-class precedent `internal/executor/executor.go`) → execution → result folds back as data. **m-9 does not evaluate the ceiling, own the check, or own the executor.** opencode's SDK-executes variant is doubly rejected: execution **inside the adapter layer** — the m-8↔m-9 seam requirement: **adapters surface tool CALLS as normalized events and carry no execution capability.** Seam: m-3 (observe-before-action), m-5 (semantics), m-7 (check/door/executor), m-8 (no-exec adapters), kickoff §5 negative leg (above-ceiling ⇒ **zero execution**). |
| C3 | **Blocking permission-ask with once/always/reject+feedback**: opencode `Permission.ask` Deferred-blocked, `reply` once/always/reject w/ `CorrectedError` feedback, cascade-reject (`permission/index.ts:28-167`) [E1-lens] | **A** (mechanics) / **R** (as a new subsystem) | The ask-shape maps onto **landed frank machinery**: ApprovalPrompter + held + park/wake/resummon. Do NOT build a second ask subsystem (product-overlap row). Promote two details into the design conversation: **reject-with-feedback** (the correction reaches the model as a typed result) and **cascade-reject of pending asks on session cancel**. "Always-allow" rule accumulation = **policy**, owned by m-5/operator config, never learned silently at runtime. Seam: m-6 (surface), m-5 (policy), m-7 (gate hosting). |
| C4 | **Opt-in gating inside tool execute** — a tool that never calls `ctx.ask` is never gated (opencode `tools.ts:81-132`) [E1-lens] | **R** | The **named anti-pattern**: the inverse of observe-as-send ("you can't forget the check that is the gate on the door", m-3 §2). The gate must sit on the **only** execution path, structurally outside tool code. This row is the fixture-source for the m-9 negative leg: a tool implementation with no gate call still cannot execute ungated. Seam: m-3 DNA applied to m-9's execution door. |
| C5 | **Typed tool-result folding** (`toolResult` role, `isError`, interrupted-orphan marking + exclusion from replay: opencode `processor.ts:539-597`, `prompt.ts:96-102` [E1-lens]; pi `createToolResultMessage` [E1-lens]) | **P** | Promote, including the abort-orphan honesty (a tool killed mid-flight is marked interrupted, never replayed as if answered). Seam: feeds m-3 observation of the turn record. |
| C6 | **Parallel tool execution w/ deterministic transcript order** (pi parallel mode, results re-emitted in source order `agent-loop.ts:491-556`) [E1-lens] | **R** (V1) / carry | V1 = serial (the simplicity rule; the vertical turn needs one real governed turn, not throughput). If parallel ever admitted, pi's source-order determinism is the promoted invariant. Seam: none yet; carry. |
| C7 | *(rev1 add — implementer missing-candidate)* **Tool-call completion/finalization**: pi refuses to execute tool calls from a `length`-truncated message — all fail unexecuted with a re-issue instruction (`agent-loop.ts:376-407` [E1-lens, implementer-reopened]); streamed tool-call deltas accumulate but only a COMPLETE call finalizes | **P** | **Partial or truncated tool-call deltas never create an executable request.** pi's truncated-refusal is the positive fixture. A request object exists only after the call is complete, schema-valid, and observed (C9). Seam: m-8 (partial-stream semantics in its contract), m-9 (finalization discipline). |
| C8 | *(rev1 add — implementer blocker 5)* **Stable tool-request identity + no-replay across recovery/retry**: the donor danger is concrete — opencode may execute tools mid-stream before lifecycle events (`processor.ts:98-102` [E1]), persists pending calls by provider-supplied ID (`processor.ts:216-253` [E1-lens]), retries the enclosing stream (`processor.ts:627-676` [E1-lens]), and labels unfinished calls interrupted only at cleanup (`processor.ts:539-597` [E1-lens]) — provider retry can duplicate an effect | **A** (new invariant) | A complete tool call yields a **stable, session/turn-bound request ID**; authorization + execution disposition are **durable** (typed states: authorized/executing/completed/interrupted/**unknown-effect**); **recovery/retry can never replay a completed or unknown-effect action** (an unknown effect is surfaced, never silently retried); duplicate/colliding call IDs **fail closed**. Landed prior art: frank's intake content-hash dedup + `intake_id`-referencing outcomes (m-7 §2.2 exactly-once-effect). Ownership split: m-9 (request identity/state) · m-5 (semantics) · m-7 (hosting/commit/recovery). |
| C9 | *(rev1 add — implementer missing-candidate; charter line made a row)* **Observe-before-action over ALL model output**: text, reasoning, and tool-call output each have a defined m-3 vantage; the charter applies the observe gate to model output and the spine names `streamed-output observation` (kickoff §5) | **A** (locked floor) | **No model-output byte becomes recipient-visible or actuator-actionable before the m-3-owned observation decision at m-3's chosen granularity.** A tool call cannot cross from model output into an actionable request on an unobserved side channel. Boundary-only observation is admissible ONLY if zero partial bytes reach any recipient or actuator first (Q3, floor fixed). Seam: m-3 owns the decision + granularity; m-9 shapes the candidate and holds the door. |

### D. Steering / interrupt / cancel

| # | candidate | verdict | disposition + seam |
|---|---|---|---|
| D1 | **Follow-up queueing** (pi `followUp` queue drained at would-stop `agent-loop.ts:263-267` [E1-lens]; opencode persist+coalesce — no queue object at all [E1]) | **R** as new queue / **A** as contract | **Product-overlapped: frank's mailbox + pipe-wake IS the follow-up queue** (m-7 §8.3; `PushTo` nudge landed). The m-9 contract is: at turn end, the loop consults the seat's governed inbox — one queue in the system. opencode's persist+coalesce validates exactly this shape. Seam: m-7 delivery, m-6 scheduling. |
| D2 | **Steer-at-boundary** (pi `steer()`: inject after current tool round, before next model call `agent-loop.ts:167,259`) [E1] anchored | **A** (hook only, spine) | Steering is a **named Step-3 carry, NOT spine** (kickoff §7). But the turn loop must land with the **boundary hook designed-in** (the poll point between tool round and next assembly) or the carry forces a re-cut. pi's steer/follow-up asymmetry is the promoted contract; injection provenance is m-1-stamped. Seam: m-5 `accepts_interjection` declaration; m-6 owns the surface; m-9 owns the injection point. |
| D3 | **Abort/cancel semantics**: signal threading; partials persisted + finalized `stopReason:"aborted"`, never discarded (pi `agent-loop.ts:348-373` [E1-lens]); in-flight tools marked interrupted after bounded wait (opencode 250ms cleanup `processor.ts:539-597` [E1-lens]); interrupt distinguished from failure (`Cause.hasInterruptsOnly` [E1-lens]) | **A** (narrowed) | Wire-level cancellation = **m-8 contract** (charter). m-9 owns turn-level disposition: a cancelled turn reaches a **typed terminal**, partial output is **committed honestly** (labeled partial, never replayed as complete), interrupted ≠ failed, and an interrupted tool's effect state is C8's `unknown-effect`, never silently retried. *(rev1 narrow:)* honest partial COMMIT passes only with the explicit rule that **no partial reaches a recipient or actuator before the m-3-owned observation decision** (C9). Soft-interrupt/cancel-redeliver = Step-3 carry riding this. Seam: m-8 (wire), m-3 (observe + honest labeling), m-7 (commit of the terminal). |
| D4 | **Transitive descendant cancel** (opencode BFS over background-job session chains `run-state.ts:111-143`) [E1-lens] | carry | No spawn in the spine ⇒ no descendants. Named carry for native governed spawn; the invariant (cancel reaches the whole governed tree) folds into the spawn design, m-5 lineage semantics. |

### E. Error / retry / timeout

| # | candidate | verdict | disposition + seam |
|---|---|---|---|
| E1 | **Errors-as-terminal-messages; the stream never throws** (pi `StreamFn` contract: failures arrive as a final message with `stopReason error/aborted` `types.ts:24-31` [E1-lens]) | **P** | Isomorphic to frank's typed-outcome discipline ({accepted, rejected, held}; m-7 §6 "a crash is not a yes"). Promote as the m-8↔m-9 event-contract posture (m-8 authors; m-9 consumer-reviews for this property). |
| E2 | **Wire-level retry/backoff**: pi delegates to vendor SDK `maxRetries` (`ai/types.ts:165-176`) [E1-lens]; opencode owns a full policy — retryable classifier (status≥500/`isRetryable`), never-retry-overflow, retry-after honored, exp backoff capped 30s, retry status published (`retry.ts:35-199`) [E1-lens] | **A** (split) | **Wire retry/idempotency = m-8's contract** (charter — m-9 never re-owns). m-9 owns only the **turn-level disposition** of a retried/failed call: bounded, **recorded** (a retry is an observable event, not silence — opencode's published retry status is the promoted shape), typed terminal on exhaustion. opencode's free-tier upsell special-casing: rejected, product-specific. **Grill candidate (§5 Q4): retry × final-wire-authorization** — is attempt N+1 the same authorized send or a fresh authorization? (kickoff §1a: no mutation after final auth; identical-bytes replay vs re-auth is an m-3 amendment consumer-review question m-9 must put on the table.) Seam: m-8 (mechanics), m-3/m-7 (the §1a boundary). |
| E3 | **Never retry context-overflow; route to compaction** (opencode `retry.ts:70`) [E1-lens] | **P** | Clean typed-cause discipline. Promote with B3's governed-compaction shape. |
| E4 | **Content-filter/refusal promoted to a typed error** (opencode `ContentFilterError` `prompt.ts:1301-1308`) [E1-lens] | **P** | A refusal is a first-class typed outcome, honestly labeled — never coerced into "empty response". Seam: m-3 honest labeling; m-8 normalizes the wire signal. |
| E5 | **Provider-call timeout** (pi `timeoutMs` — the actual HTTP request-timeout field, `ai/types.ts:153-157` [E1, rev1-verified]; frank's landed expiry scheduler w/ operator kill/extend disposition `internal/engine/expiry.go` [E1-lens]) | **A** (evidence revised) | *(rev1 blocker-6 fix: the prior cite, `maxRetryDelayMs` `ai/types.ts:165-176`, caps a server-requested RETRY delay and fails fast when it exceeds the cap — retry-delay capping is NOT call-timeout semantics and is not promoted as such.)* m-9 must not hang a turn forever: promote frank's OWN expiry pattern (operator-disposable timeout) over inventing one; **m-7 owns timer hosting, m-9 owns only the typed turn disposition.** Whether an awaiting turn parks (m-6) or holds (`ServiceWhileBlocked`) is §5 Q6 — m-9 adds no scheduler either way. Seam: m-7 (scheduler hosting), m-6 (park semantics). |
| E6 | *(rev1 add — implementer missing-candidate)* **Backpressure / bounded buffering**: timeout/backpressure + partial-stream semantics are on the kickoff's normative provider-contract list (§1) and were absent from the original 36 rows | **A** | **m-8 owns wire-level backpressure mechanics** (its contract); m-9's obligation is to **consume the normalized stream without unbounded transcript/event accumulation and without ever blocking m-7's commit loop** (bounded buffers, typed overflow disposition — never silent drop). Seam: m-8 (wire), m-7 (loop-never-blocked), m-9 (bounded consumption). |

### F. Extensibility hooks — the bypass surface (kickoff §1c "no pi-shaped escape callback")

| # | candidate | verdict | disposition + seam |
|---|---|---|---|
| F1 | **`onPayload` / `before_provider_payload`** — replace the raw provider request body just before HTTP, after all validation (pi `ai/types.ts:138` [E1]; harness last-write-wins `agent-harness.ts:277-292` [E1-lens]) | **R** | The canonical §1a violation: mutation after the final authorization point. NO seat- or extension-registered payload mutation exists in frank's runtime, period. Any legitimate per-provider shaping is m-8 **translation** — deterministic, versioned, and sequenced per §1a (pre-translation check → translate/bind → **final authorization → zero mutation → wire**). Seam: m-3 §1a amendment; m-8 translation ownership. |
| F2 | **`before_provider_request`** stream-options patch — headers/transport/timeout/fetch swap (pi `harness/types.ts:99-105`) [E1-lens] | **R** | Headers/endpoint/transport are m-7 trusted-config + m-8 endpoint-binding territory (§1b: endpoint allowlisting/binding). Nothing runtime-side patches transport. Seam: m-7 credential/config amendment. |
| F3 | **`getApiKey` dynamic per-request resolution** (pi `agent/types.ts:201` [E1]) | **R** | Credentials never enter m-9 (charter: "m-9 never handles a secret"). Sourcing/rotation/redaction = the m-7 §1b amendment; binding = trusted-side after final auth. Seam: m-1 secret boundary, m-7. |
| F4 | **The options-bag coupling**: one deep-merged struct carries factual limits AND `apiKey`/`baseURL`/`fetch`/timeouts, spread into the SDK factory (opencode `provider.ts:1029-1054` [E1, kickoff-cited], `resolveSDK :1712-1792`, spot-checked `:308/:612`; `toPublicInfo :1072` existing to strip it = the hazard acknowledged in-source) | **R** (structurally) | The exact anti-pattern the kickoff's single-writer split exists to kill: m-8 catalog = facts (no secrets in catalog bytes), m-7 config = credentials, m-4 = policy. The m-9-facing consequence: **m-9's interface to m-8 accepts a pinned lane reference/descriptor, never a merged runtime bag** — if the adapter contract exposes an options-bag, m-9's consumer review rejects it. Seam: m-8↔m-9 contract; m-8↔m-7 credential seam. |
| F5 | **`afterToolCall`/`tool_result` result rewriting** (pi `types.ts:269-281`) [E1-lens] | **R** (silent) / **A** (declared) | Silent result substitution breaks observation (the transcript lies). Any transform (truncation, image resize — opencode does both) must be **declared and recorded** on the record (m-3's `degradation_notes` vocabulary is the landed home). Seam: m-3. |
| F6 | **`transformContext`/`convertToLlm` arbitrary context rewrite** (pi `types.ts:143-191`) [E1-lens] | **R** (seat-registered) / **A** (owned projection) | Per-provider message projection is a legitimate, deterministic m-8 translation duty; pi's *caller-supplied arbitrary* transform is the rejected part. m-9's own history-projection (compaction filtering, B1/B3) is deterministic, versioned, and observable. Seam: m-8 (translation), m-3 (observability). |
| F7 | **Runtime plugin/middleware registration** (opencode `plugin.trigger("tool.execute.before")` [E1-lens]; pi harness `.on()` hook registry [E1-lens]) | **R** | frank's extension seam is **additive config/registry data** (the m-3 check-registry pattern; m-5 registry-as-config), never runtime code hooks in the governed path. Seam: m-7 config model (restart-only, digest-pinned). |

### G. Spawn topology

| # | candidate | verdict | disposition + seam |
|---|---|---|---|
| G1 | **pi subagents**: separate OS processes, anonymous (`--no-session`), stdout-JSON, SIGTERM/KILL, concurrency-capped parallel/chain topologies (`subagent/index.ts:267-429,530-664`) [E1-lens]; **opencode task**: in-process child sessions, same store, child ruleset = parent denies carried + own permissions (NOT a narrowed copy of a parent maximum — see G2) (`tool/task.ts:104-240`) [E1-lens] | **R** (both, as-is) | Neither child is a **governed seat**: no channel-stamped FROM, no minted identity, no ceiling-at-spawn record, no observe gate on child output. Native governed spawn = the named Step-3 carry (kickoff §7) riding m-5's `orchestrator_lead`/authority-ceiling-at-spawn + Part-F mechanics. **V1 bakes no spawn.** Seam: m-1 (identity), m-5 (ceiling-at-spawn), Part F (lineage). |
| G2 | **Spawn permission derivation** (opencode `deriveSubagentSessionPermission`, `agent/subagent-permissions.ts:4-26` [E1, rev1-verified]) | **R** as donor semantics / **already-closed** as principle | *(rev1 blocker-4 fix — the prior "P (principle corroboration)" was WRONG:)* the source is explicit — "Parent agent restrictions only govern that agent; the subagent's own permissions determine its capabilities." It carries selected parent **denies** plus default `todowrite`/`task` denies; it does **not** bound the child by a parent positive maximum. That is NOT m-5's locked `child ≤ parent MAX` monotonic law — opencode is a **negative fixture for incomplete inheritance**, not corroboration. The principle itself is **already-closed**: m-5 owns and locked it (ceiling tighten-below/never-loosen, §5/GL-2). Nothing to build in the spine. Seam: m-5 (already locked). |

---

## 3. Planner audit deliverable (protocol form)

```text
4-bucket verdict (rev1 — folds the implementer return):
- still-open: the m-9-OWNED logical surface — the model-turn/session/context state machine;
  deterministic context assembly + compaction mechanics (typed state, fixtures-not-policy);
  normalized-turn-event consumption; INERT, stable-identity tool-execution request construction +
  request/result state (C7/C8); typed turn-level finish/error/cancel/timeout disposition; bounded
  stream consumption (E6). The m-3 streamed-output contract and the m-5/m-7 authorization/execution
  interface are OPEN CROSS-DOMAIN SEAMS, not m-9 ownership (rev1 blocker-1: ceiling enforcement is
  m-5 semantics + m-7 hosting; m-9 consumes). Greenfield confirmed by exhaustive sweep (§1).
- already-closed: the substrate the runtime rides — §1 table (commit loop + intake, system-verb handler
  pattern, ServiceWhileBlocked + expiry, observe gate + Env.Evaluate, ApprovalPrompter, park/wake +
  PushTo, egress outbox, trusted config + digest, seat surface + FROM stamp); m-5's monotonic
  child≤parent ceiling law (G2, rev1). Consume, never rebuild.
- product-overlapped: follow-up queueing (frank mailbox/wake, D1) · ask/approval subsystem (landed
  approval/held/park, C3) · wire retry/idempotency/cancellation/finish vocab + backpressure (m-8
  charter, A4/D3/E2/E6) · spawn (m-5 + Step-3 carry, G1) · transport/credentials/endpoints (m-7 §1b,
  F2/F3) · provider-request egress (m-3 §1a amendment, F1) · facts catalog (m-8, B4) · AND (rev1) any
  proposed m-9 durable store, scheduler/lease manager, ceiling evaluator, executor, recovery ledger,
  or streaming-observer implementation — those belong to the named existing owners/seams.
- recommended-next: m-9.implementer NARROW rev1 verification (confirm blockers 1-6 closed, the four
  added candidates C7/C8/C9/E6, and the row adjudications applied — not a full re-audit); then master
  reconciliation; then DESIGN (kickoff §6 step 2) against the verified matrix.
PRIMARY_BUCKET: still-open

Duplicate/already-built gate: run (§1) — m-9 surface greenfield; adjacent landed set enumerated with
the consume-not-rebuild obligation per row. No promote-existing recommendation contradicts a bucket.

Boundary contract:
Writes: this audit artifact (design-input; no code, no frank/ bytes).
Reads: references/pi + references/opencode (lens-mined, anchors spot-checked); frank@502e06c landed
  interfaces (lens-mined); locked m-7/m-5/m-3 designs + m-8 charter + RUNTIME-RESEARCH (direct read).
Target entity: the m-9 DESIGN (kickoff §6 step 2) + the m-8↔m-9 consumer-lock seam inputs.
Downstream consumer: m-9.implementer (adversarial return) → reconciled matrix → m-9 DESIGN; m-8 pair
  (rows A3/A4/B4/B5/C2/E1/E2/F4 name their contract obligations); the three §6-step-4 owner amendments
  (rows E2/F1/F2/F3 are consumer-review inputs).
Contract: the promotion matrix rows (verdict + seam per mechanism).
Proof: E1 file:line per row ([E1] spot-checked / [E1-lens] as labeled; method note).
No-consumer action: none — every row names its consumer (design section or seam owner).

Design recommendation (NOT a design; the shape the evidence points at, for the DESIGN phase to prove; rev1):
a LOGICAL turn loop outside m-7's critical section (a model turn is seconds-to-minutes; m-7 §2.3
forbids stalling the loop) whose canonical state transitions commit THROUGH the m-7 substrate
(A1/A5/B1), holding the exactly-one-active-logical-turn invariant recovery-safely (A2), assembling
requests from the trusted ceiling-filtered exposure it consumes (B2), with the C2 rev1 sequence —
normalized output → m-3 observe decision → complete/validated inert request with stable identity
(C7/C8) → m-5/m-7 authorization → m-7-hosted unprivileged execution — recorded compaction as fixtures
(B3), typed terminals everywhere (A4/D3/E1-E5), bounded consumption (E6), zero seat-reachable
mutation hooks (F1-F7), and no spawn (G1). Process placement (in-conductor subsystem vs separate
governed runtime process) is DELIBERATELY left to DESIGN + grill (§5 Q1) — both shapes satisfy this
matrix.

Evidence by claim: inline per row; method note governs [E1] vs [E1-lens] weight.

Risks / reject-or-narrow gates (review blockers for the eventual design; rev1 — adopts the
implementer's hardened set):
1. No m-9 ceiling evaluator, executor, scheduler, durable writer, recovery path, trusted-config
   reader, or fourth terminal token (no re-own of m-7/m-5 property in any form).
2. No model-output byte becomes recipient-visible or actuator-actionable before the m-3-owned
   observation decision at the chosen granularity (C9).
3. No partial/truncated/duplicate/unknown-effect tool call executes; no retry or recovery replays a
   completed or indeterminate effect (C7/C8) — including via the adapter (C2's m-8 no-exec
   requirement) or a tool that "forgets" to ask (C4).
4. No adapter execution callback and no seat-reachable request/result/context mutation hook
   (F-section) — the §1c named negative.
5. No credential, endpoint, transport knob, routing judgment, or policy mutation in the m-9 surface
   (F2-F4).
6. No silent context mutation: compaction/truncation/injection without a committed, labeled record
   (B3/F5); no second queue/ask/retry subsystem where frank or m-8 already owns one (D1/C3/E2).
7. No donor spawn topology in V1; later spawn must preserve stamped identity, lineage, and monotonic
   ceiling-at-spawn (G1/G2).

Questions for Implementer/operator: §5 below.
```

---

## 4. What this matrix hands the m-8 pair (seam obligations surfaced, not imposed)

From the m-9 consumer side, the adapter contract m-9 will consumer-review needs: normalized events with **no execution capability** (C2); typed finish/error vocabulary incl. refusal (A4/E4/E1); cancellation semantics with partial-stream disposition (D3); **tool-call completion semantics — partial/truncated deltas distinguishable from complete calls** (C7, rev1); retry/idempotency **below** the m-9 turn boundary with recorded attempts AND effect-safe interplay with C8's no-replay (E2/Q4, rev1); **backpressure/bounded-stream mechanics** (E6, rev1); reasoning-replay round-trip opacity (B5); usage/limit facts from the catalog, not the wire-merge (B4/F4); and **no options-bag interface** (F4). These are audit-surfaced inputs to the m-8 DESIGN + the consumer-lock reconcile — m-8 owns its contract.

## 5. Questions carried to the adversarial return + DESIGN/GRILL (none block this AUDIT)

- **Q1 — Process placement of the turn loop (grill candidate #1, hard-to-reverse):** in-conductor system subsystem (the ResummonHandler-pattern arm) vs a separate governed runtime process attached seat-side. Tension to resolve: the send path must reach the §1a governed egress + §1b credential binding (trusted-side), but tool execution must be unprivileged (C2), and a stalled provider stream must never block m-7's loop (m-7 §2.3). Both shapes satisfy this matrix; the grill decides.
- **Q2 — Session/transcript record design WITHIN the m-7 substrate** *(rev1: PARTLY CLOSED — the "m-9-owned append-only session log" option is STRUCK per blocker 2; canonical transitions commit through m-7's substrate, full stop)*: what remains designable is record schema, projection shape, and volume management (261 records/4 days was the store's design point; model turns are noisier) inside the substrate — never a second durable truth.
- **Q3 — Observe granularity over streamed output** *(rev1: floor FIXED, owner answer still required from m-3)*: the landed gate is candidate-at-submit only (no streaming observer exists). m-3 chooses the granularity, but the locked floor stands: **boundary-only observation is admissible ONLY if zero partial bytes reach any recipient or actuator first** (C9). m-9 shapes the candidate; m-3 owns the decision.
- **Q4 — Retry × final-wire-authorization (feeds the m-3 §1a amendment consumer review)** *(rev1: widened per the implementer)*: is a wire-level retry of identical bytes one authorization or N? Where does the idempotency key (m-8) sit relative to the final-authorization point? **Plus the tool-effect half: provider retry must never duplicate an authorized effect (C8 no-replay).**
- **Q5 — Compaction authorization** *(rev1: ownership split fixed)*: trigger POLICY belongs to m-5/operator (RUNTIME-RESEARCH §12 #6); m-9 owns deterministic mechanics + typed state. Open: operator-gated vs archetype-policy-gated vs free-with-record.
- **Q6 — Provider-await posture** *(rev1: constraint added — m-9 adds NO scheduler either way)*: does a turn awaiting a slow provider response hold via the `ServiceWhileBlocked`/expiry pattern, or park the lane (m-6 7-state semantics)? For the m-7/m-6 seam review. (Provider latency is seconds; park/wake was designed for human latency.)

---

## 6. rev1 fold-log (m-9.implementer adversarial return `c7-audit-m-9/AUDIT-implementer-20260714-213834`, verdict revise — every finding VERIFIED against the bytes/charter before folding; none rejected)

1. **Blocker 1 — ceiling enforcement mis-assigned to m-9.** Verified against charter bytes ("does not set or enforce it"). Folded: §1 authority-headers row, B2, C2, the still-open bucket, the design recommendation — enforcement = m-5 semantics + m-7 hosting; m-9's still-open piece = the inert request + consuming the authorization.
2. **Blocker 2 — m-9-owned canonical session log struck.** Verified against m-7 one-writer/canonical substrate + the charter no-re-own of crash-atomic commit/recovery. Folded: B1 revised, Q2 narrowed to schema/projection/volume within the substrate.
3. **Blocker 3 — C2 omitted the m-3 observation stage.** Verified against charter ("the observe-gate applies to model output") + kickoff §5 (`streamed-output observation`). Folded: the C2 rev1 sequence, new row C9 (observe-before-action as a locked floor), D3 narrowed (no partial to recipient/actuator pre-observation), A1 narrowed.
4. **Blocker 4 — G2 misread opencode as monotonic narrowing.** Planner re-verified `subagent-permissions.ts:4-26` directly: the docstring is explicit — parent restrictions govern the parent; only denies carry; the child's own permissions determine capabilities. My original "P (principle corroboration)" was WRONG. Folded: G2 flipped to R-as-donor / already-closed-principle (opencode = negative fixture for incomplete inheritance); G1's "permission-narrowed" phrase corrected.
5. **Blocker 5 — missing crash-safe tool-request/effect identity.** A genuine gap (the matrix had no row binding request identity across recovery/retry). Folded: new row C8 (stable session/turn-bound request ID · durable typed disposition incl. unknown-effect · no-replay of completed/indeterminate effects · duplicate-ID fail-closed; ownership m-9 identity/state · m-5 semantics · m-7 hosting/commit/recovery), with the opencode danger chain cited; Q4 widened with the no-replay half.
6. **Blocker 6 — E5 evidence wrong.** Planner re-verified `ai/types.ts:150-176` directly: `timeoutMs` (:153-157) is the request timeout; `maxRetryDelayMs` (:165-176) caps a server-requested retry delay. My cite was wrong. Folded: E5 evidence corrected, retry-delay capping explicitly not promoted as timeout semantics; m-7 owns timer hosting.
7. **Row adjudications applied:** A1/A2 narrowed (logical lifecycle + invariant-only; no donor coalescer as the frank shape pre-Q1); B3 narrowed (fixtures, not wholesale; policy to m-5/operator); C1 narrowed (malformed-call feedback is governed context — no ungoverned side channel); C2 citation replaced with the verified direct-executor (`native-runtime.ts:169-190`) + the in-source warning (`processor.ts:98-102`).
8. **Missing candidates added:** C7 (tool-call completion/finalization; pi truncated-refusal fixture), C8 (identity/no-replay), C9 (observe-before-action), E6 (backpressure/bounded buffering — from the kickoff §1 contract list).
9. **Count correction:** the original matrix held **36** rows (my cover relay said 30 — a miscount in the cover, corrected here); rev1 = **40** rows. §3 buckets, risks (now the implementer's hardened 7-gate set), and §5 question dispositions reworked as above.

**ACTIONS_GIT_REF:** no frank/ edits; artifact = this file (rev1) + the c7-audit-m-9 AUDIT relays + INDEX rows. cwd is not a git repo (docs workspace).
**FINAL_GIT_STATUS_SHORT:** unavailable — cwd is not a git repo (docs workspace); `frank/` untouched at `502e06c` (s11-close).
