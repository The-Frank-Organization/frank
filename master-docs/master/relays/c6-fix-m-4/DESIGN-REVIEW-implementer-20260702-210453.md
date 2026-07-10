## DESIGN-REVIEW -- m-4 c6-fix-m-4 adversarial review

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
DESIGN_REVIEW_VERDICT: must-revise
FROM: m-4.implementer
TO: m-4.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
IN_REPLY_TO: c6-fix-m-4/DESIGN-planner-20260702-210000.md
BUNDLE_ID: c6-fix-m-4

## Verdict

DESIGN_REVIEW_VERDICT: must-revise

The c6 fold is directionally right, but two live m-4-owned surfaces still contradict the return claim. The fixes are narrow doc edits: update the primary m-4 FieldSpec row for the template operator exception, and remove or past-tense the stale R2 escalation status in the design header. I am not blocking on the cross-domain mirrors themselves; those are correctly flagged to their owners.

## Blocking findings

1. **F4 is not fully folded into the primary FieldSpec row.**

   The planner relay claims "`routing_assignments` seat-scope widened to admit `operator` on `template_ref`-bearing records only" (`DESIGN-planner-20260702-210000.md:29`, `:38-40`). The actual §7 fold records that exception correctly: template-spawn routing is `FROM = operator`, ordinary routing remains planner / orch-planner, and the exception is limited to `template_ref`-bearing records (`2026-06-29-routing-policy-design.md:265-279`).

   But §5, the design's primary FieldSpec table, still says `routing_assignments | seat_scoped_enum -> planner / orch-planner` with no conditional operator exception (`:204-214`). Since §5 is explicitly the `routing_decision` FieldSpec and the m-2 mirror still reads planner/orch-planner only (`2026-06-28-form-schema-design.md:304-315`), the m-4 doc currently has two competing seat-scope sources. Revise §5 to include the same conditional: `planner / orch-planner; operator only on template_ref-bearing template-spawn records` or equivalent. Keep the m-2/m-7 mirrors flagged to their owners.

2. **F6 stale R2 status remains in the live header.**

   The fold correctly marks the R2-boundary as resolved in §13.5: "RESOLVED -- ratified at the c2 lock" (`2026-06-29-routing-policy-design.md:432-438`). But the design header still says "one R2-boundary item is escalated to orchestrator/VP for ratification at the c2 lock" (`:14-22`). That is live status text, not a historical quote, and contradicts the c6 F6 cleanup claim. Revise it to past tense, for example "was escalated and ratified at the c2 lock", or remove the sentence.

## Confirmed portions

- F1 / x1-F1: the §2C routing-lane build-carry deferral marker is present in §13.6, including `gate_referenceable` per-column negative fixtures and altitude-B per-row deviation grain (`2026-06-29-routing-policy-design.md:439-447`), matching the restored c5 ledger in `RECONCILE.md:325-331`.
- F5: `deviation_reason_code` now has a config-sourced, default-seeded, operator-configurable value-set with `other` fail-safe in §5/§6 (`2026-06-29-routing-policy-design.md:204-210`, `:221-237`).
- F7 / x3-F1: the false `human_decision_required`/`routing_unavailable` A-member assertion is gone; §10 now treats them as routing-outcome triggers and relies on §J2 `other` -> A for present correctness while carrying explicit `routing_escalation` as a cross-domain clarity item (`2026-06-29-routing-policy-design.md:354-378`; `ARCHITECTURE.md:105-114`).
- F8/F9: the late-classified snapshot-provenance KEEP is recorded in §16 (`2026-06-29-routing-policy-design.md:500-506`), and ARCHITECTURE §J1 already includes the decision-5 typed ODB model-name carve-out (`ARCHITECTURE.md:99-103`).

## Required revision

- Update m-4 §5 so the FieldSpec row and §7 agree on the template-spawn operator authoring exception.
- Update the header R2 status sentence so it no longer describes a resolved c2 item as currently escalated.
- Return a rev1 `DESIGN` relay for re-review. No PLAN, IMPL, code, `pcode/`, or sibling-domain edit is authorized by this review.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-4/DESIGN-REVIEW-implementer-20260702-210453.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c6-fix-m-4 master/relays/c6-fix-m-4/DESIGN-REVIEW-implementer-20260702-210453.md` -- OK
- `git -C pcode status --short` -- clean
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote `master/relays/c6-fix-m-4/DESIGN-REVIEW-implementer-20260702-210453.md` and appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
