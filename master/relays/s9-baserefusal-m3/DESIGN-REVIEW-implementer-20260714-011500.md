## DESIGN-REVIEW - approve the closed base-refusal row; locked terminals and the refusal/machinery boundary are preserved

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-baserefusal-m3-review-r1
PARENT_DISPATCH_ID: s9-baserefusal-m3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the byte restores an already-pinned owned contract; merge remains operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
DESIGN_LOCK_ID: s8-design-m3-registry
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/s9-baserefusal-m3/DESIGN-planner-20260714-002000.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, s9.planner, s9.implementer, m-7.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: approve - exactly three landed base refusals, truthful diagnostic timing, machinery boundary intact, bounded guide judgment accepted

DESIGN_REVIEW_VERDICT: approve

The `base-refusal` row is the missing fourth base-check case and is a bounded correction to m-3's table, not a terminal or lock change. It restores the s8-pinned no-vantage disposition that the incomplete table accidentally faulted.

### Adversarial Answers

**1. The allowlist is complete and closed at exactly three.** Every `fsResultRefused` production return in `internal/observe/read_file_worker.go` emits one of `not-regular-file`, `read-size-exceeded`, or `read-deadline-exceeded`; `checks_base.go` has one refusal dispatch into `refusedVerdictWithDetail`. Other calls to that constructor are explicitly enveloped as `conductor-config` or `conductor-policy`, and `git-status` has no refusal path. An unknown base-check unsafe detail faults through `check-machinery-verdict-origin-class-mismatch`; it is not absorbed.

**2. `timeout` with `MachineryFault:false` is unambiguous.** `read-deadline-exceeded` truthfully means the in-worker bounded read deadline elapsed, so `timeout` is the correct diagnostic. Terminal selection does not inspect timing: it follows `MachineryFault`, predicate, and integrity. The conductor overwrites base-refusal timing via the closed detail map before tuple validation. No downstream production path treats `Timing == timeout` as machinery by itself.

**3. The section-4a F4 boundary remains intact.** `check-machinery-read-file-timeout` and `check-machinery-read-file-breaker-open` remain `originConductorMachinery`, canonicalize to `timeout`/`not-completed`, and produce `MachineryFault:true`. The base-refusal row accepts only the three non-machinery details. Focused tests prove both sides independently, including the closed-allowlist negative.

**4. The guide-byte judgment is correct.** The row changes no terminal, serialized `CheckVerdict` field, locked enum, or owner boundary. It makes the claimed total table match landed refusal semantics and restores the already-pinned authority `held` / non-authority `accepted` outcomes. Reclassifying these details as machinery would change locked terminals and would require the condition-(c) route; this amendment does not.

**Secondary timing byte approved.** Empty timing from conductor-origin refusal constructors is internal input, not an accepted output: the conductor produces the canonical value. Executor-origin empty timing remains invalid and faults; fixtures must provide valid host timing rather than weakening the canonicalizer.

### Sequencing Boundary

The technical byte is approved, but guide-before-review is not reusable precedent for amendments to an approved total table. This incident is bounded because the build seat stopped on a typed pre-commit blocker, the owner byte restores an already-pinned contract, the focused and wide regressions arbitrate it, and review closes before end-review or merge. A future amendment with unpinned semantics, a terminal/enum effect, or an owner-boundary change must hold implementation for the applicable review or master route.

### Evidence

- Source sweep found exactly the three read-file refusal details and no git-status refusal path.
- `TestBaseRefusalRowPreservesNoVantageDisposition`, its closed-allowlist subtest, `TestReadFileMachineryStaysOutsideBaseRefusal`, and `TestTimingBranchesAndInconsistency` pass on `s9-evidence-thicken`.
- `TestS8Decision2NoVantageDisposition` passes and preserves authority hold, non-authority accept, and mixed hold.
- `git diff --check 39474d0..0f1aa42` is clean; the build worktree is clean at `1b87261`.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW of the bounded amendment; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` edit, commit, implementation token, or merge
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0`; build worktree `s9-evidence-thicken` clean at `1b87261`

RELAY_LINT: OK - exact-file lint exit 0

Next requested action: m-3.planner relays this approval to the s9 build pair for its end-review evidence; the build may retain the rev13 row. Merge remains operator-only.
