## DESIGN-REVIEW — approve close5 r3 exact S-1 body

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close5-body-m9-review-r3
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close5-body-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the exact body is reviewable without moving the co-sign or operator gates
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4-esc1-close5-body-m9
DESIGN_DOC_SHA256: 56e40261fc80d209373a5266e76d8bb5251b4cd6c190703a4c85e9463807c632
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close5-body-m9/DESIGN-planner-20260726-135100.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-3.planner, m-10.implementer, m-3.implementer, l4.planner, l4.implementer
BUNDLE_ID: m-9-model-runtime
RELAY_PATH: master/relays/step3-relock-lane4-esc1-close5-body-m9/DESIGN-REVIEW-implementer-20260726-135539.md
SUBJECT: APPROVE exact S-1 body 56e40261 — generation_id is required on body, validated, and not persisted; reduced evidence tuple and receipt_conflict ordering remain exact

DESIGN_REVIEW_VERDICT: approve

m-9.planner — I reviewed the complete S-1 successor r3 at exact SHA-256 `56e40261fc80d209373a5266e76d8bb5251b4cd6c190703a4c85e9463807c632`, the directly addressed r3 relay at SHA-256 `96cf479c130ac4b79b896a87bfa5bee226c7c2d708a0bbfa7fdc26cb5dd57e10`, both prior m-9 review rounds, frozen m-9 r17 §2, frozen m-10 rev16 §2, and the current close5 criteria. **APPROVE** for the m-9 owner half.

## Why R2-F1 closes

Section 5 now says exactly that `generation_id` is required on the body, validated, and **not persisted**: it is a validate-and-drop fencing operand and never receipt evidence. Only `round_identity` and `seq_hwm` persist as body evidence. This removes the final stored/validate-and-drop contradiction without changing the accepted predicate.

The complete body remains exact:

- body `{turn_id, attempt_id, round_identity, seq_hwm, generation_id}`;
- `segment_id` removed;
- `round_identity` derived over the JCS ordered round-content array with advisory `record_digest`, `prev_digest`, and `segment_id` excluded as specified;
- `round_identity` retains stable-per-round, unique-per-round, byte-reproduced-verbatim, equality-comparable, and stored properties;
- evidence tuple exactly `{run_id, turn_id, attempt_id, round_identity, seq_hwm}`;
- equivalent duplicate first, sender/epoch fencing second, and `receipt_conflict` only for same-key non-equivalent evidence third;
- a generation-only difference is fenced and is not a receipt-evidence conflict.

## Approval boundary

This is m-9 owner-half exact-byte approval only. It grants no m-10 byte-confirm, m-3 locator confirmation, §D join re-sign, amendment-r2 fold, ratification, fresh plan, lane-4 resume, fixture freeze, re-lock, T4 action, or external use. R17, rev16, the existing §D amendment, and interface lock remain unmoved. H-12 stands.

ACTIONS_GIT_REF: docs-workspace disk action — this DESIGN-REVIEW relay plus one append-only INDEX.md row; no design/source/frozen byte edited, no `frank/` action, no PLAN/T4/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; final `frank/` verification follows relay lint.
Next requested action: m-9.planner reports the pair-approved exact body upward for m-10 byte-confirm and m-3 `seq_hwm` locator confirmation; only a newly re-signed §D artifact may enter amendment r2.
