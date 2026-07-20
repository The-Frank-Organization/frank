## SITREP -- master.orchestrator-reviewer / readiness-review decomposition review

ROLE: Orchestrator Reviewer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: readiness-decomp
PARENT_DISPATCH_ID: readiness-decomp
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- review of read-only build-readiness decomposition; operator on CC
GRILL_REQUIRED: no
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator

Verdict: approve.

I reviewed:
- `master/relays/readiness-decomp/SITREP-orchestrator-planner-20260630-192926.md`
- `master/relays/c3-lock/RECONCILE-orchestrator-planner-20260630-191525.md`
- `master/README.md`
- `ROADMAP.md`
- `master/ARCHITECTURE.md` §C3
- `master/relays/INDEX.md`

Approve the build-readiness review, with the amendments below incorporated into the lens prompt and final
readiness register. No second decomposition relay is required if these amendments are carried forward exactly.

Finding 1 -- the phase boundary is correct.

The proposal stays in the AUDIT band: read-only review of the locked Step-0 design-of-record, no new design lock,
no code, no spike, no Step-1 plan, and no implementation authority. That matches the current state: c3 is closed,
Step-0 is complete, and Step-1 conductor-core remains a separate operator-opened gate per `ROADMAP.md` and
`master/ARCHITECTURE.md` §C3.7.

Finding 2 -- fresh orchestrator-run lenses are the right default.

Q2: concur. Run fresh red-team lenses first. Re-engaging the six domain pairs before finding a concrete gap would
reward local ownership bias and re-open standing-down seats too early. If a lens finds a real design gap, route a
bounded owner check/fix to the relevant pair or upstream domain, VP-gated, using the c3 Seam-C pattern. Build notes
and spike candidates stay in the register and go to the operator; they do not silently become pair work.

Finding 3 -- add two explicit dimensions; do not bury them inside lenses 1/4.

Q1: the five proposed lenses are good, but two dimensions are too load-bearing to remain implicit:

1. **Operator / HITL failure-path lens.** Trace the human path specifically: A/B/C/D bucket correctness,
   ODB render/capture, park/wake, stale-answer refresh, resummon cadence, operator-FROM verdict stamping,
   away-mode token return, and the "never auto-approve" invariant. This can be a sixth lens or an explicitly named
   required subtrace inside Composition & seams, but the final register must label it separately.
2. **Versioning / migration / replay lens.** Step-1 rides zero-migrator and append-only records, while C3 adds
   later carries like open named-axis ceilings, `external_send`, `certification`, and derived `surface_intent`.
   The readiness review must check absent-defaults, replay of older records, schema/header compatibility,
   and whether any future carry secretly requires destructive migration or canonical-data repair.

If the planner keeps five spawned lenses for cost, fold these as named mandatory subreports under lens 1 and lens 4.
Do not leave them as casual considerations.

Finding 4 -- refine the register schema before running the lenses.

The proposed register fields are close. Add:
- `step`: Step-1 blocker / later-step carry / unknown
- `owner_to_reengage`: none or exact domain/seat
- `operator_decision_required`: yes/no plus the concrete decision
- `evidence_ref`: source path and line or relay path; E0 agent intuition is not enough for a design-gap label
- `staleness_risk`: whether the item can drift before build

Also require each item to distinguish:
- **design-gap**: locked design contradicts itself, lacks a writer/reader, or cannot support Step-1 without a
  bounded design fix;
- **build-risk**: design is sufficient but implementation is hard and needs test/spike planning;
- **operator-choice**: design supports alternatives but a human value/policy choice remains.

Finding 5 -- no hidden spike or plan authority.

Q3: concur. A spike is code or build exploration and must be a separate operator-opened step. The readiness review
may recommend spikes and rank them, but may not run them, draft Step-1 tasks, or issue implementation authority.

Approved next actions:
- Run the read-only readiness lenses with the two added explicit dimensions above.
- Reconcile a build-readiness register with go/no-go and recommendation.
- Bring the register to the operator and VP before any Step-1 PLAN or spike.

Not authorized:
- no PLAN phase opening;
- no implementation, source/pcode edit, branch, commit, PR, merge, live verification, or spike;
- no standing pair re-engagement unless a concrete design-gap item is found and routed through a bounded VP-gated
  owner check/fix;
- no contract changes to c1/c2/c3 without a bounded fix relay.

ACTIONS_GIT_REF: wrote this review relay and appended `master/relays/INDEX.md`; docs workspace only, no code/source/pcode edits
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
