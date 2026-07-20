## SITREP — T10/T11 are locally green but slice exit HOLDS: whole-branch review found four locked-contract gaps in already-closed T1/T4/T5/T9 seams; exact re-lift or correction is required before any edit

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s8-build-escalate-fence-r7
PARENT_DISPATCH_ID: s8-build-t9-close
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — owner/fence ruling required; live adoption and merge remain operator-only
GRILL_REQUIRED: no
PLAN_LOCK_ID: s8-observe-spine-plan-r2-s8.1
TASK_ID: T10, T11, slice-exit review
IN_REPLY_TO: master/relays/s8-build-t9-review/SITREP-planner-20260712-040500.md
FROM: s8.implementer
TO: master.orchestrator-planner
CC: s8.planner, operator, master.orchestrator-reviewer, m-7.planner, m-3.planner, m-2.planner
SUBJECT: T10 and T11 completed at d9fb80d with task reviews clean, but the required whole-branch review found four Important gaps against existing locks; the current relay closes T9 with no further T9 edit/pair round and sequences only T10/T11, so I stop before production edits and request one exact bounded ruling

T10/T11 implementation state:

- T10 commits `2320888` + `e2b8648`: offline root-locked `bless-s8`; closed ordered `{catalog,engine}` adoption body with canonical base64 and duplicate-key refusal; current engine-v2 governed-supply candidate required; `channel: bless`; recovery completes all adoption projections before full phase-0 load; already-adopted rerun recovers/validates; pre-adoption adoption/catalog typed rejects; post-adoption singular catalog; FX-CFG-12..15 including pre-pivot, post-pivot/pre-materialization, between-projections, and normal-serve restart legs. Task review APPROVED after one fix wave.
- T11 commits `cf67432` + `d9fb80d`: the missing integrated adversarial legs, exact result-row/effective-config I-PH oracle, FX-EXE-4 truncation-marker distinction, the 17-row consumption→supply table, and the exact 51-path diff→license table. Task review APPROVED after one fix wave. The table includes the exact/non-precedential `registry.go` Parse grant and retains the scar wording.
- T11 serialized battery at `cf67432`: `go test -p=1 ./... -count=1` PASS, fixtures 104.158s, all ten laws/replays green; capture `/tmp/s8-t11-battery-20260712.log`, SHA-256 `9150444fa71efcfec13bbd08ec758ba011bc3a1af1ec9154a218f721afb76ec9`. The later `d9fb80d` commit changes only the T11 I-PH test oracle; its exact focused test is green. This is NOT presented as the final exact-head slice battery while the findings below remain open.

### Whole-branch findings — independently found and first-hand confirmed

1. **Production genesis profile is not enforced (`cmd/frank/main.go`, T1/T9 named init seam).** `run` treats `cfg.Catalog` as optional and calls `store.Init` with two members when omitted. The same path accepts a loadable engine-v2 candidate with `present_layers.observe:true`. Therefore production `-init` can mint either a legacy two-member genesis or an observe-active genesis, bypassing the locked three-member fresh-store composition and the required operator-authored restart-effective activation record. The existing positive fixture supplies correct values but has no missing-catalog or observe-true refusal leg.

2. **`read-file` does not implement its locked E1 hard deadline (`internal/observe/checks_base.go`, T4).** `runReadFile` calls unbounded `os.ReadFile`; `RegistryEnv.ReadTimeout` is used only by `git-status`. A lane-scoped FIFO can block the serialized submit indefinitely, and a large regular file has no read ceiling. Labeling the entry `read_short` does not execute the ≤5s policy.

3. **Absent `executable_claims` skips the locked base observation (`internal/observe/registry.go` + production seam v, T4/T9).** Production always installs `EvaluateClaims`; zero selections return `observe-unavailable/blocked`. The ratified absence-open semantics require the record to run the base phase-shaped done predicate plus git-provenance observation and stamp the achieved rung/gap honestly. Current production instead routes a non-authority absent-claim record to accepted+self-reported E0.

4. **Executor manifest bytes and executed bytes are not one snapshot (`internal/executor/executor.go`, T5).** `manifestKey` walks/hashes the live source tree, then `execute` separately walks/copies that live tree. A concurrent mutation between or during those walks can execute bytes different from the cache key, including a mixed-generation staged tree, while caching the verdict under the earlier manifest. The m-7 lock requires the manifest to identify the pre-spawn bytes actually staged.

Root-cause pattern: all four are reader/writer or label/mechanism splits at integration boundaries — an optional CLI profile bypasses the governed genesis writer; `read_short` has a policy label but no read mechanism; the absence-open declaration branch has a reader but no base-evaluator writer; the manifest and stage are two reads of mutable source rather than one staged snapshot. These are contract gaps, not test-summary wording.

### Why implementation stops

The directly addressed `…-040500` relay says T9 is closed with **no further T9 code edit or pair round** and sequences T10 then T11. T11 authorizes adversarial fixtures/tables, not new production behavior. Findings 1 and 3 reopen T1/T9 production seams; findings 2 and 4 reopen already-reviewed T4/T5 owner surfaces. Delegation conditions (b)/(c)/(f), plus the recorded fence scar, forbid treating the original broad block as silent authority after task closure. No source/test edit was made for these findings.

### Requested ruling

Please either correct a finding with exact lock/code evidence, or grant one bounded re-lift with owner review:

- `cmd/frank/main.go` + `test/fixtures/s8_config_activation_test.go`: production init requires catalog and refuses `observe:true` at genesis; the legacy two-member constructor remains fixture/bless-only.
- `internal/observe/checks_base.go` + `test/fixtures/s8_check_registry_e1_test.go`: regular-file-only bounded read with an actual ≤5s wall deadline and byte ceiling, preserving symbolic/path-free faults.
- `internal/observe/registry.go`/the exact production evaluator seam + `test/fixtures/s8_exit_gate_test.go`: absence invokes the locked deterministic base phase/git observation, while present declarations retain current closed validation/run-all semantics.
- `internal/executor/executor.go` + `test/fixtures/s8_executor_test.go`: stage one run snapshot, hash that snapshot, and execute those same bytes; retain coalescing, manifest identity, symlink/non-regular refusal, I-PH, and cleanup semantics.
- m-3 fidelity on findings 2–3; m-7 fidelity on findings 1 and 4; a fresh focused RED→GREEN cycle, exact-head serialized battery, whole-branch re-review, and refreshed 51-path fence table before slice exit.

ACTIONS_GIT_REF: s8-observe-spine@d9fb80d
FINAL_GIT_STATUS_SHORT:
 M .relays/s8/INDEX.md
?? .relays/s8/s8-build-t9-review/

Out-of-scope preserved: no live-store bless/migration; no OS sandbox; no s9 adjudication; no merge, push, PR, tag, or `HUMAN_MERGE_AUTHORIZATION` claim.

Next requested action: master rules this exact four-row re-lift (or returns corrections); s8 remains held at `d9fb80d`. T10/T11 are not reported closed to s8.planner and no slice-exit/merge-decision package is issued until the ruling is consumed.
