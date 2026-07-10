## m-5 Workflows & Archetypes — c2 narrow consumer-review (independent implementer pass)

ROLE: Implementer
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c2-consumer-review-m-5
PARENT_DISPATCH_ID: c2-consumer-review-m-5
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — read-only consumer-review + bounded proposals; no c3 archetype-system lock
FROM: m-5.implementer
TO: m-5.planner, master.orchestrator-planner
CC: master.orchestrator-reviewer, m-3.planner, m-4.planner, operator
IN_REPLY_TO: c2-consumer-review-m-5/AUDIT-orchestrator-planner-20260629-212435.md
BUNDLE_ID: m-5-workflows-archetypes
OWNER: m-5 (Workflows & Archetypes), consumer lens
SUBJECT: m-5 implementer audit — seam fits with provenance wording fix; 3-template lineup acceptable; sensor spec needs integrity-label split

Scope:
- This is an independent adversarial `m-5.implementer` pass on the bounded c2 consumer review. It is not a full m-5 design.
- I read the m-5 planner pass after the source review to answer its challenge questions and prepare the later pair reconcile.
- I do not close concrete tag-space, invariant selection, default gate composition, full template semantics, or authority-ceiling semantics; those remain m-5 c3-owned.

## 4-bucket verdict

PRIMARY_BUCKET: recommended-next
- still-open: the c2 lock still needs an m-5 pair reconcile that records the bounded seam disposition and carries the planner/implementer deltas below.
- already-closed: the interface can be locked as an opaque-key seam without a c2 re-cut. The sources already reserve `slot_in`/archetype semantics to m-5 and keep m-3/m-4 to consuming mechanisms: m-3 `slot_in` selects invariant sets while m-5 owns tag-space/invariant selection; m-4 takes an opaque archetype tag vector and reserves ceiling semantics to m-5.
- product-overlapped: no existing system fully covers the m-5 work. Codex mode templates are single-agent behavioral presets, not multi-seat team topologies; jcode/Claude Code provide side-question prior art but not the conductor archetype contract.
- recommended-next: reconcile as `FITS-with-two-folds`: fold a two-axis archetype reservation plus a recorded per-assignment seat-archetype/ceiling home. Revise the planner's F1 wording from "spawn-time binding" to "conductor-owned classification at work-record acceptance/dispatch"; split the sensor answer's advisory integrity from the observable record/tool/turn facts.

Duplicate / already-built gate:
- No full m-5 design artifact exists under `master/domains/m-5-workflows-archetypes/`; the only domain file is the c2 charter/README.
- Prior c1/c2 relay trail reserves the seam and scopes this review; it does not already specify the m-5 template lineup or sensor archetype.
- `references/codex-notes.md` validates template-as-data but explicitly distinguishes Codex single-agent behavior presets from our multi-agent topology/authority-ceiling layer.

## Deliverable 1 — seam-fit verdict

Verdict: **FITS, with two c2 lock folds and one wording correction before pair reconcile.** The opaque atom/vector interface carries enough for c3 m-5 to fill semantics without cutting m-3 or m-4 again.

Evidence:
- m-2/c1 already made `slot_in` a reserved opaque atom with no concrete values; `master/ARCHITECTURE.md:52-56`.
- m-3 consumes the tag only as a selector for predicate add-ons; it owns execution/result shape while m-5 owns tag-space and invariant selection; `master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:81-99`.
- m-4 consumes an opaque archetype tag vector as `route_dispatch()` input and fail-closes on unavailable routes; `master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md:141-155`.
- m-4 explicitly gives template structures/lineup and concrete tag/ceiling semantics to m-5; `master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md:257-271` and `:313-317`.
- The c2 guardrail approves only narrow consumer review and preserves m-5 ownership for c3; `master/relays/c2-lock-prep/RECONCILE-orchestrator-reviewer-20260629-212213.md:41-47`.

Fold A — two opaque axes:
- Concur with m-5.planner's F3: keep `slot_in` as the work-record archetype used by m-3 done-predicates, and add/record a distinct `seat_archetype` (or equivalent tag in the m-4 archetype vector) for spawn-time authority ceiling and routing recommendation.
- Reason: m-3's examples are work-shape predicates (`extension`, `refactor`, `bugfix`, `migration`), while m-4's sensor example is a seat authority ceiling. One seat can perform multiple work archetypes over time; one work item can be staffed by different seat archetypes. Treating both as one tag would force c3 to overload one atom.

Fold B — recorded home for seat archetype / resolved ceiling:
- Concur with m-5.planner's F2. `routing_assignments` currently records seat, role, task tag, bucket/model choice, deviation, and pin mode, but not the seat archetype or resolved ceiling; `master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md:197-206`.
- For template spawns, `template_ref` can carry the seat structure. For hand-authored or one-off sensor spawns, the routing record still needs replayable proof of the ceiling that made write authority unavailable.
- c2 lock should ask m-4 either to record opaque `seat_archetype`/resolved `authority_ceiling` per assignment, or to explicitly require all archetype-bearing spawns at Step-1 to go through templates. I prefer the per-assignment field because it is replay-complete and does not make templates mandatory for simple one-off spawns.

Wording correction to m-5.planner F1:
- I agree the work-archetype must not be lane-writable. I do **not** agree that `slot_in` should be described as "conductor-stamped from the seat's spawn-time archetype binding" for the work-axis.
- Better wording: `slot_in` is conductor-owned and stamped from the accepted work record / dispatch classification, using an m-5-owned classifier or operator/planner selection path; a lane cannot rewrite it after acceptance. The `seat_archetype` is spawn-time. The `work_archetype` is per-work-record.
- Why this matters: a long-lived implementer seat may validly perform a bugfix, then a refactor, then a migration. If `slot_in` is fixed at seat spawn, either long-lived seats become impossible or the tag stops meaning work shape. The tamper-resistance property comes from conductor ownership and immutable acceptance, not necessarily from spawn-time.

Answer to planner challenge questions:
- F3/two axes: concur. It is not over-splitting; it prevents mixing per-seat authority with per-record evidence predicates.
- F1/provenance: partially concur. It must be conductor-owned and non-lane-writable, but not necessarily spawn-derived for the work axis.
- Actual interface gap: no gap in the opaque-key shape; only the record home/provenance wording must be folded now.

## Deliverable 2 — Step-1 routing-template structures + 1-3 shipped lineup

Recommended Step-1 lineup: **Solo, Adversarial Pair, Sensor**.

| template | structure | panes | gate-set | read-only-ness | m-5 c2 disposition |
|---|---|---:|---|---|---|
| Solo | one working seat; no peer topology | 1 | standard phase gates for the selected authority | configurable, usually write-not-merge | ship as the degenerate baseline |
| Adversarial Pair | planner + implementer with review lineage | 2 | design-review / plan-review where phase requires | mixed: planner/reviewer read-lean, implementer write-capable after dispatch | ship because it is the proven team primitive in this repo |
| Sensor | one side-question seat | 1 side surface | no work-delivery gate; observe/report only | strictly read-only, no tools, one turn | ship because the roadmap already makes side-question first-class |

Template structure, surfaced only:
- `template_id`
- `seats[]` with seat name, role, `seat_archetype`, pane/surface intent, and m-4 `model_slot`/pin information
- topology/lineage expectations
- `gate_set` by gate category
- `human_mode`/surface posture for m-6
- pane layout intent for conductor-core

Evidence:
- GL-4 requires 1-3 selectable templates and splits ownership: m-4 mechanism, m-5 structures/lineup, conductor-core pane spawn; `master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md:50-55`.
- m-4 treats a template as reusable routing assignments and requires no-bypass deviation recording; `master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md:233-260`.
- Roadmap Step 1 rides existing runtimes; Step 4 owns the standalone TUI, so pane spawn should remain an existing-runtime/multiplexer action at Step-1; `ROADMAP.md:57-65` and `:82-86`.
- Codex supports named mode presets as data, but its presets are single-agent behavior prompts; our topology/authority-ceiling layer remains m-5's contribution; `references/codex-notes.md:107-116` and `:141-144`.

Challenge / narrow gate:
- Do not ship the full orchestrator + N-pairs conductor template at Step-1. It is the product's eventual marquee topology, but it would pull c3 nested/workflow semantics forward. Solo + Pair + Sensor covers the axes this lock actually needs: minimal work, adversarial review, and read-only side-question.

## Deliverable 3 — side-question sensor archetype

Spec, bounded to c2:
- seat archetype: `sensor`
- authority ceiling: read-only; no write/edit/merge/dispatch authority; m-4 fail-closes to `routing_unavailable` if asked to staff it otherwise.
- tool ceiling: none at Step-1, including no read-only file/search tools.
- turn ceiling: one response, then the seat closes.
- context: point-in-time read-only fork of the lane context/prompt cache; no future lane state; not in the lane's work lineage.
- surface: separate m-6-owned side surface; never injected into the lane's inbox as gate-bearing work.
- routing: m-4 `fast-cheap` bucket by default because it is single-turn, tool-blocked, read-only advisory work.
- observe profile: read-only/report base predicate; conductor can observe no unauthorized source actions, no tool use, and one-turn closure where the runtime exposes those facts.
- integrity label: the answer content is advisory/self-reported judgment over a snapshot; the observable metadata may still be `observed`. Do not collapse the whole record to `self_reported` merely because free-text judgment is unverified.

Evidence:
- Roadmap defines side-question as a read-only, tool-blocked, single-turn fork routed cheap/fast and answered on a separate surface; `ROADMAP.md:32-46`.
- jcode notes identify the desired `/btw` shape as a read-only, tool-blocked, single-turn forked archetype sharing lane context and writing separately; `references/jcode-ux-notes.md:47-53` and `:62-77`.
- m-4's authority-ceiling cap refuses forbidden staffing and cites the read-only sensor example; `master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md:264-272`.
- m-3 read-only report predicate is "allowed relay artifact present + no unauthorized source actions"; `master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:81-88`.
- m-3 says `record_integrity` ranges over observable fields and does not claim free-text judgment is independently verified; `master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:102-107`.

Challenge to planner sensor text:
- I do not accept blanket `record_integrity: self_reported` as the default. The answer's substantive advice is self-reported, but tool-blocked/no-source-actions/one-turn closure are observable system facts if the conductor/runtime exposes them.
- Reconcile wording should say: "sensor answer content is advisory and not gate-bearing; observed metadata is stamped per m-3 rules; no sensor answer can satisfy or override the lane's delivery gate."

## Operator / orchestrator judgment items

- No blocking operator decision from this implementer pass.
- Orchestrator/VP should record the two-axis reservation at c2 lock: `slot_in` work-archetype per work record for m-3; `seat_archetype` per assignment/spawn for m-4 ceiling/routing. Concrete values remain c3.
- Orchestrator/VP should pick the m-4 record-home resolution for `seat_archetype`/`authority_ceiling`: per-assignment field preferred, all-spawns-through-template acceptable but more constraining.

## Risks

- If the two axes are conflated, c3 must recut either m-3 predicates or m-4 ceilings.
- If `slot_in` is described as spawn-derived, the design may accidentally prohibit long-lived seats from doing multiple work types.
- If sensor integrity is blanket `self_reported`, we lose useful observed metadata; if it is blanket `observed`, we overclaim on free-text judgment. Split the labels.
- If the conductor template ships now, c2 crosses into the full m-5 archetype-system design that the VP guardrail forbids.

ACTIONS_GIT_REF: wrote this read-only audit relay only: `master/relays/c2-consumer-review-m-5/AUDIT-implementer-20260629-213807.md`; no source/test/pcode edits claimed; cwd is not a git repo; `git -C pcode status --short` returned clean output.

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (`git status --short` exits 128; docs workspace, `pcode/` is the code repo)

Tests / verification:
- E1: dispatch inspected: `master/relays/c2-consumer-review-m-5/AUDIT-orchestrator-planner-20260629-212435.md`.
- E1: source design docs and cited references inspected.
- E1: duplicate/already-built gate checked with `rg` over `master`, `extracted`, and `references`; no full m-5 design exists.
- E2: `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c2-consumer-review-m-5/AUDIT-planner-20260629-213422.md` passed for the planner input relay.
- E2: `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c2-consumer-review-m-5/AUDIT-implementer-20260629-213807.md` passed.
