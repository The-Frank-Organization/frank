## DESIGN-REVIEW - m-5.implementer adversarial review of c6.1 confirm

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c6.1-confirm-m-5-review
PARENT_DISPATCH_ID: c6.1-confirm-m-5
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - owner-confirm review only; no new operator decision surfaced
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-5-workflows-archetypes
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: c6.1-confirm-m-5/DESIGN-planner-20260702-235314.md
FROM: m-5.implementer
TO: m-5.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: c6.1-confirm-m-5
OWNER: m-5 (Workflows & Archetypes)

DESIGN_REVIEW_VERDICT: approve

m-5.planner - I reviewed the c6.1 owner-confirm relay, the CTO c6.1 dispatch, the c6 differential VP instruction, `master/c61-fix.diff`, and the live m-1/m-5/m-7 design anchors. Verdict: approve.

This approval is review-only and scoped only to the single m-5 c6.1 observe-phasing delta. It grants no mechanism change, no design-lock reopen, no PLAN, no IMPL, no `pcode/`, no m-3 egress-disposition confirmation, and no m-5-F2 resolution.

## Review

1. The corrected section 9 text converges to the locked Step-1 boundary.

The c6.1 diff changes exactly the m-5 section 9 observe-phasing sentence from the old "every send observes" Step-1 claim to the current split: the send-gate/chokepoint is present by design, but the observe hook is inert in Step-1 and observe predicates land at Step-2 (`master/c61-fix.diff:63-70`; `m-5 design:160-164`). That now matches m-1's Step-1 boundary: Step-1 is store + form + lineage, with m-3 observe-as-send reserved as a Step-2 hook and "Step-1 records carry no observe gate" (`m-1 design:122-129`). It also matches m-7 NF-S5 / CQ-1(a): Step-1 submits with no observe layer are accepted without observe-owned fields, while the same submit requires those fields once the Step-2 observe layer is present (`m-7 design:151-156`).

2. I do not find a remaining section 7 vs section 9 contradiction.

The T1 table still says "observe-as-send gate always-on", but the same cell immediately scopes the `slot_in` invariant families to Step-2 and points to section 9 (`m-5 design:121-124`). Section 9 now supplies the needed disambiguation: "always-on" names the send-gate/chokepoint, not a Step-1 observe predicate (`m-5 design:163`). That is precise enough for this confirm; I do not require a T1 wording tighten. The section 4 observer-selected proof is also consistent because it is the design property for predicate selection when the observe hook exists; it does not claim Step-1 enforcement (`m-5 design:62-78`).

3. Lock invariants remain intact.

The correction does not change tag values, field ownership, F1/F2, R2, C2.4, GL-1..GL-6, authority ceilings, or the terminal enum. The live text still records `slot_in` classification at acceptance, reserves the predicate families for Step-2, and leaves the observer-selected-control claim scoped to the observe mechanism rather than Step-1 runtime enforcement (`m-5 design:17-23`, `:62-78`, `:160-164`, `:240-248`).

4. Scope boundaries are preserved.

The m-3 section 3.3 egress correction is not an m-5 delta; I am not confirming it here. The m-5-F2 posture/away-trigger scope question is also untouched by c6.1 and remains separately outstanding, exactly as the planner relay says.

No must-revise findings.

## Verification

Pre-write review evidence:
- `sed -n '1,220p' master/relays/c6.1-confirm-m-5/DESIGN-planner-20260702-235314.md` - read addressed m-5 planner confirm relay.
- `sed -n '1,220p' master/relays/c6.1-confirm-m-5/DESIGN-orchestrator-planner-20260702-233008.md` - read CTO owner-confirm dispatch.
- `sed -n '1,220p' master/relays/c6-differential/RECONCILE-orchestrator-planner-20260702-225941.md` - reviewed c6.1 differential source.
- `sed -n '1,220p' master/relays/c6-differential/RECONCILE-orchestrator-reviewer-20260702-232510.md` - reviewed VP requirement for owner confirmations before re-close.
- `sed -n '60,74p' master/c61-fix.diff` - checked exact m-5 hunk.
- `nl -ba master/domains/m-5-workflows-archetypes/design/2026-06-30-archetype-system-design.md | sed -n '60,88p;116,128p;158,166p;240,248p'` - checked section 4, T1, section 9, and c6 fold-log anchors.
- `nl -ba master/domains/m-1-trust-identity/design/2026-06-28-trust-identity-design.md | sed -n '118,134p'` - checked m-1 Step-1 build boundary.
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md | sed -n '150,158p'` - checked m-7 NF-S5/CQ-1(a) step gate.
- `rg -n "observe-as-send|observe hook|Step-1|Step-2|every send observes|c6\\.1|c61|m-5" master/c61-fix.diff master/domains/m-5-workflows-archetypes/design/2026-06-30-archetype-system-design.md master/relays/c6.1-confirm-m-5` - checked remaining observe-phasing language.
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c6.1-confirm-m-5/DESIGN-planner-20260702-235314.md` - OK.
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c6.1-confirm-m-5/DESIGN-orchestrator-planner-20260702-233008.md` - OK.
- `git -C pcode status --short` - clean.
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`.

Post-write verification:
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c6.1-confirm-m-5/DESIGN-REVIEW-implementer-20260703-000135.md` - OK (`OK master/relays/c6.1-confirm-m-5/DESIGN-REVIEW-implementer-20260703-000135.md`)
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c6.1-confirm-m-5` - OK (`OK master/relays/c6.1-confirm-m-5`)
- `perl -ne 'print "$ARGV:$.:$_" if /[^\x20-\x7E\n]/' master/relays/c6.1-confirm-m-5/DESIGN-REVIEW-implementer-20260703-000135.md` - no output (ASCII clean)
- `git -C pcode status --short` - clean
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`
- `rg -n "20260703-000135|DESIGN-REVIEW-implementer-20260703-000135|c6\\.1-confirm-m-5-review" master/relays/INDEX.md master/relays/c6.1-confirm-m-5/DESIGN-REVIEW-implementer-20260703-000135.md` - confirms the relay `DISPATCH_ID` and matching `master/relays/INDEX.md` row survived

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no domain-doc edit, no source/code/`pcode` edit, no PLAN, no IMPL, no runtime spike.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (docs workspace; `git status --short` exits 128); `git -C pcode status --short` clean.
RELAY_LINT: OK - `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c6.1-confirm-m-5/DESIGN-REVIEW-implementer-20260703-000135.md`.
DISPATCH_ROOT_LINT: OK - `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c6.1-confirm-m-5`.
Next requested action: m-5.planner returns the c6.1-confirm-m-5 completion relay to master.orchestrator-planner and operator.
