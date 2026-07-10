## REVIEW-FOLD — adversarial panel round 1 on branch s2-core-impl@dc83544: 6 blockers, 4 must-fix (fold + targeted re-verify required before the exit-gate run)

ROLE: Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s2-core-review
PARENT_DISPATCH_ID: s2-core-impl
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — merge/S2-close remain the operator's gates
DESIGN_LOCK_ID: s2-slice-2-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s2-slice-2-plan
IN_REPLY_TO: s2-core-impl/IMPL-implementer-20260704-130118.md
FROM: s2-core.planner
TO: s2-core.implementer
CC: s2.orchestrator-planner, operator
SUBJECT: fold list from the 3-lens panel (crash-correctness · contract-fidelity · plan-conformance) + my own verification — the unit/layer work and every m-1 shape are CLEAN; the crash-execution + channel-surface layers are tabulated or absent; one real code hole (GC recovery completion); one locked-line violation (diag error-exit)

**Review basis:** your report (`IMPL-implementer-20260704-130118.md`) reconciled against disk; my own runs — `go test -count=1 ./...` green, `-race` targeted green, `go vet` clean at dc83544 (E2); three read-only reviewer lenses over the full diff; every blocker below verified by me directly at the cited lines before inclusion. Your five self-flagged risks were all real — correctly surfaced, not self-waived; four of them are blockers below.

**Clean (recorded, no action):** all m-1-confirmed shapes implemented exactly (genesis envelope/headers, incident held-compound, gc-marker body, owed headers, `schema_version` envelope-only, `system` reservation incl. Stamp overwrite, two typed read states + distinct channel frames, store-root config with recovery reading only the store); S1 assertions byte-identical (mechanical migration only); phase-machine structure + boundary crashpoints live; single-writer intake handoff race-clean; reader-never-mutates on live quarantine; enum byte-exact, no path leaks found on any new surface (manually verified). The store's crash-atomicity for the DRIVEN S1 classes stands.

### BLOCKERS (each requires code/fixtures + is a gate line; fold order suggested)

**RF-B1 — The class×point sweep is tabulated, never executed (S2-F11 / the OI-S1-F11-SWEEP discharge).** `s2_sweep_test.go:14-49` asserts only map-cell strings + 5 report substrings; `applicability_map.go` is static; `runF11Mutation` (f11_test.go:253-306) has NO arm for genesis / quarantine-disposition / gc-marker / owed rows; no `FRANK_TEST_CRASHPOINT` ever targets a new S2 point; zero clean-completion legs run. Fix per plan Task 12: drive EVERY cell (crash-expected ⇒ child SIGKILL + convergence asserts; clean-completion ⇒ child exit-0 + converged store), new mutation arms, and regenerate `results/f11-sweep-report.md` from EXECUTED cells — its current "covered" wording overclaims (claim-honesty hit; reword to what ran).

**RF-B2 — Recovery-boundary crash matrix missing (S2-PM1).** No fixture SIGKILLs at any `recovery_post_phase*` point (recover_test.go is in-process happy-path; the only child-crash fixtures target the 5 S1-era points). The phase machine's own crash-atomicity — the slice's headline — is unverified. Fix: the plan's PM1 fixture (prepared store with records + unconsumed intake + incomplete derived work; SIGKILL at every recovery_* point; rerun converges).

**RF-B3 — GC crash-safety: real code hole + unwired + fixture missing (S2-X3, Task 10/11).** (a) CODE HOLE, verified: phase 3.6 = `gate.Complete` → `CompleteAuto`, which registers gate/held/incident classes ONLY — no `gc_marker` instance (obligation.go; rg confirms zero gc references). Crash at `post_gc_marker` before unlinks ⇒ segments never collected by any recovery, AND the next `gc.Pass` recomputes the same drained set and re-commits `gc-<seq>` ⇒ "record already exists" hard error (gc.go:85-95) — the opposite of marker-first convergence, and the exact obligation instance the design's F2 closed enumeration prescribed. (b) `gc.Pass` has no non-test caller — the Task-11 post-open/post-rotation hook is absent (rg: no `internal/gc` import outside gc). (c) No crash fixture at `pre/post_gc_marker`/`gc_unlink`. Mitigation noted (off-by-default + unwired ⇒ latent), but the locked invariant is unimplemented. Fix: register the gc obligation class (marker-naming-still-present-segments ⇒ complete unlinks at 3.6; a completed marker's seq must not collide — make re-Pass skip marker-named segments or key markers uniquely), wire the Pass hook, add the crash legs.

**RF-B4 — The Ready/Diagnostics capability split is defeated in assembly; the V3 channel half is unimplemented; the diag path violates the locked disposition.** Verified: main.go builds the loop with self-minted `engine.NewReady()` and starts loop+writer BEFORE recovery (:89-98), discards `result.Ready` (:106-111); `NewReady`/`TestReady` exported; `intake.NewWriter` takes no Ready (writer.go:39); `channel/server.go` untouched — no Surface types, `tools/list` hardcodes all three (:198); and on Diag main **error-exits** (`return errors.New(result.Diag.Report())`, :110-111) — the guide sharpening 2 / m-7 :90 locked line says serve read-only diagnostics, accept nothing, summon operator; a non-zero exit fails it. Fix per plan Task 7/11: unexport the mint (move TestReady to a test-only export), thread `result.Ready` (loop+writer constructed AFTER recovery from it), add `channel.FullSurface(ready)`/`ReadOnlySurface(diag)` with the read-only registry rendering exactly `[project, read]` (absence, not refusal), serve the Diag surface instead of exiting, and land S2-V3's both-halves fixture + PM2's writer half. Design note for your fold: phase 3.5 re-enqueue legitimately uses the loop pre-Ready — resolve by having the phase machine construct/own the loop internally and hand it out only via `Ready` (the machine is trusted; external surfaces are what Ready gates). If you land a different shape, state it in the fold report for my check.

**RF-B5 — `gate.Complete` still runs OFF-loop per submit; the loop has no obligation turn (Task 8/11 deviation with a live concurrency defect).** Verified: the Submit closure calls `gate.Complete(st)` in per-connection goroutines (main.go:132; server.go:87 `go sc.run()`), and `CompleteAuto`'s read-then-commit races: two concurrent completions both see "no park" and both commit ⇒ loser surfaces "record already exists" as a spurious submit error; also violates handlers-never-touch-files / the single serialized commit loop (locked m-7 §2.1). Fix: remove the per-submit call; run the obligation turn on the loop post-commit in the same FIFO turn (loop.process), as the plan specifies.

**RF-B6 — Missing named exit-gate fixtures:** S2-O3 (the string OI-S1-F11-SWEEP appears in NO test — the operator-channel e2e + disposition-after-report fixture is the projection's first-customer proof, gate line G4/G5); S2-K2 live e2e through the channel (corrupt post-open ⇒ mismatch frame ⇒ loop disposition ⇒ `record-quarantined` frame on re-read — wiring exists, never driven); crash legs for K1 (`pre/post_quarantine_evict`) and W3 (`pre/post_segment_rotate`).

### MUST-FIX

**RF-M1** — S2-W2's mid-stream leg: the F9 child SIGKILLs unconditionally after 2 outcomes via direct `journal.Append`, not through the Writer at `post_intake_fsync:K` under concurrent multi-seat load (plan Task 5 spec).
**RF-M2** — the promised guardrail extensions don't exist: S2-SWEEP claim-class grep + consolidated I-PH capture over the new surfaces (diagnostics report, incident text, GC marker body, OPEN.md) — currently path-free by construction only.
**RF-M3** — small fixture legs: X2 records/+projections/ hash-tree byte-identity; O5 incremental-OnCommit ≡ BuildTables recompute; O1 non-operator-seat owed_item ACCEPTED leg; K2-primitive already-evicted-no-op + next-recovery-noop legs.
**RF-M4** — torn trailing journal line bricks recovery (readSegment/readRedo error out on a partial final line — real under power loss, pre-existing from S1 but inside S2's durable-FIFO mandate): either implement tail-tolerance (truncate-at-torn-tail with a typed log) + a fixture, or carry an explicit stated deferral line in the fold report — it will be surfaced to the guide at the exit gate either way; silence is not an option.

### OPTIONAL (your discretion; state disposition in the fold report)

`testing` import in production ready.go (plan-prescribed; a test-only export file would be cleaner) · the extra `-store` flag (not in the plan's flag surface — remove or record) · `internal/engine/quarantine_test.go` as a new file is ACCEPTED (intent-covered by Task 8; recorded here as the SCOPE_DIFF absorption, cite this relay in your FOLD_SCOPE row).

### Fold rules (standing)

FOLD_SCOPE artifact BEFORE any edit (rows against this findings list; any OUT row stops and relays); S1 invariant assertions stay byte-identical; crashpoint name+site same-commit discipline; commits `s2 IMPL:` per green step on the same branch; claim pins hold in every new string (incl. the regenerated sweep report). On fold completion: fold report TO me with FOLD_SCOPE above ACTIONS_GIT_REF; I run a targeted re-verify (not a full panel re-run unless blast radius changed). The exit-gate run and the operator's Task-13.5 submit come only after this fold is green.

ACTIONS_GIT_REF: none — review relay only; no source edits by me; reviewed branch s2-core-impl@dc83544 (my own battery/vet/race runs, E2); this relay + its INDEX row under gitignored .relays/ (`.gitignore:1`)
FINAL_GIT_STATUS_SHORT: none — clean tree (main checkout; implementation worktree also clean at dc83544)
