## DESIGN — lane-4 plan rev5 `faa23c7f…`: folds VP r4's three gates (the r3 mechanics stayed CLOSED). F1 restores the DISTINCT independent `.implementer` CONTENT review (a rev4 regression — I'd collapsed the implementer role to byte-equality only): the .planner authors content, the .implementer does (i) byte-equality AND (ii) an adversarial content review of the full materialized set + manifest with a durable approve/revise verdict BEFORE owner-fidelity + VP freeze. F2 advances the embedded GRILL_LOCK source trail through rev5/r4 + refreshes its design-lock impact. F3 (transmittal) carries the ratified fixed values.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-lane4
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the operator authorizes the zero-authority preflight boot AND separately supplies the post-preflight activation/green-light before any lane-4 authoring; a hand-relay fallback (preflight fail) is an operator-owned B13 deviation.
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-lane4-staffing-grill-1
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260728-050000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: VP re-review lane-4 plan rev5 `faa23c7f6486a19e6960d2bb5e3a72a3f1e9b08cf4c33dee22ce1ad171e3679b` — restored independent implementer content review; refreshed GRILL_LOCK record; transmittal carries the fixed typed values; Item A lock `cbd1893c…` preserved

## What changed vs rev4 `8f47ed90…` (the three r4 gates; r3 mechanics stayed closed)
- **F1 — independent implementer content review restored (my rev4 regression).** When I added the manifest-envelope chain in rev4 I collapsed the `.implementer` to byte-equality only, dropping the content review rev3 carried. Corrected in §3 + §7: the **`.planner` is the proposal/content author**; the **`.implementer` has two distinct read-only roles — (i) proposal-to-file byte-equality (every file incl. the manifest + any chunk reassembly), and (ii) a distinct independent adversarial CONTENT review** of the full materialized fixture set + final manifest (scenario discrimination, expected canonical rows, typed expectations + fixed values, carried rows, observer/locator resolvability, weight allocation, chunk contract) with a **durable approve/revise verdict**. **Owner-fidelity + VP review/freeze occur only after the implementer content-approves.** Byte-equality ≠ content review.
- **F2 — GRILL_LOCK record advanced.** `GRILL_SOURCE` now runs through rev5 (`…rev4 8f47ed90 → rev5`) and reviews r1–r4; `Design-lock impact` refreshed to the current locked decisions (complete-manifest exact-byte chain + frame-fit/HOLD; the two implementer roles incl. the content review; the owner-real matrix; the inert-kickoff order).
- **F3 — transmittal fixed values.** The ratified §7 typed values, stated exactly (no tuple shorthand, no key-count, no per-record-total): `effect_counter_expectation { counter_before_recovery: 1, counter_after_recovery: 1, invocations_after_recovery: 0 }` bound to `xit-crash-1`; `degraded_expectation { corruption_cut, expected_disposition: "degraded", expected_resume_action }` bound to `xit-dur-2`.

## What I ask the VP to review (approach only)
- Is the distinct implementer content review restored (approve/revise verdict, before owner-fidelity/VP), and the planner-author/implementer-checker split explicit?
- Is the embedded GRILL_LOCK source + design-lock impact now current through rev5/r4?
- Does anything reopen a closed r1–r3 decision or a locked byte?

## Boundaries
No fixtures, manifest, freeze, re-lock, PLAN-lock, T4/code token, credential, provider call, `frank/` action, preflight boot, or team activation on this revision. Item A lock `cbd1893c…` + all owner/frozen/amendment bytes UNMOVED. On VP approval → master writes the INERT kickoff; the operator authorizes the zero-authority preflight boot; on pass + activation the pair authors. **H-12 hard-blocks external use.**

## Verification
Plan rev5 `master/STEP-3-LANE4-PLAN.md` = SHA-256 `faa23c7f6486a19e6960d2bb5e3a72a3f1e9b08cf4c33dee22ce1ad171e3679b` (supersedes rev4 `8f47ed90…`). The distinct implementer content review (§3 role ii + §7 step 4), the refreshed GRILL_SOURCE/Design-lock impact, and the exact fixed values verified present. Item-A lock `cbd1893c…` PRESERVED. Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — edited `master/STEP-3-LANE4-PLAN.md` to rev5 (§3 implementer two-role · §7 content-review step · §10 GRILL_LOCK source + design-lock impact · header) + this relay + one INDEX.md row; no fixtures/manifest/lock/owner/frozen/amendment byte moved, no `frank/` action, no preflight boot, no team activation, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP re-reviews plan rev5 `faa23c7f…` + GRILL_LOCK; on approve → master writes the INERT detailed kickoff brief; the operator authorizes the zero-authority preflight boot; on preflight-pass + operator activation the lane-4 pair authors. Freeze + re-lock stay Master+VP; T4 behind the re-lock + H-16/H-26. H-12 stands.
