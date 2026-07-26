## RECONCILE — STAGE-6 RE-SCOPE r4: the five executability defects fixed at architecture grain — lock digest excludes source provenance, descriptor binds the real env/interpreter, provider-lowering owned by m-8, journal is crash-total, exit gate requires verdict=pass → VP decomposition review r4

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-arch-packet
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — the amendment's contracts are now internally executable; it needs your decomposition review r4, then the operator's re-scope ratification. Joint lock `b7e1f0ef` stays HELD/superseded.
GRILL_REQUIRED: no — the F106 grill is done (§3 GRILL_LOCK); rev4 folds only the five bounded executability corrections and introduces no new product choice
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-043904.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: your r3 was right on all five — real correctness bugs, not wording; rev4 fixes each at architecture grain (soft-stable lock digest, real env+interpreter+apply_patch binding, m-8-owned provider lowering, crash-total journal with a CONTENT_LOST park, and an applicability-AND-verdict gate over frozen fixtures). Review r4.

VERDICT: revise — self-initiated: master returns amendment rev4 with the executability defects resolved

## 1. Your r3 accepted in full; rev4 folds the five fixes
Amendment rev4 `master/STEP-3-STAGE6-AMENDMENT.md` (`1c485e9d8f56e584725b6750bb7de58324f3773503815537213d572a90dad2e9`) supersedes rev3 `419c3793…`. F102 + F104-B stay closed; all five directions you accepted are preserved; no bound design byte moves.

## 2. The five corrections at architecture grain
- **F101 (soft-edit stability):** `bundle_sha256` is now computed ONLY over a `lock_payload` = `{recipe identity + ordered interface_ids + extracted-region digests}`; mixed-source full-file SHAs move to a separate `provenance` section that `--verify` checks but that does NOT feed the constitutional digest. Added the undeclared-marker full-inventory scan + pinned `recipe_sha256`/`bundle_sha256` locations. Ships a **`bundle-soft-stability` negative fixture** proving an out-of-HARD-region edit leaves `bundle_sha256` unchanged while moving `source_sha256`.
- **F103 (real context identity):** env branch stated = the **m-1-sanitized environment** (hardening #1/#7, already frozen — no behavior delta); `env_digest` over the COMPLETE presented set. Added `shell_interpreter_ref{path,version,binary_sha256?}` so `backend_id="ambient"` is never the only impl identity; `tool_impl_ref` narrowed to the wrapper. `apply_patch` split from single-path tools → `canonical_resource` = an **ordered target-set digest**. cwd canonical byte form pinned.
- **F104-E (lowering owned where the bytes are):** split into `logical_surface_digest` (m-9, pre-lowering, m-2 supplies the schema/description component) + `provider_lowered_tools_digest` (**m-8**, over the lowered `tools[]` of the frozen-core body it alone produces, carried on the m-8 terminal beside `frozen_core_digest`); `model_surface_digest` = m-3 joining the two DIGESTS, never m-9 hashing bytes it can't see. Carriage + independent observer derivation defined. No duplication of m-8's translation.
- **F105 (crash-total journal):** added `tool_call arguments` + `replay_envelope` as first-durable blobs (the ticket digest is non-reconstructive; the replay payload is in-memory-only). A monotonic `round_marker` + the round's content blobs + its outcome row commit in ONE m-10 transaction; a crash before it parks UNKNOWN; a resume-time missing/mismatched blob parks typed **`CONTENT_LOST`** — silent resume across a known-effect/lost-content gap is forbidden. Replay-envelope recovery settled (persist, else typed `REPLAY_UNRECOVERABLE`/new-attempt). Retention/GC/orphan-GC/size(§2a)/integrity added.
- **F106 (decidable gate):** every leg now requires E3 `applicable` **AND** typed `verdict=pass` **AND** (where carried) `observed_outcome=pass`; any `fail` fails, any unknown holds. Fixtures byte-bound via a frozen `STEP-3-EXIT-FIXTURES.json` manifest (spec frozen at re-lock, built at T4). Total overhead rule with NO undefined interval (p50 ≤20% PASS · 20–100% HOLD-with-operator-waiver · >100% FAIL) + per-metric p95 ceilings + EXACTLY 30 turns/100 calls + a named baseline. Operability/handoff/crash-honesty are machine predicates over structured fields (surface field-set; second-seat committed relay record; idempotency-keyed exactly-one-EXECUTED observer).

## 3. For your r4 adversarial attention
- The env branch declares no behavior delta by resting on m-1's EXISTING sanitization (#1/#7) — confirm that digesting the m-1-sanitized presented env (rather than newly clearing to an allow-list) is the honest binding you asked for, not a claim that exceeds what m-1 froze.
- `provider_lowered_tools_digest` is scoped to the lowered `tools[]` of the frozen-core body; since `frozen_core_digest` already covers the whole body, confirm the dedicated component digest is additive-not-redundant (it isolates the tool surface for the E3 join) and not a second source of truth.
- `CONTENT_LOST` + `REPLAY_UNRECOVERABLE` are new typed non-resumable states — confirm they compose with the frozen UNKNOWN/park family (they are honest terminals for a content-integrity cut, distinct from an effect-UNKNOWN) rather than needing an m-9/m-10 lifecycle amendment beyond this seam.

## 4. Requested return
Decomposition review r4 over rev4 `1c485e9d…`. On your pass the amendment goes to the operator for the re-scope gate. No PLAN, T4 token, credential, provider call, release binding, live E3, merge, deploy, or out-of-envelope use is requested.

## Verification
Recomputed from disk: amendment rev4 `1c485e9d8f56e584725b6750bb7de58324f3773503815537213d572a90dad2e9`; VP r3 `043904` is the parent. The nine design finals + H-16 rev16 `a349a329…` + census `959b1928…` are UNMOVED (rev4 makes the amendment's own contracts executable; it withdraws no approved mechanism and moves no bound byte — the referenced m-1 §1.3, m-8 §1.1/§2.2, m-10 §D.1, worker §2.5/§3.2 are cited, not edited). Exact-file lint of THIS relay OK (root-mode historical/INDEX noise disclosed per the erratum rule, not used as proof). `frank/` untouched by this authoring.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + amendment rev4 (overwrites rev3 in place; rev3 `419c3793` preserved by hash in the r3 trail) + one INDEX.md row; no design doc byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin (the vendored reviewer basis).
Next requested action: the VP returns decomposition review r4; on pass master routes the amendment to the operator for the re-scope gate.
