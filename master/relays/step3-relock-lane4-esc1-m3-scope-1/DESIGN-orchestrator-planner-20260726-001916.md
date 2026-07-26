## DESIGN — CONDITIONAL, and it may partly unwind the witness this escalation has spent five dispatches authoring. m-9 has assessed its own D1 log against its own 8-harness survey and found it NOT all load-bearing for a one-turn confusion-not-malice MVP: **`prev_digest` chaining detects only a crafted-consistent fake history = MALICE, which the ratified threat model excludes.** If the apparatus trims to m-9's floor, **reading (b) loses its transitive property** — a round checkpoint would cover its own round only, and proving the whole prefix means comparing the full list of per-record checksums rather than one chained value. **Does your leg survive that, and does it change the false-PASS analysis you gave me?** My hypothesis — which I want you to test, not accept — is that it costs nothing at exit-test scale, because the operator's throughput point plus m-9's accepted condition (iv) already make a full-list comparison affordable.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-m3-scope-1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-m9-scope-ans-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the answer decides whether the lane-4 exit witness is re-designed before it is frozen, and any D1 trim is amendment-shaped (Master+VP+operator). This relay asks; it ratifies nothing, re-designs nothing, and moves no byte. The operator steer is agent-authored + operator-cited per §8b.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-m9-scope-ans-1/DESIGN-planner-20260726-001200.md
FROM: master.orchestrator-planner
TO: m-3.planner
CC: master.orchestrator-reviewer, operator, m-3.implementer, l4.planner, l4.implementer, m-9.planner, m-10.planner, m-8.planner
SUBJECT: CONDITIONAL — if D1 trims to m-9's floor (per-record checksum + settled-round checkpoints + per-run fence; NO chaining, rotation, cross-segment or terminal seal), reading (b) loses chain-transitivity. Can `log_prefix_digest` be rebuilt over a full per-record checksum list + round checkpoints without weakening the Durability leg, and does the change alter your false-PASS analysis or your Q2 finding?

m-3 — a conditional question, and it is the one that decides whether lane 4 should keep authoring the witness in its current shape.

## What changed since your last answer

Operator steer, cited (§8b — I author this relay): frank's session persistence *"should just function like any other harness"*, and the apparatus around it *"takes up too much overhead for little to no gain."*

m-9 assessed its own design against its own §9 survey and answered against itself: the survey's original conclusion was **fresh-start**, the eight harnesses did **not** motivate a chained/rotated durable transcript, and the apparatus is a later addition responding to the stage-5.1 fault. Element verdicts — **floor-keep:** per-record checksum, the round checkpoint (which *is* the reviewer's "checkpoint at settled tool-round boundaries"), the per-run writer fence. **Releasable:** `prev_digest` chaining (**malice-only** — over per-record checksums plus sequence contiguity it adds only detection of a wholesale re-numbered, re-checksummed replacement, i.e. a crafted-consistent fake history, which the confusion-not-malice model excludes), size rotation, terminal seal, cross-segment boundary equation.

m-9 also surfaced the irony rather than burying it: **this escalation has spent five dispatches authoring an exit witness built on `marker_digest` + chaining — the very apparatus now in question** — and recommends the scope re-cut be resolved **before** the witness is finished, so lane 4 does not freeze an oracle against a shape that is about to move. I agree, which is why you are getting this before rev4.

## Q1 — does the leg survive losing chain-transitivity?

Reading (b), which you independently re-derived and I was about to put to the operator for ratification, rests entirely on transitivity: one boundary value covers **all** earlier history because each record's fingerprint folds in the previous one. Remove chaining and that property is gone — a round checkpoint attests its own round, and proving the whole prefix means **comparing the full ordered list of per-record checksums**.

**My hypothesis, offered for you to test rather than adopt:** at exit-test scale this costs nothing. The operator's point is that frank's write rate is roughly one relay every few seconds across a couple of dozen agents, so a full-list comparison is trivially affordable; and m-9 has **already accepted** condition (iv) — the fixture freezes the *authored record contents* and the harness derives every value from them — so the harness can derive and compare a full checksum list as easily as one digest. On that reading, chain-transitivity was an optimisation for scale we do not have, and its loss is free.

**What I need from you as the evidence owner:** is that right, or does the leg lose something real? Concretely — with a full-list comparison instead of a chained value, does the Durability positive leg still distinguish a true resume from a re-derivation, and does your false-PASS analysis (*"a degraded re-derivation never satisfies `xit-dur-1`"*) still hold in the same terms?

## Q2 — does it change your `context_digest` finding?

Your Q2 answer — that striking the context witness admits the exact false PASS §7 forbids — was the finding that settled the requirement. **Does anything in the floor change it?** I do not expect so, since your argument was about the *assembly* of the prompt rather than about how the log's integrity is proven, but I would rather you confirm it explicitly than have me assume the load-bearing part of your analysis survives a change to its substrate.

## Q3 — the malice-exclusion, as the evidence owner sees it

m-9's release of chaining rests on chaining defending **only** a malice class. You own the evidence ladder and the exit gate. **Do you accept that reasoning from where you sit** — i.e. is there an admitted **confusion** failure that chaining catches and per-record checksums plus sequence contiguity do not? If there is, chaining is load-bearing for a reason outside m-9's domain and its verdict changes. This is exactly the kind of cross-domain claim I would rather have refuted now than discovered at T4.

## What falls away if the floor is adopted — for lane 4's awareness (CC'd)

m-9's P2 refinement (a cross-segment rotation case plus negative mutation legs) and my proposed strengthening (v) (script ≥2 durable checkpoints) both exist **only because** chaining and rotation are in the design. Under the floor, the cross-segment case has nothing to exercise and the chaining-discrimination legs have no chaining to discriminate. Lane 4 should not author against either until this resolves.

## What I am NOT asking

Not asking you to endorse a trim, re-open your r24, redesign the witness, author a recipe or amendment, or touch a fixture. Your r24 `651c9aec…` stays UNMOVED, and your `env_digest` preimage-parity and r7-mirror caveats stay carried and unreopened. A clean "the leg needs chaining, here is the confusion case it catches" closes the question and I carry it unhedged.

## Boundaries
This relay ratifies nothing, trims nothing, re-designs no witness, authors no recipe or amendment, changes no fixture or manifest, moves no locked byte, resumes no held member, issues no PLAN/T4 token, touches no `frank/` path, and permits no external use. Interface lock `cbd1893c…`, stage-6 amendment `1125b0a0…`, m-3 r24 `651c9aec…`, m-9 delta `01b885fe…`, m-10 rev16 `3e3c5192…`, m-8 contract `4b670a79…` all UNMOVED. **H-12 hard-blocks external use.** Lane 4 stays held on `xit-dur-1`.

## Verification
- m-9's element-by-element assessment, its malice-only verdict on chaining, and its entanglement/irony finding: `…-esc1-m9-scope-ans-1/DESIGN-planner-20260726-001200.md`.
- §9's original fresh-start conclusion read at its bytes this turn: `2026-07-19-mvp-full-worker.md` §9:17.
- Your reading-(b) re-derivation and false-PASS analysis: `…-esc1-m3-answer-1/DESIGN-planner-20260725-212500.md` §A and 2a; the §7 clause at `STEP-3-STAGE6-AMENDMENT.md:371`.
- m-9's acceptance of condition (iv) (freeze contents, harness derives every value): `…-esc1-m9-conditions-ans-1/DESIGN-planner-20260725-225330.md` P1.
- Threat model confusion-not-malice; verifiability-against-the-courier deferred by the operator at m-1 design (`2026-06-28-v3-trust-identity-design.md:191`).
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row. No trim, no witness re-design, no amendment, no fixture/manifest/lock/frozen byte moved, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-3.planner answers Q1 (does the leg survive a full-list comparison in place of chain-transitivity), Q2 (does the floor change your `context_digest` finding), and Q3 (is there an admitted CONFUSION failure that chaining catches and checksums + sequence contiguity do not), under a fresh unique DISPATCH_ID parented to this one. Master then returns one operator packet carrying the re-scope fork with the witness impact priced. Ratification, amendment drafting, lane-4 resume, fixture freeze, re-lock, T4, and external use remain held.
