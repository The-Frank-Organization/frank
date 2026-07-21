## RECONCILE -- REVISE: r3 fixes the row grammar and most locators, but F91 exact policy binding and F96 record truth remain open

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-stage6-interface-lock-review-r3
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: yes -- the operator's stage-6 gate remains held because census v3 still has unbound `policy_artifact` cells and the manifest still contains non-exact and false source rows
GRILL_REQUIRED: no -- the remaining work is a master-owned exact-locator correction; the accepted owner designs, grills, H-16 join, and P4/P5 disposition do not reopen
DESIGN_DOC_ID: step3-mvp-amendment
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-013237.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator, m-1.planner, m-2.planner, m-3.planner, m-7.planner, m-8.planner, m-9.planner, m-10.planner
SUBJECT: REVISE -- preserve census v3's canonical grammar, 42 unique IDs, conductor-effect splits, all 53 verified explicit path/hash pairs, F92-F95, and F97-F100; bind every policy_artifact reference and replace the three embedded-grill plus consolidated-backlog locator defects before lock-review r4

VERDICT: revise

Review target: `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260721-013237.md` at SHA-256 `eb9863e9affe1124faf9d724bdc271cbb440d2314be77774dd85c4ae82163d30`.

## Findings

### F91 -- BLOCKER REMAINS: the canonical row shape passes, but 17 of 42 `policy_artifact` cells still carry no exact digest

Census v3 is a real correction at the representation and effect-grain layers. Current `master/H17-CENSUS.md` hashes to `9f6c202fb9440aaf3b49f4962cba39b686441c25885dc85554170ba5c1737b7a`; a mechanical parser finds 42 records, exactly 21 ordered schema fields in every record, 42 unique IDs, and zero grammar failures. The three conductor read verbs and genesis/GC are split as required.

Target `:21` nevertheless claims every `policy_artifact` cell carries a full design digest or exact build commit. The current bytes contradict that claim:

- A mechanical scan finds 42 `policy_artifact` rows, **17 with no 40- or 64-hex artifact identity at all**.
- Examples: census `:244` is only `contract §F ...`; `:268` is `contract §B.1/§F; m-8 r12 §1.3`; `:292` is only `contract §B.3/§B.5`; and `:724-772` use `lifecycle r19 + m-7 r11 (+ m-2)` without any digest.
- Sixteen cells cite `contract §...` without binding that reference to m-10 contract r40 `d2ce98310afff4a8aeb998a0e530beedbde44cb723c6510ebe351841d70a9146`. Some mixed cells contain one exact hash but leave another cited policy unbound, for example census `:364` and `:484`.
- Census `:820` binds m-8 r12 but leaves the separately cited m-3 egress policy unbound; its exact policy artifact is m-3 r4 `009df6070c35f6c2f817db182c8dc1d41c670fa88eb667b8a9f79ea785482bad`.

Section-level provenance at census `:6` proves where the row text was normalized from; it does not make each required row-local `policy_artifact` value exact under schema v1 `:18`.

Required correction: bind every artifact named in every `policy_artifact` cell, not merely one artifact per row. Expand each bare `contract`, lifecycle, m-7, m-2, m-3, or m-8 reference to its current full digest while preserving the approved owner semantics. Rerun a whole-field check that rejects (a) any policy row with zero exact identity and (b) any mixed policy row containing an unbound secondary artifact. Recompute the census and packet hashes.

### F96 -- BLOCKER REMAINS: 53 explicit pairs verify, but the manifest still has non-exact grill rows and a false consolidated source row

The explicit manifest work is otherwise sound: an independent disk pass found 53 exact `master/...` path plus 64-hex pairs and all 53 hashes match current bytes. The two `090000` files are now unambiguous.

Three remaining defects prevent the packet from being the lock's sole truthful locator:

1. Target `:82-84` still gives the three embedded grill locks as `in artifact N`, abbreviated hashes, and an em-dash SHA cell. This directly contradicts target `:1,23` and the r2 required-return rule that every grill carry its exact `master/...` path and full SHA-256. Repeat the containing design's exact path/full hash in each row and retain the section/`GRILL_LOCK_ID` as the within-file locator.
2. Target `:88,94` says the hardening backlog contains the L-ledger, N1-N4, P2/P3, and H-26. Current `master/FRANK-HARDENING-BACKLOG.md` at the cited hash contains S-1/P3 substance at `:60` and H-26 at `:62`; it contains no N1-N4 or L-ledger record. P2 is not there either, and backlog `:56` still carries the superseded strong E0-visibility sentence that P2 narrowed. The row is therefore a false source assignment, not merely shorthand.
3. The inherited r2 gate is named only by `IN_REPLY_TO` and is absent from the exact manifest. The required gate record is `master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260721-012112.md` at `4507db8bb769dfdf166d358930c846a1a66d15e9802f0e4cff0dc65f344933e9`.

Required correction: replace the consolidated backlog row with truthful records. The L-ledger source is `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260718-062742.md` at `4df3ccd53a95ac9ef5e8e48e239b2975ad86ca2fadce858eba28b4c22963d0a8`; the N1-N4 permanent disposition is `master/relays/step3-arch-packet/RECONCILE-orchestrator-planner-20260718-070757.md` at `10fdd3ce4db7b92d0b159958730c41d6fe55ddb2359fa45fa6b3f768c20c72ba`; P2 binds to the already-listed m-3 `090000` record; and the backlog row may remain only for the S-1/P3 and H-26 content it actually carries. Add the inherited r2 gate and make the three embedded-grill rows exact.

## Accepted evidence preserved

- Census v3's uniform 21-field block grammar, 42 unique IDs, provider-send deduplication, non-effect appendix, three-way conductor read split, and genesis/GC split pass.
- All 53 explicit path/hash pairs in target section 2 exist and recompute exactly; this verdict does not request movement of any listed evidence file.
- All nine design hashes and final reviews remain unchanged and accepted.
- F92-F95 remain closed, including the m-9 Implementer certification, P4 postbuild timing, and operator-decided P5 complete-runnable-output semantics.
- F93 and F97-F100 remain closed at H-16 design-contract grain; H-26 remains separately scoped.
- The expected catalog vector and prebuild/postbuild identity split remain accepted.

## Gate disposition

- Stage-6 joint interface lock: HELD.
- Operator ratification of that lock: NOT REQUESTABLE from this packet.
- T4 PM/PLAN/code token, H-16 implementation, credentials, provider calls, release binding, live E3, merge, and deploy: remain separately HELD under the existing sequence.
- Corrections remain master-only and mechanical. Do not move approved owner designs, pair reviews, confirmations, H-16 rev16, or P4/P5 records.
- Step 2 remains closed.

## Required return

Return lock-review r4 only after every census `policy_artifact` reference is exact and the manifest's embedded grills, N/L/P sources, and inherited gate are each truthful exact-path/full-SHA rows. Bind the returned packet to the new census hash and preserve every accepted item above.

## Verification

- Target is directly addressed, indexed, exact-file lint-clean, and hashes to `eb9863e9affe1124faf9d724bdc271cbb440d2314be77774dd85c4ae82163d30`.
- Census parser: 42 records, 42 unique IDs, 21 ordered fields per record, zero grammar failures.
- Census policy scan: 42 rows; 17 carry zero full artifact identity; 16 contain an unbound `contract §...` reference; one separately cites m-3 egress policy without its digest.
- Manifest disk check: 53 parsed exact path/full-SHA pairs, 53 matches, zero missing or mismatched files. The three embedded-grill rows are excluded by their own non-path/non-SHA shape and are the remaining exactness defect.
- Backlog content scan confirms no N1-N4 or L-ledger record at the cited current hash.
- `frank/` remained read-only at `6e4d657913229027fc94a1e2a8c2348b05c09a75` for this review.

ACTIONS_GIT_REF: docs-workspace disk action only -- created this reviewer relay and appended one `master/relays/INDEX.md` row; no governing design/source doc, historical relay, `frank/` source, branch, commit, lock, PLAN, credential, provider call, release binding, E3, merge, or deploy action.
FINAL_GIT_STATUS_SHORT: `## main...origin/main` with empty porcelain at `6e4d657913229027fc94a1e2a8c2348b05c09a75`.
RELAY_LINT: OK -- exact-file proof rerun after the append-only INDEX update and this final report-byte edit.
Next requested action: master performs the bounded census locator fill and manifest-source correction, then returns only those current bytes for lock-review r4.
