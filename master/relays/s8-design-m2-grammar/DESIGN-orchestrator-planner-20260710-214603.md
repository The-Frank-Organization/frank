## DESIGN — s8 pre-design dispatch to m-2.planner (sole author): the CONFIG/ATOM GRAMMAR design — the `layer_present:observe` activation semantics, the never-a-code-default rule, and the member grammar for the knob + catalog members; grill required; your Implementer design-reviews before the lock feeds the s8 PLAN

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: s8-design-m2-grammar
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — design dispatch; operator decisions the grill surfaces route to the operator as durable questions
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m2-grammar
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step2-prep/RECONCILE-orchestrator-reviewer-20260710-201733.md
FROM: master.orchestrator-planner
TO: m-2.planner
CC: operator, master.orchestrator-reviewer, m-2.implementer, m-7.planner
SUBJECT: supersedes the withdrawn `step2-prep/SITREP-…-193204` (report-only relays cannot direct design; consultation context only) — you are the SOLE AUTHOR of `s8-design-m2-grammar`; the m-7 host face is the SEPARATE artifact `s8-design-m7-config`; shared content is a named boundary contract, never a shared writer

**Your artifact (the m-2-owned grammar face, one author, your FROM):**
1. **The `layer_present:observe` activation semantics** (your intake shape, `step2-prep/SITREP-planner-…-013000` item 1): the atom stays your §5 step-gate mechanism; activation reads the governed flag — **never the `DefaultLayers()` code default** — and the rule "activation is a governed, visible act" is stated as a design invariant with its executable fixture named (the s5 dormancy sweep mirroring into activation tests).
2. **The member grammar** for the two new config members (the knob; the catalog): the §9-compatible shape — versioning/compat class of each member change, ignore-unknown behavior for older readers, MAJOR/MINOR classification rules — so m-7's host design has a grammar to bind to.
3. **The required-set consequences of activation:** what flips when observe goes live (Block A required_when firing; the A-1 re-render class) stated at your validator grain, with the migration story for records accepted pre-activation (immutability held; no re-validation).
4. **Boundary contract:** name `s8-design-m7-config` (the digest home, load-time flip, genesis composition — m-7's) where your grammar consumes it; do not author m-7's content.

**Process (binding):** the **grill** per the design-grill skill (codebase-answerable questions from your locked record; operator forks surfaced durable) → `GRILL_LOCK_ID` folded into the doc → **m-2.implementer returns `PHASE: DESIGN-REVIEW`** → the reviewed artifact to master. Master reconciles with `s8-design-m7-config`; **the s8 PLAN consumes only the reconciled locks.**

Next requested action: your grill → the design doc → your Implementer's DESIGN-REVIEW → the reviewed artifact to master.

ACTIONS_GIT_REF: none — design dispatch only.
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `2e1b4f0` (tag `s7-close`); cwd is not a git repo (docs workspace).
