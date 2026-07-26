## DESIGN-REVIEW -- MUST-REVISE-ONE-GATE: rev9 closes staffing and authorship, but its exhaustive lineage table skips the master-materialization and two-review-duty edges

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-vp-design-review-r9
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- B23 durably satisfies the operator's PAIR choice; the operator still retains the later pair-boot and kickoff-handover gate
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-lane4-staffing-grill-1
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260725-160917.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: Preserve B23 and the exact write fence; complete the unique-id chain across master materialization, equality confirmation, final content review, and escalation resume

VERDICT: revise

Review-routing target: `master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260725-160917.md` at SHA-256 `8afb4a7a49e725a9f603eef09e59daed0f548bc73f8813d42b7ea9748412e41a`.

Plan reviewed: `master/STEP-3-LANE4-PLAN.md` rev9 at SHA-256 `a0009ef930a3a8f3bc4e2edf80e601c9b6c686c8ab24a3535ced8ef1a2206a52`.

Companion records reviewed:

- `master/PROTOCOL-DEVIATIONS.md` at SHA-256 `e117d0596deadffacb2ca21a5c1453440dc12469458ab51fa164b87b5d32448b`;
- void nested kickoff `master/STEP-3-LANE4-KICKOFF-NESTED.md` at SHA-256 `302a4bb6a335f9f99e140d640852172760ad97123e49cb7e335f72e7d8c3c7cb`;
- void pair kickoff `master/STEP-3-LANE4-KICKOFF.md` at SHA-256 `3615e40252fd5ed78e57c9a0bdfb82bb785cf23d479718d8ec7cc592889f26af`; and
- unchanged interface lock `master/STEP-3-INTERFACE-LOCK.md` at SHA-256 `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.

## Finding

### LANE4-VP-R9-F1 -- GATE: the unique-id table omits actor-changing relays that the materialization and review sequence requires

Rev9 correctly forbids shared dispatch ids and assigns unique ids to proposal, generic review request, generic verdict, escalation request/disposition, and final return. The table at plan lines 40-48 nevertheless does not represent the workflow at lines 53-61 and 98-104:

1. A planner proposal is consumed by **master**, which materializes the governed artifact and recomputes its on-disk byte length and hash.
2. Only after that action can the implementer verify proposal-to-file equality.
3. After equality is established for the complete set and final manifest, the implementer performs a separate adversarial **content** review.

The table instead jumps from `...-l4-propose-<n>` directly to a **planner-authored** `...-l4-review-req-<n>`, with no master-authored materialization receipt/request in the predecessor chain. It then supplies only one generic implementer verdict even though byte-equality and content approval are explicitly different duties at different points in the sequence. The final return therefore cannot mechanically prove that its parent is the final content approval rather than an earlier equality confirmation.

The escalation branch is also incomplete: `...-esc<n>-disp` is defined, but no resume/return-down row makes the first post-disposition relay parent to that disposition. The general next-proposal row instead points to the previous review verdict and can skip the escalation outcome.

Required correction:

1. add a unique master-authored materialization receipt/request per proposal, parented to `...-l4-propose-<n>`, carrying the materialized path, recomputed byte length/hash, and disk-action evidence;
2. add a unique implementer equality confirmation parented to that materialization relay;
3. after the complete fixture set and manifest have equality confirmations, add a unique final content-review request parented to the last required equality confirmation;
4. add a unique implementer content-review verdict parented to that request;
5. parent the pair's final return to the content-review approval, not a generic last verdict; and
6. define the post-escalation resume relay so its parent is `...-esc<n>-disp` before the flow rejoins the interrupted stage.

Use separate namespace prefixes or suffixes for grouping, but keep every row's `DISPATCH_ID` unique and every `PARENT_DISPATCH_ID` equal to the actual immediate predecessor relay id. Do not rely on operator prompting or filesystem visibility as an unrecorded handoff.

## Passed scope

- **R8-F1 CLOSED:** B23 durably records the operator's `PAIR on file relays` choice using the established agent-authored/operator-cited form, supersedes B21's team-shape half, and closes B22's open item. Plan, B22/B23, GRILL_LOCK, and both void banners now agree.
- **R8-F2 CLOSED:** the pair may write only its own relay files and post-cutover INDEX rows; every governed artifact remains read-only; master alone materializes; no proxy-authored `FROM` is permitted; the kickoff must validate the actual tool configuration before handover.
- **R8-F3 principle CLOSED:** `IN_REPLY_TO` is correctly excluded from gate semantics, shared request/verdict ids are withdrawn, and each listed relay has a unique id. This finding concerns the table's missing relays, not the rule itself.
- The exact ten-record/six-leg schema, fixed values, carried obligations, owner-real fidelity matrix, two independent implementer duties, owner-fidelity-before-VP order, and Master+VP-only freeze/re-lock remain intact.
- B22's transport stand-down, the dogfood close, Item A, H-16/H-26, and H-12 remain unchanged.

## Gate disposition

- Return one bounded rev10 changing only the lineage table, matching sequence/GRILL_LOCK text, and necessary status/transmittal language. Preserve B23, B22, the write fence, sections 4/5/6/8, both void banners except any exact rev pointer needed, and every locked/owner byte.
- No kickoff, pair boot, proposal, materialization, equality review, content review, fixture, manifest, owner-fidelity request, freeze, re-lock, T4 token, or external use on rev9.
- Approval, when earned, will authorize only a fresh **inert** kickoff. The operator retains pair boot and handover.

## Verification

- Recomputed hashes: incoming `8afb4a7a49e725a9f603eef09e59daed0f548bc73f8813d42b7ea9748412e41a`; plan rev9 `a0009ef930a3a8f3bc4e2edf80e601c9b6c686c8ab24a3535ced8ef1a2206a52`; deviations `e117d0596deadffacb2ca21a5c1453440dc12469458ab51fa164b87b5d32448b`; void nested kickoff `302a4bb6a335f9f99e140d640852172760ad97123e49cb7e335f72e7d8c3c7cb`; void pair kickoff `3615e40252fd5ed78e57c9a0bdfb82bb785cf23d479718d8ec7cc592889f26af`; interface lock `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.
- Historical exact-file lint of the incoming relay is `OK` with `--no-freshness`.
- Forward index check is `OK master/relays/INDEX.md`.
- Fresh interface-lock constituent rehash is `rows=38 distinct=38 mismatches=0`.
- `frank/` is clean on `main...origin/main`, with HEAD and `origin/main` both `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row. No plan, deviation, kickoff, hardening record, export, seat, tool configuration, proposal, materialization, fixture, manifest, lock, owner/frozen artifact, `frank/` source, branch, commit, PLAN/T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` is clean at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260725-161748.md`.
Next requested action: master returns bounded rev10 with the complete unique-id materialization/equality/content-review/escalation-resume chain; all other rev9 decisions stay closed.
