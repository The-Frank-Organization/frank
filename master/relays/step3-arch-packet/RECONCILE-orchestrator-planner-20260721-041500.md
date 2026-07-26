## RECONCILE — STAGE-6 RE-SCOPE r2: F101–F106 answered; the operator grill reversed the sandbox INTO forgone-and-documented, which collapses several blockers; amendment rebuilt hashable + acyclic → VP decomposition review r2

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the operator grill (§3 GRILL_LOCK) is complete; this needs your decomposition review r2, then the operator's re-scope ratification. The joint lock record `b7e1f0ef` stays HELD/superseded.
GRILL_REQUIRED: no — the F106 grill is DONE and folded as the §3 GRILL_LOCK (`step3-stage6-rescope-grill-1`); this relay routes the corrected decomposition
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-033631.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: your r1 was right on all six (F101–F106); the operator grill then forwent the sandbox (documented, H-12 promoted) — resolving F102/F103/F106.1-2 by removal — and I rebuilt the lock as a hashable interface bundle + an acyclic DAG + a real journal source model. Review r2.

VERDICT: revise — self-initiated: master returns the corrected decomposition (amendment rev2) for VP review r2

## 1. Your r1 accepted in full
F101–F106 all land; nothing contested. The corrections are in amendment rev2 `master/STEP-3-STAGE6-AMENDMENT.md` (`29a36285b9c2c5a49deff828490e00c4b99b8a249ed2f6941f2f8821683d1bc6`), which supersedes rev1 `0634b6e4…`.

## 2. The grill outcome changed the shape (§3 GRILL_LOCK `step3-stage6-rescope-grill-1`)
The operator, after the 8-harness sandbox survey (the field mostly forgoes per-tool containment; frontier = OS-sandbox-around-the-agent; nobody mediates per-effect intent from shell semantics — independent confirmation of your F103), **forwent the sandbox in the MVP**: bash stays ambient, the claim is narrowed to exact invocation-context binding (no per-effect destructive/protected holds), the gap is documented, and **H-12 is promoted to a HARD pre-external-use blocker**. That single decision **resolves three of your blockers by removal**:
- **F102** — no fixed sandbox policy is added, so the m-10-hosts-F59 / m-9-executes boundary is UNCHANGED; m-5 stays stood down, m-6 stays out (rev1's move of enforcement into m-9 is withdrawn).
- **F103** — the exact-effect claim is narrowed exactly as you required; the unenforceable fixed-policy holds are dropped.
- **F106.1/.2** — no backend to scope; bash authority = ambient-with-narrowed-claim, grilled and recorded.

## 3. How the remaining blockers are answered in rev2
- **F101 (hashable lock):** §4 makes Tier-HARD a **mechanically-extracted canonical interface bundle** (versioned extraction recipe → one digest), plus whole-file hashes only where the whole file is hard (m-1/m-2/m-3 contracts, H-17). Prose labels/line-ranges are NOT the boundary. The F101 sub-rule is applied: effect-identity / canonicalization / teardown / UNKNOWN-visibility are HARD; only cosmetic wording + tool UX ergonomics are SOFT. `model_surface_digest` locks the recipe/field-set/ownership/carriage, not the per-attempt value.
- **F104 (DAG):** §6 replaces the pair-total order with a per-interface DAG: B `m-8→m-9(carrier)→m-10→m-3`; C `m-10→m-9`; D `m-9⇄m-10` two-sided seam; E `m-2+m-9→aggregator→m-3`; A authored last; **the m-7 study resolves FIRST** (before affected m-10/m-9 finals + the re-lock), H-24 before re-lock if cross-epoch completion survives.
- **F105 (journal source):** §5-D requires a field-level source map + the content/blob store/writer/reader/linearization/atomicity/idempotency/crash-cuts/retention/integrity/redaction contract; hard record-schema+resume-predicate vs soft rendering; m-1 reviews content + at-rest refs; no-second-truth by construction.
- **F106 (grill):** the §3 GRILL_LOCK records D1–D6 (sandbox forgone; bash claim narrowed; exit gate = property legs + overhead; utility = non-gated public dogfood [CRM — informal ask, operator retains all rights — + bivpak] + honestly-labeled agent-as-operator SWE-bench, NO threshold gate to avoid a benchmax incentive; dogfood⊥exit decoupling; effect descriptor kept as context-binding).

## 4. Where master still flags for your adversarial attention (unchanged from r1's carries)
- review-4 (broker) is a PARTIAL simplification — worker-replacement advances the epoch every crash-loop retry, so the F64 fence/linearization stay regardless; only survive-app-main-crash + adoption + cross-epoch-completion is in question. §8 places the study FIRST per your F104.
- review-6 (product proof) is a step-close gate, not an interface-lock artifact; its acceptance contract (property legs + overhead) is predeclared in §7. The utility DEMONSTRATION is deliberately non-gated (public auditable artifacts > a gameable task); the honest consequence — the exit gate proves governance-correct-and-cheap, NOT real-work-useful — is stated plainly in §7.
- I added **m-5.planner + m-6.planner to CC** (you flagged their omission); rev2's position is that they stay OUT because no fixed authority policy is added — please confirm that removal-not-reassignment is the right disposition rather than a silent ownership change.

## 5. Requested return
Decomposition review r2 over amendment rev2 `29a36285…`: (a) is the §4 interface-bundle mechanism a sound hashable lock boundary; (b) is the §6 DAG acyclic + correct against the frozen seams, incl the m-7-first placement + H-24 conditional; (c) is the §5-D journal source model sufficient; (d) does forgoing the sandbox + keeping m-5/m-6 out silently reopen or drop an approved mechanism; (e) does anything move a bound design byte? On your pass, the amendment goes to the operator for the re-scope gate. No PLAN, T4 token, credential, provider call, release binding, live E3, merge, or deploy is requested.

## Verification
Recomputed from disk: amendment rev2 `29a36285b9c2c5a49deff828490e00c4b99b8a249ed2f6941f2f8821683d1bc6`; review basis `b4e79f3b…`; VP r1 `033631` is the parent. The nine design finals + H-16 rev16 + census `959b1928…` are UNMOVED (the amendment adds obligations + re-tiers the lock; it withdraws no approved mechanism and moves no bound byte). Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched by this authoring; the c78da38 vendoring (the reviewer basis) is unchanged.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + amendment rev2 (overwrites rev1 in place; rev1 `0634b6e4` preserved by hash in the r1 trail) + one INDEX.md row; no design doc byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `6e4d657` local / `c78da38` origin (the vendored reviewer basis).
Next requested action: the VP returns decomposition review r2; on pass master routes the amendment to the operator for the re-scope gate.
