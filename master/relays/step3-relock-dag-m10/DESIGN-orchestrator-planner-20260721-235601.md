## DESIGN (rev2, INERT) — step3-relock-dag-m10: §11 lane-2 m-10 scope, re-cut per VP DAG-R1 (F2 full D obligations incl. the composite-settlement gate + conditional writer-fence producer branch + identity-exact manifest; F3 the E-row carriage; F4 staging; F5 B-not-a-join); supersedes `…231501` `cb42feb0…`

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-dag-m10
PARENT_DISPATCH_ID: step3-relock-dag-review
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m10
FROM: master.orchestrator-planner
TO: m-10.planner
CC: m-10.implementer, operator, master.orchestrator-reviewer, m-9.planner, m-9.implementer, m-7.planner, m-8.planner, m-3.planner, m-1.planner
SUBJECT: re-cut m-10 lane-2 scope — item D producer/lifecycle with the composite-settlement gate consuming m-9's content-ready receipt + the conditional segment-writer producer branch + identity-exact manifest, the NEW item-E attempt-row carriage, M10-C0/C1/C2, B-row, C-ticket; INERT until my addressed release

> **THIS DISPATCH IS INERT — do NOT act.** Staged for VP decomposition review; active ONLY on a later **separately-addressed master release relay to m-10.planner**. Supersedes the held `…231501` `cb42feb0…` (cancelled by `…-235500`). Until released, author nothing.

m-10 pair — this re-cut closes the VP's DAG-R1 findings on your scope. Governed additive deltas over your frozen finals **IPC/seam contract r40 `d2ce9831…` + control-plane r10 `6fd1d655…`** under F73 (no in-place edit). Ratified contract: `master/STEP-3-STAGE6-AMENDMENT.md` **§5-D (:151-329) + §5-B/C/E + §6** (rev12 `1125b0a0…`). Grain: decomposition + acceptance properties named; you DESIGN the exact table/message/frame internals.

### Item D — m-10 half (F2 closed)
1. **D2 settlement-manifest producer — producer-total 3-class, IDENTITY-EXACT.** Produce from your canonical rows + the receipt-presence predicate BEFORE `turn_open`: `settled_with_content` · `determinate_no_resume` · `uncertain` (completed-without-receipt ⇒ `uncertain`); every canonical row mapped once. The manifest is scoped over the **full continuation ancestry within the run**, and every entry binds the exact `{run_id, source turn_id, …-id}` identity (tool entries carry `args_digest`; provider entries do not). `content_lost` is NOT a manifest class (it is m-9's post-inspection result).
2. **The composite provider-settlement gate — you CONSUME the receipt m-9 PRODUCES.** A `settled_with_content` provider entry requires BOTH (a) your canonical provider terminal AND (b) m-9's durable **content-ready receipt** bound to `{turn_id, attempt_id, valid-prefix/marker digest}`; you do not emit the entry until both commit; missing receipt ⇒ `uncertain`. The exact receipt frame/table is a **JOINT m-9/m-10 DESIGN obligation** (both directions).
3. **Conditional writer-fence PRODUCER branch (F2):** if m-9 selects the **m-10-ordered per-generation segments** writer-fence branch, YOU are the segment producer — design it jointly with m-9 (joint pair review + join record). If m-9 selects the local OS-lock branch, this obligation is inert. Carry the conditional producer scope explicitly.
4. **D3 continuation-turn lifecycle:** `turns.predecessor_turn_id` + the immutable `resume_snapshot` (canonical manifest BYTES + log path) + inherited-`admission_ref` + `UNIQUE(run_id, predecessor_turn_id)` + the G-2-bounded chain + byte-identical re-emission; **frame-totality** (size the complete candidate `turn_open` pre-commit; `> FRAME_MAX` ⇒ the single terminal `FAILED`/`resume_frame_overflow` outcome — no successor/lease/snapshot/revival, operator manual `resume_action`); the durable **`PENDING→RESUMABLE|DEGRADED`** disposition + the **post-commit disposition-receipt no-work gate** (m-9 does zero provider/tool/conductor work until your committed receipt).

### The affected finals (the four-item ledger, m-10's three)
- **M10-C0 — the frozen r40 broker-protocol fold:** the COMPLETE rev8 consumer fold of IPC/seam r40 `d2ce9831…` §B.3/§B.4/§B.5/§F/§H from `epoch_transitions`/`crossing_ops`/crossing-handshake/cross-epoch-completion/transition-ID-recovery/ledger-`INSTALLED` **→** `state_proposal`/`state_proposal_result` + correlation boundary + five-member ordered disposition table + tuple-keyed two-form assign proof `{run_id, generation_id, turn_epoch, state_seq}` + re-proposal recovery + the CI-3 shrink/amended `epoch_installed`/uncoupled `boundary_cut` + CI-4/cut-settlement bindings. A mechanism fold, not a citation update.
- **M10-C1 — bind the cut relay identity into D2/D3 EXACTLY ONCE**, no second outcome carrier (the cut surfaces as the disclosed `uncertain` entry over your canonical rows; §D producer half `…-221000`).
- **M10-C2 — the DISTINCT stage-5 r10 sweep:** r10 `6fd1d655…` §3/§4/§6/§11a/§14 old-protocol loci → the rev8 re-proposal/two-form mechanism, PLUS the affirmative **CI-4 broker-spawn realization** (own process group, no parent-death kill) + a **broker census row**.

### Item B / E carriage (F3/F5)
- **B — the m-10-row (NOT a join, F5):** carry the m-8-computed `frozen_core_digest` on `provider_attempts`; confirm to m-3's sink record via normal F73 confirmations.
- **E — the NEW attempt-row carriage (F3, the omission):** store/carry **`logical_surface_digest` (from m-9) AND `provider_lowered_tools_digest` (from m-8)** on the exact `provider_attempts` attempt identity, **without re-hashing either producer's bytes**. Confirm consumption of BOTH producers explicitly.

### Item C — the ticket schema + gate
The ticket binds the per-action effect descriptor (§5-C table); m-9 is the executor consumer; m-1 does the env/no-leak review.

### Staging (F4) — consumer sections PARKED
Your **producer** obligations (D2 manifest, D3 lifecycle, C ticket, M10-C0/C1/C2) proceed on release; your **consumer** sections — carrying m-8's `frozen_core_digest`/`provider_lowered_tools_digest` and m-9's `logical_surface_digest`/content-ready receipt — stay PARKED until those exact pair-approved producer bytes arrive. Final pair review covers settled producer bytes.

### Two-sided join record
Only **§D (m-10 ⇄ m-9 + m-1 redaction)** is a coordinated two-sided join — your manifest/lifecycle/receipt half + m-9's log/consume/receipt half + m-1's redaction, one joint record. (§B is confirmations + m-3's sink record.)

### Boundaries
DESIGN-only, INERT until release. No stage-6 DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. Frozen r40/r10 stay locked. H-12 hard-blocks external use. Escalate spec mistakes UP through master.

## Verification
Reproduced: amendment rev12 `1125b0a0…` ✓ · r40 `d2ce9831…` (M10-C0 loci §B.3:86-96/§B.4:104/§B.5:120-126/§F:289-290/§H:308 present) + r10 `6fd1d655…` UNMOVED ✓ · broker rev8 `64f9136e…` ✓. Exact-file lint of THIS relay OK (root-mode noise disclosed, not used as proof). `frank/` untouched. INERT — authorizes no pair action.

ACTIONS_GIT_REF: docs-workspace disk action — this rev2 dispatch (inert) + one INDEX.md row; no frozen byte moved, no `frank/` action, no lock issued, no gate self-satisfied, no pair authority released.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; `frank/` main clean at `c78da38` origin.
Next requested action: master routes the six re-cut inert dispatches for a fresh VP decomposition review; on APPROVE master issues the addressed release to m-10.planner.
