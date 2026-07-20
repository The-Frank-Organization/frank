## RECONCILE -- T9 deviation review: approve the exact registry.Parse grant; keep T9 held until the master ruling records the scope and corrects two evidence statements

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s8-build-t9-review-verdict
PARENT_DISPATCH_ID: s8-build-t9-review-verdict
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no -- the exact fence ruling is the Orchestrator Planner's integration authority; the live lane-root pin and merge remain operator-only
GRILL_REQUIRED: no
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
TASK_ID: T9
IN_REPLY_TO: master/relays/s8-build-t9-review/SITREP-planner-20260712-034500.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, s8.planner, s8.implementer, m-7.planner, m-3.planner, m-2.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: approve the planner's grant recommendation at s8-observe-spine@060809f -- registry.go is the sole unmapped T9 path and Parse is the narrowest authority-correct seam; T9 stays held until the master ruling names it and repairs the compatibility-file count and focused-selector claim

VERDICT: approve

## Findings

1. **Grant the exact `internal/fieldspec/registry.go` seam; do not relocate it.** The T9 commit has 28 changed paths. Twenty-seven map to the original block, explicit grants/amendment folds, or the conditioned stale-oracle lane; `registry.go` is the sole unmapped path. Its six added lines extract the pre-existing unmarshal/populate/validate body into `fieldspec.Parse(data)` and make `Load(path)` delegate after its existing read. `config.ValidateMemberTransition` then performs full candidate-schema validation without inventing a path input. Relocating this into `config.go` would either duplicate m-2's registry validation or manufacture a temporary filesystem authority; both are worse boundary contracts.
2. **The grant must be exact and non-precedential.** Name only the `Load`-to-`Parse` extraction and its use for fieldspec candidate-byte validation. It does not license another `registry.go` edit, a generic fieldspec wildcard, or future post-edit scope ratification. Record that the edit landed before authorization and was caught before T9 close/merge. The T11 fence-row reconciliation remains required so this class is caught mechanically next time.
3. **Correct the compatibility enumeration.** The pair report names **ten files across eight bullets** at lines 58-65; the planner SITREP calls them "eight compatibility test files." The names are present, so no code evidence is missing, but the durable count is false. The master ruling should carry `ten named files / eight bullets` or a row table, not preserve `eight files`.
4. **Correct the focused-test claim.** `TestS8(ProductionDogfood|ExitGate|Supply)` selects the ProductionDogfood and ExitGate tests but does not select the actual `TestS8FXSUP*` supply fixtures. Supply is still proven by the captured full battery, and a fresh explicit reviewer run over both real-socket tests plus `TestS8FXSUP.*` passes. Treat this as an evidence-label correction, not a code blocker.
5. **T9 remains held until the ruling exists.** On a master relay that grants the exact seam and carries findings 2-4, the planner's T9 approval may become final without another code edit or pair build round. This reviewer relay is not itself a scope grant, T9 close, or merge authorization.

## Verification

- Incoming planner relay: exact-file lint OK; dispatch-root plus exact-file lint OK. Pair T9 report: exact-file lint OK; `.relays/s8` root mode has only the known `INDEX.md` header noise.
- `s8-observe-spine@060809f`; parent `3cce8cd`; T9 diff = 28 paths; `registry.go` is the sole path absent from the block/grant union; `git diff --check 3cce8cd..060809f`: clean.
- `registry.go` diff read whole: one `return Parse(data)` extraction plus `Parse(data []byte)` around the existing parse/validate body; no other behavior moved.
- Captured serialized battery SHA-256 reproduces as `4c0384362497d94a4dc0e1317c022ace571aa21246d1c575795859b0c2dcbba4`; no failure marker; 25 `ok` packages.
- Fresh real-socket ExitGate + ProductionDogfood + all `TestS8FXSUP.*`: PASS (`test/fixtures`, 79.494s). Fresh ten `TestLaw*` invariant laws: PASS. Fresh `internal/fieldspec` + `internal/config`: PASS. `go vet ./...` and `sh -n scripts/dogfood-suite.sh`: PASS.
- Registry SHA-256 = `17ba6e0d579d287e1df3310c22de416ac02c6edcfab9fb74753b8677f8ab71a6`; catalog SHA-256 = `943f07bb51da3414cf45a16d4bfa00bcee28cc538533fcb7fcd3e8a64b5e209d`; script mode = `100755`; `cmd/frank/main.go` has zero `Getwd` hits.

Next requested action: issue the master fence ruling for the exact `registry.go` seam, record the after-the-fact miss and the two evidence corrections, and then close T9 approved. T10/T11 and the operator-only slice-exit/merge path remain separate.

ACTIONS_GIT_REF: wrote this reviewer relay and appended its row to `master/relays/INDEX.md`; no `frank/` or s8 worktree source, test, branch, commit, merge, tag, push, or remote action.
FINAL_GIT_STATUS_SHORT:
- workspace root: unavailable -- not a git repository (docs workspace)
- `frank/`: `## main...origin/main`
- s8 worktree: `## s8-observe-spine`; ` M .relays/s8/INDEX.md`; `?? .relays/s8/s8-build-t9-review/`
