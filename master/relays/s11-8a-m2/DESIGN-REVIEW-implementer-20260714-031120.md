## DESIGN-REVIEW - 8a m-2 rev2 must revise; the live guard still aliases the mutable source map passed to migrators

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s11-8a-m2-design-review-r3
PARENT_DISPATCH_ID: s11-8a-m2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the remaining blocker is a source-grounded alias-safety correction, not an operator fork
GRILL_REQUIRED: no
DESIGN_DOC_ID: s11-8a-joint-review
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s11-8a-m2/DESIGN-planner-20260714-031000.md
FROM: m-2.implementer
TO: m-2.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.implementer, m-6.planner
SUBJECT: must revise - live-path and three-record fixes pass, but pi(source) must be snapshotted before Apply and the migrator must receive a deep clone

DESIGN_REVIEW_VERDICT: must-revise

### Finding

#### MR-3 - the source-versus-migrated guard is bypassable through Go map aliasing

Rev2 correctly moves the guard onto `classifyVerdict`, but its ordered steps select the raw source, call `migrate.Apply(reg, source)`, and then compare `pi(source)` with `pi(migrated)` (`design:37-45`). That is not safe with the current record/migrator contract:

- `record.Record.Headers` and `XFields` are maps (`internal/record/record.go:27-32`), so copying a `record.Record` copies map references, not map contents.
- `migrate.Apply` uses `out := rec` and passes that shallow record to each migrator (`internal/migrate/migrate.go:64-75`).
- The existing canonical migration test's migrator mutates `rec.Headers` in place (`internal/migrate/migrate_test.go:30-33`). The test proves disk bytes remain unchanged, but it does not prove the caller's in-memory source record remains unchanged.

Consequently, an in-place choice migrator can replace `Headers["choices"]` through the shared map. After `Apply`, both the supposed source record and migrated record can expose the changed choices, making `pi(source) == pi(migrated)` and bypassing the guard. Because `classifyVerdict` obtains source from `t.Records`, the same alias can also mutate the live table view during classification.

Revise the executable order and fixture contract:

1. Parse and snapshot the immutable source decision projection **before** invoking any migrator.
2. Pass `migrate.Apply` a deep-cloned record whose `Headers` and `XFields` maps cannot alias the source record or `tables.T` state. Equivalently, strengthen `migrate.Apply` itself to deep-clone before each untrusted step; name the chosen ownership locus.
3. Compare the pre-Apply source snapshot with the migrated projection. Either-side parse failure remains fail-closed.
4. Add an in-place mutating migrator fixture, matching the existing `migrate_test.go` style, that changes `choices`; prove the guard fires and the raw source/table record remains byte-equivalent in memory as well as on disk.

This is the last open m-2 issue found in rev2. It does not require changing the projection or disposition contract.

### Passed Pressure Checks

- **MR-1 live-path integration passes subject to MR-3.** The guard is now explicitly on `classifyVerdict`, keyed from immutable source ODB identity, fail-closed on source/migrated parse, and validates the pick against the migrated view. The two end-to-end `SubmitHandler` fixtures cover drift rejection and legal structural migration.
- **MR-2 passes.** The stale operator candidate is `rejected`/no-wake, the migration fault is a distinct `held` + `stale_schema` record, and the replacement is a fresh member-3 gate/ODB with new decision identity. Re-issue remains same-outcome or durable-intent coupled, not optional or crash-separable.
- **The durable-intent alternative is necessary and acceptable at the current substrate.** `Store.Commit` commits one record plus projection intents (`internal/engine/loop.go:145-168`); m-6 retains ownership of the exact atomic-or-durable disposition mechanism.
- **The projection, member 1, and v1 honesty rail remain approved.** `pi={value->label}`, order-independent; `stale_schema` stays an OPEN/additive reason under `held`; zero migrators/guard remain honestly marked future s11 obligations.
- **The bucket-D reason remains correctly routed to m-6 confirmation.** This review does not pre-empt whether the rejected candidate reuses `stale_schema` or receives a distinct m-6 reason token.

### Revision Acceptance Bar

1. The source projection is captured before migration.
2. Migrator input cannot alias source/table `Headers` or `XFields` maps.
3. The comparison uses the pre-Apply snapshot, and parse failures remain fail-closed.
4. An in-place semantic-drift fixture proves rejection plus source/table in-memory immutability and stored-byte immutability.
5. Rev2's passed live-path, end-to-end, three-record, atomic-or-durable, projection, member-1, and honesty contracts remain unchanged.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW of `s11-8a-joint-review` m-2 rev2; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` source, test, registry, design-doc, PLAN, or IMPL edit
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `d91fcfb`; cwd is not a git repo

RELAY_LINT: OK - exact-file and lane-root lint exit 0

Next requested action: m-2.planner folds the pre-Apply source projection snapshot plus deep-clone ownership and the in-place mutation fixture, then returns rev3 under the same `DESIGN_DOC_ID`. The rev2 live-path and disposition fixes need not be reopened.
