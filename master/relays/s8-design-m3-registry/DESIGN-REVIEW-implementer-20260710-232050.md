## DESIGN-REVIEW - m-3.implementer review of s8-design-m3-registry

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-design-m3-registry
PARENT_DISPATCH_ID: s8-design-m3-registry
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded m-3 design revision; no new operator decision surfaced
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m3-registry
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s8-grill-m3-registry
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: s8-design-m3-registry/DESIGN-planner-20260710-231500.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: must-revise - timeout fault-edge wording must not permit non-authority accepted+self_reported; check-param/verdict-output I-PH must be pinned; ground-truth exec.Command statement needs production-surface correction

DESIGN_REVIEW_VERDICT: must-revise

I reviewed `s8-design-m3-registry/DESIGN-planner-20260710-231500.md` and the design doc `master/domains/m-3-observation-evidence/design/2026-07-10-s8-check-registry-probe-design.md` against the corrected master dispatch, the locked c2 m-3 record, the Step-2 kickoff, current `frank@2e1b4f0`, and the m-7 claim-boundary/I-PH contracts.

The core shape is close, but I cannot approve the design lock yet. Two m-3-local blockers remain, plus one required ground-truth correction.

## Findings

### F1 - BLOCKER: timeout is a machinery fault and must not fall through to non-authority accepted+self_reported

The new design correctly says timeout is "machinery-could-not-complete," not predicate-false. But it then allows the non-authority timeout path to become either `rejected`+fault-edge or `accepted`+`self_reported` when "non-critical and honestly degradable."

Evidence:
- `master/domains/m-3-observation-evidence/design/2026-07-10-s8-check-registry-probe-design.md:76-78` classifies expiry as a machinery fault, then permits non-authority `accepted`+`self_reported` in the same timeout paragraph.
- `master/domains/m-3-observation-evidence/design/2026-07-10-s8-check-registry-probe-design.md:122` repeats that `blocked|degraded`/timeout can become non-authority `rejected`+fault-edge or `accepted`+label.
- The locked m-3 table distinguishes no-vantage from machinery-broke: no-vantage/opaque non-authority can accept+label, but "conductor check machinery ran-and-broke" is non-authority `rejected`/author-returned and authority `held` (`master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md:71-72`).
- The Step-2 kickoff is explicit: "observe-machinery internal fault/timeout: authority -> held, non-authority -> rejected/author-return" (`master/STEP-2-KICKOFF.md:47`).

Required revision:
- Make timeout/process-kill/check-machinery failure the fault edge: authority-class -> terminal `held` + escalate; non-authority -> terminal `rejected` / author-return with fault edge named.
- Keep `accepted`+`self_reported` only for the separate no-vantage/opaque-lane degradation row, not for timeout or executor machinery failure.

### F2 - BLOCKER: the check registry lacks the I-PH/path and verdict-output contract for new m-3-controlled surfaces

The design owns closed check param schemas and the `CheckVerdict` shape, but does not constrain `read-file.path`, suite target naming, or `failing_detail` output. That leaves a faithful builder room to expose raw conductor-internal paths, config/outbox/store/socket paths, command output, or file contents through a new check surface.

Evidence:
- The planner relay says param schemas are part of the requested m-3 face (`master/relays/s8-design-m3-registry/DESIGN-planner-20260710-231500.md:26`).
- The entry descriptor promises a closed, conductor-validated `param_schema` (`master/domains/m-3-observation-evidence/design/2026-07-10-s8-check-registry-probe-design.md:39`).
- The concrete E1 entry is only `{ path, expect: line | hash | schema_ref }`, with no normalization/root rule, no absolute/`..` denial, no store/config/outbox/operator-channel/socket denial, and no statement that `schema_ref` is a registry id rather than a path (`...s8-check-registry-probe-design.md:56`).
- The verdict shape contains `failing_detail` without a redaction/ceiling/path-free contract (`...s8-check-registry-probe-design.md:108-113`).
- Step-2 requires I-PH across every new surface: no store/config/socket path anywhere seat-visible (`master/STEP-2-KICKOFF.md:52`).
- m-7's locked guardrail forbids raw conductor-internal store/config/outbox/operator-channel paths and effective config values in seat-deliverable surfaces (`master/domains/m-7-conductor-core/design/2026-07-01-v3-conductor-core-design.md:125,168`).
- m-3 hands m-7 the closed param schemas and verdict shape; this is not solely m-7 host policy (`...s8-check-registry-probe-design.md:136`).

Required revision:
- Add the m-3-owned check-param and result-output hygiene contract before lock. At minimum: params are normalized and closed; paths are scoped to the declared lane/work artifact or other approved evidence roots; absolute paths and traversal are refused; raw store/config/outbox/operator-channel/socket paths and effective config values never appear in verdicts, bounces, rows, or failing details; `run-suite.target` is a registry/named-suite enum rather than command text; `failing_detail` is symbolic/bounded/path-redacted.
- Add a negative fixture or explicit acceptance item for the above, or mark the exact fixture as m-7-hosted if the execution host owns the probe while m-3 owns the schema/output contract.

### F3 - REQUIRED CORRECTION: ground truth overstates subprocess absence

The doc's greenfield statement says "no `exec.Command` anywhere in-tree." That is false as written. The production executor/check/probe surface is greenfield, but the tree already has test/fixture helper subprocess usage.

Evidence:
- `master/domains/m-3-observation-evidence/design/2026-07-10-s8-check-registry-probe-design.md:20` says no `/check`, `/probe`, `/executor` package and no `exec.Command` anywhere in-tree.
- `rg -n "exec\\.Command" frank --glob '!master-docs/**'` found existing test/fixture helper subprocesses, including `frank/test/fixtures/s4_iph_test.go:174`, `frank/test/fixtures/main_assembly_test.go:43`, and `frank/test/invariants/path_hygiene_test.go:829`.
- `rg -n "exec\\.Command" frank --glob '*.go' --glob '!**/*_test.go' --glob '!master-docs/**'` returned no production non-test hits, and `rg --files frank | rg '(^|/)(check|probe|executor)(/|$)'` returned no package directories.

Required revision:
- Replace the statement with production-surface wording, for example: no production check/probe/executor package and no production executor subprocess path exists; existing test/fixture helper subprocesses are prior local test harness patterns, not an observe-check executor.

## Confirmed Non-Blockers

1. No direct executor write path is introduced by the m-3 semantics once F2 is fixed. The design says the executor returns a typed verdict only, holds no store handle, has no §3.1 write-allowlist entry, and the conductor computes `executable_claim_results`, `achieved_evidence`, and `target_gap_result` (`...s8-check-registry-probe-design.md:106-124`). That matches the locked observer-only write-allowlist (`...2026-06-29-v3-observe-evidence-design.md:61`).

2. The VP-F4 claim boundary is mostly honored. The design claims provided-surface absence, explicitly does not claim no possible reach, and reports the same-uid ambient residual separately (`...s8-check-registry-probe-design.md:94-98`). That matches the corrected m-7 executor dispatch's claim ceiling (`master/relays/s8-design-m7-executor/DESIGN-orchestrator-planner-20260710-214607.md:21-24`) and the VP correction (`master/relays/step2-prep/RECONCILE-orchestrator-reviewer-20260710-201733.md:39-45`).

3. No c2 mechanism reopen is required. The design realizes the §13 carry: concrete registry entries, operator gate, additive-entry path, and suite-class executor isolation (`...2026-06-29-v3-observe-evidence-design.md:211-221`; `...s8-check-registry-probe-design.md:16,185-193`). The byte-exact terminal enum remains `{accepted, rejected, held}`; F1 is a wording/semantics correction inside the existing table, not a new token.

4. The m-7 executor design is not yet authored in `master/domains` or relayed for m-7. That is not an m-3 approval blocker by itself, because this artifact already leaves master reconcile with `s8-design-m7-executor` as a lock prerequisite (`...s8-check-registry-probe-design.md:138,190-193`). Do not collapse this review into overall s8 executor closure.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-design-m3-registry/DESIGN-planner-20260710-231500.md` - OK
- `nl -ba master/relays/s8-design-m3-registry/DESIGN-planner-20260710-231500.md` - confirmed `FROM: m-3.planner`, `TO: m-3.implementer`, `DESIGN_DOC_ID: s8-design-m3-registry`, and review request.
- `nl -ba master/domains/m-3-observation-evidence/design/2026-07-10-s8-check-registry-probe-design.md` - inspected the full design doc, especially lines 16, 20, 39, 56, 76-80, 108-124, 128-138, and 185-193.
- `nl -ba master/domains/m-3-observation-evidence/design/2026-06-29-v3-observe-evidence-design.md | sed -n '29,100p;180,230p'` - checked locked §2/§3.1/§3.3/§4/§13.
- `nl -ba master/STEP-2-KICKOFF.md | sed -n '44,64p'` - checked Step-2 gate/I-PH/timeout obligations.
- `nl -ba master/relays/s8-design-m7-executor/DESIGN-orchestrator-planner-20260710-214607.md` and `nl -ba master/relays/step2-prep/RECONCILE-orchestrator-reviewer-20260710-201733.md` - checked the F4 claim boundary and m-7/m-3 split.
- `git -C frank rev-parse --short HEAD` - `2e1b4f0`; `git -C frank status --short` - clean.
- `rg -n "exec\\.Command" frank --glob '!master-docs/**'` - test/fixture subprocess hits exist; production-only grep returned no hits.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no design-doc edit by this seat, no `frank/` source/test edit, no `pcode/` edit, no PLAN, no IMPL.
FINAL_GIT_STATUS_SHORT: cwd root unavailable - `git status --short` returned `fatal: not a git repository (or any of the parent directories): .git`; `git -C frank status --short` returned clean
Next requested action: m-3.planner revise the s8 design doc for F1/F2/F3, then send a revised DESIGN relay back to m-3.implementer for re-review. Master should hold reconcile/PLAN consumption of `s8-design-m3-registry` until the revised review approves it.
RELAY_LINT: OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-design-m3-registry/DESIGN-REVIEW-implementer-20260710-232050.md`; dispatch root OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s8-design-m3-registry`
