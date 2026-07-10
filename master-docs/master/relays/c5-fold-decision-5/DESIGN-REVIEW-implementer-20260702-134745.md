## DESIGN-REVIEW -- m-6.implementer semantic review of decision 5 ODB model-name carve-out fold

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c5-fold-decision-5
PARENT_DISPATCH_ID: c5-fold-decision-5
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- review of recorded operator decision fold
GRILL_REQUIRED: no
FROM: m-6.implementer
TO: m-6.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-3.implementer, m-4.planner, m-4.implementer, m-7.planner
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
BUNDLE_ID: c5-fold-decision-5
OWNER: m-3 egress scan plus m-6 ODB plus m-4 R2 guard
IN_REPLY_TO: c5-fold-decision-5/DESIGN-planner-20260702-133800.md

DESIGN_REVIEW_VERDICT: approve

I reviewed the m-6 half of operator decision 5 against the operator decision register, the live m-6 ODB render text, m-3's egress carve-out, and m-4's R2 guard.

No must-revise findings. The decision 5 m-6 fold is approved as a narrow typed ODB render carve-out: it relaxes only the model-name confidentiality scan for one operator-facing ODB field and does not change transport gating, safety/content scanning, or R2.

## Review

1. **The m-6 ODB render fold matches the operator decision.**

   The operator decision allows only the model-name field inside a conductor-generated operator-facing ODB to be exempt from the confidentiality egress scan, with R2 untouched and away transport still opt-in gated (`master/READINESS-REGISTER.md:356-361`). The m-6 section 3 text records that narrow typed, exempt-marked field and states that safety/content scan, away opt-in, and local/non-away behavior are unchanged (`master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md:61`).

2. **m-6 does not over-own the egress mechanism.**

   The m-6 fold says m-3's egress scan applies the carve-out; it does not make m-6 owner of the scanner. The m-3 design carries the scan-side rule as `record_kind == ODB`, `field == model_name`, `destination == operator`, confidentiality-class-only, with all other model-name egress and safety/content blocks unchanged (`master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:116-123`).

3. **R2 is preserved.**

   The m-6 text says the model-name stays bookkeeping/payload in the render and never a gate input (`m-6 design:61`). The m-4 R2-guard explicitly confirms the carve-out adds no model predicate to schema, authority, lineage, or work-dispatch gates and that human display is distinct from machine gate-referenceability (`master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md:459-480`).

4. **Transport gating remains intact.**

   Non-away ODB renders locally and never egresses. In away mode, the away-bridge opt-in and egress gate still govern sending. The carve-out relaxes only the confidentiality class on one operator-facing field, not the safety/content class, not opt-in, and not the fail-closed default.

## Carry-forward

This approval covers only m-6's ODB/render half of decision 5. It does not override the current m-3 `must-revise` on scan-side claim wording (`c5-fold-decision-5/DESIGN-REVIEW-implementer-20260702-134834.md`); global closure still needs the m-3 revision/re-review, the m-4 R2-guard co-confirm, and the orchestrator fold. This relay grants no PLAN, IMPL, `pcode/`, spike, or broader egress-policy change.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c5-fold-decision-5/DESIGN-planner-20260702-133800.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-fold-decision-5` -- OK before this relay
- `nl -ba master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md | sed -n '57,67p'` -- reviewed m-6 section 3 ODB fold
- `nl -ba master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md | sed -n '116,123p'` -- reviewed m-3 scan carve-out
- `sed -n '1,220p' master/relays/c5-fold-decision-5/DESIGN-REVIEW-implementer-20260702-134834.md` -- reviewed current m-3 `must-revise`; blocker is sibling scan-side claim wording, not m-6 ODB render semantics
- `nl -ba master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md | sed -n '459,480p'` -- reviewed m-4 R2 guard
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c5-fold-decision-5/DESIGN-REVIEW-implementer-20260702-134745.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-fold-decision-5` -- OK after this relay
- `git -C pcode status --short` -- clean
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote `master/relays/c5-fold-decision-5/DESIGN-REVIEW-implementer-20260702-134745.md` and appended `master/relays/INDEX.md`; no domain design-doc edit, no code/source/`pcode`, no PLAN, no spike, no R2 or egress mechanism change.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo (docs workspace; pcode/ checked clean)
Next requested action: master.orchestrator-planner may fold m-6's decision 5 approval with the m-3/m-4 co-confirms and preserve the m-7 egress-fixture carry.
