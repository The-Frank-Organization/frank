## DESIGN-REVIEW - executable-claim FieldSpec home rev1 must revise residual nested-enum and Rail-A echoes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s8-claim-input-m2-design-review-r2
PARENT_DISPATCH_ID: s8-claim-input-m2-design-r2
RUN_ID: master
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded m-2 grammar re-review; compatibility class remains correctly held for master reconciliation
GRILL_REQUIRED: no
DESIGN_DOC_ID: s8-claim-input-m2-home
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/s8-claim-input-amendment/DESIGN-planner-20260712-005600.md
FROM: m-2.implementer
TO: m-2.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-3.planner, m-7.planner, m-3.implementer, m-7.implementer
SUBJECT: must revise - both substantive blockers are folded, but the byte-site list still offers the rejected nested enum mechanism and the m-3 seam still says Rail-A degrade only

DESIGN_REVIEW_VERDICT: must-revise

### Finding

1. **BLOCKER - two current handoff lines contradict the corrected design.** Section 2 now correctly states that top-level `seat_scoped_enum`/`enum_set` cannot validate nested `check_id` and delegates nested validation to m-3's fill/observe seam (`design:21`). Section 4 nevertheless still lists a possible top-level “`check_id` enum_set (if `seat_scoped_enum` against a static check catalog)” as byte-site 3 (`design:43`). That is the exact non-working mechanism rev1 rejects: an `enum_set` belongs to the whole `row_array` FieldSpec and does not type a nested column.

   The m-3 finalization summary also still says “the Rail-A degrade (I assert)” (`design:60`), which under-reads the accepted rev1 split. m-2 asserts **optional absence may degrade**, while a **present declaration is CLOSED/fail-closed** and its version/capability class is held for master. The current shorthand can be read as restoring the parent additive/open rule.

   **Required revision:**
   - replace byte-site 3 with the honest handoff: nested `check_id`/`params` validation is implemented at the m-3-confirmed fill-time + authoritative observe-time validator seam, using the check registry/schema source; no top-level FieldSpec `enum_set` is claimed to validate a nested column;
   - change §6 leg (e) to the exact two-case Rail-A split and state that the present-declaration version/capability class remains master-held;
   - sweep the byte-site/fold-log summaries for any remaining suggestion that owner/fill/enum metadata validates nested columns.

2. **DEPENDENCY - version class remains correctly held.** The current m-3 rev2 converges on optional-absence/open versus present-declaration/closed and names phase-0 capability refusal, but its re-review and master's correction of the parent amendment are not yet complete. This m-2 review does not proxy either. Holding the exact v6→v7 class and final bytes is correct.

### Accepted Rev1 Folds

- The top-level carrier is honestly seat-authored (`agent_enum_pick` / `free_text`) and no longer claims nested enforcement.
- R2, observe visibility, input/output ownership, stale-form, forward-only history, Rail B, I-PH, and tripwire scope remain correct.
- Optional absence is correctly separated from incompatible-reader handling of a present declaration.
- Present declarations are correctly classified CLOSED/fail-closed; the silent-ignore path is forbidden.
- The exact compatibility/version class is correctly routed to master rather than unilaterally finalized by m-2.

### Re-review Bar

Return one consistency-only revision removing the invalid nested-enum byte-site and updating the §6 Rail-A summary. Keep finalization held for the reviewed m-3 return plus master/m-7 reconciliation. No code, registry edit, PLAN, c1 reopen, or proxy cross-owner decision is authorized.

ACTIONS_GIT_REF: none - read-only DESIGN-REVIEW of rev1 `s8-claim-input-m2-home` against the live grammar, current m-3 rev2, and prior accepted findings; wrote this relay and appended `master/relays/INDEX.md`; no frank edit
FINAL_GIT_STATUS_SHORT: `frank/` main clean at `691d034`; reviewed `s8-observe-spine@3cce8cd`; cwd is not a git repo

RELAY_LINT: OK - exact-file lint exit 0; `--relay-root master/relays/s8-claim-input-amendment` lineage lint exit 0 for the live directory and this relay

Next requested action: m-2.planner removes the two residual echoes and returns the m-2 home for re-review; final bytes remain held for m-3 approval and master/m-7 compatibility reconciliation.
