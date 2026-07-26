## RECONCILE — STAGE-6 RE-SCOPE r12: r11's three exact-byte consistency defects folded — the trust invariant is now TWO time-scoped properties (settlement-evidence + resume-time evidence-AND-presence), overflow commits to ONE terminal `FAILED`/`resume_frame_overflow` state (no `run-terminal/parked`), and `xit-dur-3/4/5` assert both missing-half orders + an OBSERVED post-receipt action + the exact terminal → VP decomposition review r12

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this needs your decomposition review r12, then the operator's re-scope ratification. Joint lock `b7e1f0ef` stays HELD/superseded.
GRILL_REQUIRED: no — the manifest stages, one-carrier decision, and fail-closed direction are settled; rev12 folds r11's three exact-byte consistency deltas; no new product choice
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-162721.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: you caught the last three exact-byte contradictions — the timeless trust implication forbade its own content_lost branch, run-terminal/parked was two lifecycle meanings, and two fixture mutants still passed; rev12 makes the invariant temporal, pins one terminal state, and discriminates the predicates

VERDICT: revise — self-initiated: master returns amendment rev12 with r11's three defects resolved

## 1. r11 closures acknowledged
Amendment rev12 `master/STEP-3-STAGE6-AMENDMENT.md` (`1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`) supersedes rev11 `61fe014c…`. r11 accepted the producer-total three-class manifest evidence, the completed-without-receipt ⇒ `uncertain` mapping, `content_lost` as the m-9 post-inspection reconciliation result, the source-real per-entry schemas, the no-successor overflow direction, and every r8–r10 closure. The grain boundary stays settled; no operator arbitration. No bound design byte moves.

## 2. r11's three, resolved (§1 · §5-D · §7)
- **F105-D2-R11 — the trust invariant is TWO time-scoped properties, not one timeless implication.** You were right: a timeless `settled_with_content ⇒ content-in-prefix, by construction` reads at resume time as forbidding the exact `content_lost` state the same section requires. rev12 replaces it with the two temporal + conjunctive properties: **(1) at settlement (content-ready commit)** — `settled_with_content` is durable *evidence* that the content record + its admitting durable `round_marker` (tool) / the committed content-ready receipt (provider) **had already fsync-linearized at that earlier point** (nothing more); **(2) at resume trust** — m-9 trusts a `tool_result`/`provider_output` **only** under a matching `settled_with_content` entry **AND** finding the matching content in the **current recovered valid prefix**; evidence-without-current-content yields **`content_lost`** → `DEGRADED`, never trusted or reconstructed. The class name stays; the bytes no longer simultaneously guarantee current presence and specify its absence branch.
- **F105-D3-R11 — one terminal overflow lifecycle state, `run-terminal/parked` removed.** rev12 applies the amendment's own already-safe fail-closed rule (`inability to establish the degraded path → fail-closed run-terminal`) explicitly to overflow: an oversized candidate `turn_open` commits **the run to a single TERMINAL `FAILED` state** carrying the closed `resume_frame_overflow` reason — **no same-run successor or revival**, no lease/snapshot, run GC/restart governed as terminal, operator-visible `stop_reason` + a **manual** `resume_action` (operator-initiated new run only). The terminal-vs-parked choice + the no-same-run-successor rule are master-level; the exact column/message encoding stays pair DESIGN.
- **F106-R11 — the last two passing mutants closed.** `xit-dur-3` now requires exact `uncertain` carriage on **BOTH** missing-half orders: terminal-first/receipt-absent → exactly one `uncertain` for the completed row (omission FAILS), **and receipt-first/terminal-absent → the extant canonically-UNKNOWN/PARTIAL provider row must ALSO appear as exactly one `uncertain`** (omission FAILS here too). `xit-dur-4` no longer accepts mere *reachability*: binding one selected first-action branch, its corresponding durable/wire effect must be **OBSERVED exactly once AFTER the receipt with zero observations before** (a worker exposing a runnable transition but never executing it, or holding forever, FAILS). `xit-dur-5` asserts the exact selected terminal state — run terminal **`FAILED`** with `resume_frame_overflow`, no successor/lease/snapshot/revival, operator manual projection — not the retired `run-terminal/parked` family. Still six legs.

## 3. For your r12 adversarial attention
- Confirm the **two invariant properties compose** — no residual sentence still asserts a timeless `settled ⇒ present`; the settlement-time evidence property and the resume-time evidence-AND-presence property are both present and jointly leave `content_lost` reachable.
- Confirm the **overflow outcome is ONE terminal state** — terminal `FAILED`/`resume_frame_overflow`, no parked/nonterminal alias anywhere, no same-run successor/revival path left to pair implementation.
- Confirm the **fixtures discriminate** — both missing-half orders carry exactly one `uncertain` (omission FAILS either way), the post-receipt action is an OBSERVED exactly-once effect, and `xit-dur-5` asserts the exact terminal state.

## 4. Requested return
Decomposition review r12 over rev12 `1125b0a0…`. On your pass the amendment goes to the operator for the re-scope gate. No PLAN, T4 token, credential, provider call, release binding, live E3, merge, deploy, or out-of-envelope use is requested.

## Verification
Recomputed from disk: amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`; VP r11 `162721` is the parent. The nine design finals + H-16 rev16 `a349a329…` + census `959b1928…` are UNMOVED (the time-scoped trust invariant, the single terminal overflow state, and the discriminating fixtures are PROPOSED m-9/m-10 owner-delta acceptance properties, not byte edits to frozen docs — m-8 §1.2, m-10 seam §D.4/`turns`/`FRAME_MAX`, worker §7.1 cited, not edited). Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched by this authoring.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + amendment rev12 (overwrites rev11 in place; rev11 `61fe014c` preserved by hash in the r11 trail) + one INDEX.md row; no design doc byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin (the vendored reviewer basis).
Next requested action: the VP returns decomposition review r12; on pass master routes the amendment to the operator for the re-scope gate.
