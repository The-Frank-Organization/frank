# T9 runnable exit-fixture and decision-⑤ egress evidence

Base task commit: `1b87261`

## Runnable set

The consolidated command executes the landed T1 through T6 mechanisms plus
T9's decision-⑤ pair. T7 and T8 are absent because their required owner
co-signs did not arrive; both remain CLOSED rather than appearing as skips.

- T1: every detachable filesystem-worker stage, breaker, component swap, and
  read-file machinery boundary.
- T2: complete scan, all five fail-closed incomplete rows, token boundary,
  binary domain, exact ceilings, symlink rule, saturation, timeout/breaker,
  I-PH, and ungoverned-lane refusal.
- T3: executor identity forgery, total verdict tuple matrix, output redaction,
  conductor-derived `signal_class`, and thickened claim rows.
- T4: the `lane_vcs` total input/config-transition tables and the no-serialized-
  filesystem absence floor.
- T5: lane-cannot-forge-operator attestation plus the conductor default.
- T6: decision-② includes `authority_mixed` to `held`.
- T9: a fixture-scoped, conductor-generated operator ODB labels and scans the
  `model_name` confidentiality finding as ready; the same content on a lane
  send is blocked.

## RED

The first T9 pair run used the normal renderer without the fixture-only runtime
ODB origin label. The real scanner found `model_name:confidentiality`, but the
ODB leg failed exactly at the decision-⑤ fence:

```text
Disposition = "egress_blocked", want "egress_ready"
Findings:[model_name:confidentiality]
```

The out-of-fence leg already passed by refusing the lane send.

## GREEN

The fixture-local renderer now labels only a system-produced, system-role,
operator-facing `record_kind: odb` model-name field. It feeds the existing
real `egress.Drain` and scanner; it does not import egress into production,
activate the away bridge, or make any external call.

Verification commands:

```text
go test ./test/fixtures -run '^TestS9Decision5' -count=1 -v
go test ./internal/observe ./internal/config ./test/fixtures -run 'Test(FSWorker|ReadFileRootFailure|FindReferences|ExecutorCannotForgeDifferentialIdentity|VerdictTupleMatrixRejectsContradictions|VerdictOutputRedaction|SignalClassDerivedFromSelection|ClaimRowsCarryRungAndSignalClass|DeclaredNone|UndeclaredOrGit|UngovernedRoot|RegistryEnvLaneVCS|AbsenceFloor|LaneVCS|LaneCannotForgeOperatorAttestation|S8Decision2|S9Decision5)' -count=1 -v
go test ./internal/observe/... -count=1 -v
go test -race ./test/fixtures -run '^TestS9Decision5' -count=1
rg -n 't\.Skip|B1|B-diff|B2|B3|B4|diff-shape|test-files-unchanged|operator-attestation positive' test/fixtures/s9_exit_test.go internal/observe
git diff -- internal/executor/executor.go
go test ./... -count=1
```

Observed: every focused and consolidated fixture passed; the race run passed;
the blocked-item and `executor.go` sweeps were empty; the full uncached battery
passed, with `test/fixtures` completing in `126.768s`.
