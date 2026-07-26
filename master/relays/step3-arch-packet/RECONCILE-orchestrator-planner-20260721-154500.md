## RECONCILE — STAGE-6 RE-SCOPE r9: grain boundary ACCEPTED; the two remaining resume seams fixed — provider settlement is a composite (canonical terminal AND a durable content-ready receipt), and the resume snapshot persists manifest BYTES with a PENDING→RESUMABLE|DEGRADED disposition gated by a post-commit receipt → VP decomposition review r9

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this needs your decomposition review r9, then the operator's re-scope ratification. Joint lock `b7e1f0ef` stays HELD/superseded.
GRILL_REQUIRED: no — the grain boundary + product choices are settled; rev9 fixes the two cross-owner seams; no new product choice
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-153916.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: thank you for accepting the grain boundary — that was the right call and it converged the review; rev9 folds the two master-level seams you named (provider composite-settlement + the durable-bytes snapshot/PENDING-disposition/receipt gate) at decomposition grain, internals still delegated

VERDICT: revise — self-initiated: master returns amendment rev9 with the two cross-owner resume seams made coherent

## 1. r8 closures acknowledged
Amendment rev9 `master/STEP-3-STAGE6-AMENDMENT.md` (`4e2e37506d99f69cdf1d4513734fd705e12901b3255c6bd3ab7ad9e3630b46a7`) supersedes rev8 `9d5e8a34…`. r8 **accepted the grain boundary** (no operator arbitration) and closed F105-D1 (decomposition), the D3 first-action table, and F106 — with F101/F102/F103/F104/handoff/crash-counter/K6/G-2/H-12 all still closed. No bound design byte moves.

## 2. The two seams, resolved (§5-D)
- **F105-D2 tool half — the marker gap.** You were right that content-fsync-before-outcome was one step short (the valid prefix ends at a durable `round_marker`). Fixed: **the content record AND the marker that admits it to the valid prefix fsync-linearize before `record_tool_outcome`** ⇒ `settled_tools ⇒ content in the durable valid prefix`, exactly; the residual (marker not yet durable, m-10 not settled) is `uncertain`.
- **F105-D2 provider half — the producer-order contradiction.** m-8 emits `attempt_result` at the attempt's terminal boundary INDEPENDENTLY of m-9, so raw `provider_attempts=completed` cannot enter `settled` while preserving content-in-prefix. Fixed: **a `settled_providers` entry requires BOTH (a) m-10's canonical provider terminal AND (b) a durable m-9 content-ready receipt bound to `{turn_id, attempt_id, valid-prefix/marker digest}`; m-10 emits it only when both committed.** m-10's canonical terminal is unchanged; the receipt is the new cross-owner conjunction (frame/table = pairs, §6/F73). Plus a **total provider-state partition** (`completed`+content-ready → settled; `denied`/`REJECTED_LOCAL`/`transport_failed`/`CANCELLED` → definite no-content discard; `UNKNOWN_PROVIDER_OUTCOME`/`PARTIAL_STREAM` → uncertain) and **source-`turn_id` on every entry** (frozen tool identity is `UNIQUE(run,turn,tool_call_id)`, so args_digest alone isn't identity-exact across the chain).
- **F105-D3 — snapshot bytes + disposition state machine + receipt gate.** Fixed all three: `resume_snapshot` persists the **canonical manifest BYTES** (a digest cannot reconstruct them; m-10 re-emits from the committed bytes); the disposition is a durable **`PENDING→RESUMABLE|DEGRADED`** status (m-10 cannot truthfully choose at admission — only m-9 can inspect its private log after `turn_open`), **separate** from the immutable snapshot; and a **post-commit disposition receipt** is the no-work gate — m-9 does NO provider attempt / tool effect / conductor verb until it commits, with crash-before-report / after-report-before-commit / after-commit-before-receipt all recovering to the same durable state without permitting work.
- **D1 ownership cleanup (your note):** the exclusive-writer mechanism is m-9-owned if it picks a local OS lock, but a **joint m-10-producer/m-9-consumer seam** if it picks m-10-ordered per-generation segments — worded that way.

## 3. For your r9 adversarial attention
- The **content-ready receipt** is the load-bearing new conjunction — confirm that requiring `{m-10 terminal ∧ m-9 receipt}` for `settled_providers`, with missing-receipt → uncertain/`content_lost`, is total over the provider-state partition and introduces no window where m-10 could emit `settled` before the receipt (the whole point).
- The **marker-before-outcome** ordering — confirm it does not deadlock the frozen per-tool-call `record_tool_outcome` timing (the worker fsyncs {content + the admitting marker} then records the outcome; a crash between → `uncertain`, safe).
- The **PENDING disposition + receipt gate** — confirm the immutable-snapshot/mutable-status split preserves byte-identical `turn_open` re-emission while the disposition legitimately transitions once.

## 4. Requested return
Decomposition review r9 over rev9 `4e2e3750…`. On your pass the amendment goes to the operator for the re-scope gate. No PLAN, T4 token, credential, provider call, release binding, live E3, merge, deploy, or out-of-envelope use is requested.

## Verification
Recomputed from disk: amendment rev9 `4e2e37506d99f69cdf1d4513734fd705e12901b3255c6bd3ab7ad9e3630b46a7`; VP r8 `153916` is the parent. The nine design finals + H-16 rev16 `a349a329…` + census `959b1928…` are UNMOVED (D2/D3 are PROPOSED m-9/m-10 owner deltas + named DESIGN obligations; the content-ready receipt + disposition receipt are new cross-owner seams, not byte edits to frozen docs — m-8 §1.2, m-10 seam §D.4/`turns`, worker §7.1 are cited, not edited). Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched by this authoring.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + amendment rev9 (overwrites rev8 in place; rev8 `9d5e8a34` preserved by hash in the r8 trail) + one INDEX.md row; no design doc byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin (the vendored reviewer basis).
Next requested action: the VP returns decomposition review r9; on pass master routes the amendment to the operator for the re-scope gate.
