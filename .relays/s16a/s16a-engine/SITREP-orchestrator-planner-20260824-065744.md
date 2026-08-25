## SITREP — THE VERSION HOLD IS LIFTED: resume at SCOPE_DIFF → token → branch → WP1

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s16a-engine-notice
PARENT_DISPATCH_ID: step3-t4-preflight
RUN_ID: s16a
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the operator's ruling ("new build is intentional") and restart words are on the master trail; your next human gate remains the terminal MERGE-GATE
GRILL_REQUIRED: no
FILED_AT_LOCAL: 20260824-065744
IN_REPLY_TO: s16a-engine/SITREP-orchestrator-planner-20260824-052218.md
FROM: master.orchestrator-planner
TO: s16a.planner, s16a.implementer
CC: master.orchestrator-reviewer, operator
SUBJECT: hold lifted — daemons restarted on the operator-ruled build (fp ddd5f9f5…); Codex plugins updated by the operator; r7 stands approved; SCOPE_DIFF then the token per the grant

1. The brief engine hold is OVER. What happened, for your trail: the operator's plugin reload re-shipped 2.9.1 at a new build fingerprint (`ddd5f9f5c84e…`); the engine correctly refused mixed-identity submissions; the operator RULED the new build intentional; both daemons (this root + master/relays2) were restarted onto it, state ready, pending_renders 0, epoch active. The operator has ALSO updated the Codex-side plugins, so BOTH your seats' clients now match the daemon — verify with `relay version` (expect kit 2.9.1, fingerprint `ddd5f9f5…`) before your first submit, and report any mismatch to master rather than retrying.
2. Nothing in your round moved during the hold and nothing re-opens: PLAN r7 @ `ff9838cdc07bc632763d9038155e27b7c18e3c14edd8ea4fefab54cbd1a1ff4f` stands formally APPROVED (`s16a-build/PLAN-REVIEW-implementer-20260824-053125.md`, engine-admitted). Resume exactly at your own sequence: SCOPE_DIFF all-in → the pair planner issues the BUILD TOKEN per the commission grant → branch `s16a-conformance` from `ff1193d7…` → WP1 ONLY (the 64-test battery; launch census expected 20 GREEN / 44 RED; a gate-row GREEN at first contact is a FINDING).
3. Practice reminders now that the engine is the record: every relay through `.engine/drafts/<your-address>/` + `relay submit --key <your key>`; SUBJECT mandatory; include FINAL_GIT_STATUS_SHORT (or an explicit unavailable reason) — the engine ADMITS without it but the linter flags it, a gap master hit on his own notice; a dispatch id is single-use unless you thread with --admits-against.
FINAL_GIT_STATUS_SHORT: (literal at draft authoring; the engine ledger is the authority for admission state:)
 M master/relays/CHECKPOINTS.md
