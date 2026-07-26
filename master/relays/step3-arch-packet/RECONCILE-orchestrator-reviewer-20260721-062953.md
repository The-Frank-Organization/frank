## RECONCILE -- REVISE: r5 closes interpreter/cwd and handoff, but the journal still has no executable blob protocol or legal recovery path, its disposition/E0 home is undefined, and the crash counter can still false-pass

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-stage6-rescope-decomposition-review-r5
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- operator re-scope ratification remains required, but rev5 is not ready for that gate until the journal transport/recovery contract and crash fixture are executable
GRILL_REQUIRED: no -- the existing GRILL_LOCK still covers the product choices; these are contract-completeness corrections to the selected journal and fixture branches
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-052500.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: REVISE -- accept F103 and the two-record handoff; require a closed blob/commit protocol, an owner-real journal rehydration and lifecycle transition, a durable m-3-compatible CONTENT_LOST path, and an independently decidable external-counter fixture

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-052500.md` at SHA-256 `d5d444b49a59c9146546b13362c1b29c8de9d77dd2cdf58affda9547263e992c`.

Proposed amendment rev5: `master/STEP-3-STAGE6-AMENDMENT.md` at SHA-256 `521c5eba7133a7de33310f8e8a8ef1057d9cb0a29bc05363be76c8c48d0aaeb0`.

## Findings

### F105-A -- BLOCKER: `commit_round` is a settlement reference, not the promised content transport

Amendment `:133-145` says m-9 first "streams" blobs over a "defined chunk sequence," but no such frame is defined. The only declared request, `commit_round`, carries `{kind, blob_digest}` references, not bytes. There is no upload message type/body, chunk ordinal/count/size or byte encoding, staging key, completion/ack rule, or requirement that m-10 verify every referenced blob is durably present and digest-valid before committing the round. The transaction at `:141-142` commits index rows, the outcome, and the marker after out-of-transaction staging; it does not explain when a staged blob becomes part of the resumable round or how an interrupted upload is classified. A phrase saying the chunk sequence is defined cannot substitute for the sequence.

The reference/settlement schema is also under-specified:

- `content_refs.kind` has no closed domain, canonical blob-byte representation, ordering/linkage rule, or reconstruction algorithm. The source map includes tool-call arguments, provider output items, tool results, and checkpoints; a list of kind+digest pairs does not by itself reconstruct the ordered `input[]` with its tool-call identities.
- One provider response may carry multiple tool calls, but `commit_round` has one singular `outcome:<record_tool_outcome payload>`. The amendment must pin whether one journal round contains one call, all calls in a provider round, or a sequence of per-call settlements, and make the marker atomicity/cardinality match.
- "same `round_index` is idempotent" is not a durable replay key or equivalence rule. At minimum the key must be scoped by run+turn, and the contract must define equivalent duplicate, conflicting duplicate, stale-generation delivery, missing/mismatched staged content, and crash-before/upload/during-upload/after-upload/before-commit/after-commit-before-receipt cuts. `marker_state_seq` has no source, domain, or comparison rule.
- The reply is named a durable receipt, but its body contains only `{round_index, marker_state_seq}`. The inherited A.2 `re` correlation is necessary but does not define the persisted equivalence key or prove which run/turn/content vector committed. The receipt and same-receipt replay need to bind the complete committed identity.
- The universal per-blob bound at `:135,170-172` cites the §2a **max captured tool output** ceiling, which bounds tool results, not provider-stream output or assembled tool-call arguments. Each content kind needs its actual limit plus a total per-round/run bound; chunk payload size must account for frame-envelope overhead below `FRAME_MAX`.

Required correction: define the closed CTRL-W blob staging/chunk protocol and reply/error family; canonical bytes and reconstruction order per content kind; correct per-kind and aggregate bounds; durable staging/commit validation and orphan handling; complete round key/equivalence/conflict table; multi-tool-call cardinality; and every upload/commit/receipt crash cut. A standard staged-blob design is valid, but the outcome+refs+marker transaction must reject absent/corrupt operands and make only a fully reconstructible round resumable.

### F105-B -- BLOCKER: the proposed resume path contradicts the frozen fresh-turn/private-store lifecycle

Rev5 `:154-168` makes m-9 run a resume-time integrity check "over m-10's durable journal/round rows," continue from the last marker, and re-attempt without `reasoning_replay`. That path does not exist on the frozen surfaces:

- Frozen m-9 says the transcript is in-memory, m-10 rows are private/frame-only, and the worker **cannot read them back** (`m-9 full worker:85-88`). It further says replacement starts a **fresh turn**, with no transcript reload (`:88,155`).
- Frozen m-10 makes its database private and requires every other consumer to use frames (`m-10 seam:269-273`). On retirement it parks the turn `INTERRUPTED`, mints E+1, and proceeds to replacement/new admission (`:77-108`); there is no `INTERRUPTED -> RESUMING/ACTIVE` transition or journal rehydration command.
- Rev5 defines only the write-side `commit_round`/receipt. It defines no authenticated m-10-to-m-9 journal inventory/read/chunk frames, no integrity-result writeback, and no admission/lease/epoch rule telling a replacement whether it resumes the old `turn_id` or starts a linked new turn.

Therefore the planner's requested composition check has a negative answer: `journal_resume_disposition` does **not** compose with the frozen park family "without an m-9/m-10 lifecycle amendment." K6 itself can stay unchanged, and the choice to omit an optional replay envelope is coherent only after a legal full-input reconstruction path exists. Stateless provider transport does not create that app-side path.

Required correction: choose and specify either (a) same-logical-turn resumption, including the m-10-owned `INTERRUPTED` transition, epoch/lease/admission rules, or (b) a fresh-turn continuation with an explicit durable predecessor/journal binding. Define the authenticated journal inventory/blob-read protocol and reconstruction equality, how unknown-effect disclosures are inserted before any re-attempt, and all replacement/crash cuts. Explicitly supersede the frozen no-reload/fresh-turn clauses through m-9 and m-10 owner deltas and a joint confirmation; do not describe this as only a `record_tool_outcome` seam change.

### F105-C -- BLOCKER: `journal_resume_disposition` has neither a durable record home nor a legal E0/operator path

Amendment `:154-159` names an enum and calls it durable, but gives no table/row key, writer, persisted fields, write frame, read frame, or transition that stores it. Naming m-9 as producer cannot make the value durable in m-10's private sole-writer store. It also says `content_lost` is surfaced as an "honest E0 marker," but frozen `m3.app_event.v1` is the closed provider-attempt schema `{event_kind:provider_attempt, scope:attempt, phase:...}` (`m-3:119-148`). No journal-integrity event or field exists, and m-3 is absent from §6-D's owner/confirmation edge.

The operator action "abandon-run OR start-a-new-attempt" is likewise not a transition: it does not say provider attempt, turn, or run; what state changes; whether the lost-content lineage remains visible; or how it composes with the frozen operator surface's no-clear/no-forge rule. A new attempt against missing settled context could silently repeat effects, so this ambiguity is safety-relevant.

Required correction: define the durable disposition row/key/schema, sole writer, producer request and consumer/read surface, and total transitions. Either add an m-3-owned E0 schema delta with its exact carrier/populator semantics and place m-3 in the DAG, or remove the E0 claim and use a typed m-10 operator-state projection plus separately bound exit evidence. Pin `abandon`/`new run or turn` actions and preserve the terminal `content_lost` record in every branch.

### F106 -- BLOCKER, NARROW: the external counter is present but its machine predicate still permits a false proof

The two-record stamped handoff at `:221` is now correct. Crash-honesty at `:219`, however, defines the external assertion as "effect count <= number of actual invocations." That relation does not compare the state before and after recovery and can pass when two invocations caused two effects. The separate prose assertion "no automatic retry" is not given its observer/key or an exact numeric relation, so the external counter is not yet an independent no-second-effect proof. Merely adding `effect_observer_key` to the manifest does not bind expected observations or the crash cut.

Required correction: bind `xit-crash-1` to a fault point and external tool contract, and declare machine fields/relations such as `{counter_before_recovery, counter_after_recovery, invocations_after_recovery}` with `counter_after_recovery == counter_before_recovery` and `invocations_after_recovery == 0`. If the fixture deliberately crashes after the first effect becomes externally visible and before outcome commit, require both counter snapshots to equal 1. Keep the canonical UNKNOWN row unchanged, and keep any explicit informed retry in a separate fixture/result. Put the expected counter relation, not only its key, in the frozen fixture manifest.

### Gate bookkeeping -- NARROW

Amendment §11.1 (`:262`) and §12 (`:274`) still request VP decomposition review **r4**, while the Status, relay, and actual gate are r5. Correct both before operator routing so the exact prerequisite is unambiguous.

## Accepted closures and direction preserved

- **F103 now CLOSES at architecture grain.** `content_id` is mandatory-or-unknown; cwd has one root-relative POSIX encoding with root/nonexistent behavior. The already-accepted env and `apply_patch` branches stand.
- **F101, F102, and F104 remain CLOSED.** No soft-stability, ownership, DAG, or digest-join regression was found.
- **F106 governed handoff CLOSES.** Origin and second-seat records now carry their own channel stamps and exact lineage. The applicability/verdict rule, overhead intervals, durable operator waiver, manifest/baseline direction, H-12, and non-gated utility direction stand.
- The source-specific `VOID` / `UNKNOWN_TOOL_OUTCOME` / `UNKNOWN_PROVIDER_OUTCOME` split is correct. A request/reply settled-round commit can be the right replacement for the one-way outcome path, once its full wire, lifecycle, and replay table are defined by the owners.
- Keeping `reasoning_replay` in K6 in-memory custody and omitting it after replacement is a plausible graceful-degradation branch; no `REPLAY_UNRECOVERABLE` token is required merely because that optional envelope is gone. This does not close the separate transcript/journal reconstruction gap.
- Broker-first placement, conditional H-24, H-16/H-26 gates, and every unchanged frozen artifact remain intact.

## Gate disposition

- Proposed stage-6 amendment rev5 `521c5eba...`: REVISE; not ready for operator re-scope ratification.
- Pending joint-lock proposal `b7e1f0ef...`: remains HELD/superseded; no lock issued.
- Stage-6 re-lock remains unrequestable until corrected owner deltas, bundle/fixture manifests, reviews, and joins complete.
- PLAN, T4/code token, credentials, provider calls, release binding, live E3, merge, deploy, out-of-envelope use, and Step-3 close remain HELD. Step 2 remains closed.

## Required return

Return decomposition review r6 over new amendment bytes that: (1) finish the staged-blob/write/commit/receipt protocol, including exact reconstruction and replay semantics; (2) add the owner-real read/rehydration and lifecycle transition that makes resume possible; (3) give `journal_resume_disposition` a durable m-10 home and a valid m-3-or-operator visibility path; (4) make the external-counter fixture compare pre/post recovery state; and (5) correct the stale r4 gate labels. Preserve every accepted closure above.

## Verification

- Target SHA-256: `d5d444b49a59c9146546b13362c1b29c8de9d77dd2cdf58affda9547263e992c`; exact-file lint: OK; directly addressed and indexed at pre-review EOF row 1880.
- Amendment rev5 SHA-256: `521c5eba7133a7de33310f8e8a8ef1057d9cb0a29bc05363be76c8c48d0aaeb0`; r4 reviewer parent SHA-256: `3e51b4ec3c26510b1e4b02d9e512c50b078f240cf1fd67cd0889f30659c6bacc`.
- All nine prior design finals, H-16 rev16 `a349a329...`, and H-17 census v3 `959b1928...` recompute to the r5 manifest.
- `frank/` is clean at `c78da3815a34480590071295c1e09bb7d53c10b6`; no product-source action was taken.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment/backlog/design/source byte, historical relay, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` with empty porcelain at `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof rerun after the append-only INDEX update and this final report-byte edit.
Next requested action: master folds the bounded journal/lifecycle/disposition/fixture corrections and returns amendment rev6 for decomposition review r6; operator re-scope ratification remains held.
