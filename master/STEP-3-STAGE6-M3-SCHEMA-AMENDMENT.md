# Step-3 Stage-6 m-3 Schema-Version Amendment (rev3) — the two m-3 evidence carriers go v1→v2 (version-dispatched), binding the pair-approved m-3 closed-schema contract by hash

**Status: PROPOSED rev3 — pending VP exact-byte review → operator hash-bound ratification.** rev3 is an
**amendment-wrapper-only** revision over rev2 (`5c3f604efdcb9dc920b89033f7d8ba7f76e6ec9d99a17f014db6a79742ba6809`,
VP-REVISE-NARROW `…-203000`) — **the bound contract `6e2abe40…` is byte-IDENTICAL and needs no fresh m-3 pair review.**
rev3 closes the two wrapper defects: **R3-F1** — remove the wildcard-form and brace-family schema shorthands that
contradicted the bound contract's four-exact-literal, no-family-matching discipline; the four `schema` literals are now
named explicitly wherever dispatch is summarized, and no wildcard-form or brace-family token remains in these bytes.
**R3-F2** — surface BOTH v2 census decisions the operator ratifies, not one (`logical_surface_digest` IN E0 v2;
`model_surface_digest` OUT of v2, deferred to a later v3 E-join). rev2 already
closed the three R2 blockers: **F1** the exact closed-schema contract is **bound by hash**; **F2** strict non-gating is
the **sole** D2 branch; **F3** D4 depends on the **final pair-approved m-8 producer revision** (converged at r5).
Rev1 was `edbbfb7c…`. A narrow, additive
amendment against operator-ratified **Stage-6 amendment rev12** (`master/STEP-3-STAGE6-AMENDMENT.md` @
`1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`). Its honest scope is **two m-3 schema identities
+ their parser/version-dispatch behavior + the D2 gating clarification** (real contract changes; no rev12 base byte is
edited in place). **Master authors + routes; master does NOT self-ratify** — the operator ratifies the exact reviewed
hashes (§8b: agent-authored + operator-cited). Until ratified: proposed input only — no DESIGN-lock, PLAN, T4 token,
credential, provider call, release binding, live E3, merge, deploy, or lane-2 r1 fold.

**Basis:** the m-3 lane-2 closed-parser collision (rev1 §0); VP amendment-reviews r1 (`…-100000`, reclassifying D1 as
a ratified-interface change) + r2 (`…-121500`, F1/F2/F3). The route-(b) deliverable — m-3's pair-approved
schema-version contract — is filed and bound below.

---

## §0 — Why an amendment (the closed-parser collision, unchanged)
Frozen m-3 r4 (`009df607…`) defines both evidence carriers as **closed schemas** (E3 §3, E0 §2.2) whose unknown-field
rejection structurally enforces F65 absorb-refusal. Rev12 §5-B/§5-E require adding `frozen_core_digest` (E0 + E3) and
`logical_surface_digest` (E0) to those exact v1 records. Adding fields to v1 would relax the closed-schema rule and
weaken F65; the coherent fix is a **new version with per-version closed matrices + explicit version dispatch** (v1 stays
byte-frozen + closed; v2 carries the ratified fields). Because rev12 fixes the v1 literals as Tier-HARD, this is a
ratified-interface amendment, not a master erratum.

## §1 — The two-schema v2 closed-schema contract: BOUND BY HASH (F1 closed)
The complete, executable closed-schema contract is the **pair-approved m-3 artifact**
**`master/domains/m-3-observation-evidence/design/2026-07-22-e0-e3-schema-version-contract.md` @ SHA-256
`6e2abe40be7a6031163348d47e6b9c13990c5599eb877d30ef196b59efbf3e1f`** (m-3 pair approve
`step3-stage6-m3-schema-amend/DESIGN-REVIEW-implementer-20260722-190000`, four rounds, zero surviving findings),
**bound here by hash.** Ratifying this amendment ratifies that exact contract hash; a byte change to the contract voids
the binding and requires fresh m-3 pair review + re-binding. Its closed set is IN those bytes — a binder never
interprets "same as another document except." What the bound contract establishes (summary, not a re-render):
- **The complete `m3.app_event.v2` field/status table** and the **complete six-scope `m3.e3_observation.v2` matrix** —
  v1 field reference + the full v2 required/forbidden matrices, present in the bytes.
- **Four-literal exact version dispatch** — the four byte-exact `schema` literals `m3.app_event.v1`,
  `m3.app_event.v2`, `m3.e3_observation.v1`, `m3.e3_observation.v2`, each selected by the record's own `schema`
  literal (**no wildcard, no family/brace matching** — the contract's §5 rejects family matching); an unknown version
  ⇒ `malformed`; a v1-literal record carrying any
  v2-only field, or a v2-literal record missing a v2-required field / carrying any unknown field ⇒ `malformed`
  (mechanically-decidable per-version checks, replacing rev1's fuzzy "cross-version mixture").
- **v1 byte-frozen + still fully revalidated** — v1 never migrates; every v1 read revalidates in full under the frozen
  v1 closed rules. F65 absorb-refusal preserved per-version (v1 unchanged; v2 closed to its own matrix).
- **Byte-only presence discipline (R0-F2 re-cut):** additive-field absence is schema-VALID at parse (no parser infers a
  pipeline position); the per-cut *producer* requiredness (`frozen_core_digest` present iff freeze-reached) is enforced
  by the D4/T4 fixtures, and predicate 1's `unknown` branch is reachable exactly at byte-valid absence. The run-constant
  acquisition/comparison vector + algorithm are UNCHANGED — only well-formedness/version-dispatch is added.

**The TWO census decisions in the bound hash, consciously surfaced (what the operator ratifies):** the bound contract
makes two explicit schema-sequencing choices, and ratifying both hashes ratifies both:
1. **`logical_surface_digest` is IN the v2 E0 census** at schema grain (presence/encoding/producer-identity = m-9), with
   the recipe-binding confirmation parked (D3). Master AFFIRMS keep: rev12 §5-E requires `logical_surface_digest` to
   ride E0, so a complete v2 E0 census **must** carry the field at schema grain; parking only the recipe-binding is
   exactly the D3 schema-now/binding-parked staging.
2. **`model_surface_digest` is OUT of both v2 carriers** — it lands only with the parked E-join, under a later governed
   **v3** bump (bound contract §§62-63/82). Master AFFIRMS: `model_surface_digest` is the m-3 E3 join of the two
   component digests, which is parked (D3) until the producer join settles, so it correctly does not enter v2 now.
A different choice for either is not an alternate inside this hash — it requires changed contract bytes, fresh m-3 pair
review, re-binding, VP review, and operator ratification. (Removing `logical_surface_digest` from v2 would cost exactly
one row in the contract's §2.1; master recommends keeping it, per §5-E.)

## §2 — D2: typed predicates 2 and 5 are STRICT NON-GATING — the SOLE normative branch (F2 closed)
Rev12 §7 gates the Governance-binding leg (`xit-gov-1`) on **predicates 1 ∧ 3 ∧ 4 ONLY** (`:370`). Predicates **2
(`provider_deny_caused_zero_transport`)** and **5 (`no_alternate_credentialed_provider_route_observed`)** are in the
§5-E typed-predicate SET but named in **no leg fixture**; rev2's "these feed the §7 exit legs" was a master over-claim.

**Normative disposition (sole, non-optional): STRICT NON-GATING.** Both predicate contracts exist (§5-E fixes the
five-id set), but their verdicts are **recorded/reported only** — a `fail`/`unknown`/missing predicate-2 or -5 record
does **NOT** fail or hold any §7 leg or the Step-3 exit. The **§7 six-leg gate is UNCHANGED; there is no hidden seventh
condition.** This is the operative contract of these bytes; there is no alternate branch inside this hash.
- **Distinct from the build proof (mechanical "no seventh condition"):** the independently-required
  **deny→zero-provider-transport BUILD proof** lives in `STEP-3-MVP-AMENDMENT.md` §10 (an instrumented-negative that
  the MVP must pass) — that is a **separate mechanism** from the typed E3 predicate-2 *record*. This amendment does not
  merge them: §10's build proof stands as-is; the typed predicate-2/5 records are non-gating diagnostics.
- **Honest consequence:** predicate 5's "no alternate credentialed provider route" property is therefore not
  §7-gated in the MVP — consistent with the ratified six-leg gate as written.
- **A future required-proof choice is a SEPARATE new amendment** — it would need each predicate's exact named consumer +
  leg/composite/fixture-manifest consequence, fresh VP review, and operator ratification. It is not latent in these
  bytes.

## §3 — D4: consume the FINAL pair-approved m-8 producer revision (F3 closed)
The m-8 producer lane has **converged**: the B/E producer delta is **pair-approved at r5 `c0b7b488…`** (m-8 review
`step3-relock-dag-m8/DESIGN-REVIEW-implementer-20260722-143000`, `approve`, after four must-revise rounds
r2→r3→r4→r5). D4 is a **dependency, not a reopened choice** — it preserves the accepted deny/post-freeze-reject carrier
matrix. **Disposition:** the settled m-8 producer basis is r5 `c0b7b488…`; **after this amendment ratifies**, master
routes that settled carrier matrix to m-3, and m-3 authors its parked exhaustive **cut-matrix** against those settled
bytes (never assumed carriers). The m-3 v2 carriage must be **version-compatible** with m-8's settled versioned
carriers — checked at the digest **value**, never assumed from a shared "v2" spelling (m-8's carrier literals and the
m-3 literals `m3.app_event.v2` / `m3.e3_observation.v2` are independent, per §1's bound contract). *(Rev1's stale "r3 pair approval" pin is void; the dependency is
version-agnostic and now satisfied at r5.)*

## §4 — D3: APPROVED, carried unchanged
m-3 authors the E0 schema grain (in the bound §1 contract); the recipe/binding confirmation (the m-3 B-sink record +
the E3 two-digest join) stays **parked** until exact pair-approved m-9/m-8 producer bytes exist — which now do, so the
parked lane-2 r1 becomes routable **after ratification** (§6). This amendment does not itself unpark r1.

## §5 — What this amendment does NOT change
No ratified DECISION moves: the digest still rides the m-3 E0 + E3 carriers + the composite exit proof; the observer
still derives independently; the **§7 six-leg gate is intact** (strict-non-gating D2); the run-constant
acquisition/comparison vector + algorithm are unchanged; **F65 absorb-refusal is preserved and per-version
strengthened.** Only the two m-3 schema-version identities + the evaluator/producer version-dispatch are amended (via
the bound contract). No m-8/m-10 record-version literal is touched (rev12 names those carriers generically; they remain
each pair's own frozen-final delta — m-8's is r5 `c0b7b488…`). Lane 1, broker rev8, NO-H-24, and the four-item
affected-final ledger are not reopened.

## §6 — Ratification gate
1. Master routes THIS amendment's exact bytes **+ the bound contract hash `6e2abe40…`** for **VP exact-byte review**.
2. On VP approve, master routes both hashes to the **operator for hash-bound ratification** (§8b — agent-authored +
   operator-cited; master does not self-ratify). Ratifying this amendment ratifies its bytes AND the bound m-3 contract
   `6e2abe40…` together.
3. **Only after ratification** may master route the settled m-8 r5 cut basis to m-3, and m-3 fold the parked lane-2 r1
   (cut-matrix + verdict machines + sink + E-join) against settled bytes.
Until then: m-3 r0 stays must-revised, r1 held, every downstream gate held, H-12 external-use block stands.
