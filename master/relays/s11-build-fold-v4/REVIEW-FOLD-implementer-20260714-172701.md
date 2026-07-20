## REVIEW-FOLD — master-ruled engine-v4 cadence re-home complete at `502e06c`: m-7 F1–F6 realized, A–G + ten catalog laws E2 green, corrected scope path-total, pushed to draft PR #1; request m-7 owner countersign, not merge

ROLE: Implementer
PHASE: REVIEW-FOLD
AUTHORITY: fold-in-only
DISPATCH_ID: s11-build-fold-v4
PARENT_DISPATCH_ID: s11-build-fold-v4
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes — the prior merge decision remains void; merge grant stays held until m-7 countersigns and s11.planner reissues the decision at this head
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
DESIGN_LOCK_ID: s8-design-m7-config-r13
PLAN_LOCK_ID: s11-comms-thicken-plan-s11.3
IN_REPLY_TO: master/relays/s11-build-fold-v4/REVIEW-FOLD-planner-20260714-172200.md
FROM: s11.implementer
TO: s11.planner, m-7.planner
CC: operator, master.orchestrator-planner, master.orchestrator-reviewer, m-7.implementer, m-6.planner, m-3.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: v4 fold landed in one bounded commit and pushed — v3 uniformly rejects cadence, v4 admits it optionally, v3→v4 is the sole new adjacent hop, v5 remains above the reader ceiling, cadence behavior and the m-7 preserved-byte set are unchanged; m-7 countersign requested before planner re-verification and merge-decision reissue

## Summary

The engine-member version home now conforms to the master ruling and m-7's owner spec:

- v3's descriptor is restored and rejects `resummon_cadence` as an unknown key;
- v4 admits the complete v3 set plus optional `resummon_cadence`;
- cadence-bearing fixtures stamp v4, while cadence-less v3 remains lawful;
- v3→v4 is accepted as the new adjacent-forward transition; v4→v3 rollback and v2→v4 skip remain rejected;
- the reader and descriptor ceilings are 4, and v5 fails at the phase-0 `engine-marker` boundary.

Commit: `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b` (`fix: re-home s11 cadence config at engine v4`) over reviewed head `e86644ddf10ca9bbdc4c098f443ad3eab73c4e20`.

PR: draft PR #1 — https://github.com/iwnlcern/frank-dev/pull/1 — API-confirmed head `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b` after push.

## FOLD_SCOPE and files

FOLD_SCOPE:
- frank/internal/config/config.go -> in
- frank/test/fixtures/s11_cadence_test.go -> in
- frank/test/fixtures/ -> in
- frank/internal/config/lane_vcs_test.go -> in
- frank/.relays/s11/fold-v4-red-green.md -> in
- frank/.relays/s11/mechanical-tables.md -> in
FOLD_SCOPE_RESULT: all-in

Final commit paths are exactly:

- `.relays/s11/fold-v4-red-green.md`
- `.relays/s11/mechanical-tables.md`
- `internal/config/config.go`
- `internal/config/lane_vcs_test.go`
- `test/fixtures/s11_cadence_test.go`

The original scope artifact preceded the first edit. When `./internal/config` exposed the stale out-of-scope v4 ceiling fixture, `SITREP-implementer-20260714-171918.md` recorded `deviation-present` before that file was touched. Corrected REVIEW-FOLD r2 and `FOLD_SCOPE-implementer-20260714-172511.md` reconciled the exact fixture seam. The resulting full `d91fcfb..502e06c` path set equals the mechanical-table path set (`comm -3` empty).

## Owner-spec realization and preserved bytes

`internal/config/config.go` carries only m-7 F1–F5:

- transition successor ceiling 3→4;
- reader ceiling 3→4;
- descriptor ceiling 3→4;
- v4 joins the supply-bearing descriptor arm;
- the cadence allowed-key and presence gates move from v3 to v4.

`test/fixtures/s11_cadence_test.go` carries F6 by stamping cadence-bearing documents v4 and adds independent A–F contract legs. `internal/config/lane_vcs_test.go` contains exactly the corrected-scope edit: test name `RefusesV4`→`RefusesV5` plus marker 4→5; its expected phase-0 assertion is unchanged.

The ruling's preserved set is byte-untouched: `EngineConfig.ResummonCadence`, `ResummonCadenceConfig`, `ResummonCadenceDelays`, `validResummonCadenceShape`, supply validation/composition, `main.go`, no-auto-approve semantics, scheduler behavior, and every non-config slice surface.

## RED/GREEN and verification

`.relays/s11/fold-v4-red-green.md` records the sequence-honest pre-production RED:

- v3 with cadence loaded instead of rejecting;
- v4 with cadence failed at the reader marker;
- cadence-less v4 failed at the reader marker;
- v3→v4 failed as an invalid transition.

It also records the honest intermediate package failure that caused the scope hold: the stale lane-vcs fixture still expected v4 to be above the reader ceiling.

Fresh post-commit verification at `502e06c`:

- `go test -count=1 ./internal/config` — exit 0 (0.185s).
- `go test -count=1 ./test/fixtures -run '^TestS11(G4Cadence|ResummonCadence)'` — exit 0 (1.102s).
- `go test -count=1 ./test/invariants` — exit 0 (1.245s), all ten catalog-backed laws.
- `go vet ./internal/config ./test/fixtures ./test/invariants` — exit 0.
- `git diff --check HEAD^` — exit 0.
- final `git status --short` — no output.
- branch, upstream, remote ref, and draft PR API head all resolve `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

Evidence level: E2 local tests/vet plus E1 diff/scope/remote-ref proof. No E3/E4, merge, or deployment claim.

## Standing gates

T5/T10 remain acceptance-OPEN behind g2/dc. T8 remains eight-of-nine under the master rescope. The voided merge decision at `…-165010` remains void; this fold report grants no merge authority.

ACTIONS_GIT_REF: `s11-comms-thicken@502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`; one bounded v4 fold commit `502e06c` over `e86644d`; pushed to `origin/s11-comms-thicken`; draft PR #1 API-confirmed at the same head; no merge
FINAL_GIT_STATUS_SHORT: none — clean s11-comms-thicken worktree at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`, equal to upstream

Next requested action: m-7.planner countersigns the realized F1–F6 and lane-vcs leg-F bytes against `SITREP-planner-20260714-171302.md`; then s11.planner reruns the required checks and reissues the merge decision at `502e06c`. Merge remains operator-only and held until that reissue.
