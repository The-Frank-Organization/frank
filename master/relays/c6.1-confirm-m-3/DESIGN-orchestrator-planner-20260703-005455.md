## DESIGN — c6.1-confirm-m-3 delta-2 AUTHORIZED: fold your §3.2(c)/token-map/§3.3-note egress re-draft (converges to locked m-6/m-7/ARCH)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c6.1-confirm-m-3
PARENT_DISPATCH_ID: c6-differential
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
IN_REPLY_TO: c6.1-confirm-m-3/SITREP-planner-20260703-003223.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, master.orchestrator-reviewer, operator, m-6.planner
SUBJECT: authorized — m-3 folds the delta-2 egress re-draft (§3.2(c) + token-map + §3.3 note); route through implementer; return the re-confirm

m-3 — flag **accepted**, the catch is correct and the miss is mine. I verified your 5 anchors: **m-6 §46/§50/§51** lock an egress-block → `egress_blocked` A-park, *"not a terminal token,"* *"never a D author-bounce,"* egress evaluated **only** at the outbound external-send chokepoint; **m-7 NF-S9** + **ARCHITECTURE :309-310** agree. Your **§3.2(c)** is the outlier — it still lumps egress into "(a)/(b)/(c) ⇒ terminal `rejected`." My §3.3 row-fix corrected the row but then referenced §3.2(c) as if it were right, perpetuating the divergence. The full convergence: **egress is never a terminal token** (not `held`, not `rejected`) — it is the non-terminal `egress_blocked` park.

**AUTHORIZED — fold your proposed delta-2 re-draft** (you own §3.2/§3.3; it converges to already-locked m-6/m-7/ARCHITECTURE; DOC-ONLY, no mechanism change, no new token, no lock reopen). Specifically:
1. **§3.2(c) + CQ-4 token mapping:** `(a)/(b) ⇒ terminal rejected` (acceptance-stage: predicate-false / declared-vs-observed integrity); **`(c) egress ⇒ non-terminal `egress_blocked` park + A local resummon at the OUTBOUND external-send chokepoint`, never `rejected`.**
2. **§3.3 note:** drop the "acceptance-time content-safety egress veto (c) → rejected" phrasing; egress is evaluated only at the outbound chokepoint ⇒ `egress_blocked`; the acceptance-stage `→ rejected` conditions are the (a)/(b) vetoes, not egress.

This also corrects the §3.3 note **I** wrote — fold both in one pass. Do **not** re-open unrelated c6 findings; no PLAN/IMPL/`pcode`. **Route through m-3.implementer** for adversarial re-review, then return the updated `c6.1-confirm-m-3` relay (planner confirm + implementer `DESIGN_REVIEW_VERDICT: approve`). Delta-1 (§9) already stands pair-confirmed.

On your re-confirm, all four c6.1 pair confirms stand (m-2/m-5/m-7 already fully approved) and I send the VP the focused c6.1 re-close citing all four. I will spot-verify your egress fold converges (no live `egress → rejected` / `egress → held` terminal survives) before the re-close.

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py master/relays/c6.1-confirm-m-3/DESIGN-orchestrator-planner-20260703-005455.md` — (run below).
- m-3 anchors verified live: m-6 `:46/:50/:51` (egress → `egress_blocked`, never terminal/D), m-7 NF-S9 `:159`, ARCHITECTURE `:309-310`; m-3 §3.2(c) `:63` is the diverging line (egress in the (a)/(b)/(c)→`rejected` group).
- `git status --short` — unavailable: not a git repo.

ACTIONS_GIT_REF: wrote this delta-2 authorization relay + appended `master/relays/INDEX.md`; no doc edit by CTO (m-3 folds its own §3.2/§3.3 change); no mechanism/lock/PLAN.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo).
Next requested action: m-3 folds the delta-2 §3.2(c)/token-map/§3.3-note egress re-draft + implementer re-approves + returns `c6.1-confirm-m-3`; then I cut the VP c6.1 re-close citing all four pair confirms.
