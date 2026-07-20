## DESIGN-REVIEW - s8 config/atom grammar rev6 must revise final lock-impact echo

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-m2-grammar-design-review-r7
PARENT_DISPATCH_ID: s8-m2-grammar-design-r7
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - the two operator legs remain required at the reconciled lock and are not proxied by this review
GRILL_REQUIRED: yes
GRILL_LOCK_ID: s8-m2-grammar-grill
DESIGN_DOC_ID: s8-design-m2-grammar
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m2-grammar/DESIGN-planner-20260711-002200.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-5.planner, m-6.planner, m-7.planner
SUBJECT: rev6 must revise - primary sequence and fixture now pass, but the design-lock impact still encodes the superseded two-owner seam

DESIGN_REVIEW_VERDICT: must-revise

### Finding

1. **BLOCKER - one current design-lock statement escaped the consistency sweep.** The rewritten matrix, step-3/4/4.5 sequence, fixture, §4 table, GRILL resolved decision, rev6 fold-log, and §8 reconciliation are mutually consistent and close both `001600` blockers. But the lock-bearing `Design-lock impact` still says the lock reflects “observe-fill enrichment ordering” and that only the “m-7/m-3” seams must bind (`design:139`). That is the superseded pre-reconciliation contract. The reviewed lock now also contains m-2's step-3 `authority_class`, the m-5/m-6 Option-B producer/profile manifest, the registry-row amendment removing `surface_intent`'s static predicates, and m-7's formal step-4.5 amendment vehicle. Leaving those out of the lock summary creates an incomplete handoff to master even though the detailed sections are correct.

   **Required revision:** replace the `Design-lock impact` bullet with the confirmed four-owner step-3/4/4.5 contract and explicitly name:
   - m-2 `authority_class` at step 3 with the s5 ③ tripwire;
   - m-3's confirmed step-4 manifest and precision notes;
   - m-5/m-6 Option B plus m-2's registry-row amendment removing the static `surface_intent` predicates;
   - m-7's step-4.5 amendment carried by `s8-design-m7-config`.

   Also change fixture `Actual` from “assembled committed record ... at step 4.5” (`design:76`) to “assembled in-courier candidate ... at step 4.5”; the committed record is correctly tested only by the subsequent behavioral oracle. Step 4.5 is pre-commit by the confirmed m-7 placement.

2. **HUMAN DECISIONS remain open at the reconciled lock.** This review does not proxy the operator's activation-authorization ratification or m-3's three grill defaults. Both remain required on record per the master close-out.

### Accepted Rev6 Fold

- The numbered form-validation -> lineage -> step-3 `authority_class` -> step-4 m-3 manifest -> step-4.5 Option-B derivation/completeness sequence is correct.
- The fixture now correctly separates registry, producer/profile, and actual/provenance expectations; the impossible registry=producer equality is gone.
- Matrix, seam, §4, GRILL resolved decision, fold-log, and §8 are consistent with all four owner confirms.
- No owner-confirmed decision requires reopening. The required revision is a current lock-summary/pre-commit terminology sweep only.

### Re-review Bar

Return one minimal revision correcting `Design-lock impact` and the pre-commit fixture noun. No new owner consultation is required. Keep both operator legs explicitly pending until recorded. No code, PLAN, IMPL, registry edit, c1 edit, or silent §3 amendment is authorized.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW of rev6 `s8-design-m2-grammar` against the accepted owner records and `frank@2e1b4f0`; wrote this relay and appended `master/relays/INDEX.md`; no frank edit
FINAL_GIT_STATUS_SHORT: `frank/` main clean and synchronized at `2e1b4f036c3cfb66902655dcbc3c6702eb1c4023`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0; `--relay-root master/relays/s8-design-m2-grammar` lineage lint exit 0 for the live directory and this relay

Next requested action: m-2.planner corrects the final lock-impact echo and pre-commit noun, then returns the unchanged design for final review while the two operator legs proceed in parallel.
