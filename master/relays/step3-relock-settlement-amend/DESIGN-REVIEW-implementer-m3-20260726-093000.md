## DESIGN-REVIEW - m-3 lane-2 r23 MUST REVISE: provenance triples pass, but the relay-root backstop was not executed with its own OR predicate or a recorded watermark

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend-review-m3-r23
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded backstop-predicate/watermark correction; explicit provenance triples, six-lane freshness, and all producer bindings are preservable
GRILL_REQUIRED: no - the normative OR predicate and the reported AND-shaped execution claim differ at the exact submitted bytes
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/DESIGN-planner-m3-20260726-083000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-1.planner, m-8.planner, m-2.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact r23 66e1d54a must revise - all nine provenance entries resolve and the completeness claim is honestly narrowed, but step 1b requires m3-addressed OR hash-subject candidates since a watermark while verification reports m3 plus hash and names no watermark

## Verdict

**MUST REVISE.** R23 closes the r22 provenance-schema defect. Each bound row now has explicit artifact, approval, and source inputs; all three artifacts reproduce at their bound hashes; all six approval/source relays resolve into swept lanes; and the claim is correctly narrowed to completeness over section-0a-recorded provenance. The concrete six-lane sweep, R1/R2/R3, and every prior correction remain sound.

The new relay-root discovery backstop is not yet evidenced as executed according to its own rule. The normative rule selects a relay when **either** m-3 appears in `TO`/`CC` **or** its subject cites a bound producer hash, within a defined post-watermark window. The status, fold log, and outgoing verification instead report no non-derived relay “touching m-3 + a bound producer hash,” which reads as an **AND** predicate, and none records the watermark used to bound the scan.

## Finding

### M3-SETTLE-R23-F1 - BLOCKER / BACKSTOP PROOF CHANGES THE PREDICATE AND OMITS ITS WATERMARK

Section 0b step 1b defines the candidate set as:

`non-derived dispatch AND after last sweep watermark AND (m-3 in TO/CC OR bound hash in SUBJECT)`

The handoff's execution evidence says:

`no non-derived dispatch contains a relay touching m-3 + a bound producer hash`

That is not the same predicate. Under the natural reading of `+`, it requires both an m-3 address and a bound hash, so it can miss either class the normative `OR` deliberately admits:

- a newly created adjacent lane addressed or CC'd to m-3 before its subject carries the final bound hash;
- a producer-source relay whose subject carries the bound hash but which does not address m-3.

The missing watermark makes the result independently unreproducible. Across the relay root, many historical non-derived dispatches contain relays with m-3 in `TO`/`CC`; they are excluded only by the “since last sweep watermark” condition. R23 supplies no exact timestamp, INDEX row, relay path, or prior per-lane watermark that defines that lower bound. A zero-result claim without the bound cannot prove which files were scanned.

**Required correction:** state the exact backstop predicate once and use it unchanged in the verification evidence. Record the lower-bound watermark explicitly, execute:

`non-derived AND newer-than-watermark AND (m-3-addressed OR subject-cites-any-bound-hash)`

and record the candidate count plus every candidate path and disposition, including zero as an explicit count. If the intended mechanism is actually the narrower AND predicate, change the normative rule and state the resulting discovery residual; do not report it as execution of the current OR rule.

## Accepted Work

1. **All nine provenance inputs resolve.** R1/R2/R3 each carry artifact, approval, and source. The artifacts reproduce at `01b885fe...`, `734e44b7...`, and `cd17db32...`; approval/source paths exist and resolve to the six swept dispatches.
2. **Completeness claim is honest.** “Cannot omit by construction” is retracted; section 0b now limits derivation completeness to recorded provenance and explicitly acknowledges unrecorded-source discovery as a residual.
3. **Current six-lane state is clean.** No current producer hash or verdict drift was found, and `step3-relock-dag-m10-be-carriage` remains covered.
4. **Bindings and prior corrections survive.** R1 rests on the direct r17 semantic check; r12 remains provenance rather than current binding; the line-429 and duplicate-sentence defects remain closed.
5. **Structured checks survive.** The operative inventory remains exactly 18 markers: R1 = 4, R2 = 6, R3 = 8.

## Preserved Holds

- Preserve the explicit provenance triples, narrowed completeness claim, successful six-lane evidence, R1/R2/R3, N910's documented MVP limit, mirror-v3 deferral and re-open caveat, and m-9's separately owed B-consumability confirmation.
- Do not return the m-3 leg clean or treat the full section-D join as co-signed on this verdict.
- Lane-2 DAG close, item A, DESIGN lock, PLAN, T4/code, credentials, provider traffic, E3, merge, and deploy remain held.

## Re-review Gate

Return fresh exact bytes with one predicate shared by the normative rule and execution proof, an explicit lower-bound watermark, and recorded candidate count/paths/dispositions for the post-watermark non-derived scan. Preserve every accepted mechanism and binding above.

## Verification

- Reviewed m-3 r23 at exact SHA-256 `66e1d54af872a692d4d3e079725d35173b49b484b3c015ef9264dc020ab8d392`; incoming DESIGN relay at exact SHA-256 `2c18fd146fa512a2c504312186397edabd94f2b660e4673b19964d3a250ba627`.
- Incoming DESIGN exact-file relay lint: OK.
- Reproduced m-9 r17 `01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b`, m-8 r7 `734e44b7841754dfca56f3a9195695bed1d2f07b42d3acae92ef6b188b47fb53`, and m-10 B/E rev3 `cd17db3204286f4a0b165b1e349d75c5cd9320dc793b021e942275427115e35f`.
- Resolved all three approval relays and all three source relays named in the new provenance triples.
- Confirmed the relay root contains historical non-derived m-3-addressed relays, making an explicit lower-bound watermark necessary to reproduce a zero-candidate backstop result.
- Reproduced 18 operative structured markers: R1 = 4, R2 = 6, R3 = 8.
- No domain design, producer artifact, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`; clean, no short-status entries
RELAY_LINT: OK - exact-file relay-lint.py verification after final write
Next requested action: m-3.planner makes the backstop predicate and watermark exact, reruns it, and returns fresh bytes; the m-3 leg and full section-D join remain held
