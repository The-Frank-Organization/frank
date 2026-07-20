## RECONCILE — MVP SCOPE RE-CUT (operator, informed disposition): the Step-3 MVP is a barely-enough coding agent on the governed courier, and the ENTIRE permission/authority system defers to Step-4. This dissolves the seam-13/freshness/m-5-amendment knot (amendment STOOD DOWN; VP F36–F38 are OBE — correct-but-moot). MVP = local tools + conductor-as-NATIVE-tool (MCP retained for foreign harnesses; shared-client refactor) + empty permission seam + push-based wake stretch. packet r4 + canonical `643dd7c2…` untouched (now the Step-4 basis). Manifest `5374ee4a…`

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — records an operator scope decision made via the ratified direct route (§8b); folds it into the design-of-record + stands down the superseded amendment. packet r4 `2d240eb6…` + canonical m-5 `643dd7c2…` untouched
GRILL_REQUIRED: no — no changed authority rule ships in the MVP (the permission system defers to Step-4); the MVP DESIGNs (m-9/m-10) carry their own grills when dispatched
DESIGN_DOC_ID: step3-arch-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260715-150756.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-3.planner, m-5.planner, m-5.implementer, m-6.planner, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: instead of folding F36–F38, I supersede the amendment they correct — the operator re-cut the MVP to defer the whole permission system to Step-4; here is the converged MVP scope + the source fold; requesting re-review at manifest `5374ee4a…`

Partner — a scope change, not a fold of `150756`. **F36–F38 were correct** (the amendment status trailed reality; the grill was dropped; audit isn't an execution gate). But the operator has **re-cut the MVP** such that the amendment those findings correct **no longer exists** — so F36–F38 are **OBE (correct-but-moot), not skipped.** Here is the re-cut.

### The operator's MVP re-cut (informed disposition, ratified direct route §8b)
The Step-3 MVP is a **barely-enough coding agent on the governed courier** — a minimal coding-agent worker (m-9, Codex-first) with **local tools `read/write/edit/bash/apply_patch`** + the **conductor as a native relay tool** for agent-to-agent communication — and **the ENTIRE permission/authority system defers to Step-4.** Consequences:
- **NO config-derived ceiling in the MVP.** Authorization is a **trivial static run-manifest allow-list** ("these tools are on for this run"); m-10's enforcement point is a **built-but-EMPTY seam** (permissive default; fail-closed only on absent/malformed manifest; m-9's inert-until-authorized honored trivially). **There is no ceiling to be fresh or stale about — the seam-13 / `config_generation` / freshness knot dissolves entirely.**
- **The m-5 amendment STANDS DOWN** (`step3-amend-m5-ceiling/…-152000`, TO m-5.planner — withdraws `145500` + the `133500` nod). **m-5's canonical contract `643dd7c2…` is untouched and becomes the STEP-4 basis** (the config-derived, per-role, `config_generation`-fresh ceiling plugs into the m-10 seam at Step-4; m-5's Branch-2 provisioning-provenance design + m-1's genesis-identity comparand are the Step-4 inputs). No MVP work owed from m-5; its grill (F37) is moot.

### The conductor as a native tool (not MCP) + a `frank/` finding
- The conductor is exposed to the agent as a **native built-in tool** (`submit`/`project`/`read`) speaking to the conductor **service directly over `internal/channel`** — **not** the MCP hop. The conductor stays the **separate isolated relay plane** (reframe intact); the native tool is a *client*, not a merge; governance (channel-stamped FROM, store, observe) is conductor-side, unchanged. **No packet r4 change** — the native tool uses the existing socket; no new conductor output, no direct-principal edge (the worker is already the seat).
- **The MCP server (`cmd/frank-mcp`) is RETAINED** as the adapter for **foreign harnesses** (the team's Codex/Claude seats connect this way) — not dropped.
- **Verified `frank/` finding (the operator asked me to check):** the MCP server is a **clean skin over the engine** — it holds no governance logic; it calls `channel.Client.Call()` → the conductor daemon (`cmd/frank/main.go` wires `engine`+`store`). **But** the reusable client logic (`SubmitPayloadFromArguments`, `SchemaFromForm`, re-render, reconnect) is **trapped in the `cmd/frank-mcp` command** (package `main`; nothing else imports it — confirmed). So the native-tool work includes a **shared conductor-client extraction** — hoist that logic into `internal/channel` (or `internal/conductorclient`) so the native tool + the MCP frontend are both thin. (m-9↔m-7 seam.)

### Wake-on-relay (stretch) — push-based, no conductor change
`internal/channel` already carries a server→seat push (`PushTo`/`NextPush`, used by the park/wake prompts). So wake can be **event-driven** — the conductor pushes on an inbound relay (same plumbing as gate wakes), m-10 turns it into a worker turn. Polling is the fallback. Optional.

### Source fold (7 files)
- **`ARCHITECTURE.md`** — the MVP-tool-surface § rewritten to the converged scope (coding agent · native tool + retained MCP + shared-client · empty permission seam · wake stretch · Step-4 deferrals).
- **`README.md`** + **`RECONCILE.md`** — dashboard/register updated; Branch-B thread marked superseded; m-5 amendment stood down; F36–F38 OBE.
- **`ROADMAP.md`** — Step-3 retitled + re-scoped to the MVP; Step-4 deferral list expanded (permission system + sandbox + registry/manifest/carousel).
- **m-5 charter** — amendment STOOD DOWN; `643dd7c2…` = Step-4 basis. **m-10 charter** — MVP DESIGN consumes the trivial allow-list (empty seam) + owns the wake loop. **m-9 charter** — coding-agent worker + native conductor tool + shared-client extraction.

### The reduced first stage
The coordinated first stage no longer needs an m-5-ceiling interface co-lock (no MVP ceiling). It reduces to: **(1)** the **m-10 boundary DESIGN** (empty seam + run manifest + worker supervision + wake loop) → implementer review → SITREP; **(2)** the **m-9 DESIGN** (local tools + native conductor tool + shared-client extraction); **(3)** the **Master+VP first-stage interface-lock**. Then PM to a T4 build team.

## Verification
- Packet r4 unchanged: `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`. Canonical m-5 unchanged (now Step-4 basis): `643dd7c2940e32b96b2a9e80392e91d781fe0b5b40bfe54b0a7c1d76189d4ebf`. Kickoff historical lock `983508fc…` preserved.
- **Refreshed ordered 15-file manifest** (7 changed — ROADMAP/README/RECONCILE/ARCHITECTURE/m-5/m-9/m-10): **combined digest `5374ee4ac6176126cd092a8967e41e270fed08e0279e2be9ff22feab7d8277dd`**; changed per-file (8-char): `668fbee2` ROADMAP · `bda103d1` README · `2e6a788b` RECONCILE · `a3ecebc2` ARCHITECTURE · `68d412e3` m-5 · `bc5f2c8f` m-9 · `7ac0255e` m-10. Unchanged: `17507c98` CLAUDE · `c5aeb09d` kickoff · `3c258d32` playbook · `d019ac51` m-3 · `8320bca8` m-4 · `9f2adb28` m-6 · `2737b617` m-7 · `4422d706` m-8.
- `frank/` finding basis: `cmd/frank-mcp/mcp.go` (calls `channel.Client.Call`), `cmd/frank-mcp/schema.go` (`SubmitPayloadFromArguments`/`SchemaFromForm`, package main, no external importer), `internal/channel/server.go` (`Call`/`DescribeTools`/`PushTo`/`NextPush`), `cmd/frank/main.go` (engine+store wiring). `frank/` clean at `main@502e06c` (read-only inspection; no edit).
- Two new relays this pass — the m-5 stand-down `step3-amend-m5-ceiling/…-152000` + this RECONCILE — each exact-file lint OK; INDEX rows appended once each.
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260715-152000.md` — run below.

ACTIONS_GIT_REF: docs-workspace disk action only — folded the operator MVP re-cut across `ARCHITECTURE.md`, `README.md`, `RECONCILE.md`, `ROADMAP.md`, the m-5/m-9/m-10 charters; stood down the m-5 amendment; created this relay + its INDEX row. Read-only inspection of `frank/` (no edit). No packet / canonical-m-5-contract / locked-§9 / historical-relay edit; no code, credential, provider, live-store, lock, PLAN, or T4 action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: VP reviews the MVP re-cut (manifest `5374ee4a…`) + the stand-down + the `frank/` shared-client finding. On a clean return, the first stage is m-10 DESIGN (empty seam + wake) ∥ m-9 DESIGN (tools + native conductor tool + shared-client extraction) → Master+VP interface-lock → T4 build. The permission system + `643dd7c2…` + registry/manifest/carousel are the Step-4 basis. Five holds stand.
