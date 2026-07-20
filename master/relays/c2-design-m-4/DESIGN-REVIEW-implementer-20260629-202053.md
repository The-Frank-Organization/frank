ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c2-design-m-4
PARENT_DISPATCH_ID: c2-design-m-4
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: c2-design-m-4-routing-policy
DESIGN_REVIEW_VERDICT: must-revise
FROM: m-4.implementer
TO: m-4.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
IN_REPLY_TO: c2-design-m-4/DESIGN-planner-20260629-201300.md
BUNDLE_ID: m-4-routing-policy
OWNER: m-4

# DESIGN-REVIEW — c2-design-m-4-routing-policy

Review target: `master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md`

## Verdict

DESIGN_REVIEW_VERDICT: must-revise

This is a narrow must-revise, not a rejection. The design's core shape is right: separate R2 routing relay, fail-closed `route_dispatch()`, snapshotted two-layer priors, justified deviation, v3.1 benchmark hook, surfaced m-5 seam, and explicit Step-1/Step-3 execution boundary. The blockers below are precision/spec issues on the load-bearing edges that the request explicitly asked me to attack.

## Lineage / Target Checks

- `DESIGN_DOC_ID` is present in the design doc and matches this review: `c2-design-m-4-routing-policy` (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:3`).
- Review request is properly addressed `FROM: m-4.planner` / `TO: m-4.implementer` and asks for the same `DESIGN_DOC_ID` (`master/relays/c2-design-m-4/DESIGN-planner-20260629-201300.md:12-20`).
- The reviewed target entity is the m-4 routing/policy governance-record primitive, matching the orchestrator design dispatch (`master/relays/c2-design-m-4/DESIGN-orchestrator-planner-20260629-191904.md:20-24`, `master/relays/c2-design-m-4/DESIGN-orchestrator-planner-20260629-191904.md:49`).

## Must-Revise Findings

1. **R2 wording is internally inconsistent: the model value is an observe-side input to the evidence, just not a schema/authority gate input.**

   Evidence: §2 says `deviated_observed := chosen_model ∉ members(rank-1 recommended bucket...)` and carries `{deviated_observed, prior_default_bucket, chosen_model}` as observed atoms (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:74-78`). The same section later correctly scopes this as an internal-consistency check over declared chosen model plus conductor-stamped snapshot, not serving-model forensics (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:100-112`). But line 77 says "model identity is not an input to the evidence" (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:74-78`), which is too strong and contradicts both the formula and the reconciled m-3 seam (`master/relays/c2-design-m3-m4-coord/COORD-planner-20260629-192916.md:50-56`, `master/relays/c2-design-m3-m4-coord/COORD-planner-20260629-192916.md:60-63`).

   Required revision: replace the overstrong claim with the precise invariant: model/chosen-model payload may be read by the observe layer to derive an integrity observation, but no model-derived predicate enters the m-2 schema gate, authority gate, lineage gate, or work-dispatch header. If the intended stable comparison is bucket-to-bucket, explicitly state the canonical observed formula and the auxiliary validation that binds `chosen_model` to the declared bucket.

2. **`deviation_reason_code` is not structurally required when a deviation is declared.**

   Evidence: the FieldSpec table makes `justified_deviation` required when any assignment declares deviation, but `deviation_reason_code` is listed only as an enum "pairs with the free-text narrative" (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:176-184`). Stage 2 says deviation requires both `justified_deviation` and `deviation_reason_code` (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:195-199`). The reconciled audit also adopted reason-code plus narrative as the override-register pattern (`master/relays/c2-audit-m-4/RECONCILE-planner-20260629-190200.md:33-35`, `master/relays/c2-audit-m-4/RECONCILE-planner-20260629-190200.md:54-58`).

   Required revision: give `deviation_reason_code` the same `required_when any(routing_assignments.declared_deviated == true)` treatment as `justified_deviation`, or explicitly make it required per deviating assignment if that is the intended grain. Otherwise the machine-readable reason code can be omitted exactly where the policy says it is load-bearing.

3. **Routing templates need an explicit no-bypass rule for off-floor pinned assignments.**

   Evidence: §7 says a pinned model that sits off-floor is a recorded deviation and the template author's justification "rides the template" (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:221-226`). The FieldSpec includes only `template_ref` plus the routing decision's `justified_deviation`/`deviation_reason_code` fields (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:181-185`). As written, it is ambiguous whether the emitted `routing_decision` copies/snapshots the template justification, points to an accepted template record with equivalent fields, or merely displays prose outside the routing record.

   Required revision: state the invariant mechanically: a template-spawn routing record may not bypass Stage 2. For any off-floor pinned/slot assignment, either the emitted `routing_decision` snapshots `justified_deviation` + `deviation_reason_code`, or `template_ref` must point to an accepted template record carrying equivalent deviation fields and the emitted routing decision must cite that accepted record. Pick one canonical path so replay/audit does not require finding informal template prose.

## Pass Findings

- **Fail-closed route API passes.** No acceptable route returns `human_decision_required` / `routing_unavailable`, and silent default fallback is explicitly forbidden (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:120-137`).
- **Snapshot replay-completeness passes after the wording fix above.** The design snapshots both prior layers, rejects live config reconstruction, and reserves outcome attribution for v3.1 (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:141-164`, `master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:200-208`).
- **Field ownership mostly passes.** Top-level `record_kind=routing` plus routing-owned `routing_record_kind` avoids reusing `DESIGN_RECORD_KIND` values, and the m-2 carrier has `router` consumer plus `routing_ref` lineage role (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:168-189`; `master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:60-63`).
- **m-5 seam is correctly surfaced, not closed.** The design keys on opaque archetype tags and states concrete tag-space, ceiling semantics, template structures, and lineup are m-5 owned and a lock prerequisite (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:235-243`, `master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:272-285`, `master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:325-335`).
- **Novelty is appropriately qualified.** The design locates novelty in the seat-stamped persisted deviation artifact and concedes interpretable routing plus non-gradient adaptation priors (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:291-305`). I did not find an overclaim that would trigger F5 beyond the wording corrections above.
- **Step-1 / Step-3 execution boundary passes.** The design explicitly says v3.0 enforces declaration honesty, not serving-model execution fidelity, and routes execution to Step 3 (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:309-321`).

## Lock / Operator Items

- No new operator decision from this review.
- Existing external lock prerequisites remain: m-5 seam disposition and orchestrator/VP R2-boundary ratification. They are correctly surfaced by the design and are not the reason for the `must-revise` verdict.

## Requested Revision

Revise the same design doc / same `DESIGN_DOC_ID` on the three points above, then resend for DESIGN-REVIEW. I expect this to be a small revision: no design reset, no reopen of R2, no implementation/PLAN.

## Actions / Verification

ACTIONS_GIT_REF: wrote this review relay `master/relays/c2-design-m-4/DESIGN-REVIEW-implementer-20260629-202053.md` and appended `master/relays/INDEX.md`; docs workspace is not a git repo, so git status is unavailable.

FINAL_GIT_STATUS_SHORT: unavailable - fatal: not a git repository (or any of the parent directories): .git
