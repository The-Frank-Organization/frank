## SITREP — structural lineage fix required before dispatch: the r2 gated PLAN's parent edge breaks the design-lock lineage walk (root-mode lint ERROR); bounded re-issue prescribed

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s2-core-plan
PARENT_DISPATCH_ID: s2-core-design
RUN_ID: s2
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: s2-core-plan/PLAN-planner-20260704-030751.md
FROM: s2.orchestrator-planner
TO: s2-core.planner
CC: s2-core.implementer, s2.orchestrator-reviewer, operator
SUBJECT: Fix before any delegated dispatch — gated PLAN r2 parents the r1 PLAN-REVIEW instead of the approving DESIGN-REVIEW; re-issue as r3 with the corrected parent + fresh implementer approve; content unchanged

**The finding (mine, E2 — root-mode lint this session):**
`relay-lint --relay-root .relays/s2` errors on `s2-core-plan/PLAN-planner-20260704-030751.md`:
`design-doc PLAN parent 's2-core-plan-review-implementer' does not resolve to a relay in this lineage`.
Cause: the r2 gated PLAN carries `DESIGN_LOCK_ID: s2-slice-2-design` + `DESIGN_RECORD_KIND: design-doc`, so the design-review lineage gate walks its `PARENT_DISPATCH_ID` expecting the **approving DESIGN-REVIEW** — but r2 parents `s2-core-plan-review-implementer` (the r1 must-revise PLAN-REVIEW). Your r1 gated PLAN (`…-030052.md`) had it right (`PARENT_DISPATCH_ID: s2-core-design-r2-review-implementer`); the S1 precedent held that shape across plan revisions (s1 plan-lock-r3 still parented the approving DESIGN-REVIEW; the plan-review thread rode `IN_REPLY_TO`).

**Why it must be fixed now:** a relay-lint structural error blocks delegated dispatch (protocol; operator waiver is the only bypass, and a cheap reword beats a waiver). Your eventual `DISPATCH IMPL` parents the approving PLAN-REVIEW → which parents the gated PLAN → whose design-lock edge must resolve. Left as-is, the whole dispatch chain lints dirty at the moment it matters.

**Bounded fix (content UNCHANGED — this is a header-edge correction, not a plan revision):**
1. Re-issue the gated PLAN as **r3**: same body/locks (`PLAN_LOCK_ID: s2-slice-2-plan`, `DESIGN_LOCK_ID: s2-slice-2-design`, design-doc kind, plan doc unchanged at main@c16f261), `PARENT_DISPATCH_ID: s2-core-design-r2-review-implementer`, `IN_REPLY_TO` may thread the r1/r2 review trail, note the supersession of `s2-core-plan-lock-r2` in the SUBJECT/body.
2. s2-core.implementer re-issues the approving PLAN-REVIEW parented to the r3 dispatch id (findings unchanged from `…-031243.md` — an explicit re-affirmation on the corrected lineage; state that no plan content changed).
3. Run `relay-lint --relay-root .relays/s2` over both before handoff; the r3 pair must be root-mode clean (INDEX noise excepted).
4. Then proceed exactly as before: m-1 fidelity approve (still the substantive gate; unchanged by this) → SCOPE_DIFF (README row cites the fence ruling `…-032200.md`) → delegated dispatch parenting the NEW approving review.

Nothing else is reopened: the r2 approve's substantive findings stand; the plan doc is untouched; this costs two relay files. My own r1→r2 reconcile note is amended in the ledger to record the catch.

ACTIONS_GIT_REF: none — report-only fix routing; this file + an INDEX row under gitignored .relays/.
FINAL_GIT_STATUS_SHORT: none — clean tree
