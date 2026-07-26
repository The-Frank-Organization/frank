## DESIGN (rev2, INERT) — step3-relock-dag-m9: §11 lane-2 m-9 scope, re-cut per VP DAG-R1 (F2 full D obligations + F4 staging + F5 B-not-a-join); supersedes `…231500` `af1bd19a…`

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-dag-m9
PARENT_DISPATCH_ID: step3-relock-dag-review
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m9
FROM: master.orchestrator-planner
TO: m-9.planner
CC: m-9.implementer, operator, master.orchestrator-reviewer, m-10.planner, m-10.implementer, m-2.planner, m-3.planner, m-8.planner, m-1.planner
SUBJECT: re-cut m-9 lane-2 scope — item D with the FULL ratified obligations (exclusive-writer boundary + branch ownership + identity-exact records + the content-ready RECEIPT you produce + marker-before-outcome + the total first-action table + retention) + M9-D2 + B/C/E halves; INERT until my addressed release

> **THIS DISPATCH IS INERT — do NOT act.** It is staged for VP decomposition review and becomes active ONLY on a later **separately-addressed master release relay to m-9.planner** (a VP verdict is not a release; this file is not a release). It supersedes the held `step3-relock-dag-m9/DESIGN-orchestrator-planner-20260721-231500.md` `af1bd19a…` (cancelled by `step3-relock-dag-hold/…-235500`). Until released, author nothing.

m-9 pair — this re-cut closes the VP's DAG-R1 findings on your scope. Everything is a governed additive delta over your frozen finals **worker r7 `cb7ff970…` + lifecycle r21 `4d3bd14e…`** under F73 (no in-place edit). Ratified contract: `master/STEP-3-STAGE6-AMENDMENT.md` **§5-D (:151-329) + §5-B/C/E + §6** (rev12 `1125b0a0…`). Grain: the amendment fixes the decomposition + acceptance properties + names your Tier-HARD obligations; you DESIGN the internals (record/segment/rotation grammar, the fence mechanism).

### Item D — m-9 half (the FULL ratified obligations, F2 closed)
1. **D1 acceptance properties (i)–(vi), all six:** (i) `fsync`-durable-append linearization · (ii) an **enforceable exclusive-writer boundary** — under it a **retired generation cannot extend or corrupt the successor's valid prefix, and stale/predecessor writes are decidable** (a bare `generation_id` label is explicitly INSUFFICIENT) · (iii) torn-tail valid-prefix recovery returning ONE deterministic prefix per crash cut · (iv) file identity bound `{run_id, run_manifest_digest}`, fail-closed · (v) content-before-outcome durability ordering · (vi) the log path rides `turn_open` (one carrier).
2. **Writer-fence BRANCH OWNERSHIP (name your branch):** a local **OS exclusive lock** (acquired ordered-after-predecessor-termination) is **m-9-owned**; **m-10-ordered per-generation segments** is a **JOINT m-10-producer/m-9-consumer** design (designed + pair-reviewed together, join record). If you select the segment branch, say so — m-10's re-cut carries the conditional producer obligation and you design it jointly, not m-9-only.
3. **Identity-exact records:** every content record binds its `tool_call_id`/`attempt_id` **plus source `turn_id`** (so reconciliation is exact across the continuation chain).
4. **The content-ready RECEIPT — you PRODUCE it (the F2 gap: m-10 consumes a receipt with no assigned writer).** m-9 produces the durable **content-ready receipt** bound to `{turn_id, attempt_id, valid-prefix/marker digest}` that m-10's composite provider settlement consumes. The exact receipt frame/table is a **JOINT m-9/m-10 DESIGN obligation** (both directions) — design it with m-10, join record.
5. **Tool ordering (marker-before-outcome, not generic):** the content record **AND its admitting durable `round_marker`** fsync-linearize **BEFORE** `record_tool_outcome` (so `settled_tools ⇒ content in the durable valid prefix`).
6. **D2 reconciliation-consume + the two time-scoped trust properties:** consume m-10's producer-total 3-class manifest on `turn_open`; trust content **only** under matching `settled_with_content` evidence AND presence in the current recovered valid prefix — else **`content_lost` (your post-inspection result) → `DEGRADED`**, never fabricated.
7. **The TOTAL first-action table (all five branches):** clean-positive resume · determinate-terminal (surface+terminalize, no auto-replacement) · uncertain-tool (surfaced, no silent re-execute) · uncertain-provider (surfaced, no auto-resend) · degraded/content_lost (`resume_action`=re-derive/abandon). No silent advance across any non-clean terminal.
8. **§7.1 invariant supersession (owner confirm):** "no m-9-owned durable session store" narrows to "no m-9-owned durable **OUTCOME** store" — content persists, every outcome stays m-10-canonical. Confirm you own + carry this.
9. **Retention:** the log is retained **per-run** and GC'd on **run-terminal**.

### M9-D2 — the affected final
Fold the above into worker r7 + lifecycle r21 by governed delta: consume the D2 manifest + `uncertain` + log inspection/reconciliation + the post-commit disposition-receipt no-work gate, and route a **broker-cut relay identity** (per the co-signed §D join record `step3-relock-broker-confirm/…-215500`) through the ratified `uncertain` branch.

### B / C / E (F3/F4/F5)
- **B — carriage (NOT a join, F5):** carry the m-8-computed `frozen_core_digest` on your m-9 attempt carriage; confirm consumability to m-3's **sink record** via normal F73 producer/consumer confirmations (B is not a two-sided join — only §D is).
- **C — executor derivation + record:** derive + record the per-action effect descriptor per the §5-C table (single cwd encoding, `env_digest` over the m-1-sanitized presented env, `shell_interpreter_ref{path,version,content_id}` mandatory-or-`unknown`); consume m-10's ticket schema; honor m-1's env/no-leak rule.
- **E — `logical_surface_digest` (owner m-9):** SHA-256 over JCS `{instructions, logical_tool_schemas[], tool_descriptions[], compaction_template, policy_messages}`, folding **m-2's** schema/description component; rides to the m-10 attempt row **and E0** (m-10 carries it — see m-10's E-row). **m-8's `provider_lowered_tools_digest` is an INDEPENDENT root — you do NOT reproduce m-8's lowering** (F4 DAG correction).

### Staging (F4) — consumer sections PARKED
Your **producer** obligations (D1 log, `logical_surface_digest`, the content-ready receipt design) proceed on release; your **consumer** sections stay PARKED until the exact pair-approved producer bytes arrive: m-2's E component (before you fold it), m-10's C ticket schema + D manifest + disposition-receipt (before you consume them). Your final pair review must cover the settled producer bytes, not promises.

### Two-sided join record
Only **§D (m-9 ⇄ m-10 + m-1 redaction)** is a coordinated two-sided join — your D1/D2/consume/receipt half + m-10's manifest/lifecycle/receipt half + m-1's redaction, one joint record. (§B is confirmations + m-3's sink record, not a join.)

### Boundaries
DESIGN-only, and INERT until release. No stage-6 DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. Frozen r7/r21 stay locked. H-12 hard-blocks external use. Escalate spec mistakes UP through master.

## Verification
Reproduced: amendment rev12 `1125b0a0…` ✓ · worker r7 `cb7ff970…` + lifecycle r21 `4d3bd14e…` UNMOVED ✓ · broker rev8 `64f9136e…` ✓. Exact-file lint of THIS relay OK (root-mode noise disclosed, not used as proof). `frank/` untouched. This re-cut is INERT (staged for VP review); it authorizes no pair action.

ACTIONS_GIT_REF: docs-workspace disk action — this rev2 dispatch (inert) + one INDEX.md row; no frozen byte moved, no `frank/` action, no lock issued, no gate self-satisfied, no pair authority released.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo; `frank/` main clean at `c78da38` origin.
Next requested action: master routes the six re-cut inert dispatches for a fresh VP decomposition review; on APPROVE master issues the addressed release to m-9.planner and the pair proceeds.
