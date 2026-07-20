## DESIGN SUPPLEMENT to `step3-mvp-design-m7` (VP F67 + F68 + F69, `step3-arch-packet/…-043205`) — GRILL_REQUIRED flips to YES (broker-placement grill + durable GRILL_LOCK before your implementer's final-byte review); you additionally OWN the F65 conductor-identity producer contract; m-2 joins your consumer-confirmation set. Append-only: the `…-041630` dispatch stands except as supplemented here

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-mvp-design-m7
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the placement grill is an m-7 design choice UNLESS the selected answer would alter a ratified topology or claim boundary; only that case returns to the operator (VP F67)
GRILL_REQUIRED: yes — SUPERSEDES the `…-041630` "no": broker process placement is a hard-to-reverse cross-domain boundary at large ceremony with no later m-7 grill stage; a durable GRILL_LOCK is owed IN the m-7 DESIGN record before m-7.implementer reviews the final bytes
DESIGN_DOC_ID: step3-mvp-design-m7-transport-broker
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260716-043205.md
FROM: master.orchestrator-planner
TO: m-7.planner
CC: m-7.implementer, m-1.planner, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
SUBJECT: supplement — (F67) broker-placement GRILL + GRILL_LOCK required pre-final-review; (F68) m-7 owns the canonical conductor build/config IDENTITY producer contract for the F65 relay-leg binding, m-3 consumes/confirms the scope boundary; (F69) m-2 added to your consumer set

m-7 — the VP's dispatch review (`043205`) corrects three things on your lane. Your `…-041630` dispatch otherwise stands unchanged; this supplement is append-only.

### F67 — the broker-placement GRILL (blocker; ceremony corrected)
`GRILL_REQUIRED` is **YES** for this lane. Before m-7.implementer reviews your final bytes, fold a **durable `GRILL_LOCK`** into the DESIGN record comparing **at least the two named placements** — (i) the broker as its **own process** vs (ii) a **protected thread/module in the app main process** — against, at minimum: the ratified **F57 claim boundary** (what each placement lets you honestly claim about credential exposure) · **m-1's credential semantics** (the parallel `step3-mvp-design-m1` contract) · **m-10's no-conductor-verb / no-credential-bytes rail** (sharper if the broker shares m-10's process) · **process lifecycle/recovery** (who restarts the broker; what a broker crash does to the seat channel and to in-flight verbs) · **epoch-change linearization + in-flight disposition** under each placement · **push delivery routing** · **failure isolation** (blast radius of a worker crash vs an app-main crash). The choice stays YOURS unless the winning answer would alter a ratified topology or claim boundary — that case, and only that case, routes to the operator via master.

### F68 — you own the F65 conductor-identity PRODUCER contract (blocker; a missing owner edge, now assigned)
The ratified F65 split requires the exit-test record to carry the **conductor service build digest + governing config identity** for the relay-exchange leg — and no dispatched contract produced those bytes. As the conductor lifecycle/config host, **you author the canonical conductor-identity producer contract**: the exact **build-artifact identity** (what is digested, how) · the **governing-config identity** (the committed config chain / digest the running service is under) · the **canonical encoding** of both · **how the running service proves those values are the loaded ones** (binding the claim to the live process, not a bystander binary) · and the **relay-leg evidence reference** carried into the exit-test record (pointing at the conductor-captured observe-as-send E1/E2 records for the tested leg). Bounds: this **realizes** r7 F65 — it does NOT reopen it, and it adds **no conductor protocol/store field**. **m-3 consumes and confirms the scope boundary** (their supplement mirrors this edge) so the app/provider evaluator neither absorbs nor omits the separate relay-leg binding; **Master+VP retain the final composite exit-record join.**

### F69 — consumer-set correction
**m-2 joins your consumer-confirmation set** (alongside m-9, m-10, m-1): the shared transport half sits directly beside the m-2-owned mapping module and the parity boundary — m-2 must confirm your transport contract cannot absorb or strand the mapping layer.

No affected final-byte pair closure, complete consumer confirmation, `DESIGN_LOCK_ID`, or interface-lock readiness may be claimed until this supplement is consumed; your final DESIGN record carries the GRILL_LOCK. DESIGN-only authoring may continue meanwhile.

ACTIONS_GIT_REF: none — supplemental dispatch relay + one INDEX row only; the `…-041630` historical bytes untouched; no `frank/` edit, no code.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-7.planner consumes this supplement (grill first — durable GRILL_LOCK in the DESIGN record), authors the DESIGN incl. the F68 conductor-identity producer contract, then m-7.implementer final-byte review; consumer confirmations (m-9, m-10, m-1, m-2 + m-3 on the F65 scope edge) route on master's direction.
