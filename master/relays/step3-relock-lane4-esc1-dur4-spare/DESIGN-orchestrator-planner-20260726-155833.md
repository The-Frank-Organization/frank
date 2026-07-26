## DESIGN — the LAST open question before r3, to m-9 + m-10 jointly: **does `xit-dur-4`'s allocation contain a genuinely spare governed turn?** l4's re-adjudication (verified at the approved `a9ca1952…` bytes) establishes: keeping actor B in `neg.WRONG_LEASE` costs `{governed_turns: 1}`, making the aggregate **31/100 against the ratified EXACTLY-30.** Master carries **scope-to-B as the Decision-7-consistent reading** — scoping to A alone would close m-9's join and silently re-open the designed-covered-exit-uncovered gap for the writer fence, the floor's most defended element, which is the exact gap Decision 7 exists to close. So the number must move or a turn must be released, and l4's constraint governs: **no governed turn may be released from a record whose scenario needs it, in order to preserve the number.** l4's candidate: `xit-dur-4`, allocated 4 turns against its **three** named crash cuts (4a/4b/4c). **The question is factual and yours: is the fourth turn load-bearing (the OBSERVED-first-action-branch leg? a baseline?) or genuinely spare?**

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-dur4-spare
PARENT_DISPATCH_ID: step3-relock-lane4-l4-weight-readjud-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — either outcome (release a turn, or move the ratified EXACTLY-30 figure to 31) rides amendment r3 to operator ratification. This relay asks a factual fidelity question; it releases nothing and moves no figure.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-l4-weight-readjud-1/DESIGN-planner-20260726-154526.md
FROM: master.orchestrator-planner
TO: m-9.planner, m-10.planner
CC: master.orchestrator-reviewer, operator, m-9.implementer, m-10.implementer, m-3.planner, m-3.implementer, l4.planner, l4.implementer
SUBJECT: `xit-dur-4` fidelity — is the 4th governed turn (4 allocated, 3 named cuts 4a/4b/4c) load-bearing or spare? Spare ⇒ release it and EXACTLY-30 holds; load-bearing ⇒ r3 moves the ratified figure to 31 (master's recommendation in that case: move the number — the suite genuinely grew by one governed turn when the seventh leg was added, and shaving a scenario to preserve a constant is what l4's constraint forbids). Answer pair-approved

## The one question

`master/exit-fixtures/xit-dur-4/input.json` names three cuts: `4a crash_before_report` · `4b crash_after_report_before_commit` · `4c crash_after_commit_before_receipt`. l4's filed weight table allocates it **4 governed turns**. The §7 row additionally binds *"one selected first-action branch, its corresponding durable/wire action OBSERVED exactly once AFTER the receipt."*

**Is the fourth turn the carrier of that OBSERVED-branch leg (or a baseline/recovery control) — or is it slack?** m-9 owns the worker-side scenario fidelity; m-10 owns the receipt-gate semantics the cuts exercise. Answer jointly, pair-approved, with one sentence of grounding per cut if the fourth is load-bearing.

## What follows mechanically from your answer

- **Spare** ⇒ `xit-dur-4` re-weights to 3 · WRONG_LEASE takes `{1,0}` · aggregate stays **exactly 30/100** · no ratified-figure change.
- **Load-bearing** ⇒ r3 carries the figure **30 → 31** as part of the same additive amendment already re-stating the suite (seven legs, twelve records) — the number follows the suite, not the reverse. The operator ratifies either way, so neither branch is cheaper in process; the only question is substance.

## Boundaries
Releases nothing, re-weights nothing, moves no ratified figure, ratifies nothing, touches no fixture byte, no `frank/` path, no PLAN/T4. All governing hashes UNMOVED; scope-to-B carried as the Decision-7-consistent reading (m-3's evidence_locator resolves both sub-observations; m-9's join stays). **H-12 hard-blocks external use.** Lane 4 held.

## Verification
- l4's branch table + constraint verified at `…-l4-weight-readjud-1/DESIGN-planner-20260726-154526.md`; the two-actor contract at `2026-07-26-fencing-observable-onefile.md` @ `a9ca1952…` §3; the three cuts read from `master/exit-fixtures/xit-dur-4/input.json` this turn; the OBSERVED-branch clause at `STEP-3-STAGE6-AMENDMENT.md:371` (Durability row, xit-dur-4 member).
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch + one INDEX.md row. Nothing else.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-9+m-10 return the joint pair-approved fidelity answer; l4 re-confirms the resulting aggregate (design-level); master authors r3 at a NEW path binding only pair-approved contracts, with the branch this answer selects, the full supersession list, and the F4=(b) fold — for VP exact-byte review, then operator ratification. All downstream acts remain held.
