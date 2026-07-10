## RECONCILE -- VP review: CONCUR with NO-GO/retraction; re-baseline required, with one R2 precision correction

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: design-review
PARENT_DISPATCH_ID: design-review
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-1.implementer, m-2.implementer, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer

Partner -- I concur with the NO-GO and with retracting the Step-1 gate certification. I do **not** concur with using
the R2 finding in its strongest literal form without a precision correction.

Verdict: **approve, with one bounded correction before this review inventory becomes downstream dispatch text.**

## Findings

1. **NO-GO is warranted.** The retraction has two independently sufficient grounds:
   - The operator's recorded decision ② says authority/merge A-gate classes fail closed on unobservable/self-reported
     evidence (`master/READINESS-REGISTER.md:340-344`), but the locked m-3 design still says `self_reported` records
     deliver and that delivery never gates on evidence integrity (`master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:63`,
     `:109-111`, `:128-130`). A faithful builder would ship the old universal deliver-and-label rule.
   - Step 1 is explicitly the "Conductor core / automated operator-relay" (`ROADMAP.md:57-65`), but the six-domain
     lock mostly assigns policy mechanisms. The docs name `conductor-core` for pane-spawn and future PLAN only
     (`master/ARCHITECTURE.md:171`, `:378-381`) and do not contain a buildable owner/design for process lifecycle,
     serialization, crash-atomic multi-file commit, config load/integrity, recovery, store genesis, or GC. m-1 owns the
     trusted store/stamper surface, but not the full running-program substrate.

2. **The operator-decision folding failure is real, not bookkeeping.** The dashboard and register now correctly
   mark the previous full-gate certification as retracted. The old certification treated several decision-implied fixes
   as post-gate carries; at least decision ② is behavior-changing for the write path and must be folded before a build
   reads m-3 as source of truth.

3. **Deployment posture must be re-decided or globally relabeled.** Decision ① recorded `ATTACH-FIRST` plus the honest
   "confusion-resistant" claim (`master/READINESS-REGISTER.md:326-337`), while locked architecture/domain text still
   contains stronger structural/by-construction phrasing for identity/store/egress surfaces
   (`master/ARCHITECTURE.md:28-37`, `master/domains/m-1-trust-identity/design/2026-06-28-trust-identity-design.md:76-95`,
   `master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:118-122`). The fix is a
   deployment fork plus a global claim sweep, not a local wording patch.

4. **R2 needs correction in the review wording.** I do **not** read current m-2 as simply saying "`chosen_model` remains
   gate-referenceable." m-2 §17.3 explicitly labels `chosen_model` as model-identity and non-gate-referenceable
   (`master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md:292`), and m-4's lock permits
   the observe layer to read `chosen_model` as payload for bucket-binding while keeping model-derived predicates out of
   gates (`master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md:96-104`).

   The real defect is still pre-build: the m-2 predicate grammar examples exclude `selected_model`/`model_*`
   (`m-2 ...form-schema-design.md:84`, `:91`, `:303`, `:345`) even though the live row field is `chosen_model`, and
   the row-array subfields do not yet have an explicit per-column FieldSpec / opt-in `gate_referenceable: bool` contract
   with negative fixtures. So the safe downstream wording is: **R2 is under-specified / untestable at schema grain and
   must be fixed by positive gate-referenceability plus row-column FieldSpecs**, not "confirmed chosen_model is currently
   gate-referenceable" unless a later machine schema proves that exact leak.

5. **The re-baseline sequence is right.** I concur with:
   1. deployment fork / claim relabeling first;
   2. named conductor-core design-of-record before any Step-1 PLAN;
   3. fold the five operator decisions and stale renames into locked domain docs;
   4. fix the mechanical FATALs, including R2 at schema grain, altitude-B row grain, authority-class fail-closed,
      pure-judgment A-floor/ceremony floor, config integrity, serialization/crash-atomicity, internal-fault disposition,
      away-token sibling burn/restart binding;
   5. then, and only then, open Step-1 PLAN by explicit operator authorization.

## Required Correction

Before dispatching downstream re-baseline work from `DESIGN-REVIEW-2026-07-01.md`, revise the R2 item to avoid
overclaiming the current text. Suggested replacement:

> R2 is not mechanically proved at the schema grain: current prose excludes `selected_model`/`model_*`, while the live
> row field is `chosen_model`; `chosen_model` is described as non-gate-referenceable, but row-array columns lack an
> explicit opt-in `gate_referenceable: bool` FieldSpec and negative fixtures. Fix with positive gate-referenceability
> defaults, per-column row schemas, and fixtures that reject any gate predicate over `chosen_model`, model identity, or
> single-family bucket proxies.

This correction does not weaken the NO-GO. It makes the re-baseline task sharper and prevents the team from arguing
against a straw version of its own current docs.

## Not Authorized

No Step-1 PLAN, code, `pcode/` work, spike, pair dispatch, or domain-design edit is authorized by this concurrence.
The operator still owns the re-baseline authorization and the deployment-fork grill/decision.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/design-review/RECONCILE-orchestrator-planner-20260701-141116.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/design-review/RECONCILE-orchestrator-reviewer-20260701-141436.md` -- OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/design-review` -- OK
- `git -C pcode status --short` -- clean
- `git status --short` -- unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: read-only review of `master/relays/design-review/RECONCILE-orchestrator-planner-20260701-141116.md`, `master/DESIGN-REVIEW-2026-07-01.md`, `master/README.md`, `master/READINESS-REGISTER.md`, `ROADMAP.md`, `master/ARCHITECTURE.md`, and m-1/m-2/m-3/m-4 domain design source; wrote this reviewer relay and appended `master/relays/INDEX.md`; no domain-design edits, no code/source/pcode, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: pcode clean; cwd status unavailable (not a git repo)
Next requested action: operator decides whether to authorize the re-baseline and grill the deployment fork; planner should correct the R2 wording before using the adversarial review inventory as downstream task text.
