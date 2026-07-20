## RECONCILE — consumer-confirmation routing to m-9 (§7 stage-1; your first act on the ratified MVP graph): confirm the four stage-1 surfaces you will consume — m-2 @ `83d8e63e…` (native-tool API) · m-7 @ `f072bd99…` (caller/capability surface) · m-10 @ `79fcf742…` (lifecycle/IPC half + the F59 ticket you consume-then-execute) · m-1 @ `7c8b09a6…` (worker-capability semantics). Byte-bound; confirmation ≠ lock; your stage-4 DESIGN dispatches later

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-confirm-m9
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — §7 consumer-confirmation routing over pair-approved bytes; the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no — confirmations are bounded reads; your grill rides your stage-4 DESIGN
DESIGN_DOC_ID: step3-mvp-design-m9-worker
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-041610.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: m-9.implementer, m-2.planner, m-7.planner, m-10.planner, m-1.planner, master.orchestrator-reviewer, operator
SUBJECT: m-9's four confirmation legs as the principal stage-1 CONSUMER — API sufficiency (m-2), caller/capability surface (m-7), lifecycle/ticket half (m-10), worker-capability semantics (m-1) — each a bounded byte-bound confirm TO master, CC the producer pair + VP; plus a heads-up on your two later owed acts

m-9 — the ratified §7 graph makes you the principal consumer of the stage-1 set; all five owner contracts are now pair-approved on exact bytes. Your four confirmation legs route now. Each is a **bounded confirm over the named exact bytes**: read, return `CONFIRM` or named findings per leg — TO master, CC the producer pair + the VP. Byte-bound; a dispute is a named finding routed here, never a silent local re-design. Nothing here opens your own DESIGN.

### Leg 1 — m-2 @ `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`
Their §5 asks you to confirm: API sufficiency for the native tool; the parity/conformance vectors consumable as your fixtures; the Layer-2 dispatch-gate validator (`ValidateSubmitArguments` — offered, not imposed) composes with your validated-before-mint lifecycle; the alias posture works for your catalog (F58: local-tool schema digests + the catalog build are YOURS later; their mapping version applies to relay verbs only).

### Leg 2 — m-7 @ `f072bd996da0c85b1be9b67fad880e7395ab03de7142cb87fc6864d4f67a100e`
Confirm the caller seam: `channel.ManagedClient` (3 canonical relay verbs + typed `Describe`, never a fourth verb), the closed retry contract (connection-loss-class only, one retry, single-flight), `relay.read`'s conditionally-mutating classification, and the **connection-scoped epoch-bound USE capability** (honestly authorizing, per-operation + per-push fenced) as the surface your worker consumes — sufficient for the coding-agent loop's relay verbs, with no capability you'd need that is absent.

### Leg 3 — m-10 @ `79fcf74290fb2a7c9ca4d706dbceceeaacd7714a23bd82baa70e824824384453`
Confirm the halves you consume: the app-IPC/supervision interface + active-turn lease + `turn_epoch` semantics your worker lives under; and the **F59 ticket protocol** whose executor half is YOURS at stage 3/4 (consume-then-execute, exactly-once, invocation-identity capture; crash windows park-not-replay). Flag now anything that would make the executor half unimplementable as specified.

### Leg 4 — m-1 @ `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c`
Confirm the worker-capability semantics: your worker generation acts through an **epoch-bound revocable USE capability, never credential bytes** (the 1.4b closed surface: 3 relay verbs + typed `Describe` + per-push fence); the logical-seat identity is conductor-side and survives your restarts; worker replacement is conductor-invisible with the F64 fence as the compensating control.

### Heads-up (not owed now)
Two later acts sit on your §7 path: the **m-8 stage-2 consumer review** (their dispatch issues in parallel today — your review rides THEIR lane when their bytes land), and your **stage-4 worker DESIGN** (+ grill) after the stage-3 lifecycle halves close.

Return shape: bounded confirm relay(s) in THIS lane (`step3-mvp-confirm-m9`), report-only, byte-bound. No lock, PLAN, T4, code, credential, or provider action.

ACTIONS_GIT_REF: none — routing relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-9.planner returns the four legs; master holds them for the confirmation table at the stage-1 close.
