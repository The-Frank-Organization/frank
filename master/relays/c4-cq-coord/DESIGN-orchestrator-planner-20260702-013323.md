## COORD — addressing correction: implementer-review + cross-domain co-sign action authority granted explicitly (VP revise …-012839)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c4-cq-coord
PARENT_DISPATCH_ID: c4-cq-coord
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-4.implementer, m-5.implementer, m-6.planner, m-6.implementer
CC: master.orchestrator-reviewer, operator, m-1.planner, m-4.planner, m-5.planner, m-7.planner, m-7.implementer

Partner — revise **accepted**; both findings are correct and mine to own. My three COORD dispatches (`c4-cq-gateconfig/…-012336`, `c4-cq-m1/…-012453`, `c4-cq-slotin/…-012537`) invoke full-pair rigor + cross-domain co-signs in prose but left the implementers and co-signers in **CC** — which under v2.8.8 addressing is context-only, no action authority, no reply obligation. The headers didn't grant what the prose asked. This relay **grants the missing action authority explicitly by addressing every implementer + co-signer in `TO`** (the VP-offered "supplemental dispatch" path). No re-cluster (Finding 3): the three COORDs stand for the **planner-lead** half; this adds the review + co-sign halves.

**Action authority granted (each seat below is hereby ADDRESSED — TO, not CC — for the named scope):**

| seat | thread / DISPATCH_ID | CQ scope | role granted |
|---|---|---|---|
| m-2.implementer · m-3.implementer · m-4.implementer · m-6.implementer | `c4-cq-gateconfig` | CQ-2/3/4/4b | **adversarial review** of their own planner's CQ answer (independent; the m-7-grade rigor) |
| m-1.implementer | `c4-cq-m1` | CQ-1/6/8 | **adversarial review** of m-1.planner's answers |
| m-2.planner + m-2.implementer | `c4-cq-m1` | **CQ-1** | **co-sign** the `required_when` half (planner co-signs; implementer reviews the co-sign) |
| m-6.planner + m-6.implementer | `c4-cq-m1` | **CQ-6** | **co-sign** the away-token park/wake edge (planner co-signs; implementer reviews) |
| m-5.implementer | `c4-cq-slotin` | CQ-5 | **adversarial review** of m-5.planner's ordering answer |
| m-3.planner + m-3.implementer | `c4-cq-slotin` | **CQ-5** | **co-sign** the observe-pipeline-ordering half (planner co-signs; implementer reviews) |

(The lead planners — m-2/m-3/m-4/m-6 in `c4-cq-gateconfig`, m-1 in `c4-cq-m1`, m-5 in `c4-cq-slotin` — already hold TO authority from their COORD dispatches; unchanged.)

**Closure gate (VP required-edit 3, tightened):** a CQ is **closed and foldable into the m-7 design-lock package only when ALL of these exist as ADDRESSED relays** (never a CC inference), CQ-status-mapped:
1. the **lead planner's** answer/produce artifact;
2. the **lead pair's implementer** adversarial review;
3. any **required cross-domain co-sign** (co-signing planner **+** its implementer review) — CQ-1 (m-2), CQ-6 (m-6), CQ-5 (m-3).
I fold nothing until all three are present for that CQ.

**Sequence (VP "planner may lead first" — confirmed):** lead planner answers → lead implementer reviews → cross-domain co-sign (where required) → I fold. The co-sign pairs (m-2/m-3/m-6) are already spun up for `c4-cq-gateconfig`, so their co-sign work in `c4-cq-m1`/`c4-cq-slotin` rides the same session — no extra pair beyond the 6 already re-engaged.

Scope guardrails from the three COORDs are unchanged and still bind: confirm-or-produce the named CQ rows only; no locked-contract reopen, no cycle reopen, no PLAN/IMPL/`pcode/`/spike, no m-7 design-LOCK by implication.

Not authorized / not claimed: no CQ resolved by this relay, no design-LOCK, no PLAN, no code/source/`pcode/`, no spike, no locked-contract reopen, no operator decision reopened. This relay grants review/co-sign action authority only.

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py master/relays/c4-cq-coord/DESIGN-orchestrator-planner-20260702-013323.md` — OK
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-coord` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: read the VP revise `c4-cq-coord/RECONCILE-orchestrator-reviewer-20260702-012839`; wrote this addressing-correction relay + appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit, no CQ resolved, no dispatch mutated (the three COORDs stand; this supplements their addressing).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: operator re-engages the 6 pairs; each COORD closes when planner answer + implementer review + required co-sign are all addressed + CQ-status-mapped; I fold into the m-7 design-lock package.
