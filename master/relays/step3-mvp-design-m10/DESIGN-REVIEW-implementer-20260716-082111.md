## DESIGN-REVIEW - m-10 stage-1 IPC/manifest seam r6 must revise: stale sole-reveal absolutes and non-idempotent recovery retirement

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m10-review-r7
PARENT_DISPATCH_ID: step3-mvp-design-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - both findings are bounded proof/state-machine corrections under the already-chosen generation-blind and spawn-only recovery branches; Master+VP retain interface-lock authority
GRILL_REQUIRED: no - this re-review does not reopen the operator-locked topology, threat ceiling, F59, F60, or revocation-first choice; the m-10 grill remains stage 5
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m10/DESIGN-planner-20260716-101500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260716-082111.md
SUBJECT: must-revise exact 1b950409... - generation-blind CTRL-C and spawn-only fail-closed recovery are the right branches, but the proof still denies the broker's pre-lease knowledge and recovery unconditionally re-retires states that must resume without another epoch mint

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner - I reviewed the exact revised bytes at SHA-256 `1b950409d5f5b4143df40203d013674b41ebbc0cae0691997d854fe42555e91a`. Routing, lineage, `DESIGN_DOC_ID`, incoming exact-file lint, the r7 basis hash, and both consumed sibling hashes pass. Both selected repair directions are correct. Two exact contradictions remain in their proof/state-machine wording.

## Findings

### R7-F1 - The broker-feed exception is added, but the old sole-frame/sole-process absolutes remain

The new distribution clause correctly makes m-8's `epoch_update{run_id,turn_epoch}` generation-blind and explicitly names the pre-lease reveal set as the broker feed plus post-lease worker `assign` (`2026-07-16-mvp-ipc-manifest-seam-contract.md:97`). The broker feed is a framed control channel to the separate broker process (`:21,101`; m-7 r6 `:98-101`).

Two governing sentences still state the opposite:

- worker `assign` is “the ONLY frame that reveals `generation_id`” (`:60`); and
- G+1 “is revealed to exactly one process, via `assign`, sent ONLY after step 5” (`:91`).

The broker receives G+1 before lease bind as the verifier, so both unqualified absolutes are false on the same bytes. Verifier-versus-presenter is a sound safety distinction, but it does not make a framed message cease being a frame or the broker cease being a process.

Required return: rewrite the two stale absolutes at `:60` and `:91` in role-exact terms. The safety fact is that the broker receives G+1 only as the installed-tuple verifier and exposes no presenter path, while post-lease worker `assign` is the only worker/presenter-facing disclosure; retired/candidate workers and m-8 do not learn G+1 pre-lease. Make that statement agree byte-for-byte with the exact reveal set at `:97` and preserve generation-blind CTRL-C.

### R7-F2 - App-main recovery is not conditioned on durable lifecycle state, so repeated/pre-lease crashes can mint extra epochs

The replacement sequence unconditionally says every app-main restart treats the prior pair as unproven-terminal and applies the retirement transaction, minting E+1 (`:82`). But the epoch rule says a pre-lease candidate wash-out never re-mints and each retirement mints exactly once (`:88`). The initial path also durably creates E=1 before connector bootstrap/lease bind (`:84`).

Those clauses conflict in at least two mandatory crash cuts:

1. crash after initial E=1 freeze but before the first lease: there is no active leased generation to retire, so recovery must resume/wash out under E=1, not mint E=2;
2. crash after a retirement transaction has already committed E+1/G+1 but before connector bootstrap or lease bind: the next app-main must resume the committed transition or wash out the pre-lease candidate under the same E+1, not apply a second retirement and mint E+2.

The words “exactly once” do not identify the durable predicate that distinguishes those states, and the sequence also parks rows before invoking a canonical retirement transaction that itself is defined to park them atomically (`:82,90`). A T4 implementer therefore has two incompatible transition orders and no pinned retry/resume test.

Required return: replace the unconditional recovery arrow with a durable-state recovery matrix and conditional transition rule. At minimum: (a) active LEASED generation with no committed retirement ⇒ one atomic §B.4 retirement transaction, including row parking; (b) committed `RETIRED_PENDING_REAP` + current pre-allocated successor/non-terminal epoch transition ⇒ resume/reconcile that same epoch/transition, never mint again; (c) ALLOCATED/SPAWNING/READY pre-lease candidate ⇒ mark failed/wash out and allocate the next generation under the current epoch; (d) initial E=1 with no prior lease ⇒ resume the initial step-4/5 path under E=1. Pin the transaction guard/identity that makes a crash immediately before versus after the retirement commit converge on one mint.

## What closes from review r6

- R6-F1's actual CTRL-C leak closes: `epoch_update` is generation-blind, m-8 needs only the epoch value, and the generation-bearing broker feed remains separate. R7-F1 is the incomplete proof-text sweep, not a request to restore G on CTRL-C.
- R6-F2's topology choice closes: m-8/m-9 are never adopted; CTRL EOF makes both children fail closed and exit; recovery creates fresh CTRL-W/CTRL-C/DATA-P; the broker alone has CI-1 adoption; connector READY gates lease/admission.
- The orphan residual is bounded honestly at the locked confusion-not-malice/same-user ceiling, and initial freeze-before-bootstrap ordering is now named.
- R4-F2 and the prior broker CI, manifest/F63, DATA-P pairing, PREPARING-ledger, attempt/event ingress, canonical wire-counter, and crash-outcome repairs remain present.

## Gate disposition

This verdict is byte-bound to `1b950409d5f5b4143df40203d013674b41ebbc0cae0691997d854fe42555e91a`. Any revision requires a new SHA and fresh uniquely-parented DESIGN-REVIEW. Consumer confirmations, the Master+VP interface lock, stage-3/5 designs and grill, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Verification

- Exact design SHA-256 recomputed: `1b950409d5f5b4143df40203d013674b41ebbc0cae0691997d854fe42555e91a`.
- Ratified amendment SHA-256 recomputed: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Pair-approved m-7 r6 SHA-256 recomputed: `f072bd996da0c85b1be9b67fad880e7395ab03de7142cb87fc6864d4f67a100e`.
- Pair-approved m-3 r2 SHA-256 recomputed: `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44`.
- Incoming DESIGN relay exact-file lint: OK.
- Exact contradiction sweep: sole frame at `:60`; sole process at `:91`; broker-feed reveal set at `:97`; unconditional recovery retirement at `:82`; initial E=1 at `:84`; no-remint candidate rule at `:88`; atomic park+retire at `:90`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`
RELAY_LINT: OK - exact-file mode exit 0
Next requested action: m-10.planner folds R7-F1 and R7-F2 into fresh design bytes, recomputes the SHA-256, files a new uniquely-parented DESIGN relay, and requests fresh m-10.implementer review; do not route consumer confirmations on the current bytes.
