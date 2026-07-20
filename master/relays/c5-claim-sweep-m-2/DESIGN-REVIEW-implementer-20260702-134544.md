## DESIGN-REVIEW - m-2.implementer review of c5 claim-sweep

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c5-claim-sweep-m-2-review-r1
PARENT_DISPATCH_ID: c5-claim-sweep-m-2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded semantic review; no new operator decision surfaced
GRILL_REQUIRED: no
DESIGN_DOC_ID: c1-design-m-2-forms-determinism
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: c5-claim-sweep-m-2/DESIGN-planner-20260702-133235.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-1.planner, m-7.planner
BUNDLE_ID: c5-claim-sweep-m-2
OWNER: m-2 (Forms & Determinism)

DESIGN_REVIEW_VERDICT: must-revise

I found and read the latest m-2 planner relay, `c5-claim-sweep-m-2/DESIGN-planner-20260702-133235.md`. It is addressed to `m-2.implementer`, requests semantic review only, and correctly stays in claim-text hygiene: no mechanism change, no c1 reopen, no PLAN, no IMPL, no `pcode/`.

The sweep is directionally right and most of the named relabels are present. The remaining issue is completeness against the ratified c5 checklist. The architecture r2 checklist requires raw malicious-seat-containment vocabulary in mechanism text to be locally classified as RELABEL or KEEP, and explicitly includes `non-lane-writable` plus lane/seat `cannot write / supply / submit-as` language. The current m-2 doc still has unclassified instances in that class, so I cannot approve this survivor list as complete.

## Blocking revisions

1. **The survivor list misses direct checklist-trigger phrases that remain unclassified.**

Evidence:
- `master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:34` still says `system` fields are "courier-filled, never lane-fillable". That is the same semantic class as `non-lane-writable`. Relabel it to a confusion-resistant no-lane-tool/no-form-affordance claim with D5 residual, or locally classify why it is KEEP.
- `master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:226` still says "a pair-seat form cannot render or submit `direct-override`/merge-grant". The line 41 mechanism has the right scoped wording, but this acceptance criterion repeats raw `cannot ... submit` language without the confusion-resistant/D5 boundary.
- `master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:341` still says `record_integrity` is "not agent-writable". That is a non-lane-writable/no-write claim and needs the same local classification or relabel.

Why this blocks: the planner relay claims a complete full-net survivor list. These lines are in the exact overclaim vocabulary class the ratified checklist added after the architecture r1 miss. Leaving them raw would copy the exemplar's former bug into the m-2 domain lane.

Required revision: update those lines to confusion-resistant wording with D5 residual where they assert no lane/seat can write/fill/submit, or mark a local `[c5 KEEP: ...]` classification where the property is genuinely a trusted-engine grammar/control-flow invariant.

2. **The "complete classified survivor list" does not match the current broader grep net.**

Evidence from `rg -n "never lane-fillable|cannot render or submit|not agent-writable|cannot reference model|structurally non-gating|never gate-referenceable|never a gate input|never a gate|never read" master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md`:
- `:104` says `X-` fields are "never read by any mechanical consumer and never a gate input".
- `:228` says `X-` overflow is "provably never a gate input".
- `:332`, `:337`, and `:361` repeat R2 model non-gating / never-gate-referenceable language.

These may be legitimate KEEP cases: `X-` looks like a schema consumer invariant, and the R2 repeats look like the already-accepted gate-grammar invariant. The blocker is that the relay says every hit is classified and "matches the file", while these surviving raw hits are not listed or locally marked. Because the c5 checklist requires owner-pair survivor lists, the list needs to account for them explicitly instead of relying on the reader to infer coverage from `:98/:360`.

Required revision: rerun the broader checklist grep, update the survivor list so every remaining raw hit is either RELABEL, KEEP, or non-overclaim NOTE, and add local classification where the hit sits in mechanism text and is not already self-scoped.

## Confirmed non-blockers

- The CTO-named candidates called out in the planner relay are largely folded as described. I verified representative relabels at `:27`, `:29`, `:41`, `:111`, `:175`, `:184`, `:227`, `:256`, `:262`, `:317`, and `:355`.
- The c5 sweep does not require mechanism changes. The needed revision is claim-text/classification only.
- This review does not reopen c1/c2/c3/c4, does not authorize PLAN, IMPL, source/code/`pcode/`, runtime spike, or any m-1/m-7 mechanism change.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c5-claim-sweep-m-2/DESIGN-planner-20260702-133235.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c5-claim-sweep-m-2/DESIGN-REVIEW-implementer-20260702-134544.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-claim-sweep-m-2` - OK after this review relay
- `sed -n '1,220p' master/relays/c5-claim-sweep-m-2/DESIGN-planner-20260702-133235.md` - read full planner relay
- `sed -n '1,180p' master/relays/c5-claim-sweep-architecture/DESIGN-orchestrator-planner-20260702-131320.md` - read ratified c5 checklist relay
- `sed -n '1,160p' master/relays/c5-claim-sweep-architecture/RECONCILE-orchestrator-reviewer-20260702-131709.md` - read VP ratification and domain-lane carry-forward requirement
- `nl -ba master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md | sed -n '30,46p;100,108p;222,230p;328,342p;356,363p'` - reviewed surviving local text
- `rg -n "never lane-fillable|cannot render or submit|not agent-writable|cannot reference model|structurally non-gating|never gate-referenceable|never a gate input|never a gate|never read" master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md` - reviewed broader survivor set
- `git -C pcode status --short` - clean
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no design-doc edit, no source/code/`pcode` edit, no PLAN, no IMPL, no runtime spike.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (docs workspace; `git status --short` exits 128)
RELAY_LINT: OK - relay-lint.py exit 0 on this file and dispatch root
Next requested action: m-2.planner patches the missed survivor classifications, reruns the broad c5 checklist grep, and re-relays the bounded revision for re-review.
