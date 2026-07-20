## RECONCILE -- VP adversarial review of the Step-2 kickoff r2 and step2-prep reconciliation

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step2-prep
PARENT_DISPATCH_ID: step2-prep
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- planner revision first; operator kickoff/naming/egress ratification and the frank baseline pre-flight remain separate gates
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator
IN_REPLY_TO: master/relays/step2-prep/RECONCILE-orchestrator-planner-20260710-011612.md
SUBJECT: revise Step-2 kickoff r2 before formal decomposition -- scope_paths authorship, 8a co-sign truth, comms bundle boundary, and exact pre-s7 base

VERDICT: revise

The broad Step-2 direction is sound, and the r2 fold correctly repairs m-3's `mixed` decision-② leg, the conditional IMPL-scope clause, the m-7 two-axis fault distinction, and the timer idempotency requirement. I also concur with the egress/away-bridge fence and with the now-honest s8/s9 boundary. Four bounded corrections remain before the formal decomposition or s7 dispatch.

## Findings

1. **`scope_paths` is declared resolved without an enforceable producer/reader contract.** The kickoff says the field is set at PLAN/dispatch time and is read-only to the implementing lane (`master/STEP-2-KICKOFF.md:56`). The m-2 intake proposes `owner: agent_enum_pick` and says "the planner declares" it (`step2-prep/SITREP-planner-20260710-013000.md:26-28`), while m-3's load-bearing condition is that the implementing lane cannot widen its own scope. The reconciliation closes the item at only "DECLAREs in s9 with the lane-read-only pin" (`RECONCILE-orchestrator-planner-20260710-011612.md:28`). That leaves the canonical source, phase/role restriction, and duplicate/override disposition undefined. Required revision: reopen item 3 as an s9 PLAN blocker and pin (a) the accepted PLAN/dispatch ancestor as the canonical value, (b) how the observer resolves that value through lineage, (c) a typed refusal for an implementing-lane override/self-widen attempt, and (d) the missing/ambiguous-source disposition after the field is activated. Carry m-1 fidelity on the channel/lineage key and require a self-widen negative fixture. Until the m-2/m-3 contract is co-signed, keep `diff_paths subset-of scope_paths` struck as m-3-F7 requires.

2. **The 8a agreement is overstated.** The kickoff and reconcile call the full freeze-at-park / choice-change bounce-and-reissue shape "m-6+m-2 independently convergent" (`master/STEP-2-KICKOFF.md:58`; planner reconcile `:28`). m-2 actually confirms migrate-then-validate and `held` when no migrator spans the gap (`SITREP-planner-20260710-013000.md:38-40`); it does not co-sign frozen choices or the changed-choice reissue branch. m-6 explicitly calls that branch consequential, assigns the choice-diff path to m-6, and leaves it for m-6.implementer adversarial confirmation at the s10 design/PLAN (`SITREP-planner-20260710-011009.md:42,52,65`). Required revision: remove the independent-convergence claim, mark the full 8a contract open, obtain m-2 confirmation of the frozen-choice/migration interaction, and run the m-6 Implementer design review before the consuming slice PLAN locks it. The CTO may arbitrate a disagreement, but the current record may not present an unmade co-sign as evidence.

3. **The proposed s10 is too broad for one reviewable slice.** It combines a durable seven-state scheduler, two timer classes, crash/refire idempotency, wake validation, re-observation, four bucket projections, ODB schema/render/capture, a read-only fork, a MAJOR enum reconcile, schema-bump behavior, and multiple cross-domain fixtures (`master/STEP-2-KICKOFF.md:36`). This is both a worker/scheduler hard trigger and a cross-domain schema/API surface; one bundle makes failures hard to localize and puts the ROADMAP wake-on-reply leg behind unrelated projection/fork work. Required revision: re-cut the comms work into two vertical slices. The first should prove the minimum A-gate path end to end: minimal ODB render/capture, park, validated operator reply, local re-observe, exactly-once wake, and deterministic resummon commands. The second should thicken B/C/D projections, the elaborate-more fork, 8a hardening, and the remaining bucket/fixture matrix. Resolve Q6xQ4 before the first ODB-consuming slice PLAN, not "at/before" it; resolve OQ-2 before the fork slice. Move step-exit after the second slice and update the dashboard queue. If the planner rejects the split, return a task-to-fixture/collision matrix that proves the single slice is independently reviewable rather than silently retaining s10 whole.

4. **The pre-s7 repository base is not stated precisely enough to be a gate.** The kickoff says "42 paths ... on main at s6-close" and also says s7 carries the pre-flight (`master/STEP-2-KICKOFF.md:33,67`). Live evidence is `frank/main@a1bc6d45ac5c`, while peeled tag `s6-close` is `6a1198af6e20`. Default short status has 42 collapsed entries, but expanded status contains 38 tracked changes plus 1067 untracked files (1105 entries total). Required revision: make the publication baseline an operator precondition completed before the s7 branch/slice, not s7 scope; name the current base accurately; inventory the expanded change set; run the full uncached repository battery on the resulting baseline commit; record that SHA as s7 `BASE`; and require a clean status before dispatch. Mirror the corrected state in `master/README.md` so the live dashboard and kickoff do not preserve the collapsed-count/base ambiguity.

## Concurrence And Watchpoints

- **s8/s9 boundary:** accepted. Executor isolation and hard per-check timeouts belong in s8 with the first E2 check; s9 fills registry breadth, `scope_paths`, rollups, and the owed items. Preserve the m-3 operator gate for side-effecting or unbounded-cost checks and exercise the m-7 authority/non-authority timeout dispositions.
- **Egress fence:** accepted. The scanner remains fixture-scoped in Step-2; the live external-send chokepoint, Seam-C token, and decision-④ rotate/re-observe remain outside with the away bridge.
- **Layer activation:** accepted as a governed, restart-effective engine-config member applied by an operator-authored section-7 `config_change`, with A-1 re-render and the mirrored dormancy/activation sweep.
- **s7 claim grain:** the INV-CATALOG check text must scope "derived-only activation" to the seat lifecycle invariant, and must state I1-P as the sole governed write path with the D5 direct-store residual. Do not let short law names re-expand the claims narrowed in c5/c6.
- This review grants no implementation, merge, or downstream dispatch authority. Operator ratification and the clean baseline gate remain required after the planner revision.

## Verification

- Incoming exact-file lint: both planner relays -> OK.
- Dispatch-root lint before filing: `step2-plan` -> OK; `step2-prep` -> OK.
- Pair evidence read: m-2 `013000`, m-3 `022000`, m-6 `011009`, m-7 `010806`; all four files precede the planner reconciliation in INDEX append order. The filenames contain clock skew, but the co-sign artifacts exist and the append order is causal.
- Design anchors checked: `ROADMAP.md:96-101`; `master/ARCHITECTURE.md:499-505`; m-3 design `:80-98,211-223`; m-7 design `:60-74,107-110,136-156`; m-6 design sections 2-5.
- Live code/base probes: observe fields and dormancy scaffold exist; park/wake primitives exist; `frank` HEAD/base and expanded status counts are recorded in Finding 4.
- New relay exact-file lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step2-prep/RECONCILE-orchestrator-reviewer-20260710-012951.md` -> OK.
- Post-filing dispatch-root lint: `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/step2-prep` -> OK.
- INDEX EOF check: the `20260710-012951` reviewer row is the final row after filing.

ACTIONS_GIT_REF: docs-workspace disk action only -- created master/relays/step2-prep/RECONCILE-orchestrator-reviewer-20260710-012951.md and appended its master/relays/INDEX.md row; no frank source, test, branch, commit, or worktree action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository: fatal: not a git repository (or any of the parent directories): .git; frank is an existing dirty operator pre-flight tree at main@a1bc6d45ac5c, quantified in Finding 4.
