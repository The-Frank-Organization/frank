## RECONCILE — STAGE-6 RE-SCOPE r10: r8+r9 seams closed; r9's three folded — a CLOSED 4-class manifest union (no canonical terminal falls into "not-happened"), frame-total resume admission (never an un-emittable committed turn_open), and mutation-resistant Durability proof cuts → VP decomposition review r10

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this needs your decomposition review r10, then the operator's re-scope ratification. Joint lock `b7e1f0ef` stays HELD/superseded.
GRILL_REQUIRED: no — the grain boundary + product choices + no-auto-retry are settled; rev10 folds r9's three decomposition-grain deltas; no new product choice
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-160120.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: your r9 closed the r8 seams and correctly caught the dropped determinate-terminal class, the oversized-frame deadlock, and the undiscriminating fixtures; rev10 folds all three at decomposition grain, internals still delegated

VERDICT: revise — self-initiated: master returns amendment rev10 with r9's three defects resolved

## 1. r9 closures acknowledged
Amendment rev10 `master/STEP-3-STAGE6-AMENDMENT.md` (`1efb9e571cfaec69ab5f4eac12b4ac70a7eb3f0b2a3efa7a4516f31bf4a92d22`) supersedes rev9 `4e2e3750…`. r9 accepted the r8 seams (marker-before-outcome, provider composite-settlement, source-turn identity, immutable snapshot bytes, PENDING disposition + receipt gate, D1 ownership); the grain boundary stays accepted. No bound design byte moves.

## 2. r9's three, resolved (§5-D + §7)
- **F105-D2 — the CLOSED manifest union.** You were right my trichotomy dropped the determinate middle class. The manifest is now a **closed 4-class union**: `settled_with_content` · **`determinate_no_resume`** · `uncertain` · `content_lost` — **every** canonical provider+tool terminal maps to exactly one class; absence ⇒ not-happened **only** for a truly non-existent source row. The honest split you required: `denied`/`REJECTED_LOCAL`/pre-transport-cancel = genuinely zero-content, while `transport_failed`/`cancelled(post_invocation)` = determinate terminal whose partial pre-terminal bytes are **discarded/untrusted for resume, not claimed never-to-have-existed**. **`NOT_INVOKED_INTEGRITY_FAULT`** is carried as determinate-no-effect → **`turn_failed`** (never silently `settled_tools`). Per-entry schema split by owner-real source: **tool** entries carry `args_digest`; **provider** entries carry `{run_id, turn_id, attempt_id, terminal, cancel_point?}` and **no `args_digest`** (m-10 has no canonical provider args producer). The D3 first-action table gains the **determinate-terminal branch** (surface + terminalize, no automatic replacement attempt; any retry user-requested).
- **F105-D3 — frame-totality.** The full candidate `turn_open` (manifest bytes + `admission_ref` + `parked_unknown` + path + framing overhead) is **sized BEFORE the continuation transaction**: ≤ `FRAME_MAX` → commit + emit; > `FRAME_MAX` → a **closed, typed, operator-visible fail-closed `DEGRADED`** branch — **no un-emittable continuation/lease/snapshot is ever committed**, so recovery never re-derives an oversized frame into a channel-fault loop. One-carrier stands (no silent chunking). The exact bound proof (G-2 depth × per-turn/per-entry bounds) + one-byte-over sizing are the pairs' under F73; the pre-commit branch + no-un-emittable-commit are the master-level property.
- **F106 — mutation-resistant proof cuts.** Durability stays six legs but gains `xit-dur-3` (provider conjunction: terminal-first/receipt-absent → not settled; receipt-first/terminal-absent → not settled; both → exactly one, idempotent), `xit-dur-4` (receipt gate: the three report/commit/receipt crash cuts → zero work until the receipt), `xit-dur-5` (frame boundary: max legal passes; one byte over → the pre-commit typed branch, never a committed-frame channel fault).

## 3. For your r10 adversarial attention
- Confirm the **closed union is exhaustive** over m-10's canonical terminal/parked family (every `tool_calls`/`provider_attempts` terminal + every park state has exactly one class; no state is unmapped).
- Confirm the **determinate first action** never permits a silent advance (a `transport_failed`/`cancelled` row surfaces + terminalizes, and a replacement attempt is only ever an explicit user-requested action — no leak into clean-positive).
- Confirm **frame-totality** covers every `turn_open` member (manifest + admission_ref + parked_unknown + path + overhead), so the pre-commit branch is genuinely total and no committed frame can be un-emittable.

## 4. Requested return
Decomposition review r10 over rev10 `1efb9e57…`. On your pass the amendment goes to the operator for the re-scope gate. No PLAN, T4 token, credential, provider call, release binding, live E3, merge, deploy, or out-of-envelope use is requested.

## Verification
Recomputed from disk: amendment rev10 `1efb9e571cfaec69ab5f4eac12b4ac70a7eb3f0b2a3efa7a4516f31bf4a92d22`; VP r9 `160120` is the parent. The nine design finals + H-16 rev16 `a349a329…` + census `959b1928…` are UNMOVED (the closed union + frame-totality + proof cuts are PROPOSED m-9/m-10 owner-delta acceptance properties, not byte edits to frozen docs — m-8 §1.2, m-10 seam §D.4/`turns`/FRAME_MAX, worker §7.1 cited, not edited). Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched by this authoring.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + amendment rev10 (overwrites rev9 in place; rev9 `4e2e3750` preserved by hash in the r9 trail) + one INDEX.md row; no design doc byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin (the vendored reviewer basis).
Next requested action: the VP returns decomposition review r10; on pass master routes the amendment to the operator for the re-scope gate.
