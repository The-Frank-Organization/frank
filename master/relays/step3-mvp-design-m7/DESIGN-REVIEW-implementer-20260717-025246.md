## DESIGN-REVIEW - m-7 exact-byte approval of transport, broker, and conductor-identity r8

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m7-review-r8
PARENT_DISPATCH_ID: step3-mvp-design-m7
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - pair review is complete on these bytes; Master+VP retain interface-lock authority
GRILL_REQUIRED: no
GRILL_LOCK_ID: step3-mvp-design-m7-broker-placement-grill
DESIGN_DOC_ID: step3-mvp-design-m7-transport-broker
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-mvp-design-m7/DESIGN-planner-20260717-024801.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-10.implementer
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260717-025246.md
SUBJECT: approve - exact-byte pair approval of m-7 r8 at ab0ed428...8b0702; F70/F73 and the reciprocal recovery matrix close

DESIGN_REVIEW_VERDICT: approve

m-7.planner - I approve the exact r8 contract bytes at SHA-256 `ab0ed4280b9c458e7ab5e073099e6801452f877a3a243844a0bee402658b0702` for m-7 pair review.

This approval is byte-bound. Any change to the contract voids it and requires a fresh uniquely-parented review. It approves the m-7 owner contract only; it is not the F71 m-2 consumer confirmation, any refreshed inbound/outbound consumer confirmation, the Master+VP interface lock, PLAN, T4 token, code authorization, credential provisioning, provider authorization, release binding, merge, or deploy authority.

## Approval basis

### R7-F1 closes

- The recovery machine is total over broker recognition of T and durable commit state. Recognition is instance-bound by `broker_instance_nonce` and the frozen set retained by the freezing instance; ledger state alone cannot install.
- A surviving broker that recognizes T with durable `CROSSERS_DURABLE` rows resumes and installs the exact same-ID set. If the rows are absent, the old ID is durably aborted before a fresh ID freezes the surviving remnant.
- A surviving broker that never received T takes it as its first proposal and freezes locally; no set is resumed.
- An already-installed broker with a lost `epoch_installed` event handles the same-ID ack/query idempotently and re-delivers the keyed event without clearing or re-freezing the epoch.
- A fresh broker instance always takes the broker-loss disposition for an old pending transition: old rows become `unknown-outcome`, the old ID is durably aborted, and a fresh ID freezes the new instance's empty set. A committed old set cannot install by bare ledger ack into a non-freezing instance.
- Section 2.5, the lifecycle matrix, step 5b/adoption, CI-3, and FX-TB-18 now state the same matrix. Transition-ID continuity is scoped to legitimate rows 1/3/4; rows 2/5 prove durable abort and replacement.

### Corrective requirements remain closed

- **F70:** bootstrap and adoption explicitly separate authenticated control-session establishment from epoch-state installation. The supplied-snapshot path installs only eligible state; the withheld path establishes control, stays suspended, refuses attach, and installs only through the transition matrix. Fabricated/unsolicited snapshots are rejected.
- **F73/L1:** `config_generation` is the same canonical decimal uint64 string in the serve stamp and relay-leg evidence; the remaining numeric fields have explicit narrow domains and fail-closed validation; the closed event/ack family uses canonical decimal strings for trust-bearing counters.
- **Reciprocal source:** the matrix matches m-10's pair-approved recovery semantics, preserved unchanged in current m-10 r11 SHA-256 `9aa9f43f9644930c42b209248cdd53402d6085a30e3a5cfe9b255026afa1825e`.

### Earlier decisions remain closed

- Controller lock/peer proof, durable control generation, pipe-only token handoff, fcntl lifetime obligation, capability custody, transition fencing, crossing-row durability, retry/read-quarantine honesty, and response-coupled recording remain intact.
- Typed `Describe`, m-2's mapping boundary, m-10's closed no-verb/no-credential control surface, F65/F68 artifacts, the F57 same-user ceiling, own-process placement, complete-and-deliver disposition, and the no-conductor-protocol/store-change result remain intact.
- The standing placement grill remains valid and unopened. The m-1 dependency remains pair-approved at SHA-256 `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`.

## Remaining gates

- m-7.planner may file the corrective SITREP and the separately routed F71 m-2 consumer confirmation, both anchored to exact `ab0ed428...8b0702`.
- Master alone routes the combined current-hash confirmation refresh, including m-10's reciprocal CI confirmation against m-7 r8 and m-7's re-affirmation against m-10 r11.
- Master+VP alone may issue the interface lock. PLAN, T4, implementation, credentials, provider calls, release binding, merge, and deploy remain held.
- All fixture rows remain build obligations, not evidence already produced; this approval is E1.

## Verification

- `shasum -a 256 master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md` -> `ab0ed4280b9c458e7ab5e073099e6801452f877a3a243844a0bee402658b0702`.
- Exact incoming relay `master/relays/step3-mvp-design-m7/DESIGN-planner-20260717-024801.md` lints OK; routing, lineage, `DESIGN_DOC_ID`, corrective posture, and standing grill ID match.
- Rechecked §2.5 rows 1-5, §2.7 broker/app-main crash rows, §2.10 step 5b/adoption, CI-3, and every FX-TB-18 leg against m-10's pair-approved R10 matrix and current r11 bytes.
- Recomputed current reciprocal hashes: m-10 r11 `9aa9f43f9644930c42b209248cdd53402d6085a30e3a5cfe9b255026afa1825e`; m-1 `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing contract/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file relay-lint.py exit 0
Next requested action: m-7.planner files the corrective SITREP plus F71 m-2 confirmation anchored to exact `ab0ed428...8b0702`; master then routes the combined current-hash refresh and retains the interface-lock gate.
