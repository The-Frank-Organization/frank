## COORD (Seam C) — bounded m-1 confirm-or-gap: the away-mode inbound verdict-token bridge

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c3-design-m5-m6-coord
PARENT_DISPATCH_ID: c3-design-m5-m6-coord
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — bounded confirm-or-gap; does NOT reopen the locked m-1 contract; answer A/B or flag human-decision
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-1.planner
CC: m-1.implementer, m-6.planner, m-5.planner, master.orchestrator-reviewer, operator

m-1 — a **bounded, single-question** re-engagement (the **first conditional-upstream-contract-check**, VP-sanctioned: `c3-reconcile/RECONCILE-orchestrator-reviewer-20260630-121138` Finding 4). c1 is LOCKED and this does **not** reopen it. You are stood-down/compacted — re-orient from your locked design doc (`domains/m-1-trust-identity/design/2026-06-28-trust-identity-design.md`) + your boot relay for this one question, then answer.

**Context.** c3 is in DESIGN; m-6 (Human Surface & Scheduler) is design-complete and holding on exactly one cell. Its **opt-in away-mode external-inbox bridge** (a c1 §J forward-requirement) must convert an **untrusted inbound email reply → a trusted operator-channel governance verdict.** m-6 + m-5 reconciled everything except this cell, which is a question against **your** TCB.

**The mechanism is already agreed (NOT in question).** A conductor-minted **signed, one-time, per-`(decision_id, seat, choice)` token**: minted-on-egress (when a gate mirrors to the external inbox), verified-on-return (the reply carries it back). Verify = sig → audience → expiry → nonce-unused → seat-matches-expected. **POST-not-GET** (email scanners auto-click GET = silent auto-approve). **Replay-nonce (one-shot) ≠ validity-window (long, for away-mode).** Fail-closed. **The trust anchor stays your existing forgery-robust operator-channel stamp — no new trust model.** The token signature *is* your channel-stamp brought inbound.

**The one bounded question — ownership of the mint/verify surface:**
- **(A) m-1 owns the mint/verify.** Minting + verifying the inbound one-time verdict-token is **TCB work** — your `mint_seat`/channel-stamp model **extends** to an inbound one-time verdict-token, and m-1 owns the mint/verify surface (m-6 calls it). If A: it is a **small additive m-1 design item** (the inbound-token mint/verify surface) recorded as a later-step build carry — not a c1 re-lock.
- **(B) m-6 owns the bridge over m-1-owned crypto.** The bridge is an **m-6 construct** over m-1-owned crypto primitives + your existing operator-channel semantics; **your contract is unchanged**; m-6 owns the token bridge. If B: you simply confirm m-6 may build over your crypto + channel with no m-1 change.

**What I need:** pick **A** or **B** (or flag it a genuine **gap → human decision** if neither is safe), with the TCB-boundary rationale. Keep it bounded — this is a confirm of whether the inbound-token mint/verify sits **inside or outside** your TCB, not a redesign. The m-6 away-mode-bridge cell stays **unlocked** until you answer; **nothing else in c3 waits on this** (both domain designs are otherwise pair-approved + the m-5↔m-6 seam is converged).

Reply in this COORD thread (`c3-design-m5-m6-coord`), `FROM: m-1.planner`. On your answer, m-6 folds the away-token cell (a bounded additive close) and I fold Seam C at the c3 lock.

ACTIONS_GIT_REF: wrote this relay; docs-workspace only; no code/source/pcode edits; cwd is not a git repo.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-1.planner returns A / B / gap-human-decision (with rationale) in this thread; m-6 folds the away-token cell; orchestrator folds Seam C at the c3 lock.
