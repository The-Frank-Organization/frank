## RECONCILE — item A routes through the VP FIRST (operator-directed). The six owners are HELD: do NOT nominate on `…-170000` until master releases post-VP-review (addressed hold — a live dispatch's delivered authority needs an addressed revoke, the DAG-R1-F1 lesson). VP: decomposition-review the item-A recipe `44bb27fa…` — completeness · single-ownership · no-foreign-byte-hashing · Tier-HARD/Tier-SOFT decidability + soft-stability soundness · the carried-obligation fixture legs (esp. env_digest-parity @ m-1 §5 :63) · the produce-not-lock boundary. On approve → master releases the hold → owners nominate. On revise → master re-cuts.

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-item-a
PARENT_DISPATCH_ID: step3-relock-item-a
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — holds a just-issued design dispatch pending a VP decomposition-review and routes that review; it moves no ratified/frozen byte, opens no owner action, and locks nothing
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-item-a
IN_REPLY_TO: master/relays/step3-relock-item-a/DESIGN-orchestrator-planner-20260726-170000.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer, m-1.planner, m-2.planner, m-3.planner, m-8.planner, m-9.planner, m-10.planner
CC: operator, m-1.implementer, m-2.implementer, m-3.implementer, m-8.implementer, m-9.implementer, m-10.implementer
SUBJECT: HOLD `…-170000` (item-A recipe) at owners pending VP decomposition-review; VP reviews `44bb27fa7420a00bebd4013ce2bb87fcfc66919c3865c293ebe7657d986d931a` for completeness / single-ownership / no-foreign-byte-hashing / Tier-HARD-SOFT decidability + soft-stability / carried-obligation fixtures / produce-not-lock scope; on approve master releases the hold, on revise master re-cuts

Operator directed a **VP decomposition-review of the item-A recipe before the owners act** — the same pass that caught the real blockers in lane 2 (DAG-R1..R5). Applying it here.

## Owners (m-1/m-2/m-3/m-8/m-9/m-10) — HOLD `…-170000`, addressed
**Do NOT nominate your Tier-HARD `lock_payload` yet.** The item-A opening dispatch `step3-relock-item-a/DESIGN-orchestrator-planner-20260726-170000` is **HELD pending VP decomposition-review.** This is an **addressed hold**, not a reviewer-CC: the DAG-R1-F1 lesson is that a live direct dispatch carries delivered authority a CC cannot revoke, so I revoke the nominate-now instruction to you explicitly here. The recipe itself is unchanged and stands as the review target; only your action is held. **Confirm no nomination has been filed** (as of this hold, none is on the trail — the leak is unexercised, as in lane 2). Master will issue an **addressed RELEASE** on VP approve; you nominate only then.

## VP — decomposition-review requested on the item-A recipe `44bb27fa…`
Review `master/relays/step3-relock-item-a/DESIGN-orchestrator-planner-20260726-170000` @ SHA-256 `44bb27fa7420a00bebd4013ce2bb87fcfc66919c3865c293ebe7657d986d931a` — the decomposition, before the owners nominate:
1. **Completeness** — does the recipe fully specify the deliverables (`STEP-3-INTERFACE-BUNDLE.json` schema · `bundle_sha256` · the `bundle-soft-stability` negative fixture · `STEP-3-EXIT-FIXTURES.json`) and the six-owner routing, with nothing load-bearing left implicit?
2. **Single-ownership** — each owner nominates only **its own** Tier-HARD elements; no owner's `lock_payload` names or hashes another owner's bytes.
3. **No-foreign-byte-hashing** — `bundle_sha256` covers the extracted payloads + the join records + the carried obligations; confirm it hashes no foreign or un-owned content, and that `base_hashes` is provenance-only (not part of the hashed load-bearing set in a way that would break soft-stability).
4. **Tier-HARD/Tier-SOFT decidability + soft-stability** — is the boundary rule **decidable** (an extractor can mechanically separate Tier-HARD from Tier-SOFT, not a prose judgment), and does the claimed **soft-stability property** actually hold (a Tier-SOFT edit leaves `bundle_sha256`; any Tier-HARD change moves it)? This is the load-bearing correctness claim of item A.
5. **The carried-obligation fixture legs** — are N910 (loss cut → no record + `uncertain` disclosed), r7-mirror (v3 + re-open caveat), and **env_digest preimage-parity (recipe @ m-1 §5 `:63`, realized byte-for-byte by m-9 §7 + m-3's E3 observer)** correctly bound as `STEP-3-EXIT-FIXTURES.json` legs — with the corrected locus (NOT m-9 §10)?
6. **DAG / ordering** — the six nominations are asserted independent/parallel over already-settled bases; confirm no cross-owner dependency was missed and the settled joins are referenced without re-opening (r17 §9 items 4/5 + rev16 S-1/S-2/S-4/S-5 normative; bind settled bases, never ancestry).
7. **Scope boundary** — item A **produces** the bundle; **lane 4 locks** it. Confirm the recipe issues no premature interface-lock / DESIGN-lock and stays within produce-not-lock.

## The path either way
- **VP approve** → master issues an **addressed RELEASE** to the six owners → they nominate → master assembles → `bundle_sha256` → soft-stability fixture → freeze `STEP-3-EXIT-FIXTURES.json` → VP item-A (assembled-bundle) review → lane 4.
- **VP revise** → master re-cuts the recipe (inert until a separately-addressed release, as in lane 2's rev2) → re-routes for VP re-review. The owners stay held throughout; no owner acts before the release.

## Boundaries
No DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. No bundle authored, no owner nomination opened (held), no lock issued. The nine settled bases + the three amendments UNMOVED (verified at the item-A open, `…-170000`). **H-12 continues to hard-block external/untrusted/multi-tenant use.**

## Verification
Item-A recipe `…-170000` = SHA-256 `44bb27fa7420a00bebd4013ce2bb87fcfc66919c3865c293ebe7657d986d931a` (bound as the VP review target). No owner nomination exists on the trail (the hold's leak is unexercised). Exact-file lint of THIS relay OK (root-mode historical INDEX/lineage noise disclosed, not used as proof). `git -C frank status` clean at `c78da38`.

ACTIONS_GIT_REF: docs-workspace disk action — this hold + VP-review-routing relay + one INDEX.md row; no design byte moved, no `frank/` action, no lock issued, no owner action opened, no fold performed.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: VP decomposition-reviews the item-A recipe `44bb27fa…` (the 7 criteria above); on approve master issues the addressed RELEASE to the six owners to nominate their Tier-HARD `lock_payload`s; on revise master re-cuts. The owners hold meanwhile — no nomination before the release. H-12 stands.
