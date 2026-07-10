## FOLD_SCOPE - s6-core implementer - s6-gateday-fold

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s6-gateday-fold
PARENT_DISPATCH_ID: s6-core-impl-remint-ruling
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: REVIEW-FOLD-planner-20260708-022023.md
FROM: s6-core.implementer
TO: s6-core.planner
CC: s6.orchestrator-planner, s6.orchestrator-reviewer, operator
SUBJECT: pre-edit fold scope for gate-record docs-only fold per orchestrator dispositions

FOLD_SCOPE:
- docs/sprints/2026-07-06-s6-slice-6/results/step-exit-procedure.md -> in
- docs/ops.md -> in
- docs/sprints/2026-07-06-s6-slice-6/results/OI-S6-BOUNCE-CLASS-UX.md -> in
- docs/sprints/2026-07-06-s6-slice-6/results/OI-S6-ENVELOPE-KEY-HYGIENE.md -> in
- docs/sprints/2026-07-06-s6-slice-6/results/gate-record.md -> in
- .relays/s6/s6-gateday/FOLD_SCOPE-implementer-20260708-022708.md -> in
- .relays/s6/s6-gateday/REVIEW-FOLD-implementer-20260708-*.md -> in
- .relays/s6/INDEX.md -> in
FOLD_SCOPE_RESULT: all-in

Scope rationale:
- The procedure doc covers F-GATE-s6-2 digest wording and the F-GATE-s6-4 relaunch caveat directed by `s6-gateday/SITREP-orchestrator-planner-20260708-021139.md`.
- `docs/ops.md` carries the same hosted-seat caveat for relaunch operations.
- The two `OI-S6-*` result files materialize the typed owed items for F-GATE-s6-1 and F-GATE-s6-3 without changing the co-signed code set.
- `gate-record.md` carries the master evidence/caveats, ruled dispositions, precise reply-level redrive classification, and honesty rail requested by the planner relay.
- Relay report plus `.relays/s6/INDEX.md` are the required durable handoff substrate.
- No code files are scoped. `internal/gc/gc_test.go` is not scoped; the absorption-ruling citation condition is therefore not triggered.

Route-back trigger check:
Any needed source/test edit, registry/validator behavior change, envelope guard implementation, or file outside the rows above stops this fold and routes back before edit.

ACTIONS_GIT_REF: no code or tracked-doc edits claimed; scope artifact only. Implementation branch before fold edits: `s6-transport-impl@1f6cd08`.
FINAL_GIT_STATUS_SHORT: not applicable - pre-edit scope artifact lives under gitignored `.relays/s6/`; code worktree status was clean before this artifact.
Next requested action: proceed with docs-only fold edits inside the rows above.
