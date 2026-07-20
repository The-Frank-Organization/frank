## PLAN-REVIEW - s9 evidence-thicken plan rev3 must revise; worker mechanics are closed, but Row-3 and verdict-byte invariants are not

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-plan-m3-review-r4
PARENT_DISPATCH_ID: s9-plan-m3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the findings are mechanical corrections within the locked design and master-ratified scope
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s9-evidence-thicken-plan
PLAN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s9-plan-m3/PLAN-planner-20260713-171000.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-7.planner, m-1.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must revise - T4 upgrades the locked claimless report floor from E0 to E1 and leaves no-VCS classification undefined; T3 validates enums independently but does not close contradictory verdict tuples or redact returned detail

PLAN_REVIEW_VERDICT: must-revise

Rev3 closes the prior worker, scan-totality, selected-identity, and direct-filesystem-probe findings. T1 now generalizes the descriptor-rooted detachable worker with a per-lane breaker and stage-complete fixtures. T2 pins the declared textual domain, token relation, numeric ceilings, and fail-closed incomplete-scan table. T3 binds identity from the conductor selection, and T4 removes the serialized filesystem probe. The master rulings, blocked ledger, differential hold, and item-10 hold remain correctly scoped. Approval is blocked by three remaining contract errors in T3/T4.

### Blocking Findings

#### F1 - T4 contradicts the locked Row-3 Option-2 E0 floor

The T4 table says a valid, observable, clean governed root "contributes to done" at E1 (`plan:137-145`). Locked s8 section 13 says the opposite under the selected Option 2: without a turn baseline, a declared-status/porcelain match verifies only the status string, cannot prove the phase-shaped done predicate, and cannot raise achieved evidence above E0 (`s8 design:286-315`). The landed code preserves that rule: a mismatch rejects as observed-false, while an exact match returns `turn-attribution-unavailable` / degraded (`checks_base.go:84-91`).

Restore the locked byte in T4. A valid observable root permits deterministic porcelain parsing and declared-vs-observed mismatch detection, but a clean or exactly matching claimless report remains `Degraded` / `self_reported` / E0 until item 10 supplies the complete turn-baseline fence. Only a declared executable claim may independently pass at its registered rung. Add a regression proving clean/matching porcelain cannot stamp phase-done E1.

#### F2 - T3 validates individual enums but permits internally contradictory verdicts

T3 validates membership of `Outcome`, `Predicate`, and `RungReached`, plus the entry rung ceiling (`plan:117-125`), but does not define allowed combinations. An executor can therefore return `Outcome:pass` with `Predicate:fail` or `blocked`. The terminal gate follows `Predicate` (`gate.go:121-147`), while claim rows, integrity, and achieved evidence follow `Outcome`/`RungReached` (`gate.go:191-210`), yielding a rejected or held record that simultaneously stamps a passing observed claim and achieved evidence.

Pin a total allowed-tuple matrix before implementation. It must bind `Outcome` to `Predicate`, constrain `RungReached` to passing outcomes and the selected entry's declared rung, and define the legal presence/value classes of timing and failing detail. Every impossible tuple must become one bounded typed machinery fault before the verdict reaches `PredicateResult` or `baseStamps`. Add adversarial fixtures for pass/fail, pass/blocked, fail/pass, skipped/pass, non-pass-with-rung, and any other rejected tuple.

#### F3 - T3 does not implement the locked verdict-output I-PH pass, and T4 does not source its sole no-vantage row

Locked section 6.1 requires conductor-side redaction before a returned verdict feeds section 6, including paths, command text, effective values, and unbounded output (`s8 design:138-141`). T3 currently validates identity and enums only; it does not constrain or sanitize returned timing/failing detail. An executor can therefore place a governed path or configuration value in a seat-visible detail while still satisfying every listed validation rule.

Define bounded output schemas in the binding pass: symbolic failing-detail allowlists (or discard-and-rederive), bounded timing fields, and a planted-secret/path fixture proving executor-returned residue cannot reach verdicts, rows, bounce detail, or degradation notes.

Separately, T4 names "valid but genuinely unobservable (no VCS vantage)" as the sole accepted opaque-lane row but does not specify the conductor-owned fact that distinguishes it from `git status` timeout/error or a malformed/ungoverned root. The implementation must not infer this from a candidate string or reinterpret command failure. Pin the governed descriptor/capability field or worker-produced typed result that establishes **no VCS capability by construction**, and fixture every transition. If the governed root declares VCS observable, inability to execute/read remains machinery fault; it never becomes opaque acceptance.

### Prior Findings Closed

- Find-references uses one shared descriptor-rooted detachable worker; the serialized path performs no filesystem syscall, and the breaker/race fixture surface is complete.
- Search-domain membership, token boundaries, all numeric ceilings, saturation, and incomplete-scan dispositions are pinned; an incomplete declared-domain scan cannot pass `count:0`.
- Identity derives from the conductor selection, and a forged returned `CheckID` becomes a typed machinery fault before `signal_class` derivation.
- The direct `EvalSymlinks`/`Open`/`Stat`/`Readdirnames`/`Close` sequence is removed from the serialized path; absent/malformed/ungoverned roots are config refusal, not no-vantage.
- `CheckVerdict` remains byte-exact; differential and item-10 entries remain blocked; B4, the scope reduction, governance-only blocked ledger, and owner confirmations preserve the master rulings.

### Revision Acceptance Bar

1. Restore the claimless report floor to degraded/self-reported/E0 for clean or exactly matching observable roots under Option 2; item 10 alone may add the complete E1 fence.
2. Define and fixture a total verdict tuple matrix so terminal disposition, claim rows, integrity, and achieved rung cannot disagree.
3. Complete the section-6.1 conductor binding pass with bounded, path-redacted verdict output schemas and hostile-output fixtures.
4. Source the opaque no-VCS fact from one governed conductor input/result and prove config faults, timeouts, and command errors cannot enter it.
5. Preserve T1/T2, every master ruling, blocked-ledger boundary, owner hold, and out-of-scope line; then reissue PLAN rev4. `SCOPE_DIFF` remains after approval and before any token.

ACTIONS_GIT_REF: none - read-only PLAN-REVIEW of `s9-evidence-thicken-plan` rev3; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` edit, branch, commit, implementation token, or merge
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0

Next requested action: m-3.planner corrects T3/T4 against F1-F3 without reopening T1/T2, the master rulings, or blocked ledger, then returns PLAN rev4 for review; implementation remains held.
