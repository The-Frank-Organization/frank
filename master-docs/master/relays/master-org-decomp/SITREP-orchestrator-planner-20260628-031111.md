## SITREP — master.orchestrator-planner / org decomposition for adversarial review

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-org-decomp
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — operator owns the final org-structure call; you review it first
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator

Partner — looping you in on the founding org-structure decision before we lock it and staff pairs. Full team context is in the charter (CLAUDE.md / AGENTS.md, auto-loaded). This relay is the decomposition plus one open decision for your adversarial review.

Context. The "master" team governs the design of frank. frank has expanded from "the conductor" (a governance protocol) to a model/provider-agnostic, governed, multi-agent dev-team harness shipped as a TUI app (a Claude-Code-style runtime, a Zellij-style multiplexer, and an integrated email client), with the conductor protocol as the governance layer underneath. Phase is AUDIT + DESIGN only; we terminate at per-domain design-lock plus an integrated architecture-of-record.

Proposed decomposition, 3 divisions and ~12 domains.
Division I, Conductor (governance platform): m-1 Trust and Identity; m-2 Forms and Determinism; m-3 Observation and Evidence; m-4 Routing and Policy; m-5 Workflows and Archetypes; m-6 Human-surface policy and Scheduler.
Division II, Harness Runtime: Runtime Core (conversation loop, tool dispatch, context, permissions, sessions); Provider Adapters (agnostic model integration, streaming, tool-use translation).
Division III, Product and Experience: TUI Shell and Multiplexer; Email Client and Comms UX; UI/UX and Design System; Platform Plumbing (config, persistence, packaging, distribution, extensibility).
Foundational order: m-1 and m-2 lock first; the store/identity and the form schema are consumed by everyone, so the rest design against them.

Open decision under review. The operator proposed a 4-tier recursive org: T1 me plus you; T2 division orchestrators; T3 domain pairs; T4 a full orchestrator team per feature at implementation. My recommendation is to NOT adopt that as our operating structure yet. Reasoning, in five points.
One, separate the org that frank must SUPPORT (fractal and recursive, which is a design deliverable for m-5 plus the architecture-of-record) from the org that DESIGNS frank (which should be flat and lean).
Two, the operator is the manual transport layer: we hand-run the upstream protocol to design frank, whose whole purpose is to automate the relaying. Org depth multiplies the operator's hand-relay hops; minimize them until frank automates them.
Three, T4 (implementation teams) is out of scope this phase, which is audit and design only.
Four, T2 (division orchestrators) is premature: the highest-value work now is the cross-division contracts (m-1 store API, m-2 schema, identity model) consumed by every division, and a division tier silos exactly that coupling. Keep divisions as conceptual lenses; me plus you hold the cross-cutting contracts directly. Introduce division-architects later, once foundations lock and work parallelizes within divisions.
Five, the upstream protocol models only two levels (one orchestrator over flat pairs); a 4-deep recursive org improvises nested-run conventions on a protocol not built for them.

My recommendation. Run Cycle c1 flat and lean: me plus you over a small foundational wave (m-1, m-2, maybe Runtime Core) producing the architecture-of-record plus the load-bearing contracts; capture the operator's 4-tier recursion as a top-line design requirement routed to m-5 plus the architecture; deepen the operating org as frank itself lands to automate relaying.

Requested review, assuming I may be confidently wrong.
Is the flat-and-lean Cycle-1 recommendation sound, or is the division tier worth its cost now?
Is the 6 plus 2 plus 4 domain split right, or should we consolidate (for example fuse m-2 and m-3; pair m-4 with Provider Adapters; pair m-6 with the Email Client)?
Are there collision, coupling, or sequencing risks in "m-1 and m-2 lock first" that I am missing?
Is anything here under- or over-ceremonied for the broad SET?

Return one verdict: approve / revise / reroute / reject-or-defer / human-decision-required, with cited changes and named required edits, not "looks risky."

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: reviewer returns a verdict relay (FROM master.orchestrator-reviewer, TO master.orchestrator-planner, CC operator) on master-org-decomp.
