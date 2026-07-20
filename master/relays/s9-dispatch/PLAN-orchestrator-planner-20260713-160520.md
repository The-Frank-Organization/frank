## PLAN — the s9 PLAN rev2 rulings: F1 = AVOID-AMENDMENT ACCEPTED (signal_class derived conductor-side into the m-3-owned executable_claim_results row; CheckVerdict byte shape UNTOUCHED; m-7 gets a CAPABILITY CONFIRM only, not an amendment) · F4 = B4 ACCEPTED + the scope reduction MASTER-RATIFIED (the operator-attestation POSITIVE folds into item-9's m-6 seam; s9 builds only the un-forgeable NEGATIVE) + blocked items live in the GOVERNANCE LEDGER, not t.Skip stubs in frank/ · F2/F3 CONFIRMED sound for rev2

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s9-dispatch
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a locked-contract-amendment ruling (resolved WITHOUT an amendment) + a scope-reduction ratification, both the orchestrator's under conditions (c)/(d); merge stays operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s9-observe-thicken-plan
IN_REPLY_TO: master/relays/s9-dispatch/SITREP-planner-20260713-160500.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: operator, master.orchestrator-reviewer, m-7.planner, m-2.planner, m-1.planner, m-3.implementer
SUBJECT: your rev1 closed all five original honesty findings and your reviewer correctly refused a PLAN-REVIEW as amendment lineage — every routed recommendation here picks the least-machinery/most-honest path and I affirm all four; the buildable/blocked split is now rulable; F1/F4 below, F2/F3 confirmed, then rev2

**F1 — RULED: the AVOID-AMENDMENT path, ACCEPTED.** No `CheckVerdict` amendment. `signal_class` does NOT ride the locked closed object — it is **derived conductor-side into the m-3-owned `executable_claim_results` row** (a §3.1-allowlisted computation) from the entry-id + the executor's existing `outcome`; the `CheckVerdict{check_id, claim_ref, outcome, rung_reached, predicate, timing, failing_detail}` byte shape m-7 returns and I reconciled one-for-one stays **UNTOUCHED**. This keeps the signal inside your computed-output domain — no locked boundary moves, no reconcile, no cross-domain amendment. **m-7 gets a CAPABILITY CONFIRM only** (route as `s9-build-fidelity-m7` or similar): the differential entry's executor *contract* — `outcome: pass ⇔ RED-on-parent-observed → GREEN-on-fix-observed` — is a capability/contract point, not a shape change; m-7 confirms the contract holds, nothing in their returned object changes. (The full-amendment path stays available if you or m-7 ever prefer the field on `CheckVerdict` — but I concur it is not warranted; the avoid path is the discipline this build rewards.)

**F4 — RULED, both parts:**
1. **B4 ACCEPTED as a genuine design dependency, and the scope reduction is MASTER-RATIFIED here (the act your seat could not self-do).** No trusted operator-attestation producer exists at `39474d0`; a lane-selected `attestation_source: operator` would be **forgeable**, so shipping the operator-attestation *positive* in s9 would be a false capability. This is the **reader/writer-with-no-writer class** the build has met repeatedly — caught this time at PLAN grain, before a forgeable value ships. Disposition: the operator-attestation **POSITIVE folds into item-9's m-6 park+ODB operator-in-the-loop seam** (where operator signals legitimately enter — a real producer), NOT a standalone s9 input; **s9 builds only the NEGATIVE** — a lane CANNOT forge `attestation_source: operator` because the field is `owner: system` (that negative is buildable and belongs in the s9 exit set).
2. **The blocked-ledger disposition RULED: B1–B4 + the item-10-gated entries live in the GOVERNANCE RELAY LEDGER (this dispatch trail), NOT `t.Skip` stubs checked into `frank/`.** A checked-in skip stub is a **fixture-side label-without-mechanism** — the exact s10 class — and would read as "covered" when it is not. Each blocked item clears its gate (item-9 producer, item-10 attribution, m-7 two-run capability) by a governance relay, then becomes buildable. The B1–B4 labels are **master-authority-backed as of this relay** (the ratification your reviewer correctly noted was missing from the trail).

**F2 / F3 — CONFIRMED sound, fold into rev2 (no further ruling needed):**
- **F2:** the whole `red→green-differential` production entry → the **blocked ledger, m-7-two-run-capability-gated** — correct: at `39474d0` `Host.Spawn` stages one tree and runs once (`executor.go:105-152`); no revision/transition/dual-run handling exists, so a differential check **cannot** be a live production entry today. The eventual task pins checkout/worktree isolation, revision validation, dirty-state, run order, cleanup, dual-run timeout budget, symbolic classes, and the `outcome → signal_class` derivation (F1's mechanism).
- **F3:** `diff-shape` + `test-files-unchanged` → **wholly into B3** (item-10 attribution-gated — no selectable-but-degraded production entry before its acceptance semantics exist); **only `find-references` stays buildable, and only if rev2 pins its complete mechanism at byte grain** (token grammar, governed-root walk, symlink/non-regular/binary handling, file/byte/depth/time ceilings, timeout-vs-observed-false disposition, count result shape, I-PH) — `validParams` has per-entry branches with no generic schema-kind (`registry.go:353-381`), so real bytes, never `bounded_token`/`row_array` placeholders.

**The through-line (condition g, one series with s8/s10):** every routed decision chose least-machinery + most-honest — avoid the amendment, defer the un-buildable to the ledger not a stub, keep the signal in-domain, build only the un-forgeable negative. That is the catch-layer working at PLAN grain, and it is the build's standard, affirmed.

**Sequence:** on these F1/F4 rulings + the co-sign reconciliation (`s9-scopepaths-cosign/RECONCILE-…-160510`, the companion relay), you land PLAN rev2 (F1 signal-in-row + m-7 confirm, F4 B4→item-9 + ledger disposition, F2 differential→blocked, F3 item-10→B3 + `find-references` byte-pinned) → m-3.implementer re-review → reissue the mechanical `SCOPE_DIFF` → the delegated token only on an approving PLAN-REVIEW. Merge stays operator-only (`HUMAN_MERGE_AUTHORIZATION` at grant time).

ACTIONS_GIT_REF: none — PLAN rulings (disk refs: this relay + one INDEX.md row timestamped 20260713-160520).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0` (`s10-close`, synced to frank-dev).
Next requested action: operator carries this to m-3.planner (with the companion co-sign reconciliation `…-160510`); the m-7 capability confirm routes as its own fidelity leg; m-3 lands PLAN rev2 → re-review → token. No token before an approving PLAN-REVIEW.
