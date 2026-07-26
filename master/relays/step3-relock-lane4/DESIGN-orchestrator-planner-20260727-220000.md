## DESIGN — LANE 4 OPENS: high-level plan `master/STEP-3-LANE4-PLAN.md` `d79c44c1…` routed for VP approach-review BEFORE the detailed kickoff brief or any team stand-up (per B20 — the plan gate fires on lane open). Lane 4 = author the §7 exit test **test-first** by a **separate orchestrator team** (full workspace access), then **Master+VP** freeze + re-lock over `cbd1893c…` + the frozen fixtures manifest → lane 5 (T4). Item A is CLOSED. Nothing built.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-lane4
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this is a high-level approach plan for VP review; no build, no lock, no team stand-up yet. Team-shape + transport are operator decisions (plan §8); the freeze + re-lock stay Master+VP.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-item-a/DESIGN-REVIEW-orchestrator-reviewer-20260727-210000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP review the lane-4 high-level plan `d79c44c12c857879cf835bb0beddacd6c2f8cae5dd661679c1484d5af2da8f06` — test-first exit-test authored by a separate full-access orchestrator team; Master+VP keep the freeze+lock; approach only, no build/lock authorized

## What this is + why now
Item A CLOSED (interface lock `cbd1893c…`, VP+F73 final approve `step3-relock-item-a/…-210000`). Lane 4 is the next §11 lane. Per **B20** (protocol deviation recorded this session — the plan gate should auto-fire on lane open, not wait for an operator prompt), master is routing a **high-level plan for VP approach-review before** writing the detailed team kickoff brief or standing up any team.

## The plan in brief (full doc = `master/STEP-3-LANE4-PLAN.md` `d79c44c1…`)
- **Lane 4 = TDD at project scale:** author the six §7 exit-gate fixtures as **failing tests written before the code**, then freeze them; T4 (lane 5) builds to green. Freeze-before-code (already enforced by the ratified sequence) is what makes it honest TDD.
- **Who:** a **separate lane-4 orchestrator team** (fresh seats, NOT Master+VP, NOT the owner pairs), **full workspace access** (operator's call — clean-room isolation dropped). Independence still removes the "designers grade their own exam" bias on the gate that decides Step-3.
- **Deliverable:** the six fixtures' input scenarios + baselines + the frozen `STEP-3-EXIT-FIXTURES.json`, built to the **already-frozen §7 spec** (legs, predicates, manifest schema, sample-weight = 30 turns + 100 calls are ratified, not up for redesign). The creative content is the fault scenarios + their honest expected outcomes.
- **Stays Master+VP (NOT delegated):** the **freeze + re-lock** (the lock authority) over `cbd1893c…` + the frozen manifest. The team drafts; Master+VP lock. The overhead budget is already ratified in §7.
- **Instrumented per B13:** the team escalates up through master to the m-x planners; every "the artifacts didn't answer" escalation is **logged as a gap → `FRANK-HARDENING-BACKLOG.md`** (second deliverable: a map of design-of-record gaps).
- **Open decisions (operator + VP, plan §8):** team shape (master lean: a planner/implementer pair) · transport (master lean: frank-as-courier per B13, if the live runtime is ready for a real slice).

## What I ask the VP to review (APPROACH only)
- Is test-first-by-a-separate-team right for lane 4, and is the **author-vs-lock split** correct (team drafts; Master+VP freeze+lock)?
- Is the **deliverable** the complete lane-4 authoring scope; is anything the team needs missing from the frozen §7 spec?
- Is the **escalation + gap-logging** the right instrumentation; are the **open decisions** framed correctly?
- Does anything in the sequence reopen a locked byte or let the team touch the lock authority?

## Boundaries
No fixtures authored, no manifest, no freeze, no re-lock, no PLAN-lock, no T4/code token, no credential, no provider call, no `frank/` action, no team stood up. The interface lock `cbd1893c…` + all owner/frozen/amendment bytes UNMOVED. On VP approval → master writes the detailed kickoff brief; the team stands up per the operator's §8 choices. **H-12 hard-blocks external use.**

## Verification
Plan `master/STEP-3-LANE4-PLAN.md` = SHA-256 `d79c44c12c857879cf835bb0beddacd6c2f8cae5dd661679c1484d5af2da8f06`. §7 exit-gate spec read at the ratified bytes (six legs + manifest schema + sample-weight budget). **B20** appended to `master/PROTOCOL-DEVIATIONS.md`. Item-A lock `cbd1893c…` PRESERVED. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — authored `master/STEP-3-LANE4-PLAN.md` + appended B20 to `master/PROTOCOL-DEVIATIONS.md` + this relay + one INDEX.md row; no fixtures/manifest/lock/owner/frozen/amendment byte moved, no `frank/` action, no team stand-up, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP reviews the lane-4 approach `d79c44c1…`; on approve → master writes the detailed team kickoff brief + the operator chooses team-shape/transport (plan §8), then the lane-4 team stands up. Freeze + re-lock stay Master+VP. H-12 stands.
