# s5-a — the single registry pass + the [VP-W3] enumerated negative dormancy fixture (DESIGN)

**DESIGN_DOC_ID:** s5-a-registry-design
**Owner:** pair s5-a "registry & rows" · authored s5-a.planner · for s5-a.implementer DESIGN-REVIEW
**Date:** 2026-07-06 · **RUN_ID:** s5 · **Tier:** medium · **Evidence:** E1 (every code claim file:line @ `main 67ee23e`)
**Basis (semantics authorities, in precedence order):** `.relays/s5/s5-escalations/RECONCILE-orchestrator-planner-20260706-053113.md` (**the riding-leg CLOSE: zero open m-x legs; owed operator-only; genesis in NO scope; the adjacent flag folds in-pass; MR-2**) · `.relays/s5/s5-escalations/RECONCILE-orchestrator-planner-20260706-052214.md` (M-1 idiom (i) blessed; M-3 (a)–(k); MR-1) · `.relays/s5/s5-fidelity/RECONCILE-orchestrator-planner-20260706-034602.md` (Q1–Q11) · `.relays/s5/s5-reconcile-audits/RECONCILE-orchestrator-planner-20260706-045327.md` (R-s5-1..7, DEF-1..5) · the two design dispatches (`.relays/s5/s5-design-s5-a/DESIGN-orchestrator-planner-20260706-045327.md`, `…-052753.md`) · m-2 design-of-record §4/§5/§12/§17 (grammar) · `ARCHITECTURE.md` §C4 (the ③ settled note + C1/C2 Step-3 carries — cited, not restated) + §J2.
**rev1 (2026-07-06):** folds the s5-a.implementer DESIGN-REVIEW must-revise (`.relays/s5/s5-design-s5-a/DESIGN-REVIEW-implementer-20260706-055207.md`) — the 053113 close (stale variants deleted), the omitted `on_timeout` row added, acceptance criterion 6's test-surface contradiction fixed. Fold log §12.
**Design Q&A record:** ran as the relay escalation ladder (pair audits → s5-reconcile-audits → s5-escalations → owner answers → master reconcile) — the file-relay form of the pair-design inline Q&A; no question remains open that this doc does not mark as a named confirm/riding leg.

---

## 1. Shape of the change — ONE atomic registry content pass, no code

One edit to `internal/fieldspec/registry.json` (plus its registry-content-coupled test files per R-s5-2). No `.go` code changes — every code-side item (DEF-1 byte fix, DEF-2 typed-REJECT guard, ③ raise mechanics, knownA wiring) is s5-b's surface; this doc only states the row-side interface s5-b consumes (§8).

Net counts: **47 → 83 rows** (+36) · **14 → 24 named_enums** (+10 new, 2 edited) · **4 existing rows edited in place** (EVIDENCE_TARGET, ACTIONS_GIT_REF, FINAL_GIT_STATUS_SHORT, record_kind) · **version label `s3-fieldspec-v2` → `s5-fieldspec-v3`** + provenance block update (R-s5-1). All other 43 existing rows byte-stable.

**Annotation mechanics (design decision D-1):** m-2's REQUIRED per-row annotation lands as an `"annotation"` string key on each row. The loader ignores unknown JSON keys (plain `json.Unmarshal`, no `DisallowUnknownFields` — registry.go:64-81), so annotations are bytes-only: digest-covered, machine-transparent, human/审readable. Flagged for implementer review + the m-2 shape confirm (§9) since `annotation` is not one of the 12 grammar columns.

## 2. named_enums

**Edited (the routing_escalation delta, verify-done at both audits):**
- `gate_category_A`: append `"routing_escalation"` → 9 members.
- `gate_category`: insert `"routing_escalation"` immediately before `"other"` → 14 members; `other` stays last.
- `gate_category_B`: UNTOUCHED. `routing_unavailable` / `human_decision_required` are NOT added anywhere (byte-distinctness verified: zero occurrences at 67ee23e).

**New (10):**

| enum | tokens | source |
|---|---|---|
| `achieved_evidence` | E0, E1, E2, E3, E4 | m-3 §3:50 (E0 in range ⇒ EVIDENCE_TARGET's E1–E4 set is NOT reusable) |
| `target_gap_result` | met, target_gt_achieved, not_applicable | m-3 §3:51 |
| `evidence_integrity` | observed, self_reported | R3-locked two-value (m-3 §6:122) |
| `record_integrity` | observed, self_reported, mixed | m-3 §6:123; m-2 §17.6:344 |
| `executable_claim_result` | pass, fail, skipped, unsafe | m-3 §3:54 (per-claim vocabulary; design decision D-2: declared so fixtures assert byte-exactness) |
| `egress_scan_result` | pass, blocked, not_applicable | m-3 §3:55 |
| `attestation_source` | conductor, operator | RECONCILE Q9(2) / m-3 §13:222 (O-2) |
| `deviation_reason_code` | capability_gap, cost_budget, latency_budget, bucket_unavailable, operator_directive, experiment, other | M-3(a): named_enums mirror of m-4's 7 defaults + row annotation "value-set is config-sourced (default-seeded, operator-configurable, hardcoded other fail-safe — m-4 §5/§6)" |
| `routing_record_kind` | routing_decision | m-4 §5:199-207 — Step-1 declares the one live token only; routing_deviation/routing_update stay reserved for a later release, NOT declared |
| `surface_intent` | progress, review_checkpoint, advisory, result | M-3(i); m-2 §17.6:346 |

Precedent for enums not referenced by any field's `enum_set` (evidence_integrity, executable_claim_result as vocabularies of object/row_array internals): `gate_category_A`/`gate_category_B` already exist exactly so (registry.json:37-52).

## 3. New rows — Block A: observe/evidence (idiom A, 12 rows)

Uniform treatment: `owner: system` (or `computed` where the live precedent is computed — gate_category_raised class), `fill_constraints: observed_value`/`computed_result` per column below, `gate_referenceable: false` (default — key simply absent), **`required_when` AND `visible_when` both `{"all_of":[{"layer_present":"observe"}]}`** (the validate.go:29-37 / render.go:50-54 symmetry my audit's trap analysis requires; live precedent predicate_test.go:64-77). Owner system/computed rows never render anyway (render.go:51) — the predicates are declared regardless so the idiom is uniform data, not an accident of owner.

| id | type | enum_set | owner | fill | annotation core |
|---|---|---|---|---|---|
| achieved_evidence | enum | achieved_evidence | system | observed_value | conductor fact, ladder depth (m-3 §4) |
| target_gap_result | enum | target_gap_result | system | computed_result | computed target>achieved flag |
| evidence_integrity | object | — (values = evidence_integrity enum, per-field map) | system | observed_value | per-field R3 two-value; DI-5 |
| record_integrity | enum | record_integrity | system | computed_result | pure function of per-field stamps; mixed = not-fully-verified |
| executable_claim_results | row_array | — (result column vocabulary = executable_claim_result enum) | system | observed_value | per-claim registry-check results |
| egress_scan_result | enum | egress_scan_result | system | observed_value | outbound chokepoint only; never terminal (m-3 §3.3) |
| degradation_notes | text | — | system | observed_value | bounded-observation floor |
| attestation_source | enum | attestation_source | system | observed_value | provenance marker, NOT a third integrity value (O-2) |
| authority_class | bool | bool | computed | computed_result | BOOL (Q9(4)); = record_kind ∈ authority-set ∨ gate_category ∈ A; never model-keyed |
| deviated_observed | bool | bool | system | computed_result | GL-1 bucket-vs-bucket; NEVER gate-referenced |
| bucket_binding_observed | bool | bool | system | computed_result | chosen_model ∈ members(declared_bucket); observe-side only |
| surface_intent | enum | surface_intent | system | computed_result | M-3(i); non-gate-bearing records only (Step-2+ semantics, annotation only — not render-expressible); posture value-enum stays m-5 config |

**Honesty annotation on every Block-A row (DEF-2, [VP-W1]):** "suppliability guard = the s5-b (h) typed-REJECT validator rule (channel-keyed, envelope-asymmetry preserved); until it lands, dormancy is render-absence, not submit-rejection — no non-lane-writability claim."

## 4. New rows — Block B: reserved-shape (13 rows)

Uniform: `owner: system`, `fill_constraints: system_only` (except noted), NO predicates (reserved shape — the filling mechanism arrives with its step), `gate_referenceable: false`, no schema-enumerated values beyond the stated enum_set.

| id | type | notes |
|---|---|---|
| slot_in | string | opaque work-archetype tag; conductor-classified at Step-2 acceptance (CQ-5 pin, m-3 §5.1); NO values — the predicate-atom guard (predicate.go:155-165) stays untouched; field-id vs atom-key namespaces are disjoint (atom keys are predicate-object keys; field lookup happens only for field/any_row atoms — predicate.go:121-154) |
| seat_archetype | string | opaque; conductor-STAMPED per m-5-F1; top-level row = the shape/vocabulary home; per-assignment carriage = a routing_assignments column (§5, degraded shape M-3(f)) |
| authority_ceiling | object | open named-axis map {write, dispatch, tool}; each axis's absent-default = its most-restrictive floor (m-5 §5:88); the ceiling object ITSELF is Step-1 F2-recordable (Q10) — reserved-shape does not forbid a carried value (m-2 AC15:240 boundary); external_send NOT declared |
| capability_prior_snapshot | object | fill computed_result; both prior layers, replay-complete (m-4 §4:190-192) |
| routing_record_kind | enum (routing_record_kind) | owner agent_enum_pick + fill seat_allowed_values — one of the M-1 ten, so it carries the Block-C step-gate visible_when + annotation (§5); listed here because its VALUE set is reserved-narrow |
| template_ref | id_ref | set when template-spawned, null otherwise (m-4 §5:214); lineage_role none |
| outcome_feedback_ref | id_ref | null-reserved for a later release |
| subject_ref | id_ref | ODB; system-projected from the envelope |
| decision_deadline | string | null-reserved scheduler slot; no timestamp type exists in the live type set (registry.go:233-244) — string + annotation (design decision D-3) |
| on_timeout | enum | **(rev1 — was omitted; M-3(d); rev2 — validation claim corrected)** valueless reserved slot: NO enum_set, NO options (loadable — registry.go:170-174 checks enum_set only when set; `enum` + `system_only` are valid row values). owner system, fill system_only. **Accurate Step-1 boundary (same as every Block-A/B/D system row):** the validator never reaches the enum check for this row — ignorePayloadField skips owner:system/system_only rows before required/type/enum checks (validate.go:31-53, :115-120) — so Step-1 protection is RENDER ABSENCE (render.go:51); submit-path rejection of a lane-supplied value depends on s5-b's DEF-2 typed-REJECT guard; conductor-internal value conformance is the later scheduler-writer's responsibility, not current FieldSpec validation. Annotation carries m-6's policy floor VERBATIM: "default hold_and_resummon is J1 policy, not a registry value; **no value may ever mean auto-approve/auto-resolve — only conservative block-ceiling tightening is legal (§J1)**" |
| completed_proof | evidence_ref | fill observed_value; m-3-sourced, never agent free-text (m-2 §17.2:292) |
| away_bridge_eligible | bool (bool) | owner computed, fill computed_result; m-6 policy-derived; ROW HOME ONLY — the step-(d) away-bridge mechanism stays OUT ([VP-W4]) |
| model_name | string | **model_identity: true** (load-guard proves gate_referenceable false — registry.go:162-164); consumers [egress_gate, human_surface]; the typed exempt-marked ODB field (decision ⑤); annotation carries the corrected honesty phrasing **(rev4 — the panel's F-SEC-1: the earlier "non-lane-writable via the tool surface" clause was FALSE as worded, since the tool surface includes submit() and owner:system rows are validation-skipped)**: "not offered on any rendered form (render-absent); lane-suppliable via raw submit headers until the s5-b (h) typed-REJECT guard lands; no non-lane-writability claim" |

## 5. New rows — Block C: the M-1 ten (idiom (i) — BLESSED step-gate, 9 rows here + routing_record_kind in §4)

Every row: `visible_when: {"all_of":[{"layer_present":"observe"}]}`, and `{"layer_present":"observe"}` conjoined into any required_when it carries. **m-2's REQUIRED verbatim annotation on each row:** *"gated to the post-Step-1 consumer fill-layer; NOT observe-owned (owner stays agent_enum_pick/free_text)."* Plus the documented limitation (travels once in this doc, per the escalation §1): the idiom couples consumer dormancy to observe-layer-presence; a future selective-withhold re-points to its own layer atom (grammar extension, later cycle — registered in m-2's annotation, not s5's problem). The [VP-W3] fixture covers ALL these rows at the render gate — no exclusion list.

| id | type | owner / fill | required_when |
|---|---|---|---|
| routing_assignments | row_array | seat_scoped_enum / seat_allowed_values — but see D-4 | none |
| justified_deviation | text | free_text / free_text | `{"all_of":[{"layer_present":"observe"},{"any_row":"routing_assignments.declared_deviated","op":"==","value":"yes"}]}` |
| deviation_reason_code | enum (deviation_reason_code) | agent_enum_pick / seat_allowed_values | same as justified_deviation (same grain — m-4 §5:211) |
| constraints | object | free_text / free_text | none (reserved/forward values) |
| plain_language_change · why_now · tradeoffs_risks · recommendation | text ×4 | free_text / free_text | none |
| choices | row_array | agent_enum_pick / free_text | none |

**D-4 (routing_assignments, the degraded shape — M-3(f), m-2+m-4 confirmed, C1+C2 registered in §C4):** `gate_referenceable: true` is REQUIRED on this one row so the `any_row` predicate may name it (parse validates the ARRAY id — predicate.go:145; precedent SCOPE_DIFF registry.json:106). No column specs exist in the live grammar — the row's annotation documents the 9 columns (seat, role, task_tag, declared_bucket, chosen_model, pin_mode, declared_deviated, seat_archetype, authority_ceiling) with per-column ownership (planner proposes seat/role/seat_archetype; conductor stamps seat_archetype + authority_ceiling — m-5-F1) and the m-4-F4 operator-on-template-spawn authorship widening, BOTH as prose citing the §C4 C1/C2 Step-3 carry-conditions (R2 at column grain; the any_row deviation coupling). seat_scope is NOT set on this row (design decision): seat_scope's render semantics are option-lists (render.go:124-150 — meaningless for a row_array, my audit's T4); fill-time authorship scoping for routing rows is a Step-3 carry, stated in the annotation. Rides m-4.implementer's (f)+(a) approve — an INTEGRATION gate, not this design's gate.

**D-5 (the declared_deviated byte):** the column vocabulary is the registry bool convention `"yes"`/`"no"`, and the any_row predicate value is `"yes"` — NOT `"true"`. This deliberately avoids re-importing the DEF-1 byte-mismatch class (validate.go:64's `"true"` vs bool `["no","yes"]`) into a new predicate. Explicit fixture leg in §7.

## 6. New rows — Block D: settled singles (2 rows) + the 4 in-place edits

**resolves_gate (M-3(j); rev3 — visible_when added, the IMPL blocker fix):** `id_ref` · owner free_text · fill free_text · `required_when: {"all_of":[{"record_kind_in":["gate_resolution"]}]}` (the disposes_owed/F-GATE-2 pattern, registry.json:118) · **`visible_when: {"any_of":[{"seat_is":["operator"]},{"role_in":["operator"]}]}`** — this row is NOT a dormant consumer row: it is the LIVE verdict-path field (consumed today, engine/submit.go:70/:216-245 — the DEF-5 retro-declaration), and the settled M-3(j) semantics are "operator-seat-scoped Step-1", so it renders ON the operator form (the live gate-resolution affordance) and on NO other seat's form. Both atoms are render-evaluable (render ctx carries Seat — render.go:34); the operator binding carries Name/Role "operator" (binding convention, s2 r3; cmd/frank/main.go:112). Known grain (annotated): predicate atoms match Name/Role, not the IsOperator flag (scopeOptions honors all three, render.go:128-135) — if a future operator binding ever carried a different name/role, only VISIBILITY would miss (the record_kind gate_resolution seat_scope narrowing still governs what the seat can DO). gate_referenceable false · lineage_role none. Annotation: live consumption = classifyVerdict (engine/submit.go:216-245: must equal PARENT_DISPATCH_ID; referenced record must be accepted + gate-bearing; single-resolution enforced; wake-seat = gate author). **③ interplay stated per the escalation §3(j): this row IS the S2 detector reference — its precision carries ③ semantics** (the referenced-gate-record store lookup, m-7-wired mechanism). Two observed gaps stated as s5-b interface notes, not row changes: (1) classifyVerdict today runs on ANY record with non-empty resolves_gate regardless of record_kind (submit.go:70) — presence-and-class unify post-fold (m-7's grain); (2) no operator-seat check exists in classifyVerdict at 67ee23e — the "operator-seat-scoped Step-1" constraint is not ROW-expressible (option-list seat_scope, T4), but **(rev1)** the record_kind-GRAIN enforcement now lands in this pass: `gate_resolution` goes operator-only in the §6 edit-4 scope (the 053113 adjacent flag), so a non-operator seat has no `gate_resolution` affordance at fill time; any residual submit-path check inside classifyVerdict stays s5-b's ③ wiring.

**gate_category_pick (MR-1, working name — m-2 shapes final name/type at this pass):** enum · enum_set gate_category · owner system · fill computed_result · seat_scope none · gate_referenceable false. Annotation: the original-pick provenance row — preserves the agent's pick when the ③ raise REWRITES the committed gate_category token to the detector's A member (m-7 mechanics, escalation §2); written conductor-side at the Q5 validate locus; consumer = the ODB "raised-because" render + mis-pick audit (m-6). Marked: **m-2-shape confirm**.

**In-place edits (4):**
1. `EVIDENCE_TARGET` (:90): add `required_when: {"not":{"phase_in":[]}}` — the total (always-true) predicate; genuinely Step-1-required, **NOT observe-gated** (M-3(b) + m-2's guardrail: it is intent). **D-6:** the total form is chosen over `phase_in:[all 11]` because required_when only fires on predicate-true and an empty-PHASE candidate would dodge a phase-keyed variant (validate.go:37 + evalOp semantics); `not(phase_in [])` is true on every candidate including empty headers. Parse-verified against the grammar (predicateNot over an empty-values phase_in parses — predicate.go:113-120, :166-173; only slot_in rejects non-empty... i.e. the empty list is legal for phase_in). Fixture leg (c) asserts the empty-header case. Flagged for implementer scrutiny as the one novel predicate idiom this pass introduces.
2. `ACTIONS_GIT_REF` (:111): add `visible_when: {"all_of":[{"layer_present":"observe"}]}` (M-3(c) — they ARE observe-owned; required_when already observe-gated).
3. `FINAL_GIT_STATUS_SHORT` (:112): same as 2.
4. `record_kind` (:113) seat_scope — the OI-S4 fold, **FINAL per the 053113 close (rev1: the earlier pending-m-1 variants are deleted — m-1's dual-confirm landed, stronger than asked):**
   - `genesis`: removed from **EVERY** scope, operator included — genesis is `store.Init` machinery, `FROM = system`, never accepted from the public submit path; a rendered genesis option is incoherent by construction (m-1, 053113 §3; consistency proof: incident/gc_marker are already absent from the form enum).
   - `owed_item`, `owed_disposition`: **operator-only**. Row annotation carries m-1's precision VERBATIM: "a Step-1 SCOPE posture, not permanent class semantics — the provenance axis (owed records are principal-authored via submit, never machinery) is unchanged; owed_disposition is authority-bearing in effect (it discharges exit-gating obligations), so the scope axis gets the pessimistic Step-1 floor. A future cycle granting named seats owed-authoring is a registry/config scope amendment **with m-1 route-back**."
   - `gate_resolution`: **operator-only** — the adjacent flag (053113 §3): it sat in `*` while the settled (j) declares it operator-seat-scoped in Step-1; this closes that inconsistency at the record_kind grain (the resolves_gate ROW-grain scoping stays inexpressible, §6 Block D note — this seat_scope narrowing is the Step-1 enforcement that IS expressible).
   - `disposition`: **operator-only** (pessimistic floor; my class check: NO live mechanical consumer keys on it — validateRecordKind has no arm, engine/submit.go:124-160; grep over internal/, test/, cmd/ finds no record_kind=disposition consumer — and it is principal-authored-via-submit, authority-adjacent (it disposes/resolves), so m-1's principle gives it the pessimistic floor, not the machinery treatment). **m-2 rules in-pass (marked confirm)** per the adjacent flag.
   - `diagnostics`: retained in `*` — unflagged by m-1's adjacent check; noted so the in-pass m-2 check sees the full row.
   - Net: `"*"` = `["diagnostics"]`; `"operator"` = `["owed_item", "owed_disposition", "gate_resolution", "disposition", "diagnostics", "config_change"]`. Every change is a conservative-direction NARROWING (MR-2: no extra pair round; a widening would reverse that and route back to m-1). Compat note for the replay leg: seat_scope narrowing changes fill-time affordance, not record shape — old accepted records replay untouched.

**NOT in this pass (settled exclusions):** no ODB `record_kind` member (Q6); no scope_paths row (R-s5-7); no live routing `record_kind` member; no gate_category_B change; `migrate.Current` stays 1 (migrate.go:11 — no record-shape change, Q4); no fieldspec .go code (R-s5-2); rank1_recommended_bucket / predicate_result / veto not persisted (Q9); accepts_interjection + template schema + posture enum stay m-5 config; gate_category_raised row untouched (M-3(g): stays owner:computed).

## 7. The [VP-W3] enumerated negative dormancy fixture (+ registry-content test updates)

**Files (R-s5-2 write surface):** NEW `test/fixtures/s5_registry_dormancy_test.go` (the fixture) · `internal/fieldspec/registry_test.go` (token-list byte-exact assertions extended: gate_category_A 9 / gate_category 14 with other-last + routing_escalation-before-other, the 10 new enums, row count 83, spot checks per block) · `internal/fieldspec/render_test.go` / `validate_test.go` registry-content cases as needed.

**The sweep (merged planner §5 + implementer §8 axes):** for every cell of {operator(IsOperator) · \*.orchestrator-planner · \*.planner grant-OFF · \*.planner grant-ON · implementer(→\*) · orchestrator-reviewer(→\*)} × {11 PHASE} × {5 CEREMONY_TIER}: `reg.Render(...)` and assert `form.HasField(f) == false` (fieldspec.go:28-31) for the enumerated names. **(rev3) The 38 names split by settled class:** **37 strict-absent on EVERY cell** (35 new rows + ACTIONS_GIT_REF + FINAL_GIT_STATUS_SHORT) + **`resolves_gate` absent on every NON-operator cell AND asserted PRESENT on the operator form** (it is the live verdict-path field, operator-seat-scoped per M-3(j) — a positive live-affordance leg, same pattern as the OI-S4 operator-only scope legs; sweeping it into strict absence was a rev2 transcription error against the settled semantics, caught by the IMPL blocker). No exclusion list for the M-1 ten (idiom (i) covers all of them at the render gate — m-2's blessing untouched; resolves_gate was never one of the ten). Digest determinism asserted per cell (render twice, equal); MERGE-GATE grant pruning axis inherited from the implementer's audit. Plus the OI-S4 scope legs: `genesis` absent from the record_kind options on EVERY seat's form (operator included); owed_item/owed_disposition/gate_resolution/disposition present ONLY on the operator form. Annotation-presence assertions inspect the RAW registry JSON, not loaded structs (FieldSpec has no Annotation field — reviewer note, DESIGN-REVIEW 055207).

**The four legs (m2d AC17/NF-S5) + controls:**
- (a) observe-requireds OFF: a minimal valid candidate (SUBJECT + EVIDENCE_TARGET + fresh form digest) validates with zero `required` violations across all Block A/C rows under DefaultLayers (predicate.go:17-23).
- (b) non-observe requireds STILL block: missing SUBJECT (validate.go:40-42); GRILL_REQUIRED=yes + DESIGN_LOCK_ID present + missing GRILL_LOCK_ID (registry.json:124).
- (c) STRONG variant only (M-3(b)): missing EVIDENCE_TARGET ⇒ `required` violation — including on a candidate with PHASE and CEREMONY_TIER omitted entirely (the D-6 total-predicate hole test).
- (d) layer_present-never-model: parse-negative `{"layer_present":"model"}` rejected (predicate.go:192-199, closed set :385-390); plus a mutated-registry negative — a copy whose predicate names `model_name` fails load ("predicate references non gate-referenceable field", registry.go:197-209).
- Positive control (the gate is the layer, not an accident): at the predicate level (Render hardcodes DefaultLayers — render.go:34 — so the flip is tested where predicate_test.go:64-77 tests it), each Block A/C visible_when fires with `observe: true` in PresentLayers.
- D-5 byte control: the justified_deviation/deviation_reason_code any_row predicate fires on a routing_assignments row carrying `declared_deviated: "yes"` and does NOT fire on `"true"` — pinning the DEF-1-class byte discipline into data.
- Enum floor: delivery_state stays byte-exact `{accepted, rejected, held}` (existing assertion registry_test.go:48 — untouched, re-asserted).

## 8. The §7 payload contract for s5-b (R-s5-6 — they land AFTER this pass integrates)

What s5-b's config_change fixture legs consume from my pass: the new fieldspec member bytes = the full post-pass registry.json; `version: s5-fieldspec-v3`; 83 rows; 24 enums with the §2 token tables byte-exact; ONE digest transition (single atomic content change — old digest = the current member digest at the test store's genesis, new_digest recomputed per engine/submit.go:173-186 over the new bytes). Their adapted [VP-W7] legs (operator-authorized shape · old→new digest no re-genesis · phase-0 genesis-chain acceptance · stale-form re-render + re-rendered success) extend the s4 suite (s4_config_change_test.go:21,:210 precedents) with this payload; the re-rendered-success leg asserts my 38 enumerated names are STILL absent post-re-render and the edited rows (EVIDENCE_TARGET required; record_kind narrowed scope) bite.

DEF-2/(h) interface note (design-note only, s5-b codes it): the typed-REJECT guard (non-empty system/computed-owned headers on lane submissions ⇒ typed violation, never silent-strip; channel-keyed; FROM/ROLE envelope overwrite asymmetry preserved) is what upgrades every Block A/B/D dormant row from render-absent to submit-rejected; my row annotations point at it.

## 9. Confirms ledger (rev1 — the riding legs are CLOSED per 053113; zero open m-x legs)

**CLOSED (folded, cited — nothing pends):** m-1 dual-confirm (owed operator-only + genesis nowhere + the adjacent flag → §6 edit 4 final; `DESIGN-planner-…-052940` via 053113) · m-4.implementer APPROVE (f)+(a) (`DESIGN-REVIEW-implementer-…-052840`; the §C4 C1+C2 registrations stay explicit INTEGRATION gates — cited in D-4's annotation) · m-6.implementer APPROVE of S1/S2/S3 (`DESIGN-REVIEW-implementer-…-052907`; s5-b's lane — cited here only for the resolves_gate/S2 interplay) · MR-2: no additional pair-nod on the narrowings; a scope WIDENING would reverse that and route back to m-1.

**Still open (in-pass m-2 shaping, drop-in scale — none blocks the PLAN lock per 052214/053113):**

| item | who |
|---|---|
| MR-1 final name/type of gate_category_pick | m-2, at my pass |
| D-1 annotation-key mechanics + D-6 total-predicate idiom | m-2, bundled |
| the `disposition` scope ruling (§6 edit 4 — designed operator-only, pessimistic floor) | m-2, in-pass per the adjacent flag |

## 10. Acceptance criteria (for the PLAN this design locks)

1. registry.json loads (Load, registry.go:64-81) with 83 rows, 24 enums, version s5-fieldspec-v3; `go vet` + full battery green.
2. The §2 token tables byte-exact (registry_test.go assertions updated in the same pass — R-s5-2); other stays last; routing_escalation immediately before it; B untouched.
3. The [VP-W3] sweep (§7) green: 330 render cells — 37 names absent on every cell; resolves_gate absent on non-operator cells + present on the operator form (rev3) — + digest-deterministic; legs (a)–(d) + both controls + the OI-S4 scope legs green.
4. All 43 untouched existing rows byte-stable; the 4 edits exactly as §6; the 47-row baseline's replay semantics unchanged (s5-b's zero-loss replay consumes this).
5. Every Block-C row carries m-2's verbatim annotation; the owed rows carry m-1's verbatim widening-route-back annotation; every Block-A/B/D system/computed row carries the DEF-2-qualified honesty annotation ([VP-W1] — no non-lane-writability claim anywhere in rows, tests, or this doc's derived text).
6. **(rev1)** No PRODUCTION `.go` under internal/ or cmd/ modified by s5-a; `internal/fieldspec/*_test.go` registry-content fixtures (registry_test.go, render_test.go, validate_test.go) ARE in-scope per R-s5-2, as §7 requires; `fieldspec_test.go` (classifier grain) stays s5-b's. No row/enum outside this doc's tables.
7. The MR-1/D-1/D-6 + `disposition` in-pass m-2 confirms folded before integration (§9 — the only open items; all riding legs closed).

**Out of scope (restated):** everything in the dispatch OUT list — engine/bounce/migrate/test-replay code, step-(d), transport-fix, the archived store, live record_kind widening beyond §6 edit 4, scope_paths.

## 11. Design decisions log (mine, within blessed constraints)

D-1 annotation-as-JSON-key (§1) · D-2 executable_claim_result named enum (§2) · D-3 decision_deadline as string (§4) · D-4 routing_assignments gate_referenceable:true + no seat_scope + prose columns (§5) · D-5 declared_deviated "yes" byte (§5) · D-6 EVIDENCE_TARGET total predicate `{"not":{"phase_in":[]}}` (§6) · D-7 Block-A uniform double predicate even on never-rendered system rows (§3 — idiom as data, not owner-accident) · D-8 new rows appended after :124 in block order A→B→C→D (existing rows byte-stable in place) · **D-9 (rev1)** `disposition` designed operator-only under m-1's pessimistic-floor principle (class check in §6 edit 4; m-2 rules in-pass) · **D-10 (rev1)** on_timeout as a valueless no-enum_set slot rather than a single-member enum (a declared value would make `hold_and_resummon` registry data when it is J1 policy — M-3(d)'s exact line).

## 12. Fold log

**rev1 — folding s5-a.implementer DESIGN-REVIEW `…-055207` (must-revise; all three blockers verified against evidence before folding, none performative):**
- **Blocker 1 (stale riding legs):** verified — `RECONCILE-orchestrator-planner-20260706-053113.md` predates my design request and closes all three legs. Folded: basis updated; §6 edit 4 rewritten to the FINAL scope shape (owed operator-only with m-1's verbatim annotation; genesis in NO scope; gate_resolution operator-only; disposition operator-only with my class-check + in-pass m-2 mark; diagnostics retained, noted); the fallback/G2 variants DELETED; §9 rewritten as a closed ledger; acceptance 7 updated.
- **Blocker 2 (on_timeout omitted):** verified — my Block B dropped the row my own audit proposed and M-3(d) requires. Folded: row added (§4, D-10); counts 35→36 new / 82→83 total / 37→38 enumerated names propagated through §1, §7, §8, acceptance 1/3.
- **Blocker 3 (criterion 6 vs R-s5-2):** verified — the criterion as written forbade the very test files §7 and R-s5-2 require. Folded: criterion 6 now distinguishes production `.go` (out) from `internal/fieldspec/*_test.go` registry-content fixtures (in), with `fieldspec_test.go` left to s5-b.
- Non-blocking reviewer note folded: annotation-presence assertions read raw JSON, not loaded structs (§7).
- D-1/D-4/D-5/D-6 direction approved by the review after live-code checks — unchanged.

**rev2 — folding s5-a.implementer DESIGN-REVIEW `…-060104` (must-revise; both blockers verified, both were my errors):**
- **Blocker 1 (on_timeout validation claim):** verified — my §4 row note claimed validate.go:50-53 rejects a filled value, contradicting my own T1/DEF-2 analysis (ignorePayloadField skips owner:system/system_only rows BEFORE the enum check, validate.go:31-53/:115-120). Corrected to the accurate boundary: render absence today; submit-path rejection = s5-b's DEF-2 guard; internal value conformance = the later scheduler writer.
- **Blocker 2 (stale §7 count):** verified — §7's registry_test.go instruction still said 82. Fixed to 83; literal-text scan for stale 82/37/35 run clean (the fold-log's historical "82→83"/"35→36"/"37→38" delta statements are the only remaining hits, correct as history).

**rev3 — folding the IMPL blocker `s5-a-impl/BLOCKER-implementer-20260706-070710.md` (verified: a genuine design defect, mine):**
- The implementer proved (E2, failing fixture leg) that the locked resolves_gate row shape renders on every Step-1 form while the locked 38-name sweep asserts it absent — an internal contradiction I authored in rev2 by sweeping a LIVE field into the dormancy list. The settled semantics were never ambiguous: M-3(j) says "operator-seat-scoped Step-1" (escalation RECONCILE 052214 §3(j)), and resolves_gate is the live verdict-path field (DEF-5 retro-declaration), not a dormant consumer row.
- Folded: neither of the blocker's two binary options — the SETTLED third shape instead: the row gains `visible_when any_of(seat_is operator, role_in operator)` (§6 Block D); the sweep splits 37-strict + the resolves_gate operator-only pair of legs (§7); acceptance 3 updated. No new semantics authored ([VP-W2] — this transcribes the settled "operator-seat-scoped" text); m-2's M-1 blessing untouched (resolves_gate is not one of the ten); the record_kind gate_resolution operator-only narrowing (053113 adjacent flag) is the paired fill-time enforcement already in the design.
- Process: ruled pair-side as a bounded transcription-defect correction against settled authority, orchestrator + reviewer CC'd on the ruling relay with a one-relay objection window (the run's standing pattern); the implementer's delta-approve of this amendment rides their next relay.

**rev4 — folding the adversarial review panel (post-IMPL, commit dd8189d; panel = custom team-of-4; verdicts: conformance approve · security must-fix · tests approve · semantics must-fix; two must-fix findings, BOTH annotation-text-only, zero mechanism findings):**
- **F-SEC-1 (security lens, verified two-seat):** the model_name annotation's affirmative "non-lane-writable via the tool surface" clause is false as worded — the tool surface includes submit(), and owner:system rows are validation-skipped (ignorePayloadField) with no header stripping, so a lane CAN write model_name via raw headers until the s5-b (h) guard. The clause also contradicted the disclaimer in the same sentence, and this doc's own §10 criterion 5 ("no non-lane-writability claim anywhere"). Root cause: my §4 phrasing — corrected above; the registry-byte fix rides the implementer's REVIEW-FOLD.
- **F-SEM-1 (semantics lens):** the record_kind annotation paraphrased m-1's VERBATIM-marked route-back line, dropping two semantic parentheticals ("(owed records are principal-authored via submit, never machinery)" and "(it discharges exit-gating obligations)"). Restore verbatim — the fold.
- **Punctuation ruling (F-SEM-2/3, planner):** VERBATIM-marked quotes (m-1's owed line; m-6's on_timeout floor incl. the "(§J1)" ref; m-2's Block-C line — already exact) are BYTE-verbatim, em-dashes included; my own DEF-2 honesty phrasing is not a marked quote and may stay ASCII. Folded in the same annotation touch.
- Tracked (not this pass): the security lens's optional note — classifyVerdict has no operator-seat check and keys on resolves_gate presence regardless of record_kind (pre-existing at 67ee23e; correctly deferred to s5-b's (h)/③ wiring; already annotated on the resolves_gate row). Travels to the orchestrator in my review synthesis.
- Optional (implementer discretion): the tests lens's two one-line fixture comments (grant-axis inertness; positive-control cross-reference).
