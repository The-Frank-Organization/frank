## RECONCILE -- REVISE: rev3 adds real tables, but soft edits still move the HARD digest, E's aggregator cannot see provider lowering, journal recovery loses required content, and the exit gate can pass failing evidence

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-stage6-rescope-decomposition-review-r3
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- the existing operator re-scope gate remains required, but amendment rev3 is not ready for it until the five bounded contracts below are internally executable and non-contradictory
GRILL_REQUIRED: no -- the existing GRILL_LOCK remains sufficient; these are mechanical correctness defects, and the operator may ratify or adjust one unambiguous overhead threshold set at the already-scheduled re-scope gate
DESIGN_DOC_ID: step3-stage6-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-050000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-5.planner, m-6.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: REVISE -- preserve rev3's accepted closures, but separate source provenance from the HARD digest, bind the actual shell/env/resource context, put provider-lowering aggregation where its bytes exist, make journal content crash-reconstructible, and prevent applicability-only/fixture-placeholder exit passes

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-050000.md` at SHA-256 `d8a22e1132ad64b2b8ec7be7cb29736ed43f90d4dd86dbee55b9978004e2a54a`.

Proposed amendment rev3: `master/STEP-3-STAGE6-AMENDMENT.md` at SHA-256 `419c3793ec6f722274741c5a2aca0ed4ff5841460b0c4820759f10829ce38fb2`.

## Findings

### F101 -- BLOCKER: a Tier-SOFT edit still changes the proposed Tier-HARD bundle digest

Amendment `:61-75` puts each mixed source's full-file `source_sha256` inside the JCS manifest, then defines the bundle digest over that manifest. Any cosmetic Tier-SOFT edit changes the full-file source hash, changes the JCS manifest, and therefore changes the bundle digest. That directly contradicts `:75` ("a Tier-SOFT edit does not") and recreates the all-artifact F73 behavior this re-cut exists to remove.

The contract also says the extractor recipe digest is recorded in the manifest (`:70-73`), but that field is absent from the declared entry schema, and the fail-closed list rejects missing declared markers without rejecting an undeclared hard marker omitted from the manifest.

Required correction: define a canonical lock payload whose digest includes recipe identity, ordered interface IDs, and extracted HARD bytes/digests, but excludes mixed-document full-source provenance. Keep each mixed source's full-file SHA in a separately authenticated generation/provenance section that `--verify` checks without feeding the constitutional lock digest; a whole-file-hard source remains locked by its extracted/full-file digest. Pin the top-level artifact schema and locations of `recipe_sha256` and `bundle_sha256`, and reject undeclared/duplicate marker IDs across the declared source inventory. Demonstrate with a negative fixture that changing bytes outside a HARD region leaves `bundle_sha256` unchanged while changing the source-provenance hash.

### F103 -- BLOCKER: the descriptor still does not identify the actual environment, shell interpreter, or multi-target `apply_patch` invocation

Rev3 correctly moves context identity and teardown into HARD and keeps containment out. Three table/identity defects remain:

- `env_profile_digest` is over an "allow-listed profile" (`:106-114`), but no allow-list artifact/owner or rule says the child receives exactly that set with no inherited extras. The frozen worker currently states bash cwd/env are ambient and unpinned (`m-9 full worker:45,178`). Hashing a subset does not bind the actual invocation environment; clearing to an exact allow-list would be a behavior delta that must be stated and reviewed.
- `tool_impl_ref` points at the F58 worker build, which covers the worker/backend code but not an external host shell executable. The actual command interpreter path/bytes/version are not thereby identified. `backend_id="ambient"` plus the worker digest still permits a different host shell to interpret the same command.
- The applicability table groups `apply_patch` with single-path tools and requires one `canonical_resource` (`:93-104`), while its frozen schema carries only a patch string and a patch may name multiple paths. One canonical path cannot represent that invocation even at intended-resource grain.

Required correction: state whether the child environment is exactly a closed profile or ambient; in either branch digest the complete environment actually presented, with m-1 reviewing secret-sensitive material. Bind the resolved shell executable/runtime as a release material or narrow `tool_impl_ref` explicitly to wrapper identity and add a separate interpreter reference. Split the per-action rows: define `apply_patch` as a deterministic ordered target set/digest or mark resource inapplicable and rely only on the exact patch args digest without claiming a single resource. Make cwd's canonical byte representation unambiguous. No sandbox, containment, or affected-resource inference is required.

### F104 -- BLOCKER: B is corrected, but E assigns final aggregation to a process that never sees provider lowering

Section 6 B (`:156-161`) now correctly separates m-3 schema production, m-8 digest production, sibling m-9/m-10 consumers, and the later evaluator join. C, D's ownership shape, broker-first placement, and conditional H-24 also remain sound.

E is not executable. The current topology sends m-9's logical `LLMRequest` to m-8, and m-8 alone performs provider-specific lowering immediately before the wire. Amendment `:142-154,164-165` nevertheless assigns the final manifest/digest aggregator to m-9 because it "assembles the surface it presents." m-9 neither owns nor observes the actual lowered provider tool representation. No new m-8-to-m-9 component frame, pre-send receipt, attempt-row field, or E0/E3 carriage is defined, so m-9 cannot hash that component and the stated attempt binding has no wire.

Required correction: distinguish the logical model surface from the provider-visible lowered surface, then pick one executable ownership shape. Either m-8 produces a canonical lowering component/digest and carries it through a defined pre-send/attempt record for a joined digest, or m-8 owns the final provider-visible aggregation while m-9 supplies a logical-surface component digest. Define exact component schemas, timing before provider send, attempt/m-10/E0/E3 carriers, and the independent observer derivation. Do not duplicate m-8's translation in m-9.

### F105 -- BLOCKER: the source-map rows are present, but one source is non-reconstructive and the allowed crash cut loses settled content

The table at `:123-140` is useful progress, and first-durable content is not inherently a second outcome truth. The map is not yet truthful enough to resume:

- It calls parsed tool calls reproducible from m-10 `tool_authorizations` (`:127`). That row stores `tool_call_id`, canonical name, and `canonical_args_digest`, explicitly never the argument payload (`m-10 seam contract:183-194`). A digest cannot reconstruct the `assistant_tool_call{...,arguments}` item required in model input. The exact assembled arguments need a first-durable content blob or a reference to another named canonical content record.
- The commit rule permits the journal round to commit "strictly after" canonical outcome rows (`:134-139`). A crash after `OUTCOME_RECORDED` but before result-content/checkpoint persistence leaves a known effect with lost model-visible result content. Simply resuming the prior jointly committed round can reissue or forget an already-executed effect. The design names no round-commit marker, state-sequence boundary, or fail-closed `CONTENT_LOST`/park disposition for this cut.
- The current replay envelope is provider-minted content kept only in m-9 in-memory turn state and needed for same-turn feedback (`m-8 provider contract:44-46,58-69`; m-9 lifecycle:160-166). Rev3 says provider-output persistence excludes reasoning bytes and never disposes replacement recovery when that envelope existed. It must either persist the opaque envelope under its existing no-surface/redaction constraints or explicitly make that replacement non-resumable/new-attempt with an honest typed state.
- The required retention/GC and per-record/blob size bounds remain absent. "Content-addressed" and "m-1-reviewed" do not define them.

Required correction: make every model-visible item reconstructible from a named blob/record; define a durable round marker and one crash-total transaction protocol relating content blobs, tool/attempt outcome rows, and the resumable checkpoint; forbid silent resume across a known-effect/lost-content gap; settle replay-envelope recovery; and add retention/GC, size, integrity, and orphan-blob rules. Physical schema may remain pair-owned, but these truth/linearization branches are architecture.

### F106 -- BLOCKER: the exit gate confuses evidence applicability with predicate success, and its fixture/budget identities remain post-selectable

The H-12 precedence note, trusted-local envelope, prohibited-use categories, operator gate ownership, and dogfood boundary at `:206-221` now satisfy the r2 governance request. The operator may adjust the proposed overhead numbers at the existing ratification gate; they become immutable for T4 once ratified.

The Step-3 exit contract still admits false or undefined passes:

- Governance-binding passes when all three records are merely `applicable` (`:175-177`). The frozen m-3 evaluator defines `applicable` only as current identity/vector binding; an applicable record separately carries `observed_outcome=pass|fail` (`m-3:158-165,195-205`). Rev3's new typed predicates likewise carry `verdict=pass|fail|unknown` (`:149-154`). As written, a current, correctly bound predicate with `verdict=fail` passes the exit leg.
- `xit-*` strings are labels, not byte-bound fixture identities. The fixtures are built later at T4 (`:171-183`) and have no manifest path/hash, input artifact hash, deterministic fault cut, expected canonical rows, observer, or exact evidence locator. The injection repo, local write, crash effect, and handoff can still be selected or interpreted after results are known.
- The budget calls p50 <=20% a ceiling but says hard failure occurs only above 100% (`:184-189`), leaving 20%-100% with no gate verdict. "At least" 30 turns/100 calls also leaves sample selection and fixture weighting open.
- Operability's "legible" and handoff's "consumes" are not machine predicates, and crash-honesty's "no duplicate effect" has no instrumented effect counter or canonical observer.

Required correction: require both E3 applicability and typed predicate/observed outcome `pass`; any `fail` fails and any unavailable/unknown holds. Freeze a fixture manifest or deterministic generator digest before T4 with exact inputs, fault injection points, expected records/effects, observers, and evidence locators. Give every leg a total pass/fail/unknown rule over structured fields. Choose one total overhead rule (including the 20%-100% interval), a fixed paired sample/weighting method, and baseline artifact/config identity. H-12 needs no further grill unless that work changes its accepted envelope.

## Accepted closures and direction preserved

- **F102 remains CLOSED.** m-10 remains the F59 authorization host, m-9 the executor/consumer, and m-5/m-6 stay out by removal-not-reassignment.
- **F101 direction is accepted:** explicit markers, dedicated bundle tooling, deterministic ordering/JCS, fail-closed verification, master ownership, F73 review, and a dedicated soft-design ledger are the right shape once source provenance is removed from the HARD digest.
- **F103 direction is accepted:** exact-context evidence, no containment claim, HARD teardown/UNKNOWN semantics, and the evidence-vs-H-21 authorization boundary stand.
- **F104 B is CLOSED.** C, the coordinated D seam, m-7-first placement, and H-24 conditional stand. Only E's ownership/carriage remains open.
- **F105's first-durable-content principle is accepted.** Content persistence does not become a rival outcome truth when outcomes remain in m-10 canonical rows; the defects are source completeness and crash linearization.
- **F106's product choices and H-12 boundary stand.** Utility remains non-gated; dogfood and exit remain decoupled; no new grill is requested.
- No frozen design artifact moved. The nine design finals, H-16 rev16, and H-17 census v3 remain byte-identical to the r4 set.

## Gate disposition

- Proposed stage-6 amendment rev3 `419c3793...`: REVISE; not ready for operator re-scope ratification.
- Pending joint-lock proposal `b7e1f0ef...`: remains HELD/superseded; no lock issued.
- Stage-6 re-lock remains unrequestable until the corrected domain deltas, executable bundle, reviews, and joins complete.
- PLAN, T4/code token, credentials, provider calls, release binding, live E3, merge, deploy, out-of-envelope use, and Step-3 close remain HELD. Step 2 remains closed.

## Required return

Return decomposition review r4 over new amendment bytes that: (1) excludes mixed-source SHA provenance from the HARD digest and proves soft-edit stability; (2) binds actual child env/interpreter and fixes per-action resource applicability; (3) makes m-8's provider-lowering contribution and digest carriage executable; (4) makes every journal item reconstructible with crash-total checkpoint semantics; and (5) binds fixtures and requires semantic predicate success under a total overhead verdict. Preserve every accepted closure above.

## Verification

- Target SHA-256: `d8a22e1132ad64b2b8ec7be7cb29736ed43f90d4dd86dbee55b9978004e2a54a`; exact-file lint: OK; directly addressed and indexed at pre-review EOF row 1876.
- Amendment rev3 SHA-256: `419c3793ec6f722274741c5a2aca0ed4ff5841460b0c4820759f10829ce38fb2`; backlog SHA-256: `4607b928cfdd7eeadac019baa0874b822d03d3614fe77009650a43e0ad3d1a6d`.
- r2 reviewer parent SHA-256: `03c198dea9de7d869acc244eebbcc6692894ae6854ea4d57dec07825eee32c90`.
- All nine prior design finals, H-16 rev16 `a349a329...`, and H-17 census v3 `959b1928...` recompute to the r4 manifest.
- `frank/` is clean at `c78da3815a34480590071295c1e09bb7d53c10b6`; no product-source action was taken.
- Clock disclosure: the parent self-stamps `050000` but its filesystem mtime is `04:35:09-0700`; this review uses the live wall-clock `043904`. Causality is exact through `IN_REPLY_TO`, the parent hash, and append-only index order.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no amendment/backlog/design/source byte, historical relay, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, deploy, or out-of-envelope action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` with empty porcelain at `c78da3815a34480590071295c1e09bb7d53c10b6`.
RELAY_LINT: OK -- exact-file proof rerun after the append-only INDEX update and this final report-byte edit.
Next requested action: master folds the five bounded mechanical corrections and returns amendment rev4 for decomposition review r4; operator re-scope ratification remains held.
