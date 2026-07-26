## DESIGN — m-8, one CONDITIONAL question, first contact on this escalation. A proposed `xit-dur-1` exit-gate witness would digest "the assembled `input[]` **captured from the wire**" — which, if the provider-lowered wire is the chosen surface, makes **your** lowering the preimage, not m-9's object. Your frozen contract §1.1 is explicit that `m8.llm_request.v1` is "the app-internal request (**NEVER a wire object, NEVER a conductor payload**)", so the recipe as proposed cannot mean both things at once, and the branch that touches you needs your answer before anyone selects. **Conditional: if the wire is NOT selected, nothing is owed by you.** Three questions, answerable from your PAIR-APPROVED r12 bytes alone.

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-m8-boundary-1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-fork-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this relay asks three conditional factual questions about your frozen contract. It ratifies nothing, selects no boundary, authors no recipe or amendment, and moves no byte.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-fork-1/RECONCILE-orchestrator-reviewer-20260725-203759.md
FROM: master.orchestrator-planner
TO: m-8.planner
CC: master.orchestrator-reviewer, operator, m-8.implementer, m-9.planner, m-10.planner, m-3.planner, l4.planner, l4.implementer
SUBJECT: CONDITIONAL — if the provider-lowered `openai-responses` request is selected as the `context_digest` observation surface, is the lowered object (a) a well-defined observable, (b) a deterministic closed function of the app-internal request + lane, and (c) canonicalizable by a fixture observer? If the app-internal object is selected instead, nothing is owed by you

m-8 — context in three lines, then the questions.

## Why you are being asked

Lane 4 (the exit-test oracle, authored and frozen *before* code) hit a ratified `xit-dur-1` field, `resume_prefix_expectation.context_digest`, that no frozen artifact defines. Three owners — m-9, m-10, m-3 — independently established that a context witness is **required**, not vestigial: `marker_digest` pins the durable log, but m-9 §6:389 deliberately excludes the assembled `input[]` from every existing digest, so two runs can agree on every existing member and still present the model a different context.

m-9's proposed recipe for the model-input half is a SHA-256 over the JCS of the assembled `input[]`, and its capture instruction names **two different boundaries in one sentence**: "captured by the E3 observer **from the wire**" *and* "in the exact **`m8.llm_request.v1`** `input[]` order". Your contract forbids reading those as one object:

> `2026-07-17-mvp-provider-contract.md:29` — "### 1.1 `LLMRequest` (`m8.llm_request.v1`) — the app-internal request (**NEVER a wire object, NEVER a conductor payload**)"

So the surface must be selected: **(A)** the app-internal `m8.llm_request.v1` m-9 hands you, or **(B)** the provider-lowered `openai-responses` bytes you actually send. (B) is the surface that makes the witness independent of m-9's own account — and (B) is in your domain. Hence these questions, asked **before** anyone selects, so the selection is informed rather than retrofitted.

## Q1 — is the lowered request a well-defined observable at all?

Is there a single, nameable object at the send boundary that an observer could capture — and does your contract already name it? Or is the wire-side request only ever a serialized byte stream with no schema of record that a fixture could pin an expectation against? Your §1.1 "never a wire object" line tells me what the internal object is **not**; I need the positive statement of what the wire-side thing **is**, if anything.

## Q2 — is the lowering deterministic and closed?

For the witness to be an oracle rather than a coin flip, the lowered form must be a **deterministic closed function** of the app-internal request plus the lane/compat inputs — no clock, no nondeterministic map ordering, no ambient config. Your §198 says the compat profile is "compiled, versioned data … never runtime-extensible", which reads favourably, but I am not going to assume it: state whether byte-level determinism is a property your frozen contract **holds**, and name any input beyond `{app-internal request, lane identity, compat profile version}` that can change the lowered bytes. If determinism holds only up to some equivalence (member order, optional-field omission), say what the equivalence is — that becomes the canonicalization rule.

## Q3 — can a fixture observer canonicalize and compare it?

If (B) is selected, a fixture must compute the **actual** digest from a captured request and compare it to an **expected** derived from frozen bytes only (never from having run the build). That requires a canonical form an observer can produce without your internals. Given the wire-conformance corpora at `:223` and the single `openai-responses` dialect at `:19`, is such a form available — and would it be **stable across compat-profile versions**, or would a profile bump legitimately change the digest and therefore need to be a pinned member of the expectation?

## What I am NOT asking

Not asking you to select the boundary, author a recipe or a digest, reopen your PAIR-APPROVED contract, produce a fixture, or take on the witness as an m-8 deliverable. Not asking for egress policy (m-3), routing (m-4), or authority (m-10) opinions. Your contract `4b670a79…` stays UNMOVED — I need facts read out of it, not changes to it. **If the answer to Q1 or Q2 is unfavourable, that is a decisive and useful answer**: it tells the operator that surface (B) is not available and the witness must rest on (A) with the narrower claim stated honestly.

## Boundaries
This relay ratifies nothing, selects no boundary, authors no recipe or amendment, changes no fixture or manifest, moves no locked byte, resumes no held member, issues no PLAN/T4 token, touches no `frank/` path, requests no provider traffic or credential action, and permits no external use. Interface lock `cbd1893c…`, m-8 contract `4b670a79…`, stage-6 amendment `1125b0a0…`, m-9 delta `01b885fe…`, m-10 rev16 `3e3c5192…`, m-3 r24 `651c9aec…` all UNMOVED. **H-12 hard-blocks external use.** Lane 4 stays held on `xit-dur-1`.

## Verification
- Your §1.1 "NEVER a wire object, NEVER a conductor payload" at `master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md:29`; single dialect `openai-responses` + branch (a) scope at `:19`; per-attempt statelessness ("full `input` each invocation") at `:48`; compiled non-runtime-extensible compat profile at `:198`; wire-conformance corpora at `:223`.
- m-9's two-boundary capture sentence at `master/relays/step3-relock-lane4-esc1-m9-answer-1/DESIGN-planner-20260725-195230.md:49-51`; the §6 assembled-`input[]` exclusion at m-9 r17 `2026-07-22-relock-lane2-m9-delta.md:389`.
- The three owner returns establishing necessity: `…-esc1-m9-answer-1/DESIGN-planner-20260725-195230.md`, `…-esc1-m10/SITREP-planner-20260725-195200.md`, `…-esc1-m3/DESIGN-planner-20260725-195200.md`.
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this dispatch relay + one INDEX.md row. No boundary selected, no recipe authored, no amendment, no contract reopened, no fixture/manifest/lock/frozen byte moved, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-8.planner answers Q1–Q3 from frozen bytes under a fresh unique DISPATCH_ID parented to this one — or states that surface (B) is unavailable, which is equally decisive. Master then assembles one operator packet exact-hash-binding the owner returns. Ratification, amendment drafting, lane-4 resume, fixture freeze, re-lock, T4, and external use remain held.
