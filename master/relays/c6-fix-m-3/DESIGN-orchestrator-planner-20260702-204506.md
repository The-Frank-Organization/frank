## DESIGN — c6-fix-m-3: apply 10 re-review cleanup findings to the Observation & Evidence design-of-record

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c6-fix-m-3
PARENT_DISPATCH_ID: c6-decomp
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, operator, master.orchestrator-reviewer, m-4.planner
SUBJECT: c6 re-review cleanup — 10 doc-only findings for Observation & Evidence

m-3 pair — the c6 re-review (`master/DESIGN-REREVIEW-2026-07-02.md`, VP-concurred **CONDITIONAL-GO**) routed these **10** findings to you. The CTO apply half — governance surfaces + cross-domain seams + the CQ-2/disposition token convergence + the §2C ledger restore — is **done and VP-approved** (`c6-apply/RECONCILE-orchestrator-reviewer-20260702-204236.md`). These remaining findings need **your domain judgment**. You are the **sole writer of your domain docs**; run your normal pair cycle (planner fixes, implementer adversarial-reviews).

**Constraints (c6 is DOC-ONLY):** no mechanism change; no design-lock reopen (review-driven consistency folds INTO the locked doc are in-scope — record them in your fold-log, `Observation & Evidence` lock invariants unchanged); preserve the four sanctioned by-construction claims + the swept **confusion-resistant / D5-residual** vocabulary + the byte-exact `{accepted, rejected, held}` enum; `pcode/` untouched; no PLAN/IMPL/spike.

**Your findings** (full detail: `DESIGN-REREVIEW-2026-07-02.md` appendix; canonical resolutions: its §5):

| id | tag | finding | location | fix (short) |
|---|---|---|---|---|
| m-3-F1 | B | conftest.py FATAL survives: run-named-suite executes lane-authored cod | master/domains/m-3-observation-evidence/design/2026-… | Before the observe-hook/E2 build step: (a) fold the sandbox SHOULD-FIX as a named build-step gate — suite-class registry… |
| m-4-F2 | B | m-3↔m-4 seam formula fork: locked m-3 doc defines deviated_observed by | master/domains/m-3-observation-evidence/design/2026-… | Fold the GL-1 canonical (bucket-vs-bucket on declared_bucket + the auxiliary bucket_binding_observed atom) into m-3 §9 v… |
| x1-fatal-resolution-F5 | B | Ranked finding #12 (observe registry executes repo-authored code insid | master/DESIGN-REVIEW-2026-07-01.md:197-198 (§3.12) a… | Record the §7 sandbox item the same way decision-④/re-mint-supersedes were recorded: a named non-locking carry in m-3 §1… |
| x3-seam-byte-integrity-F3 | B | CQ-2 certified CLOSED while the authority_class ∧ record_integrity==mi | master/domains/m-3-observation-evidence/design/2026-… | Run the flagged m-2/m-3 co-sign before Step-1 PLAN: either broaden the CQ-2 key to `record_integrity ∈ {self_reported, m… |
| m-3-F10 | M | Sweep-authored §11 sentence overstates D5 detectability: recompute cat | master/domains/m-3-observation-evidence/design/2026-… | Scope the sentence: 'an INCONSISTENT forgery (rollup ≠ per-field detail) is recompute-detectable; a coherent forged reco… |
| m-3-F4 | M | Decision-⑤ egress-fixture flag to m-7 has no receiving-side record: m- | master/domains/m-3-observation-evidence/design/2026-… | Append the ⑤ fixture pair (m-3 §13:204's (a)/(b)/(c) set) as a row on an m-7-inherited surface — the ARCHITECTURE §C4 bu… |
| m-3-F5 | M | E4/operator-observed evidence has no honest label in the two-value int | master/domains/m-3-observation-evidence/design/2026-… | At the observe-hook build step, define the operator-attested path without reopening R3's two-value per-field enum: e.g.,… |
| m-3-F6 | M | observe_result enum values (`blocked`, `degraded`, `skipped`, `unsafe` | master/domains/m-3-observation-evidence/design/2026-… | Add a total enum→{accepted, rejected, held} disposition table (keyed by authority class) to §3, or mark the two enums ex… |
| m-3-F7 | M | Base IMPL done-predicate clause 'no scope drift' is not conductor-obse | master/domains/m-3-observation-evidence/design/2026-… | Define the observable (e.g., diff ⊆ paths declared in a named form field with a named owner) or strike the clause from t… |
| x2-claim-honesty-F4 | M | m-3 §2's kept "AND unforgeable" on the probe result is a smuggled stre | master/domains/m-3-observation-evidence/design/2026-… | Relabel the one word: "…and lane-unsubstitutable through the tool surface (confusion-resistant; D5 ground-truth residual… |

Tags: **B**=blocker (clear before Step-1 PLAN) · **H**=hygiene · **m**=minor/carry.

**Canonical resolution guidance (CTO / §5, ◆ = VP-ratified):**
- **conftest.py sandbox carry (m-3-F1 / x1-F5, ◆):** record a **named non-locking carry** in §13 (owner m-3 + m-7-hosted execution) gating the **Step-2 observe build** — suite-class registry entries run in an **unprivileged executor with no store/config/outbox handle and no signing key**; relabel §4 honestly (the agent never selects arbitrary code, but suite-class entries still **execute repo-resident lane code**); flag the paired m-7 executor-isolation fixture (m-7 dispatch carries the receiving side).
- **CQ-2 mixed (x3-F3):** CTO already broadened your §6 to `{self_reported, mixed}` (m-3-F2, VP-approved) — **verify** it + ensure your §13 fixture exercises the `mixed` leg; do not re-touch the §6 disposition text.
- **deviated_observed GL-1 fold (m-4-F2, seam ◆):** fold the **GL-1 canonical** (bucket-vs-bucket on `declared_bucket` + the auxiliary `bucket_binding_observed` atom) into §9; coordinate m-4 deleting/conditioning its 'equivalent fallback' bracket (m-4 CC'd).

**Already applied by CTO (verify, do not redo):** CTO already applied in m-3 design: CQ-2 §6 broaden (m-3-F2), the `unforgeable`→scoped relabel §2 (m-3-F3), the stale header (m-3-F9), the §6/§15 mixed-edge home (x1-F3), + README `non-forgeable` (m-3-F8). Verify, don't redo.

**Return:** a `c6-fix-m-3` completion relay — **planner fold + implementer `DESIGN_REVIEW_VERDICT: approve`** — with a per-finding changelog + `## Verification` (relay-lint OK; each fix landed; byte-consistency + the four sanctioned claims preserved). CC `master.orchestrator-reviewer` + `operator`. On all-7 return I run the c6 verify sweep + the **c6 close (VP co-sign)** → then (e) Step-1 PLAN opens.

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py <this file>` — lint-clean at dispatch.
- Finding source of record: `master/DESIGN-REREVIEW-2026-07-02.md` (appendix + §5); CTO apply evidence `master/c6-apply.diff`.
- `git status --short` — unavailable: not a git repo.

ACTIONS_GIT_REF: wrote this c6-fix-m-3 dispatch + appended `master/relays/INDEX.md`; NO domain-doc edit by CTO in this dispatch (you own your docs).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo.
Next requested action: m-3 pair applies the 10 findings (planner fix + implementer approve) and returns the `c6-fix-m-3` completion relay.
