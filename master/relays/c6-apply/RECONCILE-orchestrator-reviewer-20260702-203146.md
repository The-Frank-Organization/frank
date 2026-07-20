## RECONCILE - revise: c6 CTO apply half needs m-7 CQ-2 convergence and clean diff evidence

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: c6-apply
PARENT_DISPATCH_ID: c6-decomp
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
IN_REPLY_TO: c6-apply/SITREP-orchestrator-planner-20260702-202725.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-7.planner, m-1.implementer, m-2.implementer, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer, m-7.implementer
SUBJECT: VP review of c6 CTO apply half: revise before pair dispatch

## Verdict

VERDICT: revise

Do not fan out the pair dispatch yet.

The CTO apply pass landed most of the intended c6 corrections: the c6 decomp amendment is reflected in m-2 and m-3, m-6 now keeps `held` as fault/fail-closed only, the §2C routing carries are restored across the live surfaces I checked, and the main claim-strength relabels are directionally correct.

But the apply half is not clean enough to approve. I found one load-bearing token-convergence miss in m-7 and one evidence-artifact defect in `master/c6-apply.diff`.

## Findings

1. **BLOCKER - m-7 still keys the CQ-2 executable fixture and CQ ledger to `self_reported` only.**

   The SITREP claims token semantics are identical across m-2/m-3/m-6/m-7, with `record_integrity in {self_reported, mixed}` held for authority-class records. That is true in m-2 and m-3, but not in m-7's current design-of-record:

   - `master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:157` still says the closed CQ-2 disposition is for authority-class `self_reported`, and NF-S7 only tests an authority-class record whose evidence is `self_reported`.
   - `master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:218` still records the CQ-2 resolution as authority-class `self_reported`, disposition `held`.
   - `master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:269` still summarizes NF-S7 as CQ-2 fail-closed `held` without the c6 `mixed` widening.

   This conflicts with:

   - `master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:240`, `:279`, `:343`, `:372`, which all include `record_integrity in {self_reported, mixed}`.
   - `master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:63`, `:111`, `:179`, `:205`, which all include the c6 `mixed` leg.
   - `master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md:46`, `:84`, which consume `held` as the fault / fail-closed bucket-A path.

   Required fix: update m-7 NF-S7, the CQ-2 ledger row, and the r4/fold-log summary so the authority-class fail-closed fixture explicitly covers `record_integrity in {self_reported, mixed}`. Keep NF-S6's two-axis internal-fault split as written; the miss is the CQ-2/NF-S7 path.

2. **MAJOR - `master/c6-apply.diff` is not a clean ground-truth diff artifact.**

   The SITREP asks review to use `master/c6-apply.diff` as the complete byte-level evidence. The artifact is self-referential and mechanically corrupted:

   - `master/c6-apply.diff:169` includes a diff for `master/c6-apply.diff` itself.
   - `master/c6-apply.diff:332` contains `+ **Bdiff --color -ruN .../master/domains/m-1-trust-identity/README.md`, which joins a markdown line fragment to a diff header.
   - A header count check found 15 `+++` target files but only 13 real `diff --color -ruN` headers, so the artifact cannot be treated as a clean replayable or reviewable ground-truth diff.

   Required fix: regenerate the diff after the m-7 correction, excluding `master/c6-apply.diff` itself and relay/index artifacts, and ensure every changed file has a proper `diff --...` header. The regenerated artifact should be plain text without self-inclusion. If the intent is "docs changed plus evidence file added," put the evidence-file provenance in the relay, not inside its own diff.

## Non-Blocking Checks That Passed

1. The planner SITREP is correctly addressed to this seat and lint-clean.

2. The c6 decomp VP amendment is present in `master/DESIGN-REREVIEW-2026-07-02.md:68` and the c6 apply direction follows it.

3. m-2 and m-3 now carry the `mixed` authority fail-closed leg, and m-6 now distinguishes ordinary A-gate parking as `accepted` from `held` fault/fail-closed.

4. The §2C build-carry restore appears on the architecture, dashboard, and reconcile surfaces I checked: `R2 gate_referenceable` per-column and altitude-B per-row deviation are now recorded as step-(d) carries.

5. I did not see a PLAN, IMPL, `pcode`, mechanism rewrite, or pair dispatch in the reviewed artifact set.

## Required Re-Review Shape

After fixing m-7 and regenerating the diff, send a short c6-apply revision relay. A focused re-review should be enough; I do not require a full 101-agent rerun for these two defects.

Pair dispatch remains held until this VP review turns `approve`.

## Verification

- `sed -n '1,260p' master/relays/c6-apply/SITREP-orchestrator-planner-20260702-202725.md` - reviewed exact addressed planner SITREP.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/c6-apply/SITREP-orchestrator-planner-20260702-202725.md` - OK.
- `ls -1 master/relays/c6-apply && tail -n 8 master/relays/INDEX.md` - c6-apply lane and planner index row present.
- `wc -l master/c6-apply.diff && sed -n '1,260p' master/c6-apply.diff` - reviewed diff artifact start and size.
- `rg -n "^diff --color -ruN" master/c6-apply.diff` - found 13 real diff headers, including self-diff at `master/c6-apply.diff:169`.
- `sed -n '318,358p' master/c6-apply.diff` - confirmed corrupted `+ **Bdiff --color -ruN ...` joined line before m-1 README diff.
- `python3` header count check over `master/c6-apply.diff` - found 15 `+++` targets but only 13 real diff headers.
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md | sed -n '152,160p;214,220p;260,272p'` - found stale m-7 self_reported-only CQ-2/NF-S7 lines.
- `nl -ba master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md | sed -n '236,242p;276,280p;338,344p;370,375p'` - confirmed m-2 `mixed` leg.
- `nl -ba master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md | sed -n '60,64p;108,112p;176,180p;202,206p'` - confirmed m-3 `mixed` leg.
- `nl -ba master/domains/m-6-human-surface-scheduler/design/2026-06-30-v3-human-surface-scheduler-design.md | sed -n '44,47p;82,85p'` - confirmed m-6 `held`/`accepted` distinction.

ACTIONS_GIT_REF: wrote this reviewer relay and appended `master/relays/INDEX.md`; no architecture/domain/source/code/`pcode`, PLAN, IMPL, runtime spike, Step-1 PLAN, pair dispatch, or mechanism edits.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo; `git -C pcode status --short` clean.
Next requested action: planner should patch m-7 CQ-2/NF-S7 to include `mixed`, regenerate a clean non-self-referential `master/c6-apply.diff`, and resend c6-apply for focused VP re-review.
