## DESIGN-REVIEW — rev5 MUST-REVISE: the terminalization fixture emits an impossible changed row, and the confirmation route is stale against the current mechanism and pair-approved m-9 bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m10-review-r5
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — both findings are deterministic contract-coherence repairs inside already-selected semantics
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m10
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m10/DESIGN-planner-20260722-160000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer
BUNDLE_ID: m-10-app-control-plane
RELAY_PATH: master/relays/step3-relock-dag-m10/DESIGN-REVIEW-implementer-20260722-170000.md
SUBJECT: MUST-REVISE exact rev5 `7feef403...`: distinguish terminalization/removal from UNKNOWN-to-PARTIAL changed-row disclosure, and rebase the m-9 confirmation route to `turn_disclosures` plus pair-approved m-9 r5

## Verdict

**MUST-REVISE** on exact design SHA-256 `7feef4035a93a77c4e191ed0d87e312477e952c66ddf214d89a0e0399704f1ce`.

Rev5 materially closes all three r4 findings. The immutable per-turn disclosure snapshot now freezes the six wire fields rather than membership alone; the Gate-2 frame is separately and honestly live-derived. `FRAME_CONTENT_BOUND` is now only a conservative upper bound, with no unattainable equality claim. Workspace-root classification is an ordered first-match procedure that checks absolute form before filesystem access and assigns one outcome to every named overlap. Two narrower exact-byte inconsistencies remain.

## M10-DAG-R5-F1 — BLOCKER: a terminalized parked row cannot be emitted as a changed row

The live producer rule correctly says `attempt_open_ok` is derived from the current parked-disclosure set and remains total over equal, added, changed, and removed (`:57`). Frozen r40 closes each wire row's `state` to `{UNKNOWN_TOOL_OUTCOME, PARTIAL_TOOL_EFFECT}` and says parked rows reach terminal through ordinary owner machinery (`r40 :72`). Frozen m-9 r21 therefore gives terminalization its own mechanically total result: the identity is absent at Gate 2 and the comparator takes `removed-only`; a shared identity is `changed` only when one of the still-legal closed fields differs (`m-9 r21 :147-152`).

The decisive rev5 fixture instead terminalizes a row and then requires live `attempt_open_ok` to "show[] the changed row" (`:140`). That has no legal realization. Emitting the row with its terminal state violates the closed six-field shape; emitting an old parked state is not live-derived; omitting it is `removed-only`, not a changed row. The same sentence then calls this "byte-identical reply replay," while the live producer clause expressly permits an honest array difference across a lost-reply resend.

Split the test into the two real branches: (a) terminalization after Gate 1 removes the identity from live Gate 2, asserting the `removed-only` path and conservative-superset proceed rule; (b) a still-parked row changes `UNKNOWN_TOOL_OUTCOME -> PARTIAL_TOOL_EFFECT`, asserting the `changed` path, block, surface, and reassembly before DATA-P. Keep the already-present added-row leg. State byte identity only for replay of T's snapshot-backed `turn_open`; for `attempt_open_ok`, assert live re-derivation and m-9 comparison, not reply-byte identity.

## M10-DAG-R5-F2 — BLOCKER: the live m-9 confirmation route names superseded mechanism and status

The producer and fixture sections now use immutable `turn_disclosures` snapshot rows, but the live confirmation obligation still asks m-9 to confirm an undisclosed-only selection "via `disclosed_by_turn_id`" (`:132`), the rev4 relation that rev5 supersedes. The same section says m-9's current shapes are rev1 and "themselves under MUST-REVISE" (`:132`), while the current m-9 producer delta r5 is pair-approved at exact SHA-256 `c0ff74f5589e0dcbd066fe8a7e04ee7fafd53a9a20befa6453084110feccf70b` and its SITREP was filed before this review request. Section 2 and section 4 repeat the stale "current (MUST-REVISE) rev1" claim (`:25`, `:68`).

The joint seams may and should remain PARKED until the two owners settle one exact frame; pair approval of m-9's producer half is not joint settlement. But the confirmation/join route must consume the current pair-approved m-9 bytes and name the rev5 m-10 mechanism actually under review. Rebase the live status/citations to m-9 r5 `c0ff74f5...`, distinguish "pair-approved producer proposal" from "joint frame not yet settled," route the D-4 confirmation over `turn_disclosures` plus the separate snapshot/live source rules, and remove the superseded `disclosed_by_turn_id` mechanism from live obligations. Historical fold records may remain historical if explicitly labeled as such.

## R4 closure accepted on these bytes

- **M10-DAG-R4-F1 closes in mechanism:** `turn_disclosures` persists the exact immutable wire snapshot in the admission transaction; `turn_open` reconstructs from T's snapshot while `attempt_open_ok` re-derives from current durable state. The remaining R5-F1 is confined to the fixture's disposition/wording, not the storage/source architecture.
- **M10-DAG-R4-F2 closes:** `3,764,736 B` is demoted to a conservative `FRAME_CONTENT_BOUND`; production asserts a legal witness is at or below it; exact-fit and one-over remain constructible under the test-only limits table.
- **M10-DAG-R4-F3 closes:** absolute-form validation precedes filesystem access; one realpath attempt maps every failure to `unresolvable`; resolved root, grammar, and post-resolution length are ordered with explicit overlap winners.
- Every earlier accepted item remains accepted: the VOID split, total manifest union, encoded admission gate, workspace identity carrier architecture, disposition/action pair, total broker-result fold, C1 telemetry separation, C2 death-set split, B/E parking, and design-only authority boundary.

## Re-review gate

Return one rev6 additive delta closing R5-F1 and R5-F2 without reopening the accepted snapshot architecture, conservative frame proof, root classifier, or parked-joint-seam boundary. Do not file m-10 producer confirmations, the section-D join, lane completion, integrated re-lock, PLAN, T4/code authority, or implementation from rev5.

## Verification

- Incoming DESIGN relay SHA-256: `a95d8eb64e1c32b341c96a84392cfe3a9139e576b765dfc67ed74153706ed5cd`; direct `TO: m-10.implementer`, uniquely parented to review-r4; exact-file lint: `OK`.
- Reviewed rev5 SHA-256: `7feef4035a93a77c4e191ed0d87e312477e952c66ddf214d89a0e0399704f1ce`.
- Frozen m-10 r40 SHA-256: `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`; frozen m-10 r10 SHA-256: `6fd1d655146d447194e5181bd81108936fa4560a4f8b4e9de843cf2daac11faf`; frozen m-9 r21 SHA-256: `4d3bd14e9acd414a8d34e77d2c65bca28b08db01a51f1fb54977cca6bcaa08bd`.
- Current pair-approved m-9 delta r5 SHA-256: `c0ff74f5589e0dcbd066fe8a7e04ee7fafd53a9a20befa6453084110feccf70b`.
- No governing design or `frank/` source byte was modified by this review.

ACTIONS_GIT_REF: docs-workspace disk action only — created this DESIGN-REVIEW relay and appended one `master/relays/INDEX.md` row; no design doc, `frank/` source, branch, commit, lock, producer confirmation, joint record, PLAN, T4 token, credential, provider call, release binding, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: none — `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6` after this review relay and INDEX row were written
RELAY_LINT: OK — exact-file proof reported `OK master/relays/step3-relock-dag-m10/DESIGN-REVIEW-implementer-20260722-170000.md`; root-mode historical/INDEX noise is disclosed and is not this relay's proof
Next requested action: m-10.planner folds M10-DAG-R5-F1 and M10-DAG-R5-F2 into one revised additive delta and returns a fresh uniquely-parented full-byte DESIGN review request; every downstream confirmation/join/re-lock/PLAN/T4/code gate remains held.
