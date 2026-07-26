## DESIGN-REVIEW — m-10 producer delta rev12 fixes S-5 wording but over-attributes S-4/S-5 fold state to receipt-only m-9 r12 loci

ROLE: Implementer
PHASE: DESIGN-REVIEW
AUTHORITY: review-only
DISPATCH_ID: step3-relock-dag-m10-review-r12
PARENT_DISPATCH_ID: step3-relock-dag-m10
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — one cross-owner provenance/state correction remains; no product choice is open at this seat
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-relock-dag-m10
DESIGN_DOC_SHA256: b82cf76edd0c0641fe0c9ca971b1b87ed0c00fc18c85d01a82815824c3be889f
DESIGN_RECORD_KIND: design-doc
DESIGN_REVIEW_VERDICT: must-revise
IN_REPLY_TO: master/relays/step3-relock-dag-m10/DESIGN-planner-20260724-193000.md
FROM: m-10.implementer
TO: m-10.planner
CC: master.orchestrator-planner, master.orchestrator-reviewer, operator, m-9.planner, m-9.implementer, m-1.planner
BUNDLE_ID: m-10-app-control-plane
SUBJECT: MUST-REVISE exact producer delta rev12 b82cf76e — S-5 wording closes, but m-9 r12 lines 299/473/511 prove only S-1/S-2 receipts and cannot support the claimed S-4/S-5 exact-fold state

DESIGN_REVIEW_VERDICT: must-revise

m-10.planner — I reviewed the complete producer delta rev12 at exact SHA-256 `b82cf76edd0c0641fe0c9ca971b1b87ed0c00fc18c85d01a82815824c3be889f`, the directly addressed request at `b108d532c5a72278139895596b6a5413adfa8cda0f3f0f9000750ea24ae552a1`, the prior rev11 verdict, current m-9 r12, and the exact master/pair settlement records. **MUST-REVISE.** R11-F1's wording and delimiter fixes pass, but rev12 attaches the S-4/S-5 status to evidence that says something narrower.

## M10-DAG-R12-F1 — receipt-only r12 loci do not prove S-4/S-5 exact-fold state

Rev12 §5 line 110 gives S-5 the correct non-normative status, but cites m-9 r12 `:299`/`:473`/`:511` as its support. §9 line 152 uses those same loci to infer that S-1/S-2/S-4/S-5 are all exact-folded at pair-approved r12 bytes. Exact r12 does not support that scope:

- line 299 is the joint-status clause for the **content-ready receipt frame**;
- line 473 covers only ledger items 4–5, the **disposition receipt and content-ready receipt** — S-2 and S-1;
- line 511 again names items 4–5 as exact-folded/joint-pending;
- line 472 explicitly keeps item 3, **m-10's D2 manifest wire schema (S-4), PARKED**, and line 474 says m-9 must rebase to m-10's eventual pair-approved exact hash before the join.

Therefore the r12 citations prove the S-1/S-2 receipt state, not S-4 or S-5. Calling S-4/S-5 `EXACT-FOLDED at pair-approved m-9 r12 bytes` erases the still-owed m-9 rebase/fold for S-4 and gives S-5 a receipt-only proof.

The correct sources are already present in this lane: master's `RECONCILE-orchestrator-planner-20260722-230000.md:35` and `…-20260723-001500.md:35` establish that S-4/S-5 are pair-settled, fold behind fresh reviews, and join only afterward; m-9's `DESIGN-planner-m9-20260722-235500.md:49` and m-10's `DESIGN-planner-20260722-234500.md:57-60` record the pair-axis settlement and no-normativity-before-both-approvals-plus-join rule.

**Required revision:** scope r12 `:299`/`:473`/`:511` only to S-1/S-2. For S-4, state the observable current truth: pair-settled and folded on m-10's side, but still PARKED in pair-approved m-9 r12 pending this m-10 pair-approved hash and m-9's ensuing rebase/fold before the join. For S-5, cite an exact live m-9 fold locus if one exists; otherwise bind its settled/pending-join status to the master and pair-settlement records above rather than the receipt ledger. Update §5, §9, and the header/status summary consistently. Keep S-3 separate; the incoming relay's terminal shorthand “S-1..S-5 going normative only at the join” must not re-sweep S-3 under that gate.

## Passed portions

- **R11-F1's semantic correction closes.** §5 now excludes the joint `assign` carrier from the independently normative m-10-local recipe and keeps the carrier joint-pending/non-normative.
- **The §9 Markdown defect closes.** The live delimiter is now well formed; `**.**` remains only in historical header prose describing the fix.
- **The concept sweep passes for live normativity assertions.** §2 S-1, §4 S-2, §5 S-5, and §9 all deny premature normativity. Line 30 is an explicitly frame-independent m-10-local evidence property; line 97 is the distinct S-3 C-item; line 156 is the separate B/E carriage.
- **Mechanism and fixtures still pass at the reviewed loci.** Receipt equivalence remains all-six-members-equal; conflict is any same-key non-equivalence with first tuple standing; FX-M10-R retains all-equal plus independent segment, sequence-high-water-mark, and marker mutations.
- **No new boundary regression found.** The m-1 path-carrier negatives, cap/terminal rules, D-4 run-wide carriage, and separate B/E carriage remain intact.

## Gate effect

No m-10 pair approval, §D join, amendment readiness, operator-visible claim change, integrated re-lock, DESIGN lock, PLAN, T4/code, credential/provider action, E3, merge, or deploy is granted by this verdict. The independently pair-approved B/E carriage remains unaffected.

## Verification

- Incoming DESIGN relay: `b108d532c5a72278139895596b6a5413adfa8cda0f3f0f9000750ea24ae552a1`.
- Reviewed producer rev12: `b82cf76edd0c0641fe0c9ca971b1b87ed0c00fc18c85d01a82815824c3be889f`.
- Pair-approved m-9 r12: `044229654f256eb3afe102ca8f4254bde5c97aa4de4b104d4e2a2024df858c35`.
- Master settlement records: `922b796d3ee0becd72d34184df50b1530f0b00dec273414d20013d79cb06ca9b` and `806370588cfc3088587aec2cc1a7b797ff92e046a51dad83dfa11bec2f9256b7`.
- Pair settlement records: m-9 `2b8b3189d4278832ff1fd255446af1b8b19e59287c8f2e47ea389cdb439b239c`; m-10 `7b45e4643bf8ae05c246a074424004ba19fdc75793c4a024c2209f90fa2da761`.

ACTIONS_GIT_REF: docs-workspace review relay + append-only INDEX row only; no design byte, frozen artifact, `frank/` byte, code, lock, PLAN, T4 token, amendment, join, credential, provider call, release binding, E3, merge, or deploy action
FINAL_GIT_STATUS_SHORT: none — clean tree (`frank/` at `c78da3815a34480590071295c1e09bb7d53c10b6`); harness root is not a git repo
RELAY_LINT: OK — exact-file lint with the dispatch root exited 0
DISPATCH_ROOT_LINT: OK — `--relay-root master/relays/step3-relock-dag-m10` exited 0
Next requested action: m-10.planner corrects the S-4/S-5 provenance and observable fold state, keeps S-3 separate, returns fresh full bytes under `step3-relock-dag-m10`, and preserves every passed mechanism and fixture unchanged.
