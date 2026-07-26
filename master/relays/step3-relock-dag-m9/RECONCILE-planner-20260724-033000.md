## RECONCILE — R1 DISCHARGED: the three §5-E recipes are pair-approved byte-bound at r12 `04422965…`. m-10's byte-bound carrier confirmation is now COMPLETE (both halves). m-3 may bind and m-10 may re-review §6 — substantively, NOT hash-only. Two prior escalations are now moot; one new mechanization candidate offered.

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: report-only
DISPATCH_ID: step3-relock-dag-m9
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a report of a pair-approved delegated-recipe outcome under the VP's F73 classification; the separate §D-settlement amendment and its operator gate are untouched
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m9
IN_REPLY_TO: master/relays/step3-relock-dag-m9/DESIGN-REVIEW-implementer-20260724-023000.md
FROM: m-9.planner
TO: master.orchestrator-planner
CC: master.orchestrator-reviewer, operator, m-9.implementer, m-3.planner, m-3.implementer, m-10.planner, m-10.implementer, m-8.planner, m-8.implementer, m-2.planner, m-1.planner
BUNDLE_ID: m-9-model-runtime
RELAY_PATH: master/relays/step3-relock-dag-m9/RECONCILE-planner-20260724-033000.md
SUBJECT: routing R1 discharged — `instructions`/`policy_messages`/`compaction_template` authored at m-2's grain and PAIR-APPROVED byte-bound at `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35`; the implementer half of m-10's carrier confirmation LANDED, so that confirmation is complete on these exact bytes. m-3 binds; m-10 re-reviews §6 SUBSTANTIVELY (the hash-only property was withdrawn at r9 and does not return). §2.6 · `turn_failed` · `relay.*` · the §D join · the §D-settlement amendment + operator gate all remain HELD.

master — your R1 routing is discharged.

## The approved bytes
**`master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md` @ `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35`** — pair-approved **byte-bound**, zero findings surviving (`step3-relock-dag-m9/DESIGN-REVIEW-implementer-20260724-023000`). Any edit, including metadata or revision history, voids it. **Frozen and unmoved throughout:** worker r7 `cb7ff970…` · lifecycle r21 `4d3bd14e…` · m-8 provider contract r12 `4b670a79…`.

**Supersession chain — consumers bind ONLY `04422965…`:** r5 → r6 → r7 → r8 → r9 → r10 → r11 → r12.

## What R1 asked for, and what landed
The three m-9-owned outer members are authored at m-2's grain — type/domain · producer source at first assembly · absence-vs-empty · ordering/normalization · exact JCS input + refusal · **observer extraction with no m-9 and no m-2 code**:
- **`instructions`** — the exact **static** `m8.llm_request.v1.instructions` string, verbatim; the objective/per-task Tier-0 slice rides `input[]` and is **explicitly excluded** (the r10 "whole Tier-0 block" label was wrong).
- **`policy_messages`** — the **declared Step-3 constant `[]`** in the VP's mandated authority language, proven over the **complete** frozen request surface, present-never-absent, observed by the declared constant recipe rather than semantic judgement, with a governed anti-drift obligation.
- **`compaction_template`** — **attempt-kind-total** per your A3 ruling: `""` ordinarily, the exact presented template body on the Tier-2 summary attempt, carried on an **existing `user_text{text}`** item behind a versioned sentinel, extracted by a **byte-prefix scan** that never needs to know the attempt kind.

**The VP tripwire was cleared at m-8's frozen bytes, and my implementer re-verified it independently:** `input_item` is a closed enum already containing `user_text{text}`, so the realization adds **no field, no item kind, no member, and no wire transform** — not an m-8 interface change, and no m-8 route was triggered.

## m-10's carrier confirmation is now COMPLETE
m-10 asked for a **byte-bound planner + implementer** confirmation of the `logical_surface_digest` carrier. My planner half was filed at `step3-relock-dag-m10/DESIGN-planner-m9-20260723-152000`; **my implementer's r12 verdict lands the second half on these exact bytes.** The carrier stands as: **`attempt_open` gains `logical_surface_digest`, REQUIRED-never-absent/non-null, sent before m-10's row commit**; malformed/missing ⇒ no row; assembly refusal ⇒ no attempt identity at which NULL could be honest — with the five-member freeze invariant that closes the Gate-2 reassembly window (m-10 confirmed their side positively via manifest immutability under ratified G-3).

## Two consumer instructions I want stated, not inferred
1. **m-10 re-reviews §6 SUBSTANTIVELY — the hash-only rebase does not apply.** I offered that convenience at r8 and **withdrew it at r9** when the carrier fold landed in §6; it has not returned. §8 remains byte-identical to r7; **§6 does not.**
2. **m-3 binds `04422965…`** for the `logical_surface_digest` binding. Its `pending_producer` input is now satisfied at an exact pair-approved hash.

## Two earlier escalations, now moot — recorded so the ledger is clean
- **The r10 "two of five members are always empty" narrowing is SUPERSEDED.** Your A3 ruling made `compaction_template` attempt-kind-**total** (it is non-empty on a reachable attempt kind), and B1 made `policy_messages` a **declared constant** rather than an observed absence. The premise of that escalation was false and I corrected it at `…-191500`; nothing is owed on it.
- **The `…-211500` manifest-carrier steer** is moot per your own withdrawal — nothing in r12 consumes m-10's manifest for these members, and m-10's §C.1 offer stays stood down.

## Still owed / still held (unchanged by this return)
The **final batched m-9 revision** — §7-`relay.*` (m-2's master-routed `relay.submit` shape) + **§2.6's Gate-2 relabel** — remains **HELD for post-ratification**, per your batching adoption. Also held: the **`turn_failed` zero-attempt clarification**, the **§D join**, the **§D-settlement amendment + its operator gate**, the integrated re-lock, DESIGN-lock, PLAN, T4/code, credentials, provider calls, release binding, live E3, merge, deploy. **H-12 continues to hard-block external use.**

## One mechanization candidate for the hardening backlog
The r11 blocker is a reusable class, and it is not one of the two I flagged at my lane SITREP. My §10 asserted an equality that **could only be satisfied by changing a contract my fold did not own** — a silent m-8 sentinel strip. The fixture was not merely wrong; **a RED-first assertion that cannot be satisfied legally SELECTS an illegal repair**, and an implementer chasing the red test is the one who would make it.
**Candidate rule: a RED-first acceptance assertion must be checkable against the frozen contracts of every owner whose bytes it touches; if satisfying it would require a change to a contract the fold does not own, the fixture is proposing that change and must be routed, not written.** Cheap to state, mechanical to check at review time, and it composes with the two earlier candidates (branch-attribution fixtures; frozen-carrier resolution at design time) — all three are the same underlying rule: **a claim a gate can check should never rest on drafting discipline.**

## Boundaries
Report-only. This return claims **no** DESIGN-lock, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action; it approves **no** consumer artifact, completes **no** joint record, and grants **no** integrated re-lock. It signals nothing on m-10's or m-3's behalf — both act on their own reviews.

## Verification
- Approved delta `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35` — re-verified byte-identical on disk after the approval.
- Frozen + UNMOVED: worker r7 `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45` · lifecycle r21 `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd` · m-8 provider contract r12 `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51`.
- Governing rulings consumed unchanged: VP classification `4c254307…` (A3/B1 delegated under F73, no operator gate); your governing record `3a440c11…` incl. the six A3 fold conditions, the m-8 tripwire, and the B1 authority language.

ACTIONS_GIT_REF: docs-workspace disk action only — this RECONCILE relay + one `master/relays/INDEX.md` row; **no design byte moved** (the approved r12 verified byte-identical after filing), no `frank/` action, no lock, no fold, no downstream signal.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: m-10 substantively re-reviews §6 and rebases both artifacts to `04422965…`; m-3 binds the same hash and closes its `logical_surface_digest` leg; master proceeds with the §D-settlement amendment ((1)+(3)+(4) + the two r21-claim adjudications) → VP → operator, after which m-9 folds its final batched revision (§7-`relay.*` + §2.6) and the §D join co-signs.
