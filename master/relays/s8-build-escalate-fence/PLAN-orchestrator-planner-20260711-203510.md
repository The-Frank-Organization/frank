## PLAN — FENCE EXPANSION GRANTED (r3): the two stale s5 fixture-oracle rows at the named-seam grain — the r2 TRIPWIRE STATEMENT is SATISFIED on the good side (genesis-at-v5 + live transition; FX-CFG-7 green on the pinned hash; no lock-pin moved) and the reported production behaviors match the locks line by line

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s8-build-escalate-fence-r3-ruling
PARENT_DISPATCH_ID: s8-build-escalate-fence-r3
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the fence ruling is the orchestrator's per the dispatch conditions (operator CC'd); the slice merge stays separately operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
FENCE_EXPANSION_AUTHORIZED: granted — exactly two fixture-oracle rows, named-seam grain
IN_REPLY_TO: master/relays/s8-build-escalate-fence/SITREP-implementer-20260711-203500.md
FROM: master.orchestrator-planner
TO: s8.implementer
CC: operator, master.orchestrator-reviewer, s8.planner, m-3.planner, m-7.planner, m-2.planner
SUBJECT: both rows granted exactly as requested — these reds are the new semantics WORKING (an s3→v6 jump is precisely the skip FX-CFG-10 exists to reject; the pre-Option-B statics are precisely what m-2's reviewed Site 4 removed) and your three no-workaround legs are each lock-faithful; the tripwire acknowledgment, one lightweight confirm, and the continuation path below

**GRANT (named seams, nothing wider):**
1. `frank/test/fixtures/s5_config_change_test.go` — retarget ONLY the old/new registry fixture basis from the historical `s3-fieldspec-v2 → landed` jump to the locked `s7a-fieldspec-v5 → s8-fieldspec-v6` transition; the five real assertion families (accepted-change · digest movement · no-re-genesis · phase-0 chain-walk · stale-form) are PRESERVED, byte-comparable in intent; remove only the now-invalid s3 snapshot/SHA helper + its explanatory comments; the rollback/skip rejection stays asserted against the production validator.
2. `frank/test/fixtures/s5_registry_dormancy_test.go` — ONLY the `surface_intent` predicate expectation moves to Option B (assert BOTH static `required_when` AND `visible_when` absent); every other Block-A/Block-C dormancy assertion stays unchanged. Noted for the record: this edit is the DESIGNED evolution of the dormancy pin, not its weakening — observe stays OFF by default, the flip stays governed, and the remaining assertions still prove it.

**Tripwire statement — SATISFIED, closed for this lane:** genesis derives and verifies the exact v5 predecessor (`1ef6abab…2485`, FX-CFG-7 hard-pinned and green), the live accepted `config_change` advances to v6 (successor SHA `ff7feb0c…a0556` on the record), the owner-supplied forward relation carries the transition, and the reported behaviors — typed `store-not-adopted: run frank -bless` on a legacy store · tokens live in the operator form only post-transition · singular catalog only on the adopted state · live adoption fail-closed as offline-bless-only — match r13 §2.4/§5/§5.1/§5.2 and the m-2 fill-gate story line by line. This is the transition machinery's first live exercise and it held.

**One lightweight confirm (a statement in the T2 report, not a stop):** the working tree shows `internal/config/config.go` modified — the §1.1 version-carrier home, so its presence is expected; confirm in the T2 report that it sits inside the plan's mechanical block (the r2 escalation's in-fence list was illustrative, and the all-in block statement you already produce at commit covers it — just make it explicit once).

**Continuation (exactly as you proposed):** make only the named edits → rerun the focused failures + the full uncached battery (file-captured, sequence-honest) → commit → the T2 report to s8.planner's review per the r2 phantom-seat disposition, with the block statement above. No merge authority granted or implied; the operator gate stands. Third clean stop of the slice — the fence discipline is now demonstrably load-bearing, and the dogfood evaluation should count all three.

ACTIONS_GIT_REF: none — a fence ruling; no `frank/` edit (disk refs: this relay + one INDEX.md row timestamped 20260711-203510; stamped after the replied-to filename per the skew convention).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `691d034`; the build branch state is the pair's to report (last reported: `s8-observe-spine` on `d94dfd4`, in-fence T2 work uncommitted).
Next requested action: operator carries this to s8.implementer; the granted edits + battery + T2 commit/report follow; master next expects the T2 report through s8.planner's review.
