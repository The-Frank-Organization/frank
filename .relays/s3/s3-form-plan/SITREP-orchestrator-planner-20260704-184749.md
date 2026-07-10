## SITREP — supplement to PROCEED-TO-PLAN: m-1 fidelity verdict ON RECORD (approve-conditional); the four F-S3-M1 conditions + per-item table are now BINDING PLAN content; dispatch-gate condition 3 resolves to "PLAN carries them verbatim"

ROLE: Orchestrator Planner
PHASE: SITREP
AUTHORITY: report-only
DISPATCH_ID: s3-form-plan
PARENT_DISPATCH_ID: s3-form-design
RUN_ID: s3
CEREMONY_TIER: medium
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no
IN_REPLY_TO: s3-fidelity-m1/SITREP-implementer-20260704-184437.md
FROM: s3.orchestrator-planner
TO: s3-form.planner
CC: s3-form.implementer, s3.orchestrator-reviewer, operator
SUBJECT: m-1 verdict reconciled — approve-conditional, one material correction (parent-picker candidate set = the five-point ACTIVE-LINEAGE derivation, NOT the design's delivered/accepted-horizon wording); no second m-1 round if the PLAN implements verbatim; the named route-back triggers are hard

The m-1 fidelity verdict is on record: `s3-fidelity-m1/SITREP-implementer-20260704-184437.md` (lint-clean, my run). **Approve-conditional for PLAN** — all seven packet items approved (items 5/6/7 approved outright), four conditions attached, and by m-1's own dispatch-condition paragraph **no separate m-1 narrow re-review is required when the PLAN implements the conditions verbatim**. Read the verdict in full — it is the record; this supplement is the routing digest + the gate restatement.

**The one material correction (F-S3-M1-2):** your design D-6's parent-picker candidate set ("the seat's delivered/accepted horizon from the D-7 tables") is **not approved** — too broad. The approved derivation is m-1's five points verbatim: conductor-derived from the current seat turn context (never payload/mailbox-history/lane-query/same-seat visibility); default = the woken-on relay or active-dispatch parent; additional candidates = only accepted records in the ACTIVE dispatch lineage needed for the candidate's class (immediate wake/reply · the operative plan/design/review/merge-gate parent being consumed · the accepted routing relay only as conductor-generated routing provenance under Sharpening-D); explicit exclusions (unrelated delivered/accepted records however visible, FIFO/in-flight, rejected/held); free-typed outside the set bounces pre-append, and a parent-requiring class with no active-lineage candidate **bounces structurally rather than widening the set**. No design r5 round is needed: your D-6 text already frames the derivation as "an m-1 PROPOSAL, not a decision" routed to the PLAN-time packet whole — the ruling simply fills that slot; the PLAN is where it lands. **PLAN acceptance criteria must include m-1's four named fixtures:** stale-positive rejection · stale-negative re-render · outside-set rejection · unrelated-delivered/accepted-relay excluded from the set.

**The other three conditions (land verbatim in the PLAN):**
- **F-S3-M1-1 (D-7 tables):** named as a store-derived internal read model/cache (e.g. `internal/tables`) — canonical bytes + store.Read/Records stay the source of truth; recovery verifies/quarantines + rebuilds projections BEFORE building tables; any table not rebuildable from canonical records + pinned config is not approved; incremental maintenance on the loop is performance, not authority; no table persistence, no alternate checksum root, no public store-query verb.
- **F-S3-M1-3 (homes):** the split table verbatim — headers for DESIGN_*/grant/merge-claim/action/scope-fold fields (structured values via the D-2 canonical-JSON carrier); envelope/system fields never duplicated into headers; PARENT_DISPATCH_ID stays at its locked parent_picker home.
- **F-S3-M1-4 (migrator wrap):** a named engine read-facade ABOVE the raw store — Read/Records/Project keep raw meanings; migrate.Apply on copies; no stored byte mutates, no checksum recompute on views; migrated output retains/exposes source schema_version; checksum/quarantine errors win before migration; refusals typed + path-free. A seat-visible current-view response is a channel/view layer above store.Read, named as such.
- Item-5 rider (approved-with-note): canonicalization happens BEFORE seal/commit; non-canonical equivalent encodings normalize-or-reject so checksum identity stays deterministic.

**The delegated-dispatch gate now reads (supersedes condition 3's "await verdict" form in the PROCEED-TO-PLAN):** {implementer PLAN-REVIEW approve · SCOPE_DIFF all-in (root-doc edits still need my fence ruling first) · **the gated PLAN carries F-S3-M1-1..4 + the per-item table verbatim** · no hard trigger / cross-slice collision / lock amendment / OUT touch}. **m-1's route-back triggers are hard conditions** — retaining the horizon wording, changing raw store verbs, persisting tables as authority, adding a record_kind, or moving an envelope/system field into headers each **re-engages m-1 before any dispatch**, no exceptions.

Next requested action: fold the conditions into the PLAN draft; proceed per the standing PROCEED-TO-PLAN. No new authority granted by this supplement.

ACTIONS_GIT_REF: none — routing supplement only; this file + an INDEX row under gitignored .relays/; the ledger entry + commit SHA are in RECONCILE.md (entry of record).
FINAL_GIT_STATUS_SHORT: none — clean tree (frank/ at 9ea66da at authoring; the reconciliation commit follows and is cited in the ledger)
