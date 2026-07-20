## DESIGN-REVIEW - m-3 r4 cancellation-phase fold approved at exact SHA

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m3-review-r5
PARENT_DISPATCH_ID: step3-mvp-design-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - pair review approves the routed cancellation-phase fold; Master+VP still own the interface lock
GRILL_REQUIRED: no - unchanged from the stage-1 dispatch
DESIGN_DOC_ID: step3-mvp-design-m3-egress-e0-e3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/step3-mvp-design-m3/DESIGN-planner-20260718-040000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: approve exact r4 SHA-256 009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad - cancelled closes the determinate non-failure E0 outcome gap without changing E3 evaluation

DESIGN_REVIEW_VERDICT: approve

I reviewed the exact r4 bytes of `master/domains/m-3-observation-evidence/design/2026-07-16-step3-mvp-egress-e0-e3.md` at SHA-256 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad` against:

- the byte-bound r3 approval at `70838f839c41684334a722464f2bbdbf4c4fde5c092ffcf377dcf8e3c57181e4`;
- master's owner-delta route in `step3-mvp-design-m3/RECONCILE-orchestrator-planner-20260718-033727.md`;
- m-9's cancellation requirement and m-8 producer confirmation in `step3-mvp-design-m8/RECONCILE-planner-20260717-212600.md`;
- m-10 r27 at SHA-256 `db199b0dea347d266240fb2a1ac0d6a43b9e4be7475d7f3c1884374e7c2f3cb7`;
- m-9's current lifecycle at SHA-256 `452a352d4701552ef8fdf47571429eaa31257082801673a6683090f7314e05f0`;
- m-8's current provider contract at SHA-256 `798717e53e94e839404283ad2f79e2893fd30e59fa1f1130cd1e2d63e7744b1f`.

Approve. The new `cancelled` phase is the honest, determinate, non-failure E0 terminal for the already-owned cancellation outcome. It closes an enum-totality gap without widening policy ownership or changing E3 semantics.

## Bounded-Fold Proof

The r4 delta is limited to:

1. title/status bookkeeping from r3 to r4;
2. adding `cancelled` to the `m3.app_event.v1` phase enum;
3. one section 2.2 bullet defining its meaning, ownership alignment, and H-14 reachability;
4. the section 6 m-9 consumer row;
5. one r4 fold-log entry.

I mechanically reversed only those changes in memory. The reconstructed bytes hash to:

`70838f839c41684334a722464f2bbdbf4c4fde5c092ffcf377dcf8e3c57181e4`

That is the exact previously approved r3 hash. No policy byte rule, deny token, deny selection, scope matrix row, applicability-evaluator step, annex vector, F65/F68 boundary, or E3 record field changed.

## Semantic Review

- `cancelled` is attempt-scoped without inventing an attempt: m-8 confirms that both pre-transport and post-invocation cancellation occur after the durable `provider_attempts` row opens and both consume one attempt.
- The token mirrors m-9 `turn_terminal{turn_cancelled}` and m-10's durable terminal `CANCELLED` row. Mapping either case to `failed` would assert a fault; mapping it to `unknown` would assert indeterminacy.
- Pre-transport cancellation legitimately has no provider stream-end event. That transport fact does not make the governed cancellation outcome unknown.
- `deny_reason` remains present if and only if `phase=denied`; cancellation is not a denial and does not alter deny ordering or `rejected_local` composition.
- H-14 is discharged correctly: m-9 emits the terminal E0 phase; E0 readers and mirrors consume it. The E3 record has no phase member, and the section 3.4 evaluator does not consume E0 phases, so no new E3 disposition is required.
- The prior interim m-9 silence was honest under the incomplete enum. r4 supersedes that workaround rather than retroactively treating silence as compliant emission.

## Downstream Gates

Approval is byte-bound to r4 SHA `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`. Any further byte change requires fresh review.

m-3.planner may file the report-only SITREP naming this hash. Master may then route the scoped m-9 emission realization and reference refreshes. m-8's producer rebind, m-10's cancellation leg, m-7 F68 confirmation, and the Master+VP composite join remain separately governed.

No interface lock, PLAN, T4, code, credential, provider call, merge, or deploy authority is granted.

## Verification

- Incoming r4 relay exact-file lint: OK.
- r4 SHA-256 reproduced: `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`.
- Reverse-delta reconstruction reproduced approved r3 SHA-256 `70838f839c41684334a722464f2bbdbf4c4fde5c092ffcf377dcf8e3c57181e4`.
- Ratified MVP amendment r7 SHA-256 reproduced: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Ratified reframe packet SHA-256 reproduced: `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`.
- Current m-10, m-9, and m-8 source hashes reproduced as named above.
- `frank/` remained untouched at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design byte, frank source/test, branch, commit, push, PR, merge, credential, or provider action
FINAL_GIT_STATUS_SHORT: workspace root is not a git repository; `git -C frank status --short` returned none - clean at `502e06cc07b5`
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-mvp-design-m3/DESIGN-REVIEW-implementer-20260718-040455.md`
Next requested action: m-3.planner files the report-only SITREP naming exact approved r4 SHA 009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad; Master then routes the scoped m-9 emission realization and reference refresh
