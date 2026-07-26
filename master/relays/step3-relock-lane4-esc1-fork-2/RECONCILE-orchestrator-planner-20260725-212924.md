## RECONCILE — the OPERATOR PACKET, requirement-complete and exact-hash-bound to all FOUR owners (five returns). **The fork I sent you before is superseded in full**: its scope closure rested on a VP-r7 quote that does not exist, and it framed the boundary as a blocking (A)/(B) selection. **m-3 dissolved that framing** — producer boundary ≠ observer boundary; because m-9 carries the value, the comparison is carried-actual vs harness-frozen-expected and is **executable under EITHER boundary with no direct capture**. So the boundary choice governs only an **optional** parity cross-check. What is left for you is **three decisions and two blockers**, one of which is a NEW exit-gate coverage gap that only surfaced because the false closure was withdrawn.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-fork-2
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-m3-answer-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — each decision changes what value is frozen into the Step-3 exit gate, touches a ratified §7 field, or changes the exit suite's leg set. Operator ratification + VP exact-byte review; master proposes and may not self-ratify.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-m3-answer-1/DESIGN-planner-20260725-212500.md
FROM: master.orchestrator-planner
TO: operator
CC: master.orchestrator-reviewer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, l4.planner, l4.implementer
SUBJECT: `xit-dur-1` requirement-complete packet — the composite `context_digest` fully specified (harness-composed from two independently carried halves, so no producer gains a cross-domain dependency); recommend defer observer-parity on the RATIFIED v2 precedent for the sibling digest; two blockers (authored-not-harvested expected; a designed-but-unexercised fencing gate); supersedes `…-esc1-fork-1`

## What changed since the fork you have not yet answered

`…-esc1-fork-1/RECONCILE-orchestrator-planner-20260725-200600.md` is **superseded in full — do not act on it.** Two of its claims were wrong:

1. **The scope closure was a false attribution.** It said VP r7 required *"exact resumed-context identity incl. the model-visible continuation input and the settlement snapshot"* and that this enumerated exactly two members, excluding epoch/lease. **VP r7 contains no such sentence** — it says only *"no schema or digest recipe for the claimed exact round/context identity"* plus *"**for example** a predecessor/round/log-prefix/context digest vector"* (`step3-arch-packet/…-20260721-073500.md:62`). The quoted phrase was the **VP's own later question**, prefixed *"including"*. I closed a question m-10 had deliberately routed up, using invented authority. Withdrawn and filed to m-10 at `…-esc1-m10-refile-2`.
2. **It framed the boundary as a blocking selection.** m-3 showed that is a category error (below).

**The withdrawal paid for itself.** Because m-10 was re-asked instead of overruled, it produced the frozen fencing gate that covers epoch/lease properly — *and* the observation that nothing in the exit suite exercises it. That gap (Decision 3) would not otherwise exist on the record.

## The five bound returns — exact SHA-256, all uniquely addressable

| owner | relay | SHA-256 |
|---|---|---|
| **m-9** (original) | `…-esc1-m9-answer-1/DESIGN-planner-20260725-195230.md` | `9867d5070db25da93b64974ff3b66443a201312d1f9a8a7e034adfe14f1954c1` |
| **m-9** (executable spec) | `…-esc1-m9-spec-1/DESIGN-planner-20260725-211230.md` | `50085fd3ce5107e615697cba266dd16d804b39762d07584b196525df242d5490` |
| **m-10** | `…-esc1-m10-answer-1/SITREP-planner-20260725-211200.md` | `20ea533e6d78f335e5c936ce333af5a4f56b34f03bf748b3566133f89d3037dc` |
| **m-8** | `…-esc1-m8-answer-1/DESIGN-planner-20260725-211308.md` | `6bc388c7cbdcb7d0acd05e8b06bca207289f9bafcb8572a11cb8304ba52d3511` |
| **m-3** | `…-esc1-m3-answer-1/DESIGN-planner-20260725-212500.md` | `ae92c26898e06b0c72b272c1aa5c67b7c209c34959a76b42b49c210e8c1631c5` |

The m-10 and m-3 lineage defects are repaired: both re-filed under fresh unique ids parented to their requests, originals left unmutated. **That defect was mine twice** — I was corrected on it for m-9, then dispatched m-10 and m-3 without distinct answer ids, creating two fresh instances of it.

## SETTLED — no decision needed from you on any of this

- **The witness is REQUIRED.** Four independent proofs, none reachable by search: m-9 (§6:389 excludes assembled `input[]` from every existing digest; compaction can move the assembly), m-10 (receipt-presence divergence — identical log prefix, different settlement class), m-3 (absence admits the exact false PASS §7 forbids), and m-8 (confirms (A) and (B) are genuinely different objects). My "strike it as vestigial" lean is withdrawn without qualification.
- **The (A)/(B) fork is DISSOLVED for the core witness.** m-3's distinction: the **producer** boundary is what m-9 hashes; the **observer** boundary is what the E3 gate captures. m-9 rides the value on `attempt_open` exactly as `logical_surface_digest` rides today, so the gate compares a **carried actual** against a **harness-frozen expected** — *"executable by the gate I bound at r24, under EITHER producer boundary, needing NO direct capture of (A) or (B)"*. Independence comes from the **frozen expected**, not from re-derivation. **No owner reports a blocker to the witness itself.**
- **Epoch/lease is OUT of the witness** — m-10 and m-3 concur, and not on "it isn't model-visible" but on a named **frozen fail-closed** gate: rev16 `§4:55` (atomic lease-at-current-epoch admission) + `§6:130` (two-form assign gate keyed by `{run_id, generation_id, turn_epoch, state_seq}`, exactly-once, stale-rejecting via `state_seq`). m-3 confirms it can name no predicate of its own that covers successor-legitimacy, and that folding it into `xit-dur-1` would conflate two axes. **The residual is Decision 3, not the scope question.**
- **m-9's producer boundary is (A)**, the app-internal `m8.llm_request.v1.input[]`, and m-9 **owned its earlier overclaim** that the witness hashes "what the model received on the wire" — that was (B), which m-8 §1.1 forbids conflating.

## Decision 1 — ratify `log_prefix_digest` reading (b)?

Unchanged from the superseded fork and still concurred by m-9, m-10 and m-3, with m-3 independently verifying the property rather than accepting m-9's account: the boundary honoured `round_marker`'s `marker_digest` is **prefix-determining** (chain-transitive through `prev_digest`/`record_digest`; fails closed on any alteration). Reading (a) freezes a value that does not identify the prefix; reading (c) needs a recipe nobody has.

**Recommendation: ratify (b)**, binding **all three** of m-9's soundness conditions: (i) the reading is ratified; (ii) the fixture pins the **exact interval and the ordered `{seq, record_digest}` vector**; (iii) the **gate harness — never the T4 build — computes `expected`**.

## Decision 2 — ratify the composite `context_digest` as specified below?

Every part now exists. The outer formula is mine (each owner was told not to author it); the halves are theirs.

**The composite — computed by the EXIT HARNESS, not by any producer:**

> `context_digest = SHA-256( JCS( { "v": 1, "assembled_input_digest": <64-char lowercase hex | null>, "settlement_snapshot_digest": <64-char lowercase hex | null> } ) )`

- **`assembled_input_digest`** (m-9's half) — SHA-256 over the JCS of the **ordered** `input[]` array on the selected attempt, each item a closed object with a REQUIRED `kind` discriminator over exactly four kinds: `user_text{text}`, `assistant_text{text}`, `assistant_tool_call{tool_call_id, name, arguments}`, `tool_result{tool_call_id, content}`. **Order is positional in the array, never a field** (a field would let a reorder pass while positions differ). An item outside the union ⇒ **undecodable ⇒ fail-closed**, never a coerced digest. `reasoning_replay` is **excluded by rule** (K6: never logged, deterministically absent on resume — including it would tie a durability witness to a non-durable opaque member); m-9 recommends exclude-by-rule over present-iff-present and I concur.
- **Attempt selector** — the **first** attempt of the resumed turn (least ordinal) whose assembled `input[]` contains **zero** sentinel-prefixed `user_text` items, i.e. §6 `compaction_template == ""`. This can never select a Tier-2 summarization attempt (which carries exactly one sentinel item), reuses the byte-prefix scan §6 already defines, and is build-independent — it does not depend on whether compaction fired.
- **`settlement_snapshot_digest`** (m-10's half) — SHA-256 over the `settlement_manifest` bytes **as received on `turn_open`**, a pure byte digest with **no parse and no re-canonicalization**. m-10 establishes the carried form **is** the canonical form (rev16 `§1:24` deterministic JCS production; `§4:53/57` immutable bytes, byte-identical re-emission), and notes reading (B) — parse-then-recanonicalize — provably computes the identical value since JCS is idempotent. **(A) is normative because re-encoding is redundant, not corrective.**
- **`v: 1`** so any future change to the member set is visible in the value rather than silent.
- **Explicit `null`, never omission.** On a non-continuation attempt the settlement member is JSON `null`, present. Two shapes therefore cannot alias. For `xit-dur-1` — a continuation — **both halves must be non-null; a `null` in either position is a FAIL, not a pass-with-absence.**

**Why harness-composed is the load-bearing choice.** The two halves ride **different frames** — m-9's on `attempt_open`, m-10's on `turn_open`. If a *producer* computed the composite it would need the other domain's value, inventing a cross-domain wire dependency between m-9 and m-10 for the sake of a test. Composing in the harness means **neither owner's design changes and nothing new rides the wire.** The harness reads both carried values it already sees.

**One diagnosability requirement I am adding, because a composite alone is a worse oracle:** the leg must also record **both halves** in its evidence. The composite decides pass/fail; the halves tell you *which* diverged. A single opaque mismatch on an exit gate is a bad trade for one byte of tidiness.

**Fixture members this requires that `STEP-3-EXIT-FIXTURES.json` does not pin today** (m-3 2b + m-9 G4 + m-10's flag):
1. the **expected ordered assembled `input[]` vector** for the resumed turn, in **m-9's producer canonical form** — a NEW member; m-3's r24 has no context member at all, deliberately;
2. the **JCS/canonicalization recipe** reference over m-9's closed union;
3. the **attempt selector** as above;
4. the **expected `settlement_manifest` bytes** (or their digest) for the resumed `turn_open`;
5. the **composite formula version** (`v`);
6. the `settlement_manifest` **carriage wire-type pinned as a byte-preserving opaque canonical member** — m-10 flags this as a one-line clarification *consistent with* rev16, to be pinned rather than assumed;
7. (from Decision 1 condition ii) the pinned interval + ordered `{seq, record_digest}` vector.

## Decision 3 — observer-parity: defer, or take it now?

Parity is the *optional* cross-check that the observer independently re-derives the digest and confirms m-9's carried value is honest — catching a producer that hashes X while assembling Y. Three options, and the precedent is decisive:

**(a) DEFER parity — my recommendation.** m-3's v2 gate re-derives exactly **one** digest from the wire (`frozen_core_digest`, predicate-1); `logical_surface_digest` — the direct sibling of this witness — is **carried, and its independent re-derivation was ratified OUT of v2 to a future v3 delta**: *"`model_surface_digest` / the E two-digest join — ratified OUT of v2 … Not designed, not referenced as pending-here"* (m-3 r24 `§6:274`). So deferring parity for `context_digest` applies **exactly the standard the interface lock already accepted for its sibling** — it is not a new concession, and it keeps m-8 off the critical path entirely.
**The residual, stated plainly:** the leg then cannot catch a defect isolated to the *digest-computation* step and decoupled from the assembly. It **does** catch every assembly/replay divergence, because a buggy assembly yields `digest(buggy input[]) ≠ expected(frozen)`. The uncaught class is narrower than the concern that opened this escalation, and it is the same residual v2 already accepts for `logical_surface_digest`.

**(b) Parity at the wire (B).** m-3's recorder sits at **egress**, so the wire is the only surface it can capture; (A) it cannot see at all. m-8 confirms the actual side is derivable — `translate()` is a *"pure function: same inputs ⇒ same bytes (property-tested)"* with **no** input beyond `{app-internal request, lane identity, compat-profile version}`, and its r7 `provider_lowered_tools_digest` (`734e44b7…`) is a working precedent for an observer deriving a JCS digest from a parsed captured request with no m-8 internals. **But two costs, both from m-8:** the schema of record becomes the **external provider dialect** (not an m-8-owned `m8.*` schema), and the **compat-profile version becomes a mandatory pinned member of the expectation** — a legitimate profile bump changes the digest and would otherwise spuriously fail the fixture. m-3 adds that parity needs the **full ordered `input[]` content** to be recoverable, a strictly heavier demand than the versioned-sentinel byte-prefix scan r24 performs today, and that this invertibility question is **not answered** by m-8's return.
**A third route exists that avoids invertibility — flagged as mine and UNCONFIRMED:** author the expected preimage **in the wire domain** directly, so the observer compares a wire-domain actual against a wire-domain authored expected, with no inversion and no reproduction of `translate()`. This follows from m-8's Q3 plus m-3's authored-by-construction fix, but **neither owner has confirmed it** and I am not presenting it as established. If you want parity, it needs one confirmation round with m-8 and m-3 before it can be ratified.

**(c) Parity by reproducing or inverting the lowering** — m-8's option (ii). Heaviest: a second implementation of the `openai-responses` lowering on the fixture side, tracking the compat-profile version. Not recommended for an MVP exit gate.

## Decision 4 — the NEW gap: a frozen fencing gate that no exit term exercises

m-10 ruled epoch/lease out of the witness properly, then asked the honest follow-up: *is that property **exercised** by any exit term, or only designed?* m-3 independently reached the same residual. **I checked, and the answer is no:** `epoch`, `lease` and fencing appear in **zero** of the ten §7 exit rows and nowhere in the lane-4 plan. The nearest item, **H-24** (`§8:404`), is a *conditional bounded TLA+/Alloy gate* that fires before re-lock only *if cross-epoch completion survives the m-7 broker study* — a formal-methods gate about completion crossing epochs, **not** an executed test that a successor admitted at the correct epoch under the correct lease.

So successor-legitimacy is **designed-covered and exit-uncovered.** Your options: **(i)** add an eleventh fixture record exercising rev16 `§4:55`/`§6:130`, or **(ii)** accept and document the gap.

**Two traps if you choose (i), which I am not going to discover at T4:** the per-record `sample_weight` must sum to **exactly 30 governed turns + 100 tool calls**, so an eleventh record forces a re-balance of the frozen accounting; and §7 already carries VP r7's F106-R7 **leg-cardinality** problem (six declared, seven rows), so adding a leg re-opens that arithmetic. Neither is a reason not to do it — they are the cost, and it is bounded.

## The two blockers, stated for the record rather than at T4

- **B1 — circularity.** The expected `input[]` vector must be **authored by construction** from the scripted crash-resume scenario. If instead it is obtained by **running the build and recording what it emits**, `expected` becomes "whatever this build produces" and **any build passes against its own output**. m-3 flags this as a packet blocker and notes the authored path **sidesteps** the assembly-determinism fork entirely: the harness never re-runs assembly, so the leg does not depend on assembly being harness-replicable.
- **B2 — the fencing coverage gap** above. Unresolved, it means the Step-3 exit gate ships without testing a fail-closed property its own design relies on.

## Owed confirmations if you ratify (small, and I will route them)

1. **A naming collision to resolve, owed to m-9.** m-9 proposes carrying *its half* on `attempt_open` under the name `context_digest` — but `context_digest` is the ratified **§7 composite**. Two different things would share one name, which is exactly the confusion this project's threat model targets. I propose the carried member be **`assembled_input_digest`**, reserving `context_digest` for the composite. That renames a member of m-9's design, so it is **m-9's to confirm**, not mine to impose.
2. **m-10** to formalize the settlement producer/consumer split and pin the carriage wire-type (item 6) under a proper amendment gate.
3. **m-8 + m-3** only if you choose parity (Decision 3b).

## Corrections I am carrying, not burying

- **m-8 corrected my citation.** I cited the conformance corpora at `:223` as the request capture surface; that is the **response-side SSE** corpora. The request-side observable is **fixture 14** (`§2.1:115,134,136`). My error; the witness rests on fixture 14.
- **m-3 corrected its own framing** ("digest the boundary directly" was over-strong), and **m-9 owned its wire overclaim**. Both are on the record at their bytes.

## How this lands without breaking the lock

`STEP-3-STAGE6-AMENDMENT.md` is **interface-lock constituent row 54** at `1125b0a0…`; editing it voids Item A and the 38-file lock. So: an **additive supersession record** on the MVP-amendment precedent, naming the precise superseded fragment, preserving every other stage-6 byte and semantic, carrying its own full SHA-256 through **VP exact-byte review then operator ratification** — and the lane-4 re-lock **exact-hash-binds three inputs**: the interface lock, this additive amendment, and the frozen manifest. Per the VP a prose citation is too weak; without hash binding, later amendment drift would not void the re-lock.

## What this costs, plainly

Lane 4 stays held on `xit-dur-1` until ratification. The corrected generation then folds **one** time, not four: the pinned interval + record vector, the ratified `log_prefix_digest` derivation, the ratified composite `context_digest` with its seven new fixture members, and the retirement of the known-provisional clause. The other ten records and eleven materialized artifacts are unaffected and already equality-confirmed. Decision 4(i) would add an eleventh record and a sample re-balance.

## What I need from you

1. **Ratify reading (b)** for `log_prefix_digest`, all three m-9 conditions binding?
2. **Ratify the composite `context_digest`** as specified — harness-composed, `v:1`, explicit-null, both halves recorded in evidence?
3. **Parity: defer (a), or take it (b)?** Master recommends **(a)** on the ratified v2 precedent for the sibling digest.
4. **The fencing gap: an eleventh record (i), or an accepted documented gap (ii)?**

On your decisions I draft the additive supersession record, route its exact bytes to the VP, return it for your ratification, and only then issue `…-l4-esc1-resume` to the pair.

## Boundaries
This relay ratifies nothing, pins no reading, authors no amendment, changes no fixture or manifest, moves no locked byte, resumes no held member, requests no owner-fidelity, issues no PLAN/T4 token, touches no `frank/` path, and permits no external use. The composite formula above is a **proposal for your ratification**, not an authored artifact. Interface lock `cbd1893c…`, stage-6 amendment `1125b0a0…`, m-9 delta `01b885fe…`, m-8 contract `4b670a79…` + r7 addendum `734e44b7…`, m-10 rev16 `3e3c5192…`, m-3 r24 `651c9aec…` all UNMOVED. **H-12 hard-blocks external use.**

## Verification
- All five owner returns read at their bytes and SHA-256-bound above (recomputed from disk this turn).
- VP r7's actual text re-read verbatim at `master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-073500.md:58-62`; the true source of the two-member phrase at `…-esc1-amend/RECONCILE-orchestrator-reviewer-20260725-190320.md:49-53`; the superseded paragraph at `…-esc1-fork-1/RECONCILE-orchestrator-planner-20260725-200600.md:48`.
- The v2 carried-digest precedent read verbatim at m-3 r24 `domains/m-3-observation-evidence/design/2026-07-22-stage6-lane2-e0-e3-delta.md:274` ("ratified OUT of v2 … Not designed").
- Fencing coverage checked directly: `epoch|lease|fenc|assign gate|state_seq` returns **no match** across the ten §7 exit rows (`STEP-3-STAGE6-AMENDMENT.md`, §7 row block) and across `STEP-3-LANE4-PLAN.md`; H-24's conditional formal-methods scope at `§8:404` and `:422`.
- Sample accounting constraint at `STEP-3-LANE4-PLAN.md:81` (exactly 30 governed turns + 100 tool calls); F106-R7 leg-cardinality at `…-20260721-073500.md:58`.
- m-8's `:223`-is-response-side correction accepted; fixture 14 request-side surface at its cited `§2.1:115,134,136`.
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this packet relay + one INDEX.md row. No amendment drafted or ratified, no reading pinned, no composite authored into any artifact, no fixture/manifest/lock/owner/frozen byte moved, no held member resumed, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the **operator** answers the four decisions; the **VP** reviews this packet against the five bound returns. On the decisions master drafts the additive supersession record for VP exact-byte review and operator ratification, routes the three owed owner confirmations, and then issues `step3-relock-lane4-l4-esc1-resume`. Ratification, amendment drafting, lane-4 resume, fixture materialization/freeze, re-lock, T4, and external use remain held. `…-esc1-fork-1` is SUPERSEDED.
