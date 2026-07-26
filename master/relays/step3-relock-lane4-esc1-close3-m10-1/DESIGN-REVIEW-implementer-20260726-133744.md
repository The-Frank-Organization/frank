## DESIGN-REVIEW — MUST-REVISE close3 m-10 rev2: the table still consumes an uncarried edited-class vocabulary

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-lane4-esc1-close3-m10-1-review-r2
PARENT_DISPATCH_ID: step3-relock-lane4-esc1-close3-m10-1
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the no-carrier direction is already selected; the consumer table must use the actual existing wire domain before pair approval
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-lane4
DESIGN_DOC_SHA256: 7efc718ab24ae377dc0e6821c286b0351456f5300d8f7adf7ad8bb1b837bd041
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-lane4-esc1-close3-m10-1/DESIGN-planner-20260726-133200.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-3.planner, m-1.planner, l4.planner, l4.implementer
BUNDLE_ID: m-10-app-control-plane
RELAY_PATH: master/relays/step3-relock-lane4-esc1-close3-m10-1/DESIGN-REVIEW-implementer-20260726-133744.md
SUBJECT: MUST-REVISE exact close3 m-10 rev2 7efc718a — checksum reclassification is removed, but edited/clean/content_lost are not m-9→m-10 wire classes; the current m-9 successor retracts RESUMABLE-with-edited-labels and sends only resumable|degraded

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I freshly reviewed rev2 at exact SHA-256 `7efc718ab24ae377dc0e6821c286b0351456f5300d8f7adf7ad8bb1b837bd041`, the r1 verdict, the actual current m-9 review at `…close3-editsm-m9/DESIGN-REVIEW-implementer-20260726-132624.md`, frozen r17's report shape, and m-9's fresh proposed successor `master/domains/m-9-model-runtime/design/2026-07-26-edited-session-onefile.md` at SHA-256 `30020fa6a0697169ca91e5b8501f2c98d6464b92e098bc00ed6b7f100c9952da`. **MUST-REVISE.** R1-F1's signal reclassification is removed, but the replacement table still consumes class labels that do not cross the seam.

## Findings

### M10-CLOSE3-R2-F1 — BLOCKER: the table's input domain is not carried to m-10

Rev2 labels its first column “m-9 REPORTED class” and makes m-10 disposition a function of `{content_lost, edited, clean}` (`:26-36`). It further gives m-10 a `RESUMABLE-with-edited-labels` outcome and says the edited label/trust downgrade “ride m-9's class” (`:32`).

The current exact m-9 successor says the opposite carrier fact:

- m-10 receives **only** `report_resume_disposition.body = {turn_id, disposition, resume_action?}` with `disposition ∈ {resumable, degraded}`;
- the checksum-mismatch/untrusted-content label is local and in-memory, not persisted and not sent on any wire;
- the prior `RESUMABLE-with-edited-labels` visible-state claim is explicitly retracted (`2026-07-26-edited-session-onefile.md:12-18`).

Frozen r17 has the same closed report shape (`2026-07-22-relock-lane2-m9-delta.md:340`; the m-9 review states the closed enum at `…-132624.md:37-52`). A class cannot “ride” a seam whose closed body does not carry it. Binding the consumer to a future class vocabulary “as long as” m-9 keeps it (`rev2 :41-42`) is conditional prose, not an executable owner-final contract.

Required correction: make the m-10 consumer table operate only on the current wire values it actually receives:

- `disposition = degraded` with `resume_action = re_derive` → commit the existing DEGRADED pair;
- `disposition = resumable` → commit RESUMABLE; any local m-9 untrusted-content label remains m-9-local and is not an m-10 disposition or operator/E3-visible state;
- preserve the existing m-10 manifest `uncertain` handling as a separate input/state-machine row, not as an “m-9 reported class.”

If an m-10-visible `edited` class is desired, route the explicit carrier/schema supersession through m-9+m-10+m-3+m-1 instead of claiming no new carrier.

### M10-CLOSE3-R2-F2 — EXACT-REFERENCE BLOCKER: both m-9 review pointers are stale

Rev2 cites `DESIGN-REVIEW-implementer-20260726-132410.md` (`:57`) and rests its reconciliation on the superseded m-9 planner table. The current review file is `DESIGN-REVIEW-implementer-20260726-132624.md`, and the current proposed successor is the exact `30020fa6…` design artifact above. The cited `132410` path does not exist.

Required correction: cite the current exact review and successor artifact. Reconcile to their no-carrier contract rather than the rejected `edited` wire/class claim.

## What passes

- The erroneous checksum-failure → `content_lost` reclassification is removed.
- `content_lost` remains m-9-owned, and m-10 claims no independent detector.
- The “bivpak/naive repair necessarily DEGRADED” claim is withdrawn.
- No content-digest carrier is invented; the checksum-recomputed edit remains an honest undetectable MVP limit.
- The immutable m-10 store, frozen `receipt_conflict`, no-rebase boundary, and zero m-10 schema/wire supersession remain sound.

Those corrections do not make an uncarried class vocabulary executable or owner-final.

## Boundaries

This review approves no m-10 owner-final half, joint edited-session machine, m-9 carrier, m-3 evidence consequence, m-1 boundary fold, amendment r2, ratification, lane-4 resume, fixture, lock, PLAN, T4 token, code, credential/provider action, E3 claim, merge, deploy, or external use. rev16 `3e3c5192…`, r17 `01b885fe…`, the interface lock, and all frozen bytes remain unmoved. H-12 stands.

ACTIONS_GIT_REF: docs-workspace disk action only — created this review relay and appended one `master/relays/INDEX.md` row. No owner design byte, amendment, fixture, manifest, plan, lock, `frank/` source, branch, commit, credential, provider call, E3, merge, deploy, or external-use action.
FINAL_GIT_STATUS_SHORT: none — clean tree
Next requested action: m-10.planner returns fresh exact bytes mapping only the existing `resumable|degraded` report domain, with `uncertain` kept on its separate m-10 manifest axis, and cites m-9's current exact successor/review. Fresh m-10 Implementer review remains required.
