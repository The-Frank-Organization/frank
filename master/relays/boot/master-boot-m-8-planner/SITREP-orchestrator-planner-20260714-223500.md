## BOOT — initialize m-8.planner (Provider Adapters design-lead, Division II greenfield) for DESIGN-ONLY AUDIT; author the audit/promotion matrix (kickoff §6 step 1) against the VP-co-signed `STEP-3-KICKOFF.md`

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-boot-m-8-planner
PARENT_DISPATCH_ID: master-boot
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no — the GRILL_LOCK obligations attach to your DESIGN (kickoff §6 step 3), not this boot or the AUDIT
FROM: master.orchestrator-planner
TO: m-8.planner
CC: operator, master.orchestrator-reviewer, m-8.implementer, m-9.planner, m-1.planner, m-3.planner, m-4.planner, m-7.planner
BUNDLE_ID: m-8-provider-adapters
SUBJECT: you are m-8.planner — the design-lead for Provider Adapters (frank's model-abstraction layer); your FIRST and ONLY authorized act is to AUTHOR the design-only AUDIT/promotion matrix, then route it to your adversarial reviewer m-8.implementer; design-lock is gated behind DESIGN → GRILL → three owner amendments → Master+VP reconcile, and no build exists until a T4 team is dispatched AFTER lock

You are **m-8.planner** — the **design-lead** of the m-8 Provider Adapters pair (Division II — Harness Runtime), a **greenfield domain stood up at Step-3** (2026-07-14). Your adversarial design-reviewer is **m-8.implementer** (booted separately). You are an independent operator-relayed session; consult by relay, you are not a subagent. Your domain charter: `master/domains/m-8-provider-adapters/README.md`.

**Consume as spec-of-record (re-derive nothing):** `master/STEP-3-KICKOFF.md` (LOCKED — operator-ratified + VP-co-signed `step3-prep/RECONCILE-orchestrator-reviewer-20260714-222000`). Read §1 (decision B + the frank-owned contract), §3 (single-writer ownership + the m-8↔m-4 seam), §4 (the four-axis lane catalog), §5 (vertical-first + the E3 live floor), §6 (the pre-build design sequence YOU run). Also `CYCLE-PLAYBOOK` Part B (the AUDIT→…→lock cycle) + Part F (the operating model you PM under, AFTER lock).

**Your scope (owns):** the frank-owned normative provider contract (request · normalized event · tool-call · reasoning-replay · usage · finish/error · cancellation · retry/idempotency · timeout/backpressure · partial-stream); provider wire translation + normalized events (incl. `reasoning_content` handling); the **factual lane catalog** (single writer — `{model_id, provider_id, serving_profile_id, compat_mode}` + spec-sheet, seeded/models.dev-shaped/pinned, NO secrets); conformance fixtures. **Does NOT own:** credentials · endpoint selection · egress policy · authority enforcement · routing judgment (the catalog is facts only; m-4 owns the policy overlay keyed to your lane IDs).

**Decision (B), as the VP sharpened it:** pi/opencode are **prior art + conformance-fixture sources, NOT the normative interface** — pi embeds assumptions frank cannot inherit (caller-supplied credentials/headers, arbitrary payload-mutation callbacks, its own retry/stream protocol); those become fixtures that must PASS frank's contract, not define it. **frank owns the contract.**

**The three MANDATORY pre-lock owner amendments your domain sits behind (§1/§3/§6):** m-7 provider-request-egress host + m-3 provider-request-egress class (the existing egress gate is the dormant away-email scanner that flags `claude-`/`gpt-` strings — unusable as-is); the m-7 credential/trusted-config amendment (none exists today); the m-4/m-2 exact-lane routing-record amendment (locked `chosen_model` can't bind the four-axis lane). **You are a CONSUMER REVIEWER on these, not the author** — surface their necessity in your AUDIT; they must close before your lock.

**Your job now — AUTHOR the AUDIT (design-only), kickoff §6 step 1:** an **audit/promotion matrix** against (1) pi (`references/pi`), (2) opencode (`references/opencode` + `references/opencode-notes.md`), (3) the landed frank interfaces (`frank/` at `s11-close@502e06c` — submit/observe/config, `internal/egress`, `internal/fieldspec/registry.json`, m-7 config), (4) the locked m-x contracts (m-1 TCB, m-3 egress, m-4 routing, m-7 conductor-core). For each candidate: **promote / adapt-with-governance / reject**, with the reason + the governance seam it crosses. Name every contract element frank must own fresh vs. every behavioral invariant/fixture worth porting. Then **route your audit to m-8.implementer for the adversarial return.**

**HARD LIMITS (this boot authorizes NONE):** no design-lock · no PLAN · no T4 code token · no implementation · no credential use · no external provider call · no writing to `frank/`. Report-only AUDIT.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/boot/master-boot-m-8-planner/SITREP-orchestrator-planner-20260714-223500.md` — run before handoff.

ACTIONS_GIT_REF: none — a boot/onboarding relay; no `frank/` edit, no code. Sole artifacts: this relay + one INDEX.md row timestamped 20260714-223500.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: operator boots the m-8.planner session; it reads `CLAUDE.md` → this boot → the m-8 charter → `STEP-3-KICKOFF.md` → authors the AUDIT/promotion matrix → routes it to m-8.implementer. m-9's audit runs concurrently; interface-locks + impl serialized.
