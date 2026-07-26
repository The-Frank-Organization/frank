# Stage-6 §5-E — the m-2 logical schema/description COMPONENT contract (`logical_tool_schemas[]` / `tool_descriptions[]`)

**DESIGN_DOC_ID:** `step3-relock-dag-m2` · **rev2** — the current revision; history in §7 (the authoritative log; this marker moves with every revision).
**Dispatch:** `master/relays/step3-relock-dag-m2/DESIGN-orchestrator-planner-20260721-235604.md` @ `342f64b6…` (rev2, RELEASED by `…/RELEASE-orchestrator-planner-20260722-004001.md`).
**Authority basis:** the ratified stage-6 amendment `master/STEP-3-STAGE6-AMENDMENT.md` §5-E + §6-E (rev12, SHA-256 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183` — verified this session).
**Governed ADDITIVE delta over the frozen m-2 final `83d8e63e…` (F73):** this doc adds a new contract; it edits NO byte of the frozen stage-1 design (whose byte-bound approve stands) and NO FieldSpec registry byte.
**Owner:** m-2.planner authors; m-2.implementer pair-reviews the final bytes; consumers = m-9 (folds the component into `logical_surface_digest`) + m-3 (independent observer derivation).

---

## §1 — What this component IS (and the non-conflation invariant)

m-9 owns `logical_surface_digest` = SHA-256 over JCS `{instructions, logical_tool_schemas[], tool_descriptions[], compaction_template, policy_messages}` — the **pre-lowering surface the model actually sees** at an assembly, riding the m-9 → m-10 attempt row / E0. This doc supplies the **component contract** for the two m-2-supplied members.

**INV-E1 (non-conflation — the load-bearing invariant):** this component and the F58 build identity are DIFFERENT artifacts answering different questions, deliberately:

| | F58 build identity (frozen stage-1 §3.2) | §5-E logical component (this doc) |
|---|---|---|
| question | *which generator/surface was BUILT?* | *what surface was PRESENTED to the model at this assembly?* |
| schema input | **pinned canonical TEMPLATE bytes** (headers slot empty, const removed — exact byte strings, frozen §3.2) | the **live presented schema as a PARSED JSON VALUE** (rendered fill, `form_digest` const, volatile option/annotation state as-of-assembly — members/values preserved, JCS-canonicalized per §2.3; **original serialization bytes are never inputs**) |
| descriptions | **EXCLUDED** (presentation, frontend-variable) | **INCLUDED** (presentation is exactly what is being bound) |
| stability | run-immutable; serve-gate verified | **attempt-grain**; legitimately varies between assemblies |
| consumer | m-10 serve gate / F63 release-binding | m-9 digest fold → attempt row / E0; m-3 observer derivation |

Neither substitutes for the other. A logical-component change across attempts is **honesty, not drift** (see §4). The two are LINKED by generator identity — by adjacency, per §3.3.

## §2 — The component recipe (normative)

### §2.1 Element shape + sourcing

- **`logical_tool_schemas[]`** — exactly one element per tool: `{"name": <canonical tool id>, "schema": <the presented schema as a PARSED JSON VALUE — the presented members and values preserved exactly, JCS-canonicalized per §2.3; the frontend's original serialization bytes are never inputs>}`.
- **`tool_descriptions[]`** — exactly one element per tool: `{"name": <canonical tool id>, "description": <the TOOL-LEVEL description STRING VALUE m-9's worker actually presents — the string value exactly, JCS-serialized per §2.3, never any particular escaping of it>}` (always a string; a tool presented with no tool-level description contributes `""`). **NO synthesis (REVISED at review-r1 F1):** the element carries only what was presented at the tool level — nothing is copied, derived, or promoted into it from anywhere else. In particular, **R-3's volatile-options annotation is a FIELD-PROPERTY `description` inside the schema object** (frozen m-2 R-3) — it is already bound, once and naturally, through `logical_tool_schemas[].schema`; an implementer must never populate a tool-level description from schema-property annotations (that would hash a surface the model never saw at that level). Tool-level description strings — the honesty banner included — are presentation surface whose governance home the frozen m-2 contract itself assigns to the m-9 catalog/m-3 surface, not to m-2.
- **Sourcing/ownership (exact, re-cut at review-r1 F1 — no re-ownership):** **m-2 owns** the component RECIPE (both arrays: shape, totality, ordering, serialization) and **produces the three relay-verb `schema` objects** (the live render at assembly time: `SchemaFromForm(form, digest)` for `relay.submit` — the R-rule render including R-5's const and the volatile schema-property annotation state; `ProjectSchema()`/`ReadSchema()` for the other two). **m-9 owns** the five local `schema` objects (their §8.3 pins are the presented schemas by construction) **and ALL EIGHT tool-level `description` strings** (whatever its worker presents, relay verbs included); **m-9 assembles and hashes** (no owner hashes bytes it cannot see — §5-E verbatim, now without exception: m-2 contributes schema objects and the recipe, never presentation strings it does not present).

### §2.2 Totality + ordering (fail-closed)

- Each array carries **exactly the ratified 8 canonical names** (`apply_patch, bash, edit, read, relay.project, relay.read, relay.submit, write`), **sorted by `name` byte order**. A missing name, an extra name, a duplicate, or a cross-array name-set mismatch ⇒ the assembly REFUSES (no digest is produced over a partial/aliased surface); aliases normalize to canonical IDs before the set check (the §3.1 table of the frozen stage-1 contract is the reference).
- Element members are exactly the two named per array — no additional members (the ratified five-member outer object is CLOSED; so are these elements — Rail A: an extra member silently folded into a digest would bind bytes no contract names).

### §2.3 Serialization (JCS-stable)

- **JCS (RFC 8785) over the PARSED presented JSON values is the SOLE logical-component encoding (REVISED at review-r1 F3).** The component members serialize under JCS as part of m-9's outer object — schema objects JCS-canonicalized like everything else (member sort, RFC 8785 §3.2.2.2 string serialization); no separate pre-hash of my component exists (the outer digest is the only digest; the LOCK binds the recipe, not a value).
- **"Verbatim" means the presented MEMBERS and JSON VALUES, not the frontend's original serialization bytes:** the recipe's input is the parsed presented surface; equivalent input serializations of the same values MUST converge on the same component bytes (a build fixture asserts convergence across differently-serialized equivalent inputs). No runtime path may depend on any non-JCS encoding reproducing JCS bytes.
- **The coincidence note, corrected and narrowly scoped (the rev0 claim was overbroad — F3):** rev0 claimed Go `encoding/json` bytes equal JCS bytes for these documents because they are ASCII without numeric literals. That premise is insufficient: Go's default marshaler HTML-escapes `<`, `>`, `&` (and U+2028/U+2029), which RFC 8785 serializes as-is — all reachable in ASCII strings, and live form-derived options/defaults/annotations are open-grammar strings this contract explicitly binds. The coincidence therefore holds ONLY as a checked fact about **specifically pinned static fixture bytes** (the §8.3 local pins and the frozen §3.2 templates, whose pinned literals contain none of the divergent characters — asserted by fixture, never extended), with **negative coverage**: a fixture string containing `<`/`>`/`&` must show the divergence and show the logical digest following JCS. The frozen F58 digests are untouched by this (they are over pinned literal bytes and self-consistent); this component simply never inherits their encoding.

## §3 — Bindings

### §3.1 Determinism claim (exact)

Given one presented surface — the ordered set of (name, schema, description) triples actually shown to the model at one assembly — the component bytes are a **pure function** of it (sort → shape → JCS). Same presented surface ⇒ same bytes, on any implementation, in any process. No hidden state, no clock, no module-internal knowledge is required to construct them.

### §3.2 Observer derivation (m-3's independence — §5-E "derives each component independently")

The observer reconstructs the component **from the worker's presented surface alone** by the same pure function: extract the triples from the presented tool list → §2.2 sort/totality → §2.1 shape → §2.3 JCS. The recipe is fully specified in this doc; the observer needs NO m-2 module code (it MAY reuse the offered helper, §5 — independence is preserved because the recipe, not the implementation, is the contract). The `form_digest` const inside the presented `relay.submit` schema additionally makes the presented-surface ↔ conductor-form linkage observable at the same grain, for free.

### §3.3 The versioning/mapping-version binding (by ADJACENCY, not payload)

The ratified outer object is a **closed five-member set** — adding a version member inside it is not mine to do and would move m-9's ratified field set. The generator binding therefore rides the **exact immutable relation that EXISTS on the frozen surfaces (REVISED at review-r1 F2 — rev0 asserted same-row catalog carriage that does not exist; verified this pass at the frozen bytes):**

- the attempt is identified by `{attempt_id, turn_id}` in m-10's `provider_attempts` (which carries **no** catalog fields — m-10 r40 §E);
- the attempt's **run relation** resolves through `turns → runs` to the run's **immutable frozen run manifest** (`runs` carries the canonical manifest bytes + `run_manifest_digest`; m-3's E0 events carry `run_manifest_digest`, and m-3's E3 acquisition obtains `tool_catalog_digest` through the named run's manifest → release-binding relation — m-3 r4, verbatim its scope table);
- that manifest's **`tool_set`** carries the F58 vector rows — the relay rows binding **`m2-mapping-v1` + the three template digests** — and its **`tool_catalog_digest`** identifies that vector, mechanically verified against the shipped registry at the F63 release-binding.

**The binding, exactly:** which generator produced the presented relay-verb schemas in a given `logical_surface_digest` is answered by resolving the attempt's run manifest and reading its relay rows — the presented schemas are those TEMPLATES filled by that SAME generator version, and drift without a version bump is already caught by the frozen F63 machinery. **Fail-closed semantics (normative):** if the run relation is unresolvable, the manifest absent/undigested, the `tool_set` relay rows missing `m2-mapping-v1`, or `tool_catalog_digest` mismatched against the locked vector — the generator-binding question is UNANSWERABLE and any claim depending on it **fails closed** (E3: non-applicable; no default, no inference). **No new version member, no ratified-field-set change, no double-binding — and no carrier is asserted that the frozen surfaces do not provide.** (Should any consumer ever require direct same-row co-residence, that is an m-10/m-3 carrier amendment routed through master — never a claim here.)

## §4 — Attempt-grain honesty (the volatility statement)

The presented `relay.submit` schema is run-varying **by construction** (the frozen stage-1 §2.3.3 grounds: `form_digest` const; `ConductorVolatile`/`DigestExempt` options move without a digest move; re-render/F-rule refreshes between attempts). Therefore `logical_surface_digest` **legitimately differs across assemblies exactly when the presented surface changed** — that is the digest doing its job (binding what the model saw THEN), not drift. No cross-attempt stability is claimed or wanted; per-assembly determinism (§3.1) is the whole claim. Anyone needing run-stable identity uses the F58 vector (INV-E1) — reachable from any attempt through the §3.3 run-manifest relation, answering its different question from its own carrier.

## §5 — The offered assembler helper (offered, not imposed — the `ValidateSubmitArguments` posture)

The m-2 module ships `LogicalToolComponent(surface []PresentedTool) (schemas, descriptions []Element, error)` implementing §2 (totality refusal included). m-9 MAY consume it for the fold; m-3 MAY reuse it at the observer; either may implement to the recipe instead — the §6 fixtures bind the BYTES, not the implementation.

## §6 — Build obligations (RED-first; nothing here exists at `frank@c78da38`)

1. The assembler helper (§5) + totality refusal tests (missing/extra/duplicate/cross-array-mismatch ⇒ refuse, no digest).
2. Determinism fixture: one pinned presented surface ⇒ byte-exact component bytes, cross-checked against an independent hand-serialization; **plus the serialization-convergence leg (F3): equivalent inputs serialized differently ⇒ the same component bytes**.
3. **The scoped coincidence + divergence pair (F3):** the pinned static fixture bytes (§8.3 pins, §3.2 templates) assert the checked coincidence; a fixture string containing `<`/`>`/`&` asserts the Go-default/JCS divergence AND that the logical digest follows JCS.
4. Volatility-visibility fixture: two assemblies across a re-render (volatile option moved) ⇒ different presented schema ⇒ different `logical_surface_digest`; same surface ⇒ same digest; **plus the F1 placement leg: the volatile change moves the schema-property annotation while the TOOL-LEVEL description stays unchanged — the digest moves solely through the schema member**.
5. Observer-parity fixture: an m-3-side reconstruction from the presented surface == the m-9-side component bytes (the §3.2 independence, executed).
6. **The manifest-join fixture (F2, rewritten):** resolve attempt → run → `run_manifest_digest` → the frozen manifest; assert the relay rows carry `m2-mapping-v1` and the presented relay schemas are the shipped generator's render of those templates; **fail-closed legs**: unresolvable run relation · absent/undigested manifest · relay rows missing the mapping version · `tool_catalog_digest` mismatched — each ⇒ the generator-binding claim is refused (E3 non-applicable), never defaulted.

## §7 — Revision log

- **rev2** (2026-07-22, m-2.planner): the single `step3-relock-e-m2-review-r2` blocker folded — **M2-E-R2-F1**: the load-bearing INV-E1 table row and both §2.1 element definitions still carried byte-verbatim wording ("schema bytes" / "verbatim") that the F3-corrected §2.3 supersedes — a builder following the table could hash original frontend serialization bytes and fork digests over semantically identical schemas (the same stale-summary class as review-r1's post-fold sweep catch; the value/byte distinction evidently needs stating at EVERY definition site, not once). Fix: the table row now distinguishes F58's pinned canonical TEMPLATE bytes from the logical side's PARSED JSON VALUE (members/values preserved, JCS-canonicalized, original serialization bytes never inputs); both §2.1 elements restated value-level (the description element too — a string value can also be escaped multiple ways, the same ambiguity one level down); the full-doc sweep re-run — remaining "verbatim" hits are citation usage (§5-E quote; m-3 scope-table reference) and the §7 history only. No recipe, fixture, frozen artifact, or sibling-owner mechanism moved.
- **rev1** (2026-07-22, m-2.planner): the three `step3-relock-e-m2-review-r1` blockers folded, each verified before folding — F1 and F3 against my own frozen bytes, F2 at the frozen m-10 r40 / m-3 r4 surfaces I had failed to read before asserting (owned: rev0's "same attempt row" carriage was assert-before-verify).
  **F1** — `tool_descriptions[]` re-cut to the exact m-9-presented TOOL-LEVEL strings with NO synthesis; the R-3 annotation stays where the frozen contract puts it (a schema-property member, bound once through `logical_tool_schemas[].schema`); ownership re-split (m-2: recipe + relay schema objects; m-9: all eight tool-level descriptions + five local schemas + assembly/hash); the placement fixture added.
  **F2** — the adjacency carrier rewritten to the relation that EXISTS: attempt → run → `run_manifest_digest` → the frozen manifest whose `tool_set` relay rows carry `m2-mapping-v1` and whose `tool_catalog_digest` identifies the vector (verified: `provider_attempts` carries no catalog fields, m-10 r40 §E; E0 carries `run_manifest_digest`/`policy_digest`, E3 acquires the catalog via the manifest→release-binding relation, m-3 r4); fail-closed semantics for absent/mismatched/unresolvable facts; the same-row co-residence claim retracted (a carrier amendment would route through master, never be asserted here).
  **F3** — JCS over parsed presented values made the SOLE encoding; "verbatim" clarified to members/values, never original serialization bytes; the coincidence claim corrected (Go's default marshaler HTML-escapes `<`/`>`/`&` — reachable in ASCII, forgotten in rev0) and scoped to pinned static fixture bytes with negative coverage; no runtime reliance anywhere.
- **rev0** (2026-07-22, m-2.planner): initial component contract for pair review. Authored under the RELEASED rev2 dispatch `342f64b6…`; amendment rev12 hash verified `1125b0a0…`; frozen m-2 `83d8e63e…` untouched (additive delta, F73).
