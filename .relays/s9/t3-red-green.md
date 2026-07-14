# T3 conductor-bound verdict evidence

Base task commit: `8a2b73d`

Plan rev13 hash:
`50759189a826572c0ba22c3d81cd222df0a495c52e63c31278208a0a5ba521c2`

Guide-byte return:
`master/relays/s9-build-blocker-t3-origin/SITREP-planner-20260714-001500.md`

## Initial RED

The original T3 tests failed before the binding pass existed. They covered:

- executor identity disagreement and forged differential identity;
- contradictory outcome/predicate/rung tuples;
- cross-origin hostile class presentation in both directions;
- all four timing states plus closed-set/tuple disagreement;
- planted path/config/secret residue;
- conductor-derived `signal_class` and thickened claim rows.

The first wide regression run then found the rev11 table-totality defect before
commit: the landed read-file refusal family had no legal origin row and
collapsed to `check-machinery-verdict-tuple-invalid`, changing the pinned
non-authority terminal. T3 stopped and the owning m-3 guide returned the rev13
closed `base-refusal` row.

Rev13 RED command:

```text
go test ./internal/observe -run '^TestBaseRefusalRowPreservesNoVantageDisposition$' -count=1 -v
```

Observed for all three allowed details:

```text
FailingDetail:"check-machinery-verdict-tuple-invalid"
MachineryFault:true
FAIL
```

## GREEN

The binding pass now:

- carries a conductor-internal origin envelope without changing
  `CheckVerdict`;
- binds identity to `Selection` and faults disagreement;
- validates the total origin/tuple/class/timing table;
- accepts exactly the three rev13 base refusals with
  `MachineryFault:false`, using `timeout` only for
  `read-deadline-exceeded`;
- keeps read-file timeout and breaker-open on the machinery edge;
- rejects empty executor timing while conductor origins produce their own;
- discards hostile residue by collapsing invalid verdicts to typed symbolic
  faults;
- derives `signal_class` and writes rung, signal, and integrity into claim
  rows.

Focused commands:

```text
go test ./internal/observe -count=1 -v
go test ./test/fixtures -run 'Test(S8Decision2|S8ExecutableClaimAggregationPrecedence|S8Adversarial|S8CheckRegistry|S8Observe|S8Verdict)' -count=1
go test -race ./internal/observe -run 'Test(ExecutorCannotForge|VerdictTuple|TruncatedPass|ExecutorOrigin|ConductorPolicy|ValidPolicy|BaseRefusal|ReadFileMachinery|TimingBranches|VerdictOutput|SignalClass|ClaimRows)' -count=1
go test -race ./test/fixtures -run 'Test(S8Decision2|S8ExecutableClaimAggregationPrecedence|S8AdversarialExecutorIsolationAndVerdictIPHIntegrated)' -count=1
```

Observed: all focused and race checks passed. Legacy fixtures were corrected to
supply valid executor timing, use real base-check origins for refusal/no-vantage
aggregation, and assert the locked six-column result rows.

## Commit-point full battery

Commands:

```text
gofmt -w internal/observe/gate.go internal/observe/registry.go internal/observe/verdict_binding.go internal/observe/verdict_binding_test.go test/fixtures/s8_adversarial_test.go test/fixtures/s8_exit_gate_test.go
git diff --check
go test ./... -count=1
```

Observed: formatting and whitespace checks were clean; every package passed in
the uncached full battery, with `test/fixtures` completing in `123.587s`.
