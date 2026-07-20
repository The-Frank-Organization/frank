## DESIGN-REVIEW -- m-6.implementer confirm of the s5-escalations known-A signal set

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s5-escalations-m6-signal-confirm
PARENT_DISPATCH_ID: s5-escalations
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- review-only confirm of a locked-design reading
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: s5-escalations/DESIGN-planner-20260706-051357.md
FROM: m-6.implementer
TO: master.orchestrator-planner
CC: operator, s5.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-7.planner, m-2.planner, m-4.planner, m-5.planner, m-1.planner
BUNDLE_ID: s5-escalations
OWNER: m-6 (Human Surface and Scheduler)
SUBJECT: adversarial approve on the m-6 S1/S2/S3 signal-set ruling; confirms the Step-1 claim boundary and does not reopen c3

DESIGN_REVIEW_VERDICT: approve

I reviewed `s5-escalations/DESIGN-planner-20260706-051357.md` plus the later master reconcile request at `s5-escalations/RECONCILE-orchestrator-planner-20260706-052214.md`. The confirm is approvable as a narrow, read-only adversarial check of the m-6 signal-set half of decision-3. It does not approve m-7's mechanics independently, does not close s5 globally, and does not authorize PLAN, IMPL, code edits, lock reopen, or a new predicate engine.

Routing note: the original planner relay CC'd `m-6.implementer`, and the master reconcile accepted the offered confirm but also left this seat in CC. I am treating the operator hand-relay of the m-6 planner artifact plus the master reconcile's explicit "m-6.implementer's signal-set confirm - REQUESTED" language as the action request, and I am answering TO master as that reconcile directs.

## Review

1. S1 is a valid pick-independent primary signal.

   The c4 CQ table defines a pure-judgment A-floor over `(phase x record_kind)` and explicitly says the floor fires from record class, not seat self-assessment (`c4-cq-gateconfig/DESIGN-planner-20260702-015800.md`). The locked m-6 design consumes `gate_category in J2 A-set` as bucket A and records the decision-3 direction invariant as raise-only, with the CQ-3 A-floor composing with m-2's monotonic MAX at fill/submit. That makes S1 the right primary signal for a B pick on an A-worthy record.

2. S2 is a valid content signal when the submission references a live gate.

   The m-6 planner correctly limits S2 to the referenced gate record's own committed `gate_category`. The live `frank/` code path already reads the referenced gate record at verdict handling time: `submit.go:216-245` checks the referenced accepted gate-bearing record, rejects dead or already-resolved refs, and wakes the gate author. m-7's companion answer sanctions that committed-record read as an engine input and identifies this row as carrying the detector interplay. This is a store-read mechanism, not m-6-authored config.

3. S3 is already locked by J2's merge split.

   `ARCHITECTURE.md` J2 derives protected-branch merge classification from `target branch x protected-branch set`: protected-branch merges are A, feature-to-feature merges are B. This is a content predicate already present in the locked J2 map, not a new category-specific predicate engine. It is safe to include in Step-1 because it already exists and has operator-configured protected-branch-set inputs.

4. The residual boundary is honest and must travel with s5-b claim surfaces.

   The planner does not overclaim. The signal set catches exactly S1, S2, S3, plus the hardcoded `other` to A fail-safe. It does not claim every possible content mis-pick. `ARCHITECTURE.md` C4 now records the same boundary: per-category content predicates beyond S1 grain are Step-2+ carry. That boundary is load-bearing and should remain mandatory wherever s5-b describes decision-3 coverage.

5. No c3 lock reopen or new m-6 mechanism is introduced.

   The m-6 answer reads from the locked c3 design, the c4 CQ A-floor table, and J2. It clarifies which existing content-derived signals feed the known-A detector, while m-7 owns the rewrite-vs-bool mechanics and m-2 owns the registry rows. The m-6 consumer requirement is valid: downstream bucket routing must see an effective A category and must retain original-pick provenance for ODB/audit render, but the implementation mechanics remain outside this review.

## Carry-forward

- Approved: m-6's S1/S2/S3 known-A signal-set ruling for decision-3.
- Preserved boundary: Step-1 claims exactly S1 plus S2 plus S3 plus fail-safe, not universal content-mis-pick detection.
- Not approved here: m-7 token-rewrite mechanics as an independent review item, m-2 registry shape, m-1 owed/genesis confirm, m-4 confirm legs, or global s5 closure.
- No work authority granted: no PLAN, IMPL, merge, `frank/` edit, `pcode/` edit, design-doc edit, spike, or lock reopen.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s5-escalations/DESIGN-planner-20260706-051357.md` -- OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s5-escalations/RECONCILE-orchestrator-planner-20260706-052214.md` -- OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s5-escalations/SITREP-planner-20260706-051409.md` -- OK.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s5-escalations/SITREP-planner-20260706-053000.md` -- OK.
- `nl -ba master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md | sed -n '36,55p'` -- reviewed bucket A/B, decision-3 direction invariant, and hardcoded fail-safe consumption.
- `sed -n '1,140p' master/relays/c4-cq-gateconfig/DESIGN-planner-20260702-015800.md` -- reviewed CQ-3 A-floor table and no-new-mechanism boundary.
- `nl -ba master/ARCHITECTURE.md | sed -n '100,138p'` -- reviewed J2 A/B map, fail-safe, merge split, and raise-only direction.
- `nl -ba master/ARCHITECTURE.md | sed -n '492,496p'` -- reviewed settled S1/S2/S3 plus Step-1 claim-boundary registration.
- `nl -ba master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md | sed -n '268,282p'` -- reviewed `gate_category` slot, effective raised category, `gate_category_raised`, and terminal outcome interaction.
- `nl -ba frank/internal/engine/submit.go | sed -n '216,248p'` -- reviewed the referenced-gate verdict-path lookup.
- `git -C frank rev-parse --short HEAD` -- `67ee23e`.
- `git -C frank status --short` -- unrelated pre-existing untracked `?? docs/sprints/2026-07-06-s5-slice-5/`; no `frank/` edits made by this review.
- `git -C pcode status --short` -- unavailable: `fatal: cannot change to 'pcode': No such file or directory`.
- `git status --short` -- unavailable: cwd is not a git repo.

ACTIONS_GIT_REF: docs workspace is not a git repo; disk refs: created `master/relays/s5-escalations/DESIGN-REVIEW-implementer-20260706-052907.md` and appended one `master/relays/INDEX.md` row; no design-doc/source/`frank/`/`pcode` edits claimed.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repo; `frank/` has pre-existing untracked `?? docs/sprints/2026-07-06-s5-slice-5/`; `pcode/` is absent.
Next requested action: master.orchestrator-planner may fold this m-6 signal-set approval into the s5-escalations riding-leg ledger, preserving the Step-1 claim boundary and the remaining non-m-6 legs.
