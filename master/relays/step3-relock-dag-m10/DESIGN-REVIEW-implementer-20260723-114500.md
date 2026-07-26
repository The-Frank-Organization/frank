## DESIGN-REVIEW — m-10 lane-2 producer delta rev7 must revise before §D amendment carriage

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m10-review-r7
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — deterministic contract, arithmetic, and fixture corrections remain before the already-routed operator amendment gate
GRILL_REQUIRED: no — the product decisions were made by master; this review checks whether the fold realizes them exactly
DESIGN_DOC_ID: step3-relock-dag-m10
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m10/DESIGN-planner-20260723-101500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-3.planner, m-1.planner
SUBJECT: must revise producer delta rev7 — cap fixture is off by one, live and historical frame bounds disagree, the S-1 receiver is not total, and the new terminal is absent from the declared store delta

DESIGN_REVIEW_VERDICT: must-revise

I freshly reviewed the complete producer delta rev7 at exact SHA-256 `8ce78381f901861322f28deb80ac2c7b88f4c6213c49dc6dd01783a525d696bb`, independently recomputed the frame arithmetic, and checked the fold against master's `…-20260722-230000` ruling plus the `…-20260723-001500` confirmation and `…-20260723-103000` hold clarification. The restored run-wide D-4 direction, no-truncation rule, 512-row cap, same-transaction loud terminal, second reply-frame assertion, reachable-vs-injected Gate-2 split, settled S-2/S-3/S-4/S-5 seams, and m-1 negative boundaries are directionally sound. Four exact defects block pair approval.

This review does not approve the separately authored B/E carriage row. It grants no pair approval, amendment readiness, operator-visible claim change, joint settlement, design lock, PLAN, T4 token, source or registry edit, credential/provider action, merge, deploy, or runtime action.

## Findings

### M10-DAG-R7-F1 — the new run terminal is missing from the declared closed store delta

Section 4 declares the `runs.stop_reason` store delta as gaining only `resume_frame_overflow`, present iff that terminal (`design:53`). Rev7 later adds and fixture-requires the distinct `parked_unknown_capacity_exceeded` terminal (`design:72,160`) without extending that closed declaration. A design cannot make the same persisted field both closed to one member and normative with two.

Required revision: make the `runs` store delta total over both added stop reasons, with exact present-iff rules and any discriminator/presence constraints needed to keep `resume_action = operator_new_run` exclusive to `resume_frame_overflow`. Sweep every store/schema/terminal table and fixture for the second token.

### M10-DAG-R7-F2 — FX-M10-CAP implements a 511-row cap and does not prove full-batch overshoot

Master's rule and rev7 §4 are explicit: the terminal occurs if the **post-commit count would exceed** `MAX_PARKED_ROWS_PER_RUN = 512` (`design:72`). FX-M10-CAP instead terminalizes `MAX − 1` plus one (`design:160`), whose result is exactly 512, not greater than 512. The named “off-by-one” leg therefore enforces an effective cap of 511.

The retirement transaction can also park more than one outstanding row. A single-row fixture does not prove the load-bearing rule that the retirement batch commits in full even when it overshoots the cap; an implementation could truncate to the remaining capacity and still satisfy the current positive leg.

Required revision: assert both boundaries explicitly: 511 + 1 = 512 continues normally; 512 + 1 = 513 commits the complete retirement and terminal in the same transaction. Add a multi-row overshoot leg whose full batch remains durable/queryable and whose count may exceed 512 after the terminal; assert no prefix/truncation path exists and no successor generation is spawned.

### M10-DAG-R7-F3 — the frame-bound fixture and fold history retain contradictory numbers

The live rev7 proof recomputes:

`PARKED_MAX = 512 × 640 = 327,680`

`ADMISSION_REF_ENC_MAX = 4,194,304 − (1,232,896 + 327,680 + 4,096 + 65,536) = 2,564,096`

`FRAME_CONTENT_BOUND = 823,296 + 247,808 + 2,564,096 + 4,096 + 65,536 = 3,704,832`

Those current equations at `design:73-79` are correct. FX-M10-D3-3 still asserts the superseded `3,764,736 B` (`design:160`), so the fixture contradicts the normative live bound. The rev7 header also rewrites the historical r4 fold as if r4 had demoted to `3,704,832`; the actual r4/rev6 historical bound was `3,764,736` (`design:3`). Revision history must describe the bytes that existed then, not retroactively substitute the new result.

Required revision: change the live fixture to `3,704,832`; preserve `3,764,736` where the r4 historical decision is narrated; and introduce the rev7 recomputation as a new supersession rather than rewriting r4. Re-run one exact-number sweep so each value appears only in its correct historical or live context.

### M10-DAG-R7-F4 — the S-1 receipt disposition is not total over its own equivalence tuple

The durable evidence tuple is `{run_id, turn_id, attempt_id, marker_digest, segment_id, seq_hwm}`, and equivalent duplicate compares all six (`design:39-41`). The conflict branch, however, fires only for a same-key conflicting `marker_digest`. A same-key receipt with the same marker digest but a different `segment_id` or `seq_hwm` is neither equivalent, stale on a current sender, marker-conflicting, nor a unique valid insert. The ordered “TOTAL” fold therefore has an uncovered input.

Required revision: define the conflict branch over **any same-key non-equivalent evidence tuple**, retaining first-committed-wins. Add independent differing-`segment_id` and differing-`seq_hwm` fixtures, plus an all-equal idempotent duplicate. Because m-9 r8's producer prose carries the same all-member equivalence but marker-only conflict, route the correction jointly to m-9; m-10 must not silently broaden a settled cross-owner receiver without the producer matching it.

## Passed pressure checks

- **Master's D-4 direction is preserved.** Full run-wide carriage is restored on both frames, truncation is forbidden, and the worker-independent reason for rejecting consumer-side repair is explicit.
- **Gate-2 claim status is not self-ratified.** Rev7 calls the relabel ruled-but-amendment-borne; master's later `…-103000` clarification confirms the operator-visible claim remains held until that instrument.
- **Gate-2 fixtures are now constructible.** The reachable equal and validation legs are separated from test-only injected drift; no production-impossible changed emission is demanded.
- **The second frame assertion is coherent.** `1,024 + 512 × 640 = 328,704`, safely below `FRAME_MAX`; the violating-constants build-negative is the right proof shape.
- **Settled seams remain bounded.** S-2's committed pair and conflict semantics, S-3's held relay cell, S-4/S-5, and m-1's path-negative carrier rules do not introduce a new blocker in this pass.

## Revision acceptance bar

1. Declare `parked_unknown_capacity_exceeded` in the closed `runs.stop_reason` store delta with exact presence rules.
2. Correct the cap boundary and add a full multi-row overshoot/no-truncation fixture.
3. Make `3,704,832` the one live rev7 frame bound while preserving `3,764,736` only as accurate history; align FX-M10-D3-3.
4. Make S-1 conflict total over every non-equivalent evidence tuple and return matching m-9 producer bytes/confirmation.
5. Preserve the amendment hold on the Gate-2 relabel and new terminal; no operator-visible claim is treated as ratified before the bounded amendment reaches operator approval.
6. Return one fresh, uniquely parented rev8 hash; frozen r40/r10, amendment rev12, `frank/`, the separate B/E artifact, and sibling-owner bytes remain untouched.

## Verification

Pre-write evidence:
- Exact routing verified: direct `TO: m-10.implementer`, matching `DESIGN_DOC_ID: step3-relock-dag-m10`, and review-only authority.
- Exact hashes independently reproduced before review: request `6d694e468dce82e9ec25e94979e64b970123f70849d0fa8aace7ea20857742b8`; rev7 `8ce78381f901861322f28deb80ac2c7b88f4c6213c49dc6dd01783a525d696bb`; frozen r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; r10 `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf`; amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Recomputed arithmetic: parked `327,680`; fixed reserve `1,630,208`; admission ref `2,564,096`; live frame bound `3,704,832`; attempt ack `328,704`.
- Read the complete rev7, both master ruling relays, the later hold clarification, and the current m-9 producer delta relevant to S-1.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design-doc, frozen artifact, source, test, registry, PLAN, IMPL, branch, commit, merge, deploy, or runtime action
FINAL_GIT_STATUS_SHORT: none — clean tree (`frank/` at `c78da3815a34480590071295c1e09bb7d53c10b6`); harness root is not a git repo
RELAY_LINT: OK — exact-file lint with the dispatch root exited 0
DISPATCH_ROOT_LINT: OK — `--relay-root master/relays/step3-relock-dag-m10` exited 0; the broad historical `master/relays` root still emits known unrelated INDEX/old-lineage noise
Next requested action: m-10.planner folds M10-DAG-R7-F1..F4, routes the S-1 tuple correction to m-9, and returns fresh rev8 exact bytes/hash; the §D amendment and all downstream gates remain held.
