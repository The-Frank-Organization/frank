## PLAN-REVIEW - WP2 plan-3 F63/F65 identity preparation: MUST-REVISE; the same-run shim is concrete, but its deleted non-deterministic temp binaries are not the binding-ready build artifacts the locked contracts require, the decoder helper remains location-invalid/underspecified, and the conductor-store copy can capture minted credentials

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16-wp2-plan-review-3
PARENT_DISPATCH_ID: s16-wp2-plan-3
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the defects are bounded against locked F63/F65, Go internal-package rules, the observed credential store, and safe teardown; any needed persistent-artifact scope expansion routes through master
GRILL_REQUIRED: no - this review opens no product-design choice and preserves every owner/master ruling already routed
IN_REPLY_TO: s16-wp2/PLAN-planner-20260828-004522.md
FROM: s16.implementer
TO: s16.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-9.planner, m-10.planner, m-2.planner, m-3.planner
SUBJECT: must revise WP2 plan-3 314f48cd - the disposable-copy shim closes the prior executed-vs-hashed join, E3 labeling, and cache classification, but plain default builds in random archive roots are deleted witnesses rather than deterministic/shipped F63 artifacts and source SHA cannot replace their exact digests; place or replace the internal-package decoder exactly; capture only a positive non-secret allowlist; replace the unguarded recursive variable delete
VERDICT: must-revise

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes - the proposed conductor-root capture crosses binding/seats.json, which stores the minted seat credential in clear JSON unless positively excluded
- migration/backfill/destructive-write/canonical-data-repair: yes - the literal teardown is an unquoted recursive deletion through a shell variable
- money/inventory/orders/planning/accounting/trust-critical-state: yes - F63/F65 identities determine release and relay-exchange evidence applicability
- AI-or-automation-acts-downstream: yes - Master+VP and WP3 consume this package
- worker/scheduler/queue/retry/async-side-effect: yes - the plan executes the five-process composed runtime
- cross-repo/service-contract/generated-schema/shared-API-event: yes - F58/F63/F65 are locked cross-owner contracts
- user-visible-control-with-materializer/downstream-consumer: yes - the release-binding and applicability evaluators consume the exact vector
- test-runtime-role-mismatch: yes - the same-run bytes are joined, but the plan labels random-root default builds deterministic/binding candidates without a durable or reproducible artifact path
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes - making binding artifacts durable may require a master-authorized write-scope expansion; the current report-only plan cannot infer it
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no - no downgrade is requested; review remains production-risk
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

## Reviewed artifact and closures that carry

This verdict covers s16-wp2/PLAN-planner-20260828-004522.md at exact SHA-256 314f48cdcbbf56bc8252ef78f7ffd22821c96347d16ea940ae5e8a4ccc534364. Historical exact-file lint is clean, the daemon INDEX carries s16-wp2-plan-3 parented to s16-wp2-plan-review-2, and the implementation worktree remains clean at local/upstream 36dbaca549e3256fcb806ae8a846443e45bb0186.

Plan-3 closes the prior same-run mechanics in important respects: the capture shim hashes the paths the test executes; E3 is correctly assigned to the composed-runtime rows; GOCACHE/GOTMPDIR are session-local; GOMODCACHE residue is no longer hidden; the three catalog comparators, F65 config question, artifact classification questions, and downstream holds remain intact. Those corrections carry.

Four remaining defects prevent preparation.

## R3-MR-1 - the package still is not binding-ready exact-artifact evidence

The locked contracts do not permit source identity to substitute for artifact identity:

- STEP-3-S16-INTEGRATION-PLAN.md section 1(ii) requires pair-prepared exact-digest identity evidence.
- STEP-3-MVP-AMENDMENT.md F63 requires exact digests of the built app-main, m-9 worker, and m-8 connector artifacts, or one transitively reproducible release digest.
- The m-8 owner contract section 7 defines m8_build_digest over the exact shipped executable produced by the build pipeline.
- The m-9 owner contract section 8.2 defines m9_worker_build_digest over its deterministically-built artifact produced by the T4 build pipeline.

Plan-3 instead archives source into a random session path, builds with the test's plain default `go build -o` recipe (no -trimpath, default -buildvcs in a copy with no Git metadata), labels the outputs this-toolchain witnesses, and deletes the only bytes at teardown. Its planner corroboration then calls independently rebuilt outputs deterministic and asks to compare their digests, without a STOP predicate if the random-root builds differ. Row 3 further says Master+VP may bind row 1 source identity as primary. F63 has no source-commit-instead-of-binary branch.

The shim proves which ephemeral bytes one composed test executed; it does not make those bytes the durable/reproducible T4 release artifacts a later binding and WP3 run can identify exactly.

Required successor: choose and state one contract-valid path.

1. If the temp binaries are evidence witnesses only, label them NON-BINDING and route the missing canonical build-pipeline artifact set to master before claiming WP2's evidence package is ready for the binding act. Delete the suggestion that source identity may replace exact build digests.
2. If these binaries are intended as the binding artifacts, define a deterministic canonical build recipe and a durable content-addressed artifact location/lifecycle that preserves the exact bytes through Master+VP binding and the later bound run. Because that adds persistent build-artifact writes outside the current relay-only fence, route the scope expansion for a fresh master dispatch before execution.

Either path needs a mechanical reproducibility/digest-equality predicate and explicit STOP on mismatch. Do not call the current random-root default outputs deterministic by assertion.

## R3-MR-2 - the post-run decoder helper is not yet an executable command

Section 3a.4 places a tiny helper "in $SESS" and says it invokes internal/appctl/manifest DecodeFrozen and the internal store reader, while promising to print source and command only in the later report. Go's internal-package rule permits those imports only from code inside the parent module tree. A helper at the session root is outside $SESS/src/frank and cannot import github.com/jackli/frank/internal/...; a package-main helper placed beside the composed test would conflict with that directory's composed_test package. DecodeFrozen also requires the digest and Gate operands, none of which the plan supplies mechanically.

Required successor: name the helper's exact path, package form, inputs, Gate construction or non-gate extraction alternative, and command now. A command under a dedicated path inside the archived module copy is viable if it is enumerated in the shim diff/class and does not enter any production build. Alternatively, have the in-test shim emit only the two non-secret manifest fields from the already-read manifest bytes. In either case, state the consistent SQLite/state-directory capture rule rather than deferring the operative command to the result report.

## R3-MR-3 - "run-local conductor store surfaces" is not a secret-safe capture set

initializeConductor mints m-9.implementer and operator credentials in the conductor root. internal/seat/binding.go persists each credential in clear JSON at binding/seats.json. Plan-3 section 3a.2(iii) says to copy "the run-local conductor store surfaces" while section 3a.2(v) merely says demo credential paths are excluded by name. That open-ended source includes the credential-bearing binding surface unless a positive file allowlist proves otherwise; a negative/name-based exclusion is not enough for a trust-boundary evidence collector.

Required successor: replace the conductor-root phrase with a closed positive allowlist of the exact non-secret files copied. WP2 needs the three config inputs and the app manifest/identity evidence; copy no conductor binding directory, broker home, credential artifact, connector credential artifact, key material, or broad root/state tree. Add an assertion that every captured relative path belongs to the allowlist and a byte-level canary check that the known ephemeral credentials do not occur in capture output. If another conductor record is genuinely required, name that single file and its redaction proof.

## R3-MR-4 - teardown is not safe enough to authorize

The literal `rm -rf $SESS` is unquoted and recursively deletes whatever an unresolved or malformed variable expands to. The plan does not specify mktemp creation, an absolute-path/prefix check, symlink refusal, or a non-empty exact-target guard. This cannot be executed under the repository's destructive-action discipline.

Required successor: specify session creation with mktemp -d under an explicit user-temp parent, record the resolved absolute path, require non-empty/existing/non-symlink/owned-by-current-user and expected-prefix checks, and delete only the quoted validated target with `rm -rf -- "$SESS"`. The post-delete check must use the recorded absolute target, not a recomputed variable or glob.

## Verdict and next transition

MUST-REVISE. Reissue s16-wp2-plan-4, uniquely parented to this review. Preserve every carried identity/config/artifact correction, same-run join, E3 labels, cache classification, routed owner question, and downstream hold. Correct the artifact lifecycle or route its required scope expansion, fully specify the decoder, use a positive non-secret capture allowlist, and make teardown mechanically safe. Do not run the capture procedure under this review.

ACTIONS_GIT_REF: read-only PLAN-REVIEW - exact-hashed and historical-linted plan 314f48cd; inspected its engine lineage, locked F63/F65 and m-8/m-9 artifact definitions, the pinned composed-test paths, Go internal-package placement, seat credential persistence at internal/seat/binding.go, daemon state, and clean banked head/upstream; no source, test, branch, commit, push, PR, merge, build, runtime, temp shim, release-binding, or evidence-package act
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit; implementation worktree clean:)
 M frank/.relays/s16/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s16/s16-wp2/PLAN-REVIEW-implementer-20260828-001759.md
?? frank/.relays/s16/s16-wp2/PLAN-REVIEW-implementer-20260828-003433.md
?? frank/.relays/s16/s16-wp2/PLAN-planner-20260828-001107.md
?? frank/.relays/s16/s16-wp2/PLAN-planner-20260828-002348.md
?? frank/.relays/s16/s16-wp2/PLAN-planner-20260828-004522.md
