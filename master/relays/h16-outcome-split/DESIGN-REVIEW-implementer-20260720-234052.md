## DESIGN-REVIEW - m-7 adversarial review of H-16 outcome split rev13

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: h16-outcome-split-review-r13
PARENT_DISPATCH_ID: h16-outcome-split
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - this finding corrects the F99 acceptance contract; owner confirmations, final master/VP join, operator merge grant, and implementation dispatch remain separate gates
GRILL_REQUIRED: no
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: h16-outcome-split/DESIGN-planner-20260720-233013.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-1.planner, m-2.planner, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-234052.md
SUBJECT: must-revise exact rev13 - the normative ceremony is lock-first, but R9-F2 still places its diagnostic probe before lock acquisition and therefore tests the rejected check-to-lock sequence

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I reviewed the self-contained rev13 design at exact SHA-256 `20333c83d0b55ec4b2e1bad1272882f650dfe7ff4b61fc67281df51fc9d5c71f`, parent relay SHA-256 `760b656d772bdb405935d6b00b815f91c88e4b00c51ba5d17a3bdee52351494f`, and current `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`.

MUST-REVISE. Rev13 closes F97 and F98 and states F99 correctly in the normative ceremony rule, but its mandatory R9-F2 acceptance leg contradicts that rule. Because the approval is exact-byte-bound, this is a design blocker even though the correction is narrow.

This review authorizes no design lock, owner confirmation, final master/VP join, PLAN, IMPL, branch, source edit, stage-6/T4 action, merge, credential action, provider action, or deploy.

## Finding

### R13-F1 - R9-F2 still encodes the forbidden probe-before-lock sequence

Section 4c correctly makes `store.AcquireRoot` phase -1, before every root/store/binding/recovery touch, and explicitly makes the conductor-socket probe post-lock diagnostic only (`2026-07-20-h16-outcome-split.md:114-115`). Decision entries 15 and 17 preserve the same rule and reject socket-probe-as-exclusion (`:191,193`).

The R9-F2 test text then asks for "a conductor starting between the ceremony's diagnostic probe and its lock acquire" (`:205`). That ordering requires the ceremony to perform its diagnostic before `AcquireRoot`. It is the exact check-to-lock shape F99 removed, and it cannot coexist with phase -1 lock acquisition or a post-lock-only diagnostic. An implementation satisfying that literal test could reintroduce a pre-lock socket/root observation while still claiming the F99 matrix is green.

Required revision:

1. Rewrite the race as concurrent ceremony/conductor startup before either process acquires the root lock, or as a conductor scheduled immediately before the ceremony's first operation, `AcquireRoot`. Exactly one process wins; the loser returns `root-lock-held` cleanly.
2. State in the fixture that no ceremony root/store/binding/recovery operation and no socket diagnostic occurs before the successful `AcquireRoot`; the socket diagnostic executes only after the ceremony owns the lock.
3. Preserve the stale-socket, alias-root, two-ceremony, conductor-start, duplicate/replay, and crash-cut legs. Re-hash and issue fresh uniquely-parented bytes for review.

## Accepted portions

- **F97 closes at the pair-review grain.** Rev13 materializes sections 1-10, all 21 decision entries, the route and consumer tables, and the complete T/R battery; no bare `unchanged` pointer carries normative semantics.
- **F98 closes at the pair-review grain.** Raw `channel.Client` drives the conductor's committed `system-owned` rejection fixtures; conforming native/MCP clients assert typed `schema_invalid` with zero conductor calls.
- **F99's normative mechanism is otherwise correct.** The ceremony acquires the shared root lock before all other relevant access and holds it through reply and exit; socket state is diagnostic only; the false current `-mint` precedent is retracted and its unlocked-writer defect is explicitly routed rather than treated as proof.
- **The live lock supports the required exclusion shape.** `store.AcquireRoot` takes a nonblocking filesystem lock and current tests prove same-root and symlink-alias exclusion. Current `-mint` still bypasses it exactly as rev13 records.
- **The previously accepted H-16 mechanism remains coherent:** outcome split, Class-G retry truth, canonical Class-D fold, commit-time transition rejection, predecessor-linked mint chain, completeness-gated upgrade anchor, effective quarantine through realization, ceremony matrix, one pivot per credential generation, no-authority-delta, and accepted-only system provenance.

## Gate disposition

Pair DESIGN-REVIEW remains HELD for rev13 `20333c83d0b55ec4b2e1bad1272882f650dfe7ff4b61fc67281df51fc9d5c71f`. F100 owner reconfirmations and the final master/VP join must not start from these bytes.

Return a narrowly corrected rev14 whose R9-F2 race is lock-first and post-lock-diagnostic throughout. Any new hash requires a fresh uniquely-parented pair review; all later gates remain held.

## Verification

- Exact incoming relay is directly addressed, indexed, and exact-file lint-clean despite unrelated root-wide historical/INDEX lint noise.
- Recomputed hashes: design `20333c83d0b55ec4b2e1bad1272882f650dfe7ff4b61fc67281df51fc9d5c71f`; parent relay `760b656d772bdb405935d6b00b815f91c88e4b00c51ba5d17a3bdee52351494f`.
- Materialization check found sections 1-10 and exactly 21 numbered decision entries; the only bare `unchanged` occurrence is the intended struct-field comment.
- Live evidence checked at `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`: `cmd/frank/main.go:112-122,577-600`; `internal/store/lock.go:43-95`; alias lock test `internal/store/lock_test.go:86-103`.
- Focused current baseline: `go test -count=1 ./internal/store ./internal/tables ./internal/intake ./internal/recover ./internal/engine ./internal/fieldspec ./internal/seat ./internal/channel` = PASS.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no design/source edit, historical relay edit, `frank/` branch, code, test, commit, design lock, PLAN, stage-6/T4 action, merge, credential, provider, or deploy action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` with empty porcelain at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK - exact-file relay-lint.py reports `OK master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-234052.md`; unrelated root-wide historical/INDEX findings remain outside this artifact.
Next requested action: m-7.planner corrects only R13-F1, re-hashes the complete contract, and sends fresh uniquely-parented bytes for pair review; owner confirmations and all later gates remain held.
