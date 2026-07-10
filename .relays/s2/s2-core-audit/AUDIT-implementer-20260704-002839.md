## AUDIT — s2-core.implementer independent audit: S1 build surface mapped to S2 mandate

ROLE: Implementer
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: s2-core-audit-implementer
PARENT_DISPATCH_ID: s2-core-audit
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s2-core.implementer
TO: s2.orchestrator-planner
CC: operator
SUBJECT: AUDIT — S2 engine-thickening surface: recovery/FIFO/GC/genesis/owed projection

Current artifact: `.relays/s2/s2-core-audit/AUDIT-orchestrator-planner-20260704-001214.md`
Scope: read-only audit only. No source, test, sprint-doc, locked-spec, PLAN, IMPL, merge, or live-verify work performed.

## Summary verdict

PRIMARY_BUCKET: still-open
still-open: All four S2 IN items need implementation work, but S1 has reusable foundations for three of them: atomic record commit, checksum/redo/projection rebuild, intake minus outcomes, gate/held derived-work scan, and crashpoint/SIGKILL harness.
already-closed: S1 already proves the thin conductor spine and several primitive properties at E2; those should be promoted and hardened, not rebuilt.
product-overlapped: The S1 gate/held derived outbox scan overlaps the owed projection and should become an instance of the owed-item projection rather than remain a second parallel mechanism.
recommended-next: Proceed to DESIGN/PLAN after reconciling this audit with the planner audit; m-1 fidelity must review store/record-shape touches before any delegated implementation dispatch.

## Spec-to-exit-gate map

1. Recovery phases 0-4:
- S2 exit gate: crash at each recovery phase boundary, correct resumption, no lost/double intake or delivery (`docs/sprints/2026-07-03-s2-slice-2/ROADMAP.md:54`).
- Locked text: ARCHITECTURE requires startup phases 0-4 and no authority consumption until recovery completes (`../master/ARCHITECTURE.md:427-429`); m-7 spells the phases: genesis validation, scan/quarantine, projection rebuild, runtime table restore, FIFO refill, then open (`the m-7 conductor-core design-of-record (2026-07-01) :87-95`).
- Current code: `internal/recover/recover.go:18-53` implements S1 dumb replay only: clean staging, rebuild projections, reopen bindings, replay unconsumed intake through a callback, run `gate.Complete`. It has no explicit phase enum/state machine, no genesis validation, no canonical checksum quarantine/held incident, no read-only diagnostic mode, and no phase-boundary crash fixtures.

2. Durable FIFO:
- S2 exit gate: order plus exactly-once preserved across crash/restart (`docs/sprints/2026-07-03-s2-slice-2/ROADMAP.md:55`).
- Locked text: m-7's honest formulation is append+fsync before FIFO, ack only with typed outcome, clear-on-pop atomic with outcome commit, recovery re-enqueues `intake - outcomes` in arrival order; "At-least-once intake, exactly-once effect, zero stale re-emission" (`the m-7 conductor-core design-of-record (2026-07-01) :56-58`).
- Current code: `internal/intake/journal.go:41-72` appends/fsyncs, assigns `intake_id`, dedupes content hash; `internal/intake/journal.go:100-127` computes unconsumed commands from committed outcomes; `internal/engine/loop.go:50-97` processes one loop job at a time. Fragility: the live server calls `journal.Append` inside each authenticated connection handler (`cmd/frank/main.go:87-95`), and the socket server runs each accepted connection in its own goroutine (`internal/channel/server.go:72-88`, `internal/channel/server.go:156-176`). That is not the locked dedicated single intake-writer task (`m-7 design:52`) and should be hardened in S2 for full ordering/concurrency claims.

3. GC/genesis:
- S2 exit gate: genesis idempotent, GC never drops a live record, crash mid-GC recovers (`docs/sprints/2026-07-03-s2-slice-2/ROADMAP.md:56`).
- Locked text: m-7 genesis is a first canonical record with schema version, config digest, address-space seed, and creation stamp; canonical records are retained forever; GC only compacts derived artifacts and drained intake/redo journal segments whose entries all have outcomes (`the m-7 conductor-core design-of-record (2026-07-01) :134-137`).
- Current code: `internal/store/store.go:38-53` creates directories only; there is no genesis record, config digest pin, address-space seed, GC marker/plan, journal segment rotation, drained-segment proof, or crash-safe compaction operation. `internal/store/store.go:100-120` reads all canonical records; there is no delete path for canonical records, which matches the retain-everything posture but does not itself implement GC.

4. Owed-item projection:
- S2 exit gate: `open = owed-record with no disposition-record`, `OI-S1-F11-SWEEP` surfaces as open, disposition drops it from the open set, materialize-first honored (`docs/sprints/2026-07-03-s2-slice-2/ROADMAP.md:57-58`).
- Locked/authorizing text: S2 roadmap places owed `record_kind` and store/API fidelity under m-1 authority (`docs/sprints/2026-07-03-s2-slice-2/ROADMAP.md:26-27`); m-1 owns the on-disk typed-envelope/store shape and `submit`/`project`/`read` API (`the m-1 trust/identity design-of-record (2026-06-28) :122-145`).
- Current code: no `record_kind` field or owed/disposition record exists; static search found only the S1 gate/held outbox envelope (`internal/gate/derived.go:11-18`). `store.Project` currently returns mailbox relay IDs for a seat (`internal/store/store.go:131-152`), not an owed open-set projection. S1 does provide the generalized seed: `gate.Complete` scans accepted gate and held records and creates derived outbox items (`internal/gate/derived.go:21-45`, `internal/gate/derived.go:77-115`), with tests for gate, partial-state convergence, mirror state, and held-derived items (`internal/gate/derived_test.go:14-111`).

5. F9/F11 re-run under new recovery:
- S2 exit gate: full class-by-point sweep discharges `OI-S1-F11-SWEEP` (`docs/sprints/2026-07-03-s2-slice-2/ROADMAP.md:58`).
- Locked text: m-7 fixture F9 and F11 require durable FIFO recovery and one-pivot-per-mutation across mutation classes including genesis/config-change/GC marker (`the m-7 conductor-core design-of-record (2026-07-01) :170-172`).
- Current code: S1 harness is reusable. `test/fixtures/f11_test.go:26-63` pins live crashpoint registry names; `test/fixtures/f11_test.go:65-112` drives real child-process SIGKILL for representative S1 mutations; `test/fixtures/f11_test.go:114-149` covers F9 whole. S1 explicitly left full cross-product breadth as the owed S2 item (`docs/sprints/2026-07-03-s1-slice-1/RECONCILE.md:143-161`).

6. No S1 regression:
- S2 exit gate: no regression of serialized-loop kill, crash atomicity, I-PH, enum byte-exactness, guardrail (`docs/sprints/2026-07-03-s2-slice-2/ROADMAP.md:59`).
- Current proof: `go test -count=1 ./...` passed this audit run (E2). The current code surface after `s1-close` has no source/test diff: `git diff --name-only s1-close..HEAD -- internal cmd test go.mod README.md` returned zero paths.

Spec gaps: none found in the S2 exit-gate lines. The open questions are implementation-shape/fidelity questions, not locked-spec gaps.

## Four S2 IN item bucket verdicts

### Recovery phases 0-4
PRIMARY_BUCKET: still-open
still-open: The explicit phase state machine is absent. S1 recovery is a flat function with staging cleanup, projection rebuild, binding restore, intake replay, and derived-work completion (`internal/recover/recover.go:18-53`), but no genesis validation, committed-record checksum quarantine, phase-boundary state, read-only diagnostics, or "open" barrier object.
already-closed: Staging cleanup, projection rebuild, binding reopen, intake replay, and derived-work completion exist and map to embryonic phases 1/2/3/3.5 (`internal/recover/recover.go:18-53`; `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md:130`).
product-overlapped: None requiring reroute; all surfaces are m-7 engine-owned while consuming m-1 store semantics.
recommended-next: Promote `recover.RunWithProcessor` into an explicit phase machine with phase-boundary crash fixtures, preserving S1 order and adding genesis/quarantine/open gating.

### Durable FIFO
PRIMARY_BUCKET: still-open
still-open: The intake journal and loop are present, but the locked "single intake-writer task" is not literal in the live server; authenticated handlers call `journal.Append` directly (`cmd/frank/main.go:87-95`) while server connections run concurrently (`internal/channel/server.go:72-88`, `internal/channel/server.go:156-176`). S2 needs a dedicated intake writer or equivalent serialized journal append path before claiming full ordering under concurrent seats.
already-closed: Append+fsync, content-hash retry dedupe, `intake_id`, unconsumed calculation, and one-job-at-a-time commit loop exist (`internal/intake/journal.go:41-72`, `internal/intake/journal.go:100-127`, `internal/engine/loop.go:50-97`).
product-overlapped: None.
recommended-next: Reuse the existing journal format and engine loop but insert the missing single-writer journal owner; test concurrent handlers, crash/restart, and content-hash retry in arrival order.

### GC/genesis
PRIMARY_BUCKET: still-open
still-open: No genesis record, config digest pinning, quarantine directory/incident flow, GC marker, journal segmentation, or derived-artifact compaction exists. Current store open creates directories only (`internal/store/store.go:38-53`).
already-closed: Canonical records are checksummed (`internal/record/record.go:35-63`) and there is no canonical delete path in current code, which is compatible with "canonical records never GC'd"; redo/projection rebuilding exists (`internal/store/projections.go:18-70`).
product-overlapped: None.
recommended-next: Implement genesis as the first canonical record, add phase-0 validation, add quarantine for checksum failures, and scope GC strictly to derived artifacts and drained journal segments.

### Owed-item projection
PRIMARY_BUCKET: still-open
still-open: There is no owed `record_kind`, disposition record, open-set projection, or `project()` over owed items. Static search found no code-level `record_kind`/owed implementation outside docs; the only code shape is the S1 gate/held outbox envelope (`internal/gate/derived.go:11-18`).
already-closed: The S1 gate/held derived-work scanner is the correct seed: it scans canonical records and materializes derived artifacts idempotently (`internal/gate/derived.go:21-45`, `internal/gate/derived.go:77-115`), and S1 has tests for convergence and held-derived visibility (`internal/gate/derived_test.go:14-111`).
product-overlapped: The S1 ODB/outbox item is adjacent but not sufficient; it should be generalized into the owed projection per the S1 guide advisory (`docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md:189`) rather than duplicated.
recommended-next: Design an m-1-fidelity-reviewed owed record/disposition shape, materialize `OI-S1-F11-SWEEP` as the first open owed item, and make disposition records remove it from the open projection without deleting canonical records.

## m-1 store-touch enumeration for fidelity review

- Add or extend canonical record typing for owed items: exact `record_kind` home, enum/value, and typed envelope placement are m-1-owned.
- Define owed record fields for at least owner, source, target surface, disposition path, created timestamp, schema version, and source reference to `OI-S1-F11-SWEEP`.
- Define disposition-record shape and linkage: how a disposition references an owed record; whether disposition has its own `record_kind`; accepted terminal state; whether disposition is a normal canonical record.
- Define `open = owed-record with no disposition-record` projection mechanics over immutable canonical records; no mutation/delete of the owed record.
- Extend or add `project()` semantics for the owed open set while preserving existing mailbox `project(seat)` semantics (`internal/store/store.go:131-152`) and m-1 API boundaries (`m-1 design:130-131`).
- Add genesis record shape: first canonical record fields, config digest, address-space seed, creation stamp, and idempotent initialization rules.
- Add quarantine surface: directory/location, held incident record shape, relationship to checksum failures, and whether quarantined canonical records remain discoverable by `read`.
- Define GC target set: derived projections, old rendered views, and drained intake/redo journal segments only; canonical records never deleted (`m-7 design:136-137`).
- Define journal segment rotation metadata and proof that all segment entries have outcomes before compaction.
- Preserve m-1 locked layout/API fidelity: `submit`, `project`, `read`, append-only immutable records, schema_version, markdown as projection/view, `INDEX.md` layout unchanged (`m-1 design:122-145`, `m-7 design:82-83`).

## Claim-boundary probes

1. Durable FIFO wording:
- Honest claim is not "exactly-once intake"; locked text says "At-least-once intake, exactly-once effect, zero stale re-emission" (`m-7 design:56-58`). S2 docs/tests should use that wording. The S2 roadmap's "exactly-once preserved" line should be read as exactly-once effect, not exactly-once queue admission.

2. GC target set:
- Canonical records are never GC'd. GC compacts only derived artifacts and drained intake/redo journal segments whose entries all have outcomes; journal rotation is size-based and operator-configured (`m-7 design:136-137`). "GC never drops a live record" is true but weaker than the locked target set; S2 should state "no canonical record deletion."

3. Materialize-first honesty:
- Owed projection guards recorded owed items only. The S2 roadmap already states it does not make an unrecorded observation impossible to miss (`docs/sprints/2026-07-03-s2-slice-2/ROADMAP.md:65-67`). S2 should preserve that line beside any projection/coverage claim.

4. Confusion-resistance:
- D5 remains accepted. ARCHITECTURE says tool-mediated guardrail removes affordance, not access; a same-uid shell-bearing seat can still write store files directly (`../master/ARCHITECTURE.md:450-463`). S2 should not turn GC/genesis/owed projection into a stronger isolation claim.

## Duplicate/already-built gate

- Reuse/promote: `record.Seal` and `record.Verify` checksum format (`internal/record/record.go:35-63`); `fsio.WriteFileAtomic` rename pivot and crashpoints (`internal/fsio/fsio.go:10-53`); `store.Commit` redo-before-record and projection application (`internal/store/store.go:55-91`, `internal/store/store.go:181-203`, `internal/store/projections.go:18-70`); `intake.Unconsumed` (`internal/intake/journal.go:100-127`); `gate.Complete` scanner (`internal/gate/derived.go:21-45`); crashpoint registry and child SIGKILL harness (`test/fixtures/f11_test.go:26-149`).
- Do not rebuild: the S1 crashpoint/SIGKILL harness is explicitly the machinery for `OI-S1-F11-SWEEP` full class-by-point sweep; extend it to all phase/mutation classes instead of creating a separate harness (`docs/sprints/2026-07-03-s1-slice-1/RECONCILE.md:160-161`).
- Not already closed: no full recovery phase machine; no m-7 exact single intake-writer task; no genesis/quarantine/GC; no owed/disposition typed record/open projection.

## Evidence and verification

- E1 source/spec audit: cited inline above.
- E2 suite: `go test -count=1 ./...` passed across cmd, internal packages, `test/fixtures`, and `test/replay`.
- E2 targeted race pass: `go test -race -count=1 ./internal/intake ./internal/channel ./internal/engine ./internal/recover ./test/fixtures` passed.
- E1 baseline check: current `HEAD` is `cc85049`; `s1-close` is `f0dcb85`; `git diff --name-only s1-close..HEAD -- internal cmd test go.mod README.md` returned zero paths, so source/test code audited here is the S1 code.
- Relay lint: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s2/s2-core-audit/AUDIT-implementer-20260704-002839.md` returned OK.
- Relay-root lint context: `python3 ~/.codex/skills/tools/relay-lint.py --relay-root .relays/s2 .relays/s2/s2-core-audit/AUDIT-implementer-20260704-002839.md` reports the known lint-exempt `INDEX.md` header noise and then `OK .relays/s2/s2-core-audit/AUDIT-implementer-20260704-002839.md`.

## Escalations / open questions

- No locked-spec gap found.
- m-1 fidelity review is required before dispatch for owed `record_kind`, disposition shape, genesis/quarantine record shape, store layout/API changes, and GC/journal segment semantics.
- The S2 plan should make the intake-writer hardening explicit; otherwise the code could overclaim durable FIFO ordering under concurrent authenticated connections.

ACTIONS_GIT_REF: no source/test/docs edits claimed; wrote audit relay `.relays/s2/s2-core-audit/AUDIT-implementer-20260704-002839.md` and appended `.relays/s2/INDEX.md` only; `.relays/` is gitignored operational substrate; final tracked `git status --short` = none — clean tree.
FINAL_GIT_STATUS_SHORT: none — clean tree
