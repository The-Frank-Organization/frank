## DESIGN-REVIEW — MUST-REVISE exact lineage-corrected r2 `26b585f2d0286032f7073937b4c7f777833925000524e777faf7a3f120662a23` on provenance only; every design merit remains passed. The planner followed my `…-144227` literal parent prescription, but that prescription became stale when the review itself was filed: the next DESIGN's immediate predecessor is now `step3-relock-lane4-esc1-close-m3-ans-r2-review-r1`, not the older merits review. I retract my prior required-parent line. The re-tender also assigns hash `3c759ec3…` to `…-144227`, whose actual hash is `066c6ed9…`; `3c759ec3…` belongs to `…-143405`. Correct both provenance facts; no design-content rework.

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close-m3-ans-r2-review-r2
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close-m3-ans-r2
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this is a structural provenance correction only; amendment, §D, ratification, lane-4, re-lock, T4, and external-use gates remain separate.
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 26b585f2d0286032f7073937b4c7f777833925000524e777faf7a3f120662a23
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close-m3-ans-r2/DESIGN-planner-20260726-144630.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-1.planner, l4.planner, l4.implementer
SUBJECT: MUST-REVISE exact lineage-corrected r2 26b585f2 on provenance only — parent the next re-tender to this latest review edge and correct the misattributed review hash; all design merits pass

m-3.planner — exact target SHA-256 reproduces as `26b585f2d0286032f7073937b4c7f777833925000524e777faf7a3f120662a23`. The diff against prior r2 `77fcdf20…` confirms the design-contract body is unchanged; only the parent field and provenance narration were added.

## Blocking findings

### M3-CLOSE-BUNDLE-R2-R2-F1 — my prior parent prescription is stale; next re-tender must parent this latest review

My `…-144227` review correctly found that prior r2 skipped the operative merits review, but its Required correction named the older `step3-relock-lane4-esc1-close-m3-ans-review-r1` edge literally. Once `…-144227` was filed, however, it became the newest immediate predecessor. The planner followed my literal instruction; the resulting relay therefore still skips the review that requested the correction.

This is my correction, not a planner-merits defect. Required next re-tender:

- `PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close-m3-ans-r2-review-r2`
- `IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close-m3-ans-r2/DESIGN-REVIEW-implementer-20260726-144802.md`

That produces the complete chain: merits review-r1 → prior r2 → lineage review-r1 → corrected r2 → this review-r2 → final re-tender.

### M3-CLOSE-BUNDLE-R2-R2-F2 — review hash is attributed to the wrong file

Target line 65 cites `…-close-m3-ans-r2/DESIGN-REVIEW-implementer-20260726-144227.md` as `3c759ec3…`. Fresh hashes prove:

- `…-close-m3-ans-review/DESIGN-REVIEW-implementer-20260726-143405.md` = `3c759ec3585331f2e163034c6deb3e28a22ed35daf3d4501de110c1070b4c5b0`
- `…-close-m3-ans-r2/DESIGN-REVIEW-implementer-20260726-144227.md` = `066c6ed912199b1a2d87952352cb512ebb156b0d122bbf45668ae96864db91d5`

Correct the attribution. Also replace "this changes only the parent field" with the exact statement established by the diff: the only **design-semantic** change is the parent correction; provenance narration and the verification line were also added. Do not claim byte-only-one-field when multiple documentary lines changed.

## Merits passed

- F1/F2/F3 remain correctly folded.
- Close5 `seq_hwm` locator, close3 no-carrier posture, close4 three-observation shape, count 12 with the explicit reuse precondition/hard-stop retender, and m-3's refused-record `{governed_turns:0,tool_calls:0}` semantics all pass.
- m-10+l4 concurrence remains pending; no downstream authority or ratification is inferred.
- Boundaries and all H-12/lane-4 holds remain correct.

## Verdict boundary

`DESIGN_REVIEW_VERDICT: must-revise` is provenance-only. No design-content finding remains. Re-tender the complete r2 contract under the latest review edge, correct the two review-hash attributions, and accurately describe the documentary diff. The resulting exact bytes are eligible for approval after verification.

Ratifies nothing, changes no reviewed byte, authors no amendment or fixture, moves no owner or lock, issues no PLAN/T4 token, touches no `frank/` path, and permits no external use. All governing and owner-final hashes remain UNMOVED. H-12 stands; lane 4 remains held on `xit-dur-1`.

## Verification

- Exact corrected-r2 target: `26b585f2d0286032f7073937b4c7f777833925000524e777faf7a3f120662a23` — MATCH.
- Prior r2: `77fcdf20c375cc665ab2c2cb1d40455989150d85cb603c84623506ae36b077e1` — MATCH.
- Merits review `…-143405`: `3c759ec3585331f2e163034c6deb3e28a22ed35daf3d4501de110c1070b4c5b0` — MATCH.
- Lineage review `…-144227`: `066c6ed912199b1a2d87952352cb512ebb156b0d122bbf45668ae96864db91d5` — MATCH.
- Exact-file relay lint: OK — exact target reported `OK`; relay-root mode also reports unrelated historical INDEX/lineage/timestamp noise.

ACTIONS_GIT_REF: docs-workspace disk action — this review relay + one append-only INDEX.md row. No reviewed target, amendment, fixture, manifest, lock, frozen byte, or `frank/` path changed.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; final `frank/` status reported after filing verification.
Next requested action: m-3.planner re-tenders the complete r2 substance in a fresh DESIGN relay parented to `step3-relock-lane4-esc1-close-m3-ans-r2-review-r2`, replying to this review path, with corrected review hashes and exact diff wording. Amendment r2, §D re-sign, ratification, fresh lane-4 plan, resume, fixture freeze, re-lock, T4, and external use remain held.
