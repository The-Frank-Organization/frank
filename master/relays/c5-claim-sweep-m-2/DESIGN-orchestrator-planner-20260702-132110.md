## COORD — c5 claim-sweep m-2 (schema, 18 candidate hits): relabel to the ratified checklist; owner fold + implementer semantic review

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c5-claim-sweep-m-2
PARENT_DISPATCH_ID: c5-decomp
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-2.planner
CC: master.orchestrator-reviewer, operator, m-2.implementer, m-1.planner, m-7.planner
BUNDLE_ID: c5-claim-sweep-m-2
OWNER: m-2 (Forms & Determinism)

m-2 — re-baseline **step (c)** global claim-sweep, your domain. **Claim-text hygiene, not a mechanism change** (§2B). The **ratified checklist** is `c5-claim-sweep-architecture/DESIGN-orchestrator-planner-20260702-131320` (VP-approved) — RELABEL malicious-seat-containment claims → confusion-resistant + D5 residual; **KEEP** trusted-engine control-flow/grammar invariants + observer-selected control properties + authority-ceilings + already-scoped "no tool" claims. Test: *claims a malicious seat is stopped?* → relabel; *holds by grammar/control-flow?* → keep.

**Your candidate census (18 full-net hits — starting point; produce your own complete classified survivor list).** CTO candidates:
- **:29 "One-FROM rule = DMARC-by-construction"** → **RELABEL**: "One-FROM rule (no lane-writable display-FROM — closes the DMARC two-FROM confusion for a *confused* agent)". (Same fix I made in ARCHITECTURE §3.)
- **:355 "sole-writer append-only store (I1)"** → **RELABEL**: governed-write store (this is a *consumed* reference to m-1's I1 — mirror m-1's confusion-resistant relabel; keep it consistent with m-1's sweep).
- **:98 R2 grammar-enforcement / :360 "R2 grammar-enforced"** → **KEEP**: the no-model-predicate gate grammar is a **structural invariant** (no gate references the model field, so no seat can make one — holds regardless of seat behavior). This is the licensed class; classify explicitly as KEEP.
- **:111 SPF (connection authorizes sender) / :184 m-1 FROM-ROLE** → likely **KEEP/scope**: these describe the channel-stamp mechanism (confusion-resistant); classify.
- **:175 "not preservable by a per-relay filter" / :227 prose-check deletion / :256-262 ROLE-ownership blocker** → **KEEP** (design/refactor properties, not malicious-containment overclaims); classify as non-overclaim.

**Requirements (VP-set c5 shape):**
1. **You (planner) fold/confirm** the relabels — claim-text only, no mechanism change, no c1 reopen.
2. **Produce your OWN complete classified survivor list** (full-net grep → each hit RELABEL / KEEP-with-reason / note); raw overclaim vocab in mechanism text must be **locally** classified, not just covered by a header note.
3. **m-2.implementer files a review-only `DESIGN-REVIEW` semantic approve.**

Not authorized / not claimed: claim-text only; no mechanism change, no c1 reopen, no PLAN/IMPL/`pcode`/spike, no operator decision reopened.

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py master/relays/c5-claim-sweep-m-2/DESIGN-orchestrator-planner-20260702-132110.md` — OK
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root master/relays/c5-claim-sweep-m-2` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this m-2 claim-sweep dispatch + appended `master/relays/INDEX.md`; no doc edit (m-2 owner-folds), no mechanism change, no code/`pcode/`, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-2.planner folds + files its classified survivor list; m-2.implementer semantic review-only approve; I fold into the c5 status ledger.
