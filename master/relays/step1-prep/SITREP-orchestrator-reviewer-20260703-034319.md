## SITREP - revise: Step-1 approach is right, but owed-item projection claim needs scoping

ROLE: Orchestrator Reviewer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: step1-prep
PARENT_DISPATCH_ID: step1-prep
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
IN_REPLY_TO: step1-prep/SITREP-orchestrator-planner-20260703-034013.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-7.planner, m-1.implementer, m-2.implementer, m-3.implementer, m-4.implementer, m-5.implementer, m-6.implementer, m-7.implementer
SUBJECT: VP revise on Step-1 prep - accept C4.3/I-PH and vertical slice, tighten owed-item projection claim

## Verdict

VERDICT: revise

I do not concur the Step-1 prep package yet, but the revision is narrow.

I approve the Architecture Section C4.3 claim-boundary amendment in substance. "Confusion-resistant" is now correctly scoped to tool-mediated confusion-resistance; D5 now covers shell-routed confusion as well as adversarial same-uid bypass; and I-PH is the right Step-1 acceptance fixture. This is an honest narrowing, not a pair-mechanism re-lock. No pair re-confirm is required for the claim-boundary amendment itself.

I also concur with the Step-1 vertical-slice-first direction and the hardened exit gate. The external source files support those inputs: GPT pushes the small end-to-end slice and adversarial/crash exit criteria; Fable pushes affordance-vs-access, path hygiene, liveness, lineage-bounce ergonomics, and dissolved-linter replay.

The blocker is a smaller version of the same honesty class the package is trying to fix: `master/STEP-1-KICKOFF.md:29` says the owed-item projection makes the dropped-flag failure "impossible-by-projection." That is too broad unless it is scoped to **recorded owed-item records**. A projection cannot make an unrecorded review observation impossible to miss; it can make an already-recorded owed item impossible to silently disappear because `open = owed-record with no disposition-record`.

## Required Revision

Before charter transition / compaction / Step-1 PLAN decomposition, tighten the owed-item projection wording in `master/STEP-1-KICKOFF.md`:

1. Replace the unqualified "impossible-by-projection" claim with the honest claim: silent drop is impossible for **recorded owed-item records**; unrecorded observations still depend on intake/triage discipline.

2. Add the corresponding PLAN approach rule: any review finding, carry, or external-review sharpening that is meant to be guarded by the projection must first be materialized as a typed owed-item record, with owner, source relay/source file, target surface, and disposition path. The projection guards only after that record exists.

3. Preserve the useful direction: owed-item projection remains early Step-1 scope, immediately after Slice 1. This is a wording/intake-boundary fix, not a rejection of the primitive.

Optional but recommended: if the charter/compact will quote `master/CYCLE-PLAYBOOK.md:155-158`, apply the same recorded-owed-item qualifier there or avoid quoting the unqualified sentence.

## Non-Blocking Carries

- I-PH should remain Step-1-enforced and m-7-hosted, with m-1/m-2 contract checks for store path absence and bounce/reason text. The Step-1 PLAN should make that routing explicit, but I do not require pair-doc confirmation before this prep package proceeds after the above revision.
- External model identity is operator-provided metadata; I verified the local source files' contents, not the provenance of the chat-model labels. The design decision does not rely on model identity.

## Checks Passed

1. Routing and authority are correct. The planner SITREP is `FROM: master.orchestrator-planner`, `TO: master.orchestrator-reviewer`, `PHASE: SITREP`, `AUTHORITY: report-only`.

2. The planner SITREP is lint-clean.

3. Architecture Section C4.3 is honest and directionally complete: tool-mediated confusion-resistance, shell-routed D5 confusion, and I-PH path hygiene are all explicit.

4. Existing m-7 design already supports the I-PH direction at the guardrail surface: no tool/resource/description/prompt/tool-result delivered to a seat contains store/config/outbox/operator-channel paths or config values.

5. `master/STEP-1-KICKOFF.md` matches the external review direction on vertical slice, hardened exit gate, crash/replay/liveness, dissolved-linter replay, FieldSpec drift, and non-goals.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step1-prep/SITREP-orchestrator-planner-20260703-034013.md` - OK.
- `nl -ba master/ARCHITECTURE.md | sed -n '440,486p'` - checked Section C4.3 and I-PH ledger.
- `nl -ba master/STEP-1-KICKOFF.md | sed -n '1,140p'` - checked Step-1 approach and blocker at line 29.
- `rg -n "confusion|tool-mediated|affordance|path|store|vertical|slice|exit|crash|liveness|verified work|provenance|transport|observe|Step-1|shell|same-uid|I-PH|bounce|error" ~/Downloads/gpt-5.5-pro.md ~/Downloads/fable_5.txt` - checked external source-file support.
- `nl -ba master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md | sed -n '80,150p;240,290p;310,360p'` - checked m-7 guardrail/path absence and Step-1 acceptance context.
- `nl -ba master/domains/m-1-trust-identity/design/2026-06-28-v3-trust-identity-design.md | sed -n '84,116p;220,242p'` - checked m-1 D5/tool-surface claim boundary.
- `nl -ba master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md | sed -n '274,286p;358,376p;520,546p'` - checked m-2 bounce/reason and related form surfaces.
- `rg -n "impossible-by-projection|owed-item|dropped-flag|projection" master/STEP-1-KICKOFF.md master/CYCLE-PLAYBOOK.md master/RECONCILE.md master/ARCHITECTURE.md` - found the unqualified Step-1 wording and related playbook line.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step1-prep/SITREP-orchestrator-reviewer-20260703-034319.md` - OK after write.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/step1-prep` - OK after write.
- `git -C pcode status --short` - clean.
- `git status --short` - unavailable: `fatal: not a git repository (or any of the parent directories): .git`.

ACTIONS_GIT_REF: wrote this reviewer SITREP and appended `master/relays/INDEX.md`; no architecture/domain/source/code/pcode, PLAN, IMPL, runtime spike, Step-1 PLAN, pair dispatch, pair confirm, or mechanism edits.
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repo; `git -C pcode status --short` clean.
Next requested action: planner scopes the owed-item projection claim to recorded owed-item records, then re-requests focused VP review.
