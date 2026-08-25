## PLAN-REVIEW — s14 m-8 connector successor plan r2: MUST-REVISE narrowly; F1/F2 close, but the later R4 ruling reverses the plan's dependency writer/fence and leaves a Go-directive collision unresolved on the concurrent launch branch

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s14-build-plan-review-2
PARENT_DISPATCH_ID: s14-build
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — s14.planner owns the successor plan and routes the branch-time go.mod seam clarification to master; no operator decision is required before that correction
GRILL_REQUIRED: no — the frozen contracts and prior review findings remain unchanged; this is a post-plan arbitration/branch-executability defect
FILED_AT_LOCAL: 20260820-235513
IN_REPLY_TO: frank/.relays/s14/s14-build/PLAN-planner-20260820-234459.md
PLAN_LOCK_ID: s14-build @ sha256 11d8d99f239cabec2af7d8e47fd927011661583874f7bb5545ba431825c2d1a4
PLAN_REVIEW_VERDICT: must-revise
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer
SUBJECT: narrow must-revise — fold R1/R4 into the plan and obtain an executable ruling for x/text's go-1.25 requirement on the pre-s13 launch branch

## Verdict and passed surface

`PLAN_REVIEW_VERDICT: must-revise`, narrowed to one post-plan blocker. Successor r2 materially closes both prior findings:

- S14-PR1-F1 closes: T1a is pure string-preserving RFC 8785; T1b is check-only artifact-specific NFC validity at the policy/catalog loci; non-NFC rejects and valid NFC non-ASCII accepts; no accepted byte is rewritten.
- S14-PR1-F2 closes: §2a classifies every r12 §8 and addendum §6 leg as LOCAL, PRODUCER, RESTACK/INTEG, or SIBLING, and §3 binds local completion only to the first two classes. Cross-lane rows and m-3 evaluator truth are not claimed locally.
- All r1 surfaces previously passed remain closed. Exact r2 hash reproduced as `11d8d99f239cabec2af7d8e47fd927011661583874f7bb5545ba431825c2d1a4`; exact-file lint passes.

No source/test byte, dependency file, branch, worktree, stage, or commit was created or changed in this review. The temporary E2 module probe lived outside the repository under `/tmp` and tested only the ruled dependency's Go-version behavior.

## S14-PR2-F1 — BLOCKER: R4 landed after r2 and makes the locked plan contradictory; its current branch recipe cannot keep the full battery green without an unassigned shared-line edit

Plan r2 was filed at `23:44:59`; master's lint-clean ruling followed at `23:45:01` (`master/relays/step3-t4-preflight/RECONCILE-orchestrator-planner-20260820-234501.md`, SHA-256 `17ea6b3bd3fc2c6db7c69a15e0a1548988131ed439912a74eaeb0074c7307455`). The ruling is addressed to `s14.planner` in TO and reaches this reviewer in CC. It changes the facts the direct plan still presents as open:

- r2's write fence says `go.mod`/`go.sum` are arbitration-only, Route A or Route B remains conditional, and “s14 edits no go.mod/go.sum byte under either route” (`PLAN:30-32`).
- R4 selects only Route A, rejects Route B, pins `golang.org/x/text@v0.41.0`, and explicitly names s14 as the branch writer of `go.mod`/`go.sum` (`RULING:9`). R1 also grants the s14 sprint-doc fence addition (`RULING:3`).

Those contradictions require a successor PLAN before delegated dispatch. More importantly, the selected module exposes an unresolved concurrent-branch collision:

- E1: the downloaded `golang.org/x/text@v0.41.0` module declares `go 1.25.0`; s14's pinned `LAUNCH_BASE` declares `go 1.22`.
- E2 in an isolated temporary module: with `go 1.22` plus `require golang.org/x/text v0.41.0`, `go mod tidy` rewrote the directive to `go 1.25.0`; after restoring `go 1.22`, `go test -mod=readonly ./...` refused with `go: updates to go.mod needed`, while `-mod=mod` changed the directive back to 1.25.
- R2 names s13 as the writer of the shared `go 1.22 → 1.25` directive (`RULING:5`), while R4 names s14 for the x/text dependency and says serialized restack resolves the join (`:9`). Restack occurs before s14's terminal merge gate, but s14 must build and run the full battery on its independent launch branch before that point. The current text does not explicitly authorize s14 to touch the directive on its branch or tell s14 to wait for the s13 restack before T1b; either silent choice violates a named rule.

Required revision:

1. Obtain a durable master successor/clarification choosing an executable branch recipe. The direct route is to authorize s14 to set the same `go 1.25` directive on its own branch alongside `x/text@v0.41.0`, name the intentional convergent overlap with s13, and retain serialized restack as the reconciliation proof. Alternatives must preserve full-battery-at-every-commit and the concurrent-branch contract explicitly; do not infer ownership from ambiguous “except under R4” prose.
2. Reissue the PLAN with R4 closed, not conditional: Route A only; `norm.NFC.IsNormalString`; x/text v0.41.0; exact s14-owned `go.mod`/`go.sum` changes and RED/GREEN commands; Route B removed from operative text.
3. Carry R1's exact s14 sprint-doc fence addition and its non-relay-artifact rule, whether the path is used or explicitly preserved unused. Update the mechanical scope inventory used by the later dispatch.
4. Correct the headline/task count while folding: T1a + T1b + T1c plus T2–T15 is seventeen tasks, not sixteen. This count defect is editorial, not an independent blocker.

## Boundary contract review

Writes: connector source/tests, pair-local relays, the R1-authorized s14 sprint-doc path if used, and the exact dependency/module bytes a clarified R4 assigns to s14.
Reads: frozen m-8/m-1/m-3/m-9/m-10 contracts, fake counterpart frames, and master R1/R4 arbitration.
Target entity: the m-8 connector plus its reproducible Go dependency graph on the s14 branch.
Downstream consumer: m-9, m-10, m-3, s16 integration, and the serialized restack/merge-gate sequence.
Contract: pure JCS, check-only NFC validity, honest fixture-boundary evidence, one x/text version, one Go directive, and deterministic shared-line reconciliation.
Proof: revised exact PLAN; master-authored branch recipe; isolated `go mod tidy`/`go test -mod=readonly` regression; full repo battery on every implementation commit and after restack.
No-consumer action: reject delegated dispatch until the branch can consume x/text and run tests without an unassigned or out-of-fence go-directive mutation.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — provider credential custody and authorize-before-attach remain in the slice
- migration/backfill/destructive-write/canonical-data-repair: no — no database/data migration or destructive repair is commissioned
- money/inventory/orders/planning/accounting/trust-critical-state: yes — provider-send authority and evidence digests remain trust-critical
- AI-or-automation-acts-downstream: yes — the connector sends model requests and normalizes provider output
- worker/scheduler/queue/retry/async-side-effect: yes — supervised process, streaming, cancellation, and one-shot provider effects
- cross-repo/service-contract/generated-schema/shared-API-event: yes — the Go directive/dependency graph and m-8/m-9/m-10/m-3 framed contracts are shared seams
- user-visible-control-with-materializer/downstream-consumer: no — no user-visible control is introduced by the plan
- test-runtime-role-mismatch: yes — s14's isolated branch must reproduce the ruled dependency under the same Go directive its tests consume
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — post-plan R1/R4 authority changes and ambiguous shared-line ownership must be folded, not inferred
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
 M .relays/s14/INDEX.md
 M ../master/relays/CHECKPOINTS.md
?? .relays/s14/s14-build/PLAN-REVIEW-implementer-20260820-235513.md

Next requested action: `s14.planner` obtains master's explicit branch-time Go-directive recipe and returns successor PLAN r3 folding R1/R4 plus the corrected task count. Only an approved successor PLAN-REVIEW may parent the literal direct implementation dispatch.
