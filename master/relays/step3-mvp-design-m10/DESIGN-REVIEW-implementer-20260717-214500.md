## DESIGN-REVIEW - m-10 r23 must revise: the cancellation-intent fold needs a total one-way result-consumption table

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r24
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - one bounded protocol-totality correction inside the accepted cancellation-intent predicate; no topology, policy, residual-risk, or product-semantic choice is reopened
GRILL_REQUIRED: no - this review does not reopen topology, policy, threat ceiling, F59, F60, revocation-first, or an operator-locked choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260717-214000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-214500.md
SUBJECT: MUST-REVISE exact r23 07db105f... - R23-F1/F2 are substantively closed, but attempt_result is a one-way notification while the new predicate leaves stale rejection, fault recording, and equivalent-duplicate precedence undefined

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner - I reviewed the exact r23 bytes at SHA-256 `07db105f5e636936bbdc1b9dc99c48ea00a30f5c8a8b1ea3694b34353da817eb`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, m-8 r8 at `b4f9146a67db4b515bd2a2d56e4ebe7793a8b1ab4db65c3f2dd74298a7e1cf2d`, and m-9's cancellation confirmation at `e9074dbf9b77fcd66fab1372382e088a32490ca7e7c4da44143985af0dc332f9` pass their identity checks.

The prior findings close substantively. R23-F1 now makes cancellation provenance-based: only a matching current-epoch pending cancellation fact can mint terminal `CANCELLED`, while bare DATA-P closure or worker crash stays on the binding UNKNOWN path. R23-F2 now makes `pre_transport` the second m-8-view-only close and pins `post_invocation` to the m-9 `stream_cancelled` view. The four requested cuts are named.

One protocol-totality blocker remains in the new R23-F1 sentence block.

## Finding

### R24-F1 - the new predicate promises unreachable rejection/recording semantics and can downgrade an equivalent duplicate

`attempt_result{attempt_id, turn_epoch, disposition}` is a one-way CTRL-C notification. The m-8 owner bytes explicitly record that no reply-class receipt/ack exists for the later `attempt_result` terminal close (`m-8 r8 :80`); m-10 likewise defines no `attempt_result` response family in §A.2 or §B.1.

r23 nevertheless says a stale-epoch `cancelled` takes “the ordinary stale rejection” and an unsolicited/mismatched result parks the row UNKNOWN “+ a typed fault is recorded” (`m-10 :61`). Neither effect is executable as written:

- there is no response carrier or closed reject token for a one-way `attempt_result`, so “stale rejection” cannot mean the reply-class `attempt_open_reject` or `turn_reject` families without crossing message ownership;
- §F has no fault/evidence table (`:227-242`), and the only general events journal is explicitly deferred (`:225`), so “recorded” names no authoritative target, fields, transaction, or retention rule; and
- the evaluation order is not total. The first matching current-epoch/PENDING `cancelled` commits terminal `CANCELLED`. A semantically equivalent duplicate delivered after that cancellation fact is no longer PENDING, or after epoch advance, fails the stated predicate and falls into mismatch/stale handling. The mismatch branch says the row parks `UNKNOWN_PROVIDER_OUTCOME`, contradicting the same paragraph's guarantee that a committed terminal `CANCELLED` row is never re-parked during retirement. “Only a genuinely ... conflicting result ... parks” does not resolve this because semantic equivalence, conflict, and precedence over epoch/PENDING checks are undefined for this family.

Required revision: make cancellation-result consumption a total ordered table, scoped to the existing one-way message:

1. Name the durable equivalence key and fact (`attempt_id`, reported epoch, `cancel_point`, and the correlated cancellation identity/fact).
2. Give an already-committed semantically equivalent `CANCELLED` result precedence over the PENDING and stale checks: idempotent no-op, no terminal downgrade, including delayed delivery after epoch advance. Define a conflicting duplicate separately and preserve the already-proven terminal fact.
3. Define fresh matching PENDING, fresh unmatched/unsolicited, and stale-uncommitted branches with exact state mutation. If stale means consume-and-drop with zero response and zero mutation, say that; if a response is intended, add an owner-real reply family and route the producer half.
4. Bind the “typed fault” to an existing exact disposition and durable effects, or state that it is non-durable/no-record and remove the unsupported recorded-evidence claim. Do not invent a durable record outside the §F schema.
5. Add fixtures for equivalent duplicate before and after epoch advance, conflicting `cancel_point`, stale uncommitted result, and fresh unsolicited/mismatched result. Every fixture must assert the wire effect (normally none for this one-way family) and that a terminal `CANCELLED` row is never downgraded.

This finding does not reject the cancellation-intent predicate, terminal `CANCELLED`, the two-view split, or the four prior boundary cuts. It only makes the newly introduced consumer branches reachable, idempotent, and state-total.

## What closes

- R23-F1's intent provenance and crash/channel-loss separation are correct.
- R23-F2's pre-transport/post-invocation view split is exact against m-9's confirmation.
- The `cancel_point` domain, attempt counting, terminal-row preservation goal, and D-5 turn-level composition remain accepted.
- The r23 edit is confined to the declared §B.1 cancellation sentence block; no unrelated design surface moved in this review.

## Gate disposition

MUST-REVISE is byte-bound to `07db105f5e636936bbdc1b9dc99c48ea00a30f5c8a8b1ea3694b34353da817eb`. The r23 SITREP and scoped letter-rebind must not route on these bytes. The m-8 fresh final-byte review, m-9 r6 fold, remaining reciprocal confirmations, Master+VP interface lock, stage-3/5 closure and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any revision requires a new SHA and a fresh uniquely-parented m-10.implementer DESIGN-REVIEW. No human gate is needed if the revision preserves the accepted provenance/two-view semantics and only totalizes the one-way consumer branches.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `a861e30df864808f8d1f6f627c166979b07d5a8fc39d34e44effeea8d379ec30`.
- Exact m-10 r23 SHA-256 recomputed: `07db105f5e636936bbdc1b9dc99c48ea00a30f5c8a8b1ea3694b34353da817eb`.
- Exact m-8 r8 SHA-256 recomputed: `b4f9146a67db4b515bd2a2d56e4ebe7793a8b1ab4db65c3f2dd74298a7e1cf2d`.
- Exact m-9 R7-F2 confirmation relay SHA-256 recomputed: `e9074dbf9b77fcd66fab1372382e088a32490ca7e7c4da44143985af0dc332f9`.
- Incoming DESIGN and m-9 confirmation exact-file lint: OK.
- Targeted one-way cancellation sweep: m-10 `:61,71,77-84,103,225,227-242`; m-8 r8 `:76,80,82-100,226`; m-9 confirmation `Leg-1` two-cut composition.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260717-214500.md`; root-wide legacy findings are outside this relay.
Next requested action: m-10.planner totalizes the one-way `cancelled` result consumer with equivalence-first idempotency, exact stale/unmatched/conflict state effects, and a reachable fault disposition; adds the duplicate/stale fixtures; recomputes the SHA-256; and requests a fresh uniquely-parented review. Do not file the r23 SITREP or route its scoped rebind.
