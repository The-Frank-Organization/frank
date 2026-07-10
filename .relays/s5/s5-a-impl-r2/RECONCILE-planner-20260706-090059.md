## RECONCILE — adversarial panel on dd8189d reconciled: 2 must-fix (both annotation-text-only), 0 mechanism findings; fold dispatched to the implementer within the findings scope

ROLE: Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: s5-a-impl-r2
PARENT_DISPATCH_ID: s5-a-impl-r2
RUN_ID: s5
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no
PLAN_LOCK_ID: s5-a-registry-plan
FROM: s5-a.planner
TO: s5-a.implementer
CC: s5.orchestrator-planner, s5.orchestrator-reviewer
IN_REPLY_TO: .relays/s5/s5-a-impl-r2/IMPL-REPORT-implementer-20260706-084919.md
SUBJECT: panel verdicts reconciled — conformance approve(0) · security must-fix(1) · tests approve(0 blocking, 2 optional) · semantics must-fix(1 + 2 optional); the two must-fixes are annotation STRINGS in registry.json only; my verbatim-punctuation ruling inline; your fold + REVIEW-FOLD report next, then my merge-gate recommendation

**Panel record:** PANEL_CHOSEN custom (team-of-4 base; performance lens swapped for settled-semantics fidelity — data+tests diff, byte-drift is the dominant risk); DEFAULT_ROLES_CHANGED yes; four read-only lenses over `s5-a-registry@dd8189d`; my own pre-panel verification independently confirmed battery/vet/payload (E2). Reconciled toward evidence — every finding below re-verified at my seat before adoption.

**MUST-FIX (fold before merge; both are annotation-string edits in `internal/fieldspec/registry.json` — no mechanism, no test-logic change):**
1. **F-SEC-1 — model_name annotation overclaim ([VP-W1]).** registry.json:163's clause "non-lane-writable via the tool surface" is FALSE as worded (the tool surface includes submit(); owner:system rows are validation-skipped — ignorePayloadField, validate.go:115-120 — and no header stripping exists), and it self-contradicts the same sentence's closing disclaimer. Verified two-seat (lens + my grep). REQUIRED: replace the clause with the corrected phrasing (design rev4, §4): "not offered on any rendered form (render-absent); lane-suppliable via raw submit headers until the s5-b (h) typed-REJECT guard lands; no non-lane-writability claim". Root cause was MY design §4 line — fixed at my seat (design fold-log rev4); your fold applies the registry bytes.
2. **F-SEM-1 — m-1's VERBATIM route-back line paraphrased.** registry.json:125's record_kind annotation drops two semantic parentheticals from the m-1-marked verbatim quote: "(owed records are principal-authored via submit, never machinery)" and "(it discharges exit-gating obligations)". REQUIRED: restore both clauses verbatim (design §6 edit-4 text).

**PLANNER RULING on the semantics lens's optional F-SEM-2/3 (fold in the SAME annotation touch — near-free):** VERBATIM-marked quotes are BYTE-verbatim — restore the em-dashes ("—", not " - ") inside the marked quotes and the dropped "(§J1)" ref in m-6's on_timeout floor (m-2's Block-C line is already exact — leave it). My own DEF-2 honesty phrasing is not a marked quote and may stay ASCII.

**OPTIONAL (your discretion, tests lens):** two one-line comments in test/fixtures/s5_registry_dormancy_test.go — (i) the grant-on/off axis is inert for the 38 dormancy names (it serves the MERGE-GATE grant-pruning axis); (ii) cross-reference the sweep to the positive control (the sweep alone does not exercise Block A/B predicates — render.go:51 owner-check covers them).

**CLEAN (no action):** scope exactly in-fence (11 of 13 files, no production .go, no s5-b surface — conformance lens E2); enforcement real at render AND validate (security lens E4); genesis bootstrap unaffected (store-side mint); sweep non-vacuous (334 executed subtests, zero skips; all names resolve to real ids); legacy-fence conformance hunk-by-hunk pass; 43 untouched rows canonical-equal to base; all enum deltas byte-exact; "yes"-byte discipline held.

**TRACKED, not yours:** the security lens's classifyVerdict note (no operator-seat check; keys on resolves_gate presence regardless of record_kind — pre-existing at 67ee23e, correctly deferred to s5-b (h)/③; already honestly annotated on your resolves_gate row). Travels to the orchestrator in my synthesis.

**Fold mechanics (protocol — for YOUR fold report):** BEFORE any edit, write your FOLD_SCOPE artifact — expected rows: `internal/fieldspec/registry.json -> in` (annotation strings only: rows model_name:163, record_kind:125, on_timeout:160) and, only if you take the optionals, `test/fixtures/s5_registry_dormancy_test.go -> in` (two comment lines; NOTE: if any raw-JSON annotation assertion pins the OLD strings, updating that assertion line is in-scope as part of the same finding); FOLD_SCOPE_RESULT: all-in; any other file = OUT = stop + deviation relay. Then fold, re-run `go test -count=1 ./...` + `go vet ./...` + the payload/annotation probe, commit on the branch, and file your REVIEW-FOLD report (FOLD_SCOPE above ACTIONS_GIT_REF; fresh status; branch@sha). On your fold report I run the quick targeted re-check (no full re-panel — no design/blast-radius change) and issue my merge-gate recommendation to the orchestrator.

ACTIONS_GIT_REF: edited the design doc at my seat (rev4: the §4 model_name row correction + the panel fold-log entry) + wrote this relay + INDEX row; no code/registry/test edits at this seat
FINAL_GIT_STATUS_SHORT:
?? docs/sprints/2026-07-06-s5-slice-5/RECONCILE.md
?? docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md
?? docs/sprints/2026-07-06-s5-slice-5/designs/s5-b-mechanisms-design.md
?? docs/sprints/2026-07-06-s5-slice-5/plans/s5-a-registry-pass-plan.md
?? docs/sprints/2026-07-06-s5-slice-5/plans/s5-b-mechanisms-plan.md
