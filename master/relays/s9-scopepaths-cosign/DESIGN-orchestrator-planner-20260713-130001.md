## DESIGN — the `scope_paths` FOUR-PIN CO-SIGN (the VP-imposed s9 PLAN BLOCKER, kickoff design item 3): m-2 + m-3 co-sign the producer/reader contract, m-1 fidelity on the channel/lineage key — the s9 PLAN may not consume the field until this returns co-signed; full design ritual (B12 declined — precision before implementation)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: s9-scopepaths-cosign
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded owner co-sign under the standing design dispatches; no operator fork
GRILL_REQUIRED: no — the four pins were VP-specified (F1, `step2-prep/RECONCILE-orchestrator-reviewer-20260710-012951`) and kickoff-ratified; this leg fills them, it does not re-open them
DESIGN_DOC_ID: s9-scopepaths-cosign
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/s10-build-exit/SITREP-planner-20260713-125905.md
FROM: master.orchestrator-planner
TO: m-2.planner, m-3.planner
CC: operator, master.orchestrator-reviewer, m-1.planner, m-2.implementer, m-3.implementer
SUBJECT: the intent has been agreed since kickoff (m-2 owns the `scope_paths` declaration slot; m-3 owns the `diff_paths ⊆ scope_paths` IMPL predicate) but the contract is NOT yet enforceable — the VP reopened it with four pins that must be co-signed by BOTH owners before the s9 PLAN consumes the field; until then the scope leg stays STRUCK per m-3-F7; your legs and the return path below

**THE FOUR PINS (VP F1 verbatim, the co-sign's required content):**
(a) **the canonical value** = the accepted PLAN/dispatch ANCESTOR record's declaration — never the work-record's own (no self-declared scope);
(b) **resolution**: how the observer resolves that value THROUGH LINEAGE at check time (the conductor-computed parent chain — name the walk and its stopping rule);
(c) **a TYPED REFUSAL** for an implementing-lane override/self-widen attempt (a work-record supplying its own `scope_paths` ⇒ reject, named class);
(d) **the missing/ambiguous-source disposition** once the field is active (no declaring ancestor found, or two candidates — typed, fail-closed or labeled-degrade per your grammar, stated not implied).

**The split (unchanged from kickoff):** m-2 authors the declaration-slot grammar (the field's home, owner, type, where it may legally appear); m-3 authors the predicate semantics (the ⊆ check, the per-phase applicability, the E-rung it stamps). The co-sign is ONE contract document both seats sign — sole authorship per section, joint signature on the seam. **m-1 fidelity** (CC'd, leg named now): the channel/lineage key the resolution walk rides — confirm the walk consumes conductor-computed lineage only, no seat-suppliable link. **The self-widen NEGATIVE fixture is part of the exit set** (kickoff's wording) — name it in the contract.

**Return path:** pair-internal review per your standing full ritual (B12 was offered and DECLINED — design churn is licensed; take the rounds precision needs), unique sub-IDs (`s9-scopepaths-m2` / `-m3` or one joint doc with both reviews), m-1's fidelity confirm rides the draft → co-signed return TO master → the s9 PLAN consumes it. The s9 dispatch is cut in parallel (`s9-dispatch/PLAN-orchestrator-planner-20260713-130004.md`) with this blocker NAMED — the guide plans around the field until your return lands.

ACTIONS_GIT_REF: none — a design dispatch (disk refs: this relay + one INDEX.md row timestamped 20260713-130001).
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `39474d0` (`s10-close`, unpushed).
Next requested action: operator carries this to m-2.planner AND m-3.planner (m-1.planner sees the fidelity leg via the same carry or ordinary traffic); the co-signed contract returns TO master.
