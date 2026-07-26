## DESIGN-REVIEW — m-9 lane-2 r6 must revise: assembly refusal has no exact lifecycle disposition; §9 still denies the fold it just records

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m9-review-r6
PARENT_DISPATCH_ID: step3-relock-dag-m9
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — both findings are internally determinate contract corrections; no operator-ratified product choice is reopened
GRILL_REQUIRED: no — the approved m-2 component fixes the product semantics; this review requires exact m-9 realization and self-consistency
DESIGN_DOC_ID: step3-relock-dag-m9
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m9/DESIGN-planner-20260723-051500.md
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-3.planner, m-3.implementer, m-2.planner, m-10.planner, m-10.implementer, m-1.planner, m-8.planner, operator
BUNDLE_ID: m-9-model-runtime
OWNER: m-9 (Model Runtime)
RELAY_PATH: master/relays/step3-relock-dag-m9/DESIGN-REVIEW-implementer-20260723-061500.md
SUBJECT: MUST-REVISE exact r6 fe0ac866 — the new ASSEMBLING refusal does not select an existing terminal/exit path, and §9's closing sentence still claims none of five inputs is consumed although item 1 is folded

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — I freshly reviewed the complete r6 artifact at exact SHA-256 `fe0ac866cd6c5a7171f1f0ccddc43ca021aa586d7fd7cf14d1a56c5c97363efe`, the directly addressed r6 relay, master's bounded routing, m-2's pair-approved component at `c3a8cd61bcbe33a3ede847b29d7416563000c66f949fa2950ea0af305edbbd2c`, the frozen m-9 worker/lifecycle bases, and m-10 rev6.

**MUST-REVISE.** The m-2 ownership/recipe fold is directionally faithful and most of §6 closes, but two live contradictions prevent byte-bound approval.

## M9-DAG-R6-F1 — `assembly REFUSES` has no exact lifecycle disposition after admission

Section §6 says a missing/extra/duplicate/cross-array mismatch makes assembly refuse, produces no digest, opens no attempt, and “surfaces through my existing fail-closed path.” That proves the negative provider boundary, but it does not define the state-machine outcome.

The failure occurs after `turn_open`: frozen lifecycle r21 §2.2 places the worker in `ADMITTED → ASSEMBLING` before `ATTEMPTING`/`attempt_open`. At this cut m-10 already has an admitted active turn, while there is intentionally no attempt row. The same frozen contract provides a closed m-9-owned machinery terminal, `turn_failed`, and the existing exact frame `turn_terminal{run_id, turn_id, turn_epoch, terminal}`. “Existing fail-closed path” does not select that route, a generation exit/supervision route, or any other one. A builder may therefore leave the admitted turn hanging or invent lifecycle behavior while still satisfying the current prose.

Required correction:

1. Select the exact existing m-9 lifecycle branch. If the intended answer is the evident one, state: totality/shape failure while `ASSEMBLING` ⇒ no `logical_surface_digest`, no `attempt_open`, no DATA-P request, then the existing `turn_terminal{..., terminal: turn_failed}` path; no provider-attempt row and no new m-10 disposition/schema. If another route is intended, it must be owner-real and routed rather than inferred here.
2. Add the m-9-owned lifecycle fixture at §10 or an explicit byte-bound fixture reference: each refusal class, including alias-normalization collision, yields zero digest, zero `attempt_open`/DATA-P, zero provider-attempt row, and exactly the selected terminal/exit behavior. m-2 §6 tests component refusal, but cannot choose or prove m-9's post-admission lifecycle transition.

This correction preserves the requested boundary: it needs no new m-10 attempt disposition and does not reopen the m-2 recipe.

## M9-DAG-R6-F2 — §9 contradicts the r6 discharge

Section §9 item 1 says the m-2 component is **DISCHARGED at r6**, §6 consumes it byte-bound, Status says four parked inputs remain, and §11 repeats four. But §9's live closing sentence still says: “Nothing in this delta consumes any of the five as though it were settled.” That statement is now false on these exact bytes.

Required correction: scope the sentence to the four remaining parked inputs, for example: “Nothing in this delta consumes any of the four remaining parked inputs as though it were settled.” Preserve item 1 as discharged and preserve items 2–5 as parked.

## Passed surfaces to preserve

- The m-2 recipe/producer split is consumed correctly: m-2 owns the recipe and three live relay-schema renders; m-9 owns five local schemas, all eight tool-level description strings, assembly, and hashing.
- The description source is buildable at this design grain: one value from the worker's presented tool-level registry per canonical name, `""` when none, with no promotion from R-3's schema-property annotation.
- Totality, canonical-name ordering, alias normalization, closed element shapes, JCS over parsed values, no component pre-hash, and the Go/JCS divergence warning match m-2 rev2.
- INV-E1 remains exact: local presented schemas are the §8.3 pins; relay presented schemas are live renders; descriptions move `logical_surface_digest` but not F58 `tool_catalog_digest`.
- The mapping-version binding uses the existing attempt→run→manifest→`tool_set` relation, adds no sixth member, and fails closed to E3 non-applicable when unresolved.
- The six §D settlement folds remain held; no m-10 B/E carriage, m-3 binding, joint record, PLAN, T4, or code authority is implied.

## Exact evidence

- Addressed r6 relay SHA-256: `44db35958d7c4abb54818aa8a0422e454c72577ccd9e3b2c13314d9ac7529c6e`; exact-file lint OK.
- r6 design SHA-256: `fe0ac866cd6c5a7171f1f0ccddc43ca021aa586d7fd7cf14d1a56c5c97363efe`.
- m-2 component SHA-256: `c3a8cd61bcbe33a3ede847b29d7416563000c66f949fa2950ea0af305edbbd2c`; its approving review exact-file lint OK.
- Frozen m-2 stage-1 SHA-256: `83d8e63ec6a32726a57f715b69d96fc3867e178b35f2aa57c19373de264a0f7d`.
- Frozen m-9 worker r7 SHA-256: `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`.
- Frozen m-9 lifecycle r21 SHA-256: `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`.
- m-10 rev6 SHA-256: `29a123fe8f153127ccfe03559850c8e5ce4c488671402cebf8325e3cafbcd7ae`.

## Authority boundary

r5's prior approval remains void because r6 superseded its bytes; r6 is not approved. m-3 must not bind `fe0ac866…`, and m-10 must not begin the dependent B/E carriage fold from this producer hash. Fresh corrected bytes and a fresh full-byte review are required. The six §D folds and all integrated re-lock, PLAN, T4/code, credential, provider, release, merge, and deploy gates remain held.

ACTIONS_GIT_REF: docs-workspace review action only — created this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design document, dependency, frozen artifact, producer confirmation, joint record, `frank/` source, branch, commit, PLAN, T4, credential, provider call, release binding, merge, or deploy action
FINAL_GIT_STATUS_SHORT: unavailable — harness root is not a git repository; `frank/` clean at `c78da3815a34480590071295c1e09bb7d53c10b6`
RELAY_LINT: OK — exact-file and dispatch-root lint exit 0
Next requested action: m-9.planner folds M9-DAG-R6-F1 and M9-DAG-R6-F2 into fresh r7 bytes, preserves the passed §6 mechanisms and every held §D boundary, and returns a fresh uniquely parented full-byte DESIGN-REVIEW request. Downstream m-10/m-3 sequencing remains held meanwhile.
