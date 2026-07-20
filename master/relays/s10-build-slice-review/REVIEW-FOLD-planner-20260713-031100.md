## REVIEW-FOLD — the corrected-authority fold dispatch you requested: the SAME bounded findings (MF-1 + AO-1 + CT-1, nothing added), now carried under PHASE REVIEW-FOLD / fold-in-only with the findings-scoped FOLD_SCOPE below; all three owner confirms are in with no contradiction — this fold is the LAST gate before the merge-decision relay

ROLE: Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s10-build-slice-review
PARENT_DISPATCH_ID: s10-build-impl
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — a pair-internal fold; merge remains operator-only via HUMAN_MERGE_AUTHORIZATION
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
PLAN_LOCK_ID: s10-comms-spine-plan-r2-s10.2
IN_REPLY_TO: s10-build-slice-review/SITREP-implementer-20260713-030123.md
FROM: s10.planner
TO: s10.implementer
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-2.planner, m-3.planner, m-7.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: your authority hold was correct — the prior relay's canonical fields said report-only and prose does not override them; this relay is the same three bounded findings from `s10-build-slice-review/SITREP-planner-20260713-024152` reissued at fold authority, with the FOLD_SCOPE fence being exactly the four rows your hold named; TDD the negative first, one fold commit, full uncached battery FILE-captured, fold report here with your own pre-edit FOLD_SCOPE above ACTIONS_GIT_REF

**The findings (restated by reference, unchanged from `…-024152`):**
1. **MF-1 (must-fix):** mirror the `Envelope.From == "operator"` guard from `approval.go:113` (live) and `:142` (replay lookup) into the expiry equivalents (`expiry.go:124` and the `Apply` path `:139-158`), plus ONE behavioral negative proving a non-operator-authored Accepted `resolves_gate` record can NOT produce `Extend` through either expiry decision path. RED first.
2. **AO-1 (accepted-optional, you accepted):** the one-line registered-carry comment at `observe/registry.go:203` marking the approved→`side-effecting-execution-refused` branch as the deliberate no-op.
3. **CT-1 (fold-report text only):** distinguish the catalog pin (`943f07…`) from the fieldspec-v5 pins (`1ef6abab…`) in the fold report's claim text; no code.

**The fold fence (findings-scoped; an OUT discovery stops the fold and returns here):**

FOLD_SCOPE:
- frank/internal/engine/expiry.go -> in
- frank/test/fixtures/s10_timeout_test.go -> in
- frank/internal/observe/registry.go -> in
- frank/.relays/s10/ -> in
FOLD_SCOPE_RESULT: all-in

**Execution shape (your own stated plan, adopted):** write your pre-edit FOLD_SCOPE artifact over these same rows; TDD the MF-1 negative (expected RED at head `9722744`); implement the mirrored guards; AO-1 comment; full uncached `go test ./... -count=1` FILE-captured under `frank/.relays/s10/`; ONE fold commit; fold report to this dispatch with FOLD_SCOPE above ACTIONS_GIT_REF and the CT-1 label correction in its text.

**Context for the record:** all three batched owner confirms are RETURNED and CONFIRM with no contradiction — m-2 (`s10-build-owner-confirms/SITREP-planner-20260713-030500`, byte-exact, tripwire held, SEQ-2), m-7 (`…-030217`, all four cells, `Spawn` re-proven at their seat, the two engineering notes accepted as named ledger residuals), m-3 (`…-030000`, all three edges, §4a invariants preserved, one Low completed-then-killed ledger note). Your fold report is therefore the last artifact before I author the merge-decision relay TO the operator.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/s10-build-slice-review/REVIEW-FOLD-planner-20260713-031100.md` — run before handoff.
- MF-1's loci re-confirmed by your own E1 verification (`…-030123` §Independent finding verification) and my lens evidence; the finding set is closed — nothing new was added at reissue.

ACTIONS_GIT_REF: none — a fold-authority dispatch; no edits by this relay
FINAL_GIT_STATUS_SHORT: unavailable — cwd (/Users/jack/Programming/harness) is not a git repo (docs workspace); the s10 worktree clean at `9722744` per your fresh status check
Next requested action: execute the fold per the shape above; fold report to this dispatch; on its green I author the merge-decision relay TO the operator (merge only via HUMAN_MERGE_AUTHORIZATION).
