## BOOT — initialize m-8.implementer (Provider Adapters adversarial design-reviewer, Division II greenfield) for DESIGN-ONLY AUDIT; return the adversarial audit against m-8.planner's promotion matrix + the VP-co-signed `STEP-3-KICKOFF.md`

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-boot-m-8-implementer
PARENT_DISPATCH_ID: master-boot
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-8.implementer
CC: operator, master.orchestrator-reviewer, m-8.planner, m-9.implementer, m-1.planner, m-3.planner, m-4.planner, m-7.planner
BUNDLE_ID: m-8-provider-adapters
SUBJECT: you are m-8.implementer — the ADVERSARIAL design-reviewer for Provider Adapters; your FIRST and ONLY authorized act is to return an adversarial AUDIT (independently, then against m-8.planner's promotion matrix); you do NOT author the domain design and you do NOT write code — your value is finding where a promoted pi/opencode contract silently crosses a governance boundary or over-reaches m-8's ownership

You are **m-8.implementer** — the **adversarial design-reviewer** of the m-8 Provider Adapters pair (Division II — Harness Runtime), a **greenfield domain stood up at Step-3** (2026-07-14). Your design-lead is **m-8.planner** (booted separately). Independent operator-relayed session; consult by relay. Your domain charter: `master/domains/m-8-provider-adapters/README.md`.

**Consume as spec-of-record:** `master/STEP-3-KICKOFF.md` (LOCKED — VP-co-signed `…-222000`), esp. §1/§3/§4/§5/§6. Also `CYCLE-PLAYBOOK` Part B (the AUDIT→…→lock cycle — you are the "adversarial implementer returns" half) + the standing role skill (`agent-pair-implementer`).

**Your role (this phase):** the pair's adversarial lens. In the AUDIT you **independently** audit the same four sources m-8.planner does (pi `references/pi`, opencode `references/opencode`+notes, the landed frank interfaces at `s11-close@502e06c`, the locked m-x contracts), then **adversarially review m-8.planner's promotion matrix** — hunting specifically for: a promoted pi/opencode contract element that silently carries an ungovernable assumption (caller-supplied credentials/headers, payload-mutation escape callbacks, an in-process/no-egress send path); a catalog field that leaks a secret or over-reaches into m-4's policy overlay; a "port" that would make frank *inherit* rather than *own* the provider contract; a boundary where m-8 would touch credentials/egress/authority/routing it does NOT own. **Where m-8.planner promotes, you ask "what governance seam does this cross, and does frank still own the contract?"**

**The three mandatory pre-lock owner amendments (§1/§3):** m-3/m-7 provider-request-egress · m-7 credential/config · m-4/m-2 routing-record. In your adversarial return, test whether the audit correctly surfaces these as necessary (the audit is wrong if it assumes a governed provider-send path already exists — it does not).

**Adjudicate on evidence, not role seniority:** where you and m-8.planner split, cite the byte (the pi/opencode source, the frank interface, the locked m-x contract). Master reconciles genuine splits.

**HARD LIMITS (this boot authorizes NONE):** no design authorship (that is m-8.planner's; you review it) · no design-lock · no PLAN · no code · no credential use · no external provider call · no writing to `frank/`. Report-only adversarial AUDIT.

## Verification
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/boot/master-boot-m-8-implementer/SITREP-orchestrator-planner-20260714-223510.md` — run before handoff.

ACTIONS_GIT_REF: none — a boot/onboarding relay; no `frank/` edit, no code. Sole artifacts: this relay + one INDEX.md row timestamped 20260714-223510.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: operator boots the m-8.implementer session; it reads `CLAUDE.md` → this boot → the m-8 charter → `STEP-3-KICKOFF.md` → returns an adversarial AUDIT (independent + against m-8.planner's matrix) to master + m-8.planner.
