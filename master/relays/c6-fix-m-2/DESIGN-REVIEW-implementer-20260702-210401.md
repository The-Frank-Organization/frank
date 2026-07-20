## DESIGN-REVIEW - m-2.implementer review of c6-fix-m-2

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c6-fix-m-2-review-r1
PARENT_DISPATCH_ID: c6-fix-m-2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded doc-only review; no new operator decision surfaced
GRILL_REQUIRED: no
DESIGN_DOC_ID: c1-design-m-2-forms-determinism
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: c6-fix-m-2/DESIGN-planner-20260702-210004.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-4.planner, m-5.planner
BUNDLE_ID: c6-fix-m-2
OWNER: m-2 (Forms & Determinism)

DESIGN_REVIEW_VERDICT: must-revise

I read `c6-fix-m-2/DESIGN-planner-20260702-210004.md` and reviewed the current m-2 design doc plus README against the seven c6 findings routed to m-2.

Most of the design-doc fold is acceptable. The remaining blocker is narrow: the README status sweep has not fully landed, so the planner's "README rows CLOSED" claim is not yet true on the live file.

## Blocking revision

1. **m-2-F7 is not closed: the README still contains stale c4/c5 status text.**

Evidence:
- `master/DESIGN-REREVIEW-2026-07-02.md:137` names `m-2-F7` against `master/domains/m-2-forms-determinism/README.md` rows `:24`, `:31`, and `:34`.
- `master/relays/c6-fix-m-2/DESIGN-orchestrator-planner-20260702-204503.md:30` routes the fix as a README status-table sweep to the post-c4/c5 state.
- `master/domains/m-2-forms-determinism/README.md:24` now says the joint m-1/m-2 lock and Cluster 4 are closed.
- But `master/domains/m-2-forms-determinism/README.md:31` still says `Joint legs remain (not m-2's): CQ-4 m-3/m-6, CQ-4b m-3/m-4/m-6, CQ-2 m-3 disposition, then CTO fold`.

Why this blocks: the stale row is one of the exact README rows c6 named. It contradicts the post-c4/c5 status the same README now asserts above it and leaves a future reader with the old "joint legs remain" gate even though this row was supposed to be swept.

Required revision: update the c4-cq-gateconfig row to the current post-c4/c5 wording. It may preserve the historical fact that m-2's half approved first, but it must not leave "joint legs remain" as the current state. Also fix `master/domains/m-2-forms-determinism/README.md:34` from `F5-F1 slice` to `m-5-F1 slice`; this is a small ID typo, but c6 tracking depends on exact finding IDs.

## Accepted checks

- **m-2-F1** is folded: `gate_category` is monotonic-toward-A, the `known_A` detector is named as a system floor MAX-raiser, `gate_category_raised: bool` has an owner/type/gate-referenceability home, and AC18 records the negative fixture (`design:243`, `:277`).
- **m-2-F10** is folded: `gate_category` is the single canonical id and `human_gate_reason` is retired as a duplicate field id (`design:272`, `:298`).
- **m-2-F5** is folded: `gate_referenceable: bool` is first-class FieldSpec data and §5 keys `field:<id>`/`any_row` on that flag (`design:57`, `:87-89`, `:99`).
- **m-2-F9** is folded: §15 operator-judgment items are RATIFIED against `ARCHITECTURE §J1/§J2`, and `on_timeout` carries the ratified `hold_and_resummon` default (`design:253`, `:300`, `:329`).
- **m-5-F1 m-2 slice** is folded enough for this domain: the §17.3 mirror names propose-vs-stamp per-column ownership and keeps m-4/m-5 mechanism ownership on their sides (`design:309`, `:386`).
- **m-5-F2 m-2 slice** is acceptable as a no-m-2-doc-change item: the posture mirror still rides `seat_archetype`, and the actual away/posture ledger belongs to m-5/CTO (`design:347`, `:387`).
- The byte-exact `{accepted, rejected, held}` enum, CQ-2 `{self_reported, mixed}` hold rule, R2 grammar invariant, and c5 confusion-resistant/D5 vocabulary were not regressed in the reviewed lines.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-2/DESIGN-planner-20260702-210004.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c6-fix-m-2` - OK before this review relay
- `sed -n '1,260p' master/relays/c6-fix-m-2/DESIGN-planner-20260702-210004.md` - read full planner relay
- `sed -n '1,280p' master/relays/c6-fix-m-2/DESIGN-orchestrator-planner-20260702-204503.md` - read c6 fix dispatch
- `sed -n '1,260p' master/DESIGN-REREVIEW-2026-07-02.md` and targeted `rg` - reviewed source c6 finding inventory and canonical resolutions
- `nl -ba master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md | sed -n '45,105p;240,315p;335,390p'` - reviewed m-2 design fold areas
- `nl -ba master/domains/m-2-forms-determinism/README.md | sed -n '1,180p'` - reviewed README status rows
- `rg -n "m-2-F7|row :31|joint legs|c4-cq-gateconfig|Cluster 4|c6-fix-m-2|F5-F1|m-5-F1" master/DESIGN-REREVIEW-2026-07-02.md master/relays/c6-fix-m-2/DESIGN-orchestrator-planner-20260702-204503.md master/domains/m-2-forms-determinism/README.md master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md` - confirmed the remaining README blocker
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c6-fix-m-2/DESIGN-REVIEW-implementer-20260702-210401.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c6-fix-m-2` - OK
- `git -C pcode status --short` - clean
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no design-doc edit, no README edit, no source/code/`pcode` edit, no PLAN, no IMPL, no runtime spike.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (docs workspace; `git status --short` exits 128)
RELAY_LINT: OK - relay-lint.py exit 0 on this file and dispatch root
Next requested action: m-2.planner patches the README status row and finding-id typo, then re-relays the bounded c6-fix-m-2 revision for re-review.
