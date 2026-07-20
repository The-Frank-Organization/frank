## PLAN-REVIEW - s9 evidence-thicken plan rev4 must revise; Row-3 is restored, but the verdict matrix and no-VCS fact still contradict landed contracts

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-plan-m3-review-r5
PARENT_DISPATCH_ID: s9-plan-m3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - two corrections are m-3 mechanics; a new governed VCS-capability field would require the named m-7 owner route
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s9-evidence-thicken-plan
PLAN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s9-plan-m3/PLAN-planner-20260713-174500.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-7.planner, m-1.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must revise - the tuple matrix rejects the locked output-truncated pass and excludes lane-ungoverned machinery rows; allowlisting does not yet source disposition independently; marker absence beneath an arbitrary governed directory does not prove no VCS vantage

PLAN_REVIEW_VERDICT: must-revise

Rev4 correctly restores the locked Option-2 floor: clean or exactly matching claimless reports remain degraded/self-reported/E0, mismatches remain observed-false, and item 10 alone can add the complete phase-done fence. It also correctly requires tuple-level validation and conductor-side output hygiene. Those prior findings are closed. Approval is blocked by three contradictions in the proposed mechanics.

### Blocking Findings

#### F1 - the five-row tuple matrix rejects an extant, explicitly accepted pass verdict

The matrix permits `pass/pass/<entry-rung>` only when `failing_detail` is empty (`plan:124-134`). The landed executor deliberately returns `pass/pass/E2` with `failing_detail: output-truncated` when bounded capture truncates successful output (`executor.go:308-323`). This is not accidental residue: `s8_adversarial_test.go:126` asserts it, the m-7 fidelity return calls the symbolic truncation marker acceptable on PASS, and the T7 fold requires it remain distinguishable from failure classes.

Because tuple validation is stage 2 and redaction/re-derivation is stage 3, rev4 would convert that valid locked verdict into `check-machinery-verdict-tuple-invalid` before redaction can act. Add a legal **pass-warning** form for the closed non-failure marker `output-truncated`, or explicitly route it through a conductor canonicalization step before tuple validation while preserving its accepted symbolic diagnostic semantics. Add the existing truncated-pass fixture to T3's regression set. Do not silently erase or fault a previously accepted bounded marker.

#### F2 - the matrix and re-derivation rules do not provide one trusted source for terminal class

T2 and T4 both specify `unsafe/blocked/none`, class `lane-ungoverned`, `MachineryFault:true`, and fail-closed terminal behavior (`plan:94-105,165-176`). T3's only machinery row, however, admits details matching `check-machinery-*` or `executor-*`; its other `unsafe/blocked` row is a policy refusal with `MachineryFault:false` (`plan:126-134`). `lane-ungoverned` matches neither machinery family, despite the plan's claim that every T2 row is legal under T3. Implementers must not choose whether to rename the class, weaken it to non-fault policy refusal, or add a third matrix family.

Pin the missing config/composition-refusal tuple family and its exact canonical class/terminal behavior. Then define `MachineryFault` derivation from a conductor-known **producer/origin plus tuple** table, not merely from whichever allowlisted detail token arrived. The current `SuiteExecutor.Spawn` boundary returns only `CheckVerdict`; an allowlist can bound a string but does not independently establish whether the verdict came from pre-spawn policy refusal, in-process machinery, or the executor. Specify the internal conductor envelope/provenance bit (without widening the locked m-7 `CheckVerdict` boundary), and map each origin to the only legal class family. The hostile fixture must try both directions: an executor-origin result presenting a policy token and a conductor policy refusal presenting a machinery token; neither may select the other disposition.

The same origin table must state how conductor timing preserves the landed closed values (`under-timeout`, `extended`, `timeout`, `not-completed`). Measuring only wall duration cannot prove that the operator selected the extension branch.

#### F3 - absence of a `.git` marker beneath a governed root is not a by-construction no-VCS fact

The governed supply schema accepts any clean absolute directory as a lane root (`config.go:482-500`); it does not require the root to be a repository root or declare a VCS kind. Git also discovers repositories through ancestor directories. Therefore a governed lane rooted at a subdirectory of a Git worktree can have no `.git` marker beneath it while `git -C <root> status` has valid vantage. Rev4 would classify that observable lane as opaque and permit non-authority `accepted+self_reported`. A bare repository and nested-repository arrangements create further marker/result mismatches.

Do not infer capability from marker absence. Either route an explicit governed `vcs_kind:{git,none}` / repository-root contract through m-7's supply owner and validate it at composition, or specify a typed conductor probe whose exact outcomes distinguish a confirmed non-repository from timeout, executable failure, malformed output, and worker indeterminacy without parsing ambient prose. Fixture at minimum: governed subdirectory inside a parent worktree, root `.git` directory, linked-worktree `.git` file, explicit non-VCS root, command timeout/error, and candidate `unavailable`. Only the independently confirmed non-VCS case may enter opaque acceptance.

### Prior Findings Closed

- The claimless report Option-2 floor is restored to degraded/self-reported/E0; clean or matching porcelain cannot stamp phase-done E1.
- Contradictory outcome/predicate/rung combinations are recognized as a conductor-boundary problem; identity remains bound from `Selection` and `signal_class` remains conductor-derived.
- Returned detail and timing are subject to an m-3-owned, bounded, path-redacted pass before any seat-visible surface.
- T1's detached descriptor-rooted worker, T2's complete declared scan domain, all ceilings, the serialized-probe removal, master rulings, owner holds, and governance-only blocked ledger remain closed and unchanged.

### Revision Acceptance Bar

1. Preserve the accepted `pass/pass/E2 + output-truncated` marker in the total tuple contract and regression suite.
2. Add a legal config/composition-fault family for `lane-ungoverned`; derive class, machinery status, and timing from a conductor-known origin table rather than a returned prefix.
3. Replace marker absence with an actual governed/typed no-VCS source, routing an m-7 supply change if that is the honest mechanism.
4. Preserve the restored Row-3 floor, T1/T2, every master ruling and blocked-ledger boundary; then reissue PLAN rev5. `SCOPE_DIFF` remains after approval and before any token.

ACTIONS_GIT_REF: none - read-only PLAN-REVIEW of `s9-evidence-thicken-plan` rev4; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` edit, branch, commit, implementation token, or merge
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0

Next requested action: m-3.planner corrects T3/T4 against F1-F3, routing any governed supply-schema change to m-7, then returns PLAN rev5 for review; implementation remains held.
