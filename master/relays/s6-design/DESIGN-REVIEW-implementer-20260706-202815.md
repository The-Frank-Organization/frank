## DESIGN-REVIEW - m-7 B-1 boot-stage addendum must revise: pre-active boot classifier is under-specified

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s6-design
PARENT_DISPATCH_ID: s6-design
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: m-7-s6-transport-amendments
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: s6-design/DESIGN-planner-20260706-202124.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-1.planner, m-2.planner
SUBJECT: must-revise - B-1 needs an exact boot-class predicate or marker before pre-active non-boot rejection is lockable
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)

## Verdict

DESIGN_REVIEW_VERDICT: must-revise

This verdict is scoped only to B-1 in `master/domains/m-7-conductor-core/design/2026-07-06-s6-transport-amendments.md` r3 (§B-1.1-B-1.5). The r2 A/D legs remain approved by `s6-design/DESIGN-REVIEW-implementer-20260706-184027.md`.

The B-1 direction is right, but the out-of-order rule is not yet mechanically tight enough to lock. B-1 says a pre-active seat renders only the boot form and that a hand-crafted non-boot submit receives a typed `boot-required` rejection. With m-2's B-2 choice of `PHASE: SITREP` and no new `BOOT` phase or `record_kind`, B-1 must define exactly how the conductor distinguishes "the boot record" from "an ordinary SITREP carrying the two boot fields plus extra non-boot work."

## Blocking finding

1. **Pre-active boot classification is under-specified.** B-1 keys `active` on the accepted BOOT record and rejects non-boot submits before active (`2026-07-06-s6-transport-amendments.md:96`, `:102-105`). m-2 B-2 deliberately avoids a new phase/record-kind and says the active transition may detect "the seat's first accepted SITREP carrying the boot required-set" (`2026-07-06-s6-transport-codec-amendment.md:152-158`). Current code evidence shows why "carrying the required set" is insufficient as the lock predicate: validation iterates the registry fields and checks required/typed/enum/seat-scope constraints, but it does not reject arbitrary extra hand-crafted headers (`frank/internal/fieldspec/validate.go:31-65`); the submit pipeline accepts once validation, lineage, and record-kind checks pass (`frank/internal/engine/submit.go:47-65`, `:82-87`). Therefore a pre-active candidate shaped as `SITREP + charter_loaded + dispatch_status + extra non-boot/authority field` is not ruled out by the written B-1 classifier. That can turn "contains boot fields" into activation and bypass the intended boot-only ordering.

## Required revision

Add one explicit B-1 rule and fixture family before approval:

- The pre-active boot candidate must be classified by an exact conductor-owned predicate, not by mere presence of the two boot fields. Acceptable shapes include either:
  - **exact rendered boot form:** before `active`, the only acceptable submit candidate is the lifecycle-rendered BOOT form, with no non-rendered/non-boot payload fields beyond the ordinary system envelope and the m-2 B-2 boot field set; or
  - **system-derived boot marker:** after validating the exact rendered boot form, the conductor stamps a system-derived `boot_ack`/equivalent marker for activation detection, with m-2 owning the field shape if needed and no new phase/record-kind.
- The authoritative rejection must name the class for a pre-active candidate that includes boot fields plus any extra non-boot field, for example `boot-required` or `non-boot-before-active`, and it must still use D-2 parity detail.
- Add a negative fixture: pre-active `SITREP` carrying `charter_loaded` + `dispatch_status` plus an extra non-boot/authority-bearing header does **not** become active and produces the typed rejection.
- Add a positive fixture: pre-active exact boot form accepts and flips active exactly once; an already-active boot-shaped SITREP is ordinary and creates no second activation edge.

This revision does not require a new seat verb, a new `BOOT` phase atom, a new `record_kind`, or recorded `bound` transitions. It only closes the classifier hole created by the otherwise-good "SITREP + lifecycle-gated form" decision.

## Grill results on the four rulings

1. **Recording shape:** approve after the blocker above. `bound` as runtime state is defensible: it represents current authenticated-channel liveness, so recording it as canonical history would add a mutation class for state recovery must distrust after a restart. Recovery deriving `minted` and `active` from records while reopening with `bound = empty` is honest and matches m-1's reconnect-as-re-bound semantics (`m-1 ...transport-amendments.md:96-100`).

2. **Out-of-order:** must revise only on the exact classifier. The selected disposition, typed refusal for pre-active non-boot submit, is the right B-1 call and composes with m-1's "activation grants no authority" line. It just needs the exact-field/marker rule above so "non-boot" is mechanically decidable.

3. **Roster:** approve. A `project(view=roster)`-grade projection scoped to operator/orchestrator seats preserves the three-verb surface and makes the scoped cross-seat metadata exception explicit. The fixture should keep the existing negative: a non-privileged seat cannot get roster data. Non-blocking wording cleanup: call the persisted part `activation_state` or equivalent and keep `bound_now` as a separate liveness bit, so an active-but-disconnected seat after restart is not misread as impossible.

4. **Seams:** approve after the blocker above. B-1 is indifferent to m-2's no-new-BOOT-phase branch if it defines the exact boot predicate/marker. It also respects m-1's B-3 edge semantics: m-1 owns what the edges assert; m-7 records/derives runtime state and enforces ordering.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s6-design/DESIGN-planner-20260706-202124.md` - OK.
- `git -C frank status --short && git -C frank rev-parse --short HEAD && git -C frank tag --points-at HEAD` - clean; `7e5c527`; `s5-close`.
- Reviewed `master/relays/s6-design/DESIGN-orchestrator-planner-20260706-200259.md`, `master/relays/s6-design/DESIGN-planner-20260706-202124.md`, `master/domains/m-7-conductor-core/design/2026-07-06-s6-transport-amendments.md:88-124`, m-2 B-2 (`master/domains/m-2-forms-determinism/design/2026-07-06-s6-transport-codec-amendment.md:141-169`), and m-1 B-3 (`master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md:91-103`).
- Reviewed code evidence in `frank/internal/fieldspec/validate.go:31-65`, `frank/internal/engine/submit.go:47-87`, `frank/cmd/frank/main.go:255-304`, and locked m-7 design lines for submit ordering, recovery, three-verb surface, and derived projections.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s6-design/DESIGN-REVIEW-implementer-20260706-202815.md` - OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s6-design master/relays/s6-design/DESIGN-planner-20260706-202124.md master/relays/s6-design/DESIGN-REVIEW-implementer-20260706-202815.md` - OK.
- `rg -n "20260706-202815|DESIGN-REVIEW-implementer-20260706-202815" master/relays/INDEX.md master/relays/s6-design/DESIGN-REVIEW-implementer-20260706-202815.md` - index row present.
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`.

ACTIONS_GIT_REF: wrote `master/relays/s6-design/DESIGN-REVIEW-implementer-20260706-202815.md`; appended `master/relays/INDEX.md`; no code/source/`frank/` edits; no DESIGN_LOCK mutation; no PLAN or IMPL dispatch.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo; `git -C frank status --short` clean at `7e5c527` / `s5-close`.
Next requested action: m-7.planner revises B-1 to add the exact pre-active boot classifier/marker rule and fixtures, then re-relays for a bounded B-1 re-review.
