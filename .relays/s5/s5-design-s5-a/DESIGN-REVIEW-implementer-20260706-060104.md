## DESIGN-REVIEW - s5-a implementer review of registry-pass design rev1

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s5-design-s5-a
PARENT_DISPATCH_ID: s5-design-s5-a
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: s5-a-registry-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: .relays/s5/s5-design-s5-a/DESIGN-planner-20260706-055846.md
FROM: s5-a.implementer
TO: s5-a.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer, operator
SUBJECT: DESIGN-REVIEW must-revise - rev1 folds main blockers, but on_timeout validation claim and row-count assertion need correction

## Verdict

DESIGN_REVIEW_VERDICT: must-revise

Rev1 fixed the three prior blockers in substance: the 053113 close is now first in the basis, stale riding-leg variants are gone, `on_timeout` is added to the row set and counts, and criterion 6 now allows the authorized `internal/fieldspec/*_test.go` registry-content fixtures. D-1/D-4/D-5/D-6 remain directionally acceptable.

Two narrow issues still block approval because they would put false assertions into the locked design and downstream tests.

## Findings

1. must-revise - the `on_timeout` row claims a validation behavior the live validator does not have.

Rev1 says the valueless `on_timeout` enum is loadable and that "a filled value then fails enumTokens conformance, validate.go:50-53." The loadability half is correct: the loader only checks `enum_set` when the field sets one, and `enum` plus `system_only` are valid row values. The validation half is false for the designed row shape. The row is `owner: system` and `fill_constraints: system_only`; `Validate` calls `ignorePayloadField` before required/type/enum checks, and `ignorePayloadField` skips every `owner == "system"` or `fill_constraints == "system_only"` row except `record_kind`. Therefore the current validate path never reaches `enumTokens` for a submitted `on_timeout` value.

Required fold: keep the valueless reserved row if that is still the intended shape, but remove the claim that current `validate.go:50-53` rejects a filled value. State the accurate boundary: Step-1 protection is render absence today, and submit-path rejection of lane-supplied system/computed headers depends on s5-b's DEF-2 typed-REJECT guard. If the design wants conductor-internal values checked later, name that as a later writer/validator responsibility rather than current FieldSpec validation.

Evidence:
- row claim: `docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md:80`.
- loader shape: `internal/fieldspec/registry.go:150-174`, `:223-256`.
- validator ordering and enum check: `internal/fieldspec/validate.go:31-53`.
- system/system_only skip: `internal/fieldspec/validate.go:115-120`.
- design already has the correct DEF-2 caveat elsewhere: `docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md:63` and `:141`.

2. must-revise - one registry-content test instruction still says row count 82 after the 83-row update.

The rev1 summary and acceptance criteria now say 83 rows and 38 enumerated absence names, but §7 still says `registry_test.go` should assert "row count 82." That is the exact registry-content test surface R-s5-2 assigns to s5-a, so the design would send the implementer to write a failing or stale assertion.

Required fold: change that §7 instruction to row count 83, and re-scan for stale 82/37/35 count text before re-requesting review.

Evidence:
- stale §7 instruction: `docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md:124`.
- corrected counts elsewhere: `docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md:16`, `:126`, `:139`, `:157-160`, `:175`.

## Checks That Now Pass

- Prior blocker 1 folded: `.relays/s5/s5-escalations/RECONCILE-orchestrator-planner-20260706-053113.md` is now first in the basis, `genesis` is removed from all scopes, owed rows are operator-only, `gate_resolution`/`disposition` are addressed, and §9 records zero open m-x riding legs.
- Prior blocker 2 mostly folded: `on_timeout` is present and the global counts moved to 83/38 except for the stale §7 test instruction above.
- Prior blocker 3 folded: acceptance criterion 6 now distinguishes production `.go` from authorized `internal/fieldspec/*_test.go` registry-content fixtures.

## Required Revision Checklist

- Correct the `on_timeout` validation rationale so it matches the current skip-before-enum-check validator path and the s5-b DEF-2 dependency.
- Change §7's `registry_test.go` row-count assertion from 82 to 83.
- Re-run a literal-text scan for stale `82`, `37`, and `35` count references before the next re-review request.

## Verification

- Reviewed the addressed rev1 DESIGN relay and current design doc.
- Re-read the prior must-revise relay and latest master riding-leg close.
- Re-checked live fieldspec loader and validator behavior for the `on_timeout` claim.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s5/s5-design-s5-a/DESIGN-REVIEW-implementer-20260706-060104.md` - OK.
- `go test ./internal/fieldspec` - OK.
- Pre-write `git status --short --untracked-files=all`:
  - `?? docs/sprints/2026-07-06-s5-slice-5/RECONCILE.md`
  - `?? docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md`
  - `?? docs/sprints/2026-07-06-s5-slice-5/designs/s5-b-mechanisms-design.md`

ACTIONS_GIT_REF: wrote `.relays/s5/s5-design-s5-a/DESIGN-REVIEW-implementer-20260706-060104.md` and appended `.relays/s5/INDEX.md`; no design-doc, registry, source, test, branch, commit, PR, merge, or store edits made by this review.
FINAL_GIT_STATUS_SHORT:
?? docs/sprints/2026-07-06-s5-slice-5/RECONCILE.md
?? docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md
?? docs/sprints/2026-07-06-s5-slice-5/designs/s5-b-mechanisms-design.md

Next requested action: s5-a.planner makes the two narrow text fixes above and re-requests DESIGN-REVIEW against `DESIGN_DOC_ID: s5-a-registry-design`.
