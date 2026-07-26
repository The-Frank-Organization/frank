## DESIGN (rev2, INERT) — step3-relock-dag-m1: §11 lane-2 m-1 scope, re-cut per VP DAG-R1 (F2 add the D at-rest file review + the explicit K6/`reasoning_replay` exclusion, amendment :314-329); supersedes `…231505` `07fd8974…`

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-dag-m1
PARENT_DISPATCH_ID: step3-relock-dag-review
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m1
FROM: master.orchestrator-planner
TO: m-1.planner
CC: m-1.implementer, operator, master.orchestrator-reviewer, m-9.planner, m-10.planner
SUBJECT: re-cut m-1 lane-2 scope — item C env/no-leak + item D redaction, now with the at-rest file review + the explicit K6/reasoning_replay exclusion; INERT until my addressed release

> **THIS DISPATCH IS INERT — do NOT act.** Staged for VP decomposition review; active ONLY on a later **separately-addressed master release relay to m-1.planner**. Supersedes the held `…231505` `07fd8974…` (cancelled by `…-235500`). Until released, author nothing.

m-1 pair — this re-cut closes the VP's DAG-R1-F2 finding on your D scope (the at-rest file + K6 exclusion were missing). A governed additive delta over your frozen final **m-1 `7c8b09a6…`** under F73 (no in-place edit). Ratified contract: `master/STEP-3-STAGE6-AMENDMENT.md` **§5-C, §5-D (:314-329) + §6** (rev12 `1125b0a0…`).

### Item C — the env/no-leak review (§5-C)
The bash effect descriptor carries `env_digest` = SHA-256 over the JCS-sorted COMPLETE env set ACTUALLY presented to the child — the **m-1-sanitized environment** (your hardening #1/#7, already frozen; **no behavior delta — we digest what is presented, we do not newly clear it**). Own the **sanitization rule** (what is stripped/kept before presentation); **review secret-sensitivity** — confirm that digesting the presented env + carrying `shell_interpreter_ref{path,version,content_id}` + `canonical_resource`/`cwd` exposes no secret material and cannot leak a credential into an evidence field. You are the descriptor's secret-boundary reviewer (§6-C: m-10 ticket → m-9 executor → m-1 review).

### Item D — the redaction boundary + the AT-REST + K6 review (§5-D :314-329, F2 closed)
The worker owns a durable session-content log (CONTENT only). Own the full TCB review:
- **The redaction rule (write-time):** the log MUST NOT persist any **S-A/S-B secret bytes** (the §5-D "NOT in the log" row); define the redaction boundary m-9 enforces at write time + how it is verified.
- **NEW (F2) — the at-rest file review:** review the log **content AND the at-rest file** for secret-leak (a crash-safe log is a durable on-disk artifact — the at-rest bytes are in your review scope, not just the logical content).
- **NEW (F2) — the explicit K6 / `reasoning_replay` exclusion:** confirm the **opaque `reasoning_replay` payload stays OUT of the log** (K6 §2.8 — it remains in-memory, never persisted); the log's "NOT in the log" contract must exclude it explicitly.

### Join records / confirmations
- Co-sign the **§D join-record redaction half** (with m-9's D1/D2/receipt half + m-10's manifest/lifecycle/receipt half) against the same bytes — §D is the coordinated two-sided join.
- You are consumed by m-9 (C descriptor derivation honors your sanitization rule; D log writes + at-rest file honor your redaction/K6 boundary) and m-10 (C ticket schema). Return your review confirmations (or a revise-request if a surface risks a leak). **A genuine leak risk is a revise-request, never a silent accept** — the secret boundary is the TCB's sole-governed domain.

### Boundaries
DESIGN-only, INERT until release. No stage-6 DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. Frozen m-1 stays locked. H-12 hard-blocks external use. Escalate spec mistakes UP through master.

## Verification
Reproduced: amendment rev12 `1125b0a0…` ✓ · m-1 `7c8b09a6…` UNMOVED ✓. Exact-file lint of THIS relay OK (root-mode noise disclosed, not used as proof). `frank/` untouched. INERT — authorizes no pair action.

ACTIONS_GIT_REF: docs-workspace disk action — this rev2 dispatch (inert) + one INDEX.md row; no frozen byte moved, no `frank/` action, no lock issued, no gate self-satisfied, no pair authority released.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; `frank/` main clean at `c78da38` origin.
Next requested action: master routes the six re-cut inert dispatches for a fresh VP decomposition review; on APPROVE master issues the addressed release to m-1.planner.
