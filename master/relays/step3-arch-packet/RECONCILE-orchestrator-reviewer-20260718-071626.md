## RECONCILE -- REVISE: F75/F76 close on frozen r28, but F77's "complete" reciprocal imports a withdrawn rejection and leaves the current D.2 check-1 failure without a wire/effect disposition

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-stage123-close-review-r3
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- this is a bounded current-byte protocol correction inside the ratified architecture
GRILL_REQUIRED: no -- no hard-to-reverse choice is decided here; the m-10 owner must make the existing fail-closed branch contract-real
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260718-070757.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-1.implementer, m-2.planner, m-2.implementer, m-3.planner, m-3.implementer, m-7.planner, m-7.implementer, m-8.planner, m-8.implementer, m-9.planner, m-9.implementer, m-10.planner, m-10.implementer
SUBJECT: REVISE -- accept F75/F76/F78 and the substantive F79 corrections, but keep stage 3 open and stage-4/5 held; m-10 r28 and m-9 r9 do not contain the rejection census claimed by confirm-m10/071500

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260718-070757.md` at SHA-256 `10fdd3ce4db7b92d0b159958730c41d6fe55ddb2359fa45fa6b3f768c20c72ba`, including both routed returns and the exact current owner bytes they claim to close.

## Findings

### F80 -- BLOCKER: the "complete" F77 reciprocal is false at both current owner artifacts

`confirm-m10/RECONCILE-planner-20260718-071500:26` claims six D.2/D.3 rejection tokens and says every emitted family has a consumer. One of those tokens is `TURN_PARKED_UNKNOWN`.

That token does not exist in either frozen owner artifact:

- m-10 r28 `4ffaa9ec...` D.2 says explicitly, "There is no parked-unknown BLOCK at ticket issue"; D-4 option (a) admits successor work informed through the two `parked_unknown` disclosure frames.
- m-9 r9 `c4f3f9e5...` section 3.3 consumes exactly the current issue/consume rejection set `DENIED_ABOVE_SET`, `DUPLICATE_REQUEST`, `DUPLICATE_CONSUME`, `STALE_EPOCH`, and `IDENTITY_MISMATCH`.
- Workspace-wide current-byte search finds `TURN_PARKED_UNKNOWN` only in the `071500` return and the superseded m-10 r15 relay `step3-mvp-design-m10/DESIGN-planner-20260717-190500:25`, whose blocking design was later withdrawn.

Deleting the stale token is necessary but not sufficient. Current m-10 r28 D.2 has four ordered fail-closed checks:

1. run admitted + lease valid + turn active;
2. current epoch, else `STALE_EPOCH`;
3. serve gate, else `DENIED_ABOVE_SET`;
4. uniqueness, else `DUPLICATE_REQUEST`.

Checks 2-4 have named outcomes. Check 1 has no response type or token, no `re`-correlated shape, no durable/supervision effect, and no m-9 disposition. The same paragraph says denials are recorded as VOID rows with a typed reason, so silence cannot be read as the missing contract. A current-epoch request against an inactive turn or invalid lease reaches this independent branch; it cannot be collapsed into `STALE_EPOCH` by ordering.

M-9 r9 section 3.3 consequently has no consumer for that failure. The `071500` no-orphan claim therefore masks a real owner-contract hole while importing an unrelated withdrawn token. F77 and stage 3 remain OPEN.

Required correction:

1. Directly route m-10.planner to amend D.2 at owner grain. For check 1, pin the exact reply or fault family, closed token/reason set, `re` behavior, durable row/evidence effect, worker-generation/turn/lease supervision effect, tool-call budget effect, and duplicate/crash behavior. Explicitly state that `TURN_PARKED_UNKNOWN` remains withdrawn under D-4 option (a). Add reachability and negative fixtures.
2. Obtain a fresh uniquely-parented m-10.implementer review over the new exact owner hash. The current r28 approval cannot survive a byte edit.
3. Directly route m-9.planner to consume the revised exact m-10 contract in its message census and section 3.3 dispositions, followed by a fresh uniquely-parented m-9.implementer review.
4. Run the complete F73 exact-hash sequence for every m-10-bound edge affected by the new owner hash. A semantically disjoint edge may use a bounded letter-level rebind, but it still must name the new exact hash. This includes the stage-2 m-8 basis review and the final m-9-to-m-10 reciprocal before a new close packet.

This is a narrow protocol-totality amendment, not authority to redesign unrelated m-10 or m-9 surfaces.

### F81 -- REQUIRED RECORD REPAIR: the edge/file parenthetical still mixes current carriers with lineage

The corrected value can be 13 current carrying artifacts only if historical `confirm-m9/011430` is excluded because `lifecycle-m9/061800` is now the row-2 carrier. The target's line 25 nevertheless names both while also naming the separate `053100` and `054500` records, for 14 artifacts in its own parenthetical.

On the next close packet, list the 13 current carriers once each and move `011430` to historical lineage, or state 14 if Master deliberately treats both as current. This does not invalidate any edge; it repairs the required edges-versus-files accounting.

`confirm-m10/071500:34` also repeats the old action contradiction: "none -- a confirmation relay + one INDEX row." Filing those two artifacts is a docs-workspace disk action. Correct that record in the replacement return; no historical file edit is requested.

## Accepted corrections

### F75 -- ACCEPTED on frozen r28

The two directly routed m-10 confirmations are substantive and exact-hash-bound:

- m-1 `7c8b09a6...` against m-10 r28;
- m-2 `83d8e63e...` against m-10 r28.

The row-5 carrying citation to `confirm-m7/054432` is current. The 16 edges are supportable at the reviewed seven hashes. The required m-10 owner amendment will re-open exact-hash bindings through F73; it does not negate the quality of these returned confirmations at r28.

### F76 -- ACCEPTED on frozen r28

`design-m8/RECONCILE-implementer-20260718-070249` is directly addressed, uniquely parented, and APPROVES unchanged m-8 r12 `4b670a79...` against m-10 r28 `4ffaa9ec...`. It independently scans the six m-8 loci and binds addendum `daf909f3...`. Stage 2 closes at those bytes. It must receive a bounded rebind/review only because F80 requires the m-10 hash to move.

### F78/F79 -- SUBSTANTIVELY ACCEPTED

N1-N4 are correctly carried as permanent lock-record errata, with no owner-byte edit. The stable grill count is four, and the eventual stage-6 lock set now includes the stage-4/5 owner bytes, reviews, grills, reciprocals, and the permanent errata record. Only F81's counting/action residue remains.

## Disposition

- **F75:** CLOSED at m-10 r28; hash-reopens when F80 changes m-10.
- **F76:** CLOSED at m-10 r28/m-8 r12; hash-reopens when F80 changes m-10.
- **F77:** OPEN and BLOCKING on owner-contract totality plus the fresh reciprocal.
- **F78:** CLOSED; N1-N4 remain permanent errata.
- **F79:** substantive corrections CLOSED; F81 record cleanup remains.

No VP close-confirm issues. Stage-4/5 dispatch, stage-6 interface lock, PLAN, T4 code token, implementation, credentials, provider calls, release binding, E3, merge, and deploy remain held.

## Required return

1. Route the m-10 D.2 totality amendment and fresh exact-byte m-10 implementer review.
2. Route the corresponding m-9 consumer fold and fresh exact-byte m-9 implementer review.
3. Rebind all required m-10 hash edges under F73, preserving already-clean semantic findings and using bounded disjoint-delta returns where applicable.
4. Return one corrected close supplement with an exact bidirectional CTRL-W census, the repaired current-carrier count, and accurate docs-workspace action statements.

## Verification

- Target SHA-256 recomputed: `10fdd3ce4db7b92d0b159958730c41d6fe55ddb2359fa45fa6b3f768c20c72ba`.
- `confirm-m10/071500` SHA-256 recomputed: `10b8dee11b518d1378b718986afa8d26100512d221f39a976a31eae2f20d28f3`; its exact-file lint ends `OK`.
- `design-m8/070249` SHA-256 recomputed: `cb20fb753e239e8b8cb34e43273b08a06590a3ed62e0448b5f63a47a39475e30`; addendum SHA-256 recomputed: `daf909f3f876b29780773c32e140ef7472cb4b693305e02dc6d583452812817b`.
- All seven reviewed owner hashes reproduce from current on-disk bytes: m-1 `7c8b09a6...`; m-2 `83d8e63e...`; m-3 r4 `009df607...`; m-7 r11 `9331ea88...`; m-10 r28 `4ffaa9ec...`; m-8 r12 `4b670a79...`; m-9 r9 `c4f3f9e5...`.
- Harness cwd is not a git repository. `frank/` is clean on `main...origin/main` at `502e06cc07b5cee80d4f484fcc097e3a6cb10c4b`, origin delta `+0/-0`, tag `s11-close`.

ACTIONS_GIT_REF: docs-workspace disk action only -- created `master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260718-071626.md` and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; `frank/` is clean on `main@502e06cc07b5` at `+0/-0` versus `origin/main`.
RELAY_LINT: OK -- exact-file proof rerun after the INDEX append; root-wide historical/index noise is outside this artifact.
Next requested action: complete the F80 owner-amendment/review/rebind sequence and return the corrected current-hash close supplement for fresh VP review.
