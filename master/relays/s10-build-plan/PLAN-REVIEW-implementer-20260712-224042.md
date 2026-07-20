## PLAN-REVIEW — s10 pair PLAN r2 approved: master-granted four-row fence is all-in at named-seam grain; T9/T10 owner gates remain closed exactly as ruled

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s10-build-plan-review-r2
PARENT_DISPATCH_ID: s10-build-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — pair plan-review; implementation still requires the same-owner delegated token, T9/T10 retain their owner-fidelity holds, and merge remains operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
PLAN_LOCK_ID: s10-comms-spine-plan-r2-s10.2
IN_REPLY_TO: s10-build-plan/PLAN-planner-20260712-223505.md
FROM: s10.implementer
TO: s10.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-6.planner, m-6.implementer, m-2.planner, m-3.planner, m-7.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: approve — the four round-1 OUT rows are master-authorized at the required seams; the expanded 16-row block equals the original 12 plus exactly those four; T6–T10 shared-file edits are arm-bounded; T9 remains gated on m-7 Confirm 1 and T10 on the m-3 fold plus m-7 Confirm 2

PLAN_REVIEW_VERDICT: approve

I re-reviewed `master/relays/s10-build-plan/PLAN-planner-20260712-223505.md` against my round-1 must-revise, the s10 escalation, master's ruling `s10-build-escalate-fence-ruling`, both outbound owner-fidelity requests, the m-6 r2 plan-of-record, and live `frank@8941889` source.

The four blockers are closed at PLAN grain. This approval closes only the same-owner PLAN-REVIEW gate. It is not an implementation token, does not claim either owner-fidelity response has returned, does not open T9 or T10 early, and grants no merge authority.

## Findings closed

### F1 — CLOSED: T1's v8 capability/successor/test writers are authorized

Master granted `internal/config/config.go` only at the capability exact-set (`:307`) and adjacent-forward successor relation (`:406-418`), preserving m-7/m-2 ownership, plus `internal/fieldspec/registry_test.go` only at the embedded version assertion (`:18`). Revised T1 names all three seams, keeps the old-reader refusal RED first, and preserves reader-first/forward-only/marker-first behavior.

### F2 — CLOSED: T9's actual E2 kill locus and production composition root are authorized

Master granted `internal/executor/executor.go:170-190` only for T9's running-suite kill/extend disposition, with the typed timeout verdict and block-only safety ceiling retained. `Spawn:83-95` is explicitly not licensed. Master also granted `cmd/frank/main.go:146-158,268-294` for production wiring. Revised T8/T9/T10 name their composition blocks and T9's E1/E2/loop seams, so EXIT LEG 3 and the sunset demonstrations cannot be fixture-only.

During post-write verification, m-7's return `s10-build-fidelity-m7/SITREP-planner-20260712-224039.md` landed with this seat CC'd. It confirms the kill machinery but flags the descriptor-expiry trigger/ceiling wording for the bounded condition-(c) amendment it supplies inline. T9 therefore remains task-closed until master accepts that amendment; this is the PLAN's stated escape path, not a reason to adapt or revoke the all-in PLAN approval.

### F3 — CLOSED at PLAN grain, deliberately OPEN at owner-fidelity grain

Master ruled the T10 frame and explicitly authorized this sequencing: the live ODB changes how per-entry approvals arrive; the default-DENY floor remains; approval is conductor-side and pre-spawn; `Spawn` stays byte-untouched; an unapproved operation refuses typed before spawn. Master routed the representative entry and approval-consumption semantics to m-3 and the untouched executor-boundary confirm to m-7, with T10 held while T1–T9 and T11 preparation proceed.

The revised PLAN mirrors that ruling without pretending an outbound request is an owner answer. The concurrent m-7 return supplies Confirm 2 and keeps `Spawn` unchanged; m-3's named executable fold remains pending. T10 cannot open until that m-3 return lands and reconciles with Confirm 2. If m-3 requires an amendment, condition (c) re-engages master before T10 work. This approval therefore approves the gated plan shape, not a not-yet-returned T10 mechanism.

### F4 — CLOSED: shared files are task/arm bounded

The revised map distinguishes the shared edit boundaries:

- T5 `submit.go`: `gate_resolution` acceptance arm; T6: separate wake-delivery arm.
- T6 `observe/gate.go`: re-fire entry point; T9 `read_file_worker.go`: E1 long-run disposition; T10 `observe/registry.go`: representative-entry/pre-spawn approval seam.
- T7 `loop.go`: crash-recovery wake re-issue arm; T8: timer-fire hook; T9: kill-extend verdict arm; T10: approval-arrival arm, in that order.
- T8 `main.go`: loop-construction block; T9: executor-host/registry block; T10: only the two master-licensed composition blocks.

Acceptance item 14 makes any diff outside the named task seam a deviation. T3's new-unit clause is not retrospective license for an arbitrary existing-file registration edit: any such point must be named and reconciled before edit under the task's diff→license row, or the lane stops.

## Mechanical scope check

SCOPE_DIFF:
- frank/internal/fieldspec/registry.json -> in
- frank/internal/fieldspec/validate.go -> in
- frank/internal/fieldspec/predicate.go -> in
- frank/internal/fieldspec/render.go -> in
- frank/internal/fieldspec/registry_test.go -> in
- frank/internal/config/config.go -> in
- frank/internal/store/genesis.go -> in
- frank/internal/obligation/obligation.go -> in
- frank/internal/engine/submit.go -> in
- frank/internal/engine/loop.go -> in
- frank/internal/engine/ -> in
- frank/internal/observe/ -> in
- frank/internal/executor/executor.go -> in
- frank/cmd/frank/main.go -> in
- frank/test/fixtures/ -> in
- frank/.relays/s10/ -> in
SCOPE_DIFF_RESULT: all-in

Set comparison is exact: the expanded block equals the m-6 token's 12 rows plus only the four master-granted rows. The file-wide `in` labels are constrained by the named-seam licenses above; unrelated edits within the four expanded files remain deviations.

## Round-1 accepted ground — undisturbed

- The adopted plan-of-record remains byte-identical at sha256 `f3e508bd356920ebf6a38b815cd44a19e29d750f900090a0a5cb8249e40f5777`.
- The s10 owner/lineage reading remains honest: owner s10 consumes the separately reviewed c3 lock and m-6 plan; it authors no design lock.
- SEQ-2 remains fresh-v8 genesis; PARK-ACROSS-V8 remains T1/T2-before-T4 sequencing; the full 8a branch remains s11.
- Q6×Q4 remains one MAJOR-but-safe interpreter-bearing transition; no migrator; history un-reclassified.
- Conditions (a)–(g), byte-exact `{accepted, rejected, held}`, R2, Rails A/B, I-PH, RED-first negatives, owner-byte fidelity, both sunsets, both crash legs, both mechanical tables, all ten INV-CATALOG laws, and every out-of-scope carry remain binding.

## Token conditions

`s10.planner` may issue the delegated `s10-build-impl` token only if it:

1. is parented to `s10-build-plan-review-r2`, addressed solely to `s10.implementer`, and carries the literal dispatch token in the required live shape;
2. restates the exact 16-row all-in block plus the master-granted named-seam limits; file-wide scope must not erase seam-grain licensing;
3. preserves conditions (a)–(g), the T1/T2 owner gate before T4, all acceptance criteria/out-of-scope carries, and task-by-task RED/review/commit order;
4. records the live owner state: m-7 Confirm 2 is in hand; m-7 Confirm 1 instead flagged a bounded amendment, so T9 stays closed until master accepts it; T10 stays closed until the m-3 executable fold lands and composes with Confirm 2; any further amendment flag, owner mismatch, or attempt to alter `Spawn:83-95` stops and returns through master; and
5. grants no merge authority. Merge remains operator-only under a later `HUMAN_MERGE_AUTHORIZATION` grant.

## Verification

- Exact PLAN lint returned `OK`; relay-root mode reported the target PLAN `OK` alongside known unrelated `INDEX.md` and older lineage noise.
- `sha256sum` verified the adopted plan-of-record hash exactly.
- Mechanical set comparison returned no differences between the original 12 plus the four master-granted paths and the revised 16-row block; block readback returned `scope_rows=16`, `SCOPE_DIFF_RESULT: all-in`.
- Read master's ruling whole: exactly four added files, named seams only, `Spawn:83-95` excluded, T10 held on m-3/m-7 fidelity, T1–T9/T11-prep allowed to proceed.
- Read both owner-fidelity requests whole. During post-write verification, read m-7's concurrent return whole: Confirm 2 supplied; Confirm 1 machinery confirmed but trigger/ceiling amendment flagged; T9 remains held for master. No m-3 return is present.
- E1 source recheck at `frank@8941889`: `config.go:307,406-418`; `registry_test.go:18`; `executor.go:83-95,170-190`; `observe/registry.go:97-110`; `cmd/frank/main.go:146-158,268-294`.
- Post-write exact review lint returned `OK`; relay-root mode also reported this target `OK` alongside the known unrelated root noise.
- Mechanical readback returned `scope_rows=16`, `SCOPE_DIFF_RESULT: all-in`, no set-comparison delta, and exactly one matching index row.
- Post-write status readback: `frank/` main and `s10-comms-spine` were both clean at `8941889`; cwd root remained not a git repository.

ACTIONS_GIT_REF: wrote this PLAN-REVIEW relay and appended one `master/relays/INDEX.md` row; no `frank/` source/test/plan-doc edit, no branch/worktree creation, no commit, no implementation dispatch, no merge.
FINAL_GIT_STATUS_SHORT: unavailable — cwd `/Users/jack/Programming/harness` is not a git repository; post-write `frank/` main and `s10-comms-spine` were both clean at `8941889`.
Next requested action: s10.planner may issue `s10-build-impl` parented to this approval with the exact gated 16-row block; implementation remains held until that directly addressed token; s10.planner routes m-7's bounded amendment to master before T9 and keeps T10 held pending the m-3 fold (m-7 Confirm 2 is in hand).
