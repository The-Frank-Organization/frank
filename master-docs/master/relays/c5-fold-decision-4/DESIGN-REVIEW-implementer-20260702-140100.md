## DESIGN-REVIEW -- m-1.implementer semantic review of decision 4 record-now fold

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c5-fold-decision-4-review
PARENT_DISPATCH_ID: c5-fold-decision-4
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- semantic review of a recorded operator decision as a non-locking build-carry
GRILL_REQUIRED: no -- no mechanism design or new product decision
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-6.planner, m-6.implementer, m-7.planner
BUNDLE_ID: c5-fold-decision-4
OWNER: m-1 (away-token) + m-6 (resummon edge) + CTO (integrated ledger)
DESIGN_DOC_ID: c1-design-m-1-trust-identity
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: c5-fold-decision-4/DESIGN-planner-20260702-135800.md

m-1.planner -- I reviewed the decision 4 record against the CTO dispatch, the VP-required split, the prior CQ-6 stable-`decision_id` recommendation, and the current m-1 design doc. Verdict: approve.

This approval is scoped only to the correctness of the record-now fold. It does not approve or lock the away-token mechanism, fixture, adversarial proof, park/wake bridge, or step-(d) design. No PLAN, no IMPL, no `pcode/`, and no spike are authorized.

## Review

1. The record is correctly non-locking and mechanism-at-(d).

The new section is explicitly titled "c5 build-carries (NON-LOCKING; mechanism + fixture + adversarial proof owed at step (d))" and says these are recorded constraints a builder inherits, not design-locked here (`m-1 design:252-255`). The residual line keeps the behavior dormant in Step-1 and names full-pair adversarial review as a (d) gate (`m-1 design:258`). That matches the CTO dispatch and VP split: record the decision now; defer detailed mechanism/proof to step (d).

2. The operator decision is recorded faithfully.

The design now records refresh as rotate `decision_id`, burn prior-cycle nonces, re-observe current state at `verify`, and bounce if the state changed since the operator last saw it (`m-1 design:256`). The owner split remains clean: m-1 owns rotate/burn/re-observe on the TCB mint/verify surface, and m-6 owns the resummon trigger edge (`m-1 design:256`).

3. The supersession note is honest and preserves the old concern as future design debt.

The record correctly says decision 4 supersedes/subsumes the earlier CQ-6 recommendation to keep `decision_id` stable because the operator chose rotate (`m-1 design:257`; prior recommendation in `c4-cq-m1/DESIGN-planner-20260702-021500.md`). It does not erase the lineage-continuity / exactly-one-wake objection. Instead it refiles that concern as a step-(d) obligation: preserve decision audit continuity across rotation and prove exactly-one-wake under rotation (`m-1 design:257`). That is the right disposition for a record-now carry.

No must-revise findings. The decision 4 m-1 record is approved as correctly scoped: recorded now as a non-locking build-carry, with the away-token mechanism and proof still owed at step (d).

## Verification

- `sed -n '1,260p' master/relays/c5-fold-decision-4/DESIGN-planner-20260702-135800.md` -- reviewed full m-1 planner record relay.
- `sed -n '1,260p' master/relays/c5-fold-decision-4/DESIGN-orchestrator-planner-20260702-132818.md` -- reviewed CTO record-now dispatch.
- `sed -n '1,260p' master/relays/c5-decomp/RECONCILE-orchestrator-reviewer-20260702-042018.md` -- reviewed VP split requiring record-now and mechanism-at-(d).
- `sed -n '1,240p' master/relays/c4-cq-m1/DESIGN-planner-20260702-021500.md` -- reviewed prior stable-`decision_id` recommendation.
- `nl -ba master/domains/m-1-trust-identity/design/2026-06-28-trust-identity-design.md | sed -n '252,270p'` -- reviewed folded §13.c5 text.
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c5-fold-decision-4/DESIGN-planner-20260702-135800.md` -- OK.
- `git -C pcode status --short` -- clean.
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this review relay and appended `master/relays/INDEX.md`; no design-doc edit, no code/source/`pcode` edit, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ checked clean)
RELAY_LINT: OK -- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c5-fold-decision-4/DESIGN-REVIEW-implementer-20260702-140100.md`
DISPATCH_ROOT_LINT: OK -- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-fold-decision-4`
