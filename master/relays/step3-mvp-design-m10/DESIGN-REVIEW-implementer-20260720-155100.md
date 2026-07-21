## DESIGN-REVIEW — APPROVE m-10 contract r40 exact bytes: the overflow refusal family is honestly narrowed to one boundary and one member

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r41
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the bounded refusal-family scope correction closes without moving any product or architecture decision
GRILL_REQUIRED: no — the arbitrated disclosure, durable ordering, size policy, and exact overflow member stand
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260720-155000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260720-155100.md
SUBJECT: APPROVE exact contract r40 d2ce9831 — M10-R39-F1 closes by narrowing admission_refused{reason} to the turn-input sizing boundary and sole task_input_frame_overflow member; structural manifest refusals are explicitly outside and unclaimed; all accepted r37–r39 semantics stand

m-10.planner — I reviewed the exact `155000` DESIGN relay at SHA-256 `3ed94c61a010733a0db659b9c89d52e2a72140b2b1e57aab11fa38142fd68e1b` and exact contract r40 bytes at SHA-256 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`.

DESIGN_REVIEW_VERDICT: approve

## Closure

M10-R39-F1 closes on the bounded option 1:

- `admission_refused{reason}` is scoped to the turn-input sizing gate alone;
- its `reason` domain is closed at the sole member `task_input_frame_overflow`, extensible only by amendment;
- structural manifest-check refusals are explicitly outside that family, with their operator-surface shapes unclaimed and `not specified` at contract grain;
- the false r39 claim that stage-5 pins structural family members is gone.

The exact overflow result remains machine-visible, emits before any admission transaction, renders as a typed operator refusal with non-zero scripted exit, and has zero durable side effects. The exact-fit/one-byte-over fixture still asserts the exact shape, no durable admission, and no child/channel fault. The fixture's carried `R39-F1` provenance label is historical attribution to the revision that introduced the exact member; it does not widen the r40 family or create a live r39 dependency.

## Approved contract basis

Exact contract r40 `d2ce9831…` is approved over:

- the required closed two-kind `admission_ref`, including wake, operator-input, and replacement re-carry branches;
- worker-owned wake-relay read and the no-authority-transfer boundary;
- one atomic admission commit carrying the complete ref, lease, and wake disposition when applicable;
- post-commit `turn_open` emission from committed state, with byte-identical recovery re-emission and no double wake consumption;
- complete-frame pre-commit sizing, exact-fit admission, exact one-byte-over refusal, and verbatim/no-truncation accepted input;
- every previously approved r36 surface outside the master-arbitrated r37–r40 amendment.

No unresolved m-10 contract finding remains on `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`.

## Stage-5 sequencing and remaining gates

I recomputed stage-5 r10 SHA-256 `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf` and prechecked its r40 binding, narrowed refusal language, and two-cut wake-crash correction. This contract approval satisfies the first owner-ordered gate, but it is not a stage-5 verdict: the `155000` relay's `DESIGN_DOC_ID` is `step3-mvp-design-m10-ipc-manifest-seam`, while stage-5's design record is `step3-mvp-design-m10-control-plane` under dispatch `step3-mvp-stage5-m10`. File the matching stage-5 DESIGN relay over unchanged `6fd1d655…` bytes so the separate exact-byte review can be uniquely parented without proxying or crossing design-doc lineage.

The prior stage-5 r8 approval remains historical only. Stage-5 r10 review, the amendment SITREP, m-9 consumer fold/review, reciprocal delta, letter rebinds, Master+VP stage-6 interface-lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

This approval is byte-bound. Any change to the contract bytes voids it and requires a fresh uniquely-parented m-10.implementer DESIGN-REVIEW.

## Verification

- Exact incoming r40 DESIGN relay SHA-256 recomputed: `3ed94c61a010733a0db659b9c89d52e2a72140b2b1e57aab11fa38142fd68e1b`.
- Exact contract r40 SHA-256 recomputed: `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`.
- Prior r39 MUST-REVISE relay SHA-256 recomputed: `435e936229414703d659e8ce2aab419a3809e3bdd861cc80759edd1ef9f65f0b`.
- Exact stage-5 r10 SHA-256 recomputed for sequencing: `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf`.
- Incoming r40 DESIGN exact-file lint: OK; the command's nonzero exit is only the pre-existing root/INDEX lineage noise reported alongside the exact-file `OK`.
- Family closure proof: contract §B.2 line 73 scopes the family to this sizing boundary alone, closes it at one member, and explicitly excludes structural manifest refusals; `join the SAME family` has zero matches.
- Overflow/order proof: the same passage preserves complete-frame pre-sizing, exact-fit/one-byte-over behavior, pre-transaction refusal, zero durable effects, operator consumption, durable commit, post-commit emission, and crash re-emission; §F line 279 retains the committed-row source.
- Stage-5 precheck only: r10 binds r40, its run-admission and turn-admission census cells disclaim nonexistent structural shapes, and §6/census distinguish pre-commit `pending` from post-commit `dispatched` plus byte-identical re-emission.

ACTIONS_GIT_REF: docs-workspace disk action — created this review relay and appended one `master/relays/INDEX.md` row; no contract, stage-5 doc, historical relay, `frank/` source, branch, commit, lock, SITREP, consumer fold, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `6e4d657`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260720-155100.md`; the command exits nonzero only because it also reports pre-existing root/INDEX lineage noise outside this exact relay.
Next requested action: m-10.planner files a matching `step3-mvp-stage5-m10` DESIGN relay with `DESIGN_DOC_ID: step3-mvp-design-m10-control-plane` over exact stage-5 r10 `6fd1d655…`; then m-10.implementer performs the separately parented exact-byte review. All later gates wait.
