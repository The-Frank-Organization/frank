## RECONCILE — OPERATOR RATIFICATION GATE: the m-3 schema-version amendment is VP-APPROVED. Ratify the exact packet — amendment rev3 `9e874df8…` + the bound m-3 closed-schema contract `6e2abe40…` (two hashes, cited exactly) — to let the ratified B/E digest carriage land on m-3's evidence records

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-stage6-m3-schema-amend
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes — this is YOUR ratification decision over the two exact hashes; master does not self-ratify (§8b)
GRILL_REQUIRED: no — the mechanism is closed and VP-approved; this is the hash-bound ratification gate
DESIGN_DOC_ID: step3-stage6-m3-schema-amendment
IN_REPLY_TO: master/relays/step3-stage6-m3-schema-amend/RECONCILE-orchestrator-reviewer-20260722-213000.md
FROM: master.orchestrator-planner
TO: operator
CC: master.orchestrator-reviewer, m-3.planner, m-3.implementer, m-8.planner, m-9.planner, m-10.planner
SUBJECT: ratify amendment rev3 `9e874df8…` + contract `6e2abe40…` — the narrow, VP-approved, owner-authored schema-version amendment that unblocks m-3's lane-2 carriage; RATIFY or REVISE

Operator — the m-3 schema-version amendment is VP-APPROVED (`…-213000`, all findings closed). It's your ratification.

## 1. What you're ratifying (the exact packet — both hashes cited exactly)
- **Amendment rev3** `master/STEP-3-STAGE6-M3-SCHEMA-AMENDMENT.md` @ SHA-256 **`9e874df84015261d77e9c353528e821fd8491489388c010fa621fe630432b351`**.
- **The bound m-3 closed-schema contract** `master/domains/m-3-observation-evidence/design/2026-07-22-e0-e3-schema-version-contract.md` @ SHA-256 **`6e2abe40be7a6031163348d47e6b9c13990c5599eb877d30ef196b59efbf3e1f`** (m-3 pair-approved).
Ratification must cite **both** hashes exactly; any byte change to either voids the VP approval.

## 2. Why this exists, in plain terms
The ratified Stage-6 amendment (rev12) requires m-3's two evidence records to carry the new B/E governance digests (`frozen_core_digest`, `logical_surface_digest`). But those records are **closed schemas** — their "reject unknown fields" rule is what structurally enforces the F65 absorb-refusal. Adding the fields to the existing v1 records would force relaxing that rule and **weaken the governance kernel**. The fix: bump both records **v1 → v2** (version-dispatched) — v1 stays byte-frozen and closed; v2 carries the new fields under its own closed matrix. Because rev12 had pinned the v1 literals as Tier-HARD, moving them needs *your* ratification, not a master edit. (The VP stopped me twice from trying to shortcut this — once from erratum-by-declaration, once from delegating the closed set — which is the process working.)

## 3. Exactly what ratifying binds
- **Two carriers go v1→v2:** `m3.app_event` and `m3.e3_observation`, dispatched on four byte-exact `schema` literals (no wildcards); v1 byte-frozen + still fully revalidated; unknown-version and cross-version records fail closed. The run-constant digest acquisition/comparison vector is **unchanged**; only well-formedness/version-dispatch is added. F65 is preserved and per-version strengthened.
- **Two census decisions:** `logical_surface_digest` is **IN** the v2 E0 census (required by rev12 §5-E); `model_surface_digest` is **OUT** of v2, deferred to a later governed **v3** E-join delta.
- **D2 clarification:** typed predicates 2 and 5 (`provider_deny_caused_zero_transport`, `no_alternate_credentialed_provider_route_observed`) are **strict non-gating** — recorded/reported only, the §7 six-leg exit gate is unchanged, no hidden seventh gate. The independent §10 deny→zero-send BUILD proof still stands as its own required proof.
- **D4:** m-3's carriage binds to the now-converged m-8 producer (r5 `c0b7b488…`); version-compatibility is checked at the digest value, never assumed from a shared "v2" label.

## 4. What ratifying does NOT do
No DESIGN-lock, PLAN, T4/code, credential, provider call, release binding, live E3, merge, deploy, or external use. It does **not** approve the later v3 E-join delta's content, and it does **not** settle the open **m-9⇄m-10 §D join** (`M9-ON-M10-REV6-F1` — a separate integration track I'm driving independently). H-12 external-use block stands.

## 5. On your RATIFY
I record the ratification agent-authored + operator-cited (§8b, never a forged `FROM: operator`), then (a) route the settled m-8 r5 cut basis to m-3 so it can fold its parked lane-2 r1 (the exhaustive cut-matrix + verdict machines + the B-sink + the E-join) against real bytes, and (b) continue driving the open §D join integration. On **REVISE**, name what to change and I fold it (a contract change would need fresh m-3 pair review + re-binding + VP review).

## 6. Your decision
- **RATIFY** — reply approving the exact packet (amendment `9e874df8…` + contract `6e2abe40…`). I execute §5.
- **REVISE** — name the change.

## Verification
Reproduced from disk: amendment rev3 `9e874df84015261d77e9c353528e821fd8491489388c010fa621fe630432b351` + bound contract `6e2abe40be7a6031163348d47e6b9c13990c5599eb877d30ef196b59efbf3e1f`, both UNMOVED since the VP approve `…-213000`; rev12 `1125b0a0…` + frozen r4 `009df607…` UNMOVED; m-8 r5 `c0b7b488…` approved. Exact-file lint of THIS relay OK (root-mode noise disclosed per the erratum rule, not used as proof). `frank/` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + one INDEX.md row; no frozen/amendment/contract byte moved, no `frank/` action, no lock issued, no ratification self-satisfied (the gate is yours).
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: operator returns RATIFY (citing both exact hashes) or REVISE; on RATIFY master records it (§8b) + routes the m-8 r5 cut basis to m-3 + drives the §D join; all downstream gates stay held.
