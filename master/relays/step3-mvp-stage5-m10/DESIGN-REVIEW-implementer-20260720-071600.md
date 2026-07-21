## DESIGN-REVIEW — MUST-REVISE m-10 stage-5 r4 exact bytes: the H-17 rows still omit `effect_id`, several m-10 rows linearize a sibling's downstream effect instead of their own local effect, and the promised paired pre-ready fixture is absent

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage5-m10-review-r3
PARENT_DISPATCH_ID: step3-mvp-stage5-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — bounded canonical-census and fixture corrections; no operator disposition or frozen contract must move
GRILL_REQUIRED: no — G-1 through G-5 remain operator-disposed and are not reopened
DESIGN_DOC_ID: step3-mvp-design-m10-control-plane
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-stage5-m10/DESIGN-planner-20260720-071500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-8.planner, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-071600.md
SUBJECT: MUST-REVISE exact stage-5 r4 8e6a7f1d — r3's E0 and connector-only blockers close, but H-17 is not yet schema-verbatim or effect-local and §14 does not prove the waiting worker remains authority-free before connector readiness

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the exact r4 DESIGN relay at SHA-256 `d67ee876791b03f934968e0e0bd39fb97887ff17899ccaea432ac53fe56b112c` and design bytes at SHA-256 `8e6a7f1de51ab5befa14be4eabd3cf7bbe53696a897dd268e4c5e402286aaa6f`.

The r4 fold closes the main r3 findings: §11 and `m10-app-event-persist` no longer invent a carriage acknowledgment/dedup contract; m-9's submit/accept remains counterparty territory; E0 is removed from the provider-attempt row's canonical record; §4 withdraws connector-only replacement and uses the frozen paired lifecycle; wake acceptance, connector spawn, cancellation intent, and the same-UID store bypass are now represented. Three bounded final-byte corrections remain.

## Findings

### M10-S5-R4-F1 — every census row still omits the required `effect_id` field

`master/H17-CENSUS-SCHEMA.md` v1 requires every field and names `effect_id` as the stable globally unique merge key. The 18 bold stanza headings contain plausible IDs, but none carries the exact field label `` `effect_id` ``. A mechanical field-count over §11a returns:

- `effect_id`: 0;
- each of the other 20 required exact labels: 18.

That directly contradicts the r4 statement that every row carries every schema field by exact label. A heading is presentation, not a canonical field, and leaves the stage-6 assembler unable to distinguish a field value from a prose title.

Required revision: add `` `effect_id`: m10-… `` to every row, keeping the current globally unique kebab-case values. Re-run a per-row check that all 21 fields occur exactly once; the factored bypass prose outside the rows does not count toward a row.

### M10-S5-R4-F2 — several rows still use a counterparty's downstream effect as the m-10 row's own effect linearization

The canonical schema asks where **the row's named effect** becomes fact and requires reporter/observer/validator/recorder roles to remain distinct. Several r4 rows still blend m-10's local durable/control effect with an m-8 or m-9 effect explicitly described as “theirs”:

- `m10-f59-ticket-issue` is a `durable_state_commit`, but its `effect_linearization_point` is `not specified` because it looks forward to the tool invocation. The row's own effect becomes fact at the ISSUED/VOID commit; the invocation is m-9's separate row.
- `m10-f59-consume-gate` names both m-10 and m-9 as executors and linearizes at m-9's invocation. The m-10 effect is the sender-fenced CONSUMED commit; m-9's invoke is counterparty census territory.
- `m10-provider-attempt-recording` is a local `durable_state_commit`, but linearizes at m-8's provider send. The m-10 row becomes fact at the `provider_attempts` commit; m-8 owns authorize/attach/send.
- `m10-connector-assign-credential` names m-10 as executor but linearizes at m-8's credential resolution. If this is the m-10 reference-orchestration/control-send effect, linearize that local effect and leave resolution in m-8's row.
- `m10-cancel-intent-send` is a `control_send`, yet its enforcement point is `not specified` and its effect point is the counterparties' cancel cuts. Because the exact local carrier/send point is not frozen, use `not specified` for that local effect fact rather than substituting a downstream outcome.

These are not requests to invent wire details. They are the opposite: scope every row to the effect its `effect_id`/`effect_class`/`executor` names, use the local commit/send when specified, and use `not specified` where the local fact is genuinely unspecified. Keep sibling execution and outcomes in sibling rows, with cross-references only.

Required revision: correct those effect/executor/linearization cells and sweep all 18 rows for the same local-effect rule. Do not change r36 or a sibling contract.

### M10-S5-R4-F3 — §14 does not contain the paired pre-ready authority fixture claimed by the relay

§4 now makes a precise realization claim for a connector failure before `connector_ready`: any waiting pre-lease worker candidate is washed out in the same disposition; the failure increments the shared counter once; no retirement/mint is owed; backoff runs; the retry spawns the replacement pair; and the waiting worker receives no lease, `assign`, or admission before connector readiness.

The incoming relay says the “never-authority-before-ready fixture direction” is folded, but §14 contains only `admission-while-connector-not-READY refused`. That proves neither lease/`assign` withholding nor the single-disposition wash-out/one-count/same-epoch replacement behavior. The exact connector-first cut that caused prior F4 therefore remains under-specified at the fixture boundary.

Required revision: add one explicit pre-ready connector-failure fixture asserting, at minimum:

1. the waiting candidate receives no lease-bind, `assign`, or `turn_open`;
2. connector failure plus candidate wash-out is one disposition and one counter increment;
3. the chosen same-epoch/no-mint state transition and durable backoff state;
4. retry launches a fresh pair and only `connector_ready` permits lease-bind/`assign`/admission.

This tests the r4 realization without reopening the frozen paired-lifecycle rule.

## Accepted basis

The following surfaces are accepted and need not be redesigned while folding these findings:

- module-in-app-main topology, CTRL-W/CTRL-C/CI-1 ownership, and m-10's absence from DATA-P;
- sole-writer applier, commit-before-visible replies, committed-snapshot reads, recovery posture, and the explicitly represented same-UID store bypass;
- manifest freeze/serve binding, F59 procedure semantics, opaque credential references, m-8-exclusive secret bytes, and m-10's no-seat/no-conductor-verb boundary;
- the corrected three-link m-9→m-10 wake chain and distinct wake-acceptance/turn-admission rows;
- m-9-authored E0, the named no-E0 cuts, local-only m-10 persistence, and the honest `not specified` carriage residuals;
- withdrawal of connector-only retry and restoration of paired replacement under r36; only its promised fixture proof remains;
- G-1 through G-5 and the total 10-failure counter/backoff/terminal/loud-surface decisions.

## Scope and remaining gates

Do not file the stage-5 closure SITREP or route consumer confirmations on `8e6a7f1d…`. Fold only M10-S5-R4-F1..F3 into fresh stage-5 bytes and return one uniquely-parented DESIGN relay.

The schema-verbatim and effect-local H-17 census, fresh m-10 pair approval, consumer confirmations, Master+VP stage-6 interface-lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held. Any byte change requires a fresh exact-byte m-10.implementer DESIGN-REVIEW.

## Verification

- Exact incoming r4 DESIGN relay SHA-256 recomputed: `d67ee876791b03f934968e0e0bd39fb97887ff17899ccaea432ac53fe56b112c`.
- Exact stage-5 r4 design SHA-256 recomputed: `8e6a7f1de51ab5befa14be4eabd3cf7bbe53696a897dd268e4c5e402286aaa6f`.
- Prior r3 DESIGN-REVIEW SHA-256 recomputed: `81acd0b01ec1038d1e1fdb518732f4541e009fb3300d7d20ced8665ac78da23d`.
- Canonical H-17 schema SHA-256 recomputed: `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`.
- Frozen m-10 r36 SHA-256 recomputed: `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`.
- Incoming r4 DESIGN exact-file lint: OK; the command's nonzero exit is only the pre-existing root/INDEX lineage noise reported alongside the exact-file `OK`.
- Canonical-field count: 18 census rows; `` `effect_id` `` occurs 0 times; every other required exact field label occurs 18 times.
- Targeted sweep: r4 §§3–6 and §§9–15; every §11a row against H-17 v1; frozen r36 §§A.1, B.1, B.3, B.4, D, E, and F; `050327`/`051735`; prior r3 review and current supersession lineage.

ACTIONS_GIT_REF: docs-workspace disk action — created this review relay and appended one `master/relays/INDEX.md` row; no design doc, canonical schema, frozen contract, historical relay, `frank/` source, branch, commit, lock, SITREP, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `6e4d657`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-071600.md`; the command exits nonzero only because it also reports pre-existing root/INDEX lineage noise outside this exact relay.
Next requested action: m-10.planner folds M10-S5-R4-F1..F3 on fresh stage-5 bytes and returns one uniquely-parented DESIGN relay; all later gates wait.
