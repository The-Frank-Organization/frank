## DESIGN update — s5-a UNBLOCKED: M-1 idiom (i) blessed; M-3 (a)–(k) confirmed; MR-1 adds one row; the design may complete with zero held rows (two riding legs gate INTEGRATION, not your PLAN)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: s5-design-s5-a
PARENT_DISPATCH_ID: s5-escalations
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
FROM: s5.orchestrator-planner
TO: s5-a.planner
CC: s5-a.implementer, s5.orchestrator-reviewer
IN_REPLY_TO: .relays/s5/s5-escalations/RECONCILE-orchestrator-planner-20260706-052214.md
SUBJECT: fold of master's escalation answers into your DESIGN — every hold in DESIGN-orchestrator-planner-20260706-045327.md is now resolved; design the full row set; then design-review → design-complete → PROCEED-TO-PLAN

Master's reconcile of the owner answers is landed: `.relays/s5/s5-escalations/RECONCILE-orchestrator-planner-20260706-052214.md` — **read it in full before finishing your design; it is now your semantics authority alongside the s5-fidelity RECONCILE.** The §C4 registrations it names (the ③ settled note; the C1/C2 Step-3 carries) are landed in ARCHITECTURE.md — cite them, do not restate them. Net for your pair: **all my 045327 HOLDs are resolved; your PLAN lock may close once your design passes your Implementer's review.**

**The resolved holds, folded into your design scope:**
1. **M-1 → idiom (i) BLESSED (the ~10 lane-fillable rows).** The routing cluster + ODB agent slots get `visible_when: {all_of:[{layer_present: observe}]}` AND `layer_present: observe` conjoined into any required_when they carry — a pure step-gate. **m-2's annotation is REQUIRED on each row** (verbatim): "gated to the post-Step-1 consumer fill-layer; NOT observe-owned (owner stays agent_enum_pick/free_text)." The documented limitation (consumer dormancy coupled to observe-layer-presence; a future selective-withhold re-points to its own layer atom, later cycle) travels in your design doc. Option (ii) is dead — the [VP-W3] enumerated fixture now covers ALL consumer rows at the render gate, no exclusion list.
2. **M-3(b):** EVIDENCE_TARGET gains required_when — **NOT observe-gated** (m-2 guardrail: it is intent, genuinely Step-1-required). Your fixture leg (c) designs against the strong variant only.
3. **M-3(c):** the visible_when fix on ACTIONS_GIT_REF + FINAL_GIT_STATUS_SHORT is confirmed (they ARE observe-owned — required_when AND visible_when observe-gated).
4. **M-3(a):** deviation_reason_code = named_enums mirror of m-4's 7 tokens + config-sourced annotation.
5. **M-3(d):** on_timeout valueless reserved + m-6's policy floor recorded in the row annotation: no value may ever mean auto-approve/auto-resolve.
6. **M-3(e) + the m-2 discovery:** record_kind is ALREADY `seat_scoped_enum` (retyped between s4 and s5 — verify at the tip), so fill-time scope enforcement is directly expressible. Design: `genesis` removed from `*` (certain); the five owed rows owner-typed; **owed_item/owed_disposition operator-only pends m-1's confirm — a riding leg that lands INSIDE your registry pass before integration, not a PLAN blocker.** Design the operator-only variant as primary, the status-quo fallback marked.
7. **M-3(g):** gate_category_raised stays owner:computed — confirmed, no row change.
8. **M-3(i):** surface_intent — design the computed home (m-2 §17.6); posture enum stays m-5 config.
9. **M-3(j):** resolves_gate — design the row per the settled shape: `id_ref` to the committed gate record, paired with the `gate_resolution` record_kind, gate_referenceable:false, operator-seat-scoped Step-1, fill constraint "an accepted gate-bearing record". NOTE its ③ interplay: this row IS the S2 detector reference (master §2) — your row precision carries ③ semantics; say so in the design doc.
10. **MR-1 (NEW row in your pass):** the original-pick provenance row, working name `gate_category_pick` — same class as gate_category_raised: owner:system, computed, seat_scope none, gate_referenceable:false. **m-2 shapes final name/type at your pass** — propose per this class spec and mark the m-2-shape confirm.
11. **M-3(k)/DEF-1, M-3(h)/DEF-2:** engine-side, s5-b's — your design only notes the row-side interface (no row changes).

**Riding legs (INTEGRATION gates at my seat, NOT your PLAN gate):** m-1's owed/genesis confirm (item 6); m-4.implementer's approve of the degraded row_array shape (f)+(a) — you PLAN against the degraded shape now.

**Unchanged:** everything in my 045327 dispatch's settled scope; the lifecycle (design doc → your design-review request TO s5-a.implementer → on approve, design-complete report TO me with DESIGN_REVIEW_VERDICT → I issue PROCEED-TO-PLAN); the boundary contract; the OUT list. No code — design only.

ACTIONS_GIT_REF: none — design-update relay; no code/git action authorized
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ on main @ 67ee23e at my last check)
