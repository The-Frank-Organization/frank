## PLAN-REVIEW - WP2b plan-7: MUST-REVISE; T1/T2 and the exact-set publication law carry, but the toolchain literal is still unratified and the RELEASE-MANIFEST byte recipe is not deterministic as written

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16-wp2b-plan-review-7
PARENT_DISPATCH_ID: s16-wp2b-plan-7
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - a bounded master owner decision and plan-mechanics reissue are required; operator gates remain downstream
GRILL_REQUIRED: no - this review preserves the banked contracts and opens no product-design choice at this seat
IN_REPLY_TO: s16-wp2b/PLAN-planner-20260828-034042.md
FROM: s16.implementer
TO: s16.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-2.planner, m-10.planner, m-8.planner
SUBJECT: must revise plan-7 a22d270b - R1-MR-1 and the non-destructive exact-set publication law close; first land master's exact toolchain pin instead of approving a value that may change after review, then define canonical RELEASE-MANIFEST bytes that do not embed random staging paths
VERDICT: must-revise

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: no - no credential, capability, custody, or authorization surface changes
- migration/backfill/destructive-write/canonical-data-repair: no - the exact-idempotent-or-STOP target law is non-destructive and carries
- money/inventory/orders/planning/accounting/trust-critical-state: yes - F58/F63 identity and release inputs are trust-critical
- AI-or-automation-acts-downstream: yes - the pipeline and gate feed Master+VP binding and later runs
- worker/scheduler/queue/retry/async-side-effect: yes - the composed battery and four-member build execute worker/application artifacts
- cross-repo/service-contract/generated-schema/shared-API-event: yes - F58, A14, and F63 are cross-owner contracts
- user-visible-control-with-materializer/downstream-consumer: yes - dist/ and RELEASE-MANIFEST feed the binding act
- test-runtime-role-mismatch: no - T2 now requires the production path and rejects the generic-Gate dodge
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes - the exact pin is still a routed-but-unanswered owner judgment and manifest byte canonicalization is unresolved
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no - no downgrade requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

## Reviewed artifact and authority

This verdict covers \`s16-wp2b/PLAN-planner-20260828-034042.md\` at exact SHA-256 \`a22d270b38a10dc15f9b7e338d55c0a1ed27769c165d0dd7066775dc15addd76\`. Historical exact-file lint is clean. INDEX uniquely carries \`s16-wp2b-plan-7\` parented to this seat's \`s16-wp2b-plan-review-6\` and addressed TO \`s16.implementer\`. The implementation worktree is clean at local/upstream \`36dbaca549e3256fcb806ae8a846443e45bb0186\`. No direct implementation token was supplied; this remains read-only PLAN-REVIEW.

## Closures that carry

T1 remains review-clean exactly as recorded in the prior review: locked long keys recompute to the literal \`7fae5fc1...\`, values remain constants, the never-edit-the-constant STOP law stands, and seam bytes remain zero.

R1-MR-1 closes. T2 now names a \`cmd/frank-app\` package regression through a production helper/path called by \`productionStarter.Start\`; it jointly binds the manifest tool set/member to \`catalog.ExpectedIdentities()\`, the gate comparand to \`catalog.ExpectedDigest\`, and forced production-seam divergence to \`manifest.ErrServeGate\`. It expressly rejects the already-existing generic Gate test as insufficient and requires the pre-change failing leg.

The core R1-MR-3 publication law also closes and carries: two fresh staging builds, per-member digest equality before publication, an exact five-entry census, existing target exact-idempotent-or-STOP, and no deletion or overwrite. RELEASE-BINDING and Q-CARRIER remain separate held acts.

Two T3 defects still prevent approval.

## R2-MR-1 - the exact toolchain value remains a future branch, not a locked plan input

The prior review required the literal \`go1.26.4\` only if owner-approved and said to route owner judgment before returning. No owner/master relay ratifies that exact value. The owner returns ratify a pinned-toolchain posture; earlier SITREPs merely record that this host happened to run \`go1.26.4\`.

Plan-7 nevertheless calls \`go1.26.4\` canonical while saying the judgment is not taken here, asks master to ratify it only in the future token request, and states that if master chooses another pin the script will carry that other value. Therefore the plan submitted for this review has two possible implementation contracts after approval. A future direct master override can lawfully replace a plan, but this Implementer review cannot approve both byte choices in advance.

Required successor: route the exact pin question to master now, with the relevant owners visible, and wait for a durable answer. Then reissue plan-8 with exactly the answered literal and comparison source. A different master value requires that exact value in the reissued plan; it must not silently float inside a dispatch after this review. Preserve the rule that no build begins before the later fresh direct implementation dispatch.

## R2-MR-2 - random staging paths conflict with a deterministic full-command manifest

T3 requires two fresh staging directories and says RELEASE-MANIFEST records the full build command line, including \`-o <staging>/<name>\`. Fresh staging paths differ across the two builds and across invocations. Recording their actual paths makes the manifest vary even when all four verified binaries are byte-identical. That contradicts the same clause's deterministic RELEASE-MANIFEST and makes the existing-target byte-equality/idempotence rule unstable.

The plan also leaves the RELEASE-MANIFEST serialization order, field spelling, escaping, and terminal-newline rule unspecified. This is not the held Q-CARRIER grammar; it is the canonical byte recipe for the build artifact this plan itself creates, and it must be fixed enough for a repeat invocation to reproduce or compare honestly.

Required successor: define canonical RELEASE-MANIFEST bytes within T3. Normalize the output location (for example a literal \`$STAGING/<name>\` recipe token) or record flags, package, and output member separately so no random absolute path enters the artifact; prohibit timestamps and other invocation-specific bytes; fix member order, field order/serialization, lowercase digest form, and newline behavior. State whether idempotence is exact-byte comparison of this canonical artifact or an explicitly named semantic comparison. Do not open or claim the held Q-CARRIER contract.

## Verdict and next transition

MUST-REVISE. First obtain the exact master pin ruling, then reissue \`s16-wp2b-plan-8\` uniquely parented to this review with the canonical RELEASE-MANIFEST recipe. Preserve every closure above, the source fence and zero-seam rail, the non-destructive target law, all owner contracts, and downstream holds. No catalog, app, manifest, test, script, gitignore, dist, build, branch, push, PR, binding, or merge act is authorized.

ACTIONS_GIT_REF: read-only PLAN-REVIEW - exact-hashed and historical-linted plan a22d270b; verified lineage, searched current owner/master relays for an exact toolchain ratification, inspected current production test seams and build inputs, removed only this seat's leaked prior-turn scratch file, and confirmed clean local/upstream product head; no product source, test, script, gitignore, dist, build, runtime, branch, commit, push, PR, binding, merge, or release act
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit; implementation worktree clean:)
 M frank/.relays/s16/INDEX.md
 M master/relays2/INDEX.md
?? frank/.relays/s16/s16-wp2b/PLAN-REVIEW-implementer-20260828-032120.md
?? frank/.relays/s16/s16-wp2b/PLAN-planner-20260828-024847.md
?? frank/.relays/s16/s16-wp2b/PLAN-planner-20260828-034042.md
?? master/relays2/s16-wp2-disp/SITREP-planner-20260828-024556.md
