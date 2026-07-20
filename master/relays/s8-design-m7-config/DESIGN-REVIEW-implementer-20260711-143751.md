## DESIGN-REVIEW - s8 config host r12 must consume reconcile check item six

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-design-m7-config-review-r12
PARENT_DISPATCH_ID: s8-design-m7-config
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - no operator fork; the catalog mutation-path call is m-7-owned and master supplied a grounded recommendation
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m7-config
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s8-config-grill-r1
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m7-config/DESIGN-planner-20260711-143335.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner
SUBJECT: r12 closes F7 and master's first five items, but it missed the r2 owner/master trail and check item six requiring a state-aware post-adoption catalog mutation path

DESIGN_REVIEW_VERDICT: must-revise

r12 correctly consumes the original five-item owner ruling: the MAJOR-safe classification, reserved-token rule, two adoption registry sites, exact count policy, wrapper decision, and old-reader fixture are all present and consistent. The live trail advanced immediately before the r12 relay, however, and added one unresolved mechanism edge.

## Finding

### F8 - The catalog is pinned but has no post-adoption mutation path

m-2's re-confirmation `s8-adoption-grammar-m2/SITREP-planner-20260711-142347.md` replaces `…-135623` as the reconcile input. It confirms the corrected `{catalog, engine}` body and flags that the shipped singular classifier accepts only `fieldspec|engine`. Master consumed that r2 record in directly addressed `SITREP-orchestrator-planner-20260711-143317.md` and added reconcile check item six: m-7 must decide and state how catalog mutations occur after adoption.

The r12 doc still says catalog changes happen through operator §7 `config_change` and its drift law assumes such changes, but defines no `member: catalog` singular arm. `classifyConfigChange` and `configTarget` therefore continue to reject the operation the design requires. r12 also cites the superseded owner record as the satisfied gate and does not mention the r2 replacement/ruling.

Required fold: take master's recommended option (a), which is the only option consistent with the existing §7 and drift-law text: add a singular `member: catalog` config-change arm. Add `catalog` beside `adoption` at both m-2 registry byte sites (`config_member` enum set and `member.seat_scope.operator`) in the same s8 changeset and classify both token additions as one MAJOR-but-safe record-schema moment.

Make the arm state-aware: `member: catalog` is lawful only when catalog is already in the pinned member set. Before adoption, it typed-rejects and cannot substitute for the atomic absent-to-initial transition. After adoption, it uses the ordinary singular body, schema/version-transition validation, recomputes the three-member header digest, emits one catalog `IntentConfig`, and remains restart-effective. Preserve the owner-fidelity/change-convention review gate and source-artifact drift rule already stated in §2/§4.

Add executable legs for pre-adoption singular catalog reject and post-adoption singular catalog accept -> one intent -> chain digest advance -> restart/load success. Sweep the GRILL_LOCK and lock-impact trail to consume `…-142347` plus `…-143317` as the current owner/master records; retain older records as superseded history.

## Confirmed

- r11/F7 is fully closed by the MAJOR-safe classification, adoption byte sites, reserved-token/count/wrapper decisions, and FX-CFG-14.
- The adoption bootstrap, canonical record, recovery, transition and capability gates remain technically approved.
- The catalog singular arm is not a new governance axis; it completes the already-stated §7 catalog lifecycle.
- All descriptor, census, step-4.5, activation, and operator F5 folds remain accepted.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, and no effective reconciled lock.

ACTIONS_GIT_REF: wrote this r12 DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `691d034`
Next requested action: m-7.planner folds check item six into config-host r13, consumes the r2 owner/master records, and returns a DESIGN relay for final re-review; master holds reconcile-A completion meanwhile.
