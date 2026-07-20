## RECONCILE -- s7b visibility review: approve the m-7 planner-first lane; pin mechanical acceptance for the flake branch and keep the gate-lift claim scoped to OI-S7A-CLOSE-ONCE-RACE

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s7b-dispatch
PARENT_DISPATCH_ID: s7b-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no -- visibility review only; s7b implementation and merge retain their existing gates
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7b-dispatch/PLAN-orchestrator-planner-20260710-193159.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-7.planner, m-7.implementer, m-1.planner, m-1.implementer, m-2.planner
SUBJECT: approve s7b routing under B10 third application -- production race, ownership, branch base, red-first requirement, delegation, m-1 fidelity, and separate merge gate are sound; two acceptance clarifications must ride the pair PLAN but do not require dispatch reissue

VERDICT: approve

## Review

1. **Authority and ownership are sound.** B10's third operator-approved application names the m-7 pair and this exact client-lifecycle race (`PROTOCOL-DEVIATIONS.md:117-122`). `TO: m-7.planner`, Implementer plan review, mechanical all-in scope diff, unique child lineage, and delegated dispatch are the correct planner-first shape.
2. **The production defect remains live at the pinned base.** `main@2e1b4f0` and peeled `s7-close` agree. A fresh `go test -race -count=20 ./cmd/frank-mcp -run '^TestShimReconnectsAndRetriesSingleCallAfterConnectionLoss$'` fails with `panic: close of closed channel` at `internal/channel/server.go:523`; source still has independent unsynchronized closes in `Client.Close` and `readLoop` (`:519-525,555-562`). The red-first requirement is real and correctly blocks live s8 MCP dogfood.
3. **Work-item separation is preserved.** The plan does not silently merge `FLAKE-SOCKET-PAR` into the runtime race. A shared root may close both only with evidence; otherwise they retain distinct mechanisms and acceptance.
4. **Scope and downstream gates are proportionate.** Channel lifecycle code/tests plus the two named fixture surfaces/shared helpers are bounded; m-1 close-ordering fidelity, VP integration, and operator merge remain explicit.

## Conditions For The Pair PLAN

- **A rationale alone cannot close `FLAKE-SOCKET-PAR`.** If the pair fixes fixture startup, require focused repeated runs of both named fixtures plus three consecutive parallel full-suite runs. If it chooses the kickoff-licensed parallelism pin, the pin must be mechanical and checked in (a canonical executable test target/script/CI entry that forces the selected parallelism), then exercised successfully; the report must state whether the underlying fixture race remains registered. The existing serialized full battery for work item 1 is necessary but does not by itself prove the pin is the gate consumers invoke.
- **Scope the lift precisely.** A successful s7b merge lifts `OI-S7A-CLOSE-ONCE-RACE` as the live-channel blocker. It does not by itself declare s8 or dogfood globally ready; the s8 design, genesis/config, implementation, and operator gates still apply.
- `IN_REPLY_TO` points to the earlier s7 revision rather than the final s7 approval/merge trail. It is display-only and does not invalidate authority, but future s7b returns should cite the current base/close chain.

No new planner approval gate is created by this relay. The addressed m-7 Planner may proceed and must carry the two acceptance clarifications into its PLAN and the Implementer's PLAN-REVIEW.

## Verification

- Incoming exact-file lint: OK; `--relay-root master/relays/s7b-dispatch` plus exact file: OK.
- `frank/main` = `2e1b4f036c3cfb66902655dcbc3c6702eb1c4023`; parents = `54420dbc` + `5e6bf83`; `s7-close` points to `2e1b4f0`.
- Fresh focused race command at current main: exit 1 with double-close panic at `Client.Close:523`.
- Both flake fixtures resolve to `test/fixtures/s6_iph_test.go` and `test/fixtures/s4_config_change_test.go`, matching the declared fence.
- `s7b-close-once` does not yet exist; no implementation has started from this review.

Next requested action: m-7.planner writes the s7b pair PLAN with the conditions above; normal Implementer plan review and delegated-dispatch lineage follow. No implementation or merge authority is granted here.

ACTIONS_GIT_REF: wrote this visibility-review relay and appended its row to `master/relays/INDEX.md`; no `frank/` source, test, branch, commit, merge, tag, or remote action.
FINAL_GIT_STATUS_SHORT:
- workspace root: unavailable -- not a git repository (docs workspace)
- `frank/`: `## main...origin/main`

