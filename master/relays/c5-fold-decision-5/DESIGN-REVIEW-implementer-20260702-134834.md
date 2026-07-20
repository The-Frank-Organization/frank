## DESIGN-REVIEW - m-3.implementer review of decision-5 m-3 egress fold

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c5-fold-decision-5
PARENT_DISPATCH_ID: c5-fold-decision-5
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded semantic review of recorded operator decision fold; no new operator decision surfaced
GRILL_REQUIRED: no
DESIGN_DOC_ID: c2-design-m-3-observation-evidence
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: c2-grill-m-3
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: c5-fold-decision-5/DESIGN-planner-20260702-133443.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-4.planner, m-6.implementer, m-4.implementer, m-7.planner, operator
BUNDLE_ID: m-3-observation-evidence

DESIGN_REVIEW_VERDICT: must-revise

I reviewed `c5-fold-decision-5/DESIGN-planner-20260702-133443.md` against the orchestrator dispatch, the current m-3 design doc, the m-6 ODB half, the m-4 R2 half, and the ratified c5 claim-sweep checklist.

The substantive policy shape is right: a narrow typed ODB model-name carve-out, confidentiality class only, safety/content still active, all other model-name egress blocked, transport still m-6-gated, and R2 untouched. The blocker is wording/classification: the fold uses fresh "Non-forgeable" / "lane cannot mark arbitrary bytes" language without local D5 scoping, in a mechanism section that the c5 checklist says must be relabeled or explicitly classified.

## Blocking revision

1. The ODB exemption's "non-forgeable" wording overclaims unless locally scoped.

Evidence:
- `master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:122` says the exemption is "Non-forgeable" and that "a lane cannot mark arbitrary bytes" exempt.
- The same wording class is repeated in the resolved decision, fixture, and fold-log at `:178`, `:204`, and `:221`.
- The ratified c5 checklist requires this exact class to become confusion-resistant/no-tool wording with D5 residual, unless locally justified as a trusted-engine grammar/control-flow or observer-selected KEEP (`c5-claim-sweep-architecture/DESIGN-orchestrator-planner-20260702-131320.md:27-31`; VP carry-forward at `c5-claim-sweep-architecture/RECONCILE-orchestrator-reviewer-20260702-131709.md:37-41`).

Required revision: replace "Non-forgeable" with the narrower property actually established here: "not lane-supplied through the governed ODB render/scan surface; a confused lane has no tool to mark arbitrary bytes as the exempt model-name field; D5 same-uid direct-store residual remains per §11." Apply that correction wherever the decision-5 fold repeats the claim (§7, §12, §13, §15), or add an explicit local KEEP classification if the planner is claiming a trusted-engine grammar/control-flow property.

## Confirmed non-blockers

1. The carve-out is correctly typed and narrow.

The m-3 doc scopes the exemption to `record_kind == ODB`, `field == model_name`, `destination == operator`, and the confidentiality scan class only (`...observe-evidence-design.md:122`). That matches the orchestrator dispatch (`c5-fold-decision-5/DESIGN-orchestrator-planner-20260702-132752.md:21-26`).

2. Safety/content still blocks.

The current m-3 fold keeps secrets, PII, auth URLs, and other content-safety hits active for any ODB field, and preserves the fail-closed `other->block` default (`...observe-evidence-design.md:120-122`, `:204`). The fixture list correctly requires ODB model-name pass plus non-ODB model-name block, any-ODB secret/PII/auth-URL block, and lane-supplied exempt-marked content block (`:204`).

3. R2 stays untouched from the m-3 side.

The design states the carve-out relaxes egress confidentiality only and does not make model identity gate-referenceable (`...observe-evidence-design.md:21`, `:122`, `:178`, `:221`). The m-4 half confirms the same two-axis separation in its current design (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:459-472`).

4. The m-6 seam exists but global close remains multi-party.

m-6 has authored the typed exempt-marked ODB model-name field and keeps non-away render local plus away-mode bridge opt-in/egress gating (`master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md:61`). That supports the m-3 scan precondition, but this review approves neither the m-6 nor m-4 implementer legs. Joint close still requires the sibling owner-pair reviews and orchestrator fold.

No R2 change, no locked-contract mechanism reopen, no PLAN, no IMPL, no `pcode/`, and no spike are authorized by this review.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c5-fold-decision-5/DESIGN-planner-20260702-133443.md` - OK.
- `sed -n '1,240p' master/relays/c5-fold-decision-5/DESIGN-planner-20260702-133443.md` - reviewed full m-3 planner relay.
- `sed -n '1,240p' master/relays/c5-fold-decision-5/DESIGN-orchestrator-planner-20260702-132752.md` - reviewed orchestrator decision-5 dispatch.
- `rg -n "model_name|model-name|ODB|Owner Decision Brief|confidentiality|safety|content|exempt|operator|away|R2|egress|destination|record_kind|field ==|scan-class|other.*block|fail-closed" master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md` - reviewed m-3 fold locations.
- `rg -n "c5-fold-decision-5|ODB model-name|model_name|R2 untouched|confidentiality" master/relays/c5-fold-decision-5 master/domains/m-4-routing-policy master/domains/m-6-human-surface-scheduler` - checked m-4/m-6 sibling authored halves.
- `sed -n '1,220p' master/relays/c5-claim-sweep-architecture/DESIGN-orchestrator-planner-20260702-131320.md` - reviewed ratified claim-sweep checklist.
- `sed -n '1,180p' master/relays/c5-claim-sweep-architecture/RECONCILE-orchestrator-reviewer-20260702-131709.md` - reviewed VP carry-forward for local classification.
- `git -C pcode status --short` - clean.
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no design-doc edit, no source/code/`pcode` edit, no PLAN, no IMPL, no runtime spike.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (docs workspace; `git status --short` exits 128)
RELAY_LINT: OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c5-fold-decision-5/DESIGN-REVIEW-implementer-20260702-134834.md`
DISPATCH_ROOT_LINT: OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-fold-decision-5`
Next requested action: m-3.planner patches the decision-5 wording/classification only, then re-relays for a narrow re-review; the m-3 substantive carve-out should be approvable after that text fix if no new raw survivor appears.
