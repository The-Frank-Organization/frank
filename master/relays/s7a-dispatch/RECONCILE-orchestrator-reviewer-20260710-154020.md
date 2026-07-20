## RECONCILE -- s7a final-byte re-review: provenance and merge routing discharged; revise disclosure because fresh verification exposed a real pre-existing channel-close race

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s7a-dispatch
PARENT_DISPATCH_ID: s7a-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes -- operator merge remains a separate downstream decision after this disclosure/ledger correction and VP re-approval; no merge now
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-2.planner, m-2.implementer, m-4.implementer, m-7.implementer
IN_REPLY_TO: master/relays/s7a-dispatch/RECONCILE-orchestrator-planner-20260710-153158.md
SUBJECT: revise disclosure only -- s7a final bytes are accepted, but FLAKE-SOCKET-PAR is incomplete and misclassified: Client.Close/readLoop has a reproducible production double-close panic on both base and branch

VERDICT: revise

## Prior Findings

1. **VP F1 is discharged.** At `2bc0763`, `s7a-fieldspec-v5` carries truthful m-2 / F-S7-R2-COLGRAIN / s7a-plan-m2 attribution and a byte-exact four-value tripwire. The pair-reviewed two-file fold changes no predicate, field row, singleton, version, or hosting behavior. The independently recomputed final member SHA-256 matches m-7's `1ef6abab4d496b11017f57ca400e8296d63824994ffce8311e4533f70cc92485`.

2. **VP F2 is discharged.** The package now states that any merge decision follows through a separate relay `TO: operator`; CC grants no action authority. This re-review grants no merge authority.

## Blocking Truth Finding

3. **The current non-blocking flake disclosure is incomplete and labels a real runtime race as harness-only.** My first fresh parallel `go test -count=1 ./...` at `2bc0763` failed in `cmd/frank-mcp/TestShimReconnectsAndRetriesSingleCallAfterConnectionLoss` with `panic: close of closed channel` at `internal/channel.(*Client).Close` (`server.go:523`). This is a third test, not either of the two socket-startup fixtures listed under `FLAKE-SOCKET-PAR` in `STEP-2-KICKOFF.md:72`.

   Source inspection identifies the production race: `Client.Close` and `Client.readLoop` each perform an unsynchronized select/default check followed by `close(c.done)` (`internal/channel/server.go:519-525,555-562`). Two concurrent closers can both observe the channel open and then both close it. The focused race run makes it deterministic enough to reproduce:

   - `go test -race -count=20 ./cmd/frank-mcp -run '^TestShimReconnectsAndRetriesSingleCallAfterConnectionLoss$'` -> **FAIL/panic** at `s7a-colgrain@2bc0763`.
   - The same command -> **FAIL/panic** at unchanged `main@1d3e92c`, once through `Client.Close:523` and once through `readLoop:560` across the two stations.

   This proves the defect is **pre-existing and not an s7a regression**, so the accepted five-file branch must stay untouched. It also proves the current wording "test-harness concurrency instability" is not accurate for this newly observed leg: MCP reconnect can exercise a production client-lifecycle double-close and panic the process.

   **Required docs/ledger fold before s7a re-approval:** register a separate named runtime owed item for the `Client.Close`/`readLoop` close-once race (do not silently subsume it under the two socket-startup fixtures), owner m-7, with disposition **before s8 dogfood opens live MCP channels**. Record the proof above, the pre-existing/non-s7a classification, and the acceptance shape: one idempotent close owner/primitive, reconnect test green under focused `-race` repetition, and full battery rerun. Preserve the original two-fixture `FLAKE-SOCKET-PAR` entry as its own harness/socket-startup issue unless later evidence unifies the roots.

   Correct the integration evidence wording to report the observed sequence honestly: one fresh parallel full-suite panic, then one parallel full-suite pass and one serialized full-suite pass; focused fieldspec/config tests and vet green. Do not keep summarizing the final-byte VP station as unqualified `24 ok / 0 FAIL`.

## Accepted S7a Package

- Final branch code at `2bc0763` remains accepted: authority chain, default-deny column guard, singleton, version/provenance, red-first commits, pair review, m-4 semantic confirm, m-7 final-byte confirm, five-file fence, and s8 genesis condition are all sound.
- The full `1d3e92c..2bc0763` diff is exactly the five authorized fieldspec files and passes `git diff --check`.
- The second parallel full suite and the serialized full suite both pass; `go vet ./...` is clean. These establish no s7a regression but do not erase the reproduced lifecycle panic.
- `FLAKE-SOCKET-PAR` remains non-blocking for the s7a diff, but the newly proven runtime race needs its own truthful carry and pre-dogfood gate.

## Verification

- Incoming `153158` exact-file lint -> OK; current `s7a-dispatch` and `s7a-fidelity-m7` relay roots -> OK before filing.
- Final worktree clean at `2bc076377076136628a303719446e46938a42cad`; final registry member SHA independently matched.
- Focused final-byte fieldspec/config checks -> PASS; `go vet ./...` -> exit 0.
- Full suite attempt 1, parallel -> FAIL with the reconnect double-close panic; attempt 2, parallel -> PASS; independent `-p=1` full suite -> PASS.
- Isolated non-race test and package repetition passed, while focused `-race -count=20` reproduced the panic on both main and s7a, establishing schedule sensitivity plus base provenance.
- `frank/main` and `s7a-colgrain` remain clean; no temporary worktree remains.

Next requested action: fold the truthful runtime-race owed item and corrected verification wording into the Step-2 record, leaving `2bc0763` unchanged, then return the package for final VP approval. On approval, route the merge decision separately `TO: operator`.

ACTIONS_GIT_REF: wrote this reviewer relay and appended `master/relays/INDEX.md`; all code/worktree actions were read-only test and source inspection; no source edit, commit, merge, or push.
FINAL_GIT_STATUS_SHORT: `frank/` main = `## main...origin/main` at `1d3e92c`; `s7a-colgrain` clean at `2bc0763`; cwd is not a Git repository.
