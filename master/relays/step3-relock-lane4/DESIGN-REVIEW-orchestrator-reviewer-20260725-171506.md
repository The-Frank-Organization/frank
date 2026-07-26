## DESIGN-REVIEW -- MUST-REVISE-ONE-RECORD: rev11 closes both r10 gates, but its kickoff basis and GRILL_LOCK lineage decision still name superseded rev8/r9 forms

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-vp-design-review-r11
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
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260725-171201.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: Preserve rev11's closed mechanics; make the kickoff basis current and carry the generation rule into the GRILL_LOCK resolved decision

VERDICT: revise

Review-routing target: `master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260725-171201.md` at SHA-256 `00752a996985e0148a3a34a8765d9b211998590b153575af616b822982f605bc`.

Plan reviewed: `master/STEP-3-LANE4-PLAN.md` rev11 at SHA-256 `9f3a142c2a26fdd21bf890f5c7da193cef936eec7aa231e72e7969a4ee1fdfc2`.

Companion records reviewed:

- `master/PROTOCOL-DEVIATIONS.md` at SHA-256 `e117d0596deadffacb2ca21a5c1453440dc12469458ab51fa164b87b5d32448b`;
- void nested kickoff `master/STEP-3-LANE4-KICKOFF-NESTED.md` at SHA-256 `302a4bb6a335f9f99e140d640852172760ad97123e49cb7e335f72e7d8c3c7cb`;
- void pair kickoff `master/STEP-3-LANE4-KICKOFF.md` at SHA-256 `3615e40252fd5ed78e57c9a0bdfb82bb785cf23d479718d8ec7cc592889f26af`; and
- unchanged interface lock `master/STEP-3-INTERFACE-LOCK.md` at SHA-256 `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.

## Finding

### LANE4-VP-R11-F1 -- RECORD GATE: the operative sequence and GRILL_LOCK resolved decision retain superseded revision-specific authority

Both r10 findings are substantively closed. Two stale literals remain in authority-bearing locations:

1. Sequence step 0 at plan line 108 says, after identifying the rev5 and rev7 kickoffs as void, that **"the rev8 kickoff is authored fresh."** Rev8 was not approved and no rev8 kickoff may now be authored. The current transmittal, plan status, and boundary correctly say approval would authorize a fresh **rev11** inert kickoff. The operative sequence therefore contradicts the approval scope at the exact point that tells master what to author.
2. The `GRILL_LOCK` resolved `Lineage` decision at line 141 stops its correction trail at r9-F1 and still prints the unsuffixed pre-generation kind names `...-l4-content-req / ...-l4-content-verdict`. `GRILL_SOURCE` and `Design-lock impact` now carry r10, but the block labelled `Resolved decisions` does not state the generation-unique `<r>` rule or the return-to-the-actual-approve edge. A downstream seat following that canonical decision paragraph can recreate the fixed-id defect rev11 otherwise closes.

Required correction:

1. replace the rev8 kickoff sentence with revision-safe authority: master authors one fresh inert kickoff bound to the exact plan revision and SHA that the VP approves;
2. carry r10-F1 into the GRILL_LOCK `Lineage` resolved decision using `...-l4-content-req-<r>` / `...-l4-content-verdict-<r>`, repeated-instance uniqueness, and return parented to the exact approving verdict; and
3. change only the necessary status/transmittal language around those two corrections.

This is a record-consistency correction. It does not reopen rev11's graph.

## Passed scope

- **R10-F1 CLOSED:** content-review request/verdict ids are parameterized by generation `<r>`; a revise verdict parents the corrective proposal; post-correction equality parents the next request; and return parents the exact verdict whose value is `approve`.
- **R10-F2 source/impact CLOSED:** `GRILL_SOURCE` now carries B23 and reviews r8-r10, assigns staffing to B23 and transport to B22, and the design-lock impact states unique ids for every mechanically distinct and repeated relay instance.
- The ten-row relay-kind graph, materialization/equality split, two implementer duties, rework edges, escalation resume, B23 staffing, B22 transport, write fence, owner-fidelity order, fixed schema/value set, Master+VP freeze reservation, and H-12 all remain intact.
- Sections 4/5/6/8, both void banners, the interface lock, and every owner/frozen byte remain unmoved.

## Gate disposition

- Return one bounded rev12 changing only sequence step 0's stale kickoff basis, the GRILL_LOCK `Lineage` resolved-decision tail, and necessary status/transmittal language.
- Preserve all rev11 lineage mechanics, the ten-row structure, B23, B22, the write fence, sections 4/5/6/8, both void banners, and every locked/owner byte.
- No kickoff, pair boot, proposal, materialization, equality review, content review, fixture, manifest, owner-fidelity request, freeze, re-lock, T4 token, or external use on rev11.
- Approval, when earned, will authorize only one fresh inert kickoff bound to the exact approved plan hash. The operator retains pair boot and handover.

## Verification

- Recomputed SHA-256 values: incoming `00752a996985e0148a3a34a8765d9b211998590b153575af616b822982f605bc`; plan rev11 `9f3a142c2a26fdd21bf890f5c7da193cef936eec7aa231e72e7969a4ee1fdfc2`; deviations `e117d0596deadffacb2ca21a5c1453440dc12469458ab51fa164b87b5d32448b`; void nested kickoff `302a4bb6a335f9f99e140d640852172760ad97123e49cb7e335f72e7d8c3c7cb`; void pair kickoff `3615e40252fd5ed78e57c9a0bdfb82bb785cf23d479718d8ec7cc592889f26af`; interface lock `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.
- Historical exact-file lint of the incoming relay is `OK` with `--no-freshness`.
- Forward index check is `OK master/relays/INDEX.md`.
- Fresh interface-lock constituent rehash is `rows=38 distinct=38 mismatches=0`.
- `frank/` is clean on `main...origin/main`, with HEAD and `origin/main` both `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row. No plan, deviation, kickoff, hardening record, export, seat, tool configuration, proposal, materialization, fixture, manifest, lock, owner/frozen artifact, `frank/` source, branch, commit, PLAN/T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` is clean at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260725-171506.md`.
Next requested action: master returns bounded rev12 with a hash-bound current kickoff basis and the r10 generation rule folded into the GRILL_LOCK resolved decision; every rev11 mechanic stays closed.
