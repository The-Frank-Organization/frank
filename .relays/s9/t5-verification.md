# T5 attestation-source negative evidence

Base task commit: `5f6a7ec`

## Verification-first result

Command:

```text
go test ./test/fixtures -run '^TestLaneCannotForgeOperatorAttestation$' -count=1 -v
```

The new fixture passed on its first run, selecting the plan's already-covered
branch rather than a production gap:

- a lane-supplied `attestation_source: operator` is rejected before the
  observe evaluator runs, with the typed
  `attestation_source:lane-supplied-system-field` bounce;
- a normal candidate is accepted with
  `attestation_source: conductor` stamped by the conductor;
- no production file changed, and the B4 operator-attestation positive remains
  deferred to item 9.

## Commit-point verification

Commands:

```text
go test -race ./test/fixtures -run 'Test(S8SuppliabilityGuard|LaneCannotForgeOperatorAttestation)' -count=1
go test ./... -count=1
git diff --check
```

Observed: focused race coverage passed; every package passed in the uncached
full battery, with `test/fixtures` completing in `127.178s`; diff whitespace
validation was clean. The four held T3 paths were shelved for this commit-point
run, so the result covers exactly HEAD plus the T5 test/evidence delta.
