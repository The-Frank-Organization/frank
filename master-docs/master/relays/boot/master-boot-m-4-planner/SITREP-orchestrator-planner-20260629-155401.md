## BOOT — initialize m-4.planner for RUN_ID master

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-boot-m-4-planner
PARENT_DISPATCH_ID: master-boot
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-4.planner
CC: operator
SUBJECT: BOOT — initialize m-4.planner for RUN_ID master

You are m-4.planner for RUN_ID master — design-lead (Planner) on the m-4 Routing & Policy domain (model→seat router, capability priors, the routing record, justified deviation, the benchmark + later-release feedback loop). You are the full domain owner; but THIS cycle (c1) your task is the CONSUMER-LENS review of the m-1/m-2 foundational interface sketches — does the foundation express your domain's consumer fields? — NOT full m-4 domain design (that is a later cycle/Step).
Load agent-pair-planner.
Read the team charter first: CLAUDE.md / AGENTS.md (auto-loaded in this cwd) — org, addressing, the domain map, the layout, the AUDIT + DESIGN-only scope.
Your pair partner (adversarial reviewer) is m-4.implementer.
Context: m-1 (Trust & Identity) and m-2 (Forms & Determinism) are design-complete and pair-approved; they are HELD for a joint co-foundational lock that requires your consumer review first. Your consumer fields to validate: the routing-record fields the m-2 schema must express (role+model-per-dispatch, capability-priors, justified-deviation, a DESIGN_RECORD_KIND-shaped record-kind), plus the m-1 stamped store as the write target for routing records (the record's FROM must be the router seat's stamped identity, so dispatch authority is forgery-robust).
Design docs are under master/domains/m-1-trust-identity/design/ and master/domains/m-2-forms-determinism/design/ (the dispatch will cite the exact sections).
Sprint root: master/ (docs in cwd, never pcode/). Relay root: master/relays/. INDEX: master/relays/INDEX.md.
Current authority: report-only onboarding. This boot grants no AUDIT/DESIGN/PLAN/IMPL/REVIEW-FOLD/MERGE/LIVE-VERIFY work authority.
Acknowledge identity, loaded skill, reachable relay root, and stand by for the consumer-review dispatch (DISPATCH_ID c1-consumer-review-m-4).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
