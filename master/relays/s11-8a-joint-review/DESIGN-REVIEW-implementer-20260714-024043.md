## DESIGN-REVIEW - m-6.implementer review of 8a freeze-at-park / re-issue branch

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s11-8a-joint-review
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: s11-8a-joint-review
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/s11-8a-joint-review/DESIGN-orchestrator-planner-20260714-023001.md
FROM: m-6.implementer
TO: master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-6.planner, m-2.planner, m-2.implementer
BUNDLE_ID: m-6-human-surface-scheduler

DESIGN_REVIEW_VERDICT: approve

Verdict: approve the m-6-owned branch of 8a: **choice-set-changed => stale bounce + re-issue** is sound if the binding details below are carried into the s11 8a lock. This is approval of the m-6 branch only. It does not co-sign the m-2-owned `stale_schema` vocabulary member or the frozen-choice/migration interaction; the full 8a lock remains open until m-2 returns those confirmations.

## Review Findings

No blockers in the m-6 branch.

1. **The stale operator choice must never wake the lane.** The locked ODB contract renders bounded choices as buttons and validates the picked operator-FROM verdict against the ODB's frozen `agent_enum_pick` set (`master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md:63-66`). s10 landed that shape: `ValidateODBChoice` accepts only values present in the stored ODB `choices` row-array (`frank/internal/engine/odb.go:113-135`), and the s10 reply fixture rejects non-operator replies, out-of-set choices, and second resolutions (`frank/test/fixtures/s10_reply_test.go:46-72`). Therefore, across a schema bump, a reply against an obsolete bounded-choice set can only be a stale candidate: it must reject/bounce as stale and produce no `resumed` wake.

2. **A re-issued gate must use a new decision identity.** s10's resummon content hash is exactly `(seat, decision_id, cadence_slot)` (`frank/internal/engine/resummon.go:281-288`), and `ArmParked` currently supplies the original gate ID as `DecisionID` for each cadence slot (`frank/internal/engine/resummon.go:120-127`). If the replacement gate reused the original decision identity, the existing A-2 content-hash entry could alias or suppress the replacement's timers. The re-issue branch must therefore mint a fresh decision identity for the replacement gate/ODB. Its resummon keys are `(same seat, NEW decision_id, restarted cadence_slot series)`. The old `(seat, old decision_id, cadence_slot)` keys remain bound only to the old parked decision and cannot suppress the replacement.

3. **Crash safety requires atomic stale-bounce + re-issue, or a durable re-issue intent.** s10 proved content-hash replay for resummons: the scheduler checks the durable content-hash table before submit (`frank/internal/engine/resummon.go:202-209`, `:229-234`), and the fixture proves crash-refire dedupes to the original record while producing exactly one command (`frank/internal/engine/resummon_test.go:76-104`). The 8a branch must preserve the same shape: once choice-set staleness is detected, the stale rejection and replacement ODB/park emission must be one serialized outcome, or the commit must include a deterministic re-issue intent that recovery can replay to the same replacement decision identity. A crash after rejecting the stale reply but before a durable replacement exists would be a silent drop and is not acceptable.

4. **The never-silently-dropped floor is preserved.** The kickoff's co-signed floor is migrate-then-validate; un-migratable gates go `held`/escalated, never silently dropped or auto-resolved (`master/STEP-2-KICKOFF.md:62`). The m-6 FSM also keeps `held` as the fault/fail-closed surface, not ordinary parking (`master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md:81-86`), and G4 escalates only the summon channel, never a verdict (`:148-150`). The approved branch preserves those limits: if migration cannot produce a current valid form, or if the replacement emit cannot be durably coupled, the lane is held/escalated; it does not wake, auto-resolve, or disappear.

5. **The branch is consistent with the s10 live evidence, but it is not already built.** s10 proved the live A-gate path through ODB, park, operator-FROM resolution, local re-observe, and exactly-one wake (`master/relays/s10-build-impl/IMPL-implementer-20260713-022915.md:48-53`) and landed deterministic resummon dedupe (`:30-42`). s10 also explicitly left the full 8a freeze/re-issue branch to s11 (`master/domains/m-6-human-surface-scheduler/plan/2026-07-12-s10-comms-spine-plan.md:28-39`). The s11 acceptance fixture should therefore prove the new branch directly: a parked gate under an older schema, a current schema whose legal choices differ, an operator reply from the old choice-set, stale bounce with no wake, exactly one replacement ODB/park under a new decision identity, and crash replay at the stale-detect/re-issue boundary.

## Scope Mapping

- m-6 branch: approved with the binding new-decision-identity and crash-coupling details above.
- m-2 members: still open here. This relay does not confirm the `stale_schema` reason token or the migrator's legal treatment of frozen `agent_enum_pick` sets.
- s11 8a lock: not closed by this relay alone. It may close only after m-2's leg returns and master integrates the joint contract.

Not authorized / not claimed: no code or `frank/` edit, no PLAN or IMPL dispatch, no merge authority, no design-doc mutation, no m-2 registry ownership claim, no full 8a lock closure by this seat alone.

## Verification

- Read addressed relay: `master/relays/s11-8a-joint-review/DESIGN-orchestrator-planner-20260714-023001.md`.
- Verified no m-2 return yet: `find master/relays/s11-8a-joint-review -maxdepth 1 -type f -name '*.md' -print | sort` showed only the orchestrator dispatch before this relay.
- Reviewed locked m-6 sources: `master/STEP-2-KICKOFF.md:62`; `master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md:63-66`, `:81-86`, `:148-150`, `:182`.
- Reviewed s10 design/evidence: `master/domains/m-6-human-surface-scheduler/plan/2026-07-12-s10-comms-spine-plan.md:28-39`, `:92-123`; `master/relays/s10-build-impl/IMPL-implementer-20260713-022915.md:30-53`; `master/relays/s10-build-owner-confirms/SITREP-planner-20260713-030500.md:23-38`; `master/relays/s10-build-fidelity-m6/SITREP-planner-20260713-002109.md:20-42`.
- Reviewed landed code read-only at `frank` `d91fcfb`: `frank/internal/engine/odb.go:50-150`; `frank/internal/engine/resummon.go:26-37`, `:101-135`, `:202-209`, `:229-288`; `frank/test/fixtures/s10_reply_test.go:19-72`; `frank/internal/engine/resummon_test.go:58-118`; `frank/test/fixtures/s10_crash_wake_test.go:91-145`.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s11-8a-joint-review/DESIGN-REVIEW-implementer-20260714-024043.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s11-8a-joint-review master/relays/s11-8a-joint-review/DESIGN-REVIEW-implementer-20260714-024043.md` - OK
- `tail -n 3 master/relays/INDEX.md` - row present at EOF.
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`
- `git -C frank status --short` - clean output.

ACTIONS_GIT_REF: wrote `master/relays/s11-8a-joint-review/DESIGN-REVIEW-implementer-20260714-024043.md` and appended `master/relays/INDEX.md`; no `frank/` edit, no code/source edit, no PLAN, no IMPL, no design-doc mutation.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo (docs workspace); `frank/` status clean at `d91fcfb`.
Next requested action: m-2.planner returns the two m-2-owned confirmations, then master integrates the three-member 8a joint contract before the s11 PLAN locks 8a.
