## SITREP - m-7.implementer carry-forward on CQ closure correction

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: c4-cq-coord
PARENT_DISPATCH_ID: c4-cq-coord
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - routing/carry-forward acknowledgement only
GRILL_REQUIRED: no
FROM: m-7.implementer
TO: m-7.planner, master.orchestrator-planner
CC: master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
IN_REPLY_TO: c4-cq-coord/RECONCILE-orchestrator-planner-20260702-032227.md

## Routing / authority check

The two orchestrator-planner CQ closure relays are addressed to `m-7.implementer` (`031533`, `032227`). The VP follow-up `032843` is CC-only to `m-7.implementer`, so I treat it as context, not an independent action dispatch.

No `DESIGN_LOCK` package is visible yet in `master/relays/c4-design-m-7`; the latest m-7 design-lane artifact remains `SITREP-planner-20260702-005512.md`. This SITREP is therefore only a carry-forward acknowledgement, not the final adversarial lock-package review.

## Carry-forward

I do **not** read the corrected CQ-6 scoping differently from `032227` / VP approval `032843`.

- CQ-6 is design-lock-bearing only on the reviewed base: persisted seat-binding table, re-attach credential proof, decision-scoped `(decision_id, seat)` sibling-burn, and atomic burn inside the commit loop.
- `re-mint-supersedes` is not pair-approved as part of CQ-6 for this lock; it is carried as a non-locking §2C away-bridge build-step item.
- For the m-7 lock package, bind NF-S1/NF-S2 only to the CQ-6 base closure and list `re-mint-supersedes` as a §2C build-carry / away-bridge build-step review item.
- Preserve the other three assembly items from `031533`: m-4 per-section version-stamp inside the single-digest config artifact, byte-exact `{accepted, rejected, held}` with stale `bounced` swept/translated, and m-3's exactly-one-outcome framing check.

## Checked evidence

- `c4-cq-coord/RECONCILE-orchestrator-planner-20260702-031533.md:32-39` - m-7 assembly instructions and three fold-integration items.
- `c4-cq-coord/RECONCILE-orchestrator-planner-20260702-032227.md:17-32` - corrected CQ-6 base closure plus `re-mint-supersedes` non-locking carry.
- `c4-cq-coord/RECONCILE-orchestrator-reviewer-20260702-032843.md:21-39` - VP approval of the correction and carry-forward conditions.
- `c4-cq-m1/DESIGN-planner-20260702-013500.md:35-41` + `c4-cq-m1/DESIGN-REVIEW-implementer-20260702-020418.md:35-41` - m-1 base answer and implementer approval, with the add-on explicitly excluded from that review.
- `c4-cq-m1/DESIGN-planner-20260702-020100.md:25-37` + `c4-cq-m1/DESIGN-REVIEW-implementer-20260702-020447.md:21-38` - m-6 co-sign/review keeps the base park/wake edge approved and treats `re-mint-supersedes` separately.
- `master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md` - current design-of-record shows partial CQ fold integration, but §15 still carries open CQ rows; final lock package review waits for the assembled `DESIGN_LOCK`.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c4-cq-coord/RECONCILE-orchestrator-planner-20260702-031533.md` - OK
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c4-cq-coord/RECONCILE-orchestrator-planner-20260702-032227.md` - OK
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c4-cq-coord/RECONCILE-orchestrator-reviewer-20260702-032843.md` - OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-coord` - OK
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c4-cq-coord/SITREP-implementer-20260702-033926.md` - OK
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`
- `git -C pcode status --short` - clean

ACTIONS_GIT_REF: read the three operator-relayed CQ closure/correction relays plus CQ-6 base/co-sign evidence; wrote `master/relays/c4-cq-coord/SITREP-implementer-20260702-033926.md`; appended `master/relays/INDEX.md`; no design-lock package reviewed, no design doc/code/source/`pcode`, no PLAN, no IMPL, no spike, no CQ resolved by this SITREP.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo; `git -C pcode status --short` clean.
Next requested action: m-7.planner assembles the DESIGN_LOCK package on the corrected CQ-6 basis; m-7.implementer performs the final adversarial pass only after that package exists.
