## DESIGN-REVIEW — APPROVE the exact m-10 r11 L7 owner fold at `9aa9f43f…`

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r12
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — pair approval advances only to the corrective SITREP and Master-routed confirmation/rebase round; Master+VP retain interface-lock authority
GRILL_REQUIRED: no — this bounded owner fold does not reopen the operator-locked topology, threat ceiling, F59, F60, or revocation-first choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260717-024500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-024700.md
SUBJECT: APPROVE exact 9aa9f43f... — the L7 six-field connector_assign fold is owner-real, copy-only, generation-blind, secret-free, and fail-closed; route SITREP plus the combined refresh round, not lock or implementation

DESIGN_REVIEW_VERDICT: approve

m-10.planner — I approve the exact revised design bytes at SHA-256 `9aa9f43f9644930c42b209248cdd53402d6085a30e3a5cfe9b255026afa1825e` for the next gate. The earlier approval at `79fcf74290fb2a7c9ca4d706dbceceeaacd7714a23bd82baa70e824824384453` is void because the bytes changed; this is the fresh uniquely-parented review required by VP F73.

## Review verdict

The two-locus L7 fold is exact and coherent:

- **§B.1 owns one six-field frame:** `connector_assign{run_id, turn_epoch, run_manifest_digest, policy_digest, provider_lane_id, lane_catalog_digest}`. The three additions are durable-state-sourced, byte-identical copies of the frozen §C.1 manifest fields; m-10 does not derive, select, or author them.
- **Encoding is closed at the seam:** both digests are 64 lowercase hex and `provider_lane_id` is the byte-exact m-8 canonical lane ID carried by the manifest. None of the additions is a counter, so §A.2 is unchanged.
- **The READY gate fails closed:** m-8 verifies loaded policy, loaded catalog, and pinned-lane membership; any mismatch withholds `connector_ready`, returns a typed fault, and prevents admission.
- **The boundary stays intact:** no `generation_id`, credential bytes, provider payload, policy judgment, or new m-10 authority rides the frame. The interregnum reveal proof and generation-blind CTRL-C rule remain unchanged.
- **§B.4 has no second drifting wire definition:** canonical replacement step 4 refers to the §B.1 six-field shape and preserves connector-ready-before-lease/admission ordering.

The fold matches m-10's accepted L7 return (`step3-mvp-design-m8/RECONCILE-planner-20260717-020500.md`) and Master's corrective (`step3-mvp-design-m10/DESIGN-orchestrator-planner-20260717-022904.md`). The current m-8 r1 contract at `c5eb7b69644b7991be6c368768041555bc982e9f66225aa6ed457e60f4c462bc` independently consumes the same exact shape, copy-only pins, and READY-withheld mismatch behavior. That m-8 rebase is consistency evidence, not a substitute for its routed consumer re-review/final-byte review.

The previously approved m-10 topology, recovery, fencing, F59, manifest, store, boundary, and deferral mechanisms remain substantively unchanged outside the two routed loci. No new blocker is introduced by L7.

## Gate disposition

APPROVE is byte-bound only to `9aa9f43f9644930c42b209248cdd53402d6085a30e3a5cfe9b255026afa1825e`. It authorizes m-10.planner's corrective SITREP and the already-routed m-3 consumer confirmation on this hash. It does not itself refresh m-9/m-7/m-3 confirmation, approve m-8, or close stage 1.

Master must still route the combined current-hash refresh only after m-7's F70+L1 owner revision reaches fresh pair approval. The live m-7 r7 bytes at `fff04fcf91290016d03e521aa02e2e6c86db776f259e9969ee91939f4d0cf214` received a `MUST-REVISE` verdict in `step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260717-024202.md`: their bootstrap recovery still collapses the distinct broker-recognition/commit cases that m-10's approved matrix preserves. m-10's reciprocal CI confirmation/rebase therefore remains held until repaired m-7 bytes pass a fresh review. The separately surfaced m-3/m-9 `turn_epoch` JSON-number versus §A.2 string mismatch is likewise a consumer/owner-resolution item for the refresh/lock round, not permission to weaken or silently amend m-10's exact counter rule.

Master+VP interface lock, stage-3/5 final closure and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `86e52d1873b8f0063ae23815ec375d86d7fd262917de4e935baeb4fc48d01976`.
- Exact design SHA-256 recomputed: `9aa9f43f9644930c42b209248cdd53402d6085a30e3a5cfe9b255026afa1825e`.
- Ratified MVP amendment SHA-256 recomputed: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Ratified architecture amendment SHA-256 recomputed: `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`.
- Pair-approved m-3 r2 SHA-256 recomputed: `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44`.
- Current m-8 r1 SHA-256 recomputed: `c5eb7b69644b7991be6c368768041555bc982e9f66225aa6ed457e60f4c462bc`.
- Incoming DESIGN relay exact-file lint: OK.
- Exact L7 sweep: §B.1 `:65`; canonical replacement reference `:103`; manifest source fields `:135-140`; m-8 r1 bootstrap/load consumer `:170-171`.

ACTIONS_GIT_REF: docs-workspace action only — created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`
RELAY_LINT: OK — exact-file relay-lint.py target result
Next requested action: m-10.planner files the corrective SITREP plus the routed m-3 confirmation against exact `9aa9f43f9644930c42b209248cdd53402d6085a30e3a5cfe9b255026afa1825e`; Master routes the combined current-hash refresh only after repaired m-7 owner bytes receive fresh pair approval. Do not claim interface lock or advance to PLAN/T4/implementation.
