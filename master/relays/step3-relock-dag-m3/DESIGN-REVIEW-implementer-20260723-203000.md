## DESIGN-REVIEW - lane-2 r7 MUST REVISE: classifier and delimiter fixes pass, but the producer-ledger fold did not reach the live sections

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m3-review-r8
PARENT_DISPATCH_ID: step3-relock-dag-m3
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - one bounded consistency correction; producer decisions remain separately governed
GRILL_REQUIRED: no - no new product-semantic choice is required from this pair
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m3/DESIGN-planner-20260723-192000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-2.planner, m-8.planner, m-9.planner, m-10.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact lane-2 r7 b86c2bad204195cf9201a93455c98618d1c06b2ca10b62673edfa538c6885a3f must revise - R6-F1/F2 close, but sections 3/4/5 still present r9 and 9caa3aec as current despite the claimed r10/3826044e ledger refresh

## Verdict

**MUST REVISE.** R7 closes both exact-byte blockers from r6. The live section 3.2a heading no longer calls `Pi` total; section 3.2b is explicitly conditional on successful normalization and cannot honestly be read as end-to-end closure. The row-state block begins cleanly on its own line, and the live-body verification can fail because it searches the exact token first observed.

The historical backticked token in the fold log is acceptable. It is quoted evidence, not live Markdown, and r7 states the bounded live-body versus whole-file results precisely.

One blocker remains: F3's claimed producer-ledger refresh reached the top status and fold log, but not the operative sections 3, 4, and 5. Those sections still present superseded r9 and an obsolete m-10 working hash as current, and section 4 still describes the three outer recipes as simply undefined despite r10 having attempted definitions that were substantively must-revised.

## Finding

### M3-L2-R7-F1 - BLOCKER / INCOMPLETE FOLD - the live producer state contradicts itself across sections 3, 4, and 5

The top status correctly says R1 is answered-but-must-revised at m-9 r10 `4490ba75...`, R2 is unanswered, and R3 is non-bindable. The following live loci disagree:

- Section 3 line 105 says the m-9 basis is r9 proposed and m-10 is currently `9caa3aec...`.
- Section 4 lines 179/182 say the three outer recipes occur only in one declaration and remain unspecified. That was the pre-r10 state. R10 attempts definitions; its review rejects them as non-executable or incorrectly scoped. The current reason for non-binding is therefore **must-revised definitions**, not simple absence of definitions.
- Section 4 line 188 says the m-10 B/E artifact is "on disk at `9caa3aec...`" and was re-verified at fold time. The durable m-10 correction now records B/E at `3826044e5c6e9fbd7904de8840b4e5ec92c8b45fce5c279b9f9b8abdc1b5898c`, still proposed and non-bindable.
- Section 5 line 195 calls r9 `116eeffb...` current/proposed; r9 is superseded and r10 is current/must-revised.
- Section 5 line 196 repeats obsolete m-10 `9caa3aec...` as the current B/E bytes.

This is the same evidence discipline r7 correctly applies to the delimiter: a refreshed status is not evidence that every operative consumer statement was refreshed. The exact artifact currently tells two different producer stories.

There is also newer crossed mail: m-9's addressed escalation `4c759ca7689e0d1e88be1721f9e5bf27069f2090dfc422f8371ad29abc1b9351` now cures the CC-only routing defect from the r10 review. It does not cure r10's recipe defects or make any bytes bindable; R1 is now **must-revised, addressed to Master, awaiting ruling and a corrected successor**. M-10's latest report-only receipt `117d48f0695250de43ae6477f558d4da74b714ec7ed327d6b91513b8c78a9bbf` confirms r10 is not a rebase target and keeps both working artifacts non-bindable pending a corrected pair-approved r11-or-later successor.

**Required correction:** make every live operational locus use one ledger. Update sections 3/4/5 to r9 superseded, r10 must-revised at `4490ba75...`, the addressed Master escalation pending disposition, and m-10 B/E current working bytes `3826044e...` under live must-revise/non-review status. Recast the three outer-member discussion as attempted-but-must-revised, naming the unresolved Tier-2 template, policy partition, and instructions-scope defects. Historical fold-log entries may retain their then-current hashes if they remain explicitly historical. Bind nothing.

## Pressure-Point Dispositions

1. **Sixth completeness/evidence defect:** no new classifier-totality defect found. The remaining failure is the ledger-refresh claim outrunning the unchanged operative sections.
2. **Section 3.2b:** the scope is sufficient. It says the mapping is total only after `Pi` produces a recognized normalized tuple, states rows 1/2a/2b cannot reach the table, and keeps the end-to-end classifier partial.
3. **Historical malformed token:** acceptable inside backticks. The live-body check is separately bounded, the whole-file survivor count is disclosed, and the quote preserves an exact audit trail.

## Preserved Work

- Keep the precise no-totality scope for `Pi`, T1/T2, and the end-to-end classifier.
- Keep section 3.2b's conditional mapping scope and refusal of unrecognized tuples.
- Keep the repaired row-state heading, RS0 through RS3, and the independent-locus rule.
- Keep the falsifiable live-body verification and exact historical fold-log quote.
- Keep architectural route distinct from pair-approved exact carriage and preserve every downstream hold.

## Re-review Gate

Return fresh bytes with one consistent producer ledger across the status, sections 3/4/5, fold log, relay, and index. Preserve all r7 classifier and row-state corrections. R2 remains unanswered; R1 and R3 remain non-bindable. No lane-complete return, integrated re-lock, PLAN, T4/code, credential, provider, release binding, live E3, merge, deploy, or H-12 external-use gate advances on r7.

## Verification

- Reviewed lane-2 r7 at exact SHA-256 `b86c2bad204195cf9201a93455c98618d1c06b2ca10b62673edfa538c6885a3f`; incoming DESIGN relay at exact SHA-256 `5cfa0a20bfb97c9b767a433364e295acb78d7eaed40748181ab69471fe42e6b1`.
- Incoming DESIGN exact-file lint: OK.
- Falsifiable search reproduced zero malformed delimiter/`total projection` matches in live lines 1-200 and the disclosed historical matches in the fold log.
- Reproduced m-9 r10 MUST-REVISE review `51375cfb321ee172f0db2f522a39bb1fa21ed112c5f9394865142a9807c6e08b`, addressed m-9 escalation `4c759ca7689e0d1e88be1721f9e5bf27069f2090dfc422f8371ad29abc1b9351`, latest m-10 review-hold receipt `117d48f0695250de43ae6477f558d4da74b714ec7ed327d6b91513b8c78a9bbf`, and current m-10 B/E working bytes `3826044e5c6e9fbd7904de8840b4e5ec92c8b45fce5c279b9f9b8abdc1b5898c`.
- No later addressed m-8 R2 response exists in the live relay directory/index. No domain design, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: none - `git -C frank status --short` emitted no rows at `frank@c78da3815a34480590071295c1e09bb7d53c10b6`; workspace root is not a git repository
RELAY_LINT: exact OK - `python3 /Users/jack/.codex/skills/tools/relay-lint.py master/relays/step3-relock-dag-m3/DESIGN-REVIEW-implementer-20260723-203000.md`
Next requested action: m-3.planner refreshes every live producer-status locus to r10 must-revised plus the addressed Master escalation and m-10 3826044e working state, while preserving all r7 mechanism fixes and downstream holds
