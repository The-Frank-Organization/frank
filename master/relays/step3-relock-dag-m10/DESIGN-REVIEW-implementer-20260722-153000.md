## DESIGN-REVIEW — rev4 MUST-REVISE: the relation preserves disclosure membership but not mutable disclosure bytes, MAX_LEGAL_FRAME still sums ceilings as exact values, and the root refusal family lacks first-match ordering

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m10-review-r4
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — the three findings are deterministic realization gaps inside already-selected semantics
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m10
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m10/DESIGN-planner-20260722-124500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer
BUNDLE_ID: m-10-app-control-plane
RELAY_PATH: master/relays/step3-relock-dag-m10/DESIGN-REVIEW-implementer-20260722-153000.md
SUBJECT: MUST-REVISE exact rev4 `25f3eefb…`: durably freeze the six disclosed fields per emitted frame, replace ceiling-sum equality with an exact constructible frame proof, and order root classification before filesystem resolution

## Verdict

**MUST-REVISE** on exact design SHA-256 `25f3eefbdca1dd0f625109a9d80943bdcb91cec4ed0e2f110bb2754aa1ec775e`.

Rev4 fixes the literal self-erasing-selector defect: after the admission commit, `disclosed_by_turn_id = T` can recover the membership selected for T. It also correctly withdraws the impossible production-`FRAME_MAX` claim, separates production and injected-limit fixtures, and gives all root failure classes typed zero-effect outcomes. Three narrower defects remain.

## M10-DAG-R4-F1 — BLOCKER: `disclosed_by_turn_id` freezes membership, not the mutable six-field disclosure bytes

Rev4 derives every first send and replay from the canonical rows stamped for T and claims the result is byte-identical and byte-stable (`:55`, fixture `:130`). But the closed disclosure row contains mutable `state`, and frozen r40 explicitly says parked rows may reach terminal states through ordinary owner machinery (`r40 :72`). Frozen m-9 r21 likewise treats the two frames as distinct snapshots because the parked set can grow, shrink, or change between `turn_open` and `attempt_open_ok` (`m-9 r21 :145-151`).

The relation stores only which identity was selected. If a stamped row changes after the first emission, reconstructing from the live canonical row either emits changed bytes, emits a now-forbidden terminal `state`, or omits the row. A crash/lost-send replay is then not byte-identical. The same defect invalidates the claim that `attempt_open_ok` is byte-stable within the turn; conversely, forcing it to the admission-time set silently removes the frozen Gate-2 added/changed detection.

Persist the exact six-field disclosure snapshot used by each replayable frame — either bytes on the owning turn/attempt row or a relation whose payload includes the immutable disclosed values — and name separate source rules for `turn_open` replay and `attempt_open_ok` replay. The Gate-2 producer must remain total over equal/added/changed/removed-only until m-9 confirms a replacement contract. Add a cut where a stamped row terminalizes after first send but before replay, plus an added/changed Gate-2 row; assert byte-identical reply replay without hiding the live change from the next pre-work comparison.

## M10-DAG-R4-F2 — BLOCKER: `MAX_LEGAL_FRAME = 3,764,736` is still not an exact constructible size

Rev4 now labels 512 B and 484 B “ACTUAL per-shape maxima” and adds the entire `OVERHEAD_MAX = 65,536 B` as though each were attainable (`:53-60`). The preceding clauses only prove **upper bounds**: member names/punctuation are `≤160` and `≤140`; enum values are bounded by generic caps rather than their longest legal literal; and `OVERHEAD_MAX` is expressly a conservative framing ceiling. JCS has no legal padding that makes those inequalities equalities.

Thus `3,764,736` proves another safe upper bound, but it does not prove a legal frame of exactly that size exists. The production fixture that must build real closed-shape bytes and assert equality to that number (`:60`, `:130`) remains unconstructible for the same reason as rev3, only at a lower target.

Either compute the true canonical maximum from exact member spellings, longest legal enum literals, exact envelopes, and simultaneous legal cardinalities, then provide a real fixture vector attaining it; or retain this number honestly as a conservative upper bound and stop requiring equality to it, while the same injected limits table supplies the constructible exact-fit/one-over pair. Do not call a sum of `≤` ceilings `MAX_LEGAL_FRAME` without an attaining legal witness.

## M10-DAG-R4-F3 — BLOCKER: the root refusal family is not an ordered total decision procedure

The recipe text starts with `realpath` and only afterward describes non-absolute input as `not_absolute` (`:77-78`). An implementation following that order can resolve a relative path against ambient process cwd instead of refusing it. The family also leaves overlaps unordered: a resolved path can be both out-of-grammar and too long, while a symlink can resolve to `/`; no clause says which single reason wins. “Realpath failure of any kind” also overlaps OS length failures with the intended `too_long` class.

Make this one explicit first-match table: required/present and absolute-path check **before any filesystem call**; one `realpath` attempt and its failure mapping; classification of the resolved root `/`; NFC normalization; then an exact precedence for grammar and post-normalization length. State whether `too_long` is only a successfully resolved output check or also classifies an OS length failure. Extend the cross-product fixtures so relative input cannot consult cwd, symlink-to-root returns `filesystem_root`, and out-of-grammar-plus-overlength yields the one prescribed token.

## Accepted on these bytes

- **R3 F1 direction:** the self-erasing boolean is gone; membership is stamped atomically, later-turn selection is disjoint, the three commit/send cuts exist, and m-9 confirmation is correctly named rather than assumed.
- **R3 F2 direction:** production `FRAME_MAX` exact-fit is withdrawn; production and injected limits are separated; the conservative compile-time proof and runtime overflow backstop remain sound.
- **R3 F3 direction:** `/` refusal is within the option allowed by review-r3; the five-member typed refusal family, zero-effects boundary, admitted-domain scoping, and refusal fixtures are the right architecture subject to ordered classification.
- Every earlier accepted item remains accepted: the canonical-row VOID split, manifest union, encoded input gate, descriptor carrier architecture, parked joint wires, disposition/action pair, total broker-result receiver, r10 sweep, M10-C1 telemetry separation, CI-4 death-set split, B/E parking, and confirmation lineage.

## Re-review gate

Return one rev5 additive delta closing R4-F1 through R4-F3. Preserve the accepted membership relation, conservative production assertion, test-only limits seam, and root refusal set. Do not file producer confirmations, the §6-D join, lane completion, F73 closure, re-lock, PLAN, T4/code authority, or implementation from rev4.

## Verification

- Reviewed design SHA-256: `25f3eefbdca1dd0f625109a9d80943bdcb91cec4ed0e2f110bb2754aa1ec775e`.
- Incoming DESIGN relay SHA-256: `7ddd43ee5ac52448b259c3a4222f755bfcdd2d144dac5881ab47dc3fd308d18c`; exact-file lint: `OK`.
- Frozen D-4 producer r40 and m-9 r21 consumer bytes were reopened at the mutable-set clauses; no pending m-9 confirmation was treated as complete.
- No governing design or `frank/` source byte was modified by this review.

ACTIONS_GIT_REF: docs-workspace disk action only — created this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design doc, `frank/` source, branch, commit, lock, PLAN, T4 token, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: none — `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6` after this review relay and INDEX row were written.
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-relock-dag-m10/DESIGN-REVIEW-implementer-20260722-153000.md`.
Next requested action: m-10.planner folds M10-DAG-R4-F1 through F3 into one revised additive delta and returns a fresh uniquely-parented full-byte DESIGN review request; every downstream confirmation/join/re-lock/PLAN/T4/code gate remains held.
