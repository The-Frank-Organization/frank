## PLAN-REVIEW — s16a WP1 PLAN r9 approved

ROLE: Implementer
PHASE: PLAN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: s16a-build-plan-review-4
PARENT_DISPATCH_ID: s16a-build-9
RUN_ID: s16a
CEREMONY_TIER: medium
EVIDENCE_TARGET: E2
HUMAN_GATE_REQUIRED: no — s16a.planner must re-run the mechanical SCOPE_DIFF and issue a successor IMPL relay re-binding the token to r9; merge remains operator-gated
PLAN_LOCK_ID: s16a-build-9
IN_REPLY_TO: s16a-build/PLAN-planner-20260824-213626.md
FROM: s16a.implementer
TO: s16a.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator
SUBJECT: formal PLAN-REVIEW approve — exact PLAN r9 def78976 faithfully enumerates the adopted m-9/m-10 fidelity strengthenings, preserves the seam-only fence, and holds WP2 for bounded PM re-confirm
PLAN_REVIEW_VERDICT: approve

`PLAN_REVIEW_VERDICT: approve` for exact PLAN r9 SHA-256 `def789764a2b5d8e8c9fff19161747665bfe1b00693a0ccaf2c981b8bb761603` at `s16a-build/PLAN-planner-20260824-213626.md`.

Findings: none.

The v2.9.1 relay engine has admitted and rendered those exact PLAN bytes at the same digest. Client and daemon both report kit `2.9.1` and fingerprint `ddd5f9f5c84e3c51dbe06ce7d1d7174604b5f0c17524c155091f3c6738bcc6c0`; daemon state is ready, epoch active, and pending renders zero. The one reported projection conflict is the inherited historical INDEX divergence, not a PLAN r9 admission or byte-integrity failure.

The authority chain is exact: PLAN r9 is FROM `s16a.planner`, TO this seat, `PHASE: PLAN`, `AUTHORITY: plan-only`, `DISPATCH_ID: s16a-build-9`, and parented to master's addressed fidelity ruling `s16a-wp1-fold`. It contains no live implementation or merge token. The ruling digest is `265050fcce7a5d5d60ef7fe7bf952b014a3c96fbb6df4efa6a1e07ea07868801`; the four normative PM-return digests re-compute exactly as r9 declares: m-9 `1e88c0a6d41e6eb9528430e1415d881ea26941a37e222232b93aabe7ed52d3a5`, m-10 `ebf2c19456d02594b4646845f385087f7c5ed5eb71fbfc171d8a86f78f9411a2`, m-7 `e8d197f0978e8933b75000fa427c160072325ae809177003078f6f0509fa3de0`, and m-8 `5f9e8d79470def42bd3bbfaeda08a474b9e815c92f8f629289c5d7f8fef6c3b1`.

Reviewed and approved surfaces:
- The work-list arithmetic is complete: m-9's 22 rows plus m-10's 10 rows, with A13/A14 shared, yield 30 unique do-not-bind row names; G04/G08 make 32 primary changed rows. The additive B09/A16/C06 riders make 35 changed test functions total. No row outside the 64-row ledger is introduced and D03 remains absent.
- Each named PM strengthening is preserved: fourteen source-token predicates gain reached behavioral legs; A03/A13/A14 gain effect/recomputability legs; A02/A05/A10/A12 gain semantics-and-locus/demux/registry-effect legs; B10 reaches the worker opaque-item path; m-10's A08/A13/A14/A19/C02/C03/C04/C05/C08/D04 deltas are carried; and G04/G08 are forward-fixed against C01/A10.
- B08 correctly retains m-8's exact enum-membership subpredicates and adds m-9's total-consumption/unknown-token behavioral leg. This resolves the two PM column rulings without deleting a binding producer-owned predicate.
- The §2.7 riders are additive sublegs only. The already-settled B09/A16/C06 predicates are not rewritten or reopened; the later bounded PM re-confirm covers only the new sublegs, preserving master's ruling that the previously binding predicates are settled.
- A02's approved locus reading is strict: the invalid nonce must be rejected before the journal-open/genesis path, proved by a pre-journal control-flow witness. Filesystem non-creation alone is insufficient because today's `journal.Open` validates the identity before creating the journal and would make that weaker predicate green at the wrong locus.
- Scope remains exactly `frank/test/seam/**` plus daemon-owned s16a lane relays. Production/source bytes, contracts/ledger/`master/**`, dependencies, `go.mod`/`go.sum`, governed s8 machinery, later WPs, conforming-side fixes, new rows, merge, push, and PR work remain forbidden.

Boundary contract: r8 §6 remains the exact inherited contract. The strengthened tagged test package writes the row-bound launch instrument; the script-derived census remains the target entity; m-9/m-10 bounded re-confirm and WP2–WP5 remain named consumers; and E2 tagged-census evidence proves the writer/reader contract. The fold changes predicates, not the target entity, row identity, census format, or consumer graph.

Acceptance/evidence check: at the unchanged implementation head `720bdd683c1c850be1077c3381c3cc8870233db0`, the exact tagged command `go test -tags seam -json -count=1 ./test/seam/ -run 'TestCT_' | python3 test/seam/census.py` freshly reproduced all G01–G20 GREEN, all 44 gate rows RED, and `SUMMARY GREEN=20 RED=44 TOTAL=64` with census exit zero. This proves the pre-fold baseline at E2; it does not pre-claim the post-fold result. PLAN r9 correctly requires a fresh post-fold derivation, treats D04 and every other color delta as a routed finding, and forbids a silent rewrite.

Remaining gates: this verdict authorizes no test edit. `s16a.planner` must mechanically re-run SCOPE_DIFF against the locked r9 scope and submit a successor IMPL relay parented to this review, carrying a live bare token re-bound to `s16a-build-9`. After the fold commit and refreshed census, m-9 and m-10 re-confirm only their changed predicates/sublegs; WP2 remains held until master's joined carriage returns. Merge remains a separate operator gate.

No implementation action: no source/test byte, branch state, dependency, stage, commit, push, PR, merge, provider, credential, store, or runtime state was created or changed by this review.

Tests / verification: E2 exact tagged baseline census; E1 full PLAN r9, ruling, and four PM returns read; exact SHA-256 checks; row-set reconciliation; direct runtime/journal-order inspection for A02 locus; header/INDEX parent and admission checks; current worktree-head check; engine version/status evidence; exact-file relay lint clean. Root-wide relay/index lint remains red only on inherited historical/projection defects and is not claimed clean.

Next requested action: `s16a.planner` runs the r9 SCOPE_DIFF and, only if every row is in, submits the successor IMPL relay with `PARENT_DISPATCH_ID: s16a-build-plan-review-4`.

ESCALATION_SCAN:
- authz/tenant/RLS/permissions/secrets: yes — the future battery probes authority and credential-bound contracts, but this act is review-only and changes no such byte
- migration/backfill/destructive-write/canonical-data-repair: no — neither this review nor the bounded test-only fold performs these actions
- money/inventory/orders/planning/accounting/trust-critical-state: yes — the reviewed seams are trust-critical, so r8's existing medium pair ceremony is preserved without a downgrade
- AI-or-automation-acts-downstream: yes — future worker/provider behavior is under test; this verdict itself performs no downstream action
- worker/scheduler/queue/retry/async-side-effect: yes — the contract battery covers worker and supervised-runtime seams; the fold remains test-only
- cross-repo/service-contract/generated-schema/shared-API-event: yes — the battery spans registered app IPC, connector, worker, and conductor contracts inside the governed repo
- user-visible-control-with-materializer/downstream-consumer: no — this fold creates no user-visible control or materializer
- test-runtime-role-mismatch: no — build-tag isolation already separates the intentionally RED instrument from the untagged product suite; the fold does not change that execution role
- broad-scope-expansion/ambiguous-product-semantics/residual-risk/live-verify-skip: no — master adopted the exact PM returns, the edit fence does not widen, and E2 is the required fold evidence
ESCALATION_SCAN_RESULT: trigger-present
PRE_SCAN_PRESSURE: none
OPERATOR_WAIVER:
- status: none
- scan_presented_to_operator: no — no ceremony downgrade or risk acceptance is requested
- operator_reply_after_scan: none
- valid_waiver: no
WAIVED_RISK_ACCEPTANCE: none — no downgrade or risk acceptance is requested

ACTIONS_GIT_REF: engine-lane governance act only — this PLAN-REVIEW is drafted under `.engine/drafts/s16a.implementer/` for submission through `relay submit`; the daemon renders the relay and INDEX row; no source/test or branch-state action claimed
FINAL_GIT_STATUS_SHORT:
 M frank/.relays/s16a/INDEX.md
 M master/relays/CHECKPOINTS.md
?? frank/.relays/s16a/s16a-build/PLAN-planner-20260824-213626.md
