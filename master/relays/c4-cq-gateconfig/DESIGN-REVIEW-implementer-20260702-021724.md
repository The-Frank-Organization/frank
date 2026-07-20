## DESIGN-REVIEW - m-3.implementer re-review of c4-cq-gateconfig m-3 revision

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c4-cq-gateconfig
PARENT_DISPATCH_ID: c4-cq-gateconfig
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded re-review of prior must-revise blockers; no new operator decision surfaced
GRILL_REQUIRED: no
DESIGN_DOC_ID: c2-design-m-3-observation-evidence
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: c2-grill-m-3
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: c4-cq-gateconfig/DESIGN-planner-20260702-021030.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-4.planner, m-6.planner, m-2.implementer, m-4.implementer, m-6.implementer, operator
BUNDLE_ID: m-3-observation-evidence

DESIGN_REVIEW_VERDICT: approve

I re-reviewed the bounded revision in `c4-cq-gateconfig/DESIGN-planner-20260702-021030.md` against the live m-3 design doc, the m-2 field home, the m-7 terminal/config contracts, the operator decision register, and the c4 CQ coordination relays.

The prior `must-revise` blockers are folded. This approval is for the m-3 CQ-2/CQ-4/CQ-4b delta only; it does not close the global CQ gate, does not design-lock m-7, and does not authorize PLAN, IMPL, `pcode/`, or a spike.

## Resolved findings

1. **Decision-2 fail-open text is now byte-consistent.**

The stale universal delivery/no-delivery-gate wording has been narrowed to the class-conditional rule the operator recorded: non-authority `self_reported` records may deliver, while authority/merge A-gate `authority_class == true && record_integrity == self_reported` is terminal `held`/escalated and not delivered. The live m-3 doc now states this in the frame, veto conditions, opaque-lane floor, resolved decision, and fold-log (`master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:15`, `:63`, `:111`, `:130`, `:175`, `:177`, `:216`). That matches operator decision 2 (`master/READINESS-REGISTER.md:340-344`), the readiness review (`master/DESIGN-REVIEW-2026-07-01.md:142-143`), and the m-2 `authority_class` home/key (`master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:329-336`). R2 remains intact because the key is record class, never model identity.

2. **CQ-4 no-`submitted`-limbo wording is now byte-consistent.**

The stale "submitted -> accepted" wording is gone from the m-3 design artifact. The current text says a candidate remains in-courier through pre-append form/lineage/observe pre-flights, then reaches exactly one terminal outcome `{accepted | rejected | held}` with no persisted `submitted` limbo (`...observe-evidence-design.md:19`). The `slot_in` carry-forward lines now phrase the open PLAN detail as submit pre-flight versus terminal-outcome commit, explicitly not a persisted `submitted` state (`...observe-evidence-design.md:101`, `:200`). That aligns with m-2's terminal `delivery_state` set (`...form-schema-design.md:270-273`) and m-7's locked terminal-state enum / NF-S16 (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:97-104`, `:165`).

## Caveat preserved

The surfaced `authority_class == true && record_integrity == mixed` edge is correctly left open rather than silently broadened. The folded m-3 disposition stays keyed to `record_integrity == self_reported`, as m-2 currently declares and as operator decision 2 names. The m-3 doc flags the possible `mixed` extension for m-2/CQ-4 co-sign (`...observe-evidence-design.md:111`). This approval does not settle that edge; it approves the bounded byte-consistency fold while preserving the cross-domain question.

## CQ-status review

- CQ-2 -> approved for the m-3 fold leg: corrected-by-artifact, keyed on m-2 `authority_class`, no model-keyed gate.
- CQ-4 -> approved for the m-3 half: Step-2 observe-bounce shares terminal `rejected` because the check ran and said no; `held` remains distinct for authority-class unobservability/check-could-not-run. Joint closure still requires the m-2/m-6 co-sign path and orchestrator fold.
- CQ-4b -> m-3 confirm remains acceptable: the section-composed single top-level digest config artifact preserves m-3 egress/evidence/registry assumptions under m-7 trusted startup/load constraints (`...conductor-core-design.md:106-110`). Global closure still depends on all required owner confirms/reviews and CTO fold.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-cq-gateconfig/DESIGN-planner-20260702-021030.md` - OK
- `nl -ba master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md | sed -n '1,45p;88,122p;160,208p;220,250p'` - reviewed revised §1, §5.1, §6, §12, §13, §15
- `nl -ba master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md | sed -n '260,342p'` - confirmed `delivery_state` and `authority_class` homes
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md | sed -n '90,115p;148,170p'` - confirmed terminal-state/config/NF-S15/NF-S16 constraints
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-cq-gateconfig/DESIGN-REVIEW-implementer-20260702-021724.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-gateconfig` - OK
- `git -C pcode status --short` - clean
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no design-doc edit, no code/source/`pcode` edit, no PLAN, no IMPL, no spike.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (docs workspace; final `git status --short` exits 128)
Next requested action: m-3.planner may relay this approval into the c4 CQ closure thread; orchestrator still needs the full co-owner confirm/review set before folding the CQ gate into the m-7 design-lock package.
