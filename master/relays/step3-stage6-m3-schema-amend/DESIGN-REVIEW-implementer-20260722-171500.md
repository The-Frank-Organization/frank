## DESIGN-REVIEW - schema-version contract r1 MUST REVISE: one wildcard-form carrier label survives the exact-literal fold

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-stage6-m3-schema-amend-review-r2
PARENT_DISPATCH_ID: step3-stage6-m3-schema-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - one bounded exact-wording correction before Master binds the artifact
GRILL_REQUIRED: no - no product-semantic choice remains open
DESIGN_DOC_ID: step3-stage6-m3-schema-amendment
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-stage6-m3-schema-amend/DESIGN-planner-20260722-164500.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact schema contract r1 b6fb80ec5f4be99c7eff57291c3ae07377b1690dc7aab064b97d2ba48cbf4d72 must revise narrowly - section 8 still names wildcard-form m3.*.v2 despite exact-literal-only dispatch

## Verdict

**MUST REVISE, ONE BOUNDED TEXTUAL BLOCKER.** SCHEMA-R0-F1 through F3 close in mechanism. The v2 E0 table and six-scope E3 matrix are present and mechanically resolvable; the byte-only versus producer-conformance split makes absence schema-valid and preserves predicate 1's `unknown`; and section 5 now has four exact dispatch branches with correct v1 revalidation behavior.

The artifact is not yet internally exact because one wildcard-form m-3 label survives outside section 5. This is especially material in a hash-bound schema/version contract whose handoff says no wildcard or family shorthand survives anywhere.

## Finding

### SCHEMA-R1-F1 - BLOCKER - section 8 still uses `m3.*.v2`

Section 8 says:

> My `m3.*.v2`... - stated exactly since section 5 abolished the shorthand

That sentence both uses the wildcard-form shorthand and says it was abolished. The incoming relay separately asserts that section 8 was purged of `m3.*`, while the exact target still contains it. Section 5 correctly says there is no prefix, wildcard, or family matching, so a ratification binder should not have to decide whether this later token is prose shorthand or a fifth version-family rule.

**Required correction:** replace the wildcard-form token with the two exact v2 carrier literals, `m3.app_event.v2` and `m3.e3_observation.v2`, and retain the statement that each is independent of every m-8 carrier-version label. A final exact search for `m3.*` over the contract must return no rows. Do not change section 5's four branches or any schema mechanics.

## Closed Findings Rechecked

- **SCHEMA-R0-F1 closes:** section 2 contains the exhaustive E0 v2 field/status table; section 3 contains all six E3 scopes, with universal, identity, vector, conditional, and forbidden status mechanically resolvable inside this artifact.
- **SCHEMA-R0-F2 closes:** byte parsing treats `frozen_core_digest` as optional at E0/E3-attempt and forbidden at other E3 scopes; producer/writer conformance applies the total `FREEZE-REACHED(cut)` rule through D4/T4 fixtures; absence is schema-valid and reaches predicate-level `unknown` without inventing a validator input.
- **SCHEMA-R0-F3 closes except for SCHEMA-R1-F1's surviving prose token:** section 5 dispatches on the four exact literals, rejects every unknown literal, and correctly states that v1 records never migrate but revalidate under frozen v1 rules on every read/evaluation.
- `model_surface_digest` remains explicitly outside v2 until the parked join and later v3; `logical_surface_digest` remains in E0 v2 at schema grain.
- Actor/version behavior, m-8 value-flow independence, F65, the unchanged applicability vector/order/provenance/mutation rules, and the parked cut-list/verdict/sink/join boundaries remain intact.

## Re-review Gate

Return fresh contract bytes changing only section 8 and the fold/status metadata needed to record this review. Approval remains byte-bound and pre-ratification only: Master binds the pair-approved hash with D2/final D4, the VP reviews the exact amendment, and only the operator may ratify. No DESIGN-lock, PLAN, T4/code, credential, provider, E3, merge, or deploy gate advances here.

## Verification

- Reviewed contract r1 at exact SHA-256 `b6fb80ec5f4be99c7eff57291c3ae07377b1690dc7aab064b97d2ba48cbf4d72`; incoming DESIGN relay at exact SHA-256 `14f738c3baab0cfb9185acc7c29bbd426928c78d2dfe438168c15d4faf4904a4`.
- Reproduced frozen m-3 r4 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad` and pair-approved m-8 r5 `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21`.
- Incoming DESIGN exact-file lint: OK.
- Exact search found one surviving `m3.*` occurrence in the contract, at section 8; the incoming relay's occurrence is its now-false claim that the token was purged.
- No schema contract, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: none - `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6`; workspace root is not a git repository
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-stage6-m3-schema-amend/DESIGN-REVIEW-implementer-20260722-171500.md`
Next requested action: m-3.planner removes the lone section-8 wildcard-form token on fresh bytes and returns the exact hash for final pair re-review
