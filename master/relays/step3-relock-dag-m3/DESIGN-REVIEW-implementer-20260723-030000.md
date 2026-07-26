## DESIGN-REVIEW - lane-2 r1 MUST REVISE: the cut table drops carrier facts, partial reads escape predicate 4, the sink consumes an undesigned m-10 row and is ambiguous on divergence, and the logical binding is not yet settled

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-review-r2
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded pair corrections plus Master sequencing of unsettled producer folds
GRILL_REQUIRED: no - no new product-semantic choice is required
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m3/DESIGN-planner-20260723-002000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact lane-2 r1 87172ca11a28279973d0b915e840f59877d1152033265126819274578c3bd74a must revise - carrier matrix incomplete, predicate-4 partial reads unclassified, m-10 sink input still parked, sink derivability ambiguous, and logical recipe binding remains only partial

## Verdict

**MUST REVISE.** The ratified v2 dispatch, phase-not-presence rule, row-2a/2b split, P2a/P2b separation, D2 gating split, predicates 1/2/3/5, F5 authority ceiling, and v3 exclusion are directionally correct. Predicate 2's non-denied result should remain `unknown`: it is recorded-only, and an out-of-scope attempt does not prove a deny property.

Four blocker groups remain. Two are local mechanical gaps. Two expose sequencing mistakes in the parent premise: exact m-10 rev6 still parks the B/E rows the sink claims to consume, and exact m-9 r5 still parks the already-pair-approved m-2 component. m-3 must not approve those producer inputs into existence.

## Findings

### M3-L2-R1-F1 - BLOCKER - the cut matrix is not the dispatched carrier matrix, and rows 9/10 do not instantiate a total freeze predicate

The parent requires each cut to name the exact carrier field as well as E0/E3 presence. Section 1 has no DATA-P or CTRL-C carrier columns. That omission hides a factual error in row 3: it says no m-8 carrier exists, while m-8 r5 section 1 defines a `m8.dataP_reply.v2` epoch reply with no digest members; only the CTRL-C `m8.attempt_result.v2` is absent. No E0 event may still be the right m-3 outcome, but the stated reason is false and the exact source route is missing.

Rows 9/10 retain m-8's `freeze completed? = maybe` but declare E3 `frozen_core_digest` forbidden if a record exists. That is not an instantiation of the ratified `present iff FREEZE-REACHED(cut)` rule: a `maybe` input is not the required total boolean, and field-forbidden is different from no E3 record. The draft simultaneously says those rows are no-record cases and permits the hypothetical record whose field it marks forbidden.

**Required correction:** reproduce the exact DATA-P and CTRL-C carrier/disposition for all ten rows, including message-absent versus carrier-present-with-field-absent. For rows 3/9/10, distinguish `NO E0/E3 RECORD` from a schema-valid record with an absent field. Do not convert m-8's `maybe` freeze into unconditional E3-forbidden. Either split the no-carrier cuts using a named attempt-bound fact that makes `FREEZE-REACHED` boolean, or route the unresolved no-record/no-field consequence to Master as a D4 contract integration issue. Preserve that P2a is never an m-3 input.

### M3-L2-R1-F2 - BLOCKER - predicate 4's three-member domain omits a completed but partial or invalid read response

The `committed | not_found | unavailable` direction is sound, and successful authoritative `not_found` should fail while machinery unavailability remains `unknown`. The discriminator is not yet total. `unavailable` is defined as a read that did not complete, but a transport can complete while yielding a truncated, malformed, unknown-kind, or otherwise non-validatable response. Such bytes are neither a committed record nor an authoritative `not_found`, and the current first-match table has no branch for them.

**Required correction:** define successful completion at the governed-read protocol grain, not transport EOF. `committed` must require one fully decoded and validated committed record; `not_found` must require the exact successful authoritative absence result; every partial, truncated, malformed, unknown, or non-validatable response must map mechanically to `unavailable`/`unknown` or to a new closed result with an explicit verdict. Then restate the five branches as exhaustive over that closed result domain.

### M3-L2-R1-F3 - BLOCKER - the sink consumes a producer row that does not exist at rev6, and its own consistency/derivability rules are not total over divergence

Section 5 confirms an m-10 rev6 `m10_row_digest` input without an exact hash. Exact m-10 rev6 is `29a123fe8f153127ccfe03559850c8e5ce4c488671402cebf8325e3cafbcd7ae`; its section 10 and the `194500` SITREP explicitly say the B-row and E-row are **DELIBERATELY NOT DESIGNED / still PARKED**. The parent dispatch's phrase "m-10 rev6 carriage" is therefore a stale integration premise. Prior m-3 legs cannot confirm a field absent from the current producer delta.

The sink algorithm also has two concrete fall-through ambiguities:

1. `not_applicable iff expected_presence=no_carrier` classifies even an impossible unexpected digest at one of the three carriage points as not-applicable rather than inconsistent. That masks exactly the invented-carrier/mixed-presence defect the sink should record.
2. `derived_equal|derived_unequal` compares the observer value with "the carriage value." When m-8=A, m-9=B, and m-10=A, there is no single carriage value; the same observer=A is equal and unequal depending on an unnamed choice.

The record also names `cut_row` as settled classification without naming the attempt-bound source/actor that establishes it. That makes `expected_presence` self-declared rather than independently reconstructible.

**Required correction:** Master routes settled m-8/m-9 producer bytes to m-10; m-10 authors and pair-approves the bounded B/E row fold; m-3 then binds the exact returned hash. Make `m3.b_sink.v1` a closed record with exact types/enums/required fields and name the attempt-bound cut-classification source. Define consistency over every expected-presence x three-carriage-value combination, including unexpected presence in a no-carrier row. Define derivability against a named canonical comparand or as a per-carrier comparison vector; do not use singular "the carriage value" when consistency is false.

### M3-L2-R1-F4 - BLOCKER / MASTER SEQUENCING - the partial logical binding is honest but does not discharge the requested binding

Section 4 correctly refuses to overclaim exact m-9 r5. That producer revision `c0ff74f5589e0dcbd066fe8a7e04ee7fafd53a9a20befa6453084110feccf70b` defines the five-member outer recipe but explicitly parks folding m-2's component in sections 6/9. A full m-3 recipe-binding confirmation cannot be made against those bytes.

The stated wait condition is stale, however: m-2's component has already landed and is pair-approved at exact SHA-256 `c3a8cd61bcbe33a3ede847b29d7416563000c66f949fa2950ea0af305edbbd2c` (`step3-relock-dag-m2/DESIGN-REVIEW-implementer-20260722-103000` and SITREP `110000`). What remains is not "after m-2 lands"; it is Master routing those settled bytes to m-9, m-9 folding them into a fresh pair-approved producer revision, and then m-3 binding that exact revision. Approving this artifact as the last m-3 lane piece would turn an explicitly partial confirmation into the complete deliverable the parent requested.

**Required correction:** keep the partial status and no-gating claim honest, but do not claim this lane deliverable complete. Escalate the exact state to Master: m-2 is settled; m-9 consumption is not. After a fresh m-9 fold and pair approval, bind the complete recipe and observer re-derivation against that exact hash. m-3 must not fold m-2-owned composition into m-9's producer contract itself.

## Preserved Work

- Ratified amendment `9e874df84015261d77e9c353528e821fd8491489388c010fa621fe630432b351`, bound v2 contract `6e2abe40be7a6031163348d47e6b9c13990c5599eb877d30ef196b59efbf3e1f`, and frozen r4 remain untouched.
- Rows 1/2a/2b correctly prove that `phase=failed` cannot encode freeze state; row 2b correctly carries the step-2 authorized digest.
- P2b keys only on m-3's resolvable complete capture; m-8's P2a counter remains context, never an E3 input.
- Predicates 1/3/4 are gating; predicates 2/5 are recorded-only. Predicate 2 non-denied remains honest `unknown`; predicate 5 retains the observed-window ceiling.
- The sink remains a record at F5 authority, never producer enforcement or the Master+VP composite join.
- `model_surface_digest` and the E join remain outside v2 and outside this fold, deferred to a separately governed v3 cycle.

## Re-review Gate

Return fresh bytes closing F1/F2 and the sink's local schema/rule defects. For F3/F4, cite Master's corrected routing and only bind producer inputs after their pair-approved bytes exist; if Master elects to split the work, mark the sink and logical binding explicitly pending and do not file a lane-complete SITREP. No integrated re-lock, PLAN, T4/code, credential, provider, release-binding, live E3, merge, deploy, or H-12 external-use gate advances on r1.

## Verification

- Reviewed lane-2 r1 at exact SHA-256 `87172ca11a28279973d0b915e840f59877d1152033265126819274578c3bd74a`; incoming DESIGN relay at exact SHA-256 `dcc4564e877c40272b66b9d6b2824b7e2109e0920f8078fbd206a3d02728bc82`.
- Reproduced ratified amendment `9e874df84015261d77e9c353528e821fd8491489388c010fa621fe630432b351`, bound contract `6e2abe40be7a6031163348d47e6b9c13990c5599eb877d30ef196b59efbf3e1f`, frozen r4 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`, m-8 r5 `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21`, m-9 r5 `c0ff74f5589e0dcbd066fe8a7e04ee7fafd53a9a20befa6453084110feccf70b`, m-10 rev6 `29a123fe8f153127ccfe03559850c8e5ce4c488671402cebf8325e3cafbcd7ae`, and m-2 rev2 `c3a8cd61bcbe33a3ede847b29d7416563000c66f949fa2950ea0af305edbbd2c`.
- Incoming DESIGN exact-file lint: OK.
- No domain design, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: none - `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6`; workspace root is not a git repository
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-relock-dag-m3/DESIGN-REVIEW-implementer-20260723-030000.md`
Next requested action: m-3.planner folds the two local mechanics findings, escalates the stale m-10/m-9 producer premises to Master for corrected sequencing, and returns only once the exact producer dependencies required by the claimed completion state exist
