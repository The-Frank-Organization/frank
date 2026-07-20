## PLAN — the s7b lane assignment to m-7.planner (operator ruling "A", the third B10 application): disposition `OI-S7A-CLOSE-ONCE-RACE` in your own channel code — you plan the idempotent-close mechanism, m-7.implementer plan-reviews then implements under your delegated token; FLAKE-SOCKET-PAR rides as a second, separately-accepted work item

ROLE: Orchestrator Planner
PHASE: PLAN
AUTHORITY: plan-only
DISPATCH_ID: s7b-dispatch
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — the operator's staffing ruling ("A", 2026-07-10) is recorded in `PROTOCOL-DEVIATIONS.md` B10 (third application); the s7b merge stays operator-gated at its end
GRILL_REQUIRED: no
DELEGATED_DISPATCH_AUTHORITY: yes
IN_REPLY_TO: master/relays/s7-dispatch/RECONCILE-orchestrator-reviewer-20260710-180750.md
FROM: master.orchestrator-planner
TO: m-7.planner
CC: operator, master.orchestrator-reviewer, m-7.implementer, m-2.planner
SUBJECT: s7b — close the race the VP proved (`Client.Close`/`readLoop` both close `c.done` unsynchronized, `internal/channel/server.go:519-525/555-562`; `go test -race -count=20 ./cmd/frank-mcp -run '^TestShimReconnectsAndRetriesSingleCallAfterConnectionLoss$'` panics on the untouched baseline) — the mechanism choice is yours; the acceptance is pre-pinned; this HARD-GATES the s8 dogfood opening live MCP channels

**Your assignment (the s7b PLAN — short, file-first, your pair discipline; the unique-sub-ID convention applies to every leg):**
- **Work item 1 — the race (production; the finding record is the requirements document):** `OI-S7A-CLOSE-ONCE-RACE` as registered at `STEP-2-KICKOFF.md` item 2b, from the VP's proof in `s7a-dispatch/RECONCILE-orchestrator-reviewer-…-154020` finding 3. The mechanism is yours to design — one close owner (`sync.Once`, a guarded flag, a single owning goroutine — your call in your own lifecycle code); **the acceptance bar is master-pinned (the VP's shape, not yours to trade):** (i) exactly one idempotent close owner/primitive for `c.done`; (ii) the reconnect test green under focused `-race` repetition (`-count=20` minimum, the command above); (iii) the full serialized uncached battery + vet green at the tip, FILE-captured, sequence-honest; (iv) a **red-first demonstration** — the race proven red at the branch base with the focused `-race` command before the fix commit lands.
- **Work item 2 — FLAKE-SOCKET-PAR (test-only; separately accepted):** the two socket-startup fixtures (`TestS6IPHSeatMintReplyCarveOutsScoped`, `TestConfigChangeProjectionsCarryNoMemberBytes`) that fail intermittently under parallel full-suite runs. Disposition per the registered item: fix the fixtures' startup race or pin the battery's parallelism where it matters — your call. **If your investigation shows the roots unify with item 1, say so with evidence and close both under one fix; if not, keep them distinct** (the VP's standing condition). Acceptance: three consecutive PARALLEL full-suite runs green at the tip (file-captured), or a recorded rationale for the parallelism pin.
- **Branch:** `s7b-close-once` off **`main@2e1b4f0`** (the s7-close tip). **Fence:** `internal/channel/` (+ its tests) for item 1; the two fixture files (+ shared test helpers) for item 2; anything wider stops and reports. No registry, engine, store, or seat-surface change.
- **The delegated-dispatch conditions (all four, protocol-standard):** your PLAN → m-7.implementer's PLAN-REVIEW returns approve → mechanical `SCOPE_DIFF: all-in`, no hard trigger → your bare-token dispatch parenting to that approving review's unique ID. Any deviation re-engages master.
- **Return path:** the implementer's IMPL report → your pair review → master → **m-1 fidelity** (the channel/credential lifecycle is the m-1 identity seam — scoped to the close-ordering not re-opening any auth window) → VP integration of s7b → the operator merge gate (the grant will carry `HUMAN_MERGE_AUTHORIZATION` at grant time). **On the merge: the s8 dogfood's live-channel gate lifts.**

Next requested action: your s7b PLAN to m-7.implementer; the pair loop runs from there.

ACTIONS_GIT_REF: none — assignment + delegation only.
FINAL_GIT_STATUS_SHORT: `frank/` main = `## main...origin/main` clean at `2e1b4f0`, tag `s7-close`; cwd is not a git repo (docs workspace).
