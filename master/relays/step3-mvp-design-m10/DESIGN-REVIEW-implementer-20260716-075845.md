## DESIGN-REVIEW - m-10 stage-1 IPC/manifest seam r3 must revise: no legal non-attachable interregnum broker tuple and the padded-TEXT CHECK admits out-of-domain counters

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r4
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - both findings are bounded exact-schema/state-machine completion defects under the chosen revocation-first model; Master+VP retain interface-lock authority
GRILL_REQUIRED: no - this stage-1 re-review does not reopen the operator-locked topology/F59/F60 decisions or the planner's revocation-first choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260716-091000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260716-075845.md
SUBJECT: must-revise exact c4ab9095... - revocation-first and padded TEXT close the prior choices, but the broker has no safe interregnum generation tuple and the SQL CHECK accepts values above uint64 max

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner - I reviewed the exact revised bytes at SHA-256 `c4ab90954882a90329822e68c35f03b73d3dd9bc0ce6fd8d9bca803cee6082b7`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, the r7 basis hash, and both consumed sibling hashes pass. The revocation-first epoch model and padded-TEXT storage direction resolve R3-F1/R3-F2 at the decision level. Two exact implementation blockers remain in those new clauses.

## Findings

### R4-F1 - The revocation interregnum has no schema-valid broker tuple that is both current and non-attachable

The retirement transaction releases the old lease, mints E+1, and leaves “successor binding = pending”; distribution immediately runs the broker epoch handshake and the general increment rule publishes `epoch_update{run_id, turn_epoch, generation_id}` (`2026-07-16-mvp-ipc-manifest-seam-contract.md:80-90`). But G+1 is not introduced until replacement spawn/assign and is not lease-bound until step 5 (`:87-88`). The consumed m-7 contract admits no missing generation: its installed snapshot is exactly `epoch_state{run_id, generation_id, turn_epoch, lease_state, state_seq}`, and worker attach grants capability when the presented `{run_id, generation_id, turn_epoch}` equals that installed state (`master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md:199-205`). It does not gate attach on `lease_state`.

The current bytes therefore leave no defined safe value for `generation_id` during steps 1-4:

- old G with E+1 can make a retired worker's guessed/presented `{G,E+1}` match the broker tuple;
- G+1 is not yet durably bound and a pre-lease candidate may wash out;
- absent/sentinel generation is not in the m-7 schema.

The same new state block also writes `terminal-pending-reap`, which is absent from the authoritative §B.1 worker-state machine (`2026-07-16-mvp-ipc-manifest-seam-contract.md:53-63,84`), and ends with the stale absolute “new generation ⇒ new epoch” even though a washed-out pre-lease generation explicitly does not re-mint (`:82`). Those leave the durable worker row and retry-generation rule ambiguous.

Required return: define the exact durable/interchange representation for the no-active-generation interregnum and prove it cannot satisfy worker attach before lease grant. Pin the `generation_id`/`lease_state` values at retirement, every same-epoch update after a washed-out candidate, and the final lease-bound tuple. If safety requires broker attach to check `lease_state`, an optional generation, or a new sentinel shape, that is an m-7 exact-contract change and must route there before m-10 can confirm CI consumption. Also map `terminal-pending-reap` to an existing §B.1 state or add it to the machine/store census, and replace the stale generation-implies-epoch shorthand with the chosen retirement-implies-epoch invariant.

### R4-F2 - The pinned SQLite CHECK accepts counters above `2^64-1`

The fixed-width padded-TEXT representation is lossless, and lexicographic ordering is valid for equal-width decimal strings. But the normative constraint checks only length and digits (`2026-07-16-mvp-ipc-manifest-seam-contract.md:81`); it does not enforce the declared unsigned-64 upper bound. Fresh SQLite proof using the exact proposed CHECK:

```text
create table c(v TEXT CHECK(length(v)=20 AND v GLOB '[0-9]*' AND v NOT GLOB '*[^0-9]*'));
insert into c values('18446744073709551615');
insert into c values('18446744073709551616');
select v from c order by v;
18446744073709551615
18446744073709551616
```

The first out-of-domain value is accepted, so a recovery `MAX` may select a value that cannot legally cross §A.2. The same paragraph also calls the m-10 transition applier “the sole increment site” immediately after listing m-7-produced `event_seq` and `op_seq`; those counters are validated/persisted by m-10 but incremented by the broker, so the ownership sentence is not exact.

Required return: add the fixed-width upper-bound constraint (for example, under the pinned decimal/BINARY ordering, `col <= '18446744073709551615'`) and require the same bound validation for every insert/update, not only local increment. Scope the transition-applier increment rule to m-10-owned counters; name m-7-owned inbound counters as validate-and-store-only and leave their increment/exhaustion behavior with m-7. Preserve wire-unpadded/store-padded conversion and fail-closed local exhaustion.

## What closes from review r3

- R3-F1 closes at the governing-choice level: epoch advancement is now revocation-first, happens only at run admission or active-generation retirement, and lease grant binds without incrementing.
- Worker failure, connector failure, and requested replacement now share one park/fence/reap/spawn/lease sequence; connector-only restart remains absent.
- R3-F2 closes at the representation-choice level: persisted counters no longer use SQLite INTEGER/REAL, and 20-digit padded decimal TEXT preserves exact full-domain identity and ordering. R4-F2 is the missing range constraint/ownership qualifier, not a rejection of padded TEXT.
- Prior broker CI, manifest/F63, DATA-P, PREPARING-ledger, attempt/event ingress, and canonical wire-counter repairs remain present.

## Gate disposition

This verdict is byte-bound to `c4ab90954882a90329822e68c35f03b73d3dd9bc0ce6fd8d9bca803cee6082b7`. Any revision requires a new SHA and fresh uniquely-parented DESIGN-REVIEW. Consumer confirmations, the Master+VP interface lock, stage-3/5 designs and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact design SHA-256 recomputed: `c4ab90954882a90329822e68c35f03b73d3dd9bc0ce6fd8d9bca803cee6082b7`.
- Ratified amendment SHA-256 recomputed: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Pair-approved m-7 r6 SHA-256 recomputed: `f072bd996da0c85b1be9b67fad880e7395ab03de7142cb87fc6864d4f67a100e`.
- Pair-approved m-3 r2 SHA-256 recomputed: `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44`.
- Incoming DESIGN relay exact-file lint: OK.
- Proposed SQLite CHECK upper-bound test: E2 fail as shown in R4-F2; `18446744073709551616` is accepted.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`
RELAY_LINT: OK - exact-file mode exit 0
Next requested action: m-10.planner folds R4-F1 and R4-F2 into fresh design bytes, recomputes the SHA-256, files a new uniquely-parented DESIGN relay, and requests fresh m-10.implementer review; do not route consumer confirmations on the current bytes.
