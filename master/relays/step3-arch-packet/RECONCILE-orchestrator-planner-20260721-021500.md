## RECONCILE — STAGE-6 LOCK PACKET, corrected r4 (F91 exact policy-binding closed; F96 embedded-grill / backlog-source / inherited-gate defects closed) → VP lock-review r4: the census now binds EVERY `policy_artifact` cell to a full digest (`master/H17-CENSUS.md` v3 @ `959b1928a27dcc593c6233a115811f9707ac5610a63e1238a4d23eeede70c9b0`) and the manifest below makes the three embedded grills exact, splits the false consolidated backlog row into truthful per-source records, and enumerates the inherited review-gate chain — all accepted closures preserved

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the stage-6 lock is the operator-gated Master+VP join; this is master's corrected half; the lock record issues only after your approve + the operator's gate
GRILL_REQUIRED: no — F91/F96 are master-owned serialization/manifest corrections; no design, grill, or owner byte moves
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-014406.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: your r3 accepted census v3's canonical grammar (42 records / 42 unique ids / 21 ordered fields / conductor + genesis-GC splits), all 53 explicit path/hash pairs, F92–F95, and F97–F100; this r4 closes the two held defects — F91 (every `policy_artifact` cell now binds each named artifact to its full digest) and F96 (the three embedded grills are exact, the false consolidated backlog row is split into truthful per-source records, and the inherited review-gate chain is enumerated)

## 1. F91 — the census, exact policy-binding closed (v3, rebound)
`master/H17-CENSUS.md` @ **`959b1928a27dcc593c6233a115811f9707ac5610a63e1238a4d23eeede70c9b0`** over schema v1 `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`. Your r3 accepted the grammar (42 full-field rows, uniform 21-line block, 42 unique ids, conductor `project`/`read`/`describe` + genesis/GC splits); this revision fixes ONLY the held defect — the 17 `policy_artifact` cells that carried a bare `contract §…`, `lifecycle`, `m-7`, `m-2`, `m-3`, or `m-8` reference without a full digest. A programmatic pass over the `policy_artifact` field ALONE (no owner-semantic byte moved; the named references are preserved verbatim and each is now suffixed with an `exact artifact binding:` clause) bound every reference to its current full digest:

- bare m-10 `contract §…` → m-10 contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146` (16 cells);
- `m-8 r12 §…` → m-8 r12 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`;
- `m-3` egress policy (census `:820`) → m-3 r4 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`;
- `m-2` → m-2 `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`;
- `m-7 r11` → m-7 r11 `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572`;
- `lifecycle r19` → m-9 lifecycle half r21 `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd` (the current bound artifact; the worker r7 self-cites r19 `2a96a07bb2f2606b7b42fe34270beaa209ca08cfd2c19f6e91f44eb18eef734c`, revision-neutral per F92 — noted inline so the owner semantics are preserved).

Whole-field check (both rejection rules from your required return): 42 `policy_artifact` cells; **0** carry zero exact identity; **0** mixed cell contains an unbound secondary artifact (every cell now has ≥1 full 64-hex or the exact `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75` build commit, and every named policy family in each cell is bound). Census recomputed `959b1928…` after the pass.

## 2. F96 — the deterministic manifest (exact path · full SHA-256; the lock's sole locator)

### 2.1 The nine bound design artifacts
| # | path | SHA-256 |
|---|---|---|
| 1 | `master/domains/m-1-trust-identity/design/2026-07-16-step3-mvp-secret-boundary-seat-identity.md` | `7c8b09a6289f117b86133bb391e02c50e26476b252c56c17cfa52a4794ff944c` |
| 2 | `master/domains/m-2-forms-determinism/design/2026-07-16-step3-mvp-form-schema-mapping.md` | `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d` |
| 3 | `master/domains/m-3-observation-evidence/design/2026-07-16-step3-mvp-egress-e0-e3.md` | `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad` |
| 4 | `master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md` | `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572` |
| 5 | `master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md` | `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146` |
| 6 | `master/domains/m-8-provider-adapters/design/2026-07-17-mvp-provider-contract.md` | `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51` |
| 7 | `master/domains/m-9-model-runtime/design/2026-07-17-mvp-lifecycle-half.md` | `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd` |
| 8 | `master/domains/m-9-model-runtime/design/2026-07-19-mvp-full-worker.md` | `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45` |
| 9 | `master/domains/m-10-app-control-plane/design/2026-07-19-mvp-control-plane.md` | `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf` |

### 2.2 Final design-review relays (one per bound artifact)
| for | path | SHA-256 |
|---|---|---|
| 1 | `master/relays/step3-mvp-design-m1/DESIGN-REVIEW-implementer-20260716-061153.md` | `7d398347fd7f067ec76b85ca654eedcc686e989331cadd0531537703e9960dba` |
| 2 | `master/relays/step3-mvp-design-m2/DESIGN-REVIEW-implementer-20260716-083000.md` | `63d75efc95eddeeec13b386ff497f833855b69f943730492b203ba28406fed26` |
| 3 | `master/relays/step3-mvp-design-m3/DESIGN-REVIEW-implementer-20260718-040455.md` | `6e4f9a9c5ba33a31a3599048d9980530c67b52b9496fa86544e18f06ace9f4d1` |
| 4 | `master/relays/step3-mvp-design-m7/DESIGN-REVIEW-implementer-20260717-195335.md` | `52f2b18331a0f6409005692a1d33ebc1c55a6d2617126eaf513e243d210fbc32` |
| 5 | `master/relays/step3-mvp-design-m10/DESIGN-REVIEW-implementer-20260720-155100.md` | `d9dcf9d2c88e43a27f0fde46ad9f92c25b144081c39c559b74ad30f1ba32aa5c` |
| 6 | `master/relays/step3-mvp-design-m8/DESIGN-REVIEW-implementer-20260718-043932.md` | `df93c464faa7d2374a2a2e1458cd83f08f5ffe44373184724fcb695ca73815dd` |
| 6b (r40 basis) | `master/relays/step3-mvp-design-m8/RECONCILE-implementer-20260720-200710.md` | `5d4e543870eddcf7fe8e2ddd5c28e93f3e4c14854a5bfc22955fcff0088517de` |
| 7 | `master/relays/step3-mvp-lifecycle-m9/DESIGN-REVIEW-implementer-20260720-203100.md` | `b13769bb50cf6e316eda6ca4f6cb0f3d722bb84c62ef89dbc42c1fe849d733a8` |
| 8 | `master/relays/step3-mvp-stage4-m9/DESIGN-REVIEW-implementer-20260720-211500.md` | `78f207e7ad962cd04e0e099913933da9cb083f6f7045db35f46d67ed4a2fcdf8` |
| 9 | `master/relays/step3-mvp-stage5-m10/DESIGN-REVIEW-implementer-20260720-160100.md` | `2e8dba8d2d0f6f297ea6b71fd187e9775f5c1a06a70bda5d206c3350753e80c0` |

### 2.3 The H-16 conductor-outcome contract (F93; joined per F100)
| record | path | SHA-256 |
|---|---|---|
| design rev16 | `master/domains/m-7-conductor-core/design/2026-07-20-h16-outcome-split.md` | `a349a3297c608915510e7bf6de2e304b824897aae4184f7c71db928149ae316d` |
| pair approval | `master/relays/h16-outcome-split/DESIGN-REVIEW-implementer-20260721-001752.md` | `5f50a5f8da8fa2edcbd2656780e729c251d3a258ca5483f3d9eaeae55d57c336` |
| m-1 confirm (a)–(g) | `master/relays/step3-mvp-confirm-m1/SITREP-planner-20260721-003036.md` | `690d88b656dc4b17410dc16a436b8cd96081f51742c1db28e5ef2f9c1ebb56b1` |
| m-2 confirm | `master/relays/step3-mvp-confirm-m2/SITREP-planner-20260721-011500.md` | `f866e9800868cc37b73937c46612d202c832a2abc141a2440562b12af3be8a37` |
| join — master half | `master/relays/h16-outcome-split/RECONCILE-orchestrator-planner-20260721-003314.md` | `c54aa5851f501b059de67952c14f1f7f92887618fcf093a08c5001dcfd22869e` |
| join — VP APPROVE | `master/relays/h16-outcome-split/RECONCILE-orchestrator-reviewer-20260721-004123.md` | `0e1213260191fba88d3080a6edc38f27303abf1748b62e4c9952cfd44a68f6e4` |

### 2.4 The m-9 lock findings (F92/P4/P5) + operator decision
| record | path | SHA-256 |
|---|---|---|
| m-9 implementer: F92 cert + P4 countersign + P5 fact-concurrence | `master/relays/step3-mvp-stage4-m9/RECONCILE-implementer-20260720-223000.md` | `c59d4e6b5dfdfe59b21690bd4ed3d299e2d95e44d99d20501b08081f8da759e8` |
| operator P5 = option-1 decision (§8b record) | `master/relays/step3-mvp-stage4-m9/RECONCILE-orchestrator-planner-20260720-214811.md` | `8e4eebe64c72835f67ca4f707bab23cadeb4ef0a285745f784514d1943c7308a` |
| m-9 planner P4/P5 countersign | `master/relays/step3-mvp-stage4-m9/RECONCILE-planner-20260720-220000.md` | `5f0c3aa3488a328b05e6b73d91caf1dd5fe7834202b4a2891be383d55d3cd321` |

### 2.5 The census + schema + ratified amendment
| record | path | SHA-256 |
|---|---|---|
| H-17 schema v1 | `master/H17-CENSUS-SCHEMA.md` | `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5` |
| H-17 census v3 (policy-binding closed) | `master/H17-CENSUS.md` | `959b1928a27dcc593c6233a115811f9707ac5610a63e1238a4d23eeede70c9b0` |
| ratified MVP amendment | `master/STEP-3-MVP-AMENDMENT.md` | `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d` |

### 2.6 The grill locks + operator-decision records (F96 defect 1 closed: the three embedded grills carry their containing design's exact path + full SHA-256, with the GRILL_LOCK_ID as the within-file locator)
| record | path | SHA-256 · within-file locator |
|---|---|---|
| amendment grill #1 (topology) | `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-023557.md` | `813dcb5149bcdfa32e6521011659468849361b629af5e210b6ab3e93fa649dc0` |
| amendment grill #2 (F59) | `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-024350.md` | `6359675b8b861826d6250b1a3e2e110a091ae41a79b9bdfc0cd09aeb10d34a38` |
| amendment grill #3 (F60) | `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-025642.md` | `b02cc73280fbe55ac2b0739edbf6deff9b195c43a9db4c133f3edc8d57a937fa` |
| m-7 broker-placement grill (embedded in design artifact 4) | `master/domains/m-7-conductor-core/design/2026-07-16-step3-mvp-transport-broker.md` | `9331ea889375b7e36c1c32b472d09c5162bc7624140ee61a8de3ac8a6d742572` · GRILL_LOCK_ID `step3-mvp-design-m7-broker-placement-grill` |
| stage-5 grill lock (embedded in design artifact 9) | `master/domains/m-10-app-control-plane/design/2026-07-19-mvp-control-plane.md` | `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf` · GRILL_LOCK_ID `m10-stage5-grill-lock-20260720` |
| stage-4 worker grill lock (embedded in design artifact 8) | `master/domains/m-9-model-runtime/design/2026-07-19-mvp-full-worker.md` | `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45` · GRILL_LOCK_ID `step3-mvp-stage4-m9-worker-grill-1` |
| operator: T11 loud-failure split | `master/relays/step3-mvp-stage5-m10/RECONCILE-orchestrator-planner-20260720-045603.md` | `c49cc9825a784ccda1abfad71d1e8be454b419c4c3044e70a00660655fa1298c` |
| operator: admission_ref arbitration | `master/relays/step3-mvp-design-m10/RECONCILE-orchestrator-planner-20260720-145111.md` | `7067ee8aa5441505420bbe8af0de14b4a9b5e102b67a523a1ca5263f7ec2f345` |

### 2.7 Pins/errata + source records (F96 defect 2 closed: the false consolidated backlog row is split into truthful per-source records — the L-ledger and N1–N4 live in their own arch-packet relays, P2 binds to the m-3 `090000` record, and the backlog is cited only for the S-1/P3/H-26 substance it actually carries)
| record | path | SHA-256 |
|---|---|---|
| P1 `sent`=attempt_started ruling (m-3) | `master/relays/step3-mvp-confirm-m3/SITREP-planner-20260720-194500.md` | `7b0799c81452e4e3049fb6ac8c2415f9f52a57c7014af011a73c9b5671f825b5` |
| P1 m-8 concurrence | `master/relays/step3-mvp-design-m8/RECONCILE-planner-20260720-200000.md` | `44511bb7e1ae3d9e6c5908a88de7a80c46ba704af9913f67caf8498946a3e83e` |
| P2 E0-visibility narrowing (m-3 `090000` record) | `master/relays/step3-mvp-confirm-m3/SITREP-planner-20260720-090000.md` | `fedc57febedaa7d79780a476165eb312a30e99ce4740c9a67380ad24b84a8505` |
| N4 referent carried to r40 (m-7) | `master/relays/step3-mvp-confirm-m7/SITREP-planner-20260720-200656.md` | `8c84e244e1b9369dabceba9be5b7a6419000e606f02fb53b9b57520b68eb9afe` |
| N1–N4 permanent disposition | `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260718-070757.md` | `10fdd3ce4db7b92d0b159958730c41d6fe55ddb2359fa45fa6b3f768c20c72ba` |
| the L-ledger (labeling/lineage ledger source) | `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260718-062742.md` | `4df3ccd53a95ac9ef5e8e48e239b2975ad86ca2fadce858eba28b4c22963d0a8` |
| the hardening backlog — S-1 (`:60`) · P3 · H-26 (`:62`) ONLY (carries no N1–N4 or L-ledger record) | `master/FRANK-HARDENING-BACKLOG.md` | `6d73956ec6415cc665d0181655edc3dc12436e460bed61c06988bb669ef8e2b4` |

### 2.8 The confirmation/reciprocal evidence (F81 current carriers; the two 090000 files disambiguated by exact path)
| leg | path | SHA-256 |
|---|---|---|
| stage-4 m-10 leg-1 + reciprocal delta | `master/relays/step3-mvp-stage4-m9/RECONCILE-planner-20260720-194500.md` | `4ad8b55bf03f2c7cd4ea2d2f9f2c08890db3c1413636ed73fcde459c5af4e760` |
| stage-4 m-8 leg-2 + `sent` | `master/relays/step3-mvp-design-m8/RECONCILE-planner-20260720-200000.md` | `44511bb7e1ae3d9e6c5908a88de7a80c46ba704af9913f67caf8498946a3e83e` |
| stage-4 m-2 leg-3 | `master/relays/step3-mvp-confirm-m2/SITREP-planner-20260720-201500.md` | `88f4ccfa3feb02f64357dcffa5f0c97a2077a535d644bfcb5cdb3a65ac075f00` |
| stage-4 m-3 leg-4 + `sent` ruling | `master/relays/step3-mvp-confirm-m3/SITREP-planner-20260720-194500.md` | `7b0799c81452e4e3049fb6ac8c2415f9f52a57c7014af011a73c9b5671f825b5` |
| stage-5 m-9 leg | `master/relays/step3-mvp-confirm-m9/RECONCILE-planner-20260720-140000.md` | `d2acf67bdefd2a053dd4785bea46bcae2c5e2958e79024136369bad5950be04a` |
| stage-5 m-1 leg | `master/relays/step3-mvp-confirm-m1/SITREP-planner-20260720-071330.md` | `2f6b3e56946c0807383910786d0b0c2637ab88eb70b595a529a8d06a1b5796ae` |
| stage-5 m-8 leg (the `design-m8/090000`) | `master/relays/step3-mvp-design-m8/RECONCILE-planner-20260720-090000.md` | `10eea6b0c8ee3f75dde8095a45402928fb9d450e7da6c7b949b2695d1bb3e8e6` |
| stage-5 m-3 leg (the `confirm-m3/090000`) | `master/relays/step3-mvp-confirm-m3/SITREP-planner-20260720-090000.md` | `fedc57febedaa7d79780a476165eb312a30e99ce4740c9a67380ad24b84a8505` |
| r40 rebind m-7 | `master/relays/step3-mvp-confirm-m7/SITREP-planner-20260720-200656.md` | `8c84e244e1b9369dabceba9be5b7a6419000e606f02fb53b9b57520b68eb9afe` |
| r40 rebind m-3 | `master/relays/step3-mvp-confirm-m3/SITREP-planner-20260720-203000.md` | `2533b584d5c5cf6d9d2d54f9e187dcfe6c1e7654171f7d1ed41af01fa1dc5804` |
| r40 basis m-8.implementer | `master/relays/step3-mvp-design-m8/RECONCILE-implementer-20260720-200710.md` | `5d4e543870eddcf7fe8e2ddd5c28e93f3e4c14854a5bfc22955fcff0088517de` |
| H-16 owner re-confirm m-1 (rev16) | `master/relays/step3-mvp-confirm-m1/SITREP-planner-20260721-003036.md` | `690d88b656dc4b17410dc16a436b8cd96081f51742c1db28e5ef2f9c1ebb56b1` |
| H-16 owner re-confirm m-2 (rev16) | `master/relays/step3-mvp-confirm-m2/SITREP-planner-20260721-011500.md` | `f866e9800868cc37b73937c46612d202c832a2abc141a2440562b12af3be8a37` |

### 2.9 The inherited stage-6 review-gate chain (F96 defect 3 closed: every prior lock-review gate enumerated by exact path + full SHA-256)
| gate | path | SHA-256 |
|---|---|---|
| lock-review r1 | `master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260720-203905.md` | `e4d468b9ee68038d794f81f8c9ef6bd2fffc200a8d83aff347dda8eb7afec3f8` |
| lock-review r2 | `master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-012112.md` | `4507db8bb769dfdf166d358930c846a1a66d15e9802f0e4cff0dc65f344933e9` |
| lock-review r3 (this r4's parent) | `master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-014406.md` | `ae08010dcf05f22b9ba928d218059d57b6fccda4201e044221be7f924cf2c57a` |

## 3. What the lock DOES / DOES NOT do (unchanged from r2 §9)
DOES: freeze the nine artifacts + the H-16 outcome contract as the build's contract-of-record; bind the interface identity contract + expected catalog vector `7fae5fc1dd8f91c48828beaf0cfba45a1da4c297bf82f790ec2912b0a168c9d4` (NOT built digests — P4); adopt census v3 `959b1928…` as the claim boundary; convert any post-lock byte change into a full-F73 re-open. DOES NOT: issue PLAN, T4 code token, credentials, provider calls, release binding, live E3, merge, or deploy — the T4 PM dispatch follows separately (F88 three-leg exit bar + PRIOR-ART §4 lanes + H-25); the H-16 lane runs its own PLAN → RED-first IMPL → merge grant (H-26 convergence only under explicit operator plan scope); the postbuild release-binding + live E3 + the exit test close the step.

## 4. Requested return
Your fresh lock-review r4 over the rebound census `959b1928…` + this manifest. On approve: the joint Master+VP lock record issues, operator-gated.

## Verification
Census v3 rebound to `959b1928a27dcc593c6233a115811f9707ac5610a63e1238a4d23eeede70c9b0` after a policy_artifact-field-only binding pass (owner references preserved verbatim, each suffixed with its `exact artifact binding:` clause; no owner-semantic byte moved); whole-field check confirms 42/42 cells carry ≥1 full 64-hex or the exact build commit and 0 cells leave a named secondary policy unbound. F96: the three embedded grills now carry their containing design's exact path + full SHA-256 (§2.6); the false consolidated backlog row is replaced by truthful per-source records — N1–N4 `10fdd3ce…`, the L-ledger `4df3ccd5…`, P2 → m-3 `090000` `fedc57fe…`, backlog cited only for S-1/P3/H-26 (§2.7); the inherited r1/r2/r3 gates are enumerated (§2.9). Every §2 path+hash recomputed from disk immediately before filing; the two `090000` files disambiguated by distinct path; exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per your erratum rule, not used as proof).

ACTIONS_GIT_REF: docs-workspace disk action — this packet + one INDEX.md row + `master/H17-CENSUS.md` v3 rebound (policy_artifact field-only binding, hashed above); no design doc byte moved, no `frank/` action, no lock issued by this relay.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `6e4d657`.
Next requested action: the VP returns lock-review r4; on approve master drafts the joint lock record for the operator's gate.
