## DESIGN-REVIEW - m-3.implementer review of CQ-5 m-3 observe-pipeline co-sign

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c4-cq-slotin
PARENT_DISPATCH_ID: c4-cq-slotin
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded CQ-5 observe-pipeline review; no new operator decision surfaced
GRILL_REQUIRED: no
DESIGN_DOC_ID: c2-design-m-3-observation-evidence
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: c2-grill-m-3
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: c4-cq-slotin/DESIGN-planner-20260702-024732.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-5.planner, m-5.implementer, m-7.planner, m-7.implementer, operator
BUNDLE_ID: m-3-observation-evidence

DESIGN_REVIEW_VERDICT: approve

I reviewed the m-3 CQ-5 observe-pipeline co-sign in `c4-cq-slotin/DESIGN-planner-20260702-024732.md` against the live m-3 design doc, the m-5 lead answer and approval, the m-7 proposed commit pipeline, the architecture seam, and the c4 coordination closure gate.

Approve. The co-sign correctly confirms the m-3-owned half: the Step-2 done-predicate reads the just-classified in-courier `slot_in`, so conductor classification must occur after form/lineage validation and before the observe hook, with classification and observation bound in the same atomic submit commit. No correction to m-7 section 3 or NF-S12 is required from the m-3 side.

## Review checks

1. **The "required" claim is defensible.**

The m-3 hook is parameterized by archetype tags and the done-predicate reads the conductor-classified `slot_in` (`master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:43-58`, `:88-103`). The observe hook's closed write-allowlist does not include `slot_in` (`...observe-evidence-design.md:61`). So a separate conductor-owned classification step must make the tag available before predicate selection. A classify-after-observe ordering would leave the predicate with no trustworthy key.

2. **The atomic classify-plus-observe bind matches m-3's TOCTOU closure.**

m-3 locks observe inside the atomic `submit()` so the passing observation is the observation stamped (`...observe-evidence-design.md:31-35`). The live CQ-5 fold pins `slot_in` classification post-form/lineage, pre-observe, and atomic with observation (`...observe-evidence-design.md:101`, `:200`, `:217`). That is the same shape m-7 proposes in its commit loop (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:68-74`) and in NF-S12 (`...conductor-core-design.md:161`).

3. **Post-form/lineage placement does not reopen c1/c2.**

The architecture seam keeps the candidate in-courier through pre-append form-validation and lineage, with no persisted `submitted` limbo (`master/ARCHITECTURE.md:58-66`). CQ-5 places classification after those checks and before observe. That adds no concrete `slot_in` values to m-2 predicates: the architecture still treats `slot_in` as an opaque reserved atom (`master/ARCHITECTURE.md:52-55`, `:187-193`), and the m-3 doc still preserves no `required_when`/`visible_when` branch on concrete slot values (`...observe-evidence-design.md:103`).

4. **Per-record classification and long-lived seats remain intact.**

m-5's locked contract says `slot_in` is conductor-classified at work-record acceptance, non-lane-writable, and per-work-record rather than spawn-fixed (`master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:17-21`, `:52-53`). Its tamper-resistance proof depends on the observer selecting the invariant, not the observed lane (`...archetype-system-design.md:64-78`). The m-5 lead answer and implementer review already approved that half (`master/relays/c4-cq-slotin/DESIGN-planner-20260702-014506.md:23-44`, `master/relays/c4-cq-slotin/DESIGN-REVIEW-implementer-20260702-020448.md:20-44`). The m-3 co-sign complements it rather than re-adjudicating m-5's tag-space.

## CQ-status review

- CQ-5 m-5 lead answer: already closed on the m-5 side.
- CQ-5 m-5 implementer adversarial review: already approved.
- CQ-5 m-3 observe-pipeline co-sign: approved by this review.
- Composite CQ-5: now has the required m-5 answer, m-5 review, m-3 co-sign, and m-3 review legs. It is foldable by `master.orchestrator-planner` under the `c4-cq-coord` closure gate, but this review does not itself fold it, does not design-lock m-7, and does not authorize PLAN, IMPL, `pcode/`, or a spike (`master/relays/c4-cq-coord/DESIGN-orchestrator-planner-20260702-013323.md:32-40`).

No c2/c3 contract reopen, no new archetype/tag-space design, no m-5 half proxied, and no m-7 design-lock by implication.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-cq-slotin/DESIGN-planner-20260702-024732.md` - OK
- `nl -ba master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md | sed -n '29,75p;84,116p;188,220p'` - reviewed m-3 §2, §3.1, §5.1, §13, §15
- `nl -ba master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md | sed -n '1,90p;190,212p'` - reviewed m-5 F1/tag-space/tamper-resistance/PLAN carry
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md | sed -n '60,80p;148,164p;214,224p'` - reviewed m-7 commit pipeline and NF-S12/CQ-5
- `nl -ba master/ARCHITECTURE.md | sed -n '48,68p;184,196p'` - reviewed c1/c2 write-path and opaque slot boundary
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-cq-slotin/DESIGN-REVIEW-implementer-20260702-030205.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-slotin` - OK
- `git -C pcode status --short` - clean
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no design-doc edit by this seat, no code/source/`pcode` edit, no PLAN, no IMPL, no spike.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (docs workspace; final `git status --short` exits 128)
Next requested action: master.orchestrator-planner may fold CQ-5 if it accepts the four addressed legs as complete; m-3.implementer stands down for this lane unless a revised artifact is relayed.
