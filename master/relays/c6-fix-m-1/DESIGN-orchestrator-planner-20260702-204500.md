## DESIGN — c6-fix-m-1: apply 5 re-review cleanup findings to the Trust & Identity design-of-record

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: c6-fix-m-1
PARENT_DISPATCH_ID: c6-decomp
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-1.planner
CC: m-1.implementer, operator, master.orchestrator-reviewer
SUBJECT: c6 re-review cleanup — 5 doc-only findings for Trust & Identity

m-1 pair — the c6 re-review (`master/DESIGN-REREVIEW-2026-07-02.md`, VP-concurred **CONDITIONAL-GO**) routed these **5** findings to you. The CTO apply half — governance surfaces + cross-domain seams + the CQ-2/disposition token convergence + the §2C ledger restore — is **done and VP-approved** (`c6-apply/RECONCILE-orchestrator-reviewer-20260702-204236.md`). These remaining findings need **your domain judgment**. You are the **sole writer of your domain docs**; run your normal pair cycle (planner fixes, implementer adversarial-reviews).

**Constraints (c6 is DOC-ONLY):** no mechanism change; no design-lock reopen (review-driven consistency folds INTO the locked doc are in-scope — record them in your fold-log, `Trust & Identity` lock invariants unchanged); preserve the four sanctioned by-construction claims + the swept **confusion-resistant / D5-residual** vocabulary + the byte-exact `{accepted, rejected, held}` enum; `pcode/` untouched; no PLAN/IMPL/spike.

**Your findings** (full detail: `DESIGN-REREVIEW-2026-07-02.md` appendix; canonical resolutions: its §5):

| id | tag | finding | location | fix (short) |
|---|---|---|---|---|
| m-1-F3 | B | DI-1 is never dispositioned for Step-1: still listed as a must-realize | master/domains/m-1-trust-identity/design/2026-06-28-… | Before Step-1 PLAN: add a DI-1 disposition line to §4 (DI-1 = D3-shelved wrap invariant; Step-1 I1 rides the m-7 interfa… |
| m-1-F12 | C | m-1's submit() enumerates only two persisted outcomes (accepted / term | master/domains/m-1-trust-identity/design/2026-06-28-… | One sentence in §5: on an engine-internal fault during commit, disposition follows m-7 C4.1 (authority record → held, pe… |
| m-1-F4 | M | DI-2's PLAN/E2 test is a wrap-grade isolation probe ('sibling must fai | master/domains/m-1-trust-identity/design/2026-06-28-… | Re-cut §12 #3 / §13 #1 / §13.rev2 #1 to the testable D4 grain (tool-surface negative fixtures: no verb/path/credential i… |
| m-1-F5 | M | Transport-selection rationale voided by the sweep: Option A now 'licen | master/domains/m-1-trust-identity/design/2026-06-28-… | One sentence in §3/§10: A ≻ B because the minted-credential mechanism is what the D3 wrap upgrade and the SO_PEERCRED/mT… |
| m-1-F6 | M | DESIGN-REVIEW §2A.5's ordered strike of the stale m-1 §13.2/§7 'm-3 ob | master/domains/m-1-trust-identity/design/2026-06-28-… | Execute the ordered strike: annotate §13.2's 'm-3 observe' as '[Step-2+, reserved hook 2c]' and add the 2c/CQ-1(a) cross… |

Tags: **B**=blocker (clear before Step-1 PLAN) · **H**=hygiene · **m**=minor/carry.

**Canonical resolution guidance (CTO / §5, ◆ = VP-ratified):**
- **DI-1 disposition (m-1-F3, ◆):** DI-1 = **D3-shelved wrap invariant**; Step-1 I1 **rides the m-7 interface guardrail** (store path absent from every seat tool surface). Re-cut §12 AC#1 to the D4 guardrail fixture (no seat tool/verb reaches the store path) **plus** an explicitly D3-scoped wrap-spike criterion; align the §4 DI list + §9 Step-1 row to say Step-1 I1 = guardrail-borne, not a lane-excluding store ACL. (ARCHITECTURE:39-40 already CTO-relabeled to match.)

**Already applied by CTO (verify, do not redo):** CTO already applied in your docs: §4 invariant counts (m-1-F7), the `unspoofable` relabel at :127 (m-1-F9), the two stale ARCHITECTURE seam-anchors (m-1-F10), and the README §Owns relabel (m-1-F1/x2-F2). Verify, don't redo.

**Return:** a `c6-fix-m-1` completion relay — **planner fold + implementer `DESIGN_REVIEW_VERDICT: approve`** — with a per-finding changelog + `## Verification` (relay-lint OK; each fix landed; byte-consistency + the four sanctioned claims preserved). CC `master.orchestrator-reviewer` + `operator`. On all-7 return I run the c6 verify sweep + the **c6 close (VP co-sign)** → then (e) Step-1 PLAN opens.

## Verification

- `python3 /Users/jack/.claude/skills/tools/relay-lint.py <this file>` — lint-clean at dispatch.
- Finding source of record: `master/DESIGN-REREVIEW-2026-07-02.md` (appendix + §5); CTO apply evidence `master/c6-apply.diff`.
- `git status --short` — unavailable: not a git repo.

ACTIONS_GIT_REF: wrote this c6-fix-m-1 dispatch + appended `master/relays/INDEX.md`; NO domain-doc edit by CTO in this dispatch (you own your docs).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo.
Next requested action: m-1 pair applies the 5 findings (planner fix + implementer approve) and returns the `c6-fix-m-1` completion relay.
