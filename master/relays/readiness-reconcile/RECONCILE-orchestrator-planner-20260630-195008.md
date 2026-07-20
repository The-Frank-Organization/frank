## RECONCILE — build-readiness register: GO-WITH-FIXES (7 lenses reconciled); VP review + operator decisions

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: readiness-reconcile
PARENT_DISPATCH_ID: readiness-decomp
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — review pass; surfaces operator-decision items (does not itself decide them); operator on CC
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator

Partner — the build-readiness review is done. Seven fresh read-only red-team lenses (composition · build-risk ·
assumptions · Step-1-readiness · adversary · operator/HITL · versioning), reconciled into
**`master/READINESS-REGISTER.md`**. The bound held: read-only, no code, no spike, no PLAN, no design changed, no
pair re-engaged. Bringing you the reconciliation for adversarial review + surfacing the operator decisions.

**VERDICT: GO-WITH-FIXES — unanimous (all 7 lenses).** The design-of-record is sound in the large and as a
*standalone-runtime target*; the core composition seams trace clean (routing_ref join, ODB bundle, B→A escalation,
egress chokepoint, observe-hook-not-a-double-writer); **nothing is no-go or unbuildable.** The review also *confirmed*
a lot works: observe-as-send TOCTOU genuinely closed by placement, J1 never-auto-approve airtight on every path, no
destructive migration, local operator-FROM unforgeable, away-token replay/forgery/expiry closed.

**Two root causes** thread the findings: **(1) ride-vs-own** — I1 sole-writer, "forgery-robust by construction," and
"sole external sender" were proven for a conductor that *owns the process tree*, but Step-1 *rides* rented runtimes,
and the honest-fallback discipline applied to DI-5 was not applied to I1/egress; **(2) cross-doc drift** — a few locked
docs fell out of sync c1→c3. The C3.6 capstone certified composition at the *architecture* altitude; these are
*doc-level* seams it did not reach — a fair and useful correction to my own capstone.

**Orchestrator-verified [V] (I read the exact lines, per your evidence bar):**
- **1a — the m-1/m-2 `submit()` write-path contradiction.** m-1 "no append-as-submitted → one atomic append-as-accepted,
  no lineage gate in the ordering" vs m-2 "append-as-`submitted` → cross-relay lineage gate → accepted." m-1's "reads
  identically" is false; the lineage gate that preserves "authority blocks before dispatch" has no hook-point in
  `submit()`. Found independently by 2 lenses. `m-1 …:91-94` vs `m-2 …:72-73`.
- **1b — the locked `submit()` embeds the m-3 observe pre-flight**, yet ROADMAP builds m-3 at Step-2 and §C2 titles it
  "Step-1 runtime-intelligence." "Store+form-gate without m-3" is not derivable from the locked pipeline. `m-1 …:93` ×
  `ROADMAP.md:57-70` × `ARCHITECTURE.md:120`.
- **4a — R2 "structurally impossible" is not grammar-enforced.** The generic `field:<id>` atom reaches `selected_model`;
  m-2 itself phrases the deviation gate as "required_when selected_model is off the prior floor"; the R2-safe
  `declared_deviated==true` lives only in m-4/ARCHITECTURE. m-2 (the doc the tool reads) is stale vs the c2 R2 lock.
  `m-2 …:84,283,285,289`.

**The gate before Step-1 (my recommended disposition — where I want your push):**
1. **MUST reconcile before any store PLAN** — **Cluster 1** (write-path; owner = **CTO+VP** + m-1/m-2, a shared-c1
   contract collision → charter-reserved to us) and **Cluster 4a/4b** (m-2 stale routing schema + the R2 grammar hole;
   owner m-2, bounded). These are design-gaps the store code hangs on.
2. **MUST get operator decisions** (surfaced, not decided here): **Cluster 2** wrap-vs-attach + "by construction vs
   honestly confusion-resistant until DI-2 is verified"; **5a** may A-gate/merge records deliver on `self_reported`;
   **6a** must agent-pick only RAISE toward A (never classify down to B); **6b** away-token refresh/re-observe binding;
   **6c** model-names-in-ODB vs the egress rule.
3. **SHOULD spike** (a separate operator-opened code step, not run here): **Cluster 3** the atomic multi-file append +
   write-serialization — one spike de-risks the nonce-burn, park/wake, and observe→append too.
4. **SHOULD reconcile as bounded VP-gated pair-fixes** (not Step-1-PLAN-blocking): **5b** (slot_in⟂opaque), **6d–g**,
   **7** (authority-containment), **8** (versioning compat-of-interpretation; 8a the parked-gate-across-bump is sharpest).

**Adversarial-review asks:**
- **Q1** — concur **GO-WITH-FIXES**, or do you read any finding as no-go / as a false positive (I flagged the deferred-
  infra items as build-risk, not gaps; push if I mis-graded one)?
- **Q2** — the **routing of fixes**: Cluster 1 → CTO+VP; Clusters 4/5b/6/7/8 → bounded owning-pair fixes (VP-gated, the
  c3 Seam-C pattern); Clusters 2/5a/6a-c → operator decisions. Right split?
- **Q3** — **sequencing**: do the MUST-reconcile (1) + MUST-decide (2) items gate opening the Step-1 PLAN, with the
  SHOULD items (3/4) runnable in parallel or folded into PLAN? Or a tighter/looser gate?

This is the operator's "measure twice" list — it is exactly what option-2 was for. On your review + the operator's
decisions, I route the bounded fixes to their owners and we re-verify, then Step-1 PLAN is the operator's to open.

ACTIONS_GIT_REF: ran 7 read-only reviewer subagents; wrote `master/READINESS-REGISTER.md` + this relay; no code/source/
pcode edits, no design-doc changes, no spike, no PLAN; cwd is not a git repo.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: VP reviews the reconciliation + the go/no-go + the fix-routing/sequencing; operator makes the Cluster-2/5a/6a-c decisions; then the orchestrator routes the bounded fixes + the Step-1 PLAN is the operator's to open.
