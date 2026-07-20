## PLAN-REVIEW — revision notice: the F14 mis-home is fixed exactly per your required shape; confirm requested (single-finding re-check)

ROLE: Orchestrator Planner
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s6-design
PARENT_DISPATCH_ID: s6-design
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — revision-confirm round of the standing pre-handoff review; the operator still holds the handoff
IN_REPLY_TO: master/relays/s6-design/RECONCILE-orchestrator-reviewer-20260706-180943.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator
SUBJECT: your must-revise finding accepted + applied in the held dispatch (joint-seam option): F14 invariant → m-1, runtime enforcement → m-7, the seam named explicitly, the dual-cite disposition rule added; nothing else changed; confirm to release the handoff

**Your blocking finding is accepted — it was a genuine mis-home.** The store lock guards the sole-governed-writer / store-isolation invariant, which is m-1's chartered contract (m-7 hosts and sequences it). My own ledger's owner split said the same; the dispatch contradicted it. Fixed in place in the held dispatch (`DESIGN-orchestrator-planner-20260706-180315.md`), taking your **joint-seam** option, exactly per your Required Revision Shape:

1. **m-1's section** now carries the **F14 store-lock INVARIANT** (lock semantics: what is locked, when held, staleness/takeover) with the VP-fold citation.
2. **m-7's section** now carries **F14 runtime enforcement only** (process choreography, startup/refusal/takeover behavior, fixture execution), pointing at m-1's invariant.
3. A new named **F14 seam** bullet sits beside the three-way parenting seam: invariant m-1 / enforcement m-7, and **F14's design-phase disposition must cite both halves** — the invariant statement and the runtime-fixture obligation.
4. The no-code / no-`frank/`-edits / no-build-slice-work constraint is untouched, as are all four of your non-blocking confirmations (in-step translation, fork gate geometry, the total-disposition bar, the amendment-then-build cadence).

No other line of the dispatch changed. Lint re-run clean (both modes) after the edits.

**On your confirm:** the operator relays the dispatch to m-1/m-7/m-2 and the s6 design-amendment phase runs.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py <this file>` + `--relay-root master/relays/s6-design` — run below.
- The revised dispatch: `master/relays/s6-design/DESIGN-orchestrator-planner-20260706-180315.md` (edited in place while HELD; the four edits enumerated above are its only deltas).

ACTIONS_GIT_REF: edited the held s6-design dispatch in place (the four F14 deltas) + wrote this notice + INDEX row; no git action, no `frank/` edit; cwd is a docs workspace (not a git repo).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main @ `7e5c527`, clean.
Next requested action: VP confirms the revision (single-finding re-check); on confirm the operator hand-relays the dispatch to the three pairs.
