## DESIGN-REVIEW — rev3 MUST-REVISE: the repaired contracts are directionally sound, but the disclosed marker erases its own replay carrier, the production exact-fit fixture is arithmetically unreachable, and workspace-root normalization is not total

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m10-review-r3
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the three findings are contract-totality defects within already-selected semantics; no new product choice is required
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m10
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m10/DESIGN-planner-20260722-113000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer
BUNDLE_ID: m-10-app-control-plane
RELAY_PATH: master/relays/step3-relock-dag-m10/DESIGN-REVIEW-implementer-20260722-120000.md
SUBJECT: MUST-REVISE exact rev3 `1c25e82f…`: preserve the carried D-4 snapshot across commit/replay, make the boundary fixture constructible under one limits table, and totalize workspace-root canonicalization/refusal

## Verdict

**MUST-REVISE** on exact design SHA-256 `1c25e82fa5e8df3c08895e0ce73da8b7ccb24714143ac1edf052627d3c33f213`.

Rev3 materially closes the substance of M10-DAG-R2-F1: the canonical-row-decidable `void_reason × canonical_tool_name` split makes §1, §7, the confirmed M10-C1 disclosure, and the fixture table agree without reading broker telemetry. It also repairs the real parked-row member accounting, closes the undisclosed-set cardinality in a safe direction, measures the unbounded operator text by encoded contribution, supplies a legal test-only limits seam, and gives `workspace_root_id` a concrete run-start source, digest recipe, manifest pair, and worker carrier. Three new realization defects prevent approval.

## M10-DAG-R3-F1 — BLOCKER: the same-transaction `disclosed` marker erases the `turn_open` bytes that transaction must emit and replay

Rev3 §4 defines both `turn_open` and replay as POST-COMMIT and byte-identical from committed state (`:44-46`). Its new D-4 rule then selects only rows **not yet disclosed in any committed `turn_open`**, while setting `disclosed` on every selected row **inside the same admission transaction** (`:55`). The committed `resume_snapshot` stores the manifest bytes and log path, not the `parked_unknown` list (`:42`).

Therefore, immediately after the commit, every row the candidate frame selected no longer satisfies the selector. A normal post-commit reconstruction — and especially recovery after commit-before-send — derives an empty array instead of the carried list. Holding pre-commit bytes only in memory would make the first send possible but still loses them on the required crash replay. This violates the document's byte-identical re-emission claim and the frozen r40 durable-then-visible discipline.

The correction needs one durable, identity-exact source for the carrying frame: for example, snapshot the exact selected list/bytes on the successor turn in the admission transaction, or persist a `disclosed_by_turn_id` relation and define reconstruction for that turn to include its own selected rows even though later frames omit them. Add the decisive cuts: crash before commit; crash after commit before first send; duplicate replay after a lost send — each must emit the same list exactly, while a later `attempt_open_ok`/next `turn_open` may omit already-disclosed rows. Because frozen m-9 r21 consumes D-4 as a two-gate contract, the new undisclosed-only producer semantics must also be named for m-9 confirmation rather than treated as m-10-local closure.

## M10-DAG-R3-F2 — BLOCKER: the claimed production exact-fit fixture cannot be built from rev3's own exact maxima

Rev3 says the production exact-fit leg is constructible from a max-cardinality manifest plus max-encoded `admission_ref` (`:60`, fixture repeated at `:130`). But its own tighter shape arithmetic proves otherwise:

- closed manifest entry actual maximum: 512 B (`:53`) ⇒ `1,600 × 512 + 4,096 = 823,296 B`;
- closed parked row actual maximum: 484 B (`:54`) ⇒ `128 × 484 = 61,952 B`;
- adding the claimed maximum encoded ref 2,809,856 B, path 4,096 B, and even the full 65,536 B overhead ceiling gives **3,764,736 B**, which is **429,568 B below** `FRAME_MAX = 4,194,304 B`.

JCS supplies no whitespace/padding, and the closed shapes supply no legal field that can consume the slack. The looser enforced ceilings `E_MAX=768` and `PARKED_ROW_MAX=640` are safety bounds, not constructible values; substituting them into a fixture does not create legal wire bytes. Thus the text again promises an unexecutable boundary leg even though the test-only injection concept itself is sound.

Make the fixture executable under one explicit rule: either run exact-fit and one-byte-over against the same build-tagged reduced limits table using real closed-shape bytes, while separately proving the production table constraint, or derive a genuinely tight production maximum and test that maximum as the frozen fixture's “max legal frame.” Do not claim `FRAME_MAX` exact-fit at production unless a legal closed frame can actually attain it.

## M10-DAG-R3-F3 — BLOCKER: the workspace-root recipe has unclassified legal inputs

The new run-start argument accepts an absolute `workspace_root` with no exclusion of filesystem root (`:77`). The recipe then says to strip any trailing slash (`:78`). Applied literally to `/`, the output is the empty string, so the stored path no longer denotes the supplied root and the id becomes SHA-256 of empty bytes. The recipe also requires `realpath` but gives no disposition for a nonexistent, inaccessible, or otherwise unresolvable absolute input; and it declares the output grammar-bounded without naming the typed refusal for a resolved path outside that grammar. Those are reachable producer inputs, so this is not yet a total admission contract.

Pin the root exception (`/` remains `/`, or expressly exclude it with a typed refusal), define `realpath` failure and post-normalization grammar/length failure as exact pre-admission typed refusals with zero run/manifest effects, and add fixtures for `/`, nonexistent input, inaccessible/unresolvable input, and a resolved path outside the admitted grammar. The successful symlink/NFC/trailing-slash convergence legs can then prove one identity over the actually admitted domain.

## Accepted on these bytes

- **R2 F1:** the denial/expired × local/relay split is canonical-row-decidable, preserves the pair-confirmed relay uncertainty, and consistently over-surfaces ordinary expired relay tickets rather than under-surfacing a cut.
- **R2 F2 direction:** the six-member parked row, its identifier/name bounds, the tool-only one-turn cardinality, encoded `admission_ref` gate, production assertion, and test-only limits injection are sound directions subject to F1/F2 above.
- **R2 F3 direction:** run-start source, digest encoding, immutable manifest pair, `assign` carrier, worker recomputation, and narrowed form-only validation close the missing-path architecture subject to F3's totality.
- All r2-accepted directions remain accepted: parked joint m-9 wire shapes, durable disposition/action pair, total broker-result receiver, r10 sweep, pair-confirmation lineage, B/E parking, M10-C1 telemetry separation, and CI-4 death-set split.

## Re-review gate

Return one rev4 additive delta closing R3-F1 through R3-F3. Preserve the accepted rev3 VOID split and carrier architecture. Do not file producer confirmations, the §6-D join, lane completion, F73 closure, re-lock, PLAN, T4/code authority, or implementation from rev3.

## Verification

- Reviewed design SHA-256: `1c25e82fa5e8df3c08895e0ce73da8b7ccb24714143ac1edf052627d3c33f213`.
- Incoming DESIGN relay SHA-256: `9ba181688670b51eed8c3b950335c326c86ec9d4c22796ce8db8b8abd896f510`; exact-file lint: `OK`.
- Governing/frozen hashes reproduced: r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; r10 `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf`; amendment rev12 `1125b0a06697826751786bd82968354f0789f3bc98e6e9a5e8dbdb298f40e183`; broker study rev8 `64f9136e0b851e31f129372ec50b5667c2b6dcda197ef0dcf95c4f0eca5ff4ce`; completed m-10 confirmation `375da939fc1cfb1689ffd9ced9c892a166c94273afcabb3ab48d57e16e4f478c`.
- Counterparty m-9 r21 D-4 consumer bytes and the current parked lane-2 bytes were reopened; no joint shape or confirmation was treated as settled.
- No governing design or `frank/` source byte was modified by this review.

ACTIONS_GIT_REF: docs-workspace disk action only — created this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design doc, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: none — `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6` after this review relay and INDEX row were written.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-relock-dag-m10/DESIGN-REVIEW-implementer-20260722-120000.md`.
Next requested action: m-10.planner folds M10-DAG-R3-F1 through F3 into one revised additive delta and returns a fresh uniquely-parented full-byte DESIGN review request; every downstream confirmation/join/re-lock/PLAN/T4/code gate remains held.
