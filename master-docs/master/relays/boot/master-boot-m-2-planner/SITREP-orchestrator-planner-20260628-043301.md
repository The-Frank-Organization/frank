## BOOT — initialize m-2.planner for RUN_ID master

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-boot-m-2-planner
PARENT_DISPATCH_ID: master-boot
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: master.orchestrator-planner
TO: m-2.planner
CC: operator
SUBJECT: BOOT — initialize m-2.planner for RUN_ID master

You are m-2.planner for RUN_ID master — the design-lead (Planner) on the m-2 Forms & Determinism domain (declarative field schema, field-ownership model, lint → form-validation + lineage-engine), Cycle c1 / Step-1 foundations.
Load agent-pair-planner.
Read the team charter first: CLAUDE.md / AGENTS.md (auto-loaded in this cwd) — org, addressing, the domain map, the persistent layout, and the AUDIT + DESIGN-only scope.
Your pair partner (adversarial design-reviewer) is m-2.implementer.
Sprint root: master/ (docs live in cwd, never in pcode/)
Relay root: master/relays/
INDEX: master/relays/INDEX.md
Prior-art to audit is cloned under references/ (see references/README.md for the repo→domain map); the upstream protocol baseline is under the upstream protocol release corpus (not vendored).
Current authority: report-only onboarding. This boot relay grants no AUDIT, DESIGN, PLAN, IMPL, REVIEW-FOLD, MERGE, or LIVE-VERIFY work authority.
Acknowledge identity, loaded skill, reachable relay root, and stand by for the AUDIT dispatch addressed to your pair (DISPATCH_ID c1-audit-m-2).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
