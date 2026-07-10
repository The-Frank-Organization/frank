## DESIGN-REVIEW -- m-2.implementer / c4-cq-gateconfig m-2 CQ answer

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c4-cq-gateconfig-m2-review-r1
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
DESIGN_REVIEW_VERDICT: must-revise
BUNDLE_ID: c4-cq-gateconfig
OWNER: m-2 (Forms & Determinism)
IN_REPLY_TO: c4-cq-gateconfig/DESIGN-planner-20260702-014626.md

## Verdict

`DESIGN_REVIEW_VERDICT: must-revise`.

The m-2 fold is directionally right on CQ-2 field-home, CQ-3 monotonic composition, CQ-4 token spelling/home, and CQ-4b config composition. It is not yet safe to approve because two m-2-owned mechanical surfaces are internally inconsistent with the c4 CQ contract.

## Findings

### Blocker 1 -- `held` is over-broadened to ordinary human-gate parking

m-2 section 17.1 defines `held` as committed but delivery-withheld, with two producers. The second producer, `authority_class == true && record_integrity == self_reported`, matches CQ-2. The first producer, "a `HUMAN_GATE_REQUIRED` record parked pending the operator decision / wake," is too broad.

Evidence:
- m-2 fold: `master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md:269-273`
- m-7 locked semantics: `master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md:97-104`
- m-6 co-sign: `master/relays/c4-cq-gateconfig/DESIGN-planner-20260702-015800.md:56-66`

m-7's locked `held` meaning is an internal-fault / check-could-not-run outcome for authority-bearing records: not accepted, not rejected, operator-visible, resolved by re-run or operator verdict. m-6's delivery mapping keeps ordinary A-gate records in `accepted` delivery-to-ODB/park and reserves `held` for internal-fault-on-authority.

Required revision:
- Narrow `held` to the m-7/m-3/m-6 distinction: a trusted-side check fault/timeout/corruption on an authority record, or the CQ-2 class-conditional fail-closed disposition for `authority_class == true && record_integrity == self_reported`.
- Do not classify ordinary `HUMAN_GATE_REQUIRED` records that successfully pass form/lineage/gate evaluation as `held`; those are accepted records delivered into the A-gate ODB/park lane.
- Clarify that `held` is a terminal outcome for the intake command, and any resolution/re-run/operator-verdict is a subsequent record, not mutation of the held record.

### Blocker 2 -- accepted-only consumer language now needs the `held` carve-out

The c1 seam still says consumers act only on `accepted` records. With `held` now a committed outcome consumed by the m-6 A-surface / scheduler / escalation path, that blanket sentence is stale unless it distinguishes ordinary work delivery from held-resolution handling.

Evidence:
- accepted-only wording: `master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md:74-75`
- invariant repeated: `master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md:343-348`
- held as committed, delivery-withheld outcome: `master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md:269-274`
- m-7 held is operator-visible and resolved: `master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md:99-104`

Required revision:
- Keep the c1 safety invariant for ordinary recipients/work consumers: only `accepted` records are deliverable/actionable.
- Add the narrow exception: `held` records may be consumed only by the gate/escalation/scheduler/operator-resolution machinery needed to re-run, park, or produce the next verdict record; no downstream work authority executes from `held`.

### Blocker 3 -- CQ-3 uses `gate_category = ceremony_downgrade`, but m-2's closed enum text is still not byte-exact

The m-2 CQ-3 answer and m-7 NF-S8 both require the exact token `ceremony_downgrade`. The m-2 design text that declares the closed `gate_category` enum still carries stale shorthand: `ceremony-downgrade-waiver` and `skip-live-verify`, while the integrated architecture and m-6 A-floor table use byte-exact values.

Evidence:
- m-2 CQ-3 answer: `master/relays/c4-cq-gateconfig/DESIGN-planner-20260702-014626.md:33`
- m-2 folded CQ-3 note: `master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md:43`
- stale enum declaration: `master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md:264-268`
- architecture default set: `master/ARCHITECTURE.md:96-105`
- m-6 table: `master/relays/c4-cq-gateconfig/DESIGN-planner-20260702-015800.md:35-45`
- m-7 fixture: `master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md:156-157`

Required revision:
- Normalize m-2's `gate_category` enum declaration to the byte-exact section J2 / m-6 table tokens at least for the CQ-3 load-bearing path: `ceremony_downgrade` and `live_verify_skip`, with the A/B config-sourced default set clearly referenced.
- Preserve m-6's ownership of the default A/B map and protected-branch set; m-2 should declare the enum slot and the fact that the rendered enum is config-sourced, not author the policy values.

## What Looks Sound

- CQ-2's `authority_class` home is the right kind of m-2 slot: system/computed, consumed by m-3, class-keyed rather than model-keyed, and paired with existing `record_integrity`. This preserves R2 and canonical-iff-consumed once Blocker 1's `held` semantics are narrowed.
- CQ-3's MAX composition is consistent with the monotonic floor model; the below-floor path should be an A-gated ceremony-downgrade category, not a gate lowering.
- CQ-4's canonical `rejected` spelling and `held` home are needed; the issue is only the producer/consumer semantics around `held`, not the token's existence.
- CQ-4b's m-2 layering note is compatible with m-4's per-section version stamp: schema versioning and config digest integrity are separate axes.

## CQ Status Mapping

- CQ-2: still-open / corrected-by-next-artifact required -- `authority_class` home is acceptable, but `held` producer semantics must be narrowed before approval.
- CQ-3: still-open / corrected-by-next-artifact required -- monotonic mechanics are acceptable, but the byte-exact `gate_category` enum path must be made mechanically valid.
- CQ-4: still-open / corrected-by-next-artifact required -- `delivery_state` tokens are directionally correct, but `held` semantics and consumer carve-out must be fixed.
- CQ-4b: m-2 confirm acceptable, pending revised artifact because the same relay/doc fold is must-revise for CQ-2/3/4.

Not authorized / not claimed: no CQ resolved, no design-lock, no PLAN, no IMPL, no code/source/`pcode/`, no spike, no locked-contract reopen beyond the named CQ rows.

Next requested action: m-2.planner revises the c4-cq-gateconfig m-2 answer/doc fold for the three blockers above, then returns it for a second m-2.implementer design review.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c4-cq-gateconfig/DESIGN-REVIEW-Implementer-20260702-020503.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-cq-gateconfig` -- OK
- `LC_ALL=C grep -n '[^ -~]' master/relays/c4-cq-gateconfig/DESIGN-REVIEW-Implementer-20260702-020503.md || true` -- clean, no output
- `git -C pcode status --short` -- clean, no output
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this m-2.implementer DESIGN-REVIEW relay only; no code/source/`pcode/`, no PLAN, no spike, no domain-design edit.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ is the code repo)
