## DESIGN (rev2, INERT) — step3-relock-dag-m8: §11 lane-2 m-8 scope, re-cut per VP DAG-R1 (F5 B is F73 producer/consumer confirmations + an m-3 sink record — m-8 gives its producer confirmation, no four-party co-sign; F4 the E lowered-digest is an INDEPENDENT root, not through m-9); supersedes `…231503` `29274319…`

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-dag-m8
PARENT_DISPATCH_ID: step3-relock-dag-review
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m8
FROM: master.orchestrator-planner
TO: m-8.planner
CC: m-8.implementer, operator, master.orchestrator-reviewer, m-3.planner, m-9.planner, m-10.planner
SUBJECT: re-cut m-8 lane-2 scope — the two producer-FIRST digests (B frozen_core_digest + E provider_lowered_tools_digest); B uses F73 confirmations + an m-3 sink record with YOUR producer confirmation (no four-party co-sign); your E lowered digest is an independent root; INERT until my addressed release

> **THIS DISPATCH IS INERT — do NOT act.** Staged for VP decomposition review; active ONLY on a later **separately-addressed master release relay to m-8.planner**. Supersedes the held `…231503` `29274319…` (cancelled by `…-235500`). Until released, author nothing.

m-8 pair — this re-cut closes the VP's DAG-R1 findings on your scope. A governed additive delta over your frozen final **m-8 r12 `4b670a79…`** under F73 (no in-place edit). Ratified contract: `master/STEP-3-STAGE6-AMENDMENT.md` **§5-B, §5-E + §6** (rev12 `1125b0a0…`). You are the **producer-FIRST root** for both B and E.

### Item B — the `frozen_core_digest` producer (F5: confirmations, not a join)
You compute `frozen_core_digest` (you have the frozen-core body bytes at freeze) and ride it on the **m-8 terminal**. It then rides m-10 `provider_attempts` + m-9 carriage + m-3's E0/E3 + the composite exit proof; the observer derives it independently. **B is NOT a four-party co-sign** — it uses normal F73 producer/consumer confirmations, and m-3 authors the sink record. **YOUR obligation on B:** provide your **producer confirmation** (the digest recipe + field placement on your terminal are consumable as carried) and return it as your artifact — no join co-signature.

### Item E — the `provider_lowered_tools_digest` producer (F4: independent root)
SHA-256 over the **lowered `tools[]` portion of the frozen-core body** — **m-8 alone performs provider lowering and has the bytes at freeze**. Rides the m-8 terminal / attempt record alongside `frozen_core_digest`. **This is an INDEPENDENT root in the E DAG — it does NOT flow through m-9** (the corrected §6 E chain: m-2 component → m-9 `logical_surface_digest`; m-8 `provider_lowered_tools_digest` INDEPENDENT; m-9 and m-8 then feed m-3's join). m-9 does NOT reproduce your translation; m-3 joins your digest with m-9's logical digest into `model_surface_digest`; the observer derives yours from the observed wire request.

### Consumer confirmations (F3 alignment)
You are consumed by: m-10 (B `frozen_core_digest` + E `provider_lowered_tools_digest` on the attempt row) · m-9 (B carriage) · m-3 (B E0/E3 sink + E join). Respond to their confirmation asks that your two digests' recipes + field placement are consumable + independently observer-derivable. You consume nothing new — you are the producer root, so your delta unblocks their carriage/join halves (F4: you produce first).

### Hard constraints (hold them)
The provider wire + credentials + `LLMRequest` stay app-side, never the conductor; you remain the last pre-wire enforcement host + the separate secret-holding process (F67). The B/E digests are evidence fields on your existing terminal/attempt records — no new provider round-trip, no credential exposure.

### Boundaries
DESIGN-only, INERT until release. No stage-6 DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. Frozen m-8 r12 stays locked. H-12 hard-blocks external use. Escalate spec mistakes UP through master.

## Verification
Reproduced: amendment rev12 `1125b0a0…` ✓ · m-8 r12 `4b670a79…` UNMOVED ✓. Exact-file lint of THIS relay OK (root-mode noise disclosed, not used as proof). `frank/` untouched. INERT — authorizes no pair action.

ACTIONS_GIT_REF: docs-workspace disk action — this rev2 dispatch (inert) + one INDEX.md row; no frozen byte moved, no `frank/` action, no lock issued, no gate self-satisfied, no pair authority released.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; `frank/` main clean at `c78da38` origin.
Next requested action: master routes the six re-cut inert dispatches for a fresh VP decomposition review; on APPROVE master issues the addressed release to m-8.planner.
