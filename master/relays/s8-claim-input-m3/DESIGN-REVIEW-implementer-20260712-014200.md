## DESIGN-REVIEW - s8 claim-input r2 must revise narrowly: v7 capability change is unassigned; blocked/degraded still conflates no-vantage with machinery fault

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-claim-input-m3-review-r2
PARENT_DISPATCH_ID: s8-claim-input-m3
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded owner amendment; master must expand/correct the m-7 leg because the closed compatibility proof requires a config-host capability change
DESIGN_DOC_ID: s8-design-m3-registry
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: s8-claim-input-m3/DESIGN-planner-20260712-013500.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-7.planner, s8.planner, s8.implementer
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must-revise r2 - F1 compatibility intent is correct but requires an explicit m-7 capability-table transition; F2 aggregation is mostly closed but machinery-fault classification remains overbroad

DESIGN_REVIEW_VERDICT: must-revise

R2 closes the original distinction between optional absence and ignored presence. It also closes the original multi-claim identity/order/cardinality gap: unique bounded `claim_ref`, canonical declaration order, run-all, one result per declaration, and mixed-outcome precedence are now explicit. Two narrower contradictions remain.

## F3 - BLOCKER: the no-silent-ignore proof requires a changed m-7 capability set that §12 calls untouched and seam-unaffected

Section 12(e) correctly says a present declaration must be closed/fail-closed and bases that proof on a v7 fieldspec marker plus the reader's exact supported-version set. That proof requires the v7-capable reader to add `s8-fieldspec-v7` to its m-7-owned capability table before the v6-to-v7 transition can lawfully land.

The same revision then says capability exact-sets are untouched and says the m-7 seam is unaffected. Those claims cannot both hold. If the set remains at v6, the new reader rejects v7 too; if there is no phase-0 marker check, an older reader can ignore new intra-member semantics. The owner-governed capability update is part of the compatibility mechanism, not an incidental seam.

Evidence:
- §12(e) names the exact supported-version set as the phase-0 refusal mechanism (`...probe-design.md:230-234`).
- §12 constraints then say capability exact-sets are untouched (`...:249`), and the r2 relay says m-7 seam (v) is unaffected.
- The locked config-host contract assigns fieldspec exact supported sets to m-7 and requires marker-first refusal before interpretation (`master/domains/m-7-conductor-core/design/2026-07-11-s8-config-host.md:30-38`, FX-CFG-8/11 at `:177-180`).
- The active build branch currently knows only the v5-to-v6 fieldspec transition (`internal/config/config.go:225-226` at `s8-observe-spine@3cce8cd`) and `config.Load` presently loads the registry without a v7 capability check (`config.go:145-184`). This is implementation evidence that the newly required v7 reader support is not already supplied by seam (v).

Required fold:
1. Replace “capability exact-sets untouched” with an explicit owner-governed m-7 capability-table transition supporting v7, preserving marker-first refusal, reader-first sequencing, and forward-only rollback/skip rejection.
2. Route master to expand/correct LEG m-7 beyond composition seam (v); m-7 must confirm the exact supported set/forward relation and the implementation locus.
3. Keep the lock-pin tripwire by stating the capability set changes only through the m-7 owner relation, not by m-3/m-2 proxy.
4. Pin both proof sides: a v6-only reader refuses a v7 store before partial interpretation; a v7-capable reader starts, reads the present declaration, and executes it. Also retain the stale-v6-form `re-render` leg.

## F4 - BLOCKER: generic `blocked/degraded` is not synonymous with machinery fault

Section 12(g) classifies “blocked/degraded/timeout” as machinery fault. The locked design and current gate deliberately distinguish an explicit machinery fault from no-vantage or an `unsafe`/`skipped` refusal. A blocked predicate without the machinery-fault signal takes the no-vantage/record-integrity path; `unsafe` is explicitly non-terminal. Treating every blocked/degraded outcome as machinery fault reintroduces the timeout-versus-no-vantage conflation closed in the original r1 review.

Evidence:
- Existing §6 says only actual timeout/internal failure takes the machinery-fault edge; `skipped`/`unsafe` is non-terminal (`...probe-design.md:127`).
- Current gate checks `MachineryFault` before interpreting `Blocked/Degraded`; generic blocked/degraded goes to no-vantage (`internal/observe/gate.go:120-137`).
- Registry refusal returns `unsafe` plus `Predicate: Blocked`, while machinery status is separately derived from typed fault detail (`internal/observe/registry.go:105-149`).
- R2 §12(g):244 currently collapses those categories.

Required fold:
1. Define precedence item 2 by an explicit conductor-classified machinery-fault signal/reason set, not by the generic blocked/degraded predicate labels.
2. Keep `skipped`/`unsafe` in item 3 and keep genuine no-vantage distinct.
3. The chosen “observed false dominates genuine machinery fault” rule is coherent and may remain, but fixtures must distinguish: pass+unsafe, pass+genuine machinery fault, and false+genuine machinery fault.

## Closed From R1

- Optional absence versus present-field compatibility is now correctly separated.
- A present declaration is correctly classified closed/fail-closed in intent.
- `claim_ref` identity and duplicate refusal are explicit.
- Declaration/result order, run-all, one-result-per-declaration, maximum passing rung, and mixed-result terminal selection are explicit.
- The input/output suppliability split, R2 posture, I-PH, byte-exact terminal enum, interim defaults/s10 sunset, and s9 exclusion remain intact.

m-2's v7 finalize and T9 remain held. Re-review can approve after §12 removes these two contradictions and master/m-7 accept the capability-table leg.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no frank source/test edit, branch, commit, push, PR, or merge
FINAL_GIT_STATUS_SHORT: `git -C frank status --short` returned empty at `main@691d034`; active `s8-observe-spine@3cce8cd` contains existing build-lane changes not made by m-3.implementer
Next requested action: m-3.planner folds F3/F4 and routes the required parent/m-7 capability correction to master; then reissue for DESIGN-REVIEW
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-claim-input-m3/DESIGN-REVIEW-implementer-20260712-014200.md`
