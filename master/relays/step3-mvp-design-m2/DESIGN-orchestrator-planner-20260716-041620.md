## DESIGN dispatch — §7 stage-1 owned contract: the form→tool-schema MAPPING module + the F58 relay-verb schema digests / mapping version + MCP↔native parity vectors (per the RATIFIED MVP amendment r7 @ `2f75f2a1…`)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-mvp-design-m2
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a stage-1 owned interface contract under the ratified amendment; the operator gates return at the Master+VP interface-lock, not per-artifact
GRILL_REQUIRED: no — stage-1 owner contracts carry pair review + consumer confirmation; the grills ride the stage-4/5 build lanes (§7)
DESIGN_DOC_ID: step3-mvp-design-m2-mapping
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-041610.md
FROM: master.orchestrator-planner
TO: m-2.planner
CC: m-2.implementer, m-7.planner, m-9.planner, m-9.implementer, m-10.planner, master.orchestrator-reviewer, operator
SUBJECT: author the m-2 stage-1 contract: the form→tool-schema + submit-payload mapping as an m-2-owned module; the relay-verb schema digests + mapping version (F58); MCP↔native parity vectors — pair-reviewed final bytes, consumers confirm

m-2 — the Step-3 MVP amendment is **ratified + operative** (`master/STEP-3-MVP-AMENDMENT.md` r7, SHA-256 `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`; operator-ratified 2026-07-16, VP byte-bound approve `step3-arch-packet/…-035505`). Your charter carries the delta section. This dispatch opens your **§7 stage-1 owned contract**.

### Author (you own these bytes; m-2.implementer pair-reviews the FINAL bytes)
1. **The mapping module contract (§5):** `SchemaFromForm` / `SubmitPayloadFromArguments` (+ re-render semantics) encode m-2-owned FieldSpec/form semantics (`frank/cmd/frank-mcp/schema.go:11-47,90-129` is the current trapped implementation — prior art for the hoist, not the contract). The mapping lands in an **m-2-owned module**; the m-7 shared transport client (`internal/channel`) must NOT absorb it — the 3-way seam is m-7 transport / m-2 mapping / m-9 consumer.
2. **The F58 build-identity components you produce:** the **relay-verb tool-schema digests** (`relay.submit`, `relay.project`, `relay.read`) + the **form→schema mapping VERSION** — canonical encoding + field applicability defined in this DESIGN (local tools carry NO mapping version; that's m-9's half). Your components feed the per-tool identity vector `{canonical name, tool-schema digest, implementation/catalog version, mapping version}` that binds at the Master+VP interface-lock and is mechanically verified at the post-build release-binding (F63).
3. **Parity conformance vectors (§5/§10):** shared vectors under which the retained MCP frontend and the native tool produce **equivalent conductor calls + re-render behavior** (the §10 "native == MCP relay path" row).

### Boundaries
No FieldSpec registry byte moves in this DESIGN (any registry need routes back as its own reviewed changeset per your standing pattern). No conductor byte/member change. Consumers to confirm: m-9 (the native-tool consumer) + m-7 (the transport boundary you sit beside). No DESIGN-lock, PLAN, T4 token, or code is authorized by this dispatch.

### Return path (§7 stage 1)
m-2.planner authors the DESIGN parented to THIS dispatch → m-2.implementer DESIGN-REVIEW as a uniquely-parented child (fresh review on any byte revision) → report-only SITREP to master naming the approved bytes + hash → consumer confirmations route on master's direction. The Master+VP interface-lock is the gate; no self-declared lock.

ACTIONS_GIT_REF: none — dispatch relay + one INDEX row only; no `frank/` edit, no code.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-2.planner opens the DESIGN on this dispatch (grounding: the ratified amendment §4/§5/§7 + your charter delta + `cmd/frank-mcp/schema.go`); pair review; report-only SITREP to master.
