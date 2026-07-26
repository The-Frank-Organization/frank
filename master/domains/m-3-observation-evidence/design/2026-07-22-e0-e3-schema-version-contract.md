# m-3 E0/E3 schema-version contract — `m3.app_event.v1→v2` + `m3.e3_observation.v1→v2` (r2; the ratification-bindable closed set)

**DESIGN_DOC_ID:** `step3-stage6-m3-schema-amendment`
**Owner:** m-3 — sole author; m-3.implementer pair-reviews the final bytes; master binds the pair-approved hash into the amendment (this artifact never self-ratifies).
**Dispatch:** `step3-stage6-m3-schema-amend/DESIGN-orchestrator-planner-20260722-130000` (items 1–7), route (b) per M3-VP-R2-F1.
**Basis (verified):** frozen m-3 r4 @ `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad` (v1 BYTE-FROZEN; governed additive contract, never an edit). D4 input now existing: pair-approved m-8 r5 @ `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21` (cited as the cut-matrix SOURCE for master's D4 fold; its per-cut evaluation stays in the parked lane-2 r1 — not authored here).
**Status:** r2 — SCHEMA-R1-F1 folded from DESIGN-REVIEW `…-171500` (r1 must-revise, one bounded textual blocker: the §8 wildcard-form token); pending final pair re-review. The cut-LIST evaluation, verdict machines, sink, and E join stay parked.

---

## §1 — The v1 reference census (hash-bound citation of frozen r4; v1 rules live THERE, byte-frozen)
**`m3.app_event.v1`** (r4 §2.2 @ `009df607…`) — exactly: `schema` · `event_kind` · `phase` · `scope` · `run_id` · `turn_id` · `attempt_id` · `turn_epoch` · `provider_lane_id` · `run_manifest_digest` · `policy_digest` · `deny_reason` (IFF `phase=denied`) · `event_evidence` · `event_integrity` · `reported_by` · `event_ts`.
**`m3.e3_observation.v1`** (r4 §3.2/§3.3 @ `009df607…`) — the universal/identity/vector census under the r4 six-scope matrix.
*(Per SCHEMA-R0-F1's allowance, v1 stays a hash-bound reference — the frozen bytes ARE its matrix. The V2 closed set is fully present below.)*

## §2 — The complete `m3.app_event.v2` closed field table (every field, every status — nothing imported)
| field | type/domain | status |
|---|---|---|
| `schema` | literal `m3.app_event.v2` | REQUIRED |
| `event_kind` | literal `provider_attempt` | REQUIRED |
| `phase` | enum `denied\|sent\|completed\|failed\|cancelled\|unknown` | REQUIRED |
| `scope` | literal `attempt` | REQUIRED |
| `run_id` · `turn_id` · `attempt_id` | string | REQUIRED |
| `turn_epoch` | canonical-decimal-uint64 STRING (`^(0\|[1-9][0-9]*)$`, < 2^64) | REQUIRED |
| `provider_lane_id` | string (m-8 canonical lane-ID form) | REQUIRED |
| `run_manifest_digest` | 64 lowercase hex | REQUIRED |
| `policy_digest` | 64 lowercase hex | REQUIRED |
| `deny_reason` | the m-3 §1.3a closed token enum | CONDITIONAL: present IFF `phase=denied` |
| `frozen_core_digest` | 64 lowercase hex | **BYTE-OPTIONAL (§4 layer 1): absence is schema-VALID; type-checked when present** |
| `logical_surface_digest` | 64 lowercase hex | REQUIRED (producer = m-9 first-hand; recipe m-9's, cited by reference; binding confirmation parked) |
| `event_evidence` | fixed literal `E0` | REQUIRED |
| `event_integrity` | fixed literal `self_reported` | REQUIRED |
| `reported_by` | string (a claim, never a proof) | REQUIRED |
| `event_ts` | RFC3339 | REQUIRED |
Any field not in this table ⇒ **malformed**. The v1 redaction rule (ids/digests/enums/timestamps only) applies verbatim.

## §3 — The complete `m3.e3_observation.v2` six-scope matrix (fully present; universal fields — `schema` (literal `m3.e3_observation.v2`) · `scope` · `claim` · `observed_outcome` · `observer_id` · `observation_ts` — REQUIRED at every scope and not repeated per row)
| scope | REQUIRED identity | REQUIRED vector | CONDITIONAL | FORBIDDEN |
|---|---|---|---|---|
| `build` | — (the release vector IS the identity) | `app_main_build_digest`+`m9_worker_build_digest`+`m8_build_digest` XOR `release_digest` · `tool_catalog_digest` | — | `run_id`, `turn_id`, `attempt_id`, `artifact_ref`, `relay_id`, `run_manifest_digest`, `policy_digest`, `provider_lane_id`, `frozen_core_digest` |
| `artifact` | `artifact_ref` (closed enum `app_main\|m9_worker\|m8_connector\|release\|tool_catalog\|policy`) | **exactly the one digest field `artifact_ref` names** (`app_main`→`app_main_build_digest` · `m9_worker`→`m9_worker_build_digest` · `m8_connector`→`m8_build_digest` · `release`→`release_digest` · `tool_catalog`→`tool_catalog_digest` · `policy`→`policy_digest`) | — | every other identity + vector field (incl. `release_digest` when `artifact_ref` ≠ `release`), `frozen_core_digest` |
| `run` | `run_id` | `run_manifest_digest` · `tool_catalog_digest` · builds XOR `release_digest` · `policy_digest` · `provider_lane_id` | — | `turn_id`, `attempt_id`, `artifact_ref`, `relay_id`, `frozen_core_digest` |
| `turn` | `run_id` + `turn_id` | same five as `run` | — | `attempt_id`, `artifact_ref`, `relay_id`, `frozen_core_digest` |
| `attempt` | `run_id` + `turn_id` + `attempt_id` | same five as `run` | **`frozen_core_digest`** — BYTE-OPTIONAL (§4 layer 1): absence schema-VALID; 64 lowercase hex when present | `artifact_ref`, `relay_id` |
| `relay_record` | `relay_id` + `run_id` | `run_manifest_digest` · `tool_catalog_digest` · builds XOR `release_digest` | — | `turn_id`, `attempt_id`, `artifact_ref`, `policy_digest`, `provider_lane_id`, `frozen_core_digest` |
Required ⇒ present; forbidden ⇒ the key does not appear (never null/empty); any violation or any field outside this census ⇒ **malformed**. `model_surface_digest` is NOT in this census (§6).

## §4 — The two-layer presence discipline for `frozen_core_digest` (SCHEMA-R0-F2 folded; the reviewer's named branch taken)
**Layer 1 — byte-only closed parsing (what every validator can decide from the record bytes alone):** `frozen_core_digest` is **syntactically OPTIONAL** at E0 (§2) and at E3 `attempt` scope (§3); when present it MUST be 64 lowercase hex; it is FORBIDDEN at every non-attempt E3 scope. **Absence is schema-VALID — never `malformed`.** No byte parser infers, or is permitted to infer, a pipeline position from `phase` or any other field (the same record bytes can carry `phase=failed` on both sides of freeze — the parser cannot and does not decide which).
**Layer 2 — producer/writer conformance (who must include it, and who checks THAT):** the E0 populator (m-9) and the E3 writer (the observer) MUST include the field **iff `FREEZE-REACHED(cut) = true` per the settled D4 cut matrix** — whose source is the pair-approved m-8 r5 (`c0b7b488…`; the per-cut evaluation rides the parked lane-2 r1 on master's D4 fold). **The enforcement locus for layer 2 is the D4/T4 conformance fixtures over the producers** — named actors checking named producers at build/test time; layer 2 is NEVER an acceptance-time byte check, and no acceptor/evaluator acquires cuts from any store (the r4 evaluator gains no acquisition step).
**The predicate coupling (consistent by construction):** at byte-valid absence, predicate 1's missing-input branch yields **`unknown`** — reachable exactly where layer 1 permits absence; a producer that OMITS where layer 2 required is a conformance failure caught by fixtures, and its record honestly evaluates `unknown`, never a fabricated pass/fail.

## §5 — Exact four-literal dispatch (SCHEMA-R0-F3 folded; no wildcard, no family matching)
The record's `schema` value is compared **byte-exact** against exactly four literals; each selects exactly one matrix and names its validators:
1. **`m3.app_event.v1`** → the frozen r4 §2.2 census/rules (@ `009df607…`) → validated by the m-10 applier at receipt + readers.
2. **`m3.app_event.v2`** → the §2 table → the same actors.
3. **`m3.e3_observation.v1`** → the frozen r4 §3.2/§3.3 matrix (@ `009df607…`) → the m-3 applicability evaluator, v1 branch.
4. **`m3.e3_observation.v2`** → the §3 matrix → the evaluator, v2 branch.
**Any other `schema` value ⇒ malformed.** There is no prefix, wildcard, family, or "starts with m3." matching — `m3.unrelated.v1` is malformed, full stop. Within a selected branch: any missing required field, any field outside the branch's census, any conditional-rule violation ⇒ malformed. No cross-version leniency or coercion exists.
**Historical v1 records (the corrected statement):** v1 bytes **never migrate, reparse, or reinterpret as v2** — and whenever a v1 record is read or evaluated, it **REVALIDATES in full under the frozen v1 rules** (canonical bytes, reference-digest check, closed-schema validity, observer provenance, applicability — the r4 machinery verbatim). "Frozen" means the RULES never change; it has never meant a record skips validation.

## §6 — `model_surface_digest`: the explicit exclusion (unchanged from r0, preserved by the review)
NOT part of the ratified v2 census for either carrier. It lands ONLY with the parked E-join work, as a subsequent governed delta (a v3 bump) keyed to the pair-approved m-9/m-8 producer recipes. A ratification over v2 is complete and decidable on the day it ratifies; bumps are cheap under §5's exact-literal dispatch (a fifth/sixth literal is one row each).

## §7 — Actors + accepted-version behavior (per master item 4; unchanged in substance from r0, restated against §5)
| role | actor | version behavior |
|---|---|---|
| E0 emitter | the m-9 worker (sole populator+carrier) | emits **v2 only** from the lane-2 build onward; no dual-emission |
| E0 acceptors | the m-10 applier (at receipt) + readers | accept the two exact E0 literals, each validated per §5 |
| E3 writer | the external observer (the r4 §3.6 registry, unchanged) | writes **v2 only** from ratification onward |
| E3 evaluator | the m-3 applicability rule, observer-side | §5 four-literal dispatch is THE added step; the run-constant acquisition/comparison vector, the acquire-then-compare order, the claim-context table, the observer-provenance step, and the mutation semantics are **byte-unchanged** |

## §8 — The m-8→m-3 digest-flow edge (label independence; unchanged from r0)
The `frozen_core_digest` VALUE flows m-8 → the carriage → my v2 field, copied verbatim (copy-never-compute). **The two exact v2 carrier literals — `m3.app_event.v2` and `m3.e3_observation.v2` — are each INDEPENDENT of every m-8 carrier-version label**; an m-8 "v2" and an `m3.app_event.v2` are unrelated strings. Compatibility is at the digest VALUE and against m-8's settled carrier matrix (the D4 source, `c0b7b488…`), never inferred from spelling.

## §9 — F65 preserved, per-version; the narrowed true claim (unchanged from r0)
v1: closed + byte-frozen; the absorb-refusal intact. v2: closed to its own §2/§3 census — no conductor field exists in either table; unknown fields malformed. Unknown literals: malformed (§5). **The one addition to the r4 evaluator is the §5 dispatch step; the run-constant vector, comparison algorithm, and every other step are byte-unchanged.**

## §10 — Fold log
- r2 (2026-07-22): SCHEMA-R1-F1 folded from `DESIGN-REVIEW-implementer-20260722-171500` — the lone surviving §8 wildcard-form token replaced with the two exact v2 carrier literals (`m3.app_event.v2`, `m3.e3_observation.v2`); §5's four branches and every schema mechanic untouched; only §8 + this header/status/fold metadata moved. The r1 relay's "purged" claim was false at the bytes — this fold's negative is grep-verified (an exact search over the contract returns zero wildcard-form occurrences). r1 SHA-256 `b6fb80ec5f4be99c7eff57291c3ae07377b1690dc7aab064b97d2ba48cbf4d72`.
- r1 (2026-07-22): SCHEMA-R0-F1..F3 folded from `DESIGN-REVIEW-implementer-20260722-160000`. **F1** — the complete v2 closed sets are now PRESENT: the §2 E0 field table (every field, every status) + the §3 six-scope E3 matrix (universal/identity/vector/conditional/forbidden, every cell explicit); v1 remains a hash-bound citation per the review's allowance. **F2** — the two-layer split: byte-optional at layer 1 (absence schema-VALID; no parser infers pipeline position — `phase=failed` exists on both sides of freeze), producer-conformance at layer 2 (present iff FREEZE-REACHED per the settled D4 source, m-8 r5 `c0b7b488…`, enforced by D4/T4 fixtures over the producers, never at acceptance time); predicate 1's `unknown` reachable exactly at byte-valid absence — the reviewer's named branch, taken whole. **F3** — four exact byte-compared literals replace the wildcard (each naming matrix + validators; `m3.unrelated.v1` malformed; no family matching), and the false "history never re-validates" claim replaced with the true one: v1 bytes never migrate/reparse as v2 AND revalidate in full under frozen v1 rules on every read/evaluation. Preserved-decisions list byte-intact. r0 SHA-256 `a09c9931047efe8fa9a52564164fd353e7c12767a09b9763f3cef0c9dc98c534`.
- r0 (2026-07-22): authored per items 1–7; owner decisions: `model_surface_digest` OUT of v2; `logical_surface_digest` IN at schema grain (flagged).
