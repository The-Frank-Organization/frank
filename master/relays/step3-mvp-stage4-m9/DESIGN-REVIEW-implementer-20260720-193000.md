## DESIGN-REVIEW — MUST-REVISE m-9 stage-4 full-worker r5 exact bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage4-m9
PARENT_DISPATCH_ID: step3-mvp-stage4-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the owner-real carrier and the operator's GRILL_LOCK stand; the remaining defect is a canonical-census classification correction
GRILL_REQUIRED: yes — satisfied by `GRILL_LOCK_ID: step3-mvp-stage4-m9-worker-grill-1`; no operator choice is reopened
DESIGN_DOC_ID: step3-mvp-design-m9-worker
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: d389eac9ecd347f244de190339ebf6867694acf11b87d44e9b661d65f41b3935
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-stage4-m9/DESIGN-planner-20260720-190000.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-10.implementer, m-8.planner, m-7.planner, m-3.planner, m-2.planner, m-1.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-mvp-stage4-m9/DESIGN-REVIEW-implementer-20260720-193000.md
SUBJECT: MUST-REVISE exact worker r5 d389eac9 — R4-F1/F2 mechanisms close, but E16 still conflates admission with authorization by naming m-10's admission commit as the authorization linearization even though the ref locates-never-grants and the wake disclosure is authorized only by the broker's per-operation seat-capability fence

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I reviewed the complete worker-r5 bytes at SHA-256 `d389eac9ecd347f244de190339ebf6867694acf11b87d44e9b661d65f41b3935`. The owner-real `admission_ref` carrier closes the missing turn→task binding, and the concrete F58 vector/digest closes the catalog assemblability blocker. One H-17 classification defect remains in the new E16 row.

## Finding

### M9-S4-R5-F1 — BLOCKER: E16 calls admission the authorization linearization for a ref that explicitly grants nothing

E16 correctly states the two branch-specific authority sources and enforcement points:

- `wake_relay`: the ref only locates a relay; the worker's own seat USE capability plus the broker's per-operation fence governs the `read`.
- `operator_input`: the task text is inert app-originated content carried on `turn_open`; the m-10 sizing gate controls whether an encodable frame is admitted.

But the row then declares `authorization_linearization_point: m-10's admission commit (the ref becomes fact on the turns row)`. That is the admission/binding linearization, not the authorization for the disclosure. It contradicts the same row's “LOCATES-never-grants” claim, the master route's authority assignment to the seat capability, and H-17's mandatory `admission ≠ authorization` distinction. On the wake branch, authorization becomes fact at the broker's per-operation capability fence when the seat `read` is served; the earlier m-10 admission commit cannot authorize a conductor read because m-10 holds no seat credential and transfers no authority. On the operator-input branch, either name the actual authority/linearization of the content carriage or use the canonical honest null token if authorization is structurally inapplicable. If the branches cannot share one honest cell, split the census family rather than collapsing them.

Required correction: make E16's `authorization_linearization_point` branch-faithful and keep the admission commit at the decision/request-freeze/durable-binding grain where it belongs. Recheck the rest of E16 for the same admission-versus-authorization distinction; no sibling amendment or operator decision is required.

## Accepted r5 substance

- Incoming relay and worker doc hashes reproduce exactly; the incoming relay passes exact-file lint.
- m-10 r40 `d2ce9831…` is owner-real and exact. The REQUIRED two-kind `turn_open.admission_ref`, admission-commit/post-commit-emission split, byte-identical replacement re-carry, and `FRAME_MAX` refusal are realized faithfully in §§2.1/7.1/10.
- The build-integrity overclaim is removed: m-10's F55/F63 serve/release gate is the validator; the worker does not claim it can resolve opaque `manifest_digest` bytes.
- The eight F58 rows are concrete, correctly encode `m2-mapping-v1` versus member absence, and the documented canonical serialization independently recomputes `tool_catalog_digest = 7fae5fc1dd8f91c48828beaf0cfba45a1da4c297bf82f790ec2912b0a168c9d4`. The surface/implementation two-layer check and §11 flag 4 are consistent with m-2.
- Census mechanics pass: 16/16 E rows carry every canonical-v1 label, with exactly two non-effect rationales. E16's mechanism, effect class, ref stability, sizing cut, no-F59 classification, and no-second-identity fixture are accepted apart from the authorization-linearization cell above.
- The other six frozen sibling hashes remain exact. The lifecycle-half r20 amendment is reviewed separately under its own lineage; later pair/reciprocal/stage-6 consumption remains held.

This verdict is byte-bound to `d389eac9ecd347f244de190339ebf6867694acf11b87d44e9b661d65f41b3935`. A corrected design requires a fresh uniquely-parented m-9.implementer DESIGN-REVIEW. No stage-4 SITREP, consumer-confirmation routing, stage-6 lock, PLAN, T4 token, implementation, release binding, E3, merge, or deploy may consume r5 as approved.

## Verification

- Exact incoming DESIGN relay SHA-256: `83a3a8156961ed8ba1bb9f242eeafc72057efa4cd57310312f0a80d3607c9088`.
- Exact reviewed worker-r5 SHA-256: `d389eac9ecd347f244de190339ebf6867694acf11b87d44e9b661d65f41b3935`.
- Exact m-10 r40 contract SHA-256: `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`.
- Exact other owner hashes: m-8 r12 `4b670a79…`; m-7 r11 `9331ea88…`; m-3 r4 `009df607…`; m-2 `83d8e63e…`; m-1 `7c8b09a6…`.
- F58 independent recomputation: canonical sorted eight-row JSON with local mapping-version member absence hashes to `7fae5fc1dd8f91c48828beaf0cfba45a1da4c297bf82f790ec2912b0a168c9d4`.
- H-17 mechanical scan: 16 effect rows, every row contains all 19 body labels plus its `effect_id` heading; exactly two non-effect rationales.
- Full-byte pass: §§0–13, GRILL_LOCK, fold log, r40 R37/R38-F1/R40-F1 owner bytes, m-2 F58 producer/version/absence bytes, and the parallel lifecycle-r20 relay/doc.
- `frank/` HEAD and cleanliness are verified immediately before handoff.

ACTIONS_GIT_REF: docs-workspace action only — created this exact-byte MUST-REVISE relay and appended one `master/relays/INDEX.md` row; no design/source doc, historical relay, frozen contract, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` with empty porcelain at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK — exact-file `relay-lint.py` verification on this relay.
Next requested action: m-9.planner corrects E16's admission/authorization classification in one bounded revision, preserves the accepted r5 mechanisms byte-for-byte, and returns a fresh uniquely-parented full-byte DESIGN relay; later gates remain held.
