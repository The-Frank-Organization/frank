## DESIGN — c6-fix-m-5: apply 3 re-review cleanup findings to the Workflows & Archetypes design-of-record

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c6-fix-m-5
PARENT_DISPATCH_ID: c6-decomp
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-5.planner
CC: m-5.implementer, operator, master.orchestrator-reviewer, m-7.planner
SUBJECT: c6 re-review cleanup — 3 doc-only findings for Workflows & Archetypes

m-5 pair — the c6 re-review (`master/DESIGN-REREVIEW-2026-07-02.md`, VP-concurred **CONDITIONAL-GO**) routed these **3** findings to you. The CTO apply half — governance surfaces + cross-domain seams + the CQ-2/disposition token convergence + the §2C ledger restore — is **done and VP-approved** (`c6-apply/RECONCILE-orchestrator-reviewer-20260702-204236.md`). These remaining findings need **your domain judgment**. You are the **sole writer of your domain docs**; run your normal pair cycle (planner fixes, implementer adversarial-reviews).

**Constraints (c6 is DOC-ONLY):** no mechanism change; no design-lock reopen (review-driven consistency folds INTO the locked doc are in-scope — record them in your fold-log, `Workflows & Archetypes` lock invariants unchanged); preserve the four sanctioned by-construction claims + the swept **confusion-resistant / D5-residual** vocabulary + the byte-exact `{accepted, rejected, held}` enum; `pcode/` untouched; no PLAN/IMPL/spike.

**Your findings** (full detail: `DESIGN-REREVIEW-2026-07-02.md` appendix; canonical resolutions: its §5):

| id | tag | finding | location | fix (short) |
|---|---|---|---|---|
| m-7-F2 | B | The trusted-config artifact's contents include m-5's archetype registr | master/domains/m-7-conductor-core/design/2026-07-01-… | A bounded COORD before Step-1 PLAN: add m-5 to the artifact's author set (an m-5 section carrying the registry, with m-5… |
| m-5-F6 | M | T1 declares 'observe-as-send always on' and §9's Step-1-rideable frami | master/domains/m-5-workflows-archetypes/design/2026-… | One-line phase annotation in §7/§9 mirroring the ceiling tiering: 'the slot_in observe-invariant families bite from Step… |
| m-5-F7 | C | Per-axis absent-defaults are mandated but declared for only the deferr | master/domains/m-5-workflows-archetypes/design/2026-… | Fold one sentence into §5 at the next bounded m-5 touch: 'every axis MUST declare an absent-default; a record-absence de… |

Tags: **B**=blocker (clear before Step-1 PLAN) · **H**=hygiene · **m**=minor/carry.

**Canonical resolution guidance (CTO / §5, ◆ = VP-ratified):**
- **Trusted-config author-set (m-7-F2, seam ◆):** add m-5's **archetype-registry section** to the trusted-config artifact author set and provide the **m-5 CQ-4b confirm**; coordinate with m-7 (CC'd) who holds §7/S15/CQ-4b.
- **Posture-vocab away-trigger (m-5-F2):** re-record the runtime-away/posture spec so the **away-mode trigger is expressible** in the locked posture model (it currently isn't); keep it within the locked lattice.

**Already applied by CTO (verify, do not redo):** CTO already applied in m-5 design: stale status header (m-5-F4), §6 table token normalization (m-5-F5), the `conductor sole-sender` relabel §12 (x2-F5/x3-F6). Verify, don't redo.

**Return:** a `c6-fix-m-5` completion relay — **planner fold + implementer `DESIGN_REVIEW_VERDICT: approve`** — with a per-finding changelog + `## Verification` (relay-lint OK; each fix landed; byte-consistency + the four sanctioned claims preserved). CC `master.orchestrator-reviewer` + `operator`. On all-7 return I run the c6 verify sweep + the **c6 close (VP co-sign)** → then (e) Step-1 PLAN opens.

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py <this file>` — lint-clean at dispatch.
- Finding source of record: `master/DESIGN-REREVIEW-2026-07-02.md` (appendix + §5); CTO apply evidence `master/c6-apply.diff`.
- `git status --short` — unavailable: not a git repo.

ACTIONS_GIT_REF: wrote this c6-fix-m-5 dispatch + appended `master/relays/INDEX.md`; NO domain-doc edit by CTO in this dispatch (you own your docs).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo.
Next requested action: m-5 pair applies the 3 findings (planner fix + implementer approve) and returns the `c6-fix-m-5` completion relay.
