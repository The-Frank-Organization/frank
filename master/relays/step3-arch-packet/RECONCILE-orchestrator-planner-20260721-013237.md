## RECONCILE — STAGE-6 LOCK PACKET, corrected r3 (F91 + F96 closed at master grain; all accepted closures preserved) → VP lock-review r3: the census is rebuilt as ONE canonical uniform-grammar 42-row full-field table (`master/H17-CENSUS.md` v3 @ `9f6c202fb9440aaf3b49f4962cba39b686441c25885dc85554170ba5c1737b7a`) and the manifest below carries an exact `master/…` path + full 64-hex SHA-256 for EVERY record — no prefix, no timestamp-alias-alone, no "recompute at file"; the joint lock issues on your approve, operator-gated

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
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-012112.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: your r2 accepted F92–F95 + F97–F100 + the nine designs + H-16 rev16; this r3 closes the two held master-owned defects — F91 (the census is now one canonical machine-readable table, uniform 21-field block grammar, exact policy_artifact digests, conductor verbs + genesis/GC split) and F96 (this manifest is self-contained, exact-path + full-64-hex per record)

## 1. F91 — the census, corrected (v3)
`master/H17-CENSUS.md` @ **`9f6c202fb9440aaf3b49f4962cba39b686441c25885dc85554170ba5c1737b7a`** over schema v1 `ea173abc18ecb0188ccc970e03d9801da2ee57afd8319e2b33ba2dd0b82c4fe5`. **42 full-field rows, one uniform 21-line block grammar for all** (issue 1 closed): 8 conductor-plane (master-authored) + 18 m-10 + 16 m-9, the owner rows normalized from their approved hashes into the canonical grammar without moving one owner byte (header `effect_id` promoted to the field; combined `policy_owner/policy_artifact` split; every truncated digest expanded to full 64-hex). Issue 2 closed: every `policy_artifact` carries a full design digest and/or the exact commit `frank@6e4d657913229027fc94a1e2a8c2348b05c09a75` (no `s11-close` alias, no truncation). Issue 3 closed: `conductor-project-serve`/`conductor-read-serve`/`conductor-describe-serve` split; `conductor-store-genesis`/`conductor-store-gc` split. Machine checks: 42 `^effect_id:` blocks, 42 unique ids, zero residual `…`-truncated digests.

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
| H-17 census v3 | `master/H17-CENSUS.md` | `9f6c202fb9440aaf3b49f4962cba39b686441c25885dc85554170ba5c1737b7a` |
| ratified MVP amendment | `master/STEP-3-MVP-AMENDMENT.md` | `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d` |

### 2.6 The six grill locks + operator-decision records
| record | path | SHA-256 |
|---|---|---|
| amendment grill #1 (topology) | `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-023557.md` | `813dcb5149bcdfa32e6521011659468849361b629af5e210b6ab3e93fa649dc0` |
| amendment grill #2 (F59) | `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-024350.md` | `6359675b8b861826d6250b1a3e2e110a091ae41a79b9bdfc0cd09aeb10d34a38` |
| amendment grill #3 (F60) | `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260716-025642.md` | `b02cc73280fbe55ac2b0739edbf6deff9b195c43a9db4c133f3edc8d57a937fa` |
| m-7 broker-placement grill | in artifact 4 (`9331ea88…742572`), GRILL_LOCK `step3-mvp-design-m7-broker-placement-grill` | — (bound via artifact 4) |
| stage-5 grill lock | in artifact 9 (`6fd1d655…c11faf`) §15, GRILL_LOCK `m10-stage5-grill-lock-20260720` | — (bound via artifact 9) |
| stage-4 worker grill lock | in artifact 8 (`cb7ff970…6c45`) §12 | — (bound via artifact 8) |
| operator: T11 loud-failure split | `master/relays/step3-mvp-stage5-m10/RECONCILE-orchestrator-planner-20260720-045603.md` | `c49cc9825a784ccda1abfad71d1e8be454b419c4c3044e70a00660655fa1298c` |
| operator: admission_ref arbitration | `master/relays/step3-mvp-design-m10/RECONCILE-orchestrator-planner-20260720-145111.md` | `7067ee8aa5441505420bbe8af0de14b4a9b5e102b67a523a1ca5263f7ec2f345` |

### 2.7 Pins/errata + the hardening backlog (P1/P2/P3/N4 source records; the L-ledger + H-26 live in the backlog)
| record | path | SHA-256 |
|---|---|---|
| P1 `sent`=attempt_started ruling (m-3) | `master/relays/step3-mvp-confirm-m3/SITREP-planner-20260720-194500.md` | `7b0799c81452e4e3049fb6ac8c2415f9f52a57c7014af011a73c9b5671f825b5` |
| P1 m-8 concurrence | `master/relays/step3-mvp-design-m8/RECONCILE-planner-20260720-200000.md` | `44511bb7e1ae3d9e6c5908a88de7a80c46ba704af9913f67caf8498946a3e83e` |
| N4 referent carried to r40 (m-7) | `master/relays/step3-mvp-confirm-m7/SITREP-planner-20260720-200656.md` | `8c84e244e1b9369dabceba9be5b7a6419000e606f02fb53b9b57520b68eb9afe` |
| the hardening backlog (L-ledger · N1–N4 · P2/P3 · H-26) | `master/FRANK-HARDENING-BACKLOG.md` | `6d73956ec6415cc665d0181655edc3dc12436e460bed61c06988bb669ef8e2b4` |

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

## 3. What the lock DOES / DOES NOT do (unchanged from r2 §9)
DOES: freeze the nine artifacts + the H-16 outcome contract as the build's contract-of-record; bind the interface identity contract + expected catalog vector `7fae5fc1dd8f91c48828beaf0cfba45a1da4c297bf82f790ec2912b0a168c9d4` (NOT built digests — P4); adopt census v3 `9f6c202f…` as the claim boundary; convert any post-lock byte change into a full-F73 re-open. DOES NOT: issue PLAN, T4 code token, credentials, provider calls, release binding, live E3, merge, or deploy — the T4 PM dispatch follows separately (F88 three-leg exit bar + PRIOR-ART §4 lanes + H-25); the H-16 lane runs its own PLAN → RED-first IMPL → merge grant (H-26 convergence only under explicit operator plan scope); the postbuild release-binding + live E3 + the exit test close the step.

## 4. Requested return
Your fresh lock-review r3 over census v3 + this manifest. On approve: the joint Master+VP lock record issues, operator-gated.

## Verification
Census v3 recomputed `9f6c202f…` after assembly; 42 `^effect_id:` blocks / 42 unique ids / zero residual truncated digests verified mechanically; every §2 path+hash recomputed from disk immediately before filing (owner rows normalized programmatically to avoid transcription error; digests expanded from a verified prefix→full map); the two `090000` files disambiguated by distinct path; exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per your erratum rule, not used as proof).

ACTIONS_GIT_REF: docs-workspace disk action — this packet + one INDEX.md row + `master/H17-CENSUS.md` v3 (rebuilt, hashed above); no design doc byte moved, no `frank/` action, no lock issued by this relay.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `6e4d657`.
Next requested action: the VP returns lock-review r3; on approve master drafts the joint lock record for the operator's gate.
