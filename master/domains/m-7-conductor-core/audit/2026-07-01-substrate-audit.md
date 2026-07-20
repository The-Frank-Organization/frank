# m-7 Conductor-Core — c4 AUDIT: the runtime substrate (planner)

**AUDIT artifact** · Owner: m-7 (Conductor-Core) · Author: m-7.planner (design-lead) · Cycle c4 · 2026-07-01
**Dispatch:** `master/relays/c4-audit-m-7/AUDIT-orchestrator-planner-20260701-153627.md` (+ CF fold `master/relays/c4-audit-m-7/RECONCILE-orchestrator-planner-20260701-154248.md`)
**Phase/authority:** AUDIT, read-only. No design locked here; design questions are SURFACED for c4 DESIGN.
**Requirement set:** `master/DESIGN-REVIEW-2026-07-01.md` §2A · claim boundary per `master/GRILL-LOCK-deployment-fork-2026-07-01.md` D1–D5 · attach facts per `master/RUNTIME-RESEARCH.md` §8/§14.
**Provenance / method (CF-2 compliance):** planner-read governing corpus + all six locked m-1..m-6 design docs (contract inputs, read in full); live-instance probe (E2, planner-run); five read-only corpus-lens subagents for the planner's OWN sweep only (v2.8.8, jcode, claude-code, codex [operator-directed addition], external web prior-art) — lenses did not simulate, replace, or proxy-author the independent m-7.implementer audit (F4 stands). Load-bearing lens cites spot-checked against source by the planner (hook non-blocking header; jcode `write_bytes_inner`; jcode wire `from_session`). CF-1: full `master/...` paths used throughout.

---

## 0. BLUF

1. **The substrate is net-new — nothing anywhere in the corpus runs a serialized, crash-atomic, guarded relay store.** v2.8.8 is verbatim "a coordination protocol, not an orchestration runtime" (its README:7): a message format + a truth-agnostic linter, **zero process**. The LIVE `master/` team is the proof by demonstration: 261 INDEX rows with out-of-order and duplicate timestamps from unserialized seat appends, no lock, no daemon, seats (including this author) holding raw store paths.
2. **Every §2A mechanism has strong, citable external prior art** — single-writer commit loops (Redis/SQLite/LMAX/Kafka-leader), crash-atomic multi-file commit (SQLite super-journal, maildir, git ref-ordering), torn-write detection + quarantine (Kafka CRC/truncate), fail-closed fault disposition (Kubernetes admission `failurePolicy: Fail`), and quantified evidence that **removing a tool from the surface (not prompting against it) is the layer that actually stops confused-agent actions** (48–68% forbidden-tool selection under adversarial framing even with in-prompt allowlists). The engine is a **synthesis of proven parts**, not an invention.
3. **One refinement to the §2A sketch (surfaced for DESIGN, not a contradiction):** crash-atomicity of the multi-file commit comes from **one designated atomic filesystem operation as the commit pivot** (SQLite deletes one super-journal; maildir does one rename), with a redo journal + idempotent replay as the *recovery* half. The design must name its pivot explicitly.
4. **The internal prior-art splits cleanly:** jcode donates a battle-tested durable-write discipline (tmp→fsync→`.bak` hardlink→rename→dir-fsync + corrupt-primary→backup recovery) and idempotency-by-replay; claude-code donates per-recipient mailbox sharding + mark-read-after-deliver + priority dequeue — and both are textbook **negative examples** on exactly m-7's pillars (no serialization; forgeable self-reported identity; no multi-file transaction; silent total-loss on corrupt read).
5. **Claim boundary held:** the single by-construction claim this audit licenses is the **concurrency invariant** (one serialized critical section ⇒ no two honest racing seats both pass a check-and-burn). Everything interface-shaped is **confusion-resistant**; adversarial containment stays shelved (D3).
6. **Verdict: PRIMARY_BUCKET still-open** — proceed to c4 DESIGN with the seam inventory below; six contract under-specifications are flagged as targeted-COORD candidates (never silent reinterpretation).

---

## 1. Evidence discipline

- Planner-read docs and planner-run probes: **E1/E2** as marked.
- Lens-agent findings: **E1 (lens-collected)** — file:line cites present; three load-bearing cites re-verified verbatim by the planner; the rest spot-checkable at the cited lines. Per protocol, un-re-verified lens detail should be treated as E1-as-cited, not planner-witnessed.
- `master/RUNTIME-RESEARCH.md` facts: primary-source-verified + VP-reviewed upstream (§14); consumed as verified inputs.

## 2. The live instance today (E2 — planner-run probes, 2026-07-01)

The standing team's own store (`master/relays/`) is the running v2.8.8-protocol instance, and it demonstrates the absence of every §2A mechanism:

| §2A mechanism | live state | evidence (E2) |
|---|---|---|
| serialized commit | **absent** — seats append INDEX.md independently | 261 rows; ≥10 out-of-order timestamp runs; duplicate timestamps (awk monotonicity + uniq -d probe, this session) |
| crash-atomic multi-file commit | **absent** — relay file + INDEX row are two separate writes by the seat | this author's own boot ACK = Write file, then `printf >>` INDEX (two ops; crash between ⇒ record with no INDEX row) |
| single writer / store isolation | **absent** — store perms `rwxr-xr-x`, same uid for every seat session | `stat` probe; no `.lock`/`.tmp` files anywhere under `master/relays/` (find probe) |
| any conductor process | **absent** | `ps` probe: no daemon; delivery = operator hand-relay (charter, `/Users/jack/Programming/harness/CLAUDE.md`) |
| interface guardrail | **absent** — seats act on raw store paths with Write/Bash | this author is the evidence; relay-lint is run manually by the seat |

## 3. Corpus findings (per source)

### 3.1 v2.8.8 release (lens: sweep-v288; E1 lens-collected, hook header planner-verified)
Corpus: `extracted/agentic-dev-team-skills-v3-export/v2.8.8-release/…/agentic-dev-team-skills-v2.8.8/`. The entire executable substrate is `tools/relay-lint.py` (1449 lines) + a 35-line PostToolUse hook.
- **Commit path:** seat's own editor tools on a shared dir; INDEX append-only is convention (`protocol.md:386`), unlinted (hook skips INDEX; `tools/adapters/README.md:27,45`); no lock/rename/fsync anywhere (whole-tree grep).
- **Config:** no config artifact read by anything (argparse takes only paths/`--relay-root`/`--templates`).
- **File-surface restriction:** none; the hook is **non-blocking by design** ("This hook never hard-blocks", `relay-lint-posttooluse.sh:3-5` — planner-verified) and bypassed by Bash/heredoc/rename writes (`tools/adapters/README.md:45`).
- **Authority checks:** relay-lint run manually + agent compliance; lineage is "confusion-robust only" because `FROM`/`PARENT_DISPATCH_ID` are agent-authored (`protocol.md:88`); forged-FROM defense "designed, deferred" (`tools/adapters/README.md:52-54`).
- **Crash/partial-write:** nothing detects or repairs; corrupt files degrade to lint findings if someone runs the linter.
- **Delivery:** operator hand-relay; addressing is semantic metadata interpreted by the reading LLM.
- **PROMOTE:** the semantic contract layer — the canonical header schema, the dispatch-token/merge/lineage/visibility **check semantics** (146/146 fixture matrix), the append-only per-dispatch store layout. These become the conductor's validation pass content (largely via the locked m-2 dissolve/survive refactor).

### 3.2 jcode (lens: sweep-jcode; E1 lens-collected, two cites planner-verified)
- **Bus:** in-process multi-task tokio server, one task per connection over shared `Arc<RwLock<HashMap>>` maps — **no single-writer loop, no global commit section**; message/event history is a bounded in-memory ring (5000), lost on restart.
- **Identity (decisive negative):** trusted `ClientConnectionInfo` binding exists (`server/debug.rs:46-60`) but the message path trusts wire `from_session` verbatim (`crates/jcode-protocol/src/wire.rs:378-390` — planner-verified; handler takes the connection table as unused `_client_connections`). Any local process on the socket forges any sender.
- **Write path:** no burn-once; closest is **idempotency-by-replay** (content-hashed request key + persisted final response + single-flight set) — exactly-once by replay, a good retry-dedup pattern.
- **Crash-atomicity (jcode's strongest):** `crates/jcode-storage/src/lib.rs:261-384` — tmp(pid+nonce)→flush→optional fsync→`.bak` **hard-link** (no ENOENT window for concurrent readers — planner-verified)→atomic rename→parent-dir fsync; corrupt-primary→`.bak` recovery; durable/fast tiers; plus a "running at reload ⇒ crashed" restart disposition.
- **Tool surface:** narrow `swarm` tool stamps sender from context (right shape) but the bus socket does no authz and agents hold raw fs+shell over the plain-JSON store — no isolation, no guardrail.
- **PROMOTE:** the storage write discipline; idempotency-replay keying; restart disposition; the stamp-inside-trusted-layer *shape*. **ANTI:** identity, serialization, store isolation.

### 3.3 claude-code (lens: sweep-ccode; E1 lens-collected)
- **Inbox:** one JSON array file per recipient (`~/.claude/teams/{team}/inboxes/{name}.json`); senders rewrite the recipient's file directly under a per-inbox `proper-lockfile` (mkdir-based) with retry/backoff; default 10s stale-lock window.
- **Poll-and-inject:** 1000ms/500ms interval polling; **at-least-once** (mark-read strictly AFTER successful delivery — correct discipline); no cross-poll dedup for regular messages (fresh UUID per poll ⇒ duplicate injection possible); priority dequeue shutdown > team-lead > FIFO (anti-starvation).
- **Crash-atomicity (sharpest gap):** every write is an in-place `writeFile` over the live file — no tmp-rename; **corrupt read is swallowed to `[]`** (`teammateMailbox.ts:98-107`) ⇒ a mid-write crash silently vaporizes the entire inbox and the next writer cements it. No journal, no quarantine.
- **Multi-file consistency:** none — broadcast = N independent locked writes (partial broadcast on crash, no record); inbox vs task-state diverge with only ad-hoc forward repair.
- **Identity:** `from` is self-reported (`getAgentName()` from the sender's own context) and privilege checks are string-compares against it (`msg.from === 'team-lead'`) — concretely forgeable.
- **PROMOTE:** per-recipient sharding axis; mark-read-after-deliver; priority dequeue; serialize-not-fail-fast. **ANTI:** in-place writes + swallow-to-[]; per-file-lock-as-transaction; self-reported `from`.

### 3.4 codex (lens: sweep-codex; operator-directed corpus addition; E1 lens-collected)
Corpus: `references/codex/codex-rs/`. The strongest single internal donor. Orienting fact: state = two tiers — an **append-only JSONL rollout file per thread (source of truth)** + a derived, rebuildable SQLite index reconciled by read-repair/backfill.
- **Persistence:** strictly-append JSONL, `SessionMeta` header line first, full-replay resume (stream-parse every line into memory); fork = copy a truncated prefix into a fresh file with recorded `forked_from_id`/`parent_thread_id` lineage (`rollout/src/recorder.rs:933-996,1756-1780`; `thread_manager.rs:711-745`).
- **Serialization (PROMOTE, with elevation):** every rollout file is owned by ONE background task draining a bounded `mpsc(256)` with oneshot-acked barriers (`recorder.rs:1718-1754,895-931`) — exactly the commit-loop shape, but **per-file only**; no global cross-file commit ordering, which is the invariant m-7 adds.
- **Crash-atomicity (ANTI):** `write_all`+`flush()`, **no fsync on the append path** (grep-verified: `sync_all` appears in compaction/secrets/certs, never the recorder); no tmp-rename for appends; a torn tail line is **silently skip-and-counted** on resume, never quarantined; record and index are **eventually consistent by design** (read-repair `state_db.rs:491-626`) — the exact opposite of a crash-atomic record+INDEX+mailbox commit.
- **Burn-once (ABSENT):** approvals are *additive* session-wide prefix grants (`session/mod.rs:2014-2019`); no check-and-consume primitive anywhere — m-7 designs that fresh.
- **Config (PROMOTE):** layered TOML merged once at startup, not re-read mid-session, not a tool; **project-local (untrusted) layer is denylisted** from security-load-bearing keys (base URLs, providers, `notify`; `loader/mod.rs:57-70`); a **config-lock replays the effective config and hard-errors on drift with a diff** (`config_lock.rs:46-74`) — reproducibility/tamper-evidence prior art for trusted config load.
- **Tool surface (PROMOTE — near-exact guardrail template):** the surface is rebuilt **trusted-side per turn** as a pure function of config+provider+environment; two separate structures — `model_visible_specs` vs the executable `ToolRegistry` — with a 4-level `ToolExposure` incl. **`Hidden`** (registered, never advertised) (`spec_plan.rs:160-273`; `tool_executor.rs:13-42`); an unknown tool call is a hard dispatch reject (`registry.rs:442-461`). "Absent from what the model sees" and "not executable" are independently controlled — precisely the remove-don't-prompt-forbid layer.
- **Lifecycle:** restart drops all in-memory threads; recovery = replay-from-rollout only (no connection-binding recovery); interrupts are modeled as appended turn-boundary markers.
- **PROMOTE:** single-owner writer + acked barriers (elevate to global); append-only JSONL + header-first + replay + fork-with-lineage; ToolExposure visible/executable split; untrusted-config denylist + config-lock replay check. **ANTI:** flush-without-fsync; eventual-consistency two-tier store; silent skip of corrupt lines; per-file-only serialization; no burn-once.

### 3.5 srt / attach facts (from `master/RUNTIME-RESEARCH.md`, primary-source-verified + VP-reviewed)
- `srt` real, Anthropic-published; jail's only egress *path* = parent-owned broker socket, kernel-enforced (§14.1 A/B) — but **fails OPEN by default** on common kernels (§14.1 D) and the broker is the whole egress TCB (§14.1 C). Wrap is **shelved** anyway (GL D3); these facts matter to m-7 only as the shelved rung's documentation.
- **Identity is conductor-owned, full stop** (§14.2.2): no rented runtime supplies forgery-robust identity; the conductor stamps FROM from its own per-seat channel. This is a *simplification* the engine design inherits.
- The wake primitive is free: a persistent seat blocked on its stdin pipe is kernel-suspended; delivery = one `write()` (§2). The pipe is the inbox; no polling loop needed on the seat side.

### 3.6 External prior-art (lens: sweep-web; E1 lens-collected, sources cited inline in the lens report)
| mechanism | prior art | takeaway for the engine |
|---|---|---|
| single-writer serialized loop | Redis event loop; SQLite single-writer; LMAX "Single Writer Principle"; Kafka partition leader | buys check-then-act atomicity + consume-once + total order; throughput ceiling irrelevant at governance-relay scale |
| crash-atomic multi-file commit | write-tmp→fsync→rename→**fsync(dir)**; maildir rename-into-place (near-exact precedent for N per-seat mailboxes); git object-before-ref; **SQLite super-journal: commit = one atomic unlink** | **name the commit pivot**: one atomic FS op decides committed-vs-not; redo journal + idempotent replay is the recovery half (ARIES CLR: recovery-of-recovery idempotent) |
| torn-write detection / quarantine | Kafka per-record CRC32 + truncate-to-last-valid; git content-addressing | checksum per record; tail-truncate is standard; **interior corruption: fail-closed/park, not silent skip** (Kafka JIRAs show skip propagates loss) — a design decision, not a citable standard |
| internal-fault disposition | Kubernetes admission webhook `failurePolicy: Fail` ("failed-to-run ≠ said-no") | validator crashed/timed out ⇒ HOLD/reject, never accept; keep the fail-closed gate minimal + self-excluding so a gate fault can't brick the conductor's own recovery |
| narrow tool surface | "Prompts Don't Protect" (arXiv 2605.18414); MCP gateway allowlisting; OWASP MCP | **remove the tool from the reachable registry** — models pick forbidden tools 48–68% under adversarial framing, up to 37% even with an explicit in-prompt allowlist; surface-narrowing is the layer that neutralizes confusion; OS containment is a different (shelved) claim |
| SQLite vs append-files | SQLite's own when-to-use + faster-than-fs; agent-audit-journal work choosing append-only `.jsonl` over SQLite | store stays local append-files (locked, GL D2) — and that is the *documented* choice for auditability/git-diffability priorities; the cost (hand-built multi-file atomicity) is exactly what the pivot+journal pattern pays |

## 4. The §2A mechanism sweep — exists-to-promote vs net-new

| §2A mechanism | exists to promote | net-new (m-7 builds) |
|---|---|---|
| **serialized commit loop** (`submit()` read-validate-append + `verify()` check-and-burn, one critical section) | codex single-owner writer task + bounded channel + oneshot-acked barriers (per-file; elevate to global); jcode idempotency-replay as retry-dedup adjunct; pattern citations (Redis/SQLite/LMAX/Kafka) | **the global loop + burn-once** — nothing internal serializes across files, and no corpus has check-and-consume |
| **crash-atomic multi-file commit** (record + INDEX + N mailboxes) | jcode `write_bytes_inner` discipline (single-file); maildir/SQLite-super-journal/git patterns | **the multi-file transaction + the named commit pivot + redo journal + recovery scan** |
| **recovery / corrupt-record quarantine** | jcode `.bak` recovery (single-file); Kafka CRC+truncate pattern | **store-wide recovery/reconciliation + interior-corruption park (fail-closed for authority records)** |
| **internal-fault disposition** | K8s `failurePolicy: Fail` pattern; jcode "running⇒crashed" restart disposition | **the {accepted, rejected(terminal), HELD(fault)} taxonomy on the authority path — the design never picked (DESIGN-REVIEW finding 11)** |
| **trusted config load + integrity** | codex layered load-once + untrusted-layer denylist + config-lock replay validation (nothing else internal reads config at all) | **conductor-owned artifact, integrity hash in genesis, absent from every seat tool surface** |
| **interface guardrail under attach** | codex `ToolExposure` visible-spec/executable-registry split incl. `Hidden` + hard reject of unknown calls (near-exact template); MCP-server-side registry construction (industry pattern; strong external evidence) | **the enforcement**: seat tool surface = only `submit()`/`project()`/`read()`; raw store/config paths absent; + the negative fixture (§4.4) |
| **local-outbox-only external send** | m-3 egress gate contract (locked, dormant until first external send) | **the chokepoint hosting**: no seat-side send path; conductor's outbox is the only egress |
| **store genesis + GC/retention** | v2.8.8 store layout (promote) | **genesis record (config hash, schema_version) + retention policy — wholly absent** |
| **restart seat-binding recovery** | jcode disconnect-cleanup thoroughness (pattern); RUNTIME-RESEARCH checkpoint-resume facts | **persisted binding table + re-attach semantics (DESIGN-REVIEW §2C names the persisted seat-binding table before park/wake)** |

### 4.1 The concurrency invariant (AC2 — the one licensed "by construction" claim)
The two-honest-seats double-accept race: seats A and B both submit records consuming the same one-shot authority (an away-token choice, a delegated dispatch, a merge grant). With per-resource locks or a multi-writer store, both can pass "check: token unused" before either writes "burn" (TOCTOU). **The kill is structural:** the conductor is one process; ALL store mutations — `submit()`'s read-validate-append and `verify()`'s check-and-burn — execute on **one single-threaded commit loop**; each request runs read→validate→append **to completion before the next begins**. Two racing submits serialize; the second's validate step sees the first's burn; no interleaving exists in which both pass. This is a property of the trusted engine's own control flow — it holds **regardless of seat behavior**, which is why it is claimable "by construction" (dispatch AC2's explicit license) while every interface claim stays confusion-resistant. It is also precisely what the prior art buys (Redis WATCH/MULTI; SQLite single-writer; LMAX). Cost: single-core write ceiling + head-of-line latency — irrelevant at relay scale (261 records in 4 days on the live team).
**Boundary honesty:** the loop serializes the *store*; it does not stop a code-executing seat writing the store file directly under attach — that is the D5 accepted-risk, claim-swept per GRILL-LOCK D4. Same-uid direct writes are adversarial-lane, out of scope.

### 4.2 Crash-atomicity + recovery + fault disposition (AC3)
Auditable shape (surfaced for DESIGN, not locked): per-record **checksum**; a **redo journal** written+fsynced first (intent: record bytes + INDEX row + mailbox deliveries); a **single named commit pivot** (one atomic rename or unlink whose presence/absence decides committed-vs-not); idempotent replay of the journal on startup; **tail truncation** for torn tails; **interior corruption or a thrown/timed-out trusted check on an authority record ⇒ HELD/fail-closed** (a distinct terminal-visible state: never silent-accept, never brick — the store keeps serving; the held record parks to the operator surface). jcode's single-file discipline (tmp→fsync→bak-hardlink→rename→dirfsync) is the per-file primitive to promote inside this; claude-code's swallow-corrupt-to-[] is the named anti-pattern.
**Refinement flag (from the external sweep):** §2A's "journal + idempotent redo" under-specifies — atomicity comes from the pivot op, the journal is recovery. Design the pivot explicitly.

### 4.3 Trusted config load + integrity (AC5)
Requirement (GRILL-LOCK D2b + DESIGN-REVIEW §2A.2): the policy-config artifact (m-6/m-3/m-4-authored: gate_category maps, egress rules, capability priors, archetype registry) is **conductor-owned, loaded once at trusted startup, integrity-checked** (hash recorded in the store — genesis/GC seam), and **absent from every seat tool surface**. Nothing in any corpus does this. Config *reload* semantics (SIGHUP vs restart-only) and integrity-mismatch disposition (fail-closed startup refusal vs held) are DESIGN questions (§7). Claim: confusion-resistant (a confused seat has no tool naming the config path); a malicious same-uid edit is the D5 accepted-risk.

### 4.4 The interface guardrail under attach (AC4)
**HOW held:** the conductor runs as an MCP server; each seat's *governance* tool surface is the tool registry the conductor presents on that seat's dedicated channel — **constructed trusted-side, per seat**. The registry contains exactly `submit()` (renders the m-2 form: forbidden options absent), `project()`, `read()`. Raw store paths and the config path **appear in no tool schema, no tool description, no resource listing** — a confused seat following "edit the relay store directly" finds no such tool to call. This is the arXiv-2605.18414 layer: remove, don't prompt-forbid.
**The honest residual (D4/D5):** under attach the seat's *host runtime* still hands it Bash/file tools; a confused-*with-a-shell* seat can be instructed around the MCP surface. Mitigations available without wrap (store dir outside seat workspaces; no store/config path ever printed into seat-visible prose) are hygiene, not containment: the claim stays **confusion-resistant**, malicious code-executing agent explicitly out of scope.
**Negative fixture shape (dispatch AC4):** (i) enumerate a seat's presented tool registry ⇒ assert exactly {submit, project, read} and zero path-bearing tools/resources; (ii) drive a scripted confused-seat turn that *attempts* a raw store write via its governance surface ⇒ assert no tool exists to express it and the turn's only store effect is via `submit()`; (iii) assert the config path string appears nowhere in any seat-deliverable projection. *(NOT the old "lane write to a config path fails closed" isolation fixture — that claim collapsed to 2B.)*

### 4.5–4.7 Local-outbox / genesis+GC / restart recovery
- **Local-outbox-only:** no seat has an external-send tool (m-5 locked: no v3.0 `external_send` axis); the only egress is the conductor's outbox at the m-3 §7 egress gate (dormant until the m-6 away bridge). m-7 hosts the chokepoint; net-new.
- **Genesis/GC:** store init = genesis record carrying config hash + schema_version + address-space seed; retention/GC policy operator-config. Wholly net-new; surfaced.
- **Restart/away touch-points (AC7, surfaced not designed):** (a) persisted seat-binding table (seat ↔ channel/credential) so a conductor restart re-binds persistent seats without re-mint (DESIGN-REVIEW §2C names it before park/wake); (b) parked-lane state (m-6 7-state machine rows) must be recoverable from the store alone — the store IS the checkpointer (locked m-6 §4); (c) in-flight `submit()` at crash resolves via the §4.2 journal (committed or not; never half); (d) away-token nonce-burn executes inside the same commit loop (Seam C, m-1-owned contract); (e) re-observe-on-wake ordering (m-6 locked) is a sequencing obligation on the loop.

## 5. SEAM INVENTORY (AC6 — precursor to the DESIGN seam matrix)

Format: `{contract owner · contract doc/section · m-7 execution obligation}` + under-specified flag. Full biting-negative-fixture matrix + contract-question-raised? column = DESIGN-lock deliverable.

| # | contract (owner) | contract doc/section | m-7 execution obligation | under-specified? |
|---|---|---|---|---|
| S1 | store append + channel-stamped FROM (m-1) | `master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md` §5 (submit/project/read/mint_seat, :97-107), §6 (on-disk + system fields), §4 (DI-1..DI-5) | run the four verbs; stamp FROM/ROLE from the channel binding; one atomic accepted-append + INDEX + mailbox deliver inside the commit loop; terminal `rejected` on fail; serve `read` to the lineage engine | no — but binding-table persistence across restart is CQ-6 |
| S2 | away-token mint/verify, Seam C (m-1) | m-1 doc §C3.4 refs + `master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md` §4 (:85-88), §12 carries | execute `verify()`'s five checks + **nonce-burn as an atomic conductor append in the same critical section as submit** (the double-accept kill, §4.1); sibling-burn per DESIGN-REVIEW §2C | sibling-burn decision-scoped resolution named in §2C but not yet in locked m-1 text — CQ-6 |
| S3 | form render + fill-time authority (m-2) | `master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md` §3 (render semantics), §4 steps 1–6 (:68-74) | render the form trusted-side per (seat, phase, tier, slot): forbidden options **absent**; validate constrained picks (set-membership / monotonic ≥ floor); courier-fill system fields | no |
| S4 | form-validation + cross-relay lineage gate (m-2) | m-2 doc §4 step 4–5, §10b/§10c, §11; canonical seam `master/ARCHITECTURE.md:58-66` | run (a) form-validation then (b) lineage walk over `persisted ∪ {candidate}` **in-courier, pre-append, nothing persisted**; blocking edges only for authority-bearing records; bounce names the failing field/edge | no |
| S5 | phase-split required-set (m-1/m-2, joint w/ m-7) | m-1 §5 Step-1 build boundary (:104); m-2 §4 step 4c; requirement `master/DESIGN-REVIEW-2026-07-01.md` §2A.5 | Step-1 form gate must not demand observe-owned fields (`ACTIONS_GIT_REF`/`FINAL_GIT_STATUS_SHORT`) that have no Step-1 writer | **YES — CQ-1.** Locked docs reserve the observe hook to Step-2 but do not visibly step-gate the observe-owned `required_when` predicates; §2A.5's fix (step-gate OR conductor-side filler) is not yet in m-1/m-2 text |
| S6 | observe-as-send hook (m-3) | `master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md` §2, §3 (:41-64) | host `observe_gate()` inside atomic submit (Step-2 hook-point named now, inert in Step-1); enforce the positive write-allowlist; bind the passing observation to the gated record (TOCTOU) | no (Step-2 hook; ordering named) |
| S7 | decision-② fail-closed on authority-class `self_reported` (m-3 contract; operator decision ②) | m-3 doc §3.2 (:63) currently says delivery **never** gates on evidence integrity — universal fail-open; requirement `master/DESIGN-REVIEW-2026-07-01.md` §2A.7 | execute class-conditional fail-closed at the gate for authority-class records | **YES — CQ-2.** The fold of ② into locked m-3 §3.2/§8/§12 (re-baseline step c, CTO-owned) has not landed; m-7 cannot execute against contradictory text |
| S8 | pure-judgment A-floor (m-2/m-6 contract) | requirement `master/DESIGN-REVIEW-2026-07-01.md` §2A.6; nearest locked anchors: m-2 §3 monotonic HUMAN_GATE floor + `master/ARCHITECTURE.md` §J2 | enforce the mandatory HUMAN_GATE floor by (phase × record_kind) at fill/submit; `CEREMONY_TIER` monotonic-with-system-floor (below-baseline pick auto-sets `gate_category=ceremony_downgrade`) | **YES — CQ-3.** The floor *table* (phase × record_kind) has no locked author yet; m-7 enforces, m-2/m-6 must own the table |
| S9 | egress / content-safety gate (m-3) | m-3 doc §7 (:116-122) | host the fail-closed scan at the sole-external-send chokepoint (the conductor outbox); dormant until the m-6 away bridge | no |
| S10 | routing record + route_dispatch fail-closed (m-4) | `master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md` §3 (:144-163), §5, §7 | sequence routing-record acceptance through submit like any relay; host `route_dispatch()`'s fail-closed refusal; never a silent default model | no |
| S11 | template spawn + pane-spawn (m-4/m-5 contracts; conductor-core named executor) | m-4 §0 GL-4 + §7 (:236-263); `master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md` §7 (:116-127), §10 | open/name panes via existing tmux/zellij/OS-terminal (Step-1-consistent); deliver boot relays; record per-assignment `seat_archetype`+`authority_ceiling` at spawn (F2 home is m-4's record) | mild — pane-spawn mechanics are Step-1 build detail; no contract gap |
| S12 | `slot_in` conductor-classification at acceptance (m-3/m-5) | m-3 doc §5.1 (:100-103; pipeline point explicitly "a PLAN detail"); m-5 §4 | classify work-archetype at work-record acceptance **inside the commit loop**, immutable thereafter, non-lane-writable; done-predicate reads it | **YES (mild) — CQ-5.** The pipeline point is m-7's to sequence; needs a joint ordering statement (classify-before-observe-predicate-selection), not a policy change |
| S13 | scheduler park/wake, 7-state machine (m-6) | m-6 doc §4 (:71-89) | persist state transitions as store records (the store IS the checkpointer); wake = one write to the seat's pipe; re-observe-on-wake sequencing; resummon timers | no — timer substrate (who fires resummon_due) is engine detail to design |
| S14 | delivery/projection buckets + ODB (m-6) | m-6 doc §2 (:31-49), §3 | deliver accepted records into TO/CC mailboxes; project buckets off locked fields; render ODB slots (m-2 §17.2) | no |
| S15 | trusted config artifact (m-6/m-3/m-4 author; m-7 loads) | `master/GRILL-LOCK-deployment-fork-2026-07-01.md` D2b; DESIGN-REVIEW §2A.2; §J config surfaces in `master/ARCHITECTURE.md` §J2 | load once at trusted startup; integrity-check; expose to policy engines internally; **absent from every seat tool surface** | **YES (mild) — CQ-4b.** Config artifact *format/composition* (one artifact vs per-domain files) has no owner statement; m-7 needs the load contract |
| S16 | terminal-state token unification (m-2/m-3/m-6) | m-2 doc §17.1 Q-E (:266); `master/ARCHITECTURE.md:287` held un-swapped (`bounced` vs `rejected`) | the engine's record-state enum must be closed: {accepted, rejected, held?, …} — m-7 executes whichever token set the owners settle | **YES — CQ-4.** Open SHOULD-fix (CTO-routed); becomes engine-blocking at m-7 DESIGN because the state enum is the loop's core type |

**Contract questions raised (targeted-COORD candidates at DESIGN; static consumption otherwise):**
- **CQ-1** (m-1+m-2+m-7 joint): phase-split required-set step-gating — §2A.5 fix not yet visible in locked text (S5).
- **CQ-2** (m-3; actually re-baseline step c dependency): decision-② fail-closed fold outstanding; m-7 design will state the class-conditional disposition and cite ②, but the m-3 text must be brought current before m-7's design-lock can claim byte-consistency (S7).
- **CQ-3** (m-2/m-6): the pure-judgment A-floor table needs a policy owner; m-7 owns only enforcement (S8).
- **CQ-4** (m-2/m-3/m-6): terminal-state token set (Q-E `bounced`→`rejected` unification + the new HELD fault state m-7 must add) — the engine's state enum forces this open item (S16); **CQ-4b**: config artifact composition/format owner (S15).
- **CQ-5** (m-3/m-5): `slot_in` classification ordering inside the commit pipeline (S12).
- **CQ-6** (m-1): persisted seat-binding table + away-token sibling-burn restart semantics (S1/S2; DESIGN-REVIEW §2C).

None of these reopens a locked contract; each is either an already-tracked open item the engine makes load-bearing (CQ-2, CQ-4), a §2A/§2C requirement not yet folded (CQ-1, CQ-6), or a missing policy-owner statement m-7 must consume (CQ-3, CQ-4b, CQ-5).

## 6. 4-bucket verdict (AC1)

```text
PRIMARY_BUCKET: still-open
still-open: the ENTIRE §2A substrate — serialized commit loop, crash-atomic multi-file commit + recovery +
  quarantine, internal-fault disposition, trusted config load, interface-guardrail enforcement,
  local-outbox-only send, genesis/GC, restart seat-binding recovery. No internal system has ANY of it
  (v2.8.8: a linter, zero process — verified; live master/: unserialized seat appends — E2; jcode: multi-task
  RwLocks + forgeable wire identity + in-memory message log; claude-code: in-place writes + swallow-corrupt +
  self-reported from). Net-new design work; this is what c4 DESIGN produces.
already-closed: none — no existing conductor exists to promote or enable. (Distinct from PROMOTE-parts, which
  are components consumed INSIDE the net-new engine: v2.8.8 semantic contract layer via the locked m-2
  dissolve/survive refactor; jcode single-file write discipline + idempotency-replay + restart disposition;
  claude-code mailbox sharding + mark-read-after-deliver + priority dequeue; codex single-owner writer +
  ToolExposure split + config denylist/config-lock; the m-1..m-6 locked contracts
  themselves — the engine hosts them, builds none of their policy.)
product-overlapped: none material. Wrap/adversarial isolation is SHELVED (GRILL-LOCK D3), not overlapped;
  the six policy domains are HOSTED, not overlapped (the one-line boundary); existing runtimes (Claude Code /
  codex app-server) are the RIDE surface Step-1 attaches to, not competing conductors.
recommended-next: proceed to c4 DESIGN of the minimal engine (the dispatch's design question), carrying:
  (a) the §4 mechanism shapes incl. the named-commit-pivot refinement; (b) the §5 seam inventory → full seam
  matrix with biting negative fixtures; (c) CQ-1..CQ-6 as targeted COORDs / tracked dependencies (CQ-2, CQ-4
  block byte-consistent design-lock, not design start); (d) GRILL before design-lock (VP constraint 3).
```

**Duplicate/already-built gate:** checked against v2.8.8 (no runtime — whole-tree grep), the live instance (E2 probes), jcode, claude-code, codex (app-server drives seats; it is a ride target, not a governance conductor), and the external field (MCP gateways do allowlisting, not governed stores; no surveyed system runs a serialized crash-atomic governance store with channel-stamped identity). Nothing to promote-in-place; no rebuild risk.

## 7. Design questions surfaced for c4 DESIGN (not locked here)

1. **The commit pivot** — which single atomic FS op is the linearization point for {record, INDEX, N mailboxes}; redo-journal shape; per-record checksum format.
2. **INDEX as derived projection?** If INDEX.md is rebuilt from records (a view), the multi-file commit shrinks to record+mailboxes+journal — simpler pivot. (Human-readability requirement says keep INDEX on disk; question is whether it is *committed* or *derived*.)
3. **Mailbox realization** — per-seat maildir-style files (rename-into-place, no lock) vs pure `project()` query over the store (m-1 doc §11 left both open); the engine's crash story differs.
4. **Process/concurrency shape** — single process: N per-seat channel tasks (I/O only) feeding one commit-loop thread (LMAX shape); `verify()` and `submit()` on the same loop (§4.1); observe/probe execution (Step-2) inside-vs-adjacent to the critical section (TOCTOU vs loop latency).
5. **Fault taxonomy** — {accepted, rejected(terminal), **held(fault)**} + quarantine; held-record operator surface (m-6 bucket A?); self-exclusion so gate faults can't brick recovery (K8s lesson).
6. **Config** — artifact composition (CQ-4b), hash-in-genesis, reload semantics (restart-only recommended), integrity-mismatch disposition (refuse-start vs held).
7. **Guardrail realization detail** — MCP server per-seat tool registry; form render as `submit()`'s tool schema (fill-time authority = the tool's input schema per seat/phase); the three-part negative fixture (§4.4).
8. **Genesis/GC** — genesis record contents; retention policy surface (operator config).
9. **Restart** — binding-table persistence (CQ-6); parked-lane recovery from store; journal replay ordering vs seat re-attach.
10. **Timer substrate** — who fires `resummon_due`/digest cadences (the one place the engine needs a clock; keep it out of the critical section).

## 8. Risks / reject-or-narrow gates

- **Over-reach risk (the boot's named failure mode):** several §5 rows border policy — the A-floor table (S8), the terminal-token set (S16), config composition (S15). m-7 must consume/flag these (CQs), never author them. This artifact deliberately assigns each CQ a policy owner.
- **Under-reach risk:** the timer substrate (§7.10), the held-record surface (§7.5), and pane-spawn (S11) are engine-homeless unless m-7 claims them — claimed here as m-7-owned engine detail.
- **Claim-boundary regression risk:** DESIGN must repeat the D4 claim verbatim wherever the guardrail is described; this artifact's §4.1/§4.4 wording is the template.
- **Dependency risk:** CQ-2/CQ-4 are open items owned elsewhere (re-baseline step c; Q-E). If unresolved by m-7 design-lock, the lock inherits a byte-inconsistency — flag to orchestrator at reconcile.
- **Narrow gate:** if DESIGN discovers the multi-file commit cannot keep INDEX.md human-append-readable AND crash-atomic, the fallback (derived INDEX, §7.2) changes the on-disk contract m-1 locked (§6 "reuse v2.8.8 layout") — that would be a targeted m-1 COORD, not a silent change.

## 9. Questions for m-7.implementer (pair reconcile) + operator

- **To m-7.implementer:** (a) does your independent sweep confirm the absence claims (especially: no serialization anywhere in v2.8.8/jcode/claude-code write paths)? (b) do you concur the licensed by-construction claim is scoped exactly to §4.1's control-flow property? (c) any hosted contract I mis-bucketed as under-specified vs consumable-as-is? (d) over-reach check: do any of my §7 design questions silently re-own policy?
- **To operator (non-blocking, surfaced only):** none requiring decision at AUDIT. CQ-2/CQ-4 land with their existing owners; the grill (VP constraint 3) is where substrate semantics get operator time.

*(Read-only audit; no source/`pcode/` writes. Relay of record: `master/relays/c4-audit-m-7/AUDIT-planner-<ts>.md`.)*
