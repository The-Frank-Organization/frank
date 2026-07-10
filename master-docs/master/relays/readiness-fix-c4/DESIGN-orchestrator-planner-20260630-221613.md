## DESIGN — bounded MUST-fix (Cluster 4a/4b): bring m-2's routing schema current with the c2 R2 lock (grammar-enforce "model ≠ gate input")

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: readiness-fix-c4
PARENT_DISPATCH_ID: readiness-fix-c4
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — bounded design reconciliation; operator on CC
GRILL_REQUIRED: no — VP-gated fold; m-4 reviews the deviation-gate contract
FROM: master.orchestrator-planner
TO: m-2.planner
CC: m-4.planner, master.orchestrator-reviewer, m-2.implementer, m-4.implementer, operator

m-2 — a **bounded** MUST-before-Step-1 fix (`READINESS-REGISTER.md` Cluster 4a/4b), verified line-by-line. Your own doc
is **internally stale**: it states R2 correctly in one place and violates it in another. This is a bring-current fix, not
a redesign — **scope is exactly the two lines below**; m-4 reviews to confirm the deviation-gate contract.

**The contradiction (verified, all in `…forms-determinism/design…`):**
- **4a — the deviation trigger reads `selected_model` (model-as-gate-input), violating R2.** `:285` —
  `justified_deviation` is **"required_when selected_model is off the prior floor."** That keys a **gate** (a required-when
  predicate) on **model identity** — exactly what R2 forbids. Yet **`:289`** (your own D-schema-half) states R2 correctly:
  *"the §5 predicate vocabulary has no `model_*` atom, so model cannot be a gate input even structurally (m-1 coord R2)."*
  And the R2-safe trigger (`declared_deviated == true`, bucket-vs-bucket) lives in **m-4/ARCHITECTURE**, not here. So
  `:285` is **stale vs the c2 R2 lock and self-contradictory with your `:289`.**
- **4b — R2 is not grammar-enforced; the generic atom is a hole.** `:84` — the predicate vocabulary includes the generic
  **`field:<id> <op> <value>`** dependent-required atom. Because `<id>` is open, it **can name `selected_model`** (or any
  model-identity field) → model reaches the gate through the generic atom, so `:289`'s "no `model_*` atom" is **necessary
  but not sufficient**. R2 is asserted, not structurally enforced.

**The bounded fix:**
1. **`:285` → R2-safe trigger.** Change `justified_deviation` **required_when** from `selected_model is off the prior
   floor` to **`declared_deviated == true`** (the agent/courier-declared deviation boolean; the *bucket-vs-bucket*
   comparison, per m-4 §2 / `ARCHITECTURE` R2). The gate keys on the **declaration**, never the model identity.
2. **`:84` → grammar-enforce R2.** Constrain the `field:<id>` atom so **`<id>` cannot name a model-identity field**
   (`selected_model` and any `model_*`): either an explicit exclusion set on the referenceable-field grammar, or a stated
   rule that model-identity fields are **not gate-referenceable**. Goal: R2 becomes **structurally impossible to violate**
   in the schema the tool reads, matching your `:289` claim.

**m-4 review (CC):** confirm the corrected trigger + grammar match the locked deviation-gate contract — `declared_deviated`
/ bucket-vs-bucket, model as payload/bookkeeping only, `routing_ref` traversed for provenance but never model-gating.

**Invariants preserved:** model = payload/bookkeeping never a gate input (R2); trust keyed to stamped seat `FROM`; the
routing decision stays a separate seat-stamped evidenced record; no change to the routing-record contract itself — only
its **schema wording + grammar** are brought current.

**Acceptance ("reconciled") =** `:285` no longer references `selected_model`; the `field:<id>` grammar cannot target a
model-identity field; m-2's schema is internally consistent with its own `:289` and with the m-4/ARCHITECTURE R2 lock;
m-4 confirms.

Not authorized by this relay: no PLAN, no code/spike, no change to the routing-record contract or any other c2 lock, no
re-scoping beyond `:84`/`:285`.

ACTIONS_GIT_REF: wrote this bounded fix-dispatch relay + appended `master/relays/INDEX.md`; no code/source/pcode edits, no design-doc edits (the fold is m-2's), no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-2 folds `:84`/`:285` to the R2-safe form; m-4 reviews/confirms; VP co-signs; then re-verify Cluster 4a/4b closed (MUST-before-Step-1-PLAN).
