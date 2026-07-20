## DESIGN-REVIEW - m-2.implementer review of c6.1-confirm-m-2

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c6.1-confirm-m-2-review-r1
PARENT_DISPATCH_ID: c6.1-confirm-m-2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - bounded c6.1 convergence review; no operator decision surfaced
GRILL_REQUIRED: no
DESIGN_DOC_ID: c1-design-m-2-forms-determinism
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: c6.1-confirm-m-2/DESIGN-planner-20260702-235346.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: c6.1-confirm-m-2
OWNER: m-2 (Forms & Determinism)

DESIGN_REVIEW_VERDICT: must-revise

I read `c6.1-confirm-m-2/DESIGN-planner-20260702-235346.md` and reviewed the live m-2 design doc against the three c6.1 deltas and their cited locked targets.

The main c6.1 target deltas are correct in the live §17.1 / §17.3 contract text, but condition (b) is not fully met: one live acceptance-criteria line still carries the pre-c6.1 routing author set. This is a narrow propagation gap, not a mechanism change and not a c6 reopen.

## Blocker

1. **AC12 still says `routing_assignments` is scoped only to planner/orch-planner.**

`master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:237` still reads that the routing decision has `routing_assignments` "seat-scoped to planner/orch-planner". That conflicts with the c6.1 target now present in §12 and §17.3: `operator` is admitted on `template_ref`-bearing template-spawn records only.

Why this blocks: the orchestrator dispatch required confirming that the deltas "introduce no contradiction elsewhere" in the doc. AC12 is a live acceptance criterion, not a historical fold-log. It should be propagated to the same bounded wording as §17.3 / m-4 §7: planner / orch-planner, plus `operator` on `template_ref`-bearing template-spawn records only.

## Accepted Checks

1. **Delta 1 - `held` on-store shape converges.**

`master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:281` now states one compound canonical record whose held-disposition record embeds the candidate and references the `intake_id`. That converges to m-7 §6 at `master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:100`, and m-7 F11 at `:172` enumerates `held (compound, candidate embedded)`.

2. **Delta 2 - template-spawn operator authoring converges in the primary contract text.**

`master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:187` and `:309` carry the bounded `operator` exception for `template_ref`-bearing template-spawn records, and `:311` carries the matching `declared_deviated` operator-on-template-spawn carve-out. That converges to m-4 `:208` and §7 `:266-272`.

3. **Delta 3 - `deviation_reason_code` config-sourced value-set converges.**

`master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:187` and `:314` carry the config-sourced enum shape with default-seeded + operator-configurable + `other` fail-safe. That converges to m-4 `:211` and §7 `:227-228`. m-2 correctly mirrors the shape and flags m-4 as value owner rather than re-authoring the shipped token list.

4. **Lock invariants remain intact except for the AC12 echo gap above.**

The byte-exact `{accepted, rejected, held}` set is still present at `:126`, `:241`, and `:282`; R2 remains grammar-enforced at `:99`, `:239`, `:313`, and `:320`; the c6.1 edits do not authorize PLAN, IMPL, `pcode`, runtime spike, or design-lock reopen.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c6.1-confirm-m-2/DESIGN-planner-20260702-235346.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c6.1-confirm-m-2` - OK before this review relay
- `sed -n '1,180p' master/relays/c6.1-confirm-m-2/DESIGN-orchestrator-planner-20260702-233000.md` - read orchestrator scope
- `sed -n '1,260p' master/relays/c6.1-confirm-m-2/DESIGN-planner-20260702-235346.md` - read planner confirmation
- `nl -ba master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md | sed -n '180,190p'` - reviewed §12 echo
- `nl -ba master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md | sed -n '232,242p'` - found AC12 blocker at `:237`
- `nl -ba master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md | sed -n '276,318p'` - reviewed §17.1 / §17.3 deltas
- `nl -ba master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md | sed -n '203,214p'` and `sed -n '246,272p'` - checked m-4 target text
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md | sed -n '96,103p'` and `sed -n '168,174p'` - checked m-7 target text
- `rg -n "planner/orch-planner|planner / orch-planner|two-record|candidate \\+ separate|deviation_reason_code.*agent_enum_pick|deviation_reason_code.*same grain|deviation_reason_code.*config-sourced|routing_assignments.*seat-scoped" master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md` - surfaced the remaining AC12 stale author-set line and confirmed the other live contract lines
- `git -C pcode status --short` - clean
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c6.1-confirm-m-2/DESIGN-REVIEW-implementer-20260703-000156.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c6.1-confirm-m-2` - OK
- `rg -n "c6.1-confirm-m-2-review-r1|DESIGN-REVIEW-implementer-20260703-000156" master/relays/INDEX.md && tail -n 5 master/relays/INDEX.md` - row present at EOF

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no design-doc edit, no README edit, no source/code/`pcode` edit, no PLAN, no IMPL, no runtime spike.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (docs workspace; `git status --short` exits 128)
RELAY_LINT: OK - relay-lint.py exit 0 on this file and dispatch root
Next requested action: m-2.planner should revise the AC12 echo to match the c6.1 author-set delta, then return a rev2 confirmation for implementer re-review. No m-2 self-advance is authorized by this review.
