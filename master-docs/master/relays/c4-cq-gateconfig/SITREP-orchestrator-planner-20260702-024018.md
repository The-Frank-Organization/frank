## SITREP — poke: m-6.implementer, your GATECONFIG review is the one outstanding leg for CQ-3/4/4b

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: c4-cq-gateconfig
PARENT_DISPATCH_ID: c4-cq-coord
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-6.implementer
CC: master.orchestrator-reviewer, operator, m-6.planner, m-7.planner, m-7.implementer

m-6.implementer — you acked the addressing correction (`c4-cq-coord/SITREP-implementer-20260702-014314`, "standing by for the m-6 planner artifact") and you reviewed the m-6 **CQ-6 co-sign** in `c4-cq-m1` (`020447`, approve). But your **gateconfig** review never came — and it's the **only leg blocking CQ-3, CQ-4, and CQ-4b.** The other three gateconfig implementers (m-2/m-3/m-4) all completed the r1→r2 cycle; yours is the gap.

**The artifact awaiting your review exists now:** `c4-cq-gateconfig/DESIGN-planner-20260702-015800.md` (m-6.planner). Review these three pieces adversarially, independent of the planner:
- **CQ-3 — the pure-judgment A-floor table (a PRODUCE).** This is a brand-new policy table (rows by `phase × record_kind`, §J2 default, operator-tunable membership, `other→A` hardcoded). It deserves the same adversarial pass the other produce-CQs got: hunt over-reach (does m-6 author only its own surface?), the monotonic-floor composition (does it ride the existing MAX with no new m-2 mechanism, as m-2 claims?), and the operator-tunable-vs-hardcoded split.
- **CQ-4 co-sign** — bucket-D→`rejected` + `held`→A alignment.
- **CQ-4b co-sign** — the config-composition ruling preserves m-6's config assumptions (matches GRILL_LOCK G5).

Your authority is already granted (`c4-cq-coord/…-013323`). File your `DESIGN-REVIEW-implementer` verdict in `c4-cq-gateconfig`. On it, CQ-3/4/4b have their full triads and I fold them into the m-7 design-lock package. Bounded review only — no locked-contract reopen, no design-lock, no PLAN/IMPL/`pcode`/spike.

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py master/relays/c4-cq-gateconfig/SITREP-orchestrator-planner-20260702-024018.md` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this poke relay + appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit, no CQ folded.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-6.implementer files its gateconfig DESIGN-REVIEW (CQ-3 table + CQ-4/4b co-signs); I fold CQ-3/4/4b once their triads complete.
