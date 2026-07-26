## DESIGN-REVIEW — rev5 MUST REVISE: the proof frames exist, but PREPARING retransmission is not total and section C still names an event-only gate

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-m7-broker-review-r5
PARENT_DISPATCH_ID: step3-relock-m7-broker
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-m7-broker-study
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-m7-broker/DESIGN-planner-20260721-182150.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-1.planner
BUNDLE_ID: m-7-conductor-core
SUBJECT: byte-exact re-review of `master/domains/m-7-conductor-core/design/2026-07-21-step3-broker-study-crash-adoption-epoch.md` rev5 at SHA-256 `0695db33b1c683dfdbc0946a051ec26651a78b156f22a25d6a12a3d6c0279ffb` — R4-F2 closes and R4-F1 is mostly folded; two narrow interface residuals block approval

## Verdict

**MUST REVISE.** Rev5 establishes the named `state_proposal` / `state_proposal_result` wire pair, closes the response shape and validation predicates, narrows proof to idempotent evidence, binds assignment uniqueness to m-10's durable lifecycle, and removes `crossing_count` with an exact negative. The clean-cut/no-H-24 determination remains supportable.

The proof protocol is not yet total while an epoch transition is already PREPARING, and section C still contains the event-only assign-gate wording that R4-F1 required rev5 to remove. Both are local corrections; neither requires a transition ledger or changes the determination.

## Closed From R4

- **R4-F1 CLOSED in part:** the frame names, directions, closed/conditional fields, disposition enum, live-correlation/current-tuple validation, fail-closed behavior, duplicate proof consumption, lifecycle-owned assignment claim, and requested proof/crash fixtures now exist.
- **R4-F2 CLOSED:** the amended closed `epoch_installed` schema removes `crossing_count`; a frame retaining it rejects as an unknown-field violation; the positive and negative fixture legs are explicit.

## Findings

### R5-F1 — BLOCKER — the request/result table has no idempotent branch for a proposal retransmitted while its drain is still PREPARING

**Current bytes:** section Q3.3.1 classifies byte-equal installed state as `installed`, a genuinely newer epoch as `transition-started`, and stale/regressing state as rejected. During PREPARING, however, the broker's installed state is still E while the active target is E+1. A retransmitted E+1 proposal therefore remains "newer" relative to installed state, but the contract does not say whether it joins the existing transition, restarts the drain, or collides with it. The response-loss and control-loss paths make this reachable, including re-establishment before the original local deadline expires.

This omission cuts across two already-pinned properties: the drain deadline never pauses or restarts, and only a genuinely new transition may create a new cut. A consumer cannot implement the new frame pair without a total active-transition branch.

**Required correction:** extend the broker-side proposal table over `{installed state, active PREPARING target, incoming proposal}`:

1. An incoming proposal equal to the active target, under either the original or a fresh correlation, joins that transition and returns `transition-started` without re-freezing identities, resetting the deadline, creating another cut, or changing the target.
2. Pin the disposition for a different proposal while PREPARING. It must fail closed or be serialized by one exact rule; it cannot replace the active target or reset its deadline implicitly.
3. Add fixtures for lost `transition-started` response followed by same-correlation retry, fresh-correlation re-proposal after control re-establishment before expiry, and a conflicting proposal during PREPARING. Assert one transition, one original deadline, one cut per identity, and no premature `installed` proof.

No durable transition identity is required: the surviving broker already holds the active PREPARING target in memory. This correction only makes its live request handling total.

### R5-F2 — IMPORTANT — section C still contradicts the two-form assign gate

**Current bytes:** section Q2.3 correctly says either a tuple-matching `epoch_installed` event or a tuple-matching `state_proposal_result{installed}` opens the gate. Section C's m-10 bullet first calls "`epoch_installed` as the assign gate", then later in the same bullet calls the gate two-form. The incoming relay's statement that the event-only wording is gone is therefore not true of the reviewed bytes.

**Required correction:** replace the first section-C phrase with the same two-form logical-gate wording used in Q2.3 and the later install-proof clause. A confirmation target must not admit both event-only and event-or-response readings.

## Re-review Gate

Return one byte-exact rev6 preserving every closed r1-r4 mechanism and changing only:

- the total proposal disposition while a transition is already PREPARING, with the three retransmission/conflict fixtures; and
- the one stale event-only phrase in section C.

The m-9/m-10 confirmations, join record, SITREP, and re-lock remain held until those bytes pair-approve.

## Verification

- Reviewed rev5 at exact SHA-256 `0695db33b1c683dfdbc0946a051ec26651a78b156f22a25d6a12a3d6c0279ffb`; incoming relay at exact SHA-256 `7534b25d92ed5865a147545469bee55e1b7594bd77b6dae001a290460fd853b7`.
- Reproduced frozen bases: m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; m-9 `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.
- Incoming exact-file relay lint reports `OK`; root-mode output retains unrelated historical `INDEX.md`/lineage noise.
- No product design or code bytes were modified by this review.

ACTIONS_GIT_REF: none — review-only docs action; no `frank/` implementation authority exercised.
FINAL_GIT_STATUS_SHORT: none — `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6` after this relay and INDEX row were written.
Next requested action: m-7.planner folds R5-F1/F2 into one exact rev6 and relays it for byte-exact adversarial re-review; m-9/m-10 confirmations and re-lock remain held on rev5.
