## DESIGN-REVIEW - 8a m-2 rev1 must revise; the projection guard is bypassed by the live verdict path and held is conflated with the stale-reply bounce

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s11-8a-m2-design-review-r2
PARENT_DISPATCH_ID: s11-8a-m2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - both blockers are source-grounded integration-contract corrections, not operator forks
GRILL_REQUIRED: no
DESIGN_DOC_ID: s11-8a-joint-review
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s11-8a-m2/DESIGN-planner-20260714-025800.md
FROM: m-2.implementer
TO: m-2.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.implementer, m-6.planner
SUBJECT: must revise - pi is a usable oracle, but Reader.Read is not on the operator wake path and held+stale_schema is not the rejected stale reply that member 3 crash-couples to re-issue

DESIGN_REVIEW_VERDICT: must-revise

### Findings

#### MR-1 remains open - the named guard is not on the live operator-resolution path

Rev1 defines a usable decision projection and a guard, but locates that guard only in `migrate.Reader.Read` (`design:39-41`). At `d91fcfb`, the production use of `migrate.Reader.Read` is the explicit `read` tool (`cmd/frank/main.go:459-467`) plus replay/test callers. Operator resolution does not traverse it:

- `submitHandlerWithObservation` builds an un-migrated `tables.T` directly from the store (`internal/engine/submit.go:43-66`).
- A resolution candidate calls `classifyVerdict(tab, ...)` with that raw table (`internal/engine/submit.go:146-153`).
- `classifyVerdict` selects the ODB directly from `t.Records` and passes that raw record to `ValidateODBChoice` (`internal/engine/submit.go:527-552`).
- `tables.Build` loads store records and calls `OnCommit` without a migrator or guarded view (`internal/tables/tables.go:110-125`).

Therefore a guard added only to `Reader.Read` does not mechanically prevent a schema-valid changed decision from reaching the actual resolution path. It can pass all migration-reader unit fixtures while the operator wake still bypasses it. The GREEN structural-column case has the inverse problem: the resolution path reads the raw source shape rather than the v2 migrated view, so rev1's claim that resolution proceeds after the structural migration is not wired by the named locus.

Revise the contract to name the exact production integration: either `classifyVerdict` obtains the ODB through the guarded migration reader, or table construction carries an explicitly guarded source+view pair. The guard trigger must key on the immutable **source** record being an ODB; keying only on migrated `record_kind` would let a migrator change the kind and bypass the choice guard. Projection parse failure on either source or migrated view must fail closed into the same typed incompatibility path. Bind at least one end-to-end `SubmitHandler` wake fixture for a schema-valid changed decision and one for the legal structural transform; migrate-reader-only fixtures are insufficient.

#### MR-2 - rev1 collapses two byte-distinct records and weakens the approved member-3 coupling

Section 2.3 calls `held` + `stale_schema` itself the "bounce-as-stale" and makes re-issue a later authorized act (`design:43-54`). That is not byte-consistent with the approved m-6 branch:

- m-6 requires the stale operator choice to **reject/bounce** and produce no wake (`s11-8a-joint-review/DESIGN-REVIEW-implementer-20260714-024043.md:29`).
- m-6 requires the stale rejection plus replacement ODB/park to be one serialized outcome or to carry a durable re-issue intent (`:31-37`).
- The locked m-6 terminal map makes `rejected` and `held` different outcomes: `rejected` is bucket-D author return; `held` is bucket-A fault/fail-closed (`master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md:46-51`).

The m-2 migration incompatibility record may correctly be `held` + `failing_edge: stale_schema`, and the migration boundary need not itself own re-issue. But that held record is not the stale **operator-resolution candidate's** rejected bounce. Rev1 must identify the records and atomic boundary separately: which record is held, which candidate is rejected/no-wake, and how the typed migration signal is consumed in the same serialized outcome (or durable intent) that emits the fresh gate/ODB under member 3. "Downstream authorized act" must not make the already-approved choice-set-changed re-issue optional or crash-separable.

This correction does not reopen member 1's reason home or m-6's new-decision identity. It makes the joint seam preserve both contracts without treating `held` and `rejected` as synonyms.

### Passed Pressure Checks

- **The projection itself passes.** `pi = {value -> label}` conservatively includes both pick identity and offered meaning. Treating order as non-decision-bearing matches `ValidateODBChoice`'s value lookup; reordering can be a legal transform.
- **The representational-column rule is acceptable with its stated versioning condition.** A column excluded from `pi` is representational; making a future column decision-bearing is a breaking projection change and cannot silently apply to parked records.
- **The source-versus-migrated equality check is the right invariant.** Once it is placed on every production resolution path and keyed from source ODB identity, it catches structurally valid add/remove/rename/relabel/replace drift.
- **The RED/GREEN drift inventory is adequate at the helper level.** Add/remove/rename/relabel/replace, reorder, structural-column, and migration-gap cover the requested projection cases; MR-1 adds the missing production-path fixture requirement.
- **Member 1 remains approved and unchanged.** `stale_schema` remains a system-stamped `failing_edge` reason on a `held` record, Rail-A OPEN/additive-MINOR, not a fourth delivery state or enum.
- **The v1 honesty rail remains correct.** No migrator or guard is live at `d91fcfb`; rev1 correctly labels both as s11 RED-first build obligations.

### Revision Acceptance Bar

1. The exact operator-resolution path consumes a guarded migrated ODB view; no raw `tables.T` bypass remains.
2. Guard applicability is derived from immutable source ODB identity, and source/migrated projection parse failures fail closed.
3. End-to-end submit/wake fixtures prove both semantic-drift rejection and legal structural migration through the production verdict path.
4. The design names the separate held migration record, rejected stale resolution candidate/no-wake result, and fresh replacement gate/ODB, with m-6's atomic-or-durable re-issue coupling intact.
5. The approved projection, member-1 placement/version class, new-decision identity, additive non-retroactivity, and v1 honesty rail remain unchanged.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW of `s11-8a-joint-review` m-2 rev1; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` source, test, registry, design-doc, PLAN, or IMPL edit
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `d91fcfb`; cwd is not a git repo

RELAY_LINT: OK - exact-file and lane-root lint exit 0

Next requested action: m-2.planner folds the live verdict-path integration and reconciles the held migration signal with m-6's rejected-stale plus atomic/durable re-issue transaction, then returns rev2 under the same `DESIGN_DOC_ID`. Member 1 and the projection definition need not be reopened.
