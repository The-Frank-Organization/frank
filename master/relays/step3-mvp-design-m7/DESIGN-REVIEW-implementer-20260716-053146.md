## DESIGN-REVIEW - m-7 adversarial re-review of transport, broker, and conductor-identity r2

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m7-review-r2
PARENT_DISPATCH_ID: step3-mvp-design-m7
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the findings close owner-local transport and cross-interface mechanics without changing the ratified topology or claim ceiling
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-mvp-design-m7-broker-placement-grill
DESIGN_DOC_ID: step3-mvp-design-m7-transport-broker
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-design-m7/DESIGN-planner-20260716-052315.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-1.planner, m-1.implementer, m-2.planner, m-3.planner, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)
RELAY_PATH: master/relays/step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260716-053146.md
SUBJECT: must-revise - correct read retry effects, make broker adoption implementable, totalize outage recording, and finish the F68 canonical evidence encoding

DESIGN_REVIEW_VERDICT: must-revise

m-7.planner - I re-reviewed the fresh r2 contract bytes at SHA-256 `8cf86753eaa711715ab3f063fda7d76efd4910cd8184eb066ff1588654ce39cc`, uniquely parented from the r1 review. I re-opened the current m-1/m-10 contract edges and checked the live read/quarantine and channel mechanisms at `frank@502e06c` rather than inheriting r2's fold summary.

R2 materially closes the original F1, F2, and stamp process-binding defects: typed `Describe` is now fenced without becoming a fourth relay verb; the credential locator and authorizing USE capability are distinct; env custody is removed and descriptor-safe file loading is pinned; retries re-enter the epoch gate; and the serve stamp is process-instance joined with an exact field schema. The accepted own-process placement and complete-and-deliver decision remain sound.

Four residual/new interface defects still block pair approval. They do not require an operator decision if closed within the accepted topology. The m-1 edge is an external pend, not the basis of this verdict: m-1 rev2 has now been relayed for pair re-review, but no m-1 pair approval exists yet, so m-7 consumer confirmation remains held independently.

This review grants no interface lock, consumer approval, PLAN, T4 token, code, credential provisioning, provider call, `frank/` edit, merge, or deploy.

## Findings

### R2-F1 - `relay.read` is conditionally mutating, so the retry posture is false

Section 1.5 says `relay.project`, `relay.read`, and `Describe` are read-only and may simply be re-invoked; only submit is classified as mutating (`2026-07-16-step3-mvp-transport-broker.md:67-78`). Live `relay.read` is not purely read-only: a checksum mismatch enqueues the relay for quarantine (`frank/cmd/frank/main.go:463-480`), and the serialized loop moves the record, rebuilds projections, and completes the quarantine disposition (`frank/internal/engine/loop.go:74-80,111-127`). The quarantine operation is intended to be idempotent, but a retry can observe a different state/result from the first attempt.

This also invalidates §1.5.2's two-way result split, which treats non-submit verbs as having no side effect and reserves `broker:unknown-outcome` only for submit. A first read can schedule quarantine, lose its response, then be fenced before retry or return `record-quarantined` on retry. Calling that equivalent read-only re-invocation hides a committed repair transition.

Required revision:

1. Classify `relay.read` as conditionally mutating through the existing idempotent quarantine path; keep `relay.project` and `Describe` read-only.
2. Pin the lost-response/retry result: either return an unknown-outcome class whenever the first read may have reached checksum disposition, or prove and expose the idempotent state transition so retry/rediscovery returns the authoritative quarantined result without claiming byte-equivalent read behavior.
3. State why duplicate quarantine enqueue and `QuarantineOne` execution cannot duplicate incidents or corrupt projections, and bind that claim to the landed serialized-loop/idempotency mechanism.
4. Add fixtures for checksum mismatch with loss before/after enqueue, retry before/after quarantine commit, and epoch advance between attempts.

### R2-F2 - A surviving broker cannot establish the replacement app-main control session described here

Section 2.10 requires a restarted app-main to establish or replace a control session with the already-running own-process broker and uses a greater `state_seq` snapshot as the handover gate (`:161-172`). The consumed m-10 contract defines only a `socketpair` created immediately before spawning each child, with the child inheriting its endpoint (`2026-07-16-mvp-ipc-manifest-seam-contract.md:17-35`). Once app-main dies, a replacement process cannot recover its lost socketpair endpoint or create a new connection to the surviving broker. Conversely, if broker supervision kills and respawns the broker with app-main, the placement grill's stated failure-isolation benefit that app-main crash leaves the seat channel up is false (`transport-broker.md:270-276`).

Even if an undisclosed listener is assumed, `state_seq > installed` proves freshness ordering, not control-peer authenticity or singleton ownership. The worker handshake also says the worker presents `{run_id,generation_id,turn_epoch}` "from its m-10 assign" (`:169-170`), while the current m-10 `assign` contains only `{run_id,turn_epoch,manifest_digest}` (`mvp-ipc-manifest-seam-contract.md:47-54`).

Required revision:

1. Choose one crash/adoption model explicitly. To preserve the accepted own-process rationale, define how a replacement app-main discovers and connects to the surviving broker.
2. Pin the control endpoint lifecycle and peer proof: listener/socket ownership and mode, app-main singleton/store lock, peer credential or launch/adoption token, stale-old-session exclusion, process discovery, reconnect deadlines, and cleanup. `state_seq` remains the state freshness check; it cannot be the authenticator.
3. If the actual decision is that broker dies with app-main, revise the grill consequence, restart matrix, and control bootstrap accordingly and route any changed placement rationale through the existing review boundary.
4. Reconcile the worker attach contract with m-10: add `generation_id` to the exact `assign` shape (and the direct broker endpoint/attach material needed by the worker), or replace the claimed source with a different authenticated handoff. Obtain m-10 consumer confirmation on the same bytes.
5. Add app-main-crash-with-broker-surviving, stale app-main racing replacement, lost socketpair endpoint, and assign/attach parity fixtures.

### R2-F3 - Mandatory recording is impossible during the control outage that requires it

Section 2.4 says control loss suspends admission and suppresses pushes, with those dispositions recorded (`transport-broker.md:110-117`). Section 2.11 makes m-10 the only durable writer and requires durable acknowledgement for `fence_reject`, `forward_suppressed`, `epoch_installed`, `control_handover`, `reauth`, retry fencing, and cross-epoch completion (`:174-181`). When the control session is down, the broker has no path to that writer. The outage rule addresses only an already-admitted operation's completed response; it does not define how control-loss rejection, push suppression, failed handover, or re-auth events satisfy their mandatory-recorded claim.

The proposed `broker_event` carrier is also not yet a total protocol. It gives one superset object but no per-event required/forbidden fields, stable event identity, duplicate/replay rule, acknowledgement correlation, or durable table/unique key. m-10's current store census has no broker-event row family (`mvp-ipc-manifest-seam-contract.md:162-177`). A commit followed by lost ack therefore has no pinned resend/dedup behavior, and a broker restart resets any merely connection-local sequence interpretation.

Required revision:

1. Define a closed per-event schema table with required/absent fields, canonical enum values, a globally stable idempotency key such as `{broker_instance_nonce,event_seq}`, and the exact ack frame (`re` correlation after durable commit).
2. Pin retry/replay and dedup: m-10 stores the unique event key and returns the same committed acknowledgement for duplicate delivery; broker restart and control handover cannot alias event identities.
3. Add the exact m-10 durable row/table obligation and obtain m-10 confirmation against its sole-writer/schema contract.
4. Totalize every no-writer window. For caller operations, return `broker:record-unavailable` when the required record cannot be acknowledged. For push/control/reauth events with no caller, either durably spool them somewhere with reviewed ownership or narrow the claim honestly; "suppressed-and-recorded" cannot hold while the sole writer is unreachable.
5. Add commit-before-ack-loss, duplicate resend, broker restart sequence reuse, control-loss fence rejection, control-loss push suppression, and m-10 recovery fixtures.

### R2-F4 - The relay-leg evidence reference still lacks the canonical encoding F68 assigned to m-7

The serve stamp is substantially pinned, but §3.3's `relay_leg_evidence` remains pseudostructure: shorthand identity fields without types, no schema/version, no canonical serialization rule, an ellipsis-bearing array with no order/cardinality/duplicate law, and no unknown/missing-field disposition (`transport-broker.md:227-236`). F68 assigned m-7 the relay-leg evidence-reference encoding, not only its conceptual field list. Master+VP own the composite join, but they cannot byte-bind this half until m-7 supplies exact bytes.

There is also a byte-level ambiguity in the stamp: §3.2 calls the file RFC 8785 JCS and newline-terminated (`:206-220`). A trailing newline is not part of JCS output. The contract must distinguish canonical payload bytes from the file terminator or remove the terminator; otherwise a consumer cannot both require byte-canonical JCS and accept the specified file.

Required revision:

1. Pin `relay_leg_evidence` as a versioned closed object with exact field types, digest/nonce/relay-ID grammars, canonical encoding, unknown/missing rejection, and an ordered unique `relay_records` array. State the deterministic order and required leg cardinality/roles for the one governed exchange.
2. Bind each item to the exact governed record identity needed for verification; define kind/role semantics and reject duplicate, missing, wrong-kind, or instance-mismatched references.
3. Resolve stamp bytes explicitly, for example `file_bytes = JCS(stamp_object) || 0x0a` with a consumer that requires exactly one terminal LF and byte-compares the prefix to regenerated JCS, or use bare JCS with no LF.
4. Add mutation vectors for array reorder/duplicate, unknown kind/field, missing leg, malformed digest/nonce, stale instance, record mismatch, and newline variance.

## Accepted portions

- **R1-F1 closes:** the caller has exactly the three relay calls plus typed metadata `Describe`; Describe uses the existing channel method, is fenced, stays out of the dispatch set, and leaves interpretation with m-2 (`:42-65,110-118,148-159`).
- **R1-F2 closes on the m-7 side:** env custody is removed; the descriptor-safe source, exact grammar, atomic rotation, locator/capability split, honest capability leak claim, connection scope, and negative census are materially complete (`:82-108,183-190`). Final m-1 consumer confirmation remains externally held.
- **R1-F3 closes except for R2-F1's read classification:** retryable failures are closed, served/context/frame errors do not retry, retries re-enter the epoch fence, and concurrent redial/push-reader replacement is single-flight (`:67-78,119-128`).
- **The accepted F67 decisions remain closed:** own-process placement and complete-and-deliver are unchanged and still supported. R2-F2 asks for the missing realization of that placement across supervisor restart, not a different placement.
- **The F65 stamp process binding is directionally accepted:** PID/socket, instance nonce, process-start time, executable identity, config identity, write-before-accept, and stale-stamp handling close the prior mutable-path/bystander defect (`:191-225`). R2-F4 is the remaining exact-byte carrier issue.
- **The route-back result remains clean** if the revision adds no broker-local durable spool or new identity authority. Choosing either would require the named route-back rather than silent expansion.

## Revision bar

Return fresh contract bytes and a fresh hash that:

1. States `relay.read`'s quarantine side effect and gives it an honest retry/unknown-outcome contract with executable fixtures.
2. Makes own-process broker control adoption and worker attach implementable against one exact m-10 IPC contract.
3. Gives every mandatory broker event a total durable-write/ack/outage protocol and m-10 schema home.
4. Completes the F68 relay-leg reference and stamp file as deterministic canonical bytes.
5. Refreshes §E against the latest m-1 relay trail and still withholds m-1 consumer confirmation until pair approval.
6. Preserves all accepted F1-F5 folds, the grill lock, the F57 ceiling, m-10 no-seat/no-secret rails, and the no-conductor-protocol/store-change result.

On revision, issue a fresh uniquely-parented DESIGN for exact-byte re-review. Consumer confirmations remain held.

## Verification

- `shasum -a 256 master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md` -> `8cf86753eaa711715ab3f063fda7d76efd4910cd8184eb066ff1588654ce39cc`.
- Exact incoming relay `master/relays/step3-mvp-design-m7/DESIGN-planner-20260716-052315.md` lints OK; routing, parent, `DESIGN_DOC_ID`, and grill lock match.
- Checked current m-1 relay trail through `DESIGN-planner-20260716-053040.md`; m-1 pair approval is still absent and its consumer edge remains held.
- Checked m-10 IPC/store contract `master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md:17-35,47-54,73-80,162-177`.
- Checked live conditional read mutation at `frank/cmd/frank/main.go:463-480` and `frank/internal/engine/loop.go:74-80,111-127` at clean `frank@502e06c`.

ACTIONS_GIT_REF: docs-workspace action only - created this review relay and appended one `master/relays/INDEX.md` row; no governing contract/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `502e06c`.
RELAY_LINT: OK - exact-file relay-lint.py exit 0
Next requested action: m-7.planner revises only R2-F1 through R2-F4, refreshes the m-1 edge, re-hashes the contract, and sends fresh uniquely-parented final bytes for pair re-review; consumer confirmation remains held.
