## DESIGN-REVIEW — MUST-REVISE m-10 stage-5 r5 exact bytes: the census is mechanically canonical and effect-local, but the new pre-ready connector fixture codifies same-epoch/no-mint against frozen r36's connector-failure E+1 transition

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage5-m10-review-r4
PARENT_DISPATCH_ID: step3-mvp-stage5-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — one bounded realization correction against frozen r36; no operator disposition or contract amendment is required
GRILL_REQUIRED: no — G-1 through G-5 remain operator-disposed and are not reopened
DESIGN_DOC_ID: step3-mvp-design-m10-control-plane
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-stage5-m10/DESIGN-planner-20260720-073000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-8.planner, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-073100.md
SUBJECT: MUST-REVISE exact stage-5 r5 47023a23 — F1/F2 and the fixture-presence portion of F3 close, but §4/§14 incorrectly apply same-epoch candidate wash-out to a connector-owner failure that r36 requires to replace both owners under E+1

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the exact r5 DESIGN relay at SHA-256 `fa0c18e605558fd1e0ccebdb00d0d4a5f4819c3f340bdcb31f6f4faac3d165d9` and design bytes at SHA-256 `47023a23451777aba5848dc7bfd94acb7d9d526e69803bb3af969b54503c782a`.

The r5 census fold is mechanically complete: all 21 canonical labels occur exactly once in each of 18 rows, all 18 `effect_id` values are unique, the seven effect-local corrections are present, and §14 now carries the requested four-part pre-ready fixture. Prior M10-S5-R4-F1 and F2 close. The fixture exposes one remaining frozen-contract conflict.

## Finding

### M10-S5-R5-F1 — connector failure cannot use the candidate-only same-epoch wash-out rule

§4 and the new §14 fixture say that when the connector fails before `connector_ready` while a worker candidate waits, the candidate is washed out in the same disposition with **no retirement and no epoch mint**, then the replacement pair is spawned under the same epoch.

Frozen r36 says the opposite for this trigger:

- §A.1: the DATA-P socketpair is bound to the worker-generation/connector-incarnation pair; death or replacement of **either owner** replaces **both under the new epoch** with a fresh pair.
- §B.1: a connector incarnation is never replaced alone; its replacement is both DATA-P owners, **new epoch**, fresh pair.
- §B.3: on connector FAILED/crash, run the same §B.4 canonical transition; step 1 retires the surviving worker generation and **mints E+1**, then both owners are reaped/spawned.
- §B.4's same-epoch wash-out is narrower: it applies when **the worker candidate itself fails pre-lease** (`no hello`, `hello-then-death`, or lease-bind failure) after revocation already occurred. A healthy waiting candidate being discarded because its connector owner failed is not that trigger.

The no-authority-before-ready portion is correct and should stay: the waiting worker receives no lease, `assign`, or `turn_open`. The one-disposition/one-counter-increment rule and durable backoff also stand. The conflict is solely the epoch transition attached to connector-owner failure.

Required revision:

1. In §4, replace the same-epoch/no-mint statement with the frozen connector-failure path: one retirement disposition for the paired owners, mint E+1, run the §B.5 distribution/install sequence, reap both, back off, and spawn the fresh pair at E+1.
2. In §14, change assertion (3) to prove the retirement commit minted exactly one E+1 and the broker installed it before successor lease-bind/`assign`; retain assertions (1), (2), and (4).
3. Keep the counter increment at exactly one for the connector failure; retiring/washing out the paired worker is part of that disposition and never a second failure count.
4. Sweep §11a failure semantics for any same-epoch connector-failure wording; the candidate-only wash-out remains valid only for candidate-originated pre-lease failures.

This is a correction to the r5 realization, not a request to amend r36 or reopen G-2.

## Accepted basis

The following surfaces are accepted and need not be redesigned:

- all 18 H-17 rows now contain all 21 exact labels with unique stable `effect_id` values;
- ticket issue, consume gate, provider-attempt recording, connector assignment, cancellation intent, attach gate, and epoch publication now linearize their own local effects honestly;
- E0 producer/visibility, app-event persistence, carriage residuals, and m-10's no-seat boundary;
- the three-link wake chain, separate wake-acceptance row, same-UID bypass accounting, and complete effect-family coverage;
- paired replacement and the absence of a connector-only retry path;
- the pre-ready no-lease/no-`assign`/no-admission gate, one disposition/one counter increment, durable backoff, replacement-pair requirement, 10th-failure terminal, and loud-surface direction;
- G-1 through G-5 and every previously accepted topology/manifest/F59/credential surface.

## Scope and remaining gates

Do not file the stage-5 closure SITREP or route consumer confirmations on `47023a23…`. Correct only M10-S5-R5-F1 on fresh stage-5 bytes and return one uniquely-parented DESIGN relay.

Fresh m-10 pair approval, the stage-5 SITREP, consumer confirmations, Master+VP stage-6 interface-lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held. Any byte change requires a fresh exact-byte m-10.implementer DESIGN-REVIEW.

## Verification

- Exact incoming r5 DESIGN relay SHA-256 recomputed: `fa0c18e605558fd1e0ccebdb00d0d4a5f4819c3f340bdcb31f6f4faac3d165d9`.
- Exact stage-5 r5 design SHA-256 recomputed: `47023a23451777aba5848dc7bfd94acb7d9d526e69803bb3af969b54503c782a`.
- Prior r4 DESIGN-REVIEW SHA-256 recomputed: `97acf5e6622542c466cdc80c58fe25d2f70d7d871959e0b14db5a39b7732900d`.
- Canonical H-17 schema SHA-256 recomputed: `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`.
- Frozen m-10 r36 SHA-256 recomputed: `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`.
- Incoming r5 DESIGN exact-file lint: OK; the command's nonzero exit is only the pre-existing root/INDEX lineage noise reported alongside the exact-file `OK`.
- Canonical-field proof over the 18 row stanzas: every one of the 21 required labels occurs 18 times; 18 unique `effect_id` values; 18 rows.
- Frozen-contract proof: r36 §A.1 line 25, §B.1 line 64, §B.3 line 82, and §B.4 lines 101–108; the candidate-only same-epoch rule is line 108, while connector failure uses the E+1 path at lines 82/102–107.
- Targeted sweep: r5 §§3–6 and §§11–15; every §11a row; frozen r36 §§A.1/B.1/B.3/B.4; prior r4 review and current supersession lineage.

ACTIONS_GIT_REF: docs-workspace disk action — created this review relay and appended one `master/relays/INDEX.md` row; no design doc, canonical schema, frozen contract, historical relay, `frank/` source, branch, commit, lock, SITREP, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `6e4d657`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-073100.md`; the command exits nonzero only because it also reports pre-existing root/INDEX lineage noise outside this exact relay.
Next requested action: m-10.planner corrects the pre-ready connector-failure epoch transition on fresh stage-5 bytes and returns one uniquely-parented DESIGN relay; all later gates wait.
