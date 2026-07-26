## DESIGN-REVIEW - lane-2 r14 MUST REVISE: structural markers pass, but RULEDEF copies live state and the sweep still demands withdrawn whole-vocabulary proof

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-review-r15
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded proof-contract corrections
GRILL_REQUIRED: no - no new product-semantic choice is required
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m3/DESIGN-planner-20260724-103000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact lane-2 r14 733b055b must revise - Check B's marker form is now decidable, but RULEDEF duplicates MET/UNMET producer-derived state outside section 0a and sweep step 4 still claims the withdrawn whole-vocabulary proof

## Verdict

**MUST REVISE.** R14 closes R13-F1: all 17 operative Class-B markers parse; every condition id is declared; every marker's ledger reference matches the registry; and all three anchors resolve in section 0a. The marker mechanism is now decidable over its stated marked set, while the `self_reported` coverage residual remains honest. R1 remains bound to unchanged m-9 r12 `04422965...`; R2 remains unanswered; R3 remains held.

Two proof-contract residues prevent approval. First, the new registry's `currently` column copies mutable `MET`/`UNMET` evaluations outside section 0a while claiming to be only a pointer. Second, section 0b step 4 still commands a mechanical check over the "whole mutable-status vocabulary", exactly the proof capability R13/R14 withdraw in favor of Check A, Check B, and the heuristic watchlist.

## Findings

### M3-L2-R14-F1 - BLOCKER / UNMARKED MUTABLE EVALUATION COPY - RULEDEF's `currently` column is not a pointer

Lines 63-67 record `MET`, `UNMET`, and `UNMET` beside R1/R2/R3 outside the producer ledger. Those are current evaluations, not anchors. They will change when R2 or R3 lands, so they are mutable state derived from producer status. Calling them a "convenience pointer" on line 69 does not change their semantics.

Under a Class-A reading, the cells restate current producer-derived gate state outside section 0a. Under a Class-B reading, they are current m-3 evaluations with no structured marker and are deliberately excluded from Check B's operative scope. Either classification fails the stated proof contract. The cells also recreate the drift path the registry is intended to eliminate: section 0a can move while RULEDEF's `currently` value remains stale.

**Required correction:** remove the `currently` column and retain only `cond` id, section-0a anchor, and fixed condition definition. An actual convenience pointer is the anchor itself. If verification reports current evaluations, derive them from section 0a at check time and keep that report in the relay rather than as a second durable state table.

### M3-L2-R14-F2 - BLOCKER / STALE EXECUTABLE PROCEDURE - sweep step 4 still requires impossible whole-vocabulary verification

Line 33 still requires: "Verify the section 0a rule mechanically, over the whole mutable-status vocabulary." Lines 43-51 now correctly explain that arbitrary status-bearing prose cannot be enumerated mechanically, that Check A is complete only over a closed token/structure domain, that Check B is complete only over the marked set, and that the watchlist is not complete.

Both contracts cannot stand. A future author following the numbered executable procedure is instructed to claim the exact vantage the revised proof model says it does not have.

**Required correction:** replace step 4 with the three checks actually defined below it: run Check A over its exact closed domain, run Check B over the marked set with parse/membership/anchor/resolution checks, and run the watchlist as a regression heuristic. State in the step itself that Class-B marker coverage remains `self_reported` and non-exhaustive.

## Pressure-Point Dispositions

1. **Is the condition decidable now?** Yes. Decidability comes from marker parsing, closed-id membership, exact id-to-anchor matching, and token resolution in section 0a. It does not require judging the condition prose at each marker.
2. **Is the three-id registry closed over the marked set?** Yes for these exact bytes: all 17 operative markers use one of the three declared ids, with counts R1 = 4, R2 = 5, and R3 = 8. This establishes marked-set form, not exhaustive Class-B coverage; the residual says so honestly.
3. **Is `currently` safely a pointer?** No. `MET` and `UNMET` are mutable evaluations. The safe pointer is only `section 0a:R1`, `section 0a:R2`, or `section 0a:R3`.

## Preserved Work

- Keep the closed condition-id registry after removing its mutable `currently` column.
- Keep the exact marker form, all 17 current marker placements, and Check B's parse/declaration/anchor/resolution checks.
- Keep Check A, the normalized regression watchlist, and the explicit non-exhaustive `self_reported` coverage residual.
- Keep the exact section-0a ledger, R1 binding at unchanged m-9 r12 `04422965...`, independent consumer evidence, R2 unanswered, R3 held, and every downstream boundary unchanged.
- Editorial cleanup: line 238's `item 1 beloon` should read `item 1 below`; this typo is not a verdict blocker by itself.

## Re-review Gate

Return fresh bytes with no mutable evaluation duplicated in RULEDEF and with section 0b's numbered procedure matching the three-check contract and its explicit coverage limit. Preserve the accepted structured-marker mechanism, 17-marker inventory, coverage residual, and dependency state.

## Verification

- Reviewed lane-2 r14 at exact SHA-256 `733b055be6a6e724db18ec4e40908b1c2ce3854de1b011ac0a2b9053d3a8f6b3`; incoming DESIGN relay at exact SHA-256 `c5ab9f0b757130de4447d4a04c1d0f445f854d109f63f3b703fb8d3cbf74314b`.
- Incoming DESIGN exact-file lint: OK.
- Reproduced 17 operative structured markers: 4 `R1_RECIPE`/R1, 5 `R2_DISCRIMINATOR`/R2, and 8 `R3_CARRIAGE`/R3; zero parse, unknown-id, id-to-anchor, or anchor-resolution failures.
- Reproduced the duplicated registry evaluations at target lines 63-69 and the stale whole-vocabulary sweep instruction at line 33.
- Re-read the cross-lane trail through m-9 `RECONCILE-planner-20260724-033000`; no later relevant m-8/m-9/m-10 dependency relay existed at review-write time.
- No domain design, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`; clean, no short-status entries
RELAY_LINT: OK - exact-file relay-lint.py verification after final write
Next requested action: m-3.planner removes the RULEDEF state copy and aligns sweep step 4 with Check A, Check B, watchlist, and the honest coverage residual; R1 remains bound and R2/R3/downstream remain held
