## SITREP — WP3 IMPLEMENTATION HOLD: the first durability driver reaches a pre-fault frozen-journal/production-reader carrier mismatch; exact RED retained, no out-of-fence translation or repair attempted

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s16-wp3-impl-finding-1
PARENT_DISPATCH_ID: s16-wp3-impl
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this is the plan-mandated production/spec finding UP; disposition and any successor authority belong upstream
GRILL_REQUIRED: no
PLAN_LOCK_ID: s16-wp3-plan-2 @ sha256 735b324ad90ee72a3c62803fa58adb9a026e68ebd65ca5004b838f8e9e690ebd
BUNDLE_ID: s16-integration
IN_REPLY_TO: s16-wp3/IMPL-orchestrator-planner-20260830-012737.md
FROM: s16.implementer
TO: s16.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-7.planner, m-9.planner, m-10.planner
SUBJECT: WP3 HOLD — xit-dur-2 honest authored base fails before its injected marker fault: frozen payload-wrapped run_open is rejected by the production flattened journal reader as missing run_id; same substrate binds xit-edit-1/2; checkpoint 3122d523 is green and pushed, exact RED remains uncommitted

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: no — the finding is an exact frozen-carrier/reader mismatch and exposes no secret or authority byte
- migration/backfill/destructive-write/canonical-data-repair: no — the frozen corpus was read only; no canonical, dist, or store byte was mutated
- money/inventory/orders/planning/accounting/trust-critical-state: yes — the mismatch blocks three trust-critical Step-3 exit records
- AI-or-automation-acts-downstream: yes — these rows would feed the exit packet and operator ratification if they were executable
- worker/scheduler/queue/retry/async-side-effect: yes — the affected reader is the m-9 resume journal path
- cross-repo/service-contract/generated-schema/shared-API-event: yes — frozen m-9 journal bytes disagree with the bound production reader contract
- user-visible-control-with-materializer/downstream-consumer: yes — m-3 evaluation cannot receive a valid xit-dur-2/xit-edit observation from this substrate
- test-runtime-role-mismatch: no — the failure invokes the production `internal/worker/journal` reader directly over the exact frozen bytes before any test-only fault
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: yes — choosing a translation or changing either side is an owner/design decision outside the approved s16 fence
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no downgrade or local absorption is requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none

## Finding S16-WP3-F1 — exact pre-fault incompatibility

The frozen corpus itself remains exact: `STEP-3-EXIT-FIXTURES.json` = `d4580c52675038049471e2fd4ea813c42604b21b0032a9ba5f39fa794f972639`; its authored substrate `master/exit-fixtures/common/edit-base-journal.jsonl` = `feb1bf6cd25ce65469cc116551bce111214f74f4260eef1a7e01dde7b6f7d6db`. The driver reads those bytes from the sibling workspace only after the corpus hash gate succeeds.

The xit-dur-2 frozen input requires an anti-vacuity honest arm BEFORE its single marker mutation: the unmodified authored base must be `resumable` at the honest marker seq `9`; only then may the injected stale `marker_digest` demonstrate `degraded` / `re_derive` to the prior marker. Both xit-edit arms bind the same checksum-bearing authored journal and require their own single-record mutations over that honest substrate.

The frozen genesis line is payload-wrapped: its `run_id`, `run_manifest_digest`, `turn_epoch`, and `create_auth_id` live under top-level `payload`. The production `journal.DecodeRecord` schema expects those kind-specific members flattened at top level. With the exact identity independently read from the frozen genesis payload, `journal.Recover` therefore returns:

`{Disposition:degraded ResumeAction:re_derive Boundary:{Kind:none Seq:0 Offset:0} NextSeq:0 FaultClass:genesis_fault GenesisFault:true Records:[]}`

The direct decode cause is exact: `run_open lacks required member "run_id"`.

Focused reproducer, from `/Users/jack/Programming/harness-s16-integration/frank`:

`go test ./test/exit -run '^TestFrozenEditBaseJournalIsConsumableByProductionReader$' -count=1 -v`

Result: FAIL at `frozen_journal_compatibility_test.go:51` with the recovery and decode cause above. This is before the xit-dur-2 fault injection, so it cannot be scored as the expected degraded arm; the anti-vacuity baseline is red. It also prevents the xit-edit pair from isolating record class as their single independent variable.

## Fence disposition and requested upstream action

No local repair was attempted. `master/exit-fixtures/**` is immutable and OUT. `internal/worker/journal/**` is production source OUTSIDE the three F65-only production rows. A harness translation from payload-wrapped frozen bytes to flattened production bytes would violate the by-reference/non-rederivation law and erase the exact mismatch being tested. Per plan-2 §§2, 3/T2, and 6, implementation is HELD and the finding routes UP through master to the owning m-9 pair for disposition. Any correction requires a separately authorized successor and the corresponding re-review/re-lock consequences; this seat does not choose which side moves.

Requested action: `s16.planner` corroborates and routes the exact mismatch UP; master routes the owner question to m-9 (with m-3 visibility because three exit predicates/locators become unexecutable). The s16 pair remains held on T2 and performs no record run.

## Completed green checkpoint, kept distinct from the RED

Before this driver exposed the mismatch, the in-fence F65 realization and exit evidence floor were committed as `3122d523cc50a06af418b4091dc7045dc75779f7` (`test(exit): realize F65 identity and exit evidence floor`) and pushed exactly to `origin/s16-integration`. Its 16-file set is confined to `cmd/frank/**`, `internal/config/**`, `internal/store/**`, and new `test/exit/**`; zero seam, script, dist, master, or other production byte moved.

Fresh pre-commit floor for `3122d523`: seam census `GREEN=64 RED=0 TOTAL=64`; `go vet ./...` exit 0; `go test ./... -count=1` exit 0, including `test/fixtures` in `251.682s`; `git diff --check` exit 0. The actual production dogfood gate also passed its focused two-adjudication run in `210.868s`. Draft PR #2 is DRAFT/OPEN at exact head `3122d523`, has no checks, and its body now records this checkpoint and hold.

The checkpoint includes the F65 serve stamp/config-generation/composite realization but does NOT claim the required m-7 chartered review has returned. That review remains a downstream prerequisite before any close could cite the conductor bytes. The exact failing driver is deliberately retained UNCOMMITTED so the green checkpoint and the new RED are not conflated.

## Honest not-done

No xit-dur-2, xit-edit-1, or xit-edit-2 verdict; no eight-leg completion; no record run; no observer capture set; no candidate E3 record; no post-capture m-3 formal ruling; no E3 admissibility or exit-closure claim; no merge, deployment, publication, release, or Step-3 ratification. H-12 remains unchanged. The bound dist was read only and remains untouched.

ACTIONS_GIT_REF: implemented and pushed only the in-fence green checkpoint 3122d523 on s16-integration; refreshed draft PR #2 at that exact head; added one uncommitted test/exit RED driver that reads the hash-correct frozen authored journal and invokes the production reader; made no out-of-fence production, frozen-corpus, seam, script, dist, master, merge, deployment, release, or CI/CD mutation
FINAL_GIT_STATUS_SHORT: (literal implementation-worktree status immediately before drafting this report; branch equals origin at 3122d523 and the retained RED is uncommitted:)
?? test/exit/frozen_journal_compatibility_test.go
