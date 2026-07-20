## DESIGN-REVIEW - m-10 r24 must revise: the total cancellation-result table names a durable key it cannot construct

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r25
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - one bounded key-constructibility and table-domain correction inside the accepted one-way cancellation consumer; no topology, policy, residual-risk, or product-semantic choice is reopened
GRILL_REQUIRED: no - this review does not reopen topology, policy, threat ceiling, F59, F60, revocation-first, or an operator-locked choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260717-215500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260718-015500.md
SUBJECT: MUST-REVISE exact r24 69264802... - the five-branch ordering closes R24-F1 behaviorally, but its durable equivalence key includes an uncarried/unpersisted cancellation identity, omits reported turn_epoch, and leaves states outside the five branches

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner - I reviewed the exact r24 bytes at SHA-256 `69264802ba370e39a2c2931fb32bcc692738904dfc15c0fffda06575c0af5c4f`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, m-8 r8 at `b4f9146a67db4b515bd2a2d56e4ebe7793a8b1ab4db65c3f2dd74298a7e1cf2d`, and m-9's cancellation confirmation at `e9074dbf9b77fcd66fab1372382e088a32490ca7e7c4da44143985af0dc332f9` pass their identity checks.

R24-F1 closes behaviorally: the consumer is explicitly one-way; equivalent committed delivery precedes PENDING/stale and cannot downgrade terminal `CANCELLED`; conflict preserves the proven terminal; fresh unmatched and stale-uncommitted have honest local effects; and the fault diagnostic is explicitly non-durable. The accepted intent predicate, crash separation, and pre/post two-view split remain intact.

One constructibility/coverage blocker remains in the new table.

## Finding

### R25-F1 - the claimed durable equivalence key is not derivable or persisted, and the five branches are not total over the received frame

r24 defines the durable key as `{attempt_id, cancel_point, the correlated cancellation identity}` and branch (1) asks whether the already-terminal row has “the same key” (`m-10 :61`). The actual producer frame carries only `{attempt_id, turn_epoch, disposition: cancelled(cancel_point)}`. It carries no `cancellation_id`. The §F `provider_attempts` row records `attempt_id`, `turn_epoch`, terminal state, and `cancel_point`, but no cancellation identity (`:232`). The separate `cancellations` row has `{cancellation_id, target, epoch at request, disposition}` with no uniqueness or idempotency rule over target+epoch (`:237`).

Therefore the consumer cannot mechanically select “the” correlated identity when two pending cancellation rows target the same attempt/epoch, and after terminal commit it has no specified persisted relation by which to compare the identity on a duplicate. The table names a durable fact the schema does not make durable.

The key also omits the frame's reported `turn_epoch`, despite r24's predecessor finding requiring the reported epoch in the durable fact. Because branch (1) precedes stale checking, a frame with the same `attempt_id`/`cancel_point` but a different reported epoch can be treated equivalent unless epoch equality is part of the equivalence predicate. Post-epoch delayed delivery should be idempotent because its **reported historical epoch still equals the committed attempt/cancellation epoch**, not because epoch is ignored.

Finally, the table calls itself TOTAL but does not state a precondition or branch for:

- a valid-shaped `cancelled` result naming an unknown `attempt_id`;
- a row already terminal in a non-`CANCELLED` state; or
- malformed identity/counter input before row lookup.

Branch (4) only says “an in-flight row” parks UNKNOWN, so it does not define these cases. The fixture phrase “zero wire effect” is also ambiguous because the incoming `attempt_result` is itself a wire frame; the intended receiver effect is no outbound response.

Required revision:

1. Make correlation constructible without inference. The narrow no-wire-change option is to make pending attempt cancellation unique/idempotent at an exact durable key such as `(target_kind=attempt, target_id=attempt_id, epoch)` and persist the selected `cancellation_id` (or an equivalent exact foreign-key fact) on the terminal `provider_attempts` row. If identity must instead ride the frame, that is an m-8-owned message-shape delta and requires routing.
2. Include the frame's reported `turn_epoch` in semantic equivalence and compare it to the persisted attempt/cancellation epoch. Keep equivalence-first precedence for delayed delivery only when that historical epoch matches; a different reported epoch is not equivalent.
3. State the table's validation/row-state domain or add ordered dispositions for malformed, unknown-attempt, and already-terminal-non-CANCELLED cases. For every case, pin row mutation and receiver response effect. Preserve terminal facts.
4. Replace “wire effect NONE” with the exact claim “no outbound response frame” so the fixture still observes the inbound CTRL-C notification.
5. Add cuts for duplicate pending cancellation requests/ambiguous correlation, same attempt+point with a different reported epoch, unknown attempt, and already-terminal non-CANCELLED row, in addition to the existing equivalent/conflict/fresh/stale cuts.

This does not reject the five chosen state dispositions. It makes their key obtainable from the message plus durable state and makes the “TOTAL” claim true over the stated domain.

## What closes

- The one-way/no-reply correction and non-durable diagnostic wording are accepted.
- Equivalence-first idempotency, terminal preservation, fresh-unmatched UNKNOWN parking, and stale-uncommitted consume/drop are directionally correct.
- The explicit-cancellation provenance predicate, raw-closure UNKNOWN rule, two-view split, attempt accounting, and D-5 turn-level composition remain accepted.
- This review remains confined to the declared §B.1 cancellation-consumption block; no unrelated design surface is reopened.

## Gate disposition

MUST-REVISE is byte-bound to `69264802ba370e39a2c2931fb32bcc692738904dfc15c0fffda06575c0af5c4f`. The r24 SITREP and scoped letter-rebind must not route on these bytes. The m-8 fresh final-byte review, m-9 r6 fold, remaining reciprocal confirmations, Master+VP interface lock, stage-3/5 closure and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

Any revision requires a new SHA and a fresh uniquely-parented m-10.implementer DESIGN-REVIEW. No human gate is needed if the revision preserves the accepted five state dispositions and only binds their key/domain to the existing message and store.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `2a6e6286ee747fcfce1965f18546692815bc30f96f6eb54b4cea81cf108c76cb`.
- Exact m-10 r24 SHA-256 recomputed: `69264802ba370e39a2c2931fb32bcc692738904dfc15c0fffda06575c0af5c4f`.
- Exact m-8 r8 SHA-256 recomputed: `b4f9146a67db4b515bd2a2d56e4ebe7793a8b1ab4db65c3f2dd74298a7e1cf2d`.
- Exact m-9 R7-F2 confirmation relay SHA-256 recomputed: `e9074dbf9b77fcd66fab1372382e088a32490ca7e7c4da44143985af0dc332f9`.
- Incoming DESIGN exact-file lint: OK.
- Targeted constructibility sweep: m-10 `:39-42,61,109,223-242`; m-8 r8 `:76,80,90-100,226`; m-9 confirmation `Leg-1` two-cut composition.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file proof reported `OK master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260718-015500.md`; root-wide legacy findings are outside this relay.
Next requested action: m-10.planner binds cancellation identity and reported epoch to an exact durable correlation key, states or completes the table domain, disambiguates receiver response effect, adds the key/domain cuts, recomputes the SHA-256, and requests a fresh uniquely-parented review. Do not file the r24 SITREP or route its scoped rebind.
