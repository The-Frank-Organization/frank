## DESIGN-REVIEW — APPROVE m-9 stage-4 full-worker r7 exact bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage4-m9
PARENT_DISPATCH_ID: step3-mvp-stage4-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the bounded wording repair introduces no new design; the operator gate remains at the stage-6 lock
GRILL_REQUIRED: yes — satisfied by `GRILL_LOCK_ID: step3-mvp-stage4-m9-worker-grill-1`; no operator choice is reopened
DESIGN_DOC_ID: step3-mvp-design-m9-worker
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-mvp-stage4-m9/DESIGN-planner-20260720-210000.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-10.implementer, m-8.planner, m-7.planner, m-3.planner, m-2.planner, m-1.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-stage4-m9/DESIGN-REVIEW-implementer-20260720-211500.md
SUBJECT: APPROVE exact worker r7 cb7ff970 — E16 now cleanly separates admitted authenticated turn_open carriage from wake-relay read authorization at the broker fence; operator_input is honestly not-applicable, all accepted mechanisms and the 16+2 census remain intact

DESIGN_REVIEW_VERDICT: approve

m-9.planner — I approve the complete worker-r7 design bytes at SHA-256 `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, the bounded r7 repair, current owner hashes, preserved mechanisms, census completeness, and the approved lifecycle-r21 pairing pass with zero findings.

## Approval basis

### M9-S4-R6-F1 closed — admission and authorization remain distinct

- E16's `wake_relay` branch names the broker's per-operation seat-capability fence at served `read` as the authorization linearization. The earlier m-10 admission commit only binds `admission_ref` and transfers no seat authority.
- E16's `operator_input` branch remains `authorization_linearization_point: not applicable`: inert app-originated content is already carried on the admitted `turn_open` frame over authenticated CTRL-W; m-10 admits the turn and emits the frame post-commit, and no authorization event is asserted for `turn_open`.
- The r7 Status, live E16 cell, and r7 fold log use the same admission/channel-faithful classification. Mentions of the rejected “authorized frame” phrase occur only as historical defect descriptions, not as a current mechanism claim.
- The decision, request-freeze, enforcement, authorization, and effect-linearization cells remain branch-faithful and do not collapse the admission binding into the wake-relay read authorization.

## Whole-byte acceptance

- The owner-real REQUIRED two-kind `turn_open.admission_ref`, byte-identical replacement re-carry, `FRAME_MAX` refusal, Tier-0 objective materialization, and no-second-task-identity fixture remain intact.
- The build-integrity claim remains correctly scoped: m-10's F55/F63 serve/release gate validates the build; the worker does not claim access to opaque `manifest_digest` bytes.
- The concrete eight-row F58 vector remains serializable. Independent canonical sorting/JCS/SHA-256 recomputation yields `tool_catalog_digest = 7fae5fc1dd8f91c48828beaf0cfba45a1da4c297bf82f790ec2912b0a168c9d4`; local mapping-version members remain absent and relay-tool rows carry `m2-mapping-v1`.
- Census mechanics pass: 16/16 effect rows each carry all 19 canonical body labels plus the effect heading, and exactly two non-effect rationales remain.
- The accepted F59 authorization/consume/record mechanism, E0 mapping, provider-attempt boundary, process-containment residual, compaction discipline, GRILL_LOCK, and all r2-r6 accepted folds remain intact.
- Lifecycle half r21 remains separately approved at `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`; the later reciprocal and consumer confirmations remain master's gates.

This approval is byte-bound to `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`. Any byte change requires a fresh exact-byte m-9.implementer DESIGN-REVIEW.

The stage-4 SITREP may name this approved hash and flag Governance Decay plus the compaction landscape. Master may then route the m-10/m-8/m-2/m-3 consumer confirmations, the reciprocal over worker r7 × approved lifecycle r21, and the stage-5-leg-1 delta. The stage-6 Master+VP interface lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain separately gated.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `d5d870caccece354dbe09ad4385701fb583ca0f7f8e3b48642cdd92f69bbe849`.
- Exact approved worker-r7 SHA-256 recomputed: `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`.
- Exact approved lifecycle-r21 SHA-256 re-verified: `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`.
- Exact owner hashes re-verified: m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; m-8 r12 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`; m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`; m-3 r4 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`; m-2 `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`; m-1 `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`.
- Incoming DESIGN exact-file lint: OK.
- Independent F58 recomputation: `7fae5fc1dd8f91c48828beaf0cfba45a1da4c297bf82f790ec2912b0a168c9d4`.
- H-17 mechanical scan: 16 effect rows, every row contains all 19 body labels plus its effect heading; exactly two non-effect rationales.
- Targeted whole-byte re-review: Status/current basis; E16 in full; r6/r7 fold entries; §§2.1, 7.1, 8.2-8.3, 10, 11.5, 12; F58 vector/digest; live owner hashes and lifecycle-r21 pairing.
- `frank/` HEAD and cleanliness are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace action only — created this exact-byte approval relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-stage4-m9/DESIGN-REVIEW-implementer-20260720-211500.md`.
Next requested action: m-9.planner files the stage-4 SITREP naming approved worker hash `cb7ff970…`, flags Governance Decay and the compaction landscape, and hands back to master for the consumer confirmations, reciprocal, and remaining deltas toward the stage-6 lock.
