## DESIGN-REVIEW - s8 config host r4 must revise removed-schema references

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-design-m7-config-review-r4
PARENT_DISPATCH_ID: s8-design-m7-config
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - technical approval remains separate from the operator-owned F5 election
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m7-config
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s8-config-grill-r1
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m7-config/DESIGN-planner-20260711-034700.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner
SUBJECT: r4 must revise - the closed twelve-row schema is correct, but provenance, ordering, and runtime-home text still name removed root_homes/config_paths fields; one schema-consistency sweep remains

DESIGN_REVIEW_VERDICT: must-revise

R4 correctly pins twelve closed canonical-family rows, including `catalog-config`, its forbidden token, its planted-leak fixture, and the current r4 lock ID. One consistency finding remains.

## Finding

### F1 - Normative text still references schema fields removed by the r4 row-table fold

The closed schema now contains only `canonical_path_families.rows` (`2026-07-11-s8-config-host.md:46-63`). Later normative text still refers to the nonexistent `canonical_path_families.config_paths` in the provenance rule (`design:67`), to nonexistent `root_homes` and `config_paths` in ordering (`design:73-75`), and to the descriptor's `config_paths` census in runtime-home text (`design:92`). The twelve row objects themselves are not assigned an ordering rule.

Required fold: replace every removed-field reference with `canonical_path_families.rows` language. Define row ordering explicitly, recommended as the pinned normative sequence shown in the table or lexicographic by `id`; ensure the displayed table follows the selected rule. Phrase the post-design exception and live-oracle comparisons as the catalog row / three config-file rows inside `rows`. Sweep the whole artifact for `root_homes` and `config_paths`; zero hits should remain because neither is in the r4 schema.

## Confirmed

- `catalog-config` is now a stable named row with pinned relative path, forbidden token, and directory flag.
- FX-CFG-5 includes its named planted-token leg.
- The lock-impact ID correctly names `s8-design-m7-config-r4`.
- All previous descriptor, digest, SHA, ordering-sequence, step-4.5, and operator-gate findings remain resolved.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, no lock, and no operator F5 choice inferred.

ACTIONS_GIT_REF: wrote this r4 DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `691d034`
Next requested action: m-7.planner folds F1 into config-host r5 and returns a final DESIGN relay for re-review.
