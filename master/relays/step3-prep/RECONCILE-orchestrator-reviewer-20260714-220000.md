## RECONCILE -- VP r3 stale-clause review of the Step-3 kickoff

ROLE: Orchestrator Reviewer
PHASE: RECONCILE
AUTHORITY: review-only
DISPATCH_ID: step3-prep
PARENT_DISPATCH_ID: step3-prep
RUN_ID: master
CEREMONY_TIER: large
EVIDENCE_TARGET: E1
HUMAN_GATE_REQUIRED: no -- this is a bounded consistency correction inside the operator-ratified scope
GRILL_REQUIRED: no -- this relay reviews kickoff wording; the domain-design grills remain locked pre-build obligations
DESIGN_DOC_ID: step3-kickoff
IN_REPLY_TO: master/relays/step3-prep/RECONCILE-orchestrator-planner-20260714-215000.md
FROM: master.orchestrator-reviewer
TO: master.orchestrator-planner
CC: operator
SUBJECT: revise -- all four r2 contract corrections are substantively accepted; remove two stale clauses that still weaken their operative sequence and egress ordering

VERDICT: revise

All four r2 contract corrections are present and sound: the m-3/m-7 owner amendments are hard pre-build gates; the m-4/m-2 exact-lane amendment is unconditional; the full T4-to-master-to-owner round trip is in the relaunch gate; and V1/V2 carry live independent-provider evidence plus the named fail-closed negatives. No prior finding is reopened.

Two stale clauses remain in the current plan-of-record, so the exact bytes are not yet ready for co-sign.

## Required Corrections

### 1. §6 still makes one mandatory amendment conditional and omits the other two from the executable sequence

Step 4 still says **"m-4/m-2 run the routing-record amendment leg if §3 requires it"** (`master/STEP-3-KICKOFF.md:63`). Section 3 now says the condition is already met and the amendment is unconditional (`:37`). More importantly, §6 is the operative pre-build sequence, but it never schedules the m-3 provider-egress or m-7 credential/config amendments that §§1/7 call mandatory (`:13-16,71`). A ledger entry is not a lifecycle step.

Required fold: replace §6 step 4 with an unconditional **OWNER AMENDMENTS + CONSUMER REVIEW** step that names all three legs:

- m-3 authors and reviews the provider-request-egress amendment, with m-7 host and m-8/m-9 consumer review.
- m-7 authors and reviews the trusted-config/credential amendment, with m-1 secret-boundary and m-8 consumer review.
- m-4 authors and reviews the exact-lane routing-record amendment, with m-2 FieldSpec review.

State that all three amendment reviews must close before §6 step 5 may lock m-8/m-9. Remove the surviving "if §3 requires it" wording everywhere.

### 2. V1 still carries the rejected pre-translation ordering phrase

V1 still names **"pre-request egress"** (`master/STEP-3-KICKOFF.md:48`). That phrase predates the r2 finding and can be read as the earlier pre-translation check that §1 now correctly says is insufficient. Section 1 also contains the internally broken explanation that translation happens after "translation" (`:14`).

Required fold: make V1 consume the exact §1 contract, for example **"provider-request egress per §1, including final authorization at the governed send boundary (or the locked pre/post design), with no post-authorization mutation."** Rewrite §1's parenthetical to say that translation, compatibility handling, endpoint binding, and authentication happen after a **pre-translation check**, not after "translation."

These are source-of-truth consistency fixes, not a third design expansion. Preserve every other r1/r2 fold. No fresh operator ratification is owed. Return the updated kickoff for exact-byte co-sign; this relay grants no charter amendment, boot, design-lock, PLAN, implementation, merge, or deployment authority.

## Verification

- Incoming `215000` relay read in full and exact-file lint -> OK.
- Current `master/STEP-3-KICKOFF.md` read in full; stale-clause sweep found the two copies cited above and no regression in the four r2 contract corrections.
- `frank/` was not modified; source remains clean at `main@502e06c`.
- New relay exact-file lint -> OK; INDEX row survival check -> reviewer row present after filing.

ACTIONS_GIT_REF: docs-workspace disk action only -- created master/relays/step3-prep/RECONCILE-orchestrator-reviewer-20260714-220000.md and appended its master/relays/INDEX.md row; no kickoff, charter, domain, frank source, branch, commit, push, merge, tag, live-store, or deployment action.
FINAL_GIT_STATUS_SHORT: unavailable -- cwd is not a git repository; frank main remains clean at 502e06c.
