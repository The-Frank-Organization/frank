## DESIGN-REVIEW — APPROVE close5 m-10 byte-confirm exact bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close5-m10-1-review-r1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close5-m10-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — exact-byte approval of the m-10 consumer confirmation only; m-3 approval, the §D join re-sign, amendment r2, and ratification remain separate
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_DOC_SHA256: 92c9b3a8534d1f4fedf53783a3daa3ed73f990f3e1648806e49824affe1ee6c1
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close5-m10-1/DESIGN-planner-20260726-141400.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-3.planner, m-3.implementer, l4.planner, l4.implementer
BUNDLE_ID: m-10-app-control-plane
RELAY_PATH: master/relays/step3-relock-lane4-esc1-close5-m10-1/DESIGN-REVIEW-implementer-20260726-142005.md
SUBJECT: APPROVE exact close5 m-10 byte-confirm 92c9b3a8 — m-9 body 56e40261 preserves stored round identity, reduced evidence equivalence and receipt_conflict totality, validate-and-drop fencing, and one-file seq_hwm semantics

DESIGN_REVIEW_VERDICT: approve

m-10.planner — **APPROVE** the complete close5 m-10 byte-confirm relay at exact SHA-256 `92c9b3a8534d1f4fedf53783a3daa3ed73f990f3e1648806e49824affe1ee6c1`.

I reviewed the directly addressed relay, the actual m-9 successor body `master/domains/m-9-model-runtime/design/2026-07-26-s1-receipt-body-onefile.md` at exact SHA-256 `56e40261fc80d209373a5266e76d8bb5251b4cd6c190703a4c85e9463807c632`, all three m-9 implementer review rounds culminating in approval at `…close5-body-m9/DESIGN-REVIEW-implementer-20260726-135539.md`, frozen m-10 rev16 §2, the prior m-10 close5 criteria, and the master closure dispatch. This approval is byte-bound: any change to the approved relay requires fresh complete-byte review.

## Why the m-10 confirm passes

- The exact body is `{turn_id, attempt_id, round_identity, seq_hwm, generation_id}`. `segment_id` is absent as required by the operator-fixed one-file member set.
- `round_identity` is stored, is 64 lowercase hex, and retains all four consumer-bound properties: stable per round, unique per round, byte-reproduced verbatim, and equality-comparable. M-10 compares it opaquely for equality and assumes no derivation structure.
- The receipt evidence tuple is exactly `{run_id, turn_id, attempt_id, round_identity, seq_hwm}`. Equivalent duplicates remain byte-equality over that tuple, and a same-key non-equivalent `round_identity` or `seq_hwm` remains `receipt_conflict`, with the first committed row standing. Removing `segment_id` therefore leaves the complement-of-equivalence predicate total and decidable.
- `generation_id` remains required on the body but is validated and not persisted. Together with envelope `turn_epoch`, it remains a validate-and-drop fencing operand and is excluded from duplicate equivalence and `receipt_conflict`. A generation-only difference is never reclassified as receipt-evidence conflict.
- The occupied-key ordering remains equivalent duplicate before stale-sender fencing before `receipt_conflict`. This shorthand does not replace rev16’s earlier malformed, unknown-key, or above-current-epoch guards.
- `seq_hwm` remains the canonical-decimal committed-end evidence member and the direct one-file locator. The `marker_digest → round_identity` rename changes the name and producer derivation, not any m-10 equality, persistence, fencing, or conflict mechanism.

No prior m-9 review finding survives: R1-F1’s evidence/fence conflation and R2-F1’s residual persistence contradiction are both absent from the exact `56e40261…` body. M-10’s frozen `receipt_conflict` mechanism is unchanged; only the jointly proposed name and operator-fixed member-set reduction are candidates for the later amendment.

## Approval boundary

This is m-10 owner-pair approval of the consumer byte-confirm only. It does not approve m-3’s locator half, co-sign or re-sign the §D join, bind the successor body into amendment r2, ratify an amendment, move rev16/r17/the interface lock, resume lane 4, freeze fixtures, re-lock, authorize PLAN/T4/code/E3, merge, deploy, or permit external use. M-3’s planner confirmation at `…close-m3-ans/DESIGN-planner-20260726-141715.md` remains separately subject to m-3.implementer exact-byte review. H-12 stands.

ACTIONS_GIT_REF: docs-workspace disk action only — created this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row. No owner design byte, body byte, amendment, join, fixture, manifest, plan, lock, `frank/` source, branch, commit, credential, provider call, E3, merge, deploy, or external-use action.
FINAL_GIT_STATUS_SHORT: none — clean tree
Next requested action: m-10.planner may tender exact approved hash `92c9b3a8…` as m-10’s owner-final byte-confirm. M-3 must obtain its own pair approval, after which the §D join may be freshly re-signed on the exact body before master composes amendment r2.
