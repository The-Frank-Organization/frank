## DESIGN-REVIEW -- m-4 c5 claim-sweep-light semantic review

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c5-claim-sweep-light
PARENT_DISPATCH_ID: c5-claim-sweep-light
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: c2-design-m-4-routing-policy
DESIGN_REVIEW_VERDICT: approve
FROM: m-4.implementer
TO: m-4.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-5.planner, m-6.planner, m-7.planner
IN_REPLY_TO: c5-claim-sweep-light/DESIGN-planner-20260702-133000.md
BUNDLE_ID: c5-claim-sweep-light

## Verdict

DESIGN_REVIEW_VERDICT: approve

The m-4 claim-sweep fold is semantically approvable. The planner applied the VP-ratified checklist as claim-text hygiene only, narrowed the two stamp-derived malice-containment claims, and left the R2 gate-grammar / replay-completeness / authority-ceiling survivor claims in defensible classes.

## Findings

1. **Required relabels were folded.**

   The two `forgery-robust-stamped` claims identified by the sweep now resolve to `confusion-resistant` plus an explicit D5 residual in the body of the m-4 design: the routing artifact attribution sentence at `2026-06-29-v3-routing-policy-design.md:69` and the novelty statement at `:359-361`. The fold-log records both relabels and states that the stamp rides m-1's `submit()` rather than an m-4-owned mechanism (`:441-449`).

2. **The execution-fidelity claim is now scoped to future Step-3 only.**

   Section 12 names the v3.0 execution-fidelity gap as a documented D5-class residual: v3.0 enforces declaration honesty, but not proof that the launched lane actually ran the recorded model (`:375-390`). The `structurally` wording is fenced to future standalone-runtime Step-3, not claimed for v3.0 / Step-1. The §16 fold-log repeats that scope boundary (`:450-452`).

3. **The survivor list is correctly classified.**

   The kept `by construction` / `structural` claims are R2 gate-grammar claims: model identity remains non-gate-referenceable by the m-2 schema grammar (`:76-90`, `:453-455`). The replay-completeness `cannot be reconstructed` statement is a snapshot rationale, not a malicious-seat-containment claim, and the authority-ceiling statements are classified as confusion-resistant enforcement (`:455-457`).

4. **No stale overclaim remains outside fold-log context.**

   A full static search over the m-4 design doc and README found the old `forgery-robust` wording only inside classified fold-log / status references (`README.md:130`, design `:447`). No unclassified `unbypassable`, `tamper-resistant`, `sole-writer`, or equivalent containment claim remains in the m-4 surface checked here.

## CQ-status mapping

- c5 claim-sweep, m-4 portion: approved / foldable.
- Scope: claim-text hygiene only; no mechanism change, no locked-contract reopen.
- No design-lock, PLAN, IMPL, code, `pcode/`, or spike is authorized by this review.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c5-claim-sweep-light/DESIGN-REVIEW-implementer-20260702-134700.md` -- OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-claim-sweep-light master/relays/c5-claim-sweep-light/DESIGN-REVIEW-implementer-20260702-134700.md` -- OK
- `git -C pcode status --short` -- clean
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote `master/relays/c5-claim-sweep-light/DESIGN-REVIEW-implementer-20260702-134700.md` and appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
