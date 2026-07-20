## DESIGN-REVIEW - s8 executor host r2 must revise residual F4 wording contradictions

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-design-m7-executor-review-r2
PARENT_DISPATCH_ID: s8-design-m7-executor
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - technical approval remains separate from the operator-owned OS-sandbox election
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m7-executor
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s8-executor-grill-r1
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m7-executor/DESIGN-planner-20260711-032229.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner
SUBJECT: r2 must revise - a removed shared cache still survives in cleanup text, and replay/network wording still exceeds the admitted same-uid residual; the five r1 mechanisms otherwise fold cleanly

DESIGN_REVIEW_VERDICT: must-revise

The r2 fold resolves the five original findings mechanically: suite-only host admission preserves m-3's taxonomy; shared cache is rejected; execution is at-least-once with one committed verdict; group death is checked before cleanup; and only the closed m-3 `CheckVerdict` crosses the boundary. Three residual wording contradictions still violate the document's own F4 rule.

## Findings

### F1 - Cleanup still says the removed shared cache survives runs

Section 2.2 says all cache variables point into the per-run workdir and no shared per-executor cache exists (`2026-07-11-s8-executor-host.md:20`). Section 2.9 still says "Cache dir survives runs" (`design:27`). Those cannot both be true, and the latter reintroduces the cross-run surface r2 rejected.

Required fold: remove the surviving-cache sentence. State that all v1 cache content is run-scoped under the workdir and follows the full §2.7 cleanup state machine; on a survivor fault the whole workdir, including cache content, is preserved for bounded diagnosis.

### F2 - The at-least-once execution rationale calls ambiently capable code unconditionally safe

Section 2.11 and the GRILL_LOCK say re-execution is "safe because v1 accepts only read-only suite-class runs" (`design:31,61`). The same document correctly admits suite code can use ambient same-uid authority to write canonical state (`design:10-15`). Read-only is the sanctioned check contract and provided surface, not OS enforcement; therefore repeated execution is not unconditionally safe under the accepted residual.

Required fold: narrow the claim to: at-least-once execution introduces no **sanctioned/provided-surface** side effect for conforming suite checks; ambient same-uid misbehavior remains the explicitly accepted operator risk until sandboxing. Preserve one committed verdict/one pivot as the actual conductor guarantee.

### F3 - Network absence is still phrased as an assertion without enforcement

Section 2.6 says network is "asserted absent at the CLAIM level only" while immediately admitting v1 has no OS-level denial (`design:24`). The GRILL_LOCK correctly rejects enforcement-level absence (`design:74`). The current phrase can still be read as an observed absence claim and is weaker than the document's otherwise precise F4 wording.

Required fold: say network use is **unsanctioned by the suite-class declaration/policy**, not absent. The executor provides no network handle or credential, but ambient network access remains possible and untested without the operator-elected sandbox.

## Confirmed Fold Evidence

- R1 F1: the host accepts only m-3 suite-class entries and refuses side-effecting entries as defense-in-depth without taking policy ownership.
- R1 F2: manifest identity is pre-spawn-byte identity; staged copies remain mutable; the shared cache is rejected normatively.
- R1 F3: FX-EXE-3 now proves one committed verdict/one pivot while admitting crash-driven re-execution.
- R1 F4: cleanup waits for direct-child reap plus bounded process-group ESRCH confirmation; survivor expiry routes to machinery fault with workdir preservation.
- R1 F5: `RunResult` stays internal and the returned byte shape is m-3's closed, redacted `CheckVerdict`.
- Operator election: approval remains technical only; the OS-sandbox risk posture still requires its operator record before lock effectiveness.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, no lock, and no sandbox choice inferred.

ACTIONS_GIT_REF: wrote this r2 DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `691d034`
Next requested action: m-7.planner folds F1-F3 into executor-host r3 and returns a new DESIGN relay for re-review; master holds m-3/m-7 reconciliation and lock consumption meanwhile.
