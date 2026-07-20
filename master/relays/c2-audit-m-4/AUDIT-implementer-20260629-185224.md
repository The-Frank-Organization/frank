ROLE: Implementer
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c2-audit-m-4
PARENT_DISPATCH_ID: c2-audit-m-4
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: m-4.implementer
TO: m-4.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-4-routing-policy
OWNER: m-4

# m-4 Implementer Independent Audit - Routing & Policy

## Verdict

PRIMARY_BUCKET: still-open

still-open: v3 still needs a first-class m-4 routing/policy primitive: a separate seat-stamped routing relay carrying role+model assignment, capability-prior basis, deviation justification, and a benchmark-feedback handle. v2.8.8 and the references provide coordination/runtime/model-selection machinery, but not an auditable governance decision record.

already-closed: reusable mechanics exist. v2.8.8 already has relay addressing, lineage, evidence discipline, and role contracts; jcode already has structured provider/model route identity and execution; Claude Code already has agent-type plus optional model override and subagent inheritance; agent-scripts already has operational decision-brief and model/provider preflight patterns. These are primitives to borrow, not a duplicate closure.

product-overlapped: jcode overlaps the runtime execution half most strongly. Its `MultiProvider`, `RouteSelection`, `RuntimeKey`, and wire-level provider/model attribution should inform Step 3 execution, but jcode does not close Step 1 governance-record design.

recommended-next: design m-4 as a thin governance layer over existing runtime surfaces: `route_dispatch()` writes an accepted R2 routing relay in Step 1; Step 3 later executes that relay through provider adapters without changing the record shape.

## Named Gap

IMPLICIT_ROUTING_GAP: surveyed systems choose models/roles through config, flags, runbook convention, runtime fallback, or coordinator actions, but they do not create a separate seat-stamped, parented, justified routing decision that downstream dispatches consume as provenance.

How the routing relay closes it: the router emits one accepted routing relay per routed dispatch. The relay is stamped by m-1 (`FROM`/seat identity), typed by m-2 (`routing` record-kind and FieldSpec), and consumed by downstream dispatches only as provenance/bookkeeping. The model assignment stays payload, never a schema-gate predicate, preserving R2 (`master/ARCHITECTURE.md:66-69`) and the identity != authority boundary (`master/ARCHITECTURE.md:74-77`).

## Evidence Findings

1. Locked local contract: R2 already requires a separate seat-stamped routing relay and explicitly bars `model_*` gate predicates (`master/ARCHITECTURE.md:66-69`). m-1 owns who; m-4 owns what stamped seats may do (`master/ARCHITECTURE.md:74-77`). Routing is category-B and model is payload (`master/ARCHITECTURE.md:94-97`).

2. Roadmap timing: routing records are designed early during Step 1, while provider adapters and router execution land in Step 3 (`ROADMAP.md:23-26`, `ROADMAP.md:59-64`, `ROADMAP.md:91-92`). Therefore the audit should not require runtime execution to lock the record/API surface.

3. v3 seed: the exported pillar locks altitude B, a 3-staged policy, and benchmark feedback as a v3.1 hook (`extracted/agentic-dev-team-skills-v3-export/v3-design/v3-adaptive-routing-pillar.md:16-19`). It later resolves model identity as non-trust-bearing payload/bookkeeping and names router as a conductor function (`extracted/agentic-dev-team-skills-v3-export/v3-design/v3-adaptive-routing-pillar.md:33-35`).

4. v2.8.8 gap: the released package is a coordination protocol, not an orchestration runtime (`extracted/agentic-dev-team-skills-v3-export/v2.8.8-release/v288-unzipped/agentic-dev-team-skills-v2.8.8/README.md:1-5`). It explicitly says roles are contracts, not model identities (`extracted/agentic-dev-team-skills-v3-export/v2.8.8-release/v288-unzipped/agentic-dev-team-skills-v2.8.8/README.md:96-100`, `extracted/agentic-dev-team-skills-v3-export/v2.8.8-release/v288-unzipped/agentic-dev-team-skills-v2.8.8/README.md:139-146`). Its route-like fields are relay `TO`/`CC`, lineage, and dispatch authority, not role+model routing.

5. jcode runtime prior art: jcode's `MultiProvider` multiplexes concrete providers (`references/jcode/crates/jcode-base/src/provider/mod.rs:301-336`), and `RouteSelection` defines structured model/runtime identity (`references/jcode/crates/jcode-provider-core/src/lib.rs:617-630`). `set_route_selection()` executes a route by converting it to a routed model spec (`references/jcode/crates/jcode-base/src/provider/mod.rs:1500-1509`). This is strong runtime prior art.

6. jcode governance gap: jcode's wire API has `SetModel`, `SetRoute`, and `RunSubagent` with optional model, but no stamped actor, policy stage, capability prior, justified deviation, benchmark handle, or R2 parent reference (`references/jcode/crates/jcode-protocol/src/wire.rs:189-207`, `references/jcode/crates/jcode-protocol/src/wire.rs:217-225`). Role assignment carries only target session and role (`references/jcode/crates/jcode-protocol/src/wire.rs:459-466`) and is operationally applied by coordinator logic (`references/jcode/crates/jcode-app-core/src/server/comm_control.rs:669-733`, `references/jcode/crates/jcode-app-core/src/server/comm_control.rs:766-818`). Its history/model events attribute provider/model (`references/jcode/crates/jcode-protocol/src/wire.rs:919-934`, `references/jcode/crates/jcode-protocol/src/wire.rs:1061-1069`) but attribution is not a governance decision.

7. Claude Code runtime prior art: main model selection is precedence over runtime override, startup flag, env/settings, and defaults (`references/claude-code/src/utils/model/model.ts:61-98`). Subagents expose `subagent_type` plus optional `model` override (`references/claude-code/src/tools/AgentTool/AgentTool.tsx:81-87`) and resolve tool-specified model before agent model/inheritance (`references/claude-code/src/utils/model/agent.ts:37-88`, `references/claude-code/src/tools/AgentTool/runAgent.ts:340-345`). Logging records selected agent type/model (`references/claude-code/src/tools/AgentTool/AgentTool.tsx:417-424`). This is routing/telemetry, not a justified R2 record.

8. agent-scripts operational prior art: maintainer orchestration says not to request a custom model and to inherit platform default (`references/agent-scripts/skills/maintainer-orchestrator/SKILL.md:24-29`), while owner decision briefs require explaining why a decision is needed and exact choices (`references/agent-scripts/skills/maintainer-orchestrator/SKILL.md:53-69`). Oracle exposes explicit model/provider route/preflight knobs (`references/agent-scripts/skills/oracle/SKILL.md:90-97`). These are useful policy and operator-load patterns, not routing-record semantics.

## Design Recommendation

Minimal primitive: `route_dispatch(dispatch_id, assignments[], policy_context) -> routing_relay_id`.

Router API surface:
- input: target dispatch id, workflow/archetype tag vector if present, seats/roles to assign, task domain hints, evidence target, budget/latency/privacy constraints, available runtime routes, and optional prior routing relay for reroute.
- output: accepted routing relay id and a read-only projection for the dispatch header/body.
- failure mode: no acceptable route emits `human_decision_required` or `routing_unavailable`; it must not silently fall back to default model.

Capability-prior representation:
- versioned policy table keyed by `role`, `task_tag`, `evidence_target`, and optionally `archetype_tag`.
- entries contain ordered candidate route families/model classes plus reason codes, not just prose.
- table is configuration/policy data; the routing relay snapshots the used policy version and matched prior row so later benchmark results can score the decision.

Routing-record FieldSpec shape:
- `record_kind`: routing
- `target_dispatch_id`: system/parent-picked dispatch id
- `assignments`: repeatable group with `seat`, `role`, `model_class_or_model`, `runtime_route_key`, `provider_family`, `capability_prior_id`, `policy_stage`, `deviation_status`, `justified_deviation`
- `constraints`: budget/latency/privacy/context/evidence target fields used by the policy
- `benchmark_feedback_handle`: reserved id for v3.1 outcome feedback
- `accepted_semantics`: accepted routing relay becomes a `parent_picker` candidate for the routed dispatch; downstream dispatch may reference it as provenance, but no `model_*` field becomes a gate input.

Three-stage policy:
1. Capability priors choose the default route assignment from the versioned policy table.
2. Justified deviation permits planner/router override only with a required free-text reason and machine-readable reason code.
3. Benchmark/outcome feedback in v3.1 consumes m-3 observed evidence and updates/prioritizes future priors; it is a forward hook now, not Step-0 closure.

## Consumer Boundary Contract

m-3 Observation & Evidence:
- m-4 exposes routing relays as evidence-addressable records. m-3 may cite whether a routing decision existed, was accepted, and later whether the routed run met observed outcomes.
- m-4 consumes m-3 observed evidence only through the benchmark feedback handle; m-4 does not own m-3 observe internals.

m-5 Workflows & Archetypes:
- m-5 owns archetype tags and authority ceiling at spawn; m-4 consumes those tags as policy inputs.
- DESIGN must settle which tag fields are stable enough for the policy table. c2 cannot lock m-4 without a m-5 seam disposition because archetype tags parameterize routing priors.

m-6 Human Surface & Scheduler:
- m-4 emits routing category/status fields for m-6 gate bucketing. Default routing is category-B/orchestrator-absorbed, but `human_decision_required`, `routing_unavailable`, and unclassified `other` must surface to the human path consistent with the c1 gate-category contract.
- m-6 does not choose models; it surfaces routing decisions and blocked policy states.

## Step 1 / Step 3 Split

The record can ride existing runtimes in Step 1: write the routing relay, then use the selected model/route manually or through host-native controls. Execution is explicitly Step 3 per `ROADMAP.md:59-64`. This means no execution dependency blocks DESIGN lock, provided the record carries enough route identity to map to future provider adapters.

## Operator Items

None for AUDIT. The m-5 seam is not an operator item yet; it is a required DESIGN coordination item because tag semantics are owned by m-5.

## Coordination Relays Sent

None. No blocking m-1/m-2 gap found; R2 is preserved rather than reopened.

## Evidence Levels

E1: local file/line citations above. Web searches were attempted for external routing prior art, but this report relies on local/bundled sources and independently spawned read-only explorer lenses for auditable citations.

## Actions

No source changes, no branches, no commits, no PRs. Relay artifact only.

ACTIONS_GIT_REF: wrote relay `master/relays/c2-audit-m-4/AUDIT-implementer-20260629-185224.md` and appended `master/relays/INDEX.md`; docs workspace is not a git repo, so git status is unavailable.

FINAL_GIT_STATUS_SHORT: unavailable - fatal: not a git repository (or any of the parent directories): .git
