## Team m-5 — BOOT + c2 NARROW consumer-lens (Workflows & Archetypes)

ROLE: Orchestrator Planner
PHASE: AUDIT
AUTHORITY: read-only
DISPATCH_ID: c2-consumer-review-m-5
PARENT_DISPATCH_ID: c2-lock-prep
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — read-only consumer-review + bounded proposals; surface operator items in findings
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-5.planner, m-5.implementer
CC: master.orchestrator-reviewer, m-3.planner, m-4.planner, operator
BUNDLE_ID: m-5-workflows-archetypes
OWNER: m-5 (Workflows & Archetypes)

**Boot.** You are the **m-5 Workflows & Archetypes** pair on the standing `master` team (design-lead `m-5.planner` = Claude Opus 4.8 high; adversarial reviewer `m-5.implementer` = GPT-5.5). Read the charter `CLAUDE.md` + your domain `master/domains/m-5-workflows-archetypes/README.md`. Your durable domain: expansion-slot presets (topology+gate-set+human-mode), the tag-space, per-archetype observe invariants, authority-ceiling-at-spawn, sensor/actuator. **Your FULL design is c3.** You are booted now for a **narrow, bounded** c2 engagement only (VP-approved, `c2-lock-prep/RECONCILE-orchestrator-reviewer-20260629-212213.md`).

Phase scope — AUDIT / consumer-review (read-only). Read the two c2 design docs, review the seam from your vantage, and **propose** the three bounded Step-1 artifacts below. Not in scope: source/test edits, branches, commits, PRs, prototype code; and — **hard guardrail** — **NOT a full m-5 archetype-system design.** No PLAN/IMPL.

Pair method: m-5.planner leads the review + proposals; m-5.implementer runs an INDEPENDENT adversarial pass and challenges; reconcile into one m-5 consumer-review relay. The two seats are independent operator-relayed sessions; the planner does not spawn/simulate the implementer.

**The three bounded deliverables (VP-approved scope — do not exceed):**
1. **Seam-fit verdict — the m-3↔m-4 opaque archetype-tag interface.** Confirm the opaque-tag interface the two locked-pending designs key on actually fits your tag-space: m-3's done-predicate `slot_in` input (the per-archetype observe-invariant selector) + m-4's capability-prior **archetype key** + the **authority-ceiling-at-spawn**. Does the opaque-atom interface carry enough for m-5 to fill the semantics in c3 *without* a c2 re-cut? Flag any interface gap NOW (it is cheap here, expensive after the lock).
2. **Step-1 routing-template structures + the 1–3 shipped lineup (GL-4).** Propose the topology for the shipped templates: seats/roles, panes, gate-set, read-only-ness per template. (m-4 owns the routing-record mechanism; conductor-core owns pane-spawn via existing tmux/zellij/OS-terminal; **you own the template structure + lineup.**) Keep it to 1–3 concrete, useful templates.
3. **The side-question sensor archetype.** Spec the read-only, tool-blocked, single-turn **sensor** archetype that backs the interjection `/btw` analog (`ROADMAP.md` interjection cross-cutting rule + `references/jcode-ux-notes.md` §3): authority ceiling = read-only; tools = none; turns = 1; shares lane context; routed by m-4 to a cheap/fast model; answers on a separate surface without interrupting the lane.

**HARD BOUNDARY (VP guardrail — read twice).** SURFACE/PROPOSE these for the c2 lock; do **NOT** close or lock the **concrete tag-space, invariant selection, default per-archetype gate composition, full template-structure semantics, or authority-ceiling semantics** — those stay **m-5-owned and are locked in c3**. The c2 lock will *reserve* them to you. If a deliverable above tempts you toward the full system, stop at the Step-1 slice and flag the rest as c3.

Sources to review (E1 — cite file:line):
- m-3 design: `master/domains/m-3-observation-evidence/design/2026-06-29-observe-evidence-design.md` (the done-predicate `slot_in` archetype hook; the observe-as-send gate; the sensor's observe profile).
- m-4 design: `master/domains/m-4-routing-policy/design/2026-06-29-routing-policy-design.md` (the capability-prior archetype key §4/§8; the authority-ceiling cap §8; GL-4 routing templates §7).
- The reconciled seam: `c2-design-m3-m4-coord/COORD-planner-20260629-192916.md`.
- The locked c1 contract `master/ARCHITECTURE.md` (do not reopen); the interjection rule in `ROADMAP.md`; `references/codex-notes.md` (codex `collaboration-mode-templates` = single-agent behavioral presets, NOT team topologies — our multi-agent template layer is novel) + `references/jcode-ux-notes.md`.

Deliverable: a reconciled m-5 consumer-review relay (planner + implementer reconciled) under `c2-consumer-review-m-5` containing — the seam-fit verdict (+ any interface gap); the proposed Step-1 template lineup (1–3); the side-question sensor archetype spec; operator-judgment items or none; E1 cites. Surface-not-close per the hard boundary. No source changes, no PR. Include ACTIONS_GIT_REF + FINAL_GIT_STATUS_SHORT (read-only; cwd is not a git repo, structured unavailable form expected).

FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
