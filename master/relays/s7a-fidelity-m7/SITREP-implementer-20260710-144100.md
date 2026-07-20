## SITREP - m-7 fidelity confirms s7a registry-load and hosting surface at d76c3ad

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s7a-fidelity-m7
PARENT_DISPATCH_ID: s7a-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - scoped fidelity return; VP integration and operator merge remain downstream
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7a-fidelity-m7/SITREP-orchestrator-planner-20260710-143436.md
FROM: m-7.implementer
TO: master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-7.planner, m-2.planner, m-2.implementer, m-4.planner, m-4.implementer
SUBJECT: CONFIRM - s7a guard fails before serve with path-free registry-load error; composite-digest and A-1 boundaries remain faithful; s8 genesis condition recorded

VERDICT: confirm

### 1. Load-time disposition - CONFIRM

The new guard executes during `fieldspec.Load`: registry validation compiles each predicate and returns the column-grain error before a `Registry` is published (`internal/fieldspec/registry.go:65-81,176-195`; `predicate.go:136-153`; `registry.go:212-223`). Production wraps that as `load fieldspec registry: <logical error>` (`internal/config/config.go:57-63`). In the conductor startup order, pinned config loads at `cmd/frank/main.go:101-105`; `channel.ServeAuthenticated` is not reached until `:242`. Therefore a forbidden column predicate prevents the conductor from opening any seat channel. It is an operator-process config-load failure on stderr, not a submitted candidate, persisted rejection, or seat-visible runtime bounce.

This is faithful to the locked trusted-startup posture: config is loaded once at trusted startup, with no hot reload, and failures close authority before serving (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:107-111`). No new disposition token or runtime surface is needed.

### 2. I-PH error surface - CONFIRM

The rejection text is constructed solely from the predicate owner and logical registry atoms: `field <owner>: predicate references non gate-referenceable row field <array>.<column>` (`registry.go:212-223`). Neither `fieldspec.Load` nor `config.Load` adds the registry path for validation failures; the only wrapper is the path-free `load fieldspec registry` class. At current call sites the error exits through process stderr before `channel.ServeAuthenticated`, so it cannot reach `submit`, `project`, `read`, descriptions, projections, or delivery payloads.

That satisfies I-PH's actual grain: canonical store/config/outbox paths are forbidden in every seat-delivered surface (`master/ARCHITECTURE.md:478`) and trusted config paths/values are absent from every seat surface (`m-7 design:111`). The logical field names are operator-actionable schema identifiers, not canonical path leakage. If a future diagnostic channel surfaces arbitrary startup errors, it must map this class to a typed/path-free diagnostic rather than forwarding `error.Error()`; no such forwarding exists at `d76c3ad`.

### 3. Version marker and composite digest - CONFIRM WITH S8 GENESIS CONDITION

`s7a-fieldspec-v5` is an internal section/version marker inside the fieldspec member. Integrity remains one root: `config.Load` reads the member bytes and `config.Digest` hashes the sorted member-name/member-hash manifest (`internal/config/config.go:37-90`). An operator-only `config_change` recomputes the composite digest with the replacement member before acceptance (`internal/engine/submit.go:331-370`), atomically materializes the member (`internal/store/config_change.go:16-44`), and Phase 0 derives its expected digest from genesis plus the accepted config-change chain (`internal/store/genesis.go:125-148,185-208`). This matches the locked single-top-level-digest plus internal section-stamp contract (`m-7 design:109,165`).

No live store exists today, so editing the repository default is the correct pre-genesis path. Condition for the s8 dogfood store: genesis must be initialized from the `s7a-fieldspec-v5` bytes and must record the exact composite `config.Load(...).Digest`; any already-initialized store must adopt those bytes through an accepted operator `config_change` carrying the recomputed composite digest before restart. Direct replacement after genesis is invalid and must remain Phase-0 fail-closed.

### 4. A-1 stable-schema digest - CONFIRM, WITH THE REQUIRED DISTINCTION

`GateReferenceableColumns` is load-time predicate metadata on `FieldSpec` (`registry.go:26-42`). `Registry.Render` copies only rendered type/options/default shape into `Form` (`internal/fieldspec/render.go:40-70,101-133`), and the digest hashes that rendered form plus `ConfigDigest`, seat pattern, phase, and tier (`render.go:220-238`). The new member is not copied into `Form`; with the same `ConfigDigest`, it cannot independently rotate a form digest.

A real §7 registry adoption does change the member bytes and therefore the composite `ConfigDigest`. That intentionally changes rendered-form digests after restart and rejects forms from the superseded registry generation; it is not a spurious A-1 bounce. This is exactly the locked distinction: digest covers stable schema plus the §7 chain head, and `re-render` occurs for a config/registry generation change, never unrelated traffic (`master/domains/m-7-conductor-core/design/2026-07-06-s6-transport-amendments.md:22-26`).

### Independent E2 verification at `s7a-colgrain@d76c3ad`

- `go test -count=1 ./internal/fieldspec -run '^TestRegistryLoadRejectsNonAllowlistedRowColumns$' -v` -> PASS, all three default-deny cases.
- Focused hosting/config battery -> PASS: `TestP1S3ExtendedSurfacesDoNotLeakPathFamilies`, `TestRunningConfigUnchangedUntilRestart`, `TestS5ConfigChangeMovesOldToNewDigestOverLandedRegistry`, `TestS5ConfigChangePhase0ValidateGenesisWalksAcceptedChain`, and `TestS5ConfigChangeStaleFormRerenderShowsRoutingEscalation`.
- `go test -count=1 -p=1 ./...` -> PASS, 24 tested packages `ok`, 2 packages with no test files.
- `go vet ./...` -> exit 0, no output.
- `git diff --check 1d3e92c..d76c3ad` -> clean; `git status --short` -> empty.

Nonblocking harness note: two non-serialized full-suite attempts each hit a different pre-existing socket-startup fixture failure (`TestS6IPHSeatMintReplyCarveOutsScoped`: `auth:invalid-credential`; `TestConfigChangeProjectionsCarryNoMemberBytes`: connection refused). Each failed test passed immediately in isolation, and the serialized uncached full suite passed. Neither fixture nor its runtime surface is in the five-file s7a diff. This is test-harness concurrency instability, not a fidelity blocker for this guard; it should remain visible to VP integration.

Scope held: read-only review of the five-file `internal/fieldspec/` diff. No branch, source, test, config, store, PR, merge, or live-store action taken.

ACTIONS_GIT_REF: read-only review of `/Users/jack/.config/superpowers/worktrees/frank/s7a-colgrain@d76c3ad660614f58d72f107330becd465b6a1eaa`; durable actions are this relay and one append-only `master/relays/INDEX.md` row only
FINAL_GIT_STATUS_SHORT: `s7a-colgrain@d76c3ad` clean (`git status --short` empty); harness cwd is not a git repository
Next requested action: combine this m-7 confirm with the m-4 fidelity verdict for VP integration. Merge remains operator-gated; the s7 row-3 fold remains downstream of the s7a merge.
