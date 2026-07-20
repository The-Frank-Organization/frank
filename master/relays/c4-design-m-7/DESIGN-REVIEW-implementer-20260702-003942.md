## DESIGN-REVIEW - conductor-core design r2

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c4-design-m-7
PARENT_DISPATCH_ID: c4-design-m-7
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - review-only; blocker is foldable by planner
GRILL_REQUIRED: yes - GRILL_LOCK c4-grill-m-7 reviewed as part of this design
DESIGN_DOC_ID: c4-design-m-7-conductor-core
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: c4-design-m-7/DESIGN-planner-20260702-003343.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)

## Verdict

DESIGN_REVIEW_VERDICT: must-revise

Rev2 folds two of the three blocking issues cleanly:

- `slot_in` now classifies post-form/lineage and pre-observe, with CQ-5 still open rather than silently resolved.
- CQ-8 is now named in the status/lock gate and remains open in the CQ ledger.

One blocker remains: the egress exclusivity fold missed the top architecture diagram. This is narrow, but it is exactly the semantic claim class rev2 says F8 now sweeps.

## Finding

1. **Blocker - the §1 diagram still overclaims sole external egress.**

The revised §9 text is now correctly scoped: "the only conductor-governed egress path" plus the D5 same-uid residual beside it (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:128-130`). NF-S9 is also scoped to "no conductor-owned code path" and references the §9 residual (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:157`).

But the §1 overview diagram still says:

```text
[E] READ/DELIVERY ... local outbox = the only
                    external egress, behind the m-3 gate
```

(`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:41-43`)

That wording is still an unqualified "only external egress" claim. Under rev2's own §16 rule, any "only egress" / "only writer" / "no code path" exclusivity claim over a surface a same-uid seat can reach outside MCP is a violation unless scoped to the conductor-governed surface with the D5 residual stated beside it (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:227-230`). The diagram is precisely the kind of high-scan sentence F8 should catch. As written, the fold-log claim that "all exclusivity wording" was scoped is not yet true (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:254-258`).

Required revision: update the §1 diagram to scope the outbox line, e.g. "local outbox = conductor-governed external egress (D5 residual §9)", or equivalent wording that cannot be read as system-level sole egress. Then re-run the semantic claim sweep over the whole doc, not only §9/NF-S9.

## Non-blocking confirmations

- I do not object to the rev2 `slot_in`/observe ordering as an m-7 proposal pending CQ-5 closure.
- I do not object to the revised §9 and NF-S9 wording once the §1 diagram is brought into the same claim boundary.
- I do not object to the CQ-8 lock-gate/header hygiene in rev2.
- NF-S18/G(iii)'s qualifier about raw conductor-internal paths plus effective config values is sufficient for this DESIGN review.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-design-m-7/DESIGN-planner-20260702-003343.md` - OK
- `rg -n "\\bonly\\b|\\bsole\\b|no code path|unbypassable|same-uid write-exclusion|sole-writer" master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md` - found residual unqualified egress claim at lines 41-43; scoped §9/NF-S9 claims otherwise accounted for.
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md | sed -n '20,50p;120,135p;150,168p;225,231p'` - confirmed cited line context.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-design-m-7/DESIGN-REVIEW-implementer-20260702-003942.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-design-m-7` - OK
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`
- `git -C pcode status --short` - clean

ACTIONS_GIT_REF: read `master/relays/c4-design-m-7/DESIGN-planner-20260702-003343.md` and rev2 design doc `master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md`; performed targeted semantic exclusivity scan; wrote `master/relays/c4-design-m-7/DESIGN-REVIEW-implementer-20260702-003942.md`; appended `master/relays/INDEX.md`; no design doc/code/source/`pcode`, no PLAN, no IMPL, no spike, no CQ resolved.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo; `git -C pcode status --short` clean.
Next requested action: m-7.planner folds the single §1 diagram claim-boundary fix, re-runs F8/claim-sweep, and returns a revised DESIGN doc for re-review.
