## DESIGN-REVIEW — MUST-REVISE m-10 stage-5 r3 exact bytes: the census is still not canonical/complete, E0 carriage invents an unowned acknowledgment contract, and the pre-pair connector-only retry contradicts frozen r36

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage5-m10-review-r2
PARENT_DISPATCH_ID: step3-mvp-stage5-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — all findings are bounded corrections against the canonical census and frozen r36; if connector-only retry is desired, that contract change must route up
GRILL_REQUIRED: no — G-1 through G-5 remain operator-disposed; this review does not reopen them
DESIGN_DOC_ID: step3-mvp-design-m10-control-plane
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-stage5-m10/DESIGN-planner-20260720-063000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-8.planner, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-063100.md
SUBJECT: MUST-REVISE exact stage-5 r3 207a6519 — prior E0-producer and wake-boundary findings close, but the H-17 inventory omits effects/required cells and masks the same-UID store bypass; E0 carriage asserts an unfrozen ack/dedup protocol; connector-only pre-pair retry conflicts with r36's never-alone replacement rule

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — the operator first handed me `DESIGN-planner-20260720-061500.md`. While its review was underway, Master filed the `051735` hold and you filed `063000`, explicitly superseding both `054000` and `061500` as non-closing lineage. This verdict therefore binds only the current, uniquely-parented r3 relay at SHA-256 `4231312fffd583d43114df57a15c2448fbaab124f571ecbfc7ebc9ecafb0bda9` and design bytes at SHA-256 `207a6519d509c9c947c8670b1c26c14665ae6624052a74de29a21c780ad1d804`.

The r3 fold closes prior findings M10-S5-R1-F2 and F3: §11 now keeps m-9 as E0 author/carrier and names the worker-death/retirement-winning no-E0 cuts, while §6 assigns conductor push plus durable `project`/`read` rediscovery only to m-9 and confines m-10 to `wake_forward` insertion/admission. The total supervised-child counter, 10th-failure commit, and expanded fixture direction also close the reachability portion of prior F4. Three final-byte blockers remain.

## Findings

### M10-S5-R3-F1 — §11a is not yet a canonical or complete H-17 inventory

The heading says every stanza has all canonical fields, but the stanzas use compressed aliases and omit required cells. Examples:

- `m10-app-event-carriage` has no `authorization_linearization_point`;
- `m10-child-terminate` has no `authorization_linearization_point` and collapses `decision_point` with `enforcement_point`;
- many rows collapse `policy_owner` and `policy_artifact` into an unlabeled `policy`, and use `credential`, `freeze`, `auth-lin`, `effect-lin`, `reporter`, `observer`, `validator`, `record`, `bypass`, and `failure` rather than the schema's required field names.

That prevents the stage-6 assembler from distinguishing an intentionally null field from an accidentally absent one. `master/H17-CENSUS-SCHEMA.md` v1 requires every field and says missing facts use exactly `unknown`, `not specified`, or `residual`; “one stanza per row, all 21 fields” is not true on these bytes.

Coverage is also incomplete:

- §6 link (2), the idempotent `wake_forward` acceptance/`wake_schedule` insert, is a durable authoritative transaction distinct from link (3)'s turn-admission transaction. `m10-turn-admission` describes only the latter and cannot silently absorb both linearization points.
- Connector OS spawn is absent. `m10-worker-spawn-assign` is worker-specific, while `m10-connector-assign-credential` begins only after connector `hello`; the connector spawn whose failures now increment the total counter has no process-lifecycle row.
- Cancellation/control sends are not mapped. `m10-run-stop-cancel` owns durable cancellation state and `m10-child-terminate` owns the shutdown/signal ladder, but no row or explicit non-effect rationale accounts for the other cancellation/control messages named by Master's H-17 bar. Where a frozen contract does not specify one, the census must say `not specified`, not omit it.

The common bypass claim is internally false. §11a says the accepted same-UID residual lives in §9 and each row, but §9 does not name direct same-UID SQLite access and most durable-state rows say `bypass_paths: none`. Under the accepted threat ceiling, another same-user process can directly open/mutate the private database; that is an alternate effect path around the applier for applicable state effects. State it in a factored common `bypass_paths` rule that explicitly applies to those rows, or repeat it in each applicable row. A common `threat_claim_scope` does not populate `bypass_paths`.

Required revision: render every row with the canonical field labels and an explicit value/null token for every required field; add distinct rows for wake acceptance and connector spawn; map cancellation/control-send effects or give explicit non-effect rationales; and correct the same-UID store bypass. Keep the stable `effect_id`s and do not invent missing facts.

### M10-S5-R3-F2 — the app-event row invents an acknowledgment/dedup contract and conflates two owners' effects

The prose correctly says m-10 persists an m-9-authored frame and m-9 carries it to the conductor. The census then combines both acts into `m10-app-event-carriage` and states that rows persist “until carriage-acknowledged” with “no duplicate submission (relay-side identity).” No cited frozen m-10/m-9/m-3 contract defines an `app_event` carriage acknowledgment, its correlation key, its m-10 receipt path, the durable transition that marks a row carried, or the replay/dedup rule. r36 fixes only m-9 delivery of `app_event` to m-10 storage and the worker seat's carriage to the conductor. The canonical schema forbids filling a missing contract fact by inference.

The same row also omits the authorization linearization point and names the conductor accept as this row's effect linearization even though that submit is an m-9/m-7 effect outside m-10's seatless boundary. Separately, `m10-provider-attempt-recording` lists E0 rows in its canonical record even though E0 creation is conditional, m-9-authored, and separately inventoried.

Required revision: make m-10's received-event validation/persistence its own exact effect row, with its local commit as the effect linearization point. Treat m-9's conductor submit/accept as counterparty census territory. Mark acknowledgment, retry, and cross-owner dedup behavior `not specified`/`residual` unless an exact frozen source defines it; do not promise no duplicate submission here. Remove E0 rows from the provider-attempt row's unconditional canonical-record claim.

### M10-S5-R3-F3 — the pre-pair connector-only retry contradicts frozen r36 and this design's own §5

§4 says a connector failure before a worker generation exists retries the connector alone because r36 is “not specified.” §5 simultaneously says a connector fault is never repaired alone and the supervisor has no restart-connector-only verb. Frozen r36 is broader than the r3 interpretation:

- §A.1: death/replacement of either DATA-P owner replaces both under a new epoch;
- §B.1: “A connector incarnation is never replaced alone”;
- §B.4: the canonical transition order is shared by connector failure, and step 4 launches the replacement pair; the worker may reach `hello`/READY but receives no `assign` until connector bootstrap and lease-bind complete.

That frozen ordering already supports connector-first startup without granting worker authority: preallocate/spawn the replacement pair, allow the worker to wait unassigned/unleased, count the connector failure exactly once, and repeat under the frozen wash-out/transition rules. r3 instead adds a connector-only replacement path that the exact contract expressly excludes.

Required revision: remove the connector-only exception and make the total counter ride the frozen paired lifecycle, including connector failure before readiness; add the connector-spawn census row and fixtures that prove the waiting worker never receives lease/assign/admission before `connector_ready`. If connector-only retry is still preferred, route a narrow r36 amendment through the owners rather than claiming no frozen byte moves.

## Accepted basis

The following surfaces are accepted and need not be redesigned while folding the findings:

- the module-in-app-main topology, CTRL-W/CTRL-C/CI-1 ownership, and absence of m-10 from DATA-P;
- the sole-writer serialized applier, commit-before-visible replies, committed-snapshot reads, and separate m-10 store/code;
- manifest freeze/serve binding, F59 issue/consume/record realization, opaque credential-reference handling, and m-8-exclusive secret bytes;
- m-10's no-seat/no-conductor-verb boundary and the corrected three-link m-9→m-10 wake chain;
- the corrected m-9-authored E0/no-E0 boundaries in §11, apart from the unfrozen carriage-ack promise;
- G-1 through G-5, including the one total 10-failure counter, durable backoff, reset-on-completed-turn, 10th-failure terminal commit, and loud-surface/fixture direction; only the frozen paired-lifecycle realization must change;
- the stated ambient-authority, same-UID, and build-identity threat ceiling, once represented honestly in the census.

## Scope and remaining gates

Do not file the stage-5 closure SITREP or route consumer confirmations on `207a6519…`. Fold only M10-S5-R3-F1..F3 into fresh stage-5 bytes, preserve the accepted grill decisions and frozen sibling contracts, and return one uniquely-parented DESIGN relay.

The canonical complete H-17 census, fresh m-10 pair approval, consumer confirmations, Master+VP stage-6 interface-lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held. Any byte change requires a fresh exact-byte m-10.implementer DESIGN-REVIEW.

## Verification

- Exact incoming r3 DESIGN relay SHA-256 recomputed: `4231312fffd583d43114df57a15c2448fbaab124f571ecbfc7ebc9ecafb0bda9`.
- Exact stage-5 r3 design SHA-256 recomputed: `207a6519d509c9c947c8670b1c26c14665ae6624052a74de29a21c780ad1d804`.
- Superseded r2 DESIGN relay SHA-256 recomputed: `0aa6d737ae6e0bd4262be76145cb588da942883df0704ee3bb1301e5b8f500d0`; its relay-bound design hash was `651a400b1aca63692b5926605521b9ae3772a633fc69c9ee91858ba9ad239417`.
- Prior r1 DESIGN-REVIEW SHA-256 recomputed: `9993c1b39d79c4774a2173ed245a72b867417f7315e0283a791e8d13eaf8ae81`.
- Master's `051735` hold SHA-256 recomputed: `3ef346c034a9a421d00fb54a7553d35292d4b587d6dfd501fc98cb2b7bac4a0f`.
- Canonical H-17 schema SHA-256 recomputed: `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`.
- Frozen m-10 r36 SHA-256 recomputed: `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`.
- Incoming r3 DESIGN exact-file lint: OK.
- Targeted sweep: current r3 §§4–6 and §§11–15; all §11a census rows against every canonical H-17 field/coverage rule; frozen r36 §§A.1, B.1, B.3, and B.4; the operative Master hold; prior pair review and superseded lineage.

ACTIONS_GIT_REF: docs-workspace disk action — created this review relay and appended one `master/relays/INDEX.md` row; no design doc, canonical schema, frozen contract, historical relay, `frank/` source, branch, commit, lock, SITREP, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `6e4d657`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-063100.md`; the command exits non-zero only because it also reports pre-existing root/INDEX lineage noise outside this exact relay.
Next requested action: m-10.planner folds M10-S5-R3-F1..F3 on fresh stage-5 bytes and returns one uniquely-parented DESIGN relay; all later gates wait.
