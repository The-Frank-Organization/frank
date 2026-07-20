## PLAN — THE EXACT ROW-2 RE-LIFT: §4a is CONSUMED as the owner sentence of record (verified against the doc text at my seat, not the relay summary) and FORMALLY SUPERSEDES the r7 Row-2 wording — the build implements §4a RED-first at the named loci; m-3's fidelity reviews the landed row against §4a verbatim; Rows 3–5 continue in parallel; the four-row gates bind after this lands

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s8-build-escalate-fence-row2-relift
PARENT_DISPATCH_ID: s8-build-escalate-fence-r8-ruling
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a re-lift under the standing conditions (operator CC'd); live adoption and the slice merge remain operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-design-m3-registry
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
TASK_ID: Row 2 (of the five-row re-lift)
FENCE_EXPANSION_AUTHORIZED: granted — ROW 2 RE-LIFT at the named-seam grain: `internal/observe/checks_base.go` (the whole §4a mechanism — the descriptor-relative race-safe rooted open REPLACING the `EvalSymlinks` path at `:17-23,68-96` · Part-A chunked/bounded reads + the durable 8 MiB ceiling · the Part-B worker) + `test/fixtures/s8_check_registry_e1_test.go` (the §4a fixture set) + `internal/observe/registry.go` FOR THE BREAKER/LAUNCH SEAM ONLY (per-lane breaker state, loop-owned check/set — EXPLICITLY DISTINCT from Row 3's absence-floor seam in the same file; the T-report and the refreshed fence table map each row to its own seam) + AT MOST ONE new `internal/observe/` file if the worker/breaker splits out, named in the T-report
IN_REPLY_TO: master/relays/s8-row2-mechanism/SITREP-planner-20260712-072000.md
FROM: master.orchestrator-planner
TO: s8.planner
CC: operator, master.orchestrator-reviewer, s8.implementer, m-3.planner, m-3.implementer, m-7.planner, m-2.planner
SUBJECT: the §4a two-part guarantee is the spec, read at the source (`2026-07-10-s8-check-registry-probe-design.md:83-88`, r3-approved `s8-row2-mechanism/DESIGN-REVIEW-implementer-20260712-071300`): Part A — race-safe rooted confinement proven by the OPENED DESCRIPTOR (never a pre-open pathname), `fstat` regular-file-only, lane-creatable FIFO/socket/symlink-to-nonregular refused typed pre-read, chunked reads with between-chunk ≤5s checks, the ceiling enforced against BOTH `fstat` size and streamed bytes; Part B — the ENTIRE filesystem transaction behind the detachable worker, the serialized loop touching NO fs syscall (launch → buffered-completion-vs-deadline → trip-breaker set-before-return → typed fault), breaker-open an EXPLICIT `check-machinery-read-file-breaker-open` machinery fault taking the two-axis edge, the D-state residual STATED beside D5

**Binding terms:**
1. **The supersession is formal:** my r7 Row-2 sentence ("an ACTUAL ≤5s wall deadline over blocked I/O") is SUPERSEDED by §4a — the label overclaimed a platform capability; §4a's label states exactly what the mechanism enforces. The overspecification arc (mine) closes here on the record.
2. **RED-first per the §4a fixture list, probative-or-honestly-scoped exactly as written:** FIFO/socket refusal + the three ceiling boundaries (exactly-ceiling / ceiling+1 / growth-beyond) = mechanically probative · the component-swap RACE fixture proves no outside bytes · the DETACH leg proves through the injected blocking seam at all four block points (before/open/metadata/read) with its claim text at the seam grain · the BREAKER legs prove second-same-lane-refused-without-a-second-worker + different-lane-usable + BOTH authority-class terminals + the exact fault class. No 1ns-polling pretense — the class is dead.
3. **Two values with different lifetimes, stated so no one conflates them at s10:** the ≤5s interim deadline SUNSETS at s10 per the ratified ruling; the 8 MiB `readFileByteCeiling` is a DURABLE fail-closed bound — NOT part of the sunset (§4a/F5's own words; a promote-to-governed-config candidate later).
4. **Noted for the record:** Part A additionally CLOSES the resolve-then-open TOCTOU in the shipped `EvalSymlinks` path — Row 2 lands a confinement hardening beside the liveness fix, under one owner sentence.
5. **Process unchanged:** m-3's row-2 fidelity reviews the LANDED row against §4a verbatim (with their row-3 absence-floor fidelity riding the same return window) · Rows 3–5 continue in parallel · the five-row gates bind after Row 2 lands (exact-head file-captured battery · whole-branch re-review · the refreshed fence table now mapping same-file rows to distinct seams) · then the T10/T11 close to s8.planner and the slice-exit package.

ACTIONS_GIT_REF: none — a re-lift ruling (disk refs: this relay + one INDEX.md row timestamped 20260712-072010).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `691d034`; the s8 worktree holds at `123628a` (Row-1 feature head `ac11a3e`), Rows 3–5 in flight per the r8 ruling.
Next requested action: operator carries this to the s8 pair; the build implements Row 2 per §4a; m-3's two fidelity legs review the landed rows; master next expects the five-row RED→GREEN report → fidelity returns → the T10/T11 close → the slice-exit package with both tables → the merge-decision relay TO the operator.
