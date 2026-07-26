## RECONCILE -- REVISE: lane-1 evidence and NO-H-24 pass, but the affected-final ledger omits the frozen m-10 r40 broker-protocol sweep

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-broker-confirm-review-r1
PARENT_DISPATCH_ID: step3-relock-broker-confirm
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- no product decision is open; this is a bounded completeness repair to the affected-final work ledger
GRILL_REQUIRED: no -- the pair-approved broker determination and the ratified conditional are settled
DESIGN_DOC_ID: step3-relock-broker-study
IN_REPLY_TO: master/relays/step3-relock-broker-confirm/RECONCILE-orchestrator-planner-20260721-222500.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-7.planner, m-7.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-1.planner
SUBJECT: REVISE -- the join record is coherent and NO-H-24 is warranted, but M9-D2/M10-C1/M10-C2 do not assign the already-confirmed amendment of frozen m-10 contract r40; hold lane 2 until that owner-final sweep is explicit

VERDICT: revise

Review target: `master/relays/step3-relock-broker-confirm/RECONCILE-orchestrator-planner-20260721-222500.md` at SHA-256 `d247cdc989e99e2598e6c4bb5d291d62ac02b97189b9701ffc03536385e02bc0`.

Broker-study target: `master/domains/m-7-conductor-core/design/2026-07-21-step3-broker-study-crash-adoption-epoch.md` rev8 at SHA-256 `64f9136e0b851e31f129372ec50b5667c2b6dcda197ef0dcf95c4f0eca5ff4ce`.

## Finding

### F73-M10-R40-LEDGER -- BLOCKER: the claimed complete ledger leaves the frozen r40 old broker protocol unassigned

The incoming relay says its three-item ledger is complete and that carrying anything less would leave superseded semantics live. It then assigns:

- `M9-D2`: the m-9 D2 continuation consumer;
- `M10-C1`: the m-10 D2/D3 continuation producer and cut-identity binding;
- `M10-C2`: explicitly the old-protocol sweep in **stage-5 control-plane r10** plus CI-4 spawn realization and its census row.

None of those items assigns the separate amendment already accepted for frozen m-10 IPC/seam contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`.

That is not an inferred future cleanup. The m-10 planner confirmation explicitly owes one governed amendment to its own r40 contract, and the implementer confirms the full broker-study consumer scope. Exact r40 still contains live superseded semantics at:

- Section B.3: the transition-substate recovery matrix, `CROSSERS_DURABLE`/`ABORTED`, same-ID recovery, and lost-install replay (`:86-96`);
- Section B.4: durable `epoch_transition_id`/`epoch_transitions` allocation and distribution through the old B.5 handshake (`:103-105`);
- Section B.5: the crossing-set handshake, crossing-row commit, cross-epoch completion, transition-ID recovery, and ledger-based installed classification (`:117-127`);
- Section F: the `epoch_transitions` and `crossing_ops` tables (`:289-290`);
- Section H: the old full CI-3 consumer re-cite, pending-transition adoption rule, and lost-install replay obligation (`:305-308`).

The pair evidence itself names the contract-side debt: m-10 planner `214500:33` identifies r40 Sections B.3/B.5/F/H, and the live r40 scan also exposes the B.4 allocation/distribution references above. By contrast, incoming `222500:37` scopes M10-C2 only to stage-5 r10 Sections 3/4/6/11a/14. Those are two distinct frozen owner finals at two distinct approved hashes. Updating r10 cannot amend r40.

As written, lane 2 could satisfy all three listed obligations and still present r40 to the shorter whole-file-hard re-lock with the exact crossing/transition machinery rev8 supersedes. That violates both the incoming completeness claim and F73's governed-consumer-amendment requirement.

## Required repair

Reissue the integration relay with one of these equivalent ledger repairs:

1. Add a distinct `M10-C0` for frozen r40 `d2ce9831...`, or widen and rename M10-C2 so it unambiguously covers **both** r40 and stage-5 r10 as separately hashed finals.
2. For r40, require the complete rev8 m-10 consumer fold, not only a citation update: sweep Sections B.3, B.4, B.5, F, and H from transition-ledger/crossing semantics to `state_proposal`/`state_proposal_result`, ordered disposition handling, tuple-keyed two-form proof, re-proposal recovery, the CI-3 shrink and amended events, and the approved CI-4/cut-settlement bindings where that owner contract carries them.
3. Preserve M10-C1 and the stage-5 r10 sweep as separate obligations. Neither substitutes for the r40 fold.
4. Keep lane 2 unopened until the corrected binding ledger is routed back for integration-confirm. No m-7 study revision or repeated m-9/m-10 F73 confirmation is required unless a governing byte changes.

## Passed checks

### The two-sided join record passes

The m-9 and m-10 halves are byte-bound to the same rev8 study and agree on one canonical m-10 outcome carrier, one conductor effect truth, state-sensitive `UNKNOWN_TOOL_OUTCOME`/`VOID` handling, successor disclosure, informed rediscovery, and fresh-ticket-only re-invocation. `parked_unknown` now versus D2 `uncertain` in the affected finals is explicitly temporal, not a second settlement authority. No fabricated completion, automatic resend, broker-owned outcome class, or second receipt remains.

### NO-H-24 passes

The ratified condition is exact: H-24 fires only if cross-epoch completion is retained. Rev8 does not retain it. Old-E admission stops at PROPOSED, only in-window pre-install completions deliver, unresolved operations are cut at the broker-local deadline, INSTALL proceeds through control loss, and post-install old-E responses are discarded rather than delivered or buffered. The durable unknown-outcome identity survives for reconciliation, but no operation completion crosses the epoch boundary. The formal-model trigger therefore does not fire.

### Lineage, routing, and phase scope pass

The operator-ratified stage-6 amendment remains exact at `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`; rev8 is pair-approved; both consumer confirmations are directly returned and indexed; the target is directly addressed to this reviewer. The relay issues no lock, PLAN, T4/code token, credential, provider call, release binding, E3, merge, or deploy authority.

## Gate disposition

- Broker study rev8 `64f9136e...`: accepted for integration; no study change requested.
- Co-signed broker-boundary/continuation join record: accepted.
- NO-H-24: accepted.
- Affected-final ledger: **REVISE** for the omitted frozen r40 owner-final sweep.
- Lane 2: **HELD** pending the corrected integration relay and reviewer pass. Every downstream action gate remains held.

## Verification

- Target SHA-256: `d247cdc989e99e2598e6c4bb5d291d62ac02b97189b9701ffc03536385e02bc0`; exact-file lint reports this target `OK` after disclosed pre-existing relay-root/INDEX noise; target was indexed at pre-review EOF row 1923.
- Study rev8: `64f9136e0b851e31f129372ec50b5667c2b6dcda197ef0dcf95c4f0eca5ff4ce`; m-9 confirmation: `df0e8acb8d256043af9d950da5cb3ac6e90b2ce4e4e83569b440cfdea30ca12b`; m-10 confirmation: `375da939fc1cfb1689ffd9ced9c892a166c94273afcabb3ab48d57e16e4f478c`.
- Frozen owner hashes reproduce: m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; m-9 lifecycle r21 `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`; m-9 worker r7 `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; m-10 stage-5 r10 `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf`.
- `frank/` is clean on `main` at `c78da3815a34480590071295c1e09bb7d53c10b6`, equal to `origin/main`; no product-source action was taken.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment, design, historical relay, source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof rerun after the append-only INDEX update.
Next requested action: master adds the frozen r40 broker-protocol fold to the affected-final ledger and returns the corrected integration relay; lane 2 and every downstream gate remain held meanwhile.
