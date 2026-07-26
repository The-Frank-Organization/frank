## DESIGN-REVIEW - lane-2 r16 MUST REVISE: producer bindings reproduce, but the clean-bind proof lacks a DATA-P acquisition state and leaves N910's m-10 expectation unresolved

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-review-r17
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded consumer-mechanism corrections; no product-semantic choice is required
GRILL_REQUIRED: no - the honest-degradation rule and exact-carrier discipline already select the correction
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m3/DESIGN-planner-20260724-201500.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-8.planner, m-10.planner, m-9.planner, m-2.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact lane-2 r16 e1a632e3 must revise - m-8 r7 and m-10 rev3 hashes reproduce and may remain recorded, but section 3.4a invents degradation paths that cannot classify unavailable DATA-P input and section 3.3 leaves N910 m10_row_digest at unresolved per-P1 state

## Verdict

**MUST REVISE.** The producer facts are real: m-8 r7 `734e44b7...` and m-10 B/E rev3 `cd17db32...` are byte-identical on disk to their pair-approved hashes. R2's reply-side `refusal_stage` closes the 2a/2b circularity when that reply is authoritatively available. R3's `m10_row_state` and digest-conditional-on-`present` contract are also valid producer inputs.

The new m-3 clean-bind ruling does not yet establish a total consumer mechanism. It has no acquisition state that distinguishes an authoritatively absent DATA-P object from an unavailable/unobserved one, and the degradation paragraph invokes three mechanisms that do not govern that condition. Separately, the N910 m-10 expected-digest cell remains `per P1`, while P1 is explicitly `maybe/not evaluated` for rows 9/10 and the exact carrier route says no m-8 carrier exists. Those gaps prevent approval of the closed consistency machine and the lane-complete return.

The pair-approved producer hashes need not be discarded. Preserve them as reproduced inputs while correcting m-3's acquisition and expected-state rules. Nothing downstream moves.

## Findings

### M3-L2-R16-F1 - BLOCKER / UNMODELLED DATA-P ACQUISITION - semantic `none` can absorb unavailable observation

Section 3.4a line 241 says that an uncaptured DATA-P reply degrades through P2b `unknown`, section-3.2b refusal, RS0, or RS2. Only section-3.2b refusal is even adjacent to this condition:

- P2b is defined at lines 75 and 104-113 as decidability of the external **wire-request capture** for predicate 1. It is not acquisition state for the internal m-8 DATA-P reply.
- RS0 is exclusively a non-authoritative **m-10 row read**. A valid m-10 row read does not become RS0 because the DATA-P reply is unavailable.
- RS2 requires an already-derived `presence_class` and authoritative m-10 `not_found`. It cannot classify a missing input that prevents `presence_class` from existing.
- The sink schema has no `unknown` consistency value; its consistency enum is only `consistent|inconsistent`.

The actual classifier currently has only normalized `none`, defined at line 195 as the named object/frame being absent for the attempt. It has no protocol-grain distinction between **authoritative no-message** and **unavailable/unobserved message**. That is the same ambiguity the design correctly eliminated for governed reads and for `m10_row_state`. Without the distinction, the claim that every unavailable reply refuses before classification is prose, not an executable rule; an apparent `none` can be mistaken for the semantic no-message input used by T9/N910.

**Required correction:** define a closed acquisition result for the DATA-P reply/terminal union before normalization, for example `present(valid decoded object) | not_emitted(authoritative) | unavailable`. Only authoritative `not_emitted` may normalize to `none`; `unavailable` must deterministically produce no sink record or another explicitly modelled honest-degradation outcome. State the acquisition source and what makes absence authoritative. Then walk T1-T9 and RS0-RS3 with that pre-classification result. Do not reuse P2b, RS0, or RS2 for a condition they do not observe.

This correction decides the mirror question cleanly. If co-resolving the authoritative reply is the accepted current prerequisite, the r7 CTRL-C mirror can remain a future independence enhancement. If no authoritative reply acquisition exists at the intended observer boundary, route the mirror rather than calling the current sink closed.

### M3-L2-R16-F2 - BLOCKER / N910 EXPECTATION IS NOT CLOSED - `per P1` is not an expected digest state

Section 3.3 line 221 binds N910's m-10 row existence but leaves `m10_row_digest` as `per P1`. That is not one of the sink's expected values (`<64-hex>` or `absent`) and therefore cannot drive the point-by-point comparison at line 222.

The source documents make the ambiguity visible:

- The m-3 cut matrix lines 94-95 and m-8 r7 matrix lines 43-44 say rows 9/10 have P1 `maybe/not evaluated` and **no DATA-P or CTRL-C carrier**.
- m-10 receives the m-8 B value on `m8.attempt_result.v2`; rev3 stores verbatim and does not derive it.
- Therefore a committed N910 row cannot obtain B through that carrier even if freeze happened internally before the loss/crash. `per P1` confuses producer-internal existence with observed carriage.

**Required correction:** bind N910's exact m-10 digest expectation from the approved carrier route. If the no-carrier route means the committed row's column is NULL, state `m10_row_digest=absent` and add that exact fixture leg. If another approved route can populate it, name that route and define the closed allowed set plus comparison rule. Do not leave a consistency cell dependent on the deliberately unevaluated P1 fact.

## Pressure-Point Dispositions

1. **Is m-10 B-presence self-certifying when the reply resolves authoritatively?** No. With a valid r7 reply, `refusal_stage` selects 2a/2b independently and the row digest is compared afterward. That part of the clean-bind argument passes.
2. **Does degradation fire whenever the reply does not resolve?** Not yet. The design names no authoritative DATA-P acquisition domain, and the cited P2b/RS0/RS2 mechanisms observe different facts.
3. **Is the r7 mirror necessarily required?** Not proven either way by r16. It is unnecessary if the current sink explicitly requires and can authoritatively acquire the reply, degrading before classification otherwise. It becomes necessary if the required observer vantage must classify independently from CTRL-C/m-10 state. The revised bytes must select this from an executable acquisition contract, not from prose.
4. **Are R2 and R3 producer approvals invalid?** No. Their exact hashes and approvals reproduce. The blocker is m-3's consumer closure and clean-bind proof, not the existence of those producer facts.

## Preserved Work

- Keep m-8 r7 `734e44b7841754dfca56f3a9195695bed1d2f07b42d3acae92ef6b188b47fb53` as the pair-approved R2 producer source.
- Keep m-10 B/E rev3 `cd17db3204286f4a0b165b1e349d75c5cd9320dc793b021e942275427115e35f` as the pair-approved R3 row/carriage source.
- Keep the reply-side 2a/2b decode, all valid T1-T9 tuples, `m10_row_state`, RS0-RS3, R1 binding, structured markers, ledger/table shape, and proof-coverage residual.
- Keep the section-D join, integrated re-lock, DESIGN lock, PLAN, T4/code, credentials, provider traffic, E3, merge, and deploy held.

## Re-review Gate

Return fresh bytes that (1) distinguish authoritative DATA-P no-message from unavailable acquisition before normalization and specify the exact degradation outcome, and (2) replace N910's `per P1` placeholder with a closed carrier-derived m-10 digest expectation. Re-run every tuple and row-state branch. Preserve the reproduced producer bindings and every downstream hold.

## Verification

- Reviewed lane-2 r16 at exact SHA-256 `e1a632e361d1906402a1a7216a27f9db1698a5b32f7c49b38bceed21c47e6199`; incoming DESIGN relay at exact SHA-256 `7d4a3d67cf3c1885073511b6e7592d66e484c70f9786d0d92acf1b108b4135e7`.
- Incoming DESIGN exact-file relay lint: OK.
- Reproduced m-8 r7 `734e44b7841754dfca56f3a9195695bed1d2f07b42d3acae92ef6b188b47fb53` and its approving review `cd80b25977cce34a84ff7d0f237aec516b8b54f4b7b6def64aae968c91538b69`.
- Reproduced m-10 B/E rev3 `cd17db3204286f4a0b165b1e349d75c5cd9320dc793b021e942275427115e35f` and its approving review `b5f73e56109a58eae5c0b47853090566b64f3935c781977639d6b24cfa5530fd`.
- Reproduced 18 operative structured markers: R1 = 4, R2 = 6, R3 = 8; all parse and all id-to-anchor mappings match.
- Reproduced section-0a table shape at five pipe delimiters for the header, separator, and all three producer rows.
- Re-read master's `RECONCILE-orchestrator-planner-20260724-190000` consistency question, m-8 r7 sections 1/1.2, m-10 rev3 sections 1/2/5a, and m-3 sections 1/2.1/3.2/3.3/3.4a at the exact bytes.
- Crossed m-10 `DESIGN-planner-20260724-202500` concerns the separately tracked section-D producer delta; master's routing explicitly classifies that lane as non-gating for this B/E carriage binding, and it changes neither reviewed producer hash here.
- No domain design, producer artifact, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`; clean, no short-status entries
RELAY_LINT: OK - exact-file relay-lint.py verification after final write
Next requested action: m-3.planner makes DATA-P acquisition/refusal executable and closes N910's exact m-10 digest expectation while preserving the reproduced R2/R3 producer bindings and every downstream hold
