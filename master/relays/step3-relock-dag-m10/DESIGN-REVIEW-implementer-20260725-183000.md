## DESIGN-REVIEW — m-10 rev15 does not yet fold Correction 2's ratified no-revival acceptance predicate

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m10-review-r15
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the operator-ratified packet already fixes the missing terminal negatives and acceptance predicate; no product choice is open at this seat
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m10
DESIGN_DOC_SHA256: 4cf52859087aed8f00a7c25585c1fb102c503d2ebb176d1d2123427e56605094
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m10/DESIGN-planner-20260725-180000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-1.planner
BUNDLE_ID: m-10-app-control-plane
SUBJECT: MUST-REVISE exact producer delta rev15 4cf52859 — Corrections 1+2 are relabeled ratified and the joint seam stays pending, but the cap-terminal mechanism/fixture omits Correction 2's ratified no-lease/no-snapshot/no-revival acceptance requirements

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the complete producer delta rev15 at exact SHA-256 `4cf52859087aed8f00a7c25585c1fb102c503d2ebb176d1d2123427e56605094`, the directly addressed request at `e02c954c681699a673a49a4ba5fd09b084dbea783fa664f97ce5c9a8ace7e731`, the operator-ratified settlement amendment packet at `1fa71cb8def2d52690cc513faa7860e0d8c2c2224b33b38ff97808990adc879b`, its §8b ratification/propagation record, and the bound m-2 cell at `5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f`. **MUST-REVISE.** The status transition is sound, but the fresh post-ratification fold does not yet carry the complete ratified Correction-2 terminal predicate.

## M10-DAG-R15-F1 — `FX-M10-CAP` proves no successor generation, but not no lease, snapshot, or revival

The ratified packet is explicit and self-complete. Correction 2 §2.2 lines 36–37 requires the `parked_unknown_capacity_exceeded` terminal to admit:

> no successor generation, same-run continuation, lease, snapshot, or revival

and requires its acceptance predicate to observe in the terminal transaction the full retirement batch, the single FAILED/stop-reason terminal, **no successor/lease/snapshot/revival**, and the durable operator projection of the terminal plus complete queryable parked set.

Rev15 §4 line 72 commits the full retirement transaction and the FAILED/stop-reason terminal, renders the loud surface, and keeps every parked identity queryable, but it does not state the no-continuation/no-lease/no-snapshot/no-revival rule. `FX-M10-CAP` at line 160 likewise asserts only `NO successor generation spawned`. The stronger `NO successor/lease/snapshot` and no-revival language at lines 59/160 belongs to the different `resume_frame_overflow` continuation-sizing terminal; it cannot prove the cap terminal's distinct ratified lifecycle.

This is exactly why the amendment's timeless-fold rule makes rev14 ancestry rather than a durable fold: a fresh successor must fold the ratified packet's complete terminal semantics and acceptance predicate, not only flip the status of the pre-ratification mechanism.

**Required revision:** add the ratified no-revival rule to the operative `parked_unknown_capacity_exceeded` terminal clause: no successor generation, same-run continuation, lease, snapshot, or revival; any next work is an ordinary new run. Extend `FX-M10-CAP` so its terminalizing boundary and multi-row/no-prefix legs observe the full retirement batch, the single FAILED/stop-reason terminal, absence of successor/continuation/lease/snapshot/revival, and the durable operator projection rendering the terminal plus the complete queryable parked set. Do not satisfy this by citing the separate `resume_frame_overflow` fixture. Preserve the packet-bound status, the two compile-time assertions, and the §D join gate.

## Passed portions

- **Correction 1 is folded as ratified and in force.** The Gate-2 claim is honestly relabeled as a fail-closed validator/drift-detector over MVP-unreachable comparison states; comparator and reachable validation bytes remain.
- **Correction 2's main producer mechanism is present.** Run-wide carriage is restored on both frames; the 512 threshold, full-batch/no-prefix rule, closed stop-reason and `resume_action` presence rules, both compile-time assertions, production `≤ FRAME_CONTENT_BOUND` witness, and reduced-table exact-fit/over fixtures remain.
- **The basis is exact.** The live doc binds packet `1fa71cb8…`; the packet and m-2 cell re-hash to the ratification record's exact values.
- **Correction 3 is cited without being folded by m-10.** §5/§9 name the ratified `relay.submit` cell and correctly leave its consumer fold to m-9 §7.
- **The §D seam remains honest.** S-1/S-2/S-4/S-5 stay JOINT-PENDING/NON-NORMATIVE until the required pair folds/approvals and two-sided join; S-4 remains parked on m-9 pending rebase, S-5 remains master/pair-record-bound, and S-3 stays separate.

## Gate effect

No m-10 pair approval, §D two-sided join, lane-2 DAG closure, integrated re-lock, DESIGN lock, PLAN, T4/code, credential/provider action, E3, merge, or deploy is granted by this verdict. The operator ratification itself stands; this finding is only against completeness of m-10's owner fold.

## Verification

- Incoming DESIGN relay: `e02c954c681699a673a49a4ba5fd09b084dbea783fa664f97ce5c9a8ace7e731`.
- Reviewed producer rev15: `4cf52859087aed8f00a7c25585c1fb102c503d2ebb176d1d2123427e56605094`.
- Ratified settlement packet: `1fa71cb8def2d52690cc513faa7860e0d8c2c2224b33b38ff97808990adc879b`.
- Bound m-2 cell: `5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f`.
- Exact search reproduced the cap-terminal gap: the only complete no-successor/no-lease/no-snapshot/no-revival rule in rev15 is attached to `resume_frame_overflow`; `FX-M10-CAP` names only no successor generation.

ACTIONS_GIT_REF: docs-workspace review relay + append-only INDEX row only; no design byte, ratified packet, bound cell, frozen artifact, `frank/` byte, code, lock, PLAN, T4 token, join, credential, provider call, release binding, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: none — clean tree (`frank/` at `c78da3815a34480590071295c1e09bb7d53c10b6`); harness root is not a git repo
RELAY_LINT: OK — exact-file lint with `--relay-root master/relays/step3-relock-dag-m10` exited 0
DISPATCH_ROOT_LINT: OK — `--relay-root master/relays/step3-relock-dag-m10` exited 0
Next requested action: m-10.planner folds Correction 2's complete no-revival mechanism and acceptance predicate into fresh full bytes, preserving the passed ratified-status and joint-pending seam clauses, then returns the successor for exact-byte re-review.
