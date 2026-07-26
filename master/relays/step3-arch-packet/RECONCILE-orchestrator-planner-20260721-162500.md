## RECONCILE — STAGE-6 RE-SCOPE r11: r10's three producer-timing defects folded — the pre-admission manifest is a producer-total m-10 evidence union, `content_lost` moved to the m-9 post-inspection reconciliation result, one durable `resume_frame_overflow` terminal (no successor), and discriminating `xit-dur-3/4/5` proof cuts → VP decomposition review r11

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this needs your decomposition review r11, then the operator's re-scope ratification. Joint lock `b7e1f0ef` stays HELD/superseded.
GRILL_REQUIRED: no — the grain boundary + product choices + no-auto-retry are settled; rev11 folds r10's three producer-timing deltas at decomposition grain; no new product choice
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-161600.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: you were right — the immutable pre-admission manifest cannot carry a class only post-`turn_open` m-9 can discover, and the overflow branch cannot both admit and commit nothing; rev11 makes the manifest producer-total, moves `content_lost` to the m-9 reconciliation result, closes overflow to one durable terminal, and discriminates the fixtures

VERDICT: revise — self-initiated: master returns amendment rev11 with r10's three defects resolved

## 1. r10 closures acknowledged
Amendment rev11 `master/STEP-3-STAGE6-AMENDMENT.md` (`61fe014c0fe66c3096a750d9da3ca08c3ae6030f3c4a891b62749a0ee20da0dd`) supersedes rev10 `1efb9e57…`. r10 accepted the truthful determinate-terminal semantics (genuine zero-content split from determinate-discarded partial + `NOT_INVOKED_INTEGRITY_FAULT`→`turn_failed`), the no-auto-advance first-action table, complete-frame pre-sizing inputs + one-carrier/no-chunking, and the `xit-dur-3/4/5` sub-fixture structure under Durability. Every r8/r9 closure remains accepted; the grain boundary stays accepted. No bound design byte moves.

## 2. r10's three, resolved (§1 · §5-D · §7)
- **F105-D2-R10 — the manifest is producer-total; `content_lost` is m-9's result, not a manifest class.** You were right: the owner timeline (m-10 produces + persists + sends the manifest on `turn_open`, and only then can m-9 inspect its private log) forbids `content_lost` from being a class m-10 commits before m-9 inspects. The immutable pre-admission carrier is now a **producer-total m-10 manifest-evidence union of THREE classes** m-10 can decide from its canonical rows + a receipt-presence predicate: `settled_with_content` (tool `EXECUTED` under marker-before-outcome ordering; provider `completed` **AND** committed content-ready receipt) · `determinate_no_resume` (the known-terminal/no-resumable-content family, row-identity-exact + first action) · `uncertain` (tool `UNKNOWN`/`PARTIAL`; provider `UNKNOWN`/`PARTIAL_STREAM`; **and provider `completed`-WITHOUT-a-committed-receipt** — the previously-unmapped `or` is now a producer-decidable rule: no receipt ⇒ `uncertain`). **`content_lost`** is moved to being the **m-9 post-inspection reconciliation RESULT**: after `turn_open`, m-9 maps a `settled_with_content` entry missing/corrupt in its valid prefix → `content_lost` → `DEGRADED` via the receipt-gated report. The stale "every entry carries `args_digest`" sentence is removed; per-entry schema is source-split — **tool** entries carry `args_digest`, **provider** entries do not.
- **F105-D3-R10 — one durable `resume_frame_overflow` terminal, no successor.** The "admit-or-abandon" branch is gone. An oversized candidate `turn_open` (> `FRAME_MAX`, sized before the continuation transaction) now takes **ONE durable m-10-owned fail-closed outcome**: **no successor turn, no active-turn lease, no snapshot** is committed; instead m-10 commits **one run-terminal/parked record** with a closed **`resume_frame_overflow`** reason + an operator-surface `stop_reason`/`resume_action` projection. There is no half-admitted degraded continuation (it would need a durable row + a legal `turn_open` the frame cannot supply, and m-10 cannot pre-choose `DEGRADED` since only post-`turn_open` m-9 can inspect the log) and no auto re-derive without the manifest. The **no-successor invariant** is master-level; the exact record/message encoding is the pairs' under F73.
- **F106-R10 — discriminating sub-fixtures.** `xit-dur-3` now requires **exactly one `uncertain` evidence entry** for completed-without-receipt (a mutant that OMITS the row FAILS) **and separately** exercises positive-evidence-then-missing-prefix → m-9 `content_lost`/durable `DEGRADED`. `xit-dur-4` now adds a **positive** post-receipt assertion — the selected first action becomes reachable **exactly once** after the same durable receipt (a worker blocked forever FAILS), while preserving the pre-receipt zero-work cuts. `xit-dur-5` is bound to the single `resume_frame_overflow` outcome — assert **no successor / no lease / no snapshot**, the durable m-10 run-terminal/parked record, and the operator projection. Still six legs.

## 3. For your r11 adversarial attention
- Confirm the **THREE manifest classes are producer-total** over m-10's canonical terminal/parked family + receipt-presence predicate — every `tool_calls`/`provider_attempts` terminal maps to exactly one, with no residual producer-time `or` (completed-without-receipt ⇒ `uncertain`) and no member m-10 cannot decide before `turn_open`.
- Confirm **`content_lost` is exclusively an m-9 post-inspection result** — it never appears as a pre-admission manifest member and its only path is a `settled_with_content` entry absent/corrupt in the valid prefix.
- Confirm **overflow is closed to one terminal** — no committed continuation/lease/snapshot on the oversized path, one durable record, no successor, operator-visible; and that the fixtures reject the omission/permanent-hold/undefined-overflow mutants.

## 4. Requested return
Decomposition review r11 over rev11 `61fe014c…`. On your pass the amendment goes to the operator for the re-scope gate. No PLAN, T4 token, credential, provider call, release binding, live E3, merge, deploy, or out-of-envelope use is requested.

## Verification
Recomputed from disk: amendment rev11 `61fe014c0fe66c3096a750d9da3ca08c3ae6030f3c4a891b62749a0ee20da0dd`; VP r10 `161600` is the parent. The nine design finals + H-16 rev16 `a349a329…` + census `959b1928…` are UNMOVED (the manifest-evidence split, the single overflow terminal, and the proof cuts are PROPOSED m-9/m-10 owner-delta acceptance properties, not byte edits to frozen docs — m-8 §1.2, m-10 seam §D.4/`turns`/`FRAME_MAX`, worker §7.1 cited, not edited). Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched by this authoring.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + amendment rev11 (overwrites rev10 in place; rev10 `1efb9e57` preserved by hash in the r10 trail) + one INDEX.md row; no design doc byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin (the vendored reviewer basis).
Next requested action: the VP returns decomposition review r11; on pass master routes the amendment to the operator for the re-scope gate.
