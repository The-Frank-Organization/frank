## DESIGN-REVIEW — m-4 s5-escalations confirm legs (f)+(a): approve with C1/C2 carried as integration gates

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s5-escalations
PARENT_DISPATCH_ID: s5-escalations
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: c2-design-m-4-routing-policy
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: s5-escalations/DESIGN-planner-20260706-052000.md
FROM: m-4.implementer
TO: m-4.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, s5.orchestrator-planner
BUNDLE_ID: s5-escalations

## Verdict

APPROVE the two requested m-4 confirm legs:

1. **(f) degraded declaration shape:** acceptable for s5/Step-1 as top-level carriers plus prose-documented row columns, because the locked m-4 design explicitly records routing in v3.0 while deferring conductor execution to Step-3. The approval is conditioned on the already-registered Step-3 carry conditions C1 and C2 remaining explicit integration gates.
2. **(a) `named_enums` placement:** acceptable. The m-4 value-set remains config-sourced, default-seeded, operator-configurable, and protected by the hardcoded `other` fail-safe; placing the seven default tokens in the registry's `named_enums` mirrors the existing `gate_category` config pattern rather than hardcoding a code enum.

No m-4 design-doc change, registry write, `frank/` edit, code edit, or PLAN is authorized or needed by this review.

## Findings

No blockers.

**Routing/authority check.** The parent `RECONCILE-orchestrator-planner-20260706-050217.md` addressed m-4 only in `CC`, despite containing a directed m-4 section. I do not treat that parent CC as standalone authority. The actionable artifact for this review is the same-owner planner relay `DESIGN-planner-20260706-052000.md`, whose header directly addresses `TO: m-4.implementer`; the operator handoff supplied that exact relay. The later master reconcile `RECONCILE-orchestrator-planner-20260706-052214.md` also registers this review as a riding integration gate, not a PLAN-lock blocker.

**(f) adversarial check.** The degraded shape does not weaken R2 for Step-1 because the currently opaque row carrier cannot expose `chosen_model` to `required_when` or `visible_when` predicates. The locked m-4 design says v3.0 writes the routing record while Step-3 executes it, and says model values appear only inside `routing_assignments` payload, never in a predicate or authority/lineage gate. The live architecture now carries the exact C1/C2 Step-3 obligations: column-grain `gate_referenceable:false` for `chosen_model` and all model-identity columns, and the `any_row` deviation-justification coupling enforced the instant the router writes the rows. That is the required fence; do not collapse it to vague "column validation later."

**(a) adversarial check.** The current registry precedent stores `gate_category` values under `named_enums`, and `ARCHITECTURE.md §J2` defines that category as config-sourced/operator-configurable with `other` as the hard fail-safe. The m-4 design uses the same pattern for `deviation_reason_code` and names the seven defaults. `frank/internal/fieldspec/registry.json` does not yet contain `deviation_reason_code`; that is consistent with this being m-2's pending registry pass, not an m-4 implementation obligation.

## Evidence

- `master/relays/s5-escalations/DESIGN-planner-20260706-052000.md:13-17` — addressed design request, same `DESIGN_DOC_ID`, `TO: m-4.implementer`.
- `master/relays/s5-escalations/RECONCILE-orchestrator-planner-20260706-050217.md:45-48` — original m-4 confirm ask, with (f) and (a).
- `master/relays/s5-escalations/RECONCILE-orchestrator-planner-20260706-052214.md:49-53` — m-4 implementer approval rides as an integration gate, not a PLAN-lock blocker.
- `master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:51-56` and `:164-166` — record-now / execute-Step-3 boundary.
- `master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:208-218` — routing row shape, deviation fields, value-set, and model-payload-only invariant.
- `master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:438-448` — existing non-locking routing build-carries for per-column R2 and altitude-B per-row deviation grain.
- `master/ARCHITECTURE.md:499-501` — C1/C2 registered as explicit Step-3 routing-enforcement carries.
- `master/ARCHITECTURE.md:105-108` and `frank/internal/fieldspec/registry.json:53-66` — `gate_category` config/named-enum precedent.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s5-escalations/DESIGN-REVIEW-implementer-20260706-052840.md` — OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s5-escalations master/relays/s5-escalations/DESIGN-REVIEW-implementer-20260706-052840.md` — OK for dispatch root and target file.
- `tail -n 5 master/relays/INDEX.md` — row present at EOF: `20260706-052840 | DESIGN-REVIEW | implementer | s5-escalations | ... | approve | s5-escalations/DESIGN-REVIEW-implementer-20260706-052840.md`.
- `git status --short` at cwd — unavailable, not a git repo.
- `git -C frank rev-parse --short HEAD` — `67ee23e`.
- `git -C frank status --short` — observed `?? docs/sprints/2026-07-06-s5-slice-5/`; no `frank/` files were edited by this review.

ACTIONS_GIT_REF: wrote `master/relays/s5-escalations/DESIGN-REVIEW-implementer-20260706-052840.md` and appended `master/relays/INDEX.md`; no m-4 design-doc edit, no registry edit, no `frank/` code/doc edit by this review.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; `git -C frank status --short` currently reports `?? docs/sprints/2026-07-06-s5-slice-5/`.
Next requested action: m-4.planner can return the approved (f)+(a) confirm to master; s5 integration should keep C1/C2 as explicit Step-3 carry gates.
