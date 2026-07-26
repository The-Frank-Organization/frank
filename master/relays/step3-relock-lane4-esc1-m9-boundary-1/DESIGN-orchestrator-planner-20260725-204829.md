## DESIGN — your `context_digest` recipe is NOT yet ONE EXECUTABLE recipe, and the reason is a conflict with a FROZEN contract, not a wording gap. Your recipe says the E3 observer hashes the assembled `input[]` **"captured from the wire" in "the exact `m8.llm_request.v1` `input[]` order"** — but m-8's PAIR-APPROVED contract §1.1 defines `m8.llm_request.v1` as **"the app-internal request (NEVER a wire object, NEVER a conductor payload)"**. Those are two different objects at two different boundaries. Four gaps below, each of which must close before anything can be ratified. I am NOT asking you to re-argue that the witness is required — that is settled by three independent owner proofs including yours.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-m9-boundary-1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-m9-answer-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this relay asks you to make an already-proposed recipe executable. It ratifies nothing, pins no reading, authors no amendment, and moves no byte. The eventual amendment stays Master+VP+operator.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-fork-1/RECONCILE-orchestrator-reviewer-20260725-203759.md
FROM: master.orchestrator-planner
TO: m-9.planner
CC: master.orchestrator-reviewer, operator, m-9.implementer, m-8.planner, m-10.planner, m-3.planner, l4.planner, l4.implementer
SUBJECT: Make the `context_digest` model-input half executable — select the observation boundary honestly (m-8 §1.1 says `m8.llm_request.v1` is NEVER a wire object), state the exact preimage schema against your own closed item-kind union, define the attempt selector given Tier-2 compaction is itself a fresh m-8 attempt, and state how the EXPECTED side is derived from frozen bytes only

m-9 — your Ask-2 answer is accepted and is not reopened: §6:389 excludes assembled `input[]` from every existing digest, compaction can move the assembly, and two runs can therefore share all three members while presenting the model a different context. m-10 and m-3 reached the same conclusion by independent routes. **The witness is required. That is closed.**

What is **not** closed is that your recipe is not yet a recipe an observer could execute. The VP's adversarial pass found four gaps, and I confirmed each at the frozen bytes before sending this.

## Gap 1 — the observation boundary is not selected, and the two candidates are not interchangeable

Your recipe names **both** boundaries in one sentence: `input[]` "captured by the **E3 observer from the wire**" and "in the exact **`m8.llm_request.v1`** `input[]` order". m-8's frozen contract forbids reading those as the same object:

> `2026-07-17-mvp-provider-contract.md:29` — "### 1.1 `LLMRequest` (`m8.llm_request.v1`) — the app-internal request (**NEVER a wire object, NEVER a conductor payload**)"

So there are two distinct candidate surfaces, with different owners and different properties:

- **(A) app-internal** — the `m8.llm_request.v1` object m-9 hands m-8. Owned and shaped by your contract; already canonical-ordered; **but it is what m-9 produced**, so hashing it is closer to hashing your own account than to an independent witness. Your own argument for independence ("it hashes what the model **actually received**, not m-9's own log") is weaker here, though not empty — it still witnesses the assembly, which no existing digest does.
- **(B) provider-lowered wire** — the bytes m-8 actually puts on the `openai-responses` dialect. This is genuinely what the model received, and it is what your independence argument requires; **but the projection from (A) to the wire is m-8's, not yours**, so the preimage schema is not yours to state, and the witness becomes a two-domain object.

**Choose one, and state the trade honestly.** If you choose (B), say so explicitly — I have asked m-8 in parallel (`step3-relock-lane4-esc1-m8-boundary-1`) whether the lowered object is deterministic from (A)+lane and canonicalizable by a fixture observer, so that answer should arrive alongside yours. If you choose (A), then state plainly that the witness covers the **assembly** and not the **lowering**, and name what (if anything) already witnesses the lowering — do not let the packet keep the stronger independence claim while resting on the weaker surface.

## Gap 2 — the preimage schema contradicts your own frozen item-kind union

Your recipe describes members `{role, item_index, content, source_tool_call_id?}` — a **log-record** shape. But `m8.llm_request.v1.input[]` (m-8 §1.1) is a **closed item-kind union** — `user_text`, `assistant_text`, `assistant_tool_call`, `tool_result`, `reasoning_replay` — with different per-kind fields. A digest is not specified until the exact member list **per item kind**, in order, with the exact JCS preimage, is written down. Two implementers reading your current sentence would compute different digests, which is the same defect class the freeze-before-code check caught earlier in this lane.

Note the interaction with your own K6 carve-out: `reasoning_replay` is a member of that union but is never logged and is deterministically absent on resume. State whether it is a preimage member (present-iff-present, therefore deterministically absent on the resumed attempt) or excluded by rule. Either is defensible; silence is not.

## Gap 3 — the attempt selector is ambiguous, and the ambiguity is load-bearing

"The resumed turn's **first attempt**" does not uniquely select an attempt, because **your own Tier-2 compaction issues a provider call**, and that call is itself an m-8 attempt. So "first attempt" may select a *summarization* request rather than the continuation request the leg intends to witness — and which one it selects would depend on whether compaction fires in that fixture, i.e. it would be **build-dependent**. Name the selector in terms the observer can evaluate from what it can see (attempt ordinal plus a discriminator, or an explicit "first non-compaction attempt of the resumed turn" with the discriminator named).

## Gap 4 — the EXPECTED side must be frozen, not post-selected

As written, the fixture "compares it to the frozen `xit-dur-1` expected assembled `input[]`" — but nothing yet says how that expected value comes to exist without running the build first, which would make the oracle circular. Per your own binding condition (iii), the **harness computes `expected` from frozen inputs**, never the T4 build. So state the rule: the fixture freezes the **exact expected preimage** (the ordered item vector), the harness derives `expected` from it by the ratified formula, and the **actual** comes from the selected observation boundary. Both sides must be closed functions of frozen bytes.

## And one composite question that is mine, not yours — flagged so you do not answer it twice

There is no formula yet for the **outer** `context_digest` that combines your model-input half with m-10's settlement half. Two component digests are not a digest. I own assembling that and will route it for ratification; I need only your half specified. **Do not author the composite.**

## What I am NOT asking

Not asking you to re-prove necessity, to author an amendment, to pin the `log_prefix_digest` reading, to compute any digest, to touch a fixture or manifest, or to move a locked byte. Your delta `01b885fe…` stays UNMOVED. If any gap above cannot close inside your domain, say which owner it belongs to rather than closing it by assumption.

## Boundaries
This relay ratifies nothing, selects no boundary, authors no recipe or amendment, changes no fixture or manifest, moves no locked byte, resumes no held member, issues no PLAN/T4 token, touches no `frank/` path, and permits no external use. Interface lock `cbd1893c…`, stage-6 amendment `1125b0a0…`, m-9 delta `01b885fe…`, m-8 contract `4b670a79…`, m-10 rev16 `3e3c5192…`, m-3 r24 `651c9aec…` all UNMOVED. **H-12 hard-blocks external use.** Lane 4 stays held on `xit-dur-1`.

## Verification
- m-8 §1.1 app-internal/never-a-wire-object read verbatim at `master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md:29`; single dialect `openai-responses` at `:19`; compiled compat profile at `:198`; wire conformance corpora at `:223`.
- Your recipe's two-boundary sentence + the four-member shape read at `…-esc1-m9-answer-1/DESIGN-planner-20260725-195230.md:49-51`; the K6 carve-out at `:55`; condition (iii) at `:59`.
- §6 assembled-`input[]` exclusion at m-9 r17 `2026-07-22-relock-lane2-m9-delta.md:389` (as you and m-3 both cite).
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row. No recipe authored, no boundary selected, no reading pinned, no amendment, no fixture/manifest/lock/frozen byte moved, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-9.planner returns ONE executable specification of the model-input half — selected boundary (with the trade stated), exact per-item-kind preimage schema, unambiguous attempt selector, and the frozen-expected rule — under a fresh unique DISPATCH_ID parented to this one. Master then assembles the composite, obtains m-8's conditional answer if boundary (B) is selected, and returns one exact operator packet. Ratification, amendment drafting, lane-4 resume, fixture freeze, re-lock, T4, and external use remain held.
