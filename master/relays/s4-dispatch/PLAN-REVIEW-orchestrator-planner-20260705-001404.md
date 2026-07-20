## PLAN-REVIEW request — VP pre-handoff gate on the s4 wire-up package (dispatch + boot held; the renumber rides this review)

ROLE: Orchestrator Planner
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s4-dispatch
PARENT_DISPATCH_ID: s4-dispatch
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — pre-handoff VP review (the S3 cadence, now standing); the operator holds the handoff until your verdict
IN_REPLY_TO: s4-dispatch/PLAN-orchestrator-planner-20260705-000914.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator
SUBJECT: review the s4 wire-up dispatch + boot before handoff — the fork election + renumber (wire-up = s4; old Section-4 → s5, built over the wired conductor), the first E3 gate, the §7 inheritance, the OUT fence

**The ask.** Pre-handoff review of the Slice-4 package — both artifacts cut, lint-clean both modes, INDEX'd, **held** (nothing relayed):
- Dispatch: `master/relays/s4-dispatch/PLAN-orchestrator-planner-20260705-000914.md`
- Boot: `master/relays/boot/s4-boot-orchestrator-planner/SITREP-orchestrator-planner-20260705-000914.md`

**What rides this review (beyond your standing checks):**
1. **The fork election + renumber.** The operator elected the wire-up as s4 (chartered as the operator's fork in the VP-approved s3-dispatch); the decomposition's old Section-4 (consumer schemas + §C4 fixtures) becomes **s5**, with the operator's sequencing intent recorded: s5 is built *with s4 in use* (its registry rows land as real §7 config-change records; its relays generate the first usage data). Confirm the renumber's record is adequate (dispatch + kickoff + dashboard) and nothing in the step1-plan approval is silently contradicted.
2. **The first E3 gate.** The exit gate claims E3 **scoped to transport/provenance only** (a real two-session relay, ideally Claude Code ↔ Codex; adversarial + crash legs live; "done" stays `self_reported`). Is the E3 scoping tight enough that no verified-work implication leaks?
3. **The §7 inheritance.** The config-change record lands here per the s3-scope-q1 ruling (m-7 guides · m-1 fidelity on the `record_kind` · crash-matrix gains the class · operator-authorized per locked §7 :109 · OI-S3-CONFIG-CHANGE dispositioned through the live mechanism). Check the conditions transferred faithfully.
4. **The seam I flagged as the slice's live design question:** second-connect semantics (a credential connecting while already connected — supersede vs reject; m-1 identity × m-7 lifecycle jointly). The dispatch routes any *new contract surface* to the amendment path. Is that guard sufficient, or should the dispatch pre-constrain the answer?
5. **The OUT fence** — federation explicitly horizon/zero-pre-work; s5 content; steer/interrupt beyond host-native; authority replacement (transport only). Tight enough for a cold reader?
6. **I-PH across the shim boundary** — promoted to an adversarial gate line (no store/config/socket path in any MCP-surfaced text). Sufficient, or does the shim need its own path-hygiene fixture class named?

## Verification
- Both artifacts lint OK exact-file + `--relay-root` at cut (run in the cutting session); INDEX rows appended; kickoff + dashboard renumber notes applied alongside.
- Baseline: `frank/` on `main`, clean, tag `s3-close`; riding owed item = `OI-S3-CONFIG-CHANGE` alone (`master/RECONCILE.md` § S3).

ACTIONS_GIT_REF: wrote this review request + an `INDEX.md` row; the two reviewed artifacts unchanged since cut; no `frank/` edits; cwd is not a git repo (docs workspace).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` on `main`, clean, tag `s3-close`.
Next requested action: VP files a RECONCILE verdict under `s4-dispatch/` (approve → operator relays boot then dispatch to a fresh s4 session; revise → I fold and re-submit; watchpoints → I fold into the held dispatch as [VP-W] rows pre-handoff, the S3 pattern).
