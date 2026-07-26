## DESIGN-REVIEW — rev8 APPROVE: correlation staging closes the final wire-totality defect

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-m7-broker-review-r8
PARENT_DISPATCH_ID: step3-relock-m7-broker
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-m7-broker-study
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/step3-relock-m7-broker/DESIGN-planner-20260721-203009.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-1.planner
BUNDLE_ID: m-7-conductor-core
SUBJECT: byte-exact approval of `master/domains/m-7-conductor-core/design/2026-07-21-step3-broker-study-crash-adoption-epoch.md` rev8 at SHA-256 `64f9136e0b851e31f129372ec50b5667c2b6dcda197ef0dcf95c4f0eca5ff4ce`

## Verdict

**APPROVE.** R7-F1 closes. Rev8 stages request handling at the only sound boundary: frames without one schema-valid correlation take the existing no-result control-frame fault, while malformed requests with a parsed correlation return the correlated `rejected-malformed`. No identity is fabricated or defaulted, m-10 fails closed on either path, and the unknown-field election plus fixtures make the split deterministic.

No unresolved pair-review finding remains on these exact bytes.

## Accepted Mechanism

- Broker survival is an m-10 supervision property; app-main recovery adopts control without moving credential bytes or rebinding the seat channel.
- Generation recovery remains lifecycle-dependent: ordinary leased-generation recovery retires exactly once, mints E+1 exactly once, and installs before successor assignment.
- Cross-epoch completion is not retained. The broker stops old-E authority at PROPOSED, drains to the fixed local deadline, cuts unresolved operations honestly, installs locally even through control loss, and discards post-install old-E responses.
- Mandatory relay-tool uncertainty remains in m-10's existing durable F59 outcome/disclosure machinery; caller and `boundary_cut` visibility are conditional or telemetry-only as specified. `Describe` retains its separate effect-free posture.
- The no-ledger install-proof join is keyed by `{run_id, generation_id, turn_epoch, state_seq}` and is two-form: a tuple-matching `epoch_installed` event or tuple-matching `state_proposal_result{installed}` opens one logical gate.
- The ordered proposal table is total: correlation-unrecoverable frame fault, correlated malformed rejection, fresh-instance bootstrap, active-target join/conflict, and the no-active comparison ladder each have one deterministic outcome. Proof is idempotent evidence; assignment uniqueness remains m-10 lifecycle-owned.
- `crossing_count` is removed from the amended closed `epoch_installed` schema, with unknown-field rejection for the former shape.
- The fixture battery discriminates response/event loss, correlation and tuple mismatch, controller/broker crash cuts, PREPARING replay/conflict, fresh bootstrap, malformed/stale precedence, and correlation-boundary faults.

## Approval Scope

This approval binds only to rev8 SHA-256 `64f9136e0b851e31f129372ec50b5667c2b6dcda197ef0dcf95c4f0eca5ff4ce`. Any byte change voids it.

Approval opens only the dispatch-prescribed next design steps under master routing: m-9/m-10 F73 confirmations, the section-D join record, and m-7.planner's SITREP to master and VP. It issues no DESIGN lock, PLAN, T4/code token, credential, provider, release, E3, merge, or deploy authority.

## Verification

- Reviewed rev8 at exact SHA-256 `64f9136e0b851e31f129372ec50b5667c2b6dcda197ef0dcf95c4f0eca5ff4ce`; incoming relay at exact SHA-256 `bdb62bc802122ca348d5692ac5faf3cdf1fe2f01b2d5567546f0c2582bdb319e`.
- Reproduced frozen bases: m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; m-9 `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Incoming exact-file relay lint reports `OK`.
- No product design or code bytes were modified by this review.

ACTIONS_GIT_REF: none — review-only docs action; no `frank/` implementation authority exercised.
FINAL_GIT_STATUS_SHORT: none — `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6` after this relay and INDEX row were written.
Next requested action: m-7.planner files the approved-byte SITREP to master and VP; master routes the m-9/m-10 confirmations and section-D join record. All later lock/action gates remain held pending that reconciliation.
