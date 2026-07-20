## DESIGN-REVIEW - m-3.implementer re-review of s8-design-m3-registry r1

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-design-m3-registry
PARENT_DISPATCH_ID: s8-design-m3-registry
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded r1 design re-review; no new operator decision surfaced
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m3-registry
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s8-grill-m3-registry
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: s8-design-m3-registry/DESIGN-planner-20260710-234500.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: approve - r1 folds F1 timeout fault edge, F2 I-PH param/verdict hygiene, F3 production-surface ground truth; m-7 reconcile still required

DESIGN_REVIEW_VERDICT: approve

I re-reviewed `s8-design-m3-registry/DESIGN-planner-20260710-234500.md` and the r1 design doc `master/domains/m-3-observation-evidence/design/2026-07-10-s8-check-registry-probe-design.md` against my prior must-revise relay `s8-design-m3-registry/DESIGN-REVIEW-implementer-20260710-232050.md`, the locked m-3 c2 table, the Step-2 kickoff, and current `frank@2e1b4f0`.

Approve for the m-3 r1 design surface. This approval does not close the master reconcile with `s8-design-m7-executor`, does not grant PLAN or IMPL authority, and does not assert that the overall s8 executor boundary is locked. It only confirms that the m-3 registry/probe/verdict-semantics document has folded the three prior m-3 review findings.

## Resolved Findings

### F1 - RESOLVED: timeout/process-kill is now the machinery-fault edge, not non-authority accepted+self_reported

The r1 doc now separates a check that was started and killed from a no-vantage/opaque-lane degradation. The timeout paragraph states that expiry takes the machinery-fault edge: authority records are `held`+escalate, and non-authority records are `rejected`/author-return with the fault edge named (`...s8-check-registry-probe-design.md:76-79`). The terminal-token line repeats that `blocked|degraded`/timeout never becomes `accepted`+label (`...:123`).

The negative fixture set now includes both sides of the distinction: killed check equals `held`/`rejected`+fault-edge, while no-vantage remains the distinct `accepted`+`self_reported` row (`...:153-154`). The resolved decisions and design-lock impact carry the same rule (`...:175`, `...:195`), and the fold-log explicitly records the F1 correction (`...:206`).

Targeted search confirmed the only remaining `accepted`+`self_reported` references around timeout are the explicit "not timeout" / paired no-vantage distinctions:
- `rg -n 'accepted.*self_reported|self_reported.*timeout|timeout.*accepted|machinery.*accepted|accepted.\+label|or `accepted`\+label' master/domains/m-3-observation-evidence/design/2026-07-10-s8-check-registry-probe-design.md`
- Hits: `:79`, `:123`, `:154`, `:175`, `:195`, `:206`; all are the corrected distinction, not an allowed timeout fallthrough.

### F2 - RESOLVED: check-param and verdict-output I-PH are pinned as an m-3-owned contract

The r1 doc now constrains the m-3-owned check surfaces before they reach the m-7 executor. The concrete rows now say `read-file.path` is lane-scoped/normalized, absolute paths and `..` are refused, `schema_ref` is a registry id rather than a path, and `run-suite.target` is a named-suite registry enum rather than command text (`...:56`, `...:65`).

New section 6.1 pins the I-PH contract: params are normalized and closed; path params are lane/evidence-root scoped; raw store/config/outbox/operator-channel/socket paths and effective config values never appear in `CheckVerdict`, bounces, `executable_claim_results`, or `failing_detail`; `failing_detail` is symbolic, bounded, and path-redacted; param refusal happens conductor-side before spawn; verdict redaction happens conductor-side after return (`...:127-130`). Section 7 now threads that exact contract into what m-3 hands to m-7 (`...:142`).

The negative fixture set, resolved decisions, design-lock impact, and fold-log all carry the I-PH requirement (`...:155`, `...:176`, `...:195`, `...:206`). Targeted search over the r1 doc confirmed these anchors for `absolute`, `..`, `schema_ref`, `named_suite`, command text, `failing_detail`, effective config, operator-channel/socket, path-free/path-redacted, conductor-side, and verdict-redaction.

### F3 - RESOLVED: ground truth is now production-surface scoped

The r1 ground-truth statement now says the production check/probe/executor surface is greenfield, with no production non-test executor subprocess path, while acknowledging existing test/fixture helper subprocesses (`...:20`). That is the correct distinction.

Verification against `frank@2e1b4f0`:
- `rg -n "exec\\.Command" frank --glob '!master-docs/**'` found only test/fixture helper subprocess uses.
- `rg -n "exec\\.Command" frank --glob '*.go' --glob '!**/*_test.go' --glob '!master-docs/**'` returned no production non-test hits.

## Confirmed Non-Blockers

1. The observer-only allowlist and byte-exact terminal tokens remain unchanged. The r1 design realizes the locked c2 mechanism; it does not reopen the `{accepted, rejected, held}` enum or give the executor a write path.

2. The VP-F4 probe boundary remains intact. The design still claims provided-surface absence, not absence of all possible same-uid reach, and it leaves the D5 ambient residual explicit.

3. The m-7 executor artifact/reconcile remains outside this approval. Section 7 consumes `spawn(check_id, params, lane_ref, timeout) -> CheckVerdict` plus m-7 guarantees, and section 10 keeps master reconcile with `s8-design-m7-executor` as a lock prerequisite. The s8 PLAN should consume only reconciled locks.

## Verification

- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-design-m3-registry/DESIGN-planner-20260710-234500.md` - OK.
- Read `s8-design-m3-registry/DESIGN-planner-20260710-234500.md` - confirmed `FROM: m-3.planner`, `TO: m-3.implementer`, `AUTHORITY: design-only`, `DESIGN_DOC_ID: s8-design-m3-registry`, and r1 re-review request.
- Read `master/domains/m-3-observation-evidence/design/2026-07-10-s8-check-registry-probe-design.md` with line numbers - checked r1 anchors at `:20`, `:56`, `:65`, `:76-79`, `:123`, `:127-130`, `:142`, `:153-155`, `:175-176`, `:195`, `:206`.
- `rg -n 'accepted.*self_reported|self_reported.*timeout|timeout.*accepted|machinery.*accepted|accepted.\+label|or `accepted`\+label' master/domains/m-3-observation-evidence/design/2026-07-10-s8-check-registry-probe-design.md` - only corrected distinction/no-vantage references.
- `rg -n "absolute|\\.\\.|schema_ref|named_suite|command text|failing_detail|effective config|operator-channel|socket|path-free|path-redacted|conductor-side|verdict-redaction|param normalization|I-PH" master/domains/m-3-observation-evidence/design/2026-07-10-s8-check-registry-probe-design.md` - confirmed I-PH anchors in the concrete rows, section 6.1, section 7, fixtures, resolved decisions, design-lock impact, and fold-log.
- `rg -n "exec\\.Command" frank --glob '!master-docs/**'` - test/fixture hits only.
- `rg -n "exec\\.Command" frank --glob '*.go' --glob '!**/*_test.go' --glob '!master-docs/**'` - no production non-test hits.
- `git -C frank rev-parse --short HEAD` - `2e1b4f0`; `git -C frank status --short` - clean.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended `master/relays/INDEX.md`; no design-doc edit by this seat, no `frank/` source/test edit, no `pcode/` edit, no PLAN, no IMPL.
FINAL_GIT_STATUS_SHORT: cwd root unavailable - `git status --short` returned `fatal: not a git repository (or any of the parent directories): .git`; `git -C frank status --short` returned clean.
Next requested action: m-3.planner return this approved r1 artifact to master for reconcile with `s8-design-m7-executor`; master should not consume the s8 PLAN until m-3/m-7 locks are reconciled.
RELAY_LINT: OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s8-design-m3-registry/DESIGN-REVIEW-implementer-20260710-235210.md`; dispatch root OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s8-design-m3-registry`.
