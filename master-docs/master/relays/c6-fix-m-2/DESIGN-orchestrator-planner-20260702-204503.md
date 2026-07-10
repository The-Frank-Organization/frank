## DESIGN — c6-fix-m-2: apply 7 re-review cleanup findings to the Forms & Determinism design-of-record

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c6-fix-m-2
PARENT_DISPATCH_ID: c6-decomp
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-2.planner
CC: m-2.implementer, operator, master.orchestrator-reviewer, m-4.planner, m-5.planner
SUBJECT: c6 re-review cleanup — 7 doc-only findings for Forms & Determinism

m-2 pair — the c6 re-review (`master/DESIGN-REREVIEW-2026-07-02.md`, VP-concurred **CONDITIONAL-GO**) routed these **7** findings to you. The CTO apply half — governance surfaces + cross-domain seams + the CQ-2/disposition token convergence + the §2C ledger restore — is **done and VP-approved** (`c6-apply/RECONCILE-orchestrator-reviewer-20260702-204236.md`). These remaining findings need **your domain judgment**. You are the **sole writer of your domain docs**; run your normal pair cycle (planner fixes, implementer adversarial-reviews).

**Constraints (c6 is DOC-ONLY):** no mechanism change; no design-lock reopen (review-driven consistency folds INTO the locked doc are in-scope — record them in your fold-log, `Forms & Determinism` lock invariants unchanged); preserve the four sanctioned by-construction claims + the swept **confusion-resistant / D5-residual** vocabulary + the byte-exact `{accepted, rejected, held}` enum; `pcode/` untouched; no PLAN/IMPL/spike.

**Your findings** (full detail: `DESIGN-REREVIEW-2026-07-02.md` appendix; canonical resolutions: its §5):

| id | tag | finding | location | fix (short) |
|---|---|---|---|---|
| m-2-F1 | B | Ratified decision ③ (RAISE-ONLY gate_category + known-A detector) has  | master/domains/m-2-forms-determinism/design/2026-06-… | Bounded m-2 fold on the CQ-3 pattern: either retype the category pick as a monotonic-toward-A hybrid (system floor = kno… |
| m-5-F1 | B | seat_archetype/authority_ceiling confused-deputy survives: F1 says con | master/domains/m-4-routing-policy/design/2026-06-29-… | Bounded m-4+m-5 fold (the register's own prescription) before the routing record ships at Step-1: (a) split propose-vs-s… |
| m-2-F10 | M | human_gate_reason vs gate_category — two field ids for one semantic, w | master/domains/m-2-forms-determinism/design/2026-06-… | Declare the canonical record-level id (one field, one id — e.g., rename the header field to `gate_category`, or state th… |
| m-2-F5 | M | The gate-referenceable-field allowlist is load-bearing for two KEPT by | master/domains/m-2-forms-determinism/design/2026-06-… | Make gate-referenceability first-class FieldSpec data (e.g., per-field `gate_referenceable: bool`, default false, with t… |
| m-2-F7 | M | m-2 domain README status rows contradict the locked record (joint lock | master/domains/m-2-forms-determinism/README.md rows … | Sweep the m-2 README status table to the post-c4/c5 state (lock CLOSED with pointer; row :31 legs closed with the F2 mix… |
| m-2-F9 | C | Stale internal markers: §15 operator-judgment items still 'proposed' t | master/domains/m-2-forms-determinism/design/2026-06-… | Mark §15 items (i)/(ii) RATIFIED → ARCH §J1/§J2 (same convention as Q-A..Q-E); repoint the cite to a stable section anch… |
| m-5-F2 | M | The away-mode trigger is unexpressable in the locked posture model — t | master/domains/m-5-workflows-archetypes/design/2026-… | Re-record the runtime-away / posture-vocab split as a named non-locking carry in the surviving §C4 integrated ledger (ne… |

Tags: **B**=blocker (clear before Step-1 PLAN) · **H**=hygiene · **m**=minor/carry.

**Canonical resolution guidance (CTO / §5, ◆ = VP-ratified):**
- **Decision-③ RAISE-ONLY typing (m-2-F1, ◆ VP-ratified):** fold the pick as **monotonic-toward-A** — system floor = known-A membership, pick constrained within `[floor, A]` — *or* keep `agent_enum_pick` and declare the known-A detector as a named MAX-raiser with a declared home for the forced/effective result; add the ③ negative fixture (B-pick over known-A ⇒ raised, recorded).
- **Archetype/authority_ceiling propose-vs-stamp (m-5-F1, seam ◆):** split propose-vs-stamp — the planner *proposes* role/archetype from a fill-time-pruned candidate set; the conductor *stamps* the resolved `seat_archetype` + `authority_ceiling` **per-column** (system ownership inside the row, alongside the §2C per-column FieldSpec work). Coordinate the m-4 (routing record) + m-5 (archetype registry) sides — CC'd.

**Already applied by CTO (verify, do not redo):** CTO already applied in m-2 design: CQ-2 §17.6 broaden (m-2-F2), `chosen_model` re-anchor (m-4-F3), `template_ref` fix (m-5-F3), `delivery_state` close (m-2-F4), the exempt ODB `model_name` slot (m-2-F6), the §8 sole-path scoping (m-2-F8). Verify, don't redo.

**Return:** a `c6-fix-m-2` completion relay — **planner fold + implementer `DESIGN_REVIEW_VERDICT: approve`** — with a per-finding changelog + `## Verification` (relay-lint OK; each fix landed; byte-consistency + the four sanctioned claims preserved). CC `master.orchestrator-reviewer` + `operator`. On all-7 return I run the c6 verify sweep + the **c6 close (VP co-sign)** → then (e) Step-1 PLAN opens.

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py <this file>` — lint-clean at dispatch.
- Finding source of record: `master/DESIGN-REREVIEW-2026-07-02.md` (appendix + §5); CTO apply evidence `master/c6-apply.diff`.
- `git status --short` — unavailable: not a git repo.

ACTIONS_GIT_REF: wrote this c6-fix-m-2 dispatch + appended `master/relays/INDEX.md`; NO domain-doc edit by CTO in this dispatch (you own your docs).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo.
Next requested action: m-2 pair applies the 7 findings (planner fix + implementer approve) and returns the `c6-fix-m-2` completion relay.
