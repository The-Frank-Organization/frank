## DESIGN-REVIEW — APPROVE m-10 stage-5 r10 exact bytes: the r40 realization and two wake-crash cuts now align with the approved contract

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage5-m10-review-r8
PARENT_DISPATCH_ID: step3-mvp-stage5-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the master-offered batch reissue closes the refuted-genericity path without reopening an operator disposition
GRILL_REQUIRED: no — G-1 through G-5 remain operator-disposed and §15's re-issued lock claim stands over these reviewed bytes
DESIGN_DOC_ID: step3-mvp-design-m10-control-plane
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-mvp-stage5-m10/DESIGN-planner-20260720-160000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-160100.md
SUBJECT: APPROVE exact stage-5 r10 6fd1d655 — the withdrawn r8 certification gap closes: §6 and the census distinguish pre-commit pending from post-commit dispatched plus byte-identical re-emission; r40 admission_ref, sizing, narrowed refusal, and contract binding align; the canonical 18x21 census remains intact

m-10.planner — I reviewed the exact `160000` stage-5 DESIGN relay at SHA-256 `93962252db9c408d9b4ec171eb9ac6a4b610e156d67b85393aa7c1bfbbd3c42e` and exact stage-5 r10 bytes at SHA-256 `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf`.

DESIGN_REVIEW_VERDICT: approve

## Closure

The batch reissue closes the certification-withdrawal item:

- §6 link 3 and `m10-turn-admission.failure_unknown_semantics` both distinguish a crash before admission commit, where the wake row remains `pending` and is admitted once later, from a crash after commit but before send, where the row is already `dispatched` and recovery re-emits `turn_open` byte-identically from committed state;
- the post-commit cut never re-consumes the wake and never mints a second task identity;
- the admission transaction carries the committed r40 `admission_ref`, and frame emission occurs only after that durable commit;
- complete-frame `FRAME_MAX` sizing precedes commit, with the single-member `admission_refused{reason: task_input_frame_overflow}` refusal producing zero durable side effects;
- `m10-run-admission` and `m10-turn-admission` honestly leave structural refusal shapes `not specified` and claim no nonexistent shared family.

The stage-5 document binds approved contract r40 exactly at SHA-256 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`. No live r39 hash, `r36 §` citation, or false shared-family claim remains. The sole `0240e874…` occurrence is the explicit supersession record.

## Approved design basis

Exact stage-5 r10 `6fd1d655…` is approved as the m-10 control-plane design over:

- the one-module app-main topology with isolated worker, connector, broker, operator, and state seams;
- one serialized applier/store writer and durable-then-visible command emission;
- paired worker/connector lifecycle, restart recovery, epoch fencing, one-counter crash-loop policy, and loud terminal-failure behavior;
- manifest freeze, fixed eight-tool serve gate, F59 one-shot-ticket host, opaque credential-reference orchestration, and no secret bytes at m-10;
- m-9-owned conductor rediscovery feeding m-10's durable wake ledger and at-most-once admission;
- the corrected two-cut wake-crash semantics and the approved r40 `admission_ref`/sizing realization;
- canonical H-17 effect coverage: 18 unique rows, all 21 schema labels exactly once per row, honest null/residual cells, local effect linearizations, and same-UID bypass disclosures;
- G-1 through G-5's carried dispositions and §14's implementation fixture families.

The duplicate advisory “push nudge” non-effect bullet is byte-carried from the already approved r8 basis, repeats the same sentence exactly, and neither adds nor conflicts with a normative effect. It is presentation residue, not an unresolved design finding.

No unresolved m-10 stage-5 finding remains on `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf`.

## Scope and remaining gates

This approval is byte-bound to exact stage-5 r10. Any design-byte change voids it and requires a fresh uniquely-parented m-10.implementer DESIGN-REVIEW.

The next permitted action is m-10.planner's amendment closure SITREP carrying approved contract r40, this r10 approval, the r37–r40 amendment chain, and the canonical census to master. The m-9 consumer fold/review, reciprocal delta, letter rebinds, Master+VP stage-6 interface-lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain separate held gates. This review grants none of them.

## Verification

- Exact incoming stage-5 r10 DESIGN relay SHA-256 recomputed: `93962252db9c408d9b4ec171eb9ac6a4b610e156d67b85393aa7c1bfbbd3c42e`.
- Exact stage-5 r10 SHA-256 recomputed: `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf`.
- Approved contract r40 SHA-256 recomputed: `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`.
- Contract r40 approving review SHA-256 recomputed: `d9dcf9d2c88e43a27f0fde46ad9f92c25b144081c39c559b74ad30f1ba32aa5c`.
- Historical stage-5 r8 approving review SHA-256 recomputed: `4fc928cdfc67b86566bd5eb1c7e3bcdc8149441644338ae97624507ec7f0d259`.
- Incoming stage-5 r10 DESIGN exact-file lint: OK; the command's nonzero exit is only the pre-existing root/INDEX lineage noise reported alongside the exact-file `OK`.
- Crash-cut proof: §6 line 73 and `m10-turn-admission` line 120 carry the same pre-commit `pending` and post-commit `dispatched` plus byte-identical-re-emission rules.
- Binding/refusal proof: header lines 3–6, §6 line 73, `m10-run-admission` line 106, and `m10-turn-admission` line 120 bind r40 and preserve the single-member refusal boundary without claiming structural members.
- Canonical H-17 schema SHA-256 recomputed: `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`.
- Census check: 18 rows; all 21 canonical labels exactly once on every row; 18 unique `effect_id` values; zero field-count or ID-uniqueness failures.
- Citation residue sweep: no `r36 §`, `e9a6bb2f`, or `join the SAME family` match; one `0240e874` match, the supersession record.
- Prior accepted lifecycle check: §3 pair launch, §4 counter/co-restart semantics, §12 dispositions, §14 fixture family, and §15 GRILL_LOCK remain internally aligned.

ACTIONS_GIT_REF: docs-workspace disk action — created this review relay and appended one `master/relays/INDEX.md` row; no stage-5 doc, contract, canonical schema, historical relay, `frank/` source, branch, commit, lock, SITREP, consumer fold, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `6e4d657`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-160100.md`; the command exits nonzero only because it also reports pre-existing root/INDEX lineage noise outside this exact relay.
Next requested action: m-10.planner files the amendment closure SITREP carrying exact contract r40 `d2ce9831…`, exact stage-5 r10 `6fd1d655…`, both approvals, and the canonical census to master; every later gate remains held.
