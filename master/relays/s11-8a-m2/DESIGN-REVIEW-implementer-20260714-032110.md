## DESIGN-REVIEW - 8a m-2 rev3 approved; frozen-choice migration guard is live-path, alias-safe, and disposition-compatible

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s11-8a-m2-design-review-r4
PARENT_DISPATCH_ID: s11-8a-m2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the m-2 contract is mechanically complete; remaining joint confirmation belongs to m-6
GRILL_REQUIRED: no
DESIGN_DOC_ID: s11-8a-joint-review
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/s11-8a-m2/DESIGN-planner-20260714-031600.md
FROM: m-2.implementer
TO: m-2.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.implementer, m-6.planner
SUBJECT: approve m-2 members 1 and 2 - MR-1 through MR-3 closed; full 8a joint lock still requires m-6 member-3 confirmation

DESIGN_REVIEW_VERDICT: approve

### Verdict

Approve the m-2 leg of `s11-8a-joint-review`: member 1 (`stale_schema` home/version class) and member 2 (frozen-choice migration interaction) are complete at E1. Rev3 closes the final alias-safety blocker while preserving the previously passed live-path, projection, and disposition contracts. This is not a solo approval of the full joint 8a lock; m-6's member-3 confirmation remains required.

### Closed Findings

- **MR-1 closed - guard is on the production verdict path.** `classifyVerdict` obtains the source ODB from the immutable table record, applies the threaded migration registry, checks projection invariance, and validates the pick against the migrated view. Raw `tables.T` choice validation is no longer an allowed bypass. End-to-end `SubmitHandler` fixtures cover semantic drift rejection and legal structural migration (`design:39-51,73-79`).
- **MR-2 closed - record outcomes are byte-distinct.** The stale operator candidate is `rejected`/no-wake, the migration fault is a separate `held` + `failing_edge: stale_schema` record, and the replacement is a fresh member-3 gate/ODB with a new decision identity. Same-outcome or durable-intent coupling preserves m-6's crash-safety requirement (`design:53-67`).
- **MR-3 closed - migrator aliasing is prevented.** The guard snapshots `pi(source)` as copied value data before any callback; `migrate.Apply` deep-clones `Headers` and `XFields` before each migrator step; comparison uses the pre-Apply snapshot. This prevents the current in-place migrator style from mutating the caller's source or live table view (`design:41-51`).
- **MR-3 fixture is sufficient.** `NF-migrate-choice-inplace-mutate` matches the observed in-place map mutation style and proves all three required outcomes: the drift guard fires, the in-memory source/table bytes remain equivalent, and stored bytes remain immutable (`design:81`).

### Approved Contract

1. `stale_schema` is a system-stamped OPEN/additive-MINOR `failing_edge` reason under `delivery_state: held`; it is neither a fourth delivery state nor a new enum.
2. Frozen choice identity is `pi = {value -> label}`; value and label are decision-bearing, order is not, and representational columns are excluded.
3. Source ODB identity gates the check; source projection is snapshotted before migration; `Apply` deep-clones map fields; either-side parse and migration failures fail closed.
4. Operator resolution validates against the guarded migrated view whose projection equals the frozen source projection, never against live registry options. Additive widening is not retroactive.
5. The stale reply, held migration fault, and replacement gate remain distinct and crash-coupled. A changed choice set never wakes or auto-resolves the old decision.
6. v1 honesty remains explicit: zero migrators and no guard fire at `d91fcfb`; these are RED-first s11 obligations.

### Remaining Joint Boundary

- m-6.implementer confirms member 3 consumes the typed staleness signal with the approved new-decision identity and atomic-or-durable re-issue behavior.
- m-6 owns the exact bucket-D reason token for the rejected stale candidate and the concrete durable-intent versus same-outcome implementation.
- Only after that confirmation may master integrate the three-member contract and unlock s11 T6. This relay does not itself close the full joint gate or authorize PLAN/IMPL.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW approval of `s11-8a-joint-review` m-2 rev3; wrote this relay and appended `master/relays/INDEX.md`; no `frank/` source, test, registry, design-doc, PLAN, or IMPL edit
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `d91fcfb`; cwd is not a git repo

RELAY_LINT: OK - exact-file and lane-root lint exit 0

Next requested action: m-2.planner returns this approved m-2 leg into the joint review; m-6.implementer confirms member 3 and the remaining m-6-owned reason/disposition details; master then integrates the full 8a contract before s11 T6 locks.
