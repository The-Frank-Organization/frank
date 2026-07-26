## DESIGN-REVIEW — close5 S-1 exact body must revise: `generation_id` is a fence operand, not receipt evidence

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close5-body-m9-review-r1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close5-body-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the blocker is a deterministic mismatch with the frozen m-10 receipt disposition; no operator choice is reopened
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4-esc1-close5-body-m9
DESIGN_DOC_SHA256: d5d79027f5ebb745e9bc47be295d0c76406de052a33d1d6f4976d56fd8ac8860
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close5-body-m9/DESIGN-planner-20260726-131200.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-3.planner, m-10.implementer, m-3.implementer, l4.planner, l4.implementer
BUNDLE_ID: m-9-model-runtime
RELAY_PATH: master/relays/step3-relock-lane4-esc1-close5-body-m9/DESIGN-REVIEW-implementer-20260726-132622.md
SUBJECT: MUST-REVISE exact S-1 successor body d5d79027 — member set, one-file derivation, stored round identity, and segment removal are sound, but §5 incorrectly puts generation_id into duplicate equivalence/receipt_conflict instead of preserving validate-and-drop fencing

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete new document `master/domains/m-9-model-runtime/design/2026-07-26-s1-receipt-body-onefile.md` at exact SHA-256 `d5d79027f5ebb745e9bc47be295d0c76406de052a33d1d6f4976d56fd8ac8860`, the directly addressed DESIGN relay at SHA-256 `c2f095cde649633dd2b31e587ede355a7cecbc71e806f2bf0e1944e776e367fa`, frozen m-9 r17 §2, frozen m-10 rev16 §2, m-10's current byte-confirm criteria, and the close5 parent. **MUST-REVISE.**

The operator-fixed body members are correct: `segment_id` is absent; `seq_hwm` and `generation_id` remain; `round_identity` is stored, encoded as 64 lowercase hex, and derived over the ordered one-file round content without the retired chain/segment terms. The rename is honestly parked on m-10 concurrence, and the §D re-sign remains held. One receiver-semantics defect blocks exact-byte pair approval.

## M9-CLOSE5-R1-F1 — `generation_id` is wrongly promoted from validate-and-drop fencing into receipt evidence

The new §5 defines the evidence tuple as `{run_id, turn_id, attempt_id, round_identity, seq_hwm, generation_id}` and fires `receipt_conflict` when `generation_id` differs (`design:30-31`).

That contradicts both frozen owners' exact predicate:

- m-9 r17 §2 excludes `generation_id` and `turn_epoch` from duplicate equivalence because they are fenced separately (`2026-07-22-relock-lane2-m9-delta.md:310-317`);
- m-10 rev16 persists the key, round-identity operand, and locators, but **validate-and-drops** `{generation_id, turn_epoch}` (`2026-07-22-stage6-lane2-producer-delta.md:39`); its ordered receiver compares the evidence tuple before stale-sender and expressly excludes both fence operands from equivalence (`:41`);
- m-10's current close5 criteria preserve exactly that split: the reduced equivalence tuple is `{run_id, turn_id, attempt_id, <round-identity>, seq_hwm}`, while `generation_id` remains a required validate-and-drop stale-sender operand (`...close5-m10-1/...131600.md:28-34`).

Required revision: keep `generation_id` in the S-1 body, but remove it from the stored/equivalence/conflict tuple. State the reduced evidence tuple exactly as `{run_id, turn_id, attempt_id, round_identity, seq_hwm}` and preserve the frozen first-match ordering: equivalent duplicate first; sender/epoch fencing on `generation_id`/`turn_epoch`; `receipt_conflict` only for a same-key non-equivalent **evidence** tuple. A generation-only difference is not a receipt-evidence conflict.

## Review boundary

This verdict pair-approves no S-1 bytes, m-10 reciprocal, m-3 locator, §D join, amendment r2, fresh lane-4 plan, fixture freeze, re-lock, T4 action, or external use. r17 `01b885fe…`, rev16 `3e3c5192…`, §D amendment `1fa71cb8…`, and interface lock `cbd1893c…` remain unmoved. H-12 stands.

ACTIONS_GIT_REF: docs-workspace disk action — this DESIGN-REVIEW relay plus one append-only INDEX.md row; no design/source/frozen byte edited, no `frank/` action, no PLAN/T4/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; `frank/` verification is recorded after relay lint.
Next requested action: m-9.planner issues a fresh exact-body revision correcting only the fence/evidence split, with a new SHA and directly addressed review relay; then m-10 byte-confirms, m-3 confirms `seq_hwm`, both pairs approve, and §D re-signs before r2.
