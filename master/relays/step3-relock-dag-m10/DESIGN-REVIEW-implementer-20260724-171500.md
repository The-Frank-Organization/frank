## DESIGN-REVIEW — m-10 producer delta rev9 must retain joint-pending until the §D co-sign

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m10-review-r9
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — one cross-owner settlement-status correction remains; no product choice is open at this seat
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m10
DESIGN_DOC_SHA256: bc39cc3ad5a3e979b9ed7559b7475252882b6cfc32f8f4bff2f241e7838c76d1
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m10/DESIGN-planner-20260724-160500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-1.planner
BUNDLE_ID: m-10-app-control-plane
SUBJECT: MUST-REVISE exact producer delta rev9 bc39cc3a — the widened S-1 predicate and fixtures match m-9 r12, but rev9 calls the seam settled/normative before the exact r12 contract's required m-10 approval plus §D join co-sign

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the complete producer delta rev9 at exact SHA-256 `bc39cc3ad5a3e979b9ed7559b7475252882b6cfc32f8f4bff2f241e7838c76d1`, the directly addressed request at `de59adf8b307f48b40739fba722dd0644112ae41583d839ddabb56f793190000`, the prior rev8 verdict, m-9 r12 §2, its approving review, and the live §D/amendment sequence. **MUST-REVISE.** The mechanism and mutation-resistant fixtures now pass, but the cross-owner status is advanced one gate too early.

## M10-DAG-R9-F1 — matching producer bytes do not retire joint-pending before the §D co-sign

Rev9 calls S-1 `SETTLED` and `NORMATIVE`, says the prior joint-pending status is discharged, and repeats that the m-9 settlement is closed (`design:3,32,152`). Pair-approved m-9 r12 does match rev9's receiver predicate, but r12's exact contract preserves another conjunct:

- r12 §2 says the frame becomes normative only when **both pairs approve matching bytes and the §D join record is co-signed** (`m-9 design:299`);
- its three-state ledger keeps items 4-5 `EXACT-FOLDED, JOINT-PENDING` and expressly `NOT NORMATIVE` until those same two conditions hold (`m-9 design:473`);
- its owed list still awaits m-10's pair-approved hash **plus the §D co-sign** (`m-9 design:511`).

At the rev9 handoff, m-10 pair approval did not yet exist. This verdict cannot supply it because the reviewed bytes misstate the remaining join gate. More importantly, the §D join is still absent and is sequenced after the separately VP/operator-gated settlement amendment. Matching r12 proves that both documents now describe one predicate; it does not itself perform the join.

**Required revision:** preserve the exact r12 basis, widened complement-of-equivalence mechanism, detector derivation, and fixtures, but label S-1 **pair-matched/exact-folded and JOINT-PENDING/non-normative until m-10 pair approval plus the §D join co-sign**. Correct the header, §2, §9 cross-owner hook, and every closed/settled/normative assertion that collapses the join. The eventual join, not this design document, retires that status.

## Passed portions

- **R8-F1 closes.** FX-M10-R now has all-equal idempotence plus independent different-`segment_id`, different-`seq_hwm`, and different-marker conflict legs; each conflict preserves the first complete tuple. A marker-only implementation no longer passes.
- **The predicates match.** Both sides define equivalence over `{run_id, turn_id, attempt_id, marker_digest, segment_id, seq_hwm}` and conflict as any same-key non-equivalence, with first-committed tuple standing. The whole-tuple not-producible derivation remains only a detector label; totality rests on the exact complement.
- **The r12 rebase is valid.** Exact m-9 r12 `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35` is pair-approved, and the relevant §2 mechanism is present at those bytes.
- **Earlier closures remain intact.** The closed stop-reason domain, 511/512/overshoot cap fixtures, corrected 3,704,832-byte bound, restored r4 history, D-4 run-wide carriage, S-2/S-3/S-4/S-5 mechanisms, and amendment-borne Gate-2/terminal claims show no new defect in this pass.

## Gate effect

No m-10 pair approval, §D join, amendment readiness, operator-visible claim change, integrated re-lock, DESIGN lock, PLAN, T4/code, credential/provider action, E3, merge, or deploy is granted by this verdict. The separate B/E carriage has its own independent review outcome and is not blocked by this S-1 status finding.

## Verification

- Incoming DESIGN relay: `de59adf8b307f48b40739fba722dd0644112ae41583d839ddabb56f793190000`.
- Reviewed producer rev9: `bc39cc3ad5a3e979b9ed7559b7475252882b6cfc32f8f4bff2f241e7838c76d1`.
- Pair-approved m-9 r12: `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35`.
- `JOINT-PENDING` residue in rev9 is zero, while exact r12 retains the status at its operative ledger and owed-list loci; the contradiction was checked from current bytes, not inferred from relay summaries.

ACTIONS_GIT_REF: docs-workspace review relay + append-only INDEX row only; no design byte, frozen artifact, `frank/` byte, code, lock, PLAN, T4 token, amendment, join, credential, provider call, release binding, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: none — clean tree (`frank/` at `c78da3815a34480590071295c1e09bb7d53c10b6`); harness root is not a git repo
RELAY_LINT: OK — exact-file lint with the dispatch root exited 0
DISPATCH_ROOT_LINT: OK — `--relay-root master/relays/step3-relock-dag-m10` exited 0
Next requested action: m-10.planner restores the exact-folded/joint-pending status until the §D co-sign, returns fresh full bytes under `step3-relock-dag-m10`, and preserves every passed mechanism and fixture unchanged.
