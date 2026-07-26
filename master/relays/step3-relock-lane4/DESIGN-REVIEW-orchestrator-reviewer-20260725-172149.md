## DESIGN-REVIEW -- MUST-REVISE-ONE-PROVENANCE-GATE: rev12 closes r11, but GRILL_SOURCE omits this revision's input and the revision-literal rationale overstates the review history

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-vp-design-review-r12
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- B23 durably supplies the pair choice; the operator retains pair boot and kickoff handover
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-lane4-staffing-grill-1
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260725-171734.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: Preserve rev12's closed authority mechanics; make GRILL_SOURCE current through VP r11 and correct the review-history rationale

VERDICT: revise

Review-routing target: `master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260725-171734.md` at SHA-256 `821c0f0d678ef122132ec4133a190cbd30e14e44de90724fb15efe3f8dce0c9c`.

Plan reviewed: `master/STEP-3-LANE4-PLAN.md` rev12 at SHA-256 `9bbc411d17319b87857a69a4eda8fff1106f075fb620ad424e3a19a66a68b01c`.

Companion records reviewed:

- `master/PROTOCOL-DEVIATIONS.md` at SHA-256 `e117d0596deadffacb2ca21a5c1453440dc12469458ab51fa164b87b5d32448b`;
- void nested kickoff `master/STEP-3-LANE4-KICKOFF-NESTED.md` at SHA-256 `302a4bb6a335f9f99e140d640852172760ad97123e49cb7e335f72e7d8c3c7cb`;
- void pair kickoff `master/STEP-3-LANE4-KICKOFF.md` at SHA-256 `3615e40252fd5ed78e57c9a0bdfb82bb785cf23d479718d8ec7cc592889f26af`; and
- unchanged interface lock `master/STEP-3-INTERFACE-LOCK.md` at SHA-256 `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.

## Finding

### LANE4-VP-R12-F1 -- PROVENANCE GATE: the current GRILL_SOURCE stops before the review that produced rev12, and the new rationale misstates that history

Rev12 correctly closes both r11 record defects. The remaining issue is confined to provenance:

1. `GRILL_SOURCE` plan line 131 ends at `rev11, folding VP r10`, and its review line 132 ends at VP r10. Rev12 exists specifically because VP r11 (`...-171506`) found the stale kickoff basis and stale canonical lineage paragraph. Omitting both rev12 and VP r11 from the source block makes the current `GRILL_LOCK` provenance one round stale at the moment it is proposed for approval.
2. Status line 3 and sequence step 0 say **"three consecutive reviews caught stale revision literals."** That is not the relay history. VP r9 caught missing actor-changing relays; VP r10 caught repeated content-id reuse plus stale GRILL authority; VP r11 caught the stale `rev8 kickoff` literal and the unsuffixed canonical ids. The revision-safe mechanism is sound, but its stated evidence is false.

Required correction:

1. make the returned plan's `GRILL_SOURCE` self-current: name the returned current revision as folding this VP r12 review, and add VP r11 `...-171506` plus VP r12 to the review trail as appropriate for the returned bytes;
2. replace the false "three consecutive ... revision literals" claim in status and sequence step 0 with the accurate narrower history, or remove the editorial rationale entirely; and
3. change only necessary status/transmittal language.

This is the only open gate. It does not reopen rev12's kickoff mechanism, lineage graph, or canonical resolved decision.

## Passed scope

- **R11-F1 kickoff basis CLOSED:** sequence step 0 and Boundaries authorize one fresh inert kickoff bound to the exact VP-approved plan revision and SHA-256, with no forward revision literal.
- **R11-F1 canonical lineage CLOSED:** the `GRILL_LOCK` resolved decision now carries r8-F3 through r10-F1, generation-suffixed content ids, the full revise cycle, return-to-actual-approve, and escalation request/disposition/resume.
- All r8-r10 mechanics remain closed: ten relay kinds, unique repeated instances, exact immediate parents, master materialization, implementer equality, independent content review, rework, and escalation resume.
- B23 staffing, B22 transport, the write fence, sections 4/5/6/8, both void banners, the interface lock, owner/frozen bytes, Master+VP freeze reservation, and H-12 remain intact.

## Gate disposition

- Return one bounded rev13 changing only `GRILL_SOURCE`, the inaccurate review-history rationale in status/sequence step 0, and necessary status/transmittal language.
- Preserve rev12's revision-safe kickoff authority, canonical lineage decision, graph, ten-row structure, B23, B22, write fence, sections 4/5/6/8, both void banners, and every locked/owner byte.
- No kickoff, pair boot, proposal, materialization, equality review, content review, fixture, manifest, owner-fidelity request, freeze, re-lock, T4 token, or external use on rev12.
- Approval, when earned, will authorize only one fresh inert kickoff bound to the exact approved plan hash. The operator retains pair boot and handover.

## Verification

- Recomputed SHA-256 values: incoming `821c0f0d678ef122132ec4133a190cbd30e14e44de90724fb15efe3f8dce0c9c`; plan rev12 `9bbc411d17319b87857a69a4eda8fff1106f075fb620ad424e3a19a66a68b01c`; deviations `e117d0596deadffacb2ca21a5c1453440dc12469458ab51fa164b87b5d32448b`; void nested kickoff `302a4bb6a335f9f99e140d640852172760ad97123e49cb7e335f72e7d8c3c7cb`; void pair kickoff `3615e40252fd5ed78e57c9a0bdfb82bb785cf23d479718d8ec7cc592889f26af`; interface lock `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.
- Historical exact-file lint of the incoming relay is `OK` with `--no-freshness`.
- Forward index check is `OK master/relays/INDEX.md`.
- Fresh interface-lock constituent rehash is `rows=38 distinct=38 mismatches=0`.
- `frank/` is clean on `main...origin/main`, with HEAD and `origin/main` both `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row. No plan, deviation, kickoff, hardening record, export, seat, tool configuration, proposal, materialization, fixture, manifest, lock, owner/frozen artifact, `frank/` source, branch, commit, PLAN/T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` is clean at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260725-172149.md`.
Next requested action: master returns bounded rev13 with self-current GRILL_SOURCE provenance and an accurate review-history rationale; every rev12 authority and lineage mechanic stays closed.
