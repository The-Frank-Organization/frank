## Team m-1 — Trust & Identity — INDEPENDENT IMPLEMENTER AUDIT

ROLE: Implementer
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c1-audit-m-1
PARENT_DISPATCH_ID: c1-audit-m-1
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — read-only audit; operator-judgment items surfaced below
FROM: m-1.implementer
TO: m-1.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-1-trust-identity
OWNER: m-1 (Trust & Identity)
IN_REPLY_TO: c1-audit-m-1/AUDIT-orchestrator-planner-20260628-125953.md

This is the independent Implementer audit for c1-audit-m-1. Method: direct source audit of the agentic dev team protocol, jcode, claude-code, export design intent, the live relay convention, three read-only explorer subagents, and a narrow web prior-art sweep over authenticated buses and transparency logs. No implementation or prototype work was performed.

## 1. 4-bucket verdict

PRIMARY_BUCKET: still-open

still-open: the frank primitive is not present in the agentic dev team protocol, jcode, or claude-code: no system owns the relay store as sole writer while stamping trusted `FROM` from the channel/connection that delivered the relay. The protocol explicitly says `PARENT_DISPATCH_ID` is agent-authored and only confusion-robust, with router-derived parent as the forgery-robust future path (`extracted/.../orchestrator-planner/protocol.md:88`). Its ROLE/FROM rule is expressly "not a harness identity mechanism" and does not defend deliberate forgery (`.../protocol.md:135-138`; `.../tools/relay-lint.py:555-575`). Claude-code stores a caller-supplied `from` in the mailbox record (`references/claude-code/src/utils/teammateMailbox.ts:43-50,134-180`). jcode has a server-side session bus, but its comm wire records still carry `from_session`/`session_id` in request bodies (`references/jcode/crates/jcode-protocol/src/wire.rs:376-390,459-466`) and its persistence is mutable swarm JSON, not an append-only stamped relay store (`references/jcode/crates/jcode-app-core/src/server/swarm_persistence.rs:216-242`).

already-closed: reuse candidates exist, but none closes the primitive. Reuse the agentic dev team protocol's human-readable relay layout and truth-agnostic lineage checks (`.../protocol.md:378-386`; `.../tools/relay-lint.py:1-8,1200-1360`), claude-code's per-recipient inbox projection and lock/poll delivery pattern (`references/claude-code/src/utils/teammateMailbox.ts:1-7,141-180`; `references/claude-code/src/hooks/useInboxPoller.ts:118-125,147-154`), and jcode's broadcast/session bus concepts (`references/jcode/crates/jcode-base/src/bus.rs:356-490`) only as transport prior art.

product-overlapped: m-1 must not absorb m-2's form-schema internals, m-3's observe predicate design, m-4's routing policy, or m-6's human-surface UX. m-1 owns the trust anchor they consume: stamped identity, append-only store, isolation, and the pre-delivery gate placement.

recommended-next: proceed to DESIGN for a minimal `submit/project/read/mint_seat` primitive with two hard invariants: only the conductor can write the store, and each inbound write channel maps to exactly one seat. Treat Merkle/transparency hardening as later optional tamper-evidence, not required for Step-1 seat-forgery robustness.

## 2. Three self-asserted identity gaps

Gap A — agentic dev team protocol agent-authored `FROM`/lineage:
The current protocol instructs relays to carry author-filled `FROM`, `TO`, and `CC` header fields (`.../protocol.md:13-25,28-40`). `relay-lint.py` parses and validates these fields structurally, but states it does not verify truth (`.../tools/relay-lint.py:1-8`) and the adapter docs explicitly defer forged-FROM identity because relay-lint cannot prove which process wrote a relay (`.../tools/adapters/README.md:52-54`). Channel-stamped `FROM` closes this by removing `FROM` from the lane-authored payload: the store accepts only a conductor-stamped record whose author is derived from the seat connection.

Gap B — jcode role/session self-assertion surface:
jcode's `swarm` tool exposes `assign_role` as an agent-facing action (`references/jcode/crates/jcode-app-core/src/tool/communicate.rs:607-652`) and the handler writes the submitted role into member state (`references/jcode/crates/jcode-app-core/src/server/comm_control.rs:669-730,766-770`). Its comm messages carry `from_session` as request data (`references/jcode/crates/jcode-protocol/src/wire.rs:376-390`) and handler code uses the supplied session id to determine swarm identity and routing (`references/jcode/crates/jcode-app-core/src/server/client_comm_message.rs:103-120,147-180`). Nuance: this is stronger than claude-code's arbitrary mailbox `from` because the server normally has a session context; however, it still lacks the frank invariant that the durable relay store is written only by a courier that stamps identity from the connection and never from the request body. Channel-stamped `FROM` closes the durable-record gap by making the seat binding a conductor input, not a request field or mutable role value.

Gap C — claude-code caller-supplied `from` plus literal team-lead guard:
Claude-code's teammate mailbox stores `from: string` (`references/claude-code/src/utils/teammateMailbox.ts:43-50`), and `writeToMailbox()` spreads the caller-supplied message into the inbox record (`references/claude-code/src/utils/teammateMailbox.ts:134-180`). Critical approval handling then tests the stored string (`msg.from === 'team-lead'`) to exit plan mode or accept mode changes (`references/claude-code/src/hooks/useInboxPoller.ts:156-164,190-194`). Channel stamping closes this directly: agents submit to the courier; only the courier writes projected inboxes; a non-lead channel cannot produce a stored `from` of `team-lead`.

## 3. Store and identity recommendation

Minimal API:
- `mint_seat(seat_address, runtime_binding) -> seat_connection`: conductor-only. Creates the channel/credential that binds one runtime lane to one seat.
- `submit(seat_connection, typed_payload) -> relay_id`: only write path. The conductor derives `FROM` and `ROLE` from the connection, fills system fields, runs m-2 form validation and m-3 observe-as-send hooks before append, writes the relay record, then projects delivery to addressed inboxes.
- `project(seat_connection) -> inbox_view`: read path scoped by the caller's stamped identity and `TO`/`CC`.
- `read/query(relay_id | lineage_selector) -> relay_record`: audit/lineage read surface for lint, reconciliation, m-3 observation, m-4 routing records, and m-6 projection.

On-disk shape:
Keep the protocol's human-readable shape for Step 1: `master/relays/<DISPATCH_ID>/<PHASE>-<ROLE>-<timestamp>.md` plus `INDEX.md`, but make it a projection of a conductor-owned record rather than a lane-authored file. `INDEX.md` can remain append-only in the convention sense (`.../protocol.md:386`), but the authoritative invariant is "single conductor writer", not "agents promise append-only".

Identity stamp:
`FROM`, `ROLE`, `PARENT_DISPATCH_ID`, timestamp, and future computed result fields are system-owned. `FROM := seat_for_connection(connection_id)`. Do not support a separate payload/display sender field in frank. That would recreate the email-style two-FROM trap where one field is visible and another is trusted.

Isolation:
I1: store write isolation. Seat runtimes cannot write relay files or `INDEX.md` directly; they can only call `submit()`.
I2: channel isolation. A seat cannot acquire another seat's connection credential. Step 1 can use a conductor-minted per-seat token or stdio/socket binding over an isolated runtime; standalone runtime can later strengthen the same interface with OS peer credentials or mTLS/SPIFFE-style workload identity. The interface should not change when the attestation backend strengthens.

Web prior-art check:
NATS authorization demonstrates a bus-side model where server policy controls which authenticated users may publish/subscribe to subjects, supporting the "gate at transport boundary" pattern (official NATS docs, Authorization, lines 82-128). Sigstore/Rekor demonstrates transparency-log hardening as a separate auditability layer, not the identity primitive itself (official Sigstore Rekor docs). Matrix room events show durable event systems can carry sender fields, but that does not by itself solve whether the sender was stamped from a trusted channel. The frank novelty is applying the known authenticated-channel pattern to the agent-seat relay store.

## 4. Boundary contract

Writes: conductor-owned relay records with stamped `FROM`, `ROLE`, lineage/system fields, and append-only projection into `master/relays/<DISPATCH_ID>/...` plus `INDEX.md`.

Reads: the agentic dev team protocol's relay shapes and linter lineage semantics; jcode session/bus and SafetySystem patterns; claude-code inbox projection; export design intent (the pre-build design-state export (not vendored), adaptive-routing pillar doc, lines 31-45).

Target entity: the frank relay store + seat-stamper primitive, still design-only.

Downstream consumers:
- m-2: consumes the stamped system-field set and must mark those fields `system`/not lane-fillable.
- m-3: consumes the pre-append `submit()` hook and store isolation to run observe-as-send from outside the lane.
- m-4: writes routing records into the stamped store; policy decisions key authority to stamped seat identity, not model claims.
- m-6: projects `TO`/`CC` into human/operator inboxes and email buckets; it must never become a second writer of relay identity.

Contract:
Fields exposed are `relay_id`, `dispatch_id`, `parent_id`, `from_seat`, `role`, `to[]`, `cc[]`, `phase`, `authority`, `created_at`, `body`, `validation_result`, and future observe evidence references. Store write operations are allowed only through `submit()`. Projection/read operations may be scoped, but cannot mutate identity.

Proof target for DESIGN/Step 1:
E2 fixture where a lane submits payload claiming another `FROM`; stored relay still has the channel seat, and no lane direct-write path can create a valid delivered relay.

No-consumer action: none. Every proposed primitive surface has a named downstream consumer.

## 5. Answers to m-1.planner questions

Q1: Concur with I1+I2 sufficiency. If only the conductor writes and only the true seat has its channel, `FROM := lookup(connection)` is enough for seat-forgery robustness. Merkle/transparency is separate defense against a malicious or corrupted conductor, not necessary for the stated Step-1 seat-forgery property.

Q2: `submit/project/read/mint_seat` is the right minimal surface. I would not add a public `attest()` call; attestation should be an internal part of `mint_seat`/connection setup and `submit()` should fail if the connection is not bound. A public `attest()` risks becoming another lane-asserted claim.

Q3: Forgery-robustness ladder: operator-attested manual relay is confusion-resistant but not by-construction against a malicious lane; minted token over an isolated per-seat connection is by-construction if tokens are unshared and lanes cannot read each other's credentials; OS peer identity or mTLS/SPIFFE-style workload identity is the stronger standalone form. Step 1 should pick minted isolated channels if the claim is "forgery-robust by construction" rather than "operator-trusted".

Q4: Reuse credit is mostly right. I would under-credit jcode as a direct identity model: its session bus is useful transport prior art, but because request bodies still carry sender/session fields and role state is mutable through `assign_role`, it is not a near-miss sole-writer identity primitive. Claude-code is the best delivery/projection prior art; the agentic dev team protocol is the best governance/lineage prior art.

## 6. Operator-judgment items

1. Step-1 transport strength: decide whether Step-1 may ship as operator-attested automation or must include minted per-seat channels before claiming by-construction seat-forgery robustness.
2. Store tamper-evidence: accept "trust the conductor as TCB" for Step-1, or require transparency/Merkle evidence in Step 1. Recommendation: defer tamper-evidence.
3. Identity/authority split: confirm m-1 owns "who wrote this relay" and the store write path; m-4/m-5 own "what this stamped seat may do".

## 7. Evidence levels

Source and design claims above are E1: file:line citations, official docs, or static search. No E2/E3/E4 claim is made because this phase has no implementation.

ACTIONS_GIT_REF: wrote expected AUDIT relay report only; no source, test, branch, commit, PR, or `pcode/` changes. cwd git status unavailable because workspace root is not a git repo.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the future code repo)
RELAY_LINT: OK — single-file lint clean. Relay-root mode also reports this relay OK; the only relay-root errors are the known lint-exempt `master/relays/INDEX.md` header errors.
