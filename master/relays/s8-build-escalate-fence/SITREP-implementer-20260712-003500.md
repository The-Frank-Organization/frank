## SITREP — T9 production-path RED exposes an unlicensed missing seam: observe activates, but cmd/frank supplies no check evaluator and no executable-claim selection surface exists

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s8-build-escalate-fence-r5
PARENT_DISPATCH_ID: s8-build-t9-relift
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — master fence/schema ruling required under dispatch conditions (b) and the named-seam license; merge remains operator-only
GRILL_REQUIRED: no
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
TASK_ID: T9
IN_REPLY_TO: master/relays/s8-build-fidelity-m7/SITREP-planner-20260712-002500.md
FROM: s8.implementer
TO: master.orchestrator-planner
CC: s8.planner, operator, master.orchestrator-reviewer, m-3.planner, m-2.planner, m-7.planner
SUBJECT: T9 held on production RED — cmd/frank restarts with observe true but passes no Registry.Evaluator into the submit handler; the locked agent-selection triple has no FieldSpec/input home; wiring it requires a new main.go seam and likely a post-T2 schema amendment, so no improvisation

DEVIATION_PRESENT: yes — the production exit leg cannot turn GREEN inside the current named-seam license without either (a) a new executable-claim selection surface or (b) an owner-specified existing mapping that is absent from the code and plan.

What the production-path test proved:
- A real `frank -init` creates the fresh three-member genesis with the amended s8-v1 catalog.
- A real authenticated operator channel submits the governed v5→v6 fieldspec `config_change`; restart succeeds.
- A real authenticated operator channel submits the engine `observe:false→true` `config_change`; restart and phase 0 succeed.
- A real lane then submits a false-done record through `DescribeTools` + `submit` on the production socket.
- RED: state is `accepted`, wanted `rejected` naming the failing predicate.

Root cause verified in code:
- `cmd/frank/main.go` constructs its production handler with `engine.SubmitHandlerWithRender(...)`.
- `SubmitHandlerWithRender` constructs `observe.Env{PresentLayers: env.PresentLayers}` only; `Evaluate` is nil.
- `observe.Gate` therefore returns the default `observe-unavailable`/`blocked` result; for a non-authority record that is the intentional no-vantage `accepted`+`self_reported` row.
- `internal/observe.Registry.Evaluator(selection)` exists and is fixture-covered, but no production call site constructs `observe.Registry` or passes an evaluator.
- The locked design says the lane selects an executable claim as `(claim_ref, check_id, params)`. No FieldSpec row, submit payload field, record header/body contract, or engine-config selector carries that triple in the current tree. `rg` finds `NewRegistry`/`Evaluator` only in observe code and fixtures.

Why this is outside the current license:
- The granted `main.go` seams are limited to three-member init/bless, production `PresentLayers` threading, and T10 boot ordering. Production observe-registry construction/evaluator wiring is a new named seam.
- Any new lane-visible selection field or structured claim surface is a registry/schema change after the completed T2 changeset, triggering dispatch condition (b): escalate to master before work.
- Inventing an implicit mapping from `ACTIONS_GIT_REF`, body text, or current working directory would contradict the locked explicit agent-selection triple and its closed param schemas.

Review findings already folded locally but NOT committed:
- Catalog load now rejects missing nested required keys as well as unknown keys; nested `directory`, `fidelity_review`, channel `ident_calls`, and `conn_receiver_calls` cases are fixture-covered.
- `TestLawPathHygiene` now consumes the owner discovery scan, boundary files, recognizers, site census, output families, sink patterns, and canonical-path rows rather than retaining stale hard-coded scan/recognizer values.
- Those focused tests are GREEN. The production dogfood test remains intentionally RED at `accepted`, so no T9 commit or completion claim was made.

Requested ruling:
1. Name the locked input home and production mapping for `(claim_ref, check_id, params)` if it already exists outside the inspected tree; authorize the exact consumer seam.
2. Otherwise route the minimal selection-surface amendment through m-3 semantics + m-2 FieldSpec ownership + m-7 production-host fidelity, then grant the exact new registry/main.go/test surfaces.
3. Preserve the existing no-vantage row; the fix must provide an actual selected evaluator, not globally reinterpret `Evaluate:nil` as failure.

Verification:
- `go test ./test/fixtures -run '^TestS8FXCFG3' -count=1` — PASS after the nested-shape fold.
- `go test ./test/invariants -run '^TestLawPathHygiene$' -count=1` — PASS after descriptor consumption.
- `go test ./test/fixtures -run '^TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate$' -count=1` — FAIL as required: `production false-done state = "accepted", want rejected`.
- Corrected catalog remains byte-exact at SHA-256 `943f07bb51da3414cf45a16d4bfa00bcee28cc538533fcb7fcd3e8a64b5e209d`.

ACTIONS_GIT_REF: none — T9 work remains uncommitted behind a production-path RED; no merge, push, PR, or live-store mutation
FINAL_GIT_STATUS_SHORT: modified T9 source/test/catalog files plus untracked test/fixtures/s8_exit_gate_test.go; no commit
Next requested action: master rules the missing selection surface and exact fence; s8.implementer resumes from the production RED only after that ruling and the required owner amendments/confirmations.
