## COORD — T9 held before edit: owner catalog bytes carry stale bounce census 16 while the live s8 tree requires 19

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s8-build-fidelity-m7-catalog-amendment
PARENT_DISPATCH_ID: s8-build-t7-close
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — m-7 owner-fidelity amendment required before T9 resumes; no merge authority
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m7-config
DESIGN_LOCK_ID: s8-design-m7-config-r13
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
TASK_ID: T9, T10
IN_REPLY_TO: master/relays/s8-build-t7-fold/SITREP-planner-20260711-235500.md
FROM: s8.implementer
TO: m-7.planner
CC: s8.planner, operator, master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-2.planner
SUBJECT: T9 stop before edit — the byte-exact owner artifact says bounce-format expected_sites 16, contradicting the gate-lift claim that the three pre-artifact s8 sites are baked in; live census and green invariant require 19; please amend owner bytes and return a new SHA

BLOCKER: the reviewed owner artifact cannot be applied byte-exact while keeping INV-CATALOG green. Per the binding census-delta rider, I made no source, fixture, catalog, or dogfood-store edit and route the discrepancy back to m-7.

Exact conflict:
- Owner artifact: `master/domains/m-7-conductor-core/design/2026-07-11-s8-v1-catalog-member.json`.
- Re-verified owner SHA-256: `14123af780ed903b3c4189f93ba166feabebf74f3a3681f2de66bff432ea57c2`.
- Owner row: `discovery.sink_patterns[bounce-format].expected_sites = 16`.
- Gate-lift relay says the T3/T6.5/T8 bounce-site moves predate the artifact and are baked into its verified counts. They are not present in that row.
- Live branch source census has 19 non-test `bounce.Format` call sites. `test/invariants/catalog.v1.json` carries 19 and `go test ./test/invariants -run '^TestLawPathHygiene$' -count=1` passes.
- Diff from `691d034..HEAD` identifies exactly three added production sites, all in `internal/engine/submit.go`: `lane-supplied-system-field`, completeness failure, and observe-predicate failure. That is the exact 16→19 delta already recorded during T3/T6.5/T8.

Why this blocks T9:
- The handoff requires byte-exact application with no reformatting, reordering, or in-slice census adjustment.
- Applying the supplied bytes would stage a catalog whose declared bounce census is three sites behind the live tree.
- Silently changing 16→19 would violate the m-7 single-writer/owner-fidelity rider. Proceeding with 16 would violate the plan's INV-CATALOG-green gate and the relay's own "baked in" claim.

Requested owner action:
1. Amend the owner artifact's bounce census from 16 to 19 after m-7 verification of the three named sites.
2. Return the complete reviewed byte artifact and its new SHA-256, confirming those exact bytes remain the SAME T9-genesis/T10-adoption member.
3. Confirm whether any artifact-shape consumer amendment is expected in T9: the current invariant source uses `path_hygiene`, while the s8-v1 owner member uses the fuller `discovery` descriptor. I will implement the locked s8-v1 shape once the owner bytes are internally current.

Verification at the s8 seat:
- `shasum -a 256` owner artifact → `14123af780ed903b3c4189f93ba166feabebf74f3a3681f2de66bff432ea57c2`.
- `rg -n 'bounce\.Format' --glob '*.go' --glob '!**/*_test.go' .` → 19 sites.
- `go test ./test/invariants -run '^TestLawPathHygiene$' -count=1` → PASS.
- Exact relay lint on `SITREP-planner-20260711-235500.md` → PASS.

ACTIONS_GIT_REF: none — report-only owner-fidelity stop; no frank edit or commit
FINAL_GIT_STATUS_SHORT: none — clean s8-observe-spine worktree at 3cce8cd
Next requested action: m-7.planner returns amended, reviewed byte-exact catalog content plus its SHA; s8.planner then re-lifts T9 on that corrected artifact, and s8.implementer resumes T9 under the standing token.
