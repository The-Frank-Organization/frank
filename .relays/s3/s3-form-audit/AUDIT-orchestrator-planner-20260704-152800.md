## AUDIT — s3-form paired independent audit: onboard to the S1+S2 build surface + map the S3 form-system mandate onto it (read-only)

ROLE: Orchestrator Planner
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: s3-form-audit
PARENT_DISPATCH_ID: s3-dispatch
RUN_ID: s3
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s3.orchestrator-planner
TO: s3-form.planner, s3-form.implementer
CC: s3.orchestrator-reviewer, operator
SUBJECT: AUDIT dispatch — paired, independent: onboard to the S1+S2 code (you built neither); map every S3 exit-gate line to the locked m-2 text + current code; census the 62 checks + 243-fixture corpus against the m-2 §10 map; enumerate the m-1 lineage/store-touch surface; 4-bucket verdicts per S3 IN item

**No operator hold is on record for s3** (unlike S1/S2; recorded in RECONCILE.md). This is the first work dispatch of RUN_ID s3.

### What this is

Paired **independent** audits (do not coordinate before writing; reconciliation is my job). Read-only: no code, no design, no plan — findings + evidence. You are auditing the **S2 baseline** (`main`, tag `s2-close` = b322b6d; code surface identical at 7a8b9d7; scaffold commit 125e4fa on top) against the **S3 mandate** (docs/sprints/2026-07-04-s3-slice-3/ROADMAP.md — the s3-dispatch scope verbatim).

### Scope (IN for this audit)

1. **Onboard to the S1+S2 code first — you built neither.** Read both sprint ledgers (docs/sprints/2026-07-03-s1-slice-1/ and …-s2-slice-2/: designs, plans, RECONCILE.md — the deviation rulings, waivers, and fidelity findings live there) and the source (internal/* + test/*). Fresh adversarial eyes are the point (S2's audits found 2 latent S1 races): anything load-bearing-but-fragile under a *changing, much larger* registry is a finding — e.g. the form-digest mechanism under registry evolution, the O(N) store scans in lineage.Check/classifyVerdict under a heavier validate path.
2. **Spec-to-exit-gate map.** Every S3 exit-gate line (ROADMAP) maps to locked m-2 text — §4 (FieldSpec shape incl. `gate_referenceable`), §5 (bounded predicate vocabulary + CQ-1 step-gate + R2 grammar rule), §6 (X- overflow), §8 (strict form-only submit), §9 (versioning/migrators), §10a/10b/10c (the dissolve/survive map), §11 (GATE-1: lineage engine + dispatch/merge authority as seat_scoped_enum), §12/§17.1/§17.6 (consumer fields S3 must *express* without choosing S4 content), §14 ACs, §18 folds — plus ARCHITECTURE §C4 (the owed-carry ledger rows S3 discharges: R2 per-column negatives, GRILL_REQUIRED row) — with file:line cites. Name any exit-gate line you cannot ground in locked text (spec gap ⇒ escalate, never improvise).
3. **The 62-check + corpus census (the disposition table's ground truth).** m-2 §10 classifies 62 checks with line cites against the upstream relay-lint.py (path in the ROADMAP). Verify that census against the actual file: does the ~33/~16/~13 split account for all 62? Name any check present in relay-lint.py that the §10 map does not cover, and any §10 row whose line cite no longer grounds. Census the fixture corpus (243 .md fixtures, 14 categories; check-relay-lint-fixtures.py = the expected-verdict oracle) — which categories map to which §10 class, and which fixtures are *pass-fixtures* (expected-OK) vs *failure-fixtures* (the replay's caught-or-obsolete population).
4. **4-bucket verdict per S3 IN item** (full FieldSpec registry · linter dissolution + disposition table · FULL replay · schema_version/migrators · R2 negatives + GRILL_REQUIRED row · re-render/drift): what do S1/S2 already ship (already-closed — do not rebuild; e.g. the form-digest/re-render seed, the seat-scope `canGrant` mechanism, X-fields carried on the record, `schema_version` stamped, the corpus walker in test/replay, the `held`/gate_category machinery), what is genuinely still-open (the owner/type/required-when/seat_scope/gate_referenceable model, the predicate evaluator, the migrator registry, the retained-check engine), what narrows. **Two named probe surfaces:** (a) the §J2 A-set is hardcoded in internal/lineage/lineage.go `isAGateCategory` AND config-sourced in registry.json — two sources for one byte-exact-locked set; (b) test/replay/classmap.go classifies by filename heuristics with a literal `uncovered-S3` deferral bucket — the S3 replay replaces this with execution.
5. **The S1 grant-narrowing carry.** S1's ratified blocker-2 narrowing (s1 RECONCILE.md 2026-07-03 design-completion entry) explicitly lands **conditional pair-Planner delegated-dispatch rendering + the m-2 lineage walk in S3**. Audit what the full §10c lineage engine needs beyond today's lineage.Check (design-review walk, dispatch-impl walk, merge-claim lineage, OUT→IN flip drift, non-addressee IMPL trap — m-2 §10c rows) and confirm this carry is on the S3 build list or name why it narrows.
6. **m-1 lineage/store-touch enumeration (the fidelity surface — VP watchpoint).** List every touch S3 plausibly makes on: `PARENT`/`parent_picker`/candidate-set derivation, system-filled lineage fields, store *query* semantics (lineage walks are store reads — enumerate which new query shapes the §10c engine needs), record envelope/header homes, `schema_version` migration read-path. Lineage movement is an m-1 fidelity trigger **even inside m-2-owned modules**. Completeness here saves a bounce later.
7. **Trusted-config seam probe (m-7 consult surface).** The full registry rides the trusted config (per-domain sections, single top-level digest, loaded once at trusted startup, restart-only). Audit how today's config.Load + genesis digest-pinning accommodates a much larger fieldspec member + versioned registry, and name exactly which questions (if any) need the m-7 consult rather than improvisation.
8. **Claim-boundary probes.** (a) The dissolution claim is proven by the executed replay, never asserted; every *obsolete* disposition must name the concrete vanished surface (VP watchpoint) — pin what "genuinely obsolete" can cite from locked text. (b) Fill-time authority = tool-mediated confusion-resistance, D5 residual — pin the honest wording for the negative fixtures. (c) CQ-1 step-gate: Step-1 forms must never require observe-owned fields (`layer_present:observe` = false) while every system/lineage/form-owned required stays required, `EVIDENCE_TARGET` included — pin the required-set consequence for the S3 registry. (d) I-PH on new surfaces: registry errors + bounce text carry no canonical store/config/outbox path.
9. **Duplicate/already-built gate** (protocol standard): anything S3-shaped already present, feature-flagged, or dead-pathed — recommend promote/reuse, never rebuild. Named candidates to confirm or refute: the digest/re-render seed, canGrant, the corpus walker, the S2 crash/battery machinery for new commit-path surfaces, the owed-item mechanism for recording any S3 owed findings (materialize-first now has teeth — the OI-S1-F11-SWEEP precedent).

### Out of scope (escalate via me, never absorb)

Everything on the ROADMAP OUT list (MCP wire-up · observe/evidence fields (Step 2) · routing execution (Step 3) · S4 consumer-field *content* — expression capacity in, content out · TUI/runtime) · any locked-contract or design-of-record amendment · any edit under ../master or ../extracted · any code/design/plan work in this phase.

### Deliverable (each seat, independently)

One lint-clean AUDIT relay: `.relays/s3/s3-form-audit/AUDIT-<role>-<YYYYMMDD-HHMMSS>.md`, FROM your seat, TO s3.orchestrator-planner, CC operator. Carrying: the spec-to-gate map (file:line), the 62-check + corpus census result, the six 4-bucket verdicts with PRIMARY_BUCKET lines, the grant-narrowing-carry disposition, the m-1 lineage/store-touch enumeration, the m-7 config-seam questions, the claim-boundary probe answers, reuse/duplicate findings, any spec gaps or escalations, evidence levels per claim (E1 cites; E2 for anything you executed — running `go test -count=1 ./...` yourself is encouraged; 18 packages ok at the baseline, my run on record), and `FINAL_GIT_STATUS_SHORT` (read-only proof). Append your INDEX row (end-of-file, write order).

### Acceptance criteria (for this audit)

- Every S3 exit-gate line mapped to locked text or flagged as a gap — none unaccounted.
- The 62-check census closed: every relay-lint.py check accounted for by a §10 row or named as unmapped; the corpus categories mapped to §10 classes.
- Six PRIMARY_BUCKET verdicts with evidence; no bookkeeping contradicting a recommendation.
- m-1 lineage/store-touch list complete enough for the fidelity packet; m-7 consult questions named or "none".
- Claim-boundary probes answered from locked text, not judgment.
- No file modified anywhere (both trees clean; frank/ at 125e4fa unless the operator advances it).

Operator-judgment items: none new — D5 residual restated (accepted, on record); the MCP deferral stands (S3 is its prerequisite, sequencing note on the master dispatch).

ACTIONS_GIT_REF: none — read-only dispatch authored as this relay file + an INDEX row; .relays/ is gitignored operational substrate, no tracked-file edit in this action.
FINAL_GIT_STATUS_SHORT: none — clean tree
