## RECONCILE -- REVISE: worker-owned placement is viable, but rev6 has no durable log protocol, its negative reconciliation is unsound, and continuation admission is undefined

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-stage6-rescope-decomposition-review-r6
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- operator re-scope ratification remains required, but rev6 is not ready for that gate until durable-log, reconciliation, continuation, and fixture-manifest contracts are executable
GRILL_REQUIRED: no -- operator D7 settles the build-resume product choice; these findings concern the completeness and soundness of the selected worker-owned-log branch
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-064500.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: REVISE -- accept the worker-owned content-log direction and F106 pre/post predicate; require an actual durable append/recovery protocol and positive durability proof, sound total reconciliation, and a complete continuation-turn lifecycle and operator surface

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-064500.md` at SHA-256 `3a75cc519d8485e62c7b40421340c729453f23f5633753b59cce32a38b860f65`.

Proposed amendment rev6: `master/STEP-3-STAGE6-AMENDMENT.md` at SHA-256 `671624512f7bdaf727b8fd14274f9ce1a5cb76906e083b3c2fee06f065aa3423`.

## Findings

### F105-D1 -- BLOCKER: the "durable JSON session log" has no durable append/recovery contract, and its gate can pass without one

Amendment `:129-173` chooses a plausible placement, but "a plain JSON file" plus "best-effort local append" does not define a crash-recoverable artifact. A single JSON document is not safely appendable without a framing/rewrite protocol. Rev6 supplies no versioned record union, canonical encoding, file/run/turn identity, record sequence and predecessor key, integrity/checksum rule, generation writer fence, append durability linearization point, flush/fsync and directory-persistence rule, torn-tail/valid-prefix recovery, duplicate/conflict semantics, stale-writer disposition, size/rotation/compaction bound, or app-main-recovery rule for rediscovering the same per-run runtime path. Those are HARD resume semantics, not T4 presentation details. "Best-effort" also contradicts D7's durable-state claim: a write that may vanish without a defined failure result cannot establish a resumable prefix. The `session-log append <= 50 ms` metric at `:227-229` is uninterpretable until it names whether it measures a buffer write or the durable append linearization.

The Durability leg is currently vacuous. `xit-dur-1` at `:215` declares both real log resume and missing/corrupt-log re-derivation to be passes. An implementation that never writes one valid log record can therefore pass the leg while providing no durable session state or resume, so the gate cannot prove operator decision D7. Require a positive crash/replacement fixture that resumes from the exact last durably committed, integrity-valid log prefix and proves the expected context/round identity. Missing/corrupt-log fallback may have a separate honest degraded fixture and operability result, but it cannot satisfy the Durability leg.

Required correction: define the closed log schema and file identity; append/durability linearization; writer fencing and replay/conflict rules; every write/crash/torn-tail/reopen/GC cut; exact bounds and recovery algorithm; and a positive durability predicate that cannot be satisfied by degraded re-derivation. Measure the same durable operation the contract promises.

### F105-D2 -- BLOCKER: `parked_unknown` cannot support the claimed absent-means-settled inference

Rev6 `:156-167` claims m-10 discloses every uncertain effect and therefore a logged effect absent from `parked_unknown` is settled. The frozen interface disproves that premise. m-10's closed overlay contains only `{turn_id, tool_call_id, ticket_id, state in {UNKNOWN_TOOL_OUTCOME, PARTIAL_TOOL_EFFECT}, canonical_tool_name, canonical_args_digest}` (`m-10 seam :72`). It contains neither `UNKNOWN_PROVIDER_OUTCOME` nor terminal/positive settlement rows. Yet the crash table parks an in-flight provider attempt as `UNKNOWN_PROVIDER_OUTCOME` (`:77-83`), and rev6 itself says logs may be ahead of or behind m-10 (`amendment :150-154`).

The resulting cuts are not total:

- **Log ahead:** a logged provider output may belong to an m-10-unknown attempt, but no provider row appears in the overlay. Absence would falsely classify it as settled. The same rule cannot compare tool effects exactly because the log contract expressly carries no ticket/epoch (`:139`) while the overlay identity includes `turn_id` and `ticket_id`; the successor also has a new `turn_id`.
- **Log behind:** m-10 may hold a terminal tool outcome while the log lacks its result/round. The overlay is empty, so resume from the earlier log prefix has no positive settlement fact and can ask the model to repeat an already-settled effect. Missing/corrupt fallback has the same blind spot.
- **Outcome smuggling:** m-8's frozen `tool_result{tool_call_id, content}` is the provider input representation of a completed result (`m-8 contract :44-45`). Persisting and trusting that item is not outcome-neutral merely because the column is labeled content. It must be bound to m-10-authoritative settlement or downgraded; otherwise the log becomes a second claimant about what completed.
- **Provider retry:** a provider-unknown or log-behind cut has no overlay member to inform the replacement. Automatically issuing the next provider attempt would conflict with the frozen "never automatic provider resend" / user-requested-new-attempt rules (`m-10 seam :68,82-83`; MVP amendment `:32`).

Required correction: provide a total, identity-exact reconciliation contract for every log-ahead/log-behind and tool/provider crash cut. Either extend the legal m-10-to-worker continuation metadata with both uncertain and positive settlement references, or impose exact durable ordering that makes the negative inference provable; in either case include the complete comparison key, provider-UNKNOWN handling, and incomplete/torn-record behavior. A `tool_result` becomes settled context only with m-10-authoritative proof, never from overlay absence alone.

### F105-D3 -- BLOCKER: the fresh continuation turn has no trigger, durable schema, wire shape, or retry semantics

Rev6 `:156-173` names a fresh continuation turn and durable predecessor link but defines none of the owner-real lifecycle needed to create or consume it. The frozen `turns` row has only turn/run/epoch/state/admission-ref (`m-10 seam :275-281`), `turn_open` has no predecessor member (`:72-73`), and m-10 admission is sourced from operator input or a consumed wake row (`m-10 control plane :69-73,120`). There is no field name/type, source and same-run validation, `UNIQUE`/idempotency rule, one-successor or chain rule under repeated crashes, admission-ref inheritance rule, lease/epoch transaction, frame carriage to the worker, or crash-before/after-admission/send table. Without an exact predecessor/log-segment carrier, the replacement cannot know which durable prefix it is authorized to load.

The trigger is safety-relevant. "On replacement m-10 admits" reads as automatic, while the frozen contract parks and requires a new informed attempt and forbids automatic provider resend. D7 decides to build resume; it does not silently settle whether continuation is operator-requested or automatic, nor which action may issue the first post-crash provider attempt. Rev6 must explicitly preserve or supersede those rules rather than relying on the word continuation.

Finally, `journal_resume_disposition in {resumable, degraded}` is worker-local and has no defined carrier to m-10's operator surface. A missing/corrupt log cannot be an "honest" degraded mode if the operator cannot observe the mode and the permitted `resume_action`; §7's operability leg already requires `{unknown_effects[], resume_action}`. This need not become an m-3 E0 or a durable m-10 disposition, but it does need an exact typed producer/carrier/consumer and fail-closed behavior before work begins.

Required correction: define the predecessor column and `turn_open` member, admission trigger, source/validation/idempotency/chain constraints, same-transaction lease and admission behavior, every repeated-crash/send cut, and the precise relationship to the no-auto-resend/user-retry rules. Define the typed resumable/degraded diagnostic path and operator action before any continuation work.

### F106 -- BLOCKER, NARROW: the crash predicate is fixed, but the frozen manifest still omits its expected relation

The machine predicate at `:216` now correctly requires `counter_after_recovery == counter_before_recovery`, `invocations_after_recovery == 0`, and both counter snapshots equal 1 at the fixed cut. That closes the predicate defect. But the manifest schema at `:221-225` still adds only `effect_observer_key`; it does not carry the expected counter relation that the same paragraph claims is frozen. This leaves the expected values/relations in prose rather than in the hashed fixture instance and repeats the exact gap returned in r5.

Required correction: add a required structured member such as `effect_counter_expectation{counter_before_recovery:1,counter_after_recovery:1,invocations_after_recovery:0}` (or an equally exact relation encoding), bound to the named fixed `fault_injection_point`, in `xit-crash-1`'s frozen manifest entry.

## Accepted closures and direction preserved

- Operator decision D7 is accepted. A worker-owned content log is a viable architecture direction and avoids the rev5 m-10 blob-upload/read-back design.
- Dropping `commit_round`/`round_committed` closes the old F105-A cross-process blob-wire defect and the private-store-read half of F105-B. It does not close the new local durability, reconciliation, or continuation contracts above.
- F106's pre/post external-counter predicate now closes; the two-record stamped handoff remains closed; the stale r4 gate labels are corrected to r6.
- F101, F102, F103, and F104 remain CLOSED. The accepted interpreter/cwd, soft-stable bundle, ownership/DAG/digest, applicability/verdict, overhead-band, H-12, and non-gated utility directions stand.
- K6 replay custody and the source-specific `VOID` / `UNKNOWN_TOOL_OUTCOME` / `UNKNOWN_PROVIDER_OUTCOME` model remain valid. A degraded mode does not inherently require m-3 E0 or a durable m-10 disposition if an exact alternative operator-visible carrier is defined.

## Gate disposition

- Proposed stage-6 amendment rev6 `67162451...`: REVISE; not ready for operator re-scope ratification.
- Pending joint-lock proposal `b7e1f0ef...`: remains HELD/superseded; no lock issued.
- Stage-6 re-lock remains unrequestable until corrected owner deltas, exact-byte reviews, affected-consumer confirmations, bundle/fixture manifests, and joins complete.
- PLAN, T4/code token, credentials, provider calls, release binding, live E3, merge, deploy, out-of-envelope use, and Step-3 close remain HELD. Step 2 remains closed.

## Required return

Return decomposition review r7 over new amendment bytes that: (1) define the worker log as a crash-recoverable, fenced, bounded artifact and make `xit-dur-1` prove an actual durable resume; (2) replace absent-means-settled with a total identity-exact reconciliation across tool/provider and log-ahead/log-behind cuts; (3) define the complete continuation-turn/predecessor/admission/retry/operator-surface lifecycle; and (4) put the exact external-counter expectation into the frozen fixture manifest. Preserve every accepted closure above.

## Verification

- Target SHA-256: `3a75cc519d8485e62c7b40421340c729453f23f5633753b59cce32a38b860f65`; exact-file lint: OK; directly addressed and indexed at pre-review EOF row 1882.
- Amendment rev6 SHA-256: `671624512f7bdaf727b8fd14274f9ce1a5cb76906e083b3c2fee06f065aa3423`; r5 reviewer parent SHA-256: `ea3fdb2b0657812106f9c0e97824126b4a0c50d0afc496756d5d695ed9b30e43`.
- Relevant frozen bases recompute unchanged: m-8 provider contract `4b670a79...`; m-9 lifecycle `4d3bd14e...`; m-9 worker `cb7ff970...`; m-10 seam `d2ce9831...`; m-10 control plane `6fd1d655...`; MVP amendment `2f75f2a1...`.
- `frank/` is clean at `c78da3815a34480590071295c1e09bb7d53c10b6`; no product-source action was taken.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment/backlog/design/source byte, historical relay, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` with empty porcelain at `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof rerun after the append-only INDEX update and this final report-byte edit.
Next requested action: master folds the bounded log/reconciliation/continuation/fixture-manifest corrections and returns amendment rev7 for decomposition review r7; operator re-scope ratification remains held.
