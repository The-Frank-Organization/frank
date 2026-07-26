## DESIGN-REVIEW — rev4 MUST REVISE: tuple-keyed recovery closes the T1/T2 race, but ACK-as-proof is not yet an exact wire contract and `crossing_count` is undefined

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-m7-broker-review-r4
PARENT_DISPATCH_ID: step3-relock-m7-broker
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-m7-broker-study
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-m7-broker/DESIGN-planner-20260721-175033.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-1.planner
BUNDLE_ID: m-7-conductor-core
SUBJECT: byte-exact re-review of `master/domains/m-7-conductor-core/design/2026-07-21-step3-broker-study-crash-adoption-epoch.md` rev4 at SHA-256 `94a08bed9febfb4c85f2524e36f7105aee8a4f0b7140349928a04a159c015b24` — the tuple join is sound in principle; two interface-definition residuals block approval

## Verdict

**MUST REVISE.** R3-F2 and R3-F3 close. R3-F1's architectural defect also closes in substance: the canonical trusted tuple is the correct lifecycle join, attempt correlations are telemetry-only, current-state re-proposal can recover from event loss, mismatched evidence fails closed, and the T1/T2/broker-death fixtures converge without a transition ledger.

Rev4 is not yet implementable at interface-lock grain because the new ACK proof exists only as prose (`installed{tuple}`), without a named frame/schema/direction, while Q2 and §C still retain the superseded event-only gate wording. The unchanged `epoch_installed` schema also retains `crossing_count` after all crossing semantics were removed.

## Closed From R3

- **R3-F1 CLOSED at mechanism level:** canonical-tuple proof, old/new correlation dedup, queue-loss recovery, broker-death reinstall, and mismatched-tuple rejection are coherent.
- **R3-F2 CLOSED:** full-state comparison now distinguishes equal state, same-epoch newer state, newer epoch, and stale/regressing state; the wash-out fixture protects first attach.
- **R3-F3 CLOSED:** F64 fencing remains universal while mandatory durable settlement and the intentional `Describe` recording supersession are stated separately and honestly.

## Findings

### R4-F1 — BLOCKER — ACK-as-proof has no exact control-wire contract, and two normative consumers still make `epoch_installed` the sole gate

**Current bytes:** §Q3.3.1 adds a re-proposal ACK carrying `installed{tuple}` and makes it sufficient install proof. No message type, direction, required fields, correlation rule, closed/open schema, malformed disposition, or duplicate handling is defined. Meanwhile §Q2.3 still says m-10 learns drain completion from the `epoch_installed` event, and §C still calls `epoch_installed` the assign gate. Those statements exclude the ACK form that exists specifically to survive event loss.

This is a boundary contract, not implementation detail. m-10 cannot confirm or build a proof form whose wire identity and consumption rule are unnamed, and the current exact bytes permit one consumer to wait forever for a lost event despite receiving the new proof ACK.

**Required correction:**

1. Name the exact m-10→broker re-proposal request and broker→m-10 proof response (or explicitly amend an existing named frame), with direction and closed required fields. The response must carry the current attempt correlation plus the canonical trusted tuple `{run_id, generation_id, turn_epoch, state_seq}` in one unambiguous shape.
2. Pin validation: correlation must match the live request; tuple must equal m-10's current durable tuple; malformed/mismatched/stale responses fail closed and never open assign. A duplicate equivalent proof is idempotent.
3. Replace §Q2.3 and §C's event-only wording with the exact two-form rule: either a valid tuple-matching `epoch_installed` event or the valid tuple-matching installed-state proof response opens the same logical gate.
4. Bind “exactly once” to existing durable lifecycle state/chokepoint behavior, or narrow the claim to idempotent proof consumption plus lifecycle-owned assignment. State why an app-main crash after proof but before/after assign cannot create a second logical lease/worker admission without adding a hidden transition row.
5. Add request/response loss and duplicate fixtures: proof response lost then re-proposal; duplicate equivalent response; response with wrong correlation; response with matching correlation but mismatched tuple; app-main crash after proof before assign; app-main crash after assign.

No durable transition ledger is required by this correction; it completes the no-ledger interface already chosen.

### R4-F2 — IMPORTANT — `epoch_installed.crossing_count` survives as an “unchanged” required member with no post-crossing semantics

Rev4 removes `crossing_set`, crossing rows, and every cross-epoch completion, but §Q3.3.3 says `epoch_installed` is unchanged. Frozen r11's closed schema requires `epoch_installed{epoch_transition_id, generation_id, turn_epoch, state_seq, crossing_count}`. The study never says what `crossing_count` now counts. Leaving it unconstrained can falsely imply retained crossings and gives m-10 no exact validation rule.

Choose and pin one governed result:

- preserve wire compatibility with `crossing_count` REQUIRED and canonically `0` for every clean-cut install, rejecting nonzero; or
- remove/rename the member in the amended closed schema and include that exact delta in m-10's confirmation scope.

Add a fixture asserting the chosen encoding and its negative. Do not leave an old mandatory member semantically orphaned.

## Re-review Gate

Return one byte-exact rev5 that preserves every closed r1-r3 mechanism and adds only:

- the exact proof request/response wire contract and validation table;
- consistent two-form gate wording in Q2, Q3, and §C;
- crash/duplicate/mismatch fixtures demonstrating idempotent proof consumption and lifecycle-owned assignment; and
- an exact `crossing_count` disposition plus negative fixture.

The clean-cut/no-H-24 determination remains supportable; neither finding requires crossing rows, retained completion, or a durable transition identity.

## Verification

- Reviewed rev4 at exact SHA-256 `94a08bed9febfb4c85f2524e36f7105aee8a4f0b7140349928a04a159c015b24`; incoming relay at exact SHA-256 `ad4313e1d4d57ea945151063739105da96ef843e3c8ddc54cddb07cc90f4d35c`.
- Reproduced frozen bases: m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; m-9 `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Incoming exact-file relay lint reports `OK`; root-mode output retains unrelated historical `INDEX.md`/lineage noise.
- No design or code bytes were modified by this review.

ACTIONS_GIT_REF: none — review-only docs action; no `frank/` implementation authority exercised.
FINAL_GIT_STATUS_SHORT: none — `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6` after this relay and INDEX row were written.
Next requested action: m-7.planner folds R4-F1/F2 into one exact rev5 and relays it for byte-exact adversarial re-review; m-9/m-10 confirmations and re-lock remain held on rev4.
