## DESIGN-REVIEW - lane-2 r11 MUST REVISE: R1's recipe binding passes, but its ledger row contains both bound and obsolete not-bindable states and the Class-B scan misses live title-case status

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-review-r12
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded ledger-integrity and verification-honesty corrections
GRILL_REQUIRED: no - no new product-semantic choice is required
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m3/DESIGN-planner-20260724-043000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact lane-2 r11 3356f421 must revise - the m-9 ledger row is a seven-pipe malformed row carrying both R1 bound and obsolete proposed/not-bindable states, while the Class-A/Class-B prose and case-sensitive status proof contradict the live operative bytes

## Verdict

**MUST REVISE.** The important semantic work passes: m-9 r12 remains pair-approved at exact `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35`, and the m-3 consumer check establishes observer-executable composition for all five `logical_surface_digest` members. R1 may bind on those bytes without moving R2, R3, m-10, the section-D join, or any downstream gate.

The exact r11 document cannot be approved. Its canonical m-9 ledger row contains the new bound state and the obsolete proposed/in-pair-review/not-bindable state simultaneously as two extra table cells. The Class-A/Class-B rewrite also retains the old stronger global guarantee, and its claimed enumerated status scan misses the live title-case `Not confirmed` token in section 5 while reporting zero status tokens.

## Findings

### M3-L2-R11-F1 - BLOCKER / CANONICAL LEDGER CORRUPTION - the m-9 row carries mutually exclusive states

Section 0a declares a four-column table. The header and ordinary rows each contain five pipe delimiters. The m-9 row at line 20 contains **seven**:

1. Its first four cells correctly say m-9 r12 is pair-approved, independently consumer-checked, and `R1 = DISCHARGED AND BOUND`.
2. Two trailing cells retain r10-era text: the VP classification followed by `SUCCESSOR PROPOSED AND UNDER PAIR REVIEW. NOT BINDABLE`.

The row is therefore structurally malformed and semantically contradictory inside the single source of truth. A Markdown renderer cannot restore the intended ownership of the extra cells, and a reader can select either current state. Line 17 also says `Nothing here is bindable` immediately before the ledger records a binding; if the intended rule is that the ledger does not self-authorize binding, it must say that rather than deny the state it records.

**Required correction:** rebuild the m-9 entry as exactly one four-cell row: producer, exact bound bytes, pair-approved plus independent-consumer-check evidence, and the narrow R1 consequence. Remove the entire obsolete proposed/in-review/not-bindable tail. Reword line 17 to state that ledger inclusion alone grants no binding and that only reproduced pair-approved bytes bind. Add a mechanical table-shape assertion: every producer row has the same pipe count/column count as the header.

### M3-L2-R11-F2 - BLOCKER / CLASS SPLIT NOT HONESTLY VERIFIED - the prose and scan disagree with live bytes

The Class-A/Class-B distinction is a valid correction in principle: producer state can be centralized while m-3's own completeness remains conditional. Its implementation is internally inconsistent:

- Section 0b line 26 still promises `no duplicated MUTABLE current-state - hash or semantic status - anywhere else`, while line 37 explicitly permits conditional Class-B status outside section 0a. Both cannot be the operative guarantee.
- Section 5 line 234 contains live `Not confirmed` Class-B status. The incoming verification reports `NOT CONFIRMED` has zero occurrences and claims zero live status tokens outside section 0a. That is another selected spelling miss, now by case rather than hyphenation.
- The same line says the next pair-approved m-9 revision `now` exists. That is current producer state wearing a Class-B label, contrary to Class A. Its `rev6` reference also violates line 47's rule that producer revisions are always qualified as `m-<n> r<N>`; if it is retained as historical context, it must be explicitly licensed or moved to the fold log.
- Section 4's ordered list is `1, 2, 2`, a smaller structural residue in the same rewritten block.

The 222/253 and 206-line scope counts may be correct, but scope does not establish a case-sensitive vocabulary claim. A real enumeration cannot report zero while an allowed Class-B token remains; it must report the hit and classify it as licensed.

**Required correction:** make line 26 state the Class-A-only guarantee and the conditional Class-B rule. Normalize case and hyphen/whitespace before status comparison, enumerate the actual hits, and require every surviving hit to carry both a condition and a section-0a citation. Report licensed Class-B hits rather than zero. Rewrite the m-10 row as a timeless condition only; remove the `now exists` producer-state restatement and move/qualify the historical revision reference. Renumber section 4's list.

## Pressure-Point Dispositions

1. **Is the Class-A/Class-B split honest?** The distinction is honest; the exact prose and proof are not yet. The old all-status-absent guarantee remains, and section 5 mixes producer state into a Class-B row.
2. **Does the five-member recipe claim rest on checking?** Yes. I rechecked the exact m-9 r12 recipe and preserve R1's binding. The ledger corruption does not invalidate the producer bytes or the independent consumer result; it invalidates this m-3 revision's representation of them.
3. **Is the RULEDEF exclusion legitimate?** Yes in principle: a rule must quote its vocabulary, and keeping RULEDEF inside the Class-A scan while excluding it only from status-hit classification is coherent. The current case-sensitive zero-result is not a valid execution of that design.

## Preserved Work

- Keep R1 bound at exact m-9 r12 `04422965...` and preserve the independent m-3 consumer evidence.
- Keep R2 unanswered and R3 unbound pending separate m-8/m-10 outcomes.
- Keep the Class-A/Class-B conceptual split, conditional consumer-state form, exact cross-lane sweep, and qualified-revision rule after fixing their execution.
- Keep all accepted cut, verdict-machine, classifier, row-state, consistency, derivability, and authority-boundary mechanics.

## Re-review Gate

Return fresh bytes with one structurally valid, non-contradictory m-9 ledger row; an exact Class-A-only guarantee; a case-normalized status enumeration that reports and classifies licensed Class-B hits; no current producer fact hidden in section 5; and corrected list numbering. R1 remains bound only to unchanged `04422965...`; R2/R3 and all downstream gates remain held.

## Verification

- Reviewed lane-2 r11 at exact SHA-256 `3356f4213b31a1353dae8765d20c97fad4ed1f9c35c3ddeb227a357dcd97ebdb`; incoming DESIGN relay at exact SHA-256 `47b0727243acfc8afedd677de07647ba24634e5007bbd377d24fd1cb3ddfc2ae`.
- Incoming DESIGN exact-file lint: OK.
- Reproduced the section-0a pipe counts: header/m-8/m-10 rows = 5 delimiters; m-9 row = 7.
- Reproduced live `Not confirmed` outside section 0a despite the incoming zero-occurrence claim for `NOT CONFIRMED`.
- Re-read the cross-lane trail through m-9 `RECONCILE-planner-20260724-033000`; no later relevant m-8/m-9/m-10 dependency relay existed at review-write time.
- No domain design, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`; clean, no short-status entries
RELAY_LINT: OK - exact-file relay-lint.py verification after final write
Next requested action: m-3.planner repairs the canonical row and Class-A/Class-B proof without moving the valid R1 binding; R2/R3/downstream remain held
