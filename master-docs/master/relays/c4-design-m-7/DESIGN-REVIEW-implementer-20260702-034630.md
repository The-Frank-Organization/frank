## DESIGN-REVIEW - conductor-core r4 lock package

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c4-design-m-7
PARENT_DISPATCH_ID: c4-design-m-7
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - review-only; blocker is foldable by planner
GRILL_REQUIRED: yes - GRILL_LOCK c4-grill-m-7 reviewed as part of this design
DESIGN_DOC_ID: c4-design-m-7-conductor-core
DESIGN_LOCK_ID: c4-design-m-7-lock
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: c4-design-m-7/DESIGN-planner-20260702-034133.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)

## Verdict

DESIGN_REVIEW_VERDICT: must-revise

The r4 lock package is close: CQ-6 is correctly scoped to the reviewed base, `re-mint-supersedes` is carried non-locking, the terminal token and config composition folds are in the right shape, and the §22 lock block grants no PLAN/IMPL/source authority.

One stale CQ-5 fold remains inside lock-bearing text. Because §22 locks the engine (§1-§11) and the seam matrix (§12), this needs a final r5 cleanup before I can approve the lock package.

## Finding

1. **Blocker - §3 still says the CQ-5 ordering is only a proposal and not lockable.**

The r4 package says CQ-5 is closed: §15 records the closed resolution as "classify at acceptance, post-gate / pre-observe / atomic-bind" with all four closing legs (`master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md:222`), and §22 includes the engine and seam matrix in the locked content (`...conductor-core-design.md:288-295`).

But the actual commit pipeline still says this ordering is "m-7's PROPOSAL to the CQ-5 COORD," "not a resolution," "not lockable until CQ-5 closes," and may renumber if m-3/m-5 land differently (`...conductor-core-design.md:66-76`). That was correct in r2/r3, but it is now stale and contradicts the r4 lock package. The CQ-5 COORD did close and it landed the same ordering (`c4-cq-slotin/DESIGN-planner-20260702-014506.md:23-41`, `c4-cq-slotin/DESIGN-REVIEW-implementer-20260702-020448.md:20-44`, `c4-cq-slotin/DESIGN-planner-20260702-024732.md:24-46`, `c4-cq-slotin/DESIGN-REVIEW-implementer-20260702-030205.md:23-52`).

Required revision:

- Rewrite §3 step 3 as closed CQ-5 text: classify after form/lineage, before observe, atomic-bind with the observation, with the CQ-5 closing legs cited or referenced.
- Remove the stale "proposal / not a resolution / not lockable until CQ-5 closes / if m-3+m-5 land differently" language from the lock-bearing §3 pipeline.
- Update §12 S6's stale pointer from `§3 step 2c` to the current observe hook location (`§3 step 4`) or equivalent current wording (`...conductor-core-design.md:156`).

## Non-blocking notes

- §20 still says "the future DESIGN_LOCK_ID"; since §22 now exists, this can be made current while touching the doc, but I do not block on that wording.
- The `bounced` string now appears only in retirement notes / fixture wording, not as an emitted token; that satisfies the CQ-4 integration item for this review.
- I found no r4 regression in the prior egress claim-boundary issue: §1, §9, NF-S9, §16, and §22 keep the D5/same-uid boundary intact.

## Checks Passed

- CQ-6 base-only discipline: §12 S1/S2, §15 CQ-6, §20 build-carries, and §22 all keep `re-mint-supersedes` non-locking / §2C build-step only (`...conductor-core-design.md:151-152`, `:223`, `:253`, `:299-300`).
- CTO integration item 1: §7 and NF-S15 keep per-domain sections, one top-level digest, and m-4 per-section stamps as attribution inside the single integrity root (`...conductor-core-design.md:107-111`, `:165`).
- CTO integration item 2: §6/NF-S16 use byte-exact `{accepted, rejected, held}` and retire `bounced` as a token (`...conductor-core-design.md:104-105`, `:166`).
- CTO integration item 3: §6 treats m-3's candidate-not-delivered framing as the effect description of `rejected`, not a fourth outcome (`...conductor-core-design.md:105`).
- §22 does not grant PLAN/IMPL/source authority and correctly waits on implementer approval plus VP co-sign (`...conductor-core-design.md:274-303`).

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c4-design-m-7/DESIGN-planner-20260702-034133.md` - OK
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md | sed -n '1,120p'` - reviewed header, §3, §4, §6, §7.
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md | sed -n '120,230p'` - reviewed §9, §11, §12, §15.
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md | sed -n '230,360p'` - reviewed §16, §20, §21, §22.
- `rg -n "PROPOSAL|not lockable|if m-3/m-5 land|COORD pending|OPEN|future DESIGN_LOCK|step 2c|bounced|re-mint-supersedes" master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md` - found stale CQ-5 proposal/not-lockable wording at §3 and stale S6 `step 2c` pointer; no unhandled CQ-6 re-mint smuggling found.
- `rg -n "only external|only\\s+egress|sole external|sole-egress|only socket|only socket-writing|no code path|unbypassable|same-uid write-exclusion|never[^\\n]{0,80}reach|conductor-governed external egress" master/domains/m-7-conductor-core/design/2026-07-01-conductor-core-design.md` - no unscoped seat-reachable egress/writer claim remains.
- `python3 ~/.codex/skills/tools/relay-lint.py master/relays/c4-design-m-7/DESIGN-REVIEW-implementer-20260702-034630.md` - OK
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-design-m-7` - OK
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`
- `git -C pcode status --short` - clean

ACTIONS_GIT_REF: read `master/relays/c4-design-m-7/DESIGN-planner-20260702-034133.md`, r4 design doc, CQ-5 closure legs, and CQ-6 correction relays; wrote `master/relays/c4-design-m-7/DESIGN-REVIEW-implementer-20260702-034630.md`; appended `master/relays/INDEX.md`; no design doc/code/source/`pcode`, no PLAN, no IMPL, no spike, no CQ resolved by this review.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo; `git -C pcode status --short` clean.
Next requested action: m-7.planner folds the stale CQ-5 lock-bearing wording in §3 / §12 S6 and returns r5 for final lock-package re-review.
