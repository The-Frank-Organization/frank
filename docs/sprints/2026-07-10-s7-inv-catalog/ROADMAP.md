# s7 INV-CATALOG execution roadmap

Status: intake co-signed; implementation in progress.

Authority of record: `master/relays/s7-dispatch/PLAN-orchestrator-planner-20260710-032426.md`, consuming r2 `...-023635.md`, the four r3 replacements in `...-030148.md`, and VP approval `RECONCILE-orchestrator-reviewer-20260710-030737.md` under operator ruling B10.

Base: `frank/main@1d3e92cc1f9f810da538b0369be9440ffd878f0a`

Branch: `s7-inv-catalog`

This file is an execution pointer. It does not amend or restate authority from the master relay trail.

## Scope

In scope after the intake co-sign:

- `test/invariants/**`: the ten executable invariant checks and versioned catalog artifact.
- `docs/sprints/2026-07-10-s7-inv-catalog/**`: local execution records.
- `.relays/s7/**`: pair intake, review, and implementation reports.

Out of scope:

- All production code under `cmd/**` and `internal/**`.
- `internal/fieldspec/registry.json`, record-kind changes, and mechanism fixes.
- Section-7 digest pinning, which remains the mandatory s8 carry.
- Merge, release, and later Step-2 slices.

Any production defect exposed by a named invariant is reported to master and is not fixed in s7.

## Named checks

| Row | Test | Binding grain |
|---|---|---|
| 1 | `TestLawTerminalEnumByteExact` | Byte-exact `{accepted, rejected, held}` plus typed rejection of a forged fourth token at submit. |
| 2 | `TestLawThreeVerbSurface` | Exact seat-visible surface `{submit, project, read}`. |
| 3 | `TestLawR2NoModelPredicate` | No model-identity field in required, visible, or gate predicate grammar. |
| 4 | `TestLawDerivedOnlyActivation` | `minted` and `active` derive from committed records; `bound_now` is runtime-only and resets across restart. No activation or bound marker persists. |
| 5 | `TestLawSoleGovernedWriter` | Sole governed write path, with the D5 direct-store residual stated. |
| 6 | `TestLawPathHygiene` | Six-family exhaustive census, including the operator-only `seat_mint` reply with credential/endpoint carve-outs; complete forbidden-path scan and planted-leak negative. |
| 7 | `TestLawCanonicalWins` | Canonical records win over corrupt projections during rebuild. |
| 8 | `TestLawOnePivotPerMutation` | Exactly one canonical rename pivot per governance mutation; crash-before absent, crash-after present. |
| 9 | `TestLawIntakeOutcomeOneToOne` | At-most-one outcome at all times; unique non-empty outcome intake IDs; settled exactly one; pending recovery at most once; duplicate-content replay/coalesce; pending-zero is valid. |
| 10 | `TestLawRebuildBeforeOpen` | Submit unavailable before recovery/rebuild completes; Ready opens only afterward. |

## Work order

1. Obtain the `m-7.planner` intake co-sign before code. Complete: `PLAN-REVIEW-planner-20260710-033918.md`.
2. Build the invariant package and catalog with red/green cycles for test harness behavior. In progress.
3. Run `go test -count=1 ./test/invariants` with all ten names green.
4. Demonstrate the command-pinned red battery on scratch, discard the weakening, then rerun green.
5. Run the full uncached repository battery and `go vet ./...`; audit the diff as test-only.
6. Send the IMPL report first to `m-7.planner` for adversarial pair review, then follow the master fidelity and integration route.
