# s5-a — the registry pass: implementation PLAN (locked to s5-a-registry-design rev2)

**PLAN_LOCK_ID:** s5-a-registry-plan
**DESIGN_LOCK_ID:** s5-a-registry-design (rev2) · **DESIGN_RECORD_KIND:** design-doc · approving review `.relays/s5/s5-design-s5-a/DESIGN-REVIEW-implementer-20260706-060559.md`
**Owner:** s5-a.planner (plan) → s5-a.implementer (execution after plan-review approve + literal dispatch)
**Build surface:** worktree `~/frank-s5-team/s5-a`, branch `s5-a-registry` off `main @ 67ee23e` (operator-cut; verify base at step 0). **Every row/enum/predicate byte comes from the design doc** `docs/sprints/2026-07-06-s5-slice-5/designs/s5-a-registry-pass-design.md` (rev2) — this plan sequences it and pins verification; it introduces NO new semantics.

## Locked scope — the only files the IMPL may touch

1. `internal/fieldspec/registry.json` — the single atomic content pass (design §§2–6).
2. `internal/fieldspec/registry_test.go` — byte-exact token/row assertions (design §7).
3. `test/fixtures/s5_registry_dormancy_test.go` — NEW; the [VP-W3] fixture (design §7).
4. `internal/fieldspec/render_test.go`, `internal/fieldspec/validate_test.go` — registry-content cases ONLY if a §7 assertion fits there better than in the fixture; default is everything in files 2–3.

Explicitly NOT touched (deviation = stop + relay, no exceptions): any production `.go` under internal/ or cmd/ · `internal/fieldspec/fieldspec_test.go` (s5-b's classifier grain) · engine/bounce/migrate/test-replay anything · the sprint docs (committing designs/plans rides integration at the orchestrator/operator gate, NOT this branch's IMPL scope) · `.relays/` substrate · the archived dogfood store.

## Steps (TDD order — assertions first where they can fail meaningfully)

**0. Base check.** `git -C <worktree> rev-parse HEAD` = 67ee23e ancestry, tree clean, `go test ./... ` green pre-change (the baseline the zero-regression floor is measured against).

**1. Write the failing assertions (red).**
   a. Extend `registry_test.go`: gate_category_A = the 9 tokens in design §2 order; gate_category = 14 with `routing_escalation` immediately before `other`, `other` last; gate_category_B unchanged 4; the 10 new enums byte-exact (design §2 table); row count 83; version `s5-fieldspec-v3`; per-block spot checks (one row per block: e.g. achieved_evidence enum_set + owner; on_timeout valueless enum + system_only; routing_assignments gate_referenceable true; gate_category_pick enum_set gate_category); delivery_state `{accepted, rejected, held}` re-asserted; record_kind seat_scope: `"*"` = `["diagnostics"]`, operator = the 6-token list, `genesis` in NEITHER.
   b. Write `test/fixtures/s5_registry_dormancy_test.go` (design §7 verbatim): the 6×11×5 sweep over the 38 enumerated names (literal slice — Block A 12 + B 13 + C 9 + D 2 + ACTIONS_GIT_REF + FINAL_GIT_STATUS_SHORT = 36 new + 2 existing), **split per design rev3: 37 names `HasField == false` on every cell; `resolves_gate` false on every NON-operator cell and asserted TRUE on the operator form** (the live verdict-path affordance — the blocker ruling of 20260706); digest determinism per cell; OI-S4 scope legs (genesis absent from record_kind options on every form incl. operator; owed_item/owed_disposition/gate_resolution/disposition options present ONLY on the operator form); legs (a) observe-requireds off · (b) SUBJECT + GRILL_LOCK_ID still block · (c) EVIDENCE_TARGET required INCLUDING on an empty-header candidate · (d) `{"layer_present":"model"}` parse-rejected + a mutated-registry copy naming model_name fails load; the predicate-level positive control (observe flipped on ⇒ Block A/C visible_when fires — the predicate_test.go:64-77 pattern); the D-5 byte control (any_row fires on declared_deviated `"yes"`, not `"true"`); annotation-presence assertions over RAW registry JSON (not loaded structs).
   Run: steps 1a/1b tests FAIL against the unedited registry (proves they bite).

**2. The registry pass (green).** Apply design §§2–6 to `registry.json` in one edit: version + provenance (§1/R-s5-1) · the 2 enum edits + 10 new enums (§2) · Block A 12 rows (§3, uniform double observe predicate, honesty annotations) · Block B 13 rows (§4, incl. on_timeout valueless with m-6's verbatim floor) · Block C 9 rows (§5, idiom (i) + m-2's verbatim annotation on each; routing_assignments gate_referenceable true, no seat_scope; the justified_deviation/deviation_reason_code any_row conjunctions with value `"yes"`) · Block D 2 rows (§6: resolves_gate; gate_category_pick working-name) · the 4 in-place edits (§6: EVIDENCE_TARGET `{"not":{"phase_in":[]}}`; visible_when on ACTIONS_GIT_REF + FINAL_GIT_STATUS_SHORT; record_kind scope FINAL — genesis nowhere, owed/gate_resolution/disposition operator-only, `"*"`=[diagnostics]) · rows appended after GRILL_LOCK_ID in block order (D-8); the 43 untouched rows byte-stable.

**3. Full verification.** `go test ./...` (whole battery, uncached) + `go vet ./...` — all green, zero regression outside the intended new/updated assertions. Then the §7 payload-contract check: `python3` one-liner row count 83 + enum count 24; the raw-JSON annotation greps.

**4. IMPL report** (relay, PHASE: IMPL, ACTIONS_GIT_REF branch@sha + diffstat, FINAL via fresh status) → plan-review chain already satisfied; report to s5-a.planner + orchestrator per protocol.

## Acceptance criteria (= design §10, restated as the gate)

Design §10 items 1–7 verbatim, measured at the worktree tip: loads/vet/battery green (1) · byte-exact tables (2) · the sweep + legs + controls + OI-S4 legs green (3) · 43 rows byte-stable + 4 edits exact (4) · verbatim annotations present incl. m-1's widening-route-back on owed rows, no non-lane-writability claim anywhere (5) · file fence held (6) · the three in-pass m-2 confirms folded IF landed before IMPL completes, else noted for integration (7, per the PROCEED-TO-PLAN).

## Anti-half-fix guards

- The enum delta and the registry_test.go token-list update land in the SAME commit (a split leaves the battery red — the byte-exact floor).
- No row without its designed annotation; no annotation paraphrase where the design says VERBATIM (m-2's Block-C line; m-1's owed line; m-6's on_timeout floor).
- The 38-name fixture slice is generated from the design's row tables, cross-checked against the registry diff — a declared-but-unenumerated row is a [VP-W3] hole (rev0 review's exact worry).
- `"yes"`/`"no"` everywhere a bool byte appears in new data — never `"true"` (the DEF-1 class; D-5).
- If any in-pass m-2 confirm (MR-1 name/type · D-1/D-6 mechanics · disposition ruling) lands mid-IMPL: fold per its content before the IMPL report; if it renames gate_category_pick, the rename is a same-commit mechanical sweep of row + tests.

## Boundary contract (unchanged from dispatch)

Writes: registry.json rows + named_enums delta + the dormancy fixture + registry-content test assertions. Reads: the settled field semantics (the design doc's cited authorities). Target entity: the FieldSpec registry. Downstream consumer: s5-b's ③/⑤/I-PH fixtures + §7 config-change legs (consume the landed registry bytes/digest per design §8, AFTER integration at the orchestrator's gate — R-s5-6); Step-2 observe consumes the declared fields. Proof: E2 (the battery + the fixture). No-consumer action: defer (already applied — scope_paths etc. excluded by design).

## Ceremony / verification target

CEREMONY_TIER medium (no downgrade — no ESCALATION_SCAN triggers touched: data + tests only, no auth/migration/money/async/cross-repo surface; the registry is consumed by the engine but every behavior change is enumerated + fixtured). EVIDENCE_TARGET at IMPL exit: **E2** (battery + fixture, uncached). Integration/live legs are s5-b's §7 fixtures + the orchestrator's gate — not this plan's claim.

## Operator-judgment items

None. The three m-2 confirms are orchestrator-held integration gates (PROCEED-TO-PLAN §"Integration gates"); no operator decision is pending inside this plan's scope.

## ADDENDUM (20260706 — the scope grant; authorizing record `.relays/s5/s5-a-impl-grant/PLAN-orchestrator-planner-20260706-081355.md`)

The settled data changes (M-3(b) EVIDENCE_TARGET required; the 053113 record_kind narrowing) break legacy battery tests outside the original five-file surface. The orchestrator GRANTED a bounded mechanical extension. This addendum carries the fence and the inventory VERBATIM; it changes nothing else in this plan.

**The fence (hard; violations are deviations).** Edit classes, exhaustively:
- **(a)** Add `EVIDENCE_TARGET` to legacy candidate constructions — assertion-preserving; value per the test's own evidence context (default E1/E2).
- **(b)** Update owed/record_kind seat-scope expectations to the 053113 settled posture.
- **(c)** Crash/applicability fixtures whose expected mutation point is no longer reached: move the candidate PAST the new required field while preserving each fixture's ORIGINAL mutation-point assertion intent — the mutation under test stays the mutation under test.
- **(d)** The ONE named non-mechanical change: `TestOwedItemAcceptsNonOperatorSeat` INVERTS (non-operator owed submission now bounces on seat-scope; operator accepted) and renames to match its new meaning. This is the settled behavior change itself being asserted — not assertion-weakening.
Forbidden inside the grant: any other assertion weakening or deletion; `t.Skip`/disabling; any production/mechanism code edit; any edit to s5-b's surfaces; any fixture whose failure is NOT one of classes (a)–(c) (an unexpected failure class = a fresh escalation, not a fold-in).

**The inventory (implementer, `.relays/s5/s5-a-impl-grant/SITREP-implementer-20260706-082132.md`; full battery run `/tmp/s5-a-full-go-test.json`; no ESC class; no s5-b collision):**

| path | class | note |
|---|---|---|
| cmd/frank-mcp/mcp_test.go | (a) | structured-carrier round-trip: add EVIDENCE_TARGET to accepted + bad submit headers; canonical/non-canonical assertion target preserved |
| internal/engine/config_change_test.go | (a) | add EVIDENCE_TARGET to configChangeRecord(...) so the intended new_digest rejection is reached |
| internal/engine/pipeline_test.go | (a) | five submit-path tests: context-appropriate EVIDENCE_TARGET on the affected records |
| internal/obligation/owed_test.go | (a)+(d) | EVIDENCE_TARGET on operator owed/disposition constructions; the named (d) inversion/rename |
| test/fixtures/f11_test.go | (c) | config-change crash matrix, S2 clean-completion, applicability rows: add/move EVIDENCE_TARGET in the mutation HELPERS so the original crash/applicability assertions remain the target |
| test/fixtures/main_assembly_test.go | (a)+(b) | binary assembly / O3 owed sweep / quarantine: EVIDENCE_TARGET on the submit records. **(b) re-tag, blocker ruling 084234:** the quarantine test's post-quarantine owed leg (`:555` seat-a/implementer submitting owed_item + owed_disposition `:598-619`) was MASKED by the EVIDENCE_TARGET bounce at inventory time; under the settled operator-only posture it needs the class-(b) update — **switch that leg's submitting channel to the operator seat, preserving EVERY downstream assertion** (accepted disposition, quarantine state, owed projection). The test's subject is the quarantine/owed-sweep flow, not owed authorship — the operator channel is the settled legitimate author; the non-operator-bounce semantics are already asserted once by the (d) inversion in owed_test.go, not duplicated here |
| test/fixtures/s4_config_change_test.go | (a) | five s4 §7 tests: EVIDENCE_TARGET on configChangeRelay(...) + the stale candidate |
| test/fixtures/s4_shim_test.go | (a) | four nudge/CC/offline tests: EVIDENCE_TARGET on s4Relay(...) |

(Class (b) has no standalone file: the record_kind posture surfaces only through the (d) inversion; the F11 owed applicability rows are (c).)

**Exit-evidence honesty ([VP-W1], from the grant):** the IMPL report enumerates every legacy-test edit under its class tag and presents them as SETTLED-CONTRACT UPDATES, not regressions; "zero regression" reads against the updated battery with the grant relay as the authorizing record.

**Surface after this addendum** = the original five files + the eight inventoried files. Steps/acceptance/guards otherwise unchanged; execution of (a)–(d) begins only under the superseding `s5-a-impl-r2` dispatch.
