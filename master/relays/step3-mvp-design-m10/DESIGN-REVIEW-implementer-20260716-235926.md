## DESIGN-REVIEW - m-10 stage-1 IPC/manifest seam r8 must revise: CI-1 recovery installs the new epoch before transition reconciliation

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r9
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - the finding is a bounded m-10/m-7 recovery-order correction under the locked topology; Master+VP retain interface-lock authority
GRILL_REQUIRED: no - this re-review does not reopen the operator-locked topology, threat ceiling, F59, F60, or revocation-first choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260716-105500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260716-235926.md
SUBJECT: must-revise exact 4e555c23... - no-ID and F59 parking folds land, but CI-1 adoption/bootstrap installs E+1 before the pending transition is reconciled

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner - I reviewed the exact revised bytes at SHA-256 `4e555c2397c2fd0baf92348ea5c6e9443046433d93dd4426d88a4b3061460fc6`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, the r7 amendment hash, and both consumed sibling hashes pass. R8-F2 closes, and allocating the transition ID in the retirement transaction closes R8-F1's no-ID cut. The new broker-first common suffix, however, conflicts with the consumed CI-1 bootstrap/adoption contract and still allows E+1 to install before the transition handshake.

## Finding

### R9-F1 - CI-1 adoption/bootstrap bypasses the pending epoch-transition barrier

The exact crash cut is:

1. The retirement transaction commits E+1, G+1, and a `PROPOSED` transition ledger row (`2026-07-16-mvp-ipc-manifest-seam-contract.md:95`).
2. The app main crashes before sending `epoch_transition_propose` (`:111`), so no crossing set has frozen and no crossing rows exist.
3. Recovery branch (b) correctly observes the durable retirement, and the revised common suffix first adopts the surviving broker or spawns a fresh one, then reconciles the transition (`:84,87`).

But “adopt/spawn” consumes m-7 CI-1 in full. A fresh broker receives and installs the durable `epoch_state` during bootstrap (`2026-07-16-step3-mvp-transport-broker.md:199-200`); an adopted broker receives the replacement controller's current durable snapshot and installs it during adoption (`:201-202`). At this crash cut that snapshot is already `{G+1,E+1,unleased}`. Therefore the common suffix installs E+1 before §B.5 proposes the transition, freezes the crossing set, commits `CROSSERS_DURABLE`, and sends the install-enabling ack.

That contradicts both contracts' hard rule that E+1 reaches `INSTALLED` only after the crossing rows are durable and acknowledged (`m-10 :109-115`; `m-7 :129-146`). It also means branch (b)'s later read of `PROPOSED` is too late: the broker has already crossed the epoch outside the identified transition machine. A surviving broker could have admitted E operations before the app crash, so an empty or late snapshot cannot retroactively establish the mandatory crossing record.

Required return: separate **authenticated control-session establishment** from **epoch-state installation** for recovery with a pending transition. On adoption and fresh-spawn cuts where the current durable epoch has a non-terminal or ABORTED transition, the broker must remain suspended while m-10 queries/reconciles that transition by ID; only the exact `CROSSERS_DURABLE` ack may install E+1. Pin the corresponding CI-1 bootstrap/adoption message/order delta against m-7's consumer contract, including:

- `PROPOSED`/`PREPARING`/`CROSSERS_DURABLE` recovery cannot take the ordinary snapshot-install path before resume-or-abort;
- `ABORTED` recovery proposes a fresh ID for the same current E+1 while the broker remains suspended;
- `INSTALLED` may take the ordinary lifecycle continuation path without a new transition; and
- initial E=1 and same-epoch pre-lease wash-out remain separately defined, so the correction does not manufacture a transition where no epoch change occurred.

The return must also name the durable/acknowledged fact that permits m-10 to classify a transition as `INSTALLED`; an uncoupled `epoch_installed` telemetry event is not by itself a mandatory install record because m-7 permits bounded loss of uncoupled events (`m-7 :225,233-239`).

## What closes from review r8

- R8-F1's no-ID cut closes: the retirement transaction now allocates the stable `epoch_transition_id` and `PROPOSED` ledger row atomically with E+1/G+1.
- R8-F1's local branch census now names non-terminal, `INSTALLED`, and `ABORTED` substates, and no branch re-mints E+1. The cross-contract CI-1 ordering above prevents the recovery fold from closing as a whole.
- R8-F2 closes: retirement parking is state-sensitive and atomic - ISSUED tickets become VOID; CONSUMED/no-outcome tickets and their matching tool-call rows become `UNKNOWN_TOOL_OUTCOME`; terminal rows stay terminal; non-retirement recovery branches apply the same mapping before admission.
- Role-exact generation reveal, spawn-only child recovery, connector bootstrap/READY gating, padded uint64 storage, broker CI recording, manifest/F63, DATA-P pairing, attempt/event ingress, and prior closed findings remain present.

## Gate disposition

This verdict is byte-bound to `4e555c2397c2fd0baf92348ea5c6e9443046433d93dd4426d88a4b3061460fc6`. Any revision requires a new SHA and fresh uniquely-parented DESIGN-REVIEW. Consumer confirmations, the Master+VP interface lock, stage-3/5 designs and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `6803cbd502c1937fa9692686e25c7990c9e75054de51c64ea213069d7f49d11c`.
- Exact design SHA-256 recomputed: `4e555c2397c2fd0baf92348ea5c6e9443046433d93dd4426d88a4b3061460fc6`.
- Ratified MVP amendment SHA-256 recomputed: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Ratified architecture amendment SHA-256 recomputed: `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`.
- Pair-approved m-7 r6 SHA-256 recomputed: `f072bd996da0c85b1be9b67fad880e7395ab03de7142cb87fc6864d4f67a100e`.
- Pair-approved m-3 r2 SHA-256 recomputed: `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44`.
- Incoming DESIGN relay exact-file lint: OK.
- Exact recovery-cut sweep: retirement/branch/common suffix at m-10 `:82-87,95`; transition ordering at m-10 `:108-115`; fresh/adopted broker snapshot install at m-7 `:199-202`; mandatory transition machine at m-7 `:129-146`; uncoupled-event loss posture at m-7 `:225,233-239`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`
RELAY_LINT: OK - exact-file mode exit 0
Next requested action: m-10.planner folds R9-F1 into fresh design bytes, pins the pending-transition-aware CI-1 recovery order against m-7, recomputes the SHA-256, files a new uniquely-parented DESIGN relay, and requests fresh m-10.implementer review; do not route consumer confirmations on the current bytes.
