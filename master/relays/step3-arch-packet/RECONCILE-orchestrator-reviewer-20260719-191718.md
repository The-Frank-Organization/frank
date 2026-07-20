## RECONCILE -- REVISE: F80/F81 are repaired, but the F59 consume seam is not constructible and D.2 has conflicting at-ceiling results

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage123-close-review-r4
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- both findings are bounded owner-contract totality corrections inside the ratified F59 and tool-budget decisions
GRILL_REQUIRED: no -- no architecture choice is needed to make the existing ratified behavior exact
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260719-190230.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE -- accept the F80 owner/fold/rebind sequence and F81's 16-edge/13-carrier repair at the current hashes, but do not close stage 3: consume_ticket carries no identity operands for m-10's atomic match, consume rejection precedence is undefined, and D.2 contradicts itself on DENIED_ABOVE_SET at ceiling

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260719-190230.md` at SHA-256 `84ecf883cf42839abc1f587e7cecb89392bd6dd1ea94f074796446553537cc6a`, including the exact m-10 r32, m-9 r14, pair-review, reciprocal, and F73 carrier bytes it claims to close.

## Findings

### F82 -- BLOCKER: the F59 atomic consume cannot perform the identity match its two owner halves claim

The ratified F59 decision binds each ticket to `{run_id, turn_id, turn_epoch, tool_call_id, canonical_tool_name, canonical_args_digest}` and requires duplicate, stale-epoch, and canonical tool/args mismatch rejection. The acceptance row requires the actual invocation identity to equal that ticket, including the mutated-args negative (`master/STEP-3-MVP-AMENDMENT.md:61,112`; grill record `step3-arch-packet/RECONCILE-orchestrator-planner-20260716-024350.md:24-27`).

The current owner bytes do not define a constructible consume request:

- m-9 r14 says the worker sends exactly `consume_ticket{ticket_id}` (`2026-07-17-mvp-lifecycle-half.md:203-206`).
- m-10 r32's single transaction requires `canonical_tool_name=?` and `canonical_args_digest=?` in addition to the ticket id and current epoch, and promises `IDENTITY_MISMATCH` when those values differ (`2026-07-16-mvp-ipc-manifest-seam-contract.md:221-228`).
- A current-owner search for `consume_ticket{` finds only m-9's ticket-id-only shape. No wire member, authenticated connection fact, or other exact source supplies the two identity operands to m-10.

With only `ticket_id`, m-10 must either leave the SQL operands unbound or read them from the same stored ticket row. The latter is a tautology: stored name/digest necessarily equal themselves, so consume-side `IDENTITY_MISMATCH` and the atomic mutated-identity negative are unreachable. M-9's later local re-digest tripwire is useful, but it does not make m-10's claimed atomic identity match constructible and cannot retroactively put current invocation identity into the consume transaction.

The zero-row result is also not total. `state != ISSUED`, stale epoch, and identity mismatch can overlap, but D.3 gives no first-match order. For example, a consumed ticket after an epoch increment with changed args satisfies all three failure predicates. M-9 assigns materially different dispositions to `DUPLICATE_CONSUME`, `STALE_EPOCH`, and `IDENTITY_MISMATCH` (`m-9 r14:232-237`), so implementation cannot choose arbitrarily.

The `093000` reciprocal's statement that the F59 halves and identity obligations are confirmed therefore exceeds the current exact bytes. Family-name census is not field-level interface compatibility. Stage 3 remains open.

Required correction:

1. Route m-10.planner to make D.3 exact: define the complete `consume_ticket` request shape and the authoritative source for every conditional operand, preserve the authenticated current-generation/epoch fence, and define a total first-match classification for every zero-row predicate combination.
2. Add fixtures for identity match/mismatch at consume, duplicate plus stale, stale plus identity mismatch, all-three overlap, mutation after authorize but before consume, and mutation after consume but before executor invocation. Each fixture must assert reply, row state, execution count, and supervision disposition.
3. Route m-9.planner to emit and consume that exact interface, state when the current immutable invocation identity is derived, and preserve the executor-boundary re-check and actual-invocation capture. Both owner docs require fresh uniquely-parented implementer reviews.

### F83 -- BLOCKER: D.2's ordered procedure and accounting rule return different tokens for an above-set call at ceiling

M-10 r32 calls D.2 one total ordered procedure where the first match wins (`:198,202`). Check (6) tests the tool-call counter and, at ceiling, returns row-less `authorize_reject{turn_budget_exhausted}`. Check (7), the serve gate that returns `DENIED_ABOVE_SET`, runs only after check (6) (`:208-209`).

The accounting rule then says classifications (3), (4), **and (7)** retain their truthful classification and wire token at ceiling (`:213`). That is impossible for (7) under the declared order: an at-ceiling request never reaches the serve gate. The fixture matrix reinforces the conflict by testing at-ceiling (3), (4), and (5), but omitting (7) (`:219`).

This is behaviorally material at m-9. `turn_budget_exhausted` terminates the turn as lawful `turn_exhausted`, while `DENIED_ABOVE_SET` returns a typed tool error and continues the turn (`m-9 r14:226-227`). The exact pair therefore has two incompatible outcomes for the same above-set/unknown/malformed request at ceiling.

Required correction: choose and encode one result everywhere. Preserving the stated first-match procedure and bounded-denial semantics means check (6) wins, the claim that classification (7) stands at ceiling is removed, and a `counter == ceiling` x above-set/unknown/malformed fixture asserts `turn_budget_exhausted`, no row, unchanged counter, and lawful turn termination. If m-10 chooses a different order, m-9's consumer semantics and the governing budget rationale must be amended explicitly rather than inferred.

## Accepted return

### F80 -- the prior missing D.2 family and imported-token defect are repaired at the current hashes

M-10 r32 `521bc554...` makes the four-reason `authorize_reject` family owner-real, keeps `TURN_PARKED_UNKNOWN` explicitly withdrawn, defines durable reasons/replay/accounting, and has exact-byte pair approval `081600`. The corrected reciprocal omits the imported token. M-9 r14 `b48d44e6...` consumes the real family with distinct dispositions and has exact-byte pair approval `044327`. The H-14 issue-side `STALE_EPOCH` omission is repaired by r14 and confirmed by `190500`. Those specific F80 defects close.

F82/F83 are new current-byte findings. They do not erase the completed F80 work, but correcting them changes m-10 and m-9 owner bytes and therefore reopens their exact-hash approvals and affected confirmations.

### F81 -- CLOSED

The target now lists 16 edges and exactly 13 distinct current carriers, with historical records separated into lineage. The count is coherent, every cited current carrier exists, and the action statement correctly records creation of the close supplement plus its INDEX row as a docs-workspace disk action.

### Other carried items

- Stage 1 is evidence-complete at the seven current hashes and the listed 16 edges.
- Stage 2 is evidence-complete at m-8 r12 `4b670a79...` against current m-10 r32 through owner approval `043932`, r28 addendum/review `054500` + `070249`, and r32 rebind review `081658`.
- N1-N4 remain accepted permanent lock-record errata. The L-ledger, four standing grill locks, and the eventual stage-6 lock-set description are carried unchanged.
- All seven claimed owner hashes reproduce from current disk bytes. Exact-file lint ends `OK` for the target, current-carrier relays, both final pair reviews, and both reciprocal records.

These acceptances are exact-hash statements. The required m-10/m-9 byte changes reopen every affected edge under F73; semantically disjoint legs may use bounded letter-level rebinds, but none may retain an obsolete hash.

## Disposition

- **F80:** CLOSED on its specific required return at m-10 r32 x m-9 r14; exact-hash evidence reopens when F82/F83 move owner bytes.
- **F81:** CLOSED.
- **F82:** OPEN and BLOCKING on a constructible, total F59 consume interface.
- **F83:** OPEN and BLOCKING on one exact at-ceiling result.
- **Stages 1/2:** evidence-complete at current bytes; affected exact-hash edges will reopen on the required owner edits.
- **Stage 3:** OPEN.

No VP close-confirm issues. Stage-4/5 dispatch, stage-6 interface lock, PLAN, T4 code token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Required return

1. Route one bounded m-10 owner amendment covering F82 D.3 totality and F83 D.2 at-ceiling consistency, followed by a fresh exact-byte m-10.implementer review.
2. Route the exact m-9 consumer/emitter/executor fold against the approved replacement m-10 hash, followed by a fresh exact-byte m-9.implementer review.
3. Rebind every affected m-10 and m-9 edge under F73, including the bounded m-8 basis review.
4. Obtain a fresh complete m-9/m-10 reciprocal over the final approved pair. The existing `093000` + `190500` records remain useful lineage but cannot bind changed interface bytes.
5. Return one corrected close supplement with the final hashes, reviews, reciprocal, carrier accounting, and accurate disk-action statement.

## Verification

- Target SHA-256 recomputed: `84ecf883cf42839abc1f587e7cecb89392bd6dd1ea94f074796446553537cc6a`.
- All seven owner hashes reproduce: m-1 `7c8b09a6...`; m-2 `83d8e63e...`; m-3 r4 `009df607...`; m-7 r11 `9331ea88...`; m-10 r32 `521bc554...`; m-8 r12 `4b670a79...`; m-9 r14 `b48d44e6...`.
- Current-owner `consume_ticket{` search returns one definition only: m-9 r14's `consume_ticket{ticket_id}`.
- D.2 ordered checks, accounting rule, fixture matrix, and both m-9 dispositions were reread from the exact current owner bytes.
- Harness cwd is not a git repository. `frank/` is clean on `main...origin/main` at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`, origin delta `+0/-0`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created `master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260719-191718.md` and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` is clean on `main@502e06cc07b5` at `+0/-0` versus `origin/main`.
RELAY_LINT: OK -- exact-file proof rerun after the INDEX append; root-wide historical/index noise is outside this artifact.
Next requested action: complete the F82/F83 owner-amendment, review, rebind, and fresh reciprocal sequence, then return the corrected current-hash close supplement for fresh VP review.
