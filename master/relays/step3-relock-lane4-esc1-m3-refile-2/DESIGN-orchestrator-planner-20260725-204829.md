## DESIGN — two asks. (1) Lineage repair: re-file your finding under a fresh unique id parented to my request — my defect, not yours (I dispatched you and m-10 without distinct answer ids after already being corrected on this class for m-9). (2) The evidence-owner question that now blocks the operator packet: **which observation boundary can your E3 evaluator actually capture and compare?** m-9's proposed recipe names two boundaries in one sentence — "captured from the wire" AND "in the exact `m8.llm_request.v1` order" — and m-8's frozen contract §1.1 says `m8.llm_request.v1` is "the app-internal request (**NEVER a wire object**)". Those are different surfaces with different owners. Your Q1 answer made the witness's *necessity* airtight; the boundary is the thing that decides whether the witness is *executable by the gate you own*. Also carried: I withdrew a false attribution that wrongly closed m-10's epoch/lease scope question — one narrow evidence question follows from that.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-m3-refile-2
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this relay requests a lineage re-file and two evidence-owner answers. It ratifies nothing, selects no boundary, authors no amendment, and moves no byte. Boundary selection and amendment stay Master+VP+operator.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-fork-1/RECONCILE-orchestrator-reviewer-20260725-203759.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: master.orchestrator-reviewer, operator, m-3.implementer, m-9.planner, m-10.planner, m-8.planner, l4.planner, l4.implementer
SUBJECT: (1) Re-file your `xit-dur-1` finding under a fresh unique DISPATCH_ID parented to `step3-relock-lane4-esc1-m3`; (2) as exit-gate owner, state which observation boundary the E3 evaluator can capture + compare (app-internal `m8.llm_request.v1` vs provider-lowered wire), whether the expected side can be frozen without running the build, and whether epoch/lease needs an evidence term

m-3 — your Q1/Q2 answer stands and is not reopened. The re-derivation-scores-PASS argument is what settled the requirement, and I have withdrawn my "strike it as vestigial" lean without qualification.

## Ask 1 — re-file under a fresh unique id (lineage repair)

Your finding was filed as `DISPATCH_ID: step3-relock-lane4-esc1-m3` with `PARENT_DISPATCH_ID: step3-relock-lane4-esc1-m3` — the same id as my request, self-parented. The resolver takes the **earliest** relay sharing an id (`CYCLE-PLAYBOOK.md:139-164`), so your answer is not uniquely addressable and its parent resolves to itself. Please **re-file the same substance** under a fresh unique id — suggested `step3-relock-lane4-esc1-m3-answer-1` — with `PARENT_DISPATCH_ID: step3-relock-lane4-esc1-m3` (my request), leaving the original in place as history, exactly as m-9 did.

**The defect is mine.** I was corrected on this class for m-9 and then dispatched you and m-10 without supplying distinct answer ids, creating two more instances of it. The cost is visible in this relay's header: I cannot parent to your *answer*, so I parent to my own request.

## Ask 2a — the boundary question, which is yours because the gate is yours

m-9's proposed model-input half is a SHA-256 over the JCS of the assembled `input[]` on the resumed turn's first attempt. Its capture instruction names two surfaces at once, and m-8's frozen contract forbids treating them as one:

> m-8 `2026-07-17-mvp-provider-contract.md:29` — "`m8.llm_request.v1` — the app-internal request (**NEVER a wire object, NEVER a conductor payload**)"

- **(A) app-internal** `m8.llm_request.v1` — the object m-9 hands m-8. Already canonically shaped, single owner, but it is **m-9's own output**, which weakens the independence that made the witness attractive (it still witnesses the assembly, which no existing digest does; it does not witness the lowering).
- **(B) provider-lowered wire** — the `openai-responses` bytes m-8 actually sends. Genuinely "what the model received", but the projection is **m-8's**, so the preimage is a two-domain object and m-8 must state its determinism and canonical form. I have asked m-8 that in parallel (`step3-relock-lane4-esc1-m8-boundary-1`).

**Your answer, as the owner of the E3 exit predicate:** which of these can the evaluator actually observe and compare within the mechanism you bound at r24 — the same capture path you already rely on for `logical_surface_digest`, or a different one? If (B) requires an observation point your evaluator does not have, that is decisive and I need to know it before the operator chooses, not after.

## Ask 2b — can the EXPECTED side be frozen without running the build?

m-9's binding condition (iii) — which I accept — is that the **harness**, not the T4 build, computes `expected` from frozen inputs. For a digest over an assembled `input[]` that means the fixture must freeze the **exact expected preimage** (the ordered item vector) and the harness derives the digest from it, with **actual** coming from the observation boundary. Confirm that this is realizable in your evaluator, and name anything the fixture would have to pin that it does not pin today. If the expected value could only be obtained by first running a build and recording what it produced, the leg is circular and I need that stated in the operator packet as a blocker rather than discovered at T4.

## Ask 2c — epoch/lease, following a withdrawal I owe you notice of

I wrongly closed m-10's open epoch/lease scope question by attributing to VP r7 a two-member enumeration it does not contain (the phrase was the VP's own later "including …" question). **That closure is withdrawn**; the branch is open. From the evidence side only: does the Durability positive leg need a term proving the successor admitted at the correct `turn_epoch` under the correct lease — or is that already covered by a frozen gate/predicate you can name? Same standard you applied to `context_digest`: if its absence admits a false PASS on the leg's headline property, say so.

## What I am NOT asking

Not asking you to re-prove necessity, select the boundary, author a recipe or amendment, compute a digest, touch a fixture or manifest, or move a locked byte. Your r24 `651c9aec…` stays UNMOVED. Your `env_digest` preimage-parity and r7-mirror caveats stay carried and unreopened.

## Boundaries
This relay ratifies nothing, selects no boundary, decides no scope, authors no recipe or amendment, changes no fixture or manifest, mutates no history, moves no locked byte, resumes no held member, issues no PLAN/T4 token, touches no `frank/` path, and permits no external use. Interface lock `cbd1893c…`, stage-6 amendment `1125b0a0…`, m-3 r24 `651c9aec…`, m-9 delta `01b885fe…`, m-8 contract `4b670a79…`, m-10 rev16 `3e3c5192…` all UNMOVED. **H-12 hard-blocks external use.** Lane 4 stays held on `xit-dur-1`.

## Verification
- m-8 §1.1 app-internal/never-a-wire-object read verbatim at `master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md:29`; single dialect `openai-responses` at `:19`.
- m-9's two-boundary capture sentence at `…-esc1-m9-answer-1/DESIGN-planner-20260725-195230.md:49-51`; condition (iii) at `:59`.
- Your §6:389 exclusion citation and the "degraded re-derivation never satisfies `xit-dur-1`" clause re-read at `…-esc1-m3/DESIGN-planner-20260725-195200.md:29`, `:37`, `:52`.
- The withdrawn paragraph: `…-esc1-fork-1/RECONCILE-orchestrator-planner-20260725-200600.md:48`; VP r7's actual text at `step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-073500.md:62`; the real source of the phrase at `…-esc1-amend/RECONCILE-orchestrator-reviewer-20260725-190320.md:49-53`.
- Lineage rule at `master/CYCLE-PLAYBOOK.md:139-164`.
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row. No history mutated, no boundary selected, no recipe authored, no amendment, no fixture/manifest/lock/frozen byte moved, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-3.planner re-files under a fresh unique id parented to `step3-relock-lane4-esc1-m3`, and answers 2a (observable boundary), 2b (frozen-expected realizability), 2c (epoch/lease evidence term). Master then assembles one operator packet exact-hash-binding all three unique owner returns. Ratification, amendment drafting, lane-4 resume, fixture freeze, re-lock, T4, and external use remain held.
