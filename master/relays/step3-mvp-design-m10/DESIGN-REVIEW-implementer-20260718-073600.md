## DESIGN-REVIEW — MUST-REVISE m-10 r29: check-1 family lacks a durable reason field, replay precedence, total classification, and consistent budget containment

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r30
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — all findings are bounded protocol-totality defects inside m-10 ownership; no operator product decision is required
GRILL_REQUIRED: no — this review does not introduce a hard-to-reverse choice
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260718-073500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-3.planner, m-7.planner, m-1.planner, m-2.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260718-073600.md
SUBJECT: MUST-REVISE exact r29 6418f820... — F80's family exists on wire now, but its durable/replay/classification/accounting contract is not yet total

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the exact r29 design bytes at SHA-256 `6418f8209c494c4967e2c676c8306244b1d375c4d2ce3702334f034d1f450108`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, and the single intended §D.2 amendment scope pass. `TURN_PARKED_UNKNOWN` is correctly withdrawn: its only occurrence in the current design is the withdrawal sentence. The new `authorize_reject` family, however, is not yet contract-total.

## Findings

### R29-F1 — BLOCKER: the claimed durable typed reason has no durable schema member

§D.2 says check-1 commits a VOID `tool_authorizations` row **with the typed reason** before replying. The normative ticket shape in §D.1 has `state`, timestamps, and `outcome_ref`, but no denial/VOID reason field; §F's `tool_authorizations` row adds no such member. A row in state VOID therefore cannot prove which `authorize_reject.reason` was committed, so the required durable evidence effect and same-reply replay claim are not implementable from the defined state.

Required revision: add the canonical durable reason member (or name an already-defined canonical row that actually stores it), its closed domain, presence/absence rule across every way a ticket reaches VOID, and the exact stored-reason-to-reply mapping. Fixtures must read the durable state, not merely observe the wire reply.

### R29-F2 — BLOCKER: idempotent replay contradicts the ordered check chain

§D.2 promises that a re-ask at an existing check-1 VOID row returns the SAME typed reply, including after the commit-before-reply crash. But uniqueness/existing-row handling is check (4), after mutable checks (1)–(3). A replay after the run/turn/lease/epoch/serve-gate state changes can therefore return a newly classified `authorize_reject`, `STALE_EPOCH`, `DENIED_ABOVE_SET`, or `DUPLICATE_REQUEST` before the stored VOID disposition governs. The crash-window promise is consequently not total.

Required revision: pin replay precedence for an existing `(run_id, turn_id, tool_call_id)` row before mutable-state reclassification, and define exact equivalent-versus-conflicting replay identity over the stored epoch/name/args digest. Cover at least: identical replay before state change, identical replay after turn/lease/epoch change, crash-after-VOID-commit replay, and same key with changed name or args digest. Every cut must return one deterministic typed result with no second row.

### R29-F3 — BLOCKER: the three reasons have no deterministic classification or total supervision rule

The three predicates are a conjunction, and their failures can overlap: a non-admitted run can also lack an active turn and valid lease; an inactive turn can coincide with a released lease. No sub-check order or exclusivity rule selects the persisted reason. The blanket “ordinary lifecycle race, not halves-divergence” rationale is also not total: a current-epoch, active-turn request whose lease alone is invalid is either unreachable under the lease invariants or evidence of divergence, not the stated turn-termination race.

Required revision: define an ordered decision table over the overlapping predicates, state which combinations are reachable, select exactly one durable reason for every reachable combination, and pin the worker/generation/turn/lease supervision effect per selected class. If a combination is invariant-impossible, give it an explicit fail-closed disposition and fixture rather than classifying it as an ordinary race.

### R29-F4 — BLOCKER: budget accounting is contradictory and the proposed fallback is not a bound

The new clause says check-1 rejects are NOT charged; the same §D.2 paragraph later says **denied authorization requests count** toward the per-turn ceiling, and §F says the ceiling is enforced at call insertion. No exception or counter predicate reconciles those statements. The rationale “no active turn exists” is false for the unclassified `lease_invalid`-with-active-turn combination. Separately, §A.3 bounded queues limit instantaneous pressure, not cumulative sequential requests: unique rejected `tool_call_id`s can commit uncharged VOID rows without ever filling a queue, so backpressure/channel-fault does not contain the stated runaway re-ask.

Required revision: state exactly which row classes increment which counter and make the general denial rule and §F insertion rule agree. For a rejection that cannot charge a turn, pin a real bounded disposition/resource rule; do not cite queue depth as a cumulative request/store bound. Add counter and resource-bound fixtures for every reason plus repeated unique rejected call IDs.

## Accepted portions and scope

- The reply-class, `re`-correlated `authorize_reject{tool_call_id, reason}` direction is an appropriate owner-level family.
- The three-token closed vocabulary is acceptable once its deterministic selection and durable storage are defined.
- Commit-before-reply, zero dispatch, and no automatic provider effect are the correct intended floor.
- `TURN_PARKED_UNKNOWN` remains withdrawn under D-4 option (a); do not resurrect it.
- No finding here reopens the pair-approved r28 surfaces outside the one §D.2 amendment. The replacement should remain one bounded r30 fold.

m-9's parallel consumer fold must wait for the corrected exact m-10 hash and consume the final reason/classification/supervision/accounting table, not the r29 draft. The m-10 SITREP, F73 rebind round, corrected reciprocal, stage-3 close, interface lock, stage-4/5, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any replacement bytes require a fresh uniquely-parented m-10.implementer DESIGN-REVIEW.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `bcdf3fedd51c98c4c217bdd76482b98c94b4ea427fb32651f13877dc184abd6e`.
- Exact m-10 r29 SHA-256 recomputed: `6418f8209c494c4967e2c676c8306244b1d375c4d2ce3702334f034d1f450108`.
- Incoming DESIGN exact-file lint: OK.
- `TURN_PARKED_UNKNOWN` occurrence count in current design: `1`, the withdrawal sentence.
- Targeted review: §A.3 queue bound; §D.1 durable ticket shape; §D.2 ordered issue checks, replay, budget, and fixtures; §D.4 VOID paths; §F table schema and store-write counter enforcement; m-9 r9 §3.3 current consumer set; VP F80 and Master `072316` correction grain.

ACTIONS_GIT_REF: docs-workspace disk action — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260718-073600.md`; root-wide legacy findings are outside this relay.
Next requested action: m-10.planner folds only R29-F1..R29-F4 as one bounded r30, binds the corrected exact hash for m-9's consumer fold, and requests a fresh uniquely-parented m-10.implementer review.
