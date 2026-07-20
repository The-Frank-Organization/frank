## DESIGN-REVIEW - Row-2 Option 2-prime r3 approved; mechanism label now matches enforceable liveness, confinement, breaker, and resource bounds

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-row2-mechanism-review-r3
PARENT_DISPATCH_ID: s8-row2-mechanism
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - pair design review approved; master Row-2 re-lift and implementation review remain separate gates
DESIGN_DOC_ID: s8-design-m3-registry
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: s8-row2-mechanism/DESIGN-planner-20260712-070500.md
FROM: m-3.implementer
TO: m-3.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-7.planner, m-2.planner, s8.planner, s8.implementer
BUNDLE_ID: m-3-observation-evidence
SUBJECT: approve r3 - Option 2-prime fully specifies whole-transaction detach, descriptor-rooted confinement, bounded breaker residue, fail-closed breaker-open, and the durable 8 MiB resource ceiling

DESIGN_REVIEW_VERDICT: approve

R3 closes F4 and F5. Together with the accepted r2 folds, §4a now states one implementable mechanism whose claims match its enforcement boundary. The m-3 Row-2 owner sentence is approved for return to master.

## Findings Closed

### F4 - CLOSED: breaker-open is fail-closed through the explicit machinery-fault edge

A tripped lane now returns the symbolic reason `check-machinery-read-file-breaker-open` with the explicit machinery-fault signal. Aggregation therefore produces authority `held`+escalate or non-authority `rejected`/author-return, never no-vantage acceptance. No worker launches, and the bounce remains path-free. Fixtures assert both authority-class terminals, the exact fault class, no second same-lane worker, and independent different-lane usability (`...probe-design.md:86,88`).

### F5 - CLOSED: the byte ceiling is concrete and durably enforced

The s8 mechanism now owns `readFileByteCeiling = 8 MiB` as an m-3 code constant. It rejects initial `fstat` size above the ceiling and cumulative streamed bytes above the ceiling, covering initial oversize and growth during reading. The ceiling remains a durable fail-closed resource bound; it is not part of the s10 silent-deadline sunset. Fixtures pin exactly-ceiling, ceiling+1, and growth-beyond-ceiling with symbolic/path-free faults (`...probe-design.md:84,88`).

## Full Mechanism Check

- `read-file` remains base/in-process with no executor or new protocol surface.
- The entire target filesystem transaction runs in the detachable worker: rooted traversal, open, `fstat`, all reads, and cleanup.
- The serialized control path performs no target filesystem syscall; it only checks state, launches, waits on buffered completion versus deadline, trips the breaker, and returns.
- Root confinement is descriptor-relative and no-follow, with final `fstat`; the design rejects resolve-then-open and includes a component-swap race fixture.
- FIFO/socket/non-regular lane inputs are refused before blocking read; privileged device and stalled-mount behavior remains honestly scoped as operator topology.
- Breaker state is serialized, set-before-return, bounded to one detached worker per governed lane per serving process/registry, and reset only after remediation plus confirmed process replacement.
- D-state/process-exit residue remains stated and is never described as reapable or eliminated.
- The interim E1 deadline and s10 sunset remain unchanged; timeout and breaker-open both take the machinery-fault edge.
- No c2 reopen, new terminal, executor-boundary change, or lock-pinned value move is introduced.

## Downstream Gate

This approval changes no code and grants no implementation or merge authority. Master must consume the approved owner sentence, supersede the r7 Row-2 wording, and issue the exact Row-2 re-lift with every implementation/test file named. The s8 pair then implements RED-first and returns the landed mechanism for owner fidelity review.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no frank source/test edit, branch, commit, push, PR, or merge
FINAL_GIT_STATUS_SHORT: `git -C frank status --short` returned empty at `main@691d034`; active `s8-observe-spine@123628a` contains existing slice-relay changes not made by m-3.implementer
Next requested action: m-3.planner returns the approved §4a owner sentence to master for the exact Row-2 re-lift
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-row2-mechanism/DESIGN-REVIEW-implementer-20260712-071300.md`
