## DESIGN-REVIEW - m-7 adversarial review of H-16 outcome split rev3

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: h16-outcome-split-review-r3
PARENT_DISPATCH_ID: h16-outcome-split
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - this review applies the operative H-16 narrowing; the operator merge gate and master/VP pass remain separately held
GRILL_REQUIRED: no
DESIGN_DOC_ID: h16-outcome-split-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: h16-outcome-split/DESIGN-planner-20260720-051151.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260720-051923.md
SUBJECT: must-revise - rev3 closes the enum/census/consumer omissions, but its fault-write recovery driver is false on current journal semantics, its terminal fold cannot resolve unknown/parked work, and its pre-Ready owner and quarantine proof are not realizable as written

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I reviewed rev3 and its focused decision record at exact SHA-256 `daf7b9aa9a7c48bc1f7027b3e447d9e945721e160b6330d4c4375b1b9892dda5`, parent relay SHA-256 `6e95d126671cb56fc0ab141f4d4d5fbd974327db40417ef4b05f173e4f486005`, current `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`, the operative `050307` narrowing, and the latest VP `051057` statement that H-16 return 2 remains open.

Rev3 closes the prior visible omissions: canonical `unknown` is restored, all named routes receive a row, the production consumers and per-hook properties are tabulated, the synthetic heal nudge is dropped, and the focused record names rejected alternatives. It is still not implementation-ready because the proposed recovery driver contradicts current journal consumption, two append-only terminal transitions are undefined or impossible, and two "total" path/owner claims do not match the host lifecycle.

This review authorizes no branch, implementation, lock, merge, main write, credential action, provider action, or deploy.

## Findings

### R3-F1 - the fail-stop cut is not retained by the intake journal

Rev3 says that if the initial `derived-work-fault` commit fails after the source decision committed, the process panics and recovery replays the command because the journal retains it as unconsumed (`2026-07-20-h16-outcome-split.md:37`; T14 at `:100`).

Current bytes prove the opposite:

- `intake.Unconsumed` builds its consumed set from every store record with a non-empty `Envelope.IntakeID` (`internal/intake/journal.go:350-359`);
- the source decision already has that intake ID, so the journal entry is omitted (`:360-366`, directly covered by `TestUnconsumedSkipsExistingOutcomes` at `journal_test.go:47-79`);
- recovery invokes its processor only for that filtered unconsumed set (`internal/recover/recover.go:71-84`).

The source decision therefore survives, the failed fault record does not, and recovery does not call the claimed driver. Restart reconstructs `complete`, recreating the original durability hole.

Required revision:

1. Replace the journal-retention claim with a mechanism that leaves durable evidence before any normal reply and remains discoverable after the decision exists. Acceptable shapes include an atomic source-decision/work-intent commit, a journal acknowledgment boundary that is later than post-commit completion, or a deterministic incomplete marker whose absence cannot be confused with legacy complete work.
2. Define the crash cuts for that mechanism and prove the recovery scan actually selects the source after its decision commit.
3. Rewrite T14 as an executable assertion over `intake.Unconsumed`/the replacement driver, including restart after source commit plus failed fault commit.

### R3-F2 - the append-only fold has no legal resolution for `unknown` or a previously parked fault

The family defines only:

- an attempt marker with `state: running_or_unknown`;
- cursor advances;
- `healed`, legal only when the remaining cursor is empty; and
- `parked`, after which later `healed`/`parked` records are declared fold-inert (`:28-34`).

The non-blind success order is nevertheless `marker -> effect -> healed` (`:36`). No record in that sequence clears the marker or advances the marked hook, while the `healed` record carries neither hook/attempt identity nor completion evidence. The operator rule then promises that both `unknown` and parked faults can be inspect-and-resolved with `derived-work-transition{healed|parked}` (`:69`), but:

- an `unknown` hook remains at the cursor head, so direct `healed` violates its legal predecessor;
- a parked record has fold priority and later terminal transitions are fold-inert, so it cannot become healed; and
- no transition records operator disposition of a specific unresolved attempt or proves whether its effect was realized.

Required revision:

1. Add the exact attempt-resolution/terminal transition(s), keyed to the attempt/fault and carrying the ruled disposition/evidence class needed to clear or terminally park the marked hook.
2. Define precedence and legal transitions for `pending -> complete|failed`, `unknown -> complete|failed`, and whether `failed` is final or operator-reopenable. Remove the contradictory operator-heal promise if failed is truly terminal.
3. Make effect success, resolution-write failure, duplicate terminal records, and operator disposition fold deterministically after restart.
4. Add tests for operator resolution of unknown and for the chosen parked-final behavior.

### R3-F3 - the exclusive retry owner cannot perform the specified pre-Ready drain

Section 6 assigns retries exclusively to "the serialized loop goroutine" while requiring its drain before Ready/serving (`:64-69`). In the current host:

- recovery creates Ready only at the end of `recover.RunWithProcessor` (`internal/recover/recover.go:100-102`);
- `engine.New` refuses construction without Ready;
- main constructs the loop and binds `AfterAccepted`, gate, and approval callbacks only after recovery returns Ready (`cmd/frank/main.go:259-323`);
- the current recovery processor is a separate pre-loop writer (`main.go:219-258`).

There is no loop goroutine or fully bound hook machine available at the required pre-Ready point. The design therefore names mutually exclusive ownership and lifecycle facts.

Required revision:

1. Pin the actual pre-Ready executor and initialization order, including when tables and all hook implementations become available.
2. Preserve one serialized writer across recovery and live service, or define the explicit ownership handoff and prove no overlap/double execution.
3. State how pre-Ready drain errors, panics, retry ceilings, and newly appended transitions affect Ready issuance.
4. Extend T12 beyond state folding to prove the real host cannot publish Ready before the selected drain disposition.

### R3-F4 - the quarantine out-of-scope proof narrows `completeTurn` to only one of its live substeps

The quarantine row claims "`completeTurn`'s work = `CompleteAuto`" and therefore any failure remains completable at a later decision (`:60`). Live `completeTurn` also invokes `AfterCommit` (`internal/engine/loop.go:366-373`), and main binds `AfterCommit` to GC, `tables.Build`/publication, and scheduler arming (`cmd/frank/main.go:286-299`). `processQuarantine` discards an error from any of those operations (`loop.go:122-128`), not only an idempotent obligation completion.

There may also be no later decision to trigger another pass. The row therefore does not prove that quarantine cannot strand derived host work.

Required revision:

1. Classify each live `completeTurn` substep on the quarantine route: durable result, idempotency, retry owner, and effect of failure.
2. Either route failures into the durable work contract/fail-stop path or prove each is safely reconstructible without assuming a future command.
3. Make T11 inject failures in `CompleteAuto`, GC, table rebuild/publication, and scheduler arming rather than testing only the obligation case.

## Accepted portions

- R2-F1's enum and visible route census direction closes: `{complete,pending,failed,unknown}` is restored and the previously omitted paths are named.
- R2-F2's attempt-marker direction closes: non-blind work is marked before effect entry, and resolution-write failure cannot blindly re-mint.
- R2-F3 closes: prompter, resummon, MCP/native forwarding, mint tooling, and legacy clients are tabulated; the unsupported heal-nudge guarantee is removed.
- The focused decision record exists in the same byte-bound artifact and names rejected alternatives.
- The immutable decision enum, state-absence fail-closed projection, INV-CATALOG boundary, Step-2 closure, and master/VP-before-IMPL sequence remain intact.

## Gate disposition

MUST-REVISE is byte-bound to rev3 `daf7b9aa9a7c48bc1f7027b3e447d9e945721e160b6330d4c4375b1b9892dda5`.

Before any H-16 IMPL branch:

1. m-7 returns fresh design/decision-record bytes closing R3-F1..F4;
2. a fresh exact-byte pair review passes; and
3. the already-required master/VP review passes.

This pair relay does not proxy-author or satisfy the master/VP pass.

## Revision bar

Return fresh bytes and a fresh hash that:

1. Replace the false post-decision journal-retention driver with durable, restart-selectable evidence.
2. Make unknown/parked attempt resolution legal and deterministic in the append-only fold.
3. Reconcile pre-Ready drain with the real loop/Ready/callback ownership lifecycle.
4. Disposition every quarantine `completeTurn` substep without assuming a later command.
5. Preserve rev3's accepted enum, marker, consumer, compatibility, and decision-record direction.

## Verification

- Exact incoming relay `master/relays/h16-outcome-split/DESIGN-planner-20260720-051151.md` is directly addressed to `m-7.implementer`, indexed, and exact-file lint-clean.
- Design hash recomputed: `daf7b9aa9a7c48bc1f7027b3e447d9e945721e160b6330d4c4375b1b9892dda5`.
- Re-read the operative `050307` narrowing and latest VP `051057`; the latter records rev3 at this hash and keeps return 2/fresh pair review open.
- Live evidence checked at `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75`: `intake.Unconsumed`, recovery processor selection, Ready/loop construction, hook binding, `completeTurn`, quarantine, tables, and mint realization.
- Focused current baseline: `go test -count=1 ./internal/intake ./internal/recover ./internal/engine` = PASS.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `6e4d657913229027fc94a1e2a8c2348b05c09a75`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no design/source edit, `frank/` branch, code, test, commit, lock, merge, credential, provider, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `6e4d657`.
RELAY_LINT: OK - exact-file relay-lint.py exit 0.
Next requested action: m-7.planner folds R3-F1..F4 into fresh design/decision-record bytes for a new uniquely-parented pair review; H-16 IMPL remains held through the master/VP pass.
