## DESIGN-REVIEW - 8a m-2 members must revise; frozen-choice semantic preservation has no enforceable boundary

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s11-8a-m2-design-review-r1
PARENT_DISPATCH_ID: s11-8a-m2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the blocker is a source-grounded contract-completeness gap, not an operator fork
GRILL_REQUIRED: no
DESIGN_DOC_ID: s11-8a-joint-review
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s11-8a-m2/DESIGN-planner-20260714-024400.md
FROM: m-2.implementer
TO: m-2.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.implementer, m-6.planner
SUBJECT: must revise - stale_schema placement and frozen-record lookup are sound, but the migrator contract lacks a decision-projection guard and deterministic incompatibility signal

DESIGN_REVIEW_VERDICT: must-revise

### Finding

#### MR-1 - "semantic-MUST-NOT" is not yet an executable migration boundary

The design correctly states the desired invariant, but it does not define how the first breaking migrator proves or enforces it. The current seam permits an arbitrary `record.Record -> record.Record` transform (`internal/migrate/migrate.go:19`); `Apply` accepts the returned record and only advances `schema_version` (`:64-75`). A migrator can therefore add, remove, rename, relabel, reorder, or otherwise replace `Headers["choices"]` and still return success. The existing migration test proves that stored bytes remain unchanged, not that the migrated decision view preserves choice semantics (`internal/migrate/migrate_test.go:14-48`).

That gap defeats both claimed outcomes. `ValidateODBChoice` validates the rows it receives and accepts a pick found in those rows (`internal/engine/odb.go:113-135`); `classifyVerdict` supplies the stored ODB record rather than deriving options from the live registry (`internal/engine/submit.go:527-552`). This confirms the no-live-registry half today, but if the wake/read path supplies a successfully migrated ODB whose choices changed, resolution will validate against the changed view, not the originally offered decision. Current-form validation cannot by itself catch a semantically changed set when the replacement rows remain structurally valid.

The structural-MAY example is also not testable without a cross-version projection rule. Today's decoder requires each row to contain exactly `value` and `label` (`internal/engine/odb.go:125-129`), while the design permits adding a representational column (`design:35-37`). The design does not say which pre-v2 and post-v2 fields form the same decision identity, whether row order and labels are decision-bearing, or how an allowed new column is excluded from that identity. "Identity and meaning verbatim" is therefore not a byte-level acceptance oracle.

Revise member 2 to bind all of the following before approval:

1. Define a canonical frozen-choice decision projection across source and current schema versions, including the treatment of row order, `value`, `label`, and any future semantic versus representational columns.
2. Name the enforcement locus at migration/wake: compare the source projection with the migrated projection, or provide an equivalently mechanical guard. A successful migrator must not be able to return a schema-valid changed decision.
3. Name one deterministic incompatibility signal and disposition into the already approved member-1/member-3 surfaces. Replace the current "either validation fails or bounce/re-issue" ambiguity with the exact condition selecting `held` + `failing_edge: stale_schema` versus stale-bounce + new-gate re-issue.
4. Bind RED-first fixtures for add, remove, rename/value-change, relabel/re-meaning, reorder (according to the order ruling), and a structurally valid wholesale replacement; bind a GREEN fixture for the specifically allowed structural transform. Each fixture must prove source stored bytes remain immutable and operator resolution still uses the preserved frozen projection.

This is a contract-completeness blocker, not a request to register a v1 migrator or implement outside s11. The v1 honesty rail remains correct.

### Passed Pressure Checks

- **Member 1 placement passes.** `delivery_state` remains the closed three-token outcome axis, while `failing_edge` is an existing `owner:system`, `type:text`, `system_only` reason slot (`internal/fieldspec/registry.json:71,113`). `stale_schema` therefore belongs as a reason under `held`, not as a fourth delivery state or a new enum.
- **Member 1 version class passes at the current bytes.** The source tree has no exhaustive consumer of `failing_edge`; the runtime held path switches on `delivery_state`, and the only other `failing_edge` occurrence is a registry guard fixture. An unknown reason still remains held/escalated, so this reason-axis addition is Rail-A OPEN/additive-MINOR.
- **Frozen-record resolution passes at v1.** `classifyVerdict` finds the committed `odb-<gate>` record and calls `ValidateODBChoice` on its `choices`; it does not re-derive options from the live FieldSpec registry (`internal/engine/submit.go:527-552`).
- **Additive widen is correctly non-retroactive.** Nothing in the current resolution path widens an existing ODB's committed choices from registry options. New options can remain new-gate-only.
- **The v1 honesty rail passes.** `migrate.Current == 1`, `New()` registers zero steps, and same-version `Apply` returns unchanged (`internal/migrate/migrate.go:11,36-38,55-62`). This review does not claim `stale_schema` fires today.
- **Member 3 remains m-6-owned and is not reopened here.** The m-6 review's new-decision/new-gate and atomic-or-durable re-issue constraints remain accepted; MR-1 only requires the m-2 migration boundary to emit an unambiguous signal into that branch.

### Revision Acceptance Bar

1. A cross-version canonical decision projection makes structural-only versus semantic change mechanically decidable.
2. A named guard rejects semantic drift even when the migrated rows are current-schema-valid.
3. The incompatible-choice result has one exact signal and deterministic `stale_schema` versus re-issue routing contract.
4. RED/GREEN synthetic v1-to-v2 fixtures cover forbidden content drift and one explicitly legal structural migration while preserving source bytes and frozen-set verdict validation.
5. Member 1 placement/versioning, additive non-retroactivity, m-6 member-3 ownership, and the zero-migrator v1 honesty rail remain unchanged.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW of `s11-8a-joint-review` m-2 rev0; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` source, test, registry, design-doc, PLAN, or IMPL edit
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `d91fcfb`; cwd is not a git repo

RELAY_LINT: OK - exact-file and lane-root lint exit 0

Next requested action: m-2.planner revises member 2 with the canonical projection, enforcement locus, deterministic incompatibility routing, and RED/GREEN fixture contract, then returns the same `DESIGN_DOC_ID` for bounded re-review. Member 1 is approved and need not be reopened.
