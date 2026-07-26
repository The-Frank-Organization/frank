## RECONCILE -- REVISE: F101 and F104-E close, but interpreter identity is optional, journal commit has no wire/receipt or total state home, and two exit predicates contradict the frozen identity/crash model

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-stage6-rescope-decomposition-review-r4
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- operator re-scope ratification remains required, but rev4 is not ready for that gate until the remaining context, journal, and exit-predicate contradictions are removed
GRILL_REQUIRED: no -- the existing GRILL_LOCK remains sufficient; these are bounded contract corrections, not new product choices
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-051500.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: REVISE -- accept the soft-stable bundle and m-8-owned lowering; require non-optional interpreter identity and one cwd encoding, define the durable journal commit wire/receipt/state family and replay custody, and replace the impossible UNKNOWN+EXECUTED and forged-origin handoff predicates

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-051500.md` at SHA-256 `db46b61e40da11a84a708d5c9f2f350132ef9f37cce33a2dc00a305262ca603a`.

Proposed amendment rev4: `master/STEP-3-STAGE6-AMENDMENT.md` at SHA-256 `1c485e9d8f56e584725b6750bb7de58324f3773503815537213d572a90dad2e9`.

## Findings

### F103 -- BLOCKER, NARROW: the environment and multi-target fixes land, but exact interpreter/cwd identity still does not

The m-1 branch is valid: its frozen contract requires a sanitized bash environment and an explicit launch allow-list (`m-1:21,25-32`), and hashing the COMPLETE set actually presented closes the environment-binding defect. Splitting `apply_patch` to an ordered resolved-target-set digest also closes the single-resource contradiction.

Two exact-context defects remain at amendment `:81-105`:

- `shell_interpreter_ref` makes `binary_sha256` optional while claiming to bind the actual external interpreter. A mutable binary can retain the same path and reported version. That is not exact implementation identity. The content digest (or an equally immutable OS/package artifact digest) must be required for a passing exact-context predicate; inability to obtain it must yield typed `unknown`/hold, not a weaker pass.
- The claimed single cwd byte form is simultaneously "absolute" and "workspace-root-relative" (`:95-96`). Those are different encodings. With `workspace_root_id` already present, choose one canonical representation or define a tuple containing both with an equality rule; do not leave one scalar with incompatible forms.

Required correction: make interpreter content/package identity mandatory-or-unknown and define symlink resolution plus the version command/source. Pin one cwd serialization, including root itself and nonexistent-path behavior where relevant. Preserve the accepted environment and `apply_patch` branches; no sandbox or affected-resource claim is requested.

### F105 -- BLOCKER: the proposed atomic round commit has no cross-process wire or durable receipt, and its new states/custody are not total

The source map now correctly treats exact tool-call arguments and replay envelopes as content, and an atomic content+outcome+marker commit is the right direction. The mechanism is not executable yet:

- Amendment `:123-135` says m-10 atomically commits blobs, outcome row, and `round_marker`, but m-9 owns and holds the content. No m-9-to-m-10 frame carries the blobs/refs/marker, no size/chunking rule exists, and no durable receipt tells m-9 that the round is now resumable. The frozen `record_tool_outcome` frame carries only ticket/epoch/outcome/identity evidence and is explicitly one-way with no reply (`m-10 seam:240-257`). Rev4 neither supersedes nor composes with that wire. Without an ack-after-commit, live m-9 can advance on an uncommitted round and recreate the same crash gap.
- "a crash before that txn commits parks UNKNOWN" collapses distinct existing branches. A consumed local ticket maps to `UNKNOWN_TOOL_OUTCOME`; an unresolved provider attempt maps to `UNKNOWN_PROVIDER_OUTCOME`; pre-effect interruption is not effect-UNKNOWN. The amendment must preserve this source-specific totality.
- `CONTENT_LOST` and `REPLAY_UNRECOVERABLE` are called typed park/non-resume states without naming their enum/record home, producer, consumer, turn/attempt transition, E0/exit mapping, or operator resume action. They must be journal-resume dispositions or explicit m-10 lifecycle states, not free-floating labels.
- Replay persistence is still optional: "if a pair proves persistence infeasible" may select `REPLAY_UNRECOVERABLE` later. That changes a Tier-HARD recovery guarantee and must return through amendment/human decision, not pair discretion. Primary persistence also conflicts with the frozen K6 custody rule that replay envelopes die at turn terminal/park and never enter m-10 storage (`m-9 lifecycle:159-166`). The amendment may supersede that for same-logical-turn crash recovery, but run-lifetime retention (`:130-132`) must not keep replay blobs after that turn's terminal boundary.

Required correction: define the Tier-HARD round-commit request/receipt schema, content transport/bounds, transaction operands, ack-after-durable-commit rule, replay/idempotency key, and every crash cut relative to send/commit/receipt. Give each new resume disposition one durable home and total transition table while preserving existing effect-specific UNKNOWN states. Require replay persistence for the approved branch, with per-turn deletion and K6 no-surface rules; infeasibility must escalate through a new amendment. If the commit supersedes `record_tool_outcome`, say so and preserve its identity/fencing checks in the new transaction.

### F106 -- BLOCKER: crash-honesty and governed-handoff predicates are impossible under the frozen keys and stamped identity

Applicability AND semantic pass is now correct, the fixture manifest direction is sound, and the overhead intervals are total. Two machine predicates at `:171-188` cannot pass truthfully:

- `xit-crash-1` requires the crashed call's `tool_calls` row to park `UNKNOWN_TOOL_OUTCOME` and the m-10 observer to count exactly one `EXECUTED`. A canonical row cannot be both states. The cited `UNIQUE(run_id,turn_id,tool_call_id)` is one-ticket idempotency, not semantic-effect dedup: recovery requires a fresh ticket/tool_call_id, and the frozen worker explicitly says a fresh ID may be a semantic duplicate and ambient bash re-runs (`m-10 seam:183-194,240-244`; m-9 worker:174-178). It cannot prove no duplicate effect across retry.
- `xit-ho-1` says the second seat commits a relay record whose channel-stamped `FROM` equals the originating seat. Channel stamping must produce the second seat on a second-seat-authored record. The predicate either forges origin or fails. The honest proof is two records: the handoff committed with `FROM=origin`, then a second-seat response/ack committed with `FROM=second seat` and an exact parent/reference to the handoff.

Required correction: make crash-honesty prove **no automatic retry** after the UNKNOWN park and use a fixture-owned external `effect_id`/counter to show recovery caused no second effect; retain the UNKNOWN canonical row and do not demand an EXECUTED state for it. If an explicit informed retry is exercised, report its duplicate-risk/result separately under a fresh ticket rather than calling F59 semantic dedup. Define handoff as the two correctly stamped records plus their lineage, and bind the relevant conductor E1/E2 evidence in the composite proof.

The fixture manifest must carry the stable effect observer/key, the two handoff record expectations, per-fixture sample weights summing to exactly 30 turns/100 calls, and the named baseline artifact/config digests. A 20%-100% overhead waiver, if retained, must be a durable operator HUMAN_GATE record citing measured evidence and thresholds, never a chat-only bypass.

## Accepted closures and direction preserved

- **F101 is CLOSED at architecture grain.** `bundle_sha256` now excludes mixed-source provenance, recipe and digest locations are explicit, undeclared markers fail closed, and `bundle-soft-stability` makes the core invariant executable.
- **F102 remains CLOSED.** The m-10/m-9 authority boundary and m-5/m-6 removal-not-reassignment stand.
- **F103 environment and `apply_patch` branches are accepted.** Only interpreter/cwd exactness remains.
- **F104 is CLOSED at architecture grain.** B's producer/consumer order and E's logical/lowered component split put ownership where bytes exist; m-3 joins digests without creating a second source of truth. The dedicated lowered-tools digest is a derived component of m-8's same frozen body, not a rival truth.
- **F105 first-durable content and one-transaction objective are accepted.** The remaining defect is the absent wire/state/custody realization at architecture grain.
- **F106 applicability+verdict, fixture-manifest, numeric p95, total p50 interval, H-12 envelope, and non-gated utility direction stand.** No additional grill is requested.
- Broker-first placement, conditional H-24, H-16/H-26 gates, and every unchanged frozen artifact remain intact.

## Gate disposition

- Proposed stage-6 amendment rev4 `1c485e9d...`: REVISE; not ready for operator re-scope ratification.
- Pending joint-lock proposal `b7e1f0ef...`: remains HELD/superseded; no lock issued.
- Stage-6 re-lock remains unrequestable until corrected owner deltas, bundle/fixture manifests, reviews, and joins complete.
- PLAN, T4/code token, credentials, provider calls, release binding, live E3, merge, deploy, out-of-envelope use, and Step-3 close remain HELD. Step 2 remains closed.

## Required return

Return decomposition review r5 over new amendment bytes that: (1) makes interpreter identity mandatory-or-unknown and pins one cwd encoding; (2) defines the journal commit wire, durable receipt, source-specific crash table, resume-state home, and turn-scoped replay custody; and (3) replaces the impossible crash/handoff predicates with external-effect/no-auto-retry and two-record stamped-lineage proofs. Preserve all closures above.

## Verification

- Target SHA-256: `db46b61e40da11a84a708d5c9f2f350132ef9f37cce33a2dc00a305262ca603a`; exact-file lint: OK; directly addressed and indexed at pre-review EOF row 1878.
- Amendment rev4 SHA-256: `1c485e9d8f56e584725b6750bb7de58324f3773503815537213d572a90dad2e9`; r3 reviewer parent SHA-256: `d88c0400fa1bc7e097bc0ceab5c0f29aec9a253de83221ca8034182d0ab683c6`.
- All nine prior design finals, H-16 rev16 `a349a329...`, and H-17 census v3 `959b1928...` recompute to the r4 manifest.
- `frank/` is clean at `c78da3815a34480590071295c1e09bb7d53c10b6`; no product-source action was taken.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment/backlog/design/source byte, historical relay, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` with empty porcelain at `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof rerun after the append-only INDEX update and this final report-byte edit.
Next requested action: master folds the three bounded contract corrections and returns amendment rev5 for decomposition review r5; operator re-scope ratification remains held.
