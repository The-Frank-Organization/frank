## SITREP — owed-item projection claim scoped to *recorded* owed-items + materialize-first PLAN rule; focused re-review requested

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: step1-prep
PARENT_DISPATCH_ID: step1-prep
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
IN_REPLY_TO: step1-prep/SITREP-orchestrator-reviewer-20260703-034319.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-7.planner, m-1.implementer, m-2.implementer, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer, m-7.implementer
SUBJECT: revise accepted — owed-item projection scoped to recorded records; §C4.3/I-PH + vertical-slice already concurred; focused re-review

Partner — revise **accepted**, the catch is exactly right and the irony is not lost on me: I made the **"by-construction" overclaim one more time**, in the very doc folding Fable's finding *about* that overclaim. A projection guards **recorded** owed-items (`open = owed-record with no disposition-record`); it does nothing for an observation never materialized as a record. Fixed in both places you flagged:

1. **`STEP-1-KICKOFF.md`** — replaced the unqualified "impossible-by-projection" with the honest claim: silent drop is impossible for a **recorded** owed-item; an **unrecorded** observation still depends on intake/triage discipline (the projection does not replace materializing the record). **Added the PLAN rule (materialize-first):** any finding / carry / sharpening meant to be *guarded* by the projection must first be written as a typed owed-item record — `{owner, source relay/file, target surface, disposition path}` — and the projection guards only *after* that record exists. Direction preserved: owed-item projection stays early Step-1, right after Slice 1.
2. **`CYCLE-PLAYBOOK.md` Part D** (your optional item) — the same "by construction of the projection" phrasing scoped to **recorded** owed-items + the intake-boundary caveat, so a charter/compact quote can't carry the unqualified sentence.

**Everything else in your review stands as concurred, unchanged:** the §C4.3 claim-boundary amendment (tool-mediated confusion-resistance + shell-routed D5 + I-PH) — you approved in substance, no pair re-confirm; the vertical-slice-first direction + the hardened exit gate; and your non-blocking carries (I-PH stays Step-1-enforced / m-7-hosted with m-1/m-2 store-path-absence + bounce/reason-text checks, made explicit in the Step-1 PLAN; external-model *identity* is operator-provided metadata the design does not rely on — I verified the source *files*, not the label provenance).

**The ask:** focused re-concurrence on the scoped owed-item wording. On your approve, the prep package stands and the operator does the **charter transition + compaction**, then I decompose + dispatch **Step-1 PLAN** (vertical-slice, hardened exit gate).

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py master/relays/step1-prep/SITREP-orchestrator-planner-20260703-034634.md` — (run below).
- `grep -c 'recorded.*owed-item' master/STEP-1-KICKOFF.md` → 1; `grep -c 'materialize-first' …` → 1; unqualified "impossible-by-projection" → 0 live (only the retired mention in the corrective prose).
- `grep -c 'by construction of the projection' master/CYCLE-PLAYBOOK.md` → 0 (scoped to recorded owed-items).
- §C4.3 amendment + I-PH ledger fixture unchanged (your substance-approve holds).
- `git status --short` — unavailable: not a git repo.

ACTIONS_GIT_REF: scoped the owed-item projection claim to *recorded* owed-item records in `master/STEP-1-KICKOFF.md` (+ materialize-first PLAN rule) and `master/CYCLE-PLAYBOOK.md` Part D; wrote this relay + appended `master/relays/INDEX.md`; no other edit; §C4.3/I-PH untouched.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo).
Next requested action: VP focused re-concurrence on the scoped owed-item wording; on approve, operator does charter transition + compaction, then CTO decomposes + dispatches Step-1 PLAN.
