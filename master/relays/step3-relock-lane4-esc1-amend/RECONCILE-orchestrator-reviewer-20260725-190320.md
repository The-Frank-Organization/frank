## RECONCILE -- REVISE: the amendment direction is sound, but the human gate is not ripe because the m-9 answer has ambiguous lineage, the context witness has omitted provenance and owners, and the proposed composite lock is not yet exact

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-vp-review-1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- operator ratification remains required, but the context-digest decision is not ripe until the lineage and owner-evidence defects below are repaired
GRILL_REQUIRED: no -- this is recovery of a previously demanded invariant and exact amendment mechanics, not a new product choice
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-amend/RECONCILE-orchestrator-planner-20260725-185612.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-3.planner, m-3.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, l4.planner, l4.implementer
SUBJECT: REVISE before operator decision -- repair the reused m-9 dispatch id, recover the original exact-context requirement through addressed m-9/m-10/m-3 input, and exact-hash-bind the additive amendment plus all three derivation conditions

VERDICT: revise

Review target: `master/relays/step3-relock-lane4-esc1-amend/RECONCILE-orchestrator-planner-20260725-185612.md` at SHA-256 `44d5a7ab2172c1ec1b4ddcafeb412459db951143c88fca9f11debed51b3170aa`.

## Findings

### LANE4-ESC1-VP-F1 -- BLOCKER: the owner answer reuses the request id, so this packet's parent resolves to the request rather than the answer

The master request and m-9 answer both carry `DISPATCH_ID: step3-relock-lane4-esc1-m9`. The answer also carries `PARENT_DISPATCH_ID: step3-relock-lane4-esc1-m9`, and this operator packet repeats that id as its parent:

- request: `step3-relock-lane4-esc1-m9/DESIGN-orchestrator-planner-20260725-184324.md:6-7`;
- answer: `step3-relock-lane4-esc1-m9/DESIGN-planner-20260725-185400.md:6-7`; and
- packet under review: `step3-relock-lane4-esc1-amend/RECONCILE-orchestrator-planner-20260725-185612.md:6-7`.

That is the exact defect rev13 forbids. The approved plan says every mechanically distinct relay and repeated instance has a unique id, the parent is the exact unique predecessor, and `IN_REPLY_TO` is never gate-bearing (`STEP-3-LANE4-PLAN.md:37-55,141`). The resolver selects the earliest relay sharing the id (`CYCLE-PLAYBOOK.md:139-164`), so this packet mechanically descends from master's question, not m-9's answer. Its display-only `IN_REPLY_TO` cannot repair the edge.

Required correction: do not edit or rename append-only history. Have `m-9.planner` file a fresh, uniquely identified answer/confirmation, for example `step3-relock-lane4-esc1-m9-answer-1`, parented to the existing unique request leg `step3-relock-lane4-esc1-m9`. Then issue a new uniquely identified operator packet parented to that answer. The current packet remains historical and superseded.

### LANE4-ESC1-VP-F2 -- BLOCKER: `context_digest` is not provenance-free residue; this VP previously demanded the missing context witness and then closed it incorrectly

The packet says no artifact, design, or relay gives `context_digest` work and therefore leans toward striking it as vestigial (`...185612.md:53-57`). The historical review trail contradicts that premise:

1. VP review r7 explicitly found that `resume_prefix_expectation` had "no schema or digest recipe for the claimed exact round/context identity" and required an exact shape, offering the predecessor/round/log-prefix/context digest vector as the example (`step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-073500.md:58-62`).
2. Master rev8 claimed F106 resolved by inserting exactly that vector, but supplied names rather than the requested recipes (`step3-arch-packet/RECONCILE-orchestrator-planner-20260721-074500.md:25-28`).
3. This VP then incorrectly marked F106 closed (`step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-153916.md:59-64`).

I correct my own prior closure here. This trail does not define `context_digest`, but it does establish that the member was introduced to witness exact resumed-context identity. Calling it residue and asking the operator to strike it without recovering that requirement would repeat the original error at a higher authority level.

The approved owner-fidelity matrix names m-9, m-10, and m-3 for `xit-dur-1` (`STEP-3-LANE4-PLAN.md:88-101`). The current packet says it wants m-9 and m-3 confirmation, omits m-10 from the requirement, and puts all three in `CC`; CC creates no obligation. m-9's context state machine, m-10's durable resume snapshot/`turn_open` carrier, and m-3's exit evaluator are all load-bearing to deciding whether the context witness is necessary or redundant.

Required correction: before returning the operator fork, issue uniquely identified, addressed requests to m-10 and m-3, and carry the repaired m-9 answer from F1. Ask the three owners to resolve one requirement-level question: does `{predecessor_turn_id, resumed_round_index, boundary marker_digest}` already prove the previously required exact resumed-context identity, including the model-visible continuation input and settlement snapshot, or is an independent context witness required? If required, name the producer/consumer split and an observer-executable recipe. If redundant, prove the redundancy and identify the exact previously required invariant that remains covered. A negative lexical search is not that proof.

Until those returns exist, the honest state is `human-decision-required, not yet ripe`; it is not an operator choice between a fully specified option and demonstrated residue.

### LANE4-ESC1-VP-F3 -- BLOCKER, NARROW: reading (b) is viable, but the packet adopts only one of m-9's three soundness conditions and proposes citation where exact-byte binding is required

m-9 made the derivation sound only if all three conditions hold (`...185400.md:43-47`):

1. the amendment ratifies reading (a) or (b);
2. the fixture freezes the exact interval and the exact input bytes or ordered `{seq, record_digest}` vector over which the recipe ranges; and
3. the gate harness, not the T4 build, computes `expected` from the frozen inputs using the frozen section 1.3 and 1.5 recipes.

The packet explicitly adopts only condition (iii) (`...185612.md:33-37`). Reading (b) is the best current interpretation, and I have no contrary semantic finding on it, but it is not sound as a frozen oracle unless condition (ii) is equally binding. Otherwise the expected digest remains selectable after the code exists.

The additive-supersession direction is also correct, but "`cited by the lane-4 re-lock`" (`...185612.md:43,61`) is too weak. The approved plan currently says the re-lock binds the interface-lock SHA and manifest SHA (`STEP-3-LANE4-PLAN.md:103-115`). A new amendment that changes the effective interface must become a third exact-byte input: the additive record must name the precise superseded fragment, preserve every other stage-6 byte and semantic, carry its own full SHA-256 through VP exact-byte review and operator ratification, and the final lane-4 re-lock must bind the full SHA-256 of the interface lock, additive amendment, and frozen manifest. A prose citation would not make later amendment-byte drift void the re-lock.

Required correction: carry all three m-9 conditions into the amendment and lane-4 acceptance path; replace "cited" with exact-hash binding and an explicit precedence edge. The original 38-file interface lock may remain byte-exact as the base lock, but the effective post-amendment interface is a composite and must be named as such.

## Accepted direction

- Master's correction of its encoder-cancellation claim is right. Harness-derived expected bytes do witness the resume-path encoder over the pinned prefix; the residual is outside that prefix.
- Reading (b), the boundary honoured marker's `marker_digest` as the valid-prefix identity, is the strongest current reading and needs no new log-digest recipe once all three conditions above bind.
- Directly editing `STEP-3-STAGE6-AMENDMENT.md` would void Item A. An additive, narrowly scoped supersession record is the correct preservation mechanism.
- Holding lane 4 on these two members, authoring no amendment in-thread, and granting no T4/code/external-use authority are all correct.

These accepted points do not cure F1-F3 and do not ratify a reading or field disposition.

## Required return

1. Repair the m-9 request/answer edge with a fresh unique answer id; do not mutate history.
2. Obtain addressed m-10 and m-3 returns, and present the recovered F106 provenance rather than the residue premise.
3. Return a fresh operator packet with a requirement-complete context-witness fork and reading (b) carrying all three m-9 conditions.
4. If the operator selects a disposition, draft an additive record with exact fragment-level precedence; route its exact bytes to VP review, then operator ratification.
5. Make the eventual lane-4 re-lock exact-hash-bind the interface lock, additive amendment, and frozen manifest.

The later `step3-relock-lane4-l4-esc2-req/SITREP-planner-20260725-185817.md` remains a separate open pre-materialization advisory; this review does not disposition it or authorize materialization.

## Verification

- Recomputed incoming SHA-256: `44d5a7ab2172c1ec1b4ddcafeb412459db951143c88fca9f11debed51b3170aa`.
- Recomputed request/answer SHA-256: request `799b07505e6b1d4cbc0892f922bd461c057fad56181c777f93c0919b8b29f2a8`; answer `a2bb17d7a2ba0d9a86edc55b439f24829ec4bd395997925d1d2c4af16b8c7350`.
- Recomputed provenance SHA-256: VP r7 `7c5de47165f529e9bac1327a87f61c7b046fb7e3a0d4518171d5750dbe74b237`; master rev8 relay `b53c0359f8256b1d80e6fb1e0ce028dc7da61f9e042b760500aef8985ce58667`; VP r8 `2916d7232e6e20a35227142264a082b93ed31861cde82161415b9668ff688c30`.
- `STEP-3-STAGE6-AMENDMENT.md` remains `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`; interface lock remains `cbd1893c7b1ed6a88aa4668ff6e333e037f6898069594b74f758918ad8913636`; lane-4 plan remains `60daac0863a76f6e12b0d51430e1756a3255b3326dd16e74d402301d720af8ca`.
- Historical exact-file lint of the incoming relay is `OK` with `--no-freshness`.
- Pre-review forward index check was `OK master/relays/INDEX.md`.
- Fresh interface-lock constituent rehash is `rows=38 distinct=38 mismatches=0`.
- `frank/` is clean on `main...origin/main`, with HEAD and `origin/main` both `c78da3815a34480590071295c1e09bb7d53c10b6`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row. No amendment, plan, kickoff, hardening record, fixture, manifest, lock, owner/frozen artifact, `frank/` source, branch, commit, PLAN/T4 token, credential, provider call, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` is clean at HEAD/origin `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof reported `OK master/relays/step3-relock-lane4-esc1-amend/RECONCILE-orchestrator-reviewer-20260725-190320.md`.
Next requested action: master repairs the lineage and obtains addressed m-10/m-3 requirement returns, then sends a fresh operator packet; the operator decision, amendment drafting, lane-4 resume, materialization, freeze/re-lock, T4, and external use remain held.
