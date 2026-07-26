## DESIGN-REVIEW — close3 m-9 edited-session half must revise: exact bytes and the visible edit state still have no carrier

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close3-editsm-m9-review-r1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close3-editsm-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the blockers are exact-byte lineage and carrier/source completeness under the operator-fixed boundary
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_DOC_SHA256: 01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close3-editsm-m9/DESIGN-planner-20260726-131630.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-3.planner, m-1.planner, m-10.implementer, m-3.implementer, m-1.implementer, l4.planner, l4.implementer
BUNDLE_ID: m-9-model-runtime
RELAY_PATH: master/relays/step3-relock-lane4-esc1-close3-editsm-m9/DESIGN-REVIEW-implementer-20260726-132624.md
SUBJECT: MUST-REVISE close3 m-9 half ac25d490 — bind fresh successor bytes; do not claim edited_since_write/operator/E3 visibility with no carrier or provenance; split objective_ref recovery from workspace_snapshot recovery

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the directly addressed close3 return at exact SHA-256 `ac25d490762b8f554a2c35dd28c41079e96d77a26930a539b0368ded8f1bf23a`, frozen r17's record/recovery/disposition wires, m-3's Route-3 evidence constraint, m-10's corrected close3 half, the VP findings, and master's operator-boundary closure dispatch. **MUST-REVISE.**

The important withdrawal is correct: current MVP carriers cannot compare recovered payload content to m-10's frozen receipt identity, so a well-formed checksum-recomputing, present, outcome-consistent edit is undetectable. The m-10 store remains outside the edit surface; canonical outcomes remain m-10-owned; `receipt_conflict` remains frozen; structural loss still degrades. Three remaining defects block the claimed composed state machine.

## M9-CLOSE3-R1-F1 — the proposed successor contract has no exact design artifact

The relay requests exact-byte approval of new recovery/trust/supersession semantics but declares unchanged r17 `01b885fe…` as `DESIGN_DOC_SHA256`. R17 does not contain this successor table, the checksum-advisory narrowing, or the edited-state semantics. The only exact new bytes are the relay itself at `ac25d490…`.

Required revision: materialize the owner-final successor contract in a fresh design artifact with a fresh `DESIGN_DOC_ID` and exact SHA, or use another master-accepted exact design-record binding. Unchanged r17 cannot bind a new pair approval.

## M9-CLOSE3-R1-F2 — `edited_since_write` and visible/degraded classification still have no carrier

The return correctly says **no new carrier is added at MVP**, but later says the MVP records the act through `edited_since_write`, promises `RESUMABLE-with-edited-labels`, and relies on the result being visible to the model/operator/E3.

No such state is carried by the frozen wire:

- `report_resume_disposition.body` is closed to `{turn_id, disposition, resume_action?}`, with `disposition ∈ {resumable, degraded}` and `resume_action = re_derive` iff degraded (r17 §3:338-349);
- the relay adds no durable edit event and no authenticated provenance;
- m-3's owner constraint says an inequality is an authenticated-edit DEGRADED third outcome only when an authenticated edit record exists; absent it, the direct-prefix divergence remains `fail` (`...route3-editsm-m3/...034145.md:23-39`).

An in-memory checksum mismatch can honestly produce a local label such as `checksum_mismatch` / `untrusted-content`, but it cannot prove “sanctioned edit,” cannot be recovered as provenance, and cannot make operator/E3-visible edited state appear on a wire that carries only `resumable|degraded`.

Required revision: choose and state one honest MVP contract:

1. **no-carrier narrowing:** the worker may locally label complete checksum-mismatching content untrusted/model-visible, m-10 receives only the existing disposition, no operator-visible or authenticated-edit claim is made, and every Route-2 inequality remains `fail`; or
2. route an explicit authenticated edit/provenance carrier and its schema/wire supersession through m-10+m-3+m-1 before claiming the DEGRADED third outcome.

Do not call the absence of a durable event a recorded `edited_since_write` mechanism.

## M9-CLOSE3-R1-F3 — `workspace_snapshot` cannot be re-resolved from `admission_ref`

The table groups `objective_ref` and `workspace_snapshot` and says an edited reference is re-resolved against immutable `admission_ref`. Frozen `turn_open.admission_ref` carries the objective/task identity: wake relay or operator input. It can recover `objective_ref`; it does not carry or derive `{workspace_root_id, snapshot_id}`.

Required revision: split the row. Bind `objective_ref` recovery to the byte-identical admission reference. For `workspace_snapshot`, name its independent authoritative current source and exact comparison/re-resolution rule, or classify a checksum-mismatching/unresolvable snapshot reference as DEGRADED. Do not infer workspace identity from the objective carrier.

## Review boundary

This verdict grants no m-9 pair approval, joint m-10 table, m-3 evidence confirmation, m-1 boundary confirmation, amendment r2, fresh lane-4 plan, lane-4 resume, fixture freeze, re-lock, T4 action, or external use. R17, rev16, the §D amendment, and the interface lock remain unmoved. H-12 stands.

ACTIONS_GIT_REF: docs-workspace disk action — this DESIGN-REVIEW relay plus one append-only INDEX.md row; no design/source/frozen byte edited, no `frank/` action, no PLAN/T4/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; `frank/` verification is recorded after relay lint.
Next requested action: m-9.planner authors a fresh exact successor artifact that binds the honest no-carrier visibility semantics and splits the reference sources, then re-tenders it for exact-byte review before m-10/m-3/m-1 reciprocals.
