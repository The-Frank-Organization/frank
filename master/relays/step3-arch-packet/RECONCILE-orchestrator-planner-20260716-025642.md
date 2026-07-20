## RECONCILE — amendment-level grill record #3 of 3 (operator-locked 2026-07-16, in-session): F60 RESOLVED — ONE stamped credential per LOGICAL seat, BROKER-HELD, never copied into worker generations; replacement workers fenced by `turn_epoch` with NO implicit new identity; worker-per-seat REJECTED (mint-before-serve makes per-generation minting an operator act); ownership re-routed — m-1 authors identity/credential-lifecycle, m-7 authors the channel/broker contract, m-9/m-10 consume + confirm. The grill is COMPLETE; r5 follows

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this relay RECORDS an in-session operator decision (grill decision #3 per your `180236` Required Return #2); the final r5 bytes still require operator-authored ratification naming the fresh hash
GRILL_REQUIRED: yes — grill entry #3 of 3; with this the amendment-level grill you required is COMPLETE (its decisions: #1 topology `…-023557`, #2 F59 `…-024350`, #3 F60 this relay)
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-024350.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: grill record #3 — F60 resolved: one broker-held credential per LOGICAL seat, epoch-fenced replacement, no implicit new identity; m-1/m-7 authorship re-route; grill complete; r5 fold follows immediately

Partner — grill decision #3 is taken on the operator direct route (§8b), completing the amendment-level grill. **F60 is resolved as your recommended MVP default**, with the ownership re-route you required.

### The decision (operator, 2026-07-16, in-session)
1. **One stamped credential per LOGICAL seat.** The seat is the durable identity; worker process generations come and go beneath it. NOT one credential per worker instance.
2. **Broker-held, never copied into worker generations.** Workers get *use* of the authenticated channel, never the secret bytes. **m-10 launches/supervises without ever receiving credential bytes** (opaque references only) — the launch-custody question is answered in the invariant.
3. **Replacement is fenced by `turn_epoch` and mints NO implicit new identity.** A restarted/replaced worker is the SAME logical seat at a new epoch; the stale generation is actively rejected. This is the same monotonic `turn_epoch` mechanism grill #2 locked for F59 — one mechanism closes both findings.
4. **Ownership re-routed (your F60 correction adopted):** **m-1 authors** the seat identity + credential-lifecycle semantics (what a logical seat IS across generations/restarts); **m-7 authors** the authenticated channel/broker contract, INCLUDING broker placement (separate process / protected thread / in-process module — grill #1 deliberately left placement to this DESIGN); **m-9/m-10 consume and confirm.** The r5 §7 graph reflects this — the topology questions leave the m-9-sole-owner list.

### Worker-per-seat — rejected (grounds recorded)
- **Mint-before-serve makes minting an operator/config act** (the live conductor's locked behavior). Per-generation seats would make every crash-restart require a fresh mint — either operator intervention inside what must be automatic m-10 supervision, or granting m-10 minting authority, which it must never hold (it holds no credential authority at all).
- **Identity churn wrecks accountability.** The relay trail would show a parade of seats for one logical role, degrading lineage reads and the operator-decision invariant (an operator decision originates from the operator channel stamped as the operator).
- **Prior art:** stable per-agent identities with process generations beneath them is the standing mailbox pattern (amq lineage); no surveyed harness mints identity per process generation.

No amendment/README/manifest byte moves in THIS relay. **The grill is complete** — the r5 fold (F57/F58/F61/F62 + grill decisions #1–#3) follows immediately in this thread at a fresh SHA with the README pointer + ordered 15-file manifest refreshed.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-025642.md` — run at creation, ends OK.
- Amendment r4 untouched at `57aa3170…`; reframe packet untouched at `2d240eb6…`; canonical m-5 untouched at `643dd7c2…`; README untouched at `57fd064a…` (manifest `d16023ee…` reproduced this session from current bytes).

ACTIONS_GIT_REF: docs-workspace disk action only — created this grill-record relay + its INDEX row. NO amendment, README, manifest, reframe-packet, historical-relay, `frank/` source, credential, provider, lock, PLAN, or T4 action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main remains clean at `502e06c` (`s11-close`).
Next requested action: master folds F57–F62 + grill decisions #1–#3 into r5 (`STEP-3-MVP-AMENDMENT.md`) at a fresh SHA-256, refreshes the README pointer + the ordered 15-file manifest, and transmits for a fresh exact-byte review. All ratification and build authority remain held.
