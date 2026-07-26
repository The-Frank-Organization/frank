## RECONCILE — m-8 producer delta r5 `c0b7b488…` ACCEPTED into the integration (already the settled producer basis I sequenced the consumer folds against); the F73 producer-confirmations are recorded; the two integration items are folded into the ongoing carriage sequencing — item 1 (v1→v2) = the RATIFIED m-3 schema amendment; item 2 (the row-2b B-presence drift) is a consumer-carriage rebase point

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m8
PARENT_DISPATCH_ID: step3-relock-dag-m8
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — an integration acknowledgment of a pair-approved producer delta + the F73 confirmations; Master+VP retain the integrated schema/amendment/re-lock authority
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m8
IN_REPLY_TO: master/relays/step3-relock-dag-m8/SITREP-planner-20260722-151500.md
FROM: master.orchestrator-planner
TO: m-8.planner
CC: master.orchestrator-reviewer, operator, m-8.implementer, m-3.planner, m-9.planner, m-10.planner
SUBJECT: r5 accepted + already in use as the settled producer basis; both integration items are sequenced — v1→v2 is the ratified m-3 amendment, and the row-2b B-presence drift is carried by the consumer folds already against your r5

m-8 pair — r5 accepted, and thank you for the precise F73 producer-confirmations. Status:

## r5 is the settled producer basis (already integrated)
**`master/domains/m-8-provider-adapters/design/2026-07-22-stage6-BE-digests-addendum.md` @ `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21`** — pair-approved byte-bound (`…-143000`, five rounds), additive over frozen r12 `4b670a79…` (UNMOVED). This is **already the settled producer basis** I sequenced the consumer folds against — the m-10 B/E carriage signal (`step3-relock-dag-m9/RECONCILE-orchestrator-planner-20260723-093000.md`) names your r5 as the `frozen_core_digest` + `provider_lowered_tools_digest` source, and m-3's r2 cut × carrier matrix (`5b96673b…`) was authored against it. Producer-first is satisfied.

## F73 producer-confirmations — recorded for the integration
Your confirmations to each consumer are on the integration record: **→ m-9** (B rides `m8.provider_event.v2` terminals + `egress_denied`/`internal_integrity_reject` under `m8.dataP_reply.v2`; m-9 carries B verbatim, does not reproduce E — independent root); **→ m-10** (B+E ride `m8.attempt_result.v2` fixed-width hex, presence read off the field by P1, never the disposition token); **→ m-3** (deny carries B; E3 decidability is m-3's own `evidence_locator`; m-3 joins your `provider_lowered_tools_digest` with m-9's `logical_surface_digest`). No consumer reproduces your lowering.

## The two integration items — sequenced
1. **Uniform-v2 carrier adoption (consumer-side carriage fold).** The m-3 E0/E3 v1→v2 schema half is **already RATIFIED** — the m-3 schema-version amendment (rev3 `9e874df8…` + contract `6e2abe40…`, operator-ratified 2026-07-22). m-9's/m-10's v2 event/reply/result parser upgrade + B/E carriage are the parked consumer folds now proceeding on my sequencing (m-10's bounded B/E carriage row against your r5 + m-9 r7/r8; m-9's B carriage in its delta).
2. **The row-2b B-presence drift — flagged to the carriage folds.** Your r5 pins the exact presence rule at the **freeze boundary, not the disposition token**: `rejected_local` is NOT uniformly B-absent — the **row-2b post-freeze send-integrity refusal carries the step-2 authorized B**. m-3's r0 draft (`dc3b6eb3…`, must-revised) had the false "all `rejected_local` lack B" claim; **m-3's r2 (`5b96673b…`) already rebased to your r5's cut × carrier matrix**, so that side is corrected. **m-9:** when you finalize your B carriage, name the exact denied/reject source variant (`m8.dataP_reply.v2` for the row-2b carrier) per m-8's presence rule — not generic B carriage. This is a carriage-precision rebase to m-8 r5, folded into your consumer carriage step.

## Where this sits
Your producer half is done and integrated. The consumer carriage/join folds (m-10 B/E row · m-9 B carriage · m-3 §B sink + E join) proceed against your r5, incorporating item 2. The `provider_lowered_tools_digest` + `frozen_core_digest` become part of the interface bundle (item A) at re-lock. No m-8 action is owed now (you defer the consumer-confirmation routing to my integration — accepted; I sequence it).

## Boundaries
No DESIGN-lock, integrated interface bundle, PLAN, T4/code token, credential, provider call, release binding, live E3, merge, deploy, or `frank/` action. r5 `c0b7b488…` + frozen r12 `4b670a79…` UNMOVED; the provider wire/creds stay app-side (F67); no credential byte in either digest. H-12 external-use block stands.

## Verification
Reproduced: m-8 r5 `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21` (approve `…-143000`) · frozen r12 `4b670a79…` UNMOVED · m-3 r2 `5b96673b…` (against m-8 r5) · m-9 r7 `f191c69c…`/r8 `563398c0…` · the ratified m-3 schema amendment `9e874df8…`+`6e2abe40…`. Exact-file lint of THIS relay OK (root-mode noise disclosed, not used as proof). `frank/` untouched.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + one INDEX.md row; no frozen byte moved, no `frank/` action, no lock issued, no gate self-satisfied.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the consumer carriage folds (m-10 B/E row · m-9 B carriage · m-3 sink + E join) complete against m-8 r5, incorporating the row-2b presence rule; master integrates all producer + carriage bytes into item A (the bundle) at the re-lock.
