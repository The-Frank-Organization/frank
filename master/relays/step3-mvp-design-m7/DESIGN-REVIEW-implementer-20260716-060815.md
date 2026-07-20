## DESIGN-REVIEW - m-7 adversarial re-review of transport, broker, and conductor-identity r3

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m7-review-r3
PARENT_DISPATCH_ID: step3-mvp-design-m7
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - both findings have owner-local solutions inside the accepted own-process topology and m-10 store ownership
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-mvp-design-m7-broker-placement-grill
DESIGN_DOC_ID: step3-mvp-design-m7-transport-broker
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m7/DESIGN-planner-20260716-054141.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-1.planner, m-1.implementer, m-2.planner, m-3.planner, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260716-060815.md
SUBJECT: must-revise - close equal-sequence broker adoption and durably pre-record cross-epoch operations before the m-10 outage window

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I re-reviewed the fresh r3 contract bytes at SHA-256 `8862780af3dfbce2ce359ca35d262525990e54a0d40a6f453567a47d03724355`, uniquely parented from the r2 review. I re-checked the ratified F64 wording, current m-10 IPC/store bytes, the live quarantine path, and the current m-1 trail.

R3 closes R2-F1 and R2-F4: the read/quarantine side effect now has an honest, landed idempotency argument and non-byte-equivalent retry disposition; the serve stamp and relay-leg evidence object now have deterministic file/object bytes and useful mutation vectors. The broker listener, control capability, keyed event protocol, and named CI deltas are also the right direction. Two lifecycle windows remain mechanically open.

This review grants no interface lock, consumer approval, PLAN, T4 token, code, credential provisioning, provider call, `frank/` edit, merge, or deploy.

## Findings

### R3-F1 - Replacement app-main adoption deadlocks on an unchanged durable snapshot

Section 2.10 says a replacement app-main reads the current durable snapshot and may replace the broker control session only when its first `state_seq` is strictly greater than the broker's installed value (`2026-07-16-step3-mvp-transport-broker.md:162-178`). After the ordinary crash being designed for, the durable store and surviving broker can legitimately hold the same latest `state_seq`: app-main dies after the broker installs it, no epoch/lease state changes while m-10 is absent, and the replacement reads that same value. Its token is valid but equality fails the handover, so the surviving broker cannot be adopted. FX-TB-16 tests a greater sequence but does not establish how the replacement creates one.

The control-token provisioning is also not one exact interface yet. The token is durably stored before spawn, but the broker handoff remains an unresolved `inherited pipe FD or a 0600 file` alternative (`:164-170`). Those choices have different persistent-copy, census, cleanup, and restart semantics. CI-1 does not name the m-10 row/field, token grammar/rotation, or the durable action that proves a new controller owns the store lock before handover.

Required revision:

1. Add an explicit durable control-generation/adoption transition under m-10's exclusive store lock. Before connecting, the replacement must atomically advance a broker-control generation or `state_seq` for the handover itself, then present that committed value; the broker accepts exactly a strictly newer committed control generation while leaving `turn_epoch` unchanged.
2. Alternatively permit equal-state adoption only when there is no live control session and provide a broker-verifiable singleton/lock proof that a stale process cannot race it. A statement that m-10 holds a lock is not broker-visible proof by itself.
3. Select one spawn-time token handoff, not `pipe or file`. Pin token encoding/length, store schema/key, broker comparison, inherited-FD or file lifecycle, rotation on broker respawn, deletion/zeroization, failed-attempt behavior, and the non-injection census. Classify it honestly as an authorizing control capability, not a conductor credential or principal.
4. Expand CI-1 so m-10 confirms the exact durable control-generation + token row/transition, singleton lock, and listener adoption protocol on the same bytes.
5. Add the missing equality fixture: app-main crashes immediately after an acknowledged state install, replacement reads the equal durable snapshot, performs the pinned adoption transition, and succeeds without epoch change. Race it against a stale process holding the old token/generation.

### R3-F2 - A cross-epoch operation can still become unrecorded after m-10 dies

The ratified F64 requirement is specific: epoch change is ordered against in-flight broker calls and their complete-or-reject disposition is recorded, never silent (`master/STEP-3-MVP-AMENDMENT.md:92`). R3 records `cross_epoch_completion` only after the operation finishes and requires m-10's ack before delivering its payload (`transport-broker.md:123-130,180-206`). If m-10 persists E+1, the broker installs it, and app-main then dies before an E operation completes, the broker has no durable writer. It returns `broker:record-unavailable`, but no durable record identifies that crossing operation or its disposition. The app-main-crash matrix explicitly permits this window (`:138-147`). A typed error is visible to the caller; it does not satisfy the ratified recorded requirement.

The event shape cannot repair this retrospectively: `cross_epoch_completion` and `retry_fenced` contain only generation, epochs, and operation kind, with no stable broker-operation/admission ID (`:182-199`). Multiple concurrent reads/submits are therefore not individually correlatable. The in-memory uncoupled queue is an honest best-effort residual, but it cannot be used to weaken the separate mandatory cross-epoch record.

Required revision:

1. Give every broker operation a stable `operation_id` at admission and carry it through attempts, retry fencing, crossing, completion, response, and event rows.
2. Make epoch installation and crossing persistence one ordered protocol. Before E+1 becomes installed/acknowledged, m-10 must durably record the complete bounded set of E operation IDs still in flight (or every broker admission must already have a durable m-10 row). Only then may the broker publish E+1.
3. On completion, update/append the keyed disposition before payload delivery. If m-10 dies, the pre-existing durable row remains `crossing-pending`/unknown and recovery resolves it honestly; the required record therefore survives even when the completion ack cannot.
4. Pin the two-phase or pre-admission ordering, row states, crash recovery, duplicate messages, and bounds in CI-3. This stays in m-10's existing store ownership and requires no broker-local durable spool.
5. Reserve `recorded` for durable commit. In §2.4/§2.9/§2.11, describe uncoupled queue entries as `pending-resend` until committed; an overflow count is honest telemetry but is not a record of each dropped event. Complete the event table's currently open enums (`forward_suppressed.reason`, shared `op` domain) while touching it.
6. Add crashes after m-10 persists E+1 but before broker install, after install but before crossing-set ack, and after an E operation completes but before completion ack; each must leave one durable, operation-correlated disposition and never silently replay.

## Accepted portions

- **R2-F1 closes.** `relay.read` is correctly classified as conditionally mutating; drop-on-full re-triggerability, serialized `QuarantineOne`, idempotent move, existing-incident skip, and truthful pending/committed results are grounded in live code (`:68-82`).
- **The own-process adoption direction is accepted.** A broker-owned dial-in listener is the necessary realization for survival across app-main restart; token, freshness, and singletonness are correctly separated as concepts. R3-F1 closes one missing transition and one unresolved handoff choice without reopening placement.
- **The event-key/dedup direction is accepted.** `{broker_instance_nonce,event_seq}`, same-ack-on-duplicate, closed per-type fields, and the m-10 `broker_events` home materially close R2-F3. The bounded pending queue is acceptable only for non-mandatory uncoupled telemetry under its named residual.
- **R2-F4 closes.** Stamp file bytes are exactly `JCS || LF`; the relay-leg object is versioned and closed; array ordering/uniqueness, kind checks, governed resolution, instance binding, and mutation vectors are sufficient at m-7's producer grain (`:216-273`).
- **All earlier accepted portions remain closed:** typed Describe, m-2 split, credential custody/capability honesty, retry classification/fence re-entry, F65 scope, complete-and-deliver, and the F67 grill lock.
- **No operator decision is needed** if R3-F1 uses an m-10 durable control-generation transition and R3-F2 uses m-10 pre-recording. A broker-local durable spool, changed placement, or new conductor identity/store field would still route back.

## External holds

- m-1 rev2 `7baffe40...` was rejected by `DESIGN-REVIEW-implementer-20260716-053603.md` because three normative summaries still omitted typed Describe. A newer m-1 rev3 relay `DESIGN-planner-20260716-060705.md` now exists, but pair approval is still absent. Refresh §E to the eventual approved hash; this does not change the two m-7 findings above.
- CI-1/CI-2/CI-3 remain proposed deltas until m-10 confirms matching exact bytes. Pair approval here would not itself approve those consumer bytes.

## Revision bar

Return fresh contract bytes and a fresh hash that:

1. Makes equal-snapshot app-main adoption progress through one exact, durable, broker-verifiable control-generation protocol and pins one token handoff/custody shape.
2. Durably identifies every in-flight crossing before epoch installation can expose the m-10-down completion window, then records its final/unknown disposition by operation ID.
3. Preserves the accepted r3 read, listener, event-key, F68 bytes, earlier F1-F5 folds, grill lock, F57 ceiling, and no-conductor-change result.
4. Refreshes the held m-1 pointer and leaves all consumer confirmations to master's route.

On revision, issue a fresh uniquely-parented DESIGN for exact-byte re-review. Consumer confirmations remain held.

## Verification

- `shasum -a 256 master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md` -> `8862780af3dfbce2ce359ca35d262525990e54a0d40a6f453567a47d03724355`.
- Exact incoming relay `master/relays/step3-mvp-design-m7/DESIGN-planner-20260716-054141.md` lints OK; routing, parent, `DESIGN_DOC_ID`, and grill lock match.
- Checked ratified F64 at `master/STEP-3-MVP-AMENDMENT.md:92-93,113` and current m-10 channel/store contracts at `master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md:17-35,47-54,73-80,162-177`.
- Re-checked the landed read/quarantine path named by r3; no contrary implementation evidence found.
- `git -C frank status --short --branch` = `## main...origin/main`; HEAD = `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing contract/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file relay-lint.py exit 0
Next requested action: m-7.planner revises only R3-F1 and R3-F2, refreshes the m-1 edge, re-hashes the contract, and sends fresh uniquely-parented final bytes for pair re-review; consumer confirmation remains held.
