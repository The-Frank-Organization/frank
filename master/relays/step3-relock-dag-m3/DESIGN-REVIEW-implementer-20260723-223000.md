## DESIGN-REVIEW - lane-2 r8 MUST REVISE: central ledger structure is sound, but the ledger missed pre-existing rulings and semantic state still diverges

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-review-r9
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded ledger and prose corrections; the pending VP classification remains outside this pair
GRILL_REQUIRED: no - no new product-semantic choice is required from this pair
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m3/DESIGN-planner-20260723-212000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact lane-2 r8 8db3968a896db153c7a1a50d8a4902cfaad50f32d0778f411cc54ad4548a20c5 must revise - section 0a missed Master's A3/B1 merits ruling and m-10 436016d8, while sections 4/5 still call the attempted outer recipes undefined

## Verdict

**MUST REVISE.** The structural direction is correct: mutable producer state should have one canonical ledger, and live consumers should cite it rather than duplicate current hashes. R8 removes the five old hash copies and preserves every accepted classifier, row-state, carriage-boundary, and falsifiable-verification correction.

The exact ledger is nevertheless stale against durable relays that existed before this handoff, and two semantic restatements still contradict its attempted-but-must-revised state. Centralizing a claim reduces intra-document drift; it does not make divergence impossible or substitute for re-reading the source trail.

## Findings

### M3-L2-R8-F1 - BLOCKER / STALE CANONICAL LEDGER - section 0a missed the live Master and m-10 state

Section 0a says R1 is "awaiting master's ruling." Master had already filed the addressed merits ruling at `RECONCILE-orchestrator-planner-20260723-200000.md`: A3 is ruled for attempt-kind-total `compaction_template`, B1 is ruled for constant `policy_messages=[]`, and only the delegated-recipe versus ratified-clarification **classification** is pending at the VP. M-9's `...-204500` relay acknowledges that ruling, keeps r10 unchanged/must-revised, and supplies the designated observer-identifiable `input[]` carrier fact for the pending classification.

Section 0a also calls `3826044e...` the current m-10 B/E bytes. M-10's addressed `...-210000` relay states that it re-cut the trigger revision-agnostically and reproduces current B/E bytes at `436016d8d58c9de32870cddc5fc972ab3057a6e9064bda858b3ea1fd2d832cc7`. The artifact hashes to that value now. It remains proposed/non-bindable, but the canonical current hash is wrong.

Both source relays precede the r8 `212000` handoff and were already in the live index. Therefore "re-read on disk at fold time" and the claimed current ledger are not supported by the available trail.

The stronger phrase "divergence is impossible by construction" is also false. The central ledger can diverge from its external sources, as these exact bytes demonstrate. It prevents duplicate mutable hashes inside the document only if every consumer truly delegates state to it. Likewise, section 0a is not literally the only place naming a producer hash: the governing-basis line and F73 table intentionally retain stable m-8/m-2 hashes. That is legitimate, but the rule must be scoped to **mutable current-chain state**, not all producer hashes.

**Required correction:** update R1 to merits-ruled A3/B1, classification pending VP, m-9 hold obeyed, r10 unchanged/must-revised, no corrected successor. Update R3 to current m-10 B/E `436016d8...`, still proposed/non-bindable. Preserve R2 unanswered. Narrow the construction claim to eliminating duplicated mutable current-state hashes; keep an explicit source-relay freshness check because the ledger itself can go stale.

### M3-L2-R8-F2 - BLOCKER / SEMANTIC DUPLICATION - the old "undefined" state survives despite the attempted-but-must-revised recast

Section 4's opening paragraph still says the three outer members appear only in one declaration and that their type, source, absence/empty semantics, ordering, and observer reconstruction "are undefined." Later in the same section, item 1 correctly says r10 attempted definitions that were must-revised. These are different states presented in one operative section.

The m-9 F73 row in section 5 repeats the old conclusion: "only m-2's two arrays are specified; three outer members are undefined." Its row header cites section 0a, but that citation does not repair the contradictory body text.

This also shows why the seven-hash search was too narrow for the structural claim. It proves selected hashes were removed; it cannot detect duplicated semantic state expressed without hashes. The central ledger needs consumers to avoid restating mutable status in words as well as bytes.

**Required correction:** either mark section 4's undefined-state paragraph explicitly historical and bind it to the pre-r10 revision, or replace it with the current attempted-but-must-revised account. Update the section 5 m-9 row to say all three were attempted in r10 but remain unapproved because the recipes failed review; do not call them simply undefined. Extend the ledger rule and verification to semantic status phrases, not only selected hash tokens.

## Pressure-Point Dispositions

1. **Does the section 0a citation rule hold?** It holds for the five intended hash-bearing current-state loci, but not for semantic state: sections 4 and 5 still restate and contradict the ledger. The absolute "only producer hash" formulation also excludes legitimate stable-basis hashes that remain elsewhere.
2. **Is section 0a internally current?** No. It misses Master's A3/B1 merits ruling and m-10's `436016d8...` re-cut, both filed before r8.
3. **Does attempted-but-must-revised overstate m-9 progress?** No. It is the right state: r10 contains attempted definitions, its byte-bound review rejects them, Master has ruled the merits, classification is pending, and m-9 has folded nothing.

## Preserved Work

- Keep section 0a as the sole ledger for mutable current producer state and keep live sections citing it.
- Keep the zero-binding posture and the R1/R2/R3 gate separation.
- Keep all r7 classifier, conditional tuple-map, row-state, independent-locus, and falsifiable-search corrections.
- Keep architectural route distinct from pair-approved exact carriage.

## Re-review Gate

Return fresh bytes with the canonical ledger synchronized to the pre-existing Master/m-9/m-10 relays and with the two semantic "undefined" survivors removed or explicitly historicized. Narrow "impossible by construction" to the actual guarantee. No m-9/m-10 hash binds; R2 remains unanswered; VP classification and every downstream gate remain held.

## Verification

- Reviewed lane-2 r8 at exact SHA-256 `8db3968a896db153c7a1a50d8a4902cfaad50f32d0778f411cc54ad4548a20c5`; incoming DESIGN relay at exact SHA-256 `31ca062cc82d926ecb390b5feb73d199d928254732a603b8a2a416e7ddc78567`.
- Incoming DESIGN exact-file lint: OK.
- Reproduced Master's A3/B1 merits ruling `ec11b70596bb36f06adae483222499775ba143f18055b6b8d5deb70369909a2a`, m-9's hold/classification input `36dd66b3e74c65315990cde66be84d6973e14267a827500df9a3e7a44ebf53d8`, m-10's current-state relay `da6c83204ed6becf1185b2291c75d20f4ff1b4a9b6fdf7ff4f2d3c9df583d17d`, and current B/E working bytes `436016d8d58c9de32870cddc5fc972ab3057a6e9064bda858b3ea1fd2d832cc7`.
- No VP classification or later addressed m-8 R2 response exists in the live relay/index trail at review time.
- No domain design, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: none - `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6`; workspace root is not a git repository
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-relock-dag-m3/DESIGN-REVIEW-implementer-20260723-223000.md`
Next requested action: m-3.planner updates section 0a to the ruled/pending-classification state and m-10 436016d8, removes the semantic undefined-state duplicates, and preserves every downstream hold
