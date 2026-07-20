## DESIGN-REVIEW - executable-claim FieldSpec home must revise nested ownership claim and open compatibility class

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-claim-input-m2-design-review
PARENT_DISPATCH_ID: s8-claim-input-m2-design
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded m-2 grammar review; master reconciliation is required for the parent Rail-A correction
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-claim-input-m2-home
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-claim-input-amendment/DESIGN-planner-20260712-005000.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-7.planner, m-3.implementer, m-7.implementer
SUBJECT: must revise - top-level seat_scoped_enum cannot constrain nested check_id, and ignored present claims make the transition closed rather than additive-open

DESIGN_REVIEW_VERDICT: must-revise

### Findings

1. **BLOCKER - the proposed owner/fill pair claims nested enforcement the live grammar does not provide.** The doc recommends `owner: seat_scoped_enum` + `fill_constraints: seat_allowed_values` “so `check_id` is picked only from the closed check catalog” (`design:21`). But `executable_claims` is a `row_array`; `check_id` is a nested column. The current grammar has no nested column schemas. `renderField` applies `seat_allowed_values` to the **whole field** through top-level `options`/`enum_set`/`seat_scope` (`render.go:101-129,154-177`), while `Validate` only seat-scope-checks the entire raw value when top-level `seat_scope` is non-empty (`validate.go:50-60`). `ParseTyped(row_array)` proves only canonical `[]map[string]string` shape; it does not validate column names or values (`canonical.go:20-30`). The existing `routing_assignments` row is a documented degraded carrier with column-grain enforcement deferred, not evidence that `seat_scoped_enum` constrains a nested token (`s5-escalations` M-3(f)).

   **Required revision:** keep the fixed m-2 claim at the honest grain: the top-level row is seat-declared and non-system-owned. Do not claim that its owner/fill pair constrains `check_id`. Either:
   - choose a top-level carrier owner/fill that accurately means seat-authored row data (for example the existing `agent_enum_pick`/`free_text` row-array pattern), with nested `check_id` validation delegated to the m-3-confirmed fill/observe validator seam; or
   - retain `seat_scoped_enum` only if the amendment also names and scopes a real nested-column grammar/enforcement mechanism. That would be a grammar extension and cannot be implied by the current bytes.

   The exact owner/fill pair therefore finalizes with the m-3 validation locus; it is not independently fixed by the routing precedent. R2 remains satisfied by top-level `gate_referenceable:false` and no `gate_referenceable_columns` either way.

2. **BLOCKER - the `v6→v7` Rail-A/additive-MINOR claim is false when a present declaration is ignored.** The doc equates an old reader ignoring the new row with valid optional absence and calls the transition OPEN/additive-MINOR (`design:23,29-33,48`). Those cases differ at acceptance. A v7-aware reader executes a present selected check; an observed false result rejects. A v6 reader that ignores the present row falls into `Evaluate:nil`; the no-vantage non-authority path can accept. Ignoring the field can therefore turn a would-be rejection into acceptance. This is a mechanically consumed, closed/fail-closed compatibility surface, not an ordinary optional header whose unknown value is semantically inert.

   The current m-3 adversarial review independently reaches the same blocker and requires correction through master because the parent amendment itself prescribed Rail-A additive (`s8-claim-input-m3/DESIGN-REVIEW-implementer-20260712-012200.md`, F1). m-2 must not finalize `s8-fieldspec-v7` as MINOR against a parent instruction now under owner review.

   **Required revision:** separate:
   - **optional absence for a compatible v7 reader**, which may remain the honest no-vantage degrade; from
   - **a present declaration encountered by an incompatible reader**, which must fail closed before submit semantics can diverge.

   Hold the exact v6→v7 compatibility/version class for master reconciliation with the revised m-3 return and m-7 reader-capability mechanism. Name the concrete refusal proof: a present `executable_claims` row cannot be silently treated as absent by a reader that cannot interpret/execute it. If any `EVIDENCE_TARGET` coupling later adds `required_when`, re-run the locked §9 classification again because requiredness changes are independently breaking/MAJOR (`v3-form-schema-design.md:126`).

### Accepted M-2 Grammar

- `executable_claims` as the seat-declared header `row_array` home paired with system-owned `executable_claim_results` is correct.
- `gate_referenceable:false` with no `gate_referenceable_columns` correctly keeps lane intent out of gate predicates and preserves R2.
- `visible_when layer_present:observe`, stale-form re-render, forward-only history, the input/output suppliability split, Rail B, I-PH, and the bounded byte-site/tripwire scope are correct.
- Exact columns, cardinality, check-catalog source, parameter schema, and validation locus remain owned by m-3; its current must-revise is not proxied here.
- The `s8-fieldspec-v6` base is verified on the build branch; main's older registry is not the source for this transition.

### Re-review Bar

Return a revision that removes the false nested enforcement claim and replaces the additive/open version lock with the master-reconciled fail-closed compatibility contract. No code, registry edit, PLAN, c1 reopen, or proxy m-3 semantic decision is authorized.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW of `s8-claim-input-m2-home` against the v6 build branch, live FieldSpec implementation, locked m-2 compatibility contract, s5 nested-row limitation, and current m-3 owner review; wrote this relay and appended `master/relays/INDEX.md`; no frank edit
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `691d034`; reviewed `s8-observe-spine@3cce8cd` carrying `s8-fieldspec-v6`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0; `--relay-root master/relays/s8-claim-input-amendment` lineage lint exit 0 for the live directory and this relay

Next requested action: m-2.planner corrects the top-level owner/enforcement claim and waits for master reconciliation of the Rail-A/version class against the revised m-3 semantics and m-7 capability gate before returning the m-2 home.
