## DESIGN — s3-form design dispatch: the S3 form system (registry · dissolution · replay · versioning), against the reconciled audits

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: s3-form-design
PARENT_DISPATCH_ID: s3-form-audit
RUN_ID: s3
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no for this dispatch — the grill (below) is operator-facing by design; the S3-close sign-off stays the operator's separate gate
GRILL_REQUIRED: yes
FROM: s3.orchestrator-planner
TO: s3-form.planner
CC: s3-form.implementer, s3.orchestrator-reviewer, operator
SUBJECT: DESIGN dispatch — design the S3 implementation against the LOCKED m-2 spec + twelve hard constraints from the reconciled audits; GRILL_REQUIRED yes (grill with the operator pre-lock; GRILL_LOCK folds into DESIGN_LOCK_ID); three question threads pending — draft provisionally, lock on their answers

### What this is

Your DESIGN phase for Slice-3. The paired audits are reconciled — **full agreement, zero blocking spec gaps** (ledger entry of record: `docs/sprints/2026-07-04-s3-slice-3/RECONCILE.md`, frank/ at b800201). The spec is LOCKED (m-2 design-of-record + ARCH §C4; escalate via me, never amend); this phase designs the **code** — data model, evaluator, engines, fixtures — not the contracts. Implementer is CC'd as design-challenger throughout; the formal DESIGN-REVIEW request goes TO `s3-form.implementer` per the lineage-gate discipline (DESIGN_DOC_ID: `s3-slice-3-design`; the eventual gated PLAN parents the approving DESIGN-REVIEW).

### GRILL_REQUIRED: yes — run design-grill with the operator BEFORE the design lock

Triggers on record (all four real): still-open work at medium tier; cross-domain boundary contracts (m-1 lineage/store seams, m-7 config seam); hard-to-reverse data decisions (registry carrier + on-disk shape, the typed-header representation, config member shape); several downstream choices hanging on the three routed question threads. **Grill agenda floor:** the typed-header/record representation decision (Headers is `map[string]string` today — row_array/object/id_ref fields need a carrier; this shapes every stored record); the registry artifact shape + its version identity; the disposition-table artifact form (in-repo generated table + report-matches-generated, per the S1 replay pattern?); the replay adjudication vocabulary (probe-(c)) if the guide answer has arrived; on-disk commitments generally. **Grill fence:** no re-opening of c1–c6 operator-grilled locks or the S1/S2 closed designs; resolved-by-guide/consult/master rows enter as resolved, not re-asked; an answer that would amend locked text escalates to master. GRILL_LOCK_ID folds into DESIGN_LOCK_ID; no design lock, no DESIGN-REVIEW-consumed-toward-PLAN, and no PROCEED-TO-PLAN until the GRILL_LOCK exists and pending-thread deltas are folded. Drafting proceeds provisionally meanwhile.

### Twelve hard constraints (from the reconciled audits; each lands in the design or the design says why not — silence fails my reconciliation)

1. **Registry data model = the full m-2 §4 column set** ({layer, owner, type, enum_set, gate_referenceable, seat_scope, required_when, visible_when, fill_constraints, consumers, lineage_role}) as a versioned pinned-config artifact; render + validate iterate the rows — no hardcoded field list survives on the live path. The §5 bounded predicate evaluator (closed atom set + all_of/any_of/not; `slot_in` reserved-shape, no Step-1 values; `layer_present` context atom) is pure data-driven boolean — never Turing-complete.
2. **Fill-time authority as registry data:** per-field seat_scope matrix; forbidden options absent from the rendered form AND re-validated by set-membership at submit. **F-P3 lands here:** `canGrant` today includes `orchestrator-reviewer` (fieldspec.go:109) — contradicts m-2 §11.2 ("operator/orchestrator form only") + the upstream grantor set; the S3 seat-scope data excludes the reviewer seat from grant/merge-gated affordances. If the design defers the fix, it becomes a typed owed item through the live S2 mechanism — not a silent carry.
3. **Disposition table built CODE-FIRST** from the actual relay-lint.py assertion inventory (1448 lines, upstream-release path in the ROADMAP) — every assertion site gets a row; the §10 map is the classification authority, not the row source. **The :840-868/:870-873 addressed-token rows are named explicitly** (the range the §10 map does not cite — reconciled finding, verified): grantor/phase halves dissolve into fill-time authority (m-2 §11.2), TO-cardinality survives as form-validation. No silent drops; every *obsolete* row names the concrete vanished surface (VP-W; the only citable ground = the vanished hand-authored-markdown channel per §10a + §8 strictness + the one-channel ROLE/FROM stamp — anything else escalates before S3 close).
4. **FULL replay executed against the real validate/lineage path** — all 243 corpus fixtures; the 96 expected-FAIL side lands caught-or-genuinely-obsolete; **the 50 expected-OK side lands NOT-over-rejected** (the non-overblocking leg — a gate dimension the exit wording doesn't name; carry it as fixtures). Reuse the corpus walker + report-matches-generated pattern; replace the name-heuristic classifier (`uncovered-S3` bucket dies). The step-gated-check adjudication rule (probe-(c)) is **provisional pending s3-guide-q1 Q2** — design the harness with an explicit evaluation-context parameter so either answer folds without re-architecture.
5. **The lineage engine = every §10c row** (design-review walk; pair-Planner dispatch walk; non-addressee IMPL trap; merge-claim lineage; OUT→IN scope-flip drift; orchestrator-review visibility + run-level waiver; the graph/index substrate). **F-P6 lands here:** edge-absence becomes a structural error per class (protocol :88 — for dispatch-class records a missing parent edge is itself the error); today's absent-PARENT-passes stays correct only for non-gated classes. **The S1 grant-narrowing carry is named explicitly in the design** (conditional pair-Planner delegated-dispatch rendering + the m-2 lineage walk — the s1 RECONCILE ratification lands here; it may not silently narrow).
6. **F-P1 lands as an architecture constraint:** the lineage engine + obligation completion ride incrementally-maintained in-memory tables (rebuilt at recovery, maintained on the single-threaded commit loop) — no per-submit full-store re-read/re-checksum. The store API stays the locked verbs (Records/Read/Project); engine-side tables over existing verbs is the recommended shape, but the choice is **m-1 fidelity-reviewed at PLAN time** regardless.
7. **schema_version + migrators:** read-time migrator registry + chain application above a migration-agnostic store; a fixture-proven v(n)→v(n+1) walk; **[VP-W] fixtured refusal/bounce legs** (unknown/future, unversioned, mismatched ⇒ bounce/re-render, never silent coercion); NO backward/downgrade migrator (introducing one = scope expansion ⇒ escalate); submission is always current-version (strict form-only submit).
8. **R2 per-column negatives + the GRILL_REQUIRED row:** `gate_referenceable: bool` (default false) as first-class registry data; the predicate grammar rejects any `field:<id>` naming a false-flagged column — negative fixtures per column class (grain provisional pending s3-guide-q1 Q3); the GRILL_REQUIRED FieldSpec row lands with guide-confirmed owner/type/values (pending Q1). Neither is expressible until constraint 1's column set exists — sequence inside the registry work.
9. **Re-render/drift: promote, never rebuild.** The digest mechanism + "re-render" bounce class exist but are dead-pathed (submit.go:33 passes `formDigest=""` — reconciled V6). Wire a live digest producer through the served surface (shape provisional pending s3-consult-m7 Q3) + a changing-registry fixture (simulation posture pending Q2/scope-q1). This closes the mandate item that has no dedicated exit-gate line — the design names its fixture explicitly.
10. **F-P4 closes by construction:** registry-driven validate enum-checks every enum field (CEREMONY_TIER, EVIDENCE_TARGET, HUMAN_GATE_REQUIRED, gate_category, DESIGN_RECORD_KIND, verdicts — the full §10b set), so the replay's enum fixtures are caught by the real path, not special-cased. The §J2 A-set collapses to ONE source (the registry/config; the hardcoded `isAGateCategory` copy in lineage.go:66-77 dies or derives).
11. **I-PH on every new surface:** all registry/predicate/migration/lineage errors route through bounce.Format's field:class grain; config.Load's path-embedding error class never reaches a seat surface; the P1 grep-family fixture extends to registry errors + re-render bounce text + migration refusals.
12. **Honesty framing in code + doc:** tool-mediated confusion-resistance (affordance, not access) + D5 residual beside every exclusivity-shaped claim; the dissolution claim is proven by the executed replay, never asserted; provenance + transport, not verified work (observe is Step 2); CQ-1 step-gate — no Step-1 form requires an observe-owned field, while every system/lineage/form-owned required (EVIDENCE_TARGET included) stays required.

### Provisional-pending threads (draft against them; de-provision on answers; nothing locks before scope-q1's ruling and the guide's Q1/Q2)

- `s3-guide-q1` → m-2.planner: GRILL_REQUIRED row owner/type/values; step-gated replay adjudication; R2 fixture grain.
- `s3-consult-m7` → m-7.planner: config member shape; the three version/digest axes; serving the rendered form over the existing socket (+ where render sits vs the loop).
- `s3-scope-q1` → master.orchestrator-planner: the §7 config-change record — build-in-S3 vs defer-with-fresh-store-posture. Draft under (a)-shaped assumptions (fresh `store.Init`, restart-based drift fixtures, fresh-store qualifier stated); fold the ruling.

### Scope boundary (unchanged from the ROADMAP — verbatim authority)

OUT: MCP wire-up · observe/evidence (Step 2) · routing execution (Step 3) · S4 consumer-field content (expression capacity in, content out) · TUI/runtime · any locked-design amendment · any edit under ../master or ../extracted. m-1 fidelity + m-7 consult boundaries per the ROADMAP. Any OUT-touch or amendment need escalates via me to master.

### Deliverable

`docs/sprints/2026-07-04-s3-slice-3/designs/s3-slice-3-design.md` (DESIGN_DOC_ID: `s3-slice-3-design`) — the implementation design: data model, evaluator, engine shapes, the disposition-table + replay harness design, fixture inventory (exit-gate-line → fixture map), the twelve constraints landed or argued, provisional sections marked. Then: design-grill with the operator (agenda floor above) → GRILL_LOCK folded → DESIGN-REVIEW request TO s3-form.implementer (per the lineage-gate discipline: review parents your DESIGN relay, same DESIGN_DOC_ID) → completion SITREP to me with the approving review. Battery must stay green at every design commit (docs-only commits expected; state it anyway). Relay + INDEX discipline as established.

Operator-judgment items: the grill itself (agenda floor above); the scope-q1 ruling is master's/operator's; D5 residual restated (accepted, on record).

ACTIONS_GIT_REF: none — dispatch relay only; this file + an INDEX row under gitignored .relays/; no tracked-file edit in this action.
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at b800201, the reconciliation ledger commit of record)
