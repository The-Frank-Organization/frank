## DESIGN-REVIEW — MUST-REVISE m-10 stage-5 r1 exact bytes: the H-17 final-review precondition is absent, and three realization claims conflict with the frozen counterparties

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage5-m10-review-r1
PARENT_DISPATCH_ID: step3-mvp-stage5-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — all findings are bounded owner corrections against already-ratified decisions and frozen contracts; any desired contract change must route up instead
GRILL_REQUIRED: no — G-1 through G-5 remain operator-disposed; this review does not reopen those choices
DESIGN_DOC_ID: step3-mvp-design-m10-control-plane
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-stage5-m10/DESIGN-planner-20260720-054000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-8.planner, m-7.planner, m-3.planner, m-2.planner, m-1.planner
BUNDLE_ID: m-10-app-control-plane
OWNER: m-10 (App Control Plane / Supervisor)
RELAY_PATH: master/relays/step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-054100.md
SUBJECT: MUST-REVISE exact stage-5 r1 b04201b6 — the superseding H-17 schema/coverage gate was not folded; attempt-linked E0 is not universally present under the m-9 reporter contract; wake rediscovery is assigned to the no-seat scheduler instead of m-9; connector-never-ready has no total path into the 10-try terminal

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the exact stage-5 r1 bytes at SHA-256 `b04201b6e4f5e2954b4d6d60164a3f9bc470951ef285ca60859fb835c17db11f`. Routing, `DESIGN_DOC_ID`, the stage-5 dispatch, the operator's five grill dispositions, and incoming exact-file lint pass. The final-byte approval precondition does not: four blockers remain.

## Findings

### M10-S5-R1-F1 — the superseding H-17 census requirement is not folded

The design's §11a still cites the superseded `043341` supplement and carries the original seven-row/eight-column table. Master's operative `050327` corrective requires the canonical `master/H17-CENSUS-SCHEMA.md` v1 at SHA-256 `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`, verbatim. The current table lacks required fields including stable `effect_id`/`effect_class`, requester/executor, authority source, policy owner/artifact, decision point, separate authorization/effect linearization, separate reporter/observer/validator/record, and threat-claim scope.

Coverage is also incomplete exactly as `050327` warned: no per-effect row or explicit non-effect rationale exists for run start/stop/recovery, process spawn/retire, turn admission, epoch publication, cancellation/control sends, provider-attempt transitions, or app-event carriage. The generic STORE meta-row cannot substitute for those authority/failure rows.

This is a gate violation, not presentation polish: `050327` says no GRILL_LOCK, final-byte review, or final SITREP until the complete inventory is folded. The operator's grill dispositions remain durable, but §15 cannot claim the final-review exit is open on these bytes.

Required revision: replace §11a with the canonical v1 field set; map every authoritative transition/effect family to a globally unique row or an explicit non-effect rationale; use `unknown`, `not specified`, or `residual` rather than inventing missing facts; and split reporter, independent observer, validator, and recorder honestly. Cite `050327` as operative and `043341` as superseded.

### M10-S5-R1-F2 — §11 overclaims E0 visibility and silently moves the event producer

§11 says `pending_app_events` rows are written by the applier at their owning transitions and that **any** attempt-linked terminal failure leaves an E0 row. The frozen counterparty contracts do not support that universal.

- r36 §B.1 says m-9 delivers the m-3-schema'd `app_event{…}` frame and m-10 stores it.
- m-3 r4 fixes `reported_by` to the m-9 worker and makes m-9 the E0 populator/carrier.
- m-8 r12's exact outcome table says connector/DATA-P loss has an E0 event only if the live worker emits it before retirement fencing; if retirement wins, no terminal E0 survives. A worker crash after an attempt produces **no E0 event**, because the populator is dead and the replacement cannot reconstruct m-10's private rows.

The canonical m-10 `provider_attempts`/UNKNOWN row survives those cuts, but it is not thereby an `m3.app_event.v1` E0 event. Having the m-10 applier synthesize a worker-attributed event would be a producer/provenance amendment, not a realization detail.

Required revision: preserve the frozen flow — m-9 authors/delivers the event, m-10 validates and durably persists the received frame — and narrow the terminal-visibility statement. Add the attempt-linked no-E0 residuals for worker crash and retirement-winning connector loss; distinguish the canonical m-10 row from the E0 copy. If m-10 synthesis is desired instead, route an explicit m-3/m-9/m-10 owner amendment rather than folding it here.

### M10-S5-R1-F3 — G-4 gives conductor rediscovery to a component with no conductor edge

§6 correctly begins with `wake_forward` from the worker. §12 then says the m-10 scheduler is the reader that checks the durable conductor inbox and “reads the unread relay.” Frozen r36 §E is explicit: only the m-9 seat receives push and performs durable `project`/`read` rediscovery at startup/reconnect/poll; m-10 has no seat credential or conductor verb and only consumes `wake_forward{relay_id}`. The current G-4 prose therefore admits an implementation that violates the NOT-a-seat boundary.

Required revision: pin the physical chain and its two durable facts:

1. m-9 reads/polls the conductor record and forwards `wake_forward`;
2. m-10 idempotently inserts its own `wake_schedule` row;
3. the m-10 scheduler atomically flips `pending → dispatched` with turn admission.

The at-most-once proof applies to step 3 in m-10's store. `dispatched` means processing began; it is not a conductor read receipt. Keep the operator's email-semantics decision, but make “checking mail” and dropped-push rediscovery m-9's obligation, never an m-10 conductor read.

### M10-S5-R1-F4 — connector-never-ready is named terminal without a reachable retry/cap transition

The start order is connector-first; worker allocation follows only after `connector_ready`. G-2's durable counter is expressly a `consecutive-failed-generation` counter, but §5 gives a pre-worker connector fault only the co-restart statement. §10/§12 nevertheless list connector-never-ready as a run-terminal loud failure. On the written machine, a connector can fail before any worker generation exists, so no named counter necessarily increments and the advertised 10-try FAILED terminal is not reachable.

The fixture list also does not support the incoming claims: §14 names only `crash-loop bound → run FAILED`; it omits backoff persistence across app restart, reset after a completed turn, cancel during backoff, the exact ninth/10th boundary, connector-never-ready reaching the cap, and persistent-alert/non-zero-exit checks for the loud-failure classes.

Required revision: make the supervised-pair attempt unit and counter transitions total across worker spawn/handshake/health failures and connector spawn/handshake/ready failures, or define a separate bounded connector policy consistent with G-2. State precisely what increments, what resets, how connector-first failure co-restarts when no worker exists yet, and which commit creates run FAILED. Add the missing backoff, connector-cap, and loud-surface fixtures. If this cannot be derived without changing the operator's G-2 disposition, route that narrow question rather than choosing silently.

## Accepted basis

The following surfaces are accepted and need not be redesigned while folding the findings:

- module-in-app-main topology with real CTRL-W/CTRL-C/CI-1 seams and no DATA-P endpoint retained by m-10;
- one serialized applier as sole writer, commit-before-visible replies, committed-snapshot reads, and no shared conductor store/code;
- frozen r36 precedence, worker/connector pairing, lease/epoch linearization, F59 replay/consume/record realization, and the immutable manifest/serve gate;
- opaque `credential_ref` handling, m-8-exclusive credential bytes, and m-10's no-seat/no-conductor-verb boundary;
- operator dispositions G-1, G-2, G-2a, G-3, G-4, and the G-5 route-up, subject only to the realization corrections above;
- honest ambient-authority, same-UID file, and build-identity residuals; prior-art provenance without vendoring.

## Scope and remaining gates

Do not file the stage-5 closure SITREP or route consumer confirmations on `b04201b6…`. Fold F1–F4 in the unlocked stage-5 design only, preserve the frozen r36 and sibling contracts, and return fresh uniquely-parented bytes.

The complete H-17 census, fresh m-10 pair approval, consumer confirmations, Master+VP stage-6 interface-lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held. Any byte change requires a fresh exact-byte m-10.implementer DESIGN-REVIEW.

## Verification

- Exact incoming DESIGN relay SHA-256 recomputed: `3771291cf00f9dabd6dfa9c4dc80d7a9137ce05c78500640f791f187911d370d`.
- Exact stage-5 r1 design SHA-256 recomputed: `b04201b6e4f5e2954b4d6d60164a3f9bc470951ef285ca60859fb835c17db11f`.
- Canonical H-17 schema SHA-256 recomputed: `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`.
- Frozen m-10 r36 SHA-256 recomputed: `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`.
- Frozen m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`.
- Frozen m-9 lifecycle r19 SHA-256 recomputed: `2a96a07bb2f2606b7b42fe34270beaa209ca08cfd2c19f6e91f44eb18eef734c`.
- Incoming DESIGN exact-file lint: OK.
- Targeted sweep: stage-5 dispatch + corrected H-17 supplement; canonical census schema; r1 §§1–15; frozen r36 lifecycle/wake/E0 boundaries; m-8 r12 loss/E0 cuts; m-3 r4 event producer/schema; grill route-ups and accepted residuals.

ACTIONS_GIT_REF: docs-workspace disk action — created this review relay and appended one `master/relays/INDEX.md` row; no design doc, frozen contract, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `6e4d657`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-054100.md`.
Next requested action: m-10.planner folds M10-S5-R1-F1..F4 on fresh stage-5 bytes and returns a uniquely-parented DESIGN relay; all later gates wait.
