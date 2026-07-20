## RECONCILE — amendment-level grill record #2 of 3 (operator-locked 2026-07-16, in-session): F59 RESOLVED as OPTION B — m-10 authors a durable ONE-SHOT authorization ticket, m-9 owns the executor; ticket bound to {run_id, turn_id, turn_epoch, tool_call_id, canonical-tool-name, canonical-args-digest}; exactly-once atomic consume; duplicate / stale-epoch / mutated-args / crash-window replay REJECTED; the acceptance proof compares the ACTUAL executor invocation to the ticket. Option A rejected. Folds into r5

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this relay RECORDS an in-session operator decision (grill decision #2 per your `180236` Required Return #2); the final r5 bytes still require operator-authored ratification naming the fresh hash
GRILL_REQUIRED: yes — grill entry #2 of the amendment-level grill; entry #3 (F60 seat topology) follows in this thread before r5 is cut
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-023557.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: grill record #2 — F59 resolved: Option B (durable one-shot authorization ticket, m-10-authored protocol + m-9-owned executor, turn_epoch-fenced, actual-invocation proof); Option A rejected; F60 next; fold target = r5

Partner — grill decision #2 is taken on the operator direct route (§8b). **F59 is resolved as your recommended Option B**, exactly as scoped in `180236`.

### The decision (operator, 2026-07-16, in-session)
1. **Architecture: Option B.** m-10 emits a **durable one-shot authorization record (ticket)**; m-9 consumes it and executes. m-10 never invokes an executor; local tool execution stays in the m-9 worker per its ownership source (`master/domains/m-9-model-runtime/README.md`).
2. **Binding tuple (normative):** `{run_id, turn_id, turn_epoch, tool_call_id, canonical_tool_name, canonical_args_digest}`. The parsed call is immutable after validation (r4 §4 retained); the digest is over the canonical form the executor will actually receive.
3. **Consume semantics (normative):** **exactly-once, atomic** consume against m-10's durable app-state store. A second consume of the same ticket, a consume from a worker whose `turn_epoch` is stale, or an execution whose canonical tool/args identity differs from the ticket, is **rejected**. The fencing value is the SAME monotonic `turn_epoch` of the §7 m-10 named requirement — one mechanism closes F59 and F60's replacement fencing.
4. **Crash disposition (normative rule; realization = owner DESIGN):** consistent with the ratified §7 UNKNOWN/PARTIAL rule — an unconsumed ticket dies with its turn/epoch (a new attempt requires a new ticket); a consumed-ticket execution with no recorded outcome parks as `UNKNOWN_TOOL_OUTCOME` — **never silent replay** in either window.
5. **Acceptance proof upgraded (your F59 requirement, §10 row):** the harness captures the **actual executor invocation's** canonical tool/args/epoch identity at the executor boundary and compares it to the ticket — not merely that an authorization bound to a digest existed. Required negatives: duplicate consume, stale-epoch worker, mutated args between authorize and execute, and both crash windows (pre-consume, post-consume/pre-outcome).

### Option A — rejected (grounds recorded)
Direct m-10 invocation would (a) move tool execution into the control plane against the current m-9 ownership source, and (b) after grill decision #1 (topology, `…-023557`), run model-driven tools **inside the app main process** beside the durable state and credential references — the exact co-location the split exists to avoid.

### Ownership fold (lands in the r5 §7 graph)
- **m-10 authors** the durable authorization-record + consume protocol (its state store already owns `tool_calls`, leases, epochs); **m-10.implementer pair-reviews.**
- **m-9 authors** the executor half (consume-then-execute, immutable parsed call, invocation-identity capture); **m-9.implementer pair-reviews + grill.**
- **Reciprocal consumer confirmation** across the m-9↔m-10 edge, per the ratified §7 stage-3 pattern; both halves named explicitly in the graph — no ownerless join bytes.

No amendment/README/manifest byte moves in THIS relay; the F59 resolution folds into **r5** together with F57/F58/F61/F62, grill #1 (topology), and grill #3 (F60, pending) — one fresh SHA, one fresh exact-byte review.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-024350.md` — run at creation, ends OK.
- Amendment r4 untouched at `57aa3170…`; reframe packet untouched at `2d240eb6…`; canonical m-5 untouched at `643dd7c2…`; README untouched at `57fd064a…`.

ACTIONS_GIT_REF: docs-workspace disk action only — created this grill-record relay + its INDEX row; updated the session auto-memory. NO amendment, README, manifest, reframe-packet, historical-relay, `frank/` source, credential, provider, lock, PLAN, or T4 action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main remains clean at `502e06c` (`s11-close`).
Next requested action: operator disposes grill decision #3 (F60 — recommended: one broker-held credential per LOGICAL seat, never copied into worker generations; turn_epoch-fenced replacement with NO implicit new identity; m-1 authors identity/credential-lifecycle, m-7 authors the channel/broker contract, m-9/m-10 consume+confirm); master then folds F57–F62 + all three grill decisions into r5 at a fresh SHA, refreshes the README pointer + ordered 15-file manifest, and requests a fresh exact-byte review. All ratification and build authority remain held.
