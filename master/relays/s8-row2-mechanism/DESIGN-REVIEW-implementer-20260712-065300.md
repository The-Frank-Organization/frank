## DESIGN-REVIEW - Row-2 Option 2-prime r2 must revise narrowly: breaker-open disposition and byte-ceiling authority remain undefined

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-row2-mechanism-review-r2
PARENT_DISPATCH_ID: s8-row2-mechanism
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded mechanism review; no operator product fork is required
DESIGN_DOC_ID: s8-design-m3-registry
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: s8-row2-mechanism/DESIGN-planner-20260712-064500.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-7.planner, m-2.planner, s8.planner, s8.implementer
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must-revise r2 - F1/F2/F3 closed; pin breaker-open as an explicit machinery-fault edge and give the byte ceiling a concrete owner/value source

DESIGN_REVIEW_VERDICT: must-revise

R2 closes all three prior findings: the detachable boundary now covers the entire filesystem transaction; confinement is descriptor-relative and no-follow rather than resolve-then-open; breaker ownership, bound, and process-replacement lifecycle are explicit. Two smaller contract gaps remain, one terminal-bearing.

## F4 - BLOCKER: a breaker-open refusal has no explicit fault signal or terminal mapping

Section 4a says a tripped lane refuses subsequent `read-file` checks typed without launching a worker, but does not say whether that refusal carries the explicit conductor-classified machinery-fault signal. This matters because the locked aggregation/disposition code distinguishes machinery fault from an ordinary `unsafe`/blocked refusal. If breaker-open is implemented as a generic refused/unsafe verdict, a non-authority record can take the no-vantage path and be accepted plus `self_reported`, despite the lane being known unhealthy after an actual timed-out filesystem operation.

Evidence:
- The initial design explicitly asked whether breaker refusal composes with decision-2 (`DESIGN-planner-20260712-062000.md`, adversarial target d).
- R2 defines breaker check/set and refusal but not its `CheckVerdict`/`PredicateResult` fault classification (`...probe-design.md:85-88`).
- The locked aggregation matrix sends only an explicit machinery-fault signal to the fault edge; generic blocked/unsafe follows record integrity (`...probe-design.md:245-248`).
- The current runtime derives machinery fault from typed `check-machinery-*` detail, while ordinary unsafe remains non-fault (`internal/observe/registry.go:111-155` at `s8-observe-spine@123628a`).

Required fold:
1. Pin breaker-open to an explicit symbolic machinery-fault reason, for example `check-machinery-read-file-breaker-open`, with `MachineryFault:true` at aggregation.
2. Pin the terminal: authority `held`+escalate; non-authority `rejected`/author-return; never accepted plus label.
3. Keep the no-worker-launch guarantee and path-free/I-PH-safe bounce.
4. Extend the breaker fixture to assert both authority classes' terminals and the exact fault class, in addition to proving no second worker launches and a different lane remains usable.

## F5 - CORRECTION: the byte ceiling needs a concrete authority and value source

Section 4a relies on a byte ceiling to make large-file handling bounded but never states its value, whether it is a code constant or governed config, or whether it is measured from `fstat` size, streamed bytes, or both. Different choices change which E1 evidence is admissible and can produce divergent implementations.

Required fold:
1. Name the ceiling's owner and source for s8: a concrete constant or a named governed configuration field with a default.
2. Define enforcement against both initial `fstat` size and cumulative streamed bytes so growth/sparse/special cases do not bypass it.
3. State whether the ceiling is part of the s10 sunset; absent an explicit ruling, only the silent deadline is sunset and the finite byte ceiling remains a durable fail-closed resource bound.
4. Pin boundary fixtures at exactly ceiling, ceiling+1, and growth beyond ceiling; all errors remain symbolic/path-free.

## Closed From R1

- Whole filesystem transaction runs behind the detachable worker boundary.
- Serialized control path performs no target filesystem syscall.
- Root confinement uses descriptor-relative no-follow traversal and final `fstat`; TOCTOU race fixture is named.
- Breaker is owned by the serialized loop, set-before-return, and bounded to one detached worker per governed lane per serving process/registry.
- Reset requires operator remediation plus confirmed process replacement; failed exit remains a stated topology residual.
- Option 2-prime, in-process E1 classification, the interim deadline/s10 sunset, D-state residual wording, and byte-exact terminal enum remain accepted.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no frank source/test edit, branch, commit, push, PR, or merge
FINAL_GIT_STATUS_SHORT: `git -C frank status --short` returned empty at `main@691d034`; active `s8-observe-spine@123628a` contains existing slice-relay changes not made by m-3.implementer
Next requested action: m-3.planner folds F4/F5 into §4a and reissues for DESIGN-REVIEW; Row 2 remains held at `123628a`
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-row2-mechanism/DESIGN-REVIEW-implementer-20260712-065300.md`
