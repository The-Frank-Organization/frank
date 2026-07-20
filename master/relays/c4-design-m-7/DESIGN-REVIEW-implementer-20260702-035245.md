## DESIGN-REVIEW - conductor-core r5 lock package

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c4-design-m-7
PARENT_DISPATCH_ID: c4-design-m-7
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - pair-side lock-package review only; VP co-sign still required
GRILL_REQUIRED: yes - GRILL_LOCK c4-grill-m-7 reviewed as part of this design
DESIGN_DOC_ID: c4-design-m-7-conductor-core
DESIGN_LOCK_ID: c4-design-m-7-lock
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: c4-design-m-7/DESIGN-planner-20260702-035009.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
BUNDLE_ID: m-7-conductor-core
OWNER: m-7 (Conductor-Core)

## Verdict

DESIGN_REVIEW_VERDICT: approve

The r5 lock package resolves the r4 blocker. I approve DESIGN_LOCK_ID `c4-design-m-7-lock` as the m-7 pair-side final lock-package pass for DESIGN_DOC_ID `c4-design-m-7-conductor-core`.

Scope of this approval: design-review only. It does not make the design lock effective without the VP co-sign, and it grants no PLAN, IMPL, code/source/`pcode`, spike, merge, or runtime authority.

## Review

The prior blocker was that the lock-bearing §3 pipeline still described the CQ-5 `slot_in` ordering as a proposal / not lockable while §15 and §22 treated CQ-5 as closed. r5 fixes that contradiction:

- §3 step 3 now states the CLOSED CQ-5 ordering: classify at acceptance, after form/lineage gates, before observe, atomic-bind with the observation, with the four CQ-5 closing legs cited (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:66-76`).
- §12 S6 now points the observe hook at §3 step 4, with predicate selection by step-3 classification, and CQ-5 is marked CLOSED (`...v3-conductor-core-design.md:156`).
- §12 S12 carries the same closed CQ-5 ordering and fixture bite (`...v3-conductor-core-design.md:162`).
- The remaining "proposal / not lockable / step 2c" hits are historical fold-log text in §21, not live mechanism text (`...v3-conductor-core-design.md:258`, `:274`).

The r4 checks-passed set remains intact:

- CQ-6 is scoped to the reviewed base; `re-mint-supersedes` stays non-locking / §2C build-carry only (`...v3-conductor-core-design.md:151-152`, `:223`, `:253`, `:301`).
- CQ-4b config composition remains one conductor-composed artifact with a single top-level digest and m-4 per-section stamps inside that artifact (`...v3-conductor-core-design.md:107-111`, `:165`).
- CQ-4 terminal tokens remain byte-exact `{accepted, rejected, held}`, and `bounced` remains retired as a value token (`...v3-conductor-core-design.md:104-105`, `:166`, `:220`).
- m-3's candidate-not-delivered framing remains an effect description of `rejected`, not a fourth outcome (`...v3-conductor-core-design.md:105`).
- The D5/same-uid egress boundary remains scoped to conductor-governed surfaces (`...v3-conductor-core-design.md:41-44`, `:130-132`, `:159`, `:231-232`).
- §22 grants nothing and stays conditional on this approval plus VP co-sign (`...v3-conductor-core-design.md:276-304`).

## Non-blocking lock-hygiene note

The status line and §22 still contain stale revision labels (`DESIGN-LOCK-READY r4`, `this doc, r4`, and "r4 = CQ-closure folds only"). I am not blocking on that metadata because the r5 relay, §21 r5 fold-log, DESIGN_DOC_ID, and DESIGN_LOCK_ID unambiguously identify the reviewed content, and the stale labels do not reintroduce a mechanism or authority contradiction. If the VP wants the lock certificate text polished before co-sign, the cleanup is label-only and does not need another m-7 pair re-review unless mechanism text changes.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-design-m-7/DESIGN-planner-20260702-035009.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c4-design-m-7` - OK
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md | sed -n '1,130p'` - reviewed header, §1, §2, §3, §4, §6, §7, §8.
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md | sed -n '130,230p'` - reviewed §9, §11, §12, §15.
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md | sed -n '230,330p'` - reviewed §16, §20, §21, §22.
- `rg -n "PROPOSAL|not lockable|land a different ordering|COORD pending|step 2c|future DESIGN_LOCK|DESIGN_DOC_ID: c4-design-m-7-conductor-core \\(this doc, r4\\)|DESIGN-LOCK-READY \\*\\*r4" master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md` - previous blocker strings now confined to historical fold-log text; stale r4 labels noted as non-blocking metadata.
- `rg -n "only external|only\\s+egress|sole external|sole-egress|only socket|only socket-writing|no code path|unbypassable|same-uid write-exclusion|never[^\\n]{0,80}reach|conductor-governed external egress|bounced|re-mint-supersedes|\\{accepted, rejected, held\\}" master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md` - no unscoped seat-reachable egress/writer claim or terminal-token regression found; expected scoped/historical hits only.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c4-design-m-7/DESIGN-REVIEW-implementer-20260702-035245.md` - OK
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`
- `git -C pcode status --short` - clean

ACTIONS_GIT_REF: read `master/relays/c4-design-m-7/DESIGN-planner-20260702-035009.md`, r5 design doc, relevant architecture seam lines, and prior r4 review; wrote `master/relays/c4-design-m-7/DESIGN-REVIEW-implementer-20260702-035245.md`; appended `master/relays/INDEX.md`; no design doc/code/source/`pcode`, no PLAN, no IMPL, no spike, no CQ resolved by this review.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo; `git -C pcode status --short` clean.
Next requested action: VP design-lock co-sign for `c4-design-m-7-lock`, with optional label-only cleanup of stale r4 revision text before/with co-sign.
