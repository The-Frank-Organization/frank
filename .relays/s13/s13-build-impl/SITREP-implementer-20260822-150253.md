## RESTACK HOLD — the first replayed commit exposed an overlap outside the ruled class; rebase is paused byte-untouched for master arbitration

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s13-build-impl-restack-1
PARENT_DISPATCH_ID: step3-t4-preflight
RUN_ID: s13
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — this is a bounded master-arbitration request under the predeclared restack stop
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/step3-t4-preflight/RECONCILE-orchestrator-planner-20260822-145641.md
FROM: s13.implementer
TO: master.orchestrator-planner
CC: s13.planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer, m-3.planner, m-7.planner, s14.planner, s15.planner
SUBJECT: s13 restack paused at unexpected design-doc add/add conflict outside frank/.relays/s13/**; exact-file successor ruling required before resolution

The restack word was verified at SHA-256
`1ca258281ed623493d7182f7f36d0fb383d9c02bdd9adc9f0f256aaaa555d4d0`. The source worktree was clean on
`s13-m10-module` at `7bbdb31181038dc36bea0425605ca91a8a18d115`; execution-time local `main` was pinned at
`571da52628af876e9320548b2a03441248c794c4`. The launch-base-to-target-main diff over
`frank/cmd/frank-app`, `frank/internal/appctl`, `frank/internal/appipc`, `frank/internal/executor`,
`frank/internal/ticket`, and `frank/internal/runmanifest` was empty.

`git rebase 571da52628af876e9320548b2a03441248c794c4` stopped at replay 1/19, `REBASE_HEAD`
`6a5930daba5df85fa9792493b648cf56ce46f430` (`s13 IMPL: bank governed build inputs`). No conflict was
resolved, no replay was skipped, and the rebase was neither continued nor aborted.

CONFLICT_CENSUS:
- EXPECTED class: `frank/.relays/s13/INDEX.md`, add/add. Stage 2 (target-main) blob
  `2e75bda0b5e2da01276838a9e2625ede18b4c5f3`; stage 3 (replayed branch commit) blob
  `91556ad1957f3746a98302159c715f903c204694`. The authorized append-only-union resolution has NOT yet
  been applied because the same replay also fired the arbitration stop below.
- UNEXPECTED / OUTSIDE-RULED-OVERLAP class:
  `frank/docs/sprints/active/2026-08-20-s13-m10-module/designs/DS-s13-m10-module-20260820.md`, add/add.
  Stage 2 (target-main) blob `e19b1eccee7401f65ba6907f37e64d92b0b12823`; stage 3 (replayed branch
  commit) blob `90d2295d80f4079b275fc8c6c8806ef171c51f86`. The target-main side is the r5 design and differs from
  the branch side by 76 changed lines (75 additions, 1 deletion when viewed mainward); the pre-restack branch
  tip itself also names blob `90d2295d80f4079b275fc8c6c8806ef171c51f86`, so taking target-main would materially change a reviewed
  design byte outside the word's sole expected overlap class.
- SOURCE conflicts at this stop: zero. The only two unmerged paths are the INDEX and exact design doc above.

ARBITRATION_REQUEST: issue an exact successor choosing either (a) extend the main-canonical restack rule to the
single design-doc path above and authorize preserving target-main's r5 bytes while the historical branch add is
replayed, with planner re-review as required, or (b) direct abort/re-cut. This seat will not infer that authority
from the `.relays/s13/**` rule. The locked rerun gates and carriage-v4 checks have not started because no
restacked tip exists.

ACTIONS_GIT_REF: attempted only `git rebase 571da52628af876e9320548b2a03441248c794c4`; rebase paused at
`6a5930daba5df85fa9792493b648cf56ce46f430`; branch ref `s13-m10-module` remains
`7bbdb31181038dc36bea0425605ca91a8a18d115`; no resolution, commit, merge, push, or gate run.
FINAL_GIT_STATUS_SHORT:
 AA frank/.relays/s13/INDEX.md
 AA frank/docs/sprints/active/2026-08-20-s13-m10-module/designs/DS-s13-m10-module-20260820.md

Next requested action: `master.orchestrator-planner` rules the exact design-doc conflict. The rebase remains
paused at the first conflicting replay so all three staged objects remain inspectable.
