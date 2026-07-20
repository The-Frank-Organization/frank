## DESIGN-REVIEW - s8 config host r6 must revise the version-gate enforcement contract

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-design-m7-config-review-r6
PARENT_DISPATCH_ID: s8-design-m7-config
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - the operator selected F5 fork (a) in `SITREP-orchestrator-planner-20260711-130825.md`; only the technical r6 findings remain
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m7-config
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s8-config-grill-r1
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m7-config/DESIGN-planner-20260711-041219.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner
SUBJECT: r6 must revise - member-set premise correction is valid, but version ceilings are bypassable without acceptance-time bump enforcement and undefined across heterogeneous version markers; F5 human gate is now satisfied

DESIGN_REVIEW_VERDICT: must-revise

The r6 premise correction is accepted: governed phase 0 calls `config.Load(StoreRootConfigPaths(root))`; a pre-s8 binary supplies the hardcoded two-member set, so it cannot compute the s8 three-member expected digest. A1's engine-member version carrier is also the correct smaller home. Two mechanism gaps remain in A2.

## Findings

### F1 - The ceiling gate depends on a version bump that the runtime does not enforce

Section 1.1 defines exact engine bump semantics and section 2.4 rejects versions above a reader ceiling (`2026-07-11-s8-config-host.md:10-21`). But FX-CFG-10 says a same-version schema change is only "refusable at review-time" and then delegates a mechanical leg to an unspecified catalog shape rule (`design:130`). If an accepted `config_change` adds engine/catalog semantics without advancing the marker, the old reader sees a version within its ceiling, ignore-unknowns the new fields, and recreates the exact same-set fail-open class A2 claims to eliminate.

Required fold: bind version-transition validation to the trusted acceptance path. Before accepting `config_change`, compare current and candidate member shape/version: schema-surface change requires exactly current+1 for numeric engine versions (or the member's defined successor); value-only change requires the same version; rollback, skip, or same-version schema drift is a typed path-free reject. Recovery may replay only already-accepted transitions. FX-CFG-10 must exercise this acceptance gate, not a review convention.

### F2 - "Per-member ceiling" is not mechanically defined for the three actual marker types

The engine proposal is integer-valued, while catalog and fieldspec currently use string markers such as `s7-v1` and `s7a-fieldspec-v5`. Section 2.4 alternates between a numeric ceiling and a handler-supported set, but does not define parsing/order, required presence, malformed-marker disposition, or each s8 reader's initial supported values. The PLAN would have to invent the compatibility relation.

Required fold: specify the closed capability table at the design grain. For each member, name the marker field, parser/comparison rule or exact supported set, required/missing behavior, initial s8 marker, and supported ceiling/set. Missing, malformed, unknown, rollback, or above-support markers must produce a typed phase-0 config-load fault before partial interpretation. If catalog/fieldspec use exact supported sets rather than ordered ceilings, name that honestly and keep the common contract as "handler-declared supported versions."

## Confirmed

- A1 correctly chooses `engine.version`, initial 1, and distinguishes schema changes from the `observe` value flip.
- Member-set additions/removals are already fail-closed in the governed runtime via hardcoded supplied paths plus expected composite digest; FX-CFG-9 is a valid executable proof of that narrower fact.
- Per-member capability checks are the right mechanism for same-set intra-member evolution once F1/F2 make their inputs trustworthy and comparable.
- The operator has selected F5 fork (a), so the previous human gate is satisfied; no attestation-at-load design is required.
- All r5 descriptor, digest, SHA, canonical-census, step-4.5, and activation folds remain approved.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, and no effective reconciled lock.

ACTIONS_GIT_REF: wrote this r6 DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `691d034`
Next requested action: m-7.planner folds F1-F2 into config-host r7 and returns a DESIGN relay for re-review; master holds reconcile-A completion meanwhile.
