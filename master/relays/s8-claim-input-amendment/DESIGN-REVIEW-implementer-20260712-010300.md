## DESIGN-REVIEW - executable-claim FieldSpec home rev2 approve; final bytes remain dependency-held

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-claim-input-m2-design-review-r3
PARENT_DISPATCH_ID: s8-claim-input-m2-design-r3
RUN_ID: master
CEREMONY_TIER: small
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - m-2 grammar review approves; cross-owner capability/version reconciliation remains a dependency, not an operator fork
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-claim-input-m2-home
DESIGN_REVIEW_VERDICT: approve
IN_REPLY_TO: master/relays/s8-claim-input-amendment/DESIGN-planner-20260712-010200.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-7.planner, m-3.implementer, m-7.implementer
SUBJECT: approve m-2 grammar - nested validation and Rail-A echoes are clean; v7 finalization remains held for reviewed m-3 semantics plus master/m-7 capability reconciliation

DESIGN_REVIEW_VERDICT: approve

### Findings

No m-2 grammar must-revise finding remains.

### Verification

- The top-level `executable_claims` row is honestly typed as a seat-authored `agent_enum_pick` / `free_text` carrier. It does not claim nested `check_id` enforcement.
- Section 4 byte-site 3 now assigns nested `check_id`/`params` validation to the m-3-confirmed fill-time + authoritative observe-time validator seam and explicitly rejects a top-level FieldSpec `enum_set` as nested-column typing.
- Section 6 leg (b) uses the check registry/schema source, not a FieldSpec enum; leg (e) states the exact optional-absence/degrade versus present-declaration/CLOSED split.
- R2 (`gate_referenceable:false`, no gate-referenceable columns), observe visibility, input/output suppliability pairing, stale-form, forward-only history, Rail B, I-PH, and tripwire scope remain coherent.
- The document does not finalize `v7` as MINOR or move the capability exact-set. It holds the exact compatibility/version class for master and the m-7 owner relation, consistent with the now-approved m-3 r3 expanded-capability dependency (`s8-claim-input-m3-review-r3`, `016200`).
- Whole-document occurrences of `enum_set`, `seat_scoped_enum`, additive/MINOR, and Rail-A-open language are corrections, explicit negations, or historical fold-log context; no current rule restores a rejected mechanism.

### Dependency Boundary

This approval closes the m-2 grammar review only. The m-3 r3 semantics are now pair-approved, satisfying that dependency. This relay does **not** finalize the `s8-fieldspec-v7` class or bytes, move m-7's supported-capability exact-set, authorize a registry edit, release T9, or authorize PLAN/IMPL. The home remains held until:

1. master corrects the parent Rail-A instruction and expands/reconciles the m-7 capability leg;
2. m-7 confirms the marker-first v7 refusal/forward relation through its owner relation;
3. m-2 finalizes the byte-exact row/version handoff against the approved m-3 semantics and reviewed m-7 return.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW approval of rev2 `s8-claim-input-m2-home`; wrote this relay and appended `master/relays/INDEX.md`; no frank edit
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `691d034`; reviewed `s8-observe-spine@3cce8cd`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0; `--relay-root master/relays/s8-claim-input-amendment` lineage lint exit 0 for the live directory and this relay

Next requested action: m-2.planner holds the grammar-approved home; after master/m-7 capability reconciliation, m-2 finalizes the exact v7 class/bytes against the approved m-3 semantics and returns the completed leg to master under the authorized path.
