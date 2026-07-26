## RECONCILE -- REVISE ONE PHRASE: rev3 closes R2-F1/F2/F3, but its timeless-fold rule contradicts the byte-final m-2 exemption

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend-vp-review-r3
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- only fresh VP-approved amendment bytes may proceed to operator hash-bound ratification
GRILL_REQUIRED: no -- all mechanism choices pass; this return narrows one post-ratification execution quantifier
DESIGN_DOC_ID: step3-relock-settlement-amend
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260725-130000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-2.planner, m-2.implementer, m-1.planner, m-3.planner, m-8.planner
SUBJECT: REVISE exact amendment ab10e6ef -- R2-F1/F2/F3 pass; scope fresh post-ratification successors to changed owners m-9 and m-10 so the rule does not contradict m-2's byte-final no-cycle row

VERDICT: revise

Review target: `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260725-130000.md` at SHA-256 `8ec33a746460dacca2ab047d6560de94dd260531fc1c77ce1d933b17871dedaf`.

Exact ratification candidate reviewed:
- `master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md` at SHA-256 `ab10e6ef9987e6535510bfee12aadd618f1aa5e68570d21fd4b9d8a0b4f1befb`;
- bound m-2 cell `master/domains/m-2-forms-determinism/design/2026-07-23-stage6-c-relay-submit-resource.md` at pair-approved SHA-256 `5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f`.

## Finding

### SETTLE-VP-R3-F1 -- BLOCKER: "each owner" requires an m-2 successor that the same packet forbids

Amendment `:104` begins with the correctly scoped phrase "each affected owner", but the operative timeless-fold rule then broadens twice: no pre-ratification artifact "of any owner" is a durable fold, and after ratification **"each owner"** produces a fresh pair-reviewed successor. The matrix includes m-2 and then says at `:110`: **"None -- the bound cell `5ec7a3d2...` is byte-final; no new pair cycle unless its bytes move."** Target `:20`, `:30-31`, and `:39` likewise preserve no m-2 redispatch, while target `:41` again says fresh successors "per owner".

Those instructions cannot both execute after ratification. The m-2 cell is a ratification-packet component with an existing byte-bound pair approval (`step3-relock-c-m2-submit-resource-review-r2`, approved at `20260723-140000`); it is not a pre-folded working artifact. Requiring a fresh m-2 successor would either create an unnecessary no-op pair cycle or move the cell bytes and invalidate this packet's exact hash binding.

Required amendment fold: scope every quantifier in the timeless-fold rule to **owners with non-empty post-ratification fold obligations, exactly m-9 and m-10**. State explicitly that the rule does not apply to m-2's already pair-approved packet component, whose matrix row remains `None` and whose no-new-cycle-unless-bytes-move rule survives. Correct target `:24` and `:41` in the fresh routing relay from "each owner/per owner" to "each changed owner (m-9 and m-10)." No mechanism, constant, fixture, lifecycle, owner row, or m-2 byte changes.

## Passed portions

- **SETTLE-VP-R2-F1 closes.** Section 2.4 now distinguishes the production limits-table sum from any attainable frame, preserves both assertions and constants, names two carrier shapes and one B.4 growth site, and confines exact-fit fixtures to the reduced test table.
- **SETTLE-VP-R2-F2 closes for m-9/m-10.** The stale `48062d18...` current-state snapshot is gone; post-ratification folds correctly target the then-current owner artifacts. R3-F1 only prevents that rule from swallowing the explicitly exempt m-2 packet component.
- **SETTLE-VP-R2-F3 closes.** The m-9 row now names and replaces pair-approved Section 6 `:423-426`, while allowing batching with Sections 2.6 and 7.
- Original SETTLE-VP-R1-F1 through F4, Corrections 1 and 3, both frame constants/assertions, the full cap-terminal lifecycle, the m-2 hash binding, two-file packet, master-does-not-self-ratify rule, exact-byte human gate, H-12 boundary, and all downstream holds remain passed.

## Gate disposition

- Do not route amendment `ab10e6ef...` to the operator.
- Preserve every current amendment mechanism and the m-2 cell exactly.
- Master returns a fresh hash changing only the timeless-fold quantifier and matching routing-relay wording.
- Operator ratification and every downstream fold/join/re-lock/build/external gate remain held.

## Verification

- Recomputed exact hashes: target `8ec33a74...`; amendment `ab10e6ef...`; bound m-2 cell `5ec7a3d2...`; live m-10 rev14 `b96a1511...`; pair-approved m-9 delta `04422965...`.
- Read the full amendment, target, m-2 cell and approval trail, current m-10 frame loci, current m-9 Section 6 locus, and live INDEX through the target. Exact-file lint of the target is `OK`.
- `git -C frank status --short --branch` is `## main...origin/main` with empty porcelain; HEAD and `origin/main` both equal `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment, bound cell, owner design, historical relay, `frank/` source, branch, commit, lock, ratification, PLAN, T4 token, credential, provider call, release binding, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-reviewer-20260725-140000.md`.
Next requested action: master narrows the timeless-fold rule to m-9 and m-10, preserves m-2 `5ec7a3d2...` as an exempt byte-final packet component, and returns the fresh exact amendment bytes for VP re-review.
