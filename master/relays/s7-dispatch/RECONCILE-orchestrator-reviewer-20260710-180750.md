## RECONCILE -- revise corrected s7 package: F2 is closed and F1 now catches unknown symbols inside current boundary files, but the hard-coded two-file discovery root still lets a new egress file ride outside the law

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s7-dispatch
PARENT_DISPATCH_ID: s7-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no -- one narrow test-only discovery-root fold remains; merge stays a later operator decision
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7-dispatch/RECONCILE-orchestrator-planner-20260710-175524.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-7.planner, m-7.implementer, m-1.implementer, m-2.implementer, m-4.implementer, m-2.planner, m-4.planner
SUBJECT: revise s7 re-review at `bbf3147` -- prior F2 fully closed; prior F1 closed at the unknown-symbol grain inside the two current files, but the discovery file set is itself an unguarded literal and the package acknowledges a new egress file remains invisible; s8 pinning cannot mechanically close a source-file addition

VERDICT: revise

## Finding

### F1-R2 -- BLOCKER: the discovery recognizes unknown symbols only after trusting a fixed two-file boundary

The prior ruling required discovery authoritative over the actual seat-egress boundary, with the r3 acceptance behavior that a future seat-visible family cannot ride outside the census (`RECONCILE-orchestrator-reviewer-20260710-030737.md:28`; `RECONCILE-orchestrator-reviewer-20260710-171215.md`, F1). The fold materially improves the old catalog-seeded scanner, but moves the unguarded assumption up one level:

- `discoverSeatEgressBoundary` hard-codes exactly `internal/channel/server.go` and `cmd/frank-mcp/mcp.go` (`test/invariants/path_hygiene_test.go:220-244`) and parses only that slice (`:246-298`).
- The scratch red added an unknown function inside `server.go`, one of those trusted files. It proves unknown function names are discovered when they call a recognized write-chain symbol inside the fixed slice. That closes the symbol-grain defect.
- A new production file is never parsed. For example, a helper in a new `internal/channel/*.go` file can call `serverConn.write`, while an existing boundary method invokes that helper through an unrecognized call in addition to its existing recognized writes. The current 17 strings and protocol-case counts remain unchanged, yet the helper can emit an additional seat-visible frame. The named law stays green.
- This is not hypothetical review invention: the incoming package explicitly calls a new egress file the tripwire's blind spot (`RECONCILE-orchestrator-planner-20260710-175524.md:25`), and the pair review records the same limitation (`<worktree>/.relays/s7/s7-inv-catalog-impl/RECONCILE-planner-20260710-173423.md:40-42`).

Deferring the literal file set to s8 section-7 pinning does not make this acceptance behavior mechanical. Adding a source file changes neither the digest-pinned catalog nor either file currently parsed by the test. The catalog can remain byte-identical while the uncensused surface rides green. Section-7 governance can govern an intentional census edit; it cannot detect a file the battery never discovers.

Required narrow fold: derive or guard the boundary file set independently across the production tree. A valid shape is a full-tree non-test AST/type-aware scan that fails when seat-egress write/encode/protocol-dispatch capability appears outside the registered centralized boundary, while retaining the exact site census inside it. Prove it with a command-pinned scratch red that adds a new production egress file plus an unknown helper call from the current boundary: the named law must fail before any expected census is edited; discard; real branch green. No production change lands.

## Closed And Accepted

- **Prior F2 is closed.** `TestLawIntakeOutcomeOneToOne` now calls `recover.RunWithProcessor` over one settled and one pending-zero intake, requires callback counts `{pending: 1, settled: 0}` with no other key, and keeps duplicate processing/cardinality as separate assertions (`test/invariants/intake_outcome_test.go:20-31,172-231`).
- **Prior F1 is closed within the current files.** Discovery no longer seeds from catalog-known symbols; it finds write-chain calls and protocol cases, and the existing-file unknown-symbol scratch red follows from the committed code.
- The current production egress topology is confined to the channel socket writer and MCP stdout encoder; the current 17-site census is coherent. The issue is future-file detection, not a missing current file.
- The fold commit is exactly two invariant files, +181, with no production delta. Row 3 and `catalog.v1.json` are byte-identical to `61cf35e`; m-4 stands.
- m-1 and m-2 round-2 confirms establish their store and bounce/reason halves were not weakened. They may stand through the next fold if those confirmed blocks and the catalog remain byte-identical; the remaining change is m-7's discovery-root mechanism.
- F3 status reporting is corrected exactly. The separate operator merge gate, staged catalog governance, s8 obligations, and disclosed race/flakes remain correctly preserved.

## Required Return

1. Route one test-only F1-R2 fold to the authorized m-7 Implementer. Preserve the six families, all captures/negatives, canonical-path census, F2, row 3, and catalog bytes.
2. Obtain m-7 pair review over the full-tree boundary guard and the genuine-new-file scratch red/discard/green proof. No new m-1/m-2/m-4 round is required if their protected blocks and catalog are proven byte-identical.
3. Return the new tip with focused row-6 verbose evidence, `go test -count=1 ./test/invariants`, serialized `go test -count=1 -p=1 ./...`, `go vet ./...`, diff/path/protected-byte proofs, scratch cleanup proof, and exact final status.
4. Keep the operator merge-decision relay unissued until VP approval.

## Verification

- Incoming exact-file lint: OK. Dispatch-root plus exact incoming lint: OK.
- `bbf3147` focused rows 6+9: PASS; row 6 enumerated all eleven canonical-path negatives.
- `go test -count=1 ./test/invariants`: PASS (`ok`, 1.286s).
- `go test -count=1 -p=1 ./...`: PASS, 25 tested packages `ok`, 2 no-test-file packages, zero failures.
- `go vet ./...`: PASS, no output.
- Protected-byte diff for `catalog.v1.json` and `terminal_surface_test.go`, `61cf35e..bbf3147`: exit 0, no output.
- Scratch branch `s7-vp-f1-egress-red`: absent; scratch worktree absent.
- Current production-wide output-primitive scan finds the channel socket writer and MCP stdout encoder in the two registered files; store/fsio writes are canonical persistence, not seat egress.

Next requested action: master routes the one discovery-root fold and returns the corrected package for VP re-review. Merge remains blocked on this verdict.

ACTIONS_GIT_REF: wrote this reviewer relay and appended its row to `master/relays/INDEX.md`; no `frank/` source, test, branch, commit, merge, tag, or remote action.
FINAL_GIT_STATUS_SHORT:
- workspace root: unavailable -- not a git repository (docs workspace)
- `frank/` main: `## main...origin/main`
- dispatched worktree `s7-inv-catalog@bbf3147`:
  ` M .relays/s7/INDEX.md`
  `?? .relays/s7/s7-inv-catalog-impl/`
