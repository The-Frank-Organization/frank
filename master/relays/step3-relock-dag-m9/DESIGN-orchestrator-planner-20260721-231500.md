## DESIGN — step3-relock-dag-m9: §11 lane-2 (interface DAG) m-9 scope — the durable-resume subsystem (item D, the critical path) + the M9-D2 affected final + the B-carriage / C-executor / E-`logical_surface_digest` deltas, over frozen worker r7 `cb7ff970…` + lifecycle r21 `4d3bd14e…`

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-dag-m9
PARENT_DISPATCH_ID: step3-relock-broker-confirm
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the ratified re-scope + the lane-1 close already authorize lane 2; a HUMAN_GATE returns only on a genuine cross-owner conflict or a spec mistake you escalate
GRILL_REQUIRED: no — product choices settled by the ratified §3 GRILL_LOCK; this executes the ratified §5/§6 contracts
DESIGN_DOC_ID: step3-relock-dag-m9
FROM: master.orchestrator-planner
TO: m-9.planner
CC: m-9.implementer, operator, master.orchestrator-reviewer, m-10.planner, m-10.implementer, m-2.planner, m-3.planner, m-8.planner, m-1.planner
SUBJECT: your lane-2 scope — the biggest of the six: item D (worker-owned crash-safe resume) + M9-D2 + B/C/E producer/consumer halves; governed additive deltas over your frozen r7/r21 under F73; two-sided D + B join records

m-9 pair — §11 lane 1 (the broker study) is CLOSED (VP integration-confirm r2 `step3-relock-broker-confirm/RECONCILE-orchestrator-reviewer-20260721-225500.md`); the broker delta rev8 `64f9136e…` is settled input. The operator opened **lane 2 (the §6 interface DAG) in full**. This is your complete lane-2 scope. Run your normal pair cycle (planner authors, implementer adversarially reviews the final bytes). Everything is a **governed additive delta over your frozen finals** — worker r7 `cb7ff970…` + lifecycle r21 `4d3bd14e…` stay the historical locks; no in-place byte edit; deltas reviewed under F73 with consumer confirmations. The ratified contracts are `master/STEP-3-STAGE6-AMENDMENT.md` **§5-B/C/D/E + §6** (rev12 `1125b0a0…`).

### Grain (amendment §5-D, binding)
The amendment fixes the **decomposition** — the load-bearing decisions, seam ownership, and acceptance properties + your named Tier-HARD DESIGN obligations. The subsystem **internals** (exact record grammar, the writer-fence mechanism, the segment/rotation state machine, exact frame/table) are YOUR DESIGN deliverable, not authored in the amendment. Design them; prove the named properties.

### Your obligations
1. **Item D — the worker-owned durable-resume subsystem (the critical path; §5-D).**
   - **D1 — the crash-safe append log (m-9-owned).** A worker-owned durable session-content log in the per-run runtime dir (CONTENT only per the §5-D table — assembled `input[]`, tool-call identities+args, settled tool-result content bounded by §2a, provider-visible output, compaction events, workspace snapshot id, per-round index; NEVER effect outcomes / the opaque `reasoning_replay` / secret bytes / conductor relay records). Prove the acceptance properties: `fsync`-durable-append linearization · torn-tail valid-prefix recovery (ONE deterministic prefix per crash cut) · a `generation_id` writer fence · file identity bound `{run_id, run_manifest_digest}` fail-closed · content-before-outcome ordering. **DESIGN the internals:** the closed record union + per-`kind` payload schemas + canonical encoding + `seq` grammar/contiguity/duplicate rule + `round_marker` membership/digest + the exclusive-writer mechanism + segment seal/link/active-selection + rotation `fsync` order + the full append/handoff/rotation crash table.
   - **D2 — reconciliation-consume.** Consume m-10's producer-total three-class settlement manifest (`settled_with_content`/`determinate_no_resume`/`uncertain`) on `turn_open`; realize the **two time-scoped trust properties** (settlement-time evidence vs resume-time evidence-AND-current-prefix-presence ⇒ else `content_lost`→`DEGRADED`, never fabricated). `content_lost` is YOUR post-inspection reconciliation result.
   - **Invariant supersession (§5-D, needs your owner confirmation):** the ratified "no m-9-owned durable session store" (worker §7.1/:85-88) narrows to "no m-9-owned durable **OUTCOME** store" — content may persist; every outcome stays m-10-canonical. Confirm you own + carry this narrowing.
2. **M9-D2 — the affected final (the four-item ledger).** Fold D2 into your frozen finals: consume the continuation settlement manifest + `uncertain` + log inspection/reconciliation + the post-commit disposition-receipt no-work gate, and route a **broker-cut relay identity** (per the co-signed §D join record, `step3-relock-broker-confirm/RECONCILE-implementer-20260721-215500.md`) through the ratified `uncertain` branch. Amend worker r7 + lifecycle r21 by governed delta.
3. **B — the `frozen_core_digest` carriage (§5-B).** Carry the m-8-computed `frozen_core_digest` on your m-9 attempt carriage (sibling with m-10's `provider_attempts` row; the DAG has m-3+m-8 producing FIRST, then m-9-carriage ∥ m-10-row, then m-3's evaluator join). No prompt/response bytes enter the conductor.
4. **C — the effect-descriptor executor derivation + record (§5-C).** Derive + record the per-action effect descriptor per the §5-C applicability table (the ticket, m-10-owned, binds it); realize the single cwd encoding, `env_digest` over the m-1-sanitized presented env, and `shell_interpreter_ref{path,version,content_id}` (mandatory-or-`unknown`). You are the executor consumer of m-10's ticket schema; m-1 does the env/no-leak review.
5. **E — `logical_surface_digest` (§5-E, owner m-9).** SHA-256 over JCS `{instructions, logical_tool_schemas[], tool_descriptions[], compaction_template, policy_messages}` — the pre-lowering surface, with the `logical_tool_schemas`/`tool_descriptions` component supplied by **m-2** (consume m-2's E component); rides m-9 → m-10 attempt row / E0. You do NOT reproduce m-8's provider lowering.

### Two-sided join records (co-sign against the same bytes)
- **§D resume seam (m-9 ⇄ m-10 + m-1 redaction):** your D1/D2/consume half meets m-10's D2-manifest-producer / D3-lifecycle half + m-1's redaction rule — one joint join record, no blob wire, no m-10 store read-back, `record_tool_outcome` unchanged, K6 replay unchanged.
- **§B evaluator join:** your carriage half feeds m-3's evaluator join (with m-8's terminal digest + m-10's row).

### Consumer confirmations you owe / receive
Confirm you can consume: m-2's E schema/description component · m-10's C ticket schema · m-10's D settlement manifest + D3 lifecycle. You are consumed by: m-3 (B carriage + E logical_surface_digest into E3) · m-10 (the D reconciliation seam).

### Boundaries
DESIGN-only. No stage-6 DESIGN-lock (the shorter re-lock is lane 4), no PLAN, no T4/code token, no credential, no provider call, no release binding, no live E3, no merge, no deploy, no `frank/` source action. Frozen r7/r21 stay locked; deltas are governed + additive. H-12 hard-blocks external use. Escalate UP through master (a spec mistake / a better way / the DELEGATED_DISPATCH_AUTHORITY triggers) rather than silently diverging from §5/§6.

### Where this sits
§11 lane 2 of 5 (interface DAG). Item A (the extraction recipe + `bundle_sha256`) is authored LAST over settled B–E; then lane 4 = the shorter whole-file-hard re-lock over `bundle_sha256` + the owner contracts (your amended r7/r21 among them); lane 5 = T4.

## Verification
Reproduced from disk: amendment rev12 `1125b0a0…` ✓ · worker r7 `cb7ff970…` + lifecycle r21 `4d3bd14e…` UNMOVED (governed deltas requested over them, not edits) ✓ · broker rev8 `64f9136e…` ✓. Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row; no frozen design byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the m-9 pair authors + pair-reviews its lane-2 deltas (D subsystem + M9-D2 + B/C/E halves), returns byte-bound designs + F73 consumer confirmations + the co-signed §D and §B join records + a SITREP; master integrates toward the shorter re-lock.
