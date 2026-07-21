## DESIGN-REVIEW — MUST-REVISE m-10 stage-5 r6 exact bytes: the E+1 connector-failure correction is right, but §3 makes the waiting paired worker used by §4/§14 unreachable

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage5-m10-review-r5
PARENT_DISPATCH_ID: step3-mvp-stage5-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — one bounded lifecycle-reachability correction against frozen r36; no operator disposition or contract amendment is required
GRILL_REQUIRED: no — G-1 through G-5 remain operator-disposed and are not reopened
DESIGN_DOC_ID: step3-mvp-design-m10-control-plane
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-stage5-m10/DESIGN-planner-20260720-074500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-8.planner, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-074600.md
SUBJECT: MUST-REVISE exact stage-5 r6 0a419442 — the paired E+1 disposition now matches r36, but §3 still waits for connector_ready before allocating/spawning G1 while §4/§14 require G1 to be waiting unassigned before connector_ready

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the exact r6 DESIGN relay at SHA-256 `183812d10582b0e97859aa482f84779bc26829efd405efb6cce22eca36dcb42c` and design bytes at SHA-256 `0a419442d31b432ac2aa53e17ea61502c365512359d4c551f4a0b2c7815865fd`.

M10-S5-R5-F1's epoch correction is substantively right: §4 and §14 now use one paired retirement disposition, exactly one E+1 mint, §B.5 install before successor authority, one counter increment, durable backoff, and a fresh pair at E+1. The candidate-only same-epoch wash-out is correctly confined. One reachability contradiction remains.

## Finding

### M10-S5-R6-F1 — §3's start sequence cannot produce §4/§14's waiting worker candidate

The normative start sequence says:

1. §3 step 4 spawns only m-8 and completes `hello → connector_assign → connector_ready`;
2. only then does step 5 allocate G1, spawn the worker, lease-bind, and send `assign`.

But the corrected §4 transition and §14 fixture require a worker candidate to be **waiting before `connector_ready`** so connector-owner failure can retire the paired worker within the one E+1 disposition. The same §4 paragraph also still says the counter covers connector failures “BEFORE any worker generation exists.” On the written start machine, the fixture's waiting candidate is unreachable.

Frozen r36 §B.4 step 4 supplies the consistent order: launch the replacement pair; the worker may reach `READY` on `hello` but receives no `assign`; bootstrap the connector to `connector_ready`; only step 5 lease-binds and sends `assign`. Initial-run §B.3/§B.4 uses the same step-4→5 order. “Connector first” is an authority/bootstrap order, not “do not allocate or spawn the paired worker yet.”

Required revision:

1. Re-cut §3 steps 4–5 so G1 is durably allocated and the fresh DATA-P pair/children are launched as the pair; worker `hello` may reach pre-lease READY but `connector_ready` must precede lease-bind, `assign`, attach, and admission.
2. Replace “BEFORE any worker generation exists” with the exact cut actually designed, such as “before any worker generation is leased/assigned.” If a truly no-generation connector-spawn cut remains possible, specify it separately rather than using the waiting-candidate disposition.
3. Make §4 and §14 use the same reachable precondition: a paired, pre-lease, unassigned worker candidate exists and holds no authority when the connector fails.
4. Preserve the now-correct E+1/§B.5/one-count/backoff semantics unchanged.

This is a consistency/reachability fix inside the r6 design, not a new choice.

## Accepted basis

Everything else in r6 is accepted and need not be redesigned:

- the complete 21×18 H-17 census with 18 unique effect IDs and effect-local linearizations;
- the corrected paired connector-failure retirement, one E+1 mint, §B.5 install ordering, one counter increment, durable backoff, and fresh-pair retry;
- the no-lease/no-`assign`/no-admission-before-`connector_ready` authority boundary;
- all E0, wake, topology, applier, manifest, F59, credential, same-UID residual, grill, fixture, and loud-surface surfaces previously accepted.

## Scope and remaining gates

Do not file the stage-5 closure SITREP or route consumer confirmations on `0a419442…`. Correct only M10-S5-R6-F1 on fresh stage-5 bytes and return one uniquely-parented DESIGN relay.

Fresh m-10 pair approval, the stage-5 SITREP, consumer confirmations, Master+VP stage-6 interface-lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held. Any byte change requires a fresh exact-byte m-10.implementer DESIGN-REVIEW.

## Verification

- Exact incoming r6 DESIGN relay SHA-256 recomputed: `183812d10582b0e97859aa482f84779bc26829efd405efb6cce22eca36dcb42c`.
- Exact stage-5 r6 design SHA-256 recomputed: `0a419442d31b432ac2aa53e17ea61502c365512359d4c551f4a0b2c7815865fd`.
- Prior r5 DESIGN-REVIEW SHA-256 recomputed: `0f0225b32db8166d83b77cadf5ff4de5f289ce36f1ea2e1ef61605b738998016`.
- Canonical H-17 schema SHA-256 recomputed: `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`.
- Frozen m-10 r36 SHA-256 recomputed: `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`.
- Incoming r6 DESIGN exact-file lint: OK; the command's nonzero exit is only the pre-existing root/INDEX lineage noise reported alongside the exact-file `OK`.
- Canonical census recheck: 18 rows; all 21 labels present on every row; 18 unique `effect_id` values.
- Reachability proof: r6 §3 lines 45–46 versus §4 line 61 and §14 line 166; frozen r36 §B.4 lines 102–108 and initial-run ordering line 97.

ACTIONS_GIT_REF: docs-workspace disk action — created this review relay and appended one `master/relays/INDEX.md` row; no design doc, canonical schema, frozen contract, historical relay, `frank/` source, branch, commit, lock, SITREP, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `6e4d657`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-074600.md`; the command exits nonzero only because it also reports pre-existing root/INDEX lineage noise outside this exact relay.
Next requested action: m-10.planner aligns §3's start sequence with the reachable paired pre-ready state and returns one uniquely-parented DESIGN relay; all later gates wait.
