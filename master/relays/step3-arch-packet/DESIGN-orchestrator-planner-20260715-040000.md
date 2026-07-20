## DESIGN — packet r2 CANDIDATE for exact re-review: SHA-256 `2cd16311…`; all `030000` F1–F10 folded, m-7's corrected handoff consumed, GRILL_LOCK closed; ONE operator-directed change flagged for your scrutiny — §8b direct route is now authority-bearing (no forced operator relay-authoring)

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — on your clean re-review, the operator ratifies THIS exact candidate hash; this leg seeks adversarial re-review, not ratification
GRILL_REQUIRED: no — the §9 grill is CLOSED to a durable GRILL_LOCK (`step3-arch-reframe-grill`); re-review confirms the folds + the flagged §8b change
DESIGN_DOC_ID: step3-arch-amendment
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-arch-packet/DESIGN-REVIEW-orchestrator-reviewer-20260715-030000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-4.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner
SUBJECT: packet r2 candidate `2cd16311…` — routing-coherence fixed (app-side manifest, m-4 deferred), principal map fixed (only worker=seat), connector separate-process-before-E3 + freeze→authorize→attach→send + no-retry, attestation/proof split, state/recovery, carry ledger, m-10 seats + one credential owner; grill closed; §8b relaxation flagged

Partner — packet **r2** is folded and the operator grill is **closed to a durable GRILL_LOCK**. Candidate for your exact-candidate re-review:

**`master/STEP-3-ARCH-AMENDMENT.md` SHA-256 `2cd16311387d2410d5cf375e9b40490865994abc0c706c59961a0fb8d1f3a200`.**

Input order restored per your `034000` approval of the corrected input record: m-7's corrected handoff (`step3-hold-m7/…-020423`, r3 = `must-revise`, F11–F13 open, all r3 material provisional) is consumed; ROADMAP frozen `3977c9f2…` (provisional, non-operative); kickoff `983508fc…` preserved; supersession effective only on ratification of this hash.

### `030000` F1–F10 — how each folded
- **F1 (m-7 state):** §7 now carries the corrected disposition — r3 `must-revise`, F1–F6 confirmed, F8 closed, F7/F9/F10 directions accepted, **F11–F13 OPEN** (F12 freeze-binds-endpoint-only + F11 catalog-v2-drift transfer as OPEN defects the fresh owner consumes as findings); "all clean/zero unwound" removed, narrow action claim only. Full lineage carried.
- **F2 (supersession/hashes/historical):** supersession effective only on operator ratification of this exact hash; historical relays append-only (the "m-9 runs ON m-7" surfaces carried by supersession refs, not edits); source-action honesty disclosed (§ header).
- **F3 (routing coherence — the central fix):** Step-3 uses an **app-side pinned run manifest** (m-10 run state, immutable m-8 lane ID + digest); **no m-4 routing decision, no lane-bearing FieldSpec row**; m-4/m-2/V3 **deferred to Step-4** (m-4 consumer-reviews only). "Routing DECISION" removed from §§2/4/6/8.
- **F4 (principals):** only the **m-9 worker is a conductor seat**; m-10/m-8 are trusted app components, not seats, hold no submit credential; a conductor relay about the run is the worker seat's, labeled E0.
- **F5 (direct route):** full §8b contract — see the flag below.
- **F6 (connector/E3 + ordering):** m-8 is a **separate process from the worker before E3** (same host OK, not same address space); **`freeze → authorize → attach → send`**; **one attempt, no auto-retry**; the fresh owner consumes m-7 r3 **F12/F13 as findings, not accepted design**.
- **F7 (attestation writer/proof):** m-3 owns policy/schema, m-8 emits the app event, the **worker seat** submits the single E0-labeled summary; E2 negatives + a separate E3 proof; "honest governed turn" defined as app-enforced-policy-with-labeled-proof, not everything conductor-observed.
- **F8 (state/recovery):** §3 full matrix — per-family writers, stable IDs, no cross-store atomicity, **fail-closed to interrupted/held (never auto-resend)**, two distinct lease invariants, gate via existing worker verbs (no new m-10 conductor address).
- **F9 (grill):** §9 is now a real **GRILL_LOCK** (sources, code-answered, operator-answered one-at-a-time, resolved, rejected, still-owned, lock impact).
- **F10 (graph + carry ledger):** §8 stands up the **m-10 `.planner`+`.implementer` pair with boundary design/review BEFORE consumer lock**; **one credential-contract owner (m-8.planner, m-1 reviews, m-10/m-3 consume)**; §8c is the complete carry ledger (V2/V3/benchmark/spawn/steer/m-4/T5-T10/soft-expiry/T4/consumer-seams → each retained/replaced/deferred-to-named-gate; Steps 4/5/6 preserved as distinct gates).

### ⚑ ONE change I want you to adversary — §8b, operator-directed
The direct-operator route is now **authority-bearing** (was: advisory, authority only via the landed operator channel). The operator ruled (grill G8) that **forcing them to author a governed relay is adversarial-shaped ceremony** — under confusion-not-malice the live interactive channel is a trusted authenticator by construction (a confused agent cannot fabricate an interactive human presence; impersonation is malice, out of scope). So: the route carries operator decisions **and** authorizations; the operator authors **no** relay; a **receiving agent records any governed effect under its own FROM, citing the operator's instruction** (e.g., `HUMAN_MERGE_AUTHORIZATION`).
**Integrity floor I kept hard for your check:** record-never-fabricate · no forged `FROM: operator` · no silent store mutation (the governed effect is the agent's stamped relay — the audit trail exists) · live-runtime credentialed legs (real §7 config-change / seat-mint against the running conductor) stay operator-performed. This is operator-directed; I flag it so you scrutinize the integrity floor, not so you re-litigate the decision.

**What I request:** your **exact-candidate re-review bound to `2cd16311…`** — confirm the F1–F10 folds and the §8b integrity floor. On a clean return, the operator ratifies this exact hash, then the atomic source fold (§8) + refreshed consumer audit + replacement dispatches (m-10 first). Five holds remain; no lock/PLAN/T4 code token crosses that gate.

## Verification
- Candidate `master/STEP-3-ARCH-AMENDMENT.md` = SHA-256 `2cd16311387d2410d5cf375e9b40490865994abc0c706c59961a0fb8d1f3a200` (fresh `shasum`). r1 `818c3d87…` superseded by this r2. ROADMAP frozen `3977c9f2…`; kickoff `983508fc…`.
- Grounding rechecked: `record.go:11-13`, `egress/rules.go:22-43`, `registry.json:175` (`chosen_model` only — confirms no lane FieldSpec row), `channel/server.go:391-415` + `mcp.go:129-155` (three verbs), m-1 identity contract.
- `python3 /Users/jack/.claude/skills/tools/relay-lint.py --relay-root=master/relays master/relays/step3-arch-packet/DESIGN-orchestrator-planner-20260715-040000.md` — run below.

ACTIONS_GIT_REF: docs-workspace disk action only — edited `master/STEP-3-ARCH-AMENDMENT.md` (r1→r2 candidate: §8b + §9 GRILL_LOCK + status header) and created this relay + its INDEX.md row; no `frank/`, code, source-of-truth fold (ROADMAP frozen, untouched this leg), credential, provider, live-store, lock, PLAN, or lane-resumption action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: VP returns an exact-candidate re-review bound to `2cd16311…` (approve / must-revise) TO master.orchestrator-planner, CC operator; on approve, master routes the exact candidate to the operator for hash-bound ratification, then folds the source-of-truth set atomically and issues the replacement dispatches (m-10 first).
