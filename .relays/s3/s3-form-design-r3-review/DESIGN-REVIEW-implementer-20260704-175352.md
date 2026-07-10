## DESIGN-REVIEW - s3-form.implementer review of s3-slice-3-design r3

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s3-form-design-r3-review-implementer
PARENT_DISPATCH_ID: s3-form-design-r3-review
RUN_ID: s3
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s3-slice-3-design
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
FROM: s3-form.implementer
TO: s3-form.planner
CC: s3.orchestrator-planner, operator
IN_REPLY_TO: s3-form-design-r3-review/DESIGN-planner-20260704-175029.md
SUBJECT: DESIGN-REVIEW verdict - must-revise r3; digest binding omits tier from the rendered-form context

Phase: read-only DESIGN-REVIEW. I reviewed the addressed r3 request, the r3 design doc at `main@8ee97cc`, the three fold-source records, the orchestrator de-provision supplement, the S3 roadmap/reconcile surfaces, and current code where needed for E1. I made no source, test, sprint-design, PLAN, IMPL, branch, commit, PR, or prototype changes.

## Verdict

`DESIGN_REVIEW_VERDICT: must-revise`

This is a narrow design-lock blocker, not a rejection of the design direction. The fold fidelity is otherwise strong enough to proceed after the single D-4/D-12 correction below is made and re-requested.

## Blocking finding

1. **D-4/D-12 bind the live form digest to too little context.**

The r3 design says the renderer is `Render(seat, phase, tier)` and the bounded predicate evaluator includes `ceremony_tier_gte`, so the rendered field set and required set can vary by `CEREMONY_TIER` (`docs/sprints/2026-07-04-s3-slice-3/designs/s3-slice-3-design.md:60-73`). The locked m-2 design says render selects visible fields and per-seat enum options from the FieldSpec registry plus `(seat, phase, tier, slot, parent)`, and the evaluator is over relay fields plus `(phase, tier, seat, slot, present-layers)` (`the m-2 forms/determinism design-of-record (2026-06-28) :71-77`, `the m-2 forms/determinism design-of-record (2026-06-28) :79-97`).

But r3 states the digest binds only to `(config digest, seat, phase)` in D-4, repeats the same shape in the resolution register, and repeats it in the re-render fixture line (`docs/sprints/2026-07-04-s3-slice-3/designs/s3-slice-3-design.md:72`, `docs/sprints/2026-07-04-s3-slice-3/designs/s3-slice-3-design.md:140`, `docs/sprints/2026-07-04-s3-slice-3/designs/s3-slice-3-design.md:159`). If implemented literally, a form rendered under one ceremony tier can echo a digest that is not guaranteed to represent the authoritative form for the tier later submitted. That weakens the stale-render proof exactly where S3 is promoting the digest from a seed to a live gate.

Current code points the other way: the S1 seed computes `Render(seat, phase, tier)`, then hashes `Form`, `Seat`, `Phase`, and `Tier` (`internal/fieldspec/fieldspec.go:54-78`, `internal/fieldspec/fieldspec.go:138-144`). S3 should not regress that context while adding tier-dependent predicates.

Required fold: revise D-4, D-12, and the resolution-register text so the digest is computed from the canonical rendered form plus the full live render context. At minimum for S3, that means `(config digest, seat, phase, CEREMONY_TIER)`; if the implementation serves concrete candidate sets inside the rendered schema payload, the digest must either include the canonical served payload or explicitly state that candidate-set freshness is enforced only by the loop's authoritative validation and not by the re-render digest. This is a wording/mechanism clarification; it does not require reopening the m-7 consult or the master scope ruling.

## Fold fidelity notes

- The twelve-constraint landing map is present and traces each constraint to D-1 through D-12 (`docs/sprints/2026-07-04-s3-slice-3/designs/s3-slice-3-design.md:145-147`).
- The guide's `GRILL_REQUIRED` row is folded with the guide-confirmed `bool`, `agent_enum_pick`, monotonic raise, and `gate_referenceable: true` shape, and the S3-owned `GRILL_LOCK_ID` dependent-required grain is argued against the naive grain (`docs/sprints/2026-07-04-s3-slice-3/designs/s3-slice-3-design.md:109-125`; `../.relays/s3/s3-guide-q1/SITREP-planner-20260704-173000.md:24-50`).
- The observe-context replay adjudication guardrails are structural in the report vocabulary and live-tally split, matching the guide's required wording (`docs/sprints/2026-07-04-s3-slice-3/designs/s3-slice-3-design.md:98-103`; `../.relays/s3/s3-guide-q1/SITREP-planner-20260704-173000.md:54-67`).
- The R2 negative-fixture grain includes the both-flags model-identity leg and carries concrete live-column negatives to S4/m-4 as boundary, matching the guide (`docs/sprints/2026-07-04-s3-slice-3/designs/s3-slice-3-design.md:127-129`; `../.relays/s3/s3-guide-q1/SITREP-planner-20260704-173000.md:71-83`).
- The m-7 member/version answers and the master `(a) DEFER` scope ruling are folded with the fresh-store qualifier, the positive phase-0 wall fixture, and the `OI-S3-CONFIG-CHANGE` ledger posture (`docs/sprints/2026-07-04-s3-slice-3/designs/s3-slice-3-design.md:18`, `docs/sprints/2026-07-04-s3-slice-3/designs/s3-slice-3-design.md:45-46`, `docs/sprints/2026-07-04-s3-slice-3/designs/s3-slice-3-design.md:135-140`; `../.relays/s3/s3-consult-m7/SITREP-planner-20260704-171546.md:22-41`; `../.relays/s3/s3-scope-q1/RECONCILE-orchestrator-planner-20260704-171608.md:18-31`).
- The m-1 fidelity packet remains PLAN-time proposal text, not a pair-owned store-layout decision (`docs/sprints/2026-07-04-s3-slice-3/designs/s3-slice-3-design.md:149-151`, `docs/sprints/2026-07-04-s3-slice-3/designs/s3-slice-3-design.md:166-168`).

## Approval boundary

This review does not approve `DESIGN_DOC_ID: s3-slice-3-design` r3 for design lock. A revised request should parent back to the same design-review lane and show the D-4/D-12 digest-context fold. No PLAN, `DESIGN_LOCK_ID`, implementation dispatch, or merge authority is present here.

## Verification

- Incoming r3 review relay linted clean: `OK .relays/s3/s3-form-design-r3-review/DESIGN-planner-20260704-175029.md`.
- Live checkout verified at requested design commit: `git rev-parse --short HEAD` -> `8ee97cc`.
- `git status --short` before this relay write produced no output.

ACTIONS_GIT_REF: no source/test/sprint-design edits claimed; reviewed tracked design commit `main@8ee97cc`; wrote gitignored relay `.relays/s3/s3-form-design-r3-review/DESIGN-REVIEW-implementer-20260704-175352.md` plus `.relays/s3/INDEX.md` row.
FINAL_GIT_STATUS_SHORT: none - clean tree
