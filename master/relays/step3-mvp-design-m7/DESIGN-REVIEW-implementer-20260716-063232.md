## DESIGN-REVIEW - m-7 adversarial re-review of transport, broker, and conductor-identity r5

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m7-review-r5
PARENT_DISPATCH_ID: step3-mvp-design-m7
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the finding is protocol-closure work inside the accepted m-7/m-10 boundary
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-mvp-design-m7-broker-placement-grill
DESIGN_DOC_ID: step3-mvp-design-m7-transport-broker
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m7/DESIGN-planner-20260716-062711.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-1.planner, m-1.implementer, m-2.planner, m-3.planner, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260716-063232.md
SUBJECT: must-revise - a PREPARING timeout cannot resume epoch E after m-10 has durably made E+1 current, and no-ack cannot distinguish pre-commit from commit-before-ack

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I re-reviewed the fresh r5 contract bytes at SHA-256 `3e88bce85128c3d69dac5827e783f77d56f2bf928132b53d591a31560221d321`, uniquely parented from the r4 review. I checked the new lock/peer predicate against Darwin's actual `F_GETLK`/`LOCAL_PEERPID` facilities and drove the transition machine through m-10's cited durable-then-visible epoch supply and the commit-before-ack ambiguity.

R5 closes R4-F1 in the required standard-primitive shape and closes R4-F2's unstable-snapshot defect: the broker proves that the connected peer is the live fcntl lock holder before evaluating the token/generation; PREPARING freezes one transition-identified set; same-ID replay is stable; crossing rows predate install; and completion-before-install has an explicit disposition. Section E is correctly rebound to m-1's approved exact SHA. One recovery edge still violates the epoch authority contract.

This review grants no interface lock, consumer approval, PLAN, T4 token, code, credential provisioning, provider call, `frank/` edit, merge, or deploy.

## Finding

### R5-F1 - PREPARING timeout re-authorizes an epoch that m-10 has already durably fenced

Section 2.5 says m-10 proposes E+1, then on m-10 loss before `CROSSERS_DURABLE` the broker times out PREPARING into local abort, lifts the barrier, and lets E continue (`2026-07-16-step3-mvp-transport-broker.md:129-140`; FX-TB-17 at `:347`). The same contract says the broker's epoch authority comes only from m-10 (`:119`), and explicitly consumes m-10 B.4. That supply contract persists the new epoch before any use, publishes the durable current state to m-8 and the broker, and makes m-10 reject the old epoch (`2026-07-16-mvp-ipc-manifest-seam-contract.md:73-80`). Therefore a proposal for E+1 is not a tentative broker-local hint: E+1 is already durable/current outside the broker. A broker timeout cannot make E authoritative again.

The timeout predicate is also unobservable as written. If m-10 commits the crossing set and crashes before its ack reaches the broker, the broker sees the same no-ack condition as a crash before commit. Local abort plus "E continues" can then both re-authorize stale E and abandon a transition whose exact crossing rows are already durable. The same-ID replay rule covers a lost ack only while the broker retains PREPARING; the timeout path defeats that protection.

Required revision:

1. Separate abort of a transition attempt from epoch authority. Once E+1 has been durably published/proposed, the broker must never admit/send/forward under E again. Loss or timeout while PREPARING leaves the broker suspended/fenced; it does not lift the barrier back to E.
2. Pin no-ack recovery without guessing whether m-10 committed. On control recovery, query/reconcile the transition ledger by `epoch_transition_id`: if the exact frozen set committed, resume that same transition/ack and install E+1; if it did not, durably abort the attempt and prepare a fresh transition for the still-current E+1. A fresh ID may replace an attempt only after the old ID is durably terminal.
3. Define what happens to any committed crossing rows when an attempt is durably aborted before install, including frozen operations that completed during PREPARING. No row may remain `crossing-pending`, and an aborted attempt must not be reported as a cross-epoch install.
4. Update CI-3 and FX-TB-17 with both indistinguishable halves: crash before crossing-set commit and commit-before-ack crash. In neither half may E resume; the latter must recover the exact committed set under the same transition ID or transactionally resolve it before a new ID.

## Accepted portions

- **R4-F1 closes.** `F_SETLK` plus connected-peer PID and `F_GETLK` holder PID is a real live-controller predicate on the pinned Linux/Darwin floor; the stale token holder without the lock fails before its fabricated generation is considered. The control token remains honestly classified, the same-user residual remains named, and CI-1/FX-TB-16 carry the negative.
- **R4-F2 closes except for the recovery edge above.** The PREPARING barrier, stable transition identity, immutable crossing set, keyed rows, duplicate/conflict rules, completion-before-install disposition, and pre-install durable ack close the missing/false-row race from r4.
- **The standard fcntl lifetime caveat is a build obligation, not a new design blocker.** Traditional POSIX record locks are process-associated and can be released when that process closes any descriptor for the same file. The build/fixture must realize the contract's "held continuously" invariant with one stable lock-file inode/region and no same-process close that drops it.
- **m-1 is closed at the owner grain.** Section E correctly names pair-approved SHA-256 `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c` while preserving m-7's separately routed consumer confirmation.
- **All earlier accepted portions remain closed:** read/quarantine honesty, typed Describe, mapping split, credential custody, retry classification/fencing, event key/dedup and honest uncoupled residual, F65/F68 canonical carriers, own-process placement, complete-and-deliver, and the grill lock.
- **No operator decision is needed.** The correction is a fail-closed recovery rule and a CI-3 delta inside the existing m-7/m-10 boundary. Placement, authority ownership, and the F57 ceiling do not change.

## Revision bar

Return fresh contract bytes and a fresh hash that:

1. Never resumes E after E+1 has been durably published by m-10.
2. Reconciles PREPARING by transition ID after no-ack without inferring whether commit occurred, and resolves every row before replacing an aborted attempt.
3. Adds the pre-commit-crash and commit-before-ack-crash fixture legs to CI-3/FX-TB-17.
4. Preserves the accepted r5 controller proof, frozen-set machine, m-1 binding, every earlier accepted decision, the F57 ceiling, m-10 no-seat/no-conductor-secret rails, and the no-conductor-change result.

On revision, issue a fresh uniquely-parented DESIGN for exact-byte re-review. CI consumer confirmations remain held.

## Verification

- `shasum -a 256 master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md` -> `3e88bce85128c3d69dac5827e783f77d56f2bf928132b53d591a31560221d321`.
- Exact incoming relay `master/relays/step3-mvp-design-m7/DESIGN-planner-20260716-062711.md` lints OK; routing, parent, `DESIGN_DOC_ID`, and grill lock match.
- Checked the consumed m-10 B.4 durable-then-visible/current-epoch supply at `master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md:73-80`.
- Verified Darwin's local `fcntl(2)` documents `F_GETLK` returning `l_pid`, process-death unlock, and traditional record-lock close semantics; the installed SDK defines `LOCAL_PEERPID` in `sys/un.h`.
- Verified exact-byte m-1 approval at `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing contract/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file relay-lint.py exit 0
Next requested action: m-7.planner revises only R5-F1, re-hashes the contract, and sends fresh uniquely-parented final bytes for pair re-review; CI consumer confirmations remain held.
