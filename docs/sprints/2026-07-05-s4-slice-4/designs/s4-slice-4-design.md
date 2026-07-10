# s4 Slice-4 — DESIGN: the wire-up (the per-seat MCP shim · seat-lifecycle hardening · the §7 config-change record)

**DESIGN_DOC_ID:** `s4-slice-4-design`
**Owner:** s4-wire pair — authored by `s4-wire.planner`; design-challenger + formal DESIGN-REVIEW addressee = `s4-wire.implementer` (per the protocol's lineage rules)
**Dispatch:** `.relays/s4/s4-wire-design/DESIGN-orchestrator-planner-20260705-013107.md` (twelve binding constraints; GRILL_REQUIRED: yes)
**Rev:** r3 (operator grill COMPLETE — GRILL_LOCK at §12; **guide answers FOLDED** — `.relays/s4/s4-guide-q1/SITREP-planner-20260705-014633.md`, all six confirmed from locked text, zero amendments; every [PG-Qn] marker resolved in place; the F2 structured-carrier rule closed at §3.2)
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
- form field → property under `headers` (the fill surface m-2 declares); `options` → `enum` (byte-exact values — the enum-byte-exactness floor rides through unmodified); `default` → `default`; parent-picker options → enum of conductor-supplied candidate relay IDs (+ default).
- **The closed carrier rule (guide-guardrail-conformant: the schema is the rendered form VERBATIM — the shim performs ZERO type reshaping):** every FieldSpec type maps to a **string-typed** JSON-Schema property. The value the host writes is byte-for-byte the value the conductor receives in `Headers map[string]string` (record.go:27-31) — the shim never parses, canonicalizes, re-encodes, or restructures a field value in either direction. Per type (the full live set — registry census: text ×12, enum ×12, id_ref ×7, bool ×4, row_array ×3, address_list ×2; plus `object` in the ParseTyped contract):

| FieldSpec type | JSON-Schema property | notes |
|---|---|---|
| `text` | `{type: string}` | — |
| `enum` / `bool` | `{type: string, enum: <rendered Options verbatim>}` | options byte-exact from the render; bool's values are whatever the registry declares — the shim invents no true/false spelling |
| `id_ref` (incl. parent-picker) | `{type: string}`; picker fields: `enum` = conductor-supplied candidate relay IDs, `default` from render | candidate sets are conductor-derived (m-1 §6); DigestExempt handling is server-side (formForDigest) — invisible to the shim |
| `row_array` / `object` / `address_list` | `{type: string}`, description states: "canonical JSON string — <shape>" (e.g. `[{"col":"val"},…]` for rows) | **the S3 locked carrier rides through unchanged**: structured values ARE canonical-JSON strings inside string headers (s3 design :52-58; canonical.go:19-75); the conductor's `ParseTyped` is the ONLY parser/validator — a non-canonical encoding gets the conductor's typed violation naming the field, surfaced like any bounce (§3.4), never a shim-side error |

  Rejected (recorded §13): exposing structured fields as native JSON objects/arrays with shim-side canonicalization into string headers — it would put a second implementation of the m-2 canonical-encoding rule inside the trusted bridge (drift risk), split violation-surfacing between shim and loop, and break "the digest proves the exact host-visible schema" (the host-visible shape would differ structurally from the rendered form the digest covers). The ergonomic cost of JSON-in-a-string is real but small, borne by models that handle it routinely, and every mis-encoding gets a typed, field-named bounce.
- envelope fill surface = exactly `{to, cc?, dispatch_id}` + `body` — recipient/dispatch selection is sender-authored per m-1 §6; **no `from`, no `role`, no `delivery_state`, no `relay_id` property exists in the schema** (system_only fields are absent, not "read-only").
- **`form_digest` = a required `const` property** carrying the digest of exactly this rendered schema (**GRILLED G-6: the model copies it back**); that copy is the proof-of-current-render the digest exists to be (m-2 drift semantics). The shim **never synthesizes or substitutes a digest** — a shim-injected digest would claim the host saw a form it may not have seen (claim-dishonest; rejected alternative §13).
- The tool description carries the honesty line: "Files a governance relay (transport/provenance only — content claims are not verified by this tool)." Never "verifies", never "authorizes".

### 3.3 Drift + phase negotiation (the re-render loop, built server-side; guide Q6(i) CONFIRMED — `tools/list_changed` + re-list is the right MCP-native mapping; the advisory/authoritative split untouched)
`Validate` re-renders for the candidate's **declared** `PHASE`/`CEREMONY_TIER` and compares the carried digest against that render (validate.go:20-25). The loop this licenses, with zero new server surface:
1. Host fills the current schema and calls `submit`.
2. If the host's declared PHASE/TIER matches the schema served, the digest matches → normal pipeline.
3. Any digest mismatch (phase switch, tier switch, or a real config/candidate drift) → typed `re-render` bounce. The shim, on seeing violation class `re-render`: re-fetches `tools/descriptions` **for the phase/tier the host declared**, rebuilds the schema, emits `notifications/tools/list_changed`, and returns the bounce as a structured non-fatal tool result: `{state: "rejected", violations: [{field: "form_digest", class: "re-render", hint: "form refreshed — re-read the submit tool schema and re-submit"}]}`. The host model re-reads and re-submits — the [VP-W4] "superseded form bounces re-render; re-rendered form succeeds" leg is this loop verbatim.
4. The shim MAY pre-fetch the declared phase's schema on a phase-mismatch **before** submitting (saves one bounce round-trip) but still never carries a digest for a schema it hasn't served to the host — the submit goes out with the digest the host copied; the pre-fetch only accelerates step 3.

### 3.4 Tool-call results
MCP `tools/call` result = `{content: [{type: "text", text: <outcome JSON>}], isError: <transport-error only>}`. Outcome JSON is the conductor's own (state/relay_id/intake_id/violations) — byte-faithful, path-free already (audit P10). `isError` is reserved for transport/auth failures; a `rejected`/`held` outcome is a NORMAL result (the model must read it, not see a tool crash).

### 3.5 The shim's own surfaces are I-PH surfaces [VP-W3; guide Q6 guardrails folded]
**Two guide guardrails on the shim as trusted surface, binding this section and §3.2:** (1) the `inputSchema` it serves is the rendered form **verbatim** — no shim-side reshaping that could re-add an affordance the render withheld; the absence set + I-PH apply to **every byte the shim re-frames**; (2) the shim holds the seat credential and IS the channel (m-1 binding grain) — its custody note rides the m-1 fidelity packet, and the D5 residual is stated wherever the shim's isolation is described. One named exception to the no-config-values rule: the §6.4 ceiling-value disclosure carve-out (exactly that one value, fixtured as such).
The shim knows the socket path and the credential. **Neither may reach any MCP-visible byte.** Every self-generated error is mapped to a typed, fixed-string class before surfacing: `shim:conductor-unreachable`, `shim:auth-failed`, `shim:connection-lost`, `shim:frame-too-large`, `shim:protocol-error`. Go's default dial error (`dial unix /path/…: connect: …`) is scrubbed at the single error-translation chokepoint (one function; every MCP-bound error passes through it; fixtured). Raw detail goes ONLY to the shim's stderr log (host-operator-facing, never model-facing) — and even stderr never prints the credential. The exit-gate I-PH matrix covers all seven bridge classes: tools/list descriptions · input schemas · tool-call results · notifications/poll hints · reconnect errors · credential-failure errors · shim diagnostics.

## 4. D-3 — second-connect: the per-credential active-channel index [VP-W1; source of record = s4-dispatch]

Engine-internal, no new seat-facing surface, **no binding-table shape change** (a needed shape change = hard stop + escalate):
- `serverConn` retains, post-`Resolve`, the SHA-256 of the presented credential (never the raw value — it lives only in the binding table and the wire read).
- `Server` gains `active map[credHash]*serverConn`, maintained under the existing `s.mu`: set on successful auth; deleted in the conn's `run()` defer (kernel close-detection — the proven-dead signal).
- **Reject leg:** `session/connect` presenting a credential whose hash is live in `active` → typed, path-free `auth:channel-active`; the existing channel is untouched. Any number of *distinct-credential* channels remain fine.
- **Recovery leg:** prior conn closed (host death, shim kill, socket drop) ⇒ its defer removed the entry ⇒ the next `session/connect` binds normally. This is reconnect — no new mechanism, no grace timer.
- **Wedged-but-alive host (GRILLED G-3, operator-accepted; guide Q5 deferral CONFIRMED with two binding conditions):** NOT provably dead (socket still open) ⇒ reject regime holds; no heartbeat/ping this slice, no grace timer, no supersede (a supersede would also be a session-hijack affordance — whoever holds the credential could yank a live seat); liveness *policy* lands with the scheduler layer (m-6, Step-2+). The guide's conditions, adopted: (1) the operator remedy — "kill the host session/shim; the kernel close frees the channel" — is a **written ops-surface deliverable** (§7; stated-not-skipped, never a tribal fact); (2) one fixture pins the escape hatch: reject-active regime → host killed → kernel close-detection → reconnect → recovery leg completes (S4-SC3). The deferral is thereby a documented boundary, not a gap.
- **Explicitly out (escalate, amendment path):** live supersede ("new connect wins"), credential rotation, re-mint-supersedes. Nothing in this design gives a second connect any effect on the first beyond the typed reject it receives.

## 5. D-4 — per-recipient wake; retiring broadcast + the global pending queue **(guide Q1 CONFIRMED — per-seat IS the locked shape; the cross-seat leg is a defect to fix, not preserve)**

The locked shape is m-7 §8.3: delivery = one write onto **the recipient's** pipe. Today's code broadcasts one global frame to every client and replays a never-cleared pending queue to every future auth (audit probes (a)/(b), live-confirmed). **The guide's three load-bearing properties — none of them the broadcast itself — bind this section:** (1) **nudge-on-auth-if-pending, derived from the STORE** at that moment, never from a mutable side-list (the retired global `pending` is exactly the side-list recovery must not depend on); (2) **fire-and-forget with mailbox truth** — a dropped/failed push is never a lost delivery; keep the write-error-continue posture; (3) **no cross-seat metadata in any seat's frames** — the audit's non-recipient-sees-pending-state finding is an absence-set-adjacent leak (§8.4 grain), fixed as part of the locked shape, with a negative fixture (seat B's frames never contain seat A's identifiers or pending state — S4-NG4). Design:

- **Wake events ride the commit's own delivery intents.** The commit already materializes per-recipient mailbox entries; the loop's post-commit hook emits `(seat, relay_id)` wake events for exactly the committed record's recipients. `Server` gains a seat→active-conns view (the §4 index already knows seat identity via the auth meta); `PushTo(seat, frame)` writes `{"kind":"delivery-nudge","relay_id":"…"}` to that seat's connection(s) only. A seat's own relay_id on its own channel — no cross-seat metadata (the audit's side-channel note dies here).
- **Offline recipient:** no connection ⇒ no queue in the channel layer at all. The mailbox IS the queue (store truth). On that seat's next auth, the server checks pending-delivery state **for that seat only** and sends one `{"kind":"recovery-nudge"}` on that connection. `PendingDeliverySeats()` (store.go:175) gains a per-seat form (or is filtered at the call site — implementer's choice at PLAN grain).
- **The global `pending` slice and `Push`/`broadcast` are RETIRED from the serving path** (audit F-W4: retire, don't patch). `QueuePush`/`broadcast` survive only if fixtures need them; the production assembly stops calling them.
- **Startup recovery-nudge:** replaced by the same per-seat on-auth check — the process-start special case disappears (a seat that connects after a conductor restart gets its own recovery nudge exactly like any reconnect).
- **Nudges are best-effort by design** (client buffer can drop, hosts may hide custom notifications — F-W2): the shim documents and implements **poll-first**: `project` at turn start is the reliable path; nudges only accelerate it. The shim surfaces a nudge as a host-visible MCP notification where the host supports it, else drops it silently and relies on the poll hint in the tool descriptions. **Guide Q6(ii) CONFIRMED: poll-hint-first is first-class BY LOCKED DESIGN, not a concession** — inbox = durable truth, pipe = nudge (§8.3); hosts that never surface notifications still converge via `project()`-on-turn.

## 6. D-5/D-6 — the §7 config-change record **(GRILLED G-1/G-2; guide Q2 CONFIRMED (i)–(iv) — the compound-record idiom; guide Q3 CONFIRMED (a))**

### 6.1 Shape (GRILLED G-1/G-2 — operator-resolved; guide Q2/Q3 confirms mechanics only)
- `record_kind: config_change` (token + provenance ride the m-1 fidelity packet — ruling condition 4; no other new record_kind).
- **Authorship = operator-channel `submit` (GRILLED G-2; guide Q3 CONFIRMED with the reason recorded):** the operator seat files a `config_change` relay through the normal pipeline — form-validated, operator-scoped (the registry restricts the record_kind value to the operator seat's rendered options; a non-operator submit carrying it bounces — fixtured), committed via the one pivot. **The accepted relay IS the operator-authorized digest-change record** — the authorization record is the mechanism working, first-class in the store. The `system`-stamped admin-time alternative is rejected on the guide's ground: it would mint a **second provenance path for an authority-bearing act on a live store**, bypassing the governed interface the record exists to prove (genesis is the sanctioned system-stamped exception only because no channel exists at store birth). **Guide semantics note, binding:** the operator's authorship IS the human authorization — the record does NOT additionally park for operator approval (a human-gate loop back to the same human is ceremony, not safety).
- **Payload embeds the new member bytes (GRILLED G-1):** headers carry `{member: fieldspec|engine, new_digest: <top-level digest after the change>}`; body carries the full new member content (registry.json ≈ 10 KiB today — comfortably inside the grilled frame bound, §6.4). Store-is-truth: config history is auditable append-only from records alone; recovery/phase-0 re-materializes `config/` from the chain after any crash or hand-edit (self-healing — the operator-facing property that decided the grill row); `config/` becomes a derived projection like INDEX. The digest-only alternative (files changed out-of-band, record merely attests) was grill-rejected: two truths, store-down failure mode, history not reconstructable from the store.

### 6.2 Mechanics under no-hot-reload (restart-boundary semantics — guide Q2 sequencing RULED: record-pivot-then-derived-bytes)
- Commit-time (running under the OLD config): validate → **one canonical pivot** (the config_change record; F11 one-pivot holds — exactly one rename) → **derived step:** materialize the embedded member bytes to `<root>/config/…` via the existing atomic write, in the same loop iteration, after the pivot. **Guide grain adopted verbatim:** the byte replacement is **derived work — a projection-grade artifact of the committed record**, replayed idempotently by recovery from the record's embedded payload (the S2 derived-work-completion mechanism generalizes; this is another derived-intent class). The member FILES are thereby demoted to derived artifacts; the append-only truth of config history is the record chain (config/ = derived projection, exactly like INDEX/mailboxes — canonical wins, unconditionally).
- The RUNNING process keeps serving under the old pinned config (no reload — locked). The new registry takes effect at the next restart. Consequence stated plainly on every claim surface: **a config change is effective-at-restart**; the [VP-W4] drift leg (old rendered form bounces `re-render`; re-rendered form succeeds) is exercised across that restart.
- **Phase-0 learns the chain (guide Q2(iii) sharpenings adopted):** expected digest = genesis's `config_digest` superseded **in commit order** by each accepted `config_change`'s `new_digest` (the "genesis chain" of m-7 §7 :109); the chain is derived from **committed records ONLY — no side ledger**. Loaded `config/` digest ≠ expected ⇒ phase-0 **re-materializes** the members from the chain's embedded bytes (idempotent redo — ARIES property), recomputes, and a persisting mismatch **anywhere in the walk** takes the same **fail-closed-serving-reads, summon-operator** disposition as today — never a brick. `ValidateGenesis`'s static compare (genesis.go:104-118) becomes the chain walk; a fresh store (no config_change records) degenerates to today's exact behavior — zero regression on S1–S3 stores.
- **Crash legs (the applicability row `config-change`):** kill between pivot and file materialization ⇒ recovery re-materializes (record = truth); kill mid-member-write ⇒ atomic file write discipline + redo; kill before pivot ⇒ nothing happened (intake re-enqueue). `f11Classes()` gains the class; every syscall boundary swept by the existing child-SIGKILL harness; **exactly one canonical rename per config change** — the derived file writes are not canonical renames.
- **Never re-genesis:** the round-trip runs on the s4 team's live store AFTER it has real records (existing store by construction); `store.Init` is untouched.

### 6.3 OI-S3-CONFIG-CHANGE discharge
After the live round-trip: the operator authors the `owed_disposition` for `OI-S3-CONFIG-CHANGE` through the operator channel on the real store (the S2 worked pattern), citing the config_change relay + the gate evidence; open owed set = empty at the exit gate. The frozen s2 store's upgrade is NOT a gate leg (optional operator call, out of gate scope).

### 6.4 Frame-transport bound rides with this **(GRILLED G-5; guide Q4 SANCTIONED — typed refusal BOTH directions; carve-out recorded)**
Raise both scanner buffers to **1 MiB**, sourced from the engine config (additive member `max_frame_bytes`; absent = 1 MiB default, so existing stores' digests are untouched). Guide rulings folded verbatim:
1. **Inbound oversize** ⇒ a typed, path-free transport-grade refusal frame (`frame-too-large` class), **connection kept alive**, nothing staged, nothing in intake — the fault-taxonomy grain applied at the transport tier (never silence; audit F-W1).
2. **Outbound oversize** — a `project`/`read` response outgrowing the ceiling as the store grows (the sneakier regression) ⇒ ALSO a typed refusal **with a narrower-re-scope hint** (re-`read` by id; paginate the projection), never a silent connection kill. Fixtured in both directions (S4-FR1/FR2).
3. **Disclosure carve-out, named (the decision-⑤ pattern):** the ceiling VALUE may appear in the refusal reason — it is actionable transport metadata, not §7 *policy* config — recorded here as an explicit, named exemption to the §8.4 "no config values on seat surfaces" absence set (an explicit carve-out beats a silent leak or a useless bounce). The bound stays in the engine member, out of the policy-config artifact; the I-PH fixture matrix asserts the carve-out covers exactly this one value and nothing else.
4. **No chunking this slice** — a dialect change; the Q6 composite gives the shim its own framing layer later without touching the private dialect.
**Grill basis (measured, 2026-07-05):** 711 relays across every trail (517 master + 194 frank s1–s4) — median ≈ 5 KB, p99 ≈ 20–25 KB, all-time max 33 KB; no relay has ever exceeded 64 KB, but the tail grows each slice and ecosystem docs already reach 93 KB — 1 MiB = ~30× max-ever headroom and swallows doc-sized bodies whole.

## 7. D-7 — custody posture + ops surface + usage posture (docs-grade, claim-honest)

- **Custody, stated honestly everywhere credentials appear** (m-1 §13.3 carry; D5): mint prints to stdout (capture discipline documented); `-operator-submit -credential` is ps-visible (admin tooling, short-lived — documented, not fixed this slice); binding table plaintext-at-rest 0600/0700 (D5-accepted, stated); shim credential env-var-first (§2). **No in-band rotation/revocation exists: a compromised credential means stopping the conductor and admin-time surgery.** Confusion-resistant, not theft-proof — verbatim on every custody surface.
- **Ops surface doc** (`docs/ops.md` or per PLAN): start (`-init` → mint → serve), stop (signal; durable FIFO + recovery make it safe), status (socket liveness + read-only diagnostics states genesis-missing/digest-mismatch), **short-socket-path rule** (darwin AF_UNix ≈104-byte cap; the default `<root>/frank.sock` join violates it for deep roots — convention: `/tmp/<team>.sock`-class paths; pre-flight length check as PLAN hardening), team-store conventions (absolute store root; the store is the governance domain), minting workflow end-to-end (who runs mint, where the credential goes, per-host wiring for Claude Code and Codex MCP configs). **Per-seat wiring (grill G-4 rider):** a host MCP `env` block is per SERVER ENTRY, not per session — so the doc pins **one seat = one credential = one config entry**, with two blessed patterns: (1) one config entry per seat in separate scopes/working dirs (matches how the team already runs seats as separate sessions); (2) `${VAR}` indirection where the host supports env expansion — the config carries `"FRANK_CREDENTIAL": "${FRANK_CREDENTIAL}"` and the secret arrives per-launch from the shell. Host support pinned per host in the doc; a same-credential double-launch gets the §4 typed reject (the fence working, documented as such).
- **Seat-occupancy model (grill G-4 follow-up, operator-raised — the paragraph the ops doc inherits):** seats are **durable identities that sessions OCCUPY at launch** by presenting the seat's credential; kill the session, relaunch with the same credential → same seat, same mailbox, history via `project` (m-7 §8.5). s4 automates the **transport**, not the **assignment**: the roster is minted admin-time; the operator launches each session as its seat (one `${VAR}` launch line in the shared team cwd); one live occupant per seat (§4). Conductor-spawned sessions with injected credentials = the m-5/m-4 spawn machinery + Step-3 standalone runtime — outside the s4 fence (spawn automation is authority-adjacent). Step-1 assignment integrity therefore rests on operator launch discipline — the m-1 wiring-time residual, stated here and in the ops doc, not implied away.
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

E2 (conductor-registry/battery): S4-MCP1 shim dialect goldens (initialize/tools-list/call shapes, id-0 pin) · S4-SCH1 Form→JSON-Schema mapping incl. system-field ABSENCE + const-digest + enum byte-exactness · **S4-SCH2 structured-carrier legs** (a `row_array`/`address_list` field round-trips as a canonical-JSON string; a non-canonical encoding gets the conductor's typed field-named violation through the MCP result; the schema description states the string carrier honestly) · S4-RR1 drift loop (phase-switch bounce → refreshed schema → success) · S4-SC1/SC2 second-connect reject (typed `auth:channel-active`) / proven-dead recovery · **S4-SC3 wedged-host escape hatch** (reject-active → kill host → kernel close → reconnect → recovery leg completes; guide Q5 condition 2) · S4-NG1..3 recipient-only nudge, non-recipient silence, offline-reconnect recovery-nudge (per-seat, store-derived) · **S4-NG4 no-cross-seat-metadata negative** (seat B's frames never contain seat A's identifiers/pending state; guide Q1 property 3) · S4-FR1 inbound frame-bound typed refusal, connection alive · **S4-FR2 outbound oversize typed refusal + narrower-re-scope hint** (`project`/`read` growth; guide Q4 ruling 2) · S4-IPH1..7 the seven bridge-surface classes each path-clean (incl. shim-diagnostic scrub with a real dial failure; **the ceiling-value carve-out asserted to cover exactly that one value**) · S4-C7-1..n config-change: operator-only refusal, chain walk (committed-records-only), re-materialization after each crash leg, applicability row, one-pivot assertion, restart re-render leg · E2 floors: full battery, enum grep, three-tool enumeration.
E3 (gate-run, operator-designated seats): the live two-session relay (A=Claude Code, B=Codex ideally) · live adversarial legs (no/bad credential, second-connect, forged submit, I-PH probes over live surfaces) · crash/liveness legs (kill frank mid-delivery; kill shim; offline-seat nudge) · the §7 round-trip + OI discharge on the real store. Every E3 record carries the transport-only qualifier.

## 10. What this design does NOT do (OUT, restated)

Consumer schema content (s5) · observe/evidence (Step-2) · routing (Step-3) · TUI/email UX (Step-4) · federation (zero pre-work) · external send (outbox dormant) · steer/interrupt beyond host-native · authority replacement (gates still park for the human — transport only) · in-band rotation/supersede (§4) · socket-dialect rewrite toward MCP (**guide Q6 CONFIRMED: §8.1's locked property is the seat-visible tool surface + per-seat channel identity, not a wire-format mandate — the composite realization stands**).

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
- native-JSON exposure of structured fields with shim-side canonicalization (r3/F2) — a second
  canonical-encoder inside the trusted bridge; splits violation-surfacing; breaks digest-proves-
  host-visible-schema; the string-carrier exposure is the verbatim-honest shape

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
- **r3** (2026-07-05): the must-revise fold (`DESIGN-REVIEW-implementer-20260705-035327`, F1+F2, both verified by me against the guide relay + source before folding). F1: the `s4-guide-q1` answers (…-014633, all six confirmed from locked text, zero amendments) folded into the BODY — §5 (Q1 three load-bearing properties + S4-NG4), §6.1 (Q3 second-provenance-path reasoning + no-extra-approval note), §6.2 (Q2 record-pivot-then-derived-bytes as a derived-intent class; chain sharpenings: committed-records-only, mismatch-anywhere ⇒ fail-closed-serving-reads), §6.4 (Q4 both-direction typed refusals + the named ceiling-value disclosure carve-out), §4 (Q5 deferral conditions: written remedy + S4-SC3), §3.3/§3.5/§5/§10 (Q6 composite + list_changed mapping + poll-first-by-design + the two shim guardrails); every [PG-Qn] marker resolved in place. F2: the closed Form→JSON-Schema carrier rule at §3.2 (all-string properties; structured types ride the S3 canonical-string carrier verbatim; ParseTyped stays the only validator; native-JSON+shim-canonicalization rejected, §13) + S4-SCH2. The queued seat-occupancy paragraph landed at §7. Grill rows untouched.
