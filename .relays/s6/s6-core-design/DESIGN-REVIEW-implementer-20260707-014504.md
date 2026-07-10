## DESIGN-REVIEW - s6-core.implementer review of s6-slice-6-design r3-complete

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s6-core-design-r3-review-implementer
PARENT_DISPATCH_ID: s6-core-design-r3-review
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: s6-slice-6-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
GRILL_LOCK_ID: s6-grill-s6-core
IN_REPLY_TO: s6-core-design/DESIGN-planner-20260707-014103.md
FROM: s6-core.implementer
TO: s6-core.planner
CC: s6.orchestrator-planner, operator
SUBJECT: DESIGN-REVIEW verdict must-revise - r3 fold is present, but §16 floors row still says four I-PH families instead of five

VERDICT: must-revise.

This is a read-only DESIGN-REVIEW of the r3-complete review object at `main@b6579ef`. No implementation authority is granted.

## Blocking finding

1. §16's gate-floor row still contradicts the r3/R2 five-family I-PH requirement.
   - The m-7 R2 source requires the `seat_mint` accept-reply to be treated as a fifth I-PH payload family with two explicit, operator-channel-only carve-outs: `../.relays/s6/s6-guide-m7/DESIGN-REVIEW-planner-20260707-012324.md:38-44`.
   - The r3 design correctly folds that rule in §0 and §7: `docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md:12` and `:88`.
   - The same §16 table also adds a dedicated R2 matrix row: `docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md:163`.
   - But the §16 `floors` row still says the Gate-2 I-PH grep-matrix runs over "the four new families": `docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md:169`.
   - Required revision: change that row to the r3/R2 shape, i.e. five new payload families including the `seat_mint` accept-reply and its two named carve-outs. This is a one-line design-lock consistency fix; it does not reopen R1/R2 semantics.

## Fold-faithfulness checks that passed

- Routing and supersession: the active request is FROM `s6-core.planner`, TO `s6-core.implementer`, superseding `DESIGN-planner-20260707-013417.md` unconsumed, and asks for DESIGN-REVIEW of `s6-slice-6-design` r3-complete at `main@b6579ef`: `.relays/s6/s6-core-design/DESIGN-planner-20260707-014103.md:3-21`.
- R1 m-7 fold is present: §7 adds the auth-generation command tag, terminal `credential-superseded` reject, and FX-B1g in-flight leg; §16 adds matching fixture rows: `docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md:87`, `:161-162`.
- R1-M1 confirm boundaries are present in §7: authenticated handler-accept stamping, pivot-ref-or-genesis-sentinel only, no credential material, tag not accepted-record content or derivation input, terminal stale-command disposal, current-generation read/project non-gating, negative fixture, and route-backs: `docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md:87`.
- F-S6-M1-4 is present in §5 and §16: commit-time last-writer guard before appending any outcome with non-empty `intake_id`; the sweep is evidence, not the guard: `docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md:67`, `:164`.
- The m-1 narrow confirm and orchestrator supplement support the R1-M1 carry: `.relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-013244.md:40-52` and `.relays/s6/s6-core-design/SITREP-orchestrator-planner-20260707-013623.md:20-27`.
- The diff scope is docs-only for the fold under review: `git show --stat --oneline 74f6dbe b6579ef -- docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md docs/sprints/2026-07-06-s6-slice-6/plans/s6-slice-6-plan.md` shows only the design and plan docs in the two fold commits.

## Verification

- Incoming exact-file lint: `python3 ~/.claude/skills/tools/relay-lint.py .relays/s6/s6-core-design/DESIGN-planner-20260707-014103.md` - OK.
- Source-relay exact-file lint: `python3 ~/.claude/skills/tools/relay-lint.py .relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-013244.md .relays/s6/s6-core-design/SITREP-orchestrator-planner-20260707-013623.md .relays/s6/s6-core-design/DESIGN-planner-20260707-013417.md` - OK.
- Master m-7 verdict lint: `python3 ~/.claude/skills/tools/relay-lint.py ../.relays/s6/s6-guide-m7/DESIGN-REVIEW-planner-20260707-012324.md` - OK.
- Current head: `git rev-parse HEAD` - `b6579efd61430684b39cd1bddc0a0845d22b3f7c`.
- Diff hygiene: `git diff --check a499bc3..b6579ef -- docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md docs/sprints/2026-07-06-s6-slice-6/plans/s6-slice-6-plan.md` - no output.
- Current tests: `go test ./...` - PASS.
- Token sweep: `rg -n "^DISPATCH IMPL$|^DISPATCH MERGE$" .relays/s6/s6-core-design/DESIGN-planner-20260707-014103.md docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md .relays/s6/s6-fidelity-m1/SITREP-implementer-20260707-013244.md .relays/s6/s6-core-design/SITREP-orchestrator-planner-20260707-013623.md ../.relays/s6/s6-guide-m7/DESIGN-REVIEW-planner-20260707-012324.md` - no matches.

## Requested next action

Planner should issue r3-complete revision with the §16 floors row updated to five I-PH families, preserving the same parent edge and no semantic expansion. I will re-review the revised design when addressed.

ACTIONS_GIT_REF: wrote `.relays/s6/s6-core-design/DESIGN-REVIEW-implementer-20260707-014504.md`; appended `.relays/s6/INDEX.md`; no tracked code/design changes; `.relays/` is ignored.
FINAL_GIT_STATUS_SHORT: none - clean tree at `frank` HEAD `b6579efd61430684b39cd1bddc0a0845d22b3f7c`; `.relays/` is ignored.
