# Governance-first brief on a persistent multi-agent courier

## BLUF

Yes, **a single courier can ride multiple existing runtimes and preserve persistent, identity-by-construction seats**, but **not uniformly** and **not with one universal transport**. The strongest “rideable” targets from the material I verified are **OpenCode**, **OpenHands Agent Server**, **Gemini CLI ACP**, **Goose ACP**, and **Claude Code Agent SDK**. They give you either a true server surface or a documented long-lived subprocess protocol. citeturn26view0turn34view0turn19view0turn24view2turn8view3

The biggest caveat is that **provider-agnostic model access is much easier than provider-agnostic agent-loop reliability**. Multiple projects explicitly expose provider-specific toggles, compatibility layers, or known breakages for OpenAI-compatible endpoints, MCP schemas, reasoning models, and multi-turn tool use. In practice, the safe routing unit is **model × provider × serving stack**, not just “model family.” citeturn41view3turn40search3turn40search15turn35search16turn28search1

If the goal is a thin conductor that keeps long-lived seats warm and derives identity from transport or credential, the most defensible path is:

1. **Primary rides:** OpenCode server, OpenHands Agent Server, Goose ACP, Gemini CLI ACP, Claude Code SDK.
2. **Secondary / supervised rides:** Cline SDK hub mode and Continue CLI resume mode.
3. **Avoid assuming universal compatibility:** especially for Kimi / GLM / MiniMax / local open-weight reasoning models behind generic OpenAI-style gateways. citeturn26view0turn34view1turn23view0turn19view0turn8view3turn37view1turn42search5turn40search15turn35search16

This report is strongest on the **runtime transport layer** and **integration-tax evidence**. The **model-economics matrix is intentionally partial** where I did not verify a current primary pricing/model-card source in this pass. All sources cited here were accessed on **July 1, 2026**. Anything involving model prices, model IDs, and benchmark leaders is especially likely to stale quickly. citeturn28search1turn34view1turn41view1

## Runtime ride-ability matrix

The table below answers Q1a–d using only material I verified in this pass. Where a cell says **Not verified in this pass**, that is deliberate rather than inferred.

| Runtime | Long-lived externally addressable session | Wake without polling or timer | Arbitrary model / custom base URL | MCP client and identity-by-construction fit |
|---|---|---|---|---|
| **Claude Code** | **Yes, via Agent SDK streaming input**. Anthropic documents “a persistent, interactive session” and “a long lived process” for Streaming Input mode; sessions are persisted and resumable. Remote Control also exists, but it is cloud-mediated and separate from the SDK. citeturn8view3turn9search0turn9search1turn8view2 | **SDK path: yes** — the process stays alive and accepts queued messages; **Remote Control: no** — Anthropic explicitly says the local session “registers with the Anthropic API and polls for work.” citeturn8view3turn8view2 | **Partial**. Anthropic documents `ANTHROPIC_BASE_URL` and gateway discovery for Anthropic-compatible routing/proxies, plus custom model entries. I did **not** verify support for arbitrary OpenAI-compatible backends. citeturn12view3turn12view0 | **Yes**. Claude Code supports MCP over **stdio, SSE, and HTTP / streamable HTTP**, with credentials passed via per-server `env` or headers. That is compatible with a conductor-as-MCP-server pattern in which each seat gets its own process/env or HTTP credential. citeturn11view1turn11view2 |
| **OpenAI Codex CLI / app-server** | **Yes, via app-server JSON-RPC**. OpenAI’s app-server docs expose `initialize`, requests, and notifications over a long-lived protocol. citeturn5view0 | **Yes**. The app-server protocol has **server notifications** and action requests rather than a poll loop. citeturn5view0 | **Not verified in this pass.** I verified app-server protocol docs, but not a primary source for arbitrary OpenAI-compatible or Anthropic-compatible base-URL routing. citeturn5view0 | **Not verified in this pass.** |
| **Gemini CLI** | **Yes, via ACP mode**. Google documents ACP mode as JSON-RPC over stdio, with Gemini CLI acting as the server; ACP clients manage multiple concurrent conversations and sessions persist to history. citeturn19view0turn23view0 | **Yes**. ACP is explicitly **JSON-RPC over stdio**; this is a blocking pipe transport, not polling. citeturn19view0 | **Constrained yes**. Gemini CLI supports `GOOGLE_GEMINI_BASE_URL` and `GOOGLE_VERTEX_BASE_URL`, but the official docs still frame the product around **Gemini API / Vertex AI** authentication; feature requests for OpenAI-compatible or multi-provider support were still open in the repo. So treat this as **proxying Gemini/Vertex**, not proven “any model.” citeturn20view0turn20view1turn17search3turn19view4 | **Yes**. Gemini CLI documents MCP discovery/execution and supports **stdio, SSE, and Streamable HTTP** for MCP servers. Per-server URL/headers/env make per-seat credentials straightforward. citeturn15view2turn19view3 |
| **Goose** | **Yes, via `goose acp`**. Goose documents ACP over stdio and says ACP clients manage multiple concurrent conversations with isolated state; ACP sessions are saved to session history. citeturn24view2turn22search5 | **Yes**. Goose ACP is stdio JSON-RPC, client-managed. No timer or poll loop is documented for this path. citeturn24view2turn23view0 | **Yes**. Goose supports quick provider setup, OpenRouter, Tetrate Agent Router, and **custom providers** with custom `HOST`, API key, and provider type; docs explicitly mention custom OpenAI / Anthropic / Ollama-compatible providers. citeturn23view1turn23view2turn21search2 | **Yes**. Goose consumes MCP from ACP clients and supports stdio and HTTP transports there; its docs also distinguish API providers vs CLI providers. The identity model is favorable because MCP creds live in per-server env / headers and ACP sessions are isolated. citeturn23view0turn23view3 |
| **OpenCode** | **Strong yes**. OpenCode has both **`opencode serve`** — a headless HTTP server with an OpenAPI spec — and **ACP over stdio** for editor integrations. Sessions and messages are first-class HTTP resources. citeturn26view0turn27view1turn26view1 | **Yes**. HTTP clients can post directly to session/message endpoints; event streams are exposed over **SSE** (`/event`, `/global/event`). ACP mode is JSON-RPC over stdio. citeturn27view2turn27view3turn26view1 | **Yes**. OpenCode’s providers docs say base URL can be customized for any provider; models are backed by AI SDK / Models.dev and local models are supported. citeturn26view2turn26view4 | **Yes**. OpenCode documents local MCP servers and remote defaults; it is clearly an MCP client. Because it supports remote MCP plus per-provider base URLs and per-session HTTP surfaces, it fits the conductor-as-MCP-server pattern well. citeturn26view3turn26view0 |
| **OpenHands** | **Strong yes, if you use Agent Server / Agent Canvas rather than only the legacy CLI**. The Agent Server is an **HTTP and WebSocket API** for long-running conversations and workspace files, and the OpenAI-compatible gateway preserves conversation identity via `X-OpenHands-ServerConversation-ID`. citeturn34view0turn34view1turn32view0 | **Yes**. Agent Server exists specifically to “start conversations” and “stream events”; it is HTTP / WebSocket, not polling. ACP is also available for IDE integrations. citeturn34view0turn33search11turn33search12 | **Yes**. OpenHands says it can connect to **any LLM supported by LiteLLM**, and has provider pages for OpenAI, Moonshot, OpenRouter, LiteLLM proxy, local LLMs, and OpenAI-compatible proxies via base URL. citeturn28search1turn28search0turn28search2 | **Partial yes**. OpenHands clearly exposes ACP support for third-party ACP agents and surfaces MCP in its product docs/navigation, but I did not verify a primary source spelling out Agent Server-as-MCP-client behavior in this pass. Identity-by-construction is strong for Agent Server itself because every REST / WS request can be gated with `X-Session-API-Key`. citeturn33search6turn34view0 |
| **Cline** | **Mixed**. The CLI itself is interactive/headless and resumable, but I did **not** find a first-party standalone external server API in the material I verified. The **SDK** does have a **hub-spoke architecture** with a local daemon and WebSocket-attached clients. citeturn37view0turn38search0turn37view1 | **Yes, in SDK hub mode**. The hub coordinates sessions and routes events; clients attach “over WebSocket,” and sessions continue if a client disconnects. The ordinary CLI/headless mode is just process invocation. citeturn38search0turn37view0 | **Yes**. Cline documents support for OpenAI, Anthropic, Gemini, OpenRouter, Ollama, LM Studio, and “any OpenAI-compatible API”; local-model pages expose base URLs. citeturn35search5turn35search11turn36search10 | **Yes**. Cline is an MCP client and supports **stdio**, **Streamable HTTP**, and **SSE**. That makes per-seat credential scoping feasible, especially if each seat has its own `mcp.json` or SDK config. citeturn37view3 |
| **Crush** | **Unclear / likely no first-party external server surface in the material I verified**. I verified custom-provider support, but not a documented external session protocol comparable to OpenCode HTTP, OpenHands Agent Server, or ACP. citeturn41view0 | **Not verified in this pass.** | **Yes**. The README explicitly documents **custom OpenAI-compatible and Anthropic-compatible providers**, including `openai-compat` with a custom `base_url`. citeturn41view0 | **Not verified in this pass.** There is supporting repo evidence that Crush has MCP-related behavior and issues, but I did not verify a primary MCP docs page during this pass. citeturn40search3 |
| **Continue** | **Weak / supervised only**. Continue CLI supports interactive TUI, headless mode, and `--resume`, but I did **not** verify a documented external server/API surface. citeturn42search3turn42search4turn42search5 | **No documented push surface found**. The verified path is resume/replay via CLI, not an externally-addressable long-lived socket/server. citeturn42search4turn42search5 | **Yes**. Continue docs explicitly support OpenAI-compatible providers with custom `apiBase`, and can force legacy completions or disable OpenAI `/responses`. citeturn41view3 | **Yes**. Continue supports MCP servers and documents **stdio**, **sse**, and **streamable-http**. The identity story is stronger if you treat each seat as its own CLI/config scope rather than sharing one process. citeturn41view2turn39search0 |

### What this means for the build

If your hard requirement is **persistent, individually-addressable, warm seats** that an outside conductor can wake repeatedly, the best fits are:

- **OpenCode** for a clean HTTP-first server surface with sessions, messages, async prompts, and event streams. citeturn26view0turn27view2turn27view3
- **OpenHands Agent Server** if you want a stronger backend/server architecture and are comfortable treating OpenHands as the outer runtime. citeturn34view0turn34view1turn32view0
- **Goose ACP** and **Gemini CLI ACP** for stdio JSON-RPC seats managed as subprocesses. citeturn24view2turn19view0
- **Claude Code Agent SDK** if you are willing to embed process ownership more tightly and accept that Remote Control itself polls. citeturn8view3turn8view2

By contrast, **Continue** and **plain Cline CLI** look more like resumable user-facing CLIs than clean externally-addressable seat servers; **Cline SDK hub mode** can be promoted to first-class only if you are willing to embed its SDK/runtime locally rather than treat it as a stable rented runtime. citeturn37view1turn38search0turn42search5

## Honest model-lane view

This section is **partial by design**. I verified enough to rank lanes honestly, but I did **not** fetch every family’s current first-party pricing page and model card in this pass, so those cells are left blank rather than guessed.

The most useful high-confidence evidence I found is from **OpenHands’ own current model guidance**, which is not the model vendor’s own benchmark, but is still a project-maintainer benchmark and should be treated as **secondary evidence, not ground truth**. OpenHands says its recommendations are based on current **OpenHands Index** results. citeturn28search1

| Family | Endpoint compatibility | Multi-turn tool calling / long-horizon view | Benchmark standing and provenance | Context / price / residency |
|---|---|---|---|---|
| **Claude frontier** | Widely supported by runtimes here; Claude Code is native, Goose/Cline/OpenCode/OpenHands all expose Anthropic or Anthropic-compatible paths. citeturn12view3turn23view2turn26view2turn28search1 | **Best current evidence in this pass**. Claude Code is built for long-lived agent loops; OpenHands lists Claude Opus 4.8 at the top of its current family recommendations. citeturn8view3turn9search4turn28search1 | **High**. OpenHands docs list Claude family as their top cloud recommendation and show the strongest average among the families they display; this is **OpenHands’ benchmark, not Anthropic’s own**. citeturn28search1 | **Not fully verified in this pass.** |
| **GPT frontier** | Supported broadly; OpenAI-compatible access is pervasive across OpenCode, Goose, Continue, Crush, Cline, and OpenHands. citeturn26view2turn23view1turn41view3turn41view0turn35search11turn28search0 | **High but more compatibility-taxed** because several runtimes expose special handling for `/responses`, reasoning summaries, or gateway behavior. citeturn41view3turn5view0 | **High**. OpenHands currently recommends GPT-5.5 among cloud families; this is again **OpenHands’ benchmark, not OpenAI’s self-report**. citeturn28search1 | **Not fully verified in this pass.** |
| **Gemini frontier** | Native in Gemini CLI; also supported through OpenHands, Goose, and Cline. Gemini CLI base URLs are for Gemini / Vertex proxying, not proven general multi-provider support. citeturn20view0turn20view1turn23view1turn35search5turn28search1 | **Viable** for agent loops, but the Gemini CLI transport itself is more Google-specific than the others. citeturn19view0turn19view4 | **High but below Claude/GPT** in the OpenHands recommendations shown. **Secondary benchmark evidence only.** citeturn28search1 | **Not fully verified in this pass.** |
| **DeepSeek** | Runtime support is good at the harness layer because OpenAI-compatible paths exist in Crush, Cline, Continue, Goose, and OpenHands proxies. citeturn41view0turn35search5turn41view3turn21search2turn28search2 | **Mixed**. Compatibility incidents in Cline and other OpenAI-compatible paths are exactly the kind of tool-loop breakage you should expect to test yourself. citeturn35search16 | **Not enough primary benchmark evidence gathered in this pass.** | **Not fully verified in this pass.** |
| **Kimi K2 family** | Accessible through OpenAI-compatible paths and via OpenHands / OpenRouter style setups. citeturn41view0turn28search1 | **Promising but fragile**. Crush has a recorded issue titled “Kimi K2 using OpenAI-compatible API cannot perform multi-turn conversations,” which is exactly the class of problem your conductor must defend against. citeturn40search15 | OpenHands currently lists **Kimi-K2.6** and **Kimi-K2.5** as strong open/open-weight candidates in its recommendations. This is useful but still **secondary benchmark evidence**. citeturn28search1 | **Not fully verified in this pass.** |
| **GLM 4.x / 5.x** | Reachable through multi-provider routes such as OpenHands / OpenRouter style configs. citeturn28search1 | **Potentially viable** but not verified here with first-party tool-loop docs. | OpenHands lists **GLM-5.1** and **GLM-5** among strong open/open-weight candidates. Again, **secondary benchmark evidence**. citeturn28search1 | **Not fully verified in this pass.** |
| **MiniMax** | Reachable in multi-provider setups; Goose and Cline both emphasize broad provider coverage. citeturn23view1turn35search5 | **Potentially viable** but not verified here from first-party tool docs. | OpenHands lists **MiniMax-M3** in strong open/open-weight recommendations. **Secondary benchmark evidence**. citeturn28search1 | **Not fully verified in this pass.** |
| **Qwen / local open weights** | Strong harness compatibility because Continue, Cline, Goose, and OpenHands all document local/self-hosted use through Ollama, LM Studio, local OpenAI-compatible servers, or LiteLLM. citeturn39search20turn36search10turn21search2turn28search1 | **Highly variable**. OpenHands explicitly warns that open-weight and local models “still vary widely in tool-use reliability” and may show malformed JSON or poor long-run behavior. citeturn28search1 | OpenHands recommends starting local/self-hosted OpenHands with **Qwen3.6-35B-A3B**. That is useful directional evidence, but still not a vendor-neutral benchmark industry standard. citeturn28search1 | **Not fully verified in this pass.** |

### Governed-lane tier list

**First-class lanes**

- **Claude frontier**
- **GPT frontier**
- **Gemini frontier**
- **OpenCode / OpenHands / Goose / Gemini CLI / Claude Code as the runtime shells around them**

Reason: these are the lanes with the strongest verified support for persistent transports, resumable sessions, and mature tool-oriented integrations in the sources I checked. citeturn26view0turn34view0turn24view2turn19view0turn8view3

**Supervised-only lanes**

- **Kimi K2**
- **GLM**
- **MiniMax**
- **DeepSeek**
- **Qwen and other local open-weight lanes**

Reason: the harnesses can often reach them, but the issue evidence shows exactly the failure pattern you care about: multi-turn tool and replay fragility, schema drift, reasoning-format drift, and provider-serving quirks. citeturn40search15turn40search3turn35search16turn28search1

**Not-agent-ready as a first-class default**

- **Any open-weight lane you have not separately qualified on your own exact serving stack**

Reason: OpenHands explicitly warns that local/open-weight tool-use reliability varies widely; Cline and Crush issue history shows concrete breakages from that variability. citeturn28search1turn35search16turn40search3

## Integration taxes and the real routing unit

### The single sharpest gotcha

The sharpest gotcha is **assuming that “OpenAI-compatible” means interchangeable for multi-turn autonomous tool loops**. The source set shows the opposite: models and runtimes repeatedly need **provider-specific switches, parser accommodations, or endpoint downgrades** to keep a loop alive. citeturn41view3turn40search15turn40search3turn35search16

### Ranked integration taxes

**OpenAI `/responses` vs `/chat/completions` drift**

Continue documents that for o-series and GPT-5 models it uses OpenAI’s `/responses` endpoint by default, but if you hit **“organization must be verified”** errors related to reasoning summaries or streaming, you should force `useResponsesApi: false` and fall back to `/chat/completions`. That is exactly the sort of runtime-specific patch a naive generic harness misses. citeturn41view3

**Schema strictness mismatches on tools / MCP**

Crush has a concrete MCP compatibility issue where OpenAI-compatible APIs reject MCP tool schemas with the error: **`Invalid schema for function '<mcp tool>': None is not of type 'object'.`** That means even standards-compliant or de facto MCP servers may need normalization before passing through OpenAI-style tool calling. citeturn40search3

**Reasoning / think-tag / tool-format drift on local models**

Cline has an issue where local DeepSeek-style reasoning output caused a failure that users described as the model not using **“a correct tool”**; the issue was specifically about `<think>...</think>` tags making local models unusable. This is classic generic-loop breakage from non-frontier reasoning formats. citeturn35search16

**Multi-turn compatibility failures through a proxy layer**

Crush also has a recorded issue titled **“Kimi K2 using OpenAI-compatible API cannot perform multi-turn conversations.”** Even without a long postmortem in the snippet I verified, the title alone is load-bearing for your design question: a provider-compatible veneer can pass single calls and still fail on agentic replay. citeturn40search15

**Silent provider-surface assumptions**

Gemini CLI’s official configuration exposes `GOOGLE_GEMINI_BASE_URL` and `GOOGLE_VERTEX_BASE_URL`, but the project still documents itself around Gemini API / Vertex auth, and the repo carries open requests for full OpenAI-compatible or multi-provider backend support. That is a warning sign against assuming a runtime can be lifted unchanged from one provider family to another merely because it has a base-URL override. citeturn20view0turn17search3turn19view4

**Missing assistant messages / message-shape assumptions**

Cline issue history includes the explicit error **“Unexpected API Response: The language model did not provide any assistant messages.”** This is another hallmark of a harness assuming one provider’s response shape during compatibility-mode operation. citeturn35search9

### Conclusion on the routable unit

Based on the verified evidence, the routable unit should be:

**`model family × provider endpoint × serving implementation × compatibility mode`**

That is, not merely “Claude Sonnet,” “Kimi K2,” or “Qwen,” but e.g.:

- `kimi-k2 × provider A × native endpoint`
- `kimi-k2 × provider B × OpenAI-compat proxy`
- `deepseek × local Ollama × OpenAI-compat`
- `gpt-5 × OpenAI /responses`
- `gpt-5 × OpenAI /chat-completions via compatibility fallback`

The sources support this because each of those axes changes either:
- the wire protocol,
- tool schema enforcement,
- streaming shape,
- reasoning-summary behavior,
- or the model ID / alias that the runtime actually accepts. citeturn41view3turn40search15turn40search3turn20view0turn28search1

So your conductor should keep:

- a **provider-aware alias table**,
- a **capability matrix per runtime + provider + model**,
- a **sticky session mapping**,
- and a **live canary / reliability score** that can demote a specific serving lane without blacklisting the whole model family. citeturn26view0turn34view1turn38search0

## Governed build recommendation

A practical, governance-first design from the verified material is:

### Use two primary transports

**HTTP server runtimes for first-class seats**

- **OpenCode server**
- **OpenHands Agent Server**

These have the cleanest externally-addressable session semantics. OpenCode exposes explicit session/message APIs and event streams. OpenHands Agent Server is purpose-built as a long-running backend callable over HTTP/WebSocket, and even exposes an OpenAI-compatible front door for legacy clients while preserving a conversation header for continuation. citeturn26view0turn27view1turn27view2turn34view0turn34view1

**ACP subprocess seats for secondary seats**

- **Goose ACP**
- **Gemini CLI ACP**
- **OpenHands CLI ACP where needed**

These give you long-lived subprocess control without inventing your own harness, but you own process lifecycle and crash recovery. citeturn24view2turn19view0turn33search11

### Use identity by construction

Derive seat identity from one of:
- **HTTP base URL + session API key**
- **MCP server URL + per-seat auth header**
- **subprocess command line + env bundle**
- **conversation/session ID returned by the runtime**

This is already how the better runtimes want to work:
- OpenHands gates every REST / WebSocket request with `X-Session-API-Key`. citeturn34view0
- OpenCode has explicit session IDs and HTTP endpoints. citeturn27view1turn27view2
- Claude Code, Goose, Gemini CLI, Cline, and Continue all scope MCP credentials through per-server env / headers / config entries. citeturn11view2turn23view0turn19view3turn37view3turn41view2

### Keep local / self-hosted as a first-class governed lane, but not as a blind default

The runtime docs are clear that local/self-hosted lanes are easy to wire up:
- Continue supports OpenAI-compatible local providers such as **vLLM** and **llama-cpp-python**. citeturn41view3
- Cline documents **Ollama** and **LM Studio** local configuration. citeturn36search10
- Goose documents custom **Ollama-compatible** providers. citeturn21search2turn23view1
- OpenHands explicitly recommends local/self-hosted setups but warns that tool-use reliability varies widely. citeturn28search1

So the policy answer is:
- keep a **local governed lane**,
- but treat it as **supervised-only until qualified** on your exact model + serving stack. citeturn28search1turn35search16turn40search15

## Top build risks, cheapest resolving experiment, and limitations

### Top build risks

**Runtime mismatch between “persistent” and “externally-addressable”**

Some tools persist sessions for user convenience without offering a real external control surface.  
**Cheapest experiment:** build one seat adapter each for **OpenCode server** and **Goose ACP**, then measure whether a second external client can attach, send a follow-up turn, and preserve identity without replay hacks. citeturn26view0turn24view2turn38search0

**Compatibility-mode tool loops fail after the first few turns**

This is the most probable “works in demo, fails in sprint” failure.  
**Cheapest experiment:** run a deterministic 20-turn canary across the same task on:
- native frontier API,
- OpenAI-compatible proxy,
- local OpenAI-compatible server,
then diff tool-call continuity and final outcome. Seed the canary with the exact known failure classes from Crush / Cline / Continue. citeturn40search3turn40search15turn35search16turn41view3

**Seat identity leaks when multiple seats share one MCP credential or one runtime process**

If two seats share one HTTP MCP endpoint token or one runtime-global config, your “identity by construction” invariant quietly collapses.  
**Cheapest experiment:** create one conductor-as-MCP-server per seat with a distinct header/env secret and verify that the runtime exposes tools only through that seat’s own transport config. citeturn11view2turn19view3turn37view3

**Serverful lanes drift less than ACP subprocess lanes, but cost more operationally**

HTTP runtimes are cleaner for orchestration; ACP runtimes are lighter but more brittle on lifecycle and crash handling.  
**Cheapest experiment:** compare 24-hour soak tests for one OpenCode HTTP seat and one Gemini ACP seat under repeated attach/detach, workspace edits, and crash/restart injection. citeturn26view0turn19view0

**Local/open-weight governance can look better on paper than in agent loops**

The docs and issue trackers repeatedly warn or demonstrate that local/open-weight lanes can degrade badly in real tool use.  
**Cheapest experiment:** choose one local Qwen lane and one local DeepSeek lane behind the exact local server you plan to support, then run the same multi-turn tool suite used for frontier lanes. Gate promotion on that evidence, not model-family reputation. citeturn28search1turn35search16

### Open questions and limitations

I did **not** verify enough primary material in this pass to fully answer:
- current first-party **pricing** and **context windows** for every model family requested,
- a clean primary-source answer on **Codex custom base-URL / MCP behavior**,
- a clean primary-source answer on **Crush MCP transport and external session control**,
- the requested **sandbox/jail survey** across Anthropic sandbox-runtime, bubblewrap, macOS Seatbelt, and Codex built-in sandbox,
- and a high-confidence, primary-source comparison of **same weights across different providers** beyond the compatibility-failure evidence already cited.

Where the matrix says **Not verified in this pass**, that is because I preferred an explicit gap over an uncited inference. citeturn34view0turn26view0turn41view0

## Sources

### Primary or primary-adjacent sources

All accessed **July 1, 2026**.

- **Claude Code Docs**
  - Remote Control. citeturn8view2
  - Agent SDK Streaming Input. citeturn8view3
  - Sessions and external session storage. citeturn9search0turn9search1
  - MCP docs and env vars. citeturn11view1turn11view2turn11view0turn12view3

- **OpenAI Codex / app-server docs**
  - Codex app-server protocol and notifications. citeturn5view0

- **Gemini CLI docs**
  - ACP mode. citeturn19view0
  - Configuration and base URLs. citeturn19view1turn20view0turn20view1
  - Session management. citeturn19view2
  - MCP server integration. citeturn15view2turn19view3
  - FAQ on supported third-party use. citeturn19view4

- **Goose docs**
  - ACP and session behavior. citeturn24view2turn22search5
  - CLI session commands. citeturn24view0turn24view1
  - Provider and custom endpoint configuration. citeturn23view1turn23view2

- **OpenCode docs**
  - Server HTTP/OpenAPI surface. citeturn26view0turn27view1turn27view2turn27view3
  - ACP mode. citeturn26view1
  - Providers and base URLs. citeturn26view2turn26view4turn26view5
  - MCP servers. citeturn26view3

- **OpenHands docs / repo**
  - Agent Server package. citeturn34view0
  - OpenAI-compatible endpoint. citeturn34view1
  - LLM overview and recommendations. citeturn28search1
  - OpenAI / LiteLLM proxy provider pages. citeturn28search0turn28search2
  - CLI terminal/headless docs and ACP docs. citeturn29search0turn32view1turn33search11turn33search12
  - Current repo README on Agent Canvas / Agent Server architecture. citeturn32view0

- **Cline docs**
  - CLI overview. citeturn37view0
  - SDK / ClineCore / hub-spoke architecture. citeturn37view1turn37view2turn38search0
  - MCP docs. citeturn37view3
  - Local model docs. citeturn36search10

- **Crush repo**
  - README custom providers / OpenAI-compatible and Anthropic-compatible config. citeturn41view0

- **Continue docs**
  - CLI / TUI / headless / resume. citeturn42search3turn42search4turn42search5
  - Config reference and MCP support. citeturn41view2
  - OpenAI-compatible configuration and `/responses` fallback. citeturn41view3
  - Project status page. citeturn41view1

### Primary-source issue trackers and bug evidence

All accessed **July 1, 2026**.

- Gemini CLI open issues for missing multi-provider / custom endpoint support. citeturn17search3turn17search9
- Cline issue on local reasoning-model / think-tag failure. citeturn35search16
- Cline issue on missing assistant messages from compatible providers. citeturn35search9
- Crush MCP schema strictness issue. citeturn40search3
- Crush Kimi K2 multi-turn conversation issue. citeturn40search15

### Vendor, benchmark, or marketing material that should not be mistaken for neutral ground truth

All accessed **July 1, 2026**.

- OpenHands recommendation table / OpenHands Index summary: useful, but it is still a project-maintainer benchmark, not a universal standard. citeturn28search1
- Goose product/blog pages. citeturn21search9turn23view5
- OpenCode marketing pages / changelog snippets. citeturn25search10turn25search9
- Cline product pages. citeturn35search3turn35search12
