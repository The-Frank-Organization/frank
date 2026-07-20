# frank Model Runtime (m-9) — DESIGN

**DESIGN_DOC_ID:** step3-design-m-9-model-runtime
**Owner:** m-9 (Model Runtime) — design-lead m-9.planner · adversarial design-reviewer m-9.implementer
**Dispatch:** `step3-design-m-9` (`DESIGN-orchestrator-planner-20260715-005510`) · **Tier:** large · **Evidence:** E1 · **GRILL_REQUIRED:** yes (agenda §11; grill NOT yet run)
**Status:** **DRAFT r0 — NOT design-of-record, NOT reviewed, NOT locked.** The topology-invariant core (§1–§8) is drafted for the operator walkthrough; the topology annex (§9) is **OPEN pending the frank-wide architecture the operator is finalizing with master** (operator hold, in-session 2026-07-15: "best not to make any solid decisions yet"). Provisional operator direction is recorded AS provisional in §9. No section below is settled until: operator walkthrough → GRILL → m-9.implementer DESIGN-REVIEW → Master+VP reconcile.
**Basis (consumed, re-derived nothing):** the pair-confirmed c7 AUDIT matrix rev1 (`master/domains/m-9-model-runtime/audit/2026-07-14-model-runtime-promotion-matrix.md`, 40 rows — treated per dispatch as HYPOTHESES this design must prove, not facts); `STEP-3-KICKOFF.md` §1/§2/§5 (SHA `983508fc…`); the locked m-7/m-5/m-3 contracts; the m-8 charter; the five-layer terminal agenda (dispatch `:33-38`); the `step3-amend-m3-egress` cue (`:26-32`).

---

## 0. Frame — the one job, and the claim boundary

The Model Runtime drives **one real governed model turn**: assemble a request from governed state, hand it to an m-8 adapter, consume the normalized event stream, surface the model's output through m-3 observation, turn parsed tool calls into **inert governed execution requests**, fold authorized results back, and land the turn on exactly one typed terminal — with every canonical state transition committed through the m-7 governance door.

**Claim boundary (held in every section):** m-9 owns *logical* runtime state and *requests*. It owns no wire (m-8), no credential/endpoint (m-7 §1b amendment), no egress disposition (m-3 §1a amendment), no ceiling semantics or authoritative check (m-5 semantics, m-7 hosting), no canonical storage/commit/recovery (m-7), no routing (m-4), no spawn (Step-3 carry). Confusion-resistance language throughout; D5 residuals inherit the m-7/m-3 statements — this doc adds no new "by construction" claim.

## 1. The turn state machine (topology-invariant core)

One session holds **at most one active logical turn** (audit A2, invariant only — enforcement placement rides §9/Q1). A turn:

```text
IDLE
  → ASSEMBLING        build the request from committed session state (§2): deterministic,
                      digest-recorded; tool set = the trusted ceiling-filtered exposure CONSUMED
                      from m-5/m-7 (never computed by m-9)
  → SEND-REQUESTED    the assembled request enters the m-3/m-7 provider-request egress path
                      (§7); a DENIAL here ⇒ terminal turn_denied with ZERO wire send
  → STREAMING         consuming m-8 normalized events (bounded buffers, §6.4); stream is
                      PRIVATE to the runtime — no byte is recipient-visible or actuator-
                      actionable in this state (C9 floor, §4)
  → OBSERVING         the m-3-owned observation decision over the turn's output at m-3's
                      granularity (§4)
  → TOOL-ROUND        zero or more tool-request lifecycles (§3); each authorized result folds
                      back; then → ASSEMBLING (next model call of the same logical turn)
  → TERMINAL          exactly one of §1.1
```

Cancellation, timeout, provider failure, refusal, and overflow enter from any non-terminal state and land on §1.1 tokens. **Every transition that changes canonical state is a committed record through the governance door; in-memory state is cache, rebuildable from committed records** (audit A5/B1).

### 1.1 Turn terminals (five-layer agenda, layer 4 — m-9-owned vocabulary; PROPOSED, grill-locks at §11)

Byte-exact closed set, exactly one per turn:

```text
turn_completed    the model finished; output observed; all folded
turn_refused      typed content-filter/refusal finish (honest, never "empty response")
turn_denied       a governance gate said no BEFORE/AT the send boundary (egress denial,
                  routing_unavailable consumed from m-4, above-ceiling turn-blocking denial)
                  — zero or partial-zero wire activity, stated per-record
turn_cancelled    operator/steer cancel; partials committed labeled-partial (D3)
turn_failed       wire/adapter/machinery failure after bounded recorded retries (§6)
turn_exhausted    bounded-resource stop: max steps, doom-loop tripwire, token budget
```

**Separation law (preserves CQ-4):** turn terminals are m-9-layer semantics carried IN records; the records themselves still reach exactly one existing relay delivery-state `{accepted, rejected, held}` at commit. **No fourth delivery_state; no turn terminal is a gate token.** Compaction is NOT a terminal — it is a recorded non-terminal detour (§5).

## 2. Session & context state (Q2, within-substrate — audit B1/blocker-2 honored)

- **Canonical = committed records in the m-7 substrate.** Session = the append-only sequence of turn-boundary records (turn opened, send authorized+attempted, observation stamped, tool requests + dispositions, compaction entries, terminal). No m-9-owned second durable truth, no m-9 recovery path: after any crash the runtime rebuilds by reading committed records (m-7 recovers the store; m-9 merely re-derives its cache — the opencode persist-and-rederive shape, audit A5, on frank's substrate).
- **The full token-level transcript is DERIVED, not canonical** (PROPOSED, grill Q2): canonical records carry the turn's semantic content (messages in/out, tool calls/results, usage, digests); the raw stream/event log is a service-local derived artifact whose **digest is bound into the turn record** — auditable (bytes provable against the digest) without pushing token-level volume through the commit loop (designed at ~relay volume, m-7 §2.3). Retention of derived logs = operator config (m-7 GC posture pattern).
- **Context assembly is a pure function** of (committed history + compaction entries) × (pinned lane facts from the m-8 catalog: context_window, max_output — audit B4) × (the consumed ceiling-filtered tool exposure) × (versioned assembly rules). The **assembly digest + input refs are recorded per send** (kickoff §4 pinning) — replay-complete. Opaque provider blobs (reasoning-replay, audit B5) round-trip uninspected via m-8's contract; the ≥8-step close/reopen canary is the conformance fixture.

## 3. The tool-request lifecycle (audit C2 rev1 sequence + C7/C8/C9 — the charter's sharpest line)

```text
model output → [m-3 OBSERVATION at its granularity (§4)]
  → COMPLETE?   only a complete, finalized tool call proceeds (C7: partial/truncated deltas
                NEVER form a request; pi truncated-refusal = the fixture)
  → VALIDATED   deterministic schema validation (m-9-owned code); malformed ⇒ typed error
                result folded back as governed context (C1) — never reaches authorization
  → INERT REQUEST minted: stable request_id = (session, turn, call-ordinal, content-hash);
                duplicate/colliding IDs FAIL CLOSED (C8)
  → AUTHORIZATION requested at the trusted m-5/m-7 door (m-5 ceiling semantics, fail-closed,
                absent-default = floor; m-9 never evaluates)
  → EXECUTION   in the m-7-hosted unprivileged executor (no store/config/outbox handle);
                adapters carry NO execution capability (m-8 seam)
  → FOLDING     typed result (isError, interrupted, transforms declared per m-3 vocabulary)
                committed + folded into context
```

**Request states (durable, typed):** `constructed → observed → authorized | denied → executing → completed | interrupted | unknown_effect`. **No-replay law (C8):** recovery/retry never re-executes `completed` or `unknown_effect`; an `unknown_effect` surfaces for disposition (operator/policy), never silent retry. V1 executes tool requests **serially** (audit C6).

## 4. Streamed-output observation (Q3 — PROPOSAL to m-3, who owns the answer)

V1 satisfies the kickoff's `streamed-output observation` with **two observation points and a structural guarantee**:
1. **Structural floor (C9, by design of §1):** the stream is consumed privately in STREAMING; nothing is delivered, rendered to a recipient, or actuator-actionable mid-stream. Zero partial bytes escape pre-observation — the floor holds by state-machine shape, not by scanning speed.
2. **Observation point A — tool-request construction:** each complete tool call passes the m-3 decision BEFORE becoming an authorizable request (§3).
3. **Observation point B — turn boundary:** the turn's assembled output is the observe candidate at commit (the landed `observe.Env.Evaluate` shape — candidate-at-submit, no new streaming observer).
**Stream-time (mid-stream) observation is an explicitly-reserved later rung** — additive (a third observation point), no re-cut. m-3 may tighten granularity in its review; this design holds under any granularity ≥ the floor.

## 5. Compaction (audit B3, as governed mechanics)

Compaction is a **committed, observable event**: trigger condition + before/after context digests + summary provenance recorded; **the summarizer call is itself a governed turn** (rides §1 + §7 like any send); the auto-continue injection (if adopted) is **system-stamped**, never author-forged (m-1). Trigger **policy** = m-5/operator (RUNTIME-RESEARCH §12 #6); m-9 owns deterministic mechanics + typed state. Donor trigger math/cut-points/tail-preservation enter as **conformance fixtures to prove, not inherited policy**. Overflow is never retried as a wire error (audit E3) — it routes here.

## 6. Errors, retry, cancellation, timeout (turn-level dispositions only)

- **6.1 Errors-as-events:** the m-8 stream never throws; failures arrive as typed normalized events (audit E1); m-9 maps them to §1.1 terminals. Refusal is `turn_refused` (E4).
- **6.2 Retry split (feeds §7):** wire-level retry/idempotency mechanics = **m-8-owned, below the turn boundary**; every attempt is **recorded** (opencode's published retry-status shape). m-9-level re-issue of a model call is **bounded and recorded**, never silent, and **never re-executes tool effects** (C8). Exhaustion ⇒ `turn_failed` with the attempt trail.
- **6.3 Cancellation:** m-8 owns wire cancel semantics; m-9 disposition: partials committed labeled-partial, interrupted ≠ failed, in-flight tool requests → `interrupted`/`unknown_effect` per C8, terminal `turn_cancelled`. Steer-at-boundary hook (audit D2) exists at the ASSEMBLING re-entry point; steering itself = Step-3 carry.
- **6.4 Timeout & backpressure:** provider-await rides the landed expiry pattern (m-7 owns timer hosting; disposition typed, operator-extendable) — m-9 adds **no scheduler** (Q6 seam: m-7/m-6 review). Stream consumption uses bounded buffers; overflow of the bound is a typed fault, never silent drop, and never blocks m-7's commit loop (E6).

## 7. The provider-send boundary — m-9 as CONSUMER of the m-3/m-7 amendments (+ the Q4 packet, §7.1)

The runtime's send path: `ASSEMBLING → [pre-translation check] → m-8 translation/binding → [FINAL AUTHORIZATION at the governed send boundary — m-3-designed, m-7-hosted] → wire`, with **zero mutation after final authorization** (kickoff §1a). m-9 consumes typed outcomes: `send_authorized | send_denied (⇒ turn_denied, zero send, no wire event) | send_failed`. Credentials/endpoints never appear in m-9's surface (§1b).

### 7.1 The Q4 consumer packet (OWED to `step3-amend-m3-egress` — m-9's requirements as the consumer; m-3 owns the disposition design)

1. **Authorization identity:** the final authorization should bind to a **content digest of what will actually leave** (final wire bytes, or pre-translation request + pinned deterministic translation version — m-3's grill choice) + an **idempotency key minted no later than authorization**, so provider-side dedup and frank's authorization identity align.
2. **Byte-identical wire retries** (same digest, same idempotency key) within a **bounded attempt count + window** should ride the SAME authorization, with each attempt **recorded** — otherwise every transient 503 forces a full re-authorization round-trip and the retry path becomes the untested path.
3. **Any byte change is a NEW authorization** — endpoint, header, body, translation version. No silent re-auth, no mutate-then-reuse.
4. **Tool effects never ride wire retries:** a retried provider call must be side-effect-free w.r.t. tool execution (C8 no-replay is m-9's own law; stated here so the amendment's retry design never assumes re-execution).
5. **Typed denial, consumable:** `send_denied` must be a typed event distinguishable from wire failure, with the mapping to existing relay delivery-states m-3 designs (no fourth token, no away-email park inheritance) — m-9 maps it to `turn_denied`.
6. **No-send = no wire event** (the locked floor): m-9's `turn_denied` record must be committable with provably-zero wire activity.

## 8. The m-8 contract requirements (consumer-review inputs; m-8 owns its contract)

Carried from audit §4, now design-binding for m-9's consumer review: normalized events with **no execution capability** · typed finish/error incl. refusal · cancellation with partial-stream disposition · **completion semantics distinguishing partial/truncated from complete tool calls** (C7) · recorded wire-retry attempts + idempotency below the turn (§7.1) · **backpressure/bounded-stream mechanics** (E6) · reasoning-replay round-trip opacity (B5) · facts from the pinned catalog, never a merged options-bag (F4).

## 9. TOPOLOGY ANNEX — **OPEN; consumes the master architecture-of-record when it lands**

**Operator hold (2026-07-15, in-session): the frank-wide service architecture is being finalized with master; no solid decision here.** Recorded provisional direction (operator, same session, explicitly provisional):
- the runtime is a **separate program**, not an in-conductor subsystem ("the conductor is merely one (differentiating) slice of frank … separate microservices build the app as a whole");
- the m-8 adapter/wire concern likewise trends toward **its own service** (operator picked this shape when asked directly);
- the conductor trends toward **governance-service + bootstrap trust root, not inter-service bus** (data plane direct service↔service; **governed effects keep the single conductor door**) — operator: "sounds like it," pending master.

**What the annex must bind once the architecture lands (and ONLY then):** the runtime's service identity/authentication mechanism · the transport for governed commits/authorization requests · the stream channel m-8→m-9 · crash/restart contract between services · where the m-3/m-7 final-authorization check physically executes. **Invariance claim (to be proven at review):** §1–§8 hold unchanged under any of the candidate topologies; only this annex's bindings vary.

## 10. Seams + fixture skeleton (bite to be authored with the implementer at review)

| seam | contract consumed | negative fixture (sketch) |
|---|---|---|
| m-9↔m-8 | §8 list | adapter offered an execute callback ⇒ no such surface exists; truncated call ⇒ no request minted (C7) |
| m-9↔m-3 | §4 observation + §7 egress disposition | unobserved tool-call bytes ⇒ cannot reach authorization; `send_denied` ⇒ zero wire event |
| m-9↔m-5/m-7 | consumed tool exposure + authorization door | above-ceiling request ⇒ zero execution; m-9 code path evaluating a ceiling ⇒ does not exist |
| m-9↔m-7 | substrate commit + expiry + recovery | kill -9 mid-turn ⇒ rebuild from committed records only; completed/unknown-effect request never re-executes (C8) |
| m-9↔m-6 | turn surface via existing buckets/ODB | no new human-surface subsystem; asks ride approval/held/park (Q6 seam review) |
| m-9↔m-4 | executes accepted routing records only | absent/invalid lane ⇒ typed `routing_unavailable`/`human_decision_required` consumed, never fallback |

## 11. GRILL agenda (owed → durable GRILL_LOCK_ID before DESIGN-REVIEW close; runs AFTER the architecture lands)

1. **Q1 placement** — adopt the master architecture-of-record; grill the m-9-specific residue (crash contract, stream channel, identity binding — §9 list).
2. **Turn-terminal vocabulary** (§1.1) — byte-exact set + the separation law.
3. **Q2 transcript locus** — canonical-semantic vs derived-with-digest split (§2), volume math shown.
4. **Q3 observation granularity** — with m-3 (the §4 proposal vs stream-time now).
5. **Q4 packet** (§7.1) — with m-3's amendment grill (final-wire vs pre/post pair interaction).
6. **Q6 provider-await + human-surface seam** — with m-7/m-6 (expiry vs park posture).

---
**ACTIONS_GIT_REF:** draft design doc only; no `frank/` edit (docs workspace, not a git repo).
**FINAL_GIT_STATUS_SHORT:** unavailable — cwd is not a git repo (docs workspace); `frank/` untouched at `502e06c`.
