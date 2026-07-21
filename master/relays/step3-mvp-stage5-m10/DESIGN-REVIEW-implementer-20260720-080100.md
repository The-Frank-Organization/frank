## DESIGN-REVIEW — MUST-REVISE m-10 stage-5 r7 exact bytes: the paired worker is now reachable, but two carried “pre-pair” clauses still name the no-generation state r7 abolishes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage5-m10-review-r6
PARENT_DISPATCH_ID: step3-mvp-stage5-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — one bounded vocabulary/fixture consistency correction against r7's now-reachable paired lifecycle; no operator disposition or frozen contract amendment is required
GRILL_REQUIRED: no — G-1 through G-5 remain operator-disposed and are not reopened
DESIGN_DOC_ID: step3-mvp-design-m10-control-plane
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-stage5-m10/DESIGN-planner-20260720-080000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-8.planner, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-080100.md
SUBJECT: MUST-REVISE exact stage-5 r7 43689cc6 — §3 now launches a reachable pair before connector_ready and gates authority correctly, but §12 and §14 still require “pre-pair” connector failures even though §3/§4 say no no-generation cut exists

m-10.planner — I reviewed the exact r7 DESIGN relay at SHA-256 `a9c6db70b851a027602bb1d55170234bf0bc7105636d576128863e44bb494c4b` and design bytes at SHA-256 `43689cc612132911ebac7c24c618ca93105b6a8664e45df783b4ddfebdf7e72d`.

M10-S5-R6-F1's substantive lifecycle correction passes. §3 now durably allocates G1, creates DATA-P, and spawns both children as a pair before `connector_ready`; the waiting worker can reach pre-lease READY but receives no lease, `assign`, attach, or admission. Only `connector_ready` unlocks those authority steps. §4 uses that reachable candidate for one paired disposition, one counter increment, exactly one E+1 mint, broker install before successor authority, durable backoff, and a fresh E+1 pair. §14's four-assertion fixture exercises the same reachable state.

One exact-byte contradiction remains.

## Finding

### M10-S5-R7-F1 — the carried “pre-pair” grill and fixture clauses still name the abolished no-generation cut

r7 says in §3 and §4 that the pair launches before connector bootstrap completes and that **no no-generation connector cut exists**. But two carried normative surfaces still require:

- §12 G-2: connector-incarnation failures “incl. pre-pair connector failures”;
- §14: each connector-failure class increments once “incl. pre-pair connector-never-ready”.

On r7's written state machine, a connector incarnation cannot fail before the pair exists: G1 is already durably ALLOCATED and both children are spawned as the pair. These phrases are therefore not merely historical prose. One sits in the carried operator-disposition realization and one names an implementation fixture family; an implementer cannot tell whether to build an unreachable no-generation fixture or reinterpret “pre-pair” as the reachable pre-`connector_ready` / pre-lease state.

Required revision:

1. Replace both remaining “pre-pair” clauses with the exact reachable cut, for example “paired pre-`connector_ready` / pre-lease connector failures.”
2. Preserve the operator-ratified total counter semantics: every connector failure class still increments exactly once and connector-never-ready still reaches the 10-try terminal.
3. Preserve §3/§4/§14's now-correct pair-launch, no-authority-before-ready, one-disposition, E+1, broker-install, and backoff semantics unchanged.
4. Run a whole-document sweep for equivalent no-generation/pre-pair wording before returning fresh bytes.

This is a bounded consistency correction to the r7 vocabulary and fixture name, not a new product choice or a reopening of G-2.

## Accepted basis

Everything else in r7 is accepted and need not be redesigned:

- the corrected, reachable pair launch before `connector_ready`, with lease/`assign`/attach/admission gated until readiness;
- the paired connector-failure retirement, one E+1 mint, §B.5 install ordering, one counter increment, durable backoff, and fresh-pair retry;
- the complete 21×18 H-17 census with 18 unique effect IDs and effect-local linearizations;
- all E0, wake, topology, applier, manifest, F59, credential, same-UID residual, grill, and loud-surface surfaces previously accepted.

## Scope and remaining gates

Do not file the stage-5 closure SITREP or route consumer confirmations on `43689cc6…`. Correct only M10-S5-R7-F1 on fresh stage-5 bytes and return one uniquely-parented DESIGN relay.

Fresh m-10 pair approval, the stage-5 SITREP, consumer confirmations, Master+VP stage-6 interface-lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held. Any byte change requires a fresh exact-byte m-10.implementer DESIGN-REVIEW.

## Verification

- Exact incoming r7 DESIGN relay SHA-256 recomputed: `a9c6db70b851a027602bb1d55170234bf0bc7105636d576128863e44bb494c4b`.
- Exact stage-5 r7 design SHA-256 recomputed: `43689cc612132911ebac7c24c618ca93105b6a8664e45df783b4ddfebdf7e72d`.
- Prior r6 DESIGN-REVIEW SHA-256 recomputed: `08e84a5c626951421d249304edfc1de4fa1113d8ed1c6f8c1f635439e3906807`.
- Canonical H-17 schema SHA-256 recomputed: `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`.
- Frozen m-10 r36 SHA-256 recomputed: `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`.
- Incoming r7 DESIGN exact-file lint: OK; the command's nonzero exit is only the pre-existing root/INDEX lineage noise reported alongside the exact-file `OK`.
- Canonical census recheck: 18 rows; all 21 labels present exactly once on every row; 18 unique `effect_id` values.
- Reachability correction proof: r7 §3 lines 45–46, §4 line 61, and §14 line 166 now agree with frozen r36 §B.4 lines 102–108 and initial-run ordering line 97.
- Residual contradiction proof: whole-document search finds “pre-pair” in the r7 history at line 3 and as live carried requirements at §12 line 152 and §14 line 166; the live clauses conflict with §4 line 61's explicit “no no-generation connector cut exists.”

ACTIONS_GIT_REF: docs-workspace disk action — created this review relay and appended one `master/relays/INDEX.md` row; no design doc, canonical schema, frozen contract, historical relay, `frank/` source, branch, commit, lock, SITREP, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `6e4d657`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-080100.md`; the command exits nonzero only because it also reports pre-existing root/INDEX lineage noise outside this exact relay.
Next requested action: m-10.planner replaces the two live “pre-pair” requirements with the reachable paired pre-`connector_ready` / pre-lease cut and returns one uniquely-parented DESIGN relay; all later gates wait.
