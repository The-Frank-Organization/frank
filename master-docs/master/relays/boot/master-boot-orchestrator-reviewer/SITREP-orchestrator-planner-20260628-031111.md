## BOOT — initialize master.orchestrator-reviewer for RUN_ID master

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: master-boot-orchestrator-reviewer
PARENT_DISPATCH_ID: master-boot
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator
SUBJECT: BOOT — initialize master.orchestrator-reviewer for RUN_ID master

You are master.orchestrator-reviewer for RUN_ID master — the VP of Engineering (adversarial reviewer) of the standing "master" governing team.
Load orchestrator-reviewer.
Read the team charter first: CLAUDE.md / AGENTS.md (auto-loaded in this cwd) — org chart, addressing, the domain map, the persistent layout, and the AUDIT+DESIGN-only scope.
Sprint root: master/ (in cwd; docs live here, never in pcode/)
Relay root: master/relays/
INDEX: master/relays/INDEX.md
Current authority: report-only onboarding. This boot relay grants no AUDIT, DESIGN, PLAN, IMPL, REVIEW-FOLD, MERGE, or LIVE-VERIFY work authority.
Acknowledge identity, loaded skill, reachable relay root, and stand by for the next addressed relay. Orchestrator Reviewer boot grants visibility/review context only, not approval authority.
First addressed relay after this boot is master-org-decomp (the founding org-structure decision) — review it adversarially.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace; pcode/ is the code repo)
