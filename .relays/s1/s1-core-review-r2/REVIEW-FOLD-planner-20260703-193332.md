## Team s1-core — s1-slice-1: targeted re-check verdict — fold round 2 (NARROW: two blocker residues; everything else verified green)

ROLE: Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s1-core-review-r2
PARENT_DISPATCH_ID: s1-core-review-implementer-report
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no for the fold itself; merge stays a separate human gate; the exit-gate SITREP waits on this round landing green
DESIGN_LOCK_ID: s1-slice-1-design
PLAN_LOCK_ID: s1-slice-1-plan
BUNDLE_ID: s1-slice-1
OWNER: s1-core
REPO: frank/ (this repo)
BASE: main@7f66057
TARGET_BRANCH: main
IN_REPLY_TO: s1-core-review/REVIEW-FOLD-implementer-20260703-192813.md
FROM: s1-core.planner
TO: s1-core.implementer
CC: s1.orchestrator-planner, operator
SUBJECT: re-check of 9c1839e..7f66057 — 12 of 14 findings verified folded on disk (do NOT re-touch those); TWO blocker residues remain: (RB1) the crash matrix is still not driven — F11 is a call-site grep, f10 is still the only SIGKILL test; (RB2) recovery wake re-issue (D-9 step 6) is still absent; plus one P1 residue (RM1)

Targeted re-check (per the round-1 relay's stated procedure; every item below re-verified by me on disk this session at E1/E2):
VERIFIED FOLDED — no further action, do not churn: B1 (loop emits `safeReason` path-free codes — loop.go:79-138); B2 core (`TestP1LoopOutcomeDoesNotLeakStorePaths` drives a real loop error path; formatter strips a path-bearing Reason); B3 (`ServeAuthenticated` + `auth.Resolve` + `auth:invalid-credential` typed reject — server.go:61,164-166,244-249); B4 first half (`RunWithProcessor` re-enqueues `intake.Unconsumed` before listener startup — recover.go:15-40); B6 (classmap.go:13 enumerates the real corpus root; report regenerated); B7 (`cmd/frank` assembles recovery→intake→loop→authenticated socket→3 tools; `TestFrankBinaryAssemblesAuthenticatedSubmitProjectRead`); M1/M2 (`resolveMu` deleted — grep count 0; verdict rides the submit/loop path with the `parks_gate` parent edge — derived.go:51-73); M3 (projection intents on accepted commit; durable park record); M4-M7 and the O-partials (`golangci-lint` 0 issues, `go vet` clean, `go test ./...` and `go test -race ./...` PASS — all re-run by me). FOLD_SCOPE discipline: artifact filed pre-edit, all rows in — clean.

### BLOCKER RESIDUES (round-1 requirements not yet met)

RB1 — The crash matrix still does not exist (round-1 B5, unfolded). Current state, verified: `test/fixtures/f11_test.go` now asserts (a) the registry names and (b) that each name has a `crashpoint.Hit("...")` call site in source — a static source-grep, not a crash test. It drives no mutation, injects no crash, runs no recovery, asserts no byte-convergence, counts no renames. `grep -rln FRANK_TEST_CRASHPOINT --include='*_test.go'` still returns ONLY `test/fixtures/f10_test.go`, which still drives a synthetic `WriteFileAtomic`, not a real `store.Commit`/submit. Consequently the exit-gate lines C1, C2, C7-crash-legs, and F9-across-a-real-crash remain unproven — the design §3 rows and plan Task 10 require them as spawn+SIGKILL+recover fixtures. Required (unchanged from round 1, now concretely enabled by your B7 assembly + B4 re-enqueue):
  1. A child-process crash harness that drives a REAL mutation through the assembled pipeline (or the loop+store directly), with `FRANK_TEST_CRASHPOINT=<name>[:<nth>]` set, per S1 mutation class {submit-accept, submit-reject, held, operator-verdict, park, outbox-enqueue}.
  2. For each (mutation class × registered crash-point on its path): SIGKILL fires, then `recover.Run`/`RunWithProcessor`, then assert the store is byte-equivalent to fully-committed or not-committed (no third state), and assert exactly one canonical rename per mutation via an instrumented fsio counter (F11's actual property).
  3. The specific named legs: C1 (crash `pre_rename` mid-commit ⇒ nothing committed, staging cleaned, re-enqueued once, exactly-once after re-run); C2 (crash post-pivot/`pre_delivery_write` ⇒ committed once, no duplicate, wake re-issued — pairs with RB2); C7 (crash after gate-accept before park/outbox completion ⇒ exactly one park + one outbox item; crash again DURING completion ⇒ still exactly one of each); F9-whole across a real crash (N enqueued, K outcomes, SIGKILL, recover ⇒ exactly N−K re-enqueued in arrival order, zero re-emission of the K).
  If a specific (class × point) pair is unreachable in S1 (e.g. delivery points on classes with no delivery), record the exclusion in the fixture with a reason — never silently skip.

RB2 — Recovery wake re-issue (D-9 step 6) is still absent (round-1 B4, second half, unfolded). Verified: no nudge/re-delivery path exists in `internal/recover/` or `cmd/frank/main.go` (grep clean); the only wake-adjacent code is verdict targeting in submit.go. A relay committed but not yet delivered when the process dies is never nudged after restart — the L1 exit-gate line ("a dropped/lost wake is re-issued at recovery; no parked lane sleeps forever") is false in code. Required: at recovery (after re-enqueue, before/at listener open), re-issue the nudge for committed-but-undelivered deliveries and for parked lanes whose verdict committed pre-crash (derive "undelivered" from store state — e.g. mailbox entry present but no delivery marker, or re-nudge idempotently on reconnect); plus the L1 recovery leg and the W1 kill-while-parked leg asserting it.

### MUST-FIX RESIDUE

RM1 — P1 surface coverage is partial (round-1 B2 residue). The real-loop-path capture is a genuine fix, but the guide's declared watch-surface says P1 covers "push frames + tool descriptions, not just bounce/error text." Verified: no P1 leg sweeps the assembled server's tool-list/tool-description strings or a captured push frame. Required: extend P1 to grep (a) the `tools/list` response + every tool description served by the assembled binary and (b) at least one captured nudge/push frame, for the path families. Cheap now that the assembly fixture exists.

Fold protocol: your fold report parents THIS relay, writes its own FOLD_SCOPE artifact before any edit, and carries FOLD_SCOPE above ACTIONS_GIT_REF. Do not re-touch the verified-folded set. Any file outside the rows below: relay the deviation first, do not edit. After this round I re-run the targeted check on RB1/RB2/RM1 only; if green, the exit-gate SITREP to the orchestrator follows.

FOLD_SCOPE:
- test/fixtures/ (new crash-matrix harness + C1/C2/C7/F9 crash-leg fixtures; f11_test.go rewrite; iph_test.go RM1 legs) -> in
- test/replay/ (only if report regeneration is touched) -> in
- internal/recover/recover.go (RB2 wake re-issue step) -> in
- internal/fsio/fsio.go (rename-counter instrumentation for F11) -> in
- internal/store/store.go (delivery-marker/undelivered derivation if RB2 needs it) -> in
- internal/engine/loop.go (crash-harness hooks on the mutation paths, if needed) -> in
- internal/channel/server.go (push-frame/tool-description capture for RM1) -> in
- cmd/frank/main.go (wake re-issue at open, if RB2 lands there) -> in
FOLD_SCOPE_RESULT: all-in

ACTIONS_GIT_REF: none — no edits made; the re-check was read-only (file reads + go test/vet/lint/race re-runs), tree clean at main@7f66057
FINAL_GIT_STATUS_SHORT: none — clean tree (`git status --short` empty at 20260703-193332; only subsequent writes are this relay + its INDEX row, both gitignored)
Next requested action: s1-core.implementer folds RB1, RB2, RM1 and reissues; no merge authority; no dispatch token exists in this relay.
