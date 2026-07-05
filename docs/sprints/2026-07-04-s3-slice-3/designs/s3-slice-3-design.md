# S3 Slice-3 — Design of the full form system (the code shape)

**DESIGN_DOC_ID:** `s3-slice-3-design`
**Owner:** s3-form — design-lead `s3-form.planner` · adversarial design-reviewer `s3-form.implementer`
**Status:** r2 — operator grill RUN 2026-07-04, **GRILL_LOCK folded at §8** (`s3-grill-s3-form`): Q1 typed-header carrier = canonical JSON-in-string (operator); Q2 disposition-table = generated pair (operator); S2-store fate = stated consequence (planner judgment, operator-deferred); the [P-GRILL] tags resolve accordingly. Still provisional: **[P-GUIDE]** (`s3-guide-q1` answer not yet on disk) and the **[P-M7]**/**[P-SCOPE]** drafted assumptions (thread replies on file TO the slice orchestrator; they enter via the orchestrator's fold). Per the dispatch: no design lock, no DESIGN-REVIEW-consumed-toward-PLAN, no PROCEED-TO-PLAN until those deltas fold. (r1 = pre-grill draft, main@c149b71.)
**Date:** 2026-07-04 · **Tier:** medium · **Evidence:** E1 (locked-spec + code cites) / E2 (battery + oracle runs from the reconciled audits)
**Basis:** reconciled paired audits (RECONCILE.md 2026-07-04 third entry; planner `s3-form-audit/AUDIT-planner-20260704-170105.md` + implementer `…-162725.md`); the twelve dispatch constraints (each landed below — §3 is the landing map).
**Locked inputs (never designed here):** the m-2 design-of-record (§3/§4/§5/§6/§8/§9/§10/§11/§12/§17/§18); ARCH §C4 (+ §C4.3 claim boundary, the owed-carry ledger); the m-7 engine design (§2/§6/§7/§8); the m-1 store contract (verbs, homes, stamping); the S1 r5 + S2 r4 closed designs. Line refs `m-2 :N` / `m-7 :N` / `relay-lint.py:N` as in the audits.

**Claim boundary (held in every sentence):** still provenance + transport, not verified work (observe is Step 2). S3 adds determinism + fill-time authority — not evidence, not done-ness. Fill-time authority is **tool-mediated confusion-resistance** (removes affordance, not access); **D5 residual** beside every exclusivity-shaped claim: a same-uid shell-routed seat — confused or malicious — bypassing the courier is the accepted Step-1 residual (ARCH §C4.3). The dissolution claim is proven **by the executed replay**, never asserted. I-PH on every new surface. CQ-1: no Step-1 form requires an observe-owned field; every system/lineage/form-owned required (`EVIDENCE_TARGET` included) stays required. *(Dispatch constraint 12 lands here and in §10.)*

---

## 1. Scope

Build exactly: the full FieldSpec registry (m-2 §4 column set, versioned pinned-config artifact) + the §5 bounded predicate evaluator + registry-driven render/validate with fill-time authority + the §10c lineage engine (all seven rows, incl. the S1 grant-narrowing carry) + the code-first 62-check disposition table + the FULL 243-fixture replay (both legs) + `schema_version` migrator registry with refusal legs + the R2 per-column negatives + the GRILL_REQUIRED row + the live re-render path — riding the S1/S2 engine (loop, store, recovery, obligation) unchanged except where a constraint names the seam.

Out (escalate via s3.orchestrator-planner, never absorb): the ROADMAP OUT list verbatim (MCP wire-up · observe/evidence Step-2 · routing execution Step-3 · S4 consumer-field *content* — expression capacity in, content out · TUI/runtime) · any locked-design amendment · any edit under ../master or ../extracted · the §7 config-change record **unless** `s3-scope-q1` rules it in **[P-SCOPE]** (drafted OUT here: fresh-`store.Init` posture, restart-based drift fixtures, the fresh-store qualifier stated wherever registry replacement is described).

## 2. Decisions

### D-1 Registry data model — the full m-2 §4 column set as versioned pinned-config data (constraint 1)

`internal/fieldspec` is rebuilt around one Go record mirroring m-2 §4 :49-67 byte-for-byte in column names:

```go
type FieldSpec struct {
    ID                string   // canonical field name
    Layer             string   // envelope | header | body
    Owner             string   // system | seat_scoped_enum | agent_enum_pick | free_text
    Type              string   // string|enum|bool|int|address|address_list|row_array|id_ref|evidence_ref|object|text
    EnumSet           string   // named-enum ref (registry.NamedEnums) or "" (inline list in Options)
    Options           []string // inline enum members (byte-exact where locked)
    GateReferenceable bool     // DEFAULT false (m-2-F5)
    SeatScope         map[string][]string // seat-pattern -> allowed options (fill-time-authority matrix)
    RequiredWhen      *Predicate // §5 bounded vocabulary (D-3)
    VisibleWhen       *Predicate
    FillConstraints   string   // monotonic|system_only|seat_allowed_values|computed_result|observed_value|parent_picker|recipient_picker
    Consumers         []string // courier|form_renderer|form_validator|lineage_engine|... (X-* = empty, enforced)
    LineageRole       string   // none|parent_edge|grant|verdict|lock_id|action_report|merge_claim|routing_ref
}
```

- **Carrier:** `config/fieldspec/registry.json` v2 — `{registry_version, named_enums{name→[]tokens}, fields[]}` — the SAME single pinned member **[P-M7 Q1: one larger member assumed; per-domain-section shape folds if m-7 answers otherwise]**. It rides the S2 store-root pinning + genesis digest unchanged (F-M1-3 placement; `configTarget` names stay `engine`/`fieldspec` — zero engine-code member changes under this assumption).
- **Registry version identity [P-M7 Q2 + P-GRILL]:** `registry_version` lives INSIDE the member (content-versioned data; the config digest pins bytes, `schema_version` versions records — the m-2 §18 :377 three-axis layering, one axis per mechanism, no conflation).
- **Named enums carry the byte-exact locked sets:** PHASE (11), AUTHORITY (9), CEREMONY_TIER (5), EVIDENCE_TARGET (4), `gate_category` A(8)/B(4) per §J2 (`routing_escalation` NOT a member — the owed §C4 carry, cited not added), grant `{dispatch-impl, dispatch-merge}`, `delivery_state` `{accepted, rejected, held}` (system-only, declared for completeness, never rendered).
- **The §J2 A-set collapses to ONE source (constraint 10):** `lineage.isAGateCategory` (lineage.go:66-77) dies; the lineage engine and the loop take a registry-backed classifier handle (constructed at startup from the pinned registry). One byte-exact set, one home; a fixture asserts no second literal copy of any A-token set exists outside the registry (grep-class, S1 SWEEP idiom).
- **The S1 flat 6-enum `Registry` struct and its hardcoded `Render`/`Validate` die.** No hardcoded field list survives on the live path; render + validate iterate rows (D-4/D-5).
- **Step-1 field rows shipped:** the S1 D-5 set (envelope: FROM/ROLE/relay_id/DISPATCH_ID/timestamp/schema_version/certification-null/PARENT/TO/CC; headers: PHASE/AUTHORITY/CEREMONY_TIER/EVIDENCE_TARGET/HUMAN_GATE_REQUIRED/gate_category/gate_category_raised/grant/delivery_state/failing_edge/SUBJECT; body + X-*) **plus** the v2.8.8 header vocabulary the exit gate names: PARENT_DISPATCH_ID (envelope home declared, D-6 realization), DESIGN_DOC_ID/DESIGN_LOCK_ID/DESIGN_RECORD_KIND/DESIGN_REVIEW_VERDICT (lineage_role: lock_id/verdict), IN_REPLY_TO (display, consumers:[]), SCOPE_DIFF/SCOPE_DIFF_RESULT/FOLD_SCOPE/FOLD_SCOPE_RESULT/ESCALATION_SCAN rows (row_array — D-2 carrier), ACTIONS_GIT_REF/FINAL_GIT_STATUS_SHORT + the observe-owned set (declared with the CQ-1 step-gate predicate — present, never Step-1-required), record_kind (system-validated tokens incl. the S2 five), ROW_TRUTH_CHECK + evidence row-arrays, ORCH_REVIEW_WAIVER (operator-FROM-only validity, D-6 row 6), GRILL_REQUIRED (D-10 **[P-GUIDE]**). Observe/routing/archetype/ODB *content* fields stay S4-OUT — where §12 consumer slots would be foreclosed by a Step-1 decision, the row is declared reserved-shape with no values (the §17.3 pattern), never given semantics.

### D-2 Typed-header carrier — canonical-JSON-in-string over the existing `Headers map[string]string` **[P-GRILL — the agenda's first item; recommendation below, decided at grill]**

Structured types (`row_array`, `object`, `address_list`) are carried as **canonical JSON strings inside the existing `Headers map[string]string`**; the registry row's `Type` drives parse + validation at submit and at read. The envelope struct (record.go:16-25) is unchanged; new headers are additive.

- **Why (recommended):** zero record-shape movement — S1/S2 records stay `schema_version: 1`-valid under the §9 additive rule (new optional headers = minor; ignore-unknown holds; no real migrator is forced, keeping "zero migrators registered" honest); the checksum canonicalization (record.go:66-73) is untouched; m-1 fidelity surface stays minimal (no envelope/store change for the carrier itself).
- **Canonical encoding rule:** one marshaler (sorted keys, no insignificant whitespace) lives beside the registry; validate re-canonicalizes and byte-compares so two encodings of the same rows cannot alias distinct checksums.
- **Rejected: `Headers map[string]any`** — a store-record shape change (every existing record's parse path + checksum canonicalization moves; forces the major-bump/migrator machinery on day one for zero expressiveness gain). **Rejected: parallel `StructuredHeaders` section** — two homes for one layer; every consumer grows a two-place lookup; §10's "typed row-arrays" need typing at the VALIDATION grain, not the carrier grain.

### D-3 The §5 bounded predicate evaluator (constraint 1, second half; R2 grammar; CQ-1)

`internal/fieldspec/predicate.go`: a typed AST parsed from registry JSON — atoms `phase_in`, `ceremony_tier_gte`, `authority_in`, `seat_is`, `role_in`, `slot_in` (**RESERVED**: shape parses, any concrete Step-1 value = registry-LOAD error), `record_kind_in`, `scan_result_in`, `claims_action`, `field:<id> <op> <value>` (op ∈ ==, in, present), `any_row:<array>.<field> <op> <value>` (bounded existential, no nesting), `layer_present:<layer>` (∈ {store, form, lineage, observe}); combinators `all_of`/`any_of`/`not`. Pure boolean over (candidate fields × context); no calls, no loops, no reflection.

- **R2 grammar enforcement AT LOAD (constraint 8 half):** a `field:`/`any_row:` atom naming a column whose `GateReferenceable` is false (the default) **fails registry load** — the predicate never exists at runtime. This is m-2 §5 :99's allowlist read off first-class data (m-2-F5), grammar-rejected exactly as AC14 :239 requires. Unknown atoms load-reject (the S1 D-6 parse-rejected posture, kept).
- **CQ-1 step-gate:** observe-owned rows (FINAL_GIT_STATUS_SHORT, ACTIONS_GIT_REF, `achieved_evidence`, the `*_RESULT` set — m-2 §5 :96) carry `all_of(<base>, layer_present:observe)`. The engine's Step-1 present-layers = `{store, form, lineage}` ⇒ never required live; the context is a **constructor parameter** so the replay harness can evaluate under `+observe` (D-8, probe-(c) **[P-GUIDE Q2]**). Guardrail fixture: every system/lineage/form-owned required still blocks (AC17(b)); `EVIDENCE_TARGET` stays required (:97).

### D-4 Render, fill-time authority, and the LIVE digest (constraints 2, 9; F-P2/F-P3/F-P5)

`Render(seat, phase, tier)` iterates registry rows: `visible_when` gates presence; owner/fill_constraints pick the affordance (system → not rendered; hybrid → constrained pick within the system candidate space/floor; seat_scoped_enum → this seat's `SeatScope` options only; agent_enum_pick → full set; free_text → text); monotonic fields render `[floor, MAX]`.

- **F-P3 fixed in data:** the grant/merge-gated `SeatScope` admits **operator + orchestrator-planner patterns only** (m-2 §11.2 :177; the v2.8.8 grantor set relay-lint.py:59/:867). `orchestrator-reviewer` and pair seats get NO grant affordance and NO `merge-gated` AUTHORITY option; `canGrant` (fieldspec.go:108-110) is replaced by the seat-scope data. Negative fixtures both ways: reviewer/pair forms omit the options (introspection); a hand-crafted submission carrying them bounces seat-scope in-loop (set-membership half).
- **Serving the form (F-P5) [P-M7 Q3]:** the per-seat form + `form_digest` are delivered through the EXISTING per-seat socket surface — drafted as: `tools/descriptions` (server.go:211-212) becomes seat-shaped, carrying `submit`'s input schema = the rendered form + its digest (render executed handler-side against the startup registry snapshot; the loop stays the authority — render is advisory, m-7 §8.2). No MCP live-adapter work; the exact wire grain folds from the m-7 consult.
- **The digest goes LIVE (F-P2, constraint 9):** the submit payload gains a `form_digest` echo; `SubmitHandler` threads it (killing the dead `""` at submit.go:33); stale digest ⇒ the existing `re-render` bounce class. **Changing-registry fixture (the mandate item with no gate line):** render a form → restart the store with a modified registry **[P-SCOPE: restart-based drift, fresh-store posture — a same-store registry change requires the §7 config-change record and is drafted OUT]** → old-digest submit bounces `re-render`; fresh render succeeds. Plus the S1-compatible same-registry leg (tampered/garbage digest bounces).

### D-5 Registry-driven form-validation (constraints 2, 10; F-P4 closed by construction)

In-loop `Validate(candidate, seat, formDigest)` iterates rows: required-set via D-3 (evaluator), **enum conformance for EVERY enum field** (PHASE, AUTHORITY, CEREMONY_TIER, EVIDENCE_TARGET, HUMAN_GATE_REQUIRED, gate_category, DESIGN_RECORD_KIND, DESIGN_REVIEW_VERDICT, grant, record_kind — the full §10b set; claude/B9- and C2-class replay fixtures are caught by this real path), seat-scope set-membership, monotonic floors (HUMAN_GATE_REQUIRED ≥ floor; gate_category known-A raise with `gate_category_raised` recorded — existing ClassifyGateCategory semantics re-homed onto registry data), typed parse per row Type (D-2 canonical JSON; a row_array that doesn't parse = a typed violation naming the field), X-* namespace (carried, stored, `Consumers` empty **enforced at load** — an X- row with consumers or a predicate referencing X- is a load reject; the AC4 negative fixture), system-owned fields in the payload **ignored byte-for-byte** (stamp wins — existing seat.Stamp posture, unchanged). Violations keep the field:class bounce grain through the one formatter (I-PH, constraint 11).

### D-6 The §10c lineage engine — all seven rows + the S1 carry (constraint 5; F-P6)

`internal/lineage` becomes an engine over the D-7 tables, run pre-append for every candidate; **blocking edges only for authority-bearing records**. The S1 pessimistic authority-superset (lineage.go:21-41) is KEPT AS THE FLOOR and never narrows; the precise classes ride `record_kind` + `LineageRole` on top.

1. **Design-review walk:** a candidate carrying `DESIGN_LOCK_ID` + `DESIGN_RECORD_KIND: design-doc` must chain: parent → an accepted `DESIGN-REVIEW` with `DESIGN_REVIEW_VERDICT: approve` + matching `DESIGN_DOC_ID` → whose parent is the same-owner accepted `DESIGN` carrying that `DESIGN_DOC_ID` (m-2 §10c :166). `audit-record` kind valid only when no same-owner DESIGN with that DESIGN_DOC_ID exists in the accepted graph; `direct-override` valid only from operator/orchestrator-planner-stamped seats (registry seat-scope on the kind's options — fill-time + validate halves).
2. **Pair-Planner dispatch walk (THE S1 GRANT-NARROWING CARRY, landed by name):** a `grant: dispatch-impl` candidate from a pair-planner-stamped seat must chain: parent → an accepted approving `PLAN-REVIEW` → whose parent is that planner's accepted `PLAN` addressed (TO) to the implementer being dispatched. **Edge-absence is a structural error for this class** (protocol :88; F-P6) — absent PARENT passes only for non-gated classes. Operator/orchestrator-planner grants bypass the pair walk (override path, per v2.8.8).
3. **Non-addressee IMPL trap:** an IMPL-phase action-claiming candidate must parent a dispatch whose TO includes the submitting seat.
4. **Merge-claim lineage:** a merge-claiming candidate (`ACTIONS_GIT_REF` merge-ref class / `grant: dispatch-merge` consumption) requires an earlier accepted same-DISPATCH_ID MERGE-GATE grant from the grantor set.
5. **OUT→IN scope-flip drift:** for candidates carrying typed SCOPE_DIFF/FOLD_SCOPE rows under `ROW_TRUTH_CHECK: required`, a path recorded OUT in an earlier accepted record of the same DISPATCH_ID flipping to IN bounces (drift class).
6. **Orchestrator-review visibility gate:** an orchestrator-planner-stamped candidate in the broad SET must carry the run's reviewer in TO/CC unless an accepted operator-FROM record with non-empty `ORCH_REVIEW_WAIVER` exists (table-backed run-level query).
7. **Graph substrate:** by-dispatch/by-relay indexes + accepted-graph views = D-7 tables (the dispatch_id_map/relay_order_key analogs, in-memory).

**Conditional pair-Planner grant RENDERING (the carry's render half):** the pair-planner form renders `grant: dispatch-impl` **only when** the D-7 lineage tables show an accepted approving PLAN-REVIEW parented to that planner's PLAN (the walk's own precondition). Render reads a loop-owned snapshot (advisory); the in-loop walk re-validates authoritatively — a stale render can never grant (m-7 §8.2 posture; render sits handler-side, the loop stays the authority) **[P-M7 Q3 grain]**.
**PARENT realization [flagged for m-1 fidelity at PLAN]:** the registry declares PARENT_DISPATCH_ID at its locked envelope/`parent_picker` home (m-2 §12 :185); the S3 realization renders a conductor-derived candidate set (the seat's delivered/accepted horizon from the D-7 tables; free-typed parent outside the set bounces — strengthening today's resolve-if-present) — the candidate-set derivation rule is written as an m-1 PROPOSAL, not a decision, and goes in the PLAN-time fidelity packet whole (VP watchpoint: lineage movement is m-1-fidelity even inside m-2-owned modules).

### D-7 Incrementally-maintained engine tables (constraint 6; F-P1 as architecture)

One `internal/tables` (extending `obligation.Tables`/`BuildTables`): records-by-relay-id + by-dispatch-id, accepted-graph parent edges, lineage-class indexes (grants, verdicts, locks, merge-gates by DISPATCH_ID), gate/park/outbox/owed completion indexes, waiver set, parked lanes, outcome-by-intake, content-hash. **Built once at recovery phase 3 (the S2 slot), maintained incrementally by the commit loop on each commit** (single-threaded ⇒ no locking; rebuild IS recovery — the store stays truth, tables are caches). `lineage.Check`, `classifyVerdict`, `validateRecordKind`, and `obligation.CompleteAuto` all move onto tables; **no per-submit `st.Records()` full-store re-read/re-checksum survives on the live path**. Store API untouched (the locked verbs); the query shapes go to m-1 fidelity at PLAN regardless (constraint 6 verbatim). S1/S2 fixtures (C7, H, S2-O5, recovery suite) run unchanged over the refactor — the no-regression proof it rides on (S2 D-7 precedent).

### D-8 Disposition table (code-first) + the FULL replay harness (constraints 3, 4; probe-(c))

- **The table artifact:** `docs/sprints/2026-07-04-s3-slice-3/results/disposition-table.md` + a machine `dispositions.json` beside the harness — **one row per assertion site in relay-lint.py's actual inventory** (line-anchored), each row = {anchor, check name, §10 class (the map is the classification AUTHORITY, never the row source), disposition ∈ dissolved-form-validation | dissolved-lineage | retained-post-submit | genuinely-obsolete, S3 surface (registry row / engine walk / fixture id), obsolete-ground}. **The :840-848/:850-868/:870-873 addressed-token rows appear explicitly:** grantor-set + phase-binding halves → dissolved into fill-time authority (grant SeatScope + `dispatch-merge` visible only in MERGE-GATE — m-2 §11.2); TO-cardinality (exactly one implementer-role addressee for a grant-bearing record) → **survives as form-validation** (a required_when/cardinality rule on the grant row); duplicate merge-verdict-line → obsolete (prose-phrase surface vanished; the typed grant field is single-valued by construction). A generator test asserts table-matches-inventory (every anchor covered, no silent drops) + report-matches-generated (the S1 replay pattern).
- **Obsolete-ground vocabulary (VP-W, closed):** only {the vanished hand-authored-markdown channel: fences/inline-code · bare own-line tokens · FIELD:-block/row grammar · ambiguous continuations/detached rows · address string grammar · prose-phrase scanning}, + §8 strict form-only submit, + the one-channel ROLE/FROM stamp. Any row needing another ground escalates before S3 close.
- **The replay harness (`test/replay` rebuilt):** for every corpus fixture (243; the walker + report-matches-generated pattern reused; `classmap.go`'s name-heuristic classifier and its `uncovered-S3` bucket DIE), a **typed-equivalent constructor** drives the REAL in-process path (render → submit → validate → lineage on a test store; never a markdown submit path — §8 preserved, importer stays test-only): **expected-FAIL fixtures (96 oracle entries)** ⇒ the equivalent bounces with the disposition row's class, or the row adjudicates genuinely-obsolete with its ground; **expected-OK fixtures (50 entries)** ⇒ the typed equivalent is ACCEPTED (the non-overblocking leg, carried as first-class gate fixtures). Root-mode/lineage fixtures build their graphs by sequential submits, then the offender bounces. The oracle (`check-relay-lint-fixtures.py` EXPECTED, 146 entries) is the expected-verdict source of record.
- **Step-gated checks (probe-(c)) [P-GUIDE Q2]:** the harness's evaluation context is parameterized (D-3); FINAL/ACTIONS-substance-class fixtures are drafted **caught-under-observe-context** — evaluated with `layer_present:observe = true` to prove the dissolved rule exists and fires, while live Step-1 forms never demand the fields; the report labels these rows `caught (observe-context)`. Either guide answer folds by flipping the parameter/label, no re-architecture.

### D-9 `schema_version` + the migrator registry (constraint 7)

`internal/migrate`: `Register(from int, fn Migrator)` + `Apply(rec, current)` walking the chain at **read/projection time** above a migration-agnostic store (no stored byte ever mutates; the view upgrades). **Zero migrators registered in production** — under D-2 the registry replacement is additive header vocabulary (minor per the §9 :126 compat contract; the major-bump trigger did not fire — stated in code doc). **Refusal legs [VP-W], all bounce/refuse, never silent-coerce:** version > current ⇒ typed refusal (unknown-future); version 0/absent ⇒ typed refusal (unversioned — every real record is stamped since S1, so this is corrupt-class, distinct from ErrChecksum); a chain that cannot reach current ⇒ typed refusal (mismatched). Submission is always current-version (strict form-only; the courier stamps — no seat-supplied version survives). **Fixtures:** a test-registered v1→v2 migrator proves the walk (mechanism first, zero real migrators — §9 :127 verbatim); the three refusal legs fixtured; a submit-side leg proves the stamp overrides any payload version. **Wrap point** (engine read-facade in front of `Read`/`Records` consumers) is an m-1 PLAN-time fidelity item (touch list §5.5 of my audit).

### D-10 The GRILL_REQUIRED FieldSpec row (constraint 8; §C4 :487 m-6-F6) **[P-GUIDE Q1 — provisional declaration]**

Drafted: `{id: GRILL_REQUIRED, layer: header, owner: agent_enum_pick, type: enum, options: [yes, no], gate_referenceable: false, required_when: phase_in [DESIGN], visible_when: phase_in [DESIGN, PLAN], fill_constraints: seat_allowed_values, consumers: [form_validator], lineage_role: none}` — the ported v2.8.8 header (default-no posture per the pair-planner skill; the m-6 meeting-lane binding stays m-6's, keyed on the locked phase atom until m-6 binds the field). Owner/type/values FOLD from the guide answer; the row ships only guide-confirmed.

### D-11 R2 per-column negative fixtures (constraint 8) **[P-GUIDE Q3 — grain provisional]**

With `gate_referenceable` first-class (D-1) and load-time grammar rejection (D-3): per-column-class fixtures assert a registry whose `required_when`/`visible_when` names (a) a false-flagged ordinary column, (b) a computed-label column (`gate_category_raised` — the Step-1-present computed field), (c) an X-* id, (d) `any_row:` over a false-flagged row-array field — each **fails registry LOAD** with a typed, path-free error naming field + rule. Model-identity columns (`chosen_model` etc.) are S4 content — the fixture grain at the live-field level folds from the guide answer; the mechanism (flag + load-reject) is field-agnostic by construction, so S4 rows inherit it without new mechanism (expression capacity in, content out).

### D-12 Fixture inventory — exit-gate-line → fixture map (the acceptance skeleton; ids stable for the PLAN)

| gate line | fixtures |
|---|---|
| Registry live end-to-end | S3-E1: a real relay in the full header vocabulary (ROLE/PHASE/AUTHORITY/DISPATCH_ID/PARENT/DESIGN_* …) renders → validates → commits → projects through the real path on a live store; S3-E2: the rendered form reaches a scripted seat over the socket surface + digest echo round-trips |
| Fill-time negatives | S3-A1: reviewer + pair forms omit grant/merge-gated (introspection); S3-A2: hand-crafted out-of-scope submission bounces pre-append (seat-scope class); S3-A3: forged system-field payload ignored byte-for-byte; S3-A4: below-floor monotonic pick bounces; known-A B-pick raised + `gate_category_raised` recorded (S1 G legs re-run over registry data) |
| Disposition table + FULL replay | S3-D1: table-matches-inventory (every relay-lint.py assertion anchor has a row, :840-873 included); S3-D2: report-matches-generated; S3-D3: 96 expected-FAIL → caught-or-obsolete, zero `uncovered-*`; S3-D4: 50 expected-OK → accepted (non-overblocking); S3-D5: every obsolete row's ground ∈ the closed vocabulary (mechanical scan) |
| R2 + GRILL_REQUIRED | S3-G1..G4: the D-11 load-reject legs; S3-G5: GRILL_REQUIRED row present + rendered per its predicate |
| schema_version + migrators | S3-V1: v1→v2 walk (test migrator); S3-V2..V4: the three refusal legs; S3-V5: stamp-overrides-payload |
| Re-render/drift | S3-RR1: stale/tampered digest bounces `re-render`; S3-RR2: registry-change-across-restart drift leg **[P-SCOPE posture]** |
| Lineage engine | S3-L1..L7: one green + one bounce leg per §10c row (design-review chain; pair dispatch walk incl. edge-absence structural + the conditional-render gate closed/open; non-addressee trap; merge-claim; OUT→IN flip; visibility gate + waiver; substrate rebuild-at-recovery) |
| No regression | S3-RE: entire S1+S2 suites green, invariant assertions untouched (mechanical call-site updates only — any assertion change is a review blocker); enum byte-exact; I-PH P1-family extended to registry errors + re-render bounces + migration refusals; owed/recovery/FIFO/GC untouched-and-green |
| materialize-first | any S3 finding meant to be guarded → typed owed record via the live mechanism (F-P3 lands in-slice per D-4; if any deferral emerges, its owed record is the fixture) |

## 3. The twelve constraints — landing map

1 → D-1/D-3 · 2 → D-4/D-5 (F-P3 fixed in data, in-slice) · 3 → D-8 (code-first; :840-873 named) · 4 → D-8 (both legs; context parameter) · 5 → D-6 (all seven rows; carry named; F-P6 structural-per-class) · 6 → D-7 (tables; store verbs untouched; m-1 fidelity at PLAN) · 7 → D-9 (walk + refusal legs; no downgrade migrator anywhere) · 8 → D-10/D-11 (sequenced inside D-1) · 9 → D-4 (promote + wire; both fixture legs; the no-gate-line mandate item gets S3-RR) · 10 → D-5 (full enum sweep) + D-1 (one §J2 source) · 11 → D-5/D-8/D-9 error grain + S3-RE's I-PH extension · 12 → the claim-boundary header + §10.

## 4. m-1 fidelity surface (PLAN-time packet; proposals, not decisions)

1. D-7 table/query shapes over the locked verbs (no store-API change proposed). 2. D-6 PARENT `parent_picker` realization (candidate-set derivation + free-typed-outside-set reject) — the VP-watchpoint item. 3. D-6 lineage-field homes (DESIGN_* et al. as headers; envelope untouched) per the F-M1-1 homes-table precedent. 4. D-9 read-facade wrap point. 5. D-2 canonical-JSON-in-string header carrier (no envelope/checksum change — confirm). 6. record_kind token additions: none proposed (the S2 five stand; the authority floor stays the S1 superset) — confirm. 7. Registry member replacement posture (fresh `store.Init`; digest change = phase-0 wall) **[P-SCOPE]**.

## 5. Provisional register (every [P-*] above, one place)

| tag | section | drafted assumption | folds from |
|---|---|---|---|
| P-M7 Q1 | D-1 carrier | one larger `fieldspec` member, member names unchanged | m-7 consult |
| P-M7 Q2 | D-1 version identity | `registry_version` inside the member; three axes stay distinct | m-7 consult |
| P-M7 Q3 | D-4 serving, D-6 render grain | seat-shaped `tools/descriptions` carries form+digest; render handler-side; loop = authority | m-7 consult |
| P-SCOPE | §1 OUT, D-4/S3-RR2, §4.7 | fresh-`store.Init` posture; restart-based drift fixture; §7 config-change record OUT | master ruling |
| P-GUIDE Q1 | D-10 | GRILL_REQUIRED = header/agent_enum_pick/enum{yes,no}/required-when DESIGN | m-2 guide |
| P-GUIDE Q2 | D-3/D-8 | step-gated checks = caught-under-observe-context, labeled | m-2 guide |
| P-GUIDE Q3 | D-11 | per-column-class grain; live-field additions inherit mechanism | m-2 guide |
| P-GRILL | D-2 (carrier), D-1 (artifact shape/version), D-8 (table artifact form), D-8 (adjudication vocabulary if guide lands first), on-disk commitments | recommendations as drafted | operator grill |

## 6. Open items carried to PLAN (none re-architect)

1. The m-1 fidelity packet (§4) routes before anything store-shape-touching dispatches (F2 condition). 2. Fixture-to-task map + SCOPE_DIFF file list. 3. S1 fixture call-site migration inventory (fixtures constructing the old 6-enum registry move to v2 construction helpers; every invariant assertion stays byte-identical — S3-RE blocker otherwise). 4. The disposition-table generator's line-anchor pinning against the frozen v2.8.8 release path (read-only input, DO-NOT-COPY).

## 7. Rejected alternatives (log)

`Headers map[string]any` / parallel structured-header section (D-2) · JSON-Schema as the registry carrier (re-litigates the locked m-2 §4 carrier decision — not ours to reopen; the bespoke FieldSpec IS the locked choice) · eval-time-only R2 checking (load-time rejection is strictly stronger and matches "rejected by the grammar", AC14) · keeping `isAGateCategory` as a verified copy (two sources for one byte-exact set is the audit's named probe — dies) · narrowing the S1 authority-superset to the precise taxonomy in one step (never-narrower rule; superset stays the floor) · a fourth `uncovered`-style replay bucket (the gate vocabulary is closed at caught-or-genuinely-obsolete; the observe-context label is an adjudication annotation, not a bucket) · per-submit store rescans retained (F-P1) · building the §7 config-change record unbidden (P-SCOPE; scope is master's) · a backward/downgrade migrator (scope expansion — escalate if ever wanted).

## 8. GRILL_LOCK

```text
GRILL_LOCK_ID: s3-grill-s3-form
GRILL_REQUIRED: yes
GRILL_SOURCE:
- plan/design/audit relay read: s3-form-design dispatch (…-170903, GRILL_REQUIRED: yes, agenda floor + fence); reconciled s3-form audits (RECONCILE.md third entry); the twelve constraints
- code/docs inspected: frank/ at 042fcd9..c149b71 (record/store/fieldspec/engine/lineage/channel + test/replay); m-2 §4/§5/§9/§10/§11; ARCH §C4 + §J2; relay-lint.py assertion inventory; the v2.8.8 oracle (146/146, own run)
- questions answered from codebase: no hardcoded field list survives the registry rebuild (D-1/D-5); the S1 authority-superset stays the floor (never-narrower); the digest mechanism exists dead-pathed (promote, not rebuild)
- questions asked operator: typed-header carrier (Q1); disposition-table artifact form (Q2)

Resolved decisions:
- Typed-header carrier — canonical JSON-in-string inside the existing Headers map[string]string; one canonical marshaler (sorted keys) beside the registry; zero record-shape movement, additive/minor under the §9 compat contract, no forced migrator — source operator (grill 2026-07-04, Q1: "Canonical JSON-in-string (Recommended)" selected)
- Disposition-table artifact — generated pair: machine dispositions.json consumed by the replay harness + generated human table at docs/sprints/2026-07-04-s3-slice-3/results/disposition-table.md; honesty legs = table-matches-inventory (every relay-lint.py assertion anchor rowed, :840-873 included) + report-matches-generated — source operator (grill Q2: "Generated pair + repo doc (Recommended)" selected)
- S2-store fate — STATED CONSEQUENCE under planner judgment (operator ruled it a non-decision: "this doesn't matter..... use your best judgement and move on", 2026-07-04): the real S2 store ($HOME/frank-s2-store) freezes at its pinned S2 config, readable forever (phase-0 walls any config swap by design); S3 gate evidence runs on fresh stores; upgrading that store becomes the first customer of the §7 config-change record if/when master mandates it — nothing forecloses it
- Registry member shape / version home / form-serving grain — RESOLVED-BY-CONSULT rows (m-7 reply on file at ../master/relays/s3-consult-m7/SITREP-planner-20260704-171546.md, TO the slice orchestrator; enters this design via the orchestrator's fold, not re-asked here; the §5 drafted assumptions stand provisional until that fold)
- §7 config-change record scope — RESOLVED-BY-MASTER row (ruling on file at ../master/relays/s3-scope-q1/RECONCILE-orchestrator-planner-20260704-171608.md, TO the slice orchestrator; same fold path; drafted fresh-store posture stands provisional until folded)

Pending (grill fence honored — not operator questions, not re-asked):
- replay adjudication vocabulary, GRILL_REQUIRED row values, R2 fixture grain — s3-guide-q1 answer not yet on disk; the P-GUIDE rows stand provisional

Rejected alternatives:
- Headers map[string]any carrier — breaking record-schema change; forces the major bump + a real migrator day-one (operator-rejected via Q1)
- parallel StructuredHeaders section — two homes for one layer (operator-rejected via Q1)
- hand-maintained single .md table / test-artifact-only table — either unexecutable or outside the gate-evidence tree (operator-rejected via Q2)
- upgrading the S2 store inside S3 — drags the config-change-record scope in through the side door (planner-rejected as part of the stated-consequence row)

Still operator-owned:
- S3-close sign-off (charter; exercised separately)
- the s3-scope-q1 ruling is master's/operator's (fold path above)

Design-lock impact:
- DESIGN_LOCK_ID for this doc must reference GRILL_LOCK_ID s3-grill-s3-form; per the dispatch, no design lock / no DESIGN-REVIEW-consumed-toward-PLAN / no PROCEED-TO-PLAN until the pending-thread deltas (P-GUIDE + the P-M7/P-SCOPE folds) land
```
