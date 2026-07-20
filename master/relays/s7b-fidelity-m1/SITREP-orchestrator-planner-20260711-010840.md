## SITREP — s7b fidelity request to m-1: the close-ordering seam at `s7b-close-once@e155aa6` — does the `sync.Once` close owner preserve your channel/credential lifecycle semantics (no auth window re-opened), checked against your B-3 / §8.5 re-attach record?

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s7b-fidelity-m1
PARENT_DISPATCH_ID: s7b-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — a scoped fidelity review; VP integration + the operator merge gate remain downstream
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7b-dispatch/PLAN-orchestrator-planner-20260710-234637.md
FROM: master.orchestrator-planner
TO: m-1.implementer
CC: operator, master.orchestrator-reviewer, m-1.planner, m-7.planner, m-7.implementer
SUBJECT: review the s7b close-once fix (worktree `/Users/jack/.config/superpowers/worktrees/frank/s7b-close-once`, tip `e155aa6`; the fix commit is `a2a6966` — `close(c.done)`/`close(s.done)` now owned solely by their `sync.Once`, `internal/channel/server.go:230/:527`) for the m-1-owned lifecycle seam; verdict TO master, CC the VP

**The lane (pair-complete; master-verified):** three commits — `a2a6966` (the idempotent close: the unsynchronized select/default close idiom replaced by `sync.Once` owners, fixing the double-close panic the VP proved on the baseline) · `5c678b4` (fixture startup hardening) · `e155aa6` (the granted delta: the one-line crashpoint block-after-kill + the invariants hardening + the child-mode short-circuits). Pair review APPROVE (`<worktree>/.relays/s7b/s7b-close-once/RECONCILE-planner-20260711-004630`); both registered defect classes closed by mechanism; master's runs at the tip: the reconnect station green `-race -count=20` · a full PARALLEL suite 25 ok / 0 FAIL (the previously flaky mode) · vet clean · the close-owner grep shows exactly the two `sync.Once` sites and no select/default idiom.

**Your scoped question (the pair framed it precisely — verify their claim, don't inherit it):** the `sync.Once` close owner changes *when* `c.done` closes relative to `conn.Close()` **in no observable way** — the owner fires first inside `Close()` exactly as the old default-arm did, and `readLoop`'s deferred close-path ordering is unchanged — so **no authentication/re-attach window re-opens**. Check that against your locked record for the channel/credential lifecycle: the B-3 boot/bind semantics and the §8.5 re-attach behavior (one active channel per credential; reconnect reoccupies the same seat; a superseded credential's channel force-close). Confirm, or must-revise citing the exact locked line.

**Return:** your verdict TO master, CC the VP. On your confirm: VP integration of s7b → the operator merge gate (`HUMAN_MERGE_AUTHORIZATION` at grant time) — and the merge lifts **only** `OI-S7A-CLOSE-ONCE-RACE` as the live-channel blocker, per the pinned boundary.

ACTIONS_GIT_REF: none — review request only; my verification runs were read-only in the worktree (evidence at `~/.claude/jobs/0908f73b/tmp/s7b-race-station.txt` + `s7b-parallel-suite.txt`).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `2e1b4f0` (tag `s7-close`); the s7b worktree at `e155aa6` — untracked `.relays/s7b/` only, implementation paths clean; cwd is not a git repo (docs workspace).
