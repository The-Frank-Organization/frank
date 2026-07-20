## Team m-2 - Forms & Determinism: DESIGN-REVIEW RESPONSE (rev1)

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: c1-design-m-2-review-r1
PARENT_DISPATCH_ID: c1-design-m-2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - approval of m-2 design review; consumer review and joint design lock remain orchestrator-owned
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
DESIGN_DOC_ID: c1-design-m-2-forms-determinism
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: approve

Verdict: approve.

I reviewed the rev1 design relay and the revised design artifact:
- `master/relays/c1-design-m-2/DESIGN-planner-20260628-174201.md`
- `master/domains/m-2-forms-determinism/design/2026-06-28-v3-form-schema-design.md`

The two previous blockers are resolved at E1:
- `ROLE` is now consistently system-owned in the envelope, stamped from the same certified connection as `FROM`, and only projected read-only if rendered in a header view. Evidence: design doc section 2 lines 23 and 27, section 3 lines 33-39, section 10a line 126, section 12 line 172, AC7 line 218.
- The lineage engine is now a pre-delivery gate over `submitted` records, with delivery and consumer action allowed only after transition to `accepted`. Lineage failure bounces to the author and is never delivered. Evidence: section 4 lines 68-74, section 10c lines 150-159, section 11.1 lines 161-164, AC8 line 219.

Review asks:
1. The two-state `submitted` -> `accepted` model and the named lineage-blocking record classes restore the v2.8.8 "blocks before dispatch" guarantee for authority-bearing records. The key invariant is consumer-side: only `accepted` records are deliverable/consumable.
2. Reclassifying `role_from_consistency_error` into section 10a is correct because the mismatch class dissolves by construction once `FROM` and `ROLE` are stamped from one channel. The updated approximate counts (~33 dissolve / ~16 form-validation / ~13 lineage) match that reclassification.
3. `recipient_picker` for `TO`/`CC` preserves the addressing checks as form-validation: the address space is system-minted, sender choice is constrained to that space, and phase/grant-specific cardinality such as dispatch-TO-exactly-one-implementer remains a typed required/cardinality rule. `FROM` remains exactly-one `system_only` identity, not a picker.
4. I found no remaining section 2-vs-section 12 inconsistency. The layer/owner model is now coherent: `FROM`/`ROLE` are system identity; `TO`/`CC` are envelope-layer recipient picks; `PARENT_DISPATCH_ID` is a constrained parent picker.

Non-blocking follow-through for the eventual PLAN:
- Add negative fixtures for `X-` attempted gate input, read-only `ROLE` projection, lineage-failed `submitted` non-delivery, and `recipient_picker` dispatch cardinality.
- Keep the consumer-review prerequisite explicit for m-3/m-4/m-6 before joint lock; this approval is only m-2 implementer design review.

ACTIONS_GIT_REF: created this DESIGN-REVIEW relay at master/relays/c1-design-m-2/DESIGN-REVIEW-Implementer-20260628-175833.md and appended the routing row to master/relays/INDEX.md; docs-workspace artifacts, no code/source implementation changes; cwd is not a git repo so a commit/ref is unavailable
FINAL_GIT_STATUS_SHORT: unavailable - cwd is not a git repository; git status reports fatal not-a-git-repository at /mnt mount boundary
