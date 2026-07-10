## COORD — c5 decision-fold ⑤ ODB model-name egress carve-out: fold into m-3 + m-6 + m-4 (owner-authored, full-pair)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c5-fold-decision-5
PARENT_DISPATCH_ID: c5-decomp
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-3.planner, m-6.planner, m-4.planner
CC: master.orchestrator-reviewer, operator, m-3.implementer, m-6.implementer, m-4.implementer, m-7.planner
BUNDLE_ID: c5-fold-decision-5
OWNER: m-3 (egress scan) + m-6 (ODB) + m-4 (model field / R2)

Fold **operator decision ⑤** (`READINESS-REGISTER.md` §Operator-decisions ⑤, RECORDED 2026-06-30) into your locked docs. A recorded decision, folded — full-pair, joint across the three owners.

**Decision ⑤ — narrow ODB model-name egress carve-out (verbatim intent):** exempt **only** the model-name field inside a conductor-generated **operator-facing ODB** from the **confidentiality** egress scan. **R2 (model ≠ gate input) is UNTOUCHED** — peer-bias protection intact. The general egress rule still blocks model-names on all other external sends; in away-mode the transport stays gated by the away-bridge opt-in; outside away-mode the ODB renders locally and never leaves.

**Fold (owner-authored, joint):**
- **m-3 (egress):** the scan **carve-out** — a typed exemption scoped to `record_kind = ODB` (operator-facing) + field = model-name + destination = operator, exempt from the **confidentiality** class only (NOT the safety/content class). Everything else still blocks model-names. Fail-closed default preserved for all other paths.
- **m-6 (ODB):** the ODB render carries the model-name in a **typed, exempt-marked** field so the scan knows the carve-out applies only here; the away-bridge opt-in gates the transport (away-mode); local render never egresses (non-away).
- **m-4 (R2 guard):** confirm the carve-out is **confidentiality-scoped only** and does **not** touch R2 — the model-name remains bookkeeping/payload, never a gate input; the carve-out is about *egress confidentiality*, not *gate-referenceability*. (This is the load-bearing separation: ⑤ relaxes a confidentiality scan on one operator-facing field; it does not make the model a gate input.)

**Requirements (VP-set c5 shape):**
1. Each owner authors its half into its locked doc; the carve-out closes only with **all three co-confirms** (m-3 scan + m-6 ODB + m-4 R2-guard).
2. Each owner **addresses its implementer** for a review-only semantic `DESIGN-REVIEW` approve (CC ≠ authority).
3. Map decision ⑤ → **folded** (all three halves) in the closure artifact. Flag to m-7 if it needs an egress-fixture update (ODB model-name field passes the confidentiality carve-out; all other model-name egress still blocked).

Not authorized / not claimed: fold a recorded decision (claim/policy text); no R2 change, no locked-contract mechanism reopen, no PLAN/IMPL/`pcode`/spike, no operator decision reopened.

## Verification

- `python3 ~/.claude/skills/tools/relay-lint.py master/relays/c5-fold-decision-5/DESIGN-orchestrator-planner-20260702-132752.md` — OK
- `python3 ~/.claude/skills/tools/relay-lint.py --relay-root master/relays/c5-fold-decision-5` — OK
- `git -C pcode status --short` — clean
- `git status --short` — unavailable: `fatal: not a git repository (or any of the parent directories): .git`

ACTIONS_GIT_REF: wrote this decision-⑤ fold dispatch + appended `master/relays/INDEX.md`; no doc edit (owners author), no R2 change, no code/`pcode/`, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-3/m-6/m-4 author their halves + address implementers for review; I fold the three co-confirms into the c5 ledger.
