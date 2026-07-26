## DESIGN-REVIEW — m-10 producer delta rev10 still contains two live premature-normativity assertions

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m10-review-r10
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — two stale live status assertions remain; no mechanism or product choice is open at this seat
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m10
DESIGN_DOC_SHA256: 56124f42c11402c915fa71e8f35c44650161f1dcb79deb7f1977b1a9a58c2540
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m10/DESIGN-planner-20260724-173500.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-1.planner
BUNDLE_ID: m-10-app-control-plane
SUBJECT: MUST-REVISE exact producer delta rev10 56124f42 — header and §2 restore joint-pending, but live §4 and §9 still call the §D seams normative before the required co-sign

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the complete producer delta rev10 at exact SHA-256 `56124f42c11402c915fa71e8f35c44650161f1dcb79deb7f1977b1a9a58c2540`, the directly addressed request at `db17b7fdeb1a1ecf77450a9f94e48e9c68c247ac80e4382e58829fb24fe432c4`, the prior rev9 verdict, and m-9 r12's exact joint-status, ledger, and owed-list bytes. **MUST-REVISE.** Rev10 correctly restores joint-pending in the revision header and §2, but two operative statements still collapse the same join gate.

## M10-DAG-R10-F1 — live §4 and §9 contradict the restored joint-pending status

The corrected rev10 status says S-1/S-2/S-4/S-5 remain `EXACT-FOLDED, JOINT-PENDING, NOT NORMATIVE` until both pairs approve matching bytes **and** the §D join is co-signed. Exact m-9 r12 says the same at §2 line 299, the ledger at line 473, and the owed list at line 511. But rev10 still says:

- §4 line 84: **“The report/receipt/conflict frames — SETTLED with m-9 (S-2), NORMATIVE at this revision.”** This advances S-2 before the join.
- §9 line 152: after correctly saying S-1/S-2/S-4/S-5 are exact-folded and non-normative, the same live paragraph says **S-1 through S-5 “are all SETTLED and folded normative above.”** That is a direct internal contradiction and also sweeps S-3 into a blanket status that should instead be stated separately.

These are operative design assertions, not merely quoted revision history. The earlier header text under “Carried from rev7” may remain as explicitly scoped historical lineage, but it must not control current status.

**Required revision:** change §4's S-2 label to `EXACT-FOLDED / JOINT-PENDING / NON-NORMATIVE` until m-10 pair approval plus the §D co-sign. In §9, replace the blanket S-1–S-5 normative sentence with seam-specific truth: S-1/S-2/S-4/S-5 remain exact-folded/joint-pending/non-normative until those gates; state S-3 independently according to its own actual status. Preserve the exact r12 basis, widened complement-of-equivalence mechanism, detector derivation, fixtures, and already-correct header/§2 wording.

## Passed portions

- **The rev10 header and §2 correction pass.** Both distinguish matching producer bytes from normativity and retain the m-10-approval plus §D-co-sign conjuncts.
- **The mechanism still passes.** Receipt equivalence is all six evidence members equal; conflict is any same-key non-equivalence; the first complete tuple stands. The whole-tuple not-producible statement remains only a detector derivation, not the basis of totality.
- **The mutation-resistant fixture still passes.** FX-M10-R contains all-equal idempotence plus independent different-`segment_id`, different-`seq_hwm`, and different-marker conflict legs.
- **The exact external basis is intact.** m-9 r12 remains `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35`, with the join gate present at lines 299/473/511.
- **Earlier closures remain intact.** No new defect was found in the stop-reason domain, cap boundaries/overshoot, 3,704,832-byte bound, D-4 run-wide carriage, carrier negatives, or the separate B/E carriage.

## Gate effect

No m-10 pair approval, §D join, amendment readiness, operator-visible claim change, integrated re-lock, DESIGN lock, PLAN, T4/code, credential/provider action, E3, merge, or deploy is granted by this verdict. The independently pair-approved B/E carriage remains unaffected.

## Verification

- Incoming DESIGN relay: `db17b7fdeb1a1ecf77450a9f94e48e9c68c247ac80e4382e58829fb24fe432c4`.
- Reviewed producer rev10: `56124f42c11402c915fa71e8f35c44650161f1dcb79deb7f1977b1a9a58c2540`.
- Pair-approved m-9 r12: `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35`.
- Operative contradiction reproduced from current bytes at rev10 lines 84 and 152; the required gate reproduced directly from current m-9 r12 lines 299, 473, and 511.

ACTIONS_GIT_REF: docs-workspace review relay + append-only INDEX row only; no design byte, frozen artifact, `frank/` byte, code, lock, PLAN, T4 token, amendment, join, credential, provider call, release binding, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: none — clean tree (`frank/` at `c78da3815a34480590071295c1e09bb7d53c10b6`); harness root is not a git repo
RELAY_LINT: OK — exact-file lint with the dispatch root exited 0
DISPATCH_ROOT_LINT: OK — `--relay-root master/relays/step3-relock-dag-m10` exited 0
Next requested action: m-10.planner corrects the two live status assertions, keeps S-3 separate, returns fresh full bytes under `step3-relock-dag-m10`, and preserves every passed mechanism and fixture unchanged.
