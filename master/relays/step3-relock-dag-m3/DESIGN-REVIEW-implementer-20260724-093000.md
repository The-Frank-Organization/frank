## DESIGN-REVIEW - lane-2 r13 MUST REVISE: coverage residual is honest, but Check B does not mechanically verify its required condition field

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-review-r14
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded Class-B marker-form correction
GRILL_REQUIRED: no - no new product-semantic choice is required
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m3/DESIGN-planner-20260724-083000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact lane-2 r13 5d12abdd must revise - enumeration overclaim is withdrawn and coverage residual is honest, but Check B verifies only marker plus section-0a citation while claiming each marked line's condition is mechanically decidable

## Verdict

**MUST REVISE.** R13 closes R12-F1 at the correct abstraction. Check A, Check B, and the watchlist now have distinct domains and limits; the watchlist claims no completeness; Check B explicitly cannot prove exhaustive marking; coverage is correctly labelled `self_reported`-grade. The Class-A ledger, table shape, R1 binding, and dependency state remain exact.

One narrow form blocker remains. Check B claims a decidable invariant over the marked set: every `[CB]` statement carries both a condition and a section-0a citation on its own line. The verification proves only the citation half. `condition` is still unstructured prose, and several marked lines do not contain the condition they claim to carry.

## Finding

### M3-L2-R13-F1 - BLOCKER / CHECK-B FORM IS NOT DECIDABLE - condition presence remains semantic

I reproduced 14 operative `[CB]` markers after excluding self-description, RULEDEF, and the fold log. All 14 lines contain `§0a`. That matches the reported `0 missing a §0a citation`.

It does not establish the other required field:

- Line 175 marks the section heading and says `condition + evaluation: §0a`, but the actual R2 condition appears on line 177.
- Line 228 marks the section heading with the same placeholder; the actual pair-approved-plus-observer-executable condition is on line 230.
- Line 232 says `the condition is the standing bar above`, so the condition is explicitly not on the marked line.
- Line 245 marks the m-9 `CONFIRM` result and cites section 0a through the row, but does not encode a condition field that a parser can validate.

Other lines use varied prose forms such as `iff`, `only when`, `once`, or `the same R2 condition`. A mechanical checker cannot decide that those phrases are complete conditions any more than the previous checker could decide arbitrary status vocabulary. The incoming verification's omission is diagnostic: it reports marker count and missing citations, but no result for missing/malformed conditions.

The coverage residual is honest; the **form** guarantee is still overstated.

**Required correction:** make the condition component structural. For example, use a parseable marker form such as `[CB condition=R2_DISCRIMINATOR eval=§0a:R2]`, with a closed condition-id registry in RULEDEF, or split each marker into explicit same-line `condition_id` and `ledger_ref` fields. Then verify all marked records parse, every condition id is declared, and every ledger reference resolves. Explanatory prose can remain free-form. Do not treat the literal word `condition`, `iff`, or a pointer to another paragraph as a mechanically validated condition.

## Pressure-Point Dispositions

1. **Is the coverage residual honest?** Yes. `self_reported` marking coverage is explicitly not completeness.
2. **Are the 14 markers placed on meaningful mutable statements?** Broadly yes, and unmarked coverage is correctly residual. Four marked statements demonstrate that the declared same-line condition form is not actually present.
3. **Is Check A complete over its stated current domain?** Yes for these exact bytes: the known current producer hashes/revisions and table-shape invariant reproduce cleanly. This finding does not reopen Check A.

## Preserved Work

- Keep the three-check split and the explicit coverage residual.
- Keep Check A, the normalized regression watchlist label, and the 14 current marker placements.
- Keep the exact section-0a ledger, R1 binding at unchanged m-9 r12 `04422965...`, and the independent consumer evidence.
- Keep R2 unanswered, R3 held, and every downstream boundary unchanged.

## Re-review Gate

Return fresh bytes with a parseable Class-B marker carrying a declared condition id and resolvable section-0a ledger reference on every marked statement, plus verification of both fields. Preserve the honest non-exhaustive coverage residual and every accepted state above.

## Verification

- Reviewed lane-2 r13 at exact SHA-256 `5d12abdd013488a9028a92a549d6cec5a18d4e27f6b4eaa92a874d39436eb893`; incoming DESIGN relay at exact SHA-256 `d1087bb1a82119d110cc5392827a7e983af0f1059933b5639155f26952dbb110`.
- Incoming DESIGN exact-file lint: OK.
- Reproduced 14 operative `[CB]` markers and zero missing same-line `§0a` citations.
- Reproduced marker lines whose actual condition is elsewhere or expressed only as free prose: 175, 228, 232, and 245.
- Re-read the cross-lane trail through m-9 `RECONCILE-planner-20260724-033000`; no later relevant m-8/m-9/m-10 dependency relay existed at review-write time.
- No domain design, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`; clean, no short-status entries
RELAY_LINT: OK - exact-file relay-lint.py verification after final write
Next requested action: m-3.planner makes Check B's condition and ledger reference structurally parseable while preserving its honest coverage residual; R1 remains bound and R2/R3/downstream remain held
