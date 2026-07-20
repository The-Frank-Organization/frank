## PLAN-REVIEW — s10 pair PLAN must revise: the 12-row fence omits four live seams required by T1/T9/T10, and the same-file/task map is not seam-named

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s10-build-plan-review
PARENT_DISPATCH_ID: s10-build-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — pair plan-review; the scope/cross-domain deviations route to master under conditions (b)/(d), while slice merge remains separately operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
PLAN_LOCK_ID: s10-comms-spine-plan-r2-s10.1
IN_REPLY_TO: s10-build-plan/PLAN-planner-20260712-201530.md
FROM: s10.implementer
TO: s10.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-6.implementer, m-2.planner, m-3.planner, m-7.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: must-revise — byte-verbatim adoption and the locked contracts are accepted, but the mechanical block is deviation-present: config.go, registry_test.go, executor.go, and cmd/frank/main.go are required OUT seams; T7–T10 overlaps and T10's live side-effect path also need executable seam naming

PLAN_REVIEW_VERDICT: must-revise

I reviewed `master/relays/s10-build-plan/PLAN-planner-20260712-201530.md` against the boot, the m-6 r2 plan-of-record, the m-6 implementation token and its 12-row fence, Q6×Q4, c3 §§2–4/9, and live `frank@8941889` source. The PLAN is correctly addressed, grants no implementation authority, and its adopted-plan sha256 is exact. The vertical order, locks, conditions (a)–(g), owner-fidelity gate, acceptance criteria, and out-of-scope carries are accepted.

Approval is blocked by the load-bearing fence. The PLAN faithfully copied the upstream 12 rows, but live source proves four additional files are required for the behavior it promises. Therefore `SCOPE_DIFF_RESULT` is `deviation-present`; this pair cannot self-authorize the files, and `s10.planner` must not issue `s10-build-impl`.

## Blocking findings

### F1 — T1's v8 marker and successor relation have no in-fence writer

The plan assigns the m-7 capability-marker move and the governed v7→v8 successor relation to T1, but their live production loci are outside the 12 rows:

- `frank/internal/config/config.go:307` is the reader capability exact-set; it currently stops at `s8-fieldspec-v7`. A v8 reader cannot accept a v8 store until this set changes.
- `frank/internal/config/config.go:406-418` is the governed adjacent-forward successor relation; it currently ends at `v6→v7`. T1's promised `v7→v8` transition cannot pass without changing it.
- `frank/internal/fieldspec/registry_test.go:15-20` asserts the embedded registry is exactly v7. Changing `registry.json` to v8 makes the ordinary package battery fail unless this existing assertion is updated.

`internal/config/config.go` and `internal/fieldspec/registry_test.go` are not covered by an `internal/config/` or `internal/fieldspec/` directory row. This is the same reader-with-no-writer / required-test-outside-fence class the s8 discipline says to stop on, not an integration-hook exception.

### F2 — T9's E2 silent-kill sunset and the live dogfood composition root are OUT

The T9 map names `observe/read_file_worker.go` and `engine/loop.go`, which covers the E1 read-file deadline and loop work, but not the live E2 kill/extend mechanism:

- `frank/internal/executor/executor.go:170-190` owns the running suite timer, `SIGKILL`, wait, and timeout verdict. An operator `extend` decision cannot affect a running E2 check without changing this owner or naming a different in-fence mechanism that actually controls it.
- `frank/cmd/frank/main.go:146-158` constructs the production executor host and observe registry, and `:268-294` constructs/wires the production loop and accepted/commit hooks. The new ODB/scheduler/kill-extend path cannot run on the claimed live fresh-v8 dogfood surface unless this composition root is changed or an already-wired in-fence consumer is identified. None is named.

Both files are OUT. Merely adding engine units and fixture calls would prove components, not the stated live production path or the sunset of the actual E2 kill locus.

### F3 — T10's RED premise is not grounded in the shipped mechanism

The PLAN says the T10 RED fails because a static `side_effecting` gate is present. At `8941889`, the production registry contains only `read-file`, `git-status`, and `run-suite` (`internal/observe/registry.go:97-110`); no `side_effecting` entry or static config allowlist is wired. Unknown entries reject during claim validation, and `executor.Host.Spawn` hard-refuses every non-`suite` class (`internal/executor/executor.go:83-95`).

The revised PLAN must name:

1. the representative registered `side_effecting` operation that makes the RED probative rather than vacuous;
2. the exact in-courier live-prompt state/reader/writer seam from refusal → ODB/park → operator verdict → allowed execution;
3. how the existing executor defense boundary consumes an approved side-effecting operation without silently reclassifying it as `suite` or weakening the fail-closed floor; and
4. the named negative proving an unapproved operation still cannot execute.

If this requires changing the executor interface/policy shape or another cross-domain surface, condition (d) routes it to master with m-3/m-7 fidelity before the PLAN is reissued. The operator sunset remains accepted; the blocker is the missing executable path, not a request to revisit the ruling.

### F4 — same-file multi-task edits are file-named, not seam-named

The boot makes this a mechanical PLAN-review condition. Current §4 maps `engine/loop.go` to T7 + T8 and again to T9/T10, while `internal/observe/` spans T6/T9/T10; it names whole files/packages but no per-task functions, new units, or non-overlapping edit boundaries. The generic statement that seams will be recorded later does not satisfy the pre-token check.

The revised task map must name the edit seam for each overlapping task, especially T7/T8/T9/T10 in `engine/loop.go`, T6/T9/T10 in `internal/observe/`, and every task wired through `cmd/frank/main.go` if master adds that file. If two tasks intentionally share one function, name the shared function and the ordered edit/commit boundary so the diff→license table is mechanically checkable.

## Mechanical scope check

SCOPE_DIFF:
- frank/internal/fieldspec/registry.json -> in
- frank/internal/fieldspec/validate.go -> in
- frank/internal/fieldspec/predicate.go -> in
- frank/internal/fieldspec/render.go -> in
- frank/internal/store/genesis.go -> in
- frank/internal/obligation/obligation.go -> in
- frank/internal/engine/submit.go -> in
- frank/internal/engine/loop.go -> in
- frank/internal/engine/ -> in
- frank/internal/observe/ -> in
- frank/test/fixtures/ -> in
- frank/.relays/s10/ -> in
- frank/internal/config/config.go -> OUT
- frank/internal/fieldspec/registry_test.go -> OUT
- frank/internal/executor/executor.go -> OUT
- frank/cmd/frank/main.go -> OUT
SCOPE_DIFF_RESULT: deviation-present

All 12 copied rows match the m-6 token byte-for-byte, and all 11 existing source/directory anchors resolve at `8941889`; that proves faithful copying and anchor existence, not row completeness. The four OUT rows are the mechanical block.

## Accepted ground to preserve

- Adopted plan hash verified exactly: `f3e508bd356920ebf6a38b815cd44a19e29d750f900090a0a5cb8249e40f5777`.
- SEQ-2 remains fresh v8 genesis; PARK-ACROSS-V8 remains avoided by T1/T2-before-T4 sequencing; the full 8a branch remains s11.
- Q6×Q4 remains one MAJOR-but-safe interpreter-bearing v7→v8 transition, marker-first, no migrator, history un-reclassified.
- Byte-exact `{accepted, rejected, held}`, R2, Rails A/B, I-PH, stop-and-hold, owner-byte fidelity, both sunsets, crash legs, both mechanical tables, and the ten INV-CATALOG laws remain binding.
- m-2 supplies/finalizes T1 registry bytes; m-7 owns the capability move and T8 fidelity; m-3 owns T6/T9/T10 fidelity; m-6 guides; merge remains operator-only.
- The s10 worktree exists, is clean, and is based at `8941889`; no implementation work has begun.

## Required revision / routing

1. Do not issue `s10-build-impl`.
2. Route the four OUT rows and the T10 cross-domain execution question to `master.orchestrator-planner` under conditions (b)/(d); request a bounded fence ruling rather than silently adding files.
3. Reissue the PLAN after the ruling with the exact authorized block, T1's real capability/successor seams, T9's E1 + E2 timer/kill-extend seams, T10's concrete registered-operation/approval/runner path, the production composition seam, and per-task same-file edit boundaries.
4. Preserve every accepted lock, condition, acceptance criterion, owner gate, and out-of-scope line above. No design ruling is reopened by this review.

## Verification

- Exact PLAN lint returned `OK`; relay-root mode reported the target PLAN `OK` alongside known unrelated `INDEX.md` and older lineage noise.
- `shasum -a 256` verified the adopted plan-of-record hash exactly.
- Process-substitution `diff` found no difference between the token's and PLAN's 12-row `SCOPE_DIFF` blocks.
- E1 source inspection at `frank@8941889`: `internal/config/config.go:307,406-418`; `internal/fieldspec/registry_test.go:15-20`; `internal/store/genesis.go:98-134`; `internal/executor/executor.go:83-95,170-190`; `internal/observe/registry.go:77-110,174-210`; `cmd/frank/main.go:146-158,268-294`.
- Worktree inspection verified `s10-comms-spine@8941889` with clean short status.
- Post-write exact relay lint returned `OK`; relay-root mode also reported this target `OK` alongside the known unrelated root noise.
- Mechanical readback returned `scope_rows=16` and `SCOPE_DIFF_RESULT: deviation-present`; the matching index row occurs exactly once.
- Post-write status readback: `frank/` main and `s10-comms-spine` were both clean at `8941889`; cwd root remained not a git repository.

ACTIONS_GIT_REF: wrote this PLAN-REVIEW relay and appended one `master/relays/INDEX.md` row; no `frank/` source/test/plan-doc edit, no branch/worktree creation, no commit, no implementation dispatch, no merge.
FINAL_GIT_STATUS_SHORT: unavailable — cwd `/Users/jack/Programming/harness` is not a git repository; post-write `frank/` main and `s10-comms-spine` worktree were both clean at `8941889`.
Next requested action: s10.planner routes the deviation and T10 execution seam to master, then reissues the PLAN after a bounded ruling; implementation remains held.
