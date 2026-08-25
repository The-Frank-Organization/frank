## DESIGN-REVIEW CORRECTION — DS-s13-m10-module r2

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s13-build-design-review-3
PARENT_DISPATCH_ID: s13-build-design
RUN_ID: s13
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the filed E-0/E-1 escalation assigns both rulings to master; no fresh operator decision is requested by this review
FILED_AT_LOCAL: 20260820-224152
IN_REPLY_TO: s13-build-design/DESIGN-REVIEW-implementer-20260820-223944.md
DESIGN_DOC_ID: DS-s13-m10-module
DESIGN_DOC_SHA256: 84444951f56ff9ace316205e23245435d4fb1bc3e48368425e5fae7d1e50d06f
DESIGN_REVIEW_VERDICT: must-revise
FROM: s13.implementer
TO: s13.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer
SUBJECT: correction to r2 review classification — must-revise pending master E-0/E-1 rulings; F2/F4 remain closed

## Correction

The immediately preceding `223944` review correctly found that r2 cannot be approved while E-0/E-1 remain
undisposed, but it incorrectly mapped "master must rule" to `HUMAN_GATE_REQUIRED: yes` and
`DESIGN_REVIEW_VERDICT: human-decision-required`. The protocol reserves that header predicate for a fresh
operator decision. The planner's filed escalation explicitly assigns both rulings to
`master.orchestrator-planner` and says no fresh operator act is requested. The live courier therefore held the
first return as `relay-518be3c52a5b9a76702bd7e3` / `intake-000051`; it is preserved as history and superseded by
this correction.

Correct verdict: `must-revise`. Master must first dispose E-0 and E-1; then s13.planner must bind those rulings
into r3 under the same `DESIGN_DOC_ID` and request fresh review. No PLAN may claim an approving design-review
parent from r2, and no implementation authority follows.

## Substantive disposition carried unchanged

- **F1 remains open and correctly escalated as E-0.** Full
  `relay-lint --relay-root .relays/s13 --no-freshness` still exits 1 on the design/plan documents beneath
  `.relays/s13/docs/**`; exact r2 relay lint and INDEX lint pass. Master must rule the location/fence, then the
  ruled owner must make full root lint clean.
- **F2 is closed.** The r2 retirement transaction is total: full park + E+1 are common, G+1 exists only on the
  ordinary branch, and the cap terminal commits no successor/continuation/lease/snapshot/revival with the full
  FX-M10-CAP predicate.
- **F3 remains open and correctly escalated as E-1.** Authoritative module metadata confirms
  `modernc.org/sqlite@v1.36.1` declares Go 1.21 while v1.50.0–v1.57.0 declare Go 1.25. Frank remains Go 1.22+.
  Master must bind either the exact floor-preserving pin or the exact Go-1.25 baseline/README ownership change.
- **F4 is closed.** Complementary production/reduced tags, P2 in both selected files, and the three-command
  compile matrix close the selector and executable negative.

The full evidence and ruling-bound successor checklist remain in the preserved `223944` review. This successor
changes only the verdict/human-gate classification and live-delivery consequence.

ACTIONS_GIT_REF: docs-workspace disk action — this append-only correction relay + one live-EOF s13 INDEX row; the held predecessor is preserved; no source/test/branch/plan/design/store/token byte changed by s13.implementer
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s13/INDEX.md
 M frank/.relays/s13/docs/designs/DS-s13-m10-module-20260820.md
 M frank/.relays/s13/docs/plans/PL-s13-build-plan-20260820.md
 M frank/.relays/s14/INDEX.md
 M frank/.relays/s15/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s13/s13-build-design/DESIGN-REVIEW-implementer-20260820-221508.md
?? frank/.relays/s13/s13-build-design/DESIGN-REVIEW-implementer-20260820-223944.md
?? frank/.relays/s13/s13-build-design/DESIGN-REVIEW-implementer-20260820-224152.md
?? frank/.relays/s13/s13-build-design/DESIGN-planner-20260820-223212.md
?? frank/.relays/s13/s13-build-design/SITREP-planner-20260820-223211.md
?? frank/.relays/s14/s14-build/
?? frank/.relays/s15/s15-build-2/
?? frank/.relays/s15/s15-build/PLAN-REVIEW-implementer-20260820-220110.md
