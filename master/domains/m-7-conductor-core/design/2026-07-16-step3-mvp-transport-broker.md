# m-7 stage-1 contract — the shared TRANSPORT/CLIENT boundary + the authenticated CHANNEL/BROKER contract (F60/F64/F66) + the conductor-identity PRODUCER contract (F68/F65)

**DESIGN_DOC_ID:** step3-mvp-design-m7-transport-broker
**Revision:** r11 — folds the r10 pair review `step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260717-193032.md` (must-revise, R10-F1: the attach-suspended predicate made TOTAL — the §2.5 PREPARING barrier joins the §2.4 causes, with its own recordable reason and the reattach-during-PREPARING fixture cut); r10 = `da1ed802…` (the R9-F1 suspension-precedence fold), r9 = `ed66e038…` (the D-3 taxonomy fold), r8 = `ab0ed428…` (pair-approved 20260717-025246; VOID from r9 on), r7 = `fff04fcf…`, r6 = `f072bd99…`, r5 = `3e88bce8…`, r4 = `28b58585…`, r3 = `8862780a…`, r2 = `8cf86753…`, r1 = `bddf868b…` (superseded bytes, preserved in the trail).
**Owner:** m-7 (Conductor-Core) · planner-authored · for m-7.implementer pair review of the FINAL bytes (§7 stage 1), then consumer confirmation (m-9 · m-10 · m-1 · m-2, + m-3 on the §3 scope edge) on master's direction.
**Date:** 2026-07-16 · **Basis:** the ratified `master/STEP-3-MVP-AMENDMENT.md` r7 (`2f75f2a1…`) §§2/2b/5/6/7/10 + F58/F60/F63/F64/F65/F66 · the dispatch `step3-mvp-design-m7/DESIGN-orchestrator-planner-20260716-041630.md` + its supplement `…-043459.md` · m-1's stage-1 contract rev2 (`7baffe40…`, in r3 pair re-review — the §E consumption rule) · m-10's stage-1 contract `master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md` (§A.1/§A.2/§B.1/§B.4/§F consumed in §2.10/§2.11; two named cross-interface deltas routed for their confirmation, §4) · landed code at `frank@502e06c`: `internal/channel/server.go` (auth `:277-309`, verb serve `:311-340`, push `:167-215`, client `:418-589`), `cmd/frank-mcp/{main,mcp,schema,errors}.go`, `cmd/frank/main.go:463-490` (the read/quarantine path), `internal/engine/loop.go:64-128` (serialized quarantine processing), `internal/store/quarantine.go:69-110` (idempotent `QuarantineOne` + incident-exists skip; `TestQuarantineOneIsIdempotent`) · my m-5-ceiling owner confirmation `step3-amend-m5-ceiling/SITREP-planner-20260715-060542.md`.
**Bounds:** design-only.
No conductor protocol verb, store record/member, or seat-surface change is required by any clause below (verified §5); the describe carriage uses the EXISTING `tools/descriptions` channel method.
The §3 producer adds one conductor OUTPUT artifact (the serve-start stamp) — new conductor serve-path code, not a protocol/store field.
No DESIGN-lock, PLAN, T4 token, or code.
The Master+VP interface-lock is the gate.

---

## §E — The m-1 edge: APPROVED source bytes (closed at r5)

m-1's contract is **pair-APPROVED, byte-bound, at SHA-256 `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`** (approve `step3-mvp-design-m1/DESIGN-REVIEW-implementer-20260716-061153.md`; completion SITREP `…-061835`).
Every clause below marked **[m-1-edge]** consumes those approved bytes: their §1.4a pins S-B resolution to exactly this contract's §2.12 sink path; their §2.7 matrix names the token-death-on-broker-restart leg; their §1.4b closed surface (three relay verbs + typed `Describe` + per-push fence) matches this contract's §2.8 exactly.
Any post-approval change to the m-1 bytes voids their approval and re-opens the marked clauses.
m-7's CONSUMER CONFIRMATION of the m-1 contract remains a separate act routed by master — this contract's consumption is not that confirmation.

## §1 — The shared transport/client boundary (amendment §5; the m-7 half of the 3-way seam)

### 1.1 What exists today (grounded at `502e06c`)

The bare protocol client ALREADY lives in the m-7 package: `channel.Client` — `Dial`/`DialAuthenticated`, `Call`, `ListTools`/`DescribeTools`, `NextPush`, the frame codec + limit, the request/response correlation loop (`internal/channel/server.go:418-589`).
What lives in `cmd/frank-mcp` (package `main`) today, mixed with the MCP skin, is the *managed* transport layer:
- lazy connect + close lifecycle — `ensureClient`/`closeClient` (`mcp.go:278-308`);
- a reconnect-and-retry-once path — `callWithReconnect` (`mcp.go:172-190`) — which §1.5 REPLACES rather than hoists: the live code retries after EVERY `Client.Call` error, including served application rejections and context cancellation, and has no single-flight rule under concurrent callers;
- the error-scrub taxonomy — five stable classes `shim:{conductor-unreachable, auth-failed, connection-lost, frame-too-large, protocol-error}` (`errors.go:11-44`);
- credential sourcing — `FRANK_CREDENTIAL` env or a partially-checked 0600 file (`main.go:24-55`) — which §2.12 REPLACES rather than hoists.

### 1.2 The hoist — `channel.ManagedClient` (the m-7-owned shared transport)

`internal/channel` gains a managed client wrapping `channel.Client`, owning EXACTLY the transport concerns:
1. **Connect lifecycle:** lazy dial-on-first-use; authenticated connect (`session/connect`) when constructed with a credential source (§2.12); explicit `Close`.
2. **Reconnect/retry:** per the §1.5 linearizable retry contract (classification + single-flight + retry gate).
3. **Error taxonomy:** the five scrub classes hoist BYTE-IDENTICAL (`shim:*` strings unchanged — frontend-observable vocabulary), plus the typed broker classes named in §2 (`broker:stale-epoch`, `broker:unknown-outcome`, `broker:record-unavailable`, `broker:suspended`).
   Scrubbing stays mandatory at the shared-client boundary: raw transport errors (paths, socket names, dial detail) never cross into a frontend/tool result (I-PH grain).
4. **Push surface:** the existing best-effort `NextPush` stream (16-slot buffer, drop-on-full, `server.go:582-586`) — best-effort/advisory semantics preserved exactly; reader continuity across connection replacement per §1.5.3.
5. **Frame limit + codec:** unchanged, already in-package.

**The caller seam (placement-agnostic consumers):** the shared package defines ONE small calling interface with exactly TWO operations:
- `Call(ctx, name canonical-relay-verb, args) → (result bytes, error-class)` — the three canonical relay verbs;
- `Describe(ctx, DescribeRequest) → (DescriptionResponse, error-class)` — the typed metadata/rediscovery operation, carried on the EXISTING `tools/descriptions` channel method; its return is the raw `DescriptionResponse`; ALL FieldSpec/form interpretation stays in m-2's module.
`Describe` is deliberately typed as a metadata operation, NOT a fourth relay verb: it never appears in the §4 8-name dispatch set, never produces a relay record, and is never aliased to a canonical relay ID.
Implementations: (a) `ManagedClient` (a credential-holding direct channel) and (b) the §2.3 worker-side capability client (the broker path).
m-2's mapping and m-9's consumer code are written against the interface, so a frontend is placement-agnostic: the retained MCP server and the native tool are thin skins over the same caller, including the re-render refresh path (which needs `Describe` — `mcp.go:209-225` is the live witness).

### 1.3 What the transport must NOT absorb (the m-2 boundary — F42/F69)

The following are m-2-owned FieldSpec/form semantics and land in an m-2-owned module, NOT in `internal/channel`:
- `SchemaFromForm` / `propertyForField` / `structuredDescription` / `skipFormField` (`schema.go:11-88`);
- `SubmitPayloadFromArguments` + `submitArguments` (`schema.go:90-129`);
- **the re-render vocabulary** — `submitNeedsReRender` / `containsReRender` / `reRenderResult` / `declaredPhaseTier` and the schema-refresh decision in `refreshSubmitSchema` (`mcp.go:209-276`): interpreting `{state, violations, form_digest, re-render}` is form-semantics interpretation, not transport.
  This is an explicit seam ruling of this contract: **the transport returns raw outcome bytes + an error class (and raw `DescriptionResponse` for `Describe`); every interpretation of those bytes in FieldSpec vocabulary is m-2's module** (m-2 confirms, F69).
  Without this ruling the re-render family would silently strand in each frontend and drift — the parity requirement (§1.4) then becomes unenforceable.

### 1.4 Frontends, parity, and the MCP conditions

- Both frontends — the retained MCP server (`cmd/frank-mcp`, foreign harnesses) and the Step-3 native tool (m-9's consumer) — become thin skins: frontend protocol handling only; transport = §1.2, mapping = m-2's module.
- **Parity tests (required, amendment §5 + §10 annex "native == MCP"):** shared conformance vectors drive both skins → equivalent stamped conductor calls, equivalent `Describe` bytes, and equivalent re-render behavior.
  The vectors live beside the m-2 mapping module; m-7 supplies the transport fake.
- **MCP off the critical path:** the native vertical passes first (§10 build order); the MCP skin ships in Step-3 only if it costs nothing from the critical path.
- **Seat-scoping condition (amendment §5, carried as a build-lane invariant):** one MCP server instance serves exactly one seat credential; NO shared global credential multiplexing caller identities.
- **Legacy custody posture:** `cmd/frank-mcp` migrates to the §2.12 hardened credential source in the same build lane; the `FRANK_CREDENTIAL` ordinary-environment sourcing path is REMOVED, not retained.

### 1.5 The retry contract (one linearizable model for every consumer; R2-F1 folded)

1. **Retryable classification (exact, closed):** a transport attempt may be retried ONLY on connection-loss/dial-class failures — `net.ErrClosed`, `io.EOF`, broken pipe, connection reset, closed-network-connection, dial/socket-absent failures.
   NEVER retried: an rpc response carrying `Error` (a SERVED application/protocol rejection — including `auth:*` and `frame:*` classes — is a terminal outcome, not a transport failure); caller context cancellation/deadline; frame-too-large.
2. **At most ONE retry per operation**, after one reconnect (close → re-dial → re-authenticate), and only through the **retry gate**: `ManagedClient` accepts a per-instance `RetryGate` hook invoked immediately before any second attempt.
   Frontend default: always-allow.
   Broker composition: the gate RE-ENTERS the §2.5 epoch fence — if the operation's admitted epoch is no longer current at retry time, the retry is refused and the operation returns `broker:stale-epoch` (for `relay.project`/`Describe`, and for `relay.read` per item 4's disposition rule) or `broker:unknown-outcome` (a submit whose first attempt may have committed) — the refused retry is recorded (§2.11).
   Both transport attempts of one operation are recorded under the one admitted operation (one admission, up to two sends — the §2.5 post-update claim is stated against ADMISSIONS and NEW sends through the gate, and the retry gate is exactly the re-entry that keeps it true).
3. **Single-flight connection replacement:** connection state carries a generation counter; concurrent failures collapse onto ONE redial/re-auth (later failures of the same connection generation join the in-progress replacement rather than starting another); the push reader re-attaches to the new connection under the same lock before the replacement is published — no detached-reader window.
4. **Side-effect posture per operation (exact; R2-F1):**
   - `relay.project` and `Describe` are read-only — a fenced-out or failed retry loses nothing (re-invoke).
   - **`relay.read` is CONDITIONALLY MUTATING** through the landed quarantine repair path: a checksum mismatch enqueues the record for quarantine and returns `checksum-mismatch` (`cmd/frank/main.go:463-481`); the serialized loop later moves the record to quarantine, rebuilds projections, and completes the incident (`engine/loop.go:89-95,111-128`).
     **No byte-equivalent retry behavior is claimed.** The retry/lost-response contract is: any subsequent read of the same record returns the AUTHORITATIVE CURRENT disposition — `checksum-mismatch` (repair pending) or `record-quarantined` with the incident identity (repair committed) — each a truthful terminal response; a caller that loses the first response learns the state from any later read (rediscovery), and a fenced-out retry returns `broker:stale-epoch` with nothing lost (the repair proceeds server-side regardless of the caller).
     **Why duplication is impossible (bound to the landed mechanism):** (a) the quarantine enqueue is a bounded drop-on-full channel (`loop.go:74-82`) — dropping is safe because any later read of the still-mismatched record re-enqueues (the repair is re-triggerable, never owed to one enqueue); (b) `QuarantineOne` executes ONLY inside the single serialized loop goroutine (`loop.go:89-95,122-128` — the same goroutine as commits; no concurrent store mutation); (c) `QuarantineOne` is idempotent by construction — record already moved ⇒ `(false, nil)` no-op (`store/quarantine.go:72-77`; `TestQuarantineOneIsIdempotent` is landed); (d) incident completion skips when the incident record already exists (`store/quarantine.go:106-107`) — duplicate enqueue can never mint a second incident or corrupt projections (projection rebuild rides the same serialized goroutine).
   - `relay.submit` is mutating: a retried submit is safe because conductor intake replays duplicates by content hash instead of re-executing accepted commands (proven by fixture FX-TB-11, not assumed); a submit whose retry is fenced out returns `broker:unknown-outcome` and its true outcome is recovered by rediscovery (the record is truth).

## §2 — The authenticated channel/BROKER contract (grill #3; F60/F64/F66; realized UNDER m-1 §2 per the §E rule)

### 2.1 Position and custody

The broker is the S-B custody + delegation point:
- it holds the LOGICAL m-9 seat's conductor credential — the only app-side holder;
- it dials ONE authenticated channel to the conductor via the §1.2 shared client (`DialAuthenticated`) — the broker is the **binding party** (m-1 §2.3);
- worker generations hold an **epoch-bound revocable USE capability** (§2.3) — never credential bytes, never a direct conductor connection;
- m-10 launches/supervises the broker but supplies NO credential bytes, NO credential path, and NO locator that itself authorizes resolution **[m-1-edge]**: m-10's launch names only the broker's config home; the operator-authored broker config inside that home names the §2.12 credential sink.

### 2.2 Placement — the GRILL decision (F67): the broker is its OWN supervised process

Resolved by the §G grill and ACCEPTED by both pair reviews: the broker runs as a separate supervised process, NOT as a module in the app main process.
Consequences pinned here:
- the secret-holding process set stays exactly `{m-8, broker}`;
- the m-10 rails become process-grain checkable: "m-10 receives no credential bytes" and "m-10 gains no conductor verb" are enforced by process boundary + interface absence (§2.8);
- worker↔broker and m-10↔broker interfaces are framed-message IPC in the private runtime directory (0700), close-on-exec, per m-10's §A.1/§A.2 discipline;
- **the broker SURVIVES an app-main crash** (no parent-death kill is configured for it): the seat channel stays bound while the fence fails closed (§2.4), and a replacement app-main ADOPTS the surviving broker (§2.10) — the grill's failure-isolation consequence, now realized rather than asserted (R2-F2).

### 2.3 The USE capability (m-1 §2.5 realized; the locator/capability split **[m-1-edge]**)

- **Type discipline (the two objects, never conflated):** an **opaque credential locator/reference** (custody orchestration; non-authorizing; m-1 §1.4 class) and the **worker USE capability** (channel delegation; deliberately AUTHORIZING) are different security objects with different claims.
  The capability is NOT an m-1 §1.4 locator and this contract does not satisfy any credential-reference requirement by pointing at it.
- **Form:** an opaque per-generation token minted by the broker at worker attach (§2.10), delivered ONLY on the worker's own broker connection — capability bytes never transit m-10.
- **Binding:** `{run_id, generation_id, turn_epoch at mint, broker_instance_nonce}` — AND connection-scoped: valid only on the broker connection it was minted on; presenting it elsewhere is rejected — re-attach mints fresh.
- **Honest authority/leak claim:** the capability carries no secret bytes, creates no principal, has no offline or credential-derived meaning, and appears nowhere in identity space (every record produced through it is stamped `FROM = the seat`).
  But it is NOT leak-inert: possession on an accessible current-epoch broker connection IS authorization for the §2.8 worker surface until the epoch advances, the connection closes, or the broker restarts.
  Its blast radius is bounded by the per-operation fence (§2.4), the connection scoping, and the 0700 runtime directory — and the amendment §2 same-user residual applies to it as to everything app-side.
- **Lifetime:** dies with the broker process (broker memory only); revoked in effect by any epoch advance and by connection close.

### 2.4 The F64 generation fence (per-operation + per-push)

Connect-time authorization is INSUFFICIENT (the live conductor authorizes once at `session/connect`, `server.go:277-309`, then serves epoch-blind, `:311-340`); the broker therefore gates EVERY operation:
- **On every `relay.submit`, `relay.project`, `relay.read` AND every `Describe`** arriving over a capability connection: compare the capability's bound epoch against the installed epoch state (§2.10); equal ⇒ forward on the seat channel; unequal ⇒ reject `broker:stale-epoch`, carried as a `fence_reject` event per §2.11's classes.
- **On every push delivery/forwarding:** a nudge arriving on the seat channel is forwarded ONLY to a capability connection whose bound epoch equals the installed epoch AT FORWARD TIME; otherwise it is suppressed and carried as `forward_suppressed` per §2.11's classes.
- **Epoch authority:** the broker's epoch view comes ONLY from m-10's `epoch_state` feed (§2.10); the broker never derives, guesses, or defaults an epoch.
- **Fail-closed floor:** no installed state — including the §2.10 step-5b snapshot-absent bootstrap/adoption branch — control session lost, or malformed update ⇒ **suspended**: reject operations (`broker:suspended`) and suppress forwards until state installs (5a) or the §2.5 reconciliation installs (5b); the suspension events are carried per §2.11's classes.
- **Monotonicity:** within a control session, only strictly increasing `state_seq` (and never a decreasing `turn_epoch`) is accepted; a regression is rejected and recorded.

### 2.5 Epoch-change linearization + in-flight disposition (F64's hard half)

- **One serialization point:** the broker serializes fence checks (including §1.5.2 retry-gate re-entries) and epoch-state installs through a single ordering point.
  An install takes effect atomically at that point; every operation gates at entry — and gates AGAIN at any retry — through the same point.
  There is no window in which two epochs are simultaneously current.
- **Operation identity (R3-F2):** every broker operation receives a stable **`operation_id` = `{broker_instance_nonce, op_seq}`** at admission (`op_seq` broker-instance-monotonic) and carries it through both transport attempts, retry fencing, crossing, completion, response correlation, and every §2.11 event row — concurrent operations are individually correlatable.
- **The epoch transition — ONE fenced, identified state machine (R3-F2 + R4-F2; the F64 "complete-or-reject, recorded" requirement made durable AND stable):**
  every transition carries a stable **`epoch_transition_id`** bound to `{run_id, from_epoch, proposed_epoch, state_seq}` through proposal, crossing rows, commit ack, install, abort, and recovery.
  States: **`PROPOSED → PREPARING → CROSSERS_DURABLE → INSTALLED`**, terminal **`ABORTED`**.
  1. **PROPOSED:** m-10 proposes E+1 (with the transition ID).
  2. **PREPARING (the barrier):** at the serialization point the broker enters PREPARING and FREEZES the crossing set — the bounded in-flight E operation IDs at that instant.
     While PREPARING: NO new E admission, NO retry send, NO stale push forward — anything arriving is rejected typed `broker:suspended` (bounded: the window is one durable commit round-trip; read-only ops re-invoke, a rejected-at-admission submit was never sent).
     The frozen set is returned as the **`crossing_set`** and CANNOT change for this transition ID: a duplicate proposal with the same ID returns the SAME frozen set; a DIFFERENT proposal is rejected until this transition reaches INSTALLED or durable ABORTED — E+1 can never install from a different crossing set under the same identity.
  3. **CROSSERS_DURABLE:** m-10 durably commits one `crossing-pending` row per frozen operation ID (keyed by transition ID + operation ID, CI-3) and acks.
  4. **INSTALLED:** only on that ack does the broker install E+1 and lift the barrier.
  **Completion while PREPARING/before install:** a frozen operation that completes before INSTALLED is durably resolved `completed-before-install` (a `cross_epoch_completion` disposition against its row) — never left as a false crossing row and never removed except by that recorded update or the abort resolution below.
  **Proposal is an epoch-authority event (R5-F1):** m-10 durably publishes E+1 BEFORE proposing (their §B.4 durable-then-visible supply; m-8 and m-10 already fence E) — so a proposal is never a tentative broker-local hint, and **once PROPOSED is received, the broker never again admits, sends, or forwards under E**, regardless of what happens to this transition attempt.
  A transition-ATTEMPT abort and epoch authority are separate things: aborting an attempt can never make E authoritative again.
  **Loss/timeout while PREPARING (no ack):** the broker CANNOT distinguish m-10-crashed-before-commit from committed-before-ack — so it never guesses: it stays in the PREPARING barrier, SUSPENDED-fenced (new arrivals reject `preparing`; frozen operations complete under the §2.11 outage rules — payloads withheld `broker:record-unavailable` where a response-coupled record is owed).
  There is no local abort and no barrier lift back to E.
  **The recovery MATRIX (R7-F1 — total over broker RECOGNITION of T × durable commit state, never ledger state alone; recognition is broker-verifiable: T's frozen set exists only in the instance that froze it, and instance continuity is carried by `broker_instance_nonce`):** on control (re)establishment, the pending `epoch_transition_id` T resolves by exactly one row —
  1. **Surviving instance, recognizes T, ledger `CROSSERS_DURABLE`** ⇒ the SAME transition resumes: the re-ack names the exact committed frozen set — which this instance froze and still holds — and the broker installs E+1 (idempotent).
  2. **Surviving instance, recognizes T, NO durable crossing rows** (ledger `PROPOSED`/`PREPARING`) ⇒ m-10 durably ABORTS the attempt — one transaction resolving EVERY row of that ID (completed ⇒ `completed-before-install`; still-in-flight ⇒ terminal `aborted-attempt`, each captured in the successor's frozen set — continuous coverage, no gap) — then proposes a FRESH ID for the STILL-CURRENT E+1, frozen over the surviving in-flight remnant.
  3. **Surviving instance, does NOT recognize T** (the proposal never arrived before the control loss) ⇒ the already-allocated T is proposed to this broker as its FIRST proposal — the ordinary phases run (this instance freezes T's set itself; nothing is resumed).
  4. **Already INSTALLED at the broker, `epoch_installed` lost** (ledger still `CROSSERS_DURABLE`) ⇒ the same-ID ack/query is idempotent; the broker re-delivers the durably keyed `epoch_installed` event and NEVER clears or re-freezes the already-installed epoch.
  5. **Fresh instance (the prior broker died — a NEW `broker_instance_nonce`; every bootstrap is this row):** every old frozen operation died with its connections ⇒ m-10 resolves T's rows `unknown-outcome` and durably aborts T — **a committed old set is NEVER installed by bare ledger ack into an instance that did not freeze it** — then proposes a NEW ID over this instance's (empty) snapshot.
  An aborted attempt is never reported as a cross-epoch install; no row of any attempt remains `crossing-pending` past its transition's terminal state; an `epoch_transition_id` survives recovery ONLY on rows 1/3/4 (the instance that froze — or will freeze — its set is at the other end).
  The mandatory record therefore EXISTS BEFORE the m-10-down window can open, every INSTALLED transition has exactly ONE durable frozen crossing set frozen by the installing instance, and E is never re-authorized after E+1's durable publication.
  On completion after install, the broker updates the keyed row (`completed{admitted_epoch, completed_epoch}` / `rejected`) via `cross_epoch_completion` and requires the durable ack BEFORE delivering the payload; ack unavailable ⇒ `broker:record-unavailable` — the pre-existing row still holds the crossing fact.
  The post-update claim, exactly: after PREPARING begins, no NEW admission and no NEW send (first attempt or retry) can occur under E; frozen operations may complete and deliver, each durably identified from before the install.
- **The replacement barrier (m-10's sequencing hook):** the frozen crossing set gives m-10 the exact in-flight identities (not merely a count); whether to drain before granting the new generation its capability is m-10's lifecycle policy.

### 2.6 Push path (amendment §6 semantics preserved exactly)

Conductor → broker seat channel (best-effort nudge, `PushTo` `server.go:167-180`) → broker forwards to the current-epoch capability connection (fence at forward time).
No new delivery guarantee is introduced anywhere on this path: best-effort + advisory end-to-end; the worker's durable rediscovery (catch-up `project`/`read` + `Describe` THROUGH the broker at startup/reconnect) is the recovery mechanism; the record is truth.
The broker adds no buffering promise beyond the existing client semantics (16-slot, drop-on-full).

### 2.7 The lifecycle event/effect matrix (m-1 §2.4 realized; the broker-side column **[m-1-edge]**)

| event | seat channel | epoch (`turn_epoch`) | capabilities | broker action | in-flight disposition |
|---|---|---|---|---|---|
| worker replacement | untouched | +1 (m-10 mints; broker installs) | prior-epoch tokens dead at the fence | install; report in-flight count (§2.5) | complete-and-deliver, recorded |
| broker crash/restart | drops (conductor unbinds on close — s6 §F) | untouched | ALL dead (broker memory + connection scope; m-1 MR2's token-death leg) | §2.10 bootstrap: listener → control session → the step-5 snapshot branch (a pending transition takes the §2.5 matrix row 5: old-ID durable abort, fresh ID over the empty set — never a bare-ledger resume) → re-dial/re-auth → accept re-attach with FRESH capability material | fail visibly to callers (`shim:connection-lost` class); rediscovery heals |
| **app-main crash (broker survives)** | **untouched — stays bound** | untouched (no writer to advance it) | untouched in form; USELESS in effect — the fence is suspended (§2.4: control loss ⇒ reject/suppress) | hold suspended; if a transition was PREPARING the barrier PERSISTS through the outage (E never resumes — §2.5 reconciliation by transition ID at adoption); accept the replacement app-main's adoption via the committed `control_generation` advance (§2.10); resume on installed/reconciled state | same-epoch admitted operations complete and deliver normally; operations needing a response-coupled record (a crossing, a fenced retry) withhold ⇒ `broker:record-unavailable` — their crossing rows pre-exist from the §2.5 install phase, so the mandatory record survives (§2.11) |
| conductor restart | drops; broker re-dials/re-binds (same credential) | untouched | untouched (broker alive; connections intact) | §1.5 reconnect (single-flight) | per §1.5.4 posture |
| re-mint (rotation) | force-closed by the conductor (`ForceCloseSeat`); broker re-reads the §2.12 sink and re-authenticates with the NEW credential (re-bind) | untouched | untouched | re-auth; recorded (§2.11) | per §1.5 (in-flight fails as connection-lost; retry re-enters the fence) |
| re-mint CONCURRENT with replacement | both effects, independently | +1 (replacement leg only) | prior-epoch dead (replacement leg only) | both actions; each leg recorded; either order yields exactly one re-bind + one epoch install | union of the two rows; no coupling (m-1 §2.2 two-counter law) |

### 2.8 Credential non-exposure + the closed interface surfaces

- S-B bytes exist ONLY in broker process memory and the §2.12 authorized sink; never in: capability tokens, worker/m-10 IPC frames, logs, typed errors, crash dumps (scrub-or-disable, m-1 §1.3.5), argv/env of any process, or any §1.2 error path.
- **The worker surface (closed, exact):** the three canonical relay calls + the typed `Describe` metadata/rediscovery operation + push receipt.
  Nothing else: no lifecycle authority, no epoch write, no credential operation.
- **The m-10 control surface (closed, exact):** the §2.10 control handshake + epoch-state install/query + attach/detach coordination + health + the §2.11 event/ack frames.
  `submit`/`project`/`read` AND `Describe` do not exist on this interface (FX-TB-9).

### 2.9 Acceptance targets (the §10 annex row, made concrete)

After a replacement (epoch E → E+1), an E-capability holder can: invoke NONE of `relay.submit`/`relay.project`/`relay.read` — and no `Describe` — through the broker (rejected `broker:stale-epoch`, carried as `fence_reject` per §2.11); receive/forward NO new push.
The full negative battery is §6.

### 2.10 Bootstrap, control sessions, adoption, and attach (R2-F2 folded)

- **The control endpoint (pinned lifecycle):** the BROKER owns a control LISTENER socket — `broker-control.sock`, mode 0600, in the private runtime directory (0700) — created at broker start before anything else is accepted (a dead predecessor's socket file is removed and replaced at bind, the conductor's own `Serve` pattern); removed on orderly exit.
  **m-10 is the CLIENT on the control channel** — this replaces any socketpair-inheritance assumption for the broker (the socketpair-before-spawn pattern in m-10's §A.1 remains theirs for m-9/m-8; the broker's control channel is dial-in BY CONSTRUCTION so that a replacement app-main can reach a surviving broker — the named cross-interface delta CI-1, §4).
- **The durable control generation (R3-F1: the handover authority, distinct from both counters):** m-10's store carries a single `broker_control` row — `{control_token, control_generation, minted_at}` — under its exclusive single-writer lock.
  `control_generation` is a strictly monotonic `uint64` advanced by ONE DURABLE TRANSACTION at every controller start: the initial spawn commits generation g before spawning; a REPLACEMENT app-main, before connecting, atomically advances g→g+1 and presents the COMMITTED g+1.
  `turn_epoch` and `state_seq` are untouched by this transition (control authority is a third counter, coupled to neither; the equal-`state_seq` crash case therefore progresses: the replacement's advance transaction manufactures the strictly-newer value the handover compares).
- **The controller lock/peer proof (R4-F1: what makes the presented generation broker-VERIFIABLE):** a bare integer from a token holder proves nothing — the broker never reads m-10's store — so the controller predicate is OS-BACKED:
  the controller holds an exclusive POSIX advisory write-lock (`fcntl F_SETLK`) on **`broker-control.lock`** in the private runtime directory, acquired BEFORE the generation-advance transaction and held continuously for its controller lifetime (kernel-released on death — the live singleton witness).
  **The traditional-record-lock caveat, pinned as a build obligation:** POSIX record locks are process-associated and drop if the process closes ANY descriptor on the same file — the build realizes "held continuously" with one stable lock-file inode and no same-process close of any descriptor on it (FX-TB-16 tests the invariant).
  On every control connection the broker: (1) reads the connected PEER's PID from the socket (`SO_PEERCRED` on Linux; `LOCAL_PEERPID` on Darwin — the platform floor, both mainstream POSIX/BSD facilities); (2) probes the lock with `fcntl F_GETLK` and requires the reported HOLDER PID to EQUAL the peer PID; (3) only then evaluates token match + strictly-greater `control_generation`.
  Any failure ⇒ reject, recorded `control_handover{outcome: rejected-lock | rejected-token | rejected-generation}`.
  **PID-reuse/TOCTOU, closed by ordering:** the probe runs AFTER the connection is accepted, so holder-PID == peer-PID means one live process is simultaneously both (two live processes cannot share a PID; a dead holder's lock is kernel-released and the probe then reports unlocked ⇒ reject) — no window in which a recycled PID passes.
  **Adoption ordering (pinned):** acquire the fcntl lock (waiting out or outliving any dying predecessor) → advance `control_generation` durably → connect → pass the peer/lock probe + token + generation.
  A stale app-main still holding the old token in memory now fails at step (2) — it does not hold the lock — regardless of what integer it fabricates; the fabrication race R3-F1 introduced the counter for is closed at OS-lock grain, not by trusting the reporter.
  Ceiling stated: a same-user process could itself acquire the lock and read the store/token — the F57 residual, unchanged; the token is lineage proof, the fcntl lock + peer probe are controller-ownership proof, `control_generation` is handover ordering; none claims cryptographic peer authentication.
- **The token (ONE handoff shape, pinned; an authorizing CONTROL capability — never a conductor credential, never a principal):** 32 bytes from the OS CSPRNG, encoded as 64 lowercase hex; stored in the `broker_control` row; handed to the broker at SPAWN via ONE INHERITED PIPE DESCRIPTOR (the broker reads exactly one line and closes the FD — no persistent copy outside the store, no file lifecycle, no argv/env, close-on-exec by construction; the non-injection census covers it as a control capability, not S-A/S-B).
  Rotation: a NEW token + generation advance is committed at every broker SPAWN (adoption presents the EXISTING token — only spawn rotates); the broker holds the token in memory only; a failed presentation (token mismatch or non-greater generation) closes the connection and is recorded as `control_handover{outcome: rejected-token | rejected-generation}`.
- **Broker start order (pinned):** (1) generate `broker_instance_nonce`; (2) read the token from the spawn pipe, close it; (3) create the control listener; (4) accept a control session (token + generation verified); (5) the SNAPSHOT BRANCH below; (6) load the credential (§2.12), dial, authenticate (bind); (7) accept worker attach ONLY once an installed state exists (attach verification requires it — §2.10 attach; until then attach returns `broker:attach-suspended`, fail-closed and transient).
  Until an installed state exists the fence is suspended (§2.4).
- **Step 5, the snapshot branch (F70 — the explicit branch, in the owner bytes):**
  - **(5a) supplied:** m-10 supplies an install-eligible durable snapshot `epoch_state{run_id, generation_id, turn_epoch, lease_state, state_seq}` ⇒ install it.
  - **(5b) WITHHELD:** when the current durable epoch carries a non-terminal or `ABORTED` transition, m-10 supplies NO install-eligible snapshot (their §B.3 recovery discipline) — the control session is nonetheless fully ESTABLISHED (token + generation + probe verified; the event/ack protocol live), the broker holds at the §2.4 suspended floor with NO installed state, and installation occurs ONLY through the §2.5 recovery MATRIX (recognition × commit state, never ledger state alone).
    At BOOTSTRAP this is by definition the fresh-instance row (matrix row 5): the old transition aborts durably (rows resolved) and a NEW ID proposes over this instance's empty snapshot — a committed old set never installs here; at ADOPTION (a surviving instance) rows 1–4 distinguish recognized-committed resume / recognized-uncommitted abort / unrecognized first-proposal / installed-replay; `ABORTED` in the ledger ⇒ the fresh-ID proposal directly.
    No snapshot, no inference, no default — the matrix's exact ack or fresh proposal is the only installer in this branch.
- **Adoption (replacement app-main, surviving broker):** the replacement acquires the `broker-control.lock` fcntl lock → advances `control_generation` durably → connects to the existing listener → passes the peer/lock probe → presents token + committed generation; the broker replaces any live session on the verified strictly-greater generation (fail-closed during handover, recorded as `control_handover{outcome: adopted}`), then takes the SAME step-5 snapshot branch: an install-eligible snapshot installs (equal `state_seq` is VALID here — the verified generation, not the state, authorizes the handover; §2.4 monotonicity resumes within the new session); a WITHHELD snapshot (pending non-terminal/`ABORTED` transition) leaves the adopted session live but the broker suspended, installing only via the §2.5 recovery matrix (5b) — a transition that was PREPARING at this surviving instance through the outage resolves by ITS `epoch_transition_id` under matrix rows 1/2/4, never inference; a transition this instance never received resolves by row 3 (first proposal).
  If the listener is absent or the connect fails, m-10 spawns a fresh broker (minting a NEW token + generation durably first).
  Reconnect deadline: a broker suspended past `CONTROL_REATTACH_DEADLINE` (compiled constant) stays suspended — it never self-authorizes; killing/respawning it is m-10 supervision policy.
- **One control session at a time:** exactly one live control session; replacement only by the peer/lock probe + token + strictly-greater committed `control_generation`; the old session closes at handover.
- **Worker attach (direct token delivery; the TYPED result taxonomy — D-3, closing the two-meanings gap m-9's half surfaced; R9-F1: SUSPENSION PRECEDENCE):** the worker connects to the broker's worker endpoint and presents `{run_id, generation_id, turn_epoch}` from its m-10 `assign`; attach returns EXACTLY ONE of three typed results, evaluated in THIS order:
  - **`broker:attach-suspended`** (TRANSIENT; evaluated FIRST — the broker's FULL attach-blocking suspension predicate takes PRECEDENCE over tuple evaluation): whenever the broker is suspended for ANY §2.4 cause — no installed state (the 5b branch, pre-install bootstrap), control-session loss (the installed state is RETAINED but NON-AUTHORIZING until a verified controller re-establishes the feed), or a malformed epoch update — **OR is inside the §2.5 PREPARING barrier** (installed E retained, control live, but E no longer authorizing after PROPOSED) — attach returns suspended, mints NO capability, and assigns NO fencing meaning to the tuple (it is never evaluated).
    The licensed worker behavior is a bounded hold-and-retry under m-10 supervision visibility; attach becomes possible when the broker leaves suspension (a state installs, verified control restores the authorizing state, or the transition completes — after which a now-stale tuple resolves TERMINAL below, honestly).
    **Totality sweep (R10-F1.5):** the attach-blocking states are EXHAUSTIVELY the §2.4 causes ∪ the §2.5 PREPARING barrier — every other broker state either has a live control session with an actively-authorizing installed state (the two branches below apply) or has no installed state at all (the no-state cause applies); PROPOSED-received-but-not-PREPARING is not an observable state (proposal processing and barrier entry are one atomic step at the serialization point), and a control-handover window is the control-lost cause until the new snapshot installs.
  - **`attach-ok`** — evaluated only when a verified control session is live, NO suspension barrier is active, AND an install-eligible current state is actively authorizing: the tuple equals the installed `epoch_state` ⇒ the broker mints the §2.3 capability and returns it ON THAT CONNECTION.
  - **`broker:attach-tuple-mismatch`** (TERMINAL for the presenting generation; the same live-control/no-barrier/actively-authorizing precondition): the presented tuple does not equal the actively-authorizing installed state — the presenting generation is FENCED; the licensed worker behavior is NO retry (re-presenting a fenced tuple is a stale generation hammering the fence; repeated presentations stay refused with no state change) and the disposition belongs to m-10 supervision.
  **No third fencing member is required (the routing's R7-F1 question, answered):** recognition-mismatch is a CONTROL-surface concept — attach compares presented-vs-installed only, and the pinned assign-after-install ordering (m-10 §B.4 step 5: publish → broker installs → THEN `assign`) means a legitimate generation can never race the install into a false mismatch.
  **Two malformed things, distinguished (R9-F1.5):** a malformed attach FRAME stays the existing protocol-error class (not a taxonomy member); a previously malformed EPOCH UPDATE is a §2.4 suspension cause and surfaces here as `attach-suspended`.
  This taxonomy TYPES the existing refusal — the F64 fence, the §2.3 binding, and the 5b branch semantics are unchanged.
  **Named cross-interface delta CI-2 (§4, m-10 confirms on exact bytes):** m-10's `assign` gains `generation_id` and the broker worker-endpoint name — `assign{run_id, turn_epoch, manifest_digest, generation_id, broker_worker_endpoint}` (their current §B.1 shape lacks both).
  A stale generation cannot re-attach: its tuple no longer equals the installed state; a replayed current tuple buys nothing — the capability it never received is connection-scoped and every operation gates per-verb anyway.
- **Control loss:** losing the control session suspends admission and forwarding immediately (§2.4) — an old feed cannot silently continue authorizing.

### 2.11 The recording protocol (R2-F3 folded: total, keyed, honest in every window)

- **Writer:** m-10's durable store — the broker keeps NO durable state and mints NO new output family (a broker-local durable spool would be a new state owner ⇒ the named route-back; deliberately not chosen).
- **Event identity (global, stable):** every event carries the idempotency key `{broker_instance_nonce, event_seq}` — `event_seq` a broker-instance-monotonic `uint64` from 1; the per-instance random nonce makes keys unaliasable across broker restarts and control handovers.
  The shared domain `op ∈ {submit, project, read, describe}` is defined once here; every operation-scoped event ALSO carries the §2.5 `operation_id` for cross-row correlation.
  **Counter encoding (L1, affirmed as part of the CLOSED family):** every trust-bearing counter in the event/ack family — `event_seq`, `op_seq`, every epoch field, `state_seq`, `control_generation`, and the transition/operation identity components — crosses JSON only as the canonical decimal uint64 STRING of m-10's §A.2 R4 rule (grammar `^(0|[1-9][0-9]*)$`, value < 2^64, numeric comparison), never as a JSON number.
- **The closed event table (within m-10's §A.2 frame envelope; enums exact and complete):**

| event type | required fields beyond the key + `run_id` | class |
|---|---|---|
| `crossing_set` | `epoch_transition_id, from_epoch, proposed_epoch, state_seq, operations: [ {operation_id, op, generation_id} … ]` (the FROZEN §2.5 snapshot; empty array valid; same transition ID ⇒ byte-identical set) | install-phase (durable BEFORE E+1 installs) |
| `cross_epoch_completion` | `epoch_transition_id, operation_id, generation_id, admitted_epoch, completed_epoch, op, disposition ∈ {completed, completed-before-install, rejected}` | response-coupled (updates the crossing row) |
| `transition_aborted` | `epoch_transition_id, reason ∈ {broker-restart, controller-recovery-uncommitted}` | uncoupled |
| `retry_fenced` | `operation_id, generation_id, admitted_epoch, current_epoch, op` | response-coupled |
| `fence_reject` | `operation_id, generation_id, presented_epoch, current_epoch (absent if suspended), op, reason ∈ {stale-epoch, suspended, preparing}` | uncoupled |
| `forward_suppressed` | `presented_epoch (of the target connection), current_epoch (absent if suspended), reason ∈ {stale-epoch, suspended, preparing}` | uncoupled |
| `epoch_installed` | `epoch_transition_id, generation_id, turn_epoch, state_seq, crossing_count` | uncoupled |
| `control_handover` | `control_generation, outcome ∈ {adopted, rejected-lock, rejected-token, rejected-generation}, state_seq (absent on rejection), dropped_event_count` | uncoupled |
| `reauth` | `reason ∈ {rotation-force-close, reconnect}` | uncoupled |
| `attach` / `detach` | `generation_id, turn_epoch`; attach additionally `outcome ∈ {ok, suspended, tuple-mismatch}` + `reason ∈ {no-installed-state, control-lost, malformed-update, preparing}` (REQUIRED iff outcome=suspended, ABSENT otherwise; the presented values recorded on refusal) | uncoupled |

  Fields not listed for a type are ABSENT (unknown fields ⇒ m-10 rejects the frame); no relay bodies, no credential-adjacent bytes anywhere.
- **Ack protocol:** m-10 durably commits the event row, then acks `broker_event_ack{broker_instance_nonce, event_seq}` — correlation by KEY, never by connection sequence; duplicate delivery of a committed key returns the SAME ack (dedup by the unique key).
  **Named cross-interface delta CI-3 (§4, m-10 confirms on exact bytes):** the `broker_events` row family (`UNIQUE(broker_instance_nonce, event_seq)`, same-ack-on-duplicate) PLUS the **crossing rows** — keyed `{epoch_transition_id, operation_id}`, states `crossing-pending → completed | completed-before-install | rejected | aborted-attempt | unknown-outcome`, written durably at the §2.5 CROSSERS_DURABLE phase, updated by `cross_epoch_completion`, duplicate-message-idempotent (same transition ID ⇒ same set, same ack), PLUS the **transition ledger** — the §2.5 state machine per `epoch_transition_id`, at most one non-terminal transition per run, **queryable by ID at control (re)establishment**, with the recovery answer evaluated JOINTLY with broker-instance continuity per the §2.5 matrix (a committed set resumes ONLY into the instance that froze it — `broker_instance_nonce` carries the recognition; ledger state alone never installs), the abort as ONE transaction resolving every row of its ID before any successor proposal, recovery resolution to `unknown-outcome` on broker loss (their §B.3 park discipline; never silent replay) — bounded by the broker's concurrent-operation constant (their current census has none of these).
- **Delivery classes + every no-writer window (total; "recorded" is RESERVED for durable commit):**
  - **Install-phase (`crossing_set`):** durably committed and acked BEFORE E+1 installs (§2.5) — the mandatory cross-epoch record pre-exists every exposure window by construction.
  - **Response-coupled events** (`cross_epoch_completion`, `retry_fenced`): durable ack BEFORE the response is delivered; ack unavailable (control down) ⇒ the payload is WITHHELD and the caller receives `broker:record-unavailable` — for reads/describe re-invoke; for submit the committed truth is recovered by rediscovery; the crossing row from the install phase still holds the durable crossing fact, so no crossing operation is ever unrecorded (R3-F2).
  - **Uncoupled events** (all others): the caller-visible rejection (`broker:suspended`/`broker:stale-epoch`) is delivered immediately; the EVENT is **pending-resend** — a bounded in-memory queue (compiled-constant depth), (re)sent with its unchanged key on session (re)establishment until acked, at-least-once + key-dedup; it becomes "recorded" only at durable commit.
    Overflow drops the OLDEST pending event and increments a drop counter reported in the next `control_handover` (`dropped_event_count`) — honest telemetry of loss, NOT a record of each dropped event.
    **The named residual:** a broker crash DURING a control outage loses the pending queue — the dual-failure window, bounded by the queue depth; the durable rows (crossing sets, installs, handovers, worker states) still bound what can have happened.
    Uncoupled pending-resend covers TELEMETRY only — nothing in the mandatory set (crossing identification + disposition) ever rides it.
    A best-effort E0 event never substitutes for this protocol (grading of what m-3 derives from the committed rows is m-3's; durability is m-10's store).

### 2.12 The credential source (the authorized sink **[m-1-edge]**)

- **The sink (one, named):** a broker-private, operator-authored credential file — provisioned from the operator submit reply / admin path (F-S6-M1-1 floor), named inside the operator-authored broker config, classified by m-1 rev2 §1.4a as S-B's ONLY resolution path (startup/re-auth) and subject to m-1's S-B census/rotation/crash-dump/backup/deletion obligations.
- **Ordinary-environment sourcing is REMOVED:** no S-B (or any seat credential) in env or argv, anywhere — the shared source has no env path (this retires `FRANK_CREDENTIAL`, §1.4).
- **Descriptor-safe loading (pinned):** open with `O_NOFOLLOW|O_CLOEXEC`; `fstat` the OPEN descriptor and verify: regular file, owner == process euid, mode exactly `0600`, size ≤ 4 KiB; read fully from that same descriptor; `fstat` again and verify identity (dev/inode) and size unchanged — no path re-open between check and read.
- **Grammar (exact):** the file contains exactly one credential token on one line — `token '\n'` or bare `token` at EOF; non-empty printable ASCII, no interior whitespace; anything else ⇒ reject, fail-visible.
- **Rotation/replacement semantics:** the operator replaces the file by atomic same-directory rename of a new 0600 file; the broker re-reads ONLY at bootstrap (§2.10) and at re-auth after a rotation force-close (§2.7) — never on a timer; a deleted/missing/failed file at those points is fail-visible (`shim:auth-failed` class), never a silent fallback.

## §3 — The conductor-identity PRODUCER contract (F68, realizing F65)

### 3.1 The two identity components and the process-bound proof

- **`conductor_build_digest`** — SHA-256 over the conductor service executable, encoded `sha256:<64 lowercase hex>`.
  **Self side (loaded-image grain):** at startup, before serving, the process opens its OWN loaded executable image — Linux: `open("/proc/self/exe")` (a descriptor to the loaded image, immune to path replacement); Darwin: the OS-reported executable path opened `O_NOFOLLOW`, with `fstat` identity (dev/inode) captured before and after hashing through that single descriptor — and records `{digest, exe_dev_inode}`.
  **Platform floor, stated:** on Linux the loaded-image binding is the kernel's own reference; on Darwin there is a residual exec-to-open window — named as this platform's floor, inside the amendment's same-user residual.
  **Observer side (the independent half):** the external observer (a) finds the serving PID from the OS socket table, (b) OS-verifies THAT process's loaded-image identity matches the deployed artifact file by dev/inode, (c) independently hashes that artifact file, (d) compares with the stamp's digest, and (e) joins stamp-to-instance via `{pid, instance_nonce}` + OS process-start-time consistency (start time ≤ `loaded_at`; no PID reuse since).
  The binding claim is the AGREEMENT of the two sides at loaded-image grain.
  **Ceiling, honest:** confusion-resistant provenance; a same-uid actor with ptrace/debug access defeats it (the amendment §2 residual); no attestation is claimed; a MAC/signature/code-sign identity is new machinery priced separately.
- **`governing_config_identity`** — the pair `{config_generation, composite_digest}`: the generation orders WHICH accepted config chain state is active (my `…-060542` confirmation — monotonic/never-reused by construction); the composite digest names the exact governing bytes (the phase-0 `ExpectedConfigDigest` walk).

### 3.2 The stamp — canonical schema, bytes, home, and write protocol (R2-F4 byte rule folded)

```
stamp_object := JCS (RFC 8785) object:
{ "stamp_schema": 1,
  "conductor_build_digest": "sha256:<64 lowercase hex>",
  "config_generation": "<canonical decimal uint64 STRING — grammar ^(0|[1-9][0-9]*)$, value < 2^64; numeric comparison over the decoded value>",
  "composite_digest": "sha256:<64 lowercase hex>",
  "member_markers": { "engine_version": <int>, "fieldspec_marker": <string>, "catalog_marker": <string> },
  "loaded_at": <RFC3339 UTC, second precision, trailing "Z">,
  "pid": <int>,
  "instance_nonce": "<32 lowercase hex — 16 bytes from the OS CSPRNG at startup>",
  "exe_dev_inode": "<decimal dev>:<decimal inode>" }

file_bytes := JCS(stamp_object) || 0x0A          # exactly one terminal LF; the LF is a FILE terminator, not part of the canonical payload
```
- **Counter encoding (L1/VP F73, the m-10 §A.2 R4 rule adopted):** `config_generation` is a full-domain trust-bearing counter and NEVER crosses JCS JSON as a number (JCS interoperable integers cap at 2^53−1, RFC 8785 §3.2.2.3) — canonical decimal string, grammar-checked, numerically compared.
  The remaining numeric fields stay JSON numbers under EXPLICITLY NARROWER domains (the L1 alternative, rationale stated): `stamp_schema`/`evidence_schema` are small fixed schema versions; `pid` is an OS process ID (< 2^31 on the pinned platforms); `member_markers.engine_version` is the small engine version integer — each < 2^31, far inside JCS interop; violation of any domain ⇒ invalid stamp (the read rule below).
- **Consumer read rule (fail-closed, byte-exact):** the file must end in exactly one `0x0A`; the consumer strips it, parses the prefix, and BYTE-COMPARES the prefix against `JCS(parsed object)` — any unknown field, missing field, wrong type, non-canonical encoding, missing/extra terminator ⇒ invalid ⇒ the consumer's deny floor.
- **Home + name (pinned):** `serve-stamp.v1.json`, beside the conductor's listening socket in the conductor's runtime directory (directory root = operator deployment config; app/observer-readable; NEVER delivered on any seat surface — I-PH; content is leak-inert).
- **Write protocol:** temp file in the same directory → `fsync(file)` → atomic `rename` → `fsync(directory)`.
- **Publication order:** written AFTER phase-0 composite-digest verification succeeds and BEFORE the listener socket accepts.
- **Failed startup + status model:** phase-0 failure writes nothing; a stamp is **current-active** only when instance-joined to a live serving process (`{pid, instance_nonce}` + start-time); an unjoined stamp is the **last-loaded record** — still true history, which is the conductor-down semantics m-10's freshness rule consumes.
- **Writer wording:** the stamp is emitted by the conductor's LIFECYCLE OUTPUT path — a serve-start artifact, not a governed-store record.

### 3.3 The relay-leg evidence reference (R2-F4 folded: a versioned closed object, exact bytes)

```
relay_leg_evidence := JCS (RFC 8785) object:
{ "evidence_schema": 1,
  "conductor_build_digest": "sha256:<64 lowercase hex>",
  "config_generation": "<canonical decimal uint64 STRING — the §3.2 counter-encoding rule verbatim>",
  "composite_digest": "sha256:<64 lowercase hex>",
  "instance_nonce": "<32 lowercase hex>",
  "relay_records": [ { "relay_id": <string: the conductor's committed relay-ID token verbatim — non-empty ASCII, no control chars, ≤ 256 bytes>,
                       "kind": "observe-as-send-E1" | "observe-as-send-E2" } … ] }

carried-as-file ⇒ the §3.2 byte rule applies verbatim (JCS || single 0x0A); embedded-as-object ⇒ bare JCS.
```
- **Array law (deterministic):** `relay_records` is sorted ascending by `relay_id` byte order; `relay_id` values are UNIQUE (a duplicate ⇒ reject); cardinality ≥ 1 with AT LEAST ONE record of kind `observe-as-send-E1` (the send leg of the tested exchange must be present); `kind` outside the two-value enum ⇒ reject; unknown/missing fields ⇒ reject.
- **Reference verification (by resolution, not grammar):** the observer verifies each `relay_id` by resolving it through the governed read surface (an operator/observer seat): the read must return a COMMITTED record that is an observe-as-send record of the named kind describing the tested exchange; unresolvable, wrong-kind, or exchange-mismatched references ⇒ reject.
- **Instance binding:** the four identity fields are copied from the stamp current-active during the tested leg; `instance_nonce` pins WHICH serve instance was tested; an identity/nonce mismatch against the observation-time stamp ⇒ non-applicable.
- **Applicability (the F65 split):** mutating the conductor binary or its governing config invalidates THIS leg's evidence (the next stamp will not match), and never the app/provider-vertical E3 — m-3 confirms this scope boundary from the evaluator side; neither absorbs nor omits the other's leg.

### 3.4 Bounds + the census obligation

No conductor protocol verb, store record/member, or seat surface is added — the stamp is a serve-start OUTPUT artifact.
It is a NEW conductor output family; its census/output-family row in my domain is OWED at the build lane with exactly the §3.2 bytes.

## §4 — Consumer obligations (what each confirms; F69 set + the named cross-interface deltas)

- **m-9:** consumes the §1.2 two-operation caller seam (including typed `Describe`) for the native tool and the §2.3/§2.10 capability + attach surface for the worker; confirms the worker path needs no credential bytes, no direct dial, and nothing beyond the §2.8 closed worker surface; confirms the §1.5.4 `relay.read` disposition contract fits the worker's rediscovery loop.
- **m-10 (including three NAMED deltas on exact bytes):** **CI-1** — the broker control endpoint is a broker-owned dial-in listener (m-10 connects; the socketpair-before-spawn pattern does not apply to the broker) with the §2.10 protocol IN FULL: the `broker_control` row `{control_token (64 lowercase hex), control_generation (uint64), minted_at}`, the durable generation-advance transaction at every controller start (spawn AND adoption), the **`broker-control.lock` exclusive fcntl lock acquired before the advance and held for the controller lifetime** (the broker's peer/lock probe verifies holder-PID == connected-peer-PID before any generation is evaluated), the spawn-pipe token handoff, rotation-on-spawn, and rejected-handover recording (`rejected-lock | rejected-token | rejected-generation`); **CI-2** — `assign` gains `generation_id` + `broker_worker_endpoint`; **CI-3** — the `broker_events` row family (`UNIQUE(broker_instance_nonce, event_seq)`, same-ack-on-duplicate) PLUS the crossing rows keyed `{epoch_transition_id, operation_id}` (states incl. `completed-before-install`) PLUS the transition ledger (the §2.5 `PROPOSED → PREPARING → CROSSERS_DURABLE → INSTALLED | ABORTED` machine per transition ID; recovery aborts durable before any new proposal).
  Plus: the §2.5 barrier ordering, the §2.11 ack ordering, the §2.10 step-5b withheld-snapshot branch as the receiving half of their recovery discipline (proven reciprocally, FX-TB-18), and the §2.8 control-surface absence (no verb, no describe, no credential bytes, no credential path).
- **m-1 [per §E]:** confirms §2 realizes their APPROVED semantics (`7c8b09a6…`) — the §2.3 locator/capability split (their §1.4a/§1.4b), the §2.12 sink (their §1.4a S-B resolution path), the §2.7 matrix (their token-death leg + the overlap row), and no coupling of the two counters.
- **m-2:** confirms the §1.3 seam ruling — the mapping module boundary (including the re-render vocabulary) is complete, `Describe`'s raw-return split leaves all interpretation theirs, and the parity-vector home works.
- **m-3 (the §3 edge):** confirms the F65 scope boundary — the relay-leg binding stays separate from the app/provider E3 vector; the §3.3 object is the exact half Master+VP byte-bind into the composite exit record.

## §5 — Boundary compliance (route-back check)

- The fence, capability, sessions, adoption, linearization, and recording are broker/app-side — no conductor gate, protocol verb, store record, member, or seat-surface change anywhere in §1–§2; `Describe` rides the EXISTING `tools/descriptions` method; the §1.5.4 read/quarantine analysis DESCRIBES landed conductor behavior and changes none of it.
- §3 adds one conductor serve-start OUTPUT artifact (required by F68); no protocol/store field; the census row is owed (§3.4).
- **No broker-local durable spool** (§2.11 chose the narrowed claim + pending-resend instead) and **no new identity authority** — the r2 review's named route-back triggers are not tripped.
- The F65 note stands: the conductor identity is bound separately in the exit-test record, never an app-release field.
- No locked m-7 text is reopened; the r13 config contract and the s11 v4 engine surface are consumed, not modified.

## §G — GRILL_LOCK (F67: broker placement — UNCHANGED from r1; accepted by both pair reviews)

```text
GRILL_LOCK_ID: step3-mvp-design-m7-broker-placement-grill
GRILL_REQUIRED: yes (supplement …-043459 F67 — supersedes the dispatch's "no")
GRILL_SOURCE:
- plan/design/audit relay read: step3-mvp-design-m7/DESIGN-orchestrator-planner-20260716-041630.md + …-043459.md; STEP-3-MVP-AMENDMENT.md r7 §§2/2b/6/7/10 (F57/F59/F60/F64/F66); m-1's contract aa90fa45… §§1-2
- code/docs inspected: frank@502e06c internal/channel/server.go (auth :277-309, verbs :311-340, push :167-215); cmd/frank-mcp/{mcp,main,errors}.go (custody + reconnect precedent)
- questions answered from codebase/docs: all eight F67 criteria (below)
- questions asked operator: none — the supplement delegates the choice to m-7 unless the winning answer alters a ratified topology or claim boundary; the selected answer is a §2b-NAMED allowed option and leaves the F57 claim unchanged, so no operator routing is triggered

Resolved decisions:
- Broker placement — OWN SUPERVISED PROCESS — criteria: (F57 claim boundary) own-process keeps the secret-holding set {m-8, broker}: two small scrub-or-disable processes; in-app-main would make the entire control plane + durable state store a secret-holding process (m-1 §1.3.5 scope explosion) and weaken "m-10 receives no credential bytes" from a process-grain property to intra-process code discipline; (m-1 semantics) the binding party is a dedicated custody process, symmetric with m-8's S-A custody — one pattern, one hardening battery shape for both secret classes; (m-10 rail) no-verb/no-bytes become interface-absence + process-boundary properties, testable at the §10 "m-10 is not a seat" grain; (lifecycle/recovery) m-10 supervises the broker like m-8/m-9; a broker crash is loud, identity-inert, and bounded (§2.7) — credential re-read from the operator 0600 file, re-bind, re-attach; (linearization) both placements admit the single-serialization-point design; own-process makes the epoch feed an explicit ordered message stream, which §2.5's monotonic re-seed rule handles; (push routing) identical hop count in both placements (an IPC hop exists either way); (failure isolation) worker crash touches nothing (F66); app-main crash does NOT take the seat channel down in own-process placement — and even when m-10 is down the fence fails closed (no epoch feed ⇒ reject), so the surviving channel grants no stale authority; (concentration precedent) the amendment's own F59 Option-A rejection refused to concentrate execution beside durable state + credential REFERENCES in app-main; placing credential BYTES there would be strictly worse — source: docs + code, decided by m-7 per the delegated authority
- In-flight disposition at epoch change — COMPLETE-AND-DELIVER, RECORDED {admitted_epoch, completed_epoch} — an admitted call carries invocation-time authority (the F59 ticket-consume temporal semantics); the record makes the crossing non-silent — source: docs (amendment §7 F64 "complete-or-reject, recorded"), decided by m-7

Rejected alternatives:
- Protected thread/module in the app main process — rejected: expands the secret-holding process set to the whole control plane; demotes the m-10 no-bytes rail to code discipline; concentrates S-B beside the durable state store + F59 tickets against the F59-rejection rationale; asymmetric with m-8's ratified custody shape; saves one supervised process at the cost of every claim-grain listed above
- Suppress-delivery in-flight disposition — rejected: manufactures a committed-but-unreported state that rediscovery must heal; adds a failure mode, adds no property (the stale worker is being torn down and cannot re-invoke)

Still operator-owned:
- none — no criterion produced a topology- or claim-boundary-altering answer

Design-lock impact:
- §2.2 pins own-process placement; §2.5 pins the disposition; both ride into the Master+VP first-stage interface-lock via this DESIGN record (no self-declared lock)
```

## §6 — Fixture obligations carried to the build lanes (m-7 grain)

FX-TB-1 stale-epoch rejection ×(3 verbs + Describe) · FX-TB-2 stale push never forwarded · FX-TB-3 `state_seq`/epoch regression rejected (in-session AND cross-session handover) · FX-TB-4 feed-absent/control-loss fail-closed (`broker:suspended`) · FX-TB-5 in-flight crossing recorded before delivery (`{admitted_epoch, completed_epoch}`) · FX-TB-6 capability dead across broker restart AND across connection replacement · FX-TB-7 rotation: force-close → re-auth with new credential, epochs untouched; PLUS the overlap leg — re-mint concurrent with replacement yields exactly one re-bind + one epoch install in either order · FX-TB-8 sentinel secret absent from tokens, IPC frames, `broker_event` rows, logs, errors · FX-TB-9 no-verb-AND-no-describe on the m-10 control interface (interface/compile grain) · FX-TB-10 parity vectors: MCP skin == native skin conductor calls + `Describe` bytes + re-render behavior · FX-TB-11 retry contract: retry only on connection-loss class; single-flight replacement with push-reader continuity; epoch-advance-between-attempts ×(3 verbs + Describe) ⇒ fenced, typed, recorded; submit content-hash replay proven · FX-TB-12 stamp: byte-exact rule (JCS prefix + single LF; unknown/missing/terminator variance ⇒ reject); absent-until-phase-0-passes; write-before-accept; crash-atomic publication; loaded-image join (pid + nonce + dev/inode + start-time) — swap-the-binary and stale-stamp-not-current negatives; self-digest == observer digest at loaded-image grain · FX-TB-13 census: capability/credential bytes never transit m-10; credential-file negatives (symlink, non-regular, wrong owner/mode, oversize, malformed/multi-line, replacement race, env/argv/FD census) · FX-TB-14 recording: response-coupled ack-before-deliver; m-10-down ⇒ `broker:record-unavailable`; uncoupled at-least-once with key-dedup — commit-before-ack-loss ⇒ duplicate resend ⇒ SAME ack; broker-restart nonce non-aliasing; control-loss rejection/suppression events committed after re-establishment; queue-overflow drop counted in `control_handover`; the dual-failure residual documented · FX-TB-15 read/quarantine (R2-F1): checksum mismatch with response loss BEFORE the enqueue (later read re-enqueues) and AFTER it; retry before quarantine commit (⇒ `checksum-mismatch`) and after (⇒ `record-quarantined`, incident identity present); duplicate enqueue ⇒ one quarantine move, one incident (idempotency bound to `QuarantineOne` + incident-exists skip); epoch advance between read attempts ⇒ fenced, repair proceeds server-side · FX-TB-16 adoption (R3-F1 + R4-F1): app-main crash immediately AFTER an acknowledged state install ⇒ replacement reads the EQUAL durable snapshot, acquires the fcntl lock, commits the `control_generation` advance, adopts successfully with NO epoch change; **the fabrication negative: a stale token holder WITHOUT the lock presents a fabricated larger generation ⇒ rejected at the peer/lock probe (`rejected-lock` recorded) while the real lock-holding replacement's committed generation is accepted**; raced stale peer loses; listener-absent ⇒ fresh spawn with a NEW token + generation; token handoff census (pipe-only, no argv/env/file copy); dead-holder probe ⇒ unlocked ⇒ reject; the record-lock lifetime invariant (a same-process close of another descriptor on the lock file must NOT drop the held lock in the build's realization); `assign`/attach parity on the CI-2 bytes · FX-TB-17 transition durability (R3-F2 + R4-F2 + R5-F1): a new E call racing in AFTER the snapshot freezes ⇒ rejected `preparing`, absent from the set; an E retry during PREPARING ⇒ retry-gate refuses; a frozen call completing before install ⇒ `completed-before-install`, never a false crossing row; lost crossing-set ack then re-proposal (same transition ID) ⇒ SAME frozen set, same ack; a CONFLICTING transition ID while one is non-terminal ⇒ rejected; broker crash after CROSSERS_DURABLE but before install ⇒ recovery resolves rows `unknown-outcome` + durable abort + fresh ID; **the two indistinguishable no-ack halves (R5-F1): m-10 crash BEFORE the crossing-set commit ⇒ the broker stays PREPARING-suspended (E rejects `preparing` throughout), reconciliation finds not-committed ⇒ durable abort resolving every row (`completed-before-install`/`aborted-attempt`) ⇒ fresh ID for the STILL-CURRENT E+1; m-10 crash AFTER commit but BEFORE ack ⇒ reconciliation finds committed ⇒ the SAME transition resumes with the exact committed set and installs E+1 — in NEITHER half does E resume, and an aborted attempt is never reported as an install**; every INSTALLED transition has exactly ONE durable frozen crossing set; concurrent crossings individually correlated by `operation_id` · FX-TB-19 the attach-result taxonomy (D-3 + R9-F1; reachability by LICENSED worker flows only): **the worker-reachable transient cut** — an assigned current-tuple worker attaches `attach-ok`; control-session loss suspends the broker (installed state RETAINED, non-authorizing); the SAME worker re-presents its otherwise-current tuple ⇒ `attach-suspended{reason: control-lost}`, NO capability minted; verified control restores ⇒ the same tuple ⇒ `attach-ok` (plus the malformed-update variant of the same cut); **suspension precedence** — during ANY §2.4 suspension a current tuple can never mint (the R9-F1 authority-regression negative); **the terminal cut** — a stale generation presents its old tuple against a live-and-authorizing successor state ⇒ `tuple-mismatch`, terminal, no capability, repeated presentations refused with no state change (the no-hammering negative); **the PREPARING cut (R10-F1)** — an assigned E worker re-presents its otherwise-current tuple during the §2.5 PREPARING barrier ⇒ `attach-suspended{reason: preparing}`, NO capability; after E+1 installs, another presentation of the old E tuple ⇒ terminal `attach-tuple-mismatch` (bounded transient handling resolving honestly into the terminal branch as authority advances); **the ordering leg** — assign-after-install means a fresh generation NEVER observes `tuple-mismatch` on its first attach; every outcome + suspended-reason recorded via the attach event · FX-TB-18 the RECIPROCAL transition-ID battery (F70 + R7-F1 — driven from BOTH sides, m-7 and m-10 asserting the SAME §2.5 recovery matrix): (a) ADOPTION, surviving instance, recognized T, committed ⇒ the same-ID `CROSSERS_DURABLE` ack is the ONLY installer (broker side proves install-on-exact-ack-only-into-the-freezing-instance; m-10 side proves withhold-until-terminal) — incl. the PREPARING-through-the-outage variant; (b) BOOTSTRAP (fresh instance) with a pending transition in ANY state incl. committed ⇒ matrix row 5: old-ID rows `unknown-outcome`, durable abort, fresh ID over the empty snapshot — **a committed old set is NEVER installed by bare ledger ack into a fresh instance**; (c) surviving instance, recognized T, NOT committed ⇒ durable abort ⇒ fresh-ID proposal ⇒ the ordinary §2.5 phases install; (d) surviving instance, UNRECOGNIZED T (proposal never arrived) ⇒ row 3: T proposed as this broker's FIRST proposal, frozen locally, ordinary phases; (e) already-installed/lost-`epoch_installed` ⇒ row 4: idempotent same-ID ack/query + durable re-delivery of the keyed event; the installed epoch is never cleared or re-frozen; (f) `ABORTED` at bootstrap/adoption ⇒ the fresh-ID proposal directly; (g) ID-continuity: one `epoch_transition_id` traceable byte-identical through proposal/frozen `crossing_set`/ledger row/crossing rows/ack/`epoch_installed` — asserted from both directions, SCOPED to the paths that legitimately retain the ID (matrix rows 1/3/4); rows 2/5 prove old-ID-abort + fresh-ID replacement instead; (h) the negative: no snapshot install, default, or inference ever occurs in the 5b branch (a fabricated/unsolicited snapshot during a pending transition is rejected).
