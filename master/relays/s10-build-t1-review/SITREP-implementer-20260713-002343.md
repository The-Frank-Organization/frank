## SITREP — s10 T1 owner bytes land verbatim and targeted GREEN, but full battery exposes a T1/T2 sequencing collision at genesis plus stale historical v7 fixture constructors; hold before commit

ROLE: Implementer
PHASE: IMPL
AUTHORITY: implementation
DISPATCH_ID: s10-build-t1-review
PARENT_DISPATCH_ID: s10-build-impl
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this needs a bounded task-seam/verification ruling through the existing planner/master path; merge remains operator-only
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
PLAN_LOCK_ID: s10-comms-spine-plan-r2-s10.2
IN_REPLY_TO: s10-build-t1-review/SITREP-planner-20260712-233011.md
FROM: s10.implementer
TO: s10.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-2.planner, m-7.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: the seven-point owner set matched OLD and landed verbatim; T1 RED is GREEN and config/fieldspec packages are green, but `go test ./... -count=1` fails seven legacy fixture legs because the live-registry-as-v7 constructors and `genesisMemberBytes` predecessor derivation do not yet understand v8; genesis.go is assigned to T2, so no commit or adaptation

Summary: The m-2/m-7 bytes from `s10-build-fidelity-m2/COORD-planner-20260712-233000.md` and `s10-build-fidelity-m7/SITREP-planner-20260712-230402.md` matched `8941889` byte-for-byte and were landed once. m-6's concurrent return `s10-build-fidelity-m6/SITREP-planner-20260713-002109.md` confirms both EMPTY seat_scope and emit-only with no correction. The original T1 phase-0 refusal fixture is GREEN and the two directly affected packages pass uncached. The required full battery is red, so T1 is not committed and no later task has started.

Plan lock: `s10-comms-spine-plan-r2-s10.2`; T1 only. T2+ remain unstarted. m-6 Cell 2 is confirmed, so T3 has no remaining owner-cell hold but stays sequence-blocked behind T1/T2. T4 remains held on both Step-5 confirms.

SCOPE_DIFF:
- frank/internal/fieldspec/registry.json -> in
- frank/internal/fieldspec/registry_test.go -> in
- frank/internal/config/config.go -> in
- frank/test/fixtures/s10_fieldspec_test.go -> in
- frank/.relays/s10/t1-red-20260712-225520.txt -> in
- frank/.relays/s10/t1-full-battery-blocker-20260713-002343.txt -> in
SCOPE_DIFF_RESULT: all-in

Landed owner bytes:
- registry marker: `s8-fieldspec-v7` to `s10-fieldspec-v8`.
- `record_kind`: append `odb`, `resummon_command` only.
- reader ceiling: append `s10-fieldspec-v8` only.
- successor relation: append the single adjacent `s8-fieldspec-v7` to `s10-fieldspec-v8` pair once, dual-licensed by m-2/m-7.
- registry version pin: true to `s10-fieldspec-v8` only.
- seat_scope unchanged; no per-kind required_when.

Files changed:
- `frank/internal/fieldspec/registry.json`
- `frank/internal/config/config.go`
- `frank/internal/fieldspec/registry_test.go`
- `frank/test/fixtures/s10_fieldspec_test.go`
- two captures under `frank/.relays/s10/`

ACTIONS_GIT_REF: `s10-comms-spine@8941889a432a0a02a4183a16fb0015c230c9561b`; uncommitted T1 diff only; no commit because full E2 battery failed

Acceptance criteria status:
- T1 RED-first proof: met at E2, then GREEN at E2.
- v8 reader accepts the v8 marker before content interpretation; v7-capability reader refuses at phase 0: met by the targeted fixture.
- ONE schema transition / no seat_scope / no required_when: diff-confirmed at E1.
- m-6 Cells 1 and 2: confirmed with no correction at E1; EMPTY seat_scope and emit-only stand.
- Full battery and T1 commit: blocked.
- Step-5 owner confirmation: not requested yet because the candidate is not commit-ready.

Root-cause analysis:
1. `internal/store/genesis.go:98-134` only recognizes a live v7 registry, strips the v7 claim row to v6, then derives the lock-pinned v5 genesis. With a live v8 registry it recognizes neither v7 nor v6 and returns the v8 source unchanged. That breaks FX-CFG-7's pinned-v5 invariant and removes the old/new form distinction used by the stale-form test.
2. `s8FieldspecV6Bytes` reads the live registry and replaces a v7 marker. After T1 that replacement is a no-op, so transition/skip tests no longer construct their intended v6/v7 history.
3. The s8 marker-refusal, exact-v7, and duplicate-marker fixtures likewise mutate/assert the live registry as though it were still v7.
4. The production dogfood failure is downstream: `scripts/dogfood-suite.sh` recursively runs the same red fixture suite, so the governed `run-suite` result correctly reports fail.

Tests/verification:
- `go test ./test/fixtures -run '^TestS10V7ReaderRefusesV8MarkerBeforeContent$' -count=1` — PASS.
- `go test ./internal/config ./internal/fieldspec -count=1` — PASS.
- `go test ./... -count=1` — FAIL: seven fixture legs; capture `frank/.relays/s10/t1-full-battery-blocker-20260713-002343.txt`.
- Six direct failures reproduced in one targeted uncached fixture command; the seventh is the recursive dogfood observation of the same red suite.
- `git diff --check` — clean before this report.

Evidence levels: targeted/package commands and full/targeted failure reproduction are E2; code-path/root-cause mapping is E1.

Out-of-scope preserved: no `genesis.go`, legacy fixture, T2+, T3+, executor, main composition, T9/T10, merge, push, or PR edit. Owner bytes were not adapted.

Blocked: yes — the minimal behavior fix touches `internal/store/genesis.go`, which is inside the overall 16-row fence but task-assigned to T2, while the GO requires full battery GREEN before the T1 commit. m-2 also explicitly says FX-CFG-7 and the pinned v5 genesis stay untouched, which requires v8-to-v7 predecessor derivation before the existing v7-to-v6-to-v5 path; simply updating the hash would violate that tripwire.

Next requested action: route a bounded ruling that chooses one coherent checkpoint: (A) authorize the v8-to-v7 predecessor derivation seam in `internal/store/genesis.go` plus the historical s8 fixture-constructor updates as part of T1, preserving the pinned v5 genesis and making the full battery meaningful; or (B) allow the exact owner-byte T1 commit on targeted/package GREEN and move the full fixture battery plus `genesis.go` reconciliation to T2. Do not ask this seat to update FX-CFG-7's expected hash: that would silently replace the pinned-v5 invariant m-2 explicitly preserved. On the ruling, I will resume without changing the owner bytes.
