# Step-3 Item A — the interface-bundle extraction recipe (re-cut r3 to ratified §4/§7) — **WITHDRAWN**

> **WITHDRAWN 2026-07-27 on ratification of the item-A simplification amendment** (`master/STEP-3-ITEM-A-SIMPLIFICATION-AMENDMENT.md` `3443f73d…`, operator-ratified; ratify relay `step3-relock-item-a/RECONCILE-orchestrator-planner-20260727-130000.md`). The interface-**bundle** mechanism this recipe describes was superseded by the plain byte-bound interface-lock **record** `master/STEP-3-INTERFACE-LOCK.md` (external SHA `cbd1893c…`). **This recipe is not built.** Retained for historical lineage only; the bytes below are the withdrawn r3 (`06e6956e…`).

**Status:** WITHDRAWN (superseded). RE-CUT **r3** by master.orchestrator-planner, folding VP item-A-review-**r2** findings F1–F4 (`step3-relock-item-a/RECONCILE-orchestrator-reviewer-20260726-200000`). Supersedes r2 (`a98e85a1…`).
**Controlling ratified contract:** `master/STEP-3-STAGE6-AMENDMENT.md` rev12 `1125b0a0…` §4 (`:82-108`) + §7 (`:363-398`). This recipe **conforms to** those bytes.
**Item A PRODUCES the bundle; lane 4 LOCKS it.**

---

## Part 1 — the bundle artifact (§4-exact; r1-F1 CLOSED, unchanged)

Top-level `master/STEP-3-INTERFACE-BUNDLE.json`, schema `step3-interface-bundle.v1`:
`{ schema, recipe_version, recipe_sha256, bundle_sha256, lock_payload, provenance }`.
- **`lock_payload`** (ONLY digest input) = `{ recipe_version, recipe_sha256, interfaces:[ {interface_id, extracted_sha256} … sorted ascending by interface_id ] }`. **`bundle_sha256` = SHA-256( JCS(`lock_payload`) )** — excludes mixed-document full-file SHAs; a Tier-SOFT edit does not move it.
- **`provenance`** (integrity only, NOT hashed) = `{ sources:[ {interface_id, source_path, source_sha256, region ∈ "marker"|"whole_file"} … ] }`.
- **Hard-region markers:** literal `<!-- HARD-BEGIN interface_id=<id> recipe=item-a.v1 -->` … `<!-- HARD-END interface_id=<id> -->`; extractor takes the enclosed bytes verbatim; `extracted_sha256` over them. A `whole_file` source's `extracted_sha256` IS its full-file SHA.
- **Extractor/verifier** = `master/tools/extract-interface-bundle.py` (master-authored at assembly; `recipe_sha256` = its own digest; `recipe_version = item-a.v1`). `--verify` recomputes every `source_sha256`, re-extracts each declared span, checks each `extracted_sha256`, runs the **undeclared-marker full-inventory scan**, recomputes `bundle_sha256`.
- **Fail-closed (`--verify` nonzero; no bundle; re-lock blocked):** declared marker span absent · duplicate `interface_id` · `source_sha256` on-disk mismatch · `extraction_recipe_version` mismatch · ill-formed span · **any UNDECLARED `HARD-BEGIN` in a declared source but absent from the manifest**.
- **Tier-SOFT ledger:** `master/SOFT-DESIGN-LEDGER.md`, not `PROTOCOL-DEVIATIONS.md`.

### The `bundle-soft-stability` negative fixture (r2-F1 — now concrete, isolated, no settled-source mutation)
- Files (shipped with the extractor, isolated — NEVER mutate a settled base):
  `master/tests/bundle-soft-stability/fixture-source.md` (a synthetic marker-bearing source: one `HARD-BEGIN/END` region + one out-of-region SOFT paragraph), `expected.json` (the frozen expected result pair), `run.sh`.
- `run.sh` computes a baseline `bundle_sha256` over the fixture source, then applies **two scripted mutations to COPIES** (`sed` scripts committed in the dir):
  - **SOFT mutation** (edit the out-of-region paragraph): asserts `provenance.source_sha256` CHANGES **and** `bundle_sha256` UNCHANGED vs baseline.
  - **HARD mutation** (edit one byte inside the `HARD-BEGIN/END` region): asserts `extracted_sha256` for that interface CHANGES **and** `bundle_sha256` MOVES vs baseline.
- `expected.json` freezes both `{soft: {source_sha256_changed:true, bundle_sha256_changed:false}, hard: {extracted_sha256_changed:true, bundle_sha256_changed:true}}`. The fixture passes only on the exact pair.

---

## Part 2 — the interface manifest (§4 inventory; r2-F1/F2/F3 CLOSED)

**Decomposition rule (kills the r2 overlap/discontiguity finding):** **ONE contiguous HARD marker span per settled source**, delimited by exact **start-heading → end-heading** anchors that wrap the source's normative sections and **exclude its Tier-SOFT tail** (fold/decision logs, boundary sections, revision logs, "what this does NOT do", fixture-obligation sections). One interface per source ⇒ single-ownership + non-overlap are decidable by construction; no two rows draw from the same section; no consumer hashes a producer's bytes. `recipe_version = item-a.v1` for every row. **On RELEASE, the owner places the two markers within the bound anchor range and its pair confirms the enclosed bytes are the right contract** (F2: owners decide the bytes; the extractor mechanically decides which are hashed).

### 2a — the eight owner interfaces (marker span; anchors = the first normative heading through the last, excluding the soft tail)
| interface_id | sole_owner | source_path | HARD span (start-heading … end-heading, exclusive of soft tail) | full source_sha256 |
|---|---|---|---|---|
| `if.m1.lane2` | m-1 | domains/m-1-trust-identity/design/2026-07-22-stage6-lane2-env-redaction.md | `## §1 — Item C: the env sanitization rule` … end of `## §2 — Item D: the redaction boundary…` (before `## §3 — What this delta does NOT do`) | `d34a7c475c5988f12872a5aaf599b201fe56970b766b935bdb42b1d2851ab1ef` |
| `if.m2.logical-component` | m-2 | domains/m-2-forms-determinism/design/2026-07-22-stage6-e-logical-component.md | `## §1 — What this component IS…` … end of `## §5 — The offered assembler helper…` (before `## §6 — Build obligations`) | `c3a8cd61bcbe33a3ede847b29d7416563000c66f949fa2950ea0af305edbbd2c` |
| `if.m2.relay-submit-cell` | m-2 | domains/m-2-forms-determinism/design/2026-07-23-stage6-c-relay-submit-resource.md | `## §1 — The gap, exactly` … end of `## §5 — The amendment cell…` (before `## §6 — Build obligations`) | `5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f` |
| `if.m3.lane2` | m-3 | domains/m-3-observation-evidence/design/2026-07-22-stage6-lane2-e0-e3-delta.md | `## §1 — The exhaustive cut × CARRIER matrix…` … end of `## §5 — F73 consumer confirmations…` (before `## §6 — Scope`) | `651c9aecd1dd27b61d1a5fd3e802b05ee350e1acb848b827ee57b27751eccb97` |
| `if.m8.lane2` | m-8 | domains/m-8-provider-adapters/design/2026-07-22-stage6-BE-digests-addendum.md | `## §0 — Three distinct predicates…` … end of `## §5 — Boundary compliance…` (before `## §6 — Fixture obligations`) | `734e44b7841754dfca56f3a9195695bed1d2f07b42d3acae92ef6b188b47fb53` |
| `if.m9.lane2` | m-9 | domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md | `## §1 — D1: the worker-owned crash-safe session-content log` … end of `## §9 — The PARKED consumer set…` (before `## §10 — Fixture obligations`) | `01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b` |
| `if.m10.producer` | m-10 | domains/m-10-app-control-plane/design/2026-07-22-stage6-lane2-producer-delta.md | `## §1 — D2: the settlement-manifest PRODUCER…` … end of `## §9 — Cross-owner review hooks…` (before `## §10 — PARKED consumer sections`) | `3e3c51926332f3068aad8c995d898a11185b0d35608213a0d97903b6cf8d53f6` |
| `if.m10.be-carriage` | m-10 | domains/m-10-app-control-plane/design/2026-07-23-lane2-be-carriage-row.md | `## §1 — The three carried members…` … end of `## §6 — What m-3 will receive…` (before `## §7 — Fixture obligations`) | `cd17db3204286f4a0b165b1e349d75c5cd9320dc793b021e942275427115e35f` |

### 2b — the five join interfaces (master-owned; r2-F2 restores the missing B-carriage join)
Master authors `master/STEP-3-INTERFACE-JOINS.md` at assembly (VP-reviewed) — five marker spans, one per join, each recording the two-sided/cross-domain binding over settled bytes (references producer interface_ids, hashes no foreign source bytes):
| interface_id | sole_owner | region | binds |
|---|---|---|---|
| `if.join.d-resume` | master | marker (§D) | the §D two-sided resume seam: m-9 r17 `01b885fe…` × m-10 rev16 `3e3c5192…` + m-1 redaction |
| `if.join.b-sink` | master | marker (§B) | the §B two-sided sink: m-9 r17 §8 ⇄ m-3 r24 `651c9aec…` `m3.b_sink.v1` |
| `if.join.b-carriage` | master | marker (§B-carriage) | **the producer→carrier→evaluator chain m-8 r7 `734e44b7…` → m-10 rev3 `cd17db32…` → m-3 r24** (the 5th normative join, r2-F2) |
| `if.join.item-e` | master | marker (§E) | item-E: m-9 r17 `logical_surface_digest` → m-3 r24 R1 binding |
| `if.join.item-c` | master | marker (§C) | item-C: m-9 r17 §7 executor + m-10 rev16 §5 C-ticket + m-1 C-confirm |

### 2c — the exit-fixtures interface (master-owned; r2-F2 gives the carried obligations a digest-bearing home; NO circular dependency)
| interface_id | sole_owner | region | source |
|---|---|---|---|
| `if.exit-fixtures` | master | **whole_file** | `master/STEP-3-EXIT-FIXTURES.json` — **frozen BEFORE the bundle** (Part 3); its full-file SHA is `extracted_sha256`. Its `carried_records` section carries N910 / env_digest-parity / r7-mirror as separately-decidable IDs, so the carried obligations enter `lock_payload` through **this one declared interface's digest**. |

**14 HARD interfaces total** (8 owner + 5 join + 1 exit-fixtures); every load-bearing item present exactly once, sole-owned, non-overlapping, contiguous. The nine settled-base SHAs are the `provenance.source_sha256` column, not raw digest inputs.

---

## Part 3 — `master/STEP-3-EXIT-FIXTURES.json` (§7-exact; r2-F3 CLOSED; frozen BEFORE the bundle)

A closed JSON manifest (schema `step3-exit-fixtures.v1`), frozen and F73/VP-reviewed **before** bundle assembly, then extracted whole_file as `if.exit-fixtures`. Fixtures are BUILT at T4 to this frozen spec (not post-selectable).

### 3a — the six gate legs (predicates bound VERBATIM to ratified §7 `:368-376`; no lossy restatement)
`legs:[ {leg_id, fixture_ids[], predicate_ref} ]` — `predicate_ref` cites the ratified §7 row by leg (the extractor/reviewer binds the exact ratified bytes, never a paraphrase):
- **`xit-gov-1`** — `provider_request_matches_frozen_core ∧ local_invocation_matches_effect_descriptor ∧ relay_record_committed_with_stamped_sender`, each applicable + `verdict=pass`.
- **Durability = `xit-dur-1…xit-dur-5`, ALL must pass** — bound verbatim to §7 `:371`: `xit-dur-1` positive resume from the exact last valid prefix + closed-manifest-union reconcile (`settled_with_content` only) reproducing `resume_prefix_expectation`; `xit-dur-2` the fixed `corruption_cut` → degrade + durable `degraded` disposition/`resume_action`; `xit-dur-3` provider conjunction + **no omission, BOTH missing-half orders** (terminal-first/receipt-absent → exactly one `uncertain`; receipt-first/terminal-absent → the UNKNOWN/PARTIAL row ALSO exactly one `uncertain`; both → exactly one `settled_with_content`, idempotent; positive-then-missing-prefix → `content_lost`+`DEGRADED`; omission mutant FAILS); `xit-dur-4` the three pre-receipt crash cuts → **zero provider/tool/conductor work until the post-commit receipt**, and the selected first-action branch **OBSERVED exactly once AFTER the receipt, zero before**; `xit-dur-5` frame boundary — max legal frame passes, one byte over → the single `resume_frame_overflow` → **TERMINAL `FAILED`, no successor/lease/snapshot/same-run revival**, operator manual `resume_action` projection.
- **`xit-crash-1`** — crashed `tool_calls` row parks `UNKNOWN_TOOL_OUTCOME` + stays; `counter_after_recovery == counter_before_recovery ∧ invocations_after_recovery == 0`.
- **`xit-inj-1`** — induced action has an F59 ticket + `FROM`-stamped record + honest recorded outcome (visibility, not prevention).
- **`xit-ho-1`** — two correctly-stamped records + lineage (origin `FROM` never forged onto the second-seat record).
- **`xit-op-1`** — operator surface exposes `{last_event, stop_reason, unknown_effects[], resume_action}` all present + non-null.

### 3b — the closed per-fixture record schema (§7 `:377-393`) + concrete frozen `sample_weight`s
Per fixture: `{ fixture_id, input_artifact_sha256, fault_injection_point, expected_canonical_rows, observer_id, evidence_locator, sample_weight:{governed_turns, tool_calls} }` **plus** the typed extensions where the leg carries them: `effect_observer_key`+`effect_counter_expectation{counter_before_recovery:1,counter_after_recovery:1,invocations_after_recovery:0}` (`xit-crash-1`), `handoff_expected_records[2]` (`xit-ho-1`), `resume_prefix_expectation{predecessor_turn_id,resumed_round_index,log_prefix_digest,context_digest}` (`xit-dur-1`), `degraded_expectation{corruption_cut,expected_disposition:"degraded",expected_resume_action}` (`xit-dur-2`). Top-level `{ baseline_artifact_digest, baseline_config_digest }`.

**Frozen `sample_weight` assignments (sum = EXACTLY 30 governed turns + 100 tool calls, r2-F3 — concrete, not a promise):**
| fixture | governed_turns | tool_calls |
|---|---|---|
| `xit-gov-1` | 3 | 12 |
| `xit-dur-1` | 3 | 10 |
| `xit-dur-2` | 3 | 10 |
| `xit-dur-3` | 3 | 10 |
| `xit-dur-4` | 3 | 10 |
| `xit-dur-5` | 3 | 10 |
| `xit-crash-1` | 3 | 12 |
| `xit-inj-1` | 2 | 8 |
| `xit-ho-1` | 4 | 6 |
| `xit-op-1` | 3 | 12 |
| **total** | **30** | **100** |

**Content-addressed artifacts + ordering (r2-F3):** the per-fixture `input_artifact_sha256` and top-level `{baseline_artifact_digest, baseline_config_digest}` name artifacts **produced and content-addressed at T4** to this frozen spec; the manifest freezes the fixture SPEC (ids, fault points, expected rows, observers, weights, predicates) and carries each `input_artifact_sha256` as a **declared required field the T4 build must fill and the re-lock must verify present + non-placeholder** — the assembled-manifest F73/VP review REJECTS any unresolved digest, placeholder, unowned observer, or arithmetic-only weight. *(The baseline/inputs are built at T4 because §7 fixes "fixtures BUILT at T4"; the manifest is frozen now with their required-field slots + the frozen expected rows/weights, and lane 4 hashes the frozen manifest. No bundle-first/fixtures-second circularity: `STEP-3-EXIT-FIXTURES.json` is frozen BEFORE the bundle and enters it as `if.exit-fixtures`.)*

### 3c — the overhead budget (r2-F4: ALREADY operator-ratified, immutable for T4)
The p95 ceilings **F59 authorize→consume ≤ 250 ms · relay round-trip ≤ 1000 ms · session-log append ≤ 50 ms** and the total added-wall-clock bands **p50 ≤ 20 % PASS · 20–100 % HOLD · > 100 % FAIL** (HOLD cleared only by a durable operator `HUMAN_GATE` relay citing the measured p50) are **ALREADY operator-ratified in rev12** (`step3-arch-packet/RECONCILE-orchestrator-planner-20260721-165500`) and **immutable for T4**. Item A restates them, it does NOT reopen them; any change is a fresh addressed amendment/ratification, never part of item A. Evidence per leg `{fixture_id, run_id, evidence_locators[], verdict, measured_metrics{}}`.

### 3d — the closed `carried_records` section (r2-F2/F3; separately-decidable IDs, distinct dispositions)
`carried_records:[ {carried_id, gating, disposition, source_locus, fixture_binding} ]`, a closed enum of exactly three IDs:
- **`n910`** — `gating:true`; disposition = a frozen expected cut mapped to `xit-dur-3` (loss cut → **no `m3.b_sink.v1` record**) + `xit-op-1` (m-10 `UNKNOWN_PROVIDER_OUTCOME` → `uncertain` present on the operator surface); `source_locus` = m-3 r24 §3.3 + m-10 rev16 §1; `fixture_binding` = the expected-canonical-rows vector on `xit-dur-3`/`xit-op-1`.
- **`env_digest_parity`** — `gating:true`; disposition = a frozen current fixture under `xit-gov-1`'s `local_invocation_matches_effect_descriptor`: canonical logical env input + expected preimage bytes + a **duplicate-name reject** vector + a **non-UTF-8 pre-spawn reject** vector; `source_locus` = **m-1 §5 `:63` (recipe)** + m-9 §7 (derivation) + m-3 E3 observer (the two-sided byte-for-byte realization); `fixture_binding` = a new `xit-gov-1` sub-vector (m-9-bytes == observer-bytes for one logical set).
- **`r7_mirror`** — **`gating:false`**, `disposition:"deferred_v3"`, `reopen_predicate:"an xit-gov-1-gating E3 predicate requires the m-10 leg to resolve 2a/2b INDEPENDENTLY of the DATA-P reply"`; NOT a current pass fixture (requires the missing independent m-10 resolution first); `source_locus` = m-3 r24 mirror-caveat; `fixture_binding` = none (recorded limit + re-open condition only).

**Honest exit claim (frozen in the manifest):** `T1–T8 live · N910 documented MVP limit · r7-mirror deferred-v3` — never "complete lane-2 coverage."

---

## Sequencing (r2-F2 — fixtures/joins frozen BEFORE the bundle; no circularity)
VP re-review approve → master **addressed RELEASE** → each owner places its two HARD markers within the Part-2a anchor range + its pair confirms the enclosed bytes → **master authors, in order:** (1) `master/STEP-3-INTERFACE-JOINS.md` (the five join sources, VP-reviewed) · (2) `master/STEP-3-EXIT-FIXTURES.json` (Part 3, F73/VP-reviewed, FROZEN) · (3) `master/tools/extract-interface-bundle.py` + the `bundle-soft-stability` fixture · (4) run `--verify` → assemble `master/STEP-3-INTERFACE-BUNDLE.json` + compute `bundle_sha256` **over all 14 interfaces incl. the frozen `if.exit-fixtures`** → **F73 review of the assembled bundle** → **lane 4** (re-lock over `bundle_sha256` + whole-file-hard owner contracts + Master+VP interface-lock) → **lane 5** (T4 builds the fixtures + input/baseline artifacts to this frozen spec, filling every `input_artifact_sha256`). The ARCHITECTURE.md D7/`relay.submit` consolidation folds at bundle time. **H-12 hard-blocks external use throughout; item A issues no DESIGN-lock/PLAN/T4 token.**
