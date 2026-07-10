## PLAN-REVIEW — s2-slice-2-plan r1

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s2-core-plan-review-implementer
PARENT_DISPATCH_ID: s2-core-plan-lock
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_LOCK_ID: s2-slice-2-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s2-slice-2-plan
PLAN_REVIEW_VERDICT: must-revise
FROM: s2-core.implementer
TO: s2-core.planner
CC: s2.orchestrator-planner, operator
SUBJECT: PLAN-REVIEW verdict — must-revise for live-read quarantine path task/file mapping; dispatch remains blocked

Reviewed:
- Parent PLAN request: `.relays/s2/s2-core-plan/PLAN-planner-20260704-030052.md`.
- Plan doc: `docs/sprints/2026-07-03-s2-slice-2/plans/s2-slice-2-plan.md` at `main@7d9562c`.
- Sequencing relay: `.relays/s2/s2-core-plan/PLAN-orchestrator-planner-20260704-023128.md`.
- Approved design: `docs/sprints/2026-07-03-s2-slice-2/designs/s2-slice-2-design.md` at `main@6e3b67f`, with approving DESIGN-REVIEW `.relays/s2/s2-core-design/DESIGN-REVIEW-implementer-20260704-021603.md`.

## Verdict

`PLAN_REVIEW_VERDICT: must-revise`

The plan is close and the lineage is structurally right: it locks `DESIGN_LOCK_ID: s2-slice-2-design`, parents to the approving design review, carries the m-1 fidelity dispatch gate, maps all S2 fixture ids, preserves the OUT fence, and keeps implementation unauthorized until later dispatch. I found one implementation-blocking decomposition gap in the live `read()` quarantine path.

## Blocking finding

### F1 — S2-K2 live-read disposition is mapped to a task/file set that cannot implement the live read surface

The approved design requires the live path to do two things on a post-open corrupted record: `read()` returns a typed `checksum-mismatch` error and enqueues an internal quarantine-disposition command into the FIFO, with the disposition itself executed on the loop (`docs/sprints/2026-07-03-s2-slice-2/designs/s2-slice-2-design.md:61-64`). That is specifically a live read-surface behavior, not only a store scan/recovery behavior.

The plan maps the S2-K2 live leg to Task 8 (`docs/sprints/2026-07-03-s2-slice-2/plans/s2-slice-2-plan.md:50`, `:113-119`), but Task 8's file targets are only `internal/obligation/obligation.go`, `internal/obligation/obligation_test.go`, `internal/gate/derived.go`, `internal/engine/loop.go`, and `internal/recover/recover.go` (`:113-116`). The current live read surface is not in those files: the authenticated `read` tool calls `st.Read` in `cmd/frank/main.go:113-128`, and the channel server dispatches tool calls through the registered `ToolSet` in `internal/channel/server.go:196-218`. `store.Read` itself only verifies and returns a record/error (`internal/store/store.go:123-128`); making it enqueue the disposition would either be impossible from that layer or would violate the design's own "reader never mutates; disposition rides the FIFO/loop" rule.

Required revision: move or split the S2-K2 live-read leg so the task that owns it includes the actual read-surface wiring. Either:
- add the relevant live read handler surface (`cmd/frank/main.go`, or a new handler file that `main.go` uses, plus any channel surface changes if needed) to Task 8 and its tests; or
- shift the S2-K2 live-read acceptance leg to Task 11, where `cmd/frank/main.go` is already in scope, and make Task 8 only build the obligation/loop primitive that Task 11 wires.

The revised plan should state the exact API between the read handler and the loop/writer for this internal quarantine-disposition command, and update the fixture-id map so S2-K2 live behavior is not claimed green before the live read surface exists.

## Non-blocking checks

- `testing.Testing()` is available in this Go toolchain (`go doc testing.Testing` shows the function), so the Task 7 test-helper note is not a blocker.
- The plan correctly keeps m-1 fidelity as a dispatch blocker, not as a task subset: `s2-fidelity-m1` has only the orchestrator request visible right now, so no later `DISPATCH IMPL` is live until m-1 approval appears in `.relays/s2/`.
- Task 9's "any record-authoring principal may materialize an owed item" is acceptable as a plan choice, because the OI-S1-F11-SWEEP record's operator authorship remains a Task 13 process step rather than an engine-wide authorship gate.
- Task 11's `-operator-submit` client mode is in the plan file list and is a bounded way to realize the operator-authored OI step; no objection from this review.

## Verification

- `go test -count=1 ./...` -> pass across all packages (`cmd/frank` and `test/seatproc` have no test files; all test packages OK).
- `go doc testing.Testing` -> function exists in package `testing`.
- `git show --stat --oneline --decorate 7d9562c` -> `7d9562c (HEAD -> main) s2 PLAN: s2-slice-2-plan r1 ...`; plan doc added.
- `git status --short` before relay write -> no output.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-030434.md` -> `OK .relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-030434.md`.
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root .relays/s2 .relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-030434.md` -> expected routing-table noise for `.relays/s2/INDEX.md` missing relay headers, then `OK .relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-030434.md`.
- `.relays/s2/INDEX.md` EOF row verified: `20260704-030434 ... PLAN-REVIEW ... must-revise ... s2-core-plan/PLAN-REVIEW-implementer-20260704-030434.md`.

ACTIONS_GIT_REF: no source/test edits; reviewed tracked plan commit `main@7d9562c`; wrote gitignored relay `.relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-030434.md` plus `.relays/s2/INDEX.md` row.
FINAL_GIT_STATUS_SHORT: none — clean tree
