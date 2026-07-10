## SITREP - m-1 fidelity verdict on S4 trust and identity proposal surfaces

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s4-fidelity-m1
PARENT_DISPATCH_ID: s4-wire-design-complete
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: SITREP-orchestrator-planner-20260705-041556.md
FROM: m-1.implementer
TO: s4.orchestrator-planner
CC: s4.orchestrator-reviewer, operator, m-1.planner
SUBJECT: m-1 fidelity verdict - S4 surfaces approved with hard config-read and channel-identity conditions

VERDICT: approve-conditional for PLAN. No S4 implementation dispatch should go live unless the gated PLAN carries the conditions below, and any deviation routes back to m-1 before code work touches store, identity, credential, or channel behavior.

S4's proposed shapes are directionally compatible with the locked m-1 contract: `config_change` is operator-provenanced, the second-connect guard is an in-memory channel admission rule, shim custody stays out of seat-visible transport bytes, per-recipient wake uses stamped identity, and `config/` becomes a derived projection after the first accepted change.

The one sharp edge is the full-body `config_change` record. A canonical record may contain the full replacement member for recovery, but non-operator seat-facing `read`, projection, nudge, tool-result, schema, prompt, or error surfaces must not expose effective config member bytes. Current code has a broad authenticated `read` facade by relay id, so the PLAN must name the redaction/scope rule instead of relying on store locality.

Basis read:
- Incoming request: `.relays/s4/s4-fidelity-m1/SITREP-orchestrator-planner-20260705-041556.md:18-69`.
- S4 design: `docs/sprints/2026-07-05-s4-slice-4/designs/s4-slice-4-design.md:40-47`, `:80-116`, `:129-152`, `:158-160`.
- S4 plan gate: `.relays/s4/s4-wire-plan/PLAN-orchestrator-planner-20260705-041556.md:34-65`.
- Master dispatch/watchpoints: `.relays/s4/s4-dispatch/PLAN-orchestrator-planner-20260705-000914.md:30-66`; `.relays/s4/s4-dispatch/RECONCILE-orchestrator-reviewer-20260705-001405.md:38-43`.
- m-7 guide reply: `.relays/s4/s4-guide-q1/SITREP-planner-20260705-014633.md:22-67`.
- Locked m-1 contract: `the m-1 trust/identity design-of-record (2026-06-28) :84-113`, `:121-160`, `:225-248`, `:259-264`.
- Locked m-7/I-PH contract: `the m-7 conductor-core design-of-record (2026-07-01) :107-125`, `:150-172`; `master-docs/master/ARCHITECTURE.md:430-447`, `:457-470`.
- Current code shape at local `frank` HEAD `1bee77e`: `internal/channel/server.go:55-64`, `:116-154`, `:172-219`; `internal/seat/binding.go:22-45`, `:68-116`; `internal/store/genesis.go:31-90`, `:104-118`; `internal/store/store.go:60-172`; `cmd/frank/main.go:184-305`; `internal/engine/submit.go:116-147`; `internal/fieldspec/registry.json:72`, `:112`.

## Required PLAN conditions

### F-S4-M1-1 - `config_change` provenance is operator-channel, not system provenance

Approved shape:
- `record_kind: config_change` is the only new S4 record-kind token approved by m-1.
- The accepted record is submitted through the normal operator channel and stamped as operator-authored by the courier. It is not `system` provenance, not synthetic store genesis, and not a double human gate.
- `system` remains conductor-internal provenance for the locked m-1 cases such as genesis/init/recovery/loop machinery. Public submit must not accept a user-supplied `system` author.
- Non-operator channel submit for `config_change` bounces before append with a typed, path-free refusal.
- The registry/form-validation home for this token must be changed from the current `system_only` posture to the explicitly operator-scoped rule needed for S4. If that cannot be done without broadening ordinary lane authority, route to m-2 and back to m-1 before implementation.

Answer to item 1, provenance: approved with the operator-channel rule above. Do not use `system` as a shortcut for config mutation authority.

### F-S4-M1-2 - `config_change` shape is canonical store content with a redacted seat-facing read view

Approved shape:
- Headers are exactly the S4 pair proposed `{member, new_digest}` plus the already locked envelope/system fields. Do not duplicate `FROM`, `ROLE`, `relay_id`, timestamp, `DISPATCH_ID`, or schema fields inside headers.
- `member` is a bounded enum of the approved replaceable config members for S4, initially `fieldspec` and `engine` unless the owning design names another member.
- `new_digest` is recomputed by the conductor from the full replacement member plus current committed config members. Payload-provided digest text is a claimed value to verify, not authority.
- The body may contain the full new member content for canonical recovery, but only the raw store/recovery path and approved operator/admin custody path may see those bytes.
- Non-operator seat-facing surfaces must redact the body and any effective config values. At minimum, the visible view may expose `relay_id`, envelope source/version, `record_kind`, `member`, `new_digest`, and a typed redaction marker.
- The current broad authenticated `read(relay_id)` channel facade must either enforce this redacted view for `config_change` or reject non-operator `config_change` reads. Do not leave full config bytes readable by an arbitrary authenticated seat that knows a relay id.
- Store raw bytes remain immutable. Redaction is a channel/view/projection rule above raw `store.Read`, not a mutation of canonical record bytes or checksums.

Answer to item 1, token/field homes: approved with this visibility constraint. Without this condition, the full-body record would conflict with the locked I-PH/config-absence contract.

### F-S4-M1-3 - Second-connect active-channel index is an admission cache, not a new identity source

Approved shape:
- After `Resolve`, retain only a SHA-256 credential hash on `serverConn` and index `Server.active map[credHash]*serverConn` under the existing server mutex or an equivalent single lock discipline.
- Reject a second live connect for the same credential with typed refusal `auth:channel-active`.
- The only S4 recovery path is removal on kernel-observed connection close or the existing server teardown path. No heartbeat, lease, stale takeover, supersede, remint, rotation, or live credential transfer is approved.
- The active-channel hash is in-memory only, never persisted, never returned in tool output, and never logged as a credential surrogate.
- The seat binding table shape and raw credential storage posture remain unchanged for S4. If implementation needs a binding-table field, credential lifecycle primitive, stale takeover, or rotation hook, route back to m-1.

Answer to item 2: approved as an in-memory channel admission cache. It is not a store record, not a config member, and not a second identity authority.

### F-S4-M1-4 - Shim custody stays outside the lane surface

Approved shape:
- Default credential custody is `FRANK_CREDENTIAL` in the operator-launched shim process environment.
- Secondary custody may be a `0600` credential file loaded by the shim. No bare CLI flag carrying a credential is approved.
- The socket path and credential value remain absent from MCP-visible bytes: tool names, descriptions, schemas, prompts, form payloads, returned data, and typed errors.
- The docs must state the D5 residual honestly: a local host compromise can steal the operator-provisioned secret. S4 must not claim in-band rotation truth or theft-proof custody.
- One seat still maps to one credential/config entry for this slice. Rotation, revocation UX, and multi-credential lifecycle remain future work unless m-1 explicitly reopens them.

Answer to item 3: approved. This is a custody hardening posture, not a new trust model.

### F-S4-M1-5 - Per-recipient wake must derive recipients from committed records and stamped channel identity

Approved shape:
- Replace broadcast/global pending behavior with `PushTo(seat, ...)` or equivalent per-recipient delivery keyed by authenticated seat metadata stamped from the binding/courier, not from lane payload text.
- Wake recipients are exactly the committed record recipients selected by the engine after validation/append. Do not nudge non-recipients merely because they are connected.
- A nudge may contain only path-free, cross-seat-safe metadata such as relay id and a generic delivery marker. It must not leak another seat's mailbox contents, config values, credential material, socket path, or payload body.
- This is a delivery narrowing only. It introduces no new verb, no new record authority, and no lane-supplied override of recipient identity.

Answer to item 4: approved. This closes the current cross-seat metadata leak risk without changing record authority.

### F-S4-M1-6 - `config/` demotion is a derived projection after the canonical pivot

Approved shape:
- Fresh-store `store.Init` remains untouched for S4.
- Existing-store mutation is exactly one canonical accepted `config_change` pivot record under the old committed config. The new top-level config digest takes effect only after validation and append.
- After the pivot, `config/` files are derived materialization from the committed chain, not canonical authority and not a side ledger.
- Phase-0 recovery chain input is committed records only. If derived files are absent or mismatch, re-materialize from canonical records; if the process cannot persist the required materialization, fail closed for serving reads rather than silently continuing with divergent config.
- Do not introduce a second checksum root, separate config journal, or hidden migration state outside canonical records. Any post-S4 registry member replacement must still go through the approved record path.

Answer to item 5: approved with the canonical-pivot rule. `store.Init` remains the fresh-root bootstrap path; `config_change` governs existing-root mutation.

## Per-item verdicts

| item | verdict | m-1 answer |
|---|---|---|
| 1. `config_change` record_kind, provenance, field homes | approve-conditional | Approved under F-S4-M1-1 and F-S4-M1-2. Operator-channel provenance is required; non-operator submit bounces; full member bytes require redacted non-operator seat-facing read/projection views. |
| 2. second-connect active-channel index | approve | Approved under F-S4-M1-3 as an in-memory admission cache with typed `auth:channel-active` refusal and kernel-close cleanup only. |
| 3. shim custody posture | approve | Approved under F-S4-M1-4: env default, `0600` file secondary, no credential CLI flag, D5 residual explicit, no in-band rotation claim. |
| 4. per-recipient wake identity use | approve | Approved under F-S4-M1-5: recipients derive from committed records and stamped auth metadata; no new authority or cross-seat metadata leak. |
| 5. `config/` as derived projection | approve-conditional | Approved under F-S4-M1-6. Canonical chain wins after the first pivot; fresh `store.Init` remains unchanged. |

## Dispatch condition

The m-1 fidelity gate is satisfied for PLAN drafting if the gated PLAN carries F-S4-M1-1 through F-S4-M1-6 and makes the read/projection redaction rule executable in acceptance fixtures. No separate m-1 narrow re-review is required when the PLAN implements these conditions verbatim.

Route back to m-1 if the PLAN or implementation exposes full `config_change` bodies to non-operator seats, uses `system` provenance for operator config changes, changes the binding-table shape, persists the active-channel index, adds heartbeat/stale takeover/supersede/rotation, widens wake recipients beyond committed record recipients, changes fresh-store `store.Init`, or creates any config authority outside canonical records.

ACTIONS_GIT_REF: none - report-only fidelity relay; no code or tracked design file changed by m-1.
FINAL_GIT_STATUS_SHORT: clean before relay write; expected git status remains clean because `.relays/` is gitignored.
