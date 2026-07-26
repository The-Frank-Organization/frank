## DESIGN — step3-relock-broker-confirm: the m-7 broker study is PAIR-APPROVED at rev8 `64f9136e…` with the SIMPLER RULE SET (NO H-24) — m-9 + m-10, confirm your §C consumer scopes on the exact bytes and co-sign the §D join record

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-broker-confirm
PARENT_DISPATCH_ID: step3-relock-m7-broker
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the study recommends NO scope change; the operator gate fired at the re-scope ratification. A HUMAN_GATE returns only if a confirmation surfaces a genuine cross-owner conflict I must arbitrate.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-broker-study
FROM: master.orchestrator-planner
TO: m-9.planner, m-10.planner
CC: m-7.planner, m-7.implementer, master.orchestrator-reviewer, operator, m-1.planner
SUBJECT: byte-bound F73 consumer-confirmation of the broker-study delta rev8 `64f9136e…` (additive over m-7 r11 `9331ea88…`) + the two-sided §D join record — on your confirmations the affected m-10/m-9 finals + the shorter re-lock proceed on this reviewed delta, no H-24

m-9 + m-10 pairs — §11 lane 1 (the m-7 broker study) has **resolved and pair-approved**: `master/domains/m-7-conductor-core/design/2026-07-21-step3-broker-study-crash-adoption-epoch.md` **rev8 @ SHA-256 `64f9136e0b851e31f129372ec50b5667c2b6dcda197ef0dcf95c4f0eca5ff4ce`** (m-7.implementer approve `step3-relock-m7-broker/DESIGN-REVIEW-implementer-20260721-205236.md`, byte-bound; a governed **additive delta over m-7's UNMOVED r11 `9331ea88…`** under F73 — the superseded r11 clauses named byte-exactly in the study's §Q3.3; **any byte change voids the approval**). Master has verified the artifact hash, the approval verdict, and the frozen bases (r11 + amendment rev12 `1125b0a0…`) — all reproduce.

**The determination (carry it as settled): the SIMPLER RULE SET.** Cross-epoch completion is **NOT retained** — the epoch boundary is a clean cut (old-E authority stops at PROPOSED; bounded drain to a fixed broker-local deadline that runs through control loss; unresolved operations cut typed + honest; install proceeds locally even mid-outage). A cut relay call's **only** identity/disposition carrier is m-10's **existing F59 outcome machinery** (terminal-stays-terminal · consumed-no-outcome ⇒ `UNKNOWN_TOOL_OUTCOME` · issued-unconsumed ⇒ `VOID` · `parked_unknown` + the item-D D2 `uncertain` class — consumed, **not** modified; ONE outcome carrier, no second settlement path). Because nothing survives the boundary to be lost or duplicated and custody is untouched, the safety properties hold **by construction** ⇒ **NO H-24 formal model fires**; the re-lock proceeds on this reviewed delta once your confirmations land.

You are each a **consumer** of this delta (F73). Confirm on the **exact rev8 bytes** — run your normal pair cycle (planner assesses, implementer adversarially checks) and return a **byte-bound confirmation** (confirm, or a specific revise-request if a scope below is not consumable as written).

### m-10 — your §C consumer scope (confirm you can consume, byte-exact)
- The **CI-3 SHRINK** — crossing rows + transition-ledger DROP; `broker_events` + ack unchanged; the event table **loses** `crossing_set`/`cross_epoch_completion`/`transition_aborted`, **gains** the uncoupled `boundary_cut`; `epoch_installed` **amended** (`crossing_count` removed).
- The **simplified transition protocol** — PROPOSED → PREPARING(bounded drain) → INSTALLED; recovery = re-proposal.
- The exact **`state_proposal` / `state_proposal_result` wire pair** — the correlation-boundary staging, the five-member disposition enum, the ORDERED eight-row table, the validation rules.
- The **two-form assign gate** keyed by the canonical tuple `{run_id, generation_id, turn_epoch, state_seq}` (proof = idempotent evidence; assignment uniqueness stays lifecycle-owned by your durable rows).
- The **full-tuple comparison ladder** (same-epoch newer-state = ordinary install).
- **CI-4** — the spawn discipline (own process group; no parent-death kill of any form; the app-main-crash survival property you supervise).
- The **express binding of the `boundary_cut`** to your §B.3/§B.4 `UNKNOWN`/`VOID`/`parked_unknown`/manifest path.

### m-9 — your §C consumer scope (confirm you can consume, byte-exact)
- The **worker-visible boundary semantics** — typed `broker:unknown-outcome` / `broker:stale-epoch` when the connection lives; **informed rediscovery** over the DISCLOSED parked identity; any re-invocation a **FRESH F59 ticket — never auto-resend**; relay/local uniformity unchanged; `Describe` effect-free + freely retryable.
- The **attach taxonomy unchanged** (FX-TB-19 carries forward whole).
- Your **reading of the §D join record** (below).

### The §D join record — two-sided, m-9 + m-10 CO-SIGN against the same rev8 bytes
The **broker-boundary ↔ continuation-lifecycle seam:** a consumed-no-outcome relay call surfaces to the **successor** turn as the disclosed **`uncertain`** identity; reconciliation is **informed rediscovery**; re-invocation is a **fresh ticket, never automatic**. This is exactly the item-D resume seam (m-10 settlement-manifest producer ⇄ m-9 reconciliation-consumer) meeting the broker's clean-cut disposition — the two must agree in one co-signed record. Produce it jointly (both pairs sign against `64f9136e…`); it is the durable artifact that binds your two halves.

### Hard constraints (already met by the study; confirm they hold in your consumption)
F67 (the secret-holding set stays `{m-8, broker}`; placement untouched) · F64 (the per-verb fence unweakened — the study only disposed of already-sent stragglers) · F60/F66 (capability-not-bytes; epoch-fenced replacement) · item-D coherence (UNKNOWN/PARTIAL **park-not-replay**, no fabricated settled effect, no auto-resend). If any of these does NOT hold as you consume the delta, that is a revise-request, not a silent accept.

### What returns to me
1. **m-10:** a byte-bound F73 confirmation on rev8 `64f9136e…` over the m-10 §C scope.
2. **m-9:** a byte-bound F73 confirmation on rev8 `64f9136e…` over the m-9 §C scope.
3. **Jointly:** the co-signed **§D join record** (both FROMs on their own records per channel-stamping; the join record references the same bytes). On all three, master carries the NO-H-24 determination into the §11 sequence and the affected m-10/m-9 finals + the shorter re-lock proceed.

### Boundaries (this dispatch does NOT authorize)
No DESIGN-lock of the stage-6 set (the shorter re-lock is a later §11 lane), no PLAN, no T4/code token, no credential, no provider call, no release binding, no live E3, no merge, no deploy, no `frank/` source action. DESIGN-only on the governance workspace. H-12 hard-blocks external use regardless. m-7's r11 stays the frozen historical lock with the delta governed over it; m-1's edges are untouched (named for completeness — no m-1 confirmation owed unless a conflict surfaces).

## Verification
Reproduced from disk this session: study artifact `64f9136e0b851e31…` ✓ · m-7 r11 `9331ea88…` UNMOVED ✓ · amendment rev12 `1125b0a0…` UNMOVED ✓ · the r8 approval relay present with `DESIGN_REVIEW_VERDICT: approve` byte-bound to rev8 ✓ · the SITREP `205544` exact-file lint OK. Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched by this authoring.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row; no design-doc frozen byte moved, no `frank/` action, no lock issued, no downstream gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the m-9 and m-10 pairs return byte-bound F73 confirmations on rev8 `64f9136e…` + the co-signed §D join record; master then carries the NO-H-24 determination into the §11 sequence and opens the affected m-10/m-9 finals feeding the shorter re-lock.
