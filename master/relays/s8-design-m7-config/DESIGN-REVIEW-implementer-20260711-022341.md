## DESIGN-REVIEW - s8 config host r1 must revise before lock

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-design-m7-config-review-r1
PARENT_DISPATCH_ID: s8-design-m7-config
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - the operator still owns the F5 owner-attestation election; technical review cannot ratify that fork
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m7-config
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s8-config-grill-r1
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m7-config/DESIGN-planner-20260711-015554.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner
SUBJECT: must revise - the descriptor is still a summary rather than a buildable full governed schema, the drift oracle confuses member and composite digests, the byte-exact SHA is abbreviated, and the operator fork remains unratified

DESIGN_REVIEW_VERDICT: must-revise

The engine-member home, restart-effective `config_change` activation, A-1 generation boundary, step-4.5 placement, and m-2 ownership fence are coherent. Four pre-lock findings remain.

## Findings

### F1 - The governed discovery descriptor is still a summary, contrary to the dispatch's explicit full-descriptor requirement

The dispatch requires every bite-determining element to become governed data: roots/exclusions, boundary files, portable idioms, context recognizers, and the exact site/family census. The design lists those categories at `2026-07-11-s8-config-host.md:23-28`, but it does not define the descriptor's closed field/schema shape or carry/pin the actual 17-site list, six-family/two-carve-out set, sink-pattern symbols/counts, recognizer forms, root, and exclusions. The current catalog schema has only `path_hygiene.families` and `sink_patterns`; roots, exclusions, boundary files, recognizers, portable idioms, and site census still live in test code (`frank/test/invariants/catalog_test.go:13-51` plus `path_hygiene_test.go`).

Required fold: make the design implementation-deterministic by specifying the closed descriptor shape and pinning every current value, either inline or by a byte-exact canonical appendix/source section. State how the law consumes each field and which mismatch class turns red. Do not leave the PLAN to invent names, normalization, or census membership.

### F2 - FX-CFG-5 names a genesis-recorded member digest that the design does not record

Section 4.3 says source bytes, runtime bytes, and the "genesis-recorded member digest" are equal (`design:34`), while section 5 records only the composite digest (`design:38-42`). The live mechanism likewise stores only `config_digest` (`frank/internal/store/genesis.go:64-79`) and later config changes advance the expected composite through `new_digest` (`genesis.go:125-140`). A catalog-member SHA cannot equal the composite digest, and after a catalog `config_change` the genesis composite is intentionally stale.

Required fold: choose and state one real oracle. Either add a governed per-member digest manifest to genesis/config-change records, or compare source bytes directly with runtime bytes and recompute the full current composite against `ExpectedConfigDigest()` (genesis plus accepted config-change chain). Update FX-CFG-5 and FX-CFG-7 to the same model.

### F3 - The supposedly byte-exact fieldspec pin is abbreviated

Section 5 and the GRILL_LOCK carry only `1ef6abab...2485`, while FX-CFG-7 requires a byte-exact assertion. The already-recorded full SHA is `1ef6abab4d496b11017f57ca400e8296d63824994ffce8311e4533f70cc92485` (`s7a-merge-gate/MERGE-GATE-orchestrator-planner-20260710-155633.md`).

Required fold: put the full 64-hex SHA in the normative genesis composition and fixture obligation; an ellipsis is display text, not a lock input.

### F4 - Review approval cannot ratify the still operator-owned F5 fork

The GRILL_LOCK correctly says the (a)/(b) owner-attestation election cannot be inferred and remains operator-owned (`design:95-96`), but its lock impact says approval alone creates `s8-design-m7-config-r1` (`design:98-99`), and the relay says the default ratifies at review/lock. This seat has technical review authority, not operator product-risk authority.

Required fold: separate technical design approval from operator ratification. Keep fork (a) as the recommended default, but make the reconciled DESIGN_LOCK conditional on an operator record selecting (a) or (b). If (b) is selected, it requires the separately-grilled design already named here. Also carry the m-2 approved artifact's separate activation-authorization ratification into master reconciliation rather than silently treating this review as that answer.

## Confirmed Non-Blockers

- `present_layers` in the engine member preserves one-writer-per-fact and exposes only an engine-computed environment fact to m-2.
- `observe:false` at genesis is a defensible dogfood default because it preserves the governed activation event.
- Step 4.5 is correctly framed as an amendment owned here while profile semantics remain m-2-owned.
- The load claim is honestly bounded to digest/history plus member/provenance shape; it does not claim review happened.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, no lock, and no operator choice inferred.

ACTIONS_GIT_REF: wrote this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `691d034`
Next requested action: m-7.planner folds F1-F4 into config-host r2 and returns a new DESIGN relay for re-review; master holds reconciliation/lock consumption meanwhile.
