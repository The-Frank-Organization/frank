## SITREP — fence ruling: root README.md IN-FENCE for the bounded S2 honesty delta (S1 ASK-1 precedent applied)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s2-core-plan
PARENT_DISPATCH_ID: s2-core-design
RUN_ID: s2
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: s2-core-plan/PLAN-REVIEW-implementer-20260704-031243.md
FROM: s2.orchestrator-planner
TO: s2-core.planner
CC: s2-core.implementer, s2.orchestrator-reviewer, operator
SUBJECT: Task-13-step-3 ASK answered ahead of the completion SITREP — README.md ruled IN-FENCE for one bounded claim-honest S2 delta; SCOPE_DIFF cites this relay; the sole remaining dispatch gate is the m-1 fidelity approve

Ruling (orchestrator's call per protocol, mirroring the S1 ASK-1 ruling verbatim in rationale): **root `README.md` is IN-FENCE for S2**, bounded to the claim-honest S2 delta — what S2 adds (the owed-item projection with the materialize-first scope stated; GC posture: journals-only, off by default, canonical records never touched; the Ready/Diagnostics recovery posture) — because the README is the repo's honesty/claim surface, S2-SWEEP-covered, and claim honesty must hold in code AND docs. Conditions, same three as S1: (i) the delta stays within the claim-honesty pins (exactly-once EFFECT; GC locked-strength wording; materialize-first; D5 beside exclusivity claims); (ii) no claim of anything unbuilt or dormant as live (render compaction dormancy stated if GC is described); (iii) the README row enters the plan's file list + SCOPE_DIFF citing THIS relay as its fence evidence. The PROCEED-TO-PLAN fence (`s2-core-plan/PLAN-orchestrator-planner-20260704-023128.md` condition 2) is amended accordingly — this relay is the amendment record.

Observed state prompting the early answer (all read on disk this session, E1): gated PLAN r2 pair-APPROVED (`PLAN-REVIEW-implementer-20260704-031243.md`, PLAN_LOCK_ID `s2-slice-2-plan` at main@c16f261; F1 quarantine-decomposition fold verified closed by the reviewer); the plan's Task 13 step 3 defers the README ASK to the completion SITREP — answered here instead; no other item waits on this seat.

Standing gates unchanged: m-1 fidelity approve on record in `.relays/s2/` (condition 3 — packet out, no verdict yet, reviewer-confirmed empty dir at review time) → then your mechanical SCOPE_DIFF (all-in, README row citing this relay) → then the delegated dispatch under the standing conditions. Merge/S2-close remain the operator's separate gate.

ACTIONS_GIT_REF: none — ruling relay; this file + an INDEX row under gitignored .relays/ (ledger entry rides its own commit).
FINAL_GIT_STATUS_SHORT: none — clean tree
