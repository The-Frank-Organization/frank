# H-17 — the governed-effect census (v3, canonical machine-readable; stage-6 lock input)

**Assembled by master 2026-07-21 per the VP's F91 r2 correction (`step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-012112`): ONE canonical machine-readable table, uniform row grammar, every row full-field under schema v1 (`master/H17-CENSUS-SCHEMA.md` @ `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`), every `policy_artifact` an exact document/digest, all non-append verbs split.**

## Canonical grammar (F91-r2 issue 1)
Every effect is one record block of **exactly 21 lines**, `<field>: <value>`, in schema-v1 field order (`effect_id` … `threat_claim_scope`), blocks separated by `---`. Every value is exact or a schema-legal null token (`not specified` / `residual` / `none (self-reported)` / `ambient` / `not applicable`). One grammar for all rows — the conductor rows (§A, master-authored), the m-10 rows (§B), and the m-9 rows (§C) all use this identical block form. The m-9 and m-10 owner rows are **normalized into this grammar from their approved artifacts without moving one owner-design byte**: m-10's 18 rows from `master/domains/m-10-app-control-plane/design/2026-07-19-mvp-control-plane.md` §11a @ `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf`; m-9's 16 rows from `master/domains/m-9-model-runtime/design/2026-07-19-mvp-full-worker.md` §11.5 @ `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45` (the header `effect_id` promoted to the named field; the combined `policy_owner/policy_artifact` split into the two fields; every truncated digest expanded to full 64-hex). A refresh re-normalizes from those exact hashes.

## Row count + F91-r2 dispositions
- **42 governed-effect rows** = 8 conductor-plane (§A) + 18 m-10 (§B) + 16 m-9 (§C). 42 unique `effect_id`s.
- **F91-r2 issue 2 (exact policy_artifact):** every `policy_artifact` cell carries a full 64-hex design digest and/or the exact build commit `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75` (never the `s11-close` tag alias, never a truncated digest).
- **F91-r2 issue 3 (verb/effect split):** conductor read/serve is split into `conductor-project-serve` / `conductor-read-serve` / `conductor-describe-serve` (distinct disclosure targets); the store boot effect is split into `conductor-store-genesis` (once-per-store) / `conductor-store-gc` (idempotent).
- **Provider-send dedup:** the wire SEND is exactly `m9-provider-attempt-send` (E8); `m10-provider-attempt-recording` is the distinct durable-row-commit effect. No invented `m8-provider-send`.
- **Claim boundary:** "governed" = the relay plane + the F59 tool-dispatch binding + the designated provider route + the app-control-plane lifecycle/store effects; local tool side effects beyond F59's invocation binding and bash egress are named residuals inside their rows.
- Non-effect rationales (merge/deploy/release/`-mint`) are in the §D appendix, outside the effect set.

---

## §A. Conductor-plane rows (8; master-authored; the Step-2 substrate + the H-16 rev16 outcome contract)

effect_id: conductor-relay-accept
effect_class: durable_state_commit
requester: any seat (via the authenticated channel)
executor: the conductor engine
authority_source: the ten INV-CATALOG laws + form/lineage/observe checks
policy_owner: m-1 (identity/store) + m-2 (form) + m-3 (observe)
policy_artifact: `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75` (the Step-2 `s11-close` engine) + the H-16 outcome contract `a349a3297c608915510e7bf6de2e304b824897aae4184f7c71db928149ae316d`
decision_point: the serialized commit loop (identity stamp → form → lineage → observe → outcome classification)
enforcement_point: the single serialized commit loop (sole governed writer)
exclusive_credential_holder: the seat channel (m-1 channel-stamp; payload identity never trusted)
request_freeze_point: intake
authorization_linearization_point: the store commit
effect_linearization_point: the store commit (the append-only record)
outcome_reporter: the conductor engine (self)
outcome_observer: observe-as-send (E1/E2 — local harness observation, not independent attestation)
outcome_validator: the observe gate + the invariant checks
canonical_record: the append-only relay store
bypass_paths: same-UID direct store-file access (accepted residual — "private" is a confusion discipline, not an OS boundary)
failure_unknown_semantics: the H-16 rev16 monotonic split — decision_state ∈ {accepted, rejected, held} fixed at commit and NEVER relabeled, × post_commit_state ∈ {complete, pending, failed, unknown}; the legacy `state` field present IFF post_commit_state == complete (fail-closed for legacy decoders); a post-commit hook failure becomes a durable derived-work fault, never a decision relabel
replay_idempotency: intake-ID idempotency (duplicate intake / restart serve the committed decision)
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: conductor-project-serve
effect_class: read_serve
requester: any seat
executor: the conductor engine
authority_source: the seat channel + the inbox render-scoping (seat_scope)
policy_owner: m-1 (addressing)
policy_artifact: `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75` (the `project` handler)
decision_point: the inbox projection (recipient-scoped)
enforcement_point: the `project` handler (I-PH: no store paths on the seat surface)
exclusive_credential_holder: the seat channel
request_freeze_point: not applicable (read; no state change)
authorization_linearization_point: not applicable (non-append)
effect_linearization_point: the read (no durable effect)
outcome_reporter: the conductor engine
outcome_observer: none (self-reported)
outcome_validator: the recipient-inbox scoping rule
canonical_record: read-history (feeds computed lineage)
bypass_paths: same-UID store-file read (accepted residual)
failure_unknown_semantics: typed empty/absent inbox result; no partial state (read)
replay_idempotency: idempotent (read)
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: conductor-read-serve
effect_class: read_serve
requester: any seat
executor: the conductor engine
authority_source: the seat channel + observability of the accepted graph
policy_owner: m-1 (addressing)
policy_artifact: `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75` (the `read` handler)
decision_point: the accepted-graph read (relay-id-scoped)
enforcement_point: the `read` handler (I-PH: no store paths)
exclusive_credential_holder: the seat channel
request_freeze_point: not applicable (read)
authorization_linearization_point: not applicable (non-append)
effect_linearization_point: the read (no durable effect); read-history is appended as lineage input
outcome_reporter: the conductor engine
outcome_observer: none (self-reported)
outcome_validator: the accepted-graph visibility rule
canonical_record: read-history (the lineage-input record — the one read that mutates derived lineage state)
bypass_paths: same-UID store-file read (accepted residual)
failure_unknown_semantics: typed absent-relay result; no partial state
replay_idempotency: idempotent (read); read-history append is set-idempotent per (seat, relay)
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: conductor-describe-serve
effect_class: read_serve
requester: any seat
executor: the conductor engine
authority_source: the seat channel + the form registry
policy_owner: m-2 (form schema/render-scoping)
policy_artifact: m-2 `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d` (the FieldSpec/Describe contract) + `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75` (the `Describe` handler)
decision_point: the form-render projection (seat-scoped `visible_when`)
enforcement_point: the `Describe` handler (render-scoping by seat)
exclusive_credential_holder: the seat channel
request_freeze_point: not applicable (read)
authorization_linearization_point: not applicable (non-append)
effect_linearization_point: the form-schema serve (no durable effect)
outcome_reporter: the conductor engine
outcome_observer: none (self-reported)
outcome_validator: the render-scoping rules (seat_scope + visible_when)
canonical_record: none (a pure serve; the form registry is the source, not a record)
bypass_paths: same-UID store-file read (accepted residual)
failure_unknown_semantics: typed absent/invalid-form result; no partial state
replay_idempotency: idempotent (read); deterministic given the form registry + seat
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: conductor-broker-attach
effect_class: credential_operation
requester: the m-9 worker generation (via the broker)
executor: the m-7 broker (outside the replaceable worker generation)
authority_source: the F60/F64 seat-credential model
policy_owner: m-1 (identity/lifecycle) + m-7 (channel/broker)
policy_artifact: m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`
decision_point: the three-token attach taxonomy + the per-verb turn_epoch fence
enforcement_point: the broker attach gate (every relay verb + push checks the current epoch)
exclusive_credential_holder: the m-7 broker (one credential per LOGICAL seat; never copied into a worker generation)
request_freeze_point: the attach presentation
authorization_linearization_point: the attach-result commit
effect_linearization_point: the attach-result commit
outcome_reporter: the broker (self)
outcome_observer: the broker's §2.5 legs (independent of the worker generation)
outcome_validator: the three-token totality + the epoch fence
canonical_record: the attach event rows
bypass_paths: same-UID inspection of broker/socket state (accepted residual — inspection, not forge-by-construction; D5)
failure_unknown_semantics: broker:attach-suspended (transient hold) / broker:attach-tuple-mismatch (terminal-for-generation) — honest typed terminals, never a masked mint
replay_idempotency: epoch-fenced; a stale generation is refused on every verb
threat_claim_scope: confusion-not-malice; same-UID residual accepted; forgery-robust-by-construction is the D3-shelved milestone, NOT claimed here

---

effect_id: conductor-outbox-send
effect_class: external_send
requester: a governed lane (park/wake, human-surface bucket)
executor: the conductor local-outbox
authority_source: the observe-gated accepted relay + the m-6 bucket rules
policy_owner: m-6 (human surface/scheduler)
policy_artifact: `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75` (the s10/s11 comms machinery)
decision_point: the bucket projection + the observe gate
enforcement_point: the local-outbox-only send path (the conductor holds no live external credential in the MVP)
exclusive_credential_holder: not specified (the MVP external-send credential surface is Step-4 m-6 work; the MVP outbox is local-only)
request_freeze_point: the accepted relay
authorization_linearization_point: the store commit that accepts the sending relay
effect_linearization_point: the local-outbox write (external delivery is out-of-band, un-attested in the MVP)
outcome_reporter: the outbox (self)
outcome_observer: none (self-reported)
outcome_validator: the observe gate on the sending relay
canonical_record: the outbox rows + the sending relay
bypass_paths: same-UID store/outbox access (accepted residual)
failure_unknown_semantics: an undelivered outbox item is durable and retryable; no false "delivered" claim
replay_idempotency: outbox-row idempotency
threat_claim_scope: confusion-not-malice; the MVP does not attest external delivery

---

effect_id: conductor-store-genesis
effect_class: durable_state_commit
requester: the conductor boot path (first boot of a store)
executor: the conductor engine (single writer)
authority_source: the trusted config load + the genesis/version rule
policy_owner: m-1 (store) + m-7 (conductor-core)
policy_artifact: `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75` (the genesis/version boot path)
decision_point: the genesis/version check at first boot (schema at user_version; higher ⇒ refuse-to-serve; lower ⇒ forward-only migration; failed integrity ⇒ refuse + operator disposition)
enforcement_point: the single boot transaction under the root lock
exclusive_credential_holder: the root lock (`conductor.lock`, phase −1, held for the writer lifetime)
request_freeze_point: the boot transaction
authorization_linearization_point: the root-lock acquisition
effect_linearization_point: the genesis/migration commit
outcome_reporter: the conductor engine (self)
outcome_observer: none (self-reported)
outcome_validator: the integrity check + the version rule
canonical_record: the store schema/version
bypass_paths: same-UID store-file access (accepted residual); NOTE — the `-mint` path's pre-lock writer is the OPEN defect H-26 (`master/FRANK-HARDENING-BACKLOG.md`), not a bypass of THIS row's phase-−1 discipline
failure_unknown_semantics: refuse-to-serve on version-higher/integrity-fail (honesty over availability; no auto-rebuild of a corrupt store)
replay_idempotency: once-per-store (genesis runs exactly once; a second boot takes the recovery/serve path, never re-genesis)
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: conductor-store-gc
effect_class: durable_state_commit
requester: the conductor GC path (steady-state maintenance)
executor: the conductor engine (single writer)
authority_source: the GC retention policy
policy_owner: m-1 (store) + m-7 (conductor-core)
policy_artifact: `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75` (the GC path)
decision_point: the retention predicate over committed state
enforcement_point: the GC transaction under the root lock (never mutates the accepted graph's truth)
exclusive_credential_holder: the root lock (`conductor.lock`, held for the writer lifetime)
request_freeze_point: the GC transaction
authorization_linearization_point: not applicable (GC removes only retention-eligible non-canonical state; no authority decision)
effect_linearization_point: the GC commit
outcome_reporter: the conductor engine (self)
outcome_observer: none (self-reported)
outcome_validator: the retention predicate (canonical rows never eligible)
canonical_record: the GC audit note (the accepted graph is never GC'd)
bypass_paths: same-UID store-file access (accepted residual)
failure_unknown_semantics: a partial GC is safe (idempotent; retention-eligible only); crash ⇒ re-run
replay_idempotency: idempotent (re-running GC over the same state is a no-op after the first pass)
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

## §B + §C. Owner rows (18 m-10 + 16 m-9; normalized into the canonical grammar from the approved artifacts)

effect_id: m10-run-admission
effect_class: durable_state_commit
requester: operator (terminal `run start`)
executor: m-10 applier
authority_source: operator grant + the ratified 8-name constant
policy_owner: operator + the manifest field owners (m-2/m-3/m-8)
policy_artifact: contract §C.1 @ r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146` + amendment r7 @ `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`
decision_point: terminal command acceptance
enforcement_point: the admission transaction (structural H-15 checks)
exclusive_credential_holder: none (the `credential_ref` rides opaque)
request_freeze_point: the admission commit
authorization_linearization_point: the admission commit
effect_linearization_point: the admission commit (one transaction)
outcome_reporter: m-10 (self)
outcome_observer: the external E3 evaluator for the manifest-digest fact (m-3-bound); otherwise none (self-reported)
outcome_validator: the structural admission checks
canonical_record: `runs`
bypass_paths: same-UID direct store write (§9)
failure_unknown_semantics: typed refusal to the operator (shape `not specified` at contract grain — no closed family claims it, R40-F1; the sizing-gate overflow member is the separate `m10-turn-admission` boundary); no partial state (one transaction)
replay_idempotency: none — a re-start mints a NEW `run_id`
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: m10-run-stop-cancel
effect_class: durable_state_commit
requester: operator
executor: m-10 applier
authority_source: operator grant
policy_owner: m-10
policy_artifact: contract §F `cancellations` + §B.2 D-5  · exact artifact binding: m-10 contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`
decision_point: terminal command acceptance
enforcement_point: the lifetime-idempotent cancellation row + the D-5-composing chokepoint transaction
exclusive_credential_holder: none
request_freeze_point: the cancellation-row commit
authorization_linearization_point: the cancellation-row commit
effect_linearization_point: the terminal-writing chokepoint transaction (cancel-ack composed, never substituting)
outcome_reporter: m-10 (self)
outcome_observer: none (self-reported)
outcome_validator: UNIQUE`(target_kind, target_id, epoch)` + the D-5 legality predicates
canonical_record: `cancellations` + `runs`/`turns`
bypass_paths: same-UID direct store write (§9)
failure_unknown_semantics: crash ⇒ PENDING cancellation recovered; never a silently-released lease
replay_idempotency: lifetime-idempotent per key
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: m10-cancel-intent-send
effect_class: control_send
requester: the run-stop/cancel machinery
executor: m-10
authority_source: the durable cancellation row (proven current-epoch intent — the contract §B.1 intent predicate; bare channel death is never intent)
policy_owner: m-10 (intent) / m-8 + m-9 (the cancel cuts)
policy_artifact: contract §B.1/§F; m-8 r12 §1.3  · exact artifact binding: m-10 contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; m-8 r12 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`
decision_point: the cancellation disposition
enforcement_point: `not specified` — the exact wire frame carrying cancel intent toward the children is not pinned by the frozen contracts at this plane (the RESULT side is frozen: m-8's `cancelled(cancel_point)` on CTRL-C)
exclusive_credential_holder: none
request_freeze_point: the durable intent row
authorization_linearization_point: the cancellation-row commit
effect_linearization_point: `not specified` — the local send fact is genuinely unfrozen; the counterparties' cancel cuts linearize in THEIR rows (cross-reference only)
outcome_reporter: m-8/m-9 (their result frames)
outcome_observer: none (self-reported by the counterparties)
outcome_validator: the §B.1 total cancellation-consumption table
canonical_record: `cancellations` + the terminal attempt/turn rows
bypass_paths: same-UID direct store write (§9)
failure_unknown_semantics: unproven intent never mints CANCELLED (channel death ⇒ UNKNOWN)
replay_idempotency: lifetime-idempotent identity; duplicate results equivalence-keyed
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: m10-app-main-recovery
effect_class: process_lifecycle + durable_state_commit
requester: OS/operator restart
executor: m-10 boot path
authority_source: the durable store itself (the contract §B.3 recovery matrix)
policy_owner: m-10
policy_artifact: contract §B.3/§B.5  · exact artifact binding: m-10 contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`
decision_point: matrix branch selection over committed state
enforcement_point: the keyed recovery transactions
exclusive_credential_holder: none
request_freeze_point: `not specified` (recovery reads, then transitions)
authorization_linearization_point: the keyed recovery-transaction commits
effect_linearization_point: the same commits
outcome_reporter: m-10 (self)
outcome_observer: m-7's broker for the transition legs (a second party's durable acks); local rows none (self-reported)
outcome_validator: the branch predicates + `broker_instance_nonce` recognition
canonical_record: `epochs`/`epoch_transitions`/`workers`
bypass_paths: same-UID direct store write (§9)
failure_unknown_semantics: integrity refusal fail-closed; the §B.5 crash matrix total; children EOF-fail-closed, never adopted
replay_idempotency: one-mint by `(run_id, retiring generation_id)`; same-ID transition acks idempotent
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: m10-worker-spawn-assign
effect_class: process_lifecycle
requester: the scheduler (initial or paired replacement)
executor: m-10 spawn path
authority_source: run admission + the contract §B.4 lifecycle rules
policy_owner: m-10
policy_artifact: contract §B.1/§B.4  · exact artifact binding: m-10 contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`
decision_point: the allocation + lease-grant transactions
enforcement_point: the sanitized-env spawn (allow-list; no secrets/handles in argv/env/FDs)
exclusive_credential_holder: the seat credential is m-7-broker-held (F60/F66) — the worker spawns holding NOTHING
request_freeze_point: `assign` content from durable state at lease-bind
authorization_linearization_point: the lease-grant commit
effect_linearization_point: OS process creation (pid)
outcome_reporter: the child (`hello`) + the OS
outcome_observer: the OS process table via `waitpid` (independent of the child's claims)
outcome_validator: `HANDSHAKE_DEADLINE` + hello grammar
canonical_record: `workers`
bypass_paths: same-UID ambient process creation exists but cannot create a SUPERVISED child (rows require the applier); same-UID direct store write (§9) could forge rows
failure_unknown_semantics: spawn-fail ⇒ FAILED wash-out same-epoch; no-`hello` ⇒ FAILED
replay_idempotency: generation ids never reused
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: m10-connector-spawn
effect_class: process_lifecycle
requester: run start / the paired-replacement path (§4 — no connector-only path exists)
executor: m-10 spawn path
authority_source: the contract §B.4 canonical order (connector-first; pair replacement)
policy_owner: m-10
policy_artifact: contract §A.1/§B.1/§B.4  · exact artifact binding: m-10 contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`
decision_point: the incarnation-allocation disposition
enforcement_point: the sanitized-env spawn (no secret bytes at spawn — the opaque ref travels later in `connector_assign`)
exclusive_credential_holder: provider credentials are m-8's at runtime; at spawn the child holds nothing
request_freeze_point: `not specified` (no assign content exists at spawn)
authorization_linearization_point: the incarnation-allocation commit
effect_linearization_point: OS process creation (pid)
outcome_reporter: the child (`hello`) + the OS
outcome_observer: the OS process table via `waitpid`
outcome_validator: `HANDSHAKE_DEADLINE` + hello grammar
canonical_record: connector incarnation state (§F)
bypass_paths: same-UID ambient process creation (unsupervised); same-UID direct store write (§9)
failure_unknown_semantics: spawn-fail/no-`hello` ⇒ incarnation FAILED ⇒ ONE G-2 counter increment ⇒ the §4 paired retry under the frozen order
replay_idempotency: incarnations never reused
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: m10-attach-gate
effect_class: admission (NOT authorization — the F87 distinction)
requester: the worker's `attach_result` report
executor: m-10 (the first-`turn_open` hold)
authority_source: m-7's r11 D-3 tokens byte-exact
policy_owner: m-7
policy_artifact: m-7 r11 @ `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572` + contract §B.2  · exact artifact binding: m-10 contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`
decision_point: receipt of `attach-ok`
enforcement_point: turn admission blocked until receipt
exclusive_credential_holder: the broker holds the verifying tuple (feed = verifier only)
request_freeze_point: `not specified`
authorization_linearization_point: the receipt commit
effect_linearization_point: the attach-ok receipt commit — the hold's own release fact (admission itself linearizes at `m10-turn-admission`)
outcome_reporter: the worker (self-reporting its broker outcome)
outcome_observer: none (self-reported) at m-10 — the broker independently ENFORCES per-operation at serve time (defense-in-depth, not observation)
outcome_validator: the closed three-token check (unknown ⇒ fault)
canonical_record: `workers` attach state
bypass_paths: none for admission itself; same-UID direct store write (§9)
failure_unknown_semantics: `ATTACH_DEADLINE` ⇒ FAILED; tuple-mismatch ⇒ FAILED no-retry
replay_idempotency: once per generation
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: m10-turn-admission
effect_class: durable_state_commit
requester: operator input OR an accepted wake row (via `m10-wake-acceptance`)
executor: m-10 scheduler
authority_source: the amendment's one-run/one-turn rule + the admission gates
policy_owner: m-10
policy_artifact: contract §B.2  · exact artifact binding: m-10 contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`
decision_point: the gate predicates (lease ∧ READY ∧ attach ∧ no-active-turn ∧ non-terminal run)
enforcement_point: the ONE admission transaction
exclusive_credential_holder: none
request_freeze_point: the `turns` row
authorization_linearization_point: the admission commit
effect_linearization_point: the admission commit (the `turn_open` command emitted post-commit)
outcome_reporter: m-10 (self)
outcome_observer: none (self-reported)
outcome_validator: the gate predicates
canonical_record: `turns` (+ the `wake_schedule` flip when wake-driven)
bypass_paths: same-UID direct store write (§9)
failure_unknown_semantics: the closed single-member `admission_refused{reason: task_input_frame_overflow}` (contract r40, the turn-input sizing gate ALONE, emitted before any admission transaction); other structural admission refusals stay typed-refusal realization surface, shape `not specified` at contract grain (no family claims them — R40-F1); **crash BEFORE the admission commit ⇒ the wake row stays `pending` (admitted once, later); crash AFTER the commit, before the `turn_open` send ⇒ the row is already `dispatched` and recovery re-emits byte-identically from the committed row — never re-consumed, no second task identity (the r9 two-cut correction)**
replay_idempotency: at-most-once per `relay_id` (the same-commit flip); operator admissions are fresh acts
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: m10-wake-acceptance
effect_class: durable_state_commit (the §6 link-2 transaction, distinct from link-3 admission — M10-S5-R3-F1)
requester: m-9 (`wake_forward{relay_id}` — the only conductor reader, contract §E)
executor: m-10 applier
authority_source: contract §E (advisory input; acceptance is unconditional recording)
policy_owner: m-10
policy_artifact: contract §E/§F  · exact artifact binding: m-10 contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`
decision_point: none — every well-formed forward is recorded (no allow/deny exists on an advisory leg)
enforcement_point: UNIQUE(`relay_id`) at the insert
exclusive_credential_holder: none
request_freeze_point: the `wake_schedule` row
authorization_linearization_point: `not specified` — no authorization leg exists (recording an advisory fact)
effect_linearization_point: the idempotent insert commit
outcome_reporter: m-10 (self)
outcome_observer: none (self-reported)
outcome_validator: frame grammar (§A.2)
canonical_record: `wake_schedule` (disposition `pending`)
bypass_paths: same-UID direct store write (§9)
failure_unknown_semantics: a dropped forward ⇒ recovered by m-9's rediscovery; malformed ⇒ channel fault
replay_idempotency: duplicate forwards no-op on UNIQUE
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: m10-worker-retirement-epoch-mint
effect_class: durable_state_commit + process_lifecycle
requester: supervision events (crash, health-timeout, invariant fault, the §D.2 (5) lease-fault)
executor: m-10 applier
authority_source: the ratified fail-closed rules
policy_owner: m-10
policy_artifact: contract §B.3/§B.4  · exact artifact binding: m-10 contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`
decision_point: the disposition classification (the total tables)
enforcement_point: the ONE atomic retirement transaction (fence + state-sensitive park + E+1 + pre-allocated G+1 + ledger row)
exclusive_credential_holder: none
request_freeze_point: `not specified`
authorization_linearization_point: the retirement commit
effect_linearization_point: the retirement commit; the broker install follows §B.5
outcome_reporter: m-10 (self)
outcome_observer: m-7's broker for the §B.5 handshake legs; the local parking none (self-reported)
outcome_validator: the committed-once key + the §B.5 substate totality
canonical_record: `workers`/`epochs`/`epoch_transitions`/`crossing_ops` + parked rows
bypass_paths: same-UID direct store write (§9)
failure_unknown_semantics: the recognition×commit matrix total; crash converges on ONE mint
replay_idempotency: keyed committed-once; same-ID acks idempotent
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: m10-epoch-publication
effect_class: serve/publication (a non-append effect, its own row per F87)
requester: the mint transaction
executor: m-10 (generation-blind `epoch_update` on CTRL-C; full `epoch_state` to the broker feed)
authority_source: source-specific epoch authority — m-10 is the sole mint (contract §B.4)
policy_owner: m-10
policy_artifact: contract §B.4  · exact artifact binding: m-10 contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`
decision_point: post-commit publication
enforcement_point: ordering (install-precedes-lease-bind; the R9 pending-transition snapshot withholding) — the VALUE fences live at the consumers by design
exclusive_credential_holder: none
request_freeze_point: the committed epoch row
authorization_linearization_point: the mint commit
effect_linearization_point: the post-commit emission — the CTRL-C `epoch_update` send + the feed publication, this row's own serve act (consumer receipt facts live with the consumers; the broker's install ack remains this row's transition validator)
outcome_reporter: m-10
outcome_observer: none independent — the consumers' own fences are cross-checks, not neutral observation
outcome_validator: the broker's install ack for transitions
canonical_record: `epochs` + `broker_events`
bypass_paths: none on the authority channel by construction (a peer-presented epoch never advances any cache); same-UID direct store write (§9) could forge the rows
failure_unknown_semantics: lost `epoch_installed` ⇒ idempotent re-delivery (R10-F1)
replay_idempotency: idempotent same-key
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: m10-connector-assign-credential
effect_class: credential_operation (reference-only at this plane)
requester: connector `hello`
executor: m-10 (`connector_assign`)
authority_source: the frozen manifest (operator-selected ref)
policy_owner: m-1
policy_artifact: m-1 @ `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c` §1.4a + contract §B.1/§C.1  · exact artifact binding: m-10 contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`
decision_point: at admission (the ref froze with the manifest)
enforcement_point: presence/grammar check at assign + m-8's bootstrap verification (their bytes)
exclusive_credential_holder: m-8 at runtime (secret bytes); m-10 holds the opaque reference only, never resolves
request_freeze_point: the admission commit
authorization_linearization_point: the assign send post-`hello`
effect_linearization_point: the local `connector_assign` send post-commit — the reference-orchestration effect this row names (m-8's resolve linearizes in THEIR row)
outcome_reporter: m-8 (`connector_ready`)
outcome_observer: none (self-reported) — the E3 evaluator externally checks digests, not credential possession
outcome_validator: m-8's loaded-policy/catalog hash verification
canonical_record: `runs`.manifest + connector incarnation state
bypass_paths: same-UID credential FILE read (accepted residual — frank's claim is *frank never moves secret bytes*, not OS isolation); same-UID direct store write (§9)
failure_unknown_semantics: mismatch ⇒ `connector_ready` withheld ⇒ zero send, no admission
replay_idempotency: once per incarnation; channels never re-established
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: m10-f59-ticket-issue
effect_class: durable_state_commit (the authorization decision)
requester: m-9 (`authorize_tool_call`)
executor: m-10 applier (the §D.2 (0)–(7) procedure)
authority_source: the operator-ratified 8-name constant via the frozen manifest (F55/F59)
policy_owner: the operator (the constant); m-10 (the procedure)
policy_artifact: contract §C.3/§D.2  · exact artifact binding: m-10 contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`
decision_point: the total ordered procedure
enforcement_point: the same ONE transaction (commit-before-reply)
exclusive_credential_holder: none (a ticket is never a bearer token)
request_freeze_point: the serve set at manifest freeze; ticket identity #1 at mint
authorization_linearization_point: the ISSUED/VOID commit
effect_linearization_point: the ISSUED/VOID commit — the row's OWN durable effect (the invocation is m-9's row; issue-authorizes-consume/invoke is a cross-reference, never this row's fact)
outcome_reporter: m-10 (self; the reply)
outcome_observer: none (self-reported)
outcome_validator: the procedure's predicates
canonical_record: `tool_authorizations` (`void_reason`; denials-are-rows)
bypass_paths: none by schema on the wire path (the only path to a ticket is this frame); same-UID direct store write (§9)
failure_unknown_semantics: the total rejection/fault table; §D.4 expiry
replay_idempotency: the §D.2 (0) stored-row mapping
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: m10-f59-consume-gate
effect_class: durable_state_commit (the authorization→effect gate for the governed tool path)
requester: m-9 (`consume_ticket` carrying comparand #2)
executor: m-10 (the gate alone — m-9's invoke is counterparty census territory, cross-referenced)
authority_source: the ISSUED ticket (identity #1)
policy_owner: m-10
policy_artifact: contract §D.3  · exact artifact binding: m-10 contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`
decision_point: the sender-fenced success predicate
enforcement_point: the single conditional UPDATE (no `consume_ok`, no execution — m-9's frozen half)
exclusive_credential_holder: `ambient` for the underlying OS effect (see bypass)
request_freeze_point: #1 at mint; comparands #2/#3 fresh (F84)
authorization_linearization_point: the CONSUMED commit
effect_linearization_point: the CONSUMED commit — the row's OWN effect (m-9's invocation linearizes in THEIR row, point B onward)
outcome_reporter: m-9 (self-report via `record_tool_outcome`)
outcome_observer: none (self-reported) — m-10 validates the report; it does NOT independently observe the invocation
outcome_validator: m-10 (the §D.4 predicates)
canonical_record: `tool_authorizations` + `tool_calls`
bypass_paths: bash-ambient (accepted, amendment §2a — the gate proves *authorized == executed* for IDENTIFIED calls; arbitrary `bash` content is same-UID ambient authority outside the fence); same-UID direct store write (§9)
failure_unknown_semantics: `UNKNOWN_TOOL_OUTCOME` park; `not_invoked_integrity_fault` definite-no-effect
replay_idempotency: one-shot atomic; `DUPLICATE_CONSUME` spent path
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: m10-f59-outcome-record
effect_class: durable_state_commit (recording, not effecting)
requester: m-9 (`record_tool_outcome`)
executor: m-10 applier (the §D.4 (0)–(8) table)
authority_source: the CONSUMED ticket
policy_owner: m-10
policy_artifact: contract §D.4  · exact artifact binding: m-10 contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`
decision_point: the total table
enforcement_point: the validated one-transaction commit
exclusive_credential_holder: none
request_freeze_point: the consumed identity
authorization_linearization_point: `not specified` — recording follows the already-authorized effect
effect_linearization_point: the terminal commit
outcome_reporter: m-9
outcome_observer: none (self-reported by m-9)
outcome_validator: m-10 (the R36-F2 predicates — validation ≠ observation)
canonical_record: `tool_calls` terminals + the labeled evidence pair
bypass_paths: none on the wire path (one-way, fenced, validated); same-UID direct store write (§9)
failure_unknown_semantics: never-consumed/current×UNKNOWN ⇒ faults; terminals never downgraded
replay_idempotency: outcome-specific equivalence keys + persisted-epoch equality
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: m10-provider-attempt-recording
effect_class: durable_state_commit (the SEND effect itself is m-8's census row; admission ≠ authorization verbatim: `attempt_open_ok` is admission/row-first ordering — provider policy authorization is m-8's, immediately before credential attach/send)
requester: m-9 (`attempt_open`)
executor: m-10 (row commit + ack)
authority_source: turn admission + the §2a bounds
policy_owner: m-10 (recording); m-8/m-3 (the send/policy — theirs)
policy_artifact: contract §B.1  · exact artifact binding: m-10 contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`
decision_point: the attempt-open checks (the closed reject family)
enforcement_point: row-commit-before-ack (m-9 issues DATA-P only after)
exclusive_credential_holder: none at this row (m-8's row holds the credential facts)
request_freeze_point: the attempt row
authorization_linearization_point: `not specified` at this plane (m-8's authorize leg)
effect_linearization_point: the `provider_attempts` row commit — the row's OWN effect (m-8's authorize/attach/send linearize in THEIR rows)
outcome_reporter: TWO self-reports — m-8 (`attempt_result`) and m-9 (`attempt_stream_end`)
outcome_observer: none (self-reported) — the two-view reconciliation cross-checks two reporters; it is not neutral observation
outcome_validator: the §B.1 total consumption tables (incl. the cancellation table)
canonical_record: `provider_attempts` (E0 rows are conditional, m-9-authored, and inventoried at `m10-app-event-persist` — removed from this row per M10-S5-R3-F2)
bypass_paths: not applicable at this plane for the send (m-8's row); same-UID direct store write (§9)
failure_unknown_semantics: `UNKNOWN_PROVIDER_OUTCOME`/`PARTIAL_STREAM` parks; `REJECTED_LOCAL`; the `CANCELLED` two-cut honesty
replay_idempotency: observable-fact equivalence keys; one row per invocation
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: m10-app-event-persist
effect_class: durable_state_commit (persist-received; renamed from "carriage" per M10-S5-R3-F2 — m-9's conductor submit/accept is counterparty census territory, see the rationale below)
requester: m-9 (the `app_event` frame — the sole E0 populator, `reported_by` = worker)
executor: m-10 applier
authority_source: packet §3a + contract §B.1
policy_owner: m-3 (the schema)
policy_artifact: m-3 r4 @ `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`
decision_point: m-3 schema validation at receipt
enforcement_point: the persist commit
exclusive_credential_holder: the seat submit capability is m-7-broker-held — m-10 never submits
request_freeze_point: the persisted frame
authorization_linearization_point: `not specified` — no authorization leg exists (persisting a received report)
effect_linearization_point: **the LOCAL persist commit**
outcome_reporter: m-9 (the event content)
outcome_observer: none (self-reported)
outcome_validator: m-10 (schema conformance only)
canonical_record: `pending_app_events`
bypass_paths: same-UID direct store write (§9)
failure_unknown_semantics: the §11 no-E0 residuals at authorship; the carried-marking transition, acknowledgment, retry, and cross-owner dedup are `not specified` (§11 residual — no frozen contract defines them; no promise is made here)
replay_idempotency: `not specified` (no frozen dedup rule exists at this leg)
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: m10-child-terminate
effect_class: process_lifecycle (control send)
requester: supervision dispositions
executor: m-10 (`shutdown` → grace → SIGTERM → SIGKILL)
authority_source: the contract §B.1 lifecycle rules
policy_owner: m-10
policy_artifact: contract §B.1  · exact artifact binding: m-10 contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`
decision_point: the disposition classification (health/crash/retire/stop)
enforcement_point: the escalation-ladder execution (a distinct act from the decision — M10-S5-R3-F1)
exclusive_credential_holder: OS signal rights (same-UID inherent)
request_freeze_point: `not specified`
authorization_linearization_point: the disposition commit that orders termination
effect_linearization_point: OS termination, proven at reap
outcome_reporter: the OS
outcome_observer: the OS process table via `waitpid` — genuinely independent (reap = proven terminal)
outcome_validator: the reap requirement (TERMINATED only on reap)
canonical_record: `workers`/connector incarnation state
bypass_paths: same-UID `kill(2)` from any same-user process (accepted residual — an external kill presents as a crash and takes the crash disposition honestly); same-UID direct store write (§9)
failure_unknown_semantics: unresponsive ⇒ ladder escalation; the orphan bound honest (contract §B.3)
replay_idempotency: kills idempotent
threat_claim_scope: confusion-not-malice; same-UID residual accepted

---

effect_id: m9-f59-local-mutate
effect_class: `reversible_workspace_mutation`
requester: the model (a parsed tool call)
executor: the m-9 worker's tool executor
authority_source: the F59 one-shot ticket under the run-manifest 8-name allow-list (policy trivial in MVP; whole permission system Step-4)
policy_owner: m-10 (the run manifest §C + the F59 protocol §D)
policy_artifact: m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`
decision_point: m-10's `authorize_tool_call` + the `consume_ticket` atomic transaction
enforcement_point: the single central F59 dispatcher (§5.1), worker-side, gated on `consume_ok`
exclusive_credential_holder: `ambient` (local host authority; the ticket is authorization, not a credential)
request_freeze_point: identity #1 at §3.1/§5.2 request construction (SHA-256/JCS over the immutable args)
authorization_linearization_point: m-10's `consume_ticket` single-row UPDATE (§D.3)
effect_linearization_point: the invocation after `consume_ok` past derivation point B
outcome_reporter: the worker (self-report via `record_tool_outcome`)
outcome_observer: `none (self-reported)`
outcome_validator: m-10 (validates `invocation_identity` == stored triple before commit, §D.4)
canonical_record: m-10's `tool_calls` row (`EXECUTED`)
bypass_paths: `residual` (no ambient bypass for a file mutation beyond the tool itself; the effect is the governed write)
failure_unknown_semantics: CONSUMED-no-record ⇒ `UNKNOWN_TOOL_OUTCOME`; integrity fault ⇒ `NOT_INVOKED_INTEGRITY_FAULT` (definite no-effect); F59 reject ⇒ no ticket/no effect
replay_idempotency: F59 one-shot ticket, `UNIQUE(run_id,turn_id,tool_call_id)`; a fresh `tool_call_id` = a new ticket + D-4 semantic-duplicate disclosure
threat_claim_scope: `confusion-not-malice; workspace-mutation is governed + recorded`

---

effect_id: m9-f59-local-read
effect_class: `local_disclosure` (pulls host bytes into model context — a governance-relevant disclosure, not a mutation)
requester: the model (a parsed tool call)
executor: the m-9 worker's tool executor
authority_source: the F59 one-shot ticket under the run-manifest 8-name allow-list (policy trivial in MVP; whole permission system Step-4)
policy_owner: m-10 (the run manifest §C + the F59 protocol §D)
policy_artifact: m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`
decision_point: m-10's `authorize_tool_call` + the `consume_ticket` atomic transaction
enforcement_point: the single central F59 dispatcher (§5.1), worker-side, gated on `consume_ok` (NO read-only fast-path — grill-ratified)
exclusive_credential_holder: `ambient` (local host authority; the ticket is authorization, not a credential)
request_freeze_point: identity #1 at §3.1/§5.2 request construction (SHA-256/JCS over the immutable args)
authorization_linearization_point: m-10's `consume_ticket` single-row UPDATE (§D.3)
effect_linearization_point: the read invocation after `consume_ok` past derivation point B
outcome_reporter: the worker (self-report via `record_tool_outcome`)
outcome_observer: `none (self-reported)`
outcome_validator: m-10 (validates `invocation_identity` == stored triple before commit, §D.4)
canonical_record: m-10's `tool_calls` row (`EXECUTED`)
bypass_paths: `residual` (no ambient bypass for a governed read beyond the tool itself; the disclosure IS the governed read)
failure_unknown_semantics: CONSUMED-no-record ⇒ `UNKNOWN_TOOL_OUTCOME`; integrity fault ⇒ `NOT_INVOKED_INTEGRITY_FAULT` (definite no-effect); F59 reject ⇒ no ticket/no effect
replay_idempotency: F59 one-shot, `UNIQUE(run_id,turn_id,tool_call_id)`; a fresh `tool_call_id` = a new ticket (reads idempotent)
threat_claim_scope: `confusion-not-malice; read provenance recorded (which read, args-digest); MVP does NOT defend against untrusted-source-content injection — Step-4 honest-labeling-by-source-trust + adjudication`

---

effect_id: m9-f59-local-bash
effect_class: `process_lifecycle` (+ arbitrary same-UID host effect)
requester: the model (a parsed tool call)
executor: the m-9 worker's tool executor (spawns the process group, §1.4)
authority_source: the F59 one-shot ticket under the run-manifest 8-name allow-list (policy trivial in MVP)
policy_owner: m-10 (the run manifest §C + the F59 protocol §D)
policy_artifact: m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`
decision_point: m-10's `authorize_tool_call` + the `consume_ticket` atomic transaction
enforcement_point: the single central F59 dispatcher (§5.1), worker-side, gated on `consume_ok` (authorizes + invokes; does NOT contain host effect)
exclusive_credential_holder: `ambient` (host authority; no sandbox)
request_freeze_point: identity #1 at §3.1/§5.2 request construction; the args-digest covers the **command string**, cwd/env ambient-and-unpinned
authorization_linearization_point: m-10's `consume_ticket` single-row UPDATE (§D.3)
effect_linearization_point: the process-group spawn after `consume_ok` past derivation point B
outcome_reporter: the worker (self-report via `record_tool_outcome`)
outcome_observer: `none (self-reported)`
outcome_validator: m-10 (validates `invocation_identity` == stored triple; `executed` = invocation occurred, NOT effect-contained)
canonical_record: m-10's `tool_calls` row (`EXECUTED`)
bypass_paths: not specified
failure_unknown_semantics: a child not provably reaped ⇒ `UNKNOWN_TOOL_OUTCOME`/`PARTIAL_TOOL_EFFECT` (§1.4); CONSUMED-no-record ⇒ `UNKNOWN_TOOL_OUTCOME`; integrity fault ⇒ `NOT_INVOKED_INTEGRITY_FAULT`
replay_idempotency: F59 one-shot, `UNIQUE(run_id,turn_id,tool_call_id)` (bash effects are not semantically idempotent — a fresh `tool_call_id` re-runs)
threat_claim_scope: `confusion-not-malice; ambient host authority; same-UID + escaped-setsid-descendant residual ACCEPTED; the record attests authorization+invocation, NOT host effect`

---

effect_id: m9-relay-submit-append
effect_class: `durable_state_commit` (a governed store append, via the courier)
requester: the model
executor: the worker (issues the verb) + the conductor (performs the append)
authority_source: TWO layers — the worker-side F59 ticket + the conductor's own governance (accepted/rejected/held + seat-stamp)
policy_owner: m-9 (the native tool) + m-1/m-7 (the courier governance)
policy_artifact: m-9 lifecycle half `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd` (r21; worker r7 self-cites r19 `2a96a07bb2f2606b7b42fe34270beaa209ca08cfd2c19f6e91f44eb18eef734c`, revision-neutral per F92) + m-7 r11 + m-2  · exact artifact binding: m-2 `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`; m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`
decision_point: worker F59 `consume` AND the conductor's admission
enforcement_point: worker F59 dispatcher + the conductor's single-threaded serialized commit loop
exclusive_credential_holder: the broker-held S-B seat credential; the worker holds only the epoch-bound USE capability (m-1, F60/F66)
request_freeze_point: the relay payload after m-2 `SubmitPayloadFromArguments`, before F59 authorize
authorization_linearization_point: worker `consume_ok`, then the conductor's commit-loop admission
effect_linearization_point: the conductor's store append (`accepted`)
outcome_reporter: the worker (`record_tool_outcome`) + the conductor (its own store record)
outcome_observer: the conductor (an independent observer of the append — distinct from the worker's self-report)
outcome_validator: m-10 (worker-side record) + the conductor (the relay's form/lineage validation)
canonical_record: the conductor's append-only store record + m-10's `tool_calls` row
bypass_paths: `residual` — the worker reaches the store ONLY through the closed verb surface (m-7 guardrail; the kimi-`klient` no-escape-hatch convergence, §4.1); no raw store path
failure_unknown_semantics: F59 reject (worker) / conductor `rejected`/`held` (courier)
replay_idempotency: F59 one-shot (worker) + the conductor's replay-by-content-hash (courier; the AXI idempotent-ack, §4.3)
threat_claim_scope: `confusion-not-malice; double-governed (worker authority + courier content-governance)`

---

effect_id: m9-relay-project-read
effect_class: `courier_read/serve` (NOT a store append — its own read/serve row, VP F87)
requester: the model
executor: the worker (issues) + the conductor (serves the projection)
authority_source: the worker-side F59 ticket (uniform) + the conductor's read/serve governance
policy_owner: m-9 (native tool) + m-1/m-7 (courier)
policy_artifact: m-9 lifecycle half `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd` (r21; worker r7 self-cites r19 `2a96a07bb2f2606b7b42fe34270beaa209ca08cfd2c19f6e91f44eb18eef734c`, revision-neutral per F92) + m-7 r11  · exact artifact binding: m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`
decision_point: worker F59 `consume` + the conductor's serve path
enforcement_point: the F59 gate (worker) + the conductor serve path
exclusive_credential_holder: broker-held S-B (worker holds the USE capability)
request_freeze_point: the query args after m-2 `ProjectSchema()` validation, before F59 authorize
authorization_linearization_point: worker `consume_ok`
effect_linearization_point: the conductor returns the projection (a read, no store mutation; **project = inbox-only scope**, m-1)
outcome_reporter: the worker (`record_tool_outcome`)
outcome_observer: the conductor (serves the read)
outcome_validator: m-10 (worker record)
canonical_record: m-10's `tool_calls` row
bypass_paths: `residual` (closed verb surface; inbox-only scope)
failure_unknown_semantics: F59 reject (worker) / conductor serve error
replay_idempotency: F59 one-shot (reads idempotent)
threat_claim_scope: `confusion-not-malice; disclosure of inbox-scoped courier state, recorded`

---

effect_id: m9-relay-read-read
effect_class: `courier_read/serve` (its own read/serve row, VP F87)
requester: the model
executor: the worker (issues) + the conductor (serves the read)
authority_source: the worker-side F59 ticket + the conductor's read/serve governance
policy_owner: m-9 + m-1/m-7
policy_artifact: m-9 lifecycle half `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd` (r21; worker r7 self-cites r19 `2a96a07bb2f2606b7b42fe34270beaa209ca08cfd2c19f6e91f44eb18eef734c`, revision-neutral per F92) + m-7 r11  · exact artifact binding: m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`
decision_point: worker F59 `consume` + the conductor serve path
enforcement_point: the F59 gate (worker) + the conductor serve path
exclusive_credential_holder: broker-held S-B (worker USE capability)
request_freeze_point: the read args after m-2 `ReadSchema()` validation, before F59 authorize
authorization_linearization_point: worker `consume_ok`
effect_linearization_point: the conductor returns the record(s) (a read, no store mutation)
outcome_reporter: the worker (`record_tool_outcome`)
outcome_observer: the conductor (serves the read)
outcome_validator: m-10 (worker record)
canonical_record: m-10's `tool_calls` row
bypass_paths: `residual` (closed verb surface)
failure_unknown_semantics: F59 reject (worker) / conductor serve error
replay_idempotency: F59 one-shot (reads idempotent)
threat_claim_scope: `confusion-not-malice; disclosure of courier state, recorded`

---

effect_id: m9-describe-serve
effect_class: `metadata_serve` (the m-2 rediscovery/tool-schema metadata op the native tool uses for re-render, §4.2)
requester: the worker (consumer logic — the re-render loop, not a model tool call)
executor: the conductor (m-7 `Describe` transport)
authority_source: the closed capability surface (m-7 §2.8; `Describe` is on the worker surface)
policy_owner: m-7 (the caller seam + `Describe` transport) + m-2 (the form/schema metadata)
policy_artifact: m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572` + m-2 `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`
decision_point: m-7's `Describe` caller seam (capability-fenced)
enforcement_point: m-7's transport, per-operation capability fence (F64)
exclusive_credential_holder: the broker-held S-B seat credential (the worker holds only the epoch-bound USE capability)
request_freeze_point: the `Describe` request keyed by `DeclaredPhaseTier(args)`, at send
authorization_linearization_point: m-7's per-operation capability fence at `Describe` time
effect_linearization_point: m-7 returns the `DescriptionResponse` (a metadata read; no store mutation)
outcome_reporter: the worker (uses the result in the re-render loop)
outcome_observer: `none (metadata read)`
outcome_validator: the m-2 module (schema validity of the rendered form)
canonical_record: `not specified` (a metadata read; not a governed store effect — the re-render decision is worker-local)
bypass_paths: `residual` (closed capability surface; no raw metadata path)
failure_unknown_semantics: an unknown/malformed description ⇒ fail-closed, no submit (m-2 §2.3.2); a transport fault ⇒ re-invoke (idempotent read)
replay_idempotency: idempotent read
threat_claim_scope: `confusion-not-malice; metadata read, no external effect`

---

effect_id: m9-provider-attempt-send
effect_class: `external_send` (to the provider, via the m-8 connector)
requester: the worker
executor: the m-8 connector (holds the credential, does the wire)
authority_source: **m-8's app-side provider-policy authorization, immediately before credential attach/send** (NOT `attempt_open_ok`, which is admission — VP F87)
policy_owner: m-8 (the provider contract)
policy_artifact: m-8 r12 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51` + m-3 (egress policy)  · exact artifact binding: m-3 r4 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`
decision_point: m-8's `freeze→authorize→attach→send`; m-3 egress policy (denied ⇒ `egress_denied`)
enforcement_point: m-8 connector (last pre-wire enforcement, app-side) + the worker's row-first `attempt_open`→`attempt_open_ok` ordering (ADMISSION, not authorization)
exclusive_credential_holder: **m-8 holds the provider credential (S-A); the worker holds NONE** (§1.1)
request_freeze_point: the `LLMRequest` assembled+validated before DATA-P (m-8 frozen-core re-check)
authorization_linearization_point: m-8's provider-policy authorization immediately before attach/send
effect_linearization_point: the DATA-P send (the wire crossing)
outcome_reporter: m-8 (`attempt_result` on CTRL-C)
outcome_observer: the worker (observes the stream terminal, emits `attempt_stream_end`) — an independent view; m-10 reconciles the two views
outcome_validator: m-10 (two-view reconciliation → the `provider_attempts` terminal)
canonical_record: m-10's `provider_attempts` row
bypass_paths: `residual` (no provider credential worker-side; cannot send except through m-8)
failure_unknown_semantics: the three loss facts (`stream_failed`/`stream_lost`/no-stream); a bare closure/crash ⇒ `UNKNOWN_PROVIDER_OUTCOME`, never `CANCELLED`
replay_idempotency: full input every attempt (`store:false`); a retry mints a new `attempt_id`
threat_claim_scope: `confusion-not-malice; provider egress governed app-side by m-8 + m-3; worker holds no credential`

---

effect_id: m9-attempt-lifecycle
effect_class: `admission/observation record` (NOT the provider send — that is E8)
requester: the worker (the turn machine, at ATTEMPTING)
executor: the worker (sends the CTRL-W frame)
authority_source: the active turn lease (m-10) at the current epoch
policy_owner: m-10 (attempt admission + the row-first ordering)
policy_artifact: m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146` §B.1
decision_point: m-10 admits (`attempt_open_ok` after the `provider_attempts` row commits — ADMISSION/row-first, VP F87) or rejects (`attempt_open_reject{stale_epoch|invalid_turn|invalid_lease}`)
enforcement_point: m-10's row-first commit — the durable row precedes the ack; the worker issues DATA-P ONLY after `attempt_open_ok`
exclusive_credential_holder: `not applicable` (a CTRL-W control frame)
request_freeze_point: the frame at send
authorization_linearization_point: m-10's `provider_attempts` row commit (admission becomes fact)
effect_linearization_point: the same row commit for `attempt_open`; the stream-view disposition write for `attempt_stream_end`
outcome_reporter: the worker (the stream-view disposition)
outcome_observer: m-10 (admits + reconciles the two views against m-8's CTRL-C view)
outcome_validator: m-10 (two-view reconciliation)
canonical_record: m-10's `provider_attempts` row
bypass_paths: `residual` (no path to a provider attempt except through admission)
failure_unknown_semantics: `invalid_turn`/`invalid_lease` ⇒ worker ceases + awaits supervision, no local retry; `stale_epoch` ⇒ fenced-cease
replay_idempotency: one row per `attempt_id`
threat_claim_scope: `confusion-not-malice; admission ordering guarantees the durable row precedes the wire`

---

effect_id: m9-turn-terminal
effect_class: `durable_state_commit` (the turn record + lease release)
requester: the worker (the turn machine, at TERMINAL)
executor: the worker (sends the CTRL-W frame)
authority_source: the turn machine (lifecycle §2.3/§2.9) under the active-turn lease
policy_owner: m-10 (the turn record + lease + the D-5 consumption)
policy_artifact: m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146` §B.2
decision_point: m-10 accepts + commits the terminal (or `turn_reject{stale_epoch|unknown_turn|conflicting_report}`)
enforcement_point: m-10's durable commit → `turn_receipt`; the composition rule (a `turn_cancelled` terminal requires the prior `turn_cancel_ack` committed)
exclusive_credential_holder: `not applicable` (a CTRL-W control frame)
request_freeze_point: the frame at send
authorization_linearization_point: m-10's durable commit of the terminal fact
effect_linearization_point: the lease release (only `turn_terminal{turn_cancelled}` releases; `turn_cancel_ack` records the partial, does not terminalize)
outcome_reporter: the worker
outcome_observer: m-10
outcome_validator: m-10 (composition rule)
canonical_record: m-10's `turns` record
bypass_paths: `residual`
failure_unknown_semantics: a lost reply ⇒ re-send (m-10 idempotent, same receipt); a crash between the worker's terminal and m-10's record ⇒ the turn stays `INTERRUPTED`, never silently completed; `conflicting_report` ⇒ internal fault
replay_idempotency: equivalence-keyed (`{terminal}` for D-5)
threat_claim_scope: `confusion-not-malice`

---

effect_id: m9-cancel-attempt
effect_class: `cancellation_intent`
requester: the worker (operator/steer cancel reaches the turn machine)
executor: the worker (sends the DATA-P `cancel_attempt`)
authority_source: proven current-epoch intent (m-8 §1.4)
policy_owner: m-8 (the sole cancellation carrier + the two honest cuts)
policy_artifact: m-8 r12 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51` §1.4
decision_point: m-8 (chooses the cut by whether transport was invoked)
enforcement_point: m-8 (`cancelled(<cancel_point>)` on CTRL-C; the worker never reads a closed channel as a cancel)
exclusive_credential_holder: `not applicable` (a DATA-P control frame; no credential)
request_freeze_point: the `cancel_attempt` frame at send
authorization_linearization_point: m-8's determination of the cut (the cancel becomes fact at `cancelled(<cancel_point>)`)
effect_linearization_point: m-8's `cancelled(<cancel_point>)` terminal on the `provider_attempts` row
outcome_reporter: m-8 (the CTRL-C `cancelled` disposition)
outcome_observer: the worker (via the stream / `attempt_stream_end`)
outcome_validator: m-10 (commits the `CANCELLED` row)
canonical_record: m-10's `CANCELLED` `provider_attempts` row
bypass_paths: `residual`
failure_unknown_semantics: a bare closure/crash ⇒ `UNKNOWN_PROVIDER_OUTCOME`, never `CANCELLED` (only proven current-epoch intent commits a cancel)
replay_idempotency: m-10 equivalence over `{attempt_id, reported turn_epoch, cancel_point}`; the worker neither carries nor compares `cancellation_id`
threat_claim_scope: `confusion-not-malice; interrupted ≠ failed`

---

effect_id: m9-e0-app-event
effect_class: `self_reported_telemetry` (E0 into `pending_app_events`)
requester: the worker (the turn machine, per the §6.1 total disposition→`phase` mapping)
executor: the worker (populator + carrier; m-8 emits none)
authority_source: the seat (the worker carries; the event is `self_reported`)
policy_owner: m-3 (the `m3.app_event.v1` schema + redaction + the `phase` token set)
policy_artifact: m-3 r4 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`
decision_point: the §6.1 total disposition→`phase` mapping (m-9's forward mapping) against the m-3 closed schema
enforcement_point: the m-3 redaction + schema validation at construction (fail-closed serve prerequisite)
exclusive_credential_holder: `none (self-reported)` (`reported_by` is a claim)
request_freeze_point: the event constructed under `m3.app_event.v1` + redaction, before the frame
authorization_linearization_point: `not applicable` (a carried event, not an authorized effect)
effect_linearization_point: the `app_event` frame on CTRL-W → the `pending_app_events` row
outcome_reporter: the worker
outcome_observer: `none (self-reported)`
outcome_validator: m-3 schema (validity) / m-10 (row commit)
canonical_record: m-10's `pending_app_events` row
bypass_paths: `residual` (redaction is fail-closed; no raw provider/prompt bytes can ride)
failure_unknown_semantics: not specified
replay_idempotency: E0 events are not digested artifacts; idempotency rides m-10's row carry
threat_claim_scope: `confusion-not-malice; E0 = self_reported, a claim not a proof`

---

effect_id: m9-process-lifecycle
effect_class: `process_lifecycle`
requester: m-10 (spawn/supervise)
executor: the worker (its own lifecycle + `setpgid`/`killpg` child teardown, §1.4)
authority_source: m-10 supervision (the worker is a replaceable generation)
policy_owner: m-10 (worker lifecycle + the lease PID-reap gate)
policy_artifact: m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146` §B.2/§B.3
decision_point: m-10 (spawn, retire, reap)
enforcement_point: m-10's lease (gates on the worker-generation PID being reaped, NOT every tool-tree process)
exclusive_credential_holder: `not applicable`
request_freeze_point: `not applicable` (a supervision lifecycle, not a frozen request)
authorization_linearization_point: m-10's generation/lease state transition (spawn admitted / retirement committed)
effect_linearization_point: the worker's exit; the group `killpg` for tool children
outcome_reporter: the worker (until EOF) / m-10 (durable)
outcome_observer: m-10
outcome_validator: m-10
canonical_record: m-10's generation/lease state
bypass_paths: not specified
failure_unknown_semantics: a child not provably reaped ⇒ the tool call parks `UNKNOWN`; a detached descendant is NOT interlocked against a successor
replay_idempotency: `not applicable`
threat_claim_scope: `confusion-not-malice; escaped-descendant residual accepted`

---

effect_id: m9-wake-forward
effect_class: not specified
requester: the worker (on receiving a push, or at a rediscovery pass)
executor: the worker (forwards `wake_forward{relay_id}`) + m-10 (inserts the schedule)
authority_source: the worker's forwarding obligation (the closed surface, m-7 §2.8) + m-10's scheduler
policy_owner: m-10 (the wake scheduler + `UNIQUE(relay_id)` at-most-once)
policy_artifact: m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146` §E
decision_point: m-10's scheduler `INSERT OR IGNORE INTO wake_schedule(relay_id, …)` (idempotent insert; m-10 r40 §E)
enforcement_point: that same `INSERT OR IGNORE` transaction + the durable `UNIQUE(relay_id)` constraint — the point that controls the declared `wake_schedule` insertion (m-10 r40 §E); *(the later `pending`→`dispatched` + `turns`-row commit is a DISTINCT downstream turn-admission effect — m-10's own `m10-wake-acceptance` census row, not this insertion)*
exclusive_credential_holder: `not applicable` (a CTRL-W signal frame)
request_freeze_point: the `wake_forward{relay_id}` frame at send
authorization_linearization_point: `not applicable` (advisory; the schedule insert is admission, not authorization)
effect_linearization_point: the committed `wake_schedule` row (`UNIQUE(relay_id)`; one schedule per `relay_id`)
outcome_reporter: the worker (forwards)
outcome_observer: m-10 (schedules)
outcome_validator: m-10 (dedup)
canonical_record: m-10's `wake_schedule` row
bypass_paths: `residual` (the worker NEVER touches the conductor to schedule — it only forwards; the scheduler never touches the conductor either, m-10 M10-S5-R1-F3)
failure_unknown_semantics: not specified
replay_idempotency: idempotent — a duplicate forward yields one schedule (m-10 `UNIQUE(relay_id)`)
threat_claim_scope: `confusion-not-malice; durability guaranteed by rediscovery, not by the signal`

---

effect_id: m9-push-receive
effect_class: not specified
requester: the conductor (a relay arrival to the seat channel triggers a push)
executor: the broker (forwards the nudge to the current-epoch capability connection) + the worker (receives it)
authority_source: the closed capability surface (push is on the worker surface, m-7 §2.8; fenced per-push at forward time, m-1 §2.3/F64)
policy_owner: m-7 (push transport) + m-1 (the per-operation fence)
policy_artifact: m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572` + m-1 `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`
decision_point: the broker's forward-to-current-epoch-connection (an epoch-fenced forward)
enforcement_point: the F64 per-push capability fence at forward time
exclusive_credential_holder: the broker-held S-B seat credential (the worker holds the epoch-bound USE capability)
request_freeze_point: `not applicable` (an inbound signal the worker does not author)
authorization_linearization_point: `not applicable` (advisory receipt, not an authorized effect)
effect_linearization_point: the worker's in-memory receipt of the nudge (it triggers an E14 `wake_forward` and/or a rediscovery pass — no durable effect of its own)
outcome_reporter: `none (an inbound signal; the worker's response is E14 / rediscovery)`
outcome_observer: the worker (receives)
outcome_validator: `not applicable` (advisory; the record is truth via rediscovery, not the push)
canonical_record: `not specified` (a push leaves no durable row of its own; the projectable record it signals is the conductor's store truth)
bypass_paths: `residual`
failure_unknown_semantics: not specified
replay_idempotency: idempotent — a duplicate push causes at most one forward/rediscovery; the record is truth
threat_claim_scope: `confusion-not-malice; push is advisory, never authoritative state`

---

effect_id: m9-objective-acquire
effect_class: `local_disclosure` (pulls the turn's objective/task into the pinned Tier-0 context — a governance-relevant disclosure, not a mutation)
requester: the worker (turn-assembly consumer logic at turn start — NOT a model tool call)
executor: the worker (reads `turn_open.admission_ref`; for `{kind: wake_relay, relay_id}` issues its own seat `read(relay_id)`; for `{kind: operator_input, task_input}` reads the verbatim `task_input` off the frame)
authority_source: for `wake_relay` the seat USE capability (the ref **LOCATES** the task, grants nothing — I-PH intact, m-1); for `operator_input` the `turn_open` frame carriage (app-originated operator-authored content)
policy_owner: m-10 (the `turn_open.admission_ref` carrier + admission commit) + m-7/m-1 (the seat read surface)
policy_artifact: m-10 r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146` R37 + m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572` + m-1 `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`
decision_point: m-10's admission commit (chooses the kind + writes the `admission_ref` value to the `turns` row)
enforcement_point: for `wake_relay` the broker per-operation capability fence on the seat `read` (m-7 §2.8/F64); for `operator_input` m-10's pre-commit `FRAME_MAX` sizing gate (`admission_refused{task_input_frame_overflow}`, before any frame exists — the worker never sees an oversized objective)
exclusive_credential_holder: the broker-held S-B seat credential for the `wake_relay` read (the worker holds the USE capability); `not applicable` for the on-frame `operator_input` read
request_freeze_point: the `admission_ref` value frozen in m-10's admission commit (the durable BINDING linearization — admission, not authorization; `turn_open` emits post-commit byte-identically — R38-F1)
authorization_linearization_point: **branch-faithful (admission ≠ authorization — the ref LOCATES, grants nothing; the m-10 admission commit binds the ref but transfers NO authority, m-10 holding no seat credential):** for `{kind: wake_relay}` the **broker's per-operation seat-capability fence at the moment the seat `read` is served** (m-7 §2.8/F64) — where the read's authorization becomes fact; for `{kind: operator_input}` `not applicable` (the `task_input` is inert app-originated content already carried on the **admitted** `turn_open` frame over authenticated CTRL-W — m-10 admits the turn and emits the frame post-commit; there is NO authorization event for `turn_open`, so reading a frame member is not a separately-authorized effect)
effect_linearization_point: the worker's in-memory capture of the objective into pinned turn state (for `wake_relay`, after the seat `read` returns)
outcome_reporter: the worker (uses the objective in Tier-0 assembly, §7.1)
outcome_observer: `none (turn-assembly consumer read)`
outcome_validator: `not applicable` (the objective enters ephemeral pinned turn state; a model-issued `relay.read` is instead the F59-gated E6 — this row is worker-internal turn assembly, like E7 `Describe`)
canonical_record: `not specified` (the objective is ephemeral pinned turn-state; the durable `wake_relay` task relay is the conductor's own record authored by the task's sender, not by this read)
bypass_paths: `residual` (the ref locates, never grants — no authority transfers; the `wake_relay` fetch runs ONLY through the worker's own seat capability, and `operator_input` is inert content on a frame)
failure_unknown_semantics: a missing/malformed `admission_ref` is a protocol fault ⇒ fail-closed (the member is REQUIRED, never absent — m-10 r40 R37); a `wake_relay` read failure ⇒ the worker's §1.6 broker-error disposition / rediscovery; a replacement re-carries the byte-identical ref ⇒ the same objective re-materializes (no second task identity)
replay_idempotency: idempotent — the ref is stable and re-admission re-carries it byte-identically (R38-F1); the objective re-materializes identically
threat_claim_scope: `confusion-not-malice; the ref LOCATES-never-grants (I-PH intact); operator_input is app-originated operator-authored content (not provider payload, not secret — the m-1 NOT-secret census)`

---

## §D. Non-effect appendix (OUTSIDE the effect set — rationales, not rows)

- **merge** — NOT an MVP governed effect. Merge onto `frank/` main is a protocol-level operator grant (a MERGE-GATE authorization + the operator's countersign), not a runtime effect the conductor or control plane mediates.
- **release-binding** — NOT a runtime effect row. The F63 post-build event binds the app-main/m-9/m-8 artifact digests (or a `release_digest`) before live E3; a build-pipeline act sequenced after the stage-6 lock, governed by the release-binding contract.
- **deploy** — ABSENT from the MVP scope.
- **`-mint` seat mint (current live path)** — a live Step-2 code path that mints before `store.AcquireRoot` (defect H-26, `master/FRANK-HARDENING-BACKLOG.md`); named as a known gap, not folded as a governed row until the H-26 fix lands.
