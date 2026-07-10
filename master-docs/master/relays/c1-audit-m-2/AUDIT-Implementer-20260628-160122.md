## AUDIT - m-2.implementer - Forms & Determinism

ROLE: Implementer
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c1-audit-m-2-implementer
PARENT_DISPATCH_ID: c1-audit-m-2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - read-only audit; no operator decision required before design
FROM: m-2.implementer
TO: master.orchestrator-planner
CC: m-2.planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-2-forms-determinism
OWNER: m-2 (Forms & Determinism)
ACTIONS_GIT_REF: relay file created at master/relays/c1-audit-m-2/AUDIT-Implementer-20260628-160122.md; cwd git status unavailable because workspace root is not a git repository
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; git status reports fatal not-a-git-repository at /mnt mount boundary

Scope: independent read-only AUDIT of the frank form schema and field-ownership model. No source, docs, pcode, branch, commit, PR, or prototype work was performed. This relay itself is the requested AUDIT artifact.

## 4-bucket verdict

PRIMARY_BUCKET: recommended-next

still-open: the frank declarative schema artifact is not yet built. The export explicitly says the next artifact is the declarative schema `(field, owner, type, required-when, enum-set)` read by tool, courier, and linter. Evidence: the pre-build design-state export (not vendored) — routing-pillar note, line 43.

already-closed: the upstream protocol already encodes most of the needed semantics as a prose protocol plus Python linter constants and checks: minimal header fields, enums, address grammar, field clusters, dispatch/merge conditions, scan rows, fold/scope rows, and lineage. Evidence: the upstream protocol release corpus (not vendored) — protocol.md:13, protocol.md:82; <upstream relay-lint tools>/relay-lint.py:18, relay-lint.py:88.

product-overlapped: boundary ownership is intentionally shared. m-2 owns the schema/form contract, while m-1 must system-fill identity/store fields, m-3 consumes observe/evidence fields, m-4 consumes routing fields, and m-6 consumes human-gate/email fields. Evidence: CLAUDE.md:36; the pre-build design-state export (not vendored) — routing-pillar note, line 43; roadmap note, line 40.

recommended-next: promote the upstream protocol into a single declarative schema source instead of rebuilding. The seed schema should be versioned data, not prose: each field has `id`, `owner`, `type`, `enum_set`, `required_when`, `visibility_when`, `fill_constraints`, `consumer`, and `lineage_role`.

## Duplicate/already-built gate

Result: partially already built. The upstream protocol is not disposable; it is the source catalog to formalize.

Evidence:
- Protocol declares the minimal phase header and address fields. Evidence: the upstream protocol release corpus (not vendored), protocol.md:13, protocol.md:28.
- Protocol classifies canonical, design-review, and local/display fields. Evidence: the upstream protocol release corpus (not vendored), protocol.md:52, protocol.md:86, protocol.md:90.
- relay-lint has the core enum sets and minimum required fields as constants. Evidence: <upstream relay-lint tools>/relay-lint.py:18, relay-lint.py:41, relay-lint.py:88.

Do not rewrite these semantics from memory. Extract them into schema data and make protocol prose generated or checked against that data.

## Dissolve-vs-survive map

Becomes moot under a typed envelope:
- Markdown fenced-block stripping and inline-code stripping for operational token interpretation. The current code must remove code/prose before token scans because a relay is ambiguous markdown. Evidence: <upstream relay-lint tools>/relay-lint.py:116, relay-lint.py:128, relay-lint.py:138.
- Ambiguous continuation detection after blank-closed field blocks. Typed field values have explicit shape and boundaries. Evidence: <upstream relay-lint tools>/relay-lint.py:173.
- Detached row detection and duplicate row-bearing block detection. Typed `SCOPE_DIFF` and `FOLD_SCOPE` are arrays, so there is no second visual row block to smuggle. Evidence: <upstream relay-lint tools>/relay-lint.py:157, relay-lint.py:211.
- Bare own-line dispatch and merge token lexical liveness. `dispatch_grant` and `merge_grant` should be typed authority records with system-visible grantor/addressee/phase, not magic strings in prose. Evidence: <upstream relay-lint tools>/relay-lint.py:622, relay-lint.py:840, relay-lint.py:850.

Survives as form-validation:
- Enum validation for `ROLE`, `PHASE`, `AUTHORITY`, `CEREMONY_TIER`, `EVIDENCE_TARGET`, `DESIGN_RECORD_KIND`, and `DESIGN_REVIEW_VERDICT`. Evidence: <upstream relay-lint tools>/relay-lint.py:812.
- Phase/authority consistency, address grammar/cardinality, required field presence, structured absence reasons, scan-row shape/result derivation, and row arrays for scope/fold. Evidence: <upstream relay-lint tools>/relay-lint.py:414, relay-lint.py:447, relay-lint.py:355, relay-lint.py:968.
- Fill-time authority: seat-scoped enums replace post-hoc rejection. `direct-override`, merge grants, and authority-lowering options are absent unless the current seat/phase can select them. Evidence: the pre-build design-state export (not vendored) — routing-pillar note, lines 41 and 43.

Survives as lineage engine:
- Design-review parent chain from `PLAN` to approving same-owner `DESIGN-REVIEW` to same-owner `DESIGN`. Evidence: <upstream relay-lint tools>/relay-lint.py:1196.
- Pair-Planner implementation dispatch chain through approving `PLAN-REVIEW`. Evidence: <upstream relay-lint tools>/relay-lint.py:1280.
- Non-addressee action trap, merge authorization lineage, row-truth drift, and orchestrator-reviewer visibility. Evidence: <upstream relay-lint tools>/relay-lint.py:1336, relay-lint.py:1362, relay-lint.py:1376, relay-lint.py:1396.

## Seed schema recommendation

Represent the schema as data with at least:

```text
FieldSpec:
  id: string
  owner: system | seat_scoped_enum | agent_enum_pick | free_text
  type: string | enum | boolean | address | address_list | row_array | evidence_ref | object | text
  enum_set: named enum or inline list
  required_when: predicates over phase, ceremony_tier, authority, grant kind, scan state, record kind, and claimed actions
  visible_when: predicates over seat, phase, and workflow/archetype
  fill_constraints: monotonic, system_only, seat_allowed_values, computed_result, observed_value, parent_picker
  consumers: courier | form_renderer | form_validator | lineage_engine | observe_gate | router | human_surface | egress_gate
  lineage_role: none | parent_edge | grant | verdict | lock_id | action_report | merge_claim
```

Render flow:
1. Courier authenticates the lane and system-fills `FROM`, timestamp, relay id, dispatch id, default parent edge, and any observed evidence fields.
2. Form renderer chooses visible fields and enum values from `FieldSpec` plus current seat, phase, ceremony tier, workflow slot, and parent relay.
3. Agent fills only owned judgment/body fields: enum picks and free text.
4. Send runs form-validation first, then observe-as-send gate, then lineage validation against the append-only store.
5. Courier renders markdown only as a view/export. The canonical object is the typed envelope.

Prior-art fit:
- jcode uses typed Rust message/tool records and JSON-schema-like tool definitions. Evidence: jcode `crates/jcode-message-types/src/lib.rs` (local reference copy, not vendored), lines 1, 17, 82, 111.
- Claude Code SDK schemas use Zod as the single source of truth for SDK data types and control protocol. Evidence: claude-code `src/entrypoints/sdk/coreSchemas.ts` (local reference copy, not vendored), line 1; claude-code `src/entrypoints/sdk/controlSchemas.ts` (local reference copy, not vendored), line 1.
- Claude Code is directionally right at the tool boundary: `SendMessage` exposes `to`, `summary`, and `message`, not caller-supplied `from`; sender is derived inside the tool. Evidence: claude-code `src/tools/SendMessageTool/SendMessageTool.ts` (local reference copy, not vendored), lines 67 and 149. The mailbox record still has a persisted `from`, so m-1 should stamp that value at the courier/store boundary, not trust any downstream persisted object. Evidence: claude-code `src/utils/teammateMailbox.ts` (local reference copy, not vendored), lines 43 and 134.
- agent-scripts validates the operational pattern of constraining at the tool boundary: the committer helper forbids `.` and resets staged state before explicitly staging named files. Evidence: agent-scripts `scripts/committer` (local reference copy, not vendored), lines 44 and 81. Its Owner Decision Brief gives m-6 a content schema for operator gates. Evidence: agent-scripts `skills/maintainer-orchestrator/SKILL.md` (local reference copy, not vendored), line 53; the pre-build design-state export (not vendored) — external-references note, line 109.

Design challenge:
- Do not include lane-supplied `FROM` in an agent-fillable schema. m-1 must stamp it from the channel. The m-2 schema should model `FROM` as `owner: system`, not as a hidden agent parameter.
- Do not make relay bodies executable. The design export correctly says CodeAct helps for executable claims, while control-plane authorization wants fixed structured data. Evidence: the pre-build design-state export (not vendored) — routing-pillar note, line 47.

## Consumer boundary contract

Writes: declarative frank field schema and form-validation contract.

Reads: the upstream protocol/linter, the design export decisions, and prior-art typed/schema contracts.

Target entity: the frank typed relay envelope and field-ownership model.

Downstream consumer: m-1 store/stamper, m-3 observe/evidence, m-4 routing/policy, m-6 human surface/scheduler.

Contract:
- m-1 fields: `from`, `relay_id`, `dispatch_id`, `parent_dispatch_id`, `timestamp`, store sequence, and addressing graph are system-owned or courier-picked. Evidence: the pre-build design-state export (not vendored) — routing-pillar note, lines 31 and 35.
- m-3 fields: `evidence_target`, `achieved_evidence`, `actions_git_ref`, `final_git_status_short`, per-phase done predicate, executable claims, and egress/content-safety result must be expressible. Observe fields are system-captured where observable. Evidence: the pre-build design-state export (not vendored) — routing-pillar note, lines 43 and 45; external-references note, line 116.
- m-4 fields: routing record kind, dispatched roles, selected model per role, static capability prior, deviation justification, and benchmark/outcome feedback handle. Evidence: the pre-build design-state export (not vendored) — roadmap note, lines 40, 44, and 52.
- m-6 fields: `human_gate_required` as monotonic floor, `human_gate_reason`, gate bucket, `TO/CC operator`, `human-decision-required`, owner decision brief fields, park/wake/summon urgency. Evidence: the pre-build design-state export (not vendored) — routing-pillar note, lines 37 and 43; external-references note, line 109.

Proof: E1 source citations above. No E2/E3/E4 proof is appropriate in this read-only design audit.

No-consumer action: any proposed field without a named consumer must be rejected or moved to free-text body until a consumer exists. The export's rule is already "canonical iff consumed." Evidence: the pre-build design-state export (not vendored) — routing-pillar note, line 41; the upstream protocol release corpus (not vendored), protocol.md:84.

## Operator-judgment items

None for this audit. The design phase should force explicit defaults for:
- whether typed envelope canonical storage is JSON, SQLite row/object, or both;
- whether markdown remains a rendered view only or a signed export;
- how strict frank's Step-1 build is about rejecting legacy hand-authored markdown relays.

## Acceptance criteria status

1. 4-bucket verdict: satisfied, E1.
2. Dissolve-vs-survive map: satisfied, E1.
3. Field-schema design recommendation: satisfied, E1.
4. Consumer boundary contract for m-3/m-4/m-6: satisfied, E1.

## Remaining risk

The largest design risk is treating `required_when` as ad hoc code instead of data. If predicates become arbitrary functions, the schema stops being auditable. Keep predicate vocabulary finite and inspectable: phase, authority, ceremony tier, grant kind, record kind, parent phase, scan result, action claim, and workflow slot.

Next requested action: m-2.planner should reconcile this independent audit with its own and draft the DESIGN artifact for implementer review.
