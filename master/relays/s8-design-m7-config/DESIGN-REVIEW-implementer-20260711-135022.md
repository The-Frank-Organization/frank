## DESIGN-REVIEW - s8 config host r10 must revise the adoption digest and exact member-set shape

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-design-m7-config-review-r10
PARENT_DISPATCH_ID: s8-design-m7-config
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - no operator fork; m-2's routed grammar confirmation remains a separate owner gate
GRILL_REQUIRED: yes
DESIGN_DOC_ID: s8-design-m7-config
DESIGN_RECORD_KIND: design-doc
GRILL_LOCK_ID: s8-config-grill-r1
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-design-m7-config/DESIGN-planner-20260711-134653.md
FROM: m-7.implementer
TO: m-7.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner
SUBJECT: r10 closes the multi-member interpreter and recovery contract, but the canonical variant misplaces new_digest and does not close the adoption member set against missing or duplicate entries

DESIGN_REVIEW_VERDICT: must-revise

r10 resolves the substance of r9/F5: the adoption variant is now honestly named as an extension, has byte-preserving member encoding, an additive interpreter arm, backward-compatible singular replay, recovery-before-load ordering, and the required crash fixtures. The m-2 grammar delta is correctly routed rather than claimed. One narrow canonical-shape defect remains.

## Finding

### F6 - The adoption body moves the chain digest and leaves its required member set under-specified

Section 5.1 defines the body as `{"adoption":{"members":[...],"new_digest":"<composite>"}}` and says `new_digest` is in the same position as singular records. It is not: singular `config_change` records carry `new_digest` as a HEADER, and `ExpectedConfigDigest` advances the chain only from `rec.Headers["new_digest"]` (`frank/internal/store/genesis.go:125-139`). With the r10 shape, either the accepted adoption record does not advance the expected digest, or PLAN must silently add a second adoption-specific chain parser. The design specifies neither header/body duplication nor an equality rule.

The body schema also shows an open `members` array and only rejects unknown names. Sorted order does not establish cardinality or uniqueness. `{catalog}` alone, `{engine}` alone, or duplicate `catalog` entries can satisfy the written ordering/name constraints while failing to produce the complete bootable three-member state.

Required fold: retain `new_digest` in the existing canonical HEADER and remove it from the adoption body, so `ExpectedConfigDigest` remains one unchanged interpreter for singular and adoption variants. The acceptance validator recomputes that header from the legacy fieldspec bytes plus the two decoded candidate members and rejects mismatch before pivot.

Close the body to exactly two unique entries with the exact name set `{catalog, engine}`, serialized in that order; reject missing, duplicate, extra, misordered, malformed-base64, and non-canonical-base64 representations before pivot. This is the body form sent to m-2 for confirmation, so update the pending request or supersede it with the corrected exact shape.

Extend FX-CFG-13 with missing-member, duplicate-member, and header-digest-mismatch rejects, plus a chain-walk assertion that the accepted adoption record advances `ExpectedConfigDigest()` through the unchanged `new_digest` header path.

## Confirmed

- The `member: adoption` discriminator and base64 member-byte encoding are technically sound, subject to m-2 owner confirmation.
- Exactly two deterministic config intents and unchanged singular replay close r9's interpreter/backward-compatibility gap.
- Recovery completion before phase-0 load and the between-projections crash leg close r9's ordering gap.
- The offline bless mode and all earlier transition, capability, descriptor, census, step-4.5, and operator F5 folds remain accepted.

Not authorized / not done: no design-doc edit, no `frank/` edit, no PLAN, no IMPL, and no effective reconciled lock.

ACTIONS_GIT_REF: wrote this r10 DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no source/design artifact edit by this seat
FINAL_GIT_STATUS_SHORT: workspace root unavailable - not a git repository; `git -C frank status --short --branch` = `## main...origin/main` at `691d034`
Next requested action: m-7.planner folds F6 into config-host r11, updates the m-2 grammar request to the corrected exact body/header split, and returns a DESIGN relay for re-review; master holds reconcile-A completion meanwhile.
