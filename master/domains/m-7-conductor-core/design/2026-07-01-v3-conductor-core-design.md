# v3 Conductor-Core — DESIGN (the engine)

**DESIGN_DOC_ID:** c4-design-m-7-conductor-core
**Owner:** m-7 (Conductor-Core) — design-lead m-7.planner · adversarial design-reviewer m-7.implementer
**Cycle/phase:** c4 / DESIGN (design-lock terminal; no PLAN/IMPL) · **Tier:** medium · **Evidence:** E1 · **Date:** 2026-07-01
**Status:** **DESIGN-LOCKED** (r5) — **VP design-lock co-sign `master/relays/c4-design-m-7/RECONCILE-orchestrator-reviewer-20260702-040327` (`VP_DESIGN_LOCK_CO_SIGN: approve`); c4 CLOSED 2026-07-02.** Design pair-approved at r3 (`master/relays/c4-design-m-7/DESIGN-REVIEW-implementer-20260702-004452.md`); **LOCK PACKAGE pair-approved at r5** (m-7.implementer `DESIGN_REVIEW_VERDICT: approve` on DESIGN_LOCK_ID `c4-design-m-7-lock`, `master/relays/c4-design-m-7/DESIGN-REVIEW-implementer-20260702-035245.md`); **CQ gate SATISFIED** (all eight design-LOCK CQs closed — §15; CQ-6 on the base with `re-mint-supersedes` carried non-locking); the three CTO fold-integration items applied (§21 r4 entry; r5 = stale-CQ-5-wording purge). **DESIGN_LOCK block at §22 — EFFECTIVE** (m-7.implementer lock-package approve `…-035245` + VP co-sign `…-040327`). Lock-gate checklist: (a) seam matrix biting (§12 — done), (b) GRILL_LOCK (§14 — done), (c) CQ ledger closed or explicitly carried non-locking (§15 — **DONE**: all eight closed; CQ-8 was surfaced by this design's grilled Q1 pivot and closed by the m-1 COORD), (d) F8 claim-sweep clean (§16 — done, r3+r4).
**Basis:** VP-approved reconciled c4 audit (`master/domains/m-7-conductor-core/audit/2026-07-01-substrate-audit.md` + `master/domains/m-7-conductor-core/audit/2026-07-01-audit-pair-reconcile.md`, implementer-confirmed `master/relays/c4-audit-m-7/RECONCILE-implementer-20260701-161137`); dispatch `master/relays/c4-design-m-7/DESIGN-orchestrator-planner-20260701-162721.md`. Per the standing guardrail, audit-surfaced mechanism shapes were treated as **hypotheses**; each is proven here against donors + fixtures, and the six hard-to-reverse decisions were operator-grilled (§14).
**GRILL_REQUIRED:** yes — run this cycle; GRILL_LOCK_ID `c4-grill-m-7` (§14).

---

## 0. Frame + locked inputs

**The one job.** The conductor-core is the ENGINE: the running program the six locked policy domains ride on. It EXECUTES their contracts — right order, right atomicity, behind the right interface — and owns no policy. This doc covers `master/DESIGN-REVIEW-2026-07-01.md` §2A in full.

**Claim boundary (LOCKED, held in every sentence of this doc):** Step-1 = attach + interface guardrail = **CONFUSION-RESISTANT** (`master/GRILL-LOCK-deployment-fork-2026-07-01.md` D1–D5). The **single** "by construction" claim in this design is the serialized-loop kill of the two-honest-seats double-accept race (§2.4) — a control-flow property of the trusted engine. A malicious same-uid code-executing seat can write the store files directly; that is the documented D5 accepted-risk, out of scope. No sentence in this doc claims "unbypassable / sole-writer by construction / same-uid write-exclusion" (§16 claim-sweep).

**Locked contracts consumed (never reopened):** m-1 verbs/stamp/store shape + DI family; m-2 FieldSpec/fill-time authority/write-path seam (`master/ARCHITECTURE.md §3 "Write path — reconciled seam"`); m-3 observe-as-send + allowlist + egress; m-4 route_dispatch/R2; m-5 archetypes/ceilings/templates; m-6 buckets/ODB/park-wake/away-bridge. Genuine contract questions ride the CQ ledger (§15) as orchestrator-fired COORDs — never silent reinterpretation.

## 1. The engine at a glance

One conductor process. Five parts:

```text
 seats (persistent runtime processes, one per seat)
   │  per-seat MCP channel (stdio pipe / bearer token)   ← identity = the channel (m-1)
   ▼
 [A] CHANNEL LAYER      N concurrent per-seat handlers; present the 3-tool registry;
                        stamp-source; hand commands to the intake-writer; await outcome
   ▼ bounded channel → SINGLE intake-writer task (append+fsync; §2.1)
 [B] INTAKE JOURNAL     durable, append-only, one writer; the FIFO's persistent half (§2.2)
   ▼ in arrival order
 [C] COMMIT LOOP        ONE thread; plain FIFO; runs EVERY mutation start-to-finish:
                        submit validate+append, verify check-and-burn, park/wake, slot_in
                        classification, outbox enqueue, genesis, GC          (§2, §3)
   ▼ staging → fsync → RENAME PIVOT → dir-fsync → projections               (§4)
 [D] STORE              canonical records (append-only, checksummed, immutable) = truth;
                        INDEX.md / rendered .md / mailboxes / burn-set = derived projections;
                        genesis record; quarantine/; intake+redo journals    (§4, §10)
   ▼ reads (concurrent — immutability makes them safe)
 [E] READ/DELIVERY      project()/read() serve committed records; delivery = one write()
                        onto the recipient seat's pipe (wake); local outbox = the
                        conductor-governed external egress, behind the m-3 gate
                        (D5 same-uid residual — §9)                          (§8, §9)
```

Trusted config is loaded once at [startup], digest-checked against genesis, and lives only inside [C]/[E] — never on a seat surface (§7).

## 2. Process/concurrency model (GRILL Q2 — locked)

### 2.1 One writer, many readers
- **Channel handlers ([A]) are mail carriers only**: they authenticate the channel (m-1 binding), accept a command, and **hand it over a bounded in-process channel to the single INTAKE-WRITER task**, then await the outcome. **Handlers never touch any file** — the intake journal has exactly ONE writer: the dedicated intake-writer task, which appends+fsyncs entries in channel-arrival order and feeds the commit loop (c6 m-7-F5 fix: the single-writer discipline is now stated at the intake tier too, resolving the apparent "handlers never touch store files, yet who appends intake?" gap). Donor: codex's single-owner writer task + bounded-channel + oneshot-acked barriers (`references/codex/codex-rs` rollout recorder) — applied at BOTH tiers (intake-writer for the journal; commit loop for the store); the LMAX single-writer principle.
- **The commit loop ([C]) is one thread** consuming a **plain FIFO** (operator-grilled: no priority classes in v3.0; claude-code's priority-dequeue recorded as a rejected-for-now future knob). Every mutation runs to completion before the next begins.
- **Reads are concurrent and lock-free** because committed records are immutable — `project()`/`read()` never race with anything. Readers of append-in-progress projection files see a shorter prefix (safe). Form renders are concurrent reads too: the render is advisory, the loop's in-critical-section validation is authoritative (locked m-2 §3/m-1 §13.2 — stale renders are caught at submit; TOCTOU stays closed).

### 2.2 The durable FIFO (operator riders, grilled)
- Every command is **appended + fsynced to the intake journal (by the single intake-writer task, §2.1) BEFORE entering the in-memory FIFO**; the seat's ack arrives only with a typed outcome. Power loss loses at most a command not yet fsynced to intake — and the seat, holding an un-acked candidate on a persistent session, retries; content-hash request keys (jcode idempotency-replay donor) dedupe the retry.
- **Clear-on-pop is atomic with the outcome commit** (operator rider — the stale-re-emission failure mode is structurally closed): every intake entry has an `intake_id`; every outcome record (`accepted`/`rejected`/`held`) references the `intake_id` it consumes; "consumed" is durably recorded by the same §4 pivot that records the outcome — no separate mark-consumed write exists to be missed. Recovery re-enqueues exactly `intake − outcomes` in arrival order (§5 phase 3.5). At-least-once intake, exactly-once effect, zero stale re-emission.

### 2.3 What waits on the loop
Step-2's observe probe runs **inside** the atomic submit (locked m-3 §2 — the passing observation binds to the gated record). A slow probe stalls the loop; the design's answer is **hard per-check timeouts → internal fault → §6 disposition** (class-conditional per CQ-2: authority-bearing → `held`; non-authority → `rejected`/author-return), never accept-without-observing, never hang. Cost accepted and recorded: during a probe other submits queue (live-team volume: 261 records/4 days).

### 2.4 The licensed by-construction claim (scoped exactly)
Two honest racing seats submit records consuming the same one-shot authority (away-token choice, delegated dispatch, merge grant). Because ALL mutations — `submit()`'s read-validate-append and `verify()`'s check-and-burn — execute on one single-threaded loop, each to completion, there exists **no interleaving** in which both pass the unused-check before either burns. The second command's validate step reads the first's committed burn. This holds regardless of seat behavior — it is the trusted engine's own control flow. **This control-flow property is the only "by construction" claim this design makes.** It does not constrain a code-executing seat writing store files directly (D5 accepted-risk).

## 3. The commit pipeline — executing the locked write-path seam

`submit()` on the loop, in the locked order (`master/ARCHITECTURE.md §3 "Write path — reconciled seam"`; m-1 §5; m-2 §4):
1. **resolve + stamp** — seat from channel binding (reject unbound); stamp `FROM`/`ROLE`; fill system envelope (`relay_id`, `DISPATCH_ID`, `timestamp`, `schema_version`, `certification=null`); validate `PARENT` ∈ conductor-derived candidate set; `TO`/`CC` ∈ minted address space.
2. **pre-append validation, candidate in-courier, nothing persisted** — (a) m-2 form-validation (required-set, enums, seat-scope); (b) m-2 lineage gate over `persisted ∪ {candidate}` (blocking edges only for authority-bearing records).
3. **`slot_in` classification — the CLOSED CQ-5 ordering** (m-5 `c4-cq-slotin/DESIGN-planner-20260702-014506` + m-5.impl `020448` + m-3 co-sign `024732` + m-3.impl `030205`): conductor-classified **at acceptance, post-gate / pre-observe / atomic-bind** — after the form/lineage gates, before the observe hook, so the (Step-2) observe predicate selection reads the just-classified in-courier value, and the classification is bound into the same atomic commit as the observation it selects. Immutable, non-lane-writable **at the tool surface** (no verb exists for a lane to set or change it; D5 same-uid residual per §8.4 — locked F1's claim grain); required by locked m-5 §4. *(History: r1 had classify-after-observe — contradiction caught in review; r2 proposed this ordering to the CQ-5 COORD; the COORD closed on exactly it — lock-bearing as of r5.)*
4. **[Step-2 hook, inert in Step-1]** m-3 `observe_gate()` under the §2.3 timeout, its predicate selected by the step-3 classification, its write-allowlist mechanically enforced (writes outside the closed set = internal fault).
5. **atomic commit** — §4: stage (record carrying classification + observation) → checksum → fsync → **rename pivot** → dir-fsync → redo-journal-driven projections (INDEX row, rendered `.md`, mailbox entries) → outcome references `intake_id`.
6. **deliver** — one `write()` per recipient pipe (§8.3). On gate failure at 2 or 4: **one atomic append of a terminal `rejected` evidenced record** (same pivot mechanics) + bounce naming the failing field/edge. On internal fault: `held` (§6).

`verify(token)` on the same loop: m-1's five checks (sig → audience → expiry → nonce-unused → seat-match) → on pass, **ONE compound operator-verdict record** committed via the same single §4 pivot, stamped `FROM: operator` via the operator channel — **the verdict record's presence IS the decision-scoped `(decision_id, seat)` burn** (m-7-F1, VP-ratified: burn+verdict merged into one canonical record; no separate burn record, no second rename). The burned-nonce set is *derived* from committed operator-verdict records. Double-redemption is killed by §2.4.

## 4. Crash-atomic multi-file commit — the named pivot (GRILL Q1 — locked)

**Package A, operator-locked.** The commit set per accepted relay = {canonical record, INDEX row, rendered `.md`, N mailbox entries, burn/outcome markers}. Atomicity design:
- **Canonical record = the ONE truth.** The typed-envelope record (m-2's canonical object; format: one self-contained checksummed record file under `store/records/`, naming carries `relay_id`) is staged in `store/staging/`, `fsync`ed, then committed by **one atomic `rename()` into `store/records/` + parent-directory fsync**. **Presence = committed** — the maildir linearization, byte-for-byte the pattern email stores have shipped for decades.
- **Everything else is a derived projection**: INDEX.md, rendered markdown, mailbox entries. Written in the same loop iteration via a **redo journal** (intent staged + fsynced *before* the pivot, replayed idempotently after any crash). Projections can never disagree with truth for longer than one recovery pass, and on conflict **canonical wins, unconditionally**.
- **INDEX.md posture (operator rider):** a pure append-only log in normal operation AND at recovery — rebuild appends missing rows in canonical commit order; a wrong row is superseded by an appended correction row; history is never rewritten. INDEX's *authority* becomes derived (canonical = records). **CLOSED via CQ-8** (m-1 `c4-cq-m1/013500` + m-1.impl `020418` approve): m-1 confirms the reading of its locked §6 — **layout UNCHANGED; only crash-recovery provenance changes.** No silent change occurred; the COORD ran and closed.
- **fsync discipline (the classic-bug guard, fixtured not assumed):** record fsync BEFORE rename; directory fsync AFTER rename; intake/redo journal fsync BEFORE the pivot depends on them. Donor evidence: jcode `write_bytes_inner` (tmp→fsync→`.bak` hardlink→rename→dir-fsync) as the per-file primitive; codex's flush-without-fsync recorded as the anti-pattern; SQLite atomic-commit + `rename(2)`/`fsync(2)` semantics as the external citations.
- **Rejected (grilled):** Package B (super-journal commit-marker with INDEX co-canonical) — same durability, more bespoke ritual, five peer truths; and codex's eventual-consistency read-repair two-tier stance — explicitly the anti-model.

## 5. Recovery state machine (GRILL Q4 — locked)

Startup after any exit, in strict order; **no authority consumption until phase 4**:
- **Phase 0 — validate genesis:** genesis record present + config digest matches (§7). Mismatch/missing ⇒ serve read-only diagnostics, accept nothing, summon operator. Fail-closed at the root.
- **Phase 1 — scan canonical records:** per-record checksum. Torn file in `staging/` ⇒ delete (never committed — routine). Checksum-fail on a *committed* record ⇒ **quarantine** (`store/quarantine/` + a HELD-class incident record); never silent-skip (Kafka interior-corruption lesson), never fail-stop the store (§2A "never brick").
- **Phase 2 — rebuild projections:** replay redo journal; ensure every committed record's INDEX row / rendered `.md` / mailbox entries exist; canonical wins; idempotent (ARIES property: recovery of the recovery is safe); INDEX repairs are appended, never rewritten (§4).
- **Phase 3 — restore runtime tables:** burned-nonce set (derived from committed operator-verdict records, whose presence implies the decision-scoped burn — a crash can never resurrect a consumed token, and there is no separate burn record to diverge from its verdict; m-7-F1), parked-lane states (m-6 7-state records; the store IS the checkpointer, locked m-6 §4), seat-binding table (§11).
- **Phase 3.5 — refill the FIFO:** re-enqueue `intake − outcomes` in arrival order (§2.2). Entries with outcomes can never re-emit.
- **Phase 4 — open:** accept `submit`/`verify`. During 1–3: reads of committed records may be served; zero authority consumption (a half-restored burn set is how stale approvals replay).

## 6. Internal-fault disposition + terminal-state enum (GRILL Q4+Q5 — locked)

**One rule:** a trusted-side check that **throws, times out, or reads corrupt data** yields a typed outcome — never silent-accept, never a brick:
- **authority-bearing record → `HELD`**: not accepted (a crash is not a yes — K8s `failurePolicy: Fail`), not rejected (the check *failed to run*, it did not *say no* — the distinction is preserved in the type), parked operator-visible (m-6 A-surface), resolved by re-run or operator verdict. Represented append-only as **ONE compound canonical record**: the held-disposition record **embeds the candidate** (full candidate bytes inside the disposition record) and references the `intake_id` — one record, one §4 rename pivot, no second canonical write (m-7-F1, VP-ratified one-pivot-per-mutation; supersedes the two-record "candidate + disposition referencing it" shape, whose two renames opened a crash window between them).
- **non-authority record → bounce** to author with the fault-edge named.
- **the engine keeps serving**: fault handling touches only the sick record; the fault path is minimal and self-excluding (its own machinery passes through no gate that could HELD it — the K8s wedge lesson).

**Terminal-state enum — CLOSED via CQ-4 (m-2/m-3/m-6 triad, certified `master/relays/c4-cq-coord/RECONCILE-orchestrator-planner-20260702-031533.md`):** every command reaches exactly one of the **byte-exact closed set `{accepted, rejected, held}`** — no fourth outcome, no outcome-free exit, no persisted `submitted` limbo. `rejected` is the **shared** terminal token across the form/lineage/observe rejection classes (the legacy `bounced` value token is retired; "bounce" survives only as the verb for the author-return action); `held` is the **distinct** fault/fail-closed token; m-6 bucket-D keys off the aligned tokens; the m-2 registry field-home for `held` landed with the CQ-2/CQ-4 folds (m-2-authored — m-7 invents no schema).
**Exactly-one-outcome check (CTO integration item 3, confirmed):** m-3's refinement framing — "nothing appended → candidate-not-delivered + terminal audit record" — is the *effect description of the `rejected` outcome*, not a new state: the candidate never becomes a deliverable record, and the terminal audit record that IS appended is the `rejected` evidenced record itself, referencing the `intake_id` (§2.2). No fourth outcome and no outcome-free exit is introduced; NF-S16's exhaustiveness fixture stands unchanged.

## 7. Trusted config load + integrity (§2A.2; CQ-4b consumer)

- The policy-config artifact — **full author set (c6 x3-F2 fix): m-2, m-3, m-4, m-5, m-6** — per-domain sections: gate_category maps + the CQ-3 A-floor table (m-6), egress rules (m-3), capability priors (m-4), **the archetype registry (m-5 — the previously omitted author of a section this doc already named)**, and **m-2's declared schema section** (the registry slice config carries). m-5's CQ-4b section confirm: **OBTAINED** — `master/relays/c6-fix-m-5/COORD-planner-20260702-205849.md` (the registry = one m-5-authored, m-5-stamped section under the single top-level digest; contents = tag-space + ceiling composition + templates + axis maps, matching m-5's locked §13 "registry as config-sourced data"; loaded at genesis; no c3 reopen). **Composition CLOSED via CQ-4b (CTO ruling, all-domain confirms):** **per-domain sections, conductor-composed into one artifact with a single top-level digest**, plus **m-4's per-section version-stamp inside the artifact** (CTO integration item 1 — compatible with the single-digest ruling; a section stamp aids per-domain change attribution without adding a second integrity root). The artifact is **conductor-owned**, read **once at trusted startup**, top-level-digest-verified against the genesis chain: genesis carries the initial config digest; a legitimate config change is itself a committed store record (operator-authorized, carrying the new digest) so config history is auditable append-only.
- **No hot reload in v3.0** (restart-only; grilled under Q2's simplicity rule) — no agent turn can cause a re-read or mutate effective config (codex donor: config-by-value, no reload machinery on the agent path). Digest mismatch at startup ⇒ Phase-0 fail-closed.
- **Absent from every seat surface**: not a tool, not a resource, no config *path or value* in any seat-deliverable schema/description/prompt/tool-result/projection (fixture F7/G(iii); implementer's strengthening folded). Donors: codex untrusted-layer denylist + config-lock replay validation.

## 8. Interface guardrail + attach/pipe lifecycle (GRILL Q3 — locked)

### 8.1 The mechanism
The conductor is an **MCP server**; each seat connects over a dedicated per-seat channel (own stdio pipe or own bearer token) — the same channel m-1's FROM stamp rides. The per-seat tool registry is **constructed trusted-side per (seat, phase)**: exactly **`submit`, `project`, `read`**. `mint_seat` is conductor-internal. Donor: codex `ToolExposure` (visible-spec vs executable-registry independently controlled; unknown call = hard dispatch reject); external evidence that removing the tool — not prompting against it — is the layer that stops confused-agent actions (arXiv 2605.18414: 48–68% forbidden-tool selection under adversarial framing, 37% even with in-prompt allowlists).

### 8.2 Fill-time authority made physical
`submit`'s **input schema IS the rendered m-2 form** for this seat/phase: forbidden options are absent from the schema itself; hybrid picks (`parent_picker`/`recipient_picker`/`monotonic`) render as constrained enums over conductor-supplied candidate sets. The render is advisory; §3 step 2 re-validates authoritatively in the critical section.

### 8.3 Delivery/wake
Delivery = one `write()` of the relay (or pointer) onto the recipient seat's open pipe; a parked/idle seat is kernel-suspended on that pipe and wakes on the write; messages queue FIFO in the kernel while busy (`master/RUNTIME-RESEARCH.md` §2, primary-source-verified). `project()` exists for re-reads; the pipe is the knock. **Per-runtime wake adapters:** Claude Code = SDK streaming-input path ONLY (Remote Control polls — do not conflate, §13.3); codex = app-server JSON-RPC (verified §14.1-F); Gemini/others = ACP; HTTP-server runtimes (OpenCode/OpenHands) = delivered request on a held connection. **Operator rider (grilled): wake is a runtime-layer property — models never touch stdin — so every concrete lane, explicitly including non-frontier hosted (Kimi/DeepSeek/GLM/MiniMax/Qwen) and local (Ollama/vLLM/llama.cpp) lanes, earns drive-loop confirmation at bring-up via the lane-qualification probe** (≥8-step close/reopen canary; `reasoning_content` replay trap + non-streamed-tool-turns rule; `master/RUNTIME-RESEARCH.md` §12 #2/#4/#5, with #4 widened to include a local-serving lane).

### 8.4 The absence set + honest residual
No tool, resource, description, prompt text, or tool result delivered to any seat contains: the store path, config path, outbox path, operator-channel path, or any config value (fixtures G(i–iii)/F6; the implementer's outbox/operator-channel/config-value additions folded). Residual, stated plainly: the seat's host runtime still hands it Bash; a shell-bearing confused or malicious seat can act outside MCP entirely. The guardrail removes the *tool-surface* path — **confusion-resistant**, D4 claim verbatim, D5 accepted-risks recorded.

### 8.5 Attach lifecycle
`mint_seat` at attach binds seat ↔ {channel, credential} in the conductor's binding table (persisted — §11); detach/crash of a seat leaves its mailbox intact (durable store) and its binding recoverable; re-attach re-presents the same credential → same seat. Credential lifecycle detail (generation/rotation/revocation) stays the m-1 PLAN carry (m-1 §13.3) — m-7 executes, m-1 owns semantics.

## 9. Local-outbox-only external send

No seat has any external-send **tool** (locked m-5: no v3.0 `external_send` axis). The **only conductor-governed egress path** is the conductor's local outbox, drained by the conductor itself, with every item passing the m-3 §7 egress/content gate (fail-closed, dormant until the m-6 away bridge activates it). Outbox enqueue is a loop mutation (§2.1); the outbox is a store-visible queue (auditable). m-7 hosts the chokepoint; rule-set membership is m-3/operator policy. **D5 residual, stated here as everywhere the outbox is called "only":** this is a governance-surface claim, not a system sole-egress claim — under same-uid attach a shell-bearing confused or malicious seat can reach the network outside the conductor entirely (host Bash/curl); that path is the documented GRILL-LOCK D5 accepted-risk, out of scope. The outbox is the only egress *the governance system offers*, and the only one the egress gate can vouch for. *(rev2: supersedes r1's unqualified "only egress" — implementer blocker 2.)*

## 10. Store genesis + GC

- **Genesis:** store init writes the genesis record — schema_version, config digest (§7), address-space seed, creation stamp — as the first canonical record; every recovery validates it (Phase 0).
- **GC/retention:** v3.0 posture = **retain everything** (the design-of-record trail is the product); GC exists only as an operator-config-gated compaction of *derived* artifacts (old rendered projections, drained intake/redo journal segments **whose entries all have outcomes**). Canonical records are never GC'd in v3.0. Journal segment rotation is size-based, operator-config.

## 11. Restart seat-binding + park/wake execution

- **Persisted seat-binding table** (CQ-6, m-1 COORD confirms semantics): seat ↔ channel/credential bindings survive conductor restart as store-adjacent conductor-owned state; re-attach without re-mint; a restart does not strand seats (`master/DESIGN-REVIEW-2026-07-01.md` §2C).
- **Park/wake:** the m-6 7-state machine rows are store records; parking consumes nothing; waking = a verdict record committed through the loop + one pipe write; **re-observe-on-wake** (locked m-6) is sequenced by the loop before the woken lane's next accepted action. Resummon/digest timers ([E]-side timer wheel, deliberately outside the critical section) fire commands INTO the FIFO — the timer never mutates anything itself.
- **Away-token restart semantics:** burned nonces are committed records (Phase-3 restores the set); sibling-burn scope per decision = CQ-6's m-1 confirm.

## 12. THE SEAM MATRIX (lock grain — biting negative fixtures authored)

Promoted from `master/domains/m-7-conductor-core/audit/2026-07-01-audit-pair-reconcile.md` §2. Columns: owner · doc/section · m-7 obligation · **biting negative fixture (authored, executable-claim form)** · CQ.

| # | contract (owner) | doc/section | m-7 obligation (engine §) | biting negative fixture | CQ |
|---|---|---|---|---|---|
| S1 | store append + stamp (m-1) | m-1 design §5 (API surface), §6, §4; CQ-6 BASE triad | §3 pipeline; §4 pivot; §11 **persisted seat-binding table + re-attach credential proof (CQ-6 base, closed)** | **NF-S1 (bound to CQ-6 BASE):** submit over an unbound channel ⇒ reject, nothing staged; submit with payload-supplied `FROM`/`ROLE` ⇒ stamped values win, payload ignored byte-for-byte; kill -9 between pivot and projections ⇒ recovery yields exactly one INDEX row + full mailboxes, zero duplicates; conductor restart ⇒ re-attach with the same credential resolves to the same seat, no re-mint | CQ-6 **CLOSED (base)** |
| S2 | away-token verify/nonce-burn, Seam C (m-1) | m-6 design §4 :85-88; m-1 §C3.4; CQ-6 BASE triad | §3 verify; §2.4; **decision-scoped `(decision_id, seat)` sibling-burn = the presence of ONE compound operator-verdict record, atomic inside the loop (CQ-6 base, closed; one-pivot-per-mutation per m-7-F1)** | **NF-S2 (=F1, bound to CQ-6 BASE):** two `verify(token)` for the same nonce enqueued concurrently ⇒ exactly one operator-verdict record after both complete (its presence = the burn); the loser fail-closes typed; resolving one choice burns the decision-scoped sibling set (implied by the same single record); replay after restart ⇒ still burned (Phase-3 derives the burn set from verdict records). *(`re-mint-supersedes` is NOT part of this fixture — §2C build-carry, unreviewed add-on, per the VP-approved CQ-6 re-scope.)* | CQ-6 **CLOSED (base)** |
| S3 | form render / fill-time authority (m-2) | m-2 design §3, §4 | §8.2 | **NF-S3:** a pair-seat's rendered `submit` schema contains no `direct-override`/merge-grant option (schema introspection); a hand-crafted MCP call supplying the forbidden enum value anyway ⇒ authoritative validation rejects at §3 step 2 (render is advisory, loop is authoritative) | — |
| S4 | form-validation + lineage gate (m-2) | m-2 §4 steps 4–5; `master/ARCHITECTURE.md §3 "Write path — reconciled seam"` | §3 step 2 | **NF-S4 (=F5):** validator forced to throw on an authority-bearing candidate ⇒ outcome `held`, author-visible, loop alive, next FIFO entry processes; no store path exists where the candidate is deliverable | — |
| S5 | phase-split required-set (m-1+m-2+m-7) | `master/DESIGN-REVIEW-2026-07-01.md` §2A.5 | §3 step 2a executes the **CQ-1(a) STEP-GATE** (closed: observe-owned `required_when` predicates are step-gated on observe-layer presence) | **NF-S5 (bound to CQ-1(a)):** a Step-1 IMPL action-report submit with no observe layer present ⇒ accepted without the observe-owned fields (their `required_when` is step-gated off); the SAME submit with the Step-2 observe layer present ⇒ those fields required + conductor-filled | CQ-1 **CLOSED** |
| S6 | observe-as-send hook (m-3) | m-3 design §2, §3 | **§3 step 4** (predicate selected by the step-3 classification — closed CQ-5 ordering); §2.3 timeout | **NF-S6:** a (Step-2) observe hook attempting to write outside its closed allowlist (e.g. `FROM`) ⇒ internal fault, allowlist violation logged, §6/CQ-2 class-conditional disposition (authority-bearing ⇒ `held`; non-authority ⇒ `rejected`/author-return); an authority-class hook timeout ⇒ `held`, never accept (a non-authority hook timeout ⇒ `rejected`/author-return per §6 — distinct from a non-authority *unobservable/no-vantage* record, which is `accepted`+`self_reported` per locked m-3) | CQ-5 **CLOSED** |
| S7 | decision-② fail-closed (m-3 + op-②) | m-3 fold (CQ-2 triad) + `master/DESIGN-REVIEW-2026-07-01.md` §2A.7 | §3/§6 execute the **closed CQ-2 disposition: class-conditional fail-closed for authority-class `record_integrity ∈ {self_reported, mixed}`, disposition token = `held`** (m-3 fold + m-2 field-home landed; c6-widened to include `mixed`) | **NF-S7 (bound to CQ-2):** an authority-class record whose `record_integrity ∈ {self_reported, mixed}` ⇒ outcome `held` (never delivered-as-accepted, never silent); a non-authority `self_reported`/`mixed` record still delivers labeled (honest-fallback preserved) | CQ-2 **CLOSED** |
| S8 | pure-judgment A-floor (m-2/m-6) | CQ-3 triad (m-6 table + m-2 mechanics); §2A.6 | enforce the **closed CQ-3 (phase × record_kind) A-floor table** at §3 step 2a via the existing monotonic MAX (no new m-2 mechanism); CEREMONY_TIER monotonic | **NF-S8 (bound to CQ-3):** a submit picking below the table's floor for its (phase × record_kind) ⇒ floor wins (monotonic MAX), `gate_category=ceremony_downgrade` auto-set; a table-absent (phase × record_kind) pair ⇒ the `other`→A fail-safe | CQ-3 **CLOSED** |
| S9 | egress gate (m-3) | m-3 §7 | §9 chokepoint | **NF-S9:** an outbox item flagged by the scan ⇒ blocked + `egress_blocked` park + local resummon; **no conductor-owned code path** sends unscanned bytes externally (the outbox drain is **the conductor's only component that writes to non-seat destinations** — seat pipes/MCP channels are seat-facing and out of this claim's grain (c6 m-7-F6 reword); a governance-surface claim; the D5 same-uid seat residual per §9 applies) | — |
| S10 | routing record + route_dispatch (m-4) | m-4 §3, §5 | §3 acceptance; host fail-closed refusal | **NF-S10:** no acceptable route ⇒ `routing_unavailable`/`human_decision_required`, never a silent default; engine gate evaluation over a routing record never reads `chosen_model` (R2 grammar — m-2 AC14 hosted) | — |
| S11 | template/pane spawn (m-4/m-5) | m-4 GL-4 §7; m-5 §7 | spawn panes (existing multiplexer), deliver boot relays, record ceiling per assignment | **NF-S11:** an off-floor template slot spawn emits a `routing_decision` (**authored `FROM=operator` on the operator-relay channel — m-4 §7/GL-4, m-4-F4; the routing_assignments seat-scope admits `operator` only on `template_ref`-bearing records**) with `declared_deviated=true` + snapshotted justification (no-bypass); spawn above any ceiling axis ⇒ fail-closed refusal | — |
| S12 | `slot_in` classification (m-3/m-5) | CQ-5 triad (m-5 + m-3 co-sign); m-3 §5.1; m-5 §4 | §3 step 3 — the **closed CQ-5 ordering: classify at acceptance, post-gate / pre-observe / atomic-bind** (m-3/m-5-confirmed; required by locked m-5 §4) | **NF-S12 (bound to CQ-5):** a lane-supplied `slot_in` value in a submit payload ⇒ ignored; the conductor-classified value is on the committed record, bound in the same atomic commit as the observation it selected, and immutable; a re-tag attempt post-acceptance has no verb to express it | CQ-5 **CLOSED** |
| S13 | scheduler park/wake (m-6) | m-6 §4 | §11 | **NF-S13:** kill -9 with three lanes parked ⇒ Phase-3 restores all three `parked_waiting_human` states from store records alone; wake before the verdict record commits is impossible (wake is pipe-write sequenced after commit) | — |
| S14 | buckets + ODB delivery (m-6) | m-6 §2, §3 | mailbox projection + ODB render slots | **NF-S14:** a bucket-D-class bounce never appears in the operator mailbox (author-facing only); an A-gate record is non-suppressible (appears regardless of digest/batch config) | CQ-4 (token) |
| S15 | trusted config (authors m-2/m-3/m-4/m-5/m-6 — c6 x3-F2 full set; composition CTO-ruled) | CQ-4b ruling + confirms (m-5 section confirm OBTAINED: `c6-fix-m-5/COORD-planner-20260702-205849`); GL D2b; §2A.2 | §7 — the **closed CQ-4b composition: per-domain sections → conductor-composed, single top-level digest, m-4 per-section version-stamp inside** | **NF-S15 (=F7, bound to CQ-4b):** top-level digest mismatch at startup ⇒ Phase-0 refuses authority ops, serves diagnostics; a per-section stamp change without a matching top-level digest change ⇒ mismatch (one integrity root); seat-side `read`/`project`/`submit` results grepped for the config path or any effective config value ⇒ zero hits | CQ-4b **CLOSED** |
| S16 | terminal-state tokens (m-2/m-3/m-6) | CQ-4 triad; m-2 §17.1 Q-E | §6 — the **closed CQ-4 byte-exact set `{accepted, rejected, held}`**; `rejected` shared across rejection classes; `held` distinct; bucket-D aligned; stale `bounced` retired | **NF-S16 (bound to CQ-4):** state-machine exhaustiveness — no code path exits without exactly one of the byte-exact `{accepted, rejected, held}`; a grep of the engine's emitted tokens finds no `bounced` value token | CQ-4 **CLOSED** |
| S17 | integrated architecture (CTO/VP) | `master/ARCHITECTURE.md` §1 (the integrated model) + §C4.1–C4.3 (the engine × six contracts + claim boundary) | execute in locked order; surface, never arbitrate | **NF-S17:** composite — NF-S1..S16 all green constitutes the integration fixture; any m-7 doc text contradicting a locked m-1..m-6 sentence = review blocker | — |
| S18 | deployment-fork claim boundary (operator) | GL D1–D5; `master/DESIGN-REVIEW-2026-07-01.md` §2 | §2.4 scope; §8.4 residual; §16 sweep | **NF-S18 (=F8+G):** G(i) registry enumeration = exactly {submit, project, read}, zero path-bearing tools/resources; G(ii) scripted confused-seat "edit the store file directly" turn ⇒ no tool can express it; only store effect via submit; G(iii)+F6 grep of every seat-deliverable surface for **raw conductor-internal** store/config/outbox/operator-channel paths + **effective config values** ⇒ zero hits (ordinary evidence citations to relay/design *files* are not hits — rev2 qualifier, prevents false positives); F8 claim-sweep (§16) clean | — |

## 13. Fixture set (executable-claim form; conductor-registry checks at build time)

F1=NF-S2 (double-verify) · F2 crash-free commit shape (submit ⇒ exactly one record + one INDEX row + one rendered `.md` + expected mailboxes, all checksums valid) · F3 staging-crash (kill -9 pre-pivot ⇒ no committed record, staging cleaned, intake entry re-enqueued once, no duplicate on re-run) · F4 projection-crash (kill -9 post-pivot ⇒ projections rebuilt, authority not re-consumed, INDEX append-only preserved) · F5=NF-S4 (validator-throw ⇒ held) · F6+G(i–iii)=NF-S18 (guardrail negatives) · F7=NF-S15 (config mismatch) · F8=§16 (claim-sweep) · **F9 (new, grilled):** durable-FIFO — kill -9 with 5 enqueued commands (2 outcomes reached) ⇒ recovery re-enqueues exactly the 3 without outcomes, in arrival order; zero re-emission of the 2 · **F10 (new):** fsync-ordering — simulated power-cut between record-fsync and dir-fsync, and between rename and projection writes ⇒ recovery converges to the correct store both times · **F11 (new, c6 m-7-F1):** one-pivot-per-mutation — enumerate every mutation class {submit-accept, submit-reject, held (compound, candidate embedded), operator-verdict (compound, burn implied), park/wake transition, outbox enqueue, genesis, config-change, GC marker} and assert each commits **exactly one canonical record via exactly one rename pivot**; a crash injected at every syscall boundary of every mutation class leaves the store in either fully-committed or not-committed state — **no mutation class has a second canonical rename to crash between**.

## 14. GRILL_LOCK

```text
GRILL_LOCK_ID: c4-grill-m-7
GRILL_REQUIRED: yes
GRILL_SOURCE:
- plan/design/audit relay read: c4-design dispatch (…-162721); the reconciled c4 audit pair (substrate-audit + audit-pair-reconcile); DESIGN-REVIEW-2026-07-01 §2A; GRILL-LOCK-deployment-fork (D1–D5); RUNTIME-RESEARCH §2/§7/§8/§12/§14
- code/docs inspected: locked m-1..m-6 design docs (contract inputs); codex rollout/ToolExposure/config-lock; jcode storage/wire; claude-code mailbox; live master/ store probes; external prior-art set (SQLite atomic-commit, maildir, LMAX, Kafka, K8s failurePolicy, arXiv 2605.18414)
- questions answered from codebase: donor patterns + absence claims (the c4 audit); wake mechanism verification (RUNTIME-RESEARCH §2/§14 primary-source pass)
- questions asked operator: the 6 below (dispatch grill agenda)

Resolved decisions:
- Q1 commit pivot — Package A: canonical-record rename() = the single linearization point; INDEX/rendered/.md/mailboxes = derived projections + redo journal; fsync-before-rename + dir-fsync-after fixtured (F10) — one truth, one proven op; recovery = dumb replay — source operator
- Q2 process model — one conductor process; N channel handlers as mail-carriers; ONE commit thread; plain FIFO (no priority classes v3.0); reads/renders concurrent on immutable committed state; Step-2 probe timeout → HELD — source operator ("keep it simple, FIFO")
- Q2-rider durable FIFO — intake journal (append+fsync) ahead of the in-memory queue; ack only on typed outcome; recovery re-enqueues intake−outcomes in arrival order — source operator ("queue does NOT get cleared on crashes")
- Q2-rider clear-on-pop — outcome records reference intake_id; consumption is atomic with the outcome commit; re-enqueue set = intake−outcomes; zero stale re-emission — source operator (prior-harness stale-re-emission incident)
- Q3 guardrail — conductor-as-MCP-server; per-seat trusted-side registry of exactly {submit, project, read}; submit schema = the rendered m-2 form; delivery/wake = pipe write (runtime-layer property); absence set incl. outbox/operator-channel paths + config values — source operator
- Q3-rider wake breadth — wake confirmed at the runtime layer (CC SDK-streaming [P], codex app-server [P §14.1-F], ACP, HTTP-server class); EVERY lane incl. non-frontier hosted (Kimi/DeepSeek/GLM/MiniMax/Qwen) + local (Ollama/vLLM/llama.cpp) earns drive-loop confirmation via the lane-qualification probe at bring-up; §12 #4 spike widened to include a local-serving lane — source operator + RUNTIME-RESEARCH primary sources
- Q4 recovery — phases 0–4 (genesis → scan/quarantine → rebuild-projections → runtime tables → open); refuse authority consumption until complete; INDEX recovery is append-only (corrections appended, history never rewritten) — source operator ("INDEX is an append only stack")
- Q4 fault taxonomy — trusted check throws/times-out/corrupt ⇒ authority → HELD (typed, operator-visible; failed-to-run ≠ said-no), non-authority → bounce; quarantine isolates; the store never bricks; fault path minimal + self-excluding — source operator ("elaborate" → confirmed)
- Q5 terminal-state enum — STRUCTURE locked now ({accepted, rejected, held}; exactly-one-outcome; no limbo); SPELLING + m-2 registry home + bucket-D naming carried to CQ-4 — source operator

Rejected alternatives:
- Package B super-journal commit-marker (five peer truths; bespoke ritual; same durability) — operator, after the durability question was answered (B saves nothing A loses; both fsync-before-pivot)
- fast-path mutations skipping the loop — two write paths = the two-truths disease; nothing gained at our volume
- priority-class dequeue (claude-code donor) — deferred future knob; v3.0 = plain FIFO — operator
- fail-stop on interior corruption — bricks the team on one bad sector; quarantine+HELD instead
- config hot-reload — restart-only; no agent-reachable reload
- bespoke (non-MCP) seat protocol — reinvents MCP, loses built-in client support

Still operator-owned:
- CQ-4 token vocabulary ratification rides the m-2/m-3/m-6 COORD (operator sees it at the c4 lock)
- GC/retention values + egress rule membership + config change authorization = operator config policy (§J pattern)
- the runtime spike gates (RUNTIME-RESEARCH §12) remain operator-opened; this design cites, never opens them

Design-lock impact:
- §2 (loop + durable FIFO + clear-on-pop), §4 (Package-A pivot + INDEX posture), §5 (recovery), §6 (fault + enum structure), §7 (config), §8 (guardrail + wake), §12 (seam matrix), §13 (fixtures) carry these decisions; the c4 DESIGN_LOCK_ID must reference GRILL_LOCK_ID c4-grill-m-7; CQ closures fold in before lock or are carried explicitly non-locking
```

## 15. CQ LEDGER — **GATE SATISFIED** (all design-LOCK rows CLOSED; certified `master/relays/c4-cq-coord/RECONCILE-orchestrator-planner-20260702-031533.md`, CQ-6 re-scoped by `...-032227.md` + VP approve `RECONCILE-orchestrator-reviewer-20260702-032843.md`)

| CQ | owner | resolution (closed) | closing legs | lands in |
|---|---|---|---|---|
| CQ-1 phase-split required-set | m-1+m-2 | **(a) STEP-GATE** — observe-owned `required_when` step-gated on observe-layer presence | m-1 `013500` + m-1.impl `020418` + m-2 co-sign `024704` + m-2.impl `030145` | §3 step 2a; NF-S5 |
| CQ-2 decision-② fail-closed | m-3 (step-c) + m-2 | class-conditional **fail-closed** for authority-class `record_integrity ∈ {self_reported, mixed}`; disposition token = **`held`**; m-2 field-home landed (c6: `mixed` leg added) | m-3 fold + m-2 `014626` + m-3.impl `021724` + m-2.impl `021604` | §3/§6; NF-S7 |
| CQ-3 A-floor table | m-2+m-6 | (phase × record_kind) table, m-6-authored; rides existing monotonic MAX (no new m-2 mechanism) | m-6 `015800` + m-2 mechanics + m-2.impl + m-6.impl `024620` | §3 step 2a; NF-S8 |
| CQ-4 terminal tokens | m-2+m-3+m-6 | byte-exact **`{accepted, rejected, held}`**; `rejected` shared / `held` distinct; bucket-D aligned; `bounced` retired | m-2/m-3/m-6 + all impl incl. `024620` | §6; NF-S16 |
| CQ-4b config composition | CTO | per-domain sections → conductor-composed, **single top-level digest** + m-4 per-section version-stamp inside (`022000`) | CTO ruling + m-2/m-3/m-4/m-6 confirms + all impl | §7; NF-S15 |
| CQ-5 slot_in ordering | m-3+m-5 | classify at acceptance, **post-gate / pre-observe / atomic-bind** (required by locked m-5 §4) | m-5 `014506` + m-5.impl `020448` + m-3 co-sign `024732` + m-3.impl `030205` | §3 step 3; NF-S12 |
| CQ-6 seat-binding + sibling-burn | m-1 (+m-6) | **CLOSED ON THE BASE:** persisted binding table + re-attach credential proof + decision-scoped `(decision_id, seat)` sibling-burn + atomic burn in the loop. **`re-mint-supersedes` = NOT part of the closure** — §2C away-bridge **build-carry**, m-1-confirmed-fit (`021500`) but adversarial review deferred to its build step; never presented as pair-approved here | base: m-1 `013500` + m-1.impl `020418` (base only) + m-6 co-sign `020100` + m-6.impl `020447` | §11; NF-S1/NF-S2 (base only) |
| CQ-7 observe row-parity | m-2 | non-locking — tracked pre-Step-1-PLAN SHOULD | (unchanged) | flag only |
| CQ-8 INDEX derived-authority | m-1 | **layout UNCHANGED**; only crash-recovery provenance changes | m-1 `013500` + m-1.impl `020418` | §4 |

The design-LOCK gate is satisfied: every `blocks: design-LOCK` row is closed (CQ-6 on the base, with the add-on explicitly carried non-locking per the VP-approved re-scope).

## 16. F8 claim-sweep (run over this doc — re-run at rev2)

Swept: the only "by construction" phrase in this design is §2.4's serialized-loop double-accept kill (control-flow, explicitly scoped, with its D5 residual stated in the same breath). The guardrail is described as confusion-resistant at every mention (§0, §8.1, §8.4, §12 S18); the D4 claim-set language is used verbatim; D5 accepted-risks (config / store-write / operator-FROM under same-uid attach) are restated in §0/§8.4/§9. No "unbypassable", no "sole-writer by construction", no "same-uid write-exclusion" anywhere.
**rev2 strengthening (from implementer blocker 2):** the sweep now targets **semantics, not just vocabulary** — any exclusivity claim ("only egress", "only writer", "no code path", "sole X") over a surface a same-uid seat can reach outside MCP is a violation unless scoped to the conductor-governed surface WITH the D5 residual stated beside it. r1's §9/NF-S9 unqualified "only egress / only socket writer" was exactly such a miss (caught in review, fixed at rev2); fixture F8 is specified to grep for the exclusivity-claim class, not merely the three banned phrases.
**c6 strengthening (m-7-F9):** the sweep's specified classes now ALSO include the **writability/reachability token family** — `non-lane-writable`, `lane-proof`, `seat-proof`, "no lane can …", "a lane cannot …" — which makes the same implicit strength claim through a different vocabulary. Any such token over state a same-uid seat can reach outside MCP is a violation unless scoped to the tool/governance surface (no verb exists) with the D5 residual stated or referenced beside it. §3 step 3's previously bare `non-lane-writable` was the motivating instance (now scoped in place).

## 17. Non-re-cut path (Step-1 → standalone)

Step-1 rides existing runtimes: seats are host-runtime processes on conductor-minted MCP channels; pane-spawn via existing tmux/zellij/OS-terminal (S11); wake via per-runtime adapters (§8.3). The standalone runtime (Step-3+) swaps the channel medium (MCP stdio/token → owned transport), the spawn mechanism, and the attestation backend — **the loop, pivot, store shape, recovery machine, fault taxonomy, config model, and the 3-verb surface are invariant**. Same non-re-cut discipline as m-1 §9: only backends swap.

## 18. Novelty / promote map (honest)

Promoted, not invented: single-writer loop (LMAX/Redis/codex), rename-pivot (maildir), redo-journal recovery (ARIES/SQLite), per-record checksum + quarantine (Kafka, minus its silent-skip), fail-closed fault gate (K8s), trusted-side tool registry (codex ToolExposure), durable intake journal (WAL commonplace), jcode per-file write discipline. **The assembly is the contribution:** one trusted engine where the commit loop is simultaneously the trust boundary (channel-stamped identity), the fail-closed governance gate (form/lineage/observe in the critical section), and the narrow-API enforcement point, over a human-readable append-only governance store — the audit's external sweep found no prior system binding these.

## 19. Acceptance criteria draft (for the eventual operator-opened PLAN — NOT this phase)

1. F1–F11 + NF-S1..S18 pass as conductor-registry checks (E2). (F11 = the c6 one-pivot-per-mutation fixture; §22's "F1–F10/G" line records the c4 lock-time set as history — the PLAN gate is THIS line, which includes every later-added fixture.)
2. kill -9 at any single point in the §3 pipeline yields, after recovery, a store byte-equivalent to either "command fully applied" or "command still queued" — never anything else (crash-point sweep harness).
3. A seat's full MCP surface enumerates to exactly three tools; the G-fixtures pass against every shipped seat archetype form.
4. The claim-sweep (F8) passes over all shipped docs.
5. Live-team dogfood: this standing team's store migrates onto the conductor (genesis + import of `master/relays/` as read-only history per m-2 §8 legacy strictness) — the operator stops being the relay transport (ROADMAP Step-1 exit).

## 20. Carries + review edges

- **To m-7.implementer (DESIGN-REVIEW):** hunt over-reach (any sentence where m-7 authors policy), claim-boundary regressions (§16), under-specified engine seams (anything an implementer couldn't build to), and the fixture set's bite.
- **CQ COORDs: all closed** (§15). **Build-carries this design adds to the §2C/C3.7 ledger:** `re-mint-supersedes` (m-1/m-6 burn-set widening on resummon-mint — m-1-confirmed-fit, adversarial review owed at the away-bridge build step; dormant in Step-1: no away-bridge/resummon exists yet, and the base decision-scoped sibling-burn + m-6's never-auto-resolve FSM hold meanwhile); CQ-7 row-parity (m-2 pre-Step-1-PLAN SHOULD); the runtime spikes cited in §8.3 (operator-gated, RUNTIME-RESEARCH §12).
- **Design-review lineage:** this doc = DESIGN_DOC_ID `c4-design-m-7-conductor-core`; **DESIGN_LOCK_ID `c4-design-m-7-lock` is assembled at §22** (references GRILL_LOCK_ID `c4-grill-m-7`; effective on implementer lock-package approve + VP co-sign); the lock package cites the audit pair-reconcile supersession note (VP carry-forward #1, already patched into the merge artifact).

## 21. rev2 fold-log (m-7.implementer DESIGN-REVIEW `c4-design-m-7/DESIGN-REVIEW-implementer-20260702-002302.md`, verdict must-revise — all three findings verified correct by the planner before folding)

1. **Blocker 1 — `slot_in`/observe ordering contradiction (§3).** r1 ran observe at step 2c and classified `slot_in` at step 3, contradicting m-3's contract that the done-predicate is keyed on `slot_in`; the parenthetical also silently picked an ordering while labeling CQ-5 open. Folded: classification moved to step 3 (post-gates, **pre-observe**); observe becomes step 4 reading the just-classified in-courier value; both bind into one atomic commit (step 5); the ordering is explicitly framed as **m-7's proposal to the CQ-5 COORD** — observe-hook placement is not lockable until CQ-5 closes. §12 S12 + §15 CQ-5 rows aligned.
2. **Blocker 2 — sole-egress overclaim (§9, NF-S9).** r1's "the only egress / the only socket writer / no code path" asserted system-level sole-egress, which is false under same-uid attach and is the claim class D4/D5 forbids. Folded: all exclusivity wording scoped to the **conductor-governed** surface with the D5 residual restated in place; §16's sweep upgraded from vocabulary-grep to semantics (exclusivity-claim class), and the miss is recorded there as the reason.
3. **Lock-hygiene — CQ-8 missing from the header lock-gate line.** Folded: header status line now names CQ-1..CQ-6 + CQ-4b + **CQ-8** and marks CQ-8 as newly surfaced by this design and part of the no-lock gate.
4. **Non-blocking note folded:** NF-S18/G(iii) qualified to raw conductor-internal paths + effective config values (evidence citations to relay/design files are not fixture hits).

No other section changed; the GRILL_LOCK (§14) decisions are untouched. Resent for re-review on the same DESIGN_DOC_ID.

**rev3 fold (m-7.implementer re-review `c4-design-m-7/DESIGN-REVIEW-implementer-20260702-003942.md`, verdict must-revise — one residual blocker, verified correct):** the §1 overview diagram still carried an unqualified "local outbox = the only external egress" — exactly the semantic class rev2's §16 rule bans, missed because the rev2 fold touched §9/NF-S9 but not the diagram. Folded: diagram line now reads "conductor-governed external egress … (D5 same-uid residual — §9)". The full-document semantic sweep was then re-run (rg over only/sole/no-code-path/never-reach classes): every remaining hit is the licensed §2.4 claim with residual, surface-scoped guardrail language, internal state-machine exhaustiveness, or the locked mechanism name "local-outbox-only" (whose §9 body carries the residual in place). The rev2 fold-log's "all exclusivity wording scoped" claim is true as of THIS rev, not rev2 — noted for lock-package honesty. Rev2's other folds (slot_in ordering, CQ-8 gate, G(iii) qualifier) were implementer-confirmed no-objection.

**r4 fold (CQ-gate closure, per `master/relays/c4-cq-coord/RECONCILE-orchestrator-planner-20260702-031533.md` + the CQ-6 re-scope `...-032227.md` + VP approve `RECONCILE-orchestrator-reviewer-20260702-032843.md`):**
1. §15 ledger flipped to CLOSED with resolutions + closing legs; the gate line updated (gate satisfied).
2. NF-Sx bindings landed: NF-S5→CQ-1(a) step-gate; NF-S7→CQ-2 fail-closed `held` (authority-class `{self_reported, mixed}`, c6-widened); NF-S8→CQ-3 table + `other`→A fail-safe; NF-S16→CQ-4 byte-exact tokens; NF-S15→CQ-4b composition (single top-level digest + per-section stamp = one integrity root); NF-S12→CQ-5 confirmed ordering; NF-S1/NF-S2→**CQ-6 BASE only** (re-mint-supersedes explicitly excluded from the fixture and never presented as pair-approved); §4→CQ-8 (layout unchanged, m-1-confirmed).
3. CTO integration items applied: (i) m-4 per-section version-stamp inside the single-digest artifact → §7 + NF-S15; (ii) byte-exact `{accepted, rejected, held}` with the `bounced` value token retired — swept §6/§12/§14 ("bounce" survives only as the author-return verb); (iii) m-3's exactly-one-outcome framing checked and confirmed non-expanding → §6.
4. `re-mint-supersedes` recorded as a §2C build-carry (§20) with its dormancy-in-Step-1 residual stated.
5. Status line → DESIGN-LOCK-READY r4; DESIGN_LOCK block written at §22.

**r5 fold (m-7.implementer lock-package review `c4-design-m-7/DESIGN-REVIEW-implementer-20260702-034630.md`, verdict must-revise — one stale-fold blocker, verified correct):** §3 step 3 still carried the r2-era "m-7's PROPOSAL / not lockable until CQ-5 closes / may renumber" language inside lock-bearing pipeline text, contradicting the r4-closed CQ-5; and §12 S6 still pointed the observe hook at the pre-rev2 "§3 step 2c". Folded: §3 step 3 rewritten as closed CQ-5 text (post-gate / pre-observe / atomic-bind) citing the four closing legs, with the r1→r2→closure history compressed into a history note; S6 repointed to §3 step 4 with the CQ marked CLOSED; §20's "future DESIGN_LOCK_ID" made current (non-blocking note, folded). The "proposal" wording surviving in THIS §21 fold-log (rev2 entry above) is historical narration of what rev2 did, not lock-bearing design text. No mechanism changed; no other section touched.

**c6 fold (re-review cleanup; VP c6-decomp amendment `c6-decomp/RECONCILE-orchestrator-reviewer-20260702-192059.md` + c6-apply review `c6-apply/RECONCILE-orchestrator-reviewer-20260702-203146.md`):** doc-only consistency corrections to this locked doc, no mechanism change, §22 LOCKED CONTENT unchanged. (i) **NF-S6 internal-fault disposition split on two axes** (authority-class × whether the trusted check could run) — authority-bearing internal fault ⇒ `held` (never silent `rejected`), non-authority ⇒ `rejected`/author-return, distinct from the non-authority *unobservable/no-vantage* record which is `accepted`+`self_reported` (m-7-F3, per the VP-ratified c6-decomp two-axis wording). (ii) **NF-S7 + the §15 CQ-2 ledger row + the r4 fold-log NF-S7 summary widened** from authority-class `self_reported` to `record_integrity ∈ {self_reported, mixed}`, so m-7's CQ-2 fixture converges with the m-2/m-3 CQ-2 canonical (VP-caught token-convergence miss, c6-apply). (iii) **Stale byte-level cites repointed** to stable section-name anchors (the `ARCHITECTURE.md` write-path seam + `m-1 design §5`), which drifted when c5 edited those files (m-7-F4). The serialized-loop two-honest-seats double-accept kill and every §22 lock invariant are untouched.

## 22. DESIGN_LOCK (assembled; effective on m-7.implementer lock-package approve + VP co-sign)

```text
DESIGN_LOCK_ID: c4-design-m-7-lock
DESIGN_DOC_ID: c4-design-m-7-conductor-core (this doc, r5)
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: c4-grill-m-7 (§14 — folded; operator decisions + riders)
PAIR_APPROVAL: design approved at r3
  (master/relays/c4-design-m-7/DESIGN-REVIEW-implementer-20260702-004452.md);
  LOCK PACKAGE approved at r5 — m-7.implementer DESIGN_REVIEW_VERDICT: approve on
  DESIGN_LOCK_ID c4-design-m-7-lock
  (master/relays/c4-design-m-7/DESIGN-REVIEW-implementer-20260702-035245.md);
  r4/r5 = CQ-closure + consistency folds only, no mechanism change from r3
CQ_GATE: SATISFIED — all eight design-LOCK CQs closed (§15; certified
  master/relays/c4-cq-coord/RECONCILE-orchestrator-planner-20260702-031533.md;
  CQ-6 re-scoped to BASE by ...-032227.md; VP approve
  master/relays/c4-cq-coord/RECONCILE-orchestrator-reviewer-20260702-032843.md)
LOCKED CONTENT: the engine (§1–§11) — durable-FIFO single-thread commit loop w/
  atomic clear-on-pop; Package-A canonical-record rename pivot + derived projections
  (INDEX layout unchanged, CQ-8); phases 0–4 recovery; byte-exact {accepted, rejected,
  held} w/ HELD fault disposition (CQ-2/CQ-4); trusted config = per-domain sections,
  single top-level digest, per-section stamps, load-once (CQ-4b); MCP {submit, project,
  read} guardrail w/ schema-as-form + pipe wake; conductor-governed local outbox;
  genesis/GC; persisted seat-binding + decision-scoped sibling-burn (CQ-6 base);
  the seam matrix at lock grain (§12) + fixtures F1–F10/G (§13)
CLAIM BOUNDARY: confusion-resistant (GL D4 verbatim); sole licensed by-construction =
  §2.4 serialized-loop double-accept kill; D5 accepted-risks restated §0/§8.4/§9;
  semantic claim-sweep clean at r3+r4 (§16)
BUILD-CARRIES (non-locking, §20): re-mint-supersedes (§2C away-bridge build step);
  CQ-7 row-parity (m-2 pre-PLAN SHOULD); operator-gated runtime spikes (§8.3)
GRANTS NOTHING: no PLAN, no IMPL, no code/pcode, no spike; Step-1 PLAN remains a
  separate operator-opened gate (re-baseline step (e))
```

**c6-fix-m-7 fold (pair slice of the c6 re-review punch-list, dispatch `master/relays/c6-fix-m-7/DESIGN-orchestrator-planner-20260702-204518.md`; doc-only, no mechanism change, §22 lock invariants unchanged; for m-7.implementer review):**
1. **m-7-F1 (B, ◆ VP-ratified) — one pivot per mutation:** the HELD outcome is now ONE compound canonical record embedding the candidate (§6); `verify()` now commits ONE compound operator-verdict record whose presence IS the decision-scoped burn — no separate burn record (§3, §5 Phase-3, NF-S2); new fixture **F11** asserts every mutation class commits exactly one canonical record via one rename (§13). Burn *semantics* (decision-scoped, atomic, CQ-6 base) unchanged; only the record shape (1 compound vs 2 records with a crash window between renames).
2. **x3-F2 (B) — trusted-config author set completed:** m-2 + m-5 added (m-5's archetype-registry section was named but its author omitted); §7 + S15; m-5's CQ-4b section confirm subsequently OBTAINED (`c6-fix-m-5/COORD-planner-20260702-205849`).
3. **m-7-F5 (M) — intake single-writer discipline:** handlers hand commands over a bounded channel to ONE intake-writer task (codex donor applied at the intake tier too); "handlers never touch any file" is now literally true; §1 diagram + §2.1 + §2.2.
4. **m-7-F6 (M) — NF-S9 grain fix:** "only socket-writing component" → "only component that writes to NON-SEAT destinations" (seat pipes/MCP channels are seat-facing sockets; the old wording contradicted §8.3 delivery).
5. **m-7-F7 + x2-claim-honesty-F8 (M/C) — README brought current:** status = DESIGN-LOCKED / c4 CLOSED with §22 + VP co-sign pointers; the stale "under VP review" consumes-note replaced with a pointer to the locked seam matrix.
6. **m-7-F9 (C) — F8 sweep widened:** the writability/reachability token family (`non-lane-writable`, `lane-proof`, `seat-proof`, "no lane can") added to §16's specified sweep classes; §3 step 3's bare instance scoped in place (tool-surface grain + D5 reference).
CTO-applied c6 items (NF-S6 two-axis, CQ-2 `{self_reported, mixed}` widen, §21 c6 entry, anchor repoints) verified present and NOT re-touched.
**c6-fix r2 (implementer review `c6-fix-m-7/DESIGN-REVIEW-implementer-20260702-210010.md`, must-revise — both findings verified + folded):** (i) §19 AC1 widened to **F1–F11** so the c6 fixture is inside the eventual Step-1 PLAN gate; §22's "F1–F10/G" line is EXPLICITLY the c4 lock-time historical set (the certificate is not retro-edited) — the PLAN gate is §19 AC1, which includes every later-added fixture; (ii) README hosts+executes line widened to `record_integrity ∈ {self_reported, mixed}` matching the CQ-2 canonical (the c6-dispatched README lines 47-49 miss). Also folded post-review: the m-5 CQ-4b section confirm recorded as OBTAINED (§7/S15; `c6-fix-m-5/COORD-planner-20260702-205849`) + the `archetype_registry` section-key/canonical-ordering answer relayed (`c6-fix-m-7/COORD-planner-20260702-210344`); ARCHITECTURE C4.1's stale author-set line flagged to the CTO (their doc).
