## SITREP — external verdicts reconciled (m-7 narrow must-revise, PRE-CONCURRED on fold · m-1 approve-conditional · m-2 approve-conditional): bounded design-r3 fold directed (R1 + R2 + the M1-4 guard); plan r2 carries all F-rows verbatim; fresh reviews per the content-changing-fold rule

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s6-core-design-verdicts
PARENT_DISPATCH_ID: s6-core-design-r2-review-implementer
RUN_ID: s6
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: s6-slice-6-design
GRILL_LOCK_ID: s6-grill-s6-core
FROM: s6.orchestrator-planner
TO: s6-core.planner
CC: s6-core.implementer, s6.orchestrator-reviewer, operator
SUBJECT: fold directive — design r3 = m-7's R1(i–iii) + R2 verbatim + m-1's F-S6-M1-4 guard line; fresh implementer DESIGN-REVIEW on the fold (content-changing — the s3 rule); plan r2 carries F-S6-M1-1..5 + F-S6-M2-1..4 verbatim + the R1/R2 task additions; the gated PLAN re-parents to the NEW approving review; m-1 narrow confirm on R1's shape routed by me in parallel

**The three verdicts (all lint-read at my seat; copies/paths):** m-7 guide `../.relays/s6/s6-guide-m7/DESIGN-REVIEW-planner-20260707-012324.md` (**must-revise, narrow, PRE-CONCURRED on faithful fold — no second guide round**; asks 1/2/3/5 confirmed, several design realizations credited above the spec's own text) · m-1 `s6-fidelity-m1/SITREP-implementer-20260707-012143.md` (**approve-conditional**: F-S6-M1-1..5; no narrow re-review needed for a plan that carries them verbatim and only adds the M1-4 guard) · m-2 `s6-fidelity-m2/SITREP-implementer-20260707-012246.md` (**approve-conditional**: F-S6-M2-1..4; **the [VP-W3] reading CONFIRMED by the domain owner**; the render-absence question RULED: `visible_when` + `gate_referenceable:false` + submit-path rejection — never seat-scope-alone).

**R1 verified real at my seat before this directive** (`intake.Cmd` = `{IntakeID, Seat, Role, IsOperator, Verb, Payload, ContentHash}` — no generation tag; my own read this session): a boot-form submit queued by the dying session before the `seat_mint` pivot passes B-1.2a admission after it and would become the NEW generation's activation record — defeating the re-mint's intent while formally satisfying the order rule. The sharpest catch of the phase; the guide-gate posture earned its keep.

### Directed fold (bounded; no re-architecture; everything below is pre-specified by its owner)

**Design r3 (one pass, docs-only):**
1. **R1(i–iii) per m-7's spec verbatim** (§7 + §16): (i) each command tagged with its **auth generation** at handler-accept time (the current `seat_mint` pivot ref for that seat — one persisted `Cmd` field, so recovery replay is byte-identical per the §4 locus principle); (ii) the loop typed-rejects generation-mismatch — class **`credential-superseded`**, path-free, D-2 parity detail; (iii) **FX-B1g gains the in-flight leg** (old-session boot form queued pre-pivot ⇒ rejected, does NOT activate; the new credential's boot does).
2. **R2 per m-7's spec verbatim** (§0 + §16): the `seat_mint` accept-reply = the FIFTH I-PH payload family with two NAMED carve-outs (the fresh credential + the new seat's endpoint), operator-channel-only, matrix-asserted never-in-record/projection/log/non-operator-reply; the FX-D2 byte-parity divergence (accept-reply ⊃ record) stated as deliberate.
3. **F-S6-M1-4 guard line** (§5): commit-time last-writer guard — before appending any outcome with a non-empty `intake_id`, the loop checks `OutcomeByIntake`; an existing outcome ⇒ replay-or-fault, never a second outcome record. The sweep fixture is evidence the guard held, not the guard.

**Review chain (the s3 ratified rule — content-changing folds need fresh approving reviews):** r3 → your DESIGN-REVIEW request TO s6-core.implementer → its fresh approving verdict (cite m-7's relay + the fold hunks in the request so the re-review verifies fold-faithfulness) → **the gated PLAN re-parents to the NEW approving review** and keeps that edge across all plan revisions.

**Plan r2 (against design r3):** carries **F-S6-M1-1..5** and **F-S6-M2-1..4 verbatim** as hard acceptance criteria (m-2's render-absence ruling lands as stated: `visible_when` render absence + `gate_referenceable:false` + submit-path rejection; the boot rows named additive-MINOR outside the seven); adds the R1 task legs (the `Cmd` field + the reject class + the FX-B1g in-flight leg) and the R2 matrix family; both m-x route-back trigger lists ride the plan text verbatim. Then the Implementer PLAN-REVIEW (the F2 plan gate).

**Gate state after this directive:** m-2 = SATISFIED on plan-carry (no re-review) · m-7 = PRE-CONCURRED on faithful fold (its citation instruction binding) · m-1 = SATISFIED on plan-carry EXCEPT the R1 shape, which m-7 routed to m-1 and m-1's no-re-review grant did not cover — **I am routing the m-1 narrow confirm in parallel** (`s6-fidelity-m1/SITREP-orchestrator-planner-20260707-012810.md`); it is a dispatch precondition alongside your PLAN-REVIEW. Delegated dispatch when: {plan-r2 PLAN-REVIEW approve · the m-1 R1-confirm on record · SCOPE_DIFF all-in · no trigger/collision/amendment/OUT}.

ACTIONS_GIT_REF: none — reconciliation + fold directive only; no code/tracked-doc edit by this relay (the ledger entry commits separately).
FINAL_GIT_STATUS_SHORT: none — clean tree at relay-write time (HEAD `fe55082`, the pair's plan-r1 commit, verified docs-only by my `git show --stat`).
Next requested action: fold r3 → fresh pair review → plan r2 → PLAN-REVIEW; hold dispatch for the m-1 R1-confirm; report completion with the full chain ids.
