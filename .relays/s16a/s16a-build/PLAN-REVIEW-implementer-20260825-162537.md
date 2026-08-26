## PLAN-REVIEW — s16a WP3 PLAN r17 approved

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16a-build-plan-review-11
PARENT_DISPATCH_ID: s16a-build-17
RUN_ID: s16a
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the planner may perform SCOPE_DIFF and issue a separately addressed bounded WP3 implementation relay; merge and ready-for-review remain operator-gated at WP5
PLAN_LOCK_ID: s16a-build-17
IN_REPLY_TO: s16a-build/PLAN-planner-20260825-161904.md
FROM: s16a.implementer
TO: s16a.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-9.planner, m-10.planner, m-8.planner
SUBJECT: approve WP3 PLAN r17 — F1 production launch-to-READY composition and F2 executable C09 server matrix are closed on the pair-approved amendments; all r16 fences remain intact
PLAN_REVIEW_VERDICT: approve

`PLAN_REVIEW_VERDICT: approve` for exact PLAN r17 SHA-256 `c44224abddf94718513768d637b26d3e780f2801837b4198d525b603b6a90ebe` at `s16a-build/PLAN-planner-20260825-161904.md`.

The relay is the correct review target: FROM `s16a.planner`, TO this seat, plan-only, parented to my indexed `s16a-build-plan-review-10`, exact-file lint clean under the active daemon-compatible client, and carrying no live implementation or merge token. The existing linked implementation worktree is clean on `s16a-conformance`; local and remote tips both equal `ae58a0ac8b74a109a57cfce0c83e8869eb5de545`.

## Findings

None. The two r16 blockers are closed without reopening its accepted scope.

### S16A-WP3-PR-F1 — closed

R17 names a NEW separate production broker binary and non-conductor server tree, with m-10 as the launch/supervision owner. Production and CT-C09 consume the same production launcher and the same supervisor SPAWNING->READY observable. The coupled test strengthening is expressly authorized: launch the real production binary through that boundary, wait for the exact 84-byte `BROKER_READY nonce=<64 lowercase hex>\n` record to drive READY, then exercise the unchanged real `brokerclient.Establish` predicate with the real socket and app-applier/control-generation advance. The plan expressly excludes a fake listener, an in-client or in-process server, exit-zero readiness, and predicate weakening. That makes the production build and the binding seam test jointly executable.

### S16A-WP3-PR-F2 — closed

Section 2c now requires executable production-server tests plus the strengthened seam test for every m-7 agenda item 4-8/10/11: telemetry-residual `Describe` and study exclusions; valid and malformed correlation-staged proposal/result; the ordered total disposition table including BOOTSTRAP and a fail-closed unknown state; both trusted-tuple assignment forms and the wrong-tuple negative; two structurally different no-durable-state checks; all four F64 operation fences with durable settlement for the three effectful F59 calls and unfenced/duplicate negatives; and F60/F66 credential-custody negatives across wire, reply, log, and exposed surfaces. Each row must land with or before a named C09 preparatory commit; the final C09 commit must green CT-C09. The plan also preserves the governing rule that the READY record proves only launch step (3), never the deeper semantics.

## Approved boundary and execution shape

The exact m-7 launch amendment SHA-256 `0063650dfd55def3e08c47ff78831e332902e35fbe2d8379596d0bf2c89101be` and m-10 supervision registration SHA-256 `1d7505805506254459732d023f78a6e6cee23c4f7f60779bac56ee7185ba1d48` were re-hashed and match. The plan-of-record r9 SHA-256 `051c31f378895a1da4c923933a69112dd86fac04c3ac82f5ed51b763eb9faab9` and charter SHA-256 `b5aee083ede4716bd1f2ab55a24b985463654dbc547250049a5ed40c2f3ec516` also match.

A12, Starter/P5, C07, C09, their named commit grain, the moving census, exact 62/2 close with only D01/D02 RED, P4/P5 inversions, per-commit plain-suite/vet evidence, push-after-each-commit discipline, draft-only PR posture, no conductor growth, no durable broker state, dependency arbitration, and the no-E3/no-merge/no-ready fences are approved as written. The C09 matrix question rule is also correct: unresolved contract meaning routes UP through master to m-7 and stops local interpretation.

Boundary-contract disposition: the production launch writer, production/test consumers, shared readiness observable, operation/custody contracts, and named executable readers now have explicit owners and E2 acceptance. No reader is allowed to substitute readiness for deeper protocol proof.

Tests / verification: full r17 read and r16->r17 diff reviewed; exact r17 digest, INDEX parentage, plan/charter pins, and both amendment digests verified; daemon-compatible exact-file lint returned zero errors and zero warnings; implementation worktree branch/tips/status verified. No source tests were rerun because this act is review-only and changes no source or test byte.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — C09 includes F60/F66 credential custody and capability fencing, now bound to named positive and negative evidence
- migration/backfill/destructive-write/canonical-data-repair: no
- money/inventory/orders/planning/accounting/trust-critical-state: yes — the governed broker and tool fence are trust-critical, with exact contracts and stop conditions
- AI-or-automation-acts-downstream: yes — the worker and broker authorize tool operations, fenced by the approved matrix
- worker/scheduler/queue/retry/async-side-effect: yes — supervised processes, broker lifecycle, and durable settlement are explicit
- cross-repo/service-contract/generated-schema/shared-API-event: yes — the m-7/m-10/m-9 shared IPC and operation contract is bound to approved amendments
- user-visible-control-with-materializer/downstream-consumer: no
- test-runtime-role-mismatch: no — F1 now requires the same production launcher and READY observable in production and CT-C09
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: no — scope, escalation stop, and executable evidence are explicit
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade or risk acceptance requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

Remaining gates: this verdict authorizes no implementation. The planner must perform SCOPE_DIFF and file a separately addressed IMPL relay carrying the bare implementation token and parented to this approval before any WP3 byte changes. Merge, ready-for-review, CI/CD, E3, and release remain unauthorized.

ACTIONS_GIT_REF: engine-lane governance act only — this review is drafted under `frank/.relays/s16a/.engine/drafts/s16a.implementer/` for daemon submission; active installed role material is v2.9.2 while the running relay daemon remains fingerprint-compatible with v2.9.1, so admission uses the daemon-matching v2.9.1 client without restarting or replacing daemon state; no source, test, implementation-worktree, branch, commit, push, PR, merge, CI/CD, provider, credential, store, or runtime state changed
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s16a/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s16a/s16a-build/PLAN-planner-20260825-161904.md
