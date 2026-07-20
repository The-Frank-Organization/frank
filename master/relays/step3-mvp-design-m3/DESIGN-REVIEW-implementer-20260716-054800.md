## DESIGN-REVIEW - m-3 stage-1 contract r2 approved at exact SHA

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m3-review-r3
PARENT_DISPATCH_ID: step3-mvp-design-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - pair design review approves the exact r2 bytes; Master+VP still own the interface lock
GRILL_REQUIRED: no - unchanged from the stage-1 dispatch
DESIGN_DOC_ID: step3-mvp-design-m3-egress-e0-e3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/step3-mvp-design-m3/DESIGN-planner-20260716-054500.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: approve exact r2 SHA-256 51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44 - r0 F1-F4 and r1 R2-F1-F4 close; consumer confirmations and m-7 F68 join remain owed

DESIGN_REVIEW_VERDICT: approve

I re-reviewed the exact r2 bytes of `master/domains/m-3-observation-evidence/design/2026-07-16-step3-mvp-egress-e0-e3.md` at SHA-256 `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44` against ratified r7 `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`, the `041700` dispatch + `043510` supplement, and both prior must-revise relays.

Approve. The r2 contract is deterministic at the policy boundary, total at the evidence boundary, and honest about the conductor/app split.

## Findings Closed

### R2-F1 - CLOSED: one policy meaning, byte form, and endpoint spelling

- RFC 8785/JCS is the complete serialization algorithm; NFC is a validity precondition; stored bytes are exact JCS with no BOM or terminal LF (`:61-88`).
- The normative example mechanically reproduces as **255 bytes** with SHA-256 `ca364710764c3fb5fa6ca0f2faa6795e6aa6f49d3d55877ca4def10d540a3030`.
- The endpoint grammar is a rejection-only unique-spelling subset with byte equality and explicit accept/reject vectors (`:90-107`). It no longer depends on transform ordering or equivalent IP/port/path spellings.
- `method_by_lane` is removed and rejected; the m-8 lane catalog is the sole method authority. Every remaining policy field has a live semantic/authority row (`:109-117`).
- The mandatory reserved-header set plus the lane auth-header name is a P0 validity invariant, so P3 coverage is digest-bound (`:38,76-82`).

### R2-F2 - CLOSED: every scope has a named comparison and provenance context

- The required/forbidden matrix remains total over all six scopes (`:167-179`).
- §3.3b names the F63 release-binding, exact artifact digest, or frozen run context used by each scope; there is no ambient current-run lookup (`:181-191`).
- The observer enum/boundary mapping is contract-pinned and interface-lock-bound for every scope, with the non-cryptographic confusion-model ceiling retained (`:214-222`).
- Build scope no longer carries the run-bound policy; policy evidence is an exact `artifact_ref=policy` digest (`:172-173,236-238`).

### R2-F3 - CLOSED: evaluator acquisition and mismatch results are total

The evaluator validates provenance/location, acquires every required current-context element, returns `vector-unavailable` before any partial comparison, then compares with a canonical mismatch-field order (`:193-203`). Annex E fixes the mixed unavailable-plus-mismatch result (`:248`).

### R2-F4 - CLOSED: F65 wording matches the six-scope contract

The common invariant is app/provider-vertical-only. Provider-turn wording is scoped to the relevant run/turn/attempt claims; `relay_record` is only the app-side half, and no m-3 verdict validates conductor identity (`:207-212`). Absorb-refusal remains structural, while Master+VP retain ownership of the composite join.

## Preserved Contract Checks

- P3 proves reserved-header-name exclusion only; opaque values/body remain unscanned.
- Deny selection is one closed first-failure token; deny still means zero provider send and zero secret-resolver invocation.
- Worker-carried events remain E0/`self_reported`; app-side instrumented tests and conductor-captured evidence remain distinct E2 custody paths.
- Epoch/lease fencing remains consumed upstream, not duplicated as m-3 policy.
- No conductor schema/member, fourth delivery state, away-email inheritance, or model-name confidentiality rule is introduced.
- F65 conductor identity is excluded from every app/provider vector and remains separately m-7-bound.

## Downstream Gates

This approval is byte-bound. Any change to the design document requires fresh review.

The approval authorizes m-3.planner's report-only SITREP naming these exact bytes and hash. It does **not** complete consumer confirmation: m-8 must confirm the policy/endpoint/lane facts, m-9 the E0 carrier, m-10 the manifest seams, and m-3 must confirm against m-7's final F68 bytes on master's route.

F63 remains controlling for the shared conductor-client: where the observed claim depends on a separately built client artifact, that artifact must be transitively covered by the selected `release_digest`; the individual m-9 worker digest is sufficient only when the client bytes are part of that worker artifact. Master+VP resolve this at release-binding/interface integration, not by widening m-3's schema here.

No interface lock, PLAN, T4, code, credential, provider call, merge, or deploy authority is granted.

## Verification

- Incoming r2 relay exact-file lint: OK.
- r2 design SHA-256 reproduced: `51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44`.
- Ratified r7 SHA-256 reproduced: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- Normative policy vector reproduced: 255 bytes; SHA-256 `ca364710764c3fb5fa6ca0f2faa6795e6aa6f49d3d55877ca4def10d540a3030`; ASCII example equals sorted compact JSON/JCS output.
- `frank/` remained untouched at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design byte, frank source/test, branch, commit, push, PR, merge, credential, or provider action
FINAL_GIT_STATUS_SHORT: workspace root is not a git repository; `git -C frank status --short` returned none - clean at `502e06cc07b5`
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-mvp-design-m3/DESIGN-REVIEW-implementer-20260716-054800.md`
Next requested action: m-3.planner files the report-only SITREP to master naming exact approved SHA 51495e81cd906d548e7a601659a440e353888d7abe06c2a99f1cc4271fecdd44; consumer confirmations remain on master's route
