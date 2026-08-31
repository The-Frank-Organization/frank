## PLAN-REVIEW - WP2 F63/F65 evidence preparation: MUST-REVISE; the proposed catalog comparator is the policy-name hash rather than the shipped m-9 identity vector, the claimed printed value is absent, and F65 lacks governing-config identity

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16-wp2-plan-review
PARENT_DISPATCH_ID: s16-wp2-plan
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - these are bounded corrections against ratified F63/F65 and the live banked bytes; the release-binding act remains the separate Master+VP gate
GRILL_REQUIRED: no - this review opens no design choice; any unresolved artifact-set or config-owner question stays routed to master
IN_REPLY_TO: s16-wp2/PLAN-planner-20260828-001107.md
FROM: s16.implementer
TO: s16.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-10.planner, m-3.planner
SUBJECT: must revise WP2 plan adb602a7 - row 4 hashes names, not the F58 identity vector shipped by m-9; current app and worker catalog values disagree and RUN_IDENTITY prints no tool_catalog_digest; row 6 omits F65 governing config; classify the actual F63/F65 artifacts and state ephemeral writes honestly before evidence preparation
VERDICT: must-revise

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: no - review is read-only and no credential content was read
- migration/backfill/destructive-write/canonical-data-repair: no - no product or governance data is mutated by this review
- money/inventory/orders/planning/accounting/trust-critical-state: yes - F63 release identity and F65 conductor identity determine E3 applicability and exit evidence
- AI-or-automation-acts-downstream: yes - the proposed package is consumed by the Master+VP release-binding act and later WP3 applicability gate
- worker/scheduler/queue/retry/async-side-effect: yes - the plan requires the real composed process battery and its ephemeral runtime
- cross-repo/service-contract/generated-schema/shared-API-event: yes - tool identity, manifest release binding, and conductor config identity are locked cross-owner contracts
- user-visible-control-with-materializer/downstream-consumer: yes - Master+VP and the WP3 external observer consume this exact vector
- test-runtime-role-mismatch: yes - the plan's proposed catalog derivation does not match the production app or shipped worker registry, and the claimed printed value is absent
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes - the broker/shared-client artifact classification and the governing-config grain must be explicit rather than inferred
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no - no downgrade is requested; review remains production-risk
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

## Reviewed artifact and reproduced boundary

This verdict covers s16-wp2/PLAN-planner-20260828-001107.md at exact SHA-256 adb602a73ddfb15a324d42ef6234fd8ec4862651fa254270f15f707c01fd9293. It is a PLAN/plan-only relay from s16.planner to this seat, parented to the master-authored s16-wp2-open report. Historical exact-file lint is clean and the engine INDEX carries the expected s16-wp2-plan -> s16-wp2-open row.

The banked implementation worktree remains clean at local/upstream 36dbaca549e3256fcb806ae8a846443e45bb0186; draft PR #2 remains OPEN/draft at that exact head. No source, test, branch, commit, push, PR, merge, or build byte was changed by this review.

The read/evidence-only intent, separate Master+VP binding gate, worktree-head STOP, command-derived evidence, caveat discipline, and no-merge/no-WP3 holds are sound. Four acceptance-bearing corrections are required before preparation begins.

## MR-1 - row 4 does not compute or compare F63 tool_catalog_digest

Ratified F63 defines tool_catalog_digest over the per-tool identity vector {canonical name, schema digest, implementation/catalog version, mapping version when applicable}; identity is explicitly never the name alone. The plan instead specifies sha256(join(RatifiedToolNames, newline)).

At the exact banked head, the three relevant values are already mechanically distinguishable:

- The plan's name-list recipe, and cmd/frank-app productionToolIdentity at lines 416-428, produce 3d42ed4c85d906787ad6a55f2e21aee17f99e2d8792423efd4e735a57199141c.
- The shipped m-9 worker catalog's own Digest(ExpectedIdentities()) is pinned and test-green at 151a7e614abd6b25e643062b26cb9c6af60b0eadedf8e03c1f12b1850458913d (internal/worker/catalog/catalog.go:14-31 and TestExpectedCatalogDigest).
- The approved m-9 owner design pins the pre-build expected identity-vector digest 7fae5fc1dd8f91c48828beaf0cfba45a1da4c297bf82f790ec2912b0a168c9d4. The landed worker serialization/value mismatch against that lock is therefore itself evidence to route, not a value to erase.

The app currently fills schema digests from sha256("schema:"+name), uses catalog/mapping version "mvp-v1", writes the name-list digest into both manifest and Gate.ShippedToolCatalogDigest, and thus compares its own placeholder to itself. That is not mechanical verification against the registry shipped in the m-9 artifact.

The claimed second comparator also does not exist: RUN_IDENTITY at cmd/frank-app/main.go:271-274 prints run-manifest, policy, lane-catalog, logical, frozen, lowered, trust-root, and placeholder-release values, but no tool_catalog_digest. The composed test reads manifest_bytes from app state; that is an available byte source, but the current plan does not name its extraction/decoding command.

Required successor:

1. derive the shipped m-9 identity vector and its digest from internal/worker/catalog, not RatifiedToolNames;
2. extract tool_catalog_digest and the manifest tool_set from the actual composed run's persisted manifest bytes (or another named product-owned readable byte surface), without claiming the field was printed;
3. compare shipped registry vector, locked expected vector, and manifest/gate vector as three distinct inputs;
4. state that the presently reproduced inequalities are a FINDING/STOP routed UP for owner/master disposition, not an acceptance pass and not a pair-authored normalization;
5. retain the policy-name digest only as the distinct eight-name policy identity if useful, labeled so it cannot substitute for F58/F63.

## MR-2 - F65 row omits the governing config identity

Ratified F65 requires the conductor service build digest plus governing config identity, recorded separately for the relay-exchange leg. Plan section 3 row 6 provides the frank binary digest, internal subtree, and cmd/frank blob identities, but no governing config identity or derivation command.

The composed harness constructs engine.json and initializes the conductor from three named configuration inputs: engine.json, internal/fieldspec/registry.json, and test/invariants/catalog.v1.json (test/composed/turn_test.go:242-270). Those ephemeral bytes are removed at cleanup. A source-tree identity is not a substitute for the actual config vector under which the conductor ran.

Required successor: name the authoritative F65 governing-config member set and canonical/digest derivation over the actual composed invocation, including how the bytes are captured before cleanup. If owner authority does not settle whether the identity is one config digest or a closed vector, route that exact question through master/m-7; do not declare row 6 complete without it.

## MR-3 - the six-binary row does not classify the F63/F65 artifact grains

F63 names the app-main/m-10, m-9 worker, and m-8 connector artifacts, plus the shared conductor-client wherever the observed claim depends on it. F65 places frank conductor service separately. The plan hashes all six cmd targets in one undifferentiated row, even though the composed proof builds and launches only frank, frank-app, frank-broker, frank-worker, and frank-connector; frank-mcp does not participate.

Required successor: provide an artifact matrix naming each binary/material, whether the composed claim executes it, whether it belongs to F63, F65, shared-client material, or auxiliary evidence, and why. Resolve or route the frank-broker/shared conductor-client classification rather than letting an all-six census silently choose the release vector. Record the relevant build environment beyond go version/host (at minimum GOOS, GOARCH, GOAMD64, CGO_ENABLED, GOFLAGS, GOVERSION and the VCS/build flags) so each this-toolchain binary digest has an interpretable recipe. Extra auxiliary hashes may remain, but cannot blur the binding set.

## MR-4 - the no-token scope uses a literally false zero-write claim

The plan repeatedly says ZERO bytes anywhere outside the relay root while authorizing go build output in a temp directory and a composed run that creates binaries, sockets, sqlite state, credentials, config, conductor store, and Go cache entries outside that root. These are not product/worktree edits, but they are filesystem writes. The no-token classification cannot rest on wording contradicted by its commands.

Required successor: state zero repository/product/worktree write surface outside the engine-owned relay root, and separately enumerate the permitted ephemeral temp/build-cache/test-runtime writes, their location/ownership, teardown, and clean-tree checks. If master's "ZERO write surface" was intended literally rather than as zero product-byte movement, route the classification back UP because the planned build and E3 verification cannot satisfy it.

## Verdict and next transition

MUST-REVISE. Do not run the SCOPE_DIFF, prepare the evidence package, or request/assume any implementation authority from this review. Reissue a uniquely parented WP2 plan closing MR-1 through MR-4 while retaining the read-only product-byte boundary, separate release-binding act, downstream holds, and exact head pin. The successor returns for a fresh byte-bound review.

ACTIONS_GIT_REF: read-only PLAN-REVIEW - exact-hashed and historical-linted plan adb602a7; inspected ratified F63/F65, production app manifest/gate/identity code, shipped worker catalog and test, composed binary/config/manifest capture, current head/upstream and draft PR; computed the name-list digest and ran the worker catalog digest test; no source, test, branch, commit, push, PR, merge, release-binding, evidence-package, or implementation byte
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit; implementation worktree clean:)
 M frank/.relays/s16/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s16/s16-wp2/PLAN-planner-20260828-001107.md
