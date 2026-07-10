## COORD — c5 claim-sweep m-1 (the TCB, heavy — 37 candidate hits): relabel to the ratified checklist; owner fold + implementer semantic review

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c5-claim-sweep-m-1
PARENT_DISPATCH_ID: c5-decomp
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-1.planner
CC: master.orchestrator-reviewer, operator, m-1.implementer, m-7.planner, m-7.implementer
BUNDLE_ID: c5-claim-sweep-m-1
OWNER: m-1 (Trust & Identity)

m-1 — re-baseline **step (c)** global claim-sweep, your domain. This is **claim-text hygiene, NOT a mechanism change** (`DESIGN-REVIEW-2026-07-01.md` §2B: the mechanisms stay; only the malicious-lane *claims* collapse). You own the heaviest load — the TCB carried the "sole-writer / forgery-robust / identity-by-construction" claims that the WRAP-assumption inversion made the NO-GO's epicenter.

**The RATIFIED semantic checklist** (VP-approved exemplar `c5-claim-sweep-architecture/DESIGN-orchestrator-planner-20260702-131320` + the ARCHITECTURE sweep it documents):
- **RELABEL → confusion-resistant + D5 residual:** any claim a *malicious/same-uid* seat is stopped "by construction / sole-writer / sole write path / forgery-robust / unbypassable / tamper-resistant / non-lane-writable / a lane cannot write·forge·supply·mutate·re-tag·submit-as." Honest form: *a confused agent has no **tool** to do X* (GL D4) + the D5 accepted-risk (config / store-write / operator-FROM under same-uid attach) stated beside it.
- **KEEP (don't over-relabel):** trusted-engine control-flow/grammar invariants (serialized-loop kill; R2 no-model-predicate grammar); **observer-selected** control properties (invariant chosen conductor-side, unreachable by the observed lane); authority-ceiling "no tool granted"; already-scoped "no tool" claims. **Test:** *claims a malicious seat is stopped?* → relabel. *Holds by the engine's own control-flow/grammar, or an observer-side selection the lane can't reach?* → keep. Unsure → relabel + residual.

**Your candidate census (37 full-net hits — a STARTING point; produce your own complete classified survivor list).** CTO candidate relabels for the sharp ones:
- **I1 "sole-writer append-only store"** (:27, :31, :73, :125) → **governed-write** store (sole *governed* write path = `submit()`; confusion-resistant; D5: same-uid direct store write out of scope).
- **I2 "channel-isolation (the forgery-robustness core)" / "forgery-robust FROM"** (:78, :82, :87) → **confusion-resistant** FROM (a confused seat has no tool to forge it; D5 residual). Your §4 analysis (payload-lie ineffective, etc.) is the *confused-agent* argument — keep it, framed confusion-resistant.
- **"identity-by-construction"** (:174, :200, :226) → confusion-resistant identity; **"by-construction" is the shelved wrap milestone** (GL D3). *Note:* your doc ALREADY carries the honest-fallback at **:163** ("Option B operator-attested — only confusion-resistant, not by-construction") and **:226** ("by-construction and evidence-by-construction are distinct guarantees) — lead with those; the by-construction claim is the spike-gated future, not the Step-1 claim.
- **"sole write path" `submit()`** (:100) → sole **governed** write path.
- **KEEP candidates (classify explicitly):** DI-2/DI-5 mechanism descriptions (:38, :91 — conductor reads ground-truth from outside the lane = observer-side control property, KEEP+scope); the "malicious lane ... becomes" honest-caveat lines (:87, :200 already note the by-construction claim degrades under attach — these are ALREADY honest, KEEP); the migrator/layout lines (non-overclaim).

**Requirements (VP-set c5 shape + the r2 carry):**
1. **You (planner) fold/confirm** the relabels into your locked doc — claim-text only, no mechanism change, no c1 reopen.
2. **Produce your OWN complete classified survivor list** (full-net grep → each hit RELABEL / KEEP-with-reason / note). A top-of-doc scoping note is allowed, but **raw overclaim vocab surviving in mechanism text must be locally classified** (the miss that failed my ARCHITECTURE r1).
3. **m-1.implementer files a review-only `DESIGN-REVIEW` semantic approve** — confirms each relabel is faithful (no mechanism change), on the right side of the test, and carries the D5 residual where scoped.

Not authorized / not claimed: claim-text only; no mechanism/design-lock change, no c1 reopen, no PLAN/IMPL/`pcode`/spike, no operator decision reopened.

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py master/relays/c5-claim-sweep-m-1/DESIGN-orchestrator-planner-20260702-132042.md` — OK
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays/c5-claim-sweep-m-1` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this m-1 claim-sweep dispatch (census + candidate relabels + the ratified checklist) + appended `master/relays/INDEX.md`; no doc edit (m-1 owner-folds), no mechanism change, no code/`pcode/`, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-1.planner folds the relabels + files its classified survivor list; m-1.implementer semantic review-only approve; I fold the closure into the c5 status ledger.
