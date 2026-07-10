## RECONCILE -- VP closure co-sign: Cluster 4 CLOSED; Step-1 PLAN still operator-gated

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: readiness-fix-c4
PARENT_DISPATCH_ID: readiness-fix-c4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-2.implementer, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer

Partner -- I co-sign Cluster 4 closure.

Verdict: **APPROVE / CLOSED.** The prior VP finding was correct to block closure while 4b/4c were outside the dispatch,
but the current `013231` re-verify now covers the full Cluster 4 surface: 4a grammar/R2, 4b stale routing FieldSpec, and
4c consumed computed-field homes.

Findings:

1. **Owner confirms are sufficient for closure.** I re-read the current owner-confirm set routed through
   `readiness-fix-c4`: m-3 confirms `record_integrity` is authorship-faithful and satisfies the "m-2-declared set";
   m-4 confirms the m-2 mirror now follows the locked routing record without extending `:203`; m-5 confirms option 1
   and does not request a c3-lock amendment; m-6 confirms `(posture x surface_intent)` reads cleanly through the recorded
   `seat_archetype`. No owner confirm leaves a remaining Cluster 4 blocker.

2. **m-2 source now matches the m-4 contract.** The folded m-2 §17.3 row at
   `master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md:292` lists the routing assignment
   fields as `seat`, `role`, `task_tag`, `declared_bucket`, `chosen_model`, `pin_mode`, `declared_deviated`,
   `seat_archetype`, and `authority_ceiling`; it also states that there is no standalone posture field. This matches
   m-4 `master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md:203`. The required
   `deviation_reason_code`, reserved `constraints`, and reserved `template_ref` are present in m-2 §17.3.

3. **4c is adequately declared.** m-2 §17.6 declares homes for `record_integrity`, `surface_intent`, and the m-5-owned
   posture/human-mode vocabulary. The important correction is preserved: posture is delivery vocabulary derived from the
   recorded `seat_archetype`, not an authored `routing_assignments` field. This satisfies the consumed-field home issue
   without reopening c3 or importing tag-value semantics into m-2.

4. **The two follow-ups remain real, but neither blocks Cluster 4 closure.** Row-level parity for the seven remaining
   m-3 observe fields is accepted as §15 Q-F, a pre-Step-1-PLAN SHOULD so fixtures can become machine-checkable. Runtime
   `away` is a legitimate m-5/m-6/m-1 design item because operator availability is runtime/global, not a frozen per-seat
   spawn default. It must be routed before the away-bridge PLAN / Step-2 work, but it does not reopen Cluster 4.

5. **Gate implication is narrow.** With Cluster 1 already VP-closed and the operator decisions recorded, this co-sign
   satisfies the current MUST-before-Step-1 gate for an operator phase decision. It does **not** authorize Step-1 PLAN,
   code, `pcode/` work, a spike, or a c3-lock amendment. The operator still owns any transition out of the current
   AUDIT/DESIGN charter boundary.

Closure record:

- Cluster 4a: CLOSED -- R2 trigger + generic model-atom hole resolved and previously VP-approved.
- Cluster 4b: CLOSED -- m-2 routing FieldSpec now mirrors m-4 `:200-210` / `:203`.
- Cluster 4c: CLOSED -- consumed computed-field homes declared, posture standalone field removed.
- Pre-PLAN carry: KEEP -- row-parity §15 Q-F before Step-1 PLAN if fixtures need the formal rows.
- Step-2 / away carry: KEEP -- runtime-`away` state home and vocabulary split before away-bridge PLAN.

Not authorized / not claimed: no unilateral Step-1 PLAN; no implementation; no code or docs under `pcode/`; no c3-lock
amendment; no deletion of the two follow-up items; no broad update to ARCHITECTURE beyond what the operator later routes.

Verification:

- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/readiness-fix-c4/RECONCILE-orchestrator-planner-20260701-013231.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/readiness-fix-c4/SITREP-planner-20260701-012000.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/readiness-fix-c4/DESIGN-planner-20260701-013000.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/readiness-fix-c4/DESIGN-planner-20260701-012910.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/readiness-fix-c4/SITREP-planner-20260701-012906.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/readiness-fix-c4/RECONCILE-orchestrator-reviewer-20260701-013613.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/readiness-fix-c4` -- OK
- `git -C pcode status --short` -- clean
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: read-only review of planner relay `readiness-fix-c4/RECONCILE-orchestrator-planner-20260701-013231.md`, owner-confirm relays, m-2 folded design §17.3/§17.6, m-4 routing record §5, and `master/READINESS-REGISTER.md`; wrote this reviewer relay and appended `master/relays/INDEX.md`; no design-doc edits, no code/source/pcode, no PLAN.
FINAL_GIT_STATUS_SHORT: pcode clean; cwd status unavailable (not a git repo)
Next requested action: operator may treat Cluster 4 as VP-closed for the Step-1 phase decision; Step-1 PLAN still requires explicit operator authorization, and the row-parity / runtime-away follow-ups remain routed carry items.
