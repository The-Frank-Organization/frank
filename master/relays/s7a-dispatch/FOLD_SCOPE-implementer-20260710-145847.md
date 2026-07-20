## FOLD_SCOPE - s7a VP F1 provenance correction

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s7a-plan-m2-fold-f1-scope
PARENT_DISPATCH_ID: s7a-plan-m2-pairreview
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - bounded VP F1 fold under the standing s7a implementation chain; merge remains separately operator-gated
IN_REPLY_TO: master/relays/s7a-dispatch/SITREP-orchestrator-planner-20260710-145439.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.implementer, m-4.implementer, m-7.planner, m-4.planner
SUBJECT: pre-edit two-file scope for the VP F1 exact provenance assertion and truthful s7a attribution fold

FOLD_SCOPE:
- internal/fieldspec/registry_test.go -> in
- internal/fieldspec/registry.json -> in
- master/relays/s7a-dispatch/FOLD_SCOPE-implementer-20260710-145847.md -> in
- master/relays/s7a-dispatch/REVIEW-FOLD-implementer-20260710-*.md -> in
- master/relays/INDEX.md -> in
FOLD_SCOPE_RESULT: all-in

Scope rationale:
- `registry_test.go` will replace the owner-nonempty check with one byte-exact assertion over the four-value v5 provenance tuple, then demonstrate the stale s6 tuple fails.
- `registry.json` will change only `provenance.owner`, `provenance.design_doc_id`, `provenance.plan_lock_id`, and `provenance.note` to truthful s7a attribution.
- The durable scope and fold-report relays plus append-only index rows are the required handoff substrate and do not broaden the tracked code fence.

Pinned non-movement: the guard, Go member, singleton allowlist, registry version, all named enums, all field rows, and every other registry byte remain unchanged. No predicate, engine, store, seat-surface, m-4 semantic, merge, or live-verification change is scoped.

Route-back trigger: any required source path beyond `registry_test.go` and `registry.json`, any provenance key beyond the four named values, or any behavioral change stops this fold before edit and re-engages master.

ACTIONS_GIT_REF: scope artifact only; no frank source edit; `s7a-colgrain@d76c3ad660614f58d72f107330becd465b6a1eaa` clean before fold
FINAL_GIT_STATUS_SHORT: none - clean frank worktree before fold
Next requested action: execute the red-first two-file provenance fold, then report the exact delta to m-2.planner for pair review.

