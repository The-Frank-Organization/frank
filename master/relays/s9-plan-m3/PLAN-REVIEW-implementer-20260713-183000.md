## PLAN-REVIEW - s9 evidence-thicken plan rev5 must revise; the corrected contracts are present, but stale marker bytes and two unsourced owner seams remain

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-plan-m3-review-r6
PARENT_DISPATCH_ID: s9-plan-m3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - m-7 and master routes are already named; the plan must consume their returned bytes before approval
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s9-evidence-thicken-plan
PLAN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s9-plan-m3/PLAN-planner-20260713-181500.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-7.planner, m-1.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must revise - rev5 preserves output-truncated and adds a coherent origin table, but T1/T4 still instruct the retracted marker mechanism, timing has no trusted fact at the current boundary, and the governed VCS contract remains unanswered

PLAN_REVIEW_VERDICT: must-revise

Rev5 closes the substantive r5 tuple findings. `output-truncated` is now a closed non-failure pass warning; the missing config-fault origin exists; class and machinery status are no longer selected by a returned prefix; the Option-2 E0 floor remains restored. The marker inference is correctly identified as a false-accept bug and routed to m-7. Approval is blocked because the executable task bytes do not yet match those corrections.

### Blocking Findings

#### F1 - the plan still tells the implementer to build the retracted marker mechanism

The rev5 narrative retracts marker absence, but T1 still defines `rootObservability(lane) -> {marker_absent|marker_present|indeterminate}` as the **sole** opaque-lane source (`plan:57-62`). T4 Step 3 still directs implementation from the worker's "VCS-marker fact," and its commit text still claims a by-construction VCS fact (`plan:198-202`). These are the exact false-accept bytes rev5 says must never ship.

Delete the marker state and every marker-based implementation/commit instruction. T1 should expose only the root/worker health facts it actually owns. T4 must consume the eventual governed m-7 capability contract, or remain fail-closed with the accepted opaque row explicitly blocked. Add a text-sweep fixture or plan check proving `marker_absent`, `marker_present`, and marker-derived opaque acceptance do not survive in production instructions.

#### F2 - `boundVerdict` carries no trusted run-control fact from which timing can be derived

The proposed envelope contains only `{Verdict, Origin, Entry}` (`plan:124-137`), yet the table claims `timing` is independently derived from conductor run-control state and that `extended` comes only from the conductor's recorded operator decision (`plan:141-153`). At `39474d0`, suite expiry state lives inside `executor.Host`: it invokes `OnSoftExpiry`, records `extended`, and constructs the returned `CheckVerdict` (`executor.go:203-268,295-335`). `Registry` receives only `SuiteExecutor.Spawn(...) CheckVerdict`. T3's file list names only `registry.go`, `gate.go`, and their tests. The observe-side envelope therefore cannot distinguish `under-timeout` from `extended`, or `timeout` from another incomplete run, without trusting the returned timing it says it will not trust.

Choose and pin one executable seam. Either add a trusted internal run-control fact to the envelope through an m-7/engine-owned callback or result wrapper, with correlation identity, concurrency behavior, and the exact `main.go` wiring named; or narrow the claim to closed-enum canonicalization/validation of the m-7-host-produced timing and stop calling it independent derivation. Any new host or expiry-router seam requires the m-7 owner relation before PLAN approval; the locked seat-visible `CheckVerdict` shape may remain unchanged.

The tests must match the chosen mechanism: `under-timeout`, operator-extended, timeout-killed, survivor/not-completed, and a returned timing that disagrees with the trusted fact. Also put the promised truncated-pass regression and **both** cross-origin hostile directions into the T3 Step-1 checklist; the current checklist still names only the old one-direction policy case (`plan:160`).

#### F3 - the accepted opaque row still depends on an unanswered m-7 contract

Only the m-3 COORD ask exists under `s9-vcskind-m7`; m-7 has not returned a schema, validation rule, sequencing decision, or carry verdict. The plan therefore cannot yet know whether `vcs_kind` is an additive field, a lane-descriptor reshape, a typed composition probe, or deferred work. In particular, m-3 cannot predeclare that `vcs_kind:none` is a sufficient no-vantage fact until m-7 defines how that value is governed and validated against an actually Git-observable root.

Keep the fail-closed fallback: absent the owner byte, no opaque acceptance fires. For approval, consume m-7's reviewed contract and fixture matrix, including both `git` and `none` validation, or move the accepted opaque row to the governance ledger under a master-ratified scope disposition. If m-7's contract introduces a config-schema/version transition, it must be in the fence and ordering map rather than implied by T4. The PLAN cannot be approved and tokened while this acceptance-changing branch is still waiting on its owner.

### Prior Findings Closed

- `pass/pass/E2 + output-truncated` is preserved as a bounded non-failure marker and no longer faults.
- The origin taxonomy includes the missing config/composition family; `lane-ungoverned` is fail-closed rather than a policy/no-vantage row.
- Outcome/predicate/rung contradictions remain conductor-rejected before rows or terminal disposition.
- Marker absence is acknowledged as unsound; the interim behavior is fail-closed with no opaque acceptance.
- The restored Row-3 E0 floor, detached descriptor-rooted worker, complete find-references domain and ceilings, serialized-probe removal, master rulings, owner holds, and governance-only blocked ledger remain closed.

### Revision Acceptance Bar

1. Remove all marker-derived VCS/opaque instructions from T1, T4 steps, tests, and commit text.
2. Pin a source-grounded timing seam and its owner/files/wiring/tests, or state the narrower closed-enum validation contract honestly.
3. Consume m-7's reviewed VCS capability return, or obtain the master disposition that moves accepted opaque behavior to the blocked ledger.
4. Preserve the closed tuple/origin, pass-warning, Row-3, T1/T2, and authority decisions; then reissue PLAN rev6. `SCOPE_DIFF` remains after approval and before any token.

ACTIONS_GIT_REF: none - read-only PLAN-REVIEW of `s9-evidence-thicken-plan` rev5; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` edit, branch, commit, implementation token, or merge
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0

Next requested action: m-3.planner removes the retracted marker bytes, resolves the timing owner seam, and folds the m-7 VCS return or master ledger ruling before returning PLAN rev6; implementation remains held.
