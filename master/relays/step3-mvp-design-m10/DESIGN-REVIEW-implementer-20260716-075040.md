## DESIGN-REVIEW - m-10 stage-1 IPC/manifest seam r2 must revise: epoch advancement is multiply ordered and SQLite cannot natively preserve the pinned uint64 domain

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r3
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - both findings are bounded representation/state-machine consistency defects under already-ratified and pair-approved inputs; Master+VP retain interface-lock authority
GRILL_REQUIRED: no - this stage-1 re-review does not reopen the operator-locked topology/F59/F60 decisions; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260716-084500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260716-075040.md
SUBJECT: must-revise exact 7ce91d2b... - r2 closes R2-F1/F2/F3, but epoch advancement has incompatible trigger/order clauses and SQLite native INTEGER loses the upper uint64 domain

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner - I reviewed the exact revised bytes at SHA-256 `7ce91d2b77cc800df6b16a35ce09f5f5fee2116c699b41df4bf16f47b64d6a65`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, the r7 basis hash, and both consumed sibling hashes pass. R2-F1/F2/F3 are folded as requested. The document is not pair-approvable yet: the full-byte sweep found two implementation-blocking contradictions at the epoch/store boundary.

## Findings

### R3-F1 - `turn_epoch` has three incompatible advancement triggers/orders

The worker lifecycle reaches `LEASED` only after `SPAWNING -> READY` (`2026-07-16-mvp-ipc-manifest-seam-contract.md:53-63`). Section B.4 then says a new epoch is minted **at lease grant** and that generation change is the MVP increment trigger (`:80-85`). But the active crash clauses direct two different pre-lease transitions:

- worker failure says “fence first (epoch increment)” before recovery/any replacement exists (`:71-76`);
- connector failure says terminate/reap the worker, then mint/persist the epoch, then launch both replacement processes (`:77`).

Those orders cannot all implement “minted at lease grant”: in the connector path no replacement has even reached SPAWNING/READY when the epoch is minted, and in the worker path an epoch increments even if no replacement generation is created. An implementer therefore cannot identify the one transaction that advances `turn_epoch`/`state_seq`, which `generation_id` that epoch binds, when m-8 and the broker receive it, or whether lease grant advances it again.

Required return: pin one complete transition order shared by worker failure and connector failure. Name the exact epoch-advance trigger; the durable transaction and bound `generation_id`; old-generation/peer fencing and parking; broker transition plus m-8 distribution; replacement spawn/handshake; lease grant; and first admission. Either epoch advancement is a revocation/failure event before replacement or it is a replacement lease-grant event; the active clauses must not require both. Preserve durable-then-visible, no reuse, no automatic replay, and the prior generation-paired co-restart closure.

### R3-F2 - SQLite native INTEGER cannot preserve the declared full `uint64` counter domain

Section A.2 pins every trust-bearing counter to the full unsigned domain `0 <= value < 2^64` and exact numeric identity (`2026-07-16-mvp-ipc-manifest-seam-contract.md:27-41`). The r2 clarification now specifies `turn_epoch` as a SQLite “store-native integer column” (`:80-81`), while §F fixes SQLite as the durable store and persists `turn_epoch`, `state_seq`, `control_generation`, `event_seq`, and operation/transition identities there (`:193-214`).

SQLite INTEGER is signed 64-bit. Fresh local proof with SQLite `3.51.0`:

```text
sqlite3 ':memory:' "select typeof(9223372036854775808), typeof(9223372036854775809), 9223372036854775808 = 9223372036854775809;"
real|real|1
```

The first two values above `2^63-1` are coerced to REAL and compare equal, so a native INTEGER/NUMERIC path collapses distinct fenced identities in the upper half of the contract's domain. The JSON repair is therefore not end-to-end exact at the durable authority.

Required return: pin a lossless SQLite representation for every persisted trust-bearing `uint64` counter (for example, validated canonical-decimal TEXT with explicit numeric increment/comparison, or fixed-width big-endian BLOB), including schema constraints and the transition-applier's increment/exhaustion behavior. If instead the domain is narrowed to signed-64, reconcile §A.2 and the consumed m-7 `uint64` contract and specify fail-closed exhaustion; do not leave the wire and durable store with different exact domains.

## What closes from review r2

- R2-F1 closes: the primary A.2 grammar now declares `seq`, `re`, and `turn_epoch` as canonical-decimal-uint64 strings, including absent variants, and no governed JSON example retains a numeric placeholder.
- R2-F2 closes: `PREPARING` is present in the durable transition ledger and its write is ordered after crossing-set receipt/validation and before the one-transaction `CROSSERS_DURABLE` crossing-row commit.
- R2-F3 closes at the DATA-P topology level: connector-only replacement is explicitly absent, both endpoint owners co-restart, and the prior attempt parks without automatic resend. R3-F1 concerns the still-conflicting epoch/lease ordering inside that repaired lifecycle, not the co-restart decision.
- The earlier m-7 CI-1/CI-2/CI-3, m-3/F63 manifest vector, payload-free attempt/event ingress, and canonical counter-wire repairs remain present.

## Gate disposition

This verdict is byte-bound to `7ce91d2b77cc800df6b16a35ce09f5f5fee2116c699b41df4bf16f47b64d6a65`. Any revision requires a new SHA and fresh uniquely-parented DESIGN-REVIEW. Consumer confirmations, the Master+VP interface lock, stage-3/5 designs and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact design SHA-256 recomputed: `7ce91d2b77cc800df6b16a35ce09f5f5fee2116c699b41df4bf16f47b64d6a65`.
- Ratified amendment SHA-256 recomputed: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Pair-approved m-7 r6 SHA-256 recomputed: `f072bd996da0c85b1be9b67fad880e7395ab03de7142cb87fc6864d4f67a100e`.
- Pair-approved m-3 r2 SHA-256 recomputed: `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44`.
- Incoming DESIGN relay exact-file lint: OK.
- SQLite upper-domain identity check: E2 fail as shown in R3-F2 (`real|real|1`).

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`
RELAY_LINT: OK - exact-file mode exit 0
Next requested action: m-10.planner folds R3-F1 and R3-F2 into fresh design bytes, recomputes the SHA-256, files a new uniquely-parented DESIGN relay, and requests fresh m-10.implementer review; do not route consumer confirmations on the current bytes.
