## DESIGN — step3-relock-dag-m1: §11 lane-2 (interface DAG) m-1 scope — item C env/no-leak review (the `env_digest` sanitization rule + descriptor secret-sensitivity) + item D redaction boundary (the session-content log carries no secret bytes), over frozen m-1 `7c8b09a6…`

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-dag-m1
PARENT_DISPATCH_ID: step3-relock-broker-confirm
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m1
FROM: master.orchestrator-planner
TO: m-1.planner
CC: m-1.implementer, operator, master.orchestrator-reviewer, m-9.planner, m-10.planner
SUBJECT: your lane-2 scope — the TCB reviews on the new evidence surfaces: item C env/no-leak (env_digest over the m-1-sanitized presented env + descriptor secret-sensitivity) + item D redaction boundary (no S-A/S-B bytes in the worker session-content log); governed additive delta over frozen m-1 under F73

m-1 pair — §11 lane 1 is CLOSED; the operator opened **lane 2 (the §6 interface DAG) in full**. Your lane-2 scope is the TCB's two secret-boundary reviews on the new evidence surfaces. Run your normal pair cycle. A **governed additive delta over your frozen final m-1 `7c8b09a6…`** (stays the historical lock; no in-place edit; reviewed under F73 with consumer confirmations). Ratified contracts: `master/STEP-3-STAGE6-AMENDMENT.md` **§5-C, §5-D + §6** (rev12 `1125b0a0…`).

### Your obligations
1. **Item C — the env/no-leak review (§5-C).** The bash effect descriptor carries `env_digest` = SHA-256 over the JCS-sorted COMPLETE env set ACTUALLY presented to the child — the **m-1-sanitized environment** (your hardening #1/#7, already frozen; **no behavior delta — we digest what is presented, we do not newly clear it**). Own the **sanitization rule** (what is stripped/kept before presentation) and **review secret-sensitivity**: confirm that digesting the presented env, and carrying `shell_interpreter_ref{path,version,content_id}` + `canonical_resource`/`cwd`, exposes no secret material and cannot leak a credential into an evidence field. Per §6, C is **m-10 descriptor/ticket → m-9 executor consumer; m-1 env/no-leak review** — you are the secret-boundary reviewer of the descriptor.
2. **Item D — the redaction boundary (§5-D).** The worker owns a durable session-content log (CONTENT only). Own the **redaction rule**: the log MUST NOT persist any S-A/S-B secret bytes (the §5-D "NOT in the log" row names it). Define the redaction boundary m-9 enforces at write time + how it is verified, so a crash-safe log can never become a secret-exfiltration surface. Per §6, D is the two-sided m-9⇄m-10 seam + **m-1 (redaction)** with a joint join record — you sign the redaction half.

### Join records / confirmations
- Co-sign the **§D join record** redaction half (with m-9's D1/D2 half + m-10's D2/D3 half) against the same bytes.
- You are consumed by: m-9 (C descriptor derivation must honor your sanitization rule; D log writes must honor your redaction boundary) · m-10 (C ticket schema). Return your review confirmations (or a revise-request if a surface risks a leak) on the m-9/m-10 deltas.

### Boundaries
DESIGN-only. No stage-6 DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. Frozen m-1 stays locked; the delta is governed + additive. H-12 hard-blocks external use. The secret-boundary is the TCB's sole-governed domain — a genuine leak risk is a revise-request, never a silent accept. Escalate spec mistakes UP through master.

### Where this sits
§11 lane 2 of 5. Your reviews gate the C descriptor + the D log as secret-safe before they ride the shorter re-lock. Item A (recipe + `bundle_sha256`) authored LAST over settled B–E; lane 4 = the shorter re-lock; lane 5 = T4.

## Verification
Reproduced from disk: amendment rev12 `1125b0a0…` ✓ · m-1 `7c8b09a6…` UNMOVED (governed delta requested, not an edit) ✓. Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row; no frozen design byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the m-1 pair authors + pair-reviews its lane-2 delta (C env/no-leak rule + D redaction boundary), returns the byte-bound review + the co-signed §D redaction half + F73 confirmations + a SITREP; master integrates toward the shorter re-lock.
