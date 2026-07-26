## DESIGN — esc2 DISPOSITION: **neither A nor B as framed — your reasoning for A is right, but the action is B's, for a reason you did not have.** VP r1 on the esc1 packet (`…-esc1-amend/RECONCILE-orchestrator-reviewer-20260725-190320.md`, filed AFTER your esc2) makes m-9's condition (ii) **binding**: the `xit-dur-1` fixture must pin the **exact interval and the exact input bytes** the recipe ranges over. Your input pins neither — it describes the shape generically ("whose interval is `[first_seq,last_seq]`") with no concrete values and no ordered `{seq, record_digest}` vector. So a fresh `xit-dur-1` input generation is **forced regardless of the clause**. Materializing it now would buy **two** propose/materialize/equality cycles where holding buys **one**. **RULING: the eleven materialize now; `xit-dur-1` is held and returns in one corrected generation that pins the interval AND retires the clause.** Your known-provisional request is **granted**.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-relock-lane4-l4-esc2-disp
PARENT_DISPATCH_ID: step3-relock-lane4-l4-esc2-req
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the operator hand-relays this. The esc1 members stay held; freeze/re-lock remain Master+VP; T4 stays behind the re-lock + H-16/H-26.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-l4-esc2-req/SITREP-planner-20260725-185817.md
FROM: master.orchestrator-planner
TO: l4.planner
CC: l4.implementer, operator, master.orchestrator-reviewer, m-9.planner, m-3.planner, m-10.planner
SUBJECT: esc2 — eleven artifacts materialize; `xit-dur-1` held for one corrected generation (interval pinning forces it, not the clause); known-provisional record granted

## Your reasoning was right; the facts moved under it

You argued that a rationale clause with no gate consequence does not justify a full hand-relay generation, and weighed it against the standing cost rule that chose file relays at B22. **That reasoning is correct and I am not overruling it.** If the clause were the only consideration, A would be the right answer.

What changed is **new information that post-dates your relay**: the VP's review of my esc1 escalation landed at `…-190320` and made **all three** of m-9's soundness conditions binding, not just the one I had adopted. Condition (ii) requires that the `xit-dur-1` fixture pin **the exact interval and the exact frozen input bytes** the digest recipe ranges over — otherwise, in the VP's words, "the expected digest remains selectable after the code exists."

**Checked against your filed bytes, not assumed:** `first_seq`, `last_seq`, `record_digest` and `marker_digest` each appear exactly **once** in the whole proposal, together in one descriptive sentence — *"each round closes with an fsync-durable round_marker whose interval is `[first_seq,last_seq]` and whose marker_digest covers the ordered `{seq, record_digest}` array."* That describes the **schema**; it pins **no concrete interval** and supplies **no ordered vector**. Condition (ii) is therefore unsatisfied today, and satisfying it is a change to the fixture input — your artifact, not the manifest.

So `xit-dur-1` needs another generation **either way**. Given that, materializing it now costs two cycles instead of one, which is the same cost argument you made, pointing the other way.

## Ruling

**Eleven artifacts materialize now.** Master proceeds to `…-l4-materialize-1` over the eleven unaffected envelopes, addressed to `l4.implementer` for its equality confirmation. Nothing about those eleven is in question.

**`xit-dur-1` is HELD** — not re-proposed now, not materialized now. When the esc1 disposition returns (reading ratified + the `context_digest` fork settled), you file **one** corrected `xit-dur-1` generation that:
1. pins the exact interval and the exact input bytes/ordered `{seq, record_digest}` vector per condition (ii);
2. retires the provisional rationale clause;
3. carries whatever the ratified reading requires.

One generation, three corrections, no wasted cycle.

## Your known-provisional request — granted, and on the record

**This relay and your esc2 stand as the record that the clause is known-provisional.** Its later removal is a **folded correction**, not a fresh finding against you, and `l4.implementer` should read it that way if it encounters the clause in any interim artifact. You found and named a defect in your own filed bytes before anyone else saw it, having gone back through them specifically because you had been corrected once — that is the behaviour this protocol is trying to produce, and it should not cost you a finding.

## Two things worth saying about the sweep itself

Your sweep result is the more valuable half of esc2: **`context_digest` appears zero times across all twelve artifacts**, and the `marker_digest` = `log_prefix_digest` identity error **never reached artifact bytes** — it lived only in the escalation prose I had already corrected. That means the inference did **not** propagate into the things that get frozen, which is exactly what you needed to check and exactly the right instinct after a correction: **verify whether the error reached the artifacts, not just the argument.**

Your hold-breach analysis is also correct as stated: naming a member and asserting a property of it is not authoring a literal, a derivation, a slot, or a placeholder. I checked the clause against the hold's own terms rather than taking your word, and I agree.

## What remains held

The two esc1 members (`log_prefix_digest`, `context_digest`) stay held in every form. Proposal 2 (the manifest) stays held. You rejoin only via `step3-relock-lane4-l4-esc1-resume`. The esc1 question is now with the operator and the VP, and has grown: the VP established that `context_digest` is **not** provenance-free residue but an **unfulfilled requirement** from VP r7 (exact resumed-context identity), so my earlier "lean toward striking it" is withdrawn pending addressed returns from m-9, m-10 and m-3.

## Boundaries
This disposition materializes nothing by itself, authors no artifact, resumes no held member, requests no owner-fidelity, freezes nothing, moves no locked byte, issues no PLAN/T4 token, touches no `frank/` path, and permits no external use. Kickoff `4f3928ec…`, plan rev13 `60daac08…`, interface lock `cbd1893c…`, amendment `1125b0a0…` UNMOVED. **H-12 hard-blocks external use.**

## Verification
Occurrence counts in `…-l4-propose-1/DESIGN-planner-20260725-183149.md` at this relay's stamp: `first_seq` 1, `last_seq` 1, `record_digest` 1, `marker_digest` 1 — all four inside a single descriptive sentence, with no concrete interval values and no ordered vector; `log_prefix_digest` 1, in the clause you quoted. m-9's three conditions are at `…-esc1-m9/DESIGN-planner-20260725-185400.md:43-47`; the VP's binding of all three is at `…-esc1-amend/RECONCILE-orchestrator-reviewer-20260725-190320.md` (F3). Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this disposition relay + one INDEX.md row. No materialization performed yet, no artifact authored or altered, no held member resumed, no lock/owner/frozen byte moved, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: master materializes the eleven and files `…-l4-materialize-1` to `l4.implementer`; `l4.planner` holds `xit-dur-1` and the two esc1 members, and rejoins only via `step3-relock-lane4-l4-esc1-resume`.
