## IMPL report - s5-b wire3 live detector wiring and archive replay leg

ROLE: Implementer
PHASE: IMPL
AUTHORITY: implementation
DISPATCH_ID: s5-b-wire3-impl
PARENT_DISPATCH_ID: s5-b-wire3-impl
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DESIGN_LOCK_ID: s5-b-mechanisms-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s5-b-mechanisms-plan
BRANCH: s5-b-wire3
IN_REPLY_TO: .relays/s5/s5-b-wire3-impl/IMPL-planner-20260706-160655.md
FROM: s5-b.implementer
TO: s5-b.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer
SUBJECT: T10 live detector wiring implemented; T11 archive replay passed; branch E2 green at 518a88f
ACTIONS_GIT_REF: s5-b-wire3@518a88f1cb34b4f6a331be78323c068b4fa23bd4; commit `518a88f s5-b: wire live detector config`
FINAL_GIT_STATUS_SHORT: none - clean code worktree

Summary:
- Implemented T10 on fresh branch `s5-b-wire3` off `main @ b30df4d`.
- Added `engine.DetectorConfigFromPinned(*config.Pinned) (DetectorConfig, error)` parsing the optional engine-member `detector` object from pinned bytes.
- Kept `internal/config/*` untouched; `EngineConfig` remains the existing unknown-key-tolerant parse.
- `cmd/frank/main.go` now parses detector config at startup and binds `env.KnownA` per submit from the current `liveTables.Snapshot()`.
- Invalid detector JSON shape or any `a_floor.member` outside `gate_category_A` now returns an error before the conductor serves.
- T11 archive replay ran read-only against `~/frank-s5-team/replay-store-dogfood-20260706`; no zeroloss harness fix was needed.

Files changed:
- `cmd/frank/main.go`
- `internal/engine/detector.go`
- `internal/engine/detector_test.go`
- `test/fixtures/s5_wire3_test.go`

TDD red evidence:
- `go test ./internal/engine -run 'TestDetectorConfigFromPinned'` failed before implementation with `undefined: engine.DetectorConfigFromPinned`.
- `go test ./test/fixtures -run 'TestS5Wire3'` failed before implementation because S1/S2 live submissions committed `gate_category:routing` without raise stamps, and invalid detector config did not fail startup with a `gate_category_A` error.

Implementation details:
- The pinned detector shape is `{detector:{a_floor:[{phase,record_kind,member}],target_branch_field,protected_branches}}`.
- Empty `a_floor` is the shipped S1 default; fixture-scoped config proves S1.
- Empty `target_branch_field` keeps S3 input-atom-pending; S3 remains mechanism-wired through existing `s3MergeSplit`.
- New binary-path fixture starts the real `cmd/frank` server, submits through the authenticated socket, and proves:
  - S1 configured hit raises from B to `authz_security`, stamps `gate_category_raised=yes`, preserves `gate_category_pick=routing`, and creates the gate outbox item.
  - Absent detector config leaves a B pick unchanged.
  - S2 referenced accepted gate raises an operator gate-resolution submission through the live path.
  - Invalid A-floor member fails before serving with a `gate_category_A` startup error.

Claim-boundary proof:
- New code/test comments carry: Step-1 detection is exactly S1 + S2 + S3 plus `other` to A.
- New code/test comments carry: S3 is mechanism-wired but input-atom-pending while `target_branch_field` is unset.
- Forbidden claim sweep over new surfaces found no `live scan`, `catches every`, `every content mis-pick`, `structurally unbypassable`, or `universal` wording.

T11 archive replay:
- Command: `FRANK_S5_REPLAY_STORE=$HOME/frank-s5-team/replay-store-dogfood-20260706 go test -count=1 ./test/replay/zeroloss/` -> exit 0.
- Verbose replay command also passed all three zeroloss tests.
- Store-derived record count: 41 `records/*.json` files.
- Redo journal count observed: one redo journal, `journal/redo/000001.jsonl`.
- Verdict: archive-compatible; no finding; archive copy not mutated.

Verification:
- Baseline in fresh worktree before edits: `go test ./...` -> exit 0.
- Target green: `go test ./internal/engine -run 'TestDetectorConfigFromPinned|TestKnownADetector'` -> exit 0.
- Target green: `go test ./test/fixtures -run 'TestS5Wire3'` -> exit 0.
- `go build ./...` -> exit 0.
- `go test -count=1 ./...` -> exit 0.
- `go vet ./...` -> exit 0.
- `git diff --check` -> exit 0.
- Combined-tree preview against unchanged `main @ b30df4d`: `git merge-tree --write-tree main HEAD` -> `f7e0898359ffbfc3f857f4799d8d9bb244ad872d`.

Out-of-scope preserved:
- No registry bytes or registry tests changed.
- No `internal/config/*` changed.
- No `cmd/frank-mcp/*` changed.
- No `internal/fieldspec/*` changed.
- No `test/replay/zeroloss/*` changed.
- No merge, push, tag, deploy, branch deletion, or worktree cleanup performed.

Next requested action:
- s5-b.planner review/reconcile this implementation report.
- Orchestrator gate can verify `s5-b-wire3@518a88f`; merge remains separately gated and is not authorized by this report.
