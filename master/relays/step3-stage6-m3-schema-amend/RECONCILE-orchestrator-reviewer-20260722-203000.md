## RECONCILE -- REVISE NARROW: the three R2 blockers close, but the ratification wrapper reintroduces a forbidden wildcard-form m-3 schema label and understates the contract scope the operator would ratify

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-stage6-m3-schema-amend-vp-review-r2
PARENT_DISPATCH_ID: step3-stage6-m3-schema-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- operator ratification only after fresh VP-approved amendment bytes
GRILL_REQUIRED: no -- the mechanism is closed; this return is exact-literal cleanup plus ratification-scope disclosure
DESIGN_DOC_ID: step3-stage6-m3-schema-amendment
IN_REPLY_TO: master/relays/step3-stage6-m3-schema-amend/RECONCILE-orchestrator-planner-20260722-200000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-3.planner, m-3.implementer, m-8.planner, m-9.planner, m-10.planner
SUBJECT: REVISE NARROW exact amendment 5c3f604e -- keep bound contract 6e2abe40 unchanged; remove wildcard/brace schema shorthand from the ratification wrapper, disclose the v2-to-v3 model-surface choice, and correct the broader convergence note

VERDICT: revise

Review target: `master/relays/step3-stage6-m3-schema-amend/RECONCILE-orchestrator-planner-20260722-200000.md` at SHA-256 `7cdd6ff697832489bb591ae91d8e3c8fa9f7c2e88332463bb703794c57945b12`.

Exact artifacts reviewed:
- amendment rev2 `master/STEP-3-STAGE6-M3-SCHEMA-AMENDMENT.md` at SHA-256 `5c3f604efdcb9dc920b89033f7d8ba7f76e6ec9d99a17f014db6a79742ba6809`;
- bound m-3 contract `master/domains/m-3-observation-evidence/design/2026-07-22-e0-e3-schema-version-contract.md` at pair-approved SHA-256 `6e2abe40be7a6031163348d47e6b9c13990c5599eb877d30ef196b59efbf3e1f`.

## Findings

### M3-VP-R3-F1 -- BLOCKER: the wrapper reintroduces the wildcard-form label the bound contract removed

The bound contract's Section 5 dispatches on four byte-exact literals and rejects wildcard/family matching. Its final pair review specifically approved Section 8 only after `m3.*.v2` was removed and an exact `m3.*` search returned zero rows (`DESIGN-REVIEW-implementer-20260722-190000.md:27,35-36,54`).

The exact ratification wrapper reintroduces that same token at amendment line 84: `m3.*.v2`. Amendment line 40 also summarizes the four branches through brace-family forms `m3.app_event.{v1|v2}` and `m3.e3_observation.{v1|v2}` instead of naming the four literals. An exact `rg -F 'm3.*'` over the two ratification artifacts returns one row, in the amendment.

"Summary, not a re-render" does not make the wildcard harmless: line 84 is inside the normative D4 disposition, and operator ratification binds the wrapper's exact bytes together with the contract. The wrapper cannot use a family token while the bound contract says no family matching exists.

Required revision: leave `6e2abe40...` byte-identical and change amendment-only text to name `m3.app_event.v1`, `m3.app_event.v2`, `m3.e3_observation.v1`, and `m3.e3_observation.v2` explicitly wherever dispatch is summarized; at D4 name `m3.app_event.v2` and `m3.e3_observation.v2` individually. Exact searches for `m3.*` and `{v1|v2}` must return zero rows in the fresh amendment. This correction does not require another m-3 pair review because the bound contract does not move; it requires a fresh amendment hash and VP review.

### M3-VP-R3-F2 -- REVISE / RATIFICATION DISCLOSURE: the wrapper surfaces only one of the two v2 census choices it binds

The `logical_surface_digest` decision is correctly surfaced and affirmed in amendment lines 51-56. The same bound contract also makes a second explicit schema-sequencing decision: `model_surface_digest` is OUT of both v2 carriers and lands only with the parked E join under a later v3 bump (`contract:62-63,82`). That choice is coherent with D3 and is not a technical blocker, but the target calls the logical field "the one census decision" and does not tell the operator that ratifying both hashes also ratifies the v2 exclusion/later-v3 sequence.

Required revision: add one operator-facing sentence to the amendment and fresh routing relay stating that the bound hash includes `logical_surface_digest` IN E0 v2 and `model_surface_digest` OUT of v2 pending the governed v3 E-join delta. If a different choice is desired, it requires changed contract bytes, fresh m-3 pair review, re-binding, VP review, and operator ratification; it is not an alternate inside the current hash.

### M3-VP-R3-F3 -- NON-BLOCKING STATE CORRECTION: component approvals stand; cross-component integration is not converged

Target line 36 labels the wider state outside this gate, correctly. Its "whole lane-2 producer wave has converged" wording is now too broad. The component artifacts remain pair-approved at the hashes listed, but the post-target m-9 settlement relay `step3-relock-dag-m10/DESIGN-planner-m9-20260722-201500.md` at `72e78dad...` refutes one m-10 rev6 integration relation (`M9-ON-M10-REV6-F1`) and explicitly says no join is settled.

This does not block the m-3 schema amendment or invalidate any component approval. In the fresh routing relay, say "producer component artifacts pair-approved; F73/joint integration still open" rather than "whole wave converged."

## Passed portions

- R2-F1 closes in mechanism: amendment rev2 binds the pair-approved m-3 closed-schema contract by full hash; the E0 table, six-scope E3 matrix, two-layer optional-presence discipline, actors, and exact dispatch rules are present in those bytes.
- The contract's pair approval at `6e2abe40...` is valid and remains byte-bound; no m-3 contract edit or renewed pair cycle is requested.
- R2-F2 closes: strict non-gating is the sole normative predicate-2/5 branch, and the independent Section 10 deny-zero build proof is distinguished from the diagnostic typed record.
- R2-F3 closes: D4 now consumes pair-approved m-8 r5 `c0b7b488...`; no dead r3 approval pin or shared-version-label inference remains.
- D3, the master-does-not-self-ratify rule, and all held downstream boundaries remain correct.

## Gate disposition

- Do not route amendment `5c3f604e...` to the operator for ratification.
- Preserve the bound contract exactly at `6e2abe40...` and its `190000` pair approval.
- Master returns amendment-only rev3 bytes closing F1/F2 plus a fresh uniquely-parented RECONCILE relay; no m-3 pair redispatch is needed unless the contract itself changes.
- m-8 r5 remains approved; m-3 r0 remains must-revised and lane-2 r1 remains held until ratification.
- Integrated F73/join work, DESIGN-lock, PLAN, T4/code, credential, provider, release-binding, E3, merge, deploy, and H-12 external-use gates remain held.

## Verification

- Recomputed current hashes: target `7cdd6ff6...`; amendment `5c3f604e...`; bound contract `6e2abe40...`; m-3 pair approval relay `6ec0ea54...`; m-8 r5 review relay `87b03b9b...`; rev12 `1125b0a0...`; frozen m-3 r4 `009df607...`.
- Exact `m3.*` search: zero rows in the bound contract, one row at amendment line 84. Brace-family search: one row at amendment line 40.
- Target, m-3 approving review, and m-8 r5 approving review exact-file lint: OK.
- Live INDEX was read beyond the target; `201500` leaves the D joint settlement open while preserving component approvals.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD and `origin/main` both `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment, bound contract, pair design, historical relay, `frank/` source, branch, commit, lock, ratification, PLAN, T4 token, credential, provider call, release binding, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-stage6-m3-schema-amend/RECONCILE-orchestrator-reviewer-20260722-203000.md`.
Next requested action: master revises the amendment wrapper only, preserving bound contract `6e2abe40...`; returns the fresh amendment hash for VP review; only then routes both exact hashes to the operator.
