## DESIGN-REVIEW - r0 MUST REVISE: v2 is rejected by the inherited parser, B presence is not total, and two predicate/staging edges remain unresolved

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-review-r1
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the pair can repair three findings; Master must dispose the contradictory staging/version instructions
GRILL_REQUIRED: no - unchanged from the released rev2 dispatch
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m3/DESIGN-planner-20260722-011500.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact r0 dc3b6eb359909fe351fb20f5aa774ba3e87ae16c2861e8e6520adb1b177a7f7e must revise - closed-schema/version incompatibility, non-total frozen-core presence, incomplete predicate machines, and unresolved release-vs-rev2 authority

## Verdict

**MUST REVISE.** The additive-document discipline, copy-never-compute rule, E0 evidence ceiling, independent observer comparison, observed-window ceiling, and parked sink/join boundaries are sound. The exact r0 bytes are not approvable because the proposed v2 record is rejected by the evaluator the document says remains unchanged, the B presence claims collapse materially different `rejected_local` cuts, and the typed-predicate set is not yet fully decidable or connected to the ratified exit table. The release-versus-rev2 staging conflict also requires an explicit Master disposition rather than a pair-local schema/binding split.

## Findings

### R0-F1 - BLOCKER - `m3.e3_observation.v2` cannot pass the inherited v1-only closed parser

The delta defines `m3.e3_observation.v2` and adds an attempt-only field, while sections 0.1 and 0.3 say the r4 evaluator carries verbatim and is byte-unchanged. Frozen r4 section 3.4 step 2 explicitly parses only `m3.e3_observation.v1`, rejects unknown fields, and applies the v1 required/forbidden matrix. A v2 literal therefore returns `non_applicable(malformed)` before any predicate can consume `frozen_core_digest`.

The release and amendment also name the v1 literals, while r0 silently chooses v2. A version bump is technically coherent with a frozen closed schema, but it is a cross-contract decision, not a local wording choice.

**Required correction:** obtain Master's explicit v1-versus-v2 disposition. If v2 is retained, define the exact v2 full field census and scope matrix and revise only evaluator well-formedness/version dispatch to accept v2; preserve the run-constant acquisition/comparison vector unchanged. Replace every claim that the whole evaluator is byte-unchanged with the narrower true claim that the applicability vector and comparison algorithm are unchanged. If v1 is required, Master must explain how adding fields is compatible with frozen r4's closed-parser rule before the pair folds it.

### R0-F2 - BLOCKER - `frozen_core_digest` presence is not total over the real m-8 outcome grammar

Section 1.1 says presence is defined by whatever the observed m-8 terminal carried, then states categorically that `rejected_local` is pre-freeze and lacks a digest. Frozen m-8 r12 disproves that category: its `rejected_local(internal_integrity_fault)` includes both duplicate-header refusal at freeze and a post-authorize digest-mismatch refusal, while other local rejects are pre-freeze. `phase=failed` therefore cannot by itself tell whether a frozen core existed. The table also has a post-authorize `cancelled(pre_transport)` cut that reached freeze but is absent from r0's consequence list.

The current sibling m-8 producer draft exposes the carrier mismatch: it adds B to normalized `completed`/`failed`/`cancelled` terminals, but policy deny and local reject are DATA-P reply shapes, not normalized terminals, while the same draft claims denied/rejected frozen attempts carry B. R0's `denied HAS it` statement has no m-9-visible carrier in those bytes.

Section 1.2 then makes B REQUIRED on every attempt-scoped E3 record, while predicate 1 assigns `unknown` when the carried digest is absent. Under a closed schema, absence is malformed before the predicate runs, so that `unknown` branch is unreachable.

**Required correction:** provide one exhaustive source-outcome/pipeline-cut matrix over all six m-8 dispositions and the local-reject sub-cuts, naming for each cut: freeze reached or not, exact DATA-P/CTRL-C source field, E0 required/forbidden status, E3 required/forbidden status, and predicate-1 result when no wire exists. Do not define schema validity tautologically as "present iff the producer happened to carry it." Align the m-8 carrier so every required E0 value is observable by m-9, and make the E3 requiredness and predicate-1 missing-input branch mutually consistent.

### R0-F3 - BLOCKER - the typed predicate machines are not total, and two have no ratified exit-leg consumer

`provider_deny_caused_zero_transport` maps denied+zero, denied+positive, and counter-unobservable, but does not assign one result for a non-denied disposition, an unavailable/malformed deny row, or disagreement between the row and terminal deny token. Merely listing a token in `observed_facts[]` does not make token equality part of the verdict.

`relay_record_committed_with_stamped_sender` relies on "definitively absent" without defining a closed governed-read result that mechanically distinguishes `not_found` from transport/read-surface unavailability. That leaves the fail/unknown boundary implementation-dependent.

Separately, the ratified section 7 table names only predicates 1, 3, and 4 in Governance-binding. Searches of the governing amendment show no exit-leg reference for predicates 2 (`provider_deny_caused_zero_transport`) or 5 (`no_alternate_credentialed_provider_route_observed`), despite rev2 saying all five feed Governance-binding, Injection-visibility, or Governed-handoff. R0 repeats the general exit-gate claim but does not discharge this reachability gap.

**Required correction:** give every predicate an ordered or mutually-exclusive total verdict table over all schema-valid input states, including missing and contradictory facts. For relay reads, pin a structured resolution domain such as `committed | not_found | unavailable` and the evidence that distinguishes its members. Route the predicate-2/predicate-5 reachability mismatch to Master+VP: either bind each to an exact section 7 leg/composite condition or label it non-gating evidence and correct the rev2 "feed" claim. The pair must not invent a new exit leg.

### R0-F4 - GOVERNANCE BLOCKER - the `logical_surface_digest` staging split is not authorized by either contradictory instruction

Rev2 explicitly parks "the E0 `logical_surface_digest` carriage" until m-9 and m-8 producer recipes are pair-approved. The later release says it does not override rev2's parking, but its author-now list includes the E0 carriage schema. R0 resolves this itself by authoring schema grain and parking recipe-binding grain. That is a plausible decomposition, but neither instrument states it, and the release expressly says its parking clause still controls.

**Required correction:** ask Master for one explicit disposition: either (a) schema grain is author-now and recipe/binding confirmation remains parked, or (b) section 2 stays wholly parked. Fold the answer verbatim and keep the losing branch absent. Until then, section 2 cannot pair-approve.

## Preserved Work

The following r0 decisions should remain unless a required correction directly touches them:

- governed additive document over frozen r4, with no frozen-byte edit;
- v1 behavior inherited except where an explicitly authorized new schema version changes field census;
- E0 remains `self_reported`, never gate-satisfying or promotable;
- carried foreign digests are copied, never recomputed; the m-9-owned logical digest is the named producer/reporter case;
- per-attempt evidence fields do not join the run-constant applicability vector;
- predicate 1 compares an independently observed wire derivation with the carried producer value;
- predicate 5 is bounded to a complete observed window and makes no claim over uninstrumented or bash traffic;
- the B sink, E two-digest join, and recipe-binding confirmation remain parked until their named producer bytes pair-approve.

## Re-review Gate

Return one exact r1 that closes R0-F1 through R0-F3 and cites Master's explicit dispositions for the schema version, predicate reachability, and section 2 staging conflicts. Preserve the frozen r4 hash and the parked sink/join set. No SITREP claiming pair approval, consumer F73 closure, stage-6 lock, PLAN, T4/code, or later gate may advance on r0.

## Verification

- Reviewed r0 at exact SHA-256 `dc3b6eb359909fe351fb20f5aa774ba3e87ae16c2861e8e6520adb1b177a7f7e`; incoming relay at exact SHA-256 `d9c617b4884e3aa076bc5afcb5e6fc2c7aa0fe4d9a31a344b6140eedc73e7526`.
- Reproduced released rev2 `4e7116deeda18ae42561fb1d38f150f7b43009dd36ddbb56d6dbd5c7fab17cde`, release relay `af83a28bb106593de534a1601804728246edd5c1cd3378bb21bcb43b336e7274`, amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`, and frozen m-3 r4 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`.
- Cross-checked frozen m-8 r12 and the current unapproved m-8 B/E draft `322b4b8554afda4e87d1dc832a3bafe0b18b9c81b9ebffc68389dd0007b5b17c` only as boundary evidence; this review grants neither sibling approval nor cross-seat authority.
- Incoming, rev2, and release relays exact-file lint: OK.
- No domain design or `frank/` product bytes were modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design byte, frank source/test, branch, commit, push, PR, merge, credential, provider, or lock action
FINAL_GIT_STATUS_SHORT: none - `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6`; workspace root is not a git repository
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-relock-dag-m3/DESIGN-REVIEW-implementer-20260722-064758.md`
Next requested action: m-3.planner routes the three Master/VP dispositions, folds one exact r1 closing R0-F1 through R0-F4, and returns it for fresh byte-bound adversarial review; all consumer closure and later gates remain held
