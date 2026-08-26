## PLAN-REVIEW — s16a WP3-coda + WP4 PLAN r18 must revise: D1's verifier does not meet the controller's current lock, durable handover recording is unscoped, and D01/D02's seam predicates are contradictory or vacuous

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16a-build-plan-review-12
PARENT_DISPATCH_ID: s16a-build-18
RUN_ID: s16a
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the Planner routes the D1 locked-contract discrepancies through master to m-7, then issues a bounded successor; the operator's next gate remains WP5 MERGE-GATE
PLAN_LOCK_ID: s16a-build-18
IN_REPLY_TO: s16a-build/PLAN-planner-20260825-213247.md
FROM: s16a.implementer
TO: s16a.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-7.implementer, m-10.planner, m-9.planner, m-8.planner
SUBJECT: must revise PLAN r18 — F_GETLK cannot verify the current flock; durable control_handover recording has no writer/reader/ack scope; replace contradictory CT-D01 and vacuous CT-D02 predicates explicitly
PLAN_REVIEW_VERDICT: must-revise

`PLAN_REVIEW_VERDICT: must-revise` for exact PLAN r18 SHA-256 `33da00b57256b8fb89201896e38fc42b2129276796c82151599171d92e60b4b7` at `s16a-build/PLAN-planner-20260825-213247.md`.

This is the correct review target: FROM `s16a.planner`, TO this seat, plan-only, parented to the joined WP34 carriage, exact-file lint clean under `--no-freshness`, with no live implementation or merge token. The implementation worktree is clean at local/remote-equal `3566d37aafed8a51a2cd9effa4abbf658c9224fa`. Plan r9 and charter re-hash to `051c31f378895a1da4c923933a69112dd86fac04c3ac82f5ed51b763eb9faab9` and `b5aee083ede4716bd1f2ab55a24b985463654dbc547250049a5ed40c2f3ec516`.

## Findings

### S16A-WP34-PR-F1 — blocker: `F_GETLK` cannot verify the real controller's current lock

The owner contract pins one lock namespace end-to-end: controller `fcntl F_SETLK`, broker `fcntl F_GETLK`, holder PID equal to connected peer PID before token/generation (`master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md:192-200`). Production instead acquires/releases `syscall.Flock(..., LOCK_EX|LOCK_NB)` at `internal/appctl/brokerclient/session.go:192-213`. A record-lock probe does not prove ownership of that separate `flock`; on Linux an honest production controller will fail the proposed verifier.

R18 scopes D1 as `internal/broker/server.go` plus tests and inherits m-7's one-file sizing. The successor must route this spec/current-byte discrepancy UP through master to m-7, then authorize either the controller conversion to the pinned `F_SETLK` lock, including stable inode, no same-process close, and Linux/Darwin realization, or an owner ruling that changes verifier and holder together. A broker-only probe would regress real CT-C09 while appearing fail-closed.

### S16A-WP34-PR-F2 — blocker: recorded `control_handover` has no scoped writer/reader/ack path

The owner contract reserves “recorded” for durable m-10-store commit through the broker event/ack protocol (`…transport-broker.md:227-255`) and defines `control_handover` outcomes including `adopted` and the three rejections (`:244`). Current bytes have no complete path:

- `internal/appipc/msgs_broker.go:63-81` does not register `control_handover`.
- `internal/appctl/brokerclient/store.go:78-80` rejects every event except `boundary_cut` and `epoch_installed`.
- `internal/broker/server.go:22-173` has no event sequence, resend queue, event/ack exchange, or live-session replacement owner; `Serve` blocks serially inside the current session.

A log, in-memory slice, or test hook is not durable recording. The successor must name the in-fence writer, closed schema, durable reader/ack, resend/identity tests, and the session-ownership proof that a verified greater generation replaces and closes the old session. If the coda is narrower than owner bytes, master/m-7 must say so explicitly, including whether accepted adoption is recorded.

### S16A-WP34-PR-F3 — blocker: CT-D01 contradicts the ruling and CT-D02 is source-grep-vacuous

`test/seam/shared_test.go:21-27` requires all three canonicalizers to ACCEPT `1.5`; the ruling requires typed refusal at encode and digest. Lines 30-32 pass D02 from a source substring and execute no codec vector.

Revise sections 3-4 to authorize seam-predicate changes explicitly: CT-D01 proves byte-verbatim arbitrary-precision integers, including above 2^53, and typed refusal for `-0`, fraction, exponent, and Go float values at every applicable encode/digest entry; CT-D02 executes the exhaustive shared vector corpus against all three codecs and requires byte equality for every vector. Name these separately from ordinary module-test updates and preserve anti-vacuity. Merely greening the present predicates is not acceptance.

## Accepted portions

D2's 32-byte OS-CSPRNG/64-lowercase-hex token; D3's sixth disclosure locus; shared-package extraction; strict-integer/refuse-everywhere policy; removal of the three duplicate implementations; D03 exclusion; 64/0 close; per-commit suite/vet/census, sentinel, bijection, clean/pushed tip, DRAFT PR; and all no-merge/no-ready/no-E3 fences are accepted unchanged.

Boundary-contract disposition: D01/D02 have a writer/consumer shape once F3 is executable. D1 does not: the verifier does not meet the lock writer, and handover recording has no registered writer or durable reader. No-consumer action: reject-narrow those seams pending the owner-routed successor.

Tests / verification: exact plan, carriage, float ruling, m-7 review, r9/charter, owner sections 2.10-2.11, and current production/test loci read at bytes. Focused baseline E2 passed: `go test -count=1 ./internal/broker ./internal/appctl/brokerclient ./internal/appipc`. This is not proof of missing D1 behavior. No source/test byte changed.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — CI-1 controls broker controller authority and token custody
- migration/backfill/destructive-write/canonical-data-repair: no
- money/inventory/orders/planning/accounting/trust-critical-state: yes — handover and canonical evidence are trust-critical
- AI-or-automation-acts-downstream: yes — the broker authorizes governed tool operations
- worker/scheduler/queue/retry/async-side-effect: yes — session replacement and resend/event semantics
- cross-repo/service-contract/generated-schema/shared-API-event: yes — m-7/m-10 broker IPC plus three-module JCS
- user-visible-control-with-materializer/downstream-consumer: no
- test-runtime-role-mismatch: yes — CT-D01 asserts the opposite policy and CT-D02 asserts only source text
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — D1 needs added controller/event surfaces or an owner narrowing
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade or risk acceptance requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

Remaining gates: no implementation is authorized. The Planner carries F1/F2 UP through master to m-7, then files PLAN r19 addressing F1-F3 for fresh review. Approval still requires SCOPE_DIFF and a separate addressed implementation token. Merge, ready-for-review, CI/CD, E3, exit, deployment, and release remain unauthorized.

ACTIONS_GIT_REF: engine-lane review act only — prepared for daemon draft submission; installed role material is v2.9.2 while the running daemon remains fingerprint-compatible with v2.9.1, so submission uses the matching v2.9.1 client; no source, test, implementation-worktree, branch, commit, push, PR, merge, CI/CD, provider, credential, store, or runtime state changed
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s16a/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s16a/s16a-build/PLAN-planner-20260825-213247.md
