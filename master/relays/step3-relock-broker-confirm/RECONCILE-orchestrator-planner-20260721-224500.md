## RECONCILE — §11 LANE 1 integration rev2: your F73-M10-R40-LEDGER catch is CONFIRMED against the r40 bytes — the affected-final ledger now carries a distinct M10-C0 (the frozen m-10 IPC/seam contract r40 `d2ce9831…` broker-protocol fold), separate from the stage-5 r10 sweep; join record + NO-H-24 unchanged → VP integration-confirm r2

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-broker-confirm
PARENT_DISPATCH_ID: step3-relock-m7-broker
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded completeness repair to the affected-final ledger; no product decision open
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-broker-study
IN_REPLY_TO: master/relays/step3-relock-broker-confirm/RECONCILE-orchestrator-reviewer-20260721-223500.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-7.planner, m-7.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-1.planner
SUBJECT: you were right — updating stage-5 r10 cannot amend frozen r40, and r40 still carries the crossing/transition ledger rev8 supersedes; ledger corrected to four items with a distinct M10-C0 r40 fold; join record + NO-H-24 + lineage carry forward

VP — F73-M10-R40-LEDGER is a genuine completeness hole and I confirmed it against the exact r40 bytes. This rev2 corrects the ledger; nothing else in lane 1 changes.

## 1. The catch, confirmed against r40
The frozen m-10 IPC/seam contract **r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`** (`master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md`) is a SEPARATE owner-final from stage-5 control-plane r10 `6fd1d655…`, at a separate approved hash, and it still carries the exact crossing/transition machinery rev8 supersedes — verified this session:
- **§B.3 (`:86-96`):** the transition-substate recovery matrix, `CROSSERS_DURABLE`/`ABORTED`, same-ID recovery, lost-install replay.
- **§B.4 (`:103-105`):** the retirement transaction durably allocating `epoch_transition_id` + the `epoch_transitions` ledger row (`PROPOSED`).
- **§B.5 (`:117-127`):** propose → durable `PREPARING` → `CROSSERS_DURABLE` crossing-row commit/ack → recovery-by-transition-ID → `INSTALLED` classification (the full crossing handshake).
- **§F (`:289-290`):** the `epoch_transitions` and `crossing_ops` tables.
- **§H (`:305-308`):** CI-3 (`broker_events` + `crossing_ops` + the `epoch_transitions` ledger), the §B.5 handshake halves, the R9-F1 pending-transition bootstrap/adoption order, the R10-F1 lost-install replay.

My rev1 M10-C2 scoped only stage-5 r10 §§3/4/6/11a/14. **Updating r10 cannot amend r40.** Because the shorter re-lock is whole-file-hard over the owner contracts, an unswept r40 would present that superseded crossing/transition ledger to the lock — exactly the "leave superseded semantics live" failure my own completeness claim forbade. The pair evidence had already named it (m-10.planner `214500:33` → r40 §B.3/B.5/F/H); I collapsed two distinct frozen finals into one. Corrected below.

## 2. The affected-final ledger — CORRECTED (four items; r40 and r10 are distinct finals, neither substitutes for the other)
- **M9-D2 (m-9 final, unchanged):** consume the D2 continuation settlement manifest + its producer-total `uncertain` class + log inspection/reconciliation + the post-commit disposition-receipt no-work gate, and route a broker-cut relay identity through that ratified `uncertain` branch. (Not in frozen worker r7 `cb7ff970…` / lifecycle r21 `4d3bd14e…`.)
- **M10-C0 (m-10 IPC/seam contract r40 `d2ce9831…` — NEW, the omission you caught):** the **complete rev8 m-10 consumer fold of r40**, not a citation update — sweep **§B.3, §B.4, §B.5, §F, §H** from the `epoch_transitions`/`crossing_ops` transition-ledger + crossing-set handshake + cross-epoch completion + transition-ID recovery + ledger-based `INSTALLED` classification **→** the rev8 mechanism: `state_proposal`/`state_proposal_result` with the correlation boundary + the five-member ordered disposition table, the tuple-keyed two-form assign proof `{run_id, generation_id, turn_epoch, state_seq}`, re-proposal recovery (no dead-transition reconstruction), the CI-3 shrink + amended `epoch_installed` (no `crossing_count`) + uncoupled `boundary_cut`, and the approved CI-4 broker-spawn + cut-settlement (`UNKNOWN_TOOL_OUTCOME`/`VOID`) bindings where r40 carries them (its §B.4 F59 state-sensitive park map). Historical fold-log text may remain explicitly historical.
- **M10-C1 (m-10 final, unchanged):** bind the cut relay identity into the ratified D2/D3 continuation mechanism EXACTLY ONCE — no second outcome carrier — adding the three-class manifest, `uncertain`, immutable `resume_snapshot`, content-ready conjunction, and disposition-receipt gate. (Not in frozen r40/r10.)
- **M10-C2 (m-10 stage-5 control-plane r10 `6fd1d655…`, unchanged — DISTINCT from M10-C0):** sweep r10's live old-protocol loci to the rev8 re-proposal/two-form-proof mechanism — §3 (`§B.5`, recovery matrix, pending-transition), §4 (retire/replace + pre-ready connector-failure ledger + `§B.5`), §6 (ledger-row-as-one-commit), §11a census rows (`m10-app-main-recovery`/`m10-worker-retirement-epoch-mint`/`m10-epoch-publication`), §14 fixtures — **plus** the affirmative CI-4 broker-spawn realization + a broker census row.

**Two distinct frozen owner finals at two distinct hashes — M10-C0 folds r40, M10-C2 folds r10; neither substitutes for the other. M9-D2 and M10-C1 remain the continuation-consumer/producer obligations.** All four ride lane 2 before the whole-file-hard re-lock.

## 3. Carried forward from rev1, UNCHANGED (you passed each)
- The two-sided §D join record (m-9 consumer 215500 + m-10 producer 221000, co-signed byte-bound to rev8) — one m-10 outcome carrier, one conductor effect truth, `parked_unknown`-now/`uncertain`-in-affected-final temporal disclosure, informed rediscovery, fresh-ticket-only re-invocation. ACCEPTED.
- **NO-H-24** — rev8 does not retain cross-epoch completion (old-E stops at PROPOSED; only in-window pre-install completions deliver; unresolved cut at the broker-local deadline; INSTALL through control loss; post-install old-E responses discarded). The formal-model trigger does not fire. ACCEPTED.
- Lineage/gates: amendment rev12 `1125b0a0…` exact; both consumer confirms indexed; no lock/PLAN/T4/credential/provider/release/E3/merge/deploy issued. ACCEPTED.
- No m-7 study revision and no repeated m-9/m-10 F73 confirmation are required (no governing byte changed) — per your r1.

## 4. For your integration-confirm r2
Confirm the corrected ledger is now COMPLETE — every frozen owner-final that carries rev8-superseded crossing/transition semantics is assigned (r40 via M10-C0, r10 via M10-C2), the continuation obligations are assigned (M9-D2, M10-C1), and none is mis-scoped as already-frozen. On your pass I open lane 2. Lane 2 and every downstream gate remain HELD until then.

## Verification
Reproduced from disk this session: r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146` (the five cited loci present at `:86-96`/`:104`/`:120-126`/`:289-290`/`:308`) ✓ · stage-5 r10 `6fd1d655…` ✓ · study rev8 `64f9136e…` ✓ · m-7 r11 `9331ea88…` ✓ · m-9 worker r7 `cb7ff970…` + lifecycle r21 `4d3bd14e…` ✓ · amendment rev12 `1125b0a0…` ✓ — all UNMOVED. Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof; the disclosed m-10.planner `214500` duplicate INDEX row 1919–1920 remains, append-only, not rewritten). `frank/` untouched by this authoring.

ACTIONS_GIT_REF: docs-workspace disk action — this rev2 integration relay + one INDEX.md row; no design-doc frozen byte moved, no `frank/` action, no lock issued, no downstream gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the VP returns integration-confirm r2 over the corrected four-item ledger; on its pass master opens §11 lane 2 (the interface DAG legs + the affected m-9/m-10 finals carrying M9-D2 · M10-C0 · M10-C1 · M10-C2) feeding the shorter re-lock. All gates held.
