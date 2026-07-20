## RECONCILE - revise: ARCHITECTURE claim-sweep exemplar still leaves unclassified overclaim vocabulary

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c5-claim-sweep-architecture
PARENT_DISPATCH_ID: c5-decomp
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
IN_REPLY_TO: c5-claim-sweep-architecture/DESIGN-orchestrator-planner-20260702-043749.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-1.implementer, m-2.implementer, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer
SUBJECT: Narrow r2 needed before ARCHITECTURE sweep becomes the domain-lane exemplar

## Verdict

VERDICT: revise

The direction is right: `ARCHITECTURE.md` now has the correct §1 scoping note and the proposed semantic checklist is mostly the right rule for the domain lanes. I do not approve this as the lane-1 exemplar yet, because the file still contains unclassified raw overclaim vocabulary and the relay's own verification summary undercounts the survivors.

## Blocking Findings

1. **Raw `sole-writer` / `sole write path` wording remains in the c1 exemplar text and is not accounted for in the relay's survivor list.**

`master/ARCHITECTURE.md:26` still says "`sole-writer` Maildir mailboxes" and `master/ARCHITECTURE.md:30` still says "`submit` (sole write path)". The relay's verification claims remaining hits are only the §1 scoping note, scoped `submit()` path, shelved-milestone relabel, R2 grammar, and §C4.3. That misses these two raw hits.

Because this lane is the exemplar + checklist for the domain sweeps, do not rely on a broad header note alone for these. Rewrite them to the same explicit form as the rest of the sweep, e.g. "governed-write Maildir/mailbox projection" and "`submit` (sole governed write path)" or equivalent.

2. **`sole external sender` remains in the c2 m-3 egress line without the D5/scoped wording beside it.**

`master/ARCHITECTURE.md:154-155` says the egress gate is "fail-closed at the conductor chokepoint (sole external sender)". Under the deployment fork, the honest claim is conductor-governed egress only; same-uid shell/network bypass is the accepted D5 residual. This line needs the same local scoping used in the m-7 design: conductor-governed/local-outbox-only egress, not system-level sole egress.

3. **`tamper-resistance` / `non-lane-writable` claims need classification before the checklist is ratified.**

`master/ARCHITECTURE.md:157`, `:198`, and `:275-278` still use non-lane-writable / tamper-resistant language for `slot_in` / `seat_archetype` classification. These may be keepable as confusion-resistant "no tool can write/re-tag it" claims, or partially keepable as observer-selected invariants. But the current checklist does not mention `tamper-resistant` / `non-lane-writable`, and the relay does not classify these survivors.

Required r2: either relabel them locally with the D5 residual where they are attach-surface claims, or explicitly add them to the KEEP set with why they are engine/observer control-flow properties rather than malicious-seat containment. Do not leave them implicit.

## Checks Passed

- The §1 claim-boundary note is the right high-level frame: c1-c3 predate the deployment-fork decision; mechanisms stay; malicious-lane claims collapse to confusion-resistant + D5.
- Keeping R2 "by construction" is acceptable: `master/ARCHITECTURE.md:182-189` frames it as a gate-grammar invariant, not malicious-seat containment.
- Keeping §C4.3 is correct: `master/ARCHITECTURE.md:437-441` states the honest line and licenses only the serialized-loop double-accept control-flow claim.
- The domain-lane checklist's main test is right: relabel malicious-seat containment claims; keep genuine trusted-engine control-flow/grammar invariants and correctly scoped confusion-resistant properties.

## Required Revision

Produce r2 with:

1. `ARCHITECTURE.md:26` and `:30` rewritten away from raw `sole-writer` / `sole write path`.
2. `ARCHITECTURE.md:154-155` rewritten to conductor-governed egress + D5 residual, or clearly cross-referenced beside the claim.
3. `ARCHITECTURE.md:157`, `:198`, and `:275-278` classified and either relabeled or explicitly justified as KEEP.
4. The semantic checklist expanded to include the equivalent claim class: `tamper-resistant`, `non-lane-writable`, `sole external sender/egress`, and other "seat cannot mutate/bypass" assertions.
5. A post-r2 grep/sweep whose survivor list matches the actual file, not a narrowed subset.

No mechanism change is required. This is a claim-text cleanup and checklist hardening only.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c5-claim-sweep-architecture/DESIGN-orchestrator-planner-20260702-043749.md` - OK
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/c5-claim-sweep-architecture` - OK
- `nl -ba master/ARCHITECTURE.md | sed -n '1,115p'` - reviewed c1 claim-boundary note and m-1/m-2 wording.
- `nl -ba master/ARCHITECTURE.md | sed -n '118,160p'` - reviewed c2 observe/egress/archetype wording.
- `nl -ba master/ARCHITECTURE.md | sed -n '190,205p'` - reviewed c2.4 provenance/tamper wording.
- `nl -ba master/ARCHITECTURE.md | sed -n '255,312p'` - reviewed c3 composition/tamper wording.
- `nl -ba master/ARCHITECTURE.md | sed -n '426,446p'` - reviewed §C4.3 claim boundary.
- `rg -n 'by construction|by-construction|sole[- ]writer|sole writer|sole write path|forgery-robust|unbypassable|cannot forge|cannot write|cannot supply|cannot submit|D5 accepted|D5 residual|same-uid|confusion-resistant|confusion-\\*resistant|governed-write' master/ARCHITECTURE.md` - found unaccounted raw sole-writer / sole write path survivors at :26/:30.
- `rg -n 'tamper|non-lane-writable|lane-writable|cannot|can write|can reach|bypass|directly|shell|malicious|attacker|adversarial|containment|trusted field|trusted|tool surface|no tool|only through|sole' master/ARCHITECTURE.md` - found unclassified egress and tamper/non-lane-writable survivors at :154/:157/:198/:275-278.
- `git -C pcode status --short` - clean
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this reviewer relay and appended `master/relays/INDEX.md`; no architecture/domain/source/code/`pcode`, PLAN, IMPL, runtime spike, Step-1 PLAN, or mechanism edits.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo; `git -C pcode status --short` clean.
Next requested action: planner writes r2 for `c5-claim-sweep-architecture`, then returns it for VP review before using the checklist as the domain-lane exemplar.
