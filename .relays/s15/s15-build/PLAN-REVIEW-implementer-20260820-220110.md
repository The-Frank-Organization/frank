## PLAN-REVIEW — s15 m-9 worker build plan r1: MUST-REVISE; the out-of-fence MCP seam is a blocking locked-contract obligation, the plan lacks a dispatchable all-in scope and executable task acceptance, and the report-only corrigendum must become a successor PLAN

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s15-build-plan-review-1
PARENT_DISPATCH_ID: s15-build
RUN_ID: s15
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the planner must revise; master must arbitrate the already-named fence/locked-surface deviation before any all-in implementation dispatch
GRILL_REQUIRED: no — the ratified GRILL_LOCK rides unchanged
FILED_AT_LOCAL: 20260820-220110
IN_REPLY_TO: frank/.relays/s15/s15-build/PLAN-planner-20260820-204124.md
PLAN_LOCK_ID: s15-build @ sha256 d5cfaefee7708f8805efa1555c7798df9efd3bc92efa017499635fff52906864
PLAN_REVIEW_VERDICT: must-revise
FROM: s15.implementer
TO: s15.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer
SUBJECT: must-revise — make the MCP/shared-module seam blocking and all-in, replace report-only supersessions with a successor PLAN, add production-risk scan + per-task acceptance/boundary proofs

## Verdict

`PLAN_REVIEW_VERDICT: must-revise`. The fourteen-task direction is broadly aligned with the frozen m-9 surfaces, and the pre-corrigendum counter bytes are inventoried faithfully, but the current PLAN is not dispatchable. Four blocking findings follow. No source/test byte was adopted, edited, staged, or committed in this review.

## S15-PR1-F1 — BLOCKER: the plan calls a required shared-module/MCP obligation non-blocking and places incompatible concerns in one package

The PLAN puts the m-2 mapping module, the channel-using shared client facade, and the native relay frontend together under `internal/seatclient` (`PLAN:34`), leaves `cmd/frank-mcp/**` untouched pending a later ruling (`:41-42`), declares MCP behavior unchanged (`:44-45`), and calls the seam non-blocking (`:53`). That cannot close the frozen m-2 contract:

- The mapping module must not import `internal/channel`; `internal/channel` must not import it; both frontends wire the two as consumers (`master/domains/m-2-forms-determinism/design/2026-07-16-step3-mvp-form-schema-mapping.md:31-40`). A channel-using facade and the mapping code therefore need an explicit package/import split; one undifferentiated `internal/seatclient` package is not a sufficient plan.
- The retained MCP frontend and native tool must consume the SAME module, mechanically enforced by an import-graph test (`mapping.md:38-39`). Temporary duplication is not contract completion.
- The build obligations explicitly require `cmd/frank-mcp` to become a consumer, the MCP dispatch gate/freshness choreography, and the MCP parity adapter (`mapping.md:253-264`). “MCP frontend behavior unchanged” contradicts those obligations because several are deliberate fail-closed deltas from the trapped implementation.

Required revision: obtain master's durable ruling before onward dispatch. If the fence is extended, name the exact `cmd/frank-mcp/**` files, the isolated mapping-module package, the channel/client package, both import directions, and the shared parity/refresh tests. If the extension is denied or deferred, narrow the dispatch before T7 and state that the slice cannot claim the m-2 module/parity obligations or close T13/T14; do not label the seam non-blocking.

## S15-PR1-F2 — BLOCKER: the plan already contains an OUT row but no mechanical scope diff

The charter fence is limited to `cmd/frank-worker/**`, `internal/worker/**`, `internal/seatclient/**`, and `.relays/s15/**`; `cmd/frank-mcp/**` is outside it. The PLAN itself identifies that deviation (`PLAN:41-42`). Under the delegated-dispatch protocol, the mechanical scope diff must therefore classify `frank/cmd/frank-mcp/** -> OUT`, yielding `SCOPE_DIFF_RESULT: deviation-present`; a planner may not issue the build token until master changes the fence or the locked plan removes that touch.

Required revision: include the complete per-path scope inventory against the charter fence and the returned master ruling. The later implementation-dispatch relay must carry exactly one `SCOPE_DIFF` and `SCOPE_DIFF_RESULT: all-in`; no “small integration hook,” “at restack,” or acceptance-needed exception converts OUT to IN.

## S15-PR1-F3 — BLOCKER: the task list is not executable acceptance criteria for this production-risk slice

The complete T1-T14 execution plan is one line (`PLAN:38-39`). It does not bind each task to exact frozen clauses, named files, RED tests, completion predicates, boundary writer/reader proof, or task-specific E2 commands. T13 is only “the §10/§6 fixture-battery sweep,” and T14 says “docs” without an in-fence path. That is insufficient for a worker/async/cross-process/authz slice and leaves two concrete hazards: accidentally treating the untouchable exit-fixture corpus as an editable build surface, and silently inventing an out-of-fence documentation path.

The handoff SITREP's JCS material (`SITREP-planner-20260820-214014.md:37-41`) is useful but is only “recommended” report-only context; it does not repair the PLAN's acceptance contract.

Required revision: for every T1-T14 row, name (a) exact governing clauses/basis, (b) exact in-fence artifact paths, (c) RED-first negative and positive batteries, (d) a hard done predicate and E2 command, (e) writer/reader/fake counterpart and proof where a boundary exists, and (f) explicit out-of-scope preservation. Integrate the JCS requirements into the successor PLAN. Name T14's documentation/store-export paths and correct T14's reviewer to `s15.planner`.

## S15-PR1-F4 — BLOCKER: the authoritative PLAN remains role-inverted and ceremony-underrated

The PLAN of record says review is asynchronous and assigns the end review to the implementer (`PLAN:18,38-39,53`). The later report-only SITREP supersedes those sentences in prose (`SITREP:24-26`) but says no other PLAN byte changes. A report-only correction is not a stable substitute for the written plan that will parent the plan review and delegated dispatch. Reissue a successor PLAN containing the corrected role flow in its own bytes.

The PLAN also declares `CEREMONY_TIER: medium` (`PLAN:9`) and carries no escalation scan, while this slice includes F59 authorization/tickets, a worker, retries/async recording, provider and app IPC contracts, a credentialed seat client, bash subprocesses, and shared API/module seams. Those are explicit hard-trigger classes. The commissioning relay correctly used `production-risk`; the pair plan must do likewise unless the operator grants a post-scan informed downgrade waiver.

Required revision: issue `s15-build-2` (or the next mechanical successor id) at `production-risk`, with the corrected planner/implementer flow in the PLAN itself and a completed escalation scan. No downgrade is requested or justified here.

## Pre-corrigendum byte disposition

The handed-off hashes match disk exactly:

- `frank/internal/worker/wire/counter.go` = `6bdd84f125fc47e02cc4de8e6cdce0738902afe8504b047382aa63fed2177c60`
- `frank/internal/worker/wire/counter_test.go` = `d52f63b42349c9ff636aaff26c28fa18a94500d16f9b81518756af86535d8be6`

They remain untracked and untouched. Adoption/re-landing is an IMPL act and waits on an approved successor PLAN plus a valid bare own-line implementation dispatch addressed only to `s15.implementer`.

## Boundary contract review

Writes: worker packages, shared seat-client artifact, native relay frontend, and the m-2 mapping implementation.
Reads: the frozen m-9/m-2/m-7/m-8/m-10/m-3 contracts and fake counterpart frames.
Target entity: the m-9 worker artifact plus the separately identified shared-client artifact.
Downstream consumer: `frank-worker`, m-10 supervisor, m-8 connector, conductor channel, retained MCP frontend, and stage-6 release binding.
Contract: the frozen frame/tool/journal/digest surfaces; especially one shared mapping module consumed by both frontends.
Proof: revised task-level E2 batteries, import-graph test, shared parity vectors, fake m-10/m-8 E2E, and full `go test ./...` at each commit.
No-consumer action: reject dispatch/close until the retained MCP frontend consumes the shared module or master explicitly narrows the commissioned completion surface.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — F59 authorization/tickets and the credentialed seat-client boundary
- migration/backfill/destructive-write/canonical-data-repair: no — no database/data migration is in the slice
- money/inventory/orders/planning/accounting/trust-critical-state: yes — governed tool authority and the durable session-content trust boundary
- AI-or-automation-acts-downstream: yes — model-requested local and relay tools produce effects
- worker/scheduler/queue/retry/async-side-effect: yes — worker lifecycle, retries, async outcome recording, and provider attempts
- cross-repo/service-contract/generated-schema/shared-API-event: yes — m-2/m-7/m-8/m-10 shared contracts and MCP/native parity
- user-visible-control-with-materializer/downstream-consumer: no — no user-visible control is introduced by this plan
- test-runtime-role-mismatch: yes — the pre-corrigendum planner-authored implementation bytes require explicit implementer adoption or re-landing
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — the unresolved out-of-fence MCP seam and vague T13/T14 completion surface
ESCALATION_SCAN_RESULT: trigger-present
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade is requested; use production-risk
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none
PRE_SCAN_PRESSURE: none

ACTIONS_GIT_REF: lane governance action only — this PLAN-REVIEW file plus its append-only s15 INDEX row; no source/test edit, stage, commit, branch move, or adoption
FINAL_GIT_STATUS_SHORT: (literal, frank/ cwd, captured after this relay and its INDEX row existed:)
 M .relays/s15/INDEX.md
 M ../master/relays/CHECKPOINTS.md
?? .relays/s15/s15-build/PLAN-REVIEW-implementer-20260820-220110.md

Next requested action: `s15.planner` returns a successor PLAN closing F1-F4. Only an approved successor review may parent the delegated implementation dispatch.
