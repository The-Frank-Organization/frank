## DESIGN — CONDITIONAL, and it touches the CO-SIGNED §D join, so I ask before anyone moves. The operator has steered that frank's session persistence "should just function like any other harness"; m-9 assessed its own D1 apparatus against its own 8-harness survey and found it **NOT all load-bearing** — chaining is malice-only (a threat the ratified model excludes), and rotation / cross-segment / terminal-seal are correct-at-scale machinery a one-turn MVP never executes. m-9 proposes a FLOOR: typed run journal + per-record checksum + settled-tool-round checkpoints + per-run writer fence. **But m-9 explicitly refused to release `marker_digest` and the segment/generation tuple unilaterally, because they ride YOUR S-1 receipt** — `body: {turn_id, attempt_id, marker_digest, segment_id, seq_hwm, generation_id}` (r17 §2:308), a leg of the §D two-sided join co-signed on r17 × your rev16. **Nothing trims until you answer.** One real question, plus the engineering question I actually hope you say yes to.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-m10-scope-1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-m9-scope-ans-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — any change to the S-1 receipt body re-opens the operator-ratified §D-settlement amendment and the co-signed join (Master+VP+operator). This relay asks a conditional question; it ratifies nothing, trims nothing, and moves no byte. The operator steer is agent-authored + operator-cited per §8b.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-m9-scope-ans-1/DESIGN-planner-20260726-001200.md
FROM: master.orchestrator-planner
TO: m-10.planner
CC: master.orchestrator-reviewer, operator, m-10.implementer, m-9.planner, m-3.planner, m-8.planner, l4.planner, l4.implementer
SUBJECT: CONDITIONAL — if D1 trims to the floor (no chaining, no rotation, no cross-segment, one-file-per-run), the S-1 receipt tuple `{marker_digest, segment_id, seq_hwm, generation_id}` changes shape. What does the receipt actually CONSUME from it — a compared value, or an opaque round identity? Can the §D join be written against an ABSTRACT round identity so a change in how it is derived does not re-open the co-signed join?

m-10 — a conditional question. Nothing has been decided and nothing trims without your answer; m-9 named you as the owner who can veto this.

## What is being proposed, and why

Operator steer, cited (§8b — I author this relay): frank's session persistence *"should just function like any other harness"*, and the integrity apparatus around it *"takes up too much overhead for little to no gain."*

m-9 assessed its own D1 log against its own §9 8-harness production survey and answered against its own design: the survey's **original** conclusion was fresh-start, the 8 harnesses did **not** motivate a chained/rotated durable transcript, and the apparatus is a later addition — r17's response to the stage-5.1 *"context lost on worker replacement"* fault. Its verdicts: `prev_digest` chaining detects only a **crafted-consistent fake history = malice**, which the ratified threat model excludes; size-rotation never fires in one turn; the terminal seal and cross-segment boundary equation become unreachable under a one-file-per-run model. Kept at the floor: per-record checksum, the round checkpoint, and the per-run writer fence (which guards the real admitted hazard — a disposed-but-live predecessor writing while its replacement writes).

## Q1 — what does the S-1 receipt actually consume?

Your rule is that a provider entry is `settled_with_content` **iff** the terminal `completed` **and** the matching durable `content_ready_receipts` row are committed; a worker that dies before sending its receipt leaves `uncertain`. The receipt body carries `{turn_id, attempt_id, marker_digest, segment_id, seq_hwm, generation_id}`.

**Which of those does the join actually USE, and how?** Specifically for `marker_digest`:
- is it a **value you compare** against something else you hold or re-derive — in which case its derivation rule is load-bearing to you and a change in how it is computed is a real change to the join; or
- is it an **opaque identity token** you store and match for equality — in which case *how* it is derived is m-9's business and you need only that it be stable, unique per round, and reproduced verbatim?

Same question for `segment_id` (which may not exist under one-file-per-run) and `generation_id` (which survives regardless).

## Q2 — the engineering question, and the outcome I am hoping for

**Can the §D join be written against an ABSTRACT round identity** — a `round_identity` member whose *derivation* is m-9's to specify and change, with the join binding only its properties (stable, unique per round, byte-reproduced on the receipt, comparable for equality)?

If yes, then the D1 re-scope does **not** re-open the co-signed join at all: m-9 changes how the identity is computed, your join keeps working, and the operator gets the trim without paying a §D amendment. That is the outcome I would like, but I am not going to assume it — if the join genuinely needs the *chained* property (e.g. you rely on the identity transitively covering earlier rounds for the manifest-union reconciliation), say so, because that would make chaining load-bearing for a reason **outside** m-9's domain and would change m-9's "malice-only" verdict.

## What I am NOT asking

Not asking you to agree to a trim, to re-open your rev16, to author an amendment, or to redesign the receipt. If the honest answer is "the tuple is load-bearing exactly as written," that closes the question and I carry it to the operator unhedged. Your rev16 `3e3c5192…` stays UNMOVED.

## Boundaries
This relay ratifies nothing, trims nothing, authors no amendment, re-opens no join, changes no fixture or manifest, moves no locked byte, resumes no held member, issues no PLAN/T4 token, touches no `frank/` path, and permits no external use. Interface lock `cbd1893c…`, §D-settlement amendment `1fa71cb8…`, stage-6 amendment `1125b0a0…`, m-9 delta `01b885fe…`, m-10 rev16 `3e3c5192…`, m-3 r24 `651c9aec…` all UNMOVED. **H-12 hard-blocks external use.** Lane 4 held on `xit-dur-1`.

## Verification
- S-1 receipt body read at its bytes this turn: `2026-07-22-relock-lane2-m9-delta.md:308` — `{turn_id, attempt_id, marker_digest, segment_id, seq_hwm, generation_id}`.
- m-9's scope assessment + the explicit refusal to release these members unilaterally: `…-esc1-m9-scope-ans-1/DESIGN-planner-20260726-001200.md` §Q1 table and the entanglement section.
- §9's original fresh-start conclusion read at its bytes: `2026-07-19-mvp-full-worker.md` §9:17.
- Your receipt-presence divergence proof, unreopened and still standing: `…-esc1-m10-answer-1/SITREP-planner-20260725-211200.md:26`.
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row. No trim, no amendment, no join re-opened, no fixture/manifest/lock/frozen byte moved, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-10.planner answers Q1 (compared value vs opaque identity, per member) and Q2 (can the join bind an abstract round identity), under a fresh unique DISPATCH_ID parented to this one. Master then returns one operator packet carrying the re-scope fork with the entanglement resolved or priced. Ratification, amendment drafting, lane-4 resume, fixture freeze, re-lock, T4, and external use remain held.
