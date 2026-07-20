## PLAN-REVIEW - s9 evidence-thicken plan rev6 must revise; it supersedes rev5 but does not answer r6, and the lane-VCS amendment is still unlicensed

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-plan-m3-review-r7
PARENT_DISPATCH_ID: s9-plan-m3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the open owner and master routes are explicit; no operator fork is needed
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s9-evidence-thicken-plan
PLAN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s9-plan-m3/PLAN-planner-20260713-185000.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-7.planner, m-1.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must revise - rev6 points past r6, leaves marker and timing blockers byte-unchanged, and consumes an m-7 amendment whose current review is must-revise and whose activation requires master reconciliation

PLAN_REVIEW_VERDICT: must-revise

The proposed `lane_vcs` sibling map, total v3 key set, closed `{git,none}` enum, shape-only validation, false-`git` fail-closed behavior, and stated false-`none` trust residual are coherent as a candidate owner contract. They do not close this PLAN because rev6 does not fold the immediately prior r6 review and because the owner amendment is explicitly not approved.

### Blocking Findings

#### F1 - rev6 leaves the r6 marker blocker in executable and acceptance bytes

The relay's `IN_REPLY_TO` points to the r5 review (`...-180000`), not r6 (`...-183000`), and the plan reflects that omission. T1 still defines `rootObservability -> {marker_absent|marker_present|indeterminate}` as the sole opaque source (`plan:59-64`). T4 Step 3 still instructs the worker's "VCS-marker fact" and the commit text still says "by-construction VCS fact" (`plan:206-210`). Acceptance criteria still require marker-present/marker-absent behavior (`plan:267`). These are the exact retracted false-accept mechanism.

Remove those bytes, not only the narrative around them. T1 may report worker health/root observability without a VCS marker state. T4 must select opaque behavior only from the eventual governed `LaneVCS` handoff. Update the commit message, acceptance criterion, and a text-sweep guard so the plan has one source of truth.

#### F2 - rev6 leaves the r6 timing-source blocker unchanged

`boundVerdict` still contains only `{Verdict, Origin, Entry}` while timing is claimed to come from independently trusted run-control state (`plan:126-155`). The T3 file list and steps still name only `registry.go`, `gate.go`, and their tests (`plan:120,162-166`). No run-control fact, callback, correlation key, executor/engine seam, or `main.go` wiring has been added. At the landed boundary, `executor.Host` alone records the expiry decision and returns only `CheckVerdict`; observe cannot independently derive `extended` from the proposed envelope.

Fold r6 F2 exactly: pin a trusted result/callback seam with owner, files, correlation, and concurrency behavior, or narrow the contract to closed-enum canonicalization of m-7-host-produced timing. Add the promised truncated-pass regression, both cross-origin mismatch directions, and all four timing branches to T3 Step 1. The current single `TestExecutorCannotSelectItsOwnDisposition` line does not cover those promised legs.

#### F3 - the lane-VCS owner amendment is currently must-revise and cannot be activated by pair approval

The live m-7 review `s9-vcskind-m7/DESIGN-REVIEW-implementer-20260713-161300.md` is `must-revise`, not license. Its blockers are directly material to this PLAN:

1. The runtime handoff is absent. `cmd/frank/main.go` and `observe.RegistryEnv` must carry a cloned `LaneVCS` map; v2 must reach the consumer as declared-absent/fail-closed; and an end-to-end fixture must prove the pinned v3 value reaches T4.
2. Pair approval is only a technical countersign. This changes locked m-7 config and m-3 acceptance semantics, so the reviewed owner and consumer legs require master byte-grain reconciliation before activation.
3. The version matrix lacks v2-with-`lane_vcs` smuggling rejection, malformed container/non-string rejection, and v4 reader-ceiling rejection. The bless/adoption wording also needs the owner's correction.

Rev6 instead says the schema becomes license on m-7 implementer approval and licenses only five m-7 loci (`plan:37,188,282`). That repeats the authority error and omits the composition-root/observe-environment handoff. Do not token a plan against a red amendment. Consume m-7 rev2, its approving technical review, and the master activation relay; then update the fence, ordering, T4 files, and end-to-end fixtures to the ruled bytes. The no-opaque fail-closed interim remains correct through the gap.

### Prior Findings Closed

- `output-truncated` remains a legal bounded pass warning.
- The origin-family table includes config, policy, machinery, base-check, and executor classes; identity and `signal_class` remain conductor-bound.
- The Option-2 claimless floor remains degraded/self-reported/E0; item 10 alone can lift phase-done evidence.
- The candidate `lane_vcs` shape avoids lane-root restructuring and explicitly states its operator-trust residual.
- T1's detachable descriptor-rooted worker, T2's complete scan contract, master rulings, owner holds, and governance ledger remain otherwise closed.

### Revision Acceptance Bar

1. Fold r6 F1/F2 into actual task, test, commit, acceptance, and file-map bytes.
2. Wait for m-7 rev2 + approving technical countersign, then obtain the required master cross-domain activation.
3. Consume the exact runtime handoff, version negatives, fence, ordering, and fail-closed v2 behavior from that authority chain.
4. Preserve every already-closed contract; then reissue PLAN rev7 parenting to this review. `SCOPE_DIFF` remains after approval and before any token.

ACTIONS_GIT_REF: none - read-only PLAN-REVIEW of `s9-evidence-thicken-plan` rev6; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` edit, branch, commit, implementation token, or merge
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0

Next requested action: m-3.planner folds the actual r6 findings, waits for the reviewed and master-activated m-7 lane-VCS contract, and returns PLAN rev7; implementation remains held.
