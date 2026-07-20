## RECONCILE — DIRECTLY ADDRESSED to m-8.implementer per VP F76: one bounded review over UNCHANGED m-8 r12 @ `4b670a79…` PLUS the exact r28 basis addendum (`design-m8/RECONCILE-planner-20260718-054500`) — name both hashes, verify the r27→r28 delta is disjoint from every m-8 seam, issue a review verdict; no doc edit, no re-hash unless you find a defect

ROLE: Orchestrator Planner
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-mvp-design-m8
PARENT_DISPATCH_ID: step3-arch-packet
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no — a bounded basis-addendum review over frozen bytes (VP F76); the operator gates at the Master+VP interface-lock
GRILL_REQUIRED: no
DESIGN_DOC_ID: step3-mvp-design-m8-provider-contract
IN_REPLY_TO: master/relays/step3-arch-packet/RECONCILE-orchestrator-reviewer-20260718-065204.md
FROM: master.orchestrator-planner
TO: m-8.implementer
CC: m-8.planner, m-10.planner, master.orchestrator-reviewer, operator
SUBJECT: the VP's close review (`065204` F76) accepted your r12 approval (`043932`, vs the r27 basis) and the planner's r28 letter addendum as strong consumer evidence, but the stage-2 dispatch condition ("if a consumed artifact rehashes … the review notes the delta", `design-m8/DESIGN-orchestrator-planner-20260717-010100:26`) requires the REVIEWER'S directly-addressed verdict — you were CC-only on the addendum, and CC creates no review action; this relay is that direct address

m-8.implementer — one bounded review, scope fixed by the VP:

1. **Targets (both frozen):** the m-8 provider contract r12 @ `4b670a79389b552d558555289dd2d19ad5ea42a3baed52946a1706eba65e4f51` (UNCHANGED — your `043932` approval stands; this is not a re-review of the contract) + the planner's basis addendum `master/relays/step3-mvp-design-m8/RECONCILE-planner-20260718-054500.md` binding it to m-10 r28 @ `4ffaa9ecb82a9e7e7b4b0e9ca82cf19de91247ed11e2672a55707692465fc406`.
2. **The one question:** is the m-10 r27→r28 delta (the `turn_terminal` `{terminal}` equivalence narrowing + the comparator confirmations, per m-10's approve `053100`) disjoint from every m-8-consumed locus — the addendum's six-row locus table (`connector_assign` 7-field · `attempt_open_ok`/`attempt_open_reject` · the `attempt_result` disposition set · the `REJECTED_LOCAL`/`CANCELLED` terminal rows · the cancellation consumer discipline · the §B.4 epoch authority)? Verify at the r28 bytes yourself — the addendum's claim is the thing under review, not the evidence for it.
3. **Return:** a review verdict (approve | must-revise) naming BOTH hashes, stating the delta disposition. On approve, stage 2 closes with no m-8 byte moving. On a defect: the finding routes through master and the F73 ladder applies. F72, the m-9 consumer reviews, and your exact r12 owner-byte approval are accepted and must NOT be replayed.

ACTIONS_GIT_REF: docs-workspace disk action — this routing relay + one INDEX.md row; no design doc, no `frank/` action.
FINAL_GIT_STATUS_SHORT: unavailable — cwd is not a git repo (docs workspace); `frank/` main clean at `502e06c` (`s11-close`).
Next requested action: m-8.implementer returns the directly-addressed verdict; master carries it into the corrected close supplement.
