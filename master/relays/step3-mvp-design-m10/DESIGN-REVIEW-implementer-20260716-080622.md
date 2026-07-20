## DESIGN-REVIEW - m-10 stage-1 IPC/manifest seam r4 must revise: replacement connector has no m-10-authoritative current-epoch bootstrap

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r5
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - the remaining finding is a bounded m-10-owned CTRL-C bootstrap omission; Master+VP retain interface-lock authority
GRILL_REQUIRED: no - this re-review does not reopen the operator-locked topology, threat ceiling, F59, F60, or the planner's revocation-first choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260716-093500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260716-080622.md
SUBJECT: must-revise exact 5e16b2fa... - R4-F1/F2 close, but co-restart discards the connector that received E+1 and no m-10-authored current-epoch bootstrap exists for its replacement

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner - I reviewed the exact revised bytes at SHA-256 `5e16b2fa30e68a74e8652c4ca56b2bffc9efdd0c4a51e31b81619453d0b86b19`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, the r7 basis hash, and both consumed sibling hashes pass. R4-F1 and R4-F2 close at the locked confusion-not-malice threat ceiling. One new implementation blocker remains in the connector half of the canonical replacement sequence.

## Finding

### R5-F1 - A newly spawned connector has no legal m-10-authoritative current epoch

The contract makes every DATA-P socketpair generation-paired: replacement of either owner replaces BOTH worker and connector (`2026-07-16-mvp-ipc-manifest-seam-contract.md:25`). In the canonical transition, m-10 sends `epoch_update(E+1)` only “where a connector survives,” then reaps the paired owners and spawns the replacement pair (`:84-89`). Therefore, any surviving connector that received E+1 in step 2 is deliberately destroyed in step 3. If the connector itself failed, there was no surviving connector to receive the update at all.

The replacement connector has no specified bootstrap that repairs this:

- the exact hello/assign handshake is the worker's CTRL-W protocol; `assign` is sent only after lease bind, includes worker/broker fields, and is explicitly the sole generation-ID reveal (`:60`);
- the connector clause says only that m-8 follows SPAWNING/READY/FAILED/TERMINATED over CTRL-C and has no lease; it defines no connector hello reply, `connector_assign`, epoch snapshot, or acknowledgment (`:64`);
- step 4 says the replacement pair reaches READY on hello but “No assign yet,” and step 5 sends only the worker `assign` (`:88-89`);
- general distribution sends `epoch_update` “on increment” (`:92`), but E+1 was incremented before the replacement connector existed; and
- m-8's current-epoch authority may advance ONLY from m-10, never from the peer's DATA-P claim, a guess, or a default (`:93`).

Thus the new connector can be marked READY without possessing the authoritative current epoch it must use to reject stale DATA-P requests. Its first provider request is undecidable under the pinned source-specific authority rule: accepting a worker-presented epoch would violate `:93`, while rejecting/holding forever leaves the replacement pair unable to serve.

Required return: pin an exact CTRL-C bootstrap/resynchronization exchange for every newly spawned or reconnected connector. The m-10 reply must be sourced from durable state and carry the current `{run_id, turn_epoch}` plus only the other identity/state fields m-8 actually needs; define the connector acknowledgment and require successful installation before connector READY, worker admission, DATA-P acceptance, or any provider-send path. State explicitly that a worker-presented DATA-P epoch cannot bootstrap this cache and that no provider payload or credential bytes cross CTRL-C. Fold the exchange into the canonical steps 4-5, app-main recovery/re-establishment where applicable, and the m-8 consumer-confirmation row. If the design instead intends the existing `assign` family to cover m-8, split and type its connector form explicitly rather than relying on the worker-only CTRL-W definition.

## What closes from review r4

- R4-F1 closes under the architecture's locked confusion-not-malice ceiling: the retirement transaction pre-allocates G+1, the broker receives the schema-exact `{G+1,E+1,unleased}` tuple, and G+1 remains unavailable to a confused worker until post-lease `assign`. `ALLOCATED` and `RETIRED_PENDING_REAP` now make the durable state sequence explicit; same-epoch candidate wash-out installs a new unrevealed generation.
- R4-F2 closes: the fixed-width decimal-TEXT CHECK now rejects values above `18446744073709551615`, the bound applies on every insert/update, m-10 increment ownership is limited to `turn_epoch`/`state_seq`/`control_generation`, and broker-owned counters are validate-and-store-only.
- Prior broker CI, manifest/F63, DATA-P pairing, PREPARING-ledger, attempt/event ingress, canonical wire-counter, and shared crash-order repairs remain present.
- This verdict does not strengthen the accepted same-user threat model: the interregnum proof is confusion-resistant, not a malicious-guessing or forgery-proof claim.

## Gate disposition

This verdict is byte-bound to `5e16b2fa30e68a74e8652c4ca56b2bffc9efdd0c4a51e31b81619453d0b86b19`. Any revision requires a new SHA and fresh uniquely-parented DESIGN-REVIEW. Consumer confirmations, the Master+VP interface lock, stage-3/5 designs and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact design SHA-256 recomputed: `5e16b2fa30e68a74e8652c4ca56b2bffc9efdd0c4a51e31b81619453d0b86b19`.
- Ratified amendment SHA-256 recomputed: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Pair-approved m-7 r6 SHA-256 recomputed: `f072bd996da0c85b1be9b67fad880e7395ab03de7142cb87fc6864d4f67a100e`.
- Pair-approved m-3 r2 SHA-256 recomputed: `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44`.
- Incoming DESIGN relay exact-file lint: OK.
- SQLite exact-bound probe: max `18446744073709551615` accepted; first out-of-domain `18446744073709551616` rejected with the proposed CHECK.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`
RELAY_LINT: OK - exact-file mode exit 0
Next requested action: m-10.planner folds R5-F1 into fresh design bytes, recomputes the SHA-256, files a new uniquely-parented DESIGN relay, and requests fresh m-10.implementer review; do not route consumer confirmations on the current bytes.
