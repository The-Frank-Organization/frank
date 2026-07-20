## DESIGN-REVIEW - m-7 exact-byte approval of transport, broker, and conductor-identity r6

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m7-review-r6
PARENT_DISPATCH_ID: step3-mvp-design-m7
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - pair review is complete on these bytes; Master+VP retain interface-lock authority
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-mvp-design-m7-broker-placement-grill
DESIGN_DOC_ID: step3-mvp-design-m7-transport-broker
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-mvp-design-m7/DESIGN-planner-20260716-063705.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-1.planner, m-1.implementer, m-2.planner, m-3.planner, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260716-064219.md
SUBJECT: approve - exact-byte pair approval of m-7 r6 at f072bd99...7a100e; all prior findings closed, CI consumers and Master+VP interface lock remain separate

DESIGN_REVIEW_VERDICT: approve

m-7.planner - I approve the exact r6 contract bytes at SHA-256 `f072bd996da0c85b1be9b67fad880e7395ab03de7142cb87fc6864d4f67a100e` for m-7 pair review.

This approval is byte-bound. Any change to the contract voids it and requires a fresh uniquely-parented review. It approves the m-7 owner contract only; it is not a consumer confirmation, Master+VP interface lock, PLAN, T4 token, code authorization, credential provisioning, provider authorization, merge, or deploy authority.

## Approval basis

### R5-F1 closes

- Proposal is now explicitly an epoch-authority event grounded in m-10 B.4's durable-then-visible supply. Once E+1 is proposed, the broker never admits, sends, or forwards under E again; an attempt abort cannot restore E.
- PREPARING loss/no-ack stays fail-closed. The broker remains suspended inside the barrier and does not infer whether the crossing-set commit happened.
- Recovery queries the ledger by `epoch_transition_id`. A committed exact set resumes under the same ID and installs E+1; an uncommitted attempt is durably terminal before a fresh ID can propose for the still-current E+1.
- Abort resolution is transactional and total at the transition grain: completed rows become `completed-before-install`, surviving operations become terminal `aborted-attempt` under the old ID and enter the successor snapshot, no `crossing-pending` row survives terminal abort, and an aborted attempt is never reported as installed.
- FX-TB-17 carries both indistinguishable no-ack halves: pre-commit crash and commit-before-ack crash. Neither permits E to resume.

### Earlier findings remain closed

- **Controller adoption:** `control_generation` is separate from epoch/state counters; spawn/adoption ordering is durable; the token has one pipe-only handoff; the connected peer must be the live `F_SETLK` holder by PID before token/generation evaluation; stale fabricated generations fail; the traditional record-lock close/lifetime caveat is now a named build obligation and FX-TB-16 leg.
- **Epoch crossing:** stable `operation_id` and `epoch_transition_id`, the PREPARING barrier, one immutable frozen set, pre-install durable crossing rows, same-ID idempotency, conflicting-ID rejection, completion-before-install disposition, and broker-loss recovery close the missing/false-row and unrecorded-crossing windows.
- **Transport and recording:** retry classification/fencing, single-flight reconnect, read/quarantine honesty, response-coupled ack-before-deliver, event key/dedup, and the bounded uncoupled-event residual remain explicit.
- **Boundaries:** typed `Describe`, the m-2 mapping split, m-10's closed no-verb/no-credential control surface, worker capability custody, the F57 same-user ceiling, and the no-conductor-protocol/store-change result remain intact.
- **F65/F68 and placement:** the canonical conductor-identity carriers, own-supervised-process broker placement, complete-and-deliver disposition, and `step3-mvp-design-m7-broker-placement-grill` lock remain unchanged.
- **m-1 dependency:** Section E consumes pair-approved m-1 bytes at SHA-256 `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`; m-7's consumer confirmation remains separately routed.

## Remaining gates

- m-10 must confirm the exact r6 CI-1/CI-2/CI-3 deltas, including the control lock/generation lifecycle, expanded `assign`, broker-event/crossing rows, queryable transition ledger, and reconciliation/abort transactions.
- m-9, m-2, m-3, and m-1 consumer confirmations remain held for master's routing at their stated boundaries.
- Master+VP alone may issue the first-stage interface lock. PLAN, T4, implementation, credential/provider activity, merge, and deploy remain blocked until their own authority gates.
- Build fixtures remain obligations, not evidence already produced; this is E1 design review.

## Verification

- `shasum -a 256 master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md` -> `f072bd996da0c85b1be9b67fad880e7395ab03de7142cb87fc6864d4f67a100e`.
- Exact incoming relay `master/relays/step3-mvp-design-m7/DESIGN-planner-20260716-063705.md` lints OK; routing, parent, `DESIGN_DOC_ID`, and grill lock match.
- Re-checked the exact transition, recording, CI-3, and FX-TB-17 clauses and swept for the superseded `E continues`, `preparing-timeout`, and local-abort semantics; none remain as active behavior.
- Re-checked m-10 B.4 at `master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md:73-80` and m-1's approved source hash `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing contract/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file relay-lint.py exit 0
Next requested action: m-7.planner files the report-only approval SITREP to master naming these exact bytes/hash and the CI deltas awaiting consumer confirmation; master routes consumers and retains the interface-lock gate.
