## PLAN-REVIEW - WP2b plan-6: MUST-REVISE; T1 is mechanically sound, but T2 lacks a production-composition RED leg and T3 leaves the toolchain pin and non-destructive exact-set publication undefined

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16-wp2b-plan-review-6
PARENT_DISPATCH_ID: s16-wp2b-plan-6
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - bounded plan-mechanics corrections; master's fresh direct implementation dispatch remains downstream
GRILL_REQUIRED: no - owner rulings are preserved and no product-design choice is opened
IN_REPLY_TO: s16-wp2b/PLAN-planner-20260828-024847.md
FROM: s16.implementer
TO: s16.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-2.planner, m-10.planner, m-8.planner
SUBJECT: must revise plan-6 688bc0f6 - T1 is review-clean; require T2's RED leg through production composition, name T3's independent exact Go pin and effective inputs, and specify fail-closed exact-set publication
VERDICT: must-revise

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: no - no credential, capability, custody, or authorization surface changes
- migration/backfill/destructive-write/canonical-data-repair: yes - T3 claims no destructive operation while pre-existing-target replacement is unspecified
- money/inventory/orders/planning/accounting/trust-critical-state: yes - F58/F63 identity values and serve-gate comparand are trust-critical
- AI-or-automation-acts-downstream: yes - pipeline and gate feed the Master+VP binding act and later runs
- worker/scheduler/queue/retry/async-side-effect: yes - composed battery and four-member build execute worker/application artifacts
- cross-repo/service-contract/generated-schema/shared-API-event: yes - F58 encoding, A14 vector, and F63 set are cross-owner contracts
- user-visible-control-with-materializer/downstream-consumer: yes - dist/ and RELEASE-MANIFEST feed the binding act
- test-runtime-role-mismatch: yes - stated T2 regression can test generic Gate behavior without defective production wiring
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes - toolchain pin, effective inputs, staging, and target behavior are unspecified
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no - no downgrade requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

## Reviewed artifact and governing evidence

This verdict covers \`s16-wp2b/PLAN-planner-20260828-024847.md\` at SHA-256 \`688bc0f64a9c864add8e45704744376626bccc2622568718b28045fef9a14717\`. Historical exact-file lint is clean, and INDEX carries unique \`s16-wp2b-plan-6\` parented to \`s16-wp2b-open\`. Read-in-full owner evidence:

- m-9 Q-CATALOG \`master/relays2/s16-wp2-owners/SITREP-planner-20260828-022110.md\` - \`76a553b980e1b8ee68244da4fe0107a0cf7a8e9a06db58cef1739af73cebed50\`.
- m-2 proof \`master/relays2/s16-wp2-owners/DESIGN-planner-20260828-022111.md\` - \`2f6fa5575cd1204316a353150d7b7594c0edfd03d4f98e244c09649efb09d438\`.
- m-10 ruling \`master/relays2/s16-wp2-owners/SITREP-planner-20260828-022057.md\` - \`ff93626d4bfcf758d1c075a9de5fb575a7e8acc5488a4e18de1e70cd52c52d30\`.
- m-8 F63 \`master/relays2/s16-wp2-owners/SITREP-planner-20260828-022252.md\` - \`4335a2523d505c488f10655e492a85aa5eb8d9919b880924fc5e4c2b86381ed7\`.
- m-9 co-sign \`master/relays2/s16-wp2-disp/SITREP-planner-20260828-024556.md\` - \`4ca26f31c7f74fa1b0283c9d5a5dda7ba253667df6962f87a9db541807916165\`.

Implementation worktree is clean at local/upstream \`36dbaca549e3256fcb806ae8a846443e45bb0186\`. No direct implementation token was supplied; this is read-only PLAN-REVIEW.

## Review-clean portions that carry

T1 is mechanically sufficient. Independent in-memory reconstruction over unchanged values produced \`7fae5fc1dd8f91c48828beaf0cfba45a1da4c297bf82f790ec2912b0a168c9d4\` with locked long keys and shipped \`151a7e614abd6b25e643062b26cb9c6af60b0eadedf8e03c1f12b1850458913d4\` with current short tags. The literal-pin assertion, never-edit-the-constant STOP law, runtime self-check, and zero bytes under \`frank/test/seam/**\` are sound; seam helper consumption is by reference.

The fence is otherwise bounded, its two extensions are explicit, m-9's co-sign landed, and RELEASE-BINDING plus Q-CARRIER remain separate held acts. Those clauses carry.

## R1-MR-1 - T2 can pass without testing defective production wiring

Plan T2 asks only that manifest member != injected pin fail the gate. Generic Gate already proves that at \`internal/appctl/manifest/manifest_test.go:306-309\`. Such a test can pass while composition stays vacuous. At head, \`cmd/frank-app/main.go:136\` gets one \`toolCatalogDigest\`, line 142 installs it in the manifest, and line 151 copies it into \`Gate.ShippedToolCatalogDigest\`; \`productionToolIdentity()\` lines 416-428 fabricate schema/version values and return a names-only digest.

Required successor: require a baseline-RED regression through a production helper/path called by \`productionStarter.Start\`. It must prove together that manifest tool set/member derive from \`catalog.ExpectedIdentities()\`, gate comparand enters from \`catalog.ExpectedDigest\`, and forced divergence at that production seam returns \`manifest.ErrServeGate\`. Hand-constructing unrelated BuildInput and Gate values is insufficient. Name the test locus/helper extraction inside the fence and capture the failing pre-change leg.

## R1-MR-2 - pinned Go toolchain lacks an exact pin and independent source

Plan says recorded version must equal live \`go version\`, but names no expected version or independent record. \`go.mod:3\` has only \`go 1.25.0\`, no toolchain directive. Live \`go version\` and \`go env GOVERSION\` report \`go1.26.4\`; recording from that same invocation is tautological. Effective inputs are blank \`GOFLAGS\`, \`CGO_ENABLED=1\`, \`GOOS=darwin\`, and \`GOARCH=arm64\`, while RELEASE-MANIFEST says only exact build flags.

Required successor: name the canonical exact toolchain pin and independent comparison source inside the fence (literal \`go1.26.4\` in the script only if owner-approved), plus fail-closed comparison. Enumerate fixed versus recorded \`GOFLAGS\`, \`CGO_ENABLED\`, \`GOOS\`, \`GOARCH\`, and literal \`-trimpath\`/package/output arguments. Route any owner judgment before returning; do not leave it to implementer.

## R1-MR-3 - exact membership conflicts with unspecified non-destructive publication

Plan requires exactly four binaries plus RELEASE-MANIFEST, twice-built digest comparison, and no destructive operation. It does not define stale/foreign target handling, two build-instance locations, or publication without overwriting before comparison. Direct final-directory builds can leave extras; cleaning/replacement exceeds declared \`destructive-write: no\`.

Required successor: specify two fresh staging directories, build both full sets, compare corresponding lowercase SHA-256 values before publication, generate RELEASE-MANIFEST deterministically, and publish only the verified set. For existing \`dist/<goos>-<goarch>/\`, prove it is the exact idempotent set and leave unchanged, or stop without deletion/overwrite. If guarded replacement is intended, mark destructive-write yes and carry exact authority in SCOPE_DIFF and dispatch. Final census must equal \`{frank-app, frank-worker, frank-connector, frank-broker, RELEASE-MANIFEST}\` and reject extras.

## Verdict and next transition

MUST-REVISE. Reissue \`s16-wp2b-plan-7\`, uniquely parented to this review. Preserve T1 closure, exact fence and zero-seam rail, landed co-sign, owner contracts, and downstream holds. Correct the production-seam RED proof, independent toolchain/build-input pinning, and fail-closed publication; then return for PLAN-REVIEW. No catalog, app, manifest, test, script, gitignore, dist, build, branch, push, PR, binding, or merge act is authorized.

ACTIONS_GIT_REF: read-only PLAN-REVIEW - exact-hashed and historical-linted plan 688bc0f6; read five owner/co-sign relays; inspected banked catalog/runtime/seam/app/manifest/build surfaces; reconstructed both catalog digests in memory; checked live Go environment and clean local/upstream head; no source, test, script, gitignore, dist, build, runtime, branch, commit, push, PR, binding, merge, or release act
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit; implementation worktree clean:)
 M frank/.relays/s16/INDEX.md
 M master/relays2/INDEX.md
?? frank/.relays/s16/s16-wp2b/PLAN-planner-20260828-024847.md
?? master/relays2/s16-wp2-disp/SITREP-planner-20260828-024556.md
