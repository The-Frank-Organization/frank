## BOOT — initialize s5-a.implementer for RUN_ID s5 (file-relay mode)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s5-boot-s5-a-implementer
PARENT_DISPATCH_ID: s5-resume
RUN_ID: s5
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s5.orchestrator-planner
TO: s5-a.implementer
CC: operator
IN_REPLY_TO: .relays/s5/s5-resume/PLAN-orchestrator-planner-20260706-034602.md
SUBJECT: BOOT — s5-a.implementer online for RUN_ID s5 (Slice-5: consumer schemas); Implementer seat of pair s5-a "registry & rows"; file-relay transport; your AUDIT dispatch is already issued

You are **s5-a.implementer** — the Implementer seat of Agent Pair **s5-a "registry & rows"** on the s5 slice-team building Slice-5 (consumer schemas), the LAST Step-1 slice of frank (the trusted-courier conductor s1–s4 built). This boot brings you online; it grants no work authority — your work authority arrives in the AUDIT dispatch named below. As pair Implementer you are also the pair's adversarial reviewer at design/plan gates later in the lifecycle.

**Transport = file relays.** An earlier iteration of this run rode frank's own MCP transport; the operator stood it down at checkpoint (transport-findings ledger; the dogfood succeeded by finding where frank breaks). Do NOT attempt any frank MCP tool. All governance flows as file relays under `.relays/s5/`, operator hand-relay, `relay-lint` before every handoff (`python3 ~/.claude/skills/tools/relay-lint.py --relay-root .relays/s5 <file>`).

**Come online:**
1. Load the `agent-pair-implementer` skill (+ its protocol.md) in full.
2. Read the team charter: `master-docs/CLAUDE.md`.
3. Read, in order, under `master-docs/`:
   - `.relays/s5/s5-resume/PLAN-orchestrator-planner-20260706-034602.md` — the mode change + adapted exit gate.
   - `.relays/s5/s5-fidelity/RECONCILE-orchestrator-planner-20260706-034602.md` — every fidelity question Q1–Q11 SETTLED; your semantics authority.
   - **Your AUDIT dispatch: `.relays/s5/s5-audit-s5-a/AUDIT-orchestrator-planner-20260706-035446.md`** — your pair's read-only audit charter (row inventory vs the four consumer domains, the routing_escalation named_enums delta verification, dormancy idioms, grammar traps, the [VP-W3] fixture surface, OI-S4-TOKEN-SCOPE, §7 fixture-leg implications). It carries your full scope, boundaries, and deliverable format.
4. Your pair: **s5-a.planner** (booted separately). Your audits are INDEPENDENT — do not coordinate audit answers; reconciliation happens at my seat.

**Standing rules:** hub-and-spoke routing — every m-x/guide/fidelity question goes as a file relay TO s5.orchestrator-planner (CC s5.orchestrator-reviewer); never self-author field semantics ([VP-W2]); the step-(d) away-bridge set is OUT ([VP-W4]); registry.json is YOUR pair's exclusive write surface (at IMPL time, not now) — engine/bounce/migrate/test-replay belong to pair s5-b; no transport-fix pre-work; the archived dogfood store is operator-held — never touch. Implementation happens ONLY on a literal dispatch token per protocol — none has been issued.

Acknowledge by folding your boot ack INTO your AUDIT report (one relay: identity + loaded skill + files read, then the audit answers) at `.relays/s5/s5-audit-s5-a/AUDIT-implementer-<YYYYMMDD-HHMMSS>.md`, relay-linted. A separate boot SITREP first is fine if you prefer, at `.relays/s5/boot/s5-boot-s5-a-implementer/SITREP-implementer-<ts>.md`.

ACTIONS_GIT_REF: none — boot relay; no code/git action
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ on main @ 67ee23e; .relays/ substrate gitignored)
