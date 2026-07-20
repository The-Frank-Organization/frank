## DESIGN-REVIEW - m-3 stage-1 contract r1 remains must-revise

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m3-review-r2
PARENT_DISPATCH_ID: step3-mvp-design-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the residuals are deterministic contract completion within the ratified m-3 scope
GRILL_REQUIRED: no - unchanged from the stage-1 dispatch
DESIGN_DOC_ID: step3-mvp-design-m3-egress-e0-e3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-mvp-design-m3/DESIGN-planner-20260716-052000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-7.planner, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: r1 at SHA-256 d98dd6ad2021bbcb6229f8bb8e51f825fefc89f4e887ceb1260ea1b9be5ab843 substantially folds r0, but policy bytes/semantics and E3 evaluation context remain non-total

DESIGN_REVIEW_VERDICT: must-revise

I re-reviewed the exact r1 bytes of `master/domains/m-3-observation-evidence/design/2026-07-16-step3-mvp-egress-e0-e3.md` at SHA-256 `d98dd6ad2021bbcb6229f8bb8e51f825fefc89f4e887ceb1260ea1b9be5ab843` against r7 `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`, the r1 relay, and my r0 F1-F4 required-return bar.

r1 closes the P3 honesty defect and the F65 seat-authority overclaim. It also supplies structured fields for all six scopes and a concrete external record-digest handle. Approval is still blocked by the residuals below.

## Blocking Findings

### R2-F1 - The policy still does not have one executable semantic form or one byte form

Section 1.5 claims one logical policy has exactly one byte string (`:75-80`), but its home-grown JSON rules do not define string escaping, the lexicographic comparator, or whether a terminal LF exists ("LF" and "no insignificant whitespace" do not settle that). The same NFC string can still be serialized as literal UTF-8 or `\u` escapes while satisfying the written rules. This is avoidable: the adjacent m-10 stage-1 contract already names RFC 8785/JCS for digested JSON (`master/domains/m-10-app-control-plane/design/2026-07-16-mvp-ipc-manifest-seam-contract.md:33`).

Section 1.6 also calls its URL a single byte-level representation while permitting IP literals "verbatim" (`:82-83`). Equivalent IPv6 spellings, non-canonical numeric ports, DNS trailing-dot/IDNA forms, and percent/dot-segment ordering are not resolved. Two implementations can produce different bytes for the same endpoint or disagree on whether input is canonical.

The policy schema also contains a semantic dead field: `method_by_lane` appears only in the example schema (`:64-72`), while P4 compares the request to the lane catalog's declared method (`:43`). The contract never says whether the policy map is authoritative, must equal the lane fact, or is ignored. Similarly, P3 describes the denied-name set as containing the mandatory four names plus the lane auth-header name (`:42`), but P0 never states that semantic membership invariant.

Required fold: use one complete named canonical-JSON algorithm (with any NFC requirement made an explicit validity precondition), and pin exact stored-byte/newline behavior. Either restrict endpoint grammar to a genuinely unique subset or adopt a complete named URL canonicalization with ordered transforms and acceptance/rejection vectors. Remove `method_by_lane` or make its exact key set and relationship to the canonical lane method normative. Make the mandatory denied-header membership invariant a P0 validity rule. No semantically inert digest field may remain.

### R2-F2 - Observer and comparison context is missing for non-run scopes

The observer registry is declared **run-scoped** (`:167-168`), but `build` and `artifact` records forbid `run_id` (`:138-139`). The evaluator therefore cannot select the registry whose boundary mapping is supposed to bind `observer_id`. The registry mapping is also absent from every bound digest, so the prose does not say whether it is interface-locked, run-manifest-bound, or mutable; changing the mapping can change provenance without any applicability-vector change.

The same context gap affects vector equality. Build identity is the release vector and artifact identity is `artifact_ref` plus one digest, yet step 5 compares both against "the currently-running vector" (`:154-156`) without identifying which run/release binding supplies that current vector. A machine cannot determine the comparison target when multiple releases/runs exist.

Required fold: add a scope-by-scope evaluation-context table. For every scope, identify (a) the immutable source/version of the observer registry and recorder-boundary mapping and (b) the canonical source of the target vector. Non-run scopes need an unambiguous release/build context, or their applicability must be defined as immutable exact-digest evidence rather than comparison to an unnamed current run. Bind or interface-lock the observer mapping so historical provenance cannot silently change.

### R2-F3 - The evaluator's mismatch/unavailable branch is not executable in its stated order

Step 5 requires equality comparison before step 6 handles a current-vector value that cannot be obtained (`:150-157`). Equality cannot be completed when an operand is unavailable. For a record with one mismatched field and a second unavailable field, two implementations can return `mismatched` or `vector-unavailable` while both plausibly follow the prose. `mismatched: [field…]` also lacks a canonical field order despite the claim that verdict reasons are closed and ordered.

Required fold: acquire all required current-vector elements before comparison, then define one precedence for any-unavailable versus mismatch. If mismatch carries a field list, order it by one explicit canonical field order. Add the mixed unavailable-plus-mismatch vector to the normative annex.

### R2-F4 - F65 wording still over-specializes all scopes as a provider turn

The record type now covers `build`, `artifact`, and app-side `relay_record` claims, but §3.5 says every evaluator `applicable` verdict is for the "provider-turn leg only" (`:161-165`). That is true for the attempt-scoped live provider record, not for all advertised scopes. The important invariant is that every scope is limited to the app/provider vertical and never validates conductor identity.

Required fold: state the common boundary as app/provider-vertical-only. Reserve "provider-turn leg" for the run/turn/attempt claims that actually represent that leg, and describe `relay_record` as only the app-side half of the relay exchange. Keep the Master+VP composite and m-7 confirmation requirements unchanged.

## Prior-Finding Disposition

- **r0 F1:** partially closed. Structured scope identities and record-digest home pass; R2-F2/F3 block total evaluator execution.
- **r0 F2:** closed. P3 now states exactly the reserved-name property it proves.
- **r0 F3:** partially closed. Deny-token precedence passes; R2-F1/F3 block canonical policy/evaluator determinism.
- **r0 F4:** closed. The schema no longer claims it enforces m-7-half presence.
- **E2 custody and epoch non-duplication:** remain closed.

## Re-review Bar

Return fresh bytes and SHA with R2-F1 through R2-F4 folded. Include exact canonical-JSON bytes, endpoint accept/reject pairs covering IP/host/port/path ambiguity, a policy-field ownership/semantic table, a six-scope registry/vector-source table, and one mixed unavailable-plus-mismatch evaluator vector. No operator decision is needed unless the fold changes the ratified scope enum, topology, or ownership split.

This verdict grants no report-only close SITREP, consumer-confirmation routing, interface lock, PLAN, T4, code, credential, provider call, merge, or deploy authority.

## Verification

- Incoming r1 relay exact-file lint: OK.
- r1 design SHA-256 reproduced: `d98dd6ad2021bbcb6229f8bb8e51f825fefc89f4e887ceb1260ea1b9be5ab843`.
- Ratified r7 SHA-256 reproduced: `2f75f2a1df00c4e330fc052bb06023547b3cfd1fe323c1e9cb0e43b9550e6e1d`.
- `frank/` remained untouched at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design byte, frank source/test, branch, commit, push, PR, merge, credential, or provider action
FINAL_GIT_STATUS_SHORT: workspace root is not a git repository; `git -C frank status --short` returned none - clean at `502e06cc07b5`
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-mvp-design-m3/DESIGN-REVIEW-implementer-20260716-052637.md`
Next requested action: m-3.planner folds R2-F1 through R2-F4 into fresh uniquely hashed design bytes and returns a uniquely parented DESIGN relay for r2 review
