## DESIGN (rev2, INERT) — step3-relock-dag-m3: §11 lane-2 m-3 scope, re-cut per VP DAG-R1 (F3 add the E0 carriage for `logical_surface_digest`; F4 evaluator sink is LAST; F5 B is F73 confirmations + an m-3 sink record, not a join); supersedes `…231502` `9c44cd75…`

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-dag-m3
PARENT_DISPATCH_ID: step3-relock-dag-review
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m3
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, operator, master.orchestrator-reviewer, m-8.planner, m-9.planner, m-10.planner
SUBJECT: re-cut m-3 lane-2 scope — item B E0/E3 carriage + the sink RECORD (not a four-party join) + the NEW E0 carriage for logical_surface_digest + the E3 two-digest join + the typed E3 predicates; INERT until my addressed release

> **THIS DISPATCH IS INERT — do NOT act.** Staged for VP decomposition review; active ONLY on a later **separately-addressed master release relay to m-3.planner**. Supersedes the held `…231502` `9c44cd75…` (cancelled by `…-235500`). Until released, author nothing.

m-3 pair — this re-cut closes the VP's DAG-R1 findings on your scope. A governed additive delta over your frozen final **m-3 r4 `009df607…`** under F73 (no in-place edit). Ratified contract: `master/STEP-3-STAGE6-AMENDMENT.md` **§5-B, §5-E + §6** (rev12 `1125b0a0…`).

### Item B — the `frozen_core_digest` carriage + the SINK RECORD (F5: NOT a four-party join)
The m-8-computed digest rides as a field on your `m3.app_event.v1` **E0** + `m3.e3_observation.v1` **E3** attempt vector + the composite exit proof; the observer derives it independently. **B is NOT a two-sided join** — it uses normal F73 producer/consumer confirmations (m-8 producer → m-9 carriage + m-10 row → your carriage), and you author an **m-3 sink record** that binds the end-to-end `frozen_core_digest` consistency + independent observer-derivability across m-8/m-9/m-10. (Only §D is a coordinated two-sided join.)
- **F4 staging: your evaluator sink is LAST within B, not concurrent with its inputs.** Author the E0/E3 schema delta on release; PARK the sink record until m-8's terminal digest + m-9's carriage + m-10's `provider_attempts` row are pair-approved.

### Item E — the `model_surface_digest` join + the NEW E0 carriage (F3)
- **NEW (F3): the E0 schema/carriage for `logical_surface_digest`.** m-9's `logical_surface_digest` rides to the m-10 attempt row **and E0** — you own the **E0 schema/carriage** for it, binding its exact producer identity (m-9). (Previously only m-10's attempt-row carriage was in view; the E0 side is yours.)
- **The E3 two-digest join:** `model_surface_digest` = SHA-256 over `{logical_surface_digest (m-9), provider_lowered_tools_digest (m-8)}` — a join of the two component DIGESTS, assembled by YOU at the E3 binding. You do NOT hash bytes you cannot see; the observer derives each component independently (logical from the worker's presented surface, lowered from the observed wire request).
- **Confirmation alignment (F3):** confirm the resulting **E0 + E3 carriage** of both components against m-9's and m-8's producer identities.
- **F4 staging:** the E3 join + the E0 `logical_surface_digest` carriage stay PARKED until m-9's `logical_surface_digest` recipe and m-8's `provider_lowered_tools_digest` recipe are pair-approved.

### Item E — the typed E3 predicates (owner m-3)
The unchanged predicate id set — `provider_request_matches_frozen_core` · `provider_deny_caused_zero_transport` · `local_invocation_matches_effect_descriptor` · `relay_record_committed_with_stamped_sender` · `no_alternate_credentialed_provider_route_observed` — each `{predicate_id, version, required_inputs[], observed_facts[], evidence_locator, verdict ∈ pass|fail|unknown, observer_id, exact_digest}`, decidable over structured fields; these feed the §7 exit legs (Governance-binding, Injection-visibility, Governed-handoff).

### Consumer confirmations
You are consumed by the §7 exit gate. Confirm you can consume: m-8's B `frozen_core_digest` + E `provider_lowered_tools_digest`; m-9's E `logical_surface_digest` + B carriage; m-10's B `provider_attempts` row + E attempt-row carriage. (Policy stays m-3-owned — the delta adds evidence fields + the E3 join, it does not move egress/observe policy.)

### Boundaries
DESIGN-only, INERT until release. No stage-6 DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. Frozen m-3 r4 stays locked. H-12 hard-blocks external use. Escalate spec mistakes UP through master.

## Verification
Reproduced: amendment rev12 `1125b0a0…` ✓ · m-3 r4 `009df607…` UNMOVED ✓. Exact-file lint of THIS relay OK (root-mode noise disclosed, not used as proof). `frank/` untouched. INERT — authorizes no pair action.

ACTIONS_GIT_REF: docs-workspace disk action — this rev2 dispatch (inert) + one INDEX.md row; no frozen byte moved, no `frank/` action, no lock issued, no gate self-satisfied, no pair authority released.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; `frank/` main clean at `c78da38` origin.
Next requested action: master routes the six re-cut inert dispatches for a fresh VP decomposition review; on APPROVE master issues the addressed release to m-3.planner.
