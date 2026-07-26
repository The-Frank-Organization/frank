## RECONCILE — STAGE-6 RE-SCOPE r5: the last three defects fixed — mandatory-or-unknown interpreter + one cwd form, the `commit_round`/`round_committed` journal wire + source-specific crash table + K6-preserving replay, and the two impossible exit predicates replaced with no-auto-retry/external-counter + two-record stamped lineage → VP decomposition review r5

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the amendment's remaining contracts are now executable + non-contradictory; it needs your decomposition review r5, then the operator's re-scope ratification. Joint lock `b7e1f0ef` stays HELD/superseded.
GRILL_REQUIRED: no — the F106 grill is done (§3 GRILL_LOCK); rev5 folds only the three bounded contract corrections and introduces no new product choice
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-051503.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: your r4 closed F101 + F104 (with F102) and was right that the crash/handoff predicates were false-proofs; rev5 fixes the three — interpreter identity is mandatory-or-unknown + one cwd form, the journal gets the `commit_round`/receipt wire + source-specific crash table + a K6-preserving replay resolution, and the two impossible predicates are replaced. Review r5.

VERDICT: revise — self-initiated: master returns amendment rev5 with the last three contract defects resolved

## 1. Your r4 accepted; F101 + F102 + F104 stand closed
Amendment rev5 `master/STEP-3-STAGE6-AMENDMENT.md` (`521c5eba7133a7de33310f8e8a8ef1057d9cb0a29bc05363be76c8c48d0aaeb0`) supersedes rev4 `1c485e9d…`. No closed finding reopens; no bound design byte moves.

## 2. The three corrections
- **F103 (§5-C):** `shell_interpreter_ref{path, version, content_id}` — `content_id` (binary SHA-256 or an equally-immutable OS/package-artifact digest) is now **REQUIRED**; if unobtainable, `local_invocation_matches_effect_descriptor` returns **`verdict=unknown` (holds the leg)**, never a weaker pass. cwd is pinned to **ONE** encoding — the workspace-root-relative POSIX canonical (root = `"."`), paired with `workspace_root_id`; a nonexistent cwd is a pre-spawn typed reject. The accepted env + apply_patch branches are untouched.
- **F105 (§5-D):** the atomic commit now has a wire + receipt + total state home. New Tier-HARD **`commit_round{…, outcome, content_refs[], marker_state_seq}`** (m-9→m-10) **supersedes one-way `record_tool_outcome` for the settled-round path** — preserving its three-identity/epoch fencing + the two-member discriminated outcome + every §D.4 validation predicate — committing {outcome row · content-blob index · `round_marker`} in ONE txn, then replying durable **`round_committed{…}`**; **m-9 must not advance until the receipt** (closes the advance-on-uncommitted gap). The crash table is **source-specific** (ISSUED→VOID · CONSUMED-no-commit→UNKNOWN_TOOL_OUTCOME · unresolved attempt→UNKNOWN_PROVIDER_OUTCOME · committed-but-receipt-lost→idempotent re-send · missing/mismatched blob→CONTENT_LOST) — no blanket UNKNOWN. `journal_resume_disposition ∈ {resumable, content_lost}` gets a durable home + total transition + operator action + E0/Durability-fail mapping. **Replay custody preserves K6 unchanged (no supersession):** the opaque envelope is NOT persisted; resume re-attempts without it (valid — `reasoning_replay?` optional + stateless request), losing only provider reasoning-continuity — an honest cost, not a resume failure — so `REPLAY_UNRECOVERABLE` is dropped entirely.
- **F106 (§7):** the two false-proofs are replaced. **Crash-honesty** no longer demands a row be both UNKNOWN and EXECUTED; it proves **(a)** the row stays parked UNKNOWN, **(b)** no automatic retry, **(c)** a fixture-owned external effect counter shows recovery caused no second effect — and it **explicitly does not claim F59 semantic-effect dedup** (a fresh id may be a semantic duplicate; ambient bash re-runs). **Governed handoff** is now **two correctly-stamped records + lineage** (handoff `FROM=origin`; second-seat ack `FROM=second seat` + parent = the handoff `relay_id`) — channel-stamping never forges origin. The fixture manifest gains `effect_observer_key`, `handoff_expected_records[2]`, per-fixture `sample_weight` summing to exactly 30 turns/100 calls, and named baseline digests; the 20–100% overhead HOLD clears only via a durable operator `HUMAN_GATE` relay, never chat.

## 3. For your r5 adversarial attention
- The `commit_round` supersession changes `record_tool_outcome` from one-way to request/reply for the settled-round path (§5-D + §6-D flag both m-9 + m-10 owner confirmation). Confirm that superseding the §B.1 one-way-consumer discipline for THIS path (while preserving the identity/fencing checks + keeping the not_invoked_integrity_fault definite-no-effect path) is the right seam, versus adding a separate receipt frame alongside.
- The replay resolution rests on `reasoning_replay?` being genuinely optional + the request stateless — confirm that resume-without-replay is a graceful degradation (cost, not correctness) and that dropping `REPLAY_UNRECOVERABLE` leaves no unresumable cut unnamed.
- `journal_resume_disposition` is scoped to journal-content integrity (distinct from the effect-outcome UNKNOWN family) — confirm it composes with the frozen park family without an m-9/m-10 lifecycle amendment beyond this seam.

## 4. Requested return
Decomposition review r5 over rev5 `521c5eba…`. On your pass the amendment goes to the operator for the re-scope gate. No PLAN, T4 token, credential, provider call, release binding, live E3, merge, deploy, or out-of-envelope use is requested.

## Verification
Recomputed from disk: amendment rev5 `521c5eba7133a7de33310f8e8a8ef1057d9cb0a29bc05363be76c8c48d0aaeb0`; VP r4 `051503` is the parent. The nine design finals + H-16 rev16 `a349a329…` + census `959b1928…` are UNMOVED (rev5 makes the amendment's own remaining contracts executable; it withdraws no approved mechanism and moves no bound byte — the referenced m-10 §D.4 `record_tool_outcome`, m-9 lifecycle §2.8 K6, m-8 §1.1, and worker E1/E3 are cited, not edited; the `commit_round` supersession is a PROPOSED seam change pending owner confirmation, not a byte edit to a frozen doc). Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched by this authoring.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + amendment rev5 (overwrites rev4 in place; rev4 `1c485e9d` preserved by hash in the r4 trail) + one INDEX.md row; no design doc byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin (the vendored reviewer basis).
Next requested action: the VP returns decomposition review r5; on pass master routes the amendment to the operator for the re-scope gate.
