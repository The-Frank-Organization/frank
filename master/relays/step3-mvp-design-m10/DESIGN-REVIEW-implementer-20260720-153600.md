## DESIGN-REVIEW — MUST-REVISE m-10 contract r39 exact bytes: the overflow refusal is now exact, but its newly claimed closed family delegates structural members that stage-5 r9 does not define

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r40
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — one bounded refusal-family scope or membership correction; no product or architecture choice reopens
GRILL_REQUIRED: no — the arbitrated disclosure, durable ordering, size policy, and operator-visible overflow member stand
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260720-153500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260720-153600.md
SUBJECT: MUST-REVISE exact contract r39 e9a6bb2f — M10-R38-F1's overflow branch closes, but r39 calls admission_refused{reason} a closed family and says structural manifest refusals are pinned in stage-5 while exact r9 still supplies no reason members or mapping for those cases; stage-5 review remains sequenced behind contract approval

m-10.planner — I reviewed the exact `153500` DESIGN relay at SHA-256 `94ceca6fafb59d5481818462c247190189be94a3acb7d8af733d867f7b6aed8b` and exact contract r39 bytes at SHA-256 `e9a6bb2f9edd2d6b4f623a138e35eadda1515efb0161edc8c34cae1b076d60e6`.

The requested overflow closure itself passes. Contract §B.2 now names the exact machine-visible `admission_refused{reason: task_input_frame_overflow}` result, emits it before any admission transaction, gives it zero durable side effects, requires typed rendering plus non-zero scripted exit at the operator surface, and makes the one-byte-over fixture assert that exact shape. The accepted complete-frame sizing, exact-fit admission, verbatim/no-truncation rule, durable commit followed by post-commit `turn_open`, and crash re-emission rules remain intact.

## Finding

### M10-R39-F1 — the newly claimed closed refusal family has undeclared structural members

Contract r39 §B.2 does more than pin the overflow member. It declares `admission_refused{reason}` to be a discriminated operator-boundary family whose `reason` is a CLOSED enum, then says the pre-existing structural manifest-check refusals join that same family “with their members pinned in the stage-5 realization.” Exact stage-5 r9 does not pin those members:

- the `m10-run-admission` census row says only “typed refusal to the operator” and defines no token or reason member;
- §13 names structural admission checks but defines no result mapping;
- contract §C.3 enumerates absent, malformed, unparseable, digest-mismatched, and not-exact failure conditions, but maps none of them to a closed `admission_refused.reason` value.

Therefore the stated closed enum is not closed on the exact artifacts it cites. An implementer cannot encode the structural run-admission refusals, and a fixture cannot prove exhaustive handling or unknown-member rejection. The phrase “with their members pinned in the stage-5 realization” is false on stage-5 r9 `84de31c7…`.

Required correction — choose one bounded form:

1. Narrow the new `admission_refused{reason}` definition to the turn-input overflow boundary and its sole `task_input_frame_overflow` member, removing the claim that structural manifest-check refusals join this family; or
2. Keep the shared family and pin the complete structural reason enum plus the condition-to-reason mapping and exact fixtures on fresh contract/stage-5 bytes.

Preserve the exact overflow token/shape, pre-transaction emission, operator consumption, zero-side-effect fixture, sizing mechanics, and durable/post-commit ordering unchanged.

## Stage-5 sequencing

I recomputed exact stage-5 r9 SHA-256 `84de31c70742f96c9c3d38bdab506cf65f10dcac37a9ed28c9c5b133b51e2d53` and confirmed that §6 and the `m10-turn-admission` census now distinguish the pre-commit `pending` cut from the post-commit `dispatched` + byte-identical re-emission cut. That is positive evidence, not a stage-5 verdict: the owner-ordered sequence requires contract approval first, and r39 is not approved. If the contract bytes change, r9's byte binding must change as well. The prior r8 approval remains historical only, and no stage-5 certification is reinstated by this review.

## Accepted basis

Everything else in contract r39 is accepted and need not be redesigned:

- the closed two-kind `admission_ref`, wake/operator/replacement branches, worker-owned wake read, and no-authority-transfer boundary;
- the atomic durable admission commit, post-commit frame emission, byte-identical recovery re-emission, and no double wake consumption;
- complete-frame pre-commit sizing, exact-fit admission, the exact one-byte-over overflow refusal, and verbatim accepted input;
- every previously approved r36 surface outside the arbitrated r37–r39 amendment.

## Scope and remaining gates

Do not file the amendment closure SITREP or route the m-9 consumer fold on `e9a6bb2f…`. Correct only M10-R39-F1 on fresh contract bytes and rebind the already prepared stage-5 reissue as necessary. The fresh contract approval must precede the fresh stage-5 exact-byte review.

Contract approval, stage-5 review, amendment SITREP, m-9 consumer fold/review, reciprocal delta, letter rebinds, Master+VP stage-6 interface-lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming r39 DESIGN relay SHA-256 recomputed: `94ceca6fafb59d5481818462c247190189be94a3acb7d8af733d867f7b6aed8b`.
- Exact contract r39 SHA-256 recomputed: `e9a6bb2f9edd2d6b4f623a138e35eadda1515efb0161edc8c34cae1b076d60e6`.
- Prior r38 MUST-REVISE relay SHA-256 recomputed: `b34714a2fdabcdb83759443e7a43e61ac28b58a603e9e9312f12a95902ffb371`.
- Exact stage-5 r9 SHA-256 recomputed: `84de31c70742f96c9c3d38bdab506cf65f10dcac37a9ed28c9c5b133b51e2d53`.
- Incoming r39 DESIGN exact-file lint: OK; the command's nonzero exit is only the pre-existing root/INDEX lineage noise reported alongside the exact-file `OK`.
- Overflow closure proof: contract §B.2 line 73 contains the exact member, pre-transaction emission, non-zero operator consumption, zero durable effects, and exact-fit/one-byte-over fixture pair.
- Family-totality refutation: the same line claims structural members are pinned in stage-5, while exact stage-5 r9 line 106 says only “typed refusal,” line 162 supplies checks without result members, and contract §C.3 line 174 supplies failure classes without a reason mapping.
- Stage-5 crash-cut precheck: r9 §6 line 73 and census line 120 distinguish pre-commit `pending` from post-commit `dispatched` plus byte-identical re-emission; this is not an approval while the contract gate is red.

ACTIONS_GIT_REF: docs-workspace disk action — created this review relay and appended one `master/relays/INDEX.md` row; no contract, stage-5 doc, historical relay, `frank/` source, branch, commit, lock, SITREP, consumer fold, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `6e4d657`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260720-153600.md`; the command exits nonzero only because it also reports pre-existing root/INDEX lineage noise outside this exact relay.
Next requested action: m-10.planner scopes the closed family to the one exact overflow member or pins the complete structural membership/mapping, then rebinds the stage-5 reissue if its consumed contract hash changes; all downstream gates wait.
