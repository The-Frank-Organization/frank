## DESIGN-REVIEW - `scope_paths` predicate rev2 approved; the two owner halves now state one executable contract

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s9-scopepaths-m3-design-review-r3
PARENT_DISPATCH_ID: s9-scopepaths-m3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: s9-scopepaths-cosign
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/s9-scopepaths-m3/DESIGN-planner-20260713-171500.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-1.planner
SUBJECT: approve - rev2 folds the master rulings and m-2 co-sign into one PLAN-only segment-prefix contract, preserves the E0 and item-10 honesty rails, and labels every resolver mechanism as RED-first build work

DESIGN_REVIEW_VERDICT: approve

Rev2 closes MR-1 through MR-4 without reopening a locked decision. The master relay validly narrows the declaration home to accepted PLAN ancestors, blesses normalized segment-prefix grammar, and places narrowing at observe. The m-2 convergence relay supplies the missing owner co-sign: `scope-self-declared` is a submit-layer violation, while `scope-exceeded` is the observe-layer predicate with bounded reasons `{plan-widens-ancestor,diff-drift}`. Both halves now agree on the legal home, grammar, owner split, narrowing locus, token model, and exit fixtures.

### Approval Grounds

1. **The declaration home is renderable and singular.** `visible_when: {phase_in:["PLAN"]}` matches the accepted-PLAN-only ruling. Grant records are no longer a second scope source or an implicit s9 implementation dependency.
2. **The path relation is mechanically complete.** Segment-wise prefix matching distinguishes `pkg/a` from `pkg/ab`; child narrowing is a finite per-row membership check. The withdrawn glob language leaves no unbounded containment algorithm or unresolved complexity choice.
3. **Layer and token semantics are coherent.** A non-PLAN candidate copy is rejected by m-2's future submit guard and is never used as the observation bound. The observe predicate has one invariant and one token, with bounded reasons for declared widening and enacted drift. The latter remains item-10-gated.
4. **The resolver claims are honest.** The recursive parent-edge walk, nearest accepted PLAN stop, declaration filter, cycle/broken-chain handling, candidate-copy veto, matching, and dispositions are all named RED-first build obligations. Only the conductor-stamped parent edge is claimed extant at `39474d0`.
5. **The no-vantage and fault boundaries remain closed.** No declaration or no attributable LHS degrades at E0 and never silently passes; evaluable overreach is observed-false; ambiguous or broken lineage is a machinery fault. The item-10-dependent `diff_paths` evaluation stays struck.
6. **I-PH is explicit.** Scope failures expose only bounded symbolic class, bounded reason, and an optional non-path scalar. Raw paths, resolved rows, and prefix values are excluded from verdicts, rows, bounce detail, and degradation notes.

This approval covers the m-3 predicate-semantics half only. The joint contract returns to master only with m-2's reviewed leg and m-1's remaining section-F lineage-fidelity confirm. It does not grant PLAN or IMPL authority, move a design lock, unstrike item 10, or assert that any resolver machinery already exists.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW of `s9-scopepaths-cosign` m-3 rev2; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` edit, registry edit, PLAN, IMPL, token, or merge
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0

Next requested action: m-3.planner returns the approved m-3 half with m-2's co-signed leg and m-1's section-F fidelity result to master; B1 remains item-10-struck for enacted-diff evaluation and implementation remains held pending PLAN approval and token.
