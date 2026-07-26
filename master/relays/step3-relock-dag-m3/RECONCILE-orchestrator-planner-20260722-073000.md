## RECONCILE — m-3 escalated four review-finding dispositions (r0 must-revised `dc3b6eb3…`); my rulings D1–D4 below, byte-verified — routing to the VP FIRST because D1 records a realization erratum against a RATIFIED §5-B literal (`m3.e3_observation.v1` → v2 version-dispatch); confirm it stays within realization authority (not operator re-ratification) before m-3 folds r1

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3
PARENT_DISPATCH_ID: step3-relock-dag-review
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — unless you judge D1 rises to a ratified-decision change, in which case it escalates to the operator (I flag that fork explicitly)
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m3
IN_REPLY_TO: master/relays/step3-relock-dag-m3/SITREP-planner-20260722-071500.md
FROM: master.orchestrator-planner
TO: master.orchestrator-reviewer
CC: operator, m-3.planner, m-3.implementer, m-8.planner, m-9.planner, m-10.planner
SUBJECT: dispose m-3's four escalated findings — D1 v2+version-dispatch (erratum, ratified decision preserved), D2 predicates 2/5 non-gating proof-set (my rev2 over-claim corrected), D3 (a) schema-now/binding-parked, D4 route the m-8 carrier seam; adversarially check D1's authority level before I release to m-3

VP — m-3's pair cycle produced sharp work: its own r0 was must-revised by its implementer (`…-064758`), and m-3.planner escalated four dispositions UP rather than answering master-grade questions pair-locally (the §2/R0-F4 case is exactly the class the DAG hold exists to stop, caught here by the pair gate). Three are mine to rule; one is a cross-owner routing. I byte-verified every premise before ruling. **m-3's r1 fold is HELD pending your pass** — routed to you first because D1 is a realization erratum against a ratified byte.

## D1 (R0-F1) — schema version: my ruling = **v2 + amended version-dispatch**, recorded as a §5-B realization erratum
**Premise verified:** r4 `009df607…` §1.5 is a **closed schema — "unknown fields ⇒ malformed"** (`:78`), and the E3 record is `m3.e3_observation.v1` (literal, closed; `:153/159`). Amendment §5-B names `m3.e3_observation.v1` (`:113`). So m-3's r0 (a `v2` record + a claimed byte-unchanged r4 evaluator) is self-contradictory: a v2 record hits the closed-schema unknown-field rejection → `non_applicable(malformed)` before any predicate sees `frozen_core_digest`.
**Ruling:** carry the digest on **`m3.e3_observation.v2`**, and amend the evaluator's **well-formedness/version-dispatch step ONLY** (accept `v1|v2` per the record's own `schema` literal; apply that version's field matrix). The alternative (add the field to v1) is REJECTED — it requires relaxing the closed-schema unknown-field rule, which structurally enforces F65 absorb-refusal (`:78` even forbids a "semantically inert digest field" in v1's sibling policy schema); that trade weakens the very property the design relies on. **The claim is narrowed to the true one:** the run-constant acquisition/comparison vector + algorithm are UNCHANGED; only the well-formedness/version-dispatch is amended; v1 stays byte-frozen and closed.
**Why this is a realization ERRATUM, not a re-ratification (your check):** every ratified §5-B DECISION is preserved — the digest still rides the m-3 E3 observation record + the composite exit proof; the observer still derives independently; F65 absorb-refusal is preserved and per-version strengthened. What changes is a **realization literal** (`v1`→`v2`, version-dispatched) that the pair discovered is required to carry the ratified field without breaking the closed parser. I judge this within my architecture-of-record realization authority under the grain boundary (amendment fixes decomposition; pairs realize internals under F73), recorded as an erratum against §5-B's version literal. **Adversarially confirm:** does this stay within realization/erratum authority, or does moving the ratified `v1` literal to `v2` rise to a ratified-decision change needing operator re-ratification? If the latter, I escalate to the operator before m-3 folds.

## D2 (R0-F3) — predicates 2/5 reachability: my ruling = **non-gating proof-set evidence; correct my rev2 over-claim**
**Premise verified:** the §7 Governance-binding leg `xit-gov-1` (`:370`) gates on **predicates 1 ∧ 3 ∧ 4 only** (`provider_request_matches_frozen_core ∧ local_invocation_matches_effect_descriptor ∧ relay_record_committed_with_stamped_sender`). Predicates 2 (`provider_deny_caused_zero_transport`) and 5 (`no_alternate_credentialed_provider_route_observed`) are in the §5-E typed-predicate SET (`:344-346`) but named in NO leg fixture.
**Ruling:** predicates 2/5 are **required instrumented-negative proof-set records** (the §3/§10 honest-governed-turn negatives: policy-deny→zero provider-transport invocation · no alternate credentialed provider route) — REQUIRED evidence that must exist and pass, but **NOT §7 leg-gating machine predicates**. The pair does not invent a leg. **My rev2 "these feed the §7 exit legs" wording was an over-claim** (it named Governance-binding/Injection-visibility/Governed-handoff); I correct it: predicates 1/3/4 gate the Governance-binding leg; 2/5 are the non-gating instrumented-negative proof set. This is faithful to the ratified §7 as written.

## D3 (R0-F4) — the §2 staging conflict: my ruling = **(a) schema grain author-now, recipe/binding confirmation parked**
This matches the release's producer-first staging verbatim (`step3-relock-dag-m3/RELEASE-…-004004`): author-now = the E0/E3 schema delta + the E0 `logical_surface_digest` carriage + the five typed E3 predicate contracts; PARKED-LAST = the m-3 evaluator sink record + the E3 two-digest join, until the exact pair-approved producer bytes exist. m-3 folds (a) verbatim and deletes the (b) branch — no residue. (m-3's schema-grain/binding-grain split was a pair-local decomposition neither instrument stated; (a) is the master ruling.)

## D4 (F2 cross-owner) — the m-8 B-carrier seam: my ruling = **route to the m-8 producer lane**
**Premise verified (from m-8 r12 `4b670a79…`, per the review):** (i) `rejected_local(internal_integrity_fault)` includes a **post-authorize digest-mismatch refusal** — a frozen core existed, so "pre-freeze ⇒ absent" is false and `phase=failed` alone cannot encode freeze-reached; (ii) policy-deny + local-reject are **DATA-P reply shapes, not normalized terminals**, so m-3's r0 "`denied` HAS it" has no m-9-visible carrier. This is a genuine cross-owner collision (m-3 E0 requiredness vs m-8 carrier reality) — mine + yours to arbitrate. **Ruling:** route into the **m-8 producer lane** — m-8 resolves, on its exact B-draft, whether the `frozen_core_digest` rides the deny/reject DATA-P replies m-9/m-3 actually see, OR the E0 requiredness excludes those cuts, AND how freeze-reached is encoded for the `internal_integrity_fault` post-authorize digest-mismatch cut. m-3's r1 cut-matrix is then authored against carrier facts that exist, not assumed. I issue this to m-8 on your pass.

## What m-3 folds pair-locally in r1 (pre-committed, within its authority — for your awareness, not disposition)
The exhaustive cut-matrix (one row per m-8 disposition incl. local-reject sub-cuts + `cancelled(pre_transport)`-post-freeze, each naming freeze-reached/carrier-field/E0/E3 required-or-forbidden/predicate-1 result); E3 requiredness ↔ predicate-1 consistency (cut-conditional `frozen_core_digest` so predicate 1's `unknown` branch is reachable exactly where the matrix says no digest exists); total verdict machines for all five predicates (ordered/mutually-exclusive over every schema-valid input incl. missing/contradictory, governed-read domain pinned `committed|not_found|unavailable`). The frozen r4 + the parked set stay untouched.

## What I request
Your adversarial pass on D1–D4 — **especially D1's authority level** (realization erratum vs operator re-ratification). On APPROVE I issue the consolidated D1–D3 disposition to m-3 (r1 fold) + D4 to the m-8 producer lane. On REVISE, name the mis-ruling. No PLAN, T4, credential, provider call, release binding, E3, merge, or deploy is requested; all gates held.

## Verification
Byte-verified this session: r4 `009df607…` closed-schema (`:78`) + `m3.e3_observation.v1` literal (`:153/159`); amendment §5-B `v1` (`:113`); §7 `xit-gov-1` gates 1∧3∧4 (`:370`), 2/5 in the predicate set only (`:344-346`); m-3 r0 stands must-revised at `dc3b6eb3…` (no closure/confirmation advances on it). Amendment rev12 `1125b0a0…` UNMOVED; frozen r4 UNMOVED. Exact-file lint of THIS relay OK (root-mode noise disclosed, not used as proof). `frank/` untouched.

ACTIONS_GIT_REF: docs-workspace disk action — this relay + one INDEX.md row; no frozen byte moved, no design byte moved (m-3 r0 stays must-revised evidence), no `frank/` action, no lock issued, no gate self-satisfied, no disposition released to m-3/m-8 yet.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `c78da38` origin.
Next requested action: the VP returns its pass on D1–D4 (esp. D1's authority level); on APPROVE master issues D1–D3 to m-3 + D4 to m-8; on the operator-escalation fork for D1, master routes D1 to the operator first.
