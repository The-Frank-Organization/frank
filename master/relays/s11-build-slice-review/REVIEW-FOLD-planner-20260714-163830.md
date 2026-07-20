## REVIEW-FOLD — the ONE B11 end-of-slice adversarial review of `d91fcfb...547ada9`: four-lens panel UNANIMOUS APPROVE, zero blockers, zero must-fix; the battery independently rerun green (uncached, my own run); eleven OPTIONAL findings consolidated below at your discretion; on your fold-or-decline report I author the merge-decision relay TO the operator

ROLE: Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s11-build-slice-review
PARENT_DISPATCH_ID: s11-build-report
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — merge remains operator-only (`HUMAN_MERGE_AUTHORIZATION` at grant time); T5/T10 remain acceptance-OPEN behind g2/dc
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
PLAN_LOCK_ID: s11-comms-thicken-plan-s11.3
IN_REPLY_TO: master/relays/s11-build-report/IMPL-implementer-20260714-162400.md
FROM: s11.planner
TO: s11.implementer
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-6.implementer, m-2.planner, m-3.planner, m-7.planner, m-5.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: whole-slice review verdict — no blocker and no must-fix across correctness/refactor-preservation/test-honesty/invariants; every optional finding is your discretion under B11 (two flagged as cheap and worth it); if you fold anything, FOLD_SCOPE discipline + a targeted check only; if you decline all, say so and I take the slice to the operator as-is

## Panel record

PANEL_CHOSEN: custom (team-of-4)
DEFAULT_ROLES_CHANGED: yes
WHY_THIS_PANEL: 16-commit / 35-path / ~2.4k-line slice on trust-critical parked-decision + operator-gate surfaces, governed by three master rulings that must be verifiable at the bytes; idiomaticity folded into the refactor lens rather than a fifth seat
ROLES:
- correctness/design-conformance — T6 §B byte-exactness, bucket/FSM/G4 contracts, T5/T10 absence
- refactor-behavior-preservation — the eight T8 items byte-identical; ruled seam compliance (executor two-seam split, five envelope sites, per-emit locus)
- test-coverage/false-green/sequence-honesty — RED-first evidence, byte-exact assertions, label→mechanism, matrix/③ NF firing power
- invariants/fence — 35/35 fence truth, terminal enum, R2, I-PH, Rail-A fail-closed, no store-write bypass

All four lenses ran read-only over the worktree at `547ada9`; each returned **approve**. Each lens's full report is retained in my session; findings below are the consolidated, deduplicated set.

## My own verification (not delegated)

- Branch shape: 16 commits, merge-base = `d91fcfb` exactly, head `547ada9`, tree clean — E1, run this session.
- Forbidden families absent from the diff name-list (`internal/observe/`, `registry.json`, store write path) — E1, run this session.
- **Full uncached battery rerun by me** at `547ada9` in the build worktree: `go test -count=1 ./... && go vet ./...` → exit 0, all packages ok (fixtures 139.2s) — E2, my own run, not the report's.
- Gate state re-audited: `s11-oq2-ceiling` still carries only the two DESIGN relays (no m-5.implementer review, no completion to master); no dc dispatch exists — the T5/T10 report-and-hold is correct, not convenient.
- Evidence package read: `t11-exit.md` carries the mandated T8 rescope wording verbatim + catch #3 (in-build, honestly logged); `t6-red-green.md` shows genuine pre-implementation REDs (the aliasing FAIL + the undefined-symbol build failure); `mechanical-tables.md` diff→license covers the 35 paths one-for-one with same-file task order named.
- One panel artifact resolved: lens 1 initially saw two socket-test FAILs — traced to its own two parallel `go test` runs contending on shared unix sockets; each test passes in isolation and my single-run battery is fully green. Not a slice defect; noted so the ledger doesn't carry a phantom.

## VERDICT

**No blocker. No must-fix. The slice conforms to the locked contracts at the bytes:** T6 builds §B exactly (three byte-distinct records, both reason tokens, no-wake, frozen π with relabel-trips/reorder-passes, fail-closed guard, real-process crash-replay to the SAME replacement identity); the eight T8 refactors are behavior-preserving with `loop.go` untouched across the entire range and persisted gate-ID/hash identity byte-for-byte stable; the buckets are true saved-query projections with the D-vs-A precedence and RAISE-ONLY promotion; the FSM is exactly seven states with `egress_blocked` local-only and no away-send; G4 has no auto-approve encoding; R2/I-PH/terminal-enum/Rail-A hold on every new surface; T5/T10 are correctly absent.

## Optional findings (your discretion under B11 — none blocks the merge decision)

**Flagged as cheap and worth folding (two):**
1. optional `internal/engine/fsm.go:48-55` + `internal/store/projections.go:360-367` — two hand-maintained failing-edge allowlists overlap intentionally (D includes `stale_choice_set`; the FSM bounce set correctly omits it) but can silently desync as edges are added. **Cross-corroborated by two lenses independently as correct-but-drift-prone.** Cheap fix: hoist the shared subset to one named set (or bind them with a comment + a relation-asserting test). If declined, this line in the relay is the ledger note so it is never mistaken for copy-paste drift.
2. optional `internal/engine/resummon_test.go:199-203` — the cadence-restart assertion is hash-tautological (any two distinct IDs differ); the real guarantee lives in the suppression check and the E2E fixture. Cheap fix: also assert the replacement's `CadenceSlot` value so the restart proof stands alone.

**Recorded, no change required (nine):**
3. optional `internal/engine/odb.go:86,91` — ineffectual first assignment to `violation`; fail-closed behavior correct; optionally tidy or comment the `sourceProjection==nil` sentinel.
4. optional `internal/engine/approval.go:80-86` (+ expiry twin) — marshal moved ahead of owner determination; diverges only on an unreachable marshal-failure path, converging to the same deny/kill fallback.
5. optional `internal/engine/resummon.go` — the snapshot path can no longer surface a `tables.Build` error from `outcomeForContentHash`; unreachable on the live in-memory snapshot.
6. optional `internal/engine/approval.go` `existingDecision` (+ expiry twin) — first-match order changed from filename order to commit order; diverges only with ≥2 accepted resolutions on one gate, which the terminal gate surface prevents.
7. optional `internal/obligation/obligation.go` `SystemOperatorRecord` — 0→1 SchemaVersion default newly applied at the held-stale-schema site; unreachable behind the `TargetSchemaVersion < 1` guard.
8. optional `internal/obligation/obligation.go:497` — host-side error wrap can carry an absolute path; the seat-facing record is path-free; optional symbolization for defense-in-depth.
9. optional `internal/config/config.go:195` — same class, host/composition-root-facing, pre-existing pattern.
10. optional `test/fixtures/s11_8a_test.go` / commit `99e945c` — T6's RED is evidence-honest but not git-provable (test + implementation in one commit); consistent with B11 straight-through; informational for the ledger.
11. optional `test/fixtures/s11_buckets_test.go:201,203` — two all-false matrix rows would pass under a no-op pipeline; backstopped by the positive-membership rows on the same projections; no change required.

## Disposition

- Fold any subset at your discretion, or decline all. **If you fold:** FOLD_SCOPE artifact BEFORE any edit (findings-scope rows only), one targeted check after (the touched packages + the s11 fixtures), NOT a full panel rerun — the fold surface is test-guarded and behavior-preserving. **If you decline all:** a one-line decline report suffices.
- Either way, report under `s11-build-slice-review` (your relay, `FROM: s11.implementer`, parented here). On it, I author the **merge-decision relay TO the operator**: the runnable slice at `547ada9` (or your folded head), T1–T4/T6–T9/T11 at E2, T5/T10 acceptance-OPEN per FINDING-4 behind g2/dc, T8 eight-of-nine under the master rescope, catches #1–#3 in the ledger. Merge stays operator-only; nothing here grants it.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/s11-build-slice-review/REVIEW-FOLD-planner-20260714-163830.md` — run before handoff (result in the inline pointer).
- Battery: my own uncached rerun at `547ada9`, exit 0 (all packages ok; fixtures 139.225s; vet clean) — E2.
- Panel: four read-only lenses, unanimous approve; findings deduplicated above with file:line evidence retained per lens.

ACTIONS_GIT_REF: none — review verdict only; no `frank/` edit, no commit, no merge, no token. Artifacts this turn: this relay + one `master/relays/INDEX.md` row timestamped 20260714-163830.
FINAL_GIT_STATUS_SHORT: unavailable — cwd (/Users/jack/Programming/harness) is not a git repo (docs workspace); `git -C frank status --short` = clean; build worktree at `547ada9` clean.
Next requested action: operator carries this to s11.implementer; the implementer folds-or-declines the optionals and reports under `s11-build-slice-review`; on that report I author the merge-decision relay TO the operator (merge operator-only, `HUMAN_MERGE_AUTHORIZATION` at grant time).
