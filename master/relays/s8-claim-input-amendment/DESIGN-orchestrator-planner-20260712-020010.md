## DESIGN — master consumes BOTH returns (m-3 §12 r3-approved · m-2 grammar-approved-held) + the two owed actions EXECUTED: (1) the Rail-A parent line is CORRECTED AND RATIFIED as ABSENCE-OPEN / PRESENT-CLOSED — my `…-004511` LEG-m3(e) instruction was wrong and both owner reviews caught it independently; (2) LEG m-7 is EXPANDED to own the `s8-fieldspec-v7` capability transition beside the seam-(v) confirm — m-2's finalize gate 1 is satisfied here, gate 2 is m-7's pending return

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: s8-claim-input-amendment-r2
PARENT_DISPATCH_ID: s8-claim-input-amendment
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — an amendment-internal correction + leg expansion under the standing dispatch (operator CC'd); the slice merge stays operator-only
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-claim-input-amendment
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/s8-claim-input-m3/SITREP-planner-20260712-020000.md
FROM: master.orchestrator-planner
TO: m-7.planner
CC: operator, master.orchestrator-reviewer, m-3.planner, m-2.planner, m-7.implementer, m-3.implementer, m-2.implementer, s8.planner, s8.implementer
SUBJECT: the actionable leg is yours — the expanded LEG m-7 below; the Rail-A ratification and the finalize sequence bind the whole amendment thread; m-3's §12 byte-grain verification is reserved for my three-leg fold; T9 stays held, T10-parallel continues

**1. THE RAIL-A CORRECTION — RATIFIED AS THE AMENDMENT'S RAIL-A OF RECORD (superseding my `…-004511` LEG-m3(e) line):** the `executable_claims` row is **ABSENCE-OPEN / PRESENT-CLOSED**. Absence of a declaration = the honest `Evaluate: nil` no-vantage degrade — open, exactly as before. A **PRESENT declaration is CLOSED/fail-closed**: a v6-capability reader ignoring a present row would fall to nil-Evaluate and ACCEPT what a v7 reader REJECTS — an old reader silently un-governing the record — and "ignore-unknown changes the meaning of acceptance" is Rail A's own closed criterion. The enforcement is the governed `v6→v7` transition + the marker-first capability ceiling: a reader that cannot enforce a present declaration REFUSES the store at phase 0; it never silently ignores the row. **The error was mine** — my dispatch line prescribed additive-open for the row wholesale; both owner reviews (m-2 rev1/rev2 · m-3 F1) independently derived the split, which is the rail applied correctly. Recorded for the trail: the rail caught its own misapplication through independent owner review — the mechanism the amendment exists to exercise.

**2. LEG m-7 — EXPANDED (supersedes the seam-(v)-only scope; r3/F3 of the m-3 review):** your leg is now TWO items, one return:
- **(i) Seam (v) confirm, as originally dispatched:** `observe.Registry` construction + evaluator injection at the composition root + the RegistryEnv wiring (lanes, schema refs, the executor at the locked `spawn(check_id, params, lane_ref, timeout) → CheckVerdict` boundary — no widening); composes with the locked §3 step order; state the §5.1.5a non-interaction.
- **(ii) The `s8-fieldspec-v7` CAPABILITY TRANSITION (your owner relation — the TRIPWIRE's named lawful path):** the fieldspec handler's exact supported set gains `s8-fieldspec-v7`; the forward relation extends over m-2's successor map (the implementation locus today knows only v5→v6 — `config.go:225-226`); marker-first refusal order preserved (the marker is read and validated BEFORE any member content, zero partial interpretation); reader-first/forward-only sequencing stated (upgrade the reader, then accept the v7 record). **Both phase-0 proof legs named in your return:** (a) a v7 store × a v6-capability reader ⇒ the typed marker-first phase-0 refusal — the PRESENT-CLOSED half made mechanical; (b) an upgraded reader × a pre-transition store ⇒ loads clean, and the governed `v6→v7` transition accepts through the standing §2.4 gate.
Return under sub-ID `s8-claim-input-m7`, pair-internal review per your standing practice.

**3. THE SEQUENCE (m-2's three gates, dispositioned):** gate 1 (the Rail-A correction) — **SATISFIED by §1 above**; gate 2 — your return per §2; gate 3 — m-2 finalizes the byte-exact row + the `v6→v7` class **against the approved m-3 §12 + your reviewed return**, then hands to the s8 build on the T2 owner-fidelity pattern. Master then reconciles the three legs at the byte grain (§12 read whole at my seat at that fold — reserved deliberately) and issues the bounded T9 fold/grant (`seam (v)` + `registry.json` v7 bytes + `s8_exit_gate_test.go`, pre-staged in `s8-build-escalate-fence/…-004510`). **T9 stays held on its preserved RED; T10-parallel continues; the exit condition is unchanged and executable** (`TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate` GREEN on the real path, or no exit).

**Consumed this relay:** m-3's return (`s8-claim-input-m3/SITREP-…-020000`; §12 approved r3 — shape/cardinality, dual-locus validation, R2, (h)-pairing, the §12g worst-wins aggregation coherent with the locked NF-S6/decision-② dispositions) · m-2's return (`s8-claim-input-m2` SITREP `…-010800`; grammar approved rev2 — the honest carrier grain with nested validation delegated to the m-3 seam, no nested-enforcement over-claim; their disclosed review arc noted with approval).

ACTIONS_GIT_REF: none — ratification + leg expansion only (disk refs: this relay + one INDEX.md row timestamped 20260712-020010; stamped after the replied-to filename per the skew convention — author wall clock 20260711-231510).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `691d034`; the s8 build lane (`s8-observe-spine@3cce8cd`) is the pair's, T10 in flight.
Next requested action: operator carries this to m-7.planner (the expanded leg); on m-7's return m-2 finalizes (no further master touch needed between those two — the sequence is authorized end-to-end); master folds the three legs and issues the T9 grant.
