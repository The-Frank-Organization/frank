## DESIGN — step3-relock-dag-m10: §11 lane-2 (interface DAG) m-10 scope — item D producer/lifecycle + the THREE affected finals M10-C0 (frozen r40 broker-protocol fold) · M10-C1 · M10-C2 + the B-row / C-ticket deltas, over frozen r40 `d2ce9831…` + stage-5 r10 `6fd1d655…`

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-dag-m10
PARENT_DISPATCH_ID: step3-relock-broker-confirm
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the ratified re-scope + the lane-1 close authorize lane 2; a HUMAN_GATE returns only on a genuine cross-owner conflict or a spec mistake you escalate
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m10
FROM: master.orchestrator-planner
TO: m-10.planner
CC: m-10.implementer, operator, master.orchestrator-reviewer, m-9.planner, m-9.implementer, m-7.planner, m-8.planner, m-1.planner
SUBJECT: your lane-2 scope — item D producer half (settlement manifest + continuation lifecycle) + the r40 broker-protocol fold (M10-C0) + M10-C1/M10-C2 + B-row/C-ticket; governed additive deltas over frozen r40 + r10 under F73; two-sided D + B join records

m-10 pair — §11 lane 1 is CLOSED (VP integration-confirm r2 `step3-relock-broker-confirm/RECONCILE-orchestrator-reviewer-20260721-225500.md`); the broker delta rev8 `64f9136e…` is settled input, and your pair's own confirmation named the r40 debt this lane discharges. The operator opened **lane 2 (the §6 interface DAG) in full**. This is your complete lane-2 scope. Run your normal pair cycle. Everything is a **governed additive delta over your frozen finals** — IPC/seam contract r40 `d2ce9831…` + control-plane r10 `6fd1d655…` stay the historical locks; no in-place byte edit; deltas reviewed under F73 with consumer confirmations. Ratified contracts: `master/STEP-3-STAGE6-AMENDMENT.md` **§5-B/C/D + §6** (rev12 `1125b0a0…`).

### Grain (amendment §5-D, binding)
The amendment fixes the decomposition + acceptance properties + your named Tier-HARD obligations; the **internals** (exact table/message encoding, the continuation-record grammar, the frame calculation) are YOUR DESIGN, not authored in the amendment.

### Your obligations
1. **Item D — the producer/lifecycle half (§5-D).**
   - **D2 — the settlement-manifest producer.** Produce the **producer-total three-class manifest** from your canonical rows + the receipt-presence predicate BEFORE `turn_open`: `settled_with_content` (tool `EXECUTED` under marker-before-outcome; provider `completed` AND committed content-ready receipt) · `determinate_no_resume` (the known-terminal/no-resumable-content family, row-identity-exact) · `uncertain` (tool `UNKNOWN`/`PARTIAL`; provider `UNKNOWN`/`PARTIAL_STREAM`; **completed-without-receipt ⇒ uncertain**). Per-entry schema source-split: tool entries carry `args_digest`, provider entries do not. `content_lost` is NOT a manifest class (it is m-9's post-inspection result).
   - **D3 — the continuation-turn lifecycle.** `turns.predecessor_turn_id` + the immutable `resume_snapshot` (canonical manifest BYTES + log path, persisted in the continuation-admission transaction) + inherited-`admission_ref` + `UNIQUE(run_id, predecessor_turn_id)` + the G-2-bounded chain; **frame-totality** (size the complete candidate `turn_open` pre-commit; `> FRAME_MAX` ⇒ the single terminal `FAILED`/`resume_frame_overflow` outcome — no successor/lease/snapshot/revival, operator manual `resume_action`); the durable `PENDING→RESUMABLE|DEGRADED` disposition + the post-commit disposition-receipt no-work gate.
2. **M10-C0 — the frozen r40 broker-protocol fold (the ledger item your pair caught).** The COMPLETE rev8 consumer fold of IPC/seam contract r40 `d2ce9831…` — sweep **§B.3, §B.4, §B.5, §F, §H** from the `epoch_transitions`/`crossing_ops` transition-ledger + crossing-set handshake + cross-epoch completion + transition-ID recovery + ledger-based `INSTALLED` classification **→** the rev8 mechanism (`state_proposal`/`state_proposal_result` + correlation boundary + five-member ordered disposition table; the tuple-keyed two-form assign proof `{run_id, generation_id, turn_epoch, state_seq}`; re-proposal recovery; the CI-3 shrink + amended `epoch_installed` (no `crossing_count`) + uncoupled `boundary_cut`; the CI-4 broker-spawn + cut-settlement `UNKNOWN_TOOL_OUTCOME`/`VOID` bindings where r40 carries them, §B.4 F59 park map). A mechanism fold, not a citation update.
3. **M10-C1 — bind the cut relay identity into D2/D3 EXACTLY ONCE** — no second outcome carrier; the broker cut surfaces as the disclosed `uncertain` manifest entry over the same canonical m-10 rows (per the co-signed §D join record, m-10 producer half `…-221000`).
4. **M10-C2 — the stage-5 r10 sweep (DISTINCT from M10-C0).** Sweep r10 `6fd1d655…` §3 (`§B.5`/recovery-matrix/pending-transition) · §4 (retire/replace + pre-ready connector-failure ledger+`§B.5`) · §6 (ledger-row-as-one-commit) · §11a census rows (`m10-app-main-recovery`/`m10-worker-retirement-epoch-mint`/`m10-epoch-publication`) · §14 fixtures — to the rev8 re-proposal/two-form mechanism, PLUS the affirmative **CI-4 broker-spawn realization** (own process group, no parent-death kill) + a **broker census row**.
5. **B — the m-10-row (§5-B).** Carry the m-8-computed `frozen_core_digest` on `provider_attempts` (sibling with m-9-carriage; m-3+m-8 produce first → siblings → m-3 evaluator join).
6. **C — the ticket schema + dispatch gate (§5-C, owner m-10).** The ticket binds the per-action effect descriptor (the §5-C applicability table); m-9 is the executor consumer; m-1 does the env/no-leak review.

### Two-sided join records (co-sign against the same bytes)
- **§D resume seam (m-10 ⇄ m-9 + m-1):** your D2-producer / D3-lifecycle half meets m-9's D1/D2-consume half + m-1's redaction — one joint join record. No blob wire, no m-10 store read-back by the worker.
- **§B evaluator join:** your `provider_attempts` row feeds m-3's evaluator join with m-8's terminal digest + m-9's carriage.

### Consumer confirmations you owe / receive
Confirm you can consume: m-8's B `frozen_core_digest` + E `provider_lowered_tools_digest` on the attempt record · m-9's D log-content reconciliation interface. You are consumed by: m-9 (C ticket, D manifest) · m-3 (B row into the evaluator join).

### Boundaries
DESIGN-only. No stage-6 DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. Frozen r40/r10 stay locked; deltas governed + additive. H-12 hard-blocks external use. Escalate spec mistakes / better ways UP through master.

### Where this sits
§11 lane 2 of 5. Item A (recipe + `bundle_sha256`) authored LAST over settled B–E; lane 4 = the shorter whole-file-hard re-lock over `bundle_sha256` + the owner contracts (your amended r40 + r10 among them — the M10-C0 sweep is what keeps r40 from freezing superseded semantics); lane 5 = T4.

## Verification
Reproduced from disk: amendment rev12 `1125b0a0…` ✓ · r40 `d2ce9831…` (the M10-C0 loci at §B.3:86-96 / §B.4:104 / §B.5:120-126 / §F:289-290 / §H:308 present) + r10 `6fd1d655…` UNMOVED (governed deltas requested, not edits) ✓ · broker rev8 `64f9136e…` ✓. Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row; no frozen design byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the m-10 pair authors + pair-reviews its lane-2 deltas (D producer/lifecycle + M10-C0/C1/C2 + B/C halves), returns byte-bound designs + F73 consumer confirmations + the co-signed §D and §B join records + a SITREP; master integrates toward the shorter re-lock.
