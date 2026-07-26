## DESIGN-REVIEW -- MUST-REVISE-NARROW: B21's full-team choice stands, but the nested address, lineage, review, and preflight contracts are not yet executable

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-vp-design-review-r6
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- the operator's B21 team-shape decision stands; the operator also retains preflight-only boot authorization and the separate post-pass activation
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-lane4-staffing-grill-1
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260728-080000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: Preserve the operator-directed full nested lane-4 team and all rev5 mechanics; return one rev7 that makes the nested seats, immediate-predecessor lineage, reviewer gate, and current-generation spawn/preflight carrier exact

VERDICT: revise

Review-routing target: `master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260728-080000.md` at SHA-256 `2447ab6e140a988f6819236efc16e0c6fc4f4c9d86f9a8cee3abe20631237208`.

Plan reviewed: `master/STEP-3-LANE4-PLAN.md` rev6 at SHA-256 `959a29aaccf3f23910cf237746acaabde6d94457e68db3937c185c1c7b329ff9`.

Upstream lock preserved: `master/STEP-3-INTERFACE-LOCK.md` at SHA-256 `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.

## Findings

### LANE4-VP-R6-F1 -- GATE: the proposed seat names do not encode the roles the plan assigns

Plan lines 20-22 call `l4.planner` an **Orchestrator Planner**, call `l4.reviewer` an **Orchestrator Reviewer**, and name workers only as `l4.w<k>`. That is not a complete relay address map:

- protocol ROLE/FROM discipline requires `ROLE: Orchestrator Planner` to use a `.orchestrator-planner` address suffix and `ROLE: Orchestrator Reviewer` to use `.orchestrator-reviewer`;
- `.reviewer` is the distinct `ROLE: Reviewer`, while `.planner` is `ROLE: Planner`;
- `l4.w<k>` has no recognized role suffix and cannot be an actual `FROM` or `TO` address; and
- the proven nested-team precedent uses `s1.orchestrator-planner`, `s1.orchestrator-reviewer`, `s1-core.planner`, and `s1-core.implementer`, preserving role-stamped identity at every tier.

This is not naming polish. Frank's registry gives orchestrator-specific authority by the `*.orchestrator-planner` suffix, and durable exported relays would otherwise fail ROLE/FROM consistency or silently run with ordinary Planner/Reviewer authority.

Required correction:

1. pin one canonical seat/role/skill map in the plan and B21, preferably `l4.orchestrator-planner` + `l4.orchestrator-reviewer`;
2. define the concrete expansion for every worker, such as `l4.w1.planner` / `l4.w1.implementer` when paired, with the role each single-seat worker uses;
3. preserve the operator's requested `l4.*` namespace while keeping the final suffix canonical; and
4. use those same exact addresses in the preflight, escalation, GRILL_LOCK, and future kickoff.

If `l4.planner` / `l4.reviewer` are intentionally retained, then their roles must be Planner / Reviewer and the plan must stop claiming they are the orchestrator pair. Do not mix the two models.

### LANE4-VP-R6-F2 -- GATE: `PARENT_DISPATCH_ID` is overloaded as a static tier-parent field

Plan line 29 says every relay's `PARENT_DISPATCH_ID` points one tier upward. The protocol defines that field as the **immediate predecessor dispatch edge**. Those are not equivalent:

- the l4 review of a decomposition must parent to the l4 decomposition relay, not directly to master;
- a worker result must parent to its addressed worker dispatch;
- the integrated-manifest review must parent to the integration/review request; and
- an escalation response must parent to the escalation it answers, even when both are within one tier.

The current convention would erase the within-tier review/dispatch chain while claiming full lineage. It also ignores the known shared-dispatch-id resolver defect recorded in `CYCLE-PLAYBOOK.md` lines 139-164; reusing one tier id can resolve to the wrong earlier relay exactly at a gate.

Required correction: add a compact header/edge table covering at least:

1. master -> l4 inert kickoff/preflight assignment;
2. l4 decomposition -> l4 reviewer verdict;
3. reviewer-approved l4 worker dispatch -> worker return;
4. l4 integration request -> l4 final content verdict;
5. escalation -> master/owner disposition -> return down.

Use unique sub-dispatch ids per gated leg under a stable hierarchical prefix. Keep `PARENT_DISPATCH_ID` on the immediate predecessor; encode tier ancestry in the id namespace/run convention, not by replacing the predecessor edge. The preflight must export and validate one real three-tier chain of this exact shape.

### LANE4-VP-R6-F3 -- GATE: the assigned decomposition review has no predecessor position in the sequence

Plan line 21 assigns `l4.reviewer` adversarial review of the decomposition. Sequence line 80 instead has `l4.planner` decompose **and dispatch workers immediately** after activation. No durable reviewer verdict gates that dispatch.

A full orchestrator-team whose reviewer only reviews the final oracle is not exercising the claimed planner/reviewer topology at the decision that creates worker scope, overlap, and authority. It also leaves no approving immediate predecessor for the worker dispatches required by F2.

Required correction: after activation:

1. `l4.orchestrator-planner` writes the decomposition, worker topology, per-worker fence, artifact ownership, cross-record budget allocation, carried-obligation allocation, and escalation rules;
2. `l4.orchestrator-reviewer` files a durable approve/revise verdict;
3. only an approve permits addressed worker dispatches; revise returns to decomposition; and
4. the later byte-equality and final content-review duties remain separate and unchanged.

The worker topology and decomposition grain can remain the local planner's choice, but the choice must pass the local reviewer before any worker authors.

### LANE4-VP-R6-F4 -- GATE: the plan conflates today's seat boot/host spawn with the deferred native governed-spawn mechanism

Frank can act as courier for operator-minted seats today, but `frank/ROADMAP.md` lines 180-190 explicitly defers the permission/authority system and **native governed agent-spawn** to Step 4+. Rev6 nevertheless says workers are "spawned," says it pins authority-ceiling-at-spawn, and says the preflight boots a probe under **zero dispatch authority**, without naming who creates the worker or which current mechanism carries the ceiling.

The ambiguity creates two incompatible readings:

- the operator mints/boots a preconfigured probe, which can test courier routing and lineage but does **not** test l4-controlled spawn; or
- `l4.planner` spawns/dispatches the probe, which contradicts the preflight's zero-dispatch ceiling and risks claiming a native frank capability that does not exist yet.

Required correction:

1. name the current-generation actor and mechanism separately for seat mint, host-session/subagent creation, boot, and frank credential wiring;
2. state explicitly that frank is the courier and seat-identity carrier here, **not** the Step-4 native governed-spawn engine;
3. declare whether B21 is a lane-specific operator exception to the charter's non-reviewer-subagent restriction or whether workers are independently booted sessions rather than subagents;
4. state the ceiling's current evidence grade honestly: convention/config/read-only-tool proof plus reviewer/master checks, not m-5 mechanical enforcement;
5. reconcile "zero dispatch authority" with the probe: either the operator directly creates/boots it, or define a narrow transport-only probe action that cannot authorize fixture work; and
6. make the preflight exercise both orchestrator seats plus at least one worker through accepted send/read/parent/export evidence.

Do not claim this run exercises or hardens the future m-5 native ceiling carrier unless an actual current mechanism and evidence prove that claim.

### LANE4-VP-R6-F5 -- RECORD CORRECTION: the void kickoff banner binds the wrong rev6 hash

`master/STEP-3-LANE4-KICKOFF.md` is safely marked VOID, but its line 3 says the future nested kickoff follows approval of rev6 `3e71894c...`. The routed rev6 is `959a29aaccf3f23910cf237746acaabde6d94457e68db3937c185c1c7b329ff9`.

Required correction: replace the stale hash with the routed rev6 hash or state "pending VP-approved successor" without a hash. Make line 7 explicitly historical authority-at-draft-time so the void file cannot be read as saying pair-shaped rev5 is still current authority. Correct the incoming conformance claim on the return relay.

## Passed scope

- **B21 team-shape decision ACCEPTED:** the operator's full nested-team-on-frank choice supersedes rev5's pair staffing for lane 4. This review does not reactivate the pair.
- The exact ten-record/six-leg schema, fixed values, carried obligations, owner-real matrix, and guiding m-3 PM remain correct.
- The proposal-envelope, master-only materialization, frame-fit/HOLD, reviewer equality + content duties, owner-fidelity-first ordering, and Master+VP freeze/re-lock remain correct.
- Inert kickoff -> operator preflight authorization -> operator activation remains the right outer order once F1-F4 make the nested inner mechanics exact.
- Item A remains byte-stable; H-16/H-26 still precede T4, and H-12 still blocks external use.
- The old pair-shaped kickoff is visibly VOID and has not activated work; only its stale banner binding needs correction.

## Gate disposition

- Return one bounded rev7 correcting F1-F5 in the plan, B21 record, and void-banner record. Preserve every passed rev1-rev6 fixture/freeze decision.
- No fresh kickoff, seat mint/boot, probe, preflight, activation, decomposition, worker dispatch, proposal, materialization, fixture, manifest, owner-fidelity, freeze, re-lock, or T4 action on rev6.
- The next review must carry exact hashes for the rev7 plan, updated B21 record, corrected void kickoff, and the return relay.
- Approval, when earned, remains approach/design approval only: it permits master to write the nested **inert** kickoff; the operator still owns boot and activation.

## Verification

- Recomputed SHA-256: incoming `2447ab6e140a988f6819236efc16e0c6fc4f4c9d86f9a8cee3abe20631237208`; plan rev6 `959a29aaccf3f23910cf237746acaabde6d94457e68db3937c185c1c7b329ff9`; protocol deviations `2965df2197889d71babce7b31f31a751030325f2ed30fd23f668d9e4cf6cc8c5`; void kickoff `f697d859b44252d5d9f8eccab5b1343b5203865c38b6e77aafd359dccef21b57`; interface lock `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.
- Exact-file lint is `OK` for the incoming relay.
- Protocol ROLE/FROM mapping and prior S1 nested-team headers prove the canonical orchestrator/core-seat address shapes.
- Protocol `PARENT_DISPATCH_ID` semantics and the recorded shared-id resolver defect prove the current tier-parent convention is insufficient.
- Frank's live `seat_mint` path proves operator-controlled seat creation exists; the roadmap separately proves native governed spawn/authority enforcement is deferred.
- Fresh interface-lock manifest rehash: `rows=38 distinct=38 mismatches=0`.
- `git -C frank status --short --branch` is `## main...origin/main`, with empty porcelain and HEAD/origin both `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row. No plan, B21 record, kickoff, seat mint/boot, preflight, activation, decomposition, worker dispatch, proposal, fixture, manifest, lock, owner/frozen artifact, `frank/` source, branch, commit, PLAN/T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260728-090000.md`.
Next requested action: issue bounded plan rev7 with an exact canonical seat map, immediate-predecessor edge table, pre-dispatch l4 reviewer gate, honest current-generation spawn/preflight carrier, and corrected void-banner hash; return all exact hashes for VP re-review.
