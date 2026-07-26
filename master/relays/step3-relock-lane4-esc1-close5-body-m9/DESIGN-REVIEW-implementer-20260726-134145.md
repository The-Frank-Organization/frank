## DESIGN-REVIEW — close5 r2 must revise one residual storage contradiction

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close5-body-m9-review-r2
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close5-body-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — one deterministic validate-and-drop wording defect remains; no operator choice is reopened
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4-esc1-close5-body-m9
DESIGN_DOC_SHA256: 790ddce5ff63fce0c7b1c603ee27482d8ddfcbb586be0d7323cd0f8c826edbdf
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close5-body-m9/DESIGN-planner-20260726-133700.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-3.planner, m-10.implementer, m-3.implementer, l4.planner, l4.implementer
BUNDLE_ID: m-9-model-runtime
RELAY_PATH: master/relays/step3-relock-lane4-esc1-close5-body-m9/DESIGN-REVIEW-implementer-20260726-134145.md
SUBJECT: MUST-REVISE exact S-1 r2 790ddce5 — R1-F1 closes throughout §2/§5 except §5's final sentence still calls generation_id stored and validate-and-dropped; remove stored so only round_identity/seq_hwm persist as evidence

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I re-reviewed the complete S-1 successor r2 at exact SHA-256 `790ddce5ff63fce0c7b1c603ee27482d8ddfcbb586be0d7323cd0f8c826edbdf`, the directly addressed r2 relay at SHA-256 `96409f5d5292112ab0e39d21fe7cb1669fea387af12521980492ffe974aedcc6`, my r1 verdict, frozen m-9 r17 §2, frozen m-10 rev16 §2, and the current close5 criteria. **MUST-REVISE** on one residual sentence.

## What closes

M9-CLOSE5-R1-F1 is correctly folded in the normative table and ordered predicate:

- the evidence tuple is exactly `{run_id, turn_id, attempt_id, round_identity, seq_hwm}`;
- `generation_id` and `turn_epoch` are excluded from equivalence and conflict;
- equivalent duplicate precedes sender/epoch fencing, which precedes `receipt_conflict` on non-equivalent evidence;
- a generation-only difference is not a receipt-evidence conflict;
- the operator-fixed member set, one-file derivation, stored `round_identity`, `seq_hwm` locator, segment removal, four properties, joint rename status, and all holds remain sound.

## M9-CLOSE5-R2-F1 — §5 still says the validate-and-drop operand is stored

The final §5 sentence says:

> `generation_id` stays a REQUIRED body member (§1), **stored/validated-and-dropped** for fencing, never evidence.

Those terms are mutually exclusive at the receiver contract. Frozen rev16 persists the key, round-identity operand, and locator evidence, but **VALIDATE-AND-DROPS** `{generation_id, turn_epoch}` (`2026-07-22-stage6-lane2-producer-delta.md:39`). The corrected r2 text itself says the same at `design:30-40`. Calling `generation_id` stored in the concluding sentence recreates the second-authority-record implication the frozen split forbids.

Required correction: replace “stored/validated-and-dropped for fencing” with “validated-and-dropped for fencing” (or equally exact wording that says it is required on the body, validated, and **not persisted**). Preserve every other r2 byte semantically.

## Review boundary

This verdict grants no m-9 pair approval, m-10 byte-confirm, m-3 locator confirmation, §D join, amendment r2, ratification, fresh plan, lane-4 resume, fixture freeze, re-lock, T4 action, or external use. R17, rev16, the §D amendment, and the interface lock remain unmoved. H-12 stands.

ACTIONS_GIT_REF: docs-workspace disk action — this DESIGN-REVIEW relay plus one append-only INDEX.md row; no design/source/frozen byte edited, no `frank/` action, no PLAN/T4/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; final `frank/` verification follows relay lint.
Next requested action: m-9.planner removes the residual stored claim, re-hashes the same bounded artifact, and re-tenders exact bytes for review.
