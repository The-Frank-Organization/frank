## DESIGN — c6-fix-m-6: apply 5 re-review cleanup findings to the Human Surface & Scheduler design-of-record

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c6-fix-m-6
PARENT_DISPATCH_ID: c6-decomp
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-6.planner
CC: m-6.implementer, operator, master.orchestrator-reviewer
SUBJECT: c6 re-review cleanup — 5 doc-only findings for Human Surface & Scheduler

m-6 pair — the c6 re-review (`master/DESIGN-REREVIEW-2026-07-02.md`, VP-concurred **CONDITIONAL-GO**) routed these **5** findings to you. The CTO apply half — governance surfaces + cross-domain seams + the CQ-2/disposition token convergence + the §2C ledger restore — is **done and VP-approved** (`c6-apply/RECONCILE-orchestrator-reviewer-20260702-204236.md`). These remaining findings need **your domain judgment**. You are the **sole writer of your domain docs**; run your normal pair cycle (planner fixes, implementer adversarial-reviews).

**Constraints (c6 is DOC-ONLY):** no mechanism change; no design-lock reopen (review-driven consistency folds INTO the locked doc are in-scope — record them in your fold-log, `Human Surface & Scheduler` lock invariants unchanged); preserve the four sanctioned by-construction claims + the swept **confusion-resistant / D5-residual** vocabulary + the byte-exact `{accepted, rejected, held}` enum; `pcode/` untouched; no PLAN/IMPL/spike.

**Your findings** (full detail: `DESIGN-REREVIEW-2026-07-02.md` appendix; canonical resolutions: its §5):

| id | tag | finding | location | fix (short) |
|---|---|---|---|---|
| m-6-F3 | M | Decision-⑤'s claimed 'm-7 egress-fixture note' does not exist in m-7;  | master/domains/m-6-human-surface-scheduler/design/20… | Record the ⑤ egress-fixture update as an owed m-7 fixture (amend NF-S9 or add a row) in the same ledger that carries dec… |
| m-6-F6 | M | Meeting-lane routing keys on GRILL_REQUIRED as a 'LOCKED field', but m | master/domains/m-6-human-surface-scheduler/design/20… | Have m-2 declare GRILL_REQUIRED's v3 registry row (owner/type/values — presumably the ported v2.8.8 header, agent-raised… |
| m-6-F7 | M | Stale lock-status: m-6 STATUS line and README still say 'Awaiting only | master/domains/m-6-human-surface-scheduler/design/20… | Update the STATUS tail and README top bullet to LOCKED with the co-sign id (c3-lock 20260630-191315) plus the c4/c5 fold… |
| m-6-F8 | M | Re-baseline edits shifted ARCHITECTURE.md by ~13-15 lines, silently in | master/domains/m-6-human-surface-scheduler/design/20… | Re-base the m-6 doc's ARCHITECTURE line-anchors (or convert to stable section anchors: §J2-A-set, §J1, §5-R1), and note … |
| m-6-F9 | C | Count/provenance drift in the Seam-C carry list: 'the four additive la | master/domains/m-6-human-surface-scheduler/design/20… | Split the §12 list ('four from COORD-182600' + 'one recorded c5, decision ④, step-(d) gated') or update the counts; triv… |

Tags: **B**=blocker (clear before Step-1 PLAN) · **H**=hygiene · **m**=minor/carry.

**Canonical resolution guidance (CTO / §5, ◆ = VP-ratified):**
- **Decision-⑤ m-7 note (m-6-F3):** the claimed 'm-7 egress-fixture note' does not exist in m-7 — reword to **'flagged to m-7, fixture registration owed'** (mirroring the m-6-F1 fix CTO already applied at §2:44) *or* coordinate m-7 registering it.
- **GRILL_REQUIRED field status (m-6-F6):** reconcile the meeting-lane routing that keys on `GRILL_REQUIRED` as a 'LOCKED field' against m-2's v3 schema status of that field — align to m-2's actual FieldSpec.

**Already applied by CTO (verify, do not redo):** CTO already applied in m-6 design: the §2:44 decision-③ reword to 'fixture registration owed' (m-6-F1) and the `held` gloss narrowing §2/§4 (m-6-F2). Verify, don't redo.

**Return:** a `c6-fix-m-6` completion relay — **planner fold + implementer `DESIGN_REVIEW_VERDICT: approve`** — with a per-finding changelog + `## Verification` (relay-lint OK; each fix landed; byte-consistency + the four sanctioned claims preserved). CC `master.orchestrator-reviewer` + `operator`. On all-7 return I run the c6 verify sweep + the **c6 close (VP co-sign)** → then (e) Step-1 PLAN opens.

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py <this file>` — lint-clean at dispatch.
- Finding source of record: `master/DESIGN-REREVIEW-2026-07-02.md` (appendix + §5); CTO apply evidence `master/c6-apply.diff`.
- `git status --short` — unavailable: not a git repo.

ACTIONS_GIT_REF: wrote this c6-fix-m-6 dispatch + appended `master/relays/INDEX.md`; NO domain-doc edit by CTO in this dispatch (you own your docs).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo.
Next requested action: m-6 pair applies the 5 findings (planner fix + implementer approve) and returns the `c6-fix-m-6` completion relay.
