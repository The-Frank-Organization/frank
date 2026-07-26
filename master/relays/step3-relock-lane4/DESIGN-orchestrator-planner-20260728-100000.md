## DESIGN — lane-4 plan rev7 `e7a333e9…`: folds VP r6's five nested-mechanics gates (B21 full-team ACCEPTED). F1 canonical role-stamped seats (l4.orchestrator-planner/l4.orchestrator-reviewer + l4.w<k>.planner/.implementer, `.implementer`→`ROLE: Implementer`). F2 PARENT_DISPATCH_ID = immediate-predecessor edge + tier ancestry in the dispatch-id namespace + a gated-leg edge map + unique per-leg ids (shared-id resolver defect). F3 a decomposition-review gate before worker dispatch. F4 the HONEST current-generation spawn model (operator mints/boots seats; frank = courier + seat-identity, NOT the Step-4 native-spawn engine; ceiling = convention/read-only/review grade, not m-5-mechanical; dogfood INFORMS not exercises m-5). F5 void-banner stale-hash fixed. A conformance pass ran before this routing (caught + fixed a lint-breaking `.implementer`/ROLE mismatch).

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-lane4
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — B21 (full nested team on frank) is an operator directive; and the operator mints + boots the seats + authorizes the zero-authority preflight + supplies the post-preflight activation before any authoring. A hand-relay fallback (preflight fail) is an operator-owned deviation.
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-lane4-staffing-grill-1
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260728-090000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review lane-4 plan rev7 `e7a333e9c4c5e34cb62dffa29c0b37f03d48022a233636a0d0c34b28006994d2` — the five nested-mechanics gates folded (canonical seats, immediate-predecessor lineage + edge map, decomposition gate, honest current-gen spawn model, void-banner hash); Item A lock `cbd1893c…` preserved

## What changed vs rev6 `959a29aa…` (the five r6 gates; B21 full-team accepted)
- **F1 — canonical role-stamped addresses.** `l4.orchestrator-planner` (`ROLE: Orchestrator Planner`) + `l4.orchestrator-reviewer` (`ROLE: Orchestrator Reviewer`) + workers `l4.w<k>.planner` (`ROLE: Planner`) / `l4.w<k>.implementer` (**`ROLE: Implementer`** — the relay-lint tripwire). The `l4.*` namespace is preserved; the suffix is protocol-canonical so frank's registry grants orchestrator authority and durable relays pass ROLE/FROM.
- **F2 — lineage.** `PARENT_DISPATCH_ID` = the **immediate-predecessor edge** (not a static tier-parent); tier ancestry lives in the hierarchical DISPATCH_ID namespace (`step3-relock-lane4` → `…-l4` → `…-l4-w<k>`) with a **unique sub-dispatch id per gated leg** — avoids the shared-dispatch-id resolver defect (`CYCLE-PLAYBOOK.md:139-164`). §3 carries the 5-row gated-leg edge map (kickoff/preflight · decomposition→verdict · approved-dispatch→return · integration→final-verdict · escalation→disposition).
- **F3 — decomposition-review gate.** §7 step 2: `l4.orchestrator-planner` writes the decomposition (topology/fence/artifact-ownership/budget+carried allocation/escalation) → `l4.orchestrator-reviewer` files a durable approve/revise verdict → **only an approve permits addressed worker dispatch** (revise returns). The byte-equality + content-review duties stay distinct/later.
- **F4 — honest current-generation spawn model.** Native governed-spawn + the permission system are **Step-4-deferred** → the **operator mints (`seat_mint`) + boots each seat as an independent session** (not l4-spawned, not a subagent); `l4.orchestrator-planner` **dispatches** by frank relay, does not create seats; **frank = courier + seat-identity carrier, NOT the native-spawn engine**; the ceiling is **convention/config/read-only-tool + review grade, honestly labeled — NOT m-5 mechanical enforcement**; this dogfood **INFORMS** (battle report) the future m-5 carrier, does not exercise/harden it. Preflight: the operator boots a **zero-authority probe** (not an l4 spawn) exercising both orchestrator seats + ≥1 worker with correct immediate-predecessor chaining + durable export.
- **F5 — void banner.** `master/STEP-3-LANE4-KICKOFF.md` no longer binds the stale rev6 hash; it says "pending VP-approved successor (rev7+)" and marks its authority-of-record line historical-at-draft-time.
- **B21 updated** to the canonical seats + honest carrier. **Unchanged/preserved:** the exact §7 schema + fixed values, the owner-real §5 matrix, the proposal→materialize→byte-equality→content-review→owner-fidelity→VP→freeze chain, Master+VP-only freeze/re-lock, H-16/H-26 before T4, H-12.

## Conformance pass (the discipline, before routing)
An independent conformance check ran over rev7 vs the five corrections + §7 no-regress. **One defect caught + fixed:** a rev7-introduced `.implementer`→`ROLE: Reviewer` mismatch (would fail relay-lint) → corrected to `ROLE: Implementer`. All other corrections + the §7 schema + the delivery chain + the lock treatment verified clean.

## What I ask the VP to review
- Are all seat addresses role-canonical (incl. the corrected `.implementer`)?
- Is the immediate-predecessor lineage + edge map + unique-per-leg-id correct, and the decomposition gate sound?
- Is the honest current-gen spawn model accurate (operator mints/boots; frank not the spawn engine; convention-grade ceiling; informs-not-hardens m-5)?
- Anything reopened that shouldn't be, or a locked byte touched?

## Boundaries
No fixtures, manifest, freeze, re-lock, PLAN-lock, T4/code token, credential, provider call, `frank/` action, seat mint/boot, preflight, or team activation on this revision. Item A lock `cbd1893c…` + all owner/frozen/amendment bytes UNMOVED. On VP approval → master writes the nested **INERT** kickoff brief; the operator mints/boots the seats + authorizes the zero-authority preflight; on pass + activation the team authors. **H-12 hard-blocks external use.**

## Verification
Plan rev7 `master/STEP-3-LANE4-PLAN.md` = SHA-256 `e7a333e9c4c5e34cb62dffa29c0b37f03d48022a233636a0d0c34b28006994d2` (supersedes rev6 `959a29aa…`). No residual bare-seat/`.implementer`→Reviewer/one-tier-parent/l4-spawn/exercise-harden language (grep 0). **B21** updated in `master/PROTOCOL-DEVIATIONS.md`; kickoff banner de-hashed. Item-A lock `cbd1893c…` PRESERVED. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — rewrote `master/STEP-3-LANE4-PLAN.md` to rev7 (canonical seats + lineage/edge-map + decomposition gate + honest spawn/preflight + GRILL_LOCK) + updated B21 in `master/PROTOCOL-DEVIATIONS.md` + de-hashed the void `master/STEP-3-LANE4-KICKOFF.md` banner + this relay + one INDEX.md row; no fixtures/manifest/lock/owner/frozen/amendment byte moved, no `frank/` action, no seat mint/boot, no activation, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP re-reviews plan rev7 `e7a333e9…` + the updated GRILL_LOCK; on approve → master writes the nested INERT kickoff brief; the operator mints/boots the seats + authorizes the zero-authority preflight; on pass + operator activation the l4 team authors. Freeze + re-lock stay Master+VP; T4 behind the re-lock + H-16/H-26. H-12 stands.
