## PLAN-REVIEW — s16a WP3 PLAN r16 must revise the C09 launch/test composition and its executable server-semantic evidence

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16a-build-plan-review-10
PARENT_DISPATCH_ID: s16a-build-16
RUN_ID: s16a
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — s16a.planner can issue a bounded plan successor; implementation and merge remain separately gated
PLAN_LOCK_ID: s16a-build-16
IN_REPLY_TO: s16a-build/PLAN-planner-20260825-084308.md
FROM: s16a.implementer
TO: s16a.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-9.planner, m-10.planner, m-8.planner
SUBJECT: must-revise C09 only — authorize the production server's launch-to-READY test composition and require E2 tests for m-7 flags 4–8/10/11; all other WP3 scope and fences accepted
PLAN_REVIEW_VERDICT: must-revise

`PLAN_REVIEW_VERDICT: must-revise` for exact PLAN r16 SHA-256 `1954b7b72386951359e9e337b5914c06c6a422f7e57d7c5f6d2f3c81b2e95943` at `s16a-build/PLAN-planner-20260825-084308.md`.

The relay is structurally the correct review target: FROM `s16a.planner`, TO this seat, plan-only, parented to the indexed WP2 close, exact-file lint clean, and carrying no live implementation token. The worktree is the existing clean linked worktree on `s16a-conformance`; local and remote tips equal `ae58a0ac8b74a109a57cfce0c83e8869eb5de545`. Two coupled C09 blockers remain before dispatch.

## Blocking findings

### S16A-WP3-PR-F1 — CT-C09 cannot turn green from the authorized production build while its launch precondition is absent and its test byte is fenced

The current binding test at `test/seam/connector_app_test.go:230-239` creates an app-store fixture and an empty temporary runtime directory, then immediately calls the real `brokerclient.Establish`. Neither that test nor `newAppFixture` starts a broker. The production client only dials `broker-control.sock`; it does not and should not spawn the server. The current non-test tree has no broker binary, no broker listener, and no broker launcher. A focused E2 run reproduced the exact contract RED: `dial unix .../broker-control.sock: connect: no such file or directory`.

R16 correctly requires the server to remain a separate supervised process/binary and forbids conductor growth. Therefore the unchanged CT-C09 can become green only through an architecturally wrong implicit/in-process server inside the client or unrelated fixture magic. But r16 authorizes a coupled seam-test strengthening only for C07 (§2.3); its general fence says coupled strengthenings ONLY under that authorization. The plan simultaneously requires CT-C09 green "from the BUILD, never from test edits." Those constraints are not jointly executable.

### S16A-WP3-PR-F2 — the C09 close proof does not mechanically cover the m-7 server checklist

The existing CT-C09 reaches only CI-1 establishment through the handshake-frame write. M-7's own fidelity return explicitly says it does not bind the full CI-1/disposition semantics. R16 imports flags 4–8, 10, and 11 as the build checklist, but acceptance names only CT-C09, the ordinary suite, vet, and a later m-7 review. It does not require new production-server tests for: `Describe` telemetry-residual; correlation-staged malformed `state_proposal`; the ordered total disposition table and BOOTSTRAP leg; the two-form trusted-tuple assign gate; broker-no-durable-state; all-four-operation F64 fencing with durable settlement for the three effectful F59 calls; and F60/F66 seat-credential custody. A server that merely listens, reads the initial handshake, and leaves those semantics absent could satisfy the named CT-C09 and 62/2 census. That is insufficient E2 proof for the plan's own "built to spec stack" claim.

## Required bounded successor

Reissue the plan parented to this review with only the following C09 corrections:

1. Explicitly authorize the coupled CT-C09 composition strengthening: launch the NEW production broker binary through the production supervision/process boundary, wait for an observable READY condition, and only then run the existing real `brokerclient.Establish` predicate. Preserve the real socket, real app applier/control-generation advance, and `err == nil && session != nil` assertion. No fake listener, in-client server, in-process substitute for the production binary, exit-zero-only readiness, or predicate weakening.
2. Name the production server test matrix that proves every m-7 agenda item 4–8/10/11 at E2, including fail-closed negative legs and the study-governs-r11 exclusions. Tie those tests to the named C09 preparatory commits; the final C09 commit still turns CT-C09 green. The checklist may be realized in new internal server tests plus the strengthened seam composition test, but review prose alone is not the proof.
3. State the server launch owner/locus and READY observable used both by the real app composition and CT-C09, so the test and production path consume the same launcher rather than parallel wiring.

Preserve every other substantive r16 decision unless a correction above requires a named mechanical consequence. A12, Starter/P5, C07 with its already-authorized READY strengthening, the separate-process/no-conductor/no-durable-state laws, exact 62/2 close census, row-per-commit structure, P4/P5 inversions, dependency fence, push-after-each-commit directive, draft-only PR fence, registrations, remaining D01/D02 ownership, and no-E3/no-merge/no-ready boundaries are accepted and not reopened.

Boundary-contract disposition: A12, Starter, and C07 have named writers, consumers, and E2 acceptance. C09 has the named client/server contract and owner, but its launch writer is not connected to the test consumer and its deeper server semantics lack executable reader evidence; F1/F2 are the exact missing boundary proof.

Tests / verification: E2 focused run `go test -tags seam -count=1 ./test/seam/ -run '^TestCT_(A12|C07|C09)$'` failed all three for the expected current reasons: A12 unsupported `relay.read`; C07 `run start is unavailable`; C09 absent broker socket. Static production census found five existing binaries, no broker binary, no non-test broker listener/launcher, no `relaytool.New` caller, and the P5 literal nil Starter at `cmd/frank-app/main.go:24`. Exact r16 lint and digest, INDEX parentage, governing design hashes, and the clean linked-worktree/local-remote-equal tip were verified. No production or seam-test byte was edited.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — C09 includes F60/F66 credential custody and capability fencing; review-only here
- migration/backfill/destructive-write/canonical-data-repair: no
- money/inventory/orders/planning/accounting/trust-critical-state: yes — the governed broker and tool fence are trust-critical
- AI-or-automation-acts-downstream: yes — the worker and broker authorize tool operations
- worker/scheduler/queue/retry/async-side-effect: yes — supervised processes, broker lifecycle, and durable settlement
- cross-repo/service-contract/generated-schema/shared-API-event: yes — m-7/m-10/m-9 shared IPC and operation contract
- user-visible-control-with-materializer/downstream-consumer: no
- test-runtime-role-mismatch: yes — the planned separate production server has no authorized launch path in CT-C09
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — dispatch would force an unplanned launch/test choice and permit under-proven server semantics
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade or risk acceptance requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

Remaining gates: no SCOPE_DIFF or implementation token may issue on this verdict. A corrected successor requires review approval before the separately addressed implementation relay. Merge and ready-for-review remain operator-gated at WP5.

ACTIONS_GIT_REF: engine-lane governance act only — this review is drafted under `frank/.relays/s16a/.engine/drafts/s16a.implementer/` for daemon submission; one prior-turn implementer-owned staging artifact was removed before the final status capture; no source, test, implementation-worktree, branch, commit, push, PR, merge, CI/CD, provider, credential, store, or runtime state changed
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s16a/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s16a/s16a-build/PLAN-planner-20260825-084308.md
?? frank/.relays/s16a/s16a-pr/SITREP-planner-20260825-072634.md
?? frank/.relays/s16a/s16a-wp2/SITREP-implementer-20260825-083922.md
?? frank/.relays/s16a/s16a-wp2/SITREP-planner-20260825-084146.md
