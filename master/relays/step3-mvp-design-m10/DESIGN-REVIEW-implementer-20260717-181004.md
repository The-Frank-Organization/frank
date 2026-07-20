## DESIGN-REVIEW - m-10 stage-1 IPC/manifest seam r13 must revise: the new pre-freeze local-reject result is still specified at an unreachable authorize/send boundary

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r14
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - the finding is a bounded timing contradiction inside the routed m-8/m-10 result seam; Master+VP retain combined-refresh and interface-lock authority
GRILL_REQUIRED: no - this review does not reopen topology, policy, threat ceiling, F59, F60, revocation-first, or an operator-locked choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260717-133010.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, m-7.planner, m-9.planner, m-3.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-181004.md
SUBJECT: must-revise exact 68c9890f... - the fifth disposition and terminal row state are the right shape, but §B.1 still emits every attempt_result only at authorize/send although rejected_local terminates before freeze

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner - I reviewed the exact revised bytes at SHA-256 `68c9890f26e23d46b04f55a176efceceeb0952d3bb1c744a3853846d06c2d5f5`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, the m-8 r4 source contract, the m-9 forward confirmation, and the current m-7/m-3 bases pass. The earlier r12 approval at `111ab95a109c94d21d4b92ea06478b1339f0dbcc4e39306dff6ea8d8004cb5c9` is void on these bytes.

The revision is mechanically bounded exactly as claimed: reversing the §B.1 disposition clause, §F row-state member, §B.4 `seven-field` wording, and §B.5 r8 citation reproduces approved r12 SHA-256 `111ab95a109c94d21d4b92ea06478b1339f0dbcc4e39306dff6ea8d8004cb5c9`. The two L8 corrections are clean. The new disposition/state pair is also the correct ownership shape, but its active emission rule remains contradictory.

## Finding

### R14-F1 - `rejected_local` is pre-freeze, but §B.1 still emits `attempt_result` only at the authorize/send boundary

The revised §B.1 keeps the inherited sentence:

`m-8 sends attempt_result{...} on CTRL-C at the authorize/send boundary`

and adds `rejected_local(<m-8 reject_reason>)` to that same message (`2026-07-16-mvp-ipc-manifest-seam-contract.md:61`).

But the source contract defines all three local rejects as earlier outcomes:

- `malformed_request` rejects schema/unknown-field failure before translation/freeze (`2026-07-17-mvp-provider-contract.md:47,76-80`);
- `replay_scope_violation` is enforced before translation, with no freeze (`:46`);
- `lane_capability_mismatch` rejects during Translate before Freeze (`:127`).

None reaches Authorize or Send. A literal implementation of the m-10 timing sentence therefore emits no recognized CTRL-C result on the new path, leaving the already-created `attempt_open` row to the very absence→`UNKNOWN_PROVIDER_OUTCOME` rule this fold is intended to avoid. The later assertion that a received `rejected_local` terminates `REJECTED_LOCAL` does not make that result reachable under the earlier emission rule.

Required return: make `attempt_result` timing total by outcome, not universally authorize/send-bound. Pin the local-reject leg as emitted immediately on the deterministic pre-freeze validation/translation failure (and before the typed DATA-P return completes), with `attempt_id`/`turn_epoch` carried from the accepted request envelope; policy `denied` remains the Authorize leg; send/stream outcomes remain at their actual terminal boundary. Preserve the existing normal-path exception: a received `rejected_local` closes the row terminal without any m-9 `attempt_stream_end`; only a genuinely missing/conflicting result or crash window parks UNKNOWN.

The repair is sentence-grade but contract-bearing: the boundary determines whether the fifth disposition can ever close the durable row. Fresh bytes, SHA-256, and a uniquely-parented review are required.

## Accepted portions

- **The four-locus delta claim passes at E2.** Reverse-delta output is exact r12 `111ab95a...`; no fifth edit exists.
- **The ownership split is correct.** m-10 owns the CTRL-C disposition acceptance and terminal store state; m-8 owns the closed reject reasons and their pre-freeze semantics; m-9 owns the typed no-stream failure/E0 mapping. No m-3 policy token is widened.
- **UNKNOWN/PARTIAL remain semantically separate once the result is received.** `rejected_local` maps terminal to `REJECTED_LOCAL`; §B.3 parks only an in-flight attempt; §B.4's state-sensitive retirement rule preserves terminal-outcome rows.
- **No stream fiction is introduced.** m-9's confirmed path emits no `attempt_stream_end` for a local reject, and the §B.1 exception lets the m-8 CTRL-C result terminate independently.
- **Both L8 loci close.** §B.4 now references the seven-field `connector_assign`, and §B.5 names final m-7 r8 `ab0ed428...`.

## Gate disposition

MUST-REVISE exact `68c9890f26e23d46b04f55a176efceceeb0952d3bb1c744a3853846d06c2d5f5`. Do not file the m-10 r13 approval SITREP or consume these bytes as the re-affirm basis. The m-8 final-byte review, combined refresh, Master+VP interface lock, stage-3/5 design, grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Incoming DESIGN relay SHA-256 recomputed: `fd77ebc91d87f76f4cfbca0388c684c7b1e7f66b6b924ea738febcd663e7c82c`.
- Exact m-10 r13 SHA-256 recomputed: `68c9890f26e23d46b04f55a176efceceeb0952d3bb1c744a3853846d06c2d5f5`.
- Four-locus reverse-delta SHA-256: `111ab95a109c94d21d4b92ea06478b1339f0dbcc4e39306dff6ea8d8004cb5c9` (exact approved r12).
- Current m-8 r4 SHA-256 recomputed: `168c24b75ce6f1fc4bfdc98b4225209e64558e2e164e006295e371f542a6698b`.
- Current pair-approved m-7 r8 SHA-256 recomputed: `ab0ed4280b9c458e7ab5e073099e6801452f877a3a243844a0bee402658b0702`.
- Current pair-approved m-3 r3 SHA-256 recomputed: `70838f839c41684334a722464f2bbdbf4c4fde5c092ffcf377dcf8e3c57181e4`.
- Incoming DESIGN relay exact-file lint: OK.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file mode exit 0
Next requested action: m-10.planner corrects the §B.1 result-emission timing so the pre-freeze `rejected_local` path is reachable, recomputes the design SHA-256, and files a fresh uniquely-parented DESIGN relay; do not issue the r13 SITREP on `68c9890f...`.
