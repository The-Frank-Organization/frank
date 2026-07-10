## SITREP — operator decision ① RECORDED: attach-first + "confusion-resistant"

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: runtime-research
PARENT_DISPATCH_ID: runtime-research
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — records a decision the operator already made; does not itself gate
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator

Partner — closing the loop you opened in `runtime-research/…-213911`. On the three-rung ladder you approved, the
operator has made **decision ① (Cluster 2 / 2a):**

**ATTACH-FIRST + "confusion-resistant."** Step-1 rides runtimes as an MCP-server with per-seat channels over
**persistent** pipes (persistent seats preserved — no per-relay spawn; the operator confirmed this was the load-bearing
concern); the courier owns identity + the gate; the honest public claim is **"confusion-resistant."** `srt`-wrap is a
*later* bounded hardening lane; **"sole external sender by construction" stays your two-stage-spike-gated milestone**
(properties 1–5 → "single mediated network path"; +6 broker-egress-gate → "destination + content control by
construction"). Recorded in `master/READINESS-REGISTER.md` → *Operator decisions — recorded log*.

**Per your Finding-3 routing, this unblocks two bounded m-1 fixes — which I am NOT opening here** (they route only as
bounded VP-gated fix relays, and the sequencing gate still holds):
1. the **Cluster-2 m-1 completeness fix** — record the "confusion-resistant" claim + the wrap-upgrade path;
2. the **identity-conductor-owned m-1 fix** — runtime identity fields (`session_source`/`clientId`) never accepted as
   `FROM`; conductor-owned per-seat channel/credential isolation is the sole stamp source (tightens the c1 identity
   lock, does not reopen it).

**Nothing else moved.** Still OPEN + unrecorded: decisions ② 5a, ③ 6a, ④ 6b, ⑤ 6c. Still unreconciled: Cluster 1
(m-1/m-2 write-path) and Cluster 4a/4b. **No Step-1 PLAN opens** until the full sequencing gate is satisfied. No spike
opened; the five/six-property security spike remains behind a separate operator-opened gate. No pair re-engaged, no
contract changed, no code — docs-only, cwd.

Next: the operator directs whether to take decisions ②–⑤ next, route the bounded MUST-reconcile fixes (Cluster 1,
4a/4b), or pause. I route nothing without the operator's go.

ACTIONS_GIT_REF: recorded decision ① in `master/READINESS-REGISTER.md`; wrote this relay + appended `master/relays/INDEX.md`; no code/source/pcode edits, no design-doc changes, no PLAN, no spike, no pair re-engaged.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: operator directs the next step (decisions ②–⑤ · route Cluster 1 + 4a/4b bounded fixes · or pause); VP has visibility on the recorded decision.
