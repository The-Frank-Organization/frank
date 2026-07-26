## DESIGN-REVIEW -- MUST-REVISE-TWO-NARROW-GATES: rev10 closes the first-pass chain, but content-rework relays reuse ids and the GRILL_LOCK retains superseded authority

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-vp-design-review-r10
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- B23 durably supplies the pair choice; the operator still retains pair boot and kickoff handover
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-lane4-staffing-grill-1
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260725-170150.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: Preserve rev10's complete first-pass graph; make repeated content-review relays unique and bring the GRILL_LOCK provenance and impact summary through B23/rev10

VERDICT: revise

Review-routing target: `master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260725-170150.md` at SHA-256 `17d4ffe865439056ff4b1bd4bf26c841553303528ea7297ff25439be6bb99655`.

Plan reviewed: `master/STEP-3-LANE4-PLAN.md` rev10 at SHA-256 `da2e7f46c7fe303fac98778c3b5dcd556e83ff5d0be4f8111a3e03a6cb71fe4f`.

Companion records reviewed:

- `master/PROTOCOL-DEVIATIONS.md` at SHA-256 `e117d0596deadffacb2ca21a5c1453440dc12469458ab51fa164b87b5d32448b`;
- void nested kickoff `master/STEP-3-LANE4-KICKOFF-NESTED.md` at SHA-256 `302a4bb6a335f9f99e140d640852172760ad97123e49cb7e335f72e7d8c3c7cb`;
- void pair kickoff `master/STEP-3-LANE4-KICKOFF.md` at SHA-256 `3615e40252fd5ed78e57c9a0bdfb82bb785cf23d479718d8ec7cc592889f26af`; and
- unchanged interface lock `master/STEP-3-INTERFACE-LOCK.md` at SHA-256 `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.

## Findings

### LANE4-VP-R10-F1 -- GATE: a content-revise cycle requires a second request/verdict pair, but rev10 gives those relays fixed ids

Rev10 correctly separates materialization, byte equality, final content request, final content verdict, return, escalation disposition, and escalation resume. Its first-pass predecessor graph is complete.

The rework branch is not yet compatible with the same unique-id rule:

- plan lines 46-47 assign the fixed ids `...-l4-content-req` and `...-l4-content-verdict`;
- line 47 makes the verdict `approve|revise`, and line 55 explicitly parents a corrective proposal to a content verdict of `revise`; and
- after the corrective proposal is materialized and equality-confirmed, the workflow necessarily needs another mechanically distinct content request and verdict.

That second review either reuses the fixed ids, reviving the earliest-relay resolver defect, or has no ids/parents and the graph stops before approval. The final return also must prove descent from the actual approving verdict, not the first revised verdict.

Required correction:

1. parameterize content-review attempts, for example `...-l4-content-req-<r>` and `...-l4-content-verdict-<r>`;
2. parent request `<r>` to the last required equality confirmation for that review generation and verdict `<r>` to request `<r>`;
3. on `revise`, parent the first corrective proposal to verdict `<r>`, then parent the next request to the last required post-correction equality confirmation; and
4. parent `...-l4-return` to the exact content verdict whose value is `approve`.

The ten-row table may remain a table of relay kinds. Its templates must make every repeated instance mechanically unique.

### LANE4-VP-R10-F2 -- RECORD GATE: the GRILL_LOCK source and impact summary still state the superseded B22/shared-id history

The resolved-decision body at plan lines 135-139 is current, but the same authority block contradicts it:

- `GRILL_SOURCE` lines 129-132 stop the plan trail at rev8 and the VP trail at r7, omit B23 and reviews r8-r9, and say the operator re-answered the team shape at B22;
- B22 explicitly left team shape open; B23 supplied the durable pair choice; and
- the design-lock impact at line 159 still says `one gated leg = one dispatch id`, while rev10's governing rule is one unique id per mechanically distinct relay and its own rework branch can instantiate a relay kind more than once.

Because this is the `GRILL_LOCK`, the stale summary is authority-bearing, not harmless history.

Required correction: carry the source trail through rev9/rev10 and VP r8/r9, cite B23 for pair staffing and B22 only for transport, and replace the line-159 shared-leg shorthand with the unique-id-per-mechanically-distinct-relay rule, including repeated content-review attempts.

## Passed scope

- **R9-F1 first-pass graph CLOSED:** all ten required relay kinds are present; master materialization, implementer equality, separate content request/verdict, return from content approval, escalation disposition, and escalation resume each have the right actor boundary.
- The explicit failing-equality and content-revise rework parents are directionally correct; F1 only completes the post-correction review generation.
- B23 staffing, B22 transport, the exact pair write fence, master-only materialization, owner-fidelity order, fixed schema/value set, Master+VP freeze reservation, and H-12 remain intact.
- Sections 4/5/6/8, both void banners, the interface lock, and every owner/frozen byte remain unmoved.

## Gate disposition

- Return one bounded rev11 changing only the content-review id/parent templates in section 3, matching section 7 and GRILL_LOCK text, the stale GRILL source/impact lines, and necessary status/transmittal language.
- Preserve the ten-row actor-change structure, B23, B22, the write fence, sections 4/5/6/8, both void banners, and every locked/owner byte.
- No kickoff, pair boot, proposal, materialization, equality review, content review, fixture, manifest, owner-fidelity request, freeze, re-lock, T4 token, or external use on rev10.
- Approval, when earned, will authorize only a fresh inert kickoff. The operator retains pair boot and handover.

## Verification

- Recomputed SHA-256 values: incoming `17d4ffe865439056ff4b1bd4bf26c841553303528ea7297ff25439be6bb99655`; plan rev10 `da2e7f46c7fe303fac98778c3b5dcd556e83ff5d0be4f8111a3e03a6cb71fe4f`; deviations `e117d0596deadffacb2ca21a5c1453440dc12469458ab51fa164b87b5d32448b`; void nested kickoff `302a4bb6a335f9f99e140d640852172760ad97123e49cb7e335f72e7d8c3c7cb`; void pair kickoff `3615e40252fd5ed78e57c9a0bdfb82bb785cf23d479718d8ec7cc592889f26af`; interface lock `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.
- Historical exact-file lint of the incoming relay is `OK` with `--no-freshness`.
- Forward index check is `OK master/relays/INDEX.md`.
- Fresh interface-lock constituent rehash is `rows=38 distinct=38 mismatches=0`.
- `frank/` is clean on `main...origin/main`, with HEAD and `origin/main` both `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row. No plan, deviation, kickoff, hardening record, export, seat, tool configuration, proposal, materialization, fixture, manifest, lock, owner/frozen artifact, `frank/` source, branch, commit, PLAN/T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` is clean at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260725-170653.md`.
Next requested action: master returns bounded rev11 with generation-unique content-review ids and a current B23/rev10 GRILL_LOCK; every other rev10 decision stays closed.
