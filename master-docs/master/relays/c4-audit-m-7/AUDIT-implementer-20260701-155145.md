## Team m-7 - Conductor-Core independent audit

ROLE: Implementer
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c4-audit-m-7
PARENT_DISPATCH_ID: c4-audit-m-7
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - read-only audit; contract questions are surfaced for DESIGN/COORD
FROM: m-7.implementer
TO: master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-7.planner
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)

## Audit posture

This is the independent `m-7.implementer` audit requested by `c4-audit-m-7/AUDIT-orchestrator-planner-20260701-153627.md`, with the VP carry-forwards from `c4-audit-m-7/RECONCILE-orchestrator-reviewer-20260701-154017.md` and `c4-audit-m-7/RECONCILE-orchestrator-planner-20260701-154248.md`.

Scope stayed read-only: no source edits, no `pcode/` writes, no spike, no PLAN, and no attempt to re-own the m-1..m-6 policy contracts. I treated `master/GRILL-LOCK-deployment-fork-2026-07-01.md` as the claim-boundary source of truth: Step-1 is attach plus interface guardrail, confusion-resistant only; adversarial containment and "sole-writer by construction" claims remain shelved.

PRIMARY_BUCKET: still-open

## Four-bucket verdict

### 1. still-open

The conductor-core substrate remains materially open. The current system has relay protocol, linting, docs, and manually-authored file relays, but no designed running engine that owns the event loop, `submit()` transaction, `verify()` consume-once path, recovery, trusted config load, seat tool surface, or restart binding.

Required open substrate work:

- Process/concurrency model: specify a conductor-owned command queue or event loop where `submit()` read-validate-append and `verify()` check-and-burn execute in one serialized critical section. This is the one claim still licensed as "by construction": two honest racing seats cannot both consume the same authority if the trusted engine serializes the consume check and burn.
- Crash-atomic multi-file commit: design the transaction over canonical record, rendered relay, `INDEX.md`, and all target mailbox projections. A plain sequence of file writes is not enough; the design needs a journal/staging state, durable flush boundaries, recovery replay, and quarantine for malformed or partial records.
- Internal-fault disposition: every trusted-side throw, timeout, corrupt-record read, config-integrity failure, or observe/check failure must end in a typed held/fail-closed outcome for authority records. It must never silently accept, and it must not brick unrelated read/projection.
- Interface guardrail under attach: seats should receive only `submit()`/`project()`/`read()` affordances. Raw store, config, outbox, and operator-channel paths must be absent from the seat tool surface. This is a confusion-resistance fixture, not same-uid filesystem isolation.
- Trusted config load: config needs a conductor-owned artifact, loaded once at trusted startup, integrity-checked, versioned, and absent from all lane-visible tools.
- Restart/park recovery: parked lanes, away-token nonce burn, restart seat binding, store genesis, and GC are only named today; the engine has not specified the persisted tables and reconciliation order.

### 2. already-closed

Several parts should be promoted, not redesigned:

- The upstream protocol gives the relay grammar, phase/addressing discipline, DESIGN-REVIEW lineage gate, evidence fields, and truth-agnostic relay-lint tripwires. Its README explicitly defines the package as coordination protocol rather than orchestration runtime, and its adapter docs say host adapters are optional hardening, not protocol replacement.
- m-1 and m-2 already close the write-path contract shape: candidate held in-courier, pre-append form and lineage gates, no persisted `submitted` limbo, one atomic accepted append plus delivery, and distinct terminal rejected records on fail.
- m-3 already specifies the observe-as-send hook shape, closed write allowlist, evidence ladder, and declared-vs-observed integrity veto. m-7 needs to host the hook and order it correctly; it must not invent alternate evidence semantics.
- m-4 already fixes R2: model identity is payload/bookkeeping only, not a schema, authority, lineage, or dispatch gate input. m-7 must preserve that split while executing routing records.
- m-5 already fixes `slot_in` and `seat_archetype` as governed archetype atoms and records the Step-1 ceiling posture as host-config best effort, with later uniform enforcement routed to conductor runtime.
- m-6 already specifies the bucket taxonomy, ODB capture, park/wake state machine, away bridge posture, and m-1-owned inbound verdict-token mint/verify boundary.
- jcode provides useful runtime prior art for live attachments, connection maps, channel subscriptions, and resume/takeover decisions. It is prior art for process lifecycle, not the conductor's governance store.
- claude-code provides useful prior art for per-agent JSON inboxes, mailbox locking, poll/inject delivery, and teammate ergonomics. It is prior art for mailbox projection and concurrent inbox writes, not conductor-stamped identity.
- External file-system prior art is directly relevant: SQLite's rollback/super-journal model demonstrates atomic commit and crash recovery for multi-file transactions; POSIX `rename()` gives atomic replacement semantics but not durable directory persistence by itself; Maildir demonstrates temp-to-new delivery to avoid exposing partial messages.

### 3. product-overlapped

The largest risk is m-7 accidentally becoming a policy owner. These are overlaps, not m-7 decisions:

- m-1 owns identity, address minting, operator-FROM channel semantics, certification seams, and the store API contract. m-7 executes append/stamp inside the loop.
- m-2 owns the FieldSpec vocabulary, ownership types, required-when grammar, lineage roles, and fill-time authority semantics. m-7 renders and validates from it.
- m-3 owns observation semantics, evidence ladder, check registry, egress scan policy, and self-reported/integrity labeling. m-7 hosts the hook and enforces the prescribed fail-closed cases.
- m-4 owns routing policy, capability priors, deviation fields, R2 boundaries, and routing records. m-7 executes route dispatch/refusal and preserves model-as-payload.
- m-5 owns archetype vocabulary, topology/gate/ceiling presets, and posture declarations. m-7 spawns/hosts according to recorded fields.
- m-6 owns human-surface policy, buckets, ODB shape, park/wake semantics, away posture, expiry policy, and POST-not-GET transport invariant. m-7 persists and resumes the state machine and calls the m-1 token surface.

Open overlap hazards to route as targeted contract questions if still ambiguous in DESIGN:

- The exact Step-1 treatment for authority-class `self_reported` needs a concrete m-3/m-2 field home and field names before m-7 can enforce decision-2 fail-closed mechanically.
- m-2 records a row-parity pre-PLAN SHOULD for remaining observe fields. m-7 should not invent those rows, but Step-1 fixtures will be cleaner if those rows land before PLAN opens.
- m-2 notes `delivery_state` token drift from `bounced` to `rejected`, while m-6 still names bucket D as observe-bounce. m-7 can support a failing-edge abstraction, but the final token vocabulary should be reconciled before implementation.
- Away-token restart binding and nonce-burn are locked as later-step build carries. m-7 should surface the persisted seat-binding table and consume-once transaction in design, but m-1/m-6 own token semantics.

### 4. recommended-next

The c4 DESIGN artifact should become the conductor-core design-of-record and lock these engine decisions:

- A single-threaded serialized commit loop over all trusted mutating operations: `submit()`, `verify()`, config reload if any is ever allowed, store genesis, GC, mailbox delivery, and outbox enqueue. The design should state which reads may run concurrently and which reads must observe a committed snapshot.
- A transaction algorithm for record plus projections: write canonical candidate to a staging area, fsync content, commit by durable marker or atomic rename, fsync parent directories, then derive projections; recovery scans staging/commit markers and either completes, rolls back, or quarantines.
- A positive interface map for seat tools: `submit()`, `project()`, `read()` only. Negative fixtures should show that a confused seat asked to write raw store/config/outbox has no such tool, no exposed path, and no config value in its prompt/tool result surface.
- A config integrity model: signed or hash-pinned policy artifact, conductor-only path, loaded once at startup, digest stamped into runtime metadata, fail-closed if absent/mismatched, and no lane-visible read/write handle.
- A recovery state machine: startup validates genesis, scans canonical records and projections, rebuilds `INDEX.md`/mailboxes from canonical records if needed, quarantines corrupt records, restores parked-lane/seat-binding tables, and refuses authority consumption while recovery is incomplete.
- A seam matrix promoted from the inventory below: each hosted contract gets owner, doc/section, m-7 obligation, fixture-positive, fixture-negative, and "contract-question-raised?" flag.
- A claim-sweep fixture: no final design text should say malicious same-uid lanes cannot bypass raw files. The honest Step-1 claim is interface-level confusion resistance.

## Substrate evidence summary

- The upstream protocol is not a runtime. The release README says the package is a coordination protocol riding on host/tooling rather than an orchestration runtime (`the upstream protocol release corpus (not vendored), README.md:1-5`, `:79-94`). `relay-lint.py` is explicitly truth-agnostic and its ROLE/FROM check is a proxy-authorship tripwire, not proof of writer identity (`the upstream relay-lint tools/relay-lint.py:1-8`, `:555-575`). The adapter docs likewise call adapters optional host hardening and list bypass residuals (`the upstream relay-lint tools/adapters/README.md:1-3`, `:43-58`).
- jcode has usable runtime shapes, especially connection state and attach/resume handling: `ClientConnectionInfo` tracks client/session binding and processing state (`references/jcode/crates/jcode-app-core/src/server/debug.rs:45-60`), `SwarmMember` carries live attachments and durable member records (`references/jcode/crates/jcode-app-core/src/server/state.rs:186-277`), and resume logic rejects unsafe duplicate/takeover cases (`references/jcode/crates/jcode-app-core/src/server/client_session.rs:1100-1210`). But its comm handler accepts a `from_session` argument and ignores `_client_connections` for authority stamping (`references/jcode/crates/jcode-app-core/src/server/client_comm_message.rs:103-128`), so it is not a conductor-stamped FROM model.
- jcode channel subscriptions are good projection/index prior art: `swarm_channels.rs` mutates forward and reverse subscription indexes together (`references/jcode/crates/jcode-app-core/src/server/swarm_channels.rs:1-22`) and lists channels from the index (`:67-88`). That informs mailbox/projection design but does not solve the conductor's atomic record-plus-INDEX-plus-mailboxes commit.
- claude-code has the nearest mailbox write prior art: each teammate inbox is a JSON file, writes acquire a lock, re-read latest state after locking, append the new message, and write the file (`references/claude-code/src/utils/teammateMailbox.ts:1-7`, `:127-191`). `proper-lockfile` is a dependency (`references/claude-code/package.json:50-60`). But sender identity is self-derived from `getAgentName()`/team context and written as `from` (`references/claude-code/src/tools/SendMessageTool/SendMessageTool.ts:149-171`, `:238-249`), matching the known self-written `from` limitation in `references/README.md:12-13`.
- Runtime research supports attach plus per-seat channels for confusion-resistant identity, while warning against adversarial overclaim: conductor-as-MCP maps per-seat transport to identity, not runtime payload (`master/RUNTIME-RESEARCH.md:499-518`), and the honest attach posture is confusion-resistant rather than sole external sender (`master/RUNTIME-RESEARCH.md:528-546`).
- External crash-atomicity prior art supports a journaled design rather than ad hoc append order. SQLite documents atomic commit, crash interruption tolerance, and multi-file commit using rollback journals and a super-journal (https://www.sqlite.org/atomiccommit.html). Linux `rename(2)` documents atomic replacement of an existing path but not full durability alone (https://man7.org/linux/man-pages/man2/rename.2.html). Linux `fsync(2)` documents that directory entries need explicit directory fsync for durability (https://man7.org/linux/man-pages/man2/fsync.2.html). Maildir writes into `tmp` before moving into `new`, avoiding partially delivered messages for readers (https://cr.yp.to/proto/maildir.html).

## Seam inventory

| contract owner | contract source | m-7 execution obligation | audit note |
|---|---|---|---|
| m-1 Trust & Identity | `master/domains/m-1-trust-identity/design/2026-06-28-trust-identity-design.md:97-106`; `:108-120` | Run `submit()` as the sole trusted write path: resolve connection, stamp FROM/ROLE, hold candidate in-courier, call m-2 gates, atomically append accepted/rejected, deliver/project, and serve `read()`. | Design must specify transaction and recovery; do not restate m-1 identity policy. |
| m-2 Forms & Determinism | `master/domains/m-2-forms-determinism/design/2026-06-28-form-schema-design.md:31-74`; `:333-353` | Render visible fields, remove forbidden options, validate seat-scoped enums and required-set, run lineage gate over accepted graph plus candidate, then commit only through m-1 transaction. | Row-parity for remaining observe fields is a pre-PLAN SHOULD; m-7 should flag, not invent. |
| m-3 Observation & Evidence | `master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md:29-78` | Host observe hook at the Step-2 point, run conductor-owned checks from outside lane, stamp observed fields, enforce declared-vs-observed veto and egress failures as specified. | Step-1 must not demand observe-owned fields unless filler/gating is explicitly resolved. |
| m-4 Routing & Policy | `master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md:96-124` | Execute route dispatch/refusal, preserve routing record lineage, and ensure no model-derived predicate enters schema/authority/lineage/dispatch gates. | R2 is policy-locked; m-7's job is to keep engine gates from accidentally reading model identity. |
| m-5 Workflows & Archetypes | `master/domains/m-5-workflows-archetypes/design/2026-06-30-archetype-system-design.md:15-23`; `:158-173` | Spawn/host panes or sessions from recorded `seat_archetype`/template data, record ceilings per assignment, and apply host-config best-effort in Step-1. | Uniform ceiling enforcement is later standalone-runtime work; Step-1 should state its host-dependent bound. |
| m-6 Human Surface & Scheduler | `master/domains/m-6-human-surface-scheduler/design/2026-06-30-human-surface-scheduler-design.md:31-88`; `:157-173` | Persist bucket/ODB/park/wake state transitions, re-observe on wake, local-outbox-only external send, and call m-1 mint/verify for inbound verdict tokens. | Restart binding, nonce burn, and parked-lane recovery must be engine-level persisted state, but token semantics stay m-1/m-6-owned. |
| CTO/VP integrated architecture | `master/ARCHITECTURE.md:17-66`; `:287-337` | Execute the integrated typed-envelope store and human-surface/away-token architecture in order, with no direct lane path into store/config/outbox. | Collision arbitration stays CTO/VP. m-7 should surface questions, not arbitrate them silently. |
| Deployment fork lock | `master/GRILL-LOCK-deployment-fork-2026-07-01.md:30-49`; `:65-88`; `master/DESIGN-REVIEW-2026-07-01.md:121-162` | Build interface-level guardrail, trusted config load, serialized commit, crash recovery, local outbox, and claim sweep fixtures. | This is the claim boundary: confusion-resistant interface mechanisms stay; adversarial security claims collapse. |

## Required DESIGN fixtures

- Positive fixture: two simultaneous `verify(token)` attempts enter the serialized loop; one commits nonce-burn/decision, the other observes already-burned and fails closed.
- Positive fixture: accepted submit produces exactly one canonical record, one rendered relay, one `INDEX.md` row, and the expected mailbox projections after crash-free commit.
- Recovery fixture: crash after staging but before commit marker leaves no accepted record visible; recovery removes or quarantines staging.
- Recovery fixture: crash after canonical commit but before all projections leads startup to rebuild `INDEX.md` and mailboxes from canonical records without duplicating authority consumption.
- Fault fixture: form validator throws on an authority-bearing relay; outcome is held/fail-closed with author-visible failure, not accepted, not process-bricking.
- Guardrail negative fixture: seat is instructed to edit raw store/config/outbox; exposed tools contain only `submit()`/`project()`/`read()`, no raw path or write capability is present.
- Config fixture: trusted config digest mismatch at startup prevents authority-bearing submit/verify while allowing a diagnostic/status projection.
- Claim-sweep fixture: design text contains no malicious-lane "unbypassable", "sole-writer by construction", or same-uid write-exclusion claims for Step-1 attach.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c4-audit-m-7/AUDIT-implementer-20260701-155145.md` - OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-audit-m-7` - OK
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`
- `git -C pcode status --short` - clean

ACTIONS_GIT_REF: wrote `master/relays/c4-audit-m-7/AUDIT-implementer-20260701-155145.md` and appended `master/relays/INDEX.md`; no code/source/`pcode/`, no PLAN, no spike, no domain design edit.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo; `git -C pcode status --short` clean.
Next requested action: m-7 planner/implementer pair reconcile this audit with the planner audit, then route c4 DESIGN with seam matrix and GRILL_REQUIRED.
