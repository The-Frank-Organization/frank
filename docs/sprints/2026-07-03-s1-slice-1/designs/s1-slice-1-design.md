# S1 Slice-1 — Design of the thin end-to-end conductor relay (the code shape)

**DESIGN_DOC_ID:** `s1-slice-1-design`
**Owner:** s1-core — design-lead `s1-core.planner` · adversarial design-reviewer `s1-core.implementer`
**Status:** DRAFT r2 — r1 must-revise blockers folded (§8 fold-log); awaiting Implementer re-review on the same DESIGN_DOC_ID
**Date:** 2026-07-03 · **Tier:** medium · **Evidence:** E1 (spec cites) / E2 (repo + corpus probes from the reconciled audits)
**Basis:** DESIGN dispatch `.relays/s1/s1-core-design/DESIGN-orchestrator-planner-20260703-140843.md` + de-provision supplement `…-142800.md`; guide answers `master/relays/s1-guide-q1/SITREP-planner-20260703-141628.md` (Q-A + Q-2 CONFIRMED — no provisional sections remain); reconciled paired audits (RECONCILE.md entries 2–3).
**Locked inputs (never designed here):** m-1 store API + system-field contract; m-2 FieldSpec envelope + predicate vocabulary; m-7 engine design (topology, pivot, fault taxonomy, guardrail); ARCHITECTURE §C4. This doc designs the *implementation* of the S1 slice against them. Line refs `m-1 :N` / `m-2 :N` / `m-7 :N` are into the locked design docs as in the audits.

**Claim boundary (held in every sentence):** S1 = provenance + transport, not verified work. The only operationally-live by-construction claims are the serialized-loop double-accept kill (m-7 §2.4) and the constrained-grammar R2. Everything identity/isolation-shaped below is **tool-mediated confusion-resistance**: it removes affordance, not access. **D5 residual, stated once here and beside every exclusivity-shaped claim below:** a same-uid shell-bearing seat can reach the store, sockets, and config outside the tool surface; that is the accepted Step-1 residual (ARCHITECTURE §C4.3).

---

## 1. Scope

Build exactly: `mint → connect → render-MVP-form → submit → stamp → validate → lineage → append(crash-atomic) → project → deliver → gate-outbox`, plus the dumb-replay recovery the guide confirmed, the S1-minimal dissolved-linter replay, and the exit-gate fixture set (namespace B1-B4 / A1-A4 / C1-C6 / R1 / P1 / L1 / W1, plus m-7 fixture ids F3/F4/F9/F10/F11 where they name the same thing).

Out (escalate, never absorb): everything on the ROADMAP scope-OUT list; genesis/quarantine/GC/segment-rotation/phase-0→4 machinery; m-2 §9 migrator machinery (`schema_version` stamping only); outbox drain/external send; observe layer; ⑤ `model_name` slot (S4-bound §C4 carry — cited, not scoped).

## 2. Decisions

### D-1 Conductor stack: Go (operator decision, 2026-07-03, inline during design)

Single static binary, no runtime dependency surface; `os.Rename` / `File.Sync()` / dir-handle `Sync()` give exact control of the fsync-before-rename discipline (F10); goroutines + channels map 1:1 onto the locked topology (D-2); crash fixtures = `go test` spawning the conductor binary and SIGKILLing it at named points — fully deterministic.

Candidates compared (dispatch-required criteria: kill -9 fixture-drivability · fsync/rename fidelity · dependency surface · team familiarity):
- **Go — SELECTED** (operator). Strong on all four; MCP served by the official MCP Go SDK (see D-3 fallback).
- **Python + official MCP SDK** — planner's pre-decision recommendation (team tooling is Python); rejected by operator choice; nothing in the locked spec favored it structurally.
- **TypeScript + MCP SDK** — most mature MCP SDK, natural single-threaded loop; rejected: heavier dependency surface, weaker syscall-grain control ergonomics.
- **Rust + rmcp** — strongest type-level single-writer enforcement; rejected: slowest iteration for a thin slice, and the S1 claim boundary (confusion-resistant, not adversarial) does not need type-level rigor yet.

### D-2 Process/concurrency shape (implements m-7 §1/§2 verbatim)

One `frank` conductor process:
- **[A] per-seat channel handler goroutines** — accept + authenticate on the seat's channel (D-3), present the 3-tool registry, forward each command over one **bounded** `chan` to the intake-writer with a per-command reply channel; **handlers never touch any file**.
- **[B] single intake-writer goroutine** — sole writer of the intake journal: assigns `intake_id`, appends + fsyncs the entry **before** it enters the in-memory FIFO, then feeds the commit loop in arrival order. Content-hash request key per entry for retry dedupe (m-7 :57).
- **[C] single commit-loop goroutine** — sole writer of store files. Consumes the FIFO one mutation at a time, start to finish: stamp → validate → lineage → stage → checksum → fsync → **rename pivot** → dir-fsync → redo-journal projections → typed outcome (referencing `intake_id`) → reply. No other goroutine opens a store file for write. *(Single-writer is enforced by package layout — the store-write API is unexported outside the loop's package — plus fixture; this is discipline, not an OS guarantee; D5 residual applies to out-of-process writers.)*
- **[D] store on disk** (D-4). **[E] reads + delivery** — `project()`/`read()` served concurrently from committed immutable records; delivery = one `write()` of a nudge frame onto the recipient seat's connected channel, sequenced after commit.

fsyncs run synchronously on their goroutine (the operator-grilled simplicity rule; at S1 volume a blocked loop during fsync is the accepted cost, m-7 :61).

### D-3 Channels, identity, mint/connect (Q-E; proposes the m-1 "fork-2 infra call" realization — recorded for the m-1 fidelity review)

- **Transport:** one **Unix domain socket per seat**, under a conductor-owned runtime dir, carrying MCP (JSON-RPC 2.0). Primary: the official **MCP Go SDK** with a socket-stream transport; **PLAN-time capability check** — if the SDK's server transport cannot ride a per-seat socket cleanly, fallback is a minimal self-hosted MCP framing (the surface is three tools; the protocol subset is small). This check is the plan's first task, not a design unknown that re-architects anything. **Fallback invariants (r2):** whichever framing ships must preserve per-seat channels, channel-stamped identity, and the exact `{submit, project, read}` per-seat registry — the fallback swaps wire framing only.
- **mint_seat (conductor-internal, no agent-facing verb):** operator/spawn tooling asks the conductor to mint; the conductor creates the seat's socket endpoint + an unguessable credential (≥128-bit random), records `seat ↔ credential` in the **persisted binding table** (conductor-owned file, store-adjacent, never a seat surface; written via the same stage→fsync→rename primitive as every conductor file write — r2, reviewer watch item), and provisions the mailbox. Credential + socket path are delivered only to that seat's lane via its private env/fd at spawn (DI-2 realization; **confusion-resistant only — D5:** same-uid processes can enumerate sockets).
- **connect:** the seat presents its credential on its socket; unbound/wrong credential ⇒ reject, nothing staged. Re-attach with the same credential resolves to the same seat, no re-mint (m-7 §8.5, CQ-6 base). `FROM`/`ROLE` are stamped from this binding — payload values ignored byte-for-byte.
- **Operator channel:** a separate socket + operator credential held by the operator's own tooling, never delivered to any lane (m-1 :254-259 carry). `FROM: operator` is stampable only via it. Verdict submissions ride it.
- **Per-seat tool registry:** exactly **`submit` / `project` / `read`**, constructed conductor-side per seat; `submit`'s input schema **is** the rendered form for that seat (D-5) — forbidden options absent from the schema itself. No tool, description, or result carries a store/config/outbox/operator-channel path or config value (the m-7 §8.4 absence set; the I-PH discipline in D-8).
- **Fixture seats:** E2 fixtures drive scripted Go test seats over real sockets. Real-runtime bring-up (Claude Code SDK streaming, codex app-server) is operator-gated spike territory, not an S1 gate (m-7 :122).

### D-4 Store shape + the commit pivot (implements m-7 §4 Package A; guide constraints R-1/R-2 built in)

```
<STORE_ROOT>/
  records/<relay_id>.json      # canonical: self-contained, CHECKSUMMED (R-1), immutable
  staging/                     # pre-pivot staging; torn files deleted at recovery
  journal/intake.jsonl         # append+fsync ahead of the FIFO; intake_id per entry
  journal/redo.jsonl           # projection intents, staged+fsynced BEFORE the pivot
  projections/INDEX.md         # append-only (corrections appended, never rewritten)
  projections/relays/<DISPATCH_ID>/<PHASE>-<ROLE>-<ts>.md   # rendered views (v2.8.8 layout, REUSE-AS-SPEC'D)
  mailboxes/<seat>.jsonl       # per-seat inbox projections
  outbox/<item_id>.json        # store-visible ODB queue items (O-1)
  binding/seats.json           # persisted binding table (conductor-owned)
```

- **Canonical record** = one JSON file: `{envelope, headers, body, x_fields, checksum}`; checksum = SHA-256 over the canonical serialization minus the checksum field (**R-1** — the field is S1 format; the quarantine disposition consuming a mismatch is S2). Naming carries `relay_id`.
- **Commit** = stage in `staging/` → fsync file → one atomic `rename()` into `records/` → dir-fsync → replay redo intents into projections. **Presence = committed.** Projections are derived; canonical wins unconditionally.
- **Every outcome record references the `intake_id` it consumes** (**R-2**); "consumed" is durably recorded by the same pivot — no separate mark-consumed write exists.
- **F11 one-pivot-per-mutation, S1 mutation classes:** `{submit-accept, submit-reject, held (compound, candidate embedded), operator-verdict, park/wake transition, outbox enqueue}` — each commits exactly one canonical record via exactly one rename. Genesis/config-change/GC classes have no S1 instances (stated, not built).

### D-5 MVP FieldSpec (Q-C) + registry realization

**Registry realization:** a conductor-owned data file (`fieldspec.json`, m-2 §4 shape: id/layer/owner/type/enum_set/gate_referenceable/seat_scope/required_when/visible_when/fill_constraints/consumers/lineage_role), loaded **once at startup, restart-only**, path and values absent from every seat surface. **Claim-honest:** S1 does *not* claim config integrity — digest-pinning rides genesis, which is S2; what S1 keeps is load-once/restart-only/absent-from-seat-surfaces (the m-7 §7 properties that don't depend on genesis).

**Field set** (selection criterion per the dispatch: maximize covered m-2 §10a/§10b classes per unit of build):

| field | layer | owner / fill | type + notes |
|---|---|---|---|
| FROM, ROLE | envelope | system_only | stamped from the channel binding; payload ignored |
| relay_id, DISPATCH_ID, timestamp, schema_version | envelope | system_only | courier-assigned; schema_version stamped per record (no migrator machinery) |
| certification | envelope | system_only | null-reserved, present in every record |
| PARENT_DISPATCH_ID | envelope | parent_picker | conductor-derived candidate set + default = woken-on relay; free-typed parent rejected |
| TO, CC | envelope | recipient_picker | validated ∈ the minted address space (incl. `operator`) |
| PHASE | header | agent_enum_pick | v2.8.8 11-token enum |
| AUTHORITY | header | seat_scoped_enum | full enum on operator/orchestrator forms; pair-seat forms omit `merge-gated` (an A2 forbidden option) |
| CEREMONY_TIER | header | agent_enum_pick | `{tiny, small, medium, large, production-risk}` |
| EVIDENCE_TARGET | header | agent_enum_pick | `{E1..E4}`; Step-1-required (intent field, not observe-owned) |
| HUMAN_GATE_REQUIRED | header | monotonic hybrid | floor = MAX(system baseline, agent raise, known-A detector); RAISE-only |
| gate_category | header | agent_enum_pick, monotonic-toward-A | **the FULL frozen §J2 default set, byte-exact (r2 — blocker 5):** A = `{merge_to_protected, irreversible_write, residual_risk_acceptance, live_verify_skip, ceremony_downgrade, authz_security, product_semantics, scope_expansion}`; B = `{merge_feature_to_feature, routing, sequencing, scope_within_bounds}`; hardcoded `other`→A fail-safe (ARCHITECTURE.md:110-115). `routing_escalation` is NOT a member (owed §C4 carry, ARCHITECTURE.md:481 — cited, not added). Ships as registry data with the A/B map + a default protected-branch set; affordance constrained to `[floor, A]`. Fixtures exercise minimal cases (one A, one B, `other`→A, known-A raise) — adopting the frozen enum is config data, not scope creep |
| gate_category_raised | header | system computed | bool; recorded whenever the floor wins (the ③ guardrail-adjacent portion, D-7) |
| grant | header | seat_scoped_enum | `{dispatch-impl, dispatch-merge}`; rendered **only on operator/orchestrator forms** — the frozen m-2 rule verbatim ("absent from pair-seat forms, present only on the operator/orchestrator form, gated to the right PHASE", m-2 :177); `dispatch-merge` visible only in `PHASE: MERGE-GATE`. **Explicit S1 narrowing (r2 — blocker 2, routed for orchestrator ratification via the design-completion SITREP):** S1 does NOT implement conditional pair-Planner delegated-dispatch rendering; the pair-Planner DISPATCH-IMPL lineage walk m-2 §10c preserves (m-2 :167) lands with the full lineage engine (S3). In S1, live grant issuance is operator/orchestrator-form only — a deliberate, stated scope line, not a silent foreclosure |
| delivery_state | header | system computed | byte-exact `{accepted, rejected, held}` — the only outcome tokens anywhere in the codebase |
| failing_edge | header | system computed | reason text; carries the lineage-bounce distinction + "re-render" (D-6); I-PH-filtered |
| SUBJECT | header | free_text | required |
| body | body | free_text | unconstrained (Format-Tax line) |
| X-* | body | free_text | namespace: `consumers: []`, `lineage_role: none`, gate_referenceable false |

**Excluded and why:** all observe-owned fields (`ACTIONS_GIT_REF`, `FINAL_GIT_STATUS_SHORT`, `*_RESULT`, …) — step-gated off in Step-1 (CQ-1); Step-1 records carry no observe gate. `ESCALATION_SCAN`/`SCOPE_DIFF`/`FOLD_SCOPE` row-arrays — their §10b check classes land with the fuller registry (S3); the R1 report names them as the uncovered remainder. Routing/archetype fields — S4. `slot_in` — reserved atom, no S1 values.

**Render:** `render(seat, phase, tier)` computes the visible field set + per-seat options from the registry and emits it as `submit`'s JSON input schema, stamped with a **form digest**; submit re-validates authoritatively in the loop (render is advisory, m-7 §8.2); a submit carrying a stale form digest bounces with reason `re-render` (STEP-1-KICKOFF :53).

### D-6 Validate + lineage (minimal engines)

- **Form-validation** (in-loop, authoritative): required-set via the bounded predicate evaluator (m-2 §5 atoms; S1 needs `phase_in`, `authority_in`, `seat_is`, `field:<id> == / present` — evaluator built to the closed vocabulary, unused atoms parse-rejected), enum membership, seat-scope, monotonic floors (HUMAN_GATE_REQUIRED, gate_category).
- **Authority-bearing (S1 rule — r2, blocker 3: a PESSIMISTIC SUPERSET of the frozen m-2 :76 set, derived from MVP-visible fields only):** a candidate is authority-bearing iff ANY of: (a) `grant` present; (b) gate-bearing (`HUMAN_GATE_REQUIRED: yes` or `gate_category ∈ A-set ∪ {other}`); (c) `PHASE ∈ {PLAN, IMPL, REVIEW-FOLD, MERGE-GATE, LIVE-VERIFY}`; (d) `AUTHORITY ∈ {implementation, merge-gated, live-verify, fold-in-only}`; (e) stamped `ROLE ∈ {orchestrator-planner}` with `PHASE ∉ {SITREP, RECONCILE}`. Mapping to m-2 :76's classes: design-doc PLAN locks + orchestrator PLAN → (c); delegated dispatch/merge grants → (a); substantive IMPL reports → (c) (S1 has no ACTIONS fields to distinguish substantive, so ALL IMPL is treated authority-bearing — pessimistic); merge claims → (c)/(d); orchestrator-authority relays → (e); A-category gates → (b). Over-inclusion is the safe direction: the costs are a lineage parent check (which parent_picker makes trivially passable for honest records) and fault→`held` instead of `rejected` (the pessimistic disposition). The precise `record_kind` taxonomy lands with the fuller registry (S3); the S1 proxy is deliberately **never narrower** than m-2 :76.
- **Lineage gate** (blocking only for authority-bearing): walks `persisted-accepted ∪ {candidate}`; S1 blocking edges: PARENT must resolve in the accepted graph, and an **operator-verdict record must parent the parked gate record it resolves** — the S1 one-shot authority (D-7/A4). Non-authority records commit on form pass (lineage is a no-op for them).
- **Bounce ergonomics** (STEP-1-KICKOFF :52): `failing_edge` distinguishes `parent-unknown — possibly in-flight; recompose + resubmit` from `parent-invalid — dead edge` (reason text, not a new state). All bounce/error text is built by one central formatter that has no store-path inputs (I-PH by construction, D-8).
- **Internal-fault disposition** (m-7 §6): every trusted check in the loop runs under panic-recover + hard timeout; fault on an authority-bearing candidate ⇒ **one compound `held` record embedding the candidate bytes**, referencing `intake_id`, one pivot; non-authority ⇒ `rejected` + bounce naming the fault edge. The fault path itself passes through no gate that could hold it; the loop keeps serving.

### D-7 Gate → park/wake → outbox (B4/W1; guide constraints O-1..O-3; the ③ portion)

- A gate-bearing candidate commits as **`accepted`** (ordinary human-gate parking is never `held` — m-2 :376). The loop then commits, as separate single-pivot mutations in the same FIFO turn: a **park-transition record** (lane → `parked_waiting_human`) and an **outbox-enqueue item** — each keyed by `gate_record_ref` (the gate record's `relay_id`), which is the idempotence discriminator.
- **Crash-window healing (r2 — blocker 1):** the committed gate-bearing `accepted` record IS the durable derived-work intent. Recovery (D-9 step 5) scans committed gate-bearing accepted records and idempotently commits any missing park-transition / outbox-enqueue follow-ups (dedupe by `gate_record_ref` — a follow-up that already exists is never re-committed) **before channels open** (R-3 extended to derived-work completion). A crash after gate-accept and before follow-up completion therefore converges to exactly one park record + one outbox item; a re-crash during completion re-converges (idempotent by the same key). Fixture C7.
- **ODB item envelope (O-3, open):** `{item_id, gate_record_ref (relay_id), seat, gate_category, created_ts, schema_version}` — the minimal envelope S1 needs; the schema is m-6/m-3-owned and explicitly **not closed** here; **no `model_name` slot is pre-built** (⑤ is the S4-bound §C4 carry, cited per the guide's ledger note).
- **O-1:** outbox enqueue is a loop mutation committing a store-visible queue item through the standard pivot; no side path can produce an outbox item (**governance-surface claim; D5 residual:** an out-of-tool-surface process could write the directory directly).
- **O-2:** S1 builds no drain and no external send; nothing in code, error text, or docs claims the egress scan is live (it is dormant by locked posture, m-7 :132).
- **Park semantics:** parked lane state lives in store records, not memory; while parked the conductor delivers nothing to that lane. **Wake:** the operator verdict record commits through the loop (one pivot) → one nudge `write()` onto the lane's socket, sequenced after commit. Restart with parked lanes: state restored from records alone (NF-S13 shape).
- **③ known-A / RAISE-ONLY, guardrail-adjacent portion:** the registry marks configured known-A conditions; the rendered form constrains the `gate_category` affordance to `[floor, A]`, and the in-loop validator applies the monotonic-toward-A MAX with `gate_category_raised = true` recorded whenever the floor wins — never delivered as B, no A→B path. Fixture in §3. The full ③ fixture (full §J2 map + detector breadth) stays a Step-1-build-wide owed item.

### D-8 Crash-fixture mechanics (Q-D) + I-PH discipline

**SELECTED: internal named crash-points.** The engine calls `crashpoint("<name>")` at every syscall boundary of the commit path; in normal operation it is a no-op; when the conductor is started with a test-only env (`FRANK_TEST_CRASHPOINT=<name>[:<nth>]`), the n-th hit issues a real `SIGKILL` to itself (never a graceful exit). The crash-point name registry — `post_intake_fsync, pre_record_fsync, post_record_fsync, pre_rename, post_rename, pre_dir_fsync, post_dir_fsync, pre_redo_fsync, post_redo_fsync, pre_projection_write, post_projection_write, pre_delivery_write, post_delivery_write, pre_outcome_reply` — doubles as the mutation-class syscall enumeration S2's full F11 sweep reuses.
- Determinism: exact, per-point, per-hit-count; `go test` spawns the binary, sets the env, drives a scripted seat, asserts the post-recovery store byte-state.
- **I-PH:** crash-point names + the env var are conductor-internal; they never appear in any seat-deliverable surface. More broadly, every seat-facing string (bounces, errors, tool descriptions, projections, delivery payloads) flows through the one formatter whose inputs are field names/reasons only — canonical store/config/outbox/operator-channel paths are not in its vocabulary. The P1 fixture greps **all fixture outputs** for the store-root path family.
- **REJECTED: external syscall-boundary injection** (dtrace/ptrace/strace supervisors) — macOS SIP blocks dtrace, ptrace is non-portable, interposition ordering is nondeterministic; revisit only if S2's exhaustive sweep needs it.

### D-9 Recovery = dumb replay (guide-confirmed line, verbatim scope)

On every start, before any channel opens (**R-3 rebuild-before-open**): (1) staging cleanup — delete torn staging files (never committed, routine); (2) projection rebuild — replay the redo journal idempotently; every committed record's INDEX row / rendered `.md` / mailbox entries exist; canonical wins; INDEX corrections appended, never rewritten; (3) binding-table restore; (4) re-enqueue `intake − outcomes` in arrival order (the F9 property, run whole); (5) **derived-work completion (r2)** — commit missing park/outbox follow-ups for committed gate-bearing records, idempotent by `gate_record_ref` (D-7); (6) wake re-issue for committed-but-undelivered deliveries and parked lanes per records. Only then accept `submit`. **Not built (S1):** genesis validation, canonical-checksum quarantine disposition, GC/segment rotation, the reified phase-0→4 machine.

### D-10 S1-minimal dissolved-linter replay (R1)

Corpus of record: the v2.8.8 fixture matrix (244 fixtures, checker-verified all-PASS — E2, reconciled audits). Harness = a `go test` that classifies every **fail**-fixture by its v2.8.8 check class (the m-2 §10 map) and, for each MVP-covered class, either:
- **caught** — construct the typed-submission equivalent and assert the MVP validator rejects it with the matching class (driving the same in-process validation functions the loop runs — never a markdown submit path; §8 strictness is preserved: the importer lives in test code only); or
- **obsolete-by-construction** — the failure shape is unexpressible through the typed envelope (fences, bare tokens, detached rows, ambiguous continuations, ROLE/FROM mismatch, address-grammar corruption), with the specific reason recorded; or
- **uncovered-S3** — the check class needs registry surface outside the MVP (scan/scope/fold row-arrays, downgrade ordering, design-review/merge lineage beyond the S1 edges) — named in the emitted report, never silently dropped.

Gate: every MVP-covered fail-fixture lands in `caught` or `obsolete-by-construction`; the report table is a fixture artifact. Organic master-trail failures are NOT gate inputs (reconciled disposition).

### D-11 The A4 double-accept instance (the serialized-loop guarantee, actually exercised)

S1's one-shot authority = **one operator verdict per parked decision**. Fixture: two racing verdict submissions resolving the same parked gate item are enqueued concurrently; because both run to completion on the one commit loop, the second's lineage/one-shot check reads the first's committed verdict ⇒ exactly one `accepted`, the loser `rejected` typed; after `kill -9` + recovery the decision is still resolved exactly once. (Plus C6's replayed-`intake_id` dedupe as the intake-grain cousin.)

## 3. Acceptance criteria draft (fixture-keyed; E2 unless noted)

| id | asserts |
|---|---|
| B1 | **(r2 reword — blocker 4)** no seat-facing tool/API path can create or serve a record except conductor `submit` (registry enumeration + the scripted confused-seat leg); only committed `accepted` records are served/delivered. Direct record-file injection detection/quarantine is S2 checksum-disposition; the S1 direct-file leg is limited to what S1 actually detects (torn/staging files cleaned at recovery, never served). **D5 stated plainly:** a same-uid shell can write a syntactically valid record file S1 will not detect — accepted residual |
| B2 | payload `FROM`/`ROLE` ignored byte-for-byte; unbound/wrong-credential connect ⇒ reject, nothing staged |
| B3 | a required-set/enum/seat-scope failure ⇒ one terminal `rejected` record + bounce naming the field/edge; candidate never delivered |
| B4 | one gate-bearing accepted record ⇒ exactly one store-visible outbox item via one pivot; no drain occurs |
| A1 | forged-FROM legs: payload lie inert; self-declared identity on a fresh socket rejected |
| A2 | pair-seat rendered schema omits `merge-gated` + the whole `grant` field (introspection); hand-crafted MCP call supplying them ⇒ rejected in-loop |
| A3 | free-typed/unknown PARENT rejected; `failing_edge` distinguishes in-flight-recompose vs dead-edge |
| A4 | D-11 whole (racing verdicts; exactly-once; survives restart) |
| C1/C2 | kill -9 mid-commit (F3 pre-pivot: nothing committed, staging cleaned, re-enqueued once) and post-pivot/mid-delivery (F4: projections rebuilt, no re-consumption, wake re-issued) |
| C3 | F9 whole: N enqueued, K outcomes ⇒ exactly N−K re-enqueued in arrival order, zero re-emission |
| C4 | F10: simulated cut between record-fsync and dir-fsync, and between rename and projection writes ⇒ store converges both times (presence=committed) |
| C5 | corrupt/delete INDEX.md, a rendered `.md`, and a mailbox ⇒ rebuilt from canonical; INDEX append-only preserved |
| C6 | replayed/duplicate `intake_id` ⇒ zero double-emission; every outcome references its `intake_id` |
| C7 | **(r2 — blocker 1)** kill -9 after gate-accept, before park/outbox completion ⇒ recovery yields exactly one park record + one outbox item; kill -9 again DURING derived-work completion ⇒ still exactly one of each (idempotence by `gate_record_ref`) |
| R1 | D-10 gate + report artifact |
| P1 | grep of every seat-deliverable surface across ALL fixture outputs for store/config/outbox/operator-channel path families ⇒ zero hits |
| L1 | busy/dead seat at delivery time receives via `project()` on reconnect; a dropped nudge is re-issued at recovery; no parked lane sleeps forever |
| W1 | gated lane parks (receives nothing while parked), wakes on the operator verdict, exactly once; kill -9 while parked ⇒ state restored from records alone |
| F11 | for each S1 mutation class: exactly one canonical record via exactly one rename; crash injected at every registered crash-point leaves fully-committed or not-committed — no second pivot to crash between |
| G | registry enumeration = exactly `{submit, project, read}` per seat; scripted confused-seat "edit the store file at <path>" turn has no tool to express it; ③ leg: a B-pick over a known-A condition ⇒ raised to A + `gate_category_raised` recorded, never delivered as B; **(r2 — blocker 5)** enum legs: one A token, one B token, and an unclassified pick ⇒ `other`→A, all byte-exact against the §J2 set; **(r2 — blocker 2)** `grant` present on the operator/orchestrator MERGE-GATE form, absent from every pair-seat form (introspection both ways) |
| H | **(r2 — blocker 3)** an orchestrator-authority relay carrying no `grant` and no gate field (e.g. `PHASE: PLAN` from an orchestrator-planner-stamped seat) is classified authority-bearing: forced validator-throw on it ⇒ `held` (not `rejected`), and its lineage edges block on failure — the r1 rule would have let it bypass |
| SWEEP | m-7 §16 claim-sweep classes (exclusivity semantics + writability-token family) run over everything seat- or user-facing S1 ships, docs and tool descriptions included (E1 review + grep classes) |

## 4. Owed carries → design homes

1. **Code-layer interface-guardrail enforcement** → D-3 registry construction + D-5 render; proven by G. 2. **I-PH fixture** → D-8 single-formatter discipline; proven by P1. 3. **③ known-A/RAISE-ONLY guardrail-adjacent portion** → D-5 `gate_category` monotonic-toward-A + D-7; proven by G's ③ leg. (⑤ cited as the S4-bound §C4 carry only; O-1..O-3 keep it un-foreclosed.)

## 5. Boundary contract

As dispatched (locked): writes = append-only store records (byte-exact `{accepted, rejected, held}`) + one local outbox item; reads = the m-1 verbs + m-2 envelope; target entity = the committed relay + rebuilt projection; consumers = seat inboxes via `project()` + the local outbox; proof = the fixture table above (E2).

## 6. Open items carried to PLAN (none re-architect)

1. MCP Go SDK per-seat-socket transport capability check — first PLAN task; fallback named in D-3 (invariants stated there).
2. *(r2: removed — blocker 5. The full frozen §J2 set is locked in D-5; PLAN names tests and code placement only.)*
3. DI-2/DI-5 realization record (D-3) — flagged to the m-1 fidelity review as the "fork-2 infra call" proposal (DI-5 itself is Step-2/observe; only DI-2 matters in S1).
4. m-1 §13 carries land as PLAN acceptance detail: TOCTOU-atomic submit (the loop already is), credential lifecycle minimal statement (mint/re-attach/revoke-on-remint), operator-channel isolation invariant.
5. **(r2)** The blocker-2 S1 narrowing (no pair-Planner grant rendering in S1) rides the design-completion SITREP for explicit orchestrator ratification.

## 7. Rejected alternatives (log)


Python/TS/Rust stacks (D-1); external syscall injection (D-8); markdown import as a submit path (violates m-2 §8 — importer is test-only); a separate burn/consumed-marker record for verdicts (two pivots — forbidden by F11/m-7-F1); closing the ODB item schema in S1 (O-3); building genesis "while we're at it" (charter-OUT; recovery is dumb replay by lock); priority-dequeue or fast-path mutations (rejected in the locked GRILL, m-7 :198-199); **(r2)** collapsing gate-accept + park + outbox into one compound pivot (rejected in favor of derived-work completion: F11 names park/wake transition and outbox enqueue as distinct mutation classes, and the compound shape would foreclose S2's independent park/wake transitions — the durable-intent + idempotent-completion mechanism heals the crash window without merging mutation classes).

## 8. r2 fold-log (s1-core.implementer DESIGN-REVIEW `.relays/s1/s1-core-design-r1-review/DESIGN-REVIEW-implementer-20260703-151318.md`, verdict must-revise — all five blockers verified correct by the planner before folding)

1. **Blocker 1 — gate/park/outbox crash window.** Folded: D-7 crash-window healing (gate record = durable intent; recovery derived-work completion, idempotent by `gate_record_ref`, before channels open); D-9 gains step 5; new fixture C7. Rejected the alternative compound-pivot shape (see §7).
2. **Blocker 2 — `grant` rendering vs pair-Planner delegated dispatch.** Folded: D-5 grant row now cites the frozen m-2 :177 rule verbatim and states the S1 narrowing explicitly (no conditional pair-Planner grant rendering in S1; the m-2 :167 pair-dispatch lineage walk lands with the full lineage engine, S3), routed for orchestrator ratification via the design-completion SITREP (§6.5). G gains the introspection-both-ways leg.
3. **Blocker 3 — authority-bearing rule too narrow.** Folded: D-6 now defines the S1 rule as a pessimistic superset of m-2 :76 derived from MVP-visible fields (grant / gate / PHASE / AUTHORITY / stamped-ROLE legs), with the explicit class mapping and the never-narrower guarantee. New fixture H.
4. **Blocker 4 — B1 overclaim.** Folded: B1 reworded to the no-tool-path property S1 actually builds; direct-file injection detection moved explicitly to S2; the S1 direct-file leg limited to torn/staging files; D5 residual stated plainly in the row.
5. **Blocker 5 — `gate_category` subset deferred to PLAN.** Folded: D-5 locks the FULL frozen §J2 default set byte-exact (A 8 / B 4 / `other`→A; `routing_escalation` cited as the owed non-member); open item 6.2 removed; G gains the enum legs. PLAN now names tests and code placement only.
6. **Watch items (non-blocking) folded:** binding-table writes use the stage→fsync→rename primitive (D-3); MCP fallback invariants stated (D-3).

No other section changed; operator decisions (D-1 Go) and the guide-locked constraints (R-1..R-3, O-1..O-3, F9-whole) are untouched. Resent for re-review on the same DESIGN_DOC_ID.
