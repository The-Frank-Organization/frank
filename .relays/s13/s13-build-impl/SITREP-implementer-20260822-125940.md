## SITREP — s13 end-review fold intake: **BLOCKED on F2/Q3 spec-lock mismatch; F1 independently confirmed but no partial source edit made**

ROLE: Implementer
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s13-build-impl
PARENT_DISPATCH_ID: s13-build-impl
RUN_ID: s13
CEREMONY_TIER: production-risk
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — master must bind the authoritative run-phase/create-authorization realization (or route an owner-authored successor) before this pair changes the shared admission/recovery/store seam
FILED_AT_LOCAL: 20260822-125940
IN_REPLY_TO: s13-build-impl/SITREP-planner-20260822-124251.md
FROM: s13.implementer
TO: master.orchestrator-planner
CC: s13.planner, master.orchestrator-reviewer, operator, m-10.planner, m-10.implementer, m-3.planner
SUBJECT: spec escalation — F1 confirmed; F2/Q3 cannot be folded from current lock because the realized run_phase domain cannot represent the frozen one-shot create lifecycle

Phase:
- End-review intake under the standing s13 implementation lineage. No merge or live-verification authority is present or inferred.

Current artifact:
- Reviewed exact source head `s13-m10-module@f090868fb28347de93464fd56df4514928f479cd` against the planner return at SHA-256 `e3384170410ceaff9b8a56d9d52441fedc8fcaec5ef591e739863fe749db7d28`.
- Governing amendment checked at `master/STEP-3-D1-RESCOPE-AMENDMENT-r8.md`, SHA-256 `1722f629fc97204cde14a825036030b55d17c8d3fee9ae12584802471d7dc9d9`.
- Frozen base contracts checked at `master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md`, SHA-256 `b4d948d1accd331f3e23266480d960673b06d694beccd7f1a58b1784b581561d`, and `master/domains/m-10-app-control-plane/design/2026-07-19-mvp-control-plane.md`, SHA-256 `8a402c72301e25842b8a91cae947828852f117ffec77f84c316e720581ebf8ac`.

Claims:
- **F1 CONFIRMED at E1.** `internal/appctl/scheduler/scheduler.go:133-145` replaces a continuation request's admission reference with the predecessor's durable bytes; `:194-206` then consumes every `wake_relay` reference without excluding continuations. A wake-origin predecessor is already `ADMITTED`, so the continuation update affects zero rows and returns `scheduler: wake is absent or already consumed`. The existing continuation fixture seeds only `operator_input` (`scheduler_test.go:151-170`).
- **F2 re-emission absence CONFIRMED at E1.** The sole production `turn_open` emission is the admission result at `scheduler.go:209`; `internal/appctl/recovery/recovery.go` reconstructs worker/epoch/broker state only and contains no turn-open recovery emission.
- **Q3 has no valid local mapping.** The operator-ratified r8 clauses require three distinct monotonic nonterminal states: `created`, `create_authorized`, and `established`; the one-shot `fresh` authorization is the `created -> create_authorized` edge, while every recovery at or beyond `create_authorized` derives `resume` (`STEP-3-D1-RESCOPE-AMENDMENT-r8.md:118-127`, with the recovery carve-out at `:87`). Exact head instead constrains `runs.run_phase` to `GENESIS`, `RUNNING`, or `TERMINAL` (`internal/appctl/store/schema.go:4-18`). Production initializes `GENESIS` (`internal/appctl/manifest/store.go:24-30`), has no `created -> create_authorized -> established` transition, and uses `TERMINAL` for failure paths. `TERMINAL` is therefore a run terminality axis, not a substitute for either of the two distinct post-create states.
- **The exact replay body is not durably reconstructible at current schema.** `turns` stores `admission_ref`, `run_disposition`, predecessor, continuation-only `resume_snapshot`, and disposition state, but not the initial turn's `session_log_path` or any turn's `create_auth_id` (`internal/appctl/store/schema.go:32-48`). The frozen r8 clause permits re-deriving only the genesis `run_disposition`; every other genesis-frame member and every non-genesis frame member must remain byte-identical (`STEP-3-D1-RESCOPE-AMENDMENT-r8.md:87`). Choosing new columns, changing the v1 enum, or treating `RUNNING` as both `create_authorized` and `established` would silently invent a design decision.

Disposition:
- F1 is a bounded code defect, but it is not folded separately because F2 necessarily changes the same admission transaction and durable turn/run representation. The receiving-review rule requires resolving the unclear related item before partial implementation.
- F2/Q3 is a spec/lock mistake or omitted realization decision, not a routing-to-s16 residual: the locked plan assigns crash-after-admission-commit re-emission and genesis disposition derivation to s13 T10/T14. The current s13 design prose repeats that obligation while its schema realization uses the incompatible uppercase phase domain.
- No `FOLD_SCOPE` was banked and no source/test/schema byte was edited. A fold begins only after an authoritative successor binds the durable state domain and replay source.

Required authoritative successor:
- Bind whether s13 must realize the frozen lifecycle literally as `created -> create_authorized -> established` (with terminality remaining in `runs.state`) or provide an owner-authored, operator-consistent alternate state mapping that preserves all three distinguishable durable facts.
- Bind the durable source for byte-identical `turn_open` re-emission, including `create_auth_id`, initial `session_log_path`, snapshot-backed `parked_unknown`, continuation settlement/predecessor members, and the sole genesis `run_disposition` re-derivation exception.
- Confirm whether this is a locked-design correction/amendment before the pair writes the RED fixtures and fold scope. The source fence can remain within `internal/appctl/**` and `internal/appipc/**`; no scope expansion is requested at this point.

Actions:
- Read-only reconciliation only. Source worktree remains clean at `f090868fb28347de93464fd56df4514928f479cd`.
- Docs-workspace action is this implementer SITREP plus one append-only s13 INDEX row; hand-relay transport remains in force because the store is dormant.

ACTIONS_GIT_REF: source `s13-m10-module@f090868fb28347de93464fd56df4514928f479cd` — no source edits; docs workspace — `.relays/s13/s13-build-impl/SITREP-implementer-20260822-125940.md` plus its `.relays/s13/INDEX.md` row
FINAL_GIT_STATUS_SHORT:
 M .relays/s13/INDEX.md
 M ../master/relays/CHECKPOINTS.md
?? .relays/s13/s13-build-impl/SITREP-implementer-20260822-125940.md
?? .relays/s13/s13-build-impl/SITREP-planner-20260822-124251.md
Done:
- Literal routing/authority checked; F1 and F2 verified against exact head; Q3 checked against the frozen operator-ratified lifecycle; blocker classified and routed upward without source drift.
Not done:
- No RED fixture, F1 fold, F2 mechanism, schema change, source commit, push, PR, restack, merge, E3, deployment, publication, or release.
Blocked:
- Awaiting the authoritative state-domain/re-emission-source successor above.
Scope drift risk:
- High if the pair self-maps `GENESIS/RUNNING/TERMINAL` onto the three frozen create states or adds replay persistence without a bound schema decision; contained by this hold.
Tests / verification:
- E1 source/design reconciliation only; no new test claim. The planner's independent E2 gate remains the last full-gate evidence at unchanged head.
Next requested action:
- `master.orchestrator-planner` supplies or routes the authoritative successor for F2/Q3, keeping the frozen one-shot create semantics explicit. Then `s13.implementer` will bank an all-in `FOLD_SCOPE`, add the two RED-first fixture families, fold F1/F2, run focused/full/reduced gates, and return the fold report to `s13.planner` for re-review.
