## SITREP — four-row build HOLDS at Row 2: the granted hard wall deadline conflicts with the locked in-process E1 boundary; the polling attempt was rejected and reverted, so m-3/master must choose the executable mechanism before Rows 3–4

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s8-build-escalate-fence-r8
PARENT_DISPATCH_ID: s8-build-four-rows
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — owner/master mechanism ruling required; live adoption and merge remain operator-only
GRILL_REQUIRED: no
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
TASK_ID: four-row re-lift, Row 2
IN_REPLY_TO: master/relays/s8-build-t10-close/SITREP-planner-20260712-045500.md
FROM: s8.implementer
TO: master.orchestrator-planner
CC: s8.planner, operator, master.orchestrator-reviewer, m-3.planner, m-7.planner, m-2.planner
SUBJECT: Row 1 landed and reviewed at ac11a3e; Row 2's first pass was correctly rejected because context polling cannot interrupt a blocked filesystem syscall, then reverted at 123628a; the remaining hard-deadline requirement cannot be satisfied inside the locked in-process E1 boundary without a mechanism decision

### Landed state

- **Row 1 landed:** `ac11a3e` (`cmd/frank/main.go`, `s8_config_activation_test.go`). Production `-init` now requires catalog, validates exactly three supplied members before any store write, and refuses `present_layers.observe:true`. Both requested negatives were RED first and GREEN after. Task review: spec-compliant; one Low test-strength note (the refusal tests do not separately assert the root remains uninitialized, although the validated control flow refuses before `store.Init`).
- **m-7 adjacent-door ACK consumed:** `s8-build-escalate-fence/SITREP-planner-20260712-042245.md` says the same observe-active refusal binds §5.1.3 bless candidate validation. That entrance is named-bound and pending because the active Row 1 license does not include `internal/store/genesis.go`; no silent bless edit was made.
- **Row 2 noncompliant attempt fully reverted:** `0f5c41a` attempted regular-only/nonblocking open, a 1 MiB ceiling, and context polling; task review found it did NOT implement the hard wall guarantee. Corrective commit `123628a` reverts it. Net Row 2 production/test bytes equal the pre-row state at `ac11a3e`; focused baseline E1 tests pass and `git diff --check ac11a3e..123628a` is clean.
- Rows 3–4 have not started. Condition (f) is honored: no pushing through a red.

### Root cause — one owner sentence conflicts with another

The r7 row requires an **ACTUAL ≤5s wall deadline** over the complete lane-file operation, including a genuinely blocked filesystem call, with no stranded work. The effective m-3 design simultaneously fixes `read-file` as **class:base, in-process, no executor, no repo-code execution** (`s8-check-registry-probe-design.md` §3/§4; lines 52, 59, 68, 182).

Inside an in-process Go call:

- `context.Context` does not cancel `open(2)`, `fstat(2)`, or `read(2)`; polling `ctx.Err()` only observes expiry between syscalls.
- `O_NONBLOCK` prevents the concrete FIFO-open wedge but is ignored/insufficient for arbitrary regular-file backing stores (for example a stalled FUSE/network mount).
- `os.File.SetReadDeadline` starts after open/stat and is not a portable answer: `go doc os.File.SetDeadline` states ordinary files generally do not support deadlines and return `ErrNoDeadline`.
- a goroutine+timer returns the serialized caller but strands a blocked goroutine/syscall, contradicting the no-leftover-liveness bar and creating an unbounded resource leak under repeated checks.
- a killable helper process or executor can isolate spawn→open/stat/read and enforce timeout→kill→reap, but that reopens the locked in-process E1 boundary and adds an execution/helper surface beyond the two licensed Row 2 files.

The reviewer also correctly rejected the attempted 1 ns local-file test: it proved an already-expired context was polled, not that a blocked syscall was interrupted. Keeping that test/mechanism would repeat the exact label→mechanism error r7 exists to eliminate.

### Requested owner/master decision

Choose and authorize one executable contract:

1. **Recommended — preserve the hard liveness guarantee:** amend the E1 boundary to permit a killable, conductor-owned read helper (or m-7 executor class) with a closed protocol, no repo-code execution, ≤5s spawn-through-reap deadline, byte ceiling, regular-file-only validation, I-PH, and a true blocked-I/O kill fixture. Name the additional production/helper/test files and m-7 boundary review required.
2. **Preserve in-process E1:** explicitly weaken the owner sentence to best-effort deadline checks over approved local regular files, accepting that an in-kernel blocked filesystem syscall can exceed 5s and wedge serialized submit. This resolves implementability but does NOT preserve the r7 hard-liveness claim.
3. **Owner-supplied alternative:** m-3 names a portable in-process cancellation primitive that covers open+stat+read on ordinary governed files without stranded work; the next re-lift binds that exact mechanism and a genuinely blocked RED fixture.

Also decide whether Row 1's now-bound bless-side observe-active refusal receives a separate named `internal/store/genesis.go` + adoption-fixture row in the same re-lift or remains a public-release carry. It was not inferred from the current grant.

### Evidence and status

- Row 1 focused config-activation tests: PASS; task review approved.
- Row 2 attempted focused tests passed but were ruled non-probative; attempt reverted.
- Post-revert focused baseline E1 tests: PASS.
- Local `go doc os.File.SetDeadline`: ordinary files generally do not support deadlines; unsupported files return `ErrNoDeadline`.
- No exact-head serialized battery or four-row report is claimed: Rows 2–4 are not complete.

ACTIONS_GIT_REF: s8-observe-spine@123628a (effective landed feature head remains Row 1 at ac11a3e; 0f5c41a is neutralized by the revert)
FINAL_GIT_STATUS_SHORT:
 M .relays/s8/INDEX.md
?? .relays/s8/s8-build-t9-review/

Out-of-scope preserved: no unlicensed helper/executor, no bless-side edit, no Row 3/4 edit, no live migration, no OS sandbox, no merge/push/PR/tag.

Next requested action: master routes the Row 2 mechanism decision to m-3 and returns an exact re-lift; s8 holds at `123628a`. On a lawful Row 2 landing, Rows 3–4 resume in order and the original fidelity/battery/review/table gates remain binding.
