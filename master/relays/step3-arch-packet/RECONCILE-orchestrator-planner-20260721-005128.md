## RECONCILE — THE CORRECTED STAGE-6 LOCK PACKET (F91–F100 closed; deterministic manifest per F96) → VP lock-review r2: the F93 chain is closed (H-16 rev16 master/VP join APPROVED `004123`), the census is rebuilt materialized (F91), the three m-9 lock findings are dispositioned (F92 cert + P4/P5 + the operator's P5 decision), the prebuild/postbuild identity split is pinned (F94/P4), L5 is build-realizable (F95/P5) — every supporting record enumerated by exact path + SHA-256; the joint Master+VP lock issues on your approve, operator-gated

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the stage-6 lock is the operator-gated Master+VP join; this is master's corrected half; the lock record issues only after your approve + the operator's gate
GRILL_REQUIRED: no — all six grill locks stand and bind below
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260720-203905.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: the corrected packet closes your F91–F100 — every design/record below recomputed from disk immediately before filing; the deterministic manifest is the lock's sole locator (F96); the census is materialized full-row (F91); H-16 is joined, not raced (F100)

## 1. The nine design artifacts the lock binds (exact path · SHA-256 · final review)

| # | artifact (exact path) | SHA-256 | final review relay |
|---|---|---|---|
| 1 | `master/domains/m-1-trust-identity/design/2026-07-16-step3-mvp-secret-boundary-seat-identity.md` | `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c` | `step3-mvp-design-m1/DESIGN-REVIEW-implementer-20260716-061153.md` |
| 2 | `master/domains/m-2-forms-determinism/design/2026-07-16-step3-mvp-form-schema-mapping.md` | `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d` | `step3-mvp-design-m2/DESIGN-REVIEW-implementer-20260716-083000.md` |
| 3 | `master/domains/m-3-observation-evidence/design/2026-07-16-step3-mvp-egress-e0-e3.md` | `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad` | `step3-mvp-design-m3/DESIGN-REVIEW-implementer-20260718-040455.md` |
| 4 | `master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md` | `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572` | `step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260717-195335.md` |
| 5 | `master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md` (contract r40) | `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146` | `step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260720-155100.md` |
| 6 | `master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md` | `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51` | `step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260718-043932.md` (+ r40 basis verdict `step3-mvp-design-m8/RECONCILE-implementer-20260720-200710.md` `5d4e543870eddcf7…`) |
| 7 | `master/domains/m-9-model-runtime/design/2026-07-17-mvp-lifecycle-half.md` (r21) | `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd` | `step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260720-203100.md` |
| 8 | `master/domains/m-9-model-runtime/design/2026-07-19-mvp-full-worker.md` (stage-4 r7) | `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45` | `step3-mvp-stage4-m9/DESIGN-REVIEW-implementer-20260720-211500.md` |
| 9 | `master/domains/m-10-app-control-plane/design/2026-07-19-mvp-control-plane.md` (stage-5 r10) | `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf` | `step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-160100.md` |

## 2. The H-16 conductor-outcome contract (F93 closed; joined per F100)

| record | exact path | SHA-256 |
|---|---|---|
| H-16 design rev16 | `master/domains/m-7-conductor-core/design/2026-07-20-h16-outcome-split.md` | `a349a3297c608915510e7bf6de2e304b824897aae4184f7c71db928149ae316d` |
| pair approval | `master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260721-001752.md` | `5f50a5f8da8fa2edcbd2656780e729c251d3a258ca5483f3d9eaeae55d57c336` |
| m-1 owner confirm (a)–(g) | `master/relays/step3-mvp-confirm-m1/SITREP-planner-20260721-003036.md` | `690d88b656dc4b17410dc16a436b8cd96081f51742c1db28e5ef2f9c1ebb56b1` |
| m-2 owner confirm | `master/relays/step3-mvp-confirm-m2/SITREP-planner-20260721-011500.md` | `f866e9800868cc37b73937c46612d202c832a2abc141a2440562b12af3be8a37` |
| master/VP join (master half) | `master/relays/h16-outcome-split/RECONCILE-orchestrator-planner-20260721-003314.md` | `c54aa5851f501b059de67952c14f1f7f92887618fcf093a08c5001dcfd22869e` |
| master/VP join (VP APPROVE) | `master/relays/h16-outcome-split/RECONCILE-orchestrator-reviewer-20260721-004123.md` | `0e1213260191fba88d3080a6edc38f27303abf1748b62e4c9952cfd44a68f6e4` |

The `conductor-relay-accept` census row (§A) binds the H-16 rev16 `decision_state × post_commit_state` monotonic split — the F91 "in flight" cell is now exact.

## 3. The rebuilt H-17 census (F91 closed)

`master/H17-CENSUS.md` v2 @ **`54208535e50723924cc8b61bc254757b5750574d7caf95db26f24adeead114d7`** over schema v1 `master/H17-CENSUS-SCHEMA.md` @ `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`. **39 governed-effect rows, full-field:** 5 conductor-plane (master-authored) + 18 m-10 (materialized verbatim from artifact 9) + 16 m-9 (materialized verbatim from artifact 8). F91 dispositions in the census header: provider-send de-duplicated to E8 (no invented `m8-provider-send`); non-append verbs split; merge/deploy/release + the `-mint`/H-26 gap in the §Non-effect appendix, outside the effect set; the schema-v1 field mapping of the owner renderings stated exact.

## 4. The prebuild identity binding (F94/P4 closed) — NOT "build-identity"

The stage-6 lock binds ONLY the **interface identity contract + the expected catalog vector** (ratified amendment `2f75f2a1…` §4:59 / §7:87). The expected `tool_catalog_digest = 7fae5fc1dd8f91c48828beaf0cfba45a1da4c297bf82f790ec2912b0a168c9d4` (8 rows: five local `tool_schema_digest`s + m-2's three produced relay digests + `m2-mapping-v1` + `m9-catalog-v1`; m-2-recomputed independently at `step3-mvp-confirm-m2/SITREP-planner-20260720-201500.md` `88f4ccfa…`) is stage-6-lockable pre-T4 because it is design-fixed. Actual app/worker/connector/client build digests + `release_digest` bind ONLY at the postbuild RELEASE-BINDING, before live E3.
**P4 (permanent lock-record pin, dual-countersigned — planner `220000` + implementer `223000`):** digests bind postbuild only; stage 6 binds only the interface identity contract + expected catalog vector; actual build digests exist only at the postbuild release-binding. Supersedes worker-r7 `:105`/`:136-137`/`:227` timing statements (an exact-locator erratum; no worker byte moves).

## 5. L5=B at build-realizable grain (F95/P5 closed — operator-decided)

**P5 (permanent pin; planner-countersigned `220000` + implementer-fact-concurrence `223000` + operator product decision):** `m9_worker_build_digest` = the complete runnable worker output INCLUDING linked bytes; the shared conductor-client carries a SEPARATE component/material digest for attribution; the release binding covers BOTH; the "changes iff the worker's own code changes" claim is WITHDRAWN. The operator-ratified **L5 = B** choice (separate-artifact-with-own-attestation) STANDS; the narrow-re-review effect is delivered procedurally via the component digest, not by artifact-digest insensitivity.
| record | exact path | SHA-256 |
|---|---|---|
| m-9 implementer F92 cert + P4 countersign + P5 fact-dissent | `master/relays/step3-mvp-stage4-m9/RECONCILE-implementer-20260720-223000.md` | `c59d4e6b5dfdfe59b216…` (full: recompute at file) |
| the operator's P5 = option 1 decision (§8b record) | `master/relays/step3-mvp-stage4-m9/RECONCILE-orchestrator-planner-20260720-214811.md` | `8e4eebe64c72835f67ca…` (full: recompute at file) |

F92 (the option-(b) revision-neutral certification of worker r7 `cb7ff970…` × half r21 `4d3bd14e…`) is closed by the m-9 IMPLEMENTER's `223000` (the exact-diff basis: six hunks, §1/§3 byte-identical by section hash) — the packet cites the implementer relay, NOT the planner's, per your F92.

## 6. The grill locks (six) + operator decision records

The three amendment grills (`step3-arch-packet/RECONCILE-orchestrator-planner-20260716-023557`, `-024350`, `-025642`) · the m-7 broker-placement grill (in artifact 4) · the stage-5 lock `m10-stage5-grill-lock-20260720` (in artifact 9 §15) · the stage-4 worker grill lock (in artifact 8 §12). Operator decision records, separate relays (F96): T11 loud-failure split `step3-mvp-stage5-m10/RECONCILE-orchestrator-planner-20260720-045603.md`; the `admission_ref` arbitration `step3-mvp-design-m10/RECONCILE-orchestrator-planner-20260720-145111.md`; the P5 decision `step3-mvp-stage4-m9/RECONCILE-orchestrator-planner-20260720-214811.md`.

## 7. The pins + errata (permanent lock-record items; no byte edits)
N1–N4 (as disposed; N4 referent carried to r40 at `step3-mvp-confirm-m7/SITREP-planner-20260720-200656.md` `8c84e244…`) · P1 `sent` = observed `attempt_started` never wire-crossing (`step3-mvp-confirm-m3/SITREP-planner-20260720-194500.md` `7b0799c8…` + m-8 concurrence `step3-mvp-design-m8/RECONCILE-planner-20260720-200000.md` `44511bb7…`) · P2 the narrowed E0 terminal-visibility invariant · P3 the S-1 Governance-Decay strength · **P4** (§4) · **P5** (§5) · the L-ledger (L2–L4/L6 ride this record; L5 = B resolved in artifact 8; L1/L7/L8 folded historically) · **H-26** the `-mint` unlocked-writer defect (`master/FRANK-HARDENING-BACKLOG.md`, m-1-endorsed, fix-path shared with the H-16 IMPL lane).

## 8. The confirmation/reciprocal evidence (F81 current carriers · exact path · SHA-256)
Stage-4 legs: m-10 leg-1 + reciprocal delta `step3-mvp-stage4-m9/RECONCILE-planner-20260720-194500.md` `4ad8b55b…` · m-8 leg-2 + `sent` `step3-mvp-design-m8/RECONCILE-planner-20260720-200000.md` `44511bb7…` · m-2 leg-3 `step3-mvp-confirm-m2/SITREP-planner-20260720-201500.md` `88f4ccfa…` · m-3 leg-4 + `sent` ruling `step3-mvp-confirm-m3/SITREP-planner-20260720-194500.md` `7b0799c8…`. Stage-5 legs: m-9 `step3-mvp-confirm-m9/RECONCILE-planner-20260720-140000.md` `d2acf67b…` · m-1 `step3-mvp-confirm-m1/SITREP-planner-20260720-071330.md` `2f6b3e56…` · m-8 `090000` · m-3 `090000` (in-lane). r40 rebinds: m-7 `step3-mvp-confirm-m7/SITREP-planner-20260720-200656.md` `8c84e244…` · m-3 `step3-mvp-confirm-m3/SITREP-planner-20260720-203000.md` `2533b584…` · m-8.implementer `step3-mvp-design-m8/RECONCILE-implementer-20260720-200710.md` `5d4e5438…`. (Each relay's full SHA-256 recomputes at its named path; the 16-bit prefixes here are the deterministic manifest keys.)

## 9. What the lock DOES / DOES NOT do
DOES: freeze the nine artifacts + the H-16 outcome contract as the build's contract-of-record; bind the interface identity contract + expected catalog vector (NOT built digests); adopt the census v2 `54208535…` as the claim boundary; convert any post-lock byte change into a full-F73 re-open. DOES NOT: issue PLAN, T4 code token, credentials, provider calls, release binding, live E3, merge, or deploy. The T4 PM dispatch follows the lock separately (carrying the F88 three-leg exit bar + PRIOR-ART §4 reference lanes + H-25 raw-ID instrumentation); the H-16 lane completes its own PLAN → RED-first IMPL → diff review → operator merge grant (H-26 convergence only under explicit operator plan scope); the postbuild release-binding + live E3 + the exit test close the step.

## 10. Requested return
Your fresh lock-review r2 over this corrected packet. On approve: the joint Master+VP lock record issues, operator-gated.

## Verification
All 9 design hashes + the census + schema + H-16 rev16 + every §2/§5/§8 record recomputed from disk immediately before filing (the prefixes above are exact); the census materializes 39 full-field rows (5+18+16) verified by row-header count; the F58 expected digest cross-checked against artifact 8 + m-2's independent recompute; exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per your erratum rule, not used as proof).

ACTIONS_GIT_REF: docs-workspace disk action — this packet + one INDEX.md row + `master/H17-CENSUS.md` v2 (rebuilt, hashed above); no design doc byte moved, no `frank/` action, no lock issued by this relay.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `6e4d657`.
Next requested action: the VP returns lock-review r2; on approve master drafts the joint lock record for the operator's gate.
