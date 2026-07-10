## AUDIT - s5-a.implementer registry and row inventory

ROLE: Implementer
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: s5-audit-s5-a
PARENT_DISPATCH_ID: s5-resume
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s5-a.implementer
TO: s5.orchestrator-planner
CC: s5.orchestrator-reviewer, operator
SUBJECT: AUDIT - s5-a registry fields, row inventory, fixture surface, and dormancy gates

IN_REPLY_TO: `.relays/s5/boot/s5-boot-s5-a-implementer/SITREP-orchestrator-planner-20260706-040203.md`; `.relays/s5/s5-audit-s5-a/AUDIT-orchestrator-planner-20260706-035446.md`

## Boot acknowledgement folded into audit

- Identity accepted: this session is `s5-a.implementer`, the Implementer seat for pair `s5-a`.
- Skill loaded: `agent-pair-implementer` plus `protocol.md` read in full.
- Team charter read: `master-docs/CLAUDE.md`.
- Boot relay read: `.relays/s5/boot/s5-boot-s5-a-implementer/SITREP-orchestrator-planner-20260706-040203.md`.
- Mode-change relay read: `.relays/s5/s5-resume/PLAN-orchestrator-planner-20260706-034602.md`.
- Fidelity reconciliation read: `.relays/s5/s5-fidelity/RECONCILE-orchestrator-planner-20260706-034602.md`.
- Audit dispatch read: `.relays/s5/s5-audit-s5-a/AUDIT-orchestrator-planner-20260706-035446.md`.
- Transport rule accepted: no frank MCP tools used or retried; this report is a file relay under `.relays/s5/`.

## Routing and authority

- The audit dispatch is addressed to `s5-a.planner, s5-a.implementer`, so this Implementer seat is directly in `TO`.
- Authority is read-only AUDIT. No source, test, registry, sprint doc, branch, commit, PR, merge, transport fix, or store archive action was taken.
- The boot relay grants report-only identity authority; the work authority comes only from the addressed AUDIT dispatch.
- This relay reports implementer findings for reconciliation. It is not a plan and does not authorize implementation.

## Summary verdict

PRIMARY_BUCKET: still-open
still-open: The live registry is still an S3/S4 seed, not the S5 consumer-schema registry. It has 47 fields and several useful rows, but most m-3/m-4/m-5/m-6 consumer slots are absent, `routing_escalation` is not in either target enum, and row-array fields cannot currently express the per-column ownership/model-identity/gate-referenceability semantics required by routing assignments and ODB choices.
already-closed: `gate_category_raised`, grill rows, owed-disposition rows, `member`, `new_digest`, `config_change` operator scoping, held/rejected delivery states, typed canonical validation for top-level row_array/object carriers, and predicate rejection of concrete Step-1 `slot_in` values already exist.
product-overlapped: s5-a owns registry rows and fixture payloads. s5-b owns submit-stage raise mechanics, dormant egress scanner/drain seams, replay/versioning mechanisms, and I-PH extensions for strings those mechanisms create. Transport-fix lineage/codec issues remain out of scope for this pass.
recommended-next: Reconcile this inventory before PLAN. The implementation should be one narrow registry/config pass plus focused registry tests and fixture payloads. Do not begin with broad engine changes from this pair.

## 1. Live registry baseline

PRIMARY_BUCKET: still-open
already-closed: `internal/fieldspec/registry.json` is parseable and current tests pass. It declares `version: "s3-fieldspec-v2"` and 47 fields. Evidence: `internal/fieldspec/registry.json:2`, registry field count by `jq`, and `go test ./internal/fieldspec`.
already-closed: The registry loader validates owners, types, fill constraints, enum references, and the R2 rule that a field cannot be both `model_identity` and `gate_referenceable`. Evidence: `internal/fieldspec/registry.go:135-182`.
already-closed: `row_array` and `object` have canonical typed carriers, but only as generic top-level maps. Evidence: `internal/fieldspec/canonical.go:19-35`.
still-open: There is no nested schema for row-array columns. That is the main blocker for faithfully declaring `routing_assignments` columns such as `chosen_model`, `seat_archetype`, and `authority_ceiling`, because those columns need their own ownership, model-identity, and gate-referenceability contracts.
recommended-next: Treat row-array nested schemas as a design reconcile item. If implementation does not extend the registry grammar, use fixture-local validation for nested columns and state the limitation explicitly.

## 2. Required enum delta

PRIMARY_BUCKET: still-open
still-open: `routing_escalation` is absent from `named_enums.gate_category_A` and `named_enums.gate_category`. Evidence: `internal/fieldspec/registry.json:37-67`.
already-closed: The existing category classifier treats unknown values as A/raised and treats `other` plus known A tokens as authority-bearing. Evidence: `internal/fieldspec/validate.go:176-197`.
recommended-next: Apply the exact fidelity delta: append `routing_escalation` to `named_enums.gate_category_A`; insert `routing_escalation` into `named_enums.gate_category` before `other`; do not change `gate_category_B`; do not add a `FieldSpec` row for `routing_escalation`; bump the registry version label only.

Compatibility note:

- The delta is additive/MINOR at the registry level.
- It is byte-distinct from `routing_unavailable` and `human_decision_required`, which are routing outcome states, not `gate_category` enum members.

## 3. Existing rows to preserve or fix

PRIMARY_BUCKET: recommended-next

| field or row | status | finding |
| --- | --- | --- |
| `gate_category_raised` | preserve | Exists as `owner:"computed"`, bool, `fill_constraints:"computed_result"`. It is hidden from Step-1 rendering because computed rows are excluded. |
| `GRILL_REQUIRED`, `GRILL_LOCK_ID` | preserve | Present and covered by registry tests. `GRILL_LOCK_ID` is dependent on `GRILL_REQUIRED == yes`. |
| `owner`, `source`, `target_surface`, `disposition_path`, `disposes_owed` | preserve | Owed/disposition rows exist and are guarded by `record_kind_in` predicates. |
| `member`, `new_digest` | preserve | Config-change rows exist and should stay with config-change fixtures. |
| `delivery_state`, `failing_edge` | preserve | Held/rejected delivery-state vocabulary exists and supports m-6 bucket mapping. |
| `ACTIONS_GIT_REF`, `FINAL_GIT_STATUS_SHORT` | fix | They are `owner:"agent_enum_pick"` with `required_when: layer_present observe`, but have no `visible_when`. Because rendering also uses `DefaultLayers()` without observe, they render in Step-1 forms while validation does not require them. |
| `record_kind` | reconcile | `config_change` is already operator-only. The `*` scope still includes `genesis`, `owed_item`, and `owed_disposition`; OI-S4-TOKEN-SCOPE calls for narrowing genesis/owed authoring from wildcard scope toward operator control. |

Recommended row-level fixes:

- Add `visible_when: layer_present observe` anywhere a field is observe-owned or observe-computed, starting with `ACTIONS_GIT_REF` and `FINAL_GIT_STATUS_SHORT`.
- Remove `genesis` from the wildcard `record_kind` scope. Narrowing `owed_item` and `owed_disposition` is policy-compatible but should be confirmed in reconcile because current docs say non-operator owed filing has no authority hole and operator authors at discretion.
- Keep `config_change` operator-only; that part is already correct.

## 4. m-3 observe and evidence rows

PRIMARY_BUCKET: still-open
still-open: The observe/evidence contract from m-3 is mostly absent from the live registry. The missing rows should be observe-gated and should not appear in Step-1 forms.

| target field | proposed shape | dormancy rule |
| --- | --- | --- |
| `achieved_evidence` | enum with E0-E4 evidence ladder values; system/observed | `required_when` and `visible_when`: `layer_present observe` |
| `target_gap_result` | enum such as `met`, `target_gt_achieved`, `not_applicable`; computed | observe-gated |
| `evidence_integrity` | object or row_array carrying observed/self-reported/mixed classification | observe-gated; needs nested-schema support if row_array |
| `record_integrity` | enum `observed`, `self_reported`, `mixed`; computed | observe-gated |
| `executable_claim_results` | row_array of claim/result tuples | observe-gated; nested columns need schema support |
| `egress_scan_result` | enum `pass`, `blocked`, `not_applicable`; system/observed | observe-gated |
| `degradation_notes` | text, system/observed | observe-gated |
| `attestation_source` | enum `conductor`, `operator`; system/computed provenance marker | observe-gated; not a third evidence-integrity class |
| `authority_class` | bool or enum bool; computed | observe-gated; mixed evidence maps to held downstream |
| `deviated_observed` | bool; computed from declared-vs-observed routing | observe-gated |
| `bucket_binding_observed` | bool; computed | observe-gated |

Do not add `predicate_result` or `veto` as persisted fields. Fidelity says these are not part of the m-3 persisted subset. `rank1_recommended_bucket` is derived, not persisted.

## 5. m-4 routing rows

PRIMARY_BUCKET: still-open
still-open: The routing-decision fields are absent from the live registry. The largest implementation risk is that the published design describes structured per-assignment columns, while the current registry can only type the top-level `routing_assignments` carrier.

| target field | proposed shape | finding |
| --- | --- | --- |
| `routing_record_kind` | enum, reserved for routing-decision records | Use a separate field row if needed; do not widen top-level `record_kind` without reconcile. |
| `routing_assignments` | row_array | Needs nested columns for seat/bucket/model/pin/deviation/archetype/ceiling; current grammar cannot encode them. |
| `capability_prior_snapshot` | object | System/computed snapshot; observe/config sourced. |
| `declared_deviated` | bool, nested under assignment | Declared by routing author; required per assignment. |
| `justified_deviation` | text | Required only when `declared_deviated == true`. Current predicate vocabulary cannot address nested row values. |
| `deviation_reason_code` | enum | Default vocabulary: `capability_gap`, `cost_budget`, `latency_budget`, `bucket_unavailable`, `operator_directive`, `experiment`, `other`. |
| `constraints` | object | Reserved-shape, no concrete Step-1 values. |
| `template_ref` | id_ref | Reserved/dormant; operator template exception requires reconcile if encoded as seat scope. |
| `outcome_feedback_ref` | id_ref | Null-reserved in this slice. |

Nested assignment columns called out by design:

- `declared_bucket`: enum from routing buckets.
- `chosen_model`: model-identity column, not gate-referenceable.
- `pin_mode`: enum for pin behavior.
- `seat_archetype`: open tag-space, not a literal actuator.
- `authority_ceiling`: object, reserved-shape open named-axis map.

Recommended reconcile gate:

- Decide whether to extend `FieldSpec` with nested row schemas now. If not, write explicit tests proving the top-level carrier is only a carrier and that downstream fixtures validate nested assignment shape separately.
- Keep `routing_unavailable` and `human_decision_required` out of `gate_category`.
- Do not create a live top-level `record_kind = routing` unless the no-live-widening ruling is clarified.

## 6. m-5 role-composition rows

PRIMARY_BUCKET: recommended-next
already-closed: The predicate engine already rejects concrete Step-1 slot values for `slot_in`, preserving the reserved-shape invariant. Evidence: `internal/fieldspec/predicate.go:155-162` and `internal/fieldspec/registry_test.go:179-206`.
still-open: The registry has no top-level `slot_in`, `seat_archetype`, or `authority_ceiling` rows.
recommended-next: Keep `slot_in` dormant unless a later slice needs it as a rendered/stored field. Model `seat_archetype` and `authority_ceiling` as columns inside `routing_assignments` rather than standalone Step-1 fields. Do not predeclare `external_send`. Keep `accepts_interjection` and template-schema behavior in config/template records, not ordinary FieldSpec rows.

## 7. m-6 human-surface and ODB rows

PRIMARY_BUCKET: still-open
still-open: ODB and away-bridge rows are absent. Fidelity says there is no live ODB `record_kind` widening in Step-1, so these should be fixture-scoped or dormant unless reconcile changes that ruling.

| target field | proposed shape | finding |
| --- | --- | --- |
| `away_bridge_eligible` | bool, computed/system | Default policy-derived gate; keep dormant until away-bridge activation. |
| `plain_language_change` | text | ODB human-surface row; fixture-scoped/dormant. |
| `why_now` | text | ODB human-surface row; fixture-scoped/dormant. |
| `completed_proof` | evidence_ref or text | ODB proof row; source should be observed/system. |
| `tradeoffs_risks` | text | ODB human-surface row; fixture-scoped/dormant. |
| `recommendation` | text or enum | ODB human-surface row; fixture-scoped/dormant. |
| `choices` | row_array | Needs nested choice schema if used beyond fixture text. |
| `subject_ref` | id_ref | System reference to the thing under decision. |
| `model_name` | string/text, system, `model_identity:true`, not gate-referenceable | ODB carve-out field; do not use outside the allowed ODB fixture context. |
| `decision_deadline` | string/text timestamp carrier | Null-reserved unless timeout behavior is implemented. |
| `on_timeout` | enum | Null-reserved unless timeout behavior is implemented. |

Recommended reconcile gate:

- If ODB rows are landed live, define the record-kind and rendering boundary first.
- If ODB rows remain fixture-scoped, keep them out of the live registry and let s5-b scanner tests use fixture-local registry views.
- Preserve the m-6 precedence rule: human-surface help does not override egress safety or authority gates.

## 8. Fixture matrix and negative legs

PRIMARY_BUCKET: recommended-next
recommended-next: Every Step-1 fixture should enumerate role, phase, ceremony tier, and grant state rather than only one happy path.

Minimum render/validate axes:

- Seats/roles: operator, orchestrator-planner, orchestrator-reviewer, pair planner, pair implementer.
- Phases: all live PHASE enum values: AUDIT, DESIGN, DESIGN-REVIEW, PLAN, PLAN-REVIEW, IMPL, REVIEW-FOLD, MERGE-GATE, LIVE-VERIFY, SITREP, RECONCILE.
- Tiers: tiny, small, medium, large, production-risk.
- Grant states: planner grant closed/open; MERGE-GATE versus non-MERGE-GATE for merge-token rendering.

Required negative legs:

- Observe-owned required fields are gated off in Step-1 forms. They must be absent, not merely optional. Current `ACTIONS_GIT_REF` and `FINAL_GIT_STATUS_SHORT` are counterexamples.
- Non-observe required fields still block submit: for example `SUBJECT`, owed rows under `record_kind_in`, and `GRILL_LOCK_ID` when `GRILL_REQUIRED == yes`.
- `EVIDENCE_TARGET` remains Step-1 intent and required; it is not observe-owned.
- `layer_present` predicates must never depend on the model or on model identity values.

## 9. Version label and section-7 implications

PRIMARY_BUCKET: recommended-next
still-open: The registry version label is still `s3-fieldspec-v2`. Evidence: `internal/fieldspec/registry.json:2`.
recommended-next: Bump the registry label as part of the S5 registry delta, for example to `s5-fieldspec-v1` after pair reconciliation. Do not bump `migrate.Current` for registry-only changes; fidelity says no envelope migrator is warranted without canonical record-shape changes.

Section-7 fixture implications for s5-a:

- Provide the new registry bytes and digest expected by the config-change fixture.
- Preserve old-to-new digest assertions.
- Do not re-genesis on config-change acceptance.
- Keep phase-0 acceptance tied to the genesis plus accepted config-change chain.
- Exercise stale-form bounce and fresh re-render success using the new registry label.

## Boundary list

Expected s5-a IN surfaces at IMPL time:

- `internal/fieldspec/registry.json`: enum delta, version label, row additions/fixes.
- `internal/fieldspec/registry_test.go`: inventory, enum, grill/owed/token-scope, nested-schema guard tests.
- `internal/fieldspec/render_test.go`: observe-gated visibility, Step-1 absence, grant/phase fixtures.
- `internal/fieldspec/validate_test.go`: required_when/visible_when and non-observe required negative legs.
- `test/fixtures/s4_config_change_test.go` or a new `test/fixtures/s5_config_change_test.go`: section-7 registry/config-change fixture payloads.

Potential s5-a IN only if reconcile approves:

- `internal/fieldspec/registry.go`, `canonical.go`, and related tests if nested row schemas are added.
- Fixture-local registry files for ODB/routing/scanner tests when live widening remains forbidden.

Expected OUT or collision surfaces:

- Submit-stage raise mechanics and classifier plumbing in engine code: s5-b.
- Dormant egress scanner and drain call sites: s5-b.
- Replay/read-path negotiation mechanics: s5-b.
- I-PH additions for new scanner/read-path strings: s5-b.
- Transport-fix lineage/parenting/codec changes: out of S5 audit implementation.
- Live external send or away-bridge activation: later slice.

## Open questions for reconcile

1. Should `row_array` gain nested schemas in S5, or should routing/ODB nested rows remain fixture-validated carriers for now?
2. Does "no live record_kind widening" apply globally to routing records, or only to ODB/model_name records? Current live `record_kind` lacks routing and ODB members.
3. Should OI-S4-TOKEN-SCOPE remove only `genesis` from wildcard scope, or also move `owed_item` and `owed_disposition` out of wildcard scope now?
4. Should `gate_category_raised` remain `owner:"computed"` or move to `owner:"system"`? Current rendering behavior works either way, but s5-b submit mechanics need a stable contract.
5. What final registry version label should be used: `s5-fieldspec-v1`, `s5-consumer-fieldspec-v1`, or another operator-preferred label?

## Actions

- Performed read-only audit of the addressed s5-a boot and AUDIT relays plus current registry/design/test surfaces.
- Wrote this audit relay and appended the s5 index row.
- No code, tests, registry, docs outside `.relays/`, branches, commits, PRs, merges, store archives, or transport fixes were changed.

ACTIONS_GIT_REF: read-only audit relay written at `.relays/s5/s5-audit-s5-a/AUDIT-implementer-20260706-041357.md`; index row appended at `.relays/s5/INDEX.md`; no source/test/runtime edits made; `.relays/` is gitignored.
FINAL_GIT_STATUS_SHORT: none - clean tree

## Verification

- E1: exact boot relay read: `.relays/s5/boot/s5-boot-s5-a-implementer/SITREP-orchestrator-planner-20260706-040203.md`.
- E1: exact AUDIT dispatch read: `.relays/s5/s5-audit-s5-a/AUDIT-orchestrator-planner-20260706-035446.md`.
- E1: master mode-change and fidelity relays read: `.relays/s5/s5-resume/PLAN-orchestrator-planner-20260706-034602.md` and `.relays/s5/s5-fidelity/RECONCILE-orchestrator-planner-20260706-034602.md`.
- E1: design records spot-checked: master `ARCHITECTURE.md`, m-2 form schema design, m-3 observe design, m-4 routing design, m-5 role-composition design, and m-6 human-surface design.
- E1: code surfaces read and cited: `internal/fieldspec/registry.json`, `internal/fieldspec/registry.go`, `internal/fieldspec/render.go`, `internal/fieldspec/validate.go`, `internal/fieldspec/predicate.go`, `internal/fieldspec/canonical.go`, and related fieldspec tests.
- E2: `go test ./internal/fieldspec` passed.
- E2: `git status --short` before relay write produced no output.
- E2: relay lint on this exact file passed after the index row was appended.
