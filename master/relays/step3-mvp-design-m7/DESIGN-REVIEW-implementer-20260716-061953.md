## DESIGN-REVIEW - m-7 adversarial re-review of transport, broker, and conductor-identity r4

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m7-review-r4
PARENT_DISPATCH_ID: step3-mvp-design-m7
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - both findings are protocol-closure work inside the accepted m-7/m-10 boundary
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-mvp-design-m7-broker-placement-grill
DESIGN_DOC_ID: step3-mvp-design-m7-transport-broker
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m7/DESIGN-planner-20260716-061547.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-1.planner, m-1.implementer, m-2.planner, m-3.planner, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260716-061953.md
SUBJECT: must-revise - make control-generation commitment broker-verifiable and freeze/idempotently identify the epoch transition before snapshotting crossers

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I re-reviewed the fresh r4 contract bytes at SHA-256 `28b585856f67a4ce47397a07318d4bd0e71ac54274fa96324eb275916e3b692f`, uniquely parented from the r3 review. I pressure-tested the new control-generation and crossing-set protocols across the stale-controller and lost-ack windows rather than accepting their nominal transaction labels.

R4 closes the handoff-choice and durable-crossing direction: the spawn token now has one pipe-FD path and explicit custody; `operation_id` exists; m-10 crossing rows predate installation; durable `recorded` vocabulary is clean; and m-1 is now pair-approved at exact SHA `7c8b09a6...ff944c`. Two protocol claims remain untrue as written.

This review grants no interface lock, consumer approval, PLAN, T4 token, code, credential provisioning, provider call, `frank/` edit, merge, or deploy.

## Findings

### R4-F1 - The broker cannot verify that a larger `control_generation` was durably committed

Section 2.10 says a replacement app-main atomically advances `control_generation` under m-10's exclusive store lock, then presents the committed value; the broker accepts token-match plus any strictly greater generation. It calls the committed generation "broker-verifiable proof" and says a stale process cannot present a fresher value without the lock (`2026-07-16-step3-mvp-transport-broker.md:166-182`). The broker never reads m-10's store and receives no commitment proof. It sees only the shared long-lived token and a caller-supplied integer.

A stale app-main that still holds the adoption token can therefore present `installed+1` (or a larger value) without committing anything and pass the exact broker predicate. The store lock prevents two conforming writers from committing; it does not make an untrusted integer broker-verifiable. This is the same confusion-resistant stale-controller case the generation was introduced to close, not the separately accepted same-user memory-inspection residual.

Required revision:

1. Add a broker-verifiable ownership/commitment proof for the controller generation. One viable OS-backed shape is: m-10 holds an exclusive broker-control lock; the broker authenticates the control socket peer, verifies that peer is the lock holder through the platform lock/peer-credential mechanism, and only then accepts the store-advanced generation. Another reviewed standard primitive is acceptable; a bare integer is not proof.
2. If the product boundary intentionally trusts every token holder to report committed generations honestly, narrow the claim and remove the stale-controller race guarantee. That would leave R3-F1's requested stale-process exclusion unsatisfied and requires explicit disposition rather than calling it closed.
3. Pin the platform floor, peer/lock identity fields, process-start/PID reuse handling, failure disposition, and adoption ordering. Keep the control token classified as an authorizing capability and retain the pipe-only handoff/census.
4. Expand CI-1 and FX-TB-16 with a negative where a stale token holder presents a fabricated larger generation without the store lock. The broker must reject it while accepting the real lock-holding replacement's committed generation.

### R4-F2 - The two-phase crossing snapshot is not stable or idempotent until the epoch transition itself is fenced

Section 2.5 says the broker snapshots the in-flight E set at its serialization point, waits for m-10 to commit crossing rows, then installs E+1 (`:123-134`). It does not say the broker holds a transition barrier or enters a PREPARING state while waiting for the durable ack. If E admissions/retries continue after the snapshot and before install, those operations are absent from the crossing set but become cross-epoch when E+1 installs. The pre-record invariant fails.

The lost-ack/re-proposal path is also not identified as one transition. A duplicate proposal causes a re-snapshot, but there is no stable `transition_id`, prepared-set version, abort state, or rule for rows from an earlier prepared snapshot whose operations completed before installation. Such rows can remain `crossing-pending` even though no crossing occurred, while newly admitted operations appear only in the later snapshot. Operation IDs deduplicate calls; they do not make changing crossing sets one idempotent epoch transition.

Required revision:

1. Add a stable `epoch_transition_id` bound to `{run_id, from_epoch, proposed_epoch, state_seq}` and carry it through proposal, crossing-set rows, commit ack, install, abort, and recovery.
2. At the serialization point, enter an explicit PREPARING barrier before snapshotting. Until commit or abort, admit no new E operation and permit no retry send or stale push forward that could cross the pending install. State the caller disposition while preparing (bounded wait or typed suspension).
3. Pin the state machine, for example `PROPOSED -> CROSSERS_DURABLE -> INSTALLED` with terminal `ABORTED`. Repeated messages for the same transition return the same prepared set/ack; a different proposal cannot replace it silently.
4. Define completion while PREPARING: if a snapshotted operation completes before install, durably resolve it as `completed-before-install` or remove it only through a transactionally recorded abort/update. It must not remain a false crossing row.
5. Pin crash recovery for broker loss and m-10 loss at every state. On recovery, either finish the same transition or abort it durably before preparing a new one; E+1 cannot install from a different crossing set under the same transition identity.
6. Expand CI-3 and FX-TB-17 with: a new E call racing after snapshot; an E retry during PREPARING; a snapshotted call completing before install; lost crossing-set ack followed by re-proposal; conflicting transition ID; broker crash after crossers durable but before install. Every installed transition must have one exact durable crossing set.

## Accepted portions

- **R3-F1 closes in shape except for broker verification.** The equal-snapshot deadlock is removed by a separate control-generation counter; spawn/adoption transitions are distinguished from epoch/state counters; the token is exactly 32 CSPRNG bytes encoded as lowercase hex and uses one inherited pipe with no second persistent copy (`:166-186`).
- **R3-F2 closes in direction except for transition fencing.** Stable per-operation IDs, pre-install m-10 crossing rows, completion updates, unknown recovery, and durable-only `recorded` terminology are the correct F64 realization (`:123-134,188-217`).
- **m-1 is now a closed owner dependency.** The exact m-1 contract SHA `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c` was pair-approved in `DESIGN-REVIEW-implementer-20260716-061153.md` and reported in `SITREP-planner-20260716-061835.md`. Refresh §E from pending to approved; m-7's consumer confirmation still remains a separate master-routed act.
- **All earlier accepted portions remain closed:** read/quarantine honesty, typed Describe, mapping split, credential custody, retry classification/fencing, event key/dedup and honest uncoupled residual, F65/F68 canonical carriers, own-process placement, complete-and-deliver, and the grill lock.
- **No operator decision is needed** if R4-F1 uses a standard OS-backed peer/lock proof and R4-F2 adds a broker transition barrier/state machine within m-10's existing store. New identity machinery or changed placement still routes back.

## Revision bar

Return fresh contract bytes and a fresh hash that:

1. Lets the broker verify controller ownership/commitment rather than trusting a token holder's generation integer.
2. Freezes one stable, transition-identified crossing set before E+1 installation and handles completion, duplicate proposal, abort, and crash recovery without false or missing rows.
3. Refreshes §E to the approved m-1 SHA while preserving consumer-confirmation authority.
4. Preserves every accepted r4 and earlier decision, the F57 ceiling, m-10 no-seat/no-conductor-secret rails, and the no-conductor-change result.

On revision, issue a fresh uniquely-parented DESIGN for exact-byte re-review. CI consumer confirmations remain held.

## Verification

- `shasum -a 256 master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md` -> `28b585856f67a4ce47397a07318d4bd0e71ac54274fa96324eb275916e3b692f`.
- Exact incoming relay `master/relays/step3-mvp-design-m7/DESIGN-planner-20260716-061547.md` lints OK; routing, parent, `DESIGN_DOC_ID`, and grill lock match.
- Checked ratified F64 at `master/STEP-3-MVP-AMENDMENT.md:92-93,113` and current m-10 IPC/store contract.
- Verified exact-byte m-1 approval at `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing contract/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file relay-lint.py exit 0
Next requested action: m-7.planner revises only R4-F1 and R4-F2, refreshes the approved m-1 edge, re-hashes the contract, and sends fresh uniquely-parented final bytes for pair re-review; CI consumer confirmations remain held.
