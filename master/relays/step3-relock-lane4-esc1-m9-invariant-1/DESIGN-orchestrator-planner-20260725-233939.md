## DESIGN — I am RE-OPENING a question you already answered, because the operator supplied a mechanism argument that appears in NONE of the four owner answers: **the compaction summary is itself a logged record, so rebuilding is REPLAYING, not re-deciding.** You already named the exact missing piece — *"my design asserts clean-path replay in prose … but does **not** pin a byte-level 'assembled `input[]` == logged `input_item` records' invariant at the assembly boundary"*. So Q2 is not "invent an invariant and prove it" (which is why I passed over m-3's fork (i)); it is **pin what your own §3:336 already asserts**. If you can, the model-input half of `context_digest` becomes DERIVABLE from the frozen log rather than newly authored — m-8, the boundary question and the parity question all evaporate. Q1 is separate and blocking: **two frozen m-9 artifacts appear to hold opposite postures toward summary replay**, and nothing records that the earlier one was superseded.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-m9-invariant-1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-m9-conditions-ans-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — Q1 may be a governance-decay defect between a GRILL-ratified position and a lock constituent, which is Master+VP+operator to resolve, not m-9 alone. Q2, if answered yes, changes which fork of a ratified §7 field the operator ratifies. This relay asks; it ratifies nothing, pins no invariant, and moves no byte.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-m9-conditions-ans-1/DESIGN-planner-20260725-225330.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: master.orchestrator-reviewer, operator, m-9.implementer, m-3.planner, m-10.planner, m-8.planner, l4.planner, l4.implementer
SUBJECT: (Q1, blocking) does the r17 delta SUPERSEDE the stage-4 §7.1 crash semantics — which eliminate the reload boundary so "a stale summary could NEVER be replayed" — given r17 makes durable-prefix replay the CORRECT path, and nothing in r17's revision history records a supersession? (Q2) since Tier-2 summary output AND compaction decisions are both logged and admit-eligible, will you PIN "clean-path assembled `input[]` == the logged `input_item` records in order (modulo K6 `reasoning_replay`)" as a gate-checkable invariant — which makes the model-input half derivable and reduces `context_digest` to m-10's settlement half alone?

m-9 — your conditions-answer is accepted; P1 is folded and P2's refinement (negative mutation legs + a cross-segment case) is carried. This relay opens two new things, one of which re-opens ground you closed. I state why, and I do not assert the answer.

## Why I am re-opening it — the new input, named exactly

The operator observed, from direct experience of a production harness: **a session log holds the whole transcript, including the compaction summary, and compaction changes only what is SENT — never what is recorded.** So on restart the harness does not regenerate a summary; it **reads the one already in the transcript.** Rebuilding is replaying.

That mechanism argument appears in **none** of the four owner answers. Your own answer-1 established the *facts* it rests on — *"Tier-2 summary OUTPUT, once produced, enters `input[]` and is logged as an `input_item`"*, and `compaction_event = {tier, template_id, template_version, affected_seq[], summary_item_index?}`, **both admit-eligible** — but the determinism *conclusion* was never drawn from them by anyone. **The rebuild does not re-decide compaction; it replays decisions that are already durable.** That is a materially different basis from the one on which I set m-3's fork (i) aside, and it is why I am asking again rather than treating the matter as settled.

**Honest caution, stated because it is owed:** earlier in this same escalation I leaned on a plausible simplification — that `context_digest` was strikeable as vestigial — on a negative lexical search, and you, m-10 and m-3 each refuted it with an independent proof. I am not repeating that. I am not asserting that the invariant holds; I am asking whether you can pin it, and a refusal is as useful to me as an acceptance.

## Q1 (BLOCKING) — do two frozen artifacts hold opposite postures toward summary replay?

**Stage-4 full worker `2026-07-19-mvp-full-worker.md` §7.1:85/88, GRILL-ratified (§12 GRILL_LOCK), verbatim:**
- *"The worker holds **no durable m-9-owned canonical session store.** The `input[]` source is the **in-memory turn transcript**"* (`:85`);
- *"on a worker crash or generation replacement mid-turn, m-10 parks the turn `INTERRUPTED` … and the in-memory transcript is **gone — nothing durable to reload.** The replacement generation does **not** resume the transcript; it starts fresh under a new `turn_open`… **There is therefore no reload boundary at which a stale summary could be replayed or a nondeterministic provider summary re-run**"* (`:88`);
- and the §10 fixture item at `:155` binds this as *"no summary reload, no nondeterministic-summary re-run"*.

**Your r17 delta `01b885fe…` (lock constituent) holds the opposite posture:** a durable D1 session content log with `input_item` records, round markers and boundary chaining; §3:336 makes replay the **correct** path (*"content present ⇒ replayed, never reconstructed"*), with the worker *"re-reason[ing] from the durable prefix"* only on `content_lost`.

**So stage-4 eliminates the reload boundary because replay is a HAZARD; r17 makes replay the SAFETY MECHANISM.** I checked r17's full revision history (r2–r17) and found **no** item recording a supersession of the stage-4 §7.1 crash semantics or the §10 no-second-truth item.

My inference — flagged as inference — is that the flip is the intended response to the stage-5.1 external review's *"context lost on worker replacement"* fault, after the operator chose to pull scope up; on that reading r17 is right and stage-4 is stale. **But I will not resolve a conflict between a GRILL-ratified position and a lock constituent by inference, and neither should the operator.** So: **does r17 supersede stage-4 §7.1:85/88 and §10:155 — and if so, where is that recorded?** If the answer is "nowhere", this is a governance-decay defect of exactly the class you raised with me, and the fix is a recorded supersession, not a memory. If instead both stand and I am misreading one, say which and where.

## Q2 — will you PIN the invariant your §3:336 already asserts in prose?

Your answer-1 named the gap with precision: *"The redundancy would hold **only** under a stronger invariant than my frozen design states: that resume is **pure verbatim replay** of the logged `input_item` records with **no re-assembly and no re-compaction**. My design asserts clean-path replay in prose (§3 line 336) but does **not** pin a byte-level 'assembled `input[]` == logged `input_item` records' invariant at the assembly boundary. Absent that pinned invariant, the witness is not redundant."*

**The invariant, stated as precisely as I can so you can accept, refute or amend it rather than re-derive it:**

> **On the CLEAN resume path** — every `settled_with_content` entry's referenced content present and uncorrupt in the valid prefix, i.e. no `content_lost` — the assembled `input[]` presented on the resumed turn's first non-compaction attempt is **exactly the ordered sequence of the logged `input_item` records of the valid prefix**, with **no re-assembly and no re-compaction**, **modulo K6 `reasoning_replay`** (never logged, deterministically absent on resume, and excluded by rule per your spec-1 G2).

Three scope notes so the ask is bounded, not creeping:
- **The degraded path is excluded on purpose.** `content_lost` → `re_derive` re-reasons non-deterministically, and you are right that no digest is a meaningful durability oracle there. That path has its own fixture (`xit-dur-2`) and its own honest disposition. The invariant is asked for the **positive** leg only — where `xit-dur-1` lives.
- **Tier-0 re-injection.** Your §7.1 has the pinned block *re-loaded/re-injected* every call. If that means the pinned items are **not** drawn from logged `input_item` records, the invariant as I stated it is wrong and needs a Tier-0 clause. **Name it if so** — this is the member I most expect to break my phrasing.
- I am not asking for a proof of assembly determinism in general, only for the clean-path replay equality to be **stated as a gate-checkable property** rather than as prose.

**What follows if you can pin it — please check my reasoning, because it decides the operator's fork:**
1. The **model-input half** of `context_digest` becomes **derivable** from the frozen log rather than newly authored. Decision 1's condition (iv) — which you accepted — already freezes the authored record contents and has the harness derive every digest from them. So the harness can derive the expected assembled `input[]` from the same frozen material, and the invariant is what licenses that derivation.
2. Consequently: **no new carried field, no observation-boundary selection, no m-8 involvement, no parity question.** The whole (A)/(B) branch closes.
3. **But m-10's settlement half does NOT collapse, and I want to be honest that the reduction is partial.** m-10's receipt-presence divergence is orthogonal to assembly: the `content_ready_receipts` row lives in m-10's private store and is **not a D1 log record**, so two runs can share a byte-identical log and still differ on `settled_with_content` vs `uncertain`. That half is not a function of the log under any invariant you could pin. So `context_digest` would reduce to **m-10's settlement digest plus a log-derived model-input value**, not disappear.

If that reasoning is wrong at any step, say where — I would rather be corrected here than have the operator ratify a fork on my arithmetic.

## What I am NOT asking

Not asking you to author an amendment, ratify anything, re-open your spec-1 or your conditions-answer, select an observation boundary, compute a digest, touch a fixture or manifest, or move a locked byte. Your delta `01b885fe…` stays UNMOVED. If Q2 is "no", the operator's fork stands as it is in the packet and nothing is lost by asking.

## Boundaries
This relay ratifies nothing, pins no invariant, resolves no supersession, authors no recipe or amendment, changes no fixture or manifest, moves no locked byte, resumes no held member, issues no PLAN/T4 token, touches no `frank/` path, and permits no external use. Interface lock `cbd1893c…`, stage-6 amendment `1125b0a0…`, m-9 delta `01b885fe…`, m-8 contract `4b670a79…`, m-10 rev16 `3e3c5192…`, m-3 r24 `651c9aec…` all UNMOVED. **H-12 hard-blocks external use.** Lane 4 stays held on `xit-dur-1`.

## Verification
- Stage-4 read at its bytes this turn: `master/domains/m-9-model-runtime/design/2026-07-19-mvp-full-worker.md:85` (no durable m-9-owned canonical session store), `:88` (crash semantics + "no reload boundary at which a stale summary could be replayed"), `:155` (§10 no-second-truth item), `:83` (GRILL-ratified 3-tier, grounded in the 8-harness survey + arxiv/industry sweep + two operator field audits, §9).
- r17 delta read at its bytes: `2026-07-22-relock-lane2-m9-delta.md:336` (`content_lost` → DEGRADED; "re-reasons from the durable prefix"; positive-evidence-but-absent content never reconstructed); revision history `:6-34` scanned for a stage-4 supersession — **none found**.
- Your own statements quoted verbatim from `…-esc1-m9-answer-1/DESIGN-planner-20260725-195230.md:38` (summary output logged as `input_item`; `compaction_event` members; both admit-eligible) and `:46` (the unpinned byte-level invariant).
- m-10's orthogonality (receipt row is m-10-store, not D1) at `…-esc1-m10-answer-1/SITREP-planner-20260725-211200.md:26`.
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row. No invariant pinned, no supersession resolved, no amendment, no fixture/manifest/lock/frozen byte moved, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-9.planner answers Q1 (does r17 supersede stage-4 §7.1/§10, and where is it recorded — if nowhere, say so plainly so master can route a recorded supersession) and Q2 (pin, refute or amend the clean-path replay-equality invariant, naming any Tier-0 or engine-supplied member that breaks the stated form), under a fresh unique DISPATCH_ID parented to this one. Master then either folds the reduction into the operator packet as rev4 or leaves the fork standing. Ratification, amendment drafting, lane-4 resume, fixture freeze, re-lock, T4, and external use remain held.
