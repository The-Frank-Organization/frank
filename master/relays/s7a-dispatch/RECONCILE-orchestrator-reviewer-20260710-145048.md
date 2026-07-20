## RECONCILE -- s7a integration review: revise before merge; v5 registry provenance still attributes the artifact to s6

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s7a-dispatch
PARENT_DISPATCH_ID: s7a-dispatch
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes -- operator merge remains a separate downstream gate after this finding is folded and VP re-approval lands; no merge now
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-2.planner, m-2.implementer, m-4.implementer, m-7.implementer
IN_REPLY_TO: master/relays/s7a-dispatch/RECONCILE-orchestrator-planner-20260710-144333.md
SUBJECT: revise -- predicate guard, authority chain, fidelity, red/green, scope, and E2 battery are sound; block merge until the s7a-fieldspec-v5 in-member provenance is truthful and pinned by test

VERDICT: revise

## Blocking Finding

1. **`s7a-fieldspec-v5` carries stale s6 change attribution.** The branch advances `registry.json.version` to `s7a-fieldspec-v5`, but the adjacent provenance block still says `owner: s6-core-registry`, `design_doc_id: s6-slice-6-design`, `plan_lock_id: s6-slice-6-plan`, and `note: s6 registry pass...` (`internal/fieldspec/registry.json:2-7`). The locked registry contract makes the in-member **version plus minimal provenance block** the change-attribution record; the top-level digest proves byte identity but is intentionally mute about attribution (`s3-consult-m7/SITREP-planner-20260704-171546.md:26-31`). Prior version advances updated both together (`s3-fieldspec-v2`, `s5-fieldspec-v3`, `s6-fieldspec-v4`). This v5 artifact therefore has intact integrity and false attribution.

   The current test does not catch the mismatch: it asserts the v5 version exactly but checks only that `Provenance["owner"]` is nonempty (`internal/fieldspec/registry_test.go:18-23`), so stale s6 metadata passes. A governance registry cannot merge with its own audit stamp pointing at the superseded plan.

   **Required narrow fold:** add a failing exact assertion for the v5 provenance tuple, then update the four provenance values to truthful s7a attribution (owner, governing/audit record, plan lock, and note). Amend AC6's "only shipped registry data delta" wording to distinguish the one semantic row delta from this required metadata attribution update. Keep the guard, singleton, version, and every other registry byte unchanged. The fold remains inside `registry.json` + `registry_test.go`, already within the five-file fence.

   After the fold: m-2 pair reviews the exact delta; m-7 narrowly re-confirms the composite-digest/s8-genesis condition against the **final bytes**. The m-4 semantic confirm may stand if the diff is provenance plus its exact assertion only. Master reruns the focused tests, full uncached battery, vet, diff-check, and five-file scope check, then returns the package here.

## Merge-Authority Correction

2. **A VP approval is not an operator merge request, and `CC: operator` carries no action authority.** The incoming relay is addressed only to this reviewer but says the operator merge gate is requested "via the operator CC" (`RECONCILE:11,15-16`). CC is visibility only. After the code finding is folded and VP approval lands, route a separate decision relay `TO: operator`; if a non-operator seat performs the merge, it must receive the valid addressed merge authorization. Do not treat a later VP approve as auto-merge authority.

## Accepted Package

- The corrected pair-first authority chain, unique dispatch IDs, approving PLAN-REVIEW, all-in five-file scope diff, and pair-issued implementation token are valid.
- The default-deny column guard is correctly called after the parent-array gate check. The shipped allowlist is exactly `routing_assignments -> [declared_deviated]`; `chosen_model` and non-allowlisted `seat` reject at registry load while the legal deviation atom compiles.
- Red-first lineage is real: detached `10ee3a2` fails all three new rejection cases because load succeeds; `d76c3ad` passes them. The green commit contains guard, member, singleton, version, positive retarget, and assertion together, preserving AC6 atomicity.
- m-4's current-grammar R2/C1 confirm and representational Step-3 residue are sound. m-7's fail-before-serve, path-free error, A-1 distinction, and section-7/composite-digest analysis are sound subject to recomputation over the final provenance-corrected bytes.
- `FLAKE-SOCKET-PAR` is non-blocking for this diff and remains correctly registered for s8. My parallel full-suite run passed; one pass does not discharge the recorded intermittent issue.

## Independent Verification

- Incoming exact-file lint -> OK; pre-filing `--relay-root master/relays/s7a-dispatch` -> OK. Both fidelity relay roots -> OK.
- Branch/worktree: `s7a-colgrain@d76c3ad660614f58d72f107330becd465b6a1eaa`; clean; two commits over `1d3e92c` (`10ee3a2` red, `d76c3ad` green).
- Diff -> exactly the five authorized fieldspec files; `git diff --check 1d3e92c..d76c3ad` -> clean.
- Fresh green at the branch tip: focused four-test fieldspec command -> PASS; `go test -count=1 ./...` -> 24 packages `ok`, 2 `[no test files]`, zero failures; `go vet ./...` -> exit 0.
- Independent red in detached worktree at `10ee3a2`: `TestRegistryLoadRejectsNonAllowlistedRowColumns` -> FAIL as intended, all three subtests reporting load succeeded when rejection was expected; temporary worktree removed afterward.
- Source review: registry-load validation runs before channel serve; error text is logical/path-free; no engine/store/seat-surface change.

Next requested action: route the two-file provenance fold through m-2 pair review, obtain the narrow m-7 final-bytes re-confirm, rerun master verification, and return the revised s7a package. No merge is authorized.

ACTIONS_GIT_REF: wrote this reviewer relay and appended `master/relays/INDEX.md`; branch/worktrees were reviewed read-only; the temporary detached red-proof worktree was created and removed; no source edit, commit, merge, or push.
FINAL_GIT_STATUS_SHORT: `frank/` main = `## main...origin/main` at `1d3e92c`; `s7a-colgrain` clean at `d76c3ad`; cwd is not a Git repository.
