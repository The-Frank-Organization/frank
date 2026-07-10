## RECONCILE -- master.orchestrator-reviewer / build-readiness register review

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: readiness-reconcile
PARENT_DISPATCH_ID: readiness-reconcile
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- review of readiness register; operator decisions surfaced but not made here
GRILL_REQUIRED: no
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator

Verdict: approve.

I reviewed:
- `master/relays/readiness-reconcile/RECONCILE-orchestrator-planner-20260630-195008.md`
- `master/READINESS-REGISTER.md`
- `master/relays/readiness-decomp/SITREP-orchestrator-reviewer-20260630-193316.md`
- `master/domains/m-1-trust-identity/design/2026-06-28-trust-identity-design.md`
- `master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md`
- `master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md`
- `ROADMAP.md`
- `master/ARCHITECTURE.md`

Finding 1 -- GO-WITH-FIXES is the right top-level verdict, with a strict reading.

Q1: concur. I do not read any finding as no-go or unbuildable. I also do not read any of the planner-verified
findings as false positives. But `GO-WITH-FIXES` means proceed to bounded reconciliation, operator decisions, and
possibly a separately authorized spike. It does not mean Step-1 PLAN may open now.

Finding 2 -- the load-bearing verified gaps are real.

The m-1/m-2 submit-state conflict is real: m-1 records no append-as-submitted step and appends only after form +
observe pre-flights, while m-2 requires a `submitted` record so the cross-relay lineage engine can walk the graph
before acceptance. The m-3 placement conflict is also real: m-1 embeds observe-as-send inside `submit()`, while the
roadmap puts m-3 mechanism in Step 2. The m-2 R2 hole is real: the generic `field:<id>` atom and the stale
`selected_model` required-when wording conflict with the m-4/R2 lock that model is payload, never gate input.

Finding 3 -- fix routing is mostly right, but Cluster 2 is not just an operator-choice bucket.

Q2: approve the proposed routing with this tightening:
- Cluster 1: CTO+VP with m-1/m-2 is correct. This is a shared c1 contract collision and gates any store PLAN.
- Cluster 4a/4b: m-2 with m-4 review is correct and gates any store/form/schema PLAN.
- Cluster 2: needs operator decisions first, but also a bounded m-1 design-completeness fix. If Step-1 is attach,
  the design must downgrade the public claim honestly. If Step-1 is wrap, m-1 must state the wrap requirement,
  native-comms suppression, and I1 fallback/no-fallback boundary explicitly.
- Clusters 5a and 6a-c: operator decisions first, then owner fixes if the operator selects the stricter posture.
- Clusters 5b, 6d-g, 7, 8: bounded owner fixes are appropriate and can be scheduled after the MUST gate is moving,
  but they must not be silently dropped from the readiness trail.

Finding 4 -- sequencing gate.

Q3: the gate should be tighter than "open PLAN while fixes run." Before any Step-1 PLAN dispatch, require:
1. Cluster 1 reconciled or explicitly converted into a scoped operator-approved Step-1 subset that excludes the
   unresolved lineage/observe path.
2. Cluster 4a/4b reconciled in m-2/m-4 or the Step-1 schema work is held.
3. Operator decisions recorded for Cluster 2, 5a, 6a, 6b, and 6c.

Cluster 3 may become the first code spike only after a separate operator-opened spike gate. The SHOULD owner-fixes
may run in parallel with each other after routing, but not as a substitute for the MUST gate above.

Finding 5 -- readiness register quality is sufficient, but routing must be normalized before dispatch.

The register does cover the two dimensions I added in the decomp review: operator/HITL and versioning/replay. Before
routing fixes, emit either a short routing relay or a register appendix with one row per cluster item carrying:
`step`, `owner_to_reengage`, `operator_decision_required`, `evidence_ref`, `staleness_risk`, and classification
`design-gap | build-risk | operator-choice`. The current register has this information in prose, but the dispatch
layer needs the normalized table to avoid ambiguous pair re-engagement.

Approved next actions:
- Bring the operator the Cluster 2 / 5a / 6a / 6b / 6c decisions.
- Emit the normalized readiness routing table.
- Route bounded fixes for Cluster 1 and Cluster 4a/4b as MUST-before-Step-1-PLAN.
- Route the remaining bounded owner fixes as SHOULD, preserving them in the register.
- Recommend Cluster 3 as the first spike candidate only under a separate operator-opened spike gate.

Not authorized:
- no Step-1 PLAN opening yet;
- no code spike;
- no implementation, source/pcode edit, branch, commit, PR, merge, or live verification;
- no pair re-engagement without a bounded fix relay naming the exact cluster item, owner, and evidence;
- no contract changes outside the bounded readiness-fix lanes.

ACTIONS_GIT_REF: wrote this review relay and appended `master/relays/INDEX.md`; docs workspace only, no code/source/pcode edits
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
