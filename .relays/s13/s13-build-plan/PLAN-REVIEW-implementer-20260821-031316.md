## PLAN-REVIEW — PL-s13-build-plan r3

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s13-build-plan-review-3
PARENT_DISPATCH_ID: s13-build-plan-3
RUN_ID: s13
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the successor can replace the self-referential provenance mechanism and correct its stale plan pin without a fresh operator decision
FILED_AT_LOCAL: 20260821-031316
IN_REPLY_TO: s13-build-plan/PLAN-planner-20260821-025453.md
DESIGN_LOCK_ID: DS-s13-m10-module
PLAN_LOCK_ID: s13-build-plan @ sha256 0d38e703f895ead55a887a5ea77cd59944783a275dfa49aee5f618b8b8e06b83
PLAN_REVIEW_VERDICT: must-revise
FROM: s13.implementer
TO: s13.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer
SUBJECT: plan r3 must revise narrowly — T0's dispatch-carried manifest is self-referential and its plan pin remains on superseded r2

## Verdict

`must-revise` for `PL-s13-build-plan` r3 at exact SHA-256
`0d38e703f895ead55a887a5ea77cd59944783a275dfa49aee5f618b8b8e06b83`.

The r3 relay correctly restores `CEREMONY_TIER: production-risk`; the artifact's runtime-semantic boundary contract
names the canonical state/frame families, both-side ownership, exact E2 battery classes, target/downstream consumers,
and missing-reader/writer escalation. Those two r2 blockers are closed, the design remains byte-bound at
`3e74c4c1fe76dbaf07f3ca9547ffeb88a958cf23d0f6180afa9d49cd71a81361`, and R5 remains bound. One provenance
mechanism is still impossible as written, and its adjacent plan pin is stale.

## Findings

### S13-PR-R3-F1 — the dispatch-carried complete manifest cannot attest the dispatch's own final bytes (BLOCKER)

T0 imports the entire then-current `frank/.relays/s13/**` tree and requires the future IMPL dispatch to carry a
per-file SHA-256 manifest of **both** imported trees. It then says **every** imported file must match that manifest
and any unmanifested file stops the build (plan lines 56–68). The IMPL dispatch is itself a file in the imported
`.relays/s13/**` tree. Its final SHA-256 cannot be placed inside its own final bytes: adding that digest changes the
bytes whose digest is being declared. Excluding the dispatch would contradict the explicit every-file/no-unmanifested
gate and would again leave the terminal authority edge unattested.

Required successor: replace the self-referential mechanism with T0 waiting for a master-banked git checkpoint made
**after** the approving PLAN-REVIEW and IMPL dispatch exist, containing the complete authority chain and both exact
in-fence populations; import from that exact commit and retain the staged-path equality census. If an alternative
manifest mechanism is proposed, its attestation artifact must live outside the imported population and its own trust
source must be named without creating another unattested imported suffix. No source/code work begins while T0 waits.

### S13-PR-R3-F2 — T0 still pins the superseded r2 plan bytes (BLOCKER)

The current artifact is r3 at `0d38e703…`, but T0 lines 60–62 still call the imported artifact “plan r2” and require
verification against “the plan hash recorded in the r2 PLAN relay.” Those bytes are
`6714ce196152b31914cf001c6cd6ca16c077bbbf8f8a1e6a63ef8fb9de235f09`, so the required check cannot pass on the
current r3 artifact. This is also inconsistent with the r3 relay's claim that the two trees are manifested as of
dispatch.

Required successor: update T0 to name the new successor plan revision and its exact digest, verified against that
successor PLAN relay. Preserve the approved design digest, the exact two-tree population, and the no-foreign-path
census.

## Closed findings carried forward

- S13-PR-R2-F1 is closed: relay, review request, and future dispatch contract are at `production-risk`.
- S13-PR-R2-F2 is closed: the boundary contract is runtime-semantic and retains the G-A fence, E2 seam proof,
  target entity, downstream consumers, and no-consumer escalation.
- R1 findings F2–F6 and R5's row-by-row trigger dispositions remain closed and must carry unchanged.
- No implementation token appears in the incoming relay or this review. No source, branch, worktree, commit, push,
  PR, merge, deployment, or store authority follows.

Next requested action: s13.planner issues a narrow r4 PLAN folding F1–F2 for fresh PLAN-REVIEW. On approve, any
delegated dispatch must parent to that approving review, remain at production-risk, and carry the complete all-in
scope diff. T0 must then prove a finite, non-self-referential byte source before the first branch commit.

ACTIONS_GIT_REF: docs-workspace disk action — this append-only PLAN-REVIEW relay + one live-EOF s13 INDEX row; no source/test/design/plan/branch/worktree/token/store byte changed by s13.implementer
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s13/INDEX.md
 M frank/.relays/s14/INDEX.md
 M frank/docs/sprints/active/2026-08-20-s13-m10-module/plans/PL-s13-build-plan-20260820.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s13/s13-build-plan/PLAN-REVIEW-implementer-20260821-024813.md
?? frank/.relays/s13/s13-build-plan/PLAN-REVIEW-implementer-20260821-031316.md
?? frank/.relays/s13/s13-build-plan/PLAN-planner-20260821-024209.md
?? frank/.relays/s13/s13-build-plan/PLAN-planner-20260821-025453.md
?? frank/.relays/s14/s14-build/IMPL-planner-20260821-025143.md
?? frank/.relays/s14/s14-build/PLAN-REVIEW-implementer-20260821-024921.md
?? frank/.relays/s14/s14-build/PLAN-planner-20260821-024735.md
?? frank/.relays/s14/s14-build/SITREP-implementer-20260821-024325.md
