## DESIGN-REVIEW - Row-2 Option 2-prime must revise: detach the whole filesystem transaction, preserve race-safe root confinement, and tighten breaker lifecycle claims

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-row2-mechanism-review-r1
PARENT_DISPATCH_ID: s8-row2-mechanism
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded mechanism review; no operator product fork is required
DESIGN_DOC_ID: s8-design-m3-registry
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: s8-row2-mechanism/DESIGN-planner-20260712-062000.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-7.planner, m-2.planner, s8.planner, s8.implementer
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must-revise - Option 2-prime direction accepted, but the detachable boundary starts too late, path confinement is raceable, and restart/breaker reset is overclaimed

DESIGN_REVIEW_VERDICT: must-revise

Option 2-prime is the right direction: it preserves the in-process E1 class, avoids a new helper protocol that cannot reap D-state anyway, distinguishes lane-confusion classes from operator-topology residuals, and maps a timeout to the machinery-fault edge rather than no-vantage. The current §4a does not yet enforce its strongest liveness and confinement labels.

## F1 - BLOCKER: detaching only `read()` leaves earlier filesystem syscalls in the serialized loop

Section 4a says the read runs in a detachable worker. The current production path performs `filepath.EvalSymlinks` over the root and target before `os.ReadFile`; the proposed mechanism also adds `open` and `fstat`. On a stalled FUSE/network mount, pathname resolution, metadata lookup, `open`, or `fstat` may block before the first `read()`. If any of those execute on the serialized commit goroutine, the claimed “loop detaches and returns” guarantee is false.

Evidence:
- §4a Part B names only a blocked `read()` and says “The read runs in a detachable worker” (`...probe-design.md:83-85`).
- The active path resolves both root and target through `filepath.EvalSymlinks` before reading (`internal/observe/checks_base.go:17-23,68-96` at `s8-observe-spine@123628a`).
- The mechanism itself requires `open` plus `fstat`; both are filesystem interactions before the chunk loop.

Required fold:
1. Put the **entire target filesystem transaction** behind the detachable boundary: rooted resolution/traversal, open, metadata verification, every read chunk, and close/cleanup ownership.
2. The serialized control path may perform only non-filesystem state checks, launch the worker, wait on a buffered completion channel versus the deadline, atomically trip the breaker, and return the typed fault.
3. Extend the injected blocking seam so it can block before/open/metadata/read, not only inside a synthetic reader. The acceptance claim is “control path returns for a blocked filesystem operation,” at the injected seam grain.

## F2 - BLOCKER: resolve-then-open does not preserve lane-root confinement against a race

The locked §6.1 contract requires the path to remain scoped to the governed lane root. Resolving symlinks, checking the resulting string, and later opening by pathname is a time-of-check/time-of-use gap: a lane can replace a checked path component between resolution and open. `fstat` proves only that the opened target is regular; it does not prove that the opened regular file remained beneath the governed root. `O_NONBLOCK` does not close this confinement gap.

Required fold:
1. Name a race-safe rooted-open mechanism, such as descriptor-relative component traversal with no-follow semantics and final `fstat`, or an equivalent platform mechanism that proves the opened descriptor belongs to the approved root traversal.
2. Preserve the current absolute/`..` refusal and symlink-escape rejection, but do not rely on a pre-open resolved pathname as the proof of confinement.
3. State the special-file boundary precisely: lane-creatable FIFO/socket/symlink classes are refused without a blocking read; privileged/pre-existing device nodes and stalled mounted filesystems belong to the stated operator-topology residual unless the rooted-open mechanism rejects them before invoking a potentially blocking device open.
4. Add a race fixture that swaps a checked component toward an outside regular file and proves no outside bytes are read.

## F3 - CORRECTION: the breaker bound and reset lifecycle need an exact owner

“One goroutine/fd per lane, until restart” is true only if breaker state is process-wide for the active registry, checked before worker launch, and tripped before the timed-out call returns. It also assumes a clean restart is possible; a process with an uninterruptible worker may not exit promptly and may retain locks or sockets.

Required fold:
1. Assign breaker check/set to the single serialized commit-loop owner and require set-before-return, so a second worker for the lane cannot launch after timeout.
2. Define the bound as at most one detached filesystem worker per governed lane **per serving process/registry instance**; total residue is bounded by the finite governed lane set.
3. Replace casual “until restart” with “for the lifetime of the serving process; reset only after operator remediation and confirmed process replacement.” Keep failed restart/process-exit as part of the stated operator-topology residual.
4. The injected fixture must prove the second same-lane call is refused without launching another worker; a different lane remains independently usable.

## Accepted Surfaces

- `read-file` remains base/in-process and does not use the m-7 executor.
- The interim E1 deadline and s10 sunset remain unchanged.
- Deadline detach returns an explicit machinery fault: authority `held`, non-authority `rejected`/author-return; never accepted plus label.
- A per-lane breaker is preferable to unlimited detached workers.
- FIFO/non-regular and byte-ceiling fixtures are valid RED-first evidence once the rooted transaction boundary is corrected.
- The D-state residue must remain stated, never described as eliminated or reapable.
- No new terminal, c2 mechanism reopen, or lock-pinned value move is needed.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no frank source/test edit, branch, commit, push, PR, or merge
FINAL_GIT_STATUS_SHORT: `git -C frank status --short` returned empty at `main@691d034`; active `s8-observe-spine@123628a` contains existing slice-relay changes not made by m-3.implementer
Next requested action: m-3.planner folds F1-F3 into §4a and reissues for DESIGN-REVIEW; Row 2 remains held at `123628a`
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-row2-mechanism/DESIGN-REVIEW-implementer-20260712-063000.md`
