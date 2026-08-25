## SITREP — WP1 draft battery reaches the required 20/44 census, but closure HALTS on the pre-existing-suite versus nested-dogfood collision

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s16a-impl
PARENT_DISPATCH_ID: s16a-impl
RUN_ID: s16a
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this is a scoped implementation finding for planner/master routing; the operator's terminal MERGE-GATE remains unopened
IN_REPLY_TO: s16a-impl/IMPL-planner-20260824-070538.md
FROM: s16a.implementer
TO: s16a.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-8.planner, m-9.planner, m-7.planner
SUBJECT: WP1 HOLD — 64-row seam draft is compile-clean and derives 20 GREEN/44 RED, but the exact pre-existing-suite command deterministically fails because its nested dogfood suite includes the intentionally RED seam package

Phase: WP1 implementation finding and bounded HOLD. No WP1-close claim, no fidelity-round request, no commit, and no WP2 byte.

Worktree mechanics disclosed: branch `s16a-conformance` lives in linked worktree `/Users/jack/Programming/harness-s16a-conformance`, created from exact authorized LAUNCH_BASE `ff1193d7e63b52a085eeea486055ae9f2c971896`. Current `HEAD` and `git merge-base s16a-conformance ff1193d7e63b52a085eeea486055ae9f2c971896` both return that exact object. Before WP1 bytes, `go test -p=1 -count=1 ./...` was green, including `test/fixtures`.

Draft implementation state, all inside the authorized `frank/test/seam/**` surface:
- one Go package across `helpers_test.go`, `agree_test.go`, `app_worker_test.go`, `provider_worker_test.go`, `connector_app_test.go`, and `shared_test.go`;
- exactly 64 `TestCT_<rowID>` functions: G01-G20, A01-A19, B01-B11, C01-C10, D01/D02/D04/D05; no D03 and no `t.Skip`;
- executable census tool `census.py`, consuming `go test -json`, rejecting missing/extra row IDs, emitting `<rowID> <GREEN|RED>`, and enforcing the launch 20/44 count;
- every test has a row/CT/selector/flag binding comment; gate failures name their CT row; A12/C07/C09 drive the worker runtime, `frank-app` run-start, and existing brokerclient dial;
- no production/source, dependency, module-file, master-doc, WP2/3/4, merge, push, or publication byte.

Script-derived launch census (`go test -json -count=1 ./test/seam/ -run 'TestCT_' | test/seam/census.py`):

```text
G01 GREEN
G02 GREEN
G03 GREEN
G04 GREEN
G05 GREEN
G06 GREEN
G07 GREEN
G08 GREEN
G09 GREEN
G10 GREEN
G11 GREEN
G12 GREEN
G13 GREEN
G14 GREEN
G15 GREEN
G16 GREEN
G17 GREEN
G18 GREEN
G19 GREEN
G20 GREEN
A01 RED
A02 RED
A03 RED
A04 RED
A05 RED
A06 RED
A07 RED
A08 RED
A09 RED
A10 RED
A11 RED
A12 RED
A13 RED
A14 RED
A15 RED
A16 RED
A17 RED
A18 RED
A19 RED
B01 RED
B02 RED
B03 RED
B04 RED
B05 RED
B06 RED
B07 RED
B08 RED
B09 RED
B10 RED
B11 RED
C01 RED
C02 RED
C03 RED
C04 RED
C05 RED
C06 RED
C07 RED
C08 RED
C09 RED
C10 RED
D01 RED
D02 RED
D04 RED
D05 RED
SUMMARY GREEN=20 RED=44 TOTAL=64
```

Verification that passes:
- `go test ./test/seam -run '^$'`: compile-clean.
- `go vet ./test/seam/`: clean.
- `git diff --check`: clean.
- mechanical `TestCT_` count: 64; `TestCT_D03` and `t.Skip`: zero hits.
- final seam census above: exactly the required 20 GREEN / 44 RED with no final row deviation.

Binding observation retained for fidelity review: an early A02 behavioral draft falsely GREENed because a non-hex 32-byte `create_auth_id` was rejected later by the journal identity validator. That did not prove the CT-A02 admission boundary. The current A02 probe includes a valid-nonce allowed control and requires the invalid nonce to fail specifically at worker `invalid assignment`; it is RED at base.

Blocking finding S16A-WP1-F1 — the plan's two required test conditions collide:
1. PLAN r7 §3.3/§3.4 require every gate row to FAIL at launch and forbid skips or silent GREEN coercion.
2. PLAN r7 §3.8 requires the exact pre-existing-suite command `go test -p=1 -count=1 $(go list ./... | grep -v '^github.com/jackli/frank/test/seam$')` to remain green; only the outer package list excludes seam.
3. The pre-existing fixture `TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate` invokes `scripts/dogfood-suite.sh`, whose live bytes export `FRANK_DOGFOOD_NESTED=1` and execute unexcluded `go test -p=1 ./... -count=1`.
4. That nested command includes the seam package; its intentional 44 assertion failures make the dogfood predicate honestly return FAIL. The outer fixture then rejects the production suite record at `test/fixtures/s8_exit_gate_test.go:312` where it requires accepted.
5. Exact-suite reproduction failed once after all other packages reported green. Isolated reproduction `go test -p=1 -count=1 ./test/fixtures -run '^TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate$' -v` failed identically in 99.052 s: `production suite state = "rejected", want accepted`, with `run-suite` outcome `fail`.
6. This is introduced solely by the authorized expected-RED package, not by a production edit: the fresh-base full suite was green before `test/seam/**`; current worktree status is only `?? frank/test/seam/`.

Why no local repair was attempted: excluding `test/seam` from `scripts/dogfood-suite.sh` changes an out-of-scope production script. Returning/short-circuiting the seam package when `FRANK_DOGFOOD_NESTED` is set would suppress the 44 required failures under a real suite invocation and appears to violate the plan's RED-means-FAIL/no-skip/anti-vacuity law. Either course needs an explicit plan/scope ruling.

Done: exact branch/worktree establishment; 64-row compile-clean draft battery; census tool; final script-derived 20/44 census; vet, bijection, D03/no-skip, and diff checks; root-cause reproduction of the pre-existing-suite collision.

Not done: WP1 acceptance criterion §7.2 is not met; no WP1-close or fidelity-round request; no commit (the plan requires the exact pre-existing suite green at every commit); no WP2 byte.

Blocked: S16A-WP1-F1 requires planner/master arbitration of the nested dogfood scope. Requested next action: rule whether the governed nested dogfood suite must exclude the intentionally RED seam package with a corresponding authorized scope change, or provide another exact mechanism that preserves the 44 real RED assertions while making §3.8 coherent. Re-dispatch after the ruling. The pair HOLDS.

ACTIONS_GIT_REF: authorized branch/worktree mechanics plus uncommitted draft bytes only under `frank/test/seam/**`; this SITREP was engine-drafted and submitted; no commit, production edit, merge, push, or publication
FINAL_GIT_STATUS_SHORT: implementation worktree `?? frank/test/seam/`; governing checkout has its inherited INDEX/checkpoint changes plus the current dispatch directory; no source/test change there
