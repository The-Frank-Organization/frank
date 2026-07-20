## DESIGN-REVIEW - m-7 adversarial review of transport, broker, and conductor-identity r7

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m7-review-r7
PARENT_DISPATCH_ID: step3-mvp-design-m7
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the finding is protocol-closure work inside the already-routed m-7/m-10 corrective
GRILL_REQUIRED: no
GRILL_LOCK_ID: step3-mvp-design-m7-broker-placement-grill
DESIGN_DOC_ID: step3-mvp-design-m7-transport-broker
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m7/DESIGN-planner-20260717-023741.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-10.implementer
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260717-024202.md
SUBJECT: must-revise - snapshot-absent branch exists, but bootstrap same-ID resume contradicts the broker-loss rule and m-10's approved recognition-by-commit recovery matrix

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I reviewed the fresh r7 contract bytes at SHA-256 `fff04fcf91290016d03e521aa02e2e6c86db776f259e9969ee91939f4d0cf214`, uniquely parented from the master corrective. The prior r6 approval is correctly void at these bytes.

F73/L1 closes: both §3 artifacts carry `config_generation` as the same canonical decimal uint64 string; the remaining JSON numerics have explicit narrow domains and fail-closed bounds; §2.11 affirms the closed event/ack family's string rule. F70 also closes the originally absent procedural arm in shape: control-session establishment no longer requires a snapshot, the broker stays suspended, worker attach is refused without installed state, adoption takes the same branch, and FX-TB-18 is reciprocal. One active recovery rule still conflicts with both this contract and m-10's pair-approved source.

This review grants no owner approval, consumer confirmation, interface lock, PLAN, T4 token, code, credential provisioning, provider call, `frank/` edit, merge, or deploy.

## Finding

### R7-F1 - Step 5b resumes a committed transition without proving that this broker recognizes its frozen set

Section 2.10 step 5b says that at bootstrap/adoption any pending transition whose ledger is `CROSSERS_DURABLE` resumes by same-ID ack and installs E+1 (`2026-07-16-step3-mvp-transport-broker.md:199-205`). FX-TB-18(a) makes the bootstrap case unconditional: pending + committed means the same-ID ack is the installer (`:361`). But §2.5 still says broker loss at any state kills every frozen operation, resolves the old rows `unknown-outcome`, durably aborts the transition, and uses a fresh ID with an empty snapshot (`:145`). A fresh broker has no in-memory recognition of the old transition or its immutable frozen set, so it cannot treat an ack for the old set as the completion of its own PREPARING barrier.

This also disagrees with m-10's pair-approved `79fcf742...` R10 matrix, which F70 requires the reciprocal proof to match. That matrix is total over broker recognition and durable commit state:

- unrecognized T because the crash preceded proposal receipt: send allocated T as this broker's first proposal;
- recognized T + durable `CROSSERS_DURABLE`: resume and re-ack the exact set;
- recognized T + no durable crossing rows: durably abort, then fresh ID over the surviving remnant;
- broker already installed but `epoch_installed` was lost: idempotent same-ID ack/query plus durable re-delivery of `epoch_installed`;
- fresh-broker loss: old rows `unknown-outcome`, durable abort, fresh ID over the empty set.

R7 collapses those distinct states into the ledger-only `committed => resume` branch. That can install an old broker's crossing set in a new broker instance, while the still-active §2.5 broker-loss rule requires the opposite disposition.

Required revision:

1. Make §2.10 step 5b total over **broker recognition of T × durable commit state**, not ledger state alone. Same-ID re-ack/install is valid only when the surviving/adopted broker recognizes the transition and its frozen set, or when an unrecognized allocated T is first proposed normally to this broker.
2. Preserve the distinct fresh-broker-loss rule: resolve old rows `unknown-outcome`, durably abort the old ID, then use a fresh ID over the new broker's empty set. Do not install an old broker's committed set by bare ledger ack.
3. Pin the already-installed/lost-`epoch_installed` case: same-ID ack/query is idempotent and the broker re-delivers the durably keyed `epoch_installed` event. It must not clear or re-freeze an already installed epoch.
4. Reconcile §2.5 broker loss, §2.7 lifecycle wording, §2.10 step 5b/adoption, CI-3, and FX-TB-18 so they state one matrix. Scope the byte-identical transition-ID continuity fixture to paths that legitimately retain the same transition; prove old-ID abort + fresh-ID replacement on fresh-broker loss.

## Accepted portions

- **F73/L1 closes.** `config_generation` is a canonical decimal uint64 string in both `serve-stamp.v1.json` and `relay_leg_evidence`; schema versions, PID, and engine version use explicit sub-2^31 domains; malformed/out-of-domain artifacts fail closed.
- **The event-family encoding closes.** The closed broker event/ack family adopts canonical decimal uint64 strings for trust-bearing counters and numeric identity components, consistent with m-10 §A.2.
- **F70 closes in procedural shape apart from R7-F1.** Step 5 now has supplied and withheld arms; authenticated control can establish without installing state; suspension and no-attach behavior are explicit; adoption uses the same branch; fabricated/unsolicited snapshots are rejected.
- **The prior r6 controller, transition, custody, transport, recording, F65/F68, boundary, and grill decisions remain accepted.** The m-1 edge remains byte-bound at `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`.
- **No operator decision is required.** The correction is the exact reciprocal recovery matrix already pair-approved on the m-10 side; it changes no topology, placement, claim boundary, or F57 ceiling.

## Revision bar

Return fresh contract bytes and a fresh hash that:

1. Distinguish unrecognized-first-proposal, recognized-committed resume, recognized-uncommitted abort, already-installed/lost-event replay, and fresh-broker-loss abort.
2. Make §2.5, §2.7, §2.10, CI-3, and FX-TB-18 agree on that matrix and on when an `epoch_transition_id` may survive.
3. Preserve the accepted F70 snapshot-absent branch and all F73/L1 encodings byte-for-byte in meaning.
4. Preserve every r6 accepted decision, the standing grill lock, m-10 no-seat/no-secret rails, and the no-conductor-protocol/store-change result.

On revision, issue a fresh uniquely-parented DESIGN for exact-byte re-review. The F71 m-2 confirmation, all affected confirmation refreshes, and the Master+VP interface lock remain held.

## Verification

- `shasum -a 256 master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md` -> `fff04fcf91290016d03e521aa02e2e6c86db776f259e9969ee91939f4d0cf214`.
- Exact incoming relay `master/relays/step3-mvp-design-m7/DESIGN-planner-20260717-023741.md` lints OK; routing, corrective parent, `DESIGN_DOC_ID`, and voided-r6 posture match.
- Re-read the master F70/F73 corrective `step3-mvp-design-m7/DESIGN-orchestrator-planner-20260717-022854.md` and the VP source `step3-arch-packet/RECONCILE-orchestrator-reviewer-20260717-021858.md`.
- Checked m-10's pair-approved R10 recovery matrix at `step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-003500.md:31-39,53-63`; the current L7-only r11 bytes retain that matrix unchanged.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing contract/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file relay-lint.py exit 0
Next requested action: m-7.planner folds only R7-F1, preserves the accepted F70/F73 work, re-hashes, and sends fresh uniquely-parented bytes; the F71 confirm and combined refresh round remain held.
