## RECONCILE — APPROVE the m-8 r12 letter rebase to m-10 r40: admission_ref and its pre-commit sizing refusal remain upstream of every m-8 connector and provider-attempt surface

ROLE: Implementer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8-basis-review-r40
PARENT_DISPATCH_ID: step3-mvp-design-m8
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this is the bounded F73 basis-rebind review; Master+VP retain stage-6 interface-lock authority
GRILL_REQUIRED: no — no design choice, topology, policy, or owner byte is reopened
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-mvp-design-m8/RECONCILE-orchestrator-planner-20260720-200339.md
FROM: m-8.implementer
TO: master.orchestrator-planner
CC: m-8.planner, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
BUNDLE_ID: m-8-provider-adapters
OWNER: m-8 (Provider Adapters / connector)
RELAY_PATH: master/relays/step3-mvp-design-m8/RECONCILE-implementer-20260720-200710.md
SUBJECT: APPROVE unchanged m-8 r12 4b670a79... against m-10 r40 d2ce9831... — turn_open.admission_ref, complete-frame sizing, and task_input_frame_overflow refusal create no m-8-visible frame or row transition

VERDICT: approve

Master — I approve the bounded basis rebind.

The two frozen owner hashes are:

- m-8 provider contract r12: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`;
- m-10 app-control-plane contract r40: `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`.

## Delta disposition

**APPROVE — disjoint.** The r36→r40 semantic delta is confined to m-10's turn-input/admission boundary:

- `turn_open` gains the required `admission_ref` member with the closed wake-relay/operator-input union;
- the complete canonical `turn_open` frame is constructed or exactly sized before the admission transaction;
- an over-limit operator task returns the operator-boundary-only `admission_refused{reason: task_input_frame_overflow}` before any admission transaction;
- accepted input commits `turns.admission_ref` with the turn/lease/wake disposition, then emits `turn_open` post-commit from that committed state;
- the post-commit/pre-send crash cut re-emits the same ref byte-identically.

These are CTRL-W turn-admission mechanics between the app supervisor and m-9, plus an operator-surface refusal. M-8 receives neither `turn_open` nor `admission_refused`, stores neither `admission_ref`, and has zero matches for either term in its frozen r12 contract.

## Named sizing/refusal check

The r40 sizing gate does not interpose on m-8's connector bootstrap or provider-attempt path:

1. **Connector order:** m-10 §B.1/§B.4 still requires connector `hello → connector_assign → connector_ready` before worker admission. The sizing gate is evaluated for a candidate `turn_open`; it does not alter, repeat, suppress, or add a field to `connector_assign`, and it sends no CTRL-C frame.
2. **Overflow path:** `admission_refused{reason: task_input_frame_overflow}` is emitted at the operator command boundary before any admission transaction. Exact r40 states zero durable effects: no `turns` row, no active-turn lease, and no admission act. The gate neither mints an epoch nor changes worker/connector lifecycle state.
3. **Attempt reachability:** without a committed/admitted turn, no lawful m-9 path reaches `attempt_open` or DATA-P. The existing defense against a confused request is unchanged: `attempt_open` naming an invalid turn receives `attempt_open_reject{invalid_turn}` with no `provider_attempts` row, no DATA-P request, and no attempt-budget charge. That pre-existing peer-fault path is not an effect of the operator refusal.
4. **Existing rows:** the refusal transaction is deliberately absent, so it cannot terminalize, park, cancel, or otherwise mutate an existing `provider_attempts` or `cancellations` row. The sole r40 schema addition is `turns.admission_ref`; the `provider_attempts` row is byte-carried.
5. **FRAME_MAX boundary:** §A.2's compiled 4 MiB maximum and the existing DATA-P chunking obligation are unchanged. R40 only applies that existing cap early enough to make required `turn_open` encodable; it adds no DATA-P size rule and no connector payload.

Thus the refusal's zero-durable-side-effect claim holds at every m-8-visible state boundary. A connector may already be READY when an operator input is refused, but the refusal neither changes that committed readiness nor creates a provider attempt.

## Independent six-locus check at exact r40

1. **Connector bootstrap:** r40 §B.1 still carries the exact seven-field `connector_assign{run_id, turn_epoch, run_manifest_digest, policy_digest, provider_lane_id, lane_catalog_digest, credential_ref}`. The new `admission_ref` is only on worker-facing `turn_open`; it is absent from CTRL-C.
2. **Attempt-open handshake:** r40 §B.1 still carries `attempt_open_ok{attempt_id, parked_unknown:[…]}` and `attempt_open_reject{attempt_id, reason ∈ {stale_epoch, invalid_turn, invalid_lease}}`, with commit-before-ack, no-row rejection, no DATA-P, no attempt charge on reject, and the separate provider-attempt ceiling intact.
3. **Attempt-result vocabulary:** r40 §B.1 still consumes the closed m-8 set `{sent_completed, denied(…), transport_failed, unknown, rejected_local(…), cancelled(pre_transport|post_invocation)}`. Neither `admission_refused` nor `task_input_frame_overflow` enters `attempt_result`.
4. **Terminal attempt rows:** r40 §B.1/§F still preserve terminal `provider_attempts.REJECTED_LOCAL` and terminal `provider_attempts.CANCELLED`; the r40 amendment adds only `turns.admission_ref` and does not modify `provider_attempts`.
5. **Cancellation consumer discipline:** r40 §B.1/§F still make `cancelled(<cancel_point>)` one-way, preserve raw closure/crash as `UNKNOWN_PROVIDER_OUTCOME`, define duplicate equivalence only over `{attempt_id, reported turn_epoch, cancel_point}`, and keep `cancellation_id` provenance/correlation-only. No admission-refusal branch reaches this table.
6. **Epoch authority:** r40 §B.4 still makes m-8 generation-blind, advances its authority only from m-10 `epoch_update`, maps below-current to `STALE_EPOCH`, and maps above-current to held/retriable `EPOCH_AHEAD` plus CTRL-C query. The sizing refusal performs no admission transaction and mints no epoch.

## Gate disposition

The r12 owner-byte approval `DESIGN-REVIEW-implementer-20260718-043932.md` remains valid without replay. No m-8 byte moves and no new m-8 hash is created. This review satisfies the final m-8 citation leg of the F73 rebind round at exact m-10 r40.

Approval advances only to Master's stage-6 lock packet after the remaining exact-r40 legs. It does not itself grant interface lock, PLAN, T4 token, implementation, credentials, provider calls, release binding, E3, merge, or deploy authority.

## Verification

- Incoming direct-address relay SHA-256 recomputed: `6576247d2ed40481e194c983f4e61b546f0f5790dc980ac40ffd8bcf6d7dccde`.
- Frozen m-8 r12 SHA-256 recomputed: `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`.
- Frozen m-10 r40 SHA-256 recomputed: `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`.
- M-10 r40 DESIGN relay SHA-256 recomputed: `3ed94c61a010733a0db659b9c89d52e2a72140b2b1e57aab11fa38142fd68e1b`.
- M-10 r40 approval relay SHA-256 recomputed: `d9dcf9d2c88e43a27f0fde46ad9f92c25b144081c39c559b74ad30f1ba32aa5c`.
- M-10 r40 SITREP relay SHA-256 recomputed: `757e96718f8ff19b5fd3156bab7b7a434c5e76dd4d860424f42adbd93d50877f`.
- Exact incoming, r40 DESIGN, r40 DESIGN-REVIEW, and r40 SITREP relays lint clean.
- `git -C frank status --short` returned empty; HEAD = `6e4d657913229027fc94a1e2a8c2348b05c09a75`.

ACTIONS_GIT_REF: docs-workspace disk action only — created this bounded review relay and appended one `master/relays/INDEX.md` row; no owner design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repository; `git -C frank status --short` returned none — clean at `6e4d657`.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-mvp-design-m8/RECONCILE-implementer-20260720-200710.md`.
Next requested action: Master binds this exact-r40 m-8 approval into the stage-6 lock packet after the remaining citation legs; keep frozen owner bytes and every later authority gate held.
