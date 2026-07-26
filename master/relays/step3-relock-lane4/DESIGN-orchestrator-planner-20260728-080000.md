## DESIGN — lane-4 plan rev6 `959a29aa…`: the operator's B21 directive — the lane-4 team is a FULL NESTED orchestrator-team on frank (l4.planner + l4.reviewer + spawned l4.w<k> sub-workers), NOT the B13 pair; driver = dogfooding the nested model + frank-as-courier. Reopens the team-structure decision VP-approved rev5 (`faa23c7f…`) settled as a pair; pins B13's deferred nested-lineage + authority-ceiling-at-spawn debt (monotone-non-increasing ceiling; 3-tier lineage). Every non-team-structure rev1–rev5 decision preserved. A conformance pass ran before this routing (5 findings folded). The pair-shaped kickoff draft is VOID.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-lane4
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — B21 (the full-nested-team-on-frank staffing) is an operator directive already given; and the operator authorizes the zero-authority preflight boot + the separate post-preflight activation before any authoring. A hand-relay fallback (preflight fail) is an operator-owned deviation.
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-lane4-staffing-grill-1
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260728-070000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review lane-4 plan rev6 `959a29aaccf3f23910cf237746acaabde6d94457e68db3937c185c1c7b329ff9` — operator-directed full nested orchestrator-team on frank (B21) replaces the pair; nested ceiling/lineage pinned; conformance-checked; Item A lock `cbd1893c…` preserved

## What changed vs the VP-approved rev5 `faa23c7f…` (a team-structure reopen, operator-directed)
The operator directed (2026-07-28): *"if throughput we're not doing it thru frank; if not we're going on the full team + frank"* → master chose **full team + frank** (driver = the dogfood, not authoring speed; the tightly-coupled single-manifest deliverable would actually favor a pair on throughput grounds — the full team is for learning + to pin the deferred nesting debt on a safe read-only task). Seats named **l4.\*** per the operator.

- **§3 team.** Pair → **full nested orchestrator-team:** `l4.planner` (orchestrator — decomposes the ten records, dispatches sub-workers, integrates the manifest, owns the 30-turn/100-call budget) + `l4.reviewer` (adversarial + the two read-only verification duties: byte-equality + content review) + **spawned `l4.w<k>` sub-workers** (author fixtures; single-seat-or-pair + decomposition grain = `l4.planner`'s local call).
- **Nested debt pinned (B13's deferred item):** authority **monotone non-increasing** down the tree (a sub-worker never exceeds `l4.planner`'s read-only/author-only/no-governed-tree-write ceiling); **three-tier lineage** `step3-relock-lane4` → `…-l4` → `…-l4-w<k>` via `PARENT_DISPATCH_ID`; three-tier escalation. Interim pin; m-5's archetype ceiling is the Step-4 carrier.
- **§7 preflight/sequence:** the preflight boots the orchestrator + ≥1 probe sub-worker and runs a real **three-tier** round-trip (sub-worker → l4 → master → back); on activation `l4.planner` decomposes/dispatches, the team authors, `l4.reviewer` verifies, master materializes.
- **Unchanged (preserved from rev1–rev5):** the exact §7 schema + fixed values; the read-only + proposal-envelope + master-materializes + frame-fit chain; the two verification duties (now `l4.reviewer`'s); the owner-real §5 matrix; guiding PM m-3; **Master+VP-only freeze + re-lock**; H-16/H-26 before T4; H-12. **`GRILL_LOCK step3-lane4-staffing-grill-1` updated** (team-shape, nested ceiling/lineage, three-tier preflight, source through rev6).

## Conformance pass (the discipline, before this routing)
An independent conformance check ran over rev6 vs §7 / B13-B21 / Part F / the lock + internal consistency. **§7 schema, lock treatment, and the core nested-team mechanics passed clean.** Five findings folded: the pre-existing pair-shaped kickoff draft **VOIDed** (`master/STEP-3-LANE4-KICKOFF.md` now banner-superseded, retained as history only); residual "out-of-pair" wording → "out-of-team" (§3/§5/§7); §7-step-5 fidelity-then-VP ordering made sequential; sub-worker notation normalized (`l4.w<k>`); a §7-step-0 note that the void draft is not reused.

## What I ask the VP to review
- Is the **nested-team structure** (§3) + the **monotone ceiling + three-tier lineage** a sound interim pin of B13's deferred debt for a read-only design-only task?
- Did the reopen **preserve** every non-team-structure rev5 decision (schema, chain, matrix, Master+VP freeze, H-gates)?
- Does the **three-tier preflight/sequence** (§7) hold, and is the delivery chain coherent for the nested team?
- Anything reopened that shouldn't be, or a locked byte touched?

## Boundaries
No fixtures, manifest, freeze, re-lock, PLAN-lock, T4/code token, credential, provider call, `frank/` action, preflight boot, or team activation on this revision. Item A lock `cbd1893c…` + all owner/frozen/amendment bytes UNMOVED. On VP approval → master writes the **nested** INERT kickoff brief (the void pair-shaped draft is replaced, not reused); the operator authorizes the zero-authority preflight boot; on pass + activation the team authors. **H-12 hard-blocks external use.**

## Verification
Plan rev6 `master/STEP-3-LANE4-PLAN.md` = SHA-256 `959a29aaccf3f23910cf237746acaabde6d94457e68db3937c185c1c7b329ff9` (supersedes the pair-shaped rev5 `faa23c7f…` on team structure). **B21** appended to `master/PROTOCOL-DEVIATIONS.md`. No residual operative "pair"/`.implementer`/out-of-pair reference (grep 0). Kickoff draft VOIDed. Item-A lock `cbd1893c…` PRESERVED (fresh rehash 38 distinct, 0 mismatch expected). Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — rewrote `master/STEP-3-LANE4-PLAN.md` to rev6 (nested team + ceiling/lineage + 3-tier preflight/sequence + GRILL_LOCK) + appended B21 to `master/PROTOCOL-DEVIATIONS.md` + VOIDed `master/STEP-3-LANE4-KICKOFF.md` + this relay + one INDEX.md row; no fixtures/manifest/lock/owner/frozen/amendment byte moved, no `frank/` action, no preflight boot, no team activation, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP re-reviews plan rev6 `959a29aa…` + the updated GRILL_LOCK; on approve → master writes the nested INERT kickoff brief; the operator authorizes the zero-authority preflight boot; on preflight-pass + operator activation the l4 team authors. Freeze + re-lock stay Master+VP; T4 behind the re-lock + H-16/H-26. H-12 stands.
