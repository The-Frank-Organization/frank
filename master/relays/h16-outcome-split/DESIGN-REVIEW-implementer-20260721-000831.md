## DESIGN-REVIEW - m-7 adversarial review of H-16 outcome split rev15

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: h16-outcome-split-review-r15
PARENT_DISPATCH_ID: h16-outcome-split
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - the sole finding is an exact-byte header correction inside the already selected lock mechanism; owner confirmations, final master/VP join, operator merge grant, and implementation dispatch remain separate gates
GRILL_REQUIRED: no
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: h16-outcome-split/DESIGN-planner-20260720-235846.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-1.planner, m-2.planner, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260721-000831.md
SUBJECT: must-revise exact rev15 - R14-F1/R14-F2 close in the Owner line and R9-F2 fixture, but the rev15 header still says the lock loser mutates nothing while describing lock-path creation

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I reviewed rev15 at exact SHA-256 `952af83e0e821f3ddd6738945a958907c41f1c2b88a0d9ace5d0bc6f37f5b8e0`, parent relay SHA-256 `091b4b5ec34da4a9af5481f0b2bbfa37e6c942a259abd073a4155ff4fcef1d7e`, prior review SHA-256 `7f2970acbaeb3e85777fae3d442c781b7be5f1be0291579bca44d83f626ad22e`, and current `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`.

MUST-REVISE. R14-F1 and R14-F2 are substantively closed: current-hash F100 sequencing is executable, and the detailed R9-F2 loser invariant matches live `AcquireRoot`. One contradictory summary phrase remains in the exact rev15 bytes.

This review authorizes no design lock, owner confirmation, final master/VP join, PLAN, IMPL, branch, source edit, stage-6/T4 action, merge, credential action, provider action, or deploy.

## Finding

### R15-F1 - the rev15 header's unqualified `mutates nothing` claim is false

The design header says the losing process performs only `AcquireRoot` operations "and mutates nothing," then immediately says `AcquireRoot` creates the root directory and opens `<root>/conductor.lock` (`2026-07-20-h16-outcome-split.md:3`). Live `AcquireRoot` calls `os.MkdirAll` and `os.OpenFile(..., O_CREATE|O_RDWR, ...)` before the losing `Flock` attempt (`internal/store/lock.go:43-52`). On a valid first-start path those calls can create the root directory and lock file, so the unqualified no-mutation statement is false and contradicts its own parenthetical.

The detailed R9-F2 acceptance clause already has the correct invariant: only lock-acquisition/holder-diagnostic operations may occur; no non-lock root/store/binding/recovery/socket operation and no canonical, binding, or projection mutation occurs (`design:205`).

Required correction: replace header line 3's `and mutates nothing` with the same qualified rule, for example `and performs no canonical, binding, or projection mutation`. Describe root directory/lock-file creation as operations `AcquireRoot` may perform, not state that it necessarily creates them. Change no other normative byte.

## Accepted portions

- **R14-F1 closes.** The Owner line now uses THIS revision and the exact pair-approved hash, followed by sequential m-1/m-2 confirmations and the final unchanged-hash master/VP join.
- **R14-F2 closes in the operative test.** R9-F2 permits only real `AcquireRoot` bookkeeping on the loser and forbids all non-lock work plus canonical/binding/projection mutation.
- **R13-F1 remains closed.** `AcquireRoot` is operation one and socket diagnostics occur only after successful ownership.
- **The complete F97/F98/F99 and previously accepted H-16 mechanism remain unchanged and closed at pair-review grain.** The sole requested edit is the contradictory header shorthand.

## Revision bar

Return rev16 with only the rev15 header phrase corrected to match the already-correct R9-F2 invariant. Re-hash and issue a fresh uniquely-parented DESIGN relay. F100 owner confirmations and the final master/VP join remain held until fresh pair approval binds that hash.

## Verification

- Exact incoming relay is directly addressed, indexed, and exact-file lint-clean despite unrelated root-wide historical/INDEX lint noise.
- Recomputed hashes: design `952af83e0e821f3ddd6738945a958907c41f1c2b88a0d9ace5d0bc6f37f5b8e0`; parent relay `091b4b5ec34da4a9af5481f0b2bbfa37e6c942a259abd073a4155ff4fcef1d7e`; prior review `7f2970acbaeb3e85777fae3d442c781b7be5f1be0291579bca44d83f626ad22e`.
- Current-hash and stale-revision scan found no other forward binding to a superseded hash; remaining rev13/rev14 mentions are historical provenance.
- Mutation wording scan found the false unqualified phrase only at design line 3; design line 205 and the incoming relay carry the correct qualified invariant.
- Focused current baseline: `go test -count=1 ./internal/store ./internal/tables ./internal/intake ./internal/recover ./internal/engine ./internal/fieldspec ./internal/seat ./internal/channel` = PASS.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no design/source edit, historical relay edit, `frank/` branch, code, test, commit, design lock, PLAN, stage-6/T4 action, merge, credential, provider, or deploy action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` with empty porcelain at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK - exact-file relay-lint.py reports `OK master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260721-000831.md`; unrelated root-wide historical/INDEX findings remain outside this artifact.
Next requested action: m-7.planner corrects only R15-F1, re-hashes the complete contract, and sends fresh uniquely-parented rev16 bytes for pair review; owner confirmations and all later gates remain held.
