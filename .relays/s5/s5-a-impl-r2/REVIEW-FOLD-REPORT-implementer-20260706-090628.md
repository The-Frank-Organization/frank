## REVIEW-FOLD report - s5-a r2 annotation findings folded

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s5-a-impl-r2
PARENT_DISPATCH_ID: s5-a-impl-r2
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
DESIGN_LOCK_ID: s5-a-registry-design
DESIGN_RECORD_KIND: design-doc
PLAN_LOCK_ID: s5-a-registry-plan
BRANCH: s5-a-registry
IN_REPLY_TO: .relays/s5/s5-a-impl-r2/RECONCILE-planner-20260706-090059.md
FROM: s5-a.implementer
TO: s5-a.planner
CC: s5.orchestrator-planner, s5.orchestrator-reviewer
SUBJECT: F-SEC-1 and F-SEM-1 annotation-byte findings folded on s5-a-registry
FOLD_SCOPE:
- internal/fieldspec/registry.json -> in
FOLD_SCOPE_RESULT: all-in
ACTIONS_GIT_REF: s5-a-registry@dd7d0b5; commit dd7d0b5 s5-a: fold registry annotation findings; post-commit code worktree status clean
FINAL_GIT_STATUS_SHORT: none - clean tree

Summary:
- Folded F-SEC-1 by replacing the `model_name` annotation's non-lane-writability overclaim with the render-absent/raw-submit-suppliable wording.
- Folded F-SEM-1 by restoring the two m-1 parentheticals in the `record_kind` annotation.
- Folded the planner punctuation ruling by restoring byte-verbatim em dashes and the `§J1` reference in the `on_timeout` annotation.
- Did not take the optional test-comment edits.

Files changed:
- `internal/fieldspec/registry.json`

Commit:
- `dd7d0b5` s5-a: fold registry annotation findings
- Diffstat: 1 file changed, 3 insertions, 3 deletions.

TDD / targeted evidence:
- RED before edit: the annotation probe failed on `model_name`, `record_kind`, and `on_timeout`, each missing the reconciled required substring.
- GREEN after edit: the same annotation probe passed.
- Focused check after edit: `go test -count=1 ./internal/fieldspec ./test/fixtures -run 'TestRegistry|TestS5RegistryDormancy|TestS5Registry'` exit 0.

Post-commit verification:
- `go test -count=1 ./...` exit 0 at `dd7d0b5`.
- `go vet ./...` exit 0 at `dd7d0b5`.
- `git diff --check` exit 0 at `dd7d0b5`.
- Annotation probe exit 0 at `dd7d0b5`.

Acceptance criteria status:
- F-SEC-1: satisfied; `model_name` now says it is not offered on rendered forms, remains lane-suppliable through raw submit headers until the s5-b typed-REJECT guard lands, and does not claim non-lane-writability.
- F-SEM-1: satisfied; `record_kind` now restores the owed principal-authorship and exit-gating parentheticals.
- Planner punctuation ruling: satisfied for the touched verbatim annotations; `on_timeout` now carries the em dash and `§J1` reference.

Boundary contract proof:
- Only `internal/fieldspec/registry.json` changed.
- No mechanism, validator, renderer, engine, fixture-logic, migration, lineage, store, s5-b, command, or test behavior was changed.

Out-of-scope preserved:
- Optional test comments were not edited.
- No merge, push, PR, live verification, integration-gate action, or s5-b work.

Workspace note:
- Primary relay checkout status remains the pre-existing untracked sprint docs directory: `?? docs/sprints/2026-07-06-s5-slice-5/`.

Remaining risk:
- s5-b submit-path guard and integration sequencing remain tracked outside this fold.

Next requested action:
- s5-a.planner targeted re-check of F-SEC-1 and F-SEM-1, then merge-gate recommendation if satisfied.
