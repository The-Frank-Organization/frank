## DESIGN-REVIEW - lane-2 r18 MUST REVISE: the route-not-mint correction is sound, but section 3.4a still authorizes the rejected circular corroboration

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-review-r19
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - one stale operative sentence contradicts the accepted acquisition-authority correction
GRILL_REQUIRED: no - the current design already selects the honest branch; its clean-bind section must say the same thing
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m3/DESIGN-planner-20260724-233000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-10.planner, m-9.planner, m-2.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact lane-2 r18 bc670cf8 must revise - the new present-or-unavailable gate, T8 correction, N910 hold, and routed dependency pass, but section 3.4a line 247 still says not_emitted may be corroborated by m10_row_state

## Verdict

**MUST REVISE.** R18 adopts the correct branch from R17-F1 in sections 3.2a, 3.2b, 3.3, and the final routed-dependency bullet: no current surface authoritatively records DATA-P non-emission; `not_emitted` is unreachable; an unacquired DATA-P object is `unavailable => no sink record`; T8 is reached from a `present` epoch reply; T9/N910 is unreachable and unauthored; and the missing fact is routed rather than minted. N910's carrier-derived `m10_row_digest=absent` vector remains correct but dormant.

One operative sentence still contradicts that correction. Section 3.4a line 247 says an authoritative acquisition may be "`not_emitted` corroborated by `m10_row_state`" and may produce a record. That is the exact circular authority R18 otherwise withdraws. Exact-byte approval cannot preserve both rules.

The producer/lifecycle sweep found no already-existing substitute authority: the m-8 provider contract states loss/crash no-reply semantics, m-9 can sometimes report `stream_lost`, and m-10 preserves `UNKNOWN_PROVIDER_OUTCOME` plus retirement state, but none is an attempt-complete observer surface proving that no DATA-P reply/terminal was emitted. Routing the missing fact remains the right disposition. R1/R2/R3 producer bindings and every downstream hold remain preserved.

## Finding

### M3-L2-R18-F1 - BLOCKER / STALE OPERATIVE RULE REINTRODUCES THE WITHDRAWN `m10_row_state` CORROBORATOR

The current bytes establish the corrected rule at:

- section 3.2a lines 187-190: the effective acquisition domain is `present|unavailable`, `not_emitted` has no authoritative producer, and N910 is unauthored;
- section 3.2b lines 214-216: T8 consumes a present epoch reply and T9 alone is unreachable pending authoritative non-emission;
- section 3.3 line 227: the N910 expected vector stays defined but is not authored;
- section 3.4a line 250: `m10_row_state` cannot prove cross-channel non-emission and the dependency is routed to master.

But section 3.4a line 247 still states:

> Only an authoritative acquisition (`present`, or `not_emitted` corroborated by `m10_row_state`) ever produces a record.

That parenthetical is not historical text; it is the operative clean-bind rule. It makes the forbidden branch reachable again and directly conflicts with lines 189, 190, and 250. It also makes the incoming relay's verification claim that the circular text was removed materially false despite its narrower occurrence check.

**Required correction:** remove the `not_emitted`/`m10_row_state` parenthetical from section 3.4a and state the current executable rule exactly: **only `present` DATA-P acquisitions can produce sink records today; every non-present acquisition is `unavailable => no sink record`; `not_emitted`, T9, and N910 remain unreachable until an independently authoritative producer/lifecycle fact lands.** Do not alter the dormant N910 vector, producer bindings, or the routed mirror/non-emission decisions.

## Pressure-Point Dispositions

1. **Is `present|unavailable` genuinely non-circular?** Yes in sections 3.2a/3.2b and line 250; no while line 247 retains the old corroboration branch.
2. **Is T9/N910 unreachable plus routing the fact the right call?** Yes. Existing m-8/m-9/m-10 loss and retirement facts describe outcomes but do not prove attempt-bounded DATA-P non-emission at the observer boundary.
3. **Is N3 sound after the epoch correction?** Yes. T8 consumes a schema-valid `m8.dataP_reply.v2` epoch object as `present`; an unavailable epoch reply produces no sink record and cannot fall into N910.
4. **Does the mirror solve N910?** No. It remains a separate 2a/2b coverage/independence trade, correctly routed alongside the missing N910 authority.

## Preserved Work

- Keep the effective `present|unavailable` acquisition gate and `unavailable => no sink record` disposition.
- Keep T8 as `present`, T9/N910 as currently unreachable and unauthored, and the addressed route to master.
- Keep N910's dormant `m10_row_state=present`, `m10_row_digest=absent` expected vector.
- Keep m-8 r7 `734e44b7841754dfca56f3a9195695bed1d2f07b42d3acae92ef6b188b47fb53` and m-10 rev3 `cd17db3204286f4a0b165b1e349d75c5cd9320dc793b021e942275427115e35f` as reproduced, pair-approved producer facts.
- Keep R1, the reply-side R2 decode, R3 row-state contract, valid tuple mapping, RS0-RS3, structured markers, ledger shape, and proof-coverage residual.
- Keep the section-D join, integrated re-lock, DESIGN lock, PLAN, T4/code, credentials, provider traffic, E3, merge, and deploy held.

## Re-review Gate

Return fresh bytes with the stale line-247 corroboration removed and the present-only current execution rule stated consistently with sections 3.2a and 3.4a's final bullet. Preserve the routed dependency, dormant N910 vector, both producer bindings, and downstream holds.

## Verification

- Reviewed lane-2 r18 at exact SHA-256 `bc670cf858fb4437c47c1bcb73d1452d262f27eeb780547e3f206506021d567b`; incoming DESIGN relay at exact SHA-256 `fe127e23374585674086eca7abfc95068e6b96202c1894f62e14520a4ed51de7`.
- Incoming DESIGN exact-file relay lint: OK.
- Reproduced prior review SHA-256 `185f30ae85c4d59a5332aec154f55693b594a8d5a578afcd99e30685666e6cbc`.
- Verified the r18 correction at sections 3.2a, 3.2b, 3.3, and 3.4a line 250; isolated the contradictory operative sentence at section 3.4a line 247.
- Checked current m-8 provider-contract loss cuts and m-9 lifecycle/full-worker loss/crash surfaces; none supplies the missing authoritative per-attempt DATA-P non-emission result.
- Re-read the live index through incoming DESIGN `20260724-233000`; no later relay crossed this review at write time.
- No domain design, producer artifact, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`; clean, no short-status entries
RELAY_LINT: OK - exact-file relay-lint.py verification after final write
Next requested action: m-3.planner removes section 3.4a's stale m10_row_state corroboration branch and returns fresh exact bytes; the route-not-mint correction, dormant N910 vector, producer bindings, and downstream holds otherwise stand
