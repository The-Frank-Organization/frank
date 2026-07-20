# DESIGN (m-2 leg) — `scope_paths` declaration-slot grammar: the m-2 section of the `s9-scopepaths-cosign` four-pin co-sign

**DESIGN_DOC_ID:** `s9-scopepaths-cosign` (m-2 leg — the declaration-slot half; m-3 authors the predicate half; both seats sign one settled contract)
**Owner:** m-2.planner (Forms & Determinism) · **Consumes:** m-3 predicate (segment-prefix ⊆, resolution walk, E-rung), m-1 fidelity (channel/lineage key)
**Base:** `s10-close@39474d0`, registry `s10-fieldspec-v8`. `scope_paths` is currently **STRUCK** (not a registry field) per m-3-F7.
**Settled by master RECONCILE `s9-scopepaths-cosign/RECONCILE-orchestrator-planner-20260713-160510` + joint converge relay `s9-scopepaths-converge`.**

> **rev3 fold-log — the SETTLED contract (master RECONCILE-160510 rulings + the m-2 grammar-granularity co-sign):**
> - **MR-1 = Option A, PLAN-ONLY (master-ruled, pin narrowed as a *correction*).** `scope_paths` is a **PLAN property**; a grant-dispatch declaring a different scope = a second source of scope truth (duplicate-axis smell). Declaration site = **accepted PLAN ancestor only.** My rev2 `role_in:[planner,orchestrator-planner]` home is superseded → **`visible_when: {phase_in:["PLAN"]}`** (render-effective; PLAN is render-seeded). The grant-bearing site is a **non-welded future door** (Rail A openable, m-2 grammar + m-7 render extension), NOT an s9 carry.
> - **Grammar = normalized segment-prefix (master-BLESSED — my leg; m-3 withdrew their glob).** Reverts my rev2 glob-adoption back to segment-prefix (`pkg/a` matches `pkg/a`, `pkg/a/x.go`; not `pkg/ab`). Trivially decidable — dissolves the containment-complexity concern.
> - **Narrowing-locus = OBSERVE (master-steered, m-3 accepted).** Reverts my rev2 declaration-time narrowing → the ⊆ predicate fires in the observe layer. The submit-time suppliability guard remains defense-in-depth (RED-first build, not extant).
> - **Class-model = TWO tokens, layer-split (my grammar-granularity call, co-signed in `s9-scopepaths-converge`):** `scope-self-declared` (submit suppliability `Violation`, condition a) + `scope-exceeded` (observe ⊆ veto, conditions b+c under master's ONE invariant). Reasoning: signal-shape determinism — a submit `Violation.Class` and an observe predicate disposition are structurally different fields; one token cannot span both coherently (condition a is caught at submit / ignored at observe, never an observe veto). See §4.
> - **MR-3 honesty rail (master-CONFIRMED):** the walk · nearest-scope-bearing-PLAN stop · declaration-site filter · cycle/broken-chain handling · candidate-copy veto · the submit suppliability guard are **RED-first s9 BUILD obligations**; only the parent edge is extant at `39474d0`.

---

## 0. What is being contracted — and what it does NOT un-strike

`scope_paths` is the **PLAN-declared path-scope** the m-3 IMPL done-predicate reads for `diff_paths ⊆ scope_paths`. This leg supplies the m-2 (home + PLAN-only render + declaration-time segment-prefix normalization) half.

**Two honest limits (retained through every revision, master-confirmed):**
1. **This co-sign fixes the RHS *contract*, not a live check.** The ⊆ *evaluation* stays STRUCK until **design item 10 (turn-baseline)** makes `diff_paths` attributable. This co-sign moves the clause from *struck (no home)* to *home-defined-but-unevaluable ⇒ honest degrade*; a live E1 predicate needs BOTH this co-sign AND item 10.
2. **Nothing here is a v8 property.** The declaration slot, the segment-prefix normalizer/validator, the ancestor-only `resolve_scope` walk, the render-legality suppliability guard (`scope-self-declared`), and the observe ⊆ veto (`scope-exceeded`) are **all RED-first s9 build obligations.** v8 has no `scope_paths` field and no scope resolver (`lineage.parents()` is one-level, `lineage.go:254-262`).

**Design principle (confusion-firewall):** the threat is a *confused* lane widening its own scope. The answer is **ancestor-only resolution + two layer-pure typed refusals + a total, trivially-decidable segment-prefix relation** — no glob metacharacter surface, no cryptographic binding.

---

## 1. The declaration slot (m-2 home — the FieldSpec byte)

```json
{"id": "scope_paths", "layer": "header", "owner": "agent_enum_pick", "type": "row_array",
 "gate_referenceable": false, "fill_constraints": "free_text", "lineage_role": "none",
 "consumers": ["observe_gate"],
 "visible_when": {"all_of": [{"phase_in": ["PLAN"]}]},
 "annotation": "PLAN-declared path-scope: the paths a governing accepted PLAN authorizes its lane to touch. Read by the m-3 IMPL done-predicate (diff_paths subset-of scope_paths, segment-prefix) — an OBSERVE-layer input, never a form gate. Each row is one lane-root-relative normalized path prefix (POSIX '/', no '..', no leading '/', empty rejected; segment-boundary match: pkg/a matches pkg/a and pkg/a/x.go, not pkg/ab). m-2 normalizes/validates each row at declaration; m-3 owns segment-prefix match/subset/narrowing. CANONICAL VALUE = the nearest scope-bearing accepted PLAN ancestor's declaration, resolved by the conductor's resolve_scope walk over PARENT_DISPATCH_ID lineage; a non-PLAN/work record NEVER hosts its own canonical scope. Declaration site = accepted PLAN ONLY (master RECONCILE-160510 MR-1 Option A); a grant-bearing declaration site is a non-welded Rail-A future door, not an s9 carry. BUILD OBLIGATIONS (none extant in v8): (1) declaration-time segment-prefix normalize/validate; (2) the resolve_scope ancestor-only walk; (3) the render-legality suppliability guard (submit Violation scope-self-declared, non-PLAN supply); (4) the observe subset veto (scope-exceeded). Until each lands with its RED-first fixture, it is not a schema property."}
```

Rationale, grounded at source:
- **`owner: agent_enum_pick`** — the **planner declares** it at a PLAN; lane-authored, not a conductor projection. (Deliberately NOT `owner: system` — *why* the submit guard cannot lean on `systemOwnedHeader`'s reject at `validate.go:35`; §4.)
- **`type: row_array`, single column `path`** — the `SCOPE_DIFF`/`FOLD_SCOPE` precedent (`registry.json:123,125`). Per my s7a R2-COLGRAIN finding, `row_array` columns carry **no** FieldSpec enum enforcement — so segment-prefix well-formedness is a **declaration-time validator** (§3, build task), not an `enum_set`.
- **`gate_referenceable: false`** — observe-layer input, never a §5 form gate (as `executable_claims`, `registry.json:151`).
- **`lineage_role: none`** — resolved *through* lineage, not itself an edge; rides `PARENT_DISPATCH_ID` `parent_edge` (`registry.json:96`).
- **`visible_when: {phase_in:["PLAN"]}` — render-effective, PLAN-only (master MR-1 Option A).** Renders on PLAN forms (planners declare), absent on IMPL work forms. **Verified render-effective:** `render.go:51-54` seeds the visibility eval with `PHASE`, so `phase_in` is decidable at render. With the declaration site narrowed to PLAN, the grant-dispatch-IMPL surface (which drove my rev2 `role_in`) is out of scope — `phase_in:[PLAN]` is now exactly correct. **Caveat (retained honesty rail):** render-scoping is render-absence, not submit-rejection — the enforced legality is §4's guard.

## 2. Pin (a) — canonical value = the accepted PLAN ancestor (never self-declared)

**m-2 home side:** the render-legal declarer is a PLAN-phase form (§1). The **canonical value is the declaration on the nearest scope-bearing accepted PLAN ancestor** — never the work-record's own copy. **The enforced legal-declaration-site test (the render-legality suppliability guard, §4, a build task) admits a `scope_paths` only on an accepted PLAN record;** a non-PLAN record (including a work record) supplying one ⇒ `scope-self-declared` submit reject. Faithful to the accepted-parent precedent (`checkParentSubstrate`, `lineage.go:179-200`). **Grant-bearing declaration** (a dispatch re-declaring a narrower scope) is a **non-welded future door** — Rail A keeps it openable via an m-2 grammar + m-7 render-context extension, but it is **not an s9 carry** (master MR-1: scope belongs in PLANs; a second scope source is the duplicate-axis smell).

**m-3 seam:** the predicate reads the RHS *exclusively* from the resolved PLAN ancestor, never from `Candidate.Record.Headers["scope_paths"]`.

## 3. Pin (b) — resolution through lineage (walk + stop) — and the segment-prefix grammar

**m-3/conductor side:** `resolve_scope(Candidate)` walks upward via `PARENT_DISPATCH_ID` (`e.parents()`, `lineage.go:254-262`), climbing accepted ancestors, **stopping at the nearest accepted PLAN ancestor bearing a `scope_paths` declaration** (innermost-wins); reaching the dispatch root with none ⇒ the (d) ∅-declared disposition. **Not in v8** (`parents()` is one-level) — a RED-first s9 build task.

**m-2 home side (what makes the walk well-defined):** at most one `scope_paths` per record (a `row_array` is a single field) ⇒ no single record self-conflicts; the render-legal declarer is a PLAN ⇒ a well-formed target chain.

**The grammar — normalized segment-prefix (master-BLESSED; owner split explicit):**
- **Entry form:** each `path` row is a **lane-root-relative path prefix**, normalized: POSIX `/`, **no `..`**, **no leading `/`**, no empty entry. **Segment-boundary match** (`pkg/a` matches `pkg/a` and `pkg/a/x.go`, NOT `pkg/ab`). `inScope(p) := ∃ row s : segmentPrefix(s, p)`; `⊆ := ∀ p ∈ diff_paths : inScope(p)` — m-3-owned, trivially decidable.
- **Validation loci (agreed, neither half assigns the other's byte):**
  - **m-2 owns declaration-time normalization + well-formedness** — a named s9 build task (the normalizer/validator) with RED-first fixture `NF-scope-malformed-entry` (a row with `..` / leading `/` / empty ⇒ typed reject at declaration). m-2 guarantees every stored row is canonical **at rest**.
  - **m-3 owns segment-prefix match/subset + nested-narrowing** over the canonical rows.

**m-1 fidelity leg:** the walk rides **only** `PARENT_DISPATCH_ID` `parent_edge` (`registry.json:96`, `system_only`; lane supply → `system-owned` reject, `validate.go:35`; conductor-stamped via `stampParent`, `engine/parent_stamp.go`). The lane-suppliable `parent_hint` (`:97`, `lineage_role:none`) is NOT authoritative — gated by computed `parent_hint_honored` (`:98`). **m-1 to confirm** (m-3 §F) the walk consumes the `parents()`/`parent_edge` chain only; `parent_hint` never substitutes.

## 4. Pin (c) — the typed refusals (TWO tokens, layer-split; my grammar-granularity call)

**Master's ONE invariant** (RECONCILE-160510): *a record's effective scope may not exceed its authorizing accepted-PLAN-ancestor's declared scope (segment-prefix ⊆).* Master steered that whether this surfaces as one token+reason or two tokens is the **m-2 grammar-granularity call.** My call: **TWO tokens, split by layer/mechanism** — because a submit `Violation` and an observe predicate disposition are **structurally different signal shapes** (different fields, different lifecycle stages, different `MachineryFault`/terminal semantics), and one token cannot span both without becoming a shape-ambiguous class (a determinism defect — my domain). The two mechanically-distinct events:

1. **`scope-self-declared` — submit-layer suppliability `Violation` (m-2/m-1 guard, condition a).** A non-PLAN record (a work record, or any non-PLAN-ancestor) supplying its own `scope_paths` ⇒ a **submit `Violation{Class:"scope-self-declared"}`** via the render-legality suppliability guard. **`MachineryFault` n/a** (a submit violation, not an observe terminal). Caught at submit — it **never reaches observe**; if the guard has not yet landed, resolution-by-construction simply **ignores** the candidate-borne copy (no observe veto for it — the honesty rail). **Honest classification (master-confirmed MR-3):** this guard is **not** in v8 — `Validate` rejects a supplied value generically only when `systemOwnedHeader(spec)` (`validate.go:35`), and `scope_paths` is `agent_enum_pick`; `seat-scope` fires only for `enum|bool` (`:55-62`), not a `row_array`'s presence; `visible_when` is render-absence (`render.go:96-108`). **Named joint s9 build task** (m-2 marker + m-1/m-2 validator), RED-first; no non-lane-writability claim before it lands.

2. **`scope-exceeded` — observe-layer ⊆ veto (m-3, conditions b + c under the ONE invariant).** In the **observe layer** (master's narrowing-locus ruling), the ⊆ predicate fires: the record's effective scope exceeds its resolved PLAN-ancestor's declared scope. This **merges** (b) a nested PLAN declaring a scope not ⊆ its ancestor and (c) `diff_paths ⊄ resolved_scope` (real drift) — both are the same invariant over different operands (declared vs enacted "effective scope"). ⇒ **observed-false / `MachineryFault:false` / terminal `rejected` both authority classes** (m-3 §D). The (b)-vs-(c) distinction, if surfaced, is a **path-free bounded internal reason** `{plan-widens-ancestor, diff-drift}` — never the paths (I-PH). Drift leg (c) is **item-10-gated**. Never widens `resolved_scope`.

*(Why not one-token+reason [m-3's rec]: condition (a) is emitted at SUBMIT as a `Violation.Class`, while (b)/(c) are emitted at OBSERVE as predicate dispositions with `MachineryFault`/terminal-by-authority. A single token spanning both value-spaces is structurally incoherent — (a) never appears at observe (caught/ignored earlier), so an observe `scope-exceeded/self-declared` reason could never actually fire. Two layer-pure tokens is the deterministic model; the confusion-firewall's "loud but not proliferating" is satisfied at TWO — one per genuine confusion: structural self-certification vs substantive over-reach.)*

**RED-first fixtures (VP exit-set; I-PH-redacted):**
- **`NF-scope-self-declared`** — a non-PLAN/work record supplies its own `scope_paths` ⇒ **`scope-self-declared` submit `Violation`** (once the guard lands). *(The pin-c work-record leg.)*
- **`NF-scope-exceeded-widen`** — a nested PLAN declares a scope not ⊆ its ancestor ⇒ observe **`scope-exceeded`** (reason `plan-widens-ancestor`); canonical bound unchanged.
- **`NF-scope-exceeded-drift`** — resolved PLAN scope + `diff_paths ⊄` it ⇒ observe **`scope-exceeded`** (reason `diff-drift`); **item-10-gated** (registered now, RED lands when the LHS exists).
- **I-PH on all:** the surface (verdict/bounce/row/`failing_detail`) carries **only** the symbolic class + a non-path offending-row **count**; raw paths, resolved bound, and prefixes do **not** appear (m-3 §H / `FX-scope-iph-redaction`).

## 5. Pin (d) — missing / ambiguous-source disposition (Rail A floor; m-3 owns the §D table)

**m-2 floor:** a missing/unresolvable declaration must **never be a silent green.** The full per-branch bytes are **m-3's §D table** (adopted, re-expressed on segment-prefix):
- **∅-declared** (walk reaches dispatch-root, none): `degraded` / `MachineryFault:false` / no-vantage → authority `held`, non-authority `accepted`+`self_reported`, **E0** — degrade, never silent green.
- **ambiguous / broken-chain**: `MachineryFault:true` → machinery-fault edge (authority `held`, non-authority `rejected`) — never an arbitrary pick.
- **drift / PLAN-widens**: observed-false → `rejected` both — **item-10-gated** for drift.

My constraint is only the no-silent-green floor + one `scope_paths` per record (a single record cannot self-conflict). Block-vs-label and the exact `MachineryFault`/terminal bytes are m-3's.

---

## 6. Boundary contract (SETTLED owner split)

| Concern | m-2 (this leg) | m-3 (predicate leg) | Joint / build-task |
|---|---|---|---|
| slot home/owner/type/`visible_when`(PLAN-only)/`gate_referenceable`/`lineage_role` | **owns** | consumes | — |
| segment-prefix grammar | declaration-time **normalize/validate** (`NF-scope-malformed-entry`) | **owns** match/subset/narrowing | one language, §3 |
| canonical = accepted PLAN ancestor (pin a) | PLAN-only render + PLAN-only guard test | **owns** RHS-only-from-ancestor | co-sign |
| `resolve_scope` walk + nearest-PLAN stop (pin b) | well-formed target chain | **owns** the walk (BUILD TASK — not v8) | co-sign; m-1 fidelity |
| `scope-self-declared` (submit suppliability, cond a) | **co-owns** registry marker + validator | — | BUILD TASK (m-2/m-1) |
| `scope-exceeded` (observe ⊆ veto, cond b+c) | — | **owns** the observe check + §D bytes | co-sign |
| missing/ambiguous disposition (pin d) | no-silent-green floor | **owns** §D block-vs-label + `MachineryFault` | co-sign |
| I-PH path redaction of surfaces | fixture assertions (class + count only) | **owns** §H output contract | co-sign |
| grant-bearing declaration site | Rail-A future door (m-2 grammar) | — | NON-WELDED (not s9) |

## 7. Out of scope / anti-half-fix guards

- **No v8-property claims:** the segment-prefix validator, `resolve_scope` walk, render-legality guard (`scope-self-declared`), and observe ⊆ veto (`scope-exceeded`) are all RED-first s9 build tasks (§0.2).
- `scope_paths` is **not** `gate_referenceable` — never a §5 form-gate key.
- The segment-prefix match/subset/narrowing, per-branch `MachineryFault`/terminal bytes, and I-PH redaction remain **m-3's**.
- `diff_paths ⊆ scope_paths` **evaluation** stays STRUCK until BOTH this co-sign AND item 10 land (§0.1).
- **Grant-bearing declaration site = non-welded future door** (master MR-1) — not an s9 carry; opening it is a future m-2 grammar + m-7 render-context amendment.
- No lock moves: c2 §5, s8 §13, `{accepted, rejected, held}`, and the s8 I-PH contract are unchanged; `scope-self-declared`/`scope-exceeded`/`scope-source-ambiguous` are `unsafe`-class detail strings under the existing §6/§12g family, not new terminals.

## 8. Open items — SETTLED (co-sign)

1. **m-3:** rev3 co-signs the two-token layer-split (`scope-self-declared` submit + `scope-exceeded` observe, reason `{plan-widens-ancestor, diff-drift}`); confirm your §D signal bytes bind to it (they do per your converge relay). Fold the settled byte into the co-sign doc.
2. **m-1:** the pin-b fidelity confirm (§3 / m-3 §F).
3. **m-2.implementer:** rev3 records the master-settled contract (segment-prefix, PLAN-only, observe-locus) — the contested rev2 axes are master-ruled, not re-decided by me; the one m-2 design byte is the §4 two-token granularity. Visible for confirmation; master's sequence puts the co-sign re-review on m-3.implementer.
