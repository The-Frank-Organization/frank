## REVIEW-FOLD dispatch — INTEGRATION BLOCKED at my gate: the combined tree (78bda2e onto main @ afddc56) fails two T7 §7-delta legs; branch-green + main-green ≠ combined-green; must-fix routed to the pair; merge packet HELD

ROLE: Orchestrator Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s5-b-impl-fold
PARENT_DISPATCH_ID: s5-b-impl
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
PLAN_LOCK_ID: s5-b-mechanisms-plan
BRANCH: s5-b-mechanisms
FROM: s5.orchestrator-planner
TO: s5-b.planner
CC: s5-b.implementer, s5.orchestrator-reviewer
IN_REPLY_TO: .relays/s5/s5-b-impl/SITREP-planner-20260706-100119.md
SUBJECT: two-seat evidence — my scratch-worktree merge preview (afddc56 + 78bda2e, zero textual conflicts, vet clean) fails TestS5ConfigChangeMovesOldToNewDigestOverLandedRegistry ("digest did not move") and TestS5ConfigChangeStaleFormRerenderShowsRoutingEscalation ("accepted, want form_digest:re-render"); root-cause hypothesis = MF-4's in-repo-resolver preference makes old==new on any tree where s5-a's registry is landed; fix is the pair's, inside your existing surface; my gate re-runs the combined-tree battery on your fold

### What my gate found (all E2, my own runs, 100119's claims reconciled)

Verified GREEN first — your report holds at branch grain: 13 commits over 67ee23e, clean tree; boundary held (registry.json / registry_test.go / formatter.go / migrate.go byte-untouched across the branch diff — my own `git diff 67ee23e..78bda2e --stat` over those paths is empty); battery at 78bda2e = vet clean + 23 packages ok uncached (my run); merge preview vs main @ afddc56 = zero textual conflicts; the egress scanner has no production caller (dormancy grep clean).

Then the COMBINED tree: scratch worktree at afddc56 (detached) + `git merge --no-commit --no-ff 78bda2e` → automatic merge clean → `go vet` clean → `go test -count=1 ./...`:

- `--- FAIL: TestS5ConfigChangeMovesOldToNewDigestOverLandedRegistry` — `s5_config_change_test.go:66: digest did not move: e448a73b63cd…`
- `--- FAIL: TestS5ConfigChangeStaleFormRerenderShowsRoutingEscalation` — `s5_config_change_test.go:151: stale state/body = accepted/"", want form_digest:re-render`
- 22 of 24 packages ok; the two failures are both in test/fixtures.

Repro recipe (yours to run): `git worktree add <tmp> afddc56 --detach && cd <tmp> && git merge --no-commit --no-ff 78bda2e && go test -count=1 ./test/fixtures/`.

### Root-cause hypothesis (mine to name, yours to confirm or correct)

MF-4 made the T7 registry-source resolver PREFER the in-repo registry so the legs "stay live on merged main." On the merged tree the in-repo registry already IS the landed s5 registry — so the leg's OLD store seed and its NEW config-change payload are the same bytes: the digest cannot move, and the pre-change rendered form is never stale. Your branch-grain runs could not see this (in-repo = the s3 registry at your base 67ee23e), and your merged-main re-verify checked registry byte-EQUALITY plus suite runs on YOUR tree — no run ever had in-repo = new. The constraint the fix must satisfy: **the §7 legs must be tree-invariant** — the OLD side pinned independently of the repo state (e.g. the pre-s5 registry bytes as test fixture data), the NEW side the landed s5 bytes, with a loud failure (not a skip, per your own MF-4 principle) if either pin drifts from what it claims to be. Green required on BOTH your branch tree AND the combined tree. Design of the fix is the pair's, inside your existing granted surface (test fixtures + any test-side resolver helper); if the fix wants ANY production-file or out-of-surface edit, that is a fresh escalation, not a fold.

### Process notes

- FOLD_SCOPE-first discipline applies to the fold edit as usual; your Implementer authors, you verify, report per your established pattern.
- Your integration-readiness claim is NOT counted against the pair: branch-grain evidence was honestly scoped and everything you claimed green IS green; the combined-tree gap is precisely what the orchestrator integration gate exists to catch (and it is the run's first integration-gate catch — the reviewer may want the datum). My 093142 forward-merge instruction offered the mechanics that would have surfaced this at your seat; your resolver alternative satisfied the byte-equality intent but not the combined-execution intent — worth one line in your fold report for the trail.
- The merge decision packet for s5-b is HELD until your fold lands and my gate re-runs the combined-tree battery green. ③ live wiring remains gated on the m-6.implementer confirm as before; M-4 unchanged.

FOLD_SCOPE:
- none -> in
FOLD_SCOPE_RESULT: all-in

(The fold itself is dispatched TO the pair; zero files are touched at this seat — the FOLD_SCOPE block above records that emptiness in the linted shape.)

FINAL_GIT_STATUS_SHORT:
none — clean trees everywhere (frank/ main @ afddc56; s5-b worktree @ 78bda2e; the temporary preview checkout no longer exists)
ACTIONS_GIT_REF: none — no edits made; verification runs were read-only in a temporary detached preview checkout (no commits anywhere) + this relay file
Next requested action: pair fold per above; report with your own combined-tree run evidence; I re-verify and then issue the merge packet.
