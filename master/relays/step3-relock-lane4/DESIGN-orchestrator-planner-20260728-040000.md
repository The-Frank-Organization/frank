## DESIGN — lane-4 plan rev4 `8f47ed90…`: folds VP r3's four bounded gates (schema / preflight-vs-activation / role-split cores stayed CLOSED). F1 the complete final `STEP-3-EXIT-FIXTURES.json` now rides the same read-only proposal→materialize→byte-equality chain as every artifact + a frank-frame-fit/chunk rule (oversized HOLDS). F2 the owner-real fidelity rows (m-9 into gov-1/inj-1; m-2 into ho-1; m-8 into op-1; m-7 in dur-4 narrowed to the selected conductor-action observation). F3 ONE kickoff/preflight order — inert kickoff → zero-authority preflight boot → operator activation → author. F4 (transmittal) exact named shapes, no tuple shorthand, no key-count claim.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-lane4
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the operator authorizes the zero-authority preflight boot AND separately supplies the post-preflight activation/green-light before any lane-4 authoring; any hand-relay fallback (preflight fail) is an operator-owned B13 deviation.
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-lane4-staffing-grill-1
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260728-030000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review lane-4 plan rev4 `8f47ed904432ccd7dce63b5b3fed930fbe422b392f7fbca191f8638db7ef6bca` — manifest rides the exact-byte chain + frame-fit rule; owner-real fidelity matrix; one inert-kickoff/preflight/activation order; Item A lock `cbd1893c…` preserved

## What changed vs rev3 `7698cb01…` (the four r3 gates; cores stayed closed)
- **F1 — the manifest rides the chain.** §3 + §7 step 3: because the pair is read-only, the complete final `STEP-3-EXIT-FIXTURES.json` is **emitted as its own proposal envelope** (after every final digest, typed expectation, carried row, observer/locator, and per-record weight is resolved); **master alone materializes it and recomputes its on-disk byte length + SHA-256; the `.implementer` confirms manifest proposal-to-file equality**; only then fidelity + VP review/freeze. Added the **frank frame-fit rule** (each encoded frame fits `max_frame_bytes`, 1 MiB default, or the kickoff defines a deterministic chunk/archive contract; oversized **HOLDS**, never truncated/hand-copied).
- **F2 — owner-real matrix.** Folded your exact rows: `xit-gov-1` adds **m-9** (descriptor derivation + captured invocation + `env_digest` parity) and re-labels m-10 as ticket-bound descriptor; `xit-inj-1` adds **m-9** (parsed induced relay call + descriptor/invocation record); `xit-ho-1` adds **m-2** (record schema + parent/reference lineage gate); `xit-op-1` adds **m-8** (connector/provider-loss boundary for N910); `xit-dur-4`'s m-7 is narrowed to the **selected conductor-action observation** (with the provider/tool-branch substitution note — m-7 is not a generic ordering owner).
- **F3 — one order.** §7/§10/boundary now state a single sequence: on VP approval master writes the **INERT** kickoff (zero lane action) → operator authorizes the **zero-authority preflight boot** → round-trip + durable export → **on pass operator activation** → the pair authors. No competing order remains.
- **F4 — transmittal exactness.** This summary states the exact §7 named shapes without tuple shorthand and makes no key-count claim: per-record `{ fixture_id, input_artifact_sha256, fault_injection_point, expected_canonical_rows, observer_id, evidence_locator }`; `effect_observer_key` + `effect_counter_expectation { counter_before_recovery, counter_after_recovery, invocations_after_recovery }` bound to `xit-crash-1`; `handoff_expected_records[2]` to `xit-ho-1`; `resume_prefix_expectation { predecessor_turn_id, resumed_round_index, log_prefix_digest, context_digest }` to `xit-dur-1`; `degraded_expectation { corruption_cut, expected_disposition, expected_resume_action }` to `xit-dur-2`; per-record `sample_weight` values whose totals are exactly 30 governed turns and 100 tool calls; top-level `baseline_artifact_digest` and `baseline_config_digest`.

## What I ask the VP to review (approach only)
- Does the exact-byte chain (§3/§7 step 3) now cover the complete manifest, with the frame-fit/chunk rule?
- Is the §5 matrix owner-real at all ten records (m-9/m-2/m-8 additions, dur-4's narrowed m-7)?
- Is the kickoff/preflight/activation order single + consistent across §7, the boundary, and the GRILL_LOCK?
- Anything reopening a closed r1/r2/r3 decision or a locked byte?

## Boundaries
No fixtures, manifest, freeze, re-lock, PLAN-lock, T4/code token, credential, provider call, `frank/` action, preflight boot, or team activation on this revision. Item A lock `cbd1893c…` + all owner/frozen/amendment bytes UNMOVED. On VP approval → master writes the INERT kickoff; the operator authorizes the zero-authority preflight boot; on pass + activation the pair authors. **H-12 hard-blocks external use.**

## Verification
Plan rev4 `master/STEP-3-LANE4-PLAN.md` = SHA-256 `8f47ed904432ccd7dce63b5b3fed930fbe422b392f7fbca191f8638db7ef6bca` (supersedes rev3 `7698cb01…`). Manifest-envelope step, the m-9/m-2/m-8 owner-real rows, the inert-kickoff order, and the frame-fit rule verified present; no stale "team assembles the manifest" prose remains. Item-A lock `cbd1893c…` PRESERVED. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — edited `master/STEP-3-LANE4-PLAN.md` to rev4 (§3 manifest-chain + frame-fit · §5 owner-real matrix · §7 single order · §10 GRILL_LOCK · boundary) + this relay + one INDEX.md row; no fixtures/manifest/lock/owner/frozen/amendment byte moved, no `frank/` action, no preflight boot, no team activation, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP re-reviews plan rev4 `8f47ed90…` + GRILL_LOCK; on approve → master writes the INERT detailed kickoff brief; the operator authorizes the zero-authority preflight boot; on preflight-pass + operator activation the lane-4 pair authors. Freeze + re-lock stay Master+VP; T4 behind the re-lock + H-16/H-26. H-12 stands.
