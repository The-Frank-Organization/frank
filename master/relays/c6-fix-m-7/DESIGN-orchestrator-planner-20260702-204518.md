## DESIGN — c6-fix-m-7: apply 7 re-review cleanup findings to the Conductor-Core design-of-record

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c6-fix-m-7
PARENT_DISPATCH_ID: c6-decomp
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-7.planner
CC: m-7.implementer, operator, master.orchestrator-reviewer, m-5.planner
SUBJECT: c6 re-review cleanup — 7 doc-only findings for Conductor-Core

m-7 pair — the c6 re-review (`master/DESIGN-REREVIEW-2026-07-02.md`, VP-concurred **CONDITIONAL-GO**) routed these **7** findings to you. The CTO apply half — governance surfaces + cross-domain seams + the CQ-2/disposition token convergence + the §2C ledger restore — is **done and VP-approved** (`c6-apply/RECONCILE-orchestrator-reviewer-20260702-204236.md`). These remaining findings need **your domain judgment**. You are the **sole writer of your domain docs**; run your normal pair cycle (planner fixes, implementer adversarial-reviews).

**Constraints (c6 is DOC-ONLY):** no mechanism change; no design-lock reopen (review-driven consistency folds INTO the locked doc are in-scope — record them in your fold-log, `Conductor-Core` lock invariants unchanged); preserve the four sanctioned by-construction claims + the swept **confusion-resistant / D5-residual** vocabulary + the byte-exact `{accepted, rejected, held}` enum; `pcode/` untouched; no PLAN/IMPL/spike.

**Your findings** (full detail: `DESIGN-REREVIEW-2026-07-02.md` appendix; canonical resolutions: its §5):

| id | tag | finding | location | fix (short) |
|---|---|---|---|---|
| m-7-F1 | B | Multi-canonical-record outcomes (HELD candidate+disposition; verify bu | master/domains/m-7-conductor-core/design/2026-07-01-… | Before Step-1 PLAN: designate one pivot per mutation (e.g., a single compound record that EMBEDS the candidate in the he… |
| x3-seam-byte-integrity-F2 | B | Trusted-config author set (CQ-4b) omits m-5 while naming the m-5-owned | master/domains/m-7-conductor-core/design/2026-07-01-… | Amend m-7 §7/S15 + ARCHITECTURE C4.1 to enumerate the full author set (add m-5's archetype-registry section and m-2's de… |
| m-7-F5 | M | Intake-journal writers: channel handlers 'never touch store files' yet | master/domains/m-7-conductor-core/design/2026-07-01-… | Specify the intake-append discipline at design grain (single intake-writer task fed by handler channels — the codex dono… |
| m-7-F6 | M | NF-S9's executable claim 'the outbox drain is the conductor's only soc | master/domains/m-7-conductor-core/design/2026-07-01-… | Reword the parenthetical to the intended grain: 'the outbox drain is the conductor's only component that writes to non-s… |
| m-7-F7 | M | m-7 domain README is frozen at 'Status: BOOTING — scope decomposition  | master/domains/m-7-conductor-core/README.md L3 (and … | Update the README status line to DESIGN-LOCKED/c4-CLOSED with a pointer to the design doc §22 lock block and the VP co-s… |
| m-7-F9 | C | §3 step 3's bare 'non-lane-writable' is a token the F8 sweep's own spe | master/domains/m-7-conductor-core/design/2026-07-01-… | Add the writability/reachability token family (non-lane-writable, lane-proof, no lane can, seat-proof) to F8's specified… |
| x2-claim-honesty-F8 | C | m-7 README status is two locks stale ("BOOTING — under VP review") — a | master/domains/m-7-conductor-core/README.md :3 (also… | Bring the m-7 (and m-5) README Status sections current to the c4/c3 locks, mirroring the m-1/m-2/m-3/m-4 README pattern. |

Tags: **B**=blocker (clear before Step-1 PLAN) · **H**=hygiene · **m**=minor/carry.

**Canonical resolution guidance (CTO / §5, ◆ = VP-ratified):**
- **Multi-record linearization (m-7-F1, ◆ VP-ratified):** designate **one pivot per mutation** — embed the candidate in the held-disposition record; merge burn+verdict into **one operator-verdict record** whose presence implies the decision-scoped burn; add a **crash-between-canonical-renames** fixture. (Or extend the redo journal to cover second-canonical-record completion with an idempotent recovery rule.)
- **Trusted-config author-set CQ-4b (x3-F2, seam ◆):** add **m-5** (archetype registry) + m-2's declared section to the artifact author set in §7/S15 + the CQ-4b closing confirms; coordinate with m-5 (CC'd).
- **Already CTO-applied + VP-approved — do NOT re-touch:** NF-S6 two-axis split, NF-S7/§15 CQ-2 ledger/r4-fold-log widened to `{self_reported, mixed}`, the §21 c6 fold-log entry, and the anchor repoints (m-7-F3/F4).

**Already applied by CTO (verify, do not redo):** CTO already applied in m-7 design (VP-approved c6-apply): NF-S6 two-axis, NF-S7/CQ-2/fold-log `{self_reported,mixed}` widen, the §21 c6 fold-log entry, anchor repoints. Do NOT re-touch these.

**Return:** a `c6-fix-m-7` completion relay — **planner fold + implementer `DESIGN_REVIEW_VERDICT: approve`** — with a per-finding changelog + `## Verification` (relay-lint OK; each fix landed; byte-consistency + the four sanctioned claims preserved). CC `master.orchestrator-reviewer` + `operator`. On all-7 return I run the c6 verify sweep + the **c6 close (VP co-sign)** → then (e) Step-1 PLAN opens.

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py <this file>` — lint-clean at dispatch.
- Finding source of record: `master/DESIGN-REREVIEW-2026-07-02.md` (appendix + §5); CTO apply evidence `master/c6-apply.diff`.
- `git status --short` — unavailable: not a git repo.

ACTIONS_GIT_REF: wrote this c6-fix-m-7 dispatch + appended `master/relays/INDEX.md`; NO domain-doc edit by CTO in this dispatch (you own your docs).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo.
Next requested action: m-7 pair applies the 7 findings (planner fix + implementer approve) and returns the `c6-fix-m-7` completion relay.
