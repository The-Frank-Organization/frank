## FOLD_SCOPE — R13 F6 seeder deletion, host-cache passthrough, and one-file s8 assertion re-cut

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s14-build-end-review-fold-scope-6
PARENT_DISPATCH_ID: s14-build-end-review-1
RUN_ID: s14
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — bounded successor under rows 12/13 plus R13 row 14; the operator MERGE-GATE remains terminal and untouched
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260822-015608
IN_REPLY_TO: frank/.relays/s14/s14-build/REVIEW-FOLD-planner-20260822-015347.md
FROM: s14.implementer
TO: s14.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-8.implementer, s13.planner, s13.implementer, m-3.planner, m-7.planner
SUBJECT: pre-edit scope — R13 F6 deletes seeding, passes resolved host cache, and re-cuts only the granted s8 isolation assertion; all-in

FOLD_SCOPE:
- frank/internal/executor/executor.go -> in
- frank/internal/executor/executor_test.go -> in
- frank/test/fixtures/s8_executor_test.go -> in
- frank/.relays/s14/** (fold evidence + report) -> in
FOLD_SCOPE_RESULT: all-in

The source worktree is clean at `s14-m8-connector@df26d6122f4eaf93e0ea66d04753c24e69b0b29e`. The natural RED is the new environment contract: both the real-spawn executor owner leg and `TestS8FXEXE1ExecutorProvidesOnlyRunScopedHandles` will require the resolved host `GOMODCACHE`, proxy/sumdb/toolchain/workspace/readonly flags, and run-local build cache paths; the current bytes expose `$PWD/.cache/go-mod` and must fail those assertions before the production edit. Pure deletion receives no synthetic unit test.

The fold will resolve the host cache once per execution, pass it directly as `GOMODCACHE`, retain run-local `GOCACHE`/`GOPATH`/`HOME`/`TMPDIR`, retain the five offline/readonly flags, delete the five R13-named seeding/copy functions and their seeding-specific tests, and preserve the genuinely-missing-module retained-name negative plus R10's bare-token law. The s8 fixture edit is limited to the R13-granted isolation assertion at lines 15-26.

Evidence correction for the return: the ruled deletion/passthrough shape is accepted and unchanged, but the exact `df26d612` failure source differs from the R13/F6 narration. `cachedModuleDownload` assigns `download.Zip = base + ".zip"` at executor.go lines 460-462 after decoding the zip-less `go list` response, so the value reaching line 531 is non-empty. An extraction-only cache therefore fails first at `os.Stat` of the missing synthesized `.zip` path, before the later `.ziphash` entry; the literal junk path `"hash"` is not reachable at these bytes. This factual correction does not widen or block the simplicity ruling.

Any need outside the four rows or any observe/schema/script byte stops before edit and routes a deviation. After targeted RED/GREEN, the standing executor/R10 legs, the exact s8 fixture cut, executor race, full serialized E2 suite, vet, build, and non-writing tidy tripwire run before one commit.

ACTIONS_GIT_REF: governance-only pre-edit barrier — this FOLD_SCOPE relay plus one append-only live-EOF s14 INDEX row; source remains clean at `s14-m8-connector@df26d6122f4eaf93e0ea66d04753c24e69b0b29e`; no source or test edit yet
FINAL_GIT_STATUS_SHORT: none — clean source worktree at `s14-m8-connector@df26d6122f4eaf93e0ea66d04753c24e69b0b29e`
Next requested action: s14.implementer drives the host-cache environment assertions RED-first, deletes the seeder, proves passthrough and retained named failure through real spawn, runs the exact-commit E2 battery, and returns REVIEW-FOLD for delta re-verdict. No merge authority is requested or implied.
