## DESIGN — step3-relock-dag-m2: §11 lane-2 (interface DAG) m-2 scope — item E producer component (the `logical_tool_schemas[]` / `tool_descriptions[]` contribution to m-9's `logical_surface_digest`), over frozen m-2 `83d8e63e…`

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-dag-m2
PARENT_DISPATCH_ID: step3-relock-broker-confirm
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m2
FROM: master.orchestrator-planner
TO: m-2.planner
CC: m-2.implementer, operator, master.orchestrator-reviewer, m-9.planner, m-3.planner
SUBJECT: your lane-2 scope — the smallest: item E producer component (the schema/description bytes feeding m-9's logical_surface_digest); governed additive delta over frozen m-2 under F73

m-2 pair — §11 lane 1 is CLOSED; the operator opened **lane 2 (the §6 interface DAG) in full**. Your lane-2 scope is the narrowest of the six: the **item E producer component**. Run your normal pair cycle. A **governed additive delta over your frozen final m-2 `83d8e63e…`** (stays the historical lock; no in-place edit; reviewed under F73 with consumer confirmations). Ratified contracts: `master/STEP-3-STAGE6-AMENDMENT.md` **§5-E + §6** (rev12 `1125b0a0…`).

### Your obligation
**Item E — the schema/description component of `logical_surface_digest` (§5-E).** m-9 owns `logical_surface_digest` = SHA-256 over JCS `{instructions, logical_tool_schemas[], tool_descriptions[], compaction_template, policy_messages}`, but its **`logical_tool_schemas[]` / `tool_descriptions[]` component is supplied by YOU** (your form→tool-schema mapping is where the canonical tool schemas + descriptions live). Deliver the exact component contract: what canonical bytes m-2 contributes, the field set, the versioning/mapping-version binding, and the JCS-stable serialization — so m-9 can fold your component into its digest deterministically and m-3 can independently derive the logical component at the observer. Per §6: **E is m-2 (schema/description) → m-9 `logical_surface_digest` → m-3 join in E3.** You are a producer-first for the logical component.

### Consumer confirmations you owe / receive
You are consumed by: m-9 (folds your component into `logical_surface_digest`) and, indirectly, m-3 (independent observer derivation of the logical component). Respond to m-9's / m-3's confirmation asks that your component contract is consumable + independently derivable. You consume nothing new here.

### Boundaries
DESIGN-only. No stage-6 DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. Frozen m-2 stays locked; the delta is governed + additive. H-12 hard-blocks external use. Escalate spec mistakes UP through master. (Your form-schema + field-ownership model is otherwise unchanged — this adds only the digest-component contract.)

### Where this sits
§11 lane 2 of 5. Your component feeds m-9's `logical_surface_digest` → m-3's `model_surface_digest` join in E3. Item A (recipe + `bundle_sha256`) authored LAST over settled B–E; lane 4 = the shorter re-lock; lane 5 = T4.

## Verification
Reproduced from disk: amendment rev12 `1125b0a0…` ✓ · m-2 `83d8e63e…` UNMOVED (governed delta requested, not an edit) ✓. Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row; no frozen design byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the m-2 pair authors + pair-reviews its lane-2 delta (the E schema/description component contract), returns the byte-bound design + F73 consumer-confirmation responses + a SITREP; master integrates toward the shorter re-lock.
