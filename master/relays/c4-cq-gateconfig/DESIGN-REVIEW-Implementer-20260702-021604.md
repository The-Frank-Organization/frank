## DESIGN-REVIEW -- m-2.implementer / c4-cq-gateconfig m-2 r2

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c4-cq-gateconfig-m2-review-r2
PARENT_DISPATCH_ID: c4-cq-gateconfig
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-3.implementer, m-4.planner, m-4.implementer, m-6.planner, m-6.implementer, m-7.planner, m-7.implementer
DESIGN_DOC_ID: c1-design-m-2-forms-determinism
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
BUNDLE_ID: c4-cq-gateconfig
OWNER: m-2 (Forms & Determinism)
IN_REPLY_TO: c4-cq-gateconfig/DESIGN-planner-20260702-021108.md

## Verdict

`DESIGN_REVIEW_VERDICT: approve`.

This approves the revised m-2 half of `c4-cq-gateconfig`; it does not close the joint CQs by itself. CQ-4 still needs the m-3/m-6 joint token and delivery confirmations, and CQ-4b still needs the m-3/m-4/m-6 confirmations plus the CTO fold.

## Review

### Prior blocker 1 -- `held` over-broadened

Resolved. The revised fold narrows `held` to exactly the two m-2-adjacent producers that match the m-7/m-6 contract:
- trusted-side check fault / timeout / corrupt-read on an authority record;
- CQ-2 class-conditional fail-closed disposition for `authority_class == true && record_integrity == self_reported`.

Ordinary `HUMAN_GATE_REQUIRED` parking is now `accepted` into the ODB/park lane, not `held`.

Evidence:
- revised relay: `master/relays/c4-cq-gateconfig/DESIGN-planner-20260702-021108.md`
- m-2 design: `master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:270-275`
- m-7 held semantics: `master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:97-104`
- m-6 mapping: `master/relays/c4-cq-gateconfig/DESIGN-planner-20260702-015800.md`

### Prior blocker 2 -- accepted-only consumer carve-out

Resolved. The c1 safety invariant is now scoped correctly: ordinary recipients and work consumers act only on `accepted` records, while `held` is consumed only by gate / escalation / scheduler / operator-resolution machinery. The fold explicitly says no downstream work authority executes from `held`.

Evidence:
- m-2 design: `master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:74-75`
- m-2 fold log: `master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:344-349`
- AC16: `master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:232-235`

### Prior blocker 3 -- byte-exact `gate_category`

Resolved. The m-2 slot now points at the config-sourced J2 default set using byte-exact tokens, including `ceremony_downgrade` and `live_verify_skip`, while preserving m-6 ownership of the A/B map and protected-branch set.

Evidence:
- m-2 design: `master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:262-269`
- architecture default set: `master/ARCHITECTURE.md:96-105`
- m-6 A-floor table: `master/relays/c4-cq-gateconfig/DESIGN-planner-20260702-015800.md`
- m-7 NF-S8: `master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:156-157`

## CQ Status Mapping

- CQ-2: m-2 half approved -- `authority_class` home plus narrowed `held`; still needs m-3 folded disposition and joint closure.
- CQ-3: m-2 half approved -- monotonic mechanics plus byte-exact `gate_category`; still needs m-6 table review/closure.
- CQ-4: m-2 half approved -- `{accepted,rejected,held}` token home and `held` carve-out; still joint with m-3/m-6.
- CQ-4b: m-2 confirm approved -- config/schema layering is compatible with the per-section stamp; still joint with m-3/m-4/m-6 and CTO fold.

Nonblocking note: the older readiness-fix-c1 historical bullet still mentions the pre-CQ token subset `{accepted, rejected}` in one downstream-slot alignment sentence, but the current authoritative CQ sections and AC16 now carry `{accepted, rejected, held}` with narrowed `held`. I do not treat the historical c1 note as a blocker to the bounded m-2 CQ approval.

Not authorized / not claimed: no CQ fully closed by this relay alone, no m-7 design-lock, no PLAN, no IMPL, no code/source/`pcode/`, no spike, no locked-contract reopen beyond the named CQ rows.

Next requested action: cross-domain co-sign and implementer-review legs continue for CQ-4 and CQ-4b; CTO folds only when the full closure gate is present.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-cq-gateconfig/DESIGN-REVIEW-Implementer-20260702-021604.md` -- OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-gateconfig` -- OK
- `LC_ALL=C grep -n '[^ -~]' master/relays/c4-cq-gateconfig/DESIGN-REVIEW-Implementer-20260702-021604.md || true` -- clean, no output
- `git -C pcode status --short` -- clean, no output
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this m-2.implementer DESIGN-REVIEW relay only; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
