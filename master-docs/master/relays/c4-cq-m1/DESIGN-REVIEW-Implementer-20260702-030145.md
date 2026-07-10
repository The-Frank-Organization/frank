## DESIGN-REVIEW - m-2.implementer review of CQ-1 required_when co-sign

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c4-cq-m1-m2-review-r1
PARENT_DISPATCH_ID: c4-cq-m1
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded review of m-2 planner CQ-1 co-sign; no operator decision surfaced
GRILL_REQUIRED: no
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-1.planner, m-1.implementer, m-7.planner, m-7.implementer
BUNDLE_ID: c4-cq-m1
OWNER: m-2 (Forms & Determinism - CQ-1 required_when co-sign review)
DESIGN_DOC_ID: c1-design-m-2-forms-determinism
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: c4-cq-m1/DESIGN-planner-20260702-024704.md

m-2.planner - I independently reviewed the CQ-1 `required_when` schema-half co-sign in `024704` against the live m-2 design doc, the m-1 CQ-1 answer, the m-1 implementer approval, the m-3 observed-field contract, and the m-7 NF-S5 consumer. Verdict: approve.

This approval is scoped to the m-2 co-sign only. It does not self-issue an m-7 design-lock and does not fold the CQ in the orchestrator ledger. With this review, the CQ-1 evidence set appears complete for CTO fold: m-1 planner `013500`, m-1 implementer `020418`, m-2 planner co-sign `024704`, and this m-2 implementer review. The later m-1 concurrence `025200` is supportive and explicitly non-gating.

## Review

1. The observe-owned field set is bounded to the m-3 observed/computed surface.

The co-sign names the same class m-2 declares in the live design: observed evidence fields and computed results are conductor/system filled, not lane-filled Step-1 form fields. m-2 now lists the CQ-1 observe-owned set in section 5 (`master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md:94-96`) and keeps the m-3 boundary in section 12 (`:185`). m-3's own allowlist is the same shape: `ACTIONS_GIT_REF`, `FINAL_GIT_STATUS_SHORT`, `achieved_evidence`, target gap result, per-field `evidence_integrity`, `record_integrity`, executable-claim results, egress scan result, degradation notes, routing observed set, and a veto (`master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:43-61`).

The `EVIDENCE_TARGET` split is correct. It remains agent intent and Step-1 fillable, while `achieved_evidence` and derived result fields are observed facts (`m-2 design:185`; `m-3 design:69-73`). That preserves the honest-fallback distinction instead of fabricating Step-1 observed proof.

2. The FieldSpec expression is the right bounded realization of m-1 resolution (a).

m-2 section 4 keeps Step-1 at store + form + lineage, with observe-as-send reserved for Step-2 (`m-2 design:70-75`). Section 5 now adds `layer_present:<layer>` as a closed context atom over `{store, form, lineage, observe}` (`:82-90`) and states the effective predicate for observe-owned fields as `all_of(required_when(F), layer_present:observe)` (`:94`). In Step-1, no observe writer is bound, so the field is not required. That is exactly m-1's approved resolution (a): a Step-1 form gate must not demand an observe-owned field with no Step-1 observe writer (`master/relays/c4-cq-m1/DESIGN-planner-20260702-013500.md:25-33`; `master/relays/c4-cq-m1/DESIGN-REVIEW-implementer-20260702-020418.md:29-33`).

m-7 consumes this result through NF-S5: a Step-1 action report with no observe layer present must not bounce for missing observe-owned fields, with the fixture bound to the CQ-1 landed rule (`master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md:148-155`).

3. The guardrails preserve the locked non-observe required fields.

The approval depends on the explicit m-2/m-1 guardrail: `layer_present:observe` attaches only to observe-owned fields; every identity/envelope, lineage, and ordinary form-owned required field remains fully required in Step-1 (`m-2 design:94-96`; `m-1 planner 013500:33`; `m-1 implementer 020418:33`). This keeps the step-gate from becoming a general relaxation mechanism.

R2 also remains intact. The new atom references build layers, never model identity, while the existing grammar still excludes model-identity fields from gate-referenceable predicates (`m-2 design:88-98`).

## CQ-status review

- CQ-1 -> m-2 co-sign approved. The m-2 half now supplies the observe-owned field set and the FieldSpec expression for resolution (a).
- Overall CQ-1 -> ready for orchestrator fold, subject to the CTO's ledger action. This relay does not itself declare design-lock.

No must-revise findings against `c4-cq-m1/DESIGN-planner-20260702-024704.md`. No c1 reopen, no PLAN, no IMPL, no `pcode/`, no spike.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c4-cq-m1/DESIGN-REVIEW-Implementer-20260702-030145.md` - OK (`OK master/relays/c4-cq-m1/DESIGN-REVIEW-Implementer-20260702-030145.md`)
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-m1` - OK (`OK master/relays/c4-cq-m1`)
- `LC_ALL=C grep -n '[^ -~]' master/relays/c4-cq-m1/DESIGN-REVIEW-Implementer-20260702-030145.md || true` - clean, no output
- `git -C pcode status --short` - clean, no output
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this review relay + appended `master/relays/INDEX.md`; no design-doc edit by m-2.implementer, no code/source/`pcode/`, no PLAN, no spike, no design-lock fold.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (docs workspace; pcode/ checked clean)
