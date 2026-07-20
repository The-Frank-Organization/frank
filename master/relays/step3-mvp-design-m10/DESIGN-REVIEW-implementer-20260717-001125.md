## DESIGN-REVIEW - m-10 stage-1 IPC/manifest seam r9 must revise: adopted-broker recovery reuses an uncommitted transition ID

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r10
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - the finding is a bounded recovery-state correction under the locked transition disposition; Master+VP retain interface-lock authority
GRILL_REQUIRED: no - this re-review does not reopen the operator-locked topology, threat ceiling, F59, F60, or revocation-first choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260717-001500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-001125.md
SUBJECT: must-revise exact 424c78dc... - control/install separation closes the snapshot bypass, but adopted-broker recovery contradicts the committed-vs-uncommitted transition rule

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner - I reviewed the exact revised bytes at SHA-256 `424c78dc76c1406fe419a03612c146d4c35a5866d53c7eda088aa2ee0c89b587`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, both amendments, and both consumed sibling hashes pass. R9-F1's control-session/snapshot-install separation is present, the broker stays suspended while a pending transition is reconciled, and the ordinary E=1/wash-out paths remain separate. One new sentence in the adopted-broker subcase now contradicts the exact recovery disposition shared by m-10 §B.5 and m-7 §2.5.

## Finding

### R10-F1 - Adopted-broker recovery advances a not-committed old ID instead of aborting it

The revised common suffix treats every non-terminal adopted-broker state alike: “re-propose the SAME ID” and then make its crossing rows durable, ack, and install (`2026-07-16-mvp-ipc-manifest-seam-contract.md:87`).

That is correct only for the committed half (`CROSSERS_DURABLE`): re-ack the exact durable set under the same ID. It is explicitly wrong for this other durable crash cut:

1. m-10 proposes transition T and the surviving broker enters PREPARING and freezes its E crossing set.
2. m-10 receives that set and durably marks T `PREPARING`, but crashes before the crossing-row transaction commits (`m-10 :112-113`).
3. On adoption, the broker still holds the PREPARING barrier and T's frozen set, while m-10's ledger proves **not committed** because no `CROSSERS_DURABLE` rows exist.

Both governing contracts pin that cut to **durably ABORT T, then propose a fresh ID** for the still-current E+1: m-10 §B.5 says not committed ⇒ abort every row/state of that ID before a fresh proposal (`:115`); m-7 says the successor transition must re-freeze the surviving in-flight remnant for continuous coverage (`2026-07-16-step3-mvp-transport-broker.md:141-145`), and FX-TB-17 names the before-commit half verbatim (`:354`). Advancing T from PREPARING to CROSSERS_DURABLE after recovery erases the required distinction between crash-before-commit and crash-after-commit.

Required return: make the adopted-broker branch total over both broker recognition and durable commit state:

- broker has no transition T because the crash preceded proposal receipt ⇒ send the already-allocated T as its first proposal;
- broker recognizes T and m-10 has `CROSSERS_DURABLE` rows ⇒ resume T and re-ack the exact committed frozen set;
- broker recognizes T but m-10 has no `CROSSERS_DURABLE` rows (`PROPOSED`/`PREPARING`) ⇒ durably ABORT T, then propose a fresh ID for the same E+1 and freeze the surviving remnant;
- ledger says `CROSSERS_DURABLE` but the broker is already INSTALLED because the install event was lost ⇒ pin the idempotent same-ID ack/query behavior and durable `epoch_installed` re-delivery as an exact m-7 confirmation requirement; and
- fresh-broker loss, already-`ABORTED`, already-`INSTALLED`, initial E=1, and same-epoch wash-out retain their current distinct dispositions.

No branch may install a not-committed old transition ID, re-mint E+1, or install an epoch snapshot before the transition disposition is durable.

## What closes from review r9

- The CI-1 snapshot bypass closes in the m-10 contract: authenticated control establishment is separate from epoch-state installation; pending transitions receive no install-eligible snapshot; the broker remains suspended until the transition disposition.
- The `INSTALLED` ledger transition is now tied to a durably committed, CI-3-keyed `epoch_installed` event row rather than an unrecorded inference. The lost-event replay leg still needs the exact consumer behavior named in R10-F1.
- The pending-transition order is explicitly routed to m-7 as a consumer-confirmation delta.
- R8's no-ID and state-sensitive F59 parking repairs, role-exact generation reveal, spawn-only child recovery, connector READY gate, manifest/F63, DATA-P pairing, counter storage, and prior closed findings remain present.

## Gate disposition

This verdict is byte-bound to `424c78dc76c1406fe419a03612c146d4c35a5866d53c7eda088aa2ee0c89b587`. Any revision requires a new SHA and fresh uniquely-parented DESIGN-REVIEW. Consumer confirmations, the Master+VP interface lock, stage-3/5 designs and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `3d8bc5558b1b3e8db9a18474ff73949e6176028b0f39dc68fbd26b75b3224a03`.
- Exact design SHA-256 recomputed: `424c78dc76c1406fe419a03612c146d4c35a5866d53c7eda088aa2ee0c89b587`.
- Ratified MVP amendment SHA-256 recomputed: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Ratified architecture amendment SHA-256 recomputed: `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`.
- Pair-approved m-7 r6 SHA-256 recomputed: `f072bd996da0c85b1be9b67fad880e7395ab03de7142cb87fc6864d4f67a100e`.
- Pair-approved m-3 r2 SHA-256 recomputed: `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44`.
- Incoming DESIGN relay exact-file lint: OK.
- Exact recovery-cut sweep: matrix/common suffix at m-10 `:82-87`; PREPARING/CROSSERS/recovery split at m-10 `:108-117`; m-7 transition and reconciliation at `:123-149`; m-7 FX-TB-17 at `:354`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`
RELAY_LINT: OK - exact-file mode exit 0
Next requested action: m-10.planner folds R10-F1 into fresh design bytes, preserves the pending-transition snapshot gate, aligns adopted-broker recovery with the committed-versus-not-committed rule, recomputes the SHA-256, files a new uniquely-parented DESIGN relay, and requests fresh m-10.implementer review; do not route consumer confirmations on the current bytes.
