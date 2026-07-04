# S1 Slice-1 Implementation Plan — the thin end-to-end conductor relay

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:executing-plans (single Implementer seat; implementation starts ONLY on a live `DISPATCH IMPL` per the dispatch conditions). Steps use checkbox (`- [ ]`) syntax for tracking.

**PLAN_LOCK_ID:** `s1-slice-1-plan` · **DESIGN_LOCK_ID:** `s1-slice-1-design` (r5; pair-approved at r2 `…-152445`, r3 = guide should-fix folds, r4 = plan-review envelope fix, r5 = m-1 F-M1-1 credential-lifecycle + guide nit) · **Owner:** s1-core · **Date:** 2026-07-03 · **Rev:** r3 (r2 = plan-review blockers; r3 = F-M1-1 lifecycle invariant + fixture in Task 5)
**Goal:** one working conductor relay in Go — `mint → connect → render → submit → stamp → validate → lineage → append(crash-atomic) → project → deliver → gate-outbox` — with the S1-scoped hardened exit gate green (E2).
**Architecture:** one `frank` process; per-seat Unix-socket MCP channels (3 tools); single intake-writer goroutine → durable intake journal → single commit-loop goroutine → Package-A rename-pivot store with derived projections; dumb-replay recovery before open; derived-work completion for gate/held follow-ups. Design doc = the spec of record; this plan cites it as `D-<n>` (docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md).
**Tech stack:** Go ≥1.22 (operator decision D-1), official MCP Go SDK (github.com/modelcontextprotocol/go-sdk) with the D-3 fallback if the capability check fails, stdlib elsewhere. `go test ./...` is the E2 umbrella.

## Global constraints (every task inherits; violations are review blockers)

- Terminal enum byte-exact `{accepted, rejected, held}` — `bounced`/`submitted` never appear as value tokens anywhere (code, fixtures, docs).
- One pivot per mutation class (F11): every mutation commits exactly one canonical record via exactly one `rename()`; fsync-before-rename, dir-fsync-after; presence = committed.
- Outcome records reference `intake_id` (R-2); canonical records carry a `checksum` field (R-1) from the very first commit.
- Rebuild-before-open (R-3): no channel accepts `submit` until staging cleanup + projection rebuild + re-enqueue + derived-work completion finish.
- I-PH: no canonical store/config/outbox/operator-channel path in ANY seat-deliverable string; all seat-facing text flows through `internal/bounce.Formatter` (D-8).
- Claim honesty (m-7 §16 sweep classes): D5 residual beside every exclusivity-shaped claim in seat-/user-facing strings, comments, docs; only the serialized-loop kill (D-11) + constrained-grammar R2 presented as operationally live.
- Store layout + record/journal formats are the D-4 shapes exactly; `gate_category` enum = the full frozen §J2 set byte-exact (D-5; `routing_escalation` NOT a member).
- **S1 narrowing (ratified, 3 conditions — dispatch `s1-core-plan` :37-41):** `grant` renders on operator/orchestrator forms only; no conditional pair-Planner grant rendering in S1 (S3 landing stated in any grant-describing surface); no schema/format decision may foreclose the S3 conditional-render landing (render is registry-data-driven — additive later, D-3/m-7 §8.2).
- Scope fence: nothing from the ROADMAP scope-OUT list; ⑤ cited only as the S4-bound carry; no edit ever under ../master or ../extracted.
- Commits: small, per green step, on `main` (greenfield repo, per dispatch TARGET_BRANCH), message prefix `s1 IMPL:`.

## File structure (the decomposition of record; SCOPE_DIFF enumerates exactly these)

```
go.mod, go.sum                          # module github.com/jackli/frank (Task 0)
cmd/frank/main.go                       # entrypoint: config load, recovery, channels, loop (Task 10)
internal/crashpoint/crashpoint.go       # named crash-point registry (Task 1)
internal/fsio/fsio.go                   # WriteFileAtomic + dir-fsync + append-fsync primitives (Task 1)
internal/record/record.go               # typed envelope, canonical JSON, checksum (Task 1)
internal/store/store.go                 # layout, pivot commit, redo journal (Task 2)
internal/store/projections.go           # INDEX.md, rendered .md, mailboxes (Task 2)
internal/intake/journal.go              # intake journal + intake-writer (Task 3)
internal/engine/loop.go                 # FIFO + single commit-loop + fault disposition (Task 3, 6, 7, 9)
internal/fieldspec/fieldspec.go         # registry load, predicate eval, render, validate (Task 4)
internal/fieldspec/registry.json        # the MVP FieldSpec data incl. §J2 set (Task 4)
internal/seat/binding.go                # mint, persisted binding table, credential check (Task 5)
internal/channel/server.go              # per-seat socket MCP server, 3-tool registry, push (Task 0 spike → Task 5)
internal/lineage/lineage.go             # parent/authority/one-shot checks (Task 6)
internal/bounce/formatter.go            # the single seat-facing text formatter (I-PH) (Task 6)
internal/gate/derived.go                # park/outbox/held derived-work + completion scan (Task 9)
internal/recover/recover.go             # dumb-replay recovery steps 1-6 (Task 10)
test/fixtures/…_test.go                 # fixture suite (each task adds its ids)
test/replay/replay_test.go + report     # R1 harness (Task 11)
test/seatproc/testseat.go               # scripted test seat (Task 5)
README.md                               # repo-root honesty surface — in-scope ONLY for Task 12 (claim-honest framing; SWEEP-checked)
docs/sprints/2026-07-03-s1-slice-1/plans/s1-slice-1-plan.md   # this plan (pair-Planner-committed)
```

**Scope-fence note (plan-review blocker 1):** `README.md` is a NEW root doc, not a Go source/test/build file — its inclusion under the dispatch fence is flagged to `s1.orchestrator-planner` for explicit confirmation in the plan-completion report (a `TO`-addressed ask, before any dispatch); if the orchestrator reads the fence as excluding it, README content moves under docs/ per their ruling and this block updates before SCOPE_DIFF runs.

Fixture-id → task map: Task 1 {F10} · 2 {C5, F2-shape} · 3 {C3/F9, C6} · 4 {B3, A2-enum, G-enum/③} · 5 {B2, A1, G(i)(ii)} · 6 {A3, H, F5-shape} · 7 {B1, C1, C2} · 8 {L1} · 9 {B4, W1, C7, A4, H-r3} · 10 {full C-matrix re-run, F11} · 11 {R1} · 12 {P1, SWEEP, G-final}.

---

### Task 0: Module bootstrap + MCP capability check (the dispatch-mandated FIRST task)

**Files:** Create `go.mod`, `internal/channel/server.go` (spike form), `internal/channel/capability_test.go`.
**Interfaces produced:** `channel.Serve(sockPath string, tools channel.ToolSet) (*channel.Server, error)`; `(*Server).Push(frame []byte) error` (server-initiated nudge); `channel.ToolSet{Submit, Project, Read func(ctx, json.RawMessage) (json.RawMessage, error)}`.

- [x] **Step 1: failing capability test** — `capability_test.go`: start a server on a temp Unix socket with a 3-tool set; from a test client: (a) `tools/list` returns exactly `["submit","project","read"]` — no more, no fewer; (b) call `submit`, get a response; (c) **server pushes an unsolicited nudge frame and the client receives it while idle** (D-3 r3 invariant); (d) two sockets = two isolated seats. Run: `go test ./internal/channel/ -run TestCapability -v` — expect FAIL (nothing exists).
- [x] **Step 2: implement against the official MCP Go SDK.** If (c) is unimplementable over the SDK's socket transport, switch to the D-3 fallback (minimal self-hosted MCP framing: JSON-RPC 2.0 over the socket, `tools/list` + `tools/call` + one server→client `notifications/nudge` method) — the fallback preserves per-seat channels, stamped identity, 3-tool registry, push (all four D-3 invariants). Record which path shipped in the task-close commit message.
- [x] **Step 3: green + commit** — `go test ./internal/channel/ -v` PASS; `git commit -m "s1 IMPL: channel transport + capability check (SDK|fallback: <which>)"`.

### Task 1: Crash-points, atomic file primitives, canonical record (R-1)

**Files:** Create `internal/crashpoint/crashpoint.go`, `internal/fsio/fsio.go`, `internal/record/record.go` + `_test.go` each.
**Interfaces produced:** `crashpoint.Hit(name string)` (no-op unless env `FRANK_TEST_CRASHPOINT=<name>[:<nth>]` at process start ⇒ nth hit = `syscall.Kill(os.Getpid(), SIGKILL)`); `crashpoint.Names()` (the D-8 registry list, exported for the F11 sweep); `fsio.WriteFileAtomic(dir, name string, data []byte) error` (stage in `<root>/staging/` → fsync → rename → dir-fsync, `crashpoint.Hit` at pre/post of each syscall); `fsio.AppendFsync(f *os.File, line []byte) error`; `record.Record{Envelope, Headers, Body, XFields, Checksum}`; `record.Seal(r) ([]byte, error)` (canonical JSON, sha256 over all-but-checksum); `record.Verify(data []byte) (Record, error)`.

- [x] **Step 1: failing tests** — checksum roundtrip (`Seal`→`Verify` OK; flip one byte ⇒ `Verify` errors); crashpoint no-op without env; `WriteFileAtomic` visible-iff-completed. Run `go test ./internal/... -v` — FAIL.
- [x] **Step 2: implement the three packages.** The atomic-write core:

```go
func WriteFileAtomic(root, rel string, data []byte) error {
    stage := filepath.Join(root, "staging", rel+".tmp")
    // write stage
    crashpoint.Hit("pre_record_fsync")
    // f.Sync()
    crashpoint.Hit("post_record_fsync"); crashpoint.Hit("pre_rename")
    // os.Rename(stage, filepath.Join(root, rel))
    crashpoint.Hit("post_rename"); crashpoint.Hit("pre_dir_fsync")
    // open parent dir, dirf.Sync()
    crashpoint.Hit("post_dir_fsync")
    return nil // (error paths return wrapped errors; no path strings in returned messages destined for seats)
}
```

- [x] **Step 3: F10 fixture** — `test/fixtures/f10_test.go`: helper `runConductorCrash(t, crashEnv, script)` spawns a child `go test`-built binary (or the Task-2 store CLI harness), sets `FRANK_TEST_CRASHPOINT`, drives one write, waits for SIGKILL death, re-runs recovery, asserts: crash at `pre_rename`/`post_record_fsync` ⇒ file absent + staging cleaned; crash at `post_rename`/`pre_dir_fsync` ⇒ file present + valid. Green, then commit `s1 IMPL: crashpoints + atomic fsio + sealed records (F10)`.

### Task 2: Store — pivot commit, redo journal, projections (D-4)

**Files:** Create `internal/store/store.go`, `internal/store/projections.go`, tests.
**Interfaces produced:** `store.Open(root string) (*Store, error)`; `(*Store).Commit(rec record.Record, redo []store.Intent) (relayID string, err error)` — ONE pivot: seal → stage redo intents (fsync, `crashpoint.Hit("pre_redo_fsync"/"post_redo_fsync")`) → `WriteFileAtomic` into `records/` → apply intents (`crashpoint.Hit("pre_projection_write"/"post_projection_write")`); `store.Intent{Kind: index|render|mailbox|outbox, Payload}`; `(*Store).RebuildProjections() error` (idempotent redo replay; canonical wins; INDEX append-only — a wrong row superseded by an appended correction, never rewritten); `(*Store).Records() iter`; `(*Store).AppendIndex(row string) error`.

- [x] **Step 1: failing tests** — commit-shape (one commit ⇒ exactly one `records/<relay_id>.json` + one INDEX row + one rendered `.md` under `projections/relays/<DISPATCH_ID>/` + a mailbox line per recipient, all checksums valid — the F2 shape); C5 (delete/corrupt INDEX.md, a rendered `.md`, a mailbox ⇒ `RebuildProjections` restores all from canonical; append-only INDEX preserved). Run — FAIL.
- [x] **Step 2: implement**; INDEX row format mirrors the v2.8.8 columns (REUSE-AS-SPEC'D); rendered `.md` = view-only serialization of the record.
- [x] **Step 3: green + commit** `s1 IMPL: store pivot + redo projections (F2-shape, C5)`.

### Task 3: Intake journal, intake-writer, FIFO, loop skeleton (D-2, §2.2 full semantics)

**Files:** Create `internal/intake/journal.go`, `internal/engine/loop.go`, tests.
**Interfaces produced:** `intake.Journal.Append(cmd intake.Cmd) (intakeID string, err error)` (assigns id, content-hash key, `AppendFsync`, `crashpoint.Hit("post_intake_fsync")`); `intake.Cmd{Seat, Verb, Payload, ContentHash}`; `engine.Loop{In chan engine.Job, …}` — `Job{Cmd, ReplyCh chan engine.Outcome}`; `engine.Outcome{State string /*accepted|rejected|held*/, RelayID, IntakeID, Reason}`; `(*Loop).Run(ctx)` — ONE goroutine, each job start-to-finish; `intake.Unconsumed(store) []Cmd` (intake − outcomes, arrival order).

- [x] **Step 1: failing tests** — F9 (spawn child, enqueue 5, crash via `FRANK_TEST_CRASHPOINT` after 2 outcomes commit, recover ⇒ exactly the 3 outcome-less re-enqueued in arrival order, zero re-emission — **F9 run whole, named F9**); C6 (duplicate content-hash retry ⇒ deduped; replayed intake-id with existing outcome ⇒ nothing emitted; every outcome references its `intake_id`). Run — FAIL.
- [x] **Step 2: implement** (ack only on typed outcome: handler's ReplyCh receives after the pivot).
- [x] **Step 3: green + commit** `s1 IMPL: durable intake + serialized loop (F9, C6)`.

### Task 4: FieldSpec registry, predicate evaluator, render, form-validate (D-5, D-6 form half)

**Files:** Create `internal/fieldspec/fieldspec.go`, `internal/fieldspec/registry.json`, tests.
**Interfaces produced:** `fieldspec.Load(path string) (*Registry, error)` (load-once at startup; the path never appears in seat-facing strings); `(*Registry).Render(seat, phase, tier string) (Form, digest string)` (visible fields + per-seat options as JSON schema; system fields not shown; forbidden options ABSENT); `(*Registry).Validate(cand record.Record, seat SeatMeta, formDigest string) []Violation` (required-set via predicate eval, enum membership, seat-scope, monotonic floors incl. `gate_category` `[floor,A]` RAISE-only + `gate_category_raised`, stale digest ⇒ `re-render` violation); `Violation{Field, Class, Reason}` — Reason text contains field names only, never paths.
**registry.json carries exactly the D-5 table:** envelope system set; PHASE (11 tokens), AUTHORITY (seat-scoped; pair forms omit `merge-gated`), CEREMONY_TIER, EVIDENCE_TARGET, HUMAN_GATE_REQUIRED (monotonic), gate_category (**A: merge_to_protected, irreversible_write, residual_risk_acceptance, live_verify_skip, ceremony_downgrade, authz_security, product_semantics, scope_expansion; B: merge_feature_to_feature, routing, sequencing, scope_within_bounds; other→A hardcoded**), gate_category_raised, grant (`{dispatch-impl, dispatch-merge}`, **operator/orchestrator forms only**, dispatch-merge phase-bound MERGE-GATE), delivery_state, failing_edge, SUBJECT, body, X-* (consumers:[], lineage_role:none).

- [x] **Step 1: failing tests** — B3 shape (missing SUBJECT / bad PHASE token / out-of-scope AUTHORITY ⇒ named violations); A2-enum (pair form renders no `merge-gated` and NO `grant` field at all — schema introspection; operator MERGE-GATE form renders both); G-enum legs (one A pick OK; one B pick OK; unknown token ⇒ classified A via other→A; below-floor gate pick ⇒ floor wins); ③ leg (registry known-A condition + agent B-pick ⇒ raised to A, `gate_category_raised=true`); stale form digest ⇒ `re-render` reason. Run — FAIL.
- [x] **Step 2: implement** (predicate atoms: `phase_in`, `authority_in`, `seat_is`, `field:<id> ==|present`; anything else in registry data ⇒ load error — closed vocabulary).
- [x] **Step 3: green + commit** `s1 IMPL: FieldSpec registry + render + validate (B3, A2-enum, G-enum, ③)`.

### Task 5: Identity — mint, binding table, connect, stamp; the 3-tool guardrail live (D-3)

**Files:** Create `internal/seat/binding.go`, wire `internal/channel/server.go` to real seats, `test/seatproc/testseat.go`, tests.
**Interfaces produced:** `seat.Mint(name string) (Cred, error)` (conductor-internal; ≥128-bit credential; persists `binding/seats.json` via `fsio.WriteFileAtomic`; **lifecycle invariant, design r5 / F-M1-1 shape (b): S1 has no remint/recycle — `Mint` on an already-bound seat returns a typed `ErrSeatAlreadyBound`, generates NO second credential, and leaves the binding table byte-identical; exactly one credential generation per seat for the life of the store**); `seat.Resolve(cred string) (SeatMeta, bool)`; `SeatMeta{Name, Role, IsOperator bool}`; per-seat socket serving `ToolSet` bound to that seat; **stamping:** envelope FROM/ROLE always overwritten from `SeatMeta` — payload values discarded before validation; operator channel = a mint with `IsOperator=true` whose socket path/credential go only to operator tooling.

- [x] **Step 1: failing tests** — B2/A1 (submit with payload FROM/ROLE set to a victim ⇒ committed record carries stamped values byte-for-byte; connect with unknown credential ⇒ reject, nothing staged, nothing in intake; fresh-socket self-declaration ⇒ reject); G(i) (per-seat `tools/list` = exactly the 3); G(ii) (scripted confused-seat turn "edit the store file directly" — the test seat enumerates its tools and asserts no tool can express a path/write — no tool input schema accepts a filesystem path); restart ⇒ re-attach same credential = same seat, no re-mint; **lifecycle legs (F-M1-1, E2):** `Mint("seat-a")` twice ⇒ second call returns `ErrSeatAlreadyBound`, `binding/seats.json` byte-identical before/after, the original credential still resolves, and a scan of the binding table + runtime state finds exactly one credential for seat-a (no stale/parallel generation exists anywhere); a credential absent from the binding table rejects before staging (shared with B2). Run — FAIL.
- [x] **Step 2: implement**; **Step 3: green + commit** `s1 IMPL: mint/bind/stamp + 3-tool guardrail (B2, A1, G(i)(ii))`.

### Task 6: Lineage, authority classification, fault→held, bounce formatter (D-6)

**Files:** Create `internal/lineage/lineage.go`, `internal/bounce/formatter.go`, extend `internal/engine/loop.go` (validate stage), tests.
**Interfaces produced:** `lineage.AuthorityBearing(cand record.Record, meta seat.SeatMeta) bool` — the D-6 pessimistic superset verbatim: grant present ∨ HGR=yes ∨ gate_category∈A∪{other} ∨ PHASE∈{PLAN,IMPL,REVIEW-FOLD,MERGE-GATE,LIVE-VERIFY} ∨ AUTHORITY∈{implementation,merge-gated,live-verify,fold-in-only} ∨ (ROLE=orchestrator-planner ∧ PHASE∉{SITREP,RECONCILE}); `lineage.Check(cand, store) *lineage.Bounce` (PARENT resolves in accepted graph — parent_picker candidate set; verdict-parents-parked-gate edge; blocking only when AuthorityBearing); `Bounce{Edge, Kind}` Kind∈{parent-unknown-recompose, parent-invalid-dead-edge}; `bounce.Format(v …Violation|Bounce) string` — the ONLY producer of seat-facing text (unit test: its output for every violation class contains no `/` path segment of the store root family); loop fault wrapper: panic/timeout in any trusted check ⇒ authority-bearing ⇒ ONE compound `held` record embedding candidate bytes + `intake_id` (one pivot), non-authority ⇒ `rejected` + bounce; loop keeps serving (next job processes).

- [x] **Step 1: failing tests** — A3 (unknown PARENT ⇒ rejected, Kind=parent-unknown-recompose; PARENT naming a rejected/nonexistent-dead edge ⇒ parent-invalid-dead-edge); H core (a `PHASE: PLAN` orchestrator-planner-stamped candidate with no grant/gate + validator forced to panic (test hook: registry field validator injected via test seam) ⇒ outcome `held` not `rejected`, loop alive — next submit processes); F5-shape (same for a gate-bearing candidate). Run — FAIL.
- [x] **Step 2: implement**; **Step 3: green + commit** `s1 IMPL: lineage + authority superset + fault→held + I-PH formatter (A3, H-core, F5-shape)`.

### Task 7: The submit pipeline end-to-end on the loop (D-2 [C] order)

**Files:** Extend `internal/engine/loop.go` (stamp→validate→lineage→commit→outcome→deliver order), tests.
**Interfaces produced:** the wired `submit` verb: handler → intake → loop → `store.Commit` → outcome to ReplyCh → delivery push. Reject path commits ONE terminal `rejected` evidenced record (same pivot mechanics) + bounce.

- [x] **Step 1: failing tests** — B1 (a spectator seat's `project()` never returns a non-accepted record; the rejected/held records are absent from work-consumer delivery; torn staging file after crash ⇒ cleaned, never served; **held carve-out asserted: the held record IS visible on the operator surface via its Task-9 outbox item — cross-referenced there**); A2 reject leg (a hand-crafted `submit` call from a pair seat supplying `grant: dispatch-impl` or `AUTHORITY: merge-gated` ⇒ rejected by the in-loop authoritative validation — render is advisory, the loop is authoritative); C1 (crash `pre_rename` mid-commit ⇒ recover ⇒ nothing committed, intake re-enqueued once, exactly-once after re-run); C2 (crash `pre_delivery_write` after pivot ⇒ recover ⇒ record committed once, wake re-issued, no duplicate record). Run — FAIL.
- [x] **Step 2: wire**; **Step 3: green + commit** `s1 IMPL: submit pipeline e2e (B1, C1, C2)`.

### Task 8: project/read + delivery push (L1)

**Files:** Extend `internal/channel/server.go` + `internal/store/projections.go` (mailbox read), tests.
**Interfaces produced:** `project` returns the caller's mailbox records (own TO/CC only — scoped by stamped identity); `read(relay_id)` returns `{record, schema_version}` for committed records; delivery = one `Push` per recipient channel, sequenced after commit; busy/disconnected seat ⇒ push skipped/failed silently, mailbox is truth.

- [x] **Step 1: failing tests** — L1 (deliver to a disconnected seat ⇒ reconnect ⇒ `project()` returns it; deliver to a busy seat ⇒ kernel/socket queues, no loss; kill conductor between commit and push ⇒ recovery re-issues the nudge — inbox durable, pipe = nudge); scope leg (seat A's `project` never returns seat B's records; `read` of another seat's relay_id: allowed for committed records — store queries serve canonical immutable records per m-1 :131 — but the mailbox scoping leg still holds for `project`). Run — FAIL.
- [x] **Step 2: implement**; **Step 3: green + commit** `s1 IMPL: project/read + push delivery (L1)`.

### Task 9: Gate → park/wake → outbox + held item + operator verdicts (D-7, O-1..O-3, r3 ①)

**Files:** Create `internal/gate/derived.go`, extend loop + recovery hooks, tests.
**Interfaces produced:** derived-intent classes `{park(source), outbox(source)}` where `source = (source_kind gate|held, source_record_ref relay_id)` — the SAME pair is the idempotence key for every derived-work class (design r4); `gate.Complete(store) error` — the idempotent scan (missing follow-ups committed, keyed by source pair; runs on the loop turn in normal operation, single-threaded pre-open at recovery); ODB item JSON `{item_id, source_kind: gate|held, source_record_ref, seat, gate_category (gate-only; null for held with no gate field), created_ts, schema_version}` (open envelope — NOT closed, no model_name; exactly one source pair per item); park state = store records; wake = verdict commit then one Push; verdict lineage: must parent the parked gate record; one-shot per decision.

- [x] **Step 1: failing tests** — B4 (one gate-bearing accepted record ⇒ exactly one `outbox/<item>.json` via one pivot; no drain code path exists — grep the binary's symbols/source tree for any network egress in `internal/gate`); W1 (parked lane receives nothing while parked; operator verdict via operator channel ⇒ wake push exactly once; kill -9 while parked ⇒ Phase-restore from records alone); C7 (crash after gate-accept before follow-ups ⇒ recovery yields exactly one park + one outbox item; crash again DURING `gate.Complete` ⇒ still exactly one of each); A4 (two racing verdicts for the same parked decision enqueued concurrently ⇒ exactly one accepted, loser typed-rejected; restart ⇒ still resolved once — D-11, the licensed §2.4 instance); H-r3 (a `held` record from Task 6 ⇒ exactly one operator-visible outbox item with `source_kind=held` + `source_record_ref` = the held record's relay_id; crash between held-commit and derivation ⇒ completed at recovery; re-crash during completion ⇒ still exactly one, idempotent by the source pair). Run — FAIL.
- [x] **Step 2: implement**; **Step 3: green + commit** `s1 IMPL: gate/park/wake/outbox + held visibility (B4, W1, C7, A4, H-r3)`.

### Task 10: Recovery assembled + the full crash matrix + F11 sweep (D-9)

**Files:** Create `internal/recover/recover.go`, `cmd/frank/main.go`, `test/fixtures/f11_test.go`, tests.
**Interfaces produced:** `recover.Run(store, intake, bindings) error` executing steps 1–6 in order (staging cleanup → projection rebuild → binding restore → re-enqueue intake−outcomes → derived-work completion → wake re-issue), all before any socket listens; `main.go` = load registry → recover → mint/attach per config → serve.

- [x] **Step 1: failing tests** — rebuild-before-open (a submit raced during recovery ⇒ connection refused/queued, never processed early); F11 (for each S1 mutation class {submit-accept, submit-reject, held, operator-verdict, park, outbox-enqueue}: drive one instance with `FRANK_TEST_CRASHPOINT` at EVERY name in `crashpoint.Names()` — after each crash+recover, store is byte-equivalent to fully-committed or not-committed; assert exactly one rename per mutation by instrumenting `fsio` counters); full C1–C7 matrix re-run against the assembled binary. Run — FAIL.
- [x] **Step 2: assemble**; **Step 3: green + commit** `s1 IMPL: recovery + crash-matrix + F11 sweep`.

### Task 11: R1 — the S1-minimal dissolved-linter replay (D-10)

**Files:** Create `test/replay/replay_test.go`, `test/replay/classmap.go`, output `test/replay/report.md` (generated artifact, committed).
**Interfaces consumed:** `fieldspec.Validate` in-process (importer is TEST-ONLY — never a submit path; m-2 §8 strictness preserved).

- [x] **Step 1: failing test** — `classmap.go` maps every FAIL fixture under `extracted/.../tools/relay-lint-fixtures/` (read-only absolute path in test config, never in seat-facing strings) to its m-2 §10 check class; test asserts every MVP-covered class fixture ends `caught` (typed equivalent constructed, validator rejects with matching class) or `obsolete-by-construction` (shape unexpressible — fences/bare tokens/rows/ROLE-FROM/address grammar — reason recorded), and every other lands `uncovered-S3` in the report, never silently dropped. Run — FAIL.
- [x] **Step 2: implement the classifier + per-fixture typed-equivalents; generate `report.md`** (three-section disposition table).
- [x] **Step 3: green + commit** `s1 IMPL: R1 replay harness + disposition report`.

### Task 12: P1 I-PH grep, SWEEP claim-honesty, G-final, honesty docs

**Files:** Create `test/fixtures/iph_test.go`, `test/fixtures/sweep_test.go`, `README.md` (repo root), tests.

- [x] **Step 1: failing tests** — P1 (run the ENTIRE fixture suite with output capture: every seat-deliverable string produced anywhere — bounces, tool results, rendered projections delivered via project, push frames — grepped for the store-root/config/outbox/operator-socket path families ⇒ zero hits; evidence citations to relay/design files are not hits per NF-S18 rev2 qualifier); SWEEP (grep seat-/user-facing strings + README + tool descriptions for the exclusivity/writability claim classes — `only writer`, `sole`, `no lane can`, `non-lane-writable`, `unbypassable` — each hit must have the D5 residual or governance-surface qualifier within the same sentence/paragraph; `bounced`/`submitted` absent as value tokens anywhere in emitted output). Run — FAIL where README doesn't exist yet.
- [x] **Step 2: write README.md** (what frank is; S1 = provenance + transport, not verified work; "done" = self_reported; tool-mediated confusion-resistance + D5 stated; grant narrowing + S3 landing stated — ratification condition 2). Fix any P1/SWEEP hits.
- [x] **Step 3: green + `go test ./...` full-suite PASS + commit** `s1 IMPL: I-PH + claim-sweep fixtures + honest README (P1, SWEEP)`.

---

### Task 13 (process, pair Planner — NOT Implementer work): external gates, then dispatch

- [ ] **Step 1:** Implementer PLAN-REVIEW of this plan → approve required (findings folded and reissued otherwise).
- [ ] **Step 2:** Orchestrator relays the locked plan to `m-7.planner` (guide) + `master.orchestrator-reviewer` (VP) via operator — carrying: the guide-gate rubric mapping (below), the r2→r3 delta note, advisory sharpening (b) (S2's owed-item projection generalizes C7 — at S2 the scan becomes an instance of it, not a parallel mechanism), and the ratified narrowing statement (condition 1).
- [ ] **Step 3:** Fidelity packet to `m-1.implementer` + `m-2.implementer`: the audits' E1 contract enumerations (planner audit §3 + implementer audit) as the review object + the DI-2 realization record (design D-3) + registry.json (§J2 byte-custody with m-2).
- [ ] **Step 4:** All four approvals present as relays in `.relays/s1/` → pair Planner runs the mechanical `SCOPE_DIFF` over this plan's file list vs the dispatch fence → `all-in` required → only then the delegated `DISPATCH IMPL` token relay (conditions verbatim from dispatch :25). Absent any approval: relay to `s1.orchestrator-planner` and wait.
- [ ] **Step 5 (at S1 close, separate):** REVIEW-FOLD → SITREP to master at exit gate → human merge gate. Merge is never implied by green fixtures.

## Guide-gate rubric mapping (checklist item → where this plan satisfies it)

1. Scope fence → Global constraints + file structure (IN list only; OUT named); escalate-don't-expand in Task-13 step 4 fallback.
2. Contract-fidelity wiring → Task 13 steps 3–4 sequence fidelity approves before any dispatch.
3. Exit gate mapped to fixtures → fixture-id→task map + every task's Step-1 test list (B1-B4, A1-A4, C1-C7, R1, P1, L1, W1, F9, F10, F11, G, H, SWEEP — each with a named test file + red→green command).
4. Byte-exact enum → Global constraints; SWEEP leg tests it in emitted output.
5. Pivot shape from slice 1 → Tasks 1–2 build rename-pivot + `intake_id`-referencing outcomes before ANY feature code; F11 (Task 10) proves one-pivot-per-mutation.
6. Owed carries materialize-first → typed records (audit §4 / design §4): guardrail enforcement = Task 5 (G legs), I-PH fixture = Task 12 (P1), ③ portion = Task 4 (③ leg); each lands as a named fixture before S1 close.
7. Claim honesty → Global constraints + Task 12 SWEEP + README content spec.

## Out of scope (verbatim fence)

Everything on the ROADMAP scope-OUT list · genesis/quarantine/GC/segment-rotation/phase-0→4 machinery · m-2 §9 migrators (`schema_version` stamping only) · outbox drain/external send/egress scan (dormant by locked posture) · ⑤ `model_name` (S4-bound carry, cited only) · pair-Planner grant rendering (ratified narrowing; S3) · organic master-trail corpus as R1 gate inputs · any `../master` or `../extracted` write · real-runtime lane qualification (operator-gated spikes).

## Verification summary

`go test ./...` = the E2 umbrella (every fixture above is a `go test` target; crash fixtures spawn+SIGKILL child processes at named crash-points). Per-fixture red→green commands are in each task. The R1 report + F11 sweep output are committed artifacts. Evidence level at S1 close: E2 (local fixtures green); merge/live-verify are separate gates.
