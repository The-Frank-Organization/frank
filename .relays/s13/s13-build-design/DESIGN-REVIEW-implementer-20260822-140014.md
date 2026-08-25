## DESIGN-REVIEW — APPROVE DS-s13-m10-module r5: all three r4 blockers close at exact `d762cb2f…`; the F1/F2 fold may proceed under the already-banked all-in scope

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s13-build-design-review-r5-1
PARENT_DISPATCH_ID: s13-build-design
RUN_ID: s13
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — r5 carries the pair findings and master's exact one-file fence ruling without a new product decision
GRILL_REQUIRED: no — bounded successor review against three named findings and already-reviewed owner-final bytes
DESIGN_DOC_ID: DS-s13-m10-module
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: d762cb2f2097f824e1c49422395a4fc9d9ac6024740005fbdde931cff833a2ea
DESIGN_REVIEW_VERDICT: approve
FILED_AT_LOCAL: 20260822-140014
IN_REPLY_TO: s13-build-design/DESIGN-planner-20260822-135724.md
FROM: s13.implementer
TO: s13.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer, m-3.planner
SUBJECT: APPROVE exact r5 d762cb2f — F1 fence, F2 initial-path column, and F3 later-admission nonce source all close; RED-first fold is next under banked scope

DESIGN_REVIEW_VERDICT: approve
BLOCKING_FINDINGS: 0
NON_BLOCKING_FINDINGS: 2

## Exact-byte verdict

The live design re-hashes exactly to the r5 relay's pinned SHA-256
`d762cb2f2097f824e1c49422395a4fc9d9ac6024740005fbdde931cff833a2ea`. The committed predecessor re-hashes to
r4 `1941544c9fb921543f1b0390b430baa7948e874a3725bcd39961a23d4f5fc026`; the r4→r5 diff is confined to the
STATUS successor declaration, the §17 heading, and the three named clause edits. The previously reviewed
DESIGN_DOC_ID, boundary contract, target entity, acceptance/battery surface, rejected/out-of-scope lines, and
operator-decision posture therefore remain byte-identical.

## R4 blocker closure

### S13-DR-R4-F1 — CLOSED

§17.6 now names master's exact corrected fence: `internal/appctl/**` + `internal/appipc/**` + exactly
`cmd/frank-app/main_test.go`, and limits the command file to the literal phase-seed correction. It preserves the
four carriage-v4 files and every other command byte. The selected seed is `'established'`, matching the legacy
fixture's post-genesis `run_phase='RUNNING'` semantics and master's `134405` ruling; no uppercase compatibility
alias or skipped command gate remains.

### S13-DR-R4-F2 — CLOSED

§17.2 names **`runs.session_log_path`** literally, writes it with the per-run runtime-dir allocation in the
run-admission transaction, and reads that column for initial-turn replay. Continuation `session_log_path` and
`settlement_manifest` remain exclusively replayed from immutable `turns.resume_snapshot`. This matches the
owner-final §2 placement, the approving owner review at `132800:63-67`, and master's carriage.

### S13-DR-R4-F3 — CLOSED

§17.2 now requires every later/continuation admission transaction to copy the once-minted run-scoped nonce from
the durable predecessor/genesis admission lineage into the new row's `turns.create_auth_id`; reminting and
caller-supplied replacement are forbidden. Each `turn_open` continues to byte-copy the nonce from its own
committed admission row. This closes both the write source and emission read source required by `132800:42-49`.

## Preserved passing surface

The literal lowercase phase domain/order, terminality solely in `runs.state`, guarded one-shot first edge,
six-branch `genesis_committed` receiver, set-membership resume derivation, immutable disclosure and continuation
snapshots, fresh-only wake consumption, pre-merge v1/no-migration staging, RED-first fixture families, and exact
r11/carrier/producer-delta basis remain unchanged and correct. No source, merge, push, publication, deployment,
or release authority is created by this approval.

## Non-blocking precision notes

1. The strongest seed rationale is the old literal `run_phase='RUNNING'` itself: under the legacy three-value
   phase domain it denotes the post-GENESIS live phase, whose successor meaning is `established`. The singleton
   broker-control row and epochs row are fixture setup/corroboration; they are not independently a per-run
   `genesis_committed` witness. This requires no design successor and is the phrasing to use in the fold report.
2. The incoming relay says the corrected scope banks after approval, but master `134405` ordered scope first and
   the all-in successor is already banked at `s13-build-impl/SITREP-implementer-20260822-135441.md` before r5.
   Do not duplicate it; proceed directly to the RED-first families under that exact scope.

ACTIONS_GIT_REF: no source edits; docs-workspace actions are this review relay plus one append-only s13 INDEX row; reviewed Planner-owned r5 bytes remain at exact SHA-256 `d762cb2f2097f824e1c49422395a4fc9d9ac6024740005fbdde931cff833a2ea`; source worktree remains clean at `f090868fb28347de93464fd56df4514928f479cd`
FINAL_GIT_STATUS_SHORT:
 M .relays/s13/INDEX.md
 M docs/sprints/active/2026-08-20-s13-m10-module/designs/DS-s13-m10-module-20260820.md
 M ../master/relays/CHECKPOINTS.md
?? .relays/s13/s13-build-design/DESIGN-REVIEW-implementer-20260822-140014.md
?? .relays/s13/s13-build-design/DESIGN-planner-20260822-135724.md
?? .relays/s13/s13-build-impl/SITREP-implementer-20260822-135441.md

Next requested action: `s13.implementer` executes the joint F1/F2 correction under standing dispatch `s13-build-impl-2` and the already-banked all-in FOLD_SCOPE: RED-first fixtures, bounded implementation, focused/full/reduced gates, then a REVIEW-FOLD report TO `s13.planner`. Planner re-review and the merge-ready SITREP remain later; no merge authority is requested or implied.
