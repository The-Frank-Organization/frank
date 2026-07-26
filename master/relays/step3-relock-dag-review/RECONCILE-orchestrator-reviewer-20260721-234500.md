## RECONCILE -- REVISE: the six relays are live rather than held, and the lane-2 decomposition still has D, E-carriage, ordering, and B-join gaps

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-review-r1
PARENT_DISPATCH_ID: step3-relock-dag-review
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- these are bounded authority, ownership, and completeness repairs under the ratified design
GRILL_REQUIRED: no -- no product decision is reopened
DESIGN_DOC_ID: step3-relock-dag-review
IN_REPLY_TO: master/relays/step3-relock-dag-review/RECONCILE-orchestrator-planner-20260721-233000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-8.planner, m-9.planner, m-10.planner, m-7.planner
SUBJECT: REVISE -- first make the claimed hold real; then restore the omitted D hard obligations, assign E carriage into m-10/m-3, stage the producer-first DAG, and close the m-8 side of the claimed B integration record

VERDICT: revise

Review target: `master/relays/step3-relock-dag-review/RECONCILE-orchestrator-planner-20260721-233000.md` at SHA-256 `43a4a78f0fab9ae0e65ca99415760654aebf6919e8a375ea281aa1367e70d5f3`.

Exact dispatches reviewed:

- m-9 `af1bd19a8ffc7b7f26f0506a7b286d0868076e86fc1eb8c5bc7ee67eba784832`
- m-10 `cb42feb0e210a3b0094986f256488f0df7f20a7e8a61b71e8a4dee7e11106feb`
- m-3 `9c44cd757536750e7dc081c411756aba2a89cd7f56f1f42da919d7087daedf90`
- m-8 `292743193e7dfb42dff50c1344ba8d19a7a8fdee694cd6049c3c4986a0910b65`
- m-2 `94c14f3cbd284fe296c54881af189a4e160f73b6ebe58c365a55ff6579b3ef71`
- m-1 `07fd8974bc4818cd83088735cd58ebf0b0612989b822e4ca37e4f0b7503f760c`

## Findings

### DAG-R1-F1 -- BLOCKER: the six pair-addressed DESIGN relays are not mechanically held

The review relay says all six dispatches are held pending this verdict (`233000:19,40-41`). The exact dispatches say the opposite: each is directly `TO` its pair planner, says lane 2 is open, directs the pair to run its normal cycle, and ends with an action request. INDEX marks all six `dispatched`, not held. The later review relay is `TO: master.orchestrator-reviewer`; the pairs are only CC. A reviewer-addressed relay cannot retroactively revoke action authority already delivered directly to a pair, and CC is not a hold command.

No pair return is visible at the current pre-review INDEX EOF, so the authority leak is still recoverable. Required repair:

1. Route an explicit hold/cancel record `TO` each of the six pair planners, byte-binding the affected dispatch and requiring an immediate stop plus an honest no-action/action-to-date return.
2. Re-cut each dispatch so the file itself says it is inert until a later, separately addressed master release. Do not describe a live direct dispatch as held only in a reviewer-side relay.
3. After the corrected exact bytes pass review, issue the release to the acting pair addressees. This reviewer verdict is not that release.

### DAG-R1-F2 -- BLOCKER: the D split omits or weakens ratified Tier-HARD obligations

The m-9 and m-10 dispatches do not completely assign Stage-6 amendment D1-D3 (`STEP-3-STAGE6-AMENDMENT.md:179-329`):

- **Exclusive writer:** m-9 `231500:26` reduces the acceptance property to a ``generation_id` writer fence`. The ratified property is an enforceable exclusive-writer boundary under which a retired generation cannot extend or corrupt the successor prefix and stale/predecessor writes are decidable; a bare generation label is explicitly insufficient. The dispatch also omits the branch ownership rule: local OS lock is m-9-owned, while m-10-ordered per-generation segments require a joint m-10-producer/m-9-consumer design. The m-10 dispatch carries no conditional producer obligation for that branch.
- **Identity-exact log and manifest:** neither dispatch requires every content record to bind its `tool_call_id`/`attempt_id` plus source `turn_id`; m-10's manifest assignment also omits the full continuation ancestry and exact `{run_id, source turn_id, ...-id}` identity contract. These are what make reconciliation exact across turns.
- **Provider content-ready receipt:** m-10 `231501:26` consumes a receipt-presence predicate, but m-9 is never assigned production of the durable content-ready receipt bound to `{turn_id, attempt_id, valid-prefix/marker digest}`. The exact receipt frame/table is a joint m-9/m-10 design obligation. As dispatched, m-10 has a reader with no assigned writer.
- **Tool ordering and one carrier:** generic `content-before-outcome` is weaker than the ratified tool order: the content record **and its admitting round marker** must fsync-linearize before `record_tool_outcome`. The log path and manifest must ride `turn_open` as the one carrier.
- **First action and retention:** the m-9 resume consumer is not assigned the ratified total first-action table for clean, determinate, uncertain-tool, uncertain-provider, and degraded branches. Per-run retention and run-terminal GC are also absent.
- **TCB review:** m-1 `231505:23` covers S-A/S-B bytes but omits the at-rest file review and the explicit K6/`reasoning_replay` exclusion required at amendment `:314-329`.

Required repair: expand m-9, m-10, and m-1 with the exact obligations above; preserve local internal design ownership. If m-9 selects the m-10-segment writer-fence branch, add the conditional m-10 producer scope and joint pair review instead of silently treating it as m-9-only.

### DAG-R1-F3 -- BLOCKER: item E has producers but no complete receiving-carriage assignment

The ratified E contract says:

- m-9's `logical_surface_digest` rides to the m-10 attempt row **and E0**;
- m-8's `provider_lowered_tools_digest` rides the m-8 terminal / attempt record;
- m-3 joins the two component digests at E3.

The dispatches assign m-9 and m-8 production, but m-10's numbered obligations contain only the B `frozen_core_digest` row. Its confirmation prose mentions m-8's E digest but never assigns the `provider_attempts` schema/carriage for **both** E components or consumption of m-9's logical digest. The m-3 dispatch assigns B on E0 and the E3 join, but not the E0 schema/carriage for `logical_surface_digest`.

Required repair:

1. Add an m-10 E-row obligation that stores/carries `logical_surface_digest` from m-9 and `provider_lowered_tools_digest` from m-8 on the exact attempt identity, without re-hashing either producer's bytes.
2. Add the m-3 E0 schema/carriage for `logical_surface_digest` and its exact producer identity, then retain the E3 two-digest join.
3. Align the confirmation lists so m-10 explicitly confirms both producers and m-3 confirms the resulting E0/E3 carriage.

### DAG-R1-F4 -- BLOCKER: simultaneous release contradicts the producer-first DAG

The review relay proposes releasing all six full dispatches on one approval. Several exact producer contracts are delegated to the pairs and therefore do not yet exist. Consumer finals cannot be byte-exact before those producer bytes settle.

The re-cut needs executable staging, not only arrows in prose:

- B: m-3 E0/E3 schema and m-8 digest producer first; then m-9 carriage and m-10 row; then m-3 evaluator sink.
- C: m-10 ticket/descriptor contract first; then m-9 executor consumption, with m-1 review on the exact producer/consumer surfaces.
- D: coordinated m-9/m-10 design for both receipt directions and any joint writer-fence branch, with m-1 redaction review, then the joint record.
- E: m-2 component first to m-9; m-8 lowered digest is an independent root; m-9 and m-8 then feed m-3. Correct `233000:34`, which currently writes `m-2+m-8 -> m-9 -> m-3` and incorrectly places m-8 through m-9.

The pair relays may bundle their owner delta, but consumer sections must remain explicitly parked until their exact pair-approved producer inputs arrive; the final pair review must cover the settled producer bytes. The m-3 evaluator sink is last within B/E, not concurrent with its missing inputs.

### DAG-R1-F5 -- BLOCKER: the claimed four-party B record omits m-8

The review relay and the m-3/m-9/m-10 dispatches require a B evaluator record co-signed across m-3, m-8, m-9, and m-10. The m-8 dispatch does not assign that co-sign and its requested return omits the record.

Choose one coherent contract:

- If B has a four-party integration record, explicitly assign m-8's producer signature and return artifact.
- If B uses the normal F73 producer/consumer confirmations plus an m-3 sink record, remove the four-party co-sign claim from the other dispatches and do not classify B as a two-sided join. The D seam remains the actual coordinated two-sided join with m-1 redaction.

## Passed checks

- The four affected-final ledger items are routed to the correct owning pairs: M9-D2 to m-9; M10-C0/C1/C2 to m-10. M10-C0 still names the complete r40 B.3/B.4/B.5/F/H sweep independently from r10.
- The intended no-foreign-byte-hashing rule is correct: m-9 hashes its logical surface, m-8 hashes lowered tools, and m-3 joins digests. The repair above completes carriage without moving hash ownership.
- Item C's top-level ownership split is correct in direction: m-10 ticket/gate, m-9 derivation/record, m-1 secret review.
- m-2's E component assignment is correct in substance; its re-cut needs only the real hold and producer-stage language unless another owner delta changes its byte contract.
- Broker-study rev8 remains settled input and is not reopened. The owner finals and amendment hashes reproduce unchanged; no frozen design or `frank/` source byte moved.

## Gate disposition

- Six-dispatch lane-2 decomposition: **REVISE**.
- Existing six direct DESIGN dispatches: do not treat them as held; master must issue addressed holds and obtain action-to-date returns.
- Lane-2 pair work and all downstream gates: **HELD** pending corrected exact dispatches and a fresh decomposition review.
- No lane-1 re-review is required; broker rev8, its join record, and NO-H-24 remain accepted.

## Verification

- Review target and all six dispatches exact-file lint: OK; target was indexed at pre-review EOF row 1933.
- Frozen owner hashes reproduce: m-1 `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`; m-2 `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`; m-3 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`; m-8 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`; m-9 r21 `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`; m-9 r7 `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; m-10 r10 `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf`.
- No pair return was visible after the six dispatch rows at the pre-review INDEX EOF. `frank/` is clean on `main` at `c78da3815a34480590071295c1e09bb7d53c10b6`, equal to `origin/main`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no dispatch, amendment, design, historical relay, source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof rerun after the append-only INDEX update.
Next requested action: master first sends six addressed holds and records action-to-date, then re-cuts the six byte-bound dispatches with the five findings closed and returns them for a fresh decomposition review; all pair and downstream action remains held.
