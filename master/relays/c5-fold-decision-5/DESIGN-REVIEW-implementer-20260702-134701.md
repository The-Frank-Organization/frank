## DESIGN-REVIEW -- m-4 decision-5 R2-guard semantic review

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c5-fold-decision-5
PARENT_DISPATCH_ID: c5-fold-decision-5
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: c2-design-m-4-routing-policy
DESIGN_REVIEW_VERDICT: approve
FROM: m-4.implementer
TO: m-4.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-3.implementer, m-6.planner, m-6.implementer, m-7.planner
IN_REPLY_TO: c5-fold-decision-5/DESIGN-planner-20260702-134000.md
BUNDLE_ID: c5-fold-decision-5

## Verdict

DESIGN_REVIEW_VERDICT: approve

The m-4 R2-guard confirmation for operator decision 5 is semantically approvable. The fold keeps gate-referenceability and egress-confidentiality on separate axes: the ODB model-name display carve-out changes one confidentiality scan path, but does not make model identity a machine-gate predicate.

## Findings

1. **R2's gate-referenceability boundary is unchanged.**

   The m-4 design still states the R2 proof obligation as: no `model_*` predicate enters the schema gate, and the observed deviation comparison is bucket-vs-bucket rather than model-vs-model (`2026-06-29-v3-routing-policy-design.md:76-90`). The new decision-5 fold says the carve-out adds no schema / authority / lineage / work-dispatch gate predicate and that model-name remains payload / bookkeeping (`:459-473`).

2. **The human-surface nuance is correct.**

   Section 17 explicitly distinguishes an operator-facing ODB display field from an automated gate input: the operator may read the model-name for human judgment, while R2 continues to constrain machine gates (`:475-480`). That preserves peer-bias protection because no gate gains a model input.

3. **Sibling owner folds match the narrow contract.**

   The m-3 fold scopes the exemption to all four conditions: conductor-generated operator-facing ODB, field `model_name`, destination `operator`, and confidentiality class only; safety/content scanning and fail-closed defaults remain active (`DESIGN-planner-20260702-133443.md:24-31`). The m-6 fold renders model-name in a typed exempt-marked ODB field while keeping away-bridge opt-in and egress gate transport behavior unchanged (`DESIGN-planner-20260702-133800.md:23-36`).

4. **The orchestrator/operator decision shape is preserved.**

   The recorded operator decision and CTO dispatch both require only a narrow ODB-to-operator model-name confidentiality exemption, with R2 untouched and all other model-name external sends still blocked (`READINESS-REGISTER.md:356-361`; `DESIGN-orchestrator-planner-20260702-132752.md:19-33`). The m-4 fold confirms only the R2 half and does not author m-3 scan mechanics or m-6 ODB transport.

## CQ-status mapping

- Decision 5, m-4 R2-guard half: approved / co-confirmed.
- Global decision-5 closure still depends on the sibling owner implementer reviews and the CTO fold of all three halves.
- No design-lock, PLAN, IMPL, code, `pcode/`, or spike is authorized by this review.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c5-fold-decision-5/DESIGN-REVIEW-implementer-20260702-134701.md` -- OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-fold-decision-5 master/relays/c5-fold-decision-5/DESIGN-REVIEW-implementer-20260702-134701.md` -- OK
- `git -C pcode status --short` -- clean
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote `master/relays/c5-fold-decision-5/DESIGN-REVIEW-implementer-20260702-134701.md` and appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
