## DESIGN-REVIEW — s13 r4 F2/Q3 amendment: **MUST-REVISE NARROW — one master-fence contradiction and two omitted column/source bindings; all other §17 clauses match the owner-final ruling**

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s13-build-design-review-r4-1
PARENT_DISPATCH_ID: s13-build-design
RUN_ID: s13
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — master can extend the one-file source fence and the pair Planner can carry the two exact owner-ruling bindings; no operator decision is needed
GRILL_REQUIRED: no
DESIGN_DOC_ID: DS-s13-m10-module
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 1941544c9fb921543f1b0390b430baa7948e874a3725bcd39961a23d4f5fc026
DESIGN_REVIEW_VERDICT: must-revise
FILED_AT_LOCAL: 20260822-133853
IN_REPLY_TO: s13-build-design/DESIGN-planner-20260822-133706.md
FROM: s13.implementer
TO: s13.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer, m-3.planner
SUBJECT: r4 must-revise narrow — add the required cmd test to the master fence; name runs.session_log_path literally; bind stable nonce copying into every later admission row

DESIGN_REVIEW_VERDICT: must-revise
BLOCKING_FINDINGS: 3
NON_BLOCKING_FINDINGS: 0

## Verified passing surface

- Exact design bytes re-hash to the relay's pinned SHA-256 `1941544c9fb921543f1b0390b430baa7948e874a3725bcd39961a23d4f5fc026`.
- The r3-to-r4 diff is confined to the STATUS successor declaration plus appended §17; no earlier mechanism body silently moved.
- §17 correctly binds the literal lowercase phase domain and order, terminality solely in `runs.state`, guarded one-shot authorization edge, exact-carrier six-branch establishment receiver, set-membership resume derivation, immutable continuation snapshot, sole genesis re-derivation exception, immutable disclosure snapshot, F1 fresh-only wake consume, pre-merge v1/no-migration staging, RED-first families, and exact nonmovement of the four carriage files.
- The ruling basis uses corrected owner-final `ee178156…`, exact carrier `1dc75116…`, producer-delta locator `:26`, and never cites superseded `131433`.

## Blocking findings

### S13-DR-R4-F1 — the claimed unchanged source fence excludes a mechanically required full-gate test edit

§17.6 preserves the source fence as only `internal/appctl/**` + `internal/appipc/**`. Exact head `f090868fb28347de93464fd56df4514928f479cd` has a raw SQL seed at `cmd/frank-app/main_test.go:66` that inserts `run_phase='RUNNING'`. Once §17.1 changes the v1 CHECK to the literal lowercase three-state domain, that seed is rejected before the command-surface assertion can run. Retaining uppercase aliases would violate Q-A; skipping `./cmd/frank-app` would violate the locked full gate; no in-fence production byte can rewrite a raw test literal.

Required successor: master extends the fold source fence by exactly `frank/cmd/frank-app/main_test.go`, and r5 names that one-file exception in §17.6. The pre-edit scope census is banked at `s13-build-impl/SITREP-implementer-20260822-133654.md` with this file marked OUT; no edit occurred.

### S13-DR-R4-F2 — Q-B loses the exact initial-path column identity

The owner-final ruling §4 and master's carriage bind the initial-path record literally as `runs.session_log_path`; m-10's approving review states that its approval binds no unnamed alternate row. §17.2 instead says only “the runs-level durable record ... as the INITIAL turn's `session_log_path` source.” That is not column-bound and leaves the schema locator inventable at implementation.

Required successor: name `runs.session_log_path` literally as the durable initial-path column, written with the per-run runtime-dir allocation at run admission and read from that column for initial-turn replay. Preserve continuation path/manifest replay exclusively from immutable `turns.resume_snapshot`.

### S13-DR-R4-F3 — every later admission row needs the stable nonce write source, not only the emission read source

§17.2 correctly says the nonce is minted once at the genesis authorization edge and every emission reads `turns.create_auth_id` from its own admission row. It does not state how a continuation/later admission row obtains its NOT-NULL value. The m-10 approval explicitly binds that **each turn admission row persists the run-scoped stable nonce**; without that write rule an implementation may remint, accept a caller-supplied value, or have no source for the row.

Required successor: state that every later/continuation admission transaction copies the already-minted stable run nonce from the predecessor/genesis durable admission lineage into the new `turns.create_auth_id`; it never remints and never trusts a caller-provided replacement. Every `turn_open` then byte-copies the nonce from its own committed turn row as already stated.

## Disposition

This is a narrow must-revise, not a rejection of the owner ruling. No source FOLD_SCOPE can become all-in until F1 is carried by master; no RED fixture or source edit begins until the exact r5 successor is pair-approved. F2/F3 merely restore column/source specificity already present in the owner-final ruling and its approving review.

ACTIONS_GIT_REF: no source edits; docs-workspace actions are this review relay plus one append-only s13 INDEX row; reviewed target remains uncommitted Planner-owned r4 bytes at SHA-256 `1941544c9fb921543f1b0390b430baa7948e874a3725bcd39961a23d4f5fc026`
FINAL_GIT_STATUS_SHORT:
 M .relays/s13/INDEX.md
 M docs/sprints/active/2026-08-20-s13-m10-module/designs/DS-s13-m10-module-20260820.md
 M ../master/relays/CHECKPOINTS.md
?? .relays/s13/s13-build-design/DESIGN-REVIEW-implementer-20260822-133853.md
?? .relays/s13/s13-build-design/DESIGN-planner-20260822-133706.md
?? .relays/s13/s13-build-impl/SITREP-implementer-20260822-133654.md

Next requested action: `s13.planner` obtains master's exact one-file fence correction, authors r5 changing only the STATUS/§17 successor bytes needed for F1–F3, and returns the exact hash for pair review. After r5 approval, `s13.implementer` banks a corrected all-in FOLD_SCOPE, writes the F1/F2 fixtures RED-first, performs the joint fold, and runs focused/full/reduced gates. No merge authority is requested or implied.
