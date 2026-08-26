## SITREP - WP3 CODA + WP4 COMPLETE at pushed tip fc585cd7: exact 64 GREEN / 0 RED / 64; shared strict-integer canonicalizer and verified broker control handover banked; draft PR refreshed; HOLD for planner carriage, m-7 coda review, and WP5 plan

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s16a-impl-7
PARENT_DISPATCH_ID: s16a-impl-7
RUN_ID: s16a
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - the authorized coda and WP4 are locally complete; this report requests planner carriage and the bounded m-7 coda review and carries no WP5, merge, ready, CI/CD, E3, exit, deployment, or release authority
IN_REPLY_TO: s16a-impl/IMPL-planner-20260826-023700.md
FROM: s16a.implementer
TO: s16a.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-7.implementer, m-10.planner, m-9.planner, m-8.planner
SUBJECT: coda + WP4 COMPLETE - pushed tip fc585cd7; exact 64/0/64; shared strict-integer canonicalizer, verified broker control replacement, 32-byte token mint; draft PR refreshed; HOLD for planner carriage, m-7 review, and WP5 plan
STATUS: COMPLETE - HOLD after filing for planner carriage, m-7 bounded coda review through master, and the planner's WP5 plan; no WP5, merge, ready, CI/CD, E3, exit, deployment, or release act is authorized by this report
PLAN_LOCK_ID: s16a-build-19

## Closed result and commit map

The WP3 coda and WP4 are complete at clean local/remote-equal branch tip `fc585cd75545d4ca7bbeec50e2ab715ac8da94ce`, pushed to `origin/s16a-conformance`. Tagged census is exactly `GREEN=64 RED=0 TOTAL=64`; D01 and D02 moved from their authorized RED state to GREEN in WP4, no other row changed color, and D03 remains absent.

1. `1d90616d56ac2b6ddeaa697630057c278391e98a` - converted the controller to one whole-file POSIX `fcntl(F_SETLK)` write lock, bounded by caller context, with the process-local one-descriptor guard and subprocess proofs of exclusion and stable inode.
2. `b4f5e28c0b51340958e3116c77f53456e0214467` - added Linux `SO_PEERCRED` / Darwin `LOCAL_PEERPID`, after-accept-before-generation `F_GETLK` holder verification, concurrent accepts, verified strictly-greater live-session replacement with old-session close, and typed in-memory `{adopted, rejected-lock, rejected-token, rejected-generation}` outcomes surfaced through the client. There is no durable `control_handover` or `adopted` recording and no m-10 recording-surface byte.
3. `4073138efbc11935278c49f9f17ce43e620db75c` - changed fresh-spawn control tokens to 32 OS-CSPRNG bytes encoded as 64 lowercase hexadecimal characters, with repeated format and distinctness proof.
4. `fc585cd75545d4ca7bbeec50e2ab715ac8da94ce` - replaced the three JCS implementations with one `internal/canonicaljson` implementation; all three old loci are thin wrappers, not shadow codecs. The shared implementation preserves arbitrary-size integer tokens byte-verbatim, refuses `-0`, fractions, exponents, and Go float values through a typed number error, and retains strict UTF-8, surrogate, duplicate-key, UTF-16 ordering, minimal-escape, and single-value rules. D01/D02 now execute an exhaustive shared vector corpus and digest checks through all three wrappers.

Every commit was pushed after the plain suite, vet, and moving-census gates. CT-C09 stayed GREEN across all coda commits. WP4 first demonstrated the authorized RED: the old float path rounded `9007199254740993` to `9007199254740992`; extraction then made D01 and D02 GREEN.

## Required coda and disclosure captures

The c4 acquisition choice is a nonblocking whole-file `F_SETLK` attempt retried every 10 ms under the caller's context deadline. This gives bounded adoption acquisition while preserving one live descriptor and never unlinking the lock file.

The D3 no-production-caller supplement now carries six loci: `f59.New`, `wire.NewCodec`, `DecodeSettlementManifest`, `SanitizeEnv`, `NewSocketPair`, and, as the sixth locus, `broker.NewCore` / `Core.Invoke`. The broker endpoint remains legitimate future work and is not claimed inverted by this close.

The broker outcome surface is deliberately in memory only. Verification precedes token and generation evaluation. A verified strictly-greater generation replaces the live session and closes the old connection; lock, token, and generation negatives fail closed with their respective typed outcomes. The existing durable `broker_control` transaction remains the controller-truth source before connect.

## Named coupled test updates

The refuse-everywhere rule invalidated old tests and fixtures that pinned floating JSON acceptance. The authorized coupled updates were:

- `internal/connector/jcs` and `internal/worker/jcs`: former RFC-float acceptance/digest cases now assert strict float refusal and arbitrary-size integer preservation/equivalence.
- `cmd/frank-connector/main_test.go`, `internal/connector/catalog/catalog_test.go`, `internal/connector/control/control_test.go`, `internal/connector/service/service_test.go`, and `test/seam/connector_app_test.go`: the optional catalog cost fixture changed from decimal `1.25` to integer `1`; production schema and predicates did not change.
- `internal/connector/request/request_test.go`: the valid fixture omits the optional deferred `temperature: 0.2`; production request schema and predicates did not change.
- `cmd/frank-app/main_test.go`: the fake broker now returns the canonical typed `adopted` handshake reply required by the new protocol.

The first shared-codec full census exposed C07 at `63/1` because the stale decimal catalog fixture was now honestly refused. The coupled fixture updates above restored C07; final census is 64/0/64. This was an expected refuse-everywhere compatibility correction within the token, not a production-contract change. One Darwin integration test initially hung because its temporary Unix socket path exceeded Darwin's limit; the test moved to a short `/tmp` socket and closes its pipe on startup error. No production deadline or protocol grammar changed.

## Verification at final committed and pushed bytes

- `go test -p=1 -count=1 ./...` PASS at `fc585cd7`; `test/fixtures` completed in 230.924 seconds.
- `go vet ./...` PASS.
- Focused broker/client/app/shared-codec packages PASS; `go test -tags seam -count=1 ./test/seam/ -run '^TestCT_(C09|D01|D02)$'` PASS.
- Tagged census PASS: `SUMMARY GREEN=64 RED=0 TOTAL=64`, all 64 rows GREEN.
- Untagged sentinel remains loud: `census invalid: seam build tag sentinel absent`, census exit 2.
- Exactly 64 `TestCT_` functions; locked ID set is bijective; D03 and `t.Skip` are absent.
- `gofmt -l cmd internal test` returned no paths; `git diff --check` PASS.
- Linux package builds for broker and brokerclient PASS; Darwin amd64 broker package cross-build PASS.
- `go.mod`, `go.sum`, D03, `fieldspec.CanonicalMarshal`, master/contracts/ledger, m-10 recording surfaces, and unrelated conductor internals were not edited.
- Worktree is clean; local `s16a-conformance`, remote `origin/s16a-conformance`, and draft PR #1 head all equal `fc585cd75545d4ca7bbeec50e2ab715ac8da94ce`.

## PR and next governed act

Draft PR #1 remains `https://github.com/The-Frank-Organization/frank-dev/pull/1`, OPEN and DRAFT, base `main`, head `s16a-conformance` at `fc585cd75545d4ca7bbeec50e2ab715ac8da94ce`. Its body is refreshed to 64/0/64 and the completed coda/WP4 state. No CI/CD, mark-ready, merge, deployment, release, E3, or exit act occurred.

Planner should carry this report through master, request m-7's bounded review of coda commits `1d90616d..4073138e`, and author the WP5 plan. Pair HOLDS. This E2 close does not claim E3 or slice exit.

ACTIONS_GIT_REF: coda/WP4 commits `1d90616d56ac2b6ddeaa697630057c278391e98a` through `fc585cd75545d4ca7bbeec50e2ab715ac8da94ce` pushed to `origin/s16a-conformance`; draft PR #1 refreshed and remains DRAFT at the same head; daemon/client kit 2.9.1 fingerprint `ddd5f9f5c84e3c51dbe06ce7d1d7174604b5f0c17524c155091f3c6738bcc6c0`; no merge, ready, CI/CD, E3, exit, deployment, or release act
FINAL_GIT_STATUS_SHORT: implementation worktree clean at local/remote/PR-equal `fc585cd75545d4ca7bbeec50e2ab715ac8da94ce`; governing checkout pre-draft has modified active INDEX, modified master CHECKPOINTS, and three daemon-rendered current relays untracked; daemon rendering for this filing appends active relay/index state
