## DESIGN-REVIEW - m-3.implementer review of c4-cq-gateconfig m-3 answer

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c4-cq-gateconfig
PARENT_DISPATCH_ID: c4-cq-gateconfig
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded review; no new operator decision surfaced
GRILL_REQUIRED: no
DESIGN_DOC_ID: c2-design-m-3-observation-evidence
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: c2-grill-m-3
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: c4-cq-gateconfig/DESIGN-planner-20260702-014846.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-2.planner, m-4.planner, m-6.planner, m-2.implementer, m-4.implementer, m-6.implementer, operator
BUNDLE_ID: m-3-observation-evidence

DESIGN_REVIEW_VERDICT: must-revise

I reviewed the m-3 CQ-2 / CQ-4 / CQ-4b answer in `c4-cq-gateconfig/DESIGN-planner-20260702-014846.md` against the folded m-3 design doc, the m-2 field-home, the m-7 CQ ledger, and the coordination dispatches.

The intended CQ shape is mostly right: `authority_class` is m-2-declared and model-free; Step-2 observe-bounce sharing `rejected` is the right terminal-token cut for "check ran and said no"; `held` is the right distinct outcome for authority-class unobservability; and the single top-level digest config composition does not break m-3's egress/check-registry/disposition assumptions. But the folded design artifact still contains stale contradictory text that would let a faithful builder ship the old fail-open behavior or a `submitted`-limbo reading. That blocks approval until patched.

## Must-revise findings

1. **Decision-2 fold is not byte-consistent: stale universal delivery/no-delivery-gate text remains.**

The new §3.2/§8 fold says authority-class `record_integrity == self_reported` is terminal `held` and not delivered (`master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:63`, `:130`). That matches the operator decision (`master/READINESS-REGISTER.md:340-344`) and the readiness review requirement (`master/DESIGN-REVIEW-2026-07-01.md:142-143`).

However, the same design doc still says self-reported completion "is delivered" without the authority-class exception (`...observe-evidence-design.md:15`), and §6 still says "No mechanical delivery gate branches on either label" (`...observe-evidence-design.md:111`). Those two sentences directly contradict condition (d). Revise them so they say the honest-fallback delivery applies only to non-authority records, while `authority_class == true && record_integrity == self_reported` is held/escalated. Also update the resolved-decision text that still describes the opaque-lane floor as unconditional "deliver-and-label" (`...observe-evidence-design.md:175`), or explicitly cross-reference the later decision-2 exception.

2. **CQ-4 no-limbo closure is undermined by stale `submitted -> accepted` wording.**

The CQ-4 answer claims the m-3 half aligns to the m-7 closed terminal set `{accepted, rejected, held}` with no persisted `submitted` limbo. The current §2 text supports that by committing a terminal `rejected` evidenced outcome record for observable failure (`...observe-evidence-design.md:31`), and m-2 declares the closed `delivery_state` set plus no `submitted` limbo (`master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:269-272`). m-7 likewise makes no-limbo part of the locked terminal enum (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:104`).

But the m-3 doc still frames the m-2 contract as a "Two-state `submitted -> accepted` write path" (`...observe-evidence-design.md:19`) and the `slot_in` carry-forward still contrasts "submit pre-flight vs the `submitted->accepted` transition" (`...observe-evidence-design.md:200`). Even if those are inherited/stale context lines, they sit inside the design artifact being folded for CQ-4 and conflict with the no-limbo closure. Replace them with the readiness-fix-c1 seam: candidate held in-courier through pre-append form/lineage/observe, then one terminal outcome.

## Confirmations

- **CQ-2 intended mechanism:** approve once the stale lines above are fixed. The fold keys on m-2's declared `authority_class`, which is `owner:system`, `computed_result`, and derived from record class or `gate_category in A`, never model identity (`master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:328-335`).
- **CQ-4 intended token cut:** approve once the no-limbo stale wording is fixed. `rejected` for checks that ran and failed, versus `held` for authority checks that could not run, preserves m-7's §6 distinction and NF-S16 exactly-one-outcome invariant (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:97-104`, `:165`).
- **CQ-4b m-3 confirm:** no correction needed from the m-3 side. The CTO section-composed single-digest artifact preserves m-3's assumptions: per-domain authorship, trusted startup load, restart-only no hot reload, digest verification, and operator-authorized append-only config changes (`master/relays/c4-cq-gateconfig/DESIGN-orchestrator-planner-20260702-012336.md:48-55`, `master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:106-110`).

## CQ-status review

- CQ-2 -> `must-revise`: intended m-3 disposition is correct, but stale fail-open/no-delivery-gate text remains in the folded design doc.
- CQ-4 -> `must-revise`: intended m-3 observe-bounce token mapping is correct, but stale `submitted -> accepted` wording must be removed or marked superseded before the no-limbo token closure is reliable.
- CQ-4b -> m-3 confirm is acceptable, but it cannot close globally until the required m-2/m-4/m-6 confirmations and implementer reviews exist and the CTO folds the composition ruling.

This review does not reopen c1/c2/c3, does not design-lock m-7, and does not authorize PLAN, IMPL, `pcode/`, or a spike.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-cq-gateconfig/DESIGN-planner-20260702-014846.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-cq-gateconfig/DESIGN-REVIEW-implementer-20260702-020509.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-gateconfig` - OK
- `nl -ba master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md | sed -n '1,180p;180,260p'` - reviewed folded §2, §3.2, §6, §8, §12, §13, §15
- `nl -ba master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md | sed -n '250,350p'` - confirmed `delivery_state` and `authority_class` homes
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md | sed -n '53,110p;153,166p;217,221p'` - confirmed terminal-state/config/CQ constraints
- `git -C pcode status --short` - clean
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no design-doc edit, no code/source/`pcode` edit, no PLAN, no IMPL, no spike.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (docs workspace; `git status --short` exits 128)
Next requested action: m-3.planner patches the stale fail-open/no-gate/no-limbo text and re-relays the bounded revision for re-review.
