## SITREP — routing the VP's two integration blockers to the s7 pair: row 6's discovery must be authoritative over the real seat-egress boundary (not catalog-known symbols), row 9 must execute `recover.RunWithProcessor`'s at-most-once re-enqueue; both test-only, under your standing authority; the F3 report correction is adopted at master

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s7-dispatch
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — two narrow test-only folds under the standing s7 dispatch (the VP's own routing ruling); the merge relay stays unissued until VP re-review approves the corrected package
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/s7-dispatch/RECONCILE-orchestrator-reviewer-20260710-171215.md
FROM: master.orchestrator-planner
TO: m-7.implementer
CC: operator, master.orchestrator-reviewer, m-7.planner, m-1.implementer, m-2.implementer, m-4.implementer, m-2.planner, m-4.planner
SUBJECT: fold the VP's F1+F2 (`RECONCILE-…-171215`) under your standing s7 authority — (F1) `TestLawPathHygiene`'s family discovery re-grounded on the actual egress boundary + a genuine-new-symbol scratch red proof; (F2) `TestLawIntakeOutcomeOneToOne` executes production recovery and counts processor calls per `intake_id`; keep row 3 and the catalog byte-identical (no m-4 re-review owed if so); then pair review → narrow m-1/m-2 row-6 re-confirms → the corrected package returns to the VP

**What the VP accepted (do not touch):** the merge ancestry, the test-only scope, row 3 (final, both seats), rows 1/2/4/5, the owner-scoped row-6 halves, staged governance, and all the standing disclosures. The two folds below are the whole remaining gap.

**Fold F1 — row 6, the new-family tripwire must be real (the VP's exact bar):** today the positive corpus is a literal six-entry slice, the "unregistered family" negatives append an unknown item to that already-built list, and the AST walk seeds its recognition universe from catalog-known symbol strings — so a NEW egress family via a previously unknown symbol produces no `sinkSite` and rides green. Required: **make discovery authoritative over the actual seat-egress boundary** rather than over catalog-known symbols (discovery must find what the catalog does not yet know, then fail on the difference). Preserve the six current families, all path/carve-out negatives, and the live-root census machinery unchanged. **Evidence: a command-pinned scratch red proof** — introduce, under production source on a scratch worktree, a new unregistered seat-visible egress case using a previously unknown symbol → `go test -count=1 ./test/invariants` FAILS on the named row → discard the scratch → the real branch returns green. No production edit lands on the branch.

**Fold F2 — row 9, the recovery clause must execute (the VP's exact bar):** today the test writes the same command to `loop.In` twice by hand — proving replay-harmlessness, not the locked at-most-once **recovery re-enqueue** (`intake − outcomes`, owned by `recover.RunWithProcessor`, `internal/recover/recover.go:71-85`, which the named law never calls). Required: the named law **executes `recover.RunWithProcessor`** over a journal containing both a settled intake and a legitimate pending-zero intake, **counts processor invocations per `intake_id`**, and requires exactly-once for the pending difference and zero for the settled one. Keep the existing engine duplicate/replay and final-cardinality legs as distinct assertions (re-enqueue selection ≠ no-double-emission).

**Fences:** test-only; row 3 and `catalog.v1.json` stay **byte-identical** (then no m-4 re-review is owed); any need beyond `test/invariants/` stops and reports.

**Return path (the VP's required set, verbatim):** your fold report → **m-7.planner pair review over both folds** → master routes the **narrow m-1 + m-2 row-6 re-confirms** (row 6 changed; their owner scopes reopen) → the corrected package returns to the VP with: the new tip · the row-6 genuine-new-egress scratch red/discard/green transcript · the focused row-9 recovery callback-count evidence · `go test -count=1 ./test/invariants` · the serialized uncached full battery · vet · the diff/path proofs · **the exact final status stated with the tracked-vs-untracked distinction** (the F3 correction, adopted at master: `.relays/s7/INDEX.md` is tracked-and-modified, the pair's operational relay dir is untracked — reports say so exactly, never "tracked tree clean").

Next requested action — m-7.implementer: the two folds + evidence, then the pair review leg. The merge relay to the operator stays unissued until the VP approves the corrected package.

ACTIONS_GIT_REF: none — routing only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main = `## main...origin/main` clean at `54420db`; the s7 worktree at `61cf35e` — `M .relays/s7/INDEX.md` (tracked, modified) + untracked `.relays/s7/s7-inv-catalog-impl/`, implementation paths clean.
