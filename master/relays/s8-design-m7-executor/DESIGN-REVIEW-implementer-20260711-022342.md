## DESIGN-REVIEW - s8 executor host r1 must revise before lock

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-design-m7-executor-review-r1
PARENT_DISPATCH_ID: s8-design-m7-executor
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - the operator still owns the OS-sandbox election; technical review cannot ratify that risk posture
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m7-executor
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s8-executor-grill-r1
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m7-executor/DESIGN-planner-20260711-015835.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner
SUBJECT: must revise - the host conflicts with the approved m-3 taxonomy/verdict contract, overclaims writable confinement, treats deterministic identity as exactly-once execution, and does not prove descendant death before cleanup

DESIGN_REVIEW_VERDICT: must-revise

The F4 top-level claim, deny-by-default environment, capture-ceiling disclosure, outer-gate fault disposition, and no-policy-in-executor boundary are directionally correct. Five cross-contract findings block lock.

## Findings

### F1 - The design contradicts the approved m-3 side-effecting taxonomy

Executor section 2.10 says "no side-effecting check class exists" and a future m-3 registration would require a new round (`2026-07-11-s8-executor-host.md:28`). The approved m-3 design already defines `class: base | suite | side_effecting`, default-denies side-effecting entries, and gives them a static engine-config allowlist (`2026-07-10-s8-check-registry-probe-design.md:32-46,81-83`).

Required fold: state that the **v1 executor accepts only suite-class read-only runs and refuses side-effecting entries**. Preserve m-3's existing side-effecting taxonomy/operator policy outside this host; do not claim the class does not exist and do not add operator-gate plumbing here.

### F2 - "Immutable/read-only inputs" and "sole writable surface" are not established under the admitted same-uid residual

Section 2.1 calls the input set immutable, section 2.2 calls staged inputs read-only and the workdir the only writable surface, then permits a writable shared per-executor cache (`design:19-20`). File modes do not make copies immutable to a same-uid process that owns the workdir, and a shared writable cache is a second provided writable surface plus a cross-run contamination channel. This wording exceeds the F4 ceiling in section 1.

Required fold: separate identity from enforcement. The manifest's input digests may immutably identify the **pre-spawn bytes**, while staged copies remain mutable ambiently unless sandboxed. For v1, remove the shared writable cache, make any cache run-scoped under the ephemeral workdir, or fully specify content-addressing/reset and include its inputs in the manifest. Replace "only writable" and "writes confined" claims with provided-surface wording that the NF can actually prove.

### F3 - A deterministic content hash does not guarantee one execution

Section 2.11 says a content-hash key causes crash-refire/duplicate dispatch to coalesce "instead of double-running", and FX-EXE-3 requires run count 1 (`design:29,33`). The cited A-2 rule says deterministic keys let timer-fired commands dedupe through governed intake; it does not make an already-spawned external process exactly-once. A crash after spawn but before atomic verdict commit can rerun a read-only suite after recovery.

Required fold: either design a durable pre-spawn reservation/result protocol with explicit crash windows and recovery semantics, or narrow the claim to one committed verdict/one pivot while admitting at-least-once suite execution under crash. The lean fit is the latter because v1 checks are read-only. FX-EXE-3 must test the selected honest property, not `run count 1` without machinery.

### F4 - Reaping the direct child does not prove descendants are gone

Section 2.7 kills the process group, waits for the child, signals the group again, and then permits workdir removal; FX-EXE-2 requires no orphan (`design:25,33`). `wait` confirms only the direct child. A descendant can survive or escape, and the design does not define group-death observation, bounded polling, subreaper behavior, or the fault path if the group remains live.

Required fold: define the post-kill state machine: signal group, reap direct child, verify process-group nonexistence (or the platform-equivalent) before cleanup/verdict, bound that verification, and route a survivor to the typed no-verdict machinery-fault edge. State the supported platform assumptions; do not claim no orphan from child reap alone.

### F5 - The returned verdict expands m-3's approved shape and can bypass its hygiene contract

Section 2.11 returns a manifest containing `exit class`, `output ref`, and `duration` verbatim (`design:29`), while the approved m-3 boundary consumes `spawn(...) -> CheckVerdict` with its closed fields and conductor-side path-free/bounded redaction (`m-3 design:107-130,134-144`). An `output ref` is not in that agreed shape and can itself expose a host path; raw stdout/stderr are host-internal evidence, not automatically m-3 verdict fields.

Required fold: distinguish the host-internal run result/manifest from the m-3 `CheckVerdict`. Return only the reconciled closed verdict shape; keep capture references internal or use a symbolic bounded identifier explicitly admitted by m-3. Preserve m-3's conductor-side redaction before any verdict feeds a seat-deliverable field.

## Operator Gate

The GRILL_LOCK correctly leaves OS sandboxing operator-owned (`design:65-66`), but approval by this seat cannot ratify the v1 residual. Keep the current v1 handle-surface design as the recommendation; require an operator record before the reconciled DESIGN_LOCK selects it or starts a sandbox design round.

## Confirmed Non-Blockers

- F4 wording at section 1 correctly distinguishes provided-handle absence from ambient same-uid authority.
- Environment allowlisting, closed stdin, captured stdout/stderr, and marked truncation are appropriate host-owned controls once the writable/cache claims are narrowed.
- Timeout/machinery failure correctly routes to the outer NF-S6 two-axis disposition; the executor adds no terminal authority.
- No c4 amendment is needed if the repaired host remains inside the locked step-4 hook.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, no lock, and no operator choice inferred.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `691d034`
Next requested action: m-7.planner folds F1-F5 into executor-host r2 and returns a new DESIGN relay for re-review; master holds m-3/m-7 reconciliation and lock consumption meanwhile.
