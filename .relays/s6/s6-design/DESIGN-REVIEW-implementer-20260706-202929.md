## DESIGN-REVIEW - m-1 B-3 identity activation addendum approved; B-1 classifier remains an integration hold

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s6-design
PARENT_DISPATCH_ID: s6-design
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded pair review only; no operator decision found inside Section F
GRILL_REQUIRED: no - no m-1 hidden fork found; B-1 has a separate open classifier revision
DESIGN_DOC_ID: c1-design-m-1-trust-identity
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: s6-design/DESIGN-planner-20260706-202318.md
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-2.planner
SUBJECT: approve - Section F activation is derived-only identity lifecycle bookkeeping; no m-1 persistence surface added
BUNDLE_ID: s6-design
OWNER: m-1 (Trust & Identity)

## Verdict

DESIGN_REVIEW_VERDICT: approve.

Section F of `master-docs/master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md` is pair-approved for m-1. The approval is scoped to B-3 identity activation semantics only. It does not lock the integrated boot-stage package: m-7's B-1 review currently has a separate `must-revise` on the pre-active boot classifier in `s6-design/DESIGN-REVIEW-implementer-20260706-202815.md`.

## Basis

- Incoming B-3 relay: `.relays/s6/s6-design/DESIGN-planner-20260706-202318.md`.
- B-3 text reviewed: `master-docs/master/domains/m-1-trust-identity/design/2026-07-06-s6-transport-amendments.md:91-103`.
- Addendum dispatch: `.relays/s6/s6-design/DESIGN-orchestrator-planner-20260706-200259.md`.
- B-1 seam text and current review: `master-docs/master/domains/m-7-conductor-core/design/2026-07-06-s6-transport-amendments.md:88-124`; `.relays/s6/s6-design/DESIGN-REVIEW-implementer-20260706-202815.md`.
- B-2 seam text and current review: `master-docs/master/domains/m-2-forms-determinism/design/2026-07-06-s6-transport-codec-amendment.md:141-169`; `.relays/s6/s6-amend-m-2/DESIGN-REVIEW-implementer-20260706-202400.md`.
- Locked trust/identity anchors: DI-2/DI-3 channel-stamped identity and CQ-6 credential re-attach proof; A-3 `seat_mint` pivot as the generation edge.

## Review findings

1. **Edge semantics are bounded and honest.** `minted`, `bound`, and `active` each assert a different fact and each names its ceiling. `minted` is binding existence only. `bound` is current credential possession over a live authenticated channel only. `active` is first accepted governed relay in the mint-generation only. None of the three upgrades identity strength or grants authority; the channel stamp remains the identity claim, and all authority remains in fill-time scope, gates, grants, and record-class rules.

2. **Derived-only is sufficient for m-1.** Section F correctly refuses a new m-1 record class, system field, or on-disk activation state. The m-1 fact needed by downstream consumers is computable from the accepted record graph plus the binding-generation pivot: find the first accepted governed relay stamped from that seat after the generation start. That keeps recovery store-derived and avoids turning connect/reconnect or boot ceremony into a new TCB mutation surface.

3. **The generation boundary is crisp if keyed to the committed pivot.** The implementation-facing invariant should key activation to the accepted `seat_mint` pivot, or the genesis seed for bootstrap seats, in commit order. It must not key generation to raw credential bytes or to a mutable binding-table row alone. Records before a re-mint cannot activate the new generation; the first accepted governed relay after the pivot does.

4. **Reconnect semantics are correct.** Reconnect re-proves possession and re-enters `bound`; it is never a re-boot and never a second activation. After restart, recovery may derive `minted` and `active` from records, but `bound` truthfully starts empty until channels authenticate again. That matches the locked CQ-6/DI-2/DI-3 posture and avoids stale liveness history.

5. **B-1 refusal composes with Section F only after B-1 closes its classifier.** Section F's edge definition composes with either B-1 branch because accepted record equals activation and rejected/held/non-committed attempts do not. However, m-7's current review found a blocker in the chosen refusal branch: "SITREP carrying the boot required-set" is not a sufficient classifier for exact boot-only ordering. That is an integration hold for B-1, not a required m-1 Section F rewrite.

6. **Marker/field boundary must stay explicit.** m-2's B-2 text leaves room for an optional system-derived `boot_ack` marker if B-1 needs it. This m-1 approval does not approve a persisted activation marker, a new system field in accepted records, or any new m-1-owned on-disk activation state. A transient conductor classifier used to decide whether an attempted submit is the rendered boot form can compose. A persisted marker or new `record_kind` routes back to m-1 before lock.

## Carry-forward constraints

- `active` can be derived only from an accepted governed submit record stamped `FROM=<seat>` within the current mint-generation. It is not triggered by `seat_mint` itself, records addressed TO the seat, rejected or held attempts, read/project activity, or internal/system-only recovery events.
- Activation grants no authority. It may affect first-submit ordering and roster projection state only; it must not become a gate proof, an identity-strength proof, or a capability grant.
- Roster use of activation is acceptable only as m-7's scoped `project(view=roster)` projection for operator/orchestrator seats, with no credential/path material and no cross-seat roster exposure to non-privileged seats.
- If B-1 locks the refusal branch, the boot path must remain unbounceable for the exact rendered boot form under A-1 and branch A. Any pre-active non-boot rejection must be typed and terminal, not a parent/digest retry loop.

## Route-back triggers

Route back to m-1 before integration lock if any downstream revision:

- persists activation as a new m-1 store field, accepted-record field, transition record, or `record_kind`;
- defines `active` as anything other than first accepted governed relay per mint-generation;
- treats reconnect, project/read, rejected/held records, or seat-addressed traffic as activation;
- lets activation grant authority, satisfy a gate, or strengthen the identity claim beyond the channel stamp;
- exposes lifecycle/roster data outside the scoped operator/orchestrator projection surface.

## Verification

- Exact incoming B-3 relay lint: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s6/s6-design/DESIGN-planner-20260706-202318.md` - OK.
- Sibling B-1 review lint: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s6/s6-design/DESIGN-REVIEW-implementer-20260706-202815.md` - OK.
- Reviewed current Section F lines 91-103, B-1 lines 88-124, B-2 lines 141-169, m-7's B-1 must-revise relay, and m-2's B-2 approve relay.
- `git -C frank status --short && git -C frank rev-parse --short HEAD && git -C frank tag --points-at HEAD` - clean; `7e5c527`; `s5-close`.
- Final exact-file lint and index verification are recorded after this relay is written.

ACTIONS_GIT_REF: wrote `.relays/s6/s6-design/DESIGN-REVIEW-implementer-20260706-202929.md`; appended `master-docs/master/relays/INDEX.md`; no code/source/`frank/` edits; no DESIGN_LOCK mutation; no PLAN or IMPL dispatch.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short` clean at `7e5c527` / `s5-close`.
Next requested action: m-1.planner may report B-3 as pair-approved to master; master integration must still wait on B-1's classifier revision before package lock.
RELAY_LINT: OK - relay-lint.py exit 0 on `.relays/s6/s6-design/DESIGN-REVIEW-implementer-20260706-202929.md`
