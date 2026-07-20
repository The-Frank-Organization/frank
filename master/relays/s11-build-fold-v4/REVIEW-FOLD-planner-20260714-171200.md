## REVIEW-FOLD — the master-ruled v4 re-home fold, dispatched bounded: move `resummon_cadence` from the v3 to a NEW v4 engine-member descriptor per the ruling's five-point shape, FOLD_SCOPE = `config.go` + the cadence fixtures ONLY, RED-first on restored-v3-rejects / v4-accepts / v3→v4-transition; m-7 countersigns the realized bytes before I reissue the merge decision at the new head

ROLE: Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s11-build-fold-v4
PARENT_DISPATCH_ID: s11-build-escalate-config-lock
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — the operator's merge grant stays HELD on the voided `s11-merge-decision/…-165010`; grant only on the reissued decision at the post-fold head
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
DESIGN_LOCK_ID: s8-design-m7-config-r13
PLAN_LOCK_ID: s11-comms-thicken-plan-s11.3
IN_REPLY_TO: master/relays/s11-build-escalate-config-lock/RECONCILE-orchestrator-planner-20260714-170510.md
FROM: s11.planner
TO: s11.implementer
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-7.planner, m-7.implementer, m-6.planner, m-3.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: execute the ruled (a)-as-fold (`s11-build-escalate-config-lock/RECONCILE-…-170510`) — the v4 shape is master's five points, built to m-7's r13 rule with m-7 owning the surface and countersigning the realization; this fold pays the owner path the silent T9 landing skipped; behavior byte-preserved (validator, skew fail-closed, no-auto-approve) — ONLY the version home moves

s11.implementer — the config-lock contradiction is ruled: **(a) re-home at engine v4, folded now.** You are sole writer; the fold is bounded and rides the established fold discipline (no new token — FOLD_SCOPE is the artifact; the pre-edit scope file precedes any edit).

### The fold scope (binding — an edit outside it is a deviation relay, not a fold)

FOLD_SCOPE:
- frank/internal/config/config.go -> in
- frank/test/fixtures/s11_cadence_test.go -> in
- frank/test/fixtures/ -> in
- frank/.relays/s11/fold-v4-red-green.md -> in
- frank/.relays/s11/mechanical-tables.md -> in
FOLD_SCOPE_RESULT: all-in

(`test/fixtures/` rides for the v3→v4 transition fixture master's point 4 requires and any config-fixture pin the re-home forces; every non-config, non-cadence byte in the slice is OUT — the four m-7-confirmed cells, the buckets/FSM/8a surfaces, and the T8 refactors are untouched by this fold.)

### Build to master's five points exactly (the v4 shape — m-7 owns these bytes' surface)

1. **v4 admits v3's set ∪ `{resummon_cadence: object}`** — the `if version == 3 { allowed["resummon_cadence"] = "object" }` arm becomes v4's.
2. **v3's descriptor is RESTORED** — a v3 store carrying the key is a typed unknown-key reject again (old and new readers agree).
3. **The engine version stamps 4 when the config carries `resummon_cadence`; a cadence-less config may remain v3** — additive-forward, no forced bump; **confirm this against r13's transition rule with m-7 during their countersign** (their call on the exact stamping rule, not ours).
4. **A v3→v4 adjacent-forward transition entry** is registered and exercised by a fixture (the day-one migration-procedure promise — the path must exist even with zero live stores). Rollback/skip stay rejected.
5. **Behavior byte-preserved:** `validResummonCadenceShape`, the fail-closed skew directions (v2 cannot carry it), and the no-auto-approve rule (present zero = immediate resummon, never a verdict) — only the version home moves.

### RED-first (capture to `frank/.relays/s11/fold-v4-red-green.md`, sequence-honest)

RED before the production edit: (i) a v3 store carrying `resummon_cadence` must be REJECTED (fails today — v3 currently accepts it); (ii) a v4 store carrying it must be ACCEPTED (fails today — no v4 arm); (iii) the v3→v4 adjacent-forward transition must validate (fails today — no entry). Then GREEN all three + the existing cadence/skew/no-auto-approve fixtures re-pinned to v4, byte-preserved assertions intact.

### After your fold commit

1. Report under this DISPATCH_ID (`FROM: s11.implementer`, parented here): the FOLD_SCOPE-conformant diff, the RED/GREEN capture, the new head, targeted results (`./internal/config` + the cadence fixtures + the ten INV-CATALOG laws), pushed to the PR branch.
2. **m-7.planner then countersigns the realized v4 bytes** (owner-fidelity over the r13 surface — the ruling's step 2; I request it on your report).
3. On the countersign, I re-run the targeted check + the full battery at the new head and **REISSUE the merge decision TO the operator**. The m-6/m-3 confirms and m-7's four clean cells stand unchanged; only the config member moved.

**Not granted:** anything outside the FOLD_SCOPE; any change to the validator/skew/no-auto-approve behavior beyond the version re-home; merge (operator-only, held on the reissued decision).

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/s11-build-fold-v4/REVIEW-FOLD-planner-20260714-171200.md` — run before handoff (result in the inline pointer).
- The ruling consumed: `s11-build-escalate-config-lock/RECONCILE-orchestrator-planner-20260714-170510` (TO this seat + m-7.planner); the contradiction's byte-verification is of record in `…/SITREP-planner-20260714-170500`.

ACTIONS_GIT_REF: none — a fold dispatch; no `frank/` edit by this relay. Artifacts this turn: this relay + one `master/relays/INDEX.md` row timestamped 20260714-171200.
FINAL_GIT_STATUS_SHORT: unavailable — cwd (/Users/jack/Programming/harness) is not a git repo (docs workspace); `git -C frank status --short` clean at `d91fcfb`; the s11 worktree clean at `e86644d`.
Next requested action: operator carries this to s11.implementer; the fold lands RED-first under FOLD_SCOPE; m-7.planner countersigns the v4 bytes on the fold report; I re-verify and reissue the merge decision at the new head. The merge grant stays held until then.
