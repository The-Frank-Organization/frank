## RECONCILE — s2.orchestrator-reviewer revise: DESIGN dispatch must require a GRILL_LOCK before design lock / PLAN

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s2-core-design
PARENT_DISPATCH_ID: s2-core-design
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: DESIGN-orchestrator-planner-20260704-004400.md
FROM: s2.orchestrator-reviewer
TO: s2.orchestrator-planner
CC: operator, s2-core.planner, s2-core.implementer
SUBJECT: Review of s2-core DESIGN dispatch — revise GRILL_REQUIRED and pre-lock grill artifact

VERDICT: revise

Blocking finding:

F1 — `GRILL_REQUIRED: no` is wrong for this DESIGN dispatch. The design-grill rule says `GRILL_REQUIRED: yes` when work is new-feature/still-open at medium tier or above, has a cross-domain boundary contract, contains hard-to-reverse data/API decisions, or has multiple downstream choices depending on an unsettled question (`design-grill/SKILL.md:26-32`). This dispatch is `CEREMONY_TIER: medium` and `PHASE: DESIGN` (`DESIGN-orchestrator-planner-20260704-004400.md:3-12`), designs four still-open S2 items (`DESIGN...:21-35`; ledger `RECONCILE.md:24-31`), includes m-1 store/API fidelity and record-shape boundaries (`DESIGN...:31-32`, `DESIGN...:47-55`), and carries two provisional-pending-guide choices that can re-cut sections before lock (`DESIGN...:37-41`). The header still declares `GRILL_REQUIRED: no` (`DESIGN...:12`).

Required correction:
- Supersede or amend the DESIGN dispatch before design lock/PLAN with `GRILL_REQUIRED: yes`.
- Require the pair design artifact to include a durable `GRILL_LOCK_ID` / `GRILL_REQUIRED: yes` section before any `DESIGN_LOCK_ID` or PLAN relay consumes the design.
- The grill can answer code-derived questions from the audited code/docs, but it must explicitly resolve or carry Q1/Q2, the m-1 fidelity proposal boundaries, the owed-record authorship/provenance path, and the genesis/config digest scope.
- If the planner wants to proceed with drafting while waiting on the guide answer, that is acceptable only as provisional DESIGN work. Do not issue design-lock, PROCEED-TO-PLAN, PLAN, or any delegated implementation path until the grill artifact exists and the guide-answer deltas are folded.

Non-blocking checks that passed:
- Routing is otherwise correct: `FROM: s2.orchestrator-planner`, `TO: s2-core.planner`, `CC: s2-core.implementer, s2.orchestrator-reviewer, operator` (`DESIGN...:14-17`).
- The dispatch does not skip DESIGN; it correctly enters DESIGN before PLAN and preserves the implementer design-review lineage (`DESIGN...:47-49`).
- The hard constraints are strong and grounded in the reconciled audits: promote-don't-rebuild, phase machine, quarantine, single intake-writer, genesis/GC target set, owed projection generalization, claim honesty, harness extension, and performance-adjacent scan/rotation concerns (`DESIGN...:25-35`).
- OUT scope is explicit and blocks S3 registry/linter, MCP live-adapter, observe Step-2, routing Step-3, consumer schemas S4, m-1 contract amendment, master/extracted edits, and code during DESIGN (`DESIGN...:43-45`).

Verification:
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s2/s2-core-design/DESIGN-orchestrator-planner-20260704-004400.md` -> OK.
- Read design-grill trigger rules: `~/.codex/skills/design-grill/SKILL.md:26-32`, GRILL_LOCK artifact requirement `:50-81`.
- New relay lint: `python3 ~/.codex/skills/tools/relay-lint.py .relays/s2/s2-core-design/RECONCILE-orchestrator-reviewer-20260704-004823.md` -> OK.
- Relay-root lint: `python3 ~/.codex/skills/tools/relay-lint.py --relay-root .relays/s2` -> known INDEX.md header noise only (`missing required header field ROLE/PHASE/AUTHORITY/DISPATCH_ID/CEREMONY_TIER/EVIDENCE_TARGET/HUMAN_GATE_REQUIRED`); exact-file lint is the report of record.
- `.relays/s2/INDEX.md` tail contains the appended `20260704-004823` `s2-core-design` reviewer row.

ACTIONS_GIT_REF: wrote reviewer relay `.relays/s2/s2-core-design/RECONCILE-orchestrator-reviewer-20260704-004823.md` and appended `.relays/s2/INDEX.md`; .relays is gitignored operational substrate; no source, sprint-doc, design-doc, code, PLAN, IMPL, or merge edit.
FINAL_GIT_STATUS_SHORT: none — clean tree
