## Team m-1 — Trust & Identity (AUDIT)

ROLE: Orchestrator Planner
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c1-audit-m-1
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — read-only audit; surface operator-judgment items in findings
FROM: master.orchestrator-planner
TO: m-1.planner, m-1.implementer
CC: master.orchestrator-reviewer, operator
BUNDLE_ID: m-1-trust-identity
OWNER: m-1 (Trust & Identity)

Phase scope — AUDIT (read-only). Inspect source and docs, run safe read-only commands, and produce an independent paired audit with findings. Not in scope: any edits, branches, commits, PRs, scaffolding, or prototype code. This is the v3 research + design phase — no implementation exists yet and none is authorized; implementation would begin only on a future bare own-line `DISPATCH IMPL` addressed to the Implementer, which is not part of this cycle.

Pair roles & research method:
- m-1.planner (Claude Opus 4.8, high thinking): lead the audit and surface design questions. Use multiple parallel agents, websearch, and a deep-research workflow for the prior-art sweep (channel-stamped / trusted-courier sender-identity prior art, agent identity-forgery defenses, append-only sole-writer stores).
- m-1.implementer (GPT-5.5): run an INDEPENDENT audit — do not mirror the planner — and challenge or answer design questions. You lack a built-in deep-research skill, so use subagents and websearches for your research.
- Independent paired audit: each member audits separately, then reconciles. The Planner does not spawn, direct, or simulate the Implementer; the two seats are independent operator-relayed sessions.

Domain context. m-1 owns the conductor's trusted core (the TCB): the sole-writer append-only relay store, the seat-stamper (channel-stamped FROM = identity by construction), store isolation, and where the inline lint/form gate runs pre-delivery. The whole v3 thesis turns on one differentiator neither prior-art harness makes: a stamping courier that is the SOLE writer of the relay store and stamps FROM from the connection a relay arrived on, never a lane-supplied field.

Sources to audit (cross-check the export's distillation against the real source):
- v2.8.8 current state: extracted/agentic-dev-team-skills-v3-export/v2.8.8-release/v288-unzipped/agentic-dev-team-skills-v2.8.8/ — the .relays store convention, tools/relay-lint.py FROM/ROLE handling and the orchestrator-review gate (confusion-robust, not forgery-robust), and protocol.md's identity/addressing model. Also our LIVE store at master/relays/, a running instance of the convention.
- references/jcode (Rust): the swarm/channels bus, MultiProvider per-seat routing, wire-level model attribution, the SafetySystem review-queue, and the agent-facing assign_role swarm action (self-asserted identity). Locate these under crates/ and src/.
- references/claude-code (TS): the Agent Teams subsystem — per-agent JSON inbox files plus lockfile, poll-and-inject delivery, roles as agentType strings, and the self-written from guarded by a literal team-lead string check. Locate under mcp-server/ and the agent/prompts sources.
- Export design intent: extracted/agentic-dev-team-skills-v3-export/v3-design/v3-adaptive-routing-pillar.md (the conductor's 6 components, channel-stamped FROM, the seat-trust model) and EXTERNAL-REFERENCES.md.

Design question to resolve:
What is the minimal store-plus-identity primitive that yields forgery-robust SEAT identity by construction — the sole-writer/stamping boundary, the store's on-disk shape, and the isolation level — such that it rides existing runtimes in Step 1 yet extends to the standalone runtime later without re-cutting?

Hard acceptance criteria:
1. A 4-bucket verdict (still-open / already-closed / product-overlapped / recommended-next) on the v3 m-1 store + identity primitive vs what v2.8.8 and the references already provide.
2. A named comparison of the three self-asserted-identity gaps (v2.8.8 agent-authored FROM/CC; jcode assign_role; claude-code self-written from) and exactly how a channel-stamped FROM closes each by construction.
3. A design recommendation (the seed of the Step-1 interface sketch): the relay-store API surface, the identity-stamp mechanism, and the isolation boundary.
4. A boundary contract naming what m-1 must expose to consumers — m-3 observe-as-send (probe-from-outside-the-lane isolation), m-4 routing records, m-6 human-gate/email projection — so the DESIGN phase can lock interfaces consumer-validated.

Boundary contract:
- Writes: the relay-store API + identity stamp that m-2/m-3/m-4/m-6 build on.
- Reads: v2.8.8 relay/identity handling; jcode and claude-code identity models; the export design intent.
- Target entity: the v3 relay store + seat-stamper primitive (a design recommendation, not code).
- Downstream consumer: m-2 (forms ride the store), m-3 (observe probes the lane), m-4 (routing records), m-6 (inbox = store projection).
- Proof: E1 source citations (file:line in v2.8.8 and the references) backing every gap and claim.
- No-consumer action: flag any proposed primitive with no downstream consumer.

AUDIT-FIRST GATES — may reject or narrow:
0. Duplicate/already-built check: does v2.8.8 or a reference already provide a sole-writer/stamping store we should promote rather than rebuild? If so, recommend promote/wire, do not rebuild.
1. If the channel-stamped-FROM primitive cannot ride an existing runtime in Step 1 (needs the standalone runtime first), say so and route the dependency to the orchestrator.

Out of scope:
- m-2 form-schema internals (sibling audit c1-audit-m-2), the router/policy, the TUI/email-client UX, and any code.

Deliverable:
A file-relay AUDIT report from each seat independently, then a reconciled view: the 4-bucket verdict; the three-gap identity comparison; the store + identity design recommendation; the consumer boundary contract; evidence levels (E1 source cites); operator-judgment items or none. No source changes, no PR. Include ACTIONS_GIT_REF and FINAL_GIT_STATUS_SHORT (read-only phase; cwd is not a git repo, so the structured unavailable form is expected).

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
