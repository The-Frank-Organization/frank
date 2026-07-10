## SITREP -- r3 approval exists; root-mode lint caveat remains visible

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s2-core-plan-review-implementer-r3-lint-caveat
PARENT_DISPATCH_ID: s2-core-plan-review-implementer-r3
RUN_ID: s2
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes - if delegated dispatch requires full root-mode lint clean beyond exact current-chain files
FROM: s2-core.implementer
TO: s2-core.planner
CC: s2.orchestrator-planner, operator
SUBJECT: post-review lint caveat -- existing r3 approval exact-file clean, but relay-root still fails on superseded r2 PLAN

Phase: post-PLAN-REVIEW verification caveat.

Current artifact:
- Existing r3 approval already present before this report: `.relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-033128.md`.
- Parent r3 PLAN: `.relays/s2/s2-core-plan/PLAN-planner-20260704-033012.md`.
- Orchestrator lineage fix request: `.relays/s2/s2-core-plan/SITREP-orchestrator-planner-20260704-032900.md`.

Claims:
- The r3 PLAN and existing r3 PLAN-REVIEW are exact-file lint clean — evidence E2 from `relay-lint.py`.
- Full root-mode lint is still not clean because the superseded r2 PLAN remains visible and still emits the same non-INDEX design-lock parent error — evidence E2 from `relay-lint.py --relay-root`.
- I am not writing a second PLAN-REVIEW verdict because `.relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-033128.md` already exists and is indexed as the r3 approving review.

Observed command output:

```text
$ python3 ~/.codex/skills/tools/relay-lint.py .relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-033128.md
OK .relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-033128.md

$ python3 ~/.codex/skills/tools/relay-lint.py --relay-root .relays/s2 .relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-033128.md
ERROR .relays/s2: INDEX.md: missing required header field ROLE
ERROR .relays/s2: INDEX.md: missing required header field PHASE
ERROR .relays/s2: INDEX.md: missing required header field AUTHORITY
ERROR .relays/s2: INDEX.md: missing required header field DISPATCH_ID
ERROR .relays/s2: INDEX.md: missing required header field CEREMONY_TIER
ERROR .relays/s2: INDEX.md: missing required header field EVIDENCE_TARGET
ERROR .relays/s2: INDEX.md: missing required header field HUMAN_GATE_REQUIRED
ERROR .relays/s2: s2-core-plan/PLAN-planner-20260704-030751.md: design-doc PLAN parent 's2-core-plan-review-implementer' does not resolve to a relay in this lineage
OK .relays/s2/s2-core-plan/PLAN-REVIEW-implementer-20260704-033128.md
```

Interpretation:
- The `INDEX.md` errors are the known routing-table noise.
- The superseded r2 PLAN error is not INDEX noise. If the dispatch gate requires full `--relay-root .relays/s2` clean output, then the active root still needs an explicit disposition for this visible superseded relay before delegated dispatch.
- If the intended dispatch gate is exact current-chain lint only, the existing r3 approval relay provides the needed approving PLAN-REVIEW parent.

Done:
- Verified the existing r3 approval relay exists and exact-file lints clean.
- Verified the current root-mode lint residue.
- Appended this report-only caveat instead of writing a duplicate r3 PLAN-REVIEW.

Not done:
- No source/test/sprint-spec edits.
- No dispatch.
- No attempt to move, delete, or rewrite the superseded r2 relay.

Blocked:
- Human/orchestrator decision if full root-mode lint cleanliness is required before delegated dispatch.

Tests / verification:
- `go test -count=1 ./...` -> pass across all packages (`cmd/frank` and `test/seatproc` have no test files; all test packages OK).
- `git status --short` before this report write -> no output.

ACTIONS_GIT_REF: no source/test/sprint-spec edits; wrote gitignored report relay `.relays/s2/s2-core-plan/SITREP-implementer-20260704-033312.md` plus `.relays/s2/INDEX.md` row.
FINAL_GIT_STATUS_SHORT: none - clean tree
