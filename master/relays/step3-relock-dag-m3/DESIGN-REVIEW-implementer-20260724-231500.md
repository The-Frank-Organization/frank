## DESIGN-REVIEW - lane-2 r17 MUST REVISE: N910 is closed, but DATA-P `not_emitted` still has no independent authoritative source

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-review-r18
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded acquisition-authority correction
GRILL_REQUIRED: no - the existing no-vantage rule selects honest unavailability unless a positive authoritative source is named
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m3/DESIGN-planner-20260724-221500.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-10.planner, m-9.planner, m-2.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact lane-2 r17 b1f05d34 must revise - N910 now has the correct carrier-derived absent expectation, but m10_row_state cannot authorize no emission on DATA-P and the epoch cut is incorrectly treated as a not_emitted case

## Verdict

**MUST REVISE.** R17 closes R16-F2: N910 now has the exact carriage-derived expectation `m10_row_state=present` plus `m10_row_digest=absent`, and the approved no-carrier route supports it. The pre-normalization `present|not_emitted|unavailable` split is also the right form, and `unavailable => no sink record` would prevent semantic `none` from absorbing an acquisition failure.

One blocker remains: the bytes do not identify a source that can authoritatively produce `not_emitted`. `m10_row_state` proves only whether an m-10 row exists; it cannot prove whether m-8 emitted a DATA-P object on a different channel. The current corroboration rule therefore does not make no-emission authoritative and reintroduces the downstream-fact circularity the acquisition gate is meant to remove.

R2/R3 producer hashes remain reproduced and may stay bound as producer facts. The m-3 clean-bind and lane-complete claims do not pass until the authority source is exact. Nothing downstream moves.

## Finding

### M3-L2-R17-F1 - BLOCKER / `not_emitted` HAS NO AUTHORITATIVE PRODUCER - row existence cannot prove cross-channel absence

Section 3.2a line 188 defines `not_emitted` as a completed read that authoritatively shows no DATA-P object, but then makes that result authoritative only when corroborated by `m10_row_state`. The named corroborator does not carry the claimed fact:

- `m10_row_state=present` means only that `provider_attempts` committed. It says nothing about whether m-8 emitted a DATA-P reply or terminal.
- `m10_row_state=not_found` means no row committed. It likewise says nothing about DATA-P emission.
- The phrase `present with no DATA-P and no CTRL-C` includes **no DATA-P** as a premise; the row state does not establish it. Using the observer's missing DATA-P capture to fill that premise is exactly the unavailable-versus-not-emitted ambiguity under review.

The epoch sentence exposes the contradiction. Row 3/T8 has a real `m8.dataP_reply.v2` epoch reply, so its DATA-P acquisition result must be `present`, never `not_emitted`. If that reply is unavailable to the observer, row state cannot distinguish the actual epoch cut from N910. Line 190 likewise says the authoritative-no-message `none` is consumed by T8/T9, but T8's `data_p_reply_kind` is `stale_epoch|epoch_ahead`; only T9 consumes DATA-P `none`.

The closed enum is therefore syntactically present but its positive-negative branch is not executable at a named authority boundary.

**Required correction:** name the exact DATA-P acquisition surface/recorder and the property that makes a completed read an authoritative per-attempt negative. The acceptable shape is a positive protocol result from a complete attempt-bounded channel capture/store, independent of the m-10 row being checked. Then:

- `present(valid reply or terminal)` feeds T1-T8 as applicable;
- authoritative DATA-P `not_emitted` feeds T9 only;
- missing, incomplete, gapped, undecodable, ambiguous, or non-authoritative acquisition is `unavailable => no sink record`;
- `m10_row_state` may be compared after classification but may not manufacture the classifier's no-message input.

If no current surface can provide an authoritative negative, say so: make `not_emitted` unreachable in the current design and author no N910 sink record until a producer/lifecycle fact is routed. The r7 CTRL-C mirror does not solve N910; it only recovers 2a/2b classification when the reply is unavailable.

## Pressure-Point Dispositions

1. **Does the new gate prevent `unavailable` from producing a record?** Yes by rule. The remaining problem is deciding when an observed absence is legitimately `not_emitted` rather than `unavailable`.
2. **Do T1-T9 align with the acquisition domain?** T1-T8 require `present` DATA-P objects/terminals. T9 alone requires authoritative `not_emitted`. R17's statement that T8 consumes DATA-P `none` is incorrect.
3. **Does N910 now have a closed m-10 expectation?** Yes. `absent` follows from no CTRL-C carrier and verbatim-only m-10 storage; R16-F2 is closed.
4. **Should the mirror route now?** That is a separate coverage decision for 2a/2b and may remain with master after the current correctness contract is closed. It cannot substitute for authoritative DATA-P absence on N910.
5. **Has the mirror trade been routed?** Not yet. Master is CC on the incoming relay, which is context only. Include the trade in the eventual addressed lane return; this is a routing note, not the current design blocker.

## Preserved Work

- Keep the closed acquisition-result form and `unavailable => no sink record` disposition.
- Keep N910's exact `m10_row_digest=absent` expectation and the approved carrier reasoning.
- Keep m-8 r7 `734e44b7841754dfca56f3a9195695bed1d2f07b42d3acae92ef6b188b47fb53` and m-10 rev3 `cd17db3204286f4a0b165b1e349d75c5cd9320dc793b021e942275427115e35f` as reproduced, pair-approved producer facts.
- Keep the reply-side R2 decode, R3 row-state contract, valid tuple mapping, RS0-RS3, R1 binding, structured markers, ledger shape, and proof-coverage residual.
- Keep the section-D join, integrated re-lock, DESIGN lock, PLAN, T4/code, credentials, provider traffic, E3, merge, and deploy held.

## Re-review Gate

Return fresh bytes naming a real authority source for the DATA-P acquisition negative, limiting `not_emitted` to T9, and treating an unavailable epoch reply as `unavailable`, not authoritative absence. If no source exists, record that limit and route the missing producer/lifecycle fact instead of claiming N910 classification. Preserve R16-F2's exact expectation and both producer bindings.

## Verification

- Reviewed lane-2 r17 at exact SHA-256 `b1f05d34abe59bf5a867ed5a57cc458ada0803d35133a4c4033313d499054ce3`; incoming DESIGN relay at exact SHA-256 `dee92847b3618519dc394b247b07c4d8b61b4b0b3c333bc5d5449cc272a911b5`.
- Incoming DESIGN exact-file relay lint: OK.
- Reproduced m-8 r7 `734e44b7841754dfca56f3a9195695bed1d2f07b42d3acae92ef6b188b47fb53` and m-10 B/E rev3 `cd17db3204286f4a0b165b1e349d75c5cd9320dc793b021e942275427115e35f` unchanged.
- Reproduced 18 operative structured markers: R1 = 4, R2 = 6, R3 = 8; all parse and all id-to-anchor mappings match.
- Reproduced R16-F2's correction at section 3.3: N910 now requires `m10_row_digest=absent`, derived from the no-CTRL-C carrier route.
- Walked the acquisition gate against T1-T9: T1-T8 consume present replies/terminals; T9 alone consumes DATA-P no-message.
- Re-read the live index through the incoming DESIGN `20260724-221500`; no later relevant m-8/m-9/m-10 producer relay existed at review-write time.
- No domain design, producer artifact, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`; clean, no short-status entries
RELAY_LINT: OK - exact-file relay-lint.py verification after final write
Next requested action: m-3.planner names the authoritative DATA-P negative source or honestly makes N910 sink-record production unavailable and routes the missing fact; N910's m-10 expectation and both producer bindings remain preserved
