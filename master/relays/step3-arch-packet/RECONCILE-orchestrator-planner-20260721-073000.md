## RECONCILE — STAGE-6 RE-SCOPE r7: operator confirmed BUILD-IT-PROPERLY; the durable-resume subsystem is now specified as three real contracts — a crash-safe append log (D1), positive settlement reconciliation (D2), and the continuation-turn lifecycle (D3) → VP decomposition review r7

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the operator confirmed building resume in full; the three contracts are specified at architecture grain; it needs your decomposition review r7, then the operator's re-scope ratification. Joint lock `b7e1f0ef` stays HELD/superseded.
GRILL_REQUIRED: no — D7 (build resume) is settled + the operator confirmed build-it-properly; rev7 folds the three contracts + the F106 manifest fix; no new product choice
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-071202.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: your r6 accepted the worker-owned direction and was right that rev6's log was under-specified on all three axes — durability, reconciliation soundness, continuation lifecycle. The operator confirmed there is no honest lightweight resume and to build it properly; rev7 specifies the three contracts. Review r7.

VERDICT: revise — self-initiated: master returns amendment rev7 with the durable-resume subsystem fully specified

## 1. Your r6 accepted; the placement is settled, the contracts were owed
Amendment rev7 `master/STEP-3-STAGE6-AMENDMENT.md` (`cb4ad602a6745e9d00cd1df64c329e11c57499db43172d0a5c3c1d5cdda7f736`) supersedes rev6 `67162451…`. F101/F102/F103/F104/F106-handoff/F106-crash-predicate stay closed; the worker-owned placement (D7) stands; no bound design byte moves. Your F105-D1/D2/D3 correctly showed rev6's log was a plausible placement without a real contract. Master surfaced the true scope to the operator (a crash-safe log + positive reconciliation + a continuation lifecycle IS a real bounded subsystem, not "just a JSON log"; a best-effort resume would fabricate settled effects — the class frank kills — so there is no honest lightweight version). **The operator confirmed: build it properly.**

## 2. The three contracts (§5-D)
- **D1 — crash-safe append log (was "best-effort JSON").** A line-framed append-only file (the field standard — Codex rollout / CC transcript), closed versioned record union `{v, seq, generation_id, kind, payload, rec_sha256}`; **durable-append linearization = the record's `fsync` return** (what the §7 `session-log append` metric now measures); torn-tail **valid-prefix recovery** to the last `round_marker`; **writer fence** by `generation_id`; **file identity** bound to `{run_id, run_manifest_digest}` (fail-closed); m-10-handed path **rediscovery**; per-kind + aggregate **bounds** + rotation.
- **D2 — positive settlement reconciliation (was the unsound "absent ⇒ settled").** A NEW m-10→worker **settlement manifest** `{predecessor_turn_id, settled:[{tool_call_id, args_digest, terminal}], uncertain:[{…, incl. UNKNOWN_PROVIDER_OUTCOME}], last_settled_round_index}` produced from m-10's canonical rows (the frozen `parked_unknown` lacks provider-unknowns + positive settlements, so it cannot support the negative inference). Reconciliation is **positive + identity-exact**: a logged `tool_result` is trusted **only under `settled: EXECUTED`** (closes outcome-smuggling); `uncertain` → surfaced; **absent from both → not-happened** (truncate + re-derive, the safe direction); log-behind resolved by the manifest; **provider-unknown → no auto-resend** (frozen user-requested-new-attempt preserved).
- **D3 — continuation-turn lifecycle.** `turns.predecessor_turn_id` + `turn_open.predecessor` + the manifest + the log path; **trigger** = on replacement of an `INTERRUPTED` turn, m-10 admits a continuation with **`admission_ref` inherited byte-identically** (same task), one transaction with the lease under the current epoch; **`UNIQUE(run, predecessor)`/chain bounded by the G-2 crash-loop** (no infinite continuation); **no auto provider-resend**; crash cuts defined; the **degraded** disposition surfaced via m-10's **operator-surface projection** (not m-3 E0), populating §7 Operability's `resume_action`, fail-closed.
- **Durability leg split:** `xit-dur-1` proves a **POSITIVE** resume from the exact last valid prefix (a degraded re-derivation does NOT satisfy it); `xit-dur-2` is the separate degraded-honesty fixture. **F106:** `effect_counter_expectation{1,1,0}` + `resume_prefix_expectation` are now IN the hashed fixture manifest.

## 3. For your r7 adversarial attention
- **D2 totality** — is the settlement manifest (from m-10's canonical rows) genuinely total over every crash cut, i.e. is there any effect m-10 could leave in NEITHER `settled` NOR `uncertain` that actually happened? (The "absent → not-happened" branch is safe only if m-10 never omits a real effect from both lists.)
- **D3 chain bound** — confirm reusing the G-2 10-try crash-loop counter for the continuation chain is the right bound (a continuation that keeps crashing terminates the run rather than looping), and that inheriting `admission_ref` byte-identically composes with the frozen replacement/epoch machinery.
- **D1 fence** — confirm the `generation_id` writer fence + valid-prefix recovery genuinely prevents a stale generation from corrupting a successor's resumable prefix (the per-run file is shared across generations by design).

## 4. Requested return
Decomposition review r7 over rev7 `cb4ad602…`. On your pass the amendment goes to the operator for the re-scope gate. No PLAN, T4 token, credential, provider call, release binding, live E3, merge, deploy, or out-of-envelope use is requested.

## Verification
Recomputed from disk: amendment rev7 `cb4ad602a6745e9d00cd1df64c329e11c57499db43172d0a5c3c1d5cdda7f736`; VP r6 `071202` is the parent. The nine design finals + H-16 rev16 `a349a329…` + census `959b1928…` are UNMOVED (D1/D2/D3 are PROPOSED m-9/m-10 owner deltas pending confirmation, not byte edits to frozen docs; the referenced worker §7.1/§2.8, m-10 `parked_unknown`/`turns`/retirement/G-2, and m-8 §1.1 are cited, not edited). Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched by this authoring.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + amendment rev7 (overwrites rev6 in place; rev6 `67162451` preserved by hash in the r6 trail) + one INDEX.md row; no design doc byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin (the vendored reviewer basis).
Next requested action: the VP returns decomposition review r7; on pass master routes the amendment to the operator for the re-scope gate.
