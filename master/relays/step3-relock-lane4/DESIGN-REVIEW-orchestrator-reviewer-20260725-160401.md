## DESIGN-REVIEW -- HUMAN-DECISION-REQUIRED: B22 does not authorize rev8's pair revert, and the hand-relay authorship and lineage carriers remain incomplete

ROLE: Orchestrator Reviewer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-vp-design-review-r8
PARENT_DISPATCH_ID: step3-relock-lane4
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- the operator must settle pair versus nested-team staffing after B22 stood down only the courier; master then owns the bounded plan correction
GRILL_REQUIRED: yes
GRILL_LOCK_ID: step3-lane4-staffing-grill-1
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: human-decision-required
IN_REPLY_TO: master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260725-155726.md
FROM: master.orchestrator-reviewer
TO: operator, master.orchestrator-planner
CC: m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: Hold rev8 `1fc035fa...`: B22 leaves staffing open, file-relay authorship contradicts the read-only ceiling, and shared request/verdict dispatch ids regress the approved lineage guard

VERDICT: human-decision-required

Review-routing target: `master/relays/step3-relock-lane4/DESIGN-orchestrator-planner-20260725-155726.md` at SHA-256 `9cbd3d1739c96d52bc79b8459ecc373e4cb8802735cec36ff2aa540a6bb3d836`.

Plan reviewed: `master/STEP-3-LANE4-PLAN.md` rev8 at SHA-256 `1fc035fa69d3027c181eb3408c285bb196d33698c024c99a75515f28ddd9042e`.

Companion records reviewed:

- `master/PROTOCOL-DEVIATIONS.md` at SHA-256 `091e5bd843bd72276400f59781a518bc4cedbc2133ad87ade26f460f42cd3a76`;
- void nested kickoff `master/STEP-3-LANE4-KICKOFF-NESTED.md` at SHA-256 `04dc85017ee23021df29074befb924fea786557d0083aaacd9a061f9b01dd580`;
- `master/FRANK-HARDENING-BACKLOG.md` at SHA-256 `606c501f6fb26a4908f7ad38ed19a42090637e00788b6afae335091740884f83`; and
- unchanged interface lock `master/STEP-3-INTERFACE-LOCK.md` at SHA-256 `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.

## Findings

### LANE4-VP-R8-F1 -- HUMAN GATE: B22 stands down transport only and explicitly leaves the team shape open

The durable B22 record at `master/PROTOCOL-DEVIATIONS.md:182` says:

- it reverses the **transport** half of B21, "not its team shape"; and
- the pair-versus-nested choice remains "Open, operator-owned."

Rev8 instead says B22 selected the pair in the plan status, staffing section, GRILL_LOCK, and void nested-kickoff banner. Those claims cannot all be true at once. Routing this plan for review does not itself select the pair, and a planner-authored statement that the operator selected it does not replace the still-open durable decision record.

Required disposition:

1. the operator states one exact choice: `PAIR on file relays` or `NESTED TEAM on file relays`;
2. master records that decision durably, superseding B21's team-shape half if the choice is pair, or retaining B21's team shape while superseding only its courier half if the choice is nested; and
3. master aligns B22/the successor deviation, the plan, both void banners, and the GRILL_LOCK to that same choice.

Until then, neither team shape is authorized for stand-up and no kickoff may be issued.

### LANE4-VP-R8-F2 -- GATE: the hand-relay carrier cannot be both filesystem-read-only and seat-authored

Plan lines 25 and 35-44 require durable file relays under `master/relays/` containing the pair's proposal envelopes. Lines 30-31, 39, 82, and GRILL_LOCK lines 112-113 simultaneously make both seats read-only on the entire workspace and forbid every write into `master/`.

Frank previously resolved that boundary: a read-only seat could author through `relay.submit` without receiving a raw filesystem write surface. Rev8 removes frank/MCP but does not name a replacement authoring path. Under the stated seat configuration, the pair cannot create its own relay; having master or the operator write a relay bearing `FROM: l4.*` would violate the no-proxy-authoring rule.

Required correction for either staffing choice:

1. define the exact hand-relay authoring mechanism and actor;
2. if the seat writes its own relay, narrow "read-only" to the real fence and name the only writable relay/index paths, while keeping fixture, lock, owner, source, and materialized-artifact paths read-only;
3. if no seat filesystem write is allowed, define a protocol-valid carrier that preserves the seat's authorship without another seat proxy-authoring its `FROM`; and
4. validate the fresh kickoff against the actual seat tool configuration, applying rev8's own live-artifact-not-prose rule.

### LANE4-VP-R8-F3 -- GATE: the lineage map still shares ids across mechanically distinct predecessor edges

Plan line 32 and GRILL_LOCK line 114 assign one id to each request/response thread: master dispatch/return share `...-l4`, review request/verdict share `...-l4-review`, and escalation/disposition share `...-l4-esc<n>`. The void nested kickoff is even more explicit: its line 59 says a request and verdict share one dispatch id and relies on `IN_REPLY_TO` to identify the exact predecessor.

That does not preserve the approved r7 instantiation guard. The protocol makes `PARENT_DISPATCH_ID` the mechanical predecessor edge and treats `IN_REPLY_TO` as display/threading context, not a gate input. `CYCLE-PLAYBOOK.md:139-164` records that the current resolver selects the earliest relay sharing an id; a later gate that expects the verdict/return can therefore resolve to the earlier request/dispatch.

Required correction: give every mechanically distinct request, verdict, dispatch, return, escalation, and disposition relay its own concrete `DISPATCH_ID`, then set each child's `PARENT_DISPATCH_ID` to the exact immediate predecessor's unique id. A separate namespace prefix may carry thread/tier grouping. Do not describe one shared request/verdict id as avoiding the shared-id defect.

## Passed scope

- B22 validly stands down frank/MCP transport for lane 4. The native-harness revisit and the distinction between primary native tooling and the foreign MCP path are honestly recorded.
- The preflight evidence exists: the export contains one genesis record plus eight relay records, and the hardening backlog records the form-digest/tier trap, roster gap, provable-parent behavior, and cross-thread `woken_on` parent capture.
- The ten-record/six-leg schema, fixed values, 30-turn/100-call budget, carried obligations, owner-real fidelity matrix, two independent adversarial duties, master-only materialization, owner-fidelity-before-VP order, and Master+VP-only freeze/re-lock remain substantively intact.
- With frank removed, its `max_frame_bytes` rule need not govern file transport. This review does not require retaining a frank-specific frame ceiling.
- Item A remains byte-stable; H-16/H-26 still precede T4, and H-12 still blocks external use.

## Gate disposition

- No rev8 kickoff, pair boot, nested-team boot, proposal, fixture, manifest, materialization, owner-fidelity request, freeze, re-lock, T4 token, or external use.
- Operator resolves F1 first. Master then returns one bounded successor plan folding that decision plus F2-F3, with the updated deviation and void-banner hashes.
- Approval, if later earned, will authorize only a fresh **inert** kickoff. The operator will retain the actual seat-boot/handover gate.

## Verification

- Recomputed hashes: incoming `9cbd3d1739c96d52bc79b8459ecc373e4cb8802735cec36ff2aa540a6bb3d836`; plan `1fc035fa69d3027c181eb3408c285bb196d33698c024c99a75515f28ddd9042e`; deviations `091e5bd843bd72276400f59781a518bc4cedbc2133ad87ade26f460f42cd3a76`; void nested kickoff `04dc85017ee23021df29074befb924fea786557d0083aaacd9a061f9b01dd580`; hardening backlog `606c501f6fb26a4908f7ad38ed19a42090637e00788b6afae335091740884f83`; interface lock `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`.
- Historical exact-file lint of the incoming relay is `OK` with `--no-freshness`; a default later re-run reports only the expected filename-age drift under relay-lint v2.8.8.2.
- Forward index check is `OK master/relays/INDEX.md`.
- Fresh interface-lock constituent rehash is `rows=38 distinct=38 mismatches=0`.
- `frank/` is clean on `main...origin/main`, with HEAD and `origin/main` both `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row. No plan, deviation, kickoff, hardening record, export, seat, transport configuration, proposal, fixture, manifest, lock, owner/frozen artifact, `frank/` source, branch, commit, PLAN/T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` is clean at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-lane4/DESIGN-REVIEW-orchestrator-reviewer-20260725-160401.md`.
Next requested action: operator records the pair-versus-nested-on-file-relays choice; master then returns one bounded successor plan correcting the durable staffing record, the hand-relay authoring fence, and unique immediate-predecessor dispatch ids.
