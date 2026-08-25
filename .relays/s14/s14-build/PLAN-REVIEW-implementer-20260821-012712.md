## PLAN-REVIEW — s14 m-8 connector successor plan r3: APPROVE; the master-owned branch recipe is closed and executable, the r2 findings remain closed, and the seventeen-task contract is ready for direct implementation dispatch

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s14-build-plan-review-3
PARENT_DISPATCH_ID: s14-build
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this is plan approval only; the operator MERGE-GATE remains terminal and unchanged
GRILL_REQUIRED: no — the successor folds the exact master clarification and changes no frozen design mechanism
FILED_AT_LOCAL: 20260821-012712
IN_REPLY_TO: frank/.relays/s14/s14-build/PLAN-planner-20260821-004514.md
PLAN_LOCK_ID: s14-build @ sha256 84205d26f88aba475ecf15b39972f3fa47e9ea8f6e36b82b1e6e274f65d887a2
PLAN_REVIEW_VERDICT: approve
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: approve plan r3 — R4/R4a/R1 are folded closed, the branch-local dependency recipe is executable, and the seventeen-task boundary is preserved

## Verdict

`PLAN_REVIEW_VERDICT: approve` for exact plan r3 SHA-256 `84205d26f88aba475ecf15b39972f3fa47e9ea8f6e36b82b1e6e274f65d887a2`.

The sole r2 blocker closes. Master's append-only R4a clarification (`master/relays/step3-t4-preflight/RECONCILE-orchestrator-planner-20260821-003714.md`, SHA-256 `6f3f8dbdecacf1c3ca1e9e9e9652a4577e11cf154b2c7e39b21bae15963ee0cc`) explicitly authorizes s14 to set the identical `go 1.25` directive on its branch alongside `golang.org/x/text v0.41.0`. Plan r3 folds that authority into an executable first T1b commit, preserves the full-battery-at-every-commit rule, and makes any divergent restack byte an arbitration stop. The dependency edit no longer requires an unassigned shared-line mutation.

The requested fold passes:

- R4/R4a are closed, not conditional: Route A alone is operative; `norm.NFC.IsNormalString`, `golang.org/x/text v0.41.0`, the branch-local `go 1.25` directive, generated `go.sum`, and RED/GREEN commands are named. Route B is absent from operative instructions.
- R1's exact sprint-doc fence addition is carried with the non-relay-artifact rule and may remain explicitly unused.
- The task count is correctly seventeen: T1a, T1b, T1c, and T2–T15.
- T2–T15 remain defined by exact r2 §2, and §2a remains defined by exact r2 §2a. R3's summaries introduce no new local claim, sibling ownership, or fixture-boundary drift.
- The earlier S14-PR1-F1 and S14-PR1-F2 closures remain intact: RFC 8785 JCS preserves strings; NFC is a check-only validity rule at the catalog/policy loci; LOCAL/PRODUCER evidence stays distinct from RESTACK/INTEG and SIBLING proof.

No source, test, dependency, branch, worktree, stage, or commit was created or changed by this review. Approval authorizes the planner to issue the direct implementation dispatch; it is not itself implementation authority.

## Boundary contract review

Writes: connector source/tests, pair-local relays, the R1-authorized sprint-doc path if used, and the exact R4/R4a module bytes on the s14 branch.
Reads: frozen m-8/m-1/m-3/m-9/m-10 contracts, fake-counterpart frames, and the exact master arbitration relays.
Target entity: the m-8 connector and its reproducible Go dependency graph.
Downstream consumer: m-9, m-10, m-3, s16 integration, and the serialized restack/merge-gate sequence.
Contract: pure JCS, check-only NFC validity, honest fixture-boundary evidence, one x/text version, identical convergent Go directives, and arbitration on divergence.
Proof: task-local RED/GREEN legs, exact dependency/directive commit evidence, full repository battery on every commit, end-of-slice adversarial review, then restack rerun and re-review.
No-consumer action: do not implement until a literal direct `DISPATCH IMPL` relay is addressed solely to `s14.implementer` with this approved review in lineage.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — provider credential custody and authorize-before-attach remain in the slice
- migration/backfill/destructive-write/canonical-data-repair: no — no database/data migration or destructive repair is commissioned
- money/inventory/orders/planning/accounting/trust-critical-state: yes — provider-send authority and evidence digests remain trust-critical
- AI-or-automation-acts-downstream: yes — the connector sends model requests and normalizes provider output
- worker/scheduler/queue/retry/async-side-effect: yes — supervised process, streaming, cancellation, and one-shot provider effects
- cross-repo/service-contract/generated-schema/shared-API-event: yes — the Go directive/dependency graph and m-8/m-9/m-10/m-3 framed contracts are shared seams
- user-visible-control-with-materializer/downstream-consumer: no — no user-visible control is introduced by the plan
- test-runtime-role-mismatch: yes — s14's isolated branch must reproduce the ruled dependency under the Go directive its tests consume
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: no — the successor resolves the prior ambiguity and retains the frozen acceptance boundary
ESCALATION_SCAN_RESULT: trigger-present
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade is requested; retain production-risk
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none
PRE_SCAN_PRESSURE: none

ACTIONS_GIT_REF: lane-governance action only — this PLAN-REVIEW file plus its append-only s14 INDEX row; no repository source/test/dependency edit, branch, worktree, stage, commit, or implementation adoption
FINAL_GIT_STATUS_SHORT: (literal, frank/ cwd, captured after this relay and its INDEX row existed:)
 M .relays/s13/INDEX.md
 M .relays/s14/INDEX.md
 M ../master/relays/CHECKPOINTS.md
?? .relays/s13/s13-build-plan/
?? .relays/s14/s14-build/PLAN-REVIEW-implementer-20260821-012712.md

Next requested action: `s14.planner` issues the literal direct implementation dispatch parented by this approved review; `s14.implementer` remains held until that dispatch exists.
