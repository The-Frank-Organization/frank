## DESIGN — step3-relock-dag-m3: §11 lane-2 (interface DAG) m-3 scope — item B (the `frozen_core_digest` E0/E3 carriage + evaluator join) + item E (the `model_surface_digest` join + typed E3 predicates), over frozen m-3 r4 `009df607…`

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-dag-m3
PARENT_DISPATCH_ID: step3-relock-broker-confirm
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m3
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, operator, master.orchestrator-reviewer, m-8.planner, m-9.planner, m-10.planner
SUBJECT: your lane-2 scope — item B (frozen_core_digest carriage on E0/E3 + the evaluator join, the DAG SINK) + item E (model_surface_digest join + the typed E3 predicates); governed additive delta over frozen m-3 r4 under F73

m-3 pair — §11 lane 1 is CLOSED; the operator opened **lane 2 (the §6 interface DAG) in full**. This is your complete lane-2 scope. Run your normal pair cycle. A **governed additive delta over your frozen final m-3 r4 `009df607…`** (stays the historical lock; no in-place edit; reviewed under F73 with consumer confirmations). Ratified contracts: `master/STEP-3-STAGE6-AMENDMENT.md` **§5-B, §5-E + §6** (rev12 `1125b0a0…`).

### Your obligations
1. **Item B — the `frozen_core_digest` join (§5-B), and you are the evaluator-join SINK of the B DAG.** The m-8-computed digest rides as a field on: your `m3.app_event.v1` **E0** app-event + your `m3.e3_observation.v1` **E3** attempt vector + the composite exit proof. The observer derives the digest **independently** (no prompt/response bytes enter the conductor). Per §6, the B DAG is: **m-3 (E0/E3 schema delta) + m-8 (terminal digest) FIRST → m-9-carriage ∥ m-10-row (siblings) → your m-3 evaluator join.** You author the E0/E3 schema delta now and the evaluator join once m-8's terminal digest + the m-9/m-10 carriage rows are settled.
2. **Item E — the `model_surface_digest` join + the typed E3 predicates (§5-E, owner m-3).**
   - `model_surface_digest` = SHA-256 over `{logical_surface_digest, provider_lowered_tools_digest}` — a join of the two component DIGESTS (not bytes), assembled by YOU at the E3 binding. The observer derives each component independently (logical from the worker's presented surface, lowered from the observed wire request). **You do NOT hash bytes you cannot see** — m-9 owns `logical_surface_digest`, m-8 owns `provider_lowered_tools_digest`; you join the two digests. LOCK = the two component recipes + field sets + producer ownership + carriage + your join recipe.
   - The typed E3 predicate ids (yours, unchanged set): `provider_request_matches_frozen_core` · `provider_deny_caused_zero_transport` · `local_invocation_matches_effect_descriptor` · `relay_record_committed_with_stamped_sender` · `no_alternate_credentialed_provider_route_observed` — each `{predicate_id, version, required_inputs[], observed_facts[], evidence_locator, verdict ∈ pass|fail|unknown, observer_id, exact_digest}`. These feed the §7 exit-gate legs (Governance-binding, Injection-visibility, Governed-handoff) — realize them so the gate predicates are decidable over structured fields.

### Two-sided join record (co-sign)
- **§B evaluator join:** your evaluator join is the sink — co-sign with m-8's terminal digest + m-9's carriage + m-10's `provider_attempts` row that the `frozen_core_digest` carriage is consistent end-to-end and independently observer-derivable.

### Consumer confirmations you owe / receive
Confirm you can consume: m-8's B `frozen_core_digest` + E `provider_lowered_tools_digest` · m-9's E `logical_surface_digest` + B carriage · m-10's B `provider_attempts` row. You are consumed by: the §7 exit gate (your E3 predicates + `model_surface_digest`).

### Boundaries
DESIGN-only. No stage-6 DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. Frozen m-3 r4 stays locked; the delta is governed + additive. H-12 hard-blocks external use. Note: policy stays m-3-owned — the amendment adds evidence fields + the E3 join, it does not move egress/observe policy. Escalate spec mistakes UP through master.

### Where this sits
§11 lane 2 of 5. Item A (recipe + `bundle_sha256`) authored LAST over settled B–E; lane 4 = the shorter re-lock; lane 5 = T4. Your E3 predicates + `model_surface_digest` are load-bearing for the §7 exit legs proven at T4.

## Verification
Reproduced from disk: amendment rev12 `1125b0a0…` ✓ · m-3 r4 `009df607…` UNMOVED (governed delta requested, not an edit) ✓. Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row; no frozen design byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the m-3 pair authors + pair-reviews its lane-2 delta (B E0/E3 carriage + evaluator join · E model_surface_digest join + E3 predicates), returns byte-bound design + F73 consumer confirmations + the co-signed §B join record + a SITREP; master integrates toward the shorter re-lock.
