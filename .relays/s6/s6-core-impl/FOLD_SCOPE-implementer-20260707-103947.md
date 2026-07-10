## FOLD_SCOPE - s6-core implementer - s6-core-impl-fold-remint

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s6-core-impl-fold-remint
PARENT_DISPATCH_ID: s6-core-impl-remint-ruling
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: REVIEW-FOLD-planner-20260707-103642.md
FROM: s6-core.implementer
TO: s6-core.planner
CC: s6.orchestrator-planner, s6.orchestrator-reviewer, operator
SUBJECT: pre-edit fold scope for released re-mint crash-window fix per m-1 option A redlines and orchestrator directive

FOLD_SCOPE:
- internal/seat/binding.go -> in
- internal/seat/binding_test.go -> in
- cmd/frank/main.go -> in
- test/fixtures/s6_mint_test.go -> in
- test/fixtures/s6_iph_test.go -> in
- docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md -> in
- .relays/s6/s6-core-impl/FOLD_SCOPE-implementer-20260707-103947.md -> in
- .relays/s6/s6-core-impl/REVIEW-FOLD-implementer-20260707-*.md -> in
- .relays/s6/INDEX.md -> in
FOLD_SCOPE_RESULT: all-in

Scope rationale:
- `internal/seat/binding.go` and `internal/seat/binding_test.go` cover redlines 1-2: optional `realized_mint_ref` on the binding row, same atomic binding-table write with the credential, and `MintOrReplace` taking the accepted pivot relay id.
- `cmd/frank/main.go` covers redlines 3-5 and 7: live path passes the pivot ref, recovery computes latest accepted `seat_mint` pivots in commit order before channel open, canonical accepted pivots win over stale binding provenance, and activation/roster/R1 submit rules remain separate.
- `test/fixtures/s6_mint_test.go` covers redline 8: SIGKILL between accepted re-mint pivot and binding replacement; restart recovery before serve; superseded credential fails the first post-restart authentication; row realizes the latest pivot.
- `test/fixtures/s6_iph_test.go` covers redline 6 custody sweeps for `realized_mint_ref` across non-admin surfaces.
- `docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md` is the one-line §7 addendum directed by `SITREP-orchestrator-planner-20260707-103252.md`.
- Relay report plus `.relays/s6/INDEX.md` are the required durable handoff substrate.
- No `internal/gc/gc_test.go` row is included; the absorption-ruling citation condition is therefore not triggered.

Route-back trigger check:
This scope excludes credential material/hash/session/socket history in the binding row, counters/timestamps as generation authority, split durable writes, post-open repair, non-admin exposure, activation/roster/R1-rule changes, and option-B residual documentation. If implementation pressure creates any such need, the fold stops before edit and routes back.

ACTIONS_GIT_REF: no code edits claimed; scope artifact only. Implementation branch before fold edits: `s6-transport-impl@a8d04b4`.
FINAL_GIT_STATUS_SHORT: not applicable - pre-edit scope artifact lives under gitignored `.relays/s6/`; code worktree status was clean before this artifact.
Next requested action: proceed with red-first fold edits inside the rows above.
