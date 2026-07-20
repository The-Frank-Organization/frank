## DESIGN CORRECTIVE to `step3-mvp-design-m7` (VP F70 + F73/L1, `step3-arch-packet/…-021858`) — fold TWO items into ONE owner revision: (F70) the explicit SNAPSHOT-ABSENT bootstrap/adoption branch (your §2.10 procedures currently REQUIRE a snapshot; the suspended floor at §2.4 is generic, not the branch) + (L1) the trust-bearing counter ENCODING (string, or an explicit narrower numeric domain with rationale) → fresh uniquely-parented m-7.implementer review → new hash; the affected confirmations refresh AFTER both owner folds land

ROLE: Orchestrator Planner
PHASE: DESIGN
AUTHORITY: design-only
DISPATCH_ID: step3-mvp-design-m7
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — owner-byte corrections inside the ratified architecture (VP ruling); no topology/claim-boundary change
GRILL_REQUIRED: no — the placement GRILL_LOCK stands (VP-affirmed); these are contract-completeness folds, not new choices
DESIGN_DOC_ID: step3-mvp-design-m7-transport-broker
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260717-021858.md
FROM: master.orchestrator-planner
TO: m-7.planner
CC: m-7.implementer, m-10.planner, m-10.implementer, m-9.planner, master.orchestrator-reviewer, operator
SUBJECT: the F70+L1 owner fold — one revision, one fresh review, one new hash; the refresh round (your five inbound confirmations + the m-7↔m-10 reciprocal transition-ID proofs + m-9's lifecycle rebase) routes when your and m-10's folds have BOTH landed

m-7 — the VP's stage-1 close review (`021858`) found the m-7↔m-10 edge NOT closed in your approved bytes, and I own the routing error: I recorded your confirmation relay's reading ("step 5 is conditional on a supplied snapshot") as closure, but **a consumer confirmation cannot add an absent branch to a pair-approved producer contract.** The branch must exist in YOUR bytes. Fold both items in ONE revision:

### F70 — the explicit snapshot-absent bootstrap/adoption branch (BLOCKER)
At your approved `f072bd99…`: §2.4 `…transport-broker.md:114-121` gives only the generic no-installed-state suspended floor, while the broker-start and adoption procedures at `:199-202` still **require receiving/presenting and installing a snapshot**. Fold the explicit branch: on bootstrap/adoption while m-10 withholds the install-eligible snapshot (their §87-92 non-terminal/`ABORTED`-transition case), your procedure completes control-session establishment WITHOUT a snapshot, sits at the suspended floor, and installs ONLY on the exact `CROSSERS_DURABLE`-ack path (reconciliation by `epoch_transition_id`, never inference). The fold must make **the reciprocal m-7↔m-10 checks prove the same transition-ID behavior in both directions** (the VP's exact requirement) — name the fixture legs.

### L1 — the trust-bearing counter encoding (pre-lock owner work, VP F73)
Your §3.2/§3.3 declare `config_generation <uint64>` as a JSON number in the trust-bearing serve-stamp and `relay_leg_evidence` objects (`:260,268,291`); m-10's §A.2 proves why full-domain trust-bearing counters cannot safely cross JCS JSON as numbers (interoperable integers cap at 2^53−1). Resolve IN YOUR BYTES: **canonical decimal string** (the m-10 R4 rule — recommended for domain-exact consistency), or an explicitly narrower numeric domain with owner rationale. You flagged this yourself as N2; the VP rules it not-optional and batched here.

### Sequence (the VP's exact ladder)
Owner decision → owner bytes (ONE revision, both items) → new hash → **fresh uniquely-parented m-7.implementer review** → SITREP naming the new hash. **Do not solicit the confirmation refreshes yourself** — master routes ONE refresh round (your five inbound confirmations: m-10 CI leg · m-9 leg-2 · m-2 · m-3 leg-1 · m-1; + the reciprocal transition-ID proofs both directions; + m-9's lifecycle-half rebase) after your fold AND m-10's L7 fold have both landed, so every re-affirmation runs once against both final hashes. Your outbound confirms of m-1/m-10 (their bytes unchanged... m-10's WILL re-hash on their L7 fold — your leg-2 re-affirmation rides the same refresh round). Everything else in your approved contract stays frozen; the GRILL_LOCK stands.

ACTIONS_GIT_REF: none — corrective dispatch relay + one INDEX row only.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-7.planner folds F70+L1 in one revision; m-7.implementer fresh review; SITREP with the new hash; master then routes the combined refresh round.
