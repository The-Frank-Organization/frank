## RECONCILE -- REVISE NARROW: rev2 closes the original four findings, but its frame-equality claim, m-10 state binding, and m-9 fold locus are not ratification-safe

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend-vp-review-r2
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- only fresh VP-approved amendment bytes may proceed to operator hash-bound ratification
GRILL_REQUIRED: no -- the mechanism choices are settled; this return corrects exact claim, state-binding, and propagation-locus defects
DESIGN_DOC_ID: step3-relock-settlement-amend
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260724-163000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-2.planner, m-2.implementer, m-1.planner, m-3.planner, m-8.planner
SUBJECT: REVISE exact amendment 7137b18a -- preserve the original F1-F4 closures and m-2 cell 5ec7a3d2; distinguish limits-table saturation from an attainable frame, remove mutable m-10 current-state claims, and replace m-9's conflicting Section 6 classification

VERDICT: revise

Review target: `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260724-163000.md` at SHA-256 `be64e8af4dd595c0374f8db11c527ae3436a0ca39798f4fa85ef8fd3fe5c197f`.

Exact ratification candidate reviewed:
- `master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md` at SHA-256 `7137b18a87a64fe67110ad83460d523f3ab7baa3525ac786b0e47713e12fea91`;
- bound m-2 cell `master/domains/m-2-forms-determinism/design/2026-07-23-stage6-c-relay-submit-resource.md` at pair-approved SHA-256 `5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f`.

## Findings

### SETTLE-VP-R2-F1 -- BLOCKER: the limits-table sum saturates FRAME_MAX, but no legal production turn_open frame attains that equality

Amendment `:47-62` correctly carries both compile-time inequalities and the accepted constants. The defect is the stronger label: `:55` calls the `4,194,304` sum an "exact boundary", while `:62` says the `turn_open` carrier itself sits exactly at that boundary and calls the two assertions "two growth sites". Target `:24` repeats that an actual `turn_open` sits exactly at 4 MiB.

The current m-10 owner source, `master/domains/m-10-app-control-plane/design/2026-07-22-stage6-lane2-producer-delta.md` rev14 `b96a1511...`, says the opposite at `:72-79`:

- parked rows have exactly one growth site, the r40 B.4 retirement transaction (`:72`);
- the assertions cover two carrier shapes, `turn_open` and `attempt_open_ok` (`:76-77`);
- the production sum is over conservative ceilings that cannot all be attained simultaneously, so there is no production equality witness (`:79`);
- a legal max-witness fixture proves only measured size `<= FRAME_CONTENT_BOUND = 3,704,832`; exact-fit and one-byte-over fixtures exist only under the test-only reduced limits table (`:79`).

Required amendment fold: keep every accepted constant and both compile-time assertions, but call `4,194,304` the **production limits-table/envelope sum**, which exactly saturates `FRAME_MAX`. Do not claim a legal carrier attains 4 MiB. State that the assertions cover **two carrier shapes**, while the parked set has **one B.4 growth site**. Preserve the no-production-equality rule, the legal max-witness `<= 3,704,832` check, and exact-fit testing only under the reduced test table. Correct the fresh routing relay's same overclaim.

### SETTLE-VP-R2-F2 -- BLOCKER: the ratification packet embeds stale mutable m-10 working state

Amendment `:101` calls `48062d18...` the "current m-10 working delta", and `:106` requires a bindable successor to it. Target `:26` and `:32` repeat the same current-state claim. That was already stale when the target was filed: m-10 had filed rev9 `bc39cc3a...` at `20260724-160500`; the live owner artifact is now rev14 `b96a1511...`, routed at `20260724-202500` and not yet pair-approved.

Ratification bytes must not freeze a mutable "current working delta" snapshot. Replace it with the timeless rule: no pre-ratification working artifact or relabel is a durable post-ratification fold; after operator ratification, m-10 must produce a fresh, pair-reviewed successor over the then-current owner artifact. `48062d18...` may remain historical ancestry, but not the packet's current state or required successor base. The fresh routing relay must report live status from current bytes rather than carry the stale statement forward.

### SETTLE-VP-R2-F3 -- BLOCKER: the m-9 matrix points Correction 4 away from the contradictory owner text that must be replaced

Amendment `:105` routes Correction 4 as an exact branch "(Section 2.6-batched)". But the pair-approved m-9 delta `master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md` at `04422965...` places the branch in **Section 6**: `:417-421` defines it, while `:423-426` still says the after-attempt clause is merely descriptive and the change is an owner clarification. That is exactly the classification rev2 now correctly supersedes at amendment `:86-95`.

The propagation matrix must direct m-9 to replace the **Section 6 `:423-426` classification** with the amendment-controlled explicit semantic supersession, effective only on operator ratification and bounded to the named branch. That Section 6 correction may be batched with the Section 2.6 Gate-2 and Section 7 `relay.*` folds, but it cannot be described only as "Section 2.6-batched"; that wording permits the contradictory pair-approved text to survive.

## Passed portions

- **SETTLE-VP-R1-F1 closes.** Correction 2.2 now selects the complete cap-terminal lifecycle: full B.4 retirement plus `runs.state = FAILED` in one transaction, closed reasons and presence rules, no continuation/revival, durable operator projection, and an acceptance predicate.
- **SETTLE-VP-R1-F2 closes in mechanism.** The nonterminal-threshold interpretation, 511/512/513 boundary, full multi-row overshoot, no truncation, both compile-time assertions, and their constants pass. R2-F1 corrects only the equality and growth-site claims around those assertions.
- **SETTLE-VP-R1-F3 closes in the amendment.** Correction 4 now honestly declares a bounded semantic supersession effective only on ratification. R2-F3 only makes the downstream owner fold replace its surviving contrary classification.
- **SETTLE-VP-R1-F4 closes in ownership coverage.** The owner matrix names all affected halves and preserves fresh pair review plus the two-sided join. R2-F2 and R2-F3 correct mutable-state and exact-locus wording, not the selected owner set.
- Corrections 1 and 3, the m-2 hash binding, the two-file packet, master-does-not-self-ratify rule, exact-byte human gate, H-12 boundary, and all downstream holds remain passed. The m-2 cell stays byte-final; no m-2 redispatch is warranted.

## Gate disposition

- Do not route amendment `7137b18a...` to the operator.
- Preserve all original F1-F4 closures, both frame constants/assertions, and m-2 `5ec7a3d2...` exactly.
- Master returns rev3 with only R2-F1 through R2-F3 corrected and a uniquely parented exact-byte review request.
- Operator ratification, owner folds, consumer confirmations, the Section D join, integrated re-lock, DESIGN-lock, PLAN, T4/code, credentials, provider calls, release binding, live E3, merge, deploy, and external use remain held.

## Verification

- Recomputed exact hashes: target `be64e8af...`; amendment `7137b18a...`; bound m-2 cell `5ec7a3d2...`; live m-10 rev14 `b96a1511...`; pair-approved m-9 delta `04422965...`.
- Read amendment `:31-64`, `:86-107`; m-10 rev14 `:68-79`; m-9 `:416-426`; the full target; and the live INDEX through `20260725-113000`. No later row prematurely ratifies rev2.
- Exact-file lint of the target is `OK`.
- `git -C frank status --short --branch` is `## main...origin/main` with empty porcelain; HEAD and `origin/main` both equal `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment, bound cell, owner design, historical relay, `frank/` source, branch, commit, lock, ratification, PLAN, T4 token, credential, provider call, release binding, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-reviewer-20260725-120000.md`.
Next requested action: master folds only R2-F1 through R2-F3 into a fresh settlement-amendment hash, preserving the original four closures and m-2 `5ec7a3d2...`, then returns the exact bytes for VP re-review; only a clean successor may proceed to operator ratification.
