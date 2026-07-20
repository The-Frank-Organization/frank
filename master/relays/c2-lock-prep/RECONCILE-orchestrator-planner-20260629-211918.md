## RECONCILE — c2 lock-prep (orchestrator → VP): both pairs design-complete; m-5 narrow engagement + R2 ratification + GL-4 + consumer-lens plan

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c2-lock-prep
PARENT_DISPATCH_ID: none
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — operator has directed the m-5 narrow engagement + the GL-4 scope addition; this is VP review of the lock-prep plan + the R2 ratification
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-3.planner, m-4.planner

Partner — both c2 pairs are **design-complete + pair-approved**, holding (no self-advance to PLAN):
- m-4 Routing & Policy — implementer `approve` rev1 (`c2-design-m-4/DESIGN-REVIEW-implementer-20260629-203329.md`); doc `master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md`.
- m-3 Observation & Evidence — implementer `approve` (`c2-design-m-3/DESIGN-REVIEW-implementer-20260629-211003.md`); doc `master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md`.

Bringing you the **c2 lock-prep plan** before I execute. Four items; on your approval I dispatch the consumer-lens round and boot m-5 narrowly. No PROCEED-TO-PLAN / implementation / merge sought.

**Item 1 — m-5 disposition = NARROW ENGAGEMENT (operator-directed; supersedes the earlier pure-reservation lean).** The operator directed booting m-5 for a **bounded** c2 scope, not deferring it wholesale to c3. Scope of the narrow engagement:
- (a) define the **1–3 v3.0 routing-template structures + lineup** (the GL-4 scope addition — seats/panes/gate-set/read-only-ness);
- (b) **consumer-review the m-3/m-4 archetype seam** — confirm the opaque archetype-tag interface (m-3 done-predicate `slot_in`; m-4 capability-prior key + authority-ceiling) fits m-5's tag-space;
- (c) the **interjection side-question sensor archetype** (read-only authority ceiling, tool-blocked, 1-turn) just folded into `ROADMAP.md` as a forward requirement.
m-5's **full** archetype design stays **c3**; this is only the v3.0 slice. This **satisfies your lock prerequisite** (the "narrow m-5 consumer review on the draft designs" option) AND realizes GL-4. **Confirm:** the scope is bounded correctly and preserves m-5 ownership of the concrete tag-space + invariants.

**Item 2 — R2-boundary ratification (requested).** Both pairs aligned; I have verified it. Ratify the cross-domain R2-preservation choice: the silent-deviation block (`declared_deviated=false`, `deviated_observed=true`) lives in m-3's **generic observe-layer integrity-veto** (clean-tree class), **not** a model-derived `required_when`. **No model-derived predicate enters the m-2 schema gate** — the gate reads only the plain `declared_deviated` boolean; the deviation comparison is **bucket-vs-bucket** (does not read model identity); it rides **snapshot-provenance** (holds for opaque lanes). Sources: m-4 design §2 + §9; m-3 design §9; reconciled seam `c2-design-m3-m4-coord/COORD-planner-20260629-192916.md` + `…COORD-RECONCILE-planner-20260629-193400.md`.

**Item 3 — GL-4 sequencing (confirm roadmap-fit).** The operator directed v3.0 ship 1–3 selectable routing templates. It splits cleanly by owner: **m-4** owns the routing-template record mechanism (a pre-filled `routing_decision` with `template_ref` — clean Step-1, locks with m-4); **m-5** owns the template structures (now in Item-1's narrow scope); **conductor-core** owns the pane-spawn. **Confirm:** the pane-spawn **rides existing multiplexer infra (tmux / zellij / OS terminal), NOT our own TUI** (Step-4) — keeping it "ride existing runtimes" Step-1-consistent. Flag any roadmap concern.

**Item 4 — the consumer-lens round (plan).** On your approval I dispatch, in parallel, a consumer-review of the m-3/m-4 designs by:
- **m-6** (warm lens, already booted) — gate mechanism + egress + the ODB evidence-summary (m-3); the `routing` gate_category + ODB content for A-escalation (m-4);
- **m-5** (newly booted, narrow scope per Item 1) — the archetype seam + the GL-4 template lineup + the side-question archetype.
Their returns → I reconcile (a `c2-consumer-reconcile` analog) → the **c2 lock** (CTO/VP co-sign, like c1), folding the m-5 reservation-of-the-rest + the R2 ratification.

**FYI (no action) — interjection forward requirement folded into `ROADMAP.md`:** a new cross-cutting rule + a Step-3 exit test — steer / side-question / interrupt; steer+interrupt baseline; the side-question = the m-5 read-only sensor archetype (Item 1c) routed by m-4. Captured in `references/jcode-ux-notes.md` (negative look) + `references/codex-notes.md` (positive look); ratified-when-designed, gates nothing now.

Requested: review the lock-prep plan and return a verdict — confirm (1) the m-5 narrow-engagement scope is sound + bounded; (2) the R2-boundary is ratified; (3) the GL-4 pane-spawn is roadmap-consistent; (4) clear to dispatch the consumer-lens round + boot m-5. Revise / human-decision-required as needed, with cited changes.

ACTIONS_GIT_REF: read both pairs' design docs + the two implementer approvals + the reconciled seam; edited `master/README.md` (lock-prep state) + `ROADMAP.md` (the interjection cross-cutting rule + Step-3 exit test) + wrote `references/codex-notes.md` + `references/jcode-ux-notes.md` + this relay. Docs-workspace only; no code/source/pcode edits; cwd is not a git repo so a git ref is unavailable.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: VP returns the lock-prep verdict; on approve I dispatch the m-6 + m-5 consumer-lens round and boot m-5 narrowly.
