# m-5↔m-10 Ceiling-Artifact Contract — CANONICAL (Step-3 owner amendment `step3-amend-m5-ceiling`)

CONTRACT_ID: step3-m5-m10-ceiling-artifact-contract
OWNER: m-5 (Workflows & Archetypes) — SOLE policy owner
CONSUMER: m-10 (App Control Plane / Supervisor) — enforcement HOST only
STATUS: PROPOSED — implementer-approved (`step3-amend-m5-ceiling/DESIGN-REVIEW-implementer-20260715-085530`, verdict approve, over rev2 DESIGN `…-084500`); PENDING m-10 hash-confirm + the Master/VP first-stage interface-lock. NOT interface-locked; the locked m-5 §9 enforcement text (`2026-06-30…:158-174`) REMAINS OPERATIVE until the staged fold (VP F20).
BASIS: ratified reframe `master/STEP-3-ARCH-AMENDMENT.md` @ 2d240eb6 (§4 step 2, §8/F12, §8c item 5); GRILL_LOCK step3-amend-m5-ceiling-grill.

This is the SINGLE canonical ceiling-artifact contract (VP F22). m-5 owns it; m-10's design consumes + confirms THIS file's exact SHA-256 — there is no second, drifting copy. The earlier COORDs (`082000`, `085000`) are superseded by this canonical file + its hash.

## 1. Ownership split (the reframe pin)
- **m-5 = SOLE policy owner:** the source policy, the resolution rule, both artifact schemas, the axis lattices, the per-axis fail-closed floors, and the typed reconfiguration/respawn contract.
- **m-10 = enforcement HOST only:** invoke the m-5 resolution at spawn, host the Layer-2 instance in the run manifest, load it at the authority point, and apply the m-5 floors. **m-10 authors no ceiling value, schema, floor, or archetype.**

## 2. Two-layer artifact
**Layer 1 — POLICY artifact (m-5-owned, static config).** The archetype registry / ceiling section of the trusted-config (CQ-4b: per-domain stamped sections under one top-level digest, loaded at genesis). Contents: the archetype→ceiling-vector map (m-5 §6), the axis lattices + per-axis absent-default floors (m-5 §5: `write→read_only, dispatch→none, tool→none`), the vocabulary (m-5 §3). Carries `policy_digest` (the section stamp) and is covered by the trusted-config's monotonic `config_generation`.

**Layer 2 — per-run BOUND artifact (runtime instance).**
```
bound_ceiling := {
  run_id,                    # the app run (m-10 manifest)
  worker_identity,           # the m-9 worker SEAT's conductor-stamped identity (the tool-executing principal)
  seat_archetype,            # the PINNED worker archetype (externally pinned — see spawn_authority_ref)
  spawn_authority_ref,       # the authority that pinned seat_archetype (F1)
  resolved_ceiling_vector,   # { write, dispatch, tool } — VALUES = m-5 policy resolution for seat_archetype
  policy_digest,             # the Layer-1 section stamp resolved from
  config_generation,         # the monotonic never-reused trusted-config epoch at bind time (F2)
  immutable: true            # a change = a NEW binding via §4, never a mutation
}
```
Keyed `(run_id, worker_identity)`; immutable.

## 3. Interface fields
| field | contract |
|---|---|
| **source** | the m-5-owned Layer-1 policy (sole source of every ceiling value + the resolution rule + both schemas + the floors). |
| **writer (Layer 2)** | m-10, at worker spawn, as a **PURE PROJECTION + COPY**: it copies the externally-pinned `seat_archetype` and invokes the m-5 resolution rule for `resolved_ceiling_vector`. **m-10 adds no ceiling value and no archetype choice.** |
| **schema home / instance home** | schema = **m-5** (both layers); instance = the **m-10 run manifest** (app-side), carrying `policy_digest` + `config_generation`. |
| **immutable binding** | `(run_id, worker_identity)`; `worker_identity` = the m-9 worker seat's conductor-stamped identity; immutable. |
| **archetype provenance (F1)** | `spawn_authority_ref` names who pinned `seat_archetype`. **Step-3 MVP = the operator/master-provisioned pinned run manifest for the single MVP lane** (routing DEFERRED, reframe VP F3). **Step-4 = the governed m-4 routing record** (same field). m-10 validation before writing `bound_ceiling`: (1) `spawn_authority_ref` present + recognized; (2) `seat_archetype` ∈ the m-5 §3 vocabulary; (3) resolves to a defined Layer-1 ceiling; (4) m-10 **copies the pinned value verbatim** — never selects/defaults/alters/substitutes. **Fail-closed** on absent/ambiguous/mismatched/**app-local (m-10-internal, no provisioning ref)**/unknown-vocabulary/unresolvable. |
| **read/load path** | at the m-10 authority enforcement point, BEFORE authorizing any m-9 tool dispatch: load `bound_ceiling` for the active `(run_id, worker_identity)`; validate **present · fresh (see freshness) · well-formed · archetype-provenance-valid**; then evaluate the parsed tool call against `resolved_ceiling_vector` (the m-5 §5 lattices — m-5 semantics, m-10 evaluates). Within ceiling → authorize; above → **deny, zero execution** (E2 negative). |
| **freshness (F2)** | fresh **iff** `policy_digest == the current active Layer-1 section stamp` **AND** `config_generation == the current active trusted-config generation` (monotonic, never-reused). Any of: absent · digest mismatch · generation mismatch (higher = superseded; **lower/rollback = regression**) · unknown/reused generation · partial/unresolvable ⇒ **stale ⇒ fail-closed**, **unless** the change arrived via the typed reconfiguration/respawn path (§4), which mints a NEW binding at the new generation. Distinguishes "same bytes still current" (generation matches) from "old bytes reintroduced" (older/reused generation ⇒ stale). |
| **fail-closed** | = the m-5 §5 per-axis absent-default floors (`write→read_only, dispatch→none, tool→none`) ⇒ **deny all tool dispatch = zero execution** (no unbounded execution). Triggered by: absent/stale/malformed artifact, invalid archetype provenance, or generation regression/rollback/reuse. |

## 4. Ceiling-change — the m-5-owned typed reconfiguration/respawn contract
`bound_ceiling` is **immutable at spawn**. A ceiling change **cannot mutate** it; it requires a **new binding via the m-5-owned typed reconfiguration/respawn contract + its gate**:
- a ceiling **raise** (loosen above the prior grant on any axis) is **authority-bearing** → a governed gate (HUMAN_GATE / the sanctioned typed-grant grammar); a **tighten** is monotonic, no raise-gate.
- the runtime/m-10 **retires the old worker and spawns a new one** with the new `bound_ceiling`. No in-place mutation exists.
- the §8b **direct-operator route CANNOT** silently or textually raise the ceiling — any change routes through this typed contract + gate (reframe §8c item 5 / VP F11.3).
m-5 owns the contract's policy shape (record contents, monotonic rules, gate class); the runtime/m-10 hosts the retire+respawn mechanism.

## 5. Cross-domain dependency (consumed, not m-5-authored)
The monotonic, never-reused `config_generation` is a **required property of the CQ-4b trusted-config** (integrity/load owner **m-7**; genesis **m-1**). This contract **consumes** it; it does not author the trusted-config mechanism. If m-7/m-1 realize monotonicity via a top-level-digest-chained counter or equivalent, the freshness rule reads whichever field carries the never-reused epoch. **To confirm with m-7/m-1 at the first-stage reconcile.**

## 6. Consumption model (VP F22)
m-5 owns this single canonical contract. m-10's design **consumes + confirms this file's exact SHA-256**; there is no second drifting copy. Neither side self-declares the join locked — the **Master/VP first-stage reconcile** issues the ONE shared ceiling-interface-lock over both approved artifacts (m-5 + m-10), and only that lock permits stage-2 m-8/m-9.
