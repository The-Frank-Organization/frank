# s6 Slice-6 Implementation Plan — the transport fix (the co-signed set, whole)

> **For the executing Implementer:** execute under Superpowers `executing-plans`, task-by-task, in plan order. Steps use checkbox (`- [ ]`) syntax. Every task ends with the battery green and a commit. Implementation authority arrives ONLY via the delegated dispatch relay under the standing F2 conditions (incl. the THREE external verdicts) — never from this document.

**Goal:** the s5-dogfood transport failures (F2–F17) are dead at their roots on a conductor whose seat surface is still exactly `submit`/`project`/`read` — branch-A parenting, one envelope codec, stable-schema digest, idempotent intake, live mint/re-mint, the store lock, scoped waivers, lifecycle/boot/roster — with every exit-gate fixture landed red-first.

**Locked design:** `docs/sprints/2026-07-06-s6-slice-6/designs/s6-slice-6-design.md` **r2 = main@a499bc3** (DESIGN_DOC_ID `s6-slice-6-design`; GRILL_LOCK `s6-grill-s6-core` = doc §18; approving review `s6-core-design-r2-review-implementer` 20260707-005333). §-references below are design sections. **The design rules; this plan sequences it.**

**Architecture (one breath):** one exported address-list codec consumed by all four judges (§2); the digest covers only the stable schema surface (§3); PARENT becomes a commit-time system stamp with a validated-hint fallback that never bounces (§4); intake replays outcomes instead of re-executing (§5); `project()` defaults to accepted (§6); mint/re-mint ride the loop as `seat_mint` pivots with binding replacement (§7); waivers become scoped records with retraction (§8); `flock` phase −1 owns the root (§10); every rejection reply carries the recorded detail (§11); lifecycle/boot/roster derive from records + runtime channel state, no persisted marker (§12).

**Tech stack:** Go, existing module, **zero new dependencies** (`syscall.Flock` is stdlib on darwin+linux); existing crashpoint/child-SIGKILL harness; existing replay + zeroloss harnesses; fixture donors named per task.

## Global constraints (bind every task)

- The §0 design constraints whole: byte-exact `{accepted, rejected, held}` · three seat verbs (roster/audit = `project` payload parameters; a new `ToolSet` entry or `names()` member is the violation shape) · I-PH `Field:Class` path-free over every surface incl. the four new payload families · `bounce.Format` stays the single formatter · claim pins verbatim (§0) · **no-perf fence**: the per-submit `tables.Build` in `cmd/frank/main.go process()` and the `gate.Complete` store scans are NOT touched, even incidentally, in any task below.
- **Grill-locked shapes land verbatim (§18):** re-mint = binding replacement · lock = `flock` on `<root>/conductor.lock` at phase −1 · intake counter = segment-header high-water line · roster = the B-1.3 seven fields · F11 = two-legs-one-claim · registry = 7 transport rows + 2 boot fields + 1 named enum + 2 record classes + the `ORCH_REVIEW_WAIVER` retirement, NO activation-marker row.
- **The three DESIGN-REVIEW watchpoints (binding task criteria):** (i) live-mint credential delivery is phrased and implemented as an operator custody HANDOFF (m-7 A-3 custody-unchanged; exact delivery path under m-1/m-7 fidelity eyes — T11 stops on any fidelity objection on record at execution time); (ii) waiver-row fill-time absence is NOT claimed mechanically solved unless m-2's verdict sanctions a render rule — **submit-path rejection is the implemented floor** and every claim surface says so until then (T9); (iii) the A-2 GC-drained-segments + restart id-reuse leg is red-first (T8).
- **External verdicts fold bounded:** the m-1/m-2/m-7 packet verdicts (routed in parallel by the orchestrator) bind every task they touch on arrival, verbatim, the s4 F-S4-M1 pattern; a must-revise re-reviews only the named surface. No `DISPATCH IMPL` before all three approves are on record.
- Red-first on every exit-gate fixture: the fixture lands and FAILS before its production change (evidence in the task log); run-FAIL-first is asserted for every negative leg.
- Commit per task (`s6:` prefix); `go test -count=1 ./...` + `go vet ./...` green before every commit; `-race` where the task says.
- Branch `s6-transport-impl` off `main`; base = the dispatch-time head (docs-only ahead of `s5-close`).

## File map (the SCOPE_DIFF universe)

| surface | files |
|---|---|
| codec + render + digest | `internal/fieldspec/canonical.go`, `render.go`, `validate.go`, `fieldspec.go` (+ `*_test.go`) |
| registry | `internal/fieldspec/registry.json` |
| lineage | `internal/lineage/lineage.go` (+ `lineage_test.go`) |
| engine | `internal/engine/submit.go`, `loop.go` (+ new `internal/engine/parent_stamp.go`, `admission.go`, `lifecycle consult` in submit; `*_test.go`) |
| intake | `internal/intake/journal.go`, `writer.go` (+ `*_test.go`) |
| store | `internal/store/store.go`, `projections.go` (+ new `internal/store/lock.go`, `lock_test.go`; `*_test.go`) |
| recover | `internal/recover/recover.go` (+ `recover_test.go`) |
| seat | `internal/seat/binding.go` (+ `binding_test.go`) — the re-mint replacement path (grill-locked; m-1 §13.3 carry) |
| channel | `internal/channel/server.go` (+ lifecycle tests) — force-close on re-mint; `bound_now` exposure |
| tables | `internal/tables/tables.go` (+ test) — ContentHash wiring; lifecycle; waiver effective-state |
| assembly | `cmd/frank/main.go` — turn context, project views, roster, mint CLI genesis-only, lock wiring |
| shim | `cmd/frank-mcp/mcp.go`, `errors.go` (+ tests) — typed-detail consumer; reconnect-retry |
| fixtures | `test/fixtures/` new: `s6_codec_test.go`, `s6_parenting_test.go`, `s6_projection_test.go`, `s6_intake_test.go`, `s6_waiver_test.go`, `s6_lock_test.go`, `s6_mint_test.go`, `s6_lifecycle_test.go`, `s6_iph_test.go`; modify `f11_test.go` (+`seat-mint` class), `applicability_map.go`, `sweep_test.go` (floor extension) |
| replay | `test/replay/` — the G-1 traffic re-drive leg (new file `test/replay/dogfood/dogfood_test.go`, env-gated like zeroloss) |
| docs | `docs/ops.md` (remint/lock/roster/boot sections), `cmd/frank-mcp/README.md` (retry note) — non-root, claim-honest. Root `README.md`: checked at plan time — NO claim line falsified by the set; no delta, no fence ask (re-checked at T15 sweep; any surprise ⇒ EARLY fence ASK to the orchestrator, never a silent edit) |
| step-exit | `docs/sprints/2026-07-06-s6-slice-6/results/step-exit-procedure.md` (new; operator legs, documented never simulated) |

Everything else is OUT for SCOPE_DIFF purposes.

---

### Task 1 — the codec: one decode for all four judges (§2; kills F6/F7)

**Files:** Modify `internal/fieldspec/canonical.go`, `internal/lineage/lineage.go`, `internal/store/projections.go`; extend tests + Create `test/fixtures/s6_codec_test.go`.
**Interfaces produced:** `fieldspec.DecodeAddressList(raw string) ([]string, error)` (empty ⇒ `nil, nil`; wraps the existing canonical path) + `fieldspec.EncodeAddressList([]string) (string, error)`; lineage `addressedTo`/`addressedInHeader`/`checkReviewerVisibility` decode via it (comma-splits DELETED); `store.DeliveryRecipients` decode-fail returns `(nil, error)` — callers surface it as an engine invariant fault via the loop's existing fault disposition, never a silent Envelope.To-only fallback; markdown render + INDEX row print the full decoded TO/CC set.

- [ ] 1.1 Red: `TestCodecFullRecipientSetEveryProjection` (multi-TO+CC canonical submit ⇒ identical set in markdown header, INDEX row, mailbox intents, and the visibility gate's match — m-2 §7.1), `TestArchiveRelaysOneEncoding` (the §7.2 archive shapes: canonical-JSON CC passes typing AND visibility), `TestDeliveryDecodeFailIsFault` (a hand-built committed record with corrupt TO reaches delivery ⇒ typed fault, zero silent delivery), `TestSubmitIllFormedRecipientBounces` (`Field:Class`, path-free).
- [ ] 1.2 FAIL → implement → PASS; battery + `-race ./internal/lineage/ ./internal/store/`.
- [ ] 1.3 Commit `s6: one address_list codec — comma-splits deleted, silent-drop dead, full recipient truth in every projection [m-2 §2]`.

### Task 2 — F13: `validateRecordKind` reduces to layer 3 (§9)

**Files:** Modify `internal/engine/submit.go`; extend `internal/engine/pipeline_test.go` or new engine test.
**Interfaces produced:** layer 1 (enum membership) + layer 2 (seat-scope) live ONLY in `reg.Validate` (already true — verified); `validateRecordKind` keeps only per-kind required checks (owed pair, config_change; T9/T11 add waiver_retraction/seat_mint); the case-list membership + `default → unknown` DELETED.

- [ ] 2.1 Red: `TestOfferedKindsAccepted` (`disposition` from operator, `diagnostics` from a plain seat — both currently bounce `unknown`), `TestGenesisRejectedAtSeatScope` (public submit of `genesis` ⇒ `record_kind:seat-scope`, NOT `unknown` — m-2 §7.3 negative), `TestUnknownTokenBouncesAtMembership`.
- [ ] 2.2 FAIL → implement → PASS; battery. Commit `s6: record_kind three-layer — engine stops second-judging membership [m-2 §2.6]`.

### Task 3 — the registry pass, additive legs (§14; grill-locked reading)

**Files:** Modify `internal/fieldspec/registry.json` (+ registry_test).
**Content:** the seven transport rows (`parent_hint` id_ref lane-optional · `parent_hint_honored` system · `parent_provenance` system enum {woken_on,active_dispatch,dispatch_root,hint} · `routing_ref_honored` system · `rationale` text · `waiver_scope` object · `retracts` id_ref) + the two boot fields (`charter_loaded` bool `agent_enum_pick` `gate_referenceable:false`; `dispatch_status` enum via new named enum `dispatch_status {read, awaiting}`) + `record_kind` enum += `waiver_retraction`, `seat_mint` (both operator-only seat_scope) + `ORCH_REVIEW_WAIVER` row REMOVED. **Sequencing note (one logical pass, two commits):** the `PARENT_DISPATCH_ID` ownership flip to `system_only` is the pass's final line and lands INSIDE Task 6 (atomically with the stamp + fixture updates — flipping here would bounce every legacy PARENT-supplying battery test with no stamp yet). NO activation-marker row (the negative is fixture-guarded at T15).

- [ ] 3.1 Red: registry_test legs — row presence/shape for all nine + the enum members + `TestNoActivationMarkerRow` (grep the registry for any marker-named row ⇒ zero) + `TestOrchReviewWaiverRowRetired`.
- [ ] 3.2 FAIL → edit → PASS; battery (all additive/dormant — zero behavior change asserted by the untouched suites). Commit `s6: registry pass (additive legs) — 7 transport rows + 2 boot fields + 2 record classes; ORCH_REVIEW_WAIVER retired [VP-W3; grill §18]`.

### Task 4 — A-1: digest = stable schema surface (§3; kills F5)

**Files:** Modify `internal/fieldspec/render.go` (+ render/validate tests, `test/fixtures/` leg).
**Interfaces produced:** `Field.ConductorVolatile bool` set by CLASS at render (`parent_picker`, `recipient_picker`, grant-state-derived options, floor-trimmed monotonic options); `formForDigest` strips options/defaults for the whole class; digest input = config digest + seat pattern + phase + tier + field shapes.

- [ ] 4.1 Red: FX-A1a (`TestForeignAcceptsNeverRotateDigest` — N accepts between render and submit ⇒ zero `re-render`), FX-A1b (`TestConfigChangeStillBouncesStaleForm` — the §7 record rotates it), `TestVolatileClassStripped` (unit: every volatile-class field's options absent from the digest input; static enums present).
- [ ] 4.2 FAIL → implement → PASS; battery. Commit `s6: digest covers the stable schema surface; conductor-volatile class exempt [m-7 A-1]`.

### Task 5 — D-2: reply detail parity, engine-side (§11; kills F3)

**Files:** Modify `internal/engine/loop.go` (+ `Outcome`), `internal/engine/submit.go`, `cmd/frank-mcp/mcp.go`; tests both sides.
**Interfaces produced:** `Outcome.Detail string` = the rejected record's Body (the `bounce.Format` output) for EVERY rejection class, byte-equal; held/fault outcomes carry their existing typed reasons unchanged; the shim's `submitNeedsReRender`/`containsReRender`/`reRenderResult` record-readback hack DELETED — the schema-refresh trigger keys on `form_digest:re-render` appearing in the typed `Detail` field.

- [ ] 5.1 Red: FX-D2 (`TestReplyDetailEqualsRecordedDetailPerClass` — table-driven over: re-render, required, enum, seat-scope, canonical-encoding, lineage class, layer-3, already-resolved; each asserts `out.Detail == rec.Body` byte-for-byte), shim test `TestShimRefreshOnTypedReRenderDetail` (no `read` call issued — assert via the fake conductor's call log).
- [ ] 5.2 FAIL → implement → PASS; battery. Commit `s6: every rejection reply carries the recorded detail; shim readback hack retired [m-7 D-2]`.

### Task 6 — branch-A parenting: the stamp, the hint, the flip (§4 + GRILL_LOCK; kills F11/F4/F2)

**Files:** Create `internal/engine/parent_stamp.go` (+ test); Modify `internal/fieldspec/registry.json` (the PARENT flip — the pass's final line), `internal/fieldspec/validate.go` (delete :69-76 parent-candidate checks), `internal/lineage/lineage.go` (delete `ActiveLineageCandidates`), `cmd/frank/main.go` (env wiring); legacy-test updates (PARENT-supplying fixtures move to `parent_hint` or drop the header — enumerate every touched test in the task log); Create `test/fixtures/s6_parenting_test.go`.
**Interfaces produced:** stamping INSIDE `SubmitHandlerWithRender` (recovery parity by construction — recovery re-execution uses the same handler): `PARENT = firstDefined(wokenOnAccepted, activeDispatchLineage, dispatchRoot)`; hint honored iff resolvable in (seat's delivered mailbox history ∪ own accepted records) ∩ accepted graph, with the generous prover (dispatch-id → accepted thread root; relay-id direct); unprovable ⇒ computed default + `parent_hint_honored: no` + verbatim hint preserved — **never a bounce**; `parent_provenance` stamped; the same generic hint-validation path stamps `routing_ref_honored` for `lineage_role: routing_ref` fields (§14.4, dormant consumer). Payload-supplied PARENT ignored byte-for-byte. Class-lineage gates untouched.

- [ ] 6.1 Red: G-2 (`TestHintHonored`, `TestHintFallbackNeverBounces` — unprovable hint ⇒ accepted, flag `no`, hint preserved verbatim), G-3 (`TestConcurrentAcceptNoParentClassBounce` — goroutine accepts landing during a submit window ⇒ zero parent-class bounces, loop-serialized), m-2 §7.4 (`TestParentAbsentFromEveryRenderedForm` — every seat×phase render), `TestPayloadParentIgnored`, `TestClassGatesStillBite` (delegated-grant relay with a fallback-stamped wrong parent ⇒ the CLASS bounce, with T5 detail), `TestRecoveryReexecutionStampsIdentically` (crash → re-drive ⇒ same PARENT stamped).
- [ ] 6.2 FAIL → implement → PASS (legacy-test updates itemized); battery + `-race ./internal/engine/`. Commit `s6: conductor-computed PARENT + validated hint, fallback-never-bounce; anchoring-bounce class deleted [m-1 §A; GRILL_LOCK]`.

### Task 7 — §B: projection default-accepted + accepted anchors (§6; kills F10)

**Files:** Modify `internal/store/store.go` (`Project` signature gains view/state awareness), `internal/store/projections.go` (rebuild filtering), `cmd/frank/main.go` (`turnContextForSeat`, project tool params); Create `test/fixtures/s6_projection_test.go`.
**Interfaces produced:** `store.Project(seat)` = accepted-addressed only (serve-time filter against canonical state); `ProjectView(seat, view)` with `view ∈ {default, audit}` (audit = the caller's own attempts incl. rejects; roster arrives T12); `RebuildProjections` writes mailbox intents for accepted only + truncates-and-rebuilds mailbox files; `turnContextForSeat` consumes the filtered view (accepted `WokenOn` by construction — feeds T6's stamp).

- [ ] 7.1 Red: m-1 §E legs — `TestPollutedArchivedMailboxFiltered` (build a store with rejected ids hand-appended to a mailbox, rebuild ⇒ default project excludes them; the REBUILT files clean), `TestAuditViewReturnsOwnAttempts`, `TestAcceptedDeliveryUnaffected`, `TestRejectNeverWokenOnAnyInterleaving` (table over commit orders), `TestOwnRejectReadable`, `TestHeldOffSeatDefault`.
- [ ] 7.2 FAIL → implement → PASS; battery. Commit `s6: project() default-accepted + audit view param; rebuild filters by delivery state; anchors accepted-only [m-1 §B]`.

### Task 8 — A-2: replay, coalesce, durable counter (§5; kills F9) **[watchpoint (iii): the GC leg is red-first]**

**Files:** Modify `internal/intake/journal.go` (segment header line; `Append` id-mint path deleted/unified), `internal/intake/writer.go` (no re-enqueue on dup), `internal/engine/loop.go`+`submit.go` (replay consult via wired `tables.ContentHash`/`OutcomeByIntake`; in-flight reply-holder), `internal/tables/tables.go` (populate `ContentHash` at OnCommit/Build from journal-carried hashes — mechanism: the outcome record's `intake_id` joins the journal entry's `content_hash` at table-build); Create `test/fixtures/s6_intake_test.go`.
**Interfaces produced:** segment header `{"segment_header":true,"high_water":N}` first-line (readers skip it; legacy headerless tolerated); `next = max(headers, entries)+1`; dup-with-outcome ⇒ replayed original outcome (same relay_id/state/Detail, nothing commits); dup-in-flight ⇒ same id, no second job, all reply channels answered by the one execution.

- [ ] 8.1 Red: FX-A2a (`TestDuplicateReplaysOriginalOutcomeByteIdentical` — incl. `Detail`), FX-A2b (`TestInFlightCoalesceSingleExecution`), FX-A2c (`TestIdUniquenessAcrossRotationCrash` + **`TestNoIdReuseAfterGCAndRestart`** — GC collects drained segments, restart, fresh submit ⇒ id strictly greater than every historical id; RED against today's max-of-surviving derivation) + `TestOneOutcomePerIntakeSweep` (the 1:1 invariant over any store), the A-2.5 duplicate-retry-across-crash leg.
- [ ] 8.2 FAIL → implement → PASS; battery + `-race ./internal/intake/ ./internal/engine/`. Commit `s6: intake idempotency — outcome replay, in-flight coalescing, segment-header high-water counter [m-7 A-2; grill §18]`.

### Task 9 — §C: scoped waivers + retraction (§8; kills F17/F12) **[watchpoint (ii): submit-path rejection is the implemented floor]**

**Files:** Modify `internal/lineage/lineage.go` (`checkReviewerVisibility` rewrite), `internal/engine/submit.go` (waiver-row operator check; `waiver_retraction` layer-3), `internal/tables/tables.go` (Waivers + retraction state); Create `test/fixtures/s6_waiver_test.go`.
**Interfaces produced:** waiver = accepted record with `waiver_scope` present + channel-stamped operator FROM; effective state = accepted matching-scope waivers − accepted retractions, commit order at gate time; retraction = `record_kind: waiver_retraction` with `retracts` id_ref (the `disposes_owed` idiom: exists + accepted + not-already-retracted); legacy unscoped `ORCH_REVIEW_WAIVER` records = run-wide until retracted; non-operator submit carrying waiver rows ⇒ typed reject (THE FLOOR — no fill-time-absence claim unless m-2's verdict sanctions a render rule; every doc line written by this task states the floor honestly).

- [ ] 9.1 Red: m-1 §E legs — `TestScopedWaiverPassesOnlyInScope` (relay/dispatch/record-class×dispatch/run legs), `TestRetractionReArmsProspectively` (post-retraction identical submit bounces), `TestPreRetractionAcceptsStand`, `TestNonOperatorWaiverRowsRejected`, `TestLegacyUnscopedRunWideUntilRetracted`, m-2 §7.5 (`TestRationaleFirstClass`, gate reads the record not the retired header).
- [ ] 9.2 FAIL → implement → PASS; battery. Commit `s6: scoped waivers + waiver_retraction; visibility gate reads effective state [m-1 §C; m-2 §4]`.

### Task 10 — A-4/§D: the store lock (§10; kills F14; grill-locked flock)

**Files:** Create `internal/store/lock.go` + `lock_test.go`; Modify `internal/recover/recover.go` (phase −1), `cmd/frank/main.go` (serve-lifetime hold + typed refusal exit); Create `test/fixtures/s6_lock_test.go` (child-process legs).
**Interfaces produced:** `store.AcquireRoot(root) (*RootLock, error)` — `flock(LOCK_EX|LOCK_NB)` on `<root>/conductor.lock`; content (diagnostic only) = holder pid + start time; typed path-free refusal error naming holder identity + the documented remedy; `RootLock.Release()` on clean shutdown; takeover = plain acquisition after holder death + an auditable store-visible diagnostics record post-recovery.

- [ ] 10.1 Red: FX-A4a (`TestSecondConductorRefusesTyped` — child process holds; second `run()` exits with the typed class, output path-free asserted by grep), FX-A4b (`TestKill9TakeoverRecoveryConverges` — SIGKILL the holder child, reacquire, full recovery, diagnostics record present), §E legs `TestAliasPathsCannotDoubleServe` (symlinked root ⇒ one winner), `TestS4GateDayLeftoverScenario` (a serving conductor + a second start on the same root ⇒ refusal), `TestNoTwoLiveClaimsInterleaving`.
- [ ] 10.2 FAIL → implement (acquire BEFORE phase 0 in `RunWithProcessor`; loser refuses reads — full exit) → PASS; battery. Commit `s6: I1-P store lock — flock phase −1, kernel-bound proof-of-death, alias-safe [m-1 §D + m-7 A-4; VP-W2 both halves; grill §18]`.

### Task 11 — A-3: live mint + re-mint (§7; kills F15; grill-locked binding replacement) **[watchpoint (i): custody handoff]**

**Files:** Modify `internal/engine/submit.go` (`seat_mint` classifier: operator-only, Body JSON `{seat, role, is_operator}`, reserved-name/self-remint guards), `internal/seat/binding.go` (`MintOrReplace` inside derived work; `Mint` keeps single-generation semantics for genesis/CLI), `internal/channel/server.go` (force-close by seat on replacement), `internal/obligation`/derived-work wiring (endpoint leg keyed by the record), `cmd/frank/main.go` (CLI `-mint` genesis-only refusal), `test/fixtures/f11_test.go` + `applicability_map.go` (`seat-mint` class); Create `test/fixtures/s6_mint_test.go`.
**Interfaces produced:** accepted `seat_mint` = the pivot; derived work = binding create-or-REPLACE (fresh credential; old dies at `Resolve`; live old-credential channel force-closed); the new credential rides ONCE in the operator's submit reply (`Outcome` extension for this class only — never in any record/projection/log; the I-PH matrix at T15 asserts it); crash-window remedy documented in ops.md (admin-time binding-table read — the custody HANDOFF wording verbatim; m-1/m-7 fidelity eyes confirmed on record before this task executes).

- [ ] 11.1 Red: FX-A3a (`TestLiveMintZeroDropsZeroRotation` — mid-traffic mint: existing channels alive, zero digest rotations), FX-A3b (`TestCrashBetweenPivotAndBindingConverges` — derived-work completion at recovery, exactly one binding row), `TestRemintReplacesBindingOldCredentialDies` (+ force-close leg), `TestSeatMintOperatorOnly`, `TestReservedNameRejected`, `TestCredentialNeverPersisted` (grep records/projections/INDEX for the credential ⇒ zero), crash class legs in f11 (applicability row falsifiable).
- [ ] 11.2 FAIL → implement → PASS; battery + `-race ./internal/channel/ ./internal/seat/`. Commit `s6: seat_mint loop mutation — live mint + binding-replacement re-mint; CLI genesis-only [m-7 A-3; m-1 §F; grill §18]`.

### Task 12 — B-1/B-2/B-3: lifecycle, boot form, roster (§12)

**Files:** Modify `internal/tables/tables.go` (lifecycle derivation), `internal/engine/submit.go` (B-1.2a admission — the LITERAL allowlist), `internal/fieldspec/render.go` (pre-active boot form), `cmd/frank/main.go` (roster view; render env lifecycle input), `internal/channel/server.go` (`ActiveSeats()` for `bound_now`); Create `test/fixtures/s6_lifecycle_test.go`.
**Interfaces produced:** lifecycle per seat from records (generation = latest committed `seat_mint` pivot or genesis-seed; `active` = first accepted governed `FROM=<seat>` submit within generation; operator always-active); admission: pre-active header set ⊆ {stamped envelope} ∪ {SUBJECT} ∪ {charter_loaded, dispatch_status}, zero others ⇒ accept, else ONE terminal `boot-required` reject with `<field>:non-boot-before-active` per-field detail (T5 parity); pre-active render = the boot form; `project(view=roster)` operator+orchestrator-only, the grill-locked seven fields, typed refusal otherwise; `project`/`read` NEVER lifecycle-gated.

- [ ] 12.1 Red: FX-B1a (mint→connect→boot roster walk, path-free leg), FX-B1b (non-boot pre-active ⇒ typed reject + the rendered form IS the boot form), FX-B1c (restart re-derivation by ORDER rule alone; all unbound at open; reconnect re-binds, NO re-boot), FX-B1d (reads never gated; non-privileged roster refused typed), FX-B1e (boot set + one extra registered header, + one unregistered ⇒ NOT active, per-field detail), FX-B1f (activation exactly once; already-active boot-shape = ordinary accept, no second edge), FX-B1g (re-mint ⇒ `minted` for the new generation; pre-re-mint accepteds do NOT activate; fresh boot accept does — [VP-W2]), m-2 §11.5 three (`TestBootFormRendersPreActiveUnbounceable`, `TestNoNewSharedVocabFromBoot`, `TestCharterLoadedSelfReportedNotGateInput`).
- [ ] 12.2 FAIL → implement → PASS; battery. Commit `s6: lifecycle minted→bound→active derived-only; literal boot admission; seven-field roster [m-7 B-1; m-2 §11; m-1 §F; grill §18]`.

### Task 13 — D-1: shim transparent reconnect-retry (§13; kills F16)

**Files:** Modify `cmd/frank-mcp/mcp.go` (+ tests).
**Interfaces produced:** on call failure over a previously-live client: close → one re-dial+re-auth → retry the SAME call once → second failure surfaces the typed class. Safe for submit by T8's replay (stated in the code comment as the constraint the retry depends on).

- [ ] 13.1 Red: FX-D1 (`TestConductorRestartNextSingleCallSucceeds` — kill+restart the served socket between calls; the next ONE call succeeds), `TestDoubleFailureSurfacesTyped`.
- [ ] 13.2 FAIL → implement → PASS; battery. Commit `s6: shim one transparent reconnect-retry [m-7 D-1]`.

### Task 14 — G-1: the archived-dogfood traffic re-drive (§16; two-legs-one-claim, grill-locked)

**Files:** Create `test/replay/dogfood/dogfood_test.go` (env-gated `FRANK_S6_DOGFOOD_STORE`, default = the archive path convention; skips when unset — the zeroloss pattern).
**Content:** reconstruct the archived store's submit SEQUENCE (records + intake journal, commit order); re-drive it seat-by-seat against a FRESH conductor at the fixed surface (render → fill from the archived record's lane fields → submit), store moving underneath exactly as recorded; assert: every archived-ACCEPTED submit accepts, ZERO parent-class bounces, ZERO digest re-render bounces (the A-1/branch-A claim on the recorded pattern). **The claim line in the test doc-comment is exactly:** "the recorded dogfood pattern completes without livelock; the race class is proven by TestConcurrentAcceptNoParentClassBounce (G-3)" — neither leg overstates. Composes with (never duplicates) the zeroloss read-integrity leg.

- [ ] 14.1 Red against a pre-fix build is impractical in-branch; the red-first evidence = running the reconstruction against the s5-close surface in a scratch worktree (expect parent/digest bounces — record the run in the task log), then green on the branch.
- [ ] 14.2 Green + battery. Commit `s6: F11 regression — archived dogfood traffic re-drive, two-legs-one-claim [GRILL_LOCK; grill §18]`.

### Task 15 — E2 floors + the I-PH matrix + sweeps (gate 2, mechanical)

**Files:** Create `test/fixtures/s6_iph_test.go`; Modify `test/fixtures/sweep_test.go` (floor extension over new outputs).
**Steps:** [ ] 15.1 I-PH grep-matrix over the four NEW payload families (roster rows · `boot-required` per-field detail · lock-refusal output · hint flags) + the seat_mint reply-custody leg — real bytes, zero path/credential occurrences (plant-a-leak double proves each scanner bites); [ ] 15.2 enum floor extended (the byte-exact triple + no new state token in any new output surface; `"bounced"`/`"submitted"` grep stays zero); [ ] 15.3 three-verb floor (tools/list through shim + server = exactly three; roster/audit reachable only as project params); [ ] 15.4 full uncached battery `go clean -testcache && go test ./...` + `go vet` + `-race` on engine/intake/channel/store/lineage — ZERO regression (s1–s5 suites untouched except the T6-itemized legacy-parent updates); [ ] 15.5 root README re-check (no falsified claim line; else STOP ⇒ fence ASK); docs/ops.md + shim README claim-honest sweep (transport-only line; the waiver floor wording; the custody handoff wording). Commit `s6: floors — I-PH matrix over new surfaces, enum/verb sweeps extended`.

### Task 16 — the step-exit procedure of record (operator legs; documented, never simulated)

**Files:** Create `docs/sprints/2026-07-06-s6-slice-6/results/step-exit-procedure.md`.
**Content contract:** the exact operator-run sequence for ROADMAP gate 3 on the FIXED conductor — fresh blessed store init; ops pre-allowlist `mcp__frank__*` in every seat session; (a) the ROADMAP:83-85 legs; (b) first live act = the operator §7-applies s5's registry (operator-authored); (c) the F11 regression leg live (the archived pattern against the live store); (d) live boot: A-3 mint (no restart) → wire → B-2 boot → `active` derived → roster shows it; every leg names its evidence artifact; every claim line transport/provenance-only. Dry-run: conductor-side steps against a scratch store ONLY (no operator leg simulated, no live-seat designation assumed).

- [ ] 16.1 Write; [ ] 16.2 scratch dry-run of conductor-side steps; [ ] 16.3 commit `s6: step-exit procedure of record (operator legs)`.

---

## Acceptance criteria (exit-gate mapping)

Gate 1: G-1/G-2/G-3 → T14/T6 · FX-A1a..FX-B1g (18) → T4/T8/T11/T10/T13/T5/T12 · m-1 §E → T7/T9/T10 (+T6 anchors) · m-2 §7/§11.5 → T1/T2/T6/T9/T12 · §F.6 lands via FX-B1c/f/g (T12 — the stated mapping). Gate 2: T15. Gate 3: T16 (operator-run). Gate 4: claim pins per task + T15.5 sweep. Anti-half-fix guards: every negative red-first with evidence; the applicability map stays falsifiable (T11); no fixture asserts through a deleted path (comma-split, candidate set, readback hack, broadcast — grep-guarded at T15).

## Out of scope (verbatim fence)

Step-2 observe · Step-3 routing execution · engine performance work of ANY kind (incl. the two named hot-path rescans inside edited files) · new seat verbs · federation · dogfood-in-slice · governance-doc edits · locked-contract changes (amendment path only) · root README (no falsified line at plan time) · the frozen s2-store.

## Operator-judgment items

Live-seat designation + every step-exit leg (gate-day); s6-close/merge (never implied by green fixtures). The grill items: RESOLVED (§18). Pending external: the m-1/m-2/m-7 verdicts (fold bounded on arrival; no dispatch before all three).

*Rev log: r1 (post-PROCEED-TO-PLAN, for Implementer PLAN-REVIEW).*
