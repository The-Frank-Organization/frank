## DESIGN dispatch — §7 stage-1 owned contract: the SHARED TRANSPORT/CLIENT boundary + the authenticated CHANNEL/BROKER contract with the F64 generation fence (per the RATIFIED MVP amendment r7 @ `2f75f2a1…`)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-mvp-design-m7
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a stage-1 owned interface contract under the ratified amendment; the operator gates return at the Master+VP interface-lock, not per-artifact
GRILL_REQUIRED: no — stage-1 owner contracts carry pair review + consumer confirmation; the grills ride the stage-4/5 build lanes (§7)
DESIGN_DOC_ID: step3-mvp-design-m7-transport-broker
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-041610.md
FROM: master.orchestrator-planner
TO: m-7.planner
CC: m-7.implementer, m-1.planner, m-2.planner, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, master.orchestrator-reviewer, operator
SUBJECT: author the m-7 stage-1 contract: hoist the shared conductor-client (transport half only) + author the channel/BROKER contract — broker-held logical-seat credential, placement OUTSIDE the replaceable worker generation, per-verb + push turn_epoch fence, epoch-change linearization + in-flight disposition — pair-reviewed final bytes, consumers confirm

m-7 — the Step-3 MVP amendment is **ratified + operative** (`master/STEP-3-MVP-AMENDMENT.md` r7, SHA-256 `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`; operator-ratified 2026-07-16, VP approve `…-035505`). Your charter carries the delta. Your conductor-host scope is untouched (the amendment changes no conductor byte/member/output); this dispatch opens your **§7 stage-1 owned contract** — two coupled artifacts.

### Author (you own these bytes; m-7.implementer pair-reviews the FINAL bytes)
1. **The shared transport/client boundary (§5):** hoist `Call`/reconnect/socket-lifecycle out of `cmd/frank-mcp` (package `main`) into the shared conductor-client at `internal/channel` — the TRANSPORT half only: the m-2-owned FieldSpec mapping must NOT be absorbed (it lands in an m-2 module; the 3-way seam is m-7 transport / m-2 mapping / m-9 consumer). Both frontends (the retained MCP server + the native tool) become thin skins over one client; MCP stays off the critical path; if it ships in Step-3 it is seat-scoped with no shared global credential multiplexing caller identities.
2. **The authenticated channel/BROKER contract (grill #3 + F64/F66), under m-1's identity semantics** (m-1's parallel stage-1 dispatch authors WHAT a logical seat is; you author HOW the channel/broker realizes it):
   - The broker holds the **LOGICAL m-9 seat's credential** — one credential per logical seat, never copied into worker generations; workers get an **epoch-bound, revocable USE capability**, never the bytes; m-10 launches/supervises with opaque references only.
   - **Placement:** the broker lives **OUTSIDE the replaceable worker generation** (F66) — beyond that constraint (own process vs protected thread/module in the app main process), placement is YOUR call in this DESIGN, made against the §2 hardening list.
   - **The generation fence (F64):** connect-time channel authorization alone is INSUFFICIENT — the live channel authorizes the credential once at connect and then serves the verbs without an app epoch (`frank/internal/channel/server.go:277-336,391-417`). The broker **checks the current worker generation / `turn_epoch` on EVERY `relay.submit`, `relay.project`, `relay.read` invocation AND on every push delivery/forwarding**. Specify the **epoch-change linearization** (the replacement's epoch increment ordered against in-flight broker calls) and the **in-flight-call disposition** at the change (complete-or-reject, recorded — never silent). m-10 supplies lifecycle/epoch state to the gate; it receives no credential bytes and gains no conductor verb.
   - Acceptance target (§10): after replacement, an old-epoch worker can invoke NONE of the three verbs through the broker nor receive/forward a new push.

### Boundaries
No conductor protocol/store change (the fence is broker-side — app-side — not a conductor gate). The F65 note stands: the conductor service identity (build digest + governing config identity) is bound separately in the exit-test record — never an app-release field. Consumers to confirm: m-9, m-10, m-1 (semantics fit). No DESIGN-lock, PLAN, T4 token, or code is authorized by this dispatch.

### Return path (§7 stage 1)
m-7.planner authors the DESIGN parented to THIS dispatch → m-7.implementer DESIGN-REVIEW as a uniquely-parented child (fresh review on any byte revision) → report-only SITREP to master naming the approved bytes + hash → consumer confirmations route on master's direction. The Master+VP interface-lock is the gate; no self-declared lock.

ACTIONS_GIT_REF: none — dispatch relay + one INDEX row only; no `frank/` edit, no code.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-7.planner opens the DESIGN on this dispatch (grounding: the ratified amendment §2b/§4-§7 + your charter delta + `internal/channel/server.go` + m-1's parallel `step3-mvp-design-m1` dispatch); pair review; report-only SITREP to master.
