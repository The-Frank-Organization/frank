## RECONCILE -- REVISE: rev7 chooses the right subsystem, but its log fence is observational, m-10 cannot produce or rehydrate the claimed settlement state, and the exit gate regresses to seven-declared-as-six

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-stage6-rescope-decomposition-review-r7
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- operator re-scope ratification remains required, but rev7 is not ready for that gate until the new durable-log, settlement, continuation, and exit-fixture contracts are owner-real and internally decidable
GRILL_REQUIRED: no -- operator D7 and the build-it-properly confirmation settle the product choice; these findings are contract-completeness and consistency defects within that selected branch
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-073000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: REVISE -- accept the full durable-resume subsystem direction; require an enforceable log handoff and closed record/rotation grammar, producer-real positive settlement that can reconstruct content, durable continuation/report snapshots with total retry semantics, and a consistent exact fixture gate

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-073000.md` at SHA-256 `bd73ff131303a690ccf0bc90c53c4a868c4f10ab1735e66b1f701375e052a647`.

Proposed amendment rev7: `master/STEP-3-STAGE6-AMENDMENT.md` at SHA-256 `cb4ad602a6745e9d00cd1df64c329e11c57499db43172d0a5c3c1d5cdda7f736`.

## Findings

### F105-D1-R7 -- BLOCKER: a generation label is not a writer fence, and the log/rotation grammar is still open

Rev7 `:155-172` correctly replaces best-effort JSON with an fsync-linearized line log and valid-prefix recovery. But its writer-fence claim is not enforced. A record's `generation_id` is authored by the same worker writing the shared file; no current-generation oracle, exclusive file lock, per-generation segment ownership, or m-10-controlled handoff participates in the append. A retired generation can therefore append a correctly checksummed old-generation record after m-10's retirement point but before it exits. On reopen, the successor must retain old-generation records written before retirement because those are the content it is resuming, so "ignore stale-generation records" cannot distinguish the valid predecessor prefix from post-retirement stale writes. With a shared per-run `seq`, stale and successor appends can also interleave or choose conflicting sequence values. The declared record field observes who claimed to write; it does not fence the write or identify the durable generation boundary.

The record union is also not closed enough to drive the claimed recovery algorithm. `payload` has no per-`kind` schemas, canonical byte encoding is unnamed, `seq` has no canonical uint grammar/contiguity/duplicate rule, and `round_marker` has no exact payload or digest binding the ordered records that form the round. Consequently "longest valid run" does not decide a checksum-valid duplicate, gap, reorder, old-generation tail, or invalid record before a later marker. Rotation adds another unresolved state machine: `header.prev_file` names a predecessor but carries no sealed-segment identity/digest, active-segment rule, create/seal/switch fsync order, or crash table. Directory fsync is specified only on initial create, not rotation. A crash can leave two plausible active segments or a new segment whose predecessor was never sealed.

Required correction: define an enforceable exclusive-writer/handoff mechanism, such as m-10-ordered per-generation segments or an OS lock whose acquisition is ordered after predecessor termination, plus the exact durable boundary that separates trusted predecessor records from stale writes. Close every record payload and canonical encoding; sequence/gap/duplicate rules; marker membership/digest; segment seal/link/active selection; rotation fsync order; and all append/handoff/rotation crash cuts. The valid-prefix algorithm must return one deterministic prefix for each cut.

### F105-D2-R7 -- BLOCKER: the settlement manifest has an ownerless round field and cannot reconstruct log-behind content

The positive-settlement direction at `:174-193` is correct, but the proposed producer cannot emit the stated manifest from its canonical rows:

- `last_settled_round_index` has no m-10 source. Frozen `record_tool_outcome` remains `{ticket_id, turn_epoch, outcome, ...identity/evidence}` with no round index (`m-10 seam :243-246`); `provider_attempts` and `tool_calls` likewise carry no round mapping (`:275-282`). m-10 does not read the worker log. Rev7 adds neither a worker-to-m-10 round carrier nor a canonical m-10 row field, so this value would be guessed or unavailable.
- The `settled` union contains only tool-call terminals, while the log contains `provider_output`. `uncertain` includes `UNKNOWN_PROVIDER_OUTCOME`, but no positive terminal-provider entry tells the worker which logged provider output is trusted. The D1 record grammar also does not require `provider_output` to carry `attempt_id`. Provider-output replay therefore still lacks positive settlement identity.
- The manifest says identity-exact over `tool_call_id/attempt_id + canonical_args_digest`, but the `uncertain` element schema carries no args digest and the provider branch has no discriminant-specific key. Entries also omit source `turn_id`; a per-run log resumed through `turn -> cont1 -> cont2` has no stated immediate-predecessor-versus-full-ancestor scope, so the same schema is not total over the chain D3 permits.

More fundamentally, m-10's terminal state proves that a tool invocation returned; it does not possess the `tool_result.content`. In the stated log-behind cut, the result/round is absent from the log. Telling the worker "EXECUTED, do not reissue" does not reconstruct the provider-valid `tool_result` item or its content, and m-10's payload-free row cannot supply it. The same problem applies when a completed provider output is absent from the log. Rev7 therefore cannot both resume the exact context and claim log-behind is repaired merely by positive status metadata.

Required correction: choose an owner-real construction. Either impose and prove durable-content-before-terminal ordering that makes a terminal m-10 outcome with missing content impossible, or classify that cut as degraded/content-lost, or add an exact digest/content custody path. Define positive terminal-provider entries and bind every log record to attempt/turn/round identity. Source or remove `last_settled_round_index`; make the settled/uncertain unions discriminated and total across the full continuation ancestry. No branch may call missing content reconstructed merely because its status is known.

### F105-D3-R7 -- BLOCKER: continuation snapshots and degraded status are not durably re-emittable, and provider retry is contradictory

D3 `:195-210` writes only `predecessor_turn_id` with the new turn and lease, while `turn_open` also carries the settlement manifest and log path. It does not say the manifest bytes/digest or path are persisted in that same admission transaction, nor bind them to a committed state sequence. The claimed crash-after-commit/before-send re-emission therefore cannot be byte-identical from the committed turn row: m-10 would have to recompute a manifest from rows that may have changed, and the exact snapshot the worker must reconcile is not durable. This also leaves repeated delivery/equivalence undefined. D1 says the path rides `assign`/`turn_open`; D3 says `turn_open`. One exact carrier and one exact command shape are required, not a slash choice.

The degraded carrier remains only a sentence: `journal_resume_disposition` is "reported" m-9 to m-10 with no frame type/body, correlation key, ack/commit order, duplicate/conflict handling, or pre-work gate. Rev7 expressly refuses a durable m-10 row, yet the frozen m-10 terminal surfaces are committed-snapshot projections (`m-10 control plane :26,91,144`). An undurable worker report cannot populate that surface after worker/app-main failure and cannot support `xit-dur-2` as a durable proof. Persist the disposition/action on the continuation turn (or define and justify an exact durable alternative), then make its report/receipt and no-work-before-disposition ordering executable.

Finally, D2 says `UNKNOWN_PROVIDER_OUTCOME` permits a fresh attempt only through the frozen user-requested path (`:190-191`), while D3 says the automatically admitted continuation's worker makes a "fresh decision" by initiating a new attempt (`:202-203`). A model decision itself requires a provider attempt; renaming it fresh does not decide whether that send is automatic or operator-requested. Define a total first-action table: at minimum clean positive resume, uncertain tool, uncertain provider, and degraded/missing-content branches, with the exact hold/surface/operator action and the condition under which a new `attempt_id` may be opened. Preserve the no-automatic-provider-resend rule literally.

The inherited `admission_ref`, one-successor uniqueness, crash-cut direction, and reuse of the already-ratified G-2 consecutive-failure bound are otherwise coherent.

### F106-R7 -- BLOCKER, NARROW: the gate now has seven rows while declaring six, and the new resume expectations are not structured

Section 7 says "Six legs" at `:253` and "the six legs" gate close at `:284`, but its table now contains seven: Governance-binding, Durability, Degraded-honesty, Crash-honesty, Injection-visibility, Governed handoff, and Operability. Either make `xit-dur-2` a required subfixture of the Durability leg or consistently define a seven-leg gate, including the applicable typed predicate/evidence binding and sample accounting. A row cannot be both a new independent gate leg and absent from the gate cardinality.

The F106 counter expectation is now correctly structured in the hashed fixture manifest and CLOSES. `resume_prefix_expectation`, however, is only a field name plus prose. It has no schema or digest recipe for the claimed exact round/context identity, and `xit-dur-2` has no structured expected corruption point, disposition, or `resume_action`. Add exact shapes, for example a predecessor/round/log-prefix/context digest vector for `xit-dur-1` and a fixed corruption cut plus expected degraded disposition/action for `xit-dur-2`. Bind each to its fixture's `fault_injection_point` and typed predicate.

## Accepted closures and direction preserved

- Operator D7 and the build-it-properly confirmation stand. Treating resume as a real D1/D2/D3 subsystem is the correct scope.
- F105-D1's fsync durability point, torn-tail-to-last-marker direction, identity header, positive durability fixture, and separate degraded fixture are accepted as directions; the best-effort false pass is gone.
- F105-D2's withdrawal of absent-means-settled is accepted. Positive settlement for trusted `tool_result`, explicit provider UNKNOWN, and absent-means-not-settled are the correct safety direction once their sources/content cuts are complete.
- F105-D3's fresh linked turn, inherited task identity, one-successor constraint, repeated-crash chain, and G-2 bound are viable. No `INTERRUPTED -> RESUMING` state is required.
- F106 crash-counter predicate and hashed `effect_counter_expectation{1,1,0}` CLOSE. The two-record handoff, corrected r7 labels, F101/F102/F103/F104, K6 custody, source-specific UNKNOWN model, H-12, and every other prior accepted closure remain closed.

## Gate disposition

- Proposed stage-6 amendment rev7 `cb4ad602...`: REVISE; not ready for operator re-scope ratification.
- Pending joint-lock proposal `b7e1f0ef...`: remains HELD/superseded; no lock issued.
- Stage-6 re-lock remains unrequestable until corrected owner deltas, exact-byte reviews, affected-consumer confirmations, bundle/fixture manifests, and joins complete.
- PLAN, T4/code token, credentials, provider calls, release binding, live E3, merge, deploy, out-of-envelope use, and Step-3 close remain HELD. Step 2 remains closed.

## Required return

Return decomposition review r8 over new amendment bytes that: (1) make D1's writer handoff enforceable and close the record/segment/rotation recovery grammar; (2) give every D2 field an actual canonical producer, cover terminal provider outputs, and eliminate or honestly classify every missing-content/log-behind cut; (3) durably snapshot and re-emit the continuation manifest/path/disposition, define their wire/receipt semantics, and make first post-crash action total without automatic provider resend; and (4) repair the exit-gate cardinality and freeze exact positive/degraded resume expectations. Preserve every accepted closure above.

## Verification

- Target SHA-256: `bd73ff131303a690ccf0bc90c53c4a868c4f10ab1735e66b1f701375e052a647`; exact-file lint: OK; directly addressed and indexed at pre-review EOF row 1884.
- Amendment rev7 SHA-256: `cb4ad602a6745e9d00cd1df64c329e11c57499db43172d0a5c3c1d5cdda7f736`; r6 reviewer parent SHA-256: `d41dd9ef7aa2c1fdf8b701a00a8c448306aaef2d16e7260b6141a282f3d98972`.
- Relevant frozen bases remain the reviewed bytes: m-8 provider contract `4b670a79...`; m-9 lifecycle `4d3bd14e...`; m-9 worker `cb7ff970...`; m-10 seam `d2ce9831...`; m-10 control plane `6fd1d655...`; MVP amendment `2f75f2a1...`.
- `frank/` is clean at `c78da3815a34480590071295c1e09bb7d53c10b6`; no product-source action was taken.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment/backlog/design/source byte, historical relay, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` with empty porcelain at `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof rerun after the append-only INDEX update and this final report-byte edit.
Next requested action: master folds the bounded D1/D2/D3 and fixture-gate corrections and returns amendment rev8 for decomposition review r8; operator re-scope ratification remains held.
