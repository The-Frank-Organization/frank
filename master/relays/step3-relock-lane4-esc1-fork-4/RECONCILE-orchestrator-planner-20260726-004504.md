## RECONCILE — the OPERATOR PACKET, **rev4**, and the shape has CHANGED: an operator-steered **D1 re-scope** now sits UPSTREAM of the four witness decisions and partly re-cuts them. Four owners assessed the durable session log against m-9's own 8-harness survey and **all four agree it is not all load-bearing**; master then verified the baseline directly against `references/` source. **The operator has ruled hash-chaining OUT**, on a ground stronger than any owner assessment: it would reject the operator's own `bivpak` repack tool as a forgery. That also answers, by decision rather than by argument, the one question m-3 scoped its agreement around. **`…-esc1-fork-3` is SUPERSEDED — act only on this rev4.** Section A is new and upstream; Decisions 2, 3 and 4 are byte-unchanged from rev2/rev3; Decision 1 is re-cut; my own strengthening (v) is WITHDRAWN.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-fork-4
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-fork-3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — each decision changes what value is frozen into the Step-3 exit gate, touches a ratified §7 field, or changes the exit suite's leg set. Operator ratification + VP exact-byte review; master proposes and may not self-ratify.
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_RECORD_KIND: design-doc
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-m3-scope-ans-1/DESIGN-planner-20260726-003000.md
FROM: master.orchestrator-planner
TO: operator
CC: master.orchestrator-reviewer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer, m-3.planner, m-3.implementer, m-8.planner, m-8.implementer, l4.planner, l4.implementer
SUBJECT: rev4 — UPSTREAM D1 re-scope (chaining OUT by operator ruling: malice-only, redundant under a frozen oracle, and it would reject the operator's own bivpak repack; runtime self-integrity ruled out of scope by the same decision) + the witness re-cut over the floor (`log_prefix_digest` becomes a full ordered list vs the frozen expected; strengthening (v) withdrawn); prior packet — the composite `context_digest` fully specified (harness-composed from two independently carried halves, so no producer gains a cross-domain dependency); recommend defer observer-parity on the RATIFIED v2 precedent for the sibling digest; two blockers (authored-not-harvested expected; a designed-but-unexercised fencing gate); rev3 adds two m-9-pending conditions to Decision 1; supersedes `…-esc1-fork-1` AND `…-esc1-fork-2`

## What changed, across three supersessions

**rev3 → rev4 (this packet):** a new **Section A** — the D1 re-scope — sits upstream of everything and partly re-cuts Decision 1. Condition (iv) is CONFIRMED by m-9 and its importance has grown (it is the frozen oracle that makes chaining redundant). **Condition (v) is WITHDRAWN by me** — it existed only to exercise chaining, which is going. Decisions 2, 3 and 4 are byte-unchanged from rev2 and can be verified as such by diff. `…-esc1-fork-3` is superseded — read only this one.

**rev2 → rev3:** Decision 1 gained conditions (iv) and (v), proposed pending m-9.

**rev1 → rev2:** `…-esc1-fork-1/RECONCILE-orchestrator-planner-20260725-200600.md` is **superseded in full — do not act on it.** Two of its claims were wrong:

1. **The scope closure was a false attribution.** It said VP r7 required *"exact resumed-context identity incl. the model-visible continuation input and the settlement snapshot"* and that this enumerated exactly two members, excluding epoch/lease. **VP r7 contains no such sentence** — it says only *"no schema or digest recipe for the claimed exact round/context identity"* plus *"**for example** a predecessor/round/log-prefix/context digest vector"* (`step3-arch-packet/…-20260721-073500.md:62`). The quoted phrase was the **VP's own later question**, prefixed *"including"*. I closed a question m-10 had deliberately routed up, using invented authority. Withdrawn and filed to m-10 at `…-esc1-m10-refile-2`.
2. **It framed the boundary as a blocking selection.** m-3 showed that is a category error (below).

**The withdrawal paid for itself.** Because m-10 was re-asked instead of overruled, it produced the frozen fencing gate that covers epoch/lease properly — *and* the observation that nothing in the exit suite exercises it. That gap (Decision 3) would not otherwise exist on the record.

## The TEN bound returns — exact SHA-256, all uniquely addressable

| owner | relay | SHA-256 |
|---|---|---|
| **m-9** (original) | `…-esc1-m9-answer-1/DESIGN-planner-20260725-195230.md` | `9867d5070db25da93b64974ff3b66443a201312d1f9a8a7e034adfe14f1954c1` |
| **m-9** (executable spec) | `…-esc1-m9-spec-1/DESIGN-planner-20260725-211230.md` | `50085fd3ce5107e615697cba266dd16d804b39762d07584b196525df242d5490` |
| **m-10** | `…-esc1-m10-answer-1/SITREP-planner-20260725-211200.md` | `20ea533e6d78f335e5c936ce333af5a4f56b34f03bf748b3566133f89d3037dc` |
| **m-8** | `…-esc1-m8-answer-1/DESIGN-planner-20260725-211308.md` | `6bc388c7cbdcb7d0acd05e8b06bca207289f9bafcb8572a11cb8304ba52d3511` |
| **m-3** | `…-esc1-m3-answer-1/DESIGN-planner-20260725-212500.md` | `ae92c26898e06b0c72b272c1aa5c67b7c209c34959a76b42b49c210e8c1631c5` |
| **m-9** (conditions) | `…-esc1-m9-conditions-ans-1/DESIGN-planner-20260725-225330.md` | `6f4e4f15023c593dd290dc30b508985c1c293572805a8333a8490b4ebf573cdf` |
| **m-9** (invariant) | `…-esc1-m9-invariant-ans-1/DESIGN-planner-20260725-234500.md` | `4dd39fc03d4d54ff8156abf5e604f6bef5a3dbe8ee8be0ef39448d874629e4af` |
| **m-9** (scope) | `…-esc1-m9-scope-ans-1/DESIGN-planner-20260726-001200.md` | `1b3e368ed3b7c35ed3f506629da9133170adb471c33eaee9da908a2bd322e726` |
| **m-10** (scope) | `…-esc1-m10-scope-ans-1/SITREP-planner-20260726-002900.md` | `46761e8892634ebaf596372957d1d9b611d70364a9d7620b885c9166f2ffc6cf` |
| **m-3** (scope) | `…-esc1-m3-scope-ans-1/DESIGN-planner-20260726-003000.md` | `2720d109850b635c60cbd82e1a4a96b8d8df8c50fccf637e7b2af9593a2fc0cb` |

The m-10 and m-3 lineage defects are repaired: both re-filed under fresh unique ids parented to their requests, originals left unmutated. **That defect was mine twice** — I was corrected on it for m-9, then dispatched m-10 and m-3 without distinct answer ids, creating two fresh instances of it.

## SETTLED — no decision needed from you on any of this

- **The witness is REQUIRED.** Four independent proofs, none reachable by search: m-9 (§6:389 excludes assembled `input[]` from every existing digest; compaction can move the assembly), m-10 (receipt-presence divergence — identical log prefix, different settlement class), m-3 (absence admits the exact false PASS §7 forbids), and m-8 (confirms (A) and (B) are genuinely different objects). My "strike it as vestigial" lean is withdrawn without qualification.
- **The (A)/(B) fork is DISSOLVED for the core witness.** m-3's distinction: the **producer** boundary is what m-9 hashes; the **observer** boundary is what the E3 gate captures. m-9 rides the value on `attempt_open` exactly as `logical_surface_digest` rides today, so the gate compares a **carried actual** against a **harness-frozen expected** — *"executable by the gate I bound at r24, under EITHER producer boundary, needing NO direct capture of (A) or (B)"*. Independence comes from the **frozen expected**, not from re-derivation. **No owner reports a blocker to the witness itself.**
- **Epoch/lease is OUT of the witness** — m-10 and m-3 concur, and not on "it isn't model-visible" but on a named **frozen fail-closed** gate: rev16 `§4:55` (atomic lease-at-current-epoch admission) + `§6:130` (two-form assign gate keyed by `{run_id, generation_id, turn_epoch, state_seq}`, exactly-once, stale-rejecting via `state_seq`). m-3 confirms it can name no predicate of its own that covers successor-legitimacy, and that folding it into `xit-dur-1` would conflate two axes. **The residual is Decision 3, not the scope question.**
- **m-9's producer boundary is (A)**, the app-internal `m8.llm_request.v1.input[]`, and m-9 **owned its earlier overclaim** that the witness hashes "what the model received on the wire" — that was (B), which m-8 §1.1 forbids conflating.

## SECTION A (NEW, UPSTREAM) — the D1 re-scope

Four owners assessed the durable session log; I then verified the baseline directly against the `references/` source rather than reciting it.

**A1 — chaining is OUT. You have already ruled; this records it and the three independent grounds.**
- **m-9, against its own design:** its §9 survey's *original* conclusion was fresh-start, the eight surveyed harnesses did **not** motivate a chained durable transcript, and the apparatus is a later addition responding to the stage-5.1 fault. Over per-record checksums plus `seq` contiguity, chaining adds detection of **only** a wholesale re-numbered, re-checksummed replacement — a crafted-consistent fake history, i.e. **malice**, which the ratified threat model excludes.
- **m-3, as evidence owner, tested rather than accepted it** and supplied a better reason than the throughput one I offered: condition (iv) gives the exit leg a **frozen oracle**. Chaining is *internal self-binding* — it lets a log attest to itself **with no external ground truth**. The exit leg *has* ground truth, so the chain is redundant there, **and that reason is robust to scale** where mine was not.
- **The operator's ground, which is decisive and stronger than either:** the `bivpak` repack tool rewrites paths inside session history and reopens it on another machine. That is a wholesale internally-consistent rewrite — **exactly and only what chaining detects** — performed legitimately by the operator's own tooling. Chaining would make correct output indistinguishable from a forged log.

**Verified against source, not recollection.** Across all eight harnesses in `references/`, every match for chaining/tamper identifiers was unrelated — JWT verification, sandbox policy, cert bundles, cache validation, CDN poisoning, release-note hashing, browser stealth. **Zero session-log hash chains.** Codex is the sharpest case: its file is `trace.jsonl`, written via `append_with_context_best_effort` (failures are `warn!`-ed and swallowed), and the entire rollout-trace module contains **zero** fsync, sync_all, checksum or crc. **The production baseline is weaker than I previously told you** — not "JSONL with engine checksums" but best-effort append with no durability guarantee. The name is the tell: a *trace*, not a journal.

**A2 — runtime self-integrity: OUT OF SCOPE, by your decision rather than by argument.** m-3 scoped its agreement precisely: it did **not** rule on oracle-free *runtime* self-integrity — production resume detecting a spliced log with no frozen expected in hand — because that is the m-1 verifiability-against-the-courier axis you already deferred. **`bivpak` answers it:** you cannot want production resume to reject rewritten logs while building a tool whose purpose is to rewrite them. Recorded as a decision so nobody re-derives the apparatus in six months from the same instincts that produced it.

**A3 — the cheap path, from m-10: a DERIVATION trim is free; a MEMBER-SET change is not.** m-10 binds **values, not derivations** — rev16 §2:41 already states the fold *"does NOT rest on that derivation"* — and it stores `marker_digest` for **equality matching only**, never re-deriving it. Its manifest union keys on **attempt identity**, not chaining, so it **affirmatively releases** the chained property. It will bind an **abstract round identity** by four properties (stable per round, unique per round, byte-reproduced verbatim, equality-comparable). **So: keep the member names and change only what is underneath, and the co-signed §D join never moves — no amendment for m-10's half.** *Removing* a member (e.g. `segment_id` under one-file-per-run) is the amendment-shaped part, and it is m-3-joined because `segment_id`/`seq_hwm` are m-3's E3 locators.

**A4 — the rest of the releasable set: OPEN, and this is a real question for you.** m-9 released, beyond chaining: **size rotation** (never fires in one turn), the **terminal seal**, and the **cross-segment boundary equation**. One honest correction to my own earlier premise: cross-segment is **not** pure dead code today — it is reachable through *generation replacement*, because the design models each generation as a new chained segment. It becomes dead only if you **also** simplify to **one-file-per-run**. `bivpak` arguably argues for that anyway: one file relocates far more cleanly than a chained segment set.

**The floor that remains** (m-9's, and reviewer-aligned): a **typed run journal** + **per-record checksum** + a **checkpoint at each settled tool-round boundary** + the **per-run writer fence**. Every element ties to an admitted hazard — torn writes, or the disposed-but-live predecessor race — and **none rests on tamper-evidence**. Note this floor is still **above** every harness in `references/`; that is not gold-plating, it is what a *tested* resume claim costs, and going below it to a bare best-effort file means giving up that claim.

**A5 — a forward constraint worth recording now, cheap now and expensive later.** If `bivpak` rewrites content inside records, **every digest over those bytes breaks — not just the chain.** Per-record checksums go invalid per rewritten record (locally recomputable). **m-10's round identity is the hard one:** it *stores* the value and a mismatch for the same key is `receipt_conflict`, **fail-closed, first-committed stands** — so a repacked session hard-conflicts rather than degrading. The standing constraint: either session digests are computed over a **path-independent canonical form**, or relocation is an **explicit authorized rebase across both stores**. m-10's abstract-round-identity concession is the seam a bolt-on adapter would plug into. *(Unverified and worth confirming before the record schema freezes: whether the log actually stores absolute paths — tool arguments, workspace snapshots and objective references are the likely carriers.)*

## Decision 1 (RE-CUT over the floor) — ratify the ordered-list form of `log_prefix_digest`?

**Reading (b) as a single chained value does not survive Section A** — it *was* the chain. m-3 named the replacement: the prefix-identity term becomes the **full ordered `{seq, record_digest}` list** (equivalently the per-round checkpoint list) over the valid prefix, compared against the **frozen expected**.

**The leg survives in the same terms, and m-3 verified rather than assumed it.** A degraded re-derivation manifests as a changed record, a missing or extra record, a reordering, or a re-numbering — each surfaces as a mismatch at ≥1 ordered position. So *"a degraded re-derivation never satisfies `xit-dur-1`"* holds unchanged, and m-3 notes it holds **more legibly**: a position-level mismatch is far more diagnosable than one opaque digest miscompare.

**Conditions, updated:**
- **(i)** the form is ratified; **(ii)** the fixture pins the exact interval; **(iii)** the **harness — never the T4 build — computes `expected`**.
- **(iv) CONFIRMED by m-9 and now load-bearing beyond its original purpose:** freeze the **authored record contents**, and the harness derives every value from them. m-9 gave the clean split — the structural digests are harness-derived, everything else is authorable, and no member is engine-assigned-and-unauthorable. It flagged one to watch: `ts_monotonic` is safe **only** because the prefix is read and replayed, never re-stamped. This condition is what supplies the frozen oracle that makes chaining redundant, so it carries more weight now than when I proposed it.
- **(v) WITHDRAWN by me.** I proposed scripting ≥2 durable checkpoints so the chaining would be exercised. With chaining gone there is nothing to exercise. m-3 confirms its leg needs one durable boundary checkpoint plus the ordered list — not the ≥2 strengthening, not the cross-segment case, not the chaining-discrimination negative legs. **m-9's P2 refinement falls away with it.**

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

**Section A — the re-scope (upstream; answer these first):**
1. **A1/A2 — confirm on the record** that chaining is out and that runtime self-integrity is out of scope for the MVP. You have said both; I am asking you to make them citable so they are not re-derived later.
2. **A4 — take the rest of the releasable set?** Size rotation, terminal seal, cross-segment boundary equation — **plus** the one-file-per-run simplification without which cross-segment is not actually dead. Master recommends **yes, all of it**, and `bivpak` favours one-file-per-run independently.
3. **A3 — the member-set question, priced separately.** Keep member names stable (free, no §D amendment) and defer removing `segment_id`/`seq_hwm`? Or remove them now as one bounded amendment, which also needs m-3 since they are its E3 locators? Master recommends **keep the names, defer the removal**.

**The witness (downstream of A):**
4. **Ratify the ordered-list form** of `log_prefix_digest` with conditions (i)–(iv) binding, (v) withdrawn?
5. **Ratify the composite `context_digest`** as specified — harness-composed, `v:1`, explicit-null, both halves recorded in evidence?
6. **Parity: defer (a), or take it (b)?** Master recommends **(a)** on the ratified v2 precedent for the sibling digest.
7. **The fencing gap: an eleventh record (i), or an accepted documented gap (ii)?**
8. **Hygiene — approve the provisional annotation?** m-9 recommends a cheap one-line marker now on the two stale stage-4 loci (*"superseded — durable replay per r17 §5; final shape pending the D1 scope re-cut"*), with the final named supersession authored once against the settled shape. It marks a GRILL-ratified document, so it needs your word. The objective-assembly pin stays **held** until A resolves — m-9's recommendation and mine, so it is authored once rather than twice.

On your decisions I draft the additive supersession record, route its exact bytes to the VP, return it for your ratification, and only then issue `…-l4-esc1-resume` to the pair.

## Boundaries
This relay ratifies nothing, pins no reading, authors no amendment, changes no fixture or manifest, moves no locked byte, resumes no held member, requests no owner-fidelity, issues no PLAN/T4 token, touches no `frank/` path, and permits no external use. The composite formula above is a **proposal for your ratification**, not an authored artifact. Interface lock `cbd1893c…`, stage-6 amendment `1125b0a0…`, m-9 delta `01b885fe…`, m-8 contract `4b670a79…` + r7 addendum `734e44b7…`, m-10 rev16 `3e3c5192…`, m-3 r24 `651c9aec…` all UNMOVED. **H-12 hard-blocks external use.**

## Verification
- All five owner returns read at their bytes and SHA-256-bound above (recomputed from disk this turn); rev2 body copied byte-exact and edited only in Decision 1, so every other section is unchanged from `…-esc1-fork-2` (SHA-256 `658ad8ec679a7841efd42e088bd67dea3b9e6db9b68347cfb4f5841a7576f8df`).
- **Section A verified against source this turn, not recited:** across all eight harnesses in `references/`, every chaining/tamper identifier match was unrelated to session logs (JWT/sandbox/certs/cache/CDN/release-notes/browser-stealth) — **zero session-log hash chains**; codex's `trace.jsonl` writes via `append_with_context_best_effort` (`codex/codex-rs/rollout-trace/src/{compaction,code_cell}.rs`) with **zero** fsync/sync_all/checksum/crc in the module. m-10's derivation-agnostic statement re-read at its cite (rev16 §2:41, §2:39, §1:16); m-9's §9 fresh-start conclusion at `2026-07-19-mvp-full-worker.md` §9:17.
- Conditions (iv)/(v) basis verified earlier: §7's "crash after **≥1** `fsync`-durable `round_marker`" at `STEP-3-STAGE6-AMENDMENT.md:371`; NO round/record count pinned across the amendment, `STEP-3-LANE4-PLAN.md`, or any of the nine materialized `master/exit-fixtures/xit-*/input.json` (`xit-dur-1/input.json` deliberately absent). Both routed to m-9 at `…-esc1-m9-conditions-1`; NEITHER confirmed.
- VP r7's actual text re-read verbatim at `master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-073500.md:58-62`; the true source of the two-member phrase at `…-esc1-amend/RECONCILE-orchestrator-reviewer-20260725-190320.md:49-53`; the superseded paragraph at `…-esc1-fork-1/RECONCILE-orchestrator-planner-20260725-200600.md:48`.
- The v2 carried-digest precedent read verbatim at m-3 r24 `domains/m-3-observation-evidence/design/2026-07-22-stage6-lane2-e0-e3-delta.md:274` ("ratified OUT of v2 … Not designed").
- Fencing coverage checked directly: `epoch|lease|fenc|assign gate|state_seq` returns **no match** across the ten §7 exit rows (`STEP-3-STAGE6-AMENDMENT.md`, §7 row block) and across `STEP-3-LANE4-PLAN.md`; H-24's conditional formal-methods scope at `§8:404` and `:422`.
- Sample accounting constraint at `STEP-3-LANE4-PLAN.md:81` (exactly 30 governed turns + 100 tool calls); F106-R7 leg-cardinality at `…-20260721-073500.md:58`.
- m-8's `:223`-is-response-side correction accepted; fixture 14 request-side surface at its cited `§2.1:115,134,136`.
- Exact-file lint of THIS relay OK. `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this packet relay + one INDEX.md row. No amendment drafted or ratified, no reading pinned, no composite authored into any artifact, no fixture/manifest/lock/owner/frozen byte moved, no held member resumed, no `frank/` action, no PLAN/T4/credential/provider/E3/merge/deploy.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the **operator** answers Section A first (it is upstream and re-cuts the rest), then the witness decisions; the **VP** reviews this packet against the five bound returns. On the decisions master drafts the additive supersession record for VP exact-byte review and operator ratification, routes the three owed owner confirmations, and then issues `step3-relock-lane4-l4-esc1-resume`. Ratification, amendment drafting, lane-4 resume, fixture materialization/freeze, re-lock, T4, and external use remain held. `…-esc1-fork-1`, `…-esc1-fork-2` AND `…-esc1-fork-3` are SUPERSEDED — act only on this rev4.
