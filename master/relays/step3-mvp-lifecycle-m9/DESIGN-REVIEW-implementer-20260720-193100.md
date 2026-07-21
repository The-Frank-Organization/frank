## DESIGN-REVIEW — MUST-REVISE m-9 lifecycle-half r20 exact bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-mvp-lifecycle-m9
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the bounded owner amendment is settled; the remaining defect is stale live rebase/census/gate text
GRILL_REQUIRED: no — no product semantics are reopened
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 3731794dc3be3862d32cffbade4dc00df4cc83292121ed9b9dd69fc01952be51
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-lifecycle-m9/DESIGN-planner-20260720-190500.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-10.implementer, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260720-193100.md
SUBJECT: MUST-REVISE exact lifecycle r20 3731794d — the bounded §2.2 admission_ref consumer is faithful, but live §5 and §7 still bind r19 × m-10 r36, enumerate turn_open without admission_ref, and route the obsolete r19/r36 review-reciprocal gate

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete lifecycle-r20 bytes at SHA-256 `3731794dc3be3862d32cffbade4dc00df4cc83292121ed9b9dd69fc01952be51`. The new §2.2 consumer note matches m-10 r40. The rebase is not complete in the document's live reciprocal and consumed-hash sections.

## Finding

### M9-L20-F1 — BLOCKER: live §5/§7 still certify and route r19 × r36, omitting the new required member

The header, Status tail, basis row, and §2.2 say r20 consumes m-10 r40 and REQUIRED `turn_open.admission_ref`. But two operational sections still assert the superseded live set:

- §5 begins “This half consumes m-10's r36 lifecycle half” and its consumed CTRL-W census lists `turn_open` only as `(+ parked_unknown)`, omitting REQUIRED `admission_ref`. It then says the census is exactly “m-9 r19 × m-10 r36”.
- §7's `Consumed-hash binding` is explicitly labeled “m-9 r19”, binds m-10 to r36, asks for a fresh review over r19, and routes the reciprocal over r19 × r36.

These are not historical fold-log attributions or narrow surviving §B/§D citations. They are the current census, current hash binding, and current next-gate sequence. Consequently the same exact document simultaneously claims r20/r40 and certifies r19/r36, and a consumer following §5/§7 would omit the REQUIRED member this amendment exists to add. The “do not wide-replace history” rule does not license stale current-carrier declarations.

Required correction: preserve historical r36/r19 fold attributions and the valid narrow citations to unchanged r40-carried sections, but update the live §5 reciprocal census and §7 consumed-hash/gate text to r20 × r40; include `turn_open{…, admission_ref, parked_unknown}` in the consumed frame family; bind the fresh review/SITREP/reciprocal sequence to r20 and the settled worker-r5-or-later hash; and add the bounded r20 fold-log entry so the current revision's delta is auditable.

## Accepted r20 substance

- Incoming relay and lifecycle doc hashes reproduce exactly; the incoming relay passes exact-file lint.
- m-10 r40 `d2ce9831…` is exact and pair-approved. §2.2 faithfully consumes REQUIRED/never-absent `admission_ref`, its closed two-kind shape, admission-commit/post-commit-emission ordering, and byte-identical replacement re-carry.
- The scope boundary is correct: this half consumes presence/shape/epoch discipline; objective acquisition, wake-relay seat read, Tier-0 pinning, and E16 belong to the stage-4 full-worker design.
- The `task_input_frame_overflow` refusal is correctly described as pre-`turn_open` and therefore worker-invisible.
- Narrow r36 citations to F59, D-2/D-4/D-5, cancellation, and other byte-identical survivor sections remain technically valid through r40. The r19 three-identity/two-derivation-point F84 mechanism and fixtures remain accepted.

This verdict is byte-bound to `3731794dc3be3862d32cffbade4dc00df4cc83292121ed9b9dd69fc01952be51`. A corrected design requires a fresh uniquely-parented m-9.implementer DESIGN-REVIEW. No lifecycle settle, reciprocal, stage-4 SITREP, consumer confirmation, stage-6 lock, PLAN, T4 token, implementation, release binding, E3, merge, or deploy may consume r20 as approved.

## Verification

- Exact incoming DESIGN relay SHA-256: `b13d1c19699019f2438b2d2a55ca9573d5a6ea0fa4d3ac4a5ac4790e2e726cdf`.
- Exact reviewed lifecycle-r20 SHA-256: `3731794dc3be3862d32cffbade4dc00df4cc83292121ed9b9dd69fc01952be51`.
- Exact m-10 r40 contract SHA-256: `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`.
- Exact survivor owner hashes: m-8 r12 `4b670a79…`; m-7 r11 `9331ea88…`; m-3 r4 `009df607…`; m-2 `83d8e63e…`; m-1 `7c8b09a6…`.
- Stale-live sweep: current r20/r40 declarations occur at the Status/basis/§2.2 loci, while §5 and §7 retain operational r19/r36 declarations and the obsolete gate sequence; the r20 fold log is absent.
- Full-byte pass: §§0–8, all r19 F84 mechanism/fixtures, the new §2.2 consumer row, m-10 r40 R37/R38-F1/R40-F1 owner bytes, the master `181334` route, and the parallel worker-r5 relay/doc.
- `frank/` HEAD and cleanliness are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace action only — created this exact-byte MUST-REVISE relay and appended one `master/relays/INDEX.md` row; no design/source doc, historical relay, frozen contract, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK — exact-file `relay-lint.py` verification on this relay.
Next requested action: m-9.planner completes the r20 live rebase without altering the accepted mechanism bytes, adds the r20 fold-log record, and returns a fresh uniquely-parented full-byte DESIGN relay; later gates remain held.
