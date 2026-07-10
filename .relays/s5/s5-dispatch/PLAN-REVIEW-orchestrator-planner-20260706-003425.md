## PLAN-REVIEW request — VP pre-handoff gate on the s5 package (dispatch + boot held; the LAST Step-1 slice, first team on frank)

ROLE: Orchestrator Planner
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s5-dispatch
PARENT_DISPATCH_ID: s5-dispatch
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — pre-handoff VP review (the standing cadence); the operator holds the handoff until your verdict, and I do not `submit` the dispatch into frank until you clear it
IN_REPLY_TO: .relays/s5/s5-dispatch/PLAN-orchestrator-planner-20260706-003425.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator
SUBJECT: review the s5 package before handoff — the last Step-1 slice (consumer schemas + the five owed §C4 fixtures + schema_version/migrator), the first team to run its governance ON frank, the m-x-via-master routing constraint, the STEP-GATED-OFF dormancy claim, the OUT fence (step-(d) away-bridge gates excluded)

**The ask.** Pre-handoff review of the Slice-5 package — both artifacts cut, lint-clean, **held** (nothing `submit`ted into frank yet):
- Dispatch: `.relays/s5/s5-dispatch/PLAN-orchestrator-planner-20260706-003425.md`
- Boot: `.relays/s5/boot/s5-boot-orchestrator-planner/SITREP-orchestrator-planner-20260706-001736.md` (already cut, lint-clean, INDEX'd; carries the load-bearing routing directive)

**What rides this review (beyond your standing checks):**
- **The dogfood framing is honest.** s5 is the first slice-team whose governance flows through frank. Does the dispatch over-claim what that proves? The ceiling is still **transport/provenance only** — consumer fields are *declared, not observed*; done-state + `record_integrity` stay `self_reported` until Step-2. Flag any surface where "ran on frank" drifts toward "observed/verified."
- **The routing constraint is sound + enforced.** m-1…m-7 are not minted on frank; s5 must route every guide/fidelity/m-x question through master. Is the hub-and-spoke stated tightly enough that s5 cannot mistake silence for permission to guess consumer *semantics*? (I bound this in the F2 plan-gate: unsettled semantics = a fidelity question = escalate.)
- **The STEP-GATED-OFF dormancy claim.** The consumer schemas are declared but must never be demanded by a Step-1 gate (CQ-1(a) phase-split required-set). The exit gate requires a **negative fixture** proving a Step-1 phase form neither requires nor renders a consumer field. Is that the right proof obligation, and is it stated as a hard acceptance rather than a hope?
- **The OUT fence — the step-(d) exclusion.** The dispatch explicitly excludes the away-bridge gates (Decision ④ away-token fixtures, R2 per-column *negative* fixtures, altitude-B per-row deviation-grain fixtures, away-mode trigger). These are §C4 "step-(d)" carries, NOT S5. Confirm I have the S5/step-(d) line in the right place — the five owed *Step-1-build* fixtures are IN; the step-(d) *away-bridge* fixtures are OUT.
- **The §J2 `routing_escalation` member ownership.** I scoped the member add as a **CTO-owned cross-domain action** (I author §J2 + m-2 mirror + m-6 confirm, via routing), with s5 only registering the resulting enum value + a fixture, and flagged it **non-blocking** (correctness holds via `other`→A). Is that the right split — should any of it gate S5, or is clarity/telemetry-only correct?
- **Addressing adaptation.** The frank `submit` of the dispatch can address only `s5.orchestrator-planner` (the sole non-master minted seat); the rich CC (you, operator, m-x) is served by the file + master routing. Confirm that split does not weaken the "authority-bearing orchestrator relays CC the reviewer" discipline — you review the file; the frank relay is the s5-facing transport only.

**Not asked of you:** the consumer *field semantics* themselves (that is m-3/m-4/m-5/m-6 fidelity, engaged via master during the build) — only the decomposition, scope fence, gate, honesty ceiling, and routing model.

**On approve:** I fold your watchpoints inline as **[VP-Wn]**, then `submit` the dispatch into frank TO `s5.orchestrator-planner`; the operator wires + boots the s5 session; s5 reads its dispatch via `project`/`read` and acks via `submit`.

## Verification
- `python3 ~/.claude/skills/tools/relay-lint.py <this file>` + `--relay-root .relays/s5/s5-dispatch` — run below.
- The package is held: no frank `submit`, no operator handoff, until your verdict. Baseline `frank/` `main` (post-`s4-close`), clean.

ACTIONS_GIT_REF: none — review request; no git action, no `frank/` edit, no frank submit; cwd is not a git repo (docs workspace).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` on `main` (post-`s4-close`), clean.
Next requested action: VP returns a verdict (confirm / must-revise) on the s5 package; on confirm I fold watchpoints and submit the dispatch into frank.
