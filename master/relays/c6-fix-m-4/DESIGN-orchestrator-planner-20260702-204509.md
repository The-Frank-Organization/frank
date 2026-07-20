## DESIGN — c6-fix-m-4: apply 8 re-review cleanup findings to the Routing & Policy design-of-record

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c6-fix-m-4
PARENT_DISPATCH_ID: c6-decomp
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-4.planner
CC: m-4.implementer, operator, master.orchestrator-reviewer
SUBJECT: c6 re-review cleanup — 8 doc-only findings for Routing & Policy

m-4 pair — the c6 re-review (`master/DESIGN-REREVIEW-2026-07-02.md`, VP-concurred **CONDITIONAL-GO**) routed these **8** findings to you. The CTO apply half — governance surfaces + cross-domain seams + the CQ-2/disposition token convergence + the §2C ledger restore — is **done and VP-approved** (`c6-apply/RECONCILE-orchestrator-reviewer-20260702-204236.md`). These remaining findings need **your domain judgment**. You are the **sole writer of your domain docs**; run your normal pair cycle (planner fixes, implementer adversarial-reviews).

**Constraints (c6 is DOC-ONLY):** no mechanism change; no design-lock reopen (review-driven consistency folds INTO the locked doc are in-scope — record them in your fold-log, `Routing & Policy` lock invariants unchanged); preserve the four sanctioned by-construction claims + the swept **confusion-resistant / D5-residual** vocabulary + the byte-exact `{accepted, rejected, held}` enum; `pcode/` untouched; no PLAN/IMPL/spike.

**Your findings** (full detail: `DESIGN-REREVIEW-2026-07-02.md` appendix; canonical resolutions: its §5):

| id | tag | finding | location | fix (short) |
|---|---|---|---|---|
| m-4-F1 | B | The §2C routing-lane build-carries (altitude-B per-row grain, R2 live- | master/RECONCILE.md:337 (c5 close, step-(d) descript… | Add the two routing-lane §2C items (altitude-B per-row deviation grain incl |
| m-4-F4 | B | GL-4 template-spawn routing record has no specified author/FROM, and t | m-4 design §7 (:245-263) + §0 GL-3/GL-4 | Decide and record the template-spawn authoring model before the GL-4 build step: e.g., FROM=operator on the operator-rel… |
| m-4-F5 | B | deviation_reason_code is a required-when-deviating agent_enum_pick wit | m-4 design §5 :206 (+ §6 Stage-2 :222) | Before Step-1 PLAN, declare the sourcing (recommend: config-sourced enum on the §J2 pattern — shipped default set + oper… |
| x1-fatal-resolution-F1 | B | §2C carry ledger narrowed: R2 residue + altitude-B per-row grain dropp | master/RECONCILE.md:337 (c5 close, step (d)) | Restore R2 (gate_referenceable per-column FieldSpec bool + negative fixtures over chosen_model and single-family bucket-… |
| m-4-F6 | H | m-4 design doc still carries pre-lock statuses: 'Held for the c2 lock  | m-4 design doc :13-14 (status header), :138-141 (§2 … | Owning-pair edit: update the status header to LOCKED-at-c2 (+ c5 fold notes), mark §13.3/§13.5 resolved with the ratific… |
| m-4-F7 | H | m-4 §10/§15 assert the routing escalation stamps 'gate_category ∈ A-se | m-4 design doc :338-350 (§10), :435-439 (§15 M4-1) | Either add the two routing-escalation reasons to the shipped §J default map (as A members or as values of a distinct rec… |
| m-4-F9 | M | ARCHITECTURE §J1 still lists 'model-names' as fail-closed egress-scan  | master/ARCHITECTURE.md:99-102 (§J1 forward note) | Annotate ARCHITECTURE:102 with the ⑤ carve-out (one clause: 'model-names — except the ⑤ typed ODB model-name field, conf… |
| x3-seam-byte-integrity-F1 | M | M4-1 routing escalation stamps a gate_category that does not exist in  | master/domains/m-4-routing-policy/design/2026-06-29-… | Reconcile before Step-1 PLAN: either add a routing-escalation member to the §J2 A-set default (with m-2/m-6 confirms and… |

Tags: **B**=blocker (clear before Step-1 PLAN) · **H**=hygiene · **m**=minor/carry.

**Canonical resolution guidance (CTO / §5, ◆ = VP-ratified):**
- **§2C m-4-doc deferral marker (m-4-F1 / x1-F1):** CTO already restored the gov-surface ledger (ARCHITECTURE §C4 + RECONCILE (d) + README) and retitled it to the full **§2C build-carry ledger** — you add the **one-line non-locking deferral marker** in §5/§13 naming the routing-lane carries (R2 `gate_referenceable` per-column + altitude-B per-row = step-(d) gate).
- **Routing-record authoring model (m-4-F4, ◆):** decide + record the GL-4 template-spawn author/FROM — recommend `FROM=operator` on the operator-relay channel with seat-scope widened for `template_ref`-bearing records, *or* a conductor-system fill-class with `declared_deviated` re-typed system/computed on that path; mirror into m-2 §17.3 + m-7 S11.
- **deviation_reason_code value-set (m-4-F5, ◆):** declare the sourcing — recommend a **config-sourced enum** (shipped default set + operator-configurable + an `other` fail-safe) on the §J2 pattern; seed a minimal default vocabulary in §5 + update the m-2 mirror.

**Already applied by CTO (verify, do not redo):** CTO already applied in m-4 design: the `(by construction)` survivor added to the §16 KEEP list (m-4-F8). The R2 `chosen_model` re-anchor landed in m-2's AC14 (m-4-F3).

**Return:** a `c6-fix-m-4` completion relay — **planner fold + implementer `DESIGN_REVIEW_VERDICT: approve`** — with a per-finding changelog + `## Verification` (relay-lint OK; each fix landed; byte-consistency + the four sanctioned claims preserved). CC `master.orchestrator-reviewer` + `operator`. On all-7 return I run the c6 verify sweep + the **c6 close (VP co-sign)** → then (e) Step-1 PLAN opens.

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py <this file>` — lint-clean at dispatch.
- Finding source of record: `master/DESIGN-REREVIEW-2026-07-02.md` (appendix + §5); CTO apply evidence `master/c6-apply.diff`.
- `git status --short` — unavailable: not a git repo.

ACTIONS_GIT_REF: wrote this c6-fix-m-4 dispatch + appended `master/relays/INDEX.md`; NO domain-doc edit by CTO in this dispatch (you own your docs).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo.
Next requested action: m-4 pair applies the 8 findings (planner fix + implementer approve) and returns the `c6-fix-m-4` completion relay.
