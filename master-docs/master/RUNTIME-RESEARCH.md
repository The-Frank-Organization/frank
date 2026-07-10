# RUNTIME-RESEARCH.md — How the conductor rides existing runtimes, provider-agnostically

**Status: RESEARCH SYNTHESIS — report-only. NOT a design lock, NOT a PLAN, NOT a spike.**
This is the Step-0 evidence base for *how* the conductor concretely attaches to and drives existing agent
runtimes (Claude Code, Codex, and the wider field) in a **model/provider-agnostic** way, with **persistent
seats** as the default. It feeds three open items and changes none of them by itself:
- **operator decision ①** (wrap-vs-attach + the honest security claim) — see §10;
- **readiness register Cluster 2** (ride-vs-own honesty) and **Cluster 3** (the first spike) — see §11;
- **Step-1 conductor-core** and **Step-3 model-agnostic runtime** planning (`ROADMAP.md:57`, `:74`).

Docs-only, in cwd. No locked design (`ARCHITECTURE.md` §1–§C3) is modified here. Where research and a locked
doc disagree, that disagreement is logged for a *bounded VP-gated fix*, not silently reconciled here.

**Cross-checked (§13)** against an operator-provided **external chat-model research brief**
(`RUNTIME-RESEARCH-GPT.md`) — an independent method that *converges* on the architecture; its one divergence
(model tiering) resolves into the tuple (§6.5). That brief is **evidence, not a seat** (no authority).

**Then primary-source-verified (§14)** — the two load-bearing claims the brief did not corroborate were checked
by a four-lens read-only pass (one adversarial). Outcome: **`srt` is real and its sole-egress *path* is confirmed
kernel-enforced by construction**, but **"sole external sender *by construction*" is refuted-as-assumed → now a
spike-gated milestone** (§14.2), and **Codex's `session_source` identity claim is refuted → identity is
conductor-owned** (§14.1-I). The security posture (§8) and decision-① framing (§10) are the verified three-rung
ladder.

**Then VP-reviewed (`runtime-research/RECONCILE-orchestrator-reviewer-20260630-213911.md`): _approve with
tightened spike gates._** The VP accepted §14 and the three-rung ladder, and (a) **expanded the by-construction
spike to two stages** — properties (1)–(5) earn only "single mediated network path"; a **6th property** (every
outbound broker request passes the conductor egress/content gate) is required for "destination + content control
by construction" (§14.2, §8, §12); (b) routed the **identity-conductor-owned** correction to a bounded m-1 fix
(§11); (c) flagged the **model-family tiers as drift-prone snapshots**, not durable architecture (§6.5).
**Decision ① is now cleared for the operator to make on verified, VP-reviewed ground** — using the three-rung
ladder (§10). Nothing is locked until the operator records it.

**Provenance:** a 3-wave read-only research fleet (≈21 subagents, all read-only web/reference lenses — the only
sanctioned subagent spawn per the charter). Roster + per-agent provenance in the Appendix. Every load-bearing
claim is tagged **[P]** primary (official docs / source / spec / first-party report), **[C]** corroborated
across ≥2 independent sources, or **[V]** vendor/SEO material treated as a claim to verify. Dates are access
dates (fleet ran 2026-06-30).

---

## 0. BLUF — the five load-bearing conclusions

1. **Attach, don't rebuild — and the wake primitive already exists.** A persistent agent process (Claude Code
   `claude -p --input-format stream-json`, Codex `app-server`) holds its stdin pipe open and sits
   **kernel-suspended at zero CPU between turns**, waking on a single delivered line. **The OS pipe *is* the
   inbox and the wake mechanism** — no polling loop, no timer. This is the technical foundation that makes the
   operator's "persistent seats, same agents all sprint" preference not just possible but *cheaper* than
   ephemeral-per-relay. [P][C]

2. **Model-agnosticism is inherited from the runtime layer, not built at the model layer.** Every serious
   runtime already abstracts providers (LiteLLM / OpenAI-compat / Anthropic-compat). The conductor integrates
   **runtimes**, not raw model APIs; the model-adapter surface it must own directly is tiny. Ride the runtime →
   inherit ~100 models for free. [P][C]

3. **Channel-stamped FROM must be *built by the conductor*, not inherited from the runtime.** The conductor runs
   as an **MCP server**; each seat reaches it over a **dedicated per-seat transport** (its own stdio pipe or its
   own bearer token), and the courier stamps FROM from *which channel the bytes arrived on*, never a
   client-supplied field. **Verification correction (§14):** the runtime does **not** supply this for free —
   Codex's `session_source` is a mutable launch-arg *telemetry* label, not a forgery-robust author stamp — so the
   conductor owns identity itself and treats the runtime as a dumb executor. This holds **by correct per-seat
   isolation** (one seat ⇔ one channel; tokens neither shared nor inherited), which a spike must prove — *not*
   "by construction." Persistence does not weaken it. [P]

4. **The honest security claim is a three-rung ladder, and "by construction" is *spike-gated* (verified §14).**
   Attach (Step-1) → **"confusion-resistant."** `srt`-wrap → **"sandboxed defense-in-depth"** — verification
   confirmed the jail's only egress *path* is a parent-owned socket, kernel-enforced. But it **refuted** "sole
   external sender / unforgeable-FROM *by construction*" as automatic: the broker is a foolable policy point (a
   null-byte allowlist bypass shipped ~5.5 months in the wild), the jail **fails open by default** on common
   kernels, and identity must be conductor-owned (Codex's `session_source` is *not* it). So the strong claim
   ships **only after a code spike proves five properties** (§14.2) — never on assumption. (§8, §10)

5. **The field is provider-agnostic by consensus — MCP is table-stakes and the model zoo is reachable — but
   a lane must be *probed*, not read off a leaderboard.** OpenAI, Google, Microsoft, and AWS all ratified MCP
   in 2025; DeepSeek, Moonshot (Kimi), Zhipu (GLM), MiniMax, and Alibaba (Qwen) all ship **first-party
   OpenAI-*and*-Anthropic-compatible endpoints**; local serving exposes the same surface. The routable unit is
   **model × provider × serving** — not the model name. And because coding scores mis-rank *agentic* reliability
   by **40+ points** (and much of the 2026 leaderboard web is fabricated), the conductor admits a lane to
   autonomy only on a **qualification probe** — which is just our m-3 observe-gate pointed at lane admission.
   **Our governance already handles model-agnosticism; it doesn't need new machinery.** (§6.5)

---

## 1. The question, and why it matters now

The operator's ask: *"research how the conductor concretely rides Claude Code / Codex,"* under two hard
constraints — **(a) model/provider-agnostic** ("most mainstream models, not just OpenAI and Anthropic") and
**(b) persistent seats strongly preferred** ("the same agents throughout an entire sprint … like a guy going
through each tab relaying the relays … some Zellij-multiplayer type stuff, but there should be a cleaner way").

This matters *now* because the build-readiness review (`READINESS-REGISTER.md`) surfaced that the c1–c3
design-of-record proved its strongest identity/egress guarantees **for a conductor that owns the process
tree**, while **Step-1 rides rented runtimes** (`ROADMAP.md:19` "Ride, then replace"). Decision ① is exactly
the question of how honest to be about that gap and how to close it. This research is the evidence the operator
asked for to make that call — *measure twice*.

The destination (`ROADMAP.md:9`) is a standalone TUI = **agent runtime × Zellij-style multiplexer × email
client**, conductor underneath. So the research also answers: what do we *ride* in Steps 1–2, and what do we
*replace* in Steps 3–4 — and the operator's "cleaner way than a Zellij multiplayer session" is answered
directly in §2–§3.

---

## 2. The wake mechanism — the OS pipe *is* the primitive (no polling, no timer)

The operator's sharpest worry: *"isn't there an issue with agents not being able to be automatically prompted
to check their inbox without a constant loop or a timer? Or can messages be queued to them?"*

**Answer: they are queued, by the kernel, for free.** This is the single most important mechanical finding.

- **Claude Code** run as `claude -p --input-format stream-json --output-format stream-json` keeps **stdin open**
  and consumes newline-delimited JSON (NDJSON) messages. Between turns the process is **blocked on a pipe read**
  — kernel-suspended, **zero CPU**, not spinning. When the conductor writes one line to that pipe, the kernel
  wakes the process and it runs the next turn. The pipe is a FIFO queue: multiple messages written while the
  seat is busy simply queue and are consumed in order. **No inbox-polling. No cron. No timer.** [P] (Claude
  Code headless / streaming-input docs; corroborated by the SDK's `stream-json` input mode.)

- **Codex** exposes **`codex app-server`** — a persistent, addressable process speaking **JSON-RPC** over
  stdio / a Unix socket / WebSocket. Seats are long-lived **ThreadId**s you send turns to. Same property:
  suspended between turns, woken by a delivered request. [P][C]

- **Cross-runtime:** Gemini CLI added **ACP** (Agent Client Protocol, `--acp`), a runtime-agnostic
  session/turn protocol that **Claude Code and Codex also speak**. ACP gives a *single* persistent-seat drive
  surface across three major runtimes. [P][C]

**Consequence for the "cleaner way than Zellij":** the operator's mental model — *"a guy going through each tab
relaying the relays"* — is exactly right, and the clean version is **not** a terminal multiplexer at all. It is:
**the conductor holds one open pipe per seat and writes the relay path (or the relay itself) onto the target
seat's pipe.** That *is* "sending the relay to the targeted session," minus the human, minus the tabs, minus
Zellij. Zellij/TUI is then a *view* over these pipes (Step-4 product surface), not the transport. The transport
is N kernel FIFOs the conductor owns both ends of. This is the "cleaner way" the operator was reaching for.

> **Design implication (feeds Step-1/Step-3, not a lock):** the conductor's courier is a process that owns, per
> seat, one long-lived bidirectional pipe. Delivering a relay = one `write()`. Waking a parked seat = one
> `write()`. The seat never "checks" anything.

---

## 3. Persistent seats — the operator's preference is the right default (with an honest caveat)

Frameworks *default* to **ephemeral turn + durable store** (spawn a worker per unit of work; state lives in the
store, not the process). That default exists for cloud-scale fan-out where workers are cattle. **Our workload is
the opposite:** a small, fixed team of reasoning seats that accrue context across a sprint — the exact shape of
*this* operation. For that, **persistent seats are the better fit**, and the operator's instinct is correct. [C]

What the research adds (the honest caveat and how to handle it):

- **Persistence does not weaken identity.** Identity = the connection (the pipe / the token), which is *more*
  stable under persistence, not less. A persistent seat = a persistent channel = a persistent stamped FROM.
  (§7) [P]

- **The real cost of persistence is context growth**, not identity or wake. A seat that runs an entire sprint
  accumulates tokens. Two mitigations, both GA today:
  1. **Auto-compaction** (what *this* session just did) — the runtime summarizes and continues. Claude Code and
     Codex both expose usage signals (`get_context_usage` / equivalents) so the conductor can *see* pressure
     coming and trigger a checkpoint. [P]
  2. **Checkpoint-resume** (`claude --resume <session>`, Codex thread reload) — a seat can be *parked to the
     store and re-hydrated*, which is also exactly the m-6 park/wake mechanism (§9). Persistence and durability
     are **not** opposed: the store is the seat's savepoint, the pipe is its liveness. [P][C]

- **Hybrid, not dogma.** The right model is **persistent by default, re-hydratable on demand**: standing seats
  (the six domain pairs, the orchestrators) stay live on their pipes across a cycle; short-lived expansion-slot
  archetypes (m-5 sensors, `/btw` side-questions per `ROADMAP.md:37`) are ephemeral forks that never needed
  persistence anyway. This maps 1:1 onto the m-5 archetype model already locked in §C3.

> **This directly answers the operator's "though this may be too much persistence":** it isn't, for the standing
> team — but the archetype layer already gives us the ephemeral escape hatch where it's actually wanted.

---

## 4. Integrate runtimes, not models — where model-agnosticism actually lives

The **critical constraint** is provider-agnosticism. The key finding is *where* to spend the agnosticism budget:

- **Runtimes already abstract providers.** Claude Code, Codex, Goose, opencode, OpenHands, Cline, Crush,
  Continue each sit on a provider-abstraction layer (LiteLLM, an OpenAI-compat client, or a native adapter set).
  Goose's model-adapter surface is famously **~one method**. So **model-agnosticism is inherited by riding a
  runtime** — the conductor does not re-implement 40 provider clients. [P][C]

- **The threat model *inverts* when the conductor is the MCP server** (not an MCP client): the entire
  tool-poisoning / malicious-server class evaporates, because the conductor is the trusted server the seats
  connect *to*, and the seats are the (sandboxed) clients. (§8) [P]

- **Two integration depths, matching the roadmap:**
  - **Step-1/2 (ride):** attach to a host runtime; inherit its model zoo; the conductor governs via MCP + the
    stdin pipe. Agnosticism is *free* — whatever the host can call, we can route to.
  - **Step-3 (own):** the conductor drives models directly through provider adapters (`ROADMAP.md:74`). Here the
    agnosticism surface is *ours*, and the finding is: **build it on the OpenAI-compat + Anthropic-compat
    baseline** that ~every vendor already speaks (§6), with a thin adapter only where a vendor deviates. This is
    a small, well-bounded surface — not a 40-provider matrix.

> **Portable-core / runtime-adapter split (the shape Step-3 should target):**
> - **Portable core (ours, model-agnostic):** the relay store (sole-writer, append-only), the seat-stamper /
>   channel-stamped FROM, the form/lint gate, the observe-as-send gate, the router + routing record, the
>   park/wake scheduler, the egress gate. *This is the entire c1–c3 design-of-record — none of it is
>   model-specific.*
> - **Runtime adapter (per host, thin):** {spawn a persistent seat · deliver a turn on its pipe · read its
>   result/usage · steer/interrupt at a step boundary · park/resume}. One small adapter per rideable runtime.
> - **Model adapter (per provider, thinner):** only where a provider isn't OpenAI-/Anthropic-compatible. Most
>   are, so this is mostly an **alias/capability table**, not code (§6).

The agnosticism budget is spent almost entirely in *config* (alias tables, capability priors) and a *thin*
adapter layer — because the runtimes and the vendors already did the hard part.

---

## 5. Runtime coverage matrix — how the conductor rides each host

How the conductor attaches to, drives, and governs each candidate runtime. "Persistent-seat drive" = can we
hold a long-lived addressable seat and wake it on a pipe/RPC. "Governs via" = how the conductor injects the
courier + gate. "Steer/interrupt" = can the human redirect a running lane (`ROADMAP.md:32`).

| Runtime | Persistent-seat drive | Model-agnostic? | MCP client | Governs via | Steer / interrupt | Ride verdict |
|---|---|---|---|---|---|---|
| **Claude Code** | ✅ SDK **streaming-input** = open stdin pipe (non-polling) **+** `--resume`; ACP-capable. *(NB: the separate "Remote Control" feature polls the Anthropic API — ride the SDK path, not that.)* | ✅ Anthropic-compat + Anthropic-compatible gateways *(arbitrary OpenAI-compat backends unverified)* | ✅ stdio/SSE/HTTP | MCP server + stdin NDJSON | ✅ both (host-native) | **Primary ride (Step-1).** Best-documented persistent+resume+steer story. |
| **Codex** | ✅ `codex app-server` (JSON-RPC over stdio/unix-socket/exp-ws; addressable resumable `thread` seats) | ⚠️ custom base-URL but **Responses-API-gated** (`wire_api="chat"` removed; non-OpenAI needs a Responses gateway) | ✅ stdio + streamable-HTTP | MCP server + app-server RPC | ✅ (host-native) | **Co-primary ride for *drive*.** But identity is **conductor-owned** — `session_source` is a telemetry label, **NOT** channel-stamped FROM (§14, verified). |
| **Gemini CLI** | ✅ **ACP** (`--acp`, JSON-RPC/stdio) — the cross-runtime seam | ⚠️ base-URL is Gemini/Vertex *proxying*, **not** proven multi-provider | ✅ stdio/SSE/HTTP | ACP + MCP | ⚠️ varies | **Ride via ACP** (as a Gemini/Vertex shell). ACP unifies CC/Codex/Gemini. |
| **Goose** (Block) | ✅ long-running sessions | ✅ **~1-method model adapter**; LiteLLM-style | ✅ | MCP server | ⚠️ | **Strong ride.** Thinnest provider surface = reference for our Step-3 adapter. |
| **OpenCode** | ✅✅ **`opencode serve`** — HTTP server, sessions/messages as first-class resources + SSE `/event`; ACP too | ✅ many providers, custom base-URL | ✅ | HTTP session API + MCP server | ⚠️ | **First-class HTTP-server ride** (cleanest externally-addressable seat). Needed a `reasoning_content` patch for thinking-models (§6). |
| **OpenHands** | ✅✅ **Agent Server** — HTTP+WebSocket, long-running convos; per-request identity via `X-Session-API-Key` | ✅ any LiteLLM model + OpenAI-compat proxies | ⚠️ ACP yes; Agent-Server-as-MCP-client unverified | HTTP+WS API (OpenAI-compat front door keeps `X-OpenHands-ServerConversation-ID`) | ⚠️ | **First-class HTTP-server ride** (strong backend; heavier). |
| **Cline / Roo** | ⚠️ editor-embedded | ✅ many providers incl. custom base-URL | ✅ | MCP server | ⚠️ | Ride-viable; preserves `reasoning_content` (good for thinking models). |
| **Crush / Continue** | ⚠️ | ✅ | ✅ | MCP server | ⚠️ | Secondary rides. |
| **aider** | ⚠️ (CLI, not a server) | ✅ LiteLLM (best leaderboard coverage) | ❌ **no MCP client** | — (no MCP) | ❌ | **The exception.** Best for *benchmarking* model×provider lanes (m-4), not for riding as a governed seat. |

**Reading of the matrix:**
- **Step-1 rides Claude Code + Codex** (both give persistent addressable seats, MCP, and host-native
  steer/interrupt). **ACP is the strategic seam** — build the persistent-seat drive against ACP and CC/Codex/
  Gemini are one adapter.
- **Two transport classes (from the external brief, §13.3):** *HTTP-server runtimes* (**OpenCode `opencode serve`**,
  **OpenHands Agent Server**) give the cleanest externally-addressable persistent seats, with per-request
  credential identity (`X-Session-API-Key`) — arguably a *cleaner* seat surface than the stdin pipe; *ACP-subprocess
  seats* (Goose, Gemini CLI, Codex app-server) are lighter but the conductor owns lifecycle + crash recovery. Ride
  either — HTTP-server is cleaner for orchestration, ACP is the cross-runtime unifier.
- **Goose is the reference for the Step-3 own-the-runtime adapter** (thinnest provider surface).
- **aider is a benchmarking instrument, not a ride target** (no MCP client) — useful to m-4's lane-scoring, not
  to the courier.
- **`reasoning_content` handling is a runtime-selection criterion**, because thinking-model lanes break on
  runtimes that drop it (§6).

---

## 6. Model coverage matrix — which models are lane-viable, at what tier, via which path

> **Fold status: COMPLETE.** Populated from the full model-layer wave — `m-open-hosted` (§6.1),
> `m-deepseek-kimi` (§6.2–6.3), `m-glm-minimax-qwen` (§6.3), `m-local-serving` (§6.4), and the honest tier-list
> `m-agentic-reality` (§6.5–6.6). The two model-layer analyses (hosted §6.5 and local §6.4) **independently
> converged** on the same architecture: the routed unit is a serving *tuple*, and the **m-3 observe-gate is the
> mechanism that makes model-agnosticism governable.**

### 6.1 The routable unit is `model × provider × serving`, not the model name

The wave's hardest finding (`m-open-hosted`): **tool-call reliability is a property of the *serving stack*, not
the weights.** The K2-Vendor-Verifier result — identical open weights scoring **72% → 100%** tool-call validity
depending only on *which provider served them* — means the conductor must route to a **`model × provider ×
serving-endpoint`** tuple carrying a **live reliability signal**, and must *re-benchmark on every provider or
model-ID bump*. This is why m-4's routing record is `model × provider × region`, not `model`. [P][C]
- **An external chat-model research brief (`RUNTIME-RESEARCH-GPT.md`, operator-provided evidence — not a seat)
  independently adds an explicit 4th axis: *compatibility mode*.** The full unit is
  **`model × provider × serving × compat-mode`**, because OpenAI's own `/responses` vs `/chat-completions` split
  changes wire protocol, tool-schema enforcement, streaming shape, and reasoning-summary behavior for the *same*
  model on the *same* provider (Continue documents forcing `useResponsesApi:false` to survive "organization must
  be verified" reasoning-summary errors [P]). Adopt the 4-axis unit — m-4's record should carry the compat-mode. [P]
- **OpenRouter is the best single front door** for hosted-model breadth: set `require_parameters: true` (drop
  providers that silently ignore `tools`) and use its Auto/quality routing; it normalizes ~one API over the zoo.
  For *governed* lanes, prefer a **named provider** (Fireworks / Together / DeepInfra / Baseten / NVIDIA NIM)
  over "cheapest wins" auto-routing, so region and reliability are pinned. [P][C]

### 6.2 The cross-cutting gotcha every courier must absorb: `reasoning_content` on thinking models

The wave's single most decision-relevant *integration* fact (`m-deepseek-kimi`, corroborated by opencode/
instructor issues): **the "thinking"/reasoning variants are the tool-calling *liability*, not the asset.** They
emit chain-of-thought in a `reasoning_content` field the OpenAI SDK doesn't surface; on **multi-turn tool-call
replay** (session reopened, history resent) strict servers **reject the request** (Moonshot: *"thinking is
enabled but reasoning_content is missing in assistant tool call message at index X"*). [P][C]
- **Mitigation (a routing rule, not just a code patch):** drive the **non-thinking / instruct** variant in
  tool loops (where these families' tool-calling is *tuned* anyway), **or** ride a runtime that preserves
  `reasoning_content` (vendor Anthropic-compat endpoints, Cline, patched opencode).
- **Spike caveat that inverts the naive intuition:** "reasoner is better for agents" is **false** here. Test any
  lane with a **≥8-step loop that closes and reopens the session** — a 2-step smoke test won't surface the
  replay rejection.

### 6.3 The hosted non-frontier families — all dual-stack, all drop-in

Every one of these ships **first-party OpenAI-compatible AND Anthropic-compatible** endpoints, so the conductor
reaches them through Claude Code / Cline / Roo / opencode / a raw loop with near-zero glue (`m-glm-minimax-qwen`,
`m-deepseek-kimi`):

| Family | Best coding/agent SKU | Dual-stack? | Lane strength | Governance flags |
|---|---|---|---|---|
| **DeepSeek** (V4-flash/pro; Terminus lineage) | `deepseek-v4-flash` **non-thinking** | ✅ OpenAI + Anthropic (`/anthropic`) first-party | **Strongest default lane in the cluster** — cheapest ($0.14/$0.28·M), ~74% Aider Polyglot / ~73% SWE-Verified, 1M ctx, tool-calling-*tuned* non-thinking line | PRC (Hangzhou) first-party → ride Western reseller for residency; aliases `deepseek-chat/reasoner` **deprecate 2026-07-24**, pin real IDs; Anthropic endpoint drops MCP-native blocks (runtime must execute MCP) |
| **Kimi** (Moonshot K2 → K2.5–2.7) | `kimi-for-coding` / K2-Instruct **non-thinking** | ✅ OpenAI + Anthropic + dedicated coding endpoint | **Strongest long-horizon *agentic* lane** — stable through **200–300 tool calls** | PRC (Beijing); no `tool_choice:"required"`; `temperature∈[0,1]`; self-host needs a **native K2 tool parser** (vLLM/SGLang/etc.); K2.6/2.7 >80% SWE claims are [V] |
| **GLM** (Zhipu, 4.6+) | GLM-4.6 coding SKU | ✅ OpenAI + Anthropic | **Most *proven* of the three in coding harnesses**; the common "just works in Claude Code" pick | PRC; ride coding SKU not general/reasoning; region = lane attribute |
| **MiniMax** (M2+) | MiniMax-M2 | ✅ OpenAI + Anthropic | Solid coding lane | PRC; ride coding SKU; model-ID churn |
| **Qwen** (Alibaba, Qwen3-Coder) | Qwen3-Coder variants | ✅ OpenAI + Anthropic; **also self-hostable** | Strong coding; the bridge to *local* (§6.4) | Self-host needs Qwen tool parser; region/hosting is a lane choice |

**Load-bearing routing rules that fall out (feed m-4):**
1. **Ride the coding-specialist SKU**, never the general/reasoning SKU, for lane work.
2. **Drive the non-thinking variant** in tool loops (§6.2).
3. **Treat model IDs as versioned** — alias table + re-benchmark on every bump (aligns with the "version schemas
   / migration from day one" stance in the public-release-intent design note).
4. **Region is a lane attribute** — PRC-first-party vs Western-reseller is a *governance* choice the router
   records, not a detail.

### 6.4 Local / self-hosted serving as governed lanes (`m-local-serving`)

**The framing correction:** every local runtime exposes a tool-capable OpenAI `/v1/chat/completions` — so "does
it have an API" is *never* the question. The load-bearing question is **"can the endpoint reliably round-trip
structured tool calls across a *multi-turn* agent loop?"** That, plus hardware endurance, decides first-class vs.
constrained — and the answer is a property of the **full serving tuple**, not the model.

**Per-runtime verdict (integration tax is in reliability-config + headless-ness, not the wire format):**

| Runtime | Tool-capable `/v1` | Native MCP host | Integration tax / verdict |
|---|---|---|---|
| **vLLM** | ✅ (+`/v1/responses`) | ✅ | Set `--enable-auto-tool-choice` + **model-matched `--tool-call-parser`**; GPU-only; **lowest tax for a headless production lane.** |
| **SGLang** | ✅ (+Anthropic `/v1/messages`) | ❌ | Same parser-matching; Anthropic path newer — **crashes on 2nd tool call of a multi-tool turn** [P]. |
| **llama.cpp** (`llama-server`) | ✅ | ❌ | `--jinja` + tools template; **GBNF grammar can constrain tool-arg JSON to *valid*** (unique strength); most portable (CPU/Metal/CUDA). |
| **Ollama** | ⚠️ shim, no `tool_choice` | ❌ (bridge) | **`/v1` silently drops streamed tool-calls** — must use native `/api/chat` (a 2nd protocol) or `stream:false`. |
| **LM Studio** | ✅ (`:1234/v1`, +Anthropic) | ✅ | Lowest glue **but a desktop app** (GUI-first; has `lms` headless) — best *supervised/HITL*, awkward for fleet automation. |
| **MLX** (`mlx_lm.server`) | ⚠️ "basic," **self-labeled not-for-production** | ❌ | Immature/model-dependent parsing → **use a 3rd-party MLX server** (mlx-openai-server, vllm-mlx) for real tool-use. |
| **TGI** | — | — | **Excluded** — maintenance mode since 2025-12-11, repo archived read-only 2026-03-21; HF says migrate to vLLM/SGLang [P]. |

**Tool-calling reality — two tiers the conductor must distinguish [P]:**
- **Native-parser tier (real):** tool-trained model + matching `--tool-call-parser` + the model's official chat
  template → calls parsed from the trained format into a structured `tool_calls` array. Reliable **only when the
  tuple matches exactly.** vLLM/SGLang carry the broadest roster (`qwen3_coder`, `kimi_k2`, `glm45`,
  `deepseek_v3`, `openai`/gpt-oss).
- **Prompt-templated shim tier (faked):** server injects a synthetic "please emit `[TOOL_REQUEST]…`" prompt and
  regex-scrapes it (LM Studio "default tier," llama.cpp "Generic" handler). **Breaks on multi-turn replay** —
  general effect measured at **~25-pt drop** when a task is revealed over turns vs. all at once (arXiv 2505.06120).
- **Where the native-parser tax bites hardest:** **Qwen3-Coder emits XML** (`<function=…>`), *not* Hermes JSON —
  point a `hermes` parser at it and the call returns as plain `content`, loop stalls; must set `qwen3_coder`.
  **Kimi K2** is the cautionary tale — vLLM needed a full engineering post to get its tool-call-ID format right,
  and it needs `-tp 8` (**datacenter-only**).
- **Cross-cutting fragility — streaming tool-calls are broken/flaky on *every* runtime right now** (vLLM, SGLang,
  Ollama, llama.cpp all have open issues; dozens of per-model streaming parsers assuming single-token deltas).
  **Robust mitigation: run local tool-turns non-streamed.** This is a hard courier rule, not a preference.

**Which local models are agent-loop-viable, by hardware tier:**

| Tier | Model @ quant | Fits on | Verdict |
|---|---|---|---|
| **Entry 24 GB** | **Qwen3-Coder-30B-A3B** Q4 | 1× 3090/4090; 32 GB Mac | **Best consumer pick** — MoE ~3B active → fast; native agentic FC; XML parser required. |
| **Entry 24 GB** | **Devstral-Small-24B** Q4 | 1× 4090; 32 GB Mac | **Viable** — purpose-built agentic coder, 53.6% SWE-Verified; mature Mistral parser. |
| **Mid 48 GB** | Llama-3.3-70B / Qwen2.5-72B Q4 | 2× 3090 | **Marginal** — fussy JSON; MoE picks are faster + less fussy. |
| **Serious 80 GB** | **gpt-oss-120b** MXFP4 | 1× H100 | **Viable-but-taxed** — strong, but harmony format buggy to self-host → non-streaming only. |
| **Serious 80–160 GB** | **GLM-4.5-Air 106B-A12B** FP8 | 2× H100 / 1× H200 | **Best open tool-use reliability tier** (GLM-4.5: 70.1% τ-bench, 64.2% SWE-Verified). |
| **Datacenter-only** | Kimi K2 (~1T, tp≥8), full DeepSeek-V3/R1 | 8×80 GB+ | Not a local lane in any normal sense. |

- **Demo-only / avoid for loops:** DeepSeek-R1-Distill (reasoning distills, not agent models — calls land in
  `content`, `</think>` breaks parsers); gpt-oss-20b (too weak, 3.4% TB2.0); Codestral (FIM, not agentic);
  Qwen3 dense "thinking" (interleaved reasoning complicates loops).
- **The KV-cache endurance trap:** a 70B model adds ~2–4 GB KV at 32K but **~40 GB at 128K** *on top of weights*
  — long autonomous trajectories OOM mid-loop unless VRAM is sized for weights **+** KV. (This is the concrete
  cost of the §6.5 rule that quantized/consumer setups hold only ~4–8K *effective* context.)

**Governance ruling — first-class-but-capability-tiered, NOT a mere escape-hatch:**
- **Upside = the strongest governance story the conductor can tell:** full data residency, air-gap-capable, zero
  third-party data-processor egress, complete auditability. **For regulated / classified / IP-sensitive work the
  frontier API is off the table, so the local lane is the *only* lane** — its governance value dominates every
  reliability cost. This is the clean answer to §6.6's PRC-residency hard flag: *self-host clears it.*
- **The cost is reliability + endurance, not capability** — the proven picks clear the capability bar; they don't
  cheaply clear parser fragility, single-stream throughput, or the KV ceiling.
- **The ruling:** residency-constrained context → **first-class primary lane** (governed identically — same
  channel-stamped FROM, same relay discipline, same observe-gates). General context with a frontier lane
  available → **capability-tiered lane**: same governance, but routed with **lower capability priors + a lower
  authority-ceiling-at-spawn (m-5) + tighter observe-gates (m-3) + mandatory non-streamed tool-turns** — viable
  for **bounded supervised-autonomous** work today, *not* set-and-forget long-horizon autonomy.

**The single most decision-relevant local finding — and it is the *same* finding as §6.5:** treat a local
endpoint as a **first-class governed lane, but tier it by an *observed tool-call-validity gate*, not by trusting
the endpoint.** The gating constraint is tool-call parsing reliability, a property of the **whole
`{model + server + parser + chat-template + streamed?}` tuple** — a parser/model mismatch or a streamed tool-turn
**silently converts a valid tool call into plain text and the loop stalls with *no error surfaced*, which is
exactly what the observe-gate exists to catch.** Pin the tuple, run tool-turns non-streamed, gate on observed
structured-tool-call validity every turn — and Qwen3-Coder-30B-A3B or Devstral on a single 24 GB card is a
production-viable supervised-autonomous lane today. **Note the convergence:** the hosted tier-list (§6.5) and the
local analysis (§6.4) *independently* landed on the same architecture — the routed unit is a tuple, and the
**m-3 observe-gate is the mechanism that makes model-agnosticism governable.**

### 6.5 The honest capability tier-list (`m-agentic-reality`)

The anchor for the whole model layer. Tiers are anchored to **neutral-agentic** data (Terminal-Bench 2.0 /
τ-airline class), at **first-party full-precision serving only** — because the tier *is* a property of the
`model × provider × serving` triple, not the model name. Successors past the Jan-2026 cutoff are shown as
trajectory, unverified on neutral harnesses.

> **⚠️ VP guardrail (`…-213911`, Finding 4): the specific model-family tiers below are drift-prone SNAPSHOTS, NOT
> durable architecture.** Do not lock current GLM/Kimi/Qwen/DeepSeek/… rankings into any design-of-record. The
> **durable rule is the tuple + canary policy** (§6.1, §6.5-probe); a model's tier belongs in a **fresh
> lane-qualification record produced when the lane is actually brought up**, and re-run on every provider/ID
> change — never read from static research prose.

**T-A — genuinely viable AUTONOMOUS agent-loop lanes** *(floor: well-formed multi-step tool calls, holds
structured-edit format, ≥~40–50% neutral-agentic, ≥32K effective context, recovers from a failed step)*
- **Frontier closed** (Claude Opus / GPT-5.x / Gemini 3.x, first-party) — the reference bar (~80–85% TB2.0 [P]),
  not "open."
- **GLM-4.6 / GLM-5 via Zhipu API** — **best all-round open agentic lane** (GLM-5 **52.4%** TB2.0 [P]); the
  community "drop-in for Claude Code" pick, matching `m-glm-minimax-qwen`'s "most proven" read.
- **Kimi K2 / K2-Thinking via Moonshot API** — RL-built for tool-use; aider **59.1%**, TB2.0 (K2.5) **43.2%** [P].
  Confirms `m-deepseek-kimi`'s "strongest endurance lane."
- **DeepSeek-V3.2 / R1-0528 via DeepSeek API** — top open *coding* (aider **74.2%** [P]) but only **borderline
  autonomy** (TB2.0 **39.6%**, τ²-airline 63.8 [P]); heaviest integration tax (§6.2) + hard residency flag (§6.6).

**T-B — constrained / supervised-only (strong executor *behind* a stronger planner; NOT autonomous)**
- **Qwen3-Coder-480B via Alibaba API** — SOTA open *coding* (~67% SWE [V]) but **craters to 23.9% TB2.0** [P] —
  the sharpest coding≠agentic gap in the data.
- **gpt-oss-120b** — good one-shot code but **#66/124 agentic tool-use** [P]; chronically verbose, ignores "be
  concise."
- **Devstral-2 / Devstral-Small-24B, MiniMax-M2.x** — SWE-tuned executors competent *in their own harness*,
  unproven cross-harness.
- **30–35B MoE (Qwen3.x-A3B class) at full precision** — "sufficient for a lot of tasks," weak file-selection
  judgment [C, Raschka].

**T-C — not agent-loop-ready**
- **Llama 3.x/4** (Maverick 15.6% aider [P]), **Codestral / Mistral-small** (11.1% [P]), **Qwen2.5-Coder-32B &
  below** (16.4% whole → **8.0% diff** = structured-edit collapse [P]), **Gemma-3**, **anything ≤~14B**.
- **Any T-A/T-B model quantized ≤4-bit drops a full tier** — quantization "loses fidelity exactly where code
  needs it" [C]; consumer setups hold only ~4–8K *effective* context despite advertised 128K+. (This is the hinge
  for §6.4 local.)

> **⚖️ Cross-check against an external chat-model research brief (`RUNTIME-RESEARCH-GPT.md`, operator-provided —
> NOT a seat; carries no authority, weighed as evidence) — the one substantive divergence, and it resolves *into*
> the tuple.** That brief put **GLM / Kimi / MiniMax / DeepSeek all in *supervised-only***, where this fleet puts
> GLM-5 / Kimi-K2.5 in **T-A**. The two are not in conflict — they scored **different axes of the same tuple**:
> this fleet scored *capability* on neutral-agentic benchmarks (first-party, full-precision); the brief scored
> *serving-reliability* on **generic OpenAI-compat-proxy issue evidence** (Crush "Kimi K2 using OpenAI-compatible
> API cannot perform multi-turn conversations" [P]; Cline `<think>`-tag breakage [P]; compat-mode multi-turn
> failures). Both are true. The resolution — which *validates* the tuple both efforts independently reached:
> **capability-tier is necessary but not sufficient for autonomy; a lane is admitted to autonomy only when its
> specific `model × provider × serving × compat-mode` tuple passes the canary.** So GLM-5 / Kimi-K2.5 are
> **autonomy-*capable*, autonomy-*admitted* only on a canary-qualified lane** (first-party or qualified reseller ·
> non-thinking variant · probe-passed) — and **supervised-only on a blind generic OpenAI-compat proxy.** This
> fleet slightly overclaimed by implying capability⇒autonomy; the brief slightly underclaimed by treating a
> fragile serving path as the model's ceiling. (Note: the brief's *own* cited evidence — the OpenHands Index
> listing Kimi-K2.5/2.6, GLM-5/5.1, MiniMax-M3 as strong candidates — supports the capability read, confirming
> the supervised-only call was serving-reliability-driven, not capability-driven.) See §13 for the full
> cross-check. **This fold is evidence only — the real VP (`master.orchestrator-reviewer`) has not yet reviewed
> it.**

**The reality gap — why leaderboards mislead (route on the right benchmark):**
- **Coding score ≠ agentic reliability — they disagree by 40+ points** on the *same* neutral harness (GLM-5 52.4
  > Kimi-K2.5 43.2 > DeepSeek-V3.2 39.6 > **Qwen3-Coder-480B 23.9**) [P]. **Route on Terminal-Bench / τ-airline,
  never SWE-bench / aider / HumanEval.**
- **Vendor SWE-bench is harness-inflated** — Scale's neutral SWE-bench-Pro put Qwen3-Coder-480B at **38.7%** vs
  the ~67% vendors cite [P]. "100% tool-call success" is a curated-demo artifact.
- **Much of the 2026 leaderboard web is fabricated** — invented model names, mangled rows from SEO aggregators
  [V]. **A router that auto-ingests scraped leaderboards ingests hallucinations** — capability priors must come
  from neutral primaries or our own probe, never a scrape.

**The single most decision-relevant finding — and it maps onto our own architecture:**
**Autonomous-lane viability is a property of the `model × provider × serving` triple, established by PROBE, not
read from a leaderboard.** The conductor should gate every lane behind a **cheap standing qualification probe** —
a canary exercising (a) a well-formed multi-step tool call, (b) structured-edit adherence, (c) forced
failed-step recovery, (d) a `reasoning_content` / `tool_choice` round-trip — **admit a lane to T-A autonomy only
on passing evidence, and re-probe on any provider/ID change.** This is not new machinery; it *is* our existing
governance, pointed at lane admission:

| Model-reality need | Our locked mechanism |
|---|---|
| the qualification probe (canary → pass/fail evidence) | **m-3** observe-as-send-gate / done-predicate / evidence ladder |
| the `model × provider × serving` triple + capability prior | **m-4** routing record + capability priors (already `model × provider × region`) |
| the per-lane adapter shim + residency/jurisdiction flags | **m-1** courier / identity boundary (the adapter is *inside* the TCB, not glue) |
| "re-probe on ID change; treat IDs as versioned" | the versioning stance in the public-release-intent design note |

**Practical floor:** *floor autonomy at the GLM-5 / Kimi-K2.5 tier on a **canary-qualified** first-party
full-precision lane; route everything below (Qwen3-Coder, gpt-oss, Devstral, quantized-anything) — and any
un-qualified generic-proxy lane — as governed executors behind a stronger planner; treat serving-precision,
provider-jurisdiction, compat-mode, and ID-pinning as first-class routing dimensions; probe, don't trust the
leaderboard.*

### 6.6 Governance flags on hosted non-frontier lanes (hard, from `m-agentic-reality`)
- **Residency/jurisdiction — a hard flag on *all* Chinese first-party APIs** (DeepSeek, Moonshot/Kimi,
  Zhipu/GLM, Alibaba/Qwen, MiniMax): data processed in the PRC under Chinese law; **no DPA/BAA, no EU/US
  residency** — "not suitable for personal/PHI/confidential data" [P policy + C]. Self-hosted weights carry no
  such flag — *this is why provider is part of the routable unit.*
- **Bans attach to the provider, not the weights** — DeepSeek is capability-viable *and* jurisdictionally
  forbidden on many govt devices at once [C]. The router must carry a *jurisdiction-permitted* bit distinct from
  *capability*.
- **Model-ID churn flips capability, not just versions** — DeepSeek silently routes legacy
  `deepseek-chat`/`reasoner` → V4-Flash (IDs deprecate **2026-07-24**), and the pointer change flips
  non-thinking → always-thinking, *re-triggering* the `reasoning_content`/`tool_choice` breakage (§6.2) [P].
  **Pinned IDs are not stable contracts on hosted open-model APIs** — hence re-probe on change.
- **Vendor-vs-weights divergence is the core tradeoff:** first-party API = best capability + residency/churn/ban/
  endpoint-security risk (DeepSeek's exposed unauthenticated ClickHouse leaked secrets + chat logs [C]);
  self-hosted weights = frozen artifact, no residency risk, but you own parser/precision fidelity;
  third-party host = same weights, different quant/parser/context-cap. **Same model name on two hosts = two
  lanes.**

---

## 7. Identity & the courier — channel-stamped FROM, made physical

The c1 lock (`ARCHITECTURE.md` §1) says seat identity = **channel-stamped FROM**: the courier stamps a relay's
author from *which channel the bytes arrived on*, never from a client-supplied field — forgery-robust by
construction. The research confirms this is not only implementable under attachment, it is **the industry's own
mandated pattern**:

- **The conductor runs as an MCP server.** Each seat connects over a **dedicated per-seat transport**: its own
  stdio pipe (for co-located seats) or its own **bearer token** (for HTTP/Streamable-HTTP seats). The courier
  maps transport → seat identity. A seat physically *cannot* present as another seat because it never transmits
  an identity to spoof — it only *is* the pipe/token it holds. **This is MCP's own security guidance**: derive
  identity from the credential, never from client payload. [P]
- **Correction (§14 verification):** an earlier draft claimed Codex "independently does this" via `session_source`
  — **refuted [P].** `session_source` is a mutable launch-arg telemetry/product-gating label (default `vscode`,
  `Custom(String)`), and Codex's per-message author (`clientId`) is *client-supplied*; Codex's crypto identity is
  bound to OpenAI's backend, not exposed as a generic FROM. **The runtime does not supply forgery-robust
  identity — the conductor stamps FROM from its own per-seat channel and treats the runtime as a dumb executor.**
  This makes the design *simpler* (one conductor-owned identity mechanism), not weaker. [P]
- **The conductor must expose Streamable HTTP** (the 2025 MCP transport OpenAI/Google/MS/AWS all adopted) so
  remote/heterogeneous seats attach over per-token channels — table-stakes for provider-agnostic reach. [P][C]
- **Per-seat credential isolation is the whole trust boundary** for the courier: N pipes / N tokens, each
  bound to one seat, none able to read or write another's. Combined with §8's sandbox, this is the physical
  realization of "sole-writer, forgery-robust FROM."

> **This resolves the readiness Cluster-1 sub-worry about identity under attachment:** channel-stamped FROM does
> **not** depend on owning the process tree. It depends on owning *the seat's channel*, which the conductor does
> by construction (it mints the pipe / issues the token). The *egress/sole-sender* half is what needs §8.

---

## 8. The two-tier security posture — confusion-resistant now → by-construction incrementally

The readiness review's root cause #1: "sole external sender / forgery-robust by construction" was proven for a
conductor that **owns the process tree**; a conductor that **attaches** to a rented runtime does not, yet, own
that tree. The research gives a concrete, off-the-shelf, incremental path — not a rebuild.

- **Tier 1 — confusion-resistant (Step-1, attach):** the conductor is the MCP server; seats reach it over
  isolated per-seat channels (§7); the gate/lint/form/observe checks run in the courier. A *cooperative* seat
  cannot forge FROM or bypass the gate. A *determined* seat sharing the host's ambient authority *might* find a
  side channel (e.g. the host's own network egress). Honest claim at this tier: **"confusion-resistant" — the
  governed path is the only *presented* path, and identity is unforgeable, but the conductor is not yet the
  *sole* external sender because it doesn't own the jail.** This is the honest Step-1 posture.

- **Tier 2 — sandboxed defense-in-depth (verified, `srt`):** launch **every seat inside a sandbox jail** using
  Anthropic's **`srt` (`@anthropic-ai/sandbox-runtime`)** — a real, Anthropic-published wrapper (bwrap + netns on
  Linux, Seatbelt on macOS), **minimal overhead** (no container; not separately benchmarked). **Verified (§14):**
  the jail's **only egress *path* is a parent-owned broker socket** — kernel-enforced *by construction* (Linux
  removes the network namespace entirely; macOS Seatbelt `EPERM`s non-localhost *even if the seat unsets proxy
  env vars*). This is genuine defense-in-depth and the transport-layer sole-egress guarantee is real. [P]
  - **Known bound (verified):** Codex's own sandbox bounds *write + network* but **not read** (Landlock
    read-restriction is a documented TODO) — so read-confinement is `srt`'s job, not the host's. [P]

- **Tier 3 — "sole external sender by construction" is a SPIKE-GATED milestone, NOT an automatic upgrade
  (verified §14; spike gate expanded by VP `…-213911`).** The adversarial pass **refuted** the by-construction
  *security* claim: sole-egress-*path* ≠ sole *exfiltrator* — the broker is the **entire egress TCB** (a null-byte
  SOCKS5 allowlist bypass shipped in the wild **~5.5 months**; no TLS inspection; DNS is a covert channel; Codex
  leaves READ open so host secrets stage freely), and the jail **fails OPEN by default** on common kernels
  (namespace-creation failure → self-disable, the Ona incident — the trigger is the *default* on Ubuntu
  23.10+/unprivileged-Docker/WSL). The claim is therefore **two-stage** (§14.2): properties (1)–(5)
  [namespace-failure fails **closed** · broker **canonicalizes hosts** · **no DNS/raw sockets** from the seat ·
  in-jail RCE **can't read a sibling's pipe/token** · conductor **stamps identity itself**] earn only **"single
  mediated network path"**; the stronger **"destination + content control by construction"** additionally requires
  **(6) every outbound broker request passes the conductor egress/content gate** on canonicalized destination,
  method, protocol, payload class, and seat/run evidence. Until the spike passes, even wrapped, the honest claim
  is **"sandboxed defense-in-depth."** [P]

**The posture is therefore a three-rung ladder, and the readiness review's overclaim finding was *right and then
some*:** ship Tier-1 "confusion-resistant" at Step-1; add Tier-2 `srt`-wrap as verified defense-in-depth
(delivering the sole-egress-*path* guarantee); and treat Tier-3 "sole external sender by construction" as a
spike-gated milestone earned by the five properties above — **never shipped on assumption.** This applies the
honest-fallback discipline to I1/egress that the readiness review said was missing — now doubly so, because
verification proved the overclaim risk extends even into the *wrapped* state.

---

## 9. Durable park/wake & away-HITL — validated against the c3 design

The wave that studied durable execution + human-in-the-loop **independently reconstructed the m-6 c3 design**,
which is strong corroboration the locked design is sound:

- **Park/wake = durable-promise / awakeable over the append-only store.** The atomic primitive is
  **compare-and-append on a RESOLVED marker** — which *is* the Cluster-3 spike candidate (`READINESS-REGISTER`).
  **Per-step memoization** (don't re-spend tokens re-running completed steps on resume) and **scan-on-register**
  (a newly-registered waiter checks for an already-resolved result) are the two implementation notes. [P][C]
- **Away-reply / HITL = durable-checkpoint + authenticated-inbound-resume.** The resume token must be
  **POST-not-GET, one-time** — which independently **validates Seam-C** (m-1 owns inbound-token mint/verify via
  the reserved `certification` seam). **The park token *is* the channel-stamped FROM** for the resuming human.
  Staleness / TOCTOU on a parked-then-woken gate is the documented pitfall — which **validates c3 decisions 2
  and 4** (away-token freshness / re-observe binding). [P][C]

> **Net:** the c3 park/wake + away-HITL design needs *no redesign* from this research; it needs the
> **Cluster-3 spike** to de-risk the one atomic primitive (compare-and-append RESOLVED) that everything else
> composes on. That spike also de-risks nonce-burn (m-1) and observe→append (m-3) — one spike, three payoffs.

---

## 10. What this says about decision ① (wrap vs attach + the honest claim)

**Recommendation (now *verified*, §14 — not just evidenced): _attach at Step-1 with the honest
"confusion-resistant" claim; add `srt`-wrap as verified "sandboxed defense-in-depth"; and treat "sole external
sender by construction" as a spike-gated milestone — never shipped on assumption._** Decision ① is not "wrap XOR
attach" — it is a **three-rung ladder**, and verification moved the top rung behind a code spike:

| | Rung 1 — Attach (Step-1) | Rung 2 — `srt`-wrap (defense-in-depth) | Rung 3 — by construction (spike-gated) |
|---|---|---|---|
| **Posture** | Conductor as MCP server; per-seat channels; gate in courier | + every seat in an `srt` jail; only egress *path* = parent-owned broker socket | + the five §14.2 properties **proven by spike** |
| **Egress claim** | **"confusion-resistant"** — governed path is the only *presented* path | **"sandboxed defense-in-depth"** — only egress *path* is the broker (kernel-enforced ✓ verified) | **"sole external sender by construction"** — destination + content controlled, fail-closed |
| **Identity claim** | channel-stamped FROM, **conductor-owned** (not inherited from runtime) | unchanged (broker socket = the channel) | unforgeable **proven** — cross-seat isolation survives in-jail RCE |
| **Honesty status** | ✅ verified honest | ✅ verified real (transport layer) | ⛔ **not claimable until the spike passes** |
| **Gate** | operator ① = *yes to posture + wording* | bounded VP-gated hardening lane | separate operator-opened **spike** gate |

**Why not wrap-first:** unchanged — wrapping before the courier + gate exist confines a seat with nothing to talk
to yet. Attach-first delivers the Step-1 goal (remove operator-as-transport, `ROADMAP.md:21`) and the honest
claim is still strong (only-presented-path + conductor-owned identity).

**What verification *changed* in the recommendation:** (a) `srt`-wrap is **defense-in-depth**, not the thing that
"earns by-construction" — the transport-layer sole-egress *path* is real and by-construction, but destination/
content control and unforgeable-FROM are **not**, and are gated on the five spike properties (§14.2);
(b) **identity is conductor-owned** — the runtime (Codex's `session_source`) does not supply it, which
*simplifies* the design; (c) the jail **fails open by default** on common kernels, so "launch-inside-jail,
fail-closed, no self-toggle" is a hard requirement, not a nicety.

**The one thing decision ① must *not* do:** ship "by construction" — **even wrapped** — before the spike proves
the five properties. The readiness review flagged this for the *attached* state; verification shows the overclaim
risk reaches into the *wrapped* state too. Ship the honest rung you're actually standing on. **This is the whole
substance of decision ①**, now on verified ground.

---

## 11. What this updates in the readiness register

This research is **evidence for**, not a change to, the register. Concretely:

- **Cluster 2 (ride-vs-own honesty) — verified + VP-reviewed, resolved-pending-operator.** §7–§10 give the honest
  **three-rung** posture (§10) and the verified `srt` path (§8, §14); the VP approved it "with tightened spike
  gates" (`runtime-research/…-213911`). On operator decision ①, the m-1 design gets a **bounded completeness fix**:
  if **attach**, record the honest Step-1 claim as **"confusion-resistant"** and state the wrap-upgrade path; if
  **wrap / construction-grade**, route the bounded m-1/m-3/m-4 follow-ups *after* the ① decision, and never
  authorize the "by construction" claim without the two-stage spike (§14.2). A bounded VP-gated fix, not a c1 reopen.
- **NEW — identity-conductor-owned → a bounded m-1 fix (VP `…-213911`, Finding 3).** The verified refutation of
  runtime-supplied identity (§14.1-I) must land in the m-1 design, not stay in research prose: the m-1 fix states
  that **runtime identity fields (`session_source`, `clientId`, …) are never accepted as `FROM` authority** and
  that **conductor-owned per-seat channel/credential isolation is the sole source of the stamp.** Sequenced after
  the operator's ① decision; bounded, VP-gated; not a c1 reopen (it *tightens* the c1 identity lock, doesn't
  contradict it).
- **Cluster 3 (first spike) — sharpened.** §9 confirms the highest-leverage first spike is
  **compare-and-append on a RESOLVED marker in the append-only store**, because it de-risks nonce-burn (m-1),
  park/wake (m-6), and observe→append (m-3) simultaneously. Still gated behind a **separate operator-opened
  spike gate** (VP's readiness-reconcile Finding 4) — this research does **not** open it.
- **No MUST-gate item is cleared by research.** Cluster 1 (m-1/m-2 `submit()` write-path) and Cluster 4a/4b
  (m-2 stale schema + R2 hole) are **doc-contradiction fixes** that research cannot resolve — they still require
  the bounded VP-gated pair-fixes. Research only *confirms* the store must support the atomic append those fixes
  will specify.
- **Versioning stance reinforced.** §6.3's model-ID churn (DeepSeek alias deprecation 2026-07-24; K2/GLM/MiniMax
  treadmills) is concrete external pressure behind the public-release-intent design note — the router must treat
  `model × provider` IDs as versioned with alias tables from day one. §6.6 sharpens it: an ID change can *flip
  capability* (non-thinking→always-thinking), so "re-benchmark on bump" is not optional.
- **New cross-domain requirement surfaced (designed-early, per `ROADMAP.md:24`).** The §6.5 lane-qualification
  probe lands on **m-4** (routing record carries the `model × provider × serving` triple + a *jurisdiction-
  permitted* bit distinct from capability + probe-evidence freshness) and **m-3** (the probe *is* an observe-gate
  / done-predicate applied to lane admission). Neither is a c2 reopen — it's a Step-3 routing/observe design note
  to carry into m-3/m-4 build planning. Logged here, not folded into a locked doc.

---

## 12. Open questions / spike candidates (recommendations only — none opened here)

1. **[Cluster-3 spike]** compare-and-append RESOLVED over the append-only store — the one primitive park/wake +
   nonce-burn + observe→append all compose on (§9). *Highest-leverage first spike.*
2. **[Runtime spike]** the persistent-seat drive against **ACP** — one integration reaches Claude Code + Codex +
   Gemini; validates the "pipe-is-the-inbox" wake with a real ≥8-step loop that **closes and reopens the
   session** (the `reasoning_content` replay trap, §6.2).
3. **[Security spike — GATES the "by construction" claim (§14.2), TWO-STAGE per VP `…-213911`]** an `srt`
   seat-launch wrap with a per-lane broker socket. Stage-A (**"single mediated network path"**): (1)
   namespace-failure fails **closed**; (2) broker **canonicalizes hosts** (rejects NUL/`%`/CRLF — the null-byte
   bypass); (3) **no DNS, no raw sockets** from the seat; (4) in-jail RCE **cannot read a sibling's token/pipe**;
   (5) conductor **stamps identity itself**. Stage-B (**"destination + content control by construction"**): (6)
   **every outbound broker request passes the conductor egress/content gate** on canonicalized destination,
   method, protocol, payload class, and seat/run evidence. Until Stage-A passes, the honest claim caps at
   "sandboxed defense-in-depth"; "sole external sender by construction" needs Stage-B too.
4. **[Model-lane bring-up]** stand up one hosted non-frontier lane end-to-end (DeepSeek-V4-flash non-thinking
   via Anthropic-compat, and one GLM/Kimi T-A lane) through the courier — smallest real proof of the
   provider-agnostic claim.
5. **[Model-governance spike]** the **lane-qualification probe** (§6.5) as an m-3 observe-gate: a canary that
   admits a lane to T-A only on passing (multi-step tool call · structured-edit adherence · forced failed-step
   recovery · `reasoning_content`/`tool_choice` round-trip), re-fired on any provider/ID change. This is the
   concrete artifact that makes "probe, don't trust the leaderboard" real.
6. **[Open]** the exact `get_context_usage`-driven checkpoint policy for persistent seats (§3) — when does a
   standing seat auto-park-and-rehydrate vs auto-compact?

---

## 13. Cross-check against the external chat-model research brief (`RUNTIME-RESEARCH-GPT.md`)

**Framing (governance).** `RUNTIME-RESEARCH-GPT.md` is an **operator-provided external research brief from a chat
model** — well-cited (primary-source-verified, accessed 2026-07-01), honest about its gaps ("Not verified in this
pass"). It is **not a seat**: no channel-stamped FROM, no relay standing, no authority. It is weighed here as
*evidence*, the same epistemic class as a local reference source (not vendored) — **not** as a peer-seat reconciliation. Its value
is that it is a genuinely **independent method** (single-pass, deep per-runtime primary verification) reaching the
same architecture as this fleet (multi-agent breadth) — convergence across independent methods is corroboration.
**The real VP (`master.orchestrator-reviewer`) has not reviewed this synthesis; that adversarial pass is still
owed before decision ①.**

### 13.1 Convergences — independent → lock in
| Conclusion | This fleet | The external brief |
|---|---|---|
| **Routable unit is a tuple, not a model** | `model × provider × serving` (+ local `{…+parser+template+streamed?}`) | `model × provider × serving × compat-mode` — adds compat-mode (adopted, §6.1) |
| **Access is easy; agent-loop *reliability* is the hard part → canary, don't trust the leaderboard** | qualification probe = m-3 observe-gate (§6.5) | "live canary / reliability score that can demote a serving lane without blacklisting the family" |
| **Identity by construction from per-seat transport/credential** | per-seat pipe/token, channel-stamped FROM (§7) | `X-Session-API-Key` / per-seat MCP header / conversation-ID |
| **Wake = blocking pipe / JSON-RPC, not polling** | OS pipe is the inbox, kernel-suspended (§2) | ACP/app-server are push notifications, "a blocking pipe transport, not polling" |
| **Local = first-class governed lane, supervised-until-qualified** | first-class-but-capability-tiered, observed-validity gate (§6.4) | "keep a local governed lane, but supervised-only until qualified on your exact serving stack" |
| **Multi-turn tool-loop breakage is THE risk** | `reasoning_content` replay trap; ≥8-step close/reopen canary (§6.2) | Crush Kimi-K2 multi-turn issue; Cline `<think>` breakage; 20-turn canary |

Six independent convergences on the load-bearing architecture. These are now high-confidence.

### 13.2 The one substantive divergence — resolved into the tuple
GLM-5 / Kimi-K2.5 tiering: **T-A (this fleet)** vs **supervised-only (the brief)** → **capability-T-A, autonomy-
admitted only on a canary-qualified lane; supervised-only on a blind generic proxy.** Full treatment in the §6.5
callout. Net effect: neither was fully right; the merged answer is stronger and *validates* the tuple.

### 13.3 What the brief adds that this fleet under-covered — **folded**
- **HTTP-server runtimes as first-class seat surfaces** (the biggest add). **OpenCode `opencode serve`** (sessions/
  messages as first-class HTTP resources + SSE `/event` stream) and **OpenHands Agent Server** (HTTP+WebSocket;
  per-request identity via `X-Session-API-Key`; an OpenAI-compat front door that preserves conversation identity
  via `X-OpenHands-ServerConversation-ID`) are arguably **cleaner "externally-addressable persistent seat"**
  surfaces than the stdin-pipe path this fleet emphasized. → **Two transport classes:** HTTP-server runtimes
  (first-class, cleaner orchestration) + ACP-subprocess seats (Goose/Gemini, lighter but you own lifecycle/crash).
  *Runtime-matrix §5 upgraded below.*
- **Precision: Claude Code *Remote Control* polls** (cloud-mediated, "registers with the Anthropic API and polls
  for work") — **only the SDK streaming-input path is the non-polling pipe** this fleet's §2 relies on. My core
  claim holds for the SDK path; Remote Control is a *different* feature and must not be conflated.
- **Precision: Gemini CLI base-URL is Gemini/Vertex *proxying*, not proven general multi-provider** (repo requests
  for OpenAI-compat backends still open) — constrains this fleet's "✅ Google + compat" to "ride it for Gemini/
  Vertex, not as an any-model shell."
- **Concrete primary-sourced issue evidence** for the failure classes this fleet described abstractly: Continue's
  `/responses`→`/chat-completions` fallback (`useResponsesApi:false`); Crush MCP schema strictness
  (`None is not of type 'object'`); Cline `<think>`-tag breakage; Crush Kimi-K2 multi-turn failure; Cline "did not
  provide any assistant messages." These *specify* the abstract classes and are excellent canary seeds.
- **Cheapest-experiment designs** (extend §12's spike list): a **deterministic 20-turn canary** diffed across
  native / OpenAI-compat-proxy / local for the same task, seeded with the known failure classes; a **per-seat
  MCP-credential-isolation test** (verify a seat sees tools *only* through its own transport); an **HTTP-vs-ACP
  24h soak** under attach/detach + crash injection.

### 13.4 What this fleet covers that the brief flagged **open** — offered back, and where my confidence must drop
- **The sandbox/jail security survey (§8) — the core of decision ①.** The brief *explicitly lists this as an
  open question it did not verify* (Anthropic sandbox-runtime, bwrap, Seatbelt, Codex sandbox). So §8 is
  **unique-to-this-fleet coverage** — which means it is **exactly what the real VP must adversarially check**
  before decision ①. High-value, low-cross-corroboration → verify the `srt` sourcing.
- **Codex custom base-URL / MCP behavior — the brief did NOT verify it; this fleet *asserted* it.** So the Codex
  "OpenAI-compat + configurable providers + MCP client + channel-stamped FROM" claims (§5, §2) are the
  **least-corroborated** in this doc. **Lowering confidence: treat as fleet-asserted, verify before any lock.**
- **Same-weights-across-providers (K2-Vendor-Verifier 72→100%, §6.1):** the brief could not get a high-confidence
  primary source beyond compatibility-failure evidence; this fleet cited the verifier. My sourcing is stronger
  here — but flag it as the kind of number to re-confirm.
- Richer-in-this-fleet and *not* contradicted: the OS-pipe wake framing (§2), persistent-seat context-growth
  handling (§3), the model-layer depth (pricing / hardware-tier table / KV trap, §6.3–6.4), and the mapping of
  the canary onto our locked domains (probe=m-3, triple=m-4, adapter=m-1, §6.5).

### 13.5 Net
Two independent methods converged on the **same build recommendation** — a thin conductor keeping warm per-seat
transports, identity-by-construction, local governed-but-qualified, everything gated by a live canary. The single
divergence sharpened the model tiering rather than breaking it. **Two items are now the highest-priority to
verify because the brief did *not* corroborate them: the `srt`/sandbox posture (§8, decision ①) and the Codex
attach claims (§5).** Those go to the real VP with the synthesis.

---

## 14. Verification pass — the four load-bearing claims re-checked against primary sources

**Why:** decision ① leaned hardest on the two claims the external brief did **not** corroborate — the `srt`/sandbox
egress posture (§8) and the Codex attach + `session_source` identity (§5). Before locking ① on them, a focused
**four-lens read-only verification** (primary-source-only; one lens adversarial-by-design) checked them against
official docs, the actual GitHub source, specs, and CVE/incident writeups (accessed 2026-06-30). **Result: the
core sandbox primitive is confirmed, but the load-bearing "by construction" language is refuted-as-assumed, and
the Codex identity claim is refuted outright.** The check changed the honest wording available to ① — exactly
what "measure twice" is for.

### 14.1 Per-claim verdicts (primary-source)
| # | Claim (as previously stated in this doc) | Verdict | What's actually true |
|---|---|---|---|
| A | `srt` = `@anthropic-ai/sandbox-runtime`, Anthropic-published, bwrap+Seatbelt | **CONFIRMED [P]** | Real npm pkg + `anthropic-experimental/sandbox-runtime`, CLI `srt`, active (still `0.0.x` preview) |
| B | The jail's only egress is a parent-owned broker socket | **CONFIRMED [P] — at the transport layer** | Linux removes the netns entirely → every external `connect()` fails; only path = parent-owned UDS/proxy. macOS Seatbelt `EPERM`s non-localhost *even if the seat unsets proxy env*. Kernel-enforced, production-shipped. |
| C | …therefore conductor = sole external *sender* **"by construction"** | **REFUTED as by-construction [P]** | Sole-egress *path* ✓; sole *exfiltrator* ✗. The broker is the entire egress TCB — a null-byte parser bypass shipped in the wild **~5.5 months**; no TLS inspection; DNS is a channel; Codex leaves READ open so secrets stage freely. **"By correct implementation," not by construction.** |
| D | Fail-safe behavior of the jail | **REFUTED — fails OPEN by default [P]** | namespace-creation failure → self-disable is the *default* condition on Ubuntu 23.10+, unprivileged Docker, WSL, hardened kernels (the Ona incident). Must be *forced* fail-CLOSED. |
| E | Codex sandbox bounds write+network but **not read** | **CONFIRMED [P]** | codex#11316: Landlock read-restriction "is a TODO"; `cat /etc/passwd` works under `workspace-write` |
| F | Codex `app-server` = persistent, addressable, push-notification seat | **CONFIRMED [P]** | JSON-RPC over stdio / unix-socket / (exp.) ws; `thread/start\|resume\|fork`; persisted, resumable |
| G | Codex is model-agnostic via custom base-URL | **PARTIAL [P] — Responses-API-gated** | Custom `base_url`+`env_key` work, but `wire_api="chat"` is **removed** on `main`; endpoints must speak the OpenAI **Responses API**. Non-OpenAI models need a Responses-compatible gateway (LiteLLM etc.). |
| H | Codex is an MCP client | **CONFIRMED [P]** | stdio + streamable-HTTP; config at global/project scope (not per-connection) |
| I | Codex "independently implements channel-stamped FROM (author from `session_source`)" | **REFUTED [P]** | `session_source` is real but a **mutable launch-arg telemetry/product-gating label** (default `vscode`, `Custom(String)`), NOT channel-derived; per-message author (`clientId`) is *client-supplied*; Codex's crypto identity is bound to OpenAI's backend. **The conductor cannot delegate identity to the runtime.** |

### 14.2 The two consequences that change decision ①
1. **"Sole external sender / unforgeable FROM *by construction*" is refuted-as-assumed → it becomes a
   spike-gated milestone, and the gate is TWO-STAGE (VP review `runtime-research/…-213911`).** The spike must
   prove: (1) namespace-creation failure fails **closed** (no unsandboxed fallback); (2) the broker
   **canonicalizes hosts** (rejects NUL/`%`/CRLF) — kill the parser differential; (3) **no DNS, no raw sockets**
   reachable from the seat; (4) a seat with in-jail RCE still **cannot read a sibling's token/pipe**; (5) the
   conductor **stamps identity itself** from the per-seat channel it owns — never delegated to the runtime.
   - **After (1)–(5), the honest claim is only "single mediated network path"** — necessary but *not sufficient*
     for destination+content control, because a compromised seat can still exfiltrate to an *allowed* endpoint
     through the one permitted path.
   - **(6) — the VP's required addition for the strongest claim:** *every outbound broker request must pass the
     conductor egress/content gate on canonicalized destination, method, protocol, payload class, and seat/run
     evidence reference; generic CONNECT/raw-proxy behavior and broad allowlisted hosts are not enough.* Only with
     (6) does the claim become **"sole external sender with destination + content control by construction."**
   Until the spike passes, even a *wrapped* conductor honestly claims only **"sandboxed defense-in-depth."**
2. **Identity is conductor-owned, full stop.** The runtime (Codex — and by the same logic any rented runtime)
   does **not** supply forgery-robust identity. This *simplifies* the design — one identity mechanism, the
   conductor's own per-seat MCP channel — and it means channel-stamped FROM is a property the conductor
   **builds**, not one it **inherits**. (Corrected in §BLUF-3, §5, §7 below.)
   - **VP routing (`…-213911`, Finding 3):** this must land as a **bounded m-1 design-completeness fix** *after*
     the operator's ① decision — not stay a research caveat in prose. The m-1 fix must state explicitly that
     **runtime identity fields (`session_source`, `clientId`, …) are NEVER accepted as `FROM` authority**, and
     that conductor-owned **per-seat channel/credential isolation is the sole source of the stamp**. (Tracked in
     §11.)

**Neither consequence blocks attach-first.** Both make the honest claim *precise*. The transport-layer sole-egress
guarantee is genuinely by-construction and kernel-enforced; what is *not* by-construction is destination/content
control and unforgeable-FROM — those are earned by the five spike properties. See §10 for the updated ① position.

### 14.3 What this says about the model-agnostic constraint
Item G is a real correction: **Codex is model-agnostic only through the OpenAI Responses API.** For the
provider-agnostic requirement this means **Claude Code is the cleaner multi-model ride** (Anthropic-compat +
gateways), and a Codex lane to a non-OpenAI model needs a **Responses-compatible translating gateway**. Not a
blocker — a routing/adapter fact for m-4 and the Step-3 adapter surface (§4).

---

## Appendix — research fleet roster & provenance

**Framing:** all subagents were **read-only web/reference lenses** (the only charter-sanctioned spawn). No
subagent held a seat, wrote a relay, edited a doc, or touched code. Findings are evidence; authority stays with
the CTO+VP and the operator.

**Wave 1–2 (mechanics & posture, ≈16 agents):** wake-mechanism (pipe-is-the-inbox); seat-model
(ephemeral-vs-persistent); identity/courier (conductor-as-MCP-server, channel-stamped FROM); MCP
provider-agnosticism & competitor ratification; integrate-runtimes-not-models; wrap-as-off-the-shelf (`srt`,
Ona caveat, Codex sandbox read-bound); durable park/wake; away-reply/HITL; model-reality (partial).

**Wave 3 (model layer):**
- `m-open-hosted` ✅ — routable-unit = model×provider×serving; K2-Vendor-Verifier 72→100%; OpenRouter front door.
- `m-glm-minimax-qwen` ✅ — GLM/MiniMax/Qwen all dual-stack; ride coding SKU; alias table; region = lane attr.
- `m-deepseek-kimi` ✅ — DeepSeek strongest default lane / Kimi strongest endurance; **`reasoning_content` replay
  trap**; drive non-thinking in loops; PRC region → Western reseller; pin IDs (alias deprec. 2026-07-24).
- `m-agentic-reality` ✅ — honest tier-list (T-A GLM-5/Kimi-K2.5/DeepSeek-borderline · T-B Qwen3-Coder/gpt-oss/
  Devstral · T-C ≤14B + all ≤4-bit-quant); coding≠agentic by 40+ pts; **probe-don't-trust-leaderboard**; probe =
  m-3 observe-gate, triple = m-4 record, adapter+residency = m-1 boundary; PRC-residency hard flag; ID-churn
  flips capability.
- `m-local-serving` ✅ — local runtimes (vLLM/SGLang lowest-tax; llama.cpp GBNF-valid args; Ollama `/v1`
  streaming-broken; LM Studio desktop/HITL; MLX not-production; TGI dead). Streaming tool-calls broken
  everywhere → **non-streamed tool-turns**. Qwen3-Coder-30B-A3B / Devstral-24B = viable on 1×24 GB;
  GLM-4.5-Air = best open reliability tier. Ruling: **first-class-but-capability-tiered**, gated by observed
  tool-call validity on the `{model+server+parser+template+streamed?}` tuple. *(3 read-only sub-lenses spawned;
  the "still-running" `af368905bd83c174d` is one of them — no separate collection needed.)*

**Tags:** [P] primary · [C] corroborated ≥2 independent · [V] vendor/SEO to-verify. **Not yet independently
re-verified by the CTO line-by-line** — this is a synthesis of read-only agent findings and is labeled
report-only until the operator and VP fold it into decisions.
