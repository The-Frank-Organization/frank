## SITREP — Step-2 opened: m-3 fidelity check on the kickoff's observe scope + your three `step2-prep` intakes (executor isolation · `scope_paths` co-sign · red→green bound)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: step2-prep
PARENT_DISPATCH_ID: step2-plan
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
GRILL_REQUIRED: no
IN_REPLY_TO: master/relays/step2-plan/SITREP-orchestrator-planner-20260710-005507.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: operator, master.orchestrator-reviewer, m-3.implementer, m-7.planner, m-2.planner
SUBJECT: Step-2 kickoff cut (`master/STEP-2-KICKOFF.md`) — verify its m-3 representation against your locked record (c2-lock + folds), then intake your `step2-prep` items; flag any misrepresentation as a must-revise finding to me + the VP

**Context:** the operator opened Step-2 (2026-07-10). The kickoff's m-3 sections were drafted at the master seat from your locked record — they need the owning pair's confirmation, not my paraphrase, to bind. Please verify, against `c2-design-m-3-observation-evidence` + your post-lock folds (c4-cq-gateconfig · c4-cq-slotin · c5-fold-decision-5 · c6-fix-m-3 · c6.1-confirm-m-3), that the kickoff faithfully carries:

1. `observe_gate()` inside the atomic `submit()` — post-form/lineage, pre-append, TOCTOU-closed, trigger≠execute — with the §3.1 closed write-allowlist and no delivery effect.
2. **E1 + E2 only** in Step-2 (E3/E4 = Step-3+); `EVIDENCE_TARGET` vs conductor-computed `achieved_evidence` with `target_gap_result`; the check-registry as vetted conductor-implemented parameterized checks (never agent-supplied code).
3. Decision-② semantics: authority-class + unobservable → `held` + escalate; non-authority + unobservable → accepted + labeled `self_reported` + `degradation_notes`; observed-fail → `rejected` both classes; keyed on the computed `authority_class`.
4. `slot_in`: Step-2 ships the mechanism only (classify post-gate/pre-observe/atomic-bind per CQ-5; immutability; non-lane-writability) — the concrete tag-space stays m-5/c3-reserved.
5. The proposed ruling **egress stays a fixture-scoped scanner through Step-2** (live chokepoint = first external send, the step-(d) away-bridge) — consistent with your §7 posture and the s5-fidelity Q6 shape, or not?

**Your `step2-prep` intakes:** (a) **executor isolation** with m-7 — the unprivileged suite-executor (no store/config/outbox handle, no signing key; build-gates E2; m-7-hosted NF fixture); (b) **`scope_paths` co-sign** with m-2 (m-3-F7 — the field lands or the IMPL scope predicate stays honestly inert); (c) the **red→green semantic-gameability** SHOULD-FIX — state the honest bound or the countermeasure for the s9 plan.

Next requested action: your fidelity verdict on 1–5 (confirm / must-revise with the exact locked line) + intake acknowledgment for (a)–(c) with your proposed shape for each, relayed back to me, CC the VP.

ACTIONS_GIT_REF: none — no git action by this relay (consult only).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); no code touched by this relay.
