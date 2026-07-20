# s9 `scope_paths` co-sign — the m-3 PREDICATE-SEMANTICS section — **rev2 (SETTLED CONTRACT)**

> **DESIGN_DOC_ID:** `s9-scopepaths-cosign` (shared contract) · **m-3 section id:** `s9-scopepaths-m3`
> **rev2 records the SETTLED joint contract**, folding: master's `s9-scopepaths-cosign/RECONCILE-orchestrator-planner-20260713-160510` (MR-1 Option A · segment-prefix blessed · the joint-convergence authorization · MR-3 confirmed) + **m-2's convergence CO-SIGN** `s9-scopepaths-converge/DESIGN-planner-20260713-153200` (the token-granularity FIXED) + m-3.implementer `…-142400`/`…-154000`. **Both halves now agree byte-for-byte.**
> **Split:** m-2 owns the declaration-slot **home** (FieldSpec byte · `visible_when` · declaration-time normalization/validation · the submit suppliability guard). **This doc owns the reader-side PREDICATE** — the match/subset/narrowing semantics over that home · per-phase applicability · the E-rung · the observe dispositions with complete signal bytes.

---

## 0. Frame — what this contract does and does NOT un-strike

The IMPL scope clause is **`diff_paths ⊆ scope_paths`** (c2 §5:98 · s8 §13 IMPL row · m-3-F7). This co-sign gives it a home + a resolution contract (**necessary**), but is **not sufficient**:

- **RHS (`scope_paths`) — this co-sign.** The declared, lineage-resolved scope set. Settled here + in the m-2 leg.
- **LHS (`diff_paths`) — rides design item 10 (the turn-baseline fence), NOT this co-sign.** Attributing tree delta to *this* candidate needs the conductor turn-entry baseline (s8 §13: "what it CANNOT observe: a turn-baseline"). Until item 10 lands, `diff_paths` is **not attributable**, so the `⊆` clause **cannot run** regardless of a co-signed RHS.

**Therefore:** the clause moves **struck (no home) → home-defined-but-unevaluable ⇒ honest degrade**, and is a **live E1 predicate only when BOTH this co-sign AND item 10 are in.** The `diff_paths ⊆ scope_paths` **evaluation stays STRUCK until item 10** — both seats and master hold this carve-out. The conductor claims no vantage it does not have.

---

## 1. The SETTLED contract (four axes — both halves co-signed)

1. **Grammar = normalized segment-prefix** (master-blessed; m-3's globs withdrawn). **Owner split:** *m-2* normalizes + validates at declaration (malformed entry — `..` / leading `/` / empty ⇒ typed reject); *m-3* matches / subsets / narrows (§2).
2. **Declaration site = the accepted PLAN ancestor ONLY** (master MR-1 **Option A**; master owned the "/dispatch" as its own over-specification — scope is a PLAN property, and a grant declaring a different scope would be a second source of scope truth). m-2's home: `visible_when: {phase_in:["PLAN"]}`. The grant-bearing site is a **non-welded Rail-A future door** (an m-2 grammar + m-7 render-context extension), **not an s9 carry**.
3. **Narrowing-locus = observe** (master's steer, both seats accept): the `⊆` predicate fires in the observe layer. m-2's submit-time suppliability guard remains **defense-in-depth**.
4. **The honesty rail (MR-3, master-confirmed):** at `39474d0` **only the conductor parent edge is extant** (`lineage.go:254-262` — one lookup, no recursion; no `scope_paths`, no `resolve_scope`). **The walk · the nearest-scope-bearing-PLAN stop · the declaration-site filter · cycle/broken-chain handling · the candidate-copy veto · m-2's submit guard are ALL RED-first s9 BUILD obligations** — never described as extant machinery.

## 2. §G — the match / subset / narrowing semantics (m-3-owned)

Over m-2's normalized segment-prefix rows:
- **`inScope(p) := ∃ s ∈ resolved_scope : segmentPrefix(s, p)`** — `segmentPrefix(s,p)` holds iff `s`'s normalized path **segments** are a prefix of `p`'s segments. `pkg/a` matches `pkg/a`, `pkg/a/x.go`, `pkg/a/b/y.go`; `pkg/a` does **NOT** match `pkg/ab` (segment-wise, not byte-wise).
- **`diff_paths ⊆ scope_paths := ∀ p ∈ diff_paths : inScope(p)`.**
- **Narrowing:** a child scope-set `S_c` narrows parent `S_p` iff **every child row is itself in-scope of the parent** (`∀ s ∈ S_c : inScope_{S_p}(s)`). Trivially decidable — **no glob-language-containment procedure and no algorithm ceiling is needed** (the reason segment-prefix was the right convergence).

## 3. Pin (a) — canonical value = the accepted PLAN ancestor, never the work-record

The RHS is read **exclusively** from the lineage-resolved **nearest scope-bearing accepted PLAN ancestor**. It is **never** read from `Candidate.Record` — a candidate-borne `scope_paths` is not an input (it is the §5 class-1 event). *A record cannot declare the bound it is then measured against — that is self-certification.*

## 4. Pin (b) — resolution through lineage: the walk + stopping rule

**Rides the conductor-computed parent edge ONLY** (m-1 fidelity §7): `PARENT_DISPATCH_ID` (`owner: system`, `fill_constraints: system_only`, `lineage_role: parent_edge`). The lane-suppliable `parent_hint` (`lineage_role: none`, gated by computed `parent_hint_honored`) is **never** substituted.

**`resolve_scope(Candidate)` — a RED-first BUILD obligation (§1.4), not extant:**
1. Start at `Candidate`; walk **`parent_edge`** upward through **accepted** parents (`checkParentSubstrate`: honored only when `DeliveryState == Accepted`). *(Today `parents()` does one lookup and returns — the recursive walk is new build.)*
2. At each ancestor: is it an **accepted PLAN** record (§1.2) **bearing a `scope_paths` declaration**?
3. **Stop at the nearest (innermost) scope-bearing accepted PLAN** — its rows are `resolved_scope`. Skip ancestors bearing none. Reach the dispatch root/genesis with none ⇒ **∅-declared** (§6).
4. Bounded by the finite committed parent DAG. An unwalkable chain (absent/corrupt edge before a root) ⇒ **broken-chain** (§6, machinery-fault).

Nearest-wins is sound **because §2's narrowing rule + the §5 `scope-exceeded` veto forbid a nested PLAN from widening** — so the innermost declaration can only ever be ⊆ its ancestors.

## 5. Pin (c) / §5 — the TWO-token class model (m-2's granularity call, master-assigned; m-3 binds)

m-2 fixed the granularity at **two tokens, split by LAYER**, diverging from my one-token-plus-reason recommendation **on the merits, correctly**: a submit `Violation.Class` and an observe predicate disposition (`MachineryFault` / terminal-by-authority) are **structurally different signal shapes in different value-spaces**; and condition (i) is *caught at submit or ignored at observe*, so a one-token observe **reason** for it **could never actually fire**. Two layer-pure tokens is the deterministic model. **m-3 binds to it exactly:**

| token | layer | owner | event | signal bytes |
|---|---|---|---|---|
| **`scope-self-declared`** | **submit** | m-2 (suppliability guard — *RED-first build, not extant*) | a **non-PLAN / work record supplies its own `scope_paths`** (structural self-certification) | a submit-layer `Violation{Class:"scope-self-declared"}`. **`MachineryFault` n/a** — it never reaches observe. **If the guard has not landed**, resolution-by-construction simply **ignores** the copy (§3) — **no observe token fires for it.** |
| **`scope-exceeded`** | **observe** | **m-3 (this predicate)** | the record's **effective scope exceeds its resolved PLAN-ancestor's scope** (substantive over-reach) — **MERGES** the two operand cases under one invariant | `predicate: fail` (**observed-false**), **`MachineryFault: false`**, terminal **`rejected` both authority classes**. A **path-free bounded reason** `{plan-widens-ancestor, diff-drift}` distinguishes the operands — **never the paths** (§8). |

**Why the merge is right (the invariant is ONE):** *a record's effective scope may not exceed its authorizing accepted-PLAN-ancestor's scope.* `plan-widens-ancestor` compares the **declared** effective scope (a nested PLAN's rows) against the ancestor's; `diff-drift` compares the **enacted** effective scope (`diff_paths`) against the resolved rows. Same invariant, different operands — **one observe token, two reasons.**

## 6. Pin (d) / §6 — the complete observe disposition table

| resolution outcome | predicate | class · reason | `MachineryFault` | terminal by authority | rung |
|---|---|---|---|---|---|
| **∅-declared** (walk reaches root with no scope-bearing accepted PLAN) | `degraded` | `scope-unbounded-no-declaration` | **false** | **no-vantage:** authority → decision-② → **`held`**; non-authority → `accepted` + `self_reported` | **E0** |
| **∅-attributable-LHS** (item-10 not landed) | `degraded` | `scope-lhs-unattributable` | **false** | as above (no-vantage) | **E0** |
| **`scope-exceeded` · `plan-widens-ancestor`** (a nested accepted PLAN declares rows exceeding its resolved PLAN-ancestor) | `fail` | `scope-exceeded` · reason `plan-widens-ancestor` | **false** | **`rejected` both classes** (observed-false) | terminal |
| **`scope-exceeded` · `diff-drift`** (attributable `diff_paths ⊄ resolved_scope`) | `fail` | `scope-exceeded` · reason `diff-drift` | **false** | **`rejected` both classes** (observed-false) | **E1** — **ITEM-10-GATED** |
| **ambiguous** (≥2 non-orderable scope-bearing PLAN declarations at the frontier) | `blocked` | `scope-source-ambiguous` | **true** | **machinery-fault edge:** authority → `held` + escalate; non-authority → `rejected` / author-return | terminal |
| **broken-chain** (`parent_edge` unwalkable mid-chain) | `blocked` | `scope-lineage-unresolved` | **true** | machinery-fault edge (as above) | terminal |
| **resolved + evaluable, `⊆` holds** | `pass` | — | false | contributes to IMPL done | **E1** |

**The line:** *unevaluable* (no declaration / no attributable LHS) → **degrade, `MachineryFault:false`, never a pass**; *evaluable-and-violated* → **observed-false veto, `MachineryFault:false`**; *ambiguous/broken machinery* → **`MachineryFault:true`, fail-closed by authority**. **No branch silently accepts an un-measured scope.**

## 7. §F — m-1 fidelity (the lineage key)

- **F-1:** the walk consumes the **conductor-resolved `parent_edge` chain ONLY** — never `parent_hint`.
- **F-2:** where an in-chain edge's `parent_provenance` is `hint`, m-1 confirms it entered the chain only because the conductor honored it (`parent_hint_honored: true`) — not a raw lane assertion.

## 8. §H — I-PH: every seat-visible scope surface is symbolic / bounded / path-redacted

Per the locked §6.1: the verdict / bounce / `executable_claim_results` row / `failing_detail` / `degradation_notes` for **every** scope branch carries **only** the bounded symbolic class + the bounded reason enum + (where correlation is needed) a **non-path bounded scalar** (e.g. an offending-row **count**). **Never** the candidate's paths, **never** the resolved ancestor rows, **never** the effective glob/prefix values.

---

## 9. Build obligations (ALL RED-first; only the parent edge is extant — §1.4)

`resolve_scope`'s recursive `parent_edge` walk · the nearest-scope-bearing-PLAN stop · the accepted-PLAN declaration-site filter · cycle / broken-chain handling · the candidate-copy veto (resolution-by-construction) · the §2 segment-prefix match/subset/narrowing · the §5/§6 dispositions · m-2's submit suppliability guard (`scope-self-declared`) · m-2's declaration-time normalization/malformed-entry reject.

## 10. Exit fixtures (co-owned; aligned with the m-2 leg's three NFs)

- **`NF-scope-self-declared`** (REQUIRED — the kickoff self-widen NF): an accepted PLAN ancestor declares `[pkg/a]`; a **work record** supplies its own `[pkg/a, pkg/b]` → **(i)** the copy rejected `scope-self-declared` at **submit** (m-2's guard, when landed); **(ii)** the predicate resolves the **ancestor** rows, so `pkg/b/y.go` is out-of-scope → **`scope-exceeded` · `diff-drift`** — the self-supplied wider set **never takes effect**. *(Leg (ii) item-10-gated.)*
- **`NF-scope-exceeded-widen`:** a nested accepted PLAN declaring rows exceeding its ancestor → `scope-exceeded` · `plan-widens-ancestor` · `fail` / `rejected`.
- **`NF-scope-exceeded-drift`:** attributable `diff_paths ⊄ resolved_scope` → `scope-exceeded` · `diff-drift` @E1 *(item-10-gated)*.
- **`NF-scope-no-ancestor`:** no scope-bearing accepted PLAN in the chain → `degraded`/`self_reported`/**E0** — **asserts nothing, does not pass** (the false-positive guard); authority-class → `held`.
- **`NF-scope-ambiguous`:** two non-orderable scope-bearing PLANs at the frontier → `blocked` / `MachineryFault:true` → `held`/`rejected` by authority.
- **`FX-scope-hint-ignored`:** `parent_hint` points at a wider-scope record; the true `parent_edge` at the narrow one → the walk resolves the **narrow** value (the m-1 F-1 proof).
- **`FX-scope-iph-redaction`:** no raw candidate path and no resolved ancestor row appears in any verdict / bounce / row / `failing_detail` / `degradation_notes` on any refusal.
- *(m-2-owned, declaration-time:* `NF-scope-malformed-entry` — `..` / leading `/` / empty ⇒ typed reject.*)*

## 11. Lock discipline

**No lock moves.** c2 §5, s8 §13, and the byte-exact `{accepted, rejected, held}` terminal enum are unchanged. `scope-self-declared` (a submit `Violation.Class`) and `scope-exceeded` + its bounded reason enum, plus `scope-unbounded-no-declaration` / `scope-lhs-unattributable` / `scope-source-ambiguous` / `scope-lineage-unresolved`, are **`unsafe`/failure-detail-class strings under the existing §6/§12g family — not new terminals.**

## Fold-log
- **rev0 (…-133000):** initial four-pin predicate fill.
- **rev1 (…-152000):** folded m-3.implementer MR-1..MR-4 (glob grammar + two-class + I-PH + layer-1-honesty).
- **rev2 (20260713-1715) — the SETTLED CONTRACT.** Folded (a) master's RECONCILE `…-160510`: **MR-1 Option A** — declaration site narrowed to **accepted-PLAN-ancestor-only** (master owned the "/dispatch" over-specification; grant site = a non-welded future door); **segment-prefix grammar blessed** (globs withdrawn — this also dissolved the glob-containment finding, no algorithm ceiling needed); **observe narrowing-locus**; **MR-3 honesty rail confirmed**. (b) **m-2's convergence CO-SIGN `…-153200`**: the token-granularity fixed at **TWO layer-split tokens** — `scope-self-declared` (submit, m-2) + **`scope-exceeded`** (observe, m-3) **merging `plan-widens-ancestor` + `diff-drift`** under one invariant with a path-free bounded reason. *m-2 diverged from my one-token recommendation and was right: a submit `Violation` and an observe disposition are different signal shapes, and a one-token observe reason for the submit case could never fire.* (c) m-3.implementer `…-154000` MR-3: **the walk / stop / filter / cycle / candidate-veto are RED-first BUILD obligations — only the parent edge is extant** (§1.4/§9). Both halves now agree byte-for-byte; the item-10 carve-out and the E0 no-silent-green floor are intact.
