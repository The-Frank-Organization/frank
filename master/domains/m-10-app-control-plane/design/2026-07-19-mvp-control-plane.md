# m-10 MVP Control-Plane DESIGN — the app-shell supervisor that RUNS the frozen seam contract

**Status:** PROPOSED r0 (stage-5; dispatch `step3-mvp-stage5-m10/DESIGN-orchestrator-planner-20260719-225207`; GRILL_REQUIRED — the §12 grill items are OPEN and this doc does not lock until GRILL_LOCK + the final-byte pair review).
**Rule of this document:** it REALIZES the stage-1 seam contract (`2026-07-16-mvp-ipc-manifest-seam-contract.md` r36 @ `0240e874ba553775a07b0b9c77be852e2cdfcbdb31fd4c489c62a87604218e01`) and re-opens NOTHING — every normative wire/store/order statement lives in r36 and is cited, not restated; any byte change to a closed artifact (including r36 itself) goes through the full F73 sequence. Where this doc and r36 could ever be read to differ, r36 governs.

**Consumed BYTE-BOUND (the stage-3-closed set):** m-10 r36 `0240e874…` (mine — the contract this design runs) · m-9 r19 `2a96a07b…` (the supervised worker counterparty: lifecycle halves, F59 executor, three-identity guard) · m-8 r12 `4b670a79…` (the supervised connector: bootstrap/ready gate, lane catalog, attempt results) · m-7 r11 `9331ea88…` (the broker I feed `epoch_state` to and dial into on CI-1; I hold NO conductor verb) · m-3 r4 `009df607…` (E0 event schema + digest/evaluator bindings) · m-1 `7c8b09a6…` (secret boundary: opaque `credential_ref` only) · m-2 `83d8e63e…` (form/tool-schema encodings the manifest vectors carry).

---

## §1 Topology — ONE module in the app main process

The m-10 control plane is a **module inside the app main process** (grill #1, ratified; the MVP amendment §2b), not a daemon. Its boundaries are **designed as-if process-separated**:

- **Real process seams (children):** CTRL-W to the worker (m-9), CTRL-C to the connector (m-8) — spawn-inherited socketpairs, §A.2 frames, exactly as contracted. DATA-P is created at spawn and **m-10 retains no endpoint** (r36 §A.1/§G.3).
- **Real network seam:** CI-1 to the broker (m-7 §2.10) — m-10 as dial-in client, control token + `control_generation`.
- **In-process seams (designed separable):** the operator terminal surface (§10) and the E0 event writer (§11) talk to the supervisor ONLY through the same command/query interface a separate process would use — no shared mutable state, no direct store handle. The Step-4 split (if ever wanted) is a transport swap, not a refactor.

Module decomposition (each a single-responsibility component; names are design vocabulary, not code):

| component | responsibility | store access |
|---|---|---|
| **applier** | THE single writer: every §F transaction, serialized (§2) | sole writer |
| **supervisor** | lifecycle state machines (worker §4, connector §5) driven by applier-committed state + channel events | via applier |
| **scheduler** | run/turn admission, lease, park/wake (§6) | via applier |
| **channel endpoints** | one reader/writer pair per live channel (CTRL-W, CTRL-C, CI-1); framing + grammar checks (§A.2) only — no policy | none |
| **terminal surface** | operator commands + state projections (§10) | read-only snapshots via applier |

## §2 Event architecture — the serialized apply loop (the §F chokepoint, running)

**One apply loop** owns the store. Channel readers, timers, and the terminal surface submit **typed events** to it; it executes each as **one §F transaction through the single-transition-chokepoint path** (r36 §F), bumping `state_seq`. Nothing else writes. This realizes single-writer discipline **by construction** (no lock protocol to get wrong) and makes every r36 "one transaction" claim structural: the retirement transaction, the F59 consume, the atomic lease-fault commit, the record commit — each is one event handled by one applier execution.

- **Ordering:** per-channel FIFO is preserved from socket to loop (one reader per channel, one queue). Cross-channel order is whatever the loop dequeues — every r36 total table is already safe under any interleaving (that is what the fences are for), so the loop imposes no cross-channel ordering promises.
- **Timers are events:** `HANDSHAKE_DEADLINE`, `ATTACH_DEADLINE`, `HEALTH_INTERVAL`, `SEND_DEADLINE`, wall-clock/turn — all fire into the same queue; a deadline disposition is applied by the same serialized path as a frame.
- **Replies/commands out** are emitted only AFTER the applier commit returns (durable-then-visible, r36 §B.1/§D) via the channel writer's bounded queue (§A.3 backpressure; full-past-deadline ⇒ channel-fault event back into the loop).
- **Reads:** the terminal surface and disclosure members read **committed snapshots** requested through the loop (or an equivalent read-transaction that can never observe an uncommitted write) — no dirty reads exist in the design.

**Provenance:** this is m-7's conductor serialized-commit-loop pattern (their §2A discipline), deliberately re-instantiated — **same pattern, separate store, separate code path** (no shared engine, no shared schema, no shared file; the conductor's store stays sole-governed by m-7). What is shared is the DISCIPLINE, not the artifact. → **GRILL item G-1.**

## §3 Run lifecycle end-to-end

**Start (operator `run start`, §10):**
1. Operator supplies the run config: the goal/task input, the provider lane selection, and the **operator-selected `credential_ref`** (m-10 writes it verbatim; never reads what it names — m-1 §1.4a).
2. **Manifest construction (§7)** → the **admission transaction**: ONE applier transaction writes `runs` (manifest bytes canonical + `run_manifest_digest` + state ADMITTED), mints **E=1 frozen**, and stamps the genesis rows (§9). **This commit is the manifest freeze point** — after it, the manifest for this `run_id` is immutable forever (r36 §C.1). → **GRILL item G-3.**
3. Broker CI-1 control session established (or already held); `epoch_state` for E=1 published per §B.4/§B.5 rules.
4. **Connector first** (r36 §B.4 step 4→5): spawn m-8, `hello` → `connector_assign` (the seven verbatim fields, §5) → `connector_ready` gates everything downstream.
5. **Worker generation machinery** (§4): allocate G1 → spawn → hello → lease-bind → `assign` → attach gate → first `turn_open` admission (§6).
6. Turns run under the frozen contract until a run-terminal state (COMPLETED/FAILED/CANCELLED) or operator stop.

**Stop (operator `run stop`):** the cancellation machinery (r36 §F `cancellations`, the D-5 composition) for a graceful path; `run cancel --hard` = retire the generation via the standard retirement transaction and mark the run CANCELLED. No bespoke teardown path exists — stop is the same machinery as supervision.

**App-main restart:** boot = **the recovery sequence, always** (there is no separate "clean boot" path — a clean store simply selects matrix branch (d)/initial-run): open store → integrity/`user_version` check → reconstruct → the r36 §B.3 recovery matrix (a)–(d) with the common suffix (control session first; pending-transition rules; broker install before lease-bind) → resume or await operator. Children of a previous incarnation are fail-closed by EOF (r36 §B.3); the design adds nothing to that contract, it just schedules the spawn-only path.

## §4 Worker-generation lifecycle (running §B.1/§B.2/§B.4)

The §B.1 machine runs in the supervisor as a per-generation state record (durable in `workers`; in-memory copies are caches of committed state, never authority):

- **Spawn:** sanitized env allow-list, no secrets/handles beyond the child's own channel endpoints (r36 §B.1); working/runtime dirs explicit. The spawn action itself is idempotent-safe: the generation row (ALLOCATED) commits BEFORE exec; a spawn failure ⇒ FAILED wash-out (§B.4 rule, same epoch).
- **Handshake:** `hello` within `HANDSHAKE_DEADLINE` → lease-grant transaction → `assign` (post-lease-bind, r36 R4-F1). The attach gate (D-2) then holds first-turn admission until the worker's `attach_result{attach-ok}` — the three m-7 tokens closed, unknown ⇒ fault (r36 §B.2).
- **Health:** `ping`/`pong` at `HEALTH_INTERVAL`; miss-past-deadline ⇒ FAILED disposition event.
- **Retire/replace:** FAILED/crash ⇒ the ONE retirement transaction (fence + park + E+1 + transition-ledger row, r36 §B.4) → §B.5 handshake with the broker → reap → next generation. Replacement is ordinary allocation under the new epoch.
- **Crash-loop honesty (→ GRILL item G-2):** each retirement is honest but unbounded replacement is not. Design: a **per-run consecutive-failed-generation counter** (durable, reset on any generation reaching a completed turn); at `MAX_CONSECUTIVE_GEN_FAILURES` (compiled constant, MVP default small) the run itself transitions **FAILED terminal** — no automatic revival; restart is an explicit operator act (a NEW run over the same store is not offered in MVP; the operator starts a new run). The counter and threshold are §2a-class termination safety, not policy.

## §5 Connector supervision + credential-reference orchestration

- Same SPAWNING/READY/FAILED/TERMINATED machine, no lease (r36 §B.1); `connector_assign` carries the **seven fields byte-verbatim from the frozen manifest** (r36 §B.1 — m-10 derives/authors nothing); READY only on `connector_ready`; the ordering gate (no admission/DATA-P/send before ready) is enforced by the scheduler refusing turn admission while the incarnation is not READY.
- **Co-restart:** a connector fault is never repaired alone — the §A.1 generation-paired co-restart via the standard retirement transaction (r36 §B.3). The supervisor has no "restart connector only" verb, by design.
- **Credential references:** the operator selects `credential_ref` at run start (§3); m-10 stores it in the manifest and copies it into `connector_assign`. **m-10 never resolves, opens, logs, or validates beyond presence/grammar** (r36 §C.1/§G.2); resolution is m-8's, inside the authorized attach only. The terminal surface displays the ref opaquely (its name, never a value). No secret-bearing column exists (§F) — this is structural, not procedural.

## §6 Scheduler — lease, epoch linearization, park/wake

- **One-active-turn lease:** `turn_open` admission = one applier transaction that checks (READY connector · leased generation at current epoch · attach-gate passed · no active turn · no run-terminal state) and writes the `turns` row + the active-turn lease together. The D-4 `parked_unknown` member is attached from committed state in the same transaction's snapshot (empty-array-never-absent).
- **Epoch linearization:** `turn_epoch` mints ONLY inside the retirement/transition transactions (r36 §B.4/§B.5) — which are applier events like everything else, so **mint, park, fence, and ledger-row are one serialized commit**; there is no second epoch authority anywhere in the process (r36 §B.4 source-specific rule realized by construction: the applier is the only writer of `epochs`).
- **Park/wake (F61, advisory):** wake relays arrive via the worker's `wake_forward` (r36 §E); the applier records `wake_schedule` with **UNIQUE(relay_id)**. **The at-most-once obligation (→ GRILL item G-4):** a wake may admit at most one turn, ever. Realization: turn admission from a wake happens in a transaction that (i) selects the `wake_schedule` row by `relay_id` with disposition `pending`, (ii) writes disposition `dispatched` + the `turns` row **in the same commit**. UNIQUE(relay_id) kills duplicate arrivals; the same-commit disposition flip kills double-admission from one row; a crash between arrival and admission leaves `pending` (recovered by rediscovery, still once). Dropped wakes are recovered by durable rediscovery, never queue growth (r36 §A.3/§E).

## §7 Run manifest — construction, freeze, and who reads what

Construction inputs, each with its authority: the **8-name tool-dispatch constant** (operator-ratified; compiled into the binary as the ONE constant the F55 serve gate and the manifest both cite — F58's policy-identity-vs-build-identity pair per r36 §C.1) · the m-3-produced `policy_digest` (m-10 writes verbatim) · the provider lane tuple + `lane_catalog_digest` from m-8's catalog (verbatim) · `credential_ref` (operator-selected, §5) · the F63 expected-catalog vector · run identity + E=1. m-10 **assembles and canonicalizes (JCS) but authors no semantic value** — every field is an operator choice, a sibling-owner digest, or a compiled constant.

**Freeze:** the §3 admission transaction. Before it, nothing exists to read; after it, readers get the frozen row only. The F55 serve gate (r36 §C.3) evaluates against the frozen row's exact canonical set at every `authorize_tool_call` — never against config files, env, or memory copies. `run_manifest_digest` (r36 §C.2) is computed at freeze and is the external E3 comparand (m-3's binding).

## §8 F59 host realization

r36 §D.1–§D.4 run as applier event types: authorize (the (0)–(7) procedure — one read+classify+conditional-insert transaction, reply after commit) · consume (the sender-fenced conditional update) · record (the (0)–(8) table + validated evidence) · expiry (the state-sensitive VOID mapping inside turn-end/retirement transactions). Two realization notes, zero new semantics: (i) the replay mapping is a **pure function of the stored row** — implemented as one lookup + one closed map, no branching on live state (that is what r36 (0) requires); (ii) the per-turn counter is the row COUNT (r36's ONE counter) — realized as a `COUNT(*)` under the same transaction, not a cached integer (no drift class exists).

## §9 Durable store — genesis, schema, recovery posture

SQLite per r36 §F (WAL, `synchronous=FULL`, private file, sole writer = the applier). **Genesis:** first boot creates the schema at `user_version=1` inside one transaction; a store with a HIGHER version than the binary ⇒ refuse to serve (fail-closed; operator message), LOWER ⇒ forward-only migration then serve (the v3 public-release rule). **Integrity:** a failed integrity check at boot ⇒ refuse + operator disposition — the MVP does not auto-rebuild a corrupt store (honesty over availability). The optional events journal stays deferred exactly as r36 §F pins (the chokepoint makes it one INSERT later).

## §10 Operator terminal surface (state-only; §8b respected)

Commands: `run start` (config + credential_ref selection) · `run stop` / `run cancel` · `status` (run/generation/turn/epoch/connector at a glance) · `attempts` / `tickets` / `parked` (the disclosure views: parked_unknown rows, attempt dispositions, ticket states — **ids, states, names, digests; never payload, never secrets** — I-PH respected) · `wakes` (schedule + dispositions). Every view is a committed-snapshot projection (§2). **No operator-authority ingress exists here**: the surface cannot clear parked rows, cannot forge dispositions, cannot touch the conductor (m-10 holds no verb) — the operator's direct route to seats stays outside this module entirely (§8b non-transitivity, r36 D-4).

## §11 E0 carriage

`pending_app_events` rows are m-3-schema'd (`m3.app_event.v1`) app-side events (attempt outcomes, phase records) written by the applier at their owning transitions (r36 §B.1); the worker seat carries them to the conductor (packet §3a — m-9 submits, m-10 never does). Durability: rows persist until the worker's carriage acknowledgment path (m-9's SITREP flow) marks them carried; no silent drop; the E0 floor is theirs to carry, mine to keep durable.

## §12 GRILL — the open items (operator, one at a time; durable GRILL_LOCK before review)

- **G-1 Store/commit discipline vs m-7's precedent:** same serialized-single-writer pattern, deliberately SEPARATE store/code (no shared engine). Shared: the discipline. Not shared: file, schema, engine code, failure domain. Is that the right cut, or should frank grow ONE reusable store engine now?
- **G-2 Crash-loop honesty:** consecutive-failed-generation counter → run FAILED terminal at a small compiled bound; no auto-revival; operator restarts explicitly. Right posture? Right that the bound is compiled (not config) in MVP?
- **G-3 Manifest freeze point:** freeze at admission (before connector spawn); NO post-start mutation of any manifest field — changing anything (even the lane) = new run. Acceptable operator ergonomics for the MVP?
- **G-4 Wake at-most-once:** the UNIQUE(relay_id) + same-commit disposition-flip proof (§6). Does the operator accept "a wake lost before durable arrival is recovered by rediscovery, possibly late" as the honest bound (no push guarantee)?

## §13 Prior-art reference lanes (PRIOR-ART §4; provenance, not vendoring)

- **deepagents** (async subagents as background runs): the background-run lifecycle SHAPE (spawn/observe/collect) informs §3/§4's run-and-generation separation; the trust model is NOT imported (deepagents trusts its subagents; frank fences them). No code vendored.
- **Talon** (per-conversation serialization; cron host): the one-loop-per-conversation serialization pattern is the §2 apply-loop's shape at different scale; the cron/wake host informs §6's park/wake realization (durable schedule + at-most-once dispatch). No code vendored.
- **jcode gate engine** (H-15 donor): the engine-enforced-gate idea lands here as **structural checks at admission**: manifest grammar/completeness and the 8-name set-equality are refused at the admission transaction (not advisory lint) — a cheap check made structural now. License check owed before any verbatim vendoring (none currently planned).

## §14 Fixture families owed at T4 (design-level names)

Boot: genesis · higher-version refuse · lower-version migrate · corrupt-store refuse · each recovery-matrix branch (a)–(d) with the §B.5 substates. Lifecycle: spawn-fail wash-out · handshake deadline · attach three-token legs · health-miss retire · co-restart pairing · crash-loop bound → run FAILED (G-2). Scheduler: one-active-turn refusal · wake at-most-once (duplicate relay · crash-between · rediscovery) · admission-while-connector-not-READY refused. Manifest: freeze immutability (post-admission mutation attempt fails structurally) · serve gate reads frozen row only · F63 vector mismatch. F59 host: the r36 fixture sets run against the real applier (not re-listed — r36 owns them). Surface: every view payload-free · no mutating verb beyond start/stop/cancel.

---
*Nothing in this document amends r36 or any consumed contract. GRILL_LOCK + the uniquely-parented m-10.implementer final-byte review + SITREP follow per the dispatch sequence.*
