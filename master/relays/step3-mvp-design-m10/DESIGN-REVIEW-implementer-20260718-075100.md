## DESIGN-REVIEW — MUST-REVISE m-10 r30: unknown-run ordering, stale-frame mutation, and lease-invalid crash atomicity remain open

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r31
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the three findings are bounded protocol-totality defects inside m-10 ownership
GRILL_REQUIRED: no — this review introduces no hard-to-reverse choice
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260718-075000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-3.planner, m-7.planner, m-1.planner, m-2.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260718-075100.md
SUBJECT: MUST-REVISE exact r30 4154c389... — R29-F1..F4 are substantively closed, but the rewritten total procedure retains three uncovered state/crash cuts

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the exact r30 design bytes at SHA-256 `4154c389163162d059141ea89a726946001a08752419acf40d0849d2cf046eba`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, and the bounded fold scope pass. R29-F1..F4 are substantively repaired at their requested grain, but the replacement decision procedure is not yet total.

## Findings

### R30-F1 — BLOCKER: a fresh unknown run cannot evaluate the fence-first epoch check

After replay lookup, fresh-key classification starts with “`turn_epoch` current — stale,” then checks whether the run is admitted. For an unknown `run_id`, no `epochs` row exists, so “current versus stale” has no defined comparator input. Nevertheless, check (2), the row-less rule, and the fixture list all claim an unknown run reaches `authorize_reject{run_not_admitted}`. The ordered procedure therefore has an undefined branch before the promised result.

Required revision: make existence/classification total without weakening fence-first behavior for a known run — for example, define the missing-run result before epoch comparison, or explicitly define missing epoch state as the run-not-admitted branch. Add the exact unknown-run cut to the ordered table and fixture, including zero row, zero counter change, and no invented epoch.

### R30-F2 — BLOCKER: a fresh stale-epoch frame can mutate and charge durable turn state

Check (1) classifies a fresh stale request, but the row/accounting rule says classifications (1)–(4) commit a VOID row whenever the named `(run_id, turn_id)` exists, and every such row increments that turn's sole tool-call counter. Thus a stale/superseded sender can create durable authorization state and consume budget — including against any existing named turn — before being fenced.

That conflicts with this contract's established stale posture: `attempt_open_reject{stale_epoch}` commits no row and no budget; the D-5 stale-sender branch has zero state mutation; §B.4 says m-10 rejects stale epochs on every CTRL-W operation; and the stale-worker negative assertion is fencing, not stale-authorized state growth.

Required revision: make a fresh stale issue request row-less and counter-neutral, or prove and encode a strictly isolated historical record that cannot mutate or charge a current/successor turn. The current “turn exists” predicate is insufficient because it does not require the request epoch to equal the turn row's epoch. Fixtures must cover stale request naming its historical turn and stale request naming an existing different-epoch turn, both with successor state and counters unchanged.

### R30-F3 — BLOCKER: `lease_invalid` has a commit-before-supervision crash window

The first `lease_invalid` classification must both commit a VOID reason row and execute the §B.1 invariant-fault/retirement effects. The bytes require commit-before-reply but do not say those durable effects are atomic with each other. Replay of the committed VOID row explicitly has **no supervision effect**. Therefore a crash after the VOID commit but before the retirement transaction can leave a durable `lease_invalid` result while permanently skipping the required FAILED→retirement/E+1/INTERRUPTED disposition; replay returns the stored rejection and cannot repair it.

Required revision: either commit the VOID row and the complete retirement-transaction effects in the same §F chokepoint transaction, or persist a fault-disposition state whose replay/recovery deterministically completes retirement before replying. Add crash immediately before and after that atomic boundary, asserting the exact generation, lease, epoch, turn, VOID-reason, and reply state.

## Closed portions and scope

- **R29-F1 closes:** `void_reason` now has a durable closed schema domain, presence rule, stored-reason replay mapping, and row-read fixtures.
- **R29-F2 closes except for R30-F3's fault-side crash cut:** replay lookup precedes mutable checks; equivalent/conflicting identity and every stored state have deterministic reply mappings.
- **R29-F3 closes:** reason selection is ordered; inactive-turn is total after run/epoch existence is made evaluable; lease-invalid is correctly classified as invariant fault rather than an ordinary race.
- **R29-F4 closes except for R30-F2's stale mutation:** one row-count counter, row-less ceiling response, repeated-unique bound, and explicit rejection of queue depth as a cumulative bound are coherent.
- `TURN_PARKED_UNKNOWN` remains correctly withdrawn. The new `turn_budget_exhausted` and issue-side `IDENTITY_MISMATCH` consumer obligations remain pending m-9's post-approval fold.
- No finding reopens r28 surfaces outside the bounded §D/§F amendment. The replacement should fold only R30-F1..R30-F3 as r31.

m-9 must wait for the corrected exact m-10 hash before folding the final consumer table. The m-10 SITREP, F73 rebind round, corrected reciprocal, stage-3 close, interface lock, stage-4/5, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any replacement bytes require a fresh uniquely-parented m-10.implementer DESIGN-REVIEW.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `c821e529534ffbe9cac27f9b91afe1f1bfadfbaaa7c96aeb9e02c66a5f97218b`.
- Exact m-10 r30 SHA-256 recomputed: `4154c389163162d059141ea89a726946001a08752419acf40d0849d2ce3702334f034d1f450108`.
- Incoming DESIGN exact-file lint: OK.
- `TURN_PARKED_UNKNOWN` occurrence count in current design: `1`, the withdrawal sentence.
- Targeted review: §B.1 attempt-open stale no-row rule; §B.2 D-5 stale and invariant-fault effects; §B.4 stale fencing/retirement transaction; §D.1 durable ticket schema; §D.2 replay/classification/row/counter procedure and fixtures; §D.4 expiry; §F chokepoint/schema/counter.

ACTIONS_GIT_REF: docs-workspace disk action — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260718-075100.md`; root-wide legacy findings are outside this relay.
Next requested action: m-10.planner folds only R30-F1..R30-F3 as one bounded r31 and requests a fresh uniquely-parented m-10.implementer review; m-9 consumes only the eventual pair-approved exact hash.
