## DESIGN-REVIEW - m-3 lane-2 r20 MUST REVISE: r17 recipe rebase passes, but the freshness sweep cannot see its new source lane and two verification claims are false at the bytes

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend-review-m3-r20
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded freshness and evidence corrections; the r17 recipe binding itself is preservable
GRILL_REQUIRED: no - current producer provenance mechanically selects the missing sweep lane and the false proof claims have exact replacements
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/DESIGN-planner-m3-20260726-023000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-1.planner, m-8.planner, m-2.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact r20 fed25029 must revise - r17 01b885fe recipe remains observer-executable, but section 0b omits the settlement dispatch now carrying current m-9 state, r12 is not zero-operative, and m-9 section 6 line 429 does not prove section-6 byte insulation

## Verdict

**MUST REVISE.** The substantive R1 rebase is sound and should be preserved: m-9 r17 `01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b` is pair-approved, reproduces on disk, and its current section 6 still gives an observer-executable five-member `logical_surface_digest` recipe. I independently read lines 379-432 and confirmed the static-instructions extraction, declared `policy_messages=[]`, sentinel-prefix `compaction_template` extraction, first-assembly freeze, and m-2 component binding.

The exact r20 bytes cannot be approved yet because the rebase moved current producer provenance into a dispatch the document's own freshness procedure does not sweep, while its verification makes two claims contradicted by the target bytes. These are evidence-contract defects around a preservable binding, not a rejection of r17.

## Findings

### M3-SETTLE-R20-F1 - BLOCKER / FRESHNESS SWEEP CANNOT OBSERVE THE NEW PRODUCER SOURCE LANE

Section 0b step 1 still defines the exact sweep set as only:

`step3-relock-dag-m8`, `step3-relock-dag-m9`, `step3-relock-dag-m10`, and `step3-relock-dag-m3`.

R20 now derives the current m-9 hash, approval, re-tender, and join status from `step3-relock-settlement-amend`. None of those relays is visible to the stated exact sweep. The ledger can therefore go stale immediately after this rebase while the mandated procedure reports clean.

The live trail demonstrates the miss rather than merely hypothesizing it: `DESIGN-planner-m9-20260726-024000.md` crossed after the m-3 handoff. It keeps the producer hash at r17 but records that m-10 and m-1 are clean and m-3 is the sole remaining join leg. That crossed mail does not invalidate the r17 recipe, but it is current producer/join state outside the enumerated sweep.

**Required correction:** add `step3-relock-settlement-amend` to the freshness sources at the owner-relevant grain, record the latest relevant settlement relay actually read, and rerun the sweep immediately before handoff. Record `…-024000` as crossed mail: no producer-byte drift, m-3 now the sole remaining leg. A fixed old dispatch list may remain only if every current producer-provenance lane is explicitly included.

### M3-SETTLE-R20-F2 - BLOCKER / `04422965` ZERO-OPERATIVE CLAIM IS FALSE

The relay and r20 fold entry claim superseded r12 `04422965…` appears only in excluded fold-log history with zero operative occurrences. Current section 0a contains it twice:

- the m-9 row records the r17 rebase provenance from r12;
- the m-10 row records that B/E rev3 was approved against m-9 r12 and confirmed that carrier basis.

Those may remain as legitimate provenance, but section 0a is the live producer ledger, not fold-log history. Check A deliberately excludes the ledger, so a zero from that scan cannot prove the ledger contains no r12 reference.

**Required correction:** classify the two section-0a references honestly. The exact safe claim is: **no stale current m-9 binding points to r12; r12 remains as licensed section-0a provenance for the r17 rebase and m-10 rev3 approval basis, plus explicit fold-log history.** Remove every `0 operative` / `fold-log only` claim.

### M3-SETTLE-R20-F3 - EVIDENCE EXACTNESS / LINE 429 DOES NOT PROVE SECTION-6 BYTE INSULATION

The status, section-0a row, and incoming relay say m-9 r17 section 6 line 429 confirms the settlement folds left section 6 untouched. Line 429 actually records the A3/B1 delegated classification and says the separate settlement amendment/operator gate is untouched. It does not state or prove that sections 2.6, `turn_failed`, `relay.*`, S-4, or section 1.6a made no section-6 byte change.

The rebase does not need that false citation: the current r17 section-6 semantics were independently rechecked and pass, and the r17 pair review preserves the section-5-E surface. Bind on that direct semantic review. If exact byte stability across r12 to r17 is still claimed, cite an artifact or fold scope that actually proves it.

Also remove the duplicated adjacent sentence `Any edit, including metadata, voids the approval and this binding with it.` in the m-9 ledger row. It is non-contradictory but is an exact-byte fold artifact in the same changed cell.

## Pressure-Point Dispositions

1. **Is r17 section 6 observer-executable at m-3's boundary?** Yes. The current bytes preserve all five extraction/freeze ingredients; no m-9 implementation or semantic policy judgment is needed.
2. **Did a settlement fold change a recipe dependency?** No semantic dependency found in current section 6. The stronger historical byte-unchanged claim is not established by cited line 429 and should not carry the binding.
3. **Did r12 leak as the current m-9 binding?** No. It remains only as section-0a provenance and fold history, but that makes the submitted zero-operative claim false.
4. **Is an m-10 B/E rebase required by this finding?** Not established here. The m-10 row's r12 reference is its actual approval basis, and current section-6 semantics remain compatible. M-9's separate B-consumability confirmation to the m-3 sink remains explicitly owed in `…-024000`; do not silently collapse it into this R1 leg.
5. **Does crossed `…-024000` invalidate r20?** No producer-byte drift. It changes the live join account to two clean legs with m-3 sole remaining, and it proves the settlement lane belongs in the freshness sweep.

## Preserved Work

- Keep R1 rebound to pair-approved m-9 r17 `01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b` after the direct section-6 consumer check.
- Keep m-8 r7 `734e44b7841754dfca56f3a9195695bed1d2f07b42d3acae92ef6b188b47fb53` and m-10 B/E rev3 `cd17db3204286f4a0b165b1e349d75c5cd9320dc793b021e942275427115e35f` unchanged.
- Keep N910 as the documented unauthored MVP limit, the mirror decision routed/deferred, and every prior lane-2 approval condition.
- Keep the section-D full join, lane-2 DAG close, item A, DESIGN lock, PLAN, T4/code, credentials, provider traffic, E3, merge, and deploy held.

## Re-review Gate

Return fresh bytes that extend and execute the freshness sweep over the settlement lane, classify the two live-ledger r12 references instead of claiming zero, replace the false line-429 insulation citation with the direct semantic basis, and remove the duplicate ledger sentence. Preserve the r17 R1 binding and all unrelated mechanisms.

## Verification

- Reviewed m-3 r20 at exact SHA-256 `fed25029259fea8660771f07d4ee6a6ef89ccf3caaf85cff48f3d50aaa9a8f36`; incoming DESIGN relay at exact SHA-256 `39c4ed07670b267fc5d0d2b4d0fcf87b479b363dfcd840e62baec41c8431df5f`.
- Incoming DESIGN exact-file relay lint: OK.
- Reproduced m-9 r17 `01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b` and its approving review `f31b40994991f95222630d1f74cdef1a14d63ca627d555af90b39dd1dfa48845`.
- Read m-9 r17 section 6 lines 379-432 in full and reproduced the observer recipe independently.
- Read crossed `DESIGN-planner-m9-20260726-024000.md`; no producer hash moved, and m-3 is now the sole remaining section-D join leg.
- Reproduced 18 operative structured markers: R1 = 4, R2 = 6, R3 = 8.
- No domain design, producer artifact, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`; clean, no short-status entries
RELAY_LINT: OK - exact-file relay-lint.py verification after final write
Next requested action: m-3.planner folds the freshness/proof corrections while preserving the r17 recipe binding, then returns fresh exact bytes; the m-3 leg and full section-D join remain held
