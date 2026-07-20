## RECONCILE — the F39–F44 required return: an EXPLICIT architecture amendment (`master/STEP-3-MVP-AMENDMENT.md` @ SHA-256 `a524bcbf…`), not a claim that r4 is architecturally unchanged. It supersedes the named r4 clauses (F39); names the honest governance boundary — comms + provider-egress governed, local tool effects NOT (F40); pins the static allow-list as an operator-fixed constant m-10 hosts-not-owns (F40); chooses branch (a) — frank owns the provider client, API-key to start (F41, operator-decided: "a harness comparable to the others"); routes the shared client to m-2 + m-7 (F42); pins the m-9→m-10 wake handoff (F43); and presents ONE non-contradictory dependency graph (F44). Requesting exact-byte VP review; operator ratifies on approve

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this amendment supersedes normative r4 clauses (F39); per the VP it requires operator ratification of the final bytes after VP exact-byte review. The substantive choices (branch-a harness, native tool) are operator-decided; this relay transmits the authored bytes for review
GRILL_REQUIRED: no — the amendment defines interface targets; the first-stage DESIGNs (m-8/m-9/m-10) carry their own grills; any Step-4 ceiling work re-opens the m-5 grill
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260715-161118.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-7.planner, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: F39–F44 folded into an explicit hashed amendment `STEP-3-MVP-AMENDMENT.md` @ `a524bcbf…` (branch-a real harness, native tool, empty seam) — requesting exact-byte VP review; on approve, operator ratifies + the §7 graph reconciles into every operative source

Partner — `161118` accepted whole; every finding was correct. I own the F39 category error: I labeled a **material architecture change** ("packet r4 untouched") as a bounded fold. **The correction is an explicit, separately-hashed amendment** — `master/STEP-3-MVP-AMENDMENT.md` @ **SHA-256 `a524bcbf0f248bbd569ff8118e857b831aa6a11d899363b846539b6b7f542f26`** — that names what it supersedes and pins the new boundary/graph. **r4 stays byte-exact `2d240eb6…` as the historical lock.**

### How the amendment answers each finding
- **F39 — supersession is explicit** (§1): a 4-row table naming the r4 clauses (`:27-29,65-71` ceiling · `:102-112` m-10+m-5 first stage · m-8/m-9-post-lock · `:65-75,112` MVP-turn) and what replaces each. r4 file bytes unchanged; this is additive + separately hashed.
- **F40 — honest boundary + allow-list provenance** (§3, §4): the MVP **governs relay comms + provider egress (E3); it does NOT govern local tool effects** (no ceiling/sandbox/audit-gate — Step-4). The static allow-list is an **operator-fixed ratified CONSTANT** (the exact 5 local tools + 3 relay verbs), provisioner-written, `run_id`-bound, immutable, deny-on-absent; **m-10 hosts enforcement, owns no policy** (preserves r4 `:28,103`). Product claim narrowed to the exact governed surface.
- **F41 — the provider vertical, chosen** (§2): **branch (a)** — operator-decided "a harness comparable to the others." frank owns the provider client (m-8 = minimum OpenAI-compatible, **API-key** to start; OAuth/subscription = Step-4). The vertical m-9→m-8→provider with **m-1 key custody (never in a transcript)** + **m-3 E3 evidence** is IN the first stage. (Empirical basis: opencode/pi/jcode all own the provider client; jcode/pi's OAuth path is the real-work Step-4 shape.)
- **F42 — shared client is a 3-way seam** (§5): **m-7** the transport/client boundary, **m-2** the form→tool-schema + submit-payload mapping (FieldSpec semantics — I had omitted m-2; now routed + CC'd), **m-9** the consumer. Package split required (channel ≠ FieldSpec policy); parity tests (MCP ≡ native) required; owners land before the consumer locks.
- **F43 — the wake handoff** (§6): the push reaches only the **connected m-9 seat** (m-10 has no credential, can't start an absent worker), so the MVP lifecycle is a **resident m-9 receiver → app-IPC nudge → m-10 schedules a turn**, with **polling the fallback**; reconnect/coalescing/dup-nudge pinned in DESIGN. No "m-10 turns the push into a turn" claim.
- **F44 — one dependency graph** (§7): a single current first stage — **BUILD** {m-10, m-9, m-8} + **GATE** {m-1, m-3, m-2, m-7}; **m-5 stood down**; m-4/m-6 out. It **supersedes** the m-10+m-5/m-9-stage-2 framing; **the source reconciliation folds on ratification** (per your directive — I did not re-fold the 7 sources pre-ratification; they remain the proposed input you flagged, now pointed at this amendment).

### The honest scope consequence (surfaced to the operator, accepted)
Choosing branch (a) grows the first stage from "m-10+m-9" to the harness floor: 3 build lanes + 4 contract/gate lanes. "Barely enough" is now **depth** (one provider, API key, no ceiling/sandbox/routing/carousel), not breadth. The operator accepted this as the cost of a real standalone harness.

## Verification
- **Amendment SHA-256:** `a524bcbf0f248bbd569ff8118e857b831aa6a11d899363b846539b6b7f542f26` (`shasum -a 256 master/STEP-3-MVP-AMENDMENT.md`).
- r4 unchanged: `2d240eb69b2f50dd47aca9c3385dca10ab144c577e8e50959f057791b72dca69`. Canonical m-5 (Step-4 ceiling basis) unchanged: `643dd7c2940e32b96b2a9e80392e91d781fe0b5b40bfe54b0a7c1d76189d4ebf`.
- The 15-file governing manifest is unchanged this pass (`5374ee4a…`) — the amendment is a NEW standalone file, not folded into the sources yet (F44 reconciliation is on ratification).
- `frank/` finding citations in the amendment verified read-only (`cmd/frank-mcp/schema.go`, `internal/channel/server.go`, r4 clause lines); `frank/` clean at `main@502e06c`.
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260715-163000.md` — run below.

ACTIONS_GIT_REF: docs-workspace disk action only — created `master/STEP-3-MVP-AMENDMENT.md` (candidate) + this transmittal relay + its INDEX row; a one-line pointer added to `master/README.md`. Read-only `frank/` inspection. No packet / r4 / canonical-m-5 / historical-relay edit; no code, credential, provider, lock, PLAN, or T4 action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: VP performs exact-byte review of `STEP-3-MVP-AMENDMENT.md` @ `a524bcbf…`; on approve, the operator ratifies (operator-authored), then master folds the §7 graph into every operative source + marks r4's §1 superseded clauses, and the first-stage DESIGNs (m-8/m-9/m-10 build + m-1/m-2/m-3/m-7 gates) dispatch. No lock/DESIGN-lock/PLAN/T4/credential/provider authorized until then.
