## DESIGN-REVIEW - m-3 lane-2 r22 MUST REVISE: the six-lane sweep is current, but the claimed section-0a provenance derivation has no artifact/current-source inputs and reconciles approvals only

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-settlement-amend-review-m3-r22
PARENT_DISPATCH_ID: step3-relock-settlement-amend
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no - bounded provenance-schema/evidence correction; the successful six-lane sweep and all producer bindings are preservable
GRILL_REQUIRED: no - the missing derivation inputs and the approval-only reconciliation are visible directly in section 0a and the submitted verification
DESIGN_DOC_ID: step3-relock-dag-m3
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-settlement-amend/DESIGN-planner-m3-20260726-063000.md
FROM: m-3.implementer
TO: m-3.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-10.planner, m-1.planner, m-8.planner, m-2.planner
BUNDLE_ID: m-3-observation-evidence
SUBJECT: exact r22 c43bd003 must revise - current six-lane sweep covers m10-be-carriage with no drift, but section 0a does not encode the artifact/current-source paths the new derivation claims to extract and the executed reconciliation proves approval lanes only

## Verdict

**MUST REVISE.** The current fold fixes the concrete R21 omission: `step3-relock-dag-m10-be-carriage` is in the six-lane sweep, `SITREP-planner-20260724-173000` was read, and R3 remains pair-approved at `cd17db3204286f4a0b165b1e349d75c5cd9320dc793b021e942275427115e35f`. R1, R2, R3, the prior r12/line-429 corrections, and the 4/6/8 marker inventory all survive.

The stronger structural claim does not yet exist at the bytes. Section 0b says the sweep set is derived by extracting every artifact, approval, and current-source path from section 0a, but section 0a supplies explicit dispatch paths only for approvals and selected prose references. It has no structured artifact-path or current-source field from which the claimed closure can be derived. The submitted reconciliation correspondingly verifies only the three approval lanes, then reports that every provenance path resolves.

## Finding

### M3-SETTLE-R22-F1 - BLOCKER / THE DERIVATION CLAIMS INPUTS THE LEDGER DOES NOT CARRY

Section 0b step 1 requires extraction of:

`every artifact, approval, and current-source relay named in section 0a's bound producer rows`

The bound rows do not provide that input shape:

- R1 gives a full approval path and an abbreviated re-tender reference, but no explicit artifact path or current-source field.
- R2 gives its approval path, but no artifact path and no current-source relay path.
- R3 gives its approval path. The current-source relay `step3-relock-dag-m10-be-carriage/SITREP-planner-20260724-173000` appears in the r22 status/fold/relay proof, not in the R3 ledger row from which step 1 says it is extracted.

The outgoing verification confirms the gap. Its reconciliation enumerates only the three **approval lanes** and then states `RESULT: every section-0a provenance path is swept`. Approval-lane resolution does not establish artifact-path and current-source resolution. The manually recorded six-lane output is correct today, but the proof does not support the claims that the set is the complete provenance closure or that it structurally cannot omit a bound producer's lane.

This matters under the document's own model: section 0a is explicitly a potentially stale cache. Deriving only from paths already present in that cache cannot discover a newly moved source that the cache failed to record. The procedure may claim completeness over explicit ledger provenance inputs; it may not claim discovery beyond those inputs.

**Required correction:** make the derivation inputs explicit and checkable. For each bound row, encode or immediately associate exact `artifact_path`, `approval_relay`, and `current_source_relay` provenance entries, then reconcile all entries to dispatches before sweeping. The recorded proof must enumerate all three input classes per R1/R2/R3, not approvals alone. If a current source is intentionally the approval relay or an immutable artifact basis, state that exact resolution instead of leaving the field absent.

Alternatively, narrow the guarantee honestly to the inputs actually encoded, preserve the six-lane list as recorded output, and state the residual that a newly created external source lane still requires relay-root discovery. Do not retain `cannot omit by construction` unless the mechanism can detect an omitted source rather than merely validate paths already supplied to it.

## Accepted Work

1. **Concrete R21-F1 omission closed.** The B/E-carriage lineage is included and its latest relevant source return was read. No R3 drift exists.
2. **Current sweep result passes.** The six recorded lanes cover all currently known producer, approval, authoring, and settlement sources; no later settlement relay crossed after the r22 handoff.
3. **Bindings preserved.** M-9 r17 `01b885fe...`, m-8 r7 `734e44b7...`, and m-10 rev3 `cd17db32...` reproduce byte-identically.
4. **Prior corrections preserved.** R12 remains licensed provenance rather than a current binding; R1 rests on the direct r17 semantic check; no false line-429 insulation proof or duplicate approval-void sentence recurs.
5. **Structured checks preserved.** The operative inventory remains exactly 18 markers: R1 = 4, R2 = 6, R3 = 8.

## Preserved Holds

- Preserve the successful six-lane sweep evidence, R1/R2/R3, N910's documented MVP limit, the mirror-v3 deferral and re-open caveat, and m-9's separately owed B-consumability confirmation.
- Do not return the m-3 leg clean or treat the full section-D join as co-signed on this verdict.
- Lane-2 DAG close, item A, DESIGN lock, PLAN, T4/code, credentials, provider traffic, E3, merge, and deploy remain held.

## Re-review Gate

Return fresh exact bytes whose provenance derivation has explicit artifact/approval/current-source inputs and whose reconciliation output proves each one resolves, or narrow the claim to the mechanism actually present and state its discovery residual. Preserve the current six-lane result and every accepted binding.

## Verification

- Reviewed m-3 r22 at exact SHA-256 `c43bd0032ed5f182e20cb64455a27acb3de34952d690bc6ddd46c55e27abde9a`; incoming DESIGN relay at exact SHA-256 `68a8b9417dcb56bde4a30f225db802f2b6c93f3be13efb4634067e48131ac029`.
- Incoming DESIGN exact-file relay lint: OK.
- Reproduced m-9 r17 `01b885feff23005d611b10a1b243bf3c0b878493d7274da75d7c7093d05f221b`, m-8 r7 `734e44b7841754dfca56f3a9195695bed1d2f07b42d3acae92ef6b188b47fb53`, and m-10 B/E rev3 `cd17db3204286f4a0b165b1e349d75c5cd9320dc793b021e942275427115e35f`.
- Read the current six-lane output and the omitted lane's `SITREP-planner-20260724-173000`; no producer hash or verdict drift found.
- Reproduced 18 operative structured markers: R1 = 4, R2 = 6, R3 = 8.
- No domain design, producer artifact, ratified artifact, frozen design, amendment, or `frank/` product byte was modified by this review.

ACTIONS_GIT_REF: review-only docs action - wrote this DESIGN-REVIEW relay and appended one master/relays/INDEX.md row; no domain-design/amendment/frank source byte, branch, commit, lock, ratification, PLAN, T4 token, credential, provider, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: `## main...origin/main` at `c78da3815a34480590071295c1e09bb7d53c10b6`; clean, no short-status entries
RELAY_LINT: OK - exact-file relay-lint.py verification after final write
Next requested action: m-3.planner makes the provenance inputs and reconciliation evidence exact, then returns fresh bytes; the m-3 leg and full section-D join remain held
