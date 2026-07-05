# s4 Slice-4 — DESIGN: the wire-up (the per-seat MCP shim · seat-lifecycle hardening · the §7 config-change record)

**DESIGN_DOC_ID:** `s4-slice-4-design`
**Owner:** s4-wire pair — authored by `s4-wire.planner`; design-challenger + formal DESIGN-REVIEW addressee = `s4-wire.implementer` (v2.8.5 lineage)
**Dispatch:** `.relays/s4/s4-wire-design/DESIGN-orchestrator-planner-20260705-013107.md` (twelve binding constraints; GRILL_REQUIRED: yes)
**Rev:** r2 (operator grill COMPLETE — GRILL_LOCK at §12; provisional-pending-guide sections still marked **[PG-Qn]** against the in-flight `s4-guide-q1` thread — guide deltas fold before lock per the dispatch grill-rule)
**Date:** 2026-07-05 · **Tier:** medium · **Evidence:** E1 (design; all code cites verified in the s4 audits at main@6987367, code = tag `s3-close`)

---

## 0. Frame + locked inputs + claim boundary

**The one job.** End the operator-as-transport: a real host session (Claude Code / Codex) files a relay through `submit()` over a per-seat MCP shim, and a second real session receives it via `project()`/`read` — no human copy-paste in the loop. Plus: live-session lifecycle hardening (second-connect, per-recipient wake, frame ceiling, custody), and the §7 config-change record (discharges `OI-S3-CONFIG-CHANGE`).

**Claim boundary, held in every sentence:** everything here is **tool-mediated confusion-resistance** (ARCHITECTURE §C4.3); D5 residual (same-uid shell-bearing process — confused or malicious) accepted and stated wherever an exclusivity or custody claim appears. **Every s4 E3 claim is "transport/provenance only"** — done-state and `record_integrity` remain `self_reported` until Step-2 observe. The shim adds **affordance, never authority**: the seat-facing verb set stays exactly `submit`/`project`/`read`.

**Locked contracts consumed (never reopened):** m-7 §2–§8 (loop, pivot, recovery, config, guardrail, attach/wake), §13 F11 (one-pivot-per-mutation; `config-change` is already in the locked class list :172); m-1 §4–§6, §13.3 (DI family, API, system-field contract, conductor-internal provenance, credential-lifecycle carry); m-2 describe-grade render + re-render bounce (S3-landed); ARCHITECTURE §C4.1/§C4.3 (I-PH). **Second-connect rule source of record = the s4-dispatch [VP-W1]** (locked m-1/m-7 are silent on concurrent connections per credential; the dispatch pre-constrains: one active channel per credential, reject-or-proven-dead-recovery only).

## 1. The bridge at a glance

```text
 host session (Claude Code / Codex — the operator designates the live seats)
   │  stdio, newline-delimited JSON-RPC 2.0 (real MCP)
   ▼
 [S] frank-mcp SHIM (one per session, one credential per shim — DI-2 across the bridge)
   │  config: {socket path, credential} — credential env-var-first (§7)
   │  dial + session/connect at start; MCP initialize/tools/list/tools/call ↔ frank frames;
   │  Form→JSON-Schema mapping (§3); drift → refreshed schema + list_changed (§3.3);
   │  every SELF-generated error scrubbed/typed (I-PH — §3.5); poll-first posture (§5.4)
   ▼  unix socket, frank's existing frame dialect (PRIVATE behind the shim — unchanged)
 [A] conductor channel layer — auth (m-1 binding) + per-credential ACTIVE-CHANNEL INDEX (§4)
 [C] serialized commit loop — unchanged pipeline; + the config_change mutation class (§6)
 [D] store — canonical records = truth; config/ members become DERIVED-from-canonical
     after the first config_change record (§6.2)
 [E] read/delivery — per-RECIPIENT nudge (§5, retires broadcast + global pending queue)
```

Everything conductor-side is a **promotion** of built substrate: `channel.Client` is the shim's conductor-side half (server.go:305-471); the S3 describe-grade render is whole server-side; the S1/S2 crash harness takes the new class; the owed mechanism carries the OI discharge.

## 2. D-1 — the shim process (`cmd/frank-mcp`)

- **One binary, stdio MCP server.** Launched by the host's MCP config (`command: frank-mcp`). One shim per session; one credential per shim; the shim never multiplexes seats (DI-2 preserved across the bridge).
- **Config (GRILLED G-4):** socket path via `FRANK_SOCKET` env or `-socket` flag; credential via `FRANK_CREDENTIAL` env — **the documented default** — with `-credential-file <path>` (0600) retained as the secondary path for hosts that cannot set per-server env; **never a bare CLI credential flag** (ps-visibility; the existing `-operator-submit -credential` stays admin-tooling and its visibility is documented, §7). Where a host stores the env value on disk in its own MCP config, the D5 custody note is stated verbatim in the shim README. **Per-config-entry consequence (operator-raised in grill):** a host's MCP `env` block belongs to the server ENTRY, not the session — every session loading that config launches a shim with the SAME credential, and the second one gets `auth:channel-active` (§4, the fence working). One seat = one credential = one config entry; the two blessed wiring patterns live in the ops doc (§7).
- **Startup:** dial socket → `session/connect` with the credential → on success, serve MCP; on failure, MCP `initialize` still answers (a host must get a coherent handshake) but every tool call returns a typed, scrubbed error (§3.5) — the shim never crash-loops the host.
- **Lifecycle:** socket drop ⇒ bounded-backoff reconnect (`session/connect` again — proven-dead recovery leg, §4); on reconnect the shim re-fetches the schema (digest may have moved) and emits `tools/list_changed`; mailbox truth means no replay is needed beyond the host calling `project` (§5.4). Shim exit ⇒ host restarts it per its own MCP supervision; re-attach with the same credential resolves to the same seat (m-7 §8.5, built).
- **Dialect pinning (audit F-W5):** the shim allocates JSON-RPC ids starting at 1 on the socket side and never emits id 0 (the conductor client-dialect treats 0 as the notification sentinel); MCP-side ids are the host's and are never forwarded.

## 3. D-2 — MCP surface: the rendered form IS the submit schema

### 3.1 Tools
`tools/list` returns exactly three tools — `submit`, `project`, `read` — never more (E2 floor; registry enumeration fixture). `project`: no arguments. `read`: `{relay_id: string}` (required). No tool, description, or schema carries a FROM/ROLE/identity input anywhere (probed at the gate).

### 3.2 Form→JSON-Schema mapping (submit's `inputSchema`)
Source = `tools/descriptions` (`DescriptionResponse{SubmitSchema, FormDigest}`) for the seat's **current context** (initial fetch: server defaults, SITREP/medium). Mapping, mechanical and honest:
- form field → property under `headers` (the fill surface m-2 declares); field `type` maps to JSON type; `options` → `enum` (byte-exact values — the enum-byte-exactness floor rides through unmodified); `default` → `default`; parent-picker options → enum of conductor-supplied candidate relay IDs (+ default).
- envelope fill surface = exactly `{to, cc?, dispatch_id}` + `body` — recipient/dispatch selection is sender-authored per m-1 §6; **no `from`, no `role`, no `delivery_state`, no `relay_id` property exists in the schema** (system_only fields are absent, not "read-only").
- **`form_digest` = a required `const` property** carrying the digest of exactly this rendered schema (**GRILLED G-6: the model copies it back**); that copy is the proof-of-current-render the digest exists to be (m-2 drift semantics). The shim **never synthesizes or substitutes a digest** — a shim-injected digest would claim the host saw a form it may not have seen (claim-dishonest; rejected alternative §13).
- The tool description carries the honesty line: "Files a governance relay (transport/provenance only — content claims are not verified by this tool)." Never "verifies", never "authorizes".

### 3.3 Drift + phase negotiation (the re-render loop, built server-side)
`Validate` re-renders for the candidate's **declared** `PHASE`/`CEREMONY_TIER` and compares the carried digest against that render (validate.go:20-25). The loop this licenses, with zero new server surface:
1. Host fills the current schema and calls `submit`.
2. If the host's declared PHASE/TIER matches the schema served, the digest matches → normal pipeline.
3. Any digest mismatch (phase switch, tier switch, or a real config/candidate drift) → typed `re-render` bounce. The shim, on seeing violation class `re-render`: re-fetches `tools/descriptions` **for the phase/tier the host declared**, rebuilds the schema, emits `notifications/tools/list_changed`, and returns the bounce as a structured non-fatal tool result: `{state: "rejected", violations: [{field: "form_digest", class: "re-render", hint: "form refreshed — re-read the submit tool schema and re-submit"}]}`. The host model re-reads and re-submits — the [VP-W4] "superseded form bounces re-render; re-rendered form succeeds" leg is this loop verbatim.
4. The shim MAY pre-fetch the declared phase's schema on a phase-mismatch **before** submitting (saves one bounce round-trip) but still never carries a digest for a schema it hasn't served to the host — the submit goes out with the digest the host copied; the pre-fetch only accelerates step 3.

### 3.4 Tool-call results
MCP `tools/call` result = `{content: [{type: "text", text: <outcome JSON>}], isError: <transport-error only>}`. Outcome JSON is the conductor's own (state/relay_id/intake_id/violations) — byte-faithful, path-free already (audit P10). `isError` is reserved for transport/auth failures; a `rejected`/`held` outcome is a NORMAL result (the model must read it, not see a tool crash).

### 3.5 The shim's own surfaces are I-PH surfaces [VP-W3]
The shim knows the socket path and the credential. **Neither may reach any MCP-visible byte.** Every self-generated error is mapped to a typed, fixed-string class before surfacing: `shim:conductor-unreachable`, `shim:auth-failed`, `shim:connection-lost`, `shim:frame-too-large`, `shim:protocol-error`. Go's default dial error (`dial unix /path/…: connect: …`) is scrubbed at the single error-translation chokepoint (one function; every MCP-bound error passes through it; fixtured). Raw detail goes ONLY to the shim's stderr log (host-operator-facing, never model-facing) — and even stderr never prints the credential. The exit-gate I-PH matrix covers all seven bridge classes: tools/list descriptions · input schemas · tool-call results · notifications/poll hints · reconnect errors · credential-failure errors · shim diagnostics.

## 4. D-3 — second-connect: the per-credential active-channel index [VP-W1; source of record = s4-dispatch]

Engine-internal, no new seat-facing surface, **no binding-table shape change** (a needed shape change = hard stop + escalate):
- `serverConn` retains, post-`Resolve`, the SHA-256 of the presented credential (never the raw value — it lives only in the binding table and the wire read).
- `Server` gains `active map[credHash]*serverConn`, maintained under the existing `s.mu`: set on successful auth; deleted in the conn's `run()` defer (kernel close-detection — the proven-dead signal).
- **Reject leg:** `session/connect` presenting a credential whose hash is live in `active` → typed, path-free `auth:channel-active`; the existing channel is untouched. Any number of *distinct-credential* channels remain fine.
- **Recovery leg:** prior conn closed (host death, shim kill, socket drop) ⇒ its defer removed the entry ⇒ the next `session/connect` binds normally. This is reconnect — no new mechanism, no grace timer.
- **Wedged-but-alive host (GRILLED G-3, operator-accepted):** NOT provably dead (socket still open) ⇒ reject regime holds; the documented operator remedy is "kill the host session/shim; the kernel close frees the channel". No heartbeat/ping this slice, no grace timer, no supersede (a supersede would also be a session-hijack affordance — whoever holds the credential could yank a live seat). Dialect untouched; the deferral confirm rides guide Q5.
- **Explicitly out (escalate, amendment path):** live supersede ("new connect wins"), credential rotation, re-mint-supersedes. Nothing in this design gives a second connect any effect on the first beyond the typed reject it receives.

## 5. D-4 — per-recipient wake; retiring broadcast + the global pending queue **[PG-Q1]**

The locked shape is m-7 §8.3: delivery = one write onto **the recipient's** pipe. Today's code broadcasts one global frame to every client and replays a never-cleared pending queue to every future auth (audit probes (a)/(b), live-confirmed). Design:

- **Wake events ride the commit's own delivery intents.** The commit already materializes per-recipient mailbox entries; the loop's post-commit hook emits `(seat, relay_id)` wake events for exactly the committed record's recipients. `Server` gains a seat→active-conns view (the §4 index already knows seat identity via the auth meta); `PushTo(seat, frame)` writes `{"kind":"delivery-nudge","relay_id":"…"}` to that seat's connection(s) only. A seat's own relay_id on its own channel — no cross-seat metadata (the audit's side-channel note dies here).
- **Offline recipient:** no connection ⇒ no queue in the channel layer at all. The mailbox IS the queue (store truth). On that seat's next auth, the server checks pending-delivery state **for that seat only** and sends one `{"kind":"recovery-nudge"}` on that connection. `PendingDeliverySeats()` (store.go:175) gains a per-seat form (or is filtered at the call site — implementer's choice at PLAN grain).
- **The global `pending` slice and `Push`/`broadcast` are RETIRED from the serving path** (audit F-W4: retire, don't patch). `QueuePush`/`broadcast` survive only if fixtures need them; the production assembly stops calling them.
- **Startup recovery-nudge:** replaced by the same per-seat on-auth check — the process-start special case disappears (a seat that connects after a conductor restart gets its own recovery nudge exactly like any reconnect).
- **Nudges are best-effort by design** (client buffer can drop, hosts may hide custom notifications — F-W2): the shim documents and implements **poll-first**: `project` at turn start is the reliable path; nudges only accelerate it. The shim surfaces a nudge as a host-visible MCP notification where the host supports it, else drops it silently and relies on the poll hint in the tool descriptions.

## 6. D-5/D-6 — the §7 config-change record **[PG-Q2/Q3; GRILL G-1/G-2]**

### 6.1 Shape (GRILLED G-1/G-2 — operator-resolved; guide Q2/Q3 confirms mechanics only)
- `record_kind: config_change` (token + provenance ride the m-1 fidelity packet — ruling condition 4; no other new record_kind).
- **Authorship = operator-channel `submit` (GRILLED G-2):** the operator seat files a `config_change` relay through the normal pipeline — form-validated, operator-scoped (the registry restricts the record_kind value to the operator seat's rendered options; a non-operator submit carrying it bounces — fixtured), committed via the one pivot. **The accepted relay IS the operator-authorized digest-change record** — the authorization record is the mechanism working, first-class in the store. The admin-time `system`-stamped alternative was grill-rejected: it puts the human authorization outside the store.
- **Payload embeds the new member bytes (GRILLED G-1):** headers carry `{member: fieldspec|engine, new_digest: <top-level digest after the change>}`; body carries the full new member content (registry.json ≈ 10 KiB today — comfortably inside the grilled frame bound, §6.4). Store-is-truth: config history is auditable append-only from records alone; recovery/phase-0 re-materializes `config/` from the chain after any crash or hand-edit (self-healing — the operator-facing property that decided the grill row); `config/` becomes a derived projection like INDEX. The digest-only alternative (files changed out-of-band, record merely attests) was grill-rejected: two truths, store-down failure mode, history not reconstructable from the store.

### 6.2 Mechanics under no-hot-reload (restart-boundary semantics)
- Commit-time (running under the OLD config): validate → **one canonical pivot** (the config_change record; F11 one-pivot holds — exactly one rename) → **derived step:** materialize the embedded member bytes to `<root>/config/…` via the existing atomic write, in the same loop iteration, after the pivot (config/ becomes a **derived projection of the latest accepted config_change**, exactly like INDEX/mailboxes — canonical wins, unconditionally).
- The RUNNING process keeps serving under the old pinned config (no reload — locked). The new registry takes effect at the next restart. Consequence stated plainly on every claim surface: **a config change is effective-at-restart**; the [VP-W4] drift leg (old rendered form bounces `re-render`; re-rendered form succeeds) is exercised across that restart.
- **Phase-0 learns the chain:** expected digest = genesis's `config_digest` superseded in commit order by each accepted `config_change`'s `new_digest` (the "genesis chain" of m-7 §7 :109). Loaded `config/` digest ≠ expected ⇒ phase-0 **re-materializes** the members from the chain's embedded bytes (idempotent redo — ARIES property), recomputes, and only a persisting mismatch fail-closes to diagnostics (the existing posture). `ValidateGenesis`'s static compare (genesis.go:104-118) becomes the chain walk; a fresh store (no config_change records) degenerates to today's exact behavior — zero regression on S1–S3 stores.
- **Crash legs (the applicability row `config-change`):** kill between pivot and file materialization ⇒ recovery re-materializes (record = truth); kill mid-member-write ⇒ atomic file write discipline + redo; kill before pivot ⇒ nothing happened (intake re-enqueue). `f11Classes()` gains the class; every syscall boundary swept by the existing child-SIGKILL harness; **exactly one canonical rename per config change** — the derived file writes are not canonical renames.
- **Never re-genesis:** the round-trip runs on the s4 team's live store AFTER it has real records (existing store by construction); `store.Init` is untouched.

### 6.3 OI-S3-CONFIG-CHANGE discharge
After the live round-trip: the operator authors the `owed_disposition` for `OI-S3-CONFIG-CHANGE` through the operator channel on the real store (the S2 worked pattern), citing the config_change relay + the gate evidence; open owed set = empty at the exit gate. The frozen s2 store's upgrade is NOT a gate leg (optional operator call, out of gate scope).

### 6.4 Frame-transport bound rides with this **(GRILLED G-5; [PG-Q4] guide delta folds if any)**
Raise both scanner buffers to **1 MiB**, sourced from the engine config (additive member `max_frame_bytes`; absent = 1 MiB default, so existing stores' digests are untouched). Oversize frame ⇒ typed, path-free refusal (`frame-too-large`) written before close where the direction permits; never silent death (audit F-W1). No chunking this slice. **Grill basis (measured, 2026-07-05):** 711 relays across every trail (517 master + 194 frank s1–s4) — median ≈ 5 KB, p99 ≈ 20–25 KB, all-time max 33 KB; no relay has ever exceeded 64 KB, but the tail grows each slice and ecosystem docs already reach 93 KB — 1 MiB = ~30× max-ever headroom and swallows doc-sized bodies whole.

## 7. D-7 — custody posture + ops surface + usage posture (docs-grade, claim-honest)

- **Custody, stated honestly everywhere credentials appear** (m-1 §13.3 carry; D5): mint prints to stdout (capture discipline documented); `-operator-submit -credential` is ps-visible (admin tooling, short-lived — documented, not fixed this slice); binding table plaintext-at-rest 0600/0700 (D5-accepted, stated); shim credential env-var-first (§2). **No in-band rotation/revocation exists: a compromised credential means stopping the conductor and admin-time surgery.** Confusion-resistant, not theft-proof — verbatim on every custody surface.
- **Ops surface doc** (`docs/ops.md` or per PLAN): start (`-init` → mint → serve), stop (signal; durable FIFO + recovery make it safe), status (socket liveness + read-only diagnostics states genesis-missing/digest-mismatch), **short-socket-path rule** (darwin AF_UNix ≈104-byte cap; the default `<root>/frank.sock` join violates it for deep roots — convention: `/tmp/<team>.sock`-class paths; pre-flight length check as PLAN hardening), team-store conventions (absolute store root; the store is the governance domain), minting workflow end-to-end (who runs mint, where the credential goes, per-host wiring for Claude Code and Codex MCP configs). **Per-seat wiring (grill G-4 rider):** a host MCP `env` block is per SERVER ENTRY, not per session — so the doc pins **one seat = one credential = one config entry**, with two blessed patterns: (1) one config entry per seat in separate scopes/working dirs (matches how the team already runs seats as separate sessions); (2) `${VAR}` indirection where the host supports env expansion — the config carries `"FRANK_CREDENTIAL": "${FRANK_CREDENTIAL}"` and the secret arrives per-launch from the shell. Host support pinned per host in the doc; a same-credential double-launch gets the §4 typed reject (the fence working, documented as such).
- **Usage posture (minimal, chartered):** the store IS the usage record; document reading it (records/, INDEX projection, `project`/`read`); at most a trivial read-only stat over `project`. Anything aggregating = s5 OUT.
- **Honesty sweep [VP-W2]:** the "transport/provenance only; done-state and `record_integrity` remain `self_reported` until Step-2 observe" line appears on: this doc (here), shim README + tool descriptions (§3.2), ops doc, the exit-gate evidence record, the §7 authorization-record wording (a provenance claim, never an integrity claim). The root README fresh-store sentence goes stale when §6 lands — **PLAN-time fence ASK** (S1 ASK-1 precedent), pre-flagged.

## 8. Boundary contracts (per surface)

**Shim ↔ host (the new boundary):**
Writes: MCP frames (tools list/results/notifications) — affordance only. Reads: host tool calls. Target entity: none (no store writes shim-side). Downstream consumer: the host session's model. Contract: JSON-RPC 2.0 MCP; three tools; schema-as-form with const digest; typed scrubbed errors. Proof: E2 fake-host golden fixtures + E3 live gate. No-consumer action: n/a (the host is the consumer by construction).

**Shim ↔ conductor:** existing frame dialect, unchanged; `channel.Client` promoted. Contract: `session/connect`/`tools/*`/`notifications/nudge`; frame bound §6.4. Proof: existing channel fixtures + new second-connect/frame-bound legs.

**§7 record ↔ store (engine boundary, m-7-guided):**
Writes: one canonical `config_change` record (pivot) + derived `config/` members + INDEX/mailbox projections. Reads: genesis chain (phase-0), operator seat binding. Target entity: the pinned-config lineage. Downstream consumer: phase-0 + every future render (new digest). Contract: §6.1 headers/body; byte-exact enum untouched. Proof: crash legs + applicability row + chain-walk fixtures + the live [VP-W4] round-trip.

**Lifecycle ↔ channel layer:** active-index (§4) + per-recipient wake (§5): engine-internal state keyed off existing auth metadata; no new verb, no binding-table shape change, no cross-seat bytes. Proof: reject/recover legs, recipient-only nudge legs, offline-reconnect leg.

## 9. Fixture plan (E2 floors + the gate's live legs)

E2 (conductor-registry/battery): S4-MCP1 shim dialect goldens (initialize/tools-list/call shapes, id-0 pin) · S4-SCH1 Form→JSON-Schema mapping incl. system-field ABSENCE + const-digest + enum byte-exactness · S4-RR1 drift loop (phase-switch bounce → refreshed schema → success) · S4-SC1/SC2 second-connect reject (typed `auth:channel-active`) / proven-dead recovery · S4-NG1..3 recipient-only nudge, non-recipient silence, offline-reconnect recovery-nudge (per-seat) · S4-FR1 frame-bound typed refusal both directions · S4-IPH1..7 the seven bridge-surface classes each path-clean (incl. shim-diagnostic scrub with a real dial failure) · S4-C7-1..n config-change: operator-only refusal, chain walk, re-materialization after each crash leg, applicability row, one-pivot assertion, restart re-render leg · E2 floors: full battery, enum grep, three-tool enumeration.
E3 (gate-run, operator-designated seats): the live two-session relay (A=Claude Code, B=Codex ideally) · live adversarial legs (no/bad credential, second-connect, forged submit, I-PH probes over live surfaces) · crash/liveness legs (kill frank mid-delivery; kill shim; offline-seat nudge) · the §7 round-trip + OI discharge on the real store. Every E3 record carries the transport-only qualifier.

## 10. What this design does NOT do (OUT, restated)

Consumer schema content (s5) · observe/evidence (Step-2) · routing (Step-3) · TUI/email UX (Step-4) · federation (zero pre-work) · external send (outbox dormant) · steer/interrupt beyond host-native · authority replacement (gates still park for the human — transport only) · in-band rotation/supersede (§4) · socket-dialect rewrite toward MCP (the composite realization stands, pending guide Q6 confirm).

## 11. Risks + honest residuals

D5 (credential theft / direct store write by a same-uid process — accepted, stated); wedged-host seat blockage (reject regime; operator remedy documented); host variance in custom-notification surfacing (poll-first is the reliable path by design); the effective-at-restart config model (stated on every §7 surface — not a limitation apology, the locked no-hot-reload semantics); MCP protocol-version drift across hosts (the shim pins the version it implements and refuses unknown-major politely — typed, scrubbed).

## 12. GRILL_LOCK (grill run 2026-07-05, operator interactive, one question at a time; COMPLETE)

```text
GRILL_LOCK_ID: s4-grill-s4-wire
GRILL_REQUIRED: yes
GRILL_SOURCE:
- plan/design/audit relay read: s4-wire-design dispatch (…-013107); the reconciled paired audits
  (AUDIT-implementer-…-012253 + AUDIT-planner-…-013000 + ledger entry 2026-07-05); the s4-dispatch
  [VP-W1..W4]; the s3-scope-q1 ruling; m-7 §7/§8/§13, m-1 §4/§5/§6/§13.3, ARCHITECTURE §C4.1/§C4.3
- code/docs inspected: server.go / binding.go / main.go / render.go / validate.go / genesis.go /
  recover.go / f11_test.go (audit-verified cites at main@6987367); the live-probe evidence
  (scratch store, P1–P10); relay-size corpus measured for G-5 (711 relays, both trails)
- questions answered from codebase: digest-validate semantics (validate.go:20-25 — re-render is
  per declared PHASE/TIER; licenses the §3.3 drift loop); registry size ≈10 KiB (G-1 feasibility);
  proven-dead = kernel close ⇒ conn leaves s.clients (server.go:181-201); max relay ever = 33 KB
- questions asked operator: G-1..G-6 below (each with recommendation + alternatives; G-1 explained
  twice at operator request — app-level consequence framing decided it; G-5 answered only after
  the operator directed a corpus measurement)

Resolved decisions:
- G-1 §7 payload — EMBED the new member bytes in the record body — store-is-truth: auditable
  config history from records alone; self-healing config/ (re-materialized from the chain);
  one-step operator workflow — source operator ("sure lets do A")
- G-2 §7 authorship — operator-channel submit through the normal pipeline; the accepted relay IS
  the authorization record — source operator
- G-3 second-connect — reject-active (auth:channel-active) + proven-dead recovery ONLY; wedged
  host = operator kills the host session; NO heartbeat/grace/supersede this slice — source
  operator (consequence explicitly accepted)
- G-4 shim custody — FRANK_CREDENTIAL env var = the documented default; 0600 credential-file
  retained as secondary for hosts without per-server env; never a bare CLI flag; D5 note verbatim
  on every custody surface. RIDER (operator question in-grill): host MCP env is per CONFIG ENTRY,
  not per session ⇒ one seat = one credential = one config entry; two blessed wiring patterns
  (per-seat config scopes; ${VAR} indirection) pinned in the ops doc — source operator
- G-5 frame bound — 1 MiB via additive engine-config member (absent = default; existing stores
  untouched) + typed path-free frame-too-large refusal; no chunking — source operator, on the
  measured corpus (median 5 KB / p99 25 KB / max-ever 33 KB / ecosystem docs 93 KB)
- G-6 digest carrier — required const field in the schema; the MODEL copies it (proof-of-read);
  the shim never synthesizes a digest — source operator

Rejected alternatives:
- digest-only §7 attestation — two truths; store-down failure mode; history not reconstructable
- hybrid per-member §7 shapes — two code paths, saves nothing
- system-stamped admin-time config change — authorization would live outside the store
- heartbeat/ping now — dialect change; not needed for the gate; deferral confirm rides guide Q5
- supersede-on-new-connect — locked-contract touch AND a session-hijack affordance
- env-var-only / file-only custody — locks out or under-serves real hosts
- keep-64KB / chunked framing — tail growth + doc-sized bodies vs needless protocol redesign
- shim-injected digest — proves only the shim saw the form; defeats proof-of-current-render

Still operator-owned:
- the live-test seat designation (which two real sessions play A and B) — gate-time, per the
  s4-dispatch operator-judgment items
- the actual §7 authorization at the gate (the mechanism IS the operator's authorization)
- the frozen s2-store upgrade (optional first customer of §7 — NOT a gate leg)

Design-lock impact:
- §2 (shim config/custody + per-entry rider), §3.2 (const digest), §4 (reject regime), §6.1
  (embed + operator authorship), §6.4 (1 MiB bound), §7 (wiring patterns + custody wording)
  carry these decisions; DESIGN_LOCK_ID must reference GRILL_LOCK_ID s4-grill-s4-wire;
  [PG-Q1..Q6] guide deltas fold before lock per the dispatch grill-rule (fence: guide answers
  enter as resolved rows, never re-asked; no c1–c6 reopen; amendments escalate)
```

## 13. Rejected alternatives (running log)

Socket-dialect rewrite to native MCP (OUT by dispatch; composite is sanctioned) · shim-side digest synthesis (claim-dishonest — defeats the proof-of-current-render) · supersede-on-new-connect (locked-contract touch; also a session-hijack affordance) · heartbeat/ping this slice (dialect change; wedged-host detection deferred with documented remedy) · chunked framing (complexity; bound + typed refusal suffices at today's payload sizes) · per-phase `oneOf` mega-schema (schema bloat; the drift loop (§3.3) is the honest negotiation at zero new surface) · patching the global pending queue (retired instead — F-W4) · a second store for usage data (the store IS the record; s5 consumes).

## 14. Revision log

- **r1** (2026-07-05): provisional draft against the twelve dispatch constraints; [PG-Q1..Q6] sections await `s4-guide-q1`; grill agenda staged (§12).
- **r2** (2026-07-05): operator grill COMPLETE — GRILL_LOCK `s4-grill-s4-wire` folded at §12; G-1/G-2 (§6.1), G-3 (§4), G-4 + per-entry rider (§2, §7), G-5 with measured basis (§6.4), G-6 (§3.2) landed in place. Still NOT lockable: [PG-Q1..Q6] guide deltas outstanding (`s4-guide-q1` unanswered at r2 time) — they fold as r3 before any DESIGN_LOCK_ID, per the dispatch grill-rule.
- **r3 queue** (recorded 2026-07-05, not yet folded): (a) the `s4-guide-q1` answer deltas; (b) a §7 seat-occupancy-model paragraph (operator G-4 follow-up): seats are durable identities sessions OCCUPY at launch via credential; s4 automates the transport, not the assignment — roster minted admin-time, per-launch `${VAR}` wiring in the shared team cwd, one live occupant per seat (G-3); conductor-spawned sessions with injected credentials = m-5/m-4 spawn machinery + Step-3, outside the s4 fence; Step-1 assignment integrity rests on operator launch discipline (the m-1 wiring-time residual, stated).
