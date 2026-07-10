## COORD — c4 CQ m-1 store-provenance cluster: CQ-1 / CQ-6 / CQ-8 (m-1 lead; m-2 + m-6 co-sign)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c4-cq-m1
PARENT_DISPATCH_ID: c4-cq-coord
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — bounded contract-question closure against locked contracts; operator directed full-pair rigor; VP-reviewed decomposition
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-1.planner
CC: master.orchestrator-reviewer, operator, m-1.implementer, m-2.planner, m-2.implementer, m-6.planner, m-6.implementer, m-7.planner, m-7.implementer
BUNDLE_ID: c4-cq-m1
OWNER: m-1 (Trust & Identity — store/provenance/binding), with m-2 (CQ-1 required_when) + m-6 (CQ-6 token edge) co-sign

**You are re-engaged for a BOUNDED, SCOPED contract-question closure — not a reopened cycle.** m-7 Conductor-Core is DESIGN-COMPLETE (pair-approved r3) and **holding at its CQ gate**. Three of its design-LOCK CQs sit on the **m-1 store/provenance/binding surface you own** — and m-7 has already **proposed the reading** for each; your job is confirm-or-correct against your locked contract. Operator directed **full-pair rigor** (m-1.planner + m-1.implementer); VP reviewed the decomposition (`c4-cq-coord/…-011714` → `…-012056`).

**SCOPE GUARDRAIL (hard):** confirm-or-correct for the THREE named CQ rows **only**. **No** reopening your locked c1 contract; **no** new domain design; **no** PLAN/IMPL/`pcode/`/spike; **no** m-7 design-LOCK by implication. A genuine contradiction with your locked contract → surface it, don't silently resolve.

**Full-pair method:** m-1.planner leads; m-1.implementer independently reviews each confirm-or-correct. The two co-signs (m-2 on CQ-1, m-6 on CQ-6) ride from the `c4-cq-gateconfig` thread where those pairs are engaged.

---

### CQ-1 — phase-split required-set  ·  m-1 (§5 boundary) + m-2 (required_when) co-sign  ·  CONFIRM-OR-CORRECT
m-7's §3-step-2a proposal: a **Step-1 form gate must not demand observe-owned fields when there is no Step-1 observe writer** (`DESIGN-REVIEW-2026-07-01.md` §2A.5). Two candidate resolutions m-7 can execute — **you pick/confirm which your contract supports:** (a) the observe-owned `required_when` predicates are **step-gated** on observe-layer presence (absent in Step-1 ⇒ not required), OR (b) a **Step-1 conductor-side filler** supplies them.
- **m-1:** confirm this against your **§5 Step-1 boundary** (store+form+lineage in Step-1, observe reserved Step-2 — the `readiness-fix-c1` seam).
- **m-2 co-sign:** the `required_when` schema half — which fields are observe-owned + how the step-gate/filler expresses in FieldSpec.
- **Closure artifact:** the confirmed resolution (a or b) + m-2 co-sign; map CQ-1 → closed/corrected. m-7 NF-S5 binds to it.

### CQ-6 — persisted seat-binding table + away-token sibling-burn  ·  m-1 + m-6 (token edge) co-sign  ·  CONFIRM-OR-CORRECT
m-7's §11 proposal: a **persisted seat-binding table** (seat ↔ channel/credential, survives conductor restart, re-attach without re-mint) + **away-token sibling-burn scope per decision** (`DESIGN-REVIEW-2026-07-01.md` §2C).
- **m-1:** confirm the binding-table semantics + the credential lifecycle boundary (m-7 executes; m-1 owns generation/rotation/revocation — m-1 §13.3 PLAN carry) + the **decision-scoped burn set** for sibling-burn (which sibling tokens a decision burns).
- **m-6 co-sign:** the away-token park/wake edge (the token's park/wake lifecycle m-6 owns).
- **Closure artifact:** confirmed binding-table + sibling-burn scope + m-6 co-sign; map CQ-6 → closed/corrected. m-7 NF-S1/NF-S2 bind to it.

### CQ-8 — INDEX derived-authority (NEW, from the grilled Q1 pivot)  ·  m-1  ·  CONFIRM-OR-CORRECT
m-7's §4 rename-pivot design makes **canonical records = truth; INDEX.md = a derived projection** (rebuilt from records at recovery, append-only, corrections appended never rewritten). m-7 reads this as **consistent with your locked m-1 §6 "reuse the upstream protocol's layout + append-only INDEX.md" — layout UNCHANGED, only crash-recovery provenance changes** (INDEX's *authority* becomes derived).
- **m-1:** confirm that reading, or flag it as a change to your locked on-disk layout (if the latter, it's a real contract question, not a silent change — that's exactly why m-7 raised CQ-8 instead of assuming).
- **Closure artifact:** confirm-or-flag; map CQ-8 → closed/corrected. m-7 §4 binds to it.

---

**CQ-status mapping (VP required-edit 3):** the closure artifact MUST map each CQ (CQ-1, CQ-6, CQ-8) to an exact status — **{closed · corrected-by-artifact · still-open / non-locking-carry}**.

**How this closes:** m-1 (+ co-signs) produces the confirm-or-correct closure; I fold it into m-7's design-lock package. m-7 stays holding until its full CQ gate closes.

Not authorized / not claimed: no locked-contract reopen, no cycle reopen, no new domain design beyond the named CQs, no PLAN/IMPL/`pcode/`/spike, no m-7 design-LOCK by implication, no operator decision reopened.

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py master/relays/c4-cq-m1/DESIGN-orchestrator-planner-20260702-012453.md` — OK
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-m1` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this COORD-2 dispatch + appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit, no CQ resolved.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-1 pair closes CQ-1/6/8 (full-pair, co-signed, CQ-status-mapped); I fold into the m-7 design-lock package.
