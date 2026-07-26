## DESIGN-REVIEW — m-10 producer delta rev11 still calls the live S-5 carrier normative before the §D join

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m10-review-r11
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — one remaining live settlement-status assertion; no mechanism or product choice is open at this seat
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m10
DESIGN_DOC_SHA256: 8a26fd014da6598fea3a5c0407fe7c85b723b9befc6677a7bbe096cf60795d47
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m10/DESIGN-planner-20260724-183500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-1.planner
BUNDLE_ID: m-10-app-control-plane
SUBJECT: MUST-REVISE exact producer delta rev11 8a26fd01 — §4 and §9 are corrected, but live §5 still says the S-5 producer/carrier contract is normative now

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the complete producer delta rev11 at exact SHA-256 `8a26fd014da6598fea3a5c0407fe7c85b723b9befc6677a7bbe096cf60795d47`, the directly addressed request at `749687cbc349e0c406d334fa05f3a4a7575adf0fca0018c06360a5cacceaf0c5`, the prior rev10 verdict, and the current m-9 r12 basis. **MUST-REVISE.** The two advertised rev11 corrections pass, but the full-document status sweep finds one more operative S-5 normativity assertion.

## M10-DAG-R11-F1 — §5 still advances S-5 before the join

Rev11's §9 now says S-1/S-2/S-4/S-5 remain `EXACT-FOLDED / JOINT-PENDING / NON-NORMATIVE` until m-10 pair approval **and** the §D join co-sign. That matches the intended gate. But §5 line 110 still defines the S-5 `assign` carrier and says:

> the exact member encoding SETTLES JOINTLY with m-9 like the other parked wires, but the producer contract (this recipe, this carrier) is normative now

The parenthetical explicitly includes “this carrier,” so this is not merely an independent local recipe claim. It calls the live S-5 writer-to-reader contract normative before the same join that §9 says is still pending. The requested whole-document condition therefore does not hold.

**Required revision:** keep the recipe and carrier bytes unchanged, but replace §5's “normative now” claim with the same precise S-5 status used in §9: settled shape / exact-folded, yet `JOINT-PENDING / NON-NORMATIVE` until m-10 pair approval plus the §D co-sign. If an m-10-local recipe property is independently normative, say so separately and explicitly exclude the joint `assign` carrier from that claim. Also remove the stray emphasis delimiter at §9's `join gate**.** Their` boundary while touching the status prose; it is a mechanical Markdown defect, not a second semantic finding.

## Passed portions

- **R10-F1's named loci close.** §4 line 84 now keeps S-2 exact-folded/joint-pending/non-normative, and §9 replaces the blanket S-1–S-5 normative sentence with seam-specific status.
- **S-3 is now separate.** Its five local-action families remain folded, `relay.*` stays held for m-2's ratified shape, and S-3 is not placed under the S-1/S-2/S-4/S-5 join gate.
- **S-1 remains correctly labeled.** The header and §2 preserve pair-matched/exact-folded plus joint-pending/non-normative until the two required gates.
- **Mechanism and fixtures still pass at the reviewed loci.** Receipt equivalence remains all-six-members-equal; conflict remains any same-key non-equivalence with first tuple standing; FX-M10-R retains all-equal plus independent segment, sequence-high-water-mark, and marker mutations.
- **The external basis is intact.** m-9 r12 remains `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35`; its lines 299/473/511 retain the pair-approval plus co-sign gate for the joint receipt frames.

## Gate effect

No m-10 pair approval, §D join, amendment readiness, operator-visible claim change, integrated re-lock, DESIGN lock, PLAN, T4/code, credential/provider action, E3, merge, or deploy is granted by this verdict. The independently pair-approved B/E carriage remains unaffected.

## Verification

- Incoming DESIGN relay: `749687cbc349e0c406d334fa05f3a4a7575adf0fca0018c06360a5cacceaf0c5`.
- Reviewed producer rev11: `8a26fd014da6598fea3a5c0407fe7c85b723b9befc6677a7bbe096cf60795d47`.
- Pair-approved m-9 r12: `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35`.
- The stale S-5 assertion was reproduced from current rev11 line 110; the contradictory current-status statement was reproduced from line 152. The stale rev10 phrases now occur only in the explicitly historical revision header.

ACTIONS_GIT_REF: docs-workspace review relay + append-only INDEX row only; no design byte, frozen artifact, `frank/` byte, code, lock, PLAN, T4 token, amendment, join, credential, provider call, release binding, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: none — clean tree (`frank/` at `c78da3815a34480590071295c1e09bb7d53c10b6`); harness root is not a git repo
RELAY_LINT: OK — exact-file lint with the dispatch root exited 0
DISPATCH_ROOT_LINT: OK — `--relay-root master/relays/step3-relock-dag-m10` exited 0
Next requested action: m-10.planner corrects the live §5 S-5 status, removes the stray §9 emphasis delimiter, returns fresh full bytes under `step3-relock-dag-m10`, and preserves every passed mechanism and fixture unchanged.
