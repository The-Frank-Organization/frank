## DESIGN-REVIEW - s8 config host r5 technically approved; operator and reconcile gates remain

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-design-m7-config-review-r5
PARENT_DISPATCH_ID: s8-design-m7-config
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes - this is technical approval only; the operator must select the F5 owner-attestation fork before lock effectiveness
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m7-config
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s8-config-grill-r1
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/s8-design-m7-config/DESIGN-planner-20260711-035207.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner
SUBJECT: approve technical r5 - all config-host findings resolved; lock remains conditional on operator F5 selection and master reconciliation with the approved m-2 grammar

DESIGN_REVIEW_VERDICT: approve

R5 is technically sound for the m-7 config-host surface. The final schema-consistency finding is closed and no new config-host blocker remains.

## Review

- `canonical_path_families.rows` is now the sole normative representation. The removed `root_homes`/`config_paths` names survive only in the non-normative status line describing their removal.
- The twelve closed rows pin IDs, relative paths, forbidden tokens, and directory flags. `catalog-config` is the explicit post-design addition and has a named planted-token fixture leg.
- Rows are truthfully set-like and sorted lexicographically by `id`; the displayed sequence obeys that rule. `output_families` and `sink_patterns` remain separately identified as order-semantic sequences.
- Provenance, runtime-home, ordering, and FX-CFG-5 text now consume `canonical_path_families.rows` consistently.
- The effective lock identifier is `s8-design-m7-config-r5`, matching this reviewed revision and `s8-config-grill-r1`.
- The prior full-descriptor, drift-oracle, full-SHA, genesis composition, activation, A-1, step-4.5, claim-ceiling, and load-boundary findings remain resolved.
- Activation authorization is already ratified by `step2-plan/SITREP-orchestrator-planner-20260711-025643.md` and is not inferred from this approval.

## Remaining Gates

- This approval does not select F5 fork (a) or (b). The reconciled lock becomes effective only after the operator record selects the lean shape-only default or elects the separately-grilled attestation design.
- Master must reconcile this config-host r5 with the approved `s8-design-m2-grammar` artifact before PLAN consumption.
- No PLAN, IMPL, or effective lock authority follows from this relay alone.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, no effective lock, and no operator F5 choice inferred.

ACTIONS_GIT_REF: wrote this approving r5 DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `691d034`
Next requested action: m-7.planner returns config-host r5 and the already-approved executor-host r4 to master for their separate m-2/m-3 reconciliations and operator ratification items; no further pair revision is requested unless reconciliation changes an m-7 contract.
