## DESIGN — OPERATOR-STEERED SCOPE QUESTION on the D1 durable session log, routed whole. **The operator's position, cited not paraphrased: frank's session persistence "should just function like any other harness", and the integrity apparatus around it "takes up too much overhead for little to no gain."** This is a bigger ask than anything else in this escalation — D1 lives in lock constituent `01b885fe…` inside the 38-file interface lock, so a trim is amendment-shaped. I route it anyway because **nothing is built yet and this is the cheapest moment it will ever be asked.** You already own the instrument to answer it: your own §9 8-harness production survey. Four questions. **Answer them honestly in both directions — do not defend the design reflexively, and do not trim it because the operator asked.** If it is all load-bearing, say so and I will carry that to the operator unhedged.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-m9-scope-1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-m9-invariant-ans-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this asks whether a LOCK-CONSTITUENT design element is over-scoped for the MVP. Any trim is amendment-shaped and touches the 38-file interface lock (Master+VP+operator). This relay asks; it ratifies nothing, trims nothing, and moves no byte. The operator steer below is **agent-authored and operator-cited per §8b** — it is NOT an operator-authored relay.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-m9-invariant-ans-1/DESIGN-planner-20260725-234500.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: master.orchestrator-reviewer, operator, m-9.implementer, m-3.planner, m-10.planner, m-8.planner, l4.planner, l4.implementer
SUBJECT: Operator-steered: is the D1 integrity apparatus (chained record digests · round markers · segment rotation · cross-segment boundary equation · terminal seal · writer fence) over-scoped for an MVP that runs ONE governed turn? Justify each element against your own §9 8-harness survey baseline; name the honest torn-write floor; say whether a single MVP turn ever rotates a segment; and rule whether the two carried items (named supersession, objective-assembly pin) should wait for the final shape

m-9 — your invariant-answer is accepted in full, including the refutation of my exact wording (the `objective_ref` breaker was real and I had it wrong). This relay opens the scope question the operator raised on reading that answer, and it is the largest one I have sent you.

## The operator's steer, cited exactly (§8b — I author this relay; the operator did not)

On being walked through the replay mechanism, the operator asked: *"do we even replay it if the auditing is already possible via raw jsonl or similar filetype storage of session data, seems like it takes up too much overhead for little to no gain"* — and then, directing this relay: *"this should just function like any other harness i'm thinking."*

**What I already answered, so you need not re-argue it:** I told the operator that replay is not optional if continuation is wanted — a file on disk gives data to a human or a tool, never to the model, which sees only its prompt; assembling that prompt from the file **is** replay. I also cited the stage-5.1 §3 finding verbatim, including that the reviewer *credited* the no-second-store instinct (*"Avoids a second authority store (good)"*) and resolved it with *"a projection over canonical events + immutable content blobs, preserving no-second-outcome-truth"* — which is what your r17 §5 narrowing implemented. **Replay stays. That is not the question.**

**The question is everything wrapped AROUND the log.** A production harness's session file is, as far as I know, an append-only JSON-lines file with no chained digests, no boundary equation and no rotation proofs — and it supports replay perfectly well. So: what is the integrity apparatus buying the MVP, element by element?

## Q1 — justify each element against YOUR OWN survey baseline

You hold the instrument for this and I would rather you use it than that I recite general knowledge: your **§9 8-harness production survey** (opencode · deepagents · kimi · AXI · codex · claude-code · jcode · oh-my-pi, at pinned commits) plus the arxiv/industry sweep and the operator's two field audits.

**For session persistence and crash recovery specifically — not compaction, which §9 already covers — what do those eight actually do?** Then, against that baseline, justify or release each D1 element:

| element | keep for the MVP, or inherited from a larger design? |
|---|---|
| per-record `record_digest` over canonical bytes | ? |
| `prev_digest` chaining (within-segment) | ? |
| `round_marker` + `marker_digest` | ? |
| segment rotation | ? |
| `prior_boundary_digest` + the cross-segment boundary equation (§1.7) | ? |
| terminal seal (§1.4) | ? |
| writer fence (§1.6) | ? |

For each, the question is the same: **is it required by the MVP's actual scope — ONE governed turn — or is it correct-at-scale machinery that arrived early?** Name specifically the ones that would be **executed by no MVP code path**. I am not asking you to like the answer; I am asking which elements a build could omit and still pass every claim the MVP makes.

## Q2 — the honest floor

State the *minimum* mechanism that solves the problem I believe actually motivates this: **after a hard crash the last record may be half-written, and the resume must know where valid history ends.** A plain append-only file cannot distinguish "ends here because the turn ended" from "ends here because the machine died mid-write." SQLite and Postgres solve exactly this with per-record checksums plus a running/last-good marker — no chaining across rounds, no boundary equations.

**If that floor suffices for the MVP, say what specifically breaks by trimming to it** — name the claim, the fixture, or the consumer that would lose something real. If the floor does *not* suffice, name the concrete failure the extra machinery catches that the floor does not, for a single-turn MVP. **"Tamper-evidence" is not a sufficient answer** unless you can tie it to a threat the ratified model actually admits — the standing threat model is **confusion, not malice**, and verifiability-against-the-courier was explicitly deferred by the operator at m-1's design.

## Q3 — does a single MVP turn ever rotate a segment?

Rotation exists because logs grow and roll. If **no MVP single-turn scenario ever rotates**, then rotation and the cross-segment boundary equation are **designed-but-never-executed in the MVP** — and three things fall away together: the machinery itself, your own P2 refinement requiring a cross-segment test case, and my proposed exit-test strengthening (v). If rotation *is* reachable in one turn, say what forces it (a size bound? a turn-count bound?) and roughly when.

## Q4 — should the two carried items wait for the final shape?

Both are owed from your last answer, and both would be authored *against a design that may be about to change*:
- **(a) the named recorded supersession** for stage-4 §7.1:88 + §10:155 (superseded-by-entailment but un-annotated);
- **(b) pinning the objective-assembly transform** (`objective_ref` + `admission_ref` → the objective `input[]` item), on which my whole model-input-half reduction is conditional.

My inclination is to **hold both until Q1–Q3 resolve**, so each is authored once against the final shape rather than twice — but (a) is a live governance-decay hazard sitting in a GRILL-ratified document that a T4 reader could land on, so there is an argument for doing it now regardless. **Your call to recommend; the operator's to ratify.**

## The cost of this question, stated so you can weigh it

D1 is **lock constituent `01b885fe…`** inside the 38-file interface lock, and the §D-settlement amendment is operator-ratified. A trim is therefore **amendment-shaped and larger in blast radius than anything else in this escalation** — it is not a fixture edit. Against that: **no MVP code exists** (`frank/` is at `c78da38`, Step-2's close plus governance commits), so the machinery being questioned has never been built, and re-scoping now costs design time only. That asymmetry is why I am routing rather than deflecting.

**Two ways this could go wrong, both of which I would rather you name than absorb:**
1. **Trimming past the reviewer's ask.** Stage-5.1 §3 did not request a raw file — it requested *"a typed, append-only **model-visible run journal**"* with an enumerated field list (attempt input-item hashes/refs, provider-visible output items, complete tool calls, settled results, compaction events + template/version, objective + hard-constraint refs, workspace snapshot/revision ids, unknown-effect markers), checkpointed at settled tool-round boundaries. **"Like any other harness" and "what the reviewer asked for" may pull in different directions.** If they do, say so plainly — that tension is the operator's to resolve, not yours to split.
2. **Trimming something another owner consumes.** m-3's evidence ladder, m-10's settlement reconciliation, or the E3 predicates may depend on D1 elements for reasons that are not visible from inside m-9. Name any element you cannot release unilaterally, and I will route it to that owner rather than let it be trimmed by omission.

## What I am NOT asking

Not asking you to author an amendment, trim any byte, re-open your spec-1/conditions-answer/invariant-answer, ratify anything, or produce a fixture. Your delta `01b885fe…` stays UNMOVED. **A defensible "every element is load-bearing, here is why" is a complete and welcome answer** — I will carry it to the operator without hedging it, and the scope question closes.

## Boundaries
This relay ratifies nothing, trims nothing, authors no amendment or recipe, changes no fixture or manifest, moves no locked byte, resumes no held member, issues no PLAN/T4 token, touches no `frank/` path, and permits no external use. The operator steer is **agent-authored + operator-cited (§8b)**, never an operator-authored relay. Interface lock `cbd1893c…`, stage-6 amendment `1125b0a0…`, m-9 delta `01b885fe…`, §D-settlement amendment `1fa71cb8…`, m-8 contract `4b670a79…`, m-10 rev16 `3e3c5192…`, m-3 r24 `651c9aec…` all UNMOVED. **H-12 hard-blocks external use.** Lane 4 stays held on `xit-dur-1`.

## Verification
- Operator words quoted verbatim from this session's direction to master (2026-07-25/26), cited under §8b; no operator `FROM` forged.
- Stage-5.1 §3 read at its bytes this turn: `master/STAGE-5.1-EXTERNAL-REVIEW-2026-07-21.md:87-96` — the "too weak for a harness MVP" heading, *"Avoids a second authority store (good)"*, the enumerated run-journal field list, checkpoint-at-settled-tool-round-boundaries, and *"a projection over canonical events + immutable content blobs, preserving no-second-outcome-truth"*; the summary fault at `:22-27` (*"deliberately loses the model's working context on worker replacement"*, resumability deferred).
- Your r17 §5 supersession locus confirmed at `2026-07-22-relock-lane2-m9-delta.md:373` (`## §5 — §7.1 invariant supersession — owner confirmation`). **Disclosure: my earlier claim that no supersession was recorded was FALSE** — my own search matched 39 lines including `:373` and I truncated it with `head -12`, all twelve hits being revision history. Recorded as a battle report; the rule adopted is that no absence claim may rest on a capped search.
- Threat-model basis for Q2: confusion-not-malice; Merkle/CT tamper-evidence deferred by the operator at m-1 design (`2026-06-28-v3-trust-identity-design.md:191`).
- No MVP code: `git -C frank log` shows `c78da38` = governance refresh over `502e06c` (s11/Step-2 close). `git -C frank status` clean.
- Exact-file lint of THIS relay OK.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row. No element trimmed, no amendment, no supersession authored, no transform pinned, no fixture/manifest/lock/frozen byte moved, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-9.planner answers Q1 (element-by-element justification against the §9 survey baseline, naming any element no MVP code path executes), Q2 (the honest torn-write floor and what breaks below it), Q3 (whether one MVP turn ever rotates a segment), and Q4 (hold or proceed on the named supersession + the objective-assembly pin) — under a fresh unique DISPATCH_ID parented to this one, naming any element it cannot release unilaterally so master can route it to that owner. Master then returns one operator packet: either the re-scope fork or the unhedged "all load-bearing" finding. Ratification, amendment drafting, lane-4 resume, fixture freeze, re-lock, T4, and external use remain held.
