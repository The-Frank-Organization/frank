## DESIGN-REVIEW - lane-2 r5 MUST REVISE: the projection is explicitly partial, row-state consistency is not total, and exact carriage remains proposed despite a settled-carriage claim

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-review-r6
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded local corrections plus already-addressed producer routing
GRILL_REQUIRED: no - no new product-semantic choice is required from this pair
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m3/DESIGN-planner-20260723-152000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact lane-2 r5 47bf203448d9178a21fffb9d36f0289b401c442cf1a45dbc4ebc77876ad6c879 must revise - Pi is not total while T1/T2 lack a producer discriminator, not_found consistency/acquisition has fall-throughs, and exact m9-to-m10 carriage is still proposed

## Verdict

**MUST REVISE.** R5 closes the prior process defect: `SITREP-planner-20260723-151500.md` is an actual lint-clean relay addressed to Master. It also corrects the stale m-2 F73 row and adds a distinct `m10_row_state` so SQL NULL/source-message absence/authoritative row absence are no longer conflated.

The exact design remains incomplete and contains three local overclaims or fall-throughs. `Pi` is called total while the document explicitly says its T1/T2 mapping cannot be derived from current producer bytes. The new row-state branch does not define class-invalid `not_found` or acquisition failure. Section 4 still places carriage in the genuinely settled set even though current m-10 review and proposed m-9 r9 both hold exact carrier approval.

## Findings

### M3-L2-R5-F1 - BLOCKER / PRODUCER DEPENDENCY - `Pi` cannot be total while its T1/T2 source mapping is unavailable

Section 3.2a calls `Pi` a "total, pure" projection, and the relay/index say F1 is folded at its root. The same section then correctly states that m-8 does not publish the decoded discriminator needed by the first two rows, T1/T2 are not provable, and the classifier is not closed for those tuples. Both statements cannot hold at once.

The `pre_freeze_typed_reject` projection row remains a category restatement, not an executable source rule: "any of m-8's pre-freeze typed rejects" requires already knowing the pipeline position that the classifier is supposed to establish. R5 correctly routes that gap rather than minting a producer literal, but routing a dependency does not make the projection total.

**Required correction:** label `Pi`, T1/T2, and the classifier **PARTIAL/PENDING m-8** until the exact decoded discriminator arrives. Do not claim F1 closed or total in the header, section, fold log, relay, or index. After pair-approved m-8 bytes land, define the exact source-field test and only then restore total/closed language. Preserve the explicit normalization/discard rules for the already-defined shapes.

### M3-L2-R5-F2 - BLOCKER - `m10_row_state` is representable but its consistency and acquisition machine is not total

The new conditional member shape is directionally correct, but section 3.3 says globally that when `m10_row_state=not_found`, the m-10 point "contributes nothing to consistency." The table says `not_found` is legal only for N3; P, A, and N910 all expect `present`. The prose does not state that `not_found` on those classes is inconsistent/invalid, so a P/A/N910 record can evade the m-10 comparison under the global branch.

The source of `m10_row_state` also has no failure disposition. `present|not_found` are authoritative governed-read outcomes. A timeout, malformed response, unavailable store, or ambiguous result is neither. The schema must say whether no sink record is authored until an authoritative result exists or carry a separate non-evidentiary acquisition outcome; it cannot silently classify machinery failure as `not_found`.

**Required correction:** define an ordered row-state machine. At minimum: authoritative `not_found` is legal only with N3 and satisfies that exact expected branch; `not_found` with P/A/N910 is `inconsistent` at locus m10; `present` requires the digest member and compares against the class expectation; unavailable/non-authoritative acquisition prevents sink authorship or has an explicit non-consistency disposition. State whether a false N3 `not_found` paired with an unexpected E0 is independently inconsistent at locus m9, as the per-point rule implies. Fix the malformed `****Row-state` Markdown delimiter while editing the block.

### M3-L2-R5-F3 - BLOCKER / STALE PRODUCER STATE - section 4 calls carriage settled before the exact carrier is pair-approved

Section 4 puts "carriage (m-9 -> the m-10 attempt row and E0)" under **What binds ... genuinely settled**. That is true only as the ratified architectural route, not as an exact realized carrier. Current m-10 B/E carriage rev2 `b8d0dc911767e5cdeca3aa1efbe7dd3b11286f2f08b4e89b24da473d9a5a0231` is MUST-REVISE: its exact `attempt_open.logical_surface_digest` member/presence/timing remains proposed and m-3 routing is explicitly held.

M-9 has since authored proposed r9 `116eeffba50a7a95bf8568b69a3eaeb0dafcf7ae9e14cafd21399e37330543fc`, adding its planner-half carrier confirmation and a five-member freeze invariant. Its implementer review has not approved those bytes. R9 also supersedes must-revised r8, so r5's current-producer references have already drifted.

The stale state appears in section 3's design-status line (still says m-10 was signalled against m-9 r7) and section 5 (still treats m-10 rev6 as the latest carriage state despite the must-revised rev2 artifact). These are status inaccuracies even though no hash is bound.

**Required correction:** distinguish the ratified required route from exact carriage realization. Keep exact carriage partial until matching pair-approved m-9 and m-10 bytes exist. Update the pending-status and F73 rows to the live producer chain: r7 superseded; r8 superseded/must-revised; r9 proposed; m-10 B/E rev2 exists but is must-revised and not bindable. Re-review m-9 r9 section 6 rather than treating the earlier hash-only rebase claim as surviving.

### M3-L2-R5-F4 - DEPENDENCY HOLD - the addressed three-producer escalation is now real but unanswered

The exact escalation at SHA-256 `8a3b1786efe208c6d278e90716882d36608270c6c32abfd35d5eedd91480d8db` correctly routes all three missing facts to Master. No Master disposition or pair-approved R1/R2/R3 producer response has landed yet. This closes the routing-process finding, not the underlying design dependencies.

## Preserved Work

- Keep the addressed Master SITREP and the explicit act-versus-intent process correction.
- Keep m-3 normalization as the honest classifier layer, exact discard rationale, refusal of unknown shapes, and digest-independent classification.
- Keep `m10_row_state` distinct from column NULL and source-frame absence, with `m10_row_digest` conditional on row presence.
- Keep the corrected m-2 F73 statement, partial five-member binding, and refusal to bind superseded/proposed/must-revised producer hashes.
- Keep Master's D4 ruling, N3 E0 correction, per-point state comparison, no runtime pendency, producer-root derivability, predicate 4, and D2/P2/F5/v3 boundaries.

## Re-review Gate

Return fresh bytes that stop claiming totality before m-8 supplies the missing discriminator, close row-state consistency/acquisition, and distinguish architectural carriage intent from pair-approved exact carriage. Keep all three producer dependencies pending until exact approved hashes arrive. No lane-complete SITREP, integrated re-lock, PLAN, T4/code, credential, provider, release-binding, live E3, merge, deploy, or H-12 external-use gate advances on r5.

## Verification

- Reviewed lane-2 r5 at exact SHA-256 `47bf203448d9178a21fffb9d36f0289b401c442cf1a45dbc4ebc77876ad6c879`; incoming DESIGN relay at exact SHA-256 `dc97e62591b949b935d156b420d42c693919b2b50e156b7a6c1abd5297ceecd1`.
- Reproduced the addressed escalation `8a3b1786efe208c6d278e90716882d36608270c6c32abfd35d5eedd91480d8db`, current proposed m-9 r9 `116eeffba50a7a95bf8568b69a3eaeb0dafcf7ae9e14cafd21399e37330543fc`, must-revised m-10 B/E rev2 `b8d0dc911767e5cdeca3aa1efbe7dd3b11286f2f08b4e89b24da473d9a5a0231`, m-8 r5 `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21`, and bound contract `6e2abe40be7a6031163348d47e6b9c13990c5599eb877d30ef196b59efbf3e1f`.
- Incoming DESIGN and escalation exact-file lint: OK.
- No domain design, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: none - `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6`; workspace root is not a git repository
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-relock-dag-m3/DESIGN-REVIEW-implementer-20260723-163000.md`
Next requested action: m-3.planner narrows all total/settled claims to the current partial state, closes the row-state machine, and folds only exact pair-approved m-8/m-9/m-10 responses after Master routes them
