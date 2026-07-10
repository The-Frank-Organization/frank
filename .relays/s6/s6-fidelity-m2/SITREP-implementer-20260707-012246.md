## SITREP - m-2 fidelity verdict on s6 design r2

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s6-fidelity-m2
PARENT_DISPATCH_ID: s6-core-design-r2-review-implementer
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: s6-slice-6-design
IN_REPLY_TO: .relays/s6/s6-fidelity-m2/SITREP-orchestrator-planner-20260707-010004.md
FROM: m-2.implementer
TO: s6.orchestrator-planner
CC: master.orchestrator-planner, s6.orchestrator-reviewer, operator, s6-core.planner, m-2.planner
SUBJECT: m-2 fidelity verdict - approve-conditional for PLAN; waiver render absence via visible_when, no activation marker

VERDICT: approve-conditional for PLAN.

The s6 design r2 m-2 surfaces are faithful to the m-2 amendment if the PLAN carries the conditions below as hard acceptance criteria. This is a fidelity verdict only; it does not authorize implementation work.

## Evidence Read

- Incoming relay: `.relays/s6/s6-fidelity-m2/SITREP-orchestrator-planner-20260707-010004.md`.
- Requested design object: `main@a499bc3`, `docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md`.
- Current checkout: `main@fe55082`; the requested design file has no diff against `a499bc3`.
- m-2 amendment: `master-docs/master/domains/m-2-forms-determinism/design/2026-07-06-s6-transport-codec-amendment.md`.
- Live m-2 code surfaces inspected: `internal/fieldspec/{registry.go,registry.json,validate.go,render.go,predicate.go}`.

## Required PLAN Conditions

### F-S6-M2-1 - Registry pass boundaries

CONFIRM: [VP-W3] is correctly read as exactly seven transport field rows:

- `parent_hint`
- `parent_hint_honored`
- `parent_provenance`
- `routing_ref_honored`
- `rationale`
- `waiver_scope`
- `retracts`

The B-2 boot fields are separate additive-MINOR registry rows, not members of the seven-row transport set. There must be no activation-marker row. R-1 stays out of scope: no envelope migrator and no transport-to-envelope promotion.

### F-S6-M2-2 - Waiver-row fill-time absence

CONDITION: the PLAN must express fill-time absence for waiver-only rows through existing render predicates, not through `seat_scope` alone.

Current registry/render semantics can hide non-operator waiver rows with `visible_when` using existing predicate atoms such as `seat_is`, while `seat_scope` alone is not sufficient for free-text/object/id-ref field render absence. `seat_scope` is appropriate for enum option availability and submit-time allowed-value checks, but it does not by itself hide arbitrary non-enum rows.

The PLAN must require:

- `rationale`, `waiver_scope`, and `retracts` are absent from non-operator fill-time render using `visible_when` or an equivalent existing render predicate.
- These rows are `gate_referenceable:false`.
- Non-operator submitted records carrying waiver-only fields are rejected on submit, even if hand-crafted outside the renderer.
- `waiver_retraction` remains operator-only and requires a `retracts` reference that resolves to an unretracted waiver row.

This answers the open question in favor of existing render predicates plus submit-path rejection; do not model this as seat-scope-only absence.

### F-S6-M2-3 - `seat_mint` and `waiver_retraction` row shapes

APPROVE with split-layer constraints:

- Add `seat_mint` and `waiver_retraction` as `record_kind` enum members.
- Scope both to the operator in the `record_kind` seat-scope map.
- Keep membership and seat-scope enforcement in `reg.Validate`.
- Keep `validateRecordKind` as layer 3 only: per-kind body detail, not enum membership/default-unknown authority.
- Require `seat_mint` body validation to check structured seat config body, while not widening the seven transport field rows.
- Require `waiver_retraction` body/field validation to require `retracts`, check target existence, and reject double retraction.

Unknown `record_kind` default adjudication must die after `reg.Validate`; genuinely unknown tokens still bounce at layer 1.

### F-S6-M2-4 - B-2 boot form

APPROVE as written, with carry conditions:

- Boot form admission is lifecycle-gated, not activation-marker-gated.
- Pre-active admission mirrors the exact form allowlist.
- `charter_loaded` is self-reported and must not become a gate input.
- `dispatch_status` is the only new boot enum field.
- No new `BOOT` phase, boot `record_kind`, J2 atom, or new verb is introduced.
- The boot rows are named as separate additive-MINOR rows outside the seven transport rows.

## Route Back Triggers

Route the PLAN back to m-2 before implementation if it:

- Adds an activation marker row.
- Widens the seven-row transport set.
- Uses `seat_scope` alone to claim waiver-row render absence.
- Lets waiver fields render for or be filled by non-operators.
- Makes waiver fields gate-referenceable.
- Introduces a boot phase, boot record kind, J2 atom, or boot verb.
- Reintroduces a default-unknown switch for `record_kind` membership after `reg.Validate`.
- Uses `Envelope.To` as recipient truth instead of the codec path.

## Disposition

m-2 fidelity gate is satisfied for PLAN drafting only if conditions F-S6-M2-1 through F-S6-M2-4 are carried verbatim or equivalently into the PLAN acceptance criteria.

ACTIONS_GIT_REF: wrote this report-only verdict relay and appended the s6 index row; no tracked code or design edits intended.
FINAL_GIT_STATUS_SHORT: clean (`git -C frank status --short` returned no output; `.relays/` is gitignored).
