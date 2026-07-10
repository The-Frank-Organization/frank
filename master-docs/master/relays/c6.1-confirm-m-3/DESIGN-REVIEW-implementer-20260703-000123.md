## DESIGN-REVIEW - m-3.implementer review of c6.1 confirm+flag

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c6.1-confirm-m-3
PARENT_DISPATCH_ID: c6.1-confirm-m-3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - bounded c6.1 confirm review; no new operator decision surfaced
GRILL_REQUIRED: no
DESIGN_DOC_ID: c2-design-m-3-observation-evidence
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: c2-grill-m-3
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: c6.1-confirm-m-3/DESIGN-planner-20260702-235604.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-4.planner, m-6.planner, operator
BUNDLE_ID: m-3-observation-evidence

DESIGN_REVIEW_VERDICT: approve

I adversarially reviewed `c6.1-confirm-m-3/DESIGN-planner-20260702-235604.md` against the CTO dispatch, the live m-3 design doc, the m-4/m-2 routing formula homes, the m-6 egress/bucket contract, m-7 NF-S9, `master/ARCHITECTURE.md`, and `master/c61-fix.diff`.

Approve the planner's confirm+flag artifact:
- Delta 1 is correctly confirmed.
- Delta 2 is correctly split: the new section 3.3 row is the right direction, but the section 3.3 note and the still-live section 3.2(c) token mapping remain contradictory and need CTO re-draft before c6.1 re-close.

This approval is NOT a confirmation of the current delta-2 m-3 text as closed. It approves the owner-pair flag that delta 2 is incomplete.

## Findings

1. Delta 1 - approve the planner confirmation.

Evidence:
- `master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:154` now defines `deviated_observed := declared_bucket != rank-1(recommended bucket for (role, task_tag))` and separately defines `bucket_binding_observed := chosen_model in members(declared_bucket)`.
- This matches the m-4 live text at `master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md:90` and `:93`.
- It also matches the m-2 field-home at `master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md:312`.
- The old c2 verbatim block is explicitly superseded, so there is no live second formula to reconcile.

2. Delta 2 - approve the planner flag.

Evidence:
- The corrected row at `master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:73` maps `egress_scan_result=blocked` to the non-terminal `egress_blocked` park + local resummon, matching m-6 and m-7.
- The same doc still says at `:63` that failed egress scan is condition (c), and that `(a)/(b)/(c)` are check-ran-said-no cases mapping to terminal `rejected`.
- The section 3.3 note at `:76` keeps the contradiction alive by distinguishing the outbound row from an "acceptance-time content-safety egress veto (c)" that maps to terminal `rejected`.
- m-6 excludes that split: `master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md:50` says egress at the conductor-governed chokepoint parks locally and raises an A resummon, never a D author-bounce; `:51` says egress is evaluated only at the external-send chokepoint on an already-A gate, while D bounces are acceptance-stage.
- m-7 NF-S9 agrees: `master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md:159` says a scanned blocked outbox item is blocked plus `egress_blocked` park plus local resummon.
- `master/ARCHITECTURE.md:309-310` also states egress is evaluated only at the external-send chokepoint on an already-A gate and resolves to `egress_blocked`, never auto-send; D bounces are separate.

Required re-draft before c6.1 close: align m-3 section 3.2(c), the CQ-4 token mapping sentence, and the section 3.3 note so egress is outbound-only: `egress_scan_result=blocked` => non-terminal `egress_blocked` park + local resummon, not terminal `rejected`. The acceptance-stage terminal `rejected` cases should remain predicate-false / declared-vs-observed integrity vetoes.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c6.1-confirm-m-3/DESIGN-planner-20260702-235604.md` - OK
- `nl -ba master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md | sed -n '46,84p;104,128p;134,154p;165,176p'` - inspected m-3 section 3.2, section 3.3, section 7, section 9, and consumer seams
- `nl -ba master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md | sed -n '84,98p'` - confirmed m-4 GL-1 formula
- `nl -ba master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md | sed -n '306,316p'` - confirmed m-2 field-home formula
- `nl -ba master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md | sed -n '40,56p;76,86p'` - confirmed m-6 egress/outbound-only disposition
- `sed -n '1,240p' master/c61-fix.diff` - checked CTO-applied c6.1 m-3 diff
- `rg -n "NF-S9|egress_blocked|egress|rejected|held|accepted" master/domains/m-7-conductor-core/design master/ARCHITECTURE.md | head -n 80` - confirmed m-7/ARCHITECTURE egress disposition anchors

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no design-doc edit by this seat, no `pcode/` edit, no PLAN, no IMPL.
FINAL_GIT_STATUS_SHORT: cwd root unavailable - `git status --short` returned `fatal: not a git repository (or any of the parent directories): .git`; `git -C pcode status --short` returned clean
RELAY_LINT: OK - `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c6.1-confirm-m-3/DESIGN-REVIEW-implementer-20260703-000123.md`; dispatch root OK - `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c6.1-confirm-m-3`
