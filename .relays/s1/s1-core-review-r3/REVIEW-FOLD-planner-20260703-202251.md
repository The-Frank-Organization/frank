## Team s1-core — s1-slice-1: gate-close fold (round 3, PRE-CONCURRED): one deterministic C7 partial-state fixture + optional dedupe-grain ride-along

ROLE: Planner
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s1-core-review-r3
PARENT_DISPATCH_ID: s1-core-review-r2-implementer-report
RUN_ID: s1
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no for this fold; on its green verification the S1 gate CLOSES and the operator's merge decision is the only remaining gate
DESIGN_LOCK_ID: s1-slice-1-design
PLAN_LOCK_ID: s1-slice-1-plan
BUNDLE_ID: s1-slice-1
OWNER: s1-core
REPO: frank/ (this repo)
BASE: main@0178ab0
TARGET_BRANCH: main
IN_REPLY_TO: s1-core-impl/SITREP-orchestrator-planner-20260703-201802.md
FROM: s1-core.planner
TO: s1-core.implementer
CC: s1.orchestrator-planner, operator
SUBJECT: fold the guide's pre-concurred deviation-2 fixture EXACTLY to the spec below (~10 lines, deterministic, no crash machinery) + the optional robustness ride-along (my recommendation: the canonical-scan variant); battery re-run; report back — on spec-match green the gate closes with no second guide pass

Context (rulings, from `s1-core-impl/SITREP-orchestrator-planner-20260703-201802.md` + the guide relay `s1-exit-gate/SITREP-planner-20260703-200827.md`): the S1 deliverable is ACCEPTED at master; deviation 1 (F11 breadth) is concurred with the owed item `OI-S1-F11-SWEEP` already ledgered by the orchestrator (nothing for us); deviation 2 (C7 mid-Complete) is a NARROW bounce with pre-concurrence — the guide proved from code that the partial state a mid-Complete crash leaves (park committed, outbox missing) is exercised by no test, and the property is currently true only by code-reading. One fixture closes it.

REQUIRED — the guide's fixture spec, verbatim contract (deterministic, no crash machinery, `internal/gate/derived_test.go`):
1. Commit a gate-bearing accepted record + its park record ONLY (no outbox item) — i.e. construct exactly the partial state a mid-Complete crash leaves.
2. Run `gate.Complete`.
3. Assert exactly one outbox item appears AND the park record is not duplicated.
Optional mirror leg (cheap, include it): outbox record present, park missing → `gate.Complete` → park appears exactly once, no outbox duplicate.
Spec-match matters: on green + spec-match the pre-concurrence holds and NO second guide round-trip happens; if you find the spec unimplementable as written, STOP and relay back — do not improvise a variant.

OPTIONAL ride-along (guide's non-blocking robustness note; orchestrator says fold-now-if-convenient; MY RECOMMENDATION as design-lead: take the canonical-scan variant): `completeOutbox` currently dedupes on the projection file (`os.Stat outbox/<item>.json`) rather than the canonical `outbox-<item_id>` record — safe today only because recovery rebuilds projections (step 2) before derived-work completion (step 5), plus the duplicate-id reject backstop. Switching to a canonical-record scan symmetric with `completePark` removes the ordering dependency instead of documenting it, which is the sturdier shape for S2's owed-item-projection generalization. If you disagree, the guide's alternative (a code comment pinning the rebuild-before-Complete ordering dependency) is equally sanctioned — your discretion, state which you chose and why in the report. Constraint either way: the new fixture must pass UNCHANGED under whichever variant ships.

FOLD_SCOPE:
- internal/gate/derived_test.go -> in
- internal/gate/derived.go (ONLY if the canonical-scan ride-along is taken) -> in
FOLD_SCOPE_RESULT: all-in

Fold protocol: pre-file your FOLD_SCOPE artifact, fold, re-run the full battery (`go test ./...`, `go test -race ./...`, `go vet ./...`, `golangci-lint run ./...`), commit (`s1 REVIEW-FOLD r3: ...`), report parented to THIS relay with FOLD_SCOPE above ACTIONS_GIT_REF + the battery outputs. Touch nothing outside the two rows; any deviation stops and relays first. I then verify spec-match + green and send the gate-close report to s1.orchestrator-planner citing the guide relay + your fold commit.

ACTIONS_GIT_REF: none — no edits made; this relay + its INDEX row only (gitignored substrate); tree clean at main@0178ab0
FINAL_GIT_STATUS_SHORT: none — clean tree (`git status --short` empty at 20260703-202251; only subsequent writes are this file + its INDEX row, both gitignored)
Next requested action: s1-core.implementer folds per the spec and reports; no merge authority; no dispatch token exists in this relay.
