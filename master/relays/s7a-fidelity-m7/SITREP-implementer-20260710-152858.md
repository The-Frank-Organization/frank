## SITREP - m-7 final-bytes re-confirm at s7a-colgrain@2bc0763

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s7a-fidelity-m7
PARENT_DISPATCH_ID: s7a-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - narrow final-bytes fidelity return; VP re-approval and operator merge remain downstream
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7a-fidelity-m7/SITREP-orchestrator-planner-20260710-151000.md
FROM: m-7.implementer
TO: master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-7.planner, m-2.planner, m-2.implementer, m-4.implementer
SUBJECT: CONFIRM final registry bytes at 2bc0763 - provenance is load-inert; composite-digest and s8-genesis condition re-confirmed over corrected artifact

VERDICT: confirm

### Final-byte disposition

The `d76c3ad..2bc0763` JSON delta changes exactly the four `provenance` values. `version: s7a-fieldspec-v5`, the column guard, `GateReferenceableColumns`, the `declared_deviated` singleton, named enums, predicates, and every field row are unchanged. The companion test replaces a nonempty-owner check with the exact four-value tuple. No hosting or semantic mechanism moved.

The final `internal/fieldspec/registry.json` member SHA-256 is `1ef6abab4d496b11017f57ca400e8296d63824994ffce8311e4533f70cc92485`. Because the member bytes changed, the composite digest value computed by `config.Load(...).Digest` also changes. The integrity rule is unchanged: the digest remains the single sorted member-manifest root, and the internal provenance/version metadata remains covered by that root.

### S8 genesis condition - RE-CONFIRMED OVER FINAL BYTES

The s8 dogfood store must be initialized from these exact final `s7a-fieldspec-v5` bytes, and genesis must record the composite `config.Load(store.StoreRootConfigPaths(root)).Digest` computed with this corrected member. Any already-initialized store must adopt these bytes through an accepted operator `config_change` carrying the recomputed composite digest before restart. Direct post-genesis replacement remains invalid and must fail Phase-0 integrity validation.

This is the same condition recorded in `SITREP-implementer-20260710-144100.md`, now explicitly bound to `2bc0763` and the final member SHA above.

### Provenance load-inertness - CONFIRMED

Repository-wide Go usage of `Registry.Provenance` is only the exact-tuple assertion in `registry_test.go`; render, validation, config loading, recovery, engine disposition, and channel serving do not consume it. The map is parsed metadata covered by the composite digest, not rendered form shape or runtime policy input. Therefore:

- A-1 remains unchanged: with the same `ConfigDigest`, provenance cannot perturb the rendered form or form digest. Adoption of the changed member intentionally rotates `ConfigDigest`, so pre-change forms re-render after restart as a real §7 generation change.
- Fail-before-serve remains unchanged: bad predicate semantics are still rejected by registry validation during `config.Load`, before `channel.ServeAuthenticated`; correcting provenance neither adds nor removes a predicate.
- I-PH remains unchanged: no provenance value is emitted on a seat surface, and the load-fault text remains the same path-free logical-field error.

### Independent E2 verification at `2bc0763`

- `go test -count=1 ./internal/fieldspec -run '^(TestRegistryV2MemberParsesAndExposesLockedEnums|TestRegistryLoadRejectsNonAllowlistedRowColumns)$' -v` -> PASS, including exact final provenance and all three default-deny cases.
- `go test -count=1 ./internal/config ./internal/store` -> PASS.
- Focused config-chain battery -> PASS: composite digest moves with replacement member, Phase 0 validates the accepted chain, and stale forms re-render after the registry generation change.
- `go test -count=1 -p=1 ./...` -> PASS, 24 tested packages `ok`, 2 packages with no test files.
- `go vet ./...` -> exit 0, no output.
- `git diff --check 1d3e92c..2bc0763` -> clean; worktree status empty.
- `git diff --name-only 1d3e92c..2bc0763` -> the original five `internal/fieldspec/` files only.

Scope held: read-only final-byte review. No branch, source, test, registry, store, PR, merge, or live-store action taken. The prior m-7 load/I-PH/A-1 confirms stand; m-4's semantic confirm is not reopened.

ACTIONS_GIT_REF: read-only review of `/Users/jack/.config/superpowers/worktrees/frank/s7a-colgrain@2bc0763`; durable actions are this relay and one append-only `master/relays/INDEX.md` row only
FINAL_GIT_STATUS_SHORT: `s7a-colgrain@2bc0763` clean (`git status --short` empty); harness cwd is not a git repository
Next requested action: return the revised final-byte package to `master.orchestrator-reviewer` for s7a re-approval. Any merge decision must then be routed separately `TO: operator`.
