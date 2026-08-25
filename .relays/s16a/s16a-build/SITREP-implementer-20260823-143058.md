## SITREP — s16a.implementer / s16a-build PLAN-REVIEW waiting on a filed parent

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s16a-build
PARENT_DISPATCH_ID: t4-s16a-commission
RUN_ID: s16a
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — s16a.planner can bank or reissue its PLAN and route the locked-ledger coverage question to master; an operator decision is needed only if master proposes a scope amendment
IN_REPLY_TO: s16a-build/PLAN-planner-20260823-141142.md
FROM: s16a.implementer
TO: s16a.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
SUBJECT: bounded wait — candidate WP1 PLAN is untracked foreign bytes and cannot parent a PLAN-REVIEW; preliminary locked-ledger and executable-test findings preserved

Phase: report-only bounded wait. No PLAN-REVIEW verdict is issued, and no implementation authority is present.

Current artifact: candidate bytes at `frank/.relays/s16a/s16a-build/PLAN-planner-20260823-141142.md` @ SHA-256 `aa9e5b01d0dd5a97d3f0bc46dd93da963e9f7049ff1c955c181d1a9c8bae48e6`.

Claims:
- The candidate is addressed `TO: s16a.implementer`, is structurally lint-clean with freshness disabled for historical verification, and names the completed four-PM round-trip. Evidence E1: exact candidate bytes plus exact-file lint.
- The candidate is not tracked in the current git object and has no git history; its INDEX row exists only in the modified worktree, while `HEAD:frank/.relays/s16a/INDEX.md` ends at the `134844` readiness row. Evidence E1: `git ls-files --error-unmatch` exit 1, empty path history, and HEAD INDEX readback.
- The s16a charter classifies a `??` path as uncommitted foreign bytes: read for awareness, never cite as a filed relay. Therefore a PLAN-REVIEW parented to this candidate would manufacture authority lineage over an unfiled parent.

Preliminary semantic findings to preserve for the filed successor:
1. **Locked-ledger coverage must resolve before onward routing.** The candidate correctly finds no r9 row whose acceptance cell binds the m-8-carried B/E producer-first digest carriage into `provider_attempts`: C08 binds the three assign digests; C04 binds `refusal_stage`. Deferring this question to the WP1-close SITREP would issue the token and build the 63-row battery before resolving a locked-surface coverage hit. Route it to master now; the successor PLAN must bind master's ruling without inventing a 64th row or silently dropping the coverage note.
2. **The pre-existing-battery command must be executable.** Acceptance criterion 2 writes `go test -p=1 -count=1 ./...` while also saying the expected-RED seam package is excluded; literal `./...` includes `./test/seam`. The successor must name the exact package-selection command that excludes only `github.com/jackli/frank/test/seam` and preserves all other pre-existing packages.

Done: read the candidate, charter, r5 plan, r9 ledger, completed round-trip, and all four PM returns; verified cited hashes, launch-base existence/ancestry, candidate lint shape, and the authority-parent defect.

Not done: no formal PLAN-REVIEW verdict; no branch, worktree, source/test byte, dependency, commit, merge, store, or credential action.

Blocked: a tracked/filed planner PLAN and a pre-token master disposition of the B/E coverage note.

Scope drift risk: approving uncommitted foreign bytes or inventing a non-ledger test would break the charter's authority and exact-close laws.

Tests / verification: E1 static/hash/lineage review only; no source tests are authorized in SITREP.

Next requested action: `s16a.planner` banks or reissues the PLAN, routes the B/E ledger question to master before token issuance, makes the pre-existing-package command exact, and sends the filed successor for PLAN-REVIEW.

ACTIONS_GIT_REF: disk action — this SITREP plus one append-only `frank/.relays/s16a/INDEX.md` row; no source/test or git-state action claimed
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s16a/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s16a/s16a-build/
