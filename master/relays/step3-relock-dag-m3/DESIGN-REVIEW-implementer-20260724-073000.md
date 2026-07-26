## DESIGN-REVIEW - lane-2 r12 MUST REVISE: ledger integrity passes, but normalized watchlist matching is still not vocabulary enumeration

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-review-r13
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded verification-contract correction
GRILL_REQUIRED: no - no new product-semantic choice is required
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m3/DESIGN-planner-20260724-063000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact lane-2 r12 00cd9951 must revise - R1 row/table shape and Class-A centralization pass, but the claimed status-vocabulary enumeration remains a selected normalized watchlist and omits live Class-B status words

## Verdict

**MUST REVISE.** R12 closes both exact R11 blockers. Section 0a now has one structurally valid four-cell m-9 row carrying only the bound state; every ledger line matches the header's five pipe delimiters; line 17 correctly distinguishes recording from authorization. The Class-A-only guarantee is stated at the top of section 0b, section 5 no longer leaks the current m-9 result into the m-10 row, and the case/separator-normalized counts for the listed phrases reproduce exactly. R1's binding at unchanged m-9 r12 `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35` remains valid.

One proof-mechanism blocker remains: normalization makes a selected vocabulary robust to spelling variants; it does not turn that selected vocabulary into an enumeration of all status present in prose. R12 still claims the latter.

## Finding

### M3-L2-R12-F1 - BLOCKER / UNFALSIFIABLE ENUMERATION CLAIM - the check still starts from a chosen status list

I reproduced the reported normalized output over the 206-line scope:

- `not confirmed` x1
- `partial` x4
- `pending` x2
- all other phrases in the incoming watchlist x0

That proves case/separator normalization works for those phrases. It does not prove that the check "enumerates the status vocabulary actually present." The operative region contains additional live status language not in the reported vocabulary, including:

- Section 4 heading: the recipe `BINDS`; the carriage row `does not`.
- Section 4: `THE CONDITION IS NOW MET`, `settled`, and `observer-executable`.
- Section 5: `CONFIRM` and `Not confirmed` consumer outcomes.
- Section 3: `incomplete`, `NOT total`, `not executable`, and `unexecutable` derived state.
- Section 6/status boundaries: `not complete`, `closed`, and `held` concepts expressed outside the selected phrases.

These are not necessarily violations; most are legitimate Class B. Their existence proves the current procedure is still watchlist matching, not vocabulary discovery. No mechanical operation described in section 0b can infer which arbitrary prose words are semantic status. Case-folding and separator-folding solve the r10/r11 spelling defects but not this abstraction gap.

There is also a direct contract ambiguity: line 39 says Class B is "explicitly OUTSIDE the mechanical check," while lines 43-48 say the check enumerates, reports, and classifies Class-B hits. The incoming verification follows the latter. The text must distinguish the Class-A violation check from the Class-B inventory check rather than call both "the mechanical check."

**Required correction:** choose one mechanically honest design:

1. **Structured inventory:** mark every mutable Class-B statement with one canonical machine token/record, then enumerate those markers and verify each carries a condition plus section-0a citation. Free prose may explain a marked statement but is not itself claimed exhaustively scanned.
2. **Narrow watchlist:** retain the normalized phrase list, call it a regression watchlist rather than enumeration, and withdraw any claim that it proves all semantic status has been found.

Resolve line 39 so Class B is outside the Class-A violation rule but inside the structured inventory, if that is the selected design. Do not add another word to the existing list and call the abstraction closed.

## Pressure-Point Dispositions

1. **Does the m-9 row carry exactly one state?** Yes. R1 is discharged/bound only; the obsolete proposed/not-bindable residue is gone.
2. **Is the Class-A-only guarantee internally consistent?** Yes, apart from the ambiguous use of "the mechanical check" for two different checks. I found no current producer state outside section 0a.
3. **Does classified enumeration establish its stated claim?** No. It establishes normalized matching for a selected phrase set. The live unlisted Class-B status words demonstrate the difference.

## Preserved Work

- Keep the valid table-shape assertion and all four-cell ledger rows.
- Keep R1 bound at exact unchanged m-9 r12 `04422965...` and preserve the independent consumer check.
- Keep the Class-A-only centralization rule, conditional Class-B semantics, normalized matching, and classified reporting as components of the corrected proof.
- Keep R2 unanswered, R3 held, and every downstream boundary unchanged.

## Re-review Gate

Return fresh bytes that either structurally mark and enumerate all Class-B status statements or honestly narrow the normalized scan to a regression watchlist. Resolve the Class-B check wording and preserve the exact current ledger, R1 binding, and all R2/R3 holds.

## Verification

- Reviewed lane-2 r12 at exact SHA-256 `00cd995186dbbcd701cf687c75ffcbe36dc388b159d6290b2c62e4e964fb71c3`; incoming DESIGN relay at exact SHA-256 `e93940d7fb3b3c13433aa4d3b95fefa844c1eac94d08fac345ebfdb78f7e805e`.
- Incoming DESIGN exact-file lint: OK.
- Reproduced section-0a table shape: header, delimiter, and three producer rows each carry five pipe delimiters.
- Reproduced the normalized status-watchlist counts: `not confirmed` x1, `partial` x4, `pending` x2.
- Re-read the cross-lane trail through m-9 `RECONCILE-planner-20260724-033000`; no later relevant m-8/m-9/m-10 dependency relay existed at review-write time.
- No domain design, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`; clean, no short-status entries
RELAY_LINT: OK - exact-file relay-lint.py verification after final write
Next requested action: m-3.planner replaces the semantic-enumeration overclaim with a structured Class-B inventory or an honestly narrow regression watchlist; R1 remains bound and R2/R3/downstream remain held
