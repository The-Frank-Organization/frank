## DESIGN — H-16 BOUNDED FIX LANE to m-7 (operator-opened 2026-07-20; the external audit's one code-verified defect): the commit loop relabels a COMMITTED decision on post-commit failure — `internal/engine/loop.go:168-185` returns `Rejected` carrying the committed relay ID when `completeTurn`/`AfterGateResolution`/`AfterApprovalResolution`/`AfterAccepted` fail, while a retry returns the persisted (possibly Accepted) state via `existingOutcomeForCommand` (:136,:299) — the same intake presents rejected-then-accepted; fix = the monotonic split `decision_state × post_commit_state`, PRE-T4

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: h16-outcome-split
PARENT_DISPATCH_ID: master
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — the operator opened this lane (in-session, 2026-07-20, recorded here §8b-style) and retains the MERGE grant; the fix lands on `frank/` main only on operator merge ratification
GRILL_REQUIRED: no — a bounded defect fix inside the existing outcome contract; no product semantic choice
DESIGN_DOC_ID: h16-outcome-split-design
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260719-224500.md
FROM: master.orchestrator-planner
TO: m-7.planner
CC: m-7.implementer, master.orchestrator-reviewer, operator
SUBJECT: the defect class is OURS (confusion, not malice): a client that sees `rejected/obligation-error`, retries, and sees `accepted` cannot know whether authority was granted — and the T4 build is about to wire the m-9/m-10/m-8 callers to exactly this API, which is why this is a PRE-T4 micro-lane (the s7b shape) and not a T4 rider; the canonical STORE is correct throughout (the append-only record never relabels) — only the RETURNED outcome lies

m-7 — the fix lane, bounded exactly:

1. **The defect (verified at the bytes by master, 2026-07-20):** after `commit` succeeds and `relayID` exists, failures in the four post-commit hooks return `Outcome{State: record.Rejected, RelayID: relayID, Reason: obligation-error|derived-work-error}` — a rejected outcome naming a committed (possibly Accepted) record. `existingOutcomeForCommand` then serves the persisted truth on retry. Also present on the `callHandler` path (:279-280).
2. **The fix shape (the audit's, master-endorsed; your bytes decide the exact form):** split the outcome into two monotonic dimensions — `decision_state ∈ {accepted, rejected, held}` (fixed at commit, NEVER relabeled by anything after) × `post_commit_state ∈ {complete, pending, failed}` (the fate of derived work: turn completion, gate/approval resolution hooks, extras). A post-commit failure becomes a DURABLE derived-work fault (a typed record or obligation the loop retries idempotently or parks) — never a decision relabel. Callers that only understand the old shape must fail closed, not misread (your call whether that's a versioned Outcome or an additive field with the old `State` preserved as `decision_state`).
3. **Scope rails:** the Outcome shape + the four hook failure paths + `existingOutcomeForCommand` coherence + tests (incl. the crash/retry cut: post-commit failure → restart → retry serves the committed decision + the pending derived-work fault). The ten INV-CATALOG laws stay green untouched; no refactors beyond the seam; Step-2 remains CLOSED — this is a defect fix inside its contract, not a re-open. Branch work in `frank/`; pair review per discipline; the MERGE onto main = operator grant on your SITREP.
4. **Sequence:** your small design note (the exact shape + compat posture) → m-7.implementer concurrence → IMPL on a branch + tests → the uniquely-parented implementer review of the diff → SITREP with the branch/commit + test evidence → operator merge grant → master records H-16 CLOSED in the backlog.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row; no design doc, no `frank/` action yet (the lane authorizes branch work, not merge).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `6e4d657`.
Next requested action: m-7.planner returns the fix design note; the VP's review of this dispatch (routed in parallel) may narrow it before work begins — check the lane before starting IMPL.
