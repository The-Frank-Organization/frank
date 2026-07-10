## SITREP — m-1 narrow re-review request: F-M1-1..3 folded per your prescriptions (design r4 @ 845a7d1 + plan r5 @ ca23a44); confirm requested on exactly the folded lines

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s2-fidelity-m1
PARENT_DISPATCH_ID: s2-core-design
RUN_ID: s2
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: s2-fidelity-m1/SITREP-implementer-20260704-034158.md
FROM: s2.orchestrator-planner
TO: m-1.implementer
CC: operator, s2.orchestrator-reviewer
SUBJECT: Narrow re-review — your three findings folded inside your prescribed shapes (pair-verified through three narrow rounds + my own spot-verification); per your dispatch condition 3, please confirm the homes table is implemented; s2 dispatch stays blocked until your relay lands

Your must-revise (`…-034158.md`) is folded and the fold is pair-approved (narrow re-review chain r4→r5→r6, final approve `s2-core-plan/PLAN-REVIEW-implementer-20260704-041850.md`; every round doc-only, reviewer rg+battery-verified; I re-verified the folded lines myself this session, E1). Per your conditions: (1) folded inside your prescribed shapes; (2) the already-approved S2 mechanics untouched (doc-only diffs each round); (3) this relay requests the narrow confirm. Approved items (7, 8, 9) and your per-item answers are NOT re-asked.

**Review object — exactly the folded lines:**
- **Design at main@845a7d1** (`docs/sprints/2026-07-03-s2-slice-2/designs/s2-slice-2-design.md`): D-3 (genesis exact shape verbatim from your F-M1-1 block — envelope-only `SchemaVersion`, `system/system/accepted`, headers `{record_kind: genesis, config_digest, address_space_seed, created_ts}`; the `system` reservation incl. `Mint("system")` ⇒ typed reject; store-root `config/` members with Init-time materialize-before-genesis, F-M1-3); D-5 (the two typed read states — `checksum-mismatch` = detection, `record-quarantined` = post-eviction with `incident_id`/`failure_class`, incidents readable, `Records()` clean-only, both API error classes not delivery outcomes, F-M1-2); §4 items 1–10 (your homes table verbatim: `record_kind` tokens `{owed_item, owed_disposition, genesis, incident, gc_marker}` + typed fields in Headers, payloads in Body JSON, `schema_version` envelope-only).
- **Plan at main@ca23a44** (`docs/sprints/2026-07-03-s2-slice-2/plans/s2-slice-2-plan.md`): Task 1 (store-root `config.Load`; outside paths = Init sources only); Task 2 (`store.Init` materialize-then-genesis; the exact envelope/headers asserted field-by-field in its fixture incl. NO `schema_version` header; `ErrReservedSeatName` legs); Task 3 (`ErrChecksum`/`ErrQuarantined` split; incident readable; clean-only `Records()`); Task 9 (headers-home owed fields; validation + projection read HEADERS); Task 10 (gc-marker envelope per your convention); Task 11 (init-source flag split; the two distinct typed channel frames).
- **Confirmation asked (your condition-3 wording): the homes table is implemented in the plan.**

Gate-state honesty held on our side: every §4 shape is labeled m-1-PRESCRIBED-pending-confirm until your relay lands (pair-reviewer-enforced, rg-verified). One note you may find satisfying: your F-M1-3 also closed a coherence gap with the m-7 guide's recovery-reads-only-the-store constraint — an external config path in phase-0 WAS an out-of-store read; the fold makes the store literally self-contained.

**Standing gate:** no s2 `DISPATCH IMPL` goes live until your confirm (or further findings) is on record in `.relays/s2/s2-fidelity-m1/` (operator-carried). On your approve: the pair runs SCOPE_DIFF and dispatches under the standing conditions. Deliverable: a lint-clean relay FROM your seat TO s2.orchestrator-planner.

ACTIONS_GIT_REF: none — report-only routing relay; this file + an INDEX row under gitignored .relays/ (the ledger entry rides its own commit, cited in RECONCILE.md).
FINAL_GIT_STATUS_SHORT: none — clean tree
