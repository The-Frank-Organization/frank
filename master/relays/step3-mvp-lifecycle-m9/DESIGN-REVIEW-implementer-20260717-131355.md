## DESIGN-REVIEW — m-9 lifecycle half r2 full-byte re-review: MUST REVISE

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-lifecycle-m9
PARENT_DISPATCH_ID: step3-mvp-lifecycle-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — four findings have bounded technical resolutions; F4 becomes human-decision-required only if the planner retains the unmitigated diagnostic residual
GRILL_REQUIRED: no — the m-9 grill remains with the stage-4 full-worker design
DESIGN_DOC_ID: step3-mvp-design-m9-lifecycle-half
DESIGN_RECORD_KIND: design-doc
DESIGN_DOC_SHA256: 9c1b8ac0ac89174f42a0f55b4960a51db64c2d025576a7a49a82fc883a61ba62
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: step3-mvp-lifecycle-m9/DESIGN-planner-20260717-130400.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-7.planner, m-1.planner, m-3.planner, m-8.planner, operator
BUNDLE_ID: m-9-model-runtime
SUBJECT: must-revise — r2 fixes F3/F7, but local-reject accounting conflicts with itself and m-8 r4; recovery still has a reader/replay gap; EOF containment cannot enforce its successor interlock; K6 accepts an ungated residual; route-back is not CLEAN

DESIGN_REVIEW_VERDICT: must-revise

I re-reviewed the complete r2 document `master/domains/m-9-model-runtime/design/2026-07-17-mvp-lifecycle-half.md`, not only the F1–F7 edits. The reviewed bytes reproduce SHA-256 `9c1b8ac0ac89174f42a0f55b4960a51db64c2d025576a7a49a82fc883a61ba62`; `DESIGN_DOC_ID` and the directly-addressed Planner relay match.

The r2 fold cleanly resolves the old F3 denial-terminal contradiction and F7 stale normative hash. Its ordinary F59 ordering, canonical counter encoding, narrow F61 push guarantee, exact-turn/exact-lane replay custody, and non-re-ownership lines survive. Five blockers remain.

## Blocking findings

### R2-F1 — The local-reject path is internally contradictory and now stale against m-8 r4

Section 2.2 first says `egress_denied` produces no stream and therefore **no** `attempt_stream_end` (`...mvp-lifecycle-half.md:89`). The next paragraph says **every** DATA-P reject, expressly including `egress_denied`, emits `attempt_stream_end{disposition: rejected_local}` (`...:90`). Those cannot both be true. Policy denial is also not a local-validation rejection: it owns m-3's deny token and `phase=denied`, while local rejects own a disjoint reject reason and `phase=failed`.

The live m-8 r4 owner bytes at `168c24b75ce6f1fc4bfdc98b4225209e64558e2e164e006295e371f542a6698b` postdate r2 and make the split explicit (`...mvp-provider-contract.md:76-87`):

- policy deny ⇒ CTRL-C `denied(<m-3 token>)`, no stream;
- local pre-stream reject ⇒ CTRL-C `rejected_local(<reject_reason>)`, where the closed reason set is `{malformed_request, lane_capability_mismatch, replay_scope_violation}`, no stream, E0 `phase=failed`, no `deny_reason`;
- m-9's stream expectation for a local reject is **none**; the typed DATA-P reply is its terminal answer.

r2 instead bases its repair on m-8 r2/r3, omits `lane_capability_mismatch`, puts the local-reject member on m-9's **stream-end** enum, and asks m-10 to reconcile it against m-8 `denied`/absent (`...mvp-lifecycle-half.md:89-90,190,200`). That would collapse policy and validation classes and no longer matches the owner proposal routed in `step3-mvp-design-m8/RECONCILE-orchestrator-planner-20260717-131203.md` (CC context only to this seat; the acting return remains m-9.planner's).

Required revision: consume the final owner-disposed m-8/m-10 shape, keep policy `denied` disjoint from local `rejected_local`, include all three local reason tokens, and pin m-9's actual forward mapping: typed DATA-P reply, no `attempt_started`, no stream-end fiction, attempt failure, E0 `phase=failed` with absent `deny_reason`, and one attempt-budget count. If m-10 still needs an m-9 observation frame for a no-stream reply, m-10 must name a semantically honest frame/type; do not overload `attempt_stream_end` while claiming no stream.

### R2-F2 — Replacement recovery still has no proven record reader, and parked-unknown context is not optional

r2 correctly withdraws the nonexistent read of m-10-private rows, but replaces it with “committed conductor records (via `project`/`read`, its own mailbox)” (`...mvp-lifecycle-half.md:121-124`). The current conductor default projection is recipient-based: mailbox intents are created only for `TO`/`CC` recipients (`frank/internal/store/projections.go:137-174`), and default `project` reads that recipient mailbox (`frank/internal/store/store.go:202-247`). A record authored `FROM: m-9` is not automatically in m-9's own mailbox. The separate audit projection does list records by `FROM` (`...store.go:250-257`), but this design neither names that view nor proves it is the intended recovery reader. The corrected F1 path therefore still asserts a reader without specifying the read contract.

The second half overclaims “no-double-execute” from m-10's ticket fence. `UNIQUE(run_id, turn_id, tool_call_id)` and epoch fencing prevent consuming the **same identified ticket/call** twice; a fresh attempt can receive a new `tool_call_id` and re-request the same semantic side effect. If the replacement starts fresh without the parked `UNKNOWN_TOOL_OUTCOME` or an operator disposition, the model can unknowingly duplicate an effect. The amendment's requirement is honest park-and-new-attempt recovery, never silent replay; hiding the unknown from the next actor makes D-4 safety-significant, not merely optional (`...mvp-lifecycle-half.md:124,177,193`).

Required revision: pin an actual conductor recovery reader (for example, the authorized audit projection plus exact record/reference rules, or an explicit self-addressed record contract) and its session-fact writer. For parked unknown effects, require one of two owner-real gates before successor work: (a) a state-only m-10 restore/disposition frame that reaches the worker/model, or (b) m-10 withholds new turn/attempt admission until an operator explicitly disposes the unknown. Narrow the guarantee to identifier-exact one-shot execution; do not claim semantic no-double-execute from the ticket key.

### R2-F3 — The EOF tool-child “successor interlock” is not enforced by the consumed m-10 lease

Process-group SIGTERM→SIGKILL is a useful first containment mechanism, but r2's fallback is not executable across the named app-main-death window (`...mvp-lifecycle-half.md:116-119,199`):

- CTRL-W EOF means m-10/app-main is already gone, so the worker cannot “surface” an un-reapable child to that m-10 instance.
- The replacement app-main cannot `waitpid` old non-children; m-10 says this explicitly (`...mvp-ipc-manifest-seam-contract.md:80-93`).
- m-10's lease blocks on the **worker generation process** being reaped (`...:67-69`), not on every process in a tool process group. If the worker exits after `TOOL_TERMINATE_DEADLINE` while a descendant remains, the worker is terminal and a successor may lease while the old effect continues.
- A worker can reap its direct tool child; “reaps the tool process group” does not by itself prove all grandchildren are dead, especially when a subprocess creates another session/group. The one delayed-sentinel fixture proves one friendly tree, not containment of the allowed `bash` surface.

Required revision: define a platform-real containment and proof boundary for the whole effecting process tree (for example, a kill-on-close job/cgroup or equivalent supervised process-tree primitive, plus fail-closed platform support), and bind successor admission to that boundary—not merely to the worker PID. If the MVP accepts an uncontained-effect residual, it must be an explicit operator decision and the design must stop claiming overlap is blocked. The fixture battery needs direct child, grandchild/background, re-group/session, and unkillable/timeout legs.

### R2-F4 — K6 turns an opacity violation into an ungated platform residual

Section 2.8 adds the requested diagnostic census, but permits an “un-scrubbing platform” as a recorded residual (`...mvp-lifecycle-half.md:129-135`) while the relay declares `HUMAN_GATE_REQUIRED: no`. A core/heap/debug artifact containing the replay payload is a durable/cross-surface propagation path, not merely the amendment's same-UID secret-inspection residual. The m-8 custody law still says the opaque payload is never logged or surfaced.

Required revision: for the supported MVP platform set, make dump/profile/debug disable-or-scrub a serve/admission prerequisite and fail closed when it cannot be established. If the planner instead wants to ship a platform where the payload may persist, route `human-decision-required` to the operator with the exact residual; a pair seat cannot silently accept it. The canary fixture must distinguish “surface disabled/absent” from “artifact emitted and proven scrubbed.”

### R2-F5 — The route-back section contradicts the document's four required owner deltas

Sections 4–5 correctly say D-1/D-2/D-3 (and potentially D-4) require m-10/m-7-owned contract changes and that those seams remain owner-delta-gated (`...mvp-lifecycle-half.md:180-193`). Section 7 nevertheless says “route-back check: CLEAN” and “No locked m-x text is reopened” (`...:207-208`). New CTRL-W attach messages, a new attempt disposition/store state, and a broker attach-result taxonomy are exactly cross-owner interface changes.

Required revision: make §7 truthful: route-back is pending, list each owner/reader/writer and the exact acceptance needed, bind the consumed hashes only after their owner folds/reviews land, and retain the no-closure/no-lock gate until those bytes are confirmed. D-4's disposition becomes non-conditional under R2-F2 unless m-10 supplies the alternative operator-admission gate.

## Confirmed resolutions and non-blockers

1. **Old F3 resolved:** `turn_denied` is now pre-wire provider-send/routing denial; `DENIED_ABOVE_SET` continues and only terminates through `turn_exhausted`. Provider-wire and tool-effect facts are independent.
2. **Old F7 resolved:** §5 cites m-10 r12 `111ab95a…`; the remaining `79fcf742…` references are explicitly historical.
3. The standard F59 path remains correctly ordered and identifier-exact: complete/observed call → digest → authorize → atomic consume → re-digest → execute → invocation-identity outcome.
4. The string-counter rule remains complete for current m-9-emitted trust-bearing counters. The F61 push claim remains narrow and accurate.
5. Exact-turn/exact-lane replay-envelope custody remains correct apart from the F4 unsupported-platform residual. The live m-8 r4 also normalizes `reasoning_end.replay_envelope?` and fails closed on the legacy field; r2 should re-affirm that in its next current-basis fold.
6. The F70 problem is now honestly labeled owner-delta-gated rather than silently claimed solved. D-2/D-3 still must land before this seam closes; the bounded-retry exhaustion/fault result belongs in those owner bytes.

## Verification

- Target SHA-256: `9c1b8ac0ac89174f42a0f55b4960a51db64c2d025576a7a49a82fc883a61ba62`.
- Exact current basis hashes reproduced: amendment `2f75f2a1…`; m-10 r12 `111ab95a…`; m-7 r8 `ab0ed428…`; m-1 `7c8b09a6…`; m-2 `83d8e63e…`; m-3 r3 `70838f83…`. Current m-8 is r4 `168c24b7…`, which postdates and materially supersedes r2's local-reject seam basis.
- Parent DESIGN relay exact-file lint: `OK`.
- Full target read with line numbers; targeted sweeps covered all prior F1–F7 surfaces, stale hashes, local-reject enums, owner deltas, EOF descendants, unknown-effect recovery, K6 diagnostics, and route-back claims.
- Live code evidence for the recovery reader: `frank/internal/store/projections.go:137-174`; `frank/internal/store/store.go:202-257`.
- `frank/` source was not edited by this review.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design-doc edit, no `frank/` source/test edit, no PLAN, no IMPL, no credential/provider action.
FINAL_GIT_STATUS_SHORT: cwd root is not a git repo; `frank/` clean at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.
Next requested action: m-9.planner folds R2-F1–R2-F5 in one bounded r3 after the m-8/m-10 local-reject owner disposition is known, routes the recovery/containment/attach owner deltas through master, rebases the consumed seam hashes, then requests a fresh uniquely-parented final-byte review. No closure SITREP, reciprocal approval, interface lock, PLAN, or T4 authority should consume `9c1b8ac0…`.
