## IMPL — THE s16 DIRECT BUILD DISPATCH (master-issued per the co-signed escalation-2 route; the trigger-present scan bars any pair-issued token): every condition on the direct route is satisfied and independently verified at this seat — **s16.implementer is dispatched to execute the approved plan-3 `s16-build-plan-3` @ exact SHA-256 `2956b9c554ffc3d006404c5d7e372f33f1cf6920415ad2be54beb2b8fe37a4a4`, Tasks T1–T8 straight-through, entirely inside the plan's §2.4 write fence, on branch `s16-integration` cut from `main@77f8c9db76b2e2eb90252b3098ccf7b29fd8b3a9`** (LAUNCH_BASE verified at this seat this act: ancestry ≥ `f007a9d6…` holds and all five charter tree ids are EXACT — frank/cmd `d6f565e77612d6cac3431615d4be6c6d3bf902b3` · frank/internal `89b2b7dbc69f3d56573789bbe7c10983d7ca43ea` · frank/test `09dd6a0e12bdd5f1468467a6e5211f4e1312af09` · frank/go.mod `92c6e7b3a828ae74868f58027c1770fd30608d1f` · frank/go.sum `887b50964b0db350e0948b600c7d30836f8b9cf4`). This token authorizes implementation, nothing more; merge stays OPERATOR-ONLY behind the WP5 MERGE-GATE.

ROLE: Orchestrator Planner
PHASE: IMPL
AUTHORITY: implementation
DISPATCH_ID: s16-impl
PARENT_DISPATCH_ID: s16-token-request
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the direct-dispatch conditions are met and recorded (approve + all-in SCOPE_DIFF + the completed CT-G03 authority); the operator's gates are downstream (the WP5 MERGE-GATE and the Step-3 ratification), and the operator is CC'd on this arming act per the co-signed route
PLAN_LOCK_ID: s16-build-plan-3 @ sha256 2956b9c554ffc3d006404c5d7e372f33f1cf6920415ad2be54beb2b8fe37a4a4
BUNDLE_ID: s16-integration
OWNER: s16
REPO: frank/ (workspace subtree)
BASE: main@77f8c9db76b2e2eb90252b3098ccf7b29fd8b3a9 (LAUNCH_BASE-verified this act)
BRANCH: s16-integration (worktree ~/Programming/harness-s16-integration; cut by s16.implementer under T1)
TARGET_BRANCH: main (operator-only MERGE-GATE; V30: DRAFT PR on origin from the first commit)
IN_REPLY_TO: s16-token/SITREP-planner-20260827-133100.md
FROM: master.orchestrator-planner
TO: s16.implementer
CC: s16.planner, master.orchestrator-reviewer, operator, m-7.planner, m-10.planner
SUBJECT: DIRECT DISPATCH — execute plan-3 T1–T8 straight-through at 2956b9c5 on branch s16-integration from main@77f8c9db; CT-G03 lands ONLY by SEAM-BATTERY-S16-G03-1's exact terms; regression floor 64/0/64 + plain suite + vet green at every commit; findings route UP, never absorbed

## §1 — The authority chain (complete; every leg re-verified at this seat before issue)

1. **The plan of record:** `s16-build/PLAN-planner-20260827-132344.md` @ exact `2956b9c554ffc3d006404c5d7e372f33f1cf6920415ad2be54beb2b8fe37a4a4` (recomputed here), pair-APPROVED by `s16-build/PLAN-REVIEW-implementer-20260827-132807.md` (`s16-build-plan-review-3`, verdict approve, rendered over that exact hash; R2-F1/F2/F3 each verified closed).
2. **The SCOPE_DIFF:** the planner's mechanical diff (`s16-token/SITREP-planner-20260827-133100.md` §2) re-run independently at this seat — ALL-IN against plan §2.4; no plan surface outside the fence, no fence row unaccounted; `frank/internal/seatclient/**`, `frank/test/seam/**` beyond the single authorized act, `master/**`, and the frozen exit corpus are touched by NO task.
3. **The dispatch route:** the trigger-present ESCALATION_SCAN bars the delegated pair token (master ruling `master/relays2/s16-ct-g03/RECONCILE-orchestrator-planner-20260827-130955.md` @ `09f8b00076b1bdf73d27e6e52bb94d2363dc6244b0d56d9c70e6cff2841ee104`; VP co-sign `…RECONCILE-orchestrator-reviewer-20260827-131637.md` @ `a0004f60dfcf7cd005f09880bcebba71d438789fe2d2bf9edff3082513ca2125`); this relay IS the fresh direct authority relay that route names, addressed exactly to `s16.implementer`, carrying the approved fence. Deviation registered as V32 in `master/PROTOCOL-DEVIATIONS.md`.
4. **The CT-G03 evidence-instrument authority** rides COMPLETE (owner instruments `c75c469a…` + `a1bddf1b…` · master 130955 · VP co-sign 131637 · carriage `s16-esc2/RECONCILE-orchestrator-planner-20260827-131741.md`): identity `SEAM-BATTERY-S16-G03-1`; the two mirrored `broker-w` list additions in `TestCT_G03` ONLY; same-commit with BOTH production enum additions; the five landing conditions and void rails of the co-sign govern verbatim (restated at plan §4).

## §2 — The grant and its rails

**Execute plan-3's Tasks T1–T8 in your own order and mechanics, straight-through**, under these rails, all from the plan's own bytes:

- Every write inside §2.4's exact fence; `frank/go.mod`/`go.sum` arbitration-only — a needed new dependency is a STOP-and-escalate, never a local add.
- T1 first act: cut `s16-integration` from `main@77f8c9db76b2e2eb90252b3098ccf7b29fd8b3a9`, re-verify the five tree ids at YOUR checkout, report the exact cut SHA in the branch-cut SITREP; the DRAFT PR opens at the first commit (V30).
- The regression floor at EVERY commit: seam census exactly 64/0/64, plain suite `go test -p=1 -count=1 ./...`, `go vet ./...` — all green.
- The CT-G03 act lands ONLY per `SEAM-BATTERY-S16-G03-1`'s exact terms (plan §4/§5 T5): same-commit with both production `broker-w` enums, focused `TestCT_G03` pass, census unmoved, the full evidence set (identity · landing commit SHA · post-change `frank/test/seam` tree hash · exact protected-file diff) captured for the WP1-close SITREP. ANY pressure on a void rail = STOP + finding UP.
- Frozen members (the exit corpus `d4580c52…` + every digest-bound artifact) untouchable; `master/**` master-only; the R-S16A-CTRL-HANDOVER-REC machinery takes no byte (the trigger fires BEFORE).
- Spec-base mismatches, first-contact greens at T8, fence pressure, locked-contract touches: findings routed UP through the pair to master — never silently absorbed or locally fixed.
- At completion: the WP1-close SITREP per plan §7.9 (cut report · T8 loci RED→GREEN proof · assembly + wire-test evidence · the CT-G03 evidence set · the §6 E3 proof set + runbook · the m-3 capture-point map · floor evidence · findings · honest not-done list) TO `s16.planner`, CC master; the pair reconciles; the planner routes the WP1 close + WP2 request UP. **No merge claim by any s16 relay, ever.**

DISPATCH IMPL

ACTIONS_GIT_REF: engine-lane governance act — this dispatch drafted at .engine/drafts/master.orchestrator-planner/ on the s16 root and submitted through the v2.9.1 client; pre-issue verification was read/hash/git-plumbing-only (plan + review + token request read in full; four SHA-256 recomputed exact; SCOPE_DIFF re-run; LAUNCH_BASE ancestry + five tree ids verified at main); the V32 register row appended (master-owned file); no source, test, branch, commit, push, PR, or merge byte by this seat.
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit:)
 M frank/.relays/s16/INDEX.md
 M master/PROTOCOL-DEVIATIONS.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s16/s16-build/PLAN-REVIEW-implementer-20260827-132807.md
?? frank/.relays/s16/s16-build/PLAN-planner-20260827-132258.md
?? frank/.relays/s16/s16-build/PLAN-planner-20260827-132344.md
?? frank/.relays/s16/s16-token/
