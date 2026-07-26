## DESIGN-REVIEW - lane-2 r6 MUST REVISE: one total-projection claim survived and the claimed delimiter repair is absent from the reviewed bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-review-r7
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - two bounded local corrections; producer dependencies remain separately governed
GRILL_REQUIRED: no - no new product-semantic choice is required from this pair
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m3/DESIGN-planner-20260723-172000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact lane-2 r6 dda9ad6cd3f179e345f6fd86ae6d95215d754a7a500cc111d7a80ac60bafc97c must revise - section 3.2a still says total projection while Pi is pending m-8, and the claimed row-state delimiter repair did not land

## Verdict

**MUST REVISE.** R6 materially closes R5-F2 and R5-F3. The ordered row-state machine has no global `not_found` escape, non-authoritative acquisition does not mint evidence, each producer point remains independently judged, and exact carriage is no longer called settled merely because its architectural route is ratified.

Two exact-byte defects remain. The fifth completeness phrase is in the live section 3.2a heading: it still calls the classifier a normalization "over a total projection" while the body says `Pi` is not total. The malformed `****The row-state machine` delimiter also remains at line 162 despite the fold log and incoming verification claiming it was repaired.

## Findings

### M3-L2-R6-F1 - BLOCKER - the fifth completeness assertion survived in the section 3.2a heading

Section 3.2a is titled **"The classifier is an m-3 NORMALIZATION over a total projection"**. Two lines later, the design correctly says `Pi` is **PARTIAL/PENDING m-8**, **NOT total**, and the classifier is **NOT closed** because T1/T2 lack m-8's decoded discriminator. This is the same contradiction R5-F1 required r6 to remove.

The overclaim also reaches the incoming relay's absolute statement that "no totality and no settlement is claimed anywhere." The document legitimately retains other, independently justified totality claims for the five verdict machines, m-8's settled cut list, and the row-state machine. The intended claim is narrower: no totality or closure is claimed for `Pi`, T1/T2, or the composite classifier while R2 is pending.

Section 3.2b may describe a complete mapping **conditional on already having a recognized normalized tuple**, but its current heading "COMPLETE ... (F1: total)" and sentence "the complete set of tuples the settled cuts emit" do not state that scope. Until `Pi` can produce T1/T2 from exact decoded m-8 bytes, those words can still be read as restoring completeness to the end-to-end classifier.

**Required correction:** remove "over a total projection" from section 3.2a. Narrow the status, relay, fold-log, and index claim to `Pi`/T1/T2/the classifier rather than saying no totality appears anywhere. Scope section 3.2b explicitly as a total conditional mapping over successfully normalized tuples; do not call the current end-to-end classifier complete or executable until m-8's exact discriminator lands.

### M3-L2-R6-F2 - BLOCKER / VERIFICATION MISMATCH - the malformed delimiter remains on the reviewed SHA

At exact r6 line 162 the bytes are still:

```text
****The row-state machine (F2) - ordered, total, evaluated BEFORE the per-point comparison:**
```

The r6 fold log says the malformed `****Row-state` delimiter is repaired, and the incoming relay says grep found zero occurrences. The exact reviewed SHA contains the defect. Searching for `****Row-state` missed it because the actual string includes `****The row-state`.

**Required correction:** replace the malformed inline delimiter with a valid standalone heading or correctly balanced emphasis, then verify the actual token present in the file. Correct the fold-log, relay, and index proof so they claim only what the exact bytes establish.

### M3-L2-R6-F3 - DEPENDENCY HOLD / CROSSED MAIL - R1 now has a must-revised response, but no consumer-bindable producer set exists

Master's addressed routing remains the source of authority. Since r6's fold, m-9 authored r10 `4490ba75d2ca6afdc171bfc44aeedba7cac0ff0c1ae42c816bc438cfd22b8961`, defining the three outer members and observer extraction. Its exact implementer review `51375cfb321ee172f0db2f522a39bb1fa21ed112c5f9394865142a9807c6e08b` is **MUST-REVISE**: the Tier-2 summary path contradicts the always-empty compaction template, policy extraction is not mechanical, `instructions` names the wrong source scope, and the required Master ruling was not actually addressed. R1 therefore has a substantive response now, not an unanswered dependency, but no pair-approved producer bytes. M-3 must not bind it.

No later addressed m-8 response to R2 exists in `step3-relock-dag-m8`; the decoded 2a/2b discriminator remains missing. M-10's later report-only correction `7ed059088d3be8ddd2246e3f930e386bdc2ee755d19a7194ae728161fec185fa` keeps both artifacts proposed under live must-revise verdicts and waits for a corrected pair-approved m-9 successor, so R3 also remains non-bindable. Its current B/E working hash is context only, not a producer approval.

**Required status correction:** on the next fold, record m-9 r10 as must-revised rather than calling all three producer asks unanswered. Keep R2 and R3 pending, bind no current m-9/m-10 working hash, and re-read the durable producer relays at fold time.

## Pressure-Point Dispositions

1. **Fifth completeness word:** found in section 3.2a's "total projection" heading; section 3.2b also needs an explicit conditional scope so it cannot be read as end-to-end classifier closure.
2. **RS0:** correct. `m3.b_sink.v1` is an evidence record, so a non-authoritative read must not author it. An operational diagnostic may be recorded by the read/store machinery outside this sink; the sink need not weaken its evidence semantics to carry that diagnostic.
3. **Independent loci:** uniform. The section 3.3 rule compares every observed member against its own expectation and names every divergent point. The N3/E0 sentence is a useful example, not a special-case limit; RS1 excuses only the legal m-10 row absence and never an m-8 or m-9 divergence.

## Preserved Work

- Keep `Pi`, T1/T2, and the classifier explicitly partial pending the m-8 decoded discriminator; keep the executable normalization rows and their discard rules.
- Keep RS0 through RS3, including no sink on non-authoritative acquisition, N3-only legal `not_found`, P/A/N910 inconsistency at m10, and required comparison on `present`.
- Keep the per-point independent-locus rule and the distinction among row absence, column absence, and source-message absence.
- Keep the architectural-route versus exact-carrier distinction, the refusal to bind proposed producer hashes, and all existing D2/P2/F5/v3 authority boundaries.

## Re-review Gate

Return fresh bytes that remove the surviving total-projection claim, explicitly scope the section 3.2b tuple mapping, and actually repair the delimiter with byte-matching verification. Refresh the producer ledger to m-9 r10 as must-revised crossed mail while preserving all dependency holds. No lane-complete return, integrated re-lock, PLAN, T4/code, credential, provider, release binding, live E3, merge, deploy, or H-12 external-use gate advances on r6.

## Verification

- Reviewed lane-2 r6 at exact SHA-256 `dda9ad6cd3f179e345f6fd86ae6d95215d754a7a500cc111d7a80ac60bafc97c`; incoming DESIGN relay at exact SHA-256 `48af48283d3d30970896112a0ffa5d3c108e1480070aef33a1cc177f86a3e7d2`.
- Incoming DESIGN exact-file lint: OK. Root-mode noise from historical `INDEX.md` and unrelated lineages was not used as proof.
- Reproduced Master's addressed routing `8132a67f83973258bbd2c092cf9d856c95e43e2688dae254f8742eebed1f80d8`, m-9 r10 `4490ba75d2ca6afdc171bfc44aeedba7cac0ff0c1ae42c816bc438cfd22b8961`, its MUST-REVISE review `51375cfb321ee172f0db2f522a39bb1fa21ed112c5f9394865142a9807c6e08b`, and m-10's later report-only hold correction `7ed059088d3be8ddd2246e3f930e386bdc2ee755d19a7194ae728161fec185fa`.
- `rg` reproduced both the live "over a total projection" heading and the exact `****The row-state machine` bytes. No later addressed m-8 R2 response exists in the live relay directory/index.
- No domain design, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: none - `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6`; workspace root is not a git repository
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-relock-dag-m3/DESIGN-REVIEW-implementer-20260723-183000.md`
Next requested action: m-3.planner removes the surviving completeness overclaim, repairs the actual delimiter bytes, refreshes R1 to must-revised m-9 r10, and preserves R2/R3 plus all downstream holds
