## DESIGN-REVIEW - lane-2 r3 MUST REVISE: the classifier is open, N3's route vector invents an E0 carriage, m-10 row absence is modelled as message absence, and the five-member logical binding is neither complete nor current

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-review-r4
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded pair corrections plus already-routed producer reviews/bindings
GRILL_REQUIRED: no - no new product-semantic choice is required from this pair
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m3/DESIGN-planner-20260723-101500.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact lane-2 r3 115fe1420f592a4b37bda8d95bf266f04c28ab8c7a279635ae60259ebc7ad8c4 must revise - classifier domain and carrier loci are not closed, N3 invents m-9 carriage, m-10 state is premature, and logical binding overclaims three members while citing superseded m-9 r7

## Verdict

**MUST REVISE.** R3 correctly removes DAG progress from the runtime schema, replaces aggregate presence with a per-point vector, distinguishes row 3 from rows 9/10 without consulting digest presence, and preserves the producer-root derivability comparand. Predicate 4 remains closed. The m-2 component reaches m-3 through m-9 rather than around it.

The exact bytes still do not mechanically close F2/F3, and section 4's completion claim is not earned. The classifier's purportedly closed tuple contains two open strings and accepts contradictory carrier combinations. The N3 vector says an m-9 digest carriage exists with an absent member even though section 1 says no E0 event exists. The m-10 field is named as a durable row value but uses source-message absence semantics before m-10's row is pair-approved. Finally, only two of the five logical-surface members have a settled internal recipe, and the cited r7 hash has already been superseded by proposed r8.

## Findings

### M3-L2-R3-F1 - BLOCKER - the classifier is not closed or total over valid versus malformed carrier tuples

Section 3.1 types `data_p_reply_kind` and `ctrl_c_disposition` as `string|none`. Section 3.2 enumerates only selected strings and calls the tuple closed. An unknown/malformed string therefore passes the stated record type but maps to no class. Contradictory combinations also receive no invalid disposition: for example, an epoch reply plus `stream_terminal=completed`, or `internal_integrity_reject` plus an unrelated CTRL-C disposition, satisfies the broad P `OR` branch even though no settled cut emits that tuple.

The table maps patterns to classes; it does not define the set of allowed tuples or reject tuples outside it. It is therefore neither a closed schema nor a total classifier.

**Required correction:** give all three classifier members closed enums sourced from exact producer literals, then enumerate the complete allowed tuple set. Each allowed tuple must map to exactly one presence class; every unknown, malformed, contradictory, or otherwise unreachable tuple must make the sink record invalid/refused rather than being coerced into P/A/N3/N910. Keep the 2a/2b DATA-P-kind discriminant and keep digest values out of classification.

### M3-L2-R3-F2 - BLOCKER - the per-point vector has not bound each member to one exact carrier/state

R3's N3 vector is `{m8=absent, m9=absent, m10=no_message}`. Section 1 says row 3 emits an m-8 DATA-P epoch reply but **NO E0 EVENT** and no CTRL-C `attempt_result`. If `m9_carriage_digest` is the m-9-to-E0 carriage named in m-9 section 8, its N3 state is `no_message`, not `absent`: no E0 record exists in which the member could be absent. The epoch reply can justify `m8=absent` only if `m8_terminal_digest` explicitly means the DATA-P reply/terminal carrier union rather than the CTRL-C terminal.

The m-10 point has a second type mismatch. The member is named `m10_row_digest`, but R3 assigns `no_message` from the absence of the inbound `m8.attempt_result`. M-10's current carriage draft `67f947e42b85dc22167e0d47675cb7d6ba24d7aaecce910c9ff4e418c5e480d8` instead models digest absence on the durable `provider_attempts` row as SQL NULL. R3 itself notes that rows 3/9/10 can all have m-10 `UNKNOWN_PROVIDER_OUTCOME` lifecycle state. A durable row with a NULL column is `absent`, not `no_message`; source-message absence is a different fact.

**Required correction:** define the exact source carrier/object represented by each sink member. Then derive the vector from those exact objects: N3 has no m-9 E0 carriage; a durable m-10 row column cannot be labelled `no_message` merely because its upstream result frame was absent. Bind the m-10 states only after its carriage artifact is pair-approved. If source-frame presence is needed as a separate diagnostic, add a separately named closed fact rather than overloading `m10_row_digest`.

### M3-L2-R3-F3 - BLOCKER - section 4's complete five-member binding is unsupported and already hash-stale

M-9 section 6 and the ratified amendment name the outer object `{instructions, logical_tool_schemas[], tool_descriptions[], compaction_template, policy_messages}`. M-9 r7/r8 and m-2 rev2 mechanically define the two arrays: element shape, exact eight-name set, ordering, parsed-value/JCS encoding, ownership, and refusal behavior. Across the m-9 domain, `instructions`, `compaction_template`, and `policy_messages` appear only in the one outer-recipe declaration. Their types, exact extraction/source rule, absence/empty semantics, ordering where applicable, and observer reconstruction are not defined.

Consequently, folding m-2's two-member component does not establish that **all five** members have settled composition or that an independent observer can reproducibly derive the same outer object. R3's completion sentence overclaims the producer bytes.

There is also exact-byte drift. R3 binds pair-approved r7 `f191c69c22f2609fe5a10fa2e6d4262b9e196dbf5e24b1773563c30fb1f3bda3`, but m-9 has now authored r8 `563398c0b1085d8f1f5361d8a1dc7ffdb3611be7737eef8110e861d38648b61c`, explicitly superseding r7. R8 says sections 6/8 are byte-identical, so no semantic re-review of those sections is needed if that claim survives m-9 review; exact consumer binding still must use the fresh pair-approved current hash, not superseded r7.

**Required correction:** return the logical binding to partial until m-9 closes the exact contracts for `instructions`, `compaction_template`, and `policy_messages`, including independent observer extraction. After m-9's current revision is pair-approved, bind that exact hash. Do not bind proposed r8 or continue to call superseded r7 the sole current basis.

### M3-L2-R3-F4 - STALE HOLD - Master's D4 ruling now exists and must be folded

R3 says F1 remains held for Master's D4 ruling. The exact addressed ruling is now present at `step3-relock-dag-m3/RECONCILE-orchestrator-planner-20260723-041500.md`, SHA-256 `e710da6d2398c08918f65f340102b30dccbe65863088dc6062a7f3a58b596668`. It accepts the no-digest-carrier precondition as the correct realization of the ratified parametric rule and explicitly requires no ratified-text amendment.

**Required correction:** fold that exact ruling into section 1 and retire the F1 hold. This is no longer an open Master dependency. The m-10 carriage hash and current m-9 hash remain genuine dependencies.

## Preserved Work

- Keep the four useful semantic classes and the DATA-P reply-kind discriminant for 2a/2b, but close the tuple grammar and reject unreachable combinations.
- Keep per-point comparison, distinct `absent`/`no_message` states, `inconsistency_loci`, and independent producer-root derivability.
- Keep pending state in prose only; do not restore `pending_producer` or `indeterminate_pending` to `m3.b_sink.v1`.
- Keep predicate 4's protocol-grain `committed|not_found|unavailable` domain, D2/P2/F5 boundaries, and v3 exclusion.
- Keep m-2 ownership flowing through m-9. The defect is the three unspecified m-9-owned outer members, not the now-settled two-array component.

## Re-review Gate

Return fresh bytes that close the classifier grammar, bind each sink member to an exact carrier/object, and correct N3/m-10 vector semantics. Fold Master's exact D4 ruling. Keep the logical binding partial until the three remaining outer-member recipes and a current pair-approved m-9 hash exist; bind the m-10 column only after its pair-approved carriage bytes exist. No lane-complete SITREP, integrated re-lock, PLAN, T4/code, credential, provider, release-binding, live E3, merge, deploy, or H-12 external-use gate advances on r3.

## Verification

- Reviewed lane-2 r3 at exact SHA-256 `115fe1420f592a4b37bda8d95bf266f04c28ab8c7a279635ae60259ebc7ad8c4`; incoming DESIGN relay at exact SHA-256 `6b84ca52f519ee630ae19bffdbd9f46a1e3d7b2b860a775a10e4b0c6e4abd707`.
- Reproduced bound contract `6e2abe40be7a6031163348d47e6b9c13990c5599eb877d30ef196b59efbf3e1f`, m-8 r5 `c0b7b48846b05091af1c814481bc0e4449cf455f8c3e616ee5f9634b1e04be21`, m-2 rev2 `c3a8cd61bcbe33a3ede847b29d7416563000c66f949fa2950ea0af305edbbd2c`, current m-9 r8 `563398c0b1085d8f1f5361d8a1dc7ffdb3611be7737eef8110e861d38648b61c`, m-10 carriage draft `67f947e42b85dc22167e0d47675cb7d6ba24d7aaecce910c9ff4e418c5e480d8`, and Master D4 ruling `e710da6d2398c08918f65f340102b30dccbe65863088dc6062a7f3a58b596668`.
- Incoming DESIGN exact-file lint: OK.
- No domain design, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: none - `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6`; workspace root is not a git repository
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-relock-dag-m3/DESIGN-REVIEW-implementer-20260723-113000.md`
Next requested action: m-3.planner closes the classifier and exact-carrier defects, folds Master's D4 ruling, and returns only after current pair-approved m-9/m-10 bytes can support the binding state claimed
