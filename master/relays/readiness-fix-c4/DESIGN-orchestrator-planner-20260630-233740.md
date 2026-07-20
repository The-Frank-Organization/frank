## DESIGN — bounded MUST-fix (Cluster 4 completion: 4b + 4c): m-2 routing FieldSpec → m-4's full record + declare computed-field homes

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: readiness-fix-c4
PARENT_DISPATCH_ID: readiness-fix-c4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — bounded design reconciliation; operator on CC (operator chose "complete the fold")
GRILL_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-2.planner
CC: m-4.planner, m-3.planner, m-5.planner, m-6.planner, master.orchestrator-reviewer, m-2.implementer, m-4.implementer, operator

m-2 — completing **Cluster 4** after the VP's correct revise (`readiness-fix-c4/…-232925`): **4a is closed
(VP-approved)**; this dispatch closes **4b + 4c**, which my earlier c4 dispatch under-scoped (my error, owned). The
operator chose **complete the fold** (not narrow). **Bounded to exactly the FieldSpec declarations below** — no
routing-record *contract* change (m-4-owned), no tag-value micro-fold, no other c2 change.

**4b — mirror m-2's routing FieldSpec (§12/§17.3) to m-4's LOCKED routing record (`m-4 …routing-policy…:200-210`).**
m-2's `routing_assignments` is a `row_array` but does not enumerate the per-row fields, and the record omits several
m-4 fields. Bring current:
- **Enumerate the per-row `routing_assignments` fields** (m-4:203): `task_tag`, `declared_bucket`, `chosen_model`
  (`pinned|resolved`), `pin_mode` (`pinned|slot`), **`seat_archetype`** (opaque — concrete values m-5-owned, c3),
  **`authority_ceiling`** (resolved-at-spawn, **recorded per assignment** for replay/audit — the F2 fold). `declared_deviated`
  is already per-row (done in the retype). Model-identity fields stay **non-gate-referenceable** (R2, unchanged).
- **Add `deviation_reason_code`** (m-4:206): `agent_enum_pick`, **`required_when any(routing_assignments.declared_deviated
  == true)`** — the *same grain/treatment* as `justified_deviation`, so the reason code can't be omitted where policy makes
  it load-bearing. (This is the field most coupled to the 4a deviation gate — it belongs with the fix, not deferred.)
- **Add `constraints`** (m-4:207 — budget/latency/privacy, **reserved/forward**) and **`template_ref`** (m-4:209 —
  `system`/`id_ref`, set when spawned from a template, null otherwise) as **reserved-shape** declarations (shape only, no
  Step-1 values — same reserved-seam pattern as `slot_in`/`certification`).

**4c — declare the conductor-computed fields' schema home** (register 4c; evidence `m-3 …:110`, `m-1 …:222-223`,
`ARCHITECTURE.md:351-352`, `m-5 …:142`): **`record_integrity`**, **`surface_intent`**, **`posture`** are consumed (ODB /
m-6 delivery) but not enumerated in m-2. Declare them as **`system`/`computed_result` field declarations** so the m-1
requirement ("m-3 allowlist resolves to an m-2-declared set") + canonical-iff-consumed (registry entry) are satisfied.
Coordinate the exact home/typing with the **owners on CC**: m-3 (`record_integrity`/`surface_intent`), m-5
(`posture`/archetype), m-6 (delivery consumption). If `posture` has no clean F2 slot, say so — that's a real gap to
surface, not paper over.

**CTO/VP clarification (the register's explicit ask, now sanctioned):** these are **additive FIELD declarations** — the
fields consumers already require — and are **distinct from the forbidden "tag-value micro-fold."** The m-2 "no new field"
guardrail forbade smuggling tag-VALUE semantics into the schema, **not** declaring the additive fields the locked m-4
record + consumers need. This dispatch sanctions the additive field declarations; it does **not** permit any tag-value
micro-fold.

**Review edges:** **m-4** confirms the routing-record mirror matches its locked record; **m-3/m-5/m-6** confirm their
computed/opaque fields' declared homes/typing. (Same pattern as m-4's 4a confirm.)

**Invariants (must-not-change):** routing-record *contract* is m-4-owned and unchanged (m-2 mirrors it as FieldSpec);
**R2 preserved** — new fields are declarations, the gate-referenceable allowlist still excludes model-identity fields, no
predicate keys on model; reserved fields carry shape only, no Step-1 values; trust keyed to stamped seat `FROM`.

**Acceptance ("Cluster 4 reconciled") =** m-2 §17.3/§12 enumerates m-4:200-210's full record (per-row fields +
`deviation_reason_code` + reserved `constraints`/`template_ref`); `record_integrity`/`surface_intent`/`posture` have
declared homes; m-4 confirms the mirror; m-3/m-5/m-6 confirm their fields; then CTO/VP re-verify → **Cluster 4 CLOSED**
(4a already closed).

Not authorized: no PLAN, no code/pcode/spike, no routing-record *contract* change, no tag-value micro-fold, no re-scope
beyond §12/§17.3 + the named fields, no R2 regression.

ACTIONS_GIT_REF: wrote this bounded fix-dispatch relay + appended `master/relays/INDEX.md`; no design-doc edits (the fold is m-2's), no code/source/pcode, no PLAN, no spike.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
Next requested action: m-2 folds the routing FieldSpec to m-4's full record + declares the computed-field homes; m-4 + m-3/m-5/m-6 confirm their fields; then CTO re-verify + VP closure co-sign → Cluster 4 CLOSED. Then the full MUST gate is genuinely satisfied.
