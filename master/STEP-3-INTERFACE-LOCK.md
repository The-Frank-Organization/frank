# STEP-3 INTERFACE LOCK — the byte-bound record of the settled Step-3 MVP interface (item A)

**This is item A** of the §11 re-lock sequence, authored to the ratified contract `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md`
rev7 SHA-256 `3443f73d30240c90aa6b1b46cbcd3763444592c422ff8ab655307bc1ebbd7373` (operator-ratified 2026-07-27, ratification relay
`master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260727-130000.md`), which superseded the §4 interface-bundle mechanism.

**Invalidation rule (the lock).** This record names each constituent file at its exact SHA-256. **Any change to any named byte — or to
this record — voids the lock and requires a re-lock.** The recorded hash set *is* the lock; there is no separate `bundle_sha256` and no
extractor.

**External binding (no self-hash).** This record contains no field naming its own SHA-256. Its full hash is named **externally** by the VP
item-A review relay and by the lane-4 Master+VP interface-lock relay; any later record cites that same external hash. (A file cannot
contain its own hash.)

**Status:** AUTHORED for VP + F73 review — that review *completes item A*. This record is not itself an approval or a re-lock; lane 4 is
the Master+VP re-lock over this record (named by its external SHA) plus the frozen exit-fixtures manifest. **H-12 hard-blocks external use.**

**Row model.** Each row is identified by `{role, path, clause}`; `sha256` is the file's on-disk hash. A file MAY appear in multiple rows
under different `{role, clause}` (the lane-2 close appears in five). The **set of distinct `path`s is the byte-bound file set** — **38
distinct files across 42 semantic `{role, path, clause}` rows** (37 files each in a single row + the lane-2 close contributing five rows).

---

## 1. `owner_base` (8) — the settled lane-2 owner bases · clause `whole_file`

| path | sha256 |
|---|---|
| `master/domains/m-1-trust-identity/design/2026-07-22-stage6-lane2-env-redaction.md` | `d34a7c475c5988f12872a5aaf599b201fe56970b766b935bdb42b1d2851ab1ef` |
| `master/domains/m-2-forms-determinism/design/2026-07-22-stage6-e-logical-component.md` | `c3a8cd61bcbe33a3ede847b29d7416563000c66f949fa2950ea0af305edbbd2c` |
| `master/domains/m-2-forms-determinism/design/2026-07-23-stage6-c-relay-submit-resource.md` | `5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f` |
| `master/domains/m-3-observation-evidence/design/2026-07-22-stage6-lane2-e0-e3-delta.md` | `651c9aecd1dd27b61d1a5fd3e802b05ee350e1acb848b827ee57b27751eccb97` |
| `master/domains/m-8-provider-adapters/design/2026-07-22-stage6-BE-digests-addendum.md` | `734e44b7841754dfca56f3a9195695bed1d2f07b42d3acae92ef6b188b47fb53` |
| `master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md` | `01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b` |
| `master/domains/m-10-app-control-plane/design/2026-07-22-stage6-lane2-producer-delta.md` | `3e3c51926332f3068aad8c995d898a11185b0d35608213a0d97903b6cf8d53f6` |
| `master/domains/m-10-app-control-plane/design/2026-07-23-lane2-be-carriage-row.md` | `cd17db3204286f4a0b165b1e349d75c5cd9320dc793b021e942275427115e35f` |

## 2. `frozen_final` (8) — the frozen owner finals the bases rest on · clause `whole_file`

| path | sha256 |
|---|---|
| `master/domains/m-1-trust-identity/design/2026-07-16-step3-mvp-secret-boundary-seat-identity.md` | `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c` |
| `master/domains/m-2-forms-determinism/design/2026-07-16-step3-mvp-form-schema-mapping.md` | `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d` |
| `master/domains/m-3-observation-evidence/design/2026-07-16-step3-mvp-egress-e0-e3.md` | `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad` |
| `master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md` | `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51` |
| `master/domains/m-9-model-runtime/design/2026-07-19-mvp-full-worker.md` | `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45` |
| `master/domains/m-9-model-runtime/design/2026-07-17-mvp-lifecycle-half.md` | `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd` |
| `master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md` | `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146` |
| `master/domains/m-10-app-control-plane/design/2026-07-19-mvp-control-plane.md` | `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf` |

## 3. `governing_amendment` — each governing amendment + its operative/ratification relay · clause `whole_file`

| what | path | sha256 |
|---|---|---|
| stage-6 rev12 (the controlling contract) | `master/STEP-3-STAGE6-AMENDMENT.md` | `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183` |
| ↳ its ratification relay | `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-165500.md` | `7c367c7f27b41f162b5a433e934a5debf0cba242d6d68e7be60d8141bad175dd` |
| m-3 schema amendment | `master/STEP-3-STAGE6-M3-SCHEMA-AMENDMENT.md` | `9e874df84015261d77e9c353528e821fd8491489388c010fa621fe630432b351` |
| ↳ its bound contract | `master/domains/m-3-observation-evidence/design/2026-07-22-e0-e3-schema-version-contract.md` | `6e2abe40be7a6031163348d47e6b9c13990c5599eb877d30ef196b59efbf3e1f` |
| ↳ its ratification relay | `master/relays/step3-stage6-m3-schema-amend/RECONCILE-orchestrator-planner-20260722-234500.md` | `49c811fd0e04577a783fd3a3945101a69dd8f89d0a6e4cea16ad66288ff7d82e` |
| §D-settlement amendment | `master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md` | `1fa71cb8def2d52690cc513faa7860e0d8c2c2224b33b38ff97808990adc879b` |
| ↳ its operative relay | `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260725-170000.md` | `984071fb0b6b093767caeaf2a84321d3eb91171b2e359b7545cad1be20d9112a` |
| item-A simplification amendment (this lock's governing contract) | `master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` | `3443f73d30240c90aa6b1b46cbcd3763444592c422ff8ab655307bc1ebbd7373` |
| ↳ its operator-ratification relay (the resolved single future slot) | `master/relays/step3-relock-item-a/RECONCILE-orchestrator-planner-20260727-130000.md` | `cabae8bd16ed179bc1df8e261c10ecba8472f230e9afd1961e846ea5058b6f8c` |

## 4. `join_or_settlement` — the cross-owner seam bindings · clause `whole_file` except the close-file rows in §5

| seam / role | path | sha256 |
|---|---|---|
| §D two-sided resume seam — co-sign | `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-123000.md` | `2f3fb651d833f4c804af8a2a8e628da12affa69ab4f4d0cc042ac28674eb3e13` |
| §D leg — m-9 | `master/relays/step3-relock-settlement-amend/DESIGN-planner-m9-20260726-024000.md` | `c9f32c632602d27dba9886e2e2d06f753e50f468ea759e18f25aa3abd189332c` |
| §D leg — m-10 | `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260726-020000.md` | `2a8774f83f6dc4fba8794f7d1266de4c719070a5fe7f9eb9a5d7b161048dd6b7` |
| §D leg — m-1 (also discharges m-1 §4 half #2 m-9 D + half #4 co-sign) | `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260723-023020.md` | `d096a4b357742f4ef6005207891e094ec52f35486585cf7bde37e127081dae3d` |
| §B two-sided sink — m-9 | `master/relays/step3-relock-settlement-amend/DESIGN-planner-m9-20260726-131500.md` | `95e8c6aaf085685122400d1c621fdea092d88e9fd9f720cff48f380eaaa35e90` |
| §B two-sided sink — m-3 | `master/relays/step3-relock-settlement-amend/DESIGN-planner-m3-20260726-133000.md` | `185acf9e0a762d0743affc0926ab50f7b11f81d484f685c5be084eb228e20fc3` |
| B-carriage (m-8→m-10→m-3) | `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260725-150000.md` | `f4cd3fab0d959682e77a79a94be9cc3f210a9156adadd96dc0ed6223f2ee228c` |
| item-C — m-1 C-confirm | `master/relays/step3-relock-settlement-amend/SITREP-planner-20260723-041943.md` | `3070f34cf2dc74fe254e66f8ac8beb1701ec6a48d2c57eda1723cf6a4dcd7f6d` |
| item-C — m-10 | `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260726-140000.md` | `8efa3b020629f9c98dffb155f429b549e84e0f7f797f8534ae435f0cddb0a3a4` |
| item-C — m-9 | `master/relays/step3-relock-settlement-amend/DESIGN-planner-m9-20260726-144500.md` | `fba9f6f54af6eb11918ed7c8cb8f3870d8469f125d05ccc98d9aaad10db6ace7` |
| m-10-C confirmation (discharges m-1 §4 half #3 m-10 C) | `master/relays/step3-relock-dag-m10/SITREP-planner-20260722-015123.md` | `774cd3809b222ca89c18ffc120b6d7d465d5574e2b67add5ea08bcbc05c53f5b` |
| env-digest-locus correction | `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-163000.md` | `84c6d6ab3816c7976e44b54396915b2a123e3e104480cad51552631466f65384` |

## 5. Multi-role rows on the lane-2 close (one file, five `{role, clause}` rows) · sha256 `fa2a634f396e71dd3ce5de3f4dbf2e1ac3651fc156b8dde0edada90df8df3c6f`
`path` = `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md`

| role | clause |
|---|---|
| `join_or_settlement` | item-E R1@r24 settled inventory |
| `join_or_settlement` | lane-2 DAG close of record |
| `carried_source` | N910 documented-MVP-limit disposition |
| `carried_source` | r7-mirror v3-deferred + re-open-predicate disposition |
| `carried_source` | env_digest-parity accepted disposition |

## 6. Precedence edges (typed, exhaustive — verbatim from ratified amendment §5.3, full literal paths) — later governing/join records supersede named owner-base status; **owner bytes UNCHANGED**
No general "later governs" rule; manifest order and filename time carry no authority — only these listed edges govern.

1. **m-1 §4 "The PARKED producer-attaching halves"** — ( `master/domains/m-1-trust-identity/design/2026-07-22-stage6-lane2-env-redaction.md`, §4 `:57–60`: **all four** halves — m-9-C (`:58`), **m-9-D** (`:58`, the §2.2 writer gate / §2.3 create-open-verify + RED battery / §2.4 route-labeled sentinel legs / K6 exclusion), m-10-C (`:59`), and §D-redaction-co-sign (`:60`) ) **→ superseded by, per half:**
   - m-9-C half → the **item-C** rows `master/relays/step3-relock-settlement-amend/SITREP-planner-20260723-041943.md` + `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260726-140000.md` + `master/relays/step3-relock-settlement-amend/DESIGN-planner-m9-20260726-144500.md`;
   - **m-9-D half** (parked half #2, m-9 D at-rest/redaction — distinct from the co-sign half) → `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260723-023020.md` (m-1 §D leg, `:42–45` "discharges parked half #2, m-9 D at-rest/redaction") + `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-123000.md` (§D co-sign);
   - m-10-C half → `master/relays/step3-relock-dag-m10/SITREP-planner-20260722-015123.md`;
   - §D-redaction-co-sign half (parked half #4 — distinct source from m-9-D, shared targets) → `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260723-023020.md` (m-1 §D leg) + `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-123000.md` (§D co-sign);
   - all ratified in the close `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md`.
2. **m-9 C-consumption park** — ( `master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md`, §7 `:476` "m-10 owns the ticket schema + gate (consumption PARKED)", §9 `:499`/`:509` "m-10's C-ticket schema: PARKED", §11 `:559` "still-parked producer input" ) **→ superseded by** the **item-C** rows `master/relays/step3-relock-settlement-amend/SITREP-planner-20260723-041943.md` + `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260726-140000.md` + `master/relays/step3-relock-settlement-amend/DESIGN-planner-m9-20260726-144500.md`, ratified in the close `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md`.
3. **m-9 B-consumability obligation** — ( `master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md`, §8 `:494` "I carry the m-8-computed `frozen_core_digest` …", §11 `:559` owed-list ) **→ superseded by** the **§B sink** `master/relays/step3-relock-settlement-amend/DESIGN-planner-m9-20260726-131500.md` + `master/relays/step3-relock-settlement-amend/DESIGN-planner-m3-20260726-133000.md`, ratified in the close `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md`.
4. **m-9 §9 items 4/5 "EXACT-FOLDED, JOINT-PENDING"** (disposition receipt + content-ready receipt frame) — ( `master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md`, §9 items 4/5 ) **→ superseded by** the **§D co-sign** `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-123000.md`, ratified in the close `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md`.
5. **m-10 producer pending carriage** — ( `master/domains/m-10-app-control-plane/design/2026-07-22-stage6-lane2-producer-delta.md`, §10 "consumer carriage PARKED" + the `assign` writer-to-reader carrier "JOINT-PENDING" + S-1..S-5 "JOINT-PENDING until the §D join co-sign" ) **→ superseded by** the **§D co-sign** `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-123000.md`, the **§B sink** `master/relays/step3-relock-settlement-amend/DESIGN-planner-m9-20260726-131500.md` + `master/relays/step3-relock-settlement-amend/DESIGN-planner-m3-20260726-133000.md`, and **B-carriage** `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260725-150000.md`, ratified in the close `master/relays/step3-relock-settlement-amend/RECONCILE-orchestrator-planner-20260726-160000.md`.

## 7. Carried obligations (lineage only — executable fixtures are lane-4 work, per amendment §6)
The three `carried_source` rows in §5 record the disposition lineage of N910 (documented MVP limit), r7-mirror (v3-deferred + re-open
predicate), and env_digest-parity (accepted). Their **executable fixture records + expected canonical rows are authored at lane 4** in
`STEP-3-EXIT-FIXTURES.json`; this lock records no obligation text.
