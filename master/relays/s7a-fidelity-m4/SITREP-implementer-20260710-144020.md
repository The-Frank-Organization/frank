## SITREP — m-4 fidelity verdict on s7a column-grain guard: confirm at the current grammar grain

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s7a-fidelity-m4
PARENT_DISPATCH_ID: s7a-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — scoped fidelity verdict; VP integration and operator merge remain downstream
IN_REPLY_TO: master/relays/s7a-fidelity-m4/SITREP-orchestrator-planner-20260710-143432.md
FROM: m-4.implementer
TO: master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-4.planner, m-2.planner, m-2.implementer, m-7.planner, m-7.implementer
SUBJECT: confirm the R2 routing face at `s7a-colgrain@d76c3ad`; prior row-3 blocker discharged, C1 enforced for the live any-row grammar, Step-3 nested-column representation still owed

VERDICT: confirm

### Scoped answers

1. **The prior row-3 must-revise is discharged by the production guard.** The any-row parser now validates both the parent array and the addressed column before compiling the predicate (`internal/fieldspec/predicate.go:136-155`). The column check is default-deny: only a column present in `GateReferenceableColumns` passes (`internal/fieldspec/registry.go:212-222`). The shipped routing row allowlists exactly `declared_deviated` (`internal/fieldspec/registry.json:174`), so `chosen_model` is non-predicate while the legal declared-policy atom remains available to both justification fields (`registry.json:175-176`). This restores the locked split: model values are payload only (`master/domains/m-4-routing-policy/design/2026-06-29-v3-routing-policy-design.md:216-218`) while deviation justification keys on the planner-declared boolean (`:208-211`; `master/ARCHITECTURE.md:200-207`).

2. **The fixture grain is sufficient and bites on both requested legs.** `TestRegistryLoadRejectsNonAllowlistedRowColumns` rejects `routing_assignments.chosen_model` under both `required_when` and `visible_when`, pins the registry-load error class text and full column reference, and rejects non-model `routing_assignments.seat` through the same path (`internal/fieldspec/registry_test.go:361-394`). The separate live-registry assertion pins the allowlist to the byte-exact singleton `declared_deviated`, and the positive evaluator runs that legal column. This is a structural allowlist, not a remembered `chosen_model` blocklist.

3. **C1 is discharged at the complete grain the current grammar can express.** C1 requires `chosen_model` and every model-identity column to be non-predicate once row-array columns are addressable (`master/ARCHITECTURE.md:501-503`). Empty-or-absent allowlist means deny, so every current or newly named routing column is non-predicate unless governance explicitly adds it to the positive allowlist. That satisfies C1 behavior now and is stronger against forgotten future model columns than a model-identity-only blocklist.

4. **C1 residue is representational, not an open enforcement gap.** Step-3 still owes the full first-class nested-column FieldSpec schema reserved by s5-(f). At that transition, `chosen_model` and any model-identity column must preserve the same effective false default through explicit per-column `model_identity` / `gate_referenceable:false` metadata or an equivalent default-deny representation, and the named row-3 invariant must remain. No additional C1 behavior is owed before that schema transition. C2, the routing-record deviation-justification coupling at write time, remains a separate registered Step-3 carry and is not closed by this C1 verdict.

5. **Claim wording is bounded and accurate.** “R2 enforced at column grain” is supported when it means the live `any_row:<array>.<column>` grammar and the shipped registry at `d76c3ad`. It must not be expanded to “the full Step-3 nested-column schema is complete” or “all future schema representations are already covered.”

### Verification

- Worktree tip and cleanliness: `git rev-parse --short HEAD` -> `d76c3ad`; `git status --short` -> empty.
- Focused E2: `go test -count=1 ./internal/fieldspec -run 'TestRegistryLoadRejectsNonAllowlistedRowColumns|TestPredicateAtomsAndCombinators|TestRegistryS5MemberContainsRegistryPassRows' -v` -> PASS, including all three default-deny negatives and the legal declared-deviation positive.
- Full E2: `go test -count=1 ./...` -> exit 0; 24 packages `ok`, 2 packages `[no test files]`, zero failures.
- Static E2: `go vet ./...` -> exit 0 with no output; `git diff --check 1d3e92c..d76c3ad` -> clean.
- Red/green lineage E1: red `10ee3a2` is parented directly on `1d3e92c`; green `d76c3ad` is parented directly on `10ee3a2`. The red commit contains the three rejecting fixtures; the green commit contains guard, allowlist member, singleton metadata, version marker, and positive retarget atomically.
- Scope held: no `frank/` source, test, branch, commit, worktree, PR, or merge action by m-4; this relay and the append-only master index row are the only artifacts.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/s7a-fidelity-m4/SITREP-implementer-20260710-144020.md` -> `OK`.
- `python3 /Users/jack/.codex/skills/tools/relay-lint.py --relay-root master/relays/s7a-fidelity-m4 master/relays/s7a-fidelity-m4/SITREP-implementer-20260710-144020.md` -> dispatch root and exact file both `OK`.

ACTIONS_GIT_REF: wrote `master/relays/s7a-fidelity-m4/SITREP-implementer-20260710-144020.md` and appended its `master/relays/INDEX.md` row; read-only review of `s7a-colgrain@d76c3ad`; no `frank/` edit or git action
FINAL_GIT_STATUS_SHORT: unavailable — `/Users/jack/Programming/harness` is not a git repository; dispatched worktree `s7a-colgrain@d76c3ad` is clean

Next requested action: combine this confirm with the m-7 hosting fidelity leg for VP integration; operator merge remains separate, after which the s7 row-3 invariant fold returns to m-2 and m-4 as already routed.
