## DESIGN-REVIEW - m-10 stage-1 IPC/manifest seam r5 must revise: CTRL-C still discloses the pre-lease generation and app-main recovery has no realizable pair handover

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r6
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - both remaining findings are bounded m-10-owned IPC/supervision contradictions; Master+VP retain interface-lock authority
GRILL_REQUIRED: no - this re-review does not reopen the operator-locked topology, threat ceiling, F59, F60, or revocation-first choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260716-095500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260716-081505.md
SUBJECT: must-revise exact bc7298b8... - ordinary replacement connector bootstrap closes R5-F1, but epoch_update contradicts the sole-generation-reveal proof and app-main recovery cannot re-establish spawn-only child channels as specified

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner - I reviewed the exact revised bytes at SHA-256 `bc7298b88ff74e76d57a2f574c638163e08cbfc4ed5c3803e3bc21976cce00a5`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, the r7 basis hash, and both consumed sibling hashes pass. The new connector bootstrap closes R5-F1 for the ordinary canonical replacement path. A fresh whole-contract sweep found two exact blockers, including one contradiction I incorrectly overclosed in review r5.

## Findings

### R6-F1 - `epoch_update` reveals G+1 before lease bind, contradicting the interregnum proof

The worker handshake states that post-lease CTRL-W `assign` is the ONLY frame that reveals `generation_id` (`2026-07-16-mvp-ipc-manifest-seam-contract.md:60`). The interregnum proof relies on the stronger exact fact that G+1 is revealed to exactly one process, only after step 5 (`:87`). The new `connector_assign` deliberately omits `generation_id` so it remains legal before lease bind (`:65,89`).

But the general distribution message is still exactly `epoch_update{run_id, turn_epoch, generation_id}` (`:93`), and canonical step 2 sends that message to a surviving connector before step 3 reaps it (`:87-88`). On worker failure or requested replacement, that old connector therefore learns G+1 during the supposedly unrevealed interregnum. The document simultaneously asserts both “only worker assign reveals generation_id” and “CTRL-C epoch_update reveals generation_id”; the information-flow proof is false on the current bytes even at the locked confusion-not-malice ceiling.

Required return: make the m-8 epoch-authority message generation-blind unless m-8 has a separately proven need for G. The direct repair is `epoch_update{run_id, turn_epoch}` on CTRL-C while retaining the full generation-bearing `epoch_state` on the broker feed. Sweep §B.1, §B.4, and §H so the sole-reveal invariant is exact. If m-8 truly needs G, replace the secrecy proof with a complete role/channel non-attachability proof and route the changed field requirement through m-8 confirmation; do not retain the contradictory absolute.

### R6-F2 - App-main recovery promises CTRL-C re-establishment without a channel or fencing path

CTRL-W, CTRL-C, and DATA-P are unnamed `socketpair` endpoints created immediately before m-10 spawns the children; only those processes inherit the endpoints (`:23`). The broker is explicitly different: its dial-in listener exists so a replacement app-main can adopt a surviving broker (`:21`; m-7 r6 `:181-203`). No listener, descriptor handoff, or other rendezvous exists for a replacement app-main to reach a surviving m-8/m-9 pair.

The new clause nevertheless says the connector bootstrap runs at every CTRL-C “re-establishment, including app-main recovery” (`:65`). The app-main crash clause only parks in-flight rows and says leases are not assumed carried over (`:80`); it does not require the orphaned pair to fail closed on CTRL loss, prove them terminal, retire the old generation, mint E+1, or spawn a fresh pair. This matters on both verticals: m-7 correctly suspends relay verbs when controller state supply disappears, but the old worker and connector still share DATA-P and the connector has no pinned rule forbidding provider send after CTRL-C EOF. Merely declining to assume the lease in the replacement process does not revoke authority already held by surviving children.

Required return: choose and pin one realizable MVP recovery path. The bounded spawn-only path is: CTRL-W/CTRL-C loss makes both children immediately fail closed (no DATA-P acceptance/provider send, tool execution, or broker use) and terminate; replacement app-main reconciles/proves the old pair terminal, applies the retirement transaction and E+1 fence exactly once, then creates a fresh CTRL-W/CTRL-C/DATA-P set, runs `connector_assign`/`connector_ready`, lease-binds, worker-assigns, and admits. State how PID reuse/orphan proof is avoided or bounded. If instead the pair survives and is adopted, define a named authenticated reconnect/rendezvous and controller-handover protocol rather than relying on spawn-only socketpairs. Also pin the initial E=1 path: run freeze must durably produce `{run_id, turn_epoch, run_manifest_digest}` before connector bootstrap, and `connector_ready` must precede worker lease/admission.

## What closes from review r5

- R5-F1 closes for an ordinary replacement once authoritative durable run state exists: `hello → connector_assign{run_id,turn_epoch,run_manifest_digest} → connector_ready` is a distinct CTRL-C exchange, the connector reaches READY only after installing/acking it, and worker admission/DATA-P/provider-send are gated on that ack.
- A worker-presented DATA-P epoch is explicitly unable to bootstrap or advance the connector cache; CTRL-C carries no provider payload or credential bytes.
- R4-F2 remains closed: fixed-width bounded decimal TEXT and split increment ownership are exact.
- R4-F1 is NOT closed on these bytes because of R6-F1. I overclosed it in review r5 by checking the new `connector_assign` field set without reconciling the pre-existing generation-bearing `epoch_update`; this relay corrects that verdict explicitly.
- Prior broker CI, manifest/F63, DATA-P pairing, PREPARING-ledger, attempt/event ingress, canonical wire-counter, and crash-outcome repairs remain present.

## Gate disposition

This verdict is byte-bound to `bc7298b88ff74e76d57a2f574c638163e08cbfc4ed5c3803e3bc21976cce00a5`. Any revision requires a new SHA and fresh uniquely-parented DESIGN-REVIEW. Consumer confirmations, the Master+VP interface lock, stage-3/5 designs and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact design SHA-256 recomputed: `bc7298b88ff74e76d57a2f574c638163e08cbfc4ed5c3803e3bc21976cce00a5`.
- Ratified amendment SHA-256 recomputed: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Pair-approved m-7 r6 SHA-256 recomputed: `f072bd996da0c85b1be9b67fad880e7395ab03de7142cb87fc6864d4f67a100e`.
- Pair-approved m-3 r2 SHA-256 recomputed: `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44`.
- Incoming DESIGN relay exact-file lint: OK.
- Exact contradiction sweep: `assign` sole reveal at `:60`; G+1 exactly-one-process proof + pre-reap CTRL-C update at `:87`; generation-bearing `epoch_update` at `:93`; spawn-only child socketpairs at `:23`; app-main recovery without pair disposition at `:80`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`
RELAY_LINT: OK - exact-file mode exit 0
Next requested action: m-10.planner folds R6-F1 and R6-F2 into fresh design bytes, recomputes the SHA-256, files a new uniquely-parented DESIGN relay, and requests fresh m-10.implementer review; do not route consumer confirmations on the current bytes.
