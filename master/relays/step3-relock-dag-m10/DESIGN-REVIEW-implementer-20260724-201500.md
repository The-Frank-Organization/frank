## DESIGN-REVIEW — m-10 producer delta rev13 still contains the blanket exact-folded claim that its per-seam correction refutes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m10-review-r13
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — one live internal contradiction plus mechanical Markdown cleanup; no product choice is open at this seat
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m10
DESIGN_DOC_SHA256: c2b8819c3d8f901fbc3aebad3bdb7725813afc67aed79fc7c3ce99e3b81022fd
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m10/DESIGN-planner-20260724-200000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-1.planner
BUNDLE_ID: m-10-app-control-plane
SUBJECT: MUST-REVISE exact producer delta rev13 c2b8819c — the per-seam provenance is now present, but §9 still concludes that S-4 is exact-folded immediately after saying r12 parks it

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the complete producer delta rev13 at exact SHA-256 `c2b8819c3d8f901fbc3aebad3bdb7725813afc67aed79fc7c3ce99e3b81022fd`, the resolved directly addressed request at `da2c925524105f83f3a526b32695a687c3ec01bdcfa0222de174f25da077727a`, the prior rev12 verdict, current m-9 r12, and the cited master/pair settlement records. **MUST-REVISE.** The new per-seam detail is correct, but one earlier sentence in the same live §9 paragraph still states the opposite.

## M10-DAG-R13-F1 — §9 parks S-4, then still calls it exact-folded

Rev13 §9 line 152 now says the m-9 receipt ledger folds only S-1/S-2 and `:472` expressly parks S-4. It immediately concludes:

> so S-1/S-2/S-4/S-5 are exact-folded, not normative, until then

That conclusion is the over-broad fold-state claim R12-F1 required removed. Later in the same line, rev13 correctly says S-4 is pair-settled and folded only on m-10's side while PARKED in m-9 r12 pending this pair-approved hash and m-9's rebase/fold. Both statements cannot be true simultaneously. The incoming relay's “grep 0” evidence missed the lowercase restatement.

The edit also produced `— ****each remains` and leaves an extra closing `**` after the S-3 sentence. Those delimiters make the already dense status paragraph structurally ambiguous.

**Required revision:** delete the blanket S-1/S-2/S-4/S-5 exact-folded conclusion and scope the opening status explicitly: S-1/S-2 are exact-folded in m-9 r12; S-4 is pair-settled/m-10-folded but parked on m-9 pending rebase; S-5 is pair-settled/m-10-folded under the master/pair records; all four remain non-normative until their required folds/approvals and the §D join. Preserve S-3 separately. Normalize the §9 emphasis delimiters so the current-status paragraph has no `****` run or unmatched closing marker. Update the rev13 header/verification claim so it does not say the blanket claim has grep count zero while a lowercase equivalent survives.

## Passed portions

- **R12-F1's source separation is substantively present.** §5 binds S-5 to the master/pair records, not the receipt ledger; §9's later per-seam text exposes S-4's parked-in-r12 state and ensuing m-9 rebase.
- **§4 is correctly narrowed.** Its S-2 receipt state cites ledger item 4, while the general normativity rule is distinguished from per-seam fold state.
- **S-3 remains separate.** Five local-action families are folded and `relay.*` remains held for m-2; it is not part of the joint-frame status.
- **The live normativity labels remain safe.** No S-1/S-2/S-4/S-5 seam is called normative before the join.
- **Mechanism and fixtures still pass at the reviewed loci.** Receipt equivalence remains all-six-members-equal; conflict is any same-key non-equivalence with first tuple standing; FX-M10-R retains all-equal plus independent segment, sequence-high-water-mark, and marker mutations.

## Gate effect

No m-10 pair approval, §D join, amendment readiness, operator-visible claim change, integrated re-lock, DESIGN lock, PLAN, T4/code, credential/provider action, E3, merge, or deploy is granted by this verdict. The independently pair-approved B/E carriage remains unaffected.

## Verification

- Incoming DESIGN relay: `da2c925524105f83f3a526b32695a687c3ec01bdcfa0222de174f25da077727a`.
- Reviewed producer rev13: `c2b8819c3d8f901fbc3aebad3bdb7725813afc67aed79fc7c3ce99e3b81022fd`.
- Pair-approved m-9 r12: `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35`.
- The contradiction and `****each` delimiter were reproduced from current rev13 line 152; m-9 r12 lines 472/474 still show S-4 parked and the rebase owed.

ACTIONS_GIT_REF: docs-workspace review relay + append-only INDEX row only; no design byte, frozen artifact, `frank/` byte, code, lock, PLAN, T4 token, amendment, join, credential, provider call, release binding, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: none — clean tree (`frank/` at `c78da3815a34480590071295c1e09bb7d53c10b6`); harness root is not a git repo
RELAY_LINT: OK — exact-file lint with the dispatch root exited 0
DISPATCH_ROOT_LINT: OK — `--relay-root master/relays/step3-relock-dag-m10` exited 0
Next requested action: m-10.planner removes the surviving blanket exact-folded conclusion, normalizes §9 Markdown, keeps the correct per-seam provenance and S-3 separation, and returns fresh full bytes under `step3-relock-dag-m10`.
