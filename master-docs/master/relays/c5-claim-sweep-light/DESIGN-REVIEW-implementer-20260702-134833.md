## DESIGN-REVIEW - m-3.implementer review of c5 claim-sweep-light m-3 fold

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c5-claim-sweep-light
PARENT_DISPATCH_ID: c5-claim-sweep-light
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded semantic review of claim-text hygiene; no new operator decision surfaced
GRILL_REQUIRED: no
DESIGN_DOC_ID: c2-design-m-3-observation-evidence
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: c2-grill-m-3
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: c5-claim-sweep-light/DESIGN-planner-20260702-133039.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-4.planner, m-5.planner, m-6.planner, m-4.implementer, m-5.implementer, m-6.implementer, m-7.planner, operator
BUNDLE_ID: m-3-observation-evidence

DESIGN_REVIEW_VERDICT: must-revise

I reviewed `c5-claim-sweep-light/DESIGN-planner-20260702-133039.md` against the ratified c5 checklist, the current m-3 design doc, and the adjacent decision-5 fold that is now present in the same document.

The named five relabels are directionally correct and preserve the licensed m-3 primitives. I cannot approve the "full-net survivor list" as complete against the current document because the later decision-5 fold introduced new raw overclaim vocabulary in mechanism text that falls inside the ratified checklist's relabel net.

## Blocking revision

1. Fresh decision-5 "non-forgeable" / "lane cannot mark" wording is not locally classified.

Evidence:
- `master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:122` says the ODB exemption is "Non-forgeable" and that "a lane cannot mark arbitrary bytes" as exempt.
- The same survivor class appears in the resolved decision, fixture, and fold-log at `:178`, `:204`, and `:221`.
- The ratified c5 checklist explicitly puts "non-lane-writable" and "a lane/seat cannot write / forge / supply / mutate / re-tag / bypass / submit-as" in the RELABEL net unless locally justified as trusted-engine control-flow, grammar, observer-selection, or already-scoped no-tool wording (`c5-claim-sweep-architecture/DESIGN-orchestrator-planner-20260702-131320.md:27-31`; VP ratification at `c5-claim-sweep-architecture/RECONCILE-orchestrator-reviewer-20260702-131709.md:37-41`).

Why this blocks: the planner relay asks for approval of a complete full-net survivor list. The current m-3 file still has raw mechanism text in the exact claim class the c5 checklist hardened after r1. The global §11 claim-boundary helps, but the VP ratification says raw overclaim vocabulary in mechanism text still needs local classification. Leaving this as "Non-forgeable" reads stronger than the D5-bounded governed-surface claim.

Required revision: relabel the decision-5 ODB exemption wording in §7, §12, §13, and §15 to the honest shape, for example: "governed-field scoped / not lane-supplied through the ODB render path; a confused lane has no tool to mark arbitrary bytes exempt; D5 same-uid direct-store residual remains per §11." If the planner wants to KEEP part of it as trusted-engine control-flow, say that locally and keep the D5 residual beside the underlying field-surface claim.

## Confirmed non-blockers

The five named claim-sweep relabels are present and faithful:
- §1 now mirrors m-1 as a governed sole-writer governance-surface claim with D5 residual (`...observe-evidence-design.md:20`).
- §5/§5.1 now frames archetype invariants as observer-selected / confusion-resistant rather than tamper-proof (`:98`, `:101`).
- §7 now scopes egress to the conductor-governed outbox, not system-level sole egress, and names best-effort content safety plus the D5 network residual (`:118`).
- §11 now centralizes the D5 claim boundary and keeps `record_integrity` as recomputable/detectable, not malicious-containment (`:159`).

The kept classes are also defensible: atomic submit binding and trigger/executor separation are trusted-courier control-flow claims (`:31-37`); the F1 `slot_in` property is observer-selected and no-tool/non-lane-writable with D5 residual (`:98-101`); R2 stays untouched (`:21`, `:141-143`); and the `record_integrity` line is a recomputable pure-function consistency claim scoped to auditability (`:110`, `:159`).

No mechanism change, no c2/c3 contract reopen, no PLAN, no IMPL, no `pcode/`, and no spike are authorized by this review.

## Verification

- `rg -n "\| m-3\.planner \|" master/relays/INDEX.md | tail -n 12` - latest addressed m-3 planner relays were `c5-claim-sweep-light` and `c5-fold-decision-5`.
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c5-claim-sweep-light/DESIGN-planner-20260702-133039.md` - OK.
- `sed -n '1,220p' master/relays/c5-claim-sweep-light/DESIGN-planner-20260702-133039.md` - reviewed full m-3 planner relay.
- `sed -n '1,220p' master/relays/c5-claim-sweep-architecture/DESIGN-orchestrator-planner-20260702-131320.md` - reviewed ratified checklist source.
- `sed -n '1,180p' master/relays/c5-claim-sweep-architecture/RECONCILE-orchestrator-reviewer-20260702-131709.md` - reviewed VP ratification and domain-lane carry-forward.
- `rg -n "non-forgeable|unforgeable|forgery|tamper-proof|tamper-resistant|cannot mark|never lane-supplied|non-lane-writable|D5|claim boundary" master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md master/relays/c5-claim-sweep-architecture` - reviewed broader survivor set.
- `rg -n "governed|sole-writer|observer-selected|confusion-resistant|tamper|tamper-proof|non-lane-writable|immutable|conductor-governed|unbypassable|sole external|structural chokepoint|D5|record_integrity|detectable|unforgeable|forg" master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md` - checked m-3 current relabels and remaining strong claims.
- `git -C pcode status --short` - clean.
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no design-doc edit, no source/code/`pcode` edit, no PLAN, no IMPL, no runtime spike.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (docs workspace; `git status --short` exits 128)
RELAY_LINT: OK - `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c5-claim-sweep-light/DESIGN-REVIEW-implementer-20260702-134833.md`
DISPATCH_ROOT_LINT: OK - `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-claim-sweep-light`
Next requested action: m-3.planner patches the decision-5 survivor wording/classification, reruns the c5 claim-sweep grep against the current document, and re-relays the bounded revision for re-review.
