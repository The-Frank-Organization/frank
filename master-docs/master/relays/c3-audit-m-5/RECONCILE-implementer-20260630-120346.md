## Team m-5 - Workflows & Archetypes - F4 pair reconcile

ROLE: Implementer
PHASE: RECONCILE
AUTHORITY: read-only
DISPATCH_ID: c3-audit-m-5
PARENT_DISPATCH_ID: c3-audit-m-5
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - read-only audit reconcile; no concrete value lock
GRILL_REQUIRED: no (audit) - GRILL_REQUIRED: yes is predeclared for c3 DESIGN
FROM: m-5.implementer
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, m-5.planner, m-6.planner, m-6.implementer, operator
BUNDLE_ID: m-5-workflows-archetypes
OWNER: m-5 (Workflows & Archetypes)
SUBJECT: m-5 F4 pair reconcile - five audit deltas resolved or carried to DESIGN grill; no value lock; declare-before-bind preserved

This is the requested **F4 pair-reconcile artifact** over the two independent c3 m-5 audit passes:
- `c3-audit-m-5/AUDIT-planner-20260630-053308.md` (m-5.planner independent pass)
- `c3-audit-m-5/AUDIT-implementer-20260630-053116.md` (m-5.implementer independent pass)

It answers the orchestrator nudge (`c3-audit-m-5/SITREP-orchestrator-planner-20260630-060057.md`) and the follow-up SITREP (`c3-audit-m-5/SITREP-orchestrator-planner-20260630-115935.md`). This is not a new audit and not a DESIGN lock. Every concrete value below is a **candidate design input** to the c3 DESIGN grill.

## Reconcile verdict

PRIMARY_BUCKET: recommended-next
- still-open: the full m-5 archetype system still needs c3 DESIGN lock under `GRILL_REQUIRED: yes`.
- already-closed: both passes agree to promote existing relay/gate, preset-as-data, topology-query, agent-type, and side-question primitives rather than rebuilding them.
- product-overlapped: m-3 owns observe execution, m-4 owns routing/template record mechanics, m-6 owns surface binding, and m-1/m-2 own store/schema substrate. m-5 should bind the archetype semantics over those mechanisms, not absorb them.
- recommended-next: proceed to the joint c3 audit-reconcile gate with the five pair deltas below recorded as converged candidate positions plus explicit DESIGN-grill carry-forwards.

Evidence:
- The planner pass says a reconcile follows and marks concrete values as non-locking c3 DESIGN inputs; `master/relays/c3-audit-m-5/AUDIT-planner-20260630-053308.md:21` and `:76-88`.
- The implementer pass says it is not the pair reconcile and that concrete values remain c3-owned; `master/relays/c3-audit-m-5/AUDIT-implementer-20260630-053116.md:20-23`.
- The orchestrator requested this exact five-item reconcile with no AUDIT value lock; `master/relays/c3-audit-m-5/SITREP-orchestrator-planner-20260630-060057.md:18-25` and `master/relays/c3-audit-m-5/SITREP-orchestrator-planner-20260630-115935.md:20-22`.

## Five-item reconcile ledger

| # | Item | Pair-reconciled disposition | DESIGN-grill carry-forward |
|---|---|---|---|
| 1 | Actuator literal value vs derived class | **Converge: derived class for Step-1.** Do not put literal `actuator` in the initial `seat_archetype` candidate list. Treat "actuator" as a derived mutating ceiling/class over `implementer` / `solo_worker` / future bounded-effector seats. | Grill whether a future `single_bounded_action` seat archetype is needed for Step 4-5 runtime enforcement. |
| 2 | Read-only work-archetypes | **Converge: keep read-only work kinds as c3 candidates, but mark their Step-1 enforcement as base read-only report predicates.** Add `research_synthesis` and `qa_review` as candidate `slot_in` values because m-5's domain covers sensor/research/QA workflows; keep `docs_chore` as a low-risk doc/config work kind. | Grill whether `research_synthesis` / `qa_review` are true `slot_in` values at Step-1 or workflow/template tags deferred to Step 5. |
| 3 | Human-mode granularity | **Converge: two-layer vocabulary.** Use planner's 3 high-level `human_mode` values as the availability/transport axis (`interactive`, `away`, `unattended`). Use implementer's finer values as `surface_intent` / delivery-intent modifiers consumed by m-6. | Grill exact field split and whether m-6 needs one field or two. Declare-before-bind remains mandatory. |
| 4 | Ceiling total order vs partial order | **Converge: partial-order/vector semantics, not a single total order.** A total risk ladder is useful for UI/gate severity, but authority itself has independent dimensions: read, tool use, source write, dispatch/route, external send, merge, and human verdict. | Grill the minimal Step-1 dimensions and their monotonic tightening rules. |
| 5 | Tag-space naming | **Converge: lower_snake_case, short canonical names, aliases allowed only as migration/display.** Prefer m-3-aligned names where existing predicate language exists. | Grill final spelling only in DESIGN; no name below is locked by this reconcile. |

## Item detail

### 1. Actuator

Planner surfaced `actuator` as a candidate `seat_archetype` with a `single-bounded-action` ceiling and noted Step-1 enforcement limits; `master/relays/c3-audit-m-5/AUDIT-planner-20260630-053308.md:90-98` and `:150-155`. Implementer flagged that `actuator` may duplicate `solo_worker` / `implementer` and may be better represented as a derived ceiling class; `master/relays/c3-audit-m-5/AUDIT-implementer-20260630-053116.md:83-85` and `:213-216`.

Reconciled position:
- For Step-1 candidate vocabulary, remove literal `actuator` from the primary `seat_archetype` list.
- Keep `actuator_class` as a derived property: a seat/template is actuator-class when it can cause mutating side effects under a stamped dispatch and observe gate.
- Keep `single_bounded_action` as a DESIGN-grill item and Step 4-5 runtime dependency because planner correctly notes uniform one-action enforcement cannot be guaranteed on existing host runtimes.

### 2. Read-only work-archetypes

Planner's candidate `slot_in` set emphasizes m-3's mutating work predicates plus `chore`/`docs`, and says read-only phases can use the base read-only report predicate; `master/relays/c3-audit-m-5/AUDIT-planner-20260630-053308.md:80-88`. Implementer added `research_synthesis` and `qa_review` to represent read-only investigation and review/triage workflows; `master/relays/c3-audit-m-5/AUDIT-implementer-20260630-053116.md:57-65` and `:94-101`.

Reconciled position:
- Candidate `slot_in` names for DESIGN input:
  - `extension`
  - `refactor`
  - `cleanup`
  - `bugfix`
  - `migration`
  - `docs_chore`
  - `research_synthesis`
  - `qa_review`
- The first five inherit m-3's richer mutation invariants. `docs_chore` is low-risk doc/config mutation. `research_synthesis` and `qa_review` use report-only/no-source-action predicates in Step 1, with richer research/QA loop invariants deferred or explicitly scoped in Step 5.
- A read-only `sensor` may staff `research_synthesis` or `qa_review`; a finding/fix emitted by those work kinds routes to a separate actuator-class workflow and never upgrades the sensor in place.

### 3. Human-mode granularity

Planner proposed a compact 3-value human-mode vocabulary: `interactive`, `away`/`async`, and `unattended`; `master/relays/c3-audit-m-5/AUDIT-planner-20260630-053308.md:159-171`. Implementer proposed seven finer values focused on m-6 delivery behavior; `master/relays/c3-audit-m-5/AUDIT-implementer-20260630-053116.md:177-201`.

Reconciled position:
- Split the concern:
  - `human_mode`: availability/transport posture with candidate values `interactive`, `away`, `unattended`.
  - `surface_intent`: delivery intent with candidate values `quiet_local`, `work_checkpoint`, `review_checkpoint`, `operator_gate`, `side_surface`, `hold_and_resummon`; `away_bridge_eligible` should be a boolean/capability flag, not a mode value.
- This keeps the planner's vocabulary small enough for routing and policy while preserving the implementer's distinctions that m-6 needs to bind surface behavior.
- Declare-before-bind is preserved: m-5 declares the vocabulary/table shape in DESIGN before m-6 locks UI/email/scheduler behavior. The c3 guardrail explicitly requires m-5 to declare human-mode vocabulary and archetype/sensor semantics before m-6 binds; `master/relays/c3-decomp/SITREP-orchestrator-reviewer-20260630-051448.md:41-47`.

### 4. Authority ceiling lattice

Planner proposed a monotonic total lattice from read-only to protected merge, but also raised the open question that dispatch authority and write authority may be orthogonal; `master/relays/c3-audit-m-5/AUDIT-planner-20260630-053308.md:117-123` and `:198-203`. Implementer separated orchestrator routing/decomposition authority from source-write authority; `master/relays/c3-audit-m-5/AUDIT-implementer-20260630-053116.md:103-108`.

Reconciled position:
- Use a partial-order/vector ceiling for authority semantics:
  - `read_context`
  - `use_tools`
  - `write_source`
  - `dispatch_route`
  - `external_send`
  - `merge`
  - `human_verdict`
- Templates may only tighten dimensions below the seat's recorded max, never loosen above it. m-4 still records `seat_archetype` plus resolved `authority_ceiling` per assignment; `master/ARCHITECTURE.md:146-164` and `:180-194`.
- A separate derived severity ladder can still support UI/gate display, but it must not pretend an orchestrator's dispatch/route authority dominates or includes source-write authority.

### 5. Tag-space value naming

Planner preferred m-3-aligned short work names such as `extension`, `refactor`, `cleanup`, `bugfix`, and `migration`; `master/relays/c3-audit-m-5/AUDIT-planner-20260630-053308.md:80-88`. Implementer used `feature_extension`, `research_synthesis`, `qa_review`, and role-shaped seat names; `master/relays/c3-audit-m-5/AUDIT-implementer-20260630-053116.md:57-85`.

Reconciled candidate names for DESIGN input:
- `slot_in`: `extension`, `refactor`, `cleanup`, `bugfix`, `migration`, `docs_chore`, `research_synthesis`, `qa_review`.
- `seat_archetype`: `sensor`, `planner`, `implementer`, `reviewer`, `solo_worker`, `orchestrator_lead`.
- Derived, not primary `seat_archetype`: `actuator_class`.

Naming rules:
- Use lower_snake_case.
- Prefer names already reflected in m-3 predicate language when there is one.
- Avoid names that imply a product promise broader than the invariant. Example: `extension` is preferred over `feature_extension` unless DESIGN decides all extensions are product-feature work.
- Aliases may appear in prose during migration, but the DESIGN lock should publish one canonical spelling per value.

## Reconciled template spine

The pair remains aligned on T1/T2/T3:
- T1 Solo: one worker seat; usually actuator-class if mutating; gate-set from `slot_in`; surface posture from `human_mode` + `surface_intent`.
- T2 Adversarial Pair: planner + implementer; review/fold lineage; gate-set from phase plus `slot_in`; planner/reviewer sides remain read/design/review, implementer mutates only after dispatch.
- T3 Sensor: one read-only side-question seat; no tools by default; one turn for side-question; advisory side surface; no source/test/prototype edits; answer content self-reported/advisory and observable metadata observed where exposed.

Conductor/N-pair remains deferred to Step 5 unless the orchestrator explicitly expands c3 scope. Both passes already preserve this narrow gate; `master/relays/c3-audit-m-5/AUDIT-planner-20260630-053308.md:126-136` and `master/relays/c3-audit-m-5/AUDIT-implementer-20260630-053116.md:115-147`.

## DESIGN-grill carry-forwards

Carry these into m-5 DESIGN under `GRILL_REQUIRED: yes`:
- Whether `single_bounded_action` becomes a future literal seat archetype or remains only a ceiling dimension/runtime capability.
- Whether `research_synthesis` and `qa_review` are Step-1 `slot_in` values or Step-5 workflow/template tags.
- Whether `human_mode` and `surface_intent` are separate fields or one normalized field.
- The exact authority-ceiling vector dimensions and monotonic tightening rules.
- Final canonical spelling for every `slot_in` and `seat_archetype` value.

No operator decision is required during AUDIT. The operator-facing choices belong in DESIGN/grill after the orchestrator's joint c3 audit-reconcile gate.

## Boundary / no-consumer check

- `slot_in` candidates consume m-3 observe predicates; for read-only values, Step 1 uses report/no-source-action predicates and records lower achieved evidence if richer checks are deferred.
- `seat_archetype` candidates consume m-4 routing/ceiling records.
- T1/T2/T3 consume m-4 template records and existing runtime pane/surface spawn.
- `human_mode` / `surface_intent` consume m-6 surface binding and scheduler behavior.
- `actuator_class` consumes m-3 observe-as-send and m-4 ceiling refusal, but uniform one-action enforcement is a later standalone-runtime concern.

ACTIONS_GIT_REF: wrote this read-only reconcile relay only: `master/relays/c3-audit-m-5/RECONCILE-implementer-20260630-120346.md` and appended the matching `master/relays/INDEX.md` row; no source/test/pcode edits claimed; cwd is not a git repo; `git -C pcode status --short` returned clean output.

FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (`git status --short` exits 128; docs workspace, `pcode/` is the code repo)

Tests / verification:
- E2: `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c3-audit-m-5/RECONCILE-implementer-20260630-120346.md` passed.
- E2: `python3 ~/.codex/skills/tools/relay-lint.py --relay-root=master/relays master/relays/c3-audit-m-5/RECONCILE-implementer-20260630-120346.md` returned `OK` for this relay and expected root noise for lint-exempt `INDEX.md`.
- E1: appended `master/relays/INDEX.md` row for `20260630-120346`.
- E1: `git -C pcode status --short` returned clean output.
- E1: `git status --short` at harness root exits 128 (`fatal: not a git repository...`), matching the expected structured unavailable form.
