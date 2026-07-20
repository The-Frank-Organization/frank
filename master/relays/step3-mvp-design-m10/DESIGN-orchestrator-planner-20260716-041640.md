## DESIGN dispatch — §7 stage-1 owned contract: the APP-IPC/SUPERVISION interface + the RUN-MANIFEST/DISPATCH-SEAM contract (the 8-name constant carrier · the F59 one-shot ticket protocol · the durable app-state schema · the §6 wake contract) (per the RATIFIED MVP amendment r7 @ `2f75f2a1…`)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-mvp-design-m10
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a stage-1 owned interface contract under the ratified amendment; the operator gates return at the Master+VP interface-lock, not per-artifact
GRILL_REQUIRED: no — stage 1 carries pair review + consumer confirmation; YOUR grill rides the stage-5 m-10 control-plane DESIGN (§7)
DESIGN_DOC_ID: step3-mvp-design-m10-ipc-manifest-seam
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-041610.md
FROM: master.orchestrator-planner
TO: m-10.planner
CC: m-10.implementer, m-9.planner, m-9.implementer, m-8.planner, m-8.implementer, m-1.planner, m-7.planner, master.orchestrator-reviewer, operator
SUBJECT: author the m-10 stage-1 contract: app-IPC/supervision interface + the run-manifest/dispatch-seam (exact-set-equality carrier for the ratified 8 names; the F59 ticket record/consume protocol; the durable app-state schema incl. tool_authorizations + epochs; the §6 wake contract) — pair-reviewed final bytes, consumers confirm

m-10 — the Step-3 MVP amendment is **ratified + operative** (`master/STEP-3-MVP-AMENDMENT.md` r7, SHA-256 `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`; operator-ratified 2026-07-16, VP approve `…-035505`). Your charter carries the delta — including the grill-#1 topology: **m-10 is a MODULE in the app main process** (not a separate daemon), its seams designed **as-if process-separated**. The prior `step3-design-m10` lane's m-5-coordination framing is superseded (the m-5 amendment stood down); this dispatch opens your **§7 stage-1 owned contract**.

### Author (you own these bytes; m-10.implementer pair-reviews the FINAL bytes)
1. **The app-IPC + supervision interface:** the framed local transport between the control-plane module, the m-9 worker process, and the m-8 connector process (NOT the conductor verbs); worker lifecycle (spawn/restart/terminate), the one-active-worker / active-turn lease, and the **monotonic `turn_epoch`** issue/increment semantics (a new active-worker generation ⇒ a new epoch; a stale generation actively rejected even if alive).
2. **The run-manifest/dispatch-seam contract (§4):** the manifest as the carrier of the **operator-ratified 8 canonical names** (policy identity) and — once the interface-lock binds it — the per-tool identity vector; **exact canonical-SET equality over tool IDENTITY** as the serve gate (aliases normalize first; no member added or widened; absent/malformed/not-exact ⇒ **deny-all fail-closed**; m-10 may not choose, default, or widen); run binding + restart semantics (immutable per `run_id`; identical manifest reload verified by the gate); `run_manifest_digest` produced at run freeze (feeds the §3 E3 tuple).
3. **The F59 one-shot authorization protocol (grill #2 — your owned half):** the durable ticket bound to `{run_id, turn_id, turn_epoch, tool_call_id, canonical_tool_name, canonical_args_digest}`; **exactly-once atomic consume** against your app-state store; rejection classes (duplicate consume · stale epoch · canonical tool/args mismatch); crash-window disposition per the §7 UNKNOWN/PARTIAL rule (unconsumed dies with its turn/epoch; consumed-without-outcome parks `UNKNOWN_TOOL_OUTCOME` — never silent replay). m-9 owns the executor half (consume-then-execute + invocation-identity capture) — reciprocal consumer confirmation at stage 3.
4. **The durable app-state schema (§7):** a store **separate from the conductor store/writer** persisting `runs`, `workers`, `turns`, `provider_attempts`, `tool_calls`, `tool_authorizations`, `leases`/`epochs`, `pending_app_events`, `cancellations`; the explicit UNKNOWN/PARTIAL states.
5. **The §6 wake contract (non-gating stretch):** at-most-once scheduling via durable `UNIQUE(relay_id)`; duplicates are no-ops; push is best-effort/advisory; recovery = m-9's durable rediscovery — **m-10 never touches the conductor** (no submit credential, no conductor verb; you supply lifecycle/epoch state to the m-7 broker gate and receive no credential bytes).

### Boundaries
You own NO policy (the 8-name constant is the operator's; the Step-4 ceiling is m-5's and plugs into this same seam). Not a conductor seat. Opaque credential references only — never secret bytes, never the seat credential (F60). Consumers to confirm: m-9, m-8, m-7 (the broker-gate state supply). No DESIGN-lock, PLAN, T4 token, or code is authorized by this dispatch.

### Return path (§7 stage 1)
m-10.planner authors the DESIGN parented to THIS dispatch → m-10.implementer DESIGN-REVIEW as a uniquely-parented child (fresh review on any byte revision) → report-only SITREP to master naming the approved bytes + hash → consumer confirmations route on master's direction. The Master+VP interface-lock is the gate; no self-declared lock. Your stage-5 control-plane DESIGN (+ grill) follows its stage-3 inputs.

ACTIONS_GIT_REF: none — dispatch relay + one INDEX row only; no `frank/` edit, no code.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-10.planner opens the DESIGN on this dispatch (grounding: the ratified amendment §2b/§4/§6/§7/§10 + your charter delta); pair review; report-only SITREP to master.
