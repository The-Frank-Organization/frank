## DESIGN-REVIEW - m-3.implementer review of c6-fix-m-3 cleanup

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c6-fix-m-3
PARENT_DISPATCH_ID: c6-fix-m-3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded c6 doc-only review; no new operator decision surfaced
GRILL_REQUIRED: no
DESIGN_DOC_ID: c2-design-m-3-observation-evidence
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: c2-grill-m-3
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: c6-fix-m-3/DESIGN-planner-20260702-210322.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-4.planner, operator
BUNDLE_ID: m-3-observation-evidence

DESIGN_REVIEW_VERDICT: must-revise

I reviewed `c6-fix-m-3/DESIGN-planner-20260702-210322.md` against the orchestrator dispatch, the live m-3 design doc, the c5 superseding review relays, and the c6 re-review targets. Most of the c6 fixes are directionally right and preserve the lock shape. I cannot approve yet because two local consistency defects remain builder-visible in the m-3 design doc.

## Blocking revisions

1. The STATUS header cites superseded c5 `must-revise` relays as the implementer-approved c5 folds.

Evidence:
- `master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:7` says the c5 folds are implementer-approved and cites `(134834 / 134833)`.
- Those IDs are the prior `must-revise` reviews: `master/relays/c5-fold-decision-5/DESIGN-REVIEW-implementer-20260702-134834.md:16` and `master/relays/c5-claim-sweep-light/DESIGN-REVIEW-implementer-20260702-134833.md:16`.
- The superseding approval relays are `master/relays/c5-fold-decision-5/DESIGN-REVIEW-implementer-20260702-140748.md:16` and `master/relays/c5-claim-sweep-light/DESIGN-REVIEW-implementer-20260702-140749.md:16`.
- The c6 fold-log also says m-3-F9 header was verified at `master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:241`, so this is not merely a missing nicety; it is the exact item marked verified by the incoming relay.

Required revision: update the header to cite `140748 / 140749`, or remove the c5 review IDs and cite the fold names without stale relay suffixes. Update the c6 fold-log's m-3-F9 claim so it no longer records the stale header as verified before correction.

2. The resolved decision for executable-claim execution still says "arbitrary lane code rejected."

Evidence:
- The c6 fix correctly narrows section 4: arbitrary agent-supplied code is rejected, while suite-class entries execute repo-resident lane code under the non-locking unprivileged-executor carry (`master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:88`).
- The build-step carry repeats that suite-class entries execute repo-resident lane code and must run without store/config/outbox handles or signing key (`master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:221`).
- The resolved decision remains broader: "conductor-owned check-registry; arbitrary lane code rejected" (`master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:189`).

Required revision: align the resolved decision with section 4 and section 13, for example: "conductor-owned check-registry; arbitrary agent-supplied code rejected; suite-class repo-resident lane code may run only through the non-locking unprivileged-executor carry." Keep it explicitly DOC-ONLY and do not reopen the lock.

## Confirmed non-blockers

- The section 3.3 disposition table is total enough for the c6 target and preserves the byte-exact terminal enum `{accepted, rejected, held}`.
- The section 4 `attestation_source {conductor | operator}` marker does not reopen R3; it stays separate from the two-value `evidence_integrity` enum.
- The section 5 `scope_paths` observable condition is correctly conditioned on an m-2-declared field.
- The section 11 detectability text is now scoped to inconsistent forgery; coherent D5 direct-store forgery remains a residual, not a guarantee.
- The section 13 entries cover the decision-5 fixture receiving surface, the `mixed` authority-class fixture leg, the suite-class executor carry, `attestation_source`, and `scope_paths`.
- The m-3 side of GL-1 is R2-clean: section 9 canonically uses bucket-vs-bucket. The m-4 fallback-bracket edit remains a cross-owner item, not an m-3 approval blocker by itself.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-3/DESIGN-planner-20260702-210322.md` - OK
- `nl -ba master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md | sed -n '1,18p;80,95p;182,194p;214,244p'` - inspected c6 landing points and blockers
- `nl -ba master/relays/c5-fold-decision-5/DESIGN-REVIEW-implementer-20260702-134834.md | sed -n '1,36p'` - confirmed old decision-5 relay is `must-revise`
- `nl -ba master/relays/c5-claim-sweep-light/DESIGN-REVIEW-implementer-20260702-134833.md | sed -n '1,36p'` - confirmed old claim-sweep relay is `must-revise`
- `nl -ba master/relays/c5-fold-decision-5/DESIGN-REVIEW-implementer-20260702-140748.md | sed -n '1,32p'` - confirmed superseding decision-5 approval
- `nl -ba master/relays/c5-claim-sweep-light/DESIGN-REVIEW-implementer-20260702-140749.md | sed -n '1,32p'` - confirmed superseding claim-sweep approval

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no design-doc edit by this seat, no `pcode/` edit, no PLAN, no IMPL.
FINAL_GIT_STATUS_SHORT: cwd root unavailable - `git status --short` returned `fatal: not a git repository (or any of the parent directories): .git`; `git -C pcode status --short` returned clean
RELAY_LINT: OK - `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-3/DESIGN-REVIEW-implementer-20260702-210904.md`; dispatch root OK - `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c6-fix-m-3`
