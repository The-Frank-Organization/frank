## AUDIT - s5-b.implementer independent audit: mechanisms and versioning

ROLE: Implementer
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: s5-audit-s5-b
PARENT_DISPATCH_ID: s5-resume
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s5-b.implementer
TO: s5.orchestrator-planner
CC: s5.orchestrator-reviewer, operator
SUBJECT: AUDIT - s5-b mechanisms/versioning survey: decision-3 raise, decision-5 scanner, section-7 fixtures, replay, I-PH, boundary list

IN_REPLY_TO: `.relays/s5/s5-audit-s5-b/AUDIT-orchestrator-planner-20260706-035446.md`

## Boot acknowledgement folded into audit

- Identity accepted: this session is `s5-b.implementer`, the Implementer seat for pair `s5-b` mechanisms and versioning.
- Skill loaded: `agent-pair-implementer` plus `protocol.md` read in full.
- Team charter read: `master-docs/CLAUDE.md`.
- Boot relay read: `.relays/s5/boot/s5-boot-s5-b-implementer/SITREP-orchestrator-planner-20260706-040203.md`.
- Mode-change relay read: `.relays/s5/s5-resume/PLAN-orchestrator-planner-20260706-034602.md`.
- Fidelity reconciliation read: `.relays/s5/s5-fidelity/RECONCILE-orchestrator-planner-20260706-034602.md`.
- Audit dispatch read: `.relays/s5/s5-audit-s5-b/AUDIT-orchestrator-planner-20260706-035446.md`.
- Transport rule accepted: no frank MCP tools used or retried; this report is a file relay under `.relays/s5/`.

## Routing and authority

- The audit dispatch is addressed to `s5-b.planner, s5-b.implementer`, so this Implementer seat is directly in `TO`.
- Authority is read-only AUDIT. No source, test, registry, sprint doc, branch, commit, PR, merge, transport fix, or store archive action was taken.
- The active dispatch contains no bare own-line `DISPATCH IMPL`; implementation remains unauthorized.
- Independence observed: before filing this audit, `.relays/s5/s5-audit-s5-b/` contained only the orchestrator dispatch, so I did not read a sibling s5-b planner audit.

## Summary verdict

PRIMARY_BUCKET: still-open
still-open: s5-b has real implementation work. Decision-3 has only a partial helper and partial submit stamping; decision-5 has no scanner or drain call site; section-7 has strong s4 coverage but not the adapted s5 fixture bundle; replay/versioning has a seam and refusal unit tests but not the required zero-loss store replay or read-path negotiation proof; I-PH fixtures need the new s5 surfaces.
already-closed: The submit path already stamps identity and schema_version before validation, validates before lineage and commit, records `gate_category_raised` in a computed registry row, materializes outbox items through store intents, has config-change mechanics, has migration seams and refusal errors, and has existing I-PH fixtures.
product-overlapped: Registry rows, enum member changes, and live `model_name` declaration belong to s5-a. s5-b should consume those rows and use fixture-local registry views where needed, not edit `internal/fieldspec/registry.json` in this phase.
recommended-next: Proceed to PLAN only after reconciling the s5-a/s5-b boundary. Build s5-b as small production mechanisms plus tests: submit-stage raise, dormant scanner at the real outbox/drain seam, section-7 fixture extensions, replay/read-path versioning tests, and I-PH extensions. Do not touch transport-fix lineage/parenting/codec issues.

## 1. Decision-3 placement survey

PRIMARY_BUCKET: still-open
still-open: The full known-A detector is not wired into the submit critical section. `Validate` always calls `ClassifyGateCategory(raw, false)`, so no submit path currently supplies the `knownA=true` case. A B-token such as `routing` is therefore not raised by submit validation today. Evidence: `internal/fieldspec/validate.go:61-65` and `internal/fieldspec/validate.go:176-187`.
already-closed: The correct insertion zone exists. `SubmitHandlerWithRender` decodes payload, stamps seat identity, assigns relay ID and intake ID, stamps `schema_version = migrate.Current`, builds tables, then calls `reg.Validate` before lineage, record_kind dispatch, config-change handling, verdict handling, and commit intents. Evidence: `internal/engine/submit.go:25-82`.
already-closed: The live registry has `gate_category` and `gate_category_raised` rows, and `gate_category_raised` is not rendered because computed rows are hidden. Evidence: `internal/fieldspec/registry.json:92-93` and `internal/fieldspec/render.go:50-52`.
product-overlapped: The dispatch says `gate_category_raised` is conductor/system-stamped; the live row is `owner:"computed"` with `fill_constraints:"computed_result"`, not `owner:"system"`. This may be semantically acceptable, but the row itself is s5-a registry ownership and should be reconciled there, not changed by s5-b. Evidence: `internal/fieldspec/registry.json:93`.
recommended-next: Implement the raise at the existing validate point after stamp/schema_version and before lineage/record_kind dispatch. The engine should supply the known-A floor from the trusted registry/config view rather than a second detector list. Keep it raise-only: do not create an A-to-B path, and record the raise in `gate_category_raised`.

Ordering constraints:

- Post-stamp: `seat.Stamp` runs before validation at `internal/engine/submit.go:29`.
- Post-schema stamp: `cand.Envelope.SchemaVersion = migrate.Current` runs at `internal/engine/submit.go:36`.
- Pre-lineage: `reg.Validate` runs before lineage at `internal/engine/submit.go:47-53`.
- Pre-record-kind/config/verdict/commit: record_kind dispatch and accepted projection intents happen only after validation and lineage at `internal/engine/submit.go:58-82`.

Duplicate gate:

- `ClassifyGateCategory` and `GateCategoryAuthorityBearing` exist and should be reused rather than rebuilt (`internal/fieldspec/validate.go:176-197`).
- Current helper tests cover raw class behavior, including manual `knownA=true`, but they do not prove submit-stage raise because submit never passes that flag (`internal/fieldspec/fieldspec_test.go:61-75`).

## 2. Decision-5 chokepoint survey

PRIMARY_BUCKET: still-open
still-open: No production scanner, ODB-specific scanner, `model_name`, or real drain/egress-render call site exists in current code. A source scan of `internal/`, `cmd/`, and `test/fixtures/` found only `bufio.Scanner` uses and no egress scanner or model-name code.
already-closed: The local outbox materialization path exists. `completeOutbox` derives an outbox record body and an `IntentOutbox`; `Store.applyIntents` writes `outbox/<item_id>.json`; `canonicalProjectionIntents` can rebuild outbox projection files from canonical record bodies. Evidence: `internal/obligation/obligation.go:143-185`, `internal/store/projections.go:76-81`, and `internal/store/projections.go:169-184`.
product-overlapped: The present path is outbox production/materialization, not external send. The ruled R-2 scanner must sit at a real future drain/egress-render seam, present but dormant, not inside local item production. The fixture should drive that real function at its call site with a fixture-enabled drain leg.
recommended-next: Add a small production scanner function and a dormant drain/render call site that accepts an outbox item plus destination/context. Keep live activation off. Fixture the three ruled legs: same model-name outside the ODB carve-out blocks; secret/PII/auth-URL inside ODB still blocks; lane-supplied exempt-marked content still blocks. Use fixture-local ODB/model_name inputs if s5-a has not yet landed live rows.

Chokepoint map:

- Outbox record construction: `internal/obligation/obligation.go:149-181`.
- Outbox file projection: `internal/store/projections.go:76-81`.
- Canonical outbox rebuild from accepted record body: `internal/store/projections.go:177-184`.
- No drain call site: code has `Project`, `Read`, outbox projection, and wake/poll paths, but no external sender or egress-render function under `internal/` or `cmd/`.

Honesty sweep obligation:

- New scanner docs, code comments, errors, and tests must say present-but-dormant or fixture-enabled. They must not say live scanning, live external send, or sole system egress.
- The scanner verdict/error strings become seat-deliverable if surfaced through read/project/bounce and must be included in I-PH tests.

## 3. Section-7 fixture legs

PRIMARY_BUCKET: recommended-next
still-open: The exact adapted s5 section-7 fixture bundle is not present as one test: operator-authorized record shape, old-to-new digest, no re-genesis, phase-0 digest acceptance through the genesis plus config-change chain, stale-form re-render plus re-rendered success.
already-closed: The s4 suite covers much of the mechanism. It proves operator config_change materializes new registry bytes, running config is not hot-reloaded, wrong `new_digest` is rejected, non-operator reads are redacted, projections carry no member bytes, and stale form bounces after restart before a fresh render succeeds. Evidence: `test/fixtures/s4_config_change_test.go:21-65`, `test/fixtures/s4_config_change_test.go:167-208`, and `test/fixtures/s4_config_change_test.go:210-281`.
already-closed: Production config-change validation and materialization exist. Evidence: `internal/engine/submit.go:164-205`, `internal/store/config_change.go:11-24`, and `internal/store/genesis.go:125-140`.
product-overlapped: Do not rebuild s4 coverage. Extend it with the missing adapted assertions, especially no re-genesis and phase-0 acceptance from the accepted config-change chain.
recommended-next: Add an s5 fixture file that reuses s4 harness helpers where possible. Construct a fresh store with `store.Init`, commit an operator fieldspec config_change, restart or run phase-0 validation, assert one genesis record remains, assert `ExpectedConfigDigest` walks from genesis digest to the latest config_change digest, assert stale form rejection, then assert fresh render acceptance.

Duplicate gate by leg:

- Operator-authorized shape: mostly covered by `configChangeRelay` and operator submit in s4; s5 should assert envelope/headers explicitly.
- Old-to-new digest: covered indirectly by `fixtureDigestWithMember`; s5 should assert before digest differs from after digest and accepted record carries the after digest.
- No re-genesis: not explicitly asserted in s4; s5 should count/verify a single `genesis` record.
- Phase-0 acceptance through genesis chain: partially covered by restart success; s5 should call or force the `ValidateGenesis`/startup path after accepted config_change.
- Stale-form re-render plus re-rendered success: already covered at `test/fixtures/s4_config_change_test.go:210-281`; s5 can extend rather than duplicate.

## 4. Versioning artifact set

PRIMARY_BUCKET: still-open
still-open: The required registry version-label bump is not landed in current `internal/fieldspec/registry.json`; it still reads `s3-fieldspec-v2`. This file belongs to s5-a. Evidence: `internal/fieldspec/registry.json:2`.
already-closed: The migration seam exists with `Current = 1`, a registry of version steps, read-time `migrate.Reader`, source-version reporting, and refusal errors for future/unversioned/gap. Evidence: `internal/migrate/migrate.go:11-17`, `internal/migrate/migrate.go:36-90`, and `cmd/frank/main.go:295-333`.
already-closed: Unit tests already cover copy-on-read migration, refusal errors at `migrate.Apply`, and store checksum errors winning before migration. Evidence: `internal/migrate/migrate_test.go:14-86`.
still-open: No new zero-loss replay over a real or constructed store exists. No test asserts count, relay identity, and canonical-wins behavior through `migrate.Reader` over a multi-record store. No fixture proves negotiation/refusal through the actual channel/read path.
product-overlapped: Do not add an envelope migrator. The fidelity ruling says `migrate.Current` changes only on record-shape changes; s5 consumer-schema additions are registry/config changes, not canonical record envelope migrations.
recommended-next: Leave `migrate.Current` at 1. Add a zero-loss replay test independent of `test/replay/harness.go`'s archived relay-lint oracle paths. Prefer an operator-supplied copy of `~/frank-archives/frank-team-store-s5-dogfood-20260706`; fallback is a constructed store with accepted, rejected, held, config_change, and outbox records. Assert record count, relay IDs, source versions, migrated view identity at Current=1, and canonical projection wins after rebuild. Add read-path refusal legs via `migrate.Reader.Read` or channel `read`, not only `migrate.Apply`.

Current read path:

- `cmd/frank` `read` uses `migrate.Reader{Store: st, Registry: migrate.New()}` and returns both `schema_version` and `source_schema_version`. Evidence: `cmd/frank/main.go:295-333`.
- `Store.Read` verifies sealed canonical record bytes before migration. Evidence: `internal/store/store.go:134-148`.

False migrator trigger to avoid:

- A registry enum/member/version-label change alone is not a record-shape change.
- A MAJOR/MINOR compatibility label from m-2 is not by itself a `migrate.Current` bump.
- Adding an ODB fixture-scoped registry view is not a live `record_kind` widening and does not justify an envelope migrator.

## 5. I-PH extension inventory

PRIMARY_BUCKET: still-open
still-open: s5-b-added scanner verdict/refusal text, drain fixture text, and read-path version refusal text are not yet in the I-PH matrix because the surfaces do not exist.
already-closed: Base I-PH and s4 bridge I-PH fixtures exist. They cover bounce strings, migration refusal strings, registry-load errors, tool descriptions, submit schema, tool-call results, auth failure, diagnostics, reconnect errors, and the one-value frame-size carve-out. Evidence: `test/fixtures/iph_test.go:25-115` and `test/fixtures/s4_iph_test.go:26-146`.
already-closed: Bounce formatting preserves the `Field:Class` and lineage `Edge:Kind` shape and drops raw reasons. Evidence: `internal/bounce/formatter.go:11-26`.
product-overlapped: ODB `model_name` row declaration is s5-a/m-6; s5-b only needs scanner fixtures and any scanner result strings it creates.
recommended-next: Extend `iph_test.go` and/or `s4_iph_test.go` with every new s5-b seat-deliverable string: scanner verdicts/classes, any egress refusal payloads, any fixture-enabled drain readback, and read-path migration refusal text. Keep `bounce.Format` unchanged unless a later dispatch explicitly authorizes bounce-surface changes.

Seat-deliverable surfaces s5-b is likely to add:

- Decision-3 rejection or raise proof surfaces: no new bounce text required if `gate_category_raised` is only a computed header.
- Decision-5 scanner verdicts: block/pass reasons or classes from the dormant drain fixture.
- Version negotiation/refusal: future schema_version, unversioned record, and migration gap read-path surfaces.
- Section-7 stale form: existing `form_digest:re-render` bounce already path-clean.

I-PH forbidden families to continue using include store/config/outbox path families, binding artifacts, socket paths, credentials, and config values. Existing assertions use these families at `test/fixtures/iph_test.go:31-35`, `test/fixtures/iph_test.go:106-114`, and `test/fixtures/s4_iph_test.go:234-245`.

## 6. Boundary file list

PRIMARY_BUCKET: recommended-next
still-open: s5-b should produce a narrow IMPL file list during PLAN. This audit names expected surfaces but does not authorize edits.
already-closed: Boundaries are clear in the dispatch: engine/bounce/migrate/test-replay/new fixture files are s5-b; registry rows and enum changes are s5-a.
product-overlapped: `internal/fieldspec/registry.json` is a collision if s5-b touches it. Any model_name, record_kind ODB, or routing_escalation row work belongs to s5-a.
recommended-next: Keep s5-b production edits to mechanisms and tests. Use fixture-local registry copies for scanner/replay tests when s5-a rows are not landed yet.

Expected s5-b IN surfaces at IMPL time:

- `internal/engine/submit.go`: possible source of submit-stage known-A floor plumbing.
- `internal/fieldspec/validate.go`: possible source of raise-only classification behavior if the API needs to accept known-A context.
- `internal/fieldspec/fieldspec_test.go` or `internal/engine/pipeline_test.go`: decision-3 unit/submit tests.
- New scanner package or file such as `internal/egress/...` or a narrowly named internal package: decision-5 production scanner, present but dormant.
- `internal/obligation/obligation.go` or a new drain adapter only if PLAN identifies this as the real dormant drain call site; do not put scanner in local item production as a live always-on check.
- `internal/migrate/migrate_test.go` and/or new fixture tests: read-path negotiation/refusal and zero-loss replay.
- `test/fixtures/s5_config_change_test.go`: section-7 adapted fixture legs.
- `test/fixtures/s5_egress_scan_test.go`: decision-5 (a)/(b)/(c) fixture-enabled drain legs.
- `test/fixtures/iph_test.go` and/or `test/fixtures/s4_iph_test.go`: I-PH extension for new surfaces.
- `test/fixtures/applicability_map.go`: only if a new crash/application fixture class is added and PLAN scopes it.

Expected OUT or collision surfaces:

- `internal/fieldspec/registry.json`: s5-a single-writer.
- `internal/fieldspec/registry_test.go`: s5-a if the test asserts live registry row inventory or version-label bump.
- Any live `record_kind` ODB widening: out of s5-b and out of Step-1 live system.
- Any transport-fix work: lineage/parenting/codec fixes are next master cycle, not s5-b.
- Archived dogfood store: operator-held evidence; propose a copied input only.
- Live external send activation or away bridge: later-step carry, not this slice.

## Open questions for reconcile

1. Does s5-a intend `gate_category_raised` to remain `owner:"computed"` or change to `owner:"system"`? This is registry ownership, but s5-b needs the stable contract before PLAN.
2. Should decision-3 raise overwrite the effective `gate_category` token, or only stamp `gate_category_raised` while leaving the agent's raw token? The current code only stamps the boolean. The design text says effective/forced result has a home in `gate_category`.
3. Where should the dormant drain function live? Current code has local outbox materialization but no external drain. The plan should name a real function boundary so the scanner is production code, not a test stub.
4. Should zero-loss replay use an operator-supplied copy of the archived dogfood store, or is a constructed test store enough for this checkpoint? The archive is operator-held and was not touched.

## Actions

- Performed read-only audit of the addressed s5-b boot and AUDIT relays plus current source/design surfaces.
- Wrote this audit relay and appended the s5 index row.
- No code, tests, registry, docs outside `.relays/`, branches, commits, PRs, merges, store archives, or transport fixes were changed.

ACTIONS_GIT_REF: read-only audit relay written at `.relays/s5/s5-audit-s5-b/AUDIT-implementer-20260706-041025.md`; index row appended at `.relays/s5/INDEX.md`; no source/test/runtime edits made; `.relays/` is gitignored.
FINAL_GIT_STATUS_SHORT: none - clean tree

## Verification

- E1: exact boot relay read: `.relays/s5/boot/s5-boot-s5-b-implementer/SITREP-orchestrator-planner-20260706-040203.md`.
- E1: exact AUDIT dispatch read: `.relays/s5/s5-audit-s5-b/AUDIT-orchestrator-planner-20260706-035446.md`.
- E1: master mode-change and fidelity relays read: `.relays/s5/s5-resume/PLAN-orchestrator-planner-20260706-034602.md` and `.relays/s5/s5-fidelity/RECONCILE-orchestrator-planner-20260706-034602.md`.
- E1: design records spot-checked: master `ARCHITECTURE.md` sections C4/J2, m-2 versioning/decision-3, m-3 decision-5, m-6 J2/ODB, m-7 config/I-PH/version seams.
- E1: code surfaces read and cited: `internal/engine/submit.go`, `internal/fieldspec`, `internal/store`, `internal/obligation`, `internal/migrate`, `cmd/frank/main.go`, `test/replay/harness.go`, `test/fixtures/iph_test.go`, `test/fixtures/s4_iph_test.go`, `test/fixtures/s4_config_change_test.go`, `internal/bounce/formatter.go`.
- E2: `git status --short` before relay write produced no output.
- E2: `git status --short` after relay write and index append produced no output.
- E2: relay lint on this exact file passed.
