## DESIGN-REVIEW — rev7 MUST REVISE: the state table is total, but its unparseable-input row cannot satisfy the mandatory correlation echo

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-m7-broker-review-r7
PARENT_DISPATCH_ID: step3-relock-m7-broker
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-m7-broker-study
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-m7-broker/DESIGN-planner-20260721-200218.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-1.planner
BUNDLE_ID: m-7-conductor-core
SUBJECT: byte-exact re-review of `master/domains/m-7-conductor-core/design/2026-07-21-step3-broker-study-crash-adoption-epoch.md` rev7 at SHA-256 `c33c91919591e2725453e23b7358f3f90dc651c8e35fc0a28811c65730669a30` — R6-F1 closes; one correlation-totality defect remains in the malformed-input branch

## Verdict

**MUST REVISE.** R6-F1 closes: the ordered table covers bootstrap, makes PREPARING precedence deterministic, preserves the active deadline/target, and supplies discriminating fixtures. The clean-cut/no-H-24 determination and all earlier closures remain supportable.

The malformed-first row now exposes one exact wire impossibility. It requires `rejected-malformed` for an unparseable request, while every `state_proposal_result` requires an echoed `proposal_correlation`. A request whose correlation is missing, malformed, or unparseable provides no valid value the broker can echo.

## Closed From R6

- **R6-F1 CLOSED:** bootstrap installs synchronously with `installed_state`; ordered first-match evaluation removes reject overlap; target-equal replay, active-other rejection, no-active comparison, and the three requested fixtures are deterministic.

## Finding

### R7-F1 — BLOCKER — frame-level malformed input cannot emit the required correlated result

**Current bytes:** `state_proposal_result` has a closed required `proposal_correlation (echoed)` field for every disposition. Ordered row 1 classifies “unparseable frame, unknown fields, grammar violation” as `rejected-malformed`. The set includes frames with no parseable valid correlation, so the required response cannot be constructed without inventing or defaulting an identity. Either choice would violate the exact correlation rule and m-10's live-request validation.

**Required correction:** stage malformed handling at the correlation boundary:

1. A frame that cannot yield one schema-valid `proposal_correlation` takes the existing exact control-frame/channel fault disposition and emits no `state_proposal_result`, unless a separate uncorrelated error frame is deliberately defined. m-10 fails closed on the absent result.
2. `rejected-malformed` is available only after a valid live correlation is parsed, for a request malformed in another field or closed-shape rule; it echoes that parsed correlation and remains non-proof.
3. State whether an unknown-field violation with an otherwise valid correlation belongs to the correlated `rejected-malformed` branch or the frame-fault branch. One reading only.
4. Add fixtures for missing correlation, invalid correlation encoding, and a correlated shape violation. Assert no fabricated/default correlation, no state mutation, and exactly the selected response-or-fault behavior.

This correction is entirely before the ordered state comparison. It does not reopen the eight state rows, add a disposition token, or require durable transition state.

## Re-review Gate

Return one byte-exact rev8 preserving every closed r1-r6 mechanism and splitting only uncorrelatable frame faults from correlated malformed rejections, with the three fixtures above. Consumer confirmations, join record, SITREP, and re-lock remain held until those bytes pair-approve.

## Verification

- Reviewed rev7 at exact SHA-256 `c33c91919591e2725453e23b7358f3f90dc651c8e35fc0a28811c65730669a30`; incoming relay at exact SHA-256 `c938f3a758810da72c69e61f37511213a3d55b946fed253bc5f6eeda4c316f70`.
- Reproduced frozen bases: m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; m-9 `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Incoming exact-file relay lint reports `OK`.
- No product design or code bytes were modified by this review.

ACTIONS_GIT_REF: none — review-only docs action; no `frank/` implementation authority exercised.
FINAL_GIT_STATUS_SHORT: none — `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6` after this relay and INDEX row were written.
Next requested action: m-7.planner folds R7-F1 into one exact rev8 and relays it for byte-exact adversarial re-review; m-9/m-10 confirmations and re-lock remain held on rev7.
