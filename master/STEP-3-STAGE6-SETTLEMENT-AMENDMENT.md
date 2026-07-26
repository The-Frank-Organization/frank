# Step-3 Stage-6 §D-Settlement Amendment — bounded mechanism-corrections (rev4)

**Status:** AUTHORED (master.orchestrator-planner) — rev4 folds the single VP review-r3 finding SETTLE-VP-R3-F1 (`step3-relock-settlement-amend/RECONCILE-orchestrator-reviewer-20260725-140000`) — scoping the timeless-fold rule to the changed owners (m-9, m-10) so it does not contradict m-2's byte-final exemption — preserving all SETTLE-VP-R2-F1…F3 (rev3) and SETTLE-VP-R1-F1…F4 (rev2) closures and the m-2 cell binding byte-exact. Routed for VP exact-byte re-review → operator hash-bound ratification. Master does NOT self-ratify (§8b).
**Instrument type:** additive amendment. It amends **no bytes in place.**
**Amends (additively):** the ratified Stage-6 amendment **rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`** and the frozen m-10 App-Control-Plane contract **r40 `d2ce9831…`**; it explicitly supersedes the scope of one clause in the frozen m-9 lifecycle **r21 `4d3bd14e…`** (Correction 4) effective only on operator ratification, without editing r21.
**Byte-exactness rule (binding):** rev12, r40, and r21 remain **byte-exact and unmoved** — historical source. This amendment is **self-complete**: every lifecycle outcome, closed domain, presence rule, frame assertion, and acceptance predicate it depends on is stated in these bytes and does not rest on any owner's working (non-bindable) delta. On operator ratification the owners fold these corrections into their own artifacts through their normal F73 pair reviews (see the propagation matrix).
**Provenance:** the four §D-settlement escalations were pair-traced (`step3-relock-dag-m10/RECONCILE-planner-20260722-224500`), master-ruled (`…-230000` `922b796d…`); Correction 3's exact cell is owner-authored + pair-approved by m-2 (`…SITREP-planner-20260723-150000`, `5ec7a3d2…`).

---

## Correction 1 — D-4 Gate-2 claim: honest relabel *(VP-passed; unchanged)*

**What rev12/r40 ratified:** the D-4 disclosure mechanism as a **two-gate guarantee** (Gate 1 = `turn_open`, Gate 2 = `attempt_open_ok`), r40 `:72`/`:81`, §D.4 row 8; m-9 mirrors it at §2.6.

**The settled fact (both pairs' independent traces accepted):** the Gate-2 **comparison** is **vacuous on every reachable MVP state** — the run's parked set cannot grow, and a parked row's fields cannot change, while a worker is alive between Gate 1 and Gate 2; the `UNKNOWN_TOOL_OUTCOME` / `PARTIAL_TOOL_EFFECT` writes sit inside §B.4 step-1 retirement, which kills the awaiting worker, so the §D.4 comparison terminals are unreachable for a parked row's retired generation.

**Correction (relabel the claim; comparator bytes stay):**
> Gate 1 delivers the disclosure guarantee. Gate 2 is a **fail-closed validator + drift-detector over states unreachable on the MVP's bytes** — it retains its well-formedness / malformed / duplicate validation (reachable and useful) and its comparator (which costs nothing and catches exactly the producer drift a future m-10 revision could introduce), but it is **not** claimed as a second independent disclosure gate on the MVP.

The comparator bytes in r40 §D.4 and m-9 §2.6 are **not deleted** — historical, unchanged. Only the ratified operator-visible *claim* ("two-gate guarantee") is corrected to the honest form above. It changes what a ratified operator-visible mechanism claims → operator ratification required; an additive claim amendment is the sufficient instrument (VP-confirmed).

---

## Correction 2 — run-wide parked-set restore + a bounded loud terminal *(rev2: full lifecycle + structural frame bounds)*

**The defect (Governance-Decay instance #2, backlogged):** m-9's rev6 had turn-scoped the parked set, breaking the ratified **worker-independent** disclosure guarantee (r40 `:72`). A parked-unknown disclosure landing in assembled content is evictable by ratified compaction, so a consumer-side re-derivation is weakest exactly on the `content_lost` / `DEGRADED` path.

### 2.1 Run-wide restore (direction, VP-passed)
Restore the ratified **run-wide, worker-independent** carriage on **both** D-4 frames (`turn_open` and `attempt_open_ok`), delivering the ratified guarantee verbatim with m-9's comparator domain restored (no worker-dependent re-derivation). m-9's turn-scoped / consumer-side fallback is **rejected**.

### 2.2 The bounded loud terminal — full durable lifecycle *(closes SETTLE-VP-R1-F1)*
The run-wide parked set has **no static bound** (G-2 resets on a completed turn; §2a bounds are per-turn; no turns-per-run constant exists — both pairs' negative search accepted). Truncation is **forbidden** (silent under-disclosure is the exact harm D-4 prevents); unbounded growth makes the frame-overflow terminal silently reachable. Therefore a bounded loud terminal, specified to the same completeness bar the ratified `resume_frame_overflow` terminal met:

1. **Atomic transaction.** The ordinary r40 §B.4 retirement transaction commits **in full** — every parked row and all normal fence/retirement effects — and **that same transaction** commits `runs.state = FAILED` and `stop_reason = parked_unknown_capacity_exceeded`. No partial commit; fencing is never traded for the bound.
2. **Closed reason domain + present-IFF.** This amendment adds `parked_unknown_capacity_exceeded` to the closed D-4/resume run-terminal `stop_reason` set, which becomes exactly **`{resume_frame_overflow, parked_unknown_capacity_exceeded}`**. Present-IFF: `stop_reason = parked_unknown_capacity_exceeded` **iff** the post-commit run-wide parked count would exceed `MAX_PARKED_ROWS_PER_RUN`; `stop_reason = resume_frame_overflow` **iff** the resume frame overflows (its ratified trigger, unchanged). The two are mutually exclusive on a run. **`resume_action = operator_new_run` remains EXCLUSIVE to `resume_frame_overflow`; the cap terminal carries NO `resume_action`.**
3. **No revival.** The cap terminal admits **no** successor generation, same-run continuation, lease, snapshot, or revival. The operator surface renders the terminal **and the complete queryable parked set** (every parked identity retained and queryable); any next work is an **ordinary new run**. The deferred Step-4 direct-operator parked-unknown clearing path remains **non-MVP** (backlogged).
4. **Acceptance predicate (in this amendment, not deferred to the F73 fold).** A predicate observes, in ONE transaction: the full retirement batch (all parked rows committed), the single `runs.state=FAILED` + `stop_reason=parked_unknown_capacity_exceeded` terminal, **no** successor/lease/snapshot/revival, and the durable operator projection rendering the terminal + the complete queryable parked set. Terminality and revival are fixed here, not left to the later fold.

### 2.3 Threshold interpretation *(closes SETTLE-VP-R1-F2, part B)*
`MAX_PARKED_ROWS_PER_RUN` is the **maximum parked count a run may hold and still remain NONTERMINAL and emit another D-4 frame** — **not** a hard storage-truncation limit:
- `511 + 1 = 512` → the run **continues** (512 is admissible and nonterminal);
- `512 + 1 = 513` → the retirement commits the full row **and** the terminal **atomically** (§2.2);
- a **multi-row** retirement may legitimately **overshoot** past 512 (e.g. commit many rows in one transaction), with **every row retained and queryable**, then commit the terminal; no later D-4 frame is emitted because the run is terminal.

The acceptance predicate must include the `511/512/513` boundary, the multi-row-overshoot, and the no-prefix cases, so `MAX_PARKED_ROWS_PER_RUN` **cannot** be implemented as a 512-row storage cap.

### 2.4 Structural frame bounds — two compile-time assertions over two carrier shapes *(closes SETTLE-VP-R1-F2 part A; corrected per SETTLE-VP-R2-F1)*
Both D-4 **carrier shapes** — `turn_open` and `attempt_open_ok` — must be proven **structurally** (compile-time assertions over the production limits table, not prose arithmetic). The parked set itself has **exactly one growth site** — the §B.4 step-1 parking inside the retirement transaction (r40 `:81`); the two assertions are per-carrier-shape, not per-growth-site. With `PARKED_ROW_MAX = 640` and the ratified frame budget `FRAME_MAX = 4,194,304` (4 MiB):

```text
PARKED_MAX = MAX_PARKED_ROWS_PER_RUN * PARKED_ROW_MAX = 512 * 640 = 327,680

# turn_open carrier shape — the PRODUCTION limits-table / envelope sum:
ADMISSION_REF_ENC_MAX + M_MAX + PARKED_MAX + PATH_MAX_M10 + OVERHEAD_MAX  <=  FRAME_MAX
  2,564,096 + 1,232,896 + 327,680 + 4,096 + 65,536 = 4,194,304  <=  4,194,304   ✓
  # equals FRAME_MAX BY CONSTRUCTION: ADMISSION_REF_ENC_MAX is DEFINED as the residual
  # (FRAME_MAX − the other four ceilings). The envelope sum SATURATES FRAME_MAX; it is
  # NOT a claim that any legal production frame attains 4 MiB (see below).

# attempt_open_ok carrier shape:
ATTEMPT_ACK_MEMBERS_MAX + PARKED_MAX  <=  FRAME_MAX
  1,024 + 327,680 = 328,704  <=  4,194,304   ✓  (~12× headroom)
```

The fresh amendment **requires both compile-time assertions**; if either envelope sum would exceed `FRAME_MAX` the build fails, so raising any ceiling (e.g. `MAX_PARKED_ROWS_PER_RUN`) that broke the envelope is a build error, not a runtime surprise — that is the intended tripwire. **No legal production frame attains `FRAME_MAX`.** The `4,194,304` figure is the **production limits-table / envelope sum over conservative ceilings that cannot all be attained simultaneously** — it saturates `FRAME_MAX` by construction, but no production-equality witness exists. The proven bound on an actual legal frame is `FRAME_CONTENT_BOUND = 3,704,832 B`, a **conservative upper bound** on any legal production frame (m-10's owner artifact): the production max-witness fixture asserts its measured size **≤ `FRAME_CONTENT_BOUND`**; **exact-fit-at-the-limit and one-byte-over fixtures exist only under the build-tagged, test-only reduced limits table** (JCS has no padding to turn the `≤`-bounds into an attained size, so no production fixture asserts equality). Under the production assertion `resume_frame_overflow` is statically unreachable and remains the fail-closed runtime backstop — the no-un-emittable-commit invariant is never conditional on the arithmetic being right.

The run-wide restore is back-to-ratified; the **new operator-visible terminal `parked_unknown_capacity_exceeded`** + the two compile-time frame assertions are the changes requiring operator ratification.

---

## Correction 3 — §5-C `relay.submit` `canonical_resource` cell *(VP-passed; unchanged; bound by hash)*

**The gap:** ratified §5-C marks `canonical_resource` REQUIRED for `relay.*` as "relay verb + target id", but **`relay.submit` structurally has no pre-existing target id** — it creates the record it names (`relay.read` has `relay_id`; `relay.project` a view token; `submit` has neither), and the table admits no `∅` for the relay column.

**Correction — bind the owner-authored, pair-approved cell by hash.** `master/domains/m-2-forms-determinism/design/2026-07-23-stage6-c-relay-submit-resource.md` rev1, SHA-256 **`5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f`** (review-r2 byte-bound approve); its §5 carries the amendment replacement-cell bytes verbatim. Normative shape:

```
canonical_resource(relay.submit) = "relay.submit:" || SHA-256(JCS{ form_digest, dispatch_id?, to?, cc? | cc_unparsed? })
```

- **`form_digest`** REQUIRED → the cell is **total** (no MVP branch yields an unfillable value); realizes my `230000` (4) **form-schema-derived target**, not `∅` (∅ discards real invocation context — considered and rejected, m-2 §1; VP-confirmed rejection of `canonical_resource = null`).
- **`dispatch_id` / `to` / `cc`** — destination coordinates, omitted when absent; CC branch = decoded string-array **or** `cc_unparsed` (mutually exclusive, distinct member names).
- **Owner boundary (m-2 §2.3a):** binds targets **as the invocation names them**; does NOT re-derive the store delivery-set computation. Binding CC as an effect-delivery target **confers no relay authority** — the standing TO/CC authority protocol is untouched.

The other five §5-C action families are settled and unchanged. This corrects only the `relay.submit` cell. The m-2 cell stays **byte-bound at `5ec7a3d2…`** and needs no new pair cycle unless its bytes move.

---

## Correction 4 — `turn_failed` zero-attempt branch: explicit semantic supersession *(rev2: reframed per SETTLE-VP-R1-F3)*

**Frozen r21 `:115` is scope-limiting text:** it says `turn_failed` is a machinery failure **after** the bounded attempt(s). The fact that the distinct `turn_denied` enum member has one pre-attempt branch proves only that the terminal-frame family **can** operate before an attempt; it does **not** make `turn_failed`'s own "after" clause descriptive. **This amendment does NOT claim frozen r21 already meant this.**

**Correction — an explicit, bounded semantic supersession (r21 byte-exact as historical source; no in-place edit):**
> Effective **only upon operator ratification**, this amendment **supersedes** r21 `:115`'s after-attempt restriction on `turn_failed` **for exactly one named branch**: the pre-attempt assembly-refusal branch. r21 remains byte-frozen historical source; nothing edits it.

**The one superseded branch, bounded to the pair-reviewed shape:** an **ADMITTED** turn that fails totality / shape / alias validation in **ASSEMBLING** yields — with **zero** `logical_surface_digest`, `attempt_open`, DATA-P request, and provider-attempt row — **exactly one** existing `turn_terminal{…, terminal: turn_failed}` followed by the ordinary `turn_receipt{terminal_recorded}`; **no auto-retry and no second assembly.** **Any other new zero-attempt use of `turn_failed` remains a fresh lifecycle amendment** — this supersession does not generalize beyond the named branch.

**Adjudication (VP-confirmed):** the additive-amendment instrument is correct and r21 stays byte-exact; the framing is corrected from "already-descriptive" to "explicit supersession effective on ratification."

---

## Propagation — explicit owner fold matrix *(closes SETTLE-VP-R1-F4)*

On operator ratification, each **changed owner (m-9 and m-10)** folds its halves through a **fresh F73 pair review**; affected consumer confirmations and the two-sided §D join occur **only afterward**. (The shorthand "owners fold all corrections" is replaced by this matrix.) **Timeless-fold rule (SETTLE-VP-R2-F2, scoped per SETTLE-VP-R3-F1 — the packet binds no changed-owner working-state snapshot):** for each **changed owner (m-9, m-10)** — i.e. an owner with a non-empty post-ratification fold obligation in the matrix below — no pre-ratification working artifact, in-flight revision, or relabel is a durable post-ratification fold, regardless of how far its content has advanced; after operator ratification that owner produces a **fresh, pair-reviewed successor over its then-current owner artifact**, and any revision that pre-folded this correction while the amendment iterated is historical ancestry, never a substitute for that successor. **This rule does NOT apply to m-2:** the bound cell `5ec7a3d2…` is an **already pair-approved ratification-packet component** (`step3-relock-c-m2-submit-resource-review-r2`, approved `20260723-140000`), **not** a pre-folded working artifact — its matrix row stays `None`, and its "no new pair cycle unless its bytes move" rule survives (requiring an m-2 successor would either force a no-op cycle or move the cell bytes and void this packet's exact hash binding).

| Owner | Folds |
|---|---|
| **m-9** | Correction 1's **consumer-side** Gate-2 label; Correction 2's **full run-wide consumer + comparator** semantics (restore) **and** the Correction-2 terminal's consumer posture; Correction 3's bound `relay.submit` cell (into §7-`relay.*`); **Correction 4 — REPLACE the pair-approved Section 6 classification (m-9 delta `04422965…` `:423-426`, "after the bounded attempt(s) is DESCRIPTIVE … an owner clarification") with the amendment-controlled explicit semantic supersession (Correction 4, effective only on operator ratification, bounded to the named pre-attempt assembly-refusal branch).** That §6 replacement may be batched with the §2.6 Gate-2 label and the §7-`relay.*` folds, but it is the **§6 `:423-426` text** that must be superseded — not merely a §2.6 addition. One batched revision under F73. |
| **m-10** | Correction 1's **producer-side** claim relabel; Correction 2's **run-wide producer**, the `parked_unknown_capacity_exceeded` terminal, the schema/presence rules, the **two compile-time frame assertions** (§2.4), and fixtures (incl. the §2.3 boundary/multi-row/no-prefix acceptance predicate + the §2.4 production max-witness `≤ FRAME_CONTENT_BOUND` and reduced-table exact-fit fixtures). Fresh full-byte pair review over a successor to the **then-current** m-10 owner artifact (per the timeless-fold rule above — no pre-ratification revision is the durable fold). |
| **m-2** | None — the bound cell `5ec7a3d2…` is byte-final; no new pair cycle unless its bytes move. |

---

## What this amendment does NOT do

- It moves **no** byte of rev12, r40, r21, or any frozen owner final.
- It authorizes **no** m-10 schema change beyond Correction 2's named terminal, the restored run-wide carriage, and the two frame assertions (owner-folded under F73 post-ratification).
- It does **not** touch the ratified §5-E five-member set / JCS-SHA-256 formula (the delegated A3/B1 recipes settled separately at m-9 r12 `04422965…` under the VP's F73 classification `4c254307…`).
- It licenses **no** §D join, integrated re-lock, DESIGN-lock, PLAN, T4/code, credential, provider call, release binding, live E3, merge, or deploy. **H-12 continues to hard-block external use.**

## Ratification packet + gate

- **Ratification unit:** this amendment file (by its own SHA-256) **+ the bound m-2 cell `5ec7a3d2…`** as one packet.
- **Gate:** VP exact-byte re-review → operator hash-bound ratification (§8b agent-authored + operator-cited; master does not self-ratify).
- **On ratification:** the propagation matrix above executes; the §D two-sided join co-signs after the folds; consumers rebase once.
- **Independent instrument:** the m-3 schema-version amendment (operator-ratified 2026-07-22, rev3 `9e874df8…` + bound contract `6e2abe40…`) is separate; this is the second Stage-6 amendment.
