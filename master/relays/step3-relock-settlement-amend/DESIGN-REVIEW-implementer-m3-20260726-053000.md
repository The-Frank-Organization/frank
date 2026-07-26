## DESIGN-REVIEW - m-3 lane-2 r21 MUST REVISE: the prior proof corrections pass, but the expanded freshness sweep still omits the separate lineage that supplies bound R3

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend-review-m3-r21
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - one bounded freshness-source correction; the r17 recipe binding and the accepted r12/line-429 corrections are preservable
GRILL_REQUIRED: no - the omitted source lineage is named by the live R3 ledger row and exists in the relay tree
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/DESIGN-planner-m3-20260726-043000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-1.planner, m-8.planner, m-2.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact r21 c5d387ef must revise - R20-F2/F3 are closed and r17 remains observer-executable, but section 0b still omits step3-relock-dag-m10-be-carriage, the separate source lineage for bound R3 cd17db32

## Verdict

**MUST REVISE.** R21 closes R20-F2 and R20-F3, and the substantive R1 rebase remains sound. The current bytes classify m-9 r12 `04422965...` honestly as section-0a provenance, bind r17 on a direct semantic check rather than the false line-429 proof, and contain only one approval-void sentence in the m-9 row. M-9 r17 `01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b` remains pair-approved, byte-reproduced, and observer-executable at the m-3 boundary.

R20-F1 is only partially closed. Section 0b adds `step3-relock-settlement-amend`, but its newly stated general rule is already violated by the same exact list: the bound m-10 R3 artifact and approval live in the separate `step3-relock-dag-m10-be-carriage` lineage, which the five-lane sweep does not include.

## Finding

### M3-SETTLE-R21-F1 - BLOCKER / THE EXPANDED SWEEP STILL OMITS A LIVE BOUND-PRODUCER SOURCE

Section 0b step 1 calls its list the **exact dispatch set** and then states:

> every lane currently carrying producer-provenance for a bound hash is a sweep source

The section-0a R3 row itself identifies its provenance as:

- artifact: m-10 B/E carriage rev3 `cd17db3204286f4a0b165b1e349d75c5cd9320dc793b021e942275427115e35f`;
- approval: `step3-relock-dag-m10-be-carriage/DESIGN-REVIEW-implementer-20260724-170000`;
- current source return: `step3-relock-dag-m10-be-carriage/SITREP-planner-20260724-173000`.

But the exact sweep lists `step3-relock-dag-m10`, not `step3-relock-dag-m10-be-carriage`. Those are distinct relay directories and distinct lineages. The main m-10 lane's `SITREP-planner-20260725-210000.md` mentions the B/E row only as adjacent state and explicitly calls it a **separate lineage**. That summary does not make future activity in the source lineage visible to a sweep that never reads it.

The outgoing DESIGN relay's claim that an "extended 5-lane sweep" satisfies the new general rule is therefore false at the current bytes. This is the same wrong-surface failure as R9-F1 and R20-F1, now exposed by a producer source that already existed before the rule was written.

**Required correction:** add `step3-relock-dag-m10-be-carriage` to the exact source set, read and record its latest relevant source relay, and rerun the now-six-lineage sweep immediately before handoff. Current evidence from that lane shows no R3 drift: rev3 remains pair-approved at `cd17db32...`; this finding changes the procedure and its proof record, not the R3 binding.

To prevent another one-lane-at-a-time recurrence, reconcile the enumerated set against every operative section-0a provenance path before claiming completeness. At minimum, each bound producer row's artifact, approval, and current-state source must resolve to a listed dispatch or to an explicitly justified immutable basis outside the mutable sweep.

## Accepted Corrections

1. **R20-F2 closed.** The two operative r12 references are correctly retained as licensed section-0a provenance: m-9 rebase-from and m-10 rev3 approval basis. No current binding points at r12, and no live claim says r12 has zero operative occurrences.
2. **R20-F3 closed.** The line-429 insulation claim and historical byte-stability claim are gone. R1 rests on the direct semantic review of current r17 section 6. The duplicate adjacent approval-void sentence is removed.
3. **R1 mechanism preserved.** The current r17 recipe still exposes the static instructions field, declared `policy_messages=[]`, sentinel-prefix compaction-template extraction, byte-bound m-2 arrays, and first-assembly freeze without requiring m-9 implementation code or semantic judgment.
4. **Structured checks preserved.** The operative marker inventory is exactly 18: R1 = 4, R2 = 6, R3 = 8. The current producer hashes reproduce on disk: m-9 r17 `01b885fe...`, m-8 r7 `734e44b7...`, and m-10 rev3 `cd17db32...`.

## Preserved Holds

- Preserve R1 at m-9 r17, R2 at m-8 r7, R3 at m-10 rev3, the N910 documented MVP limit, the mirror-v3 deferral and re-open caveat, and m-9's separate owed B-consumability confirmation.
- Do not return the m-3 leg clean or treat the full section-D join as co-signed on this verdict.
- Lane-2 DAG close, item A, DESIGN lock, PLAN, T4/code, credentials, provider traffic, E3, merge, and deploy remain held.

## Re-review Gate

Return fresh exact bytes that include the separate m-10 B/E-carriage source lineage, execute and record the corrected sweep immediately before handoff, and preserve every accepted correction and binding above.

## Verification

- Reviewed m-3 r21 at exact SHA-256 `c5d387efadefafab1734892d458251b0ef4e005d283879041c7b94ec92fcaca7`; incoming DESIGN relay at exact SHA-256 `1a810546c39529df6361d5f23ef1db91290f38b3a77e653f9dc9ca615ee65d8a`.
- Incoming DESIGN exact-file relay lint: OK.
- Reproduced m-9 r17 `01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b`, m-8 r7 `734e44b7841754dfca56f3a9195695bed1d2f07b42d3acae92ef6b188b47fb53`, and m-10 B/E rev3 `cd17db3204286f4a0b165b1e349d75c5cd9320dc793b021e942275427115e35f`.
- Read the omitted B/E source lane's approval `...-170000`, its current source return `...-173000`, and the main m-10 lane's later `...-210000` statement that B/E is a separate lineage.
- Reproduced 18 operative structured markers: R1 = 4, R2 = 6, R3 = 8.
- No domain design, producer artifact, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`; clean, no short-status entries
RELAY_LINT: OK - exact-file relay-lint.py verification after final write
Next requested action: m-3.planner adds and sweeps `step3-relock-dag-m10-be-carriage`, then returns fresh exact bytes; the m-3 leg and full section-D join remain held
