## DESIGN-REVIEW - lane-2 r2 MUST REVISE: the no-carrier precondition is not ratified, cut provenance is not injective, and the pending sink conflates message absence with absent members

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-review-r3
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded pair corrections plus Master's already-requested contract-integration and producer sequencing
GRILL_REQUIRED: no - no new product-semantic choice is required from the pair
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m3/DESIGN-planner-20260723-035000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact lane-2 r2 5b96673b2abdc7616fe9954df8bc7a669d9318b8b76b86e5e64b3a040e183b51 must revise - no-carrier precondition awaits Master, cut-row provenance is not mechanically unique, and the pending sink schema erases exact carrier-route distinctions

## Verdict

**MUST REVISE.** R2 closes R1-F2: predicate 4 now has a genuinely closed protocol-grain result domain, including completed-but-partial or invalid responses in `unavailable`. It also corrects row 3's DATA-P carrier, names exact DATA-P/CTRL-C routes, keeps all absence schema-valid, preserves the D2/P2/F5/v3 boundaries, and correctly identifies `m8_terminal_digest` as the producer-root derivability comparand. The m-10 and m-9 producer gaps are now reported honestly and escalated rather than approved into existence.

The exact bytes still cannot be approved. Section 1 locally adds a precondition to the ratified iff rule while correctly admitting that only Master can dispose it. Sections 3.1-3.3 then claim a mechanically closed sink but do not carry an injective cut-classification source, encode temporary design progress as runtime evidence, and treat `absent` and `no_message` as interchangeable despite the exact route matrix distinguishing them.

## Findings

### M3-L2-R2-F1 - BLOCKER / MASTER INTEGRATION - the no-carrier precondition is not an instantiation of the ratified total iff rule

The bound contract at exact SHA-256 `6e2abe40be7a6031163348d47e6b9c13990c5599eb877d30ef196b59efbf3e1f` says the E0 populator and E3 writer MUST include `frozen_core_digest` iff `FREEZE-REACHED(cut) = true`. R2 section 1 instead says the rule has no subject when no digest-bearing carrier exists and makes absence follow "by construction, not by the rule." For rows 9/10, whose settled m-8 input is `freeze = maybe`, that is a new rule precondition, not an evaluation of the ratified boolean.

The escalation in `SITREP-planner-20260723-034500.md` is the correct authority route, and the draft correctly refuses to self-declare the change into the ratified text. That same admission means these exact presence bytes are not yet binding.

**Required correction:** hold rows 9/10's producer-conformance consequence pending Master's exact D4 disposition. Fold the returned ruling verbatim. If the ruling changes the bound contract rather than merely supplying its settled cut evaluation, route the change through the governing amendment/ratification path; m-3 cannot create the precondition locally.

### M3-L2-R2-F2 - BLOCKER - `cut_source` plus open `cut_source_value` does not mechanically reconstruct a unique `cut_row`

Section 3.2 says an independent reader can derive `cut_row`, but the record carries only `cut_source = m8_disposition|m10_terminal_state` plus an unconstrained string. That mapping is neither closed nor injective:

- Rows 2a and 2b share the CTRL-C disposition `rejected_local(internal_integrity_fault)` while requiring opposite digest presence. The prose allows the named m-8 carrier to be either DATA-P or CTRL-C, but the record does not identify which carrier/kind supplied the string. Presence cannot be used to choose 2a versus 2b without making the expected-presence derivation circular.
- The exact m-10 terminal-state values for rows 3/9/10 are not enumerated or mapped to row ids, and their producer contract is still pending. An open string cannot prove those rows uniquely classifiable.

**Required correction:** define a closed, attempt-bound classification tuple and a complete mapping from that tuple to each cut row. The tuple must identify the source carrier and exact typed fact needed to distinguish 2a from 2b without consulting any sink digest member. Enumerate the eventual m-10 terminal-state inputs and prove whether they uniquely distinguish 3/9/10 after the producer fold exists; if they do not, route the missing discriminant to the owning producer rather than self-asserting `cut_row`.

### M3-L2-R2-F3 - BLOCKER - the sink schema encodes integration status and its consistency machine erases the exact carrier route

`pending_producer` and `indeterminate_pending` describe the current document/DAG state, not evidence from an attempt. Baking them into `m3.b_sink.v1` would permanently weaken the runtime schema after the m-10 producer lands. Before that producer contract exists, no runtime sink record can validly claim to have observed its carriage point; the section should remain pending as design status, not model the pending work as an attempt value.

The remaining consistency branches are total only over a coarse aggregate. For both `no_carrier` and `absent_members`, every mixture of `absent` and `no_message` is declared consistent. That loses the exact route R2 just established. It accepts a missing expected message (`no_message` where a carrier must exist with absent members) and an invented message with no digest (`absent` where no message may exist). The three members intentionally distinguish those states, so consistency must preserve them.

**Required correction:** remove `pending_producer` and `indeterminate_pending` from the final runtime schema; keep the whole producer-dependent sink section explicitly pending until exact pair-approved m-10 bytes can be bound. Derive an expected per-carriage-point state vector from the uniquely reconstructed cut row, with each point exactly one of `<64-hex-required> | absent | no_message`, and compare all three observed members point by point. A wrong route is `inconsistent` even when no hex digest appears. Preserve derivability as the independent comparison of the observer derivation to the producer-root m-8 value.

### M3-L2-R2-F4 - DEPENDENCY HOLD - the two producer folds remain unresolved, as R2 now states accurately

M-10 rev6 `29a123fe8f153127ccfe03559850c8e5ce4c488671402cebf8325e3cafbcd7ae` still does not design the B/E rows. M-2 rev2 `c3a8cd61bcbe33a3ede847b29d7416563000c66f949fa2950ea0af305edbbd2c` is pair-approved, but m-9 r5 `c0ff74f5589e0dcbd066fe8a7e04ee7fafd53a9a20befa6453084110feccf70b` still parks consuming that component. R2 section 4 and its `034500` escalation now state that chain correctly and do not claim lane completion.

This is a hold, not permission to synthesize either producer contract in m-3. Exact approval remains unavailable until Master routes the producer work, those pairs approve fresh bytes, and m-3 binds those hashes.

## Preserved Work

- Keep the exact ten-row DATA-P/CTRL-C carrier facts, especially row 3's existing epoch reply and absent CTRL-C result.
- Keep predicate 4's closed `committed|not_found|unavailable` protocol-grain domain and five exhaustive branches.
- Keep `m8_terminal_digest` as the canonical producer-root comparand. Consistency diagnoses copy/carriage disagreement; derivability independently answers whether the observer reproduced the producer-root digest.
- Keep ratified v2 byte parsing, presence-read-not-inferred, P2a/P2b separation, D2's gating split, F5's record-only authority ceiling, and exclusion of `model_surface_digest`/the E join to v3.
- Keep the honest partial logical binding and the Master -> m-9 -> m-3 sequencing. M-3 must not fold m-2-owned composition directly into m-9's producer contract.

## Re-review Gate

Return fresh bytes only after F2/F3 are mechanically closed. For F1/F4, bind Master's exact disposition and the fresh pair-approved m-9/m-10 producer hashes; if those responses have not landed, leave the affected design sections pending and do not represent temporary DAG state in the runtime schema. No lane-complete SITREP, integrated re-lock, PLAN, T4/code, credential, provider, release-binding, live E3, merge, deploy, or H-12 external-use gate advances on r2.

## Verification

- Reviewed lane-2 r2 at exact SHA-256 `5b96673b2abdc7616fe9954df8bc7a669d9318b8b76b86e5e64b3a040e183b51`; incoming DESIGN relay at exact SHA-256 `fbc6c2c2a292699556839c40cfa23690cc84f259bd98b326d39d495bbe9f2426`.
- Reproduced ratified amendment `9e874df84015261d77e9c353528e821fd8491489388c010fa621fe630432b351`, bound contract `6e2abe40be7a6031163348d47e6b9c13990c5599eb877d30ef196b59efbf3e1f`, m-8 r5 `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21`, m-9 r5 `c0ff74f5589e0dcbd066fe8a7e04ee7fafd53a9a20befa6453084110feccf70b`, m-10 rev6 `29a123fe8f153127ccfe03559850c8e5ce4c488671402cebf8325e3cafbcd7ae`, and m-2 rev2 `c3a8cd61bcbe33a3ede847b29d7416563000c66f949fa2950ea0af305edbbd2c`.
- Incoming DESIGN exact-file lint: OK.
- No domain design, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: none - `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6`; workspace root is not a git repository
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-relock-dag-m3/DESIGN-REVIEW-implementer-20260723-050000.md`
Next requested action: m-3.planner closes the local cut-provenance and per-carriage consistency defects, keeps producer-pending state out of the runtime schema, and folds only the exact Master/producer dispositions once they exist
