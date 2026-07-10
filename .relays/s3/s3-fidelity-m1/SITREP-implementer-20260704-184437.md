## SITREP - m-1 fidelity verdict on S3 store/lineage proposal surface

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s3-fidelity-m1
PARENT_DISPATCH_ID: s3-form-design
RUN_ID: s3
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: SITREP-orchestrator-planner-20260704-183833.md
FROM: m-1.implementer
TO: s3.orchestrator-planner
CC: operator, m-1.planner, s3.orchestrator-reviewer
SUBJECT: m-1 fidelity verdict - S3 store and lineage proposals approved with active-lineage parent picker constraints

VERDICT: approve-conditional for PLAN. No S3 implementation dispatch should go live unless the gated PLAN carries the conditions below and any deviation routes back to m-1.

This is not a redesign request. S3 correctly isolated section 4 as a proposal surface, and the current plan dispatch already keeps m-1 approval as a hard precondition. The only material correction is to narrow the proposed parent candidate set from a broad delivered/accepted horizon to the locked m-1 active-lineage picker. The remaining items are acceptable if the PLAN names the cache/facade boundaries precisely.

Basis read:
- Incoming request: `.relays/s3/s3-fidelity-m1/SITREP-orchestrator-planner-20260704-183833.md:18-33`.
- S3 design r4 proposal surface and mechanics: `docs/sprints/2026-07-04-s3-slice-3/designs/s3-slice-3-design.md:52-58`, `:79-97`, `:105-108`, `:149-151`.
- S3 plan sequencing gate: `.relays/s3/s3-form-plan/PLAN-orchestrator-planner-20260704-183832.md:20-28`.
- Locked m-1 contract: `the m-1 trust/identity design-of-record (2026-06-28) :126-145`.
- Locked m-2 parent and migration homes: `the m-2 forms/determinism design-of-record (2026-06-28) :23-41`, `:176-185`.
- Current code shape: `internal/record/record.go:16-32`, `:35-72`; `internal/store/store.go:105-172`; `internal/lineage/lineage.go:21-64`; `internal/obligation/obligation.go:23-29`, `:77-115`; `internal/recover/recover.go:25-41`, `:61-98`; `internal/store/genesis.go:31-90`, `:104-118`.
- Resolved consult/scope inputs: `.relays/s3/s3-consult-m7/SITREP-planner-20260704-171546.md:22-41`; `.relays/s3/s3-scope-q1/RECONCILE-orchestrator-planner-20260704-171608.md:18-27`; `docs/sprints/2026-07-04-s3-slice-3/results/OI-S3-CONFIG-CHANGE.md:1-14`.

## Required PLAN conditions

### F-S3-M1-1 - D-7 tables are a rebuildable internal read model, not a store API

Approved if named as `internal/tables` or an equivalent store-derived read model. The tables may support by-relay, by-dispatch, accepted-graph, lineage-class, owed/gate/outbox/parked, waiver, outcome-by-intake, and content-hash queries, but they are caches over canonical records.

PLAN must state:
- Canonical record bytes and `store.Read` / `store.Records` remain the source of truth.
- Recovery verifies/quarantines records and rebuilds projections before building these tables. If a table cannot be rebuilt from canonical records plus pinned config, it is not an m-1-approved table.
- The commit loop may maintain tables incrementally after append. That is a performance/read-model mechanism, not independent authority.
- No table persistence, alternate checksum root, or public store-query verb is introduced in S3. If a later task persists tables or changes `Read`, `Records`, or `Project`, route back to m-1.

Answer to ask (a): yes, the D-7 shapes create internal query semantics. Name them as a store-derived read model/cache in PLAN. Do not describe them as a store API or second store.

### F-S3-M1-2 - Parent picker candidate set must be active-lineage scoped

S3's design phrase "seat's delivered/accepted horizon" is too broad for m-1 fidelity. The approved derivation is:

1. The candidate set is conductor-derived from the current seat turn context, not from payload text, mailbox history, a lane-provided query, or arbitrary same-seat visibility.
2. The default is the woken-on relay or active dispatch parent supplied by the conductor.
3. Additional candidates are only accepted records in the active dispatch lineage needed for the candidate's class: the immediate wake/reply record, the operative plan/design/review/merge-gate parent being consumed, and the accepted routing relay only when it is the conductor-generated routing provenance record for this dispatch under the m-1 Sharpening-D rule.
4. The set explicitly excludes unrelated delivered records, unrelated accepted records merely visible through `project` or `read`, FIFO/in-flight commands, and rejected/held records.
5. Free-typed values outside that conductor-supplied set bounce before append. If a record class requires a parent and no active-lineage candidate exists, the candidate bounces structurally rather than widening the set.

This preserves the locked parent-picker property: the lane can select within a conductor-selected set, but cannot invent a false authority edge. PLAN acceptance criteria should include a stale-positive rejection, stale-negative re-render, outside-set rejection, and an unrelated delivered/accepted relay excluded from the candidate set.

Answer to ask (b): hold PLAN to the five-point active-lineage derivation above, not the broad delivered/accepted horizon wording.

### F-S3-M1-3 - Lineage field homes are headers, except the m-1 envelope/system fields

Approved with this split:

| surface | m-1-approved home |
|---|---|
| `PARENT_DISPATCH_ID` | locked `parent_picker` home; conductor-supplied candidate set plus default; do not treat it as free text |
| `schema_version` | envelope only, stamped by the courier/store |
| `FROM`, `ROLE`, `relay_id`, `DISPATCH_ID`, timestamp/certification | system envelope fields per m-1 |
| `record_kind` | header slot, system-validated by S3/S2 taxonomy |
| `DESIGN_DOC_ID`, `DESIGN_LOCK_ID`, `DESIGN_RECORD_KIND`, `DESIGN_REVIEW_VERDICT` | headers with m-2 lineage roles |
| grant, merge-claim refs, action refs, scope/fold rows | headers; structured values use the D-2 canonical JSON string carrier |
| record payloads that are not header fields | body JSON/text as already approved by the owning record class |

Answer to ask (c): the new lineage fields named in the S3 packet may live in headers. Do not duplicate envelope/system fields in headers, and do not relax the `PARENT_DISPATCH_ID` picker.

### F-S3-M1-4 - Migrator wrap point is above the raw store

Approved if the migrator is an engine/read-view facade over raw store reads:
- `store.Read`, `store.Records`, and `store.Project` keep their current raw meanings. `Read`/`Records` verify and return immutable canonical records with original `Envelope.SchemaVersion`; `Project` remains mailbox relay-id projection.
- A new engine/view layer may call raw `Read` or `Records`, apply `migrate.Apply` to copies, and feed lineage, form validation, projection rebuild helpers, replay, or seat-facing current-view responses.
- No stored byte mutates, no checksum is recomputed for a migrated view, and migrated output must retain or expose the source schema version for diagnostics.
- Store-level checksum/quarantine errors win before any migration attempt. Unknown-future, unversioned, and missing-chain refusals remain typed, path-free refusals.

Answer to ask (d): the sanctioned wrap point is a named engine read facade in front of consumers, not inside `store.Store` and not a changed raw store contract. If S3 wants a seat-visible `read` response to return current views rather than canonical raw records, name that as a channel/view response above store `Read`, with the raw store contract preserved.

## Per-item verdicts

| item | verdict | m-1 answer |
|---|---|---|
| 1. D-7 tables | approve-conditional | Internal store-derived read model approved under F-S3-M1-1. Name the query/cache semantics; no public store API or persisted table authority. |
| 2. PARENT `parent_picker` realization | approve-conditional | Free-typed outside-set reject approved, but candidate derivation must be active-lineage scoped per F-S3-M1-2. Broad delivered/accepted horizon is not approved. |
| 3. lineage-field homes | approve-conditional | Header homes approved for DESIGN/grant/merge/action fields; envelope/system fields and parent picker stay in their locked homes per F-S3-M1-3. |
| 4. D-9 read-facade wrap | approve-conditional | Engine/view facade above raw store approved per F-S3-M1-4; store `Read`/`Records`/`Project` semantics stay raw. |
| 5. canonical JSON string header carrier | approve | Approved. Canonicalization must happen before seal/commit; noncanonical equivalent encodings must be normalized or rejected so checksum identity is deterministic. Envelope and checksum algorithms stay unchanged. |
| 6. `record_kind` token additions | approve | No new S3 token is approved or needed here. S2's five tokens stand; D-6 precise classes ride `record_kind` plus `LineageRole`, while the S1 authority superset remains the floor. Any new token routes back to m-1. |
| 7. registry member replacement posture | approve | Fresh `store.Init` posture approved. Changed registry member means changed config digest and phase-0 wall on existing stores. The future config-change record remains OUT for S3 and must carry m-1 fidelity on any new `record_kind`. |

## Dispatch condition

The m-1 fidelity gate is satisfied for PLAN drafting if the gated PLAN carries F-S3-M1-1 through F-S3-M1-4 and the per-item table above. No separate m-1 narrow re-review is required when the PLAN implements those conditions verbatim. Route back to m-1 if the PLAN retains the broad delivered/accepted horizon, changes the raw store verbs, persists D-7 tables as authority, adds a `record_kind`, or moves any envelope/system field into headers.

ACTIONS_GIT_REF: none - report-only fidelity relay; no code or tracked design file changed by m-1.
FINAL_GIT_STATUS_SHORT: clean before relay write; expected git status remains clean because `.relays/` is gitignored.
