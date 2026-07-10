## SITREP — s3 → m-1 fidelity packet: the S3 form-system design's seven store/lineage proposals (operator-carried; parallel with PLAN drafting)

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s3-fidelity-m1
PARENT_DISPATCH_ID: s3-form-design
RUN_ID: s3
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: s3-form-design/SITREP-planner-20260704-183144.md
FROM: s3.orchestrator-planner
TO: m-1.implementer
CC: operator, m-1.planner
SUBJECT: fidelity review request — review object = design §4 (seven proposals) at frank/ main@291ab08; the S2 packet pattern; delegated dispatch is BLOCKED until your verdict is on record in .relays/s3/; expected-light (S3 sits above the store) but the VP lineage-movement watchpoint routes it to you regardless

**What this is.** The S3 slice (full FieldSpec registry + linter dissolution; guide m-2) has a pair-approved design (`docs/sprints/2026-07-04-s3-slice-3/designs/s3-slice-3-design.md` r4 at main@291ab08, DESIGN_DOC_ID `s3-slice-3-design`). Its §4 enumerates every store/lineage touch as **proposals, not decisions** — your authority per the master s3-dispatch + the [VP-W] watchpoint (lineage movement is an m-1 fidelity trigger even inside m-2-owned modules). The pair drafts its PLAN in parallel; **no `DISPATCH IMPL` issues before your verdict is on record** (the F2 condition, restated in the PROCEED-TO-PLAN).

**Review object — design §4's seven proposals (the design's own numbering; supporting detail in the named D-sections):**
1. **D-7 engine tables over the locked verbs** — the §10c lineage engine + obligation completion ride incrementally-maintained in-memory tables (rebuilt at recovery, maintained on the commit loop); **no store-API change proposed** (Records/Read/Project stand).
2. **D-6 PARENT `parent_picker` realization** — conductor-derived candidate set + free-typed-outside-set reject; the canonical VP-watchpoint item.
3. **D-6 lineage-field homes** — DESIGN_DOC_ID/DESIGN_LOCK_ID/DESIGN_RECORD_KIND/DESIGN_REVIEW_VERDICT/grant/merge-claim refs as **headers**; envelope untouched — per your S2 F-M1-1 homes-table precedent.
4. **D-9 migrator read-facade wrap point** — read-time migration above a migration-agnostic store; where the chain wraps Read/Records/Project consumers.
5. **D-2 canonical-JSON-in-string header carrier** (operator-decided at the grill) — structured values ride inside the existing `Headers map[string]string`; **no envelope/checksum change** — confirm.
6. **record_kind token additions: none proposed** (the S2 five stand; the authority-bearing floor stays the S1 pessimistic superset) — confirm.
7. **Registry member replacement posture** — fresh `store.Init`; a digest change = the phase-0 wall, fixtured **positive** (master-ratified fresh-store posture, `s3-scope-q1` ruling; the §7 config-change record is OUT with OI-S3-CONFIG-CHANGE carrying it).

**Specific asks (answer per-item approve / approve-conditional / must-revise, the S2 form):** (a) does any D-7 table shape amount to a de-facto store-query semantics change you want named? (b) the exact parent_picker candidate-set derivation you'll hold the PLAN to; (c) the homes split for the new lineage fields (headers vs envelope); (d) the read-facade wrap point you sanction for the migrator chain; (e) confirm 5/6/7 as stated.

Context (read-only): the reconciled audits' 8-surface m-1 enumeration (planner audit §5, `.relays/s3/s3-form-audit/AUDIT-planner-20260704-170105.md`); the m-7 consult answers (`../.relays/s3/s3-consult-m7/SITREP-planner-20260704-171546.md` — render-side lineage reads touch committed records only); the master scope ruling (`../.relays/s3/s3-scope-q1/RECONCILE-orchestrator-planner-20260704-171608.md`).

Next requested action: your fidelity verdict via operator hand-relay into `.relays/s3/s3-fidelity-m1/` (the S2 pattern); a must-revise folds bounded and comes back narrow.

ACTIONS_GIT_REF: none — fidelity request relay only; this file + an INDEX row under gitignored .relays/; no tracked-file edit in this action.
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at 291ab08)
