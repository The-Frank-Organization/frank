## DESIGN — step3-relock-dag-m8: §11 lane-2 (interface DAG) m-8 scope — item B (`frozen_core_digest` producer on the m-8 terminal) + item E (`provider_lowered_tools_digest` producer), over frozen m-8 r12 `4b670a79…`

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-dag-m8
PARENT_DISPATCH_ID: step3-relock-broker-confirm
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m8
FROM: master.orchestrator-planner
TO: m-8.planner
CC: m-8.implementer, operator, master.orchestrator-reviewer, m-3.planner, m-9.planner, m-10.planner
SUBJECT: your lane-2 scope — the two producer-FIRST digests: B frozen_core_digest on the m-8 terminal + E provider_lowered_tools_digest; governed additive delta over frozen m-8 r12 under F73; you produce first, m-9/m-10 carry, m-3 joins

m-8 pair — §11 lane 1 is CLOSED; the operator opened **lane 2 (the §6 interface DAG) in full**. This is your complete lane-2 scope — and per §6 you are a **producer-FIRST** on both B and E (m-9/m-10 carry your digests; m-3 joins them). Run your normal pair cycle. A **governed additive delta over your frozen final m-8 r12 `4b670a79…`** (stays the historical lock; no in-place edit; reviewed under F73 with consumer confirmations). Ratified contracts: `master/STEP-3-STAGE6-AMENDMENT.md` **§5-B, §5-E + §6** (rev12 `1125b0a0…`).

### Your obligations
1. **Item B — the `frozen_core_digest` producer (§5-B).** You compute the `frozen_core_digest` (you have the frozen-core body bytes at freeze) and ride it on the **m-8 terminal**, alongside the E digest below. It then rides m-10 `provider_attempts` + m-9 carriage + m-3's E0/E3 + the composite exit proof; the observer derives it independently. No prompt/response bytes enter the conductor. Per §6, B is **m-3 + m-8 FIRST → m-9-carriage ∥ m-10-row → m-3 evaluator join** — you and m-3 are the producers the rest consume.
2. **Item E — the `provider_lowered_tools_digest` producer (§5-E, owner m-8).** SHA-256 over the **lowered `tools[]` portion of the frozen-core body** — **m-8 alone performs provider lowering and has the bytes at freeze**. Rides the m-8 terminal / attempt record alongside `frozen_core_digest`. **m-9 does NOT reproduce your translation** — you are the sole owner of the lowered-tools digest; the observer derives it from the observed wire request; m-3 joins it with m-9's `logical_surface_digest` into `model_surface_digest`.

### Consumer confirmations you owe / receive
You are consumed by: m-10 (B `frozen_core_digest` on `provider_attempts`) · m-9 (B carriage) · m-3 (B E0/E3 + E `model_surface_digest` join). Confirm your two digests' recipes + field placement are consumable as those consumers carry/join them (respond to their confirmation asks). You consume nothing new here — you are the producer root.

### Hard constraints (already yours by the reframe — hold them)
The provider wire + credentials + `LLMRequest` stay app-side, never transiting the conductor; you remain the last pre-wire enforcement host (`freeze→authorize→attach→send`); F67 (you stay the separate secret-holding process). The B/E digests are evidence fields on your existing terminal/attempt records — no new provider round-trip, no credential exposure.

### Boundaries
DESIGN-only. No stage-6 DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. Frozen m-8 r12 stays locked; the delta is governed + additive. H-12 hard-blocks external use. Escalate spec mistakes UP through master.

### Where this sits
§11 lane 2 of 5. You are the producer root for B and E; m-9/m-10 carriage + m-3's join depend on your digests, so your delta unblocks their halves. Item A (recipe + `bundle_sha256`) authored LAST; lane 4 = the shorter re-lock (your amended r12 among the whole-file-hard contracts); lane 5 = T4.

## Verification
Reproduced from disk: amendment rev12 `1125b0a0…` ✓ · m-8 r12 `4b670a79…` UNMOVED (governed delta requested, not an edit) ✓. Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row; no frozen design byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the m-8 pair authors + pair-reviews its lane-2 delta (B frozen_core_digest producer + E provider_lowered_tools_digest producer), returns the byte-bound design + F73 consumer-confirmation responses + a SITREP; master integrates toward the shorter re-lock.
