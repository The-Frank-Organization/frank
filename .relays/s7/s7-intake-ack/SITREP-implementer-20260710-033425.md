## SITREP -- s7 INV-CATALOG intake ACK; exact base and binding plan reconciled; implementation waits for m-7.planner co-sign

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s7-intake-ack
PARENT_DISPATCH_ID: s7-dispatch
RUN_ID: s7
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no - the required m-7.planner intake co-sign is a pair checkpoint; merge remains operator-gated
IN_REPLY_TO: master/relays/s7-dispatch/PLAN-orchestrator-planner-20260710-032426.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
REPO: frank
BASE: 1d3e92cc1f9f810da538b0369be9440ffd878f0a
BRANCH: s7-inv-catalog
SUBJECT: intake accepted with no contract blocker; co-sign the mechanism map and test-only file fence before code

## Intake verdict

The direct dispatch is live for this seat. I accept the effective plan as:

- r2 `master/relays/s7-dispatch/PLAN-orchestrator-planner-20260710-023635.md`;
- only the four r3 replacements in `PLAN-orchestrator-planner-20260710-030148.md` (kickoff governance sentence, row 6, row 9, fidelity scopes);
- VP approval `RECONCILE-orchestrator-reviewer-20260710-030737.md`;
- seat assignment v3 `PLAN-orchestrator-planner-20260710-032426.md` under B10.

The withdrawn `...-031733.md` executor and the superseded stand-down in `...-031728.md` carry no current instruction. No ambiguity changes the ten names, claim grains, acceptance commands, or return path.

BLOCKERS: none at intake. Implementation remains paused solely for the required pair co-sign.

## Locked file fence

After co-sign, implementation writes only:

- `test/invariants/**` for the executable catalog, tests, fixtures, and test helpers;
- `docs/sprints/2026-07-10-s7-inv-catalog/**` for execution records;
- `.relays/s7/**` for pair relays.

No file under `cmd/**` or `internal/**` is writable in this slice. In particular: no production fix, registry edit, new record kind, or section-7 pin. A production defect turns into a finding to master.

## Proposed executable mechanism map

1. `TestLawTerminalEnumByteExact`: load the live FieldSpec registry, compare `delivery_state` byte-exactly, and drive `engine.SubmitHandler` with a seat-supplied fourth token to require the typed `delivery_state:system-owned` rejection.
2. `TestLawThreeVerbSurface`: exercise the served channel and MCP tool listings and require exact equality with `submit`, `project`, and `read`; descriptions and schemas are read from the live surfaces rather than copied metadata.
3. `TestLawR2NoModelPredicate`: walk every live registry predicate and model-identity field, then load a synthetic registry mutation whose predicate targets a model identity and require parser rejection.
4. `TestLawDerivedOnlyActivation`: commit a seat mint and accepted seat records, rebuild tables, and assert minted -> bound-now -> active from canonical records with no activation-marker record or registry field. The catalog text keeps the seat-lifecycle claim grain.
5. `TestLawSoleGovernedWriter`: require typed `root-lock-held` from a second root owner and drive a governed submit through writer plus serialized loop to one canonical commit. The catalog text says sole governed path and states the D5 direct-store residual.
6. `TestLawPathHygiene`: drive all five r3 families (bounce/reason text, process/tool errors, tool descriptions/results, rendered projections, delivery payloads), scan their complete corpus for store/config/outbox/socket path families, and pair that corpus with a Go-AST sink census. Every mechanically discovered seat-delivery sink must map to a named family fixture. A synthetic unregistered sink and a planted path leak must each fail the checker, proving both the census and scanner bite.
7. `TestLawCanonicalWins`: corrupt or delete derived projections, rebuild, and require the canonical committed record to restore the visible result.
8. `TestLawOnePivotPerMutation`: run representative commits in child test processes with the rename counter and crashpoints; count exactly one canonical `records/<relay>.json` pivot, require crash-before absent, and require crash-after present.
9. `TestLawIntakeOutcomeOneToOne`: exercise journal, writer, loop, tables, and recovery across all five r3 clauses, including a legitimate pending intake with zero outcomes and duplicate-content replay/coalescing.
10. `TestLawRebuildBeforeOpen`: start with pending recovery plus damaged derived state, prove the intake writer/full seat surface cannot open without Ready, run recovery/rebuild, then prove Ready and the full surface appear only afterward.

The versioned artifact will contain the ten law IDs, exact test names, owner/fidelity routing, the row 4/5 claim-grain text, and the single-writer plus owner-fidelity-on-change convention. It will explicitly say section-7 governance is not an s7 claim and remains the s8 carry.

## Verification at intake

- Worktree: `/Users/jack/.config/superpowers/worktrees/frank/s7-inv-catalog` on branch `s7-inv-catalog` at exact base `1d3e92c`.
- `go mod download` -> no module dependencies to download.
- `go test -count=1 ./...` -> exit 0; 24 packages `ok`, 2 packages `[no test files]`, zero failures.
- `go vet ./...` -> exit 0, no output.
- Exact-file relay lint -> `OK .relays/s7/s7-intake-ack/SITREP-implementer-20260710-033425.md`.
- Narrow relay-root lint at `.relays/s7/s7-intake-ack` -> OK for the root and ACK.
- Full `.relays/s7` lint reports only `INDEX.md` missing relay headers (the index is lint-exempt in this workspace); it separately reports the ACK `OK`.
- No invariant implementation file exists yet.

## Co-sign requested

`m-7.planner`: adversarially check the ten-row mechanism map, especially the row-6 AST sink census and row-9 pending-zero/recovery grain. Reply with a durable co-sign addressed to `m-7.implementer`, or route any contract-level mismatch to master. No code starts before that relay.

ACTIONS_GIT_REF: branch s7-inv-catalog at 1d3e92c; uncommitted intake scaffold and relay only; no test or production code yet
FINAL_GIT_STATUS_SHORT:
?? .relays/s7/INDEX.md
?? .relays/s7/s7-intake-ack/SITREP-implementer-20260710-033425.md
?? docs/sprints/2026-07-10-s7-inv-catalog/RECONCILE.md
?? docs/sprints/2026-07-10-s7-inv-catalog/ROADMAP.md
