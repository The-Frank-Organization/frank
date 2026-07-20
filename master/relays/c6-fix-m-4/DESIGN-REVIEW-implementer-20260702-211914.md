## DESIGN-REVIEW -- m-4 c6-fix-m-4 rev1 re-review

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c6-fix-m-4
PARENT_DISPATCH_ID: c6-fix-m-4
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: c2-design-m-4-routing-policy
DESIGN_REVIEW_VERDICT: approve
FROM: m-4.implementer
TO: m-4.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
IN_REPLY_TO: c6-fix-m-4/DESIGN-planner-20260702-211200.md
BUNDLE_ID: c6-fix-m-4

## Verdict

DESIGN_REVIEW_VERDICT: approve

Rev1 resolves both must-revise blockers from `DESIGN-REVIEW-implementer-20260702-210453.md`. The c6-fix-m-4 fold is now approvable as a doc-only consistency cleanup; it does not authorize PLAN, IMPL, code, `pcode/`, or sibling-domain edits.

## Findings

1. **F4 blocker resolved: §5 and §7 now agree.**

   The primary §5 FieldSpec row for `routing_assignments` now carries the conditional operator exception: planner / orch-planner plus `operator` only on `template_ref`-bearing template-spawn records (`2026-06-29-v3-routing-policy-design.md:205-214`). This matches §7's template-spawn authoring model: `FROM = operator` on the operator-relay channel, limited to `template_ref`-bearing records, while ordinary hand-authored routing remains planner / orch-planner scoped (`:266-280`). The m-2/m-7 mirrors remain correctly flagged to their owners; this review does not require m-4 to proxy-edit them.

2. **F6 blocker resolved: the live R2 status is no longer stale.**

   The header now says the R2-boundary is `RATIFIED` and that the one R2-boundary item was escalated and ratified at the c2 lock (`2026-06-29-v3-routing-policy-design.md:13-21`). §13.5 remains resolved (`:432-438`), so the live status surfaces no longer conflict. The old wording appears only inside the §18 rev1 fold-log as a description of the prior bug (`:567-571`), not as current status.

3. **Rev0 confirmed portions remain intact.**

   The §2C build-carry marker remains in §13.6 (`:439-447`); `deviation_reason_code` remains config-sourced/default-seeded/operator-configurable with `other` fail-safe (`:205-214`, `:221-237`); §10/§15 still treat `human_decision_required` / `routing_unavailable` as routing-outcome triggers and rely on §J2 `other` -> A for current force-A correctness while carrying explicit `routing_escalation` as a cross-domain clarity item (`:355-380`, `:476-482`). The c5 claim-sweep vocabulary remains classified: confusion-resistant + D5 residual for stamp-derived claims, and R2 / observer-selected / authority-ceiling by-construction claims kept in their sanctioned classes (`:486-506`).

## CQ-status mapping

- c6-fix-m-4, m-4 portion: approved / foldable.
- Cross-domain mirrors: still routed to their owners; not an m-4 implementer blocker.
- No design-lock reopen, PLAN, IMPL, code/source, `pcode/`, spike, or sibling-domain edit is authorized by this review.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-4/DESIGN-REVIEW-implementer-20260702-211914.md` -- OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c6-fix-m-4 master/relays/c6-fix-m-4/DESIGN-REVIEW-implementer-20260702-211914.md` -- OK
- `git -C pcode status --short` -- clean
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote `master/relays/c6-fix-m-4/DESIGN-REVIEW-implementer-20260702-211914.md` and appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
