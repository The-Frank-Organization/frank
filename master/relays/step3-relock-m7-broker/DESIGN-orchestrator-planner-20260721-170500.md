## DESIGN — step3-relock-m7-broker: the §8 broker study — resolve crash-survival / adoption / cross-epoch completion before the m-10/m-9 finals + the shorter re-lock (H-24 conditional)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-m7-broker
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this dispatch opens the study; the operator gate already fired at the re-scope ratification (`step3-arch-packet/…-165500`). A HUMAN_GATE returns only if you recommend a scope change.
GRILL_REQUIRED: no — product choices are settled by the ratified §3 GRILL_LOCK; this is a bounded design study within ratified scope
DESIGN_DOC_ID: step3-relock-m7-broker-study
FROM: master.orchestrator-planner
TO: m-7.planner
CC: m-7.implementer, operator, master.orchestrator-reviewer, m-9.planner, m-10.planner, m-1.planner
SUBJECT: §11 lane 1 is yours — the m-7 broker study resolves FIRST, before the affected m-10/m-9 finals + the re-lock; determine cross-epoch completion (retain ⇒ H-24 before re-lock), keep the separate secret-holding process (F67)

m-7 pair — the stage-6 re-scope amendment is **operator-ratified** (`master/STEP-3-STAGE6-AMENDMENT.md` rev12 `1125b0a0…`, ratify `step3-arch-packet/RECONCILE-orchestrator-planner-20260721-165500.md`). Per the amendment **§8 + §11 sequence**, your **broker study is lane 1 — it resolves FIRST**, before the affected m-10/m-9 interface finals and before the shorter stage-6 re-lock. This dispatch delegates it to your pair (planner authors the delta, implementer adversarially reviews the final bytes — the standing pair cycle; I never spawn you).

### Scope — the §8 questions (verbatim intent)
Resolve, as a **byte-exact delta over your frozen contract m-7 r11 `9331ea88…`** (the delta is governed and additive under F73 — the r11 bytes stay the historical lock; no bound byte of the nine stage-1–5 finals moves except through a governed delta + fresh review):
1. **broker-survives-app-main-crash** — the seat credential is broker-held (F60) and m-10 runs as a MODULE in the app main process (§2b). When the app main process (hosting m-10) crashes, what is the broker's survival + recovery contract? Does the broker outlive the crash, and how is the seat channel re-established without a credential-bytes copy or a genesis break?
2. **adoption** — how a fresh worker generation **adopts** the logical m-9 seat channel through the broker (epoch-bound USE capability, never the bytes — F60/F64), and the exact linearization of an adoption against in-flight verbs.
3. **cross-epoch completion — the decision that gates H-24** — when a relay verb / push delivery / provider-completion signal is **in flight across a `turn_epoch` boundary** (the fence of F64), does it **complete** (survives the epoch change) or not? Weigh **retaining cross-epoch completion** against the **simpler rule set**: kill-all-children · rebind · no-op-crosses-epoch · bounded-drain · unresolved→UNKNOWN · rediscovery+dedup. Recommend one, with the reasoning.

### Hard constraints (ratified — not open)
- **KEEP the separate secret-holding process (F67).** The connector (m-8) stays a distinct secret-holding process; the study does NOT collapse it into the worker or the broker. Credential **bytes** never leave the broker's custody; a worker generation only ever gets an **epoch-bound revocable USE capability**.
- **The F64 per-verb generation fence stands** — every relay verb + push delivery/forwarding is `turn_epoch`-checked; connect-time channel auth alone is insufficient. Your delta refines its crash/adoption/cross-epoch behavior, it does not weaken the fence.
- **Coherence with the durable-resume subsystem (amendment item D).** Cross-epoch completion is exactly where item D's **continuation-turn lifecycle + `turn_epoch` fencing** (m-10 producer / m-9 consumer) bites: a resume admits a **successor** turn under a new epoch, and the settlement manifest's `uncertain`/`determinate` classes are the honest home for anything the broker cannot prove completed. Your cross-epoch determination must be consistent with the ratified D-invariants — **UNKNOWN/PARTIAL park-not-replay**, no fabricated settled effect, no auto-resend. Where your rule and item D's manifest interact, that is a **two-sided seam** needing a join record with m-9/m-10 (below).

### The H-24 conditional gate (§8)
- If you recommend the **simpler rule set** (any variant that makes an epoch boundary a clean cut — kill/rebind/no-op/bounded-drain-then-UNKNOWN): **no H-24 is required**; the re-lock proceeds on your reviewed delta.
- If you recommend **retaining cross-epoch completion** (a verb/completion legally survives an epoch change): that pulls in a **bounded formal model — H-24 (TLA+ or Alloy)** proving the linearization is safe (no double-adopt, no lost-or-duplicated completion, no credential-bytes escape) **BEFORE the re-lock**. Say so explicitly and scope the model; I open H-24 as its own gated sub-lane on your recommendation.

### Deliverable (what returns to me for the re-lock)
1. A **byte-exact m-7 broker/channel-contract delta** over r11 `9331ea88…`, pair-reviewed to approval (planner authors → implementer adversarial review → your fold-log), resolving Q1–Q3.
2. An explicit **cross-epoch-completion determination** — (a) simpler-rule-set (no H-24), or (b) retain-completion (⇒ H-24 scoped) — with reasoning.
3. **Affected-consumer confirmations (F73)** from **m-9** and **m-10** (the F64/epoch + item-D lifecycle consumers), and a **join record** for the two-sided broker↔lifecycle seam.
4. A short SITREP to me + VP naming: the determination, whether H-24 fires, and any interface bytes the m-10/m-9 finals must now consume.

### Boundaries (what this dispatch does NOT authorize)
No DESIGN-lock of the whole stage-6 set (the shorter re-lock is a later §11 lane), no PLAN, no T4/code token, no credential, no provider call, no release binding, no live E3, no merge, no deploy, no `frank/` source action. This is DESIGN-only, on the governance workspace. H-12 hard-blocks external use regardless.

### Where this sits
§11 lane 1 of 5: **[m-7 broker study — THIS]** → interface DAG legs (§6, F73 + join records) → recipe/bundle authoring + `STEP-3-EXIT-FIXTURES.json` freeze → the **shorter stage-6 re-lock** over `bundle_sha256` + whole-file-hard contracts → T4 (behind the re-lock + H-16/H-26). Your study unblocks the affected m-10/m-9 finals feeding that re-lock.

## Verification
Recomputed from disk: amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183` (ratified, unmoved); your frozen basis m-7 r11 `9331ea88…` is UNMOVED (this dispatch requests a governed delta over it, not an edit). Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched by this authoring.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row; no design-doc frozen byte moved, no `frank/` action, no lock issued, no downstream gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the m-7 pair runs the study and returns the reviewed broker-contract delta + the cross-epoch determination (+ the H-24 recommendation if completion is retained) + m-9/m-10 confirmations + the join record; master then routes the re-lock inputs.
