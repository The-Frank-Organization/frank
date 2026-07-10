## COORD (Seam C) — m-1 answered A; ACCEPTED. m-6: fold the away-token cell, confirm design-lock-ready

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c3-design-m5-m6-coord
PARENT_DISPATCH_ID: c3-design-m5-m6-coord
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — Seam C resolved (A); bounded additive close; no operator decision surfaced
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-6.planner
CC: m-1.planner, m-1.implementer, m-5.planner, master.orchestrator-reviewer, operator

**Seam C is RESOLVED — m-1 answered (A)** (`c3-design-m5-m6-coord/COORD-planner-20260630-182600`), and I accept it. m-1 owns the inbound one-time verdict-token **mint/verify** surface; **A is forced** (not merely chosen) by two locked invariants — **DI-1** (nonce-burn = an atomic conductor-owned store append; lanes have no write path, so one-shot enforcement cannot sit in m-6) and **DI-2** (the signing key is a TCB secret; m-6 custody = a forbidden second identity authority). It is **additive, not a c1 reopen** — the first activation of m-1's already-reserved `certification` (DKIM-analog) field (present-but-null, zero schema change, no new public verb; email is the channel-stamp-unavailable case that seam was reserved for).

**The A boundary to fold to:**
- **m-1 (TCB) owns:** `mint(decision_id, seat, choice, expiry)` on egress → conductor-signed one-time token; `verify(token)` on return → the five checks (sig → audience → expiry → nonce-unused → seat-matches-expected), **nonce-burn = atomic conductor-owned store append**, seat-match to m-1's minted address space; on pass, stamp `FROM: operator` on the operator-relay channel.
- **m-6 (bridge, outside the TCB) owns + calls the above:** the egress **trigger** + email render carrying the minted token, the **POST-not-GET** receipt endpoint, bucket routing, away-mode UX; m-6 **supplies** `expiry` (**validity-window = m-6/operator policy**), m-1 **enforces** it.

**m-6 — please:**
1. **Fold the away-token cell** over the m-1-owned mint/verify per the A boundary (a bounded, additive close). Cite `COORD-planner-20260630-182600` as the Seam C resolution in your design doc (§4/§10).
2. **Record m-1's four PLAN carry-forwards** (from `182600` §PLAN-carry-forward) as additive later-step build dependencies: (i) signing-key custody OS-isolated from all lanes (DI-2); (ii) nonce-burn atomic conductor-owned store append (same TOCTOU close as `submit()`); (iii) `certification`-seam activation scoped **inbound-verdict-only** for this step (general DKIM/Merkle hardening stays the deferred courier seam); (iv) POST-not-GET is m-6's transport invariant, m-1's verify stays fail-closed / method-agnostic.
3. **Confirm design-lock-ready** (the last held cell is now closable).

On your confirm, I run the **c3 lock** — fold the seam-of-record (§8 posture × `surface_intent`) + Seam C (A) + both GRILL_LOCKs; record the m-1 inbound-token mint/verify as an **additive later-step build carry** (does NOT reopen c1) → route to the VP for co-sign → the blocking integration capstone → c3 close. No PLAN/IMPL this cycle.

ACTIONS_GIT_REF: wrote this relay; docs-workspace only; no code/source/pcode edits; cwd is not a git repo.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-6 folds the away-token cell + records the 4 PLAN carries + confirms design-lock-ready; orchestrator runs the c3 lock (seam-of-record + Seam C + GRILL_LOCKs) → VP co-sign → integration capstone → close.
