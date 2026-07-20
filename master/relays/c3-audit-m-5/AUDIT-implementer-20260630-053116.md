## m-5 Workflows & Archetypes - c3 focused audit (independent implementer pass)

ROLE: Implementer
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c3-audit-m-5
PARENT_DISPATCH_ID: c3-audit-m-5
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - read-only audit; concrete archetype values surfaced for c3 DESIGN only
FROM: m-5.implementer
TO: m-5.planner, master.orchestrator-planner
CC: master.orchestrator-reviewer, m-6.planner, m-6.implementer, operator
IN_REPLY_TO: c3-audit-m-5/AUDIT-orchestrator-planner-20260630-051950.md
BUNDLE_ID: m-5-workflows-archetypes
OWNER: m-5 (Workflows & Archetypes)
SUBJECT: m-5 implementer audit - promote existing gate/sensor patterns; design the tag-space, ceilings, templates, and human-mode seam in c3

Scope:
- This is an independent `m-5.implementer` audit pass for c3. It is not the pair reconcile and not a DESIGN lock.
- I treat the c1/c2 substrate as locked: `slot_in` is a per-work-record work-archetype, `seat_archetype` is per-assignment/spawn, and concrete values remain m-5 c3-owned.
- I inspected local prior art in v2.8.8, Codex, Claude Code, the v3 export, the c2 m-5 reserved proposals, and the locked m-3/m-4 design docs. Web search was attempted per dispatch, but no external source added a better cite than the checked-out local source trail.

## 4-bucket verdict

PRIMARY_BUCKET: recommended-next

still-open:
- The full v3 m-5 archetype system is not already built or locked. C2 only locked the two opaque axes and reserved concrete tag-space, invariant selection, default gate composition, template structures/lineup, and authority-ceiling semantics to m-5 c3; `master/ARCHITECTURE.md:180-205` and `master/ARCHITECTURE.md:231-232`.
- `master/domains/m-5-workflows-archetypes/README.md` is stale as a domain status surface: it still frames m-5 as a c2 narrow consumer lens and does not show the c3 full-domain audit/design surface. Orchestrator should refresh it during c3 reconcile/design prep, but this audit does not edit it.
- The concrete values below are therefore design inputs only. They must be grilled and locked in the m-5 DESIGN artifact, then exposed before m-6 binds surfaces.

already-closed:
- v2.8.8 already gives mature relay/gate primitives: file-first relays, phase/authority/evidence headers, role separation, pair lifecycle, design/plan/review/merge gates, and reviewer visibility. Promote those as the substrate, not as the archetype system; `extracted/agentic-dev-team-skills-v3-export/v2.8.8-release/v288-unzipped/agentic-dev-team-skills-v2.8.8/README.md:5-45`, `:79-101`, and `:137-154`.
- C2 already closed the integration seam: m-3 executes observe predicates keyed by opaque `slot_in`; m-4 records `seat_archetype` plus resolved `authority_ceiling` per assignment and owns template record mechanics; m-5 owns structures/lineup and concrete semantics; `master/ARCHITECTURE.md:131-164`.
- The c2 sensor integrity split is already reserved: answer content is advisory/self-reported and never gate-bearing; observable metadata such as tool-blocked/no-source-actions/one-turn can be observed; `master/ARCHITECTURE.md:196-205`.

product-overlapped:
- Codex collaboration-mode templates overlap the idea of behavior presets, but they are single-agent prompt/style modes, not topology-bearing workflow records; `references/codex/codex-rs/collaboration-mode-templates/src/lib.rs:1-4`, `references/codex/codex-rs/collaboration-mode-templates/templates/plan.md:1-20`, and `references/codex-notes.md:107-116`.
- Codex `agent-graph-store` overlaps persisted topology, but it is a mutable parent/child graph with open/closed status, not an append-only conductor lineage with seat archetypes, gate-sets, human modes, or ceiling semantics; `references/codex/codex-rs/agent-graph-store/src/store.rs:13-27`, `references/codex/codex-rs/agent-graph-store/src/types.rs:4-12`, and `references/codex/codex-rs/agent-graph-store/src/local.rs:31-60`.
- Claude Code overlaps strongly on agents-as-data and side-question sensors. Its `/btw` side question forks a separate lightweight one-turn agent, denies tools, and does not interrupt the main conversation, which is direct sensor prior art; `references/claude-code/src/utils/sideQuestion.ts:48-52` and `:60-96`. Its custom/built-in agent definitions are useful "templates as data" prior art, but not m-5's two-axis, gate-composed, human-mode-bearing archetype system.

recommended-next:
- Proceed to DESIGN with a minimal m-5-owned registry: concrete work tags (`slot_in` values), seat tags (`seat_archetype` values), invariant/gate/ceiling composition, and three shipped templates: T1 Solo, T2 Adversarial Pair, T3 Sensor.
- Do not ship a full conductor-plus-N-pairs or continuous research/QA swarm template in v3.0. Those are valid Step 5 extensions, but the c3 lock only needs the minimal structures that m-3, m-4, and m-6 must consume.
- Run the required design grill on the tag values, authority ceilings, and m-6 human-mode vocabulary before lock.

Duplicate / already-built gate:
- Promote v2.8.8 relay/gate primitives, Codex data/template and topology patterns, and Claude Code side-question mechanics.
- Do not rebuild those primitives, but also do not pretend they satisfy m-5. None provides the full v3 tuple: `(topology + gate-set + human-mode)`, concrete tag-space, per-archetype observe invariants, authority-ceiling-at-spawn, and sensor/actuator composition. The v3 export states that an expansion slot is exactly that shared-substrate preset; `extracted/agentic-dev-team-skills-v3-export/v3-design/v3-adaptive-routing-pillar.md:53-58`.

## Concrete tag-space proposal (surfaced, not locked)

Work-archetype axis: `slot_in` values, stamped at accepted work-record/dispatch classification, non-lane-writable.

| value | intended use | default observe invariant |
|---|---|---|
| `feature_extension` | additive user-visible capability | additive-only diff where feasible; existing suite green; no test edits unless test-addition explicitly declared |
| `refactor` | behavior-preserving structure change | suite green; test files unchanged by default; diff explains moved/renamed behavior |
| `cleanup` | dead-code/docs/config hygiene with no behavior intent | references eliminated or intentionally redirected; no new public behavior; focused suite/smoke green |
| `bugfix` | defect correction with a reproducible symptom | red-on-parent repro, green-on-fix where feasible; suite green; no unrelated feature drift |
| `migration` | schema/runtime/dependency/path migration | reversibility or rollback note; compatibility window; higher human-gate tier for production/data risk |
| `research_synthesis` | read-only investigation or design/audit synthesis | cited artifact/report present; no source actions; minority/uncertainty preserved when applicable |
| `qa_review` | read-only or low-mutation verification/triage | finding evidence/repro or explicit no-finding report; dedup/novelty where applicable; no source actions unless explicitly escalated |

Notes:
- These are named presets over a tag vector, not a taxonomy tree. The export argues gates compose from mutation properties and named archetypes are presets; `extracted/agentic-dev-team-skills-v3-export/v3-design/v3-adaptive-routing-pillar.md:53-56`.
- For Step 1, the registry can be declared in docs/forms and enforced through m-3's existing observe-as-send hook. Some richer checks such as differential red/green or reversibility may be operator-gated descriptors until executable-claim infrastructure exists; this is a later runtime dependency, not a c3 blocker.

Seat-archetype axis: `seat_archetype` values, stamped per assignment/spawn, non-lane-writable, consumed by m-4 routing/ceiling.

| value | authority ceiling | routing prior | m-6 surface posture |
|---|---|---|---|
| `solo_worker` | may edit only after valid dispatch; never self-merge | coding-capable model bucket for mutating work | local work lane with normal checkpoint gates |
| `planner` | read/design/dispatch-authoring as delegated; no source edits | high-reasoning planning bucket | collaboration lane, design checkpoint friendly |
| `implementer` | source edits only inside dispatch authority; no merge without fold/authorization | coding-capable bucket | work lane, evidence-first send gate |
| `reviewer` | read-only review/audit/fold unless explicitly delegated otherwise | adversarial/reasoning bucket | review lane, findings-first output |
| `sensor` | read-only, tool-blocked, single-turn by default; cannot become writer | fast-cheap by default unless domain needs stronger lens | side surface, non-interrupting |
| `actuator` | mutating team/seat class; edits only through dispatched workflow, gates heavy | coding-capable or pair topology routing | work lane with operator gates for authority raises |
| `orchestrator_lead` | routing/decomposition/reconcile authority only; no direct implementation | high-reasoning orchestration bucket | dashboard/reconcile lane |

Notes:
- `sensor` and `actuator` are structural classes, not role names only. A `solo_worker` or `implementer` can be an actuator seat; a `research_synthesis` or `qa_review` work item can be staffed by one or more sensors.
- `actuator` should be treated carefully in DESIGN: it may be better represented as a derived ceiling class rather than a literal seat value if the grill finds it duplicates `solo_worker`/`implementer`.

## Per-archetype invariant composition

Composition rule:
- `slot_in` selects the work-record invariant family that m-3 observes.
- `seat_archetype` selects the spawn-time authority ceiling and m-4 routing prior.
- The template combines both axes with a default gate-set and human-mode. The conductor records both; the lane does not rewrite either.

Work invariant map:
- `feature_extension`: base source-change predicate + additive-diff check + suite green. HUMAN_GATE raises only for production/data/egress or policy-sensitive surfaces.
- `refactor`: base source-change predicate + suite green + test-files-unchanged default. Any test edit requires explicit declaration and higher review.
- `cleanup`: base source-change predicate + reference/consumer check (`find-references=0` or equivalent evidence) + no behavior claim. Minimal gate unless deleting public API/config.
- `bugfix`: repro artifact preferred; parent SHA fails and fixed SHA passes when executable. If differential run is unavailable, achieved evidence degrades and must be explicit.
- `migration`: rollback/reversibility note + compatibility evidence + broader smoke/integration check; HUMAN_GATE tends to A for irreversible/data-affecting moves.
- `research_synthesis`: report artifact with citations, scope, uncertainties, and no source actions. Output may feed a later actuator dispatch but cannot satisfy implementation gates.
- `qa_review`: finding/no-finding artifact with repro/evidence, false-positive triage, and dedup/novelty. A real fix routes to a separate actuator spawn.

Seat/ceiling map:
- `sensor`: read-only; no tools by default; one turn; advisory content only. It can recommend or draft a dispatch request, but cannot stamp or perform write authority.
- `planner` / `orchestrator_lead`: may author design/routing/reconcile relays within their addressed authority; cannot turn planning into implementation.
- `reviewer`: read-only and findings-first. May block or recommend folds through the existing relay protocol, not through private vetoes.
- `solo_worker` / `implementer`: write-capable only after an accepted dispatch; subject to observe-as-send and fold/merge gates.
- `actuator`: mutating structural class. Read-only-to-write transition is a hard human-gated boundary; sensors emit work into actuators rather than upgrading themselves.

Evidence:
- m-3 already lists candidate done-predicate add-ons for extension/refactor/cleanup/bugfix/migration and owns execution of the check registry; `master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:75-99`.
- The export states read-only to write is a hard human-gated boundary and sensors emit into a separately spawned actuator loop; `extracted/agentic-dev-team-skills-v3-export/v3-design/v3-adaptive-routing-pillar.md:56-58`.
- m-4 already fails closed when a requested staffing violates the resolved ceiling; `master/ARCHITECTURE.md:157-164`.

## Template lineup (T1/T2/T3)

T1 - Solo:
- topology: one addressed working seat; no peer seat.
- seats: `solo_worker` or role-specific `implementer`; optional `planner` only as upstream author, not part of the template.
- panes/surfaces: one work pane/lane in Step 1 using existing host runtime or multiplexer.
- gate-set: selected by `slot_in`; observe-as-send always active; merge/authority raises remain separate.
- read-onlyness: usually mutating, but bounded by dispatch authority and selected ceiling.
- default human-mode: `work_checkpoint` for normal progress and `operator_gate` for A-category raises.

T2 - Adversarial Pair:
- topology: planner + implementer, with review/fold lineage.
- seats: `planner` plus `implementer`; optional reviewer visibility remains orchestrator/VP by protocol, not a third pair seat.
- panes/surfaces: two work panes/lanes in existing runtime.
- gate-set: design/plan/review/fold gates as phase requires; work invariant still comes from `slot_in`.
- read-onlyness: planner/reviewer side is read/design/review; implementer side mutates only after dispatch.
- default human-mode: `review_checkpoint` for pair decisions; `operator_gate` for scope/authority changes.

T3 - Sensor:
- topology: one side-question/read-only fork; not in the lane's delivery lineage.
- seats: `sensor`.
- panes/surfaces: m-6-owned side surface, dismissible/non-interrupting; no work pane.
- gate-set: no delivery gate; report-only observe predicate; no source/test/prototype edits.
- read-onlyness: strict read-only, tool-blocked, single-turn.
- default human-mode: `side_surface`; no gate-bearing interruption.

Deferred:
- `T4 conductor_team` / orchestrator-plus-N-pairs remains Step 5 or a later c3 design extension. It would introduce nested topology, recursion, scheduler, and capstone integration pressure not needed to bind the c2 opaque atoms.

Evidence:
- C2 already reserved T1 Solo, T2 Adversarial Pair, and T3 Sensor as the proposed v3.0 lineup while deferring the conductor template; `master/ARCHITECTURE.md:196-205`.
- m-4 owns the routing-template record mechanism and m-5 owns structures/lineup; pane spawn rides existing tmux/zellij/OS terminal in Step 1; `master/ARCHITECTURE.md:162-164`.
- v2.8.8 proves the adversarial-pair relay lifecycle works as a protocol primitive; `extracted/agentic-dev-team-skills-v3-export/v2.8.8-release/v288-unzipped/agentic-dev-team-skills-v2.8.8/README.md:137-154`.

## Sensor / actuator design

Sensor:
- authority ceiling: read-only, no dispatch/implementation/merge authority.
- tool ceiling: no tools by default for side-question; if later research sensors get read tools, that is a distinct template with observed tool-use metadata and a bounded episode.
- turn ceiling: one response for side-question; bounded episodes for future research/QA sensors.
- context: point-in-time read-only context snapshot or prompt-cache fork.
- lineage: advisory side output, not gate-bearing work lineage.
- integrity: answer content is `self_reported`/advisory; runtime-observable facts such as one-turn closure, no tool use, no source actions, and fork metadata are `observed` where exposed.
- routing: default `fast-cheap`; can route stronger by domain if m-4 prior says the question needs it.

Actuator:
- authority ceiling: mutating only through a stamped workflow/template and accepted dispatch.
- tool ceiling: normal source/test tools allowed only inside dispatch scope.
- gate-set: heavy enough for the selected `slot_in`; observe-as-send required.
- lineage: owns work delivery and can satisfy implementation gates.
- sensor composition: sensors may emit candidate work items or decision briefs into a human/conductor queue; a human or authorized orchestrator spawns an actuator. A sensor never upgrades in place.

Interjection home:
- `steer`: m-6/runtime surface for user course correction; m-5 only declares whether a template accepts steer messages without a respawn.
- `side_question`: m-5 `sensor` archetype plus m-4 routing; m-6 owns the side surface.
- `interrupt`: m-6/runtime cancellation/hold surface; m-5 declares whether the template is interruptible and what happens to its episode state.

Evidence:
- Roadmap defines side-question as read-only, tool-blocked, single-turn, cheap/fast, and separate surface; `ROADMAP.md:32-46`.
- Claude Code implements the same basic sensor shape through `runSideQuestion`: prompt-cache fork, no tools, one turn, separate from the main conversation; `references/claude-code/src/utils/sideQuestion.ts:48-96`.
- C2 records side-question as a forward interjection requirement with m-6 surface, runtime mechanism, and m-5/m-4 archetype/routing ownership; `master/ARCHITECTURE.md:216-220`.

## m-6 boundary contract

m-5 exposes this vocabulary; m-6 binds UI/email/meeting/scheduler behavior to it.

Human-mode values to declare in DESIGN:
- `quiet_local`: no expected human interaction beyond normal final/result relay; appropriate for low-risk solo work or passive sensors.
- `work_checkpoint`: local in-session checkpoint for progress, scope, or evidence review.
- `review_checkpoint`: adversarial pair/design-review checkpoint where discussion is part of the value.
- `operator_gate`: blocking governance verdict; should map to the prepared decision-brief style when asynchronous.
- `away_bridge_eligible`: A-category blocking gate may bridge to email/scheduler when operator is away and policy allows it.
- `side_surface`: non-blocking sensor answer surface, dismissible and not gate-bearing.
- `hold_and_resummon`: pause until operator attention is available; used when neither local checkpoint nor away bridge is appropriate.

Seam A - template to human surface:
- m-5 declares `human_mode`, interruptibility, and whether output is gate-bearing.
- m-6 decides how that appears in TUI/email/meeting/scheduler surfaces.
- m-5 should not specify m-6 layout, copy, notification buckets, or scheduler implementation.

Seam B - interjection:
- m-5 declares semantic classes (`steer`, `side_question`, `interrupt`) and the allowed effect per template.
- m-6 owns the input affordance and response surface.
- Runtime/conductor owns actual injection, cancellation, fork, and state handling.

Ordering requirement:
- m-5 DESIGN must publish this vocabulary before m-6 locks surface behavior. The c3 decomp guardrail explicitly says m-5 declares human-mode vocabulary and archetype/sensor semantics before m-6 binds; `master/relays/c3-decomp/SITREP-orchestrator-reviewer-20260630-051448.md:41-47`.

## Reject / narrow gates

- Reject locking concrete archetype values in AUDIT. This file only surfaces inputs for c3 DESIGN.
- Reject any m-2 `required_when` or `visible_when` rules over these concrete values in c3. C2 reserved opaque atoms only and explicitly avoided an m-2 micro-fold; `master/ARCHITECTURE.md:189-194`.
- Reject sensor-to-actuator in-place upgrades. Read-only to write is a hard human-gated boundary.
- Reject a full conductor/N-pair template for v3.0 unless the orchestrator explicitly expands c3 scope. It has no immediate downstream consumer needed to close m-3/m-4/m-6 seams.
- Narrow any invariant that cannot ride Step 1 into an evidence-target/degradation note plus later runtime dependency. Do not make standalone runtime enforcement a c3 design blocker.

## Operator / orchestrator judgment items

- No blocking operator decision from this implementer audit.
- Orchestrator should refresh or queue refresh of `master/domains/m-5-workflows-archetypes/README.md`; it is stale against the c3 audit/design scope.
- Pair reconcile should decide whether `actuator` is a literal `seat_archetype` value or a derived ceiling class over `solo_worker`/`implementer`.
- Pair reconcile should decide whether `research_synthesis` and `qa_review` ship as c3 work-archetypes now or are marked Step 5 extension values. My recommendation is to surface them now but lock only if m-6 and orchestrator need them for the final Step-0 architecture-of-record.

## Evidence levels

- E1: dispatch inspected: `master/relays/c3-audit-m-5/AUDIT-orchestrator-planner-20260630-051950.md`.
- E1: c2 lock inspected: `master/ARCHITECTURE.md:120-235`.
- E1: m-3/m-4 locked designs inspected: `master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md`; `master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md`.
- E1: prior-art source inspected: v2.8.8 release tree, Codex collaboration templates and graph store, Claude Code side-question and agent definitions, v3 adaptive-routing pillar.
- E1: read-only status checked. Harness root is not a git repo; `pcode/` is the code repo and was clean before this relay write.

ACTIONS_GIT_REF: wrote this read-only audit relay only: `master/relays/c3-audit-m-5/AUDIT-implementer-20260630-053116.md`; no source/test/pcode edits claimed; cwd is not a git repo; `git -C pcode status --short` returned clean output.

FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (`git status --short` exits 128; docs workspace, `pcode/` is the code repo)

Tests / verification:
- E2: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c3-audit-m-5/AUDIT-implementer-20260630-053116.md` passed.
- E1: `git -C pcode status --short` returned clean output.
- E1: `git status --short` at harness root exits 128 (`fatal: not a git repository...`), matching the expected structured unavailable form.
