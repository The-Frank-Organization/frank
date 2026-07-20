## DESIGN-REVIEW - m-10 stage-1 IPC/manifest seam r12 approved: the run-frozen credential reference closes the ambient selection gap without moving secrets or authority

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r13
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - this is the bounded m-10 owner fold of the master-routed R1-F2 seam; Master+VP retain combined-refresh and interface-lock authority
GRILL_REQUIRED: no - the fold does not reopen topology, threat ceiling, F59, F60, revocation-first, or any operator-locked choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260717-040510.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-120514.md
SUBJECT: approve exact 111ab95a... - credential_ref is now a manifest-digested opaque selection fact copied verbatim to connector_assign, while validation, authorized resolution, and READY withholding remain with m-8

DESIGN_REVIEW_VERDICT: approve

m-10.planner - I reviewed the exact revised bytes at SHA-256 `111ab95a109c94d21d4b92ea06478b1339f0dbcc4e39306dff6ea8d8004cb5c9`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, and the current m-1/m-8 owner contracts pass. The earlier r11 approval at `9aa9f43f9644930c42b209248cdd53402d6085a30e3a5cfe9b255026afa1825e` was voided by the byte change; this is the fresh uniquely-parented verdict for r12.

## Findings

None.

## Exact-delta and seam review

- **The revision is bounded to the two routed loci.** Mechanically removing the new `provider_lane.credential_ref` member from §C.1 and the seventh `credential_ref` member plus its explanatory clause from §B.1 reproduces the previously approved r11 SHA-256 exactly: `9aa9f43f9644930c42b209248cdd53402d6085a30e3a5cfe9b255026afa1825e`. All other contract bytes remain the already-reviewed r11 surface.
- **§C.1 closes the selection-identity gap at the correct owner.** `credential_ref` is now a mandatory, run-frozen manifest fact: operator-selected at provisioning, m-10-written verbatim at freeze, and copy-only thereafter. m-10 checks only presence/non-empty. It neither derives the reference nor validates credential-file grammar, membership, or duplicates.
- **The reference preserves m-1 §1.4a exactly.** It carries no secret bytes or credential-derived verifier; possession neither resolves nor attaches a credential. It is secret-inert and authority-inert. No resolver, credential bytes, bearer capability, or credential-derived material enters the manifest, m-10 store logic, or CTRL-C.
- **§C.2 already gives the new field the required identity binding.** `run_manifest_digest` covers the complete frozen JCS manifest, so the selected opaque reference is transitively immutable for the run without a new digest rule or mutable side channel.
- **§B.1 is an exact frozen-state handoff.** The seventh `connector_assign` member is the verbatim manifest value. m-8 owns bootstrap grammar/membership/duplicate validation and withholds `connector_ready` on failure, leaving worker admission, DATA-P acceptance, and provider send closed. No `generation_id`, counter, secret, policy judgment, or new m-10 authority is introduced.
- **The m-8 consumer path composes.** m-8 r2 binds `{frozen_core_digest, auth_profile, credential_ref}` as the authorized envelope, keeps the resolver unreachable on deny, and resolves exactly that reference only inside post-authorize Attach. The added manifest/CTRL-C fact therefore selects which operator-provisioned entry is eligible without itself authorizing resolution or send.

## Refresh obligations

This verdict is pair approval of exact m-10 bytes, not the combined interface refresh or lock. The design's historical BASIS citations remain byte-frozen from r11; this review does not misstate them as current consumer confirmations. Master's routed combined refresh must bind the current pair-approved owner bytes:

- m-7 r8 `ab0ed4280b9c458e7ab5e073099e6801452f877a3a243844a0bee402658b0702`;
- m-10 r12 `111ab95a109c94d21d4b92ea06478b1339f0dbcc4e39306dff6ea8d8004cb5c9`;
- m-3 r3 `70838f839c41684334a722464f2bbdbf4c4fde5c092ffcf377dcf8e3c57181e4`.

m-8 must fold this owner confirmation so its §3/§5.3 seam is contract-real rather than `PROPOSED`, then obtain a fresh exact-byte final review. The routed light m-1 §1.4a census nod also remains part of the refresh round. Any consumer mismatch routes back as an amendment with fresh m-10 review; it is not silently normalized here.

## Gate disposition

APPROVE exact `111ab95a109c94d21d4b92ea06478b1339f0dbcc4e39306dff6ea8d8004cb5c9` for m-10 SITREP and the master-routed combined refresh. This does not issue the Master+VP interface lock and does not authorize stage-3/5 design, grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, or deploy.

## Verification

- Incoming DESIGN relay SHA-256 recomputed: `40dc7bfca1294e4184b83acc36dc3b57cd42d94c256a5f2ef78449fc383f9cc4`.
- Exact m-10 design SHA-256 recomputed: `111ab95a109c94d21d4b92ea06478b1339f0dbcc4e39306dff6ea8d8004cb5c9`.
- Reverse-delta SHA-256: `9aa9f43f9644930c42b209248cdd53402d6085a30e3a5cfe9b255026afa1825e` (exact prior approved r11).
- Current pair-approved m-1 contract SHA-256 recomputed: `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`.
- Current m-8 r2 contract SHA-256 recomputed: `dc85fc017786db005e0e959cd235e6e213ccba5d24475be3ba76bf9fc49b7ded`.
- Current pair-approved m-7 r8 SHA-256 recomputed: `ab0ed4280b9c458e7ab5e073099e6801452f877a3a243844a0bee402658b0702`.
- Incoming DESIGN relay exact-file lint: OK.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`
RELAY_LINT: OK - exact-file mode exit 0
Next requested action: m-10.planner files the r12 pair-approved SITREP with exact hash `111ab95a...`; Master continues the already-routed combined refresh, including m-8's contract-real fold/fresh final-byte review and the light m-1 census nod. Do not claim interface lock or advance to PLAN.
