## PLAN-REVIEW -- s4-slice-4-plan current r3 review

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s4-wire-plan-review-implementer-r3
PARENT_DISPATCH_ID: s4-wire-plan-lock
RUN_ID: s4
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: yes - SCOPE_DIFF, delegated dispatch, live-seat designation, section-7 authorization, E3 gate, and s4-close remain
DESIGN_LOCK_ID: s4-slice-4-design
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s4-grill-s4-wire
PLAN_LOCK_ID: s4-slice-4-plan
PLAN_REVIEW_VERDICT: approve
FROM: s4-wire.implementer
TO: s4-wire.planner
CC: s4.orchestrator-planner, operator
IN_REPLY_TO: s4-wire-plan/PLAN-planner-20260705-124044.md
BASE: main@28dfa33
SUBJECT: PLAN-REVIEW verdict -- approve current r3 plan body; m-1 carry and README ruling fold verified; dispatch remains gated

## Phase Discipline

This is a read-only PLAN-REVIEW response. I made no source, test, design, plan, or sprint-spec edits. I did not start implementation; no live implementation dispatch token is present in the reviewed relays.

## Reviewed

- Operator-routed trigger: `.relays/s4/s4-wire-plan/SITREP-orchestrator-planner-20260705-043600.md`.
- Parent PLAN request: `.relays/s4/s4-wire-plan/PLAN-planner-20260705-124044.md`.
- Binding m-1 verdict: `.relays/s4/s4-fidelity-m1/SITREP-implementer-20260705-042308.md`.
- Binding m-1 supplement: `.relays/s4/s4-wire-plan/SITREP-orchestrator-planner-20260705-042636.md`.
- README fence ruling: `.relays/s4/s4-wire-plan-fence-ask/SITREP-orchestrator-planner-20260705-042636.md`.
- Plan-currentness context: `.relays/s4/s4-wire-plan/SITREP-planner-20260705-043557.md`.
- Plan doc: `docs/sprints/2026-07-05-s4-slice-4/plans/s4-slice-4-plan.md`; current plan content last changed at `main@e29e0b2`, live `HEAD/main` is `28dfa33` with only a later RECONCILE append.
- Locked design doc: `docs/sprints/2026-07-05-s4-slice-4/designs/s4-slice-4-design.md` r3 at `main@2ef9437`, with approving review `.relays/s4/s4-wire-design/DESIGN-REVIEW-implementer-20260705-040925.md`.

## Verdict

`PLAN_REVIEW_VERDICT: approve`

This approval is scoped to the current plan body for `PLAN_LOCK_ID: s4-slice-4-plan`, parented to `s4-wire-plan-lock`. It approves the plan gate only. It does not authorize implementation, merge, SCOPE_DIFF deviation, an OUT touch, design amendment, or any section-7 live authorization.

## Review Object Currentness

The parent PLAN relay is the only non-superseded pair PLAN relay and is addressed to this seat (`PLAN-planner-20260705-124044.md:6-20`). It locks the current r3 plan object directly at `main@e29e0b2` and supersedes the two earlier unconsumed PLAN lock relays, `042444` r1 and `043152` r2 (`PLAN-planner-20260705-124044.md:22-37`). The planner's hygiene SITREP had already recorded the pre-review r3 advance and the lack of any prior PLAN-REVIEW consumption (`SITREP-planner-20260705-043557.md:20-30`); the 124044 relay makes that review object unambiguous.

I verified that delta: `git diff d0ea59b..e29e0b2 -- docs/sprints/2026-07-05-s4-slice-4/plans/s4-slice-4-plan.md` changes only the README row from PENDING-RULING to RULED-IN, cites `s4-wire-plan-fence-ask/SITREP-orchestrator-planner-20260705-042636.md`, and restates the three conditions in Task 11. Current `28dfa33` adds only `RECONCILE.md`; the plan body is unchanged after `e29e0b2`.

Result: I am approving the plan r3 body locked by `PLAN-planner-20260705-124044.md`, not the superseded r1/r2 PLAN lock relays or the stale pending-ruling text.

## Checks

### 1. Design lock and boundary contract

The PLAN relay carries the expected design lineage: `DESIGN_LOCK_ID: s4-slice-4-design`, `DESIGN_RECORD_KIND: design-doc`, `GRILL_LOCK_ID: s4-grill-s4-wire`, and parent `s4-wire-design-r3-review-implementer` (`PLAN-planner-20260705-124044.md:6-17`, `:39-40`). The plan points at the approved design r3 (`s4-slice-4-plan.md:7`) and keeps the design's exact seat-facing verb set: `submit`/`project`/`read` (`s4-slice-4-plan.md:13-18`).

The OUT fence is preserved: no s5 consumer content, Step-2 observe, Step-3 routing, TUI, federation, external send, authority replacement, in-band rotation/supersede, socket-dialect rewrite, binding-table shape change, frozen s2-store upgrade, or analytics work enters scope (`s4-slice-4-plan.md:155-161`).

### 2. m-1 verbatim carry and route-back triggers

The m-1 verdict requires F-S4-M1-1 through F-S4-M1-6 to appear in the gated PLAN and makes the config-read/projection redaction rule executable in fixtures (`SITREP-implementer-20260705-042308.md:18-22`, `:36-117`). The plan carries all six conditions in the global constraint block and names the same route-back triggers (`s4-slice-4-plan.md:18`).

The task mapping is executable:

- F-S4-M1-1 maps to Task 4, including operator-channel provenance, non-operator typed refusal, operator-scoped registry home, and the stop-and-route trigger if ordinary lane authority would broaden (`s4-slice-4-plan.md:77-86`).
- F-S4-M1-2 maps to new Task 4b, including exact redacted non-operator read view, operator full read, no member bytes in projections/nudges/errors, and digest-as-claim fixture (`s4-slice-4-plan.md:88-95`).
- F-S4-M1-3 maps to Task 2, including SHA-256 in-memory active index, `auth:channel-active`, kernel-close recovery, no persistence/output/logging, and no binding-table change (`s4-slice-4-plan.md:55-64`).
- F-S4-M1-4 maps to Task 7 and Task 11 custody/docs surfaces (`s4-slice-4-plan.md:111-121`, `:133-137`).
- F-S4-M1-5 maps to Task 3, with committed-record recipients, stamped auth metadata, recipient-only nudges, and no cross-seat metadata (`s4-slice-4-plan.md:66-75`).
- F-S4-M1-6 maps to Tasks 4 and 5, with fresh `store.Init` untouched, one canonical pivot, committed-records-only chain, no second root/journal/hidden state, and fail-closed serving reads (`s4-slice-4-plan.md:77-86`, `:97-104`).

The acceptance map repeats the same mapping and makes the 4b redaction fixtures gate acceptance, not optional hardening (`s4-slice-4-plan.md:151-153`). This satisfies m-1's dispatch-condition paragraph.

### 3. Task 4b redacted-view shape

m-1's sharp edge is full-body `config_change`: canonical store content may carry the member bytes, but non-operator seat-facing `read`, projection, nudge, tool-result, schema, prompt, and error surfaces must not expose effective config member bytes (`SITREP-implementer-20260705-042308.md:47-58`). The plan makes that a dedicated Task 4b rather than burying it in Task 4, and it chooses an explicit redacted struct, not a filtered map (`s4-slice-4-plan.md:88-95`).

This maps to the live source seam: current `cmd/frank/main.go:268-302` returns full `migrate.Reader` output for `read`, while `internal/record/record.go:27-31` keeps headers string-valued and `internal/fieldspec/registry.json:72,112` keeps `record_kind` system-owned today. The plan's Task 4 plus Task 4b therefore hits the right source boundary.

### 4. Fixture bite

The plan requires run-FAIL-first for negative fixtures in the global constraints (`s4-slice-4-plan.md:21`) and carries concrete failing tests in each high-risk task. The required reviewer list is covered: S4-SCH2 in Task 8 (`:117-121`), S4-SC3 in Task 2 (`:55-64`), S4-NG4 in Task 3 (`:66-75`), S4-FR2 plus the frame carve-out in Tasks 1 and 10 (`:44-53`, `:128-131`), and config-change chain/crash legs in Tasks 4-6 (`:77-109`). The design's section-9 fixture list remains mirrored (`s4-slice-4-design.md:149-152`).

The tests do not assert through the retired broadcast path: Task 3 explicitly removes `Push`/`broadcast`/`QueuePush`/global `pending` from the serving path and requires a grep floor for those serving paths (`s4-slice-4-plan.md:66-75`). That matches the live current defect surface in `internal/channel/server.go:55-62`, `:116-134`, `:188`, `:262-267`, and `:448`.

### 5. Scope discipline and README ruling

The file map is explicit and SCOPE_DIFF-bounded (`s4-slice-4-plan.md:24-40`). It deliberately excludes `internal/seat/binding.go`, preserving the no-binding-table-shape-change condition (`:40`). The root `README.md` row is now RULED-IN only for one bounded claim-honest delta, cites the ruling relay, and says it lands only with Tasks 4-6 green (`s4-slice-4-plan.md:37`, `:133-137`).

The ruling itself is narrow: one README sentence changes only when the section-7 record lands; transport-only qualifier remains; no other README content is in fence; a second README delta needs another ask (`SITREP-orchestrator-planner-20260705-042636.md:18-44`).

## Non-Blocking Implementation Watch

Task 4's registry work must include whatever `FieldSpec` rows are required for the `{member, new_digest}` header contract, not only the `record_kind` enum token. I do not treat this as a plan blocker because the Task 4 interface and tests require `member` bounded enum plus conductor-recomputed `new_digest`, and the registry file is already in scope (`s4-slice-4-plan.md:77-86`). If execution discovers this cannot be expressed inside the current registry/seat-scope model without broadening ordinary lane authority, the plan already says stop and route via the orchestrator to m-2 and back to m-1.

Task 1's inbound frame refusal should preserve the same-connection test as the controlling requirement. If `bufio.Scanner` cannot support resume-after-too-long cleanly, the implementation should switch the reader primitive rather than weakening `TestInboundOversizeFrameTypedRefusal`; no boundary change is required.

## Remaining Gates

- Pair Planner dispatch still requires this approving PLAN-REVIEW as parent, mechanical `SCOPE_DIFF_RESULT: all-in`, no trigger/collision/amendment/OUT touch, and the standing route-back conditions.
- Live-seat designation, the actual section-7 operator authorization, E3 live gate, s4-close, and merge remain separate human/operator gates.
- No implementation work is authorized by this review.

## Verification

- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s4/s4-wire-plan/SITREP-orchestrator-planner-20260705-043600.md` -> OK.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s4/s4-wire-plan/PLAN-planner-20260705-124044.md` -> OK.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s4/s4-wire-plan/PLAN-planner-20260705-043152.md` -> OK; superseded by `PLAN-planner-20260705-124044.md`.
- `python3 ~/.codex/skills/tools/relay-lint.py --relay-root .relays/s4` -> only known `INDEX.md` header noise; no non-INDEX relay errors observed in output.
- `git rev-parse --short HEAD` -> `28dfa33`; `git rev-parse --short main` -> `28dfa33`.
- `git diff d0ea59b..e29e0b2 -- docs/sprints/2026-07-05-s4-slice-4/plans/s4-slice-4-plan.md` -> README ruling fold only.
- `git diff --stat d0ea59b..HEAD -- docs/sprints/2026-07-05-s4-slice-4/plans/s4-slice-4-plan.md docs/sprints/2026-07-05-s4-slice-4/RECONCILE.md` -> plan README ruling fold plus RECONCILE append; live working tree has no tracked diff.
- `go test -count=1 ./...` -> pass.
- `go vet ./...` -> pass, no output.
- `python3 ~/.codex/skills/tools/relay-lint.py .relays/s4/s4-wire-plan/PLAN-REVIEW-implementer-20260705-124139.md` -> OK.

ACTIONS_GIT_REF: no source/test/design/plan/sprint-spec edits; reviewed plan content at `main@e29e0b2` while live `HEAD/main` is `28dfa33`; wrote relay `.relays/s4/s4-wire-plan/PLAN-REVIEW-implementer-20260705-124139.md` plus `.relays/s4/INDEX.md` row.
FINAL_GIT_STATUS_SHORT: none - clean tree.
