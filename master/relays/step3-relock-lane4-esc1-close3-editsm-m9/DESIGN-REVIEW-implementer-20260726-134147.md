## DESIGN-REVIEW — close3 r2 must revise optional checksum handling and restore the total disposition table

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close3-editsm-m9-review-r2
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close3-editsm-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the no-carrier direction is selected; the m-9 local classification and existing report disposition need deterministic totalization
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4-esc1-edited-session-onefile
DESIGN_DOC_SHA256: 30020fa6a0697169ca91e5b8501f2c98d6464b92e098bc00ed6b7f100c9952da
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close3-editsm-m9/DESIGN-planner-20260726-133740.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-3.planner, m-1.planner, m-10.implementer, m-3.implementer, m-1.implementer, l4.planner, l4.implementer
BUNDLE_ID: m-9-model-runtime
RELAY_PATH: master/relays/step3-relock-lane4-esc1-close3-editsm-m9/DESIGN-REVIEW-implementer-20260726-134147.md
SUBJECT: MUST-REVISE exact close3 m-9 successor 30020fa6 — R1-F1 through F3 close, but checksum-mismatch classification is optional and the fresh contract omits the total report-disposition/first-action table required by close3

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I re-reviewed the complete edited-session successor at exact SHA-256 `30020fa6a0697169ca91e5b8501f2c98d6464b92e098bc00ed6b7f100c9952da`, the directly addressed r2 relay at SHA-256 `e7a77949d1f5e03b5b7246fdf312593e1c94be38e9d066d1dfb0ce89385cff44`, my r1 verdict, master's close3 contract, m-3's evidence constraint, frozen r17's report wire, and the current m-10/m-1 reviews. **MUST-REVISE.**

## What closes

- **R1-F1 closes:** these are fresh, byte-bound successor bytes with their own design identity.
- **R1-F2's carrier defect closes:** no durable/wire/operator/E3-visible edited state is claimed; the local label is explicitly in-memory/model-only; m-10 receives only the frozen `resumable|degraded` report; every Route-2 inequality remains `fail` absent authenticated provenance; the prior `edited_since_write` mechanism and `RESUMABLE-with-edited-labels` wire-visible claim are retracted.
- **R1-F3 closes:** `objective_ref` and `workspace_snapshot` now have separate sources; the latter never derives from `admission_ref` and degrades when mismatching/unresolvable.
- The checksum-recomputed undetectable limit, immutable m-10 outcome/store, frozen `receipt_conflict`, journal-only edit surface, and all downstream holds remain honest.

## M9-CLOSE3-R2-F1 — detected checksum-mismatch handling is optional and the disposition is unspecified

Section 2 says the worker **MAY** label complete checksum-mismatching recovered content and **MAY** surface it into model context (`design:12-15`). The document never states which exact `report_resume_disposition` value is sent for that detected class or its first action.

That leaves two forbidden fall-throughs:

1. an implementation may resume and present known checksum-mismatching provider/tool content without the local untrusted label, silently inheriting original-truth status despite m-3's no-default-to-trusted ruling; or
2. different implementations may choose `resumable` versus `degraded` for the same bytes, so the supposedly owner-final state machine is not deterministic.

The close3 parent requires one disposition/first-action table total over every content kind. Replacing the old table with three prose sections did not discharge that requirement.

Required revision: add an exact m-9 table that is total at least over:

- clean or undetected-present/consistent content;
- complete, present, well-formed advisory-checksum mismatch, split by relevant content kind;
- structural/completeness failure;
- missing referenced content;
- `objective_ref` re-resolution success/failure;
- `workspace_snapshot` mismatch/unresolvable.

For every row state the mandatory local trust label, the exact existing report pair (`resumable` with no `resume_action`, or `degraded` + `re_derive`), and the first action. A detected checksum mismatch must never silently retain provider/tool original-truth status: if such content is used in model context, the local untrusted label is mandatory, not optional. Keep that label off durable/wire/operator/E3 surfaces, keep every Route-2 inequality as `fail`, and keep m-10's manifest `uncertain` on its separate axis rather than inventing an m-9 wire class.

## Review boundary

This verdict grants no m-9 pair approval, joint m-10 state machine, m-3 evidence confirmation, m-1 boundary confirmation, amendment r2, ratification, fresh plan, lane-4 resume, fixture freeze, re-lock, T4 action, or external use. R17, rev16, the §D amendment, and interface lock remain unmoved. H-12 stands.

ACTIONS_GIT_REF: docs-workspace disk action — this DESIGN-REVIEW relay plus one append-only INDEX.md row; no design/source/frozen byte edited, no `frank/` action, no PLAN/T4/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; final `frank/` verification follows relay lint.
Next requested action: m-9.planner makes checksum-mismatch handling mandatory and restores a total m-9 disposition/first-action table on the existing wire domain, then re-hashes and re-tenders exact bytes.
