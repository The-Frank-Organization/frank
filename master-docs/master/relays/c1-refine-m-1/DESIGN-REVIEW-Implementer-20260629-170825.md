## Team m-1 - Trust & Identity: DESIGN-REVIEW rev2

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c1-design-m-1-review-r2
PARENT_DISPATCH_ID: c1-refine-m-1
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - read-only DESIGN-REVIEW of rev2 fold; operator decisions unchanged from rev1 grill
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, m-2.planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-1-trust-identity
OWNER: m-1 (Trust & Identity)
DESIGN_DOC_ID: c1-design-m-1-trust-identity
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: c1-refine-m-1/DESIGN-planner-20260629-170308.md

### Verdict

APPROVE. The rev2 fold preserves the rev1 approved design and correctly incorporates the consumer-review
findings without changing the m-1 target entity, public verbs, system-field split, or non-re-cut path.

This is a read-only design review. I made no source or `pcode/` edits.

### Scope reviewed

Reviewed:
- `master/relays/c1-refine-m-1/DESIGN-planner-20260629-170308.md`
- `master/domains/m-1-trust-identity/design/2026-06-28-trust-identity-design.md` rev2
- `master/relays/c1-refine-m-1/DESIGN-orchestrator-planner-20260629-164925.md`
- m-1/m-2 rev2 reaffirm/concur relays (`165702`, `170032`, `170308`)
- VP fold-confirm `master/relays/c1-consumer-reconcile/RECONCILE-orchestrator-reviewer-20260629-164726.md`

### Challenge review

1. **DI-5 / I3 distinctness passes.** DI-5 is not double-counted with DI-2. The design states DI-2 as sibling
   credential/connection confidentiality and DI-5 as the conductor's out-of-lane read vantage over workspace,
   git, or process state. A lane can hold its own valid credential while still lying about its own workspace if
   the conductor accepts self-report, so I3 needs DI-5 as a separate invariant. The fallback label is also
   honest: without DI-5, evidence is "self-reported / confusion-resistant," not observe-by-construction.

2. **Observe spoof surface is bounded.** Under DI-5, the lane cannot substitute or intercept the observed value
   because the conductor reads and fills the observed fields at submit. The remaining PLAN detail is snapshot
   semantics: an observed value is the conductor's point-in-time read at submit; later lane state drift needs a
   new observe rather than mutating the accepted record. This is consistent with append-only records.

3. **Operator-FROM stamping is acceptable with the special-channel framing.** Rev2 correctly avoids treating
   `operator` as a lane-minted credential. The operator is both a delivery target and a special stamped `FROM`
   over the operator-relay channel, so a lane cannot present `operator` via `submit()`. This does add a trusted
   path, but it is the existing charter-level operator trust path, not a new lane capability. PLAN should keep
   the invariant explicit: lane processes must have no write access to the operator-relay channel.

4. **Routing-relay parent candidate passes.** Admitting an accepted routing relay into the conductor-derived
   `parent_picker` candidate set does not create a free-typed-parent surface. The routing relay must already be
   an accepted, stamped record in the seat lineage; the dispatch can reference it as provenance/bookkeeping.
   Rev2 keeps model identity payload-only and never a gate input, and m-2's concurrence confirms no `model_*`
   predicate enters the schema gate.

5. **m-3 write-allowlist prevents a second identity authority.** The allowlist says the hook may write only the
   closed observed/computed evidence set and a pass/fail veto, with no system-only, identity, envelope, or
   delivery writes. That is the right mechanical shape. PLAN should bind the written `*_RESULT` names to the
   m-2-declared closed field set so the wildcard-looking notation cannot become an extension point by accident.

6. **Submit ordering preserves I1/I2 and closes the rev1 TOCTOU note at design level.** Rev2 now says a relay is
   lane-proposed while submitted and becomes accepted only after resolve/stamp, m-2 validation, m-3 observe, and
   atomic conductor-owned append plus mailbox projection. Nothing observable is appended before pre-flights pass,
   and there is no lane re-entry between validation and append.

7. **No rev1 regression found.** The folds add DI-5/I3, special-address handling, routing provenance, allowlist
   wording, and clearer submit ordering. They do not weaken DI-1 through DI-4, `FROM`/`ROLE` stamping, `PARENT`
   forgery resistance, recipient validation, append-only store semantics, or the token-to-SO_PEERCRED/mTLS
   non-re-cut path.

### Acceptance carry-forward

Carry these into the eventual PLAN/E2 criteria; they do not reopen DESIGN:
- DI-2 and DI-5 must both be tested as independent isolation/read-vantage properties.
- The operator-relay channel must be unreachable by lane processes; lanes must not be able to submit as
  `operator`.
- The m-3 observed/computed field allowlist must resolve to an explicit m-2 schema-declared set.
- Observed evidence is a point-in-time conductor snapshot attached to an immutable accepted record.

### Disposition

`DESIGN_REVIEW_VERDICT: approve`

This approval is for the m-1 rev2 pair-review boundary. The joint m-1/m-2 co-foundational lock remains with
the orchestrator/VP after both rev2 pair reviews.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and updated master/relays/INDEX.md; no source or pcode edits.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (git status --short returns fatal); pcode git status --short is clean.
RELAY_LINT: OK - relay-lint.py exit 0 on this file; relay-root mode still reports the standing INDEX.md header-field errors.
