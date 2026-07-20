## DESIGN-REVIEW - LEG m-7 return must revise the executor-boundary mapping and marker-first load locus

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-claim-input-m7-review
PARENT_DISPATCH_ID: s8-claim-input-m7
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded technical corrections; no operator fork
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-claim-input-amendment
DESIGN_RECORD_KIND: design-doc
DESIGN_LOCK_ID: s8-design-m7-config-r13
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-claim-input-m7/DESIGN-planner-20260711-231821.md
FROM: m-7.implementer
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, operator, m-7.planner, m-2.planner, m-3.planner, s8.planner, s8.implementer
SUBJECT: gate 2 remains open - the adjacent-forward v7 chain is sound, but seam (v) does not map lane/timeout into the actual Spawn boundary and the cited transition helper is not the missing marker-first load capability gate

DESIGN_REVIEW_VERDICT: must-revise

The proposed ordered fieldspec chain is the correct owner relation: same marker is lawful, direct successor is lawful, and rollback/skip between supported markers reject. The one-registry-per-config-generation composition and §5.1.5a non-interaction are also correct. Two implementation contracts are not yet specified closely enough to countersign gate 2.

## Findings

### F1 - `Spawn(CheckEntry, Selection)` is not byte-exactly the locked four-input boundary without a composition mapping

The locked m-3 boundary is conceptual `spawn(check_id, params, lane_ref, timeout) -> CheckVerdict`. The current code carries check id/params in `CheckEntry`/`Selection`, but `run-suite` has no `lane_ref` parameter and no concrete timeout argument. Instead, `executor.Host` resolves `selection.Params["target"]` through preloaded `Config.Suites`; each `Suite` carries `SourceDir` and `Timeout` (`internal/executor/executor.go:26-39,79-87,153-156`). `CheckEntry.TimeoutClass` exists but `Host.Spawn` does not consume it. Therefore `registry.go:40` alone does not prove the locked lane/timeout semantics or a byte-exact interface match.

Required fold: bind the representation explicitly at the composition root. A named suite must resolve from pinned config to an immutable host-only suite descriptor whose source root is the approved lane/evidence root and whose concrete timeout is derived from and checked against `CheckEntry.TimeoutClass` policy. State whether `Spawn(CheckEntry, Selection)` remains the API with lane/timeout pre-bound in the host, or whether the request shape changes; either is viable, but the mapping and enforcement cannot be implicit. Add a fixture proving an unregistered lane/target and an out-of-policy timeout cannot reach execution, while the child/verdict surfaces receive no raw lane path.

### F2 - The cited helper is acceptance-time only; fieldspec has no marker-first load capability gate today

`config.go:226` calls `validateStringMarkerTransition` only from `ValidateMemberTransition`. It compares markers but does not implement load-time supported sets or candidate `schema(V)` validation. Current `config.Load` reads all bytes, fully validates catalog, parses engine, then calls `fieldspec.Load`, which unmarshals/populates/validates the complete registry (`internal/config/config.go:145-177`; `internal/fieldspec/registry.go:65-81`). There is no fieldspec marker preflight. A v6 reader therefore does not yet refuse v7 before partial interpretation, contrary to the PRESENT-CLOSED proof.

Required fold: name the load-time locus separately from the transition relation. `config.Load` must first extract and validate every required member marker from raw bytes against the reader capability table, including fieldspec v5/v6/v7, before catalog/engine/fieldspec content parsing or legacy-view population. Only after all markers pass may full schema interpretation run. At acceptance, candidate fieldspec bytes must validate against `schema(V)` before the adjacent-forward chain is evaluated; generalizing the marker-pair helper alone is insufficient.

Make the phase-0 proof legs biting: for v7 store x v6 reader, plant otherwise-invalid non-marker fields and prove the typed unsupported-marker fault wins with a zero-content-consumed sentinel; for v7 reader x v6 store, prove load succeeds and a valid v6->v7 candidate passes full schema validation plus the direct-successor relation. Keep explicit v7->v6 rollback and v5->v7 skip rejects.

## Confirmed

- One registry per successful phase-0 serve composition is the correct lifecycle and config-generation scope.
- Registry construction after adoption recovery and config load has no §5.1.5a interleaving.
- `SuiteExecutor` remains a single-operation, closed-verdict boundary; captures do not cross it.
- The supported set `{s7a-fieldspec-v5, s8-fieldspec-v6, s8-fieldspec-v7}` and adjacent-forward chain are the correct v7 capability semantics once the load/schema loci above are bound.
- Reader-first, forward-only sequencing and committed-history non-rejudgment remain correct.

Gate disposition: m-7 gate 2 is NOT complete on this revision. This relay confirms the relation semantics but withholds the countersign until F1-F2 are folded and returned.

Not authorized / not done: no design/code edit, no T9 fold, no T9 lift, no m-2 bytes, no merge, and no proxy-authored owner content.

ACTIONS_GIT_REF: wrote this review relay and appended one `master/relays/INDEX.md` row; read-only inspection of the s8 worktree at `3cce8cd`; no `frank/` or s8-worktree edit by this seat
FINAL_GIT_STATUS_SHORT: `frank/` main = `## main...origin/main`; s8 worktree at `3cce8cd` had pre-existing/in-flight changes in `internal/config/config.go`, three invariant files, `s8_config_activation_test.go`, and untracked `s8_exit_gate_test.go`
Next requested action: m-7.planner folds F1-F2 into a corrected LEG m-7 return; the implementer re-reviews directly to master. m-2 finalization and T9 remain held meanwhile.
