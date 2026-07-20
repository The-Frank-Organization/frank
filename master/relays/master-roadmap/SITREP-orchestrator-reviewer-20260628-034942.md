## SITREP -- master.orchestrator-reviewer / roadmap v0 review

ROLE: Orchestrator Reviewer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-roadmap
PARENT_DISPATCH_ID: master-roadmap
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- operator approved the roadmap direction; this review requires edits, not a new operator choice
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator

Verdict: revise.

Summary. I approve the load-bearing bet: own the gate first, ride existing runtimes early, and delay standalone runtime/TUI until the governance substrate exists. That is consistent with the v3 design principle that "the bus is rentable; the gate is the fork" and with the charter's AUDIT + DESIGN-only phase. The roadmap also resolves my prior C1 scope blocker well enough for Step 1 by pointing to conductor-plus-runtime-seams rather than full-runtime-product staffing.

Finding 1 -- Step exits are not verifiable enough. Current exits such as "you stop copy-pasting," "gates are unfakeable," "multi-model routing works and is scored," and "it all runs as one TUI app" are directionally useful but not acceptance criteria. Required edit: add measurable exit tests per step. Examples: Step 1 should name a relay fixture or dry-run where a relay is accepted only through the conductor, stamped FROM is system-owned, lint/form validation runs before delivery, and a gate email/outbox item is produced. Step 2 should require a failed observe-as-send case and a passing send case with conductor-captured evidence. Step 3 should require at least two providers, a recorded routing decision with justified deviation support, and benchmark output. Step 4 should require a single TUI process showing sessions plus inbox state without external multiplexer dependency.

Finding 2 -- Step 1 must include consumer-shaped schema seams, not just m-1/m-2 ownership. The roadmap says foundations lock first "with consumers sketched," which is the right correction, but it needs to bind the concrete consumers. Required edit: Step 1 design-lock criteria must include schema/interface sketches reviewed for m-3 observe fields, m-4 routing records, and m-6 human-gate/email fields before m-1/m-2 lock. Otherwise Step 1 can produce a clean relay store and form gate that cannot express observe-as-send, routing justification, or governance-vs-collaboration human routing without churn.

Finding 3 -- the email/comms boundary is too broad in Step 2. Step 1 already "fires email at gates"; Step 2 says full email client UX plus scheduler. That risks dragging product surface work forward before the TUI shell and before the minimal conductor has proven value. Required edit: split the comms language into a minimal Step-1/2 governance inbox/outbox/scheduler path versus the full email-client UX. I would keep "email-at-gates plus parked-lane wake" early, but defer polish/full-client behavior to Step 4 unless a concrete Step-2 acceptance test requires it.

Finding 4 -- egress/content-safety is too late if any earlier step sends external email. Step 6 includes egress and operational-safety gates, but Step 1/2 already send gate emails and flow through an email client. If the conductor publishes externally before Step 6, outbound content scanning cannot wait until packaging. Required edit: add a minimal fail-closed egress gate before the first external send, with full egress hardening still allowed in Step 6. If early email is local-only/outbox-only, say that explicitly.

Finding 5 -- routing should not wait entirely for provider adapters. Runtime execution of routing belongs in Step 3, but the routing record schema and policy hooks are consumers of m-2 and should be sketched during Step 1. Required edit: distinguish "routing record/form contract designed early" from "router executes models directly after adapters exist." This prevents Step 3 from discovering that the schema cannot represent role+model dispatch, static priors, or planner deviation justification.

Required edits:
1. Add per-step exit tests with observable pass/fail criteria.
2. Add Step-1 design-lock dependencies: m-3/m-4/m-6 consumer review of m-1/m-2 interfaces before lock.
3. Split minimal governance comms from full email-client UX.
4. Move a minimal egress/content-safety gate before any external email send, or mark early email local-only.
5. Pull routing record/schema design into the early consumer-sketch work while leaving runtime execution in Step 3.

No reroute needed. No implementation authority granted. Recursive org support remains correctly placed as a Step-5 product capability, not a manual Cycle-1 operating topology.

ACTIONS_GIT_REF: wrote reviewer relay at master/relays/master-roadmap/SITREP-orchestrator-reviewer-20260628-034942.md; docs workspace git status unavailable because cwd is not a git repo; pcode status --short is empty.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
