## DESIGN-REVIEW - conductor-core design r1

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c4-design-m-7
PARENT_DISPATCH_ID: c4-design-m-7
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - review-only; blockers are foldable by planner
GRILL_REQUIRED: yes - GRILL_LOCK c4-grill-m-7 reviewed as part of this design
DESIGN_DOC_ID: c4-design-m-7-conductor-core
DESIGN_REVIEW_VERDICT: must-revise
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)

## Verdict

DESIGN_REVIEW_VERDICT: must-revise

The design is directionally strong: the claim boundary is mostly held, the Package-A pivot is concrete, recovery is sequenced, the seam matrix has real negative fixtures, and the CQ ledger correctly prevents design-lock while upstream policy seams remain open.

Two issues block approval: one pipeline ordering contradiction and one claim-boundary regression in the local-outbox/egress wording. A third lock-hygiene issue should be folded in the same pass.

## Findings

1. **Blocker - `slot_in` is classified after the observe gate that needs it.**

The commit pipeline runs m-3 `observe_gate()` in step 2c, then classifies `slot_in` only in step 3 "after gates pass, before the pivot" (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:67-74`). That conflicts with the m-3/m-5 contract: m-3 says observe runs as the fixed second pre-flight inside atomic submit (`master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:31-35`), and the `slot_in` done-predicate reads the conductor-classified work-archetype (`master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:100-102`; `master/domains/m-5-workflows-archetypes/design/2026-06-30-v3-archetype-system-design.md:17-21`). As written, the Step-2 observe hook cannot select the right predicate before `slot_in` exists, while the doc also says done-predicate selection reads a committed value. The revision needs one coherent ordering: either classify after form/lineage and before observe-predicate selection, then pivot the record carrying both `slot_in` and observation; or keep observe inert until CQ-5 closes and say no Step-2 observe gate is lockable in this design yet. The current ordering silently resolves CQ-5 while still labeling it open.

2. **Blocker - local-outbox wording overclaims sole egress under attach.**

The design correctly states the honest residual in §8.4: host runtimes still hand seats Bash, and shell-bearing seats can act outside MCP (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:121-122`). But §9 and NF-S9 then say the conductor outbox is "the only egress" and "the only socket writer" with "no code path" sending unscanned bytes externally (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:127-129`, `:156`). That reads as a forbidden adversarial/sole-external-sender claim, not just a governance-surface claim. The deployment grill permits only interface-level confusion resistance: seats act through `submit`/`project`/`read`, malicious same-uid bypass is explicitly out of scope, and local-outbox-only is a retained mechanism, not an adversarial containment proof (`master/GRILL-LOCK-deployment-fork-2026-07-01.md:35-49`, `:71-75`). Revise §9/NF-S9 to scope the statement to conductor-governed external send: "the only conductor-owned/governance egress path" or "the only m-7 outbox drain path." Add the same D5 residual beside it so F8 cannot pass while the stronger wording remains.

3. **Lock-hygiene - the status line omits CQ-8, but CQ-8 is design-lock blocking later.**

The design header says design-lock is gated on CQ-1..CQ-6+4b (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:5-7`), but §15 adds `CQ-8 INDEX derived-authority` as design-LOCK blocking (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:210-224`). Future readers tend to scan the header first. Fold the header/status line to include CQ-8, or explicitly mark CQ-8 as a post-dispatch m-1 COORD newly added by the design and part of the no-lock gate. This is not a design-start blocker, but it is a review blocker because approval would otherwise bless a stale lock package summary.

## Non-blocking notes

- I do not object to locking Package A, the canonical-record rename pivot, derived projections, redo replay, append-only INDEX repair posture, or the durable FIFO shape, provided CQ-8 is routed to m-1 before design-lock.
- The terminal-state enum structure `{accepted, rejected, held}` is acceptable as an m-7 engine shape only with the existing CQ-4 caveat preserved: final token spelling and m-2 field home remain owner-controlled.
- The "no config path/value in seat-deliverable surfaces" fixture should be interpreted as raw conductor-internal config/store/outbox/operator-channel paths and effective config values, not ordinary evidence citations to relay/design files. If the planner touches §8.4 or NF-S18, adding that qualifier would prevent a false-positive fixture later.

## Required revision conditions

- Move or qualify `slot_in` classification so the Step-2 observe predicate never needs a value that is classified only after observe has already run.
- Scope all local-outbox/egress exclusivity claims to the conductor-governed/MCP surface and restate the attach D5 residual anywhere "only egress" appears.
- Update the status/lock-gate summary to include CQ-8 or explicitly explain why CQ-8 is not part of the lock gate.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-design-m-7/DESIGN-REVIEW-implementer-20260702-002302.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-design-m-7` - OK
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`
- `git -C pcode status --short` - clean

ACTIONS_GIT_REF: read design doc `master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md`, c4 design dispatch and planner relay, relevant c4 audit reconcile relays, protocol, and m-3/m-5/grill source anchors; wrote `master/relays/c4-design-m-7/DESIGN-REVIEW-implementer-20260702-002302.md`; appended `master/relays/INDEX.md`; no code/source/`pcode`, no PLAN, no IMPL, no spike, no CQ resolved.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo; `git -C pcode status --short` clean.
Next requested action: m-7.planner folds the three revision conditions and returns a revised DESIGN doc for re-review.
