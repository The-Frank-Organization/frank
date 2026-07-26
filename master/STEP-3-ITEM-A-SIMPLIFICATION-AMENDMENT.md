# Step-3 Item-A Simplification Amendment (rev7) — the plain interface-lock record replaces §4's soft-edit-stable bundle

**Status:** AUTHORED by master.orchestrator-planner on the **operator's directive** (2026-07-26: "do the dead simple way").
rev7 folds VP review r9 (`step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260727-100000.md`, one mechanical blocker: expand three
surviving path shorthands — the §5.1 close-file reference + two `same file` edge-source selectors — to full literal paths; R8's m-9 D
correction CLOSED). Routed for VP re-review → operator hash-bound ratification (§8b; recorded agent-authored + operator-cited; master does
not self-ratify). The owners' item-A hold stands until ratification, then RELEASED — no owner action.

**Amends** `master/STEP-3-STAGE6-AMENDMENT.md` rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`,
replacing the **interface-bundle mechanism** wherever operative — **§4** (item A), **§6's item-A edge** (`:359`), **§11 steps 4–5**
(`:424–427`), and **§12's bundle-specific VP criterion** (`:432–435`). §7 keeps its bytes; this amendment fixes only the *ordering* by
which its manifest is frozen (§4 below). Every other ratified section stands unchanged.

---

## 1. The change (mechanism)
**OLD (bundle):** a hashable, soft-edit-stable interface **bundle** (`STEP-3-INTERFACE-BUNDLE.json`, `bundle_sha256`, extractor + `--verify`,
dedicated artifacts, `bundle-soft-stability` fixture). **NEW (record lock):** a plain **interface-lock record** —
`master/STEP-3-INTERFACE-LOCK.md` — a **closed manifest** naming every constituent file at its **full SHA-256**, under **"named at this
exact hash; any change to a named byte voids the lock and requires a re-lock."** No extractor, markers, `bundle_sha256`, dedicated
artifacts, or soft-stability fixture. This is the **same byte-bound-hash lock the team has used for every approval in this project.**

## 2. Why (rationale)
The settled bases are already frozen + pair-approved + byte-bound; frozen design-of-record is not cosmetically edited, so
soft-edit-stability (F101) solves a problem that does not arise. The plain fingerprint-record lock is proven, and aligns with the
operator's **confusion-firewall + MVP-minimality** philosophy — cut ceremony that does not earn its cost.

---

## 3. Supersession surface (exact — r4-F1 / r5-F2 / r6 all confirmed complete)
On ratification, the bundle mechanism is superseded wherever rev12 makes it operative; no other rev12 byte moves.

| rev12 locus | replaced by |
|---|---|
| **§4** (`:82–108`) | the **§5 record-lock contract** of this amendment |
| **§6 item-A edge** (`:359`) | "A: the interface-lock **record** authored last, byte-binding settled B–E + the joins + the governing amendments." |
| **§11 step 4** (`:424–425`) | "Author `master/STEP-3-INTERFACE-LOCK.md` (§5 record); VP + F73-review it — **that is item A.** No fixture freeze in item A." |
| **§11 step 5** (`:426–427`) | "**Lane 4, in this exact order:** (i) author + content-address the immutable §7 fixture-input + baseline artifacts; (ii) freeze `STEP-3-EXIT-FIXTURES.json` with final non-placeholder digests; (iii) Master+VP re-lock over **both** the interface-lock record (named by its **externally-recorded** full SHA-256, §5.4) **and** that frozen manifest. Then lane 5/T4." |
| **§12 VP criterion** (`:432–435`, bundle clause) | "does `master/STEP-3-INTERFACE-LOCK.md` **close** the manifest (§5 — every constituent file at exact path + full SHA-256 under the repeatable `{role, path, clause}` row model), carry the **exhaustive typed precedence edges** of §5.3, and rely on **external** hash-binding (no self-hash)." All other §12 criteria stand unchanged. |

**Post-ratification source-fold manifest** (owed master work AFTER ratification, NOT part of the lock; historical relays/ledger stay
append-only): `ROADMAP.md`; `master/README.md`; `master/ARCHITECTURE.md` (with the owed D7/`relay.submit` consolidation);
`master/domains/m-3-observation-evidence/README.md`; `master/CYCLE-PLAYBOOK.md:408`; `master/domains/m-1-trust-identity/README.md:111`;
`master/domains/m-2-forms-determinism/README.md:59`; **withdrawn:** `master/STEP-3-ITEM-A-RECIPE.md` r3 `06e6956e…`.

## 4. Fixture-digest ordering (r4-F2 / r5-F2 — un-fusing MOVES the freeze to lane 4, it does not dissolve the circularity)
§7 (`:377–387`) requires concrete `input_artifact_sha256` + baseline digests in the frozen manifest, so lane 4 runs one order (identical
to §3's §11-step-5 row): **(i) author + content-address the immutable fixture-input + baseline artifacts → (ii) freeze the manifest with
final digests → (iii) re-lock over the externally-named record + the frozen manifest → (iv) T4 builds executable fixtures, filling no
hash-bound slot.** If an input genuinely cannot exist before T4, that is a §7/§11 ordering amendment through this same gate — not a mutable slot.

## 5. The §5 record-lock contract (r4-F3 / r5-F1,F3 / r6-F1 — a closed, literal, externally-bound manifest)
`master/STEP-3-INTERFACE-LOCK.md` is a **closed manifest**: the settled interface is exactly reproducible from the named files alone.
No free-form design prose — only rows, the precedence edges of §5.3, and the invalidation rule.

**Amendment-vs-record split.** This amendment fixes the **complete row set** — every row's `{role, path, clause}` and the exhaustive
precedence edges (§5.2/§5.3) — with **exactly one** future slot: this amendment's own operator-ratification relay, which does not exist
until ratification. Every other path is literal here. The item-A record is the literal instance that fills each row's verified full
`sha256` (recomputed at authoring) and is bound externally (§5.4); the VP reviews the filled record at item A before lane 4. **No path,
role, clause, or precedence edge is deferred to authoring.**

### 5.1 Row model (r6-F1 — repeatable semantic binding; r7-F1 — every row carries a clause)
Each row's identity is **`{role, path, clause}`**; `sha256` is that file's on-disk hash. A file **MAY appear in multiple rows** under
different `{role, clause}` (e.g., the lane-2 close is both the item-E binding and the carried-source of record). The **set of distinct
`path`s is the byte-bound file set.** `role ∈ { owner_base, frozen_final, governing_amendment, join_or_settlement, carried_source }`.
**Clause rule (r7-F1):** every row's `clause` is **`whole_file`** — the lock binds the entire file — **except** the rows on the close file
`master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md`, whose distinct clauses are stated inline in §5.2. (The precedence edges of §5.3 name specific *sections* of an owner base as
the superseded locus; that is the edge's source selector, not the owner-base row's clause, which is `whole_file`.)

### 5.2 The closed row set (literal paths + clause; record fills verified full sha256)
**`owner_base` (8) — clause `whole_file` each:**
- `master/domains/m-1-trust-identity/design/2026-07-22-stage6-lane2-env-redaction.md` (`d34a7c47…`)
- `master/domains/m-2-forms-determinism/design/2026-07-22-stage6-e-logical-component.md` (`c3a8cd61…`)
- `master/domains/m-2-forms-determinism/design/2026-07-23-stage6-c-relay-submit-resource.md` (`5ec7a3d2…`)
- `master/domains/m-3-observation-evidence/design/2026-07-22-stage6-lane2-e0-e3-delta.md` (`651c9aec…`)
- `master/domains/m-8-provider-adapters/design/2026-07-22-stage6-BE-digests-addendum.md` (`734e44b7…`, incorporating integrated r5 `c0b7b488…`)
- `master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md` (`01b885fe…`)
- `master/domains/m-10-app-control-plane/design/2026-07-22-stage6-lane2-producer-delta.md` (`3e3c5192…`)
- `master/domains/m-10-app-control-plane/design/2026-07-23-lane2-be-carriage-row.md` (`cd17db32…`)

**`frozen_final` (8) — clause `whole_file` each:**
- `master/domains/m-1-trust-identity/design/2026-07-16-step3-mvp-secret-boundary-seat-identity.md` (`7c8b09a6…`)
- `master/domains/m-2-forms-determinism/design/2026-07-16-step3-mvp-form-schema-mapping.md` (`83d8e63e…`)
- `master/domains/m-3-observation-evidence/design/2026-07-16-step3-mvp-egress-e0-e3.md` (`009df607…`)
- `master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md` (`4b670a79…`)
- `master/domains/m-9-model-runtime/design/2026-07-19-mvp-full-worker.md` (`cb7ff970…`)
- `master/domains/m-9-model-runtime/design/2026-07-17-mvp-lifecycle-half.md` (`4d3bd14e…`)
- `master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md` (`d2ce9831…`)
- `master/domains/m-10-app-control-plane/design/2026-07-19-mvp-control-plane.md` (`6fd1d655…`)

**`governing_amendment` (each amendment + its operative ratification relay — clause `whole_file` each; all literal except the one future slot):**
- `master/STEP-3-STAGE6-AMENDMENT.md` (rev12 `1125b0a0…`) — ratified by `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-165500.md` (`7c367c7f…`)
- `master/STEP-3-STAGE6-M3-SCHEMA-AMENDMENT.md` (`9e874df8…`) + bound contract `master/domains/m-3-observation-evidence/design/2026-07-22-e0-e3-schema-version-contract.md` (`6e2abe40…`) — ratified by `master/relays/step3-stage6-m3-schema-amend/RECONCILE-orchestrator-planner-20260722-234500.md` (`49c811fd…`)
- `master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md` (`1fa71cb8…`) — operative by `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260725-170000.md` (`984071fb…`)
- `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` (this amendment, at its ratified hash) — **the single future slot:** its operator-ratification relay, resolved to one literal path + full SHA-256 by the item-A record post-ratification

**`join_or_settlement` (clause `whole_file` each, except the two close-file rows noted):**
- §D two-sided resume seam — co-sign `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-123000.md`; legs `master/relays/step3-relock-settlement-amend/DESIGN-planner-m9-20260726-024000.md`, `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260726-020000.md`, `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260723-023020.md`
- §B two-sided sink — `master/relays/step3-relock-settlement-amend/DESIGN-planner-m9-20260726-131500.md`; `master/relays/step3-relock-settlement-amend/DESIGN-planner-m3-20260726-133000.md`
- B-carriage (m-8→m-10→m-3) — `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260725-150000.md`
- item-C — `master/relays/step3-relock-settlement-amend/SITREP-planner-20260723-041943.md`; `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260726-140000.md`; `master/relays/step3-relock-settlement-amend/DESIGN-planner-m9-20260726-144500.md`
- **m-10-C confirmation** (discharges the m-1 §4 m-10-C half; r7-F1) — `master/relays/step3-relock-dag-m10/SITREP-planner-20260722-015123.md` (`774cd380…`)
- item-E (R1@r24) — `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md` (**clause "item-E R1@r24 settled inventory"**)
- lane-2 close of record — `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md` (**clause "lane-2 DAG close of record"**); env-digest-locus correction (clause `whole_file`) — `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-163000.md`

**`carried_source` (disposition lineage only, §6 — all three on the close file under distinct clauses):**
- N910 disposition — `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md` (**clause "N910 documented-MVP-limit disposition"**)
- r7-mirror disposition — `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md` (**clause "r7-mirror v3-deferred + re-open-predicate disposition"**)
- env_digest-parity disposition — `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md` (**clause "env_digest-parity accepted disposition"**; locus in the m-1 `owner_base`; realized by m-9 §7 + m-3's E3 observer)

### 5.3 Precedence edges (r5-F3c / r6-F1 / r7-F1 — EXHAUSTIVE; owner-base conflict census over ALL operative sections, full literal paths)
Census run over all 8 `owner_base` files at their **operative sections** (not just status-ledger vocabulary), separating live prose from
revision history. Complete set below: m-1 (§4) + m-9 (C, B, receipts) + m-10 (§10/S/`assign`). **m-2/m-3/m-8 carry no operative cross-seam
status; the m-10 B/E carriage row's m-9-r12 reference is a resolved forward reference, not pending; revision-history statements stay
history.** No general "later governs" rule; manifest order and filename time carry no authority — only these listed edges govern. **Owner-base
bytes are UNCHANGED in every edge** (the edge governs the conflicting status; it moves no byte). All paths literal.

1. **m-1 §4 "The PARKED producer-attaching halves"** — ( `master/domains/m-1-trust-identity/design/2026-07-22-stage6-lane2-env-redaction.md`, §4 `:57–60`: **all four** halves — m-9-C (`:58`), **m-9-D** (`:58`, the §2.2 writer gate / §2.3 create-open-verify + RED battery / §2.4 route-labeled sentinel legs / K6 exclusion), m-10-C (`:59`), and §D-redaction-co-sign (`:60`) ) **→ superseded by, per half (r8-F1):**
   - m-9-C half → the **item-C** rows `master/relays/step3-relock-settlement-amend/SITREP-planner-20260723-041943.md` + `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260726-140000.md` + `master/relays/step3-relock-settlement-amend/DESIGN-planner-m9-20260726-144500.md`;
   - **m-9-D half** (parked half #2, m-9 D at-rest/redaction — distinct from the co-sign half) → `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260723-023020.md` (m-1 §D leg, `:42–45` "discharges parked half #2, m-9 D at-rest/redaction") + `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-123000.md` (§D co-sign);
   - m-10-C half → `master/relays/step3-relock-dag-m10/SITREP-planner-20260722-015123.md`;
   - §D-redaction-co-sign half (parked half #4 — distinct source from m-9-D, shared targets) → `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260723-023020.md` (m-1 §D leg) + `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-123000.md` (§D co-sign);
   - all ratified in the close `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md`.
2. **m-9 C-consumption park** — ( `master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md`, §7 `:476` "m-10 owns the ticket schema + gate (consumption PARKED)", §9 `:499`/`:509` "m-10's C-ticket schema: PARKED", §11 `:559` "still-parked producer input" ) **→ superseded by** the **item-C** rows `master/relays/step3-relock-settlement-amend/SITREP-planner-20260723-041943.md` + `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260726-140000.md` + `master/relays/step3-relock-settlement-amend/DESIGN-planner-m9-20260726-144500.md`, ratified in the close `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md`.
3. **m-9 B-consumability obligation** — ( `master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md`, §8 `:494` "I carry the m-8-computed `frozen_core_digest` …", §11 `:559` owed-list ) **→ superseded by** the **§B sink** `master/relays/step3-relock-settlement-amend/DESIGN-planner-m9-20260726-131500.md` + `master/relays/step3-relock-settlement-amend/DESIGN-planner-m3-20260726-133000.md`, ratified in the close `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md`.
4. **m-9 §9 items 4/5 "EXACT-FOLDED, JOINT-PENDING"** (disposition receipt + content-ready receipt frame) — ( `master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md`, §9 items 4/5 ) **→ superseded by** the **§D co-sign** `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-123000.md`, ratified in the close `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md`.
5. **m-10 producer pending carriage** — ( `master/domains/m-10-app-control-plane/design/2026-07-22-stage6-lane2-producer-delta.md`, §10 "consumer carriage PARKED" + the `assign` writer-to-reader carrier "JOINT-PENDING" + S-1..S-5 "JOINT-PENDING until the §D join co-sign" ) **→ superseded by** the **§D co-sign** `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-123000.md`, the **§B sink** `master/relays/step3-relock-settlement-amend/DESIGN-planner-m9-20260726-131500.md` + `master/relays/step3-relock-settlement-amend/DESIGN-planner-m3-20260726-133000.md`, and **B-carriage** `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260725-150000.md`, ratified in the close `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md`.

### 5.4 External binding (r5-F1 / r6-F2 — no self-hash, no future-hash citation)
The record contains **no** field naming its own SHA-256; it states only the invalidation rule: *"any change to this record or to any named
byte voids the lock and requires a re-lock."* The record's full SHA-256 is named **externally** by the VP item-A review relay and by the
lane-4 Master+VP interface-lock relay. **This amendment's own operator-ratification record binds only this amendment's exact hash** — it
does not and cannot cite the later interface-lock record hash (r6-F2). No distinct post-item-A operator gate is added; the lane-4 Master+VP
lock binds the record (as the earlier joint interface-lock `b7e1f0ef` was Master+VP, not operator-ratified).

## 6. Carried-obligation boundary (r4-F4 — one boundary)
The interface lock records only the **`carried_source` disposition relays** as governing lineage; it does not restate obligation text.
**Lane 4 alone** owns their executable fixture records + expected canonical rows (in `STEP-3-EXIT-FIXTURES.json`, per §7 + §4). The
env_digest-parity **locus** is captured by the m-1 `owner_base` hash itself — no free-form locus entry. No obligation text lives in both.

## 7. Sequencing
Master authors `master/STEP-3-INTERFACE-LOCK.md` (the §5 closed manifest) → VP + F73 review naming its external SHA → **that is item A.**
Then **lane 4** runs the §4 order (fixture inputs/baselines → freeze `STEP-3-EXIT-FIXTURES.json` → Master+VP interface-lock over the
externally-named record + the frozen manifest). Then **lane 5** (T4). **H-12 hard-blocks external use throughout.**

## 8. Ratification + withdrawal
VP re-review → **operator hash-bound ratification of THIS amendment's exact hash** (§8b; recorded agent-authored + operator-cited). On
ratification: the §4/§6/§11/§12 bundle mechanism is superseded per §3; `master/STEP-3-ITEM-A-RECIPE.md` r3 `06e6956e…` + the bundle
apparatus are **WITHDRAWN** and not built. Ratified rev12 `1125b0a0…` keeps every non-superseded byte; the eight settled bases + frozen
finals + governing amendments are UNMOVED. No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3,
merge, deploy, or `frank/` action is authorized here.
