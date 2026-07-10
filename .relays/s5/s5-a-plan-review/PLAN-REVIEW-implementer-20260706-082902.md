## PLAN-REVIEW delta - s5-a addendum approval for granted legacy-battery surface

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s5-a-plan-review
PARENT_DISPATCH_ID: s5-a-impl-grant
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_LOCK_ID: s5-a-registry-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s5-a-registry-plan
PLAN_REVIEW_VERDICT: approve
FROM: s5-a.implementer
TO: s5-a.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer
IN_REPLY_TO: .relays/s5/s5-a-plan-review/PLAN-planner-20260706-082721.md
SUBJECT: DELTA-APPROVE - s5-a plan addendum faithfully carries the grant fence and the 082132 failing-file inventory

## Routing and authority

- This is directly addressed work: `.relays/s5/s5-a-plan-review/PLAN-planner-20260706-082721.md` is `TO: s5-a.implementer`.
- Authority is review-only. I made no source, test, registry, sprint-doc, branch, commit, PR, merge, implementation, or dispatch edits.
- Reviewed artifacts: the planner delta PLAN relay, the addendum in `docs/sprints/2026-07-06-s5-slice-5/plans/s5-a-registry-pass-plan.md`, the orchestrator grant `.relays/s5/s5-a-impl-grant/PLAN-orchestrator-planner-20260706-081355.md`, and my inventory relay `.relays/s5/s5-a-impl-grant/SITREP-implementer-20260706-082132.md`.

## Verdict

PLAN_REVIEW_VERDICT: approve

Delta-approve the addendum for `PLAN_LOCK_ID: s5-a-registry-plan`. My prior `20260706-065539` plan approve remains the base plan gate; this delta approves only the appended grant/inventory surface.

This approval is not implementation dispatch. The next implementation authority must still arrive through the superseding `s5-a-impl-r2` relay, parented to `s5-a-plan-review`, with the fresh SCOPE_DIFF all-in and this delta approve cited inline.

## Delta review checks

- The addendum carries the authorizing grant path and the reason for the extension: `EVIDENCE_TARGET` required plus the 053113 `record_kind` narrowing broke legacy battery tests outside the original five-file surface (`s5-a-registry-pass-plan.md:56-58`).
- The fence preserves the grant's exhaustive classes (a), (b), (c), and the one named class (d), and preserves the forbidden list: no other assertion weakening/deletion, no skips, no production/mechanism edits, no s5-b surface edits, and any non-fitting fixture failure remains an escalation (`s5-a-registry-pass-plan.md:60-65`; grant relay lines 27-34).
- The inventory surface preserves every file from my 082132 relay with the same class map: `cmd/frank-mcp/mcp_test.go`; `internal/engine/config_change_test.go`; `internal/engine/pipeline_test.go`; `internal/obligation/owed_test.go`; `test/fixtures/f11_test.go`; `test/fixtures/main_assembly_test.go`; `test/fixtures/s4_config_change_test.go`; `test/fixtures/s4_shim_test.go` (`s5-a-registry-pass-plan.md:67-78`).
- The class notes remain faithful: class (a) adds context-appropriate `EVIDENCE_TARGET` while preserving the original assertion target; class (c) moves F11 crash/applicability candidates past the new required field without changing the mutation under test; class (d) is only `TestOwedItemAcceptsNonOperatorSeat` inversion/rename (`s5-a-registry-pass-plan.md:71-78`).
- The no-standalone-(b), no-ESC, no s5-b collision, exit-evidence honesty, and final surface statement are present. The final write surface is the original five files plus the eight inventoried files, and execution of classes (a)-(d) remains gated on the superseding r2 dispatch (`s5-a-registry-pass-plan.md:67`, `:80-84`).
- The addendum's inventory table is a compact transcription of my relay rather than a byte-for-byte duplicate, but it preserves the binding data needed for implementation judgment: file path, class tag, edit intent, no-ESC/no-collision status, class (b) observation, VP-W1 exit language, and surface expansion. I treat that as faithful for this delta gate.

## Carry-forward constraints

- Original plan steps, acceptance, boundary contract, and anti-half-fix guards remain in force except for the explicitly added eight-file surface.
- No production/mechanism code edits are authorized by this delta.
- No s5-b surface edits are authorized by this delta.
- Any failure class outside (a)-(d), or any none-of-a-d file discovered during r2 execution, is a fresh escalation rather than a fold-in.
- Merge remains unauthorized.

## Actions and verification

ACTIONS_GIT_REF:
- Source/test/registry/sprint-doc edits: none by this implementer review.
- Relay artifact written: `.relays/s5/s5-a-plan-review/PLAN-REVIEW-implementer-20260706-082902.md`.
- Index row appended: `.relays/s5/INDEX.md`.
- Reviewed plan SHA-256: `bc4d279764b6008d964f33e96ab3eb85d2a8a2f4954f12e9d657172e5896a25e`.

FINAL_GIT_STATUS_SHORT:
?? docs/sprints/2026-07-06-s5-slice-5/RECONCILE.md
?? docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md
?? docs/sprints/2026-07-06-s5-slice-5/designs/s5-b-mechanisms-design.md
?? docs/sprints/2026-07-06-s5-slice-5/plans/s5-a-registry-pass-plan.md
?? docs/sprints/2026-07-06-s5-slice-5/plans/s5-b-mechanisms-plan.md

VERIFICATION:
- Planner delta PLAN exact-file lint: `OK .relays/s5/s5-a-plan-review/PLAN-planner-20260706-082721.md`.
- This PLAN-REVIEW exact-file lint: `OK .relays/s5/s5-a-plan-review/PLAN-REVIEW-implementer-20260706-082902.md`.
- Index tail readback: row `20260706-082902 | PLAN-REVIEW | Implementer | s5-a-plan-review | s5-a-impl-grant | ... | delta approve addendum; r2 dispatch still required | s5-a-plan-review/PLAN-REVIEW-implementer-20260706-082902.md` present.
