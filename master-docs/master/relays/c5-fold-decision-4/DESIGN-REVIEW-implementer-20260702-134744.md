## DESIGN-REVIEW -- m-6.implementer semantic review of decision 4 record-now fold

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c5-fold-decision-4
PARENT_DISPATCH_ID: c5-fold-decision-4
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- semantic review of non-locking build-carry record
GRILL_REQUIRED: no
FROM: m-6.implementer
TO: m-6.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-1.planner, m-1.implementer, m-7.planner
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
BUNDLE_ID: c5-fold-decision-4
OWNER: m-1 away-token plus m-6 resummon edge plus CTO integrated ledger
IN_REPLY_TO: c5-fold-decision-4/DESIGN-planner-20260702-133700.md

DESIGN_REVIEW_VERDICT: approve

I reviewed the m-6 record of operator decision 4 against the VP-required split, the live m-6 design doc, the operator decision register, and the parallel m-1 record/review state.

No must-revise findings. The decision 4 m-6 record is approved as correctly scoped record-now / mechanism-at-step-d. It is not an approval of the rotate/re-observe mechanism, fixture, away bridge, or park/wake implementation.

## Review

1. **The operator decision is recorded faithfully.**

   The operator register says refresh rotates `decision_id`, burns prior nonces, re-observes current state at `verify`, and bounces approval if state changed since the operator last saw it (`master/READINESS-REGISTER.md:351-354`). The m-6 section 12 carry records the same rotate/burn/re-observe rule and ties it to the resummon trigger edge (`master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md:177-182`).

2. **The record is explicitly non-locking.**

   The m-6 text says the mechanism, fixture, and full-pair adversarial review are owed at build-step d and that the carry is not design-locked here (`m-6 design:182`). That matches the VP-accepted c5 split: record decision 4 now as a build-carry, defer mechanism/proof before park/wake or away-bridge ships.

3. **The owner split is preserved.**

   m-6 owns the resummon trigger edge. It does not claim ownership of m-1's rotate/burn/re-observe mechanism. The parallel m-1 record states the TCB-owned rotate/burn/re-observe side and preserves decision-audit-continuity/exactly-one-wake as step-d design debt (`master/domains/m-1-trust-identity/design/2026-06-28-trust-identity-design.md:252-258`), with m-1.implementer approval now present in `c5-fold-decision-4/DESIGN-REVIEW-implementer-20260702-140100.md`.

4. **Step-1 residual behavior is not hidden.**

   The m-6 carry states the behavior is dormant until step d and that base sibling-burn plus m-6's never-auto-resolve-on-expiry FSM remain operative until then (`m-6 design:182`). That is the right non-locking record-now posture.

## Carry-forward

This approval covers only m-6's semantic record of decision 4. The rotate/re-observe mechanism, fixture, exactly-one-wake proof, and adversarial review remain owed at step d before the park/wake or away bridge ships. No PLAN, IMPL, `pcode/`, spike, or design-lock of decision 4 is authorized by this review.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c5-fold-decision-4/DESIGN-planner-20260702-133700.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-fold-decision-4` -- OK before this relay
- `nl -ba master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md | sed -n '169,182p'` -- reviewed m-6 section 12 carry
- `nl -ba master/domains/m-1-trust-identity/design/2026-06-28-trust-identity-design.md | sed -n '252,258p'` -- reviewed parallel m-1 record
- `sed -n '1,120p' master/relays/c5-fold-decision-4/DESIGN-REVIEW-implementer-20260702-140100.md` -- reviewed m-1 implementer approval for the m-1 half
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c5-fold-decision-4/DESIGN-REVIEW-implementer-20260702-134744.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-fold-decision-4` -- OK after this relay
- `git -C pcode status --short` -- clean
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote `master/relays/c5-fold-decision-4/DESIGN-REVIEW-implementer-20260702-134744.md` and appended `master/relays/INDEX.md`; no domain design-doc edit, no code/source/`pcode`, no PLAN, no spike, no mechanism design.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ checked clean)
Next requested action: master.orchestrator-planner may add decision 4 to the integrated build-carry ledger as non-locking record-now, preserving the step-d mechanism/proof gate.
