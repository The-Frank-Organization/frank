## DESIGN-REVIEW - lane-2 r15 APPROVE: registry state copy removed and executable sweep now matches the bounded three-check proof contract

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-review-r16
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - prior proof-contract blockers are closed without a product-semantic decision
GRILL_REQUIRED: no - no new product-semantic choice is required
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/step3-relock-dag-m3/DESIGN-planner-20260724-123000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact lane-2 r15 d004dbc7 approve - RULEDEF now contains only fixed condition definitions and anchors, step 4 executes the three bounded checks with the coverage residual inline, and R1 remains byte-bound while R2/R3 stay open

## Verdict

**APPROVE** the exact r15 design bytes at SHA-256 `d004dbc77e70477175a78a7aa649ad8cdb5bf33224ea4abff38ab3d818e54300` for `DESIGN_DOC_ID: step3-relock-dag-m3`.

R15 closes both R14 findings. RULEDEF is now a three-column registry containing only condition id, section-0a anchor, and fixed condition meaning. It carries no current evaluation cell. Section 0b step 4 now commands Check A over its closed token/structure domain, Check B over the marked set, and the watchlist as a heuristic, while stating inline that Class-B marker coverage is `self_reported` and non-exhaustive.

This approval is bounded to the reviewed design bytes and proof contract. It does not declare the lane complete: R1 remains bound to unchanged m-9 r12 `04422965...`; R2 remains unanswered; R3 remains non-bindable; no m-10 byte or downstream gate is approved.

## Finding Dispositions

### M3-L2-R14-F1 - CLOSED / registry no longer stores a current evaluation

The registry at lines 63-67 has exactly three columns and no `currently`, `MET`, or `UNMET` evaluation cell. Line 69's `MET/UNMET` text is a negated rule explanation, not a live evaluation. The only registry pointer is its section-0a anchor.

Current derived Class-B evaluations still appear where the design must state its own completeness, for example the marked section-4 heading and its marked `THE CONDITION IS MET` statement. That is intentional under the accepted class split: each is structurally marked with a declared condition id and exact ledger anchor. The removed defect was the unmarked RULEDEF copy, not the existence of conditional Class-B consequences outside the ledger.

Reporting R1/R2/R3 evaluations in the exact-byte DESIGN relay is a point-in-time derivation and evidence record, not a second source of truth inside the design. Any later review must refresh that report against the live ledger and trail; this approval does not make an old relay current forever.

### M3-L2-R14-F2 - CLOSED / numbered procedure matches the declared proof limits

Step 4 no longer commands a whole-status-vocabulary proof. It names the same three checks defined immediately below it, gives each its actual domain, and repeats the non-exhaustive `self_reported` Class-B coverage limit. There is no residual executable instruction claiming mechanical completeness over arbitrary prose.

## Pressure-Point Dispositions

1. **Any unsafe current evaluation outside section 0a?** No unmarked operative copy was found. Marked Class-B consequences remain deliberately outside the ledger and satisfy the accepted id-plus-anchor form; self-description and the fold log remain delimiter-licensed history.
2. **Does step 4 match the definitions below it?** Yes: Check A = closed tokens/structure; Check B = marked-set parse, membership, anchor match, and resolution; watchlist = regression heuristic; coverage = non-exhaustive.
3. **Is relay-time derivation a hidden state copy?** No, provided it remains exact-byte, timestamped evidence and is refreshed on each fold. It does not replace or amend the ledger and grants no future-current status.

## Preserved Boundaries

- The structured marker form and all 17 current marker placements remain unchanged.
- Check A, Check B, the watchlist, and the explicit coverage residual retain distinct claims.
- The section-0a ledger and table-shape rule remain authoritative.
- R1 stays byte-bound only at m-9 r12 `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35`.
- R2, R3, m-10 carriage, the section-D join, DESIGN lock, PLAN, T4/code, credentials, provider traffic, E3, merge, and deploy remain outside this approval.

## Verification

- Reviewed lane-2 r15 at exact SHA-256 `d004dbc77e70477175a78a7aa649ad8cdb5bf33224ea4abff38ab3d818e54300`; incoming DESIGN relay at exact SHA-256 `d6483a323c8b841b880ed779d19d59a68867b9ee33b7a5100f097bc97b94212c`.
- Incoming DESIGN exact-file relay lint: OK.
- Reproduced RULEDEF's three-column shape: four pipe delimiters on its header, separator, and three rows; no current evaluation cell exists.
- Reproduced 17 operative markers: 4 `R1_RECIPE`/R1, 5 `R2_DISCRIMINATOR`/R2, and 8 `R3_CARRIAGE`/R3; all 17 parse and all id-to-anchor mappings match.
- Reproduced the section-0a table shape: five pipe delimiters on the header, separator, and all three producer rows.
- Re-hashed the bound m-9 design artifact at `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35`; its approving review remains `DESIGN_REVIEW_VERDICT: approve` over that exact hash.
- Re-read the latest relevant trail: m-8 through `RECONCILE-orchestrator-planner-20260723-104500`, m-9 through `RECONCILE-planner-20260724-033000`, m-10 through `SITREP-planner-20260723-223000`, and m-3 through the incoming DESIGN `20260724-123000`. No later producer relay existed at review-write time.
- No domain design, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`; clean, no short-status entries
RELAY_LINT: OK - exact-file relay-lint.py verification after final write
Next requested action: m-3.planner returns the approved exact r15 hash to master while preserving R2/R3 and every downstream hold; any later design lock must bind this exact DESIGN_DOC_ID and SHA
