## SITREP — v4 fold scope deviation discovered by the required config check: the pre-existing reader-ceiling fixture still names v4, so `internal/config/lane_vcs_test.go` must re-pin v4→v5 but is OUT of the binding FOLD_SCOPE; fold held before that edit, commit, or push

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s11-build-fold-v4-deviation-lane-vcs-test
PARENT_DISPATCH_ID: s11-build-fold-v4
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — scope correction only; merge remains held for the reissued post-fold decision
GRILL_REQUIRED: no
DESIGN_DOC_ID: c3-design-m-6-human-surface-scheduler
DESIGN_LOCK_ID: s8-design-m7-config-r13
PLAN_LOCK_ID: s11-comms-thicken-plan-s11.3
IN_REPLY_TO: master/relays/s11-build-fold-v4/REVIEW-FOLD-planner-20260714-171200.md
FROM: s11.implementer
TO: s11.planner, m-7.planner, master.orchestrator-planner
CC: operator, master.orchestrator-reviewer, m-7.implementer, m-6.planner, m-3.planner
BUNDLE_ID: m-6-human-surface-scheduler
SUBJECT: required scope correction — add exactly internal/config/lane_vcs_test.go so its reader-ceiling fixture moves from v4 to v5; the ruled v4 bytes and new A-F fixture are otherwise green, but the mandated ./internal/config target cannot pass while the stale out-of-scope assertion remains

## Mechanical scope result

FOLD_SCOPE:
- frank/internal/config/config.go -> in
- frank/test/fixtures/s11_cadence_test.go -> in
- frank/test/fixtures/ -> in
- frank/.relays/s11/fold-v4-red-green.md -> in
- frank/.relays/s11/mechanical-tables.md -> in
- frank/internal/config/lane_vcs_test.go -> OUT
FOLD_SCOPE_RESULT: deviation-present

No edit was made to the OUT row.

## Evidence

The sequence-honest A–F RED was observed before production edits:

- v3 carrying cadence incorrectly loaded successfully;
- v4 with cadence failed at `engine-marker`;
- cadence-less v4 failed at `engine-marker`;
- v3→v4 failed with `config-version-transition`.

After carrying m-7's F1–F5 production anchors and F6 test stamp inside the authorized scope:

- `go test -count=1 ./test/fixtures -run '^TestS11ResummonCadenceUsesEngineV4Contract$'` — exit 0;
- `go test -count=1 ./test/fixtures -run '^TestS11G4Cadence'` — exit 0;
- `go test -count=1 ./internal/config` — exit 1 only at `TestLaneVCSReaderCeilingRefusesV4BeforeSupply`: observed `config-load: engine`, while the stale fixture expects phase-0 `engine-marker` rejection for v4.

That failure is correct under the ruled bytes: v4 is now a supported reader marker, so invalid v4 supply reaches member validation. The ceiling fixture must become `TestLaneVCSReaderCeilingRefusesV5BeforeSupply` with its document marker changed from 4 to 5; the expected phase-0 `engine-marker` result remains unchanged. Making production reject v4 at preflight to preserve the stale test would directly violate the master ruling and the owner F1–F6 spec.

Current authorized working diff is limited to `internal/config/config.go` and `test/fixtures/s11_cadence_test.go`; no evidence or mechanical-table file has yet been written, no commit or push occurred, and the existing PR remains at `e86644d`.

ACTIONS_GIT_REF: uncommitted authorized REVIEW-FOLD diff in `internal/config/config.go` and `test/fixtures/s11_cadence_test.go` over `s11-comms-thicken@e86644ddf10ca9bbdc4c098f443ad3eab73c4e20`; no OUT-row edit, commit, push, PR-head change, or merge
FINAL_GIT_STATUS_SHORT:
 M internal/config/config.go
 M test/fixtures/s11_cadence_test.go

Next requested action: issue a corrected directly addressed REVIEW-FOLD scope adding only `frank/internal/config/lane_vcs_test.go` for the v4→v5 reader-ceiling re-pin. On that correction, s11.implementer completes the evidence/mechanical files, targeted config + cadence + ten-law checks, commit, push, and fold report for m-7 countersign. Merge remains held.
