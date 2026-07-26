## RECONCILE — REQUIREMENT-COMPLETE FORK for the operator. All three addressed owners returned, and **all three independently answer: an independent context witness IS REQUIRED.** Not one of them could be reached by lexical search — each supplies a distinct mechanical proof: **m-9** (§6 explicitly excludes assembled `input[]` from every existing digest; compaction can move the assembly), **m-10** (an airtight receipt-presence divergence — identical log prefix, different settlement class), **m-3** (without it the leg cannot distinguish "resumed from the exact prefix" from "re-derived a matching prefix", which is the exact case §7 says must FAIL). **My "strike it as vestigial" lean is withdrawn in full; the VP's correction was right and my premise was unsafe.** Both halves of the witness already have observer-executable mechanisms, so option (ii) is a **binding-and-scope decision, not new invention** — and m-10's one open scope question is **closed by VP r7's own wording.** Two operator decisions below.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-fork-1
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-m9-answer-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — both decisions change what value is frozen into the Step-3 exit gate and touch a ratified §7 field. Operator ratification + VP exact-byte review; master proposes and may not self-ratify.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-m9-answer-1/DESIGN-planner-20260725-195230.md
FROM: master.orchestrator-planner
TO: operator
CC: master.orchestrator-reviewer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-3.planner, m-3.implementer, l4.planner, l4.implementer
SUBJECT: All three owners return "witness required" with independent proofs; recommend ratifying reading (b) + option (ii) composite witness; VP r7's wording closes the scope question

## The three returns — independent, and none reachable by search

| owner | finding | the proof it supplied |
|---|---|---|
| **m-9** (`…-esc1-m9-answer-1`) | witness **warranted** | The three members transitively pin the durable **log** (chain proof: `record_digest` covers `prev_digest`, so any divergence anywhere changes the boundary `marker_digest`). But **§6 line 389 deliberately excludes assembled `input[]`** from `logical_surface_digest`, so **no existing digest witnesses what the model receives**. Governed 3-tier compaction can produce a different assembly from the same log ⇒ **two runs can share all three members and present the model a different context.** |
| **m-10** (`…-esc1-m10`) | witness **required** | Receipt-presence divergence at rev16 §2:30: an entry is `settled_with_content` **iff** a `content_ready_receipts` row is committed — and that row lives in **m-10's private store, not the D1 log**. So it is **orthogonal to `marker_digest`**. Two runs, **byte-identical log prefix**, same round, same predecessor: one yields a trusted `provider_output`, the other `uncertain`. Materially different resumed context, same three members. |
| **m-3** (`…-esc1-m3`, evidence owner) | witness **load-bearing** | Without it the leg has **no mechanical means** to separate a true resume from a re-derivation that reconstructs a matching-looking prefix while assembling a different `input[]` — and such a run **scores PASS**. That is precisely what §7's own row forbids: *"A degraded re-derivation never satisfies `xit-dur-1`."* The leg could not enforce its own headline requirement. |

**I withdraw my earlier lean without qualification.** I proposed striking `context_digest` on the strength of a negative lexical search. Three owners reading their own frozen bytes found three independent reasons it is load-bearing. The VP was right that a negative search is not proof of redundancy, and right to correct its own prior closure of F106 rather than let mine stand on it.

## Decision 1 — `log_prefix_digest`: ratify reading (b)?

All three concur, and m-3 independently verified the property rather than accepting m-9's account: the boundary honoured `round_marker`'s `marker_digest` is **prefix-determining** (chain-transitive, fails closed on any alteration). Reading (a) freezes a value that does not identify the prefix; reading (c) requires a recipe nobody has.

**Recommendation: ratify (b)**, binding **all three** of m-9's soundness conditions — not just the one I originally adopted:
1. the reading is ratified;
2. the `xit-dur-1` fixture pins the **exact interval and the ordered `{seq, record_digest}` vector** (this forces the corrected fixture generation already ruled at esc2);
3. the **gate harness — not the T4 build — computes `expected`** from the frozen inputs using the frozen §1.3 + §1.5 recipes.

## Decision 2 — `context_digest`: the requirement-complete fork

**(i) Assembly-determinism proof.** m-3 named the one exit: if **m-9 freezes and proves**, as a gate-checkable lock property, that the assembled model-visible input is a deterministic closed function of the frozen prefix + frozen manifest/build, then context-identity follows from prefix-identity and the witness is derivable. **But m-9's own bytes argue against assuming it:** §6 explicitly excludes assembled content, and m-9 states its design does **not** pin a byte-level "assembled `input[]` == logged `input_item` records" invariant. So (i) means **authoring a new invariant and proving it** — a larger design act than it sounds, in m-9's domain, and it must be *proven* rather than asserted.

**(ii) Author the composite witness.** Per m-10, a digest proving *exact resumed-context identity* is a **composite**, and **both halves already have observer-executable mechanisms — no new capture path is owed:**
- **model-visible continuation input** (m-9): SHA-256 over the JCS of the assembled `input[]` on the resumed turn's **first attempt**, captured by the **E3 observer from the wire** — independent precisely because it hashes what the model actually received, not m-9's own log, so an assembly bug cannot self-cancel;
- **settlement snapshot** (m-10): `SHA-256(JCS settlement_manifest bytes)`, computable by a fixture observer **from the `turn_open` wire alone** (payload-free, no m-10 internals).

**m-10's open scope question is closed by the requirement's own wording.** m-10 asked whether epoch/lease durable state is a third member. **VP r7 required "exact resumed-context identity incl. the model-visible continuation input and the settlement snapshot"** — it enumerates exactly **two** components, and epoch/lease is durable state that is not among them. So the composite is **two members**, settled by the original requirement rather than by preference.

**Recommendation: (ii).** It discharges VP r7's requirement as written, needs no new mechanism, and each half is owned by the seat that already produces it. (i) trades a bounded binding decision for an unbounded proof obligation in m-9's domain, and if the proof fails we return here having spent the time.

## How it lands without breaking the lock
`STEP-3-STAGE6-AMENDMENT.md` is **interface-lock constituent row 54** at `1125b0a0…`; editing it voids Item A and the 38-file lock. So: an **additive supersession record** (MVP-amendment precedent) naming the precise superseded fragment, preserving every other stage-6 byte and semantic, carrying its own full SHA-256 through **VP exact-byte review then operator ratification** — and the lane-4 re-lock **exact-hash-binds three inputs**: the interface lock, this additive amendment, and the frozen manifest. Per the VP, a prose citation is too weak: without hash binding, later amendment drift would not void the re-lock.

## What this costs, stated plainly
Lane 4 stays held on these two members until ratification. The corrected `xit-dur-1` generation then folds: the pinned interval + record vector (condition ii), the ratified `log_prefix_digest` derivation, the ratified `context_digest` recipe, and the retirement of the known-provisional clause — **one generation, not four**. The other ten records and eleven materialized artifacts are unaffected and already equality-confirmed.

## What I need from you
1. **Ratify reading (b)** for `log_prefix_digest`, with all three m-9 conditions binding?
2. **Choose (i) or (ii)** for `context_digest` — master and all three owners point to **(ii)**.

On your decision I draft the additive record, route its exact bytes to the VP, return it for your ratification, and only then issue `…-l4-esc1-resume` to the pair.

## Boundaries
This relay ratifies nothing, pins no reading, authors no recipe or amendment, changes no fixture or manifest, moves no locked byte, resumes no held member, requests no owner-fidelity, issues no PLAN/T4 token, touches no `frank/` path, and permits no external use. Interface lock `cbd1893c…`, stage-6 amendment `1125b0a0…`, m-9 delta `01b885fe…`, m-10 rev16 `3e3c5192…`, m-3 r24 `651c9aec…` all UNMOVED. **H-12 hard-blocks external use.**

## Verification
Three owner returns read at their bytes: `…-esc1-m9-answer-1/DESIGN-planner-20260725-195230.md` (re-filed under the fresh unique id as requested — the F1 lineage repair is discharged and this packet parents to the **answer**, not the question), `…-esc1-m10/SITREP-planner-20260725-195200.md`, `…-esc1-m3/DESIGN-planner-20260725-195200.md`. VP r7's requirement wording verified verbatim at `step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-073500.md:62`. §7 Durability row + the "degraded re-derivation never satisfies `xit-dur-1`" clause at `STEP-3-STAGE6-AMENDMENT.md:371`; the field vector at `:383`. Amendment confirmed lock row 54. Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this fork relay + one INDEX.md row. No amendment drafted or ratified, no reading pinned, no recipe authored, no fixture/manifest/lock/owner/frozen byte moved, no held member resumed, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the **operator** ratifies reading (b) and selects (i) or (ii); the **VP** reviews this fork. On the decision master drafts the additive supersession record for VP exact-byte review and operator ratification, then issues `step3-relock-lane4-l4-esc1-resume`.
