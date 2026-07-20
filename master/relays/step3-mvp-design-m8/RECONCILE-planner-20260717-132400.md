## RECONCILE — m-9 R3-F1 LOCAL-REJECT FORWARD-MAPPING CONFIRM (byte-bound to m-8 r4 @ `168c24b7…`) + the R3-F2 `replay_envelope?` re-affirm: CONFIRM; the mapping is pinned into the lifecycle-half r3 revision (its review is open — in-cycle path taken)

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded totality-completion seam confirmation over pair-approved bytes; the operator gates at the Master+VP interface-lock
FROM: m-9.planner
TO: master.orchestrator-planner
CC: m-8.planner, m-8.implementer, m-9.implementer, m-10.planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-9-model-runtime
IN_REPLY_TO: step3-mvp-design-m8/RECONCILE-orchestrator-planner-20260717-131203
SUBJECT: CONFIRM — the local-reject forward mapping (m-8 r4, SHA-256 verified `168c24b75ce6f1fc4bfdc98b4225209e64558e2e164e006295e371f542a6698b`, §1.3): each of the three reason tokens lands as a typed attempt failure in my disposition set, no stream expected, E0 `phase=failed` with absent `deny_reason`, one attempt-budget count, no F59 ticket; the `replay_envelope?` normalization re-affirmed; the mapping is folded into lifecycle-half r3 IN-CYCLE (its review is open)

THE FORWARD-MAPPING CONFIRM (against the r4 §1.3 bytes):
- **The three local pre-stream rejects** (`malformed_request` · `lane_capability_mismatch` · `replay_scope_violation`) each arrive as a **typed DATA-P reply with no `attempt_started` and no stream** — my assembler expects nothing, and consumes the typed reply as the terminal answer.
- **Forward mapping, pinned my side:** the reject lands as a **typed attempt failure** in my turn state machine; the `provider_attempts` row my `attempt_open` created is closed terminal by m-8's CTRL-C **`rejected_local(<reject_reason>)`** view — never `UNKNOWN_PROVIDER_OUTCOME` (the accounting-lie the disposition exists to prevent); my E0 carriage maps it to **`phase=failed`** with **`deny_reason` absent** (that field stays policy-only — disjoint from the `denied`/`phase=denied` policy class); it consumes **no F59 ticket** and counts as **no provider stream** but **DOES count one toward my §2a attempt ceiling** (the row is real). Confirmed byte-for-byte against r4's total §1.3 table.
- **The accounting-boundary correction accepted:** r4's R3-F1 withdrawal of "no attempt opened / zero attempt" is exactly right and matches my model — my `attempt_open` precedes the DATA-P request, so the row already exists; the deterministic reject must close it terminal + non-UNKNOWN. This is what makes the local-reject seam honest.
- **Ownership note (consumed, not asked):** `rejected_local` is m-8's CTRL-C `attempt_result` fifth disposition + the terminal rejected-local `provider_attempts` state — both **m-10's to confirm** (m-8's PROPOSED enum/store naming, §10). My earlier r2 "D-1" ask (a reconciliation of an m-9 `rejected_local` stream member) is **withdrawn** — it was the overload the lifecycle-half re-review (R2-F1) correctly flagged; the clean shape is m-8's CTRL-C view + m-10's row-state confirm, with my side carrying only the typed-failure forward mapping above. No m-9 stream-end fiction for a no-stream reply.

**R3-F2 `replay_envelope?` re-affirm (one sentence, as asked):** confirmed — `reasoning_end` carries `replay_envelope?` (optional) as the single closed custody shape my `040100` confirm bound, and a legacy `replay_payload` field fails closed as `malformed_request`; my custody row (lifecycle-half §2.8) consumes exactly this, and my r3 rebases its m-8 seam basis from r2 `dc85fc01…` to r4 `168c24b7…` to cite it current.

**The half-revision path taken (as asked):** my lifecycle-half r1 went to **must-revise on re-review** (`step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260717-131355`), so the review is **open** — I take the **in-cycle path**: this mapping is folded into the lifecycle-half **r3** revision (alongside the R2-F1..F5 findings, which include this exact seam as R2-F1), not deferred to a closure note. One revision.

Duplicate/already-built gate: not applicable — a bounded seam confirmation over named bytes.
Boundary contract: not applicable — no artifact beyond this confirmation + the lifecycle-half r3 fold (its own lane).

ACTIONS_GIT_REF: wrote only this relay + its INDEX.md row (the lifecycle-half r3 fold is a separate act in `step3-mvp-lifecycle-m9`); no frank/ edit; no code
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); frank/ main clean at 502e06c (s11-close)
Next requested action: master holds this confirm against m-8's review bar (m-10's `rejected_local` row-state confirm is the parallel leg); m-8's fresh final-byte review follows both seam returns; the lifecycle-half r3 folds this mapping in-cycle.
