## DESIGN-REVIEW — correction and must-revise m-1 lane-4 edited-session boundary rev2

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close3-editsm-m1-review-r2-correction-1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close3-editsm-m1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the correction is a determinate cross-owner vocabulary repair; amendment-r2 composition and ratification remain Master+VP+operator gates
GRILL_REQUIRED: no — no product choice is opened by this correction
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close3-editsm-m1/DESIGN-planner-20260726-133133.md
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-3.planner, m-9.implementer, m-10.implementer, m-3.implementer, l4.planner, l4.implementer
BUNDLE_ID: m-1-trust-identity
OWNER: m-1 (Trust & Identity)
RELAY_PATH: master/relays/step3-relock-lane4-esc1-close3-editsm-m1/DESIGN-REVIEW-implementer-20260726-133821.md
SUPERSEDES_RELAY: master/relays/step3-relock-lane4-esc1-close3-editsm-m1/DESIGN-REVIEW-implementer-20260726-133628.md
REVOKED_RELAY_SHA256: 62e70099cf35eb3f9bd46500e1b805cb672dc4885a6131af0e71706434847075
SUBJECT: correction — revoke the premature r2 approval and require m-1 to preserve m-9's distinct edited versus content_lost classes

DESIGN_REVIEW_VERDICT: must-revise

## Revocation

The earlier relay `DESIGN-REVIEW-implementer-20260726-133628.md` at exact SHA-256 `62e70099cf35eb3f9bd46500e1b805cb672dc4885a6131af0e71706434847075` is **WITHDRAWN, REVOKED, and INERT**. It was authored before the concurrent m-9 and m-10 pair-review bytes were consumed. It MUST NOT satisfy owner-final review, amendment-fold, PLAN-lineage, or any downstream approval gate, and no successor may parent an authority claim to it.

This correction supersedes that relay's verdict. The current verdict on the exact m-1 rev2 planner return at SHA-256 `554ab21e236e246230479fd97276cf3a7a800d1787dd1242c7a52a51f33ec5fd` is **must-revise**.

The rev2 return correctly closes M1-CLOSE3-R1-F1 through F3 except for one residual cross-owner classification error. BR-INV prime, the narrowed edit surface, the honest undetectable checksum-recomputed limit, the three guarantee-grain answers, zero m-1 MVP supersession, repairability scope, the m-9+m-10 Step-4 pair, frozen hashes, `receipt_conflict`, H-12, and all downstream holds remain accepted.

## Finding

### M1-CLOSE3-R2-F1 — advisory-checksum mismatch is `edited`, not `content_lost`

The m-1 rev2 groups advisory-checksum mismatch with structural/missing/file-integrity faults, then says all such detections enter m-9's `content_lost`/`DEGRADED` handling (`DESIGN-planner-20260726-133133.md:31-36`). That collapses two distinct m-9-owned classes.

The current m-9 planner bytes distinguish them exactly (`…close3-editsm-m9/DESIGN-planner-20260726-131630.md:40-47,56-59`):

- structural/completeness failure or missing referenced content is `content_lost` and enters `DEGRADED + re_derive`;
- a complete, present, well-formed record with an advisory-checksum mismatch is `edited`, untrusted-but-model-visible, and is not `content_lost`;
- a complete, present, consistent edit with its checksum recomputed remains undetected and may resume as clean under the accepted MVP limit.

The m-9 implementer separately marked the prior planner successor must-revise because its durable/wire carrier and provenance claims were not supported. A fresh m-9 successor then arrived during this review (`DESIGN-planner-20260726-133740.md`, design artifact `30020fa6…`): it narrows checksum mismatch to a local, in-memory, model-visible-only class and expressly retracts any durable/wire or operator/E3-visible edited state. It remains pair-review pending. Neither state authorizes m-1 to replace the distinct m-9 vocabulary with `content_lost`. m-1 owns the trust/isolation boundary, not the edited-session disposition or carrier.

## Required revision

Return fresh exact bytes that:

1. Preserve every accepted rev2 correction and all existing holds.
2. Split the detected branch: structural/completeness failure and missing content may enter m-9's `content_lost`/`DEGRADED` path; a complete, present, well-formed advisory-checksum mismatch remains the distinct m-9-owned `edited` / untrusted-but-model-visible class and is not `content_lost`.
3. Keep the checksum-recomputed present/consistent class as the undetected clean-resume MVP limit.
4. State that final disposition, label visibility, and any durable or wire carrier are m-9/m-10-owned and remain subject to fresh pair-approved successor bytes. Do not promise operator-visible, E3-visible, or durable edited labels from m-1.
5. Issue a fresh uniquely-parented DESIGN for exact-byte review. Do not call rev2 owner-final, and do not use the revoked approval as lineage.

No human decision is required for this owner-local correction. Master and m-1.planner must treat the revoked approval as invalid and keep amendment-r2 composition, owner-final fold, lane resume, fixture materialization/freeze, re-lock, PLAN, T4/code, external use, credential, provider action, live E3, `frank/` edits, merge, and deploy held.

## Verification

- Reproduced exact SHA-256: reviewed m-1 rev2 `554ab21e236e246230479fd97276cf3a7a800d1787dd1242c7a52a51f33ec5fd`; revoked review `62e70099cf35eb3f9bd46500e1b805cb672dc4885a6131af0e71706434847075`; prior m-9 planner half `ac25d490762b8f554a2c35dd28c41079e96d77a26930a539b0368ded8f1bf23a`; fresh pair-review-pending m-9 relay `e7a77949d1f5e03b5b7246fdf312593e1c94be38e9d066d1dfb0ce89385cff44`, binding design artifact `30020fa6a0697169ca91e5b8501f2c98d6464b92e098bc00ed6b7f100c9952da`.
- Re-read the concurrent m-9 implementer must-revise, m-10 implementer cross-owner finding, and fresh no-carrier m-9 successor before finalizing this correction.
- Exact-file and dispatch-root relay-lint pass. The index row is unique; full index lint remains blocked by append-only ordering defects at lines 2412, 2418, and 2420, including two concurrent rows appended after this correction with earlier timestamps.

ACTIONS_GIT_REF: docs-workspace action only — appended this correction relay and one `master/relays/INDEX.md` row; the revoked relay and its historical index row were not mutated; no design/source contract, frozen byte, `frank/` source, branch, commit, amendment, ratification, lock, PLAN, T4, credential, provider action, external use, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` was `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK — exact-file and dispatch-root relay-lint.py exit 0 (`--no-freshness` used for dispatch-root historical files)
INDEX_LINT: pre-existing/concurrent failure — line 2412 retains the earlier timestamp-order defect; after this row was appended, concurrent writers added older-timestamp rows causing additional failures at lines 2418 and 2420; this correction's row is unique and no historical row was rewritten
Next requested action: m-1.planner folds M1-CLOSE3-R2-F1 only, preserves every accepted rev2 correction, re-hashes the standalone return, and sends fresh exact bytes for pair review; master treats the earlier approval as revoked and keeps every downstream gate held.
