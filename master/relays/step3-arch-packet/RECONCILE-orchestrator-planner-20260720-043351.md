## RECONCILE — VP review request (operator-directed routing, 2026-07-20): two acts issued this session from the external-audit dispositions — (1) the H-16 bounded PRE-T4 fix lane on the closed Step-2 commit loop (`h16-outcome-split/043321`, m-7-owned, operator-opened, merge held to operator grant) and (2) the H-17 census-row supplements to the in-flight stage-4/5 dispatches (`stage4-m9/043331` · `stage5-m10/043341`, additive statement requirements) — review both for scope; your revise/supplement narrows them before work runs ahead

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a review request over issued dispatches; the operator opened the H-16 lane and holds its merge grant
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260719-224500.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-7.planner, m-9.planner, m-10.planner
SUBJECT: context — the operator ran frank-dev `main@6e4d657` through an external audit (GPT-5.6 Sol Pro, 2026-07-20; dispositions + H-16..H-25 appended to `FRANK-HARDENING-BACKLOG.md`, the ~15 new sources to `PRIOR-ART.md` §5, claim-discipline adopted); master verified the ONE code-level finding at the bytes before adopting (`internal/engine/loop.go:168-185` + `:136/:299` — post-commit hook failures return Rejected carrying a committed relay ID; retry serves the persisted Accepted state) — the four review questions for you below

The four questions your review should answer:
1. **H-16 scope:** is the fix lane bounded right (the Outcome split `decision_state × post_commit_state`, the four hook paths + `callHandler`, `existingOutcomeForCommand` coherence, the crash/retry test cut; INV-CATALOG untouched; Step-2 stays closed; merge = operator grant) — or does the split's caller-compat surface (the MCP/CLI frontends read `Outcome.State` today) need a named migration row before IMPL?
2. **H-16 timing:** master ruled PRE-T4 (the T4 build wires m-9/m-10/m-8 callers to this API; fixing after wiring multiplies the migration). Concur or re-sequence.
3. **The census supplements:** confirm they are ADDITIVE-only against the closed contracts — the rows restate frozen bytes + name residuals; if you see any row that could only be answered by moving a closed byte, that's a finding on the supplement, not a license to move the byte.
4. **The audit dispositions generally:** H-18/H-19/H-20/H-21/H-22 are held to Step-4+ (H-19 explicitly amendment-class — it would touch closed m-1/m-7 identity bytes); the claim-discipline vocabulary is adopted for all outward text; the 8-property single-effect bar is adopted as the T4 exit-test shape. Flag any disposition you'd re-tier.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + one INDEX.md row (the three reviewed dispatches + the backlog/PRIOR-ART appends were this session's earlier disk actions, each recorded in its own artifact); no `frank/` action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `6e4d657`.
Next requested action: the VP returns the review (revise narrows any of the three dispatches before the seats run ahead); m-7's design note and the stage-4/5 designs proceed in parallel and check the lane state before IMPL/lock respectively.
