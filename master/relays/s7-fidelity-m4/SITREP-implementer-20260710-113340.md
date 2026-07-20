## SITREP — m-4 fidelity verdict on s7 INV-CATALOG row 3: must revise the routing-column negative

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s7-fidelity-m4
PARENT_DISPATCH_ID: s7-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — bounded fidelity finding routes to master; VP integration and operator merge remain downstream
IN_REPLY_TO: master/relays/s7-fidelity-m4/SITREP-orchestrator-planner-20260710-112547.md
FROM: m-4.implementer
TO: master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-4.planner, m-7.planner, m-7.implementer
SUBJECT: m-4 fidelity at `s7-inv-catalog@35aabb9` — row 3 preserves the R2 wording and top-level model guard, but misses the load-bearing `routing_assignments.chosen_model` any-row negative

VERDICT: must-revise

### Scoped finding

1. **The law text and top-level leg are faithful.** `test/invariants/catalog.v1.json:95-99` says model identity is not gate-referenceable and cannot drive required, visible, or gate predicates. That matches the locked R2 boundary: model values are payload only (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:216-218`; `master/ARCHITECTURE.md:178-188,200-207`). `TestLawR2NoModelPredicate` also finds every top-level `model_identity` field, requires `gate_referenceable:false`, scans live required/visible predicates for those ids, and proves synthetic direct-field predicates over `model_name` fail load (`test/invariants/terminal_surface_test.go:187-252`).

2. **The routing-column negative is absent and the current mechanism accepts the forbidden shape.** The locked routing row puts `chosen_model` inside `routing_assignments` and says model values never enter required/visible/authority/lineage gates (`m-4 design:208,216-218`). The registered C1 carry requires `chosen_model` and every model-identity column to be non-predicate at column grain as soon as the grammar can address row-array columns (`master/ARCHITECTURE.md:501-503`). The live grammar does address `ARRAY.field`: it splits `routing_assignments.chosen_model`, then validates only the parent array (`internal/fieldspec/predicate.go:136-152`). Because the registry marks the carrier `routing_assignments` gate-referenceable (`internal/fieldspec/registry.json:174`), a synthetic `any_row:routing_assignments.chosen_model` reaches a compiled predicate. The s7 invariant derives its guarded ids only from top-level `FieldSpec.ModelIdentity` rows and never plants this any-row negative, so its absolute catalog claim is not executable at the routing grain.

3. **The allowed deviation atom is correctly distinguished, but does not close the negative.** The live registry keys both justification fields on `any_row:routing_assignments.declared_deviated == yes` (`registry.json:175-176`), matching the conceptual `declared_deviated == true` / bucket-vs-bucket contract (`m-4 design:208-211`; `ARCHITECTURE.md:200-207`). That is the allowed declared-policy atom. It does not make `chosen_model` a legal predicate or prove the C1 column guard.

### Required revision

- Keep the catalog claim unchanged; narrowing it to top-level model fields would drift from R2.
- Add synthetic required/visible `any_row:routing_assignments.chosen_model` negatives to `TestLawR2NoModelPredicate` and require registry load to reject them as non-gate-referenceable.
- The new negative should turn red on the current parser, exposing a pre-existing production enforcement gap. Preserve the s7 test-only fence: return that production correction to master for the proper FieldSpec/routing lane, then fold the test after the mechanism rejects the forbidden column predicate.
- Re-route row 3 to both m-2 and m-4 fidelity after the production guard and exact named invariant negative land.

### Verification

- `git rev-parse --short HEAD` in the dispatched worktree -> `35aabb9`.
- `go test -count=1 -run '^TestLawR2NoModelPredicate$' -v ./test/invariants` -> PASS, including only synthetic direct `required_when` and `visible_when` legs; this is evidence of the coverage gap, not closure of the omitted any-row leg.
- `go test -count=1 ./internal/fieldspec -run 'TestRegistryLoadRejectsBadPredicateReferences|TestRegistryLoadRejectsInvalidRows' -v` -> PASS; the existing any-row rejection test covers only a non-gate-referenceable parent array, while model-identity rejection covers only top-level fields.
- `git diff --check 1d3e92c..35aabb9` -> clean.
- Independent m-2 fidelity reached the same row-3 finding at `master/relays/s7-fidelity-m2/SITREP-implementer-20260710-113112.md:24`; this m-4 verdict is grounded independently in the routing lock and C1 carry.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s7-fidelity-m4/SITREP-implementer-20260710-113340.md` -> `OK`.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s7-fidelity-m4 master/relays/s7-fidelity-m4/SITREP-implementer-20260710-113340.md` -> dispatch root and exact file both `OK`.
- Scope held: no `frank/` source, test, branch, commit, worktree, or merge action by m-4; this relay and the append-only master index row are the only artifacts.

ACTIONS_GIT_REF: wrote `master/relays/s7-fidelity-m4/SITREP-implementer-20260710-113340.md` and appended its `master/relays/INDEX.md` row; no `frank/` edit or git action
FINAL_GIT_STATUS_SHORT: unavailable — `/Users/jack/Programming/harness` is not a git repository; dispatched worktree `35aabb9` remains `M .relays/s7/INDEX.md` plus untracked `.relays/s7/s7-inv-catalog-impl/` from the expected operational handoff, with no m-4 edits

Next requested action: master blocks row-3 fidelity, routes the production column-grain rejection to the proper lane without breaching s7's test-only fence, then returns the exact revised named invariant to m-2 and m-4.
