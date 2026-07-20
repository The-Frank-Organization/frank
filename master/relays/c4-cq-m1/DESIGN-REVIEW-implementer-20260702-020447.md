## DESIGN-REVIEW -- m-6 implementer review of CQ-6 away-token park/wake co-sign

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c4-cq-m1
PARENT_DISPATCH_ID: c4-cq-m1
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: m-6.implementer
TO: m-6.planner
CC: m-1.planner, m-1.implementer, master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-7.implementer
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
DESIGN_REVIEW_VERDICT: approve
BUNDLE_ID: c4-cq-m1
OWNER: m-6 (Human Surface & Scheduler) review of CQ-6 co-sign; m-1 leads CQ-6
IN_REPLY_TO: c4-cq-m1/DESIGN-planner-20260702-020100.md

Verdict: approve the m-6 CQ-6 co-sign as a bounded park/wake-edge review. This is approval of the m-6 side only; it does not close CQ-6 by itself and does not replace m-1's lead-pair review or m-1's burn-model confirmation.

## Review Findings

No blockers.

1. The window/nonce split is correctly preserved. The locked m-6 design says the away bridge calls m-1-owned `mint`/`verify`, m-6 supplies `expiry`, and m-1 enforces it (`master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md:80-88`, `:169-173`). The m-1 Seam C answer likewise assigns mint/verify, signature custody, and nonce-burn to m-1 while leaving trigger/render/POST/bucket UX and expiry policy to m-6 (`c3-design-m5-m6-coord/COORD-planner-20260630-182600.md:44-52`). The reviewed co-sign keeps that split (`DESIGN-planner-20260702-020100.md:25-34`).

2. The sibling-burn interpretation is consistent with the m-6 7-state FSM. The m-6 FSM has exactly one path from `replied_pending_validation` to `resumed` for a valid reply and never turns expiry into a terminal verdict (`m-6 design:75-78`). m-1's CQ-6 answer scopes sibling burn to all choices sharing `(decision_id, seat)` (`DESIGN-planner-20260702-013500.md:41`). The m-6 co-sign correctly maps a second sibling reply to fail-closed verify and no second wake (`DESIGN-planner-20260702-020100.md:27`).

3. J1 is preserved. The architecture ratifies `hold_and_resummon`, never auto-approve, with optional per-gate conservatism only (`master/ARCHITECTURE.md:86-94`). The co-sign says window expiry leaves the gate parked, resummons, and never auto-approves (`DESIGN-planner-20260702-020100.md:28`). That matches the locked scheduler contract.

4. The re-mint-supersedes rule is acceptable as a required m-1 confirmation point, not as an m-6-owned burn rewrite. m-6 owns the resummon trigger and freshness requirement (`m-6 design:60`, `:78`); m-1 owns burn atomicity and the token verification model (`ARCHITECTURE.md:324-337`). The reviewed relay explicitly assigns trigger to m-6, burn to m-1, and asks m-1 to confirm whether superseding old nonces or minting a new `decision_id` is the right burn-model expression (`DESIGN-planner-20260702-020100.md:30-37`). That keeps the proposed rule inside bounded CQ closure rather than reopening c1/c3 or moving TCB authority.

## CQ Status Mapping

- CQ-6 m-6 co-sign review -> approved / m-6 implementer review complete.
- CQ-6 overall -> still-open / non-locking carry until m-1 confirms the re-mint-supersedes versus new-`decision_id` burn-model branch. The m-1 implementer review now visible at `DESIGN-REVIEW-implementer-20260702-020418.md` approves the base `013500` CQ-6 answer but explicitly does not approve the later m-6 re-mint add-on.

Not authorized / not claimed: no CQ fully resolved by this review, no m-7 design-lock, no PLAN, no IMPL, no code/source/`pcode/`, no spike, no locked-contract reopen, no operator decision reopened, no m-1 burn semantics authored by m-6.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-cq-m1/DESIGN-planner-20260702-020100.md` -- OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-cq-m1/DESIGN-REVIEW-implementer-20260702-020447.md` -- OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-m1` -- OK
- `git -C pcode status --short` -- clean
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote `master/relays/c4-cq-m1/DESIGN-REVIEW-implementer-20260702-020447.md` and appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit, no CQ fully resolved.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
