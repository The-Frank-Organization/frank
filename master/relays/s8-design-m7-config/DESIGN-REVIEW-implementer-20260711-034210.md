## DESIGN-REVIEW - s8 config host r3 must revise two final lock-text defects

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-design-m7-config-review-r3
PARENT_DISPATCH_ID: s8-design-m7-config
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - technical approval remains separate from the operator-owned F5 election; activation authorization is already ratified
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m7-config
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s8-config-grill-r1
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m7-config/DESIGN-planner-20260711-033341.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner
SUBJECT: r3 must revise - the catalog path is now present but its canonical-family ID remains implicit, and the lock-impact line still names superseded r2; all behavioral/census folds otherwise approve

DESIGN_REVIEW_VERDICT: must-revise

R3 resolves both r2 findings mechanically: the post-design census now contains all three config paths and the order contract is field-specific. Two final lock-text defects remain.

## Findings

### F1 - The catalog canonical-path-family ID is still unspecified

The descriptor now pins `config/catalog/catalog.json` and says it contributes its own forbidden token (`2026-07-11-s8-config-host.md:46-61,78`). But the closed schema carries only path strings, while the current law's canonical families have stable IDs such as `engine-config` and `fieldspec-registry`. The new family ID is not named. The prior required fold explicitly required the corresponding ID/forbidden token; leaving the ID implicit makes the PLAN invent a governed census value and weakens the named-family planted-leak fixture.

Required fold: pin the new stable ID, recommended `catalog-config`, and state the path→ID/forbidden-token derivation for all three config paths (or represent canonical families as closed `{id, relative, forbidden, directory}` rows). Add the named `catalog-config` planted-token leg to the fixture obligation.

### F2 - The design-lock impact still names r2

The artifact status and incoming relay are r3, but the GRILL_LOCK says the reconciled lock is `s8-design-m7-config-r2` (`design:145`). A lock identifier cannot point at the superseded revision.

Required fold: change the effective lock ID to `s8-design-m7-config-r3` and sweep all lock-impact/reconcile text for stale revision identifiers.

## Confirmed

- Post-design `config_paths` is exactly catalog, engine, fieldspec; the catalog path is included in the canonical oracle and leak corpus.
- Set-like lists and ordered census sequences now have coherent per-field ordering semantics.
- All r1 and r2 digest, SHA, descriptor, step-4.5, claim-boundary, and operator-gate folds remain intact.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, no lock, and no operator F5 choice inferred.

ACTIONS_GIT_REF: wrote this r3 DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `691d034`
Next requested action: m-7.planner folds F1-F2 into config-host r4 and returns a final DESIGN relay for re-review.
