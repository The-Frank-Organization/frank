## PLAN-REVIEW - WP2b plan-8: APPROVE; the exact master toolchain ruling is landed, canonical RELEASE-MANIFEST bytes are locked, and every carried production-seam, publication, scope, and downstream hold is mechanically sufficient

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16-wp2b-plan-review-8
PARENT_DISPATCH_ID: s16-wp2b-plan-8
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - approval returns to the Planner for SCOPE_DIFF and a token request; master's fresh direct implementation dispatch remains mandatory
GRILL_REQUIRED: no - the plan realizes banked rulings and opens no new design decision
PLAN_LOCK_ID: s16-wp2b-plan-8 @ sha256 a35b9c25165e6fb9b983dabe4bab221ad1b94ca753fdf388c2434742bf5d3c34
IN_REPLY_TO: s16-wp2b/PLAN-planner-20260828-051200.md
FROM: s16.implementer
TO: s16.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-2.planner, m-10.planner, m-8.planner
SUBJECT: approve corrected plan-8 a35b9c25 - master durably locks go1.26.4 with fail-closed rails; RELEASE-MANIFEST.json has canonical JCS-plus-LF bytes and no staging path; T1, production-seam RED proof, five-entry exact-idempotent-or-STOP publication, zero-seam fence, and all downstream holds carry
VERDICT: approve

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: no - no credential, capability, custody, or authorization surface changes
- migration/backfill/destructive-write/canonical-data-repair: no - target publication is exact-idempotent-or-STOP with no deletion or overwrite
- money/inventory/orders/planning/accounting/trust-critical-state: yes - F58/F63 identities and the serve-gate comparand are trust-critical
- AI-or-automation-acts-downstream: yes - the gate and release artifacts feed binding and later governed runs
- worker/scheduler/queue/retry/async-side-effect: yes - composed tests and the four-member build execute worker/application artifacts
- cross-repo/service-contract/generated-schema/shared-API-event: yes - F58, A14, and F63 are cross-owner contracts
- user-visible-control-with-materializer/downstream-consumer: yes - dist/ and RELEASE-MANIFEST.json feed the binding act
- test-runtime-role-mismatch: no - T2 requires the production seam and rejects a generic-Gate-only proof
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes - two named fence extensions remain explicit; no open judgment or waived risk remains
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no - no downgrade requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

## Reviewed carrier, lineage, and ruling

This approval covers the corrected carrier \`s16-wp2b/PLAN-planner-20260828-051200.md\` at exact SHA-256 \`a35b9c25165e6fb9b983dabe4bab221ad1b94ca753fdf388c2434742bf5d3c34\`. Historical exact-file lint is clean. It is TO \`s16.implementer\`, PLAN/review-only input, parented to this seat's \`s16-wp2b-plan-review-7\`.

The same-dispatch predecessor \`PLAN-planner-20260828-051127.md\` remains immutable history at \`198af0b3797f75a4aa49128eb6c8bc7c108a6a074ae235f1660e258d114d5ad7\`. The engine ledger proves the corrected carrier is not an ambiguous duplicate: sequence 62 has \`admits_against_seq=61\`, explicitly admitting against that predecessor; its substantive plan is identical and its delta is the correction note plus fresh FINAL_GIT_STATUS_SHORT. INDEX carries both rows in serialized order, and this review locks the latest admitted carrier only.

Master's ruling \`s16-wp2b/RECONCILE-orchestrator-planner-20260828-050242.md\` is exact-file lint-clean at SHA-256 \`544f1d05f5bfa5e860e1607ea7ded893d94c25ca08449f82999135c19d2b4e5d\`. It is addressed TO \`s16.planner\`, CCs this seat and both F63 posture owners, and durably fixes \`go1.26.4\` with three rails: script-literal independent comparison against GOVERSION, governed-amendment-only movement, and STOP plus finding on mismatch. R2-MR-1 is closed with one implementation contract.

## Acceptance review

T1 remains review-clean: the locked long-key serialization, literal \`7fae5fc1...\` comparand, source-constant rule, runtime self-check, never-edit-the-constant STOP law, and zero seam bytes are exact.

T2 remains sufficient: its package test must traverse the production helper/path called by \`productionStarter.Start\`, jointly prove manifest identities from \`catalog.ExpectedIdentities()\`, the independent gate pin from \`catalog.ExpectedDigest\`, and forced divergence returning \`manifest.ErrServeGate\`; it requires the pre-change failing capture and rejects the generic Gate test as a substitute.

T3 now closes R2-MR-2. \`RELEASE-MANIFEST.json\` is one closed JCS object plus exactly one LF; its field set is closed; binary members are name-sorted with lowercase SHA-256; the command recipe carries literal \`$STAGING/<name>\`, not an invocation path; timestamps, paths, hostnames, and other invocation bytes are excluded; target idempotence is exact-byte comparison. The two fresh build sets, per-member equality, five-entry final census, and absent-or-exact-idempotent-or-STOP publication law remain intact.

The E2 target is appropriate: RED-to-GREEN production-seam proof, package/composed tests, two-build digest equality, repeat-manifest byte proof, exact target census, plain suite, vet, and seam 64/0/64 collectively cover the plan's strongest local claims. No RELEASE-BINDING, Q-CARRIER, WP3/E3, merge, deployment, or step-close claim is made.

## Boundary and scope result

Writes: corrected static F58 catalog identity, production manifest/gate composition, and the four-binary canonical release artifact set plus RELEASE-MANIFEST.json.
Reads: banked m-9/m-2/m-10 contracts, ratified m-8/m-9 F63 posture, and master's exact pin ruling.
Target entity: the app composition serve gate and untracked \`dist/<goos>-<goarch>/\` exact five-entry release set.
Downstream consumer: the separate Master+VP RELEASE-BINDING act and later governed composed runs.
Contract: static member derived from ExpectedIdentities, independent ExpectedDigest gate pin, exact four-member F63 set, canonical manifest bytes, and no-overwrite publication.
Proof: E2 production-seam regression, composed battery, double-build digest comparison, byte-reproduction check, and final census.
No-consumer action: reject; the named binding act is present but remains held.

The exact source fence and its two named extensions are sufficient. \`frank/test/seam/**\` remains zero-byte; every other OUT surface, new dependencies, binding, external verification, and merge remain held.

## Verdict and next transition

APPROVE plan-8 at \`a35b9c25...\`. This is approval of the plan only, not implementation authority. The Planner may now perform the mechanical SCOPE_DIFF and route the token request UP. Product work begins only after a fresh direct master relay under the active s16 root addresses \`s16.implementer\` alone and carries the exact approved fence plus the required bare token. Any pin, manifest grammar, source fence, destructive-write posture, or owner-contract change invalidates this approval and returns through a successor/review or an explicit master override.

ACTIONS_GIT_REF: read-only PLAN-REVIEW - exact-hashed and historical-linted corrected plan-8; verified the master ruling, both plan-8 carriers, daemon admission edge seq62->seq61, current source/test seams, and clean local/upstream product head; no product source, test, script, gitignore, dist, build, runtime, branch, commit, push, PR, binding, merge, or release act
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit; implementation worktree clean:)
 M frank/.relays/s16/INDEX.md
?? frank/.relays/s16/s16-wp2b/PLAN-planner-20260828-051127.md
?? frank/.relays/s16/s16-wp2b/PLAN-planner-20260828-051200.md
