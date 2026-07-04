# S2 Slice-2 — Design of the thickened engine (recovery phase machine · intake-writer · genesis/GC · the generalized owed-item projection)

**DESIGN_DOC_ID:** `s2-slice-2-design`
**Owner:** s2-core — design-lead `s2-core.planner` · adversarial design-reviewer `s2-core.implementer`
**Status:** r3 — r1 drafted against dispatch r2 (`.relays/s2/s2-core-design/DESIGN-orchestrator-planner-20260704-005310.md`, GRILL_REQUIRED: yes) + the de-provision supplement (`…-005315.md`); operator grill run 2026-07-04, GRILL_LOCK at §8 (`s2-grill-s2-core`); r2 = the two pair-review folds (§9, `DESIGN-REVIEW-implementer-20260704-014646.md`), pair-APPROVED at r2 (`…-021603.md`); r3 = the m-1 fidelity folds F-M1-1..3 (§9 second entry, `s2-fidelity-m1/SITREP-implementer-20260704-034158.md`) — bounded shape-fixes inside the m-1-prescribed forms, no mechanism change; r4 = state-label normalization (pair narrow-review `PLAN-REVIEW-implementer-20260704-040358.md` F2: "m-1-approved" → m-1-PRESCRIBED-pending-confirm throughout — the shapes become approved only when m-1's narrow re-review is on record). Both confirms now ON RECORD: pair narrow approve `s2-core-plan/PLAN-REVIEW-implementer-20260704-041850.md` (r6) + m-1 narrow confirm APPROVE `s2-fidelity-m1/SITREP-implementer-20260704-043529.md` (per the green-light relay `s2-core-plan/SITREP-orchestrator-planner-20260704-044011.md`) — the §4 shapes are m-1-APPROVED as folded
**Date:** 2026-07-04 · **Tier:** medium · **Evidence:** E1 (locked-spec + code cites) / E2 (battery + rg probes from the reconciled audits)
**Basis:** reconciled paired audits (RECONCILE.md 2026-07-04 entry; planner `s2-core-audit/AUDIT-planner-20260704-003144.md` + implementer `…-002839.md`); m-7 guide answers (`../master/relays/s2-guide-q1/SITREP-planner-20260704-004750.md` — Q1=(a), Q2=(i), both from locked text); the s2-dispatch charter (`../master/relays/s2-dispatch/PLAN-orchestrator-planner-20260703-223913.md`).
**Locked inputs (never designed here):** m-7 engine design (§2.1/§2.2/§5/§6/§10/§13); m-1 store contract (verbs, system fields, layout authority); ARCH §C4; the S1 design r5 decisions this slice builds on. Line refs `m-7 :N` / `m-1 :N` as in the audits. **m-1 owns the shape of every store-layout/record-kind item below — §4 is a PROPOSAL surface, not a decision.**

**Claim boundary (held in every sentence):** still provenance + transport, not verified work ("done" = `self_reported`; observe is Step 2). S2 adds durability hardening + the first governance primitive. The owed-item projection **guards recorded owed-items only — it does NOT make an unrecorded observation impossible to miss** (materialize-first). FIFO claims are **at-least-once intake, exactly-once EFFECT** (m-7 :58) — never unqualified "exactly-once". Confusion-resistance stays tool-mediated; **D5 shell-routed-confusion residual** stated beside every exclusivity-shaped claim. I-PH holds for every new surface this slice ships (quarantine incident text, phase-0 diagnostics, GC output, projection artifacts).

---

## 1. Scope

Build exactly: the reified recovery phase machine 0–4 (promoting every S1 piece into its slot), the single intake-writer + segmented durable FIFO, genesis + pinned-config digest + phase-0 fail-closed diagnostics, checksum quarantine disposition (recovery AND live path), drained-journal-segment GC (off by default; grill rows), the generalized obligation projection with the owed-item `record_kind` as its governance instance (OI-S1-F11-SWEEP = first customer), the crash-harness extension (recovery-boundary crashpoints, applicability map, new mutation-class arms), and the S2 exit-gate fixture set incl. F9/F11 re-run whole.

Out (escalate, never absorb): the ROADMAP OUT list verbatim (S3 registry/linter · MCP live-adapter · observe Step-2 · routing Step-3 · consumer schemas S4); **rendered-projection compaction** (grill row 1 — named-but-dormant, below); the §7 config-change record (S3 — S2 config is fixed at store creation, stated); outbox drain/external send; any m-1 contract amendment.

## 2. Decisions

### D-1 The reified recovery phase machine (m-7 §5 :89-95 verbatim; promotes S1 pieces)

`internal/recover` becomes a phase machine with named, crashpointed boundaries:

| phase | does | S1 piece promoted (already-closed inventory — audits, binding) |
|---|---|---|
| **0 validate-genesis** | genesis record present + config digest matches (D-3); fail ⇒ diagnostics mode (D-4), accept nothing | — (new) |
| **1 scan/quarantine** | staging cleanup; per-record checksum scan; mismatch ⇒ quarantine + incident (D-5) | staging cleanup (recover.go:19-24); `record.Verify` as the scan primitive |
| **2 rebuild projections** | **canonical-driven**: ensure every committed record's INDEX row / rendered `.md` / mailbox / outbox artifacts exist, replaying surviving redo segments as an optimization; canonical wins; idempotent; INDEX corrections appended, never rewritten | `RebuildProjections` + `appendUnique` (projections.go) — extended per D-6 canonical-sufficiency |
| **3 restore runtime tables** | binding table; outcome-by-`intake_id` set; content-hash map; obligation-completion indexes (D-7); parked-lane set; owed open-set | `seat.Open` (binding.go:44-63); the rest formalizes what S1 recomputed by rescanning per operation |
| **3.5 refill FIFO** | re-enqueue `intake − outcomes` in arrival order across all segments | `intake.Unconsumed` (journal.go:100-127) |
| **3.6 derived-work completion** | complete open auto-completable obligations (park/outbox/incident, D-7) — the S1 D-9 step-5 discipline, now an instance of the projection | `gate.Complete` (derived.go), refactored onto D-7 |
| **4 open** | mint the `Ready` capability; channels accept `submit` | rebuild-before-open ordering (main.go:78→:87), now structural |

- **No-authority-until-open as STRUCTURE, not call order — the two-capability split (r2, blocker F1):** `recover.Run(root, cfg)` returns a mode-typed terminal state: **`*engine.Ready`** (all phases complete) or **`*engine.Diagnostics`** (phase-0 failure — D-4). Both are unexported-field structs only the phase machine constructs. The gating is split so mutation is phase-4-gated WITHOUT suppressing the locked read-only surface: **`Submit` (the rendered tool), the intake writer, and the commit loop require `Ready`** — unconstructible in any other state; **`project`/`read` (+ the operator diagnostics report) require only a recovery terminal state** — `channel.ServeAuthenticated` takes a toolset factory built from either `Ready` (full 3-tool registry) or `Diagnostics` (read-only registry, no `submit` rendered). Channels open only in a terminal state; mid-recovery (phases 1–3) no channel is open in S2 — m-7 :95's "reads MAY be served" is permissive, and S2 exercises it only via the two terminal capabilities (stated, so the locked line is honored, not narrowed).
- **Phase-boundary crashpoints:** new registry names `recovery_post_phase0` … `recovery_post_phase4` + `pre/post` at each phase's mutation sites (quarantine evict, incident commit, derived-completion commits). The audits found ZERO crashpoints in the recovery path; the names-live-in-source pin (f11_test.go:49-63) extends its file list to recover.go. Recovery-of-recovery (crash mid-phase ⇒ rerun converges) is the ARIES property m-7 :92 licenses — every phase is idempotent by construction (scan, ensure-exists, rebuild-set, re-enqueue-set, complete-by-key).
- **Recovery reads ONLY the store** (guide constraint, binding): no docs-file read anywhere in the trusted recovery/genesis path; the engine never mints obligation records from out-of-store text.

### D-2 Single intake-writer + segmented durable FIFO (m-7 §2.1 :52 / §2.2 :56-58 verbatim)

- **`intake.Writer` goroutine — the ONE journal writer.** Channel handlers become literal mail-carriers: they send `{cmd, replyCh}` over a **bounded** channel to the Writer and await the typed outcome; **handlers never touch any file** (closes audit finding F-2: today per-connection goroutines call the unlocked `journal.Append` directly — main.go:90, server.go:87). The Writer: assigns `intake_id` from an in-memory monotonic counter (initialized at phase 3 — no more ReadAll-per-append), dedupes by the in-memory content-hash map (same init), appends+fsyncs to the **active segment**, then feeds the commit loop **in journal order** — ordering is now by construction, not by accident of goroutine scheduling.
- **Segmented journals:** `journal/intake/<seq>.jsonl` + `journal/redo/<seq>.jsonl` (naming = m-1 proposal, §4.7); append+fsync semantics unchanged (`post_intake_fsync` stays); rotation at the configured size (**default 4 MiB — grill row 2**), new crashpoints `pre/post_segment_rotate`. Rotation is bookkeeping only: a crash mid-rotation leaves either the old segment active or the new one empty-active — both converge (fixture).
- **Ack-on-outcome, content-hash retry dedupe, outcome-references-`intake_id` (R-2)** all carry forward unchanged (already-closed inventory).
- **Claim pin (blocker-by-construction):** at-least-once intake, exactly-once effect, zero stale re-emission — m-7 :58 wording verbatim in every FIFO-adjacent doc/fixture/string.

### D-3 Genesis + the pinned config (m-7 §10 :136, §7 :109-111; guide Q1=(a) + four sharpenings)

- **Genesis record** = the first canonical record, committed via the standard pivot by `store.Init` — **exact shape m-1-PRESCRIBED (r3, F-M1-1):** `Envelope{RelayID: "genesis", DispatchID: "genesis", From: "system", Role: "system", DeliveryState: "accepted", SchemaVersion: 1}` + `Headers{record_kind: "genesis", config_digest, address_space_seed, created_ts}` — `schema_version` lives ONLY in the envelope, never duplicated as a header. **Idempotence** = the existing duplicate-relay-id rejection (store.go:67-71): re-init against an existing genesis is a typed no-op refusal — nothing written, fixture-proven (the "genesis idempotent" gate line, pinned as this property).
- **Internal provenance convention (r3, F-M1-1 — m-1-prescribed):** conductor-authored records (genesis, incident, gc-marker, derived park/outbox) carry `From: "system", Role: "system"`, `DeliveryState: "accepted"` (incident: `"held"`), `SchemaVersion: 1`. **`system` is reserved: it is not a lane address** — `seat.Mint("system")` rejects typed, and no public-`submit`-path stamping can produce it (the channel binding is the only stamp source; no binding may hold the name). Owed items/dispositions ride the ordinary submit path stamped from the submitting channel; the real OI-S1-F11-SWEEP record stays `FROM=operator` via the operator channel (Q2 unchanged).
- **The pinned artifact set at S2** (Q1=(a)) — **store-root members, m-1-PRESCRIBED (r3, F-M1-3):** `config/fieldspec/registry.json` + `config/engine.json` (+ optional persisted `config/manifest.json`; the digest algorithm may compute the manifest in memory). `store.Init(root, sources)` accepts operator-supplied SOURCE paths but **materializes the pinned bytes under the store root BEFORE writing genesis**; phase-0 validation compares genesis against the store-root members, **never an outside path** — recovery reads only the store, now literally including its config. The members are not canonical records, not projections, not seat-visible resources, not domain-stamped sections in S2. `engine.json`: `{gc_enabled: false, segment_rotate_bytes: 4194304}` (**grill row 2 defaults**). **Deterministic digest input** (sharpening 3): SHA-256 over a canonical manifest — members sorted by name, each entry `name '\0' sha256(bytes) '\n'` — so the S3 change-record can extend the member set unambiguously.
- **Honest attribution** (sharpening 1): the pinned artifacts are described as **operator-ratified build config in the m-2-locked shape** — no domain-author stamp on any section a domain didn't author; the CQ-4b composed artifact + per-section stamps land with the consumer sections (S3/S4). Stated wherever the config is described (SWEEP class).
- **Claim scope** (sharpening 4): load-once / digest-verified / phase-0-fail-closed is claimed **for the pinned artifacts only**. Config evolution: **none in S2** — the §7 config-change record is S3; a store's config is fixed at `store.Init` for the life of the slice (stated plainly; grill row 2 consequence, operator-seen).
- **Load-once, restart-only** carries from S1 (D-5 there); the S1 claim-honesty gap ("digest-pinning rides genesis, which is S2") **closes here** — that sentence comes out of the S1 design's successor surfaces.

### D-4 Phase-0 diagnostics mode (guide sharpening 2: disposition, not error-exit)

On genesis missing / digest mismatch: the conductor **stays up, fail-closed** in the **`Diagnostics` capability** (the D-1 r2 split — the read-only terminal state the phase machine mints instead of `Ready`): per-seat registries render **`project`/`read` only** (no `submit` tool exists in the rendered registry — the guardrail idiom: absence, not refusal; structurally, `Submit`/writer/loop cannot be built from `Diagnostics`); no mutation of any store file; the operator channel additionally serves a typed **diagnostics report** — failure class + expected-vs-found digest values, **never a path** (I-PH; the report flows through the bounce formatter). "Summon operator" = that report + the process holding read-only until the operator intervenes and restarts. A non-zero exit is the REJECTED alternative (fails m-7 :90; §7 fold-log). Neighbor rule held: quarantine-never-brick is D-5's job — phase-0 failure is the one deliberate fail-closed stop (m-7 :90 "fail-closed at the root"), and even it serves reads.

### D-5 Quarantine replaces fail-stop (m-7 §5 :91, §6 :102; audit convergent finding F-1)

- **Phase-1 scan:** checksum-fail on a committed record ⇒ **evict** (one `rename()` of `records/<id>.json` → `quarantine/<id>.json`) + **ONE compound HELD-class incident record** committed via the standard pivot (references the quarantined `relay_id` + failure class; shape = m-1 proposal §4.4). Never silent-skip, never fail-stop.
- **Crash-window healing by the D-7 mechanism:** the quarantined file IS the durable intent; `open incident = quarantine/ member with no incident-record` is an auto-completable obligation instance — crash between evict and incident-commit converges at the next phase 3.6 (idempotent by `relay_id`). Crash between scan and evict re-detects at the next phase 1. One canonical record per disposition (F11-compatible: the evict rename is not a canonical-record commit; the incident is the mutation's one pivot).
- **The LIVE path is swept** (the audits' F-1 bite): after phase 1, all in-store records are verified; post-open operations run on the phase-3 runtime tables (D-6), so the per-operation full-store rescans that today propagate a mismatch as a brick (store.go:114-117 → lineage.go:51-53, derived.go, recover.go) disappear structurally. **Two typed read states, m-1-PRESCRIBED (r3, F-M1-2):** (i) live read of a corrupt file still present in `records/` ⇒ `checksum-mismatch` (carries `relay_id`; path-free) + enqueue the internal quarantine-disposition command — the DETECTION event; the reader never mutates. (ii) read AFTER eviction ⇒ `record-quarantined` (carries `relay_id`, `incident_id` when present, `failure_class: checksum-mismatch`; path-free) — the stable post-disposition state; never `checksum-mismatch` for this state. `read("incident-<relay_id>")` serves the HELD incident record normally. `Records()` and projection rebuilds operate over clean canonical records only — quarantined bytes are never re-included. Both states are API error classes, not delivery outcomes — the enum stays byte-exact `{accepted, rejected, held}`.
- **"The store never bricks"** is the property under test: corrupt one committed record ⇒ recovery completes, record quarantined + incident present, submit path alive, every other record served (fixture S2-K2).

### D-6 Canonical sufficiency + runtime tables (constraint 9 made structural)

- **Canonical-sufficiency rule (new, load-bearing for GC):** every derived artifact must be a **pure function of committed canonical records**. S1 violates it in one place, found while drafting: the outbox RECORD (derived.go:105-114) carries the item payload only in its redo intent — collect that redo segment and the outbox projection file is unrebuildable from canonical. Fix: the outbox/derived record **embeds its item payload in the record body** (as the held record already embeds its candidate); phase-2 regenerates any projection from records alone. This is what licenses redo-segment GC and makes "canonical wins unconditionally" true rather than aspirational. *(m-1 packet note §4.8 — body usage, no envelope/layout change.)*
- **Runtime tables (phase 3):** outcome-set, content-hash map, obligation-completion indexes, parked lanes, owed open-set — built once at recovery, maintained **incrementally by the loop** on each commit. Consequences: the F-3 hot path (gate.Complete's three full-store scans per submit, main.go:98 + derived.go:22/50/80) disappears as a consequence of the design, not as tuning; `appendUnique`'s full-file reread stays only on the cold rebuild path. The store remains truth; tables are derived caches whose rebuild IS recovery.

### D-7 The generalized obligation projection (constraint 6; guide advisory binding: ONE mechanism)

- **The mechanism (r2, clarification F2):** an obligation class = `{source predicate over DURABLE STORE FACTS, completion key, completer}`; **`open(class) = source-fact with no completion-record`**, computed at phase 3, maintained on the loop. A committed canonical record is the NORMAL source; two instances have explicitly-enumerated **file-backed durable intents** as their source — a `quarantine/` member (evicted-but-no-incident, D-5) and a GC-marker record naming still-present segment files (marker-but-not-unlinked, D-8). The enumeration is closed: exactly these two non-record sources exist; anything else is a design change. One mechanism either way — the completion walk is identical. Instances:
  | instance | source | completion record | completer |
  |---|---|---|---|
  | gate→park | gate-bearing `accepted` | park record (`parks_gate`) | **auto** (loop / phase 3.6) |
  | gate/held→outbox | gate or `held` record | outbox record (`(source_kind, source_record_ref)`) | auto |
  | quarantine→incident | quarantine/ member | incident record | auto |
  | **owed-item** | `record_kind: owed_item` | disposition record (`disposes_owed`) | **NONE — human/agent-authored only** |
- `gate.Complete` refactors onto this engine (C7's scan becomes an instance — s1 design §9.4 advisory, now landed); its S1 fixtures (C7, H-r3, the s1-close mirror legs) keep passing unchanged over the refactor — that is the no-regression proof the refactor rides on.
- **The owed-item instance is the governance primitive:** it has NO auto-completer by design — an open owed-item stays open until a record-authoring principal commits a disposition record. **Materialize-first, verbatim scope:** the projection makes silent drop impossible for a RECORDED item only; materializing is an intake/triage act by a record-authoring principal. An empty owed-projection at genesis is correct, not a bug (guide sharpening 4).
- **Surfaces:** the open set is served via the existing verbs — a derived projection artifact `projections/owed/OPEN.md` (rebuildable, canonical-sufficient) + `read`/`project` on the operator channel. No ODB/consumer-schema closure (S4-OUT): the artifact carries the typed-record fields only, schema explicitly not closed here.
- **OI-S1-F11-SWEEP (first customer; guide Q2=(i) + sharpenings):** the **operator authors** the owed-record through the operator channel during S2 IMPL — `FROM=operator` via the operator-relay stamp (locked m-1 §6 first-class-operator model), never a synthetic system stamp. Payload = the typed record `{owner, source, target_surface, disposition_path}` with `source` citing the s1 ledger entry (s1 RECONCILE.md :160-161) AND the guide's deviation-1 ruling (`../master/relays/s1-exit-gate/SITREP-planner-20260703-200827.md`); `target_surface` = the F11 full class×point sweep on the existing harness; `disposition_path` = the S2 exit gate. Fixture proof: surfaces as open → closes at the gate with a disposition record.

### D-8 GC — drained journal segments only (grill rows 1–2, operator-ratified 2026-07-04)

- **Target set (grill row 1):** **fully-drained journal segments ONLY** — an intake segment is drained when every entry has a committed outcome record; a redo segment when every entry's record is committed (canonical-sufficiency D-6 makes redo entries disposable then). **Rendered-projection compaction is explicitly NOT built** — the locked class (m-7 :137 "old rendered projections") has no definable membership in v3.0 (nothing supersedes a render under retain-everything); it stays **named-but-dormant** with this sentence as the record that we did not touch rendered copies, exactly the "stated, not built" pattern S1 used for unbuilt mutation classes. Canonical records are never GC'd in v3.0 (locked; the fixture asserts `records/` byte-untouched).
- **Posture (grill row 2):** collector **OFF by default** (`gc_enabled: false` in the pinned config) — out of the box nothing is ever deleted; rotation still runs (harmless bookkeeping). Fixtures create stores with tiny rotation sizes + collector on.
- **Mechanism — marker-first (materialize-first for deletion):** when enabled, a GC pass runs as a loop mutation post-open and after each rotation: compute the drained set against the runtime tables (serialized with commits — no scan/commit race), commit **ONE GC-marker canonical record** whose body names the collected segments (canonical-sufficient), then unlink the segment files. The marker is the durable intent: `open gc = marker naming a still-present segment` is another auto-completable obligation instance — crash mid-GC (before/after any unlink) converges at phase 3.6. Never unlink before the marker commits. New crashpoints `pre/post_gc_marker`, `pre/post_gc_unlink`.
- **Claim pins:** the exit-gate wording "never drops a live record" is claimed at the locked strength — GC never touches canonical records AT ALL; for journals, "live" = an outcome-less entry ⇒ its segment is not drained ⇒ untouchable (fixture asserts no outcome-less entry in any collected segment, checked pre-unlink).

### D-9 Crash-harness extension (constraint 8 — extension, not replacement)

- **Registry additions (additive):** recovery phase boundaries (D-1), rotation (D-2), quarantine evict/incident (D-5), GC marker/unlink (D-8). The child-SIGKILL machinery, `FRANK_TEST_RENAME_COUNTER`, wait-status assertion, and post-recovery asserts (f11_test.go) are reused unchanged.
- **The applicability map (audit gap 1, designed in):** a declared map `mutation class → expected-reachable crashpoint set`. The full sweep iterates **every class × every registered point**: an expected-reachable point must produce a SIGKILL child + convergent recovery; a not-expected point must produce a **clean-completion child + convergent store** (asserted, not skipped). Every cell is therefore covered by one of the two legs — this is what "full class×point sweep" honestly means, and the map itself is a fixture artifact (reviewable, so a wrongly-excluded point is visible).
- **Mutation classes swept at S2:** the six S1 classes + `{genesis, quarantine-disposition, gc-marker}`; owed-item and owed-disposition commits ride the submit-accept mutation path but appear as **explicit named rows in the applicability-map artifact** mapping to that path (r2, reviewer carry-forward) — the mapping is visible in the reviewed fixture artifact, never hidden behind the broad class name. **F9 and F11 re-run whole** under the phase machine (gate line G5); the S1 suite runs unmodified as the regression floor (G6).
- **OI-S1-F11-SWEEP is discharged by this design's sweep** and its disposition record commits at the S2 exit gate (D-7).

## 3. Acceptance criteria draft (fixture-keyed; E2 unless noted; ids prefixed S2-; gate ids G1–G6 per the planner audit's spec-to-exit-gate map, `AUDIT-planner-20260704-003144.md` §1)

| id | asserts | gate |
|---|---|---|
| S2-V1 | `store.Init` writes genesis (one record, one rename); re-init against existing genesis ⇒ typed refusal, nothing written | G3 |
| S2-V2 | digest deterministic: same artifact set ⇒ same digest across orderings; any member byte-flip ⇒ mismatch | G3 |
| S2-V3 | genesis missing/digest mismatch ⇒ **`Diagnostics` capability** (r2): `recover.Run` returns `Diagnostics`, not `Ready`; `submit` absent from every rendered registry while `project`/`read` serve, zero store mutation, operator diagnostics report present + path-free (I-PH grep) — **both halves asserted in one fixture** (read surface alive AND submit/writer/loop unconstructed), so neither half can be satisfied by deleting the other | G1/G3 |
| S2-K1 | corrupt one committed record ⇒ phase 1 evicts to quarantine/ + exactly one incident record; crash between evict and incident ⇒ next recovery completes it (idempotent) | G1 |
| S2-K2 | never-bricks: with one corrupt record, recovery completes, submit path serves, all other records readable; live `read` of a post-open-corrupted record ⇒ typed error + loop-executed disposition | G1/G6 |
| S2-W1 | N concurrent seats submitting ⇒ journal ids gap-free monotonic, journal order = loop order = outcome order; race detector clean | G2 |
| S2-W2 | F9-concurrent: crash mid-stream under concurrent multi-seat load ⇒ re-enqueue = exactly intake−outcomes in arrival order, zero re-emission (exactly-once EFFECT wording) | G2 |
| S2-W3 | rotation: segment rolls at configured size; crash at `pre/post_segment_rotate` ⇒ converges; entries never split/lost across the boundary | G2 |
| S2-X1 | GC off by default: collector never runs without `gc_enabled`; store byte-stable modulo new commits | G3 |
| S2-X2 | GC on: collects ONLY fully-drained segments; a segment with one outcome-less entry survives; `records/` hash-tree byte-identical before/after | G3 |
| S2-X3 | crash at `pre/post_gc_marker` and between marker and unlink ⇒ recovery converges; collected set = marker set exactly | G3 |
| S2-O1 | operator-authored owed-record commits via operator channel ⇒ surfaces in the open set (projection artifact + read/project) | G4 |
| S2-O2 | disposition record commits ⇒ item drops from open set; a second disposition for the same owed id ⇒ typed reject (one-shot idiom) | G4 |
| S2-O3 | OI-S1-F11-SWEEP end-to-end: the real record (guide-specified payload) open at IMPL → dispositioned at the S2 exit gate | G4/G5 |
| S2-O4 | empty owed-projection at genesis is served correctly (not an error) | G4 |
| S2-O5 | ONE mechanism: gate→park, gate/held→outbox, quarantine→incident all complete via the obligation engine; the S1 C7/H fixtures pass unchanged over the refactor | G4/G6 |
| S2-PM1 | phase-boundary crash matrix: SIGKILL at every `recovery_*` crashpoint ⇒ rerun converges (recovery-of-recovery); no lost/double intake or delivery across any boundary | G1 |
| S2-PM2 | structural no-authority-before-open under the **two-capability split** (r2): `Submit`/intake-writer/commit-loop unconstructible without the phase-4 `Ready` (compile-level); the read-only registry constructible from `Diagnostics` alone; a runtime probe confirms no channel serves before a terminal state exists | G1 |
| S2-F11 | the full class×point sweep per the applicability map (every cell = crash-leg or clean-completion-leg); F9 + F11 whole re-run under the phase machine | G5 |
| S2-RE | entire S1 fixture suite green with its **invariant assertions untouched** (serialized-loop kill, crash-atomicity, I-PH, enum byte-exactness, guardrail); mechanical call-site updates only where engine signatures changed (recover.Run→Ready; store setup gains store.Init for genesis) — any assertion change is a review blocker | G6 |
| S2-SWEEP | claim-sweep classes over everything S2 ships (docs, strings, fixtures): exactly-once-EFFECT wording, GC locked-strength wording, materialize-first beside projection claims, D5 beside exclusivity claims, I-PH on new surfaces | G6 |

## 4. m-1 fidelity surface (r3: the verdict's prescriptions folded from `s2-fidelity-m1/SITREP-implementer-20260704-034158.md` — items marked ✔ carry the **m-1-PRESCRIBED shape verbatim** (the homes table + per-item answers); they become m-1-APPROVED only when m-1's narrow re-review confirm is on record — **pending as of r4**; the F-M1-1 homes table governs: `schema_version` = envelope-only; `record_kind` + typed fields = headers; payloads = body JSON; internal provenance = the D-3 `system` convention)

1. ✔ **`record_kind: owed_item`** — `Headers{record_kind, owner, source, target_surface, disposition_path}`; validation + the open-set projection read the HEADERS (body may be empty/narrative); normal submit stamping applies; envelope compliance unchanged (checksum, envelope SchemaVersion, intake_id when loop-committed).
2. ✔ **Disposition record** — `Headers{record_kind: owed_disposition, disposes_owed: <owed relay_id>}`; one disposition per owed id (duplicate ⇒ typed reject).
3. ✔ **Genesis record** — the exact D-3 shape (m-1-prescribed): fixed `relay_id`/`DispatchID` `genesis`, `system/system/accepted/SchemaVersion 1` envelope, headers `{record_kind: genesis, config_digest, address_space_seed, created_ts}`; record #1 in `records/`.
4. ✔ **Incident record** — `system/system/held` envelope + `Headers{record_kind: incident, quarantined_ref, failure_class}`; compound single-pivot.
5. ✔ **`quarantine/`** — store-root member (locked location m-7 :91); eviction = name-preserving rename; post-eviction `read` ⇒ `record-quarantined` (never `checksum-mismatch` for that state) per D-5's two typed states.
6. ✔ **GC-marker record** — `Headers{record_kind: gc_marker}`, `system/system/accepted` envelope, Body JSON naming the collected segments (canonical-sufficient); one per GC pass; canonical records never GC'd.
7. ✔ **Journal segmentation layout** — `journal/intake/000001.jsonl` + `journal/redo/000001.jsonl` shape, six-digit zero-padded monotonic `<seq>`, highest = active; rotation at pinned-config bytes (m-1: approve as proposed).
8. ✔ **Derived-record body embedding** (canonical-sufficiency, D-6) — Body JSON is canonical record content; required for canonical sufficiency (m-1: approve).
9. ✔ **`projections/owed/OPEN.md`** — derived, rebuildable, carries no independent authority; empty-at-genesis valid (m-1: approve).
10. ✔ **Pinned config** — STORE-ROOT members `config/fieldspec/registry.json` + `config/engine.json` (+ optional `config/manifest.json`) per D-3 (F-M1-3); operator-supplied paths are Init-time SOURCES only, never phase-0 dependencies; operator-ratified build config in the m-2-locked shape, no domain-author stamps.

## 5. Boundary contract

Writes: append-only store records (byte-exact `{accepted, rejected, held}` — unchanged enum) incl. the new kinds above; quarantine evictions; GC unlinks of drained segments only. Reads: the m-1 verbs; the locked m-7 engine surfaces. Target entity: the recovered/opened store + the owed open-set. Downstream consumers: seats via `project()`/`read`; the operator (diagnostics, owed items); the S2 exit-gate fixtures (E2). Proof: §3 table. No-consumer action: not applicable (all surfaces have named consumers above).

## 6. Open items carried to PLAN (none re-architect)

1. Crashpoint name list finalization (exact registry strings) + the applicability map as a reviewed fixture artifact.
2. `Ready`-capability plumbing through cmd/frank/main.go (assembly order refactor; no behavior change pre-phase-4).
3. m-1 fidelity packet = §4 verbatim, routed at PLAN per the dispatch; nothing store-shape-touching dispatches before m-1's approve.
4. The operator's OI-S1-F11-SWEEP submit is an IMPL-phase step (Q2 sharpening 4 sequencing) — the plan schedules it after the operator channel is up in the fixture/live store, before the exit-gate run.
5. Fixture-to-task map + SCOPE_DIFF file list.
6. S1 fixture call-site migration inventory (fixtures create stores via `store.Init` so phase-0 passes; `recover.Run` call sites take the new signature) — mechanical only; every S1 invariant assertion stays byte-identical (S2-RE blocker otherwise).

## 7. Rejected alternatives (log)

Render compaction now, or renders-as-cache with regenerate-on-miss (B/B-lite) — **operator-rejected at grill** (row 1): no definable "old" membership in v3.0; segments-only keeps never-drops trivially strong; revisit when supersession exists. · GC on by default — cleanup-by-default contradicts retain-everything; operator picked off (row 2). · Error-exit on phase-0 failure — fails m-7 :90 (guide sharpening 2); diagnostics disposition instead. · Config values outside the digest — violates the one-integrity-root rule (§7 locked); pinned, with S2 immutability stated. · Config-change record in S2 — S3's; building it now = scope creep on the OUT fence. · Engine-transcribed owed-record from the s1 ledger (Q2 shape (ii)) — guide-rejected: docs-file read inside the trusted recovery path breaks dumb-replay idempotence + store-is-truth; fakes provenance. · Two mechanisms (derived-work scan + a separate owed projection) — guide advisory says instance-of; one engine (D-7). · Grandfathering genesis-less stores — no persistent store exists to grandfather (repo-verified); phase-0 requires genesis unconditionally. · Per-operation full-store rescans kept (status quo) — carries the F-1 brick and the F-3 O(N²) hot path; runtime tables are the designed shape. · Live quarantine executed by the reader — readers never mutate; disposition rides the FIFO to the loop.

## 8. GRILL_LOCK

```text
GRILL_LOCK_ID: s2-grill-s2-core
GRILL_REQUIRED: yes
GRILL_SOURCE:
- plan/design/audit relay read: s2-core-design dispatch r2 (…-005310) + de-provision supplement (…-005315); reconciled s2-core audits; s2-dispatch (r2); s2-guide-q1 answer (…-004750); s1 design r5 + s1 RECONCILE
- code/docs inspected: frank/ at 6ceeb5d (store/intake/engine/recover/gate/fsio/crashpoint + fixtures); m-7 :52/:56-58/:89-95/:109-111/:134-137/:172; m-1 :122-145; ARCH §C4
- questions answered from codebase: no persistent frank store exists outside test tmpdirs ⇒ no genesis-less grandfathering needed (phase-0 requires genesis unconditionally); quarantine location + GC gating + config-in-digest answered from locked text, not asked
- questions asked operator: GC target scope; GC default posture + rotation size

Resolved decisions:
- S2 GC target scope — drained journal segments ONLY; rendered-projection compaction NOT built, documented dormant (no definable "old" membership under v3.0 retain-everything) — source operator (2026-07-04, "lets just do journals only, and then document that we did NOT touch rendered copies")
- GC default posture — collector OFF by default (gc_enabled: false); rotation always on — source operator ("off by default … wont get cleared by default" confirmed)
- Journal rotation size default — 4 MiB (fixtures use tiny configured sizes) — source operator ("sure lets do 4mib then", after the relays≈messages clarification)
- GC/rotation values live INSIDE the digest-pinned per-store config; fixed at store creation for all of S2 (config-change record = S3) — source locked text (§7 one-integrity-root; consequence stated to operator during grill)
- Q1 genesis digest scope = (a) pin-what-exists, deterministic manifest, honest attribution, claim scoped — source guide (resolved-by-guide row; s2-guide-q1 …-004750)
- Q2 owed-record authorship = (i) operator-authored via operator channel; recovery reads only the store — source guide (resolved-by-guide row)
- m-1 proposal boundary — every store-layout/record-shape item is a PROPOSAL routed to m-1 fidelity before dispatch; the pair fixes nothing — source dispatch (F2 condition, restated; not an operator question)

Rejected alternatives:
- render compaction now / renders-as-cache (B, B-lite) — operator-rejected for S2; revisit when supersession exists
- GC on by default — contradicts retain-everything posture
- error-exit on phase-0 failure — fails the locked disposition (m-7 :90)
- engine-transcribed owed-record (Q2 shape (ii)) — guide-rejected (out-of-store read in the trusted path)

Still operator-owned:
- S2-close sign-off (exercised separately at the exit gate — charter)
- future retention/GC tuning values once the S3 config-change path exists (operator config policy per m-7 §14)

Design-lock impact:
- D-8 carries the two operator rows verbatim; D-3 carries the defaults + S2-immutability statement; §1/§7 carry the render-compaction dormancy line; the DESIGN_LOCK_ID for this doc must reference GRILL_LOCK_ID s2-grill-s2-core
```

## 9. r2 fold-log (s2-core.implementer DESIGN-REVIEW `.relays/s2/s2-core-design/DESIGN-REVIEW-implementer-20260704-014646.md`, verdict must-revise — both findings verified correct by the planner before folding)

1. **F1 (blocker) — `Ready` over-gated the read-only/diagnostics surface.** r1's D-1 required `Ready` for `channel.ServeAuthenticated` while D-4 required channel-served read-only diagnostics exactly when phase 4 is never reached — a structural contradiction against m-7 :89-95's read-only-before-open allowance. Folded: the D-1 two-capability split (`Ready` gates `Submit`/writer/loop; `Diagnostics` licenses the read-only registry; channels open only in a terminal state; the phases-1–3 "may serve reads" line honored as permissive). D-4 names the capability; S2-V3 asserts both halves in one fixture; S2-PM2 names the split. No mechanism change beyond the capability boundary.
2. **F2 (required clarification) — obligation source widened to durable store facts.** r1's "source predicate over committed records" excluded two of its own instances (quarantine/ member; GC marker + still-present segments). Folded: D-7 source = durable store fact, committed record as the normal case, the two file-backed intents explicitly enumerated and CLOSED (a third non-record source = design change). One mechanism unchanged.
3. **Carry-forwards adopted (non-blocking):** D-9 now requires owed-item/owed-disposition as explicit named rows in the applicability-map artifact; the §4 m-1 fidelity gate stays hard (already §6.3 — unchanged).

No other section changed; the GRILL_LOCK (§8) rows and every operator/guide decision are untouched. Resent for re-review on the same DESIGN_DOC_ID.

### r3 fold-log (m-1 fidelity `s2-fidelity-m1/SITREP-implementer-20260704-034158.md`, verdict must-revise — all three findings verified correct by the planner before folding)

1. **F-M1-1 — canonical field homes + internal provenance.** Folded as the m-1 homes table verbatim: `schema_version` envelope-only (my r2 had it in genesis headers — wrong home); `record_kind` tokens `{owed_item, owed_disposition, genesis, incident, gc_marker}` + all typed fields in Headers; payloads in Body JSON; the D-3 internal-provenance convention (`system/system`, accepted — incident held); `system` reserved (mint-reject + unstampable from the public submit path). D-3 genesis block now carries the exact m-1-prescribed shape; §4 items 1–4, 6, 8 carry the prescribed shapes verbatim (pending m-1's narrow re-review confirm — r4 label normalization).
2. **F-M1-2 — two typed corruption states.** Folded into D-5: `checksum-mismatch` = the detection event (corrupt file still in `records/`); `record-quarantined` = the stable post-eviction state (with `incident_id`, `failure_class`); incident records readable; `Records()`/rebuilds over clean records only; both are API error classes, no fourth delivery outcome.
3. **F-M1-3 — store-root config placement.** Folded into D-3: pinned members live under `<root>/config/`; `store.Init` materializes source bytes before genesis; phase-0 compares store-root members only — closing the coherence gap with the guide's recovery-reads-only-the-store constraint (an external config path in phase-0 WAS an out-of-store read; caught by m-1, verified by me).

Already-approved S2 mechanics otherwise unchanged (m-1 dispatch condition 2 honored); GRILL_LOCK rows untouched. Both narrow re-reviews subsequently landed APPROVE on exactly these folded lines (pair r6 `…-041850.md`; m-1 confirm `s2-fidelity-m1/SITREP-implementer-20260704-043529.md`).

## 10. Claim-sweep note (m-7 §16 classes, run over this doc)

Exclusivity-shaped claims in this doc: "handlers never touch any file" (tool/engine-internal control-flow, m-7 :52 verbatim); "GC never touches canonical records" (locked posture, fixture-scoped); "no submit tool exists in the rendered registry" (tool-surface grain). Each is a governance/tool-surface claim; the D5 shell-routed residual (same-uid processes can reach store/journals/sockets outside the tool surface) applies to all of them and is accepted Step-1 posture (ARCH §C4.3). No "by construction" claim is made here beyond citing the locked §2.4 serialized-loop kill; the projection's guarantee is scoped to recorded items in every mention; FIFO claims say exactly-once EFFECT throughout.
