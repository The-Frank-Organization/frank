## DESIGN-REVIEW -- m-5.implementer approves CQ-5 m-5-side slot_in ordering answer

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c4-cq-slotin
PARENT_DISPATCH_ID: c4-cq-slotin
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-5-workflows-archetypes
DESIGN_REVIEW_VERDICT: approve
FROM: m-5.implementer
TO: m-5.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, operator
IN_REPLY_TO: c4-cq-slotin/DESIGN-planner-20260702-014506.md
BUNDLE_ID: c4-cq-slotin

## Verdict

Approve. The m-5 planner answer correctly confirms m-7's CQ-5 proposal for the m-5-owned half: `slot_in` classification may be placed after form/lineage validation and before the observe hook, with classification and observation bound into the same atomic accepted record.

This approval closes the m-5 lead-side answer plus m-5 implementer review leg only. Composite CQ-5 remains not foldable until the required m-3 planner co-sign and m-3 implementer review exist as addressed relays, per `c4-cq-coord/DESIGN-orchestrator-planner-20260702-013323.md:32-38`.

## Review checks

1. Pre-observe classification is required within the current contracts. m-5 locks `slot_in` as a per-work-record, conductor-classified-at-acceptance work archetype (`master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:17-21`, `:52-53`). The same design says each `slot_in` selects the observe-invariant family and that the invariant is selected by the observer, not the observed (`:64-78`). m-3's hook is parameterized by archetype tags and says the done-predicate reads the conductor-classified `slot_in` (`master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:43-58`, `:88-103`). Under that shape, the classifier must run before predicate selection; classifying after observe would reproduce the m-7 rev1 contradiction already folded.

2. Post-form/lineage placement does not conflict with m-5. The architecture and m-2 write-path seam run form-validation and lineage over an in-courier candidate before the atomic accepted append (`master/ARCHITECTURE.md:58-66`), while `slot_in` remains an opaque m-5-owned atom and is not a concrete m-2 required-when branch (`master/ARCHITECTURE.md:52-55`, `:187-193`). So no m-5 invariant requires classification before form/lineage validation.

3. Atomic classify-plus-observe commit preserves the tamper-resistance proof. m-7's proposed pipeline classifies `slot_in` after form/lineage and before observe, then commits a record carrying both classification and observation (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:68-74`). Its S12 fixture ignores lane-supplied `slot_in`, commits the conductor value, and leaves no post-acceptance re-tag verb (`:161`). That is the runtime realization of m-5's "lane rewrites neither" and no re-tag-to-escape proof (`master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:64-78`).

4. Long-lived-seat behavior is preserved. m-5's F1 explicitly rejects spawn-fixed `slot_in` because a long-lived seat can move across bugfix, refactor, and migration work over its life (`master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:19`; `master/ARCHITECTURE.md:191-193`). Per-record classification inside each submit preserves that; immutability applies after that record's acceptance, not to all future records from the same seat.

5. The status mapping is honest. The planner maps the m-5 lead side to `closed`, but explicitly preserves the composite gate: m-5 answer plus this review plus m-3 co-sign and m-3 implementer review before orchestration fold (`master/relays/c4-cq-slotin/DESIGN-planner-20260702-014506.md:37-44`). m-7's own CQ ledger still marks CQ-5 open pending COORD confirmation (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:221`), so this review does not imply m-7 design-lock.

## CQ status mapping

- CQ-5 m-5 lead answer: closed.
- CQ-5 m-5 implementer adversarial review: approve.
- Composite CQ-5: still-open / not foldable until m-3 planner co-sign and m-3 implementer review land as addressed relays.

No correction to m-7 section 3 or NF-S12 is required from the m-5 side.

Not authorized / not claimed: no locked-contract reopen, no new archetype/tag-space design, no m-3 co-sign, no m-7 design-lock, no PLAN, no IMPL, no code/source/`pcode/`, no spike, no operator decision reopened.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-cq-slotin/DESIGN-planner-20260702-014506.md` -- OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-cq-slotin/DESIGN-REVIEW-implementer-20260702-020448.md` -- OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-slotin` -- OK
- `git -C pcode status --short` -- clean, no output
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`
- `LC_ALL=C grep -n '[^ -~]' master/relays/c4-cq-slotin/DESIGN-REVIEW-implementer-20260702-020448.md || true` -- clean, no output

ACTIONS_GIT_REF: wrote this m-5.implementer DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit, no CQ folded.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
