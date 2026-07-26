## DESIGN-REVIEW - schema-version contract r0 MUST REVISE: the closed matrix is still delegated, conditional presence is not byte-decidable, and wildcard dispatch is not exact

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-stage6-m3-schema-amend-review-r1
PARENT_DISPATCH_ID: step3-stage6-m3-schema-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded owner corrections before Master binds the artifact into amendment bytes
GRILL_REQUIRED: no - no product-semantic choice remains open in this review
DESIGN_DOC_ID: step3-stage6-m3-schema-amendment
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-stage6-m3-schema-amend/DESIGN-planner-20260722-134500.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact schema contract r0 a09c9931047efe8fa9a52564164fd353e7c12767a09b9763f3cef0c9dc98c534 must revise - full matrix absent, FREEZE-REACHED validation context unnamed, and wildcard/history dispatch wording is not mechanical

## Verdict

**MUST REVISE.** The central direction is correct: both m-3 carriers move to independently closed v2 schemas; v1 stays frozen; `logical_surface_digest` is in E0 v2 at schema grain; `model_surface_digest` waits for a later v3 with the settled join; m-8/m-3 version labels are independent; and only evaluator well-formedness/version dispatch changes. Those decisions close the original r0-F1 architecture defect.

The exact artifact still fails the VP's ratification-bindability bar in three bounded places. It does not actually contain the complete E3 per-scope matrix, its conditional-required rule cannot be evaluated from either record or a named validation input, and its supposedly mechanical dispatch uses a wildcard plus an incorrect historical-record rule.

## Findings

### SCHEMA-R0-F1 - BLOCKER - the complete v2 E3 matrix is still delegated to frozen r4

Section 2.2 says the r4 six-scope matrix is "reproduced unchanged" except for the attempt addition, but no six-scope matrix appears in this artifact. Section 1 provides a field census and cites r4 for the required/forbidden cells. This is a precise delta, but it is still delegation by reference - the exact defect M3-VP-R2-F1 and the `130000` route said a ratifiable closed-set artifact must eliminate.

The ratification binder should not need to interpret "same as another document except" to discover whether `policy_digest`, `provider_lane_id`, `artifact_ref`, `relay_id`, or each build/release XOR is required or forbidden at a scope.

**Required correction:** reproduce the full six-row `m3.e3_observation.v2` matrix in this artifact, with universal, identity, vector, and additive fields explicitly required, forbidden, or conditional in every scope. Keep the v1 census as a hash-bound reference if desired, but the v2 closed matrix itself must be present and complete. For E0, likewise state one closed field table or equivalent exhaustive census whose required/conditional/forbidden status is visible without importing prose rules.

### SCHEMA-R0-F2 - BLOCKER - `FREEZE-REACHED(cut)` is total as a function but not mechanically available to either closed validator

Section 3 defines a total boolean over a producer cut, and section 5 makes a missing conditionally-required field `malformed`. Neither `m3.app_event.v2` nor `m3.e3_observation.v2` carries `cut`, the contract names no attempt-bound external source from which an E0 acceptor or E3 evaluator obtains it, and the unchanged r4 evaluator has no m-10-store or m-8-record acquisition step. A function can be mathematically total while still being unevaluable at the validation boundary.

This matters because the same record bytes can contain `phase=failed` both before and after freeze. Presence cannot be validated from `phase`, and label-independent m-8 versioning does not supply the missing relation.

**Required correction:** separate the two layers explicitly:

1. **Byte-only closed parsing:** state whether `frozen_core_digest` is syntactically optional/conditional at E0 and E3 attempt scope, with correct type when present and forbidden at non-attempt E3 scopes.
2. **Producer/writer conformance:** require presence iff the settled D4 cut matrix says freeze reached, naming the exact attempt-bound input and actor that performs this check.

If absence is to be `malformed`, name a constructible validation context that supplies the exact cut and binds it to `{run_id, turn_id, attempt_id}` for each acceptor/evaluator. Otherwise, absence must remain schema-valid and yield predicate-level `unknown`, while D4 fixtures enforce producer conformance. Do not imply the byte parser can infer a hidden pipeline position.

The now pair-approved m-8 r5 `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21` supplies the exact P1 cut matrix for Master's D4 fold; it does not by itself create an m-3 validator input.

### SCHEMA-R0-F3 - BLOCKER - exact dispatch uses wildcard literals and historical v1 behavior is misstated

Section 5 dispatches on `m3.*.v1` and `m3.*.v2`. In a ratifiable closed-schema contract, `*` is not a literal: read mechanically, `m3.unrelated.v1` could select the v1 branch despite section 5.3 saying every other value is malformed. The two carriers also have different matrices, so one wildcard branch does not identify which matrix runs.

Section 4 says v1 history "never re-validates." Frozen r4 requires a presented E3 record to be rechecked for canonical bytes, reference digest, closed-schema validity, observer provenance, and applicability whenever it is evaluated. Historical records do not migrate or reinterpret as v2, but they absolutely revalidate under v1.

**Required correction:** replace the wildcard rules with four exact literal branches: `m3.app_event.v1`, `m3.app_event.v2`, `m3.e3_observation.v1`, and `m3.e3_observation.v2`, each naming its exact matrix and validator/acceptor. State that historical v1 bytes never migrate or reparse as v2; when read/evaluated, they validate again under the frozen v1 rules. Unknown literals and fields remain malformed with no prefix/family matching.

## Preserved Decisions

Preserve these r0 decisions unchanged unless a required correction directly refines their mechanics:

- both carriers use new v2 literals and keep v1 byte-frozen;
- `logical_surface_digest` is in E0 v2 at schema grain, with recipe-binding confirmation still parked;
- `model_surface_digest` is excluded from v2 and lands with the settled E join under a later v3;
- m-9 emits E0 v2 only after the lane-2 build; m-10 and readers accept exact v1/v2 literals by their own matrices;
- the E3 writer emits v2 after ratification while the evaluator accepts historical v1 and current v2;
- no dual emission, coercion, cross-version leniency, or shared-label inference with m-8;
- the run-constant applicability vector, acquire-then-compare ordering, provenance checks, F65 split, and mutation semantics stay unchanged;
- the cut list, verdict machines, sink, and E join remain outside this bounded artifact.

## Re-review Gate

Return one exact r1 containing the complete matrices and closing SCHEMA-R0-F1 through F3. Pair approval remains only a pre-ratification input: Master must bind the approved hash with D2 and final D4, the VP must review the exact amendment, and only the operator may ratify. The held lane-2 r1, integrated re-lock, PLAN, T4/code, credential, provider, release-binding, E3, merge, deploy, and H-12 external-use gates do not advance on these bytes.

## Verification

- Reviewed contract r0 at exact SHA-256 `a09c9931047efe8fa9a52564164fd353e7c12767a09b9763f3cef0c9dc98c534`; incoming DESIGN relay at exact SHA-256 `2a8cf91b52680e0d6767b4201efef5dbab03cdc2bd7f991e2585db80f19cb230`.
- Reproduced route relay `aa8ed0f18cb1b58ac0b252eb9db8bf6b0e6e21cc71f880d2aabc250c97f90c88`, frozen m-3 r4 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`, and pair-approved m-8 r5 `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21`.
- Incoming DESIGN exact-file lint: OK.
- No schema contract, frozen design, amendment, or `frank/` product bytes were modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: none - `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6`; workspace root is not a git repository
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-stage6-m3-schema-amend/DESIGN-REVIEW-implementer-20260722-160000.md`
Next requested action: m-3.planner folds one exact schema-contract r1 closing SCHEMA-R0-F1 through F3 and returns it for fresh byte-bound review; only an approved hash goes to Master for amendment binding
