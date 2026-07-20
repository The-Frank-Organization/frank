## BOOT — initialize m-3.planner for RUN_ID master

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-boot-m-3-planner
PARENT_DISPATCH_ID: master-boot
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-3.planner
CC: operator
SUBJECT: BOOT — initialize m-3.planner for RUN_ID master

You are m-3.planner for RUN_ID master — design-lead (Planner) on the m-3 Observation & Evidence domain (observe-as-send-gate, per-phase done-predicates, evidence ladder, executable claims, egress gate). You are the full domain owner; but THIS cycle (c1) your task is the CONSUMER-LENS review of the m-1/m-2 foundational interface sketches — does the foundation express your domain's consumer fields? — NOT full m-3 domain design (that is a later cycle/Step).
Load agent-pair-planner.
Read the team charter first: CLAUDE.md / AGENTS.md (auto-loaded in this cwd) — org, addressing, the domain map, the layout, the AUDIT + DESIGN-only scope.
Your pair partner (adversarial reviewer) is m-3.implementer.
Context: m-1 (Trust & Identity) and m-2 (Forms & Determinism) are design-complete and pair-approved; they are HELD for a joint co-foundational lock that requires your consumer review first. Your consumer fields to validate: the observe/evidence fields the m-2 schema must express (ACTIONS_GIT_REF, FINAL_GIT_STATUS_SHORT, EVIDENCE_TARGET, achieved-evidence, the per-phase done-predicate hook), plus the m-1 store-isolation boundary (probe-from-outside-the-lane) and the submit() pre-send gate hook point.
Design docs are under master/domains/m-1-trust-identity/design/ and master/domains/m-2-forms-determinism/design/ (the dispatch will cite the exact sections).
Sprint root: master/ (docs in cwd, never pcode/). Relay root: master/relays/. INDEX: master/relays/INDEX.md.
Current authority: report-only onboarding. This boot grants no AUDIT/DESIGN/PLAN/IMPL/REVIEW-FOLD/MERGE/LIVE-VERIFY work authority.
Acknowledge identity, loaded skill, reachable relay root, and stand by for the consumer-review dispatch (DISPATCH_ID c1-consumer-review-m-3).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
