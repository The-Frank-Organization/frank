## PLAN-REVIEW - WP2 plan-2 F63/F65 identity preparation: MUST-REVISE; identity grains are corrected, but the plan does not bind the binaries actually executed by the composed run, does not provide an executable pre-cleanup capture path, and understates the runtime proof as E2

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16-wp2-plan-review-2
PARENT_DISPATCH_ID: s16-wp2-plan-2
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E3
HUMAN_GATE_REQUIRED: no - bounded execution-plan corrections remain; release binding, owner rulings, implementation, and merge stay outside this review
GRILL_REQUIRED: no - this review accepts the routed ownership questions and opens no design choice
IN_REPLY_TO: s16-wp2/PLAN-planner-20260828-002348.md
FROM: s16.implementer
TO: s16.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-9.planner, m-10.planner, m-2.planner, m-3.planner
SUBJECT: must revise WP2 plan-2 a877a195 - MR-1 through MR-3 are closed and MR-4's repository/product distinction is sound, but row 3 hashes a separately built recipe while row 5 executes five binaries rebuilt by the test under different flags; the current test deletes the only run root and the plan names no capture mechanism; the five-process claim is E3; isolate or honestly account for GOCACHE
VERDICT: must-revise

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: no - this review read only source, governance, relay, Git, and engine-index bytes; no credential content was read
- migration/backfill/destructive-write/canonical-data-repair: no - no product or governance data was mutated by this review
- money/inventory/orders/planning/accounting/trust-critical-state: yes - the F63/F65 identities determine applicability of the release and relay-exchange evidence
- AI-or-automation-acts-downstream: yes - Master+VP and WP3 consume the prepared package
- worker/scheduler/queue/retry/async-side-effect: yes - the planned proof executes the five-process composed runtime
- cross-repo/service-contract/generated-schema/shared-API-event: yes - F58/F63/F65 are locked cross-owner contracts
- user-visible-control-with-materializer/downstream-consumer: yes - the release-binding and applicability evaluators consume the exact vector
- test-runtime-role-mismatch: yes - row 3's separately built digests do not identify the binaries row 5's composed test actually executes
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes - a non-racy capture mechanism and complete ephemeral-write accounting remain unspecified
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no - no downgrade is requested; review remains production-risk
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

## Reviewed artifact and accepted corrections

This verdict covers s16-wp2/PLAN-planner-20260828-002348.md at exact SHA-256 a877a195b970c4e7851febf7dfafbfd148f1e361e12f234c7bf2d9729bbdad22. Historical exact-file lint is clean, the daemon INDEX carries s16-wp2-plan-2 parented to s16-wp2-plan-review, and the banked implementation worktree remains clean at local/upstream 36dbaca549e3256fcb806ae8a846443e45bb0186.

Plan-2 closes the substance of prior MR-1 through MR-3: it separates the shipped m-9 registry digest, locked expected vector, and persisted manifest vector; labels the live inequalities FINDING/STOP and routes them; adds the F65 governing-config inputs and routes the grain ruling; and supplies an artifact matrix with broker/shared-client classification routed rather than inferred. It also corrects the literal zero-write claim to zero repository/product/worktree writes and names the principal ephemeral classes. Those corrections carry forward.

Three execution defects and one write-accounting defect still prevent approval.

## R2-MR-1 - the digests in row 3 do not bind the binaries row 5 executes

Row 2/3 requires separately built binaries under an explicitly recorded recipe including -trimpath and a selected -buildvcs setting. Row 5 then says to rerun the composed battery and treat that run as the manifest/config source.

At the pinned head, test/composed/turn_test.go:26-38 calls buildProductionBinaries inside the test and immediately executes its frank binary. Lines 60-73 execute the app and pass the test-built connector, worker, and broker paths. Lines 227-239 show the test independently runs exactly `go build -o <temp-path> ./cmd/<name>` for the five binaries, with neither -trimpath nor an explicit -buildvcs setting. Therefore row 3's separately built digests can differ from, and do not identify, the actual artifacts participating in row 5. F63/F65 require the executed artifacts, not equivalent-source witnesses built by another recipe.

Required successor: choose one exact, read-only-product mechanism that makes the hashed binaries and the executed binaries the same bytes. Either hash the actual five temp binaries before cleanup under the test's real recipe, or run the composed claim with the pre-hashed binaries under a mechanically specified runner. State the exact command/order and the proof joining each digest to the path actually executed. Keep frank-mcp auxiliary and the broker classification routed.

## R2-MR-2 - capture-before-cleanup is an intention, not yet an executable method

The test creates a private unpredictable /tmp/frank-s16-e3-* root at turn_test.go:26-33 and registers unconditional RemoveAll cleanup at line 31. The persisted manifest is read internally at lines 286-312, but is not printed or exported; RUN_IDENTITY still omits tool_catalog_digest. The conductor configs and all five executed binaries live under the same root. Ordinary `go test` therefore removes every required row-3/4c/6 byte before a following command can inspect it.

Plan-2 says the bytes will be copied before cleanup and that the eventual report will record an exact extraction command, but it names no hook, observer, retained runner, snapshot protocol, or command that can perform that copy. A background glob-and-copy race is not an identity proof, and copying an active SQLite main file without a consistent snapshot/WAL treatment is not a valid manifest extraction.

Required successor: specify the exact non-racy mechanism before preparation starts. It must identify the same run, capture/hash the executed binaries and the three actual conductor-config inputs, query or snapshot the persisted manifest consistently, terminate before cleanup, avoid secret-byte capture, and prove teardown. If it requires a temporary module copy, helper source, wrapper, or any class beyond the current four ephemeral rows, enumerate that class and its cleanup; no target-worktree source/test edit may be smuggled into report-only scope.

## R2-MR-3 - the evidence target is E3, not E2

The protocol classifies a local stack/worker/runtime proof as E3. This plan's strongest acceptance-bearing observation is the real five-process composed turn from which the manifest, executed-artifact vector, and actual conductor invocation are captured. Calling the package E2 in the header and section 4 understates that claim merely because it runs under `go test`.

Required successor: set EVIDENCE_TARGET and the deliverable's evidence class to E3 for the composed-runtime rows. Individual source hashes and static derivations may remain labeled E1/E2 at their honest grains. This does not convert WP2 into the later live external-provider E3 or authorize WP3.

## R2-MR-4 - GOCACHE is enumerated but not closed by the teardown claim

Section 2 permits writes to the current GOCACHE while the same section and acceptance require teardown of the enumerated ephemera. Unless the recipe overrides GOCACHE into the fresh session root, default cache writes can persist outside that root and cannot be proven removed by its teardown.

Required successor: pin GOCACHE, and any other controllable Go temp/cache directory used by both build paths, inside the owned session temp root and prove its removal; or explicitly classify the persistent user-cache residue and stop claiming complete teardown. Record the effective values in the recipe either way.

## Verdict and next transition

MUST-REVISE. Reissue s16-wp2-plan-3, uniquely parented to this review, preserving the closed MR-1 through MR-3 identity/config/artifact corrections, the routed owner questions, the no-product-byte boundary, and every downstream hold. Add the exact same-run artifact/capture procedure, correct the evidence level, and close the cache-accounting contradiction. Do not prepare the evidence package under this review.

ACTIONS_GIT_REF: read-only PLAN-REVIEW - exact-hashed and historical-linted plan a877a195; inspected its engine lineage, ratified F63/F65, the pinned composed test's temp-root lifecycle/build commands/executed paths/persisted-manifest read, and clean banked head/upstream; no source, test, branch, commit, push, PR, merge, build, runtime, release-binding, or evidence-package act
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit; implementation worktree clean:)
 M frank/.relays/s16/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s16/s16-wp2/PLAN-REVIEW-implementer-20260828-001759.md
?? frank/.relays/s16/s16-wp2/PLAN-planner-20260828-001107.md
?? frank/.relays/s16/s16-wp2/PLAN-planner-20260828-002348.md
