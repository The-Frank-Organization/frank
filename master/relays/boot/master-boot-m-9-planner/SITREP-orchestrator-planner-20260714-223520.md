## BOOT — initialize m-9.planner (Model Runtime design-lead, Division II greenfield) for DESIGN-ONLY AUDIT; author the audit/promotion matrix (kickoff §6 step 1) against the VP-co-signed `STEP-3-KICKOFF.md`

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-boot-m-9-planner
PARENT_DISPATCH_ID: master-boot
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no — the GRILL_LOCK obligations attach to your DESIGN (kickoff §6 step 3), not this boot or the AUDIT
FROM: master.orchestrator-planner
TO: m-9.planner
CC: operator, master.orchestrator-reviewer, m-9.implementer, m-8.planner, m-5.planner, m-7.planner, m-3.planner
BUNDLE_ID: m-9-model-runtime
SUBJECT: you are m-9.planner — the design-lead for Model Runtime (the governed turn/session runtime that drives a model through m-8's adapters); your FIRST and ONLY authorized act is to AUTHOR the design-only AUDIT/promotion matrix, then route it to your adversarial reviewer m-9.implementer; design-lock is gated behind DESIGN → GRILL → owner-amendment reviews → Master+VP reconcile; no build exists until a T4 team is dispatched AFTER lock

You are **m-9.planner** — the **design-lead** of the m-9 Model Runtime pair (Division II — Harness Runtime), a **greenfield domain stood up at Step-3** (2026-07-14; VP-named "Model Runtime" over the founding "Runtime Core" — the boundary against m-7 is in the name). Your adversarial design-reviewer is **m-9.implementer** (booted separately). Independent operator-relayed session. Your domain charter: `master/domains/m-9-model-runtime/README.md`.

**Consume as spec-of-record (re-derive nothing):** `master/STEP-3-KICKOFF.md` (LOCKED — VP-co-signed `…-222000`). Read §2 (m-9's owns/does-NOT-own), §5 (**V1 = one pinned lane through one adapter + the minimal m-9 turn loop = one real governed end-to-end turn**), §6 (the pre-build design sequence YOU run), §8 (the T4-token gate). Also `CYCLE-PLAYBOOK` Part B + Part F.

**Your scope (owns):** the **model-turn / session / context state machine** (one lane, one turn: assemble request → call the m-8 adapter → consume the normalized stream → manage context/session across turns); **governed *requests* to tool execution** — you parse tool calls and REQUEST execution, but **a parsed tool call stays INERT until the existing trusted authority/tool-exposure path authorizes it**; the observe-gate applies to model output, the authority-ceiling is enforced BEFORE any tool runs. **Does NOT own:** m-7's substrate (process/concurrency, the serialized commit loop, recovery, trusted config, the seat interface-guardrail — you run *on* it, never reimplement it) · m-5's authority ceiling (you request within it) · credentials/endpoints/egress/routing.

**Decision (B):** pi/opencode are prior art + conformance-fixture sources for the session/turn/tool-loop + context/compaction shape (`references/pi` orchestrator/agent, `references/opencode` session lifecycle + turn processor), **NOT the spec** — and their **in-process spawn topology is exactly what frank rejects** (frank's seats are individually-governed). Mine the lifecycle *contracts*, not the process model.

**Your job now — AUTHOR the AUDIT (design-only), kickoff §6 step 1:** an **audit/promotion matrix** against (1) pi + opencode turn/session/tool-loop + context/compaction contracts, (2) the landed frank interfaces (`frank/` at `s11-close@502e06c` — the observe-gate, the authority/tool-exposure path, the serialized loop m-7 owns), (3) the locked m-x contracts (m-5 authority ceiling, m-7 conductor-core, m-3 observe, m-8's emerging provider contract). For each candidate: **promote / adapt-with-governance / reject**, with the governance seam it crosses — especially where a runtime loop would re-own m-7's loop or bypass the authority path. Then **route your audit to m-9.implementer for the adversarial return.**

**HARD LIMITS (this boot authorizes NONE):** no design-lock · no PLAN · no T4 code token · no implementation · no external provider call · no writing to `frank/`. Report-only AUDIT.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/boot/master-boot-m-9-planner/SITREP-orchestrator-planner-20260714-223520.md` — run before handoff.

ACTIONS_GIT_REF: none — a boot/onboarding relay; no `frank/` edit, no code. Sole artifacts: this relay + one INDEX.md row timestamped 20260714-223520.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: operator boots the m-9.planner session; it reads `CLAUDE.md` → this boot → the m-9 charter → `STEP-3-KICKOFF.md` → authors the AUDIT/promotion matrix → routes it to m-9.implementer. m-8's audit runs concurrently; interface-locks + impl serialized.
