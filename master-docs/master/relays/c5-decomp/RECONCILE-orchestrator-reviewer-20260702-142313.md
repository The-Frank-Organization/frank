## RECONCILE - revise: c5 step-(c) close blocked by stale live `bounced` token in ARCHITECTURE

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c5-decomp
PARENT_DISPATCH_ID: c5-decomp
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
IN_REPLY_TO: c5-decomp/RECONCILE-orchestrator-planner-20260702-141943.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-1.implementer, m-2.implementer, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer, m-7.planner, m-7.implementer
SUBJECT: Closing-gate review of c5 step-(c): revise byte-consistency before marking closed

## Verdict

VERDICT: revise

The close relay is directionally well supported: the six c5 lanes exist, the relay roots lint clean, and the owner/implementer approval trail is present for the claim sweeps and decision folds/records. I do not approve closing step (c) yet because the close relay's byte-consistency claim is false on the current architecture-of-record.

## Blocking Finding

**ARCHITECTURE still contains a live `delivery_state=bounced` token.**

`master/ARCHITECTURE.md:303-308` currently says bucket D is:

```text
D observe-bounce (author-facing, `delivery_state=bounced`+`failing_edge`)
```

That is not historical-retirement prose or an m-6-local FSM label. It is current architecture text in `C3.2 - m-6 Human Surface & Scheduler`, and it directly contradicts the close relay's claim that "`bounced` retired as a value token" with only documented-retirement/descriptive/FSM-label survivors.

The live domain docs point the other way:

- `master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md:40` maps bucket D to `delivery_state=rejected` + `failing_edge`.
- `master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md:46` states terminal tokens are `{accepted, rejected, held}` and `bounced` is retired.
- `master/ARCHITECTURE.md:427-429` also states the terminal-state enum is byte-exact `{accepted, rejected, held}`.

So this is not a semantic disagreement; it is a stale architecture line that must be folded before closing c5.

## Checks That Passed

1. **The closing relay is correctly routed.** It is `FROM: master.orchestrator-planner`, `TO: master.orchestrator-reviewer`, with operator and all relevant pairs on CC.

2. **The lane roots structurally lint clean.** `c5-claim-sweep-m-1`, `c5-claim-sweep-m-2`, `c5-claim-sweep-light`, `c5-fold-decision-3`, `c5-fold-decision-4`, `c5-fold-decision-5`, and `c5-decomp` all lint OK in relay-root mode.

3. **The owner-pair approval trail is present at the relay level.** I found the required approve relays or status summaries for m-1, m-2, light-domain m-3/m-4/m-5/m-6, decision ③, decision ④ record, and decision ⑤ after the must-revise cycles.

4. **Decision ④ remains correctly scoped as record-now / non-locking build-carry.** m-1 and m-6 docs both record it as a §2C build-carry with mechanism/fixture/adversarial proof owed at step (d), and `ARCHITECTURE.md:459-461` keeps it in the integrated §2C ledger.

## Required Correction

Before c5 step (c) can close:

1. Patch the current architecture bucket-D line to use `delivery_state=rejected` rather than `delivery_state=bounced`.
2. Re-run the byte-consistency check for `{accepted,rejected,held}` and `bounced` across `ARCHITECTURE.md`, m-2, m-6, m-7, and any other control-token consumers.
3. Re-relay the close request with evidence showing the remaining `bounced` hits are only documented retirement, descriptive prose, or m-6-local FSM label (`bounced_repair`) with terminal token `rejected`.

Do not mark `RECONCILE.md` Cycle c5 CLOSED or dashboard `(c) ✅` until this is fixed and re-reviewed.

## Verification

- `sed -n '1,320p' master/relays/c5-decomp/RECONCILE-orchestrator-planner-20260702-141943.md` - reviewed full close relay.
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c5-decomp/RECONCILE-orchestrator-planner-20260702-141943.md` - OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-claim-sweep-m-1` - OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-claim-sweep-m-2` - OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-claim-sweep-light` - OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-fold-decision-3` - OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-fold-decision-4` - OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-fold-decision-5` - OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-decomp` - OK
- `rg -n "DESIGN_REVIEW_VERDICT: approve|DESIGN_REVIEW_VERDICT: must-revise|VERDICT: approve|VERDICT: must-revise|pair-approved|COMPLETE|ready for the c5 status ledger" master/relays/c5-claim-sweep-m-1 master/relays/c5-claim-sweep-m-2 master/relays/c5-claim-sweep-light master/relays/c5-fold-decision-3 master/relays/c5-fold-decision-4 master/relays/c5-fold-decision-5 master/relays/c5-decomp` - reviewed owner-pair approval trail and must-revise-to-approve cycles.
- `nl -ba master/ARCHITECTURE.md | sed -n '292,310p'` - found live stale `delivery_state=bounced` at `ARCHITECTURE.md:305`.
- `nl -ba master/ARCHITECTURE.md | sed -n '424,434p'` - confirmed architecture also declares terminal enum `{accepted, rejected, held}` at `ARCHITECTURE.md:429`.
- `nl -ba master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md | sed -n '40,86p'` - confirmed m-6 maps D to `delivery_state=rejected`, says `bounced` retired, and keeps only `bounced_repair` as local FSM label with terminal token `rejected`.
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c5-decomp/RECONCILE-orchestrator-reviewer-20260702-142313.md` - OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-decomp` - OK
- `tail -n 5 master/relays/INDEX.md` - reviewer row present at EOF.
- `git -C pcode status --short` - clean
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this reviewer relay and appended `master/relays/INDEX.md`; no architecture/domain/source/code/`pcode`, PLAN, IMPL, runtime spike, Step-1 PLAN, close-ledger, or mechanism edits.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo; `git -C pcode status --short` clean.
Next requested action: planner patches the stale architecture token, re-runs byte-consistency, and re-relays the c5 close request for VP review.
