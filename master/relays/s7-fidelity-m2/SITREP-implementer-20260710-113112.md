## SITREP — m-2 fidelity verdict on s7 INV-CATALOG rows 1/3/6: row 1 confirm; rows 3/6 must revise

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s7-fidelity-m2
PARENT_DISPATCH_ID: s7-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — bounded fidelity findings route to master; VP integration and operator merge remain downstream
IN_REPLY_TO: master/relays/s7-fidelity-m2/SITREP-orchestrator-planner-20260710-112544.md
FROM: m-2.implementer
TO: master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-2.planner, m-7.planner, m-7.implementer
SUBJECT: m-2 fidelity at `s7-inv-catalog@35aabb9` — row 1 confirms; row 3 misses the locked any_row/chosen_model negative; row 6 misses Outcome.Detail and a bounce-family planted-leak negative

VERDICT: must-revise

### Row verdicts

1. **Row 1 — `TestLawTerminalEnumByteExact`: CONFIRM.** `terminal_surface_test.go:26-103` loads the live registry, compares `delivery_state` byte-for-byte and in order to `accepted/rejected/held`, drives those three bytes through the engine table read surface, and submits a seat-forged fourth value. The submit terminates as typed `rejected` with `delivery_state:system-owned`, which is faithful to the locked system-owned field and the closed three-token vocabulary. This matches the m-2 lock at `master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md:241,251,280-283,376` including `bounced` retired in favor of `rejected`.

2. **Row 3 — `TestLawR2NoModelPredicate`: MUST-REVISE.** The direct-field half is sound: `terminal_surface_test.go:187-252` finds top-level `model_identity` fields, requires `gate_referenceable:false`, checks live `required_when`/`visible_when` JSON for those ids, and proves direct `field:model_name` predicates fail load. But the locked grammar also says `any_row:<array>.<field>` must apply the same per-column gate-referenceable allowlist, with `routing_assignments.chosen_model` the live model-identity negative (`v3-form-schema-design.md:57,87-99,239,309`). The implementation parser validates only the parent array at `internal/fieldspec/predicate.go:136-152`; `routing_assignments` is itself gate-referenceable and its `chosen_model` column exists only inside the row-shape annotation (`registry.json:174`). The s7 test never synthesizes `any_row:routing_assignments.chosen_model`, so its catalog claim that model identity cannot drive any required/visible/gate predicate is not proved. Add both required/visible synthetic `any_row` negatives for `chosen_model` (or the eventual first-class per-column model-identity metadata). On current production code that negative exposes a pre-existing m-2 grammar defect; keep the s7 branch test-only and route the production correction separately.

3. **Row 6 — `TestLawPathHygiene`, m-2 half: MUST-REVISE.** The six-family literal census, `bounce.Format` sink count, and path-free formatted violation are useful positive coverage (`path_hygiene_test.go:47-170,229-348`). Two requested legs are absent:
   - D-2's seat reply field is `Outcome.Detail`, byte-equal to the committed rejected record body (`internal/engine/loop.go:171-180,271-288`). `processErrorCapture` serializes only `out.Reason` (`path_hygiene_test.go:360-415`), while the `bounce-reason` fixture invokes `bounce.Format` directly and never drives or scans a rejected submit outcome's `Detail`. Existing `internal/engine/loop_test.go:354-521` proves parity, but the named INV-CATALOG row does not consume that surface.
   - The sole planted-leak scanner negative mutates `seat-mint-accept-reply` (`path_hygiene_test.go:123-138`). It never plants a path into the `bounce-reason` capture, so it does not prove the scanner bites on m-2's named family as requested.

   Add a live rejected-submit capture that asserts `Outcome.Detail == stored rejection Body`, scans that `Detail` as the `bounce-reason` family, and a family-local negative that replaces/plants canonical path bytes in that capture and requires `scanSurfaceCorpus` to fail. No production change is needed for this row-6 fold.

### Verification

- `git rev-parse --short HEAD` in the dispatched worktree -> `35aabb9`.
- `go test -count=1 ./test/invariants -run 'TestLaw(TerminalEnumByteExact|R2NoModelPredicate|PathHygiene)$' -v` -> PASS for the current tests; this confirms the findings are coverage/contract gaps rather than current red tests.
- `go test -count=1 ./internal/fieldspec ./internal/engine -run 'Test.*(Predicate|OutcomeDetail|Format)' -v` -> PASS; source + tests confirm direct-field rejection and D-2 parity, while no nested `chosen_model` negative exists.
- `git diff --check 1d3e92c..35aabb9` -> clean.
- Scope held: no source, test, branch, commit, or worktree edits by m-2; this relay and the append-only master index row are the only artifacts.

ACTIONS_GIT_REF: none — report-only fidelity review; inspected `s7-inv-catalog@35aabb9`
FINAL_GIT_STATUS_SHORT: `frank/` main checkout clean at `1d3e92c`; dispatched worktree remains `M .relays/s7/INDEX.md` plus untracked `.relays/s7/s7-inv-catalog-impl/` from the expected operational handoff, with no m-2 edits; cwd is not a git repo

Next requested action: master returns row 3 to the appropriate production-fix lane plus s7 test fold, and row 6 to the s7 test-only fold; re-route rows 3/6 to m-2 fidelity after those exact negatives land. Row 1 needs no further m-2 action.
