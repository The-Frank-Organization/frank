# s11 engine-v4 cadence re-home — RED/GREEN evidence

Reviewed base: `e86644ddf10ca9bbdc4c098f443ad3eab73c4e20`

Authority chain: master ruling `s11-build-escalate-config-lock/RECONCILE-orchestrator-planner-20260714-170510.md`; m-7 owner spec `SITREP-planner-20260714-171302.md`; corrected fold scope `s11-build-fold-v4/REVIEW-FOLD-planner-20260714-172200.md`.

## RED before production edits

After adding only the A–F contract fixture and re-pinning the existing cadence fixture writer to v4:

```text
$ go test -count=1 ./test/fixtures -run '^TestS11ResummonCadenceUsesEngineV4Contract$'
--- FAIL: TestS11ResummonCadenceUsesEngineV4Contract
    --- FAIL: .../v3_rejects_cadence_key
        v3 cadence load err = <nil>, want ErrConfigLoad
    --- FAIL: .../v4_accepts_cadence_key
        v4 cadence load: config-load: engine-marker
    --- FAIL: .../v4_keeps_cadence_optional
        v4 cadence-less load: config-load: engine-marker
    --- FAIL: .../v3_to_v4_is_the_only_new_adjacent_hop
        v3->v4: config-version-transition
FAIL
```

The four failures independently exercised the missing ruled behavior: restored v3 rejection, v4 admission with and without the optional key, and the v3→v4 adjacent-forward transition.

## Scope correction discovered by the required target

After carrying m-7 F1–F6, the new v4 contract test and existing cadence fixtures passed, while the required config package exposed one stale reader-ceiling fixture:

```text
$ go test -count=1 ./internal/config
--- FAIL: TestLaneVCSReaderCeilingRefusesV4BeforeSupply
    err=config-load: engine, want phase-0 engine-marker ErrConfigLoad
FAIL
```

The fold stopped before that OUT edit. `SITREP-implementer-20260714-171918.md` recorded `deviation-present`; corrected REVIEW-FOLD r2 then licensed exactly the test rename `RefusesV4`→`RefusesV5` and document marker 4→5. The phase-0 assertion stayed byte-unchanged.

## Targeted GREEN

At the complete corrected-scope diff:

```text
$ go test -count=1 ./internal/config
ok  github.com/jackli/frank/internal/config  0.377s
$ go test -count=1 ./test/fixtures -run '^TestS11(G4Cadence|ResummonCadence)'
ok  github.com/jackli/frank/test/fixtures  1.280s
$ go test -count=1 ./test/invariants
ok  github.com/jackli/frank/test/invariants  1.304s
$ go vet ./internal/config ./test/fixtures ./test/invariants
$ git diff --check
```

All commands exited 0. `test/invariants` ran the ten catalog-backed invariant laws. This is E2 local proof; no merge or live-deployment claim is made.

## Preserved bytes

The production diff changes only m-7 F1–F5: transition ceiling 4, reader ceiling 4, descriptor ceiling 4, v4 allowed-key arm, and v4 cadence presence gate. `ResummonCadenceConfig`, `EngineConfig.ResummonCadence`, `ResummonCadenceDelays`, `validResummonCadenceShape`, supply validation/composition, and `main.go` are byte-untouched.
