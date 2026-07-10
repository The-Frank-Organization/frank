## DESIGN-REVIEW - s1-core.implementer re-review of s1-slice-1-design r2

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s1-core-design-r2-review-implementer
PARENT_DISPATCH_ID: s1-core-design-r2-review
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: s1-core.implementer
TO: s1-core.planner
CC: s1.orchestrator-planner, operator
IN_REPLY_TO: s1-core-design-r2-review/DESIGN-planner-20260703-152116.md
DESIGN_DOC_ID: s1-slice-1-design
DESIGN_REVIEW_VERDICT: approve
BASE: main@5622516
SUBJECT: r2 design re-review approve - five r1 blockers folded, grant narrowing routed

Phase:
Read-only DESIGN-REVIEW. I reviewed `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md` at main@5622516 against the addressed r2 review relay. I made no source/test/sprint-spec edits, no implementation branch, no commit, no PR, and no prototype.

Verdict:
approve.

Findings:
- No blocking findings.

Review scope:
- Incoming r2 review request: `.relays/s1/s1-core-design-r2-review/DESIGN-planner-20260703-152116.md`.
- Prior r1 review: `.relays/s1/s1-core-design-r1-review/DESIGN-REVIEW-implementer-20260703-151318.md`.
- Candidate design: `docs/sprints/2026-07-03-s1-slice-1/designs/s1-slice-1-design.md` at main@5622516.
- Locked surfaces checked for the folds: m-2 FieldSpec/lineage gate (`m-2 design:76,166-177`), m-7 conductor topology/pivot/recovery (`m-7 design:56-58,80-95,97-104,116-132`), and ARCHITECTURE J2 / owed-carry lines (`ARCHITECTURE.md:105-115,477-481`).

Fold verification:
1. Crash-window blocker closed. R2 makes the accepted gate record the durable derived-work intent, completes missing park/outbox follow-ups before channels open, dedupes by `gate_record_ref`, and adds C7 for kill after gate-accept plus re-crash during completion. Evidence: design lines 112-113, 129, 161. This preserves distinct park/outbox mutation classes instead of absorbing a compound pivot. Evidence: design lines 69, 190.

2. `grant` rendering blocker closed as an explicit routed narrowing. R2 now states S1 grants are operator/orchestrator-form only, pair-seat forms omit `grant`, and conditional pair-Planner delegated-dispatch rendering is not in S1; it routes that narrowing for orchestrator ratification in the design-completion SITREP. Evidence: design lines 91, 167, 185, 195; m-2 supports the operator/orchestrator grant-field reading at lines 166-177. This approval confirms the design is explicit and buildable; it does not itself ratify the product/protocol narrowing.

3. Authority-bearing blocker closed. R2 replaces the grant/gate-only rule with a pessimistic superset over MVP-visible fields: grant, gate, authority phases, authority values, and orchestrator-planner role. It maps those legs to m-2's authority-bearing classes and adds fixture H. Evidence: design lines 105, 168, 196; m-2 line 76.

4. B1 overclaim closed. R2 limits B1 to the no-tool-path property, moves direct record-file injection detection/quarantine to S2, narrows the direct-file leg to torn/staging cleanup, and states the D5 same-uid residual plainly. Evidence: design lines 10, 18, 129, 148, 197.

5. `gate_category` blocker closed. R2 locks the full frozen J2 default set in DESIGN, preserves `routing_escalation` as an owed non-member, removes the PLAN enum-selection item, and adds enum fixture legs. Evidence: design lines 89, 167, 181-182, 198; ARCHITECTURE lines 105-115 and 477-481.

Watch items:
- Binding-table persistence now uses the same stage/fsync/rename primitive as conductor file writes. Evidence: design line 45.
- MCP fallback invariants now preserve per-seat channels, stamped identity, and exactly `{submit, project, read}`. Evidence: design line 44.

Scope absorption check:
- OUT remains out: genesis/quarantine/GC/segment rotation, phase-0-to-4 machinery, migrators, observe, outbox drain/external send, and `model_name` are still excluded. Evidence: design lines 18, 98, 114-116, 129, 173, 179-185.
- The full FieldSpec registry and wider replay remain later work; R1 stays MVP-covered caught/obsolete/uncovered-S3. Evidence: design lines 98, 133-138.
- The grant narrowing must be ratified before PLAN proceeds; PLAN must not silently reintroduce pair-Planner grant rendering or treat it as implemented in S1 without an explicit orchestrator/operator decision. Evidence: design line 185 and r2 request lines 28, 39.

Approval boundary:
This approve is for `DESIGN_DOC_ID: s1-slice-1-design` r2 as a buildable S1 design. It grants no PLAN, IMPL, merge, or live-verification authority. It assumes the design-completion SITREP carries the blocker-2 narrowing to `s1.orchestrator-planner` / operator exactly as r2 states.

Actions:
- Wrote this DESIGN-REVIEW relay.
- Appended the corresponding routing row to `.relays/s1/INDEX.md`.

ACTIONS_GIT_REF: no source/test/sprint-spec edits claimed; relay artifact written at `.relays/s1/s1-core-design-r2-review/DESIGN-REVIEW-implementer-20260703-152445.md`; index row appended at `.relays/s1/INDEX.md`; final git status captured below.
FINAL_GIT_STATUS_SHORT: none - clean tree

Tests / verification:
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s1/s1-core-design-r2-review/DESIGN-planner-20260703-152116.md` passed.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s1/s1-core-design-r2-review/DESIGN-REVIEW-implementer-20260703-152445.md` passed.
- `git status --short --untracked-files=all` final output captured in `FINAL_GIT_STATUS_SHORT`.

Next requested action:
- `s1-core.planner` may send the design-completion SITREP with the blocker-2 narrowing called out for orchestrator/operator ratification, then hold for the next explicit phase dispatch.
