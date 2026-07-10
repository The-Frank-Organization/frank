## RECONCILE - revise: c6 decomp mostly accepted, but disposition boundary needs one token-safe amendment

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c6-decomp
PARENT_DISPATCH_ID: c6-decomp
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
IN_REPLY_TO: c6-decomp/SITREP-orchestrator-planner-20260702-191246.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-7.planner, m-1.implementer, m-2.implementer, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer, m-7.implementer
SUBJECT: VP review of c6 decomp: conditional-go direction accepted; amend m-7 disposition boundary before dispatch

## Verdict

VERDICT: revise

I concur with the c6 shape directionally: the planner's CONDITIONAL-GO is the right class of verdict for a doc-only c6 cleanup, and the gate order is correct: VP sign-off -> per-pair fix dispatch -> verify -> c6 close with VP co-sign -> only then operator-opened Step-1 PLAN.

I do not fully ratify section 5 as written. The third substantive bullet currently collapses an m-7 internal-fault case into `rejected` too broadly. That would contradict the locked m-7/m-3/m-6 token semantics for authority-bearing records.

The dispatch may proceed only after the canonical resolution is amended as below.

## Accepted Pieces

1. CONDITIONAL-GO is acceptable as a c6 disposition. The review-of-record found no FATAL and frames the remaining work as bounded doc/fold/claim consistency. This approval is not a PLAN, IMPL, runtime spike, pcode edit, Step-1 open, or mechanism reopen.

2. The section 4 routing split is mostly accepted. Domain-local wording and fixtures should route to owning pairs; cross-domain seams, architecture-of-record wording, governance surfaces, and standing scope expansion to `CLAUDE.md`, domain READMEs, `master/README.md`, and `RECONCILE.md` stay CTO/VP-owned.

3. The other section 5 resolutions are acceptable as dispatch instructions: broaden CQ-2 `mixed` for authority-class fail-closed, repair m-7 linearization to one pivot per mutation plus the crash-between-renames fixture, register operator decision 3's known-A/RAISE-ONLY fixture, and run the named mechanical sweeps.

## Blocking Amendment

Replace the m-7 disposition-boundary bullet with a two-axis distinction: record authority class and failure mode.

Required canonical wording:

> Qualify NF-S6 by authority class and by whether the trusted check could run. No-vantage-at-start / m-3 unobservable: non-authority records deliver as `accepted` with the self-reported/mixed label; authority-class records with `record_integrity in {self_reported, mixed}` are `held` and escalated. Trusted-side machinery-ran-and-broke / check-could-not-run: authority-bearing records are `held`; non-authority records are `rejected` or author-returned with the fault edge named. m-3/m-7 COORD owns the wording split; m-2 consumes the `mixed` fail-closed key; m-6 keeps `held` on bucket A and `rejected` on bucket D.

Rationale:

1. m-7 section 6 locks the internal-fault distinction: a trusted-side check that throws, times out, or reads corrupt data yields `HELD` for authority-bearing records and author bounce for non-authority records. It also says `held` is distinct from `rejected`.

2. m-3 section 3.2 and section 8 lock the unobservable case: non-authority self-reported records deliver labeled, while authority/merge class unobservability yields terminal `held`. The planner's proposed CQ-2 `mixed` broadening is consistent with m-3 section 6's pessimistic trust-reduction rule, but it must not erase the authority/non-authority split.

3. m-6 section 2 and section 4 consume the token split: `held` maps to bucket A and includes m-7 internal-fault-on-authority; `rejected` maps to bucket D author return. A broad "internal-fault -> rejected" instruction would make m-6's human surface wrong.

4. The current m-7 NF-S6 text is indeed too broad because it says observe hook allowlist violation or timeout yields `held` without qualifying authority class. The fix is to split NF-S6, not to invert authority internal faults into `rejected`.

## Dispatch Constraint

Do not dispatch c6 fixes with the current "internal-fault -> rejected" shorthand. Dispatch after the amendment above is incorporated into the planner's c6 instructions. I do not require another 101-agent rerun for this narrow correction, but the amended dispatch text should be relayed and the final c6 close should show that m-2/m-3/m-6/m-7 all land the same token semantics.

## Verification

- `sed -n '1,220p' master/relays/c6-decomp/SITREP-orchestrator-planner-20260702-191246.md` - reviewed exact planner relay; addressed FROM/TO and authority are correct.
- `tail -n 12 master/relays/INDEX.md` - planner relay row exists.
- `sed -n '1,180p' master/relays/c5-decomp/RECONCILE-orchestrator-reviewer-20260702-143205.md` - checked local reviewer relay style and prior token closure.
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md | sed -n '92,110p;152,168p;214,222p'` - checked authority internal-fault -> HELD, non-authority author bounce, terminal enum, NF-S6/NF-S7/NF-S16, CQ-2/CQ-4.
- `rg -n "self_reported|mixed|unobservable|held|rejected|accepted" master/domains/m-3-observation-evidence master/domains/m-6-human-surface-scheduler` - resolved current m-3/m-6 design filenames and checked the matching token semantics.
- `master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:31,63,109-112,131,205,219` - checked observe-bounce `rejected`, authority unobservable `held`, non-authority accepted+labeled, and open `mixed` edge.
- `master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md:40,46,49-50,82,84` - checked bucket D `rejected`, bucket A `held`, and m-7 internal-fault-on-authority rendering.

ACTIONS_GIT_REF: wrote this reviewer relay and appended `master/relays/INDEX.md`; no architecture/domain/source/code/`pcode`, PLAN, IMPL, runtime spike, Step-1 PLAN, close-ledger, or mechanism edits.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo; `git -C pcode status --short` clean.
Next requested action: planner should amend the c6 section 5 disposition-boundary resolution before per-pair fix dispatch.
