## RECONCILE — THE TOOLCHAIN PIN RULED DURABLY: the canonical F63 release-build toolchain is **`go1.26.4`** (exactly the proposed literal); the Q-F63 ratification is now COMPLETE (m-9's co-sign banked); plan-8 reissues with the locked literal + the R2-MR-2 canonical RELEASE-MANIFEST recipe

**The ruling (durable, not dispatch-time):** the pinned Go toolchain for the canonical F63 release build is the exact literal **`go1.26.4`**. Basis, verified at master's own seat this act: `go env GOVERSION` in the commissioned worktree = `go1.26.4` (darwin/arm64) — the toolchain under which the ENTIRE slice evidence trail exists (the WP1 composed proof and its three-seat corroborations including master's own battery/census runs, both WP2 capture runs, every gate run); `go.mod:3` carries only the minimum `go 1.25.0` with no `toolchain` directive, so the pin is the release contract's to carry, and the RELEASE-MANIFEST records it per the ratified posture. **Rails on the pin:** (i) it is a LOCKED plan input at the plan-8 reissue — the script carries the literal as the independent comparison source and FAIL-CLOSES against `go env GOVERSION` before any build byte; (ii) the pin MOVES only by a governed amendment through the same chain that ratified the posture (m-8/m-9 visible, master ruling) — never by drifting to whatever toolchain is installed; (iii) a GOVERSION mismatch at build time is a STOP + finding UP, never a local toolchain swap inside the dispatched work.

**The ratification record, updated:** m-9's co-sign (`master/relays2/s16-wp2-disp/SITREP-planner-20260828-024556.md` — §8.2 worker grain exact; the canonical-build posture accepted as strictly stronger; the four-member set joined) COMPLETES the Q-F63 ratification per my stated condition. The WP2b realization dispatch is therefore UNBLOCKED on the ratification side — it now waits only on the plan-8 → approving-review → all-in SCOPE_DIFF chain.

**Sequencing confirmed as the escalation laid it:** plan-8 reissues uniquely parented to `s16-wp2b-plan-review-7`, carrying (a) exactly `go1.26.4` + the comparison source, and (b) the R2-MR-2 canonical RELEASE-MANIFEST byte recipe as specified (canonical JSON · members sorted by name w/ lowercase-hex SHA-256 · the pinned literal + verified GOVERSION/GOOS/GOARCH/CGO_ENABLED · `GOFLAGS=""` · flags `["-trimpath"]` · a LITERAL output-recipe token never an absolute staging path · no timestamps/invocation-specific bytes · fixed field order, stated serialization, one trailing LF · idempotence = exact-byte comparison). The reviewer's law is endorsed: a pair review never approves a floating byte choice. Then the review, the SCOPE_DIFF, the token request; implementation only on master's fresh direct dispatch (V32). Every carried closure and hold stands; no code/script/build byte before the dispatch.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: s16-wp2b-pin
PARENT_DISPATCH_ID: s16-wp2b-pin-esc
RUN_ID: s16
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the routed pin ruling (master's, owners visible); the operator's gates remain the WP5 MERGE-GATE and the Step-3 ratification
IN_REPLY_TO: s16-wp2b/SITREP-planner-20260828-045950.md
FROM: master.orchestrator-planner
TO: s16.planner
CC: s16.implementer, master.orchestrator-reviewer, operator, m-8.planner, m-9.planner
SUBJECT: the toolchain pin RULED - go1.26.4 exactly (verified at master's own GOVERSION; the trail's toolchain), locked plan input with fail-closed GOVERSION comparison + governed-amendment-only movement + STOP-on-mismatch; Q-F63 ratification COMPLETE on m-9's co-sign; plan-8 reissues with the literal + the canonical RELEASE-MANIFEST recipe; dispatch on the V32 chain

ACTIONS_GIT_REF: engine-lane governance act — this ruling drafted at .engine/drafts/master.orchestrator-planner/ on the s16 root and submitted through the v2.9.1 client; verification this act was read-only (the escalation + review chain + co-sign read in full and exact-file-linted; GOVERSION/go.mod verified in the worktree); no source, test, script, branch, commit, push, PR, build, or merge byte.
FINAL_GIT_STATUS_SHORT: (literal, captured immediately pre-submit:)
 M frank/.relays/s16/INDEX.md
 M master/relays2/INDEX.md
?? frank/.relays/s16/s16-wp2b/PLAN-REVIEW-implementer-20260828-032120.md
?? frank/.relays/s16/s16-wp2b/PLAN-REVIEW-implementer-20260828-034510.md
?? frank/.relays/s16/s16-wp2b/PLAN-planner-20260828-024847.md
?? frank/.relays/s16/s16-wp2b/PLAN-planner-20260828-034042.md
?? frank/.relays/s16/s16-wp2b/SITREP-planner-20260828-045950.md
?? master/relays2/s16-wp2-disp/SITREP-planner-20260828-024556.md
