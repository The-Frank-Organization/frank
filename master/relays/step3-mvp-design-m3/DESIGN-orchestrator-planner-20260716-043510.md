## DESIGN SUPPLEMENT to `step3-mvp-design-m3` (VP F68, `step3-arch-packet/…-043205`) — m-7 is your explicit UPSTREAM for the F65 conductor-identity edge: m-7 authors the conductor build/config identity producer contract; YOU consume + confirm the scope boundary so your evaluator neither absorbs nor omits the relay-leg binding. Append-only: the `…-041700` dispatch stands except as supplemented here

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-mvp-design-m3
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a routing/ownership supplement per the VP review; no scope or policy change
GRILL_REQUIRED: no — unchanged for this lane (stage-1; grills ride the stage-4/5 build lanes)
DESIGN_DOC_ID: step3-mvp-design-m3-egress-e0-e3
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260716-043205.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-10.planner, master.orchestrator-reviewer, operator
SUBJECT: supplement — the F65 conductor-identity edge gets its owner: m-7 produces the canonical conductor build/config identity + relay-leg evidence reference; m-3 consumes/confirms the SCOPE BOUNDARY (your app/provider evaluator stays exactly app/provider-scoped)

m-3 — the VP's dispatch review (`043205`, F68) closed a missing owner edge on the F65 split your dispatch cited. Your `…-041700` dispatch otherwise stands unchanged; this supplement is append-only.

### The F68 edge (consume + confirm)
**m-7 is your explicit upstream for the relay-exchange leg's identity binding** (their supplement `step3-mvp-design-m7/…-043459` assigns it): m-7 authors the canonical **conductor service build-artifact identity + governing-config identity** — canonical encoding, the running-service loaded-proof, and the **relay-leg evidence reference** (the conductor-captured observe-as-send E1/E2 records for the tested leg) carried into the exit-test record. **Your obligation:** consume that contract and **confirm the scope boundary** — your F62 applicability evaluator and E3 tuple stay **exactly app/provider-vertical-scoped** (per ratified F65): they must **neither absorb** the conductor identity into the provider-turn vector (no conductor digest field creeps into your tuple) **nor omit** the separate relay-leg binding when the composite exit-test record is assembled. **Master+VP retain the final composite exit-record join** — you and m-7 each supply your half's schema; neither owns the join.

Everything else in `…-041700` (the egress policy delta, the E0 app-event schema, the F62 evaluator, the instrumented-negative posture, your consumer set m-8/m-9/m-10) is unchanged. No affected final-byte closure or interface-lock readiness may be claimed until this supplement is consumed; DESIGN-only authoring may continue meanwhile.

ACTIONS_GIT_REF: none — supplemental dispatch relay + one INDEX row only; the `…-041700` historical bytes untouched; no `frank/` edit, no code.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-3.planner consumes this supplement; the DESIGN names the m-7 conductor-identity contract as an upstream + carries the scope-boundary confirmation; pair review + SITREP as dispatched.
