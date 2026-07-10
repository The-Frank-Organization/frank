## AUDIT - s4-wire.implementer / s4-wire-audit

ROLE: Implementer
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: s4-wire-audit
PARENT_DISPATCH_ID: s4-dispatch
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
FROM: s4-wire.implementer
TO: s4.orchestrator-planner
CC: s4.orchestrator-reviewer, operator
SUBJECT: AUDIT - S4 wire-up ground truth, reuse gates, and open implementation scope

IN_REPLY_TO: `.relays/s4/s4-wire-audit/AUDIT-orchestrator-planner-20260705-010636.md`

## Routing and authority

- Routing accepted: the dispatch is addressed to `s4-wire.planner, s4-wire.implementer`, so this Implementer seat is directly in `TO`.
- Authority accepted: read-only AUDIT only. I did not modify source code, tests, sprint docs, design docs, or runtime behavior.
- Independence observed: before filing this audit, I did not read any sibling audit from `s4-wire.planner`. The audit directory contained only the orchestrator dispatch when checked.
- Output shape observed: this relay is filed before implementation, indexed in `.relays/s4/INDEX.md`, and linted as an exact file.

## Onboarding and baseline proof

- HEAD: `6987367`.
- `s3-close`: `b5a2c95`.
- `git diff --stat s3-close..HEAD`: docs-only sprint ledger delta, 3 files and 83 insertions:
  - `docs/sprints/2026-07-04-s3-slice-3/RECONCILE.md`
  - `docs/sprints/2026-07-05-s4-slice-4/RECONCILE.md`
  - `docs/sprints/2026-07-05-s4-slice-4/ROADMAP.md`
- Fresh full test battery: `go test -count=1 ./...` passed.
- Fresh vet battery: `go vet ./...` passed with no output.
- Fresh rendered-form size probe: virtual-overlay test passed with `rendered_description_response_bytes=2222 fields=28 digest=e6db1acff826fe72c9c53e81b6ce13916116120ed5672f7434a3aefa4cb231c5`.
- S1-S3 ledger/source onboarding covered: S1 provenance/transport and D5 self-reported boundary, S2 field/form validation and bounce posture, S3 durable store/recovery/owed-item posture, current S4 ROADMAP/RECONCILE, and current source surfaces in `cmd/frank`, `internal/channel`, `internal/seat`, `internal/fieldspec`, `internal/store`, `internal/recover`, `internal/config`, `internal/obligation`, and fixtures.

## Spec-to-exit-gate map

- S4 exit gate asks for six IN items: per-seat MCP shim, rendered-form schema through MCP, seat lifecycle hardening, section 7 config-change record, operational surface, and usage-data posture.
- m-7 section 8 guides S4 on conductor attach lifecycle, pipe wake, MCP surface, guardrail preservation, and schema-as-form.
- m-7 section 7 requires trusted committed config, no hot reload, config absent from seat-authorable surface, and fresh-store recovery diagnostics for config mismatch.
- m-1 sections 4 and 5 own channel identity, stamped provenance, confusion resistance, and the `submit`/`project`/`read` plus admin-time `mint_seat` boundary.
- m-1 section 6 owns `FROM`, `ROLE`, delivery/diagnostic fields, and conductor-internal provenance.
- m-1 section 13.3 explicitly leaves DI-2 tool surface, submit TOCTOU, and credential lifecycle generation/rotation/revocation as plan items. S4 must not pretend those lifecycle gaps are already closed.
- ARCHITECTURE C4.1 and C4.3 bind S4 to the trusted config and wire-up boundary: conductor host affordance only, no new seat authority, and D5 remains `self_reported` until observe.
- `s3-scope-q1` says the config-change record is deferred to wire-up under m-7 guidance, m-1 fidelity on `record_kind`, and crash-matrix coverage for the class.
- The S4 IN list matches these gates, but the current code already closes several sub-surfaces. The implementation phase should promote existing primitives instead of rebuilding them.

## Four-bucket verdicts

### IN-1: Per-seat MCP shim that exposes only submit/project/read and emits nudge/poll hints

PRIMARY_BUCKET: still-open

- still-open: No real per-seat MCP shim exists. Current transport is a bespoke socket JSON-RPC dialect with `session/connect`, `tools/list`, `tools/descriptions`, `tools/call`, and `notifications/nudge`.
- already-closed: The conductor-side tool handlers and authenticated socket client exist and are reusable.
- product-overlapped: A shim can translate between MCP stdio and existing `internal/channel.Client`; the conductor does not need a new authority model.
- recommended-next: Build the shim as a translation layer over `internal/channel.Client`; do not expand the tool list or make conductor itself claim MCP fidelity unless the MCP handshake and schema contracts are actually implemented.

### IN-2: Rendered-form schema and digest through MCP, including hidden FROM/ROLE and bounce fidelity

PRIMARY_BUCKET: still-open

- still-open: The schema is not yet delivered through an MCP shim because the shim does not exist.
- already-closed: `cmd/frank` already serves a describe response with `Tools: [submit project read]`, per-seat `SubmitSchema`, and `FormDigest`; `internal/fieldspec` hides system/computed fields and validates by re-rendering the digest.
- product-overlapped: The real work is promotion across the shim, not a second renderer.
- recommended-next: Reuse `DescribeTools` and the existing render/validate path; add tests that prove MCP `tools/list` or equivalent exposes the existing submit schema without offering `FROM` or `ROLE`, and that stale digest bounces keep current path-clean behavior.

### IN-3: Seat lifecycle hardening for credential minting, reconnect, second-connect policy, custody, restart recovery, and rejection evidence

PRIMARY_BUCKET: still-open

- still-open: The active second-connect constraint is absent, rotation/revocation surfaces are absent, and custody is weak because mint prints the credential to stdout and operator-submit accepts `-credential` on the command line.
- already-closed: `seat.Mint` creates one durable credential per seat, rejects duplicate mint and reserved `system`, `Resolve` authenticates by credential, durable mailbox recovery reissues pending delivery nudges, and fixtures cover minted-seat auth and duplicate mint rejection.
- product-overlapped: Restart recovery and persistent identity are partially handled by existing store/seat/mailbox mechanisms; active-channel policy is the missing piece.
- recommended-next: Implement one-active-channel-per-credential with reject-active-duplicate and recover-only-proven-dead semantics. Treat live supersede, remint, rotation, or revocation as locked-contract touches requiring escalation.

### IN-4: Section 7 config-change record path

PRIMARY_BUCKET: still-open

- still-open: There is no `config-change` record kind, mutation class, crash-matrix row, recovery class, or applicability-map row for post-genesis config updates.
- already-closed: `store.Init` pins committed config bytes and digest at genesis, phase-0 recovery validates genesis/config match, the config digest is path-clean, crash harness classes exist for genesis/gc/owed, and the owed-item/disposition mechanism can carry the existing OI-S3 config-change obligation.
- product-overlapped: This should evolve the existing store/config/recovery/owed harnesses, not introduce an independent configuration subsystem or re-genesis path.
- recommended-next: Add a first-class config-change mutation/record that commits new config state through the existing serialized store path, updates recovery/applicability/crash fixtures, and closes the owed item with explicit disposition evidence.

### IN-5: Operational surface for boot, socket discovery, minting, status, logs, and stop/restart

PRIMARY_BUCKET: still-open

- still-open: The user-facing operational runbook/status surface is incomplete; there is no documented start/stop/status convention for team-store and socket paths.
- already-closed: `cmd/frank` has `-init`, `-mint`, `-operator-submit`, `-store`, `-socket`, `-registry`, `-credential`, and server modes, and fixtures cover binary-level mint, submit/project/read, read-only diagnostics, and startup quarantine.
- product-overlapped: Existing flags are enough for the narrow CLI substrate; the gap is operator ergonomics, naming, and evidence.
- recommended-next: Document and fixture the exact boot/mint/operator-submit/status path while preserving the current authority split between admin minting and seat-level submit/project/read.

### IN-6: Usage-data posture: no analytics, expose existing records read-only

PRIMARY_BUCKET: already-closed

- already-closed: The current surface exposes committed records through `project` and `read`; there is no analytics collector or extra usage telemetry surface in the audited code.
- still-open: The posture needs documentation and probably a read-only operator example showing how to inspect usage-relevant committed records without adding analytics.
- product-overlapped: Existing record projection is the right mechanism; a new analytics subsystem would be scope drift.
- recommended-next: State the posture plainly in the operational docs and tests: usage evidence comes from committed records/projections only, with no new collection path.

## Duplicate and reuse gate

- `internal/channel.Client`: reuse. It already authenticates and calls conductor tools over the local socket. It is not an MCP shim and should not be described as MCP-compliant until wrapped or replaced by one.
- `test/seatproc/testseat.go`: do not promote as implementation substrate. It is a stub that prints `["submit","project","read"]` and is useful only as a test-seat placeholder.
- `cmd/frank -operator-submit`: reuse for binary-level custody/workflow tests, but do not bless `-credential` as secure custody; command-line credentials are inspectable by local process observers.
- `cmd/frank -mint`: reuse for admin-time minting and duplicate/reserved-seat rejection. Missing rotation/revocation remains open.
- S3 describe render: reuse. `DescriptionResponse.SubmitSchema` and `FormDigest` are the correct schema/digest source for the shim.
- S1/S2 crash harness: reuse by adding the config-change mutation class instead of creating a parallel harness.
- Owed mechanism: reuse for OI-S3 config-change closure and disposition proof.
- `store.Init` and phase-0 recovery: reuse as trusted-config baseline, but do not use re-init or re-genesis as the config-change story for an existing store.

## m-1 / m-7 boundary surface enumeration

- m-1-owned, must preserve: seat identity, credential binding, authenticated seat metadata, stamped `FROM`/`ROLE`, `submit`/`project`/`read`, admin-time mint semantics, provenance fields, `record_kind`, and path-clean bounce posture.
- m-7-guided, may implement under S4: host-side attach lifecycle, local socket/stdio bridge, poll/nudge hints, recovery wake across the bridge, schema-as-form carried through the bridge, trusted config diagnostics, and section 7 config-change record mechanics.
- Locked-contract or escalation surfaces: adding seat-authorable tools beyond `submit`/`project`/`read`, live credential supersede, rotation/revocation semantics, remint, hot config reload, new authority-bearing fields, and any claim that D5/`record_integrity` is observed rather than `self_reported`.

## Claim probes

### A. Nudge broadcast and recovery nudge all-seats list

- Current submit nudge is broadcast to all connected clients, not targeted. `Server.Push` calls `broadcast`, and `broadcast` iterates all active clients.
- Recovery wake queues one push with `seats` set to the full pending-delivery seat list from `PendingDeliverySeats`, then `QueuePush` sends it to all active clients and also stores it globally for later authenticated clients.
- Implementation implication: if the shim claims per-seat nudge semantics, it must filter or route at the shim/client layer. Current conductor nudge is all-clients broadcast.

### B. Global pending flushed to every auth and never cleared

- Confirmed. `Server.pending` is global `[][]byte`; `QueuePush` appends to it; `flushPending` copies and writes all pending messages to every newly authenticated connection.
- There is no clear-on-delivery or per-seat acknowledgement path in current code.
- Implementation implication: startup/recovery wake can duplicate across clients and reconnects. That is acceptable only as a poll hint, not as exactly-once delivery evidence.

### C. No second-connect constraint

- Confirmed. Auth accepts any credential that `seat.Bindings.Resolve` recognizes, sets `c.authed`, and flushes global pending. There is no active credential/session map and no reject-active-duplicate path.
- Existing tests cover duplicate mint rejection and unknown credential rejection, but not active duplicate connect rejection.
- Implementation implication: VP-W1 remains open. The policy should be reject-active-duplicate or recover-only-proven-dead, with broader supersede semantics escalated.

### D. No rotation or revocation surface

- Confirmed. `seat.Mint` creates one credential per seat and rejects duplicate mint. `Resolve` is lookup-only. There is no rotate, revoke, generation counter, expiry, or multi-credential history.
- Implementation implication: do not claim lifecycle hardening covers rotation/revocation. Either leave explicitly out of S4 or escalate if required.

### E. Credential custody stdout and CLI flag

- Confirmed. `cmd/frank -mint` prints `credential=<hex>` to stdout. `cmd/frank -operator-submit` accepts `-credential`.
- Fixtures prove credentials do not leak into records/mailboxes/projections/outbox, but local stdout/process-argument custody remains weak.
- Implementation implication: operational docs must call this out; a stronger custody channel is a separate product/security decision.

### F. Scanner default ceiling and rendered form payload size

- Confirmed. Both server and client read loops use `bufio.NewScanner` with default token sizing; no custom buffer is configured.
- Fresh measurement: the current rendered description response for `SITREP` at medium tier is 2,222 bytes with 28 fields, well below the default scanner ceiling.
- Implementation implication: current form size is safe, but MCP/shim payload growth should either stay bounded by tests or move off scanner/default-token assumptions.

### G. I-PH bridge surface census

- Existing path-clean positives: tool-call errors are sanitized through `safeError`; tests cover sanitized unknown credential/tool errors, config digest path independence, read-only diagnostics on digest mismatch, and no credential leakage into records/mailboxes/projections/outbox.
- Bridge gaps: there is no MCP shim, so shim diagnostics and MCP error paths have no I-PH fixtures. Auth/connect failures are path-clean in observed strings, but the shim must preserve that guarantee for stdio/MCP-facing errors.
- Implementation implication: add path-clean assertions for shim connect/auth/tool-call errors and stale schema bounces, not just the existing conductor socket path.

### H. MCP fidelity versus bespoke JSON-RPC dialect

- Current protocol is not MCP. It uses `session/connect`, `tools/list`, `tools/descriptions`, `tools/call`, and `notifications/nudge` with local structs. It does not implement MCP `initialize`, standard tool schema presentation, or official notification semantics.
- This is not fatal if S4 builds a per-seat shim that translates MCP to the existing conductor socket. It is a blocker if the implementation claims the conductor socket itself is MCP.
- Implementation implication: keep the conductor internal dialect private behind the shim unless a locked-contract change explicitly authorizes a direct MCP server rewrite.

## Second-connect ground truth

Ground truth: active duplicate connections are currently allowed.

- A second connection using the same valid credential passes the same `Resolve` path as the first connection.
- No current map tracks credential-to-active-connection, seat-to-active-connection, or generation.
- No existing test expects second-connect rejection.
- Therefore any S4 plan that says the second-connect constraint is already implemented is false. It must be implemented and tested, or explicitly deferred/escalated.

## Honesty-line census

- S4 may honestly claim transport/provenance improvements only after implementation evidence.
- S4 may not claim observed done-state or observed `record_integrity`; README and S4 ROADMAP keep D5/`record_integrity` self-reported until Step-2 observe.
- Recovery nudge is a poll hint, not delivery proof.
- Global pending wake is a wake/recheck mechanism, not per-seat durable acknowledgement.
- Credentials are durable identity bindings, not a complete lifecycle system.
- Current describe schema is conductor-rendered and digest-validated, not yet MCP-carried.
- Config is genesis-pinned and phase-0 validated, not yet post-genesis config-change capable.

## Recommended implementation order

1. Add tests for the missing active duplicate credential policy and implement the smallest active-channel registry that rejects live duplicates while allowing proven-dead recovery.
2. Build the per-seat MCP shim as a translation layer over `internal/channel.Client`, preserving only `submit`, `project`, and `read`.
3. Promote existing `DescribeTools` schema/digest through the shim and add MCP-facing schema/bounce/path-clean tests.
4. Add config-change as a first-class store mutation/record with recovery/applicability/crash-matrix coverage and OI-S3 disposition.
5. Document and fixture the operational flow around init, mint, socket path, operator-submit, status/readback, and restart.
6. Document usage-data posture as read-only committed-record inspection; do not add analytics.

## Actions

- Performed read-only audit of the addressed S4 dispatch and current source/spec/fixture surfaces.
- Ran fresh local verification commands listed below.
- Wrote this audit relay and appended the S4 index row.

ACTIONS_GIT_REF: read-only audit relay written at `.relays/s4/s4-wire-audit/AUDIT-implementer-20260705-012253.md`; index row appended at `.relays/s4/INDEX.md`; no source/test/spec/runtime behavior changes made.

FINAL_GIT_STATUS_SHORT: none - clean tree

## Verification

- E2: `git rev-parse --short HEAD` -> `6987367`.
- E2: `git rev-parse --short s3-close^{commit}` -> `b5a2c95`.
- E2: `git diff --stat s3-close..HEAD` -> docs-only sprint ledger delta, 3 files and 83 insertions.
- E2: `git status --short --untracked-files=all` before audit write -> clean.
- E2: `go test -count=1 ./...` -> pass.
- E2: `go vet ./...` -> pass.
- E2: `go test -overlay="$tmp/overlay.json" ./internal/fieldspec -run TestAuditMeasureRenderedResponseSize -count=1 -v` -> pass; logged `rendered_description_response_bytes=2222 fields=28`.
- E2: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s4/s4-wire-audit/AUDIT-implementer-20260705-012253.md` -> pass.
- E2: `git status --short --untracked-files=all` after audit write and index append -> clean.
