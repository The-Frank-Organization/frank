## PLAN-REVIEW — s15 m-9 worker successor plan r2: MUST-REVISE narrowly; prior F1–F4 close, but Branch B lacks a semantic-narrowing/follow-up gate and T14 mistakes transport for the required store export

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s15-build-plan-review-2
PARENT_DISPATCH_ID: s15-build-2
RUN_ID: s15
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the planner can repair T14; master must make any Branch-B semantic narrowing and follow-up assignment explicit in its pending ruling
GRILL_REQUIRED: no — the ratified GRILL_LOCK rides unchanged
FILED_AT_LOCAL: 20260820-223848
IN_REPLY_TO: frank/.relays/s15/s15-build-2/PLAN-planner-20260820-222809.md
PLAN_LOCK_ID: s15-build-2 @ sha256 5a045663547838384ec43d42d7f4073ff68d5ff1113be9091ade575849c2dcb5
PLAN_REVIEW_VERDICT: must-revise
FROM: s15.implementer
TO: s15.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer
SUBJECT: narrow must-revise — make Branch B require explicit semantic narrowing plus a named MCP follow-up gate; restore the master-owned verbatim store export at T14

## Verdict and passed surface

`PLAN_REVIEW_VERDICT: must-revise`, narrowed to two blockers. Successor r2 materially closes S15-PR1-F1 through F4: the role flow is authoritative in-plan, production-risk ceremony and scan are correct, the package/import split is explicit, the seam is blocking, the path inventory is mechanical, and T1–T14 now carry usable acceptance/boundary/preservation criteria. Those surfaces do not reopen.

## S15-PR2-F1 — BLOCKER: a Branch-B fence ruling does not itself narrow a locked semantic obligation

The plan correctly says the frozen m-2 contract requires both frontends to consume the same module (`PLAN:28-40`), but Branch B then allows a fence denial/defer to exclude the MCP consumer, dispatch-gate/freshness, and parity-adapter obligations from the slice completion claim (`:41`, `:111-117`, `:159-170`). Section 5 gates dispatch only on master's durable “fence ruling” (`:194-199`).

A path ruling can make the touched-path `SCOPE_DIFF` all-in; it cannot by itself make mandatory m-2 §6 obligations non-applicable. The frozen build contract requires `cmd/frank-mcp` to become a consumer and names the MCP halves as build obligations (`master/domains/m-2-forms-determinism/design/2026-07-16-step3-mvp-form-schema-mapping.md:253-264`).

Required revision: Branch B is dispatchable only if master's returned ruling explicitly does all three: (1) narrows s15's completion claim rather than merely denying/defering the fence, (2) names the owner/dispatch/gate that must land the omitted MCP obligations, and (3) requires T13/T14 to return the slice as partial/held on that named obligation rather than complete. Otherwise Branch A is the only valid completion path. Carry this semantic condition in §1, the scope table, T7, T13, T14, and §5; an all-in path diff is necessary but not sufficient.

## S15-PR2-F2 — BLOCKER: T14 conflates live-store transport with the required durable export

T14 currently says “the store export = the closing lane SITREP submitted via the seat channel to master” (`PLAN:167-170`). That is transport, not export. CYCLE-PLAYBOOK F.7.2 requires master at each slice close to export the slice's verbatim store-record JSONs, content-addressed by `relay_id`, into `frank/.relays/s15/store-export/` and cite the export in the close relay (`master/CYCLE-PLAYBOOK.md:431-437`). The charter separately binds “the store records export at slice close” (`master/subteams/s15-m9-worker/CHARTER.md:32-35`).

Required revision: T14 must distinguish (a) the pair's closing SITREP submitted through the courier from (b) master's export act. The pair's return requests the master-owned export; slice close waits for a master receipt citing the populated `frank/.relays/s15/store-export/` artifacts, or labels the export explicitly outstanding and therefore does not claim slice close. Do not assign master-authored export bytes to either pair seat.

## Pre-corrigendum bytes

The two handed-off counter files remain untracked and untouched. Their disposition remains an IMPL act only after an approved successor and a valid implementation dispatch.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — F59 and the credentialed seat-client boundary remain in scope
- migration/backfill/destructive-write/canonical-data-repair: no — no migration or canonical-data repair
- money/inventory/orders/planning/accounting/trust-critical-state: yes — governed authority and durable trust state
- AI-or-automation-acts-downstream: yes — model-requested tools act downstream
- worker/scheduler/queue/retry/async-side-effect: yes — worker, attempts, retries, and async outcome recording
- cross-repo/service-contract/generated-schema/shared-API-event: yes — m-2/m-7/m-8/m-10 contracts and MCP/native parity
- user-visible-control-with-materializer/downstream-consumer: no — no user-visible control introduced
- test-runtime-role-mismatch: yes — pre-corrigendum planner-authored bytes remain pending implementer adoption/re-landing
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — Branch-B semantic narrowing and the close/export gate are unresolved
ESCALATION_SCAN_RESULT: trigger-present
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade requested; production-risk stands
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none
PRE_SCAN_PRESSURE: none

ACTIONS_GIT_REF: lane governance action only — this PLAN-REVIEW file plus its append-only s15 INDEX row; no source/test edit, stage, commit, branch move, or adoption
FINAL_GIT_STATUS_SHORT: (literal, frank/ cwd, captured after this relay and its INDEX row existed:)
 M .relays/s13/INDEX.md
 M .relays/s13/docs/designs/DS-s13-m10-module-20260820.md
 M .relays/s13/docs/plans/PL-s13-build-plan-20260820.md
 M .relays/s14/INDEX.md
 M .relays/s15/INDEX.md
 M ../master/relays/CHECKPOINTS.md
?? .relays/s13/s13-build-design/DESIGN-REVIEW-implementer-20260820-221508.md
?? .relays/s13/s13-build-design/DESIGN-planner-20260820-223212.md
?? .relays/s13/s13-build-design/SITREP-planner-20260820-223211.md
?? .relays/s14/s14-build/
?? .relays/s15/s15-build-2/
?? .relays/s15/s15-build/PLAN-REVIEW-implementer-20260820-220110.md

Next requested action: `s15.planner` returns successor PLAN r3 closing only S15-PR2-F1/F2. The prior-review closures stand.
