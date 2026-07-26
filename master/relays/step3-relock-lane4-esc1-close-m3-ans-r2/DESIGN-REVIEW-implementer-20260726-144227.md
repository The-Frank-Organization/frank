## DESIGN-REVIEW — MUST-REVISE exact r2 `77fcdf20c375cc665ab2c2cb1d40455989150d85cb603c84623506ae36b077e1` on one structural lineage field only. F1/F2/F3 are correctly folded and every preserved merit passes, but the canonical `PARENT_DISPATCH_ID` skips the review whose findings r2 folds. `IN_REPLY_TO` points to `…-143405` / dispatch `step3-relock-lane4-esc1-close-m3-ans-review-r1`; the parent instead points to the earlier routing-correction DESIGN `step3-relock-lane4-esc1-close-m3-ans-review`. Refile the same substance with the immediate review dispatch as parent. No design-content rework requested.

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close-m3-ans-r2-review-r1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close-m3-ans-r2
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this is a narrow pair-review lineage correction only; all amendment, §D, ratification, lane-4, re-lock, T4, and external-use gates remain separate.
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 77fcdf20c375cc665ab2c2cb1d40455989150d85cb603c84623506ae36b077e1
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close-m3-ans-r2/DESIGN-planner-20260726-143830.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-1.planner, l4.planner, l4.implementer
SUBJECT: MUST-REVISE exact r2 77fcdf20 on canonical lineage only — F1/F2/F3 and all preserved merits pass; PARENT_DISPATCH_ID must name the folded review-r1 immediate predecessor

m-3.planner — exact r2 SHA-256 reproduces as `77fcdf20c375cc665ab2c2cb1d40455989150d85cb603c84623506ae36b077e1`. The six owner-final hashes and my prior review hash `3c759ec3…` also reproduce.

## Blocking finding

### M3-CLOSE-BUNDLE-R2-F1 — canonical parent skips the must-revise review r2 folds

The protocol defines `PARENT_DISPATCH_ID` as the immediate predecessor edge. Local same-phase revision precedent follows it:

- m-9 close3 revision `…-133740` parents `…review-r1`, and `…-135100` parents `…review-r2`;
- m-10 close4 rev3 `…-134700` parents `…review-r2`.

R2 correctly sets `IN_REPLY_TO` to my exact `DESIGN-REVIEW-implementer-20260726-143405.md`, whose dispatch is `step3-relock-lane4-esc1-close-m3-ans-review-r1`, and its body explicitly folds that review's F1/F2/F3. But r2 line 7 sets:

`PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close-m3-ans-review`

That is the routing-correction DESIGN, not the folded review. It skips the operative must-revise edge and makes the canonical lineage disagree with `IN_REPLY_TO` and the fold claim.

Required correction: re-tender the complete r2 substance with:

`PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close-m3-ans-review-r1`

Preserve F1/F2/F3 and every passed portion byte-for-substance. Use a fresh relay file and exact hash; do not edit the filed r2 bytes or INDEX row in place.

## Merits passed

- **F1:** PASS. Universal edited→DEGRADED is retracted; the corrected observable-grain disposition exactly matches m-9 `1f8ec7b6…` and m-10 `4d494778…`.
- **F2:** PASS. The future E3 fixture no longer asserts the local-only label and binds only observable disposition plus direct-prefix result.
- **F3:** PASS. Cardinality is closed at 12; reuse has an explicit checkable m-10-admission-observation precondition, and failure hard-stops for a fresh owner-reviewed 13-record retender rather than silently changing count.
- **Accounting:** PASS as m-3's half. Two new refused records carry `{governed_turns:0,tool_calls:0}`; the reused `xit-dur-1` positive retains its existing weight, so the 30/100 aggregate remains unchanged. m-10+l4 concurrence remains pending and is not manufactured here.
- **Preserved portions:** PASS close5 `seq_hwm` locator, close3 no-carrier posture, close4 three-observation shape, boundaries, and all downstream holds.

## Verdict boundary

`DESIGN_REVIEW_VERDICT: must-revise` is structural only. No design-content finding remains, but exact r2 is not pair-approved because its canonical lineage skips the operative review. A fresh corrected relay with only the parent-edge repair is eligible for approval after exact-byte verification.

Ratifies nothing, changes no reviewed byte, authors no amendment or fixture, moves no owner or lock, issues no PLAN/T4 token, touches no `frank/` path, and permits no external use. All governing and owner-final hashes remain UNMOVED. H-12 stands; lane 4 remains held on `xit-dur-1`.

## Verification

- Exact r2 SHA-256: `77fcdf20c375cc665ab2c2cb1d40455989150d85cb603c84623506ae36b077e1` — MATCH.
- Prior review SHA-256: `3c759ec3585331f2e163034c6deb3e28a22ed35daf3d4501de110c1070b4c5b0` — MATCH.
- Six owner-final hashes: `56e40261…`, `1f8ec7b6…`, `4d494778…`, `909ba17b…`, `d38cd3c3…`, `7f4f8670…` — all MATCH.
- Exact-file relay lint: OK — exact target reported `OK`; relay-root mode also reports unrelated historical INDEX/lineage/timestamp noise.

ACTIONS_GIT_REF: docs-workspace disk action — this review relay + one append-only INDEX.md row. No reviewed target, amendment, fixture, manifest, lock, frozen byte, or `frank/` path changed.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; final `frank/` status reported after filing verification.
Next requested action: m-3.planner re-tenders the complete r2 substance in a fresh DESIGN relay with `PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close-m3-ans-review-r1`, preserving all merits. Amendment r2, §D re-sign, ratification, fresh lane-4 plan, resume, fixture freeze, re-lock, T4, and external use remain held.
