## DESIGN-REVIEW - s8 config host r2 must revise two remaining census contradictions

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-design-m7-config-review-r2
PARENT_DISPATCH_ID: s8-design-m7-config
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - technical approval remains separate from the operator-owned F5 election; activation authorization is now ratified on the master record
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m7-config
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s8-config-grill-r1
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m7-config/DESIGN-planner-20260711-032228.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner
SUBJECT: r2 must revise - the new catalog runtime path is absent from its own canonical census, and the global sort rule contradicts the pinned family/pattern orders; the four r1 findings otherwise fold cleanly

DESIGN_REVIEW_VERDICT: must-revise

The r2 fold resolves r1 F1-F4 in substance: the descriptor now has a closed buildable shape and consumption map; the drift oracle uses direct bytes plus `ExpectedConfigDigest()`; the fieldspec SHA is full; technical approval no longer ratifies operator choices. Two new normative contradictions remain.

## Findings

### F1 - The descriptor omits the catalog path that this design adds

The design makes `<root>/config/catalog/catalog.json` a canonical runtime member and says it joins the canonical path census (`2026-07-11-s8-config-host.md:74-78`). But the closed descriptor pins only `config/engine.json` and `config/fieldspec/registry.json` (`design:46-49,57`). Once the PLAN adds `catalog` to `StoreRootConfigPaths`, the live oracle will return three paths while the descriptor expects two, so the design's own canonical-equality leg turns red on every conforming implementation.

Required fold: add `config/catalog/catalog.json` to the normative `config_paths`, add its corresponding canonical-path-family ID/forbidden token to the law contract, update the stated count from two to three, and carry it into FX-CFG-3/5 or the canonical-census fixture as appropriate. Pin the **post-design expected census**, not only the pre-implementation `691d034` census.

### F2 - The global sorting rule conflicts with the pinned family and sink orders

Section 3 says "every list sorted lexicographically" (`design:59`). The normative `output_families` order (`bounce`, `process`, `tool`, `rendered`, `delivery`, `seat-mint`) and `sink_patterns` order (`bounce`, `safe`, `tool`, `projection`, `delivery`, `seat-mint`) at `design:55-56` preserve current catalog/census order and are not lexicographic. The live law also explicitly expects the family census order. A PLAN cannot satisfy both requirements.

Required fold: define ordering per field. Keep `site_census`, boundary paths, roots, and other set-like lists lexicographically sorted; explicitly mark `output_families` and `sink_patterns` as ordered normative sequences if current order is semantically pinned, or sort their pinned values and change the consumer contract accordingly. Do not retain a universal rule with exceptions only inferable from examples.

## Confirmed Fold Evidence

- R1 F1: the 17-site multiset, six families/carve-outs, six sink patterns/counts, recognizers, roots/exclusions, normalization, and consumption map are now explicit.
- R1 F2: source/runtime byte equality plus recomputed current composite against `ExpectedConfigDigest()` is a real post-config-change oracle.
- R1 F3: the full SHA is `1ef6abab4d496b11017f57ca400e8296d63824994ffce8311e4533f70cc92485`.
- R1 F4: the lock is correctly conditional on the operator F5 record. The separate activation-authorization leg is now satisfied by `step2-plan/SITREP-orchestrator-planner-20260711-025643.md`; master reconciliation can consume that record.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, no lock, and no operator F5 choice inferred.

ACTIONS_GIT_REF: wrote this r2 DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `691d034`
Next requested action: m-7.planner folds F1-F2 into config-host r3 and returns a new DESIGN relay for re-review; master holds config reconciliation/lock consumption meanwhile.
