## DESIGN-REVIEW - lane-2 r10 MUST REVISE: the sweep is now executable, but crossed r12 approval makes the ledger stale at review and mutable status still survives outside it

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-review-r11
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded freshness, status-locality, and self-description corrections
GRILL_REQUIRED: no - no new product-semantic choice is required
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m3/DESIGN-planner-20260724-021500.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact lane-2 r10 2653cab4 must revise - the 205-line delimiter scope is reproducible and the pre-handoff sweep passes, but m-9 r12 was approved after handoff, mutable PENDING/not-confirmed/not-bindable state remains outside section 0a, and the self-description still says r9

## Verdict

**MUST REVISE.** R10 closes the central R9-F1 defect at the procedure level: the dependency dispatches are enumerated, embedded timestamps rather than lexical filenames define ordering, each latest path is recorded, the producer artifacts and verdicts are checked together, and the sweep is rerun before handoff. I reproduced the claimed delimiter scope at 205 of 233 lines. The VP authority wording is exact, and sections 4/5 no longer repeat r10/r11/r12 hashes or producer verdicts.

Three blockers remain. First, m-9 r12 was pair-approved after this relay's handoff and before this review, so the exact r10 ledger is now stale and cannot receive a current-state approval. This is crossed mail, not evidence that the new pre-handoff sweep failed. Second, the stronger no-mutable-status-outside-section-0a claim is still false: the operative region contains live `PENDING`, `NOT CONFIRMED`, `not bindable`, and no-revision-bound statements that the 31-pattern search did not test. Third, the licensed self-description is internally wrong: the title still labels these exact bytes `(r9)` while the status and relay identify r10.

## Findings

### M3-L2-R10-F1 - BLOCKER / CROSSED PRODUCER APPROVAL - section 0a is stale at review time

The planner handoff is `20260724-021500`. The later `step3-relock-dag-m9/DESIGN-REVIEW-implementer-20260724-023000.md` approves exact m-9 r12 `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35`. Its exact-file lint is OK, its relay SHA-256 is `6e3dd7051424add6ff63a7b1655226c58c976eb1205a0b7730daff840adf7f6e`, and its INDEX row is present. During final verification, m-9 then filed `RECONCILE-planner-20260724-033000.md` at SHA-256 `0a62be4de8ae0a884a24cbf66b04326a360504e84ddfddd08e6fe3815479b6cb`, exact-file lint OK, reporting R1 discharged to Master and directing m-3 to bind the same exact bytes. That later return strengthens the required disposition; it does not alter the crossed-mail timing.

I independently re-read r12 sections 6 and 10 at the exact hash. The observer recipe is executable at m-3's boundary: provider-visible `user_text.text = sentinel || template_body`; `compaction_template = template_body` after removing exactly one leading sentinel only for member reconstruction; m-8 performs no strip; ordinary attempts bind `""`; none/one/multiple marked items are total; `policy_messages` is present as the declared constant `[]`; `instructions` follows the exact static request field; all five members freeze at first assembly before `attempt_open`. I found no new m-3 consumer defect in those bytes.

**Required correction:** update R1 to pair-approved r12 and record m-3's independent consumer check/binding against exact `04422965...`. Preserve the separate m-10 rebase/review and R3 hold; this approval does not approve m-10, settle the section-D join, or close R2. Add both the `023000` approval and `033000` R1-discharge return to the m-9 watermark. Because these producer relays crossed after handoff, do not record them as another failure of the pre-handoff sweep.

### M3-L2-R10-F2 - BLOCKER / STATUS-LOCALITY CLAIM STILL FALSE - the vocabulary check omits live mutable phrases

Section 0b promises that no mutable current producer state, hash or semantic status, exists outside section 0a, and the incoming relay strengthens the claim to "sections 4/5 now carry no mutable status at all." The operative bytes still say:

- Section 3 heading and line 130: `SECTION PENDING m-10's carriage row`; the artifact "exists but is not bindable" and "the m-9 basis has moved again."
- Section 3.2a: `PARTIAL/PENDING m-8`, two `PENDING - not executable` rows, and the current claim that m-8 has not published the discriminator.
- Section 3.3: three m-10 digest cells are `PENDING`, and line 197 says the expected states "remain deliberately unstated" until producer approval.
- Section 4: the exact carrier is `PENDING` matching producer bytes; item 2 says `I bind none of them`.
- Section 5: m-9 recipe-binding is `PARTIAL` and `NOT discharged`; m-10 is `NOT CONFIRMED - PENDING` and "an artifact now exists"; the m-2 row says no m-9 revision is bound.

Some are m-3's derived consumer state rather than a producer's own verdict, but they are still mutable current status and they change when section 0a changes. That is the multi-locus drift the rule says it removes. The verification searched `non-bindable` with a hyphen, but the live phrase is `not bindable`; it did not search plain `PENDING`, `NOT CONFIRMED`, `I bind none`, `PARTIAL`, or `no m-9 revision is bound`. The 205/233 line-count assertion proves the intended region was inspected; it does not prove the status vocabulary was complete.

**Required correction:** either make every operative dependency statement conditional and timeless, with its current evaluation solely in section 0a, or narrow section 0b's guarantee to the exact state classes actually centralized and stop claiming all mutable semantic status is absent. Under the current stronger rule, sections 3/4/5 must cite section 0a rather than restating `PENDING`/binding state. Make the verification derive or enumerate the forbidden vocabulary including whitespace/hyphen variants; a selected list that omits the live words cannot establish the claim.

### M3-L2-R10-F3 - BLOCKER / SELF-DESCRIPTION CONTRADICTION - the title still says r9

Line 1 ends in `(r9)`. Line 7, the incoming relay, and the exact target hash identify r10. `SELFDESC-END` is an honest category boundary - document revision identity is not producer state - but licensing self-description does not license stale or contradictory self-description.

**Required correction:** update the title to r10. Keep the self-description delimiter, but make the excluded header internally exact before using it as a verification exception.

## Pressure-Point Dispositions

1. **Is section 0b executable by another reviewer?** Yes at the structural level. I reproduced the dispatch ordering method and the 205/233 operative-line scope. Its mutable-vocabulary assertion is incomplete and therefore does not yet prove the claimed invariant.
2. **Does anything in sections 4/5 still carry status that can stale?** Yes: `PENDING`, `I bind none`, `PARTIAL`, `NOT discharged`, `NOT CONFIRMED`, and `no m-9 revision is bound` are live state outside section 0a.
3. **Is `SELFDESC-END` honest?** The boundary category is honest; the enclosed title is not. It says r9 for r10 bytes. I found no current producer hash hidden above the delimiter beyond the explicitly stable governing basis, but the exception still requires correct document identity.

## Preserved Work

- Keep the enumerated cross-lane sweep, embedded-timestamp ordering, per-dispatch watermark, artifact-plus-verdict check, and immediate pre-handoff rerun.
- Keep the anchored delimiters and operative-line-count assertion.
- Keep the VP-mandated B1 authority wording and m-10's assembly-before-`attempt_open` condition.
- Keep every accepted cut matrix, verdict machine, classifier hold, row-state rule, consistency/derivability rule, and authority boundary.
- Treat m-9 `023000` as crossed mail that advances R1, not as evidence against the pre-handoff procedure.

## Re-review Gate

Return fresh bytes with R1 bound to pair-approved, independently consumer-checked m-9 r12 `04422965...`; R2 still unanswered; R3 still held pending m-10's separately pair-approved rebase. Remove or honestly narrow every mutable-status claim outside section 0a, extend the check to the words actually present, and correct the title to r10/r11 as appropriate. Bind no m-10 byte and close no downstream gate.

## Verification

- Reviewed lane-2 r10 at exact SHA-256 `2653cab4f08115b9b617796559136a145d7d4ba9ee4062ebddd5f564ee9bb542`; incoming DESIGN relay at exact SHA-256 `151a4ab368844080145f12effb5eec65c6b0fd5ebbb1f7522be3543a0e887a57`.
- Incoming DESIGN exact-file lint: OK.
- Reproduced the delimiter-scoped operative region at 205 of 233 lines.
- Reproduced m-9 r12 `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35`; reviewed its sections 6/10, exact `023000` approval relay, and exact `033000` R1-discharge return.
- No later relevant m-8, m-10, or m-3 dependency relay existed at final verification time; the same-timestamp m-9 return is folded above.
- No domain design, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`; clean, no short-status entries
RELAY_LINT: OK - exact-file relay-lint.py verification after final write
Next requested action: m-3.planner folds crossed m-9 r12 approval and this consumer check, centralizes or narrows every remaining mutable-status statement, corrects the title, and returns fresh exact bytes; R2/R3/downstream remain held
