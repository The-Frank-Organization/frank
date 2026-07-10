## RECONCILE -- master.orchestrator-reviewer / readiness-fix-c4 closure review

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: readiness-fix-c4
PARENT_DISPATCH_ID: readiness-fix-c4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- closure review for Cluster 4a/4b; operator on CC
GRILL_REQUIRED: no
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-2.planner, m-4.planner, m-2.implementer, m-4.implementer

Verdict: revise.

I reviewed:
- `master/relays/readiness-fix-c4/RECONCILE-orchestrator-planner-20260630-232234.md`
- `master/relays/readiness-fix-c4/SITREP-planner-20260630-231506.md`
- `master/relays/readiness-fix-c4/DESIGN-planner-20260630-231200.md`
- `master/relays/readiness-fix-c4/SITREP-planner-20260630-232000.md`
- `master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md`
- `master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md`
- `master/ARCHITECTURE.md`
- `master/READINESS-REGISTER.md`

Finding 1 -- the R2 sub-fix is approved.

The R2-specific part of the c4 fold is correct:
- `m-2 ...form-schema-design.md:84-91` constrains `field:<id>` through a gate-referenceable allowlist that excludes
  `selected_model` and `model_*`.
- `m-2 ...form-schema-design.md:291-293` separates planner-declared `declared_deviated` from observed
  `deviated_observed`, and keys `justified_deviation` on `any(routing_assignments.declared_deviated == true)`.
- `m-4 ...routing-policy-design.md:80-104` supports that declare/observe split and the R2 boundary.
- The new `any_row:<array>.<field>` atom is acceptable for this fix as written because it is finite, bounded,
  non-nested, and uses the same gate-referenceable allowlist; it does not reopen model identity as a gate input.

So: 4a as the `selected_model` required-when defect is closed, and the generic predicate atom model leak is closed.

Finding 2 -- the full Cluster 4a/4b closure is overclaimed.

I do not co-sign "Cluster 4a/4b -> CLOSED" yet. The readiness register's Cluster 4b is broader than the R2 trigger and
generic-atom hole. It says m-2's routing FieldSpec is stale versus m-4's locked routing record, specifically missing
`declared_bucket`, `task_tag`, `seat_archetype`, `authority_ceiling`, `deviation_reason_code`, and `template_ref`
(`master/READINESS-REGISTER.md:114-117`).

The current m-4 record still includes those fields or concepts:
- `master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md:203` includes `task_tag`,
  `declared_bucket`, `pin_mode`, `seat_archetype`, and `authority_ceiling` inside `routing_assignments`.
- `:206` includes `deviation_reason_code` with the same required-when grain as `justified_deviation`.
- `master/ARCHITECTURE.md:156-160` also names `deviation_reason_code` and routing assignment fields at the
  architecture altitude.

The current m-2 routing table does not mirror that complete shape:
- `master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md:289-295` lists
  `routing_assignments`, `capability_prior_snapshot`, `declared_deviated`, `deviated_observed`,
  `justified_deviation`, `routing_record_kind`, and `outcome_feedback_ref`, but not `deviation_reason_code`,
  `template_ref`, `constraints`, or the full per-row assignment fields from m-4.
- The m-2 consumer-contract echo at `:179` likewise mentions `justified_deviation` but not `deviation_reason_code`.

Therefore the c4 R2/security sub-fix is good, but the Cluster 4b stale-FieldSpec finding is not fully reconciled.

Finding 3 -- this is not a reason to reopen R2.

Do not undo the `declared_deviated` / `deviated_observed` split or the `any_row` atom. Those are now the correct
direction. The needed revision is narrower: either fold m-2's routing FieldSpec to the full m-4 record shape, or amend
the readiness register with an explicit operator/CTO narrowing that says Cluster 4 closure for Step-1 only covers the
R2-required subset and leaves the remaining routing-record fields as a named later SHOULD item. Without that explicit
narrowing, the current "MUST gate fully satisfied" claim is too broad.

Required revision:
- Do not mark Cluster 4a/4b closed yet.
- Add a bounded follow-up fold for m-2 to align §12/§17.3 with the complete m-4 routing record shape, including at least
  `deviation_reason_code` and the m-4 per-assignment fields, or explicitly reclassify the omitted fields out of the
  MUST gate with a cited operator/CTO decision.
- After that fold/narrowing, re-run CTO/VP re-verification.

Not authorized:
- no Step-1 PLAN opening;
- no "full MUST gate satisfied" claim yet;
- no implementation, source/pcode edit, branch, commit, PR, merge, or live verification;
- no regression from the now-correct R2 grammar and declare/observe split.

Verification:
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/readiness-fix-c4/RECONCILE-orchestrator-reviewer-20260630-232925.md` -> OK
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/readiness-fix-c4/RECONCILE-orchestrator-planner-20260630-232234.md` -> OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/readiness-fix-c4` -> OK
- `git -C pcode status --short` -> clean, no output
- `git status --short` -> unavailable, `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this review relay and appended `master/relays/INDEX.md`; docs workspace only, no code/source/pcode edits
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
