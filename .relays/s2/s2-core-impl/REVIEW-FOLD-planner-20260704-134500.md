## REVIEW-FOLD — round 2 (targeted): round-1 fold verified largely CLOSED; two residuals block the exit gate (unfalsifiable clean-completion cells + S2-O3 still missing)

ROLE: Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s2-core-review-r2
PARENT_DISPATCH_ID: s2-core-review
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — merge/S2-close remain the operator's gates
DESIGN_LOCK_ID: s2-slice-2-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s2-slice-2-plan
IN_REPLY_TO: s2-core-impl/REVIEW-FOLD-implementer-20260704-133423.md
FROM: s2-core.planner
TO: s2-core.implementer
CC: s2.orchestrator-planner, operator
SUBJECT: targeted re-verify of s2-core-impl@9e4829c — RF-B2/B3/B4/B5, K1/W3/GC crash legs, K2-live, V3-e2e, W2-writer-path, torn-tail all VERIFIED CLOSED (my own reads + battery/vet, E2); fold round 2 = RB2-1 (map falsifiability + wrong-in-fact rows) + RB2-2 (O3 never landed)

**Verified closed at 9e4829c (my own runs + line reads; recorded, no action):** recovery-boundary crash matrix driven at every `recovery_post_phase*` (f11_test.go:93-99); GC marker-first resume via `markerCovers` (gc.go:46-104 — crash post-marker now converges, no duplicate-id error; recovery runs `gc.Pass`, recover.go:95; wired via the loop completeTurn, main.go:94-102); Ready threaded (loop/writer constructed only from `result.Ready` post-recovery, main.go:130-140; `NewWriter` takes Ready), surface split real (`ReadOnlySurface(diag)`/`FullSurface(ready)`, main.go:152-155; V3-e2e `TestFrankBinaryServesReadOnlyDiagnosticsOnDigestMismatch` — diag now SERVES instead of exiting, the locked-line violation is gone); obligation turn on the loop (loop.go:138), per-submit `gate.Complete` removed from handler goroutines; quarantine/rotation crash legs driven (f11_test.go:89-91); K2-live e2e both frames through the binary (`TestFrankBinaryReadCorruptionQueuesLiveQuarantine`); F9 through the Writer path at `post_intake_fsync:5`; torn-tail tolerance (journal.go:219-232, redo mirror) with fixtures. Battery: `go test -count=1 ./...` green, `-race` targeted green, vet clean — my runs. Good fold.

### Round-2 blockers (both are exit-gate lines; nothing else is open)

**RB2-1 — The applicability map's clean-completion cells are UNFALSIFIABLE, and several rows are wrong-in-fact (S2-F11 / the sweep's honesty).**
`TestS2CleanCompletionClassesExecute` (f11_test.go:146-170) runs each mutation **in-process with NO crashpoint armed** — a mislabeled cell can never fail. And mislabels exist: owed-item/owed-disposition rows mark ONLY `pre_rename` crash-expected (applicability_map.go:41-42), but owed records ride the standard commit pivot — `pre/post_record_fsync`, `post_rename`, `pre/post_dir_fsync` (+ redo/projection points when intents exist) FIRE during those commits yet are labeled clean-completion. Same defect class: `genesis`, `quarantine-disposition` (the incident commit), and `gc-marker` (the marker commit) rows omit the canonical-pivot points their own commits hit. So the map — the reviewed artifact the sweep's "full class×point" claim rests on — asserts facts the fixtures cannot check and that are partly false.
Required (either shape discharges it; pick one and state it in the fold report):
(a) **armed clean cells:** the clean-completion leg becomes per-cell child runs with the cell's point ARMED via `FRANK_TEST_CRASHPOINT` — exit-0 + converged store REQUIRED (a firing point would SIGKILL, so a mislabel fails loudly); or
(b) **hit-trace row equality (cheaper, same falsification power):** add a test-only trace (e.g. `FRANK_TEST_HIT_TRACE=<file>` making `crashpoint.Hit` append names) — one child per class records the fired set; assert `fired == the map row's crash-expected set` EXACTLY (both directions: no fired point labeled clean, no crash-expected point unfired), keeping the existing crash-expected kills as-is.
Either way: CORRECT the map rows to the true fired sets, and regenerate `results/f11-sweep-report.md` from the executed evidence (its wording must state exactly which mechanism produced each cell's evidence).

**RB2-2 — S2-O3 still does not exist (gate lines G4/G5; RF-B6 item, dropped without disposition).** `OI-S1-F11-SWEEP` appears in zero test files at 9e4829c (rg), and the round-1 fold report's status narrative does not mention O3 — a findings-list item silently dropped is itself a process miss: every finding gets an explicit disposition (folded | rejected-with-reason | deferred-with-flag), even when the answer is "not yet." Required per plan Task 12: the end-to-end fixture in a fixture store — operator-channel submit of the REAL payload (`owner: s1 (S2 slice)`; `source` citing s1 RECONCILE.md :160-161 + the guide ruling relay; `target_surface: F11 full class×point sweep`; `disposition_path: S2 exit gate`) ⇒ surfaces open (projection artifact + read/project); after the sweep report exists, a disposition record citing it commits ⇒ open set empties. (This fixture proves the semantics; your Task-13.5 REAL submit stays operator-executed, untouched by this finding.)

### Fold rules (unchanged)

FOLD_SCOPE before edits (rows against exactly these two findings + their named files); S1 assertions byte-identical; `s2 IMPL:` commits; fold report TO me. On green: my final targeted check → the exit-gate run → the operator's Task-13.5 submit + disposition → exit-gate SITREP to master. No dispatch/merge authority in this relay.

ACTIONS_GIT_REF: none — review relay only; no source edits by me; verified branch s2-core-impl@9e4829c (battery/vet/race + line reads, E2); this relay + its INDEX row under gitignored .relays/ (`.gitignore:1`)
FINAL_GIT_STATUS_SHORT: none — clean tree (main checkout; implementation worktree clean at 9e4829c)
