## PLAN-REVIEW - s9 evidence-thicken plan rev2 must revise; authority is closed, but the new E1 scan and conductor derivation are not yet mechanically safe

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-plan-m3-review-r3
PARENT_DISPATCH_ID: s9-plan-m3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the remaining findings are mechanical task corrections within the master-ratified buildable scope
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s9-evidence-thicken-plan
PLAN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s9-plan-m3/PLAN-planner-20260713-163500.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-7.planner, m-1.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must revise - master authority and blocked scope are now correct, but find-references repeats the blocked-filesystem/TOCTOU defect, its byte contract leaves ceilings and skip semantics open, signal_class lacks a conductor-bound entry identity, and opaque-lane can still wedge or misclassify config faults

PLAN_REVIEW_VERDICT: must-revise

Rev2 closes every authority and scope blocker from r2. The `160520` master relay validly selects the no-`CheckVerdict`-amendment path, ratifies B4, and orders governance-only blocked tracking; `160510` validly settles B1's PLAN-only/segment-prefix shape. The whole differential entry and item-10-dependent entries are now correctly out of buildable source scope. Approval is blocked only by the mechanics of Tasks 1-3.

### Blocking Findings

#### F1 - Task 1 repeats the blocked-filesystem and resolve-then-open defects already eliminated for E1 reads

Task 1 specifies an "in-process governed-root walk" with a five-second ceiling and says to skip symlinks/non-regular files (`plan:66-80`). A recursive walk performs root resolution, directory traversal, open, metadata, read, and close operations; any can block on the operator-topology residual. A timer checked between files cannot bound a blocked syscall. The locked s8 mechanism requires the entire filesystem transaction behind a detachable worker while the serialized path performs only state checks, launch, buffered-channel/deadline wait, breaker trip, and return (`s8 design:83-88`). The landed `read_file_worker.go:51-95,181-269` implements that rule for one file; Task 1 neither reuses nor extends the worker/breaker seam.

Likewise, "skip symlinks" on a path-based walk does not prove rooted confinement against a component swap between inspection and open. The accepted E1 mechanism proves confinement by descriptor-relative no-follow traversal and final-fd metadata, not `EvalSymlinks`/`WalkDir`-style resolve then open.

Keep `find-references` classed as in-process/base, but put every filesystem syscall in a detachable, descriptor-rooted worker. Pin per-lane active/breaker lifecycle, set-before-return behavior, the at-most-one-detached-worker bound, and injected blocking fixtures at root-open/directory-read/file-open/metadata/read/close. Add a component-swap fixture proving no outside bytes are scanned. The serialized commit loop must touch no filesystem syscall.

#### F2 - the claimed byte-grain scan contract still leaves acceptance-changing choices open

The plan names `findRefFileCeiling` and `findRefDepthCeiling` without values and gives only an 8 MiB per-file limit (`plan:70-76`). It does not bind a total-byte ceiling, maximum result count/saturation behavior, or whether symlink/non-regular/binary/oversize artifacts are excluded from the search domain versus make the observation incomplete. "Skipped" and "ceiling-refused as machinery-fault" are not interchangeable: silently skipping a potentially matching artifact can produce a false `count:0` pass.

The literal-token relation is also undefined for the accepted dotted token grammar: substring count, identifier-boundary count, and language-aware tokenization produce different results for `foo`, `foo.bar`, and embedded names.

Revise Task 1 with concrete numeric depth/file/total-byte/per-file/count ceilings; an exact textual-file domain and deterministic binary rule; exact token-boundary/count semantics; and a total disposition table. Any artifact required by the declared search domain that cannot be safely scanned must prevent E1 pass and return a symbolic bounded fault/refusal. Pin exact-ceiling/ceiling-plus-one, count-saturation, binary, symlink, non-regular, unreadable, and malformed-UTF-8 fixtures. No raw path or matched line may escape.

#### F3 - `signal_class` is derived from an executor-returned ID, not yet a conductor-bound entry identity

Master ruled derivation from the entry id plus existing outcome, with the returned `CheckVerdict` shape unchanged. Task 2 modifies only `gate.go:baseStamps` and consumes `CheckVerdict{CheckID,Outcome}` (`plan:82-95`). At the landed seam, `Registry.Run` returns the executor's verdict unchanged and `Evaluator` forwards it into `PredicateResult.Verdicts` (`internal/observe/registry.go:185-240`). Therefore `baseStamps` sees the returned `CheckID`, not an independently bound conductor selection identity.

Bind the identity before derivation: after `Run`, the conductor must validate/rewrite `CheckID` and `ClaimRef` from the selected `CheckEntry`/`Selection`, validate the closed outcome/rung/predicate combinations, and only then feed rows. Alternatively carry a conductor-internal selected-entry id alongside the verdict without changing the m-7 boundary object. Add a fixture whose executor returns `CheckID:red-to-green-differential` for a selected `run-suite`; it must never stamp `signal_class:differential`. This preserves master's no-amendment ruling while making the derivation genuinely conductor-side.

#### F4 - Task 3 does not define the no-vantage boundary or remove the current direct filesystem probe

Task 3 says an "unobservable lane" takes no-vantage, but does not distinguish a valid governed root that cannot be observed from a missing/malformed/unapproved governed-root configuration. The locked supply-set rule treats the latter as typed startup/composition failure or check refusal, never a silent ambient fallback; it is not automatically the non-authority `accepted+self_reported` opaque-lane row.

Current `evaluateAbsenceFloor` directly calls `EvalSymlinks`, `Open`, `Stat`, `Readdirnames`, and `Close` before running git (`internal/observe/checks_base.go:55-73`). Those calls can wedge and use the same resolve-then-open pattern the s8 mechanism rejected. Since Task 3 edits this function, leaving them in place would make the new opaque-lane feature preserve a known mechanism violation.

Pin the input table: invalid/absent governed config, valid-but-unobservable root, observable clean/dirty root, timeout, malformed git output, and other machinery error. Source valid-root observability from a detachable race-safe worker/fact; do not perform direct filesystem probing on the serialized path. Fixtures must prove no candidate string can request no-vantage and that timeout/machinery never reaches the opaque-lane accepted path.

### Prior Findings Closed

- The `CheckVerdict` shape remains byte-exact; `signal_class` is an m-3 row derivation under master's explicit ruling.
- The differential production entry is wholly B-diff/m-7 gated; no inert entry lands.
- `diff-shape` and `test-files-unchanged` are wholly B3/item-10 gated.
- B4 and the scope reduction are master-ratified; blocked items stay in governance relays, not `t.Skip` source stubs.
- B1 consumes the master-settled PLAN-only/segment-prefix/observe-locus direction and remains blocked pending the joint co-sign.
- Task 5 is verification-only; Tasks 6/7 correctly hold on owner confirmations.

### Revision Acceptance Bar

1. Make `find-references` descriptor-rooted, fully detachable, breaker-bounded, and race-fixtured; pin all numeric and semantic limits.
2. Define one exact search domain/token relation and fail closed whenever the complete declared domain was not observed.
3. Bind `signal_class` to conductor-selected entry identity and validate returned verdict bytes without changing the `CheckVerdict` boundary shape.
4. Give opaque-lane detection a total source-grounded input table and remove/directly replace the serialized filesystem probe.
5. Preserve every master ruling, blocked-ledger boundary, owner hold, and out-of-scope line; then reissue PLAN rev3 for review. `SCOPE_DIFF` remains after approval and before any token.

ACTIONS_GIT_REF: none - read-only PLAN-REVIEW of `s9-evidence-thicken-plan` rev2; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` edit, branch, commit, implementation token, or merge
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0

Next requested action: m-3.planner revises Tasks 1-3 against F1-F4 without reopening the master rulings or blocked ledger, then returns PLAN rev3 for review; implementation remains held.
