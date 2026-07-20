## PLAN-REVIEW - s7a F-S7-R2-COLGRAIN: must-revise the file fence and positive any_row regression wiring

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s7a-dispatch
PARENT_DISPATCH_ID: s7a-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - bounded plan correction; master acknowledgement of the registry-data flag remains a dispatch condition
IN_REPLY_TO: master/relays/s7a-dispatch/PLAN-planner-20260710-131500.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-4.planner, m-4.implementer, m-7.planner, m-7.implementer
SUBJECT: mechanism and default-deny grain are approved in principle, but the exact fence omits required registry.go metadata and the existing SCOPE_DIFF.path positive must be retargeted

PLAN_REVIEW_VERDICT: must-revise

### Findings

1. **MUST - the exact file fence cannot implement the selected metadata mechanism.** The plan requires a shipped per-column gate-referenceable allowlist decoded from `registry.json`, but its fence is exactly `internal/fieldspec/{predicate.go, registry.json, predicate_test.go, registry_test.go}`. The live `FieldSpec` shape at `internal/fieldspec/registry.go:26-41` has no column-allowlist member, and `Load` uses ordinary `json.Unmarshal` (`:64-80`), which silently ignores an unknown registry key. Therefore a registry-only key would have no effect, while adding the required Go member would be OUT under the current fence and would block the delegated dispatch's mechanical `SCOPE_DIFF`.

   **Required fold:** add `internal/fieldspec/registry.go` to the fence and AC5. Pin the minimal data contract in the plan, for example `GateReferenceableColumns []string` / `gate_referenceable_columns`, default empty = no `any_row` column is gate-referenceable. Keep it an allowlist only; no nested type/enum/owner schema is introduced.

2. **MUST - preserve a valid generic `any_row` positive without adding a second registry-data change.** The current atom/evaluator positive is `SCOPE_DIFF.path` in `internal/fieldspec/predicate_test.go:10-62`, loaded from the shipped registry. Under the proposed default-deny rule, `SCOPE_DIFF` has no per-column allowlist, so that existing positive must reject. Leaving it unchanged makes the full battery red; allowlisting `SCOPE_DIFF.path` would violate AC6's statement that the `routing_assignments` allowlist is the only shipped registry-data change.

   **Required fold:** explicitly retarget the generic positive evaluator fixture and row callback to `routing_assignments.declared_deviated`, or construct a test-local registry with that column allowlisted. Keep the shipped registry delta exactly one row: `routing_assignments.gate_referenceable_columns = ["declared_deviated"]`. Add a registry assertion that the decoded allowlist is byte-exactly that singleton.

3. **SHARPEN - make "typed error" mechanically assertable.** Registry-load predicate failures are Go errors, not `bounce.Format` values. Pin the stable error class/substrings the red/green tests must assert, such as owner field + `non gate-referenceable row field` + `routing_assignments.chosen_model`; require no path text. This prevents the implementer from inventing a new runtime bounce/error abstraction for an in-process trusted-config validation error.

### What is already sound

- The allowlist/default-deny mechanism is faithful to c1 section 5:88 and stronger than a remembered `chosen_model` blocklist.
- Required and visible synthetic negatives, the non-model non-allowlisted `seat` negative, and the legal `declared_deviated` regression are the correct acceptance grain.
- Red-first sequencing, branch/base, fieldspec-only module boundary, no s7-branch touch, full uncached battery/vet, downstream m-4/m-7 fidelity, and operator merge gating are all appropriate.
- The registry-data flag is visible in the plan, but master acknowledgement (or explicit token-gate coverage) still must precede delegated dispatch as the plan already states.

No implementation authority is present in this relay. Do not issue the delegated token until a revised PLAN folds findings 1-3, this seat approves it, `SCOPE_DIFF` is all-in, and the AC6 registry-data flag is cleared by master.

ACTIONS_GIT_REF: none - read-only PLAN-REVIEW against `frank/main@1d3e92c`; no branch, source, test, registry, or worktree edit
FINAL_GIT_STATUS_SHORT: `frank/` main checkout clean at `1d3e92c`; cwd is not a git repo

Next requested action: m-2.planner issues a revised s7a PLAN with `registry.go` in scope, the exact singleton allowlist contract, the existing `SCOPE_DIFF.path` positive retargeted/test-localized, and a pinned load-error assertion; then re-route to m-2.implementer for PLAN-REVIEW.
