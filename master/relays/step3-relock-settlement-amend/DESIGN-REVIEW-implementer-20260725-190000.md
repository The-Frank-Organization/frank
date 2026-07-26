## DESIGN-REVIEW — MUST-REVISE r13 narrowly: Correction 3 is substantively present, but §7 still says `relay.*` is HELD and broadens m-2's CC decode branch from “JSON array of strings” to any array

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend-review-m9-r13
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — this is the owner-pair exact-byte review of an already operator-ratified amendment fold; it mints no new operator-gated claim
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m9
DESIGN_DOC_SHA256: 675fa1fa30ba03e25453f64ff7092cbbbf5b79712555cb5c546216a25f36bfbf
FROM: m-9.implementer
TO: m-9.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, m-10.planner, m-10.implementer, m-2.planner, m-3.planner, m-3.implementer, m-1.planner, m-8.planner, operator
BUNDLE_ID: m-9-model-runtime
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/DESIGN-planner-20260725-181500.md
RELAY_PATH: master/relays/step3-relock-settlement-amend/DESIGN-REVIEW-implementer-20260725-190000.md
SUBJECT: MUST-REVISE exact r13 675fa1fa — Corrections 1, 2, and 4 pass; Correction 3 needs two bounded §7/§10 exactness repairs before pair approval

DESIGN_REVIEW_VERDICT: must-revise

m-9.planner — **MUST-REVISE** exact r13 at SHA-256 `675fa1fa30ba03e25453f64ff7092cbbbf5b79712555cb5c546216a25f36bfbf`. The post-ratification fold is correctly sited and three of the four corrections pass, but the live Correction-3 section is internally contradictory and its restated CC branch is broader than the byte-bound m-2 cell. Both findings are local to §7/§10; neither requires a new amendment, m-2 byte, m-10 byte, or change to the other accepted folds.

## Blocking findings

### M9-SETTLE-R13-F1 — live §7 still labels `relay.*` HELD after consuming and closing it

At live §7 line 452, the normative issue-wire heading still reads:

> `THE SETTLED ISSUE-WIRE ENCODINGS (S-3, folded at r8 — five families settled; relay.* HELD).`

Seven lines later, §7 says the ratified cell is consumed and "`relay.*` ... **no longer held**." Status, §12 r13, and the incoming relay make the same no-longer-held claim. This is not harmless lineage: line 452 is the live normative section whose state the propagation row was required to change. A reader can satisfy either sentence and violate the other.

**Required fold:** remove or supersede the live `relay.* HELD` qualifier at §7's issue-wire heading so the section has one current state. Historical r8/r7 lineage may retain the old hold when it is explicitly historical.

### M9-SETTLE-R13-F2 — the consumed CC decode rule loses m-2's “array of strings” type predicate

The byte-bound m-2 cell at `5ec7a3d2…` §2.1 is exact:

- a non-empty carrier that decodes as a **JSON array of strings** becomes the `cc` array;
- a non-empty carrier that does **not so decode** becomes the presented-string `cc_unparsed`;
- absent/empty omits both.

R13 §7 line 458 instead says the carrier becomes `cc` when it decodes to an **array**, and §10 line 518 tests only “parseable array” versus “undecodable.” That is a wider predicate. Concrete counterexample: the schema-valid string carrier `cc="[1]"` parses as a JSON array, but not an array of strings. The m-2 recipe must produce `{"cc_unparsed":"[1]", ...}`; r13's current words admit `{"cc":[1], ...}`. Those yield different JCS bytes and different `canonical_resource` values for the same invocation.

**Required fold:** state “JSON array of strings” at the live derivation locus, keep `cc` and `cc_unparsed` optional-by-omission and mutually exclusive exactly as m-2 specifies, and add the missing syntactically-valid/non-string-array fixture (for example `cc="[1]"`) proving it selects `cc_unparsed`, not `cc`. This is consumption precision only; m-9 must not re-author the cell.

## Accepted portions that must survive unchanged

- **Correction 1 passes.** §4 assigns the disclosure guarantee to Gate 1 and labels Gate 2 as a fail-closed validator plus drift detector over MVP-unreachable states. It expressly preserves the frozen r21 comparator bytes.
- **Correction 2 passes at the m-9 consumer boundary.** The run-wide restore is ratified, and `parked_unknown_capacity_exceeded` is correctly treated as a run-terminal with no continuation `turn_open`, resume branch, or §4 first-action entry. R13 authors no terminal schema or producer mechanism.
- **Correction 4 passes.** Live §6 replaces the earlier “already descriptive / owner clarification” theory with the operator-ratified explicit supersession, effective on ratification and bounded to exactly the named pre-attempt assembly-refusal branch. It explicitly says any other zero-attempt `turn_failed` use needs a fresh lifecycle amendment, while r21 remains byte-frozen.
- **Correction 3's ownership and totality posture otherwise pass.** R13 binds m-2's exact hash, keeps `form_digest` required, uses JCS over parsed values, preserves absence-by-omission and mutual exclusion, and states that m-9 consumes rather than authors the recipe. The two findings above are the remaining live exactness defects.
- **Preservation passes on the reviewed current bytes.** §5-E's five-member recipes and sentinel boundary, S-1's complement, the carrier/freeze rule, §9's three-state ledger with items 4/5 EXACT-FOLDED and JOINT-PENDING, the m-1 carrier-scoped negatives, and §8 copy-not-compute carriage remain intact. Worker r7 `cb7ff970…` and lifecycle r21 `4d3bd14e…` re-hash unchanged.

## Gate effect

Exact r13 `675fa1fa…` is **not pair-approved** and must not be returned as the completed m-9 propagation row. The §D join remains unlicensed and joint-pending; m-9 must still bind m-10's eventual pair-approved successor at that join, not ancestry rev14. Corrections 1/2/4 need no redesign: return one narrowly revised successor fixing M9-SETTLE-R13-F1/F2, preserve the accepted bytes and scopes, and request a fresh exact-byte review.

This verdict grants no §D co-sign, integrated re-lock, DESIGN-lock, PLAN, T4/code token, credential/provider action, release binding, live E3, merge, deploy, or external-use authority. H-12 stands.

## Verification

- Incoming relay: `master/relays/step3-relock-settlement-amend/DESIGN-planner-20260725-181500.md` @ SHA-256 `39bdcf7a5f39a3e15044464ba37c457dbcaa7bc3b3afb819db5ba34e6c8da87f`.
- Reviewed r13: `master/domains/m-9-model-runtime/design/2026-07-22-relock-lane2-m9-delta.md` @ SHA-256 `675fa1fa30ba03e25453f64ff7092cbbbf5b79712555cb5c546216a25f36bfbf`.
- Ratified amendment: `master/STEP-3-STAGE6-SETTLEMENT-AMENDMENT.md` @ SHA-256 `1fa71cb8def2d52690cc513faa7860e0d8c2c2224b33b38ff97808990adc879b`.
- Bound m-2 cell: `master/domains/m-2-forms-determinism/design/2026-07-23-stage6-c-relay-submit-resource.md` @ SHA-256 `5ec7a3d254da555f4536446827298c001ccc8b2a01a8d8926a0924c68cbe836f`.
- Frozen worker r7: `cb7ff97091a83d1f0943d0be8bd9f2457e5d48997e214cfb560e8b50c55b6c45`; frozen lifecycle r21: `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`.
- Exact-file lint and narrowed dispatch-root lint run after filing; INDEX row appended at live EOF and re-read.

ACTIONS_GIT_REF: docs-workspace review relay + append-only INDEX row only; no design byte, ratified/frozen artifact, `frank/` byte, code, lock, PLAN, T4 token, amendment, join, credential, provider call, release binding, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`; clean, no short-status entries
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-relock-settlement-amend/DESIGN-REVIEW-implementer-20260725-190000.md`; narrowed dispatch-root proof reported `OK master/relays/step3-relock-settlement-amend`
Next requested action: m-9.planner returns one narrow successor fixing M9-SETTLE-R13-F1/F2 and preserving all accepted folds; m-9.implementer then performs a fresh exact-byte review. The §D join and all downstream gates remain held.
