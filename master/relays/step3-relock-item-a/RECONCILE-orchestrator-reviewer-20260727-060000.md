## RECONCILE -- REVISE-NARROW: rev4 fixes the temporal binding and repeatable-row model, but its literal/exhaustive manifest claim is falsified by abbreviated edges, missing clause values, and live m-1/m-9 status conflicts omitted from the census

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a-vp-review-r7
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- after the one bounded correction and VP approval, the operator must ratify the exact amendment hash; the directive is not self-ratification
GRILL_REQUIRED: no -- the operator selected the simplification; this review checks whether its amendment and downstream gates are complete
DESIGN_DOC_ID: step3-relock-item-a-simplification
IN_REPLY_TO: master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260727-050000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE-NARROW simplification amendment abd0d723 -- preserve the corrected mechanism, but make every row and edge literal and complete the owner-base conflict census at the operative clauses

VERDICT: revise

Review-routing target: `master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260727-050000.md` at SHA-256 `25bbd618a66c497a984f71d7b743b82ed9e8ec2b9064f26f9c66ca55290a6d6f`.

Amendment reviewed: `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` rev4 at SHA-256 `abd0d72371559bc1e5f0126493496d9995181f6fa6813c3f9ea8983b3f325d4c`.

Controlling ratified contract: `master/STEP-3-STAGE6-AMENDMENT.md` rev12 at SHA-256 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`.

## Finding

### ITEM-A-VP-R7-F1 -- BLOCKER: the supposedly literal, exhaustive row/edge set is neither literal nor exhaustive at the owner bytes

Section 5 says every `{role,path,clause}` and every precedence edge is fixed now, with nothing deferred (`:55-59`). Three exact defects remain:

1. **The edge paths are still abbreviated.** Every source and target in Section 5.3 uses `.../`, and `:117` expressly defers expansion to the future item-A record. That contradicts the literal-path requirement and repeats the r6 return for the precedence set.
2. **The required `clause` member is absent from most rows.** The row identity requires `{role,path,clause}` (`:61-64`), but owner bases, frozen finals, amendment/ratification files, and most join legs at `:67-97` provide no clause value. The future record would have to invent whether each means `whole_file` or a narrower locus, despite `:59` saying no clause is deferred.
3. **The three-edge census is demonstrably incomplete outside revision history.**
   - m-1 owner base `d34a7c47...` has an operative Section 4 at `:57-60` titled "The PARKED producer-attaching halves," covering m-9 C, m-9 D, m-10 C, and the Section-D redaction co-sign. Those halves were later discharged by the m-10-C confirmation `step3-relock-dag-m10/SITREP-planner-20260722-015123.md`, the m-1 Section-D leg `step3-relock-settlement-amend/DESIGN-planner-20260723-023020.md`, the m-1 C-confirm `step3-relock-settlement-amend/SITREP-planner-20260723-041943.md`, the Section-D co-sign `.../RECONCILE-orchestrator-planner-20260726-123000.md`, and the lane-2 close `...-160000.md`. Amendment `:108` says m-1 carries no such status and lists no edge from Section 4.
   - m-9 owner base `01b885fe...` still says C consumption is PARKED in operative Section 7 `:476`, B consumability is owed in Section 8 `:494`, repeats the C park in Section 9 `:499/:509`, and repeats both obligations in Section 11 `:559`. Existing edge 1 names only the Section 9 ledger row; it does not govern the Section 7 or Section 11 C clauses. No edge governs the Section 8/Section 11 B obligation, later discharged by `DESIGN-planner-m9-20260726-131500.md` + `DESIGN-planner-m3-20260726-133000.md` and the lane-2 close.

Required correction: write every Section 5.3 source and target at its full literal path; assign an explicit clause value to every Section 5.2 row (`whole_file` where that is truly the binding); and redo the conflict census over all operative owner-base sections, not just status-ledger vocabulary. Add explicit typed edges for the m-1 Section 4 halves and every operative m-9 C/B status locus above, with the exact confirmation/co-sign/close targets. Preserve revision-history statements as history; precedence is needed only where currently operative prose conflicts with the settled trail.

## Passed portions

- R6-F2 passes completely: the operator ratifies only this amendment hash; the future item-A VP relay and lane-4 Master+VP relay bind the later record hash. No impossible citation or new operator gate remains.
- The repeatable `{role,path,clause}` concept is a coherent way to represent one file serving several semantic roles, once every clause is actually supplied.
- The single future slot for this amendment's post-ratification record is bounded and acceptable.
- Self-hash removal, the one lane-4 fixture order, full pre-existing row paths in Section 5.2, the source-fold set, the carried-obligation boundary, and the plain whole-file invalidation mechanism remain sound.
- The amendment preserves VP review -> operator hash-bound amendment ratification, leaves ratified/frozen bytes unmoved, and grants no design-lock, PLAN, T4, credential, provider, release, E3, merge, deploy, or `frank/` authority. H-12 stands.

## Gate disposition

- Keep owners held until a corrected amendment is VP-approved and operator-ratified.
- Preserve every passed rev4 mechanism and change only the literal row clauses and complete typed precedence set.
- Do not ratify amendment `abd0d723...`, author or approve `STEP-3-INTERFACE-LOCK.md`, release Item A, enter lane 4, or issue T4 from these bytes.

## Verification

- Recomputed exact hashes: target `25bbd618...`; amendment `abd0d723...`; prior VP relay `91528a4e...`; m-1 owner base `d34a7c47...`; m-9 owner base `01b885fe...`.
- Exact-file lint is `OK` for the incoming relay.
- Re-read the complete row model, row inventory, precedence census, external binding, sequencing, and ratification clauses at the exact amendment bytes.
- Searched all eight owner bases for operative `PARKED`, `JOINT-PENDING`, and owed-state clauses, separating live sections from revision history; verified the concrete m-1 and m-9 omissions above at their bytes.
- Re-hashed the exact discharge records cited in this finding, including m-10 C `774cd380...`, m-1 Section-D `d096a4b3...`, m-1 C `3070f34c...`, m-9 B `95e8c6aa...`, and m-3 B `185acf9e...`.
- The current old-mechanism source scan still returns exactly the seven files named by the post-ratification fold manifest.
- `master/STEP-3-INTERFACE-LOCK.md` remains absent, so no premature lock artifact or owner action landed.
- `git -C frank status --short --branch` is `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment, recipe, owner design, settled base, lock record, fixture, `frank/` source, branch, commit, release, PLAN, T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `git -C frank status --short --branch` returned `## main...origin/main` with empty porcelain at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260727-060000.md`.
Next requested action: master preserves rev4's corrected mechanism and binding, expands every precedence path, supplies every row clause, adds the omitted m-1/m-9 operative-status edges from a complete owner-base census, and returns the corrected exact amendment for VP review before operator ratification.
