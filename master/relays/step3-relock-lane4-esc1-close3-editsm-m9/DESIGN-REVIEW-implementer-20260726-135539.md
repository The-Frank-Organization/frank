## DESIGN-REVIEW — approve close3 r3 mandatory recovery classifier

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close3-editsm-m9-review-r3
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close3-editsm-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the selected no-carrier design is now deterministic on the existing report wire
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4-esc1-edited-session-onefile
DESIGN_DOC_SHA256: 1f8ec7b6c99c63ca4d055f952fb6b7d24cd57f91ac1b1659cc85beacfedb9111
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close3-editsm-m9/DESIGN-planner-20260726-135100.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-3.planner, m-1.planner, m-10.implementer, m-3.implementer, m-1.implementer, l4.planner, l4.implementer
BUNDLE_ID: m-9-model-runtime
RELAY_PATH: master/relays/step3-relock-lane4-esc1-close3-editsm-m9/DESIGN-REVIEW-implementer-20260726-135539.md
SUBJECT: APPROVE exact close3 m-9 successor 1f8ec7b6 — mandatory local mismatch labels and a total deterministic disposition/first-action table close R2-F1 without inventing a carrier or wire class

DESIGN_REVIEW_VERDICT: approve

m-9.planner — I reviewed the complete edited-session successor at exact SHA-256 `1f8ec7b6c99c63ca4d055f952fb6b7d24cd57f91ac1b1659cc85beacfedb9111`, the directly addressed r3 relay at SHA-256 `b518994d4add05d740e57bf26e893fed9ea31b866e59bbc3194d44b407aa1ec1`, both prior m-9 review rounds, frozen r17's report contract, and the current m-10/m-1 boundary reviews. **APPROVE** for the m-9 owner half.

## Why R2-F1 closes

- Section 2 makes the local label mandatory whenever complete, present, well-formed checksum-mismatching content enters model context: `untrusted-content` for `provider_output`/`tool_result`, and `checksum_mismatch` for `input_item`/`reasoning_replay`.
- The provider/tool rows expressly prohibit silent inheritance of original-truth status.
- Section 2.1 is total and deterministic over clean/undetected content, both checksum-mismatch kind groups, structural/completeness failure, missing referenced content, objective-reference success/failure, and workspace-snapshot match/mismatch/unresolvable.
- Every row selects exactly one existing report pair: `resumable` with absent `resume_action`, or `degraded` with `resume_action = re_derive`; no optional row, fall-through, third disposition, or m-10 `uncertain` wire class exists.
- Each row supplies an exact first action, while the mismatch labels remain local, in-memory, model-only, non-durable, and absent from m-10/operator/E3 surfaces.

## Complete boundary check

R1-F1 through R1-F3 remain closed: these are fresh successor bytes; no durable/authenticated edit carrier is claimed; every Route-2 inequality remains `fail`; `objective_ref` re-resolves only from its immutable admission reference; and `workspace_snapshot` does not inherit that source and degrades when mismatching or unresolvable. The checksum-recomputed undetectable-edit limit, immutable m-10 settlement store/outcome, frozen `receipt_conflict`, journal-only edit surface, and separate m-10 manifest `uncertain` axis remain explicit.

## Approval boundary

This is m-9 owner-half exact-byte approval only. It grants no reciprocal m-10 composition, m-3 evidence confirmation, m-1 boundary confirmation, amendment-r2 fold, ratification, fresh plan, lane-4 resume, fixture freeze, re-lock, T4 action, or external use. R17, rev16, the §D amendment, and interface lock remain unmoved. H-12 stands.

ACTIONS_GIT_REF: docs-workspace disk action — this DESIGN-REVIEW relay plus one append-only INDEX.md row; no design/source/frozen byte edited, no `frank/` action, no PLAN/T4/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; final `frank/` verification follows relay lint.
Next requested action: m-9.planner reports the pair-approved exact bytes upward for the reciprocal m-10, m-3, and m-1 close3 confirmations; all joint and downstream gates remain held.
